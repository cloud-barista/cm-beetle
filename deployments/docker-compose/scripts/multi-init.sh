#!/bin/bash

# A wrapper to run initialization scripts with a single password prompt

SCRIPT_DIR=$(cd $(dirname "$0") && pwd)

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "CM-Beetle (with CB-Tumblebug) Initialization"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
read -s -p "Enter the password for credentials.yaml.enc: " MULTI_INIT_PWD
echo ""
export MULTI_INIT_PWD

# 1. OpenBao
OPENBAO_SH="$SCRIPT_DIR/../openbao/openbao-register-creds.sh"
if [ ! -f "$OPENBAO_SH" ]; then
    echo "Error: Cannot find openbao-register-creds.sh"
    exit 1
fi

echo ""
echo "Step 1. Registering credentials to OpenBao..."
chmod +x "$OPENBAO_SH" 2>/dev/null || true
bash "$OPENBAO_SH"
if [ $? -ne 0 ]; then exit 1; fi

# 2. Tumblebug
TB_SH="$SCRIPT_DIR/../cb-tumblebug/init/init.sh"
if [ ! -f "$TB_SH" ]; then
    echo "Error: Cannot find init.sh"
    exit 1
fi

echo ""
echo "Step 2. Registering credentials to Tumblebug..."
chmod +x "$TB_SH" 2>/dev/null || true
bash "$TB_SH"
if [ $? -ne 0 ]; then exit 1; fi

echo ""
echo "Initialization completed successfully."
