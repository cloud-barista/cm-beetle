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

## v0.13.5 Sync (2026-10-12)

Based on CB-Tumblebug tag `v0.13.5` (`16407239df37b38e8e81ce64b81bff0e6a42e370`). Upgrade path: **v0.13.3 &rarr; v0.13.5**.

### Model Changes (v0.13.3 &rarr; v0.13.5)

| Change | Description |
| --- | --- |
| **`NodeInfo`**: `RootDeviceName` | Updated JSON serialization tag from `json:"RootDeviceName"` to `json:"rootDeviceName"` matching upstream CB-Tumblebug standard casing |
| **`k8sClusterDetailExt`** | Removed redundant `NodeSpecNamingRule` field from extension struct as it is now natively included in upstream `tbmodel.K8sClusterDetail` |
| **Model Version Headers** | Updated `copied-tb-model.go`, `copied-tb-k8s-model.go`, and `copied-tb-rdbms-model.go` version headers to CB-Tumblebug `v0.13.5` (`16407239df37b38e8e81ce64b81bff0e6a42e370`) |

### Deployment File Changes

| File | Action |
| --- | --- |
| **Model files (`imdl/`)** | |
| `imdl/cloud-model/copied-tb-model.go` | **Updated** — Synced version header to `v0.13.5` (`16407239`), updated `RootDeviceName` tag to `json:"rootDeviceName"` |
| `imdl/cloud-model/copied-tb-k8s-model.go` | **Updated** — Synced version header to `v0.13.5` (`16407239`) |
| `imdl/rdbms-model/copied-tb-rdbms-model.go` | **Updated** — Synced version header to `v0.13.5` (`16407239`) |
| **Go Module Dependencies (`go.mod`)** | |
| `go.mod` | **Updated** — `github.com/cloud-barista/cb-tumblebug` updated from `v0.13.3` to `v0.13.5` |
| **Docker Compose (`deployments/docker-compose/`)** | |
| `docker-compose.yaml` | **Updated** — Synced service images: `cb-spider:0.13.14`, `cb-mapui:0.13.13`, `cb-tumblebug:0.13.5` |
| **Assets & Scripts (`deployments/docker-compose/cb-tumblebug/`)** | |
| `assets/assets.dump.gz` | **No change** — Verified MD5 `dba8da8e89b5ebdd203daf7c4480147b` (matches upstream) |
| `assets/assets.dump.gz.info` | **No change** — Verified matches upstream sidecar manifest |
| `assets/diskinfo.yaml` | **Updated** — Synced KT default root/data disk type to SSD (and diskTypeOrder [SSD, HDD]) |
| `assets/k8sclusterinfo.yaml` | **Updated** — Synced K8s version matrices (AWS, Azure, GCP, Alibaba, IBM) and nodeGroupNamingRule pattern |
| `assets/rdbmsinfo.yaml` | **Maintained** — Preserved latest CM-Beetle RDBMS metadata |
| `conf/cloud_conf.yaml` | **Updated** — Synced nlbInfraSpecId (`aws+ap-northeast-2+t2.small`) and nlbInfraImageId (`ubuntu 22.04`) |
| `init/templates/*` | **Updated** — Synced standardized templates (including `infra-aws-global-extream.json`, `infra-global-extream.json`, and CSP region adjustments) |
| `init/*` | **Updated** — Synced `init.sh`, `cleanDB.sh`, `init.py`, `multi-init.sh`, and `template.credentials.yaml` |
| `scripts/*` | **No change** — Verified `backup-assets.sh`, `restore-assets.sh`, and `lib/pg-backend.sh` match upstream |
| `interface/mcp/*` | **No change** — Verified 1:1 match with upstream `src/interface/mcp` |

> **Note:** OpenBao files are maintained independently at `deployments/docker-compose/openbao/` with its own `SYNC.md`, not under `cb-tumblebug/init/openbao/`.

## Upstream Source Paths

| Local path | Upstream path |
| --- | --- |
| `conf/` | `conf/` |
| `assets/` | `assets/` |
| `init/` | `init/` |
| `scripts/` | `scripts/` |
| `interface/mcp/` | `src/interface/mcp/` |
