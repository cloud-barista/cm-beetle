# Target Cloud Infrastructure Summary

**Generated At:** 2025-12-10 07:24:02

**Namespace:** mig01

**MCI Name:** mmci01

---

## Overview

| Property | Value |
|----------|-------|
| **MCI Name** | mmci01 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | AZURE |
| **Target Region** | koreasouth |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| Standard_B2als_v2 | 2 | 3.9 | - | x86_64 |  | $0.0432 | 1 |
| Standard_B2as_v2 | 2 | 7.8 | - | x86_64 |  | $0.0865 | 1 |
| Standard_B4as_v2 | 4 | 15.6 | - | x86_64 |  | $0.1730 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230 | 0001-com-ubuntu-server-jammy:22_04-lts-gen2 | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| migrated-ec268ed7-821e-9d73-e79f-961262161624-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4shuti5npi2mbdfo080 | Running | 2 vCPU, 3.9 GiB | 0001-com-ubuntu-server-jammy:22_04-lts-gen2 (0001-com-ubuntu-server-jammy:22_04-lts-gen2) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 52.231.218.136<br>**Private IP:** 10.0.1.5<br>**SGs:** mig-sg-01<br>**SSH:** mig-sshkey-01 |
| migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4shuti5npi2mbdfo0a0 | Running | 2 vCPU, 7.8 GiB | 0001-com-ubuntu-server-jammy:22_04-lts-gen2 (0001-com-ubuntu-server-jammy:22_04-lts-gen2) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 52.231.229.168<br>**Private IP:** 10.0.1.4<br>**SGs:** mig-sg-03<br>**SSH:** mig-sshkey-01 |
| migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4shuti5npi2mbdfo090 | Running | 4 vCPU, 15.6 GiB | 0001-com-ubuntu-server-jammy:22_04-lts-gen2 (0001-com-ubuntu-server-jammy:22_04-lts-gen2) | **VNet:** mig-vnet-01<br>**Subnet:** mig-subnet-01<br>**Public IP:** 52.231.218.239<br>**Private IP:** 10.0.1.6<br>**SGs:** mig-sg-02<br>**SSH:** mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | mig-vnet-01 |
| **CSP VNet ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4shttq5npi2mbdfo040 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | azure-koreasouth |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| mig-subnet-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4shttq5npi2mbdfo040/subnets/d4shttq5npi2mbdfo04g | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| mig-sshkey-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/d4shu0i5npi2mbdfo050 |  |  |

### Security Groups

#### Security Group: mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | mig-sg-01 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/d4shu6q5npi2mbdfo05g |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 14 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 80 | 0.0.0.0/0 |
| inbound | TCP | 443 | 0.0.0.0/0 |
| inbound | TCP | 8080 | 0.0.0.0/0 |
| inbound | TCP | 9113 | 10.0.0.0/16 |
| inbound | UDP | 9113 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | mig-sg-02 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/d4shueq5npi2mbdfo060 |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 2049 | 0.0.0.0/0 |
| inbound | UDP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 111 | 0.0.0.0/0 |
| inbound | UDP | 111 | 0.0.0.0/0 |
| inbound | TCP | 20048 | 10.0.0.0/16 |
| inbound | UDP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 32803 | 10.0.0.0/16 |
| inbound | UDP | 32803 | 10.0.0.0/16 |
| inbound | TCP | 9100 | 10.0.0.0/16 |
| inbound | UDP | 9100 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: mig-sg-03

| Property | Value |
|----------|-------|
| **Name** | mig-sg-03 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/d4shum25npi2mbdfo06g |
| **VNet** | mig-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 3306 | 10.0.0.0/16 |
| inbound | UDP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 4567 | 10.0.0.0/16 |
| inbound | UDP | 4567 | 10.0.0.0/16 |
| inbound | TCP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4568 | 10.0.0.0/16 |
| inbound | TCP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 4444 | 10.0.0.0/16 |
| inbound | TCP | 9104 | 10.0.0.0/16 |
| inbound | UDP | 9104 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.3027 |
| **Per Day** | $7.26 |
| **Per Month (30 days)** | $217.94 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AZURE | koreasouth | 3 | $0.3027 | $217.94 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| migrated-ec268ed7-821e-9d73-e79f-961262161624-1 | Standard_B2als_v2 | $0.0432 | $31.10 |
| migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | Standard_B2as_v2 | $0.0865 | $62.28 |
| migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | Standard_B4as_v2 | $0.1730 | $124.56 |


