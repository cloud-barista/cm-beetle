# CM-Beetle test results for NCP (object storage)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle object storage integration with NCP.

## Environment and scenario

### Environment

- CM-Beetle: imdl/v0.1.3
- CB-Tumblebug: vunknown
- Target CSP: NCP
- Target Region: kr
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Name Seed: my
- Test CLI: Custom automated testing tool
- Test Date: May 7, 2026
- Test Time: 21:58:40 KST
- Test Execution: 2026-05-07 21:58:40 KST

### Scenario

1. Recommend target object storage (buckets) via Beetle
1. Migrate (create) object storages via Beetle
1. List all object storages via Beetle
1. Check existence of first bucket (HEAD) via Beetle
1. Get first bucket details via Beetle
1. Delete all buckets (cleanup) via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for NCP

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/middleware/objectStorage` | ✅ **PASS** | 6ms | Pass |
| 2 | `POST /beetle/migration/middleware/ns/mig01/objectStorage` | ✅ **PASS** | 1.931s | Pass |
| 3 | `GET /beetle/migration/middleware/ns/mig01/objectStorage` | ✅ **PASS** | 5ms | Pass |
| 4 | `HEAD /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 198ms | Pass |
| 5 | `GET /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 297ms | Pass |
| 6 | `DELETE /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 4.229s | Pass |

**Overall Result**: 6/6 tests passed ✅

**Total Duration**: 11.67103443s

*Test executed on May 7, 2026 at 21:58:40 KST (2026-05-07 21:58:40 KST) using CM-Beetle automated test CLI*

---

## Detailed Test Case Results

> [!NOTE]
> This section provides detailed information for each test case, including API request information and response details.

### Test Case 1: Recommend target object storage

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/recommendation/middleware/objectStorage`
- **Purpose**: Get object storage recommendations for migration
- **Target CSP**: `ncp`
- **Target Region**: `kr`

**Request Body**:

<details>
  <summary> <ins>Click to see the request body</ins> </summary>

```json
{
  "nameSeed": "my",
  "sourceObjectStorages": [
    {
      "bucketName": "source-bucket-01",
      "corsEnabled": true,
      "corsRule": [
        {
          "allowedOrigin": [
            "https://example.com",
            "https://app.example.com"
          ],
          "allowedMethod": [
            "GET",
            "PUT",
            "POST",
            "DELETE"
          ],
          "allowedHeader": [
            "*"
          ],
          "exposeHeader": [
            "ETag",
            "x-amz-request-id"
          ],
          "maxAgeSeconds": 3600
        }
      ],
      "totalSizeBytes": 10737418240,
      "objectCount": 1000,
      "accessFrequency": "frequent",
      "tags": {
        "env": "production",
        "team": "platform"
      }
    },
    {
      "bucketName": "source-bucket-02",
      "versioningEnabled": true,
      "corsEnabled": true,
      "corsRule": [
        {
          "allowedOrigin": [
            "*"
          ],
          "allowedMethod": [
            "GET"
          ],
          "allowedHeader": [
            "Content-Type",
            "Authorization"
          ],
          "exposeHeader": [
            "ETag"
          ],
          "maxAgeSeconds": 86400
        }
      ],
      "encryptionEnabled": true,
      "totalSizeBytes": 1073741824,
      "objectCount": 100,
      "accessFrequency": "infrequent",
      "tags": {
        "env": "staging",
        "team": "data"
      }
    }
  ]
}
```

</details>

#### 1.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Object storage recommendation generated successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "nameSeed": "my",
  "status": "partial",
  "description": "Successfully recommended 2 object storage configuration(s)",
  "warnings": [
    "Bucket 'source-bucket-01': CORS disabled (not supported on ncp)",
    "Bucket 'source-bucket-02': versioning disabled (not supported on ncp)",
    "Bucket 'source-bucket-02': CORS disabled (not supported on ncp)"
  ],
  "targetCloud": {
    "csp": "ncp",
    "region": "kr"
  },
  "targetObjectStorages": [
    {
      "sourceBucketName": "source-bucket-01",
      "bucketName": "os-01",
      "versioningEnabled": false,
      "corsEnabled": false
    },
    {
      "sourceBucketName": "source-bucket-02",
      "bucketName": "os-02",
      "versioningEnabled": false,
      "corsEnabled": false
    }
  ]
}
```

</details>

### Test Case 2: Migrate (create) object storages

#### 2.1 API Request Information

- **API Endpoint**: `POST /beetle/migration/middleware/ns/mig01/objectStorage`
- **Purpose**: Create object storages (buckets) in the target cloud
- **Namespace ID**: `mig01`

**Request Body** (recommendation result):

<details>
  <summary> <ins>Click to see the request body</ins> </summary>

```json
{
  "nameSeed": "my",
  "status": "partial",
  "description": "Successfully recommended 2 object storage configuration(s)",
  "warnings": [
    "Bucket 'source-bucket-01': CORS disabled (not supported on ncp)",
    "Bucket 'source-bucket-02': versioning disabled (not supported on ncp)",
    "Bucket 'source-bucket-02': CORS disabled (not supported on ncp)"
  ],
  "targetCloud": {
    "csp": "ncp",
    "region": "kr"
  },
  "targetObjectStorages": [
    {
      "sourceBucketName": "source-bucket-01",
      "bucketName": "os-01",
      "versioningEnabled": false,
      "corsEnabled": false
    },
    {
      "sourceBucketName": "source-bucket-02",
      "bucketName": "os-02",
      "versioningEnabled": false,
      "corsEnabled": false
    }
  ]
}
```

</details>

#### 2.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Object storages created successfully

### Test Case 3: List object storages

#### 3.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/middleware/ns/mig01/objectStorage`
- **Purpose**: Retrieve all object storages in the namespace
- **Request Body**: None (GET request)

#### 3.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Object storage list retrieved successfully

**Response Body**:

```json
{
  "data": {
    "objectStorage": null
  },
  "success": true
}
```

### Test Case 4: Check existence of first bucket

#### 4.1 API Request Information

- **API Endpoint**: `HEAD /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}`
- **Purpose**: Verify the first migrated bucket exists
- **Request Body**: None (HEAD request)

#### 4.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **HTTP Status**: 200

### Test Case 5: Get first bucket details

#### 5.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}`
- **Purpose**: Retrieve detailed information for the first migrated bucket
- **Request Body**: None (GET request)

#### 5.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Bucket details retrieved successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "connectionName": "ncp-kr",
    "description": "Created by CM-Beetle",
    "id": "my-os-01",
    "name": "my-os-01",
    "status": "Available"
  },
  "success": true
}
```

</details>

### Test Case 6: Delete all buckets (cleanup)

#### 6.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}`
- **Purpose**: Delete all migrated buckets as cleanup
- **Note**: Always runs regardless of previous test failures

#### 6.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: All buckets deleted successfully

