# CM-Beetle test results for AZURE (object storage)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle object storage integration with AZURE.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.4+ (95cc61d)
- CB-Tumblebug: v0.13.8
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Name Seed: my
- Test CLI: Custom automated testing tool
- Test Date: November 11, 2026
- Test Time: 12:58:03 KST
- Test Execution: 2026-11-11 12:58:03 KST

### Scenario

1. Recommend target object storage (buckets) via Beetle
1. Migrate (create) object storages via Beetle
1. List all object storages via Beetle
1. Check existence of first bucket (HEAD) via Beetle
1. Get first bucket details via Beetle
1. Delete all buckets (cleanup) via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for AZURE

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/middleware/objectStorage` | ✅ **PASS** | 4ms | Pass |
| 2 | `POST /beetle/migration/middleware/ns/mig01/objectStorage` | ✅ **PASS** | 1.02s | Pass |
| 3 | `GET /beetle/migration/middleware/ns/mig01/objectStorage` | ✅ **PASS** | 4ms | Pass |
| 4 | `HEAD /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 39ms | Pass |
| 5 | `GET /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 77ms | Pass |
| 6 | `DELETE /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 6.354s | Pass |

**Overall Result**: 6/6 tests passed ✅

**Total Duration**: 12.504580615s

*Test executed on November 11, 2026 at 12:58:03 KST (2026-11-11 12:58:03 KST) using CM-Beetle automated test CLI*

---

## Detailed Test Case Results

> [!NOTE]
> This section provides detailed information for each test case, including API request information and response details.

### Test Case 1: Recommend target object storage

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/recommendation/middleware/objectStorage`
- **Purpose**: Get object storage recommendations for migration
- **Target CSP**: `azure`
- **Target Region**: `koreasouth`

**Request Body**:

<details>
  <summary> <ins>Click to see the request body</ins> </summary>

```json
{
  "desiredCloud": {
    "csp": "azure",
    "region": "koreasouth"
  },
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
  "status": "partial",
  "description": "Successfully recommended 2 object storage configuration(s)",
  "warnings": [
    "Bucket 'source-bucket-01': CORS disabled (not supported on azure)",
    "Bucket 'source-bucket-02': versioning disabled (not supported on azure)",
    "Bucket 'source-bucket-02': CORS disabled (not supported on azure)"
  ],
  "targetCloud": {
    "csp": "azure",
    "region": "koreasouth"
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
  "status": "partial",
  "description": "Successfully recommended 2 object storage configuration(s)",
  "warnings": [
    "Bucket 'source-bucket-01': CORS disabled (not supported on azure)",
    "Bucket 'source-bucket-02': versioning disabled (not supported on azure)",
    "Bucket 'source-bucket-02': CORS disabled (not supported on azure)"
  ],
  "targetCloud": {
    "csp": "azure",
    "region": "koreasouth"
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

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "Successfully created 2 object storage(s)",
    "objectStorage": [
      {
        "conditions": [
          {
            "lastTransitionTime": "2026-11-11T03:58:06Z",
            "reason": "Available",
            "status": "True",
            "type": "Ready"
          },
          {
            "lastTransitionTime": "2026-11-11T03:58:06Z",
            "reason": "Available",
            "status": "True",
            "type": "Synced"
          }
        ],
        "connectionConfig": {
          "configName": "azure-koreasouth",
          "credentialHolder": "admin",
          "credentialName": "azure",
          "driverName": "azure-driver-v1.0.so",
          "providerName": "azure",
          "regionDetail": {
            "description": "Korea South",
            "location": {
              "display": "Korea South",
              "latitude": 35.1796,
              "longitude": 129.0756
            },
            "regionId": "koreasouth",
            "regionName": "koreasouth",
            "zones": []
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "koreasouth",
            "assignedZone": ""
          },
          "regionZoneInfoName": "azure-koreasouth",
          "verified": true
        },
        "connectionName": "azure-koreasouth",
        "cspResourceId": "tbg2d8ikd9udfti3ndc0",
        "cspResourceName": "tbg2d8ikd9udfti3ndc0",
        "description": "Created by CM-Beetle",
        "id": "my-os-01",
        "maxKeys": 1000,
        "name": "my-os-01",
        "resourceType": "objectStorage",
        "status": "Available",
        "uid": "tbg2d8ikd9udfti3ndc0"
      },
      {
        "conditions": [
          {
            "lastTransitionTime": "2026-11-11T03:58:07Z",
            "reason": "Available",
            "status": "True",
            "type": "Ready"
          },
          {
            "lastTransitionTime": "2026-11-11T03:58:07Z",
            "reason": "Available",
            "status": "True",
            "type": "Synced"
          }
        ],
        "connectionConfig": {
          "configName": "azure-koreasouth",
          "credentialHolder": "admin",
          "credentialName": "azure",
          "driverName": "azure-driver-v1.0.so",
          "providerName": "azure",
          "regionDetail": {
            "description": "Korea South",
            "location": {
              "display": "Korea South",
              "latitude": 35.1796,
              "longitude": 129.0756
            },
            "regionId": "koreasouth",
            "regionName": "koreasouth",
            "zones": []
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "koreasouth",
            "assignedZone": ""
          },
          "regionZoneInfoName": "azure-koreasouth",
          "verified": true
        },
        "connectionName": "azure-koreasouth",
        "cspResourceId": "tbmd242rd6jqro7pubdk",
        "cspResourceName": "tbmd242rd6jqro7pubdk",
        "description": "Created by CM-Beetle",
        "id": "my-os-02",
        "maxKeys": 1000,
        "name": "my-os-02",
        "resourceType": "objectStorage",
        "status": "Available",
        "uid": "tbmd242rd6jqro7pubdk"
      }
    ],
    "status": "created"
  },
  "message": "Successfully created 2 object storage(s)",
  "success": true
}
```

</details>

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
    "objectStorage": [
      {
        "connectionName": "azure-koreasouth",
        "id": "my-os-01",
        "name": "my-os-01"
      },
      {
        "connectionName": "azure-koreasouth",
        "id": "my-os-02",
        "name": "my-os-02"
      }
    ]
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
    "conditions": [
      {
        "lastTransitionTime": "2026-11-11T03:58:06Z",
        "reason": "Available",
        "status": "True",
        "type": "Ready"
      },
      {
        "lastTransitionTime": "2026-11-11T03:58:06Z",
        "reason": "Available",
        "status": "True",
        "type": "Synced"
      }
    ],
    "connectionConfig": {
      "configName": "azure-koreasouth",
      "credentialHolder": "admin",
      "credentialName": "azure",
      "driverName": "azure-driver-v1.0.so",
      "providerName": "azure",
      "regionDetail": {
        "description": "Korea South",
        "location": {
          "display": "Korea South",
          "latitude": 35.1796,
          "longitude": 129.0756
        },
        "regionId": "koreasouth",
        "regionName": "koreasouth",
        "zones": []
      },
      "regionRepresentative": true,
      "regionZoneInfo": {
        "assignedRegion": "koreasouth",
        "assignedZone": ""
      },
      "regionZoneInfoName": "azure-koreasouth",
      "verified": true
    },
    "connectionName": "azure-koreasouth",
    "cspResourceId": "tbg2d8ikd9udfti3ndc0",
    "cspResourceName": "tbg2d8ikd9udfti3ndc0",
    "description": "Created by CM-Beetle",
    "id": "my-os-01",
    "maxKeys": 1000,
    "name": "my-os-01",
    "resourceType": "objectStorage",
    "status": "Available",
    "uid": "tbg2d8ikd9udfti3ndc0"
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

