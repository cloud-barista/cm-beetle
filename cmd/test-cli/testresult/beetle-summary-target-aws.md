# Target Cloud Infrastructure Summary

**Generated At:** 2026-04-29 11:19:39

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
| t3a.small | 2 | 2.0 | - | x86_64 |  | $0.0234 | 1 |
| t3a.large | 2 | 8.0 | - | x86_64 |  | $0.0936 | 1 |
| t3a.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1872 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ami-08a1c21841a4c7a5f | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260410 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | i-051ac9230d0bef1c9 | Running | 2 vCPU, 2.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260410 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260410) | **VNet:** my01-vnet-01<br>**Subnet:** my01-subnet-01<br>**Public IP:** 52.79.76.8<br>**Private IP:** 10.0.1.54<br>**SGs:** my01-sg-01<br>**SSH:** my01-sshkey-01 |
| my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | i-037f076d3e4a7e58c | Running | 2 vCPU, 8.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260410 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260410) | **VNet:** my01-vnet-01<br>**Subnet:** my01-subnet-01<br>**Public IP:** 13.124.154.54<br>**Private IP:** 10.0.1.121<br>**SGs:** my01-sg-03<br>**SSH:** my01-sshkey-01 |
| my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | i-0c804300e25a9e689 | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260410 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260410) | **VNet:** my01-vnet-01<br>**Subnet:** my01-subnet-01<br>**Public IP:** 43.203.222.40<br>**Private IP:** 10.0.1.146<br>**SGs:** my01-sg-02<br>**SSH:** my01-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my01-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my01-vnet-01 |
| **CSP VNet ID** | vpc-0700b9f9aa0f1f4ee |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my01-subnet-01 | subnet-06ade1584a33bb44a | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my01-sshkey-01 | tbd7ouhsp0bhc4mlvtc480 |  | b5:b9:09:d2:f0:15:a8:83:bc:d9:2b:89:9f:53:ef:cb:99:d4:0f:4b |

### Security Groups

#### Security Group: my01-sg-01

| Property | Value |
|----------|-------|
| **Name** | my01-sg-01 |
| **CSP Security Group ID** | sg-007b9eccfcda456c8 |
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
| **CSP Security Group ID** | sg-0fc813b405734c9e2 |
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
| **CSP Security Group ID** | sg-0e188003866e69890 |
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


