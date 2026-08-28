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

## v0.6.0 Sync (2026-08-20)

- Initialized `cm-honeybee/openbao/openbao-config.hcl` based on upstream `cm-honeybee` v0.6.0.
- `cm-honeybee` self-manages its dedicated OpenBao container (`openbao-honeybee`) by performing automated initialization and unsealing on startup.
