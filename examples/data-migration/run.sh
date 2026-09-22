#!/usr/bin/env bash
# ==============================================================================
# CM-Beetle: Data Migration (Reference Example Script)
#
# NOTE: This script is provided strictly as an illustrative reference example
# demonstrating client-side credential encryption and asynchronous job dispatch.
# To learn and follow the step-by-step workflow, refer to the individual curl
# commands documented in README.md.
# ==============================================================================
set -euo pipefail

BEETLE_URL="${BEETLE_URL:-http://localhost:8056/beetle}"
AUTH="${BEETLE_AUTH:-default:default}"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ENCRYPT="${ENCRYPT:-true}" # Default to true for standard security workflow
PLAINTEXT_FILE="${1:-${SCRIPT_DIR}/data-migration-request.json}"
PAYLOAD_FILE="${PLAINTEXT_FILE}"

if [ "${ENCRYPT}" = "true" ]; then
  echo "=== [1/3] Encrypting sensitive credentials with one-time RSA public key ==="
  KEY_BUNDLE=$(curl -s -X GET "${BEETLE_URL}/migration/data/encryptionKey" -u "${AUTH}" | jq '.data')
  PLAIN_MODEL=$(cat "${PLAINTEXT_FILE}")

  ENCRYPTED_FILE="${SCRIPT_DIR}/data-migration-request-encrypted.json"
  jq -n --argjson kb "${KEY_BUNDLE}" --argjson m "${PLAIN_MODEL}" \
    '{publicKeyBundle: $kb, model: $m}' | \
    curl -s -X POST "${BEETLE_URL}/migration/data/test/encrypt" \
      -H "Content-Type: application/json" \
      -u "${AUTH}" \
      -d @- | jq '.data' > "${ENCRYPTED_FILE}"

  PAYLOAD_FILE="${ENCRYPTED_FILE}"
  echo "Payload encrypted successfully -> ${PAYLOAD_FILE}"
  echo ""
fi

echo "=== [2/3] Triggering Data Migration Asynchronously (Prefer: respond-async) ==="
echo "Using payload file: ${PAYLOAD_FILE}"
RES=$(curl -s -X POST "${BEETLE_URL}/migration/data" \
  -H "Content-Type: application/json" \
  -H "Prefer: respond-async" \
  -u "${AUTH}" \
  -d @"${PAYLOAD_FILE}")

echo "${RES}" | jq .
REQ_ID=$(echo "${RES}" | jq -r '.data.reqId // empty')

if [ -n "${REQ_ID}" ]; then
  echo ""
  echo "=== [3/3] Tracking Migration Job Status (reqId: ${REQ_ID}) ==="
  echo "Waiting 3 seconds for initial worker progress..."
  sleep 3
  curl -s -X GET "${BEETLE_URL}/request/${REQ_ID}" \
    -u "${AUTH}" | jq .
  echo ""
  echo "You can check status at any time with:"
  echo "curl -s -X GET \"${BEETLE_URL}/request/${REQ_ID}\" -u \"${AUTH}\" | jq ."
fi

