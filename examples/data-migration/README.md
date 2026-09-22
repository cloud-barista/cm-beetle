# Data Migration (Source KT Cloud kr1 → Destination NCP kr) Example

This example demonstrates how to transfer data from a **Source KT Cloud Object Storage (`kt`, `kr1` region)** to a target cloud Object Storage bucket managed by **CB-Tumblebug (NCP `kr` region)** using the CM-Beetle Data Migration REST API (`POST /beetle/migration/data`).

- **API Documentation:** [Data Migration API (incubating) Spec](<https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BMigration%5D%20Data%20(incubating)>)

---

## Architecture & Scenario Overview

```text
┌────────────────────────────────────────┐         ┌───────────────────────────┐         ┌──────────────────────────────────────┐
│       Source: KT Cloud kr1             │ ──────> │  CM-Beetle Relay Worker   │ ──────> │        Target: NCP kr                │
│ (S3-compatible via accessType: minio)  │         │ (Memory Staging / Relay)  │         │ (Managed by Tumblebug: test-os-01)   │
└────────────────────────────────────────┘         └───────────────────────────┘         └──────────────────────────────────────┘
```

- **Source:** KT Cloud Object Storage in region `kr1` accessed via MinIO S3 SDK engine (`accessType: "minio"`).
- **Destination:** Target NCP bucket (`test-os-01`) in namespace `mig01`, provisioned via the `object-storage-migration` step (`accessType: "tumblebug"`).
- **Payload File:** `data-migration-request.json`

---

## Step 0: Prerequisites and Preparation for Data Migration

> [!IMPORTANT]
> Before triggering data migration, **both the source S3 credentials and the destination target bucket must be prepared in advance**.
> Unlike the Object Storage provisioning API, the Data Migration API transfers data into an **already existing** destination bucket and will not create buckets automatically.

Ensure the following prerequisites and items are prepared before executing migration:

### 1. Source Storage Preparation (KT Cloud kr1)
- **Bucket and Path**: The source bucket (e.g., `kt-test-source-bucket`) must exist and contain the data to be migrated. Specify the directory prefix properly (e.g., `bucket-name/` for the whole bucket, or `bucket-name/data/` for a specific subfolder).
- **S3 API Credentials**: Valid `accessKeyId` and `secretAccessKey` with read permissions (`s3:GetObject`, `s3:ListBucket`).
- **Endpoint Reachability**: CM-Beetle must be able to reach the KT Cloud S3 endpoint (`s3.ap-northeast-1.ktcloud.com`).

### 2. Destination Storage Preparation (NCP kr via Tumblebug)
- **Pre-provisioned Target Bucket**: The target bucket (e.g., `test-os-01`) **must be provisioned beforehand** (completed in the `object-storage-migration` step).
- **Registered IDs**: The target namespace ID (`mig01`) and object storage ID (`test-os-01`) must be valid in CB-Tumblebug.
- **Write Permissions**: The connection config in CB-Tumblebug must have write access (`s3:PutObject`) to the target bucket.

### 3. Server Staging Disk Space & Filtering
- **Relay Transfer Staging**: CM-Beetle uses a secure relay architecture (`Source S3 → CM-Beetle Local Staging → Destination S3`). Ensure adequate disk space is available in `/tmp/transx-staging` for in-transit data chunks.
- **Filter Configuration**: Use `filter.include` and `filter.exclude` to skip temporary files, caches, or trash directories (`*.tmp`, `.trash/*`).

---

## Configuration (`data-migration-request.json`)

```json
{
  "source": {
    "storageType": "objectstorage",
    "path": "kt-test-source-bucket/data/",
    "objectStorage": {
      "accessType": "minio",
      "minio": {
        "endpoint": "s3.ap-northeast-1.ktcloud.com",
        "accessKeyId": "YOUR_KT_S3_ACCESS_KEY",
        "secretAccessKey": "YOUR_KT_S3_SECRET_KEY",
        "region": "kr1",
        "useSSL": true
      }
    },
    "filter": {
      "include": ["*"],
      "exclude": ["*.tmp", ".trash/*"]
    }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "test-os-01/data/",
    "objectStorage": {
      "accessType": "tumblebug",
      "tumblebug": {
        "endpoint": "http://localhost:1323/tumblebug",
        "nsId": "mig01",
        "osId": "test-os-01",
        "expires": 3600,
        "auth": {
          "authType": "basic",
          "basic": {
            "username": "default",
            "password": "default"
          }
        }
      }
    }
  },
  "strategy": "auto"
}
```

### Configuration Fields Explained

- `source.objectStorage.minio.endpoint`: KT Cloud S3 endpoint (`s3.ap-northeast-1.ktcloud.com`).
- `source.objectStorage.minio.accessKeyId` & `secretAccessKey`: Replace `YOUR_KT_S3_ACCESS_KEY` and `YOUR_KT_S3_SECRET_KEY` with your actual KT Cloud credentials for real data transfers.
- `destination.objectStorage.tumblebug.osId`: Points to `test-os-01` created in the previous [Object Storage Provisioning](../object-storage-migration/) step.
- `destination.path`: Destination folder path inside the target bucket (e.g., `test-os-01/data/`).

> [!NOTE]
> **First-Time Testing Without Real KT Cloud Credentials (Dry-Run Behavior)**:
> You can test the step-by-step curl commands below immediately even without editing credentials!
> CM-Beetle will demonstrate:
> 1. Fetching the one-time RSA public key bundle (`GET /encryptionKey`).
> 2. Encrypting sensitive access keys in transit into RSA-OAEP ciphertext (`POST /test/encrypt`).
> 3. Registering the async migration request (`HTTP 202 Accepted`).
> 4. Launching the background worker and tracking job status (`GET /request/{reqId}`).
> The background worker will report a connection error against the placeholder KT host, confirming that the entire security and orchestration lifecycle functions as designed.

---

## Security & Sensitive Credential Encryption (Standard Workflow)

> [!IMPORTANT]
> **Production Security Standard**: Sensitive credentials (`accessKeyId`, `secretAccessKey`, `password`) are **required to be encrypted in transit** using a server-issued one-time RSA public key before invoking the migration API.

### 🔍 Technical Implementation vs. Production Security Policy

| Category | Backend API Server Behavior | Production & Beetle UX Lab Policy (Standard) |
| :--- | :--- | :--- |
| **Credential Encryption** | **Optional (`req.IsEncrypted()`)**<br>- Flexibly accepts plaintext models for local development and curl debugging convenience | **Mandatory Standard**<br>- Enforces one-time public key encryption to prevent credential exposure in transit |
| **Beetle UX Lab Behavior** | - | Prohibits plaintext transmission in `DataTransferCenter.tsx`; always fetches a one-time RSA key via `GET /encryptionKey` to encrypt credentials before execution |

### Automated 3-Step Encryption & Execution Workflow

#### Step 1: Fetch One-Time Public Key Bundle

```bash
KEY_BUNDLE=$(curl -s -X GET "http://localhost:8056/beetle/migration/data/encryptionKey" -u "default:default" | jq '.data')
echo "$KEY_BUNDLE" | jq .
```

- **Expected Response:**
  ```json
  {
    "keyId": "0692f2ed-2982-4561-9610-903d371f80b3",
    "publicKey": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----\n",
    "algorithm": "RSA-OAEP-256",
    "expiresAt": "2026-09-22T08:25:06Z"
  }
  ```

#### Step 2: Encrypt Sensitive Credential Fields

Encrypt the plaintext `data-migration-request.json` using the one-time key:

```bash
PLAINTEXT_MODEL=$(cat data-migration-request.json)

jq -n --argjson kb "$KEY_BUNDLE" --argjson m "$PLAINTEXT_MODEL" \
  '{publicKeyBundle: $kb, model: $m}' | \
  curl -s -X POST "http://localhost:8056/beetle/migration/data/test/encrypt" \
    -H "Content-Type: application/json" \
    -u "default:default" \
    -d @- | jq '.data' > data-migration-request-encrypted.json
```

_(The generated `data-migration-request-encrypted.json` contains encrypted ciphertext credentials and a unique `encryptionKeyId`.)_

> [!IMPORTANT]
> **One-Time Key Lifecycle**:
> Each RSA public key is strictly single-use to prevent replay attacks. Once CM-Beetle decrypts the request payload on the server, the key is immediately purged from the server cache. To send another migration request, fetch a fresh key via `GET /encryptionKey`.

#### Step 3: Launch Data Migration (Asynchronous Recommended)

```bash
curl -s -X POST "http://localhost:8056/beetle/migration/data" \
  -H "Content-Type: application/json" \
  -H "Prefer: respond-async" \
  -u "default:default" \
  -d @data-migration-request-encrypted.json | jq .
```

- **Expected Async Response (HTTP 202 Accepted):**
  ```json
  {
    "success": true,
    "data": {
      "reqId": "1790064906103665774",
      "status": "Handling",
      "statusUrl": "/beetle/request/1790064906103665774"
    },
    "message": "Migration started. Use GET /request/{reqId} to check status."
  }
  ```

---

## Check Migration Job Progress

Poll the status of the background migration job:

```bash
curl -s -X GET "http://localhost:8056/beetle/request/{reqId}" \
  -u "default:default" | jq .
```

- **Expected Response while Transferring:**
  ```json
  {
    "success": true,
    "data": {
      "startTime": "2026-09-22T08:15:06Z",
      "status": "Handling"
    }
  }
  ```

Status transitions from `Handling` → `Success` (or `Error` if network or authentication fails).

---

## Reference Automation Script (`run.sh`)

An illustrative script [`run.sh`](./run.sh) is provided as an automation reference example demonstrating how to chain the one-time key retrieval, credential encryption, asynchronous transfer request, and status polling in bash.

> [!NOTE]
> This script is shared strictly for **reference and educational purposes**. To understand the data migration workflow step-by-step and inspect intermediate states, follow the individual curl commands described in the sections above.
