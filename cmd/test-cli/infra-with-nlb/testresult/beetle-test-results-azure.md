# CM-Beetle test results for AZURE (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.2+ (83efe14)
- imdl: v0.1.19+ (83efe14)
- CB-Tumblebug: v0.13.8
- CB-Spider: v0.13.20
- CB-MapUI: v0.13.14
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: October 30, 2026
- Test Time: 21:45:54 KST
- Test Execution: 2026-10-30 21:45:54 KST

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
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 2.835s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 125ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 58.565s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 14ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 3ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 9ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 2.255s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 9.189s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 5ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 7ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 9m25.464s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.082s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.082s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 39.389s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m39.637s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 13m57.982052914s

*Test executed on October 30, 2026 at 21:45:54 KST (2026-10-30 21:45:54 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=51.2% Max=100.0% Avg=85.9% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "azure",
        "region": "koreasouth"
      },
      "targetInfra": {
        "name": "infra101",
        "installMonAgent": "",
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_d4s_v4",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
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
            "name": "ng-ec268ed7-821e-9d73-e79f-961262161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=83.3%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_f2s_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
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
          "id": "azure+koreasouth+standard_d4s_v4",
          "uid": "tboubh1qij20enfs0695",
          "cspSpecName": "Standard_D4s_v4",
          "name": "azure+koreasouth+standard_d4s_v4",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 15.625,
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
              "value": "Standard_D4s_v4"
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
              "value": "True"
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
              "key": "CachedDiskBytes",
              "value": "107374182400"
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
              "value": "standardDSv4Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "D4s_v4"
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
          "uid": "tbdk2p4rel9h9a71mibl",
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
          "memoryGiB": 3.90625,
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "uid": "tbm5bq012h0vl4rmsqjj",
          "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.08.21 14:01:58 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
              "value": "22.04.202608060"
            },
            {
              "key": "ID",
              "value": "/Subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202608060"
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
          "systemLabel": "",
          "description": "",
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=25.6% Max=100.0% Avg=73.3% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "azure",
        "region": "koreasouth"
      },
      "targetInfra": {
        "name": "infra101",
        "installMonAgent": "",
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=50.0% Memory=97.7% Image=83.3%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_e2s_v3",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
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
            "name": "ng-ec268ed7-821e-9d73-e79f-961262161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=25.6% Image=83.3%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_d2_v5",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
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
          "id": "azure+koreasouth+standard_e2s_v3",
          "uid": "tbl4f9lg3uc227ohnu4c",
          "cspSpecName": "Standard_E2s_v3",
          "name": "azure+koreasouth+standard_e2s_v3",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 15.625,
          "diskSizeGB": 34,
          "costPerHour": 0.144,
          "orderInFilteredResult": 4,
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
          "rootDiskSize": 34,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MaxDataDiskCount",
              "value": "4"
            },
            {
              "key": "MemoryInMB",
              "value": "16384"
            },
            {
              "key": "Name",
              "value": "Standard_E2s_v3"
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
              "value": "32768"
            },
            {
              "key": "MaxResourceVolumeMB",
              "value": "32768"
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
              "value": "16"
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
              "value": "160"
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
              "value": "53687091200"
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
              "value": "standardESv3Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "E2s_v3"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-10-30T12:46:02.137658081Z",
          "availabilityRefreshAfter": "2026-10-30T13:46:02.137658081Z"
        },
        {
          "id": "azure+koreasouth+standard_d2_v5",
          "uid": "tbvinsjtnc0dvpmikoef",
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
          "memoryGiB": 7.8125,
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "uid": "tbm5bq012h0vl4rmsqjj",
          "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.08.21 14:01:58 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
              "value": "22.04.202608060"
            },
            {
              "key": "ID",
              "value": "/Subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202608060"
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
          "systemLabel": "",
          "description": "",
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
  "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=51.2% Max=100.0% Avg=85.9% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
  "targetCloud": {
    "csp": "azure",
    "region": "koreasouth"
  },
  "targetInfra": {
    "name": "my02-infra101",
    "installMonAgent": "",
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
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
        "connectionName": "azure-koreasouth",
        "specId": "azure+koreasouth+standard_d4s_v4",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
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
        "name": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=83.3%",
        "connectionName": "azure-koreasouth",
        "specId": "azure+koreasouth+standard_f2s_v2",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
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
      "id": "azure+koreasouth+standard_d4s_v4",
      "uid": "tboubh1qij20enfs0695",
      "cspSpecName": "Standard_D4s_v4",
      "name": "azure+koreasouth+standard_d4s_v4",
      "namespace": "system",
      "connectionName": "azure-koreasouth",
      "providerName": "azure",
      "regionName": "koreasouth",
      "regionLatitude": 35.1796,
      "regionLongitude": 129.0756,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 15.625,
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
          "value": "Standard_D4s_v4"
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
          "value": "True"
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
          "key": "CachedDiskBytes",
          "value": "107374182400"
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
          "value": "standardDSv4Family"
        },
        {
          "key": "Tier",
          "value": "Standard"
        },
        {
          "key": "Size",
          "value": "D4s_v4"
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
      "uid": "tbdk2p4rel9h9a71mibl",
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
      "memoryGiB": 3.90625,
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
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "regionList": [
        "common"
      ],
      "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "uid": "tbm5bq012h0vl4rmsqjj",
      "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "azure-australiacentral",
      "infraType": "",
      "fetchedTime": "2026.08.21 14:01:58 Fri",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
          "value": "22.04.202608060"
        },
        {
          "key": "ID",
          "value": "/Subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202608060"
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
      "systemLabel": "",
      "description": "",
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
  "uid": "tbm6mfhb6nabn1k7qgvm",
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
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "cm-migrated-at": "2026-10-30T12:46:27Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my02-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my02-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbm6mfhb6nabn1k7qgvm"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbqv3mailnhp49dco6b2",
      "cspResourceName": "tbqv3mailnhp49dco6b2",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqv3mailnhp49dco6b2",
      "name": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:49",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
        "createdBy": "tbqv3mailnhp49dco6b2",
        "keypair": "tb3napbekhvj5kamg1d6",
        "publicip": "tbqv3mailnhp49dco6b2-77240-PublicIP"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.214.57.24",
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
        "memoryGiB": 3.90625,
        "costPerHour": 0.0961
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83",
      "networkInterface": "tbqv3mailnhp49dco6b2-44374-VNic",
      "securityGroupIds": [
        "my02-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEvMLJHK1Rx/pBoS2rXAGOs8OcdWuCA1rO258G7uC9JBkI3Tg4QkKITbf/2Rmhiu2ab3YBx4rqiONhlx3/YJY0s=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:q3YnsgWTzO7JKvCxgl6zv9NEUjqFFrVAVNEOEYqcizU",
        "firstUsedAt": "2026-10-30T12:46:56Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:55Z",
          "completedTime": "2026-10-30T12:47:04Z",
          "elapsedTime": 9,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my02-infra101-tb55j82oqa955as5p9ei",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:04Z",
          "completedTime": "2026-10-30T12:47:05Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbqv3mailnhp49dco6b2 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my02-ng-influxdb-back-1",
      "uid": "tbpdi0ugkueg94k19mdc",
      "cspResourceName": "tbpdi0ugkueg94k19mdc",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbpdi0ugkueg94k19mdc",
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
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:53",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "createdBy": "tbpdi0ugkueg94k19mdc",
        "keypair": "tb3napbekhvj5kamg1d6",
        "nlbBackend": "influxdb_back",
        "publicip": "tbpdi0ugkueg94k19mdc-12756-PublicIP"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.178.85",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
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
      "specId": "azure+koreasouth+standard_d4s_v4",
      "cspSpecName": "Standard_D4s_v4",
      "spec": {
        "cspSpecName": "Standard_D4s_v4",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.221
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83",
      "networkInterface": "tbpdi0ugkueg94k19mdc-30796-VNic",
      "securityGroupIds": [
        "my02-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMZmfMID1w/jtk3f5pXOXj7DJIe3rqXb9eCp6CIRvOMQWzqx+BdXNO0niFL1etcMCbHhuBsaScNgJSVAKZzQPkA=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:WTluzax9OOTKDrBdciATYKl91gtqoKkCw00TmIGlf9U",
        "firstUsedAt": "2026-10-30T12:46:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:55Z",
          "completedTime": "2026-10-30T12:47:02Z",
          "elapsedTime": 7,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my02-infra101-tb55j82oqa955as5p9ei",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:04Z",
          "completedTime": "2026-10-30T12:47:05Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbpdi0ugkueg94k19mdc 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my02-ng-influxdb-back-2",
      "uid": "tbnhrp3lhf257rnbhamg",
      "cspResourceName": "tbnhrp3lhf257rnbhamg",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbnhrp3lhf257rnbhamg",
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
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:55",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "createdBy": "tbnhrp3lhf257rnbhamg",
        "keypair": "tb3napbekhvj5kamg1d6",
        "nlbBackend": "influxdb_back",
        "publicip": "tbnhrp3lhf257rnbhamg-71487-PublicIP"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "40.89.195.136",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
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
      "specId": "azure+koreasouth+standard_d4s_v4",
      "cspSpecName": "Standard_D4s_v4",
      "spec": {
        "cspSpecName": "Standard_D4s_v4",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.221
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83",
      "networkInterface": "tbnhrp3lhf257rnbhamg-6243-VNic",
      "securityGroupIds": [
        "my02-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHLy3+9h8bCNXZS5r+9AtVPJdMzCRKqFYSiwZ5vn37pyZqja1Kr8nVFilb31f4TpJQ9xkLbk77pLHmxW3PMTKLc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:91Ko8/YfZQa/lXp9bFlwWB5fuQ/mOkmDylztvjG89Hw",
        "firstUsedAt": "2026-10-30T12:46:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:55Z",
          "completedTime": "2026-10-30T12:47:03Z",
          "elapsedTime": 8,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my02-infra101-tb55j82oqa955as5p9ei",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:04Z",
          "completedTime": "2026-10-30T12:47:04Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbnhrp3lhf257rnbhamg 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
        "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my02-ng-influxdb-back"
      ],
      "nodeIds": [
        "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my02-ng-influxdb-back-1",
        "my02-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
    }
  ],
  "newNodeList": null,
  "postCommands": [
    {
      "userName": "cb-user",
      "command": [
        "uname -a"
      ]
    }
  ],
  "postCommandResults": [
    {
      "phase": 1,
      "target": "all nodes",
      "status": "Completed",
      "results": {
        "results": [
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-influxdb-back-2",
            "nodeIp": "40.89.195.136",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbnhrp3lhf257rnbhamg 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-influxdb-back-1",
            "nodeIp": "20.200.178.85",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbpdi0ugkueg94k19mdc 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "20.214.57.24",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbqv3mailnhp49dco6b2 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandStatus": "Completed",
  "postCommandRequestId": "pc-my02-infra101-tb55j82oqa955as5p9ei"
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
- **Count**: 2

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "infra": [
    {
      "resourceType": "infra",
      "id": "my01-infra101",
      "uid": "tbmank9k49ik11qfvtl9",
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
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "cm-migrated-at": "2026-10-30T12:45:49Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my01-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my01-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbmank9k49ik11qfvtl9"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbk92qu15e8j9kjos3op",
          "cspResourceName": "tbk92qu15e8j9kjos3op",
          "cspResourceId": "i-0e9a8dff79587097e",
          "name": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:46:09",
          "label": {
            "Name": "tbk92qu15e8j9kjos3op",
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-10-30 12:46:09",
            "sys.cspResourceId": "i-0e9a8dff79587097e",
            "sys.cspResourceName": "tbk92qu15e8j9kjos3op",
            "sys.id": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tbk92qu15e8j9kjos3op",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "54.180.141.244",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.201",
          "privateDNS": "ip-10-0-1-201.ap-northeast-2.compute.internal",
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
          "imageId": "ami-012a353bb3afb92ee",
          "cspImageName": "ami-012a353bb3afb92ee",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-012a353bb3afb92ee",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731"
          },
          "vNetId": "my01-mig-vnet-01",
          "cspVNetId": "vpc-02e0dfd3cb3ff1fcb",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-09bb75a1817ae8dbb",
          "networkInterface": "eni-00017a3663e0da5ce",
          "securityGroupIds": [
            "my01-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tbq5thtqubnf05ts4une"
        },
        {
          "resourceType": "node",
          "id": "my01-ng-influxdb-back-1",
          "uid": "tb3e0r0s18nekarl2ulm",
          "cspResourceName": "tb3e0r0s18nekarl2ulm",
          "cspResourceId": "i-0dee049fb9412465b",
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
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:46:07",
          "label": {
            "Name": "tb3e0r0s18nekarl2ulm",
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-10-30 12:46:07",
            "sys.cspResourceId": "i-0dee049fb9412465b",
            "sys.cspResourceName": "tb3e0r0s18nekarl2ulm",
            "sys.id": "my01-ng-influxdb-back-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-influxdb-back",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tb3e0r0s18nekarl2ulm",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "13.125.79.237",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.221",
          "privateDNS": "ip-10-0-1-221.ap-northeast-2.compute.internal",
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
          "imageId": "ami-012a353bb3afb92ee",
          "cspImageName": "ami-012a353bb3afb92ee",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-012a353bb3afb92ee",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731"
          },
          "vNetId": "my01-mig-vnet-01",
          "cspVNetId": "vpc-02e0dfd3cb3ff1fcb",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-09bb75a1817ae8dbb",
          "networkInterface": "eni-01ef47616e5e27737",
          "securityGroupIds": [
            "my01-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tbq5thtqubnf05ts4une"
        },
        {
          "resourceType": "node",
          "id": "my01-ng-influxdb-back-2",
          "uid": "tbeu73dsbj12jnttj7iq",
          "cspResourceName": "tbeu73dsbj12jnttj7iq",
          "cspResourceId": "i-0cf1d4327934156dd",
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
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:46:11",
          "label": {
            "Name": "tbeu73dsbj12jnttj7iq",
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-10-30 12:46:11",
            "sys.cspResourceId": "i-0cf1d4327934156dd",
            "sys.cspResourceName": "tbeu73dsbj12jnttj7iq",
            "sys.id": "my01-ng-influxdb-back-2",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-influxdb-back",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tbeu73dsbj12jnttj7iq",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "3.35.54.160",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.253",
          "privateDNS": "ip-10-0-1-253.ap-northeast-2.compute.internal",
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
          "imageId": "ami-012a353bb3afb92ee",
          "cspImageName": "ami-012a353bb3afb92ee",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-012a353bb3afb92ee",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731"
          },
          "vNetId": "my01-mig-vnet-01",
          "cspVNetId": "vpc-02e0dfd3cb3ff1fcb",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-09bb75a1817ae8dbb",
          "networkInterface": "eni-07e998286a093231e",
          "securityGroupIds": [
            "my01-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tbq5thtqubnf05ts4une"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my02-infra101",
      "uid": "tbm6mfhb6nabn1k7qgvm",
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
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "cm-migrated-at": "2026-10-30T12:46:27Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my02-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbm6mfhb6nabn1k7qgvm"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbqv3mailnhp49dco6b2",
          "cspResourceName": "tbqv3mailnhp49dco6b2",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqv3mailnhp49dco6b2",
          "name": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:46:49",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "createdBy": "tbqv3mailnhp49dco6b2",
            "keypair": "tb3napbekhvj5kamg1d6",
            "publicip": "tbqv3mailnhp49dco6b2-77240-PublicIP",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-10-30 12:46:49",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqv3mailnhp49dco6b2",
            "sys.cspResourceName": "tbqv3mailnhp49dco6b2",
            "sys.id": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tbqv3mailnhp49dco6b2",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.57.24",
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
            "memoryGiB": 3.90625,
            "costPerHour": 0.0961
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83",
          "networkInterface": "tbqv3mailnhp49dco6b2-44374-VNic",
          "securityGroupIds": [
            "my02-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6"
        },
        {
          "resourceType": "node",
          "id": "my02-ng-influxdb-back-1",
          "uid": "tbpdi0ugkueg94k19mdc",
          "cspResourceName": "tbpdi0ugkueg94k19mdc",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbpdi0ugkueg94k19mdc",
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
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:46:53",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "createdBy": "tbpdi0ugkueg94k19mdc",
            "keypair": "tb3napbekhvj5kamg1d6",
            "nlbBackend": "influxdb_back",
            "publicip": "tbpdi0ugkueg94k19mdc-12756-PublicIP",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-10-30 12:46:53",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbpdi0ugkueg94k19mdc",
            "sys.cspResourceName": "tbpdi0ugkueg94k19mdc",
            "sys.id": "my02-ng-influxdb-back-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-influxdb-back",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tbpdi0ugkueg94k19mdc",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.178.85",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
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
          "specId": "azure+koreasouth+standard_d4s_v4",
          "cspSpecName": "Standard_D4s_v4",
          "spec": {
            "cspSpecName": "Standard_D4s_v4",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.221
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83",
          "networkInterface": "tbpdi0ugkueg94k19mdc-30796-VNic",
          "securityGroupIds": [
            "my02-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6"
        },
        {
          "resourceType": "node",
          "id": "my02-ng-influxdb-back-2",
          "uid": "tbnhrp3lhf257rnbhamg",
          "cspResourceName": "tbnhrp3lhf257rnbhamg",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbnhrp3lhf257rnbhamg",
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
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:46:55",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "createdBy": "tbnhrp3lhf257rnbhamg",
            "keypair": "tb3napbekhvj5kamg1d6",
            "nlbBackend": "influxdb_back",
            "publicip": "tbnhrp3lhf257rnbhamg-71487-PublicIP",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-10-30 12:46:55",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbnhrp3lhf257rnbhamg",
            "sys.cspResourceName": "tbnhrp3lhf257rnbhamg",
            "sys.id": "my02-ng-influxdb-back-2",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-influxdb-back",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tbnhrp3lhf257rnbhamg",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "40.89.195.136",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
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
          "specId": "azure+koreasouth+standard_d4s_v4",
          "cspSpecName": "Standard_D4s_v4",
          "spec": {
            "cspSpecName": "Standard_D4s_v4",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.221
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83",
          "networkInterface": "tbnhrp3lhf257rnbhamg-6243-VNic",
          "securityGroupIds": [
            "my02-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6"
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
- **IDs**: [my01-infra101 my02-infra101]

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
  "uid": "tbm6mfhb6nabn1k7qgvm",
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
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "cm-migrated-at": "2026-10-30T12:46:27Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my02-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my02-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbm6mfhb6nabn1k7qgvm"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbqv3mailnhp49dco6b2",
      "cspResourceName": "tbqv3mailnhp49dco6b2",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqv3mailnhp49dco6b2",
      "name": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:49",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
        "createdBy": "tbqv3mailnhp49dco6b2",
        "keypair": "tb3napbekhvj5kamg1d6",
        "publicip": "tbqv3mailnhp49dco6b2-77240-PublicIP"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.214.57.24",
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
        "memoryGiB": 3.90625,
        "costPerHour": 0.0961
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83",
      "networkInterface": "tbqv3mailnhp49dco6b2-44374-VNic",
      "securityGroupIds": [
        "my02-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEvMLJHK1Rx/pBoS2rXAGOs8OcdWuCA1rO258G7uC9JBkI3Tg4QkKITbf/2Rmhiu2ab3YBx4rqiONhlx3/YJY0s=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:q3YnsgWTzO7JKvCxgl6zv9NEUjqFFrVAVNEOEYqcizU",
        "firstUsedAt": "2026-10-30T12:46:56Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:55Z",
          "completedTime": "2026-10-30T12:47:04Z",
          "elapsedTime": 9,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my02-infra101-tb55j82oqa955as5p9ei",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:04Z",
          "completedTime": "2026-10-30T12:47:05Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbqv3mailnhp49dco6b2 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my02-ng-influxdb-back-1",
      "uid": "tbpdi0ugkueg94k19mdc",
      "cspResourceName": "tbpdi0ugkueg94k19mdc",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbpdi0ugkueg94k19mdc",
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
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:53",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "createdBy": "tbpdi0ugkueg94k19mdc",
        "keypair": "tb3napbekhvj5kamg1d6",
        "nlbBackend": "influxdb_back",
        "publicip": "tbpdi0ugkueg94k19mdc-12756-PublicIP"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.178.85",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
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
      "specId": "azure+koreasouth+standard_d4s_v4",
      "cspSpecName": "Standard_D4s_v4",
      "spec": {
        "cspSpecName": "Standard_D4s_v4",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.221
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83",
      "networkInterface": "tbpdi0ugkueg94k19mdc-30796-VNic",
      "securityGroupIds": [
        "my02-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMZmfMID1w/jtk3f5pXOXj7DJIe3rqXb9eCp6CIRvOMQWzqx+BdXNO0niFL1etcMCbHhuBsaScNgJSVAKZzQPkA=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:WTluzax9OOTKDrBdciATYKl91gtqoKkCw00TmIGlf9U",
        "firstUsedAt": "2026-10-30T12:46:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:55Z",
          "completedTime": "2026-10-30T12:47:02Z",
          "elapsedTime": 7,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my02-infra101-tb55j82oqa955as5p9ei",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:04Z",
          "completedTime": "2026-10-30T12:47:05Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbpdi0ugkueg94k19mdc 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my02-ng-influxdb-back-2",
      "uid": "tbnhrp3lhf257rnbhamg",
      "cspResourceName": "tbnhrp3lhf257rnbhamg",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbnhrp3lhf257rnbhamg",
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
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:55",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "createdBy": "tbnhrp3lhf257rnbhamg",
        "keypair": "tb3napbekhvj5kamg1d6",
        "nlbBackend": "influxdb_back",
        "publicip": "tbnhrp3lhf257rnbhamg-71487-PublicIP"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "40.89.195.136",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
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
      "specId": "azure+koreasouth+standard_d4s_v4",
      "cspSpecName": "Standard_D4s_v4",
      "spec": {
        "cspSpecName": "Standard_D4s_v4",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.221
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
      },
      "vNetId": "my02-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2",
      "subnetId": "my02-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83",
      "networkInterface": "tbnhrp3lhf257rnbhamg-6243-VNic",
      "securityGroupIds": [
        "my02-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHLy3+9h8bCNXZS5r+9AtVPJdMzCRKqFYSiwZ5vn37pyZqja1Kr8nVFilb31f4TpJQ9xkLbk77pLHmxW3PMTKLc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:91Ko8/YfZQa/lXp9bFlwWB5fuQ/mOkmDylztvjG89Hw",
        "firstUsedAt": "2026-10-30T12:46:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:55Z",
          "completedTime": "2026-10-30T12:47:03Z",
          "elapsedTime": 8,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my02-infra101-tb55j82oqa955as5p9ei",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:04Z",
          "completedTime": "2026-10-30T12:47:04Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbnhrp3lhf257rnbhamg 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
        "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my02-ng-influxdb-back"
      ],
      "nodeIds": [
        "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my02-ng-influxdb-back-1",
        "my02-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
    }
  ],
  "newNodeList": null,
  "postCommands": [
    {
      "userName": "cb-user",
      "command": [
        "uname -a"
      ]
    }
  ],
  "postCommandResults": [
    {
      "phase": 1,
      "target": "all nodes",
      "status": "Completed",
      "results": {
        "results": [
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-influxdb-back-2",
            "nodeIp": "40.89.195.136",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbnhrp3lhf257rnbhamg 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-influxdb-back-1",
            "nodeIp": "20.200.178.85",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbpdi0ugkueg94k19mdc 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "20.214.57.24",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbqv3mailnhp49dco6b2 6.8.0-1070-azure #78~22.04.1-Ubuntu SMP Thu Sep 24 18:46:45 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandStatus": "Completed",
  "postCommandRequestId": "pc-my02-infra101-tb55j82oqa955as5p9ei"
}
```

</details>

### Test Case 7: Remote Command Accessibility Check

#### 7.1 Test Information

- **Test Type**: SSH Connectivity Test for All VMs
- **Command Executed**: `uname -a` (to verify system information)

#### 7.2 Test Result Information

- **Status**: ✅ **SUCCESS**
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
      "cspResourceName": "tbmtschbvlir8fh87rcu",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu",
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
        "ip": "20.200.128.237",
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
          },
          {
            "key": "BackendAddressPools",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811}"
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582}"
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
          },
          {
            "key": "BackendAddressPools",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811}"
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582}"
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
      "createdTime": "2026-10-30T12:47:40Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu"
        },
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{backendAddressPools:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682,name:backend-423682,properties:{loadBalancerBackendAddresses:[{name:backend-42368210.0.1.4,properties:{ipAddress:10.0.1.4,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2}}},{name:backend-42368210.0.1.6,properties:{ipAddress:10.0.1.6,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2}}}],loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773}],provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/backendAddressPools}],frontendIPConfigurations:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811,name:frontEndIp-252811,properties:{loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773}],privateIPAllocationMethod:Dynamic,provisioningState:Succeeded,publicIPAddress:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/publicIPAddresses/tbmtschbvlir8fh87rcu}},type:Microsoft.Network/loadBalancers/frontendIPConfigurations}],inboundNatPools:[],inboundNatRules:[],loadBalancingRules:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773,name:lbrule-431773,properties:{backendAddressPool:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682},backendAddressPools:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}],backendPort:8086,disableOutboundSnat:false,enableConnectionTracking:false,enableFloatingIP:false,enableTcpReset:false,frontendIPConfiguration:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811},frontendPort:9999,idleTimeoutInMinutes:4,loadDistribution:Default,probe:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582},protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/loadBalancingRules}],outboundRules:[],probes:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582,name:probe-339582,properties:{intervalInSeconds:10,loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773}],noHealthyBackendsBehavior:AllProbedDown,numberOfProbes:3,port:8086,probeThreshold:1,protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/probes}],provisioningState:Succeeded,resourceGuid:4c1e9c1a-8f18-4030-9963-0598e42a9131}"
        },
        {
          "key": "SKU",
          "value": "{name:Standard,tier:Regional}"
        },
        {
          "key": "Tags",
          "value": "{createdAt:1793364460}"
        },
        {
          "key": "Etag",
          "value": "W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\"
        },
        {
          "key": "Name",
          "value": "tbmtschbvlir8fh87rcu"
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
    "cspResourceName": "tbmtschbvlir8fh87rcu",
    "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu",
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
      "ip": "20.200.128.237",
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
        },
        {
          "key": "BackendAddressPools",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811}"
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582}"
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
        },
        {
          "key": "BackendAddressPools",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811}"
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582}"
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
    "createdTime": "2026-10-30T12:47:40Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "ID",
        "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu"
      },
      {
        "key": "Location",
        "value": "koreasouth"
      },
      {
        "key": "Properties",
        "value": "{backendAddressPools:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682,name:backend-423682,properties:{loadBalancerBackendAddresses:[{name:backend-42368210.0.1.4,properties:{ipAddress:10.0.1.4,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2}}},{name:backend-42368210.0.1.6,properties:{ipAddress:10.0.1.6,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2}}}],loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773}],provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/backendAddressPools}],frontendIPConfigurations:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811,name:frontEndIp-252811,properties:{loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773}],privateIPAllocationMethod:Dynamic,provisioningState:Succeeded,publicIPAddress:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/publicIPAddresses/tbmtschbvlir8fh87rcu}},type:Microsoft.Network/loadBalancers/frontendIPConfigurations}],inboundNatPools:[],inboundNatRules:[],loadBalancingRules:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773,name:lbrule-431773,properties:{backendAddressPool:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682},backendAddressPools:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}],backendPort:8086,disableOutboundSnat:false,enableConnectionTracking:false,enableFloatingIP:false,enableTcpReset:false,frontendIPConfiguration:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811},frontendPort:9999,idleTimeoutInMinutes:4,loadDistribution:Default,probe:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582},protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/loadBalancingRules}],outboundRules:[],probes:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582,name:probe-339582,properties:{intervalInSeconds:10,loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773}],noHealthyBackendsBehavior:AllProbedDown,numberOfProbes:3,port:8086,probeThreshold:1,protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/probes}],provisioningState:Succeeded,resourceGuid:4c1e9c1a-8f18-4030-9963-0598e42a9131}"
      },
      {
        "key": "SKU",
        "value": "{name:Standard,tier:Regional}"
      },
      {
        "key": "Tags",
        "value": "{createdAt:1793364460}"
      },
      {
        "key": "Etag",
        "value": "W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\"
      },
      {
        "key": "Name",
        "value": "tbmtschbvlir8fh87rcu"
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
  "cspResourceName": "tbmtschbvlir8fh87rcu",
  "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu",
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
    "ip": "20.200.128.237",
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
      },
      {
        "key": "BackendAddressPools",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811}"
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582}"
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
      },
      {
        "key": "BackendAddressPools",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}"
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811}"
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582}"
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
  "createdTime": "2026-10-30T12:47:40Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "ID",
      "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu"
    },
    {
      "key": "Location",
      "value": "koreasouth"
    },
    {
      "key": "Properties",
      "value": "{backendAddressPools:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682,name:backend-423682,properties:{loadBalancerBackendAddresses:[{name:backend-42368210.0.1.4,properties:{ipAddress:10.0.1.4,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2}}},{name:backend-42368210.0.1.6,properties:{ipAddress:10.0.1.6,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2}}}],loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773}],provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/backendAddressPools}],frontendIPConfigurations:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811,name:frontEndIp-252811,properties:{loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773}],privateIPAllocationMethod:Dynamic,provisioningState:Succeeded,publicIPAddress:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/publicIPAddresses/tbmtschbvlir8fh87rcu}},type:Microsoft.Network/loadBalancers/frontendIPConfigurations}],inboundNatPools:[],inboundNatRules:[],loadBalancingRules:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773,name:lbrule-431773,properties:{backendAddressPool:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682},backendAddressPools:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/backendAddressPools/backend-423682}],backendPort:8086,disableOutboundSnat:false,enableConnectionTracking:false,enableFloatingIP:false,enableTcpReset:false,frontendIPConfiguration:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/frontendIPConfigurations/frontEndIp-252811},frontendPort:9999,idleTimeoutInMinutes:4,loadDistribution:Default,probe:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582},protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/loadBalancingRules}],outboundRules:[],probes:[{etag:W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/probes/probe-339582,name:probe-339582,properties:{intervalInSeconds:10,loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbmtschbvlir8fh87rcu/loadBalancingRules/lbrule-431773}],noHealthyBackendsBehavior:AllProbedDown,numberOfProbes:3,port:8086,probeThreshold:1,protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/probes}],provisioningState:Succeeded,resourceGuid:4c1e9c1a-8f18-4030-9963-0598e42a9131}"
    },
    {
      "key": "SKU",
      "value": "{name:Standard,tier:Regional}"
    },
    {
      "key": "Tags",
      "value": "{createdAt:1793364460}"
    },
    {
      "key": "Etag",
      "value": "W/\\bed24e02-4814-4c41-b5f1-c309a0869946\\"
    },
    {
      "key": "Name",
      "value": "tbmtschbvlir8fh87rcu"
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

### Test Case 12: Target Infrastructure Summary

#### 12.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

### Test Case 13: Migration Report

#### 13.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 13.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-10-30 12:57:28

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
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| Standard_F2s_v2 | 2 | 3.9 | - | x86_64 |  | $0.0961 | 1 |
| Standard_D4s_v4 | 4 | 15.6 | - | x86_64 |  | $0.2210 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060 | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqv3mailnhp49dco6b2 | Running | 2 vCPU, 3.9 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 20.214.57.24<br>**Private IP:** 10.0.1.5<br>**SGs:** my02-mig-sg-02<br>**SSH:** my02-mig-sshkey-01 |
| my02-ng-influxdb-back-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbpdi0ugkueg94k19mdc | Running | 4 vCPU, 15.6 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 20.200.178.85<br>**Private IP:** 10.0.1.4<br>**SGs:** my02-mig-sg-01<br>**SSH:** my02-mig-sshkey-01 |
| my02-ng-influxdb-back-2 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbnhrp3lhf257rnbhamg | Running | 4 vCPU, 15.6 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060) | **VNet:** my02-mig-vnet-01<br>**Subnet:** my02-mig-subnet-01<br>**Public IP:** 40.89.195.136<br>**Private IP:** 10.0.1.6<br>**SGs:** my02-mig-sg-01<br>**SSH:** my02-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my02-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my02-mig-vnet-01 |
| **CSP VNet ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | azure-koreasouth |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my02-mig-subnet-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb5rs0pmcvhbce8amrk2/subnets/tbll1eb5fi5hb8p1gt83 | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my02-mig-sshkey-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tb3napbekhvj5kamg1d6 |  |  |

### Security Groups

#### Security Group: my02-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my02-mig-sg-01 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbriet3sm4j9st9l3ls7 |
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
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbn462dpau0gr390pvuu |
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
| my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | Standard_F2s_v2 | $0.0961 | $69.19 |
| my02-ng-influxdb-back-1 | Standard_D4s_v4 | $0.2210 | $159.12 |
| my02-ng-influxdb-back-2 | Standard_D4s_v4 | $0.2210 | $159.12 |




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

