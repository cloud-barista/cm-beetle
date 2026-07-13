# CM-Beetle test results for GCP (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with GCP cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: b418c24
- imdl: unknown
- CB-Tumblebug: Unknown (Fallback to Latest)
- CB-Spider: Unknown (Fallback to Latest)
- CB-MapUI: Unknown (Fallback to Latest)
- Target CSP: GCP
- Target Region: asia-northeast3
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: July 13, 2026
- Test Time: 18:15:27 KST
- Test Execution: 2026-07-13 18:15:27 KST

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

## Test result for GCP

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 11.98s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 3m43.747s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 17ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 4ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 21ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 10.858s | Pass |
| 7 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 43.754s | Pass |
| 8 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 4ms | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 5ms | Pass |
| 10 | NLB Load Balancing Verification | ✅ **PASS** | 1m35.144s | Pass |
| 11 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.215s | Pass |
| 12 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.208s | Pass |
| 13 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 50.438s | Pass |
| 14 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 6m13.93s | Pass |

**Overall Result**: 14/14 tests passed ✅

**Total Duration**: 14m25.665842032s

*Test executed on July 13, 2026 at 18:15:27 KST (2026-07-13 18:15:27 KST) using CM-Beetle automated test CLI*

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
  "desiredCsp": "gcp",
  "desiredRegion": "asia-northeast3",
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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=80.0% Max=100.0% Avg=92.6% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "gcp",
        "region": "asia-northeast3"
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+e2-standard-4",
            "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
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
              "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+e2-highcpu-2",
            "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 10,
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
        "connectionName": "gcp-asia-northeast3",
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
        "connectionName": "gcp-asia-northeast3",
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
          "id": "gcp+asia-northeast3+e2-standard-4",
          "uid": "tbjkou3d0g0gqutn1g4u",
          "cspSpecName": "e2-standard-4",
          "name": "gcp+asia-northeast3+e2-standard-4",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 15.625,
          "diskSizeGB": -1,
          "costPerHour": 0.171931,
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
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "Efficient Instance, 4 vCPUs, 16 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "4"
            },
            {
              "key": "Id",
              "value": "335004"
            },
            {
              "key": "ImageSpaceGb",
              "value": "0"
            },
            {
              "key": "IsSharedCpu",
              "value": "false"
            },
            {
              "key": "Kind",
              "value": "compute#machineType"
            },
            {
              "key": "MaximumPersistentDisks",
              "value": "128"
            },
            {
              "key": "MaximumPersistentDisksSizeGb",
              "value": "263168"
            },
            {
              "key": "MemoryMb",
              "value": "16384"
            },
            {
              "key": "Name",
              "value": "e2-standard-4"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/seokho-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        },
        {
          "id": "gcp+asia-northeast3+e2-highcpu-2",
          "uid": "tb5vdhg2g76dc1jdvku8",
          "cspSpecName": "e2-highcpu-2",
          "name": "gcp+asia-northeast3+e2-highcpu-2",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1.953125,
          "diskSizeGB": -1,
          "costPerHour": 0.063531,
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
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "Efficient Instance, 2 vCPUs, 2 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "2"
            },
            {
              "key": "Id",
              "value": "337002"
            },
            {
              "key": "ImageSpaceGb",
              "value": "0"
            },
            {
              "key": "IsSharedCpu",
              "value": "false"
            },
            {
              "key": "Kind",
              "value": "compute#machineType"
            },
            {
              "key": "MaximumPersistentDisks",
              "value": "128"
            },
            {
              "key": "MaximumPersistentDisksSizeGb",
              "value": "263168"
            },
            {
              "key": "MemoryMb",
              "value": "2048"
            },
            {
              "key": "Name",
              "value": "e2-highcpu-2"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/seokho-etri/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "gcp",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
          "regionList": [
            "common"
          ],
          "id": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
          "uid": "tbgi22itnhj6oho5semi",
          "name": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "gcp-africa-south1",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:06:43 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
          "osDiskType": "NA",
          "osDiskSizeGB": 10,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Architecture",
              "value": "X86_64"
            },
            {
              "key": "ArchiveSizeBytes",
              "value": "3633714624"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-06-23T08:11:14.769-07:00"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
            },
            {
              "key": "DiskSizeGb",
              "value": "10"
            },
            {
              "key": "EnableConfidentialCompute",
              "value": "false"
            },
            {
              "key": "Family",
              "value": "ubuntu-2204-lts"
            },
            {
              "key": "GuestOsFeatures",
              "value": "{type:VIRTIO_SCSI_MULTIQUEUE}; {type:SEV_CAPABLE}; {type:SEV_SNP_CAPABLE}; {type:SEV_LIVE_MIGRATABLE}; {type:SEV_LIVE_MIGRATABLE_V2}; {type:IDPF}; {type:TDX_CAPABLE}; {type:UEFI_COMPATIBLE}; {type:GVNIC}"
            },
            {
              "key": "Id",
              "value": "768974430924119293"
            },
            {
              "key": "Kind",
              "value": "compute#image"
            },
            {
              "key": "LabelFingerprint",
              "value": "iNBmVNCFF9w="
            },
            {
              "key": "Labels",
              "value": "{public-image:true}"
            },
            {
              "key": "LicenseCodes",
              "value": "5511465778777431107"
            },
            {
              "key": "Licenses",
              "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts"
            },
            {
              "key": "Name",
              "value": "ubuntu-2204-jammy-v20260623"
            },
            {
              "key": "RawDisk",
              "value": "{containerType:TAR}"
            },
            {
              "key": "SatisfiesPzi",
              "value": "false"
            },
            {
              "key": "SatisfiesPzs",
              "value": "false"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623"
            },
            {
              "key": "SourceType",
              "value": "RAW"
            },
            {
              "key": "Status",
              "value": "READY"
            },
            {
              "key": "StorageLocations",
              "value": "eu; us; asia"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
          "commandHistory": null
        },
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "gcp",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
          "regionList": [
            "common"
          ],
          "id": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
          "uid": "tb5h42jnp3crmnrvk1qh",
          "name": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "gcp-africa-south1",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:06:52 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23",
          "osDiskType": "NA",
          "osDiskSizeGB": 10,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Architecture",
              "value": "X86_64"
            },
            {
              "key": "ArchiveSizeBytes",
              "value": "3921927744"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-06-23T08:09:32.727-07:00"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23"
            },
            {
              "key": "DiskSizeGb",
              "value": "10"
            },
            {
              "key": "EnableConfidentialCompute",
              "value": "false"
            },
            {
              "key": "Family",
              "value": "ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e"
            },
            {
              "key": "GuestOsFeatures",
              "value": "{type:VIRTIO_SCSI_MULTIQUEUE}; {type:SEV_CAPABLE}; {type:SEV_SNP_CAPABLE}; {type:SEV_LIVE_MIGRATABLE}; {type:SEV_LIVE_MIGRATABLE_V2}; {type:IDPF}; {type:TDX_CAPABLE}; {type:UEFI_COMPATIBLE}; {type:GVNIC}"
            },
            {
              "key": "Id",
              "value": "6725596184982265667"
            },
            {
              "key": "Kind",
              "value": "compute#image"
            },
            {
              "key": "LabelFingerprint",
              "value": "iNBmVNCFF9w="
            },
            {
              "key": "Labels",
              "value": "{public-image:true}"
            },
            {
              "key": "LicenseCodes",
              "value": "5511465778777431107; 3491824550911535745"
            },
            {
              "key": "Licenses",
              "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts; https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/licenses/ubuntu-2204-lts-tpu"
            },
            {
              "key": "Name",
              "value": "ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623"
            },
            {
              "key": "RawDisk",
              "value": "{containerType:TAR}"
            },
            {
              "key": "SatisfiesPzi",
              "value": "false"
            },
            {
              "key": "SatisfiesPzs",
              "value": "false"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623"
            },
            {
              "key": "SourceType",
              "value": "RAW"
            },
            {
              "key": "Status",
              "value": "READY"
            },
            {
              "key": "StorageLocations",
              "value": "eu; asia; us"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "gcp-asia-northeast3",
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
          "connectionName": "gcp-asia-northeast3",
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
            "port": "8086"
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=80.0% Max=100.0% Avg=92.6% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "gcp",
        "region": "asia-northeast3"
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+n2d-standard-4",
            "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
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
              "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+n2d-highcpu-2",
            "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 10,
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
        "connectionName": "gcp-asia-northeast3",
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
        "connectionName": "gcp-asia-northeast3",
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
          "id": "gcp+asia-northeast3+n2d-standard-4",
          "uid": "tbg6v2c6adbnbld6nc05",
          "cspSpecName": "n2d-standard-4",
          "name": "gcp+asia-northeast3+n2d-standard-4",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 15.625,
          "diskSizeGB": -1,
          "costPerHour": 0.217708,
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
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "4 vCPUs 16 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "4"
            },
            {
              "key": "Id",
              "value": "911004"
            },
            {
              "key": "ImageSpaceGb",
              "value": "0"
            },
            {
              "key": "IsSharedCpu",
              "value": "false"
            },
            {
              "key": "Kind",
              "value": "compute#machineType"
            },
            {
              "key": "MaximumPersistentDisks",
              "value": "128"
            },
            {
              "key": "MaximumPersistentDisksSizeGb",
              "value": "263168"
            },
            {
              "key": "MemoryMb",
              "value": "16384"
            },
            {
              "key": "Name",
              "value": "n2d-standard-4"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/seokho-etri/zones/asia-northeast3-a/machineTypes/n2d-standard-4"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        },
        {
          "id": "gcp+asia-northeast3+n2d-highcpu-2",
          "uid": "tbsbe60fntfghnkvihur",
          "cspSpecName": "n2d-highcpu-2",
          "name": "gcp+asia-northeast3+n2d-highcpu-2",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1.953125,
          "diskSizeGB": -1,
          "costPerHour": 0.080366,
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
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "2 vCPUs 2 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "2"
            },
            {
              "key": "Id",
              "value": "910002"
            },
            {
              "key": "ImageSpaceGb",
              "value": "0"
            },
            {
              "key": "IsSharedCpu",
              "value": "false"
            },
            {
              "key": "Kind",
              "value": "compute#machineType"
            },
            {
              "key": "MaximumPersistentDisks",
              "value": "128"
            },
            {
              "key": "MaximumPersistentDisksSizeGb",
              "value": "263168"
            },
            {
              "key": "MemoryMb",
              "value": "2048"
            },
            {
              "key": "Name",
              "value": "n2d-highcpu-2"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/seokho-etri/zones/asia-northeast3-a/machineTypes/n2d-highcpu-2"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "gcp",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
          "regionList": [
            "common"
          ],
          "id": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
          "uid": "tbgi22itnhj6oho5semi",
          "name": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "gcp-africa-south1",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:06:43 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
          "osDiskType": "NA",
          "osDiskSizeGB": 10,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Architecture",
              "value": "X86_64"
            },
            {
              "key": "ArchiveSizeBytes",
              "value": "3633714624"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-06-23T08:11:14.769-07:00"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
            },
            {
              "key": "DiskSizeGb",
              "value": "10"
            },
            {
              "key": "EnableConfidentialCompute",
              "value": "false"
            },
            {
              "key": "Family",
              "value": "ubuntu-2204-lts"
            },
            {
              "key": "GuestOsFeatures",
              "value": "{type:VIRTIO_SCSI_MULTIQUEUE}; {type:SEV_CAPABLE}; {type:SEV_SNP_CAPABLE}; {type:SEV_LIVE_MIGRATABLE}; {type:SEV_LIVE_MIGRATABLE_V2}; {type:IDPF}; {type:TDX_CAPABLE}; {type:UEFI_COMPATIBLE}; {type:GVNIC}"
            },
            {
              "key": "Id",
              "value": "768974430924119293"
            },
            {
              "key": "Kind",
              "value": "compute#image"
            },
            {
              "key": "LabelFingerprint",
              "value": "iNBmVNCFF9w="
            },
            {
              "key": "Labels",
              "value": "{public-image:true}"
            },
            {
              "key": "LicenseCodes",
              "value": "5511465778777431107"
            },
            {
              "key": "Licenses",
              "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts"
            },
            {
              "key": "Name",
              "value": "ubuntu-2204-jammy-v20260623"
            },
            {
              "key": "RawDisk",
              "value": "{containerType:TAR}"
            },
            {
              "key": "SatisfiesPzi",
              "value": "false"
            },
            {
              "key": "SatisfiesPzs",
              "value": "false"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623"
            },
            {
              "key": "SourceType",
              "value": "RAW"
            },
            {
              "key": "Status",
              "value": "READY"
            },
            {
              "key": "StorageLocations",
              "value": "eu; us; asia"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
          "commandHistory": null
        },
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "gcp",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
          "regionList": [
            "common"
          ],
          "id": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
          "uid": "tb5h42jnp3crmnrvk1qh",
          "name": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "gcp-africa-south1",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:06:52 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23",
          "osDiskType": "NA",
          "osDiskSizeGB": 10,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Architecture",
              "value": "X86_64"
            },
            {
              "key": "ArchiveSizeBytes",
              "value": "3921927744"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-06-23T08:09:32.727-07:00"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23"
            },
            {
              "key": "DiskSizeGb",
              "value": "10"
            },
            {
              "key": "EnableConfidentialCompute",
              "value": "false"
            },
            {
              "key": "Family",
              "value": "ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e"
            },
            {
              "key": "GuestOsFeatures",
              "value": "{type:VIRTIO_SCSI_MULTIQUEUE}; {type:SEV_CAPABLE}; {type:SEV_SNP_CAPABLE}; {type:SEV_LIVE_MIGRATABLE}; {type:SEV_LIVE_MIGRATABLE_V2}; {type:IDPF}; {type:TDX_CAPABLE}; {type:UEFI_COMPATIBLE}; {type:GVNIC}"
            },
            {
              "key": "Id",
              "value": "6725596184982265667"
            },
            {
              "key": "Kind",
              "value": "compute#image"
            },
            {
              "key": "LabelFingerprint",
              "value": "iNBmVNCFF9w="
            },
            {
              "key": "Labels",
              "value": "{public-image:true}"
            },
            {
              "key": "LicenseCodes",
              "value": "5511465778777431107; 3491824550911535745"
            },
            {
              "key": "Licenses",
              "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts; https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/licenses/ubuntu-2204-lts-tpu"
            },
            {
              "key": "Name",
              "value": "ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623"
            },
            {
              "key": "RawDisk",
              "value": "{containerType:TAR}"
            },
            {
              "key": "SatisfiesPzi",
              "value": "false"
            },
            {
              "key": "SatisfiesPzs",
              "value": "false"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623"
            },
            {
              "key": "SourceType",
              "value": "RAW"
            },
            {
              "key": "Status",
              "value": "READY"
            },
            {
              "key": "StorageLocations",
              "value": "eu; asia; us"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "gcp-asia-northeast3",
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
          "connectionName": "gcp-asia-northeast3",
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
            "port": "8086"
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
  "uid": "tb6llm23eif1q7e2v20j",
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
    "sys.uid": "tb6llm23eif1q7e2v20j"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb1pu9f78qonoie5ajod",
      "cspResourceName": "tb1pu9f78qonoie5ajod",
      "cspResourceId": "tb1pu9f78qonoie5ajod",
      "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-13 09:19:02",
      "label": {
        "keypair": "tb0ffm8vtk5l3ahog7h1",
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-07-13 09:19:02",
        "sys.cspResourceId": "tb1pu9f78qonoie5ajod",
        "sys.cspResourceName": "tb1pu9f78qonoie5ajod",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb1pu9f78qonoie5ajod",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "34.64.195.188",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 10,
      "RootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+e2-highcpu-2",
      "cspSpecName": "e2-highcpu-2",
      "spec": {
        "cspSpecName": "e2-highcpu-2",
        "vCPU": 2,
        "memoryGiB": 1.953125,
        "costPerHour": 0.063531
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "tb6qc3e779ji5e5kqg7k",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "tb08efe9m19o55jbt2pp",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tb0ffm8vtk5l3ahog7h1",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGqm7KRJjGvBKrIt6Exfjzucx67wuUm/nj+gaOV9Ky6kP1o0gIJxdDPsEgLQDWLTWyIkQBD06kUCHIWenkF806Y=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:AW1o1PfIfEsqC0Q956shJ5kMBWEbmKsxnd/txVkAPM4",
        "firstUsedAt": "2026-07-13T09:19:13Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:19:12Z",
          "completedTime": "2026-07-13T09:19:14Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb1pu9f78qonoie5ajod 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-07-13T02:18:20.464-07:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts,https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/licenses/ubuntu-2204-lts-tpu],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/tb1pu9f78qonoie5ajod,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "YdsSlUtErnc="
        },
        {
          "key": "Id",
          "value": "69656031750853555"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "hpJPIa5xQEk="
        },
        {
          "key": "Labels",
          "value": "{keypair:tb0ffm8vtk5l3ahog7h1}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-07-13T02:18:28.104-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:EaVRkxGYCT8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNYQTsR6YY0L8RJs96TMpl5gPye9R63LiKkn3OqgmFStcBfQyThBeoteg37qHvu3ZHriU8a27iZg8jfhi9OqJidbwljur/rQF4nwXLsrSj6wfagobNIo1LUSxLH0dDUbmDz4CvgC9e+d57mWyPfIkalRfoORxyxQpO1jD8TCwXXPkBLxlxRoftvZiGiqmqobThpWkrLjUQVPNohXciJJQNPJEOvo6UV/X2aECsgvr3fXA/vnZij71tE9MBRWH1cZCjyhGchsTvxaUJL8OjiCjM5VKCSmgZ3KhvGQC5QP+wDogAm/9XClt2fgNlnhS2suKpPXT7gnqxqHGRcb6cEEe/hOAmn8rsrWA0aiT1/YiPXSik3hd35mJ/4nFposCH4d7n323GnMimjvy4nfc1oJr3w9wHXp3xzoDYEDykAqml8Q50xmBW/TJeQimBg8uCbZvSGJtpACxuivu2e62LCUnlhJTfPoPJS3OUEm5NAVr02gEoWuWUvKWmzXxWzRG2kGTVI0MLEUtBQzvx+32ZawMY/Lt88GcprahddXXDEYDdVugxAXbOxKkua99YgJ6YPL7nlHtWEbNI+iB4ndWvh4vLuDdHc5LMamD82nG3wRlGbRtsy7XdYycZa5jveKIZAR8bs07vd2uK3E0op5c/TgWEuiW0d6G22whWMopulZcPAQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tb1pu9f78qonoie5ajod"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.195.188,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:wBXwrarXbvk=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/tb6qc3e779ji5e5kqg7k,networkIP:10.0.1.4,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/tb08efe9m19o55jbt2pp}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/tb1pu9f78qonoie5ajod"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:fXE212sEmT4=,items:[tb2kpd806s3hu3mifdtj]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tbrah6jesqlv0jaeqia5",
      "cspResourceName": "tbrah6jesqlv0jaeqia5",
      "cspResourceId": "tbrah6jesqlv0jaeqia5",
      "name": "my-ng-influxdb-back-1",
      "nodeGroupId": "my-ng-influxdb-back",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-13 09:18:59",
      "label": {
        "keypair": "tb0ffm8vtk5l3ahog7h1",
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-07-13 09:18:59",
        "sys.cspResourceId": "tbrah6jesqlv0jaeqia5",
        "sys.cspResourceName": "tbrah6jesqlv0jaeqia5",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbrah6jesqlv0jaeqia5",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "34.47.124.15",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.3",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 30,
      "RootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+e2-standard-4",
      "cspSpecName": "e2-standard-4",
      "spec": {
        "cspSpecName": "e2-standard-4",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.171931
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "tb6qc3e779ji5e5kqg7k",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "tb08efe9m19o55jbt2pp",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tb0ffm8vtk5l3ahog7h1",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBDkAdCzKgYDJPyL7tedFntUSar5CmzrXe1F9Zy/hWbQoPMUFAyu/f+p7/5nWMBstIbvt5xMQeVryRG7Gtv+OFQ=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:tTZZ/sg0gbpE3c3MGAbQPZw1sKMEpPC0b7/hMKaqnxE",
        "firstUsedAt": "2026-07-13T09:19:13Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:19:12Z",
          "completedTime": "2026-07-13T09:19:14Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbrah6jesqlv0jaeqia5 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-07-13T02:18:18.359-07:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:30,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/tbrah6jesqlv0jaeqia5,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "f76-KtKrepE="
        },
        {
          "key": "Id",
          "value": "6820223286706246581"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "hpJPIa5xQEk="
        },
        {
          "key": "Labels",
          "value": "{keypair:tb0ffm8vtk5l3ahog7h1}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-07-13T02:18:26.195-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:EaVRkxGYCT8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNYQTsR6YY0L8RJs96TMpl5gPye9R63LiKkn3OqgmFStcBfQyThBeoteg37qHvu3ZHriU8a27iZg8jfhi9OqJidbwljur/rQF4nwXLsrSj6wfagobNIo1LUSxLH0dDUbmDz4CvgC9e+d57mWyPfIkalRfoORxyxQpO1jD8TCwXXPkBLxlxRoftvZiGiqmqobThpWkrLjUQVPNohXciJJQNPJEOvo6UV/X2aECsgvr3fXA/vnZij71tE9MBRWH1cZCjyhGchsTvxaUJL8OjiCjM5VKCSmgZ3KhvGQC5QP+wDogAm/9XClt2fgNlnhS2suKpPXT7gnqxqHGRcb6cEEe/hOAmn8rsrWA0aiT1/YiPXSik3hd35mJ/4nFposCH4d7n323GnMimjvy4nfc1oJr3w9wHXp3xzoDYEDykAqml8Q50xmBW/TJeQimBg8uCbZvSGJtpACxuivu2e62LCUnlhJTfPoPJS3OUEm5NAVr02gEoWuWUvKWmzXxWzRG2kGTVI0MLEUtBQzvx+32ZawMY/Lt88GcprahddXXDEYDdVugxAXbOxKkua99YgJ6YPL7nlHtWEbNI+iB4ndWvh4vLuDdHc5LMamD82nG3wRlGbRtsy7XdYycZa5jveKIZAR8bs07vd2uK3E0op5c/TgWEuiW0d6G22whWMopulZcPAQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tbrah6jesqlv0jaeqia5"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.47.124.15,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:Gp6cH14WOLg=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/tb6qc3e779ji5e5kqg7k,networkIP:10.0.1.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/tb08efe9m19o55jbt2pp}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/tbrah6jesqlv0jaeqia5"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:7VBbY9V0wIg=,items:[tb74avigphtcuf3958vc]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tbq36hhnk3jjdfpfuu73",
      "cspResourceName": "tbq36hhnk3jjdfpfuu73",
      "cspResourceId": "tbq36hhnk3jjdfpfuu73",
      "name": "my-ng-influxdb-back-2",
      "nodeGroupId": "my-ng-influxdb-back",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-13 09:18:58",
      "label": {
        "keypair": "tb0ffm8vtk5l3ahog7h1",
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-07-13 09:18:58",
        "sys.cspResourceId": "tbq36hhnk3jjdfpfuu73",
        "sys.cspResourceName": "tbq36hhnk3jjdfpfuu73",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbq36hhnk3jjdfpfuu73",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "8.230.19.201",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.2",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 30,
      "RootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+e2-standard-4",
      "cspSpecName": "e2-standard-4",
      "spec": {
        "cspSpecName": "e2-standard-4",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.171931
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "tb6qc3e779ji5e5kqg7k",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "tb08efe9m19o55jbt2pp",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tb0ffm8vtk5l3ahog7h1",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBN0hMfrekkyROVSveoZGsCeo0oPeeDORUexhxHc5peKEzw1/abt2djbg20vUUqqGrbH1d15K1gpSj6AZBQqj/V8=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:EdgVwUKxu+60R9ACsX+IxKJ2XBexCbLL494NU7S4fnk",
        "firstUsedAt": "2026-07-13T09:19:12Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:19:12Z",
          "completedTime": "2026-07-13T09:19:15Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbq36hhnk3jjdfpfuu73 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-07-13T02:18:17.910-07:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:30,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/tbq36hhnk3jjdfpfuu73,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "9ZZizxjggT8="
        },
        {
          "key": "Id",
          "value": "3947188695974089654"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "hpJPIa5xQEk="
        },
        {
          "key": "Labels",
          "value": "{keypair:tb0ffm8vtk5l3ahog7h1}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-07-13T02:18:25.872-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:EaVRkxGYCT8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNYQTsR6YY0L8RJs96TMpl5gPye9R63LiKkn3OqgmFStcBfQyThBeoteg37qHvu3ZHriU8a27iZg8jfhi9OqJidbwljur/rQF4nwXLsrSj6wfagobNIo1LUSxLH0dDUbmDz4CvgC9e+d57mWyPfIkalRfoORxyxQpO1jD8TCwXXPkBLxlxRoftvZiGiqmqobThpWkrLjUQVPNohXciJJQNPJEOvo6UV/X2aECsgvr3fXA/vnZij71tE9MBRWH1cZCjyhGchsTvxaUJL8OjiCjM5VKCSmgZ3KhvGQC5QP+wDogAm/9XClt2fgNlnhS2suKpPXT7gnqxqHGRcb6cEEe/hOAmn8rsrWA0aiT1/YiPXSik3hd35mJ/4nFposCH4d7n323GnMimjvy4nfc1oJr3w9wHXp3xzoDYEDykAqml8Q50xmBW/TJeQimBg8uCbZvSGJtpACxuivu2e62LCUnlhJTfPoPJS3OUEm5NAVr02gEoWuWUvKWmzXxWzRG2kGTVI0MLEUtBQzvx+32ZawMY/Lt88GcprahddXXDEYDdVugxAXbOxKkua99YgJ6YPL7nlHtWEbNI+iB4ndWvh4vLuDdHc5LMamD82nG3wRlGbRtsy7XdYycZa5jveKIZAR8bs07vd2uK3E0op5c/TgWEuiW0d6G22whWMopulZcPAQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tbq36hhnk3jjdfpfuu73"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:8.230.19.201,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:BG1-eHXkxAs=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/tb6qc3e779ji5e5kqg7k,networkIP:10.0.1.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/tb08efe9m19o55jbt2pp}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/tbq36hhnk3jjdfpfuu73"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:7VBbY9V0wIg=,items:[tb74avigphtcuf3958vc]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
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
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "34.47.124.15",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbrah6jesqlv0jaeqia5 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "34.64.195.188",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tb1pu9f78qonoie5ajod 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "8.230.19.201",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbq36hhnk3jjdfpfuu73 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "tb6llm23eif1q7e2v20j",
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
        "sys.uid": "tb6llm23eif1q7e2v20j"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb1pu9f78qonoie5ajod",
          "cspResourceName": "tb1pu9f78qonoie5ajod",
          "cspResourceId": "tb1pu9f78qonoie5ajod",
          "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-07-13 09:19:02",
          "label": {
            "keypair": "tb0ffm8vtk5l3ahog7h1",
            "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-07-13 09:19:02",
            "sys.cspResourceId": "tb1pu9f78qonoie5ajod",
            "sys.cspResourceName": "tb1pu9f78qonoie5ajod",
            "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tb1pu9f78qonoie5ajod",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.64.195.188",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": 10,
          "RootDeviceName": "persistent-disk-0",
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
          "specId": "gcp+asia-northeast3+e2-highcpu-2",
          "cspSpecName": "e2-highcpu-2",
          "spec": {
            "cspSpecName": "e2-highcpu-2",
            "vCPU": 2,
            "memoryGiB": 1.953125,
            "costPerHour": 0.063531
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "tb6qc3e779ji5e5kqg7k",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "tb08efe9m19o55jbt2pp",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tb0ffm8vtk5l3ahog7h1",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGqm7KRJjGvBKrIt6Exfjzucx67wuUm/nj+gaOV9Ky6kP1o0gIJxdDPsEgLQDWLTWyIkQBD06kUCHIWenkF806Y=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:AW1o1PfIfEsqC0Q956shJ5kMBWEbmKsxnd/txVkAPM4",
            "firstUsedAt": "2026-07-13T09:19:13Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-13T09:19:12Z",
              "completedTime": "2026-07-13T09:19:14Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tb1pu9f78qonoie5ajod 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "CanIpForward",
              "value": "false"
            },
            {
              "key": "CpuPlatform",
              "value": "Intel Broadwell"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-07-13T02:18:20.464-07:00"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "Description",
              "value": "compute sample instance"
            },
            {
              "key": "Disks",
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts,https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/licenses/ubuntu-2204-lts-tpu],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/tb1pu9f78qonoie5ajod,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "YdsSlUtErnc="
            },
            {
              "key": "Id",
              "value": "69656031750853555"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "hpJPIa5xQEk="
            },
            {
              "key": "Labels",
              "value": "{keypair:tb0ffm8vtk5l3ahog7h1}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2026-07-13T02:18:28.104-07:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:EaVRkxGYCT8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNYQTsR6YY0L8RJs96TMpl5gPye9R63LiKkn3OqgmFStcBfQyThBeoteg37qHvu3ZHriU8a27iZg8jfhi9OqJidbwljur/rQF4nwXLsrSj6wfagobNIo1LUSxLH0dDUbmDz4CvgC9e+d57mWyPfIkalRfoORxyxQpO1jD8TCwXXPkBLxlxRoftvZiGiqmqobThpWkrLjUQVPNohXciJJQNPJEOvo6UV/X2aECsgvr3fXA/vnZij71tE9MBRWH1cZCjyhGchsTvxaUJL8OjiCjM5VKCSmgZ3KhvGQC5QP+wDogAm/9XClt2fgNlnhS2suKpPXT7gnqxqHGRcb6cEEe/hOAmn8rsrWA0aiT1/YiPXSik3hd35mJ/4nFposCH4d7n323GnMimjvy4nfc1oJr3w9wHXp3xzoDYEDykAqml8Q50xmBW/TJeQimBg8uCbZvSGJtpACxuivu2e62LCUnlhJTfPoPJS3OUEm5NAVr02gEoWuWUvKWmzXxWzRG2kGTVI0MLEUtBQzvx+32ZawMY/Lt88GcprahddXXDEYDdVugxAXbOxKkua99YgJ6YPL7nlHtWEbNI+iB4ndWvh4vLuDdHc5LMamD82nG3wRlGbRtsy7XdYycZa5jveKIZAR8bs07vd2uK3E0op5c/TgWEuiW0d6G22whWMopulZcPAQ== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "tb1pu9f78qonoie5ajod"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.195.188,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:wBXwrarXbvk=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/tb6qc3e779ji5e5kqg7k,networkIP:10.0.1.4,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/tb08efe9m19o55jbt2pp}"
            },
            {
              "key": "ResourceStatus",
              "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
            },
            {
              "key": "SatisfiesPzi",
              "value": "true"
            },
            {
              "key": "SatisfiesPzs",
              "value": "false"
            },
            {
              "key": "Scheduling",
              "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/tb1pu9f78qonoie5ajod"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
            },
            {
              "key": "ShieldedInstanceConfig",
              "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
            },
            {
              "key": "ShieldedInstanceIntegrityPolicy",
              "value": "{updateAutoLearnPolicy:true}"
            },
            {
              "key": "StartRestricted",
              "value": "false"
            },
            {
              "key": "Status",
              "value": "RUNNING"
            },
            {
              "key": "Tags",
              "value": "{fingerprint:fXE212sEmT4=,items:[tb2kpd806s3hu3mifdtj]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-1",
          "uid": "tbrah6jesqlv0jaeqia5",
          "cspResourceName": "tbrah6jesqlv0jaeqia5",
          "cspResourceId": "tbrah6jesqlv0jaeqia5",
          "name": "my-ng-influxdb-back-1",
          "nodeGroupId": "my-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-07-13 09:18:59",
          "label": {
            "keypair": "tb0ffm8vtk5l3ahog7h1",
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-07-13 09:18:59",
            "sys.cspResourceId": "tbrah6jesqlv0jaeqia5",
            "sys.cspResourceName": "tbrah6jesqlv0jaeqia5",
            "sys.id": "my-ng-influxdb-back-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbrah6jesqlv0jaeqia5",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.47.124.15",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.3",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": 30,
          "RootDeviceName": "persistent-disk-0",
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
          "specId": "gcp+asia-northeast3+e2-standard-4",
          "cspSpecName": "e2-standard-4",
          "spec": {
            "cspSpecName": "e2-standard-4",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.171931
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "tb6qc3e779ji5e5kqg7k",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "tb08efe9m19o55jbt2pp",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tb0ffm8vtk5l3ahog7h1",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBDkAdCzKgYDJPyL7tedFntUSar5CmzrXe1F9Zy/hWbQoPMUFAyu/f+p7/5nWMBstIbvt5xMQeVryRG7Gtv+OFQ=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:tTZZ/sg0gbpE3c3MGAbQPZw1sKMEpPC0b7/hMKaqnxE",
            "firstUsedAt": "2026-07-13T09:19:13Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-13T09:19:12Z",
              "completedTime": "2026-07-13T09:19:14Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbrah6jesqlv0jaeqia5 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "CanIpForward",
              "value": "false"
            },
            {
              "key": "CpuPlatform",
              "value": "Intel Broadwell"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-07-13T02:18:18.359-07:00"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "Description",
              "value": "compute sample instance"
            },
            {
              "key": "Disks",
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:30,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/tbrah6jesqlv0jaeqia5,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "f76-KtKrepE="
            },
            {
              "key": "Id",
              "value": "6820223286706246581"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "hpJPIa5xQEk="
            },
            {
              "key": "Labels",
              "value": "{keypair:tb0ffm8vtk5l3ahog7h1}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2026-07-13T02:18:26.195-07:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:EaVRkxGYCT8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNYQTsR6YY0L8RJs96TMpl5gPye9R63LiKkn3OqgmFStcBfQyThBeoteg37qHvu3ZHriU8a27iZg8jfhi9OqJidbwljur/rQF4nwXLsrSj6wfagobNIo1LUSxLH0dDUbmDz4CvgC9e+d57mWyPfIkalRfoORxyxQpO1jD8TCwXXPkBLxlxRoftvZiGiqmqobThpWkrLjUQVPNohXciJJQNPJEOvo6UV/X2aECsgvr3fXA/vnZij71tE9MBRWH1cZCjyhGchsTvxaUJL8OjiCjM5VKCSmgZ3KhvGQC5QP+wDogAm/9XClt2fgNlnhS2suKpPXT7gnqxqHGRcb6cEEe/hOAmn8rsrWA0aiT1/YiPXSik3hd35mJ/4nFposCH4d7n323GnMimjvy4nfc1oJr3w9wHXp3xzoDYEDykAqml8Q50xmBW/TJeQimBg8uCbZvSGJtpACxuivu2e62LCUnlhJTfPoPJS3OUEm5NAVr02gEoWuWUvKWmzXxWzRG2kGTVI0MLEUtBQzvx+32ZawMY/Lt88GcprahddXXDEYDdVugxAXbOxKkua99YgJ6YPL7nlHtWEbNI+iB4ndWvh4vLuDdHc5LMamD82nG3wRlGbRtsy7XdYycZa5jveKIZAR8bs07vd2uK3E0op5c/TgWEuiW0d6G22whWMopulZcPAQ== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "tbrah6jesqlv0jaeqia5"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.47.124.15,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:Gp6cH14WOLg=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/tb6qc3e779ji5e5kqg7k,networkIP:10.0.1.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/tb08efe9m19o55jbt2pp}"
            },
            {
              "key": "ResourceStatus",
              "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
            },
            {
              "key": "SatisfiesPzi",
              "value": "true"
            },
            {
              "key": "SatisfiesPzs",
              "value": "false"
            },
            {
              "key": "Scheduling",
              "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/tbrah6jesqlv0jaeqia5"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
            },
            {
              "key": "ShieldedInstanceConfig",
              "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
            },
            {
              "key": "ShieldedInstanceIntegrityPolicy",
              "value": "{updateAutoLearnPolicy:true}"
            },
            {
              "key": "StartRestricted",
              "value": "false"
            },
            {
              "key": "Status",
              "value": "RUNNING"
            },
            {
              "key": "Tags",
              "value": "{fingerprint:7VBbY9V0wIg=,items:[tb74avigphtcuf3958vc]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-2",
          "uid": "tbq36hhnk3jjdfpfuu73",
          "cspResourceName": "tbq36hhnk3jjdfpfuu73",
          "cspResourceId": "tbq36hhnk3jjdfpfuu73",
          "name": "my-ng-influxdb-back-2",
          "nodeGroupId": "my-ng-influxdb-back",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-07-13 09:18:58",
          "label": {
            "keypair": "tb0ffm8vtk5l3ahog7h1",
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-07-13 09:18:58",
            "sys.cspResourceId": "tbq36hhnk3jjdfpfuu73",
            "sys.cspResourceName": "tbq36hhnk3jjdfpfuu73",
            "sys.id": "my-ng-influxdb-back-2",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbq36hhnk3jjdfpfuu73",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "8.230.19.201",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.2",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": 30,
          "RootDeviceName": "persistent-disk-0",
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
          "specId": "gcp+asia-northeast3+e2-standard-4",
          "cspSpecName": "e2-standard-4",
          "spec": {
            "cspSpecName": "e2-standard-4",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.171931
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "tb6qc3e779ji5e5kqg7k",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "tb08efe9m19o55jbt2pp",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tb0ffm8vtk5l3ahog7h1",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBN0hMfrekkyROVSveoZGsCeo0oPeeDORUexhxHc5peKEzw1/abt2djbg20vUUqqGrbH1d15K1gpSj6AZBQqj/V8=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:EdgVwUKxu+60R9ACsX+IxKJ2XBexCbLL494NU7S4fnk",
            "firstUsedAt": "2026-07-13T09:19:12Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-13T09:19:12Z",
              "completedTime": "2026-07-13T09:19:15Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbq36hhnk3jjdfpfuu73 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "CanIpForward",
              "value": "false"
            },
            {
              "key": "CpuPlatform",
              "value": "Intel Broadwell"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-07-13T02:18:17.910-07:00"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "Description",
              "value": "compute sample instance"
            },
            {
              "key": "Disks",
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:30,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/tbq36hhnk3jjdfpfuu73,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "9ZZizxjggT8="
            },
            {
              "key": "Id",
              "value": "3947188695974089654"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "hpJPIa5xQEk="
            },
            {
              "key": "Labels",
              "value": "{keypair:tb0ffm8vtk5l3ahog7h1}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2026-07-13T02:18:25.872-07:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:EaVRkxGYCT8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNYQTsR6YY0L8RJs96TMpl5gPye9R63LiKkn3OqgmFStcBfQyThBeoteg37qHvu3ZHriU8a27iZg8jfhi9OqJidbwljur/rQF4nwXLsrSj6wfagobNIo1LUSxLH0dDUbmDz4CvgC9e+d57mWyPfIkalRfoORxyxQpO1jD8TCwXXPkBLxlxRoftvZiGiqmqobThpWkrLjUQVPNohXciJJQNPJEOvo6UV/X2aECsgvr3fXA/vnZij71tE9MBRWH1cZCjyhGchsTvxaUJL8OjiCjM5VKCSmgZ3KhvGQC5QP+wDogAm/9XClt2fgNlnhS2suKpPXT7gnqxqHGRcb6cEEe/hOAmn8rsrWA0aiT1/YiPXSik3hd35mJ/4nFposCH4d7n323GnMimjvy4nfc1oJr3w9wHXp3xzoDYEDykAqml8Q50xmBW/TJeQimBg8uCbZvSGJtpACxuivu2e62LCUnlhJTfPoPJS3OUEm5NAVr02gEoWuWUvKWmzXxWzRG2kGTVI0MLEUtBQzvx+32ZawMY/Lt88GcprahddXXDEYDdVugxAXbOxKkua99YgJ6YPL7nlHtWEbNI+iB4ndWvh4vLuDdHc5LMamD82nG3wRlGbRtsy7XdYycZa5jveKIZAR8bs07vd2uK3E0op5c/TgWEuiW0d6G22whWMopulZcPAQ== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "tbq36hhnk3jjdfpfuu73"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:8.230.19.201,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:BG1-eHXkxAs=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/tb6qc3e779ji5e5kqg7k,networkIP:10.0.1.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/tb08efe9m19o55jbt2pp}"
            },
            {
              "key": "ResourceStatus",
              "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
            },
            {
              "key": "SatisfiesPzi",
              "value": "true"
            },
            {
              "key": "SatisfiesPzs",
              "value": "false"
            },
            {
              "key": "Scheduling",
              "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/tbq36hhnk3jjdfpfuu73"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
            },
            {
              "key": "ShieldedInstanceConfig",
              "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
            },
            {
              "key": "ShieldedInstanceIntegrityPolicy",
              "value": "{updateAutoLearnPolicy:true}"
            },
            {
              "key": "StartRestricted",
              "value": "false"
            },
            {
              "key": "Status",
              "value": "RUNNING"
            },
            {
              "key": "Tags",
              "value": "{fingerprint:7VBbY9V0wIg=,items:[tb74avigphtcuf3958vc]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
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
            "nodeId": "my-ng-influxdb-back-1",
            "nodeIp": "34.47.124.15",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbrah6jesqlv0jaeqia5 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "34.64.195.188",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb1pu9f78qonoie5ajod 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-2",
            "nodeIp": "8.230.19.201",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbq36hhnk3jjdfpfuu73 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "tb6llm23eif1q7e2v20j",
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
    "sys.uid": "tb6llm23eif1q7e2v20j"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb1pu9f78qonoie5ajod",
      "cspResourceName": "tb1pu9f78qonoie5ajod",
      "cspResourceId": "tb1pu9f78qonoie5ajod",
      "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-13 09:19:02",
      "label": {
        "keypair": "tb0ffm8vtk5l3ahog7h1",
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-07-13 09:19:02",
        "sys.cspResourceId": "tb1pu9f78qonoie5ajod",
        "sys.cspResourceName": "tb1pu9f78qonoie5ajod",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb1pu9f78qonoie5ajod",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "34.64.195.188",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 10,
      "RootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+e2-highcpu-2",
      "cspSpecName": "e2-highcpu-2",
      "spec": {
        "cspSpecName": "e2-highcpu-2",
        "vCPU": 2,
        "memoryGiB": 1.953125,
        "costPerHour": 0.063531
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "tb6qc3e779ji5e5kqg7k",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "tb08efe9m19o55jbt2pp",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tb0ffm8vtk5l3ahog7h1",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGqm7KRJjGvBKrIt6Exfjzucx67wuUm/nj+gaOV9Ky6kP1o0gIJxdDPsEgLQDWLTWyIkQBD06kUCHIWenkF806Y=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:AW1o1PfIfEsqC0Q956shJ5kMBWEbmKsxnd/txVkAPM4",
        "firstUsedAt": "2026-07-13T09:19:13Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:19:12Z",
          "completedTime": "2026-07-13T09:19:14Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb1pu9f78qonoie5ajod 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-07-13T02:18:20.464-07:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts,https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/licenses/ubuntu-2204-lts-tpu],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/tb1pu9f78qonoie5ajod,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "YdsSlUtErnc="
        },
        {
          "key": "Id",
          "value": "69656031750853555"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "hpJPIa5xQEk="
        },
        {
          "key": "Labels",
          "value": "{keypair:tb0ffm8vtk5l3ahog7h1}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-07-13T02:18:28.104-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:EaVRkxGYCT8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNYQTsR6YY0L8RJs96TMpl5gPye9R63LiKkn3OqgmFStcBfQyThBeoteg37qHvu3ZHriU8a27iZg8jfhi9OqJidbwljur/rQF4nwXLsrSj6wfagobNIo1LUSxLH0dDUbmDz4CvgC9e+d57mWyPfIkalRfoORxyxQpO1jD8TCwXXPkBLxlxRoftvZiGiqmqobThpWkrLjUQVPNohXciJJQNPJEOvo6UV/X2aECsgvr3fXA/vnZij71tE9MBRWH1cZCjyhGchsTvxaUJL8OjiCjM5VKCSmgZ3KhvGQC5QP+wDogAm/9XClt2fgNlnhS2suKpPXT7gnqxqHGRcb6cEEe/hOAmn8rsrWA0aiT1/YiPXSik3hd35mJ/4nFposCH4d7n323GnMimjvy4nfc1oJr3w9wHXp3xzoDYEDykAqml8Q50xmBW/TJeQimBg8uCbZvSGJtpACxuivu2e62LCUnlhJTfPoPJS3OUEm5NAVr02gEoWuWUvKWmzXxWzRG2kGTVI0MLEUtBQzvx+32ZawMY/Lt88GcprahddXXDEYDdVugxAXbOxKkua99YgJ6YPL7nlHtWEbNI+iB4ndWvh4vLuDdHc5LMamD82nG3wRlGbRtsy7XdYycZa5jveKIZAR8bs07vd2uK3E0op5c/TgWEuiW0d6G22whWMopulZcPAQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tb1pu9f78qonoie5ajod"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.195.188,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:wBXwrarXbvk=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/tb6qc3e779ji5e5kqg7k,networkIP:10.0.1.4,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/tb08efe9m19o55jbt2pp}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/tb1pu9f78qonoie5ajod"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:fXE212sEmT4=,items:[tb2kpd806s3hu3mifdtj]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tbrah6jesqlv0jaeqia5",
      "cspResourceName": "tbrah6jesqlv0jaeqia5",
      "cspResourceId": "tbrah6jesqlv0jaeqia5",
      "name": "my-ng-influxdb-back-1",
      "nodeGroupId": "my-ng-influxdb-back",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-13 09:18:59",
      "label": {
        "keypair": "tb0ffm8vtk5l3ahog7h1",
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-07-13 09:18:59",
        "sys.cspResourceId": "tbrah6jesqlv0jaeqia5",
        "sys.cspResourceName": "tbrah6jesqlv0jaeqia5",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbrah6jesqlv0jaeqia5",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "34.47.124.15",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.3",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 30,
      "RootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+e2-standard-4",
      "cspSpecName": "e2-standard-4",
      "spec": {
        "cspSpecName": "e2-standard-4",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.171931
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "tb6qc3e779ji5e5kqg7k",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "tb08efe9m19o55jbt2pp",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tb0ffm8vtk5l3ahog7h1",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBDkAdCzKgYDJPyL7tedFntUSar5CmzrXe1F9Zy/hWbQoPMUFAyu/f+p7/5nWMBstIbvt5xMQeVryRG7Gtv+OFQ=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:tTZZ/sg0gbpE3c3MGAbQPZw1sKMEpPC0b7/hMKaqnxE",
        "firstUsedAt": "2026-07-13T09:19:13Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:19:12Z",
          "completedTime": "2026-07-13T09:19:14Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbrah6jesqlv0jaeqia5 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-07-13T02:18:18.359-07:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:30,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/tbrah6jesqlv0jaeqia5,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "f76-KtKrepE="
        },
        {
          "key": "Id",
          "value": "6820223286706246581"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "hpJPIa5xQEk="
        },
        {
          "key": "Labels",
          "value": "{keypair:tb0ffm8vtk5l3ahog7h1}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-07-13T02:18:26.195-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:EaVRkxGYCT8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNYQTsR6YY0L8RJs96TMpl5gPye9R63LiKkn3OqgmFStcBfQyThBeoteg37qHvu3ZHriU8a27iZg8jfhi9OqJidbwljur/rQF4nwXLsrSj6wfagobNIo1LUSxLH0dDUbmDz4CvgC9e+d57mWyPfIkalRfoORxyxQpO1jD8TCwXXPkBLxlxRoftvZiGiqmqobThpWkrLjUQVPNohXciJJQNPJEOvo6UV/X2aECsgvr3fXA/vnZij71tE9MBRWH1cZCjyhGchsTvxaUJL8OjiCjM5VKCSmgZ3KhvGQC5QP+wDogAm/9XClt2fgNlnhS2suKpPXT7gnqxqHGRcb6cEEe/hOAmn8rsrWA0aiT1/YiPXSik3hd35mJ/4nFposCH4d7n323GnMimjvy4nfc1oJr3w9wHXp3xzoDYEDykAqml8Q50xmBW/TJeQimBg8uCbZvSGJtpACxuivu2e62LCUnlhJTfPoPJS3OUEm5NAVr02gEoWuWUvKWmzXxWzRG2kGTVI0MLEUtBQzvx+32ZawMY/Lt88GcprahddXXDEYDdVugxAXbOxKkua99YgJ6YPL7nlHtWEbNI+iB4ndWvh4vLuDdHc5LMamD82nG3wRlGbRtsy7XdYycZa5jveKIZAR8bs07vd2uK3E0op5c/TgWEuiW0d6G22whWMopulZcPAQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tbrah6jesqlv0jaeqia5"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.47.124.15,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:Gp6cH14WOLg=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/tb6qc3e779ji5e5kqg7k,networkIP:10.0.1.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/tb08efe9m19o55jbt2pp}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/tbrah6jesqlv0jaeqia5"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:7VBbY9V0wIg=,items:[tb74avigphtcuf3958vc]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tbq36hhnk3jjdfpfuu73",
      "cspResourceName": "tbq36hhnk3jjdfpfuu73",
      "cspResourceId": "tbq36hhnk3jjdfpfuu73",
      "name": "my-ng-influxdb-back-2",
      "nodeGroupId": "my-ng-influxdb-back",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-07-13 09:18:58",
      "label": {
        "keypair": "tb0ffm8vtk5l3ahog7h1",
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-07-13 09:18:58",
        "sys.cspResourceId": "tbq36hhnk3jjdfpfuu73",
        "sys.cspResourceName": "tbq36hhnk3jjdfpfuu73",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbq36hhnk3jjdfpfuu73",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=80.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "8.230.19.201",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.2",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 30,
      "RootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+e2-standard-4",
      "cspSpecName": "e2-standard-4",
      "spec": {
        "cspSpecName": "e2-standard-4",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.171931
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "tb6qc3e779ji5e5kqg7k",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "tb08efe9m19o55jbt2pp",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tb0ffm8vtk5l3ahog7h1",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBN0hMfrekkyROVSveoZGsCeo0oPeeDORUexhxHc5peKEzw1/abt2djbg20vUUqqGrbH1d15K1gpSj6AZBQqj/V8=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:EdgVwUKxu+60R9ACsX+IxKJ2XBexCbLL494NU7S4fnk",
        "firstUsedAt": "2026-07-13T09:19:12Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:19:12Z",
          "completedTime": "2026-07-13T09:19:15Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbq36hhnk3jjdfpfuu73 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-07-13T02:18:17.910-07:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:30,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/tbq36hhnk3jjdfpfuu73,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "9ZZizxjggT8="
        },
        {
          "key": "Id",
          "value": "3947188695974089654"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "hpJPIa5xQEk="
        },
        {
          "key": "Labels",
          "value": "{keypair:tb0ffm8vtk5l3ahog7h1}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-07-13T02:18:25.872-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:EaVRkxGYCT8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDNYQTsR6YY0L8RJs96TMpl5gPye9R63LiKkn3OqgmFStcBfQyThBeoteg37qHvu3ZHriU8a27iZg8jfhi9OqJidbwljur/rQF4nwXLsrSj6wfagobNIo1LUSxLH0dDUbmDz4CvgC9e+d57mWyPfIkalRfoORxyxQpO1jD8TCwXXPkBLxlxRoftvZiGiqmqobThpWkrLjUQVPNohXciJJQNPJEOvo6UV/X2aECsgvr3fXA/vnZij71tE9MBRWH1cZCjyhGchsTvxaUJL8OjiCjM5VKCSmgZ3KhvGQC5QP+wDogAm/9XClt2fgNlnhS2suKpPXT7gnqxqHGRcb6cEEe/hOAmn8rsrWA0aiT1/YiPXSik3hd35mJ/4nFposCH4d7n323GnMimjvy4nfc1oJr3w9wHXp3xzoDYEDykAqml8Q50xmBW/TJeQimBg8uCbZvSGJtpACxuivu2e62LCUnlhJTfPoPJS3OUEm5NAVr02gEoWuWUvKWmzXxWzRG2kGTVI0MLEUtBQzvx+32ZawMY/Lt88GcprahddXXDEYDdVugxAXbOxKkua99YgJ6YPL7nlHtWEbNI+iB4ndWvh4vLuDdHc5LMamD82nG3wRlGbRtsy7XdYycZa5jveKIZAR8bs07vd2uK3E0op5c/TgWEuiW0d6G22whWMopulZcPAQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tbq36hhnk3jjdfpfuu73"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:8.230.19.201,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:BG1-eHXkxAs=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/tb6qc3e779ji5e5kqg7k,networkIP:10.0.1.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/tb08efe9m19o55jbt2pp}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/tbq36hhnk3jjdfpfuu73"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:7VBbY9V0wIg=,items:[tb74avigphtcuf3958vc]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
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
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "34.47.124.15",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbrah6jesqlv0jaeqia5 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "34.64.195.188",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tb1pu9f78qonoie5ajod 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "8.230.19.201",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbq36hhnk3jjdfpfuu73 6.8.0-1060-gcp #63~22.04.1-Ubuntu SMP Wed May 27 08:12:44 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "cspResourceName": "tb1r38gq84m6n07noqfm",
      "cspResourceId": "tb1r38gq84m6n07noqfm",
      "name": "my-ng-influxdb-back",
      "connectionName": "gcp-asia-northeast3",
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
      "type": "",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "ip": "34.47.73.75",
        "port": "8086",
        "keyValueList": [
          {
            "key": "IPAddress",
            "value": "34.47.73.75"
          },
          {
            "key": "IPProtocol",
            "value": "TCP"
          },
          {
            "key": "AllPorts",
            "value": "false"
          },
          {
            "key": "AllowGlobalAccess",
            "value": "false"
          },
          {
            "key": "AllowPscGlobalAccess",
            "value": "false"
          },
          {
            "key": "CreationTimestamp",
            "value": "2026-07-13T02:20:05.554-07:00"
          },
          {
            "key": "ExternalManagedBackendBucketMigrationTestingPercentage",
            "value": "0.00"
          },
          {
            "key": "Fingerprint",
            "value": "IrMhuTszlUc="
          },
          {
            "key": "Id",
            "value": "3947952749181211434"
          },
          {
            "key": "IsMirroringCollector",
            "value": "false"
          },
          {
            "key": "Kind",
            "value": "compute#forwardingRule"
          },
          {
            "key": "LabelFingerprint",
            "value": "42WmSpB8rSM="
          },
          {
            "key": "LoadBalancingScheme",
            "value": "EXTERNAL"
          },
          {
            "key": "Name",
            "value": "tb1r38gq84m6n07noqfm"
          },
          {
            "key": "NetworkTier",
            "value": "PREMIUM"
          },
          {
            "key": "NoAutomateDnsZone",
            "value": "false"
          },
          {
            "key": "PortRange",
            "value": "8086-8086"
          },
          {
            "key": "PscConnectionId",
            "value": "0"
          },
          {
            "key": "Region",
            "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3"
          },
          {
            "key": "SelfLink",
            "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/tb1r38gq84m6n07noqfm"
          },
          {
            "key": "SelfLinkWithId",
            "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/3947952749181211434"
          },
          {
            "key": "Target",
            "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/targetPools/tb1r38gq84m6n07noqfm"
          }
        ]
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
        "protocol": "HTTP",
        "port": "8086",
        "interval": 10,
        "threshold": 3,
        "timeout": 10
      },
      "createdTime": "0001-01-01T00:00:00Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "isAutoGenerated": false,
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
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
    "cspResourceName": "tb1r38gq84m6n07noqfm",
    "cspResourceId": "tb1r38gq84m6n07noqfm",
    "name": "my-ng-influxdb-back",
    "connectionName": "gcp-asia-northeast3",
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
    "type": "",
    "scope": "REGION",
    "listener": {
      "protocol": "TCP",
      "ip": "34.47.73.75",
      "port": "8086",
      "keyValueList": [
        {
          "key": "IPAddress",
          "value": "34.47.73.75"
        },
        {
          "key": "IPProtocol",
          "value": "TCP"
        },
        {
          "key": "AllPorts",
          "value": "false"
        },
        {
          "key": "AllowGlobalAccess",
          "value": "false"
        },
        {
          "key": "AllowPscGlobalAccess",
          "value": "false"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-07-13T02:20:05.554-07:00"
        },
        {
          "key": "ExternalManagedBackendBucketMigrationTestingPercentage",
          "value": "0.00"
        },
        {
          "key": "Fingerprint",
          "value": "IrMhuTszlUc="
        },
        {
          "key": "Id",
          "value": "3947952749181211434"
        },
        {
          "key": "IsMirroringCollector",
          "value": "false"
        },
        {
          "key": "Kind",
          "value": "compute#forwardingRule"
        },
        {
          "key": "LabelFingerprint",
          "value": "42WmSpB8rSM="
        },
        {
          "key": "LoadBalancingScheme",
          "value": "EXTERNAL"
        },
        {
          "key": "Name",
          "value": "tb1r38gq84m6n07noqfm"
        },
        {
          "key": "NetworkTier",
          "value": "PREMIUM"
        },
        {
          "key": "NoAutomateDnsZone",
          "value": "false"
        },
        {
          "key": "PortRange",
          "value": "8086-8086"
        },
        {
          "key": "PscConnectionId",
          "value": "0"
        },
        {
          "key": "Region",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/tb1r38gq84m6n07noqfm"
        },
        {
          "key": "SelfLinkWithId",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/3947952749181211434"
        },
        {
          "key": "Target",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/targetPools/tb1r38gq84m6n07noqfm"
        }
      ]
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
      "protocol": "HTTP",
      "port": "8086",
      "interval": 10,
      "threshold": 3,
      "timeout": 10
    },
    "createdTime": "0001-01-01T00:00:00Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "isAutoGenerated": false,
    "location": {
      "display": "South Korea (Seoul)",
      "latitude": 37.2,
      "longitude": 127
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
  "cspResourceName": "tb1r38gq84m6n07noqfm",
  "cspResourceId": "tb1r38gq84m6n07noqfm",
  "name": "my-ng-influxdb-back",
  "connectionName": "gcp-asia-northeast3",
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
  "type": "",
  "scope": "REGION",
  "listener": {
    "protocol": "TCP",
    "ip": "34.47.73.75",
    "port": "8086",
    "keyValueList": [
      {
        "key": "IPAddress",
        "value": "34.47.73.75"
      },
      {
        "key": "IPProtocol",
        "value": "TCP"
      },
      {
        "key": "AllPorts",
        "value": "false"
      },
      {
        "key": "AllowGlobalAccess",
        "value": "false"
      },
      {
        "key": "AllowPscGlobalAccess",
        "value": "false"
      },
      {
        "key": "CreationTimestamp",
        "value": "2026-07-13T02:20:05.554-07:00"
      },
      {
        "key": "ExternalManagedBackendBucketMigrationTestingPercentage",
        "value": "0.00"
      },
      {
        "key": "Fingerprint",
        "value": "IrMhuTszlUc="
      },
      {
        "key": "Id",
        "value": "3947952749181211434"
      },
      {
        "key": "IsMirroringCollector",
        "value": "false"
      },
      {
        "key": "Kind",
        "value": "compute#forwardingRule"
      },
      {
        "key": "LabelFingerprint",
        "value": "42WmSpB8rSM="
      },
      {
        "key": "LoadBalancingScheme",
        "value": "EXTERNAL"
      },
      {
        "key": "Name",
        "value": "tb1r38gq84m6n07noqfm"
      },
      {
        "key": "NetworkTier",
        "value": "PREMIUM"
      },
      {
        "key": "NoAutomateDnsZone",
        "value": "false"
      },
      {
        "key": "PortRange",
        "value": "8086-8086"
      },
      {
        "key": "PscConnectionId",
        "value": "0"
      },
      {
        "key": "Region",
        "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3"
      },
      {
        "key": "SelfLink",
        "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/tb1r38gq84m6n07noqfm"
      },
      {
        "key": "SelfLinkWithId",
        "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/3947952749181211434"
      },
      {
        "key": "Target",
        "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/targetPools/tb1r38gq84m6n07noqfm"
      }
    ]
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
    "protocol": "HTTP",
    "port": "8086",
    "interval": 10,
    "threshold": 3,
    "timeout": 10
  },
  "createdTime": "0001-01-01T00:00:00Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "isAutoGenerated": false,
  "location": {
    "display": "South Korea (Seoul)",
    "latitude": 37.2,
    "longitude": 127
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

**Generated At:** 2026-07-13 09:22:10

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | GCP |
| **Target Region** | asia-northeast3 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| e2-highcpu-2 | 2 | 2.0 | - | x86_64 |  | $0.0635 | 1 |
| e2-standard-4 | 4 | 15.6 | - | x86_64 |  | $0.1719 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623 | Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23 | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 1 |
| https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623 | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23 | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 2 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | tb1pu9f78qonoie5ajod | Running | 2 vCPU, 2.0 GiB | Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23 (Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 34.64.195.188<br>**Private IP:** 10.0.1.4<br>**SGs:** my-mig-sg-02<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-1 | tbrah6jesqlv0jaeqia5 | Running | 4 vCPU, 15.6 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 34.47.124.15<br>**Private IP:** 10.0.1.3<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-2 | tbq36hhnk3jjdfpfuu73 | Running | 4 vCPU, 15.6 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 8.230.19.201<br>**Private IP:** 10.0.1.2<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-vnet-01 |
| **CSP VNet ID** | tb6qc3e779ji5e5kqg7k |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | gcp-asia-northeast3 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-mig-subnet-01 | tb08efe9m19o55jbt2pp | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-mig-sshkey-01 | tb0ffm8vtk5l3ahog7h1 |  |  |

### Security Groups

#### Security Group: my-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-01 |
| **CSP Security Group ID** | tb74avigphtcuf3958vc |
| **VNet** | my-mig-vnet-01 |
| **Rule Count** | 5 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-02 |
| **CSP Security Group ID** | tb2kpd806s3hu3mifdtj |
| **VNet** | my-mig-vnet-01 |
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
| **Per Hour** | $0.4074 |
| **Per Day** | $9.78 |
| **Per Month (30 days)** | $293.32 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| GCP | asia-northeast3 | 3 | $0.4074 | $293.32 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | e2-highcpu-2 | $0.0635 | $45.74 |
| my-ng-influxdb-back-1 | e2-standard-4 | $0.1719 | $123.79 |
| my-ng-influxdb-back-2 | e2-standard-4 | $0.1719 | $123.79 |




### Test Case 12: Migration Report

#### 12.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-07-13 09:22:16*

---

## 📊 Migration Summary

**Target Cloud:** GCP

**Target Region:** asia-northeast3

**Namespace:** mig01 | **Infra ID:** my-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $293.32 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | e2-standard-4, e2-highcpu-2 |
| 3 | **VM OS Image** | 2 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 9 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** tb1pu9f78qonoie5ajod<br>**Label(sourceMachineId):** ng-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my-ng-influxdb-back-1<br>**VM ID:** tbrah6jesqlv0jaeqia5<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |
| 3 | **VM Name:** my-ng-influxdb-back-2<br>**VM ID:** tbq36hhnk3jjdfpfuu73<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** e2-highcpu-2<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2 | my-ng-influxdb-back-1 | **Spec ID:** e2-standard-4<br>**vCPUs:** 4<br>**Memory:** 15.6 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3 | my-ng-influxdb-back-2 | **Spec ID:** e2-standard-4<br>**vCPUs:** 4<br>**Memory:** 15.6 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 2 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** https://www.googleapis.com/compute/v1/projects/ubuntu-os-accelerator-images/global/images/ubuntu-accel-2204-amd64-tpu-v5e-v5p-v6e-v20260623<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical, Ubuntu, 22.04 LTS TPU version(s): v5e/v5p/v6e, amd64 jammy image built on 2026-06-23 | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my-ng-influxdb-back-1 | **Image ID:** https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23 | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my-ng-influxdb-back-2 | **Image ID:** https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260623<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23 | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 9 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my-mig-sg-01

**CSP ID:** tb74avigphtcuf3958vc | **VNet:** my-mig-vnet-01 | **Rules:** 5

**Assigned VMs:**

- **VM:** my-ng-influxdb-back-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng
- **VM:** my-ng-influxdb-back-2
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 2 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | ALL |  | 10.0.0.0/16 | inbound * * from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 5 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my-mig-sg-02

**CSP ID:** tb2kpd806s3hu3mifdtj | **VNet:** my-mig-vnet-01 | **Rules:** 4

**Assigned VMs:**

- **VM:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** ng-ec268ed7-821e-9d73-e79f

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
| 1 | **Name:** my-mig-vnet-01<br>**ID:** tb6qc3e779ji5e5kqg7k | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my-mig-subnet-01<br>**ID:** tb08efe9m19o55jbt2pp | 10.0.1.0/24 | my-mig-vnet-01 |

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
| 1 | my-mig-sshkey-01 | tb0ffm8vtk5l3ahog7h1 |  | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.4074 |
| Daily | $9.78 |
| Monthly | $293.32 |
| Yearly | $3519.88 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | e2-highcpu-2 | $45.74 | 15.6% |

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

