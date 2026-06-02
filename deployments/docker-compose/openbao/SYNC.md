# OpenBao Files Sync Guide

> **Note:** These files originate from [MC-Terrarium](https://github.com/cloud-barista/mc-terrarium) and are distributed through CB-Tumblebug.

Sync from CB-Tumblebug's bundled copy (`cb-tumblebug/init/openbao/`), not directly from MC-Terrarium, to ensure version compatibility with the integrated Tumblebug release.

## Upstream Source

```
cb-tumblebug/init/openbao/
```

CM-Beetle's copy lives at:

```
deployments/docker-compose/openbao/
```

## Sync Procedure

1. Diff each file against the upstream:

   ```bash
   TB=/path/to/cb-tumblebug/init/openbao
   OPENBAO=deployments/docker-compose/openbao

   diff $OPENBAO/openbao-config.hcl           $TB/openbao-config.hcl
   diff $OPENBAO/openbao-init.sh              $TB/openbao-init.sh
   diff $OPENBAO/openbao-unseal.sh            $TB/openbao-unseal.sh
   diff $OPENBAO/openbao-register-creds.py    $TB/openbao-register-creds.py
   diff $OPENBAO/openbao-register-creds.sh    $TB/openbao-register-creds.sh
   diff $OPENBAO/pyproject.toml               $TB/pyproject.toml
   ```

2. Apply functional changes while preserving any CM-Beetle-specific customizations (e.g., comment style).

## v0.12.13 Sync (2026-06-02)

Based on TB v0.12.13 `555a29bd` (tagged release).

| File                        | Action                                                                                                           |
| --------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `pyproject.toml`            | **Updated** — Python minimum version raised from 3.8 to 3.10; requests library updated from 2.32.4 to 2.33.1    |
| `openbao-config.hcl`        | No change                                                                                                        |
| `openbao-init.sh`           | No change (upstream comment style differs — CM-Beetle unicode style kept)                                        |
| `openbao-unseal.sh`         | No change (upstream comment style differs — CM-Beetle unicode style kept)                                        |
| `openbao-register-creds.py` | No change (upstream has minor comment text rewording and import reordering — CM-Beetle unicode style kept)       |
| `openbao-register-creds.sh` | No change                                                                                                        |

## v0.12.12 Sync (2026-05-20)

Based on TB main HEAD `92979e93` (upstream/main).

| File                        | Action                                                                                    |
| --------------------------- | ----------------------------------------------------------------------------------------- |
| `openbao-config.hcl`        | No change                                                                                 |
| `openbao-init.sh`           | No change (upstream comment style differs — CM-Beetle unicode style kept)                 |
| `openbao-unseal.sh`         | No change (upstream comment style differs — CM-Beetle unicode style kept)                 |
| `openbao-register-creds.py` | No change (upstream has minor comment text rewording only — CM-Beetle unicode style kept) |
| `openbao-register-creds.sh` | No change                                                                                 |

## v0.12.10 Sync (2026-05-14)

| File                        | Action                                                                                                              |
| --------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `openbao-register-creds.py` | **Updated** — added `registered_providers` set for full path tracking; refined placeholder comment and print format |
| `openbao-register-creds.sh` | **Updated** — minimum Python version raised from 3.8 to 3.10                                                        |
| `openbao-config.hcl`        | No change                                                                                                           |
| `openbao-init.sh`           | No change (upstream differs only in comment style — CM-Beetle unicode style kept)                                   |
| `openbao-unseal.sh`         | No change (upstream differs only in comment style — CM-Beetle unicode style kept)                                   |
