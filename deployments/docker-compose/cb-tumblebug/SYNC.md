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

## v0.13.8 Sync (2026-10-30)

Based on CB-Tumblebug tag `v0.13.8` (`19e8e8fdc1a0668ddb8f4abf3f66c44eb25d782a`). Upgrade path: **v0.13.7 &rarr; v0.13.8**.

### Model Changes (v0.13.7 &rarr; v0.13.8)

| Change | Description |
| --- | --- |
| **`SpecInfo`** & **`AvailabilityStatus`** | Added `AvailabilityStatus` string enum (`Available`, `SoldOut`, `QuotaExceeded`, `Unsupported`, `Unverified`) and runtime availability fields (`AvailabilityStatus`, `AvailabilityReason`, `AvailableZones`, `UnavailableZones`, `AvailabilityCheckedAt`, `AvailabilityRefreshAfter`) to `SpecInfo` |
| **`K8sClusterReq`**, **`K8sClusterDynamicReq`**, **`K8sClusterInfo`** | Added optional `Storage` fields (`*K8sClusterStorageReq`, `*K8sClusterStorageInfo`) and connected K8s abstract storage dependency types (`K8sStorageStatus`, `K8sStorageDriverRef`, `K8sStorageTypeState`, `K8sStorageAction`) |
| **`RDBMSInfo`** (`imdl/rdbms-model`) | Added optional `KeyValueList` (`json:"keyValueList,omitempty"`) field to propagate CSP-internal metadata (e.g., auto-created DB security groups) |
| **Model Version Headers** | Updated `copied-tb-model.go`, `copied-tb-k8s-model.go`, and `copied-tb-rdbms-model.go` version headers to CB-Tumblebug `v0.13.8` (`19e8e8fdc1a0668ddb8f4abf3f66c44eb25d782a`) |

### Deployment File Changes

| File | Action |
| --- | --- |
| **Model files (`imdl/`)** | |
| `imdl/cloud-model/copied-tb-model.go` | **Updated** - Synced version header to `v0.13.8` (`19e8e8fd`), added `AvailabilityStatus` and runtime availability fields on `SpecInfo` |
| `imdl/cloud-model/copied-tb-k8s-model.go` | **Updated** - Synced version header to `v0.13.8` (`19e8e8fd`), added `Storage` fields and K8s abstract storage dependency types |
| `imdl/rdbms-model/copied-tb-rdbms-model.go` | **Updated** - Synced version header to `v0.13.8` (`19e8e8fd`) and added `KeyValueList` to `RDBMSInfo` |
| **Go Module Dependencies (`go.mod`)** | |
| `go.mod` | **Updated** - `github.com/cloud-barista/cb-tumblebug` updated from `v0.13.7` to `v0.13.8`, `github.com/cloud-barista/cm-beetle/imdl` updated from `v0.1.18` to `v0.1.19` |
| **Docker Compose (`deployments/docker-compose/`)** | |
| `docker-compose.yaml` | **Updated** - Synced service images: `cb-spider:0.13.20`, `cb-mapui:0.13.14`, `cb-tumblebug:0.13.8` |
| **Assets & Scripts (`deployments/docker-compose/cb-tumblebug/`)** | |
| `assets/assets.dump.gz` | **Updated** - Synced latest asset dump (MD5 `d738ba9cd2b21ccdbcc3f5259dec9ea2`, 98,960 specs / 234,602 images across 10 CSPs) |
| `assets/assets.dump.gz.info` | **Updated** - Synced sidecar manifest matching `2026-10-29 16:41:54 UTC` asset dump |
| `assets/cloudimage.csv` | **Updated** - Synced enriched OS image catalog entries |
| `assets/cloudinfo.yaml` | **Updated** - Removed deprecated `cn-shenzhen-a` zone in Alibaba Cloud |
| `assets/cloudspec.csv` | **Updated** - Synced enriched VM spec and pricing catalog entries |
| `assets/diskinfo.yaml` | **No change** - Verified matches upstream |
| `assets/extractionpatterns.yaml` | **Updated** - Synced updated spec/image extraction patterns |
| `assets/k8sclusterinfo.yaml` | **No change** - Verified matches upstream |
| `assets/k8sstorageinfo.yaml` | **New** - Added Kubernetes abstract storage (`ssd-general`) and CSP StorageClass catalog |
| `assets/rdbmsinfo.yaml` | **Updated** - Added Alibaba `adminUserPasswordRequirement` and updated OpenStack Trove 2025.2 (Flamingo) active datastore versions (`mysql: 8.4`, `mariadb: 11.4`) |
| `conf/cloud_conf.yaml` | **Updated** - Updated `nlb.timeout` from `9` to `-1` to disable custom NLB health check timeout incompatible with Azure |
| `init/templates/*` | **Updated** - Synced verified CSP specs and Ubuntu image IDs across 8 `infra-*.json` templates |
| `init/*` | **Updated** - Synced `init.py` and `README.md` (`init.sh`, `cleanDB.sh`, `multi-init.sh`, `template.credentials.yaml` unchanged) |
| `scripts/*` | **Updated** - Synced `scripts/lib/pg-backend.sh` with SSH remote PostgreSQL backend support (`backup-assets.sh` and `restore-assets.sh` unchanged) |
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
