#!/usr/bin/env bash
set -euo pipefail

# Auto-detect test environment and run appropriate tests
# This script checks for Incus/LXD availability and falls back to Docker

PROJECTNAME=$(basename "$PWD")
PROJECTORG=$(basename "$(dirname "$PWD")")

echo "=== ${PROJECTNAME} Test Runner ==="
echo "Project: ${PROJECTORG}/${PROJECTNAME}"
echo

# Check which container runtime is available
if command -v incus &>/dev/null; then
  echo "✓ Incus detected - running full OS tests (preferred)"
  exec "$(dirname "$0")/incus.sh"
elif command -v lxc &>/dev/null; then
  echo "✓ LXD detected - running full OS tests"
  echo "Note: Using LXD (replace with Incus when available)"
  exec "$(dirname "$0")/incus.sh"
elif command -v docker &>/dev/null; then
  echo "⚠ Docker detected - running container tests (fallback)"
  echo "Note: Incus preferred for full systemd testing"
  exec "$(dirname "$0")/docker.sh"
else
  echo "✗ ERROR: No container runtime found"
  echo "Please install: incus (preferred) or docker (fallback)"
  exit 1
fi
