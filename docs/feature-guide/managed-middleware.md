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

> NLB recommendation and migration are available as a **preview** feature. Tested on AWS, Azure, GCP, and IBM.

| Feature        | AWS | Azure | GCP | Alibaba | Tencent | IBM | (OpenStack) | NCP | (NHN) | (KT) |
| -------------- | :-: | :---: | :-: | :-----: | :-----: | :-: | :---------: | :-: | :---: | :--: |
| Recommendation | ✅  |  🚧   | 🚧  |   🚧    |   📅    | 🚧  |     📅      | 🚧  |  📅   |  📅  |
| Migration      | ✅  |  🚧   | 🚧  |   🚧    |   📅    | 🚧  |     📅      | 🚧  |  📅   |  📅  |

> - Recommendation groups source nodes by NLB backend topology (N:1) and returns Pareto-frontier spec candidates.
> - Migration provisions NLBs on the target cloud after VM infrastructure is deployed.

### CSP-Specific Characteristics and Constraints

| CSP   | Port translation | Endpoint        | Key constraint / auto-adjustment |
|-------|:---------------:|-----------------|----------------------------------|
| AWS   | ✅              | DNS name        | SG: backend port opened from `0.0.0.0/0` automatically. DNS takes ~5 min to propagate. |
| Azure | ✅              | DNS + static IP | Health check timeout not supported — omitted automatically. |
| GCP   | ❌              | IP only         | Listener port forced = backend (application) port. Clients connect on the application port (e.g., 8086, not 9999). |
| IBM   | ✅              | IP (async)      | Listener address assigned asynchronously — re-query if empty after migration. Timeout < interval enforced automatically. |

> **GCP**: External Passthrough NLB uses target pools with no port translation. Traffic arrives at backend VMs on the listener port as-is, so the recommendation overrides the listener port to match the application port.
>
> **IBM**: The NLB listener address may be empty immediately after creation due to asynchronous provisioning. Retry the GET NLB call after a short wait.

### NLB Deletion Behavior

- Some CSPs delete NLBs asynchronously — the API returns success before ENIs are fully released.
- Deleting VNet/subnets too soon may cause dependency errors (e.g., `DependencyViolation` on AWS).
- CM-Beetle mitigations:
  - `DeleteNlb`: waits after a successful deletion response before returning (e.g., 15s).
  - `DeleteVMInfra` (VNet deletion): retries on failure with a back-off interval (e.g., 10s × up to 10 times).

## Roadmap

- 🚧 **In progress**: NLB — Azure, GCP, Alibaba, IBM, NCP
- 📅 **Planned**: NLB — Tencent, OpenStack, NHN, KT · DBMS — all CSPs
