# VM Infrastructure

This document covers recommendation and migration support for VM infrastructure resources by target CSP.

## Support Status

| Resource        | AWS | Azure | GCP | Alibaba | Tencent | IBM | (OpenStack) | NCP | (NHN) | (KT) |
| --------------- | :-: | :---: | :-: | :-----: | :-----: | :-: | :---------: | :-: | :---: | :--: |
| VNet            | ✅  |  ✅   | ✅  |   ✅    |   📅    | 🚧  |     📅      | ✅  |   —   |  —   |
| Subnet          | ✅  |  ✅   | ✅  |   ✅    |   📅    | 🚧  |     📅      | ✅  |   —   |  —   |
| Security Group  | ✅  |  ✅   | ✅  |   ✅    |   📅    | 🚧  |     📅      | ✅  |   —   |  —   |
| SSH Key         | ✅  |  ✅   | ✅  |   ✅    |   📅    | 🚧  |     📅      | ✅  |   —   |  —   |
| VM Spec         | ✅  |  ✅   | ✅  |   ✅    |   📅    | 🚧  |     📅      | ✅  |   —   |  —   |
| OS Image        | ✅  |  ✅   | ✅  |   ✅    |   📅    | 🚧  |     📅      | ✅  |   —   |  —   |
| Virtual Machine | ✅  |  ✅   | ✅  |   ⚠️    |   📅    | 🚧  |     📅      | ✅  |   —   |  —   |

## Supported Operations

| Operation      | Supported Resources                                              |
| -------------- | ---------------------------------------------------------------- |
| Recommendation | VNet, Subnet, Security Group, VM Spec, OS Image, Virtual Machine |
| Migration      | VNet, Subnet, Security Group, SSH Key, Virtual Machine           |

## Roadmap

- ⚠️ **Alibaba VM**: supported with known issues/limitations
- 🚧 **In progress**: IBM — all resources
- 📅 **Planned**: Tencent, OpenStack
