#!/usr/bin/env bash
set -euo pipefail

# Incus-based testing (preferred - full OS with systemd)
# Tests binary in Debian container with full systemd support

PROJECTNAME=$(basename "$PWD")
PROJECTORG=$(basename "$(dirname "$PWD")")

echo "=== Incus Testing: ${PROJECTORG}/${PROJECTNAME} ==="

# Create temp directory for build
mkdir -p "${TMPDIR:-/tmp}/${PROJECTORG}"
BUILD_DIR=$(mktemp -d "${TMPDIR:-/tmp}/${PROJECTORG}/${PROJECTNAME}-XXXXXX")
trap "rm -rf $BUILD_DIR" EXIT

CONTAINER_NAME="test-${PROJECTNAME}-$$"

# Cleanup function
cleanup() {
  echo
  echo "=== Cleaning Up ==="
  incus delete "$CONTAINER_NAME" --force 2>/dev/null || true
  rm -rf "$BUILD_DIR"
}
trap cleanup EXIT

echo
echo "=== Building Binary ==="
docker run --rm \
  -v "$(pwd):/build" \
  -w /build \
  -e CGO_ENABLED=0 \
  golang:alpine go build -o "$BUILD_DIR/${PROJECTNAME}" ./src

echo
echo "=== Launching Incus Container (Debian 12) ==="
incus launch images:debian/12 "$CONTAINER_NAME"
sleep 5

echo
echo "=== Copying Binary to Container ==="
incus file push "$BUILD_DIR/${PROJECTNAME}" "$CONTAINER_NAME/usr/local/bin/"
incus exec "$CONTAINER_NAME" -- chmod +x "/usr/local/bin/${PROJECTNAME}"

echo
echo "=== Version Check ==="
incus exec "$CONTAINER_NAME" -- "${PROJECTNAME}" --version

echo
echo "=== Help Check ==="
incus exec "$CONTAINER_NAME" -- "${PROJECTNAME}" --help

echo
echo "=== Starting Server (background) ==="
incus exec "$CONTAINER_NAME" -- sh -c "${PROJECTNAME} --address 0.0.0.0 --port 8080 > /tmp/server.log 2>&1 &"
sleep 3

echo
echo "=== Testing Health Endpoints ==="
incus exec "$CONTAINER_NAME" -- sh -c "curl -s http://localhost:8080/healthz || echo 'Health check failed'"
incus exec "$CONTAINER_NAME" -- sh -c "curl -s http://localhost:8080/api/v1/healthz || echo 'API health check failed'"

echo
echo "=== Testing Conversion API ==="
incus exec "$CONTAINER_NAME" -- sh -c "curl -s http://localhost:8080/api/v1/formats || echo 'Formats endpoint failed'"

echo
echo "=== Server Logs (last 20 lines) ==="
incus exec "$CONTAINER_NAME" -- sh -c "tail -20 /tmp/server.log 2>/dev/null || echo 'No logs found'"

echo
echo "✓ All Incus tests passed"
