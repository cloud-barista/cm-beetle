# Computing Infrastructure Migration

This repository provides computing infrastructure migration features.
This is a sub-system on [Cloud-Barista platform](https://github.com/cloud-barista/docs)
and utilizes [CB-Tumblebug](https://github.com/cloud-barista/cb-tumblebug)
to depoly a multi-cloud infra as a target computing infrastructure.

## Overview

Computing Infrastructure Migration framework (codename: cm-beetle) supports:

- recommendation of optimal target cloud infrastructure (compute, GPU, network, Kubernetes, and managed middleware),
- end-to-end infrastructure migration execution and sub-request tracking via CB-Tumblebug and CB-Spider, and
- cross-cloud filesystem and object storage data transfer (`transx`) with symmetric pre/post workload hooks.

### Architecture and Multi-Cloud Workflow

CM-Beetle operates on standardized `imdl` (`cm-model`) artifacts (`OnpremInfra`, `RecommendedInfra`), decoupling source infrastructure models from target CSP specifics:

```
Source Infra (On-Premise or Cloud)
   │  CM-Honeybee discovery -> Source / Planning Model (imdl)
   ▼
CM-Beetle Recommendation & Validation
   │  Target Blueprint (RecommendedInfra + cm-* lineage labels)
   ▼
CM-Beetle Migration & Data Transfer (transx)
   │  Transport-AOP sub-request tracking -> CB-Tumblebug -> CB-Spider
   ▼
Target Cloud Infra (AWS, Azure, GCP, Alibaba, Tencent, IBM, NCP, NHN, KT, OpenStack)
```

Because deployed resources preserve source machine traceability (`cm-source-machine-ids` and `cm-*` lineage labels), users can also re-target an already migrated workload to another CSP (Cloud-to-Cloud transfer) using the same artifact pipeline.

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
  - [Request Tracking and Sub-Request Progress](docs/feature-guide/core-request-tracking.md)
  - [Migration Lineage and Resource Traceability](docs/feature-guide/core-lineage.md)
  - [Naming Rules (NameSeed)](docs/feature-guide/core-naming-rule.md)
  - [Validation](docs/feature-guide/core-validation.md)
- **Infrastructure**
  - [Computing Infrastructure](docs/feature-guide/infra-computing.md)
  - [GPU Infrastructure Recommendation & Migration](docs/feature-guide/gpu-infra-migration-api-guide.md)
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
