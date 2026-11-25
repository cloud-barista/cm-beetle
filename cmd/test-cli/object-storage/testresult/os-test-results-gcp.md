# CM-Beetle test results for GCP (object storage)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle object storage integration with GCP.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.5+ (83361e8)
- CB-Tumblebug: v0.14.0
- Target CSP: GCP
- Target Region: asia-northeast3
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Name Seed: my
- Test CLI: Custom automated testing tool
- Test Date: November 25, 2026
- Test Time: 18:35:28 KST
- Test Execution: 2026-11-25 18:35:28 KST

### Scenario

1. Recommend target object storage (buckets) via Beetle
1. Migrate (create) object storages via Beetle
1. List all object storages via Beetle
1. Check existence of first bucket (HEAD) via Beetle
1. Get first bucket details via Beetle
1. Delete all buckets (cleanup) via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for GCP

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/middleware/objectStorage` | ✅ **PASS** | 16ms | Pass |
| 2 | `POST /beetle/migration/middleware/ns/mig01/objectStorage` | ✅ **PASS** | 10.617s | Pass |
| 3 | `GET /beetle/migration/middleware/ns/mig01/objectStorage` | ✅ **PASS** | 7ms | Pass |
| 4 | `HEAD /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 696ms | Pass |
| 5 | `GET /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 1.236s | Pass |
| 6 | `DELETE /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 11.544s | Pass |

**Overall Result**: 6/6 tests passed ✅

**Total Duration**: 29.123444285s

*Test executed on November 25, 2026 at 18:35:28 KST (2026-11-25 18:35:28 KST) using CM-Beetle automated test CLI*

---

## Detailed Test Case Results

> [!NOTE]
> This section provides detailed information for each test case, including API request information and response details.

### Test Case 1: Recommend target object storage

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/recommendation/middleware/objectStorage`
- **Purpose**: Get object storage recommendations for migration
- **Target CSP**: `gcp`
- **Target Region**: `asia-northeast3`

**Request Body**:

<details>
  <summary> <ins>Click to see the request body</ins> </summary>

```json
{
  "desiredCloud": {
    "csp": "gcp",
    "region": "asia-northeast3"
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
    "csp": "gcp",
    "region": "asia-northeast3"
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
    "csp": "gcp",
    "region": "asia-northeast3"
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
            "lastTransitionTime": "2026-11-25T09:35:35Z",
            "reason": "Available",
            "status": "True",
            "type": "Ready"
          },
          {
            "lastTransitionTime": "2026-11-25T09:35:35Z",
            "reason": "Available",
            "status": "True",
            "type": "Synced"
          }
        ],
        "connectionConfig": {
          "configName": "gcp-asia-northeast3",
          "credentialHolder": "admin",
          "credentialName": "gcp",
          "driverName": "gcp-driver-v1.0.so",
          "providerName": "gcp",
          "regionDetail": {
            "description": "Seoul South Korea",
            "location": {
              "display": "South Korea (Seoul)",
              "latitude": 37.2,
              "longitude": 127
            },
            "regionId": "asia-northeast3",
            "regionName": "asia-northeast3",
            "zones": [
              "asia-northeast3-a",
              "asia-northeast3-b",
              "asia-northeast3-c"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "asia-northeast3",
            "assignedZone": "asia-northeast3-a"
          },
          "regionZoneInfoName": "gcp-asia-northeast3",
          "verified": true
        },
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "tbdkbtcmbjnk2hkh16fn",
        "cspResourceName": "tbdkbtcmbjnk2hkh16fn",
        "description": "Created by CM-Beetle",
        "id": "my-os-01",
        "maxKeys": 1000,
        "name": "my-os-01",
        "resourceType": "objectStorage",
        "status": "Available",
        "uid": "tbdkbtcmbjnk2hkh16fn"
      },
      {
        "conditions": [
          {
            "lastTransitionTime": "2026-11-25T09:35:39Z",
            "reason": "Available",
            "status": "True",
            "type": "Ready"
          },
          {
            "lastTransitionTime": "2026-11-25T09:35:39Z",
            "reason": "Available",
            "status": "True",
            "type": "Synced"
          }
        ],
        "connectionConfig": {
          "configName": "gcp-asia-northeast3",
          "credentialHolder": "admin",
          "credentialName": "gcp",
          "driverName": "gcp-driver-v1.0.so",
          "providerName": "gcp",
          "regionDetail": {
            "description": "Seoul South Korea",
            "location": {
              "display": "South Korea (Seoul)",
              "latitude": 37.2,
              "longitude": 127
            },
            "regionId": "asia-northeast3",
            "regionName": "asia-northeast3",
            "zones": [
              "asia-northeast3-a",
              "asia-northeast3-b",
              "asia-northeast3-c"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "asia-northeast3",
            "assignedZone": "asia-northeast3-a"
          },
          "regionZoneInfoName": "gcp-asia-northeast3",
          "verified": true
        },
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "tbijp6t91bllcaeaehg8",
        "cspResourceName": "tbijp6t91bllcaeaehg8",
        "description": "Created by CM-Beetle",
        "id": "my-os-02",
        "maxKeys": 1000,
        "name": "my-os-02",
        "resourceType": "objectStorage",
        "status": "Available",
        "uid": "tbijp6t91bllcaeaehg8"
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
        "connectionName": "gcp-asia-northeast3",
        "id": "my-os-01",
        "name": "my-os-01"
      },
      {
        "connectionName": "gcp-asia-northeast3",
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
        "lastTransitionTime": "2026-11-25T09:35:35Z",
        "reason": "Available",
        "status": "True",
        "type": "Ready"
      },
      {
        "lastTransitionTime": "2026-11-25T09:35:35Z",
        "reason": "Available",
        "status": "True",
        "type": "Synced"
      }
    ],
    "connectionConfig": {
      "configName": "gcp-asia-northeast3",
      "credentialHolder": "admin",
      "credentialName": "gcp",
      "driverName": "gcp-driver-v1.0.so",
      "providerName": "gcp",
      "regionDetail": {
        "description": "Seoul South Korea",
        "location": {
          "display": "South Korea (Seoul)",
          "latitude": 37.2,
          "longitude": 127
        },
        "regionId": "asia-northeast3",
        "regionName": "asia-northeast3",
        "zones": [
          "asia-northeast3-a",
          "asia-northeast3-b",
          "asia-northeast3-c"
        ]
      },
      "regionRepresentative": true,
      "regionZoneInfo": {
        "assignedRegion": "asia-northeast3",
        "assignedZone": "asia-northeast3-a"
      },
      "regionZoneInfoName": "gcp-asia-northeast3",
      "verified": true
    },
    "connectionName": "gcp-asia-northeast3",
    "cspResourceId": "tbdkbtcmbjnk2hkh16fn",
    "cspResourceName": "tbdkbtcmbjnk2hkh16fn",
    "description": "Created by CM-Beetle",
    "id": "my-os-01",
    "maxKeys": 1000,
    "name": "my-os-01",
    "resourceType": "objectStorage",
    "status": "Available",
    "uid": "tbdkbtcmbjnk2hkh16fn"
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

