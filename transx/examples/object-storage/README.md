# Object Storage Migration Example

This example demonstrates data migration using Object Storage with the `transx` library.

## Overview

The `transx` library supports two Object Storage handlers:

1. **Spider Handler**: Uses CB-Spider's presigned URL API (recommended)

   - No direct credentials needed (managed by CB-Spider)
   - Supports multiple cloud providers (AWS S3, Azure Blob, GCS, etc.)
   - Fast and secure

2. **MinIO Handler**: Direct S3-compatible API access
   - Requires AWS credentials (Access Key + Secret Key)
   - Standard S3 protocol
   - Direct connection to S3/MinIO servers

## Features

- **Direct Mode**: Bidirectional transfers (Local ↔ Object Storage)
- **Relay Mode**: Object Storage → Object Storage (via intermediate storage)
- **Two Handlers**: Spider (presigned URL) or MinIO SDK (direct S3)
- **Step-by-step Execution**: Supports individual backup, transfer, and restore steps
- **Multi-file Support**: Automatic handling of directory structures

## Quick Start

### Prerequisites

1. **For Spider Handler**:

   - CB-Spider running on `localhost:1024`
   - Cloud connection configured in CB-Spider
   - Bucket exists

2. **For MinIO Handler**:
   - S3/MinIO endpoint accessible
   - AWS Access Key and Secret Key
   - Bucket exists

### Configuration Files

Main configuration files are provided:

| File                                      | Handler   | Direction      | Description                   |
| ----------------------------------------- | --------- | -------------- | ----------------------------- |
| `config-spider-upload.json`               | Spider    | Local → S3     | Upload via CB-Spider          |
| `config-spider-download.json`             | Spider    | S3 → Local     | Download via CB-Spider        |
| `config-minio-upload.json`                | MinIO SDK | Local → S3     | Upload direct to S3           |
| `config-minio-download.json`              | MinIO SDK | S3 → Local     | Download direct from S3       |
| `config-rsync-upload.json`                | rsync     | Local → Remote | Upload via rsync/SSH          |
| `config-rsync-download.json`              | rsync     | Remote → Local | Download via rsync/SSH        |
| **Filtering Examples:**                   |           |                |                               |
| `config-filtering-combined-upload.json`   | Spider    | Local → S3     | Upload with include/exclude   |
| `config-filtering-combined-download.json` | MinIO     | S3 → Local     | Download with include/exclude |
| `config-filtering-exclude-logs.json`      | MinIO     | S3 → Local     | Exclude log/tmp files         |
| `config-filtering-include-data.json`      | MinIO     | S3 → Local     | Include only JSON/CSV         |
| `config-filtering-exclude-dirs.json`      | Spider    | Local → S3     | Exclude common directories    |
| `config-rsync-filtering.json`             | rsync     | Local → Remote | Rsync with include/exclude    |

### Running Tests

```bash
# Spider handler (via CB-Spider)
./migrate.sh -c config-spider-upload.json -v
./migrate.sh -c config-spider-download.json -v

# MinIO SDK (direct S3 connection)
./migrate.sh -c config-minio-upload.json -v
./migrate.sh -c config-minio-download.json -v

# rsync (SSH-based file transfer)
./migrate.sh -c config-rsync-upload.json -v
./migrate.sh -c config-rsync-download.json -v

# File filtering examples
./migrate.sh -c config-filtering-combined-upload.json -v
./migrate.sh -c config-filtering-combined-download.json -v
./migrate.sh -c config-filtering-exclude-logs.json -v
./migrate.sh -c config-filtering-include-data.json -v
./migrate.sh -c config-filtering-exclude-dirs.json -v
./migrate.sh -c config-rsync-filtering.json -v

# Step-by-step execution
./migrate.sh -c config-spider-upload.json -s transfer -v
```

## Configuration Examples

### Spider Handler Configuration

**Upload (Local → Object Storage)**:

```json
{
  "source": {
    "endpoint": "",
    "dataPath": "/tmp/minio-test-data/"
  },
  "destination": {
    "endpoint": "http://localhost:1024/spider/s3",
    "dataPath": "my-bucket/path/"
  },
  "destinationTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "handler": "spider",
      "accessKeyId": "aws-ap-northeast-2",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3
    }
  }
}
```

**Download (Object Storage → Local)**:

```json
{
  "source": {
    "endpoint": "http://localhost:1024/spider/s3",
    "dataPath": "my-bucket/path/"
  },
  "destination": {
    "endpoint": "",
    "dataPath": "/tmp/download/"
  },
  "sourceTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "handler": "spider",
      "accessKeyId": "aws-ap-northeast-2",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3
    }
  }
}
```

### MinIO SDK Configuration

**Upload (Local → S3)**:

```json
{
  "source": {
    "endpoint": "",
    "dataPath": "/tmp/minio-test-data/"
  },
  "destination": {
    "endpoint": "http://s3.ap-northeast-2.amazonaws.com",
    "dataPath": "my-bucket/path/"
  },
  "destinationTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "handler": "minio",
      "accessKeyId": "AKIAIOSFODNN7EXAMPLE",
      "secretAccessKey": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
      "region": "ap-northeast-2",
      "useSSL": true,
      "timeout": 300,
      "maxRetries": 3
    }
  }
}
```

**Download (S3 → Local)**:

```json
{
  "source": {
    "endpoint": "http://s3.ap-northeast-2.amazonaws.com",
    "dataPath": "my-bucket/path/"
  },
  "destination": {
    "endpoint": "",
    "dataPath": "/tmp/download/"
  },
  "sourceTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "handler": "minio",
      "accessKeyId": "AKIAIOSFODNN7EXAMPLE",
      "secretAccessKey": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
      "region": "ap-northeast-2",
      "useSSL": true,
      "timeout": 300,
      "maxRetries": 3
    }
  }
}
```

### Rsync Configuration

**Upload (Local → Remote Server)**:

```json
{
  "source": {
    "endpoint": "",
    "dataPath": "/tmp/minio-test-data/"
  },
  "destination": {
    "endpoint": "user@remote-server.example.com",
    "dataPath": "/tmp/rsync-upload-test/"
  },
  "destinationTransferOptions": {
    "method": "rsync",
    "rsyncOptions": {
      "username": "user",
      "sshPrivateKeyPath": "~/.ssh/id_rsa",
      "insecureSkipHostKeyVerification": false,
      "connectTimeout": 30,
      "verbose": false,
      "progress": true,
      "compress": true,
      "archive": true,
      "delete": false,
      "exclude": ["*.tmp", "*.log"],
      "transferDirContentsOnly": true
    }
  }
}
```

**Download (Remote Server → Local)**:

```json
{
  "source": {
    "endpoint": "user@remote-server.example.com",
    "dataPath": "/tmp/rsync-upload-test/"
  },
  "destination": {
    "endpoint": "",
    "dataPath": "/tmp/rsync-download-test/"
  },
  "sourceTransferOptions": {
    "method": "rsync",
    "rsyncOptions": {
      "username": "user",
      "sshPrivateKeyPath": "~/.ssh/id_rsa",
      "insecureSkipHostKeyVerification": false,
      "connectTimeout": 30,
      "verbose": false,
      "progress": true,
      "compress": true,
      "archive": true,
      "delete": false,
      "exclude": ["*.tmp", "*.log"],
      "transferDirContentsOnly": true
    }
  }
}
```

## Configuration Fields

### Common Fields

- **`endpoint`**:
  - Spider: `http://localhost:1024/spider/s3`
  - MinIO: S3 endpoint (e.g., `http://s3.amazonaws.com`)
  - Local: Empty string `""`
- **`dataPath`**:
  - Object Storage: `bucket-name/path/`
  - Local: `/absolute/path/`

### Spider Handler Options

- **`handler`**: `"spider"` (default)
- **`accessKeyId`**: CB-Spider connection name (e.g., `"aws-ap-northeast-2"`)
- **`expiresIn`**: Presigned URL expiration in seconds (default: 3600)
- **`timeout`**: Request timeout in seconds (default: 300)
- **`maxRetries`**: Retry attempts (default: 3)

### MinIO SDK Options

- **`handler`**: `"minio"` (**required**)
- **`accessKeyId`**: AWS Access Key ID (**required**)
- **`secretAccessKey`**: AWS Secret Access Key (**required**)
- **`region`**: AWS region (default: `"us-east-1"`)
- **`useSSL`**: Use HTTPS (default: `true`, **required for AWS S3**)
- **`timeout`**: Request timeout in seconds (default: 300)
- **`maxRetries`**: Retry attempts (default: 3)

### Rsync Options

- **`username`**: SSH username for remote server
- **`sshPrivateKeyPath`**: Path to SSH private key (e.g., `~/.ssh/id_rsa`)
- **`insecureSkipHostKeyVerification`**: Skip SSH host key verification (default: `false`)
- **`connectTimeout`**: SSH connection timeout in seconds (default: 30)
- **`verbose`**: Enable verbose output (default: `false`)
- **`progress`**: Show progress during transfer (default: `true`)
- **`compress`**: Compress data during transfer (default: `true`)
- **`archive`**: Archive mode, preserves permissions (default: `true`)
- **`delete`**: Delete files in destination not in source (default: `false`)
- **`filter`**: File filtering options with `include`/`exclude` patterns (e.g., `{"exclude": ["*.tmp", "*.log"]}`)
- **`transferDirContentsOnly`**: Transfer directory contents only (default: `true`)

## Handler Comparison

| Feature                   | Spider Handler              | MinIO SDK          | Rsync              |
| ------------------------- | --------------------------- | ------------------ | ------------------ |
| **Authentication**        | CB-Spider connection name   | AWS credentials    | SSH key            |
| **Credential Management** | Managed by CB-Spider        | User managed       | User managed       |
| **SSL/TLS**               | Handled by Spider           | Must configure     | SSH encrypted      |
| **Multi-cloud Support**   | Yes (AWS, Azure, GCS, etc.) | S3-compatible only | Any SSH server     |
| **Direct S3 Access**      | No (via Spider)             | Yes                | No                 |
| **Protocol**              | HTTP/HTTPS                  | S3 API             | SSH/rsync          |
| **Performance**           | ~23% faster (test result)   | Baseline           | Depends on network |
| **Use Case**              | Cloud object storage        | Direct S3 access   | Remote servers     |

## Advanced Usage

### Step-by-step Execution

```bash
# Execute only specific steps
./migrate.sh -c config.json -s backup    # Backup only
./migrate.sh -c config.json -s transfer  # Transfer only
./migrate.sh -c config.json -s restore   # Restore only
```

### Multiple Files and Directories

The library automatically handles:

- Directory uploads/downloads
- Prefix-based listing
- Directory structure preservation

```json
{
  "source": {
    "dataPath": "/tmp/backup-dir/" // Will upload all files in directory
  },
  "destination": {
    "dataPath": "my-bucket/backups/" // All files uploaded to this prefix
  }
}
```

## Legacy Configuration Files

Additional configuration files for reference:

- `config-basic-direct.json`: Basic remote to local
- `config-basic-relay.json`: Basic relay mode
- `config-local-to-objectstorage.json`: Local to Object Storage
- `config-objectstorage-to-local.json`: Object Storage to local
- `config-objectstorage-to-objectstorage.json`: Object Storage migration
- `config-rsync-to-objectstorage.json`: Rsync to Object Storage

## Troubleshooting

### Spider Handler Issues

**Connection refused:**

```bash
# Check if CB-Spider is running
curl http://localhost:1024/spider/readyz
```

**Invalid connection:**

- Verify connection name exists in CB-Spider
- Check CB-Spider connection configuration

### MinIO SDK Issues

**Authentication failed:**

- Verify `accessKeyId` and `secretAccessKey` are correct
- Ensure credentials have appropriate permissions

**SSL/TLS errors:**

- Set `"useSSL": true` for AWS S3
- Set `"useSSL": false` for local MinIO without SSL

**No objects found:**

- Verify bucket name and path are correct
- Check bucket exists and is accessible
- Ensure credentials have list permissions

### Rsync Issues

**Connection refused:**

- Verify remote server is accessible
- Check SSH service is running on remote server
- Verify username and hostname are correct

**Permission denied:**

- Ensure SSH key has correct permissions (600)
- Verify SSH key is authorized on remote server
- Check user has read/write permissions on remote paths

**Host key verification failed:**

- Add remote host to `~/.ssh/known_hosts`
- Or set `"insecureSkipHostKeyVerification": true` (not recommended for production)

## Testing

### Test Data Setup

#### Quick Test Data

```bash
# Create simple test data
mkdir -p /tmp/minio-test-data
echo "test content" > /tmp/minio-test-data/test-file.txt
```

#### Comprehensive Test Data (Recommended)

```bash
# Create test data with multiple depths and file types
./create-test-data.sh /tmp/transx-test-data

# This creates a structure with:
# - 4 directory depth levels
# - Multiple file types (.txt, .json, .csv, .log, .tmp, .go, .md, .pdf, etc.)
# - ~70+ files across various directories
# - Common directories to exclude (.git, node_modules, backup, temp)
```

The test data includes:

- **Root level**: Config files, documentation
- **Level 1**: src/, docs/, data/, logs/, temp/, backup/
- **Level 2**: Nested subdirectories (models/, tests/, images/, raw/, processed/)
- **Level 3**: Deep structures (entities/, 2025/, reports/)
- **Level 4**: Very deep nesting (Q1/, monthly data)

### Run Tests

```bash
# Basic upload/download tests
./migrate.sh -c config-spider-upload.json -v
./migrate.sh -c config-spider-download.json -v

# Filtering tests
# MinIO download with filtering (exclude logs)
./migrate.sh -c config-filtering-exclude-logs.json -v

# MinIO download with filtering (include only data files)
./migrate.sh -c config-filtering-include-data.json -v

# Spider upload with filtering (exclude directories)
./migrate.sh -c config-filtering-exclude-dirs.json -v

# Spider upload with filtering (combined include/exclude)
./migrate.sh -c config-filtering-combined.json -v

# Verify data integrity
md5sum /tmp/transx-test-data/*
md5sum /tmp/transx-download-test/*
```

### Filtering Test Configurations

Four filtering configuration files demonstrate filtering in different scenarios:

| File                                 | Description                  | Handler | Direction  | Filters                                                                             |
| ------------------------------------ | ---------------------------- | ------- | ---------- | ----------------------------------------------------------------------------------- |
| `config-filtering-exclude-logs.json` | Exclude log and temp files   | MinIO   | S3 → Local | `"exclude": ["*.log", "*.tmp"]`                                                     |
| `config-filtering-include-data.json` | Include only data files      | MinIO   | S3 → Local | `"include": ["*.json", "*.csv"]`                                                    |
| `config-filtering-exclude-dirs.json` | Exclude common directories   | Spider  | Local → S3 | `"exclude": [".git/*", "node_modules/*", "backup/*", "temp/*"]`                     |
| `config-filtering-combined.json`     | Combined include and exclude | Spider  | Local → S3 | `"include": ["data/*", "docs/*"]`<br>`"exclude": ["*.log", "*.tmp", "*/archive/*"]` |

**Scenarios:**

- **MinIO Download**: Tests filtering when downloading from S3-compatible storage (sourceTransferOptions)
- **Spider Upload**: Tests filtering when uploading to Object Storage via CB-Spider (destinationTransferOptions)

### Verify Filtering Results

```bash
# Count files before and after filtering
echo "Source files:"
find /tmp/transx-test-data -type f | wc -l

echo "Files matching include pattern (*.json, *.csv):"
find /tmp/transx-test-data -type f \( -name "*.json" -o -name "*.csv" \) | wc -l

echo "Files after excluding (*.log, *.tmp):"
find /tmp/transx-test-data -type f ! \( -name "*.log" -o -name "*.tmp" \) | wc -l
```

## Integration Example

## Integration Example

```go
package main

import (
    "log"
    "github.com/cloud-barista/cm-beetle/transx"
)

func main() {
    // Load configuration from file
    task, err := loadConfig("config-spider-upload.json")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Execute migration
    err = transx.MigrateData(task)
    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    }

    log.Println("Migration completed successfully")
}
```

## File Filtering

Object Storage transfers support include/exclude patterns for selective file transfer, similar to rsync's filtering mechanism.

### Filtering Options

Add `filter` object with `include` and/or `exclude` arrays to `objectStorageTransferOptions`:

```json
{
  "objectStorageTransferOptions": {
    "handler": "spider",
    "accessKeyId": "aws-config01",
    "filter": {
      "exclude": ["*.log", "*.tmp", ".git/**"],
      "include": ["*.json", "*.csv"]
    }
  }
}
```

### Filtering Logic

1. **Include patterns** (whitelist): If specified, only objects matching at least one pattern are transferred
2. **Exclude patterns** (blacklist): **Always takes precedence** - matching files are excluded even if they match include patterns
3. **No patterns**: All objects are transferred

**Important**: Exclude has priority over Include for safety (prevents accidental transfer of sensitive files)

### Pattern Syntax

Supports glob patterns with recursive matching:

- `*` - matches any sequence of characters (in single path segment)
- `?` - matches any single character
- `[...]` - matches any character in brackets
- `**` - matches zero or more directories recursively (like rsync)

**Pattern Examples:**

| Pattern          | Matches                                  | Description                       |
| ---------------- | ---------------------------------------- | --------------------------------- |
| `*.log`          | `app.log`, `error.log`, `dir/test.log`   | Any .log file (any depth)         |
| `data/*`         | `data/file.txt`                          | Files directly in data/           |
| `data/**`        | `data/file.txt`, `data/sub/file.txt`     | All files under data/ recursively |
| `data/**/*.json` | `data/config.json`, `data/sub/data.json` | All .json files under data/       |
| `.git/**`        | `.git/config`, `.git/objects/abc`        | Everything under .git/            |
| `**/test/**`     | `src/test/file.go`, `lib/test/util.go`   | Any path containing /test/        |

### Examples

**Exclude log and temporary files (recursive):**

```json
{
  "filter": {
    "exclude": ["*.log", "*.tmp", "temp/**"]
  }
}
```

**Include only data and docs directories (recursive):**

```json
{
  "filter": {
    "include": ["data/**", "docs/**"]
  }
}
```

**Include only specific file types:**

```json
{
  "filter": {
    "include": ["*.pdf", "*.docx", "*.txt"]
  }
}
```

**Exclude directories recursively:**

```json
{
  "filter": {
    "exclude": [".git/**", "node_modules/**", "backup/**"]
  }
}
```

**Combine include and exclude (recursive):**

```json
{
  "filter": {
    "include": ["data/**", "docs/**"],
    "exclude": ["*.log", "*.tmp", "**/archive/**"]
  }
}
```

**Result**:

- ✅ Transfers: `data/config.json`, `data/sub/output.csv`, `docs/readme.txt`
- ❌ Excludes: `data/process.log`, `data/backup.tmp`, `data/archive/old.json`

**Explanation**: Only files under `data/` or `docs/` are included (recursive), but log/tmp files and anything in archive directories are excluded.

### When Filtering is Applied

- **Download**: After listing objects, before transfer
- **Upload**: During file system walk, before transfer
- **Performance**: In-memory filtering (no extra API calls)

## Performance Notes

Based on test results with 4 files (~11.5KB total):

- **Spider Handler**: ~590ms average (23% faster)
- **MinIO SDK**: ~770ms average

Both handlers provide 100% data integrity with identical MD5 checksums.

## References

- [CB-Spider Documentation](https://github.com/cloud-barista/cb-spider)
- [MinIO Go Client](https://github.com/minio/minio-go)
- [transx Library](../../README.md)
