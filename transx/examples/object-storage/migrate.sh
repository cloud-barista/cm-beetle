#!/bin/bash

# Object Storage Migration Script for transx
# This script demonstrates data migration using Object Storage (CB-Spider compatible) with the transx library.

set -e

# Color codes for output formatting
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Default configuration
DEFAULT_CONFIG="config-basic-direct.json"
DEFAULT_MODE="direct"

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to show usage
show_usage() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Object Storage Migration with transx"
    echo ""
    echo "Options:"
    echo "  -m, --mode MODE           Migration mode:"
    echo "                              direct: local-to-remote, remote-to-local"  
    echo "                              relay: object-storage-to-object-storage, rsync-to-object-storage"
    echo "  -c, --config FILE         Migration configuration file"
    echo "  -v, --verbose             Enable verbose logging"
    echo "  -h, --help                Show this help message"
    echo ""
    echo "Pre-defined configurations:"
    echo "  config-basic-direct.json                 # Basic remote to local (default)"
    echo "  config-local-to-objectstorage.json      # Local to Object Storage upload"
    echo "  config-objectstorage-to-local.json      # Object Storage to local download"
    echo "  config-basic-relay.json                  # Basic relay mode"
    echo "  config-objectstorage-to-objectstorage.json  # Object Storage to Object Storage"
    echo "  config-rsync-to-objectstorage.json      # Rsync to Object Storage"
    echo "  config-spider-upload.json                # Upload via Spider (presigned URL)"
    echo "  config-spider-download.json              # Download via Spider (presigned URL)"
    echo "  config-minio-upload.json                 # Upload via MinIO SDK (direct S3)"
    echo "  config-minio-download.json               # Download via MinIO SDK (direct S3)"
    echo "  config-rsync-upload.json                 # Upload via rsync (local to remote)"
    echo "  config-rsync-download.json               # Download via rsync (remote to local)"
    echo ""
    echo "Step-by-step execution:"
    echo "  -s, --step STEP           Execute specific step: 'backup', 'transfer', 'restore'"
    echo ""
    echo "Examples:"
    echo "  $0                                               # Default: basic direct mode"
    echo "  $0 -c config-local-to-objectstorage.json -v     # Upload to Object Storage"
    echo "  $0 -c config-objectstorage-to-local.json -v     # Download from Object Storage"  
    echo "  $0 -c config-objectstorage-to-objectstorage.json -v # Object Storage migration"
    echo "  $0 -c config-rsync-to-objectstorage.json -v     # Rsync to Object Storage"
    echo "  $0 -c config-spider-upload.json -v               # Upload via Spider"
    echo "  $0 -c config-spider-download.json -v             # Download via Spider"
    echo "  $0 -c config-minio-upload.json -v                # Upload via MinIO SDK"
    echo "  $0 -c config-minio-download.json -v              # Download via MinIO SDK"
    echo "  $0 -c config-rsync-upload.json -v                # Upload via rsync"
    echo "  $0 -c config-rsync-download.json -v              # Download via rsync"
    echo "  $0 -s backup                                     # Execute only backup step"
    echo "  $0 -s transfer                                   # Execute only transfer step"
    echo "  $0 -s restore                                    # Execute only restore step"
    echo ""
}

# Parse command line arguments
MODE="$DEFAULT_MODE"
CONFIG=""
VERBOSE=""
STEP=""

while [[ $# -gt 0 ]]; do
    case $1 in
        -m|--mode)
            MODE="$2"
            shift 2
            ;;
        -c|--config)
            CONFIG="$2"
            shift 2
            ;;
        -v|--verbose)
            VERBOSE="-verbose"
            shift
            ;;
        -s|--step)
            STEP="$2"
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

# Determine configuration file
if [[ -z "$CONFIG" ]]; then
    if [[ "$MODE" == "relay" ]]; then
        CONFIG="config-basic-relay.json"
    else
        CONFIG="$DEFAULT_CONFIG"
    fi
fi

# Validate configuration file exists
if [[ ! -f "$CONFIG" ]]; then
    print_error "Configuration file not found: $CONFIG"
    exit 1
fi

# Build command line arguments for the Go program
CMD_ARGS=("-config" "$CONFIG")

if [[ -n "$VERBOSE" ]]; then
    CMD_ARGS+=("$VERBOSE")
fi

# Add step-specific arguments
case "$STEP" in
    backup)
        CMD_ARGS+=("-backup")
        ;;
    transfer)
        CMD_ARGS+=("-transfer")
        ;;
    restore)
        CMD_ARGS+=("-restore")
        ;;
    "")
        # No specific step, run complete migration
        ;;
    *)
        print_error "Invalid step: $STEP. Valid steps are: backup, transfer, restore"
        exit 1
        ;;
esac

print_status "Starting Object Storage Migration"
print_status "Mode: $MODE"
print_status "Configuration: $CONFIG"

if [[ -n "$STEP" ]]; then
    print_status "Step: $STEP"
fi

# Always build the binary
print_status "Building migration binary..."
go build -o main . || {
    print_error "Failed to build migration binary"
    exit 1
}

# Create required directories
print_status "Creating required directories..."
mkdir -p /tmp/object-storage-migration
mkdir -p /tmp/minio-download-test

# Test environment setup
if [[ "$CONFIG" == *"minio"* ]] || [[ "$CONFIG" == *"spider"* ]] || [[ "$CONFIG" == *"rsync"* ]]; then
    print_status "Setting up test environment..."
    
    # Create test data if uploading
    if [[ "$CONFIG" == *"upload"* ]]; then
        if [[ ! -d "/tmp/minio-test-data" ]] || [[ -z "$(ls -A /tmp/minio-test-data 2>/dev/null)" ]]; then
            print_warning "Test data not found. Please ensure test files exist in /tmp/minio-test-data/"
            print_status "Expected files: sample-database.sql, application.conf, test-data.json, README.txt"
        else
            print_status "Found test data:"
            ls -lh /tmp/minio-test-data/
        fi
    fi
    
    # Prepare download directory
    if [[ "$CONFIG" == *"download"* ]]; then
        if [[ "$CONFIG" == *"rsync"* ]]; then
            print_status "Preparing download directory: /tmp/rsync-download-test/"
            rm -rf /tmp/rsync-download-test/*
            mkdir -p /tmp/rsync-download-test/
        else
            print_status "Preparing download directory: /tmp/minio-download-test/"
            rm -rf /tmp/minio-download-test/*
        fi
    fi
fi

# Execute the migration
print_status "Executing migration..."
START_TIME=$(date +%s)

if ./main "${CMD_ARGS[@]}"; then
    END_TIME=$(date +%s)
    DURATION=$((END_TIME - START_TIME))
    print_success "Migration completed successfully in ${DURATION} seconds"
    
    # Post-migration actions
    if [[ "$CONFIG" == *"minio-download"* ]] || [[ "$CONFIG" == *"spider-download"* ]]; then
        print_status "Downloaded files:"
        echo "----------------------------------------"
        ls -lh /tmp/minio-download-test/ || echo "No files found"
        echo "----------------------------------------"
        
        # Show sample file content if available
        if [[ -f "/tmp/minio-download-test/README.txt" ]]; then
            print_status "Sample content (README.txt):"
            echo "----------------------------------------"
            head -n 20 /tmp/minio-download-test/README.txt
            echo "----------------------------------------"
        fi
    fi
    
    if [[ "$CONFIG" == *"rsync-download"* ]]; then
        print_status "Downloaded files (rsync):"
        echo "----------------------------------------"
        ls -lh /tmp/rsync-download-test/ || echo "No files found"
        echo "----------------------------------------"
        
        # Show sample file content if available
        if [[ -f "/tmp/rsync-download-test/README.txt" ]]; then
            print_status "Sample content (README.txt):"
            echo "----------------------------------------"
            head -n 20 /tmp/rsync-download-test/README.txt
            echo "----------------------------------------"
        fi
    fi
    
    # Show results if it's a download operation
    if [[ -f "/tmp/object-storage-migration/data.sql" ]]; then
        print_status "Downloaded file content:"
        echo "----------------------------------------"
        cat /tmp/object-storage-migration/data.sql
        echo "----------------------------------------"
    fi
else
    print_error "Migration failed"
    exit 1
fi
