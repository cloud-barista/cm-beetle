# CM-Beetle test results for ALIBABA (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with ALIBABA cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.2+ (83efe14)
- imdl: v0.1.19+ (83efe14)
- CB-Tumblebug: v0.13.8
- CB-Spider: v0.13.20
- CB-MapUI: v0.13.14
- Target CSP: ALIBABA
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: October 30, 2026
- Test Time: 21:46:54 KST
- Test Execution: 2026-10-30 21:46:54 KST

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

## Test result for ALIBABA

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 1.678s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 120ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 56.735s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 22ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 3ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 9ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 2.755s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 10.421s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 7ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 6ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 1m35.541s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.085s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.093s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 16.171s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 49.313s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 4m53.336893195s

*Test executed on October 30, 2026 at 21:46:54 KST (2026-10-30 21:46:54 KST) using CM-Beetle automated test CLI*

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
  "desiredCsp": "alibaba",
  "desiredRegion": "ap-northeast-2",
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
        "csp": "alibaba",
        "region": "ap-northeast-2"
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
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "cloud_auto",
            "rootDiskSize": 40,
            "dataDiskIds": null
          },
          {
            "name": "ng-ec268ed7-821e-9d73-e79f-961262161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "cloud_auto",
            "rootDiskSize": 40,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "mig-vnet-01",
        "connectionName": "alibaba-ap-northeast-2",
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
        "connectionName": "alibaba-ap-northeast-2",
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
          "id": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "uid": "tbkvn2idlu4v9cdmvqsr",
          "cspSpecName": "ecs.e-c1m4.xlarge",
          "name": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1582,
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
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "ap-northeast-2a",
            "ap-northeast-2b",
            "ap-northeast-2c"
          ],
          "availabilityCheckedAt": "2026-10-30T12:47:00.486640572Z",
          "availabilityRefreshAfter": "2026-10-30T13:47:00.486640572Z"
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "uid": "tbs24l3o5mhfvco8hi0l",
          "cspSpecName": "ecs.e-c1m1.large",
          "name": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.0178,
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
          "rootDiskType": "",
          "rootDiskSize": -1,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "ap-northeast-2a",
            "ap-northeast-2b",
            "ap-northeast-2c"
          ],
          "availabilityCheckedAt": "2026-10-30T12:47:00.920832995Z",
          "availabilityRefreshAfter": "2026-10-30T13:47:00.920832995Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
          "regionList": [
            "ap-northeast-1",
            "ap-northeast-2",
            "ap-southeast-1",
            "ap-southeast-3",
            "ap-southeast-5",
            "ap-southeast-6",
            "ap-southeast-7",
            "ap-southeast-8",
            "cn-beijing",
            "cn-chengdu",
            "cn-fuzhou",
            "cn-guangzhou",
            "cn-hangzhou",
            "cn-heyuan",
            "cn-hongkong",
            "cn-huhehaote",
            "cn-nanjing",
            "cn-qingdao",
            "cn-shanghai",
            "cn-shenzhen",
            "cn-wuhan-lr",
            "cn-wulanchabu",
            "cn-zhangjiakou",
            "cn-zhongwei",
            "eu-central-1",
            "eu-west-1",
            "eu-west-2",
            "me-central-1",
            "me-east-1",
            "na-south-1",
            "us-east-1",
            "us-west-1"
          ],
          "id": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
          "uid": "tbnb1rmgtanm1d1cou0u",
          "name": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-us-west-1",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:58:28 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu  22.04 64 bit",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "BootMode",
              "value": "UEFI-Preferred"
            },
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260810.vhd"
            },
            {
              "key": "ImageOwnerAlias",
              "value": "system"
            },
            {
              "key": "OSName",
              "value": "Ubuntu  22.04 64位"
            },
            {
              "key": "OSNameEn",
              "value": "Ubuntu  22.04 64 bit"
            },
            {
              "key": "ImageFamily",
              "value": "acs:ubuntu_22_04_x64"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "IsSupportIoOptimized",
              "value": "true"
            },
            {
              "key": "Size",
              "value": "20"
            },
            {
              "key": "Description",
              "value": "Kernel version is 5.15.0-187-generic, 2026.8.13"
            },
            {
              "key": "Usage",
              "value": "instance"
            },
            {
              "key": "IsCopied",
              "value": "false"
            },
            {
              "key": "LoginAsNonRootSupported",
              "value": "true"
            },
            {
              "key": "ImageVersion",
              "value": "v2026.8.13"
            },
            {
              "key": "OSType",
              "value": "linux"
            },
            {
              "key": "IsSubscribed",
              "value": "false"
            },
            {
              "key": "IsSupportCloudinit",
              "value": "true"
            },
            {
              "key": "CreationTime",
              "value": "2026-08-13T01:52:07Z"
            },
            {
              "key": "Progress",
              "value": "100%"
            },
            {
              "key": "Platform",
              "value": "Ubuntu"
            },
            {
              "key": "ImageName",
              "value": "ubuntu_22_04_x64_20G_alibase_20260810.vhd"
            },
            {
              "key": "Status",
              "value": "Available"
            },
            {
              "key": "ImageOwnerId",
              "value": "0"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "DetectionOptions",
              "value": "{Status:,Items:{Item:null}}"
            },
            {
              "key": "Features",
              "value": "{MemoryOnlineUpgrade:unsupported,NvmeSupport:supported,CpuOnlineDowngrade:unsupported,ImdsSupport:v2,MemoryOnlineDowngrade:unsupported,CpuOnlineUpgrade:unsupported}"
            },
            {
              "key": "Tags",
              "value": "{Tag:[]}"
            },
            {
              "key": "DiskDeviceMappings",
              "value": "{DiskDeviceMapping:[]}"
            }
          ],
          "systemLabel": "",
          "description": "Kernel version is 5.15.0-187-generic, 2026.8.13",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "alibaba-ap-northeast-2",
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
          "connectionName": "alibaba-ap-northeast-2",
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
        "csp": "alibaba",
        "region": "ap-northeast-2"
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
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.g6.xlarge",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "cloud_auto",
            "rootDiskSize": 40,
            "dataDiskIds": null
          },
          {
            "name": "ng-ec268ed7-821e-9d73-e79f-961262161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "cloud_auto",
            "rootDiskSize": 40,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "mig-vnet-01",
        "connectionName": "alibaba-ap-northeast-2",
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
        "connectionName": "alibaba-ap-northeast-2",
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
          "id": "alibaba+ap-northeast-2+ecs.g6.xlarge",
          "uid": "tbkj6acd1vp4hlvbir3q",
          "cspSpecName": "ecs.g6.xlarge",
          "name": "alibaba+ap-northeast-2+ecs.g6.xlarge",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.160848,
          "orderInFilteredResult": 11,
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
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ],
          "availabilityStatus": "Unverified"
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "uid": "tbb6ajcho9h9f2g3quc7",
          "cspSpecName": "ecs.t6-c1m1.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.02139,
          "orderInFilteredResult": 6,
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
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ],
          "availabilityStatus": "Available",
          "availabilityReason": "zone(s) [ap-northeast-2b] unavailable/SoldOut; available in [ap-northeast-2a]",
          "availableZones": [
            "ap-northeast-2a"
          ],
          "unavailableZones": [
            "ap-northeast-2b"
          ],
          "availabilityCheckedAt": "2026-10-30T12:47:00.89089867Z",
          "availabilityRefreshAfter": "2026-10-30T13:47:00.89089867Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
          "regionList": [
            "ap-northeast-1",
            "ap-northeast-2",
            "ap-southeast-1",
            "ap-southeast-3",
            "ap-southeast-5",
            "ap-southeast-6",
            "ap-southeast-7",
            "ap-southeast-8",
            "cn-beijing",
            "cn-chengdu",
            "cn-fuzhou",
            "cn-guangzhou",
            "cn-hangzhou",
            "cn-heyuan",
            "cn-hongkong",
            "cn-huhehaote",
            "cn-nanjing",
            "cn-qingdao",
            "cn-shanghai",
            "cn-shenzhen",
            "cn-wuhan-lr",
            "cn-wulanchabu",
            "cn-zhangjiakou",
            "cn-zhongwei",
            "eu-central-1",
            "eu-west-1",
            "eu-west-2",
            "me-central-1",
            "me-east-1",
            "na-south-1",
            "us-east-1",
            "us-west-1"
          ],
          "id": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
          "uid": "tbnb1rmgtanm1d1cou0u",
          "name": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-us-west-1",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:58:28 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu  22.04 64 bit",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "BootMode",
              "value": "UEFI-Preferred"
            },
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260810.vhd"
            },
            {
              "key": "ImageOwnerAlias",
              "value": "system"
            },
            {
              "key": "OSName",
              "value": "Ubuntu  22.04 64位"
            },
            {
              "key": "OSNameEn",
              "value": "Ubuntu  22.04 64 bit"
            },
            {
              "key": "ImageFamily",
              "value": "acs:ubuntu_22_04_x64"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "IsSupportIoOptimized",
              "value": "true"
            },
            {
              "key": "Size",
              "value": "20"
            },
            {
              "key": "Description",
              "value": "Kernel version is 5.15.0-187-generic, 2026.8.13"
            },
            {
              "key": "Usage",
              "value": "instance"
            },
            {
              "key": "IsCopied",
              "value": "false"
            },
            {
              "key": "LoginAsNonRootSupported",
              "value": "true"
            },
            {
              "key": "ImageVersion",
              "value": "v2026.8.13"
            },
            {
              "key": "OSType",
              "value": "linux"
            },
            {
              "key": "IsSubscribed",
              "value": "false"
            },
            {
              "key": "IsSupportCloudinit",
              "value": "true"
            },
            {
              "key": "CreationTime",
              "value": "2026-08-13T01:52:07Z"
            },
            {
              "key": "Progress",
              "value": "100%"
            },
            {
              "key": "Platform",
              "value": "Ubuntu"
            },
            {
              "key": "ImageName",
              "value": "ubuntu_22_04_x64_20G_alibase_20260810.vhd"
            },
            {
              "key": "Status",
              "value": "Available"
            },
            {
              "key": "ImageOwnerId",
              "value": "0"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "DetectionOptions",
              "value": "{Status:,Items:{Item:null}}"
            },
            {
              "key": "Features",
              "value": "{MemoryOnlineUpgrade:unsupported,NvmeSupport:supported,CpuOnlineDowngrade:unsupported,ImdsSupport:v2,MemoryOnlineDowngrade:unsupported,CpuOnlineUpgrade:unsupported}"
            },
            {
              "key": "Tags",
              "value": "{Tag:[]}"
            },
            {
              "key": "DiskDeviceMappings",
              "value": "{DiskDeviceMapping:[]}"
            }
          ],
          "systemLabel": "",
          "description": "Kernel version is 5.15.0-187-generic, 2026.8.13",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "alibaba-ap-northeast-2",
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
          "connectionName": "alibaba-ap-northeast-2",
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
    "csp": "alibaba",
    "region": "ap-northeast-2"
  },
  "targetInfra": {
    "name": "my04-infra101",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "NLB-aware recommended infrastructure for cloud migration",
    "nodeGroups": [
      {
        "name": "my04-ng-influxdb-back",
        "nodeGroupSize": 2,
        "label": {
          "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "nlbBackend": "influxdb_back"
        },
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
        "connectionName": "alibaba-ap-northeast-2",
        "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
        "imageId": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
        "vNetId": "my04-mig-vnet-01",
        "subnetId": "my04-mig-subnet-01",
        "securityGroupIds": [
          "my04-mig-sg-01"
        ],
        "sshKeyId": "my04-mig-sshkey-01",
        "rootDiskType": "cloud_auto",
        "rootDiskSize": 40,
        "dataDiskIds": null
      },
      {
        "name": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
        "connectionName": "alibaba-ap-northeast-2",
        "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
        "imageId": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
        "vNetId": "my04-mig-vnet-01",
        "subnetId": "my04-mig-subnet-01",
        "securityGroupIds": [
          "my04-mig-sg-02"
        ],
        "sshKeyId": "my04-mig-sshkey-01",
        "rootDiskType": "cloud_auto",
        "rootDiskSize": 40,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "my04-mig-vnet-01",
    "connectionName": "alibaba-ap-northeast-2",
    "cidrBlock": "10.0.0.0/21",
    "subnetInfoList": [
      {
        "name": "my04-mig-subnet-01",
        "ipv4_CIDR": "10.0.1.0/24",
        "description": "a recommended subnet for migration"
      }
    ],
    "description": "a recommended vNet for migration"
  },
  "targetSshKey": {
    "name": "my04-mig-sshkey-01",
    "connectionName": "alibaba-ap-northeast-2",
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
      "id": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
      "uid": "tbkvn2idlu4v9cdmvqsr",
      "cspSpecName": "ecs.e-c1m4.xlarge",
      "name": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
      "namespace": "system",
      "connectionName": "alibaba-ap-northeast-2",
      "providerName": "alibaba",
      "regionName": "ap-northeast-2",
      "regionLatitude": 37.36,
      "regionLongitude": 126.78,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 16,
      "diskSizeGB": -1,
      "costPerHour": 0.1582,
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
          "key": "CpuArchitecture",
          "value": "X86"
        }
      ],
      "availabilityStatus": "Available",
      "availableZones": [
        "ap-northeast-2a",
        "ap-northeast-2b",
        "ap-northeast-2c"
      ],
      "availabilityCheckedAt": "2026-10-30T12:47:00.486640572Z",
      "availabilityRefreshAfter": "2026-10-30T13:47:00.486640572Z"
    },
    {
      "id": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
      "uid": "tbs24l3o5mhfvco8hi0l",
      "cspSpecName": "ecs.e-c1m1.large",
      "name": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
      "namespace": "system",
      "connectionName": "alibaba-ap-northeast-2",
      "providerName": "alibaba",
      "regionName": "ap-northeast-2",
      "regionLatitude": 37.36,
      "regionLongitude": 126.78,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 2,
      "diskSizeGB": -1,
      "costPerHour": 0.0178,
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
      "rootDiskType": "",
      "rootDiskSize": -1,
      "systemLabel": "auto-gen",
      "details": [
        {
          "key": "CpuArchitecture",
          "value": "X86"
        }
      ],
      "availabilityStatus": "Available",
      "availableZones": [
        "ap-northeast-2a",
        "ap-northeast-2b",
        "ap-northeast-2c"
      ],
      "availabilityCheckedAt": "2026-10-30T12:47:00.920832995Z",
      "availabilityRefreshAfter": "2026-10-30T13:47:00.920832995Z"
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "alibaba",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
      "regionList": [
        "ap-northeast-1",
        "ap-northeast-2",
        "ap-southeast-1",
        "ap-southeast-3",
        "ap-southeast-5",
        "ap-southeast-6",
        "ap-southeast-7",
        "ap-southeast-8",
        "cn-beijing",
        "cn-chengdu",
        "cn-fuzhou",
        "cn-guangzhou",
        "cn-hangzhou",
        "cn-heyuan",
        "cn-hongkong",
        "cn-huhehaote",
        "cn-nanjing",
        "cn-qingdao",
        "cn-shanghai",
        "cn-shenzhen",
        "cn-wuhan-lr",
        "cn-wulanchabu",
        "cn-zhangjiakou",
        "cn-zhongwei",
        "eu-central-1",
        "eu-west-1",
        "eu-west-2",
        "me-central-1",
        "me-east-1",
        "na-south-1",
        "us-east-1",
        "us-west-1"
      ],
      "id": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
      "uid": "tbnb1rmgtanm1d1cou0u",
      "name": "ubuntu_22_04_x64_20G_alibase_20260810.vhd",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "alibaba-us-west-1",
      "infraType": "",
      "fetchedTime": "2026.08.21 13:58:28 Fri",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Ubuntu  22.04 64 bit",
      "osDiskType": "NA",
      "osDiskSizeGB": 20,
      "imageStatus": "Available",
      "details": [
        {
          "key": "BootMode",
          "value": "UEFI-Preferred"
        },
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260810.vhd"
        },
        {
          "key": "ImageOwnerAlias",
          "value": "system"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "ImageFamily",
          "value": "acs:ubuntu_22_04_x64"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "IsSupportIoOptimized",
          "value": "true"
        },
        {
          "key": "Size",
          "value": "20"
        },
        {
          "key": "Description",
          "value": "Kernel version is 5.15.0-187-generic, 2026.8.13"
        },
        {
          "key": "Usage",
          "value": "instance"
        },
        {
          "key": "IsCopied",
          "value": "false"
        },
        {
          "key": "LoginAsNonRootSupported",
          "value": "true"
        },
        {
          "key": "ImageVersion",
          "value": "v2026.8.13"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "IsSubscribed",
          "value": "false"
        },
        {
          "key": "IsSupportCloudinit",
          "value": "true"
        },
        {
          "key": "CreationTime",
          "value": "2026-08-13T01:52:07Z"
        },
        {
          "key": "Progress",
          "value": "100%"
        },
        {
          "key": "Platform",
          "value": "Ubuntu"
        },
        {
          "key": "ImageName",
          "value": "ubuntu_22_04_x64_20G_alibase_20260810.vhd"
        },
        {
          "key": "Status",
          "value": "Available"
        },
        {
          "key": "ImageOwnerId",
          "value": "0"
        },
        {
          "key": "IsPublic",
          "value": "true"
        },
        {
          "key": "DetectionOptions",
          "value": "{Status:,Items:{Item:null}}"
        },
        {
          "key": "Features",
          "value": "{MemoryOnlineUpgrade:unsupported,NvmeSupport:supported,CpuOnlineDowngrade:unsupported,ImdsSupport:v2,MemoryOnlineDowngrade:unsupported,CpuOnlineUpgrade:unsupported}"
        },
        {
          "key": "Tags",
          "value": "{Tag:[]}"
        },
        {
          "key": "DiskDeviceMappings",
          "value": "{DiskDeviceMapping:[]}"
        }
      ],
      "systemLabel": "",
      "description": "Kernel version is 5.15.0-187-generic, 2026.8.13",
      "commandHistory": null
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "my04-mig-sg-01",
      "connectionName": "alibaba-ap-northeast-2",
      "vNetId": "my04-mig-vnet-01",
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
      "name": "my04-mig-sg-02",
      "connectionName": "alibaba-ap-northeast-2",
      "vNetId": "my04-mig-vnet-01",
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
- **MCI ID**: `my04-infra101`
- **MCI Name**: `my04-infra101`
- **Status**: `Running:3 (R:3/3)`

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:47:32",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "/dev/xvda",
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
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
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
      "cspSshKeyId": "tb5aoempt96t422gdonc",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "eAthv3!k8s$bs1",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHpI4oMvPNvfmXHJ8g2GY0iiAbDeyTvkoFqlE5QjvE+lbTa+uCTZ1VivCfM/VFpcEW8lVebG4mbjatzns1bY+f8=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:c3AJdLaRmNxuvOPlgOuLZ1wE4y+D7aiUvI93hakrDfU",
        "firstUsedAt": "2026-10-30T12:47:54Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:34Z",
          "completedTime": "2026-10-30T12:48:00Z",
          "elapsedTime": 26,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my04-infra101-tbgo3rtbdn3dhp7adie6",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:00Z",
          "completedTime": "2026-10-30T12:48:02Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7c2uilk7pplm0pa0qoZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:47:33",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "/dev/xvda",
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
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
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
      "cspSshKeyId": "tb5aoempt96t422gdonc",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "!b0kuAlt15dd$g",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBABkaQO4fMpNka900n69E8KrOiT/P23ddXvHRcTCNQNnLUiznjAm5OGCmG2vIDu4oB6TGit5AwSSfuGXqbJRlcs=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:XOeJB152l4Ingr/wJLpx8cK+jCEDXa+jx6XVqQTbH9w",
        "firstUsedAt": "2026-10-30T12:47:57Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:34Z",
          "completedTime": "2026-10-30T12:47:58Z",
          "elapsedTime": 24,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my04-infra101-tbgo3rtbdn3dhp7adie6",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:00Z",
          "completedTime": "2026-10-30T12:48:01Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7hb03d1frhdo6k7uttZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:47:34",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "/dev/xvda",
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
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
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
      "cspSshKeyId": "tb5aoempt96t422gdonc",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "etA2$egtvf5!b1",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDNPyCK1zxJgwjzkgNZ0TvTM3CO2ke/PwmQy8SNl1g2RUwbMiL/rjolVaTeqTW0OcbSS2xXnrc7+FgduFPcI4Bg=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:NTvqaLeKUcdemg0p4tBt2uWg0hybNgSLdaSoTPOtYTM",
        "firstUsedAt": "2026-10-30T12:47:58Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:34Z",
          "completedTime": "2026-10-30T12:47:59Z",
          "elapsedTime": 25,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my04-infra101-tbgo3rtbdn3dhp7adie6",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:00Z",
          "completedTime": "2026-10-30T12:48:01Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7hb03d1frhdo6k7utuZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my04-mig-vnet-01",
      "name": "my04-mig-vnet-01",
      "infraId": "my04-infra101",
      "vNetId": "my04-mig-vnet-01",
      "connectionNames": [
        "alibaba-ap-northeast-2"
      ],
      "providerNames": [
        "alibaba"
      ],
      "regionNames": [
        "ap-northeast-2"
      ],
      "nodeGroupIds": [
        "my04-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my04-ng-influxdb-back"
      ],
      "nodeIds": [
        "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my04-ng-influxdb-back-1",
        "my04-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-influxdb-back-1",
            "nodeIp": "8.213.158.198",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7hb03d1frhdo6k7uttZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-influxdb-back-2",
            "nodeIp": "47.80.21.196",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7hb03d1frhdo6k7utuZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "8.220.204.23",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7c2uilk7pplm0pa0qoZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my04-infra101-tbgo3rtbdn3dhp7adie6"
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
- **Count**: 4

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
          "id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbuku56fj76km7vqa0su",
          "name": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
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
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
          "cspVNetId": "",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my05-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-1",
          "uid": "tb42prmq38mkramqj7kf",
          "name": "my05-ng-influxdb-back-1",
          "nodeGroupId": "my05-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
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
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
          "cspVNetId": "",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-2",
          "uid": "tbho8c20nshb3d7pklqn",
          "name": "my05-ng-influxdb-back-2",
          "nodeGroupId": "my05-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
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
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
          "cspVNetId": "",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
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
- **IDs**: [my01-infra101 my02-infra101 my04-infra101 my05-infra101]

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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:47:32",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "/dev/xvda",
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
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
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
      "cspSshKeyId": "tb5aoempt96t422gdonc",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "eAthv3!k8s$bs1",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHpI4oMvPNvfmXHJ8g2GY0iiAbDeyTvkoFqlE5QjvE+lbTa+uCTZ1VivCfM/VFpcEW8lVebG4mbjatzns1bY+f8=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:c3AJdLaRmNxuvOPlgOuLZ1wE4y+D7aiUvI93hakrDfU",
        "firstUsedAt": "2026-10-30T12:47:54Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:34Z",
          "completedTime": "2026-10-30T12:48:00Z",
          "elapsedTime": 26,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my04-infra101-tbgo3rtbdn3dhp7adie6",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:00Z",
          "completedTime": "2026-10-30T12:48:02Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7c2uilk7pplm0pa0qoZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:47:33",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "/dev/xvda",
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
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
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
      "cspSshKeyId": "tb5aoempt96t422gdonc",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "!b0kuAlt15dd$g",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBABkaQO4fMpNka900n69E8KrOiT/P23ddXvHRcTCNQNnLUiznjAm5OGCmG2vIDu4oB6TGit5AwSSfuGXqbJRlcs=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:XOeJB152l4Ingr/wJLpx8cK+jCEDXa+jx6XVqQTbH9w",
        "firstUsedAt": "2026-10-30T12:47:57Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:34Z",
          "completedTime": "2026-10-30T12:47:58Z",
          "elapsedTime": 24,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my04-infra101-tbgo3rtbdn3dhp7adie6",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:00Z",
          "completedTime": "2026-10-30T12:48:01Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7hb03d1frhdo6k7uttZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:47:34",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
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
      "rootDeviceName": "/dev/xvda",
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
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
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
      "cspSshKeyId": "tb5aoempt96t422gdonc",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "etA2$egtvf5!b1",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDNPyCK1zxJgwjzkgNZ0TvTM3CO2ke/PwmQy8SNl1g2RUwbMiL/rjolVaTeqTW0OcbSS2xXnrc7+FgduFPcI4Bg=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:NTvqaLeKUcdemg0p4tBt2uWg0hybNgSLdaSoTPOtYTM",
        "firstUsedAt": "2026-10-30T12:47:58Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:47:34Z",
          "completedTime": "2026-10-30T12:47:59Z",
          "elapsedTime": 25,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my04-infra101-tbgo3rtbdn3dhp7adie6",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:48:00Z",
          "completedTime": "2026-10-30T12:48:01Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7hb03d1frhdo6k7utuZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my04-mig-vnet-01",
      "name": "my04-mig-vnet-01",
      "infraId": "my04-infra101",
      "vNetId": "my04-mig-vnet-01",
      "connectionNames": [
        "alibaba-ap-northeast-2"
      ],
      "providerNames": [
        "alibaba"
      ],
      "regionNames": [
        "ap-northeast-2"
      ],
      "nodeGroupIds": [
        "my04-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my04-ng-influxdb-back"
      ],
      "nodeIds": [
        "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my04-ng-influxdb-back-1",
        "my04-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-influxdb-back-1",
            "nodeIp": "8.213.158.198",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7hb03d1frhdo6k7uttZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-influxdb-back-2",
            "nodeIp": "47.80.21.196",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7hb03d1frhdo6k7utuZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "8.220.204.23",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7c2uilk7pplm0pa0qoZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my04-infra101-tbgo3rtbdn3dhp7adie6"
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
      "id": "my04-ng-influxdb-back",
      "cspResourceName": "tb41oq05vbnl3nvecm0m",
      "cspResourceId": "lb-mj76vt39sjkyn7jraj5gv",
      "name": "my04-ng-influxdb-back",
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
      "type": "PUBLIC",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "ip": "43.108.98.145",
        "port": "9999"
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my04-ng-influxdb-back",
        "nodes": [
          "my04-ng-influxdb-back-1",
          "my04-ng-influxdb-back-2"
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
            "key": "IId",
            "value": "{NameId:tb41oq05vbnl3nvecm0m,SystemId:lb-mj76vt39sjkyn7jraj5gv}"
          },
          {
            "key": "VpcIID",
            "value": "{NameId:,SystemId:}"
          },
          {
            "key": "Listener",
            "value": "{Protocol:,IP:,Port:,DNSName:}"
          },
          {
            "key": "VMGroup",
            "value": "{Protocol:tcp,Port:8086,VMs:null}"
          },
          {
            "key": "HealthChecker",
            "value": "{Protocol:tcp,Port:8086,Interval:10,Timeout:10,Threshold:3}"
          },
          {
            "key": "CreatedTime",
            "value": "0001-01-01T00:00:00Z"
          },
          {
            "key": "KeyValueList",
            "value": "{Key:BaseResponse,Value:{}}; {Key:Status,Value:running}; {Key:ConnectionDrainTimeout,Value:0}; {Key:HealthCheckTcpFastCloseEnabled,Value:false}; {Key:FullNatEnabled,Value:false}; {Key:RequestId,Value:01A18CDA-6A62-3376-BBF3-8DE9BFA5E74B}; {Key:HealthCheckConnectPort,Value:8086}; {Key:Bandwidth,Value:-1}; {Key:HealthCheckType,Value:tcp}; {Key:BackendServerPort,Value:8086}; {Key:AclStatus,Value:off}; {Key:UnhealthyThreshold,Value:3}; {Key:MaxConnection,Value:0}; {Key:ProxyProtocolV2Enabled,Value:false}; {Key:PersistenceTimeout,Value:0}; {Key:ListenerPort,Value:9999}; {Key:HealthCheckInterval,Value:10}; {Key:FailoverThreshold,Value:0}; {Key:Scheduler,Value:wrr}; {Key:EstablishedTimeout,Value:900}; {Key:HealthCheckConnectTimeout,Value:10}; {Key:MasterSlaveModeEnabled,Value:false}; {Key:HealthyThreshold,Value:3}; {Key:HealthCheck,Value:on}; {Key:LoadBalancerId,Value:lb-mj76vt39sjkyn7jraj5gv}; {Key:AclIds,Value:{AclId:null}}; {Key:PortRanges,Value:{PortRange:[]}}; {Key:Tags,Value:{Tag:[]}}"
          }
        ]
      },
      "createdTime": "2026-10-30T12:48:03Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "BaseResponse",
          "value": "{}"
        },
        {
          "key": "CreateTimeStamp",
          "value": "1793364483000"
        },
        {
          "key": "HasReservedInfo",
          "value": "false"
        },
        {
          "key": "CreateTime",
          "value": "2026-10-30T12:48:03Z"
        },
        {
          "key": "LoadBalancerId",
          "value": "lb-mj76vt39sjkyn7jraj5gv"
        },
        {
          "key": "PayType",
          "value": "PayOnDemand"
        },
        {
          "key": "AddressType",
          "value": "internet"
        },
        {
          "key": "SupportPrivateLink",
          "value": "false"
        },
        {
          "key": "NetworkType",
          "value": "classic"
        },
        {
          "key": "SpecBpsFlag",
          "value": "false"
        },
        {
          "key": "AddressIPVersion",
          "value": "ipv4"
        },
        {
          "key": "RequestId",
          "value": "01A18CDA-69B3-3179-9D44-816628E5D303"
        },
        {
          "key": "Bandwidth",
          "value": "5120"
        },
        {
          "key": "LoadBalancerName",
          "value": "tb41oq05vbnl3nvecm0m"
        },
        {
          "key": "Address",
          "value": "43.108.98.145"
        },
        {
          "key": "SlaveZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "EndTimeStamp",
          "value": "32493801600000"
        },
        {
          "key": "MasterZoneId",
          "value": "ap-northeast-2b"
        },
        {
          "key": "LoadBalancerSpec",
          "value": "slb.s1.small"
        },
        {
          "key": "AutoReleaseTime",
          "value": "0"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "ModificationProtectionStatus",
          "value": "NonProtection"
        },
        {
          "key": "EndTime",
          "value": "2999-09-08T16:00:00Z"
        },
        {
          "key": "LoadBalancerStatus",
          "value": "active"
        },
        {
          "key": "ResourceGroupId",
          "value": "rg-acfnvekhilw5kmy"
        },
        {
          "key": "InternetChargeType",
          "value": "paybytraffic"
        },
        {
          "key": "BusinessStatus",
          "value": "Normal"
        },
        {
          "key": "DeleteProtection",
          "value": "off"
        },
        {
          "key": "RegionIdAlias",
          "value": "ap-northeast-2"
        },
        {
          "key": "RenewalDuration",
          "value": "0"
        },
        {
          "key": "CloudInstanceUid",
          "value": "0"
        },
        {
          "key": "InstanceChargeType",
          "value": "PayBySpec"
        },
        {
          "key": "Labels",
          "value": "{Label:null}"
        },
        {
          "key": "ListenerPorts",
          "value": "{ListenerPort:[9999]}"
        },
        {
          "key": "Tags",
          "value": "{Tag:[]}"
        },
        {
          "key": "ListenerPortsAndProtocal",
          "value": "{ListenerPortAndProtocal:[{ListenerProtocal:tcp,ListenerPort:9999}]}"
        },
        {
          "key": "ListenerPortsAndProtocol",
          "value": "{ListenerPortAndProtocol:[{ListenerPort:9999,ListenerProtocol:tcp,ListenerForward:,Description:,ForwardPort:0}]}"
        },
        {
          "key": "BackendServers",
          "value": "{BackendServer:[{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7hb03d1frhdo6k7utu},{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7hb03d1frhdo6k7utt}]}"
        }
      ],
      "isAutoGenerated": false,
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
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
    "id": "my04-ng-influxdb-back",
    "cspResourceName": "tb41oq05vbnl3nvecm0m",
    "cspResourceId": "lb-mj76vt39sjkyn7jraj5gv",
    "name": "my04-ng-influxdb-back",
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
    "type": "PUBLIC",
    "scope": "REGION",
    "listener": {
      "protocol": "TCP",
      "ip": "43.108.98.145",
      "port": "9999"
    },
    "targetGroup": {
      "protocol": "TCP",
      "port": "8086",
      "nodeGroupId": "my04-ng-influxdb-back",
      "nodes": [
        "my04-ng-influxdb-back-1",
        "my04-ng-influxdb-back-2"
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
          "key": "IId",
          "value": "{NameId:tb41oq05vbnl3nvecm0m,SystemId:lb-mj76vt39sjkyn7jraj5gv}"
        },
        {
          "key": "VpcIID",
          "value": "{NameId:,SystemId:}"
        },
        {
          "key": "Listener",
          "value": "{Protocol:,IP:,Port:,DNSName:}"
        },
        {
          "key": "VMGroup",
          "value": "{Protocol:tcp,Port:8086,VMs:null}"
        },
        {
          "key": "HealthChecker",
          "value": "{Protocol:tcp,Port:8086,Interval:10,Timeout:10,Threshold:3}"
        },
        {
          "key": "CreatedTime",
          "value": "0001-01-01T00:00:00Z"
        },
        {
          "key": "KeyValueList",
          "value": "{Key:BaseResponse,Value:{}}; {Key:Status,Value:running}; {Key:ConnectionDrainTimeout,Value:0}; {Key:HealthCheckTcpFastCloseEnabled,Value:false}; {Key:FullNatEnabled,Value:false}; {Key:RequestId,Value:01A18CDA-6A62-3376-BBF3-8DE9BFA5E74B}; {Key:HealthCheckConnectPort,Value:8086}; {Key:Bandwidth,Value:-1}; {Key:HealthCheckType,Value:tcp}; {Key:BackendServerPort,Value:8086}; {Key:AclStatus,Value:off}; {Key:UnhealthyThreshold,Value:3}; {Key:MaxConnection,Value:0}; {Key:ProxyProtocolV2Enabled,Value:false}; {Key:PersistenceTimeout,Value:0}; {Key:ListenerPort,Value:9999}; {Key:HealthCheckInterval,Value:10}; {Key:FailoverThreshold,Value:0}; {Key:Scheduler,Value:wrr}; {Key:EstablishedTimeout,Value:900}; {Key:HealthCheckConnectTimeout,Value:10}; {Key:MasterSlaveModeEnabled,Value:false}; {Key:HealthyThreshold,Value:3}; {Key:HealthCheck,Value:on}; {Key:LoadBalancerId,Value:lb-mj76vt39sjkyn7jraj5gv}; {Key:AclIds,Value:{AclId:null}}; {Key:PortRanges,Value:{PortRange:[]}}; {Key:Tags,Value:{Tag:[]}}"
        }
      ]
    },
    "createdTime": "2026-10-30T12:48:03Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "BaseResponse",
        "value": "{}"
      },
      {
        "key": "CreateTimeStamp",
        "value": "1793364483000"
      },
      {
        "key": "HasReservedInfo",
        "value": "false"
      },
      {
        "key": "CreateTime",
        "value": "2026-10-30T12:48:03Z"
      },
      {
        "key": "LoadBalancerId",
        "value": "lb-mj76vt39sjkyn7jraj5gv"
      },
      {
        "key": "PayType",
        "value": "PayOnDemand"
      },
      {
        "key": "AddressType",
        "value": "internet"
      },
      {
        "key": "SupportPrivateLink",
        "value": "false"
      },
      {
        "key": "NetworkType",
        "value": "classic"
      },
      {
        "key": "SpecBpsFlag",
        "value": "false"
      },
      {
        "key": "AddressIPVersion",
        "value": "ipv4"
      },
      {
        "key": "RequestId",
        "value": "01A18CDA-69B3-3179-9D44-816628E5D303"
      },
      {
        "key": "Bandwidth",
        "value": "5120"
      },
      {
        "key": "LoadBalancerName",
        "value": "tb41oq05vbnl3nvecm0m"
      },
      {
        "key": "Address",
        "value": "43.108.98.145"
      },
      {
        "key": "SlaveZoneId",
        "value": "ap-northeast-2a"
      },
      {
        "key": "EndTimeStamp",
        "value": "32493801600000"
      },
      {
        "key": "MasterZoneId",
        "value": "ap-northeast-2b"
      },
      {
        "key": "LoadBalancerSpec",
        "value": "slb.s1.small"
      },
      {
        "key": "AutoReleaseTime",
        "value": "0"
      },
      {
        "key": "RegionId",
        "value": "ap-northeast-2"
      },
      {
        "key": "ModificationProtectionStatus",
        "value": "NonProtection"
      },
      {
        "key": "EndTime",
        "value": "2999-09-08T16:00:00Z"
      },
      {
        "key": "LoadBalancerStatus",
        "value": "active"
      },
      {
        "key": "ResourceGroupId",
        "value": "rg-acfnvekhilw5kmy"
      },
      {
        "key": "InternetChargeType",
        "value": "paybytraffic"
      },
      {
        "key": "BusinessStatus",
        "value": "Normal"
      },
      {
        "key": "DeleteProtection",
        "value": "off"
      },
      {
        "key": "RegionIdAlias",
        "value": "ap-northeast-2"
      },
      {
        "key": "RenewalDuration",
        "value": "0"
      },
      {
        "key": "CloudInstanceUid",
        "value": "0"
      },
      {
        "key": "InstanceChargeType",
        "value": "PayBySpec"
      },
      {
        "key": "Labels",
        "value": "{Label:null}"
      },
      {
        "key": "ListenerPorts",
        "value": "{ListenerPort:[9999]}"
      },
      {
        "key": "Tags",
        "value": "{Tag:[]}"
      },
      {
        "key": "ListenerPortsAndProtocal",
        "value": "{ListenerPortAndProtocal:[{ListenerProtocal:tcp,ListenerPort:9999}]}"
      },
      {
        "key": "ListenerPortsAndProtocol",
        "value": "{ListenerPortAndProtocol:[{ListenerPort:9999,ListenerProtocol:tcp,ListenerForward:,Description:,ForwardPort:0}]}"
      },
      {
        "key": "BackendServers",
        "value": "{BackendServer:[{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7hb03d1frhdo6k7utu},{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7hb03d1frhdo6k7utt}]}"
      }
    ],
    "isAutoGenerated": false,
    "location": {
      "display": "South Korea (Seoul)",
      "latitude": 37.36,
      "longitude": 126.78
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
  "id": "my04-ng-influxdb-back",
  "cspResourceName": "tb41oq05vbnl3nvecm0m",
  "cspResourceId": "lb-mj76vt39sjkyn7jraj5gv",
  "name": "my04-ng-influxdb-back",
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
  "type": "PUBLIC",
  "scope": "REGION",
  "listener": {
    "protocol": "TCP",
    "ip": "43.108.98.145",
    "port": "9999"
  },
  "targetGroup": {
    "protocol": "TCP",
    "port": "8086",
    "nodeGroupId": "my04-ng-influxdb-back",
    "nodes": [
      "my04-ng-influxdb-back-1",
      "my04-ng-influxdb-back-2"
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
        "key": "IId",
        "value": "{NameId:tb41oq05vbnl3nvecm0m,SystemId:lb-mj76vt39sjkyn7jraj5gv}"
      },
      {
        "key": "VpcIID",
        "value": "{NameId:,SystemId:}"
      },
      {
        "key": "Listener",
        "value": "{Protocol:,IP:,Port:,DNSName:}"
      },
      {
        "key": "VMGroup",
        "value": "{Protocol:tcp,Port:8086,VMs:null}"
      },
      {
        "key": "HealthChecker",
        "value": "{Protocol:tcp,Port:8086,Interval:10,Timeout:10,Threshold:3}"
      },
      {
        "key": "CreatedTime",
        "value": "0001-01-01T00:00:00Z"
      },
      {
        "key": "KeyValueList",
        "value": "{Key:BaseResponse,Value:{}}; {Key:Status,Value:running}; {Key:ConnectionDrainTimeout,Value:0}; {Key:HealthCheckTcpFastCloseEnabled,Value:false}; {Key:FullNatEnabled,Value:false}; {Key:RequestId,Value:01A18CDA-6A62-3376-BBF3-8DE9BFA5E74B}; {Key:HealthCheckConnectPort,Value:8086}; {Key:Bandwidth,Value:-1}; {Key:HealthCheckType,Value:tcp}; {Key:BackendServerPort,Value:8086}; {Key:AclStatus,Value:off}; {Key:UnhealthyThreshold,Value:3}; {Key:MaxConnection,Value:0}; {Key:ProxyProtocolV2Enabled,Value:false}; {Key:PersistenceTimeout,Value:0}; {Key:ListenerPort,Value:9999}; {Key:HealthCheckInterval,Value:10}; {Key:FailoverThreshold,Value:0}; {Key:Scheduler,Value:wrr}; {Key:EstablishedTimeout,Value:900}; {Key:HealthCheckConnectTimeout,Value:10}; {Key:MasterSlaveModeEnabled,Value:false}; {Key:HealthyThreshold,Value:3}; {Key:HealthCheck,Value:on}; {Key:LoadBalancerId,Value:lb-mj76vt39sjkyn7jraj5gv}; {Key:AclIds,Value:{AclId:null}}; {Key:PortRanges,Value:{PortRange:[]}}; {Key:Tags,Value:{Tag:[]}}"
      }
    ]
  },
  "createdTime": "2026-10-30T12:48:03Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "BaseResponse",
      "value": "{}"
    },
    {
      "key": "CreateTimeStamp",
      "value": "1793364483000"
    },
    {
      "key": "HasReservedInfo",
      "value": "false"
    },
    {
      "key": "CreateTime",
      "value": "2026-10-30T12:48:03Z"
    },
    {
      "key": "LoadBalancerId",
      "value": "lb-mj76vt39sjkyn7jraj5gv"
    },
    {
      "key": "PayType",
      "value": "PayOnDemand"
    },
    {
      "key": "AddressType",
      "value": "internet"
    },
    {
      "key": "SupportPrivateLink",
      "value": "false"
    },
    {
      "key": "NetworkType",
      "value": "classic"
    },
    {
      "key": "SpecBpsFlag",
      "value": "false"
    },
    {
      "key": "AddressIPVersion",
      "value": "ipv4"
    },
    {
      "key": "RequestId",
      "value": "01A18CDA-69B3-3179-9D44-816628E5D303"
    },
    {
      "key": "Bandwidth",
      "value": "5120"
    },
    {
      "key": "LoadBalancerName",
      "value": "tb41oq05vbnl3nvecm0m"
    },
    {
      "key": "Address",
      "value": "43.108.98.145"
    },
    {
      "key": "SlaveZoneId",
      "value": "ap-northeast-2a"
    },
    {
      "key": "EndTimeStamp",
      "value": "32493801600000"
    },
    {
      "key": "MasterZoneId",
      "value": "ap-northeast-2b"
    },
    {
      "key": "LoadBalancerSpec",
      "value": "slb.s1.small"
    },
    {
      "key": "AutoReleaseTime",
      "value": "0"
    },
    {
      "key": "RegionId",
      "value": "ap-northeast-2"
    },
    {
      "key": "ModificationProtectionStatus",
      "value": "NonProtection"
    },
    {
      "key": "EndTime",
      "value": "2999-09-08T16:00:00Z"
    },
    {
      "key": "LoadBalancerStatus",
      "value": "active"
    },
    {
      "key": "ResourceGroupId",
      "value": "rg-acfnvekhilw5kmy"
    },
    {
      "key": "InternetChargeType",
      "value": "paybytraffic"
    },
    {
      "key": "BusinessStatus",
      "value": "Normal"
    },
    {
      "key": "DeleteProtection",
      "value": "off"
    },
    {
      "key": "RegionIdAlias",
      "value": "ap-northeast-2"
    },
    {
      "key": "RenewalDuration",
      "value": "0"
    },
    {
      "key": "CloudInstanceUid",
      "value": "0"
    },
    {
      "key": "InstanceChargeType",
      "value": "PayBySpec"
    },
    {
      "key": "Labels",
      "value": "{Label:null}"
    },
    {
      "key": "ListenerPorts",
      "value": "{ListenerPort:[9999]}"
    },
    {
      "key": "Tags",
      "value": "{Tag:[]}"
    },
    {
      "key": "ListenerPortsAndProtocal",
      "value": "{ListenerPortAndProtocal:[{ListenerProtocal:tcp,ListenerPort:9999}]}"
    },
    {
      "key": "ListenerPortsAndProtocol",
      "value": "{ListenerPortAndProtocol:[{ListenerPort:9999,ListenerProtocol:tcp,ListenerForward:,Description:,ForwardPort:0}]}"
    },
    {
      "key": "BackendServers",
      "value": "{BackendServer:[{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7hb03d1frhdo6k7utu},{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7hb03d1frhdo6k7utt}]}"
    }
  ],
  "isAutoGenerated": false,
  "location": {
    "display": "South Korea (Seoul)",
    "latitude": 37.36,
    "longitude": 126.78
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

**Generated At:** 2026-10-30 12:50:37

**Namespace:** mig01

**Infra Name:** my04-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my04-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | ALIBABA |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| ecs.e-c1m1.large | 2 | 2.0 | - | x86_64 |  | $0.0178 | 1 |
| ecs.e-c1m4.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1582 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ubuntu_22_04_x64_20G_alibase_20260810.vhd | Ubuntu  22.04 64 bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | i-mj7c2uilk7pplm0pa0qo | Running | 2 vCPU, 2.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 8.220.204.23<br>**Private IP:** 10.0.1.142<br>**SGs:** my04-mig-sg-02<br>**SSH:** my04-mig-sshkey-01 |
| my04-ng-influxdb-back-1 | i-mj7hb03d1frhdo6k7utt | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 8.213.158.198<br>**Private IP:** 10.0.1.141<br>**SGs:** my04-mig-sg-01<br>**SSH:** my04-mig-sshkey-01 |
| my04-ng-influxdb-back-2 | i-mj7hb03d1frhdo6k7utu | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 47.80.21.196<br>**Private IP:** 10.0.1.140<br>**SGs:** my04-mig-sg-01<br>**SSH:** my04-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my04-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my04-mig-vnet-01 |
| **CSP VNet ID** | vpc-mj7xwzw3fbv61c7bto7le |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | alibaba-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my04-mig-subnet-01 | vsw-mj77bwqspdjcxy9d4i88j | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my04-mig-sshkey-01 | tb5aoempt96t422gdonc |  | d467b5e0b43d2054cf8965e2708fbea3 |

### Security Groups

#### Security Group: my04-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my04-mig-sg-01 |
| **CSP Security Group ID** | sg-mj7gjbh1a4atjhkrj8wt |
| **VNet** | my04-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my04-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my04-mig-sg-02 |
| **CSP Security Group ID** | sg-mj7c2uilk7pplm0lr4kp |
| **VNet** | my04-mig-vnet-01 |
| **Rule Count** | 4 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.3342 |
| **Per Day** | $8.02 |
| **Per Month (30 days)** | $240.62 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| ALIBABA | ap-northeast-2 | 3 | $0.3342 | $240.62 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my04-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | ecs.e-c1m1.large | $0.0178 | $12.82 |
| my04-ng-influxdb-back-1 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |
| my04-ng-influxdb-back-2 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |




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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my04-infra101)",
  "success": true
}
```

