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

## v0.12.22 Sync (2026-07-01)

Based on TB v0.12.22 `50c213b8` (tagged release).

| File                                  | Action                                                                                                                                                          |
| ------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Model files (imdl/cloud-model/)**   |                                                                                                                                                                 |
| `copied-tb-model.go`                  | **Updated** — Added `CommandStatusCompletedWithError` constant to distinguish remote command non-zero exits from SSH/transport failures                         |
| **docker-compose.yaml**               |                                                                                                                                                                 |
| `docker-compose.yaml`                 | **Updated** — cb-tumblebug `0.12.19→0.12.22`, cb-mapui `0.12.43→0.12.47`, cb-spider unchanged `0.12.32`                                                        |
| **Assets**                            |                                                                                                                                                                 |
| `assets/assets.dump.gz`               | **Updated** — MD5 changed from `77888ce732683fef...` to `9db6eb7212e1450918145d664bbef465` (copied to beetle)                                                  |
| `assets/cloudimage_ignore.yaml`       | **Updated** — Added Azure Hyper-V Generation 1 metadata filters to exclude legacy Gen1 images (modern Azure VMs require Gen2)                                  |
| `assets/cloudinfo.yaml`               | **Updated** — CSP information and region updates (copied to beetle)                                                                                             |
| `assets/k8sclusterinfo.yaml`          | **Updated** — K8s cluster version and configuration updates (copied to beetle)                                                                                  |
| **Init**                              |                                                                                                                                                                 |
| `init/decCredential.sh`               | **Differs** — TB version updated but not copied (beetle uses existing version)                                                                                  |
| `init/openbao/`                       | **New in TB** — OpenBao-related files (not copied to beetle - not needed)                                                                                       |
| `init/templates/*.json`               | **New templates in TB** — Many new infra/usecase templates added upstream; not copied (beetle manages its own templates)                                        |
| **Config**                            |                                                                                                                                                                 |
| `conf/setup.env`, `conf/traefik.yaml` | **New in TB** — TB-specific configuration files (not copied to beetle - not needed for beetle's deployment)                                                     |
| **Scripts**                           |                                                                                                                                                                 |
| `scripts/*`                           | **Many new scripts in TB** — Operational scripts for backup, etcd, NLB, usecases, etc. (not copied to beetle)                                                  |

**Summary of Changes:**

- **Models**: Added `CommandStatusCompletedWithError` status constant for better SSH command result handling
- **Docker Compose**: Updated tumblebug and mapui versions; spider version remains unchanged
- **Assets**: Updated binary dump, cloud image ignore patterns with Azure Gen1 filters, cloud info, and K8s cluster configurations
- **Upstream additions not copied**: New infraAutopilot model structs (16 new types for declarative resilient infra provisioning), init templates, TB-specific configs, and operational scripts

**Note**: TB v0.12.22 introduces extensive infraAutopilot functionality (availability query GPU support, image metadata filters, and 16 new structs for resilient provisioning). These are standalone features without dependency chains to beetle's existing copied structs, so they were not synchronized per dependency-based sync rules.

## v0.12.19 Sync (2026-06-24)

Based on TB v0.12.19 `6b0b1102` (tagged release).

| File                                  | Action                                                                                                                                                          |
| ------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Model files (imdl/cloud-model/)**   |                                                                                                                                                                 |
| `copied-tb-model.go`                  | **Header updated** — No struct changes for copied types; upstream added `NLBFeatureSupport`/`NLBSupportResponse` in nlb.go but not copied (not needed by beetle) |
| **docker-compose.yaml**               |                                                                                                                                                                 |
| `docker-compose.yaml`                 | **Updated** — cb-tumblebug `0.12.15→0.12.19`, cb-spider `0.12.30→0.12.32`, cb-mapui `0.12.39→0.12.43`                                                         |
| **Assets**                            |                                                                                                                                                                 |
| `assets/assets.dump.gz`               | No change (MD5 identical: `77888ce732683fef...`)                                                                                                                |
| `assets/cloudimage.csv`               | **Updated** — Added 2 Alibaba K8s node image entries (`AliyunLinux3ContainerOptimized`, `Ubuntu`) for ACK node pool support                                    |
| `assets/k8sclusterinfo.yaml`          | **Updated** — Added `initialNodeGroupManagedByCluster: true` for Alibaba and Tencent clusters (initial node group lifecycle-bound to cluster)                   |
| **Interface/MCP**                     |                                                                                                                                                                 |
| `interface/mcp/tb-mcp.py`             | **Updated** — Migrated from SSE to Streamable HTTP transport; updated imports (`mcp.server.fastmcp` → `fastmcp`); removed monitoring policy comment             |
| `interface/mcp/Dockerfile`            | **Updated** — Changed run command from `fastmcp run --transport sse` to `uv run ./tb-mcp.py`                                                                   |
| `interface/mcp/architecture.md`       | **Updated** — Diagrams updated to reflect Streamable HTTP transport replacing SSE                                                                               |
| `interface/mcp/README.md`             | **Updated** — Documentation updated for new transport mode                                                                                                      |
| `interface/mcp/PROXY_README.md`       | **New** — Documentation for mcp-simple-proxy.py usage (copied to beetle)                                                                                       |
| `interface/mcp/claude_desktop_config.json` | **New** — Claude Desktop configuration example (copied to beetle)                                                                                          |
| `interface/mcp/mcp-simple-proxy.py`   | **New** — Simple stdio-to-Streamable-HTTP proxy for Claude Desktop integration (copied to beetle)                                                               |
| **Init**                              |                                                                                                                                                                 |
| `init/init.py`                        | No change                                                                                                                                                       |
| `init/templates/`                     | **New templates in TB** — Many new infra/usecase/sg/vnet templates added; not copied (beetle manages its own templates)                                         |
| **Config**                            |                                                                                                                                                                 |
| `conf/cloud_conf.yaml`                | No change                                                                                                                                                       |

**Summary of Changes:**

- **Models**: No struct changes for beetle's copied types; only new standalone NLB support structs added upstream
- **Docker Compose**: Updated all three service versions (tumblebug, spider, mapui)
- **Assets**: K8s cluster info updated for Alibaba ACK and Tencent TKE initial node group management; cloudimage.csv updated with Alibaba image types
- **MCP Interface**: Major migration from SSE to Streamable HTTP transport; added proxy support for Claude Desktop; three new files added

## v0.12.15 Sync (2026-06-15)

Based on TB v0.12.15 `4f01927b` (tagged release).

| File                                  | Action                                                                                                                                                 |
| ------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Model files (imdl/cloud-model/)**   |                                                                                                                                                        |
| `copied-tb-model.go`                  | **Updated** — Added `IsBasicGpuImage` field to `ImageInfo` struct; updated example values for `VNetTemplateId` and `SgTemplateId` in `InfraDynamicReq` |
| **Assets**                            |                                                                                                                                                        |
| `assets/assets.dump.gz`               | **Updated** — MD5 changed from `9beccbd54b29...` to `77888ce732683fef...` (copied to beetle)                                                           |
| `assets/cloudimage_ignore.yaml`       | **New** — Cloud image ignore patterns configuration (copied to beetle)                                                                                 |
| `assets/extractionpatterns.yaml`      | **Updated** — Added `gpuExcludePatterns` section and `basicGpuImageRules` with per-CSP GPU image identification rules (copied to beetle)               |
| `assets/k8sclusterinfo.yaml`          | **Updated** — IBM Cloud K8s versions updated (added 1.35, updated 1.34, 1.33; removed older versions) (copied to beetle)                               |
| **Init**                              |                                                                                                                                                        |
| `init/init.py`                        | **Updated** — Added `k8sCluster` template type detection for K8s multi-cluster dynamic provisioning (copied to beetle)                                 |
| `init/openbao/`                       | **New** — OpenBao-related initialization files (not copied to beetle - not needed for beetle's deployment)                                             |
| `init/templates/*.json`               | **Reorganized** — TB renamed templates with resource-type prefixes (e.g., `default-sg.json` → `sg-default.json`). Beetle keeps current template names. |
| **Scripts**                           |                                                                                                                                                        |
| `scripts/*`                           | **Many new scripts** — TB added many operational scripts (not copied to beetle - beetle only needs `restore-assets.sh`)                                |
| **Config**                            |                                                                                                                                                        |
| `conf/setup.env`, `conf/traefik.yaml` | **New in TB** — TB-specific configuration files (not copied to beetle - not needed for beetle's deployment)                                            |

**Summary of Changes:**

- **Models**: Synchronized `IsBasicGpuImage` field addition and example value updates
- **Assets**: Updated binary dump, extraction patterns for GPU image detection, K8s cluster info, and added new cloudimage_ignore.yaml (all copied to beetle)
- **Init**: Updated init.py with K8s cluster template support (copied to beetle)
- **K8s**: Added K8s cluster template support and version updates
- **Templates**: TB reorganized template naming (beetle keeps current names for now)
- **Not Copied**: TB-specific scripts, configs (setup.env, traefik.yaml), openbao init files, and operational scripts not needed for beetle's docker-compose deployment

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
