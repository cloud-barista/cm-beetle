#!/bin/bash

# Object Storage Migration Script for transx
# Supports: MinIO (direct S3), Spider (presigned URL), Tumblebug (multi-cloud abstraction)

set -e

# Color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

print_status() { echo -e "${BLUE}[INFO]${NC} $1"; }
print_success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
print_warning() { echo -e "${YELLOW}[WARNING]${NC} $1"; }
print_error() { echo -e "${RED}[ERROR]${NC} $1"; }

show_usage() {
    cat << EOF
Usage: $0 [OPTIONS]

Object Storage Migration with transx

Options:
  -c, --config FILE    Migration configuration file (required)
  -v, --verbose        Enable verbose logging
  -s, --step STEP      Execute specific step: 'backup', 'transfer', 'restore'
  -h, --help           Show this help message

Available Configurations:

  Upload/Download (Local ↔ Object Storage):
    config-minio-upload.json       MinIO: Local → S3
    config-minio-download.json     MinIO: S3 → Local
    config-spider-upload.json      Spider: Local → S3 (via CB-Spider)
    config-spider-download.json    Spider: S3 → Local (via CB-Spider)

  Object Storage to Object Storage:
    config-tumblebug-os2os.json         Tumblebug: S3 → S3 (via CB-Tumblebug)
    config-tumblebug-os2os-filter.json  Tumblebug: S3 → S3 with filtering

Examples:
  # Upload test data to source Object Storage (using MinIO)
  $0 -c config-minio-upload.json -v

  # Transfer between Object Storages (using Tumblebug)
  $0 -c config-tumblebug-os2os.json -v

  # Transfer with file filtering
  $0 -c config-tumblebug-os2os-filter.json -v

  # Download from Object Storage
  $0 -c config-spider-download.json -v

EOF
}

# Default values
CONFIG=""
VERBOSE=""
STEP_ARG=""

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -c|--config)
            CONFIG="$2"
            shift 2
            ;;
        -v|--verbose)
            VERBOSE="-verbose"
            shift
            ;;
        -s|--step)
            case "$2" in
                backup)   STEP_ARG="-backup" ;;
                transfer) STEP_ARG="-transfer" ;;
                restore)  STEP_ARG="-restore" ;;
                *)
                    print_error "Invalid step: $2 (use: backup, transfer, restore)"
                    exit 1
                    ;;
            esac
            shift 2
            ;;
        -h|--help)
            show_usage
            exit 0
            ;;
        *)
            print_error "Unknown option: $1"
            show_usage
            exit 1
            ;;
    esac
done

# Validate config
if [[ -z "$CONFIG" ]]; then
    print_error "Configuration file is required. Use -c <config-file>"
    show_usage
    exit 1
fi

if [[ ! -f "$CONFIG" ]]; then
    print_error "Configuration file not found: $CONFIG"
    exit 1
fi

# Get script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Build if needed
if [[ ! -f "$SCRIPT_DIR/main" ]] || [[ "$SCRIPT_DIR/main.go" -nt "$SCRIPT_DIR/main" ]]; then
    print_status "Building migration tool..."
    cd "$SCRIPT_DIR"
    go build -o main main.go
    print_success "Build complete"
fi

# Run migration
print_status "Starting migration with config: $CONFIG"
cd "$SCRIPT_DIR"

CMD_ARGS=("-config" "$CONFIG")
[[ -n "$VERBOSE" ]] && CMD_ARGS+=("$VERBOSE")
[[ -n "$STEP_ARG" ]] && CMD_ARGS+=("$STEP_ARG")

./main "${CMD_ARGS[@]}"

print_success "Migration completed"
