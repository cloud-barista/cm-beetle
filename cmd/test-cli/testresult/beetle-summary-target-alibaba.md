# Target Cloud Infrastructure Summary

**Generated At:** 2026-03-25 13:31:33

**Namespace:** mig01

**MCI Name:** mig-4-mci101

---

## Overview

| Property             | Value                                                 |
| -------------------- | ----------------------------------------------------- |
| **MCI Name**         | mig-4-mci101                                          |
| **Description**      | Recommended VMs comprising multi-cloud infrastructure |
| **Status**           | Failed:3 (R:0/3)                                      |
| **Target Cloud**     | ALIBABA                                               |
| **Target Region**    |                                                       |
| **Total VMs**        | 3                                                     |
| **Running VMs**      | 0                                                     |
| **Stopped VMs**      | 0                                                     |
| **Monitoring Agent** |                                                       |

## Compute Resources

### VM Specifications

| Name               | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
| ------------------ | ----- | ------------ | --- | ------------ | --------- | --------------- | ------------------- |
| ecs.t6-c1m4.xlarge | 4     | 16.0         | -   | x86_64       |           | $0.1693         | 1                   |
| ecs.t6-c1m1.large  | 2     | 2.0          | -   | x86_64       |           | $0.0214         | 1                   |
| ecs.t6-c1m4.large  | 2     | 8.0          | -   | x86_64       |           | $0.0846         | 1                   |

### VM Images

| Name                                      | Distribution        | OS Type      | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
| ----------------------------------------- | ------------------- | ------------ | ----------- | ------------ | -------------- | -------------- | -------------------- |
| ubuntu_22_04_x64_20G_alibase_20260119.vhd | Ubuntu 22.04 64 bit | Ubuntu 22.04 | Linux/UNIX  | x86_64       | NA             | 20 GB          | 3                    |

### Virtual Machines

| VM Name                                         | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image                                     | Misc                                                                                                                                              |
| ----------------------------------------------- | --------- | ------ | ----------------------- | ----------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1 |           | Failed | 2 vCPU, 2.0 GiB         | Ubuntu 22.04 64 bit (Ubuntu 22.04 64 bit) | **VNet:** mig-4-vnet-01<br>**Subnet:** mig-4-subnet-01<br>**Public IP:** <br>**Private IP:** <br>**SGs:** mig-4-sg-01<br>**SSH:** mig-4-sshkey-01 |
| mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 |           | Failed | 2 vCPU, 8.0 GiB         | Ubuntu 22.04 64 bit (Ubuntu 22.04 64 bit) | **VNet:** mig-4-vnet-01<br>**Subnet:** mig-4-subnet-01<br>**Public IP:** <br>**Private IP:** <br>**SGs:** mig-4-sg-03<br>**SSH:** mig-4-sshkey-01 |
| mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 |           | Failed | 4 vCPU, 16.0 GiB        | Ubuntu 22.04 64 bit (Ubuntu 22.04 64 bit) | **VNet:** mig-4-vnet-01<br>**Subnet:** mig-4-subnet-01<br>**Public IP:** <br>**Private IP:** <br>**SGs:** mig-4-sg-02<br>**SSH:** mig-4-sshkey-01 |

## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: mig-4-vnet-01

| Property         | Value                     |
| ---------------- | ------------------------- |
| **Name**         | mig-4-vnet-01             |
| **CSP VNet ID**  | vpc-mj72oyy2s6tly01dcchqs |
| **CIDR Block**   | 10.0.0.0/21               |
| **Connection**   | alibaba-ap-northeast-2    |
| **Subnet Count** | 1                         |

**Subnets:**

| Name            | CSP Subnet ID             | CIDR Block  | Zone            |
| --------------- | ------------------------- | ----------- | --------------- |
| mig-4-subnet-01 | vsw-mj782snkv4i7r0see4na0 | 10.0.1.0/24 | ap-northeast-2a |

## Security Resources

### SSH Keys

| Name            | CSP SSH Key ID       | Username | Fingerprint                      |
| --------------- | -------------------- | -------- | -------------------------------- |
| mig-4-sshkey-01 | d71u75v693a119sjo6g0 |          | e88593506421268e6551072903407d33 |

### Security Groups

#### Security Group: mig-4-sg-01

| Property                  | Value                   |
| ------------------------- | ----------------------- |
| **Name**                  | mig-4-sg-01             |
| **CSP Security Group ID** | sg-mj7bmek530yys7apfwdi |
| **VNet**                  | mig-4-vnet-01           |
| **Rule Count**            | 14 rules                |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR        |
| --------- | -------- | ---------- | ----------- |
| inbound   | ALL      |            | 10.0.0.0/16 |
| inbound   | UDP      | 9113       | 10.0.0.0/16 |
| inbound   | TCP      | 9113       | 10.0.0.0/16 |
| inbound   | TCP      | 8080       | 0.0.0.0/0   |
| inbound   | TCP      | 443        | 0.0.0.0/0   |
| inbound   | TCP      | 80         | 0.0.0.0/0   |
| inbound   | TCP      | 22         | 0.0.0.0/0   |
| inbound   | UDP      | 1900       | 0.0.0.0/0   |
| inbound   | UDP      | 5353       | 0.0.0.0/0   |
| inbound   | UDP      | 68         | 0.0.0.0/0   |
| inbound   | ICMP     |            | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | ALL      |            | 0.0.0.0/0   |

#### Security Group: mig-4-sg-02

| Property                  | Value                   |
| ------------------------- | ----------------------- |
| **Name**                  | mig-4-sg-02             |
| **CSP Security Group ID** | sg-mj7eu0g9c9oo47ounzk6 |
| **VNet**                  | mig-4-vnet-01           |
| **Rule Count**            | 19 rules                |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR        |
| --------- | -------- | ---------- | ----------- |
| inbound   | ALL      |            | 10.0.0.0/16 |
| inbound   | UDP      | 9100       | 10.0.0.0/16 |
| inbound   | TCP      | 9100       | 10.0.0.0/16 |
| inbound   | UDP      | 32803      | 10.0.0.0/16 |
| inbound   | TCP      | 32803      | 10.0.0.0/16 |
| inbound   | UDP      | 20048      | 10.0.0.0/16 |
| inbound   | TCP      | 20048      | 10.0.0.0/16 |
| inbound   | UDP      | 111        | 0.0.0.0/0   |
| inbound   | TCP      | 111        | 0.0.0.0/0   |
| inbound   | UDP      | 2049       | 0.0.0.0/0   |
| inbound   | TCP      | 2049       | 0.0.0.0/0   |
| inbound   | TCP      | 22         | 0.0.0.0/0   |
| inbound   | UDP      | 1900       | 0.0.0.0/0   |
| inbound   | UDP      | 5353       | 0.0.0.0/0   |
| inbound   | UDP      | 68         | 0.0.0.0/0   |
| inbound   | ICMP     |            | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | ALL      |            | 0.0.0.0/0   |

#### Security Group: mig-4-sg-03

| Property                  | Value                   |
| ------------------------- | ----------------------- |
| **Name**                  | mig-4-sg-03             |
| **CSP Security Group ID** | sg-mj7i5crahq8i5rzcer4u |
| **VNet**                  | mig-4-vnet-01           |
| **Rule Count**            | 19 rules                |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR        |
| --------- | -------- | ---------- | ----------- |
| inbound   | ALL      |            | 10.0.0.0/16 |
| inbound   | UDP      | 9104       | 10.0.0.0/16 |
| inbound   | TCP      | 9104       | 10.0.0.0/16 |
| inbound   | UDP      | 4444       | 10.0.0.0/16 |
| inbound   | TCP      | 4444       | 10.0.0.0/16 |
| inbound   | UDP      | 4568       | 10.0.0.0/16 |
| inbound   | TCP      | 4568       | 10.0.0.0/16 |
| inbound   | UDP      | 4567       | 10.0.0.0/16 |
| inbound   | TCP      | 4567       | 10.0.0.0/16 |
| inbound   | UDP      | 3306       | 10.0.0.0/16 |
| inbound   | TCP      | 3306       | 10.0.0.0/16 |
| inbound   | TCP      | 22         | 0.0.0.0/0   |
| inbound   | UDP      | 1900       | 0.0.0.0/0   |
| inbound   | UDP      | 5353       | 0.0.0.0/0   |
| inbound   | UDP      | 68         | 0.0.0.0/0   |
| inbound   | ICMP     |            | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | ALL      |            | 0.0.0.0/0   |

## Cost Estimation

### Total Cost Summary

| Period                  | Cost (USD) |
| ----------------------- | ---------- |
| **Per Hour**            | $0.0000    |
| **Per Day**             | $0.00      |
| **Per Month (30 days)** | $0.00      |

### Cost by Region

| CSP     | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
| ------- | ------ | -------- | --------------- | ---------------- |
| ALIBABA |        | 3        | $0.0000         | $0.00            |

### Cost by Virtual Machine

| VM Name                                         | Spec | Cost/Hour (USD) | Cost/Month (USD) |
| ----------------------------------------------- | ---- | --------------- | ---------------- |
| mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1 |      | $0.0000         | $0.00            |
| mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 |      | $0.0000         | $0.00            |
| mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 |      | $0.0000         | $0.00            |
