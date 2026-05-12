# Data Migration Feature Guide

## Overview

CM-Beetle provides a data migration API that transfers data between storage systems across cloud providers. This feature enables you to:

- **Migrate** files and objects between remote filesystems (SSH/rsync) and object storages (S3-compatible)
- **Filter** files by include/exclude glob patterns during transfer
- **Encrypt** sensitive credentials in transit using RSA-OAEP + AES-256-GCM
- **Track** migration progress via asynchronous request status polling

Data migration is independent of infrastructure migration. You can migrate data between any two supported endpoints regardless of which CSP hosts them.

> [!NOTE]
> This feature is **incubating**. APIs may change in future versions.

---

## Supported Transfer Combinations

The transfer planner selects the appropriate pipeline based on source and destination `storageType`:

| Source             | Destination        | Pipeline       | Transfer method                               |
| ------------------ | ------------------ | -------------- | --------------------------------------------- |
| `filesystem` (SSH) | `filesystem` (SSH) | Filesystem     | rsync (direct or relay via CM-Beetle host)    |
| `filesystem` (SSH) | `objectstorage`    | Cross-storage  | rsync pull to staging → S3 upload             |
| `objectstorage`    | `objectstorage`    | Object Storage | Presigned URL download → presigned URL upload |
| `objectstorage`    | `filesystem` (SSH) | Cross-storage  | S3 download → rsync push                      |

> [!NOTE]
> Local filesystem access (`accessType: local`) is **not allowed** via the API for security reasons.
> The CM-Beetle server rejects any request where source or destination refers to the server's local filesystem.

### Object storage access types

| `accessType` | Description                                                                           |
| ------------ | ------------------------------------------------------------------------------------- |
| `tumblebug`  | Access via CB-Tumblebug presigned URL API (recommended for Cloud-Barista deployments) |
| `spider`     | Access via CB-Spider presigned URL API                                                |
| `minio`      | Direct S3 SDK access using minio-go (requires S3 credentials)                         |

---

## Transfer Strategies

The `strategy` field controls how the transfer is orchestrated when both endpoints are SSH filesystems:

| Strategy | Description                                                                                               |
| -------- | --------------------------------------------------------------------------------------------------------- |
| `auto`   | Default. Automatically selects the most efficient method (agent forwarding if possible, relay otherwise). |
| `direct` | Force direct SSH agent forwarding between source and destination servers.                                 |
| `relay`  | Force transfer via the CM-Beetle host as a relay (download to staging, then upload).                      |

For all other combinations (cross-storage, object-to-object), the strategy field has no effect.

---

## API Reference

### GET `/beetle/migration/data/encryptionKey`

Returns a one-time RSA public key bundle. Use this to encrypt sensitive fields before sending a migration request.

**Response `200 OK`:**

```json
{
  "success": true,
  "data": {
    "keyId": "key-abc123",
    "algorithm": "RSA-OAEP-256+AES-256-GCM",
    "publicKey": "-----BEGIN PUBLIC KEY-----\n...\n-----END PUBLIC KEY-----",
    "expiresAt": "2026-05-08T12:30:00Z"
  }
}
```

- Keys expire **30 minutes** after generation.
- Keys are **one-time use** — invalidated immediately after the server decrypts a request.

---

### POST `/beetle/migration/data`

Start a data migration. Returns `202 Accepted` immediately; migration runs asynchronously.

**Request headers:**

| Header         | Required | Description                                                        |
| -------------- | -------- | ------------------------------------------------------------------ |
| `X-Request-Id` | No       | Custom request ID. Auto-generated if omitted. Used to poll status. |

**Request body (`DataMigrationModel`):**

```json
{
  "source": { ... },
  "destination": { ... },
  "strategy": "auto",
  "encryptionKeyId": ""
}
```

**Response `202 Accepted`:**

```json
{
  "success": true,
  "message": "Migration started. Use GET /request/{reqId} to check status.",
  "data": {
    "reqId": "req-xyz789",
    "status": "Handling",
    "statusUrl": "/beetle/request/req-xyz789"
  }
}
```

**Poll for status:**

```
GET /beetle/request/{reqId}
```

```json
{
  "success": true,
  "data": {
    "status": "Success",
    "responseData": {
      "message": "Data migrated successfully (4.231s)",
      "elapsedTime": "4.231s"
    }
  }
}
```

Status values: `Handling` → `Success` | `Error`

---

### POST `/beetle/migration/data/test/encrypt` _(test only)_

Encrypts a plaintext `DataMigrationModel` server-side for testing the encryption workflow.

> [!WARNING]
> **Do not use in production.** Sending plaintext credentials to the server defeats the purpose of encryption. Use client-side encryption in production.

**Request body:**

```json
{
  "publicKeyBundle": { "keyId": "...", "publicKey": "..." },
  "model": { "source": { ... }, "destination": { ... } }
}
```

**Response `200 OK`:** Returns the encrypted `DataMigrationModel` with `encryptionKeyId` set.

---

### POST `/beetle/migration/data/test/decrypt` _(test only)_

Decrypts an encrypted `DataMigrationModel` without executing migration. Use to verify the encryption round-trip.

> [!NOTE]
> Consumes the encryption key. Generate a new key before calling `POST /migration/data`.

---

## Encryption Workflow

Sensitive fields (SSH private keys, S3 credentials, auth passwords/tokens) are encrypted using a hybrid scheme:

1. **RSA-OAEP-256** to exchange a random AES key
2. **AES-256-GCM** to encrypt the sensitive field values

### Production workflow (client-side encryption)

```
1. GET  /beetle/migration/data/encryptionKey   → get publicKey + keyId
2. [Client] encrypt sensitive fields using publicKey
3. POST /beetle/migration/data                 → send encrypted model (encryptionKeyId set)
4. GET  /beetle/request/{reqId}                → poll until Success or Error
```

### Development workflow (server-side test)

```
1. GET  /beetle/migration/data/encryptionKey        → get publicKey + keyId
2. POST /beetle/migration/data/test/encrypt         → server encrypts for you
3. POST /beetle/migration/data/test/decrypt         → verify decryption (optional)
4. GET  /beetle/migration/data/encryptionKey        → get NEW key (previous key was consumed)
5. POST /beetle/migration/data/test/encrypt         → re-encrypt with new key
6. POST /beetle/migration/data                      → run actual migration
```

For Go clients, use the `transx` package helpers:

```go
bundle, _ := transx.ParsePublicKeyBundle(keyBundle)
encryptedModel, _ := transx.EncryptModel(model, bundle, keyBundle.KeyID)
```

---

## Request Body Reference

### Object storage → Object storage (via Tumblebug)

```json
{
  "source": {
    "storageType": "objectstorage",
    "path": "source-os-id",
    "objectStorage": {
      "accessType": "tumblebug",
      "tumblebug": {
        "endpoint": "http://cb-tumblebug:1323/tumblebug",
        "nsId": "mig01",
        "osId": "source-os-id",
        "auth": {
          "authType": "basic",
          "basic": { "username": "default", "password": "default" }
        }
      }
    }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "target-os-id",
    "objectStorage": {
      "accessType": "tumblebug",
      "tumblebug": {
        "endpoint": "http://cb-tumblebug:1323/tumblebug",
        "nsId": "mig01",
        "osId": "target-os-id",
        "auth": {
          "authType": "basic",
          "basic": { "username": "default", "password": "default" }
        }
      }
    }
  },
  "strategy": "auto"
}
```

### SSH filesystem → Object storage (via Tumblebug)

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/data/app",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "192.168.1.10",
        "port": 22,
        "username": "ubuntu",
        "privateKey": "-----BEGIN RSA PRIVATE KEY-----\n...\n-----END RSA PRIVATE KEY-----"
      }
    },
    "filter": {
      "include": ["**/*.txt", "**/*.json"],
      "exclude": ["**/*.log"]
    }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "target-os-id",
    "objectStorage": {
      "accessType": "tumblebug",
      "tumblebug": {
        "endpoint": "http://cb-tumblebug:1323/tumblebug",
        "nsId": "mig01",
        "osId": "target-os-id",
        "auth": {
          "authType": "basic",
          "basic": { "username": "default", "password": "default" }
        }
      }
    }
  },
  "strategy": "auto"
}
```

### SSH filesystem → SSH filesystem

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/data/app",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "source-server.example.com",
        "port": 22,
        "username": "ubuntu",
        "privateKey": "-----BEGIN RSA PRIVATE KEY-----\n...\n-----END RSA PRIVATE KEY-----",
        "archive": true,
        "compress": true
      }
    }
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/data/migrated",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "target-server.example.com",
        "port": 22,
        "username": "ubuntu",
        "privateKey": "-----BEGIN RSA PRIVATE KEY-----\n...\n-----END RSA PRIVATE KEY-----"
      }
    }
  },
  "strategy": "auto"
}
```

### Object storage via direct S3 credentials (minio)

```json
{
  "source": {
    "storageType": "objectstorage",
    "path": "my-source-bucket",
    "objectStorage": {
      "accessType": "minio",
      "minio": {
        "endpoint": "s3.amazonaws.com",
        "accessKeyId": "AKIAIOSFODNN7EXAMPLE",
        "secretAccessKey": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
        "region": "ap-northeast-2",
        "useSSL": true
      }
    }
  },
  "destination": { ... },
  "strategy": "auto"
}
```

---

## File Filtering

Apply glob patterns to control which files are transferred. Filtering is supported on the source location.

```json
"filter": {
  "include": ["**/*.txt", "**/*.json", "**/*.csv"],
  "exclude": ["**/*.log", "temp/**"]
}
```

- `include`: Only transfer files matching at least one pattern. Empty means all files.
- `exclude`: Skip files matching any pattern. Applied after include.
- Patterns follow standard glob syntax (`*`, `**`, `?`, `[...]`).

---

## DataMigrationModel Field Reference

### `DataLocation`

| Field           | Type   | Required | Description                                            |
| --------------- | ------ | -------- | ------------------------------------------------------ |
| `storageType`   | string | ✅       | `filesystem` or `objectstorage`                        |
| `path`          | string | ✅       | File path (filesystem) or OS ID (objectstorage)        |
| `filesystem`    | object | —        | Required when `storageType=filesystem`                 |
| `objectStorage` | object | —        | Required when `storageType=objectstorage`              |
| `filter`        | object | —        | File filter (include/exclude glob patterns)            |
| `preCmd`        | string | —        | Shell command to run before transfer (source side)     |
| `postCmd`       | string | —        | Shell command to run after transfer (destination side) |

### `SSHConfig`

| Field            | Type   | Default | Description                                         |
| ---------------- | ------ | ------- | --------------------------------------------------- |
| `host`           | string | —       | SSH host address                                    |
| `port`           | int    | `22`    | SSH port                                            |
| `username`       | string | —       | SSH username                                        |
| `privateKey`     | string | —       | PEM private key content (newlines as `\n`)          |
| `privateKeyPath` | string | —       | Path to private key file on CM-Beetle host          |
| `useAgent`       | bool   | `false` | Use SSH agent for authentication                    |
| `connectTimeout` | int    | `30`    | Connection timeout in seconds                       |
| `archive`        | bool   | `false` | rsync `-a` flag (preserves permissions, timestamps) |
| `compress`       | bool   | `false` | rsync `-z` flag                                     |
| `delete`         | bool   | `false` | rsync `--delete` flag (remove files not in source)  |
| `dryRun`         | bool   | `false` | rsync `--dry-run` flag                              |

### `TumblebugConfig`

| Field      | Type   | Default | Description                                          |
| ---------- | ------ | ------- | ---------------------------------------------------- |
| `endpoint` | string | —       | Tumblebug API base URL (include `/tumblebug` prefix) |
| `nsId`     | string | —       | Namespace ID                                         |
| `osId`     | string | —       | Object Storage ID                                    |
| `expires`  | int    | `3600`  | Presigned URL expiration in seconds                  |
| `auth`     | object | —       | Authentication configuration                         |

### `AuthConfig`

| Field            | Type   | Description             |
| ---------------- | ------ | ----------------------- |
| `authType`       | string | `basic` or `jwt`        |
| `basic.username` | string | Username for basic auth |
| `basic.password` | string | Password for basic auth |
| `jwt.token`      | string | JWT token               |

---

## Error Reference

| HTTP Status | Condition                                                                    |
| ----------- | ---------------------------------------------------------------------------- |
| `400`       | Invalid request format, local filesystem access attempted, decryption failed |
| `400`       | Encryption key not found (already used or never generated)                   |
| `400`       | Encryption key expired                                                       |
| `202`       | Migration started successfully                                               |
| `500`       | Encryption key generation failed                                             |

Migration errors after `202` are reported via `GET /beetle/request/{reqId}`:

```json
{
  "data": {
    "status": "Error",
    "errorResponse": "Data migration failed: step 1 (upload-to-s3) failed: ..."
  }
}
```

---

## Integration with Object Storage Migration

A typical end-to-end workflow combining Object Storage infrastructure migration with data migration:

```
1. POST /recommendation/middleware/objectStorage   → get recommended target bucket config
2. POST /migration/middleware/ns/{nsId}/objectStorage → create source and target buckets
3. GET  /beetle/migration/data/encryptionKey       → get public key
4. POST /beetle/migration/data                     → migrate data source OS → target OS
5. GET  /beetle/request/{reqId}                    → poll until complete
6. DELETE /migration/middleware/ns/{nsId}/objectStorage/{osId} → cleanup
```

See the [Object Storage Feature Guide](object-storage-migration-feature-guide.md) for bucket recommendation and management details.

---

## Test Results

End-to-end test results across all supported CSPs are available in:
[docs/test-results-data-migration.md](../test-results-data-migration.md)
