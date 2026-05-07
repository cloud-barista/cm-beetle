# Target Cloud Infrastructure Summary

**Generated At:** 2026-03-25 12:49:46

**Namespace:** mig01

**MCI Name:** mig-2-mci101

---

## Overview

| Property             | Value                                                 |
| -------------------- | ----------------------------------------------------- |
| **MCI Name**         | mig-2-mci101                                          |
| **Description**      | Recommended VMs comprising multi-cloud infrastructure |
| **Status**           | Running:3 (R:3/3)                                     |
| **Target Cloud**     | GCP                                                   |
| **Target Region**    | asia-northeast3                                       |
| **Total VMs**        | 3                                                     |
| **Running VMs**      | 3                                                     |
| **Stopped VMs**      | 0                                                     |
| **Monitoring Agent** |                                                       |

## Compute Resources

### VM Specifications

| Name          | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
| ------------- | ----- | ------------ | --- | ------------ | --------- | --------------- | ------------------- |
| e2-standard-2 | 2     | 7.8          | -   | x86_64       |           | $0.0860         | 1                   |
| e2-standard-4 | 4     | 15.6         | -   | x86_64       |           | $0.1719         | 1                   |
| e2-small      | 2     | 2.0          | -   | x86_64       |           | $0.0215         | 1                   |

### VM Images

| Name                                                                                                    | Distribution                                                        | OS Type      | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------ | ----------- | ------------ | -------------- | -------------- | -------------------- |
| https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260210 | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-02-10 | Ubuntu 22.04 | Linux/UNIX  | x86_64       | NA             | 10 GB          | 3                    |

### Virtual Machines

| VM Name                                         | CSP VM ID            | Status  | Spec (vCPU, Memory GiB) | Image                                                                                                                                     | Misc                                                                                                                                                                    |
| ----------------------------------------------- | -------------------- | ------- | ----------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| mig-2-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | d71tiu7693a119qk821g | Running | 2 vCPU, 2.0 GiB         | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-02-10 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-02-10) | **VNet:** mig-2-vnet-01<br>**Subnet:** mig-2-subnet-01<br>**Public IP:** 34.64.147.105<br>**Private IP:** 10.0.1.4<br>**SGs:** mig-2-sg-01<br>**SSH:** mig-2-sshkey-01  |
| mig-2-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | d71tiu7693a119qk823g | Running | 2 vCPU, 7.8 GiB         | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-02-10 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-02-10) | **VNet:** mig-2-vnet-01<br>**Subnet:** mig-2-subnet-01<br>**Public IP:** 34.158.212.143<br>**Private IP:** 10.0.1.3<br>**SGs:** mig-2-sg-03<br>**SSH:** mig-2-sshkey-01 |
| mig-2-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | d71tiu7693a119qk822g | Running | 4 vCPU, 15.6 GiB        | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-02-10 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-02-10) | **VNet:** mig-2-vnet-01<br>**Subnet:** mig-2-subnet-01<br>**Public IP:** 34.22.67.125<br>**Private IP:** 10.0.1.2<br>**SGs:** mig-2-sg-02<br>**SSH:** mig-2-sshkey-01   |

## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: mig-2-vnet-01

| Property         | Value                |
| ---------------- | -------------------- |
| **Name**         | mig-2-vnet-01        |
| **CSP VNet ID**  | d71tf97693a119qk816g |
| **CIDR Block**   | 10.0.0.0/21          |
| **Connection**   | gcp-asia-northeast3  |
| **Subnet Count** | 1                    |

**Subnets:**

| Name            | CSP Subnet ID        | CIDR Block  | Zone |
| --------------- | -------------------- | ----------- | ---- |
| mig-2-subnet-01 | d71tf97693a119qk8170 | 10.0.1.0/24 |      |

## Security Resources

### SSH Keys

| Name            | CSP SSH Key ID       | Username | Fingerprint |
| --------------- | -------------------- | -------- | ----------- |
| mig-2-sshkey-01 | d71tfjv693a119qk81ig |          |             |

### Security Groups

#### Security Group: mig-2-sg-01

| Property                  | Value                |
| ------------------------- | -------------------- |
| **Name**                  | mig-2-sg-01          |
| **CSP Security Group ID** | d71tfk7693a119qk81j0 |
| **VNet**                  | mig-2-vnet-01        |
| **Rule Count**            | 14 rules             |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR        |
| --------- | -------- | ---------- | ----------- |
| inbound   | ICMP     |            | 0.0.0.0/0   |
| inbound   | UDP      | 68         | 0.0.0.0/0   |
| inbound   | UDP      | 5353       | 0.0.0.0/0   |
| inbound   | UDP      | 1900       | 0.0.0.0/0   |
| inbound   | TCP      | 22         | 0.0.0.0/0   |
| inbound   | TCP      | 80         | 0.0.0.0/0   |
| inbound   | TCP      | 443        | 0.0.0.0/0   |
| inbound   | TCP      | 8080       | 0.0.0.0/0   |
| inbound   | TCP      | 9113       | 10.0.0.0/16 |
| inbound   | UDP      | 9113       | 10.0.0.0/16 |
| inbound   | ALL      |            | 10.0.0.0/16 |
| outbound  | ALL      |            | 0.0.0.0/0   |
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |

#### Security Group: mig-2-sg-02

| Property                  | Value                |
| ------------------------- | -------------------- |
| **Name**                  | mig-2-sg-02          |
| **CSP Security Group ID** | d71tggn693a119qk81qg |
| **VNet**                  | mig-2-vnet-01        |
| **Rule Count**            | 19 rules             |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR        |
| --------- | -------- | ---------- | ----------- |
| inbound   | ICMP     |            | 0.0.0.0/0   |
| inbound   | UDP      | 68         | 0.0.0.0/0   |
| inbound   | UDP      | 5353       | 0.0.0.0/0   |
| inbound   | UDP      | 1900       | 0.0.0.0/0   |
| inbound   | TCP      | 22         | 0.0.0.0/0   |
| inbound   | TCP      | 2049       | 0.0.0.0/0   |
| inbound   | UDP      | 2049       | 0.0.0.0/0   |
| inbound   | TCP      | 111        | 0.0.0.0/0   |
| inbound   | UDP      | 111        | 0.0.0.0/0   |
| inbound   | TCP      | 20048      | 10.0.0.0/16 |
| inbound   | UDP      | 20048      | 10.0.0.0/16 |
| inbound   | TCP      | 32803      | 10.0.0.0/16 |
| inbound   | UDP      | 32803      | 10.0.0.0/16 |
| inbound   | TCP      | 9100       | 10.0.0.0/16 |
| inbound   | UDP      | 9100       | 10.0.0.0/16 |
| inbound   | ALL      |            | 10.0.0.0/16 |
| outbound  | ALL      |            | 0.0.0.0/0   |
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |

#### Security Group: mig-2-sg-03

| Property                  | Value                |
| ------------------------- | -------------------- |
| **Name**                  | mig-2-sg-03          |
| **CSP Security Group ID** | d71thjn693a119qk8200 |
| **VNet**                  | mig-2-vnet-01        |
| **Rule Count**            | 19 rules             |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR        |
| --------- | -------- | ---------- | ----------- |
| inbound   | ICMP     |            | 0.0.0.0/0   |
| inbound   | UDP      | 68         | 0.0.0.0/0   |
| inbound   | UDP      | 5353       | 0.0.0.0/0   |
| inbound   | UDP      | 1900       | 0.0.0.0/0   |
| inbound   | TCP      | 22         | 0.0.0.0/0   |
| inbound   | TCP      | 3306       | 10.0.0.0/16 |
| inbound   | UDP      | 3306       | 10.0.0.0/16 |
| inbound   | TCP      | 4567       | 10.0.0.0/16 |
| inbound   | UDP      | 4567       | 10.0.0.0/16 |
| inbound   | TCP      | 4568       | 10.0.0.0/16 |
| inbound   | UDP      | 4568       | 10.0.0.0/16 |
| inbound   | TCP      | 4444       | 10.0.0.0/16 |
| inbound   | UDP      | 4444       | 10.0.0.0/16 |
| inbound   | TCP      | 9104       | 10.0.0.0/16 |
| inbound   | UDP      | 9104       | 10.0.0.0/16 |
| inbound   | ALL      |            | 10.0.0.0/16 |
| outbound  | ALL      |            | 0.0.0.0/0   |
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |

## Cost Estimation

### Total Cost Summary

| Period                  | Cost (USD) |
| ----------------------- | ---------- |
| **Per Hour**            | $0.2794    |
| **Per Day**             | $6.71      |
| **Per Month (30 days)** | $201.16    |

### Cost by Region

| CSP | Region          | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
| --- | --------------- | -------- | --------------- | ---------------- |
| GCP | asia-northeast3 | 3        | $0.2794         | $201.16          |

### Cost by Virtual Machine

| VM Name                                         | Spec          | Cost/Hour (USD) | Cost/Month (USD) |
| ----------------------------------------------- | ------------- | --------------- | ---------------- |
| mig-2-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | e2-small      | $0.0215         | $15.47           |
| mig-2-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | e2-standard-2 | $0.0860         | $61.90           |
| mig-2-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | e2-standard-4 | $0.1719         | $123.79          |
