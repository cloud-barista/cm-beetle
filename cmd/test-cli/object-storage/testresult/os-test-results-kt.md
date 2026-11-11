# CM-Beetle test results for KT (object storage)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle object storage integration with KT.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.4+ (95cc61d)
- CB-Tumblebug: v0.13.8
- Target CSP: KT
- Target Region: kr1
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Name Seed: my
- Test CLI: Custom automated testing tool
- Test Date: November 11, 2026
- Test Time: 13:04:11 KST
- Test Execution: 2026-11-11 13:04:11 KST

### Scenario

1. Recommend target object storage (buckets) via Beetle
1. Migrate (create) object storages via Beetle
1. List all object storages via Beetle
1. Check existence of first bucket (HEAD) via Beetle
1. Get first bucket details via Beetle
1. Delete all buckets (cleanup) via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for KT

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/middleware/objectStorage` | ✅ **PASS** | 3ms | Pass |
| 2 | `POST /beetle/migration/middleware/ns/mig01/objectStorage` | ✅ **PASS** | 20.866s | Pass |
| 3 | `GET /beetle/migration/middleware/ns/mig01/objectStorage` | ✅ **PASS** | 6ms | Pass |
| 4 | `HEAD /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 199ms | Pass |
| 5 | `GET /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 580ms | Pass |
| 6 | `DELETE /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 8.734s | Pass |

**Overall Result**: 6/6 tests passed ✅

**Total Duration**: 35.396807248s

*Test executed on November 11, 2026 at 13:04:11 KST (2026-11-11 13:04:11 KST) using CM-Beetle automated test CLI*

---

## Detailed Test Case Results

> [!NOTE]
> This section provides detailed information for each test case, including API request information and response details.

### Test Case 1: Recommend target object storage

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/recommendation/middleware/objectStorage`
- **Purpose**: Get object storage recommendations for migration
- **Target CSP**: `kt`
- **Target Region**: `kr1`

**Request Body**:

<details>
  <summary> <ins>Click to see the request body</ins> </summary>

```json
{
  "desiredCloud": {
    "csp": "kt",
    "region": "kr1"
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
  "status": "success",
  "description": "Successfully recommended 2 object storage configuration(s)",
  "targetCloud": {
    "csp": "kt",
    "region": "kr1"
  },
  "targetObjectStorages": [
    {
      "sourceBucketName": "source-bucket-01",
      "bucketName": "os-01",
      "versioningEnabled": false,
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
      ]
    },
    {
      "sourceBucketName": "source-bucket-02",
      "bucketName": "os-02",
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
      ]
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
  "status": "success",
  "description": "Successfully recommended 2 object storage configuration(s)",
  "targetCloud": {
    "csp": "kt",
    "region": "kr1"
  },
  "targetObjectStorages": [
    {
      "sourceBucketName": "source-bucket-01",
      "bucketName": "os-01",
      "versioningEnabled": false,
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
      ]
    },
    {
      "sourceBucketName": "source-bucket-02",
      "bucketName": "os-02",
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
      ]
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
            "lastTransitionTime": "2026-11-11T04:04:15Z",
            "reason": "Available",
            "status": "True",
            "type": "Ready"
          },
          {
            "lastTransitionTime": "2026-11-11T04:04:15Z",
            "reason": "Available",
            "status": "True",
            "type": "Synced"
          }
        ],
        "connectionConfig": {
          "configName": "kt-kr1",
          "credentialHolder": "admin",
          "credentialName": "kt",
          "driverName": "kt-driver-v1.0.so",
          "providerName": "kt",
          "regionDetail": {
            "description": "Mok-dong - Seoul (South Korea)",
            "location": {
              "display": "Mok-dong - Seoul (South Korea)",
              "latitude": 37.536726,
              "longitude": 126.874278
            },
            "regionId": "KR1",
            "regionName": "kr1",
            "zones": [
              "DX-M1"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "KR1",
            "assignedZone": "DX-M1"
          },
          "regionZoneInfoName": "kt-kr1",
          "verified": true
        },
        "connectionName": "kt-kr1",
        "cspResourceId": "tbv3mkuplp70lr4b1s6g",
        "cspResourceName": "tbv3mkuplp70lr4b1s6g",
        "description": "Created by CM-Beetle",
        "id": "my-os-01",
        "maxKeys": 1000,
        "name": "my-os-01",
        "resourceType": "objectStorage",
        "status": "Available",
        "uid": "tbv3mkuplp70lr4b1s6g"
      },
      {
        "conditions": [
          {
            "lastTransitionTime": "2026-11-11T04:04:15Z",
            "reason": "Available",
            "status": "True",
            "type": "Ready"
          },
          {
            "lastTransitionTime": "2026-11-11T04:04:15Z",
            "reason": "Available",
            "status": "True",
            "type": "Synced"
          }
        ],
        "connectionConfig": {
          "configName": "kt-kr1",
          "credentialHolder": "admin",
          "credentialName": "kt",
          "driverName": "kt-driver-v1.0.so",
          "providerName": "kt",
          "regionDetail": {
            "description": "Mok-dong - Seoul (South Korea)",
            "location": {
              "display": "Mok-dong - Seoul (South Korea)",
              "latitude": 37.536726,
              "longitude": 126.874278
            },
            "regionId": "KR1",
            "regionName": "kr1",
            "zones": [
              "DX-M1"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "KR1",
            "assignedZone": "DX-M1"
          },
          "regionZoneInfoName": "kt-kr1",
          "verified": true
        },
        "connectionName": "kt-kr1",
        "cspResourceId": "tbcv17kn49kiagbt7d7d",
        "cspResourceName": "tbcv17kn49kiagbt7d7d",
        "description": "Created by CM-Beetle",
        "id": "my-os-02",
        "maxKeys": 1000,
        "name": "my-os-02",
        "resourceType": "objectStorage",
        "status": "Available",
        "uid": "tbcv17kn49kiagbt7d7d"
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
        "connectionName": "kt-kr1",
        "id": "my-os-01",
        "name": "my-os-01"
      },
      {
        "connectionName": "kt-kr1",
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
        "lastTransitionTime": "2026-11-11T04:04:15Z",
        "reason": "Available",
        "status": "True",
        "type": "Ready"
      },
      {
        "lastTransitionTime": "2026-11-11T04:04:15Z",
        "reason": "Available",
        "status": "True",
        "type": "Synced"
      }
    ],
    "connectionConfig": {
      "configName": "kt-kr1",
      "credentialHolder": "admin",
      "credentialName": "kt",
      "driverName": "kt-driver-v1.0.so",
      "providerName": "kt",
      "regionDetail": {
        "description": "Mok-dong - Seoul (South Korea)",
        "location": {
          "display": "Mok-dong - Seoul (South Korea)",
          "latitude": 37.536726,
          "longitude": 126.874278
        },
        "regionId": "KR1",
        "regionName": "kr1",
        "zones": [
          "DX-M1"
        ]
      },
      "regionRepresentative": true,
      "regionZoneInfo": {
        "assignedRegion": "KR1",
        "assignedZone": "DX-M1"
      },
      "regionZoneInfoName": "kt-kr1",
      "verified": true
    },
    "connectionName": "kt-kr1",
    "cspResourceId": "tbv3mkuplp70lr4b1s6g",
    "cspResourceName": "tbv3mkuplp70lr4b1s6g",
    "description": "Created by CM-Beetle",
    "id": "my-os-01",
    "maxKeys": 1000,
    "name": "my-os-01",
    "resourceType": "objectStorage",
    "status": "Available",
    "uid": "tbv3mkuplp70lr4b1s6g"
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

