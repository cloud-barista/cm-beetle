# Target Cloud Infrastructure Summary

**Generated At:** 2026-10-30 12:51:45

**Namespace:** mig01

**Infra Name:** my05-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my05-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | TENCENT |
| **Target Region** | ap-seoul |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| S6.LARGE16 | 4 | 16.0 | - | x86_64 |  | $0.2300 | 2 |
| SA5.MEDIUM2 | 2 | 2.0 | - | x86_64 |  | $0.0600 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| img-487zeit5 | Ubuntu Server 22.04 LTS 64bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 1 |
| img-7rotv4ux | Ubuntu Server 22.04 LTS 64bit UEFI | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 2 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | ins-mhkac4j9 | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04 LTS 64bit (Ubuntu Server 22.04 LTS 64bit) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.166.11.218<br>**Private IP:** 10.0.1.8<br>**SGs:** my05-mig-sg-02<br>**SSH:** my05-mig-sshkey-01 |
| my05-ng-influxdb-back-1 | ins-4ap62gvz | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.131.254.242<br>**Private IP:** 10.0.1.10<br>**SGs:** my05-mig-sg-01<br>**SSH:** my05-mig-sshkey-01 |
| my05-ng-influxdb-back-2 | ins-q40emar7 | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.166.9.19<br>**Private IP:** 10.0.1.2<br>**SGs:** my05-mig-sg-01<br>**SSH:** my05-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my05-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my05-mig-vnet-01 |
| **CSP VNet ID** | vpc-rtznlhxm |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | tencent-ap-seoul |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my05-mig-subnet-01 | subnet-19h6xb1b | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my05-mig-sshkey-01 | skey-7ic7ibdp |  |  |

### Security Groups

#### Security Group: my05-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my05-mig-sg-01 |
| **CSP Security Group ID** | sg-d2jm3yvx |
| **VNet** | my05-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my05-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my05-mig-sg-02 |
| **CSP Security Group ID** | sg-hgf81iwn |
| **VNet** | my05-mig-vnet-01 |
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
| **Per Hour** | $0.5200 |
| **Per Day** | $12.48 |
| **Per Month (30 days)** | $374.40 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| TENCENT | ap-seoul | 3 | $0.5200 | $374.40 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | SA5.MEDIUM2 | $0.0600 | $43.20 |
| my05-ng-influxdb-back-1 | S6.LARGE16 | $0.2300 | $165.60 |
| my05-ng-influxdb-back-2 | S6.LARGE16 | $0.2300 | $165.60 |


