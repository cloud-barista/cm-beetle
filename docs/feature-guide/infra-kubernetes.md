# Kubernetes Infrastructure

This document covers recommendation and migration support for Kubernetes infrastructure resources by target CSP.

## Support Status

| Resource       | AWS | Azure | GCP | Alibaba | Tencent | IBM | (OpenStack) | NCP | NHN | (KT) |
| -------------- | :-: | :---: | :-: | :-----: | :-----: | :-: | :---------: | :-: | :-: | :--: |
| VNet           | 📅  |  📅   | 📅  |   📅    |   📅    | 📅  |     ❌      |  -  |  -  |  ❌  |
| Subnet         | 📅  |  📅   | 📅  |   📅    |   📅    | 📅  |     ❌      |  -  |  -  |  ❌  |
| Security Group | 📅  |  📅   | 📅  |   📅    |   📅    | 📅  |     ❌      |  -  |  -  |  ❌  |
| SSH Key        | 📅  |  📅   | 📅  |   📅    |   📅    | 📅  |     ❌      |  -  |  -  |  ❌  |
| VM Spec        | 🚧  |  📅   | 📅  |   📅    |   📅    | 📅  |     ❌      |  -  |  -  |  ❌  |
| Node Image     | ✅  |  ✅   | ✅  |   ✅    |   🚧    | ✅  |     ❌      | ✅  | ✅  |  ❌  |
| K8s Cluster    | ✅  |  ✅   | ✅  |   ✅    |   🚧    | ✅  |     ❌      | ✅  | ✅  |  ❌  |

## Supported Operations

TBD

## Roadmap

- 🚧 **In progress**: Tencent
- ❌ OpenStack, KT: Not supported
