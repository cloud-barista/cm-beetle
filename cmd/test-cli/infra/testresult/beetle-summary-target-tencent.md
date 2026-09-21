# Target Cloud Infrastructure Summary

**Generated At:** 2026-09-21 04:21:08

**Namespace:** mig01

**Infra Name:** my05-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my05-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
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
| BF1.MEDIUM2 | 2 | 2.0 | - | x86_64 |  | $0.0200 | 1 |
| BF1.MEDIUM8 | 2 | 8.0 | - | x86_64 |  | $0.0700 | 1 |
| BF1.LARGE16 | 4 | 16.0 | - | x86_64 |  | $0.1500 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| img-7rotv4ux | Ubuntu Server 22.04 LTS 64bit UEFI | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | ins-pjjkbdwz | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-vnet-01<br>**Subnet:** my05-subnet-01<br>**Public IP:** 43.155.208.156<br>**Private IP:** 10.0.1.13<br>**SGs:** my05-sg-01<br>**SSH:** my05-sshkey-01 |
| my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | ins-pl7luz01 | Running | 2 vCPU, 8.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-vnet-01<br>**Subnet:** my05-subnet-01<br>**Public IP:** 43.155.222.67<br>**Private IP:** 10.0.1.3<br>**SGs:** my05-sg-03<br>**SSH:** my05-sshkey-01 |
| my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | ins-olowx4jd | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-vnet-01<br>**Subnet:** my05-subnet-01<br>**Public IP:** 119.28.150.173<br>**Private IP:** 10.0.1.12<br>**SGs:** my05-sg-02<br>**SSH:** my05-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my05-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my05-vnet-01 |
| **CSP VNet ID** | vpc-ktohe25c |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | tencent-ap-seoul |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my05-subnet-01 | subnet-0ilffklb | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my05-sshkey-01 | skey-c8que8al |  |  |

### Security Groups

#### Security Group: my05-sg-01

| Property | Value |
|----------|-------|
| **Name** | my05-sg-01 |
| **CSP Security Group ID** | sg-m9z8xo0l |
| **VNet** | my05-vnet-01 |
| **Rule Count** | 14 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 80 | 0.0.0.0/0 |
| inbound | TCP | 443 | 0.0.0.0/0 |
| inbound | TCP | 8080 | 0.0.0.0/0 |
| inbound | TCP | 9113 | 10.0.0.0/16 |
| inbound | UDP | 9113 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my05-sg-02

| Property | Value |
|----------|-------|
| **Name** | my05-sg-02 |
| **CSP Security Group ID** | sg-3z1hp8d1 |
| **VNet** | my05-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 2049 | 0.0.0.0/0 |
| inbound | UDP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 111 | 0.0.0.0/0 |
| inbound | UDP | 111 | 0.0.0.0/0 |
| inbound | TCP | 20048 | 10.0.0.0/16 |
| inbound | UDP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 32803 | 10.0.0.0/16 |
| inbound | UDP | 32803 | 10.0.0.0/16 |
| inbound | TCP | 9100 | 10.0.0.0/16 |
| inbound | UDP | 9100 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my05-sg-03

| Property | Value |
|----------|-------|
| **Name** | my05-sg-03 |
| **CSP Security Group ID** | sg-4au7mzat |
| **VNet** | my05-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 3306 | 10.0.0.0/16 |
| inbound | UDP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 4567 | 10.0.0.0/16 |
| inbound | UDP | 4567 | 10.0.0.0/16 |
| inbound | TCP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4568 | 10.0.0.0/16 |
| inbound | TCP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 4444 | 10.0.0.0/16 |
| inbound | TCP | 9104 | 10.0.0.0/16 |
| inbound | UDP | 9104 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.2400 |
| **Per Day** | $5.76 |
| **Per Month (30 days)** | $172.80 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| TENCENT | ap-seoul | 3 | $0.2400 | $172.80 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | BF1.MEDIUM2 | $0.0200 | $14.40 |
| my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | BF1.MEDIUM8 | $0.0700 | $50.40 |
| my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | BF1.LARGE16 | $0.1500 | $108.00 |


