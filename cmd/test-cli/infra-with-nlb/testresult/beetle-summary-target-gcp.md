# Target Cloud Infrastructure Summary

**Generated At:** 2026-10-12 11:01:03

**Namespace:** mig01

**Infra Name:** my03-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my03-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | GCP |
| **Target Region** | asia-northeast3 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| n4-standard-4 | 4 | 15.6 | - | x86_64 |  | $0.2330 | 2 |
| n1-highcpu-2 | 2 | 1.8 | - | x86_64 |  | $0.0910 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820 | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20 | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | tb34cj51eke012cm5vec | Running | 2 vCPU, 1.8 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20) | **VNet:** my03-mig-vnet-01<br>**Subnet:** my03-mig-subnet-01<br>**Public IP:** 34.50.50.20<br>**Private IP:** 10.0.1.3<br>**SGs:** my03-mig-sg-02<br>**SSH:** my03-mig-sshkey-01 |
| my03-ng-influxdb-back-1 | tb6iphm7mug3sutkbmrj | Running | 4 vCPU, 15.6 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20) | **VNet:** my03-mig-vnet-01<br>**Subnet:** my03-mig-subnet-01<br>**Public IP:** 8.230.11.202<br>**Private IP:** 10.0.1.4<br>**SGs:** my03-mig-sg-01<br>**SSH:** my03-mig-sshkey-01 |
| my03-ng-influxdb-back-2 | tb0vbtatce8oe2gk7l55 | Running | 4 vCPU, 15.6 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20) | **VNet:** my03-mig-vnet-01<br>**Subnet:** my03-mig-subnet-01<br>**Public IP:** 34.64.253.99<br>**Private IP:** 10.0.1.2<br>**SGs:** my03-mig-sg-01<br>**SSH:** my03-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my03-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my03-mig-vnet-01 |
| **CSP VNet ID** | tb6ei259h9mvdch4rumt |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | gcp-asia-northeast3 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my03-mig-subnet-01 | tbvssd98bnj6fq03668o | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my03-mig-sshkey-01 | tbo0ov0jqt6lbp7d7e4m |  |  |

### Security Groups

#### Security Group: my03-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my03-mig-sg-01 |
| **CSP Security Group ID** | tbjs9h2h9ossf4jdga7u |
| **VNet** | my03-mig-vnet-01 |
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

#### Security Group: my03-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my03-mig-sg-02 |
| **CSP Security Group ID** | tbk826ghvc2um7m13tup |
| **VNet** | my03-mig-vnet-01 |
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
| **Per Hour** | $0.5569 |
| **Per Day** | $13.37 |
| **Per Month (30 days)** | $400.97 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| GCP | asia-northeast3 | 3 | $0.5569 | $400.97 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | n1-highcpu-2 | $0.0910 | $65.51 |
| my03-ng-influxdb-back-1 | n4-standard-4 | $0.2330 | $167.73 |
| my03-ng-influxdb-back-2 | n4-standard-4 | $0.2330 | $167.73 |


