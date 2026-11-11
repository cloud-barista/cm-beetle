# CM-Beetle test results for IBM (object storage)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle object storage integration with IBM.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.4+ (95cc61d)
- CB-Tumblebug: v0.13.8
- Target CSP: IBM
- Target Region: au-syd
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Name Seed: my
- Test CLI: Custom automated testing tool
- Test Date: November 11, 2026
- Test Time: 13:00:09 KST
- Test Execution: 2026-11-11 13:00:09 KST

### Scenario

1. Recommend target object storage (buckets) via Beetle
1. Migrate (create) object storages via Beetle
1. List all object storages via Beetle
1. Check existence of first bucket (HEAD) via Beetle
1. Get first bucket details via Beetle
1. Delete all buckets (cleanup) via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for IBM

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/middleware/objectStorage` | ✅ **PASS** | 3ms | Pass |
| 2 | `POST /beetle/migration/middleware/ns/mig01/objectStorage` | ✅ **PASS** | 2m41.75s | Pass |
| 3 | `GET /beetle/migration/middleware/ns/mig01/objectStorage` | ✅ **PASS** | 7ms | Pass |
| 4 | `HEAD /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 2.106s | Pass |
| 5 | `GET /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 4.61s | Pass |
| 6 | `DELETE /beetle/migration/middleware/ns/mig01/objectStorage/{{osId}}` | ✅ **PASS** | 28.168s | Pass |

**Overall Result**: 6/6 tests passed ✅

**Total Duration**: 3m21.649301493s

*Test executed on November 11, 2026 at 13:00:09 KST (2026-11-11 13:00:09 KST) using CM-Beetle automated test CLI*

---

## Detailed Test Case Results

> [!NOTE]
> This section provides detailed information for each test case, including API request information and response details.

### Test Case 1: Recommend target object storage

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/recommendation/middleware/objectStorage`
- **Purpose**: Get object storage recommendations for migration
- **Target CSP**: `ibm`
- **Target Region**: `au-syd`

**Request Body**:

<details>
  <summary> <ins>Click to see the request body</ins> </summary>

```json
{
  "desiredCloud": {
    "csp": "ibm",
    "region": "au-syd"
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
    "csp": "ibm",
    "region": "au-syd"
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
    "csp": "ibm",
    "region": "au-syd"
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
            "lastTransitionTime": "2026-11-11T04:00:26Z",
            "reason": "Available",
            "status": "True",
            "type": "Ready"
          },
          {
            "lastTransitionTime": "2026-11-11T04:00:26Z",
            "reason": "Available",
            "status": "True",
            "type": "Synced"
          }
        ],
        "connectionConfig": {
          "configName": "ibm-au-syd",
          "credentialHolder": "admin",
          "credentialName": "ibm",
          "driverName": "ibm-driver-v1.0.so",
          "providerName": "ibm",
          "regionDetail": {
            "description": "Sydney (Australia)",
            "location": {
              "display": "Australia (Sydney)",
              "latitude": -33.86882,
              "longitude": 151.209296
            },
            "regionId": "au-syd",
            "regionName": "au-syd",
            "zones": [
              "au-syd-1",
              "au-syd-2",
              "au-syd-3"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "au-syd",
            "assignedZone": "au-syd-1"
          },
          "regionZoneInfoName": "ibm-au-syd",
          "verified": true
        },
        "connectionName": "ibm-au-syd",
        "cspResourceId": "tbns38hdccs95asu615f",
        "cspResourceName": "tbns38hdccs95asu615f",
        "description": "Created by CM-Beetle",
        "id": "my-os-01",
        "maxKeys": 1000,
        "name": "my-os-01",
        "resourceType": "objectStorage",
        "status": "Available",
        "uid": "tbns38hdccs95asu615f"
      },
      {
        "conditions": [
          {
            "lastTransitionTime": "2026-11-11T04:00:39Z",
            "reason": "Available",
            "status": "True",
            "type": "Ready"
          },
          {
            "lastTransitionTime": "2026-11-11T04:00:39Z",
            "reason": "Available",
            "status": "True",
            "type": "Synced"
          }
        ],
        "connectionConfig": {
          "configName": "ibm-au-syd",
          "credentialHolder": "admin",
          "credentialName": "ibm",
          "driverName": "ibm-driver-v1.0.so",
          "providerName": "ibm",
          "regionDetail": {
            "description": "Sydney (Australia)",
            "location": {
              "display": "Australia (Sydney)",
              "latitude": -33.86882,
              "longitude": 151.209296
            },
            "regionId": "au-syd",
            "regionName": "au-syd",
            "zones": [
              "au-syd-1",
              "au-syd-2",
              "au-syd-3"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "au-syd",
            "assignedZone": "au-syd-1"
          },
          "regionZoneInfoName": "ibm-au-syd",
          "verified": true
        },
        "connectionName": "ibm-au-syd",
        "cspResourceId": "tbei2j4r62g6uv7eppsq",
        "cspResourceName": "tbei2j4r62g6uv7eppsq",
        "description": "Created by CM-Beetle",
        "id": "my-os-02",
        "maxKeys": 1000,
        "name": "my-os-02",
        "resourceType": "objectStorage",
        "status": "Available",
        "uid": "tbei2j4r62g6uv7eppsq"
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
        "connectionName": "ibm-au-syd",
        "id": "my-os-01",
        "name": "my-os-01"
      },
      {
        "connectionName": "ibm-au-syd",
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
        "lastTransitionTime": "2026-11-11T04:00:26Z",
        "reason": "Available",
        "status": "True",
        "type": "Ready"
      },
      {
        "lastTransitionTime": "2026-11-11T04:00:26Z",
        "reason": "Available",
        "status": "True",
        "type": "Synced"
      }
    ],
    "connectionConfig": {
      "configName": "ibm-au-syd",
      "credentialHolder": "admin",
      "credentialName": "ibm",
      "driverName": "ibm-driver-v1.0.so",
      "providerName": "ibm",
      "regionDetail": {
        "description": "Sydney (Australia)",
        "location": {
          "display": "Australia (Sydney)",
          "latitude": -33.86882,
          "longitude": 151.209296
        },
        "regionId": "au-syd",
        "regionName": "au-syd",
        "zones": [
          "au-syd-1",
          "au-syd-2",
          "au-syd-3"
        ]
      },
      "regionRepresentative": true,
      "regionZoneInfo": {
        "assignedRegion": "au-syd",
        "assignedZone": "au-syd-1"
      },
      "regionZoneInfoName": "ibm-au-syd",
      "verified": true
    },
    "connectionName": "ibm-au-syd",
    "cspResourceId": "tbns38hdccs95asu615f",
    "cspResourceName": "tbns38hdccs95asu615f",
    "description": "Created by CM-Beetle",
    "id": "my-os-01",
    "maxKeys": 1000,
    "name": "my-os-01",
    "resourceType": "objectStorage",
    "status": "Available",
    "uid": "tbns38hdccs95asu615f"
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

