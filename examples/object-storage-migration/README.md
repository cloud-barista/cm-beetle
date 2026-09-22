# Object Storage Recommendation and Migration Example

This example demonstrates how to use the CM-Beetle REST APIs to:

1. **Recommend** optimal Object Storage (bucket) configurations for a target CSP/region.
2. **Migrate (Provision)** target cloud buckets with naming customization (Late Binding).
3. **Verify and Manage** migrated buckets.

---

## Prerequisites

- CM-Beetle server running on `http://localhost:8056`
- Basic authentication credentials: `default:default`
- Target namespace (e.g., `mig01`) created or available in CB-Tumblebug

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
<summary><strong>Alternative Discovery Methods (Temporary Test APIs & Beetle UX Lab)</strong></summary>

<br>

> [!NOTE]
> **Temporary Test APIs**:
> The `scan` and `inspect` endpoints (`POST .../objectStorage/scan` and `POST .../objectStorage/inspect`) are **temporary internal APIs** introduced in CM-Beetle specifically for testing, feature evaluation, and UI prototyping. They may be deprecated or replaced in future releases.

#### Method B: Automated Discovery via Scan & Inspect APIs (Temporary Test APIs)

If you have source cloud credentials, you can discover and inspect buckets automatically:

1. **Scan bucket list:**
   ```bash
   curl -s -X POST "http://localhost:8056/beetle/migration/middleware/objectStorage/scan" \
     -H "Content-Type: application/json" \
     -u "default:default" \
     -d '{
       "csp": "kt",
       "region": "kr1",
       "accessKeyId": "YOUR_SOURCE_KEY_ID",
       "secretAccessKey": "YOUR_SOURCE_SECRET_KEY"
     }' | jq .
   ```

2. **Inspect selected bucket metadata (CORS, Versioning, etc.):**
   ```bash
   curl -s -X POST "http://localhost:8056/beetle/migration/middleware/objectStorage/inspect" \
     -H "Content-Type: application/json" \
     -u "default:default" \
     -d '{
       "csp": "kt",
       "region": "kr1",
       "accessKeyId": "YOUR_SOURCE_KEY_ID",
       "secretAccessKey": "YOUR_SOURCE_SECRET_KEY",
       "selectedBucketNames": ["kt-test-source-bucket"]
     }' | jq .
   ```

#### Method C: Visual Discovery via Beetle UX Lab Portal

You can also use the Beetle UX Lab GUI (`http://localhost:8056/beetle/ui/`) under **Object Storage Migration** to enter credentials, scan buckets, and inspect properties with a single click.

</details>

---

## Step 1: Recommend Object Storage

Analyze source bucket configurations (versioning, CORS, encryption, capacity) and recommend compatible bucket configurations for the target cloud.

- **API Endpoint:** `POST /beetle/recommendation/middleware/objectStorage`
- **API Documentation:** [Managed Object Storage Recommendation API Spec](https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BRecommendation%5D%20Managed%20Object%20Storage)
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
  "data": {
    "status": "partial",
    "description": "Successfully recommended 1 object storage configuration(s)",
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
> The status is reported as `partial` because NCP Object Storage does not support bucket versioning or CORS in the same manner as AWS or KT Cloud. CM-Beetle automatically reconciles these CSP feature differences, disables unsupported options, and records the changes in `warnings` so that the output model can be immediately provisioned without manual fixes.

### Save Recommendation Directly for Migration

You can pipe the recommendation `data` directly into `object-storage-migration-request.json`:

```bash
curl -s -X POST "http://localhost:8056/beetle/recommendation/middleware/objectStorage" \
  -H "Content-Type: application/json" \
  -u "default:default" \
  -d @object-storage-recommendation-request.json | jq '.data' > object-storage-migration-request.json
```

---

## Step 2: Migrate (Provision) Object Storage Buckets

Create the recommended buckets in the target cloud provider.

- **API Endpoint:** `POST /beetle/migration/middleware/ns/{nsId}/objectStorage`
- **API Documentation:** [Managed Object Storage Migration API Spec](https://cloud-barista.github.io/api/?url=https://raw.githubusercontent.com/cloud-barista/cm-beetle/v0.6.1/api/swagger.yaml#/%5BMigration%5D%20Managed%20Object%20Storage)
- **Payload File:** `object-storage-migration-request.json`

> [!NOTE]
> **About `nameSeed` (Late Binding Naming):**  
> The `nameSeed` query parameter prefixes target bucket names at migration time (e.g., `os-01` becomes `test-os-01`).
>
> - **Validation Rules:** 1 to 20 characters, alphanumeric and hyphens (`-`) only, must start with an alphanumeric character.
> - **Omission:** If omitted, recommended bucket names (`os-01`, etc.) are used as-is.

### Option 2A: Synchronous Execution (with `nameSeed=test`)

Using `?nameSeed=test` prefixes target bucket names (`os-01` becomes `test-os-01`):

```bash
curl -s -i -X POST "http://localhost:8056/beetle/migration/middleware/ns/mig01/objectStorage?nameSeed=test" \
  -H "Content-Type: application/json" \
  -u "default:default" \
  -d @object-storage-migration-request.json
```

- **Expected Response (`HTTP/1.1 201 Created`):**
  ```json
  {
    "data": {
      "targetCloud": { "csp": "ncp", "region": "kr" },
      "targetObjectStorages": [
        {
          "id": "test-os-01",
          "name": "test-os-01",
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

- **Response:** HTTP 202 with `reqId`
- **Check Status:**
  ```bash
  curl -s -X GET "http://localhost:8056/beetle/request/{reqId}" -u "default:default" | jq .
  ```

---

## Step 3: Verify and Inspect Migrated Buckets

### List All Buckets in Namespace

```bash
curl -s -X GET "http://localhost:8056/beetle/migration/middleware/ns/mig01/objectStorage" \
  -u "default:default" | jq .
```

- **Expected Output:**
  ```json
  [
    {
      "id": "test-os-01",
      "name": "test-os-01",
      "connectionName": "ncp-kr",
      "status": "Available"
    }
  ]
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

### Delete Bucket (Cleanup)

> [!TIP]
> **Planning to test Data Migration next?**
> Keep `test-os-01` active! Do **NOT** run the delete command yet, as the next tutorial ([Data Migration](../data-migration/)) transfers objects directly into this `test-os-01` bucket. Delete it only after data migration testing is complete.

```bash
curl -s -X DELETE "http://localhost:8056/beetle/migration/middleware/ns/mig01/objectStorage/test-os-01" \
  -u "default:default" | jq .
```

---

## Reference Automation Script (`run.sh`)

An illustrative script [`run.sh`](./run.sh) is provided as an automation reference example demonstrating how to chain the recommendation, provisioning, and inspection API calls in a bash workflow.

> [!NOTE]
> This script is shared strictly for **reference and educational purposes**. To understand each phase of the Object Storage lifecycle and verify intermediate outputs, follow the individual curl commands described in the sections above.

