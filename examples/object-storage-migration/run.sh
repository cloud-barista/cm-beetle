#!/usr/bin/env bash
# ==============================================================================
# CM-Beetle: Object Storage Recommendation & Migration (Reference Example Script)
#
# NOTE: This script is provided strictly as an illustrative reference example
# of automating CM-Beetle REST API calls. To learn and follow the step-by-step
# workflow, refer to the individual curl commands documented in README.md.
# ==============================================================================
set -euo pipefail

BEETLE_URL="${BEETLE_URL:-http://localhost:8056/beetle}"
AUTH="${BEETLE_AUTH:-default:default}"
NS_ID="${NS_ID:-mig01}"
NAME_SEED="${NAME_SEED:-test}"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "=== [1] Recommend Object Storage for target cloud ==="
curl -s -X POST "${BEETLE_URL}/recommendation/middleware/objectStorage" \
  -H "Content-Type: application/json" \
  -u "${AUTH}" \
  -d @"${SCRIPT_DIR}/object-storage-recommendation-request.json" | jq '.data' > "${SCRIPT_DIR}/object-storage-migration-request.json"

echo "Recommendation saved to: ${SCRIPT_DIR}/object-storage-migration-request.json"

echo ""
echo "=== [2] Provision target Object Storage (Bucket) with nameSeed='${NAME_SEED}' ==="
curl -s -i -X POST "${BEETLE_URL}/migration/middleware/ns/${NS_ID}/objectStorage?nameSeed=${NAME_SEED}" \
  -H "Content-Type: application/json" \
  -u "${AUTH}" \
  -d @"${SCRIPT_DIR}/object-storage-migration-request.json"

echo ""
echo "=== [3] List migrated Object Storage Buckets in namespace '${NS_ID}' ==="
curl -s -X GET "${BEETLE_URL}/migration/middleware/ns/${NS_ID}/objectStorage" \
  -u "${AUTH}" | jq .

echo ""
echo "=== [4] Clean up (Delete test bucket '${NAME_SEED}-os-01') ==="
echo "To delete the test bucket, run the following command:"
echo "curl -s -X DELETE \"${BEETLE_URL}/migration/middleware/ns/${NS_ID}/objectStorage/${NAME_SEED}-os-01\" -u \"${AUTH}\""
