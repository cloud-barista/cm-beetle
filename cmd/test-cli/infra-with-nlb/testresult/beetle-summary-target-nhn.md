# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:28:42

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
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| r2.c4m16 | 4 | 16.0 | - | x86_64 | default | $0.1954 | 2 |
| c2.c2m2 | 2 | 2.0 | - | x86_64 | default | $0.0422 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| 0f07c795-2a46-44fc-a61b-fa0d96763ce2 | Ubuntu Server 22.04.5 LTS (2026.03.10) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my09-ng-influxdb-back-1 | 1161a76a-b53f-4325-b20c-2d4507c069eb | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 133.186.153.17<br>**Private IP:** 10.0.1.19<br>**SGs:** my09-mig-sg-01<br>**SSH:** my09-mig-sshkey-01 |
| my09-ng-influxdb-back-2 | d68821e9-9436-4815-9890-63d0f84c5beb | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 125.6.37.131<br>**Private IP:** 10.0.1.50<br>**SGs:** my09-mig-sg-01<br>**SSH:** my09-mig-sshkey-01 |
| my09-ng-ip-10-0-1-30-62161624-1 | a06b8a06-eb5c-4842-96ff-752c8b4afa2e | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 125.6.36.178<br>**Private IP:** 10.0.1.63<br>**SGs:** my09-mig-sg-02<br>**SSH:** my09-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my09-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my09-mig-vnet-01 |
| **CSP VNet ID** | 3cf6f879-d1de-49d2-a60e-a08a87423b32 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | nhn-kr1 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my09-mig-subnet-01 | 8618fe2a-2640-4777-8c22-7f9e6adecef1 | 10.0.1.0/24 |  |

### Network Load Balancers (NLB)

#### NLB: my09-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my09-ng-influxdb-back |
| **CSP NLB Name** | tbfpvpl7ovve2g0acn9a |
| **CSP NLB ID** | 55b474f6-1ec8-495c-a9a8-a4550305844d |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | nhn-kr1 |
| **Listener** | TCP:9999 (IP: 125.6.37.15) |
| **Target Group** | TCP:8086 (Node Group: my09-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my09-ng-influxdb-back-1 | 1161a76a-b53f-4325-b20c-2d4507c069eb | 10.0.1.19 | 133.186.153.17 | Running |
| my09-ng-influxdb-back-2 | d68821e9-9436-4815-9890-63d0f84c5beb | 10.0.1.50 | 125.6.37.131 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my09-mig-sshkey-01 | tb3d1kt2cp0pjrblu83g | cb-user | 69:ab:70:c9:c6:f4:b1:8f:0b:d1:12:26:ac:a7:e4:6e |

### Security Groups

#### Security Group: my09-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my09-mig-sg-01 |
| **CSP Security Group ID** | 126aeb07-4848-44e6-b4ce-314dfb7240ef |
| **VNet** | my09-mig-vnet-01 |
| **Rule Count** | 8 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 8086 | 10.0.0.0/21 |
| inbound | TCP | 8086 | 10.0.1.0/24 |

#### Security Group: my09-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my09-mig-sg-02 |
| **CSP Security Group ID** | 3b7c1ed5-dca7-4cdb-b18e-a766a72ade33 |
| **VNet** | my09-mig-vnet-01 |
| **Rule Count** | 4 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 9999 | 0.0.0.0/0 |


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
| my09-ng-influxdb-back-1 | r2.c4m16 | $0.1954 | $140.66 |
| my09-ng-influxdb-back-2 | r2.c4m16 | $0.1954 | $140.66 |
| my09-ng-ip-10-0-1-30-62161624-1 | c2.c2m2 | $0.0422 | $30.37 |


