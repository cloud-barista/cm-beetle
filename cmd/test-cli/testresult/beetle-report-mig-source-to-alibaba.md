# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

_Report generated: 2026-03-25 13:31:38_

---

## 📊 Migration Summary

**Target Cloud:** ALIBABA

**Target Region:**

**Namespace:** mig01 | **MCI ID:** mig-4-mci101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| #   | Resource Type       | Count             | Status          | Details                          |
| --- | ------------------- | ----------------- | --------------- | -------------------------------- |
| 1   | **Virtual Machine** | 3                 | ✅ Created      | 0 running, 3 total               |
| 2   | **VM Spec**         | 0                 | ⚠️ Not selected | No specs used                    |
| 3   | **VM OS Image**     | 1                 | ✅ Selected     | Ubuntu 22.04                     |
| 4   | **VNet (VPC)**      | 1                 | ✅ Created      | mig-4-vnet-01, CIDR: 10.0.0.0/21 |
| 5   | **Subnet**          | 1                 | ✅ Created      | 10.0.1.0/24 (in mig-4-vnet-01)   |
| 6   | **Security Group**  | 3 security groups | ✅ Created      | Total 52 rules in 3 sgs          |
| 7   | **SSH Key**         | 1 keys            | ✅ Created      | For VM access control            |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source server(s) in the on-premise infrastructure.

| No. | Migrated VM                                                                                                                        | Source Server                                                |
| --- | ---------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| 1   | **VM Name:** mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** <br>**Label(sourceMachineId):** 4-vm-ec268ed7-821e-9d73 | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec268ed7-821e-9d73 |
| 2   | **VM Name:** mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** <br>**Label(sourceMachineId):** 4-vm-ec288dd0-c6fa-8a49 | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec288dd0-c6fa-8a49 |
| 3   | **VM Name:** mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** <br>**Label(sourceMachineId):** 4-vm-ec2d32b5-98fb-5a96 | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec2d32b5-98fb-5a96 |

---

## ⚙️ VM Specs

**Summary:** 0 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM                                     | VM Spec                                                                      | Source Server                                                | Source Server Spec                                                         |
| --- | ----------------------------------------------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------ | -------------------------------------------------------------------------- |
| 1   | mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** <br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB  | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec268ed7-821e-9d73 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2   | mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Spec ID:** <br>**vCPUs:** 2<br>**Memory:** 8.0 GB<br>**Root Disk:** 50 GB  | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec288dd0-c6fa-8a49 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3   | mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Spec ID:** <br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec2d32b5-98fb-5a96 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM                                     | VM OS Image Info                                                                                                                 | Source Server                                                | Source OS                                                |
| --- | ----------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ | -------------------------------------------------------- |
| 1   | mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260119.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu 22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec268ed7-821e-9d73 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2   | mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260119.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu 22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec288dd0-c6fa-8a49 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3   | mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260119.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu 22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec2d32b5-98fb-5a96 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: mig-4-sg-01

**CSP ID:** sg-mj7bmek530yys7apfwdi | **VNet:** mig-4-vnet-01 | **Rules:** 14

**Assigned VMs:**

- **VM:** mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 4-vm-ec268ed7-821e-9d73

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule              | Note                 |
| --- | --------- | -------- | ------- | ----------- | --------------------------------- | -------------------- |
| 1   | inbound   | ALL      |         | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 2   | inbound   | UDP      | 9113    | 10.0.0.0/16 | inbound udp 9113 from 10.0.0.0/16 | Migrated from source |
| 3   | inbound   | TCP      | 9113    | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 4   | inbound   | TCP      | 8080    | 0.0.0.0/0   | inbound tcp 8080                  | Migrated from source |
| 5   | inbound   | TCP      | 443     | 0.0.0.0/0   | inbound tcp 443                   | Migrated from source |
| 6   | inbound   | TCP      | 80      | 0.0.0.0/0   | inbound tcp 80                    | Migrated from source |
| 7   | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                    | Migrated from source |
| 8   | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                  | Migrated from source |
| 9   | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                  | Migrated from source |
| 10  | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                    | Migrated from source |
| 11  | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                   | Migrated from source |
| 12  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 13  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 14  | outbound  | ALL      |         | 0.0.0.0/0   | outbound \* \*                    | Migrated from source |

### Security Group: mig-4-sg-02

**CSP ID:** sg-mj7eu0g9c9oo47ounzk6 | **VNet:** mig-4-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 4-vm-ec2d32b5-98fb-5a96

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule               | Note                 |
| --- | --------- | -------- | ------- | ----------- | ---------------------------------- | -------------------- |
| 1   | inbound   | ALL      |         | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16  | Migrated from source |
| 2   | inbound   | UDP      | 9100    | 10.0.0.0/16 | inbound udp 9100 from 10.0.0.0/16  | Migrated from source |
| 3   | inbound   | TCP      | 9100    | 10.0.0.0/16 | inbound tcp 9100 from 10.0.0.0/16  | Migrated from source |
| 4   | inbound   | UDP      | 32803   | 10.0.0.0/16 | inbound udp 32803 from 10.0.0.0/16 | Migrated from source |
| 5   | inbound   | TCP      | 32803   | 10.0.0.0/16 | inbound tcp 32803 from 10.0.0.0/16 | Migrated from source |
| 6   | inbound   | UDP      | 20048   | 10.0.0.0/16 | inbound udp 20048 from 10.0.0.0/16 | Migrated from source |
| 7   | inbound   | TCP      | 20048   | 10.0.0.0/16 | inbound tcp 20048 from 10.0.0.0/16 | Migrated from source |
| 8   | inbound   | UDP      | 111     | 0.0.0.0/0   | inbound udp 111                    | Migrated from source |
| 9   | inbound   | TCP      | 111     | 0.0.0.0/0   | inbound tcp 111                    | Migrated from source |
| 10  | inbound   | UDP      | 2049    | 0.0.0.0/0   | inbound udp 2049                   | Migrated from source |
| 11  | inbound   | TCP      | 2049    | 0.0.0.0/0   | inbound tcp 2049                   | Migrated from source |
| 12  | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                     | Migrated from source |
| 13  | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                   | Migrated from source |
| 14  | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                   | Migrated from source |
| 15  | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                     | Migrated from source |
| 16  | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                    | Migrated from source |
| 17  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                  | Created by system    |
| 18  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                  | Created by system    |
| 19  | outbound  | ALL      |         | 0.0.0.0/0   | outbound \* \*                     | Migrated from source |

### Security Group: mig-4-sg-03

**CSP ID:** sg-mj7i5crahq8i5rzcer4u | **VNet:** mig-4-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 4-vm-ec288dd0-c6fa-8a49

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule              | Note                 |
| --- | --------- | -------- | ------- | ----------- | --------------------------------- | -------------------- |
| 1   | inbound   | ALL      |         | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 2   | inbound   | UDP      | 9104    | 10.0.0.0/16 | inbound udp 9104 from 10.0.0.0/16 | Migrated from source |
| 3   | inbound   | TCP      | 9104    | 10.0.0.0/16 | inbound tcp 9104 from 10.0.0.0/16 | Migrated from source |
| 4   | inbound   | UDP      | 4444    | 10.0.0.0/16 | inbound udp 4444 from 10.0.0.0/16 | Migrated from source |
| 5   | inbound   | TCP      | 4444    | 10.0.0.0/16 | inbound tcp 4444 from 10.0.0.0/16 | Migrated from source |
| 6   | inbound   | UDP      | 4568    | 10.0.0.0/16 | inbound udp 4568 from 10.0.0.0/16 | Migrated from source |
| 7   | inbound   | TCP      | 4568    | 10.0.0.0/16 | inbound tcp 4568 from 10.0.0.0/16 | Migrated from source |
| 8   | inbound   | UDP      | 4567    | 10.0.0.0/16 | inbound udp 4567 from 10.0.0.0/16 | Migrated from source |
| 9   | inbound   | TCP      | 4567    | 10.0.0.0/16 | inbound tcp 4567 from 10.0.0.0/16 | Migrated from source |
| 10  | inbound   | UDP      | 3306    | 10.0.0.0/16 | inbound udp 3306 from 10.0.0.0/16 | Migrated from source |
| 11  | inbound   | TCP      | 3306    | 10.0.0.0/16 | inbound tcp 3306 from 10.0.0.0/16 | Migrated from source |
| 12  | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                    | Migrated from source |
| 13  | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                  | Migrated from source |
| 14  | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                  | Migrated from source |
| 15  | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                    | Migrated from source |
| 16  | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                   | Migrated from source |
| 17  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 18  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 19  | outbound  | ALL      |         | 0.0.0.0/0   | outbound \* \*                    | Migrated from source |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source server network information.

### VPC(VNet)

| No. | VPC(VNet)                                                    | CIDR Block  |
| --- | ------------------------------------------------------------ | ----------- |
| 1   | **Name:** mig-4-vnet-01<br>**ID:** vpc-mj72oyy2s6tly01dcchqs | 10.0.0.0/21 |

### Subnets

| No. | Subnet                                                         | CIDR Block  | Associated VPC(VNet) |
| --- | -------------------------------------------------------------- | ----------- | -------------------- |
| 1   | **Name:** mig-4-subnet-01<br>**ID:** vsw-mj782snkv4i7r0see4na0 | 10.0.1.0/24 | mig-4-vnet-01        |

### Source Network Information

**CIDR:** 10.0.1.0/24 | **Gateway:** 10.0.1.1 | **Connected Servers:** 3

### Network Details by Server (3 servers)

#### 1. ip-10-0-1-30

**Active Interfaces:**

| Interface | IP Address  | State |
| --------- | ----------- | ----- |
| lo        | 127.0.0.1/8 | up    |

#### 2. ip-10-0-1-221

**Active Interfaces:**

| Interface | IP Address  | State |
| --------- | ----------- | ----- |
| lo        | 127.0.0.1/8 | up    |

#### 3. ip-10-0-1-138

**Active Interfaces:**

| Interface | IP Address  | State |
| --------- | ----------- | ----- |
| lo        | 127.0.0.1/8 | up    |

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name    | CSP Key ID           | Fingerprint                      | Usage             |
| --- | --------------- | -------------------- | -------------------------------- | ----------------- |
| 1   | mig-4-sshkey-01 | d71u75v693a119sjo6g0 | e88593506421268e6551072903407d33 | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period  | Cost (USD) |
| ------- | ---------- |
| Hourly  | $0.0000    |
| Daily   | $0.00      |
| Monthly | $0.00      |
| Yearly  | $0.00      |

### Cost Breakdown by Component

| Component                | Spec | Monthly Cost | Percentage |
| ------------------------ | ---- | ------------ | ---------- |
| ip-10-0-1-30 (migrated)  | N/A  | $0.00        | 0.0%       |
| ip-10-0-1-221 (migrated) | N/A  | $0.00        | 0.0%       |
| ip-10-0-1-138 (migrated) | N/A  | $0.00        | 0.0%       |

---

---

_Report generated by CM-Beetle_
