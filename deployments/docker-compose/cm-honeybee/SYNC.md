# CM-Honeybee Assets & Configurations Sync Guide

> **Note:** These files originate from [CM-Honeybee](https://github.com/cloud-barista/cm-honeybee).

## Upstream Source

```
https://github.com/cloud-barista/cm-honeybee/tree/main/server/
```

CM-Beetle's copy lives at:

```
deployments/docker-compose/cm-honeybee/
```

## Structure & Mappings

| Local Path in CM-Beetle | Upstream Source Path in CM-Honeybee | Description |
| :--- | :--- | :--- |
| `openbao/openbao-config.hcl` | `server/openbao/openbao-config.hcl` | OpenBao configuration for Honeybee's dedicated secrets backend |

## v0.6.3 Sync (2026-11-04)

- Bumped `cm-honeybee` service image to `cloudbaristaorg/cm-honeybee:0.6.3` (includes `gpuCards` refinement in `/infra/refined`).
- Aligned `cm-honeybee` environment variables, dependencies, and volumes in `docker-compose.ui.yaml` with upstream `server/docker-compose.yaml` and `server/conf/cm-honeybee.yaml` (`HONEYBEE_VAULT_ADDR`, plus commented optional overrides `HONEYBEE_LISTEN_PORT`, `HONEYBEE_AGENT_PORT`, `HONEYBEE_AGENT_REQUEST_TIMEOUT`; removed obsolete `HONEYBEE_SPIDER_*`, `HONEYBEE_API_ALLOW_ORIGINS`, `HONEYBEE_LOG_LEVEL`, and `/app/db/` mount).
- Verified OpenBao configuration `cm-honeybee/openbao/openbao-config.hcl` remains identical with upstream v0.6.3.

## v0.6.1 Sync (2026-09-08)

- Bumped `cm-honeybee` service image to `cloudbaristaorg/cm-honeybee:0.6.1`.
- Removed fixed agent port mapping (`8082:8082`) in `docker-compose.ui.yaml`, as the agent now dynamically allocates and binds to a kernel-chosen loopback port over SSH.
- Verified OpenBao configuration `cm-honeybee/openbao/openbao-config.hcl` remains identical with upstream v0.6.1.
