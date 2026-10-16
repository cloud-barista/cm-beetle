# CB-Tumblebug Files Sync Guide

Files under this directory are copied from the [CB-Tumblebug](https://github.com/cloud-barista/cb-tumblebug) repository.
When upgrading CB-Tumblebug, check each file against the upstream source and sync as needed.

## Sync Procedure

1. Check out the target CB-Tumblebug version:

   ```bash
   cd /path/to/cb-tumblebug
   git checkout v0.13.2
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
     echo "=== Checking new files in $dir/ ==="
     diff -qr $BEETLE/$dir $TB/$dir 2>&1 | grep "^Only in $TB"
   done

   # Note: assets/spider has been deprecated and removed
   # Note: assets/rdbmsinfo.yaml is maintained directly in cm-beetle for managed RDBMS feature
   ```

3. Review individual file changes and copy updated files:

   ```bash
   cp $TB/assets/*.yaml $BEETLE/assets/
   cp $TB/assets/*.csv $BEETLE/assets/
   cp $TB/assets/assets.dump.gz* $BEETLE/assets/
   cp -r $TB/src/interface/mcp/* $BEETLE/interface/mcp/
   ```

4. For binary assets (`assets.dump.gz`), compare checksums:
   ```bash
   md5sum $BEETLE/assets/assets.dump.gz $TB/assets/assets.dump.gz
   ```

---

## v0.13.7 Sync (2026-10-16)

Based on CB-Tumblebug tag `v0.13.7` (`09a796e61ef128e9ac0dbddc7911e8a86e55b2f7`). Upgrade path: **v0.13.5 &rarr; v0.13.7**.

### Model Changes (v0.13.5 &rarr; v0.13.7)

| Change | Description |
| --- | --- |
| **`RDBMSDatabaseCreateReq`** | Added optional `Charset` (`json:"charset,omitempty"`) and `Collation` (`json:"collation,omitempty"`) fields for MySQL/MariaDB logical database creation |
| **`TargetDatabase`** (`imdl/rdbms-model`) | Added `Charset` and `Collation` fields to propagate recommended database character set and collation into the RDBMS migration pipeline |
| **Model Version Headers** | Updated `copied-tb-model.go`, `copied-tb-k8s-model.go`, and `copied-tb-rdbms-model.go` version headers to CB-Tumblebug `v0.13.7` (`09a796e61ef128e9ac0dbddc7911e8a86e55b2f7`) |

### Deployment File Changes

| File | Action |
| --- | --- |
| **Model files (`imdl/`)** | |
| `imdl/cloud-model/copied-tb-model.go` | **Updated** - Synced version header to `v0.13.7` (`09a796e6`) |
| `imdl/cloud-model/copied-tb-k8s-model.go` | **Updated** - Synced version header to `v0.13.7` (`09a796e6`) |
| `imdl/rdbms-model/copied-tb-rdbms-model.go` | **Updated** - Synced version header to `v0.13.7` (`09a796e6`) and added `Charset`/`Collation` fields to `RDBMSDatabaseCreateReq` |
| **Go Module Dependencies (`go.mod`)** | |
| `go.mod` | **Updated** - `github.com/cloud-barista/cb-tumblebug` updated from `v0.13.5` to `v0.13.7` |
| **Docker Compose (`deployments/docker-compose/`)** | |
| `docker-compose.yaml` | **Updated** - Synced service images: `cb-spider:0.13.17`, `cb-mapui:0.13.13`, `cb-tumblebug:0.13.7` |
| **Assets & Scripts (`deployments/docker-compose/cb-tumblebug/`)** | |
| `assets/assets.dump.gz` | **No change** - Verified MD5 `dba8da8e89b5ebdd203daf7c4480147b` (matches upstream) |
| `assets/assets.dump.gz.info` | **No change** - Verified matches upstream sidecar manifest |
| `assets/diskinfo.yaml` | **No change** - Verified matches upstream |
| `assets/k8sclusterinfo.yaml` | **No change** - Verified matches upstream |
| `assets/rdbmsinfo.yaml` | **Updated** - Synced `databaseRequirements.note` (charset/collation support details across Azure, GCP, Alibaba, Tencent, IBM, NCP, NHN, OpenStack) and CSP notes |
| `conf/cloud_conf.yaml` | **No change** - Verified matches upstream |
| `init/templates/*` | **Updated** - Synced Alibaba spec updates (`infra-across-global.json`, `infra-across-scale-test.json`, `infra-alibaba-global.json`) and added `infra-all-csp.json`, `infra-db-bastion.json` |
| `init/*` | **No change** - Verified `init.sh`, `cleanDB.sh`, `init.py`, `multi-init.sh`, and `template.credentials.yaml` match upstream |
| `scripts/*` | **No change** - Verified `backup-assets.sh`, `restore-assets.sh`, and `lib/pg-backend.sh` match upstream |
| `interface/mcp/*` | **No change** - Verified 1:1 match with upstream `src/interface/mcp` |

> **Note:** OpenBao files are maintained independently at `deployments/docker-compose/openbao/` with its own `SYNC.md`, not under `cb-tumblebug/init/openbao/`.

## Upstream Source Paths

| Local path | Upstream path |
| --- | --- |
| `conf/` | `conf/` |
| `assets/` | `assets/` |
| `init/` | `init/` |
| `scripts/` | `scripts/` |
| `interface/mcp/` | `src/interface/mcp/` |
