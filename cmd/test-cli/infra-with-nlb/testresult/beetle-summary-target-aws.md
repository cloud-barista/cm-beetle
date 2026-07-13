# Target Cloud Infrastructure Summary

**Generated At:** 2026-07-13 09:03:53

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
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
| t3a.small | 2 | 2.0 | - | x86_64 |  | $0.0234 | 1 |
| t3a.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1872 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ami-0afe1fd15675c3f15 | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | i-07f638bb18bbb13e6 | Running | 2 vCPU, 2.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 13.209.82.179<br>**Private IP:** 10.0.1.7<br>**SGs:** my-mig-sg-02<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-1 | i-0f63a3a34b581895d | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 3.38.214.237<br>**Private IP:** 10.0.1.208<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-2 | i-05fa2460fe53c2da4 | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 54.180.128.3<br>**Private IP:** 10.0.1.177<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-vnet-01 |
| **CSP VNet ID** | vpc-0e032ee5b40ca2cdd |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-mig-subnet-01 | subnet-0818566fb5fde1e2e | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-mig-sshkey-01 | tbt1bt5hiiqoqi38lfv2 |  | 2a:61:89:ab:5d:44:cb:cf:66:1d:da:b7:30:d6:8f:2d:d7:8a:57:81 |

### Security Groups

#### Security Group: my-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-01 |
| **CSP Security Group ID** | sg-04f9aacbfff1339fd |
| **VNet** | my-mig-vnet-01 |
| **Rule Count** | 5 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-02 |
| **CSP Security Group ID** | sg-06bff0ecdf53345f6 |
| **VNet** | my-mig-vnet-01 |
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
| **Per Hour** | $0.3978 |
| **Per Day** | $9.55 |
| **Per Month (30 days)** | $286.42 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AWS | ap-northeast-2 | 3 | $0.3978 | $286.42 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | t3a.small | $0.0234 | $16.85 |
| my-ng-influxdb-back-1 | t3a.xlarge | $0.1872 | $134.78 |
| my-ng-influxdb-back-2 | t3a.xlarge | $0.1872 | $134.78 |


