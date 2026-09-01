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

## v0.13.2 Sync (2026-09-01)

Based on CB-Tumblebug tag `v0.13.2` (`2a7436583f889cc794ebf37d151362a2e684e871`). Upgrade path: **v0.13.1 &rarr; v0.13.2**.

### Model Changes (v0.13.1 &rarr; v0.13.2)

| Change | Description |
| --- | --- |
| **`NodeInfo`**: `Failure` | Added structured `Failure *ProvisioningFailure` field representing node creation failure classification, zone attempts, and retryability |
| **`ProvisioningFailure` & `ZoneCapability`** | Added failure classification struct and zone shifting capability model with `Failure*` and `RetryHint*` constants |
| **`RDBMSDBMSRequirement`**: `DeprecatedVersions`, `EndOfLifeVersions` | Added version deprecation and EOL status tracking in managed RDBMS metadata |
| **`RDBMSCreateRequest` / `RDBMSInfo`**: `NHNDBSGToAllowAllInbound` | Added NHN Cloud DB security group inbound rule option |
| **`NodeSummary` & `InfraInfoSummary`** | Maintained lightweight projection structs for Infra list views (`GET /ns/{nsId}/infra`) |

### Deployment File Changes

| File | Action |
| --- | --- |
| **Model files (`imdl/cloud-model/`)** | |
| `copied-tb-model.go` | **Updated** — Synced with v0.13.2 (`2a743658`): added `NodeInfo.Failure`, `ProvisioningFailure`, `ZoneCapability`, failure constants, version header |
| `copied-tb-k8s-model.go` | **Updated** — Synced version header to `v0.13.2` (`2a743658`) |
| **Go Module Dependencies (`go.mod`)** | |
| `go.mod` | **Updated** — `github.com/cloud-barista/cb-tumblebug` updated from `v0.13.1` to `v0.13.2` |
| **Docker Compose (`deployments/docker-compose/`)** | |
| `docker-compose.yaml` | **Updated** — Synced service images: `cb-spider:0.13.2`, `cb-mapui:0.13.6`, `cb-tumblebug:0.13.2` |
| **Assets & Scripts (`deployments/docker-compose/cb-tumblebug/`)** | |
| `assets/assets.dump.gz` | **No change** — Verified MD5 `dba8da8e89b5ebdd203daf7c4480147b` (39MB dump matches upstream) |
| `assets/assets.dump.gz.info` | **No change** — Verified matches upstream sidecar manifest |
| `assets/k8sclusterinfo.yaml` | **No change** — Verified matches upstream K8s version and CSP support metadata |
| `assets/diskinfo.yaml` | **No change** — Verified matches upstream disk specifications and IOPS metadata |
| `assets/rdbmsinfo.yaml` | **No change** — Verified matches upstream multi-cloud RDBMS capability matrix |
| `assets/cloudinfo.yaml`, `extractionpatterns.yaml`, etc. | **No change** — Verified matches upstream provider configurations |
| `init/templates/*` | **No change** — Verified matches 28 standardized templates (`infra-*.json`, `k8scluster-across.json`, `sg-*.json`, `vnet-*.json`) |
| `scripts/lib/pg-backend.sh` | **No change** — Verified matches upstream Postgres backend detection library |
| `scripts/restore-assets.sh`, `backup-assets.sh` | **No change** — Verified matches upstream asset utilities |
| `interface/mcp/*` | **No change** — Verified matches upstream stateless HTTP mode and proxy configurations |

## Upstream Source Paths

| Local path | Upstream path |
| --- | --- |
| `conf/` | `conf/` |
| `assets/` | `assets/` |
| `init/` | `init/` |
| `scripts/` | `scripts/` |
| `interface/mcp/` | `src/interface/mcp/` |
