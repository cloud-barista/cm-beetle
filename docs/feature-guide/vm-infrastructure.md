# VM Infrastructure

This document covers recommendation and migration support for VM infrastructure resources by target CSP.

## Support Status

| Resource        | AWS | Azure | GCP | Alibaba | Tencent | IBM | (OpenStack) | NCP | (NHN) | (KT) |
| --------------- | :-: | :---: | :-: | :-----: | :-----: | :-: | :---------: | :-: | :---: | :--: |
| VNet            | ✅  |  ✅   | ✅  |   ✅    |   📅    | ✅  |     📅      | ✅  |   —   |  —   |
| Subnet          | ✅  |  ✅   | ✅  |   ✅    |   📅    | ✅  |     📅      | ✅  |   —   |  —   |
| Security Group  | ✅  |  ✅   | ✅  |   ✅    |   📅    | ✅  |     📅      | ✅  |   —   |  —   |
| SSH Key         | ✅  |  ✅   | ✅  |   ✅    |   📅    | ✅  |     📅      | ✅  |   —   |  —   |
| Node Spec       | ✅  |  ✅   | ✅  |   ⚠️    |   📅    | ✅  |     📅      | ✅  |   —   |  —   |
| OS Image        | ✅  |  ✅   | ✅  |   ✅    |   📅    | ✅  |     📅      | ✅  |   —   |  —   |
| Virtual Machine | ✅  |  ✅   | ✅  |   ✅    |   📅    | ✅  |     📅      | ✅  |   —   |  —   |

## Supported Operations

| Operation      | Supported Resources                                              |
| -------------- | ---------------------------------------------------------------- |
| Recommendation | VNet, Subnet, Security Group, VM Spec, OS Image, Virtual Machine |
| Migration      | VNet, Subnet, Security Group, SSH Key, Virtual Machine           |

## Roadmap

- **Planned**: Tencent, OpenStack
