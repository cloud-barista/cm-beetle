# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-06-24 09:56:44*

---

## 📊 Migration Summary

**Target Cloud:** IBM

**Target Region:** au-syd

**Namespace:** mig01 | **Infra ID:** my06-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $321.12 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 3 | ✅ Selected | nxf-2x2, bxf-2x8, bxf-4x16 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my06-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my06-vnet-01) |
| 6 | **Security Group** | 3 security groups | ✅ Created | Total 52 rules in 3 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** 02h7_f8178175-8693-4c3c-bb11-296961b82139<br>**Label(sourceMachineId):** vm-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** 02h7_faa52016-7f78-46a7-bd3e-550531ee6879<br>**Label(sourceMachineId):** vm-ec288dd0-c6fa-8a49-2f60 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 |
| 3 | **VM Name:** my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** 02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58<br>**Label(sourceMachineId):** vm-ec2d32b5-98fb-5a96-7913 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 |

---

## ⚙️ VM Specs

**Summary:** 3 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** nxf-2x2<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2 | my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Spec ID:** bxf-2x8<br>**vCPUs:** 2<br>**Memory:** 8.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3 | my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Spec ID:** bxf-4x16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my06-sg-01

**CSP ID:** r026-4a4bcf87-47c6-4e4a-aef9-14c9b0081370 | **VNet:** my06-vnet-01 | **Rules:** 14

**Assigned VMs:**

- **VM:** my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec268ed7-821e-9d73-e79f

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 2 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 3 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 4 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | inbound | TCP | 80 | 0.0.0.0/0 | inbound tcp 80 | Migrated from source |
| 7 | inbound | TCP | 443 | 0.0.0.0/0 | inbound tcp 443 | Migrated from source |
| 8 | inbound | TCP | 8080 | 0.0.0.0/0 | inbound tcp 8080 | Migrated from source |
| 9 | inbound | TCP | 9113 | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 10 | inbound | UDP | 9113 | 10.0.0.0/16 | inbound udp 9113 from 10.0.0.0/16 | Migrated from source |
| 11 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 12 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 13 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 14 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my06-sg-02

**CSP ID:** r026-51dd6208-ab6c-486a-b2c9-b1e68dcf7791 | **VNet:** my06-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec2d32b5-98fb-5a96-7913

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 2 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 3 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 4 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | inbound | TCP | 2049 | 0.0.0.0/0 | inbound tcp 2049 | Migrated from source |
| 7 | inbound | UDP | 2049 | 0.0.0.0/0 | inbound udp 2049 | Migrated from source |
| 8 | inbound | TCP | 111 | 0.0.0.0/0 | inbound tcp 111 | Migrated from source |
| 9 | inbound | UDP | 111 | 0.0.0.0/0 | inbound udp 111 | Migrated from source |
| 10 | inbound | TCP | 20048 | 10.0.0.0/16 | inbound tcp 20048 from 10.0.0.0/16 | Migrated from source |
| 11 | inbound | UDP | 20048 | 10.0.0.0/16 | inbound udp 20048 from 10.0.0.0/16 | Migrated from source |
| 12 | inbound | TCP | 32803 | 10.0.0.0/16 | inbound tcp 32803 from 10.0.0.0/16 | Migrated from source |
| 13 | inbound | UDP | 32803 | 10.0.0.0/16 | inbound udp 32803 from 10.0.0.0/16 | Migrated from source |
| 14 | inbound | TCP | 9100 | 10.0.0.0/16 | inbound tcp 9100 from 10.0.0.0/16 | Migrated from source |
| 15 | inbound | UDP | 9100 | 10.0.0.0/16 | inbound udp 9100 from 10.0.0.0/16 | Migrated from source |
| 16 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 17 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 18 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 19 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my06-sg-03

**CSP ID:** r026-fca6d682-9c93-4321-8801-d4436cf67c14 | **VNet:** my06-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec288dd0-c6fa-8a49-2f60

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 2 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 3 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 4 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | inbound | TCP | 3306 | 10.0.0.0/16 | inbound tcp 3306 from 10.0.0.0/16 | Migrated from source |
| 7 | inbound | UDP | 3306 | 10.0.0.0/16 | inbound udp 3306 from 10.0.0.0/16 | Migrated from source |
| 8 | inbound | TCP | 4567 | 10.0.0.0/16 | inbound tcp 4567 from 10.0.0.0/16 | Migrated from source |
| 9 | inbound | UDP | 4567 | 10.0.0.0/16 | inbound udp 4567 from 10.0.0.0/16 | Migrated from source |
| 10 | inbound | TCP | 4568 | 10.0.0.0/16 | inbound tcp 4568 from 10.0.0.0/16 | Migrated from source |
| 11 | inbound | UDP | 4568 | 10.0.0.0/16 | inbound udp 4568 from 10.0.0.0/16 | Migrated from source |
| 12 | inbound | TCP | 4444 | 10.0.0.0/16 | inbound tcp 4444 from 10.0.0.0/16 | Migrated from source |
| 13 | inbound | UDP | 4444 | 10.0.0.0/16 | inbound udp 4444 from 10.0.0.0/16 | Migrated from source |
| 14 | inbound | TCP | 9104 | 10.0.0.0/16 | inbound tcp 9104 from 10.0.0.0/16 | Migrated from source |
| 15 | inbound | UDP | 9104 | 10.0.0.0/16 | inbound udp 9104 from 10.0.0.0/16 | Migrated from source |
| 16 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 17 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 18 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 19 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source node network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** my06-vnet-01<br>**ID:** r026-4650a745-f71b-4cfc-9c2f-bb503634ac18 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my06-subnet-01<br>**ID:** 02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179 | 10.0.1.0/24 | my06-vnet-01 |

### Source Network Information

**CIDR:** 10.0.1.0/24 | **Gateway:** 10.0.1.1 | **Connected Servers:** 3

### Network Details by Server (3 servers)

#### 1. ip-10-0-1-30

**Active Interfaces:**

| Interface | IP Address | State |
|-----------|------------|-------|
| lo | 127.0.0.1/8 | up |

#### 2. ip-10-0-1-221

**Active Interfaces:**

| Interface | IP Address | State |
|-----------|------------|-------|
| lo | 127.0.0.1/8 | up |

#### 3. ip-10-0-1-138

**Active Interfaces:**

| Interface | IP Address | State |
|-----------|------------|-------|
| lo | 127.0.0.1/8 | up |

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my06-sshkey-01 | r026-f149af7c-f2ec-41c2-937f-f4f592fe4602 | SHA256:hpm7J8C7RXdoPYHRNRKt9tB0B04Zuncu03GzqKEshzM | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.4460 |
| Daily | $10.70 |
| Monthly | $321.12 |
| Yearly | $3853.44 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | nxf-2x2 | $67.68 | 21.1% |
| ip-10-0-1-221 (migrated) | bxf-4x16 | $169.20 | 52.7% |
| ip-10-0-1-138 (migrated) | bxf-2x8 | $84.24 | 26.2% |

---


---

*Report generated by CM-Beetle*
