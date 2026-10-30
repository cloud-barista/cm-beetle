# Target Cloud Infrastructure Summary

**Generated At:** 2026-10-30 12:51:03

**Namespace:** mig01

**Infra Name:** my01-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my01-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
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
| t3.small | 2 | 2.0 | - | x86_64 |  | $0.0260 | 1 |
| t3.xlarge | 4 | 16.0 | - | x86_64 |  | $0.2080 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ami-012a353bb3afb92ee | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | i-0e9a8dff79587097e | Running | 2 vCPU, 2.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 54.180.141.244<br>**Private IP:** 10.0.1.201<br>**SGs:** my01-mig-sg-02<br>**SSH:** my01-mig-sshkey-01 |
| my01-ng-influxdb-back-1 | i-0dee049fb9412465b | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 13.125.79.237<br>**Private IP:** 10.0.1.221<br>**SGs:** my01-mig-sg-01<br>**SSH:** my01-mig-sshkey-01 |
| my01-ng-influxdb-back-2 | i-0cf1d4327934156dd | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 3.35.54.160<br>**Private IP:** 10.0.1.253<br>**SGs:** my01-mig-sg-01<br>**SSH:** my01-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my01-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my01-mig-vnet-01 |
| **CSP VNet ID** | vpc-02e0dfd3cb3ff1fcb |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my01-mig-subnet-01 | subnet-09bb75a1817ae8dbb | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my01-mig-sshkey-01 | tbq5thtqubnf05ts4une |  | ef:95:1c:cc:1d:6c:3f:1a:f5:c7:87:5e:bf:4d:8e:0d:fd:75:16:d5 |

### Security Groups

#### Security Group: my01-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my01-mig-sg-01 |
| **CSP Security Group ID** | sg-0704eaa2a1108babc |
| **VNet** | my01-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my01-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my01-mig-sg-02 |
| **CSP Security Group ID** | sg-082244dd31713d22a |
| **VNet** | my01-mig-vnet-01 |
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
| **Per Hour** | $0.4420 |
| **Per Day** | $10.61 |
| **Per Month (30 days)** | $318.24 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AWS | ap-northeast-2 | 3 | $0.4420 | $318.24 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | t3.small | $0.0260 | $18.72 |
| my01-ng-influxdb-back-1 | t3.xlarge | $0.2080 | $149.76 |
| my01-ng-influxdb-back-2 | t3.xlarge | $0.2080 | $149.76 |


