# Object Storage Migration Example

This example demonstrates data migration using Object Storage with the `transx` library.

## Overview

The `transx` library supports three Object Storage access methods:

| Access Type   | Description          | Use Case                        |
| ------------- | -------------------- | ------------------------------- |
| **MinIO**     | Direct S3 SDK access | Direct connection to S3/MinIO   |
| **Spider**    | Via CB-Spider API    | Multi-cloud with presigned URLs |
| **Tumblebug** | Via CB-Tumblebug API | Multi-cloud abstraction layer   |

## Configuration Files

### Upload/Download (Local ↔ Object Storage)

| File                          | Direction  | Description            |
| ----------------------------- | ---------- | ---------------------- |
| `config-minio-upload.json`    | Local → S3 | Upload via MinIO SDK   |
| `config-minio-download.json`  | S3 → Local | Download via MinIO SDK |
| `config-spider-upload.json`   | Local → S3 | Upload via CB-Spider   |
| `config-spider-download.json` | S3 → Local | Download via CB-Spider |

### Object Storage to Object Storage

| File                                 | Description                                    |
| ------------------------------------ | ---------------------------------------------- |
| `config-tumblebug-os2os.json`        | Transfer between Object Storages via Tumblebug |
| `config-tumblebug-os2os-filter.json` | Same as above with file filtering              |

## Quick Start

### 1. Prepare Test Data

```bash
./create-test-data.sh
```

### 2. Upload Test Data to Source Object Storage

```bash
# Using MinIO (direct S3)
./migrate.sh -c config-minio-upload.json -v

# Or using Spider (via CB-Spider)
./migrate.sh -c config-spider-upload.json -v
```

### 3. Transfer Between Object Storages

```bash
# Basic transfer
./migrate.sh -c config-tumblebug-os2os.json -v

# With file filtering
./migrate.sh -c config-tumblebug-os2os-filter.json -v
```

### 4. Download to Local

```bash
./migrate.sh -c config-minio-download.json -v
```

## Configuration Examples

### MinIO (Direct S3 SDK)

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/tmp/transx-test-data/",
    "filesystem": { "accessType": "local" }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "my-bucket/backup/",
    "objectStorage": {
      "accessType": "minio",
      "minio": {
        "endpoint": "s3.ap-northeast-2.amazonaws.com",
        "accessKeyId": "AKIAXXXXXXXX",
        "secretAccessKey": "xxxxxxxx",
        "region": "ap-northeast-2",
        "useSSL": true
      }
    }
  },
  "strategy": "auto"
}
```

### Spider (CB-Spider Presigned URL)

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/tmp/transx-test-data/",
    "filesystem": { "accessType": "local" }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "my-bucket/",
    "objectStorage": {
      "accessType": "spider",
      "spider": {
        "endpoint": "http://localhost:1024/spider",
        "connectionName": "aws-connection",
        "expires": 3600
      }
    }
  },
  "strategy": "auto"
}
```

### Tumblebug (Object Storage to Object Storage)

```json
{
  "source": {
    "storageType": "objectstorage",
    "path": "source-bucket/data",
    "objectStorage": {
      "accessType": "tumblebug",
      "tumblebug": {
        "endpoint": "http://localhost:1323/tumblebug",
        "nsId": "ns01",
        "osId": "source-object-storage",
        "expires": 3600
      }
    }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "dest-bucket/data",
    "objectStorage": {
      "accessType": "tumblebug",
      "tumblebug": {
        "endpoint": "http://localhost:1323/tumblebug",
        "nsId": "ns01",
        "osId": "dest-object-storage",
        "expires": 3600
      }
    }
  },
  "strategy": "auto"
}
```

### With File Filtering

```json
{
  "source": {
    "storageType": "objectstorage",
    "path": "source-bucket/data",
    "objectStorage": { ... },
    "filter": {
      "include": ["*.csv", "*.json", "reports/**"],
      "exclude": ["*.tmp", "*.log", "temp/**", "cache/**"]
    }
  },
  "destination": { ... },
  "strategy": "auto"
}
```

## Test Workflow

```
┌─────────────────┐     Upload      ┌─────────────────┐
│   Local Data    │ ───────────────→│  Source Bucket  │
│ /tmp/test-data/ │  (minio/spider) │   (AWS S3)      │
└─────────────────┘                 └────────┬────────┘
                                             │
                                    OS2OS Transfer
                                    (tumblebug)
                                             │
                                             ▼
                                    ┌─────────────────┐
                                    │  Dest Bucket    │
                                    │   (GCP GCS)     │
                                    └────────┬────────┘
                                             │
                                      Download
                                    (minio/spider)
                                             │
                                             ▼
                                    ┌─────────────────┐
                                    │  Local Verify   │
                                    │ /tmp/download/  │
                                    └─────────────────┘
```

## Template Files

Template files with placeholder values are provided for creating custom configurations:

- `template-config-minio-upload.json`
- `template-config-minio-download.json`
- `template-config-spider-upload.json`
- `template-config-spider-download.json`
- `template-config-tumblebug-os2os.json`
- `template-config-tumblebug-os2os-filter.json`

## Prerequisites

### For MinIO Access

- S3/MinIO endpoint accessible
- AWS Access Key and Secret Key
- Bucket exists

### For Spider Access

- CB-Spider running (default: `localhost:1024`)
- Cloud connection configured in CB-Spider
- Bucket exists

### For Tumblebug Access

- CB-Tumblebug running (default: `localhost:1323`)
- Namespace created
- Object Storage resources registered
