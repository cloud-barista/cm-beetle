# Target Cloud Infrastructure Summary

**Generated At:** 2026-01-23 06:01:38

**Namespace:** mig01

**MCI Name:** mmci01

---

## Overview

| Property | Value |
|----------|-------|
| **MCI Name** | mmci01 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | AWS |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| t3a.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1872 | 1 |
| t3a.small | 2 | 2.0 | - | x86_64 |  | $0.0234 | 1 |
| t3a.large | 2 | 8.0 | - | x86_64 |  | $0.0936 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ami-068550d7714642b14 | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251122 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| migrated-ec268ed7-821e-9d73-e79f-961262161624-1 | i-071732b058553f34a | Running | 2 vCPU, 2.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251122 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251122) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 43.201.5.115<br>**Private IP:** 10.0.1.244<br>**SGs:** mig-sg-01<br>**SSH:** mig-sshkey-01 |
| migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | i-0c19cca3d0c1b884b | Running | 2 vCPU, 8.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251122 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251122) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 3.38.96.26<br>**Private IP:** 10.0.1.231<br>**SGs:** mig-sg-03<br>**SSH:** mig-sshkey-01 |
| migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | i-0334fe36c78bec854 | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251122 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251122) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 3.34.200.178<br>**Private IP:** 10.0.1.61<br>**SGs:** mig-sg-02<br>**SSH:** mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | mig-vnet-01 |
| **CSP VNet ID** | vpc-07a606b9a0d6e3e28 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| mig-subnet-01 | subnet-0267320d200451643 | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| mig-sshkey-01 | d5pgssl5q76kouk46s3g |  | a3:fd:75:98:41:17:97:43:8d:4a:ca:8e:7a:cf:11:7f:e8:7d:7f:65 |

### Security Groups

#### Security Group: mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | mig-sg-01 |
| **CSP Security Group ID** | sg-06a92d28f61a34ab3 |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 14 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 80 | 0.0.0.0/0 |
| inbound | TCP | 8080 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | UDP | 9113 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 9113 | 10.0.0.0/16 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | TCP | 443 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | mig-sg-02 |
| **CSP Security Group ID** | sg-052e086c77bfb884c |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | UDP | 9100 | 10.0.0.0/16 |
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 32803 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 111 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 32803 | 10.0.0.0/16 |
| inbound | TCP | 9100 | 10.0.0.0/16 |
| inbound | UDP | 111 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: mig-sg-03

| Property | Value |
|----------|-------|
| **Name** | mig-sg-03 |
| **CSP Security Group ID** | sg-0d5f2fa9080967bd4 |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | UDP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 9104 | 10.0.0.0/16 |
| inbound | TCP | 9104 | 10.0.0.0/16 |
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 3306 | 10.0.0.0/16 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4567 | 10.0.0.0/16 |
| inbound | TCP | 4567 | 10.0.0.0/16 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | TCP | 3306 | 10.0.0.0/16 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.3042 |
| **Per Day** | $7.30 |
| **Per Month (30 days)** | $219.02 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AWS | ap-northeast-2 | 3 | $0.3042 | $219.02 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| migrated-ec268ed7-821e-9d73-e79f-961262161624-1 | t3a.small | $0.0234 | $16.85 |
| migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | t3a.large | $0.0936 | $67.39 |
| migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | t3a.xlarge | $0.1872 | $134.78 |


