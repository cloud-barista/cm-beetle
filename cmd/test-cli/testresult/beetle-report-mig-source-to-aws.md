# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2025-11-17 04:48:17*

---

## 📊 Migration Summary

**Target Cloud:** AWS

**Target Region:** ap-northeast-2

**Namespace:** mig01 | **MCI ID:** mmci01

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
| 2 | **VM Spec** | 3 | ✅ Selected | t3a.xlarge, t3a.small, t3a.large |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in mig-vnet-01) |
| 6 | **Security Group** | 3 security groups | ✅ Created | Total 52 rules in 3 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source server(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1<br>**VM ID:** i-0dfa1890fcefba4be<br>**Label(sourceMachineId):** ec21fd51-16bb-7e10-5e23-12ef283b2204 | **Hostname:** ip-10-0-1-85<br>**Machine ID:** ec21fd51-16bb-7e10-5e23-12ef283b2204 |
| 2 | **VM Name:** migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1<br>**VM ID:** i-00c25957c7fdb5cfc<br>**Label(sourceMachineId):** ec2643f0-9388-3d97-f3a4-f387cd52696c | **Hostname:** ip-10-0-1-66<br>**Machine ID:** ec2643f0-9388-3d97-f3a4-f387cd52696c |
| 3 | **VM Name:** migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1<br>**VM ID:** i-0a5882b217e54f74f<br>**Label(sourceMachineId):** ec2876a6-3c84-7e62-aaf9-c3203f12e0b8 | **Hostname:** ip-10-0-1-9<br>**Machine ID:** ec2876a6-3c84-7e62-aaf9-c3203f12e0b8 |

---

## ⚙️ VM Specs

**Summary:** 3 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1 | **Spec ID:** t3a.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-85<br>**Machine ID:** ec21fd51-16bb-7e10-5e23-12ef283b2204 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 0 GB |
| 2 | migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1 | **Spec ID:** t3a.small<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-66<br>**Machine ID:** ec2643f0-9388-3d97-f3a4-f387cd52696c | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 0 GB |
| 3 | migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1 | **Spec ID:** t3a.large<br>**vCPUs:** 2<br>**Memory:** 8.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-9<br>**Machine ID:** ec2876a6-3c84-7e62-aaf9-c3203f12e0b8 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 8 GB<br>**Root Disk:** 0 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1 | **Image ID:** ami-010be25c3775061c9<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015 | **Hostname:** ip-10-0-1-85<br>**Machine ID:** ec21fd51-16bb-7e10-5e23-12ef283b2204 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1 | **Image ID:** ami-010be25c3775061c9<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015 | **Hostname:** ip-10-0-1-66<br>**Machine ID:** ec2643f0-9388-3d97-f3a4-f387cd52696c | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1 | **Image ID:** ami-010be25c3775061c9<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015 | **Hostname:** ip-10-0-1-9<br>**Machine ID:** ec2876a6-3c84-7e62-aaf9-c3203f12e0b8 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: mig-sg-01

**CSP ID:** sg-08d51712b18745856 | **VNet:** mig-vnet-01 | **Rules:** 14

**Assigned VMs:**

- **VM:** migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1
  - **Source Server:** **Hostname:** ip-10-0-1-66, **Machine ID:** ec2643f0-9388-3d97-f3a4-f387cd52696c

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

### Security Group: mig-sg-02

**CSP ID:** sg-0893e1a7e93bcf006 | **VNet:** mig-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1
  - **Source Server:** **Hostname:** ip-10-0-1-85, **Machine ID:** ec21fd51-16bb-7e10-5e23-12ef283b2204

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

### Security Group: mig-sg-03

**CSP ID:** sg-023e0f0b2c789b8f8 | **VNet:** mig-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1
  - **Source Server:** **Hostname:** ip-10-0-1-9, **Machine ID:** ec2876a6-3c84-7e62-aaf9-c3203f12e0b8

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

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source server network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** mig-vnet-01<br>**ID:** vpc-0330b15879aede8d2 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** mig-subnet-01<br>**ID:** subnet-00440b2d1ddf6b379 | 10.0.1.0/24 | mig-vnet-01 |

### Source Network Information

**CIDR:** 10.0.1.0/24 | **Gateway:** 10.0.1.1 | **Connected Servers:** 3

### Network Details by Server (3 servers)

#### 1. ip-10-0-1-66

**Active Interfaces:**

| Interface | IP Address | State |
|-----------|------------|-------|
| lo | 127.0.0.1/8 | up |

#### 2. ip-10-0-1-85

**Active Interfaces:**

| Interface | IP Address | State |
|-----------|------------|-------|
| lo | 127.0.0.1/8 | up |

#### 3. ip-10-0-1-9

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
| 1 | mig-sshkey-01 | d4dahgap2foc73bcejd0 | d5:4b:aa:0f:d4:eb:7d:20:44:88:f0:5f:fc:e4:c8:1a:ee:9b:82:34 | Used by all 3 VMs |

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
| ip-10-0-1-66 (migrated) | t3a.small | $16.85 | 7.7% |
| ip-10-0-1-85 (migrated) | t3a.xlarge | $134.78 | 61.5% |
| ip-10-0-1-9 (migrated) | t3a.large | $67.39 | 30.8% |

---


---

*Report generated by CM-Beetle*
