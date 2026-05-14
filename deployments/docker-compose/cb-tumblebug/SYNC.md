# CB-Tumblebug Files Sync Guide

Files under this directory are copied from the [CB-Tumblebug](https://github.com/cloud-barista/cb-tumblebug) repository.
When upgrading CB-Tumblebug, check each file against the upstream source and sync as needed.

## Sync Procedure

1. Check out the target CB-Tumblebug version:
   ```bash
   cd /path/to/cb-tumblebug
   git checkout v<VERSION>
   ```

2. Diff each tracked file against the upstream:
   ```bash
   TB=/path/to/cb-tumblebug
   BEETLE=deployments/docker-compose/cb-tumblebug

   diff $BEETLE/conf/cloud_conf.yaml       $TB/conf/cloud_conf.yaml
   diff $BEETLE/assets/cloudinfo.yaml      $TB/assets/cloudinfo.yaml
   diff $BEETLE/assets/cloudspec.csv       $TB/assets/cloudspec.csv
   diff $BEETLE/init/template.credentials.yaml $TB/init/template.credentials.yaml
   diff $BEETLE/init/init.sh               $TB/init/init.sh
   diff $BEETLE/init/genCredential.sh      $TB/init/genCredential.sh
   diff $BEETLE/init/initMetabase.sh       $TB/init/initMetabase.sh
   diff $BEETLE/scripts/restore-assets.sh  $TB/deployments/docker-compose/scripts/restore-assets.sh
   ```

3. Copy new or updated files:
   ```bash
   cp $TB/assets/azure-publisher-filters.yaml $BEETLE/assets/
   ```

4. For binary assets (`assets.dump.gz`), compare checksums:
   ```bash
   md5sum $BEETLE/assets/assets.dump.gz $TB/assets/assets.dump.gz
   ```

## v0.12.10 Sync (2026-05-14)

| File | Action |
|------|--------|
| `assets/azure-publisher-filters.yaml` | **Added** — new TB file for Azure VM image publisher filtering |
| `init/template.credentials.yaml` | **Updated** — added S3AccessKey/S3SecretKey fields for Azure, IBM, KT, NHN, OpenStack |
| `conf/cloud_conf.yaml` | No change |
| `assets/assets.dump.gz` | No change (MD5 identical) |
| `init/init.sh`, `genCredential.sh`, `initMetabase.sh` | No change |
| `scripts/restore-assets.sh` | No change |

## Upstream Source Paths

| Local path | Upstream path |
|------------|---------------|
| `conf/` | `conf/` |
| `assets/` | `assets/` |
| `init/` | `init/` |
| `scripts/` | `deployments/docker-compose/scripts/` |
