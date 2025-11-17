# Target Cloud Infrastructure Summary

**Generated At:** 2025-11-17 05:06:42

**Namespace:** mig01

**MCI Name:** mmci01

---

## Overview

| Property | Value |
|----------|-------|
| **MCI Name** | mmci01 |
| **Description** | a recommended multi-cloud infrastructure |
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
| e2-standard-2 | 2 | 7.8 | - | x86_64 |  | $0.0860 | 1 |
| e2-standard-4 | 4 | 15.6 | - | x86_64 |  | $0.1719 | 1 |
| e2-small | 2 | 2.0 | - | x86_64 |  | $0.0215 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251111 | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2025-11-11 | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1 | d4daq1qp2foc73bcek10 | Running | 4 vCPU, 15.6 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2025-11-11 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2025-11-11) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 34.64.97.113<br>**Private IP:** 10.0.1.3<br>**SGs:** mig-sg-02<br>**SSH:** mig-sshkey-01 |
| migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1 | d4daq1qp2foc73bcek00 | Running | 2 vCPU, 2.0 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2025-11-11 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2025-11-11) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 34.64.253.204<br>**Private IP:** 10.0.1.4<br>**SGs:** mig-sg-01<br>**SSH:** mig-sshkey-01 |
| migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1 | d4daq1qp2foc73bcek20 | Running | 2 vCPU, 7.8 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2025-11-11 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2025-11-11) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 34.47.82.114<br>**Private IP:** 10.0.1.2<br>**SGs:** mig-sg-03<br>**SSH:** mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | mig-vnet-01 |
| **CSP VNet ID** | d4danaip2foc73bcejs0 |
| **CIDR Block** | GCP VPC does not support IPv4_CIDR |
| **Connection** | gcp-asia-northeast3 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| mig-subnet-01 | d4danaip2foc73bcejsg | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| mig-sshkey-01 | d4daniip2foc73bcejt0 |  |  |

### Security Groups

#### Security Group: mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | mig-sg-01 |
| **CSP Security Group ID** | d4danj2p2foc73bcejtg |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 14 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 80 | 0.0.0.0/0 |
| inbound | TCP | 443 | 0.0.0.0/0 |
| inbound | TCP | 8080 | 0.0.0.0/0 |
| inbound | TCP | 9113 | 10.0.0.0/16 |
| inbound | UDP | 9113 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |

#### Security Group: mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | mig-sg-02 |
| **CSP Security Group ID** | d4dao7qp2foc73bceju0 |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 2049 | 0.0.0.0/0 |
| inbound | UDP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 111 | 0.0.0.0/0 |
| inbound | UDP | 111 | 0.0.0.0/0 |
| inbound | TCP | 20048 | 10.0.0.0/16 |
| inbound | UDP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 32803 | 10.0.0.0/16 |
| inbound | UDP | 32803 | 10.0.0.0/16 |
| inbound | TCP | 9100 | 10.0.0.0/16 |
| inbound | UDP | 9100 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |

#### Security Group: mig-sg-03

| Property | Value |
|----------|-------|
| **Name** | mig-sg-03 |
| **CSP Security Group ID** | d4dap5ip2foc73bcejug |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 3306 | 10.0.0.0/16 |
| inbound | UDP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 4567 | 10.0.0.0/16 |
| inbound | UDP | 4567 | 10.0.0.0/16 |
| inbound | TCP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4568 | 10.0.0.0/16 |
| inbound | TCP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 4444 | 10.0.0.0/16 |
| inbound | TCP | 9104 | 10.0.0.0/16 |
| inbound | UDP | 9104 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.2794 |
| **Per Day** | $6.71 |
| **Per Month (30 days)** | $201.16 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| GCP | asia-northeast3 | 3 | $0.2794 | $201.16 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1 | e2-standard-4 | $0.1719 | $123.79 |
| migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1 | e2-small | $0.0215 | $15.47 |
| migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1 | e2-standard-2 | $0.0860 | $61.90 |


