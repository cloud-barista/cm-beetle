# CM-Beetle test results for NCP (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with NCP cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.2+ (83efe14)
- imdl: v0.1.19+ (83efe14)
- CB-Tumblebug: v0.13.8
- CB-Spider: v0.13.20
- CB-MapUI: v0.13.14
- Target CSP: NCP
- Target Region: kr
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: October 30, 2026
- Test Time: 21:48:24 KST
- Test Execution: 2026-10-30 21:48:24 KST

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

## Test result for NCP

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 765ms | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 153ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 3m28.929s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 28ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 4ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 10ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 3.137s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 1m36.314s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 5ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 5ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 1m37.063s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.083s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.081s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 1m1.44s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m51.566s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 10m39.893093338s

*Test executed on October 30, 2026 at 21:48:24 KST (2026-10-30 21:48:24 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=80.6% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
              "nlbBackend": "influxdb_back"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=66.7%",
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
          "uid": "tb4im0u7ac4tdo5hk3gi",
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
          "orderInFilteredResult": 10,
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
              "key": "ServerSpecNo",
              "value": "242"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "143396791,109093105,107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-10-30T12:48:30.225151936Z",
          "availabilityRefreshAfter": "2026-10-30T13:48:30.225151936Z"
        },
        {
          "id": "ncp+kr+ci2-g3",
          "uid": "tblkiitfba2b02bk7bar",
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
          "orderInFilteredResult": 5,
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
              "key": "ServerSpecNo",
              "value": "1161"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "143396791,109093105,107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-10-30T12:48:30.374314944Z",
          "availabilityRefreshAfter": "2026-10-30T13:48:30.374314944Z"
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
          "uid": "tbhn395o7dg4d7h5mnqm",
          "name": "23214590",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ncp-kr",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:29 Fri",
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
              "value": "kernel version : 5.15.0-181-generic"
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
            },
            {
              "key": "ServerImageLifeCyclePhase",
              "value": "{code:IN_SERVICE,codeName:In Service}"
            },
            {
              "key": "PlatformCategoryCode",
              "value": "OS"
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=25.0% Max=100.0% Avg=68.1% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
              "nlbBackend": "influxdb_back"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=50.0% Memory=100.0% Image=66.7%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+m2-g3",
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
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=25.0% Image=66.7%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+s2-g3",
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
          "id": "ncp+kr+m2-g3",
          "uid": "tbvvvrn25r61f88391jo",
          "cspSpecName": "m2-g3",
          "name": "ncp+kr+m2-g3",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1172,
          "orderInFilteredResult": 5,
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
              "value": "m2-g3"
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
              "value": "SVR.VSVR.HIMEM.C002.M016.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 2EA, Memory 16GB"
            },
            {
              "key": "ServerSpecNo",
              "value": "237"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "143396791,109093105,107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-10-30T12:48:30.210995266Z",
          "availabilityRefreshAfter": "2026-10-30T13:48:30.210995266Z"
        },
        {
          "id": "ncp+kr+s2-g3",
          "uid": "tba9k28sh7kfolvo4fi8",
          "cspSpecName": "s2-g3",
          "name": "ncp+kr+s2-g3",
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
          "orderInFilteredResult": 9,
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
              "value": "s2-g3"
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
              "value": "SVR.VSVR.STAND.C002.M008.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 2EA, Memory 8GB"
            },
            {
              "key": "ServerSpecNo",
              "value": "239"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "143396791,109093105,107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-10-30T12:48:30.405826804Z",
          "availabilityRefreshAfter": "2026-10-30T13:48:30.405826804Z"
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
          "uid": "tbhn395o7dg4d7h5mnqm",
          "name": "23214590",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ncp-kr",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:29 Fri",
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
              "value": "kernel version : 5.15.0-181-generic"
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
            },
            {
              "key": "ServerImageLifeCyclePhase",
              "value": "{code:IN_SERVICE,codeName:In Service}"
            },
            {
              "key": "PlatformCategoryCode",
              "value": "OS"
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
  "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=80.6% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
  "targetCloud": {
    "csp": "ncp",
    "region": "kr"
  },
  "targetInfra": {
    "name": "my08-infra101",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "NLB-aware recommended infrastructure for cloud migration",
    "nodeGroups": [
      {
        "name": "my08-ng-influxdb-back",
        "nodeGroupSize": 2,
        "label": {
          "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "nlbBackend": "influxdb_back"
        },
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
        "connectionName": "ncp-kr",
        "specId": "ncp+kr+s4-g3",
        "imageId": "23214590",
        "vNetId": "my08-mig-vnet-01",
        "subnetId": "my08-mig-subnet-01",
        "securityGroupIds": [
          "my08-mig-sg-01"
        ],
        "sshKeyId": "my08-mig-sshkey-01",
        "rootDiskSize": 50,
        "dataDiskIds": null
      },
      {
        "name": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=66.7%",
        "connectionName": "ncp-kr",
        "specId": "ncp+kr+ci2-g3",
        "imageId": "23214590",
        "vNetId": "my08-mig-vnet-01",
        "subnetId": "my08-mig-subnet-01",
        "securityGroupIds": [
          "my08-mig-sg-02"
        ],
        "sshKeyId": "my08-mig-sshkey-01",
        "rootDiskSize": 50,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "my08-mig-vnet-01",
    "connectionName": "ncp-kr",
    "cidrBlock": "10.0.0.0/21",
    "subnetInfoList": [
      {
        "name": "my08-mig-subnet-01",
        "ipv4_CIDR": "10.0.1.0/24",
        "description": "a recommended subnet for migration"
      }
    ],
    "description": "a recommended vNet for migration"
  },
  "targetSshKey": {
    "name": "my08-mig-sshkey-01",
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
      "uid": "tb4im0u7ac4tdo5hk3gi",
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
      "orderInFilteredResult": 10,
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
          "key": "ServerSpecNo",
          "value": "242"
        },
        {
          "key": "CorrespondingImageIds",
          "value": "143396791,109093105,107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675"
        }
      ],
      "availabilityStatus": "Available",
      "availabilityCheckedAt": "2026-10-30T12:48:30.225151936Z",
      "availabilityRefreshAfter": "2026-10-30T13:48:30.225151936Z"
    },
    {
      "id": "ncp+kr+ci2-g3",
      "uid": "tblkiitfba2b02bk7bar",
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
      "orderInFilteredResult": 5,
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
          "key": "ServerSpecNo",
          "value": "1161"
        },
        {
          "key": "CorrespondingImageIds",
          "value": "143396791,109093105,107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675"
        }
      ],
      "availabilityStatus": "Available",
      "availabilityCheckedAt": "2026-10-30T12:48:30.374314944Z",
      "availabilityRefreshAfter": "2026-10-30T13:48:30.374314944Z"
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
      "uid": "tbhn395o7dg4d7h5mnqm",
      "name": "23214590",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "ncp-kr",
      "infraType": "",
      "fetchedTime": "2026.08.21 13:57:29 Fri",
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
          "value": "kernel version : 5.15.0-181-generic"
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
        },
        {
          "key": "ServerImageLifeCyclePhase",
          "value": "{code:IN_SERVICE,codeName:In Service}"
        },
        {
          "key": "PlatformCategoryCode",
          "value": "OS"
        }
      ],
      "systemLabel": "",
      "description": "",
      "commandHistory": null
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "my08-mig-sg-01",
      "connectionName": "ncp-kr",
      "vNetId": "my08-mig-vnet-01",
      "description": "Recommended security group for NLB backend influxdb_back",
      "firewallRules": [
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
      "name": "my08-mig-sg-02",
      "connectionName": "ncp-kr",
      "vNetId": "my08-mig-vnet-01",
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
- **MCI ID**: `my08-infra101`
- **MCI Name**: `my08-infra101`
- **Status**: `Running:3 (R:3/3)`

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my08-infra101",
  "uid": "tbqisg04kagcpshkp314",
  "name": "my08-infra101",
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
    "cm-migrated-at": "2026-10-30T12:50:18Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my08-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my08-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbqisg04kagcpshkp314"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb666sqqpo5nodvufp1h",
      "cspResourceName": "tb666sqqpo5nodvufp1h",
      "cspResourceId": "146612733",
      "name": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-10-30 12:51:57",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "211.233.221.207",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.8",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": 50,
      "rootDeviceName": "/dev/vda",
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
      "vNetId": "my08-mig-vnet-01",
      "cspVNetId": "151836",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "332997",
      "networkInterface": "nic-6113010",
      "securityGroupIds": [
        "my08-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tblv95nnkjn34ni5rsgb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-10-30T12:51:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:51:57Z",
          "completedTime": "2026-10-30T12:52:00Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my08-infra101-tbbnrantelivpl3hni6r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:52:01Z",
          "completedTime": "2026-10-30T12:52:02Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb666sqqpo5nodvufp1h 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my08-ng-influxdb-back-1",
      "uid": "tb48gt5jcbctsb16hcdt",
      "cspResourceName": "tb48gt5jcbctsb16hcdt",
      "cspResourceId": "146612727",
      "name": "my08-ng-influxdb-back-1",
      "nodeGroupId": "my08-ng-influxdb-back",
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
      "createdTime": "2026-10-30 12:51:57",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "211.233.221.202",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": 50,
      "rootDeviceName": "/dev/vda",
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
      "vNetId": "my08-mig-vnet-01",
      "cspVNetId": "151836",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "332997",
      "networkInterface": "nic-6113008",
      "securityGroupIds": [
        "my08-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tblv95nnkjn34ni5rsgb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-10-30T12:51:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:51:57Z",
          "completedTime": "2026-10-30T12:52:00Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my08-infra101-tbbnrantelivpl3hni6r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:52:01Z",
          "completedTime": "2026-10-30T12:52:03Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb48gt5jcbctsb16hcdt 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my08-ng-influxdb-back-2",
      "uid": "tbaqmvhvg3phq0tnoa1b",
      "cspResourceName": "tbaqmvhvg3phq0tnoa1b",
      "cspResourceId": "146612730",
      "name": "my08-ng-influxdb-back-2",
      "nodeGroupId": "my08-ng-influxdb-back",
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
      "createdTime": "2026-10-30 12:51:57",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "211.233.221.206",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.7",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": 50,
      "rootDeviceName": "/dev/vda",
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
      "vNetId": "my08-mig-vnet-01",
      "cspVNetId": "151836",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "332997",
      "networkInterface": "nic-6113009",
      "securityGroupIds": [
        "my08-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tblv95nnkjn34ni5rsgb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-10-30T12:51:57Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:51:57Z",
          "completedTime": "2026-10-30T12:52:01Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my08-infra101-tbbnrantelivpl3hni6r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:52:01Z",
          "completedTime": "2026-10-30T12:52:04Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbaqmvhvg3phq0tnoa1b 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my08-mig-vnet-01",
      "name": "my08-mig-vnet-01",
      "infraId": "my08-infra101",
      "vNetId": "my08-mig-vnet-01",
      "connectionNames": [
        "ncp-kr"
      ],
      "providerNames": [
        "ncp"
      ],
      "regionNames": [
        "KR"
      ],
      "nodeGroupIds": [
        "my08-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my08-ng-influxdb-back"
      ],
      "nodeIds": [
        "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my08-ng-influxdb-back-1",
        "my08-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "211.233.221.207",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb666sqqpo5nodvufp1h 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-influxdb-back-1",
            "nodeIp": "211.233.221.202",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb48gt5jcbctsb16hcdt 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-influxdb-back-2",
            "nodeIp": "211.233.221.206",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbaqmvhvg3phq0tnoa1b 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my08-infra101-tbbnrantelivpl3hni6r"
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
- **Count**: 5

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "infra": [
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
    },
    {
      "resourceType": "infra",
      "id": "my03-infra101",
      "uid": "tbdcanpcp70pd7mapdn4",
      "name": "my03-infra101",
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
        "cm-migrated-at": "2026-10-30T12:48:39Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my03-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my03-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbdcanpcp70pd7mapdn4"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb0ogg2jdctsppoih0d1",
          "cspResourceName": "tb0ogg2jdctsppoih0d1",
          "cspResourceId": "tb0ogg2jdctsppoih0d1",
          "name": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:49:01",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "keypair": "tb2h282hf7m5h9nubvn7",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-10-30 12:49:01",
            "sys.cspResourceId": "tb0ogg2jdctsppoih0d1",
            "sys.cspResourceName": "tb0ogg2jdctsppoih0d1",
            "sys.id": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tb0ogg2jdctsppoih0d1",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.47.115.170",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.3",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": 10,
          "rootDeviceName": "",
          "connectionName": "gcp-asia-northeast3",
          "connectionConfig": {
            "configName": "gcp-asia-northeast3",
            "providerName": "gcp",
            "driverName": "gcp-driver-v1.0.so",
            "credentialName": "gcp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "gcp-asia-northeast3",
            "regionZoneInfo": {
              "assignedRegion": "asia-northeast3",
              "assignedZone": "asia-northeast3-a"
            },
            "regionDetail": {
              "regionId": "asia-northeast3",
              "regionName": "asia-northeast3",
              "description": "Seoul South Korea",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.2,
                "longitude": 127
              },
              "zones": [
                "asia-northeast3-a",
                "asia-northeast3-b",
                "asia-northeast3-c"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "gcp+asia-northeast3+n1-highcpu-2",
          "cspSpecName": "n1-highcpu-2",
          "spec": {
            "cspSpecName": "n1-highcpu-2",
            "vCPU": 2,
            "memoryGiB": 1.7568359,
            "costPerHour": 0.09099
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20"
          },
          "vNetId": "my03-mig-vnet-01",
          "cspVNetId": "tbplbufb5140g64c4e6a",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tb19rvmuqe5cd87gv5eh",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tb2h282hf7m5h9nubvn7"
        },
        {
          "resourceType": "node",
          "id": "my03-ng-influxdb-back-1",
          "uid": "tb0i6hj6t9k5dgoo7360",
          "cspResourceName": "tb0i6hj6t9k5dgoo7360",
          "cspResourceId": "tb0i6hj6t9k5dgoo7360",
          "name": "my03-ng-influxdb-back-1",
          "nodeGroupId": "my03-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:49:00",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "keypair": "tb2h282hf7m5h9nubvn7",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-10-30 12:49:00",
            "sys.cspResourceId": "tb0i6hj6t9k5dgoo7360",
            "sys.cspResourceName": "tb0i6hj6t9k5dgoo7360",
            "sys.id": "my03-ng-influxdb-back-1",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-influxdb-back",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tb0i6hj6t9k5dgoo7360",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.64.190.160",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.2",
          "privateDNS": "",
          "rootDiskType": "hyperdisk-balanced",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "gcp-asia-northeast3",
          "connectionConfig": {
            "configName": "gcp-asia-northeast3",
            "providerName": "gcp",
            "driverName": "gcp-driver-v1.0.so",
            "credentialName": "gcp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "gcp-asia-northeast3",
            "regionZoneInfo": {
              "assignedRegion": "asia-northeast3",
              "assignedZone": "asia-northeast3-a"
            },
            "regionDetail": {
              "regionId": "asia-northeast3",
              "regionName": "asia-northeast3",
              "description": "Seoul South Korea",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.2,
                "longitude": 127
              },
              "zones": [
                "asia-northeast3-a",
                "asia-northeast3-b",
                "asia-northeast3-c"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "gcp+asia-northeast3+n4-standard-4",
          "cspSpecName": "n4-standard-4",
          "spec": {
            "cspSpecName": "n4-standard-4",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.232954
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20"
          },
          "vNetId": "my03-mig-vnet-01",
          "cspVNetId": "tbplbufb5140g64c4e6a",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tb19rvmuqe5cd87gv5eh",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tb2h282hf7m5h9nubvn7"
        },
        {
          "resourceType": "node",
          "id": "my03-ng-influxdb-back-2",
          "uid": "tb4ceb9m4usomf7ac0m5",
          "cspResourceName": "tb4ceb9m4usomf7ac0m5",
          "cspResourceId": "tb4ceb9m4usomf7ac0m5",
          "name": "my03-ng-influxdb-back-2",
          "nodeGroupId": "my03-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:49:02",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "keypair": "tb2h282hf7m5h9nubvn7",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-10-30 12:49:02",
            "sys.cspResourceId": "tb4ceb9m4usomf7ac0m5",
            "sys.cspResourceName": "tb4ceb9m4usomf7ac0m5",
            "sys.id": "my03-ng-influxdb-back-2",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-influxdb-back",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tb4ceb9m4usomf7ac0m5",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.50.27.124",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "hyperdisk-balanced",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "gcp-asia-northeast3",
          "connectionConfig": {
            "configName": "gcp-asia-northeast3",
            "providerName": "gcp",
            "driverName": "gcp-driver-v1.0.so",
            "credentialName": "gcp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "gcp-asia-northeast3",
            "regionZoneInfo": {
              "assignedRegion": "asia-northeast3",
              "assignedZone": "asia-northeast3-a"
            },
            "regionDetail": {
              "regionId": "asia-northeast3",
              "regionName": "asia-northeast3",
              "description": "Seoul South Korea",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.2,
                "longitude": 127
              },
              "zones": [
                "asia-northeast3-a",
                "asia-northeast3-b",
                "asia-northeast3-c"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "gcp+asia-northeast3+n4-standard-4",
          "cspSpecName": "n4-standard-4",
          "spec": {
            "cspSpecName": "n4-standard-4",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.232954
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20"
          },
          "vNetId": "my03-mig-vnet-01",
          "cspVNetId": "tbplbufb5140g64c4e6a",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tb19rvmuqe5cd87gv5eh",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tb2h282hf7m5h9nubvn7"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my05-infra101",
      "uid": "tb8u5qft8v04a887h53q",
      "name": "my05-infra101",
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
        "cm-migrated-at": "2026-10-30T12:47:54Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my05-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tb8u5qft8v04a887h53q"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbuku56fj76km7vqa0su",
          "cspResourceName": "tbuku56fj76km7vqa0su",
          "cspResourceId": "ins-mhkac4j9",
          "name": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:48:23",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-10-30 12:48:23",
            "sys.cspResourceId": "ins-mhkac4j9",
            "sys.cspResourceName": "tbuku56fj76km7vqa0su",
            "sys.id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tbuku56fj76km7vqa0su",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.166.11.218",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.8",
          "privateDNS": "",
          "rootDiskType": "CLOUD_HSSD",
          "rootDiskSize": 50,
          "rootDeviceName": "",
          "connectionName": "tencent-ap-seoul",
          "connectionConfig": {
            "configName": "tencent-ap-seoul",
            "providerName": "tencent",
            "driverName": "tencent-driver-v1.0.so",
            "credentialName": "tencent",
            "credentialHolder": "admin",
            "regionZoneInfoName": "tencent-ap-seoul",
            "regionZoneInfo": {
              "assignedRegion": "ap-seoul",
              "assignedZone": "ap-seoul-1"
            },
            "regionDetail": {
              "regionId": "ap-seoul",
              "regionName": "ap-seoul",
              "description": "Seoul",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.566536,
                "longitude": 126.977966
              },
              "zones": [
                "ap-seoul-1",
                "ap-seoul-2"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "tencent+ap-seoul+sa5.medium2",
          "cspSpecName": "SA5.MEDIUM2",
          "spec": {
            "cspSpecName": "SA5.MEDIUM2",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.06
          },
          "imageId": "img-487zeit5",
          "cspImageName": "img-487zeit5",
          "image": {
            "resourceType": "image",
            "cspImageName": "img-487zeit5",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Server 22.04 LTS 64bit"
          },
          "vNetId": "my05-mig-vnet-01",
          "cspVNetId": "vpc-rtznlhxm",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-19h6xb1b",
          "networkInterface": "eni-9sa7jv46",
          "securityGroupIds": [
            "my05-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-7ic7ibdp"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-1",
          "uid": "tb42prmq38mkramqj7kf",
          "cspResourceName": "tb42prmq38mkramqj7kf",
          "cspResourceId": "ins-4ap62gvz",
          "name": "my05-ng-influxdb-back-1",
          "nodeGroupId": "my05-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:48:20",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-10-30 12:48:20",
            "sys.cspResourceId": "ins-4ap62gvz",
            "sys.cspResourceName": "tb42prmq38mkramqj7kf",
            "sys.id": "my05-ng-influxdb-back-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tb42prmq38mkramqj7kf",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.131.254.242",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.10",
          "privateDNS": "",
          "rootDiskType": "CLOUD_PREMIUM",
          "rootDiskSize": 50,
          "rootDeviceName": "",
          "connectionName": "tencent-ap-seoul",
          "connectionConfig": {
            "configName": "tencent-ap-seoul",
            "providerName": "tencent",
            "driverName": "tencent-driver-v1.0.so",
            "credentialName": "tencent",
            "credentialHolder": "admin",
            "regionZoneInfoName": "tencent-ap-seoul",
            "regionZoneInfo": {
              "assignedRegion": "ap-seoul",
              "assignedZone": "ap-seoul-1"
            },
            "regionDetail": {
              "regionId": "ap-seoul",
              "regionName": "ap-seoul",
              "description": "Seoul",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.566536,
                "longitude": 126.977966
              },
              "zones": [
                "ap-seoul-1",
                "ap-seoul-2"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "tencent+ap-seoul+s6.large16",
          "cspSpecName": "S6.LARGE16",
          "spec": {
            "cspSpecName": "S6.LARGE16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.23
          },
          "imageId": "img-7rotv4ux",
          "cspImageName": "img-7rotv4ux",
          "image": {
            "resourceType": "image",
            "cspImageName": "img-7rotv4ux",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI"
          },
          "vNetId": "my05-mig-vnet-01",
          "cspVNetId": "vpc-rtznlhxm",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-19h6xb1b",
          "networkInterface": "eni-h8jn4zvw",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-7ic7ibdp"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-2",
          "uid": "tbho8c20nshb3d7pklqn",
          "cspResourceName": "tbho8c20nshb3d7pklqn",
          "cspResourceId": "ins-q40emar7",
          "name": "my05-ng-influxdb-back-2",
          "nodeGroupId": "my05-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:48:20",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-10-30 12:48:20",
            "sys.cspResourceId": "ins-q40emar7",
            "sys.cspResourceName": "tbho8c20nshb3d7pklqn",
            "sys.id": "my05-ng-influxdb-back-2",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tbho8c20nshb3d7pklqn",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.166.9.19",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.2",
          "privateDNS": "",
          "rootDiskType": "CLOUD_PREMIUM",
          "rootDiskSize": 50,
          "rootDeviceName": "",
          "connectionName": "tencent-ap-seoul",
          "connectionConfig": {
            "configName": "tencent-ap-seoul",
            "providerName": "tencent",
            "driverName": "tencent-driver-v1.0.so",
            "credentialName": "tencent",
            "credentialHolder": "admin",
            "regionZoneInfoName": "tencent-ap-seoul",
            "regionZoneInfo": {
              "assignedRegion": "ap-seoul",
              "assignedZone": "ap-seoul-1"
            },
            "regionDetail": {
              "regionId": "ap-seoul",
              "regionName": "ap-seoul",
              "description": "Seoul",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.566536,
                "longitude": 126.977966
              },
              "zones": [
                "ap-seoul-1",
                "ap-seoul-2"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "tencent+ap-seoul+s6.large16",
          "cspSpecName": "S6.LARGE16",
          "spec": {
            "cspSpecName": "S6.LARGE16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.23
          },
          "imageId": "img-7rotv4ux",
          "cspImageName": "img-7rotv4ux",
          "image": {
            "resourceType": "image",
            "cspImageName": "img-7rotv4ux",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI"
          },
          "vNetId": "my05-mig-vnet-01",
          "cspVNetId": "vpc-rtznlhxm",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-19h6xb1b",
          "networkInterface": "eni-ajmm99d6",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-7ic7ibdp"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my06-infra101",
      "uid": "tbm6c5c56h6nf7roqjvh",
      "name": "my06-infra101",
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
        "cm-migrated-at": "2026-10-30T12:49:09Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my06-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbm6c5c56h6nf7roqjvh"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbar0eeronn7me24e1jr",
          "cspResourceName": "tbar0eeronn7me24e1jr",
          "cspResourceId": "02h7_8466ac07-73a0-45a7-947e-f6025a81d369",
          "name": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:49:27",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-10-30 12:49:27",
            "sys.cspResourceId": "02h7_8466ac07-73a0-45a7-947e-f6025a81d369",
            "sys.cspResourceName": "tbar0eeronn7me24e1jr",
            "sys.id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbar0eeronn7me24e1jr",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.100.42",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.5",
          "privateDNS": "",
          "rootDiskType": "general-purpose",
          "rootDiskSize": 100,
          "rootDeviceName": "",
          "connectionName": "ibm-au-syd",
          "connectionConfig": {
            "configName": "ibm-au-syd",
            "providerName": "ibm",
            "driverName": "ibm-driver-v1.0.so",
            "credentialName": "ibm",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ibm-au-syd",
            "regionZoneInfo": {
              "assignedRegion": "au-syd",
              "assignedZone": "au-syd-1"
            },
            "regionDetail": {
              "regionId": "au-syd",
              "regionName": "au-syd",
              "description": "Sydney (Australia)",
              "location": {
                "display": "Australia (Sydney)",
                "latitude": -33.86882,
                "longitude": 151.209296
              },
              "zones": [
                "au-syd-1",
                "au-syd-2",
                "au-syd-3"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ibm+au-syd+cx2-2x4",
          "cspSpecName": "cx2-2x4",
          "spec": {
            "cspSpecName": "cx2-2x4",
            "vCPU": 2,
            "memoryGiB": 4,
            "costPerHour": 0.106
          },
          "imageId": "r026-b879575b-8209-4e6a-a46e-0e01adf56aaf",
          "cspImageName": "r026-b879575b-8209-4e6a-a46e-0e01adf56aaf",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-b879575b-8209-4e6a-a46e-0e01adf56aaf",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-mig-vnet-01",
          "cspVNetId": "r026-2f361af0-73d0-46a0-9d2b-b8da85809476",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-cc88798d-0035-4f70-bc4d-80b630a00c4c",
          "networkInterface": "cannon-unending-glancing-clumsily",
          "securityGroupIds": [
            "my06-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-f6cfb962-7eff-4047-8665-12fe4e8077bf"
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-1",
          "uid": "tbf0cqgntsded5k3h3ln",
          "cspResourceName": "tbf0cqgntsded5k3h3ln",
          "cspResourceId": "02h7_3957719b-2881-4faa-a682-8c6f5569a4f5",
          "name": "my06-ng-influxdb-back-1",
          "nodeGroupId": "my06-ng-influxdb-back",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:49:28",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-10-30 12:49:28",
            "sys.cspResourceId": "02h7_3957719b-2881-4faa-a682-8c6f5569a4f5",
            "sys.cspResourceName": "tbf0cqgntsded5k3h3ln",
            "sys.id": "my06-ng-influxdb-back-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-influxdb-back",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbf0cqgntsded5k3h3ln",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.90.173",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "general-purpose",
          "rootDiskSize": 100,
          "rootDeviceName": "",
          "connectionName": "ibm-au-syd",
          "connectionConfig": {
            "configName": "ibm-au-syd",
            "providerName": "ibm",
            "driverName": "ibm-driver-v1.0.so",
            "credentialName": "ibm",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ibm-au-syd",
            "regionZoneInfo": {
              "assignedRegion": "au-syd",
              "assignedZone": "au-syd-1"
            },
            "regionDetail": {
              "regionId": "au-syd",
              "regionName": "au-syd",
              "description": "Sydney (Australia)",
              "location": {
                "display": "Australia (Sydney)",
                "latitude": -33.86882,
                "longitude": 151.209296
              },
              "zones": [
                "au-syd-1",
                "au-syd-2",
                "au-syd-3"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ibm+au-syd+bx2-4x16",
          "cspSpecName": "bx2-4x16",
          "spec": {
            "cspSpecName": "bx2-4x16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.241
          },
          "imageId": "r026-b879575b-8209-4e6a-a46e-0e01adf56aaf",
          "cspImageName": "r026-b879575b-8209-4e6a-a46e-0e01adf56aaf",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-b879575b-8209-4e6a-a46e-0e01adf56aaf",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-mig-vnet-01",
          "cspVNetId": "r026-2f361af0-73d0-46a0-9d2b-b8da85809476",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-cc88798d-0035-4f70-bc4d-80b630a00c4c",
          "networkInterface": "motivate-shape-slinky-turbulent",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-f6cfb962-7eff-4047-8665-12fe4e8077bf"
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-2",
          "uid": "tbvicb2g6blfup7t2rn1",
          "cspResourceName": "tbvicb2g6blfup7t2rn1",
          "cspResourceId": "02h7_72a79220-56f5-411a-a2f8-073d38e9848a",
          "name": "my06-ng-influxdb-back-2",
          "nodeGroupId": "my06-ng-influxdb-back",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:49:28",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-10-30 12:49:28",
            "sys.cspResourceId": "02h7_72a79220-56f5-411a-a2f8-073d38e9848a",
            "sys.cspResourceName": "tbvicb2g6blfup7t2rn1",
            "sys.id": "my06-ng-influxdb-back-2",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-influxdb-back",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbvicb2g6blfup7t2rn1",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.97.164",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "general-purpose",
          "rootDiskSize": 100,
          "rootDeviceName": "",
          "connectionName": "ibm-au-syd",
          "connectionConfig": {
            "configName": "ibm-au-syd",
            "providerName": "ibm",
            "driverName": "ibm-driver-v1.0.so",
            "credentialName": "ibm",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ibm-au-syd",
            "regionZoneInfo": {
              "assignedRegion": "au-syd",
              "assignedZone": "au-syd-1"
            },
            "regionDetail": {
              "regionId": "au-syd",
              "regionName": "au-syd",
              "description": "Sydney (Australia)",
              "location": {
                "display": "Australia (Sydney)",
                "latitude": -33.86882,
                "longitude": 151.209296
              },
              "zones": [
                "au-syd-1",
                "au-syd-2",
                "au-syd-3"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ibm+au-syd+bx2-4x16",
          "cspSpecName": "bx2-4x16",
          "spec": {
            "cspSpecName": "bx2-4x16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.241
          },
          "imageId": "r026-b879575b-8209-4e6a-a46e-0e01adf56aaf",
          "cspImageName": "r026-b879575b-8209-4e6a-a46e-0e01adf56aaf",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-b879575b-8209-4e6a-a46e-0e01adf56aaf",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-mig-vnet-01",
          "cspVNetId": "r026-2f361af0-73d0-46a0-9d2b-b8da85809476",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-cc88798d-0035-4f70-bc4d-80b630a00c4c",
          "networkInterface": "axiom-groovy-haven-isolated",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-f6cfb962-7eff-4047-8665-12fe4e8077bf"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my08-infra101",
      "uid": "tbqisg04kagcpshkp314",
      "name": "my08-infra101",
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
        "cm-migrated-at": "2026-10-30T12:50:18Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my08-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my08-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbqisg04kagcpshkp314"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb666sqqpo5nodvufp1h",
          "cspResourceName": "tb666sqqpo5nodvufp1h",
          "cspResourceId": "146612733",
          "name": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:51:57",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2026-10-30 12:51:57",
            "sys.cspResourceId": "146612733",
            "sys.cspResourceName": "tb666sqqpo5nodvufp1h",
            "sys.id": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my08-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my08-mig-subnet-01",
            "sys.uid": "tb666sqqpo5nodvufp1h",
            "sys.vNetId": "my08-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "211.233.221.207",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.8",
          "privateDNS": "",
          "rootDiskType": "HDD",
          "rootDiskSize": 50,
          "rootDeviceName": "",
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
          "vNetId": "my08-mig-vnet-01",
          "cspVNetId": "151836",
          "subnetId": "my08-mig-subnet-01",
          "cspSubnetId": "332997",
          "networkInterface": "nic-6113010",
          "securityGroupIds": [
            "my08-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my08-mig-sshkey-01",
          "cspSshKeyId": "tblv95nnkjn34ni5rsgb"
        },
        {
          "resourceType": "node",
          "id": "my08-ng-influxdb-back-1",
          "uid": "tb48gt5jcbctsb16hcdt",
          "cspResourceName": "tb48gt5jcbctsb16hcdt",
          "cspResourceId": "146612727",
          "name": "my08-ng-influxdb-back-1",
          "nodeGroupId": "my08-ng-influxdb-back",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:51:57",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2026-10-30 12:51:57",
            "sys.cspResourceId": "146612727",
            "sys.cspResourceName": "tb48gt5jcbctsb16hcdt",
            "sys.id": "my08-ng-influxdb-back-1",
            "sys.infraId": "my08-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my08-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my08-ng-influxdb-back",
            "sys.subnetId": "my08-mig-subnet-01",
            "sys.uid": "tb48gt5jcbctsb16hcdt",
            "sys.vNetId": "my08-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "211.233.221.202",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "HDD",
          "rootDiskSize": 50,
          "rootDeviceName": "",
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
          "vNetId": "my08-mig-vnet-01",
          "cspVNetId": "151836",
          "subnetId": "my08-mig-subnet-01",
          "cspSubnetId": "332997",
          "networkInterface": "nic-6113008",
          "securityGroupIds": [
            "my08-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my08-mig-sshkey-01",
          "cspSshKeyId": "tblv95nnkjn34ni5rsgb"
        },
        {
          "resourceType": "node",
          "id": "my08-ng-influxdb-back-2",
          "uid": "tbaqmvhvg3phq0tnoa1b",
          "cspResourceName": "tbaqmvhvg3phq0tnoa1b",
          "cspResourceId": "146612730",
          "name": "my08-ng-influxdb-back-2",
          "nodeGroupId": "my08-ng-influxdb-back",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-30 12:51:57",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2026-10-30 12:51:57",
            "sys.cspResourceId": "146612730",
            "sys.cspResourceName": "tbaqmvhvg3phq0tnoa1b",
            "sys.id": "my08-ng-influxdb-back-2",
            "sys.infraId": "my08-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my08-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my08-ng-influxdb-back",
            "sys.subnetId": "my08-mig-subnet-01",
            "sys.uid": "tbaqmvhvg3phq0tnoa1b",
            "sys.vNetId": "my08-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "211.233.221.206",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.7",
          "privateDNS": "",
          "rootDiskType": "HDD",
          "rootDiskSize": 50,
          "rootDeviceName": "",
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
          "vNetId": "my08-mig-vnet-01",
          "cspVNetId": "151836",
          "subnetId": "my08-mig-subnet-01",
          "cspSubnetId": "332997",
          "networkInterface": "nic-6113009",
          "securityGroupIds": [
            "my08-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my08-mig-sshkey-01",
          "cspSshKeyId": "tblv95nnkjn34ni5rsgb"
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
- **IDs**: [my02-infra101 my03-infra101 my05-infra101 my06-infra101 my08-infra101]

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
  "id": "my08-infra101",
  "uid": "tbqisg04kagcpshkp314",
  "name": "my08-infra101",
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
    "cm-migrated-at": "2026-10-30T12:50:18Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my08-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my08-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbqisg04kagcpshkp314"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb666sqqpo5nodvufp1h",
      "cspResourceName": "tb666sqqpo5nodvufp1h",
      "cspResourceId": "146612733",
      "name": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-10-30 12:51:57",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "211.233.221.207",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.8",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": 50,
      "rootDeviceName": "/dev/vda",
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
      "vNetId": "my08-mig-vnet-01",
      "cspVNetId": "151836",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "332997",
      "networkInterface": "nic-6113010",
      "securityGroupIds": [
        "my08-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tblv95nnkjn34ni5rsgb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-10-30T12:51:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:51:57Z",
          "completedTime": "2026-10-30T12:52:00Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my08-infra101-tbbnrantelivpl3hni6r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:52:01Z",
          "completedTime": "2026-10-30T12:52:02Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb666sqqpo5nodvufp1h 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my08-ng-influxdb-back-1",
      "uid": "tb48gt5jcbctsb16hcdt",
      "cspResourceName": "tb48gt5jcbctsb16hcdt",
      "cspResourceId": "146612727",
      "name": "my08-ng-influxdb-back-1",
      "nodeGroupId": "my08-ng-influxdb-back",
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
      "createdTime": "2026-10-30 12:51:57",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "211.233.221.202",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": 50,
      "rootDeviceName": "/dev/vda",
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
      "vNetId": "my08-mig-vnet-01",
      "cspVNetId": "151836",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "332997",
      "networkInterface": "nic-6113008",
      "securityGroupIds": [
        "my08-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tblv95nnkjn34ni5rsgb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-10-30T12:51:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:51:57Z",
          "completedTime": "2026-10-30T12:52:00Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my08-infra101-tbbnrantelivpl3hni6r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:52:01Z",
          "completedTime": "2026-10-30T12:52:03Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb48gt5jcbctsb16hcdt 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my08-ng-influxdb-back-2",
      "uid": "tbaqmvhvg3phq0tnoa1b",
      "cspResourceName": "tbaqmvhvg3phq0tnoa1b",
      "cspResourceId": "146612730",
      "name": "my08-ng-influxdb-back-2",
      "nodeGroupId": "my08-ng-influxdb-back",
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
      "createdTime": "2026-10-30 12:51:57",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "211.233.221.206",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.7",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": 50,
      "rootDeviceName": "/dev/vda",
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
      "vNetId": "my08-mig-vnet-01",
      "cspVNetId": "151836",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "332997",
      "networkInterface": "nic-6113009",
      "securityGroupIds": [
        "my08-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tblv95nnkjn34ni5rsgb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-10-30T12:51:57Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:51:57Z",
          "completedTime": "2026-10-30T12:52:01Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my08-infra101-tbbnrantelivpl3hni6r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:52:01Z",
          "completedTime": "2026-10-30T12:52:04Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbaqmvhvg3phq0tnoa1b 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my08-mig-vnet-01",
      "name": "my08-mig-vnet-01",
      "infraId": "my08-infra101",
      "vNetId": "my08-mig-vnet-01",
      "connectionNames": [
        "ncp-kr"
      ],
      "providerNames": [
        "ncp"
      ],
      "regionNames": [
        "KR"
      ],
      "nodeGroupIds": [
        "my08-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my08-ng-influxdb-back"
      ],
      "nodeIds": [
        "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my08-ng-influxdb-back-1",
        "my08-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "211.233.221.207",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb666sqqpo5nodvufp1h 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-influxdb-back-1",
            "nodeIp": "211.233.221.202",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb48gt5jcbctsb16hcdt 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-influxdb-back-2",
            "nodeIp": "211.233.221.206",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbaqmvhvg3phq0tnoa1b 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my08-infra101-tbbnrantelivpl3hni6r"
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
      "id": "my08-ng-influxdb-back",
      "cspResourceName": "tbkv2afhv9qgtsg0qb3t",
      "cspResourceId": "146612759",
      "name": "my08-ng-influxdb-back",
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
      "type": "PUBLIC",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "ip": "211.233.204.202",
        "port": "9999",
        "keyValueList": [
          {
            "key": "LoadBalancerInstanceNo",
            "value": "146612759"
          },
          {
            "key": "LoadBalancerListenerNo",
            "value": "720859"
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
            "value": "798871"
          }
        ]
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my08-ng-influxdb-back",
        "nodes": [
          "my08-ng-influxdb-back-1",
          "my08-ng-influxdb-back-2"
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
      "createdTime": "2026-10-30T21:52:09+09:00",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "LoadBalancerInstanceNo",
          "value": "146612759"
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
          "value": "2026-10-30T21:52:09+0900"
        },
        {
          "key": "LoadBalancerName",
          "value": "tbkv2afhv9qgtsg0qb3t"
        },
        {
          "key": "LoadBalancerDomain",
          "value": "tbkv2afhv9qgtsg0qb3t-146612759-6e462c3ca2ce.kr.lb.naverncp.com"
        },
        {
          "key": "LoadBalancerIpList",
          "value": "211.233.204.202"
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
          "value": "151836"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "SubnetNoList",
          "value": "332998"
        },
        {
          "key": "LoadBalancerSubnetList",
          "value": "{zoneCode:KR-1,subnetNo:332998,publicIpInstanceNo:146612765}"
        },
        {
          "key": "LoadBalancerListenerNoList",
          "value": "720859"
        },
        {
          "key": "ListenerId",
          "value": "720859"
        },
        {
          "key": "HealthCheckerId",
          "value": "6505114"
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
    "id": "my08-ng-influxdb-back",
    "cspResourceName": "tbkv2afhv9qgtsg0qb3t",
    "cspResourceId": "146612759",
    "name": "my08-ng-influxdb-back",
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
    "type": "PUBLIC",
    "scope": "REGION",
    "listener": {
      "protocol": "TCP",
      "ip": "211.233.204.202",
      "port": "9999",
      "keyValueList": [
        {
          "key": "LoadBalancerInstanceNo",
          "value": "146612759"
        },
        {
          "key": "LoadBalancerListenerNo",
          "value": "720859"
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
          "value": "798871"
        }
      ]
    },
    "targetGroup": {
      "protocol": "TCP",
      "port": "8086",
      "nodeGroupId": "my08-ng-influxdb-back",
      "nodes": [
        "my08-ng-influxdb-back-1",
        "my08-ng-influxdb-back-2"
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
    "createdTime": "2026-10-30T21:52:09+09:00",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "LoadBalancerInstanceNo",
        "value": "146612759"
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
        "value": "2026-10-30T21:52:09+0900"
      },
      {
        "key": "LoadBalancerName",
        "value": "tbkv2afhv9qgtsg0qb3t"
      },
      {
        "key": "LoadBalancerDomain",
        "value": "tbkv2afhv9qgtsg0qb3t-146612759-6e462c3ca2ce.kr.lb.naverncp.com"
      },
      {
        "key": "LoadBalancerIpList",
        "value": "211.233.204.202"
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
        "value": "151836"
      },
      {
        "key": "RegionCode",
        "value": "KR"
      },
      {
        "key": "SubnetNoList",
        "value": "332998"
      },
      {
        "key": "LoadBalancerSubnetList",
        "value": "{zoneCode:KR-1,subnetNo:332998,publicIpInstanceNo:146612765}"
      },
      {
        "key": "LoadBalancerListenerNoList",
        "value": "720859"
      },
      {
        "key": "ListenerId",
        "value": "720859"
      },
      {
        "key": "HealthCheckerId",
        "value": "6505114"
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
  "id": "my08-ng-influxdb-back",
  "cspResourceName": "tbkv2afhv9qgtsg0qb3t",
  "cspResourceId": "146612759",
  "name": "my08-ng-influxdb-back",
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
  "type": "PUBLIC",
  "scope": "REGION",
  "listener": {
    "protocol": "TCP",
    "ip": "211.233.204.202",
    "port": "9999",
    "keyValueList": [
      {
        "key": "LoadBalancerInstanceNo",
        "value": "146612759"
      },
      {
        "key": "LoadBalancerListenerNo",
        "value": "720859"
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
        "value": "798871"
      }
    ]
  },
  "targetGroup": {
    "protocol": "TCP",
    "port": "8086",
    "nodeGroupId": "my08-ng-influxdb-back",
    "nodes": [
      "my08-ng-influxdb-back-1",
      "my08-ng-influxdb-back-2"
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
  "createdTime": "2026-10-30T21:52:09+09:00",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "LoadBalancerInstanceNo",
      "value": "146612759"
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
      "value": "2026-10-30T21:52:09+0900"
    },
    {
      "key": "LoadBalancerName",
      "value": "tbkv2afhv9qgtsg0qb3t"
    },
    {
      "key": "LoadBalancerDomain",
      "value": "tbkv2afhv9qgtsg0qb3t-146612759-6e462c3ca2ce.kr.lb.naverncp.com"
    },
    {
      "key": "LoadBalancerIpList",
      "value": "211.233.204.202"
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
      "value": "151836"
    },
    {
      "key": "RegionCode",
      "value": "KR"
    },
    {
      "key": "SubnetNoList",
      "value": "332998"
    },
    {
      "key": "LoadBalancerSubnetList",
      "value": "{zoneCode:KR-1,subnetNo:332998,publicIpInstanceNo:146612765}"
    },
    {
      "key": "LoadBalancerListenerNoList",
      "value": "720859"
    },
    {
      "key": "ListenerId",
      "value": "720859"
    },
    {
      "key": "HealthCheckerId",
      "value": "6505114"
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

**Generated At:** 2026-10-30 12:56:06

**Namespace:** mig01

**Infra Name:** my08-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my08-infra101 |
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
| s4-g3 | 4 | 16.0 | - | x86_64 | default | $0.1747 | 2 |
| ci2-g3 | 2 | 4.0 | - | x86_64 | default | $0.0730 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| 23214590 | ubuntu-22.04-base (Hypervisor:KVM) | Ubuntu 22.04 | Linux/UNIX | x86_64 | Common BlockStorage 1 | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | 146612733 | Running | 2 vCPU, 4.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 211.233.221.207<br>**Private IP:** 10.0.1.8<br>**SGs:** my08-mig-sg-02<br>**SSH:** my08-mig-sshkey-01 |
| my08-ng-influxdb-back-1 | 146612727 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 211.233.221.202<br>**Private IP:** 10.0.1.6<br>**SGs:** my08-mig-sg-01<br>**SSH:** my08-mig-sshkey-01 |
| my08-ng-influxdb-back-2 | 146612730 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 211.233.221.206<br>**Private IP:** 10.0.1.7<br>**SGs:** my08-mig-sg-01<br>**SSH:** my08-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my08-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my08-mig-vnet-01 |
| **CSP VNet ID** | 151836 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ncp-kr |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my08-mig-subnet-01 | 332997 | 10.0.1.0/24 | KR-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my08-mig-sshkey-01 | tblv95nnkjn34ni5rsgb | cb-user | da:5f:e5:5f:50:de:5c:ff:8f:2d:86:c0:a2:da:85:e5 |

### Security Groups

#### Security Group: my08-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my08-mig-sg-01 |
| **CSP Security Group ID** | 413150 |
| **VNet** | my08-mig-vnet-01 |
| **Rule Count** | 10 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my08-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my08-mig-sg-02 |
| **CSP Security Group ID** | 413151 |
| **VNet** | my08-mig-vnet-01 |
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
| my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | ci2-g3 | $0.0730 | $52.56 |
| my08-ng-influxdb-back-1 | s4-g3 | $0.1747 | $125.78 |
| my08-ng-influxdb-back-2 | s4-g3 | $0.1747 | $125.78 |




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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my08-infra101)",
  "success": true
}
```

