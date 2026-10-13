# Target Cloud Infrastructure Summary

**Generated At:** 2026-10-12 11:02:10

**Namespace:** mig01

**Infra Name:** my08-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my08-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | NCP |
| **Target Region** | KR |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| ci2-g3 | 2 | 4.0 | - | x86_64 | default | $0.0730 | 1 |
| s4-g3 | 4 | 16.0 | - | x86_64 | default | $0.1747 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| 23214590 | ubuntu-22.04-base (Hypervisor:KVM) | Ubuntu 22.04 | Linux/UNIX | x86_64 | Common BlockStorage 1 | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | 146264585 | Running | 2 vCPU, 4.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 211.233.199.16<br>**Private IP:** 10.0.1.6<br>**SGs:** my08-mig-sg-02<br>**SSH:** my08-mig-sshkey-01 |
| my08-ng-influxdb-back-1 | 146264588 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 211.233.204.234<br>**Private IP:** 10.0.1.8<br>**SGs:** my08-mig-sg-01<br>**SSH:** my08-mig-sshkey-01 |
| my08-ng-influxdb-back-2 | 146264582 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 211.233.199.17<br>**Private IP:** 10.0.1.7<br>**SGs:** my08-mig-sg-01<br>**SSH:** my08-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my08-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my08-mig-vnet-01 |
| **CSP VNet ID** | 150301 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ncp-kr |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my08-mig-subnet-01 | 329426 | 10.0.1.0/24 | KR-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my08-mig-sshkey-01 | tbbrmposh4nl13mhejhr | cb-user | a6:34:d8:d8:ae:41:ef:2d:44:d6:5d:e8:df:2a:9a:4e |

### Security Groups

#### Security Group: my08-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my08-mig-sg-01 |
| **CSP Security Group ID** | 406332 |
| **VNet** | my08-mig-vnet-01 |
| **Rule Count** | 10 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my08-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my08-mig-sg-02 |
| **CSP Security Group ID** | 406333 |
| **VNet** | my08-mig-vnet-01 |
| **Rule Count** | 8 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.4224 |
| **Per Day** | $10.14 |
| **Per Month (30 days)** | $304.13 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| NCP | KR | 3 | $0.4224 | $304.13 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | ci2-g3 | $0.0730 | $52.56 |
| my08-ng-influxdb-back-1 | s4-g3 | $0.1747 | $125.78 |
| my08-ng-influxdb-back-2 | s4-g3 | $0.1747 | $125.78 |


