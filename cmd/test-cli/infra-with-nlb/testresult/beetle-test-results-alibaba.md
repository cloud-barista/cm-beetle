# CM-Beetle test results for ALIBABA (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with ALIBABA cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: 3950cc5
- imdl: unknown
- CB-Tumblebug: Unknown (Fallback to Latest)
- CB-Spider: Unknown (Fallback to Latest)
- CB-MapUI: Unknown (Fallback to Latest)
- Target CSP: ALIBABA
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: July 1, 2026
- Test Time: 15:26:33 KST
- Test Execution: 2026-07-01 15:26:33 KST

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

## Test result for ALIBABA

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 4.061s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m5.63s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 19ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 10ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 17ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 1.848s | Pass |
| 7 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 13.548s | Pass |
| 8 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 11ms | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 6ms | Pass |
| 10 | NLB Load Balancing Verification | ✅ **PASS** | 1m28.112s | Pass |
| 11 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.412s | Pass |
| 12 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.216s | Pass |
| 13 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 15.768s | Pass |
| 14 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 26.68s | Pass |

**Overall Result**: 14/14 tests passed ✅

**Total Duration**: 4m31.674650367s

*Test executed on July 1, 2026 at 15:26:33 KST (2026-07-01 15:26:33 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=60.0% Max=100.0% Avg=86.7% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "nlbBackend": "influxdb_back",
              "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
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
              "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "cloud_essd_entry",
            "rootDiskSize": 40,
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
          "uid": "tbbghi471c985rbqe33v",
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
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "uid": "tbai41dc0rp9b7gpcck5",
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
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "regionList": [
            "ap-northeast-1",
            "ap-northeast-2",
            "ap-southeast-1",
            "ap-southeast-3",
            "ap-southeast-5",
            "ap-southeast-6",
            "ap-southeast-7",
            "cn-beijing",
            "cn-chengdu",
            "cn-fuzhou",
            "cn-guangzhou",
            "cn-hangzhou",
            "cn-heyuan",
            "cn-hongkong",
            "cn-nanjing",
            "cn-qingdao",
            "cn-shanghai",
            "cn-shenzhen",
            "cn-wuhan-lr",
            "cn-wulanchabu",
            "cn-zhangjiakou",
            "eu-central-1",
            "eu-west-1",
            "me-central-1",
            "me-east-1",
            "na-south-1",
            "us-east-1",
            "us-west-1"
          ],
          "id": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "uid": "tb15cl4rjlgtva8v38op",
          "name": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.06.08 09:13:03 Mon",
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
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
              "value": "Kernel version is 5.15.0-179-generic, 2026.5.29"
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
              "value": "v2026.5.29"
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
              "value": "2026-05-29T07:00:58Z"
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
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
          "description": "Kernel version is 5.15.0-179-generic, 2026.5.29",
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
      ]
    },
    {
      "status": "partially-matched",
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=60.0% Max=100.0% Avg=86.7% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "nlbBackend": "influxdb_back",
              "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
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
              "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
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
        "postCommand": {
          "userName": "",
          "command": null
        },
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
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "uid": "tbj7kinu2tgei794cq0i",
          "cspSpecName": "ecs.t6-c1m4.xlarge",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
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
          "costPerHour": 0.16926,
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
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "uid": "tbc3l7erbrt3lpp0obgq",
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
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "regionList": [
            "ap-northeast-1",
            "ap-northeast-2",
            "ap-southeast-1",
            "ap-southeast-3",
            "ap-southeast-5",
            "ap-southeast-6",
            "ap-southeast-7",
            "cn-beijing",
            "cn-chengdu",
            "cn-fuzhou",
            "cn-guangzhou",
            "cn-hangzhou",
            "cn-heyuan",
            "cn-hongkong",
            "cn-nanjing",
            "cn-qingdao",
            "cn-shanghai",
            "cn-shenzhen",
            "cn-wuhan-lr",
            "cn-wulanchabu",
            "cn-zhangjiakou",
            "eu-central-1",
            "eu-west-1",
            "me-central-1",
            "me-east-1",
            "na-south-1",
            "us-east-1",
            "us-west-1"
          ],
          "id": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "uid": "tb15cl4rjlgtva8v38op",
          "name": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.06.08 09:13:03 Mon",
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
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
              "value": "Kernel version is 5.15.0-179-generic, 2026.5.29"
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
              "value": "v2026.5.29"
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
              "value": "2026-05-29T07:00:58Z"
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
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
          "description": "Kernel version is 5.15.0-179-generic, 2026.5.29",
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
  "uid": "tbpkg3ju83mv2cgvh506",
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
    "sys.uid": "tbpkg3ju83mv2cgvh506"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbv94e1717k2psbrl8vj",
      "cspResourceName": "tbv94e1717k2psbrl8vj",
      "cspResourceId": "i-mj72v2gcuu4t9s3674t1",
      "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-07-01 06:27:35",
      "label": {
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-07-01 06:27:35",
        "sys.cspResourceId": "i-mj72v2gcuu4t9s3674t1",
        "sys.cspResourceName": "tbv94e1717k2psbrl8vj",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbv94e1717k2psbrl8vj",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "47.80.7.206",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.202",
      "privateDNS": "",
      "rootDiskType": "cloud_essd_entry",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-mj7bs5xicpvmokn3j8rx6",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "vsw-mj725y065dz05gozxllw0",
      "networkInterface": "eni-mj72v2gcuu4t9s3424nw",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbbj6mafh101dhopefl6",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "b1eA$kubcth!bt",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBLNqm+Msu+iliZXZkAKO3/F2WYFmtfUNY+5gIp7aPgiwkHhvp8stPxMK1vc9hlE3dNBUt7SYQ6/xuqFHfzaaBEg=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:ECIEfWtxe4k0u9UtOVx8yCxfPsdc9E9xjl6rn6Jz4WM",
        "firstUsedAt": "2026-07-01T06:27:43Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:27:43Z",
          "completedTime": "2026-07-01T06:27:44Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj72v2gcuu4t9s3674t1Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260615.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m1.large"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "tbv94e1717k2psbrl8vj"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "200"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj72v2gcuu4t9s3674t1Z"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "39f38f96-7573-4f69-8a9e-55f8b14cf7bd"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj72v2gcuu4t9s3674t1"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "2048"
        },
        {
          "key": "CreationTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "KeyPairName",
          "value": "tbbj6mafh101dhopefl6"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj7c5h3pubp9gdjeyqfa]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[47.80.7.206]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:ON,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj725y065dz05gozxllw0,VpcId:vpc-mj7bs5xicpvmokn3j8rx6,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.202]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:04:9e:0c,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj72v2gcuu4t9s3424nw,PrimaryIpAddress:10.0.1.202,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.202,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tb2fq8ommbbfh1fargco",
      "cspResourceName": "tb2fq8ommbbfh1fargco",
      "cspResourceId": "i-mj7el4jrcshrdb07v7y5",
      "name": "my-ng-influxdb-back-1",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-01 06:27:35",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-07-01 06:27:35",
        "sys.cspResourceId": "i-mj7el4jrcshrdb07v7y5",
        "sys.cspResourceName": "tb2fq8ommbbfh1fargco",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb2fq8ommbbfh1fargco",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "8.213.149.169",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.201",
      "privateDNS": "",
      "rootDiskType": "cloud_auto",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-mj7bs5xicpvmokn3j8rx6",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "vsw-mj725y065dz05gozxllw0",
      "networkInterface": "eni-mj7el4jrcshrdb08jgp9",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbbj6mafh101dhopefl6",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "1k71o$An!tbvoj",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAKioTZZw/2dfI/QdXjWUQleWeAUzavpZ3kZUZHs0eQhQ2COYCsbhcIIGI9TAEEQYmtYZaOS48Duq89MrxcpLC4=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:AyF5fPdElYdcvkifZ2mFnHNiVVUBEM8gtGIPXfJDyig",
        "firstUsedAt": "2026-07-01T06:27:44Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:27:43Z",
          "completedTime": "2026-07-01T06:27:45Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7el4jrcshrdb07v7y5Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260615.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m4.xlarge"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "tb2fq8ommbbfh1fargco"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "800"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj7el4jrcshrdb07v7y5Z"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "4"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "babc67b5-b7b7-4cf4-ad05-c4e958fb481f"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj7el4jrcshrdb07v7y5"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "16384"
        },
        {
          "key": "CreationTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "KeyPairName",
          "value": "tbbj6mafh101dhopefl6"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj7gfnq2bzbgd0f3ief0]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.213.149.169]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:2,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj725y065dz05gozxllw0,VpcId:vpc-mj7bs5xicpvmokn3j8rx6,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.201]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:04:9e:0b,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7el4jrcshrdb08jgp9,PrimaryIpAddress:10.0.1.201,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.201,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tbokbh6onufboogkr50d",
      "cspResourceName": "tbokbh6onufboogkr50d",
      "cspResourceId": "i-mj72g1bx55n3xahf3n2t",
      "name": "my-ng-influxdb-back-2",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-01 06:27:35",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-07-01 06:27:35",
        "sys.cspResourceId": "i-mj72g1bx55n3xahf3n2t",
        "sys.cspResourceName": "tbokbh6onufboogkr50d",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbokbh6onufboogkr50d",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "47.80.12.1",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.203",
      "privateDNS": "",
      "rootDiskType": "cloud_auto",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-mj7bs5xicpvmokn3j8rx6",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "vsw-mj725y065dz05gozxllw0",
      "networkInterface": "eni-mj72g1bx55n3xahb0sya",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbbj6mafh101dhopefl6",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "6speqA1$g8bt!s",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMRxYuuuLxlJ8Eh65OCBqIqfs9NkiJS6GtEj8f6kR/QF+/NVaS82GJwRDNSEyuCIZ5o7730JARezdspPRNciISs=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:iGkWYzvDV9D8ZbPzuS39rKieaSLfz73YanlXEc+fkXI",
        "firstUsedAt": "2026-07-01T06:27:44Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:27:43Z",
          "completedTime": "2026-07-01T06:27:45Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj72g1bx55n3xahf3n2tZ 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260615.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m4.xlarge"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "tbokbh6onufboogkr50d"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "800"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj72g1bx55n3xahf3n2tZ"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "4"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "cc9158e2-2681-4c37-b0ac-cac17d1da9aa"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj72g1bx55n3xahf3n2t"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "16384"
        },
        {
          "key": "CreationTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "KeyPairName",
          "value": "tbbj6mafh101dhopefl6"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj7gfnq2bzbgd0f3ief0]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[47.80.12.1]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:2,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj725y065dz05gozxllw0,VpcId:vpc-mj7bs5xicpvmokn3j8rx6,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.203]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:04:9e:0d,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj72g1bx55n3xahb0sya,PrimaryIpAddress:10.0.1.203,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.203,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
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
        "nodeIp": "47.80.7.206",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj72v2gcuu4t9s3674t1Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "47.80.12.1",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj72g1bx55n3xahf3n2tZ 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "8.213.149.169",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj7el4jrcshrdb07v7y5Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "tbpkg3ju83mv2cgvh506",
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
        "sys.uid": "tbpkg3ju83mv2cgvh506"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbv94e1717k2psbrl8vj",
          "cspResourceName": "tbv94e1717k2psbrl8vj",
          "cspResourceId": "i-mj72v2gcuu4t9s3674t1",
          "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-07-01 06:27:35",
          "label": {
            "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-07-01 06:27:35",
            "sys.cspResourceId": "i-mj72v2gcuu4t9s3674t1",
            "sys.cspResourceName": "tbv94e1717k2psbrl8vj",
            "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbv94e1717k2psbrl8vj",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.7.206",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.202",
          "privateDNS": "",
          "rootDiskType": "cloud_essd_entry",
          "rootDiskSize": 40,
          "RootDeviceName": "/dev/xvda",
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
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "vpc-mj7bs5xicpvmokn3j8rx6",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "vsw-mj725y065dz05gozxllw0",
          "networkInterface": "eni-mj72v2gcuu4t9s3424nw",
          "securityGroupIds": [
            "my-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tbbj6mafh101dhopefl6",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "b1eA$kubcth!bt",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBLNqm+Msu+iliZXZkAKO3/F2WYFmtfUNY+5gIp7aPgiwkHhvp8stPxMK1vc9hlE3dNBUt7SYQ6/xuqFHfzaaBEg=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:ECIEfWtxe4k0u9UtOVx8yCxfPsdc9E9xjl6rn6Jz4WM",
            "firstUsedAt": "2026-07-01T06:27:43Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-01T06:27:43Z",
              "completedTime": "2026-07-01T06:27:44Z",
              "elapsedTime": 1,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj72v2gcuu4t9s3674t1Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260615.vhd"
            },
            {
              "key": "InstanceType",
              "value": "ecs.e-c1m1.large"
            },
            {
              "key": "DeviceAvailable",
              "value": "true"
            },
            {
              "key": "InstanceNetworkType",
              "value": "vpc"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "IsSpot",
              "value": "false"
            },
            {
              "key": "InstanceChargeType",
              "value": "PostPaid"
            },
            {
              "key": "InstanceName",
              "value": "tbv94e1717k2psbrl8vj"
            },
            {
              "key": "DeploymentSetGroupNo",
              "value": "0"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "Connected",
              "value": "false"
            },
            {
              "key": "InvocationCount",
              "value": "0"
            },
            {
              "key": "StartTime",
              "value": "2026-07-01T06:27Z"
            },
            {
              "key": "ZoneId",
              "value": "ap-northeast-2a"
            },
            {
              "key": "InternetMaxBandwidthIn",
              "value": "200"
            },
            {
              "key": "InternetChargeType",
              "value": "PayByBandwidth"
            },
            {
              "key": "HostName",
              "value": "iZmj72v2gcuu4t9s3674t1Z"
            },
            {
              "key": "Status",
              "value": "Running"
            },
            {
              "key": "CPU",
              "value": "0"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "SpotPriceLimit",
              "value": "0.00"
            },
            {
              "key": "OSName",
              "value": "Ubuntu  22.04 64位"
            },
            {
              "key": "InstanceOwnerId",
              "value": "0"
            },
            {
              "key": "OSNameEn",
              "value": "Ubuntu  22.04 64 bit"
            },
            {
              "key": "SerialNumber",
              "value": "39f38f96-7573-4f69-8a9e-55f8b14cf7bd"
            },
            {
              "key": "RegionId",
              "value": "ap-northeast-2"
            },
            {
              "key": "IoOptimized",
              "value": "true"
            },
            {
              "key": "InternetMaxBandwidthOut",
              "value": "5"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.e"
            },
            {
              "key": "InstanceId",
              "value": "i-mj72v2gcuu4t9s3674t1"
            },
            {
              "key": "Recyclable",
              "value": "false"
            },
            {
              "key": "ExpiredTime",
              "value": "2099-12-31T15:59Z"
            },
            {
              "key": "OSType",
              "value": "linux"
            },
            {
              "key": "Memory",
              "value": "2048"
            },
            {
              "key": "CreationTime",
              "value": "2026-07-01T06:27Z"
            },
            {
              "key": "KeyPairName",
              "value": "tbbj6mafh101dhopefl6"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "StoppedMode",
              "value": "Not-applicable"
            },
            {
              "key": "SpotStrategy",
              "value": "NoSpot"
            },
            {
              "key": "SpotDuration",
              "value": "0"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "SecurityGroupIds",
              "value": "{SecurityGroupId:[sg-mj7c5h3pubp9gdjeyqfa]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[47.80.7.206]}"
            },
            {
              "key": "RdmaIpAddress",
              "value": "{IpAddress:null}"
            },
            {
              "key": "DedicatedHostAttribute",
              "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
            },
            {
              "key": "EcsCapacityReservationAttr",
              "value": "{CapacityReservationPreference:,CapacityReservationId:}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:ON,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "DedicatedInstanceAttribute",
              "value": "{Affinity:,Tenancy:}"
            },
            {
              "key": "PrivateDnsNameOptions",
              "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
            },
            {
              "key": "AdditionalInfo",
              "value": "{EnableHighDensityMode:false}"
            },
            {
              "key": "ImageOptions",
              "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
            },
            {
              "key": "EipAddress",
              "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
            },
            {
              "key": "VpcAttributes",
              "value": "{VSwitchId:vsw-mj725y065dz05gozxllw0,VpcId:vpc-mj7bs5xicpvmokn3j8rx6,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.202]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:04:9e:0c,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj72v2gcuu4t9s3424nw,PrimaryIpAddress:10.0.1.202,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.202,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
            },
            {
              "key": "OperationLocks",
              "value": "{LockReason:[]}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-1",
          "uid": "tb2fq8ommbbfh1fargco",
          "cspResourceName": "tb2fq8ommbbfh1fargco",
          "cspResourceId": "i-mj7el4jrcshrdb07v7y5",
          "name": "my-ng-influxdb-back-1",
          "nodeGroupId": "my-ng-influxdb-back",
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
          "createdTime": "2026-07-01 06:27:35",
          "label": {
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-07-01 06:27:35",
            "sys.cspResourceId": "i-mj7el4jrcshrdb07v7y5",
            "sys.cspResourceName": "tb2fq8ommbbfh1fargco",
            "sys.id": "my-ng-influxdb-back-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tb2fq8ommbbfh1fargco",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.213.149.169",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.201",
          "privateDNS": "",
          "rootDiskType": "cloud_auto",
          "rootDiskSize": 40,
          "RootDeviceName": "/dev/xvda",
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
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "vpc-mj7bs5xicpvmokn3j8rx6",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "vsw-mj725y065dz05gozxllw0",
          "networkInterface": "eni-mj7el4jrcshrdb08jgp9",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tbbj6mafh101dhopefl6",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "1k71o$An!tbvoj",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAKioTZZw/2dfI/QdXjWUQleWeAUzavpZ3kZUZHs0eQhQ2COYCsbhcIIGI9TAEEQYmtYZaOS48Duq89MrxcpLC4=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:AyF5fPdElYdcvkifZ2mFnHNiVVUBEM8gtGIPXfJDyig",
            "firstUsedAt": "2026-07-01T06:27:44Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-01T06:27:43Z",
              "completedTime": "2026-07-01T06:27:45Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj7el4jrcshrdb07v7y5Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260615.vhd"
            },
            {
              "key": "InstanceType",
              "value": "ecs.e-c1m4.xlarge"
            },
            {
              "key": "DeviceAvailable",
              "value": "true"
            },
            {
              "key": "InstanceNetworkType",
              "value": "vpc"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "IsSpot",
              "value": "false"
            },
            {
              "key": "InstanceChargeType",
              "value": "PostPaid"
            },
            {
              "key": "InstanceName",
              "value": "tb2fq8ommbbfh1fargco"
            },
            {
              "key": "DeploymentSetGroupNo",
              "value": "0"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "Connected",
              "value": "false"
            },
            {
              "key": "InvocationCount",
              "value": "0"
            },
            {
              "key": "StartTime",
              "value": "2026-07-01T06:27Z"
            },
            {
              "key": "ZoneId",
              "value": "ap-northeast-2a"
            },
            {
              "key": "InternetMaxBandwidthIn",
              "value": "800"
            },
            {
              "key": "InternetChargeType",
              "value": "PayByBandwidth"
            },
            {
              "key": "HostName",
              "value": "iZmj7el4jrcshrdb07v7y5Z"
            },
            {
              "key": "Status",
              "value": "Running"
            },
            {
              "key": "CPU",
              "value": "0"
            },
            {
              "key": "Cpu",
              "value": "4"
            },
            {
              "key": "SpotPriceLimit",
              "value": "0.00"
            },
            {
              "key": "OSName",
              "value": "Ubuntu  22.04 64位"
            },
            {
              "key": "InstanceOwnerId",
              "value": "0"
            },
            {
              "key": "OSNameEn",
              "value": "Ubuntu  22.04 64 bit"
            },
            {
              "key": "SerialNumber",
              "value": "babc67b5-b7b7-4cf4-ad05-c4e958fb481f"
            },
            {
              "key": "RegionId",
              "value": "ap-northeast-2"
            },
            {
              "key": "IoOptimized",
              "value": "true"
            },
            {
              "key": "InternetMaxBandwidthOut",
              "value": "5"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.e"
            },
            {
              "key": "InstanceId",
              "value": "i-mj7el4jrcshrdb07v7y5"
            },
            {
              "key": "Recyclable",
              "value": "false"
            },
            {
              "key": "ExpiredTime",
              "value": "2099-12-31T15:59Z"
            },
            {
              "key": "OSType",
              "value": "linux"
            },
            {
              "key": "Memory",
              "value": "16384"
            },
            {
              "key": "CreationTime",
              "value": "2026-07-01T06:27Z"
            },
            {
              "key": "KeyPairName",
              "value": "tbbj6mafh101dhopefl6"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "StoppedMode",
              "value": "Not-applicable"
            },
            {
              "key": "SpotStrategy",
              "value": "NoSpot"
            },
            {
              "key": "SpotDuration",
              "value": "0"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "SecurityGroupIds",
              "value": "{SecurityGroupId:[sg-mj7gfnq2bzbgd0f3ief0]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[8.213.149.169]}"
            },
            {
              "key": "RdmaIpAddress",
              "value": "{IpAddress:null}"
            },
            {
              "key": "DedicatedHostAttribute",
              "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
            },
            {
              "key": "EcsCapacityReservationAttr",
              "value": "{CapacityReservationPreference:,CapacityReservationId:}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:2,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "DedicatedInstanceAttribute",
              "value": "{Affinity:,Tenancy:}"
            },
            {
              "key": "PrivateDnsNameOptions",
              "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
            },
            {
              "key": "AdditionalInfo",
              "value": "{EnableHighDensityMode:false}"
            },
            {
              "key": "ImageOptions",
              "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
            },
            {
              "key": "EipAddress",
              "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
            },
            {
              "key": "VpcAttributes",
              "value": "{VSwitchId:vsw-mj725y065dz05gozxllw0,VpcId:vpc-mj7bs5xicpvmokn3j8rx6,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.201]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:04:9e:0b,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7el4jrcshrdb08jgp9,PrimaryIpAddress:10.0.1.201,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.201,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
            },
            {
              "key": "OperationLocks",
              "value": "{LockReason:[]}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-2",
          "uid": "tbokbh6onufboogkr50d",
          "cspResourceName": "tbokbh6onufboogkr50d",
          "cspResourceId": "i-mj72g1bx55n3xahf3n2t",
          "name": "my-ng-influxdb-back-2",
          "nodeGroupId": "my-ng-influxdb-back",
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
          "createdTime": "2026-07-01 06:27:35",
          "label": {
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-07-01 06:27:35",
            "sys.cspResourceId": "i-mj72g1bx55n3xahf3n2t",
            "sys.cspResourceName": "tbokbh6onufboogkr50d",
            "sys.id": "my-ng-influxdb-back-2",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbokbh6onufboogkr50d",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.12.1",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.203",
          "privateDNS": "",
          "rootDiskType": "cloud_auto",
          "rootDiskSize": 40,
          "RootDeviceName": "/dev/xvda",
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
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "vpc-mj7bs5xicpvmokn3j8rx6",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "vsw-mj725y065dz05gozxllw0",
          "networkInterface": "eni-mj72g1bx55n3xahb0sya",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tbbj6mafh101dhopefl6",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "6speqA1$g8bt!s",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMRxYuuuLxlJ8Eh65OCBqIqfs9NkiJS6GtEj8f6kR/QF+/NVaS82GJwRDNSEyuCIZ5o7730JARezdspPRNciISs=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:iGkWYzvDV9D8ZbPzuS39rKieaSLfz73YanlXEc+fkXI",
            "firstUsedAt": "2026-07-01T06:27:44Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-01T06:27:43Z",
              "completedTime": "2026-07-01T06:27:45Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj72g1bx55n3xahf3n2tZ 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260615.vhd"
            },
            {
              "key": "InstanceType",
              "value": "ecs.e-c1m4.xlarge"
            },
            {
              "key": "DeviceAvailable",
              "value": "true"
            },
            {
              "key": "InstanceNetworkType",
              "value": "vpc"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "IsSpot",
              "value": "false"
            },
            {
              "key": "InstanceChargeType",
              "value": "PostPaid"
            },
            {
              "key": "InstanceName",
              "value": "tbokbh6onufboogkr50d"
            },
            {
              "key": "DeploymentSetGroupNo",
              "value": "0"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "Connected",
              "value": "false"
            },
            {
              "key": "InvocationCount",
              "value": "0"
            },
            {
              "key": "StartTime",
              "value": "2026-07-01T06:27Z"
            },
            {
              "key": "ZoneId",
              "value": "ap-northeast-2a"
            },
            {
              "key": "InternetMaxBandwidthIn",
              "value": "800"
            },
            {
              "key": "InternetChargeType",
              "value": "PayByBandwidth"
            },
            {
              "key": "HostName",
              "value": "iZmj72g1bx55n3xahf3n2tZ"
            },
            {
              "key": "Status",
              "value": "Running"
            },
            {
              "key": "CPU",
              "value": "0"
            },
            {
              "key": "Cpu",
              "value": "4"
            },
            {
              "key": "SpotPriceLimit",
              "value": "0.00"
            },
            {
              "key": "OSName",
              "value": "Ubuntu  22.04 64位"
            },
            {
              "key": "InstanceOwnerId",
              "value": "0"
            },
            {
              "key": "OSNameEn",
              "value": "Ubuntu  22.04 64 bit"
            },
            {
              "key": "SerialNumber",
              "value": "cc9158e2-2681-4c37-b0ac-cac17d1da9aa"
            },
            {
              "key": "RegionId",
              "value": "ap-northeast-2"
            },
            {
              "key": "IoOptimized",
              "value": "true"
            },
            {
              "key": "InternetMaxBandwidthOut",
              "value": "5"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.e"
            },
            {
              "key": "InstanceId",
              "value": "i-mj72g1bx55n3xahf3n2t"
            },
            {
              "key": "Recyclable",
              "value": "false"
            },
            {
              "key": "ExpiredTime",
              "value": "2099-12-31T15:59Z"
            },
            {
              "key": "OSType",
              "value": "linux"
            },
            {
              "key": "Memory",
              "value": "16384"
            },
            {
              "key": "CreationTime",
              "value": "2026-07-01T06:27Z"
            },
            {
              "key": "KeyPairName",
              "value": "tbbj6mafh101dhopefl6"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "StoppedMode",
              "value": "Not-applicable"
            },
            {
              "key": "SpotStrategy",
              "value": "NoSpot"
            },
            {
              "key": "SpotDuration",
              "value": "0"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "SecurityGroupIds",
              "value": "{SecurityGroupId:[sg-mj7gfnq2bzbgd0f3ief0]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[47.80.12.1]}"
            },
            {
              "key": "RdmaIpAddress",
              "value": "{IpAddress:null}"
            },
            {
              "key": "DedicatedHostAttribute",
              "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
            },
            {
              "key": "EcsCapacityReservationAttr",
              "value": "{CapacityReservationPreference:,CapacityReservationId:}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:2,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "DedicatedInstanceAttribute",
              "value": "{Affinity:,Tenancy:}"
            },
            {
              "key": "PrivateDnsNameOptions",
              "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
            },
            {
              "key": "AdditionalInfo",
              "value": "{EnableHighDensityMode:false}"
            },
            {
              "key": "ImageOptions",
              "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
            },
            {
              "key": "EipAddress",
              "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
            },
            {
              "key": "VpcAttributes",
              "value": "{VSwitchId:vsw-mj725y065dz05gozxllw0,VpcId:vpc-mj7bs5xicpvmokn3j8rx6,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.203]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:04:9e:0d,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj72g1bx55n3xahb0sya,PrimaryIpAddress:10.0.1.203,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.203,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
            },
            {
              "key": "OperationLocks",
              "value": "{LockReason:[]}"
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
            "nodeIp": "47.80.7.206",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj72v2gcuu4t9s3674t1Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-2",
            "nodeIp": "47.80.12.1",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj72g1bx55n3xahf3n2tZ 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-1",
            "nodeIp": "8.213.149.169",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7el4jrcshrdb07v7y5Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "tbpkg3ju83mv2cgvh506",
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
    "sys.uid": "tbpkg3ju83mv2cgvh506"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbv94e1717k2psbrl8vj",
      "cspResourceName": "tbv94e1717k2psbrl8vj",
      "cspResourceId": "i-mj72v2gcuu4t9s3674t1",
      "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-07-01 06:27:35",
      "label": {
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-07-01 06:27:35",
        "sys.cspResourceId": "i-mj72v2gcuu4t9s3674t1",
        "sys.cspResourceName": "tbv94e1717k2psbrl8vj",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbv94e1717k2psbrl8vj",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "47.80.7.206",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.202",
      "privateDNS": "",
      "rootDiskType": "cloud_essd_entry",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-mj7bs5xicpvmokn3j8rx6",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "vsw-mj725y065dz05gozxllw0",
      "networkInterface": "eni-mj72v2gcuu4t9s3424nw",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbbj6mafh101dhopefl6",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "b1eA$kubcth!bt",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBLNqm+Msu+iliZXZkAKO3/F2WYFmtfUNY+5gIp7aPgiwkHhvp8stPxMK1vc9hlE3dNBUt7SYQ6/xuqFHfzaaBEg=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:ECIEfWtxe4k0u9UtOVx8yCxfPsdc9E9xjl6rn6Jz4WM",
        "firstUsedAt": "2026-07-01T06:27:43Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:27:43Z",
          "completedTime": "2026-07-01T06:27:44Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj72v2gcuu4t9s3674t1Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260615.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m1.large"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "tbv94e1717k2psbrl8vj"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "200"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj72v2gcuu4t9s3674t1Z"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "39f38f96-7573-4f69-8a9e-55f8b14cf7bd"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj72v2gcuu4t9s3674t1"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "2048"
        },
        {
          "key": "CreationTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "KeyPairName",
          "value": "tbbj6mafh101dhopefl6"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj7c5h3pubp9gdjeyqfa]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[47.80.7.206]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:ON,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj725y065dz05gozxllw0,VpcId:vpc-mj7bs5xicpvmokn3j8rx6,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.202]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:04:9e:0c,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj72v2gcuu4t9s3424nw,PrimaryIpAddress:10.0.1.202,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.202,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tb2fq8ommbbfh1fargco",
      "cspResourceName": "tb2fq8ommbbfh1fargco",
      "cspResourceId": "i-mj7el4jrcshrdb07v7y5",
      "name": "my-ng-influxdb-back-1",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-01 06:27:35",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-07-01 06:27:35",
        "sys.cspResourceId": "i-mj7el4jrcshrdb07v7y5",
        "sys.cspResourceName": "tb2fq8ommbbfh1fargco",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb2fq8ommbbfh1fargco",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "8.213.149.169",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.201",
      "privateDNS": "",
      "rootDiskType": "cloud_auto",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-mj7bs5xicpvmokn3j8rx6",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "vsw-mj725y065dz05gozxllw0",
      "networkInterface": "eni-mj7el4jrcshrdb08jgp9",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbbj6mafh101dhopefl6",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "1k71o$An!tbvoj",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAKioTZZw/2dfI/QdXjWUQleWeAUzavpZ3kZUZHs0eQhQ2COYCsbhcIIGI9TAEEQYmtYZaOS48Duq89MrxcpLC4=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:AyF5fPdElYdcvkifZ2mFnHNiVVUBEM8gtGIPXfJDyig",
        "firstUsedAt": "2026-07-01T06:27:44Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:27:43Z",
          "completedTime": "2026-07-01T06:27:45Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7el4jrcshrdb07v7y5Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260615.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m4.xlarge"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "tb2fq8ommbbfh1fargco"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "800"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj7el4jrcshrdb07v7y5Z"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "4"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "babc67b5-b7b7-4cf4-ad05-c4e958fb481f"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj7el4jrcshrdb07v7y5"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "16384"
        },
        {
          "key": "CreationTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "KeyPairName",
          "value": "tbbj6mafh101dhopefl6"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj7gfnq2bzbgd0f3ief0]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.213.149.169]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:2,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj725y065dz05gozxllw0,VpcId:vpc-mj7bs5xicpvmokn3j8rx6,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.201]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:04:9e:0b,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7el4jrcshrdb08jgp9,PrimaryIpAddress:10.0.1.201,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.201,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tbokbh6onufboogkr50d",
      "cspResourceName": "tbokbh6onufboogkr50d",
      "cspResourceId": "i-mj72g1bx55n3xahf3n2t",
      "name": "my-ng-influxdb-back-2",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-01 06:27:35",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-07-01 06:27:35",
        "sys.cspResourceId": "i-mj72g1bx55n3xahf3n2t",
        "sys.cspResourceName": "tbokbh6onufboogkr50d",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbokbh6onufboogkr50d",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=60.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "47.80.12.1",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.203",
      "privateDNS": "",
      "rootDiskType": "cloud_auto",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260615.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-mj7bs5xicpvmokn3j8rx6",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "vsw-mj725y065dz05gozxllw0",
      "networkInterface": "eni-mj72g1bx55n3xahb0sya",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbbj6mafh101dhopefl6",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "6speqA1$g8bt!s",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMRxYuuuLxlJ8Eh65OCBqIqfs9NkiJS6GtEj8f6kR/QF+/NVaS82GJwRDNSEyuCIZ5o7730JARezdspPRNciISs=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:iGkWYzvDV9D8ZbPzuS39rKieaSLfz73YanlXEc+fkXI",
        "firstUsedAt": "2026-07-01T06:27:44Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-01T06:27:43Z",
          "completedTime": "2026-07-01T06:27:45Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj72g1bx55n3xahf3n2tZ 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260615.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m4.xlarge"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "tbokbh6onufboogkr50d"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "800"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj72g1bx55n3xahf3n2tZ"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "4"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "cc9158e2-2681-4c37-b0ac-cac17d1da9aa"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj72g1bx55n3xahf3n2t"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "16384"
        },
        {
          "key": "CreationTime",
          "value": "2026-07-01T06:27Z"
        },
        {
          "key": "KeyPairName",
          "value": "tbbj6mafh101dhopefl6"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj7gfnq2bzbgd0f3ief0]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[47.80.12.1]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:2,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj725y065dz05gozxllw0,VpcId:vpc-mj7bs5xicpvmokn3j8rx6,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.203]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:04:9e:0d,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj72g1bx55n3xahb0sya,PrimaryIpAddress:10.0.1.203,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.203,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
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
        "nodeIp": "47.80.7.206",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj72v2gcuu4t9s3674t1Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "47.80.12.1",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj72g1bx55n3xahf3n2tZ 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "8.213.149.169",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj7el4jrcshrdb07v7y5Z 5.15.0-181-generic #191-Ubuntu SMP Fri May 22 19:09:02 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "cspResourceName": "tbiqgtsc58od6omfq21d",
      "cspResourceId": "lb-mj7shdy89ri41caxe3xq6",
      "name": "my-ng-influxdb-back",
      "connectionName": "alibaba-ap-northeast-2",
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
        "ip": "43.108.15.106",
        "port": "9999"
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my-ng-influxdb-back",
        "nodes": [
          "my-ng-influxdb-back-1",
          "my-ng-influxdb-back-2"
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
            "value": "{NameId:tbiqgtsc58od6omfq21d,SystemId:lb-mj7shdy89ri41caxe3xq6}"
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
            "value": "{Key:BaseResponse,Value:{}}; {Key:Status,Value:running}; {Key:ConnectionDrainTimeout,Value:0}; {Key:HealthCheckTcpFastCloseEnabled,Value:false}; {Key:FullNatEnabled,Value:false}; {Key:RequestId,Value:8E51D1A7-C835-33FC-AA72-31D91F04A8AC}; {Key:HealthCheckConnectPort,Value:8086}; {Key:Bandwidth,Value:-1}; {Key:HealthCheckType,Value:tcp}; {Key:BackendServerPort,Value:8086}; {Key:AclStatus,Value:off}; {Key:UnhealthyThreshold,Value:3}; {Key:MaxConnection,Value:0}; {Key:ProxyProtocolV2Enabled,Value:false}; {Key:PersistenceTimeout,Value:0}; {Key:ListenerPort,Value:9999}; {Key:HealthCheckInterval,Value:10}; {Key:FailoverThreshold,Value:0}; {Key:Scheduler,Value:wrr}; {Key:EstablishedTimeout,Value:900}; {Key:HealthCheckConnectTimeout,Value:10}; {Key:MasterSlaveModeEnabled,Value:false}; {Key:HealthyThreshold,Value:3}; {Key:HealthCheck,Value:on}; {Key:LoadBalancerId,Value:lb-mj7shdy89ri41caxe3xq6}; {Key:AclIds,Value:{AclId:null}}; {Key:PortRanges,Value:{PortRange:[]}}; {Key:Tags,Value:{Tag:[]}}"
          }
        ]
      },
      "createdTime": "2026-07-01T06:28:18Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "BaseResponse",
          "value": "{}"
        },
        {
          "key": "CreateTimeStamp",
          "value": "1782887298000"
        },
        {
          "key": "HasReservedInfo",
          "value": "false"
        },
        {
          "key": "CreateTime",
          "value": "2026-07-01T06:28:18Z"
        },
        {
          "key": "LoadBalancerId",
          "value": "lb-mj7shdy89ri41caxe3xq6"
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
          "value": "09E13231-5DAB-3B2E-996F-329C1B88F06C"
        },
        {
          "key": "Bandwidth",
          "value": "5120"
        },
        {
          "key": "LoadBalancerName",
          "value": "tbiqgtsc58od6omfq21d"
        },
        {
          "key": "Address",
          "value": "43.108.15.106"
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
          "value": "{BackendServer:[{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7el4jrcshrdb07v7y5},{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj72g1bx55n3xahf3n2t}]}"
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
    "cspResourceName": "tbiqgtsc58od6omfq21d",
    "cspResourceId": "lb-mj7shdy89ri41caxe3xq6",
    "name": "my-ng-influxdb-back",
    "connectionName": "alibaba-ap-northeast-2",
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
      "ip": "43.108.15.106",
      "port": "9999"
    },
    "targetGroup": {
      "protocol": "TCP",
      "port": "8086",
      "nodeGroupId": "my-ng-influxdb-back",
      "nodes": [
        "my-ng-influxdb-back-1",
        "my-ng-influxdb-back-2"
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
          "value": "{NameId:tbiqgtsc58od6omfq21d,SystemId:lb-mj7shdy89ri41caxe3xq6}"
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
          "value": "{Key:BaseResponse,Value:{}}; {Key:Status,Value:running}; {Key:ConnectionDrainTimeout,Value:0}; {Key:HealthCheckTcpFastCloseEnabled,Value:false}; {Key:FullNatEnabled,Value:false}; {Key:RequestId,Value:8E51D1A7-C835-33FC-AA72-31D91F04A8AC}; {Key:HealthCheckConnectPort,Value:8086}; {Key:Bandwidth,Value:-1}; {Key:HealthCheckType,Value:tcp}; {Key:BackendServerPort,Value:8086}; {Key:AclStatus,Value:off}; {Key:UnhealthyThreshold,Value:3}; {Key:MaxConnection,Value:0}; {Key:ProxyProtocolV2Enabled,Value:false}; {Key:PersistenceTimeout,Value:0}; {Key:ListenerPort,Value:9999}; {Key:HealthCheckInterval,Value:10}; {Key:FailoverThreshold,Value:0}; {Key:Scheduler,Value:wrr}; {Key:EstablishedTimeout,Value:900}; {Key:HealthCheckConnectTimeout,Value:10}; {Key:MasterSlaveModeEnabled,Value:false}; {Key:HealthyThreshold,Value:3}; {Key:HealthCheck,Value:on}; {Key:LoadBalancerId,Value:lb-mj7shdy89ri41caxe3xq6}; {Key:AclIds,Value:{AclId:null}}; {Key:PortRanges,Value:{PortRange:[]}}; {Key:Tags,Value:{Tag:[]}}"
        }
      ]
    },
    "createdTime": "2026-07-01T06:28:18Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "BaseResponse",
        "value": "{}"
      },
      {
        "key": "CreateTimeStamp",
        "value": "1782887298000"
      },
      {
        "key": "HasReservedInfo",
        "value": "false"
      },
      {
        "key": "CreateTime",
        "value": "2026-07-01T06:28:18Z"
      },
      {
        "key": "LoadBalancerId",
        "value": "lb-mj7shdy89ri41caxe3xq6"
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
        "value": "09E13231-5DAB-3B2E-996F-329C1B88F06C"
      },
      {
        "key": "Bandwidth",
        "value": "5120"
      },
      {
        "key": "LoadBalancerName",
        "value": "tbiqgtsc58od6omfq21d"
      },
      {
        "key": "Address",
        "value": "43.108.15.106"
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
        "value": "{BackendServer:[{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7el4jrcshrdb07v7y5},{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj72g1bx55n3xahf3n2t}]}"
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
  "cspResourceName": "tbiqgtsc58od6omfq21d",
  "cspResourceId": "lb-mj7shdy89ri41caxe3xq6",
  "name": "my-ng-influxdb-back",
  "connectionName": "alibaba-ap-northeast-2",
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
    "ip": "43.108.15.106",
    "port": "9999"
  },
  "targetGroup": {
    "protocol": "TCP",
    "port": "8086",
    "nodeGroupId": "my-ng-influxdb-back",
    "nodes": [
      "my-ng-influxdb-back-1",
      "my-ng-influxdb-back-2"
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
        "value": "{NameId:tbiqgtsc58od6omfq21d,SystemId:lb-mj7shdy89ri41caxe3xq6}"
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
        "value": "{Key:BaseResponse,Value:{}}; {Key:Status,Value:running}; {Key:ConnectionDrainTimeout,Value:0}; {Key:HealthCheckTcpFastCloseEnabled,Value:false}; {Key:FullNatEnabled,Value:false}; {Key:RequestId,Value:8E51D1A7-C835-33FC-AA72-31D91F04A8AC}; {Key:HealthCheckConnectPort,Value:8086}; {Key:Bandwidth,Value:-1}; {Key:HealthCheckType,Value:tcp}; {Key:BackendServerPort,Value:8086}; {Key:AclStatus,Value:off}; {Key:UnhealthyThreshold,Value:3}; {Key:MaxConnection,Value:0}; {Key:ProxyProtocolV2Enabled,Value:false}; {Key:PersistenceTimeout,Value:0}; {Key:ListenerPort,Value:9999}; {Key:HealthCheckInterval,Value:10}; {Key:FailoverThreshold,Value:0}; {Key:Scheduler,Value:wrr}; {Key:EstablishedTimeout,Value:900}; {Key:HealthCheckConnectTimeout,Value:10}; {Key:MasterSlaveModeEnabled,Value:false}; {Key:HealthyThreshold,Value:3}; {Key:HealthCheck,Value:on}; {Key:LoadBalancerId,Value:lb-mj7shdy89ri41caxe3xq6}; {Key:AclIds,Value:{AclId:null}}; {Key:PortRanges,Value:{PortRange:[]}}; {Key:Tags,Value:{Tag:[]}}"
      }
    ]
  },
  "createdTime": "2026-07-01T06:28:18Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "BaseResponse",
      "value": "{}"
    },
    {
      "key": "CreateTimeStamp",
      "value": "1782887298000"
    },
    {
      "key": "HasReservedInfo",
      "value": "false"
    },
    {
      "key": "CreateTime",
      "value": "2026-07-01T06:28:18Z"
    },
    {
      "key": "LoadBalancerId",
      "value": "lb-mj7shdy89ri41caxe3xq6"
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
      "value": "09E13231-5DAB-3B2E-996F-329C1B88F06C"
    },
    {
      "key": "Bandwidth",
      "value": "5120"
    },
    {
      "key": "LoadBalancerName",
      "value": "tbiqgtsc58od6omfq21d"
    },
    {
      "key": "Address",
      "value": "43.108.15.106"
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
      "value": "{BackendServer:[{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7el4jrcshrdb07v7y5},{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj72g1bx55n3xahf3n2t}]}"
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

**Generated At:** 2026-07-01 06:30:00

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
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
| ecs.e-c1m4.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1582 | 2 |
| ecs.e-c1m1.large | 2 | 2.0 | - | x86_64 |  | $0.0178 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ubuntu_22_04_x64_20G_alibase_20260522.vhd | Ubuntu  22.04 64 bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | i-mj72v2gcuu4t9s3674t1 | Running | 2 vCPU, 2.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 47.80.7.206<br>**Private IP:** 10.0.1.202<br>**SGs:** my-mig-sg-02<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-1 | i-mj7el4jrcshrdb07v7y5 | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 8.213.149.169<br>**Private IP:** 10.0.1.201<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-2 | i-mj72g1bx55n3xahf3n2t | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 47.80.12.1<br>**Private IP:** 10.0.1.203<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-vnet-01 |
| **CSP VNet ID** | vpc-mj7bs5xicpvmokn3j8rx6 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | alibaba-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-mig-subnet-01 | vsw-mj725y065dz05gozxllw0 | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-mig-sshkey-01 | tbbj6mafh101dhopefl6 |  | 938129357923abb15fa633e7158ce0ff |

### Security Groups

#### Security Group: my-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-01 |
| **CSP Security Group ID** | sg-mj7gfnq2bzbgd0f3ief0 |
| **VNet** | my-mig-vnet-01 |
| **Rule Count** | 5 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 8086 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-02 |
| **CSP Security Group ID** | sg-mj7c5h3pubp9gdjeyqfa |
| **VNet** | my-mig-vnet-01 |
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
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | ecs.e-c1m1.large | $0.0178 | $12.82 |
| my-ng-influxdb-back-1 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |
| my-ng-influxdb-back-2 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |




### Test Case 12: Migration Report

#### 12.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-07-01 06:30:05*

---

## 📊 Migration Summary

**Target Cloud:** ALIBABA

**Target Region:** ap-northeast-2

**Namespace:** mig01 | **Infra ID:** my-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $240.62 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | ecs.e-c1m1.large, ecs.e-c1m4.xlarge |
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
| 1 | **VM Name:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** i-mj72v2gcuu4t9s3674t1<br>**Label(sourceMachineId):** ng-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my-ng-influxdb-back-1<br>**VM ID:** i-mj7el4jrcshrdb07v7y5<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |
| 3 | **VM Name:** my-ng-influxdb-back-2<br>**VM ID:** i-mj72g1bx55n3xahf3n2t<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** ecs.e-c1m1.large<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2 | my-ng-influxdb-back-1 | **Spec ID:** ecs.e-c1m4.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3 | my-ng-influxdb-back-2 | **Spec ID:** ecs.e-c1m4.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260522.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my-ng-influxdb-back-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260522.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my-ng-influxdb-back-2 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260522.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 9 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my-mig-sg-01

**CSP ID:** sg-mj7gfnq2bzbgd0f3ief0 | **VNet:** my-mig-vnet-01 | **Rules:** 5

**Assigned VMs:**

- **VM:** my-ng-influxdb-back-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng
- **VM:** my-ng-influxdb-back-2
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 2 | inbound | ALL |  | 10.0.0.0/16 | inbound * * from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 5 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my-mig-sg-02

**CSP ID:** sg-mj7c5h3pubp9gdjeyqfa | **VNet:** my-mig-vnet-01 | **Rules:** 4

**Assigned VMs:**

- **VM:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng-ec268ed7-821e-9d73-e79f

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ALL |  | 10.0.0.0/16 | inbound * * from 10.0.0.0/16 | Migrated from source |
| 2 | inbound | TCP | 9999 | 0.0.0.0/0 | inbound tcp 9999 | Migrated from source |
| 3 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 4 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source node network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** my-mig-vnet-01<br>**ID:** vpc-mj7bs5xicpvmokn3j8rx6 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my-mig-subnet-01<br>**ID:** vsw-mj725y065dz05gozxllw0 | 10.0.1.0/24 | my-mig-vnet-01 |

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
| 1 | my-mig-sshkey-01 | tbbj6mafh101dhopefl6 | 938129357923abb15fa633e7158ce0ff | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.3342 |
| Daily | $8.02 |
| Monthly | $240.62 |
| Yearly | $2887.49 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | ecs.e-c1m1.large | $12.82 | 5.3% |

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

