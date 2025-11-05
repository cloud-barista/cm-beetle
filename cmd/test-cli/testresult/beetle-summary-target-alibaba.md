# Target Cloud Infrastructure Summary

**Generated At:** 2025-11-05 04:36:06

**Namespace:** mig01

**MCI Name:** mmci01

---

## Overview

| Property | Value |
|----------|-------|
| **MCI Name** | mmci01 |
| **Description** | a recommended multi-cloud infrastructure |
| **Status** | Running:2 (R:2/2) |
| **Target Cloud** | ALIBABA |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 2 |
| **Running VMs** | 2 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| ecs.t6-c1m4.2xlarge | 8 | 32.0 | - | x86_64 |  | $0.3385 | 1 |
| ecs.t6-c1m4.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1693 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ubuntu_22_04_x64_20G_alibase_20250917.vhd | Ubuntu  22.04 64 bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 2 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1 | i-mj7bz4ymj6qw6yvwsl5f | Running | 8 vCPU, 32.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 8.213.149.114<br>**Private IP:** 192.168.110.39<br>**SGs:** mig-sg-02<br>**SSH:** mig-sshkey-01 |
| migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1 | i-mj7ajjr4yvgvpqlx0vts | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 8.220.200.156<br>**Private IP:** 192.168.110.40<br>**SGs:** mig-sg-01<br>**SSH:** mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | mig-vnet-01 |
| **CSP VNet ID** | vpc-mj7qgs1luk3szocsy5uqz |
| **CIDR Block** | 192.168.96.0/19 |
| **Connection** | alibaba-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| mig-subnet-01 | vsw-mj74458hhp0qq26mcaaed | 192.168.110.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| mig-sshkey-01 | d45d7reqjs728podgjhg |  | b58b791ee2aa2cbe49fd397c905eae08 |

### Security Groups

#### Security Group: mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | mig-sg-01 |
| **CSP Security Group ID** | sg-mj7ajjr4yvgvpqlsnc9t |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 53 | 192.168.110.0/24 |
| inbound | TCP | 8082 | 0.0.0.0/0 |
| inbound | TCP | 8081 | 0.0.0.0/0 |
| inbound | TCP | 10022 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | mig-sg-02 |
| **CSP Security Group ID** | sg-mj7c90xduxac8di2xg53 |
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
| **Per Hour** | $0.5078 |
| **Per Day** | $12.19 |
| **Per Month (30 days)** | $365.60 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| ALIBABA | ap-northeast-2 | 2 | $0.5078 | $365.60 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1 | ecs.t6-c1m4.2xlarge | $0.3385 | $243.73 |
| migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1 | ecs.t6-c1m4.xlarge | $0.1693 | $121.87 |


