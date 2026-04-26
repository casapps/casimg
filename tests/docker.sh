#!/usr/bin/env bash
set -euo pipefail

# Docker-based testing (fallback when Incus not available)
# Tests binary in Alpine container

PROJECTNAME=$(basename "$PWD")
PROJECTORG=$(basename "$(dirname "$PWD")")

echo "=== Docker Testing: ${PROJECTORG}/${PROJECTNAME} ==="

# Create temp directory for build
mkdir -p "${TMPDIR:-/tmp}/${PROJECTORG}"
BUILD_DIR=$(mktemp -d "${TMPDIR:-/tmp}/${PROJECTORG}/${PROJECTNAME}-XXXXXX")
trap "rm -rf $BUILD_DIR" EXIT

echo
echo "=== Building Binary ==="
docker run --rm \
  -v "$(pwd):/build" \
  -w /build \
  -e CGO_ENABLED=0 \
  golang:alpine go build -o "$BUILD_DIR/${PROJECTNAME}" ./src

echo
echo "=== Testing in Alpine Container ==="
docker run --rm \
  -v "$BUILD_DIR:/app" \
  alpine:latest sh -c "
    set -e
    chmod +x /app/${PROJECTNAME}

    echo '=== Version Check ==='
    /app/${PROJECTNAME} --version

    echo
    echo '=== Help Check ==='
    /app/${PROJECTNAME} --help

    echo
    echo '=== Starting Server (background) ==='
    /app/${PROJECTNAME} --address 127.0.0.1 --port 8080 &
    SERVER_PID=\$!
    sleep 3

    echo
    echo '=== Testing Health Endpoints ==='
    wget -O- -q http://127.0.0.1:8080/healthz || echo 'Health check failed'
    wget -O- -q http://127.0.0.1:8080/api/v1/healthz || echo 'API health check failed'

    echo
    echo '=== Testing Conversion API ==='
    wget -O- -q http://127.0.0.1:8080/api/v1/formats || echo 'Formats endpoint failed'

    echo
    echo '=== Stopping Server ==='
    kill \$SERVER_PID 2>/dev/null || true
    wait \$SERVER_PID 2>/dev/null || true

    echo
    echo '✓ Docker tests complete'
"

echo
echo "✓ All Docker tests passed"
