# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:20:17

**Namespace:** mig01

**Infra Name:** my05-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my05-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | TENCENT |
| **Target Region** | ap-seoul |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| SA5.LARGE16 | 4 | 16.0 | - | x86_64 |  | $0.2200 | 2 |
| SA5.MEDIUM2 | 2 | 2.0 | - | x86_64 |  | $0.0600 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| img-487zeit5 | Ubuntu Server 22.04 LTS 64bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my05-ng-influxdb-back-1 | ins-fxa70pv9 | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit (Ubuntu Server 22.04 LTS 64bit) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.166.24.5<br>**Private IP:** 10.0.1.2<br>**SGs:** my05-mig-sg-01<br>**SSH:** my05-mig-sshkey-01 |
| my05-ng-influxdb-back-2 | ins-6t8n96er | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit (Ubuntu Server 22.04 LTS 64bit) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.155.212.36<br>**Private IP:** 10.0.1.6<br>**SGs:** my05-mig-sg-01<br>**SSH:** my05-mig-sshkey-01 |
| my05-ng-ip-10-0-1-30-62161624-1 | ins-px4j8x8x | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04 LTS 64bit (Ubuntu Server 22.04 LTS 64bit) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.155.211.124<br>**Private IP:** 10.0.1.11<br>**SGs:** my05-mig-sg-02<br>**SSH:** my05-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my05-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my05-mig-vnet-01 |
| **CSP VNet ID** | vpc-p2o55rts |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | tencent-ap-seoul |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my05-mig-subnet-01 | subnet-4kcu8oj9 | 10.0.1.0/24 | ap-seoul-2 |

### Network Load Balancers (NLB)

#### NLB: my05-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my05-ng-influxdb-back |
| **CSP NLB Name** | tb4lovtl25d27niar0j5 |
| **CSP NLB ID** | lb-80vzpm0r |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | tencent-ap-seoul |
| **Listener** | TCP:9999 (DNS: lb-80vzpm0r-z1pcdkye6afvigxn.clb.kr-tencentclb.com) |
| **Target Group** | TCP:8086 (Node Group: my05-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my05-ng-influxdb-back-1 | ins-fxa70pv9 | 10.0.1.2 | 43.166.24.5 | Running |
| my05-ng-influxdb-back-2 | ins-6t8n96er | 10.0.1.6 | 43.155.212.36 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my05-mig-sshkey-01 | skey-b1dxxqvp |  |  |

### Security Groups

#### Security Group: my05-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my05-mig-sg-01 |
| **CSP Security Group ID** | sg-43psw731 |
| **VNet** | my05-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my05-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my05-mig-sg-02 |
| **CSP Security Group ID** | sg-gk3esedz |
| **VNet** | my05-mig-vnet-01 |
| **Rule Count** | 4 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.5000 |
| **Per Day** | $12.00 |
| **Per Month (30 days)** | $360.00 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| TENCENT | ap-seoul | 3 | $0.5000 | $360.00 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my05-ng-influxdb-back-1 | SA5.LARGE16 | $0.2200 | $158.40 |
| my05-ng-influxdb-back-2 | SA5.LARGE16 | $0.2200 | $158.40 |
| my05-ng-ip-10-0-1-30-62161624-1 | SA5.MEDIUM2 | $0.0600 | $43.20 |


