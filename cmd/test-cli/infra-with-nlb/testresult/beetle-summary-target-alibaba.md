# Target Cloud Infrastructure Summary

**Generated At:** 2026-10-12 10:58:16

**Namespace:** mig01

**Infra Name:** my04-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my04-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
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
| ecs.e-c1m1.large | 2 | 2.0 | - | x86_64 |  | $0.0178 | 1 |
| ecs.e-c1m4.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1582 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ubuntu_22_04_x64_20G_alibase_20260810.vhd | Ubuntu  22.04 64 bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | i-mj78vnkdsw3hgl5t0j1i | Running | 2 vCPU, 2.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 47.80.249.240<br>**Private IP:** 10.0.1.73<br>**SGs:** my04-mig-sg-02<br>**SSH:** my04-mig-sshkey-01 |
| my04-ng-influxdb-back-1 | i-mj75o7wmbqfttpuy12n7 | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 47.80.251.194<br>**Private IP:** 10.0.1.72<br>**SGs:** my04-mig-sg-01<br>**SSH:** my04-mig-sshkey-01 |
| my04-ng-influxdb-back-2 | i-mj75o7wmbqfttpuy12n8 | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 47.80.250.25<br>**Private IP:** 10.0.1.74<br>**SGs:** my04-mig-sg-01<br>**SSH:** my04-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my04-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my04-mig-vnet-01 |
| **CSP VNet ID** | vpc-mj7r2m3wqm09p88ugm6if |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | alibaba-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my04-mig-subnet-01 | vsw-mj77koiujuio5e71st964 | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my04-mig-sshkey-01 | tbghcomebhflfpstm793 |  | 5b366c3fab1f3601a78f3e16aa933205 |

### Security Groups

#### Security Group: my04-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my04-mig-sg-01 |
| **CSP Security Group ID** | sg-mj755w3u7hdkfhqme0dt |
| **VNet** | my04-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my04-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my04-mig-sg-02 |
| **CSP Security Group ID** | sg-mj78vnkdsw3hgl5xrwmf |
| **VNet** | my04-mig-vnet-01 |
| **Rule Count** | 4 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.3342 |
| **Per Day** | $8.02 |
| **Per Month (30 days)** | $240.62 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| ALIBABA | ap-northeast-2 | 3 | $0.3342 | $240.62 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | ecs.e-c1m1.large | $0.0178 | $12.82 |
| my04-ng-influxdb-back-1 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |
| my04-ng-influxdb-back-2 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |


