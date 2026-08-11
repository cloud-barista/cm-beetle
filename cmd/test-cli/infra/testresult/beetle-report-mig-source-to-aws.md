# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-08-11 04:11:23*

---

## 📊 Migration Summary

**Target Cloud:** AWS

**Target Region:** ap-northeast-2

**Namespace:** mig01 | **Infra ID:** my-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $219.02 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 3 | ✅ Selected | t3a.small, t3a.large, t3a.xlarge |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my-vnet-01) |
| 6 | **Security Group** | 3 security groups | ✅ Created | Total 52 rules in 3 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** i-0ab2bb9a9968a2e8f<br>**Label(sourceMachineId):** ec268ed7-821e-9d73-e79f-961262161624 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 |
| 2 | **VM Name:** my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** i-05d74744926086bcc<br>**Label(sourceMachineId):** ec288dd0-c6fa-8a49-2f60-bc898311febf | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf |
| 3 | **VM Name:** my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** i-015fb4dc098de22ea<br>**Label(sourceMachineId):** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 |

---

## ⚙️ VM Specs

**Summary:** 3 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** t3a.small<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 0 GB |
| 2 | my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Spec ID:** t3a.large<br>**vCPUs:** 2<br>**Memory:** 8.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 8 GB<br>**Root Disk:** 0 GB |
| 3 | my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Spec ID:** t3a.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 0 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** ami-0afe1fd15675c3f15<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** ami-0afe1fd15675c3f15<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** ami-0afe1fd15675c3f15<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my-sg-01

**CSP ID:** sg-05ef7aea1f88dfc72 | **VNet:** my-vnet-01 | **Rules:** 14

**Assigned VMs:**

- **VM:** my-vm-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** ip-10-0-1-30, **Machine ID:** ec268ed7-821e-9d73-e79f-961262161624

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 80 | 0.0.0.0/0 | inbound tcp 80 | Migrated from source |
| 2 | inbound | TCP | 8080 | 0.0.0.0/0 | inbound tcp 8080 | Migrated from source |
| 3 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | UDP | 9113 | 10.0.0.0/16 | inbound udp 9113 from 10.0.0.0/16 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 7 | inbound | TCP | 9113 | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 8 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 9 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 10 | inbound | TCP | 443 | 0.0.0.0/0 | inbound tcp 443 | Migrated from source |
| 11 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 12 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |
| 13 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 14 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |

### Security Group: my-sg-02

**CSP ID:** sg-079200a3354909306 | **VNet:** my-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1
  - **Source Server:** **Hostname:** ip-10-0-1-221, **Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | UDP | 9100 | 10.0.0.0/16 | inbound udp 9100 from 10.0.0.0/16 | Migrated from source |
| 2 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 3 | inbound | UDP | 20048 | 10.0.0.0/16 | inbound udp 20048 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | TCP | 2049 | 0.0.0.0/0 | inbound tcp 2049 | Migrated from source |
| 5 | inbound | TCP | 32803 | 10.0.0.0/16 | inbound tcp 32803 from 10.0.0.0/16 | Migrated from source |
| 6 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 7 | inbound | TCP | 20048 | 10.0.0.0/16 | inbound tcp 20048 from 10.0.0.0/16 | Migrated from source |
| 8 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 9 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 10 | inbound | UDP | 2049 | 0.0.0.0/0 | inbound udp 2049 | Migrated from source |
| 11 | inbound | TCP | 111 | 0.0.0.0/0 | inbound tcp 111 | Migrated from source |
| 12 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 13 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 14 | inbound | UDP | 32803 | 10.0.0.0/16 | inbound udp 32803 from 10.0.0.0/16 | Migrated from source |
| 15 | inbound | TCP | 9100 | 10.0.0.0/16 | inbound tcp 9100 from 10.0.0.0/16 | Migrated from source |
| 16 | inbound | UDP | 111 | 0.0.0.0/0 | inbound udp 111 | Migrated from source |
| 17 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |
| 18 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 19 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |

### Security Group: my-sg-03

**CSP ID:** sg-0628a002e31e30412 | **VNet:** my-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1
  - **Source Server:** **Hostname:** ip-10-0-1-138, **Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | UDP | 4568 | 10.0.0.0/16 | inbound udp 4568 from 10.0.0.0/16 | Migrated from source |
| 2 | inbound | UDP | 4444 | 10.0.0.0/16 | inbound udp 4444 from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | UDP | 9104 | 10.0.0.0/16 | inbound udp 9104 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | TCP | 9104 | 10.0.0.0/16 | inbound tcp 9104 from 10.0.0.0/16 | Migrated from source |
| 5 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 6 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 7 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 8 | inbound | TCP | 4444 | 10.0.0.0/16 | inbound tcp 4444 from 10.0.0.0/16 | Migrated from source |
| 9 | inbound | UDP | 3306 | 10.0.0.0/16 | inbound udp 3306 from 10.0.0.0/16 | Migrated from source |
| 10 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 11 | inbound | TCP | 4568 | 10.0.0.0/16 | inbound tcp 4568 from 10.0.0.0/16 | Migrated from source |
| 12 | inbound | UDP | 4567 | 10.0.0.0/16 | inbound udp 4567 from 10.0.0.0/16 | Migrated from source |
| 13 | inbound | TCP | 4567 | 10.0.0.0/16 | inbound tcp 4567 from 10.0.0.0/16 | Migrated from source |
| 14 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 15 | inbound | TCP | 3306 | 10.0.0.0/16 | inbound tcp 3306 from 10.0.0.0/16 | Migrated from source |
| 16 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 17 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |
| 18 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 19 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source node network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** my-vnet-01<br>**ID:** vpc-031a5f92d193a4999 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my-subnet-01<br>**ID:** subnet-0b4cc6639f869e6f1 | 10.0.1.0/24 | my-vnet-01 |

### Source Network Information

**CIDR:** 10.0.1.0/24 | **Gateway:** 10.0.1.1 | **Connected Servers:** 3

### Network Details by Server (3 servers)

#### 1. ip-10-0-1-30

**Active Interfaces:**

| Interface | IP Address | State |
|-----------|------------|-------|
| lo | 127.0.0.1/8 | up |
| ens5 | 10.0.1.30/24 | up |

**Main Routes:**

| Destination | Gateway | Interface |
|-------------|---------|-----------|
| 0.0.0.0/0 | 10.0.1.1 | ens5 |
| 10.0.0.2/32 | 10.0.1.1 | ens5 |
| 10.0.1.0/24 | 10.0.1.1 | ens5 |
| 10.0.1.1/32 | 10.0.1.1 | ens5 |

#### 2. ip-10-0-1-221

**Active Interfaces:**

| Interface | IP Address | State |
|-----------|------------|-------|
| lo | 127.0.0.1/8 | up |
| ens5 | 10.0.1.221/24 | up |

**Main Routes:**

| Destination | Gateway | Interface |
|-------------|---------|-----------|
| 0.0.0.0/0 | 10.0.1.1 | ens5 |
| 10.0.0.2/32 | 10.0.1.1 | ens5 |
| 10.0.1.0/24 | 10.0.1.1 | ens5 |
| 10.0.1.1/32 | 10.0.1.1 | ens5 |

#### 3. ip-10-0-1-138

**Active Interfaces:**

| Interface | IP Address | State |
|-----------|------------|-------|
| lo | 127.0.0.1/8 | up |
| ens5 | 10.0.1.138/24 | up |

**Main Routes:**

| Destination | Gateway | Interface |
|-------------|---------|-----------|
| 0.0.0.0/0 | 10.0.1.1 | ens5 |
| 10.0.0.2/32 | 10.0.1.1 | ens5 |
| 10.0.1.0/24 | 10.0.1.1 | ens5 |
| 10.0.1.1/32 | 10.0.1.1 | ens5 |

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my-sshkey-01 | tbnr4ecavk041qi7s4a6 | 3d:2a:0f:ac:8e:d7:66:de:d2:da:e6:be:04:f2:b0:58:6a:d0:12:74 | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.3042 |
| Daily | $7.30 |
| Monthly | $219.02 |
| Yearly | $2628.29 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | t3a.small | $16.85 | 7.7% |
| ip-10-0-1-221 (migrated) | t3a.xlarge | $134.78 | 61.5% |
| ip-10-0-1-138 (migrated) | t3a.large | $67.39 | 30.8% |

---


---

*Report generated by CM-Beetle*
