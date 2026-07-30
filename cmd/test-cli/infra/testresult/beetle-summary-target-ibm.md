# Target Cloud Infrastructure Summary

**Generated At:** 2026-07-30 01:19:37

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
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
| bxf-2x8 | 2 | 8.0 | - | x86_64 |  | $0.1170 | 1 |
| bxf-4x16 | 4 | 16.0 | - | x86_64 |  | $0.2350 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599 | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | 02h7_7cb7d68f-a4a5-4166-8512-df4115f4c450 | Running | 2 vCPU, 2.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 159.23.93.105<br>**Private IP:** 10.0.1.5<br>**SGs:** my-sg-01<br>**SSH:** my-sshkey-01 |
| my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | 02h7_67e33f45-6a8f-48ab-9096-8754e4d44c56 | Running | 2 vCPU, 8.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 159.23.91.219<br>**Private IP:** 10.0.1.6<br>**SGs:** my-sg-03<br>**SSH:** my-sshkey-01 |
| my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | 02h7_2b2e2194-5a77-4d14-8f9d-b62ef5e1f0bc | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 159.23.102.119<br>**Private IP:** 10.0.1.4<br>**SGs:** my-sg-02<br>**SSH:** my-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-vnet-01 |
| **CSP VNet ID** | r026-417f5571-d0c7-487c-b354-42b85989f3a0 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-subnet-01 | 02h7-d8cffe55-1735-4724-819e-5c4a4cd0766f | 10.0.1.0/24 | au-syd-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-sshkey-01 | r026-6d8513e1-f27f-419b-b227-4038896760f2 |  | SHA256:h33RymjkE3ER05uVqIxxZk0OFimOkNu9hMBcScOqD8Q |

### Security Groups

#### Security Group: my-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-sg-01 |
| **CSP Security Group ID** | r026-b6cdb1e3-8f19-4e7f-acd3-2a6d9b0a72bb |
| **VNet** | my-vnet-01 |
| **Rule Count** | 14 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 80 | 0.0.0.0/0 |
| inbound | TCP | 443 | 0.0.0.0/0 |
| inbound | TCP | 8080 | 0.0.0.0/0 |
| inbound | TCP | 9113 | 10.0.0.0/16 |
| inbound | UDP | 9113 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-sg-02 |
| **CSP Security Group ID** | r026-9f6b756c-c57f-4d24-86be-32633c197efe |
| **VNet** | my-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 2049 | 0.0.0.0/0 |
| inbound | UDP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 111 | 0.0.0.0/0 |
| inbound | UDP | 111 | 0.0.0.0/0 |
| inbound | TCP | 20048 | 10.0.0.0/16 |
| inbound | UDP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 32803 | 10.0.0.0/16 |
| inbound | UDP | 32803 | 10.0.0.0/16 |
| inbound | TCP | 9100 | 10.0.0.0/16 |
| inbound | UDP | 9100 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-sg-03

| Property | Value |
|----------|-------|
| **Name** | my-sg-03 |
| **CSP Security Group ID** | r026-6920486f-c8de-48ac-bf7c-ce35b1f4bde5 |
| **VNet** | my-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 3306 | 10.0.0.0/16 |
| inbound | UDP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 4567 | 10.0.0.0/16 |
| inbound | UDP | 4567 | 10.0.0.0/16 |
| inbound | TCP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4568 | 10.0.0.0/16 |
| inbound | TCP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 4444 | 10.0.0.0/16 |
| inbound | TCP | 9104 | 10.0.0.0/16 |
| inbound | UDP | 9104 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.4460 |
| **Per Day** | $10.70 |
| **Per Month (30 days)** | $321.12 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| IBM | au-syd | 3 | $0.4460 | $321.12 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | nxf-2x2 | $0.0940 | $67.68 |
| my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | bxf-2x8 | $0.1170 | $84.24 |
| my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | bxf-4x16 | $0.2350 | $169.20 |


