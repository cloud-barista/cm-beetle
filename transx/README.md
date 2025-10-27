# transx

A Go-based data migration library that supports multiple transfer methods for moving data between databases and storage systems.

## Features

- **Error-Only Approach**: Structured error handling for seamless integration with logging frameworks (zerolog, logrus, etc.)
- **Multi-Protocol Transfer**: Core data transfer functionality with support for `rsync` and `object-storage-api` methods
- **Optional Data Processing**: Backup and restore operations as supplementary features when needed
- **Explicit Configuration**: User-specified transfer methods (no auto-detection)
- **Direct & Relay Modes**: Flexible migration patterns supporting both direct transfers and relay node scenarios
- **Data Integrity**: Built-in validation and verification with comprehensive error context

## Quick Start

```go
import "github.com/yunkon-kim/transx"

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
go get github.com/yunkon-kim/transx
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
