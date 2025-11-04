# Source Infrastructure Summary

**Generated At:** 2025-11-04 15:11:23

**Infrastructure Name:** infra-2-servers

---

## Overview

| Metric | Value |
|--------|-------|
| Infrastructure Name | infra-2-servers |
| Total Servers | 2 |
| Total CPU Cores | 6 |
| Total Memory (GB) | 32 |
| Total Disk (GB) | 2405 |
| Total Networks | 2 |

## Compute Resources

### Servers (2)

#### 1. cm-nfs

| Component | Details |
|-----------|----------|
| CPU | Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz |
| **CPU CPUs** | 1 |
| CPU Cores | 2 |
| **CPU Threads** | 4 |
| **Equivalent vCPUs** | 4 (CPUs × Threads) |
| CPU Speed | 2.50 GHz |
| Architecture | x86_64 |
| **Memory** | 16 GB (DDR4) |
| **Root Disk** | 1093 GB (HDD) |
| **OS** | Ubuntu 22.04.5 LTS (Jammy Jellyfish) |
| **Primary IP** | 172.29.0.102/24 |

#### 2. cm-web

| Component | Details |
|-----------|----------|
| CPU | Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz |
| **CPU CPUs** | 1 |
| CPU Cores | 4 |
| **CPU Threads** | 8 |
| **Equivalent vCPUs** | 8 (CPUs × Threads) |
| CPU Speed | 3.10 GHz |
| Architecture | x86_64 |
| **Memory** | 16 GB (DDR4) |
| **Root Disk** | 1312 GB (HDD) |
| **OS** | Ubuntu 22.04.5 LTS (Jammy Jellyfish) |
| **Primary IP** | 172.29.0.103/24 |


## Network Resources

### Networks (1)

#### 1. network-1

| Property | Value |
|----------|-------|
| Network CIDR | 192.168.110.0/24 |
| Gateway | 192.168.110.254 |
| Connected Servers | 2 |
| Description | Network with gateway 192.168.110.254 |

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
|-------------|---------|-----------|
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
|-------------|---------|-----------|
| 0.0.0.0/0 | 192.168.110.254 | br-ex |
| 172.29.0.0/24 | on-link | eno1np0 |
| 192.168.110.0/24 | 192.168.110.254 | br-ex |


## Security Resources

### Firewall Rules by Server (2 servers)

#### 1. cm-nfs

**Firewall Rules:** (3 rules)

| Direction | Protocol | Src CIDR | Src Ports | Dst CIDR | Dst Ports | Action |
|-----------|----------|----------|-----------|----------|-----------|--------|
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 10022 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 8081,8082 | allow |
| inbound | udp | 192.168.110.0/24 | * | 0.0.0.0/0 | 53 | allow |

#### 2. cm-web

*No firewall rules*


## Storage Resources

### Storage by Server (2 servers)

| Hostname | RootDisk (GB) | Type |
|----------|---------------|------|
| cm-nfs | 1093 | HDD |
| cm-web | 1312 | HDD |

### Storage by Type

| Type | Total (GB) | Servers |
|------|------------|----------|
| HDD | 2405 | 2 |

---

## 📝 Important Notes for Cloud Migration

### CPU to vCPU Mapping

When migrating to cloud VMs, the **Equivalent vCPUs** value is calculated as:

```
Equivalent vCPUs = CPUs × CPU Threads
```

**Example:**
- Source Server: 2 CPUs, 2 Threads per CPU
- Calculation: 2 CPUs × 2 Threads = **4 vCPUs**
- Target VM Spec: Select a VM with **4 vCPUs** (e.g., AWS t3.xlarge)

This calculation ensures that the target VM has sufficient processing capacity equivalent to the source server.

