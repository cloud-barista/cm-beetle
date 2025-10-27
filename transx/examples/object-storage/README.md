# Object Storage Migration Example

This example demonstrates how to perform data migration using Object Storage with the `transx` library and CB-Spider integration.

## Overview

The `transx` library supports Object Storage migration through CB-Spider presigned URLs, providing a unified API for various cloud object storage services (AWS S3, Azure Blob, GCS, etc.).

## Features

- **Direct Mode**: Bidirectional transfers (Local ↔ Object Storage)
- **Relay Mode**: Object Storage → Object Storage (via intermediate storage)
- **CB-Spider Integration**: Uses CB-Spider's presigned URL API for secure transfers
- **Step-by-step Execution**: Supports individual backup, transfer, and restore steps
- **Multi-file Support**: Automatic handling of directory structures and multiple files

## Configuration

### Migration Configuration Files

The library uses simplified configuration files that specify endpoints and transfer options. Multiple pre-configured examples are provided for different migration scenarios.

### Direct Mode Configurations

#### 1. Remote to Local (`config-objectstorage-to-local.json`)

Downloads data from Object Storage to local filesystem:

```json
{
  "source": {
    "endpoint": "http://localhost:1024/spider/s3",
    "dataPath": "spider-test-bucket/backup/data.sql",
    "backupCmd": "echo 'Sample backup data' > /tmp/data.sql"
  },
  "destination": {
    "endpoint": "",
    "dataPath": "/tmp/object-storage-migration/data.sql",
    "restoreCmd": "cat /tmp/object-storage-migration/data.sql"
  },
  "sourceTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "accessKeyId": "conn-kimy-aws",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3,
      "verifySSL": false
    }
  },
  "destinationTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "accessKeyId": "conn-kimy-aws",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3,
      "verifySSL": false
    }
  }
}
```

#### 2. Local to Remote (`config-local-to-objectstorage.json`)

Uploads local files to Object Storage:

```json
{
  "source": {
    "endpoint": "",
    "dataPath": "/tmp/local-data/backup.sql",
    "backupCmd": "echo 'Local backup data created at $(date)' > /tmp/local-data/backup.sql"
  },
  "destination": {
    "endpoint": "http://localhost:1024/spider/s3",
    "dataPath": "spider-test-bucket/uploads/backup.sql"
  },
  "sourceTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "accessKeyId": "conn-kimy-aws",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3,
      "verifySSL": false
    }
  },
  "destinationTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "accessKeyId": "conn-kimy-aws",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3,
      "verifySSL": false
    }
  }
}
```

### Relay Mode Configurations

#### 1. Basic Relay Mode (`config-basic-relay.json`)

Transfers data between different Object Storage systems:

```json
{
  "source": {
    "endpoint": "http://localhost:1024/spider/s3",
    "dataPath": "source-bucket/backup/data.sql",
    "backupCmd": "echo 'Sample source data' > /tmp/source-data.sql"
  },
  "destination": {
    "endpoint": "http://localhost:1024/spider/s3",
    "dataPath": "dest-bucket/backup/data.sql",
    "restoreCmd": "cat /tmp/dest-data.sql"
  },
  "sourceTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "accessKeyId": "conn-source-aws",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3,
      "verifySSL": false
    }
  },
  "destinationTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "accessKeyId": "conn-dest-aws",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3,
      "verifySSL": false
    }
  }
}
```

#### 2. Object Storage to Object Storage (`config-relay-objectstorage.json`)

Advanced Object Storage to Object Storage migration with different connection configurations:

```json
{
  "source": {
    "endpoint": "http://localhost:1024/spider/s3",
    "dataPath": "source-bucket/backup/database-backup.sql"
  },
  "destination": {
    "endpoint": "http://localhost:1024/spider/s3",
    "dataPath": "destination-bucket/migrated/database-backup.sql",
    "restoreCmd": "echo 'Data migration to destination bucket completed successfully'"
  },
  "sourceTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "accessKeyId": "conn-source-aws",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3,
      "verifySSL": false
    }
  },
  "destinationTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "accessKeyId": "conn-destination-aws",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3,
      "verifySSL": false
    }
  }
}
```

#### 3. Rsync to Object Storage (`rsync-to-object-storage-relay-mode-config.json`)

Transfers data from remote server via rsync to Object Storage:

```json
{
  "source": {
    "endpoint": "user@remote-server.example.com",
    "dataPath": "/var/backups/database/backup.sql",
    "backupCmd": "mysqldump -u backup_user -p'backup_password' production_db > /var/backups/database/backup.sql"
  },
  "destination": {
    "endpoint": "http://localhost:1024/spider/s3",
    "dataPath": "backup-storage-bucket/rsync-backups/backup.sql",
    "restoreCmd": "echo 'Backup successfully uploaded to Object Storage via rsync relay'"
  },
  "sourceTransferOptions": {
    "method": "rsync",
    "rsyncOptions": {
      "sshKeyPath": "~/.ssh/id_rsa",
      "sshPort": 22,
      "timeout": 300,
      "maxRetries": 3,
      "verbose": false,
      "preservePermissions": true,
      "compressData": true,
      "deleteAfterTransfer": false,
      "additionalOptions": ["--exclude=*.tmp", "--exclude=*.log"]
    }
  },
  "destinationTransferOptions": {
    "method": "object-storage-api",
    "objectStorageTransferOptions": {
      "accessKeyId": "conn-backup-aws",
      "expiresIn": 3600,
      "timeout": 300,
      "maxRetries": 3,
      "verifySSL": false
    }
  }
}
```

### Configuration Fields

#### Required Fields

- **`endpoint`**: CB-Spider API endpoint (e.g., `"http://localhost:1024/spider/s3"`)
- **`dataPath`**: Bucket and object path (e.g., `"bucket-name/path/to/file.sql"`)
- **`accessKeyId`**: CB-Spider connection name (e.g., `"conn-kimy-aws"`)

#### Optional Fields

- **`expiresIn`**: Presigned URL expiration in seconds (default: 3600)
- **`timeout`**: HTTP request timeout in seconds (default: 300)
- **`maxRetries`**: Maximum retry attempts (default: 3)
- **`verifySSL`**: SSL certificate verification (default: false for CB-Spider)

## Usage

### Using the Shell Script

The provided shell script offers an easy way to run migrations:

```bash
# Direct Mode Examples
./migrate.sh                                                # Default: remote-to-local
./migrate.sh -c remote-to-local-direct-mode-config.json -v  # Download from Object Storage
./migrate.sh -c local-to-remote-direct-mode-config.json -v  # Upload to Object Storage

# Relay Mode Examples
# Direct mode examples
./migrate.sh -c config-objectstorage-to-local.json -v       # Download from Object Storage
./migrate.sh -c config-local-to-objectstorage.json -v       # Upload to Object Storage

# Relay mode examples
./migrate.sh -c config-basic-relay.json -v                  # Basic relay mode
./migrate.sh -c config-relay-objectstorage.json -v          # Object Storage migration

# Step-by-step execution
./migrate.sh -s backup    # Run only backup step
./migrate.sh -s transfer  # Run only transfer step
./migrate.sh -s restore   # Run only restore step
```

### Using the Go Program Directly

```bash
# Build the program
go build -o main .

# Run direct mode
./main -config direct-mode-config.json -verbose

# Run relay mode
./main -config relay-mode-config.json -verbose

# Run specific steps
./main -config direct-mode-config.json -backup
./main -config direct-mode-config.json -transfer
./main -config direct-mode-config.json -restore
```

## CB-Spider Setup

### Prerequisites

1. CB-Spider server running on `localhost:1024`
2. Configured cloud connections (AWS, Azure, GCS, etc.)
3. Valid connection names for Object Storage access

### Connection Configuration Example

For AWS S3 through CB-Spider:

```json
{
  "ConnectionName": "conn-kimy-aws",
  "ProviderName": "AWS",
  "DriverName": "aws-driver-v1.0.so",
  "CredentialName": "aws-credential01",
  "RegionName": "us-east-1"
}
```

## Migration Modes

### Direct Mode

- **Upload (Local → Object Storage)**: Local filesystem to Object Storage
- **Download (Object Storage → Local)**: Object Storage to local filesystem
- **Use Case**: File uploads, downloading backups, data export/import

### Relay Mode

- **Source**: Object Storage (via CB-Spider presigned URL)
- **Destination**: Object Storage (via CB-Spider presigned URL)
- **Use Case**: Cross-cloud migration, backup replication between different Object Storage providers

## Features and Capabilities

### Automatic Bucket Validation

The library automatically validates bucket existence before attempting transfers.

### Multi-file Support

When downloading directories or using prefix patterns, the library automatically:

- Lists all objects matching the prefix
- Downloads each file individually
- Maintains directory structure

### URL Encoding Handling

Presigned URLs are automatically decoded to handle XML encoding issues (e.g., `&amp;` → `&`).

### Error Handling and Retries

- Configurable retry logic with exponential backoff
- Detailed error reporting
- Timeout handling for long-running transfers

## Testing

### Prerequisites

Before running the example, ensure:

1. CB-Spider is running and accessible
2. Cloud provider connections are configured
3. Target buckets exist and are accessible
4. Network connectivity to CB-Spider endpoint

### Running Tests

```bash
# Test direct mode (Object Storage to local)
./migrate.sh -v

# Test relay mode (Object Storage to Object Storage)
./migrate.sh -m relay -v

# Test individual steps
./migrate.sh -s backup -v
./migrate.sh -s transfer -v
./migrate.sh -s restore -v
```

### Troubleshooting

**Connection Issues:**

- Verify CB-Spider is running on the configured endpoint
- Check connection name exists in CB-Spider
- Validate network connectivity

**Authentication Errors:**

- Ensure connection credentials are properly configured in CB-Spider
- Verify connection name matches the configuration

**Bucket Access Issues:**

- Confirm bucket exists and is accessible
- Check IAM permissions for the configured credentials
- Verify bucket region matches the connection configuration

## Integration

This example can be easily integrated into larger data migration workflows:

```go
import "github.com/yunkon-kim/transx"

// Load configuration
config := loadConfig("my-config.json")

// Execute migration
err := transx.MigrateData(config)
if err != nil {
    log.Fatalf("Migration failed: %v", err)
}
```
