# Target Cloud Infrastructure Summary

**Generated At:** 2026-03-25 09:26:25

**Namespace:** mig01

**MCI Name:** mmci01

---

## Overview

| Property             | Value                                                 |
| -------------------- | ----------------------------------------------------- |
| **MCI Name**         | mmci01                                                |
| **Description**      | Recommended VMs comprising multi-cloud infrastructure |
| **Status**           | Running:3 (R:3/3)                                     |
| **Target Cloud**     | AZURE                                                 |
| **Target Region**    | koreasouth                                            |
| **Total VMs**        | 3                                                     |
| **Running VMs**      | 3                                                     |
| **Stopped VMs**      | 0                                                     |
| **Monitoring Agent** |                                                       |

## Compute Resources

### VM Specifications

| Name              | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
| ----------------- | ----- | ------------ | --- | ------------ | --------- | --------------- | ------------------- |
| Standard_B2als_v2 | 2     | 3.9          | -   | x86_64       |           | $0.0432         | 1                   |
| Standard_B2as_v2  | 2     | 7.8          | -   | x86_64       |           | $0.0865         | 1                   |
| Standard_B4as_v2  | 4     | 15.6         | -   | x86_64       |           | $0.1730         | 1                   |

### VM Images

| Name                                                                              | Distribution                                            | OS Type      | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
| --------------------------------------------------------------------------------- | ------------------------------------------------------- | ------------ | ----------- | ------------ | -------------- | -------------- | -------------------- |
| Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310 | 0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2 | Ubuntu 22.04 | Linux/UNIX  | x86_64       | NA             | -              | 2                    |
| Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310      | 0001-com-ubuntu-server-jammy-daily:22_04-daily-lts      | Ubuntu 22.04 | Linux/UNIX  | x86_64       | NA             | -              | 1                    |

### Virtual Machines

| VM Name                                   | CSP VM ID                                                                                                                       | Status  | Spec (vCPU, Memory GiB) | Image                                                                                                             | Misc                                                                                                                                            |
| ----------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------- | ----------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| vm-ec268ed7-821e-9d73-e79f-961262161624-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8jg | Running | 2 vCPU, 3.9 GiB         | 0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2 (0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2) | **VNet:** vnet-01<br>**Subnet:** subnet-01<br>**Public IP:** 20.214.57.165<br>**Private IP:** 10.0.1.5<br>**SGs:** sg-01<br>**SSH:** sshkey-01  |
| vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8lg | Running | 2 vCPU, 7.8 GiB         | 0001-com-ubuntu-server-jammy-daily:22_04-daily-lts (0001-com-ubuntu-server-jammy-daily:22_04-daily-lts)           | **VNet:** vnet-01<br>**Subnet:** subnet-01<br>**Public IP:** 20.200.184.197<br>**Private IP:** 10.0.1.6<br>**SGs:** sg-03<br>**SSH:** sshkey-01 |
| vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8kg | Running | 4 vCPU, 15.6 GiB        | 0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2 (0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2) | **VNet:** vnet-01<br>**Subnet:** subnet-01<br>**Public IP:** 20.200.169.139<br>**Private IP:** 10.0.1.4<br>**SGs:** sg-02<br>**SSH:** sshkey-01 |

## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: vnet-01

| Property         | Value                                                                                                                           |
| ---------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| **Name**         | vnet-01                                                                                                                         |
| **CSP VNet ID**  | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg |
| **CIDR Block**   | 10.0.0.0/21                                                                                                                     |
| **Connection**   | azure-koreasouth                                                                                                                |
| **Subnet Count** | 1                                                                                                                               |

**Subnets:**

| Name      | CSP Subnet ID                                                                                                                                                | CIDR Block  | Zone |
| --------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------- | ---- |
| subnet-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg/subnets/d71qj0f693a119t3b8g0 | 10.0.1.0/24 |      |

## Security Resources

### SSH Keys

| Name      | CSP SSH Key ID                                                                                                                | Username | Fingerprint |
| --------- | ----------------------------------------------------------------------------------------------------------------------------- | -------- | ----------- |
| sshkey-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/d71qj2f693a119t3b8gg |          |             |

### Security Groups

#### Security Group: sg-01

| Property                  | Value                                                                                                                                 |
| ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| **Name**                  | sg-01                                                                                                                                 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/d71qj57693a119t3b8h0 |
| **VNet**                  | vnet-01                                                                                                                               |
| **Rule Count**            | 14 rules                                                                                                                              |

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
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |
| outbound  | ALL      |            | 0.0.0.0/0   |

#### Security Group: sg-02

| Property                  | Value                                                                                                                                 |
| ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| **Name**                  | sg-02                                                                                                                                 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/d71qj7n693a119t3b8hg |
| **VNet**                  | vnet-01                                                                                                                               |
| **Rule Count**            | 19 rules                                                                                                                              |

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
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |
| outbound  | ALL      |            | 0.0.0.0/0   |

#### Security Group: sg-03

| Property                  | Value                                                                                                                                 |
| ------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| **Name**                  | sg-03                                                                                                                                 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/d71qjaf693a119t3b8i0 |
| **VNet**                  | vnet-01                                                                                                                               |
| **Rule Count**            | 19 rules                                                                                                                              |

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
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |
| outbound  | ALL      |            | 0.0.0.0/0   |

## Cost Estimation

### Total Cost Summary

| Period                  | Cost (USD) |
| ----------------------- | ---------- |
| **Per Hour**            | $0.3027    |
| **Per Day**             | $7.26      |
| **Per Month (30 days)** | $217.94    |

### Cost by Region

| CSP   | Region     | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
| ----- | ---------- | -------- | --------------- | ---------------- |
| AZURE | koreasouth | 3        | $0.3027         | $217.94          |

### Cost by Virtual Machine

| VM Name                                   | Spec              | Cost/Hour (USD) | Cost/Month (USD) |
| ----------------------------------------- | ----------------- | --------------- | ---------------- |
| vm-ec268ed7-821e-9d73-e79f-961262161624-1 | Standard_B2als_v2 | $0.0432         | $31.10           |
| vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | Standard_B2as_v2  | $0.0865         | $62.28           |
| vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | Standard_B4as_v2  | $0.1730         | $124.56          |
