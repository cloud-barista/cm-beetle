# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-03 07:46:54

**Namespace:** mig01

**Infra Name:** my01-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my01-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | AWS |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Total NLBs** | 1 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| t3.small | 2 | 2.0 | - | x86_64 |  | $0.0260 | 1 |
| t3.xlarge | 4 | 16.0 | - | x86_64 |  | $0.2080 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ami-046333f9ecafdf406 | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | i-0be90f3aabdd99191 | Running | 2 vCPU, 2.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 43.203.237.20<br>**Private IP:** 10.0.1.74<br>**SGs:** my01-mig-sg-02<br>**SSH:** my01-mig-sshkey-01 |
| my01-ng-influxdb-back-1 | i-017eb01b702b8df89 | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 3.34.142.234<br>**Private IP:** 10.0.1.24<br>**SGs:** my01-mig-sg-01<br>**SSH:** my01-mig-sshkey-01 |
| my01-ng-influxdb-back-2 | i-0eca101eb0ca896bc | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 52.78.237.145<br>**Private IP:** 10.0.1.65<br>**SGs:** my01-mig-sg-01<br>**SSH:** my01-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my01-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my01-mig-vnet-01 |
| **CSP VNet ID** | vpc-0530bd0f6f0f28bf1 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my01-mig-subnet-01 | subnet-0f0d9639c2fdd86cf | 10.0.1.0/24 | ap-northeast-2a |

### Network Load Balancers (NLB)

#### NLB: my01-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my01-ng-influxdb-back |
| **CSP NLB Name** | tbfrdc4bp48s53urlj0a |
| **CSP NLB ID** | arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbfrdc4bp48s53urlj0a/f9a14c3ef763e5c9 |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | aws-ap-northeast-2 |
| **Listener** | TCP:9999 (DNS: tbfrdc4bp48s53urlj0a-f9a14c3ef763e5c9.elb.ap-northeast-2.amazonaws.com) |
| **Target Group** | TCP:8086 (Node Group: my01-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my01-ng-influxdb-back-1 | i-017eb01b702b8df89 | 10.0.1.24 | 3.34.142.234 | Running |
| my01-ng-influxdb-back-2 | i-0eca101eb0ca896bc | 10.0.1.65 | 52.78.237.145 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my01-mig-sshkey-01 | tbisjldj8j6botepff4r |  | 31:bc:1b:74:16:c9:85:3c:c7:ce:c0:dc:92:73:d7:47:21:1f:49:d8 |

### Security Groups

#### Security Group: my01-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my01-mig-sg-01 |
| **CSP Security Group ID** | sg-08a1b09b202034ecb |
| **VNet** | my01-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my01-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my01-mig-sg-02 |
| **CSP Security Group ID** | sg-0ba080c26634f6c47 |
| **VNet** | my01-mig-vnet-01 |
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
| **Per Hour** | $0.4420 |
| **Per Day** | $10.61 |
| **Per Month (30 days)** | $318.24 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AWS | ap-northeast-2 | 3 | $0.4420 | $318.24 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | t3.small | $0.0260 | $18.72 |
| my01-ng-influxdb-back-1 | t3.xlarge | $0.2080 | $149.76 |
| my01-ng-influxdb-back-2 | t3.xlarge | $0.2080 | $149.76 |


