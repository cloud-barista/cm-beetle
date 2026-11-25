# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:29:44

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
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| bx2-4x16 | 4 | 16.0 | - | x86_64 |  | $0.2470 | 2 |
| cx2-2x4 | 2 | 4.0 | - | x86_64 |  | $0.1090 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| r026-6245bce7-1353-421f-9eb8-47aa0ed18c50 | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my06-ng-influxdb-back-1 | 02h7_65a21f53-e4b2-4464-b89c-2d1ad1261b16 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-mig-vnet-01<br>**Subnet:** my06-mig-subnet-01<br>**Public IP:** 159.23.97.204<br>**Private IP:** 10.0.1.4<br>**SGs:** my06-mig-sg-01<br>**SSH:** my06-mig-sshkey-01 |
| my06-ng-influxdb-back-2 | 02h7_ffe75cff-d1c8-40d1-822f-1ece581c6b69 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-mig-vnet-01<br>**Subnet:** my06-mig-subnet-01<br>**Public IP:** 159.23.91.103<br>**Private IP:** 10.0.1.6<br>**SGs:** my06-mig-sg-01<br>**SSH:** my06-mig-sshkey-01 |
| my06-ng-ip-10-0-1-30-62161624-1 | 02h7_1984afd7-a085-4a9e-82d2-e003457d44e3 | Running | 2 vCPU, 4.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-mig-vnet-01<br>**Subnet:** my06-mig-subnet-01<br>**Public IP:** 159.23.91.104<br>**Private IP:** 10.0.1.5<br>**SGs:** my06-mig-sg-02<br>**SSH:** my06-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my06-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my06-mig-vnet-01 |
| **CSP VNet ID** | r026-07690ed0-796b-433a-b7ae-30a7ffa211a7 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my06-mig-subnet-01 | 02h7-b9f42d6c-389f-44ee-9c18-82a38e5ce5a7 | 10.0.1.0/24 | au-syd-1 |

### Network Load Balancers (NLB)

#### NLB: my06-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my06-ng-influxdb-back |
| **CSP NLB Name** | tbqofi5aon8o516gc3pk |
| **CSP NLB ID** | r026-479f3c09-227a-40cf-9855-61794958e8bb |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | ibm-au-syd |
| **Listener** | TCP:9999 (DNS: 479f3c09-au-syd.lb.appdomain.cloud, IP: 159.23.91.106) |
| **Target Group** | TCP:8086 (Node Group: my06-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 5s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my06-ng-influxdb-back-1 | 02h7_65a21f53-e4b2-4464-b89c-2d1ad1261b16 | 10.0.1.4 | 159.23.97.204 | Running |
| my06-ng-influxdb-back-2 | 02h7_ffe75cff-d1c8-40d1-822f-1ece581c6b69 | 10.0.1.6 | 159.23.91.103 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my06-mig-sshkey-01 | r026-12c6f57d-43ad-4812-a953-07c0a581a0ec |  | SHA256:YazISpGV/XQCiwMPwXCO8CAe32YGoahFfAcKLRyVHM4 |

### Security Groups

#### Security Group: my06-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my06-mig-sg-01 |
| **CSP Security Group ID** | r026-17b253ae-cdf8-42dd-8aec-7e87ac4dbb85 |
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
| **CSP Security Group ID** | r026-481fe27e-00ff-457e-9ccc-ec10465a9d66 |
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
| **Per Hour** | $0.6030 |
| **Per Day** | $14.47 |
| **Per Month (30 days)** | $434.16 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| IBM | au-syd | 3 | $0.6030 | $434.16 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my06-ng-influxdb-back-1 | bx2-4x16 | $0.2470 | $177.84 |
| my06-ng-influxdb-back-2 | bx2-4x16 | $0.2470 | $177.84 |
| my06-ng-ip-10-0-1-30-62161624-1 | cx2-2x4 | $0.1090 | $78.48 |


