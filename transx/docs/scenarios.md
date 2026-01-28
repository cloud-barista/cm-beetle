# transx Usage Scenarios

This document describes various data migration scenarios supported by the transx library. Each scenario explains the data flow from source to destination, optionally through relay nodes.

## Introduction

**transx** is a Go-based data migration library that supports multiple transfer methods for moving data between systems. The library provides:

- **Multi-Protocol Transfer**: Supports `rsync` (SSH-based) and `object-storage-api` (minio, CB-Spider, CB-Tumblebug) methods
- **Flexible Migration Modes**: Direct transfers and relay node scenarios
- **Optional Pre/Post Processing**: Backup and restore commands as hooks
- **Data Integrity**: Built-in validation and verification

## Terminology

| Term                | Description                                                                                              |
| ------------------- | -------------------------------------------------------------------------------------------------------- |
| **transx Host**     | The machine where transx runs. It orchestrates all transfers and may serve as a staging area.            |
| **Server**          | A remote machine accessed via SSH (e.g., Server A, Server B).                                            |
| **Object Storage**  | Cloud storage accessed via minio SDK (S3-compatible), CB-Spider API, or CB-Tumblebug API.                |
| **Staging**         | Temporary storage on the transx Host (`/tmp/transx-staging`), used when direct transfer is not possible. |
| **Direct Transfer** | Data flows directly between source and destination (transx Host only sends commands).                    |
| **Staged Transfer** | Data passes through transx Host's staging area before reaching the destination.                          |

## Table of Contents

- [1. Filesystem Transfers](#1-filesystem-transfers)
  - [1.1 Local → Remote (Push)](#11-local--remote-push)
  - [1.2 Remote → Local (Pull)](#12-remote--local-pull)
  - [1.3 Remote → Remote (Agent Forward)](#13-remote--remote-agent-forward)
  - [1.4 Remote → Remote (Relay)](#14-remote--remote-relay)
- [2. Object Storage Transfers](#2-object-storage-transfers)
  - [2.1 Local → Object Storage (Upload)](#21-local--object-storage-upload)
  - [2.2 Object Storage → Local (Download)](#22-object-storage--local-download)
  - [2.3 Object Storage → Object Storage (Cross-Cloud)](#23-object-storage--object-storage-cross-cloud)
- [3. Cross-Storage Transfers](#3-cross-storage-transfers)
  - [3.1 Remote Filesystem → Object Storage](#31-remote-filesystem--object-storage)
  - [3.2 Object Storage → Remote Filesystem](#32-object-storage--remote-filesystem)
- [4. Database Migration with Hooks](#4-database-migration-with-hooks)
  - [4.1 MariaDB Direct Mode Migration](#41-mariadb-direct-mode-migration)
  - [4.2 MariaDB Relay Mode Migration](#42-mariadb-relay-mode-migration)

---

## 1. Filesystem Transfers

Filesystem transfers use `rsync` over SSH for efficient file synchronization between systems.

### 1.1 Local → Remote (Push)

Transfer files from the local machine to a remote server.

**Use Case**: Deploy local files to a production server.

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/data/app",
    "filesystem": { "accessType": "local" }
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/var/www/app",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "prod-server.example.com",
        "port": 22,
        "username": "deploy",
        "privateKey": "-----BEGIN RSA PRIVATE KEY-----\n...\n-----END RSA PRIVATE KEY-----"
      }
    }
  }
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────┐
│                     Push Transfer                             │
├───────────────────────────────────────────────────────────────┤
│                                                               │
│   transx HOST                       SERVER                    │
│   ───────────                       ──────                    │
│   [Source]                          [Destination]             │
│                                                               │
│   /data/app ─────── rsync ────────► /var/www/app              │
│                      (SSH)                                    │
│                                                               │
│   Pipeline: filesystem-transfer                               │
│   Step 1: rsync-transfer (push)                               │
│                                                               │
└───────────────────────────────────────────────────────────────┘
```

### 1.2 Remote → Local (Pull)

Transfer files from a remote server to the local machine.

**Use Case**: Backup remote server data to local storage.

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/var/log/app",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "prod-server.example.com",
        "port": 22,
        "username": "backup",
        "privateKeyPath": "~/.ssh/id_rsa"
      }
    }
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/backup/logs",
    "filesystem": { "accessType": "local" }
  }
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────┐
│                     Pull Transfer                             │
├───────────────────────────────────────────────────────────────┤
│                                                               │
│   SERVER                            transx HOST               │
│   ──────                            ───────────               │
│   [Source]                          [Destination]             │
│                                                               │
│   /var/log/app ───── rsync ───────► /backup/logs              │
│                       (SSH)                                   │
│                                                               │
│   Pipeline: filesystem-transfer                               │
│   Step 1: rsync-transfer (pull)                               │
│                                                               │
└───────────────────────────────────────────────────────────────┘
```

### 1.3 Remote → Remote (Agent Forward)

Transfer files between two remote servers using SSH agent forwarding. The local machine orchestrates the transfer without staging data locally.

**Use Case**: Migrate data between servers without using local disk space.

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/data/source",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "server-a.example.com",
        "username": "admin",
        "useAgent": true
      }
    }
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/data/destination",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "server-b.example.com",
        "username": "admin",
        "useAgent": true
      }
    }
  },
  "strategy": "direct"
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────────────┐
│                 Agent Forward Transfer (Direct)                       │
├───────────────────────────────────────────────────────────────────────┤
│                                                                       │
│   transx HOST                                                         │
│   ───────────                                                         │
│   [Orchestrator]                                                      │
│    ┌──────────────┐                                                   │
│    │ SSH Agent    │  ← Keys stored here                               │
│    └──────┬───────┘                                                   │
│           │                                                           │
│           │ SSH + Agent Forwarding                                    │
│           ▼                                                           │
│      SERVER A ═══════════════════════════════════════► SERVER B       │
│      ────────           rsync (data transfer)          ────────       │
│      [Source]       (auth via forwarded agent)         [Destination]  │
│      /data/source ═══════════════════════════════════► /data/dest     │
│                                                                       │
│   • transx Host connects to Server A with agent forwarding            │
│   • transx Host runs rsync command on Server A                        │
│   • Server A authenticates to Server B using forwarded agent          │
│   • Data transfers directly: Server A → Server B                      │
│                                                                       │
│   Pipeline: filesystem-transfer                                       │
│   Step 1: rsync-transfer (agent-forward)                              │
│                                                                       │
└───────────────────────────────────────────────────────────────────────┘
```

### 1.4 Remote → Remote (Relay)

Transfer files between two remote servers via local staging. The local machine downloads from source, then uploads to destination.

**Use Case**: Migrate data between servers when agent forwarding is not available.

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/data/source",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "server-a.example.com",
        "username": "admin",
        "privateKey": "..."
      }
    }
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/data/destination",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "server-b.example.com",
        "username": "admin",
        "privateKey": "..."
      }
    }
  },
  "strategy": "relay"
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────────────┐
│                      Relay Transfer (2-Step)                          │
├───────────────────────────────────────────────────────────────────────┤
│                                                                       │
│   SERVER A               transx HOST                     SERVER B     │
│   ────────               ───────────                     ────────     │
│   [Source]               [Staging]                       [Destination]│
│                                                                       │
│   /data/source ── rsync ──► /tmp/transx-staging ── rsync ──► /data    │
│                   (pull)                            (push)            │
│                                                                       │
│   Pipeline: filesystem-transfer                                       │
│   Step 1: pull-to-staging (Server A → transx Host)                    │
│   Step 2: push-from-staging (transx Host → Server B)                  │
│                                                                       │
└───────────────────────────────────────────────────────────────────────┘
```

---

## 2. Object Storage Transfers

Object Storage transfers use minio-go (S3-compatible), CB-Spider API, or CB-Tumblebug API.

### 2.1 Local → Object Storage (Upload)

Upload local files to object storage using minio-go (S3-compatible SDK).

**Use Case**: Backup local data to cloud storage.

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/data/backup",
    "filesystem": { "accessType": "local" }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "my-bucket/backups/2024-01",
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
  }
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────┐
│                  Upload to Object Storage                     │
├───────────────────────────────────────────────────────────────┤
│                                                               │
│   transx HOST                       OBJECT STORAGE            │
│   ───────────                       ──────────────            │
│   [Source]                          [Destination]             │
│                                                               │
│   /data/backup ──── minio SDK ───► my-bucket/backups/        │
│                        (HTTPS)                                │
│                                                               │
│   Pipeline: cross-storage-transfer                            │
│   Step 1: upload-to-s3                                        │
│                                                               │
└───────────────────────────────────────────────────────────────┘
```

### 2.2 Object Storage → Local (Download)

Download files from object storage to local filesystem using CB-Spider API.

**Use Case**: Restore data from cloud backup.

```json
{
  "source": {
    "storageType": "objectstorage",
    "path": "my-bucket/backups/2024-01",
    "objectStorage": {
      "accessType": "spider",
      "spider": {
        "endpoint": "http://localhost:1024/spider",
        "connectionName": "aws-connection",
        "expires": 3600,
        "auth": {
          "authType": "basic",
          "basic": { "username": "admin", "password": "secret" }
        }
      }
    }
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/data/restored",
    "filesystem": { "accessType": "local" }
  }
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────┐
│              Download from Object Storage                     │
├───────────────────────────────────────────────────────────────┤
│                                                               │
│   OBJECT STORAGE                    transx HOST               │
│   ──────────────                    ───────────               │
│   [Source]                          [Destination]             │
│                                                               │
│   my-bucket/backups/ ─ Spider API ─► /data/restored            │
│                          (HTTPS)                              │
│                                                               │
│   Pipeline: cross-storage-transfer                            │
│   Step 1: download-from-s3                                    │
│                                                               │
└───────────────────────────────────────────────────────────────┘
```

### 2.3 Object Storage → Object Storage (Cross-Cloud)

Transfer data between two different object storage systems (e.g., AWS S3 → GCP Cloud Storage).

**Use Case**: Migrate data between cloud providers.

```json
{
  "source": {
    "storageType": "objectstorage",
    "path": "aws-bucket/data",
    "objectStorage": {
      "accessType": "minio",
      "minio": {
        "endpoint": "s3.amazonaws.com",
        "accessKeyId": "AWS_ACCESS_KEY",
        "secretAccessKey": "AWS_SECRET_KEY",
        "region": "us-east-1"
      }
    }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "gcp-bucket/data",
    "objectStorage": {
      "accessType": "minio",
      "minio": {
        "endpoint": "storage.googleapis.com",
        "accessKeyId": "GCP_ACCESS_KEY",
        "secretAccessKey": "GCP_SECRET_KEY",
        "region": "us-central1"
      }
    }
  }
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────────────┐
│                 Object Storage to Object Storage                      │
├───────────────────────────────────────────────────────────────────────┤
│                                                                       │
│   AWS S3                 transx HOST                     GCP Storage  │
│   ──────                 ───────────                     ───────────  │
│   [Source]               [Staging]                       [Destination]│
│                                                                       │
│   aws-bucket ─ minio SDK ─► /tmp/transx-staging ─ minio SDK ─► gcp    │
│               (download)                             (upload)         │
│                                                                       │
│   Pipeline: objectstorage-transfer                                    │
│   Step 1: download-from-s3 (AWS → transx Host)                        │
│   Step 2: upload-to-s3 (transx Host → GCP)                            │
│                                                                       │
└───────────────────────────────────────────────────────────────────────┘
```

---

## 3. Cross-Storage Transfers

Cross-storage transfers move data between filesystem and object storage, automatically handling the protocol differences.

### 3.1 Remote Filesystem → Object Storage

Transfer files from a remote server to object storage.

**Use Case**: Archive server data to cloud storage for long-term retention.

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/var/log/app",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "app-server.example.com",
        "username": "backup",
        "privateKey": "..."
      }
    }
  },
  "destination": {
    "storageType": "objectstorage",
    "path": "archive-bucket/logs/app-server",
    "objectStorage": {
      "accessType": "tumblebug",
      "tumblebug": {
        "endpoint": "http://localhost:1323/tumblebug",
        "nsId": "ns-01",
        "osId": "os-aws-01",
        "auth": {
          "authType": "basic",
          "basic": { "username": "default", "password": "default" }
        }
      }
    }
  }
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────────────┐
│               Remote Filesystem → Object Storage                      │
├───────────────────────────────────────────────────────────────────────┤
│                                                                       │
│   SERVER                 transx HOST                 OBJECT STORAGE   │
│   ──────                 ───────────                 ──────────────   │
│   [Source]               [Staging]                   [Destination]    │
│                                                                       │
│   /var/log/app ── rsync ──► /tmp/transx-staging ─ Tumblebug API ► bucket│
│                  (pull)                              (upload)         │
│                                                                       │
│   Pipeline: cross-storage-transfer                                    │
│   Step 1: rsync-from-server (Server → transx Host)                    │
│   Step 2: upload-to-s3 (transx Host → Object Storage)                 │
│                                                                       │
└───────────────────────────────────────────────────────────────────────┘
```

### 3.2 Object Storage → Remote Filesystem

Transfer files from object storage to a remote server.

**Use Case**: Restore archived data to a newly provisioned server.

```json
{
  "source": {
    "storageType": "objectstorage",
    "path": "archive-bucket/data/snapshot-2024",
    "objectStorage": {
      "accessType": "minio",
      "minio": {
        "endpoint": "s3.amazonaws.com",
        "accessKeyId": "AKIAIOSFODNN7EXAMPLE",
        "secretAccessKey": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
        "region": "ap-northeast-2"
      }
    }
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/data/restored",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "new-server.example.com",
        "username": "admin",
        "privateKey": "..."
      }
    }
  }
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────────────┐
│               Object Storage → Remote Filesystem                      │
├───────────────────────────────────────────────────────────────────────┤
│                                                                       │
│   OBJECT STORAGE         transx HOST                 SERVER           │
│   ──────────────         ───────────                 ──────           │
│   [Source]               [Staging]                   [Destination]    │
│                                                                       │
│   bucket/data ─ minio SDK ─► /tmp/transx-staging ── rsync ──► /data    │
│                (download)                            (push)           │
│                                                                       │
│   Pipeline: cross-storage-transfer                                    │
│   Step 1: download-from-s3 (Object Storage → transx Host)             │
│   Step 2: rsync-to-server (transx Host → Server)                      │
│                                                                       │
└───────────────────────────────────────────────────────────────────────┘
```

---

## 4. Database Migration with Hooks

Database migrations use `preCmd` and `postCmd` hooks to execute backup and restore commands around the transfer.

### 4.1 MariaDB Direct Mode Migration

Migrate a MariaDB database from one server to another using direct transfer.

**Use Case**: Database migration between servers in the same network.

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/tmp/db-backup",
    "preCmd": "mysqldump -u root -p'password' --all-databases > /tmp/db-backup/dump.sql",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "db-source.example.com",
        "username": "dba",
        "privateKey": "..."
      }
    }
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/tmp/db-restore",
    "postCmd": "mysql -u root -p'password' < /tmp/db-restore/dump.sql",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "db-dest.example.com",
        "username": "dba",
        "privateKey": "..."
      }
    }
  },
  "strategy": "direct"
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────────────────────┐
│                  MariaDB Migration (Direct Mode)                              │
├───────────────────────────────────────────────────────────────────────────────┤
│                                                                               │
│   SERVER A (DB Source)                                SERVER B (DB Dest)      │
│   ────────────────────                                ──────────────────      │
│   [Source]                                            [Destination]           │
│                                                                               │
│   ① PreCmd: mysqldump                                                         │
│      MariaDB ──► /tmp/db-backup/dump.sql                                      │
│                                                                               │
│   ② Transfer: rsync (direct)                                                  │
│      /tmp/db-backup/dump.sql ─────────────────────► /tmp/db-restore/dump.sql  │
│                                                                               │
│                                                      ③ PostCmd: mysql         │
│                                                         /tmp/db-restore ──► DB│
│                                                                               │
│   MigrateData() Stages:                                                       │
│   ├─ Stage 1: Backup (preCmd on Server A)                                     │
│   ├─ Stage 2: Transfer (rsync)                                                │
│   └─ Stage 3: Restore (postCmd on Server B)                                   │
│                                                                               │
└───────────────────────────────────────────────────────────────────────────────┘
```

### 4.2 MariaDB Relay Mode Migration

Migrate a MariaDB database through a relay node when direct transfer is not possible.

**Use Case**: Database migration between networks that cannot communicate directly.

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/tmp/db-backup",
    "preCmd": "mysqldump -u root -p'password' mydb > /tmp/db-backup/mydb.sql",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "10.0.1.50",
        "username": "dba",
        "privateKey": "..."
      }
    }
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/tmp/db-restore",
    "postCmd": "mysql -u root -p'password' mydb < /tmp/db-restore/mydb.sql",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "10.0.2.50",
        "username": "dba",
        "privateKey": "..."
      }
    }
  },
  "strategy": "relay"
}
```

**Data Flow**:

```
┌───────────────────────────────────────────────────────────────────────────────┐
│                   MariaDB Migration (Relay Mode)                              │
├───────────────────────────────────────────────────────────────────────────────┤
│                                                                               │
│   SERVER A (10.0.1.50)   transx HOST           SERVER B (10.0.2.50)           │
│   ────────────────────   ───────────           ─────────────────────          │
│   [Source]               [Staging]             [Destination]                  │
│                                                                               │
│   ① PreCmd: mysqldump                                                         │
│      MariaDB ──► /tmp/db-backup/mydb.sql                                      │
│                                                                               │
│   ② Transfer Step 1: rsync (pull)                                             │
│      /tmp/db-backup ──────────────────► /tmp/transx-staging/                  │
│                                                                               │
│   ③ Transfer Step 2: rsync (push)                                             │
│                        /tmp/transx-staging/ ──────────► /tmp/db-restore       │
│                                                                               │
│                                                ④ PostCmd: mysql               │
│                                                   /tmp/db-restore ──► MariaDB │
│                                                                               │
│   MigrateData() Stages:                                                       │
│   ├─ Stage 1: Backup (preCmd on Server A)                                     │
│   ├─ Stage 2: Transfer                                                        │
│   │   ├─ Step 1: pull-to-staging                                              │
│   │   └─ Step 2: push-from-staging                                            │
│   └─ Stage 3: Restore (postCmd on Server B)                                   │
│                                                                               │
└───────────────────────────────────────────────────────────────────────────────┘
```

---

## Summary

| Scenario                        | Pipeline               | Steps                 | Key Components                         |
| ------------------------------- | ---------------------- | --------------------- | -------------------------------------- |
| Local → Remote                  | filesystem-transfer    | 1 (push)              | rsync over SSH                         |
| Remote → Local                  | filesystem-transfer    | 1 (pull)              | rsync over SSH                         |
| Remote → Remote (Direct)        | filesystem-transfer    | 1 (agent-forward)     | rsync with SSH agent                   |
| Remote → Remote (Relay)         | filesystem-transfer    | 2 (pull + push)       | rsync via local staging                |
| Local → Object Storage          | cross-storage-transfer | 1 (upload)            | minio SDK / Spider API / Tumblebug API |
| Object Storage → Local          | cross-storage-transfer | 1 (download)          | minio SDK / Spider API / Tumblebug API |
| Object Storage → Object Storage | objectstorage-transfer | 2 (download + upload) | Object Storage API via local staging   |
| Remote FS → Object Storage      | cross-storage-transfer | 2 (rsync + upload)    | rsync + Object Storage API             |
| Object Storage → Remote FS      | cross-storage-transfer | 2 (download + rsync)  | Object Storage API + rsync             |
| Database Migration              | Any of above           | + preCmd/postCmd      | Hooks for backup/restore               |

## Related Resources

- [transx README](../README.md) - Main documentation
- [MariaDB Migration Example](../examples/mariadb-migration/) - Complete database migration example
- [Object Storage Example](../examples/object-storage/) - Object storage transfer examples
