# CB-Tumblebug Files Sync Guide

Files under this directory are copied from the [CB-Tumblebug](https://github.com/cloud-barista/cb-tumblebug) repository.
When upgrading CB-Tumblebug, check each file against the upstream source and sync as needed.

## Sync Procedure

1. Check out the target CB-Tumblebug version:

   ```bash
   cd /path/to/cb-tumblebug
   git checkout v<VERSION>
   ```

2. Discover all changed files systematically:

   ```bash
   TB=/path/to/cb-tumblebug
   BEETLE=deployments/docker-compose/cb-tumblebug

   # Detect changed files in each directory
   for dir in assets conf init scripts; do
     echo "=== Checking $dir/ ==="
     diff -qr $BEETLE/$dir $TB/$dir 2>&1 | grep "^Files.*differ$"
   done

   # Check MCP files (different path in TB)
   diff -qr $BEETLE/interface/mcp $TB/src/interface/mcp

   # Detect new files (TB only)
   for dir in assets conf init scripts; do
     diff -qr $BEETLE/$dir $TB/$dir 2>&1 | grep "^Only in $TB"
   done

   # Note: assets/spider is excluded from synchronization (cb-tumblebug issue #2694)
   # Note: assets/rdbmsinfo.yaml is required for Tumblebug managed RDBMS feature

   # Detect removed files (Beetle only)
   for dir in assets conf init scripts; do
     diff -qr $BEETLE/$dir $TB/$dir 2>&1 | grep "^Only in $BEETLE"
   done
   ```

3. Review individual file changes:

   ```bash
   # Example: review specific file diff
   diff -u $BEETLE/assets/cloudimage.csv $TB/assets/cloudimage.csv
   ```

4. Copy new or updated files:

   ```bash
   cp $TB/assets/new-file.yaml $BEETLE/assets/

   # MCP files (note different source path)
   cp $TB/src/interface/mcp/tb-mcp.py $BEETLE/interface/mcp/
   ```

5. For binary assets (`assets.dump.gz`), compare checksums:
   ```bash
   md5sum $BEETLE/assets/assets.dump.gz $TB/assets/assets.dump.gz
   ```

---

## v0.12.30 Sync (2026-08-07)

Based on TB v0.12.30 `c2c4e76b` (tagged release). Upgrade path: **v0.12.25 → v0.12.30**.

### Model Changes (v0.12.25→v0.12.30)

| Change                               | Version  | Description                                             |
| ------------------------------------ | -------- | ------------------------------------------------------- |
| **PostCommand → PostCommands**       | v0.12.29 | Multi-phase bootstrap refactor (BREAKING)               |
| **NodeUserPassword removal**         | v0.12.29 | Field removed from CreateNodeGroupDynamicReq (BREAKING) |
| **DistributeSubnets**                | v0.12.29 | Multi-AZ subnet distribution feature                    |
| **Resource Pruning structs**         | v0.12.29 | Added ResourcePruneResult, ResourcePruneResults         |
| **PostCommandReq**                   | v0.12.29 | New struct for phase-based command targeting            |
| **AddNodeGroupDynamicReq**           | v0.12.29 | New struct for adding nodes to existing infra           |
| **PostCommandStatus**                | v0.12.29 | New enum for command execution status                   |
| **PostCommandPhaseResult**           | v0.12.29 | New struct for phase execution results                  |
| **SshCmdResultForAPI**               | v0.12.29 | New API response type for SSH results                   |
| **InfraSshCmdResultForAPI**          | v0.12.29 | New wrapper for multiple SSH results                    |
| **StatusCountInfo.CountReconciling** | v0.12.29 | Added reconciling status count                          |
| **ConnConfig.VerifiedMessage**       | v0.12.30 | Credential verification error details                   |
| **NodeInfoInNs**                     | v0.12.30 | Namespace-wide Node listing                             |
| **NLBInfoInNs**                      | v0.12.30 | Namespace-wide NLB listing                              |

### Deployment File Changes

| File                                  | Action                                                                                                                             |
| ------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| **Model files (imdl/cloud-model/)**   |                                                                                                                                    |
| `copied-tb-model.go`                  | **Updated** — v0.12.25→v0.12.30 model synchronization (PostCommand→PostCommands, VerifiedMessage, NodeInfoInNs, NLBInfoInNs, etc.) |
| **go.mod**                            |                                                                                                                                    |
| `go.mod`                              | **Updated** — cb-tumblebug `v0.12.25→v0.12.30`                                                                                     |
| **docker-compose.yaml**               |                                                                                                                                    |
| `docker-compose.yaml`                 | **Updated** — cb-tumblebug `0.12.25→0.12.30`, cb-spider `0.12.35→0.12.42`, cb-mapui `0.12.50→0.12.56`                              |
|                                       | **Added** — cb-spider-postgres service for Spider METADB (replaces SQLite)                                                         |
| **.env**                              |                                                                                                                                    |
| `.env`                                | **Updated** — Added SP_POSTGRES_USER, SP_POSTGRES_PASSWORD, SP_POSTGRES_DB for Spider PostgreSQL                                   |
| **Assets**                            |                                                                                                                                    |
| `assets/*`                            | **No changes** — All asset files remain unchanged from v0.12.25                                                                    |
| **Config**                            |                                                                                                                                    |
| `conf/cloud_conf.yaml`                | **Changed in TB** — Updated, but not synced (beetle manages its own cloud configuration)                                           |
| `conf/setup.env`, `conf/traefik.yaml` | **TB-specific configs** — Not copied (not needed for beetle's deployment)                                                          |
| **Init**                              |                                                                                                                                    |
| `init/cleanDB.sh`                     | **Changed in TB** — Not synced (beetle manages its own init flow)                                                                  |
| `init/init.py`                        | **Changed in TB** — Not synced (beetle manages its own init flow)                                                                  |
| `init/init.sh`                        | **Changed in TB** — Not synced (beetle manages its own init flow)                                                                  |
| `init/multi-init.sh`                  | **Changed in TB** — Not synced (beetle manages its own init flow)                                                                  |
| `init/openbao/`                       | **TB-specific** — Not copied (not needed for beetle deployment)                                                                    |
| `init/templates/*.json`               | **Reorganized in TB** — Many new templates added, beetle maintains its own templates                                               |
| **Scripts**                           |                                                                                                                                    |
| `scripts/restore-assets.sh`           | **Changed in TB** — Not synced (beetle manages its own scripts)                                                                    |
| `scripts/*`                           | **Many new TB-specific scripts** — Not copied (operational utilities for TB)                                                       |

### Infrastructure Changes

#### CB-Spider PostgreSQL METADB

**Major Change:** CB-Spider now uses PostgreSQL instead of SQLite for metadata storage (v0.12.30).

**docker-compose.yaml additions:**

```yaml
cb-spider:
  depends_on:
    cb-spider-postgres:
      condition: service_healthy
  environment:
    - SPIDER_METADB_ENDPOINT=cb-spider-postgres:5432
    - SPIDER_METADB_USER=${SP_POSTGRES_USER}
    - SPIDER_METADB_PASSWORD=${SP_POSTGRES_PASSWORD}
    - SPIDER_METADB_DATABASE=${SP_POSTGRES_DB}

cb-spider-postgres:
  image: postgres:16-alpine
  volumes:
    - ./data/cb-spider-container/meta_db/postgres:/var/lib/postgresql/data
```

**.env additions:**

```bash
SP_POSTGRES_USER=spider
SP_POSTGRES_PASSWORD=spider
SP_POSTGRES_DB=cb_spider
```

**Migration Impact:**

- Existing Spider SQLite data will be ignored
- CSP credentials need to be re-registered after upgrade
- Volume mount point changed: `meta_db/` → `meta_db/postgres/`

### Key Features Summary (v0.12.25→v0.12.30)

1. **Multi-Phase Bootstrap** (v0.12.29) - Sequential command execution with targeting
2. **Multi-AZ Distribution** (v0.12.29) - DistributeSubnets for high availability
3. **Resource Pruning** (v0.12.29) - Clean up orphaned metadata
4. **Credential Diagnostics** (v0.12.30) - Detailed verification error messages
5. **Namespace-Wide Listing** (v0.12.30) - Single-call resource queries across infras
6. **Default Security Group SSH-Only** (v0.12.30) - Breaking change, all ports now closed by default except SSH

**Comprehensive Upgrade Documentation:** See [docs/changes/BREAKING_CHANGES_v0.5.10.md](../../docs/changes/BREAKING_CHANGES_v0.5.10.md)

**Note:** Previous version history is tracked through Git. This document shows only the latest sync status.

- **CreateNodeGroupDynamicReq.DistributeSubnets**: Round-robin VM distribution across VNet subnets for multi-AZ spread
- **CreateNodeGroupDynamicReq.NodeUserPassword**: Field removed (replaced with comment explaining SSH-based access)
- **StatusCountInfo.CountReconciling**: New field for tracking reconciling status
- **AddNodeGroupDynamicReq**: New request type for adding node groups to existing infra
- **PostCommandStatus**: New type with constants (None, Completed, CompletedWithErrors, Failed, Skipped, Running)
- **PostCommandPhaseResult**: Result type for post-command phase execution
- **SshCmdResultForAPI / InfraSshCmdResultForAPI**: API response types for SSH command results
- **ResourcePruneResult / ResourcePruneResults**: New structs for orphaned metadata cleanup operations

- **Docker Compose**: Updated tumblebug (0.12.29), spider (0.12.40), and mapui (0.12.55) to latest compatible releases

- **Go Module**: Updated cb-tumblebug dependency to v0.12.29

- **Deployment Files**: Minor changes in conf/cloud_conf.yaml, init scripts, and templates not synced (beetle maintains independent deployment configuration)

**Key Feature Changes (v0.12.25→v0.12.29)**:

1. **Multi-phase Bootstrap**: PostCommands array replaces single PostCommand, enabling sequential bootstrap phases with per-phase targeting
2. **Subnet Distribution**: DistributeSubnets flag enables automatic VM distribution across availability zones via subnet round-robin
3. **Resource Pruning**: New API structures for cleaning up orphaned metadata entries
4. **Async Execution**: PostCommandAsync enables background command execution for long-running bootstrap operations

- `ResourceCountOverview` and `ResourcesByManageType` fields for Spider-Tumblebug resource reconciliation (not needed in beetle)
- OpenBao-related initialization scripts and preflight checks
- TB-specific operational scripts and configuration files

**Note**: TB v0.12.25 introduces enhanced OpenBao credential store validation and resource reconciliation features. These structs and scripts were not synchronized as they have no dependency chains to beetle's existing copied structs and represent TB-specific operational concerns.

## v0.12.13 Sync (2026-06-02)

Based on TB v0.12.13 `555a29bd` (tagged release).

| File                                       | Action                                                                                              |
| ------------------------------------------ | --------------------------------------------------------------------------------------------------- |
| `assets/assets.dump.gz`                    | **Updated** — MD5 changed to `9beccbd54b29...`                                                      |
| `assets/cloudimage.csv`                    | **Updated** — added 10 Tencent K8s node images (TencentOS, CentOS, Ubuntu, RHEL)                    |
| `assets/k8sclusterinfo.yaml`               | **Updated** — K8s version updates: AWS 1.35/1.34 added, Alibaba 1.35 added, Tencent 1.34/1.32 added |
| `assets/spider/.cloud-init-ibm/cloud-init` | **Updated** — migrated from bash script to cloud-config YAML format                                 |
| `init/README.md`                           | **Updated** — Python minimum version raised from 3.8 to 3.10                                        |
| `init/init.py`                             | **Updated** — improved statistics output format with percentage display                             |
| `scripts/restore-assets.sh`                | No change                                                                                           |
| `conf/cloud_conf.yaml`                     | No change                                                                                           |
| `assets/cloudinfo.yaml`                    | No change                                                                                           |
| `assets/cloudspec.csv`                     | No change                                                                                           |
| `assets/azure-publisher-filters.yaml`      | No change                                                                                           |
| `init/template.credentials.yaml`           | No change                                                                                           |
| `init/init.sh`                             | No change                                                                                           |
| `init/genCredential.sh`                    | No change                                                                                           |
| `init/initMetabase.sh`                     | No change                                                                                           |

## v0.12.12 Sync (2026-05-20)

Based on TB main HEAD `92979e93` (upstream/main).

| File                                  | Action                                                                                                                                                                                                                    |
| ------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `scripts/restore-assets.sh`           | **Updated** — container name auto-detection via `TB_POSTGRES_CONTAINER` env var; `DB_USER`/`DB_NAME` env-var-izable; identifier validation added; exact container name matching (`grep -Fxq`); SQL identifiers now quoted |
| `conf/cloud_conf.yaml`                | No change                                                                                                                                                                                                                 |
| `assets/cloudinfo.yaml`               | No change                                                                                                                                                                                                                 |
| `assets/cloudspec.csv`                | No change                                                                                                                                                                                                                 |
| `assets/azure-publisher-filters.yaml` | No change                                                                                                                                                                                                                 |
| `assets/assets.dump.gz`               | No change (MD5 identical)                                                                                                                                                                                                 |
| `init/template.credentials.yaml`      | No change                                                                                                                                                                                                                 |
| `init/init.sh`                        | No change                                                                                                                                                                                                                 |
| `init/genCredential.sh`               | No change                                                                                                                                                                                                                 |
| `init/initMetabase.sh`                | No change                                                                                                                                                                                                                 |

## v0.12.10 Sync (2026-05-14)

| File                                                  | Action                                                                                |
| ----------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `assets/azure-publisher-filters.yaml`                 | **Added** — new TB file for Azure VM image publisher filtering                        |
| `init/template.credentials.yaml`                      | **Updated** — added S3AccessKey/S3SecretKey fields for Azure, IBM, KT, NHN, OpenStack |
| `conf/cloud_conf.yaml`                                | No change                                                                             |
| `assets/assets.dump.gz`                               | No change (MD5 identical)                                                             |
| `init/init.sh`, `genCredential.sh`, `initMetabase.sh` | No change                                                                             |
| `scripts/restore-assets.sh`                           | No change                                                                             |

## Upstream Source Paths

| Local path       | Upstream path        |
| ---------------- | -------------------- |
| `conf/`          | `conf/`              |
| `assets/`        | `assets/`            |
| `init/`          | `init/`              |
| `scripts/`       | `scripts/`           |
| `interface/mcp/` | `src/interface/mcp/` |
