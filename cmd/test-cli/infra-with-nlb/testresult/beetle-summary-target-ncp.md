# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:23:58

**Namespace:** mig01

**Infra Name:** my08-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my08-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | NCP |
| **Target Region** | KR |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| s4-g3 | 4 | 16.0 | - | x86_64 | default | $0.1745 | 2 |
| c2-g3 | 2 | 4.0 | - | x86_64 | default | $0.0677 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| 23214590 | ubuntu-22.04-base (Hypervisor:KVM) | Ubuntu 22.04 | Linux/UNIX | x86_64 | Common BlockStorage 1 | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my08-ng-influxdb-back-1 | 147164275 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 223.130.150.80<br>**Private IP:** 10.0.1.8<br>**SGs:** my08-mig-sg-01<br>**SSH:** my08-mig-sshkey-01 |
| my08-ng-influxdb-back-2 | 147164269 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 175.45.195.206<br>**Private IP:** 10.0.1.6<br>**SGs:** my08-mig-sg-01<br>**SSH:** my08-mig-sshkey-01 |
| my08-ng-ip-10-0-1-30-62161624-1 | 147164272 | Running | 2 vCPU, 4.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 223.130.159.128<br>**Private IP:** 10.0.1.7<br>**SGs:** my08-mig-sg-02<br>**SSH:** my08-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my08-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my08-mig-vnet-01 |
| **CSP VNet ID** | 153788 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ncp-kr |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my08-mig-subnet-01 | 336709 | 10.0.1.0/24 | KR-1 |

### Network Load Balancers (NLB)

#### NLB: my08-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my08-ng-influxdb-back |
| **CSP NLB Name** | tb4eeq1ff30hug6d5pqv |
| **CSP NLB ID** | 147164402 |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | ncp-kr |
| **Listener** | TCP:9999 (IP: 211.233.218.78) |
| **Target Group** | TCP:8086 (Node Group: my08-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 0s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my08-ng-influxdb-back-1 | 147164275 | 10.0.1.8 | 223.130.150.80 | Running |
| my08-ng-influxdb-back-2 | 147164269 | 10.0.1.6 | 175.45.195.206 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my08-mig-sshkey-01 | tbmo1vkhfnoa7ta0p33p | cb-user | 2a:8f:c2:e3:53:12:d9:33:23:0b:fc:24:37:da:ff:85 |

### Security Groups

#### Security Group: my08-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my08-mig-sg-01 |
| **CSP Security Group ID** | 423537 |
| **VNet** | my08-mig-vnet-01 |
| **Rule Count** | 10 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my08-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my08-mig-sg-02 |
| **CSP Security Group ID** | 423538 |
| **VNet** | my08-mig-vnet-01 |
| **Rule Count** | 8 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.4167 |
| **Per Day** | $10.00 |
| **Per Month (30 days)** | $300.02 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| NCP | KR | 3 | $0.4167 | $300.02 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my08-ng-influxdb-back-1 | s4-g3 | $0.1745 | $125.64 |
| my08-ng-influxdb-back-2 | s4-g3 | $0.1745 | $125.64 |
| my08-ng-ip-10-0-1-30-62161624-1 | c2-g3 | $0.0677 | $48.74 |


