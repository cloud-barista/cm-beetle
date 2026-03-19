#!/bin/bash
# ==============================================================================
# cleanDB.sh — Clean all metadata and persistent data (excluding OpenBao)
# ==============================================================================
set -euo pipefail

# Resolve project root (parent of this script's directory)
SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
DEPLOY_DIR=$(cd "$SCRIPT_DIR/.." && pwd)

RED='\033[0;31m'
LGREEN='\033[1;32m'
NC='\033[0m' # No Color

DATA_DIR="${DEPLOY_DIR}/data"

echo -e "=========================================================="
echo -e "[Info]"
echo -e "=========================================================="
echo -e "This script will ${RED}REMOVE PERSISTENT DATA${NC} in the local deployment."
echo -e "Target directory: ${LGREEN}${DATA_DIR}${NC} (Excluding OpenBao data)"
echo ""

if [ ! -d "$DATA_DIR" ]; then
    echo -e "Data directory not found. Nothing to clean."
    exit 0
fi

# Identify subdirectories to remove (excluding openbao-data)
CLEAN_TARGETS=$(ls -A1 "$DATA_DIR" 2>/dev/null | grep -v "^openbao-data$" || true)

if [ -z "$CLEAN_TARGETS" ]; then
    echo -e "No data found to clean (excluding openbao-data)."
    exit 0
fi

echo -e "The following service data will be REMOVED:"
echo "$CLEAN_TARGETS" | sed 's/^/  - /'
echo ""

while true; do
    read -p 'Do you want to proceed? (y/n) : ' CHECKPROCEED
    case $CHECKPROCEED in
    [Yy]*)
        break
        ;;
    [Nn]*)
        echo -e "\nCleanup cancelled."
        exit 1
        ;;
    *)
        echo "Please answer yes or no."
        ;;
    esac
done

echo -e "\nCleaning up data..."
cd "$DATA_DIR"
for target in $CLEAN_TARGETS; do
    sudo rm -rf "$target"
done

echo -e "${LGREEN}Done!${NC} Selected persistent data has been removed."
