# MariaDB Data Migration Example with transx

This example demonstrates how to perform MariaDB database migration using the transx library with Docker containers simulating remote servers.

## Overview

This example performs the following tasks:

1. Database backup from source MariaDB (mariadb-dump via SSH)
2. Transfer backup files to the target server (rsync via SSH)
3. Database restoration on the target MariaDB (via SSH)

### Supported Strategies & Transfer Modes

The `transx` library supports 5 strategies (`auto`, `direct`, `relay`, `ssh-agent-fwd-push`, `ssh-agent-fwd-pull`) and generalized `TransferMode` execution directions:

| TransferMode         | Source       | Destination  | Description                                                                   |
| -------------------- | ------------ | ------------ | ----------------------------------------------------------------------------- |
| `pull`               | SSH / S3     | Local        | Pulls data from remote endpoint to local filesystem                           |
| `push`               | Local        | SSH / S3 / Local | Pushes data from local filesystem to destination                          |
| `ssh-agent-fwd-auto` | SSH (remote) | SSH (remote) | Probes `ssh-agent-fwd-push` first, then falls back to `ssh-agent-fwd-pull`    |
| `ssh-agent-fwd-push` | SSH (remote) | SSH (remote) | Connects to Source with SSH agent forwarding and pushes directly to Destination |
| `ssh-agent-fwd-pull` | SSH (remote) | SSH (remote) | Connects to Destination with SSH agent forwarding and pulls directly from Source |

### Example Configurations

This example provides 3 configuration files:

| Config                       | Strategy | Flow                                     | Use Case                      |
| ---------------------------- | -------- | ---------------------------------------- | ----------------------------- |
| `direct-mode-config.json`    | `direct` | SSH source -> Local (`pull`)             | Backup only (no restore)      |
| `relay-mode-config.json`     | `relay`  | SSH source -> Local -> SSH dest (`pull` + `push`) | Full migration via local |
| `agent-fwd-mode-config.json` | `direct` | SSH source -> SSH dest (`ssh-agent-fwd-auto`) | Direct container-to-container |

## Prerequisites

- Go 1.21 or higher
- Docker installed and running
- SSH client installed

## Quick Start

### 1. Setup Environment

```bash
# Make scripts executable
chmod +x setup_environment.sh migrate.sh

# Setup both MariaDB containers with SSH access and test data
./setup_environment.sh all
```

This will:
- Build a custom MariaDB image with SSH server and rsync
- Start source container (MariaDB: 3306, SSH: 2222)
- Start target container (MariaDB: 3307, SSH: 2223)
- Generate SSH keypair in `./ssh_keys/`
- Configure SSH access between containers
- Populate source database with test data

### 2. Run Migration

```bash
# Direct mode: Pull backup from source container to local
./migrate.sh direct --verbose

# Relay mode: Full migration via local relay (source → local → target)
./migrate.sh relay --verbose

# Agent-Forward mode: Direct transfer between containers
./migrate.sh agent-fwd --verbose
```

## Configuration Files

### Direct Mode (`direct-mode-config.json`)

**Pull mode:** SSH source → Local filesystem (backup only)

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/backup",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "localhost",
        "port": 2222,
        "username": "root",
        "privateKeyPath": "./ssh_keys/id_rsa"
      }
    },
    "preCmd": "mariadb-dump ... > /backup/testdb_backup.sql"
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/home/ubuntu/mariadb_backup",
    "filesystem": {
      "accessType": "local"
    }
  },
  "strategy": "direct"
}
```

### Relay Mode (`relay-mode-config.json`)

**Relay mode:** SSH source → Local → SSH destination (full migration via local)

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/backup",
    "filesystem": {
      "accessType": "ssh",
      "ssh": { "host": "localhost", "port": 2222, ... }
    },
    "preCmd": "mariadb-dump ..."
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/backup",
    "filesystem": {
      "accessType": "ssh",
      "ssh": { "host": "localhost", "port": 2223, ... }
    },
    "postCmd": "mariadb ... < /backup/testdb_backup.sql"
  },
  "strategy": "relay"
}
```

### Agent-Forward Mode (`agent-fwd-mode-config.json`)

**AgentForward mode:** SSH source → SSH destination (direct transfer)

```json
{
  "source": {
    "storageType": "filesystem",
    "path": "/backup",
    "filesystem": {
      "accessType": "ssh",
      "ssh": { "host": "localhost", "port": 2222, ... }
    },
    "preCmd": "mariadb-dump ..."
  },
  "destination": {
    "storageType": "filesystem",
    "path": "/backup",
    "filesystem": {
      "accessType": "ssh",
      "ssh": {
        "host": "HOST_IP_PLACEHOLDER",
        "port": 2223,
        "privateKeyPath": "/root/.ssh/id_rsa"
      }
    },
    "postCmd": "mariadb ..."
  },
  "strategy": "direct"
}
```

> **Note:** 
> - `HOST_IP_PLACEHOLDER` is replaced with host IP by `migrate.sh`
> - `source.privateKeyPath`: Host path (for host → source connection)
> - `destination.privateKeyPath`: Source container's internal path (for source → target rsync)

## Architecture

```
┌─────────────────────────────────────────────────────────────────────────┐
│                              Host Machine                                │
│                                                                          │
│  ┌─────────────────────┐                    ┌─────────────────────┐     │
│  │  mariadb_source     │                    │  mariadb_target     │     │
│  │  (SSH: 2222)        │                    │  (SSH: 2223)        │     │
│  │  (MariaDB: 3306)    │                    │  (MariaDB: 3307)    │     │
│  │                     │                    │                     │     │
│  │  /backup/           │                    │  /backup/           │     │
│  │    testdb_backup.sql│                    │    testdb_backup.sql│     │
│  └─────────────────────┘                    └─────────────────────┘     │
│           │                                           ▲                  │
│           │                                           │                  │
│           ▼                                           │                  │
│  ┌─────────────────────────────────────────────────────┐                │
│  │                    Transfer Modes                    │                │
│  │                                                      │                │
│  │  [Direct]     source ──────────────▶ local           │                │
│  │  [Relay]      source ──▶ local ──▶ destination       │                │
│  │  [AgentFwd]   source ──────────────▶ destination     │                │
│  └─────────────────────────────────────────────────────┘                │
│                                                                          │
│  ./ssh_keys/id_rsa  ─── SSH Key for container access                    │
│  /home/ubuntu/mariadb_backup  ─── Local staging area (for relay)        │
└─────────────────────────────────────────────────────────────────────────┘
```

## Container Setup Details

After running `./setup_environment.sh all`:

| Container       | MariaDB Port | SSH Port | Root Password    | SSH User |
| --------------- | ------------ | -------- | ---------------- | -------- |
| mariadb_source  | 3306         | 2222     | source_password  | root     |
| mariadb_target  | 3307         | 2223     | target_password  | root     |

**SSH Key Location:** `./ssh_keys/id_rsa`

The source database contains:
- **users** table: 5 sample users
- **orders** table: 5 sample orders

## Troubleshooting

### SSH connection failed

```bash
# Test SSH to source container
ssh -i ./ssh_keys/id_rsa -p 2222 root@localhost

# Test SSH to target container
ssh -i ./ssh_keys/id_rsa -p 2223 root@localhost

# Re-run SSH connectivity test
./setup_environment.sh test-ssh
```

### Docker not running

```bash
sudo systemctl start docker
```

### Container not starting

```bash
# Check container logs
docker logs mariadb_source
docker logs mariadb_target

# Restart environment
./setup_environment.sh cleanup
./setup_environment.sh all
```

### Permission denied on SSH key

```bash
chmod 600 ./ssh_keys/id_rsa
```

## Cleanup

```bash
# Stop containers and remove SSH keys
./setup_environment.sh cleanup
```

This will:
- Stop and remove both MariaDB containers
- Clean up backup directories
- Remove generated SSH keys from `./ssh_keys/`

## Building the Example

```bash
cd /path/to/transx/examples/mariadb-migration
go build -o mariadb-migration main.go
```

## Running Manually

```bash
# Direct mode
./mariadb-migration --config=direct-mode-config.json --verbose

# Individual steps
./mariadb-migration --config=direct-mode-config.json --backup
./mariadb-migration --config=direct-mode-config.json --transfer
./mariadb-migration --config=direct-mode-config.json --restore
```
