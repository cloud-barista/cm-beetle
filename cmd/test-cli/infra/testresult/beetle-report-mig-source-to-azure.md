# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

_Report generated: 2026-03-25 12:44:18_

---

## 📊 Migration Summary

**Target Cloud:** AZURE

**Target Region:** koreasouth

**Namespace:** mig01 | **MCI ID:** mig-1-mci101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $217.94 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| #   | Resource Type       | Count             | Status      | Details                                               |
| --- | ------------------- | ----------------- | ----------- | ----------------------------------------------------- |
| 1   | **Virtual Machine** | 3                 | ✅ Created  | 3 running, 3 total                                    |
| 2   | **VM Spec**         | 3                 | ✅ Selected | Standard_B2als_v2, Standard_B2as_v2, Standard_B4as_v2 |
| 3   | **VM OS Image**     | 2                 | ✅ Selected | Ubuntu 22.04                                          |
| 4   | **VNet (VPC)**      | 1                 | ✅ Created  | mig-1-vnet-01, CIDR: 10.0.0.0/21                      |
| 5   | **Subnet**          | 1                 | ✅ Created  | 10.0.1.0/24 (in mig-1-vnet-01)                        |
| 6   | **Security Group**  | 3 security groups | ✅ Created  | Total 52 rules in 3 sgs                               |
| 7   | **SSH Key**         | 1 keys            | ✅ Created  | For VM access control                                 |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source server(s) in the on-premise infrastructure.

| No. | Migrated VM                                                                                                                                                                                                                                                       | Source Server                                                |
| --- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| 1   | **VM Name:** mig-1-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71tfun693a119qk81m0<br>**Label(sourceMachineId):** 1-vm-ec268ed7-821e-9d73 | **Hostname:** N/A<br>**Machine ID:** 1-vm-ec268ed7-821e-9d73 |
| 2   | **VM Name:** mig-1-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71tfun693a119qk81o0<br>**Label(sourceMachineId):** 1-vm-ec288dd0-c6fa-8a49 | **Hostname:** N/A<br>**Machine ID:** 1-vm-ec288dd0-c6fa-8a49 |
| 3   | **VM Name:** mig-1-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71tfun693a119qk81n0<br>**Label(sourceMachineId):** 1-vm-ec2d32b5-98fb-5a96 | **Hostname:** N/A<br>**Machine ID:** 1-vm-ec2d32b5-98fb-5a96 |

---

## ⚙️ VM Specs

**Summary:** 3 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM                                     | VM Spec                                                                                      | Source Server                                                | Source Server Spec                                                         |
| --- | ----------------------------------------------- | -------------------------------------------------------------------------------------------- | ------------------------------------------------------------ | -------------------------------------------------------------------------- |
| 1   | mig-1-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** Standard_B2als_v2<br>**vCPUs:** 2<br>**Memory:** 3.9 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** 1-vm-ec268ed7-821e-9d73 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2   | mig-1-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Spec ID:** Standard_B2as_v2<br>**vCPUs:** 2<br>**Memory:** 7.8 GB<br>**Root Disk:** 50 GB  | **Hostname:** N/A<br>**Machine ID:** 1-vm-ec288dd0-c6fa-8a49 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3   | mig-1-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Spec ID:** Standard_B4as_v2<br>**vCPUs:** 4<br>**Memory:** 15.6 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** 1-vm-ec2d32b5-98fb-5a96 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 2 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM                                     | VM OS Image Info                                                                                                                                                                                             | Source Server                                                | Source OS                                                |
| --- | ----------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------ | -------------------------------------------------------- |
| 1   | mig-1-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** 0001-com-ubuntu-server-jammy-daily:22_04-daily-lts           | **Hostname:** N/A<br>**Machine ID:** 1-vm-ec268ed7-821e-9d73 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2   | mig-1-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** 0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2 | **Hostname:** N/A<br>**Machine ID:** 1-vm-ec288dd0-c6fa-8a49 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3   | mig-1-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** 0001-com-ubuntu-server-jammy-daily:22_04-daily-lts           | **Hostname:** N/A<br>**Machine ID:** 1-vm-ec2d32b5-98fb-5a96 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: mig-1-sg-01

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/d71tfg7693a119qk81h0 | **VNet:** mig-1-vnet-01 | **Rules:** 14

**Assigned VMs:**

- **VM:** mig-1-vm-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 1-vm-ec268ed7-821e-9d73

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule              | Note                 |
| --- | --------- | -------- | ------- | ----------- | --------------------------------- | -------------------- |
| 1   | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                   | Migrated from source |
| 2   | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                    | Migrated from source |
| 3   | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                  | Migrated from source |
| 4   | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                  | Migrated from source |
| 5   | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                    | Migrated from source |
| 6   | inbound   | TCP      | 80      | 0.0.0.0/0   | inbound tcp 80                    | Migrated from source |
| 7   | inbound   | TCP      | 443     | 0.0.0.0/0   | inbound tcp 443                   | Migrated from source |
| 8   | inbound   | TCP      | 8080    | 0.0.0.0/0   | inbound tcp 8080                  | Migrated from source |
| 9   | inbound   | TCP      | 9113    | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 10  | inbound   | UDP      | 9113    | 10.0.0.0/16 | inbound udp 9113 from 10.0.0.0/16 | Migrated from source |
| 11  | inbound   | ALL      |         | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 12  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 13  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 14  | outbound  | ALL      |         | 0.0.0.0/0   | outbound \* \*                    | Migrated from source |

### Security Group: mig-1-sg-02

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/d71tfkv693a119qk81jg | **VNet:** mig-1-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** mig-1-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 1-vm-ec2d32b5-98fb-5a96

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule               | Note                 |
| --- | --------- | -------- | ------- | ----------- | ---------------------------------- | -------------------- |
| 1   | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                    | Migrated from source |
| 2   | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                     | Migrated from source |
| 3   | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                   | Migrated from source |
| 4   | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                   | Migrated from source |
| 5   | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                     | Migrated from source |
| 6   | inbound   | TCP      | 2049    | 0.0.0.0/0   | inbound tcp 2049                   | Migrated from source |
| 7   | inbound   | UDP      | 2049    | 0.0.0.0/0   | inbound udp 2049                   | Migrated from source |
| 8   | inbound   | TCP      | 111     | 0.0.0.0/0   | inbound tcp 111                    | Migrated from source |
| 9   | inbound   | UDP      | 111     | 0.0.0.0/0   | inbound udp 111                    | Migrated from source |
| 10  | inbound   | TCP      | 20048   | 10.0.0.0/16 | inbound tcp 20048 from 10.0.0.0/16 | Migrated from source |
| 11  | inbound   | UDP      | 20048   | 10.0.0.0/16 | inbound udp 20048 from 10.0.0.0/16 | Migrated from source |
| 12  | inbound   | TCP      | 32803   | 10.0.0.0/16 | inbound tcp 32803 from 10.0.0.0/16 | Migrated from source |
| 13  | inbound   | UDP      | 32803   | 10.0.0.0/16 | inbound udp 32803 from 10.0.0.0/16 | Migrated from source |
| 14  | inbound   | TCP      | 9100    | 10.0.0.0/16 | inbound tcp 9100 from 10.0.0.0/16  | Migrated from source |
| 15  | inbound   | UDP      | 9100    | 10.0.0.0/16 | inbound udp 9100 from 10.0.0.0/16  | Migrated from source |
| 16  | inbound   | ALL      |         | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16  | Migrated from source |
| 17  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                  | Created by system    |
| 18  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                  | Created by system    |
| 19  | outbound  | ALL      |         | 0.0.0.0/0   | outbound \* \*                     | Migrated from source |

### Security Group: mig-1-sg-03

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/d71tfpn693a119qk81k0 | **VNet:** mig-1-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** mig-1-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 1-vm-ec288dd0-c6fa-8a49

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule              | Note                 |
| --- | --------- | -------- | ------- | ----------- | --------------------------------- | -------------------- |
| 1   | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                   | Migrated from source |
| 2   | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                    | Migrated from source |
| 3   | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                  | Migrated from source |
| 4   | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                  | Migrated from source |
| 5   | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                    | Migrated from source |
| 6   | inbound   | TCP      | 3306    | 10.0.0.0/16 | inbound tcp 3306 from 10.0.0.0/16 | Migrated from source |
| 7   | inbound   | UDP      | 3306    | 10.0.0.0/16 | inbound udp 3306 from 10.0.0.0/16 | Migrated from source |
| 8   | inbound   | TCP      | 4567    | 10.0.0.0/16 | inbound tcp 4567 from 10.0.0.0/16 | Migrated from source |
| 9   | inbound   | UDP      | 4567    | 10.0.0.0/16 | inbound udp 4567 from 10.0.0.0/16 | Migrated from source |
| 10  | inbound   | TCP      | 4568    | 10.0.0.0/16 | inbound tcp 4568 from 10.0.0.0/16 | Migrated from source |
| 11  | inbound   | UDP      | 4568    | 10.0.0.0/16 | inbound udp 4568 from 10.0.0.0/16 | Migrated from source |
| 12  | inbound   | TCP      | 4444    | 10.0.0.0/16 | inbound tcp 4444 from 10.0.0.0/16 | Migrated from source |
| 13  | inbound   | UDP      | 4444    | 10.0.0.0/16 | inbound udp 4444 from 10.0.0.0/16 | Migrated from source |
| 14  | inbound   | TCP      | 9104    | 10.0.0.0/16 | inbound tcp 9104 from 10.0.0.0/16 | Migrated from source |
| 15  | inbound   | UDP      | 9104    | 10.0.0.0/16 | inbound udp 9104 from 10.0.0.0/16 | Migrated from source |
| 16  | inbound   | ALL      |         | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 17  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 18  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 19  | outbound  | ALL      |         | 0.0.0.0/0   | outbound \* \*                    | Migrated from source |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source server network information.

### VPC(VNet)

| No. | VPC(VNet)                                                                                                                                                          | CIDR Block  |
| --- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------- |
| 1   | **Name:** mig-1-vnet-01<br>**ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71tf9v693a119qk818g | 10.0.0.0/21 |

### Subnets

| No. | Subnet                                                                                                                                                                                            | CIDR Block  | Associated VPC(VNet) |
| --- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------- | -------------------- |
| 1   | **Name:** mig-1-subnet-01<br>**ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71tf9v693a119qk818g/subnets/d71tf9v693a119qk8190 | 10.0.1.0/24 | mig-1-vnet-01        |

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

| No. | SSH Key Name    | CSP Key ID                                                                                                                    | Fingerprint | Usage             |
| --- | --------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------- | ----------------- |
| 1   | mig-1-sshkey-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/d71tfcf693a119qk81b0 |             | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period  | Cost (USD) |
| ------- | ---------- |
| Hourly  | $0.3027    |
| Daily   | $7.26      |
| Monthly | $217.94    |
| Yearly  | $2615.33   |

### Cost Breakdown by Component

| Component                | Spec              | Monthly Cost | Percentage |
| ------------------------ | ----------------- | ------------ | ---------- |
| ip-10-0-1-30 (migrated)  | Standard_B2als_v2 | $31.10       | 14.3%      |
| ip-10-0-1-221 (migrated) | Standard_B4as_v2  | $124.56      | 57.2%      |
| ip-10-0-1-138 (migrated) | Standard_B2as_v2  | $62.28       | 28.6%      |

---

---

_Report generated by CM-Beetle_
