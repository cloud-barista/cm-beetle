# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:18:52

**Namespace:** mig01

**Infra Name:** my04-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my04-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | ALIBABA |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| ecs.e-c1m4.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1582 | 2 |
| ecs.e-c1m1.large | 2 | 2.0 | - | x86_64 |  | $0.0178 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ubuntu_22_04_x64_20G_alibase_20260916.vhd | Ubuntu  22.04 64 bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my04-ng-influxdb-back-1 | i-mj7f24f44ih73lbcha5e | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 8.213.148.200<br>**Private IP:** 10.0.1.3<br>**SGs:** my04-mig-sg-01<br>**SSH:** my04-mig-sshkey-01 |
| my04-ng-influxdb-back-2 | i-mj7803i8w5ifq7vxuxuc | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 47.80.61.166<br>**Private IP:** 10.0.1.2<br>**SGs:** my04-mig-sg-01<br>**SSH:** my04-mig-sshkey-01 |
| my04-ng-ip-10-0-1-30-62161624-1 | i-mj7ja9zhp5zysim4twue | Running | 2 vCPU, 2.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 8.220.218.42<br>**Private IP:** 10.0.1.1<br>**SGs:** my04-mig-sg-02<br>**SSH:** my04-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my04-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my04-mig-vnet-01 |
| **CSP VNet ID** | vpc-mj7abluu7bv81vyj6k43j |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | alibaba-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my04-mig-subnet-01 | vsw-mj7ipqju3acsuca2qfa7e | 10.0.1.0/24 | ap-northeast-2a |

### Network Load Balancers (NLB)

#### NLB: my04-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my04-ng-influxdb-back |
| **CSP NLB Name** | tb084ephbj3j3ao7c7m4 |
| **CSP NLB ID** | lb-mj70hxl0s9hyk87qunerv |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | alibaba-ap-northeast-2 |
| **Listener** | TCP:9999 (IP: 43.108.8.225) |
| **Target Group** | TCP:8086 (Node Group: my04-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my04-ng-influxdb-back-1 | i-mj7f24f44ih73lbcha5e | 10.0.1.3 | 8.213.148.200 | Running |
| my04-ng-influxdb-back-2 | i-mj7803i8w5ifq7vxuxuc | 10.0.1.2 | 47.80.61.166 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my04-mig-sshkey-01 | tbskggpajj56tq2usk2e |  | ab6135bc1f579a0742563628dc47d647 |

### Security Groups

#### Security Group: my04-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my04-mig-sg-01 |
| **CSP Security Group ID** | sg-mj73k9coy7sks4yjzisr |
| **VNet** | my04-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my04-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my04-mig-sg-02 |
| **CSP Security Group ID** | sg-mj77a670vw20ufh90ntg |
| **VNet** | my04-mig-vnet-01 |
| **Rule Count** | 4 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.3342 |
| **Per Day** | $8.02 |
| **Per Month (30 days)** | $240.62 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| ALIBABA | ap-northeast-2 | 3 | $0.3342 | $240.62 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my04-ng-influxdb-back-1 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |
| my04-ng-influxdb-back-2 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |
| my04-ng-ip-10-0-1-30-62161624-1 | ecs.e-c1m1.large | $0.0178 | $12.82 |


