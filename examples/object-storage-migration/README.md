# Object Storage Recommendation and Migration Example

This example demonstrates how to use the CM-Beetle REST APIs to:

1. **Recommend** optimal Object Storage (bucket) configurations for a target CSP and region.
2. **Migrate (Provision)** target cloud buckets with naming customization (Late Binding).
3. **Verify and Clean Up** migrated buckets and objects.

---

## Prerequisites

- CM-Beetle server running on `http://localhost:8056`
- Basic authentication credentials: `default:default`
- Target namespace (for example, `mig01`) created in CB-Tumblebug

---

## Step 0: Prepare Source Object Storage Information

> [!IMPORTANT]
> To request an Object Storage recommendation, **information about your existing source bucket(s) must be prepared in advance**.
> CM-Beetle evaluates these source properties (versioning, CORS, encryption, capacity) against the target CSP feature matrix to recommend compatible bucket configurations.

You can prepare the source bucket information in one of the following ways:

### Method A: Manual JSON Preparation (Recommended)

Directly edit `object-storage-recommendation-request.json` with your source bucket details.
At minimum, the `bucketName` field is required:

```json
{
  "desiredCloud": { "csp": "ncp", "region": "kr" },
  "sourceObjectStorages": [
    {
      "bucketName": "kt-test-source-bucket",
      "versioningEnabled": true,
      "corsEnabled": true,
      "totalSizeBytes": 1073741824,
      "objectCount": 100,
      "accessFrequency": "frequent",
      "tags": {
        "SourceCsp": "kt",
        "SourceRegion": "kr1"
      }
    }
  ]
}
```

<details>
<summary><strong>Alternative Discovery Methods ([TEST ONLY] Client Helpers and Beetle UX Lab)</strong></summary>

<br>

> [!WARNING]
> **Non-Production Test Helpers Only (`[Test] Client Helper (Non-Production)`)**:
> The `publicKey`, `scan`, and `inspect` endpoints (`GET /beetle/migration/security/publicKey`, `POST /beetle/migration/middleware/objectStorage/scan`, and `POST /beetle/migration/middleware/objectStorage/inspect`) are **non-production test helpers** provided solely for local evaluation and Beetle UX Lab prototyping.
>
> - Do **NOT** send cloud credentials to these endpoints in production workflows; production clients should inspect source S3 buckets directly on the client side.
> - Supported S3-compatible CSPs for `/scan` and `/inspect`: `aws`, `gcp`, `alibaba`, `tencent`, `ibm`, `ncp` (unsupported CSPs return `400 Bad Request`).

#### Method B: Automated Discovery via `[TEST ONLY]` Scan and Inspect Helpers

For supported S3-compatible CSPs (`aws`, `gcp`, `alibaba`, `tencent`, `ibm`, `ncp`), you can test bucket discovery locally:

1. **Scan bucket list:**
   ```bash
   curl -s -X POST "http://localhost:8056/beetle/migration/middleware/objectStorage/scan" \
     -H "Content-Type: application/json" \
     -u "default:default" \
     -d '{
       "csp": "aws",
       "region": "ap-northeast-2",
       "accessKeyId": "YOUR_SOURCE_KEY_ID",
       "secretAccessKey": "YOUR_SOURCE_SECRET_KEY"
     }' | jq .
   ```

2. **Inspect selected bucket metadata (CORS rules, Versioning, encryption, object summary):**
   ```bash
   curl -s -X POST "http://localhost:8056/beetle/migration/middleware/objectStorage/inspect" \
     -H "Content-Type: application/json" \
     -u "default:default" \
     -d '{
       "csp": "aws",
       "region": "ap-northeast-2",
       "accessKeyId": "YOUR_SOURCE_KEY_ID",
       "secretAccessKey": "YOUR_SOURCE_SECRET_KEY",
       "selectedBucketNames": ["my-source-bucket"]
     }' | jq .
   ```

#### Method C: Visual Discovery via Beetle UX Lab Portal

You can also use the Beetle UX Lab GUI (`http://localhost:8056/beetle/ui/`) under **Object Storage Migration** for local evaluation.

</details>

---

## Step 1: Recommend Object Storage

Analyze source bucket configurations (versioning, CORS, encryption, capacity) and recommend compatible bucket configurations for the target cloud.

- **API Endpoint:** `POST /beetle/recommendation/middleware/objectStorage`
- **Payload File:** `object-storage-recommendation-request.json`

### curl Command

```bash
curl -s -X POST "http://localhost:8056/beetle/recommendation/middleware/objectStorage" \
  -H "Content-Type: application/json" \
  -u "default:default" \
  -d @object-storage-recommendation-request.json | jq .
```

### Expected Response

```json
{
  "success": true,
  "data": {
    "status": "partial",
    "description": "Recommended 1 object storage configuration(s) for ncp kr with 2 warning(s)",
    "warnings": [
      "Bucket 'kt-test-source-bucket': versioning disabled (not supported on ncp)",
      "Bucket 'kt-test-source-bucket': CORS disabled (not supported on ncp)"
    ],
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetObjectStorages": [
      {
        "sourceBucketName": "kt-test-source-bucket",
        "bucketName": "os-01",
        "versioningEnabled": false,
        "corsEnabled": false
      }
    ]
  }
}
```

> [!NOTE]
> **Understanding `"status": "partial"` and `warnings`**:
> The status is reported as `partial` because NCP Object Storage does not support bucket-level versioning or CORS via CB-Tumblebug. CM-Beetle automatically disables unsupported options and records the adjustments in `warnings` so the output model can be provisioned directly.

### Save Recommendation Directly for Migration

Pipe the recommendation `data` directly into `object-storage-migration-request.json`:

```bash
curl -s -X POST "http://localhost:8056/beetle/recommendation/middleware/objectStorage" \
  -H "Content-Type: application/json" \
  -u "default:default" \
  -d @object-storage-recommendation-request.json | jq '.data' > object-storage-migration-request.json
```

---

## Step 2: Migrate (Provision) Object Storage Buckets

Create the recommended buckets in the target cloud provider. Each bucket in `targetObjectStorages` is provisioned independently.

- **API Endpoint:** `POST /beetle/migration/middleware/ns/{nsId}/objectStorage`
- **Payload File:** `object-storage-migration-request.json`

> [!NOTE]
> **About `nameSeed` (Late Binding Naming):**
> The `nameSeed` query parameter prefixes target bucket names at migration time (for example, `os-01` becomes `test-os-01`).
>
> - **Validation Rules:** 1 to 20 characters, alphanumeric and hyphens (`-`) only, must start with an alphanumeric character.
> - **Omission:** If omitted, recommended base bucket names (`os-01`, etc.) are used as-is.

### Option 2A: Synchronous Execution (with `nameSeed=test`)

```bash
curl -s -X POST "http://localhost:8056/beetle/migration/middleware/ns/mig01/objectStorage?nameSeed=test" \
  -H "Content-Type: application/json" \
  -u "default:default" \
  -d @object-storage-migration-request.json | jq .
```

- **Expected Response (`HTTP/1.1 201 Created`):**
  ```json
  {
    "success": true,
    "data": {
      "status": "created",
      "description": "Successfully created 1 of 1 object storage(s) for ncp kr",
      "objectStorage": [
        {
          "id": "test-os-01",
          "name": "test-os-01",
          "connectionName": "ncp-kr",
          "cspResourceName": "test-os-01-...",
          "status": "Available"
        }
      ]
    }
  }
  ```

### Option 2B: Asynchronous Execution (`Prefer: respond-async`)

For multi-bucket provisioning, trigger async execution:

```bash
curl -s -X POST "http://localhost:8056/beetle/migration/middleware/ns/mig01/objectStorage?nameSeed=test" \
  -H "Content-Type: application/json" \
  -H "Prefer: respond-async" \
  -u "default:default" \
  -d @object-storage-migration-request.json | jq .
```

- **Response:** `HTTP/1.1 202 Accepted` with `reqId`
- **Poll Status (`MigratedObjectStorageResult` in `responseData`):**
  ```bash
  curl -s -X GET "http://localhost:8056/beetle/request/{reqId}" -u "default:default" | jq .
  ```

---

## Step 3: Verify and Clean Up Migrated Buckets

### List All Buckets in Namespace

```bash
curl -s -X GET "http://localhost:8056/beetle/migration/middleware/ns/mig01/objectStorage" \
  -u "default:default" | jq .
```

- **Expected Output:**
  ```json
  {
    "success": true,
    "data": {
      "objectStorage": [
        {
          "id": "test-os-01",
          "name": "test-os-01",
          "connectionName": "ncp-kr",
          "status": "Available"
        }
      ]
    }
  }
  ```

### Get Specific Bucket Details

```bash
curl -s -X GET "http://localhost:8056/beetle/migration/middleware/ns/mig01/objectStorage/test-os-01" \
  -u "default:default" | jq .
```

### Check Bucket Existence (HEAD)

```bash
curl -s -I -X HEAD "http://localhost:8056/beetle/migration/middleware/ns/mig01/objectStorage/test-os-01" \
  -u "default:default"
```

- **Expected Output:** `HTTP/1.1 200 OK`

### Post-Migration Object Verification & Day-2 Operations

After transferring objects (see [`../data-migration/`](../data-migration/)), you can verify transferred objects directly via CM-Beetle:

- **List objects:** `GET /beetle/migration/middleware/ns/mig01/objectStorage/test-os-01/object`
- **Inspect object headers (`ETag`, `Last-Modified`, `Content-Length`, `X-Storage-Class`):** `HEAD /beetle/migration/middleware/ns/mig01/objectStorage/test-os-01/object/{objectKey}`
- **Delete verification object:** `DELETE /beetle/migration/middleware/ns/mig01/objectStorage/test-os-01/object/{objectKey}`

For **Day-2 operations** after migration (such as generating Presigned URLs or modifying CORS/Versioning policies), use the CB-Tumblebug Object Storage APIs (`/tumblebug/ns/{nsId}/resources/objectStorage/{osId}/...`).

### Delete Bucket (Cleanup)

> [!TIP]
> **Planning to test Data Migration next?**
> Keep `test-os-01` active! Do **NOT** delete the bucket yet, as the next tutorial ([Data Migration](../data-migration/)) transfers objects directly into `test-os-01`.

`DELETE .../objectStorage/{osId}` supports three explicit modes:

- **Standard delete (empty bucket):** omit `option` to delete the bucket itself.
- **Empty contents only (keep bucket):** `?option=empty` removes all objects/versions while keeping `test-os-01` active.
- **Force delete (non-empty bucket):** `?option=force` empties all objects/versions first and then deletes `test-os-01`.

```bash
# Delete an empty bucket (returns HTTP 204 No Content)
curl -i -X DELETE "http://localhost:8056/beetle/migration/middleware/ns/mig01/objectStorage/test-os-01" \
  -u "default:default"

# Or empty all objects/versions and delete the bucket in one step (returns HTTP 204 No Content)
curl -i -X DELETE "http://localhost:8056/beetle/migration/middleware/ns/mig01/objectStorage/test-os-01?option=force" \
  -u "default:default"
```

---

## Reference Automation Script (`run.sh`)

An illustrative script [`run.sh`](./run.sh) is provided as an automation reference example demonstrating how to chain the recommendation, provisioning, and inspection API calls in a bash workflow.
