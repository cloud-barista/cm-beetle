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
