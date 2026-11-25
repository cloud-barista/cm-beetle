# CM-Beetle test results for NCP (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with NCP cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.5+ (83361e8)
- imdl: v0.1.22+ (83361e8)
- CB-Tumblebug: v0.14.0
- CB-Spider: v0.14.0
- CB-MapUI: v0.14.0
- Target CSP: NCP
- Target Region: kr
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: November 25, 2026
- Test Time: 18:17:02 KST
- Test Execution: 2026-11-25 18:17:02 KST

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
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 1.28s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 453ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 3m8.244s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 173ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 21ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 31ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 968ms | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 1m45.128s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 9ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 13ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 1m28.639s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.253s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.257s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 52.487s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m19.334s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 9m37.658133585s

*Test executed on November 25, 2026 at 18:17:02 KST (2026-11-25 18:17:02 KST) using CM-Beetle automated test CLI*

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
            "name": "ng-ip-10-0-1-30-62161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=66.7%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+c2-g3",
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
          "uid": "tb4u3q0oag92chophfq8",
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
          "costPerHour": 0.1745,
          "orderInFilteredResult": 8,
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
              "value": "145429462,143396791,109093105,107029409,104630229,25495367,23221307,23221289,23214590,19463675"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-11-25T09:17:08.028543004Z",
          "availabilityRefreshAfter": "2026-11-25T10:17:08.028543004Z"
        },
        {
          "id": "ncp+kr+c2-g3",
          "uid": "tb2k6lp6lcnqpv796ku0",
          "cspSpecName": "c2-g3",
          "name": "ncp+kr+c2-g3",
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
          "costPerHour": 0.0677,
          "orderInFilteredResult": 3,
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
              "value": "c2-g3"
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
              "value": "SVR.VSVR.HICPU.C002.M004.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 2EA, Memory 4GB"
            },
            {
              "key": "ServerSpecNo",
              "value": "236"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "145429462,143396791,109093105,107029409,104630229,25495367,23221307,23221289,23214590,19463675"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-11-25T09:17:08.155059066Z",
          "availabilityRefreshAfter": "2026-11-25T10:17:08.155059066Z"
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
          "uid": "tbmthl0a6gh7673kfp2q",
          "name": "23214590",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ncp-kr",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:24:53 Thu",
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=72.2% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "ncp",
        "region": "kr"
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
            "name": "ng-ip-10-0-1-30-62161624",
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
          "id": "ncp+kr+m2-g3",
          "uid": "tb3lah09928tdi13klp0",
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
          "costPerHour": 0.1171,
          "orderInFilteredResult": 1,
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
              "value": "145429462,143396791,109093105,107029409,104630229,25495367,23221307,23221289,23214590,19463675"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-11-25T09:17:08.034442611Z",
          "availabilityRefreshAfter": "2026-11-25T10:17:08.034442611Z"
        },
        {
          "id": "ncp+kr+ci2-g3",
          "uid": "tbi0g37odil1dfkeaorh",
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
          "costPerHour": 0.0729,
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
              "value": "145429462,143396791,109093105,107029409,104630229,25495367,23221307,23221289,23214590,19463675"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-11-25T09:17:08.162187816Z",
          "availabilityRefreshAfter": "2026-11-25T10:17:08.162187816Z"
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
          "uid": "tbmthl0a6gh7673kfp2q",
          "name": "23214590",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ncp-kr",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:24:53 Thu",
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
        "name": "my08-ng-ip-10-0-1-30-62161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=66.7%",
        "connectionName": "ncp-kr",
        "specId": "ncp+kr+c2-g3",
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
      "uid": "tb4u3q0oag92chophfq8",
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
      "costPerHour": 0.1745,
      "orderInFilteredResult": 8,
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
          "value": "145429462,143396791,109093105,107029409,104630229,25495367,23221307,23221289,23214590,19463675"
        }
      ],
      "availabilityStatus": "Available",
      "availabilityCheckedAt": "2026-11-25T09:17:08.028543004Z",
      "availabilityRefreshAfter": "2026-11-25T10:17:08.028543004Z"
    },
    {
      "id": "ncp+kr+c2-g3",
      "uid": "tb2k6lp6lcnqpv796ku0",
      "cspSpecName": "c2-g3",
      "name": "ncp+kr+c2-g3",
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
      "costPerHour": 0.0677,
      "orderInFilteredResult": 3,
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
          "value": "c2-g3"
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
          "value": "SVR.VSVR.HICPU.C002.M004.G003"
        },
        {
          "key": "ServerSpecDescription",
          "value": "vCPU 2EA, Memory 4GB"
        },
        {
          "key": "ServerSpecNo",
          "value": "236"
        },
        {
          "key": "CorrespondingImageIds",
          "value": "145429462,143396791,109093105,107029409,104630229,25495367,23221307,23221289,23214590,19463675"
        }
      ],
      "availabilityStatus": "Available",
      "availabilityCheckedAt": "2026-11-25T09:17:08.155059066Z",
      "availabilityRefreshAfter": "2026-11-25T10:17:08.155059066Z"
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
      "uid": "tbmthl0a6gh7673kfp2q",
      "name": "23214590",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "ncp-kr",
      "infraType": "",
      "fetchedTime": "2026.10.29 04:24:53 Thu",
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
        "nodeGroupId": "my08-ng-influxdb-back"
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
  "uid": "tbqa7keur5fguo4egmrp",
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
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "cm-migrated-at": "2026-11-25T09:18:44Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my08-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my08-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbqa7keur5fguo4egmrp"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my08-ng-influxdb-back-1",
      "uid": "tba0vpa8a9koncq8jbhl",
      "cspResourceName": "tba0vpa8a9koncq8jbhl",
      "cspResourceId": "147164275",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:19:57",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "223.130.150.80",
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
      "specId": "ncp+kr+s4-g3",
      "cspSpecName": "s4-g3",
      "spec": {
        "cspSpecName": "s4-g3",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1745
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
      "cspVNetId": "153788",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "336709",
      "networkInterface": "nic-6190936",
      "securityGroupIds": [
        "my08-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tbmo1vkhfnoa7ta0p33p",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-11-25T09:20:07Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:06Z",
          "completedTime": "2026-11-25T09:20:09Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my08-infra101-tb2bpgqmq5j8mvp21njh",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:09Z",
          "completedTime": "2026-11-25T09:20:11Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tba0vpa8a9koncq8jbhl 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my08-ng-influxdb-back-2",
      "uid": "tb5rf1thbdblvaeqlndp",
      "cspResourceName": "tb5rf1thbdblvaeqlndp",
      "cspResourceId": "147164269",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:20:05",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "175.45.195.206",
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
        "costPerHour": 0.1745
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
      "cspVNetId": "153788",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "336709",
      "networkInterface": "nic-6190934",
      "securityGroupIds": [
        "my08-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tbmo1vkhfnoa7ta0p33p",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-11-25T09:20:07Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:06Z",
          "completedTime": "2026-11-25T09:20:08Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my08-infra101-tb2bpgqmq5j8mvp21njh",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:09Z",
          "completedTime": "2026-11-25T09:20:10Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb5rf1thbdblvaeqlndp 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my08-ng-ip-10-0-1-30-62161624-1",
      "uid": "tb78egghlm68bc17jcop",
      "cspResourceName": "tb78egghlm68bc17jcop",
      "cspResourceId": "147164272",
      "name": "my08-ng-ip-10-0-1-30-62161624-1",
      "nodeGroupId": "my08-ng-ip-10-0-1-30-62161624",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:20:06",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "223.130.159.128",
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
      "specId": "ncp+kr+c2-g3",
      "cspSpecName": "c2-g3",
      "spec": {
        "cspSpecName": "c2-g3",
        "vCPU": 2,
        "memoryGiB": 4,
        "costPerHour": 0.0677
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
      "cspVNetId": "153788",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "336709",
      "networkInterface": "nic-6190935",
      "securityGroupIds": [
        "my08-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tbmo1vkhfnoa7ta0p33p",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-11-25T09:20:07Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:06Z",
          "completedTime": "2026-11-25T09:20:09Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my08-infra101-tb2bpgqmq5j8mvp21njh",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:09Z",
          "completedTime": "2026-11-25T09:20:11Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb78egghlm68bc17jcop 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
        "my08-ng-influxdb-back",
        "my08-ng-ip-10-0-1-30-62161624"
      ],
      "nodeIds": [
        "my08-ng-influxdb-back-1",
        "my08-ng-influxdb-back-2",
        "my08-ng-ip-10-0-1-30-62161624-1"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my08-ng-influxdb-back",
      "representativeNodeId": "my08-ng-influxdb-back-1"
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
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-influxdb-back-2",
            "nodeIp": "175.45.195.206",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb5rf1thbdblvaeqlndp 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-ip-10-0-1-30-62161624-1",
            "nodeIp": "223.130.159.128",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb78egghlm68bc17jcop 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-influxdb-back-1",
            "nodeIp": "223.130.150.80",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tba0vpa8a9koncq8jbhl 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postScriptRequestId": "ps-my08-infra101-tb2bpgqmq5j8mvp21njh"
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
- **Count**: 7

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
      "status": "Terminating:3 (R:0/3)",
      "statusCount": {
        "countTotal": 3,
        "countCreating": 0,
        "countRunning": 0,
        "countFailed": 0,
        "countSuspended": 0,
        "countRebooting": 0,
        "countTerminated": 0,
        "countSuspending": 0,
        "countResuming": 0,
        "countTerminating": 3,
        "countRegistering": 0,
        "countReconciling": 0,
        "countUndefined": 0
      },
      "targetStatus": "Terminated",
      "targetAction": "Terminate",
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
          "status": "Terminating",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
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
          "status": "Terminating",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
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
          "status": "Terminating",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
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
      "id": "my03-infra101",
      "uid": "tbq870bov9ecui7jr7gv",
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
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "cm-migrated-at": "2026-11-25T09:17:21Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my03-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my03-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbq870bov9ecui7jr7gv"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my03-ng-influxdb-back-1",
          "uid": "tbjg90vgd6kbujirukfr",
          "cspResourceName": "tbjg90vgd6kbujirukfr",
          "cspResourceId": "tbjg90vgd6kbujirukfr",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:17:43",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "keypair": "tbdnf6sndttnf336lgim",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-11-25 09:17:43",
            "sys.cspResourceId": "tbjg90vgd6kbujirukfr",
            "sys.cspResourceName": "tbjg90vgd6kbujirukfr",
            "sys.id": "my03-ng-influxdb-back-1",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-influxdb-back",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tbjg90vgd6kbujirukfr",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.64.48.98",
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
            "memoryGiB": 16,
            "costPerHour": 0.232954
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21"
          },
          "vNetId": "my03-mig-vnet-01",
          "cspVNetId": "tb5dn9aum2qip3phdv8n",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tbgi6cba5avpr77flia4",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tbdnf6sndttnf336lgim"
        },
        {
          "resourceType": "node",
          "id": "my03-ng-influxdb-back-2",
          "uid": "tb1lee0l2cp9nso74pq6",
          "cspResourceName": "tb1lee0l2cp9nso74pq6",
          "cspResourceId": "tb1lee0l2cp9nso74pq6",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:17:40",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "keypair": "tbdnf6sndttnf336lgim",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-11-25 09:17:40",
            "sys.cspResourceId": "tb1lee0l2cp9nso74pq6",
            "sys.cspResourceName": "tb1lee0l2cp9nso74pq6",
            "sys.id": "my03-ng-influxdb-back-2",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-influxdb-back",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tb1lee0l2cp9nso74pq6",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.158.205.11",
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
            "memoryGiB": 16,
            "costPerHour": 0.232954
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21"
          },
          "vNetId": "my03-mig-vnet-01",
          "cspVNetId": "tb5dn9aum2qip3phdv8n",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tbgi6cba5avpr77flia4",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tbdnf6sndttnf336lgim"
        },
        {
          "resourceType": "node",
          "id": "my03-ng-ip-10-0-1-30-62161624-1",
          "uid": "tbr066qbkvoha4fjthgs",
          "cspResourceName": "tbr066qbkvoha4fjthgs",
          "cspResourceId": "tbr066qbkvoha4fjthgs",
          "name": "my03-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my03-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:17:41",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "keypair": "tbdnf6sndttnf336lgim",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-11-25 09:17:41",
            "sys.cspResourceId": "tbr066qbkvoha4fjthgs",
            "sys.cspResourceName": "tbr066qbkvoha4fjthgs",
            "sys.id": "my03-ng-ip-10-0-1-30-62161624-1",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-ip-10-0-1-30-62161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-ip-10-0-1-30-62161624",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tbr066qbkvoha4fjthgs",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.22.89.144",
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
          "specId": "gcp+asia-northeast3+n2-highcpu-2",
          "cspSpecName": "n2-highcpu-2",
          "spec": {
            "cspSpecName": "n2-highcpu-2",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.092072
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21"
          },
          "vNetId": "my03-mig-vnet-01",
          "cspVNetId": "tb5dn9aum2qip3phdv8n",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tbgi6cba5avpr77flia4",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tbdnf6sndttnf336lgim"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my05-infra101",
      "uid": "tbf56iebln7rh1kjspoe",
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
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "cm-migrated-at": "2026-11-25T09:16:32Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my05-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbf56iebln7rh1kjspoe"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-1",
          "uid": "tb9q3als393k50fpj5ps",
          "cspResourceName": "tb9q3als393k50fpj5ps",
          "cspResourceId": "ins-fxa70pv9",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:58",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-11-25 09:16:58",
            "sys.cspResourceId": "ins-fxa70pv9",
            "sys.cspResourceName": "tb9q3als393k50fpj5ps",
            "sys.id": "my05-ng-influxdb-back-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tb9q3als393k50fpj5ps",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-2"
          },
          "publicIP": "43.166.24.5",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.2",
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
          "specId": "tencent+ap-seoul+sa5.large16",
          "cspSpecName": "SA5.LARGE16",
          "spec": {
            "cspSpecName": "SA5.LARGE16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.22
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
          "cspVNetId": "vpc-p2o55rts",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-4kcu8oj9",
          "networkInterface": "eni-n9fgsyh0",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-b1dxxqvp"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-2",
          "uid": "tb1qavefj8k9l3vjdh1q",
          "cspResourceName": "tb1qavefj8k9l3vjdh1q",
          "cspResourceId": "ins-6t8n96er",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:58",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-11-25 09:16:58",
            "sys.cspResourceId": "ins-6t8n96er",
            "sys.cspResourceName": "tb1qavefj8k9l3vjdh1q",
            "sys.id": "my05-ng-influxdb-back-2",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tb1qavefj8k9l3vjdh1q",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-2"
          },
          "publicIP": "43.155.212.36",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
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
          "specId": "tencent+ap-seoul+sa5.large16",
          "cspSpecName": "SA5.LARGE16",
          "spec": {
            "cspSpecName": "SA5.LARGE16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.22
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
          "cspVNetId": "vpc-p2o55rts",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-4kcu8oj9",
          "networkInterface": "eni-ewr48fqa",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-b1dxxqvp"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-ip-10-0-1-30-62161624-1",
          "uid": "tb5qb43uh91ckbefruhm",
          "cspResourceName": "tb5qb43uh91ckbefruhm",
          "cspResourceId": "ins-px4j8x8x",
          "name": "my05-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my05-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:17:01",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-11-25 09:17:01",
            "sys.cspResourceId": "ins-px4j8x8x",
            "sys.cspResourceName": "tb5qb43uh91ckbefruhm",
            "sys.id": "my05-ng-ip-10-0-1-30-62161624-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-ip-10-0-1-30-62161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-ip-10-0-1-30-62161624",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tb5qb43uh91ckbefruhm",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-2"
          },
          "publicIP": "43.155.211.124",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.11",
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
          "cspVNetId": "vpc-p2o55rts",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-4kcu8oj9",
          "networkInterface": "eni-hs3bmg2o",
          "securityGroupIds": [
            "my05-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-b1dxxqvp"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my06-infra101",
      "uid": "tbb0t27stpupc73nksn6",
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
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "cm-migrated-at": "2026-11-25T09:17:45Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my06-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbb0t27stpupc73nksn6"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-1",
          "uid": "tbrec6bj4qkpctso5mci",
          "cspResourceName": "tbrec6bj4qkpctso5mci",
          "cspResourceId": "02h7_65a21f53-e4b2-4464-b89c-2d1ad1261b16",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:18:09",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-11-25 09:18:09",
            "sys.cspResourceId": "02h7_65a21f53-e4b2-4464-b89c-2d1ad1261b16",
            "sys.cspResourceName": "tbrec6bj4qkpctso5mci",
            "sys.id": "my06-ng-influxdb-back-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-influxdb-back",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbrec6bj4qkpctso5mci",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.97.204",
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
            "costPerHour": 0.247
          },
          "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-mig-vnet-01",
          "cspVNetId": "r026-07690ed0-796b-433a-b7ae-30a7ffa211a7",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-b9f42d6c-389f-44ee-9c18-82a38e5ce5a7",
          "networkInterface": "grafting-pavestone-plunder-broker",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-12c6f57d-43ad-4812-a953-07c0a581a0ec"
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-2",
          "uid": "tbklm8p51kb4oqdr8o4f",
          "cspResourceName": "tbklm8p51kb4oqdr8o4f",
          "cspResourceId": "02h7_ffe75cff-d1c8-40d1-822f-1ece581c6b69",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:18:08",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-11-25 09:18:08",
            "sys.cspResourceId": "02h7_ffe75cff-d1c8-40d1-822f-1ece581c6b69",
            "sys.cspResourceName": "tbklm8p51kb4oqdr8o4f",
            "sys.id": "my06-ng-influxdb-back-2",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-influxdb-back",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbklm8p51kb4oqdr8o4f",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.91.103",
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
            "costPerHour": 0.247
          },
          "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-mig-vnet-01",
          "cspVNetId": "r026-07690ed0-796b-433a-b7ae-30a7ffa211a7",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-b9f42d6c-389f-44ee-9c18-82a38e5ce5a7",
          "networkInterface": "ion-skimmer-ending-drip",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-12c6f57d-43ad-4812-a953-07c0a581a0ec"
        },
        {
          "resourceType": "node",
          "id": "my06-ng-ip-10-0-1-30-62161624-1",
          "uid": "tbo141372eh8ff0pdrif",
          "cspResourceName": "tbo141372eh8ff0pdrif",
          "cspResourceId": "02h7_1984afd7-a085-4a9e-82d2-e003457d44e3",
          "name": "my06-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my06-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:18:09",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-11-25 09:18:09",
            "sys.cspResourceId": "02h7_1984afd7-a085-4a9e-82d2-e003457d44e3",
            "sys.cspResourceName": "tbo141372eh8ff0pdrif",
            "sys.id": "my06-ng-ip-10-0-1-30-62161624-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-ip-10-0-1-30-62161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-ip-10-0-1-30-62161624",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbo141372eh8ff0pdrif",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.91.104",
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
            "costPerHour": 0.109
          },
          "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-mig-vnet-01",
          "cspVNetId": "r026-07690ed0-796b-433a-b7ae-30a7ffa211a7",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-b9f42d6c-389f-44ee-9c18-82a38e5ce5a7",
          "networkInterface": "zoology-attention-probational-stabilize",
          "securityGroupIds": [
            "my06-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-12c6f57d-43ad-4812-a953-07c0a581a0ec"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my08-infra101",
      "uid": "tbqa7keur5fguo4egmrp",
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
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "cm-migrated-at": "2026-11-25T09:18:44Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my08-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my08-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbqa7keur5fguo4egmrp"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my08-ng-influxdb-back-1",
          "uid": "tba0vpa8a9koncq8jbhl",
          "cspResourceName": "tba0vpa8a9koncq8jbhl",
          "cspResourceId": "147164275",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:19:57",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2026-11-25 09:19:57",
            "sys.cspResourceId": "147164275",
            "sys.cspResourceName": "tba0vpa8a9koncq8jbhl",
            "sys.id": "my08-ng-influxdb-back-1",
            "sys.infraId": "my08-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my08-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my08-ng-influxdb-back",
            "sys.subnetId": "my08-mig-subnet-01",
            "sys.uid": "tba0vpa8a9koncq8jbhl",
            "sys.vNetId": "my08-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "223.130.150.80",
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
          "specId": "ncp+kr+s4-g3",
          "cspSpecName": "s4-g3",
          "spec": {
            "cspSpecName": "s4-g3",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1745
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
          "cspVNetId": "153788",
          "subnetId": "my08-mig-subnet-01",
          "cspSubnetId": "336709",
          "networkInterface": "nic-6190936",
          "securityGroupIds": [
            "my08-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my08-mig-sshkey-01",
          "cspSshKeyId": "tbmo1vkhfnoa7ta0p33p"
        },
        {
          "resourceType": "node",
          "id": "my08-ng-influxdb-back-2",
          "uid": "tb5rf1thbdblvaeqlndp",
          "cspResourceName": "tb5rf1thbdblvaeqlndp",
          "cspResourceId": "147164269",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:20:05",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2026-11-25 09:20:05",
            "sys.cspResourceId": "147164269",
            "sys.cspResourceName": "tb5rf1thbdblvaeqlndp",
            "sys.id": "my08-ng-influxdb-back-2",
            "sys.infraId": "my08-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my08-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my08-ng-influxdb-back",
            "sys.subnetId": "my08-mig-subnet-01",
            "sys.uid": "tb5rf1thbdblvaeqlndp",
            "sys.vNetId": "my08-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "175.45.195.206",
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
            "costPerHour": 0.1745
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
          "cspVNetId": "153788",
          "subnetId": "my08-mig-subnet-01",
          "cspSubnetId": "336709",
          "networkInterface": "nic-6190934",
          "securityGroupIds": [
            "my08-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my08-mig-sshkey-01",
          "cspSshKeyId": "tbmo1vkhfnoa7ta0p33p"
        },
        {
          "resourceType": "node",
          "id": "my08-ng-ip-10-0-1-30-62161624-1",
          "uid": "tb78egghlm68bc17jcop",
          "cspResourceName": "tb78egghlm68bc17jcop",
          "cspResourceId": "147164272",
          "name": "my08-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my08-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:20:06",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2026-11-25 09:20:06",
            "sys.cspResourceId": "147164272",
            "sys.cspResourceName": "tb78egghlm68bc17jcop",
            "sys.id": "my08-ng-ip-10-0-1-30-62161624-1",
            "sys.infraId": "my08-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my08-ng-ip-10-0-1-30-62161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my08-ng-ip-10-0-1-30-62161624",
            "sys.subnetId": "my08-mig-subnet-01",
            "sys.uid": "tb78egghlm68bc17jcop",
            "sys.vNetId": "my08-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "223.130.159.128",
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
          "specId": "ncp+kr+c2-g3",
          "cspSpecName": "c2-g3",
          "spec": {
            "cspSpecName": "c2-g3",
            "vCPU": 2,
            "memoryGiB": 4,
            "costPerHour": 0.0677
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
          "cspVNetId": "153788",
          "subnetId": "my08-mig-subnet-01",
          "cspSubnetId": "336709",
          "networkInterface": "nic-6190935",
          "securityGroupIds": [
            "my08-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my08-mig-sshkey-01",
          "cspSshKeyId": "tbmo1vkhfnoa7ta0p33p"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my09-infra101",
      "uid": "tblq4eqit0en53dj5db4",
      "name": "my09-infra101",
      "status": "Creating:3 (R:0/3)",
      "statusCount": {
        "countTotal": 3,
        "countCreating": 3,
        "countRunning": 0,
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
      "targetStatus": "Running",
      "targetAction": "Create",
      "configureCloudAdaptiveNetwork": "",
      "label": null,
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my09-ng-influxdb-back-1",
          "uid": "tbl514uefdglqljjbjua",
          "name": "my09-ng-influxdb-back-1",
          "nodeGroupId": "my09-ng-influxdb-back",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "nhn-kr1",
          "connectionConfig": {
            "configName": "nhn-kr1",
            "providerName": "nhn",
            "driverName": "nhn-driver-v1.0.so",
            "credentialName": "nhn",
            "credentialHolder": "admin",
            "regionZoneInfoName": "nhn-kr1",
            "regionZoneInfo": {
              "assignedRegion": "KR1",
              "assignedZone": "kr-pub-a"
            },
            "regionDetail": {
              "regionId": "KR1",
              "regionName": "kr1",
              "description": "Pangyo (South Korea)",
              "location": {
                "display": "Pangyo (South Korea)",
                "latitude": 37.390889,
                "longitude": 127.096792
              },
              "zones": [
                "kr-pub-a",
                "kr-pub-b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "nhn+kr1+r2.c4m16",
          "cspSpecName": "r2.c4m16",
          "spec": {
            "cspSpecName": "r2.c4m16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.19535999
          },
          "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "image": {
            "resourceType": "image",
            "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
          },
          "vNetId": "my09-mig-vnet-01",
          "cspVNetId": "",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my09-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my09-ng-influxdb-back-2",
          "uid": "tbapjacflov280qo67i6",
          "name": "my09-ng-influxdb-back-2",
          "nodeGroupId": "my09-ng-influxdb-back",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "nhn-kr1",
          "connectionConfig": {
            "configName": "nhn-kr1",
            "providerName": "nhn",
            "driverName": "nhn-driver-v1.0.so",
            "credentialName": "nhn",
            "credentialHolder": "admin",
            "regionZoneInfoName": "nhn-kr1",
            "regionZoneInfo": {
              "assignedRegion": "KR1",
              "assignedZone": "kr-pub-a"
            },
            "regionDetail": {
              "regionId": "KR1",
              "regionName": "kr1",
              "description": "Pangyo (South Korea)",
              "location": {
                "display": "Pangyo (South Korea)",
                "latitude": 37.390889,
                "longitude": 127.096792
              },
              "zones": [
                "kr-pub-a",
                "kr-pub-b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "nhn+kr1+r2.c4m16",
          "cspSpecName": "r2.c4m16",
          "spec": {
            "cspSpecName": "r2.c4m16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.19535999
          },
          "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "image": {
            "resourceType": "image",
            "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
          },
          "vNetId": "my09-mig-vnet-01",
          "cspVNetId": "",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my09-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my09-ng-ip-10-0-1-30-62161624-1",
          "uid": "tbvu2ueg9vml0gan1d54",
          "name": "my09-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my09-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
          "rootDiskSize": 20,
          "rootDeviceName": "",
          "connectionName": "nhn-kr1",
          "connectionConfig": {
            "configName": "nhn-kr1",
            "providerName": "nhn",
            "driverName": "nhn-driver-v1.0.so",
            "credentialName": "nhn",
            "credentialHolder": "admin",
            "regionZoneInfoName": "nhn-kr1",
            "regionZoneInfo": {
              "assignedRegion": "KR1",
              "assignedZone": "kr-pub-a"
            },
            "regionDetail": {
              "regionId": "KR1",
              "regionName": "kr1",
              "description": "Pangyo (South Korea)",
              "location": {
                "display": "Pangyo (South Korea)",
                "latitude": 37.390889,
                "longitude": 127.096792
              },
              "zones": [
                "kr-pub-a",
                "kr-pub-b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "nhn+kr1+c2.c2m2",
          "cspSpecName": "c2.c2m2",
          "spec": {
            "cspSpecName": "c2.c2m2",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.042179998
          },
          "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "image": {
            "resourceType": "image",
            "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
          },
          "vNetId": "my09-mig-vnet-01",
          "cspVNetId": "",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my09-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": ""
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
- **IDs**: [my01-infra101 my02-infra101 my03-infra101 my05-infra101 my06-infra101 my08-infra101 my09-infra101]

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
  "uid": "tbqa7keur5fguo4egmrp",
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
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "cm-migrated-at": "2026-11-25T09:18:44Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my08-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my08-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbqa7keur5fguo4egmrp"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my08-ng-influxdb-back-1",
      "uid": "tba0vpa8a9koncq8jbhl",
      "cspResourceName": "tba0vpa8a9koncq8jbhl",
      "cspResourceId": "147164275",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:19:57",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "223.130.150.80",
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
      "specId": "ncp+kr+s4-g3",
      "cspSpecName": "s4-g3",
      "spec": {
        "cspSpecName": "s4-g3",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1745
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
      "cspVNetId": "153788",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "336709",
      "networkInterface": "nic-6190936",
      "securityGroupIds": [
        "my08-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tbmo1vkhfnoa7ta0p33p",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-11-25T09:20:07Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:06Z",
          "completedTime": "2026-11-25T09:20:09Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my08-infra101-tb2bpgqmq5j8mvp21njh",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:09Z",
          "completedTime": "2026-11-25T09:20:11Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tba0vpa8a9koncq8jbhl 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my08-ng-influxdb-back-2",
      "uid": "tb5rf1thbdblvaeqlndp",
      "cspResourceName": "tb5rf1thbdblvaeqlndp",
      "cspResourceId": "147164269",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:20:05",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "175.45.195.206",
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
        "costPerHour": 0.1745
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
      "cspVNetId": "153788",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "336709",
      "networkInterface": "nic-6190934",
      "securityGroupIds": [
        "my08-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tbmo1vkhfnoa7ta0p33p",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-11-25T09:20:07Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:06Z",
          "completedTime": "2026-11-25T09:20:08Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my08-infra101-tb2bpgqmq5j8mvp21njh",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:09Z",
          "completedTime": "2026-11-25T09:20:10Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb5rf1thbdblvaeqlndp 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my08-ng-ip-10-0-1-30-62161624-1",
      "uid": "tb78egghlm68bc17jcop",
      "cspResourceName": "tb78egghlm68bc17jcop",
      "cspResourceId": "147164272",
      "name": "my08-ng-ip-10-0-1-30-62161624-1",
      "nodeGroupId": "my08-ng-ip-10-0-1-30-62161624",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:20:06",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=66.7%",
      "region": {
        "region": "KR",
        "zone": "KR-1"
      },
      "publicIP": "223.130.159.128",
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
      "specId": "ncp+kr+c2-g3",
      "cspSpecName": "c2-g3",
      "spec": {
        "cspSpecName": "c2-g3",
        "vCPU": 2,
        "memoryGiB": 4,
        "costPerHour": 0.0677
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
      "cspVNetId": "153788",
      "subnetId": "my08-mig-subnet-01",
      "cspSubnetId": "336709",
      "networkInterface": "nic-6190935",
      "securityGroupIds": [
        "my08-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my08-mig-sshkey-01",
      "cspSshKeyId": "tbmo1vkhfnoa7ta0p33p",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHWlg+tXjkjnBfk0KfXIZnsAnFZeTiRoXl6rO1F9LAx0QmlAhB6H9S23i9Ye/A9ACdUpS363A6UvMT7MLqDkFNU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2WSq7XAjDTgX/DGvHQOWkvXaevoO2kb8tl2cApRpCVE",
        "firstUsedAt": "2026-11-25T09:20:07Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:06Z",
          "completedTime": "2026-11-25T09:20:09Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my08-infra101-tb2bpgqmq5j8mvp21njh",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:20:09Z",
          "completedTime": "2026-11-25T09:20:11Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb78egghlm68bc17jcop 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
        "my08-ng-influxdb-back",
        "my08-ng-ip-10-0-1-30-62161624"
      ],
      "nodeIds": [
        "my08-ng-influxdb-back-1",
        "my08-ng-influxdb-back-2",
        "my08-ng-ip-10-0-1-30-62161624-1"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my08-ng-influxdb-back",
      "representativeNodeId": "my08-ng-influxdb-back-1"
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
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-influxdb-back-2",
            "nodeIp": "175.45.195.206",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb5rf1thbdblvaeqlndp 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-ip-10-0-1-30-62161624-1",
            "nodeIp": "223.130.159.128",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb78egghlm68bc17jcop 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my08-infra101",
            "nodeId": "my08-ng-influxdb-back-1",
            "nodeIp": "223.130.150.80",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tba0vpa8a9koncq8jbhl 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postScriptRequestId": "ps-my08-infra101-tb2bpgqmq5j8mvp21njh"
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
      "id": "my08-ng-influxdb-back",
      "cspResourceName": "tb4eeq1ff30hug6d5pqv",
      "cspResourceId": "147164402",
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
        "ip": "211.233.218.78",
        "port": "9999",
        "keyValueList": [
          {
            "key": "LoadBalancerInstanceNo",
            "value": "147164402"
          },
          {
            "key": "LoadBalancerListenerNo",
            "value": "757965"
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
            "value": "837925"
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
      "createdTime": "2026-11-25T18:20:52+09:00",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "LoadBalancerInstanceNo",
          "value": "147164402"
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
          "value": "2026-11-25T18:20:52+0900"
        },
        {
          "key": "LoadBalancerName",
          "value": "tb4eeq1ff30hug6d5pqv"
        },
        {
          "key": "LoadBalancerDomain",
          "value": "tb4eeq1ff30hug6d5pqv-147164402-8640f5437342.kr.lb.naverncp.com"
        },
        {
          "key": "LoadBalancerIpList",
          "value": "211.233.218.78"
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
          "value": "153788"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "SubnetNoList",
          "value": "336710"
        },
        {
          "key": "LoadBalancerSubnetList",
          "value": "{zoneCode:KR-1,subnetNo:336710,publicIpInstanceNo:147164410}"
        },
        {
          "key": "LoadBalancerListenerNoList",
          "value": "757965"
        },
        {
          "key": "ListenerId",
          "value": "757965"
        },
        {
          "key": "HealthCheckerId",
          "value": "6648179"
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
    "cspResourceName": "tb4eeq1ff30hug6d5pqv",
    "cspResourceId": "147164402",
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
      "ip": "211.233.218.78",
      "port": "9999",
      "keyValueList": [
        {
          "key": "LoadBalancerInstanceNo",
          "value": "147164402"
        },
        {
          "key": "LoadBalancerListenerNo",
          "value": "757965"
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
          "value": "837925"
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
    "createdTime": "2026-11-25T18:20:52+09:00",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "LoadBalancerInstanceNo",
        "value": "147164402"
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
        "value": "2026-11-25T18:20:52+0900"
      },
      {
        "key": "LoadBalancerName",
        "value": "tb4eeq1ff30hug6d5pqv"
      },
      {
        "key": "LoadBalancerDomain",
        "value": "tb4eeq1ff30hug6d5pqv-147164402-8640f5437342.kr.lb.naverncp.com"
      },
      {
        "key": "LoadBalancerIpList",
        "value": "211.233.218.78"
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
        "value": "153788"
      },
      {
        "key": "RegionCode",
        "value": "KR"
      },
      {
        "key": "SubnetNoList",
        "value": "336710"
      },
      {
        "key": "LoadBalancerSubnetList",
        "value": "{zoneCode:KR-1,subnetNo:336710,publicIpInstanceNo:147164410}"
      },
      {
        "key": "LoadBalancerListenerNoList",
        "value": "757965"
      },
      {
        "key": "ListenerId",
        "value": "757965"
      },
      {
        "key": "HealthCheckerId",
        "value": "6648179"
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
  "cspResourceName": "tb4eeq1ff30hug6d5pqv",
  "cspResourceId": "147164402",
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
    "ip": "211.233.218.78",
    "port": "9999",
    "keyValueList": [
      {
        "key": "LoadBalancerInstanceNo",
        "value": "147164402"
      },
      {
        "key": "LoadBalancerListenerNo",
        "value": "757965"
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
        "value": "837925"
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
  "createdTime": "2026-11-25T18:20:52+09:00",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "LoadBalancerInstanceNo",
      "value": "147164402"
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
      "value": "2026-11-25T18:20:52+0900"
    },
    {
      "key": "LoadBalancerName",
      "value": "tb4eeq1ff30hug6d5pqv"
    },
    {
      "key": "LoadBalancerDomain",
      "value": "tb4eeq1ff30hug6d5pqv-147164402-8640f5437342.kr.lb.naverncp.com"
    },
    {
      "key": "LoadBalancerIpList",
      "value": "211.233.218.78"
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
      "value": "153788"
    },
    {
      "key": "RegionCode",
      "value": "KR"
    },
    {
      "key": "SubnetNoList",
      "value": "336710"
    },
    {
      "key": "LoadBalancerSubnetList",
      "value": "{zoneCode:KR-1,subnetNo:336710,publicIpInstanceNo:147164410}"
    },
    {
      "key": "LoadBalancerListenerNoList",
      "value": "757965"
    },
    {
      "key": "ListenerId",
      "value": "757965"
    },
    {
      "key": "HealthCheckerId",
      "value": "6648179"
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

**Summary**: Load balancing confirmed: traffic distributed across 2 unique VMs

### Test Case 12: Target Infrastructure Summary

#### 12.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Target Infrastructure Summary**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:23:58

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
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| s4-g3 | 4 | 16.0 | - | x86_64 | default | $0.1745 | 2 |
| c2-g3 | 2 | 4.0 | - | x86_64 | default | $0.0677 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| 23214590 | ubuntu-22.04-base (Hypervisor:KVM) | Ubuntu 22.04 | Linux/UNIX | x86_64 | Common BlockStorage 1 | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my08-ng-influxdb-back-1 | 147164275 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 223.130.150.80<br>**Private IP:** 10.0.1.8<br>**SGs:** my08-mig-sg-01<br>**SSH:** my08-mig-sshkey-01 |
| my08-ng-influxdb-back-2 | 147164269 | Running | 4 vCPU, 16.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 175.45.195.206<br>**Private IP:** 10.0.1.6<br>**SGs:** my08-mig-sg-01<br>**SSH:** my08-mig-sshkey-01 |
| my08-ng-ip-10-0-1-30-62161624-1 | 147164272 | Running | 2 vCPU, 4.0 GiB | ubuntu-22.04-base (Hypervisor:KVM) (ubuntu-22.04-base (Hypervisor:KVM)) | **VNet:** my08-mig-vnet-01<br>**Subnet:** my08-mig-subnet-01<br>**Public IP:** 223.130.159.128<br>**Private IP:** 10.0.1.7<br>**SGs:** my08-mig-sg-02<br>**SSH:** my08-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my08-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my08-mig-vnet-01 |
| **CSP VNet ID** | 153788 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ncp-kr |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my08-mig-subnet-01 | 336709 | 10.0.1.0/24 | KR-1 |

### Network Load Balancers (NLB)

#### NLB: my08-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my08-ng-influxdb-back |
| **CSP NLB Name** | tb4eeq1ff30hug6d5pqv |
| **CSP NLB ID** | 147164402 |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | ncp-kr |
| **Listener** | TCP:9999 (IP: 211.233.218.78) |
| **Target Group** | TCP:8086 (Node Group: my08-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 0s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my08-ng-influxdb-back-1 | 147164275 | 10.0.1.8 | 223.130.150.80 | Running |
| my08-ng-influxdb-back-2 | 147164269 | 10.0.1.6 | 175.45.195.206 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my08-mig-sshkey-01 | tbmo1vkhfnoa7ta0p33p | cb-user | 2a:8f:c2:e3:53:12:d9:33:23:0b:fc:24:37:da:ff:85 |

### Security Groups

#### Security Group: my08-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my08-mig-sg-01 |
| **CSP Security Group ID** | 423537 |
| **VNet** | my08-mig-vnet-01 |
| **Rule Count** | 10 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 1-65535 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my08-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my08-mig-sg-02 |
| **CSP Security Group ID** | 423538 |
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
| **Per Hour** | $0.4167 |
| **Per Day** | $10.00 |
| **Per Month (30 days)** | $300.02 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| NCP | KR | 3 | $0.4167 | $300.02 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my08-ng-influxdb-back-1 | s4-g3 | $0.1745 | $125.64 |
| my08-ng-influxdb-back-2 | s4-g3 | $0.1745 | $125.64 |
| my08-ng-ip-10-0-1-30-62161624-1 | c2-g3 | $0.0677 | $48.74 |




### Test Case 13: Migration Report

#### 13.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 13.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-11-25 09:24:03*

---

## 📊 Migration Summary

**Target Cloud:** NCP

**Target Region:** KR

**Namespace:** mig01 | **Infra ID:** my08-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**NLB Migration Status:** Completed

**Total Load Balancers:** 1

**Migrated Load Balancers:** 1

**💰 Estimated Monthly Cost:** $300.02 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | s4-g3, c2-g3 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my08-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my08-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 18 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |
| 8 | **Network Load Balancer** | 1 | ✅ Created | my08-ng-influxdb-back (TCP:9999 -> TCP:8086) |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my08-ng-influxdb-back-1<br>**VM ID:** 147164275<br>**Label(sourceMachineId):** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 |
| 2 | **VM Name:** my08-ng-influxdb-back-2<br>**VM ID:** 147164269<br>**Label(sourceMachineId):** ec288dd0-c6fa-8a49-2f60-bc898311febf | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf |
| 3 | **VM Name:** my08-ng-ip-10-0-1-30-62161624-1<br>**VM ID:** 147164272<br>**Label(sourceMachineId):** ec268ed7-821e-9d73-e79f-961262161624 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my08-ng-influxdb-back-1 | **Spec ID:** s4-g3<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 30 GB |
| 2 | my08-ng-influxdb-back-2 | **Spec ID:** s4-g3<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 8 GB<br>**Root Disk:** 30 GB |
| 3 | my08-ng-ip-10-0-1-30-62161624-1 | **Spec ID:** c2-g3<br>**vCPUs:** 2<br>**Memory:** 4.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 8 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my08-ng-influxdb-back-1 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | my08-ng-influxdb-back-2 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | my08-ng-ip-10-0-1-30-62161624-1 | **Image ID:** 23214590<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu-22.04-base (Hypervisor:KVM) | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 18 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my08-mig-sg-01

**CSP ID:** 423537 | **VNet:** my08-mig-vnet-01 | **Rules:** 10

**Assigned VMs:**

- **VM:** my08-ng-influxdb-back-1
  - **Source Server:** **Hostname:** ip-10-0-1-221, **Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932
- **VM:** my08-ng-influxdb-back-2
  - **Source Server:** **Hostname:** ip-10-0-1-138, **Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 2 | inbound | TCP | 9999 | 0.0.0.0/0 | - | Created by system |
| 3 | inbound | ICMP |  | 0.0.0.0/0 | - | Created by system |
| 4 | inbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 5 | inbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 6 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 7 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 8 | outbound | ICMP |  | 0.0.0.0/0 | - | Created by system |
| 9 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 10 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |

### Security Group: my08-mig-sg-02

**CSP ID:** 423538 | **VNet:** my08-mig-vnet-01 | **Rules:** 8

**Assigned VMs:**

- **VM:** my08-ng-ip-10-0-1-30-62161624-1
  - **Source Server:** **Hostname:** ip-10-0-1-30, **Machine ID:** ec268ed7-821e-9d73-e79f-961262161624

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
| 1 | **Name:** my08-mig-vnet-01<br>**ID:** 153788 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my08-mig-subnet-01<br>**ID:** 336709 | 10.0.1.0/24 | my08-mig-vnet-01 |

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

### 1. influxdb_back (haproxy) → my08-ng-influxdb-back

**Configuration Comparison:**

| Property | Source (On-Premise) | Target (Cloud NLB) | Note |
|----------|---------------------|--------------------|------|
| **Load Balancer / Name** | haproxy (Host: ip-10-0-1-30) | my08-ng-influxdb-back (CSP ID: tb4eeq1ff30hug6d5pqv) | Replaced by managed NLB |
| **Type / Scope** | Software LB | PUBLIC / REGION | Cloud-native NLB |
| **Frontend Listener** | *:9999 (TCP) | TCP:9999 | IP: 211.233.218.78 |
| **Backend / Target Group** | influxdb_back (TCP:8086, roundrobin) | my08-ng-influxdb-back (TCP:8086, 2 nodes) | Target Group attached to NodeGroup |
| **Health Check** | Enabled (Interval: 10s, Timeout: 10s, Threshold: 3) | TCP:8086 (Interval: 10s, Timeout: 0s, Threshold: 3) | Preserved |

**Backend Node Mappings:**

| No. | Source Backend Server | Source Endpoint | Target Node | Target CSP VM ID | Target Private IP | Target Public IP | Target Port | Status |
|-----|-----------------------|-----------------|-------------|------------------|-------------------|------------------|-------------|--------|
| 1 | influx1 (ip-10-0-1-221) | 10.0.1.221:8086 | my08-ng-influxdb-back-1 | 147164275 | 10.0.1.8 | 223.130.150.80 | 8086 | Running |
| 2 | influx2 (ip-10-0-1-138) | 10.0.1.138:8086 | my08-ng-influxdb-back-2 | 147164269 | 10.0.1.6 | 175.45.195.206 | 8086 | Running |

**Migration Notes:**
- On-premise haproxy on ip-10-0-1-30 (migrated to my08-ng-ip-10-0-1-30-62161624-1) is replaced by cloud-managed NLB.

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my08-mig-sshkey-01 | tbmo1vkhfnoa7ta0p33p | 2a:8f:c2:e3:53:12:d9:33:23:0b:fc:24:37:da:ff:85 | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.4167 |
| Daily | $10.00 |
| Monthly | $300.02 |
| Yearly | $3600.29 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | c2-g3 | $48.74 | 16.2% |
| ip-10-0-1-221 (migrated) | s4-g3 | $125.64 | 41.9% |
| ip-10-0-1-138 (migrated) | s4-g3 | $125.64 | 41.9% |

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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my08-infra101)",
  "success": true
}
```

