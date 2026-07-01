# CM-Beetle test results for NCP (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with NCP cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: 3950cc5
- imdl: unknown
- CB-Tumblebug: Unknown (Fallback to Latest)
- CB-Spider: Unknown (Fallback to Latest)
- CB-MapUI: Unknown (Fallback to Latest)
- Target CSP: NCP
- Target Region: kr
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: July 1, 2026
- Test Time: 15:30:42 KST
- Test Execution: 2026-07-01 15:30:42 KST

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

## Test result for NCP

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 1m2.755s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 6m47.619s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 21ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 3ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 16ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 9.656s | Pass |
| 7 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 1m45.881s | Pass |
| 8 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 8ms | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 15ms | Pass |
| 10 | NLB Load Balancing Verification | ✅ **PASS** | 1m37.874s | Pass |
| 11 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.222s | Pass |
| 12 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.22s | Pass |
| 13 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 58.762s | Pass |
| 14 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 2m12.186s | Pass |

**Overall Result**: 14/14 tests passed ✅

**Total Duration**: 15m30.576855769s

*Test executed on July 1, 2026 at 15:30:42 KST (2026-07-01 15:30:42 KST) using CM-Beetle automated test CLI*

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
  "desiredCsp": "ncp",
  "desiredRegion": "kr",
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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=78.3% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "ncp",
        "region": "kr"
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+s4-g3",
            "imageId": "23214590",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "ng-ec268ed7-821e-9d73-e79f-961262161624",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=60.0%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+ci2-g3",
            "imageId": "23214590",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 50,
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
        "connectionName": "ncp-kr",
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
        "connectionName": "ncp-kr",
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
          "id": "ncp+kr+s4-g3",
          "uid": "tbm9opoi4hdt1krq5t5s",
          "cspSpecName": "s4-g3",
          "name": "ncp+kr+s4-g3",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1747,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "s4-g3"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "4"
            },
            {
              "key": "MemorySize",
              "value": "17179869184"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "7875"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "150994944"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.STAND.C004.M016.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 4EA, Memory 16GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "109093105,107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318,16946033"
            }
          ]
        },
        {
          "id": "ncp+kr+ci2-g3",
          "uid": "tb3kd4r9fjj7qc3r5acu",
          "cspSpecName": "ci2-g3",
          "name": "ncp+kr+ci2-g3",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 4,
          "diskSizeGB": -1,
          "costPerHour": 0.073,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "ci2-g3"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "2"
            },
            {
              "key": "MemorySize",
              "value": "4294967296"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "5250"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "84934656"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.CPU.C002.M004.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 2EA, Memory 4GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "109093105,107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318,16946033"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "ncp",
          "cspImageName": "23214590",
          "regionList": [
            "kr"
          ],
          "id": "23214590",
          "uid": "tbuaurmepoesrrcfn1lt",
          "name": "23214590",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ncp-kr",
          "infraType": "",
          "fetchedTime": "2026.06.08 09:13:04 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)",
          "osDiskType": "Common BlockStorage 1",
          "osDiskSizeGB": 10,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ServerImageNo",
              "value": "23214590"
            },
            {
              "key": "ServerImageName",
              "value": "ubuntu-22.04-base"
            },
            {
              "key": "ServerImageDescription",
              "value": "kernel version : 5.15.0-140-generic"
            },
            {
              "key": "ServerImageProductCode",
              "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
            },
            {
              "key": "ServerImageType",
              "value": "{code:NCP,codeName:NCP 서버이미지}"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "OsCategoryType",
              "value": "{code:LINUX,codeName:LINUX}"
            },
            {
              "key": "OsType",
              "value": "{code:UBUNTU,codeName:UBUNTU}"
            },
            {
              "key": "ServerImageStatus",
              "value": "{code:CREAT,codeName:NSI CREATED state}"
            },
            {
              "key": "ServerImageOperation",
              "value": "{code:NULL,codeName:NSI NULL OP}"
            },
            {
              "key": "ServerImageStatusName",
              "value": "created"
            },
            {
              "key": "CreateDate",
              "value": "2024-03-21T18:22:55+0900"
            },
            {
              "key": "ShareStatus",
              "value": "{code:NULL,codeName:NSI Share NULL State}"
            },
            {
              "key": "BlockStorageMappingList",
              "value": "{order:0,blockStorageSnapshotInstanceNo:23214591,blockStorageSnapshotName:snapshot of ubuntu-22.04-base,blockStorageSize:10737418240,blockStorageVolumeType:{code:CB1,codeName:Common BlockStorage 1},iops:100,throughput:104857600,isEncryptedVolume:false}"
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
          "connectionName": "ncp-kr",
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
          "connectionName": "ncp-kr",
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
            "timeout": 10
          }
        }
      ]
    },
    {
      "status": "partially-matched",
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=25.0% Max=100.0% Avg=74.2% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "ncp",
        "region": "kr"
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+s4-g3a",
            "imageId": "23214590",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "ng-ec268ed7-821e-9d73-e79f-961262161624",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=25.0% Image=60.0%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+s2-g3a",
            "imageId": "23214590",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 50,
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
        "connectionName": "ncp-kr",
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
        "connectionName": "ncp-kr",
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
          "id": "ncp+kr+s4-g3a",
          "uid": "tb3nu29jb3b462tjqqf4",
          "cspSpecName": "s4-g3a",
          "name": "ncp+kr+s4-g3a",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1747,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "s4-g3a"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "4"
            },
            {
              "key": "MemorySize",
              "value": "17179869184"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "7875"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "150994944"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.AMD.STAND.C004.M016.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 4EA, Memory 16GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "109093105,107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318"
            }
          ]
        },
        {
          "id": "ncp+kr+s2-g3a",
          "uid": "tbde6sbvnj1m13s0kfl0",
          "cspSpecName": "s2-g3a",
          "name": "ncp+kr+s2-g3a",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.0848,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "s2-g3a"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "2"
            },
            {
              "key": "MemorySize",
              "value": "8589934592"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "4725"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "84934656"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.AMD.STAND.C002.M008.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 2EA, Memory 8GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "109093105,107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "ncp",
          "cspImageName": "23214590",
          "regionList": [
            "kr"
          ],
          "id": "23214590",
          "uid": "tbuaurmepoesrrcfn1lt",
          "name": "23214590",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ncp-kr",
          "infraType": "",
          "fetchedTime": "2026.06.08 09:13:04 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)",
          "osDiskType": "Common BlockStorage 1",
          "osDiskSizeGB": 10,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ServerImageNo",
              "value": "23214590"
            },
            {
              "key": "ServerImageName",
              "value": "ubuntu-22.04-base"
            },
            {
              "key": "ServerImageDescription",
              "value": "kernel version : 5.15.0-140-generic"
            },
            {
              "key": "ServerImageProductCode",
              "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
            },
            {
              "key": "ServerImageType",
              "value": "{code:NCP,codeName:NCP 서버이미지}"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "OsCategoryType",
              "value": "{code:LINUX,codeName:LINUX}"
            },
            {
              "key": "OsType",
              "value": "{code:UBUNTU,codeName:UBUNTU}"
            },
            {
              "key": "ServerImageStatus",
              "value": "{code:CREAT,codeName:NSI CREATED state}"
            },
            {
              "key": "ServerImageOperation",
              "value": "{code:NULL,codeName:NSI NULL OP}"
            },
            {
              "key": "ServerImageStatusName",
              "value": "created"
            },
            {
              "key": "CreateDate",
              "value": "2024-03-21T18:22:55+0900"
            },
            {
              "key": "ShareStatus",
              "value": "{code:NULL,codeName:NSI Share NULL State}"
            },
            {
              "key": "BlockStorageMappingList",
              "value": "{order:0,blockStorageSnapshotInstanceNo:23214591,blockStorageSnapshotName:snapshot of ubuntu-22.04-base,blockStorageSize:10737418240,blockStorageVolumeType:{code:CB1,codeName:Common BlockStorage 1},iops:100,throughput:104857600,isEncryptedVolume:false}"
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
          "connectionName": "ncp-kr",
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
          "connectionName": "ncp-kr",
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
            "timeout": 10
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
  "uid": "tbcka9ob5ngmj3o64de7",
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
    "sys.uid": "tbcka9ob5ngmj3o64de7"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb0d6j1sb7gueaallvh5",
      "cspResourceName": "tb0d6j1sb7gueaallvh5",
      "cspResourceId": "142596691",
      "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-01 06:37:37",
      "label": {
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ncp-kr",
        "sys.createdTime": "2026-07-01 06:37:37",
        "sys.cspResourceId": "142596691",
        "sys.cspResourceName": "tb0d6j1sb7gueaallvh5",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb0d6j1sb7gueaallvh5",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=60.0%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "101.79.27.179",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "SSD",
      "rootDiskSize": 50,
      "RootDeviceName": "/dev/vda",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+ci2-g3",
      "cspSpecName": "ci2-g3",
      "spec": {
        "cspSpecName": "ci2-g3",
        "vCPU": 2,
        "memoryGiB": 4,
        "costPerHour": 0.073
      },
      "imageId": "23214590",
      "cspImageName": "23214590",
      "image": {
        "resourceType": "image",
        "cspImageName": "23214590",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "141972",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "307646",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbq9b9ffgls9jobs6mc3",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-07-01T06:38:05Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:38:05Z",
          "completedTime": "2026-07-01T06:38:06Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb0d6j1sb7gueaallvh5 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "142596691"
        },
        {
          "key": "ServerName",
          "value": "tb0d6j1sb7gueaallvh5"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize",
          "value": "4294967296"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "tbq9b9ffgls9jobs6mc3"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2026-07-01T15:33:42+0900"
        },
        {
          "key": "Uptime",
          "value": "2026-07-01T15:35:11+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.CPU.C002.M004.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "141972"
        },
        {
          "key": "SubnetNo",
          "value": "307646"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5796956"
        },
        {
          "key": "InitScriptNo",
          "value": "179364"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:CPU,codeName:CPU-Intensive}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "ci2-g3"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tbrc4osvf89mr7vj8nkn",
      "cspResourceName": "tbrc4osvf89mr7vj8nkn",
      "cspResourceId": "142596698",
      "name": "my-ng-influxdb-back-1",
      "nodeGroupId": "my-ng-influxdb-back",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-01 06:38:00",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ncp-kr",
        "sys.createdTime": "2026-07-01 06:38:00",
        "sys.cspResourceId": "142596698",
        "sys.cspResourceName": "tbrc4osvf89mr7vj8nkn",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbrc4osvf89mr7vj8nkn",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "101.79.27.181",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.8",
      "privateDNS": "",
      "rootDiskType": "SSD",
      "rootDiskSize": 50,
      "RootDeviceName": "/dev/vda",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+s4-g3",
      "cspSpecName": "s4-g3",
      "spec": {
        "cspSpecName": "s4-g3",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1747
      },
      "imageId": "23214590",
      "cspImageName": "23214590",
      "image": {
        "resourceType": "image",
        "cspImageName": "23214590",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "141972",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "307646",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbq9b9ffgls9jobs6mc3",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-07-01T06:38:06Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:38:05Z",
          "completedTime": "2026-07-01T06:38:07Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbrc4osvf89mr7vj8nkn 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "142596698"
        },
        {
          "key": "ServerName",
          "value": "tbrc4osvf89mr7vj8nkn"
        },
        {
          "key": "CpuCount",
          "value": "4"
        },
        {
          "key": "MemorySize",
          "value": "17179869184"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "tbq9b9ffgls9jobs6mc3"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2026-07-01T15:33:45+0900"
        },
        {
          "key": "Uptime",
          "value": "2026-07-01T15:35:21+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.STAND.C004.M016.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "141972"
        },
        {
          "key": "SubnetNo",
          "value": "307646"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5796958"
        },
        {
          "key": "InitScriptNo",
          "value": "179366"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:STAND,codeName:Standard}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "s4-g3"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tb2oeo8u3ds034ptjavc",
      "cspResourceName": "tb2oeo8u3ds034ptjavc",
      "cspResourceId": "142596695",
      "name": "my-ng-influxdb-back-2",
      "nodeGroupId": "my-ng-influxdb-back",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-01 06:37:54",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ncp-kr",
        "sys.createdTime": "2026-07-01 06:37:54",
        "sys.cspResourceId": "142596695",
        "sys.cspResourceName": "tb2oeo8u3ds034ptjavc",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb2oeo8u3ds034ptjavc",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "101.79.27.180",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.7",
      "privateDNS": "",
      "rootDiskType": "SSD",
      "rootDiskSize": 50,
      "RootDeviceName": "/dev/vda",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+s4-g3",
      "cspSpecName": "s4-g3",
      "spec": {
        "cspSpecName": "s4-g3",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1747
      },
      "imageId": "23214590",
      "cspImageName": "23214590",
      "image": {
        "resourceType": "image",
        "cspImageName": "23214590",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "141972",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "307646",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbq9b9ffgls9jobs6mc3",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-07-01T06:38:06Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:38:05Z",
          "completedTime": "2026-07-01T06:38:07Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb2oeo8u3ds034ptjavc 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "142596695"
        },
        {
          "key": "ServerName",
          "value": "tb2oeo8u3ds034ptjavc"
        },
        {
          "key": "CpuCount",
          "value": "4"
        },
        {
          "key": "MemorySize",
          "value": "17179869184"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "tbq9b9ffgls9jobs6mc3"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2026-07-01T15:33:45+0900"
        },
        {
          "key": "Uptime",
          "value": "2026-07-01T15:35:14+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.STAND.C004.M016.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "141972"
        },
        {
          "key": "SubnetNo",
          "value": "307646"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5796957"
        },
        {
          "key": "InitScriptNo",
          "value": "179365"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:STAND,codeName:Standard}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "s4-g3"
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
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "101.79.27.179",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tb0d6j1sb7gueaallvh5 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "101.79.27.181",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbrc4osvf89mr7vj8nkn 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "101.79.27.180",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tb2oeo8u3ds034ptjavc 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "tbcka9ob5ngmj3o64de7",
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
        "sys.uid": "tbcka9ob5ngmj3o64de7"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb0d6j1sb7gueaallvh5",
          "cspResourceName": "tb0d6j1sb7gueaallvh5",
          "cspResourceId": "142596691",
          "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-07-01 06:37:37",
          "label": {
            "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2026-07-01 06:37:37",
            "sys.cspResourceId": "142596691",
            "sys.cspResourceName": "tb0d6j1sb7gueaallvh5",
            "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tb0d6j1sb7gueaallvh5",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=60.0%",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "101.79.27.179",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "SSD",
          "rootDiskSize": 50,
          "RootDeviceName": "/dev/vda",
          "connectionName": "ncp-kr",
          "connectionConfig": {
            "configName": "ncp-kr",
            "providerName": "ncp",
            "driverName": "ncp-driver-v1.0.so",
            "credentialName": "ncp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ncp-kr",
            "regionZoneInfo": {
              "assignedRegion": "KR",
              "assignedZone": "KR-1"
            },
            "regionDetail": {
              "regionId": "KR",
              "regionName": "kr",
              "description": "Korea 1",
              "location": {
                "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
                "latitude": 37.4754,
                "longitude": 126.8831
              },
              "zones": [
                "KR-1",
                "KR-2"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ncp+kr+ci2-g3",
          "cspSpecName": "ci2-g3",
          "spec": {
            "cspSpecName": "ci2-g3",
            "vCPU": 2,
            "memoryGiB": 4,
            "costPerHour": 0.073
          },
          "imageId": "23214590",
          "cspImageName": "23214590",
          "image": {
            "resourceType": "image",
            "cspImageName": "23214590",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "141972",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "307646",
          "networkInterface": "eth0",
          "securityGroupIds": [
            "my-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tbq9b9ffgls9jobs6mc3",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
            "firstUsedAt": "2026-07-01T06:38:05Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-01T06:38:05Z",
              "completedTime": "2026-07-01T06:38:06Z",
              "elapsedTime": 1,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tb0d6j1sb7gueaallvh5 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ServerInstanceNo",
              "value": "142596691"
            },
            {
              "key": "ServerName",
              "value": "tb0d6j1sb7gueaallvh5"
            },
            {
              "key": "CpuCount",
              "value": "2"
            },
            {
              "key": "MemorySize",
              "value": "4294967296"
            },
            {
              "key": "PlatformType",
              "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
            },
            {
              "key": "LoginKeyName",
              "value": "tbq9b9ffgls9jobs6mc3"
            },
            {
              "key": "ServerInstanceStatus",
              "value": "{code:RUN,codeName:서버 RUN 상태}"
            },
            {
              "key": "ServerInstanceOperation",
              "value": "{code:NULL,codeName:서버 NULL OP}"
            },
            {
              "key": "ServerInstanceStatusName",
              "value": "running"
            },
            {
              "key": "CreateDate",
              "value": "2026-07-01T15:33:42+0900"
            },
            {
              "key": "Uptime",
              "value": "2026-07-01T15:35:11+0900"
            },
            {
              "key": "ServerImageProductCode",
              "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.CPU.C002.M004.G003"
            },
            {
              "key": "IsProtectServerTermination",
              "value": "false"
            },
            {
              "key": "ZoneCode",
              "value": "KR-1"
            },
            {
              "key": "RegionCode",
              "value": "KR"
            },
            {
              "key": "VpcNo",
              "value": "141972"
            },
            {
              "key": "SubnetNo",
              "value": "307646"
            },
            {
              "key": "NetworkInterfaceNoList",
              "value": "5796956"
            },
            {
              "key": "InitScriptNo",
              "value": "179364"
            },
            {
              "key": "ServerInstanceType",
              "value": "{code:CPU,codeName:CPU-Intensive}"
            },
            {
              "key": "BaseBlockStorageDiskType",
              "value": "{code:NET,codeName:네트웍 스토리지}"
            },
            {
              "key": "BaseBlockStorageDiskDetailType",
              "value": "{code:SSD,codeName:SSD}"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "ServerImageNo",
              "value": "23214590"
            },
            {
              "key": "ServerSpecCode",
              "value": "ci2-g3"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-1",
          "uid": "tbrc4osvf89mr7vj8nkn",
          "cspResourceName": "tbrc4osvf89mr7vj8nkn",
          "cspResourceId": "142596698",
          "name": "my-ng-influxdb-back-1",
          "nodeGroupId": "my-ng-influxdb-back",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-07-01 06:38:00",
          "label": {
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2026-07-01 06:38:00",
            "sys.cspResourceId": "142596698",
            "sys.cspResourceName": "tbrc4osvf89mr7vj8nkn",
            "sys.id": "my-ng-influxdb-back-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbrc4osvf89mr7vj8nkn",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "101.79.27.181",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.8",
          "privateDNS": "",
          "rootDiskType": "SSD",
          "rootDiskSize": 50,
          "RootDeviceName": "/dev/vda",
          "connectionName": "ncp-kr",
          "connectionConfig": {
            "configName": "ncp-kr",
            "providerName": "ncp",
            "driverName": "ncp-driver-v1.0.so",
            "credentialName": "ncp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ncp-kr",
            "regionZoneInfo": {
              "assignedRegion": "KR",
              "assignedZone": "KR-1"
            },
            "regionDetail": {
              "regionId": "KR",
              "regionName": "kr",
              "description": "Korea 1",
              "location": {
                "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
                "latitude": 37.4754,
                "longitude": 126.8831
              },
              "zones": [
                "KR-1",
                "KR-2"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ncp+kr+s4-g3",
          "cspSpecName": "s4-g3",
          "spec": {
            "cspSpecName": "s4-g3",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1747
          },
          "imageId": "23214590",
          "cspImageName": "23214590",
          "image": {
            "resourceType": "image",
            "cspImageName": "23214590",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "141972",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "307646",
          "networkInterface": "eth0",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tbq9b9ffgls9jobs6mc3",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
            "firstUsedAt": "2026-07-01T06:38:06Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-01T06:38:05Z",
              "completedTime": "2026-07-01T06:38:07Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbrc4osvf89mr7vj8nkn 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ServerInstanceNo",
              "value": "142596698"
            },
            {
              "key": "ServerName",
              "value": "tbrc4osvf89mr7vj8nkn"
            },
            {
              "key": "CpuCount",
              "value": "4"
            },
            {
              "key": "MemorySize",
              "value": "17179869184"
            },
            {
              "key": "PlatformType",
              "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
            },
            {
              "key": "LoginKeyName",
              "value": "tbq9b9ffgls9jobs6mc3"
            },
            {
              "key": "ServerInstanceStatus",
              "value": "{code:RUN,codeName:서버 RUN 상태}"
            },
            {
              "key": "ServerInstanceOperation",
              "value": "{code:NULL,codeName:서버 NULL OP}"
            },
            {
              "key": "ServerInstanceStatusName",
              "value": "running"
            },
            {
              "key": "CreateDate",
              "value": "2026-07-01T15:33:45+0900"
            },
            {
              "key": "Uptime",
              "value": "2026-07-01T15:35:21+0900"
            },
            {
              "key": "ServerImageProductCode",
              "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.STAND.C004.M016.G003"
            },
            {
              "key": "IsProtectServerTermination",
              "value": "false"
            },
            {
              "key": "ZoneCode",
              "value": "KR-1"
            },
            {
              "key": "RegionCode",
              "value": "KR"
            },
            {
              "key": "VpcNo",
              "value": "141972"
            },
            {
              "key": "SubnetNo",
              "value": "307646"
            },
            {
              "key": "NetworkInterfaceNoList",
              "value": "5796958"
            },
            {
              "key": "InitScriptNo",
              "value": "179366"
            },
            {
              "key": "ServerInstanceType",
              "value": "{code:STAND,codeName:Standard}"
            },
            {
              "key": "BaseBlockStorageDiskType",
              "value": "{code:NET,codeName:네트웍 스토리지}"
            },
            {
              "key": "BaseBlockStorageDiskDetailType",
              "value": "{code:SSD,codeName:SSD}"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "ServerImageNo",
              "value": "23214590"
            },
            {
              "key": "ServerSpecCode",
              "value": "s4-g3"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-2",
          "uid": "tb2oeo8u3ds034ptjavc",
          "cspResourceName": "tb2oeo8u3ds034ptjavc",
          "cspResourceId": "142596695",
          "name": "my-ng-influxdb-back-2",
          "nodeGroupId": "my-ng-influxdb-back",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-07-01 06:37:54",
          "label": {
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2026-07-01 06:37:54",
            "sys.cspResourceId": "142596695",
            "sys.cspResourceName": "tb2oeo8u3ds034ptjavc",
            "sys.id": "my-ng-influxdb-back-2",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tb2oeo8u3ds034ptjavc",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "101.79.27.180",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.7",
          "privateDNS": "",
          "rootDiskType": "SSD",
          "rootDiskSize": 50,
          "RootDeviceName": "/dev/vda",
          "connectionName": "ncp-kr",
          "connectionConfig": {
            "configName": "ncp-kr",
            "providerName": "ncp",
            "driverName": "ncp-driver-v1.0.so",
            "credentialName": "ncp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ncp-kr",
            "regionZoneInfo": {
              "assignedRegion": "KR",
              "assignedZone": "KR-1"
            },
            "regionDetail": {
              "regionId": "KR",
              "regionName": "kr",
              "description": "Korea 1",
              "location": {
                "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
                "latitude": 37.4754,
                "longitude": 126.8831
              },
              "zones": [
                "KR-1",
                "KR-2"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ncp+kr+s4-g3",
          "cspSpecName": "s4-g3",
          "spec": {
            "cspSpecName": "s4-g3",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1747
          },
          "imageId": "23214590",
          "cspImageName": "23214590",
          "image": {
            "resourceType": "image",
            "cspImageName": "23214590",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "141972",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "307646",
          "networkInterface": "eth0",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tbq9b9ffgls9jobs6mc3",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
            "firstUsedAt": "2026-07-01T06:38:06Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-01T06:38:05Z",
              "completedTime": "2026-07-01T06:38:07Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tb2oeo8u3ds034ptjavc 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ServerInstanceNo",
              "value": "142596695"
            },
            {
              "key": "ServerName",
              "value": "tb2oeo8u3ds034ptjavc"
            },
            {
              "key": "CpuCount",
              "value": "4"
            },
            {
              "key": "MemorySize",
              "value": "17179869184"
            },
            {
              "key": "PlatformType",
              "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
            },
            {
              "key": "LoginKeyName",
              "value": "tbq9b9ffgls9jobs6mc3"
            },
            {
              "key": "ServerInstanceStatus",
              "value": "{code:RUN,codeName:서버 RUN 상태}"
            },
            {
              "key": "ServerInstanceOperation",
              "value": "{code:NULL,codeName:서버 NULL OP}"
            },
            {
              "key": "ServerInstanceStatusName",
              "value": "running"
            },
            {
              "key": "CreateDate",
              "value": "2026-07-01T15:33:45+0900"
            },
            {
              "key": "Uptime",
              "value": "2026-07-01T15:35:14+0900"
            },
            {
              "key": "ServerImageProductCode",
              "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.STAND.C004.M016.G003"
            },
            {
              "key": "IsProtectServerTermination",
              "value": "false"
            },
            {
              "key": "ZoneCode",
              "value": "KR-1"
            },
            {
              "key": "RegionCode",
              "value": "KR"
            },
            {
              "key": "VpcNo",
              "value": "141972"
            },
            {
              "key": "SubnetNo",
              "value": "307646"
            },
            {
              "key": "NetworkInterfaceNoList",
              "value": "5796957"
            },
            {
              "key": "InitScriptNo",
              "value": "179365"
            },
            {
              "key": "ServerInstanceType",
              "value": "{code:STAND,codeName:Standard}"
            },
            {
              "key": "BaseBlockStorageDiskType",
              "value": "{code:NET,codeName:네트웍 스토리지}"
            },
            {
              "key": "BaseBlockStorageDiskDetailType",
              "value": "{code:SSD,codeName:SSD}"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "ServerImageNo",
              "value": "23214590"
            },
            {
              "key": "ServerSpecCode",
              "value": "s4-g3"
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
            "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "101.79.27.179",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb0d6j1sb7gueaallvh5 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-1",
            "nodeIp": "101.79.27.181",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbrc4osvf89mr7vj8nkn 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-2",
            "nodeIp": "101.79.27.180",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb2oeo8u3ds034ptjavc 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "tbcka9ob5ngmj3o64de7",
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
    "sys.uid": "tbcka9ob5ngmj3o64de7"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb0d6j1sb7gueaallvh5",
      "cspResourceName": "tb0d6j1sb7gueaallvh5",
      "cspResourceId": "142596691",
      "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-01 06:37:37",
      "label": {
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ncp-kr",
        "sys.createdTime": "2026-07-01 06:37:37",
        "sys.cspResourceId": "142596691",
        "sys.cspResourceName": "tb0d6j1sb7gueaallvh5",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb0d6j1sb7gueaallvh5",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=60.0%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "101.79.27.179",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "SSD",
      "rootDiskSize": 50,
      "RootDeviceName": "/dev/vda",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+ci2-g3",
      "cspSpecName": "ci2-g3",
      "spec": {
        "cspSpecName": "ci2-g3",
        "vCPU": 2,
        "memoryGiB": 4,
        "costPerHour": 0.073
      },
      "imageId": "23214590",
      "cspImageName": "23214590",
      "image": {
        "resourceType": "image",
        "cspImageName": "23214590",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "141972",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "307646",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbq9b9ffgls9jobs6mc3",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-07-01T06:38:05Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:38:05Z",
          "completedTime": "2026-07-01T06:38:06Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb0d6j1sb7gueaallvh5 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "142596691"
        },
        {
          "key": "ServerName",
          "value": "tb0d6j1sb7gueaallvh5"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize",
          "value": "4294967296"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "tbq9b9ffgls9jobs6mc3"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2026-07-01T15:33:42+0900"
        },
        {
          "key": "Uptime",
          "value": "2026-07-01T15:35:11+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.CPU.C002.M004.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "141972"
        },
        {
          "key": "SubnetNo",
          "value": "307646"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5796956"
        },
        {
          "key": "InitScriptNo",
          "value": "179364"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:CPU,codeName:CPU-Intensive}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "ci2-g3"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tbrc4osvf89mr7vj8nkn",
      "cspResourceName": "tbrc4osvf89mr7vj8nkn",
      "cspResourceId": "142596698",
      "name": "my-ng-influxdb-back-1",
      "nodeGroupId": "my-ng-influxdb-back",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-01 06:38:00",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ncp-kr",
        "sys.createdTime": "2026-07-01 06:38:00",
        "sys.cspResourceId": "142596698",
        "sys.cspResourceName": "tbrc4osvf89mr7vj8nkn",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbrc4osvf89mr7vj8nkn",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "101.79.27.181",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.8",
      "privateDNS": "",
      "rootDiskType": "SSD",
      "rootDiskSize": 50,
      "RootDeviceName": "/dev/vda",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+s4-g3",
      "cspSpecName": "s4-g3",
      "spec": {
        "cspSpecName": "s4-g3",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1747
      },
      "imageId": "23214590",
      "cspImageName": "23214590",
      "image": {
        "resourceType": "image",
        "cspImageName": "23214590",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "141972",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "307646",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbq9b9ffgls9jobs6mc3",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-07-01T06:38:06Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:38:05Z",
          "completedTime": "2026-07-01T06:38:07Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbrc4osvf89mr7vj8nkn 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "142596698"
        },
        {
          "key": "ServerName",
          "value": "tbrc4osvf89mr7vj8nkn"
        },
        {
          "key": "CpuCount",
          "value": "4"
        },
        {
          "key": "MemorySize",
          "value": "17179869184"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "tbq9b9ffgls9jobs6mc3"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2026-07-01T15:33:45+0900"
        },
        {
          "key": "Uptime",
          "value": "2026-07-01T15:35:21+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.STAND.C004.M016.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "141972"
        },
        {
          "key": "SubnetNo",
          "value": "307646"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5796958"
        },
        {
          "key": "InitScriptNo",
          "value": "179366"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:STAND,codeName:Standard}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "s4-g3"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tb2oeo8u3ds034ptjavc",
      "cspResourceName": "tb2oeo8u3ds034ptjavc",
      "cspResourceId": "142596695",
      "name": "my-ng-influxdb-back-2",
      "nodeGroupId": "my-ng-influxdb-back",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-01 06:37:54",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ncp-kr",
        "sys.createdTime": "2026-07-01 06:37:54",
        "sys.cspResourceId": "142596695",
        "sys.cspResourceName": "tb2oeo8u3ds034ptjavc",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb2oeo8u3ds034ptjavc",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "101.79.27.180",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.7",
      "privateDNS": "",
      "rootDiskType": "SSD",
      "rootDiskSize": 50,
      "RootDeviceName": "/dev/vda",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+s4-g3",
      "cspSpecName": "s4-g3",
      "spec": {
        "cspSpecName": "s4-g3",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1747
      },
      "imageId": "23214590",
      "cspImageName": "23214590",
      "image": {
        "resourceType": "image",
        "cspImageName": "23214590",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "141972",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "307646",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbq9b9ffgls9jobs6mc3",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-07-01T06:38:06Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:38:05Z",
          "completedTime": "2026-07-01T06:38:07Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb2oeo8u3ds034ptjavc 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "142596695"
        },
        {
          "key": "ServerName",
          "value": "tb2oeo8u3ds034ptjavc"
        },
        {
          "key": "CpuCount",
          "value": "4"
        },
        {
          "key": "MemorySize",
          "value": "17179869184"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "tbq9b9ffgls9jobs6mc3"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2026-07-01T15:33:45+0900"
        },
        {
          "key": "Uptime",
          "value": "2026-07-01T15:35:14+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.STAND.C004.M016.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "141972"
        },
        {
          "key": "SubnetNo",
          "value": "307646"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5796957"
        },
        {
          "key": "InitScriptNo",
          "value": "179365"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:STAND,codeName:Standard}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "s4-g3"
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
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "101.79.27.179",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tb0d6j1sb7gueaallvh5 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "101.79.27.181",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbrc4osvf89mr7vj8nkn 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "101.79.27.180",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tb2oeo8u3ds034ptjavc 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "cspResourceName": "tbun6uth37c7rfbnvqga",
      "cspResourceId": "142596939",
      "name": "my-ng-influxdb-back",
      "connectionName": "ncp-kr",
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
        "ip": "101.79.27.197",
        "port": "9999",
        "keyValueList": [
          {
            "key": "LoadBalancerInstanceNo",
            "value": "142596939"
          },
          {
            "key": "LoadBalancerListenerNo",
            "value": "589089"
          },
          {
            "key": "ProtocolType",
            "value": "{code:TCP,codeName:TCP protocol}"
          },
          {
            "key": "Port",
            "value": "9999"
          },
          {
            "key": "UseHttp2",
            "value": "false"
          },
          {
            "key": "TlsMinVersionType",
            "value": "{}"
          },
          {
            "key": "LoadBalancerRuleNoList",
            "value": "660483"
          }
        ]
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my-ng-influxdb-back",
        "nodes": [
          "my-ng-influxdb-back-2",
          "my-ng-influxdb-back-1"
        ],
        "keyValueList": [
          {
            "key": "AlgorithmType",
            "value": "Maglev Hash"
          },
          {
            "key": "TargetType",
            "value": "Server (VPC)"
          }
        ]
      },
      "healthChecker": {
        "protocol": "TCP",
        "port": "8086",
        "interval": 10,
        "threshold": 3,
        "timeout": 0
      },
      "createdTime": "2026-07-01T15:38:52+09:00",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "LoadBalancerInstanceNo",
          "value": "142596939"
        },
        {
          "key": "LoadBalancerInstanceStatus",
          "value": "{code:USED,codeName:LB USED state}"
        },
        {
          "key": "LoadBalancerInstanceOperation",
          "value": "{code:NULL,codeName:LB NULL OP}"
        },
        {
          "key": "LoadBalancerInstanceStatusName",
          "value": "Running"
        },
        {
          "key": "CreateDate",
          "value": "2026-07-01T15:38:52+0900"
        },
        {
          "key": "LoadBalancerName",
          "value": "tbun6uth37c7rfbnvqga"
        },
        {
          "key": "LoadBalancerDomain",
          "value": "tbun6uth37c7rfbnvqga-142596939-56e7db94d387.kr.lb.naverncp.com"
        },
        {
          "key": "LoadBalancerIpList",
          "value": "101.79.27.197"
        },
        {
          "key": "LoadBalancerType",
          "value": "{code:NETWORK,codeName:Network Load Balancer}"
        },
        {
          "key": "LoadBalancerNetworkType",
          "value": "{code:PUBLIC,codeName:Public}"
        },
        {
          "key": "ThroughputType",
          "value": "{code:DYNAMIC,codeName:Dynamic-Sizing}"
        },
        {
          "key": "VpcNo",
          "value": "141972"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "SubnetNoList",
          "value": "307652"
        },
        {
          "key": "LoadBalancerSubnetList",
          "value": "{zoneCode:KR-1,subnetNo:307652,publicIpInstanceNo:142596957}"
        },
        {
          "key": "LoadBalancerListenerNoList",
          "value": "589089"
        },
        {
          "key": "ListenerId",
          "value": "589089"
        },
        {
          "key": "HealthCheckerId",
          "value": "5822613"
        }
      ],
      "isAutoGenerated": false,
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
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
    "cspResourceName": "tbun6uth37c7rfbnvqga",
    "cspResourceId": "142596939",
    "name": "my-ng-influxdb-back",
    "connectionName": "ncp-kr",
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
      "ip": "101.79.27.197",
      "port": "9999",
      "keyValueList": [
        {
          "key": "LoadBalancerInstanceNo",
          "value": "142596939"
        },
        {
          "key": "LoadBalancerListenerNo",
          "value": "589089"
        },
        {
          "key": "ProtocolType",
          "value": "{code:TCP,codeName:TCP protocol}"
        },
        {
          "key": "Port",
          "value": "9999"
        },
        {
          "key": "UseHttp2",
          "value": "false"
        },
        {
          "key": "TlsMinVersionType",
          "value": "{}"
        },
        {
          "key": "LoadBalancerRuleNoList",
          "value": "660483"
        }
      ]
    },
    "targetGroup": {
      "protocol": "TCP",
      "port": "8086",
      "nodeGroupId": "my-ng-influxdb-back",
      "nodes": [
        "my-ng-influxdb-back-2",
        "my-ng-influxdb-back-1"
      ],
      "keyValueList": [
        {
          "key": "AlgorithmType",
          "value": "Maglev Hash"
        },
        {
          "key": "TargetType",
          "value": "Server (VPC)"
        }
      ]
    },
    "healthChecker": {
      "protocol": "TCP",
      "port": "8086",
      "interval": 10,
      "threshold": 3,
      "timeout": 0
    },
    "createdTime": "2026-07-01T15:38:52+09:00",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "LoadBalancerInstanceNo",
        "value": "142596939"
      },
      {
        "key": "LoadBalancerInstanceStatus",
        "value": "{code:USED,codeName:LB USED state}"
      },
      {
        "key": "LoadBalancerInstanceOperation",
        "value": "{code:NULL,codeName:LB NULL OP}"
      },
      {
        "key": "LoadBalancerInstanceStatusName",
        "value": "Running"
      },
      {
        "key": "CreateDate",
        "value": "2026-07-01T15:38:52+0900"
      },
      {
        "key": "LoadBalancerName",
        "value": "tbun6uth37c7rfbnvqga"
      },
      {
        "key": "LoadBalancerDomain",
        "value": "tbun6uth37c7rfbnvqga-142596939-56e7db94d387.kr.lb.naverncp.com"
      },
      {
        "key": "LoadBalancerIpList",
        "value": "101.79.27.197"
      },
      {
        "key": "LoadBalancerType",
        "value": "{code:NETWORK,codeName:Network Load Balancer}"
      },
      {
        "key": "LoadBalancerNetworkType",
        "value": "{code:PUBLIC,codeName:Public}"
      },
      {
        "key": "ThroughputType",
        "value": "{code:DYNAMIC,codeName:Dynamic-Sizing}"
      },
      {
        "key": "VpcNo",
        "value": "141972"
      },
      {
        "key": "RegionCode",
        "value": "KR"
      },
      {
        "key": "SubnetNoList",
        "value": "307652"
      },
      {
        "key": "LoadBalancerSubnetList",
        "value": "{zoneCode:KR-1,subnetNo:307652,publicIpInstanceNo:142596957}"
      },
      {
        "key": "LoadBalancerListenerNoList",
        "value": "589089"
      },
      {
        "key": "ListenerId",
        "value": "589089"
      },
      {
        "key": "HealthCheckerId",
        "value": "5822613"
      }
    ],
    "isAutoGenerated": false,
    "location": {
      "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
      "latitude": 37.4754,
      "longitude": 126.8831
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
  "cspResourceName": "tbun6uth37c7rfbnvqga",
  "cspResourceId": "142596939",
  "name": "my-ng-influxdb-back",
  "connectionName": "ncp-kr",
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
    "ip": "101.79.27.197",
    "port": "9999",
    "keyValueList": [
      {
        "key": "LoadBalancerInstanceNo",
        "value": "142596939"
      },
      {
        "key": "LoadBalancerListenerNo",
        "value": "589089"
      },
      {
        "key": "ProtocolType",
        "value": "{code:TCP,codeName:TCP protocol}"
      },
      {
        "key": "Port",
        "value": "9999"
      },
      {
        "key": "UseHttp2",
        "value": "false"
      },
      {
        "key": "TlsMinVersionType",
        "value": "{}"
      },
      {
        "key": "LoadBalancerRuleNoList",
        "value": "660483"
      }
    ]
  },
  "targetGroup": {
    "protocol": "TCP",
    "port": "8086",
    "nodeGroupId": "my-ng-influxdb-back",
    "nodes": [
      "my-ng-influxdb-back-2",
      "my-ng-influxdb-back-1"
    ],
    "keyValueList": [
      {
        "key": "AlgorithmType",
        "value": "Maglev Hash"
      },
      {
        "key": "TargetType",
        "value": "Server (VPC)"
      }
    ]
  },
  "healthChecker": {
    "protocol": "TCP",
    "port": "8086",
    "interval": 10,
    "threshold": 3,
    "timeout": 0
  },
  "createdTime": "2026-07-01T15:38:52+09:00",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "LoadBalancerInstanceNo",
      "value": "142596939"
    },
    {
      "key": "LoadBalancerInstanceStatus",
      "value": "{code:USED,codeName:LB USED state}"
    },
    {
      "key": "LoadBalancerInstanceOperation",
      "value": "{code:NULL,codeName:LB NULL OP}"
    },
    {
      "key": "LoadBalancerInstanceStatusName",
      "value": "Running"
    },
    {
      "key": "CreateDate",
      "value": "2026-07-01T15:38:52+0900"
    },
    {
      "key": "LoadBalancerName",
      "value": "tbun6uth37c7rfbnvqga"
    },
    {
      "key": "LoadBalancerDomain",
      "value": "tbun6uth37c7rfbnvqga-142596939-56e7db94d387.kr.lb.naverncp.com"
    },
    {
      "key": "LoadBalancerIpList",
      "value": "101.79.27.197"
    },
    {
      "key": "LoadBalancerType",
      "value": "{code:NETWORK,codeName:Network Load Balancer}"
    },
    {
      "key": "LoadBalancerNetworkType",
      "value": "{code:PUBLIC,codeName:Public}"
    },
    {
      "key": "ThroughputType",
      "value": "{code:DYNAMIC,codeName:Dynamic-Sizing}"
    },
    {
      "key": "VpcNo",
      "value": "141972"
    },
    {
      "key": "RegionCode",
      "value": "KR"
    },
    {
      "key": "SubnetNoList",
      "value": "307652"
    },
    {
      "key": "LoadBalancerSubnetList",
      "value": "{zoneCode:KR-1,subnetNo:307652,publicIpInstanceNo:142596957}"
    },
    {
      "key": "LoadBalancerListenerNoList",
      "value": "589089"
    },
    {
      "key": "ListenerId",
      "value": "589089"
    },
    {
      "key": "HealthCheckerId",
      "value": "5822613"
    }
  ],
  "isAutoGenerated": false,
  "location": {
    "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
    "latitude": 37.4754,
    "longitude": 126.8831
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

**Generated At:** 2026-07-01 06:42:02

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | NCP |
| **Target Region** | KR |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| ci2-g3 | 2 | 4.0 | - | x86_64 | default | $0.0730 | 1 |
| s4-g3 | 4 | 16.0 | - | x86_64 | default | $0.1747 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| 23214590 | ubuntu-22.04-base (Hypervisor:KVM) | Ubuntu 22.04 | Linux/UNIX | x86_64 | Common BlockStorage 1 | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | 142596691 | Running | 2 vCPU, 4.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 101.79.27.179<br>**Private IP:** 10.0.1.6<br>**SGs:** my-mig-sg-02<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-1 | 142596698 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 101.79.27.181<br>**Private IP:** 10.0.1.8<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-2 | 142596695 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 101.79.27.180<br>**Private IP:** 10.0.1.7<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-vnet-01 |
| **CSP VNet ID** | 141972 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ncp-kr |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-mig-subnet-01 | 307646 | 10.0.1.0/24 | KR-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-mig-sshkey-01 | tbq9b9ffgls9jobs6mc3 | cb-user |  |

### Security Groups

#### Security Group: my-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-01 |
| **CSP Security Group ID** | 363817 |
| **VNet** | my-mig-vnet-01 |
| **Rule Count** | 9 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-02 |
| **CSP Security Group ID** | 363818 |
| **VNet** | my-mig-vnet-01 |
| **Rule Count** | 8 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.4224 |
| **Per Day** | $10.14 |
| **Per Month (30 days)** | $304.13 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| NCP | KR | 3 | $0.4224 | $304.13 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | ci2-g3 | $0.0730 | $52.56 |
| my-ng-influxdb-back-1 | s4-g3 | $0.1747 | $125.78 |
| my-ng-influxdb-back-2 | s4-g3 | $0.1747 | $125.78 |




### Test Case 12: Migration Report

#### 12.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-07-01 06:42:07*

---

## 📊 Migration Summary

**Target Cloud:** NCP

**Target Region:** KR

**Namespace:** mig01 | **Infra ID:** my-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $304.13 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | ci2-g3, s4-g3 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 17 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** 142596691<br>**Label(sourceMachineId):** ng-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my-ng-influxdb-back-1<br>**VM ID:** 142596698<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |
| 3 | **VM Name:** my-ng-influxdb-back-2<br>**VM ID:** 142596695<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** ci2-g3<br>**vCPUs:** 2<br>**Memory:** 4.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2 | my-ng-influxdb-back-1 | **Spec ID:** s4-g3<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3 | my-ng-influxdb-back-2 | **Spec ID:** s4-g3<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my-ng-influxdb-back-1 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my-ng-influxdb-back-2 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 17 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my-mig-sg-01

**CSP ID:** 363817 | **VNet:** my-mig-vnet-01 | **Rules:** 9

**Assigned VMs:**

- **VM:** my-ng-influxdb-back-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng
- **VM:** my-ng-influxdb-back-2
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 2 | inbound | ICMP |  | 0.0.0.0/0 | - | Created by system |
| 3 | inbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 4 | inbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 5 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 6 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 7 | outbound | ICMP |  | 0.0.0.0/0 | - | Created by system |
| 8 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 9 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |

### Security Group: my-mig-sg-02

**CSP ID:** 363818 | **VNet:** my-mig-vnet-01 | **Rules:** 8

**Assigned VMs:**

- **VM:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng-ec268ed7-821e-9d73-e79f

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | - | Created by system |
| 2 | inbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 3 | inbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 4 | inbound | TCP | 9999 | 0.0.0.0/0 | inbound tcp 9999 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | outbound | ICMP |  | 0.0.0.0/0 | - | Created by system |
| 7 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 8 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source node network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** my-mig-vnet-01<br>**ID:** 141972 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my-mig-subnet-01<br>**ID:** 307646 | 10.0.1.0/24 | my-mig-vnet-01 |

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
| 1 | my-mig-sshkey-01 | tbq9b9ffgls9jobs6mc3 |  | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.4224 |
| Daily | $10.14 |
| Monthly | $304.13 |
| Yearly | $3649.54 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | ci2-g3 | $52.56 | 17.3% |

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

