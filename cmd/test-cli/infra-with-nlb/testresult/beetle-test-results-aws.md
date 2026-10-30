# CM-Beetle test results for AWS (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AWS cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.2+ (83efe14)
- imdl: v0.1.19+ (83efe14)
- CB-Tumblebug: v0.13.8
- CB-Spider: v0.13.20
- CB-MapUI: v0.13.14
- Target CSP: AWS
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: October 30, 2026
- Test Time: 21:45:24 KST
- Test Execution: 2026-10-30 21:45:24 KST

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

## Test result for AWS

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 2.978s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 140ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m2.058s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 12ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 3ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 11ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 3.047s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 2.184s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 5ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 5ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 3m33.228s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.099s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.099s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 16.193s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 59.264s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 6m59.613082555s

*Test executed on October 30, 2026 at 21:45:24 KST (2026-10-30 21:45:24 KST) using CM-Beetle automated test CLI*

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
  "desiredCsp": "aws",
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
      "status": "highly-matched",
      "description": "Candidate #1 | highly-matched | 1 NLB(s) | Overall Match Rate: Min=100.0% Max=100.0% Avg=100.0% | VMs: 2 total, 2 matched, 0 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "aws",
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3.xlarge",
            "imageId": "ami-012a353bb3afb92ee",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3.small",
            "imageId": "ami-012a353bb3afb92ee",
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
        "connectionName": "aws-ap-northeast-2",
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
        "connectionName": "aws-ap-northeast-2",
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
          "id": "aws+ap-northeast-2+t3.xlarge",
          "uid": "tbt1via2qor5okve6pun",
          "cspSpecName": "t3.xlarge",
          "name": "aws+ap-northeast-2+t3.xlarge",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.208,
          "orderInFilteredResult": 17,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "true"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "false"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3.xlarge"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:16384}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:15,Ipv6AddressesPerInterface:15,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:4,NetworkCards:[{MaximumNetworkInterfaces:4,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.5}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:2,DefaultThreadsPerCore:2,DefaultVCpus:4,ValidCores:[2],ValidThreadsPerCore:[1,2]}"
            }
          ],
          "availabilityStatus": "Unverified"
        },
        {
          "id": "aws+ap-northeast-2+t3.small",
          "uid": "tbpf96u7vfjaf20v9o5q",
          "cspSpecName": "t3.small",
          "name": "aws+ap-northeast-2+t3.small",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.026,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "true"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:174,BaselineIops:1000,BaselineThroughputInMBps:21.75,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "true"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3.small"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:2048}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:4,Ipv6AddressesPerInterface:4,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.5}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "ap-northeast-2a",
            "ap-northeast-2b",
            "ap-northeast-2c",
            "ap-northeast-2d"
          ],
          "availabilityCheckedAt": "2026-10-30T12:45:32.100139859Z",
          "availabilityRefreshAfter": "2026-10-30T13:45:32.100139859Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "aws",
          "cspImageName": "ami-012a353bb3afb92ee",
          "regionList": [
            "ap-northeast-2"
          ],
          "id": "ami-012a353bb3afb92ee",
          "uid": "tbbe6d8vhrqe5l5nne2u",
          "name": "ami-012a353bb3afb92ee",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "aws-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.08.21 14:20:39 Fri",
          "creationDate": "2026-07-31T07:25:12.000Z",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731",
          "osDiskType": "ebs",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "BlockDeviceMappings",
              "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-047f4bba8e035b38e,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CreationDate",
              "value": "2026-07-31T07:25:12.000Z"
            },
            {
              "key": "DeprecationTime",
              "value": "2028-07-31T07:25:12.000Z"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04, amd64 jammy image"
            },
            {
              "key": "EnaSupport",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "ImageId",
              "value": "ami-012a353bb3afb92ee"
            },
            {
              "key": "ImageLocation",
              "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731"
            },
            {
              "key": "ImageOwnerAlias",
              "value": "amazon"
            },
            {
              "key": "ImageType",
              "value": "machine"
            },
            {
              "key": "Name",
              "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731"
            },
            {
              "key": "OwnerId",
              "value": "099720109477"
            },
            {
              "key": "PlatformDetails",
              "value": "Linux/UNIX"
            },
            {
              "key": "Public",
              "value": "true"
            },
            {
              "key": "RootDeviceName",
              "value": "/dev/sda1"
            },
            {
              "key": "RootDeviceType",
              "value": "ebs"
            },
            {
              "key": "SriovNetSupport",
              "value": "simple"
            },
            {
              "key": "State",
              "value": "available"
            },
            {
              "key": "UsageOperation",
              "value": "RunInstances"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04, amd64 jammy image",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "aws-ap-northeast-2",
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
          "connectionName": "aws-ap-northeast-2",
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=91.7% | VMs: 2 total, 1 matched, 1 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "aws",
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t2.xlarge",
            "imageId": "ami-012a353bb3afb92ee",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3.micro",
            "imageId": "ami-012a353bb3afb92ee",
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
        "connectionName": "aws-ap-northeast-2",
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
        "connectionName": "aws-ap-northeast-2",
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
          "id": "aws+ap-northeast-2+t2.xlarge",
          "uid": "tbvn7578t2f2cik2qsem",
          "cspSpecName": "t2.xlarge",
          "name": "aws+ap-northeast-2+t2.xlarge",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.2304,
          "orderInFilteredResult": 21,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "false"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:null,EbsOptimizedSupport:unsupported,EncryptionSupport:supported,NvmeSupport:unsupported}"
            },
            {
              "key": "FreeTierEligible",
              "value": "false"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t2.xlarge"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:16384}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:unsupported,Ipv4AddressesPerInterface:15,Ipv6AddressesPerInterface:15,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Moderate}],NetworkPerformance:Moderate}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.3}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:4,DefaultThreadsPerCore:1,DefaultVCpus:4,ValidCores:null,ValidThreadsPerCore:null}"
            }
          ],
          "availabilityStatus": "Unverified"
        },
        {
          "id": "aws+ap-northeast-2+t3.micro",
          "uid": "tb73efokamcsgeqateu1",
          "cspSpecName": "t3.micro",
          "name": "aws+ap-northeast-2+t3.micro",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1,
          "diskSizeGB": -1,
          "costPerHour": 0.013,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "true"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:87,BaselineIops:500,BaselineThroughputInMBps:10.875,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "true"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3.micro"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:1024}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:2,Ipv6AddressesPerInterface:2,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:2,NetworkCards:[{MaximumNetworkInterfaces:2,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.5}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "ap-northeast-2a",
            "ap-northeast-2b",
            "ap-northeast-2c",
            "ap-northeast-2d"
          ],
          "availabilityCheckedAt": "2026-10-30T12:45:32.100139878Z",
          "availabilityRefreshAfter": "2026-10-30T13:45:32.100139878Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "aws",
          "cspImageName": "ami-012a353bb3afb92ee",
          "regionList": [
            "ap-northeast-2"
          ],
          "id": "ami-012a353bb3afb92ee",
          "uid": "tbbe6d8vhrqe5l5nne2u",
          "name": "ami-012a353bb3afb92ee",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "aws-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.08.21 14:20:39 Fri",
          "creationDate": "2026-07-31T07:25:12.000Z",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731",
          "osDiskType": "ebs",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "BlockDeviceMappings",
              "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-047f4bba8e035b38e,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CreationDate",
              "value": "2026-07-31T07:25:12.000Z"
            },
            {
              "key": "DeprecationTime",
              "value": "2028-07-31T07:25:12.000Z"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04, amd64 jammy image"
            },
            {
              "key": "EnaSupport",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "ImageId",
              "value": "ami-012a353bb3afb92ee"
            },
            {
              "key": "ImageLocation",
              "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731"
            },
            {
              "key": "ImageOwnerAlias",
              "value": "amazon"
            },
            {
              "key": "ImageType",
              "value": "machine"
            },
            {
              "key": "Name",
              "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731"
            },
            {
              "key": "OwnerId",
              "value": "099720109477"
            },
            {
              "key": "PlatformDetails",
              "value": "Linux/UNIX"
            },
            {
              "key": "Public",
              "value": "true"
            },
            {
              "key": "RootDeviceName",
              "value": "/dev/sda1"
            },
            {
              "key": "RootDeviceType",
              "value": "ebs"
            },
            {
              "key": "SriovNetSupport",
              "value": "simple"
            },
            {
              "key": "State",
              "value": "available"
            },
            {
              "key": "UsageOperation",
              "value": "RunInstances"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04, amd64 jammy image",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "aws-ap-northeast-2",
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
          "connectionName": "aws-ap-northeast-2",
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
  "status": "highly-matched",
  "description": "Candidate #1 | highly-matched | 1 NLB(s) | Overall Match Rate: Min=100.0% Max=100.0% Avg=100.0% | VMs: 2 total, 2 matched, 0 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
  "targetCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
  },
  "targetInfra": {
    "name": "my01-infra101",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "NLB-aware recommended infrastructure for cloud migration",
    "nodeGroups": [
      {
        "name": "my01-ng-influxdb-back",
        "nodeGroupSize": 2,
        "label": {
          "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "nlbBackend": "influxdb_back"
        },
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
        "connectionName": "aws-ap-northeast-2",
        "specId": "aws+ap-northeast-2+t3.xlarge",
        "imageId": "ami-012a353bb3afb92ee",
        "vNetId": "my01-mig-vnet-01",
        "subnetId": "my01-mig-subnet-01",
        "securityGroupIds": [
          "my01-mig-sg-01"
        ],
        "sshKeyId": "my01-mig-sshkey-01",
        "rootDiskSize": 30,
        "dataDiskIds": null
      },
      {
        "name": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
        "connectionName": "aws-ap-northeast-2",
        "specId": "aws+ap-northeast-2+t3.small",
        "imageId": "ami-012a353bb3afb92ee",
        "vNetId": "my01-mig-vnet-01",
        "subnetId": "my01-mig-subnet-01",
        "securityGroupIds": [
          "my01-mig-sg-02"
        ],
        "sshKeyId": "my01-mig-sshkey-01",
        "rootDiskSize": 10,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "my01-mig-vnet-01",
    "connectionName": "aws-ap-northeast-2",
    "cidrBlock": "10.0.0.0/21",
    "subnetInfoList": [
      {
        "name": "my01-mig-subnet-01",
        "ipv4_CIDR": "10.0.1.0/24",
        "description": "a recommended subnet for migration"
      }
    ],
    "description": "a recommended vNet for migration"
  },
  "targetSshKey": {
    "name": "my01-mig-sshkey-01",
    "connectionName": "aws-ap-northeast-2",
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
      "id": "aws+ap-northeast-2+t3.xlarge",
      "uid": "tbt1via2qor5okve6pun",
      "cspSpecName": "t3.xlarge",
      "name": "aws+ap-northeast-2+t3.xlarge",
      "namespace": "system",
      "connectionName": "aws-ap-northeast-2",
      "providerName": "aws",
      "regionName": "ap-northeast-2",
      "regionLatitude": 37.36,
      "regionLongitude": 126.78,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 16,
      "diskSizeGB": -1,
      "costPerHour": 0.208,
      "orderInFilteredResult": 17,
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
          "key": "AutoRecoverySupported",
          "value": "true"
        },
        {
          "key": "BareMetal",
          "value": "false"
        },
        {
          "key": "BurstablePerformanceSupported",
          "value": "true"
        },
        {
          "key": "CurrentGeneration",
          "value": "true"
        },
        {
          "key": "DedicatedHostsSupported",
          "value": "true"
        },
        {
          "key": "EbsInfo",
          "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
        },
        {
          "key": "FreeTierEligible",
          "value": "false"
        },
        {
          "key": "HibernationSupported",
          "value": "true"
        },
        {
          "key": "Hypervisor",
          "value": "nitro"
        },
        {
          "key": "InstanceStorageSupported",
          "value": "false"
        },
        {
          "key": "InstanceType",
          "value": "t3.xlarge"
        },
        {
          "key": "MemoryInfo",
          "value": "{SizeInMiB:16384}"
        },
        {
          "key": "NetworkInfo",
          "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:15,Ipv6AddressesPerInterface:15,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:4,NetworkCards:[{MaximumNetworkInterfaces:4,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
        },
        {
          "key": "PlacementGroupInfo",
          "value": "{SupportedStrategies:[partition,spread]}"
        },
        {
          "key": "ProcessorInfo",
          "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.5}"
        },
        {
          "key": "SupportedBootModes",
          "value": "legacy-bios; uefi"
        },
        {
          "key": "SupportedRootDeviceTypes",
          "value": "ebs"
        },
        {
          "key": "SupportedUsageClasses",
          "value": "on-demand; spot"
        },
        {
          "key": "SupportedVirtualizationTypes",
          "value": "hvm"
        },
        {
          "key": "VCpuInfo",
          "value": "{DefaultCores:2,DefaultThreadsPerCore:2,DefaultVCpus:4,ValidCores:[2],ValidThreadsPerCore:[1,2]}"
        }
      ],
      "availabilityStatus": "Unverified"
    },
    {
      "id": "aws+ap-northeast-2+t3.small",
      "uid": "tbpf96u7vfjaf20v9o5q",
      "cspSpecName": "t3.small",
      "name": "aws+ap-northeast-2+t3.small",
      "namespace": "system",
      "connectionName": "aws-ap-northeast-2",
      "providerName": "aws",
      "regionName": "ap-northeast-2",
      "regionLatitude": 37.36,
      "regionLongitude": 126.78,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 2,
      "diskSizeGB": -1,
      "costPerHour": 0.026,
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
          "key": "AutoRecoverySupported",
          "value": "true"
        },
        {
          "key": "BareMetal",
          "value": "false"
        },
        {
          "key": "BurstablePerformanceSupported",
          "value": "true"
        },
        {
          "key": "CurrentGeneration",
          "value": "true"
        },
        {
          "key": "DedicatedHostsSupported",
          "value": "true"
        },
        {
          "key": "EbsInfo",
          "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:174,BaselineIops:1000,BaselineThroughputInMBps:21.75,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
        },
        {
          "key": "FreeTierEligible",
          "value": "true"
        },
        {
          "key": "HibernationSupported",
          "value": "true"
        },
        {
          "key": "Hypervisor",
          "value": "nitro"
        },
        {
          "key": "InstanceStorageSupported",
          "value": "false"
        },
        {
          "key": "InstanceType",
          "value": "t3.small"
        },
        {
          "key": "MemoryInfo",
          "value": "{SizeInMiB:2048}"
        },
        {
          "key": "NetworkInfo",
          "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:4,Ipv6AddressesPerInterface:4,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
        },
        {
          "key": "PlacementGroupInfo",
          "value": "{SupportedStrategies:[partition,spread]}"
        },
        {
          "key": "ProcessorInfo",
          "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.5}"
        },
        {
          "key": "SupportedBootModes",
          "value": "legacy-bios; uefi"
        },
        {
          "key": "SupportedRootDeviceTypes",
          "value": "ebs"
        },
        {
          "key": "SupportedUsageClasses",
          "value": "on-demand; spot"
        },
        {
          "key": "SupportedVirtualizationTypes",
          "value": "hvm"
        },
        {
          "key": "VCpuInfo",
          "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
        }
      ],
      "availabilityStatus": "Available",
      "availableZones": [
        "ap-northeast-2a",
        "ap-northeast-2b",
        "ap-northeast-2c",
        "ap-northeast-2d"
      ],
      "availabilityCheckedAt": "2026-10-30T12:45:32.100139859Z",
      "availabilityRefreshAfter": "2026-10-30T13:45:32.100139859Z"
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "aws",
      "cspImageName": "ami-012a353bb3afb92ee",
      "regionList": [
        "ap-northeast-2"
      ],
      "id": "ami-012a353bb3afb92ee",
      "uid": "tbbe6d8vhrqe5l5nne2u",
      "name": "ami-012a353bb3afb92ee",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "aws-ap-northeast-2",
      "infraType": "",
      "fetchedTime": "2026.08.21 14:20:39 Fri",
      "creationDate": "2026-07-31T07:25:12.000Z",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731",
      "osDiskType": "ebs",
      "osDiskSizeGB": -1,
      "imageStatus": "Available",
      "details": [
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-047f4bba8e035b38e,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CreationDate",
          "value": "2026-07-31T07:25:12.000Z"
        },
        {
          "key": "DeprecationTime",
          "value": "2028-07-31T07:25:12.000Z"
        },
        {
          "key": "Description",
          "value": "Canonical, Ubuntu, 22.04, amd64 jammy image"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-012a353bb3afb92ee"
        },
        {
          "key": "ImageLocation",
          "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731"
        },
        {
          "key": "ImageOwnerAlias",
          "value": "amazon"
        },
        {
          "key": "ImageType",
          "value": "machine"
        },
        {
          "key": "Name",
          "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731"
        },
        {
          "key": "OwnerId",
          "value": "099720109477"
        },
        {
          "key": "PlatformDetails",
          "value": "Linux/UNIX"
        },
        {
          "key": "Public",
          "value": "true"
        },
        {
          "key": "RootDeviceName",
          "value": "/dev/sda1"
        },
        {
          "key": "RootDeviceType",
          "value": "ebs"
        },
        {
          "key": "SriovNetSupport",
          "value": "simple"
        },
        {
          "key": "State",
          "value": "available"
        },
        {
          "key": "UsageOperation",
          "value": "RunInstances"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        }
      ],
      "systemLabel": "",
      "description": "Canonical, Ubuntu, 22.04, amd64 jammy image",
      "commandHistory": null
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "my01-mig-sg-01",
      "connectionName": "aws-ap-northeast-2",
      "vNetId": "my01-mig-vnet-01",
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
      "name": "my01-mig-sg-02",
      "connectionName": "aws-ap-northeast-2",
      "vNetId": "my01-mig-vnet-01",
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
- **MCI ID**: `my01-infra101`
- **MCI Name**: `my01-infra101`
- **Status**: `Running:3 (R:3/3)`

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:09",
      "label": {
        "Name": "tbk92qu15e8j9kjos3op",
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "rootDeviceName": "/dev/sda1",
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
      "cspSshKeyId": "tbq5thtqubnf05ts4une",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFziZflp6YkjKAHOxUGTXAkQztKDTLb/PobF3s0x/0wo/+1fyHgQFrmdUWZ2X811yrEDLAINKjgTp8UIJaRXQIo=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:c+dhn5mvwt4uabcY25+o7Ne66/SiyKv1zSJDnYfBX3E",
        "firstUsedAt": "2026-10-30T12:46:35Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-10-30T12:46:11Z",
          "completedTime": "2026-10-30T12:46:21Z",
          "elapsedTime": 10,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.201:22 (as \"cb-user\") via bastion Node \"my01-ng-influxdb-back-1\" at 13.125.79.237:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 13.125.79.237:22 as user \"cb-user\" (bastionNodeId=my01-ng-influxdb-back-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:34Z",
          "completedTime": "2026-10-30T12:46:36Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my01-infra101-tbnjm3s9k56b5vmsobgu",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:37Z",
          "completedTime": "2026-10-30T12:46:38Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-201 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:07",
      "label": {
        "Name": "tb3e0r0s18nekarl2ulm",
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "rootDeviceName": "/dev/sda1",
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
      "cspSshKeyId": "tbq5thtqubnf05ts4une",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGvh0rfhvr8i5F+cqCDXM3PqF+ZoQwq8WuCE8H8/eHp9xWMb5/fvAoDZKaZPkIwCu73WaQ9x7WXEvcSo2KCLkXQ=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:dGPrUYFDrEOpRLMgpP4gwFVUBGPRMOtzoFmNxM2hdxY",
        "firstUsedAt": "2026-10-30T12:46:11Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:11Z",
          "completedTime": "2026-10-30T12:46:24Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-10-30T12:46:37Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my01-infra101-tbnjm3s9k56b5vmsobgu",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:37Z",
          "completedTime": "2026-10-30T12:46:39Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-221 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:11",
      "label": {
        "Name": "tbeu73dsbj12jnttj7iq",
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "rootDeviceName": "/dev/sda1",
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
      "cspSshKeyId": "tbq5thtqubnf05ts4une",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBIPazjRPMB1whuS1eIIHrL0LqModIykiykHXJQE3J5SK+jVD+LjI5g1hqUSOrjjBpyJ+fG8M1F2UfXA4XG8dyM4=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:mR2YU5Kf7cVze+Ba0W74y1hM+tHWAIF1NyO4WgY3kQ8",
        "firstUsedAt": "2026-10-30T12:46:35Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-10-30T12:46:11Z",
          "completedTime": "2026-10-30T12:46:21Z",
          "elapsedTime": 10,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my01-ng-influxdb-back-2\" at 10.0.1.253:22 (as \"cb-user\") via bastion Node \"my01-ng-influxdb-back-1\" at 13.125.79.237:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 13.125.79.237:22 as user \"cb-user\" (bastionNodeId=my01-ng-influxdb-back-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:34Z",
          "completedTime": "2026-10-30T12:46:36Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my01-infra101-tbnjm3s9k56b5vmsobgu",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:37Z",
          "completedTime": "2026-10-30T12:46:38Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-253 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my01-mig-vnet-01",
      "name": "my01-mig-vnet-01",
      "infraId": "my01-infra101",
      "vNetId": "my01-mig-vnet-01",
      "connectionNames": [
        "aws-ap-northeast-2"
      ],
      "providerNames": [
        "aws"
      ],
      "regionNames": [
        "ap-northeast-2"
      ],
      "nodeGroupIds": [
        "my01-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my01-ng-influxdb-back"
      ],
      "nodeIds": [
        "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my01-ng-influxdb-back-1",
        "my01-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "54.180.141.244",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-201 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-influxdb-back-2",
            "nodeIp": "3.35.54.160",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-253 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-influxdb-back-1",
            "nodeIp": "13.125.79.237",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-221 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my01-infra101-tbnjm3s9k56b5vmsobgu"
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
- **Count**: 2

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
      "status": "Partial-Creating:2 (R:1/3)",
      "statusCount": {
        "countTotal": 3,
        "countCreating": 2,
        "countRunning": 1,
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
          "name": "my02-ng-influxdb-back-1",
          "nodeGroupId": "my02-ng-influxdb-back",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
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
            "region": "koreasouth"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my02-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my02-ng-influxdb-back-2",
          "uid": "tbnhrp3lhf257rnbhamg",
          "name": "my02-ng-influxdb-back-2",
          "nodeGroupId": "my02-ng-influxdb-back",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
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
            "region": "koreasouth"
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
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
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my02-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
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
- **IDs**: [my01-infra101 my02-infra101]

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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:09",
      "label": {
        "Name": "tbk92qu15e8j9kjos3op",
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "rootDeviceName": "/dev/sda1",
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
      "cspSshKeyId": "tbq5thtqubnf05ts4une",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFziZflp6YkjKAHOxUGTXAkQztKDTLb/PobF3s0x/0wo/+1fyHgQFrmdUWZ2X811yrEDLAINKjgTp8UIJaRXQIo=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:c+dhn5mvwt4uabcY25+o7Ne66/SiyKv1zSJDnYfBX3E",
        "firstUsedAt": "2026-10-30T12:46:35Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-10-30T12:46:11Z",
          "completedTime": "2026-10-30T12:46:21Z",
          "elapsedTime": 10,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.201:22 (as \"cb-user\") via bastion Node \"my01-ng-influxdb-back-1\" at 13.125.79.237:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 13.125.79.237:22 as user \"cb-user\" (bastionNodeId=my01-ng-influxdb-back-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:34Z",
          "completedTime": "2026-10-30T12:46:36Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my01-infra101-tbnjm3s9k56b5vmsobgu",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:37Z",
          "completedTime": "2026-10-30T12:46:38Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-201 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:07",
      "label": {
        "Name": "tb3e0r0s18nekarl2ulm",
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "rootDeviceName": "/dev/sda1",
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
      "cspSshKeyId": "tbq5thtqubnf05ts4une",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGvh0rfhvr8i5F+cqCDXM3PqF+ZoQwq8WuCE8H8/eHp9xWMb5/fvAoDZKaZPkIwCu73WaQ9x7WXEvcSo2KCLkXQ=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:dGPrUYFDrEOpRLMgpP4gwFVUBGPRMOtzoFmNxM2hdxY",
        "firstUsedAt": "2026-10-30T12:46:11Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:11Z",
          "completedTime": "2026-10-30T12:46:24Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-10-30T12:46:37Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my01-infra101-tbnjm3s9k56b5vmsobgu",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:37Z",
          "completedTime": "2026-10-30T12:46:39Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-221 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-10-30 12:46:11",
      "label": {
        "Name": "tbeu73dsbj12jnttj7iq",
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "rootDeviceName": "/dev/sda1",
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
      "cspSshKeyId": "tbq5thtqubnf05ts4une",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBIPazjRPMB1whuS1eIIHrL0LqModIykiykHXJQE3J5SK+jVD+LjI5g1hqUSOrjjBpyJ+fG8M1F2UfXA4XG8dyM4=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:mR2YU5Kf7cVze+Ba0W74y1hM+tHWAIF1NyO4WgY3kQ8",
        "firstUsedAt": "2026-10-30T12:46:35Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-10-30T12:46:11Z",
          "completedTime": "2026-10-30T12:46:21Z",
          "elapsedTime": 10,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my01-ng-influxdb-back-2\" at 10.0.1.253:22 (as \"cb-user\") via bastion Node \"my01-ng-influxdb-back-1\" at 13.125.79.237:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 13.125.79.237:22 as user \"cb-user\" (bastionNodeId=my01-ng-influxdb-back-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:34Z",
          "completedTime": "2026-10-30T12:46:36Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "pc-my01-infra101-tbnjm3s9k56b5vmsobgu",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-10-30T12:46:37Z",
          "completedTime": "2026-10-30T12:46:38Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-253 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my01-mig-vnet-01",
      "name": "my01-mig-vnet-01",
      "infraId": "my01-infra101",
      "vNetId": "my01-mig-vnet-01",
      "connectionNames": [
        "aws-ap-northeast-2"
      ],
      "providerNames": [
        "aws"
      ],
      "regionNames": [
        "ap-northeast-2"
      ],
      "nodeGroupIds": [
        "my01-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my01-ng-influxdb-back"
      ],
      "nodeIds": [
        "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my01-ng-influxdb-back-1",
        "my01-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "54.180.141.244",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-201 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-influxdb-back-2",
            "nodeIp": "3.35.54.160",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-253 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-influxdb-back-1",
            "nodeIp": "13.125.79.237",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-221 6.8.0-1061-aws #64~22.04.1-Ubuntu SMP Thu Jul 16 14:38:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my01-infra101-tbnjm3s9k56b5vmsobgu"
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
      "id": "my01-ng-influxdb-back",
      "cspResourceName": "tbsd5l543sjssm4r416n",
      "cspResourceId": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e",
      "name": "my01-ng-influxdb-back",
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
      "type": "PUBLIC",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "port": "9999",
        "dnsName": "tbsd5l543sjssm4r416n-ef4968adbbe89e0e.elb.ap-northeast-2.amazonaws.com",
        "keyValueList": [
          {
            "key": "DefaultActions",
            "value": "{AuthenticateCognitoConfig:null,AuthenticateOidcConfig:null,FixedResponseConfig:null,ForwardConfig:{TargetGroupStickinessConfig:{DurationSeconds:null,Enabled:false},TargetGroups:[{TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675,Weight:null}]},Order:null,RedirectConfig:null,TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675,Type:forward}"
          },
          {
            "key": "ListenerArn",
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:listener/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e/784175fe6ff692d1"
          },
          {
            "key": "LoadBalancerArn",
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
          },
          {
            "key": "Port",
            "value": "9999"
          },
          {
            "key": "Protocol",
            "value": "TCP"
          }
        ]
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my01-ng-influxdb-back",
        "nodes": [
          "my01-ng-influxdb-back-1",
          "my01-ng-influxdb-back-2"
        ],
        "keyValueList": [
          {
            "key": "HealthCheckEnabled",
            "value": "true"
          },
          {
            "key": "HealthCheckIntervalSeconds",
            "value": "10"
          },
          {
            "key": "HealthCheckPort",
            "value": "8086"
          },
          {
            "key": "HealthCheckProtocol",
            "value": "TCP"
          },
          {
            "key": "HealthCheckTimeoutSeconds",
            "value": "10"
          },
          {
            "key": "HealthyThresholdCount",
            "value": "3"
          },
          {
            "key": "LoadBalancerArns",
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
          },
          {
            "key": "Port",
            "value": "8086"
          },
          {
            "key": "Protocol",
            "value": "TCP"
          },
          {
            "key": "TargetGroupArn",
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675"
          },
          {
            "key": "TargetGroupName",
            "value": "tbsd5l543sjssm4r416n"
          },
          {
            "key": "TargetType",
            "value": "instance"
          },
          {
            "key": "UnhealthyThresholdCount",
            "value": "3"
          },
          {
            "key": "VpcId",
            "value": "vpc-02e0dfd3cb3ff1fcb"
          }
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
            "key": "HealthCheckEnabled",
            "value": "true"
          },
          {
            "key": "HealthCheckIntervalSeconds",
            "value": "10"
          },
          {
            "key": "HealthCheckPort",
            "value": "8086"
          },
          {
            "key": "HealthCheckProtocol",
            "value": "TCP"
          },
          {
            "key": "HealthCheckTimeoutSeconds",
            "value": "10"
          },
          {
            "key": "HealthyThresholdCount",
            "value": "3"
          },
          {
            "key": "LoadBalancerArns",
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
          },
          {
            "key": "Port",
            "value": "8086"
          },
          {
            "key": "Protocol",
            "value": "TCP"
          },
          {
            "key": "TargetGroupArn",
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675"
          },
          {
            "key": "TargetGroupName",
            "value": "tbsd5l543sjssm4r416n"
          },
          {
            "key": "TargetType",
            "value": "instance"
          },
          {
            "key": "UnhealthyThresholdCount",
            "value": "3"
          },
          {
            "key": "VpcId",
            "value": "vpc-02e0dfd3cb3ff1fcb"
          }
        ]
      },
      "createdTime": "2026-10-30T12:46:34.188Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "AvailabilityZones",
          "value": "{LoadBalancerAddresses:null,OutpostId:null,SubnetId:subnet-09bb75a1817ae8dbb,ZoneName:ap-northeast-2a}"
        },
        {
          "key": "CanonicalHostedZoneId",
          "value": "ZIBE1TIR4HY56"
        },
        {
          "key": "CreatedTime",
          "value": "2026-10-30T12:46:34.188Z"
        },
        {
          "key": "DNSName",
          "value": "tbsd5l543sjssm4r416n-ef4968adbbe89e0e.elb.ap-northeast-2.amazonaws.com"
        },
        {
          "key": "IpAddressType",
          "value": "ipv4"
        },
        {
          "key": "LoadBalancerArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
        },
        {
          "key": "LoadBalancerName",
          "value": "tbsd5l543sjssm4r416n"
        },
        {
          "key": "Scheme",
          "value": "internet-facing"
        },
        {
          "key": "State",
          "value": "{Code:provisioning,Reason:null}"
        },
        {
          "key": "Type",
          "value": "network"
        },
        {
          "key": "VpcId",
          "value": "vpc-02e0dfd3cb3ff1fcb"
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
    "id": "my01-ng-influxdb-back",
    "cspResourceName": "tbsd5l543sjssm4r416n",
    "cspResourceId": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e",
    "name": "my01-ng-influxdb-back",
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
    "type": "PUBLIC",
    "scope": "REGION",
    "listener": {
      "protocol": "TCP",
      "port": "9999",
      "dnsName": "tbsd5l543sjssm4r416n-ef4968adbbe89e0e.elb.ap-northeast-2.amazonaws.com",
      "keyValueList": [
        {
          "key": "DefaultActions",
          "value": "{AuthenticateCognitoConfig:null,AuthenticateOidcConfig:null,FixedResponseConfig:null,ForwardConfig:{TargetGroupStickinessConfig:{DurationSeconds:null,Enabled:false},TargetGroups:[{TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675,Weight:null}]},Order:null,RedirectConfig:null,TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675,Type:forward}"
        },
        {
          "key": "ListenerArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:listener/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e/784175fe6ff692d1"
        },
        {
          "key": "LoadBalancerArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
        },
        {
          "key": "Port",
          "value": "9999"
        },
        {
          "key": "Protocol",
          "value": "TCP"
        }
      ]
    },
    "targetGroup": {
      "protocol": "TCP",
      "port": "8086",
      "nodeGroupId": "my01-ng-influxdb-back",
      "nodes": [
        "my01-ng-influxdb-back-1",
        "my01-ng-influxdb-back-2"
      ],
      "keyValueList": [
        {
          "key": "HealthCheckEnabled",
          "value": "true"
        },
        {
          "key": "HealthCheckIntervalSeconds",
          "value": "10"
        },
        {
          "key": "HealthCheckPort",
          "value": "8086"
        },
        {
          "key": "HealthCheckProtocol",
          "value": "TCP"
        },
        {
          "key": "HealthCheckTimeoutSeconds",
          "value": "10"
        },
        {
          "key": "HealthyThresholdCount",
          "value": "3"
        },
        {
          "key": "LoadBalancerArns",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
        },
        {
          "key": "Port",
          "value": "8086"
        },
        {
          "key": "Protocol",
          "value": "TCP"
        },
        {
          "key": "TargetGroupArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675"
        },
        {
          "key": "TargetGroupName",
          "value": "tbsd5l543sjssm4r416n"
        },
        {
          "key": "TargetType",
          "value": "instance"
        },
        {
          "key": "UnhealthyThresholdCount",
          "value": "3"
        },
        {
          "key": "VpcId",
          "value": "vpc-02e0dfd3cb3ff1fcb"
        }
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
          "key": "HealthCheckEnabled",
          "value": "true"
        },
        {
          "key": "HealthCheckIntervalSeconds",
          "value": "10"
        },
        {
          "key": "HealthCheckPort",
          "value": "8086"
        },
        {
          "key": "HealthCheckProtocol",
          "value": "TCP"
        },
        {
          "key": "HealthCheckTimeoutSeconds",
          "value": "10"
        },
        {
          "key": "HealthyThresholdCount",
          "value": "3"
        },
        {
          "key": "LoadBalancerArns",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
        },
        {
          "key": "Port",
          "value": "8086"
        },
        {
          "key": "Protocol",
          "value": "TCP"
        },
        {
          "key": "TargetGroupArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675"
        },
        {
          "key": "TargetGroupName",
          "value": "tbsd5l543sjssm4r416n"
        },
        {
          "key": "TargetType",
          "value": "instance"
        },
        {
          "key": "UnhealthyThresholdCount",
          "value": "3"
        },
        {
          "key": "VpcId",
          "value": "vpc-02e0dfd3cb3ff1fcb"
        }
      ]
    },
    "createdTime": "2026-10-30T12:46:34.188Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "AvailabilityZones",
        "value": "{LoadBalancerAddresses:null,OutpostId:null,SubnetId:subnet-09bb75a1817ae8dbb,ZoneName:ap-northeast-2a}"
      },
      {
        "key": "CanonicalHostedZoneId",
        "value": "ZIBE1TIR4HY56"
      },
      {
        "key": "CreatedTime",
        "value": "2026-10-30T12:46:34.188Z"
      },
      {
        "key": "DNSName",
        "value": "tbsd5l543sjssm4r416n-ef4968adbbe89e0e.elb.ap-northeast-2.amazonaws.com"
      },
      {
        "key": "IpAddressType",
        "value": "ipv4"
      },
      {
        "key": "LoadBalancerArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
      },
      {
        "key": "LoadBalancerName",
        "value": "tbsd5l543sjssm4r416n"
      },
      {
        "key": "Scheme",
        "value": "internet-facing"
      },
      {
        "key": "State",
        "value": "{Code:provisioning,Reason:null}"
      },
      {
        "key": "Type",
        "value": "network"
      },
      {
        "key": "VpcId",
        "value": "vpc-02e0dfd3cb3ff1fcb"
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
  "id": "my01-ng-influxdb-back",
  "cspResourceName": "tbsd5l543sjssm4r416n",
  "cspResourceId": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e",
  "name": "my01-ng-influxdb-back",
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
  "type": "PUBLIC",
  "scope": "REGION",
  "listener": {
    "protocol": "TCP",
    "port": "9999",
    "dnsName": "tbsd5l543sjssm4r416n-ef4968adbbe89e0e.elb.ap-northeast-2.amazonaws.com",
    "keyValueList": [
      {
        "key": "DefaultActions",
        "value": "{AuthenticateCognitoConfig:null,AuthenticateOidcConfig:null,FixedResponseConfig:null,ForwardConfig:{TargetGroupStickinessConfig:{DurationSeconds:null,Enabled:false},TargetGroups:[{TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675,Weight:null}]},Order:null,RedirectConfig:null,TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675,Type:forward}"
      },
      {
        "key": "ListenerArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:listener/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e/784175fe6ff692d1"
      },
      {
        "key": "LoadBalancerArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
      },
      {
        "key": "Port",
        "value": "9999"
      },
      {
        "key": "Protocol",
        "value": "TCP"
      }
    ]
  },
  "targetGroup": {
    "protocol": "TCP",
    "port": "8086",
    "nodeGroupId": "my01-ng-influxdb-back",
    "nodes": [
      "my01-ng-influxdb-back-1",
      "my01-ng-influxdb-back-2"
    ],
    "keyValueList": [
      {
        "key": "HealthCheckEnabled",
        "value": "true"
      },
      {
        "key": "HealthCheckIntervalSeconds",
        "value": "10"
      },
      {
        "key": "HealthCheckPort",
        "value": "8086"
      },
      {
        "key": "HealthCheckProtocol",
        "value": "TCP"
      },
      {
        "key": "HealthCheckTimeoutSeconds",
        "value": "10"
      },
      {
        "key": "HealthyThresholdCount",
        "value": "3"
      },
      {
        "key": "LoadBalancerArns",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
      },
      {
        "key": "Port",
        "value": "8086"
      },
      {
        "key": "Protocol",
        "value": "TCP"
      },
      {
        "key": "TargetGroupArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675"
      },
      {
        "key": "TargetGroupName",
        "value": "tbsd5l543sjssm4r416n"
      },
      {
        "key": "TargetType",
        "value": "instance"
      },
      {
        "key": "UnhealthyThresholdCount",
        "value": "3"
      },
      {
        "key": "VpcId",
        "value": "vpc-02e0dfd3cb3ff1fcb"
      }
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
        "key": "HealthCheckEnabled",
        "value": "true"
      },
      {
        "key": "HealthCheckIntervalSeconds",
        "value": "10"
      },
      {
        "key": "HealthCheckPort",
        "value": "8086"
      },
      {
        "key": "HealthCheckProtocol",
        "value": "TCP"
      },
      {
        "key": "HealthCheckTimeoutSeconds",
        "value": "10"
      },
      {
        "key": "HealthyThresholdCount",
        "value": "3"
      },
      {
        "key": "LoadBalancerArns",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
      },
      {
        "key": "Port",
        "value": "8086"
      },
      {
        "key": "Protocol",
        "value": "TCP"
      },
      {
        "key": "TargetGroupArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbsd5l543sjssm4r416n/d62c019f14fe8675"
      },
      {
        "key": "TargetGroupName",
        "value": "tbsd5l543sjssm4r416n"
      },
      {
        "key": "TargetType",
        "value": "instance"
      },
      {
        "key": "UnhealthyThresholdCount",
        "value": "3"
      },
      {
        "key": "VpcId",
        "value": "vpc-02e0dfd3cb3ff1fcb"
      }
    ]
  },
  "createdTime": "2026-10-30T12:46:34.188Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "AvailabilityZones",
      "value": "{LoadBalancerAddresses:null,OutpostId:null,SubnetId:subnet-09bb75a1817ae8dbb,ZoneName:ap-northeast-2a}"
    },
    {
      "key": "CanonicalHostedZoneId",
      "value": "ZIBE1TIR4HY56"
    },
    {
      "key": "CreatedTime",
      "value": "2026-10-30T12:46:34.188Z"
    },
    {
      "key": "DNSName",
      "value": "tbsd5l543sjssm4r416n-ef4968adbbe89e0e.elb.ap-northeast-2.amazonaws.com"
    },
    {
      "key": "IpAddressType",
      "value": "ipv4"
    },
    {
      "key": "LoadBalancerArn",
      "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbsd5l543sjssm4r416n/ef4968adbbe89e0e"
    },
    {
      "key": "LoadBalancerName",
      "value": "tbsd5l543sjssm4r416n"
    },
    {
      "key": "Scheme",
      "value": "internet-facing"
    },
    {
      "key": "State",
      "value": "{Code:provisioning,Reason:null}"
    },
    {
      "key": "Type",
      "value": "network"
    },
    {
      "key": "VpcId",
      "value": "vpc-02e0dfd3cb3ff1fcb"
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

**Generated At:** 2026-10-30 12:51:03

**Namespace:** mig01

**Infra Name:** my01-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my01-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | AWS |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| t3.small | 2 | 2.0 | - | x86_64 |  | $0.0260 | 1 |
| t3.xlarge | 4 | 16.0 | - | x86_64 |  | $0.2080 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ami-012a353bb3afb92ee | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | i-0e9a8dff79587097e | Running | 2 vCPU, 2.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 54.180.141.244<br>**Private IP:** 10.0.1.201<br>**SGs:** my01-mig-sg-02<br>**SSH:** my01-mig-sshkey-01 |
| my01-ng-influxdb-back-1 | i-0dee049fb9412465b | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 13.125.79.237<br>**Private IP:** 10.0.1.221<br>**SGs:** my01-mig-sg-01<br>**SSH:** my01-mig-sshkey-01 |
| my01-ng-influxdb-back-2 | i-0cf1d4327934156dd | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260731) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 3.35.54.160<br>**Private IP:** 10.0.1.253<br>**SGs:** my01-mig-sg-01<br>**SSH:** my01-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my01-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my01-mig-vnet-01 |
| **CSP VNet ID** | vpc-02e0dfd3cb3ff1fcb |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my01-mig-subnet-01 | subnet-09bb75a1817ae8dbb | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my01-mig-sshkey-01 | tbq5thtqubnf05ts4une |  | ef:95:1c:cc:1d:6c:3f:1a:f5:c7:87:5e:bf:4d:8e:0d:fd:75:16:d5 |

### Security Groups

#### Security Group: my01-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my01-mig-sg-01 |
| **CSP Security Group ID** | sg-0704eaa2a1108babc |
| **VNet** | my01-mig-vnet-01 |
| **Rule Count** | 6 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | TCP | 9999 | 0.0.0.0/0 |
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my01-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my01-mig-sg-02 |
| **CSP Security Group ID** | sg-082244dd31713d22a |
| **VNet** | my01-mig-vnet-01 |
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
| **Per Hour** | $0.4420 |
| **Per Day** | $10.61 |
| **Per Month (30 days)** | $318.24 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AWS | ap-northeast-2 | 3 | $0.4420 | $318.24 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | t3.small | $0.0260 | $18.72 |
| my01-ng-influxdb-back-1 | t3.xlarge | $0.2080 | $149.76 |
| my01-ng-influxdb-back-2 | t3.xlarge | $0.2080 | $149.76 |




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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my01-infra101)",
  "success": true
}
```

