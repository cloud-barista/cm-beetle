# Target Cloud Infrastructure Summary

**Generated At:** 2026-10-30 12:57:28

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
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| Standard_F2s_v2 | 2 | 3.9 | - | x86_64 |  | $0.0961 | 1 |
| Standard_D4s_v4 | 4 | 15.6 | - | x86_64 |  | $0.2210 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060 | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqv3mailnhp49dco6b2 | Running | 2 vCPU, 3.9 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 20.214.57.24<br>**Private IP:** 10.0.1.5<br>**SGs:** my02-mig-sg-02<br>**SSH:** my02-mig-sshkey-01 |
| my02-ng-influxdb-back-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbpdi0ugkueg94k19mdc | Running | 4 vCPU, 15.6 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 20.200.178.85<br>**Private IP:** 10.0.1.4<br>**SGs:** my02-mig-sg-01<br>**SSH:** my02-mig-sshkey-01 |
| my02-ng-influxdb-back-2 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbnhrp3lhf257rnbhamg | Running | 4 vCPU, 15.6 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 40.89.195.136<br>**Private IP:** 10.0.1.6<br>**SGs:** my02-mig-sg-01<br>**SSH:** my02-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my02-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my02-mig-vnet-01 |
| **CSP VNet ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | azure-koreasouth |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my02-mig-subnet-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83 | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my02-mig-sshkey-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6 |  |  |

### Security Groups

#### Security Group: my02-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my02-mig-sg-01 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbriet3sm4j9st9l3ls7 |
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
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbn462dpau0gr390pvuu |
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
| my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | Standard_F2s_v2 | $0.0961 | $69.19 |
| my02-ng-influxdb-back-1 | Standard_D4s_v4 | $0.2210 | $159.12 |
| my02-ng-influxdb-back-2 | Standard_D4s_v4 | $0.2210 | $159.12 |


