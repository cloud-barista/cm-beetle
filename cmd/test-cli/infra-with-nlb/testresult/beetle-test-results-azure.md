# CM-Beetle test results for AZURE (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.5+ (83361e8)
- imdl: v0.1.22+ (83361e8)
- CB-Tumblebug: v0.14.0
- CB-Spider: v0.14.0
- CB-MapUI: v0.14.0
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: November 25, 2026
- Test Time: 18:14:41 KST
- Test Execution: 2026-11-25 18:14:41 KST

### Scenario

1. Recommend target model for computing infra with NLB via Beetle
1. Validate the target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Remote Command Accessibility Check
1. Migrate NLBs to the cloud infra via Beetle
1. Get a list of migrated NLBs via Beetle
1. Get details of a specific migrated NLB via Beetle
1. NLB Load Balancing Verification
1. Target Infrastructure Summary via Beetle
1. Migration Report via Beetle
1. Delete the migrated NLBs via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for AZURE

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 2.634s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 444ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m25.228s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 80ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 8ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 52ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 1.1s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 8.216s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 18ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 17ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 5m56.303s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.476s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.226s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 39.084s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m39.67s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 10m53.928023114s

*Test executed on November 25, 2026 at 18:14:41 KST (2026-11-25 18:14:41 KST) using CM-Beetle automated test CLI*

---

## Detailed Test Case Results

> [!INFO]
> This section provides detailed information for each test case, including API request information and response details.

### Test Case 1: Recommend target model for computing infra with NLB

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/recommendation/infraWithNlb`
- **Purpose**: Get NLB-aware infrastructure recommendations for migration
- **Required Parameters**: `desiredCsp` and `desiredRegion` in request body

**Request Body**:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "lineage": {},
  "desiredCsp": "azure",
  "desiredRegion": "koreasouth",
  "sourceInfra": {
    "network": {
      "ipv4Networks": {
        "defaultGateways": [
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec268ed7-821e-9d73-e79f-961262161624"
          },
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
          },
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
          }
        ]
      },
      "ipv6Networks": {}
    },
    "nodes": [
      {
        "hostname": "ip-10-0-1-30",
        "machineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 1,
          "cores": 1,
          "threads": 2,
          "maxSpeed": 2.499,
          "vendor": "GenuineIntel",
          "model": "Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz"
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 2,
          "available": 1
        },
        "rootDisk": {
          "label": "",
          "type": "SSD",
          "totalSize": 8
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": [
              "127.0.0.1/8"
            ],
            "ipv6CidrBlocks": [
              "::1/128"
            ],
            "mtu": 65536,
            "state": "up"
          },
          {
            "name": "ens5",
            "macAddress": "02:6f:de:fc:71:b1",
            "ipv4CidrBlocks": [
              "10.0.1.30/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::6f:deff:fefc:71b1/64"
            ],
            "mtu": 9001,
            "state": "up"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0/0",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "10.0.1.0/24",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "firewallTable": [
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "22",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9999",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "allow"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.3 LTS",
          "version": "22.04.3 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "ip-10-0-1-221",
        "machineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 1,
          "cores": 2,
          "threads": 4,
          "maxSpeed": 2.499,
          "vendor": "GenuineIntel",
          "model": "Intel(R) Xeon(R) Platinum 8175M CPU @ 2.50GHz"
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 16,
          "available": 15
        },
        "rootDisk": {
          "label": "",
          "type": "SSD",
          "totalSize": 30
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": [
              "127.0.0.1/8"
            ],
            "ipv6CidrBlocks": [
              "::1/128"
            ],
            "mtu": 65536,
            "state": "up"
          },
          {
            "name": "ens5",
            "macAddress": "02:08:96:7d:f4:17",
            "ipv4CidrBlocks": [
              "10.0.1.221/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::8:96ff:fe7d:f417/64"
            ],
            "mtu": 9001,
            "state": "up"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0/0",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "10.0.1.0/24",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "firewallTable": [
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "22",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8086",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "allow"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.3 LTS",
          "version": "22.04.3 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "ip-10-0-1-138",
        "machineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 1,
          "cores": 2,
          "threads": 4,
          "maxSpeed": 2.499,
          "vendor": "GenuineIntel",
          "model": "Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz"
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 8,
          "available": 7
        },
        "rootDisk": {
          "label": "",
          "type": "SSD",
          "totalSize": 30
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": [
              "127.0.0.1/8"
            ],
            "ipv6CidrBlocks": [
              "::1/128"
            ],
            "mtu": 65536,
            "state": "up"
          },
          {
            "name": "ens5",
            "macAddress": "02:bf:6e:6c:6e:31",
            "ipv4CidrBlocks": [
              "10.0.1.138/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::bf:6eff:fe6c:6e31/64"
            ],
            "mtu": 9001,
            "state": "up"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0/0",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "10.0.1.0/24",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "firewallTable": [
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "22",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8086",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "allow"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.3 LTS",
          "version": "22.04.3 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      }
    ],
    "nlbs": [
      {
        "hostMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "software": "haproxy",
        "listener": {
          "bindAddress": "*",
          "port": 9999,
          "protocol": "tcp"
        },
        "backend": {
          "name": "influxdb_back",
          "balance": "roundrobin",
          "protocol": "tcp",
          "servers": [
            {
              "name": "influx1",
              "ip": "10.0.1.221",
              "port": 8086
            },
            {
              "name": "influx2",
              "ip": "10.0.1.138",
              "port": 8086
            }
          ]
        },
        "healthCheck": {
          "enabled": true,
          "interval": 10,
          "timeout": 10,
          "threshold": 3
        }
      }
    ]
  }
}
```

</details>

#### 1.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infrastructure recommendation generated successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "success": true,
  "data": [
    {
      "status": "partially-matched",
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=86.1% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "azure",
        "region": "koreasouth"
      },
      "targetInfra": {
        "name": "infra101",
        "label": null,
        "systemLabel": "",
        "description": "NLB-aware recommended infrastructure for cloud migration",
        "nodeGroups": [
          {
            "name": "ng-influxdb-back",
            "nodeGroupSize": 2,
            "label": {
              "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
              "nlbBackend": "influxdb_back"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_d4_v4",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 30,
            "dataDiskIds": null
          },
          {
            "name": "ng-ip-10-0-1-30-62161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=83.3%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_f2s_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 30,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "mig-vnet-01",
        "connectionName": "azure-koreasouth",
        "cidrBlock": "10.0.0.0/21",
        "subnetInfoList": [
          {
            "name": "mig-subnet-01",
            "ipv4_CIDR": "10.0.1.0/24",
            "description": "a recommended subnet for migration"
          }
        ],
        "description": "a recommended vNet for migration"
      },
      "targetSshKey": {
        "name": "mig-sshkey-01",
        "connectionName": "azure-koreasouth",
        "description": "SSH key pair for migration (Note: provided ONLY once, MUST be downloaded)",
        "cspResourceId": "",
        "fingerprint": "",
        "username": "",
        "verifiedUsername": "",
        "publicKey": "",
        "privateKey": ""
      },
      "targetSpecList": [
        {
          "id": "azure+koreasouth+standard_d4_v4",
          "uid": "tb2f522bo2kgbr6t0n9g",
          "cspSpecName": "Standard_D4_v4",
          "name": "azure+koreasouth+standard_d4_v4",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "costPerHour": 0.221,
          "orderInFilteredResult": 30,
          "evaluationScore01": -1,
          "evaluationScore02": -1,
          "evaluationScore03": -1,
          "evaluationScore04": -1,
          "evaluationScore05": -1,
          "evaluationScore06": -1,
          "evaluationScore07": -1,
          "evaluationScore08": -1,
          "evaluationScore09": -1,
          "evaluationScore10": -1,
          "rootDiskType": "",
          "rootDiskSize": 0,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MaxDataDiskCount",
              "value": "8"
            },
            {
              "key": "MemoryInMB",
              "value": "16384"
            },
            {
              "key": "Name",
              "value": "Standard_D4_v4"
            },
            {
              "key": "NumberOfCores",
              "value": "4"
            },
            {
              "key": "OSDiskSizeInMB",
              "value": "1047552"
            },
            {
              "key": "ResourceDiskSizeInMB",
              "value": "0"
            },
            {
              "key": "MaxResourceVolumeMB",
              "value": "0"
            },
            {
              "key": "OSVhdSizeMB",
              "value": "1047552"
            },
            {
              "key": "vCPUs",
              "value": "4"
            },
            {
              "key": "MemoryPreservingMaintenanceSupported",
              "value": "True"
            },
            {
              "key": "HyperVGenerations",
              "value": "V1,V2"
            },
            {
              "key": "SupportedCapacityReservationTypes",
              "value": "Open,Targeted"
            },
            {
              "key": "MemoryGB",
              "value": "16"
            },
            {
              "key": "MaxDataDiskCount",
              "value": "8"
            },
            {
              "key": "CpuArchitectureType",
              "value": "x64"
            },
            {
              "key": "LowPriorityCapable",
              "value": "True"
            },
            {
              "key": "PremiumIO",
              "value": "False"
            },
            {
              "key": "VMDeploymentTypes",
              "value": "IaaS"
            },
            {
              "key": "vCPUsConstraintsAllowed",
              "value": "1, 2, 4"
            },
            {
              "key": "vCPUsAvailable",
              "value": "4"
            },
            {
              "key": "ACUs",
              "value": "195"
            },
            {
              "key": "vCPUsPerCore",
              "value": "2"
            },
            {
              "key": "CombinedTempDiskAndCachedIOPS",
              "value": "38500"
            },
            {
              "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
              "value": "242221056"
            },
            {
              "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
              "value": "242221056"
            },
            {
              "key": "UncachedDiskIOPS",
              "value": "6400"
            },
            {
              "key": "UncachedDiskBytesPerSecond",
              "value": "96000000"
            },
            {
              "key": "EphemeralOSDiskSupported",
              "value": "False"
            },
            {
              "key": "EncryptionAtHostSupported",
              "value": "False"
            },
            {
              "key": "CapacityReservationSupported",
              "value": "True"
            },
            {
              "key": "AcceleratedNetworkingEnabled",
              "value": "True"
            },
            {
              "key": "RdmaEnabled",
              "value": "False"
            },
            {
              "key": "MaxNetworkInterfaces",
              "value": "2"
            },
            {
              "key": "LocationInfo_0_Location",
              "value": "KoreaSouth"
            },
            {
              "key": "Family",
              "value": "standardDv4Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "D4_v4"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ],
          "availabilityStatus": "Unverified"
        },
        {
          "id": "azure+koreasouth+standard_f2s_v2",
          "uid": "tbr3k3hgqboa0n905jlg",
          "cspSpecName": "Standard_F2s_v2",
          "name": "azure+koreasouth+standard_f2s_v2",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 4,
          "diskSizeGB": 17,
          "costPerHour": 0.0961,
          "orderInFilteredResult": 14,
          "evaluationScore01": -1,
          "evaluationScore02": -1,
          "evaluationScore03": -1,
          "evaluationScore04": -1,
          "evaluationScore05": -1,
          "evaluationScore06": -1,
          "evaluationScore07": -1,
          "evaluationScore08": -1,
          "evaluationScore09": -1,
          "evaluationScore10": -1,
          "rootDiskType": "",
          "rootDiskSize": 17,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MaxDataDiskCount",
              "value": "4"
            },
            {
              "key": "MemoryInMB",
              "value": "4096"
            },
            {
              "key": "Name",
              "value": "Standard_F2s_v2"
            },
            {
              "key": "NumberOfCores",
              "value": "2"
            },
            {
              "key": "OSDiskSizeInMB",
              "value": "1047552"
            },
            {
              "key": "ResourceDiskSizeInMB",
              "value": "16384"
            },
            {
              "key": "MaxResourceVolumeMB",
              "value": "16384"
            },
            {
              "key": "OSVhdSizeMB",
              "value": "1047552"
            },
            {
              "key": "vCPUs",
              "value": "2"
            },
            {
              "key": "MemoryPreservingMaintenanceSupported",
              "value": "True"
            },
            {
              "key": "HyperVGenerations",
              "value": "V1,V2"
            },
            {
              "key": "SupportedCapacityReservationTypes",
              "value": "Open,Targeted"
            },
            {
              "key": "SupportedEphemeralOSDiskPlacements",
              "value": "ResourceDisk,CacheDisk"
            },
            {
              "key": "MemoryGB",
              "value": "4"
            },
            {
              "key": "MaxDataDiskCount",
              "value": "4"
            },
            {
              "key": "CpuArchitectureType",
              "value": "x64"
            },
            {
              "key": "LowPriorityCapable",
              "value": "True"
            },
            {
              "key": "PremiumIO",
              "value": "True"
            },
            {
              "key": "VMDeploymentTypes",
              "value": "IaaS"
            },
            {
              "key": "vCPUsConstraintsAllowed",
              "value": "1, 2"
            },
            {
              "key": "vCPUsAvailable",
              "value": "2"
            },
            {
              "key": "ACUs",
              "value": "195"
            },
            {
              "key": "vCPUsPerCore",
              "value": "2"
            },
            {
              "key": "CombinedTempDiskAndCachedIOPS",
              "value": "4000"
            },
            {
              "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
              "value": "32768000"
            },
            {
              "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
              "value": "32768000"
            },
            {
              "key": "CachedDiskBytes",
              "value": "34359738368"
            },
            {
              "key": "UncachedDiskIOPS",
              "value": "3200"
            },
            {
              "key": "UncachedDiskBytesPerSecond",
              "value": "48000000"
            },
            {
              "key": "RetirementDateUtc",
              "value": "11/15/2028"
            },
            {
              "key": "EphemeralOSDiskSupported",
              "value": "True"
            },
            {
              "key": "EncryptionAtHostSupported",
              "value": "True"
            },
            {
              "key": "CapacityReservationSupported",
              "value": "True"
            },
            {
              "key": "AcceleratedNetworkingEnabled",
              "value": "True"
            },
            {
              "key": "RdmaEnabled",
              "value": "False"
            },
            {
              "key": "MaxNetworkInterfaces",
              "value": "2"
            },
            {
              "key": "UltraSSDAvailable",
              "value": "True"
            },
            {
              "key": "LocationInfo_0_Location",
              "value": "KoreaSouth"
            },
            {
              "key": "Family",
              "value": "standardFSv2Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "F2s_v2"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ],
          "availabilityStatus": "Unverified"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "uid": "tbmm9857u5pvqm24elb2",
          "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:34:25 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "osDiskType": "default",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Location",
              "value": "australiacentral"
            },
            {
              "key": "Publisher",
              "value": "Canonical"
            },
            {
              "key": "Offer",
              "value": "0001-com-ubuntu-server-jammy"
            },
            {
              "key": "SKU",
              "value": "22_04-lts-gen2"
            },
            {
              "key": "Version",
              "value": "22.04.202610220"
            },
            {
              "key": "ID",
              "value": "/Subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202610220"
            },
            {
              "key": "HyperVGeneration",
              "value": "V2"
            },
            {
              "key": "Features",
              "value": "SecurityType=TrustedLaunchSupported, IsAcceleratedNetworkSupported=True, DiskControllerTypes=SCSI, NVMe, IsHibernateSupported=True"
            },
            {
              "key": "FeatureCount",
              "value": "4"
            },
            {
              "key": "ImageDeprecationState",
              "value": "Active"
            }
          ],
          "systemLabel": "from-assets",
          "description": "Server LTS Gen2",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "azure-koreasouth",
          "vNetId": "mig-vnet-01",
          "description": "Recommended security group for NLB backend influxdb_back",
          "firewallRules": [
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "8086",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "",
              "Protocol": "ALL",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9999",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "8086",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            }
          ],
          "cspResourceId": ""
        },
        {
          "name": "mig-sg-02",
          "connectionName": "azure-koreasouth",
          "vNetId": "mig-vnet-01",
          "description": "Recommended security group for ec268ed7-821e-9d73-e79f-961262161624",
          "firewallRules": [
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "9999",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "",
              "Protocol": "ALL",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            }
          ],
          "cspResourceId": ""
        }
      ],
      "targetNlbList": [
        {
          "description": "Migrated from HAProxy backend: influxdb_back",
          "type": "PUBLIC",
          "scope": "REGION",
          "listener": {
            "protocol": "TCP",
            "port": "9999"
          },
          "targetGroup": {
            "protocol": "TCP",
            "port": "8086",
            "nodeGroupId": "ng-influxdb-back"
          },
          "healthChecker": {
            "interval": 10,
            "threshold": 3,
            "timeout": -1
          }
        }
      ],
      "targetK8sCluster": {
        "connectionName": "",
        "description": "",
        "name": "",
        "version": "",
        "vNetId": "",
        "subnetIds": null,
        "securityGroupIds": null,
        "k8sNodeGroupList": null,
        "cspResourceId": "",
        "label": null,
        "systemLabel": ""
      }
    },
    {
      "status": "partially-matched",
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=25.0% Max=100.0% Avg=81.9% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "azure",
        "region": "koreasouth"
      },
      "targetInfra": {
        "name": "infra101",
        "label": null,
        "systemLabel": "",
        "description": "NLB-aware recommended infrastructure for cloud migration",
        "nodeGroups": [
          {
            "name": "ng-influxdb-back",
            "nodeGroupSize": 2,
            "label": {
              "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
              "nlbBackend": "influxdb_back"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_d4_v5",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 30,
            "dataDiskIds": null
          },
          {
            "name": "ng-ip-10-0-1-30-62161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=25.0% Image=83.3%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_d2_v5",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 30,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "mig-vnet-01",
        "connectionName": "azure-koreasouth",
        "cidrBlock": "10.0.0.0/21",
        "subnetInfoList": [
          {
            "name": "mig-subnet-01",
            "ipv4_CIDR": "10.0.1.0/24",
            "description": "a recommended subnet for migration"
          }
        ],
        "description": "a recommended vNet for migration"
      },
      "targetSshKey": {
        "name": "mig-sshkey-01",
        "connectionName": "azure-koreasouth",
        "description": "SSH key pair for migration (Note: provided ONLY once, MUST be downloaded)",
        "cspResourceId": "",
        "fingerprint": "",
        "username": "",
        "verifiedUsername": "",
        "publicKey": "",
        "privateKey": ""
      },
      "targetSpecList": [
        {
          "id": "azure+koreasouth+standard_d4_v5",
          "uid": "tbqo5gtrl1h7vt8u5d1b",
          "cspSpecName": "Standard_D4_v5",
          "name": "azure+koreasouth+standard_d4_v5",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "costPerHour": 0.221,
          "orderInFilteredResult": 29,
          "evaluationScore01": -1,
          "evaluationScore02": -1,
          "evaluationScore03": -1,
          "evaluationScore04": -1,
          "evaluationScore05": -1,
          "evaluationScore06": -1,
          "evaluationScore07": -1,
          "evaluationScore08": -1,
          "evaluationScore09": -1,
          "evaluationScore10": -1,
          "rootDiskType": "",
          "rootDiskSize": 0,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MaxDataDiskCount",
              "value": "8"
            },
            {
              "key": "MemoryInMB",
              "value": "16384"
            },
            {
              "key": "Name",
              "value": "Standard_D4_v5"
            },
            {
              "key": "NumberOfCores",
              "value": "4"
            },
            {
              "key": "OSDiskSizeInMB",
              "value": "1047552"
            },
            {
              "key": "ResourceDiskSizeInMB",
              "value": "0"
            },
            {
              "key": "MaxResourceVolumeMB",
              "value": "0"
            },
            {
              "key": "OSVhdSizeMB",
              "value": "1047552"
            },
            {
              "key": "vCPUs",
              "value": "4"
            },
            {
              "key": "MemoryPreservingMaintenanceSupported",
              "value": "True"
            },
            {
              "key": "HyperVGenerations",
              "value": "V1,V2"
            },
            {
              "key": "SupportedCapacityReservationTypes",
              "value": "Open,Targeted"
            },
            {
              "key": "MemoryGB",
              "value": "16"
            },
            {
              "key": "MaxDataDiskCount",
              "value": "8"
            },
            {
              "key": "CpuArchitectureType",
              "value": "x64"
            },
            {
              "key": "LowPriorityCapable",
              "value": "True"
            },
            {
              "key": "PremiumIO",
              "value": "False"
            },
            {
              "key": "VMDeploymentTypes",
              "value": "IaaS"
            },
            {
              "key": "vCPUsConstraintsAllowed",
              "value": "1, 2, 4"
            },
            {
              "key": "vCPUsAvailable",
              "value": "4"
            },
            {
              "key": "vCPUsPerCore",
              "value": "2"
            },
            {
              "key": "CombinedTempDiskAndCachedIOPS",
              "value": "19000"
            },
            {
              "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
              "value": "250000000"
            },
            {
              "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
              "value": "250000000"
            },
            {
              "key": "UncachedDiskIOPS",
              "value": "6400"
            },
            {
              "key": "UncachedDiskBytesPerSecond",
              "value": "145000000"
            },
            {
              "key": "EphemeralOSDiskSupported",
              "value": "False"
            },
            {
              "key": "EncryptionAtHostSupported",
              "value": "True"
            },
            {
              "key": "CapacityReservationSupported",
              "value": "True"
            },
            {
              "key": "AcceleratedNetworkingEnabled",
              "value": "True"
            },
            {
              "key": "RdmaEnabled",
              "value": "False"
            },
            {
              "key": "MaxNetworkInterfaces",
              "value": "2"
            },
            {
              "key": "UltraSSDAvailable",
              "value": "True"
            },
            {
              "key": "LocationInfo_0_Location",
              "value": "KoreaSouth"
            },
            {
              "key": "Family",
              "value": "standardDv5Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "D4_v5"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ],
          "availabilityStatus": "Unverified"
        },
        {
          "id": "azure+koreasouth+standard_d2_v5",
          "uid": "tbt145k3cn5i4f8veqtm",
          "cspSpecName": "Standard_D2_v5",
          "name": "azure+koreasouth+standard_d2_v5",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "costPerHour": 0.11,
          "orderInFilteredResult": 23,
          "evaluationScore01": -1,
          "evaluationScore02": -1,
          "evaluationScore03": -1,
          "evaluationScore04": -1,
          "evaluationScore05": -1,
          "evaluationScore06": -1,
          "evaluationScore07": -1,
          "evaluationScore08": -1,
          "evaluationScore09": -1,
          "evaluationScore10": -1,
          "rootDiskType": "",
          "rootDiskSize": 0,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MaxDataDiskCount",
              "value": "4"
            },
            {
              "key": "MemoryInMB",
              "value": "8192"
            },
            {
              "key": "Name",
              "value": "Standard_D2_v5"
            },
            {
              "key": "NumberOfCores",
              "value": "2"
            },
            {
              "key": "OSDiskSizeInMB",
              "value": "1047552"
            },
            {
              "key": "ResourceDiskSizeInMB",
              "value": "0"
            },
            {
              "key": "MaxResourceVolumeMB",
              "value": "0"
            },
            {
              "key": "OSVhdSizeMB",
              "value": "1047552"
            },
            {
              "key": "vCPUs",
              "value": "2"
            },
            {
              "key": "MemoryPreservingMaintenanceSupported",
              "value": "True"
            },
            {
              "key": "HyperVGenerations",
              "value": "V1,V2"
            },
            {
              "key": "SupportedCapacityReservationTypes",
              "value": "Open,Targeted"
            },
            {
              "key": "MemoryGB",
              "value": "8"
            },
            {
              "key": "MaxDataDiskCount",
              "value": "4"
            },
            {
              "key": "CpuArchitectureType",
              "value": "x64"
            },
            {
              "key": "LowPriorityCapable",
              "value": "True"
            },
            {
              "key": "PremiumIO",
              "value": "False"
            },
            {
              "key": "VMDeploymentTypes",
              "value": "IaaS"
            },
            {
              "key": "vCPUsConstraintsAllowed",
              "value": "1, 2"
            },
            {
              "key": "vCPUsAvailable",
              "value": "2"
            },
            {
              "key": "vCPUsPerCore",
              "value": "2"
            },
            {
              "key": "CombinedTempDiskAndCachedIOPS",
              "value": "9000"
            },
            {
              "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
              "value": "125000000"
            },
            {
              "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
              "value": "125000000"
            },
            {
              "key": "UncachedDiskIOPS",
              "value": "3750"
            },
            {
              "key": "UncachedDiskBytesPerSecond",
              "value": "85000000"
            },
            {
              "key": "EphemeralOSDiskSupported",
              "value": "False"
            },
            {
              "key": "EncryptionAtHostSupported",
              "value": "True"
            },
            {
              "key": "CapacityReservationSupported",
              "value": "True"
            },
            {
              "key": "AcceleratedNetworkingEnabled",
              "value": "True"
            },
            {
              "key": "RdmaEnabled",
              "value": "False"
            },
            {
              "key": "MaxNetworkInterfaces",
              "value": "2"
            },
            {
              "key": "UltraSSDAvailable",
              "value": "True"
            },
            {
              "key": "LocationInfo_0_Location",
              "value": "KoreaSouth"
            },
            {
              "key": "Family",
              "value": "standardDv5Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "D2_v5"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ],
          "availabilityStatus": "Unverified"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "uid": "tbmm9857u5pvqm24elb2",
          "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:34:25 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "osDiskType": "default",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Location",
              "value": "australiacentral"
            },
            {
              "key": "Publisher",
              "value": "Canonical"
            },
            {
              "key": "Offer",
              "value": "0001-com-ubuntu-server-jammy"
            },
            {
              "key": "SKU",
              "value": "22_04-lts-gen2"
            },
            {
              "key": "Version",
              "value": "22.04.202610220"
            },
            {
              "key": "ID",
              "value": "/Subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202610220"
            },
            {
              "key": "HyperVGeneration",
              "value": "V2"
            },
            {
              "key": "Features",
              "value": "SecurityType=TrustedLaunchSupported, IsAcceleratedNetworkSupported=True, DiskControllerTypes=SCSI, NVMe, IsHibernateSupported=True"
            },
            {
              "key": "FeatureCount",
              "value": "4"
            },
            {
              "key": "ImageDeprecationState",
              "value": "Active"
            }
          ],
          "systemLabel": "from-assets",
          "description": "Server LTS Gen2",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "azure-koreasouth",
          "vNetId": "mig-vnet-01",
          "description": "Recommended security group for NLB backend influxdb_back",
          "firewallRules": [
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "8086",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "",
              "Protocol": "ALL",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9999",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "8086",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            }
          ],
          "cspResourceId": ""
        },
        {
          "name": "mig-sg-02",
          "connectionName": "azure-koreasouth",
          "vNetId": "mig-vnet-01",
          "description": "Recommended security group for ec268ed7-821e-9d73-e79f-961262161624",
          "firewallRules": [
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "9999",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "",
              "Protocol": "ALL",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            }
          ],
          "cspResourceId": ""
        }
      ],
      "targetNlbList": [
        {
          "description": "Migrated from HAProxy backend: influxdb_back",
          "type": "PUBLIC",
          "scope": "REGION",
          "listener": {
            "protocol": "TCP",
            "port": "9999"
          },
          "targetGroup": {
            "protocol": "TCP",
            "port": "8086",
            "nodeGroupId": "ng-influxdb-back"
          },
          "healthChecker": {
            "interval": 10,
            "threshold": 3,
            "timeout": -1
          }
        }
      ],
      "targetK8sCluster": {
        "connectionName": "",
        "description": "",
        "name": "",
        "version": "",
        "vNetId": "",
        "subnetIds": null,
        "securityGroupIds": null,
        "k8sNodeGroupList": null,
        "cspResourceId": "",
        "label": null,
        "systemLabel": ""
      }
    }
  ],
  "message": "2 candidate(s) recommended — each with 1 NLB(s) and 2 NodeGroup(s)"
}
```

</details>

### Test Case 2: Validate the target model for computing infra

#### 2.1 API Request Information

- **API Endpoint**: `POST /beetle/validation/ns/mig01/infra`
- **Purpose**: Validate the recommended target model before migration (name collisions, spec/image compatibility, resource availability)

**Request Body**:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "status": "partially-matched",
  "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=86.1% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
  "targetCloud": {
    "csp": "azure",
    "region": "koreasouth"
  },
  "targetInfra": {
    "name": "my02-infra101",
    "label": null,
    "systemLabel": "",
    "description": "NLB-aware recommended infrastructure for cloud migration",
    "nodeGroups": [
      {
        "name": "my02-ng-influxdb-back",
        "nodeGroupSize": 2,
        "label": {
          "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "nlbBackend": "influxdb_back"
        },
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
        "connectionName": "azure-koreasouth",
        "specId": "azure+koreasouth+standard_d4_v4",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
        "vNetId": "my02-mig-vnet-01",
        "subnetId": "my02-mig-subnet-01",
        "securityGroupIds": [
          "my02-mig-sg-01"
        ],
        "sshKeyId": "my02-mig-sshkey-01",
        "rootDiskSize": 30,
        "dataDiskIds": null
      },
      {
        "name": "my02-ng-ip-10-0-1-30-62161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=83.3%",
        "connectionName": "azure-koreasouth",
        "specId": "azure+koreasouth+standard_f2s_v2",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
        "vNetId": "my02-mig-vnet-01",
        "subnetId": "my02-mig-subnet-01",
        "securityGroupIds": [
          "my02-mig-sg-02"
        ],
        "sshKeyId": "my02-mig-sshkey-01",
        "rootDiskSize": 30,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "my02-mig-vnet-01",
    "connectionName": "azure-koreasouth",
    "cidrBlock": "10.0.0.0/21",
    "subnetInfoList": [
      {
        "name": "my02-mig-subnet-01",
        "ipv4_CIDR": "10.0.1.0/24",
        "description": "a recommended subnet for migration"
      }
    ],
    "description": "a recommended vNet for migration"
  },
  "targetSshKey": {
    "name": "my02-mig-sshkey-01",
    "connectionName": "azure-koreasouth",
    "description": "SSH key pair for migration (Note: provided ONLY once, MUST be downloaded)",
    "cspResourceId": "",
    "fingerprint": "",
    "username": "",
    "verifiedUsername": "",
    "publicKey": "",
    "privateKey": ""
  },
  "targetSpecList": [
    {
      "id": "azure+koreasouth+standard_d4_v4",
      "uid": "tb2f522bo2kgbr6t0n9g",
      "cspSpecName": "Standard_D4_v4",
      "name": "azure+koreasouth+standard_d4_v4",
      "namespace": "system",
      "connectionName": "azure-koreasouth",
      "providerName": "azure",
      "regionName": "koreasouth",
      "regionLatitude": 35.1796,
      "regionLongitude": 129.0756,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 16,
      "costPerHour": 0.221,
      "orderInFilteredResult": 30,
      "evaluationScore01": -1,
      "evaluationScore02": -1,
      "evaluationScore03": -1,
      "evaluationScore04": -1,
      "evaluationScore05": -1,
      "evaluationScore06": -1,
      "evaluationScore07": -1,
      "evaluationScore08": -1,
      "evaluationScore09": -1,
      "evaluationScore10": -1,
      "rootDiskType": "",
      "rootDiskSize": 0,
      "systemLabel": "auto-gen",
      "details": [
        {
          "key": "MaxDataDiskCount",
          "value": "8"
        },
        {
          "key": "MemoryInMB",
          "value": "16384"
        },
        {
          "key": "Name",
          "value": "Standard_D4_v4"
        },
        {
          "key": "NumberOfCores",
          "value": "4"
        },
        {
          "key": "OSDiskSizeInMB",
          "value": "1047552"
        },
        {
          "key": "ResourceDiskSizeInMB",
          "value": "0"
        },
        {
          "key": "MaxResourceVolumeMB",
          "value": "0"
        },
        {
          "key": "OSVhdSizeMB",
          "value": "1047552"
        },
        {
          "key": "vCPUs",
          "value": "4"
        },
        {
          "key": "MemoryPreservingMaintenanceSupported",
          "value": "True"
        },
        {
          "key": "HyperVGenerations",
          "value": "V1,V2"
        },
        {
          "key": "SupportedCapacityReservationTypes",
          "value": "Open,Targeted"
        },
        {
          "key": "MemoryGB",
          "value": "16"
        },
        {
          "key": "MaxDataDiskCount",
          "value": "8"
        },
        {
          "key": "CpuArchitectureType",
          "value": "x64"
        },
        {
          "key": "LowPriorityCapable",
          "value": "True"
        },
        {
          "key": "PremiumIO",
          "value": "False"
        },
        {
          "key": "VMDeploymentTypes",
          "value": "IaaS"
        },
        {
          "key": "vCPUsConstraintsAllowed",
          "value": "1, 2, 4"
        },
        {
          "key": "vCPUsAvailable",
          "value": "4"
        },
        {
          "key": "ACUs",
          "value": "195"
        },
        {
          "key": "vCPUsPerCore",
          "value": "2"
        },
        {
          "key": "CombinedTempDiskAndCachedIOPS",
          "value": "38500"
        },
        {
          "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
          "value": "242221056"
        },
        {
          "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
          "value": "242221056"
        },
        {
          "key": "UncachedDiskIOPS",
          "value": "6400"
        },
        {
          "key": "UncachedDiskBytesPerSecond",
          "value": "96000000"
        },
        {
          "key": "EphemeralOSDiskSupported",
          "value": "False"
        },
        {
          "key": "EncryptionAtHostSupported",
          "value": "False"
        },
        {
          "key": "CapacityReservationSupported",
          "value": "True"
        },
        {
          "key": "AcceleratedNetworkingEnabled",
          "value": "True"
        },
        {
          "key": "RdmaEnabled",
          "value": "False"
        },
        {
          "key": "MaxNetworkInterfaces",
          "value": "2"
        },
        {
          "key": "LocationInfo_0_Location",
          "value": "KoreaSouth"
        },
        {
          "key": "Family",
          "value": "standardDv4Family"
        },
        {
          "key": "Tier",
          "value": "Standard"
        },
        {
          "key": "Size",
          "value": "D4_v4"
        },
        {
          "key": "ResourceType",
          "value": "virtualMachines"
        }
      ],
      "availabilityStatus": "Unverified"
    },
    {
      "id": "azure+koreasouth+standard_f2s_v2",
      "uid": "tbr3k3hgqboa0n905jlg",
      "cspSpecName": "Standard_F2s_v2",
      "name": "azure+koreasouth+standard_f2s_v2",
      "namespace": "system",
      "connectionName": "azure-koreasouth",
      "providerName": "azure",
      "regionName": "koreasouth",
      "regionLatitude": 35.1796,
      "regionLongitude": 129.0756,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 4,
      "diskSizeGB": 17,
      "costPerHour": 0.0961,
      "orderInFilteredResult": 14,
      "evaluationScore01": -1,
      "evaluationScore02": -1,
      "evaluationScore03": -1,
      "evaluationScore04": -1,
      "evaluationScore05": -1,
      "evaluationScore06": -1,
      "evaluationScore07": -1,
      "evaluationScore08": -1,
      "evaluationScore09": -1,
      "evaluationScore10": -1,
      "rootDiskType": "",
      "rootDiskSize": 17,
      "systemLabel": "auto-gen",
      "details": [
        {
          "key": "MaxDataDiskCount",
          "value": "4"
        },
        {
          "key": "MemoryInMB",
          "value": "4096"
        },
        {
          "key": "Name",
          "value": "Standard_F2s_v2"
        },
        {
          "key": "NumberOfCores",
          "value": "2"
        },
        {
          "key": "OSDiskSizeInMB",
          "value": "1047552"
        },
        {
          "key": "ResourceDiskSizeInMB",
          "value": "16384"
        },
        {
          "key": "MaxResourceVolumeMB",
          "value": "16384"
        },
        {
          "key": "OSVhdSizeMB",
          "value": "1047552"
        },
        {
          "key": "vCPUs",
          "value": "2"
        },
        {
          "key": "MemoryPreservingMaintenanceSupported",
          "value": "True"
        },
        {
          "key": "HyperVGenerations",
          "value": "V1,V2"
        },
        {
          "key": "SupportedCapacityReservationTypes",
          "value": "Open,Targeted"
        },
        {
          "key": "SupportedEphemeralOSDiskPlacements",
          "value": "ResourceDisk,CacheDisk"
        },
        {
          "key": "MemoryGB",
          "value": "4"
        },
        {
          "key": "MaxDataDiskCount",
          "value": "4"
        },
        {
          "key": "CpuArchitectureType",
          "value": "x64"
        },
        {
          "key": "LowPriorityCapable",
          "value": "True"
        },
        {
          "key": "PremiumIO",
          "value": "True"
        },
        {
          "key": "VMDeploymentTypes",
          "value": "IaaS"
        },
        {
          "key": "vCPUsConstraintsAllowed",
          "value": "1, 2"
        },
        {
          "key": "vCPUsAvailable",
          "value": "2"
        },
        {
          "key": "ACUs",
          "value": "195"
        },
        {
          "key": "vCPUsPerCore",
          "value": "2"
        },
        {
          "key": "CombinedTempDiskAndCachedIOPS",
          "value": "4000"
        },
        {
          "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
          "value": "32768000"
        },
        {
          "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
          "value": "32768000"
        },
        {
          "key": "CachedDiskBytes",
          "value": "34359738368"
        },
        {
          "key": "UncachedDiskIOPS",
          "value": "3200"
        },
        {
          "key": "UncachedDiskBytesPerSecond",
          "value": "48000000"
        },
        {
          "key": "RetirementDateUtc",
          "value": "11/15/2028"
        },
        {
          "key": "EphemeralOSDiskSupported",
          "value": "True"
        },
        {
          "key": "EncryptionAtHostSupported",
          "value": "True"
        },
        {
          "key": "CapacityReservationSupported",
          "value": "True"
        },
        {
          "key": "AcceleratedNetworkingEnabled",
          "value": "True"
        },
        {
          "key": "RdmaEnabled",
          "value": "False"
        },
        {
          "key": "MaxNetworkInterfaces",
          "value": "2"
        },
        {
          "key": "UltraSSDAvailable",
          "value": "True"
        },
        {
          "key": "LocationInfo_0_Location",
          "value": "KoreaSouth"
        },
        {
          "key": "Family",
          "value": "standardFSv2Family"
        },
        {
          "key": "Tier",
          "value": "Standard"
        },
        {
          "key": "Size",
          "value": "F2s_v2"
        },
        {
          "key": "ResourceType",
          "value": "virtualMachines"
        }
      ],
      "availabilityStatus": "Unverified"
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "azure",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "regionList": [
        "common"
      ],
      "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "uid": "tbmm9857u5pvqm24elb2",
      "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "azure-australiacentral",
      "infraType": "",
      "fetchedTime": "2026.10.29 04:34:25 Thu",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "osDiskType": "default",
      "osDiskSizeGB": -1,
      "imageStatus": "Available",
      "details": [
        {
          "key": "Location",
          "value": "australiacentral"
        },
        {
          "key": "Publisher",
          "value": "Canonical"
        },
        {
          "key": "Offer",
          "value": "0001-com-ubuntu-server-jammy"
        },
        {
          "key": "SKU",
          "value": "22_04-lts-gen2"
        },
        {
          "key": "Version",
          "value": "22.04.202610220"
        },
        {
          "key": "ID",
          "value": "/Subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202610220"
        },
        {
          "key": "HyperVGeneration",
          "value": "V2"
        },
        {
          "key": "Features",
          "value": "SecurityType=TrustedLaunchSupported, IsAcceleratedNetworkSupported=True, DiskControllerTypes=SCSI, NVMe, IsHibernateSupported=True"
        },
        {
          "key": "FeatureCount",
          "value": "4"
        },
        {
          "key": "ImageDeprecationState",
          "value": "Active"
        }
      ],
      "systemLabel": "from-assets",
      "description": "Server LTS Gen2",
      "commandHistory": null
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "my02-mig-sg-01",
      "connectionName": "azure-koreasouth",
      "vNetId": "my02-mig-vnet-01",
      "description": "Recommended security group for NLB backend influxdb_back",
      "firewallRules": [
        {
          "Ports": "22",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "8086",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "",
          "Protocol": "ALL",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "9999",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "8086",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        }
      ],
      "cspResourceId": ""
    },
    {
      "name": "my02-mig-sg-02",
      "connectionName": "azure-koreasouth",
      "vNetId": "my02-mig-vnet-01",
      "description": "Recommended security group for ec268ed7-821e-9d73-e79f-961262161624",
      "firewallRules": [
        {
          "Ports": "22",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "9999",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "",
          "Protocol": "ALL",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        }
      ],
      "cspResourceId": ""
    }
  ],
  "targetNlbList": [
    {
      "description": "Migrated from HAProxy backend: influxdb_back",
      "type": "PUBLIC",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "port": "9999"
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my02-ng-influxdb-back"
      },
      "healthChecker": {
        "interval": 10,
        "threshold": 3,
        "timeout": -1
      }
    }
  ],
  "targetK8sCluster": {
    "connectionName": "",
    "description": "",
    "name": "",
    "version": "",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": null,
    "cspResourceId": "",
    "label": null,
    "systemLabel": ""
  }
}
```

</details>

#### 2.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Target model is valid, no issues found

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "valid": true,
  "issues": []
}
```

</details>

### Test Case 3: Migrate the computing infra as defined in the target model

#### 3.1 API Request Information

- **API Endpoint**: `POST /beetle/migration/ns/mig01/infra`
- **Purpose**: Create and migrate infrastructure based on recommendation
- **Namespace ID**: `mig01`
- **Request Body**: Uses the response from the previous recommendation step

#### 3.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **MCI ID**: `my02-infra101`
- **MCI Name**: `my02-infra101`
- **Status**: `Running:3 (R:3/3)`

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my02-infra101",
  "uid": "tb3t1dqr76ce6f3l49ug",
  "name": "my02-infra101",
  "status": "Running:3 (R:3/3)",
  "statusCount": {
    "countTotal": 3,
    "countCreating": 0,
    "countRunning": 3,
    "countFailed": 0,
    "countSuspended": 0,
    "countRebooting": 0,
    "countTerminated": 0,
    "countSuspending": 0,
    "countResuming": 0,
    "countTerminating": 0,
    "countRegistering": 0,
    "countReconciling": 0,
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "cm-migrated-at": "2026-11-25T09:15:11Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my02-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my02-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tb3t1dqr76ce6f3l49ug"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my02-ng-influxdb-back-1",
      "uid": "tb2mefhc7qac6sg6gmoh",
      "cspResourceName": "tb2mefhc7qac6sg6gmoh",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb2mefhc7qac6sg6gmoh",
      "name": "my02-ng-influxdb-back-1",
      "nodeGroupId": "my02-ng-influxdb-back",
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:03",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "createdBy": "tb2mefhc7qac6sg6gmoh",
        "keypair": "tbbm9d00plcuabc5g4jg",
        "nlbBackend": "influxdb_back",
        "publicip": "tb2mefhc7qac6sg6gmoh-91708-PublicIP"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "52.147.125.128",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "StandardHDD",
      "rootDiskSize": 30,
      "rootDeviceName": "Not visible in Azure",
      "connectionName": "azure-koreasouth",
      "connectionConfig": {
        "configName": "azure-koreasouth",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreasouth",
        "regionZoneInfo": {
          "assignedRegion": "koreasouth",
          "assignedZone": ""
        },
        "regionDetail": {
          "regionId": "koreasouth",
          "regionName": "koreasouth",
          "description": "Korea South",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "zones": []
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreasouth+standard_d4_v4",
      "cspSpecName": "Standard_D4_v4",
      "spec": {
        "cspSpecName": "Standard_D4_v4",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.221
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm",
      "networkInterface": "tb2mefhc7qac6sg6gmoh-61943-VNic",
      "securityGroupIds": [
        "my02-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJ66k1oYuwBtYA5NOE/x6ieyZkjOF/fzk5ySLbd9GfcC/PJnflIiuyOBELFTzQtzN8k+Z6Sqwyad2RitrdzH+Xw=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:9AeffaLS0Nj2SNqwUICjEFi7i04E4T3ap6yI3NXJYow",
        "firstUsedAt": "2026-11-25T09:16:07Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:03Z",
          "completedTime": "2026-11-25T09:16:09Z",
          "elapsedTime": 6,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my02-infra101-tb7q121q173qpuoqlvva",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:11Z",
          "completedTime": "2026-11-25T09:16:12Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb2mefhc7qac6sg6gmoh 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my02-ng-influxdb-back-2",
      "uid": "tb6lhvd6gifs72gf64ok",
      "cspResourceName": "tb6lhvd6gifs72gf64ok",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb6lhvd6gifs72gf64ok",
      "name": "my02-ng-influxdb-back-2",
      "nodeGroupId": "my02-ng-influxdb-back",
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:03",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "createdBy": "tb6lhvd6gifs72gf64ok",
        "keypair": "tbbm9d00plcuabc5g4jg",
        "nlbBackend": "influxdb_back",
        "publicip": "tb6lhvd6gifs72gf64ok-39512-PublicIP"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "74.227.129.35",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "StandardHDD",
      "rootDiskSize": 30,
      "rootDeviceName": "Not visible in Azure",
      "connectionName": "azure-koreasouth",
      "connectionConfig": {
        "configName": "azure-koreasouth",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreasouth",
        "regionZoneInfo": {
          "assignedRegion": "koreasouth",
          "assignedZone": ""
        },
        "regionDetail": {
          "regionId": "koreasouth",
          "regionName": "koreasouth",
          "description": "Korea South",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "zones": []
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreasouth+standard_d4_v4",
      "cspSpecName": "Standard_D4_v4",
      "spec": {
        "cspSpecName": "Standard_D4_v4",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.221
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm",
      "networkInterface": "tb6lhvd6gifs72gf64ok-32847-VNic",
      "securityGroupIds": [
        "my02-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMRFzcn/w5rroLqFcUp0u5fnG2nZSHGUDEzCBNRmmFtQySyMPoQjhIXdh4WP9ordWlEmwgX/8rwHyOuWuGlbBlI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:j/BfQxwdAH7y6IarUawWD7py8RYnF3uvhw4bxtAz9Rw",
        "firstUsedAt": "2026-11-25T09:16:07Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:03Z",
          "completedTime": "2026-11-25T09:16:10Z",
          "elapsedTime": 7,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my02-infra101-tb7q121q173qpuoqlvva",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:11Z",
          "completedTime": "2026-11-25T09:16:13Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb6lhvd6gifs72gf64ok 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my02-ng-ip-10-0-1-30-62161624-1",
      "uid": "tbd32ihs5kcb30u4ok9r",
      "cspResourceName": "tbd32ihs5kcb30u4ok9r",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd32ihs5kcb30u4ok9r",
      "name": "my02-ng-ip-10-0-1-30-62161624-1",
      "nodeGroupId": "my02-ng-ip-10-0-1-30-62161624",
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:03",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
        "createdBy": "tbd32ihs5kcb30u4ok9r",
        "keypair": "tbbm9d00plcuabc5g4jg",
        "publicip": "tbd32ihs5kcb30u4ok9r-53594-PublicIP"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "74.227.128.223",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 30,
      "rootDeviceName": "Not visible in Azure",
      "connectionName": "azure-koreasouth",
      "connectionConfig": {
        "configName": "azure-koreasouth",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreasouth",
        "regionZoneInfo": {
          "assignedRegion": "koreasouth",
          "assignedZone": ""
        },
        "regionDetail": {
          "regionId": "koreasouth",
          "regionName": "koreasouth",
          "description": "Korea South",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "zones": []
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreasouth+standard_f2s_v2",
      "cspSpecName": "Standard_F2s_v2",
      "spec": {
        "cspSpecName": "Standard_F2s_v2",
        "vCPU": 2,
        "memoryGiB": 4,
        "costPerHour": 0.0961
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm",
      "networkInterface": "tbd32ihs5kcb30u4ok9r-15740-VNic",
      "securityGroupIds": [
        "my02-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHI+eba+Z/Dhzv6Jm8MotVrH8Kx1nA1pM1RNUFZqkym1Mz372/ksu/DU8NcqcUPNw3/JdHZLSC8GTlwjQHeXZCI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:RVYMe+sr8Bqh2p9XkxhkP47Be/vb/LbfCGmnjWUblf4",
        "firstUsedAt": "2026-11-25T09:16:04Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:03Z",
          "completedTime": "2026-11-25T09:16:11Z",
          "elapsedTime": 8,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my02-infra101-tb7q121q173qpuoqlvva",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:11Z",
          "completedTime": "2026-11-25T09:16:14Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd32ihs5kcb30u4ok9r 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my02-mig-vnet-01",
      "name": "my02-mig-vnet-01",
      "infraId": "my02-infra101",
      "vNetId": "my02-mig-vnet-01",
      "connectionNames": [
        "azure-koreasouth"
      ],
      "providerNames": [
        "azure"
      ],
      "regionNames": [
        "koreasouth"
      ],
      "nodeGroupIds": [
        "my02-ng-influxdb-back",
        "my02-ng-ip-10-0-1-30-62161624"
      ],
      "nodeIds": [
        "my02-ng-influxdb-back-1",
        "my02-ng-influxdb-back-2",
        "my02-ng-ip-10-0-1-30-62161624-1"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my02-ng-influxdb-back",
      "representativeNodeId": "my02-ng-influxdb-back-1"
    }
  ],
  "newNodeList": null,
  "postScript": {
    "phases": [
      {
        "userName": "cb-user",
        "command": [
          "uname -a"
        ]
      }
    ]
  },
  "postScriptResults": [
    {
      "phase": 1,
      "target": "all nodes",
      "status": "Completed",
      "results": {
        "results": [
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-influxdb-back-1",
            "nodeIp": "52.147.125.128",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb2mefhc7qac6sg6gmoh 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-influxdb-back-2",
            "nodeIp": "74.227.129.35",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb6lhvd6gifs72gf64ok 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-ip-10-0-1-30-62161624-1",
            "nodeIp": "74.227.128.223",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbd32ihs5kcb30u4ok9r 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          }
        ]
      }
    }
  ],
  "postScriptStatus": "Completed",
  "postScriptRequestId": "ps-my02-infra101-tb7q121q173qpuoqlvva"
}
```

</details>

### Test Case 4: Get a list of infras

#### 4.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra`
- **Purpose**: Get a list of all migrated infrastructures
- **Namespace ID**: `mig01`

#### 4.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Count**: 3

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "infra": [
    {
      "resourceType": "infra",
      "id": "my01-infra101",
      "uid": "tbohd5haepfsekdac2ko",
      "name": "my01-infra101",
      "status": "Running:3 (R:3/3)",
      "statusCount": {
        "countTotal": 3,
        "countCreating": 0,
        "countRunning": 3,
        "countFailed": 0,
        "countSuspended": 0,
        "countRebooting": 0,
        "countTerminated": 0,
        "countSuspending": 0,
        "countResuming": 0,
        "countTerminating": 0,
        "countRegistering": 0,
        "countReconciling": 0,
        "countUndefined": 0
      },
      "targetStatus": "None",
      "targetAction": "None",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "cm-migrated-at": "2026-11-25T09:14:38Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my01-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my01-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbohd5haepfsekdac2ko"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my01-ng-influxdb-back-1",
          "uid": "tb9kb21aeetl91i0qfdj",
          "cspResourceName": "tb9kb21aeetl91i0qfdj",
          "cspResourceId": "i-0bbb8c166a68d3f41",
          "name": "my01-ng-influxdb-back-1",
          "nodeGroupId": "my01-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:14:55",
          "label": {
            "Name": "tb9kb21aeetl91i0qfdj",
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-11-25 09:14:55",
            "sys.cspResourceId": "i-0bbb8c166a68d3f41",
            "sys.cspResourceName": "tb9kb21aeetl91i0qfdj",
            "sys.id": "my01-ng-influxdb-back-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-influxdb-back",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tb9kb21aeetl91i0qfdj",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "3.38.252.17",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.123",
          "privateDNS": "ip-10-0-1-123.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "aws-ap-northeast-2",
          "connectionConfig": {
            "configName": "aws-ap-northeast-2",
            "providerName": "aws",
            "driverName": "aws-driver-v1.0.so",
            "credentialName": "aws",
            "credentialHolder": "admin",
            "regionZoneInfoName": "aws-ap-northeast-2",
            "regionZoneInfo": {
              "assignedRegion": "ap-northeast-2",
              "assignedZone": "ap-northeast-2a"
            },
            "regionDetail": {
              "regionId": "ap-northeast-2",
              "regionName": "ap-northeast-2",
              "description": "Asia Pacific (Seoul)",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.36,
                "longitude": 126.78
              },
              "zones": [
                "ap-northeast-2a",
                "ap-northeast-2b",
                "ap-northeast-2c",
                "ap-northeast-2d"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "aws+ap-northeast-2+t3.xlarge",
          "cspSpecName": "t3.xlarge",
          "spec": {
            "cspSpecName": "t3.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.208
          },
          "imageId": "ami-046333f9ecafdf406",
          "cspImageName": "ami-046333f9ecafdf406",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-046333f9ecafdf406",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021"
          },
          "vNetId": "my01-mig-vnet-01",
          "cspVNetId": "vpc-09aa8042d1c874768",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-063b3e54670c7a6f5",
          "networkInterface": "eni-0f4c6e1a23c3cae52",
          "securityGroupIds": [
            "my01-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tb0tqsk2ejcg4plsgf7r"
        },
        {
          "resourceType": "node",
          "id": "my01-ng-influxdb-back-2",
          "uid": "tbn8mcdp8a5j8e8918um",
          "cspResourceName": "tbn8mcdp8a5j8e8918um",
          "cspResourceId": "i-0f6efd5a8c05d59e1",
          "name": "my01-ng-influxdb-back-2",
          "nodeGroupId": "my01-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:14:59",
          "label": {
            "Name": "tbn8mcdp8a5j8e8918um",
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-11-25 09:14:59",
            "sys.cspResourceId": "i-0f6efd5a8c05d59e1",
            "sys.cspResourceName": "tbn8mcdp8a5j8e8918um",
            "sys.id": "my01-ng-influxdb-back-2",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-influxdb-back",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tbn8mcdp8a5j8e8918um",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "52.79.141.27",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.193",
          "privateDNS": "ip-10-0-1-193.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "aws-ap-northeast-2",
          "connectionConfig": {
            "configName": "aws-ap-northeast-2",
            "providerName": "aws",
            "driverName": "aws-driver-v1.0.so",
            "credentialName": "aws",
            "credentialHolder": "admin",
            "regionZoneInfoName": "aws-ap-northeast-2",
            "regionZoneInfo": {
              "assignedRegion": "ap-northeast-2",
              "assignedZone": "ap-northeast-2a"
            },
            "regionDetail": {
              "regionId": "ap-northeast-2",
              "regionName": "ap-northeast-2",
              "description": "Asia Pacific (Seoul)",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.36,
                "longitude": 126.78
              },
              "zones": [
                "ap-northeast-2a",
                "ap-northeast-2b",
                "ap-northeast-2c",
                "ap-northeast-2d"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "aws+ap-northeast-2+t3.xlarge",
          "cspSpecName": "t3.xlarge",
          "spec": {
            "cspSpecName": "t3.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.208
          },
          "imageId": "ami-046333f9ecafdf406",
          "cspImageName": "ami-046333f9ecafdf406",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-046333f9ecafdf406",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021"
          },
          "vNetId": "my01-mig-vnet-01",
          "cspVNetId": "vpc-09aa8042d1c874768",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-063b3e54670c7a6f5",
          "networkInterface": "eni-0cfff99b6b9719dd8",
          "securityGroupIds": [
            "my01-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tb0tqsk2ejcg4plsgf7r"
        },
        {
          "resourceType": "node",
          "id": "my01-ng-ip-10-0-1-30-62161624-1",
          "uid": "tbqfgi5tbnlnpkk7ec6e",
          "cspResourceName": "tbqfgi5tbnlnpkk7ec6e",
          "cspResourceId": "i-07196fca4d742d8ca",
          "name": "my01-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my01-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:14:56",
          "label": {
            "Name": "tbqfgi5tbnlnpkk7ec6e",
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-11-25 09:14:56",
            "sys.cspResourceId": "i-07196fca4d742d8ca",
            "sys.cspResourceName": "tbqfgi5tbnlnpkk7ec6e",
            "sys.id": "my01-ng-ip-10-0-1-30-62161624-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-ip-10-0-1-30-62161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-ip-10-0-1-30-62161624",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tbqfgi5tbnlnpkk7ec6e",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "15.165.15.17",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.194",
          "privateDNS": "ip-10-0-1-194.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 10,
          "rootDeviceName": "",
          "connectionName": "aws-ap-northeast-2",
          "connectionConfig": {
            "configName": "aws-ap-northeast-2",
            "providerName": "aws",
            "driverName": "aws-driver-v1.0.so",
            "credentialName": "aws",
            "credentialHolder": "admin",
            "regionZoneInfoName": "aws-ap-northeast-2",
            "regionZoneInfo": {
              "assignedRegion": "ap-northeast-2",
              "assignedZone": "ap-northeast-2a"
            },
            "regionDetail": {
              "regionId": "ap-northeast-2",
              "regionName": "ap-northeast-2",
              "description": "Asia Pacific (Seoul)",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.36,
                "longitude": 126.78
              },
              "zones": [
                "ap-northeast-2a",
                "ap-northeast-2b",
                "ap-northeast-2c",
                "ap-northeast-2d"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "aws+ap-northeast-2+t3.small",
          "cspSpecName": "t3.small",
          "spec": {
            "cspSpecName": "t3.small",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.026
          },
          "imageId": "ami-046333f9ecafdf406",
          "cspImageName": "ami-046333f9ecafdf406",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-046333f9ecafdf406",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021"
          },
          "vNetId": "my01-mig-vnet-01",
          "cspVNetId": "vpc-09aa8042d1c874768",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-063b3e54670c7a6f5",
          "networkInterface": "eni-02dd0387f7ddfee69",
          "securityGroupIds": [
            "my01-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tb0tqsk2ejcg4plsgf7r"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my02-infra101",
      "uid": "tb3t1dqr76ce6f3l49ug",
      "name": "my02-infra101",
      "status": "Running:3 (R:3/3)",
      "statusCount": {
        "countTotal": 3,
        "countCreating": 0,
        "countRunning": 3,
        "countFailed": 0,
        "countSuspended": 0,
        "countRebooting": 0,
        "countTerminated": 0,
        "countSuspending": 0,
        "countResuming": 0,
        "countTerminating": 0,
        "countRegistering": 0,
        "countReconciling": 0,
        "countUndefined": 0
      },
      "targetStatus": "None",
      "targetAction": "None",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "cm-migrated-at": "2026-11-25T09:15:11Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my02-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tb3t1dqr76ce6f3l49ug"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my02-ng-influxdb-back-1",
          "uid": "tb2mefhc7qac6sg6gmoh",
          "cspResourceName": "tb2mefhc7qac6sg6gmoh",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb2mefhc7qac6sg6gmoh",
          "name": "my02-ng-influxdb-back-1",
          "nodeGroupId": "my02-ng-influxdb-back",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:03",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "createdBy": "tb2mefhc7qac6sg6gmoh",
            "keypair": "tbbm9d00plcuabc5g4jg",
            "nlbBackend": "influxdb_back",
            "publicip": "tb2mefhc7qac6sg6gmoh-91708-PublicIP",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-11-25 09:16:03",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb2mefhc7qac6sg6gmoh",
            "sys.cspResourceName": "tb2mefhc7qac6sg6gmoh",
            "sys.id": "my02-ng-influxdb-back-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-influxdb-back",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tb2mefhc7qac6sg6gmoh",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "52.147.125.128",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "StandardHDD",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "azure-koreasouth",
          "connectionConfig": {
            "configName": "azure-koreasouth",
            "providerName": "azure",
            "driverName": "azure-driver-v1.0.so",
            "credentialName": "azure",
            "credentialHolder": "admin",
            "regionZoneInfoName": "azure-koreasouth",
            "regionZoneInfo": {
              "assignedRegion": "koreasouth",
              "assignedZone": ""
            },
            "regionDetail": {
              "regionId": "koreasouth",
              "regionName": "koreasouth",
              "description": "Korea South",
              "location": {
                "display": "Korea South",
                "latitude": 35.1796,
                "longitude": 129.0756
              },
              "zones": []
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "azure+koreasouth+standard_d4_v4",
          "cspSpecName": "Standard_D4_v4",
          "spec": {
            "cspSpecName": "Standard_D4_v4",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.221
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm",
          "networkInterface": "tb2mefhc7qac6sg6gmoh-61943-VNic",
          "securityGroupIds": [
            "my02-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg"
        },
        {
          "resourceType": "node",
          "id": "my02-ng-influxdb-back-2",
          "uid": "tb6lhvd6gifs72gf64ok",
          "cspResourceName": "tb6lhvd6gifs72gf64ok",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb6lhvd6gifs72gf64ok",
          "name": "my02-ng-influxdb-back-2",
          "nodeGroupId": "my02-ng-influxdb-back",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:03",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "createdBy": "tb6lhvd6gifs72gf64ok",
            "keypair": "tbbm9d00plcuabc5g4jg",
            "nlbBackend": "influxdb_back",
            "publicip": "tb6lhvd6gifs72gf64ok-39512-PublicIP",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-11-25 09:16:03",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb6lhvd6gifs72gf64ok",
            "sys.cspResourceName": "tb6lhvd6gifs72gf64ok",
            "sys.id": "my02-ng-influxdb-back-2",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-influxdb-back",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tb6lhvd6gifs72gf64ok",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "74.227.129.35",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "StandardHDD",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "azure-koreasouth",
          "connectionConfig": {
            "configName": "azure-koreasouth",
            "providerName": "azure",
            "driverName": "azure-driver-v1.0.so",
            "credentialName": "azure",
            "credentialHolder": "admin",
            "regionZoneInfoName": "azure-koreasouth",
            "regionZoneInfo": {
              "assignedRegion": "koreasouth",
              "assignedZone": ""
            },
            "regionDetail": {
              "regionId": "koreasouth",
              "regionName": "koreasouth",
              "description": "Korea South",
              "location": {
                "display": "Korea South",
                "latitude": 35.1796,
                "longitude": 129.0756
              },
              "zones": []
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "azure+koreasouth+standard_d4_v4",
          "cspSpecName": "Standard_D4_v4",
          "spec": {
            "cspSpecName": "Standard_D4_v4",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.221
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm",
          "networkInterface": "tb6lhvd6gifs72gf64ok-32847-VNic",
          "securityGroupIds": [
            "my02-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg"
        },
        {
          "resourceType": "node",
          "id": "my02-ng-ip-10-0-1-30-62161624-1",
          "uid": "tbd32ihs5kcb30u4ok9r",
          "cspResourceName": "tbd32ihs5kcb30u4ok9r",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd32ihs5kcb30u4ok9r",
          "name": "my02-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my02-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:03",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "createdBy": "tbd32ihs5kcb30u4ok9r",
            "keypair": "tbbm9d00plcuabc5g4jg",
            "publicip": "tbd32ihs5kcb30u4ok9r-53594-PublicIP",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-11-25 09:16:03",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd32ihs5kcb30u4ok9r",
            "sys.cspResourceName": "tbd32ihs5kcb30u4ok9r",
            "sys.id": "my02-ng-ip-10-0-1-30-62161624-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-ip-10-0-1-30-62161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-ip-10-0-1-30-62161624",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tbd32ihs5kcb30u4ok9r",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "74.227.128.223",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.5",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "azure-koreasouth",
          "connectionConfig": {
            "configName": "azure-koreasouth",
            "providerName": "azure",
            "driverName": "azure-driver-v1.0.so",
            "credentialName": "azure",
            "credentialHolder": "admin",
            "regionZoneInfoName": "azure-koreasouth",
            "regionZoneInfo": {
              "assignedRegion": "koreasouth",
              "assignedZone": ""
            },
            "regionDetail": {
              "regionId": "koreasouth",
              "regionName": "koreasouth",
              "description": "Korea South",
              "location": {
                "display": "Korea South",
                "latitude": 35.1796,
                "longitude": 129.0756
              },
              "zones": []
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "azure+koreasouth+standard_f2s_v2",
          "cspSpecName": "Standard_F2s_v2",
          "spec": {
            "cspSpecName": "Standard_F2s_v2",
            "vCPU": 2,
            "memoryGiB": 4,
            "costPerHour": 0.0961
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm",
          "networkInterface": "tbd32ihs5kcb30u4ok9r-15740-VNic",
          "securityGroupIds": [
            "my02-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my04-infra101",
      "uid": "tbeamcb9hb2d3rp8n5a3",
      "name": "my04-infra101",
      "status": "Running:3 (R:3/3)",
      "statusCount": {
        "countTotal": 3,
        "countCreating": 0,
        "countRunning": 3,
        "countFailed": 0,
        "countSuspended": 0,
        "countRebooting": 0,
        "countTerminated": 0,
        "countSuspending": 0,
        "countResuming": 0,
        "countTerminating": 0,
        "countRegistering": 0,
        "countReconciling": 0,
        "countUndefined": 0
      },
      "targetStatus": "None",
      "targetAction": "None",
      "configureCloudAdaptiveNetwork": "",
      "label": null,
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my04-ng-influxdb-back-1",
          "uid": "tbpthhohn7kjt5n614un",
          "cspResourceName": "tbpthhohn7kjt5n614un",
          "cspResourceId": "i-mj7f24f44ih73lbcha5e",
          "name": "my04-ng-influxdb-back-1",
          "nodeGroupId": "my04-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:12",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-11-25 09:16:12",
            "sys.cspResourceId": "i-mj7f24f44ih73lbcha5e",
            "sys.cspResourceName": "tbpthhohn7kjt5n614un",
            "sys.id": "my04-ng-influxdb-back-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-influxdb-back",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbpthhohn7kjt5n614un",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.213.148.200",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.3",
          "privateDNS": "",
          "rootDiskType": "cloud_essd_entry",
          "rootDiskSize": 40,
          "rootDeviceName": "",
          "connectionName": "alibaba-ap-northeast-2",
          "connectionConfig": {
            "configName": "alibaba-ap-northeast-2",
            "providerName": "alibaba",
            "driverName": "alibaba-driver-v1.0.so",
            "credentialName": "alibaba",
            "credentialHolder": "admin",
            "regionZoneInfoName": "alibaba-ap-northeast-2",
            "regionZoneInfo": {
              "assignedRegion": "ap-northeast-2",
              "assignedZone": "ap-northeast-2a"
            },
            "regionDetail": {
              "regionId": "ap-northeast-2",
              "regionName": "ap-northeast-2",
              "description": "South Korea (Seoul)",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.36,
                "longitude": 126.78
              },
              "zones": [
                "ap-northeast-2a",
                "ap-northeast-2b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "cspSpecName": "ecs.e-c1m4.xlarge",
          "spec": {
            "cspSpecName": "ecs.e-c1m4.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1582
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
          "networkInterface": "eni-mj7f24f44ih73lb9htrf",
          "securityGroupIds": [
            "my04-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tbskggpajj56tq2usk2e"
        },
        {
          "resourceType": "node",
          "id": "my04-ng-influxdb-back-2",
          "uid": "tbbj2qfadajevs4b9t50",
          "cspResourceName": "tbbj2qfadajevs4b9t50",
          "cspResourceId": "i-mj7803i8w5ifq7vxuxuc",
          "name": "my04-ng-influxdb-back-2",
          "nodeGroupId": "my04-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:11",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-11-25 09:16:11",
            "sys.cspResourceId": "i-mj7803i8w5ifq7vxuxuc",
            "sys.cspResourceName": "tbbj2qfadajevs4b9t50",
            "sys.id": "my04-ng-influxdb-back-2",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-influxdb-back",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbbj2qfadajevs4b9t50",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.61.166",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.2",
          "privateDNS": "",
          "rootDiskType": "cloud_essd_entry",
          "rootDiskSize": 40,
          "rootDeviceName": "",
          "connectionName": "alibaba-ap-northeast-2",
          "connectionConfig": {
            "configName": "alibaba-ap-northeast-2",
            "providerName": "alibaba",
            "driverName": "alibaba-driver-v1.0.so",
            "credentialName": "alibaba",
            "credentialHolder": "admin",
            "regionZoneInfoName": "alibaba-ap-northeast-2",
            "regionZoneInfo": {
              "assignedRegion": "ap-northeast-2",
              "assignedZone": "ap-northeast-2a"
            },
            "regionDetail": {
              "regionId": "ap-northeast-2",
              "regionName": "ap-northeast-2",
              "description": "South Korea (Seoul)",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.36,
                "longitude": 126.78
              },
              "zones": [
                "ap-northeast-2a",
                "ap-northeast-2b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "cspSpecName": "ecs.e-c1m4.xlarge",
          "spec": {
            "cspSpecName": "ecs.e-c1m4.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1582
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
          "networkInterface": "eni-mj7803i8w5ifq7vyy4l2",
          "securityGroupIds": [
            "my04-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tbskggpajj56tq2usk2e"
        },
        {
          "resourceType": "node",
          "id": "my04-ng-ip-10-0-1-30-62161624-1",
          "uid": "tbrvt5o72aejgu84q2d9",
          "cspResourceName": "tbrvt5o72aejgu84q2d9",
          "cspResourceId": "i-mj7ja9zhp5zysim4twue",
          "name": "my04-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my04-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:11",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-11-25 09:16:11",
            "sys.cspResourceId": "i-mj7ja9zhp5zysim4twue",
            "sys.cspResourceName": "tbrvt5o72aejgu84q2d9",
            "sys.id": "my04-ng-ip-10-0-1-30-62161624-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-ip-10-0-1-30-62161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-ip-10-0-1-30-62161624",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbrvt5o72aejgu84q2d9",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.220.218.42",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.1",
          "privateDNS": "",
          "rootDiskType": "cloud_auto",
          "rootDiskSize": 40,
          "rootDeviceName": "",
          "connectionName": "alibaba-ap-northeast-2",
          "connectionConfig": {
            "configName": "alibaba-ap-northeast-2",
            "providerName": "alibaba",
            "driverName": "alibaba-driver-v1.0.so",
            "credentialName": "alibaba",
            "credentialHolder": "admin",
            "regionZoneInfoName": "alibaba-ap-northeast-2",
            "regionZoneInfo": {
              "assignedRegion": "ap-northeast-2",
              "assignedZone": "ap-northeast-2a"
            },
            "regionDetail": {
              "regionId": "ap-northeast-2",
              "regionName": "ap-northeast-2",
              "description": "South Korea (Seoul)",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.36,
                "longitude": 126.78
              },
              "zones": [
                "ap-northeast-2a",
                "ap-northeast-2b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "cspSpecName": "ecs.e-c1m1.large",
          "spec": {
            "cspSpecName": "ecs.e-c1m1.large",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.0178
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
          "networkInterface": "eni-mj7ja9zhp5zysim0y47s",
          "securityGroupIds": [
            "my04-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tbskggpajj56tq2usk2e"
        }
      ],
      "newNodeList": null
    }
  ]
}
```

</details>

### Test Case 5: Get a list of infra IDs

#### 5.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra?option=id`
- **Purpose**: Get a list of IDs of all migrated infrastructures

#### 5.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **IDs**: [my01-infra101 my02-infra101 my04-infra101]

### Test Case 6: Get a specific infra

#### 6.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra/{{infraId}}`
#### 6.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my02-infra101",
  "uid": "tb3t1dqr76ce6f3l49ug",
  "name": "my02-infra101",
  "status": "Running:3 (R:3/3)",
  "statusCount": {
    "countTotal": 3,
    "countCreating": 0,
    "countRunning": 3,
    "countFailed": 0,
    "countSuspended": 0,
    "countRebooting": 0,
    "countTerminated": 0,
    "countSuspending": 0,
    "countResuming": 0,
    "countTerminating": 0,
    "countRegistering": 0,
    "countReconciling": 0,
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "cm-migrated-at": "2026-11-25T09:15:11Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my02-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my02-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tb3t1dqr76ce6f3l49ug"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my02-ng-influxdb-back-1",
      "uid": "tb2mefhc7qac6sg6gmoh",
      "cspResourceName": "tb2mefhc7qac6sg6gmoh",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb2mefhc7qac6sg6gmoh",
      "name": "my02-ng-influxdb-back-1",
      "nodeGroupId": "my02-ng-influxdb-back",
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:03",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "createdBy": "tb2mefhc7qac6sg6gmoh",
        "keypair": "tbbm9d00plcuabc5g4jg",
        "nlbBackend": "influxdb_back",
        "publicip": "tb2mefhc7qac6sg6gmoh-91708-PublicIP"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "52.147.125.128",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "StandardHDD",
      "rootDiskSize": 30,
      "rootDeviceName": "Not visible in Azure",
      "connectionName": "azure-koreasouth",
      "connectionConfig": {
        "configName": "azure-koreasouth",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreasouth",
        "regionZoneInfo": {
          "assignedRegion": "koreasouth",
          "assignedZone": ""
        },
        "regionDetail": {
          "regionId": "koreasouth",
          "regionName": "koreasouth",
          "description": "Korea South",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "zones": []
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreasouth+standard_d4_v4",
      "cspSpecName": "Standard_D4_v4",
      "spec": {
        "cspSpecName": "Standard_D4_v4",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.221
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm",
      "networkInterface": "tb2mefhc7qac6sg6gmoh-61943-VNic",
      "securityGroupIds": [
        "my02-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJ66k1oYuwBtYA5NOE/x6ieyZkjOF/fzk5ySLbd9GfcC/PJnflIiuyOBELFTzQtzN8k+Z6Sqwyad2RitrdzH+Xw=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:9AeffaLS0Nj2SNqwUICjEFi7i04E4T3ap6yI3NXJYow",
        "firstUsedAt": "2026-11-25T09:16:07Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:03Z",
          "completedTime": "2026-11-25T09:16:09Z",
          "elapsedTime": 6,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my02-infra101-tb7q121q173qpuoqlvva",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:11Z",
          "completedTime": "2026-11-25T09:16:12Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb2mefhc7qac6sg6gmoh 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my02-ng-influxdb-back-2",
      "uid": "tb6lhvd6gifs72gf64ok",
      "cspResourceName": "tb6lhvd6gifs72gf64ok",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb6lhvd6gifs72gf64ok",
      "name": "my02-ng-influxdb-back-2",
      "nodeGroupId": "my02-ng-influxdb-back",
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:03",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "createdBy": "tb6lhvd6gifs72gf64ok",
        "keypair": "tbbm9d00plcuabc5g4jg",
        "nlbBackend": "influxdb_back",
        "publicip": "tb6lhvd6gifs72gf64ok-39512-PublicIP"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "74.227.129.35",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "StandardHDD",
      "rootDiskSize": 30,
      "rootDeviceName": "Not visible in Azure",
      "connectionName": "azure-koreasouth",
      "connectionConfig": {
        "configName": "azure-koreasouth",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreasouth",
        "regionZoneInfo": {
          "assignedRegion": "koreasouth",
          "assignedZone": ""
        },
        "regionDetail": {
          "regionId": "koreasouth",
          "regionName": "koreasouth",
          "description": "Korea South",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "zones": []
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreasouth+standard_d4_v4",
      "cspSpecName": "Standard_D4_v4",
      "spec": {
        "cspSpecName": "Standard_D4_v4",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.221
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm",
      "networkInterface": "tb6lhvd6gifs72gf64ok-32847-VNic",
      "securityGroupIds": [
        "my02-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMRFzcn/w5rroLqFcUp0u5fnG2nZSHGUDEzCBNRmmFtQySyMPoQjhIXdh4WP9ordWlEmwgX/8rwHyOuWuGlbBlI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:j/BfQxwdAH7y6IarUawWD7py8RYnF3uvhw4bxtAz9Rw",
        "firstUsedAt": "2026-11-25T09:16:07Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:03Z",
          "completedTime": "2026-11-25T09:16:10Z",
          "elapsedTime": 7,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my02-infra101-tb7q121q173qpuoqlvva",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:11Z",
          "completedTime": "2026-11-25T09:16:13Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb6lhvd6gifs72gf64ok 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my02-ng-ip-10-0-1-30-62161624-1",
      "uid": "tbd32ihs5kcb30u4ok9r",
      "cspResourceName": "tbd32ihs5kcb30u4ok9r",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd32ihs5kcb30u4ok9r",
      "name": "my02-ng-ip-10-0-1-30-62161624-1",
      "nodeGroupId": "my02-ng-ip-10-0-1-30-62161624",
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:03",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
        "createdBy": "tbd32ihs5kcb30u4ok9r",
        "keypair": "tbbm9d00plcuabc5g4jg",
        "publicip": "tbd32ihs5kcb30u4ok9r-53594-PublicIP"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "74.227.128.223",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 30,
      "rootDeviceName": "Not visible in Azure",
      "connectionName": "azure-koreasouth",
      "connectionConfig": {
        "configName": "azure-koreasouth",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreasouth",
        "regionZoneInfo": {
          "assignedRegion": "koreasouth",
          "assignedZone": ""
        },
        "regionDetail": {
          "regionId": "koreasouth",
          "regionName": "koreasouth",
          "description": "Korea South",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "zones": []
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreasouth+standard_f2s_v2",
      "cspSpecName": "Standard_F2s_v2",
      "spec": {
        "cspSpecName": "Standard_F2s_v2",
        "vCPU": 2,
        "memoryGiB": 4,
        "costPerHour": 0.0961
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm",
      "networkInterface": "tbd32ihs5kcb30u4ok9r-15740-VNic",
      "securityGroupIds": [
        "my02-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHI+eba+Z/Dhzv6Jm8MotVrH8Kx1nA1pM1RNUFZqkym1Mz372/ksu/DU8NcqcUPNw3/JdHZLSC8GTlwjQHeXZCI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:RVYMe+sr8Bqh2p9XkxhkP47Be/vb/LbfCGmnjWUblf4",
        "firstUsedAt": "2026-11-25T09:16:04Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:03Z",
          "completedTime": "2026-11-25T09:16:11Z",
          "elapsedTime": 8,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my02-infra101-tb7q121q173qpuoqlvva",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:11Z",
          "completedTime": "2026-11-25T09:16:14Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd32ihs5kcb30u4ok9r 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my02-mig-vnet-01",
      "name": "my02-mig-vnet-01",
      "infraId": "my02-infra101",
      "vNetId": "my02-mig-vnet-01",
      "connectionNames": [
        "azure-koreasouth"
      ],
      "providerNames": [
        "azure"
      ],
      "regionNames": [
        "koreasouth"
      ],
      "nodeGroupIds": [
        "my02-ng-influxdb-back",
        "my02-ng-ip-10-0-1-30-62161624"
      ],
      "nodeIds": [
        "my02-ng-influxdb-back-1",
        "my02-ng-influxdb-back-2",
        "my02-ng-ip-10-0-1-30-62161624-1"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my02-ng-influxdb-back",
      "representativeNodeId": "my02-ng-influxdb-back-1"
    }
  ],
  "newNodeList": null,
  "postScript": {
    "phases": [
      {
        "userName": "cb-user",
        "command": [
          "uname -a"
        ]
      }
    ]
  },
  "postScriptResults": [
    {
      "phase": 1,
      "target": "all nodes",
      "status": "Completed",
      "results": {
        "results": [
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-influxdb-back-1",
            "nodeIp": "52.147.125.128",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb2mefhc7qac6sg6gmoh 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-influxdb-back-2",
            "nodeIp": "74.227.129.35",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb6lhvd6gifs72gf64ok 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-ip-10-0-1-30-62161624-1",
            "nodeIp": "74.227.128.223",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbd32ihs5kcb30u4ok9r 6.8.0-1073-azure #82~22.04.1-Ubuntu SMP Tue Oct 27 22:13:07 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          }
        ]
      }
    }
  ],
  "postScriptStatus": "Completed",
  "postScriptRequestId": "ps-my02-infra101-tb7q121q173qpuoqlvva"
}
```

</details>

### Test Case 7: Remote Command Accessibility Check

#### 7.1 Test Information

- **Test Type**: SSH Connectivity Test for All VMs
- **Command Executed**: `uname -a` (to verify system information)

#### 7.2 Test Result Information

- **Status**: ✅ **SUCCESS**
**Summary**: 3/3 VMs accessible via SSH

### Test Case 8: Migrate NLBs to the cloud infra

#### 8.1 API Request Information

- **API Endpoint**: `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb`
- **Purpose**: Create target load balancers mapped from source HAProxy configuration

#### 8.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **NLB Status**: `created`
- **Description**: `1 NLB(s) created successfully`

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "status": "created",
  "description": "1 NLB(s) created successfully",
  "nlbList": [
    {
      "resourceType": "",
      "id": "my02-ng-influxdb-back",
      "cspResourceName": "tbbjvscnut4iq5kjrmj4",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4",
      "name": "my02-ng-influxdb-back",
      "connectionName": "azure-koreasouth",
      "connectionConfig": {
        "configName": "azure-koreasouth",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreasouth",
        "regionZoneInfo": {
          "assignedRegion": "koreasouth",
          "assignedZone": ""
        },
        "regionDetail": {
          "regionId": "koreasouth",
          "regionName": "koreasouth",
          "description": "Korea South",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "zones": []
        },
        "regionRepresentative": true,
        "verified": true
      },
      "type": "PUBLIC",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "ip": "20.200.190.98",
        "port": "9999",
        "keyValueList": [
          {
            "key": "FrontendPort",
            "value": "9999"
          },
          {
            "key": "Protocol",
            "value": "Tcp"
          },
          {
            "key": "BackendAddressPool",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
          },
          {
            "key": "BackendAddressPools",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
          },
          {
            "key": "BackendPort",
            "value": "8086"
          },
          {
            "key": "DisableOutboundSnat",
            "value": "false"
          },
          {
            "key": "EnableConnectionTracking",
            "value": "false"
          },
          {
            "key": "EnableFloatingIP",
            "value": "false"
          },
          {
            "key": "EnableTCPReset",
            "value": "false"
          },
          {
            "key": "FrontendIPConfiguration",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404}"
          },
          {
            "key": "IdleTimeoutInMinutes",
            "value": "4"
          },
          {
            "key": "LoadDistribution",
            "value": "Default"
          },
          {
            "key": "Probe",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695}"
          },
          {
            "key": "ProvisioningState",
            "value": "Succeeded"
          }
        ]
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my02-ng-influxdb-back",
        "nodes": [
          "my02-ng-influxdb-back-1",
          "my02-ng-influxdb-back-2"
        ],
        "keyValueList": [
          {
            "key": "FrontendPort",
            "value": "9999"
          },
          {
            "key": "Protocol",
            "value": "Tcp"
          },
          {
            "key": "BackendAddressPool",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
          },
          {
            "key": "BackendAddressPools",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
          },
          {
            "key": "BackendPort",
            "value": "8086"
          },
          {
            "key": "DisableOutboundSnat",
            "value": "false"
          },
          {
            "key": "EnableConnectionTracking",
            "value": "false"
          },
          {
            "key": "EnableFloatingIP",
            "value": "false"
          },
          {
            "key": "EnableTCPReset",
            "value": "false"
          },
          {
            "key": "FrontendIPConfiguration",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404}"
          },
          {
            "key": "IdleTimeoutInMinutes",
            "value": "4"
          },
          {
            "key": "LoadDistribution",
            "value": "Default"
          },
          {
            "key": "Probe",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695}"
          },
          {
            "key": "ProvisioningState",
            "value": "Succeeded"
          }
        ]
      },
      "healthChecker": {
        "protocol": "TCP",
        "port": "8086",
        "interval": 10,
        "threshold": 3,
        "timeout": -1
      },
      "createdTime": "2026-11-25T09:16:44Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4"
        },
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{backendAddressPools:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965,name:backend-25965,properties:{loadBalancerBackendAddresses:[{name:backend-2596510.0.1.6,properties:{ipAddress:10.0.1.6,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9}}},{name:backend-2596510.0.1.4,properties:{ipAddress:10.0.1.4,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9}}}],loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666}],provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/backendAddressPools}],frontendIPConfigurations:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404,name:frontEndIp-363404,properties:{loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666}],privateIPAllocationMethod:Dynamic,provisioningState:Succeeded,publicIPAddress:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/publicIPAddresses/tbbjvscnut4iq5kjrmj4}},type:Microsoft.Network/loadBalancers/frontendIPConfigurations}],inboundNatPools:[],inboundNatRules:[],loadBalancingRules:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666,name:lbrule-667666,properties:{backendAddressPool:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965},backendAddressPools:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}],backendPort:8086,disableOutboundSnat:false,enableConnectionTracking:false,enableFloatingIP:false,enableTcpReset:false,frontendIPConfiguration:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404},frontendPort:9999,idleTimeoutInMinutes:4,loadDistribution:Default,probe:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695},protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/loadBalancingRules}],outboundRules:[],probes:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695,name:probe-602695,properties:{intervalInSeconds:10,loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666}],noHealthyBackendsBehavior:AllProbedDown,numberOfProbes:3,port:8086,probeThreshold:1,protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/probes}],provisioningState:Succeeded,resourceGuid:a74d37b4-5311-4b31-8ae9-0825972de969}"
        },
        {
          "key": "SKU",
          "value": "{name:Standard,tier:Regional}"
        },
        {
          "key": "Tags",
          "value": "{createdAt:1795598204}"
        },
        {
          "key": "Etag",
          "value": "W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\"
        },
        {
          "key": "Name",
          "value": "tbbjvscnut4iq5kjrmj4"
        },
        {
          "key": "Type",
          "value": "Microsoft.Network/loadBalancers"
        }
      ],
      "isAutoGenerated": false,
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      }
    }
  ]
}
```

</details>

### Test Case 9: Get a list of migrated NLBs

#### 9.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb`

#### 9.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
[
  {
    "resourceType": "",
    "id": "my02-ng-influxdb-back",
    "cspResourceName": "tbbjvscnut4iq5kjrmj4",
    "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4",
    "name": "my02-ng-influxdb-back",
    "connectionName": "azure-koreasouth",
    "connectionConfig": {
      "configName": "azure-koreasouth",
      "providerName": "azure",
      "driverName": "azure-driver-v1.0.so",
      "credentialName": "azure",
      "credentialHolder": "admin",
      "regionZoneInfoName": "azure-koreasouth",
      "regionZoneInfo": {
        "assignedRegion": "koreasouth",
        "assignedZone": ""
      },
      "regionDetail": {
        "regionId": "koreasouth",
        "regionName": "koreasouth",
        "description": "Korea South",
        "location": {
          "display": "Korea South",
          "latitude": 35.1796,
          "longitude": 129.0756
        },
        "zones": []
      },
      "regionRepresentative": true,
      "verified": true
    },
    "type": "PUBLIC",
    "scope": "REGION",
    "listener": {
      "protocol": "TCP",
      "ip": "20.200.190.98",
      "port": "9999",
      "keyValueList": [
        {
          "key": "FrontendPort",
          "value": "9999"
        },
        {
          "key": "Protocol",
          "value": "Tcp"
        },
        {
          "key": "BackendAddressPool",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
        },
        {
          "key": "BackendAddressPools",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
        },
        {
          "key": "BackendPort",
          "value": "8086"
        },
        {
          "key": "DisableOutboundSnat",
          "value": "false"
        },
        {
          "key": "EnableConnectionTracking",
          "value": "false"
        },
        {
          "key": "EnableFloatingIP",
          "value": "false"
        },
        {
          "key": "EnableTCPReset",
          "value": "false"
        },
        {
          "key": "FrontendIPConfiguration",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404}"
        },
        {
          "key": "IdleTimeoutInMinutes",
          "value": "4"
        },
        {
          "key": "LoadDistribution",
          "value": "Default"
        },
        {
          "key": "Probe",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695}"
        },
        {
          "key": "ProvisioningState",
          "value": "Succeeded"
        }
      ]
    },
    "targetGroup": {
      "protocol": "TCP",
      "port": "8086",
      "nodeGroupId": "my02-ng-influxdb-back",
      "nodes": [
        "my02-ng-influxdb-back-1",
        "my02-ng-influxdb-back-2"
      ],
      "keyValueList": [
        {
          "key": "FrontendPort",
          "value": "9999"
        },
        {
          "key": "Protocol",
          "value": "Tcp"
        },
        {
          "key": "BackendAddressPool",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
        },
        {
          "key": "BackendAddressPools",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
        },
        {
          "key": "BackendPort",
          "value": "8086"
        },
        {
          "key": "DisableOutboundSnat",
          "value": "false"
        },
        {
          "key": "EnableConnectionTracking",
          "value": "false"
        },
        {
          "key": "EnableFloatingIP",
          "value": "false"
        },
        {
          "key": "EnableTCPReset",
          "value": "false"
        },
        {
          "key": "FrontendIPConfiguration",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404}"
        },
        {
          "key": "IdleTimeoutInMinutes",
          "value": "4"
        },
        {
          "key": "LoadDistribution",
          "value": "Default"
        },
        {
          "key": "Probe",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695}"
        },
        {
          "key": "ProvisioningState",
          "value": "Succeeded"
        }
      ]
    },
    "healthChecker": {
      "protocol": "TCP",
      "port": "8086",
      "interval": 10,
      "threshold": 3,
      "timeout": -1
    },
    "createdTime": "2026-11-25T09:16:44Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "ID",
        "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4"
      },
      {
        "key": "Location",
        "value": "koreasouth"
      },
      {
        "key": "Properties",
        "value": "{backendAddressPools:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965,name:backend-25965,properties:{loadBalancerBackendAddresses:[{name:backend-2596510.0.1.6,properties:{ipAddress:10.0.1.6,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9}}},{name:backend-2596510.0.1.4,properties:{ipAddress:10.0.1.4,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9}}}],loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666}],provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/backendAddressPools}],frontendIPConfigurations:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404,name:frontEndIp-363404,properties:{loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666}],privateIPAllocationMethod:Dynamic,provisioningState:Succeeded,publicIPAddress:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/publicIPAddresses/tbbjvscnut4iq5kjrmj4}},type:Microsoft.Network/loadBalancers/frontendIPConfigurations}],inboundNatPools:[],inboundNatRules:[],loadBalancingRules:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666,name:lbrule-667666,properties:{backendAddressPool:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965},backendAddressPools:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}],backendPort:8086,disableOutboundSnat:false,enableConnectionTracking:false,enableFloatingIP:false,enableTcpReset:false,frontendIPConfiguration:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404},frontendPort:9999,idleTimeoutInMinutes:4,loadDistribution:Default,probe:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695},protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/loadBalancingRules}],outboundRules:[],probes:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695,name:probe-602695,properties:{intervalInSeconds:10,loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666}],noHealthyBackendsBehavior:AllProbedDown,numberOfProbes:3,port:8086,probeThreshold:1,protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/probes}],provisioningState:Succeeded,resourceGuid:a74d37b4-5311-4b31-8ae9-0825972de969}"
      },
      {
        "key": "SKU",
        "value": "{name:Standard,tier:Regional}"
      },
      {
        "key": "Tags",
        "value": "{createdAt:1795598204}"
      },
      {
        "key": "Etag",
        "value": "W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\"
      },
      {
        "key": "Name",
        "value": "tbbjvscnut4iq5kjrmj4"
      },
      {
        "key": "Type",
        "value": "Microsoft.Network/loadBalancers"
      }
    ],
    "isAutoGenerated": false,
    "location": {
      "display": "Korea South",
      "latitude": 35.1796,
      "longitude": 129.0756
    }
  }
]
```

</details>

### Test Case 10: Get details of a specific migrated NLB

#### 10.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}`

#### 10.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "",
  "id": "my02-ng-influxdb-back",
  "cspResourceName": "tbbjvscnut4iq5kjrmj4",
  "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4",
  "name": "my02-ng-influxdb-back",
  "connectionName": "azure-koreasouth",
  "connectionConfig": {
    "configName": "azure-koreasouth",
    "providerName": "azure",
    "driverName": "azure-driver-v1.0.so",
    "credentialName": "azure",
    "credentialHolder": "admin",
    "regionZoneInfoName": "azure-koreasouth",
    "regionZoneInfo": {
      "assignedRegion": "koreasouth",
      "assignedZone": ""
    },
    "regionDetail": {
      "regionId": "koreasouth",
      "regionName": "koreasouth",
      "description": "Korea South",
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      },
      "zones": []
    },
    "regionRepresentative": true,
    "verified": true
  },
  "type": "PUBLIC",
  "scope": "REGION",
  "listener": {
    "protocol": "TCP",
    "ip": "20.200.190.98",
    "port": "9999",
    "keyValueList": [
      {
        "key": "FrontendPort",
        "value": "9999"
      },
      {
        "key": "Protocol",
        "value": "Tcp"
      },
      {
        "key": "BackendAddressPool",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
      },
      {
        "key": "BackendAddressPools",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
      },
      {
        "key": "BackendPort",
        "value": "8086"
      },
      {
        "key": "DisableOutboundSnat",
        "value": "false"
      },
      {
        "key": "EnableConnectionTracking",
        "value": "false"
      },
      {
        "key": "EnableFloatingIP",
        "value": "false"
      },
      {
        "key": "EnableTCPReset",
        "value": "false"
      },
      {
        "key": "FrontendIPConfiguration",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404}"
      },
      {
        "key": "IdleTimeoutInMinutes",
        "value": "4"
      },
      {
        "key": "LoadDistribution",
        "value": "Default"
      },
      {
        "key": "Probe",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695}"
      },
      {
        "key": "ProvisioningState",
        "value": "Succeeded"
      }
    ]
  },
  "targetGroup": {
    "protocol": "TCP",
    "port": "8086",
    "nodeGroupId": "my02-ng-influxdb-back",
    "nodes": [
      "my02-ng-influxdb-back-1",
      "my02-ng-influxdb-back-2"
    ],
    "keyValueList": [
      {
        "key": "FrontendPort",
        "value": "9999"
      },
      {
        "key": "Protocol",
        "value": "Tcp"
      },
      {
        "key": "BackendAddressPool",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
      },
      {
        "key": "BackendAddressPools",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}"
      },
      {
        "key": "BackendPort",
        "value": "8086"
      },
      {
        "key": "DisableOutboundSnat",
        "value": "false"
      },
      {
        "key": "EnableConnectionTracking",
        "value": "false"
      },
      {
        "key": "EnableFloatingIP",
        "value": "false"
      },
      {
        "key": "EnableTCPReset",
        "value": "false"
      },
      {
        "key": "FrontendIPConfiguration",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404}"
      },
      {
        "key": "IdleTimeoutInMinutes",
        "value": "4"
      },
      {
        "key": "LoadDistribution",
        "value": "Default"
      },
      {
        "key": "Probe",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695}"
      },
      {
        "key": "ProvisioningState",
        "value": "Succeeded"
      }
    ]
  },
  "healthChecker": {
    "protocol": "TCP",
    "port": "8086",
    "interval": 10,
    "threshold": 3,
    "timeout": -1
  },
  "createdTime": "2026-11-25T09:16:44Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "ID",
      "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4"
    },
    {
      "key": "Location",
      "value": "koreasouth"
    },
    {
      "key": "Properties",
      "value": "{backendAddressPools:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965,name:backend-25965,properties:{loadBalancerBackendAddresses:[{name:backend-2596510.0.1.6,properties:{ipAddress:10.0.1.6,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9}}},{name:backend-2596510.0.1.4,properties:{ipAddress:10.0.1.4,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9}}}],loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666}],provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/backendAddressPools}],frontendIPConfigurations:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404,name:frontEndIp-363404,properties:{loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666}],privateIPAllocationMethod:Dynamic,provisioningState:Succeeded,publicIPAddress:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/publicIPAddresses/tbbjvscnut4iq5kjrmj4}},type:Microsoft.Network/loadBalancers/frontendIPConfigurations}],inboundNatPools:[],inboundNatRules:[],loadBalancingRules:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666,name:lbrule-667666,properties:{backendAddressPool:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965},backendAddressPools:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/backendAddressPools/backend-25965}],backendPort:8086,disableOutboundSnat:false,enableConnectionTracking:false,enableFloatingIP:false,enableTcpReset:false,frontendIPConfiguration:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/frontendIPConfigurations/frontEndIp-363404},frontendPort:9999,idleTimeoutInMinutes:4,loadDistribution:Default,probe:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695},protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/loadBalancingRules}],outboundRules:[],probes:[{etag:W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/probes/probe-602695,name:probe-602695,properties:{intervalInSeconds:10,loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4/loadBalancingRules/lbrule-667666}],noHealthyBackendsBehavior:AllProbedDown,numberOfProbes:3,port:8086,probeThreshold:1,protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/probes}],provisioningState:Succeeded,resourceGuid:a74d37b4-5311-4b31-8ae9-0825972de969}"
    },
    {
      "key": "SKU",
      "value": "{name:Standard,tier:Regional}"
    },
    {
      "key": "Tags",
      "value": "{createdAt:1795598204}"
    },
    {
      "key": "Etag",
      "value": "W/\\fdcd6318-a3d6-4a8b-a4f6-e02a197d43c5\\"
    },
    {
      "key": "Name",
      "value": "tbbjvscnut4iq5kjrmj4"
    },
    {
      "key": "Type",
      "value": "Microsoft.Network/loadBalancers"
    }
  ],
  "isAutoGenerated": false,
  "location": {
    "display": "Korea South",
    "latitude": 35.1796,
    "longitude": 129.0756
  }
}
```

</details>

### Test Case 11: NLB Load Balancing Verification

#### 11.1 Test Information

- **Test Type**: Active Traffic Distribution Verification via NLB Endpoint
- **Target Port**: `8086` (Backend Mock Web Server)
- **Listener Port**: `9999` (NLB Listener)
- **Requests Sent**: 15 HTTP GET requests

#### 11.2 Test Result Information

- **Status**: ✅ **SUCCESS**

**Summary**: Load balancing confirmed: traffic distributed across 2 unique VMs

### Test Case 12: Target Infrastructure Summary

#### 12.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Target Infrastructure Summary**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:22:43

**Namespace:** mig01

**Infra Name:** my02-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my02-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | AZURE |
| **Target Region** | koreasouth |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| Standard_D4_v4 | 4 | 16.0 | - | x86_64 |  | $0.2210 | 2 |
| Standard_F2s_v2 | 2 | 4.0 | - | x86_64 |  | $0.0961 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my02-ng-influxdb-back-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb2mefhc7qac6sg6gmoh | Running | 4 vCPU, 16.0 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 52.147.125.128<br>**Private IP:** 10.0.1.6<br>**SGs:** my02-mig-sg-01<br>**SSH:** my02-mig-sshkey-01 |
| my02-ng-influxdb-back-2 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb6lhvd6gifs72gf64ok | Running | 4 vCPU, 16.0 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 74.227.129.35<br>**Private IP:** 10.0.1.4<br>**SGs:** my02-mig-sg-01<br>**SSH:** my02-mig-sshkey-01 |
| my02-ng-ip-10-0-1-30-62161624-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd32ihs5kcb30u4ok9r | Running | 2 vCPU, 4.0 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 74.227.128.223<br>**Private IP:** 10.0.1.5<br>**SGs:** my02-mig-sg-02<br>**SSH:** my02-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my02-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my02-mig-vnet-01 |
| **CSP VNet ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | azure-koreasouth |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my02-mig-subnet-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm | 10.0.1.0/24 |  |

### Network Load Balancers (NLB)

#### NLB: my02-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my02-ng-influxdb-back |
| **CSP NLB Name** | tbbjvscnut4iq5kjrmj4 |
| **CSP NLB ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbbjvscnut4iq5kjrmj4 |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | azure-koreasouth |
| **Listener** | TCP:9999 (IP: 20.200.190.98) |
| **Target Group** | TCP:8086 (Node Group: my02-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: CSP Default, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my02-ng-influxdb-back-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb2mefhc7qac6sg6gmoh | 10.0.1.6 | 52.147.125.128 | Running |
| my02-ng-influxdb-back-2 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb6lhvd6gifs72gf64ok | 10.0.1.4 | 74.227.129.35 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my02-mig-sshkey-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg |  |  |

### Security Groups

#### Security Group: my02-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my02-mig-sg-01 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tb77r547caqdgeolc338 |
| **VNet** | my02-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my02-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my02-mig-sg-02 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd34sc0anir1ka300ki |
| **VNet** | my02-mig-vnet-01 |
| **Rule Count** | 4 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.5381 |
| **Per Day** | $12.91 |
| **Per Month (30 days)** | $387.43 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AZURE | koreasouth | 3 | $0.5381 | $387.43 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my02-ng-influxdb-back-1 | Standard_D4_v4 | $0.2210 | $159.12 |
| my02-ng-influxdb-back-2 | Standard_D4_v4 | $0.2210 | $159.12 |
| my02-ng-ip-10-0-1-30-62161624-1 | Standard_F2s_v2 | $0.0961 | $69.19 |




### Test Case 13: Migration Report

#### 13.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 13.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-11-25 09:22:48*

---

## 📊 Migration Summary

**Target Cloud:** AZURE

**Target Region:** koreasouth

**Namespace:** mig01 | **Infra ID:** my02-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**NLB Migration Status:** Completed

**Total Load Balancers:** 1

**Migrated Load Balancers:** 1

**💰 Estimated Monthly Cost:** $387.43 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | Standard_D4_v4, Standard_F2s_v2 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my02-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my02-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 10 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |
| 8 | **Network Load Balancer** | 1 | ✅ Created | my02-ng-influxdb-back (TCP:9999 -> TCP:8086) |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my02-ng-influxdb-back-1<br>**VM ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb2mefhc7qac6sg6gmoh<br>**Label(sourceMachineId):** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 |
| 2 | **VM Name:** my02-ng-influxdb-back-2<br>**VM ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb6lhvd6gifs72gf64ok<br>**Label(sourceMachineId):** ec288dd0-c6fa-8a49-2f60-bc898311febf | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf |
| 3 | **VM Name:** my02-ng-ip-10-0-1-30-62161624-1<br>**VM ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd32ihs5kcb30u4ok9r<br>**Label(sourceMachineId):** ec268ed7-821e-9d73-e79f-961262161624 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my02-ng-influxdb-back-1 | **Spec ID:** Standard_D4_v4<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 30 GB |
| 2 | my02-ng-influxdb-back-2 | **Spec ID:** Standard_D4_v4<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 8 GB<br>**Root Disk:** 30 GB |
| 3 | my02-ng-ip-10-0-1-30-62161624-1 | **Spec ID:** Standard_F2s_v2<br>**vCPUs:** 2<br>**Memory:** 4.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 8 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my02-ng-influxdb-back-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | my02-ng-influxdb-back-2 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | my02-ng-ip-10-0-1-30-62161624-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 10 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my02-mig-sg-01

**CSP ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tb77r547caqdgeolc338 | **VNet:** my02-mig-vnet-01 | **Rules:** 6

**Assigned VMs:**

- **VM:** my02-ng-influxdb-back-1
  - **Source Server:** **Hostname:** ip-10-0-1-221, **Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932
- **VM:** my02-ng-influxdb-back-2
  - **Source Server:** **Hostname:** ip-10-0-1-138, **Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 2 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | TCP | 9999 | 0.0.0.0/0 | - | Created by system |
| 5 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 6 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my02-mig-sg-02

**CSP ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd34sc0anir1ka300ki | **VNet:** my02-mig-vnet-01 | **Rules:** 4

**Assigned VMs:**

- **VM:** my02-ng-ip-10-0-1-30-62161624-1
  - **Source Server:** **Hostname:** ip-10-0-1-30, **Machine ID:** ec268ed7-821e-9d73-e79f-961262161624

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 2 | inbound | TCP | 9999 | 0.0.0.0/0 | inbound tcp 9999 | Migrated from source |
| 3 | inbound | ALL |  | 10.0.0.0/16 | inbound * * from 10.0.0.0/16 | Migrated from source |
| 4 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source node network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** my02-mig-vnet-01<br>**ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my02-mig-subnet-01<br>**ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbn5tmrldklp5kklaje9/subnets/tbeh520erobu2h1n1nfm | 10.0.1.0/24 | my02-mig-vnet-01 |

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

### 1. influxdb_back (haproxy) → my02-ng-influxdb-back

**Configuration Comparison:**

| Property | Source (On-Premise) | Target (Cloud NLB) | Note |
|----------|---------------------|--------------------|------|
| **Load Balancer / Name** | haproxy (Host: ip-10-0-1-30) | my02-ng-influxdb-back (CSP ID: tbbjvscnut4iq5kjrmj4) | Replaced by managed NLB |
| **Type / Scope** | Software LB | PUBLIC / REGION | Cloud-native NLB |
| **Frontend Listener** | *:9999 (TCP) | TCP:9999 | IP: 20.200.190.98 |
| **Backend / Target Group** | influxdb_back (TCP:8086, roundrobin) | my02-ng-influxdb-back (TCP:8086, 2 nodes) | Target Group attached to NodeGroup |
| **Health Check** | Enabled (Interval: 10s, Timeout: 10s, Threshold: 3) | TCP:8086 (Interval: 10s, Timeout: CSP Default, Threshold: 3) | Timeout uses CSP default |

**Backend Node Mappings:**

| No. | Source Backend Server | Source Endpoint | Target Node | Target CSP VM ID | Target Private IP | Target Public IP | Target Port | Status |
|-----|-----------------------|-----------------|-------------|------------------|-------------------|------------------|-------------|--------|
| 1 | influx1 (ip-10-0-1-221) | 10.0.1.221:8086 | my02-ng-influxdb-back-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb2mefhc7qac6sg6gmoh | 10.0.1.6 | 52.147.125.128 | 8086 | Running |
| 2 | influx2 (ip-10-0-1-138) | 10.0.1.138:8086 | my02-ng-influxdb-back-2 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb6lhvd6gifs72gf64ok | 10.0.1.4 | 74.227.129.35 | 8086 | Running |

**Migration Notes:**
- On-premise haproxy on ip-10-0-1-30 (migrated to my02-ng-ip-10-0-1-30-62161624-1) is replaced by cloud-managed NLB.
- Health check timeout uses CSP default setting on AZURE.

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my02-mig-sshkey-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tbbm9d00plcuabc5g4jg |  | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.5381 |
| Daily | $12.91 |
| Monthly | $387.43 |
| Yearly | $4649.18 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | Standard_F2s_v2 | $69.19 | 17.9% |
| ip-10-0-1-221 (migrated) | Standard_D4_v4 | $159.12 | 41.1% |
| ip-10-0-1-138 (migrated) | Standard_D4_v4 | $159.12 | 41.1% |

---


---

*Report generated by CM-Beetle*


### Test Case 14: Delete the migrated NLBs

#### 14.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}`

#### 14.2 API Response Information

- **Status**: ✅ **SUCCESS**
### Test Case 15: Delete the migrated computing infra

#### 15.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}?option=terminate`

#### 15.2 API Response Information

- **Status**: ✅ **SUCCESS**
**Response Body**:

```json
{
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my02-infra101)",
  "success": true
}
```

