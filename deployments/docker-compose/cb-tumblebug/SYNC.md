# CB-Tumblebug Files Sync Guide

Files under this directory are copied from the [CB-Tumblebug](https://github.com/cloud-barista/cb-tumblebug) repository.
When upgrading CB-Tumblebug, check each file against the upstream source and sync as needed.

## Sync Procedure

1. Check out the target CB-Tumblebug version:

   ```bash
   cd /path/to/cb-tumblebug
   git checkout v0.14.0
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

## v0.14.0 Sync (2026-11-25)

Based on CB-Tumblebug tag `v0.14.0` (`c05bc2c5b926859a04eac4d4e2a8c0db9e32ad4a`). Upgrade path: **v0.13.8 -> v0.14.0**.

### Model Changes (v0.13.8 -> v0.14.0)

| Change | Description |
| --- | --- |
| **`InfraReq`**, **`InfraDynamicReq`**, **`InfraInfo`**, **`InfraInfoSummary`** | Replaced `PostCommands`, `PostCommandAsync`, `PostCommandResults`, `PostCommandStatus`, `PostCommandRequestId` with `PostScript` (`*PostScriptReq`), `PostScriptAsync` (`json:"-"`), `PostScriptResults` (`[]PostScriptPhaseResult`), `PostScriptStatus` (added `PostScriptStatusInterrupted`), `PostScriptRequestId`; removed deprecated `InstallMonAgent` |
| **`PostScriptReq`**, **`PostScriptPhase`**, **`PostScriptPhaseResult`**, **`PostScriptStatus`** | Renamed from `PostCommandReq` / `PostCommandPhaseResult` / `PostCommandStatus` and wrapped in `PostScriptReq{Phases}` |
| **`InitScriptReq`**, **`InitScriptStep`**, **`InfraCmdReq`** | Added `InitScript *InitScriptReq` to `CreateNodeGroupReq`, `InfraDynamicReq`, `CreateNodeGroupDynamicReq`, `NodeInfo`, `NodeSummary` (with `InitScriptStatus`, `InitScriptMessage`, `InitScriptSteps`); added `TemplateId`, `Scripts`, `Env` to `InfraCmdReq` (`Command` tag updated to `required_without_all=TemplateId Scripts`) |
| **`NodeInfo`**, **`NodeSummary`** | Removed deprecated `MonAgentStatus` |
| **`SpecInfo`**, **`SpecSummary`**, **`AvailabilityQuota`** | Added `AcceleratorPartitionSize float32` and `AvailabilityQuota *AvailabilityQuota` |
| **Model Version Headers** | Updated `copied-tb-model.go`, `copied-tb-k8s-model.go` (no struct changes), and `copied-tb-rdbms-model.go` (no struct changes) to CB-Tumblebug `v0.14.0` (`c05bc2c5b926859a04eac4d4e2a8c0db9e32ad4a`) |

### Deployment File Changes

| File | Action |
| --- | --- |
| **Model files (`imdl/`)** | |
| `imdl/cloud-model/copied-tb-model.go` | **Updated** - Synced structs and version header to `v0.14.0` (`c05bc2c5`) |
| `imdl/cloud-model/copied-tb-k8s-model.go` | **Updated** - Synced version header to `v0.14.0` (`c05bc2c5`, no struct changes) |
| `imdl/rdbms-model/copied-tb-rdbms-model.go` | **Updated** - Synced version header to `v0.14.0` (`c05bc2c5`, no struct changes) |
| **Go Module Dependencies (`go.mod`)** | |
| `go.mod` | **Updated** - `github.com/cloud-barista/cb-tumblebug` updated from `v0.13.8` to `v0.14.0` |
| **Docker Compose (`deployments/docker-compose/`)** | |
| `docker-compose.yaml` | **Updated** - Synced service images (`cb-tumblebug:0.14.0`, `cb-spider:0.14.0`, `cb-mapui:0.14.0`, commented `cb-tumblebug-mcp-server:0.14.0`) and removed legacy `./data/cb-tumblebug-container/meta_db/:/app/meta_db/` volume mount |
| **Assets & Scripts (`deployments/docker-compose/cb-tumblebug/`)** | |
| `assets/assets.dump.gz` | **Updated** - Synced latest asset dump (MD5 `f0e35d5e35cb4760736e002ef5c0d595`, 45M, `2026-10-30 05:26:01 UTC`, 98,960 specs / 234,602 images across 10 CSPs, KT priced 30% -> 35%) |
| `assets/assets.dump.gz.info` | **Updated** - Synced sidecar manifest matching updated asset dump |
| `assets/cloudspec.csv` | **Updated** - Added `acceleratorPartitionSize` column and updated KT/cloud spec entries (533 -> 543 lines) |
| `assets/rdbmsinfo.yaml` | **Updated** - Added CB-Spider #1886/#1890 and CB-Tumblebug #2829 private-mode HTTP 422 fast-fail notes and NHN `nhnDBSGToAllowAllInbound` details |
| `assets/*.yaml`, `assets/*.csv` (others) | **No change** - Verified matches upstream (`cloudimage.csv`, `cloudinfo.yaml`, `diskinfo.yaml`, `extractionpatterns.yaml`, `k8sclusterinfo.yaml`, `k8sstorageinfo.yaml`, `networkinfo.yaml`, etc.) |
| `conf/cloud_conf.yaml` | **No change** - Verified matches upstream |
| `init/*` | **Updated** - Synced `init.py` (initScript template detection and `/readyz/init` `assetsSummary` check), `multi-init.sh`, and `template.credentials.yaml` (NHN endpoint normalized) |
| `init/templates/*` | **Updated** - Mirrored 1:1 with upstream; updated `infra-*.json` templates and added `infra-ntfy.json`, `infra-ntfy-global.json`, `infra-obs.json`, `initscript-hello.json`, `initscript-nginx.json`, `initscript-otel-agent.json`, `initscript-swap.json`, `sg-ntfy.json`, `sg-obs.json` |
| `scripts/restore-assets.sh`, `scripts/lib/pg-backend.sh` | **Updated** - Synced compressed `.gz` direct SSH transfer and `LogLevel=ERROR` improvements (`backup-assets.sh` unchanged) |
| `interface/mcp/*` | **Updated** - Mirrored upstream `src/interface/mcp/*` (`Dockerfile` +`pyyaml==6.0.4`, `tb-mcp.py`, `mcp-simple-proxy.py`, `README.md`, `PROXY_README.md`, `architecture.md`, and `tests/`) |

> **Note:** OpenBao files are maintained independently at `deployments/docker-compose/openbao/` with its own `SYNC.md`, not under `cb-tumblebug/init/openbao/`. (Verified 0 changes in upstream `init/openbao/` between `v0.13.8` and `v0.14.0`.)

## Upstream Source Paths

| Local path | Upstream path |
| --- | --- |
| `conf/` | `conf/` |
| `assets/` | `assets/` |
| `init/` | `init/` |
| `scripts/` | `scripts/` |
| `interface/mcp/` | `src/interface/mcp/` |
