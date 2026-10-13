# CM-Beetle test results for TENCENT (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with TENCENT cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.1+ (3649ed7)
- imdl: v0.1.16+ (3649ed7)
- CB-Tumblebug: unknown
- CB-Spider: unknown
- CB-MapUI: v0.13.13
- Target CSP: TENCENT
- Target Region: ap-seoul
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: October 12, 2026
- Test Time: 21:19:08 KST
- Test Execution: 2026-10-12 21:19:08 KST

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
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 2.439s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 1.698s | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m18.757s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 12ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 11ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 20ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 1.599s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 19.511s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 5ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 6ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 1m23.447s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.351s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.409s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 1m19.257s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 36.802s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 6m4.571922511s

*Test executed on October 12, 2026 at 21:19:08 KST (2026-10-12 21:19:08 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=60.0% Max=100.0% Avg=86.7% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
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
          ]
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
          ]
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=60.0% Max=100.0% Avg=86.7% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
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
          ]
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
          ]
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
  "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=60.0% Max=100.0% Avg=86.7% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
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
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
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
      ]
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
      ]
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
  "uid": "tbi9qv0015b5fmf6vtoe",
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
    "cm-migrated-at": "2026-10-12T12:19:37Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my05-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my05-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbi9qv0015b5fmf6vtoe"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbk8gdrar02p4pigcka0",
      "cspResourceName": "tbk8gdrar02p4pigcka0",
      "cspResourceId": "ins-rn0jiohn",
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
      "createdTime": "2026-10-12 12:20:08",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.155.158.249",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.9",
      "privateDNS": "",
      "rootDiskType": "CLOUD_HSSD",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-8mjrvptj",
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
      "cspVNetId": "vpc-38jxtmcs",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-chmtrzmn",
      "networkInterface": "eni-dn4z9ep8",
      "securityGroupIds": [
        "my05-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-jab0kioj",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGZjPhkh17/50PlGCufKUubnq8CnFcpqNS8FkU8AUMf7X1AS2Jz5yVufIAuxbDMBrrcIb3JXXH5S4Cnw0HW2xbI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:9tMCcTUkCOUr07bqTKn9R46Mn5HyBPEyXEhfSEQBGPM",
        "firstUsedAt": "2026-10-12T12:20:26Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-10-12T12:20:08Z",
          "completedTime": "2026-10-12T12:20:26Z",
          "elapsedTime": 18,
          "resultSummary": "Command execution failed",
          "errorMessage": "SSH authentication failed. Please check: 1) private key is valid 2) user 'cb-user' exists on target 3) authorized_keys is properly configured"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:38Z",
          "completedTime": "2026-10-12T12:20:39Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my05-infra101-tb1j1uhhsjhem26nd0lo",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:40Z",
          "completedTime": "2026-10-12T12:20:40Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-9-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-ng-influxdb-back-1",
      "uid": "tb92iq0gift5smd0fq95",
      "cspResourceName": "tb92iq0gift5smd0fq95",
      "cspResourceId": "ins-fgook4wl",
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
      "createdTime": "2026-10-12 12:20:04",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.133.249.165",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-nb4pjhtr",
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
      "cspVNetId": "vpc-38jxtmcs",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-chmtrzmn",
      "networkInterface": "eni-nlimpves",
      "securityGroupIds": [
        "my05-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-jab0kioj",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMVUtYjl0VBYfKOi7NRrxLVkdxR592nVGT381GGM/ZPxhTAWtnpQIzdjm/3VDbcMpH7bgiBlKBtsKWst/95OrGI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:9nk7CLx4+0yPblsRzmqpx2P444ojwQDuC7A4+CoWsTM",
        "firstUsedAt": "2026-10-12T12:20:26Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:08Z",
          "completedTime": "2026-10-12T12:20:27Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-10-12T12:20:39Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb1j1uhhsjhem26nd0lo",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:40Z",
          "completedTime": "2026-10-12T12:20:40Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-4-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-ng-influxdb-back-2",
      "uid": "tbslr9p77kpm3kll6bfn",
      "cspResourceName": "tbslr9p77kpm3kll6bfn",
      "cspResourceId": "ins-flev5l93",
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
      "createdTime": "2026-10-12 12:20:03",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "119.28.161.163",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.12",
      "privateDNS": "",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-l3zo8ymz",
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
      "cspVNetId": "vpc-38jxtmcs",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-chmtrzmn",
      "networkInterface": "eni-8o9mmz2e",
      "securityGroupIds": [
        "my05-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-jab0kioj",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJXmyUlX8xCKniPmzLkQHGyw7v+4p24oB72Y61YMT5hu1B16i38XL1MWnS/D56Oz2Zc0C19ZWozWLDx9JZJpysM=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:64GyW9qw8ev2qpOzJiz4qkm3TpYSQI0rbA4ab2HU18A",
        "firstUsedAt": "2026-10-12T12:20:24Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:08Z",
          "completedTime": "2026-10-12T12:20:28Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-10-12T12:20:40Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb1j1uhhsjhem26nd0lo",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:40Z",
          "completedTime": "2026-10-12T12:20:41Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-12-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
            "nodeId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "43.155.158.249",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-9-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-influxdb-back-1",
            "nodeIp": "43.133.249.165",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-4-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-influxdb-back-2",
            "nodeIp": "119.28.161.163",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-12-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my05-infra101-tb1j1uhhsjhem26nd0lo"
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
      "uid": "tbi9qv0015b5fmf6vtoe",
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
        "cm-migrated-at": "2026-10-12T12:19:37Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my05-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbi9qv0015b5fmf6vtoe"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbk8gdrar02p4pigcka0",
          "cspResourceName": "tbk8gdrar02p4pigcka0",
          "cspResourceId": "ins-rn0jiohn",
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
          "createdTime": "2026-10-12 12:20:08",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-10-12 12:20:08",
            "sys.cspResourceId": "ins-rn0jiohn",
            "sys.cspResourceName": "tbk8gdrar02p4pigcka0",
            "sys.id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tbk8gdrar02p4pigcka0",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.155.158.249",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.9",
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
          "cspVNetId": "vpc-38jxtmcs",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-chmtrzmn",
          "networkInterface": "eni-dn4z9ep8",
          "securityGroupIds": [
            "my05-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-jab0kioj"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-1",
          "uid": "tb92iq0gift5smd0fq95",
          "cspResourceName": "tb92iq0gift5smd0fq95",
          "cspResourceId": "ins-fgook4wl",
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
          "createdTime": "2026-10-12 12:20:04",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-10-12 12:20:04",
            "sys.cspResourceId": "ins-fgook4wl",
            "sys.cspResourceName": "tb92iq0gift5smd0fq95",
            "sys.id": "my05-ng-influxdb-back-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tb92iq0gift5smd0fq95",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.133.249.165",
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
          "cspVNetId": "vpc-38jxtmcs",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-chmtrzmn",
          "networkInterface": "eni-nlimpves",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-jab0kioj"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-2",
          "uid": "tbslr9p77kpm3kll6bfn",
          "cspResourceName": "tbslr9p77kpm3kll6bfn",
          "cspResourceId": "ins-flev5l93",
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
          "createdTime": "2026-10-12 12:20:03",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-10-12 12:20:03",
            "sys.cspResourceId": "ins-flev5l93",
            "sys.cspResourceName": "tbslr9p77kpm3kll6bfn",
            "sys.id": "my05-ng-influxdb-back-2",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tbslr9p77kpm3kll6bfn",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "119.28.161.163",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.12",
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
          "cspVNetId": "vpc-38jxtmcs",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-chmtrzmn",
          "networkInterface": "eni-8o9mmz2e",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-jab0kioj"
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
  "uid": "tbi9qv0015b5fmf6vtoe",
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
    "cm-migrated-at": "2026-10-12T12:19:37Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my05-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my05-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbi9qv0015b5fmf6vtoe"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbk8gdrar02p4pigcka0",
      "cspResourceName": "tbk8gdrar02p4pigcka0",
      "cspResourceId": "ins-rn0jiohn",
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
      "createdTime": "2026-10-12 12:20:08",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.155.158.249",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.9",
      "privateDNS": "",
      "rootDiskType": "CLOUD_HSSD",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-8mjrvptj",
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
      "cspVNetId": "vpc-38jxtmcs",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-chmtrzmn",
      "networkInterface": "eni-dn4z9ep8",
      "securityGroupIds": [
        "my05-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-jab0kioj",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGZjPhkh17/50PlGCufKUubnq8CnFcpqNS8FkU8AUMf7X1AS2Jz5yVufIAuxbDMBrrcIb3JXXH5S4Cnw0HW2xbI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:9tMCcTUkCOUr07bqTKn9R46Mn5HyBPEyXEhfSEQBGPM",
        "firstUsedAt": "2026-10-12T12:20:26Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-10-12T12:20:08Z",
          "completedTime": "2026-10-12T12:20:26Z",
          "elapsedTime": 18,
          "resultSummary": "Command execution failed",
          "errorMessage": "SSH authentication failed. Please check: 1) private key is valid 2) user 'cb-user' exists on target 3) authorized_keys is properly configured"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:38Z",
          "completedTime": "2026-10-12T12:20:39Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my05-infra101-tb1j1uhhsjhem26nd0lo",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:40Z",
          "completedTime": "2026-10-12T12:20:40Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-9-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-ng-influxdb-back-1",
      "uid": "tb92iq0gift5smd0fq95",
      "cspResourceName": "tb92iq0gift5smd0fq95",
      "cspResourceId": "ins-fgook4wl",
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
      "createdTime": "2026-10-12 12:20:04",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.133.249.165",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-nb4pjhtr",
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
      "cspVNetId": "vpc-38jxtmcs",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-chmtrzmn",
      "networkInterface": "eni-nlimpves",
      "securityGroupIds": [
        "my05-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-jab0kioj",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMVUtYjl0VBYfKOi7NRrxLVkdxR592nVGT381GGM/ZPxhTAWtnpQIzdjm/3VDbcMpH7bgiBlKBtsKWst/95OrGI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:9nk7CLx4+0yPblsRzmqpx2P444ojwQDuC7A4+CoWsTM",
        "firstUsedAt": "2026-10-12T12:20:26Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:08Z",
          "completedTime": "2026-10-12T12:20:27Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-10-12T12:20:39Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb1j1uhhsjhem26nd0lo",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:40Z",
          "completedTime": "2026-10-12T12:20:40Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-4-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-ng-influxdb-back-2",
      "uid": "tbslr9p77kpm3kll6bfn",
      "cspResourceName": "tbslr9p77kpm3kll6bfn",
      "cspResourceId": "ins-flev5l93",
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
      "createdTime": "2026-10-12 12:20:03",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "119.28.161.163",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.12",
      "privateDNS": "",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 50,
      "rootDeviceName": "disk-l3zo8ymz",
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
      "cspVNetId": "vpc-38jxtmcs",
      "subnetId": "my05-mig-subnet-01",
      "cspSubnetId": "subnet-chmtrzmn",
      "networkInterface": "eni-8o9mmz2e",
      "securityGroupIds": [
        "my05-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-mig-sshkey-01",
      "cspSshKeyId": "skey-jab0kioj",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJXmyUlX8xCKniPmzLkQHGyw7v+4p24oB72Y61YMT5hu1B16i38XL1MWnS/D56Oz2Zc0C19ZWozWLDx9JZJpysM=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:64GyW9qw8ev2qpOzJiz4qkm3TpYSQI0rbA4ab2HU18A",
        "firstUsedAt": "2026-10-12T12:20:24Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:08Z",
          "completedTime": "2026-10-12T12:20:28Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-10-12T12:20:40Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb1j1uhhsjhem26nd0lo",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-12T12:20:40Z",
          "completedTime": "2026-10-12T12:20:41Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-12-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
            "nodeId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "43.155.158.249",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-9-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-influxdb-back-1",
            "nodeIp": "43.133.249.165",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-4-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-ng-influxdb-back-2",
            "nodeIp": "119.28.161.163",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-12-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my05-infra101-tb1j1uhhsjhem26nd0lo"
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
      "cspResourceName": "tbulrqdjeuctf5pl65g1",
      "cspResourceId": "lb-3zeo4wdb",
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
        "dnsName": "lb-3zeo4wdb-6401ixupewkjcl89.clb.kr-tencentclb.com",
        "keyValueList": [
          {
            "key": "LoadBalancerId",
            "value": "lb-3zeo4wdb"
          },
          {
            "key": "LoadBalancerName",
            "value": "tbulrqdjeuctf5pl65g1"
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
            "value": "lb-3zeo4wdb-6401ixupewkjcl89.clb.kr-tencentclb.com"
          },
          {
            "key": "Status",
            "value": "1"
          },
          {
            "key": "CreateTime",
            "value": "2026-10-12 20:21:20"
          },
          {
            "key": "StatusTime",
            "value": "2026-10-12 20:21:29"
          },
          {
            "key": "ProjectId",
            "value": "0"
          },
          {
            "key": "VpcId",
            "value": "vpc-38jxtmcs"
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
            "value": "{Region:ap-seoul,VpcId:vpc-38jxtmcs}"
          },
          {
            "key": "AddressIPVersion",
            "value": "ipv4"
          },
          {
            "key": "NumericalVpcId",
            "value": "15832685"
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
            "value": "2026-11-12 20:21:21"
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
      "createdTime": "2026-10-12T20:21:20Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "LoadBalancerId",
          "value": "lb-3zeo4wdb"
        },
        {
          "key": "LoadBalancerName",
          "value": "tbulrqdjeuctf5pl65g1"
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
          "value": "lb-3zeo4wdb-6401ixupewkjcl89.clb.kr-tencentclb.com"
        },
        {
          "key": "Status",
          "value": "1"
        },
        {
          "key": "CreateTime",
          "value": "2026-10-12 20:21:20"
        },
        {
          "key": "StatusTime",
          "value": "2026-10-12 20:21:29"
        },
        {
          "key": "ProjectId",
          "value": "0"
        },
        {
          "key": "VpcId",
          "value": "vpc-38jxtmcs"
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
          "value": "{Region:ap-seoul,VpcId:vpc-38jxtmcs}"
        },
        {
          "key": "AddressIPVersion",
          "value": "ipv4"
        },
        {
          "key": "NumericalVpcId",
          "value": "15832685"
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
          "value": "2026-11-12 20:21:21"
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
    "cspResourceName": "tbulrqdjeuctf5pl65g1",
    "cspResourceId": "lb-3zeo4wdb",
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
      "dnsName": "lb-3zeo4wdb-6401ixupewkjcl89.clb.kr-tencentclb.com",
      "keyValueList": [
        {
          "key": "LoadBalancerId",
          "value": "lb-3zeo4wdb"
        },
        {
          "key": "LoadBalancerName",
          "value": "tbulrqdjeuctf5pl65g1"
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
          "value": "lb-3zeo4wdb-6401ixupewkjcl89.clb.kr-tencentclb.com"
        },
        {
          "key": "Status",
          "value": "1"
        },
        {
          "key": "CreateTime",
          "value": "2026-10-12 20:21:20"
        },
        {
          "key": "StatusTime",
          "value": "2026-10-12 20:21:29"
        },
        {
          "key": "ProjectId",
          "value": "0"
        },
        {
          "key": "VpcId",
          "value": "vpc-38jxtmcs"
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
          "value": "{Region:ap-seoul,VpcId:vpc-38jxtmcs}"
        },
        {
          "key": "AddressIPVersion",
          "value": "ipv4"
        },
        {
          "key": "NumericalVpcId",
          "value": "15832685"
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
          "value": "2026-11-12 20:21:21"
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
    "createdTime": "2026-10-12T20:21:20Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "LoadBalancerId",
        "value": "lb-3zeo4wdb"
      },
      {
        "key": "LoadBalancerName",
        "value": "tbulrqdjeuctf5pl65g1"
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
        "value": "lb-3zeo4wdb-6401ixupewkjcl89.clb.kr-tencentclb.com"
      },
      {
        "key": "Status",
        "value": "1"
      },
      {
        "key": "CreateTime",
        "value": "2026-10-12 20:21:20"
      },
      {
        "key": "StatusTime",
        "value": "2026-10-12 20:21:29"
      },
      {
        "key": "ProjectId",
        "value": "0"
      },
      {
        "key": "VpcId",
        "value": "vpc-38jxtmcs"
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
        "value": "{Region:ap-seoul,VpcId:vpc-38jxtmcs}"
      },
      {
        "key": "AddressIPVersion",
        "value": "ipv4"
      },
      {
        "key": "NumericalVpcId",
        "value": "15832685"
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
        "value": "2026-11-12 20:21:21"
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
  "cspResourceName": "tbulrqdjeuctf5pl65g1",
  "cspResourceId": "lb-3zeo4wdb",
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
    "dnsName": "lb-3zeo4wdb-6401ixupewkjcl89.clb.kr-tencentclb.com",
    "keyValueList": [
      {
        "key": "LoadBalancerId",
        "value": "lb-3zeo4wdb"
      },
      {
        "key": "LoadBalancerName",
        "value": "tbulrqdjeuctf5pl65g1"
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
        "value": "lb-3zeo4wdb-6401ixupewkjcl89.clb.kr-tencentclb.com"
      },
      {
        "key": "Status",
        "value": "1"
      },
      {
        "key": "CreateTime",
        "value": "2026-10-12 20:21:20"
      },
      {
        "key": "StatusTime",
        "value": "2026-10-12 20:21:29"
      },
      {
        "key": "ProjectId",
        "value": "0"
      },
      {
        "key": "VpcId",
        "value": "vpc-38jxtmcs"
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
        "value": "{Region:ap-seoul,VpcId:vpc-38jxtmcs}"
      },
      {
        "key": "AddressIPVersion",
        "value": "ipv4"
      },
      {
        "key": "NumericalVpcId",
        "value": "15832685"
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
        "value": "2026-11-12 20:21:21"
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
  "createdTime": "2026-10-12T20:21:20Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "LoadBalancerId",
      "value": "lb-3zeo4wdb"
    },
    {
      "key": "LoadBalancerName",
      "value": "tbulrqdjeuctf5pl65g1"
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
      "value": "lb-3zeo4wdb-6401ixupewkjcl89.clb.kr-tencentclb.com"
    },
    {
      "key": "Status",
      "value": "1"
    },
    {
      "key": "CreateTime",
      "value": "2026-10-12 20:21:20"
    },
    {
      "key": "StatusTime",
      "value": "2026-10-12 20:21:29"
    },
    {
      "key": "ProjectId",
      "value": "0"
    },
    {
      "key": "VpcId",
      "value": "vpc-38jxtmcs"
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
      "value": "{Region:ap-seoul,VpcId:vpc-38jxtmcs}"
    },
    {
      "key": "AddressIPVersion",
      "value": "ipv4"
    },
    {
      "key": "NumericalVpcId",
      "value": "15832685"
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
      "value": "2026-11-12 20:21:21"
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

**Generated At:** 2026-10-12 12:23:11

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
| SA5.MEDIUM2 | 2 | 2.0 | - | x86_64 |  | $0.0600 | 1 |
| S6.LARGE16 | 4 | 16.0 | - | x86_64 |  | $0.2300 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| img-7rotv4ux | Ubuntu Server 22.04 LTS 64bit UEFI | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | ins-rn0jiohn | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.155.158.249<br>**Private IP:** 10.0.1.9<br>**SGs:** my05-mig-sg-02<br>**SSH:** my05-mig-sshkey-01 |
| my05-ng-influxdb-back-1 | ins-fgook4wl | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 43.133.249.165<br>**Private IP:** 10.0.1.4<br>**SGs:** my05-mig-sg-01<br>**SSH:** my05-mig-sshkey-01 |
| my05-ng-influxdb-back-2 | ins-flev5l93 | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-mig-vnet-01<br>**Subnet:** my05-mig-subnet-01<br>**Public IP:** 119.28.161.163<br>**Private IP:** 10.0.1.12<br>**SGs:** my05-mig-sg-01<br>**SSH:** my05-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my05-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my05-mig-vnet-01 |
| **CSP VNet ID** | vpc-38jxtmcs |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | tencent-ap-seoul |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my05-mig-subnet-01 | subnet-chmtrzmn | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my05-mig-sshkey-01 | skey-jab0kioj |  |  |

### Security Groups

#### Security Group: my05-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my05-mig-sg-01 |
| **CSP Security Group ID** | sg-jcefansv |
| **VNet** | my05-mig-vnet-01 |
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

#### Security Group: my05-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my05-mig-sg-02 |
| **CSP Security Group ID** | sg-op1dxapn |
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

