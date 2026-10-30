# CM-Beetle test results for GCP (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with GCP cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.2+ (83efe14)
- imdl: v0.1.19+ (83efe14)
- CB-Tumblebug: v0.13.8
- CB-Spider: v0.13.20
- CB-MapUI: v0.13.14
- Target CSP: GCP
- Target Region: asia-northeast3
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: October 30, 2026
- Test Time: 21:46:24 KST
- Test Execution: 2026-10-30 21:46:24 KST

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

## Test result for GCP

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 948ms | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 123ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 2m53.899s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 30ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 3ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 9ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 2.751s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 27.848s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 4ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 4ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 1m36.738s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.094s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.099s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 41.829s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 4m8.605s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 10m53.310353439s

*Test executed on October 30, 2026 at 21:46:24 KST (2026-10-30 21:46:24 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=83.3% Max=100.0% Avg=92.0% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
              "nlbBackend": "influxdb_back"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+n4-standard-4",
            "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=87.8% Image=83.3%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+n1-highcpu-2",
            "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
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
          "id": "gcp+asia-northeast3+n4-standard-4",
          "uid": "tbtg35suro3tjfhfk2at",
          "cspSpecName": "n4-standard-4",
          "name": "gcp+asia-northeast3+n4-standard-4",
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
          "costPerHour": 0.232954,
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
              "key": "Architecture",
              "value": "X86_64"
            },
            {
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "4 vCPUs, 16 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "4"
            },
            {
              "key": "Id",
              "value": "1750002"
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
              "value": "n4-standard-4"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/seokho-etri/zones/asia-northeast3-a/machineTypes/n4-standard-4"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "asia-northeast3-a",
            "asia-northeast3-b",
            "asia-northeast3-c"
          ],
          "availabilityCheckedAt": "2026-10-30T12:46:30.583180225Z",
          "availabilityRefreshAfter": "2026-10-30T13:46:30.583180225Z"
        },
        {
          "id": "gcp+asia-northeast3+n1-highcpu-2",
          "uid": "tbuaoej0i7lm24gthk7e",
          "cspSpecName": "n1-highcpu-2",
          "name": "gcp+asia-northeast3+n1-highcpu-2",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1.7568359,
          "diskSizeGB": -1,
          "costPerHour": 0.09099,
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
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "2 vCPUs, 1.8 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "2"
            },
            {
              "key": "Id",
              "value": "4002"
            },
            {
              "key": "ImageSpaceGb",
              "value": "10"
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
              "value": "1843"
            },
            {
              "key": "Name",
              "value": "n1-highcpu-2"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/seokho-etri/zones/asia-northeast3-a/machineTypes/n1-highcpu-2"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "asia-northeast3-a",
            "asia-northeast3-b",
            "asia-northeast3-c"
          ],
          "availabilityCheckedAt": "2026-10-30T12:46:30.610224369Z",
          "availabilityRefreshAfter": "2026-10-30T13:46:30.610224369Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "gcp",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "regionList": [
            "common"
          ],
          "id": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "uid": "tbta9p6tv2cjgah4d9rr",
          "name": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "gcp-africa-south1",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:58:59 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20",
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
              "value": "63898864000"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-08-20T06:17:20.257-07:00"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20"
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
              "value": "6494416155768996016"
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
              "value": "ubuntu-2204-jammy-v20260820"
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
              "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820"
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
              "value": "asia; europe-west8; us-east7; northamerica-south1; northamerica-northeast2; asia-east1; asia-southeast3; southamerica-east1; europe-west2; europe-north1; us-west3; europe-west15; us-central2; asia-northeast1; asia-southeast2; me-central2; europe-west4; us-east5; us-west8; northamerica-northeast1; us-west4; us-central1; southamerica-west1; asia-southeast1; me-central1; australia-southeast1; europe-central2; asia-northeast3; australia-southeast2; asia-south1; us-east4; us-south1; africa-south1; us; europe-west6; europe-north2; me-west1; us-west1; europe-west9; europe-west1; asia-northeast2; us-east1; europe-west3; asia-south2; europe-west10; europe-southwest1; europe-west12; eu; asia-east2; us-west2"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20",
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=83.3% Max=100.0% Avg=93.7% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
              "nlbBackend": "influxdb_back"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+n2-standard-4",
            "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+n2-highcpu-2",
            "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
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
          "id": "gcp+asia-northeast3+n2-standard-4",
          "uid": "tb2krgi3rsv0p6pfoeeu",
          "cspSpecName": "n2-standard-4",
          "name": "gcp+asia-northeast3+n2-standard-4",
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
          "costPerHour": 0.249436,
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
              "value": "901004"
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
              "value": "n2-standard-4"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/seokho-etri/zones/asia-northeast3-a/machineTypes/n2-standard-4"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ],
          "availabilityStatus": "Unverified"
        },
        {
          "id": "gcp+asia-northeast3+n2-highcpu-2",
          "uid": "tbr0pd6ifdsag3bm25g0",
          "cspSpecName": "n2-highcpu-2",
          "name": "gcp+asia-northeast3+n2-highcpu-2",
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
          "costPerHour": 0.092072,
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
              "value": "903002"
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
              "value": "n2-highcpu-2"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/seokho-etri/zones/asia-northeast3-a/machineTypes/n2-highcpu-2"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "asia-northeast3-a",
            "asia-northeast3-b",
            "asia-northeast3-c"
          ],
          "availabilityCheckedAt": "2026-10-30T12:46:30.610286249Z",
          "availabilityRefreshAfter": "2026-10-30T13:46:30.610286249Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "gcp",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "regionList": [
            "common"
          ],
          "id": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "uid": "tbta9p6tv2cjgah4d9rr",
          "name": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "gcp-africa-south1",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:58:59 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20",
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
              "value": "63898864000"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-08-20T06:17:20.257-07:00"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20"
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
              "value": "6494416155768996016"
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
              "value": "ubuntu-2204-jammy-v20260820"
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
              "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820"
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
              "value": "asia; europe-west8; us-east7; northamerica-south1; northamerica-northeast2; asia-east1; asia-southeast3; southamerica-east1; europe-west2; europe-north1; us-west3; europe-west15; us-central2; asia-northeast1; asia-southeast2; me-central2; europe-west4; us-east5; us-west8; northamerica-northeast1; us-west4; us-central1; southamerica-west1; asia-southeast1; me-central1; australia-southeast1; europe-central2; asia-northeast3; australia-southeast2; asia-south1; us-east4; us-south1; africa-south1; us; europe-west6; europe-north2; me-west1; us-west1; europe-west9; europe-west1; asia-northeast2; us-east1; europe-west3; asia-south2; europe-west10; europe-southwest1; europe-west12; eu; asia-east2; us-west2"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20",
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
  "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=83.3% Max=100.0% Avg=92.0% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
  "targetCloud": {
    "csp": "gcp",
    "region": "asia-northeast3"
  },
  "targetInfra": {
    "name": "my03-infra101",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "NLB-aware recommended infrastructure for cloud migration",
    "nodeGroups": [
      {
        "name": "my03-ng-influxdb-back",
        "nodeGroupSize": 2,
        "label": {
          "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "nlbBackend": "influxdb_back"
        },
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
        "connectionName": "gcp-asia-northeast3",
        "specId": "gcp+asia-northeast3+n4-standard-4",
        "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
        "vNetId": "my03-mig-vnet-01",
        "subnetId": "my03-mig-subnet-01",
        "securityGroupIds": [
          "my03-mig-sg-01"
        ],
        "sshKeyId": "my03-mig-sshkey-01",
        "rootDiskSize": 30,
        "dataDiskIds": null
      },
      {
        "name": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=87.8% Image=83.3%",
        "connectionName": "gcp-asia-northeast3",
        "specId": "gcp+asia-northeast3+n1-highcpu-2",
        "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
        "vNetId": "my03-mig-vnet-01",
        "subnetId": "my03-mig-subnet-01",
        "securityGroupIds": [
          "my03-mig-sg-02"
        ],
        "sshKeyId": "my03-mig-sshkey-01",
        "rootDiskSize": 10,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "my03-mig-vnet-01",
    "connectionName": "gcp-asia-northeast3",
    "cidrBlock": "10.0.0.0/21",
    "subnetInfoList": [
      {
        "name": "my03-mig-subnet-01",
        "ipv4_CIDR": "10.0.1.0/24",
        "description": "a recommended subnet for migration"
      }
    ],
    "description": "a recommended vNet for migration"
  },
  "targetSshKey": {
    "name": "my03-mig-sshkey-01",
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
      "id": "gcp+asia-northeast3+n4-standard-4",
      "uid": "tbtg35suro3tjfhfk2at",
      "cspSpecName": "n4-standard-4",
      "name": "gcp+asia-northeast3+n4-standard-4",
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
      "costPerHour": 0.232954,
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
          "key": "Architecture",
          "value": "X86_64"
        },
        {
          "key": "CreationTimestamp",
          "value": "1969-12-31T16:00:00.000-08:00"
        },
        {
          "key": "Description",
          "value": "4 vCPUs, 16 GB RAM"
        },
        {
          "key": "GuestCpus",
          "value": "4"
        },
        {
          "key": "Id",
          "value": "1750002"
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
          "value": "n4-standard-4"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/seokho-etri/zones/asia-northeast3-a/machineTypes/n4-standard-4"
        },
        {
          "key": "Zone",
          "value": "asia-northeast3-a"
        }
      ],
      "availabilityStatus": "Available",
      "availableZones": [
        "asia-northeast3-a",
        "asia-northeast3-b",
        "asia-northeast3-c"
      ],
      "availabilityCheckedAt": "2026-10-30T12:46:30.583180225Z",
      "availabilityRefreshAfter": "2026-10-30T13:46:30.583180225Z"
    },
    {
      "id": "gcp+asia-northeast3+n1-highcpu-2",
      "uid": "tbuaoej0i7lm24gthk7e",
      "cspSpecName": "n1-highcpu-2",
      "name": "gcp+asia-northeast3+n1-highcpu-2",
      "namespace": "system",
      "connectionName": "gcp-asia-northeast3",
      "providerName": "gcp",
      "regionName": "asia-northeast3",
      "regionLatitude": 37.2,
      "regionLongitude": 127,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 1.7568359,
      "diskSizeGB": -1,
      "costPerHour": 0.09099,
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
          "key": "CreationTimestamp",
          "value": "1969-12-31T16:00:00.000-08:00"
        },
        {
          "key": "Description",
          "value": "2 vCPUs, 1.8 GB RAM"
        },
        {
          "key": "GuestCpus",
          "value": "2"
        },
        {
          "key": "Id",
          "value": "4002"
        },
        {
          "key": "ImageSpaceGb",
          "value": "10"
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
          "value": "1843"
        },
        {
          "key": "Name",
          "value": "n1-highcpu-2"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/seokho-etri/zones/asia-northeast3-a/machineTypes/n1-highcpu-2"
        },
        {
          "key": "Zone",
          "value": "asia-northeast3-a"
        }
      ],
      "availabilityStatus": "Available",
      "availableZones": [
        "asia-northeast3-a",
        "asia-northeast3-b",
        "asia-northeast3-c"
      ],
      "availabilityCheckedAt": "2026-10-30T12:46:30.610224369Z",
      "availabilityRefreshAfter": "2026-10-30T13:46:30.610224369Z"
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "gcp",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
      "regionList": [
        "common"
      ],
      "id": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
      "uid": "tbta9p6tv2cjgah4d9rr",
      "name": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "gcp-africa-south1",
      "infraType": "",
      "fetchedTime": "2026.08.21 13:58:59 Fri",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20",
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
          "value": "63898864000"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-08-20T06:17:20.257-07:00"
        },
        {
          "key": "Description",
          "value": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20"
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
          "value": "6494416155768996016"
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
          "value": "ubuntu-2204-jammy-v20260820"
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
          "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820"
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
          "value": "asia; europe-west8; us-east7; northamerica-south1; northamerica-northeast2; asia-east1; asia-southeast3; southamerica-east1; europe-west2; europe-north1; us-west3; europe-west15; us-central2; asia-northeast1; asia-southeast2; me-central2; europe-west4; us-east5; us-west8; northamerica-northeast1; us-west4; us-central1; southamerica-west1; asia-southeast1; me-central1; australia-southeast1; europe-central2; asia-northeast3; australia-southeast2; asia-south1; us-east4; us-south1; africa-south1; us; europe-west6; europe-north2; me-west1; us-west1; europe-west9; europe-west1; asia-northeast2; us-east1; europe-west3; asia-south2; europe-west10; europe-southwest1; europe-west12; eu; asia-east2; us-west2"
        }
      ],
      "systemLabel": "",
      "description": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20",
      "commandHistory": null
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "my03-mig-sg-01",
      "connectionName": "gcp-asia-northeast3",
      "vNetId": "my03-mig-vnet-01",
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
      "name": "my03-mig-sg-02",
      "connectionName": "gcp-asia-northeast3",
      "vNetId": "my03-mig-vnet-01",
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
- **MCI ID**: `my03-infra101`
- **MCI Name**: `my03-infra101`
- **Status**: `Running:3 (R:3/3)`

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
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
  "label": {
    "cm-migrated-at": "2026-10-30T12:48:39Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my03-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my03-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbdcanpcp70pd7mapdn4"
  },
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:49:01",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
        "keypair": "tb2h282hf7m5h9nubvn7"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=87.8% Image=83.3%",
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
      "rootDeviceName": "persistent-disk-0",
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
      "cspSshKeyId": "tb2h282hf7m5h9nubvn7",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAwh+aT5cfKTYj0GomN/rpYSvMvrPbt20o9n6OY/9IyiQeaYR1ilD9K4+HZEh2Yt+j20f+8E8QZ0W8PdmOIKYcE=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:cwO67T6MerUdC9KrYQh7vo9LyA/G/ju4tQI5mSfRGmM",
        "firstUsedAt": "2026-10-30T12:49:23Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:02Z",
          "completedTime": "2026-10-30T12:49:25Z",
          "elapsedTime": 23,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my03-infra101-tbgjiopqash1t69uo2aq",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:26Z",
          "completedTime": "2026-10-30T12:49:28Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb0ogg2jdctsppoih0d1 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:49:00",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "keypair": "tb2h282hf7m5h9nubvn7",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
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
      "rootDeviceName": "persistent-disk-0",
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
      "cspSshKeyId": "tb2h282hf7m5h9nubvn7",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBF2xsj1nW4aopZ/SjpQmSIvuYTIi8B1DutOHNMrAUqZteky7+N22rOEPTrUdCAgCnqF8+FimiHirXgYaq1rpB+o=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:Cde4DUYUxNrBVw1DuRHrfPi5xQ2C3zV30M52hlBZTjM",
        "firstUsedAt": "2026-10-30T12:49:14Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:02Z",
          "completedTime": "2026-10-30T12:49:26Z",
          "elapsedTime": 24,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my03-infra101-tbgjiopqash1t69uo2aq",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:26Z",
          "completedTime": "2026-10-30T12:49:29Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb0i6hj6t9k5dgoo7360 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:49:02",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "keypair": "tb2h282hf7m5h9nubvn7",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
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
      "rootDeviceName": "persistent-disk-0",
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
      "cspSshKeyId": "tb2h282hf7m5h9nubvn7",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBWzWhxTMlXcJoxa1vh+z0Eg6QCFwvXn9pBPGc34dcTB3FMy4IrcQ05T4YrnM1n1cCZSseOC57zVImBXoEneMR0=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:0YuQiiQGeJHli6MCpkLxhSsvXKfZLoSpcOvZZtiOsrY",
        "firstUsedAt": "2026-10-30T12:49:23Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:02Z",
          "completedTime": "2026-10-30T12:49:24Z",
          "elapsedTime": 22,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my03-infra101-tbgjiopqash1t69uo2aq",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:26Z",
          "completedTime": "2026-10-30T12:49:28Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb4ceb9m4usomf7ac0m5 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my03-mig-vnet-01",
      "name": "my03-mig-vnet-01",
      "infraId": "my03-infra101",
      "vNetId": "my03-mig-vnet-01",
      "connectionNames": [
        "gcp-asia-northeast3"
      ],
      "providerNames": [
        "gcp"
      ],
      "regionNames": [
        "asia-northeast3"
      ],
      "nodeGroupIds": [
        "my03-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my03-ng-influxdb-back"
      ],
      "nodeIds": [
        "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my03-ng-influxdb-back-1",
        "my03-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my03-infra101",
            "nodeId": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "34.47.115.170",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb0ogg2jdctsppoih0d1 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my03-infra101",
            "nodeId": "my03-ng-influxdb-back-2",
            "nodeIp": "34.50.27.124",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb4ceb9m4usomf7ac0m5 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my03-infra101",
            "nodeId": "my03-ng-influxdb-back-1",
            "nodeIp": "34.64.190.160",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb0i6hj6t9k5dgoo7360 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my03-infra101-tbgjiopqash1t69uo2aq"
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
      "label": {
        "cm-migrated-at": "2026-10-30T12:48:39Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my03-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my03-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbdcanpcp70pd7mapdn4"
      },
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
          "id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbar0eeronn7me24e1jr",
          "cspResourceName": "tbar0eeronn7me24e1jr",
          "cspResourceId": "02h7_8466ac07-73a0-45a7-947e-f6025a81d369",
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
          "createdTime": "2026-10-30 12:49:27",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-10-30 12:49:27",
            "sys.cspResourceId": "02h7_8466ac07-73a0-45a7-947e-f6025a81d369",
            "sys.cspResourceName": "tbar0eeronn7me24e1jr",
            "sys.id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbar0eeronn7me24e1jr",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.100.42",
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
          "cspVNetId": "r026-2f361af0-73d0-46a0-9d2b-b8da85809476",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-cc88798d-0035-4f70-bc4d-80b630a00c4c",
          "networkInterface": "cannon-unending-glancing-clumsily",
          "securityGroupIds": [
            "my06-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-f6cfb962-7eff-4047-8665-12fe4e8077bf"
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-1",
          "uid": "tbf0cqgntsded5k3h3ln",
          "cspResourceName": "tbf0cqgntsded5k3h3ln",
          "cspResourceId": "02h7_3957719b-2881-4faa-a682-8c6f5569a4f5",
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
          "createdTime": "2026-10-30 12:49:28",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-10-30 12:49:28",
            "sys.cspResourceId": "02h7_3957719b-2881-4faa-a682-8c6f5569a4f5",
            "sys.cspResourceName": "tbf0cqgntsded5k3h3ln",
            "sys.id": "my06-ng-influxdb-back-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-influxdb-back",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbf0cqgntsded5k3h3ln",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.90.173",
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
          "cspVNetId": "r026-2f361af0-73d0-46a0-9d2b-b8da85809476",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-cc88798d-0035-4f70-bc4d-80b630a00c4c",
          "networkInterface": "motivate-shape-slinky-turbulent",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-f6cfb962-7eff-4047-8665-12fe4e8077bf"
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-2",
          "uid": "tbvicb2g6blfup7t2rn1",
          "cspResourceName": "tbvicb2g6blfup7t2rn1",
          "cspResourceId": "02h7_72a79220-56f5-411a-a2f8-073d38e9848a",
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
          "createdTime": "2026-10-30 12:49:28",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-10-30 12:49:28",
            "sys.cspResourceId": "02h7_72a79220-56f5-411a-a2f8-073d38e9848a",
            "sys.cspResourceName": "tbvicb2g6blfup7t2rn1",
            "sys.id": "my06-ng-influxdb-back-2",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-influxdb-back",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbvicb2g6blfup7t2rn1",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.97.164",
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
          "cspVNetId": "r026-2f361af0-73d0-46a0-9d2b-b8da85809476",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-cc88798d-0035-4f70-bc4d-80b630a00c4c",
          "networkInterface": "axiom-groovy-haven-isolated",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-f6cfb962-7eff-4047-8665-12fe4e8077bf"
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
  "label": {
    "cm-migrated-at": "2026-10-30T12:48:39Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my03-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my03-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbdcanpcp70pd7mapdn4"
  },
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:49:01",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
        "keypair": "tb2h282hf7m5h9nubvn7"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=87.8% Image=83.3%",
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
      "rootDeviceName": "persistent-disk-0",
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
      "cspSshKeyId": "tb2h282hf7m5h9nubvn7",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAwh+aT5cfKTYj0GomN/rpYSvMvrPbt20o9n6OY/9IyiQeaYR1ilD9K4+HZEh2Yt+j20f+8E8QZ0W8PdmOIKYcE=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:cwO67T6MerUdC9KrYQh7vo9LyA/G/ju4tQI5mSfRGmM",
        "firstUsedAt": "2026-10-30T12:49:23Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:02Z",
          "completedTime": "2026-10-30T12:49:25Z",
          "elapsedTime": 23,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my03-infra101-tbgjiopqash1t69uo2aq",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:26Z",
          "completedTime": "2026-10-30T12:49:28Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb0ogg2jdctsppoih0d1 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:49:00",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "keypair": "tb2h282hf7m5h9nubvn7",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
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
      "rootDeviceName": "persistent-disk-0",
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
      "cspSshKeyId": "tb2h282hf7m5h9nubvn7",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBF2xsj1nW4aopZ/SjpQmSIvuYTIi8B1DutOHNMrAUqZteky7+N22rOEPTrUdCAgCnqF8+FimiHirXgYaq1rpB+o=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:Cde4DUYUxNrBVw1DuRHrfPi5xQ2C3zV30M52hlBZTjM",
        "firstUsedAt": "2026-10-30T12:49:14Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:02Z",
          "completedTime": "2026-10-30T12:49:26Z",
          "elapsedTime": 24,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my03-infra101-tbgjiopqash1t69uo2aq",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:26Z",
          "completedTime": "2026-10-30T12:49:29Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb0i6hj6t9k5dgoo7360 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:49:02",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "keypair": "tb2h282hf7m5h9nubvn7",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=97.7% Image=83.3%",
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
      "rootDeviceName": "persistent-disk-0",
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
      "cspSshKeyId": "tb2h282hf7m5h9nubvn7",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBWzWhxTMlXcJoxa1vh+z0Eg6QCFwvXn9pBPGc34dcTB3FMy4IrcQ05T4YrnM1n1cCZSseOC57zVImBXoEneMR0=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:0YuQiiQGeJHli6MCpkLxhSsvXKfZLoSpcOvZZtiOsrY",
        "firstUsedAt": "2026-10-30T12:49:23Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:02Z",
          "completedTime": "2026-10-30T12:49:24Z",
          "elapsedTime": 22,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my03-infra101-tbgjiopqash1t69uo2aq",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:49:26Z",
          "completedTime": "2026-10-30T12:49:28Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb4ceb9m4usomf7ac0m5 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my03-mig-vnet-01",
      "name": "my03-mig-vnet-01",
      "infraId": "my03-infra101",
      "vNetId": "my03-mig-vnet-01",
      "connectionNames": [
        "gcp-asia-northeast3"
      ],
      "providerNames": [
        "gcp"
      ],
      "regionNames": [
        "asia-northeast3"
      ],
      "nodeGroupIds": [
        "my03-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my03-ng-influxdb-back"
      ],
      "nodeIds": [
        "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my03-ng-influxdb-back-1",
        "my03-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my03-infra101",
            "nodeId": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "34.47.115.170",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb0ogg2jdctsppoih0d1 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my03-infra101",
            "nodeId": "my03-ng-influxdb-back-2",
            "nodeIp": "34.50.27.124",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb4ceb9m4usomf7ac0m5 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my03-infra101",
            "nodeId": "my03-ng-influxdb-back-1",
            "nodeIp": "34.64.190.160",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb0i6hj6t9k5dgoo7360 6.8.0-1066-gcp #74~22.04.1-Ubuntu SMP Fri Aug  7 21:51:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my03-infra101-tbgjiopqash1t69uo2aq"
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
      "id": "my03-ng-influxdb-back",
      "cspResourceName": "tbk08q2hn14apj9njl1l",
      "cspResourceId": "tbk08q2hn14apj9njl1l",
      "name": "my03-ng-influxdb-back",
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
      "type": "",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "ip": "34.50.23.184",
        "port": "8086",
        "keyValueList": [
          {
            "key": "IPAddress",
            "value": "34.50.23.184"
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
            "value": "2026-10-30T05:49:32.957-07:00"
          },
          {
            "key": "ExternalManagedBackendBucketMigrationTestingPercentage",
            "value": "0.00"
          },
          {
            "key": "Fingerprint",
            "value": "pE_Fo108I38="
          },
          {
            "key": "Id",
            "value": "8639558531198561459"
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
            "value": "tbk08q2hn14apj9njl1l"
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
            "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/tbk08q2hn14apj9njl1l"
          },
          {
            "key": "SelfLinkWithId",
            "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/8639558531198561459"
          },
          {
            "key": "Target",
            "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/targetPools/tbk08q2hn14apj9njl1l"
          }
        ]
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my03-ng-influxdb-back",
        "nodes": [
          "my03-ng-influxdb-back-1",
          "my03-ng-influxdb-back-2"
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
    "id": "my03-ng-influxdb-back",
    "cspResourceName": "tbk08q2hn14apj9njl1l",
    "cspResourceId": "tbk08q2hn14apj9njl1l",
    "name": "my03-ng-influxdb-back",
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
    "type": "",
    "scope": "REGION",
    "listener": {
      "protocol": "TCP",
      "ip": "34.50.23.184",
      "port": "8086",
      "keyValueList": [
        {
          "key": "IPAddress",
          "value": "34.50.23.184"
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
          "value": "2026-10-30T05:49:32.957-07:00"
        },
        {
          "key": "ExternalManagedBackendBucketMigrationTestingPercentage",
          "value": "0.00"
        },
        {
          "key": "Fingerprint",
          "value": "pE_Fo108I38="
        },
        {
          "key": "Id",
          "value": "8639558531198561459"
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
          "value": "tbk08q2hn14apj9njl1l"
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
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/tbk08q2hn14apj9njl1l"
        },
        {
          "key": "SelfLinkWithId",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/8639558531198561459"
        },
        {
          "key": "Target",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/targetPools/tbk08q2hn14apj9njl1l"
        }
      ]
    },
    "targetGroup": {
      "protocol": "TCP",
      "port": "8086",
      "nodeGroupId": "my03-ng-influxdb-back",
      "nodes": [
        "my03-ng-influxdb-back-1",
        "my03-ng-influxdb-back-2"
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
  "id": "my03-ng-influxdb-back",
  "cspResourceName": "tbk08q2hn14apj9njl1l",
  "cspResourceId": "tbk08q2hn14apj9njl1l",
  "name": "my03-ng-influxdb-back",
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
  "type": "",
  "scope": "REGION",
  "listener": {
    "protocol": "TCP",
    "ip": "34.50.23.184",
    "port": "8086",
    "keyValueList": [
      {
        "key": "IPAddress",
        "value": "34.50.23.184"
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
        "value": "2026-10-30T05:49:32.957-07:00"
      },
      {
        "key": "ExternalManagedBackendBucketMigrationTestingPercentage",
        "value": "0.00"
      },
      {
        "key": "Fingerprint",
        "value": "pE_Fo108I38="
      },
      {
        "key": "Id",
        "value": "8639558531198561459"
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
        "value": "tbk08q2hn14apj9njl1l"
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
        "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/tbk08q2hn14apj9njl1l"
      },
      {
        "key": "SelfLinkWithId",
        "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/forwardingRules/8639558531198561459"
      },
      {
        "key": "Target",
        "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/targetPools/tbk08q2hn14apj9njl1l"
      }
    ]
  },
  "targetGroup": {
    "protocol": "TCP",
    "port": "8086",
    "nodeGroupId": "my03-ng-influxdb-back",
    "nodes": [
      "my03-ng-influxdb-back-1",
      "my03-ng-influxdb-back-2"
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

**Generated At:** 2026-10-30 12:52:22

**Namespace:** mig01

**Infra Name:** my03-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my03-infra101 |
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
| n1-highcpu-2 | 2 | 1.8 | - | x86_64 |  | $0.0910 | 1 |
| n4-standard-4 | 4 | 15.6 | - | x86_64 |  | $0.2330 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20260820 | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20 | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | tb0ogg2jdctsppoih0d1 | Running | 2 vCPU, 1.8 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20) | **VNet:** my03-mig-vnet-01<br>**Subnet:** my03-mig-subnet-01<br>**Public IP:** 34.47.115.170<br>**Private IP:** 10.0.1.3<br>**SGs:** my03-mig-sg-02<br>**SSH:** my03-mig-sshkey-01 |
| my03-ng-influxdb-back-1 | tb0i6hj6t9k5dgoo7360 | Running | 4 vCPU, 15.6 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20) | **VNet:** my03-mig-vnet-01<br>**Subnet:** my03-mig-subnet-01<br>**Public IP:** 34.64.190.160<br>**Private IP:** 10.0.1.2<br>**SGs:** my03-mig-sg-01<br>**SSH:** my03-mig-sshkey-01 |
| my03-ng-influxdb-back-2 | tb4ceb9m4usomf7ac0m5 | Running | 4 vCPU, 15.6 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-08-20) | **VNet:** my03-mig-vnet-01<br>**Subnet:** my03-mig-subnet-01<br>**Public IP:** 34.50.27.124<br>**Private IP:** 10.0.1.4<br>**SGs:** my03-mig-sg-01<br>**SSH:** my03-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my03-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my03-mig-vnet-01 |
| **CSP VNet ID** | tbplbufb5140g64c4e6a |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | gcp-asia-northeast3 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my03-mig-subnet-01 | tb19rvmuqe5cd87gv5eh | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my03-mig-sshkey-01 | tb2h282hf7m5h9nubvn7 |  |  |

### Security Groups

#### Security Group: my03-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my03-mig-sg-01 |
| **CSP Security Group ID** | tbl1o935e0dico40og7k |
| **VNet** | my03-mig-vnet-01 |
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

#### Security Group: my03-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my03-mig-sg-02 |
| **CSP Security Group ID** | tbreh130bhrk8ekqh6a8 |
| **VNet** | my03-mig-vnet-01 |
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
| **Per Hour** | $0.5569 |
| **Per Day** | $13.37 |
| **Per Month (30 days)** | $400.97 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| GCP | asia-northeast3 | 3 | $0.5569 | $400.97 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | n1-highcpu-2 | $0.0910 | $65.51 |
| my03-ng-influxdb-back-1 | n4-standard-4 | $0.2330 | $167.73 |
| my03-ng-influxdb-back-2 | n4-standard-4 | $0.2330 | $167.73 |




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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my03-infra101)",
  "success": true
}
```

