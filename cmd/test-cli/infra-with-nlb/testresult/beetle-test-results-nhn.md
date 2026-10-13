# CM-Beetle test results for NHN (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with NHN cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.1+ (3649ed7)
- imdl: v0.1.16+ (3649ed7)
- CB-Tumblebug: unknown
- CB-Spider: unknown
- CB-MapUI: v0.13.13
- Target CSP: NHN
- Target Region: kr1
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: October 12, 2026
- Test Time: 19:54:35 KST
- Test Execution: 2026-10-12 19:54:35 KST

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

## Test result for NHN

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 4.632s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 1.895s | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 2m53.568s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 49ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 9ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 16ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 5.912s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 1m1.058s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 7ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 8ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 1m33.46s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.311s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.249s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 22.293s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m45.152s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 8m49.210237302s

*Test executed on October 12, 2026 at 19:54:35 KST (2026-10-12 19:54:35 KST) using CM-Beetle automated test CLI*

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
  "desiredCsp": "nhn",
  "desiredRegion": "kr1",
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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=60.0% Max=100.0% Avg=86.7% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "nhn",
        "region": "kr1"
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
            "connectionName": "nhn-kr1",
            "specId": "nhn+kr1+r2.c4m16",
            "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
            "connectionName": "nhn-kr1",
            "specId": "nhn+kr1+c2.c2m2",
            "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 20,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "mig-vnet-01",
        "connectionName": "nhn-kr1",
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
        "connectionName": "nhn-kr1",
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
          "id": "nhn+kr1+r2.c4m16",
          "uid": "tb9ngrgm4pd16v667jgg",
          "cspSpecName": "r2.c4m16",
          "name": "nhn+kr1+r2.c4m16",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "costPerHour": 0.19535999,
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
              "key": "ID",
              "value": "8339613a-2837-4c72-906b-a8a47aa17c3f"
            },
            {
              "key": "Name",
              "value": "r2.c4m16"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/8339613a-2837-4c72-906b-a8a47aa17c3f,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/8339613a-2837-4c72-906b-a8a47aa17c3f,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "16384"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "4"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:memory}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        },
        {
          "id": "nhn+kr1+c2.c2m2",
          "uid": "tbcqsc0oqtknhnnqaunv",
          "cspSpecName": "c2.c2m2",
          "name": "nhn+kr1+c2.c2m2",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "costPerHour": 0.042179998,
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
              "key": "ID",
              "value": "f5ebd3a7-f978-477d-8dcf-2a1578e29df3"
            },
            {
              "key": "Name",
              "value": "c2.c2m2"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/f5ebd3a7-f978-477d-8dcf-2a1578e29df3,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/f5ebd3a7-f978-477d-8dcf-2a1578e29df3,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "2048"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "2"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:core}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "nhn",
          "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "regionList": [
            "kr1"
          ],
          "id": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "uid": "tb76dofaj88ahqu4op2v",
          "name": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "nhn-kr1",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:30 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": true,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ID",
              "value": "0f07c795-2a46-44fc-a61b-fa0d96763ce2"
            },
            {
              "key": "Name",
              "value": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
            },
            {
              "key": "Status",
              "value": "active"
            },
            {
              "key": "Tags",
              "value": "BASE"
            },
            {
              "key": "ContainerFormat",
              "value": "bare"
            },
            {
              "key": "DiskFormat",
              "value": "qcow2"
            },
            {
              "key": "MinDiskGigabytes",
              "value": "20"
            },
            {
              "key": "MinRAMMegabytes",
              "value": "0"
            },
            {
              "key": "Owner",
              "value": "c289b99209ca4e189095cdecebbd092d"
            },
            {
              "key": "Protected",
              "value": "true"
            },
            {
              "key": "Visibility",
              "value": "public"
            },
            {
              "key": "Hidden",
              "value": "false"
            },
            {
              "key": "Checksum",
              "value": "4c563c67af5d73b15e7003af724d6b55"
            },
            {
              "key": "Properties",
              "value": "{deprecate_date:null,description:Ubuntu Server 22.04.5 LTS (2026.03.10),hw_cpu_sockets:1,hw_qemu_guest_agent:yes,hw_vif_multiqueue_enabled:true,login_username:ubuntu,max_cpu:,min_cpu:0,monitoring_agent:sysmon,nhncloud_allow_autoscale:true,nhncloud_allow_cgroup:v2,nhncloud_allow_compute_flavor:true,nhncloud_allow_cpu_flavor:true,nhncloud_allow_download:false,nhncloud_allow_gpu_flavor:false,nhncloud_allow_image_create:true,nhncloud_allow_imagebuilder:true,nhncloud_allow_instance_template:true,nhncloud_allow_local_bootdisk_flavor:true,nhncloud_allow_nks_cpu_flavor:false,nhncloud_allow_nks_gpu_flavor:false,nhncloud_allow_user_script:true,nhncloud_category:OS,nhncloud_product:compute,os_architecture:amd64,os_distro:ubuntu,os_type:linux,os_version:Server 22.04 LTS,project_domain:WDI;NORMAL,release_date:2026.03.10,tc_env:cloudmon,sysmon}"
            },
            {
              "key": "CreatedAt",
              "value": "2026-03-09T01:34:35Z"
            },
            {
              "key": "UpdatedAt",
              "value": "2026-03-09T21:14:13Z"
            },
            {
              "key": "File",
              "value": "/v2/images/0f07c795-2a46-44fc-a61b-fa0d96763ce2/file"
            },
            {
              "key": "Schema",
              "value": "/v2/schemas/image"
            },
            {
              "key": "VirtualSize",
              "value": "0"
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
          "connectionName": "nhn-kr1",
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
            },
            {
              "Ports": "8086",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/21"
            },
            {
              "Ports": "8086",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.1.0/24"
            }
          ],
          "cspResourceId": ""
        },
        {
          "name": "mig-sg-02",
          "connectionName": "nhn-kr1",
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=70.0% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "nhn",
        "region": "kr1"
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=50.0% Memory=100.0% Image=60.0%",
            "connectionName": "nhn-kr1",
            "specId": "nhn+kr1+m2.c8m16",
            "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=60.0%",
            "connectionName": "nhn-kr1",
            "specId": "nhn+kr1+m2.c2m4",
            "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 20,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "mig-vnet-01",
        "connectionName": "nhn-kr1",
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
        "connectionName": "nhn-kr1",
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
          "id": "nhn+kr1+m2.c8m16",
          "uid": "tbstcl8ocf1p2kvit3nv",
          "cspSpecName": "m2.c8m16",
          "name": "nhn+kr1+m2.c8m16",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 8,
          "memoryGiB": 16,
          "costPerHour": 0.29525998,
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
              "key": "ID",
              "value": "eaddf5bd-2725-409a-ac57-044935747183"
            },
            {
              "key": "Name",
              "value": "m2.c8m16"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/eaddf5bd-2725-409a-ac57-044935747183,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/eaddf5bd-2725-409a-ac57-044935747183,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "16384"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "8"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:general}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        },
        {
          "id": "nhn+kr1+m2.c2m4",
          "uid": "tbpnvei7jqqgj4d9678f",
          "cspSpecName": "m2.c2m4",
          "name": "nhn+kr1+m2.c2m4",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 4,
          "costPerHour": 0.06882,
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
              "key": "ID",
              "value": "35a73b57-58a7-434d-aa08-5249aaa95b3e"
            },
            {
              "key": "Name",
              "value": "m2.c2m4"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/35a73b57-58a7-434d-aa08-5249aaa95b3e,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/35a73b57-58a7-434d-aa08-5249aaa95b3e,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "4096"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "2"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:general}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "nhn",
          "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "regionList": [
            "kr1"
          ],
          "id": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "uid": "tb76dofaj88ahqu4op2v",
          "name": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "nhn-kr1",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:30 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": true,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ID",
              "value": "0f07c795-2a46-44fc-a61b-fa0d96763ce2"
            },
            {
              "key": "Name",
              "value": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
            },
            {
              "key": "Status",
              "value": "active"
            },
            {
              "key": "Tags",
              "value": "BASE"
            },
            {
              "key": "ContainerFormat",
              "value": "bare"
            },
            {
              "key": "DiskFormat",
              "value": "qcow2"
            },
            {
              "key": "MinDiskGigabytes",
              "value": "20"
            },
            {
              "key": "MinRAMMegabytes",
              "value": "0"
            },
            {
              "key": "Owner",
              "value": "c289b99209ca4e189095cdecebbd092d"
            },
            {
              "key": "Protected",
              "value": "true"
            },
            {
              "key": "Visibility",
              "value": "public"
            },
            {
              "key": "Hidden",
              "value": "false"
            },
            {
              "key": "Checksum",
              "value": "4c563c67af5d73b15e7003af724d6b55"
            },
            {
              "key": "Properties",
              "value": "{deprecate_date:null,description:Ubuntu Server 22.04.5 LTS (2026.03.10),hw_cpu_sockets:1,hw_qemu_guest_agent:yes,hw_vif_multiqueue_enabled:true,login_username:ubuntu,max_cpu:,min_cpu:0,monitoring_agent:sysmon,nhncloud_allow_autoscale:true,nhncloud_allow_cgroup:v2,nhncloud_allow_compute_flavor:true,nhncloud_allow_cpu_flavor:true,nhncloud_allow_download:false,nhncloud_allow_gpu_flavor:false,nhncloud_allow_image_create:true,nhncloud_allow_imagebuilder:true,nhncloud_allow_instance_template:true,nhncloud_allow_local_bootdisk_flavor:true,nhncloud_allow_nks_cpu_flavor:false,nhncloud_allow_nks_gpu_flavor:false,nhncloud_allow_user_script:true,nhncloud_category:OS,nhncloud_product:compute,os_architecture:amd64,os_distro:ubuntu,os_type:linux,os_version:Server 22.04 LTS,project_domain:WDI;NORMAL,release_date:2026.03.10,tc_env:cloudmon,sysmon}"
            },
            {
              "key": "CreatedAt",
              "value": "2026-03-09T01:34:35Z"
            },
            {
              "key": "UpdatedAt",
              "value": "2026-03-09T21:14:13Z"
            },
            {
              "key": "File",
              "value": "/v2/images/0f07c795-2a46-44fc-a61b-fa0d96763ce2/file"
            },
            {
              "key": "Schema",
              "value": "/v2/schemas/image"
            },
            {
              "key": "VirtualSize",
              "value": "0"
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
          "connectionName": "nhn-kr1",
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
            },
            {
              "Ports": "8086",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/21"
            },
            {
              "Ports": "8086",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.1.0/24"
            }
          ],
          "cspResourceId": ""
        },
        {
          "name": "mig-sg-02",
          "connectionName": "nhn-kr1",
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
  "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=60.0% Max=100.0% Avg=86.7% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
  "targetCloud": {
    "csp": "nhn",
    "region": "kr1"
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
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
        "connectionName": "nhn-kr1",
        "specId": "nhn+kr1+r2.c4m16",
        "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
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
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
        "connectionName": "nhn-kr1",
        "specId": "nhn+kr1+c2.c2m2",
        "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
        "vNetId": "mig-vnet-01",
        "subnetId": "mig-subnet-01",
        "securityGroupIds": [
          "mig-sg-02"
        ],
        "sshKeyId": "mig-sshkey-01",
        "rootDiskSize": 20,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "mig-vnet-01",
    "connectionName": "nhn-kr1",
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
    "connectionName": "nhn-kr1",
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
      "id": "nhn+kr1+r2.c4m16",
      "uid": "tb9ngrgm4pd16v667jgg",
      "cspSpecName": "r2.c4m16",
      "name": "nhn+kr1+r2.c4m16",
      "namespace": "system",
      "connectionName": "nhn-kr1",
      "providerName": "nhn",
      "regionName": "kr1",
      "regionLatitude": 37.390889,
      "regionLongitude": 127.096792,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 16,
      "costPerHour": 0.19535999,
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
          "key": "ID",
          "value": "8339613a-2837-4c72-906b-a8a47aa17c3f"
        },
        {
          "key": "Name",
          "value": "r2.c4m16"
        },
        {
          "key": "Links",
          "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/8339613a-2837-4c72-906b-a8a47aa17c3f,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/8339613a-2837-4c72-906b-a8a47aa17c3f,rel:bookmark}"
        },
        {
          "key": "RAM",
          "value": "16384"
        },
        {
          "key": "Disabled",
          "value": "false"
        },
        {
          "key": "VCPUs",
          "value": "4"
        },
        {
          "key": "ExtraSpecs",
          "value": "{flavor_type:memory}"
        },
        {
          "key": "IsPublic",
          "value": "true"
        },
        {
          "key": "RxTxFactor",
          "value": "1.00"
        },
        {
          "key": "Ephemeral",
          "value": "0"
        },
        {
          "key": "Disk",
          "value": "0"
        },
        {
          "key": "Notice!!",
          "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
        }
      ]
    },
    {
      "id": "nhn+kr1+c2.c2m2",
      "uid": "tbcqsc0oqtknhnnqaunv",
      "cspSpecName": "c2.c2m2",
      "name": "nhn+kr1+c2.c2m2",
      "namespace": "system",
      "connectionName": "nhn-kr1",
      "providerName": "nhn",
      "regionName": "kr1",
      "regionLatitude": 37.390889,
      "regionLongitude": 127.096792,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 2,
      "costPerHour": 0.042179998,
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
          "key": "ID",
          "value": "f5ebd3a7-f978-477d-8dcf-2a1578e29df3"
        },
        {
          "key": "Name",
          "value": "c2.c2m2"
        },
        {
          "key": "Links",
          "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/f5ebd3a7-f978-477d-8dcf-2a1578e29df3,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/f5ebd3a7-f978-477d-8dcf-2a1578e29df3,rel:bookmark}"
        },
        {
          "key": "RAM",
          "value": "2048"
        },
        {
          "key": "Disabled",
          "value": "false"
        },
        {
          "key": "VCPUs",
          "value": "2"
        },
        {
          "key": "ExtraSpecs",
          "value": "{flavor_type:core}"
        },
        {
          "key": "IsPublic",
          "value": "true"
        },
        {
          "key": "RxTxFactor",
          "value": "1.00"
        },
        {
          "key": "Ephemeral",
          "value": "0"
        },
        {
          "key": "Disk",
          "value": "0"
        },
        {
          "key": "Notice!!",
          "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
        }
      ]
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "nhn",
      "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
      "regionList": [
        "kr1"
      ],
      "id": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
      "uid": "tb76dofaj88ahqu4op2v",
      "name": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "nhn-kr1",
      "infraType": "",
      "fetchedTime": "2026.08.21 13:57:30 Fri",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": true,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)",
      "osDiskType": "NA",
      "osDiskSizeGB": 20,
      "imageStatus": "Available",
      "details": [
        {
          "key": "ID",
          "value": "0f07c795-2a46-44fc-a61b-fa0d96763ce2"
        },
        {
          "key": "Name",
          "value": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
        },
        {
          "key": "Status",
          "value": "active"
        },
        {
          "key": "Tags",
          "value": "BASE"
        },
        {
          "key": "ContainerFormat",
          "value": "bare"
        },
        {
          "key": "DiskFormat",
          "value": "qcow2"
        },
        {
          "key": "MinDiskGigabytes",
          "value": "20"
        },
        {
          "key": "MinRAMMegabytes",
          "value": "0"
        },
        {
          "key": "Owner",
          "value": "c289b99209ca4e189095cdecebbd092d"
        },
        {
          "key": "Protected",
          "value": "true"
        },
        {
          "key": "Visibility",
          "value": "public"
        },
        {
          "key": "Hidden",
          "value": "false"
        },
        {
          "key": "Checksum",
          "value": "4c563c67af5d73b15e7003af724d6b55"
        },
        {
          "key": "Properties",
          "value": "{deprecate_date:null,description:Ubuntu Server 22.04.5 LTS (2026.03.10),hw_cpu_sockets:1,hw_qemu_guest_agent:yes,hw_vif_multiqueue_enabled:true,login_username:ubuntu,max_cpu:,min_cpu:0,monitoring_agent:sysmon,nhncloud_allow_autoscale:true,nhncloud_allow_cgroup:v2,nhncloud_allow_compute_flavor:true,nhncloud_allow_cpu_flavor:true,nhncloud_allow_download:false,nhncloud_allow_gpu_flavor:false,nhncloud_allow_image_create:true,nhncloud_allow_imagebuilder:true,nhncloud_allow_instance_template:true,nhncloud_allow_local_bootdisk_flavor:true,nhncloud_allow_nks_cpu_flavor:false,nhncloud_allow_nks_gpu_flavor:false,nhncloud_allow_user_script:true,nhncloud_category:OS,nhncloud_product:compute,os_architecture:amd64,os_distro:ubuntu,os_type:linux,os_version:Server 22.04 LTS,project_domain:WDI;NORMAL,release_date:2026.03.10,tc_env:cloudmon,sysmon}"
        },
        {
          "key": "CreatedAt",
          "value": "2026-03-09T01:34:35Z"
        },
        {
          "key": "UpdatedAt",
          "value": "2026-03-09T21:14:13Z"
        },
        {
          "key": "File",
          "value": "/v2/images/0f07c795-2a46-44fc-a61b-fa0d96763ce2/file"
        },
        {
          "key": "Schema",
          "value": "/v2/schemas/image"
        },
        {
          "key": "VirtualSize",
          "value": "0"
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
      "connectionName": "nhn-kr1",
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
        },
        {
          "Ports": "8086",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/21"
        },
        {
          "Ports": "8086",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.1.0/24"
        }
      ],
      "cspResourceId": ""
    },
    {
      "name": "mig-sg-02",
      "connectionName": "nhn-kr1",
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
- **MCI ID**: `my09-infra101`
- **MCI Name**: `my09-infra101`
- **Status**: `Running:3 (R:3/3)`

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my09-infra101",
  "uid": "tbamt74trtru61e9kjd9",
  "name": "my09-infra101",
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
    "cm-migrated-at": "2026-10-12T10:55:45Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my09-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my09-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbamt74trtru61e9kjd9"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbumqdg2mjttqidbelrn",
      "cspResourceName": "tbumqdg2mjttqidbelrn",
      "cspResourceId": "c54494b3-165c-4fe0-a924-32b4a0fec1fc",
      "name": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "Pangyo (South Korea)",
        "latitude": 37.390889,
        "longitude": 127.096792
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-12 10:57:38",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "133.186.223.136",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.11",
      "privateDNS": "",
      "rootDiskType": "General_HDD",
      "rootDiskSize": 20,
      "rootDeviceName": "/dev/vda",
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
      "cspVNetId": "63f62c96-b2f4-4602-8385-62d049b22434",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "b31a5d3f-d214-4469-92ca-c8588598023d",
      "networkInterface": "ac3c341f-2011-4db1-9259-75cc5083ab5b",
      "securityGroupIds": [
        "my09-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tbhjbjpobn3enr3num8l",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-10-12T10:57:41Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:38Z",
          "completedTime": "2026-10-12T10:58:00Z",
          "elapsedTime": 22,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my09-infra101-tb6v49d6kgnom0kh568c",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:43Z",
          "completedTime": "2026-10-12T10:57:44Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbumqdg2mjttqidbelrn 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my09-ng-influxdb-back-1",
      "uid": "tbjimg37g8bs1tp9a3ct",
      "cspResourceName": "tbjimg37g8bs1tp9a3ct",
      "cspResourceId": "246fe00a-bfd2-402b-b140-22545a2f1c6c",
      "name": "my09-ng-influxdb-back-1",
      "nodeGroupId": "my09-ng-influxdb-back",
      "location": {
        "display": "Pangyo (South Korea)",
        "latitude": 37.390889,
        "longitude": 127.096792
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-12 10:57:21",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "133.186.209.144",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.87",
      "privateDNS": "",
      "rootDiskType": "General_HDD",
      "rootDiskSize": 30,
      "rootDeviceName": "/dev/vda",
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
      "cspVNetId": "63f62c96-b2f4-4602-8385-62d049b22434",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "b31a5d3f-d214-4469-92ca-c8588598023d",
      "networkInterface": "c7d34596-85ef-452e-9a4a-362053792dec",
      "securityGroupIds": [
        "my09-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tbhjbjpobn3enr3num8l",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-10-12T10:57:39Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:38Z",
          "completedTime": "2026-10-12T10:57:43Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my09-infra101-tb6v49d6kgnom0kh568c",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:43Z",
          "completedTime": "2026-10-12T10:57:45Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbjimg37g8bs1tp9a3ct 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my09-ng-influxdb-back-2",
      "uid": "tbtcvsqdbqus3diq9190",
      "cspResourceName": "tbtcvsqdbqus3diq9190",
      "cspResourceId": "b0993afa-0c7a-44c7-809b-236ca241113f",
      "name": "my09-ng-influxdb-back-2",
      "nodeGroupId": "my09-ng-influxdb-back",
      "location": {
        "display": "Pangyo (South Korea)",
        "latitude": 37.390889,
        "longitude": 127.096792
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-12 10:57:26",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "133.186.209.114",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.100",
      "privateDNS": "",
      "rootDiskType": "General_HDD",
      "rootDiskSize": 30,
      "rootDeviceName": "/dev/vda",
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
      "cspVNetId": "63f62c96-b2f4-4602-8385-62d049b22434",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "b31a5d3f-d214-4469-92ca-c8588598023d",
      "networkInterface": "9509e6a0-72e1-4a56-8104-f9a230f89e4f",
      "securityGroupIds": [
        "my09-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tbhjbjpobn3enr3num8l",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-10-12T10:57:41Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:38Z",
          "completedTime": "2026-10-12T10:57:42Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my09-infra101-tb6v49d6kgnom0kh568c",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:43Z",
          "completedTime": "2026-10-12T10:57:44Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbtcvsqdbqus3diq9190 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my09-mig-vnet-01",
      "name": "my09-mig-vnet-01",
      "infraId": "my09-infra101",
      "vNetId": "my09-mig-vnet-01",
      "connectionNames": [
        "nhn-kr1"
      ],
      "providerNames": [
        "nhn"
      ],
      "regionNames": [
        "KR1"
      ],
      "nodeGroupIds": [
        "my09-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my09-ng-influxdb-back"
      ],
      "nodeIds": [
        "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my09-ng-influxdb-back-1",
        "my09-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-influxdb-back-2",
            "nodeIp": "133.186.209.114",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbtcvsqdbqus3diq9190 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "133.186.223.136",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbumqdg2mjttqidbelrn 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-influxdb-back-1",
            "nodeIp": "133.186.209.144",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbjimg37g8bs1tp9a3ct 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my09-infra101-tb6v49d6kgnom0kh568c"
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
- **Count**: 8

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "infra": [
    {
      "resourceType": "infra",
      "id": "my01-infra101",
      "uid": "tb7jivaqmhqeqh8mpt2e",
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
        "cm-migrated-at": "2026-10-12T10:54:50Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my01-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my01-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tb7jivaqmhqeqh8mpt2e"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb285m5mg1j774dvvu4k",
          "cspResourceName": "tb285m5mg1j774dvvu4k",
          "cspResourceId": "i-01e11280c6d560912",
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
          "createdTime": "2026-10-12 10:55:13",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-10-12 10:55:13",
            "sys.cspResourceId": "i-01e11280c6d560912",
            "sys.cspResourceName": "tb285m5mg1j774dvvu4k",
            "sys.id": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tb285m5mg1j774dvvu4k",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "43.202.65.168",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.45",
          "privateDNS": "ip-10-0-1-45.ap-northeast-2.compute.internal",
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
          "cspVNetId": "vpc-08afb74d5d25b8401",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-07356212594969c9e",
          "networkInterface": "eni-0523e5626d1021d81",
          "securityGroupIds": [
            "my01-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tblqeie98atc8phigbn8"
        },
        {
          "resourceType": "node",
          "id": "my01-ng-influxdb-back-1",
          "uid": "tb0rphiopmk4jbj63gu9",
          "cspResourceName": "tb0rphiopmk4jbj63gu9",
          "cspResourceId": "i-0ccffb9271ef8b826",
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
          "createdTime": "2026-10-12 10:55:12",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-10-12 10:55:12",
            "sys.cspResourceId": "i-0ccffb9271ef8b826",
            "sys.cspResourceName": "tb0rphiopmk4jbj63gu9",
            "sys.id": "my01-ng-influxdb-back-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-influxdb-back",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tb0rphiopmk4jbj63gu9",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "3.35.0.184",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.252",
          "privateDNS": "ip-10-0-1-252.ap-northeast-2.compute.internal",
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
          "cspVNetId": "vpc-08afb74d5d25b8401",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-07356212594969c9e",
          "networkInterface": "eni-003a26766f6ec524f",
          "securityGroupIds": [
            "my01-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tblqeie98atc8phigbn8"
        },
        {
          "resourceType": "node",
          "id": "my01-ng-influxdb-back-2",
          "uid": "tbvicduatqk2kbkh4sh1",
          "cspResourceName": "tbvicduatqk2kbkh4sh1",
          "cspResourceId": "i-0fe239de8057ad524",
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
          "createdTime": "2026-10-12 10:55:13",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-10-12 10:55:13",
            "sys.cspResourceId": "i-0fe239de8057ad524",
            "sys.cspResourceName": "tbvicduatqk2kbkh4sh1",
            "sys.id": "my01-ng-influxdb-back-2",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-influxdb-back",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tbvicduatqk2kbkh4sh1",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "13.124.59.129",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.157",
          "privateDNS": "ip-10-0-1-157.ap-northeast-2.compute.internal",
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
          "cspVNetId": "vpc-08afb74d5d25b8401",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-07356212594969c9e",
          "networkInterface": "eni-0bdf4f149596775ca",
          "securityGroupIds": [
            "my01-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tblqeie98atc8phigbn8"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my02-infra101",
      "uid": "tbhgfnet31pr9ql9j7bk",
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
        "cm-migrated-at": "2026-10-12T10:55:25Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my02-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbhgfnet31pr9ql9j7bk"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbgo8nk6f2vdi4ud1di3",
          "cspResourceName": "tbgo8nk6f2vdi4ud1di3",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbgo8nk6f2vdi4ud1di3",
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
          "createdTime": "2026-10-12 10:55:38",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-10-12 10:55:38",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbgo8nk6f2vdi4ud1di3",
            "sys.cspResourceName": "tbgo8nk6f2vdi4ud1di3",
            "sys.id": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tbgo8nk6f2vdi4ud1di3",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "52.231.199.148",
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbr90eqdo0i08mloa9d4",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbr90eqdo0i08mloa9d4/subnets/tb65ekb0igsd7f9m5um8",
          "networkInterface": "tbgo8nk6f2vdi4ud1di3-76427-VNic",
          "securityGroupIds": [
            "my02-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tben35a2d0tikbeoom6u"
        },
        {
          "resourceType": "node",
          "id": "my02-ng-influxdb-back-1",
          "uid": "tbj14c6t6l7e1i7n3o19",
          "cspResourceName": "tbj14c6t6l7e1i7n3o19",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbj14c6t6l7e1i7n3o19",
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
          "createdTime": "2026-10-12 10:55:36",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-10-12 10:55:36",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbj14c6t6l7e1i7n3o19",
            "sys.cspResourceName": "tbj14c6t6l7e1i7n3o19",
            "sys.id": "my02-ng-influxdb-back-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-influxdb-back",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tbj14c6t6l7e1i7n3o19",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.14.38",
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbr90eqdo0i08mloa9d4",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbr90eqdo0i08mloa9d4/subnets/tb65ekb0igsd7f9m5um8",
          "networkInterface": "tbj14c6t6l7e1i7n3o19-84301-VNic",
          "securityGroupIds": [
            "my02-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tben35a2d0tikbeoom6u"
        },
        {
          "resourceType": "node",
          "id": "my02-ng-influxdb-back-2",
          "uid": "tbs7emlcpd5nlm03jso5",
          "cspResourceName": "tbs7emlcpd5nlm03jso5",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbs7emlcpd5nlm03jso5",
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
          "createdTime": "2026-10-12 10:55:37",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-10-12 10:55:37",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbs7emlcpd5nlm03jso5",
            "sys.cspResourceName": "tbs7emlcpd5nlm03jso5",
            "sys.id": "my02-ng-influxdb-back-2",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-influxdb-back",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tbs7emlcpd5nlm03jso5",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.157.253",
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbr90eqdo0i08mloa9d4",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbr90eqdo0i08mloa9d4/subnets/tb65ekb0igsd7f9m5um8",
          "networkInterface": "tbs7emlcpd5nlm03jso5-17738-VNic",
          "securityGroupIds": [
            "my02-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tben35a2d0tikbeoom6u"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my03-infra101",
      "uid": "tbaciag0qhpm0el9ei0q",
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
      "label": null,
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb34cj51eke012cm5vec",
          "cspResourceName": "tb34cj51eke012cm5vec",
          "cspResourceId": "tb34cj51eke012cm5vec",
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
          "createdTime": "2026-10-12 10:57:45",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-10-12 10:57:45",
            "sys.cspResourceId": "tb34cj51eke012cm5vec",
            "sys.cspResourceName": "tb34cj51eke012cm5vec",
            "sys.id": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tb34cj51eke012cm5vec",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.50.50.20",
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
          "cspVNetId": "tb6ei259h9mvdch4rumt",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tbvssd98bnj6fq03668o",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tbo0ov0jqt6lbp7d7e4m"
        },
        {
          "resourceType": "node",
          "id": "my03-ng-influxdb-back-1",
          "uid": "tb6iphm7mug3sutkbmrj",
          "cspResourceName": "tb6iphm7mug3sutkbmrj",
          "cspResourceId": "tb6iphm7mug3sutkbmrj",
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
          "createdTime": "2026-10-12 10:57:46",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-10-12 10:57:46",
            "sys.cspResourceId": "tb6iphm7mug3sutkbmrj",
            "sys.cspResourceName": "tb6iphm7mug3sutkbmrj",
            "sys.id": "my03-ng-influxdb-back-1",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-influxdb-back",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tb6iphm7mug3sutkbmrj",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "8.230.11.202",
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
          "cspVNetId": "tb6ei259h9mvdch4rumt",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tbvssd98bnj6fq03668o",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tbo0ov0jqt6lbp7d7e4m"
        },
        {
          "resourceType": "node",
          "id": "my03-ng-influxdb-back-2",
          "uid": "tb0vbtatce8oe2gk7l55",
          "cspResourceName": "tb0vbtatce8oe2gk7l55",
          "cspResourceId": "tb0vbtatce8oe2gk7l55",
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
          "createdTime": "2026-10-12 10:57:44",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-10-12 10:57:44",
            "sys.cspResourceId": "tb0vbtatce8oe2gk7l55",
            "sys.cspResourceName": "tb0vbtatce8oe2gk7l55",
            "sys.id": "my03-ng-influxdb-back-2",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-influxdb-back",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tb0vbtatce8oe2gk7l55",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.64.253.99",
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
          "cspVNetId": "tb6ei259h9mvdch4rumt",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tbvssd98bnj6fq03668o",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tbo0ov0jqt6lbp7d7e4m"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my04-infra101",
      "uid": "tbr6e4i4thobqvnp77cd",
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
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "cm-migrated-at": "2026-10-12T10:54:58Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my04-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my04-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbr6e4i4thobqvnp77cd"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbs1873dgjmra2nm490f",
          "cspResourceName": "tbs1873dgjmra2nm490f",
          "cspResourceId": "i-mj78vnkdsw3hgl5t0j1i",
          "name": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-10-12 10:55:10",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-10-12 10:55:10",
            "sys.cspResourceId": "i-mj78vnkdsw3hgl5t0j1i",
            "sys.cspResourceName": "tbs1873dgjmra2nm490f",
            "sys.id": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbs1873dgjmra2nm490f",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.249.240",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.73",
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
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7r2m3wqm09p88ugm6if",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj77koiujuio5e71st964",
          "networkInterface": "eni-mj78vnkdsw3hgl5yycxw",
          "securityGroupIds": [
            "my04-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tbghcomebhflfpstm793"
        },
        {
          "resourceType": "node",
          "id": "my04-ng-influxdb-back-1",
          "uid": "tbstgdpvnav7m9kapobr",
          "cspResourceName": "tbstgdpvnav7m9kapobr",
          "cspResourceId": "i-mj75o7wmbqfttpuy12n7",
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
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-12 10:55:06",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-10-12 10:55:06",
            "sys.cspResourceId": "i-mj75o7wmbqfttpuy12n7",
            "sys.cspResourceName": "tbstgdpvnav7m9kapobr",
            "sys.id": "my04-ng-influxdb-back-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-influxdb-back",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbstgdpvnav7m9kapobr",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.251.194",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.72",
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
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "cspSpecName": "ecs.e-c1m4.xlarge",
          "spec": {
            "cspSpecName": "ecs.e-c1m4.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1582
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7r2m3wqm09p88ugm6if",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj77koiujuio5e71st964",
          "networkInterface": "eni-mj75o7wmbqfttpv47h21",
          "securityGroupIds": [
            "my04-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tbghcomebhflfpstm793"
        },
        {
          "resourceType": "node",
          "id": "my04-ng-influxdb-back-2",
          "uid": "tbpsj9r6espo7v0rf7ff",
          "cspResourceName": "tbpsj9r6espo7v0rf7ff",
          "cspResourceId": "i-mj75o7wmbqfttpuy12n8",
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
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-12 10:55:11",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-10-12 10:55:11",
            "sys.cspResourceId": "i-mj75o7wmbqfttpuy12n8",
            "sys.cspResourceName": "tbpsj9r6espo7v0rf7ff",
            "sys.id": "my04-ng-influxdb-back-2",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-influxdb-back",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbpsj9r6espo7v0rf7ff",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.250.25",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.74",
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
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "cspSpecName": "ecs.e-c1m4.xlarge",
          "spec": {
            "cspSpecName": "ecs.e-c1m4.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1582
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7r2m3wqm09p88ugm6if",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj77koiujuio5e71st964",
          "networkInterface": "eni-mj75o7wmbqfttpv47h24",
          "securityGroupIds": [
            "my04-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tbghcomebhflfpstm793"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my05-infra101",
      "uid": "tbmsm54ptrs85bg2d7sv",
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
        "cm-migrated-at": "2026-10-12T10:55:20Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my05-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbmsm54ptrs85bg2d7sv"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb7cfn4jaa4fg0i9flen",
          "cspResourceName": "tb7cfn4jaa4fg0i9flen",
          "cspResourceId": "ins-hzhqfrzx",
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
          "createdTime": "2026-10-12 10:55:34",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-10-12 10:55:34",
            "sys.cspResourceId": "ins-hzhqfrzx",
            "sys.cspResourceName": "tb7cfn4jaa4fg0i9flen",
            "sys.id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tb7cfn4jaa4fg0i9flen",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.133.248.198",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.9",
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
          "specId": "tencent+ap-seoul+bf1.medium2",
          "cspSpecName": "BF1.MEDIUM2",
          "spec": {
            "cspSpecName": "BF1.MEDIUM2",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.02
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
          "cspVNetId": "vpc-2hz6b4vc",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-mn81rac7",
          "networkInterface": "eni-1579pusm",
          "securityGroupIds": [
            "my05-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-0u5rnpeb"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-1",
          "uid": "tb2ah0srjvbcm440j9gg",
          "cspResourceName": "tb2ah0srjvbcm440j9gg",
          "cspResourceId": "ins-c6vp1w4f",
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
          "createdTime": "2026-10-12 10:55:49",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-10-12 10:55:49",
            "sys.cspResourceId": "ins-c6vp1w4f",
            "sys.cspResourceName": "tb2ah0srjvbcm440j9gg",
            "sys.id": "my05-ng-influxdb-back-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tb2ah0srjvbcm440j9gg",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.155.200.27",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.8",
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
          "specId": "tencent+ap-seoul+bf1.large16",
          "cspSpecName": "BF1.LARGE16",
          "spec": {
            "cspSpecName": "BF1.LARGE16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.15
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
          "cspVNetId": "vpc-2hz6b4vc",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-mn81rac7",
          "networkInterface": "eni-bbgu30vi",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-0u5rnpeb"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-2",
          "uid": "tb5kvr71nkkbsfmd33v8",
          "cspResourceName": "tb5kvr71nkkbsfmd33v8",
          "cspResourceId": "ins-90viujqt",
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
          "createdTime": "2026-10-12 10:55:35",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-10-12 10:55:35",
            "sys.cspResourceId": "ins-90viujqt",
            "sys.cspResourceName": "tb5kvr71nkkbsfmd33v8",
            "sys.id": "my05-ng-influxdb-back-2",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tb5kvr71nkkbsfmd33v8",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.155.157.6",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
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
          "specId": "tencent+ap-seoul+bf1.large16",
          "cspSpecName": "BF1.LARGE16",
          "spec": {
            "cspSpecName": "BF1.LARGE16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.15
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
          "cspVNetId": "vpc-2hz6b4vc",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-mn81rac7",
          "networkInterface": "eni-27rmbrhk",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-0u5rnpeb"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my06-infra101",
      "uid": "tbdbo98h4av8i7ot3ms7",
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
        "cm-migrated-at": "2026-10-12T10:55:51Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my06-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbdbo98h4av8i7ot3ms7"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbicmmgcumo0subq1boh",
          "cspResourceName": "tbicmmgcumo0subq1boh",
          "cspResourceId": "02h7_b73ff4f6-7d9a-4c6d-93ab-a0d4c6ef31c5",
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
          "createdTime": "2026-10-12 10:56:11",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-10-12 10:56:11",
            "sys.cspResourceId": "02h7_b73ff4f6-7d9a-4c6d-93ab-a0d4c6ef31c5",
            "sys.cspResourceName": "tbicmmgcumo0subq1boh",
            "sys.id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbicmmgcumo0subq1boh",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.93.202",
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
          "cspVNetId": "r026-d88398f2-3008-4600-a9b7-09de248c9d82",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-0cdc9592-fc29-40bb-a412-28b334f8caae",
          "networkInterface": "neatness-plunder-sacked-babbling",
          "securityGroupIds": [
            "my06-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-2afea093-cd50-4d8f-bb32-0bd32af60c50"
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-1",
          "uid": "tbnsjqfoqfr7q502ba95",
          "cspResourceName": "tbnsjqfoqfr7q502ba95",
          "cspResourceId": "02h7_70df7833-98d3-456c-8aa7-43a3f51f73d6",
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
          "createdTime": "2026-10-12 10:56:11",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-10-12 10:56:11",
            "sys.cspResourceId": "02h7_70df7833-98d3-456c-8aa7-43a3f51f73d6",
            "sys.cspResourceName": "tbnsjqfoqfr7q502ba95",
            "sys.id": "my06-ng-influxdb-back-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-influxdb-back",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbnsjqfoqfr7q502ba95",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.99.21",
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
          "cspVNetId": "r026-d88398f2-3008-4600-a9b7-09de248c9d82",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-0cdc9592-fc29-40bb-a412-28b334f8caae",
          "networkInterface": "sanitizer-ethically-majestic-unrevised",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-2afea093-cd50-4d8f-bb32-0bd32af60c50"
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-2",
          "uid": "tbosamc62njbprdica60",
          "cspResourceName": "tbosamc62njbprdica60",
          "cspResourceId": "02h7_d3acbb7a-5ae4-4bb3-9d9e-62c1a321294f",
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
          "createdTime": "2026-10-12 10:56:11",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-10-12 10:56:11",
            "sys.cspResourceId": "02h7_d3acbb7a-5ae4-4bb3-9d9e-62c1a321294f",
            "sys.cspResourceName": "tbosamc62njbprdica60",
            "sys.id": "my06-ng-influxdb-back-2",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-influxdb-back",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbosamc62njbprdica60",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.103.103",
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
          "cspVNetId": "r026-d88398f2-3008-4600-a9b7-09de248c9d82",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-0cdc9592-fc29-40bb-a412-28b334f8caae",
          "networkInterface": "sterilize-prankster-ensure-likewise",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-2afea093-cd50-4d8f-bb32-0bd32af60c50"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my08-infra101",
      "uid": "tbu9j7gf6ltd4dvfnn2q",
      "name": "my08-infra101",
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
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": null,
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb0i614dcarcujrioaab",
          "name": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my08-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
          },
          "description": "",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
          "cspVNetId": "",
          "subnetId": "my08-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my08-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my08-mig-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my08-ng-influxdb-back-1",
          "uid": "tbu2m9qgqvljp90p6ik3",
          "name": "my08-ng-influxdb-back-1",
          "nodeGroupId": "my08-ng-influxdb-back",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back"
          },
          "description": "",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
          "cspVNetId": "",
          "subnetId": "my08-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my08-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my08-mig-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my08-ng-influxdb-back-2",
          "uid": "tbt6qgooqti510ghbklm",
          "name": "my08-ng-influxdb-back-2",
          "nodeGroupId": "my08-ng-influxdb-back",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back"
          },
          "description": "",
          "region": {
            "region": "KR",
            "zone": "KR-1"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
          "cspVNetId": "",
          "subnetId": "my08-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my08-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my08-mig-sshkey-01",
          "cspSshKeyId": ""
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my09-infra101",
      "uid": "tbamt74trtru61e9kjd9",
      "name": "my09-infra101",
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
        "cm-migrated-at": "2026-10-12T10:55:45Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my09-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my09-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbamt74trtru61e9kjd9"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbumqdg2mjttqidbelrn",
          "cspResourceName": "tbumqdg2mjttqidbelrn",
          "cspResourceId": "c54494b3-165c-4fe0-a924-32b4a0fec1fc",
          "name": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-12 10:57:38",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "nhn-kr1",
            "sys.createdTime": "2026-10-12 10:57:38",
            "sys.cspResourceId": "c54494b3-165c-4fe0-a924-32b4a0fec1fc",
            "sys.cspResourceName": "tbumqdg2mjttqidbelrn",
            "sys.id": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my09-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my09-mig-subnet-01",
            "sys.uid": "tbumqdg2mjttqidbelrn",
            "sys.vNetId": "my09-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "133.186.223.136",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.11",
          "privateDNS": "",
          "rootDiskType": "General_HDD",
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
          "cspVNetId": "63f62c96-b2f4-4602-8385-62d049b22434",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "b31a5d3f-d214-4469-92ca-c8588598023d",
          "networkInterface": "ac3c341f-2011-4db1-9259-75cc5083ab5b",
          "securityGroupIds": [
            "my09-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": "tbhjbjpobn3enr3num8l"
        },
        {
          "resourceType": "node",
          "id": "my09-ng-influxdb-back-1",
          "uid": "tbjimg37g8bs1tp9a3ct",
          "cspResourceName": "tbjimg37g8bs1tp9a3ct",
          "cspResourceId": "246fe00a-bfd2-402b-b140-22545a2f1c6c",
          "name": "my09-ng-influxdb-back-1",
          "nodeGroupId": "my09-ng-influxdb-back",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-12 10:57:21",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "nhn-kr1",
            "sys.createdTime": "2026-10-12 10:57:21",
            "sys.cspResourceId": "246fe00a-bfd2-402b-b140-22545a2f1c6c",
            "sys.cspResourceName": "tbjimg37g8bs1tp9a3ct",
            "sys.id": "my09-ng-influxdb-back-1",
            "sys.infraId": "my09-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my09-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my09-ng-influxdb-back",
            "sys.subnetId": "my09-mig-subnet-01",
            "sys.uid": "tbjimg37g8bs1tp9a3ct",
            "sys.vNetId": "my09-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "133.186.209.144",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.87",
          "privateDNS": "",
          "rootDiskType": "General_HDD",
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
          "cspVNetId": "63f62c96-b2f4-4602-8385-62d049b22434",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "b31a5d3f-d214-4469-92ca-c8588598023d",
          "networkInterface": "c7d34596-85ef-452e-9a4a-362053792dec",
          "securityGroupIds": [
            "my09-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": "tbhjbjpobn3enr3num8l"
        },
        {
          "resourceType": "node",
          "id": "my09-ng-influxdb-back-2",
          "uid": "tbtcvsqdbqus3diq9190",
          "cspResourceName": "tbtcvsqdbqus3diq9190",
          "cspResourceId": "b0993afa-0c7a-44c7-809b-236ca241113f",
          "name": "my09-ng-influxdb-back-2",
          "nodeGroupId": "my09-ng-influxdb-back",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-10-12 10:57:26",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "nhn-kr1",
            "sys.createdTime": "2026-10-12 10:57:26",
            "sys.cspResourceId": "b0993afa-0c7a-44c7-809b-236ca241113f",
            "sys.cspResourceName": "tbtcvsqdbqus3diq9190",
            "sys.id": "my09-ng-influxdb-back-2",
            "sys.infraId": "my09-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my09-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my09-ng-influxdb-back",
            "sys.subnetId": "my09-mig-subnet-01",
            "sys.uid": "tbtcvsqdbqus3diq9190",
            "sys.vNetId": "my09-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "133.186.209.114",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.100",
          "privateDNS": "",
          "rootDiskType": "General_HDD",
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
          "cspVNetId": "63f62c96-b2f4-4602-8385-62d049b22434",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "b31a5d3f-d214-4469-92ca-c8588598023d",
          "networkInterface": "9509e6a0-72e1-4a56-8104-f9a230f89e4f",
          "securityGroupIds": [
            "my09-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": "tbhjbjpobn3enr3num8l"
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
- **IDs**: [my01-infra101 my02-infra101 my03-infra101 my04-infra101 my05-infra101 my06-infra101 my08-infra101 my09-infra101]

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
  "id": "my09-infra101",
  "uid": "tbamt74trtru61e9kjd9",
  "name": "my09-infra101",
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
    "cm-migrated-at": "2026-10-12T10:55:45Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my09-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my09-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbamt74trtru61e9kjd9"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbumqdg2mjttqidbelrn",
      "cspResourceName": "tbumqdg2mjttqidbelrn",
      "cspResourceId": "c54494b3-165c-4fe0-a924-32b4a0fec1fc",
      "name": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "Pangyo (South Korea)",
        "latitude": 37.390889,
        "longitude": 127.096792
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-12 10:57:38",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "133.186.223.136",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.11",
      "privateDNS": "",
      "rootDiskType": "General_HDD",
      "rootDiskSize": 20,
      "rootDeviceName": "/dev/vda",
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
      "cspVNetId": "63f62c96-b2f4-4602-8385-62d049b22434",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "b31a5d3f-d214-4469-92ca-c8588598023d",
      "networkInterface": "ac3c341f-2011-4db1-9259-75cc5083ab5b",
      "securityGroupIds": [
        "my09-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tbhjbjpobn3enr3num8l",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-10-12T10:57:41Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:38Z",
          "completedTime": "2026-10-12T10:58:00Z",
          "elapsedTime": 22,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my09-infra101-tb6v49d6kgnom0kh568c",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:43Z",
          "completedTime": "2026-10-12T10:57:44Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbumqdg2mjttqidbelrn 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my09-ng-influxdb-back-1",
      "uid": "tbjimg37g8bs1tp9a3ct",
      "cspResourceName": "tbjimg37g8bs1tp9a3ct",
      "cspResourceId": "246fe00a-bfd2-402b-b140-22545a2f1c6c",
      "name": "my09-ng-influxdb-back-1",
      "nodeGroupId": "my09-ng-influxdb-back",
      "location": {
        "display": "Pangyo (South Korea)",
        "latitude": 37.390889,
        "longitude": 127.096792
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-12 10:57:21",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "133.186.209.144",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.87",
      "privateDNS": "",
      "rootDiskType": "General_HDD",
      "rootDiskSize": 30,
      "rootDeviceName": "/dev/vda",
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
      "cspVNetId": "63f62c96-b2f4-4602-8385-62d049b22434",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "b31a5d3f-d214-4469-92ca-c8588598023d",
      "networkInterface": "c7d34596-85ef-452e-9a4a-362053792dec",
      "securityGroupIds": [
        "my09-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tbhjbjpobn3enr3num8l",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-10-12T10:57:39Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:38Z",
          "completedTime": "2026-10-12T10:57:43Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my09-infra101-tb6v49d6kgnom0kh568c",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:43Z",
          "completedTime": "2026-10-12T10:57:45Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbjimg37g8bs1tp9a3ct 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my09-ng-influxdb-back-2",
      "uid": "tbtcvsqdbqus3diq9190",
      "cspResourceName": "tbtcvsqdbqus3diq9190",
      "cspResourceId": "b0993afa-0c7a-44c7-809b-236ca241113f",
      "name": "my09-ng-influxdb-back-2",
      "nodeGroupId": "my09-ng-influxdb-back",
      "location": {
        "display": "Pangyo (South Korea)",
        "latitude": 37.390889,
        "longitude": 127.096792
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-12 10:57:26",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "133.186.209.114",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.100",
      "privateDNS": "",
      "rootDiskType": "General_HDD",
      "rootDiskSize": 30,
      "rootDeviceName": "/dev/vda",
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
      "cspVNetId": "63f62c96-b2f4-4602-8385-62d049b22434",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "b31a5d3f-d214-4469-92ca-c8588598023d",
      "networkInterface": "9509e6a0-72e1-4a56-8104-f9a230f89e4f",
      "securityGroupIds": [
        "my09-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tbhjbjpobn3enr3num8l",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-10-12T10:57:41Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:38Z",
          "completedTime": "2026-10-12T10:57:42Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my09-infra101-tb6v49d6kgnom0kh568c",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T10:57:43Z",
          "completedTime": "2026-10-12T10:57:44Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbtcvsqdbqus3diq9190 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my09-mig-vnet-01",
      "name": "my09-mig-vnet-01",
      "infraId": "my09-infra101",
      "vNetId": "my09-mig-vnet-01",
      "connectionNames": [
        "nhn-kr1"
      ],
      "providerNames": [
        "nhn"
      ],
      "regionNames": [
        "KR1"
      ],
      "nodeGroupIds": [
        "my09-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my09-ng-influxdb-back"
      ],
      "nodeIds": [
        "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my09-ng-influxdb-back-1",
        "my09-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-influxdb-back-2",
            "nodeIp": "133.186.209.114",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbtcvsqdbqus3diq9190 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "133.186.223.136",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbumqdg2mjttqidbelrn 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-influxdb-back-1",
            "nodeIp": "133.186.209.144",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbjimg37g8bs1tp9a3ct 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my09-infra101-tb6v49d6kgnom0kh568c"
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
      "id": "my09-ng-influxdb-back",
      "cspResourceName": "tbqgm6k1n6dig9lbdcjj",
      "cspResourceId": "86618fc4-dac7-4585-9364-e31bbd8c5b99",
      "name": "my09-ng-influxdb-back",
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
      "type": "PUBLIC",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "ip": "133.186.244.227",
        "port": "9999"
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my09-ng-influxdb-back",
        "nodes": [
          "my09-ng-influxdb-back-1",
          "my09-ng-influxdb-back-2"
        ]
      },
      "healthChecker": {
        "protocol": "TCP",
        "port": "8086",
        "interval": 10,
        "threshold": 3,
        "timeout": 10
      },
      "createdTime": "2026-10-12T10:59:18.384935578Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "ProvisioningStatus",
          "value": "ACTIVE"
        },
        {
          "key": "TenantID",
          "value": "6fc6c2ef568f45f7a9bdc6be211f52f9"
        },
        {
          "key": "Provider",
          "value": "haproxy"
        },
        {
          "key": "Name",
          "value": "tbqgm6k1n6dig9lbdcjj"
        },
        {
          "key": "Listeners",
          "value": "{default_pool_id:,protocol:,description:,name:,loadbalancers:null,tenant_id:,project_id:,admin_state_up:false,connection_limit:0,keepalive_timeout:0,default_tls_container_ref:,sni_container_refs:null,protocol_port:0,id:22952323-d5c8-414f-bf27-bddbc1b3883f,allowed_cidrs:null,tls_versions:null,timeout_client_data:0,timeout_member_data:0,timeout_member_connect:0,timeout_tcp_inspect:0,insert_headers:null,provisioning_status:,l7policies:null,pools:null}"
        },
        {
          "key": "Pools",
          "value": "{lb_algorithm:,protocol:,description:,admin_state_up:false,tenant_id:,member_port:0,session_persistence:null,healthmonitor_id:,loadbalancers:null,listeners:null,members:null,id:26941b32-7cb6-4706-8210-9313772c2bd5,name:,project_id:,provider:,provisioning_status:,healthmonitor:{admin_state_up:false,health_check_port:0,delay:0,expected_codes:,max_retries:0,http_method:,timeout:0,pools:null,url_path:,type:,id:,name:,project_id:,max_retries_down:0,host_header:,provisioning_status:}}"
        },
        {
          "key": "VipAddress",
          "value": "10.0.1.40"
        },
        {
          "key": "VipPortID",
          "value": "43f40508-8c5e-4198-8f9a-a72ee0b27f27"
        },
        {
          "key": "VipSubnetID",
          "value": "b31a5d3f-d214-4469-92ca-c8588598023d"
        },
        {
          "key": "ID",
          "value": "86618fc4-dac7-4585-9364-e31bbd8c5b99"
        },
        {
          "key": "OperatingStatus",
          "value": "ONLINE"
        },
        {
          "key": "AdminStateUp",
          "value": "true"
        },
        {
          "key": "LoadBalancerType",
          "value": "shared"
        }
      ],
      "isAutoGenerated": false,
      "location": {
        "display": "Pangyo (South Korea)",
        "latitude": 37.390889,
        "longitude": 127.096792
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
    "id": "my09-ng-influxdb-back",
    "cspResourceName": "tbqgm6k1n6dig9lbdcjj",
    "cspResourceId": "86618fc4-dac7-4585-9364-e31bbd8c5b99",
    "name": "my09-ng-influxdb-back",
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
    "type": "PUBLIC",
    "scope": "REGION",
    "listener": {
      "protocol": "TCP",
      "ip": "133.186.244.227",
      "port": "9999"
    },
    "targetGroup": {
      "protocol": "TCP",
      "port": "8086",
      "nodeGroupId": "my09-ng-influxdb-back",
      "nodes": [
        "my09-ng-influxdb-back-1",
        "my09-ng-influxdb-back-2"
      ]
    },
    "healthChecker": {
      "protocol": "TCP",
      "port": "8086",
      "interval": 10,
      "threshold": 3,
      "timeout": 10
    },
    "createdTime": "2026-10-12T10:59:18.384935578Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "ProvisioningStatus",
        "value": "ACTIVE"
      },
      {
        "key": "TenantID",
        "value": "6fc6c2ef568f45f7a9bdc6be211f52f9"
      },
      {
        "key": "Provider",
        "value": "haproxy"
      },
      {
        "key": "Name",
        "value": "tbqgm6k1n6dig9lbdcjj"
      },
      {
        "key": "Listeners",
        "value": "{default_pool_id:,protocol:,description:,name:,loadbalancers:null,tenant_id:,project_id:,admin_state_up:false,connection_limit:0,keepalive_timeout:0,default_tls_container_ref:,sni_container_refs:null,protocol_port:0,id:22952323-d5c8-414f-bf27-bddbc1b3883f,allowed_cidrs:null,tls_versions:null,timeout_client_data:0,timeout_member_data:0,timeout_member_connect:0,timeout_tcp_inspect:0,insert_headers:null,provisioning_status:,l7policies:null,pools:null}"
      },
      {
        "key": "Pools",
        "value": "{lb_algorithm:,protocol:,description:,admin_state_up:false,tenant_id:,member_port:0,session_persistence:null,healthmonitor_id:,loadbalancers:null,listeners:null,members:null,id:26941b32-7cb6-4706-8210-9313772c2bd5,name:,project_id:,provider:,provisioning_status:,healthmonitor:{admin_state_up:false,health_check_port:0,delay:0,expected_codes:,max_retries:0,http_method:,timeout:0,pools:null,url_path:,type:,id:,name:,project_id:,max_retries_down:0,host_header:,provisioning_status:}}"
      },
      {
        "key": "VipAddress",
        "value": "10.0.1.40"
      },
      {
        "key": "VipPortID",
        "value": "43f40508-8c5e-4198-8f9a-a72ee0b27f27"
      },
      {
        "key": "VipSubnetID",
        "value": "b31a5d3f-d214-4469-92ca-c8588598023d"
      },
      {
        "key": "ID",
        "value": "86618fc4-dac7-4585-9364-e31bbd8c5b99"
      },
      {
        "key": "OperatingStatus",
        "value": "ONLINE"
      },
      {
        "key": "AdminStateUp",
        "value": "true"
      },
      {
        "key": "LoadBalancerType",
        "value": "shared"
      }
    ],
    "isAutoGenerated": false,
    "location": {
      "display": "Pangyo (South Korea)",
      "latitude": 37.390889,
      "longitude": 127.096792
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
  "id": "my09-ng-influxdb-back",
  "cspResourceName": "tbqgm6k1n6dig9lbdcjj",
  "cspResourceId": "86618fc4-dac7-4585-9364-e31bbd8c5b99",
  "name": "my09-ng-influxdb-back",
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
  "type": "PUBLIC",
  "scope": "REGION",
  "listener": {
    "protocol": "TCP",
    "ip": "133.186.244.227",
    "port": "9999"
  },
  "targetGroup": {
    "protocol": "TCP",
    "port": "8086",
    "nodeGroupId": "my09-ng-influxdb-back",
    "nodes": [
      "my09-ng-influxdb-back-1",
      "my09-ng-influxdb-back-2"
    ]
  },
  "healthChecker": {
    "protocol": "TCP",
    "port": "8086",
    "interval": 10,
    "threshold": 3,
    "timeout": 10
  },
  "createdTime": "2026-10-12T10:59:18.384935578Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "ProvisioningStatus",
      "value": "ACTIVE"
    },
    {
      "key": "TenantID",
      "value": "6fc6c2ef568f45f7a9bdc6be211f52f9"
    },
    {
      "key": "Provider",
      "value": "haproxy"
    },
    {
      "key": "Name",
      "value": "tbqgm6k1n6dig9lbdcjj"
    },
    {
      "key": "Listeners",
      "value": "{default_pool_id:,protocol:,description:,name:,loadbalancers:null,tenant_id:,project_id:,admin_state_up:false,connection_limit:0,keepalive_timeout:0,default_tls_container_ref:,sni_container_refs:null,protocol_port:0,id:22952323-d5c8-414f-bf27-bddbc1b3883f,allowed_cidrs:null,tls_versions:null,timeout_client_data:0,timeout_member_data:0,timeout_member_connect:0,timeout_tcp_inspect:0,insert_headers:null,provisioning_status:,l7policies:null,pools:null}"
    },
    {
      "key": "Pools",
      "value": "{lb_algorithm:,protocol:,description:,admin_state_up:false,tenant_id:,member_port:0,session_persistence:null,healthmonitor_id:,loadbalancers:null,listeners:null,members:null,id:26941b32-7cb6-4706-8210-9313772c2bd5,name:,project_id:,provider:,provisioning_status:,healthmonitor:{admin_state_up:false,health_check_port:0,delay:0,expected_codes:,max_retries:0,http_method:,timeout:0,pools:null,url_path:,type:,id:,name:,project_id:,max_retries_down:0,host_header:,provisioning_status:}}"
    },
    {
      "key": "VipAddress",
      "value": "10.0.1.40"
    },
    {
      "key": "VipPortID",
      "value": "43f40508-8c5e-4198-8f9a-a72ee0b27f27"
    },
    {
      "key": "VipSubnetID",
      "value": "b31a5d3f-d214-4469-92ca-c8588598023d"
    },
    {
      "key": "ID",
      "value": "86618fc4-dac7-4585-9364-e31bbd8c5b99"
    },
    {
      "key": "OperatingStatus",
      "value": "ONLINE"
    },
    {
      "key": "AdminStateUp",
      "value": "true"
    },
    {
      "key": "LoadBalancerType",
      "value": "shared"
    }
  ],
  "isAutoGenerated": false,
  "location": {
    "display": "Pangyo (South Korea)",
    "latitude": 37.390889,
    "longitude": 127.096792
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

**Generated At:** 2026-10-12 11:01:11

**Namespace:** mig01

**Infra Name:** my09-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my09-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | NHN |
| **Target Region** | KR1 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| c2.c2m2 | 2 | 2.0 | - | x86_64 | default | $0.0422 | 1 |
| r2.c4m16 | 4 | 16.0 | - | x86_64 | default | $0.1954 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| 0f07c795-2a46-44fc-a61b-fa0d96763ce2 | Ubuntu Server 22.04.5 LTS (2026.03.10) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | c54494b3-165c-4fe0-a924-32b4a0fec1fc | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 133.186.223.136<br>**Private IP:** 10.0.1.11<br>**SGs:** my09-mig-sg-02<br>**SSH:** my09-mig-sshkey-01 |
| my09-ng-influxdb-back-1 | 246fe00a-bfd2-402b-b140-22545a2f1c6c | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 133.186.209.144<br>**Private IP:** 10.0.1.87<br>**SGs:** my09-mig-sg-01<br>**SSH:** my09-mig-sshkey-01 |
| my09-ng-influxdb-back-2 | b0993afa-0c7a-44c7-809b-236ca241113f | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 133.186.209.114<br>**Private IP:** 10.0.1.100<br>**SGs:** my09-mig-sg-01<br>**SSH:** my09-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my09-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my09-mig-vnet-01 |
| **CSP VNet ID** | 63f62c96-b2f4-4602-8385-62d049b22434 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | nhn-kr1 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my09-mig-subnet-01 | b31a5d3f-d214-4469-92ca-c8588598023d | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my09-mig-sshkey-01 | tbhjbjpobn3enr3num8l | cb-user | 31:03:28:01:29:5a:fd:4d:b6:4f:08:58:97:1b:0c:ad |

### Security Groups

#### Security Group: my09-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my09-mig-sg-01 |
| **CSP Security Group ID** | ed3c8491-1514-4bce-8e7d-e982e62f49cf |
| **VNet** | my09-mig-vnet-01 |
| **Rule Count** | 8 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 9999 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.1.0/24 |
| inbound | TCP | 8086 | 10.0.0.0/21 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |

#### Security Group: my09-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my09-mig-sg-02 |
| **CSP Security Group ID** | 8b50179f-bdbc-4e39-aff3-1848ec9913c9 |
| **VNet** | my09-mig-vnet-01 |
| **Rule Count** | 4 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.4329 |
| **Per Day** | $10.39 |
| **Per Month (30 days)** | $311.69 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| NHN | KR1 | 3 | $0.4329 | $311.69 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | c2.c2m2 | $0.0422 | $30.37 |
| my09-ng-influxdb-back-1 | r2.c4m16 | $0.1954 | $140.66 |
| my09-ng-influxdb-back-2 | r2.c4m16 | $0.1954 | $140.66 |




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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my09-infra101)",
  "success": true
}
```

