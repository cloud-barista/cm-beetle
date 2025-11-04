# Target Cloud Infrastructure Summary

**Generated At:** 2025-11-04 15:13:01

**Namespace:** mig01

**MCI Name:** mmci01

---

## Overview

| Property | Value |
|----------|-------|
| **MCI Name** | mmci01 |
| **Description** | a recommended multi-cloud infrastructure |
| **Status** | Running:2 (R:2/2) |
| **Target Cloud** | AWS |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 2 |
| **Running VMs** | 2 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| c5a.2xlarge | 8 | 16.0 | - | x86_64 |  | $0.3440 | 1 |
| t3a.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1872 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ami-010be25c3775061c9 | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | - | 2 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1 | i-0af60931229c7ed22 | Running | 8 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 43.203.207.119<br>**Private IP:** 192.168.110.152<br>**SGs:** mig-sg-02<br>**SSH:** mig-sshkey-01 |
| migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1 | i-0f42c57ec9c0984b0 | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 3.35.18.54<br>**Private IP:** 192.168.110.128<br>**SGs:** mig-sg-01<br>**SSH:** mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | mig-vnet-01 |
| **CSP VNet ID** | vpc-02eba736ffc6c3057 |
| **CIDR Block** | 192.168.96.0/19 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| mig-subnet-01 | subnet-0f185c6541a811d62 | 192.168.110.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| mig-sshkey-01 | d451fbmqjs728pra18vg |  | 05:24:1b:8e:6c:18:45:9c:1d:5c:a5:f9:7b:ce:18:cd:76:ed:d9:27 |

### Security Groups

#### Security Group: mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | mig-sg-01 |
| **CSP Security Group ID** | sg-0a751c7bb2497f7b3 |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 10022 | 0.0.0.0/0 |
| inbound | TCP | 8082 | 0.0.0.0/0 |
| inbound | UDP | 53 | 192.168.110.0/24 |
| inbound | TCP | 8081 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | mig-sg-02 |
| **CSP Security Group ID** | sg-00808f305bd1562e3 |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 2 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.5312 |
| **Per Day** | $12.75 |
| **Per Month (30 days)** | $382.46 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AWS | ap-northeast-2 | 2 | $0.5312 | $382.46 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1 | c5a.2xlarge | $0.3440 | $247.68 |
| migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1 | t3a.xlarge | $0.1872 | $134.78 |


