# CM-Beetle test results for TENCENT (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with TENCENT cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.2+ (83efe14)
- imdl: v0.1.19+ (83efe14)
- CB-Tumblebug: v0.13.8
- CB-Spider: v0.13.20
- CB-MapUI: v0.13.14
- Target CSP: TENCENT
- Target Region: ap-seoul
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: October 30, 2026
- Test Time: 21:47:24 KST
- Test Execution: 2026-10-30 21:47:24 KST

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

## Test result for TENCENT

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 3.427s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 231ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m25.425s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 33ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 3ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 11ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 3.273s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 21.318s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 4ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 6ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 1m31.261s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.101s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.1s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 1m18.522s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 37.151s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 6m21.242118694s

*Test executed on October 30, 2026 at 21:47:24 KST (2026-10-30 21:47:24 KST) using CM-Beetle automated test CLI*

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
  "desiredCsp": "tencent",
  "desiredRegion": "ap-seoul",
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
        "csp": "tencent",
        "region": "ap-seoul"
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
            "connectionName": "tencent-ap-seoul",
            "specId": "tencent+ap-seoul+s6.large16",
            "imageId": "img-7rotv4ux",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
            "connectionName": "tencent-ap-seoul",
            "specId": "tencent+ap-seoul+sa5.medium2",
            "imageId": "img-487zeit5",
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
        "connectionName": "tencent-ap-seoul",
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
        "connectionName": "tencent-ap-seoul",
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
          "id": "tencent+ap-seoul+s6.large16",
          "uid": "tbo92c4co0ojd6mfkp54",
          "cspSpecName": "S6.LARGE16",
          "name": "tencent+ap-seoul+s6.large16",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.23,
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
          "rootDiskType": "",
          "rootDiskSize": -1,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "S6.LARGE16"
            },
            {
              "key": "InstanceChargeType",
              "value": "SPOTPAID"
            },
            {
              "key": "NetworkCard",
              "value": "100"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "4"
            },
            {
              "key": "Memory",
              "value": "16"
            },
            {
              "key": "InstanceFamily",
              "value": "S6"
            },
            {
              "key": "TypeName",
              "value": "S6"
            },
            {
              "key": "Status",
              "value": "SELL"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.23,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.046,UnitPriceSecondStep:0.23,UnitPriceDiscountSecondStep:0.046,UnitPriceThirdStep:0.23,UnitPriceDiscountThirdStep:0.046}"
            },
            {
              "key": "InstanceBandwidth",
              "value": "4"
            },
            {
              "key": "InstancePps",
              "value": "60"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "Intel Ice Lake"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "2.7GHz/3.3GHz"
            },
            {
              "key": "StatusCategory",
              "value": "EnoughStock"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityReason": "zone(s) [ap-seoul-2] unavailable/SoldOut; available in [ap-seoul-1]",
          "availableZones": [
            "ap-seoul-1"
          ],
          "unavailableZones": [
            "ap-seoul-2"
          ],
          "availabilityCheckedAt": "2026-10-30T12:47:30.630714939Z",
          "availabilityRefreshAfter": "2026-10-30T13:47:30.630714939Z"
        },
        {
          "id": "tencent+ap-seoul+sa5.medium2",
          "uid": "tbhbumjcksmmvgckua7p",
          "cspSpecName": "SA5.MEDIUM2",
          "name": "tencent+ap-seoul+sa5.medium2",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.06,
          "orderInFilteredResult": 15,
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
          "rootDiskSize": -1,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "SA5.MEDIUM2"
            },
            {
              "key": "InstanceChargeType",
              "value": "SPOTPAID"
            },
            {
              "key": "NetworkCard",
              "value": "100"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "Memory",
              "value": "2"
            },
            {
              "key": "InstanceFamily",
              "value": "SA5"
            },
            {
              "key": "TypeName",
              "value": "SA5"
            },
            {
              "key": "Status",
              "value": "SELL"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.06,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.012,UnitPriceSecondStep:0.06,UnitPriceDiscountSecondStep:0.012,UnitPriceThirdStep:0.06,UnitPriceDiscountThirdStep:0.012}"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "25"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "AMD EPYC Bergamo"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "-/3.1GHz"
            },
            {
              "key": "StatusCategory",
              "value": "NormalStock"
            }
          ],
          "availabilityStatus": "Unverified"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "tencent",
          "cspImageName": "img-7rotv4ux",
          "regionList": [
            "ap-bangkok",
            "ap-beijing",
            "ap-chengdu",
            "ap-chongqing",
            "ap-guangzhou",
            "ap-hongkong",
            "ap-jakarta",
            "ap-nanjing",
            "ap-seoul",
            "ap-shanghai",
            "ap-singapore",
            "ap-tokyo",
            "eu-frankfurt",
            "me-saudi-arabia",
            "na-ashburn",
            "na-siliconvalley",
            "sa-saopaulo"
          ],
          "id": "img-7rotv4ux",
          "uid": "tbm5u74berjs1hfbpgol",
          "name": "img-7rotv4ux",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "tencent-sa-saopaulo",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:36 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ImageId",
              "value": "img-7rotv4ux"
            },
            {
              "key": "OsName",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageType",
              "value": "PUBLIC_IMAGE"
            },
            {
              "key": "ImageName",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageDescription",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageSize",
              "value": "20"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "ImageState",
              "value": "NORMAL"
            },
            {
              "key": "Platform",
              "value": "Ubuntu"
            },
            {
              "key": "ImageSource",
              "value": "OFFICIAL"
            },
            {
              "key": "IsSupportCloudinit",
              "value": "true"
            },
            {
              "key": "ImageDeprecated",
              "value": "false"
            }
          ],
          "systemLabel": "",
          "description": "",
          "commandHistory": null
        },
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "tencent",
          "cspImageName": "img-487zeit5",
          "regionList": [
            "ap-bangkok",
            "ap-beijing",
            "ap-chengdu",
            "ap-chongqing",
            "ap-guangzhou",
            "ap-hongkong",
            "ap-jakarta",
            "ap-nanjing",
            "ap-seoul",
            "ap-shanghai",
            "ap-singapore",
            "ap-tokyo",
            "eu-frankfurt",
            "me-saudi-arabia",
            "na-ashburn",
            "na-siliconvalley",
            "sa-saopaulo"
          ],
          "id": "img-487zeit5",
          "uid": "tb0lihs2v8sp4akmc3ss",
          "name": "img-487zeit5",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "tencent-sa-saopaulo",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:36 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Server 22.04 LTS 64bit",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ImageId",
              "value": "img-487zeit5"
            },
            {
              "key": "OsName",
              "value": "Ubuntu Server 22.04 LTS 64bit"
            },
            {
              "key": "ImageType",
              "value": "PUBLIC_IMAGE"
            },
            {
              "key": "ImageName",
              "value": "Ubuntu Server 22.04 LTS 64bit"
            },
            {
              "key": "ImageDescription",
              "value": "Ubuntu Server 22.04 LTS 64bit"
            },
            {
              "key": "ImageSize",
              "value": "20"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "ImageState",
              "value": "NORMAL"
            },
            {
              "key": "Platform",
              "value": "Ubuntu"
            },
            {
              "key": "ImageSource",
              "value": "OFFICIAL"
            },
            {
              "key": "IsSupportCloudinit",
              "value": "true"
            },
            {
              "key": "ImageDeprecated",
              "value": "false"
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
          "connectionName": "tencent-ap-seoul",
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
          "connectionName": "tencent-ap-seoul",
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=66.7% Max=100.0% Avg=88.9% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "tencent",
        "region": "ap-seoul"
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
            "connectionName": "tencent-ap-seoul",
            "specId": "tencent+ap-seoul+sa4.large16",
            "imageId": "img-7rotv4ux",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
            "connectionName": "tencent-ap-seoul",
            "specId": "tencent+ap-seoul+s8.medium2",
            "imageId": "img-487zeit5",
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
        "connectionName": "tencent-ap-seoul",
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
        "connectionName": "tencent-ap-seoul",
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
          "id": "tencent+ap-seoul+sa4.large16",
          "uid": "tbfc3c32gmjude8685o2",
          "cspSpecName": "SA4.LARGE16",
          "name": "tencent+ap-seoul+sa4.large16",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.26,
          "orderInFilteredResult": 12,
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
          "rootDiskSize": -1,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "SA4.LARGE16"
            },
            {
              "key": "InstanceChargeType",
              "value": "SPOTPAID"
            },
            {
              "key": "NetworkCard",
              "value": "100"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "4"
            },
            {
              "key": "Memory",
              "value": "16"
            },
            {
              "key": "InstanceFamily",
              "value": "SA4"
            },
            {
              "key": "TypeName",
              "value": "SA4"
            },
            {
              "key": "Status",
              "value": "SELL"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.26,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.052,UnitPriceSecondStep:0.26,UnitPriceDiscountSecondStep:0.052,UnitPriceThirdStep:0.26,UnitPriceDiscountThirdStep:0.052}"
            },
            {
              "key": "InstanceBandwidth",
              "value": "2.5"
            },
            {
              "key": "InstancePps",
              "value": "40"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "AMD EPYC Genoa"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "-/3.7GHz"
            },
            {
              "key": "StatusCategory",
              "value": "EnoughStock"
            }
          ],
          "availabilityStatus": "Unverified"
        },
        {
          "id": "tencent+ap-seoul+s8.medium2",
          "uid": "tb7bc43289ri6ij8h4i0",
          "cspSpecName": "S8.MEDIUM2",
          "name": "tencent+ap-seoul+s8.medium2",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.08,
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
          "rootDiskSize": -1,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "S8.MEDIUM2"
            },
            {
              "key": "InstanceChargeType",
              "value": "SPOTPAID"
            },
            {
              "key": "NetworkCard",
              "value": "100"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "Memory",
              "value": "2"
            },
            {
              "key": "InstanceFamily",
              "value": "S8"
            },
            {
              "key": "TypeName",
              "value": "S8"
            },
            {
              "key": "Status",
              "value": "SELL"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.08,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.016,UnitPriceSecondStep:0.08,UnitPriceDiscountSecondStep:0.016,UnitPriceThirdStep:0.08,UnitPriceDiscountThirdStep:0.016}"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "40"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "Intel Emerald Rapids"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "-/3.0GHz"
            },
            {
              "key": "StatusCategory",
              "value": "EnoughStock"
            }
          ],
          "availabilityStatus": "Unverified"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "tencent",
          "cspImageName": "img-7rotv4ux",
          "regionList": [
            "ap-bangkok",
            "ap-beijing",
            "ap-chengdu",
            "ap-chongqing",
            "ap-guangzhou",
            "ap-hongkong",
            "ap-jakarta",
            "ap-nanjing",
            "ap-seoul",
            "ap-shanghai",
            "ap-singapore",
            "ap-tokyo",
            "eu-frankfurt",
            "me-saudi-arabia",
            "na-ashburn",
            "na-siliconvalley",
            "sa-saopaulo"
          ],
          "id": "img-7rotv4ux",
          "uid": "tbm5u74berjs1hfbpgol",
          "name": "img-7rotv4ux",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "tencent-sa-saopaulo",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:36 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ImageId",
              "value": "img-7rotv4ux"
            },
            {
              "key": "OsName",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageType",
              "value": "PUBLIC_IMAGE"
            },
            {
              "key": "ImageName",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageDescription",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageSize",
              "value": "20"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "ImageState",
              "value": "NORMAL"
            },
            {
              "key": "Platform",
              "value": "Ubuntu"
            },
            {
              "key": "ImageSource",
              "value": "OFFICIAL"
            },
            {
              "key": "IsSupportCloudinit",
              "value": "true"
            },
            {
              "key": "ImageDeprecated",
              "value": "false"
            }
          ],
          "systemLabel": "",
          "description": "",
          "commandHistory": null
        },
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "tencent",
          "cspImageName": "img-487zeit5",
          "regionList": [
            "ap-bangkok",
            "ap-beijing",
            "ap-chengdu",
            "ap-chongqing",
            "ap-guangzhou",
            "ap-hongkong",
            "ap-jakarta",
            "ap-nanjing",
            "ap-seoul",
            "ap-shanghai",
            "ap-singapore",
            "ap-tokyo",
            "eu-frankfurt",
            "me-saudi-arabia",
            "na-ashburn",
            "na-siliconvalley",
            "sa-saopaulo"
          ],
          "id": "img-487zeit5",
          "uid": "tb0lihs2v8sp4akmc3ss",
          "name": "img-487zeit5",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "tencent-sa-saopaulo",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:36 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Server 22.04 LTS 64bit",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ImageId",
              "value": "img-487zeit5"
            },
            {
              "key": "OsName",
              "value": "Ubuntu Server 22.04 LTS 64bit"
            },
            {
              "key": "ImageType",
              "value": "PUBLIC_IMAGE"
            },
            {
              "key": "ImageName",
              "value": "Ubuntu Server 22.04 LTS 64bit"
            },
            {
              "key": "ImageDescription",
              "value": "Ubuntu Server 22.04 LTS 64bit"
            },
            {
              "key": "ImageSize",
              "value": "20"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "ImageState",
              "value": "NORMAL"
            },
            {
              "key": "Platform",
              "value": "Ubuntu"
            },
            {
              "key": "ImageSource",
              "value": "OFFICIAL"
            },
            {
              "key": "IsSupportCloudinit",
              "value": "true"
            },
            {
              "key": "ImageDeprecated",
              "value": "false"
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
          "connectionName": "tencent-ap-seoul",
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
          "connectionName": "tencent-ap-seoul",
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
    "csp": "tencent",
    "region": "ap-seoul"
  },
  "targetInfra": {
    "name": "my05-infra101",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "NLB-aware recommended infrastructure for cloud migration",
    "nodeGroups": [
      {
        "name": "my05-ng-influxdb-back",
        "nodeGroupSize": 2,
        "label": {
          "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "nlbBackend": "influxdb_back"
        },
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
        "connectionName": "tencent-ap-seoul",
        "specId": "tencent+ap-seoul+s6.large16",
        "imageId": "img-7rotv4ux",
        "vNetId": "my05-mig-vnet-01",
        "subnetId": "my05-mig-subnet-01",
        "securityGroupIds": [
          "my05-mig-sg-01"
        ],
        "sshKeyId": "my05-mig-sshkey-01",
        "rootDiskSize": 50,
        "dataDiskIds": null
      },
      {
        "name": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
        "connectionName": "tencent-ap-seoul",
        "specId": "tencent+ap-seoul+sa5.medium2",
        "imageId": "img-487zeit5",
        "vNetId": "my05-mig-vnet-01",
        "subnetId": "my05-mig-subnet-01",
        "securityGroupIds": [
          "my05-mig-sg-02"
        ],
        "sshKeyId": "my05-mig-sshkey-01",
        "rootDiskSize": 50,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "my05-mig-vnet-01",
    "connectionName": "tencent-ap-seoul",
    "cidrBlock": "10.0.0.0/21",
    "subnetInfoList": [
      {
        "name": "my05-mig-subnet-01",
        "ipv4_CIDR": "10.0.1.0/24",
        "description": "a recommended subnet for migration"
      }
    ],
    "description": "a recommended vNet for migration"
  },
  "targetSshKey": {
    "name": "my05-mig-sshkey-01",
    "connectionName": "tencent-ap-seoul",
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
      "id": "tencent+ap-seoul+s6.large16",
      "uid": "tbo92c4co0ojd6mfkp54",
      "cspSpecName": "S6.LARGE16",
      "name": "tencent+ap-seoul+s6.large16",
      "namespace": "system",
      "connectionName": "tencent-ap-seoul",
      "providerName": "tencent",
      "regionName": "ap-seoul",
      "regionLatitude": 37.566536,
      "regionLongitude": 126.977966,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 16,
      "diskSizeGB": -1,
      "costPerHour": 0.23,
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
      "rootDiskType": "",
      "rootDiskSize": -1,
      "systemLabel": "auto-gen",
      "details": [
        {
          "key": "Zone",
          "value": "ap-seoul-1"
        },
        {
          "key": "InstanceType",
          "value": "S6.LARGE16"
        },
        {
          "key": "InstanceChargeType",
          "value": "SPOTPAID"
        },
        {
          "key": "NetworkCard",
          "value": "100"
        },
        {
          "key": "Externals",
          "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
        },
        {
          "key": "Cpu",
          "value": "4"
        },
        {
          "key": "Memory",
          "value": "16"
        },
        {
          "key": "InstanceFamily",
          "value": "S6"
        },
        {
          "key": "TypeName",
          "value": "S6"
        },
        {
          "key": "Status",
          "value": "SELL"
        },
        {
          "key": "Price",
          "value": "{UnitPrice:0.23,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.046,UnitPriceSecondStep:0.23,UnitPriceDiscountSecondStep:0.046,UnitPriceThirdStep:0.23,UnitPriceDiscountThirdStep:0.046}"
        },
        {
          "key": "InstanceBandwidth",
          "value": "4"
        },
        {
          "key": "InstancePps",
          "value": "60"
        },
        {
          "key": "StorageBlockAmount",
          "value": "0"
        },
        {
          "key": "CpuType",
          "value": "Intel Ice Lake"
        },
        {
          "key": "Gpu",
          "value": "0"
        },
        {
          "key": "Fpga",
          "value": "0"
        },
        {
          "key": "GpuCount",
          "value": "0"
        },
        {
          "key": "Frequency",
          "value": "2.7GHz/3.3GHz"
        },
        {
          "key": "StatusCategory",
          "value": "EnoughStock"
        }
      ],
      "availabilityStatus": "Available",
      "availabilityReason": "zone(s) [ap-seoul-2] unavailable/SoldOut; available in [ap-seoul-1]",
      "availableZones": [
        "ap-seoul-1"
      ],
      "unavailableZones": [
        "ap-seoul-2"
      ],
      "availabilityCheckedAt": "2026-10-30T12:47:30.630714939Z",
      "availabilityRefreshAfter": "2026-10-30T13:47:30.630714939Z"
    },
    {
      "id": "tencent+ap-seoul+sa5.medium2",
      "uid": "tbhbumjcksmmvgckua7p",
      "cspSpecName": "SA5.MEDIUM2",
      "name": "tencent+ap-seoul+sa5.medium2",
      "namespace": "system",
      "connectionName": "tencent-ap-seoul",
      "providerName": "tencent",
      "regionName": "ap-seoul",
      "regionLatitude": 37.566536,
      "regionLongitude": 126.977966,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 2,
      "diskSizeGB": -1,
      "costPerHour": 0.06,
      "orderInFilteredResult": 15,
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
      "rootDiskSize": -1,
      "systemLabel": "auto-gen",
      "details": [
        {
          "key": "Zone",
          "value": "ap-seoul-1"
        },
        {
          "key": "InstanceType",
          "value": "SA5.MEDIUM2"
        },
        {
          "key": "InstanceChargeType",
          "value": "SPOTPAID"
        },
        {
          "key": "NetworkCard",
          "value": "100"
        },
        {
          "key": "Externals",
          "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "Memory",
          "value": "2"
        },
        {
          "key": "InstanceFamily",
          "value": "SA5"
        },
        {
          "key": "TypeName",
          "value": "SA5"
        },
        {
          "key": "Status",
          "value": "SELL"
        },
        {
          "key": "Price",
          "value": "{UnitPrice:0.06,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.012,UnitPriceSecondStep:0.06,UnitPriceDiscountSecondStep:0.012,UnitPriceThirdStep:0.06,UnitPriceDiscountThirdStep:0.012}"
        },
        {
          "key": "InstanceBandwidth",
          "value": "1.5"
        },
        {
          "key": "InstancePps",
          "value": "25"
        },
        {
          "key": "StorageBlockAmount",
          "value": "0"
        },
        {
          "key": "CpuType",
          "value": "AMD EPYC Bergamo"
        },
        {
          "key": "Gpu",
          "value": "0"
        },
        {
          "key": "Fpga",
          "value": "0"
        },
        {
          "key": "GpuCount",
          "value": "0"
        },
        {
          "key": "Frequency",
          "value": "-/3.1GHz"
        },
        {
          "key": "StatusCategory",
          "value": "NormalStock"
        }
      ],
      "availabilityStatus": "Unverified"
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "tencent",
      "cspImageName": "img-7rotv4ux",
      "regionList": [
        "ap-bangkok",
        "ap-beijing",
        "ap-chengdu",
        "ap-chongqing",
        "ap-guangzhou",
        "ap-hongkong",
        "ap-jakarta",
        "ap-nanjing",
        "ap-seoul",
        "ap-shanghai",
        "ap-singapore",
        "ap-tokyo",
        "eu-frankfurt",
        "me-saudi-arabia",
        "na-ashburn",
        "na-siliconvalley",
        "sa-saopaulo"
      ],
      "id": "img-7rotv4ux",
      "uid": "tbm5u74berjs1hfbpgol",
      "name": "img-7rotv4ux",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "tencent-sa-saopaulo",
      "infraType": "",
      "fetchedTime": "2026.08.21 13:57:36 Fri",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI",
      "osDiskType": "NA",
      "osDiskSizeGB": 20,
      "imageStatus": "Available",
      "details": [
        {
          "key": "ImageId",
          "value": "img-7rotv4ux"
        },
        {
          "key": "OsName",
          "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
        },
        {
          "key": "ImageType",
          "value": "PUBLIC_IMAGE"
        },
        {
          "key": "ImageName",
          "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
        },
        {
          "key": "ImageDescription",
          "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
        },
        {
          "key": "ImageSize",
          "value": "20"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "ImageState",
          "value": "NORMAL"
        },
        {
          "key": "Platform",
          "value": "Ubuntu"
        },
        {
          "key": "ImageSource",
          "value": "OFFICIAL"
        },
        {
          "key": "IsSupportCloudinit",
          "value": "true"
        },
        {
          "key": "ImageDeprecated",
          "value": "false"
        }
      ],
      "systemLabel": "",
      "description": "",
      "commandHistory": null
    },
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "tencent",
      "cspImageName": "img-487zeit5",
      "regionList": [
        "ap-bangkok",
        "ap-beijing",
        "ap-chengdu",
        "ap-chongqing",
        "ap-guangzhou",
        "ap-hongkong",
        "ap-jakarta",
        "ap-nanjing",
        "ap-seoul",
        "ap-shanghai",
        "ap-singapore",
        "ap-tokyo",
        "eu-frankfurt",
        "me-saudi-arabia",
        "na-ashburn",
        "na-siliconvalley",
        "sa-saopaulo"
      ],
      "id": "img-487zeit5",
      "uid": "tb0lihs2v8sp4akmc3ss",
      "name": "img-487zeit5",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "tencent-sa-saopaulo",
      "infraType": "",
      "fetchedTime": "2026.08.21 13:57:36 Fri",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Ubuntu Server 22.04 LTS 64bit",
      "osDiskType": "NA",
      "osDiskSizeGB": 20,
      "imageStatus": "Available",
      "details": [
        {
          "key": "ImageId",
          "value": "img-487zeit5"
        },
        {
          "key": "OsName",
          "value": "Ubuntu Server 22.04 LTS 64bit"
        },
        {
          "key": "ImageType",
          "value": "PUBLIC_IMAGE"
        },
        {
          "key": "ImageName",
          "value": "Ubuntu Server 22.04 LTS 64bit"
        },
        {
          "key": "ImageDescription",
          "value": "Ubuntu Server 22.04 LTS 64bit"
        },
        {
          "key": "ImageSize",
          "value": "20"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "ImageState",
          "value": "NORMAL"
        },
        {
          "key": "Platform",
          "value": "Ubuntu"
        },
        {
          "key": "ImageSource",
          "value": "OFFICIAL"
        },
        {
          "key": "IsSupportCloudinit",
          "value": "true"
        },
        {
          "key": "ImageDeprecated",
          "value": "false"
        }
      ],
      "systemLabel": "",
      "description": "",
      "commandHistory": null
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "my05-mig-sg-01",
      "connectionName": "tencent-ap-seoul",
      "vNetId": "my05-mig-vnet-01",
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
      "name": "my05-mig-sg-02",
      "connectionName": "tencent-ap-seoul",
      "vNetId": "my05-mig-vnet-01",
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
- **MCI ID**: `my05-infra101`
- **MCI Name**: `my05-infra101`
- **Status**: `Running:3 (R:3/3)`

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:48:23",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "disk-i2yv7fx5",
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
      "cspSshKeyId": "skey-7ic7ibdp",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBITkZGQHws525lt9X0HxHWisKUXT4aac91VHsa0WYA4F6un5vQUge0+4/pkUCGMMouFRWt8ZeaAsVefDAjAxg6s=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:8vyMr/XIwPVgGY/9n5g1wUK4f4UL1tE6Dj1Nmr/S0NA",
        "firstUsedAt": "2026-10-30T12:48:57Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-10-30T12:48:23Z",
          "completedTime": "2026-10-30T12:48:45Z",
          "elapsedTime": 22,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.8:22 (as \"cb-user\") via bastion Node \"my05-ng-influxdb-back-2\" at 43.166.9.19:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 43.166.9.19:22 as user \"cb-user\" (bastionNodeId=my05-ng-influxdb-back-2): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:56Z",
          "completedTime": "2026-10-30T12:48:58Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my05-infra101-tb7jgohfu1vdvfh1strl",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:00Z",
          "completedTime": "2026-10-30T12:49:02Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-8-ubuntu 5.15.0-190-generic #200-Ubuntu SMP Fri Aug 7 15:06:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:48:20",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "disk-4wy8w515",
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
      "cspSshKeyId": "skey-7ic7ibdp",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBH7eHqHhzfP5DlkwJce0NI2PFbq7mcVQF7Y2v6pzr68EKrF2nD0aixJSH/hTPwQTs1Q5wM5vF21fTc79xIjZYkU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:axKG5jc57e6dPqdXFh6mnXaZ9TKhowx9PmuAOFhDuSY",
        "firstUsedAt": "2026-10-30T12:48:58Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-10-30T12:48:23Z",
          "completedTime": "2026-10-30T12:48:44Z",
          "elapsedTime": 21,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my05-ng-influxdb-back-1\" at 10.0.1.10:22 (as \"cb-user\") via bastion Node \"my05-ng-influxdb-back-2\" at 43.166.9.19:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 43.166.9.19:22 as user \"cb-user\" (bastionNodeId=my05-ng-influxdb-back-2): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:56Z",
          "completedTime": "2026-10-30T12:48:59Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my05-infra101-tb7jgohfu1vdvfh1strl",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:00Z",
          "completedTime": "2026-10-30T12:49:02Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-10-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:48:20",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "disk-2edepf6f",
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
      "cspSshKeyId": "skey-7ic7ibdp",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGLYVBOl8ArO1dzdiQubuOcVn7X2JeFiXmohk5nDDkafbJ4CLFCgSz8twDHm92Ct++K9pO7WcFKaAxMM1ESNDKY=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:Y4QjrxZ80G7Nw8opYJ70ttPXH6OUWpFd+vFz6NDmUv0",
        "firstUsedAt": "2026-10-30T12:48:35Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:23Z",
          "completedTime": "2026-10-30T12:48:46Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-10-30T12:49:00Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb7jgohfu1vdvfh1strl",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:00Z",
          "completedTime": "2026-10-30T12:49:03Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-2-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my05-mig-vnet-01",
      "name": "my05-mig-vnet-01",
      "infraId": "my05-infra101",
      "vNetId": "my05-mig-vnet-01",
      "connectionNames": [
        "tencent-ap-seoul"
      ],
      "providerNames": [
        "tencent"
      ],
      "regionNames": [
        "ap-seoul"
      ],
      "nodeGroupIds": [
        "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my05-ng-influxdb-back"
      ],
      "nodeIds": [
        "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my05-ng-influxdb-back-1",
        "my05-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-influxdb-back-1",
            "nodeIp": "43.131.254.242",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-10-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "43.166.11.218",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-8-ubuntu 5.15.0-190-generic #200-Ubuntu SMP Fri Aug 7 15:06:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-influxdb-back-2",
            "nodeIp": "43.166.9.19",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-2-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my05-infra101-tb7jgohfu1vdvfh1strl"
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
- **Count**: 6

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
      "label": null,
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
      "id": "my04-infra101",
      "uid": "tbqh2604smo06ces9d4e",
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
        "cm-migrated-at": "2026-10-30T12:47:22Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my04-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my04-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbqh2604smo06ces9d4e"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbck9ft8t7h92en7cc5d",
          "cspResourceName": "tbck9ft8t7h92en7cc5d",
          "cspResourceId": "i-mj7c2uilk7pplm0pa0qo",
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
          "createdTime": "2026-10-30 12:47:32",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-10-30 12:47:32",
            "sys.cspResourceId": "i-mj7c2uilk7pplm0pa0qo",
            "sys.cspResourceName": "tbck9ft8t7h92en7cc5d",
            "sys.id": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbck9ft8t7h92en7cc5d",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.220.204.23",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.142",
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
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7xwzw3fbv61c7bto7le",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj77bwqspdjcxy9d4i88j",
          "networkInterface": "eni-mj7c2uilk7pplm0qn8yj",
          "securityGroupIds": [
            "my04-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tb5aoempt96t422gdonc"
        },
        {
          "resourceType": "node",
          "id": "my04-ng-influxdb-back-1",
          "uid": "tbq36qlq7g4k5h8pd3iq",
          "cspResourceName": "tbq36qlq7g4k5h8pd3iq",
          "cspResourceId": "i-mj7hb03d1frhdo6k7utt",
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
          "createdTime": "2026-10-30 12:47:33",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-10-30 12:47:33",
            "sys.cspResourceId": "i-mj7hb03d1frhdo6k7utt",
            "sys.cspResourceName": "tbq36qlq7g4k5h8pd3iq",
            "sys.id": "my04-ng-influxdb-back-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-influxdb-back",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbq36qlq7g4k5h8pd3iq",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.213.158.198",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.141",
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
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7xwzw3fbv61c7bto7le",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj77bwqspdjcxy9d4i88j",
          "networkInterface": "eni-mj7hb03d1frhdo6m7drv",
          "securityGroupIds": [
            "my04-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tb5aoempt96t422gdonc"
        },
        {
          "resourceType": "node",
          "id": "my04-ng-influxdb-back-2",
          "uid": "tbnkp7imlebkp8bhh3os",
          "cspResourceName": "tbnkp7imlebkp8bhh3os",
          "cspResourceId": "i-mj7hb03d1frhdo6k7utu",
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
          "createdTime": "2026-10-30 12:47:34",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-10-30 12:47:34",
            "sys.cspResourceId": "i-mj7hb03d1frhdo6k7utu",
            "sys.cspResourceName": "tbnkp7imlebkp8bhh3os",
            "sys.id": "my04-ng-influxdb-back-2",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-influxdb-back",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbnkp7imlebkp8bhh3os",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.21.196",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.140",
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
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7xwzw3fbv61c7bto7le",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj77bwqspdjcxy9d4i88j",
          "networkInterface": "eni-mj7hb03d1frhdo6m7dry",
          "securityGroupIds": [
            "my04-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tb5aoempt96t422gdonc"
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
          "id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbar0eeronn7me24e1jr",
          "name": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
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
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
          "cspVNetId": "",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my06-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-1",
          "uid": "tbf0cqgntsded5k3h3ln",
          "name": "my06-ng-influxdb-back-1",
          "nodeGroupId": "my06-ng-influxdb-back",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
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
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
          "cspVNetId": "",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-2",
          "uid": "tbvicb2g6blfup7t2rn1",
          "name": "my06-ng-influxdb-back-2",
          "nodeGroupId": "my06-ng-influxdb-back",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
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
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
          "cspVNetId": "",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
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
- **IDs**: [my01-infra101 my02-infra101 my03-infra101 my04-infra101 my05-infra101 my06-infra101]

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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:48:23",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "disk-i2yv7fx5",
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
      "cspSshKeyId": "skey-7ic7ibdp",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBITkZGQHws525lt9X0HxHWisKUXT4aac91VHsa0WYA4F6un5vQUge0+4/pkUCGMMouFRWt8ZeaAsVefDAjAxg6s=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:8vyMr/XIwPVgGY/9n5g1wUK4f4UL1tE6Dj1Nmr/S0NA",
        "firstUsedAt": "2026-10-30T12:48:57Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-10-30T12:48:23Z",
          "completedTime": "2026-10-30T12:48:45Z",
          "elapsedTime": 22,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.8:22 (as \"cb-user\") via bastion Node \"my05-ng-influxdb-back-2\" at 43.166.9.19:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 43.166.9.19:22 as user \"cb-user\" (bastionNodeId=my05-ng-influxdb-back-2): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:56Z",
          "completedTime": "2026-10-30T12:48:58Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my05-infra101-tb7jgohfu1vdvfh1strl",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:00Z",
          "completedTime": "2026-10-30T12:49:02Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-8-ubuntu 5.15.0-190-generic #200-Ubuntu SMP Fri Aug 7 15:06:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:48:20",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "disk-4wy8w515",
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
      "cspSshKeyId": "skey-7ic7ibdp",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBH7eHqHhzfP5DlkwJce0NI2PFbq7mcVQF7Y2v6pzr68EKrF2nD0aixJSH/hTPwQTs1Q5wM5vF21fTc79xIjZYkU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:axKG5jc57e6dPqdXFh6mnXaZ9TKhowx9PmuAOFhDuSY",
        "firstUsedAt": "2026-10-30T12:48:58Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-10-30T12:48:23Z",
          "completedTime": "2026-10-30T12:48:44Z",
          "elapsedTime": 21,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my05-ng-influxdb-back-1\" at 10.0.1.10:22 (as \"cb-user\") via bastion Node \"my05-ng-influxdb-back-2\" at 43.166.9.19:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 43.166.9.19:22 as user \"cb-user\" (bastionNodeId=my05-ng-influxdb-back-2): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:56Z",
          "completedTime": "2026-10-30T12:48:59Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my05-infra101-tb7jgohfu1vdvfh1strl",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:00Z",
          "completedTime": "2026-10-30T12:49:02Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-10-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:48:20",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "disk-2edepf6f",
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
      "cspSshKeyId": "skey-7ic7ibdp",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGLYVBOl8ArO1dzdiQubuOcVn7X2JeFiXmohk5nDDkafbJ4CLFCgSz8twDHm92Ct++K9pO7WcFKaAxMM1ESNDKY=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:Y4QjrxZ80G7Nw8opYJ70ttPXH6OUWpFd+vFz6NDmUv0",
        "firstUsedAt": "2026-10-30T12:48:35Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:23Z",
          "completedTime": "2026-10-30T12:48:46Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-10-30T12:49:00Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb7jgohfu1vdvfh1strl",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:00Z",
          "completedTime": "2026-10-30T12:49:03Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-2-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my05-mig-vnet-01",
      "name": "my05-mig-vnet-01",
      "infraId": "my05-infra101",
      "vNetId": "my05-mig-vnet-01",
      "connectionNames": [
        "tencent-ap-seoul"
      ],
      "providerNames": [
        "tencent"
      ],
      "regionNames": [
        "ap-seoul"
      ],
      "nodeGroupIds": [
        "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my05-ng-influxdb-back"
      ],
      "nodeIds": [
        "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my05-ng-influxdb-back-1",
        "my05-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-influxdb-back-1",
            "nodeIp": "43.131.254.242",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-10-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "43.166.11.218",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-8-ubuntu 5.15.0-190-generic #200-Ubuntu SMP Fri Aug 7 15:06:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-influxdb-back-2",
            "nodeIp": "43.166.9.19",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-2-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my05-infra101-tb7jgohfu1vdvfh1strl"
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
      "id": "my05-ng-influxdb-back",
      "cspResourceName": "tb02oudm3htkij1pcl4k",
      "cspResourceId": "lb-ac8ouxdn",
      "name": "my05-ng-influxdb-back",
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
      "type": "PUBLIC",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "port": "9999",
        "dnsName": "lb-ac8ouxdn-5tbe8u7opci1wahm.clb.kr-tencentclb.com",
        "keyValueList": [
          {
            "key": "LoadBalancerId",
            "value": "lb-ac8ouxdn"
          },
          {
            "key": "LoadBalancerName",
            "value": "tb02oudm3htkij1pcl4k"
          },
          {
            "key": "LoadBalancerType",
            "value": "OPEN"
          },
          {
            "key": "Forward",
            "value": "1"
          },
          {
            "key": "Domain",
            "value": "lb-ac8ouxdn-5tbe8u7opci1wahm.clb.kr-tencentclb.com"
          },
          {
            "key": "Status",
            "value": "1"
          },
          {
            "key": "CreateTime",
            "value": "2026-10-30 20:49:04"
          },
          {
            "key": "StatusTime",
            "value": "2026-10-30 20:49:13"
          },
          {
            "key": "ProjectId",
            "value": "0"
          },
          {
            "key": "VpcId",
            "value": "vpc-rtznlhxm"
          },
          {
            "key": "OpenBgp",
            "value": "0"
          },
          {
            "key": "Snat",
            "value": "false"
          },
          {
            "key": "Isolation",
            "value": "0"
          },
          {
            "key": "TargetRegionInfo",
            "value": "{Region:ap-seoul,VpcId:vpc-rtznlhxm}"
          },
          {
            "key": "AddressIPVersion",
            "value": "ipv4"
          },
          {
            "key": "NumericalVpcId",
            "value": "15840670"
          },
          {
            "key": "VipIsp",
            "value": "BGP"
          },
          {
            "key": "MasterZone",
            "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
          },
          {
            "key": "BackupZoneSet",
            "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
          },
          {
            "key": "ExpireTime",
            "value": "2026-11-30 20:49:04"
          },
          {
            "key": "ChargeType",
            "value": "POSTPAID_BY_HOUR"
          },
          {
            "key": "NetworkAttributes",
            "value": "{InternetChargeType:TRAFFIC_POSTPAID_BY_HOUR,InternetMaxBandwidthOut:10}"
          },
          {
            "key": "IsDDos",
            "value": "false"
          },
          {
            "key": "LoadBalancerPassToTarget",
            "value": "false"
          },
          {
            "key": "ExclusiveCluster",
            "value": "{}"
          },
          {
            "key": "SnatPro",
            "value": "false"
          },
          {
            "key": "SlaType",
            "value": "clb.c2.medium"
          },
          {
            "key": "IsBlock",
            "value": "false"
          },
          {
            "key": "LocalBgp",
            "value": "false"
          },
          {
            "key": "MixIpTarget",
            "value": "false"
          },
          {
            "key": "AttributeFlags",
            "value": "IsDomainCLB"
          }
        ]
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my05-ng-influxdb-back",
        "nodes": [
          "my05-ng-influxdb-back-1",
          "my05-ng-influxdb-back-2"
        ]
      },
      "healthChecker": {
        "protocol": "TCP",
        "port": "8086",
        "interval": 10,
        "threshold": 3,
        "timeout": 10,
        "keyValueList": [
          {
            "key": "HealthSwitch",
            "value": "1"
          },
          {
            "key": "TimeOut",
            "value": "10"
          },
          {
            "key": "IntervalTime",
            "value": "10"
          },
          {
            "key": "HealthNum",
            "value": "3"
          },
          {
            "key": "UnHealthNum",
            "value": "3"
          },
          {
            "key": "CheckPort",
            "value": "8086"
          },
          {
            "key": "CheckType",
            "value": "TCP"
          },
          {
            "key": "SourceIpType",
            "value": "1"
          }
        ]
      },
      "createdTime": "2026-10-30T20:49:04Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "LoadBalancerId",
          "value": "lb-ac8ouxdn"
        },
        {
          "key": "LoadBalancerName",
          "value": "tb02oudm3htkij1pcl4k"
        },
        {
          "key": "LoadBalancerType",
          "value": "OPEN"
        },
        {
          "key": "Forward",
          "value": "1"
        },
        {
          "key": "Domain",
          "value": "lb-ac8ouxdn-5tbe8u7opci1wahm.clb.kr-tencentclb.com"
        },
        {
          "key": "Status",
          "value": "1"
        },
        {
          "key": "CreateTime",
          "value": "2026-10-30 20:49:04"
        },
        {
          "key": "StatusTime",
          "value": "2026-10-30 20:49:13"
        },
        {
          "key": "ProjectId",
          "value": "0"
        },
        {
          "key": "VpcId",
          "value": "vpc-rtznlhxm"
        },
        {
          "key": "OpenBgp",
          "value": "0"
        },
        {
          "key": "Snat",
          "value": "false"
        },
        {
          "key": "Isolation",
          "value": "0"
        },
        {
          "key": "TargetRegionInfo",
          "value": "{Region:ap-seoul,VpcId:vpc-rtznlhxm}"
        },
        {
          "key": "AddressIPVersion",
          "value": "ipv4"
        },
        {
          "key": "NumericalVpcId",
          "value": "15840670"
        },
        {
          "key": "VipIsp",
          "value": "BGP"
        },
        {
          "key": "MasterZone",
          "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
        },
        {
          "key": "BackupZoneSet",
          "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
        },
        {
          "key": "ExpireTime",
          "value": "2026-11-30 20:49:04"
        },
        {
          "key": "ChargeType",
          "value": "POSTPAID_BY_HOUR"
        },
        {
          "key": "NetworkAttributes",
          "value": "{InternetChargeType:TRAFFIC_POSTPAID_BY_HOUR,InternetMaxBandwidthOut:10}"
        },
        {
          "key": "IsDDos",
          "value": "false"
        },
        {
          "key": "LoadBalancerPassToTarget",
          "value": "false"
        },
        {
          "key": "ExclusiveCluster",
          "value": "{}"
        },
        {
          "key": "SnatPro",
          "value": "false"
        },
        {
          "key": "SlaType",
          "value": "clb.c2.medium"
        },
        {
          "key": "IsBlock",
          "value": "false"
        },
        {
          "key": "LocalBgp",
          "value": "false"
        },
        {
          "key": "MixIpTarget",
          "value": "false"
        },
        {
          "key": "AttributeFlags",
          "value": "IsDomainCLB"
        }
      ],
      "isAutoGenerated": false,
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.566536,
        "longitude": 126.977966
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
    "id": "my05-ng-influxdb-back",
    "cspResourceName": "tb02oudm3htkij1pcl4k",
    "cspResourceId": "lb-ac8ouxdn",
    "name": "my05-ng-influxdb-back",
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
    "type": "PUBLIC",
    "scope": "REGION",
    "listener": {
      "protocol": "TCP",
      "port": "9999",
      "dnsName": "lb-ac8ouxdn-5tbe8u7opci1wahm.clb.kr-tencentclb.com",
      "keyValueList": [
        {
          "key": "LoadBalancerId",
          "value": "lb-ac8ouxdn"
        },
        {
          "key": "LoadBalancerName",
          "value": "tb02oudm3htkij1pcl4k"
        },
        {
          "key": "LoadBalancerType",
          "value": "OPEN"
        },
        {
          "key": "Forward",
          "value": "1"
        },
        {
          "key": "Domain",
          "value": "lb-ac8ouxdn-5tbe8u7opci1wahm.clb.kr-tencentclb.com"
        },
        {
          "key": "Status",
          "value": "1"
        },
        {
          "key": "CreateTime",
          "value": "2026-10-30 20:49:04"
        },
        {
          "key": "StatusTime",
          "value": "2026-10-30 20:49:13"
        },
        {
          "key": "ProjectId",
          "value": "0"
        },
        {
          "key": "VpcId",
          "value": "vpc-rtznlhxm"
        },
        {
          "key": "OpenBgp",
          "value": "0"
        },
        {
          "key": "Snat",
          "value": "false"
        },
        {
          "key": "Isolation",
          "value": "0"
        },
        {
          "key": "TargetRegionInfo",
          "value": "{Region:ap-seoul,VpcId:vpc-rtznlhxm}"
        },
        {
          "key": "AddressIPVersion",
          "value": "ipv4"
        },
        {
          "key": "NumericalVpcId",
          "value": "15840670"
        },
        {
          "key": "VipIsp",
          "value": "BGP"
        },
        {
          "key": "MasterZone",
          "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
        },
        {
          "key": "BackupZoneSet",
          "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
        },
        {
          "key": "ExpireTime",
          "value": "2026-11-30 20:49:04"
        },
        {
          "key": "ChargeType",
          "value": "POSTPAID_BY_HOUR"
        },
        {
          "key": "NetworkAttributes",
          "value": "{InternetChargeType:TRAFFIC_POSTPAID_BY_HOUR,InternetMaxBandwidthOut:10}"
        },
        {
          "key": "IsDDos",
          "value": "false"
        },
        {
          "key": "LoadBalancerPassToTarget",
          "value": "false"
        },
        {
          "key": "ExclusiveCluster",
          "value": "{}"
        },
        {
          "key": "SnatPro",
          "value": "false"
        },
        {
          "key": "SlaType",
          "value": "clb.c2.medium"
        },
        {
          "key": "IsBlock",
          "value": "false"
        },
        {
          "key": "LocalBgp",
          "value": "false"
        },
        {
          "key": "MixIpTarget",
          "value": "false"
        },
        {
          "key": "AttributeFlags",
          "value": "IsDomainCLB"
        }
      ]
    },
    "targetGroup": {
      "protocol": "TCP",
      "port": "8086",
      "nodeGroupId": "my05-ng-influxdb-back",
      "nodes": [
        "my05-ng-influxdb-back-1",
        "my05-ng-influxdb-back-2"
      ]
    },
    "healthChecker": {
      "protocol": "TCP",
      "port": "8086",
      "interval": 10,
      "threshold": 3,
      "timeout": 10,
      "keyValueList": [
        {
          "key": "HealthSwitch",
          "value": "1"
        },
        {
          "key": "TimeOut",
          "value": "10"
        },
        {
          "key": "IntervalTime",
          "value": "10"
        },
        {
          "key": "HealthNum",
          "value": "3"
        },
        {
          "key": "UnHealthNum",
          "value": "3"
        },
        {
          "key": "CheckPort",
          "value": "8086"
        },
        {
          "key": "CheckType",
          "value": "TCP"
        },
        {
          "key": "SourceIpType",
          "value": "1"
        }
      ]
    },
    "createdTime": "2026-10-30T20:49:04Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "LoadBalancerId",
        "value": "lb-ac8ouxdn"
      },
      {
        "key": "LoadBalancerName",
        "value": "tb02oudm3htkij1pcl4k"
      },
      {
        "key": "LoadBalancerType",
        "value": "OPEN"
      },
      {
        "key": "Forward",
        "value": "1"
      },
      {
        "key": "Domain",
        "value": "lb-ac8ouxdn-5tbe8u7opci1wahm.clb.kr-tencentclb.com"
      },
      {
        "key": "Status",
        "value": "1"
      },
      {
        "key": "CreateTime",
        "value": "2026-10-30 20:49:04"
      },
      {
        "key": "StatusTime",
        "value": "2026-10-30 20:49:13"
      },
      {
        "key": "ProjectId",
        "value": "0"
      },
      {
        "key": "VpcId",
        "value": "vpc-rtznlhxm"
      },
      {
        "key": "OpenBgp",
        "value": "0"
      },
      {
        "key": "Snat",
        "value": "false"
      },
      {
        "key": "Isolation",
        "value": "0"
      },
      {
        "key": "TargetRegionInfo",
        "value": "{Region:ap-seoul,VpcId:vpc-rtznlhxm}"
      },
      {
        "key": "AddressIPVersion",
        "value": "ipv4"
      },
      {
        "key": "NumericalVpcId",
        "value": "15840670"
      },
      {
        "key": "VipIsp",
        "value": "BGP"
      },
      {
        "key": "MasterZone",
        "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
      },
      {
        "key": "BackupZoneSet",
        "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
      },
      {
        "key": "ExpireTime",
        "value": "2026-11-30 20:49:04"
      },
      {
        "key": "ChargeType",
        "value": "POSTPAID_BY_HOUR"
      },
      {
        "key": "NetworkAttributes",
        "value": "{InternetChargeType:TRAFFIC_POSTPAID_BY_HOUR,InternetMaxBandwidthOut:10}"
      },
      {
        "key": "IsDDos",
        "value": "false"
      },
      {
        "key": "LoadBalancerPassToTarget",
        "value": "false"
      },
      {
        "key": "ExclusiveCluster",
        "value": "{}"
      },
      {
        "key": "SnatPro",
        "value": "false"
      },
      {
        "key": "SlaType",
        "value": "clb.c2.medium"
      },
      {
        "key": "IsBlock",
        "value": "false"
      },
      {
        "key": "LocalBgp",
        "value": "false"
      },
      {
        "key": "MixIpTarget",
        "value": "false"
      },
      {
        "key": "AttributeFlags",
        "value": "IsDomainCLB"
      }
    ],
    "isAutoGenerated": false,
    "location": {
      "display": "South Korea (Seoul)",
      "latitude": 37.566536,
      "longitude": 126.977966
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
  "id": "my05-ng-influxdb-back",
  "cspResourceName": "tb02oudm3htkij1pcl4k",
  "cspResourceId": "lb-ac8ouxdn",
  "name": "my05-ng-influxdb-back",
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
  "type": "PUBLIC",
  "scope": "REGION",
  "listener": {
    "protocol": "TCP",
    "port": "9999",
    "dnsName": "lb-ac8ouxdn-5tbe8u7opci1wahm.clb.kr-tencentclb.com",
    "keyValueList": [
      {
        "key": "LoadBalancerId",
        "value": "lb-ac8ouxdn"
      },
      {
        "key": "LoadBalancerName",
        "value": "tb02oudm3htkij1pcl4k"
      },
      {
        "key": "LoadBalancerType",
        "value": "OPEN"
      },
      {
        "key": "Forward",
        "value": "1"
      },
      {
        "key": "Domain",
        "value": "lb-ac8ouxdn-5tbe8u7opci1wahm.clb.kr-tencentclb.com"
      },
      {
        "key": "Status",
        "value": "1"
      },
      {
        "key": "CreateTime",
        "value": "2026-10-30 20:49:04"
      },
      {
        "key": "StatusTime",
        "value": "2026-10-30 20:49:13"
      },
      {
        "key": "ProjectId",
        "value": "0"
      },
      {
        "key": "VpcId",
        "value": "vpc-rtznlhxm"
      },
      {
        "key": "OpenBgp",
        "value": "0"
      },
      {
        "key": "Snat",
        "value": "false"
      },
      {
        "key": "Isolation",
        "value": "0"
      },
      {
        "key": "TargetRegionInfo",
        "value": "{Region:ap-seoul,VpcId:vpc-rtznlhxm}"
      },
      {
        "key": "AddressIPVersion",
        "value": "ipv4"
      },
      {
        "key": "NumericalVpcId",
        "value": "15840670"
      },
      {
        "key": "VipIsp",
        "value": "BGP"
      },
      {
        "key": "MasterZone",
        "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
      },
      {
        "key": "BackupZoneSet",
        "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
      },
      {
        "key": "ExpireTime",
        "value": "2026-11-30 20:49:04"
      },
      {
        "key": "ChargeType",
        "value": "POSTPAID_BY_HOUR"
      },
      {
        "key": "NetworkAttributes",
        "value": "{InternetChargeType:TRAFFIC_POSTPAID_BY_HOUR,InternetMaxBandwidthOut:10}"
      },
      {
        "key": "IsDDos",
        "value": "false"
      },
      {
        "key": "LoadBalancerPassToTarget",
        "value": "false"
      },
      {
        "key": "ExclusiveCluster",
        "value": "{}"
      },
      {
        "key": "SnatPro",
        "value": "false"
      },
      {
        "key": "SlaType",
        "value": "clb.c2.medium"
      },
      {
        "key": "IsBlock",
        "value": "false"
      },
      {
        "key": "LocalBgp",
        "value": "false"
      },
      {
        "key": "MixIpTarget",
        "value": "false"
      },
      {
        "key": "AttributeFlags",
        "value": "IsDomainCLB"
      }
    ]
  },
  "targetGroup": {
    "protocol": "TCP",
    "port": "8086",
    "nodeGroupId": "my05-ng-influxdb-back",
    "nodes": [
      "my05-ng-influxdb-back-1",
      "my05-ng-influxdb-back-2"
    ]
  },
  "healthChecker": {
    "protocol": "TCP",
    "port": "8086",
    "interval": 10,
    "threshold": 3,
    "timeout": 10,
    "keyValueList": [
      {
        "key": "HealthSwitch",
        "value": "1"
      },
      {
        "key": "TimeOut",
        "value": "10"
      },
      {
        "key": "IntervalTime",
        "value": "10"
      },
      {
        "key": "HealthNum",
        "value": "3"
      },
      {
        "key": "UnHealthNum",
        "value": "3"
      },
      {
        "key": "CheckPort",
        "value": "8086"
      },
      {
        "key": "CheckType",
        "value": "TCP"
      },
      {
        "key": "SourceIpType",
        "value": "1"
      }
    ]
  },
  "createdTime": "2026-10-30T20:49:04Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "LoadBalancerId",
      "value": "lb-ac8ouxdn"
    },
    {
      "key": "LoadBalancerName",
      "value": "tb02oudm3htkij1pcl4k"
    },
    {
      "key": "LoadBalancerType",
      "value": "OPEN"
    },
    {
      "key": "Forward",
      "value": "1"
    },
    {
      "key": "Domain",
      "value": "lb-ac8ouxdn-5tbe8u7opci1wahm.clb.kr-tencentclb.com"
    },
    {
      "key": "Status",
      "value": "1"
    },
    {
      "key": "CreateTime",
      "value": "2026-10-30 20:49:04"
    },
    {
      "key": "StatusTime",
      "value": "2026-10-30 20:49:13"
    },
    {
      "key": "ProjectId",
      "value": "0"
    },
    {
      "key": "VpcId",
      "value": "vpc-rtznlhxm"
    },
    {
      "key": "OpenBgp",
      "value": "0"
    },
    {
      "key": "Snat",
      "value": "false"
    },
    {
      "key": "Isolation",
      "value": "0"
    },
    {
      "key": "TargetRegionInfo",
      "value": "{Region:ap-seoul,VpcId:vpc-rtznlhxm}"
    },
    {
      "key": "AddressIPVersion",
      "value": "ipv4"
    },
    {
      "key": "NumericalVpcId",
      "value": "15840670"
    },
    {
      "key": "VipIsp",
      "value": "BGP"
    },
    {
      "key": "MasterZone",
      "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
    },
    {
      "key": "BackupZoneSet",
      "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
    },
    {
      "key": "ExpireTime",
      "value": "2026-11-30 20:49:04"
    },
    {
      "key": "ChargeType",
      "value": "POSTPAID_BY_HOUR"
    },
    {
      "key": "NetworkAttributes",
      "value": "{InternetChargeType:TRAFFIC_POSTPAID_BY_HOUR,InternetMaxBandwidthOut:10}"
    },
    {
      "key": "IsDDos",
      "value": "false"
    },
    {
      "key": "LoadBalancerPassToTarget",
      "value": "false"
    },
    {
      "key": "ExclusiveCluster",
      "value": "{}"
    },
    {
      "key": "SnatPro",
      "value": "false"
    },
    {
      "key": "SlaType",
      "value": "clb.c2.medium"
    },
    {
      "key": "IsBlock",
      "value": "false"
    },
    {
      "key": "LocalBgp",
      "value": "false"
    },
    {
      "key": "MixIpTarget",
      "value": "false"
    },
    {
      "key": "AttributeFlags",
      "value": "IsDomainCLB"
    }
  ],
  "isAutoGenerated": false,
  "location": {
    "display": "South Korea (Seoul)",
    "latitude": 37.566536,
    "longitude": 126.977966
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

**Generated At:** 2026-10-30 12:51:45

**Namespace:** mig01

**Infra Name:** my05-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my05-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | TENCENT |
| **Target Region** | ap-seoul |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| S6.LARGE16 | 4 | 16.0 | - | x86_64 |  | $0.2300 | 2 |
| SA5.MEDIUM2 | 2 | 2.0 | - | x86_64 |  | $0.0600 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| img-487zeit5 | Ubuntu Server 22.04 LTS 64bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 1 |
| img-7rotv4ux | Ubuntu Server 22.04 LTS 64bit UEFI | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 2 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | ins-mhkac4j9 | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04 LTS 64bit (Ubuntu Server 22.04 LTS 64bit) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.166.11.218<br>**Private IP:** 10.0.1.8<br>**SGs:** my05-mig-sg-02<br>**SSH:** my05-mig-sshkey-01 |
| my05-ng-influxdb-back-1 | ins-4ap62gvz | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.131.254.242<br>**Private IP:** 10.0.1.10<br>**SGs:** my05-mig-sg-01<br>**SSH:** my05-mig-sshkey-01 |
| my05-ng-influxdb-back-2 | ins-q40emar7 | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.166.9.19<br>**Private IP:** 10.0.1.2<br>**SGs:** my05-mig-sg-01<br>**SSH:** my05-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my05-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my05-mig-vnet-01 |
| **CSP VNet ID** | vpc-rtznlhxm |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | tencent-ap-seoul |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my05-mig-subnet-01 | subnet-19h6xb1b | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my05-mig-sshkey-01 | skey-7ic7ibdp |  |  |

### Security Groups

#### Security Group: my05-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my05-mig-sg-01 |
| **CSP Security Group ID** | sg-d2jm3yvx |
| **VNet** | my05-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my05-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my05-mig-sg-02 |
| **CSP Security Group ID** | sg-hgf81iwn |
| **VNet** | my05-mig-vnet-01 |
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
| **Per Hour** | $0.5200 |
| **Per Day** | $12.48 |
| **Per Month (30 days)** | $374.40 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| TENCENT | ap-seoul | 3 | $0.5200 | $374.40 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | SA5.MEDIUM2 | $0.0600 | $43.20 |
| my05-ng-influxdb-back-1 | S6.LARGE16 | $0.2300 | $165.60 |
| my05-ng-influxdb-back-2 | S6.LARGE16 | $0.2300 | $165.60 |




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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my05-infra101)",
  "success": true
}
```

