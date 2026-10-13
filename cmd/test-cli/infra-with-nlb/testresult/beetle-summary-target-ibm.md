# Target Cloud Infrastructure Summary

**Generated At:** 2026-10-12 11:07:48

**Namespace:** mig01

**Infra Name:** my06-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my06-infra101 |
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
| bx2-4x16 | 4 | 16.0 | - | x86_64 |  | $0.2410 | 2 |
| cx2-2x4 | 2 | 4.0 | - | x86_64 |  | $0.1060 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| r026-b879575b-8209-4e6a-a46e-0e01adf56aaf | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | 02h7_b73ff4f6-7d9a-4c6d-93ab-a0d4c6ef31c5 | Running | 2 vCPU, 4.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-mig-vnet-01<br>**Subnet:** my06-mig-subnet-01<br>**Public IP:** 159.23.93.202<br>**Private IP:** 10.0.1.5<br>**SGs:** my06-mig-sg-02<br>**SSH:** my06-mig-sshkey-01 |
| my06-ng-influxdb-back-1 | 02h7_70df7833-98d3-456c-8aa7-43a3f51f73d6 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-mig-vnet-01<br>**Subnet:** my06-mig-subnet-01<br>**Public IP:** 159.23.99.21<br>**Private IP:** 10.0.1.6<br>**SGs:** my06-mig-sg-01<br>**SSH:** my06-mig-sshkey-01 |
| my06-ng-influxdb-back-2 | 02h7_d3acbb7a-5ae4-4bb3-9d9e-62c1a321294f | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-mig-vnet-01<br>**Subnet:** my06-mig-subnet-01<br>**Public IP:** 159.23.103.103<br>**Private IP:** 10.0.1.4<br>**SGs:** my06-mig-sg-01<br>**SSH:** my06-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my06-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my06-mig-vnet-01 |
| **CSP VNet ID** | r026-d88398f2-3008-4600-a9b7-09de248c9d82 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my06-mig-subnet-01 | 02h7-0cdc9592-fc29-40bb-a412-28b334f8caae | 10.0.1.0/24 | au-syd-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my06-mig-sshkey-01 | r026-2afea093-cd50-4d8f-bb32-0bd32af60c50 |  | SHA256:N5q9VF3QjxUyO0XMkHwWOkfVy0gzX0BtbXimv4+j3bg |

### Security Groups

#### Security Group: my06-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my06-mig-sg-01 |
| **CSP Security Group ID** | r026-09cc8d9b-d306-4c90-bf02-b201694e955a |
| **VNet** | my06-mig-vnet-01 |
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

#### Security Group: my06-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my06-mig-sg-02 |
| **CSP Security Group ID** | r026-0a6d6edb-8e9f-4250-a320-b0138925d42a |
| **VNet** | my06-mig-vnet-01 |
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
| **Per Hour** | $0.5880 |
| **Per Day** | $14.11 |
| **Per Month (30 days)** | $423.36 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| IBM | au-syd | 3 | $0.5880 | $423.36 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | cx2-2x4 | $0.1060 | $76.32 |
| my06-ng-influxdb-back-1 | bx2-4x16 | $0.2410 | $173.52 |
| my06-ng-influxdb-back-2 | bx2-4x16 | $0.2410 | $173.52 |


