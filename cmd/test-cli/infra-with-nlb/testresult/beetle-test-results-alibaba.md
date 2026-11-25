# CM-Beetle test results for ALIBABA (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with ALIBABA cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.5+ (83361e8)
- imdl: v0.1.22+ (83361e8)
- CB-Tumblebug: v0.14.0
- CB-Spider: v0.14.0
- CB-MapUI: v0.14.0
- Target CSP: ALIBABA
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: November 25, 2026
- Test Time: 18:15:38 KST
- Test Execution: 2026-11-25 18:15:38 KST

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
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 2.68s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 471ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 50.707s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 139ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 5ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 27ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 716ms | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 9.517s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 17ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 13ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 1m26.955s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.211s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.313s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 15.712s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m15.883s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 5m3.736830958s

*Test executed on November 25, 2026 at 18:15:38 KST (2026-11-25 18:15:38 KST) using CM-Beetle automated test CLI*

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
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "cloud_essd_entry",
            "rootDiskSize": 40,
            "dataDiskIds": null
          },
          {
            "name": "ng-ip-10-0-1-30-62161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
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
            "zone": "ap-northeast-2a",
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
          "uid": "tb43rhd31thmkk2e8qar",
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
          "id": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "uid": "tbqnoq1dth35qc6b0u8q",
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
          "availabilityCheckedAt": "2026-11-25T09:15:43.11254734Z",
          "availabilityRefreshAfter": "2026-11-25T10:15:43.11254734Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "regionList": [
            "ap-northeast-1",
            "ap-northeast-2",
            "ap-southeast-3",
            "ap-southeast-5",
            "ap-southeast-6",
            "ap-southeast-7",
            "ap-southeast-8",
            "cn-beijing",
            "cn-fuzhou",
            "cn-guangzhou",
            "cn-heyuan",
            "cn-huhehaote",
            "cn-nanjing",
            "cn-shanghai",
            "cn-shenzhen",
            "cn-wuhan-lr",
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
          "id": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "uid": "tb5j7limcvqqqhpebue3",
          "name": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-us-west-1",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:25:32 Thu",
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260916.vhd"
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
              "value": "Kernel version is 5.15.0-191-generic, 2026.9.17"
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
              "value": "v2026.9.17"
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
              "value": "2026-09-17T05:15:48Z"
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260916.vhd"
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
          "description": "Kernel version is 5.15.0-191-generic, 2026.9.17",
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
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 40,
            "dataDiskIds": null
          },
          {
            "name": "ng-ip-10-0-1-30-62161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
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
            "zone": "ap-northeast-2a",
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
          "uid": "tbd6m2ehlkm9i96a7ct6",
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
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ],
          "availabilityStatus": "Unverified"
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "uid": "tbgbunq9i5rdkmv9cmes",
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
          "availabilityCheckedAt": "2026-11-25T09:15:43.110631744Z",
          "availabilityRefreshAfter": "2026-11-25T10:15:43.110631744Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "regionList": [
            "ap-northeast-1",
            "ap-northeast-2",
            "ap-southeast-3",
            "ap-southeast-5",
            "ap-southeast-6",
            "ap-southeast-7",
            "ap-southeast-8",
            "cn-beijing",
            "cn-fuzhou",
            "cn-guangzhou",
            "cn-heyuan",
            "cn-huhehaote",
            "cn-nanjing",
            "cn-shanghai",
            "cn-shenzhen",
            "cn-wuhan-lr",
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
          "id": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "uid": "tb5j7limcvqqqhpebue3",
          "name": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-us-west-1",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:25:32 Thu",
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260916.vhd"
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
              "value": "Kernel version is 5.15.0-191-generic, 2026.9.17"
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
              "value": "v2026.9.17"
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
              "value": "2026-09-17T05:15:48Z"
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260916.vhd"
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
          "description": "Kernel version is 5.15.0-191-generic, 2026.9.17",
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
        "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
        "vNetId": "my04-mig-vnet-01",
        "subnetId": "my04-mig-subnet-01",
        "securityGroupIds": [
          "my04-mig-sg-01"
        ],
        "sshKeyId": "my04-mig-sshkey-01",
        "rootDiskType": "cloud_essd_entry",
        "rootDiskSize": 40,
        "dataDiskIds": null
      },
      {
        "name": "my04-ng-ip-10-0-1-30-62161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
        "connectionName": "alibaba-ap-northeast-2",
        "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
        "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
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
        "zone": "ap-northeast-2a",
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
      "uid": "tb43rhd31thmkk2e8qar",
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
      "id": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
      "uid": "tbqnoq1dth35qc6b0u8q",
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
      "availabilityCheckedAt": "2026-11-25T09:15:43.11254734Z",
      "availabilityRefreshAfter": "2026-11-25T10:15:43.11254734Z"
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "alibaba",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
      "regionList": [
        "ap-northeast-1",
        "ap-northeast-2",
        "ap-southeast-3",
        "ap-southeast-5",
        "ap-southeast-6",
        "ap-southeast-7",
        "ap-southeast-8",
        "cn-beijing",
        "cn-fuzhou",
        "cn-guangzhou",
        "cn-heyuan",
        "cn-huhehaote",
        "cn-nanjing",
        "cn-shanghai",
        "cn-shenzhen",
        "cn-wuhan-lr",
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
      "id": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
      "uid": "tb5j7limcvqqqhpebue3",
      "name": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "alibaba-us-west-1",
      "infraType": "",
      "fetchedTime": "2026.10.29 04:25:32 Thu",
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
          "value": "ubuntu_22_04_x64_20G_alibase_20260916.vhd"
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
          "value": "Kernel version is 5.15.0-191-generic, 2026.9.17"
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
          "value": "v2026.9.17"
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
          "value": "2026-09-17T05:15:48Z"
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
          "value": "ubuntu_22_04_x64_20G_alibase_20260916.vhd"
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
      "description": "Kernel version is 5.15.0-191-generic, 2026.9.17",
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
        "nodeGroupId": "my04-ng-influxdb-back"
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
  "uid": "tbeamcb9hb2d3rp8n5a3",
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
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "cm-migrated-at": "2026-11-25T09:16:04Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my04-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my04-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbeamcb9hb2d3rp8n5a3"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my04-ng-influxdb-back-1",
      "uid": "tbpthhohn7kjt5n614un",
      "cspResourceName": "tbpthhohn7kjt5n614un",
      "cspResourceId": "i-mj7f24f44ih73lbcha5e",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:12",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "8.213.148.200",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.3",
      "privateDNS": "",
      "rootDiskType": "cloud_essd_entry",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-mig-vnet-01",
      "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
      "subnetId": "my04-mig-subnet-01",
      "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
      "networkInterface": "eni-mj7f24f44ih73lb9htrf",
      "securityGroupIds": [
        "my04-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-mig-sshkey-01",
      "cspSshKeyId": "tbskggpajj56tq2usk2e",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "0911bt9!Amkve$",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBM55WXSUP/yVOggAwIbBTLBngJQ0BNUTBLVi3a3q9E2NOszVpblxGM29p9giUykJG84gT5HUlZfyIBXE5sQziU0=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:EQjmnijNTvofKKnxIiyJcmCkk4szxi/1hU0Z6NKv4UA",
        "firstUsedAt": "2026-11-25T09:16:34Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:12Z",
          "completedTime": "2026-11-25T09:16:34Z",
          "elapsedTime": 22,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my04-infra101-tbiu6vo8l46p19n3amm1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:35Z",
          "completedTime": "2026-11-25T09:16:36Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7f24f44ih73lbcha5eZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my04-ng-influxdb-back-2",
      "uid": "tbbj2qfadajevs4b9t50",
      "cspResourceName": "tbbj2qfadajevs4b9t50",
      "cspResourceId": "i-mj7803i8w5ifq7vxuxuc",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:11",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "47.80.61.166",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.2",
      "privateDNS": "",
      "rootDiskType": "cloud_essd_entry",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-mig-vnet-01",
      "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
      "subnetId": "my04-mig-subnet-01",
      "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
      "networkInterface": "eni-mj7803i8w5ifq7vyy4l2",
      "securityGroupIds": [
        "my04-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-mig-sshkey-01",
      "cspSshKeyId": "tbskggpajj56tq2usk2e",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "7v!71ktmsbA$4j",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBN6OaZ5vQG2wEQhKxZ8c9kT7y8Tg2giVYJ59yup69YsjzK/v8je48dMDyMWGDopUfN+NUKC838shb3Jqq5w3jRM=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:30i/DI/mKm1x8N+osKzXeujj4ZwpL5LeQscZkjKOXYs",
        "firstUsedAt": "2026-11-25T09:16:34Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:12Z",
          "completedTime": "2026-11-25T09:16:34Z",
          "elapsedTime": 22,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my04-infra101-tbiu6vo8l46p19n3amm1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:35Z",
          "completedTime": "2026-11-25T09:16:36Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7803i8w5ifq7vxuxucZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my04-ng-ip-10-0-1-30-62161624-1",
      "uid": "tbrvt5o72aejgu84q2d9",
      "cspResourceName": "tbrvt5o72aejgu84q2d9",
      "cspResourceId": "i-mj7ja9zhp5zysim4twue",
      "name": "my04-ng-ip-10-0-1-30-62161624-1",
      "nodeGroupId": "my04-ng-ip-10-0-1-30-62161624",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:11",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "8.220.218.42",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.1",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-mig-vnet-01",
      "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
      "subnetId": "my04-mig-subnet-01",
      "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
      "networkInterface": "eni-mj7ja9zhp5zysim0y47s",
      "securityGroupIds": [
        "my04-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-mig-sshkey-01",
      "cspSshKeyId": "tbskggpajj56tq2usk2e",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "b!h5918eAe$vtb",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBH9xztJQ8TbUAmdLK19nrs5q5MeywZcoHCtZg7vDFW0Ds/sGKeDlLxL5zl7FvtMCL2TCk5ZBxXbFef+9TLJKHYo=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:diFFp3Xe/fz1T2LAw64O7Q/X+Xbb3NqPY7b6iwD2fIo",
        "firstUsedAt": "2026-11-25T09:16:29Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:12Z",
          "completedTime": "2026-11-25T09:16:35Z",
          "elapsedTime": 23,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my04-infra101-tbiu6vo8l46p19n3amm1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:35Z",
          "completedTime": "2026-11-25T09:16:37Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7ja9zhp5zysim4twueZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
        "my04-ng-influxdb-back",
        "my04-ng-ip-10-0-1-30-62161624"
      ],
      "nodeIds": [
        "my04-ng-influxdb-back-1",
        "my04-ng-influxdb-back-2",
        "my04-ng-ip-10-0-1-30-62161624-1"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my04-ng-influxdb-back",
      "representativeNodeId": "my04-ng-influxdb-back-1"
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
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-influxdb-back-2",
            "nodeIp": "47.80.61.166",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7803i8w5ifq7vxuxucZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-influxdb-back-1",
            "nodeIp": "8.213.148.200",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7f24f44ih73lbcha5eZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-ip-10-0-1-30-62161624-1",
            "nodeIp": "8.220.218.42",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7ja9zhp5zysim4twueZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postScriptRequestId": "ps-my04-infra101-tbiu6vo8l46p19n3amm1"
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
      "uid": "tbohd5haepfsekdac2ko",
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
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
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
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
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
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
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
      "id": "my04-infra101",
      "uid": "tbeamcb9hb2d3rp8n5a3",
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
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "cm-migrated-at": "2026-11-25T09:16:04Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my04-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my04-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbeamcb9hb2d3rp8n5a3"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my04-ng-influxdb-back-1",
          "uid": "tbpthhohn7kjt5n614un",
          "cspResourceName": "tbpthhohn7kjt5n614un",
          "cspResourceId": "i-mj7f24f44ih73lbcha5e",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:12",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-11-25 09:16:12",
            "sys.cspResourceId": "i-mj7f24f44ih73lbcha5e",
            "sys.cspResourceName": "tbpthhohn7kjt5n614un",
            "sys.id": "my04-ng-influxdb-back-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-influxdb-back",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbpthhohn7kjt5n614un",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.213.148.200",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.3",
          "privateDNS": "",
          "rootDiskType": "cloud_essd_entry",
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
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
          "networkInterface": "eni-mj7f24f44ih73lb9htrf",
          "securityGroupIds": [
            "my04-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tbskggpajj56tq2usk2e"
        },
        {
          "resourceType": "node",
          "id": "my04-ng-influxdb-back-2",
          "uid": "tbbj2qfadajevs4b9t50",
          "cspResourceName": "tbbj2qfadajevs4b9t50",
          "cspResourceId": "i-mj7803i8w5ifq7vxuxuc",
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
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:11",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-11-25 09:16:11",
            "sys.cspResourceId": "i-mj7803i8w5ifq7vxuxuc",
            "sys.cspResourceName": "tbbj2qfadajevs4b9t50",
            "sys.id": "my04-ng-influxdb-back-2",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-influxdb-back",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbbj2qfadajevs4b9t50",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.61.166",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.2",
          "privateDNS": "",
          "rootDiskType": "cloud_essd_entry",
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
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
          "networkInterface": "eni-mj7803i8w5ifq7vyy4l2",
          "securityGroupIds": [
            "my04-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tbskggpajj56tq2usk2e"
        },
        {
          "resourceType": "node",
          "id": "my04-ng-ip-10-0-1-30-62161624-1",
          "uid": "tbrvt5o72aejgu84q2d9",
          "cspResourceName": "tbrvt5o72aejgu84q2d9",
          "cspResourceId": "i-mj7ja9zhp5zysim4twue",
          "name": "my04-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my04-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-25 09:16:11",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-11-25 09:16:11",
            "sys.cspResourceId": "i-mj7ja9zhp5zysim4twue",
            "sys.cspResourceName": "tbrvt5o72aejgu84q2d9",
            "sys.id": "my04-ng-ip-10-0-1-30-62161624-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-ng-ip-10-0-1-30-62161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-ng-ip-10-0-1-30-62161624",
            "sys.subnetId": "my04-mig-subnet-01",
            "sys.uid": "tbrvt5o72aejgu84q2d9",
            "sys.vNetId": "my04-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.220.218.42",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.1",
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
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-mig-vnet-01",
          "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
          "subnetId": "my04-mig-subnet-01",
          "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
          "networkInterface": "eni-mj7ja9zhp5zysim0y47s",
          "securityGroupIds": [
            "my04-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-mig-sshkey-01",
          "cspSshKeyId": "tbskggpajj56tq2usk2e"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my05-infra101",
      "uid": "tbf56iebln7rh1kjspoe",
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
      "configureCloudAdaptiveNetwork": "",
      "label": null,
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-1",
          "uid": "tb9q3als393k50fpj5ps",
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
          "uid": "tb1qavefj8k9l3vjdh1q",
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
          "id": "my05-ng-ip-10-0-1-30-62161624-1",
          "uid": "tb5qb43uh91ckbefruhm",
          "name": "my05-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my05-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
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
  "uid": "tbeamcb9hb2d3rp8n5a3",
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
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "cm-migrated-at": "2026-11-25T09:16:04Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my04-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my04-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbeamcb9hb2d3rp8n5a3"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my04-ng-influxdb-back-1",
      "uid": "tbpthhohn7kjt5n614un",
      "cspResourceName": "tbpthhohn7kjt5n614un",
      "cspResourceId": "i-mj7f24f44ih73lbcha5e",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:12",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "8.213.148.200",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.3",
      "privateDNS": "",
      "rootDiskType": "cloud_essd_entry",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-mig-vnet-01",
      "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
      "subnetId": "my04-mig-subnet-01",
      "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
      "networkInterface": "eni-mj7f24f44ih73lb9htrf",
      "securityGroupIds": [
        "my04-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-mig-sshkey-01",
      "cspSshKeyId": "tbskggpajj56tq2usk2e",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "0911bt9!Amkve$",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBM55WXSUP/yVOggAwIbBTLBngJQ0BNUTBLVi3a3q9E2NOszVpblxGM29p9giUykJG84gT5HUlZfyIBXE5sQziU0=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:EQjmnijNTvofKKnxIiyJcmCkk4szxi/1hU0Z6NKv4UA",
        "firstUsedAt": "2026-11-25T09:16:34Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:12Z",
          "completedTime": "2026-11-25T09:16:34Z",
          "elapsedTime": 22,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my04-infra101-tbiu6vo8l46p19n3amm1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:35Z",
          "completedTime": "2026-11-25T09:16:36Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7f24f44ih73lbcha5eZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my04-ng-influxdb-back-2",
      "uid": "tbbj2qfadajevs4b9t50",
      "cspResourceName": "tbbj2qfadajevs4b9t50",
      "cspResourceId": "i-mj7803i8w5ifq7vxuxuc",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:11",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "47.80.61.166",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.2",
      "privateDNS": "",
      "rootDiskType": "cloud_essd_entry",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-mig-vnet-01",
      "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
      "subnetId": "my04-mig-subnet-01",
      "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
      "networkInterface": "eni-mj7803i8w5ifq7vyy4l2",
      "securityGroupIds": [
        "my04-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-mig-sshkey-01",
      "cspSshKeyId": "tbskggpajj56tq2usk2e",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "7v!71ktmsbA$4j",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBN6OaZ5vQG2wEQhKxZ8c9kT7y8Tg2giVYJ59yup69YsjzK/v8je48dMDyMWGDopUfN+NUKC838shb3Jqq5w3jRM=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:30i/DI/mKm1x8N+osKzXeujj4ZwpL5LeQscZkjKOXYs",
        "firstUsedAt": "2026-11-25T09:16:34Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:12Z",
          "completedTime": "2026-11-25T09:16:34Z",
          "elapsedTime": 22,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my04-infra101-tbiu6vo8l46p19n3amm1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:35Z",
          "completedTime": "2026-11-25T09:16:36Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7803i8w5ifq7vxuxucZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my04-ng-ip-10-0-1-30-62161624-1",
      "uid": "tbrvt5o72aejgu84q2d9",
      "cspResourceName": "tbrvt5o72aejgu84q2d9",
      "cspResourceId": "i-mj7ja9zhp5zysim4twue",
      "name": "my04-ng-ip-10-0-1-30-62161624-1",
      "nodeGroupId": "my04-ng-ip-10-0-1-30-62161624",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:16:11",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=66.7%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "8.220.218.42",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.1",
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
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260916.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20261016.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-mig-vnet-01",
      "cspVNetId": "vpc-mj7abluu7bv81vyj6k43j",
      "subnetId": "my04-mig-subnet-01",
      "cspSubnetId": "vsw-mj7ipqju3acsuca2qfa7e",
      "networkInterface": "eni-mj7ja9zhp5zysim0y47s",
      "securityGroupIds": [
        "my04-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-mig-sshkey-01",
      "cspSshKeyId": "tbskggpajj56tq2usk2e",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "b!h5918eAe$vtb",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBH9xztJQ8TbUAmdLK19nrs5q5MeywZcoHCtZg7vDFW0Ds/sGKeDlLxL5zl7FvtMCL2TCk5ZBxXbFef+9TLJKHYo=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:diFFp3Xe/fz1T2LAw64O7Q/X+Xbb3NqPY7b6iwD2fIo",
        "firstUsedAt": "2026-11-25T09:16:29Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:12Z",
          "completedTime": "2026-11-25T09:16:35Z",
          "elapsedTime": 23,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "ps-my04-infra101-tbiu6vo8l46p19n3amm1",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:16:35Z",
          "completedTime": "2026-11-25T09:16:37Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7ja9zhp5zysim4twueZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
        "my04-ng-influxdb-back",
        "my04-ng-ip-10-0-1-30-62161624"
      ],
      "nodeIds": [
        "my04-ng-influxdb-back-1",
        "my04-ng-influxdb-back-2",
        "my04-ng-ip-10-0-1-30-62161624-1"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my04-ng-influxdb-back",
      "representativeNodeId": "my04-ng-influxdb-back-1"
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
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-influxdb-back-2",
            "nodeIp": "47.80.61.166",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7803i8w5ifq7vxuxucZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-influxdb-back-1",
            "nodeIp": "8.213.148.200",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7f24f44ih73lbcha5eZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-ng-ip-10-0-1-30-62161624-1",
            "nodeIp": "8.220.218.42",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7ja9zhp5zysim4twueZ 5.15.0-198-generic #208-Ubuntu SMP Fri Sep 4 10:22:38 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postScriptRequestId": "ps-my04-infra101-tbiu6vo8l46p19n3amm1"
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
      "id": "my04-ng-influxdb-back",
      "cspResourceName": "tb084ephbj3j3ao7c7m4",
      "cspResourceId": "lb-mj70hxl0s9hyk87qunerv",
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
        "ip": "43.108.8.225",
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
            "value": "{NameId:tb084ephbj3j3ao7c7m4,SystemId:lb-mj70hxl0s9hyk87qunerv}"
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
            "value": "{Key:BaseResponse,Value:{}}; {Key:Status,Value:running}; {Key:ConnectionDrainTimeout,Value:0}; {Key:HealthCheckTcpFastCloseEnabled,Value:false}; {Key:FullNatEnabled,Value:false}; {Key:RequestId,Value:01A211FE-B693-3687-8793-76DBB97316FD}; {Key:HealthCheckConnectPort,Value:8086}; {Key:Bandwidth,Value:-1}; {Key:HealthCheckType,Value:tcp}; {Key:BackendServerPort,Value:8086}; {Key:AclStatus,Value:off}; {Key:UnhealthyThreshold,Value:3}; {Key:MaxConnection,Value:0}; {Key:ProxyProtocolV2Enabled,Value:false}; {Key:PersistenceTimeout,Value:0}; {Key:ListenerPort,Value:9999}; {Key:HealthCheckInterval,Value:10}; {Key:FailoverThreshold,Value:0}; {Key:Scheduler,Value:wrr}; {Key:EstablishedTimeout,Value:900}; {Key:HealthCheckConnectTimeout,Value:10}; {Key:MasterSlaveModeEnabled,Value:false}; {Key:HealthyThreshold,Value:3}; {Key:HealthCheck,Value:on}; {Key:LoadBalancerId,Value:lb-mj70hxl0s9hyk87qunerv}; {Key:AclIds,Value:{AclId:null}}; {Key:PortRanges,Value:{PortRange:[]}}; {Key:Tags,Value:{Tag:[]}}"
          }
        ]
      },
      "createdTime": "2026-11-25T09:17:12Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "BaseResponse",
          "value": "{}"
        },
        {
          "key": "CreateTimeStamp",
          "value": "1795598232000"
        },
        {
          "key": "HasReservedInfo",
          "value": "false"
        },
        {
          "key": "CreateTime",
          "value": "2026-11-25T09:17:12Z"
        },
        {
          "key": "LoadBalancerId",
          "value": "lb-mj70hxl0s9hyk87qunerv"
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
          "value": "01A211FE-B623-325E-9E89-F32CCB1F4252"
        },
        {
          "key": "Bandwidth",
          "value": "5120"
        },
        {
          "key": "LoadBalancerName",
          "value": "tb084ephbj3j3ao7c7m4"
        },
        {
          "key": "Address",
          "value": "43.108.8.225"
        },
        {
          "key": "SlaveZoneId",
          "value": "ap-northeast-2c"
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
          "value": "{BackendServer:[{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7803i8w5ifq7vxuxuc},{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7f24f44ih73lbcha5e}]}"
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
    "cspResourceName": "tb084ephbj3j3ao7c7m4",
    "cspResourceId": "lb-mj70hxl0s9hyk87qunerv",
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
      "ip": "43.108.8.225",
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
          "value": "{NameId:tb084ephbj3j3ao7c7m4,SystemId:lb-mj70hxl0s9hyk87qunerv}"
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
          "value": "{Key:BaseResponse,Value:{}}; {Key:Status,Value:running}; {Key:ConnectionDrainTimeout,Value:0}; {Key:HealthCheckTcpFastCloseEnabled,Value:false}; {Key:FullNatEnabled,Value:false}; {Key:RequestId,Value:01A211FE-B693-3687-8793-76DBB97316FD}; {Key:HealthCheckConnectPort,Value:8086}; {Key:Bandwidth,Value:-1}; {Key:HealthCheckType,Value:tcp}; {Key:BackendServerPort,Value:8086}; {Key:AclStatus,Value:off}; {Key:UnhealthyThreshold,Value:3}; {Key:MaxConnection,Value:0}; {Key:ProxyProtocolV2Enabled,Value:false}; {Key:PersistenceTimeout,Value:0}; {Key:ListenerPort,Value:9999}; {Key:HealthCheckInterval,Value:10}; {Key:FailoverThreshold,Value:0}; {Key:Scheduler,Value:wrr}; {Key:EstablishedTimeout,Value:900}; {Key:HealthCheckConnectTimeout,Value:10}; {Key:MasterSlaveModeEnabled,Value:false}; {Key:HealthyThreshold,Value:3}; {Key:HealthCheck,Value:on}; {Key:LoadBalancerId,Value:lb-mj70hxl0s9hyk87qunerv}; {Key:AclIds,Value:{AclId:null}}; {Key:PortRanges,Value:{PortRange:[]}}; {Key:Tags,Value:{Tag:[]}}"
        }
      ]
    },
    "createdTime": "2026-11-25T09:17:12Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "BaseResponse",
        "value": "{}"
      },
      {
        "key": "CreateTimeStamp",
        "value": "1795598232000"
      },
      {
        "key": "HasReservedInfo",
        "value": "false"
      },
      {
        "key": "CreateTime",
        "value": "2026-11-25T09:17:12Z"
      },
      {
        "key": "LoadBalancerId",
        "value": "lb-mj70hxl0s9hyk87qunerv"
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
        "value": "01A211FE-B623-325E-9E89-F32CCB1F4252"
      },
      {
        "key": "Bandwidth",
        "value": "5120"
      },
      {
        "key": "LoadBalancerName",
        "value": "tb084ephbj3j3ao7c7m4"
      },
      {
        "key": "Address",
        "value": "43.108.8.225"
      },
      {
        "key": "SlaveZoneId",
        "value": "ap-northeast-2c"
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
        "value": "{BackendServer:[{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7803i8w5ifq7vxuxuc},{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7f24f44ih73lbcha5e}]}"
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
  "cspResourceName": "tb084ephbj3j3ao7c7m4",
  "cspResourceId": "lb-mj70hxl0s9hyk87qunerv",
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
    "ip": "43.108.8.225",
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
        "value": "{NameId:tb084ephbj3j3ao7c7m4,SystemId:lb-mj70hxl0s9hyk87qunerv}"
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
        "value": "{Key:BaseResponse,Value:{}}; {Key:Status,Value:running}; {Key:ConnectionDrainTimeout,Value:0}; {Key:HealthCheckTcpFastCloseEnabled,Value:false}; {Key:FullNatEnabled,Value:false}; {Key:RequestId,Value:01A211FE-B693-3687-8793-76DBB97316FD}; {Key:HealthCheckConnectPort,Value:8086}; {Key:Bandwidth,Value:-1}; {Key:HealthCheckType,Value:tcp}; {Key:BackendServerPort,Value:8086}; {Key:AclStatus,Value:off}; {Key:UnhealthyThreshold,Value:3}; {Key:MaxConnection,Value:0}; {Key:ProxyProtocolV2Enabled,Value:false}; {Key:PersistenceTimeout,Value:0}; {Key:ListenerPort,Value:9999}; {Key:HealthCheckInterval,Value:10}; {Key:FailoverThreshold,Value:0}; {Key:Scheduler,Value:wrr}; {Key:EstablishedTimeout,Value:900}; {Key:HealthCheckConnectTimeout,Value:10}; {Key:MasterSlaveModeEnabled,Value:false}; {Key:HealthyThreshold,Value:3}; {Key:HealthCheck,Value:on}; {Key:LoadBalancerId,Value:lb-mj70hxl0s9hyk87qunerv}; {Key:AclIds,Value:{AclId:null}}; {Key:PortRanges,Value:{PortRange:[]}}; {Key:Tags,Value:{Tag:[]}}"
      }
    ]
  },
  "createdTime": "2026-11-25T09:17:12Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "BaseResponse",
      "value": "{}"
    },
    {
      "key": "CreateTimeStamp",
      "value": "1795598232000"
    },
    {
      "key": "HasReservedInfo",
      "value": "false"
    },
    {
      "key": "CreateTime",
      "value": "2026-11-25T09:17:12Z"
    },
    {
      "key": "LoadBalancerId",
      "value": "lb-mj70hxl0s9hyk87qunerv"
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
      "value": "01A211FE-B623-325E-9E89-F32CCB1F4252"
    },
    {
      "key": "Bandwidth",
      "value": "5120"
    },
    {
      "key": "LoadBalancerName",
      "value": "tb084ephbj3j3ao7c7m4"
    },
    {
      "key": "Address",
      "value": "43.108.8.225"
    },
    {
      "key": "SlaveZoneId",
      "value": "ap-northeast-2c"
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
      "value": "{BackendServer:[{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7803i8w5ifq7vxuxuc},{Type:ecs,VpcId:,Weight:50,Description:,ServerIp:,ServerId:i-mj7f24f44ih73lbcha5e}]}"
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

**Summary**: Load balancing confirmed: traffic distributed across 2 unique VMs

### Test Case 12: Target Infrastructure Summary

#### 12.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Target Infrastructure Summary**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-11-25 09:18:52

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
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| ecs.e-c1m4.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1582 | 2 |
| ecs.e-c1m1.large | 2 | 2.0 | - | x86_64 |  | $0.0178 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ubuntu_22_04_x64_20G_alibase_20260916.vhd | Ubuntu  22.04 64 bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my04-ng-influxdb-back-1 | i-mj7f24f44ih73lbcha5e | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 8.213.148.200<br>**Private IP:** 10.0.1.3<br>**SGs:** my04-mig-sg-01<br>**SSH:** my04-mig-sshkey-01 |
| my04-ng-influxdb-back-2 | i-mj7803i8w5ifq7vxuxuc | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 47.80.61.166<br>**Private IP:** 10.0.1.2<br>**SGs:** my04-mig-sg-01<br>**SSH:** my04-mig-sshkey-01 |
| my04-ng-ip-10-0-1-30-62161624-1 | i-mj7ja9zhp5zysim4twue | Running | 2 vCPU, 2.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-mig-vnet-01<br>**Subnet:** my04-mig-subnet-01<br>**Public IP:** 8.220.218.42<br>**Private IP:** 10.0.1.1<br>**SGs:** my04-mig-sg-02<br>**SSH:** my04-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my04-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my04-mig-vnet-01 |
| **CSP VNet ID** | vpc-mj7abluu7bv81vyj6k43j |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | alibaba-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my04-mig-subnet-01 | vsw-mj7ipqju3acsuca2qfa7e | 10.0.1.0/24 | ap-northeast-2a |

### Network Load Balancers (NLB)

#### NLB: my04-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my04-ng-influxdb-back |
| **CSP NLB Name** | tb084ephbj3j3ao7c7m4 |
| **CSP NLB ID** | lb-mj70hxl0s9hyk87qunerv |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | alibaba-ap-northeast-2 |
| **Listener** | TCP:9999 (IP: 43.108.8.225) |
| **Target Group** | TCP:8086 (Node Group: my04-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my04-ng-influxdb-back-1 | i-mj7f24f44ih73lbcha5e | 10.0.1.3 | 8.213.148.200 | Running |
| my04-ng-influxdb-back-2 | i-mj7803i8w5ifq7vxuxuc | 10.0.1.2 | 47.80.61.166 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my04-mig-sshkey-01 | tbskggpajj56tq2usk2e |  | ab6135bc1f579a0742563628dc47d647 |

### Security Groups

#### Security Group: my04-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my04-mig-sg-01 |
| **CSP Security Group ID** | sg-mj73k9coy7sks4yjzisr |
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
| **CSP Security Group ID** | sg-mj77a670vw20ufh90ntg |
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
| my04-ng-influxdb-back-1 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |
| my04-ng-influxdb-back-2 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |
| my04-ng-ip-10-0-1-30-62161624-1 | ecs.e-c1m1.large | $0.0178 | $12.82 |




### Test Case 13: Migration Report

#### 13.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 13.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-11-25 09:18:57*

---

## 📊 Migration Summary

**Target Cloud:** ALIBABA

**Target Region:** ap-northeast-2

**Namespace:** mig01 | **Infra ID:** my04-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**NLB Migration Status:** Completed

**Total Load Balancers:** 1

**Migrated Load Balancers:** 1

**💰 Estimated Monthly Cost:** $240.62 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | ecs.e-c1m4.xlarge, ecs.e-c1m1.large |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my04-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my04-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 10 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |
| 8 | **Network Load Balancer** | 1 | ✅ Created | my04-ng-influxdb-back (TCP:9999 -> TCP:8086) |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my04-ng-influxdb-back-1<br>**VM ID:** i-mj7f24f44ih73lbcha5e<br>**Label(sourceMachineId):** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 |
| 2 | **VM Name:** my04-ng-influxdb-back-2<br>**VM ID:** i-mj7803i8w5ifq7vxuxuc<br>**Label(sourceMachineId):** ec288dd0-c6fa-8a49-2f60-bc898311febf | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf |
| 3 | **VM Name:** my04-ng-ip-10-0-1-30-62161624-1<br>**VM ID:** i-mj7ja9zhp5zysim4twue<br>**Label(sourceMachineId):** ec268ed7-821e-9d73-e79f-961262161624 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my04-ng-influxdb-back-1 | **Spec ID:** ecs.e-c1m4.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 30 GB |
| 2 | my04-ng-influxdb-back-2 | **Spec ID:** ecs.e-c1m4.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 8 GB<br>**Root Disk:** 30 GB |
| 3 | my04-ng-ip-10-0-1-30-62161624-1 | **Spec ID:** ecs.e-c1m1.large<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 8 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my04-ng-influxdb-back-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260916.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | my04-ng-influxdb-back-2 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260916.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | my04-ng-ip-10-0-1-30-62161624-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260916.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 10 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my04-mig-sg-01

**CSP ID:** sg-mj73k9coy7sks4yjzisr | **VNet:** my04-mig-vnet-01 | **Rules:** 6

**Assigned VMs:**

- **VM:** my04-ng-influxdb-back-1
  - **Source Server:** **Hostname:** ip-10-0-1-221, **Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932
- **VM:** my04-ng-influxdb-back-2
  - **Source Server:** **Hostname:** ip-10-0-1-138, **Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 2 | inbound | TCP | 9999 | 0.0.0.0/0 | - | Created by system |
| 3 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my04-mig-sg-02

**CSP ID:** sg-mj77a670vw20ufh90ntg | **VNet:** my04-mig-vnet-01 | **Rules:** 4

**Assigned VMs:**

- **VM:** my04-ng-ip-10-0-1-30-62161624-1
  - **Source Server:** **Hostname:** ip-10-0-1-30, **Machine ID:** ec268ed7-821e-9d73-e79f-961262161624

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
| 1 | **Name:** my04-mig-vnet-01<br>**ID:** vpc-mj7abluu7bv81vyj6k43j | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my04-mig-subnet-01<br>**ID:** vsw-mj7ipqju3acsuca2qfa7e | 10.0.1.0/24 | my04-mig-vnet-01 |

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

### 1. influxdb_back (haproxy) → my04-ng-influxdb-back

**Configuration Comparison:**

| Property | Source (On-Premise) | Target (Cloud NLB) | Note |
|----------|---------------------|--------------------|------|
| **Load Balancer / Name** | haproxy (Host: ip-10-0-1-30) | my04-ng-influxdb-back (CSP ID: tb084ephbj3j3ao7c7m4) | Replaced by managed NLB |
| **Type / Scope** | Software LB | PUBLIC / REGION | Cloud-native NLB |
| **Frontend Listener** | *:9999 (TCP) | TCP:9999 | IP: 43.108.8.225 |
| **Backend / Target Group** | influxdb_back (TCP:8086, roundrobin) | my04-ng-influxdb-back (TCP:8086, 2 nodes) | Target Group attached to NodeGroup |
| **Health Check** | Enabled (Interval: 10s, Timeout: 10s, Threshold: 3) | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) | Preserved |

**Backend Node Mappings:**

| No. | Source Backend Server | Source Endpoint | Target Node | Target CSP VM ID | Target Private IP | Target Public IP | Target Port | Status |
|-----|-----------------------|-----------------|-------------|------------------|-------------------|------------------|-------------|--------|
| 1 | influx1 (ip-10-0-1-221) | 10.0.1.221:8086 | my04-ng-influxdb-back-1 | i-mj7f24f44ih73lbcha5e | 10.0.1.3 | 8.213.148.200 | 8086 | Running |
| 2 | influx2 (ip-10-0-1-138) | 10.0.1.138:8086 | my04-ng-influxdb-back-2 | i-mj7803i8w5ifq7vxuxuc | 10.0.1.2 | 47.80.61.166 | 8086 | Running |

**Migration Notes:**
- On-premise haproxy on ip-10-0-1-30 (migrated to my04-ng-ip-10-0-1-30-62161624-1) is replaced by cloud-managed NLB.

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my04-mig-sshkey-01 | tbskggpajj56tq2usk2e | ab6135bc1f579a0742563628dc47d647 | Used by all 3 VMs |

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
| ip-10-0-1-221 (migrated) | ecs.e-c1m4.xlarge | $113.90 | 47.3% |
| ip-10-0-1-138 (migrated) | ecs.e-c1m4.xlarge | $113.90 | 47.3% |

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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my04-infra101)",
  "success": true
}
```

