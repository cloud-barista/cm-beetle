# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

_Report generated: 2026-03-25 12:48:49_

---

## 📊 Migration Summary

**Target Cloud:** NCP

**Target Region:** KR

**Namespace:** mig01 | **MCI ID:** mig-3-mci101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $239.40 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| #   | Resource Type       | Count             | Status      | Details                          |
| --- | ------------------- | ----------------- | ----------- | -------------------------------- |
| 1   | **Virtual Machine** | 3                 | ✅ Created  | 3 running, 3 total               |
| 2   | **VM Spec**         | 3                 | ✅ Selected | ci2-g3, s2-g3, s4-g3a            |
| 3   | **VM OS Image**     | 1                 | ✅ Selected | Ubuntu 22.04                     |
| 4   | **VNet (VPC)**      | 1                 | ✅ Created  | mig-3-vnet-01, CIDR: 10.0.0.0/21 |
| 5   | **Subnet**          | 1                 | ✅ Created  | 10.0.1.0/24 (in mig-3-vnet-01)   |
| 6   | **Security Group**  | 3 security groups | ✅ Created  | Total 55 rules in 3 sgs          |
| 7   | **SSH Key**         | 1 keys            | ✅ Created  | For VM access control            |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source server(s) in the on-premise infrastructure.

| No. | Migrated VM                                                                                                                                 | Source Server                                                |
| --- | ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| 1   | **VM Name:** mig-3-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** 134455079<br>**Label(sourceMachineId):** 3-vm-ec268ed7-821e-9d73 | **Hostname:** N/A<br>**Machine ID:** 3-vm-ec268ed7-821e-9d73 |
| 2   | **VM Name:** mig-3-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** 134455087<br>**Label(sourceMachineId):** 3-vm-ec288dd0-c6fa-8a49 | **Hostname:** N/A<br>**Machine ID:** 3-vm-ec288dd0-c6fa-8a49 |
| 3   | **VM Name:** mig-3-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** 134455073<br>**Label(sourceMachineId):** 3-vm-ec2d32b5-98fb-5a96 | **Hostname:** N/A<br>**Machine ID:** 3-vm-ec2d32b5-98fb-5a96 |

---

## ⚙️ VM Specs

**Summary:** 3 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM                                     | VM Spec                                                                            | Source Server                                                | Source Server Spec                                                         |
| --- | ----------------------------------------------- | ---------------------------------------------------------------------------------- | ------------------------------------------------------------ | -------------------------------------------------------------------------- |
| 1   | mig-3-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** ci2-g3<br>**vCPUs:** 2<br>**Memory:** 4.0 GB<br>**Root Disk:** 50 GB  | **Hostname:** N/A<br>**Machine ID:** 3-vm-ec268ed7-821e-9d73 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2   | mig-3-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Spec ID:** s2-g3<br>**vCPUs:** 2<br>**Memory:** 8.0 GB<br>**Root Disk:** 50 GB   | **Hostname:** N/A<br>**Machine ID:** 3-vm-ec288dd0-c6fa-8a49 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3   | mig-3-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Spec ID:** s4-g3a<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** 3-vm-ec2d32b5-98fb-5a96 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM                                     | VM OS Image Info                                                                                               | Source Server                                                | Source OS                                                |
| --- | ----------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ | -------------------------------------------------------- |
| 1   | mig-3-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** N/A<br>**Machine ID:** 3-vm-ec268ed7-821e-9d73 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2   | mig-3-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** N/A<br>**Machine ID:** 3-vm-ec288dd0-c6fa-8a49 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3   | mig-3-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** N/A<br>**Machine ID:** 3-vm-ec2d32b5-98fb-5a96 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 55 security rule(s) have been created and configured for the migrated VMs.

### Security Group: mig-3-sg-01

**CSP ID:** 343878 | **VNet:** mig-3-vnet-01 | **Rules:** 15

**Assigned VMs:**

- **VM:** mig-3-vm-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 3-vm-ec268ed7-821e-9d73

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule              | Note                 |
| --- | --------- | -------- | ------- | ----------- | --------------------------------- | -------------------- |
| 1   | inbound   | UDP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 2   | inbound   | TCP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 3   | inbound   | UDP      | 9113    | 10.0.0.0/16 | inbound udp 9113 from 10.0.0.0/16 | Migrated from source |
| 4   | inbound   | TCP      | 9113    | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 5   | inbound   | TCP      | 8080    | 0.0.0.0/0   | inbound tcp 8080                  | Migrated from source |
| 6   | inbound   | TCP      | 443     | 0.0.0.0/0   | inbound tcp 443                   | Migrated from source |
| 7   | inbound   | TCP      | 80      | 0.0.0.0/0   | inbound tcp 80                    | Migrated from source |
| 8   | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                    | Migrated from source |
| 9   | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                  | Migrated from source |
| 10  | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                  | Migrated from source |
| 11  | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                    | Migrated from source |
| 12  | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                   | Migrated from source |
| 13  | outbound  | ICMP     |         | 0.0.0.0/0   | -                                 | Created by system    |
| 14  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 15  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |

### Security Group: mig-3-sg-02

**CSP ID:** 343879 | **VNet:** mig-3-vnet-01 | **Rules:** 20

**Assigned VMs:**

- **VM:** mig-3-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 3-vm-ec2d32b5-98fb-5a96

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule               | Note                 |
| --- | --------- | -------- | ------- | ----------- | ---------------------------------- | -------------------- |
| 1   | inbound   | UDP      | 1-65535 | 0.0.0.0/0   | -                                  | Created by system    |
| 2   | inbound   | TCP      | 1-65535 | 0.0.0.0/0   | -                                  | Created by system    |
| 3   | inbound   | UDP      | 9100    | 10.0.0.0/16 | inbound udp 9100 from 10.0.0.0/16  | Migrated from source |
| 4   | inbound   | TCP      | 9100    | 10.0.0.0/16 | inbound tcp 9100 from 10.0.0.0/16  | Migrated from source |
| 5   | inbound   | UDP      | 32803   | 10.0.0.0/16 | inbound udp 32803 from 10.0.0.0/16 | Migrated from source |
| 6   | inbound   | TCP      | 32803   | 10.0.0.0/16 | inbound tcp 32803 from 10.0.0.0/16 | Migrated from source |
| 7   | inbound   | UDP      | 20048   | 10.0.0.0/16 | inbound udp 20048 from 10.0.0.0/16 | Migrated from source |
| 8   | inbound   | TCP      | 20048   | 10.0.0.0/16 | inbound tcp 20048 from 10.0.0.0/16 | Migrated from source |
| 9   | inbound   | UDP      | 111     | 0.0.0.0/0   | inbound udp 111                    | Migrated from source |
| 10  | inbound   | TCP      | 111     | 0.0.0.0/0   | inbound tcp 111                    | Migrated from source |
| 11  | inbound   | UDP      | 2049    | 0.0.0.0/0   | inbound udp 2049                   | Migrated from source |
| 12  | inbound   | TCP      | 2049    | 0.0.0.0/0   | inbound tcp 2049                   | Migrated from source |
| 13  | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                     | Migrated from source |
| 14  | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                   | Migrated from source |
| 15  | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                   | Migrated from source |
| 16  | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                     | Migrated from source |
| 17  | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                    | Migrated from source |
| 18  | outbound  | ICMP     |         | 0.0.0.0/0   | -                                  | Created by system    |
| 19  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                  | Created by system    |
| 20  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                  | Created by system    |

### Security Group: mig-3-sg-03

**CSP ID:** 343880 | **VNet:** mig-3-vnet-01 | **Rules:** 20

**Assigned VMs:**

- **VM:** mig-3-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 3-vm-ec288dd0-c6fa-8a49

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule              | Note                 |
| --- | --------- | -------- | ------- | ----------- | --------------------------------- | -------------------- |
| 1   | inbound   | UDP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 2   | inbound   | TCP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 3   | inbound   | UDP      | 9104    | 10.0.0.0/16 | inbound udp 9104 from 10.0.0.0/16 | Migrated from source |
| 4   | inbound   | TCP      | 9104    | 10.0.0.0/16 | inbound tcp 9104 from 10.0.0.0/16 | Migrated from source |
| 5   | inbound   | UDP      | 4444    | 10.0.0.0/16 | inbound udp 4444 from 10.0.0.0/16 | Migrated from source |
| 6   | inbound   | TCP      | 4444    | 10.0.0.0/16 | inbound tcp 4444 from 10.0.0.0/16 | Migrated from source |
| 7   | inbound   | UDP      | 4568    | 10.0.0.0/16 | inbound udp 4568 from 10.0.0.0/16 | Migrated from source |
| 8   | inbound   | TCP      | 4568    | 10.0.0.0/16 | inbound tcp 4568 from 10.0.0.0/16 | Migrated from source |
| 9   | inbound   | UDP      | 4567    | 10.0.0.0/16 | inbound udp 4567 from 10.0.0.0/16 | Migrated from source |
| 10  | inbound   | TCP      | 4567    | 10.0.0.0/16 | inbound tcp 4567 from 10.0.0.0/16 | Migrated from source |
| 11  | inbound   | UDP      | 3306    | 10.0.0.0/16 | inbound udp 3306 from 10.0.0.0/16 | Migrated from source |
| 12  | inbound   | TCP      | 3306    | 10.0.0.0/16 | inbound tcp 3306 from 10.0.0.0/16 | Migrated from source |
| 13  | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                    | Migrated from source |
| 14  | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                  | Migrated from source |
| 15  | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                  | Migrated from source |
| 16  | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                    | Migrated from source |
| 17  | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                   | Migrated from source |
| 18  | outbound  | ICMP     |         | 0.0.0.0/0   | -                                 | Created by system    |
| 19  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 20  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source server network information.

### VPC(VNet)

| No. | VPC(VNet)                                 | CIDR Block  |
| --- | ----------------------------------------- | ----------- |
| 1   | **Name:** mig-3-vnet-01<br>**ID:** 135959 | 10.0.0.0/21 |

### Subnets

| No. | Subnet                                      | CIDR Block  | Associated VPC(VNet) |
| --- | ------------------------------------------- | ----------- | -------------------- |
| 1   | **Name:** mig-3-subnet-01<br>**ID:** 293642 | 10.0.1.0/24 | mig-3-vnet-01        |

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

| No. | SSH Key Name    | CSP Key ID           | Fingerprint | Usage             |
| --- | --------------- | -------------------- | ----------- | ----------------- |
| 1   | mig-3-sshkey-01 | d71tfgv693a119qk81hg |             | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period  | Cost (USD) |
| ------- | ---------- |
| Hourly  | $0.3325    |
| Daily   | $7.98      |
| Monthly | $239.40    |
| Yearly  | $2872.80   |

### Cost Breakdown by Component

| Component                | Spec   | Monthly Cost | Percentage |
| ------------------------ | ------ | ------------ | ---------- |
| ip-10-0-1-30 (migrated)  | ci2-g3 | $52.56       | 22.0%      |
| ip-10-0-1-221 (migrated) | s4-g3a | $125.78      | 52.5%      |
| ip-10-0-1-138 (migrated) | s2-g3  | $61.06       | 25.5%      |

---

---

_Report generated by CM-Beetle_
