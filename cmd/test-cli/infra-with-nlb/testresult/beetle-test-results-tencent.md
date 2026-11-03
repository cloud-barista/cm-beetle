# CM-Beetle test results for TENCENT (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with TENCENT cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.3+ (d67a8e3)
- imdl: v0.1.19+ (d67a8e3)
- CB-Tumblebug: v0.13.8
- CB-Spider: v0.13.20
- CB-MapUI: v0.13.14
- Target CSP: TENCENT
- Target Region: ap-seoul
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: November 3, 2026
- Test Time: 17:55:07 KST
- Test Execution: 2026-11-03 17:55:07 KST

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
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 7.248s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 378ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m31.791s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 14ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 3ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 18ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 2.227s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 19.295s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 8ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 11ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 1m23.391s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.191s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.208s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 1m17.739s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 34.766s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 6m17.5184278s

*Test executed on November 3, 2026 at 17:55:07 KST (2026-11-03 17:55:07 KST) using CM-Beetle automated test CLI*

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
            "specId": "tencent+ap-seoul+sa5.large16",
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
            "imageId": "img-7rotv4ux",
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
          "id": "tencent+ap-seoul+sa5.large16",
          "uid": "tbj7e4ekpmehjld13pi6",
          "cspSpecName": "SA5.LARGE16",
          "name": "tencent+ap-seoul+sa5.large16",
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
          "costPerHour": 0.22,
          "orderInFilteredResult": 7,
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
              "value": "SA5.LARGE16"
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
              "value": "{UnitPrice:0.22,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.044,UnitPriceSecondStep:0.22,UnitPriceDiscountSecondStep:0.044,UnitPriceThirdStep:0.22,UnitPriceDiscountThirdStep:0.044}"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "30"
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
          "availabilityStatus": "Available",
          "availableZones": [
            "ap-seoul-1",
            "ap-seoul-2"
          ],
          "availabilityCheckedAt": "2026-11-03T08:55:00.811872587Z",
          "availabilityRefreshAfter": "2026-11-03T09:55:00.811872587Z"
        },
        {
          "id": "tencent+ap-seoul+sa5.medium2",
          "uid": "tbk2u12n0igvrispt92u",
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
          "orderInFilteredResult": 13,
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
          "uid": "tbrtv38m36npv5tcj4qk",
          "name": "img-7rotv4ux",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "tencent-sa-saopaulo",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:25:24 Thu",
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
            "specId": "tencent+ap-seoul+s8.medium2",
            "imageId": "img-7rotv4ux",
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
          "uid": "tb722u163vtjo7lfom0v",
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
          "availabilityCheckedAt": "2026-11-03T08:55:00.817243487Z",
          "availabilityRefreshAfter": "2026-11-03T09:55:00.817243487Z"
        },
        {
          "id": "tencent+ap-seoul+s8.medium2",
          "uid": "tb7ks8ma414p48ovhhmg",
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
          "orderInFilteredResult": 22,
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
          "uid": "tbrtv38m36npv5tcj4qk",
          "name": "img-7rotv4ux",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "tencent-sa-saopaulo",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:25:24 Thu",
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
        "specId": "tencent+ap-seoul+sa5.large16",
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
        "imageId": "img-7rotv4ux",
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
      "id": "tencent+ap-seoul+sa5.large16",
      "uid": "tbj7e4ekpmehjld13pi6",
      "cspSpecName": "SA5.LARGE16",
      "name": "tencent+ap-seoul+sa5.large16",
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
      "costPerHour": 0.22,
      "orderInFilteredResult": 7,
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
          "value": "SA5.LARGE16"
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
          "value": "{UnitPrice:0.22,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.044,UnitPriceSecondStep:0.22,UnitPriceDiscountSecondStep:0.044,UnitPriceThirdStep:0.22,UnitPriceDiscountThirdStep:0.044}"
        },
        {
          "key": "InstanceBandwidth",
          "value": "1.5"
        },
        {
          "key": "InstancePps",
          "value": "30"
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
      "availabilityStatus": "Available",
      "availableZones": [
        "ap-seoul-1",
        "ap-seoul-2"
      ],
      "availabilityCheckedAt": "2026-11-03T08:55:00.811872587Z",
      "availabilityRefreshAfter": "2026-11-03T09:55:00.811872587Z"
    },
    {
      "id": "tencent+ap-seoul+sa5.medium2",
      "uid": "tbk2u12n0igvrispt92u",
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
      "orderInFilteredResult": 13,
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
      "uid": "tbrtv38m36npv5tcj4qk",
      "name": "img-7rotv4ux",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "tencent-sa-saopaulo",
      "infraType": "",
      "fetchedTime": "2026.10.29 04:25:24 Thu",
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
        "nodeGroupId": "my05-ng-influxdb-back"
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
  "uid": "tbf3tsbhtkhbgi88ndkt",
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
    "cm-migrated-at": "2026-11-03T08:55:25Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my05-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my05-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbf3tsbhtkhbgi88ndkt"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbmmh0ojen41pao6hsik",
      "cspResourceName": "tbmmh0ojen41pao6hsik",
      "cspResourceId": "ins-7h4dbejf",
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
      "createdTime": "2026-11-03 08:55:56",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "101.33.73.9",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.13",
      "privateDNS": "",
      "rootDiskType": "CLOUD_HSSD",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-evs80xkr",
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
      "cspVNetId": "vpc-6sjytz4a",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-1poy0xlj",
      "networkInterface": "eni-6e62gb9c",
      "securityGroupIds": [
        "my05-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-ofv7yi59",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDVQAhryPeAzpaXy9qRZ8FMdJLaYqmdCMS82MsqKspzxteX/YsklowDQYB4efjzxUBYqmmv6x5x1uLhng5VhEuo=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:0Q2hhwHMBEEK1txJsFCcKRVAYWYYGTQbCbnPVgBt4UE",
        "firstUsedAt": "2026-11-03T08:56:39Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-03T08:56:01Z",
          "completedTime": "2026-11-03T08:56:26Z",
          "elapsedTime": 25,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.13:22 (as \"cb-user\") via bastion Node \"my05-ng-influxdb-back-1\" at 43.133.92.116:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 43.133.92.116:22 as user \"cb-user\" (bastionNodeId=my05-ng-influxdb-back-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:39Z",
          "completedTime": "2026-11-03T08:56:40Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my05-infra101-tb6h71qjs7l2ln9q33p1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:40Z",
          "completedTime": "2026-11-03T08:56:41Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-13-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-ng-influxdb-back-1",
      "uid": "tbs6a4suvroi915pj8md",
      "cspResourceName": "tbs6a4suvroi915pj8md",
      "cspResourceId": "ins-ncurja61",
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
      "createdTime": "2026-11-03 08:55:54",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.133.92.116",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.12",
      "privateDNS": "",
      "rootDiskType": "CLOUD_HSSD",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-ad5w7al3",
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
      "cspVNetId": "vpc-6sjytz4a",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-1poy0xlj",
      "networkInterface": "eni-gcjn4nl0",
      "securityGroupIds": [
        "my05-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-ofv7yi59",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBF8TcwNVd2LZiKP4+ROrPOH7g38PqZjb3+b9ale9TjP55kWOjkVi8NzX7044uOH63ZmCrQFP851HZkYI04aUcpQ=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:R44brYbZ1TiTOv+ZgkxyFise8lFcn9jubPPAYNw4fl8",
        "firstUsedAt": "2026-11-03T08:56:17Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:01Z",
          "completedTime": "2026-11-03T08:56:29Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-11-03T08:56:40Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb6h71qjs7l2ln9q33p1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:40Z",
          "completedTime": "2026-11-03T08:56:42Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-12-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-ng-influxdb-back-2",
      "uid": "tbu8ptqpubtv73mdejs5",
      "cspResourceName": "tbu8ptqpubtv73mdejs5",
      "cspResourceId": "ins-2uveewfz",
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
      "createdTime": "2026-11-03 08:56:01",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.166.14.150",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.2",
      "privateDNS": "",
      "rootDiskType": "CLOUD_HSSD",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-06zi2y8n",
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
      "cspVNetId": "vpc-6sjytz4a",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-1poy0xlj",
      "networkInterface": "eni-6laeqnew",
      "securityGroupIds": [
        "my05-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-ofv7yi59",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBE4WRcAiAu1JCFwmTart1ssAQ4GGfM76L1zrGqV39995D/Fr9gkfepS6UFt12OXZN/WbhZmodAB7GgYZCZOJArQ=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:Dcg+YfpF8AZHid5AxpCI3C0ejHbDPPHP3wJy+jkZPyE",
        "firstUsedAt": "2026-11-03T08:56:39Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-03T08:56:01Z",
          "completedTime": "2026-11-03T08:56:26Z",
          "elapsedTime": 25,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my05-ng-influxdb-back-2\" at 10.0.1.2:22 (as \"cb-user\") via bastion Node \"my05-ng-influxdb-back-1\" at 43.133.92.116:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 43.133.92.116:22 as user \"cb-user\" (bastionNodeId=my05-ng-influxdb-back-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:39Z",
          "completedTime": "2026-11-03T08:56:40Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my05-infra101-tb6h71qjs7l2ln9q33p1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:40Z",
          "completedTime": "2026-11-03T08:56:41Z",
          "elapsedTime": 1,
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
            "nodeId": "my05-ng-influxdb-back-2",
            "nodeIp": "43.166.14.150",
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
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "101.33.73.9",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-13-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-influxdb-back-1",
            "nodeIp": "43.133.92.116",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-12-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my05-infra101-tb6h71qjs7l2ln9q33p1"
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
- **Count**: 1

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "infra": [
    {
      "resourceType": "infra",
      "id": "my05-infra101",
      "uid": "tbf3tsbhtkhbgi88ndkt",
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
        "cm-migrated-at": "2026-11-03T08:55:25Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my05-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbf3tsbhtkhbgi88ndkt"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbmmh0ojen41pao6hsik",
          "cspResourceName": "tbmmh0ojen41pao6hsik",
          "cspResourceId": "ins-7h4dbejf",
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
          "createdTime": "2026-11-03 08:55:56",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-11-03 08:55:56",
            "sys.cspResourceId": "ins-7h4dbejf",
            "sys.cspResourceName": "tbmmh0ojen41pao6hsik",
            "sys.id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tbmmh0ojen41pao6hsik",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "101.33.73.9",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.13",
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
          "cspVNetId": "vpc-6sjytz4a",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-1poy0xlj",
          "networkInterface": "eni-6e62gb9c",
          "securityGroupIds": [
            "my05-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-ofv7yi59"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-1",
          "uid": "tbs6a4suvroi915pj8md",
          "cspResourceName": "tbs6a4suvroi915pj8md",
          "cspResourceId": "ins-ncurja61",
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
          "createdTime": "2026-11-03 08:55:54",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-11-03 08:55:54",
            "sys.cspResourceId": "ins-ncurja61",
            "sys.cspResourceName": "tbs6a4suvroi915pj8md",
            "sys.id": "my05-ng-influxdb-back-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tbs6a4suvroi915pj8md",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.133.92.116",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.12",
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
          "cspVNetId": "vpc-6sjytz4a",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-1poy0xlj",
          "networkInterface": "eni-gcjn4nl0",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-ofv7yi59"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-2",
          "uid": "tbu8ptqpubtv73mdejs5",
          "cspResourceName": "tbu8ptqpubtv73mdejs5",
          "cspResourceId": "ins-2uveewfz",
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
          "createdTime": "2026-11-03 08:56:01",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-11-03 08:56:01",
            "sys.cspResourceId": "ins-2uveewfz",
            "sys.cspResourceName": "tbu8ptqpubtv73mdejs5",
            "sys.id": "my05-ng-influxdb-back-2",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tbu8ptqpubtv73mdejs5",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.166.14.150",
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
          "cspVNetId": "vpc-6sjytz4a",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-1poy0xlj",
          "networkInterface": "eni-6laeqnew",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-ofv7yi59"
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
- **IDs**: [my05-infra101]

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
  "uid": "tbf3tsbhtkhbgi88ndkt",
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
    "cm-migrated-at": "2026-11-03T08:55:25Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my05-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my05-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbf3tsbhtkhbgi88ndkt"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbmmh0ojen41pao6hsik",
      "cspResourceName": "tbmmh0ojen41pao6hsik",
      "cspResourceId": "ins-7h4dbejf",
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
      "createdTime": "2026-11-03 08:55:56",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "101.33.73.9",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.13",
      "privateDNS": "",
      "rootDiskType": "CLOUD_HSSD",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-evs80xkr",
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
      "cspVNetId": "vpc-6sjytz4a",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-1poy0xlj",
      "networkInterface": "eni-6e62gb9c",
      "securityGroupIds": [
        "my05-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-ofv7yi59",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDVQAhryPeAzpaXy9qRZ8FMdJLaYqmdCMS82MsqKspzxteX/YsklowDQYB4efjzxUBYqmmv6x5x1uLhng5VhEuo=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:0Q2hhwHMBEEK1txJsFCcKRVAYWYYGTQbCbnPVgBt4UE",
        "firstUsedAt": "2026-11-03T08:56:39Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-03T08:56:01Z",
          "completedTime": "2026-11-03T08:56:26Z",
          "elapsedTime": 25,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.13:22 (as \"cb-user\") via bastion Node \"my05-ng-influxdb-back-1\" at 43.133.92.116:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 43.133.92.116:22 as user \"cb-user\" (bastionNodeId=my05-ng-influxdb-back-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:39Z",
          "completedTime": "2026-11-03T08:56:40Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my05-infra101-tb6h71qjs7l2ln9q33p1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:40Z",
          "completedTime": "2026-11-03T08:56:41Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-13-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-ng-influxdb-back-1",
      "uid": "tbs6a4suvroi915pj8md",
      "cspResourceName": "tbs6a4suvroi915pj8md",
      "cspResourceId": "ins-ncurja61",
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
      "createdTime": "2026-11-03 08:55:54",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.133.92.116",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.12",
      "privateDNS": "",
      "rootDiskType": "CLOUD_HSSD",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-ad5w7al3",
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
      "cspVNetId": "vpc-6sjytz4a",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-1poy0xlj",
      "networkInterface": "eni-gcjn4nl0",
      "securityGroupIds": [
        "my05-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-ofv7yi59",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBF8TcwNVd2LZiKP4+ROrPOH7g38PqZjb3+b9ale9TjP55kWOjkVi8NzX7044uOH63ZmCrQFP851HZkYI04aUcpQ=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:R44brYbZ1TiTOv+ZgkxyFise8lFcn9jubPPAYNw4fl8",
        "firstUsedAt": "2026-11-03T08:56:17Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:01Z",
          "completedTime": "2026-11-03T08:56:29Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-11-03T08:56:40Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb6h71qjs7l2ln9q33p1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:40Z",
          "completedTime": "2026-11-03T08:56:42Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-12-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-ng-influxdb-back-2",
      "uid": "tbu8ptqpubtv73mdejs5",
      "cspResourceName": "tbu8ptqpubtv73mdejs5",
      "cspResourceId": "ins-2uveewfz",
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
      "createdTime": "2026-11-03 08:56:01",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.166.14.150",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.2",
      "privateDNS": "",
      "rootDiskType": "CLOUD_HSSD",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-06zi2y8n",
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
      "cspVNetId": "vpc-6sjytz4a",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-1poy0xlj",
      "networkInterface": "eni-6laeqnew",
      "securityGroupIds": [
        "my05-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-ofv7yi59",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBE4WRcAiAu1JCFwmTart1ssAQ4GGfM76L1zrGqV39995D/Fr9gkfepS6UFt12OXZN/WbhZmodAB7GgYZCZOJArQ=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:Dcg+YfpF8AZHid5AxpCI3C0ejHbDPPHP3wJy+jkZPyE",
        "firstUsedAt": "2026-11-03T08:56:39Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-03T08:56:01Z",
          "completedTime": "2026-11-03T08:56:26Z",
          "elapsedTime": 25,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my05-ng-influxdb-back-2\" at 10.0.1.2:22 (as \"cb-user\") via bastion Node \"my05-ng-influxdb-back-1\" at 43.133.92.116:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 43.133.92.116:22 as user \"cb-user\" (bastionNodeId=my05-ng-influxdb-back-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:39Z",
          "completedTime": "2026-11-03T08:56:40Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my05-infra101-tb6h71qjs7l2ln9q33p1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T08:56:40Z",
          "completedTime": "2026-11-03T08:56:41Z",
          "elapsedTime": 1,
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
            "nodeId": "my05-ng-influxdb-back-2",
            "nodeIp": "43.166.14.150",
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
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "101.33.73.9",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-13-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-influxdb-back-1",
            "nodeIp": "43.133.92.116",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-12-ubuntu 5.15.0-191-generic #201-Ubuntu SMP Fri Aug 7 18:39:04 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my05-infra101-tb6h71qjs7l2ln9q33p1"
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
      "id": "my05-ng-influxdb-back",
      "cspResourceName": "tbv73dv4ckbrlo6pn8t2",
      "cspResourceId": "lb-o8lkwd2z",
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
        "dnsName": "lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com",
        "keyValueList": [
          {
            "key": "LoadBalancerId",
            "value": "lb-o8lkwd2z"
          },
          {
            "key": "LoadBalancerName",
            "value": "tbv73dv4ckbrlo6pn8t2"
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
            "value": "lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com"
          },
          {
            "key": "Status",
            "value": "1"
          },
          {
            "key": "CreateTime",
            "value": "2026-11-03 16:57:23"
          },
          {
            "key": "StatusTime",
            "value": "2026-11-03 16:57:31"
          },
          {
            "key": "ProjectId",
            "value": "0"
          },
          {
            "key": "VpcId",
            "value": "vpc-6sjytz4a"
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
            "value": "{Region:ap-seoul,VpcId:vpc-6sjytz4a}"
          },
          {
            "key": "AddressIPVersion",
            "value": "ipv4"
          },
          {
            "key": "NumericalVpcId",
            "value": "15841761"
          },
          {
            "key": "VipIsp",
            "value": "BGP"
          },
          {
            "key": "MasterZone",
            "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
          },
          {
            "key": "BackupZoneSet",
            "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
          },
          {
            "key": "ExpireTime",
            "value": "2026-12-03 16:57:23"
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
      "createdTime": "2026-11-03T16:57:23Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "LoadBalancerId",
          "value": "lb-o8lkwd2z"
        },
        {
          "key": "LoadBalancerName",
          "value": "tbv73dv4ckbrlo6pn8t2"
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
          "value": "lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com"
        },
        {
          "key": "Status",
          "value": "1"
        },
        {
          "key": "CreateTime",
          "value": "2026-11-03 16:57:23"
        },
        {
          "key": "StatusTime",
          "value": "2026-11-03 16:57:31"
        },
        {
          "key": "ProjectId",
          "value": "0"
        },
        {
          "key": "VpcId",
          "value": "vpc-6sjytz4a"
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
          "value": "{Region:ap-seoul,VpcId:vpc-6sjytz4a}"
        },
        {
          "key": "AddressIPVersion",
          "value": "ipv4"
        },
        {
          "key": "NumericalVpcId",
          "value": "15841761"
        },
        {
          "key": "VipIsp",
          "value": "BGP"
        },
        {
          "key": "MasterZone",
          "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
        },
        {
          "key": "BackupZoneSet",
          "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
        },
        {
          "key": "ExpireTime",
          "value": "2026-12-03 16:57:23"
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
    "cspResourceName": "tbv73dv4ckbrlo6pn8t2",
    "cspResourceId": "lb-o8lkwd2z",
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
      "dnsName": "lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com",
      "keyValueList": [
        {
          "key": "LoadBalancerId",
          "value": "lb-o8lkwd2z"
        },
        {
          "key": "LoadBalancerName",
          "value": "tbv73dv4ckbrlo6pn8t2"
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
          "value": "lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com"
        },
        {
          "key": "Status",
          "value": "1"
        },
        {
          "key": "CreateTime",
          "value": "2026-11-03 16:57:23"
        },
        {
          "key": "StatusTime",
          "value": "2026-11-03 16:57:31"
        },
        {
          "key": "ProjectId",
          "value": "0"
        },
        {
          "key": "VpcId",
          "value": "vpc-6sjytz4a"
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
          "value": "{Region:ap-seoul,VpcId:vpc-6sjytz4a}"
        },
        {
          "key": "AddressIPVersion",
          "value": "ipv4"
        },
        {
          "key": "NumericalVpcId",
          "value": "15841761"
        },
        {
          "key": "VipIsp",
          "value": "BGP"
        },
        {
          "key": "MasterZone",
          "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
        },
        {
          "key": "BackupZoneSet",
          "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
        },
        {
          "key": "ExpireTime",
          "value": "2026-12-03 16:57:23"
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
    "createdTime": "2026-11-03T16:57:23Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "LoadBalancerId",
        "value": "lb-o8lkwd2z"
      },
      {
        "key": "LoadBalancerName",
        "value": "tbv73dv4ckbrlo6pn8t2"
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
        "value": "lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com"
      },
      {
        "key": "Status",
        "value": "1"
      },
      {
        "key": "CreateTime",
        "value": "2026-11-03 16:57:23"
      },
      {
        "key": "StatusTime",
        "value": "2026-11-03 16:57:31"
      },
      {
        "key": "ProjectId",
        "value": "0"
      },
      {
        "key": "VpcId",
        "value": "vpc-6sjytz4a"
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
        "value": "{Region:ap-seoul,VpcId:vpc-6sjytz4a}"
      },
      {
        "key": "AddressIPVersion",
        "value": "ipv4"
      },
      {
        "key": "NumericalVpcId",
        "value": "15841761"
      },
      {
        "key": "VipIsp",
        "value": "BGP"
      },
      {
        "key": "MasterZone",
        "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
      },
      {
        "key": "BackupZoneSet",
        "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
      },
      {
        "key": "ExpireTime",
        "value": "2026-12-03 16:57:23"
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
  "cspResourceName": "tbv73dv4ckbrlo6pn8t2",
  "cspResourceId": "lb-o8lkwd2z",
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
    "dnsName": "lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com",
    "keyValueList": [
      {
        "key": "LoadBalancerId",
        "value": "lb-o8lkwd2z"
      },
      {
        "key": "LoadBalancerName",
        "value": "tbv73dv4ckbrlo6pn8t2"
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
        "value": "lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com"
      },
      {
        "key": "Status",
        "value": "1"
      },
      {
        "key": "CreateTime",
        "value": "2026-11-03 16:57:23"
      },
      {
        "key": "StatusTime",
        "value": "2026-11-03 16:57:31"
      },
      {
        "key": "ProjectId",
        "value": "0"
      },
      {
        "key": "VpcId",
        "value": "vpc-6sjytz4a"
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
        "value": "{Region:ap-seoul,VpcId:vpc-6sjytz4a}"
      },
      {
        "key": "AddressIPVersion",
        "value": "ipv4"
      },
      {
        "key": "NumericalVpcId",
        "value": "15841761"
      },
      {
        "key": "VipIsp",
        "value": "BGP"
      },
      {
        "key": "MasterZone",
        "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
      },
      {
        "key": "BackupZoneSet",
        "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
      },
      {
        "key": "ExpireTime",
        "value": "2026-12-03 16:57:23"
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
  "createdTime": "2026-11-03T16:57:23Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "LoadBalancerId",
      "value": "lb-o8lkwd2z"
    },
    {
      "key": "LoadBalancerName",
      "value": "tbv73dv4ckbrlo6pn8t2"
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
      "value": "lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com"
    },
    {
      "key": "Status",
      "value": "1"
    },
    {
      "key": "CreateTime",
      "value": "2026-11-03 16:57:23"
    },
    {
      "key": "StatusTime",
      "value": "2026-11-03 16:57:31"
    },
    {
      "key": "ProjectId",
      "value": "0"
    },
    {
      "key": "VpcId",
      "value": "vpc-6sjytz4a"
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
      "value": "{Region:ap-seoul,VpcId:vpc-6sjytz4a}"
    },
    {
      "key": "AddressIPVersion",
      "value": "ipv4"
    },
    {
      "key": "NumericalVpcId",
      "value": "15841761"
    },
    {
      "key": "VipIsp",
      "value": "BGP"
    },
    {
      "key": "MasterZone",
      "value": "{ZoneId:180001,Zone:ap-seoul-1,ZoneName:ap-seoul-1,ZoneRegion:ap-seoul,LocalZone:false}"
    },
    {
      "key": "BackupZoneSet",
      "value": "{ZoneId:180002,Zone:ap-seoul-2,ZoneName:ap-seoul-2,ZoneRegion:ap-seoul,LocalZone:false}"
    },
    {
      "key": "ExpireTime",
      "value": "2026-12-03 16:57:23"
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

**Summary**: Load balancing confirmed: traffic distributed across 2 unique VMs

### Test Case 12: Target Infrastructure Summary

#### 12.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Target Infrastructure Summary**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-03 08:59:13

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
| **Total NLBs** | 1 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| SA5.LARGE16 | 4 | 16.0 | - | x86_64 |  | $0.2200 | 2 |
| SA5.MEDIUM2 | 2 | 2.0 | - | x86_64 |  | $0.0600 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| img-7rotv4ux | Ubuntu Server 22.04 LTS 64bit UEFI | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | ins-7h4dbejf | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 101.33.73.9<br>**Private IP:** 10.0.1.13<br>**SGs:** my05-mig-sg-02<br>**SSH:** my05-mig-sshkey-01 |
| my05-ng-influxdb-back-1 | ins-ncurja61 | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.133.92.116<br>**Private IP:** 10.0.1.12<br>**SGs:** my05-mig-sg-01<br>**SSH:** my05-mig-sshkey-01 |
| my05-ng-influxdb-back-2 | ins-2uveewfz | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.166.14.150<br>**Private IP:** 10.0.1.2<br>**SGs:** my05-mig-sg-01<br>**SSH:** my05-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my05-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my05-mig-vnet-01 |
| **CSP VNet ID** | vpc-6sjytz4a |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | tencent-ap-seoul |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my05-mig-subnet-01 | subnet-1poy0xlj | 10.0.1.0/24 |  |

### Network Load Balancers (NLB)

#### NLB: my05-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my05-ng-influxdb-back |
| **CSP NLB Name** | tbv73dv4ckbrlo6pn8t2 |
| **CSP NLB ID** | lb-o8lkwd2z |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | tencent-ap-seoul |
| **Listener** | TCP:9999 (DNS: lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com) |
| **Target Group** | TCP:8086 (Node Group: my05-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my05-ng-influxdb-back-1 | ins-ncurja61 | 10.0.1.12 | 43.133.92.116 | Running |
| my05-ng-influxdb-back-2 | ins-2uveewfz | 10.0.1.2 | 43.166.14.150 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my05-mig-sshkey-01 | skey-ofv7yi59 |  |  |

### Security Groups

#### Security Group: my05-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my05-mig-sg-01 |
| **CSP Security Group ID** | sg-b31fuorh |
| **VNet** | my05-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my05-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my05-mig-sg-02 |
| **CSP Security Group ID** | sg-dmpck3qp |
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
| **Per Hour** | $0.5000 |
| **Per Day** | $12.00 |
| **Per Month (30 days)** | $360.00 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| TENCENT | ap-seoul | 3 | $0.5000 | $360.00 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | SA5.MEDIUM2 | $0.0600 | $43.20 |
| my05-ng-influxdb-back-1 | SA5.LARGE16 | $0.2200 | $158.40 |
| my05-ng-influxdb-back-2 | SA5.LARGE16 | $0.2200 | $158.40 |




### Test Case 13: Migration Report

#### 13.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 13.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-11-03 08:59:18*

---

## 📊 Migration Summary

**Target Cloud:** TENCENT

**Target Region:** ap-seoul

**Namespace:** mig01 | **Infra ID:** my05-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**NLB Migration Status:** Completed

**Total Load Balancers:** 1

**Migrated Load Balancers:** 1

**💰 Estimated Monthly Cost:** $360.00 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | SA5.MEDIUM2, SA5.LARGE16 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my05-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my05-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 10 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |
| 8 | **Network Load Balancer** | 1 | ✅ Created | my05-ng-influxdb-back (TCP:9999 -> TCP:8086) |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** ins-7h4dbejf<br>**Label(sourceMachineId):** ec268ed7-821e-9d73-e79f-961262161624 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 |
| 2 | **VM Name:** my05-ng-influxdb-back-1<br>**VM ID:** ins-ncurja61<br>**Label(sourceMachineId):** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 |
| 3 | **VM Name:** my05-ng-influxdb-back-2<br>**VM ID:** ins-2uveewfz<br>**Label(sourceMachineId):** ec288dd0-c6fa-8a49-2f60-bc898311febf | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** SA5.MEDIUM2<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 8 GB |
| 2 | my05-ng-influxdb-back-1 | **Spec ID:** SA5.LARGE16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 30 GB |
| 3 | my05-ng-influxdb-back-2 | **Spec ID:** SA5.LARGE16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 8 GB<br>**Root Disk:** 30 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** img-7rotv4ux<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Server 22.04 LTS 64bit UEFI | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | my05-ng-influxdb-back-1 | **Image ID:** img-7rotv4ux<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Server 22.04 LTS 64bit UEFI | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | my05-ng-influxdb-back-2 | **Image ID:** img-7rotv4ux<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Server 22.04 LTS 64bit UEFI | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 10 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my05-mig-sg-01

**CSP ID:** sg-b31fuorh | **VNet:** my05-mig-vnet-01 | **Rules:** 6

**Assigned VMs:**

- **VM:** my05-ng-influxdb-back-1
  - **Source Server:** **Hostname:** ip-10-0-1-221, **Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932
- **VM:** my05-ng-influxdb-back-2
  - **Source Server:** **Hostname:** ip-10-0-1-138, **Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 2 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 4 | inbound | TCP | 9999 | 0.0.0.0/0 | - | Created by system |
| 5 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 6 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my05-mig-sg-02

**CSP ID:** sg-dmpck3qp | **VNet:** my05-mig-vnet-01 | **Rules:** 4

**Assigned VMs:**

- **VM:** my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1
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
| 1 | **Name:** my05-mig-vnet-01<br>**ID:** vpc-6sjytz4a | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my05-mig-subnet-01<br>**ID:** subnet-1poy0xlj | 10.0.1.0/24 | my05-mig-vnet-01 |

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

### 1. influxdb_back (haproxy) → my05-ng-influxdb-back

**Configuration Comparison:**

| Property | Source (On-Premise) | Target (Cloud NLB) | Note |
|----------|---------------------|--------------------|------|
| **Load Balancer / Name** | haproxy (Host: ip-10-0-1-30) | my05-ng-influxdb-back (CSP ID: tbv73dv4ckbrlo6pn8t2) | Replaced by managed NLB |
| **Type / Scope** | Software LB | PUBLIC / REGION | Cloud-native NLB |
| **Frontend Listener** | *:9999 (TCP) | TCP:9999 | DNS: lb-o8lkwd2z-9xmsd0r2j3vfnl5i.clb.kr-tencentclb.com |
| **Backend / Target Group** | influxdb_back (TCP:8086, roundrobin) | my05-ng-influxdb-back (TCP:8086, 2 nodes) | Target Group attached to NodeGroup |
| **Health Check** | Enabled (Interval: 10s, Timeout: 10s, Threshold: 3) | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) | Preserved |

**Backend Node Mappings:**

| No. | Source Backend Server | Source Endpoint | Target Node | Target CSP VM ID | Target Private IP | Target Public IP | Target Port | Status |
|-----|-----------------------|-----------------|-------------|------------------|-------------------|------------------|-------------|--------|
| 1 | influx1 (ip-10-0-1-221) | 10.0.1.221:8086 | my05-ng-influxdb-back-1 | ins-ncurja61 | 10.0.1.12 | 43.133.92.116 | 8086 | Running |
| 2 | influx2 (ip-10-0-1-138) | 10.0.1.138:8086 | my05-ng-influxdb-back-2 | ins-2uveewfz | 10.0.1.2 | 43.166.14.150 | 8086 | Running |

**Migration Notes:**
- On-premise haproxy on ip-10-0-1-30 (migrated to my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1) is replaced by cloud-managed NLB.

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my05-mig-sshkey-01 | skey-ofv7yi59 |  | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.5000 |
| Daily | $12.00 |
| Monthly | $360.00 |
| Yearly | $4320.00 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | SA5.MEDIUM2 | $43.20 | 12.0% |
| ip-10-0-1-221 (migrated) | SA5.LARGE16 | $158.40 | 44.0% |
| ip-10-0-1-138 (migrated) | SA5.LARGE16 | $158.40 | 44.0% |

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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my05-infra101)",
  "success": true
}
```

