# CM-Beetle test results for AZURE (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: b418c24
- imdl: unknown
- CB-Tumblebug: Unknown (Fallback to Latest)
- CB-Spider: Unknown (Fallback to Latest)
- CB-MapUI: Unknown (Fallback to Latest)
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: July 13, 2026
- Test Time: 18:05:52 KST
- Test Execution: 2026-07-13 18:05:52 KST

### Scenario

1. Recommend target model for computing infra with NLB via Beetle
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
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 3.643s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m57.933s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 14ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 5ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 14ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 14.617s | Pass |
| 7 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 19.957s | Pass |
| 8 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 5ms | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 6ms | Pass |
| 10 | NLB Load Balancing Verification | ✅ **PASS** | 4m18.94s | Pass |
| 11 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.642s | Pass |
| 12 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.217s | Pass |
| 13 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 39.358s | Pass |
| 14 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m53.795s | Pass |

**Overall Result**: 14/14 tests passed ✅

**Total Duration**: 10m24.426978518s

*Test executed on July 13, 2026 at 18:05:52 KST (2026-07-13 18:05:52 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=51.2% Max=100.0% Avg=84.8% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "nlbBackend": "influxdb_back",
              "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b4as_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
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
              "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=80.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2als_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
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
        "postCommand": {
          "userName": "",
          "command": null
        },
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
          "id": "azure+koreasouth+standard_b4as_v2",
          "uid": "tbsncjt92f4d837mfelh",
          "cspSpecName": "Standard_B4as_v2",
          "name": "azure+koreasouth+standard_b4as_v2",
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
          "costPerHour": 0.173,
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
              "value": "Standard_B4as_v2"
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
              "key": "HibernationSupported",
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
              "value": "3"
            },
            {
              "key": "UltraSSDAvailable",
              "value": "False"
            },
            {
              "key": "LocationInfo_0_Location",
              "value": "KoreaSouth"
            },
            {
              "key": "Family",
              "value": "standardBasv2Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "B4as_v2"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b2als_v2",
          "uid": "tbp2modtn0bgras2hitq",
          "cspSpecName": "Standard_B2als_v2",
          "name": "azure+koreasouth+standard_b2als_v2",
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
          "costPerHour": 0.0432,
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
              "value": "4096"
            },
            {
              "key": "Name",
              "value": "Standard_B2als_v2"
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
              "key": "HibernationSupported",
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
              "value": "False"
            },
            {
              "key": "LocationInfo_0_Location",
              "value": "KoreaSouth"
            },
            {
              "key": "Family",
              "value": "standardBasv2Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "B2als_v2"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
          "uid": "tb1s10b0f3hlee1uavl0",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:09:29 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
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
              "value": "0001-com-ubuntu-server-jammy-daily"
            },
            {
              "key": "SKU",
              "value": "22_04-daily-lts-gen2"
            },
            {
              "key": "Version",
              "value": "22.04.202606270"
            },
            {
              "key": "ID",
              "value": "/Subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy-daily/Skus/22_04-daily-lts-gen2/Versions/22.04.202606270"
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
              "Ports": "22",
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
      ]
    },
    {
      "status": "partially-matched",
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=51.2% Max=100.0% Avg=84.8% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "nlbBackend": "influxdb_back",
              "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b4s_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
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
              "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=80.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2ls_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
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
        "postCommand": {
          "userName": "",
          "command": null
        },
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
          "id": "azure+koreasouth+standard_b4s_v2",
          "uid": "tb53tmibahilpks21ooj",
          "cspSpecName": "Standard_B4s_v2",
          "name": "azure+koreasouth+standard_b4s_v2",
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
          "costPerHour": 0.191,
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
              "value": "Standard_B4s_v2"
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
              "key": "HibernationSupported",
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
              "value": "3"
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
              "value": "standardBsv2Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "B4s_v2"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b2ls_v2",
          "uid": "tb8sm6o1g1ch6hjctbrc",
          "cspSpecName": "Standard_B2ls_v2",
          "name": "azure+koreasouth+standard_b2ls_v2",
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
          "costPerHour": 0.0478,
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
              "value": "4096"
            },
            {
              "key": "Name",
              "value": "Standard_B2ls_v2"
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
              "key": "HibernationSupported",
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
              "value": "standardBsv2Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "B2ls_v2"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
          "uid": "tb1s10b0f3hlee1uavl0",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:09:29 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
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
              "value": "0001-com-ubuntu-server-jammy-daily"
            },
            {
              "key": "SKU",
              "value": "22_04-daily-lts-gen2"
            },
            {
              "key": "Version",
              "value": "22.04.202606270"
            },
            {
              "key": "ID",
              "value": "/Subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy-daily/Skus/22_04-daily-lts-gen2/Versions/22.04.202606270"
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
              "Ports": "22",
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
      ]
    }
  ],
  "message": "2 candidate(s) recommended — each with 1 NLB(s) and 2 NodeGroup(s)"
}
```

</details>

### Test Case 2: Migrate the computing infra as defined in the target model

#### 2.1 API Request Information

- **API Endpoint**: `POST /beetle/migration/ns/mig01/infra`
- **Purpose**: Create and migrate infrastructure based on recommendation
- **Namespace ID**: `mig01`
- **Request Body**: Uses the response from the previous recommendation step

#### 2.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **MCI ID**: `my-infra101`
- **MCI Name**: `my-infra101`
- **Status**: `Running:3 (R:3/3)`

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my-infra101",
  "uid": "tbd2e5b9emntpa5r4cfi",
  "name": "my-infra101",
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
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbd2e5b9emntpa5r4cfi"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbcl9tmmuljpq6ch85bu",
      "cspResourceName": "tbcl9tmmuljpq6ch85bu",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu",
      "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-07-13 09:07:36",
      "label": {
        "createdBy": "tbcl9tmmuljpq6ch85bu",
        "keypair": "tb2u75q8oapec5k38b4g",
        "publicip": "tbcl9tmmuljpq6ch85bu-13542-PublicIP",
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-07-13 09:07:36",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu",
        "sys.cspResourceName": "tbcl9tmmuljpq6ch85bu",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbcl9tmmuljpq6ch85bu",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=80.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.154.27",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 30,
      "RootDeviceName": "Not visible in Azure",
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
      "specId": "azure+koreasouth+standard_b2als_v2",
      "cspSpecName": "Standard_B2als_v2",
      "spec": {
        "cspSpecName": "Standard_B2als_v2",
        "vCPU": 2,
        "memoryGiB": 3.90625,
        "costPerHour": 0.0432
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca",
      "networkInterface": "tbcl9tmmuljpq6ch85bu-43200-VNic",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOE/7RoEi+AzzZ1kpoAWio1noNE+xZEwNk+gdvVvcXYjKWg9VU63Tr3tY2NBg8t9GGLpBJy/caxls10/DUJwrfo=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:teJyjWv5Ck9PjCQTVmatU6g/pWeR7Gn/lsbqgStc1lw",
        "firstUsedAt": "2026-07-13T09:07:44Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:07:43Z",
          "completedTime": "2026-07-13T09:07:54Z",
          "elapsedTime": 11,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbcl9tmmuljpq6ch85bu 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbcl9tmmuljpq6ch85bu-43200-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbcl9tmmuljpq6ch85bu,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCs8MgIejibfA1aLUcHOjd/je1e94Qxr/VrlIaUhF0JVcjir98yzDvBS++48cVJL/OCYJAFVG3qpNOdmeAM/l7Xxwm1NsS7Yw/HSlw8YFPqySDNLVvtiB65bA5Bo3hd6sY9dK5LoyD3dHLHiN5qE0dUbT/Pj426Zt1YLcZ0STODZjhzsgaxB0gkOHV1avfQtTRsqZzrDrwM9Jduvr1du0bCqZnudFSZlCmKLx4n4cZaqLpW/u/CjTiod1b+lTSxDKAiG+bSIpliV8dfjzetr7ngy5A2JLI+3KTAqfOnUn2btPhx63Kf6/xe38DItxTa0E2PlYu70ItcB3sKEeJYtpSiGZJwouSjp5W0sr5zscsi9wS9RrAzS66KpI6Fyj+rmq0sUkgoRK75Xq6DJXyTUyZbWw5sfosSDaj457dnLyQWRP9eiYxK03WOnz0I1IC8DedCATbXHcMLY7jP5eTp1g7+srJJ4kQwyKHz/XTV94stTRoTclJHWm6pnEX0Og73ja9ldfvyrA3YoVmMJF2qjvCMe2e+El1a/m3eVs85/AUv9Ftu5Of6dyiLuHsLRvEuJiRmQIrnSED3S9hjyd9+wvRzarrgivBNNgUDq50w5U6A9orDTYbVCbOwFPNScrspBWlBagN9ddshNc+knDGSWEG/ozLaTDUkHc6To8iIiQOa2Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202607110,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202607110},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbcl9tmmuljpq6ch85bu_OsDisk_1_2c8d3f955b3946d88c381112443e49ab,storageAccountType:Premium_LRS},name:tbcl9tmmuljpq6ch85bu_OsDisk_1_2c8d3f955b3946d88c381112443e49ab,osType:Linux}},timeCreated:2026-07-13T09:06:46.7816512Z,vmId:49705c05-5f42-411d-8314-cdd21611a599}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbcl9tmmuljpq6ch85bu,keypair:tb2u75q8oapec5k38b4g,publicip:tbcl9tmmuljpq6ch85bu-13542-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu"
        },
        {
          "key": "Name",
          "value": "tbcl9tmmuljpq6ch85bu"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tb1svcl4me4e92tv6sgf",
      "cspResourceName": "tb1svcl4me4e92tv6sgf",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf",
      "name": "my-ng-influxdb-back-1",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-13 09:07:36",
      "label": {
        "createdBy": "tb1svcl4me4e92tv6sgf",
        "keypair": "tb2u75q8oapec5k38b4g",
        "nlbBackend": "influxdb_back",
        "publicip": "tb1svcl4me4e92tv6sgf-44370-PublicIP",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-07-13 09:07:36",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf",
        "sys.cspResourceName": "tb1svcl4me4e92tv6sgf",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb1svcl4me4e92tv6sgf",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.169.114",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 30,
      "RootDeviceName": "Not visible in Azure",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "spec": {
        "cspSpecName": "Standard_B4as_v2",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.173
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca",
      "networkInterface": "tb1svcl4me4e92tv6sgf-7508-VNic",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBLP62jWw5FeY05LpIfxOXldynLIedRpF9RQpp9q1vEFPWSXLpwmaukBSFk7gBSJHL82ZXly+jEOj0Tpibt5AfZs=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:CZNx3S2AWjmmxVgafBcy/Dn77CYobNX5Fxv02bE8KPw",
        "firstUsedAt": "2026-07-13T09:07:47Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:07:43Z",
          "completedTime": "2026-07-13T09:07:52Z",
          "elapsedTime": 9,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb1svcl4me4e92tv6sgf 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb1svcl4me4e92tv6sgf-7508-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb1svcl4me4e92tv6sgf,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCs8MgIejibfA1aLUcHOjd/je1e94Qxr/VrlIaUhF0JVcjir98yzDvBS++48cVJL/OCYJAFVG3qpNOdmeAM/l7Xxwm1NsS7Yw/HSlw8YFPqySDNLVvtiB65bA5Bo3hd6sY9dK5LoyD3dHLHiN5qE0dUbT/Pj426Zt1YLcZ0STODZjhzsgaxB0gkOHV1avfQtTRsqZzrDrwM9Jduvr1du0bCqZnudFSZlCmKLx4n4cZaqLpW/u/CjTiod1b+lTSxDKAiG+bSIpliV8dfjzetr7ngy5A2JLI+3KTAqfOnUn2btPhx63Kf6/xe38DItxTa0E2PlYu70ItcB3sKEeJYtpSiGZJwouSjp5W0sr5zscsi9wS9RrAzS66KpI6Fyj+rmq0sUkgoRK75Xq6DJXyTUyZbWw5sfosSDaj457dnLyQWRP9eiYxK03WOnz0I1IC8DedCATbXHcMLY7jP5eTp1g7+srJJ4kQwyKHz/XTV94stTRoTclJHWm6pnEX0Og73ja9ldfvyrA3YoVmMJF2qjvCMe2e+El1a/m3eVs85/AUv9Ftu5Of6dyiLuHsLRvEuJiRmQIrnSED3S9hjyd9+wvRzarrgivBNNgUDq50w5U6A9orDTYbVCbOwFPNScrspBWlBagN9ddshNc+knDGSWEG/ozLaTDUkHc6To8iIiQOa2Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202607110,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202607110},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb1svcl4me4e92tv6sgf_OsDisk_1_95ec66a1018647aeb23188c2721d2a71,storageAccountType:Premium_LRS},name:tb1svcl4me4e92tv6sgf_OsDisk_1_95ec66a1018647aeb23188c2721d2a71,osType:Linux}},timeCreated:2026-07-13T09:06:46.9036104Z,vmId:04b5e918-c063-4fce-ad99-a807f3243c17}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tb1svcl4me4e92tv6sgf,keypair:tb2u75q8oapec5k38b4g,publicip:tb1svcl4me4e92tv6sgf-44370-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf"
        },
        {
          "key": "Name",
          "value": "tb1svcl4me4e92tv6sgf"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tbqhrgj5prs41irtpup7",
      "cspResourceName": "tbqhrgj5prs41irtpup7",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7",
      "name": "my-ng-influxdb-back-2",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-13 09:07:36",
      "label": {
        "createdBy": "tbqhrgj5prs41irtpup7",
        "keypair": "tb2u75q8oapec5k38b4g",
        "nlbBackend": "influxdb_back",
        "publicip": "tbqhrgj5prs41irtpup7-60794-PublicIP",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-07-13 09:07:36",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7",
        "sys.cspResourceName": "tbqhrgj5prs41irtpup7",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbqhrgj5prs41irtpup7",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.176.75",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 30,
      "RootDeviceName": "Not visible in Azure",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "spec": {
        "cspSpecName": "Standard_B4as_v2",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.173
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca",
      "networkInterface": "tbqhrgj5prs41irtpup7-9598-VNic",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMLPZ5VHxZr60q+bAM02p3TJpjr7sPsNubfkl+wDCtRURWAwcdMlJbi2bEMa/qclq8cis7vkJRpiNyPw1CtvD5k=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:9ROddh5SEDiKRk8HzgR9uJ4LqV3M8k1Ki2iNYBuY2zY",
        "firstUsedAt": "2026-07-13T09:07:47Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:07:43Z",
          "completedTime": "2026-07-13T09:07:50Z",
          "elapsedTime": 7,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbqhrgj5prs41irtpup7 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbqhrgj5prs41irtpup7-9598-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbqhrgj5prs41irtpup7,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCs8MgIejibfA1aLUcHOjd/je1e94Qxr/VrlIaUhF0JVcjir98yzDvBS++48cVJL/OCYJAFVG3qpNOdmeAM/l7Xxwm1NsS7Yw/HSlw8YFPqySDNLVvtiB65bA5Bo3hd6sY9dK5LoyD3dHLHiN5qE0dUbT/Pj426Zt1YLcZ0STODZjhzsgaxB0gkOHV1avfQtTRsqZzrDrwM9Jduvr1du0bCqZnudFSZlCmKLx4n4cZaqLpW/u/CjTiod1b+lTSxDKAiG+bSIpliV8dfjzetr7ngy5A2JLI+3KTAqfOnUn2btPhx63Kf6/xe38DItxTa0E2PlYu70ItcB3sKEeJYtpSiGZJwouSjp5W0sr5zscsi9wS9RrAzS66KpI6Fyj+rmq0sUkgoRK75Xq6DJXyTUyZbWw5sfosSDaj457dnLyQWRP9eiYxK03WOnz0I1IC8DedCATbXHcMLY7jP5eTp1g7+srJJ4kQwyKHz/XTV94stTRoTclJHWm6pnEX0Og73ja9ldfvyrA3YoVmMJF2qjvCMe2e+El1a/m3eVs85/AUv9Ftu5Of6dyiLuHsLRvEuJiRmQIrnSED3S9hjyd9+wvRzarrgivBNNgUDq50w5U6A9orDTYbVCbOwFPNScrspBWlBagN9ddshNc+knDGSWEG/ozLaTDUkHc6To8iIiQOa2Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202607110,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202607110},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbqhrgj5prs41irtpup7_OsDisk_1_301e40569c7c44b6976ee64593bd5153,storageAccountType:Premium_LRS},name:tbqhrgj5prs41irtpup7_OsDisk_1_301e40569c7c44b6976ee64593bd5153,osType:Linux}},timeCreated:2026-07-13T09:06:46.4640082Z,vmId:ca4bae65-c6c5-4459-9ae3-2d550dd7da9c}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbqhrgj5prs41irtpup7,keypair:tb2u75q8oapec5k38b4g,publicip:tbqhrgj5prs41irtpup7-60794-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7"
        },
        {
          "key": "Name",
          "value": "tbqhrgj5prs41irtpup7"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    }
  ],
  "newNodeList": null,
  "postCommand": {
    "userName": "cb-user",
    "command": [
      "uname -a"
    ]
  },
  "postCommandResult": {
    "results": [
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "20.200.176.75",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbqhrgj5prs41irtpup7 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "20.200.169.114",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tb1svcl4me4e92tv6sgf 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "20.200.154.27",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbcl9tmmuljpq6ch85bu 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      }
    ]
  }
}
```

</details>

### Test Case 3: Get a list of infras

#### 3.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra`
- **Purpose**: Get a list of all migrated infrastructures
- **Namespace ID**: `mig01`

#### 3.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Count**: 1

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "infra": [
    {
      "resourceType": "infra",
      "id": "my-infra101",
      "uid": "tbd2e5b9emntpa5r4cfi",
      "name": "my-infra101",
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
        "countUndefined": 0
      },
      "targetStatus": "None",
      "targetAction": "None",
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbd2e5b9emntpa5r4cfi"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbcl9tmmuljpq6ch85bu",
          "cspResourceName": "tbcl9tmmuljpq6ch85bu",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu",
          "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-07-13 09:07:36",
          "label": {
            "createdBy": "tbcl9tmmuljpq6ch85bu",
            "keypair": "tb2u75q8oapec5k38b4g",
            "publicip": "tbcl9tmmuljpq6ch85bu-13542-PublicIP",
            "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-07-13 09:07:36",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu",
            "sys.cspResourceName": "tbcl9tmmuljpq6ch85bu",
            "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbcl9tmmuljpq6ch85bu",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=80.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.154.27",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 30,
          "RootDeviceName": "Not visible in Azure",
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
          "specId": "azure+koreasouth+standard_b2als_v2",
          "cspSpecName": "Standard_B2als_v2",
          "spec": {
            "cspSpecName": "Standard_B2als_v2",
            "vCPU": 2,
            "memoryGiB": 3.90625,
            "costPerHour": 0.0432
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca",
          "networkInterface": "tbcl9tmmuljpq6ch85bu-43200-VNic",
          "securityGroupIds": [
            "my-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOE/7RoEi+AzzZ1kpoAWio1noNE+xZEwNk+gdvVvcXYjKWg9VU63Tr3tY2NBg8t9GGLpBJy/caxls10/DUJwrfo=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:teJyjWv5Ck9PjCQTVmatU6g/pWeR7Gn/lsbqgStc1lw",
            "firstUsedAt": "2026-07-13T09:07:44Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-13T09:07:43Z",
              "completedTime": "2026-07-13T09:07:54Z",
              "elapsedTime": 11,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbcl9tmmuljpq6ch85bu 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreasouth"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbcl9tmmuljpq6ch85bu-43200-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbcl9tmmuljpq6ch85bu,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCs8MgIejibfA1aLUcHOjd/je1e94Qxr/VrlIaUhF0JVcjir98yzDvBS++48cVJL/OCYJAFVG3qpNOdmeAM/l7Xxwm1NsS7Yw/HSlw8YFPqySDNLVvtiB65bA5Bo3hd6sY9dK5LoyD3dHLHiN5qE0dUbT/Pj426Zt1YLcZ0STODZjhzsgaxB0gkOHV1avfQtTRsqZzrDrwM9Jduvr1du0bCqZnudFSZlCmKLx4n4cZaqLpW/u/CjTiod1b+lTSxDKAiG+bSIpliV8dfjzetr7ngy5A2JLI+3KTAqfOnUn2btPhx63Kf6/xe38DItxTa0E2PlYu70ItcB3sKEeJYtpSiGZJwouSjp5W0sr5zscsi9wS9RrAzS66KpI6Fyj+rmq0sUkgoRK75Xq6DJXyTUyZbWw5sfosSDaj457dnLyQWRP9eiYxK03WOnz0I1IC8DedCATbXHcMLY7jP5eTp1g7+srJJ4kQwyKHz/XTV94stTRoTclJHWm6pnEX0Og73ja9ldfvyrA3YoVmMJF2qjvCMe2e+El1a/m3eVs85/AUv9Ftu5Of6dyiLuHsLRvEuJiRmQIrnSED3S9hjyd9+wvRzarrgivBNNgUDq50w5U6A9orDTYbVCbOwFPNScrspBWlBagN9ddshNc+knDGSWEG/ozLaTDUkHc6To8iIiQOa2Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202607110,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202607110},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbcl9tmmuljpq6ch85bu_OsDisk_1_2c8d3f955b3946d88c381112443e49ab,storageAccountType:Premium_LRS},name:tbcl9tmmuljpq6ch85bu_OsDisk_1_2c8d3f955b3946d88c381112443e49ab,osType:Linux}},timeCreated:2026-07-13T09:06:46.7816512Z,vmId:49705c05-5f42-411d-8314-cdd21611a599}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbcl9tmmuljpq6ch85bu,keypair:tb2u75q8oapec5k38b4g,publicip:tbcl9tmmuljpq6ch85bu-13542-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu"
            },
            {
              "key": "Name",
              "value": "tbcl9tmmuljpq6ch85bu"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-1",
          "uid": "tb1svcl4me4e92tv6sgf",
          "cspResourceName": "tb1svcl4me4e92tv6sgf",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf",
          "name": "my-ng-influxdb-back-1",
          "nodeGroupId": "my-ng-influxdb-back",
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
          "createdTime": "2026-07-13 09:07:36",
          "label": {
            "createdBy": "tb1svcl4me4e92tv6sgf",
            "keypair": "tb2u75q8oapec5k38b4g",
            "nlbBackend": "influxdb_back",
            "publicip": "tb1svcl4me4e92tv6sgf-44370-PublicIP",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-07-13 09:07:36",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf",
            "sys.cspResourceName": "tb1svcl4me4e92tv6sgf",
            "sys.id": "my-ng-influxdb-back-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tb1svcl4me4e92tv6sgf",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.169.114",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.5",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 30,
          "RootDeviceName": "Not visible in Azure",
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
          "specId": "azure+koreasouth+standard_b4as_v2",
          "cspSpecName": "Standard_B4as_v2",
          "spec": {
            "cspSpecName": "Standard_B4as_v2",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.173
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca",
          "networkInterface": "tb1svcl4me4e92tv6sgf-7508-VNic",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBLP62jWw5FeY05LpIfxOXldynLIedRpF9RQpp9q1vEFPWSXLpwmaukBSFk7gBSJHL82ZXly+jEOj0Tpibt5AfZs=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:CZNx3S2AWjmmxVgafBcy/Dn77CYobNX5Fxv02bE8KPw",
            "firstUsedAt": "2026-07-13T09:07:47Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-13T09:07:43Z",
              "completedTime": "2026-07-13T09:07:52Z",
              "elapsedTime": 9,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tb1svcl4me4e92tv6sgf 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreasouth"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb1svcl4me4e92tv6sgf-7508-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb1svcl4me4e92tv6sgf,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCs8MgIejibfA1aLUcHOjd/je1e94Qxr/VrlIaUhF0JVcjir98yzDvBS++48cVJL/OCYJAFVG3qpNOdmeAM/l7Xxwm1NsS7Yw/HSlw8YFPqySDNLVvtiB65bA5Bo3hd6sY9dK5LoyD3dHLHiN5qE0dUbT/Pj426Zt1YLcZ0STODZjhzsgaxB0gkOHV1avfQtTRsqZzrDrwM9Jduvr1du0bCqZnudFSZlCmKLx4n4cZaqLpW/u/CjTiod1b+lTSxDKAiG+bSIpliV8dfjzetr7ngy5A2JLI+3KTAqfOnUn2btPhx63Kf6/xe38DItxTa0E2PlYu70ItcB3sKEeJYtpSiGZJwouSjp5W0sr5zscsi9wS9RrAzS66KpI6Fyj+rmq0sUkgoRK75Xq6DJXyTUyZbWw5sfosSDaj457dnLyQWRP9eiYxK03WOnz0I1IC8DedCATbXHcMLY7jP5eTp1g7+srJJ4kQwyKHz/XTV94stTRoTclJHWm6pnEX0Og73ja9ldfvyrA3YoVmMJF2qjvCMe2e+El1a/m3eVs85/AUv9Ftu5Of6dyiLuHsLRvEuJiRmQIrnSED3S9hjyd9+wvRzarrgivBNNgUDq50w5U6A9orDTYbVCbOwFPNScrspBWlBagN9ddshNc+knDGSWEG/ozLaTDUkHc6To8iIiQOa2Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202607110,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202607110},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb1svcl4me4e92tv6sgf_OsDisk_1_95ec66a1018647aeb23188c2721d2a71,storageAccountType:Premium_LRS},name:tb1svcl4me4e92tv6sgf_OsDisk_1_95ec66a1018647aeb23188c2721d2a71,osType:Linux}},timeCreated:2026-07-13T09:06:46.9036104Z,vmId:04b5e918-c063-4fce-ad99-a807f3243c17}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tb1svcl4me4e92tv6sgf,keypair:tb2u75q8oapec5k38b4g,publicip:tb1svcl4me4e92tv6sgf-44370-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf"
            },
            {
              "key": "Name",
              "value": "tb1svcl4me4e92tv6sgf"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-2",
          "uid": "tbqhrgj5prs41irtpup7",
          "cspResourceName": "tbqhrgj5prs41irtpup7",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7",
          "name": "my-ng-influxdb-back-2",
          "nodeGroupId": "my-ng-influxdb-back",
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
          "createdTime": "2026-07-13 09:07:36",
          "label": {
            "createdBy": "tbqhrgj5prs41irtpup7",
            "keypair": "tb2u75q8oapec5k38b4g",
            "nlbBackend": "influxdb_back",
            "publicip": "tbqhrgj5prs41irtpup7-60794-PublicIP",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-07-13 09:07:36",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7",
            "sys.cspResourceName": "tbqhrgj5prs41irtpup7",
            "sys.id": "my-ng-influxdb-back-2",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbqhrgj5prs41irtpup7",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.176.75",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 30,
          "RootDeviceName": "Not visible in Azure",
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
          "specId": "azure+koreasouth+standard_b4as_v2",
          "cspSpecName": "Standard_B4as_v2",
          "spec": {
            "cspSpecName": "Standard_B4as_v2",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.173
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca",
          "networkInterface": "tbqhrgj5prs41irtpup7-9598-VNic",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMLPZ5VHxZr60q+bAM02p3TJpjr7sPsNubfkl+wDCtRURWAwcdMlJbi2bEMa/qclq8cis7vkJRpiNyPw1CtvD5k=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:9ROddh5SEDiKRk8HzgR9uJ4LqV3M8k1Ki2iNYBuY2zY",
            "firstUsedAt": "2026-07-13T09:07:47Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-13T09:07:43Z",
              "completedTime": "2026-07-13T09:07:50Z",
              "elapsedTime": 7,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbqhrgj5prs41irtpup7 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreasouth"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbqhrgj5prs41irtpup7-9598-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbqhrgj5prs41irtpup7,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCs8MgIejibfA1aLUcHOjd/je1e94Qxr/VrlIaUhF0JVcjir98yzDvBS++48cVJL/OCYJAFVG3qpNOdmeAM/l7Xxwm1NsS7Yw/HSlw8YFPqySDNLVvtiB65bA5Bo3hd6sY9dK5LoyD3dHLHiN5qE0dUbT/Pj426Zt1YLcZ0STODZjhzsgaxB0gkOHV1avfQtTRsqZzrDrwM9Jduvr1du0bCqZnudFSZlCmKLx4n4cZaqLpW/u/CjTiod1b+lTSxDKAiG+bSIpliV8dfjzetr7ngy5A2JLI+3KTAqfOnUn2btPhx63Kf6/xe38DItxTa0E2PlYu70ItcB3sKEeJYtpSiGZJwouSjp5W0sr5zscsi9wS9RrAzS66KpI6Fyj+rmq0sUkgoRK75Xq6DJXyTUyZbWw5sfosSDaj457dnLyQWRP9eiYxK03WOnz0I1IC8DedCATbXHcMLY7jP5eTp1g7+srJJ4kQwyKHz/XTV94stTRoTclJHWm6pnEX0Og73ja9ldfvyrA3YoVmMJF2qjvCMe2e+El1a/m3eVs85/AUv9Ftu5Of6dyiLuHsLRvEuJiRmQIrnSED3S9hjyd9+wvRzarrgivBNNgUDq50w5U6A9orDTYbVCbOwFPNScrspBWlBagN9ddshNc+knDGSWEG/ozLaTDUkHc6To8iIiQOa2Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202607110,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202607110},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbqhrgj5prs41irtpup7_OsDisk_1_301e40569c7c44b6976ee64593bd5153,storageAccountType:Premium_LRS},name:tbqhrgj5prs41irtpup7_OsDisk_1_301e40569c7c44b6976ee64593bd5153,osType:Linux}},timeCreated:2026-07-13T09:06:46.4640082Z,vmId:ca4bae65-c6c5-4459-9ae3-2d550dd7da9c}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbqhrgj5prs41irtpup7,keypair:tb2u75q8oapec5k38b4g,publicip:tbqhrgj5prs41irtpup7-60794-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7"
            },
            {
              "key": "Name",
              "value": "tbqhrgj5prs41irtpup7"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        }
      ],
      "newNodeList": null,
      "postCommand": {
        "userName": "cb-user",
        "command": [
          "uname -a"
        ]
      },
      "postCommandResult": {
        "results": [
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-2",
            "nodeIp": "20.200.176.75",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbqhrgj5prs41irtpup7 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-1",
            "nodeIp": "20.200.169.114",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb1svcl4me4e92tv6sgf 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "20.200.154.27",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbcl9tmmuljpq6ch85bu 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          }
        ]
      }
    }
  ]
}
```

</details>

### Test Case 4: Get a list of infra IDs

#### 4.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra?option=id`
- **Purpose**: Get a list of IDs of all migrated infrastructures

#### 4.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **IDs**: [my-infra101]

### Test Case 5: Get a specific infra

#### 5.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra/{{infraId}}`
#### 5.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my-infra101",
  "uid": "tbd2e5b9emntpa5r4cfi",
  "name": "my-infra101",
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
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbd2e5b9emntpa5r4cfi"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbcl9tmmuljpq6ch85bu",
      "cspResourceName": "tbcl9tmmuljpq6ch85bu",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu",
      "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-07-13 09:07:36",
      "label": {
        "createdBy": "tbcl9tmmuljpq6ch85bu",
        "keypair": "tb2u75q8oapec5k38b4g",
        "publicip": "tbcl9tmmuljpq6ch85bu-13542-PublicIP",
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-07-13 09:07:36",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu",
        "sys.cspResourceName": "tbcl9tmmuljpq6ch85bu",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbcl9tmmuljpq6ch85bu",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=80.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.154.27",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 30,
      "RootDeviceName": "Not visible in Azure",
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
      "specId": "azure+koreasouth+standard_b2als_v2",
      "cspSpecName": "Standard_B2als_v2",
      "spec": {
        "cspSpecName": "Standard_B2als_v2",
        "vCPU": 2,
        "memoryGiB": 3.90625,
        "costPerHour": 0.0432
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca",
      "networkInterface": "tbcl9tmmuljpq6ch85bu-43200-VNic",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOE/7RoEi+AzzZ1kpoAWio1noNE+xZEwNk+gdvVvcXYjKWg9VU63Tr3tY2NBg8t9GGLpBJy/caxls10/DUJwrfo=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:teJyjWv5Ck9PjCQTVmatU6g/pWeR7Gn/lsbqgStc1lw",
        "firstUsedAt": "2026-07-13T09:07:44Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:07:43Z",
          "completedTime": "2026-07-13T09:07:54Z",
          "elapsedTime": 11,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbcl9tmmuljpq6ch85bu 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbcl9tmmuljpq6ch85bu-43200-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbcl9tmmuljpq6ch85bu,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCs8MgIejibfA1aLUcHOjd/je1e94Qxr/VrlIaUhF0JVcjir98yzDvBS++48cVJL/OCYJAFVG3qpNOdmeAM/l7Xxwm1NsS7Yw/HSlw8YFPqySDNLVvtiB65bA5Bo3hd6sY9dK5LoyD3dHLHiN5qE0dUbT/Pj426Zt1YLcZ0STODZjhzsgaxB0gkOHV1avfQtTRsqZzrDrwM9Jduvr1du0bCqZnudFSZlCmKLx4n4cZaqLpW/u/CjTiod1b+lTSxDKAiG+bSIpliV8dfjzetr7ngy5A2JLI+3KTAqfOnUn2btPhx63Kf6/xe38DItxTa0E2PlYu70ItcB3sKEeJYtpSiGZJwouSjp5W0sr5zscsi9wS9RrAzS66KpI6Fyj+rmq0sUkgoRK75Xq6DJXyTUyZbWw5sfosSDaj457dnLyQWRP9eiYxK03WOnz0I1IC8DedCATbXHcMLY7jP5eTp1g7+srJJ4kQwyKHz/XTV94stTRoTclJHWm6pnEX0Og73ja9ldfvyrA3YoVmMJF2qjvCMe2e+El1a/m3eVs85/AUv9Ftu5Of6dyiLuHsLRvEuJiRmQIrnSED3S9hjyd9+wvRzarrgivBNNgUDq50w5U6A9orDTYbVCbOwFPNScrspBWlBagN9ddshNc+knDGSWEG/ozLaTDUkHc6To8iIiQOa2Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202607110,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202607110},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbcl9tmmuljpq6ch85bu_OsDisk_1_2c8d3f955b3946d88c381112443e49ab,storageAccountType:Premium_LRS},name:tbcl9tmmuljpq6ch85bu_OsDisk_1_2c8d3f955b3946d88c381112443e49ab,osType:Linux}},timeCreated:2026-07-13T09:06:46.7816512Z,vmId:49705c05-5f42-411d-8314-cdd21611a599}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbcl9tmmuljpq6ch85bu,keypair:tb2u75q8oapec5k38b4g,publicip:tbcl9tmmuljpq6ch85bu-13542-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu"
        },
        {
          "key": "Name",
          "value": "tbcl9tmmuljpq6ch85bu"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tb1svcl4me4e92tv6sgf",
      "cspResourceName": "tb1svcl4me4e92tv6sgf",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf",
      "name": "my-ng-influxdb-back-1",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-13 09:07:36",
      "label": {
        "createdBy": "tb1svcl4me4e92tv6sgf",
        "keypair": "tb2u75q8oapec5k38b4g",
        "nlbBackend": "influxdb_back",
        "publicip": "tb1svcl4me4e92tv6sgf-44370-PublicIP",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-07-13 09:07:36",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf",
        "sys.cspResourceName": "tb1svcl4me4e92tv6sgf",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb1svcl4me4e92tv6sgf",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.169.114",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 30,
      "RootDeviceName": "Not visible in Azure",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "spec": {
        "cspSpecName": "Standard_B4as_v2",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.173
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca",
      "networkInterface": "tb1svcl4me4e92tv6sgf-7508-VNic",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBLP62jWw5FeY05LpIfxOXldynLIedRpF9RQpp9q1vEFPWSXLpwmaukBSFk7gBSJHL82ZXly+jEOj0Tpibt5AfZs=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:CZNx3S2AWjmmxVgafBcy/Dn77CYobNX5Fxv02bE8KPw",
        "firstUsedAt": "2026-07-13T09:07:47Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:07:43Z",
          "completedTime": "2026-07-13T09:07:52Z",
          "elapsedTime": 9,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb1svcl4me4e92tv6sgf 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb1svcl4me4e92tv6sgf-7508-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb1svcl4me4e92tv6sgf,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCs8MgIejibfA1aLUcHOjd/je1e94Qxr/VrlIaUhF0JVcjir98yzDvBS++48cVJL/OCYJAFVG3qpNOdmeAM/l7Xxwm1NsS7Yw/HSlw8YFPqySDNLVvtiB65bA5Bo3hd6sY9dK5LoyD3dHLHiN5qE0dUbT/Pj426Zt1YLcZ0STODZjhzsgaxB0gkOHV1avfQtTRsqZzrDrwM9Jduvr1du0bCqZnudFSZlCmKLx4n4cZaqLpW/u/CjTiod1b+lTSxDKAiG+bSIpliV8dfjzetr7ngy5A2JLI+3KTAqfOnUn2btPhx63Kf6/xe38DItxTa0E2PlYu70ItcB3sKEeJYtpSiGZJwouSjp5W0sr5zscsi9wS9RrAzS66KpI6Fyj+rmq0sUkgoRK75Xq6DJXyTUyZbWw5sfosSDaj457dnLyQWRP9eiYxK03WOnz0I1IC8DedCATbXHcMLY7jP5eTp1g7+srJJ4kQwyKHz/XTV94stTRoTclJHWm6pnEX0Og73ja9ldfvyrA3YoVmMJF2qjvCMe2e+El1a/m3eVs85/AUv9Ftu5Of6dyiLuHsLRvEuJiRmQIrnSED3S9hjyd9+wvRzarrgivBNNgUDq50w5U6A9orDTYbVCbOwFPNScrspBWlBagN9ddshNc+knDGSWEG/ozLaTDUkHc6To8iIiQOa2Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202607110,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202607110},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb1svcl4me4e92tv6sgf_OsDisk_1_95ec66a1018647aeb23188c2721d2a71,storageAccountType:Premium_LRS},name:tb1svcl4me4e92tv6sgf_OsDisk_1_95ec66a1018647aeb23188c2721d2a71,osType:Linux}},timeCreated:2026-07-13T09:06:46.9036104Z,vmId:04b5e918-c063-4fce-ad99-a807f3243c17}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tb1svcl4me4e92tv6sgf,keypair:tb2u75q8oapec5k38b4g,publicip:tb1svcl4me4e92tv6sgf-44370-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf"
        },
        {
          "key": "Name",
          "value": "tb1svcl4me4e92tv6sgf"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tbqhrgj5prs41irtpup7",
      "cspResourceName": "tbqhrgj5prs41irtpup7",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7",
      "name": "my-ng-influxdb-back-2",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-13 09:07:36",
      "label": {
        "createdBy": "tbqhrgj5prs41irtpup7",
        "keypair": "tb2u75q8oapec5k38b4g",
        "nlbBackend": "influxdb_back",
        "publicip": "tbqhrgj5prs41irtpup7-60794-PublicIP",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-07-13 09:07:36",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7",
        "sys.cspResourceName": "tbqhrgj5prs41irtpup7",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbqhrgj5prs41irtpup7",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.176.75",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 30,
      "RootDeviceName": "Not visible in Azure",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "spec": {
        "cspSpecName": "Standard_B4as_v2",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.173
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202607110",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca",
      "networkInterface": "tbqhrgj5prs41irtpup7-9598-VNic",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMLPZ5VHxZr60q+bAM02p3TJpjr7sPsNubfkl+wDCtRURWAwcdMlJbi2bEMa/qclq8cis7vkJRpiNyPw1CtvD5k=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:9ROddh5SEDiKRk8HzgR9uJ4LqV3M8k1Ki2iNYBuY2zY",
        "firstUsedAt": "2026-07-13T09:07:47Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:07:43Z",
          "completedTime": "2026-07-13T09:07:50Z",
          "elapsedTime": 7,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbqhrgj5prs41irtpup7 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbqhrgj5prs41irtpup7-9598-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbqhrgj5prs41irtpup7,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCs8MgIejibfA1aLUcHOjd/je1e94Qxr/VrlIaUhF0JVcjir98yzDvBS++48cVJL/OCYJAFVG3qpNOdmeAM/l7Xxwm1NsS7Yw/HSlw8YFPqySDNLVvtiB65bA5Bo3hd6sY9dK5LoyD3dHLHiN5qE0dUbT/Pj426Zt1YLcZ0STODZjhzsgaxB0gkOHV1avfQtTRsqZzrDrwM9Jduvr1du0bCqZnudFSZlCmKLx4n4cZaqLpW/u/CjTiod1b+lTSxDKAiG+bSIpliV8dfjzetr7ngy5A2JLI+3KTAqfOnUn2btPhx63Kf6/xe38DItxTa0E2PlYu70ItcB3sKEeJYtpSiGZJwouSjp5W0sr5zscsi9wS9RrAzS66KpI6Fyj+rmq0sUkgoRK75Xq6DJXyTUyZbWw5sfosSDaj457dnLyQWRP9eiYxK03WOnz0I1IC8DedCATbXHcMLY7jP5eTp1g7+srJJ4kQwyKHz/XTV94stTRoTclJHWm6pnEX0Og73ja9ldfvyrA3YoVmMJF2qjvCMe2e+El1a/m3eVs85/AUv9Ftu5Of6dyiLuHsLRvEuJiRmQIrnSED3S9hjyd9+wvRzarrgivBNNgUDq50w5U6A9orDTYbVCbOwFPNScrspBWlBagN9ddshNc+knDGSWEG/ozLaTDUkHc6To8iIiQOa2Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202607110,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202607110},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbqhrgj5prs41irtpup7_OsDisk_1_301e40569c7c44b6976ee64593bd5153,storageAccountType:Premium_LRS},name:tbqhrgj5prs41irtpup7_OsDisk_1_301e40569c7c44b6976ee64593bd5153,osType:Linux}},timeCreated:2026-07-13T09:06:46.4640082Z,vmId:ca4bae65-c6c5-4459-9ae3-2d550dd7da9c}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbqhrgj5prs41irtpup7,keypair:tb2u75q8oapec5k38b4g,publicip:tbqhrgj5prs41irtpup7-60794-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7"
        },
        {
          "key": "Name",
          "value": "tbqhrgj5prs41irtpup7"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    }
  ],
  "newNodeList": null,
  "postCommand": {
    "userName": "cb-user",
    "command": [
      "uname -a"
    ]
  },
  "postCommandResult": {
    "results": [
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "20.200.176.75",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbqhrgj5prs41irtpup7 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "20.200.169.114",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tb1svcl4me4e92tv6sgf 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "20.200.154.27",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbcl9tmmuljpq6ch85bu 6.8.0-1062-azure #69~22.04.1-Ubuntu SMP Mon Jun 29 19:32:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      }
    ]
  }
}
```

</details>

### Test Case 6: Remote Command Accessibility Check

#### 6.1 Test Information

- **Test Type**: SSH Connectivity Test for All VMs
- **Command Executed**: `uname -a` (to verify system information)

#### 6.2 Test Result Information

- **Status**: ✅ **SUCCESS**
**Summary**: 3/3 VMs accessible via SSH

### Test Case 7: Migrate NLBs to the cloud infra

#### 7.1 API Request Information

- **API Endpoint**: `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb`
- **Purpose**: Create target load balancers mapped from source HAProxy configuration

#### 7.2 API Response Information

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
      "id": "my-ng-influxdb-back",
      "cspResourceName": "tbleesnoogg54iu0eun1",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1",
      "name": "my-ng-influxdb-back",
      "connectionName": "azure-koreasouth",
      "connectionConfig": {
        "configName": "",
        "providerName": "",
        "driverName": "",
        "credentialName": "",
        "credentialHolder": "",
        "regionZoneInfoName": "",
        "regionZoneInfo": {
          "assignedRegion": "",
          "assignedZone": ""
        },
        "regionDetail": {
          "regionId": "",
          "regionName": "",
          "description": "",
          "location": {
            "display": "",
            "latitude": 0,
            "longitude": 0
          },
          "zones": null
        },
        "regionRepresentative": false,
        "verified": false
      },
      "type": "PUBLIC",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "ip": "20.214.56.99",
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
          },
          {
            "key": "BackendAddressPools",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592}"
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556}"
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
        "nodeGroupId": "my-ng-influxdb-back",
        "nodes": [
          "my-ng-influxdb-back-1",
          "my-ng-influxdb-back-2"
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
          },
          {
            "key": "BackendAddressPools",
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592}"
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
            "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556}"
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
      "createdTime": "2026-07-13T09:08:35Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1"
        },
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{backendAddressPools:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914,name:backend-93914,properties:{loadBalancerBackendAddresses:[{name:backend-9391410.0.1.5,properties:{ipAddress:10.0.1.5,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7}}},{name:backend-9391410.0.1.4,properties:{ipAddress:10.0.1.4,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7}}}],loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520}],provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/backendAddressPools}],frontendIPConfigurations:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592,name:frontEndIp-663592,properties:{loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520}],privateIPAllocationMethod:Dynamic,provisioningState:Succeeded,publicIPAddress:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/publicIPAddresses/tbleesnoogg54iu0eun1}},type:Microsoft.Network/loadBalancers/frontendIPConfigurations}],inboundNatPools:[],inboundNatRules:[],loadBalancingRules:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520,name:lbrule-501520,properties:{backendAddressPool:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914},backendAddressPools:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}],backendPort:8086,disableOutboundSnat:false,enableConnectionTracking:false,enableFloatingIP:false,enableTcpReset:false,frontendIPConfiguration:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592},frontendPort:9999,idleTimeoutInMinutes:4,loadDistribution:Default,probe:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556},protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/loadBalancingRules}],outboundRules:[],probes:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556,name:probe-340556,properties:{intervalInSeconds:10,loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520}],noHealthyBackendsBehavior:AllProbedDown,numberOfProbes:3,port:8086,probeThreshold:1,protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/probes}],provisioningState:Succeeded,resourceGuid:4f798827-8fa7-4e68-adda-d8eb05f4f13a}"
        },
        {
          "key": "SKU",
          "value": "{name:Standard,tier:Regional}"
        },
        {
          "key": "Tags",
          "value": "{createdAt:1783933715}"
        },
        {
          "key": "Etag",
          "value": "W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\"
        },
        {
          "key": "Name",
          "value": "tbleesnoogg54iu0eun1"
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

### Test Case 8: Get a list of migrated NLBs

#### 8.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb`

#### 8.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
[
  {
    "resourceType": "",
    "id": "my-ng-influxdb-back",
    "cspResourceName": "tbleesnoogg54iu0eun1",
    "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1",
    "name": "my-ng-influxdb-back",
    "connectionName": "azure-koreasouth",
    "connectionConfig": {
      "configName": "",
      "providerName": "",
      "driverName": "",
      "credentialName": "",
      "credentialHolder": "",
      "regionZoneInfoName": "",
      "regionZoneInfo": {
        "assignedRegion": "",
        "assignedZone": ""
      },
      "regionDetail": {
        "regionId": "",
        "regionName": "",
        "description": "",
        "location": {
          "display": "",
          "latitude": 0,
          "longitude": 0
        },
        "zones": null
      },
      "regionRepresentative": false,
      "verified": false
    },
    "type": "PUBLIC",
    "scope": "REGION",
    "listener": {
      "protocol": "TCP",
      "ip": "20.214.56.99",
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
        },
        {
          "key": "BackendAddressPools",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592}"
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556}"
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
      "nodeGroupId": "my-ng-influxdb-back",
      "nodes": [
        "my-ng-influxdb-back-1",
        "my-ng-influxdb-back-2"
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
        },
        {
          "key": "BackendAddressPools",
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592}"
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
          "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556}"
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
    "createdTime": "2026-07-13T09:08:35Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "ID",
        "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1"
      },
      {
        "key": "Location",
        "value": "koreasouth"
      },
      {
        "key": "Properties",
        "value": "{backendAddressPools:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914,name:backend-93914,properties:{loadBalancerBackendAddresses:[{name:backend-9391410.0.1.5,properties:{ipAddress:10.0.1.5,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7}}},{name:backend-9391410.0.1.4,properties:{ipAddress:10.0.1.4,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7}}}],loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520}],provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/backendAddressPools}],frontendIPConfigurations:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592,name:frontEndIp-663592,properties:{loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520}],privateIPAllocationMethod:Dynamic,provisioningState:Succeeded,publicIPAddress:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/publicIPAddresses/tbleesnoogg54iu0eun1}},type:Microsoft.Network/loadBalancers/frontendIPConfigurations}],inboundNatPools:[],inboundNatRules:[],loadBalancingRules:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520,name:lbrule-501520,properties:{backendAddressPool:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914},backendAddressPools:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}],backendPort:8086,disableOutboundSnat:false,enableConnectionTracking:false,enableFloatingIP:false,enableTcpReset:false,frontendIPConfiguration:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592},frontendPort:9999,idleTimeoutInMinutes:4,loadDistribution:Default,probe:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556},protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/loadBalancingRules}],outboundRules:[],probes:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556,name:probe-340556,properties:{intervalInSeconds:10,loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520}],noHealthyBackendsBehavior:AllProbedDown,numberOfProbes:3,port:8086,probeThreshold:1,protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/probes}],provisioningState:Succeeded,resourceGuid:4f798827-8fa7-4e68-adda-d8eb05f4f13a}"
      },
      {
        "key": "SKU",
        "value": "{name:Standard,tier:Regional}"
      },
      {
        "key": "Tags",
        "value": "{createdAt:1783933715}"
      },
      {
        "key": "Etag",
        "value": "W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\"
      },
      {
        "key": "Name",
        "value": "tbleesnoogg54iu0eun1"
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

### Test Case 9: Get details of a specific migrated NLB

#### 9.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}`

#### 9.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "",
  "id": "my-ng-influxdb-back",
  "cspResourceName": "tbleesnoogg54iu0eun1",
  "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1",
  "name": "my-ng-influxdb-back",
  "connectionName": "azure-koreasouth",
  "connectionConfig": {
    "configName": "",
    "providerName": "",
    "driverName": "",
    "credentialName": "",
    "credentialHolder": "",
    "regionZoneInfoName": "",
    "regionZoneInfo": {
      "assignedRegion": "",
      "assignedZone": ""
    },
    "regionDetail": {
      "regionId": "",
      "regionName": "",
      "description": "",
      "location": {
        "display": "",
        "latitude": 0,
        "longitude": 0
      },
      "zones": null
    },
    "regionRepresentative": false,
    "verified": false
  },
  "type": "PUBLIC",
  "scope": "REGION",
  "listener": {
    "protocol": "TCP",
    "ip": "20.214.56.99",
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
      },
      {
        "key": "BackendAddressPools",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592}"
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556}"
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
    "nodeGroupId": "my-ng-influxdb-back",
    "nodes": [
      "my-ng-influxdb-back-1",
      "my-ng-influxdb-back-2"
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
      },
      {
        "key": "BackendAddressPools",
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}"
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592}"
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
        "value": "{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556}"
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
  "createdTime": "2026-07-13T09:08:35Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "ID",
      "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1"
    },
    {
      "key": "Location",
      "value": "koreasouth"
    },
    {
      "key": "Properties",
      "value": "{backendAddressPools:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914,name:backend-93914,properties:{loadBalancerBackendAddresses:[{name:backend-9391410.0.1.5,properties:{ipAddress:10.0.1.5,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7}}},{name:backend-9391410.0.1.4,properties:{ipAddress:10.0.1.4,virtualNetwork:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7}}}],loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520}],provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/backendAddressPools}],frontendIPConfigurations:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592,name:frontEndIp-663592,properties:{loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520}],privateIPAllocationMethod:Dynamic,provisioningState:Succeeded,publicIPAddress:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/publicIPAddresses/tbleesnoogg54iu0eun1}},type:Microsoft.Network/loadBalancers/frontendIPConfigurations}],inboundNatPools:[],inboundNatRules:[],loadBalancingRules:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520,name:lbrule-501520,properties:{backendAddressPool:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914},backendAddressPools:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/backendAddressPools/backend-93914}],backendPort:8086,disableOutboundSnat:false,enableConnectionTracking:false,enableFloatingIP:false,enableTcpReset:false,frontendIPConfiguration:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/frontendIPConfigurations/frontEndIp-663592},frontendPort:9999,idleTimeoutInMinutes:4,loadDistribution:Default,probe:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556},protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/loadBalancingRules}],outboundRules:[],probes:[{etag:W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\,id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/probes/probe-340556,name:probe-340556,properties:{intervalInSeconds:10,loadBalancingRules:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/loadBalancers/tbleesnoogg54iu0eun1/loadBalancingRules/lbrule-501520}],noHealthyBackendsBehavior:AllProbedDown,numberOfProbes:3,port:8086,probeThreshold:1,protocol:Tcp,provisioningState:Succeeded},type:Microsoft.Network/loadBalancers/probes}],provisioningState:Succeeded,resourceGuid:4f798827-8fa7-4e68-adda-d8eb05f4f13a}"
    },
    {
      "key": "SKU",
      "value": "{name:Standard,tier:Regional}"
    },
    {
      "key": "Tags",
      "value": "{createdAt:1783933715}"
    },
    {
      "key": "Etag",
      "value": "W/\\dc303697-bf87-4cbe-9264-cc321951eaf5\\"
    },
    {
      "key": "Name",
      "value": "tbleesnoogg54iu0eun1"
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

### Test Case 10: NLB Load Balancing Verification

#### 10.1 Test Information

- **Test Type**: Active Traffic Distribution Verification via NLB Endpoint
- **Target Port**: `8086` (Backend Mock Web Server)
- **Listener Port**: `9999` (NLB Listener)
- **Requests Sent**: 15 HTTP GET requests

#### 10.2 Test Result Information

- **Status**: ✅ **SUCCESS**

**Summary**: Load balancing confirmed: traffic distributed across 2 unique VMs

### Test Case 11: Target Infrastructure Summary

#### 11.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`

#### 11.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Target Infrastructure Summary**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-07-13 09:13:01

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
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
| Standard_B2als_v2 | 2 | 3.9 | - | x86_64 |  | $0.0432 | 1 |
| Standard_B4as_v2 | 4 | 15.6 | - | x86_64 |  | $0.1730 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270 | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu | Running | 2 vCPU, 3.9 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 20.200.154.27<br>**Private IP:** 10.0.1.6<br>**SGs:** my-mig-sg-02<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-1 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf | Running | 4 vCPU, 15.6 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 20.200.169.114<br>**Private IP:** 10.0.1.5<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-2 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7 | Running | 4 vCPU, 15.6 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 20.200.176.75<br>**Private IP:** 10.0.1.4<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-vnet-01 |
| **CSP VNet ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | azure-koreasouth |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-mig-subnet-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-mig-sshkey-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g |  |  |

### Security Groups

#### Security Group: my-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-01 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tb1s1dcuc20roc78vo4h |
| **VNet** | my-mig-vnet-01 |
| **Rule Count** | 5 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-02 |
| **CSP Security Group ID** | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tb4fd9iclc3sk70i3qgb |
| **VNet** | my-mig-vnet-01 |
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
| **Per Hour** | $0.3892 |
| **Per Day** | $9.34 |
| **Per Month (30 days)** | $280.22 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AZURE | koreasouth | 3 | $0.3892 | $280.22 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | Standard_B2als_v2 | $0.0432 | $31.10 |
| my-ng-influxdb-back-1 | Standard_B4as_v2 | $0.1730 | $124.56 |
| my-ng-influxdb-back-2 | Standard_B4as_v2 | $0.1730 | $124.56 |




### Test Case 12: Migration Report

#### 12.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-07-13 09:13:06*

---

## 📊 Migration Summary

**Target Cloud:** AZURE

**Target Region:** koreasouth

**Namespace:** mig01 | **Infra ID:** my-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $280.22 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | Standard_B4as_v2, Standard_B2als_v2 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 9 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbcl9tmmuljpq6ch85bu<br>**Label(sourceMachineId):** ng-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my-ng-influxdb-back-1<br>**VM ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb1svcl4me4e92tv6sgf<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |
| 3 | **VM Name:** my-ng-influxdb-back-2<br>**VM ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqhrgj5prs41irtpup7<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** Standard_B2als_v2<br>**vCPUs:** 2<br>**Memory:** 3.9 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2 | my-ng-influxdb-back-1 | **Spec ID:** Standard_B4as_v2<br>**vCPUs:** 4<br>**Memory:** 15.6 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3 | my-ng-influxdb-back-2 | **Spec ID:** Standard_B4as_v2<br>**vCPUs:** 4<br>**Memory:** 15.6 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270 | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my-ng-influxdb-back-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270 | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my-ng-influxdb-back-2 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606270 | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 9 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my-mig-sg-01

**CSP ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tb1s1dcuc20roc78vo4h | **VNet:** my-mig-vnet-01 | **Rules:** 5

**Assigned VMs:**

- **VM:** my-ng-influxdb-back-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng
- **VM:** my-ng-influxdb-back-2
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 2 | inbound | ALL |  | 10.0.0.0/16 | inbound * * from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 4 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 5 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my-mig-sg-02

**CSP ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tb4fd9iclc3sk70i3qgb | **VNet:** my-mig-vnet-01 | **Rules:** 4

**Assigned VMs:**

- **VM:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng-ec268ed7-821e-9d73-e79f

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
| 1 | **Name:** my-mig-vnet-01<br>**ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my-mig-subnet-01<br>**ID:** /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb2vomneagjpv1nbslv7/subnets/tbhekuk5plt1l0batrca | 10.0.1.0/24 | my-mig-vnet-01 |

### Source Network Information

**CIDR:** 10.0.1.0/24 | **Gateway:** 10.0.1.1 | **Connected Servers:** 3

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

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my-mig-sshkey-01 | /subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tb2u75q8oapec5k38b4g |  | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.3892 |
| Daily | $9.34 |
| Monthly | $280.22 |
| Yearly | $3362.69 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | Standard_B2als_v2 | $31.10 | 11.1% |

---


---

*Report generated by CM-Beetle*


### Test Case 13: Delete the migrated NLBs

#### 13.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}`

#### 13.2 API Response Information

- **Status**: ✅ **SUCCESS**
### Test Case 14: Delete the migrated computing infra

#### 14.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}?option=terminate`

#### 14.2 API Response Information

- **Status**: ✅ **SUCCESS**
**Response Body**:

```json
{
  "message": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: my-infra101)",
  "success": true
}
```

