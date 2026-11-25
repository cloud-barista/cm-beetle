# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:20:42

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
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| n2-highcpu-2 | 2 | 2.0 | - | x86_64 |  | $0.0921 | 1 |
| n4-standard-4 | 4 | 16.0 | - | x86_64 |  | $0.2330 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021 | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21 | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my03-ng-influxdb-back-1 | tbjg90vgd6kbujirukfr | Running | 4 vCPU, 16.0 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21) | **VNet:** my03-mig-vnet-01<br>**Subnet:** my03-mig-subnet-01<br>**Public IP:** 34.64.48.98<br>**Private IP:** 10.0.1.4<br>**SGs:** my03-mig-sg-01<br>**SSH:** my03-mig-sshkey-01 |
| my03-ng-influxdb-back-2 | tb1lee0l2cp9nso74pq6 | Running | 4 vCPU, 16.0 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21) | **VNet:** my03-mig-vnet-01<br>**Subnet:** my03-mig-subnet-01<br>**Public IP:** 34.158.205.11<br>**Private IP:** 10.0.1.2<br>**SGs:** my03-mig-sg-01<br>**SSH:** my03-mig-sshkey-01 |
| my03-ng-ip-10-0-1-30-62161624-1 | tbr066qbkvoha4fjthgs | Running | 2 vCPU, 2.0 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21) | **VNet:** my03-mig-vnet-01<br>**Subnet:** my03-mig-subnet-01<br>**Public IP:** 34.22.89.144<br>**Private IP:** 10.0.1.3<br>**SGs:** my03-mig-sg-02<br>**SSH:** my03-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my03-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my03-mig-vnet-01 |
| **CSP VNet ID** | tb5dn9aum2qip3phdv8n |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | gcp-asia-northeast3 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my03-mig-subnet-01 | tbgi6cba5avpr77flia4 | 10.0.1.0/24 |  |

### Network Load Balancers (NLB)

#### NLB: my03-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my03-ng-influxdb-back |
| **CSP NLB Name** | tbiu3c4ki47g9gl6fr82 |
| **CSP NLB ID** | tbiu3c4ki47g9gl6fr82 |
| **Type / Scope** |  / REGION |
| **Connection** | gcp-asia-northeast3 |
| **Listener** | TCP:8086 (IP: 34.50.41.226) |
| **Target Group** | TCP:8086 (Node Group: my03-ng-influxdb-back, 2 nodes) |
| **Health Checker** | HTTP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my03-ng-influxdb-back-1 | tbjg90vgd6kbujirukfr | 10.0.1.4 | 34.64.48.98 | Running |
| my03-ng-influxdb-back-2 | tb1lee0l2cp9nso74pq6 | 10.0.1.2 | 34.158.205.11 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my03-mig-sshkey-01 | tbdnf6sndttnf336lgim |  |  |

### Security Groups

#### Security Group: my03-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my03-mig-sg-01 |
| **CSP Security Group ID** | tb9uu0hejtsjn923sfvd |
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
| **CSP Security Group ID** | tbplq2umhkn1upbkpp57 |
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
| **Per Hour** | $0.5580 |
| **Per Day** | $13.39 |
| **Per Month (30 days)** | $401.75 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| GCP | asia-northeast3 | 3 | $0.5580 | $401.75 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my03-ng-influxdb-back-1 | n4-standard-4 | $0.2330 | $167.73 |
| my03-ng-influxdb-back-2 | n4-standard-4 | $0.2330 | $167.73 |
| my03-ng-ip-10-0-1-30-62161624-1 | n2-highcpu-2 | $0.0921 | $66.29 |


