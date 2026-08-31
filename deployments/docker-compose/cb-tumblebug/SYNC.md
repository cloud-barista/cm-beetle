# CB-Tumblebug Files Sync Guide

Files under this directory are copied from the [CB-Tumblebug](https://github.com/cloud-barista/cb-tumblebug) repository.
When upgrading CB-Tumblebug, check each file against the upstream source and sync as needed.

## Sync Procedure

1. Check out the target CB-Tumblebug version:

   ```bash
   cd /path/to/cb-tumblebug
   git checkout main && git pull upstream main
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

   # Note: assets/spider is excluded from synchronization (cb-tumblebug issue #2694)
   # Note: assets/rdbmsinfo.yaml is required for Tumblebug managed RDBMS feature
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

## Latest Main Branch Sync (2026-08-31)

Based on CB-Tumblebug latest `main` branch `a9a10473` (`Merge pull request #2741 from seokho-son/main`). Upgrade path: **v0.12.30 → latest main**.

### Model Changes (v0.12.30 → latest main)

| Change | Description |
| --- | --- |
| **`ImageInfo`**: `SystemMessage`, `DeletionRequestedAt` | Added error message tracking and deletion tombstone tracking |
| **`NLBInfo`**: `SystemMessage`, `DeletionRequestedAt` | Added system message and deletion tombstone tracking |
| **`NodeSummary`** & **`InfraInfoSummary`** | Added lightweight projection structs for Infra list views (`GET /ns/{nsId}/infra`) |
| **`SecurityGroupInfo` / `SshKeyInfo`**: `Conditions`, `DeletionRequestedAt`, `Status` | Added lifecycle status, K8s-style conditions, and tombstone tracking |
| **`DataDiskInfo`**: `DiskFailed`, `Conditions`, `DeletionRequestedAt` | Added failure enum and condition tracking |
| **`RDBMS` Models**: `RDBMSInfo`, `RDBMSCreateRequest`, `RDBMSCapabilityResponse`, `RDBMSSupportResponse`, `RDBMSCSPSupportInfo`, `RDBMSDatabaseCreateReq`, `RDBMSDatabaseInfo`, `RDBMSDatabaseListResponse` | Managed RDBMS models synchronized with Spider v0.13.1 contract and multi-cloud DB engines |

### Deployment File Changes

| File | Action |
| --- | --- |
| **Model files (`imdl/cloud-model/`)** | |
| `copied-tb-model.go` | **Updated** — Synced with latest main `a9a10473`: added `NodeSummary`, `InfraInfoSummary`, `ImageInfo` and `NLBInfo` fields, version header |
| `copied-tb-k8s-model.go` | **Updated** — Synced full K8s cluster models (`K8sClusterInfo`, `K8sNodeGroupInfo`, `K8sClusterDynamicReq`, `K8sNodeGroupDynamicReq`, `K8sMultiClusterDynamicReq`, `K8sMultiClusterInfo`, `K8sAccessInfo`, `K8sClusterTokenResponse`, `K8sClusterKubeconfigResponse`, `IID`) |
| **Docker Compose (`deployments/docker-compose/`)** | |
| `docker-compose.yaml` | **Updated** — cb-spider `0.13.1` with `GOMEMLIMIT=8GiB`, cb-mapui `0.13.5` with `MAPUI_PARAM_*` optional variables, openbao `2.5.1` |
| **Assets & Scripts (`deployments/docker-compose/cb-tumblebug/`)** | |
| `assets/assets.dump.gz` | **Updated** — Synced with latest 39MB dump |
| `assets/assets.dump.gz.info` | **Updated** — Synced manifest sidecar |
| `assets/k8sclusterinfo.yaml` | **Updated** — Synced K8s version and CSP support metadata |
| `assets/diskinfo.yaml` | **Added** — Synced disk specifications and IOPS metadata |
| `assets/rdbmsinfo.yaml` | **Updated** — Synced multi-cloud RDBMS capability matrix |
| `assets/cloudinfo.yaml`, `extractionpatterns.yaml`, etc. | **Updated** — Synced latest provider configurations |
| `init/templates/*` | **Updated** — Synced 28 standardized templates (`infra-*.json`, `k8scluster-across.json`, `sg-*.json`, `vnet-*.json`) |
| `scripts/lib/pg-backend.sh` | **Added** — Synced Postgres backend detection library for Docker/Local/K8s environments |
| `scripts/restore-assets.sh`, `backup-assets.sh` | **Updated** — Synced Postgres asset restore and backup utilities |
| `interface/mcp/*` | **Updated** — Synced stateless HTTP mode, label management, search options, and proxy configurations |

## Upstream Source Paths

| Local path | Upstream path |
| --- | --- |
| `conf/` | `conf/` |
| `assets/` | `assets/` |
| `init/` | `init/` |
| `scripts/` | `scripts/` |
| `interface/mcp/` | `src/interface/mcp/` |
