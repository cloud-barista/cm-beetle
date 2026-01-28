# transx

A Go-based data migration library that supports multiple transfer methods for moving data between databases and storage systems.

## Features

- **Error-Only Approach**: Structured error handling for seamless integration with logging frameworks (zerolog, logrus, etc.)
- **Multi-Protocol Transfer**: Core data transfer functionality with support for `rsync` and `object-storage-api` methods
- **Optional Data Processing**: Backup and restore operations as supplementary features when needed
- **Explicit Configuration**: User-specified transfer methods (no auto-detection)
- **Direct & Relay Modes**: Flexible migration patterns supporting both direct transfers and relay node scenarios
- **Data Integrity**: Built-in validation and verification with comprehensive error context
- **Sensitive Data Encryption**: Hybrid encryption (AES-256-GCM + RSA-OAEP) for secure transmission of credentials

## Quick Start

```go
import "github.com/cloud-barista/cm-beetle/transx"

// Core transfer operation
err := transx.Transfer(dataModel)
if err != nil {
    log.Error().Err(err).Msg("Transfer failed")
}

// Complete migration with optional backup/restore
err = transx.MigrateData(dataModel)
if err != nil {
    // Rich error context available for logging frameworks
    log.Error().Err(err).Msg("Migration failed")
}
```

## Transfer Methods

| Method               | Description                                   | Use Case                                        |
| -------------------- | --------------------------------------------- | ----------------------------------------------- |
| `rsync`              | SSH-based file transfers using rsync          | Remote server migrations with authentication    |
| `object-storage-api` | HTTP-based transfers with Object Storage APIs | S3-compatible storage (CB-Spider, AWS S3, etc.) |

**Note**: File operations on the same host are handled automatically when endpoint is empty.

## Architecture

The library implements a **Transfer-Centric Data Migration Model** with the following components:

**Core Functionality**:

- **Transfer**: Move data between systems using specified transfer methods (rsync, object-storage-api)

**Optional Operations** (when backup/restore commands are provided):

- **Backup**: Export data from source systems before transfer
- **Restore**: Import transferred data into destination systems

The transfer operation is the primary focus, with backup and restore serving as optional pre/post-processing steps.

### Transfer Modes

- **Direct Mode**: Source → Destination (at least one endpoint is local)
- **Relay Mode**: Source → Relay Node → Destination (both endpoints are remote)

## Error Handling

The library implements an **Error-Only Approach** with unified error handling:

```go
// Unified OperationError provides rich context for all operations
type OperationError struct {
    Operation   string            // "backup", "restore", "transfer"
    Method      string            // transfer method (for transfer operations)
    Source      string            // source path/endpoint
    Destination string            // destination path/endpoint
    Command     string            // executed command (for backup/restore)
    Output      string            // command output (for backup/restore)
    IsRelayMode bool              // relay mode flag (for transfer)
    Context     map[string]string // additional context information
    Err         error             // underlying error
}

// Simple interface for users - just use err.Error()
// Advanced users can access rich context via type assertion
```

The transfer operation always executes, while backup and restore are conditional based on configuration.

## Examples

### MariaDB Migration

See [examples/mariadb-migration](examples/mariadb-migration/) for a complete database migration example:

- Direct mode migration (local ↔ remote)
- Relay mode migration (remote ↔ remote)
- Individual step execution (backup, transfer, restore)
- Comprehensive testing and validation

### Object Storage Migration

See [examples/object-storage](examples/object-storage/) for Object Storage integration:

- CB-Spider presigned URL integration
- Bidirectional transfers (local ↔ object storage)
- Cross-cloud migration scenarios
- AWS S3 compatibility

## Installation

```bash
go get github.com/cloud-barista/cm-beetle/transx
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

---

_Appendix_

## Sensitive Data Encryption

> [!NOTE]
> transx was developed as an internal tool/package for easy data transfer between systems, so secure transmission of sensitive data was not a primary design goal. However, we provide a minimal encryption option for users who need additional security.

transx provides built-in encryption for sensitive fields like SSH private keys and S3 credentials. The encryption uses a **hybrid approach** combining AES-256-GCM for data encryption and RSA-OAEP for key exchange, supporting data of any size.

### Sensitive Fields

The following fields are automatically encrypted when using `EncryptModel`:

| Field Path                                      | Description                   |
| ----------------------------------------------- | ----------------------------- |
| `*.filesystem.ssh.privateKey`                   | SSH private key content       |
| `*.objectStorage.minio.accessKeyId`             | S3 access key ID              |
| `*.objectStorage.minio.secretAccessKey`         | S3 secret access key          |
| `*.objectStorage.spider.auth.basic.password`    | Spider Basic auth password    |
| `*.objectStorage.spider.auth.jwt.token`         | Spider JWT token              |
| `*.objectStorage.tumblebug.auth.basic.password` | Tumblebug Basic auth password |
| `*.objectStorage.tumblebug.auth.jwt.token`      | Tumblebug JWT token           |

### Encryption Workflow

```
┌─────────────────────────────────────────────────────────────────────────┐
│                     Encryption Workflow                                 │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│   CLIENT                                SERVER                          │
│   ──────                                ──────                          │
│                                                                         │
│   ① Request Key ─────────────────────► Generate KeyPair                │
│                                         Store in KeyStore               │
│              ◄───────────────────────── Return PublicKeyBundle          │
│                                                                         │
│   ② EncryptModel()                                                      │
│      - Parse PublicKey                                                  │
│      - Encrypt sensitive fields                                         │
│                                                                         │
│   ③ Send encrypted model ────────────► Receive                         │
│                                                                         │
│                                         ④ DecryptModelWithStore()       │
│                                            - Lookup KeyPair by keyId    │
│                                            - Decrypt fields             │
│                                            - Auto-delete key (one-time) │
│                                                                         │
│                                         ⑤ Execute migration             │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

### Usage: Plaintext Transmission (No Encryption)

For internal/trusted environments where encryption is not required:

**Server Side**

```go
package main

import (
    "encoding/json"
    "net/http"

    "github.com/cloud-barista/cm-beetle/transx"
)

// main initializes the HTTP server for plaintext data migration.
func main() {
    http.HandleFunc("/api/v1/migration", handleMigration)
    http.ListenAndServe(":8080", nil)
}

// handleMigration processes plaintext migration requests.
// It validates the request body and executes the migration directly.
func handleMigration(w http.ResponseWriter, r *http.Request) {
    var model transx.DataMigrationModel
    if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Validate the model
    if err := transx.Validate(model); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Execute migration directly with plaintext model
    if err := transx.MigrateData(model); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "completed"})
}
```

**Client Side**

```go
package main

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"

    "github.com/cloud-barista/cm-beetle/transx"
)

// main demonstrates sending a plaintext migration request to the server.
// Use this approach only in trusted/internal environments.
func main() {
    // Create migration model with sensitive data (plaintext)
    model := transx.DataMigrationModel{
        Source: transx.DataLocation{
            StorageType: transx.StorageTypeFilesystem,
            Path:        "/data/source",
            Filesystem: &transx.FilesystemAccess{
                AccessType: transx.AccessTypeSSH,
                SSH: &transx.SSHConfig{
                    Host:       "192.168.1.100",
                    Port:       22,
                    Username:   "ubuntu",
                    PrivateKey: "-----BEGIN RSA PRIVATE KEY-----\nMIIE...\n-----END RSA PRIVATE KEY-----",
                },
            },
        },
        Destination: transx.DataLocation{
            StorageType: transx.StorageTypeObjectStorage,
            Path:        "my-bucket/backup",
            ObjectStorage: &transx.ObjectStorageAccess{
                AccessType: transx.AccessTypeMinio,
                Minio: &transx.S3MinioConfig{
                    Endpoint:        "s3.amazonaws.com",
                    AccessKeyId:     "AKIAIOSFODNN7EXAMPLE",
                    SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
                    Region:          "us-east-1",
                },
            },
        },
    }

    // Send plaintext model directly
    jsonData, _ := json.Marshal(model)
    resp, err := http.Post("http://localhost:8080/api/v1/migration",
        "application/json", bytes.NewReader(jsonData))
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    log.Printf("Migration status: %d", resp.StatusCode)
}
```

### Usage: Encrypted Transmission (Recommended for Production)

For secure transmission of sensitive data over untrusted networks:

**Server Side**

```go
package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/cloud-barista/cm-beetle/transx"
)

// Global KeyStore for managing encryption keys
var keyStore = transx.NewKeyStore()

// main initializes the HTTP server with encryption key management.
// It starts a background routine to clean up expired keys.
func main() {
    // Start background cleanup for expired keys (every 10 minutes)
    stopCleanup := make(chan struct{})
    keyStore.StartCleanupRoutine(10*time.Minute, stopCleanup)
    defer close(stopCleanup)

    http.HandleFunc("/api/v1/encryption/key", handleGetKey)
    http.HandleFunc("/api/v1/migration/secure", handleSecureMigration)
    http.ListenAndServe(":8080", nil)
}

// handleGetKey generates a new RSA key pair and returns the public key bundle.
// The key is stored in KeyStore and valid for 30 minutes (one-time use).
func handleGetKey(w http.ResponseWriter, r *http.Request) {
    // Generate a one-time use key (30 minutes validity)
    keyPair, err := keyStore.GenerateKeyPair(30 * time.Minute)
    if err != nil {
        http.Error(w, "Key generation failed", http.StatusInternalServerError)
        return
    }

    // Export public key bundle for client
    bundle, err := keyPair.ExportPublicBundle()
    if err != nil {
        http.Error(w, "Key export failed", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(bundle)
}

// handleSecureMigration processes encrypted migration requests.
// It decrypts the model using the one-time key, then executes the migration.
// The encryption key is automatically deleted after successful decryption.
func handleSecureMigration(w http.ResponseWriter, r *http.Request) {
    var model transx.DataMigrationModel
    if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Check if model is encrypted
    if !model.IsEncrypted() {
        http.Error(w, "Encrypted model required", http.StatusBadRequest)
        return
    }

    // Decrypt the model (key is auto-deleted after decryption - one-time use)
    decryptedModel, err := transx.DecryptModelWithStore(model, keyStore)
    if err != nil {
        switch err {
        case transx.ErrKeyNotFound:
            http.Error(w, "Key not found or already used", http.StatusBadRequest)
        case transx.ErrKeyExpired:
            http.Error(w, "Key has expired", http.StatusBadRequest)
        default:
            http.Error(w, "Decryption failed", http.StatusInternalServerError)
        }
        return
    }

    // Now 'decryptedModel' contains plaintext sensitive data
    // Execute migration
    if err := transx.MigrateData(decryptedModel); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "completed"})
}
```

**Client Side**

```go
package main

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"

    "github.com/cloud-barista/cm-beetle/transx"
)

// main demonstrates the complete encrypted transmission workflow:
// 1. Request public key from server
// 2. Encrypt sensitive fields in the model
// 3. Send encrypted model to server
func main() {
    // Step 1: Request public key from server
    resp, err := http.Get("http://localhost:8080/api/v1/encryption/key")
    if err != nil {
        log.Fatal("Failed to get encryption key:", err)
    }
    defer resp.Body.Close()

    var bundle transx.PublicKeyBundle
    if err := json.NewDecoder(resp.Body).Decode(&bundle); err != nil {
        log.Fatal("Failed to parse key bundle:", err)
    }

    log.Printf("Received key: %s (expires: %v)", bundle.KeyID, bundle.ExpiresAt)

    // Step 2: Parse the public key
    publicKey, err := transx.ParsePublicKeyBundle(bundle)
    if err != nil {
        log.Fatal("Failed to parse public key:", err)
    }

    // Step 3: Create migration model with sensitive data
    model := transx.DataMigrationModel{
        Source: transx.DataLocation{
            StorageType: transx.StorageTypeFilesystem,
            Path:        "/data/source",
            Filesystem: &transx.FilesystemAccess{
                AccessType: transx.AccessTypeSSH,
                SSH: &transx.SSHConfig{
                    Host:       "192.168.1.100",
                    Port:       22,
                    Username:   "ubuntu",
                    // This 3KB+ private key will be encrypted using hybrid encryption
                    PrivateKey: "-----BEGIN RSA PRIVATE KEY-----\nMIIEpQIBAAKCAQEA...(long key content)...\n-----END RSA PRIVATE KEY-----",
                },
            },
        },
        Destination: transx.DataLocation{
            StorageType: transx.StorageTypeObjectStorage,
            Path:        "my-bucket/backup",
            ObjectStorage: &transx.ObjectStorageAccess{
                AccessType: transx.AccessTypeMinio,
                Minio: &transx.S3MinioConfig{
                    Endpoint:        "s3.amazonaws.com",
                    AccessKeyId:     "AKIAIOSFODNN7EXAMPLE",
                    SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
                    Region:          "us-east-1",
                },
            },
        },
    }

    // Step 4: Encrypt sensitive fields
    encryptedModel, err := transx.EncryptModel(model, publicKey, bundle.KeyID)
    if err != nil {
        log.Fatal("Failed to encrypt model:", err)
    }

    log.Printf("Model encrypted with key: %s", encryptedModel.EncryptionKeyID)

    // Step 5: Send encrypted model to server
    jsonData, _ := json.Marshal(encryptedModel)
    resp, err = http.Post("http://localhost:8080/api/v1/migration/secure",
        "application/json", bytes.NewReader(jsonData))
    if err != nil {
        log.Fatal("Failed to send request:", err)
    }
    defer resp.Body.Close()

    log.Printf("Migration status: %d", resp.StatusCode)
}
```

### Encrypted JSON Example

When transmitted, the encrypted model looks like this:

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/data/source",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "192.168.1.100",
        "port": 22,
        "username": "ubuntu",
        "privateKey": "eyJ2IjoxLCJlayI6IkFBRUJBd1FGQmdj..."
      }
    }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "my-bucket/backup",
    "objectStorage": {
      "accessType": "minio",
      "minio": {
        "endpoint": "s3.amazonaws.com",
        "accessKeyId": "eyJ2IjoxLCJlayI6IkhJSktMTU5P...",
        "secretAccessKey": "eyJ2IjoxLCJlayI6IlBRUlNUVVZX...",
        "region": "us-east-1"
      }
    }
  },
  "encryptionKeyId": "key-a1b2c3d4e5f6..."
}
```

**Note**: The `encryptionKeyId` field indicates which key was used for encryption.
Sensitive fields are automatically encrypted based on the predefined list in `transx-sec.go`.
No need to specify which fields are encrypted - the server knows from the internal `sensitiveFields` definition.

### Security Considerations

| Aspect                 | Implementation                                            |
| ---------------------- | --------------------------------------------------------- |
| **Algorithm**          | Hybrid: AES-256-GCM (data) + RSA-2048-OAEP (key exchange) |
| **Key Validity**       | Default 30 minutes (configurable)                         |
| **One-Time Use**       | Keys are automatically deleted after decryption           |
| **Large Data Support** | Hybrid encryption handles data of any size                |
| **Transport Security** | Always use HTTPS in production                            |
| **Key Storage**        | In-memory only (keys don't survive restart)               |
