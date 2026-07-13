# Target Cloud Infrastructure Summary

**Generated At:** 2026-07-13 09:42:43

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | IBM |
| **Target Region** | au-syd |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| nxf-2x2 | 2 | 2.0 | - | x86_64 |  | $0.0940 | 1 |
| bxf-4x16 | 4 | 16.0 | - | x86_64 |  | $0.2350 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599 | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | 02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3 | Running | 2 vCPU, 2.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 159.23.93.219<br>**Private IP:** 10.0.1.4<br>**SGs:** my-mig-sg-02<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-1 | 02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 159.23.98.228<br>**Private IP:** 10.0.1.5<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-2 | 02h7_57bb1846-1e20-4101-b3a3-351f585f49d7 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 159.23.93.220<br>**Private IP:** 10.0.1.6<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-vnet-01 |
| **CSP VNet ID** | r026-973a751e-d6bc-4105-a2b1-e76438bea613 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-mig-subnet-01 | 02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a | 10.0.1.0/24 | au-syd-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-mig-sshkey-01 | r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0 |  | SHA256:W9+bkBcpsLeBdDiDOPUXfFQJ/q1cPg31WqKwqVudX/E |

### Security Groups

#### Security Group: my-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-01 |
| **CSP Security Group ID** | r026-e9472d3e-f8b2-4c3e-a8e5-be6a646610fa |
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
| **CSP Security Group ID** | r026-95fb8ca5-af0a-4f5c-b195-cee7dfe0f2d8 |
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
| **Per Hour** | $0.5640 |
| **Per Day** | $13.54 |
| **Per Month (30 days)** | $406.08 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| IBM | au-syd | 3 | $0.5640 | $406.08 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | nxf-2x2 | $0.0940 | $67.68 |
| my-ng-influxdb-back-1 | bxf-4x16 | $0.2350 | $169.20 |
| my-ng-influxdb-back-2 | bxf-4x16 | $0.2350 | $169.20 |


