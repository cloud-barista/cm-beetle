# Target Cloud Infrastructure Summary

**Generated At:** 2026-01-05 00:55:03

**Namespace:** mig01

**MCI Name:** mmci01

---

## Overview

| Property | Value |
|----------|-------|
| **MCI Name** | mmci01 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | ALIBABA |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| ecs.t6-c1m4.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1693 | 1 |
| ecs.t6-c1m1.large | 2 | 2.0 | - | x86_64 |  | $0.0214 | 1 |
| ecs.t6-c1m4.large | 2 | 8.0 | - | x86_64 |  | $0.0846 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ubuntu_22_04_x64_20G_alibase_20251126.vhd | Ubuntu  22.04 64 bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| migrated-ec268ed7-821e-9d73-e79f-961262161624-1 | i-mj7hkcwwo3sw5o0yj4ez | Running | 2 vCPU, 2.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 8.220.214.65<br>**Private IP:** 10.0.1.20<br>**SGs:** mig-sg-01<br>**SSH:** mig-sshkey-01 |
| migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | i-mj75o04wa4yj53jao06z | Running | 2 vCPU, 8.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 47.80.8.142<br>**Private IP:** 10.0.1.19<br>**SGs:** mig-sg-03<br>**SSH:** mig-sshkey-01 |
| migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | i-mj7b0w9cdl65cen5k7qd | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 8.220.217.145<br>**Private IP:** 10.0.1.18<br>**SGs:** mig-sg-02<br>**SSH:** mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | mig-vnet-01 |
| **CSP VNet ID** | vpc-mj749b5u3pvfmthms4ebu |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | alibaba-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| mig-subnet-01 | vsw-mj7mswvsxitpmjaknyemm | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| mig-sshkey-01 | d5dgn3q5npi2mb8i9jd0 |  | b766065efcbd62b8e8a4654598aee662 |

### Security Groups

#### Security Group: mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | mig-sg-01 |
| **CSP Security Group ID** | sg-mj7bnzgqyrbe328p0jyu |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 14 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
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
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | mig-sg-02 |
| **CSP Security Group ID** | sg-mj7i9ov4dzvjxv6sglta |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
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
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: mig-sg-03

| Property | Value |
|----------|-------|
| **Name** | mig-sg-03 |
| **CSP Security Group ID** | sg-mj7f0gosmwx4215asuai |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
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
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.2753 |
| **Per Day** | $6.61 |
| **Per Month (30 days)** | $198.20 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| ALIBABA | ap-northeast-2 | 3 | $0.2753 | $198.20 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| migrated-ec268ed7-821e-9d73-e79f-961262161624-1 | ecs.t6-c1m1.large | $0.0214 | $15.40 |
| migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | ecs.t6-c1m4.large | $0.0846 | $60.93 |
| migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | ecs.t6-c1m4.xlarge | $0.1693 | $121.87 |


