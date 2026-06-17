# Target Cloud Infrastructure Summary

**Generated At:** 2026-06-17 09:50:44

**Namespace:** mig01

**Infra Name:** my01-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my01-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | AWS |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| t3a.large | 2 | 8.0 | - | x86_64 |  | $0.0936 | 1 |
| t3a.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1872 | 1 |
| t3a.small | 2 | 2.0 | - | x86_64 |  | $0.0234 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ami-09a72717a566d88fa | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260602 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | i-0f6ffab8087fecf9c | Running | 2 vCPU, 2.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260602 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260602) | **VNet:** my01-vnet-01<br>**Subnet:** my01-subnet-01<br>**Public IP:** 3.35.138.48<br>**Private IP:** 10.0.1.193<br>**SGs:** my01-sg-01<br>**SSH:** my01-sshkey-01 |
| my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | i-0d54437998cb267cd | Running | 2 vCPU, 8.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260602 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260602) | **VNet:** my01-vnet-01<br>**Subnet:** my01-subnet-01<br>**Public IP:** 13.124.69.82<br>**Private IP:** 10.0.1.9<br>**SGs:** my01-sg-03<br>**SSH:** my01-sshkey-01 |
| my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | i-0652cd78bfd506b86 | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260602 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260602) | **VNet:** my01-vnet-01<br>**Subnet:** my01-subnet-01<br>**Public IP:** 43.203.249.33<br>**Private IP:** 10.0.1.41<br>**SGs:** my01-sg-02<br>**SSH:** my01-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my01-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my01-vnet-01 |
| **CSP VNet ID** | vpc-091b19f48c70a6558 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my01-subnet-01 | subnet-04b1ae9b4fa56d880 | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my01-sshkey-01 | tbc4esm8m35ft1plmok6 |  | 76:48:24:f6:32:16:21:18:75:da:1e:b5:76:9e:8d:dc:ee:1c:30:15 |

### Security Groups

#### Security Group: my01-sg-01

| Property | Value |
|----------|-------|
| **Name** | my01-sg-01 |
| **CSP Security Group ID** | sg-0822b166884ea3f10 |
| **VNet** | my01-vnet-01 |
| **Rule Count** | 14 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 80 | 0.0.0.0/0 |
| inbound | TCP | 8080 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | UDP | 9113 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 9113 | 10.0.0.0/16 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | TCP | 443 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my01-sg-02

| Property | Value |
|----------|-------|
| **Name** | my01-sg-02 |
| **CSP Security Group ID** | sg-077b765d4b7eb2a1c |
| **VNet** | my01-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | UDP | 9100 | 10.0.0.0/16 |
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 32803 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 111 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 32803 | 10.0.0.0/16 |
| inbound | TCP | 9100 | 10.0.0.0/16 |
| inbound | UDP | 111 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my01-sg-03

| Property | Value |
|----------|-------|
| **Name** | my01-sg-03 |
| **CSP Security Group ID** | sg-0c3570f091852d364 |
| **VNet** | my01-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | UDP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 9104 | 10.0.0.0/16 |
| inbound | TCP | 9104 | 10.0.0.0/16 |
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 3306 | 10.0.0.0/16 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4567 | 10.0.0.0/16 |
| inbound | TCP | 4567 | 10.0.0.0/16 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | TCP | 3306 | 10.0.0.0/16 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.3042 |
| **Per Day** | $7.30 |
| **Per Month (30 days)** | $219.02 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AWS | ap-northeast-2 | 3 | $0.3042 | $219.02 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | t3a.small | $0.0234 | $16.85 |
| my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | t3a.large | $0.0936 | $67.39 |
| my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | t3a.xlarge | $0.1872 | $134.78 |


