# Target Cloud Infrastructure Summary

**Generated At:** 2026-10-12 11:01:11

**Namespace:** mig01

**Infra Name:** my09-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my09-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | NHN |
| **Target Region** | KR1 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| c2.c2m2 | 2 | 2.0 | - | x86_64 | default | $0.0422 | 1 |
| r2.c4m16 | 4 | 16.0 | - | x86_64 | default | $0.1954 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| 0f07c795-2a46-44fc-a61b-fa0d96763ce2 | Ubuntu Server 22.04.5 LTS (2026.03.10) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | c54494b3-165c-4fe0-a924-32b4a0fec1fc | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 133.186.223.136<br>**Private IP:** 10.0.1.11<br>**SGs:** my09-mig-sg-02<br>**SSH:** my09-mig-sshkey-01 |
| my09-ng-influxdb-back-1 | 246fe00a-bfd2-402b-b140-22545a2f1c6c | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 133.186.209.144<br>**Private IP:** 10.0.1.87<br>**SGs:** my09-mig-sg-01<br>**SSH:** my09-mig-sshkey-01 |
| my09-ng-influxdb-back-2 | b0993afa-0c7a-44c7-809b-236ca241113f | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 133.186.209.114<br>**Private IP:** 10.0.1.100<br>**SGs:** my09-mig-sg-01<br>**SSH:** my09-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my09-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my09-mig-vnet-01 |
| **CSP VNet ID** | 63f62c96-b2f4-4602-8385-62d049b22434 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | nhn-kr1 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my09-mig-subnet-01 | b31a5d3f-d214-4469-92ca-c8588598023d | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my09-mig-sshkey-01 | tbhjbjpobn3enr3num8l | cb-user | 31:03:28:01:29:5a:fd:4d:b6:4f:08:58:97:1b:0c:ad |

### Security Groups

#### Security Group: my09-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my09-mig-sg-01 |
| **CSP Security Group ID** | ed3c8491-1514-4bce-8e7d-e982e62f49cf |
| **VNet** | my09-mig-vnet-01 |
| **Rule Count** | 8 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 9999 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.1.0/24 |
| inbound | TCP | 8086 | 10.0.0.0/21 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |

#### Security Group: my09-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my09-mig-sg-02 |
| **CSP Security Group ID** | 8b50179f-bdbc-4e39-aff3-1848ec9913c9 |
| **VNet** | my09-mig-vnet-01 |
| **Rule Count** | 4 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.4329 |
| **Per Day** | $10.39 |
| **Per Month (30 days)** | $311.69 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| NHN | KR1 | 3 | $0.4329 | $311.69 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | c2.c2m2 | $0.0422 | $30.37 |
| my09-ng-influxdb-back-1 | r2.c4m16 | $0.1954 | $140.66 |
| my09-ng-influxdb-back-2 | r2.c4m16 | $0.1954 | $140.66 |


