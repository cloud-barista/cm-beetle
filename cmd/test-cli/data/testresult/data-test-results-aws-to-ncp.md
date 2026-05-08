# CM-Beetle test results for NCP kr (data migration)

> [!NOTE]
> This document presents test results for CM-Beetle data migration to NCP (kr).

## Environment and scenario

### Environment

- CM-Beetle: transx/v0.1.2
- CB-Tumblebug: vunknown
- Source CSP: AWS
- Source Region: ap-northeast-2
- Target CSP: NCP
- Target Region: kr
- Source OS ID: src01-cm-beetle-test-bucket
- Target OS ID: data01-cm-beetle-test-bucket
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Name Seed: data01
- Test CLI: CM-Beetle data migration automated test CLI
- Test Date: May 8, 2026
- Test Time: 20:39:44 KST
- Test Execution: 2026-05-08 20:39:44 KST

### Scenario

**Pre-flight (shared, runs once before all target tests):**

1. Create source object storage via CM-Beetle
1. Upload dummy data to source object storage (local → source OS, encrypted, with filter, async)

**Per target CSP-Region (Steps 1–4 below):**

1. Create target object storage via CM-Beetle
1. Migrate data: source OS → target OS (encrypted, async)
1. Verify migrated data: compare source and target object lists
1. Delete target object storage (cleanup) via CM-Beetle

**Post-flight (shared, runs once after all target tests):**

1. Delete source object storage (cleanup) via CM-Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for NCP kr

### Test Results Summary

| Step | Endpoint / Description | Status | Duration | Details |
|------|------------------------|--------|----------|---------|
| 1 | `POST /beetle/migration/middleware/ns/mig01/objectStorage` (target) | ✅ **PASS** | 1.118s | Pass |
| 2 | `POST /beetle/migration/data` (migrate: source OS → target OS, encrypted, async) | ✅ **PASS** | 10.34s | Pass |
| 3 | Verify migrated data (compare source and target object lists) | ✅ **PASS** | 688ms | Pass |
| 4 | `DELETE /beetle/migration/middleware/ns/mig01/objectStorage/{targetOsId}` (cleanup) | ✅ **PASS** | 11.523s | Pass |

**Overall Result**: 4/4 steps passed ✅

**Total Duration**: 29.674535703s

*Test executed on May 8, 2026 at 20:39:44 KST (2026-05-08 20:39:44 KST) using CM-Beetle automated test CLI*

---

## Detailed Test Results

> [!NOTE]
> This section provides detailed information for each test step, including API request and response details.

### Step 1: Create target object storage

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/migration/middleware/ns/mig01/objectStorage`
- **Purpose**: Create target object storage bucket directly from JSON spec
- **Target CSP**: `ncp`
- **Target Region**: `kr`
- **Target OS ID**: `data01-cm-beetle-test-bucket`
- **Source OS ID**: `src01-cm-beetle-test-bucket` (shared pre-flight)

<details>
<summary>Request Body</summary>

```json
{
  "nameSeed": "data01",
  "status": "recommended",
  "description": "Direct creation (no recommendation step)",
  "targetCloud": {
    "csp": "ncp",
    "region": "kr"
  },
  "targetObjectStorages": [
    {
      "sourceBucketName": "test-source",
      "bucketName": "cm-beetle-test-bucket",
      "versioningEnabled": false,
      "corsEnabled": false
    }
  ]
}
```

</details>

#### 1.2 API Response Information

- **Status**: ✅ **SUCCESS** (HTTP 201 Created)

### Step 2: Migrate data (source OS → target OS, encrypted)

#### 2.1 API Request Information

- **API Endpoint**: `POST /beetle/migration/data`
- **Purpose**: Migrate data from source object storage to target object storage
- **Transfer Direction**: Source Object Storage (Tumblebug) → Target Object Storage (Tumblebug)
- **Encryption**: Tumblebug auth credentials encrypted with one-time RSA public key
- **Strategy**: auto

<details>
<summary>Request Body (sanitized — passwords masked)</summary>

```json
{
  "source": {
    "storageType": "objectstorage",
    "path": "src01-cm-beetle-test-bucket",
    "objectStorage": {
      "accessType": "tumblebug",
      "tumblebug": {
        "endpoint": "http://cb-tumblebug:1323/tumblebug",
        "nsId": "mig01",
        "osId": "src01-cm-beetle-test-bucket",
        "auth": {
          "authType": "basic",
          "username": "default",
          "password": "***"
        }
      }
    }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "data01-cm-beetle-test-bucket",
    "objectStorage": {
      "accessType": "tumblebug",
      "tumblebug": {
        "endpoint": "http://cb-tumblebug:1323/tumblebug",
        "nsId": "mig01",
        "osId": "data01-cm-beetle-test-bucket",
        "auth": {
          "authType": "basic",
          "username": "default",
          "password": "***"
        }
      }
    }
  },
  "strategy": "auto",
  "_note": "Sensitive fields (passwords) are masked with *** in this report"
}
```

</details>

#### 2.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Duration**: 10.34s
- **Note**: Initial response was `202 Accepted` (status: `Handling`). The test CLI polled `GET /beetle/request/1778240392251121859` every 10s (timeout: 600s) until migration completed with status `Success`. Step 3 was then executed.
<details>
<summary>Initial Response Body (202 Accepted)</summary>

```json
{
  "data": {
    "reqId": "1778240392251121859",
    "status": "Handling",
    "statusUrl": "/beetle/request/1778240392251121859"
  },
  "message": "Migration started. Use GET /request/{reqId} to check status.",
  "success": true
}
```

</details>

### Step 3: Verify migrated data

#### 3.1 Action

- **Purpose**: Compare source and target bucket object lists to confirm migration completeness
- **Method**: List all objects in source and target via CM-Beetle API, then compare keys

<details>
<summary>Object Lists (source and target)</summary>

```json
{
  "source": {
    "osId": "src01-cm-beetle-test-bucket",
    "count": 25,
    "objects": [
      "docs/file001.txt",
      "docs/file002.txt",
      "docs/file003.txt",
      "medium/file001.csv",
      "medium/file001.json",
      "medium/file002.csv",
      "medium/file002.json",
      "medium/file003.csv",
      "medium/file003.json",
      "medium/file004.json",
      "medium/file005.json",
      "nested/level1/level2/file001.txt",
      "nested/level1/level2/file002.txt",
      "nested/level1/level3/file001.txt",
      "nested/level1/level3/file002.txt",
      "small/file001.txt",
      "small/file002.txt",
      "small/file003.txt",
      "small/file004.txt",
      "small/file005.txt",
      "small/file006.txt",
      "small/file007.txt",
      "small/file008.txt",
      "small/file009.txt",
      "small/file010.txt"
    ]
  },
  "target": {
    "osId": "data01-cm-beetle-test-bucket",
    "count": 25,
    "objects": [
      "docs/file001.txt",
      "docs/file002.txt",
      "docs/file003.txt",
      "medium/file001.csv",
      "medium/file001.json",
      "medium/file002.csv",
      "medium/file002.json",
      "medium/file003.csv",
      "medium/file003.json",
      "medium/file004.json",
      "medium/file005.json",
      "nested/level1/level2/file001.txt",
      "nested/level1/level2/file002.txt",
      "nested/level1/level3/file001.txt",
      "nested/level1/level3/file002.txt",
      "small/file001.txt",
      "small/file002.txt",
      "small/file003.txt",
      "small/file004.txt",
      "small/file005.txt",
      "small/file006.txt",
      "small/file007.txt",
      "small/file008.txt",
      "small/file009.txt",
      "small/file010.txt"
    ]
  }
}
```

</details>

<details>
<summary>Source data tree (25 objects)</summary>

```
.
├── docs
│   ├── file001.txt
│   ├── file002.txt
│   └── file003.txt
├── medium
│   ├── file001.csv
│   ├── file001.json
│   ├── file002.csv
│   ├── file002.json
│   ├── file003.csv
│   ├── file003.json
│   ├── file004.json
│   └── file005.json
├── nested
│   └── level1
│       ├── level2
│       │   ├── file001.txt
│       │   └── file002.txt
│       └── level3
│           ├── file001.txt
│           └── file002.txt
└── small
    ├── file001.txt
    ├── file002.txt
    ├── file003.txt
    ├── file004.txt
    ├── file005.txt
    ├── file006.txt
    ├── file007.txt
    ├── file008.txt
    ├── file009.txt
    └── file010.txt
```

</details>

<details>
<summary>Target data tree (25 objects)</summary>

```
.
├── docs
│   ├── file001.txt
│   ├── file002.txt
│   └── file003.txt
├── medium
│   ├── file001.csv
│   ├── file001.json
│   ├── file002.csv
│   ├── file002.json
│   ├── file003.csv
│   ├── file003.json
│   ├── file004.json
│   └── file005.json
├── nested
│   └── level1
│       ├── level2
│       │   ├── file001.txt
│       │   └── file002.txt
│       └── level3
│           ├── file001.txt
│           └── file002.txt
└── small
    ├── file001.txt
    ├── file002.txt
    ├── file003.txt
    ├── file004.txt
    ├── file005.txt
    ├── file006.txt
    ├── file007.txt
    ├── file008.txt
    ├── file009.txt
    └── file010.txt
```

</details>

#### 3.2 Result

- **Status**: ✅ **SUCCESS**
- **Source objects**: 25
- **Target objects**: 25
- **Matched**: 25/25 ✅
- **Duration**: 688ms

### Step 4: Delete target object storage (cleanup)

#### 4.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/middleware/ns/mig01/objectStorage/data01-cm-beetle-test-bucket`
- **Purpose**: Delete target object storage bucket as cleanup
- **Note**: Always runs regardless of previous step failures

#### 4.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Target object storage deleted successfully

