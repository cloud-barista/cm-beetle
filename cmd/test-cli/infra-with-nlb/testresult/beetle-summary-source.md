# Source Infrastructure Summary

**Generated At:** 2026-07-01 06:08:54

**Infrastructure Name:** infra-3-nodes

---

## Overview

| Metric | Value |
|--------|-------|
| Infrastructure Name | infra-3-nodes |
| Total Servers | 3 |
| Total CPU Cores | 5 |
| Total Memory (GB) | 26 |
| Total Disk (GB) | 68 |
| Total Networks | 3 |

## Compute Resources

### Servers (3)

#### 1. ip-10-0-1-30

| Component | Details |
|-----------|----------|
| CPU | Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz |
| **CPU CPUs** | 1 |
| CPU Cores | 1 |
| **CPU Threads** | 2 |
| **Equivalent vCPUs** | 2 (CPUs × Threads) |
| CPU Speed | 2.50 GHz |
| Architecture | x86_64 |
| **Memory** | 2 GB (DDR4) |
| **Root Disk** | 8 GB (SSD) |
| **OS** | Ubuntu 22.04.3 LTS (Jammy Jellyfish) |
| **Primary IP** | 10.0.1.30/24 |

#### 2. ip-10-0-1-221

| Component | Details |
|-----------|----------|
| CPU | Intel(R) Xeon(R) Platinum 8175M CPU @ 2.50GHz |
| **CPU CPUs** | 1 |
| CPU Cores | 2 |
| **CPU Threads** | 4 |
| **Equivalent vCPUs** | 4 (CPUs × Threads) |
| CPU Speed | 2.50 GHz |
| Architecture | x86_64 |
| **Memory** | 16 GB (DDR4) |
| **Root Disk** | 30 GB (SSD) |
| **OS** | Ubuntu 22.04.3 LTS (Jammy Jellyfish) |
| **Primary IP** | 10.0.1.221/24 |

#### 3. ip-10-0-1-138

| Component | Details |
|-----------|----------|
| CPU | Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz |
| **CPU CPUs** | 1 |
| CPU Cores | 2 |
| **CPU Threads** | 4 |
| **Equivalent vCPUs** | 4 (CPUs × Threads) |
| CPU Speed | 2.50 GHz |
| Architecture | x86_64 |
| **Memory** | 8 GB (DDR4) |
| **Root Disk** | 30 GB (SSD) |
| **OS** | Ubuntu 22.04.3 LTS (Jammy Jellyfish) |
| **Primary IP** | 10.0.1.138/24 |


## Network Resources

### Networks (1)

#### 1. network-1

| Property | Value |
|----------|-------|
| Network CIDR | 10.0.1.0/24 |
| Gateway | 10.0.1.1 |
| Connected Servers | 3 |
| Description | Network with gateway 10.0.1.1 |

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


## Security Resources

### Firewall Rules by Server (3 servers)

#### 1. ip-10-0-1-30

**Firewall Rules:** (4 rules)

| Direction | Protocol | Src CIDR | Src Ports | Dst CIDR | Dst Ports | Action |
|-----------|----------|----------|-----------|----------|-----------|--------|
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 22 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 9999 | allow |
| inbound | * | 10.0.0.0/16 | * | 0.0.0.0/0 | * | allow |
| outbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |

#### 2. ip-10-0-1-221

**Firewall Rules:** (4 rules)

| Direction | Protocol | Src CIDR | Src Ports | Dst CIDR | Dst Ports | Action |
|-----------|----------|----------|-----------|----------|-----------|--------|
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 22 | allow |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 8086 | allow |
| inbound | * | 10.0.0.0/16 | * | 0.0.0.0/0 | * | allow |
| outbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |

#### 3. ip-10-0-1-138

**Firewall Rules:** (4 rules)

| Direction | Protocol | Src CIDR | Src Ports | Dst CIDR | Dst Ports | Action |
|-----------|----------|----------|-----------|----------|-----------|--------|
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 22 | allow |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 8086 | allow |
| inbound | * | 10.0.0.0/16 | * | 0.0.0.0/0 | * | allow |
| outbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |


## Storage Resources

### Storage by Server (3 servers)

| Hostname | RootDisk (GB) | Type |
|----------|---------------|------|
| ip-10-0-1-30 | 8 | SSD |
| ip-10-0-1-221 | 30 | SSD |
| ip-10-0-1-138 | 30 | SSD |

### Storage by Type

| Type | Total (GB) | Servers |
|------|------------|----------|
| SSD | 68 | 3 |

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

This calculation ensures that the target VM has sufficient processing capacity equivalent to the source node.

