#!/bin/bash
# =============================================================================
# Field Encryption and Decryption Test Script
# =============================================================================
# This script tests the complete encryption/decryption workflow:
#   1. Get encryption public key
#   2. Encrypt model using public key (compare before/after)
#   3. Decrypt model using server-side key (compare before/after)
#
# Usage:
#   ./scripts/test-field-encryption-and-decryption.sh
#   ./scripts/test-field-encryption-and-decryption.sh http://localhost:8056
# =============================================================================

set -e

# Configuration
BASE_URL="${1:-http://localhost:8056}/beetle"
AUTH="default:default"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Helper functions
print_header() {
    echo -e "\n${BLUE}================================================================${NC}"
    echo -e "${BLUE}$1${NC}"
    echo -e "${BLUE}================================================================${NC}"
}

print_subheader() {
    echo -e "\n${CYAN}--- $1 ---${NC}"
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

print_error() {
    echo -e "${RED}✗ $1${NC}"
}

# Check if jq is available
if ! command -v jq &> /dev/null; then
    print_error "jq is required but not installed. Please install jq first."
    exit 1
fi

echo -e "${GREEN}"
echo "╔═══════════════════════════════════════════════════════════════╗"
echo "║         Encryption Workflow Test Script                       ║"
echo "╚═══════════════════════════════════════════════════════════════╝"
echo -e "${NC}"
echo "Base URL: $BASE_URL"
echo "Auth: $AUTH"

# =============================================================================
# Test Data - Plaintext Model with Sensitive Fields
# =============================================================================
PLAINTEXT_MODEL='{
  "source": {
    "storageType": "filesystem",
    "path": "/data/source",
    "filesystem": {
      "ssh": {
        "host": "192.168.1.100",
        "port": 22,
        "user": "ubuntu",
        "privateKey": "-----BEGIN OPENSSH PRIVATE KEY-----\nMIIEpAIBAAKCAQEAtest1234567890abcdef\n-----END OPENSSH PRIVATE KEY-----"
      }
    }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "my-bucket/backup",
    "objectStorage": {
      "minio": {
        "endpoint": "https://minio.example.com",
        "accessKeyId": "AKIAIOSFODNN7EXAMPLE",
        "secretAccessKey": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
      }
    }
  }
}'

# =============================================================================
# Step 1: Get Encryption Public Key
# =============================================================================
print_header "Step 1: Get Encryption Public Key"

KEY_RESPONSE=$(curl -s -X GET "$BASE_URL/migration/data/encryptionKey" \
  -u "$AUTH" \
  -H "Content-Type: application/json")

# Check if request was successful
if echo "$KEY_RESPONSE" | jq -e '.success' > /dev/null 2>&1; then
    SUCCESS=$(echo "$KEY_RESPONSE" | jq -r '.success')
    if [ "$SUCCESS" = "true" ]; then
        print_success "Encryption key generated successfully"
        
        KEY_ID=$(echo "$KEY_RESPONSE" | jq -r '.data.keyId')
        ALGORITHM=$(echo "$KEY_RESPONSE" | jq -r '.data.algorithm')
        EXPIRES_AT=$(echo "$KEY_RESPONSE" | jq -r '.data.expiresAt')
        
        echo ""
        echo "Key ID:     $KEY_ID"
        echo "Algorithm:  $ALGORITHM"
        echo "Expires At: $EXPIRES_AT"
        echo ""
        echo "Public Key (first 100 chars):"
        echo "$KEY_RESPONSE" | jq -r '.data.publicKey' | head -c 100
        echo "..."
    else
        print_error "Failed to get encryption key"
        echo "$KEY_RESPONSE" | jq .
        exit 1
    fi
else
    print_error "Invalid response from server"
    echo "$KEY_RESPONSE"
    exit 1
fi

# =============================================================================
# Step 2: Encrypt Model
# =============================================================================
print_header "Step 2: Encrypt Model (Test Encryption)"

print_subheader "BEFORE Encryption - Plaintext Model"
echo "Sensitive fields are visible in plaintext:"
echo ""
echo "$PLAINTEXT_MODEL" | jq .

# Build encrypt request
ENCRYPT_REQUEST=$(jq -n \
  --argjson bundle "$(echo "$KEY_RESPONSE" | jq '.data')" \
  --argjson model "$PLAINTEXT_MODEL" \
  '{publicKeyBundle: $bundle, model: $model}')

# Send encryption request
ENCRYPT_RESPONSE=$(curl -s -X POST "$BASE_URL/migration/data/test/encrypt" \
  -u "$AUTH" \
  -H "Content-Type: application/json" \
  -d "$ENCRYPT_REQUEST")

print_subheader "AFTER Encryption - Encrypted Model"

if echo "$ENCRYPT_RESPONSE" | jq -e '.success' > /dev/null 2>&1; then
    SUCCESS=$(echo "$ENCRYPT_RESPONSE" | jq -r '.success')
    if [ "$SUCCESS" = "true" ]; then
        print_success "Model encrypted successfully"
        echo ""
        echo "Sensitive fields are now encrypted (base64-encoded ciphertext):"
        echo ""
        
        # Store encrypted model for next step
        ENCRYPTED_MODEL=$(echo "$ENCRYPT_RESPONSE" | jq '.data')
        echo "$ENCRYPTED_MODEL" | jq .
    else
        print_error "Encryption failed"
        echo "$ENCRYPT_RESPONSE" | jq .
        exit 1
    fi
else
    print_error "Invalid response from server"
    echo "$ENCRYPT_RESPONSE"
    exit 1
fi

# =============================================================================
# Step 3: Decrypt Model
# =============================================================================
print_header "Step 3: Decrypt Model (Test Decryption)"

print_subheader "INPUT - Using encrypted model from Step 2"
echo "The encrypted model from Step 2 will be sent to the decryption API."
echo ""

# Send decryption request
DECRYPT_RESPONSE=$(curl -s -X POST "$BASE_URL/migration/data/test/decrypt" \
  -u "$AUTH" \
  -H "Content-Type: application/json" \
  -d "$ENCRYPTED_MODEL")

print_subheader "AFTER Decryption - Plaintext Model Restored"

if echo "$DECRYPT_RESPONSE" | jq -e '.success' > /dev/null 2>&1; then
    SUCCESS=$(echo "$DECRYPT_RESPONSE" | jq -r '.success')
    if [ "$SUCCESS" = "true" ]; then
        print_success "Model decrypted successfully"
        echo ""
        echo "Sensitive fields are restored to plaintext:"
        echo ""
        
        DECRYPTED_MODEL=$(echo "$DECRYPT_RESPONSE" | jq '.data')
        echo "$DECRYPTED_MODEL" | jq .
        
        # Extract values for verification
        DECRYPTED_SSH=$(echo "$DECRYPTED_MODEL" | jq -r '.source.filesystem.ssh.privateKey // "N/A"')
        DECRYPTED_ACCESS=$(echo "$DECRYPTED_MODEL" | jq -r '.destination.objectStorage.minio.accessKeyId // "N/A"')
        DECRYPTED_SECRET=$(echo "$DECRYPTED_MODEL" | jq -r '.destination.objectStorage.minio.secretAccessKey // "N/A"')
        
        # Verify decryption
        print_subheader "Verification - Compare Original vs Decrypted"
        
        ORIGINAL_SSH=$(echo "$PLAINTEXT_MODEL" | jq -r '.source.filesystem.ssh.privateKey')
        ORIGINAL_ACCESS=$(echo "$PLAINTEXT_MODEL" | jq -r '.destination.objectStorage.minio.accessKeyId')
        ORIGINAL_SECRET=$(echo "$PLAINTEXT_MODEL" | jq -r '.destination.objectStorage.minio.secretAccessKey')
        
        MATCH=true
        
        if [ "$ORIGINAL_SSH" = "$DECRYPTED_SSH" ]; then
            print_success "SSH privateKey: MATCH"
        else
            print_error "SSH privateKey: MISMATCH"
            MATCH=false
        fi
        
        if [ "$ORIGINAL_ACCESS" = "$DECRYPTED_ACCESS" ]; then
            print_success "Minio accessKeyId: MATCH"
        else
            print_error "Minio accessKeyId: MISMATCH"
            MATCH=false
        fi
        
        if [ "$ORIGINAL_SECRET" = "$DECRYPTED_SECRET" ]; then
            print_success "Minio secretAccessKey: MATCH"
        else
            print_error "Minio secretAccessKey: MISMATCH"
            MATCH=false
        fi
        
    else
        print_error "Decryption failed"
        echo "$DECRYPT_RESPONSE" | jq .
        exit 1
    fi
else
    print_error "Invalid response from server"
    echo "$DECRYPT_RESPONSE"
    exit 1
fi

# =============================================================================
# Summary
# =============================================================================
print_header "Test Summary"

if [ "$MATCH" = true ]; then
    echo -e "${GREEN}"
    echo "╔═══════════════════════════════════════════════════════════════╗"
    echo "║  ✓ ALL TESTS PASSED                                           ║"
    echo "║                                                               ║"
    echo "║  The encryption workflow is working correctly:                ║"
    echo "║  1. Public key generation: OK                                 ║"
    echo "║  2. Encryption (client-side simulation): OK                   ║"
    echo "║  3. Decryption (server-side): OK                              ║"
    echo "║  4. Data integrity verification: OK                           ║"
    echo "╚═══════════════════════════════════════════════════════════════╝"
    echo -e "${NC}"
else
    echo -e "${RED}"
    echo "╔═══════════════════════════════════════════════════════════════╗"
    echo "║  ✗ SOME TESTS FAILED                                          ║"
    echo "║                                                               ║"
    echo "║  Please check the error messages above.                       ║"
    echo "╚═══════════════════════════════════════════════════════════════╝"
    echo -e "${NC}"
    exit 1
fi
