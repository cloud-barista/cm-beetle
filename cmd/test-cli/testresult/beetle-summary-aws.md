# Infrastructure Summary

**Generated At:** 2025-11-03 07:15:11

**Namespace:** mig01

**MCI Name:** mmci01

**Summary Version:** 1.0

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
| **Label** | sys.id=mmci01, sys.labelType=mci, sys.manager=cb-tumblebug, sys.name=mmci01, sys.namespace=mig01, sys.uid=d445cguqjs728pvqbejg, sys.description=a recommended multi-cloud infrastructure |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | Usage Count |
|------|-------|--------------|-----|--------------|-----------|-----------------|-------------|
| c5a.2xlarge | 8 | 16.0 | - | x86_64 |  | $0.3440 | 1 |
| t3a.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1872 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Disk Type | Disk Size | Usage Count |
|------|--------------|---------|-------------|--------------|-----------|-----------|-------------|
| ami-010be25c3775061c9 | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | -1 GB | 2 |

### Virtual Machines

| Instance Name | CSP Instance ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------------|-----------------|--------|-------------------------|-------|------|
| migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1 | i-093b4b7c722ed9ec5 | Running | 8 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015) | VNet: mig-vnet-01, Subnet: mig-subnet-01, Public IP: 43.201.85.138, Private IP: 192.168.110.167, SGs: mig-sg-02, SSH: mig-sshkey-01 |
| migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1 | i-042a0b12d56ac1aba | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015) | VNet: mig-vnet-01, Subnet: mig-subnet-01, Public IP: 3.38.135.44, Private IP: 192.168.110.39, SGs: mig-sg-01, SSH: mig-sshkey-01 |


## Network Resources

### Virtual Networks

#### VNet: mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | mig-vnet-01 |
| **CSP VNet ID** | vpc-0200cd398ed7c7f17 |
| **CIDR Block** | 192.168.96.0/19 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| mig-subnet-01 | subnet-0c464077de36a8e01 | 192.168.110.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| mig-sshkey-01 | d445cemqjs728pvqbei0 |  | 79:cf:60:4f:65:b8:15:93:d2:5d:45:9f:c5:96:f8:88:34:68:32:06 |

### Security Groups

#### Security Group: mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | mig-sg-01 |
| **CSP Security Group ID** | sg-092b0916e4f0b9a3c |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 6 |

**Security Group Rules:**

| Direction | Protocol | Port | CIDR |
|-----------|----------|------|------|
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
| **CSP Security Group ID** | sg-0e7fae5aa026fb64a |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 2 |

**Security Group Rules:**

| Direction | Protocol | Port | CIDR |
|-----------|----------|------|------|
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


