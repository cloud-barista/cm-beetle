# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-03 07:48:45

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
| my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | 538b3c77-3777-4e37-be29-1004fcda8048 | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 103.218.158.79<br>**Private IP:** 10.0.1.4<br>**SGs:** my09-mig-sg-02<br>**SSH:** my09-mig-sshkey-01 |
| my09-ng-influxdb-back-1 | cb04341b-8725-47fd-b5e4-804c24dd8bba | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 133.186.222.44<br>**Private IP:** 10.0.1.101<br>**SGs:** my09-mig-sg-01<br>**SSH:** my09-mig-sshkey-01 |
| my09-ng-influxdb-back-2 | f3c5c55a-727c-450f-92c9-e1f6263e3ef7 | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 133.186.213.214<br>**Private IP:** 10.0.1.100<br>**SGs:** my09-mig-sg-01<br>**SSH:** my09-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my09-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my09-mig-vnet-01 |
| **CSP VNet ID** | 5607e2b9-d31d-445b-8235-826821764519 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | nhn-kr1 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my09-mig-subnet-01 | e9cf711b-2444-44d6-9963-14630f58b051 | 10.0.1.0/24 |  |

### Network Load Balancers (NLB)

#### NLB: my09-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my09-ng-influxdb-back |
| **CSP NLB Name** | tbsc944nleu70b5070vb |
| **CSP NLB ID** | 52664d9e-3fbe-4d84-8d33-a76fc970d5f8 |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | nhn-kr1 |
| **Listener** | TCP:9999 (IP: 133.186.251.215) |
| **Target Group** | TCP:8086 (Node Group: my09-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my09-ng-influxdb-back-1 | cb04341b-8725-47fd-b5e4-804c24dd8bba | 10.0.1.101 | 133.186.222.44 | Running |
| my09-ng-influxdb-back-2 | f3c5c55a-727c-450f-92c9-e1f6263e3ef7 | 10.0.1.100 | 133.186.213.214 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my09-mig-sshkey-01 | tbm9mtmk1plihtn1bjvn | cb-user | 10:b5:7d:7b:dd:65:52:32:69:6c:b8:d3:3d:d0:f9:a0 |

### Security Groups

#### Security Group: my09-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my09-mig-sg-01 |
| **CSP Security Group ID** | 9deb893f-653e-4aa8-8309-aaf905bc683c |
| **VNet** | my09-mig-vnet-01 |
| **Rule Count** | 8 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | ALL |  | 0.0.0.0/0 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/21 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.1.0/24 |

#### Security Group: my09-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my09-mig-sg-02 |
| **CSP Security Group ID** | b7e42c3f-6ce0-4a99-adb3-8c43a9add9a7 |
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
| my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | c2.c2m2 | $0.0422 | $30.37 |
| my09-ng-influxdb-back-1 | r2.c4m16 | $0.1954 | $140.66 |
| my09-ng-influxdb-back-2 | r2.c4m16 | $0.1954 | $140.66 |


