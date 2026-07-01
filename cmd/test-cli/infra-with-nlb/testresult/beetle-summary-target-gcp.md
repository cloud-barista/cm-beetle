# Target Cloud Infrastructure Summary

**Generated At:** 2026-07-01 05:59:23

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
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
| e2-standard-4 | 4 | 15.6 | - | x86_64 |  | $0.1719 | 2 |
| e2-highcpu-2 | 2 | 2.0 | - | x86_64 |  | $0.0635 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260530 | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-05-30 | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 1 |
| https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260530 | Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-05-30 | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 2 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | tbebk1kfvnmb8etin43s | Running | 2 vCPU, 2.0 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-05-30 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-05-30) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 34.50.15.190<br>**Private IP:** 10.0.1.3<br>**SGs:** my-mig-sg-02<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-1 | tbpjdbmbl77kl71d5tpq | Running | 4 vCPU, 15.6 GiB | Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-05-30 (Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-05-30) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 34.64.43.29<br>**Private IP:** 10.0.1.4<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-2 | tbo27a2ne7re56lv6ldl | Running | 4 vCPU, 15.6 GiB | Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-05-30 (Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-05-30) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 34.50.12.245<br>**Private IP:** 10.0.1.2<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-vnet-01 |
| **CSP VNet ID** | tb8o38ikh0oe0ra0gc5c |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | gcp-asia-northeast3 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-mig-subnet-01 | tbmsmd7ts1rblt5lrfp0 | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-mig-sshkey-01 | tbgtbihj6p3uke51otql |  |  |

### Security Groups

#### Security Group: my-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-01 |
| **CSP Security Group ID** | tbdg1d0n4ipeqjodp89d |
| **VNet** | my-mig-vnet-01 |
| **Rule Count** | 5 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-02 |
| **CSP Security Group ID** | tb0ih3nphdhfkpihqmpv |
| **VNet** | my-mig-vnet-01 |
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
| **Per Hour** | $0.4074 |
| **Per Day** | $9.78 |
| **Per Month (30 days)** | $293.32 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| GCP | asia-northeast3 | 3 | $0.4074 | $293.32 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | e2-highcpu-2 | $0.0635 | $45.74 |
| my-ng-influxdb-back-1 | e2-standard-4 | $0.1719 | $123.79 |
| my-ng-influxdb-back-2 | e2-standard-4 | $0.1719 | $123.79 |


