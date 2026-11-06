# Computing Infrastructure Migration

This repository provides computing infrastructure migration features.
This is a sub-system on [Cloud-Barista platform](https://github.com/cloud-barista/docs)
and utilizes [CB-Tumblebug](https://github.com/cloud-barista/cb-tumblebug)
to depoly a multi-cloud infra as a target computing infrastructure.

## Overview

Computing Infrastructure Migration framework (codename: cm-beetle) is going to support:

- migration execution and control from source to target computing infrastructure, and
- recommendation of optimal configuration of target cloud infrastructure.

## Migration Support Status Summary

The tables below show recommendation and migration support by target CSP and resource type.

<details>
<summary>Legend</summary>

| Icon | Meaning         |
| :--: | --------------- |
|  ✅  | Fully supported |
|  ⚠️  | Has limitations |
|  🚧  | In progress     |
|  📅  | Planned         |
|  ❌  | Not available   |
|  —   | Not planned     |

> **(CSP)** — Low priority or not pursued this year.

</details>

| Resource        | AWS | Azure | GCP | Alibaba | Tencent | IBM | OpenStack | NCP | NHN | KT  |
| --------------- | :-: | :---: | :-: | :-----: | :-----: | :-: | :-------: | :-: | :-: | :-: |
| VNet            | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    📅     | ✅  | ✅  |  —  |
| Subnet          | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    📅     | ✅  | ✅  |  —  |
| Security Group  | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    📅     | ✅  | ✅  |  —  |
| SSH Key         | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    📅     | ✅  | ✅  |  —  |
| Node Spec       | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    📅     | ✅  | ✅  |  —  |
| OS Image        | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    📅     | ✅  | ✅  |  —  |
| Virtual Machine | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    📅     | ✅  | ✅  |  —  |
| K8s Cluster     | ✅  |  ✅   | ✅  |   ✅    |   🚧    | ✅  |    ❌     | ✅  | ✅  | ❌  |
| K8s Node Image  | ✅  |  ✅   | ✅  |   ✅    |   🚧    | ✅  |    ❌     | ✅  | ✅  | ❌  |
| Object Storage  | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    ✅     | ✅  | ✅  | ✅  |
| NLB             | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    📅     | ✅  | ✅  | 📅  |
| RDBMS (MySQL)   | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    ✅     | ✅  | ✅  | ❌  |
| RDBMS (MariaDB) | ✅  |  ❌   | ❌  |   ✅    |   ❌    | ❌  |    ✅     | ❌  | ✅  | ❌  |

**Data Migration · Incubating**

Data migration transfers data between storage endpoints regardless of which CSP hosts them.

| Source                         | Destination                    | Supported |
| ------------------------------ | ------------------------------ | :-------: |
| Filesystem (SSH/rsync)         | Filesystem (SSH/rsync)         |    ✅     |
| Filesystem (SSH/rsync)         | Object Storage (S3-compatible) |    ✅     |
| Object Storage (S3-compatible) | Object Storage (S3-compatible) |    ✅     |
| Object Storage (S3-compatible) | Filesystem (SSH/rsync)         |    ✅     |

### Feature Guides

- **Core & System**
  - [Async Responses](docs/feature-guide/core-async-response.md)
  - [Naming Rules (NameSeed)](docs/feature-guide/core-naming-rule.md)
  - [Validation](docs/feature-guide/core-validation.md)
- **Infrastructure**
  - [Computing Infrastructure](docs/feature-guide/infra-computing.md)
  - [Kubernetes Infrastructure](docs/feature-guide/infra-kubernetes.md)
- **Managed Middleware**
  - [Managed Middleware Overview](docs/feature-guide/middleware-overview.md)
  - [Network Load Balancer (NLB)](docs/feature-guide/middleware-nlb.md)
  - [Managed Object Storage](docs/feature-guide/middleware-object-storage.md)
  - [Managed RDBMS (RDS)](docs/feature-guide/middleware-rdbms.md)
- **Data**
  - [Data Transfer](docs/feature-guide/data-transfer.md)

## Execution and development environment

- Operating system (OS):
  - Ubuntu 22.04
- Languages:
  - Go: 1.26.2+
  - Python: 3.10.12
- Container runtime:
  - Docker: 24.0.7

## How to run and use CM-Beetle

1. [Installation and execution guide](docs/installation-and-execution.md)
2. [How to deploy CM-Beetle, along with its related components](https://github.com/cloud-barista/cm-beetle/discussions/105)

Note - Deprecated, [How to run and use the CM-Beetle](https://github.com/cloud-barista/cm-beetle/discussions/73)
