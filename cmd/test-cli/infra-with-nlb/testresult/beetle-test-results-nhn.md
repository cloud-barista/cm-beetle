# CM-Beetle test results for NHN (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with NHN cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.5+ (83361e8)
- imdl: v0.1.22+ (83361e8)
- CB-Tumblebug: v0.14.0
- CB-Spider: v0.14.0
- CB-MapUI: v0.14.0
- Target CSP: NHN
- Target Region: kr1
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: November 25, 2026
- Test Time: 18:17:30 KST
- Test Execution: 2026-11-25 18:17:30 KST

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
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 2.436s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 536ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 8m6.334s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 72ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 5ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 55ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 867ms | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 1m15.051s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 12ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 12ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 1m27.926s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.247s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.446s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 29.361s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 2m4.127s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 14m27.847905811s

*Test executed on November 25, 2026 at 18:17:30 KST (2026-11-25 18:17:30 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=66.7% Max=100.0% Avg=88.9% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "nhn",
        "region": "kr1"
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
            "name": "ng-ip-10-0-1-30-62161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
          "uid": "tb28tmdpadmhbh75cnuo",
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
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-11-25T09:17:35.97589952Z",
          "availabilityRefreshAfter": "2026-11-25T10:17:35.97589952Z"
        },
        {
          "id": "nhn+kr1+c2.c2m2",
          "uid": "tbdc48vd4vups57fcrem",
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
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-11-25T09:17:37.696674878Z",
          "availabilityRefreshAfter": "2026-11-25T10:17:37.696674878Z"
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
          "uid": "tbf1lu751r8jpcodi27t",
          "name": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "nhn-kr1",
          "infraType": "",
          "fetchedTime": "2026.10.29 07:15:53 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=72.2% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "nhn",
        "region": "kr1"
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
            "name": "ng-ip-10-0-1-30-62161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=66.7%",
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
          "uid": "tb0bgprrr18sil1dvls3",
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
          "orderInFilteredResult": 2,
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
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-11-25T09:17:35.974960728Z",
          "availabilityRefreshAfter": "2026-11-25T10:17:35.974960728Z"
        },
        {
          "id": "nhn+kr1+m2.c2m4",
          "uid": "tbjd179s4cm2g5ij21tt",
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
          ],
          "availabilityStatus": "Available",
          "availabilityCheckedAt": "2026-11-25T09:17:37.696737458Z",
          "availabilityRefreshAfter": "2026-11-25T10:17:37.696737458Z"
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
          "uid": "tbf1lu751r8jpcodi27t",
          "name": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "nhn-kr1",
          "infraType": "",
          "fetchedTime": "2026.10.29 07:15:53 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
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
  "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=66.7% Max=100.0% Avg=88.9% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
  "targetCloud": {
    "csp": "nhn",
    "region": "kr1"
  },
  "targetInfra": {
    "name": "my09-infra101",
    "label": null,
    "systemLabel": "",
    "description": "NLB-aware recommended infrastructure for cloud migration",
    "nodeGroups": [
      {
        "name": "my09-ng-influxdb-back",
        "nodeGroupSize": 2,
        "label": {
          "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "nlbBackend": "influxdb_back"
        },
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
        "connectionName": "nhn-kr1",
        "specId": "nhn+kr1+r2.c4m16",
        "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
        "vNetId": "my09-mig-vnet-01",
        "subnetId": "my09-mig-subnet-01",
        "securityGroupIds": [
          "my09-mig-sg-01"
        ],
        "sshKeyId": "my09-mig-sshkey-01",
        "rootDiskSize": 30,
        "dataDiskIds": null
      },
      {
        "name": "my09-ng-ip-10-0-1-30-62161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
        "connectionName": "nhn-kr1",
        "specId": "nhn+kr1+c2.c2m2",
        "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
        "vNetId": "my09-mig-vnet-01",
        "subnetId": "my09-mig-subnet-01",
        "securityGroupIds": [
          "my09-mig-sg-02"
        ],
        "sshKeyId": "my09-mig-sshkey-01",
        "rootDiskSize": 20,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "my09-mig-vnet-01",
    "connectionName": "nhn-kr1",
    "cidrBlock": "10.0.0.0/21",
    "subnetInfoList": [
      {
        "name": "my09-mig-subnet-01",
        "ipv4_CIDR": "10.0.1.0/24",
        "description": "a recommended subnet for migration"
      }
    ],
    "description": "a recommended vNet for migration"
  },
  "targetSshKey": {
    "name": "my09-mig-sshkey-01",
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
      "uid": "tb28tmdpadmhbh75cnuo",
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
      ],
      "availabilityStatus": "Available",
      "availabilityCheckedAt": "2026-11-25T09:17:35.97589952Z",
      "availabilityRefreshAfter": "2026-11-25T10:17:35.97589952Z"
    },
    {
      "id": "nhn+kr1+c2.c2m2",
      "uid": "tbdc48vd4vups57fcrem",
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
      ],
      "availabilityStatus": "Available",
      "availabilityCheckedAt": "2026-11-25T09:17:37.696674878Z",
      "availabilityRefreshAfter": "2026-11-25T10:17:37.696674878Z"
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
      "uid": "tbf1lu751r8jpcodi27t",
      "name": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "nhn-kr1",
      "infraType": "",
      "fetchedTime": "2026.10.29 07:15:53 Thu",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
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
      "name": "my09-mig-sg-01",
      "connectionName": "nhn-kr1",
      "vNetId": "my09-mig-vnet-01",
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
      "name": "my09-mig-sg-02",
      "connectionName": "nhn-kr1",
      "vNetId": "my09-mig-vnet-01",
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
        "nodeGroupId": "my09-ng-influxdb-back"
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
  "uid": "tblq4eqit0en53dj5db4",
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
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "cm-migrated-at": "2026-11-25T09:18:38Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my09-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my09-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tblq4eqit0en53dj5db4"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my09-ng-influxdb-back-1",
      "uid": "tbl514uefdglqljjbjua",
      "cspResourceName": "tbl514uefdglqljjbjua",
      "cspResourceId": "1161a76a-b53f-4325-b20c-2d4507c069eb",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:25:15",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "133.186.153.17",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.19",
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
      "cspVNetId": "3cf6f879-d1de-49d2-a60e-a08a87423b32",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "8618fe2a-2640-4777-8c22-7f9e6adecef1",
      "networkInterface": "8bd4211b-4e28-41ad-8235-3064ad2875a1",
      "securityGroupIds": [
        "my09-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tb3d1kt2cp0pjrblu83g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-11-25T09:25:20Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:15Z",
          "completedTime": "2026-11-25T09:25:21Z",
          "elapsedTime": 6,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my09-infra101-tbcige1hg3ut5159fb9d",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:22Z",
          "completedTime": "2026-11-25T09:25:22Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbl514uefdglqljjbjua 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my09-ng-influxdb-back-2",
      "uid": "tbapjacflov280qo67i6",
      "cspResourceName": "tbapjacflov280qo67i6",
      "cspResourceId": "d68821e9-9436-4815-9890-63d0f84c5beb",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:25:08",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "125.6.37.131",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.50",
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
      "cspVNetId": "3cf6f879-d1de-49d2-a60e-a08a87423b32",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "8618fe2a-2640-4777-8c22-7f9e6adecef1",
      "networkInterface": "4be5ad23-e23b-49d6-844b-9d37ee4f3017",
      "securityGroupIds": [
        "my09-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tb3d1kt2cp0pjrblu83g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-11-25T09:25:19Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:15Z",
          "completedTime": "2026-11-25T09:25:22Z",
          "elapsedTime": 7,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my09-infra101-tbcige1hg3ut5159fb9d",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:22Z",
          "completedTime": "2026-11-25T09:25:23Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbapjacflov280qo67i6 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my09-ng-ip-10-0-1-30-62161624-1",
      "uid": "tbvu2ueg9vml0gan1d54",
      "cspResourceName": "tbvu2ueg9vml0gan1d54",
      "cspResourceId": "a06b8a06-eb5c-4842-96ff-752c8b4afa2e",
      "name": "my09-ng-ip-10-0-1-30-62161624-1",
      "nodeGroupId": "my09-ng-ip-10-0-1-30-62161624",
      "location": {
        "display": "Pangyo (South Korea)",
        "latitude": 37.390889,
        "longitude": 127.096792
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:25:13",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "125.6.36.178",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.63",
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
      "cspVNetId": "3cf6f879-d1de-49d2-a60e-a08a87423b32",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "8618fe2a-2640-4777-8c22-7f9e6adecef1",
      "networkInterface": "12918d1e-1b05-485a-813a-90439ce580e8",
      "securityGroupIds": [
        "my09-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tb3d1kt2cp0pjrblu83g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-11-25T09:25:20Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:15Z",
          "completedTime": "2026-11-25T09:25:21Z",
          "elapsedTime": 6,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my09-infra101-tbcige1hg3ut5159fb9d",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:22Z",
          "completedTime": "2026-11-25T09:25:22Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbvu2ueg9vml0gan1d54 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
        "my09-ng-influxdb-back",
        "my09-ng-ip-10-0-1-30-62161624"
      ],
      "nodeIds": [
        "my09-ng-influxdb-back-1",
        "my09-ng-influxdb-back-2",
        "my09-ng-ip-10-0-1-30-62161624-1"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my09-ng-influxdb-back",
      "representativeNodeId": "my09-ng-influxdb-back-1"
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
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-ip-10-0-1-30-62161624-1",
            "nodeIp": "125.6.36.178",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbvu2ueg9vml0gan1d54 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-influxdb-back-1",
            "nodeIp": "133.186.153.17",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbl514uefdglqljjbjua 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-influxdb-back-2",
            "nodeIp": "125.6.37.131",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbapjacflov280qo67i6 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postScriptRequestId": "ps-my09-infra101-tbcige1hg3ut5159fb9d"
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
          "status": "Terminating",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
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
          "status": "Terminating",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
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
          "status": "Terminating",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
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
        "cm-migrated-at": "2026-11-25T09:18:38Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my09-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my09-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tblq4eqit0en53dj5db4"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my09-ng-influxdb-back-1",
          "uid": "tbl514uefdglqljjbjua",
          "cspResourceName": "tbl514uefdglqljjbjua",
          "cspResourceId": "1161a76a-b53f-4325-b20c-2d4507c069eb",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:25:15",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "nhn-kr1",
            "sys.createdTime": "2026-11-25 09:25:15",
            "sys.cspResourceId": "1161a76a-b53f-4325-b20c-2d4507c069eb",
            "sys.cspResourceName": "tbl514uefdglqljjbjua",
            "sys.id": "my09-ng-influxdb-back-1",
            "sys.infraId": "my09-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my09-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my09-ng-influxdb-back",
            "sys.subnetId": "my09-mig-subnet-01",
            "sys.uid": "tbl514uefdglqljjbjua",
            "sys.vNetId": "my09-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "133.186.153.17",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.19",
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
          "cspVNetId": "3cf6f879-d1de-49d2-a60e-a08a87423b32",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "8618fe2a-2640-4777-8c22-7f9e6adecef1",
          "networkInterface": "8bd4211b-4e28-41ad-8235-3064ad2875a1",
          "securityGroupIds": [
            "my09-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": "tb3d1kt2cp0pjrblu83g"
        },
        {
          "resourceType": "node",
          "id": "my09-ng-influxdb-back-2",
          "uid": "tbapjacflov280qo67i6",
          "cspResourceName": "tbapjacflov280qo67i6",
          "cspResourceId": "d68821e9-9436-4815-9890-63d0f84c5beb",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:25:08",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "nhn-kr1",
            "sys.createdTime": "2026-11-25 09:25:08",
            "sys.cspResourceId": "d68821e9-9436-4815-9890-63d0f84c5beb",
            "sys.cspResourceName": "tbapjacflov280qo67i6",
            "sys.id": "my09-ng-influxdb-back-2",
            "sys.infraId": "my09-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my09-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my09-ng-influxdb-back",
            "sys.subnetId": "my09-mig-subnet-01",
            "sys.uid": "tbapjacflov280qo67i6",
            "sys.vNetId": "my09-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "125.6.37.131",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.50",
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
          "cspVNetId": "3cf6f879-d1de-49d2-a60e-a08a87423b32",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "8618fe2a-2640-4777-8c22-7f9e6adecef1",
          "networkInterface": "4be5ad23-e23b-49d6-844b-9d37ee4f3017",
          "securityGroupIds": [
            "my09-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": "tb3d1kt2cp0pjrblu83g"
        },
        {
          "resourceType": "node",
          "id": "my09-ng-ip-10-0-1-30-62161624-1",
          "uid": "tbvu2ueg9vml0gan1d54",
          "cspResourceName": "tbvu2ueg9vml0gan1d54",
          "cspResourceId": "a06b8a06-eb5c-4842-96ff-752c8b4afa2e",
          "name": "my09-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my09-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:25:13",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "nhn-kr1",
            "sys.createdTime": "2026-11-25 09:25:13",
            "sys.cspResourceId": "a06b8a06-eb5c-4842-96ff-752c8b4afa2e",
            "sys.cspResourceName": "tbvu2ueg9vml0gan1d54",
            "sys.id": "my09-ng-ip-10-0-1-30-62161624-1",
            "sys.infraId": "my09-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my09-ng-ip-10-0-1-30-62161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my09-ng-ip-10-0-1-30-62161624",
            "sys.subnetId": "my09-mig-subnet-01",
            "sys.uid": "tbvu2ueg9vml0gan1d54",
            "sys.vNetId": "my09-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "125.6.36.178",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.63",
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
          "cspVNetId": "3cf6f879-d1de-49d2-a60e-a08a87423b32",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "8618fe2a-2640-4777-8c22-7f9e6adecef1",
          "networkInterface": "12918d1e-1b05-485a-813a-90439ce580e8",
          "securityGroupIds": [
            "my09-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": "tb3d1kt2cp0pjrblu83g"
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
- **IDs**: [my06-infra101 my08-infra101 my09-infra101]

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
  "uid": "tblq4eqit0en53dj5db4",
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
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "cm-migrated-at": "2026-11-25T09:18:38Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my09-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my09-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tblq4eqit0en53dj5db4"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my09-ng-influxdb-back-1",
      "uid": "tbl514uefdglqljjbjua",
      "cspResourceName": "tbl514uefdglqljjbjua",
      "cspResourceId": "1161a76a-b53f-4325-b20c-2d4507c069eb",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:25:15",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "133.186.153.17",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.19",
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
      "cspVNetId": "3cf6f879-d1de-49d2-a60e-a08a87423b32",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "8618fe2a-2640-4777-8c22-7f9e6adecef1",
      "networkInterface": "8bd4211b-4e28-41ad-8235-3064ad2875a1",
      "securityGroupIds": [
        "my09-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tb3d1kt2cp0pjrblu83g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-11-25T09:25:20Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:15Z",
          "completedTime": "2026-11-25T09:25:21Z",
          "elapsedTime": 6,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my09-infra101-tbcige1hg3ut5159fb9d",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:22Z",
          "completedTime": "2026-11-25T09:25:22Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbl514uefdglqljjbjua 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my09-ng-influxdb-back-2",
      "uid": "tbapjacflov280qo67i6",
      "cspResourceName": "tbapjacflov280qo67i6",
      "cspResourceId": "d68821e9-9436-4815-9890-63d0f84c5beb",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:25:08",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "125.6.37.131",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.50",
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
      "cspVNetId": "3cf6f879-d1de-49d2-a60e-a08a87423b32",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "8618fe2a-2640-4777-8c22-7f9e6adecef1",
      "networkInterface": "4be5ad23-e23b-49d6-844b-9d37ee4f3017",
      "securityGroupIds": [
        "my09-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tb3d1kt2cp0pjrblu83g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-11-25T09:25:19Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:15Z",
          "completedTime": "2026-11-25T09:25:22Z",
          "elapsedTime": 7,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my09-infra101-tbcige1hg3ut5159fb9d",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:22Z",
          "completedTime": "2026-11-25T09:25:23Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbapjacflov280qo67i6 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my09-ng-ip-10-0-1-30-62161624-1",
      "uid": "tbvu2ueg9vml0gan1d54",
      "cspResourceName": "tbvu2ueg9vml0gan1d54",
      "cspResourceId": "a06b8a06-eb5c-4842-96ff-752c8b4afa2e",
      "name": "my09-ng-ip-10-0-1-30-62161624-1",
      "nodeGroupId": "my09-ng-ip-10-0-1-30-62161624",
      "location": {
        "display": "Pangyo (South Korea)",
        "latitude": 37.390889,
        "longitude": 127.096792
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:25:13",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "KR1",
        "zone": "kr-pub-a"
      },
      "publicIP": "125.6.36.178",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.63",
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
      "cspVNetId": "3cf6f879-d1de-49d2-a60e-a08a87423b32",
      "subnetId": "my09-mig-subnet-01",
      "cspSubnetId": "8618fe2a-2640-4777-8c22-7f9e6adecef1",
      "networkInterface": "12918d1e-1b05-485a-813a-90439ce580e8",
      "securityGroupIds": [
        "my09-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my09-mig-sshkey-01",
      "cspSshKeyId": "tb3d1kt2cp0pjrblu83g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGa1g0d+PoJkZEaFabCKVEPhhkamD5Ex+izqPW5kzonD7wps0FxMCb8qqR7WYKRh9c6laQgdJJx+UWxjvFmVEqc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wYX5KgnVCWVtqg3wsST6Fk9FuB3D27rQy6e7Yt3jqmw",
        "firstUsedAt": "2026-11-25T09:25:20Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:15Z",
          "completedTime": "2026-11-25T09:25:21Z",
          "elapsedTime": 6,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my09-infra101-tbcige1hg3ut5159fb9d",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:25:22Z",
          "completedTime": "2026-11-25T09:25:22Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbvu2ueg9vml0gan1d54 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
        "my09-ng-influxdb-back",
        "my09-ng-ip-10-0-1-30-62161624"
      ],
      "nodeIds": [
        "my09-ng-influxdb-back-1",
        "my09-ng-influxdb-back-2",
        "my09-ng-ip-10-0-1-30-62161624-1"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my09-ng-influxdb-back",
      "representativeNodeId": "my09-ng-influxdb-back-1"
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
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-ip-10-0-1-30-62161624-1",
            "nodeIp": "125.6.36.178",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbvu2ueg9vml0gan1d54 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-influxdb-back-1",
            "nodeIp": "133.186.153.17",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbl514uefdglqljjbjua 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my09-infra101",
            "nodeId": "my09-ng-influxdb-back-2",
            "nodeIp": "125.6.37.131",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbapjacflov280qo67i6 5.15.0-164-generic #174-Ubuntu SMP Fri Nov 14 20:25:16 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postScriptRequestId": "ps-my09-infra101-tbcige1hg3ut5159fb9d"
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
      "id": "my09-ng-influxdb-back",
      "cspResourceName": "tbfpvpl7ovve2g0acn9a",
      "cspResourceId": "55b474f6-1ec8-495c-a9a8-a4550305844d",
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
        "ip": "125.6.37.15",
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
      "createdTime": "2026-11-25T09:26:57.965815092Z",
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
          "value": "tbfpvpl7ovve2g0acn9a"
        },
        {
          "key": "Listeners",
          "value": "{default_pool_id:,protocol:,description:,name:,loadbalancers:null,tenant_id:,project_id:,admin_state_up:false,connection_limit:0,keepalive_timeout:0,default_tls_container_ref:,sni_container_refs:null,protocol_port:0,id:14c762c1-8288-41a3-8a0e-6135040d1db5,allowed_cidrs:null,tls_versions:null,timeout_client_data:0,timeout_member_data:0,timeout_member_connect:0,timeout_tcp_inspect:0,insert_headers:null,provisioning_status:,l7policies:null,pools:null}"
        },
        {
          "key": "Pools",
          "value": "{lb_algorithm:,protocol:,description:,admin_state_up:false,tenant_id:,member_port:0,session_persistence:null,healthmonitor_id:,loadbalancers:null,listeners:null,members:null,id:28122adc-d6e8-4f89-8068-a10dcafab11d,name:,project_id:,provider:,provisioning_status:,healthmonitor:{admin_state_up:false,health_check_port:0,delay:0,expected_codes:,max_retries:0,http_method:,timeout:0,pools:null,url_path:,type:,id:,name:,project_id:,max_retries_down:0,host_header:,provisioning_status:}}"
        },
        {
          "key": "VipAddress",
          "value": "10.0.1.66"
        },
        {
          "key": "VipPortID",
          "value": "e1594e56-d7d0-4cf9-831a-23df1a14130c"
        },
        {
          "key": "VipSubnetID",
          "value": "8618fe2a-2640-4777-8c22-7f9e6adecef1"
        },
        {
          "key": "ID",
          "value": "55b474f6-1ec8-495c-a9a8-a4550305844d"
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
    "cspResourceName": "tbfpvpl7ovve2g0acn9a",
    "cspResourceId": "55b474f6-1ec8-495c-a9a8-a4550305844d",
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
      "ip": "125.6.37.15",
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
    "createdTime": "2026-11-25T09:26:57.965815092Z",
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
        "value": "tbfpvpl7ovve2g0acn9a"
      },
      {
        "key": "Listeners",
        "value": "{default_pool_id:,protocol:,description:,name:,loadbalancers:null,tenant_id:,project_id:,admin_state_up:false,connection_limit:0,keepalive_timeout:0,default_tls_container_ref:,sni_container_refs:null,protocol_port:0,id:14c762c1-8288-41a3-8a0e-6135040d1db5,allowed_cidrs:null,tls_versions:null,timeout_client_data:0,timeout_member_data:0,timeout_member_connect:0,timeout_tcp_inspect:0,insert_headers:null,provisioning_status:,l7policies:null,pools:null}"
      },
      {
        "key": "Pools",
        "value": "{lb_algorithm:,protocol:,description:,admin_state_up:false,tenant_id:,member_port:0,session_persistence:null,healthmonitor_id:,loadbalancers:null,listeners:null,members:null,id:28122adc-d6e8-4f89-8068-a10dcafab11d,name:,project_id:,provider:,provisioning_status:,healthmonitor:{admin_state_up:false,health_check_port:0,delay:0,expected_codes:,max_retries:0,http_method:,timeout:0,pools:null,url_path:,type:,id:,name:,project_id:,max_retries_down:0,host_header:,provisioning_status:}}"
      },
      {
        "key": "VipAddress",
        "value": "10.0.1.66"
      },
      {
        "key": "VipPortID",
        "value": "e1594e56-d7d0-4cf9-831a-23df1a14130c"
      },
      {
        "key": "VipSubnetID",
        "value": "8618fe2a-2640-4777-8c22-7f9e6adecef1"
      },
      {
        "key": "ID",
        "value": "55b474f6-1ec8-495c-a9a8-a4550305844d"
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
  "cspResourceName": "tbfpvpl7ovve2g0acn9a",
  "cspResourceId": "55b474f6-1ec8-495c-a9a8-a4550305844d",
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
    "ip": "125.6.37.15",
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
  "createdTime": "2026-11-25T09:26:57.965815092Z",
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
      "value": "tbfpvpl7ovve2g0acn9a"
    },
    {
      "key": "Listeners",
      "value": "{default_pool_id:,protocol:,description:,name:,loadbalancers:null,tenant_id:,project_id:,admin_state_up:false,connection_limit:0,keepalive_timeout:0,default_tls_container_ref:,sni_container_refs:null,protocol_port:0,id:14c762c1-8288-41a3-8a0e-6135040d1db5,allowed_cidrs:null,tls_versions:null,timeout_client_data:0,timeout_member_data:0,timeout_member_connect:0,timeout_tcp_inspect:0,insert_headers:null,provisioning_status:,l7policies:null,pools:null}"
    },
    {
      "key": "Pools",
      "value": "{lb_algorithm:,protocol:,description:,admin_state_up:false,tenant_id:,member_port:0,session_persistence:null,healthmonitor_id:,loadbalancers:null,listeners:null,members:null,id:28122adc-d6e8-4f89-8068-a10dcafab11d,name:,project_id:,provider:,provisioning_status:,healthmonitor:{admin_state_up:false,health_check_port:0,delay:0,expected_codes:,max_retries:0,http_method:,timeout:0,pools:null,url_path:,type:,id:,name:,project_id:,max_retries_down:0,host_header:,provisioning_status:}}"
    },
    {
      "key": "VipAddress",
      "value": "10.0.1.66"
    },
    {
      "key": "VipPortID",
      "value": "e1594e56-d7d0-4cf9-831a-23df1a14130c"
    },
    {
      "key": "VipSubnetID",
      "value": "8618fe2a-2640-4777-8c22-7f9e6adecef1"
    },
    {
      "key": "ID",
      "value": "55b474f6-1ec8-495c-a9a8-a4550305844d"
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

**Summary**: Load balancing confirmed: traffic distributed across 2 unique VMs

### Test Case 12: Target Infrastructure Summary

#### 12.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Target Infrastructure Summary**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:28:42

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
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| r2.c4m16 | 4 | 16.0 | - | x86_64 | default | $0.1954 | 2 |
| c2.c2m2 | 2 | 2.0 | - | x86_64 | default | $0.0422 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| 0f07c795-2a46-44fc-a61b-fa0d96763ce2 | Ubuntu Server 22.04.5 LTS (2026.03.10) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my09-ng-influxdb-back-1 | 1161a76a-b53f-4325-b20c-2d4507c069eb | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 133.186.153.17<br>**Private IP:** 10.0.1.19<br>**SGs:** my09-mig-sg-01<br>**SSH:** my09-mig-sshkey-01 |
| my09-ng-influxdb-back-2 | d68821e9-9436-4815-9890-63d0f84c5beb | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 125.6.37.131<br>**Private IP:** 10.0.1.50<br>**SGs:** my09-mig-sg-01<br>**SSH:** my09-mig-sshkey-01 |
| my09-ng-ip-10-0-1-30-62161624-1 | a06b8a06-eb5c-4842-96ff-752c8b4afa2e | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04.5 LTS (2026.03.10) (Ubuntu Server 22.04.5 LTS (2026.03.10)) | **VNet:** my09-mig-vnet-01<br>**Subnet:** my09-mig-subnet-01<br>**Public IP:** 125.6.36.178<br>**Private IP:** 10.0.1.63<br>**SGs:** my09-mig-sg-02<br>**SSH:** my09-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my09-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my09-mig-vnet-01 |
| **CSP VNet ID** | 3cf6f879-d1de-49d2-a60e-a08a87423b32 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | nhn-kr1 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my09-mig-subnet-01 | 8618fe2a-2640-4777-8c22-7f9e6adecef1 | 10.0.1.0/24 |  |

### Network Load Balancers (NLB)

#### NLB: my09-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my09-ng-influxdb-back |
| **CSP NLB Name** | tbfpvpl7ovve2g0acn9a |
| **CSP NLB ID** | 55b474f6-1ec8-495c-a9a8-a4550305844d |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | nhn-kr1 |
| **Listener** | TCP:9999 (IP: 125.6.37.15) |
| **Target Group** | TCP:8086 (Node Group: my09-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my09-ng-influxdb-back-1 | 1161a76a-b53f-4325-b20c-2d4507c069eb | 10.0.1.19 | 133.186.153.17 | Running |
| my09-ng-influxdb-back-2 | d68821e9-9436-4815-9890-63d0f84c5beb | 10.0.1.50 | 125.6.37.131 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my09-mig-sshkey-01 | tb3d1kt2cp0pjrblu83g | cb-user | 69:ab:70:c9:c6:f4:b1:8f:0b:d1:12:26:ac:a7:e4:6e |

### Security Groups

#### Security Group: my09-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my09-mig-sg-01 |
| **CSP Security Group ID** | 126aeb07-4848-44e6-b4ce-314dfb7240ef |
| **VNet** | my09-mig-vnet-01 |
| **Rule Count** | 8 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 8086 | 10.0.0.0/21 |
| inbound | TCP | 8086 | 10.0.1.0/24 |

#### Security Group: my09-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my09-mig-sg-02 |
| **CSP Security Group ID** | 3b7c1ed5-dca7-4cdb-b18e-a766a72ade33 |
| **VNet** | my09-mig-vnet-01 |
| **Rule Count** | 4 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 9999 | 0.0.0.0/0 |


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
| my09-ng-influxdb-back-1 | r2.c4m16 | $0.1954 | $140.66 |
| my09-ng-influxdb-back-2 | r2.c4m16 | $0.1954 | $140.66 |
| my09-ng-ip-10-0-1-30-62161624-1 | c2.c2m2 | $0.0422 | $30.37 |




### Test Case 13: Migration Report

#### 13.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 13.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-11-25 09:28:45*

---

## 📊 Migration Summary

**Target Cloud:** NHN

**Target Region:** KR1

**Namespace:** mig01 | **Infra ID:** my09-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**NLB Migration Status:** Completed

**Total Load Balancers:** 1

**Migrated Load Balancers:** 1

**💰 Estimated Monthly Cost:** $311.69 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | r2.c4m16, c2.c2m2 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my09-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my09-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 12 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |
| 8 | **Network Load Balancer** | 1 | ✅ Created | my09-ng-influxdb-back (TCP:9999 -> TCP:8086) |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my09-ng-influxdb-back-1<br>**VM ID:** 1161a76a-b53f-4325-b20c-2d4507c069eb<br>**Label(sourceMachineId):** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 |
| 2 | **VM Name:** my09-ng-influxdb-back-2<br>**VM ID:** d68821e9-9436-4815-9890-63d0f84c5beb<br>**Label(sourceMachineId):** ec288dd0-c6fa-8a49-2f60-bc898311febf | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf |
| 3 | **VM Name:** my09-ng-ip-10-0-1-30-62161624-1<br>**VM ID:** a06b8a06-eb5c-4842-96ff-752c8b4afa2e<br>**Label(sourceMachineId):** ec268ed7-821e-9d73-e79f-961262161624 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my09-ng-influxdb-back-1 | **Spec ID:** r2.c4m16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 30 GB |
| 2 | my09-ng-influxdb-back-2 | **Spec ID:** r2.c4m16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 8 GB<br>**Root Disk:** 30 GB |
| 3 | my09-ng-ip-10-0-1-30-62161624-1 | **Spec ID:** c2.c2m2<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 8 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my09-ng-influxdb-back-1 | **Image ID:** 0f07c795-2a46-44fc-a61b-fa0d96763ce2<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Server 22.04.5 LTS (2026.03.10) | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | my09-ng-influxdb-back-2 | **Image ID:** 0f07c795-2a46-44fc-a61b-fa0d96763ce2<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Server 22.04.5 LTS (2026.03.10) | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | my09-ng-ip-10-0-1-30-62161624-1 | **Image ID:** 0f07c795-2a46-44fc-a61b-fa0d96763ce2<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Server 22.04.5 LTS (2026.03.10) | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 12 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my09-mig-sg-01

**CSP ID:** 126aeb07-4848-44e6-b4ce-314dfb7240ef | **VNet:** my09-mig-vnet-01 | **Rules:** 8

**Assigned VMs:**

- **VM:** my09-ng-influxdb-back-1
  - **Source Server:** **Hostname:** ip-10-0-1-221, **Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932
- **VM:** my09-ng-influxdb-back-2
  - **Source Server:** **Hostname:** ip-10-0-1-138, **Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 2 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 4 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |
| 5 | inbound | TCP | 9999 | 0.0.0.0/0 | - | Created by system |
| 6 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 7 | inbound | TCP | 8086 | 10.0.0.0/21 | - | Created by system |
| 8 | inbound | TCP | 8086 | 10.0.1.0/24 | - | Created by system |

### Security Group: my09-mig-sg-02

**CSP ID:** 3b7c1ed5-dca7-4cdb-b18e-a766a72ade33 | **VNet:** my09-mig-vnet-01 | **Rules:** 4

**Assigned VMs:**

- **VM:** my09-ng-ip-10-0-1-30-62161624-1
  - **Source Server:** **Hostname:** ip-10-0-1-30, **Machine ID:** ec268ed7-821e-9d73-e79f-961262161624

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 2 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |
| 3 | inbound | ALL |  | 10.0.0.0/16 | inbound * * from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | TCP | 9999 | 0.0.0.0/0 | inbound tcp 9999 | Migrated from source |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source node network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** my09-mig-vnet-01<br>**ID:** 3cf6f879-d1de-49d2-a60e-a08a87423b32 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my09-mig-subnet-01<br>**ID:** 8618fe2a-2640-4777-8c22-7f9e6adecef1 | 10.0.1.0/24 | my09-mig-vnet-01 |

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

### 1. influxdb_back (haproxy) → my09-ng-influxdb-back

**Configuration Comparison:**

| Property | Source (On-Premise) | Target (Cloud NLB) | Note |
|----------|---------------------|--------------------|------|
| **Load Balancer / Name** | haproxy (Host: ip-10-0-1-30) | my09-ng-influxdb-back (CSP ID: tbfpvpl7ovve2g0acn9a) | Replaced by managed NLB |
| **Type / Scope** | Software LB | PUBLIC / REGION | Cloud-native NLB |
| **Frontend Listener** | *:9999 (TCP) | TCP:9999 | IP: 125.6.37.15 |
| **Backend / Target Group** | influxdb_back (TCP:8086, roundrobin) | my09-ng-influxdb-back (TCP:8086, 2 nodes) | Target Group attached to NodeGroup |
| **Health Check** | Enabled (Interval: 10s, Timeout: 10s, Threshold: 3) | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) | Preserved |

**Backend Node Mappings:**

| No. | Source Backend Server | Source Endpoint | Target Node | Target CSP VM ID | Target Private IP | Target Public IP | Target Port | Status |
|-----|-----------------------|-----------------|-------------|------------------|-------------------|------------------|-------------|--------|
| 1 | influx1 (ip-10-0-1-221) | 10.0.1.221:8086 | my09-ng-influxdb-back-1 | 1161a76a-b53f-4325-b20c-2d4507c069eb | 10.0.1.19 | 133.186.153.17 | 8086 | Running |
| 2 | influx2 (ip-10-0-1-138) | 10.0.1.138:8086 | my09-ng-influxdb-back-2 | d68821e9-9436-4815-9890-63d0f84c5beb | 10.0.1.50 | 125.6.37.131 | 8086 | Running |

**Migration Notes:**
- On-premise haproxy on ip-10-0-1-30 (migrated to my09-ng-ip-10-0-1-30-62161624-1) is replaced by cloud-managed NLB.

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my09-mig-sshkey-01 | tb3d1kt2cp0pjrblu83g | 69:ab:70:c9:c6:f4:b1:8f:0b:d1:12:26:ac:a7:e4:6e | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.4329 |
| Daily | $10.39 |
| Monthly | $311.69 |
| Yearly | $3740.26 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | c2.c2m2 | $30.37 | 9.7% |
| ip-10-0-1-221 (migrated) | r2.c4m16 | $140.66 | 45.1% |
| ip-10-0-1-138 (migrated) | r2.c4m16 | $140.66 | 45.1% |

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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my09-infra101)",
  "success": true
}
```

