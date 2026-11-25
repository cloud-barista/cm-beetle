# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:22:43

**Namespace:** mig01

**Infra Name:** my02-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my02-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | AZURE |
| **Target Region** | koreasouth |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| Standard_D4_v4 | 4 | 16.0 | - | x86_64 |  | $0.2210 | 2 |
| Standard_F2s_v2 | 2 | 4.0 | - | x86_64 |  | $0.0961 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my02-ng-influxdb-back-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb2mefhc7qac6sg6gmoh | Running | 4 vCPU, 16.0 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 52.147.125.128<br>**Private IP:** 10.0.1.6<br>**SGs:** my02-mig-sg-01<br>**SSH:** my02-mig-sshkey-01 |
| my02-ng-influxdb-back-2 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb6lhvd6gifs72gf64ok | Running | 4 vCPU, 16.0 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 74.227.129.35<br>**Private IP:** 10.0.1.4<br>**SGs:** my02-mig-sg-01<br>**SSH:** my02-mig-sshkey-01 |
| my02-ng-ip-10-0-1-30-62161624-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd32ihs5kcb30u4ok9r | Running | 2 vCPU, 4.0 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 74.227.128.223<br>**Private IP:** 10.0.1.5<br>**SGs:** my02-mig-sg-02<br>**SSH:** my02-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my02-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my02-mig-vnet-01 |
| **CSP VNet ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | azure-koreasouth |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my02-mig-subnet-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm | 10.0.1.0/24 |  |

### Network Load Balancers (NLB)

#### NLB: my02-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my02-ng-influxdb-back |
| **CSP NLB Name** | tbbjvscnut4iq5kjrmj4 |
| **CSP NLB ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4 |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | azure-koreasouth |
| **Listener** | TCP:9999 (IP: 20.200.190.98) |
| **Target Group** | TCP:8086 (Node Group: my02-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: CSP Default, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my02-ng-influxdb-back-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb2mefhc7qac6sg6gmoh | 10.0.1.6 | 52.147.125.128 | Running |
| my02-ng-influxdb-back-2 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb6lhvd6gifs72gf64ok | 10.0.1.4 | 74.227.129.35 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my02-mig-sshkey-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg |  |  |

### Security Groups

#### Security Group: my02-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my02-mig-sg-01 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tb77r547caqdgeolc338 |
| **VNet** | my02-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my02-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my02-mig-sg-02 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd34sc0anir1ka300ki |
| **VNet** | my02-mig-vnet-01 |
| **Rule Count** | 4 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.5381 |
| **Per Day** | $12.91 |
| **Per Month (30 days)** | $387.43 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AZURE | koreasouth | 3 | $0.5381 | $387.43 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my02-ng-influxdb-back-1 | Standard_D4_v4 | $0.2210 | $159.12 |
| my02-ng-influxdb-back-2 | Standard_D4_v4 | $0.2210 | $159.12 |
| my02-ng-ip-10-0-1-30-62161624-1 | Standard_F2s_v2 | $0.0961 | $69.19 |


