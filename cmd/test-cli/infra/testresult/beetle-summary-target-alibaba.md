# Target Cloud Infrastructure Summary

**Generated At:** 2026-05-19 05:22:28

**Namespace:** mig01

**Infra Name:** my04-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my04-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | ALIBABA |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| ecs.e-c1m4.large | 2 | 8.0 | - | x86_64 |  | $0.0791 | 1 |
| ecs.e-c1m4.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1582 | 1 |
| ecs.e-c1m1.large | 2 | 2.0 | - | x86_64 |  | $0.0178 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ubuntu_22_04_x64_20G_alibase_20260316.vhd | Ubuntu  22.04 64 bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | i-mj7c2m22qkgbnbe7713p | Running | 2 vCPU, 2.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-vnet-01<br>**Subnet:** my04-subnet-01<br>**Public IP:** 47.80.240.250<br>**Private IP:** 10.0.1.190<br>**SGs:** my04-sg-01<br>**SSH:** my04-sshkey-01 |
| my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | i-mj75hai0hn9bimlmq1s9 | Running | 2 vCPU, 8.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-vnet-01<br>**Subnet:** my04-subnet-01<br>**Public IP:** 47.80.19.178<br>**Private IP:** 10.0.1.192<br>**SGs:** my04-sg-03<br>**SSH:** my04-sshkey-01 |
| my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | i-mj7fxe8c2pijgu7ylnrz | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-vnet-01<br>**Subnet:** my04-subnet-01<br>**Public IP:** 47.80.57.53<br>**Private IP:** 10.0.1.191<br>**SGs:** my04-sg-02<br>**SSH:** my04-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my04-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my04-vnet-01 |
| **CSP VNet ID** | vpc-mj751zvhsvt0aghr7etbc |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | alibaba-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my04-subnet-01 | vsw-mj7nvmifpblqb1zo7bd9p | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my04-sshkey-01 | tbf5horl7idh7mirobfl |  | 5906c2acffce8f791805dfe64a0c76c9 |

### Security Groups

#### Security Group: my04-sg-01

| Property | Value |
|----------|-------|
| **Name** | my04-sg-01 |
| **CSP Security Group ID** | sg-mj76s1u22n45bkbt8j4h |
| **VNet** | my04-vnet-01 |
| **Rule Count** | 14 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | UDP | 9113 | 10.0.0.0/16 |
| inbound | TCP | 9113 | 10.0.0.0/16 |
| inbound | TCP | 8080 | 0.0.0.0/0 |
| inbound | TCP | 443 | 0.0.0.0/0 |
| inbound | TCP | 80 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my04-sg-02

| Property | Value |
|----------|-------|
| **Name** | my04-sg-02 |
| **CSP Security Group ID** | sg-mj75r0e6o8hdal192p4q |
| **VNet** | my04-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | UDP | 9100 | 10.0.0.0/16 |
| inbound | TCP | 9100 | 10.0.0.0/16 |
| inbound | UDP | 32803 | 10.0.0.0/16 |
| inbound | TCP | 32803 | 10.0.0.0/16 |
| inbound | UDP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 20048 | 10.0.0.0/16 |
| inbound | UDP | 111 | 0.0.0.0/0 |
| inbound | TCP | 111 | 0.0.0.0/0 |
| inbound | UDP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my04-sg-03

| Property | Value |
|----------|-------|
| **Name** | my04-sg-03 |
| **CSP Security Group ID** | sg-mj70qfjkj53ki1csq8pf |
| **VNet** | my04-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | UDP | 9104 | 10.0.0.0/16 |
| inbound | TCP | 9104 | 10.0.0.0/16 |
| inbound | UDP | 4444 | 10.0.0.0/16 |
| inbound | TCP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 4568 | 10.0.0.0/16 |
| inbound | TCP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4567 | 10.0.0.0/16 |
| inbound | TCP | 4567 | 10.0.0.0/16 |
| inbound | UDP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.2551 |
| **Per Day** | $6.12 |
| **Per Month (30 days)** | $183.67 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| ALIBABA | ap-northeast-2 | 3 | $0.2551 | $183.67 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | ecs.e-c1m1.large | $0.0178 | $12.82 |
| my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | ecs.e-c1m4.large | $0.0791 | $56.95 |
| my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |


