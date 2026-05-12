# Target Cloud Infrastructure Summary

**Generated At:** 2026-05-12 09:52:07

**Namespace:** mig01

**Infra Name:** my04-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my04-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
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
| s2-g3a | 2 | 8.0 | - | x86_64 | default | $0.0848 | 1 |
| s4-g3a | 4 | 16.0 | - | x86_64 | default | $0.1747 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| 23214590 | ubuntu-22.04-base (Hypervisor:KVM) | Ubuntu 22.04 | Linux/UNIX | x86_64 | Common BlockStorage 1 | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | 138234481 | Running | 2 vCPU, 4.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my04-vnet-01<br>**Subnet:** my04-subnet-01<br>**Public IP:** 101.79.19.173<br>**Private IP:** 10.0.1.8<br>**SGs:** my04-sg-01<br>**SSH:** my04-sshkey-01 |
| my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | 138234474 | Running | 2 vCPU, 8.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my04-vnet-01<br>**Subnet:** my04-subnet-01<br>**Public IP:** 101.79.20.23<br>**Private IP:** 10.0.1.7<br>**SGs:** my04-sg-03<br>**SSH:** my04-sshkey-01 |
| my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | 138234467 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my04-vnet-01<br>**Subnet:** my04-subnet-01<br>**Public IP:** 101.79.19.56<br>**Private IP:** 10.0.1.6<br>**SGs:** my04-sg-02<br>**SSH:** my04-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my04-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my04-vnet-01 |
| **CSP VNet ID** | 138792 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ncp-kr |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my04-subnet-01 | 300131 | 10.0.1.0/24 | KR-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my04-sshkey-01 | tbd81fd3e6aji65qtg7seg | cb-user |  |

### Security Groups

#### Security Group: my04-sg-01

| Property | Value |
|----------|-------|
| **Name** | my04-sg-01 |
| **CSP Security Group ID** | 352888 |
| **VNet** | my04-vnet-01 |
| **Rule Count** | 15 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| inbound | UDP | 9113 | 10.0.0.0/16 |
| inbound | TCP | 9113 | 10.0.0.0/16 |
| inbound | TCP | 8080 | 0.0.0.0/0 |
| inbound | TCP | 443 | 0.0.0.0/0 |
| inbound | TCP | 80 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my04-sg-02

| Property | Value |
|----------|-------|
| **Name** | my04-sg-02 |
| **CSP Security Group ID** | 352889 |
| **VNet** | my04-vnet-01 |
| **Rule Count** | 20 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| inbound | UDP | 9100 | 10.0.0.0/16 |
| inbound | TCP | 9100 | 10.0.0.0/16 |
| inbound | UDP | 32803 | 10.0.0.0/16 |
| inbound | TCP | 32803 | 10.0.0.0/16 |
| inbound | UDP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 20048 | 10.0.0.0/16 |
| inbound | UDP | 111 | 0.0.0.0/0 |
| inbound | TCP | 111 | 0.0.0.0/0 |
| inbound | UDP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my04-sg-03

| Property | Value |
|----------|-------|
| **Name** | my04-sg-03 |
| **CSP Security Group ID** | 352890 |
| **VNet** | my04-vnet-01 |
| **Rule Count** | 20 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| inbound | UDP | 9104 | 10.0.0.0/16 |
| inbound | TCP | 9104 | 10.0.0.0/16 |
| inbound | UDP | 4444 | 10.0.0.0/16 |
| inbound | TCP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 4568 | 10.0.0.0/16 |
| inbound | TCP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4567 | 10.0.0.0/16 |
| inbound | TCP | 4567 | 10.0.0.0/16 |
| inbound | UDP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.3325 |
| **Per Day** | $7.98 |
| **Per Month (30 days)** | $239.40 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| NCP | KR | 3 | $0.3325 | $239.40 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | ci2-g3 | $0.0730 | $52.56 |
| my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | s2-g3a | $0.0848 | $61.06 |
| my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | s4-g3a | $0.1747 | $125.78 |


