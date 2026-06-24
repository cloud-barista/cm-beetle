# Managed Middleware

This document covers support for managed middleware resources (NLB, DBMS, Object Storage) by target CSP.

## Overview

| Resource       | AWS | Azure | GCP | Alibaba | Tencent | IBM | (OpenStack) | NCP | (NHN) | (KT) |
| -------------- | :-: | :---: | :-: | :-----: | :-----: | :-: | :---------: | :-: | :---: | :--: |
| Object Storage | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |     ✅      | ✅  |  ✅   |  ✅  |
| NLB            | ✅  |  🚧   | 🚧  |   🚧    |   📅    | 🚧  |     📅      | 🚧  |  📅   |  📅  |
| DBMS           | 📅  |  📅   | 📅  |   📅    |   📅    | 📅  |     📅      | 📅  |  📅   |  📅  |

## Object Storage

See also: [Object Storage Migration Feature Guide](object-storage-migration-feature-guide.md)

| Feature       | AWS | Azure | GCP | Alibaba | Tencent | IBM | OpenStack | NCP | NHN | KT  |
| ------------- | :-: | :---: | :-: | :-----: | :-----: | :-: | :-------: | :-: | :-: | :-: |
| CORS          | ✅  |  ❌   | ✅  |   ✅    |   ✅    | ✅  |    ✅     | ❌  | ❌  | ✅  |
| Versioning    | ✅  |  ❌   | ✅  |   ✅    |   ✅    | ✅  |    ❌     | ❌  | ❌  | ✅  |
| Presigned URL | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |    ✅     | ✅  | ✅  | ✅  |

> - Recommendation and migration are supported for all listed CSPs.
> - ❌ Not configurable per bucket — available at the provider level only.

## NLB (Preview)

> NLB recommendation and migration are available as a **preview** feature. Tested on AWS.

| Feature        | AWS | Azure | GCP | Alibaba | Tencent | IBM | (OpenStack) | NCP | (NHN) | (KT) |
| -------------- | :-: | :---: | :-: | :-----: | :-----: | :-: | :---------: | :-: | :---: | :--: |
| Recommendation | ✅  |  🚧   | 🚧  |   🚧    |   📅    | 🚧  |     📅      | 🚧  |  📅   |  📅  |
| Migration      | ✅  |  🚧   | 🚧  |   🚧    |   📅    | 🚧  |     📅      | 🚧  |  📅   |  📅  |

> - Recommendation groups source nodes by NLB backend topology (N:1) and returns Pareto-frontier spec candidates.
> - Migration provisions NLBs on the target cloud after VM infrastructure is deployed.

## Roadmap

- 🚧 **In progress**: NLB — Azure, GCP, Alibaba, IBM, NCP
- 📅 **Planned**: NLB — Tencent, OpenStack, NHN, KT · DBMS — all CSPs
