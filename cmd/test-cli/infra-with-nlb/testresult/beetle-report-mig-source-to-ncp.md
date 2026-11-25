# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-11-25 09:24:03*

---

## 📊 Migration Summary

**Target Cloud:** NCP

**Target Region:** KR

**Namespace:** mig01 | **Infra ID:** my08-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**NLB Migration Status:** Completed

**Total Load Balancers:** 1

**Migrated Load Balancers:** 1

**💰 Estimated Monthly Cost:** $300.02 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | s4-g3, c2-g3 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my08-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my08-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 18 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |
| 8 | **Network Load Balancer** | 1 | ✅ Created | my08-ng-influxdb-back (TCP:9999 -> TCP:8086) |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my08-ng-influxdb-back-1<br>**VM ID:** 147164275<br>**Label(sourceMachineId):** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 |
| 2 | **VM Name:** my08-ng-influxdb-back-2<br>**VM ID:** 147164269<br>**Label(sourceMachineId):** ec288dd0-c6fa-8a49-2f60-bc898311febf | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf |
| 3 | **VM Name:** my08-ng-ip-10-0-1-30-62161624-1<br>**VM ID:** 147164272<br>**Label(sourceMachineId):** ec268ed7-821e-9d73-e79f-961262161624 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my08-ng-influxdb-back-1 | **Spec ID:** s4-g3<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 30 GB |
| 2 | my08-ng-influxdb-back-2 | **Spec ID:** s4-g3<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 8 GB<br>**Root Disk:** 30 GB |
| 3 | my08-ng-ip-10-0-1-30-62161624-1 | **Spec ID:** c2-g3<br>**vCPUs:** 2<br>**Memory:** 4.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 8 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my08-ng-influxdb-back-1 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | my08-ng-influxdb-back-2 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | my08-ng-ip-10-0-1-30-62161624-1 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 18 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my08-mig-sg-01

**CSP ID:** 423537 | **VNet:** my08-mig-vnet-01 | **Rules:** 10

**Assigned VMs:**

- **VM:** my08-ng-influxdb-back-1
  - **Source Server:** **Hostname:** ip-10-0-1-221, **Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932
- **VM:** my08-ng-influxdb-back-2
  - **Source Server:** **Hostname:** ip-10-0-1-138, **Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 2 | inbound | TCP | 9999 | 0.0.0.0/0 | - | Created by system |
| 3 | inbound | ICMP |  | 0.0.0.0/0 | - | Created by system |
| 4 | inbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 5 | inbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 6 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 7 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 8 | outbound | ICMP |  | 0.0.0.0/0 | - | Created by system |
| 9 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 10 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |

### Security Group: my08-mig-sg-02

**CSP ID:** 423538 | **VNet:** my08-mig-vnet-01 | **Rules:** 8

**Assigned VMs:**

- **VM:** my08-ng-ip-10-0-1-30-62161624-1
  - **Source Server:** **Hostname:** ip-10-0-1-30, **Machine ID:** ec268ed7-821e-9d73-e79f-961262161624

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | - | Created by system |
| 2 | inbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 3 | inbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 4 | inbound | TCP | 9999 | 0.0.0.0/0 | inbound tcp 9999 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | outbound | ICMP |  | 0.0.0.0/0 | - | Created by system |
| 7 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 8 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source node network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** my08-mig-vnet-01<br>**ID:** 153788 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my08-mig-subnet-01<br>**ID:** 336709 | 10.0.1.0/24 | my08-mig-vnet-01 |

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
| 10.0.1.0/24 | 10.0.1.1 | ens5 |

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
| 10.0.1.0/24 | 10.0.1.1 | ens5 |

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
| 10.0.1.0/24 | 10.0.1.1 | ens5 |

---

## ⚖️ Network Load Balancers (NLBs)

**Summary:** 1 of 1 load balancer(s) migrated from on-premise software load balancers to cloud-managed NLBs.

### 1. influxdb_back (haproxy) → my08-ng-influxdb-back

**Configuration Comparison:**

| Property | Source (On-Premise) | Target (Cloud NLB) | Note |
|----------|---------------------|--------------------|------|
| **Load Balancer / Name** | haproxy (Host: ip-10-0-1-30) | my08-ng-influxdb-back (CSP ID: tb4eeq1ff30hug6d5pqv) | Replaced by managed NLB |
| **Type / Scope** | Software LB | PUBLIC / REGION | Cloud-native NLB |
| **Frontend Listener** | *:9999 (TCP) | TCP:9999 | IP: 211.233.218.78 |
| **Backend / Target Group** | influxdb_back (TCP:8086, roundrobin) | my08-ng-influxdb-back (TCP:8086, 2 nodes) | Target Group attached to NodeGroup |
| **Health Check** | Enabled (Interval: 10s, Timeout: 10s, Threshold: 3) | TCP:8086 (Interval: 10s, Timeout: 0s, Threshold: 3) | Preserved |

**Backend Node Mappings:**

| No. | Source Backend Server | Source Endpoint | Target Node | Target CSP VM ID | Target Private IP | Target Public IP | Target Port | Status |
|-----|-----------------------|-----------------|-------------|------------------|-------------------|------------------|-------------|--------|
| 1 | influx1 (ip-10-0-1-221) | 10.0.1.221:8086 | my08-ng-influxdb-back-1 | 147164275 | 10.0.1.8 | 223.130.150.80 | 8086 | Running |
| 2 | influx2 (ip-10-0-1-138) | 10.0.1.138:8086 | my08-ng-influxdb-back-2 | 147164269 | 10.0.1.6 | 175.45.195.206 | 8086 | Running |

**Migration Notes:**
- On-premise haproxy on ip-10-0-1-30 (migrated to my08-ng-ip-10-0-1-30-62161624-1) is replaced by cloud-managed NLB.

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my08-mig-sshkey-01 | tbmo1vkhfnoa7ta0p33p | 2a:8f:c2:e3:53:12:d9:33:23:0b:fc:24:37:da:ff:85 | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.4167 |
| Daily | $10.00 |
| Monthly | $300.02 |
| Yearly | $3600.29 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | c2-g3 | $48.74 | 16.2% |
| ip-10-0-1-221 (migrated) | s4-g3 | $125.64 | 41.9% |
| ip-10-0-1-138 (migrated) | s4-g3 | $125.64 | 41.9% |

---


---

*Report generated by CM-Beetle*
