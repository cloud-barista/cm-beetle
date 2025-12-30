# Source Infrastructure Summary

**Generated At:** 2026-01-05 00:11:52

**Infrastructure Name:** infra-3-servers

---

## Overview

| Metric | Value |
|--------|-------|
| Infrastructure Name | infra-3-servers |
| Total Servers | 3 |
| Total CPU Cores | 4 |
| Total Memory (GB) | 26 |
| Total Disk (GB) | 0 |
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
| **OS** | Ubuntu 22.04.3 LTS (Jammy Jellyfish) |
| **Primary IP** | 10.0.1.221/24 |

#### 3. ip-10-0-1-138

| Component | Details |
|-----------|----------|
| CPU | Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz |
| **CPU CPUs** | 1 |
| CPU Cores | 1 |
| **CPU Threads** | 2 |
| **Equivalent vCPUs** | 2 (CPUs × Threads) |
| CPU Speed | 2.50 GHz |
| Architecture | x86_64 |
| **Memory** | 8 GB (DDR4) |
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

**Firewall Rules:** (45 rules)

| Direction | Protocol | Src CIDR | Src Ports | Dst CIDR | Dst Ports | Action |
|-----------|----------|----------|-----------|----------|-----------|--------|
| inbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| inbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | deny |
| inbound | icmp | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| inbound | udp | 0.0.0.0/0 | 67 | 0.0.0.0/0 | 68 | allow |
| inbound | udp | 0.0.0.0/0 | * | 224.0.0.251/32 | 5353 | allow |
| inbound | udp | 0.0.0.0/0 | * | 239.255.255.250/32 | 1900 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 22 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 80 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 443 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 8080 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 3306 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 5432 | deny |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 9113 | allow |
| inbound | udp | 10.0.0.0/16 | * | 0.0.0.0/0 | 9113 | allow |
| inbound | * | 10.0.0.0/16 | * | 0.0.0.0/0 | * | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 23 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 135 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 139 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 445 | deny |
| outbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| outbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| outbound | udp | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| inbound | * | ::/0 | * | ::/0 | * | allow |
| inbound | * | ::/0 | * | ::/0 | * | deny |
| inbound | icmpv6 | ::/0 | * | ::/0 | * | allow |
| inbound | icmpv6 | fe80::/10 | * | ::/0 | * | allow |
| inbound | udp | fe80::/10 | 547 | fe80::/10 | 546 | allow |
| inbound | udp | ::/0 | * | ff02::fb/128 | 5353 | allow |
| inbound | udp | ::/0 | * | ff02::f/128 | 1900 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 22 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 80 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 443 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 8080 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 3306 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 5432 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 23 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 135 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 139 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 445 | deny |
| outbound | * | ::/0 | * | ::/0 | * | allow |
| outbound | * | ::/0 | * | ::/0 | * | deny |
| outbound | icmpv6 | ::/0 | * | ::/0 | * | allow |
| outbound | icmpv6 | fe80::/10 | * | ::/0 | * | allow |
| outbound | tcp | ::/0 | * | ::/0 | * | allow |
| outbound | udp | ::/0 | * | ::/0 | * | allow |

#### 2. ip-10-0-1-221

**Firewall Rules:** (51 rules)

| Direction | Protocol | Src CIDR | Src Ports | Dst CIDR | Dst Ports | Action |
|-----------|----------|----------|-----------|----------|-----------|--------|
| inbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| inbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | deny |
| inbound | icmp | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| inbound | udp | 0.0.0.0/0 | 67 | 0.0.0.0/0 | 68 | allow |
| inbound | udp | 0.0.0.0/0 | * | 224.0.0.251/32 | 5353 | allow |
| inbound | udp | 0.0.0.0/0 | * | 239.255.255.250/32 | 1900 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 22 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 2049 | allow |
| inbound | udp | 0.0.0.0/0 | * | 0.0.0.0/0 | 2049 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 111 | allow |
| inbound | udp | 0.0.0.0/0 | * | 0.0.0.0/0 | 111 | allow |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 20048 | allow |
| inbound | udp | 10.0.0.0/16 | * | 0.0.0.0/0 | 20048 | allow |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 32803 | allow |
| inbound | udp | 10.0.0.0/16 | * | 0.0.0.0/0 | 32803 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 80 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 443 | deny |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 9100 | allow |
| inbound | udp | 10.0.0.0/16 | * | 0.0.0.0/0 | 9100 | allow |
| inbound | * | 10.0.0.0/16 | * | 0.0.0.0/0 | * | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 23 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 135 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 139 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 445 | deny |
| outbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| outbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| outbound | udp | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| inbound | * | ::/0 | * | ::/0 | * | allow |
| inbound | * | ::/0 | * | ::/0 | * | deny |
| inbound | icmpv6 | ::/0 | * | ::/0 | * | allow |
| inbound | icmpv6 | fe80::/10 | * | ::/0 | * | allow |
| inbound | udp | fe80::/10 | 547 | fe80::/10 | 546 | allow |
| inbound | udp | ::/0 | * | ff02::fb/128 | 5353 | allow |
| inbound | udp | ::/0 | * | ff02::f/128 | 1900 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 22 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 2049 | allow |
| inbound | udp | ::/0 | * | ::/0 | 2049 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 111 | allow |
| inbound | udp | ::/0 | * | ::/0 | 111 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 80 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 443 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 23 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 135 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 139 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 445 | deny |
| outbound | * | ::/0 | * | ::/0 | * | allow |
| outbound | * | ::/0 | * | ::/0 | * | deny |
| outbound | icmpv6 | ::/0 | * | ::/0 | * | allow |
| outbound | icmpv6 | fe80::/10 | * | ::/0 | * | allow |
| outbound | tcp | ::/0 | * | ::/0 | * | allow |
| outbound | udp | ::/0 | * | ::/0 | * | allow |

#### 3. ip-10-0-1-138

**Firewall Rules:** (53 rules)

| Direction | Protocol | Src CIDR | Src Ports | Dst CIDR | Dst Ports | Action |
|-----------|----------|----------|-----------|----------|-----------|--------|
| inbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| inbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | deny |
| inbound | icmp | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| inbound | udp | 0.0.0.0/0 | 67 | 0.0.0.0/0 | 68 | allow |
| inbound | udp | 0.0.0.0/0 | * | 224.0.0.251/32 | 5353 | allow |
| inbound | udp | 0.0.0.0/0 | * | 239.255.255.250/32 | 1900 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 22 | allow |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 3306 | allow |
| inbound | udp | 10.0.0.0/16 | * | 0.0.0.0/0 | 3306 | allow |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 4567 | allow |
| inbound | udp | 10.0.0.0/16 | * | 0.0.0.0/0 | 4567 | allow |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 4568 | allow |
| inbound | udp | 10.0.0.0/16 | * | 0.0.0.0/0 | 4568 | allow |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 4444 | allow |
| inbound | udp | 10.0.0.0/16 | * | 0.0.0.0/0 | 4444 | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 80 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 443 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 8080 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 3306 | deny |
| inbound | udp | 0.0.0.0/0 | * | 0.0.0.0/0 | 3306 | deny |
| inbound | tcp | 10.0.0.0/16 | * | 0.0.0.0/0 | 9104 | allow |
| inbound | udp | 10.0.0.0/16 | * | 0.0.0.0/0 | 9104 | allow |
| inbound | * | 10.0.0.0/16 | * | 0.0.0.0/0 | * | allow |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 23 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 135 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 139 | deny |
| inbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | 445 | deny |
| outbound | * | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| outbound | tcp | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| outbound | udp | 0.0.0.0/0 | * | 0.0.0.0/0 | * | allow |
| inbound | * | ::/0 | * | ::/0 | * | allow |
| inbound | * | ::/0 | * | ::/0 | * | deny |
| inbound | icmpv6 | ::/0 | * | ::/0 | * | allow |
| inbound | icmpv6 | fe80::/10 | * | ::/0 | * | allow |
| inbound | udp | fe80::/10 | 547 | fe80::/10 | 546 | allow |
| inbound | udp | ::/0 | * | ff02::fb/128 | 5353 | allow |
| inbound | udp | ::/0 | * | ff02::f/128 | 1900 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 22 | allow |
| inbound | tcp | ::/0 | * | ::/0 | 80 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 443 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 8080 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 3306 | deny |
| inbound | udp | ::/0 | * | ::/0 | 3306 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 23 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 135 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 139 | deny |
| inbound | tcp | ::/0 | * | ::/0 | 445 | deny |
| outbound | * | ::/0 | * | ::/0 | * | allow |
| outbound | * | ::/0 | * | ::/0 | * | deny |
| outbound | icmpv6 | ::/0 | * | ::/0 | * | allow |
| outbound | icmpv6 | fe80::/10 | * | ::/0 | * | allow |
| outbound | tcp | ::/0 | * | ::/0 | * | allow |
| outbound | udp | ::/0 | * | ::/0 | * | allow |


## Storage Resources

### Storage by Type

| Type | Total (GB) | Servers |
|------|------------|----------|
|  | 0 | 0 |

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

