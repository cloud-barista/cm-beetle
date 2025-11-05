# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2025-11-05 04:36:10*

---

## 📊 Migration Summary

**Target Cloud:** ALIBABA

**Target Region:** ap-northeast-2

**Namespace:** mig01 | **MCI ID:** mmci01

**Migration Status:** Completed

**Total Servers:** 2

**Migrated Servers:** 2

**💰 Estimated Monthly Cost:** $365.60 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 2 | ✅ Created | 2 running, 2 total |
| 2 | **VM Spec** | 2 | ✅ Selected | ecs.t6-c1m4.2xlarge, ecs.t6-c1m4.xlarge |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | mig-vnet-01, CIDR: 192.168.96.0/19 |
| 5 | **Subnet** | 1 | ✅ Created | 192.168.110.0/24 (in mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 8 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 2 VM(s) have been successfully created in the target cloud, migrated from 2 source server(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1<br>**VM ID:** i-mj7bz4ymj6qw6yvwsl5f<br>**Label(sourceMachineId):** 0036e4b9-c8b4-e811-906e-000ffee02d5c | **Hostname:** cm-web<br>**Machine ID:** 0036e4b9-c8b4-e811-906e-000ffee02d5c |
| 2 | **VM Name:** migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1<br>**VM ID:** i-mj7ajjr4yvgvpqlx0vts<br>**Label(sourceMachineId):** 00a9f3d4-74b6-e811-906e-000ffee02d5c | **Hostname:** cm-nfs<br>**Machine ID:** 00a9f3d4-74b6-e811-906e-000ffee02d5c |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1 | **Spec ID:** ecs.t6-c1m4.2xlarge<br>**vCPUs:** 8<br>**Memory:** 32.0 GB<br>**Root Disk:** 50 GB | **Hostname:** cm-web<br>**Machine ID:** 0036e4b9-c8b4-e811-906e-000ffee02d5c | **CPUs:** 1<br>**Threads:** 8<br>**Memory:** 16 GB<br>**Root Disk:** 1312 GB |
| 2 | migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1 | **Spec ID:** ecs.t6-c1m4.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** cm-nfs<br>**Machine ID:** 00a9f3d4-74b6-e811-906e-000ffee02d5c | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 1093 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20250917.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** cm-web<br>**Machine ID:** 0036e4b9-c8b4-e811-906e-000ffee02d5c | **PrettyName:** Ubuntu 22.04.5 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.5 LTS (Jammy Jellyfish) |
| 2 | migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20250917.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** cm-nfs<br>**Machine ID:** 00a9f3d4-74b6-e811-906e-000ffee02d5c | **PrettyName:** Ubuntu 22.04.5 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.5 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 8 security rule(s) have been created and configured for the migrated VMs.

### Security Group: mig-sg-01

**CSP ID:** sg-mj7ajjr4yvgvpqlsnc9t | **VNet:** mig-vnet-01 | **Rules:** 6

**Assigned VMs:**

- **VM:** migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1
  - **Source Server:** **Hostname:** cm-nfs, **Machine ID:** 00a9f3d4-74b6-e811-906e-000ffee02d5c

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 22 | 0.0.0.0/0 | - | Created by system |
| 2 | inbound | UDP | 53 | 192.168.110.0/24 | inbound udp 53 from 192.168.110.0/24 | Migrated from source |
| 3 | inbound | TCP | 8082 | 0.0.0.0/0 | inbound tcp 8081,8082 | Migrated from source |
| 4 | inbound | TCP | 8081 | 0.0.0.0/0 | inbound tcp 8081,8082 | Migrated from source |
| 5 | inbound | TCP | 10022 | 0.0.0.0/0 | inbound tcp 10022 | Migrated from source |
| 6 | outbound | ALL |  | 0.0.0.0/0 | - | Created by system |

### Security Group: mig-sg-02

**CSP ID:** sg-mj7c90xduxac8di2xg53 | **VNet:** mig-vnet-01 | **Rules:** 2

**Assigned VMs:**

- **VM:** migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1
  - **Source Server:** **Hostname:** cm-web, **Machine ID:** 0036e4b9-c8b4-e811-906e-000ffee02d5c

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 22 | 0.0.0.0/0 | - | Created by system |
| 2 | outbound | ALL |  | 0.0.0.0/0 | - | Created by system |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source server network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** mig-vnet-01<br>**ID:** vpc-mj7qgs1luk3szocsy5uqz | 192.168.96.0/19 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** mig-subnet-01<br>**ID:** vsw-mj74458hhp0qq26mcaaed | 192.168.110.0/24 | mig-vnet-01 |

### Source Network Information

**CIDR:** 192.168.110.0/24 | **Gateway:** 192.168.110.254 | **Connected Servers:** 2

### Network Details by Server (2 servers)

#### 1. cm-nfs

**Active Interfaces:**

| Interface | IP Address | State |
|-----------|------------|-------|
| lo | 127.0.0.1/8 | up |
| eno1np0 | 172.29.0.102/24, 172.29.0.200/32 | up |
| eno2np1 | 192.168.110.200/32 | up |
| br-189b10762332 | 172.20.0.1/16 | down |
| br-f67138586d47 | 172.19.0.1/16 | down |
| br-068801a3f047 | 172.17.0.1/16 | up |
| br-ex | 192.168.110.102/24 | up |

**Main Routes:**

| Destination | Gateway | Interface |
|-------------|---------||-----------|
| 0.0.0.0/0 | 192.168.110.254 | br-ex |
| 172.17.0.0/16 | on-link | br-068801a3f047 |
| 172.19.0.0/16 | on-link | br-f67138586d47 |
| 172.20.0.0/16 | on-link | br-189b10762332 |
| 172.29.0.0/24 | on-link | eno1np0 |
| 192.168.110.0/24 | 192.168.110.254 | br-ex |

#### 2. cm-web

**Active Interfaces:**

| Interface | IP Address | State |
|-----------|------------|-------|
| lo | 127.0.0.1/8 | up |
| eno1np0 | 172.29.0.103/24 | up |
| eno2np1 | - | up |
| br-ex | 192.168.110.103/24 | up |

**Main Routes:**

| Destination | Gateway | Interface |
|-------------|---------||-----------|
| 0.0.0.0/0 | 192.168.110.254 | br-ex |
| 172.29.0.0/24 | on-link | eno1np0 |
| 192.168.110.0/24 | 192.168.110.254 | br-ex |

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | mig-sshkey-01 | d45d7reqjs728podgjhg | b58b791ee2aa2cbe49fd397c905eae08 | Used by all 2 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.5078 |
| Daily | $12.19 |
| Monthly | $365.60 |
| Yearly | $4387.22 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| cm-nfs (migrated) | ecs.t6-c1m4.xlarge | $121.87 | 33.3% |
| cm-web (migrated) | ecs.t6-c1m4.2xlarge | $243.73 | 66.7% |

---


---

*Report generated by CM-Beetle*
