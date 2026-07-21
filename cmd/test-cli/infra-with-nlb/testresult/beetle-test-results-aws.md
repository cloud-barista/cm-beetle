# CM-Beetle test results for AWS (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AWS cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.5.5+ (726f2e8)
- imdl: v0.1.10+ (726f2e8)
- CB-Tumblebug: v0.12.25
- CB-Spider: v0.12.35
- CB-MapUI: v0.12.50
- Target CSP: AWS
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: July 21, 2026
- Test Time: 13:13:58 KST
- Test Execution: 2026-07-21 13:13:58 KST

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

## Test result for AWS

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 2.499s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m5.27s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 16ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 4ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 20ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 2.888s | Pass |
| 7 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 1.925s | Pass |
| 8 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 5ms | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 5ms | Pass |
| 10 | NLB Load Balancing Verification | ✅ **PASS** | 4m5.423s | Pass |
| 11 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.77s | Pass |
| 12 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.274s | Pass |
| 13 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 16.615s | Pass |
| 14 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m31.014s | Pass |

**Overall Result**: 14/14 tests passed ✅

**Total Duration**: 8m2.064529913s

*Test executed on July 21, 2026 at 13:13:58 KST (2026-07-21 13:13:58 KST) using CM-Beetle automated test CLI*

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
              "nlbBackend": "influxdb_back",
              "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3a.xlarge",
            "imageId": "ami-0afe1fd15675c3f15",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3a.small",
            "imageId": "ami-0afe1fd15675c3f15",
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
          "id": "aws+ap-northeast-2+t3a.xlarge",
          "uid": "tbssh2pg8r7mima22oe2",
          "cspSpecName": "t3a.xlarge",
          "name": "aws+ap-northeast-2+t3a.xlarge",
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
          "costPerHour": 0.1872,
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
              "value": "t3a.xlarge"
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
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
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
          ]
        },
        {
          "id": "aws+ap-northeast-2+t3a.small",
          "uid": "tbebd1j28q5o5ioi26d6",
          "cspSpecName": "t3a.small",
          "name": "aws+ap-northeast-2+t3a.small",
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
          "costPerHour": 0.0234,
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
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:175,BaselineIops:1000,BaselineThroughputInMBps:21.875,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
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
              "value": "t3a.small"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:2048}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:4,Ipv6AddressesPerInterface:4,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:2,NetworkCards:[{MaximumNetworkInterfaces:2,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
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
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "aws",
          "cspImageName": "ami-0afe1fd15675c3f15",
          "regionList": [
            "ap-northeast-2"
          ],
          "id": "ami-0afe1fd15675c3f15",
          "uid": "tbpehspo42fvpin86adv",
          "name": "ami-0afe1fd15675c3f15",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "aws-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:05:23 Mon",
          "creationDate": "2026-06-10T06:57:25.000Z",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-0060c8629d239ea6b,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CreationDate",
              "value": "2026-06-10T06:57:25.000Z"
            },
            {
              "key": "DeprecationTime",
              "value": "2028-06-10T06:57:25.000Z"
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
              "value": "ami-0afe1fd15675c3f15"
            },
            {
              "key": "ImageLocation",
              "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
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
              "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
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
      ]
    },
    {
      "status": "highly-matched",
      "description": "Candidate #2 | highly-matched | 1 NLB(s) | Overall Match Rate: Min=100.0% Max=100.0% Avg=100.0% | VMs: 2 total, 2 matched, 0 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "nlbBackend": "influxdb_back",
              "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3.xlarge",
            "imageId": "ami-0afe1fd15675c3f15",
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
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3.small",
            "imageId": "ami-0afe1fd15675c3f15",
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
          "uid": "tb09l809e55le6fu9t5n",
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
          ]
        },
        {
          "id": "aws+ap-northeast-2+t3.small",
          "uid": "tbn2t1dq7f03abs1m54h",
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
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "aws",
          "cspImageName": "ami-0afe1fd15675c3f15",
          "regionList": [
            "ap-northeast-2"
          ],
          "id": "ami-0afe1fd15675c3f15",
          "uid": "tbpehspo42fvpin86adv",
          "name": "ami-0afe1fd15675c3f15",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "aws-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:05:23 Mon",
          "creationDate": "2026-06-10T06:57:25.000Z",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-0060c8629d239ea6b,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CreationDate",
              "value": "2026-06-10T06:57:25.000Z"
            },
            {
              "key": "DeprecationTime",
              "value": "2028-06-10T06:57:25.000Z"
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
              "value": "ami-0afe1fd15675c3f15"
            },
            {
              "key": "ImageLocation",
              "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
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
              "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
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
  "uid": "tb168engtfup19aher7k",
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
    "sys.uid": "tb168engtfup19aher7k"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb7u6fnsqi1moudbbif4",
      "cspResourceName": "tb7u6fnsqi1moudbbif4",
      "cspResourceId": "i-0054a217572ac064b",
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
      "createdTime": "2026-07-21 04:14:43",
      "label": {
        "Name": "tb7u6fnsqi1moudbbif4",
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2026-07-21 04:14:43",
        "sys.cspResourceId": "i-0054a217572ac064b",
        "sys.cspResourceName": "tb7u6fnsqi1moudbbif4",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb7u6fnsqi1moudbbif4",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "43.203.180.137",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.211",
      "privateDNS": "ip-10-0-1-211.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": 10,
      "RootDeviceName": "/dev/sda1",
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
      "specId": "aws+ap-northeast-2+t3a.small",
      "cspSpecName": "t3a.small",
      "spec": {
        "cspSpecName": "t3a.small",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.0234
      },
      "imageId": "ami-0afe1fd15675c3f15",
      "cspImageName": "ami-0afe1fd15675c3f15",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-0afe1fd15675c3f15",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-0201fe7dacc5b3501",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "subnet-0b401ce3e440d6a88",
      "networkInterface": "eni-00965fb7b340143ce",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbh9rdeut22if8j5cv6k",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJUyoe5mfCp1erfPjl3PyvVKByna3hSBCZfBexUKqGSCWPF+HIOnA4SGSO6StoOOpBbIFOOX2/Ey3TYWVaWJ7QI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:F8fQbYegCXDO7uswMxBAHK0gZ6O80QfpEksjXhvJFLw",
        "firstUsedAt": "2026-07-21T04:14:50Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-21T04:14:49Z",
          "completedTime": "2026-07-21T04:15:10Z",
          "elapsedTime": 21,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-211 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-07-21T04:14:22Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-02afa74102200f386}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "E4154BA3-C8B5-404A-BDA3-BA6B2EBDBBDB"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:1,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-0afe1fd15675c3f15"
        },
        {
          "key": "InstanceId",
          "value": "i-0054a217572ac064b"
        },
        {
          "key": "InstanceType",
          "value": "t3a.small"
        },
        {
          "key": "KeyName",
          "value": "tbh9rdeut22if8j5cv6k"
        },
        {
          "key": "LaunchTime",
          "value": "2026-07-21T04:14:21Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.180.137},Attachment:{AttachTime:2026-07-21T04:14:21Z,AttachmentId:eni-attach-043bbf6a874585d2d,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b08a1f2cc70c3840,GroupName:tb2sp42g38j3600jpagj}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:94:76:f8:09:4f,NetworkInterfaceId:eni-00965fb7b340143ce,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.211,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.180.137},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.211}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0b401ce3e440d6a88,VpcId:vpc-0201fe7dacc5b3501}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-211.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.211"
        },
        {
          "key": "PublicIpAddress",
          "value": "43.203.180.137"
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
          "key": "SecurityGroups",
          "value": "{GroupId:sg-0b08a1f2cc70c3840,GroupName:tb2sp42g38j3600jpagj}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0b401ce3e440d6a88"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tb7u6fnsqi1moudbbif4}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0201fe7dacc5b3501"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tbneuo0rtp598mde00kl",
      "cspResourceName": "tbneuo0rtp598mde00kl",
      "cspResourceId": "i-0a922ef6e86d46017",
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
      "createdTime": "2026-07-21 04:14:38",
      "label": {
        "Name": "tbneuo0rtp598mde00kl",
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2026-07-21 04:14:38",
        "sys.cspResourceId": "i-0a922ef6e86d46017",
        "sys.cspResourceName": "tbneuo0rtp598mde00kl",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbneuo0rtp598mde00kl",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "3.38.153.7",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.233",
      "privateDNS": "ip-10-0-1-233.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": 30,
      "RootDeviceName": "/dev/sda1",
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
      "specId": "aws+ap-northeast-2+t3a.xlarge",
      "cspSpecName": "t3a.xlarge",
      "spec": {
        "cspSpecName": "t3a.xlarge",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1872
      },
      "imageId": "ami-0afe1fd15675c3f15",
      "cspImageName": "ami-0afe1fd15675c3f15",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-0afe1fd15675c3f15",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-0201fe7dacc5b3501",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "subnet-0b401ce3e440d6a88",
      "networkInterface": "eni-073ba4bba97dcc800",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbh9rdeut22if8j5cv6k",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBX1UeZ8Cvxy2X5GEJ5wGahwzx0xQQqAJ3ISprs4x9V6c841b1MM6mJD2Q1yL69L1G/xEh4fHhZ7vgk8vfGz5vA=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:cCWmDt/ho5e8VTO6vLsMygOlkc2X6wSbSRb21STGLcA",
        "firstUsedAt": "2026-07-21T04:14:49Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-21T04:14:49Z",
          "completedTime": "2026-07-21T04:15:10Z",
          "elapsedTime": 21,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-233 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-07-21T04:14:18Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0eb13b9837198cd82}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "49FA7C53-59BB-4609-A57B-F19D3761CE5C"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:2,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-0afe1fd15675c3f15"
        },
        {
          "key": "InstanceId",
          "value": "i-0a922ef6e86d46017"
        },
        {
          "key": "InstanceType",
          "value": "t3a.xlarge"
        },
        {
          "key": "KeyName",
          "value": "tbh9rdeut22if8j5cv6k"
        },
        {
          "key": "LaunchTime",
          "value": "2026-07-21T04:14:17Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.38.153.7},Attachment:{AttachTime:2026-07-21T04:14:17Z,AttachmentId:eni-attach-0446610409a40d148,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:9b:bd:e3:d6:59,NetworkInterfaceId:eni-073ba4bba97dcc800,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.233,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.38.153.7},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.233}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0b401ce3e440d6a88,VpcId:vpc-0201fe7dacc5b3501}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-233.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.233"
        },
        {
          "key": "PublicIpAddress",
          "value": "3.38.153.7"
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
          "key": "SecurityGroups",
          "value": "{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0b401ce3e440d6a88"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbneuo0rtp598mde00kl}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0201fe7dacc5b3501"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tbfcjl18emj99142cipm",
      "cspResourceName": "tbfcjl18emj99142cipm",
      "cspResourceId": "i-06f3cc5c469f82f40",
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
      "createdTime": "2026-07-21 04:14:41",
      "label": {
        "Name": "tbfcjl18emj99142cipm",
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2026-07-21 04:14:41",
        "sys.cspResourceId": "i-06f3cc5c469f82f40",
        "sys.cspResourceName": "tbfcjl18emj99142cipm",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbfcjl18emj99142cipm",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "3.39.251.226",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.145",
      "privateDNS": "ip-10-0-1-145.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": 30,
      "RootDeviceName": "/dev/sda1",
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
      "specId": "aws+ap-northeast-2+t3a.xlarge",
      "cspSpecName": "t3a.xlarge",
      "spec": {
        "cspSpecName": "t3a.xlarge",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1872
      },
      "imageId": "ami-0afe1fd15675c3f15",
      "cspImageName": "ami-0afe1fd15675c3f15",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-0afe1fd15675c3f15",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-0201fe7dacc5b3501",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "subnet-0b401ce3e440d6a88",
      "networkInterface": "eni-00e59e53d2438bedf",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbh9rdeut22if8j5cv6k",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEUlFwV+LidXdVFjp33gpGja29R5Zllc2VcCP7KZq1ZCI9XqDBp/pOh+VeM7fpiSGPpJZOtVZc+kdv497A6ja3I=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:jF0XS+JJjHfi4QK/kgIBKgBL79pWqPan4ncwAbaPubE",
        "firstUsedAt": "2026-07-21T04:14:50Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-21T04:14:49Z",
          "completedTime": "2026-07-21T04:14:52Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-145 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-07-21T04:14:21Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0a446b7810b00d3d5}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "C3F8023D-4CF6-4E4F-B17A-08FB56D4EDE6"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:2,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-0afe1fd15675c3f15"
        },
        {
          "key": "InstanceId",
          "value": "i-06f3cc5c469f82f40"
        },
        {
          "key": "InstanceType",
          "value": "t3a.xlarge"
        },
        {
          "key": "KeyName",
          "value": "tbh9rdeut22if8j5cv6k"
        },
        {
          "key": "LaunchTime",
          "value": "2026-07-21T04:14:20Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.251.226},Attachment:{AttachTime:2026-07-21T04:14:20Z,AttachmentId:eni-attach-06a41c8ec19fbdce5,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:43:e9:6b:c4:b1,NetworkInterfaceId:eni-00e59e53d2438bedf,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.145,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.251.226},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.145}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0b401ce3e440d6a88,VpcId:vpc-0201fe7dacc5b3501}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-145.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.145"
        },
        {
          "key": "PublicIpAddress",
          "value": "3.39.251.226"
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
          "key": "SecurityGroups",
          "value": "{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0b401ce3e440d6a88"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbfcjl18emj99142cipm}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0201fe7dacc5b3501"
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
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "3.39.251.226",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-145 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "43.203.180.137",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-211 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "3.38.153.7",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-233 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "tb168engtfup19aher7k",
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
        "sys.uid": "tb168engtfup19aher7k"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb7u6fnsqi1moudbbif4",
          "cspResourceName": "tb7u6fnsqi1moudbbif4",
          "cspResourceId": "i-0054a217572ac064b",
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
          "createdTime": "2026-07-21 04:14:43",
          "label": {
            "Name": "tb7u6fnsqi1moudbbif4",
            "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-07-21 04:14:43",
            "sys.cspResourceId": "i-0054a217572ac064b",
            "sys.cspResourceName": "tb7u6fnsqi1moudbbif4",
            "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tb7u6fnsqi1moudbbif4",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "43.203.180.137",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.211",
          "privateDNS": "ip-10-0-1-211.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 10,
          "RootDeviceName": "/dev/sda1",
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
          "specId": "aws+ap-northeast-2+t3a.small",
          "cspSpecName": "t3a.small",
          "spec": {
            "cspSpecName": "t3a.small",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.0234
          },
          "imageId": "ami-0afe1fd15675c3f15",
          "cspImageName": "ami-0afe1fd15675c3f15",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-0afe1fd15675c3f15",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "vpc-0201fe7dacc5b3501",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "subnet-0b401ce3e440d6a88",
          "networkInterface": "eni-00965fb7b340143ce",
          "securityGroupIds": [
            "my-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tbh9rdeut22if8j5cv6k",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJUyoe5mfCp1erfPjl3PyvVKByna3hSBCZfBexUKqGSCWPF+HIOnA4SGSO6StoOOpBbIFOOX2/Ey3TYWVaWJ7QI=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:F8fQbYegCXDO7uswMxBAHK0gZ6O80QfpEksjXhvJFLw",
            "firstUsedAt": "2026-07-21T04:14:50Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-21T04:14:49Z",
              "completedTime": "2026-07-21T04:15:10Z",
              "elapsedTime": 21,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-211 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "AmiLaunchIndex",
              "value": "0"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "BlockDeviceMappings",
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-07-21T04:14:22Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-02afa74102200f386}}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CapacityReservationSpecification",
              "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
            },
            {
              "key": "ClientToken",
              "value": "E4154BA3-C8B5-404A-BDA3-BA6B2EBDBBDB"
            },
            {
              "key": "CpuOptions",
              "value": "{CoreCount:1,ThreadsPerCore:2}"
            },
            {
              "key": "EbsOptimized",
              "value": "false"
            },
            {
              "key": "EnaSupport",
              "value": "true"
            },
            {
              "key": "EnclaveOptions",
              "value": "{Enabled:false}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "ImageId",
              "value": "ami-0afe1fd15675c3f15"
            },
            {
              "key": "InstanceId",
              "value": "i-0054a217572ac064b"
            },
            {
              "key": "InstanceType",
              "value": "t3a.small"
            },
            {
              "key": "KeyName",
              "value": "tbh9rdeut22if8j5cv6k"
            },
            {
              "key": "LaunchTime",
              "value": "2026-07-21T04:14:21Z"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
            },
            {
              "key": "Monitoring",
              "value": "{State:disabled}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.180.137},Attachment:{AttachTime:2026-07-21T04:14:21Z,AttachmentId:eni-attach-043bbf6a874585d2d,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b08a1f2cc70c3840,GroupName:tb2sp42g38j3600jpagj}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:94:76:f8:09:4f,NetworkInterfaceId:eni-00965fb7b340143ce,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.211,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.180.137},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.211}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0b401ce3e440d6a88,VpcId:vpc-0201fe7dacc5b3501}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-211.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.211"
            },
            {
              "key": "PublicIpAddress",
              "value": "43.203.180.137"
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
              "key": "SecurityGroups",
              "value": "{GroupId:sg-0b08a1f2cc70c3840,GroupName:tb2sp42g38j3600jpagj}"
            },
            {
              "key": "SourceDestCheck",
              "value": "true"
            },
            {
              "key": "State",
              "value": "{Code:16,Name:running}"
            },
            {
              "key": "SubnetId",
              "value": "subnet-0b401ce3e440d6a88"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tb7u6fnsqi1moudbbif4}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0201fe7dacc5b3501"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-1",
          "uid": "tbneuo0rtp598mde00kl",
          "cspResourceName": "tbneuo0rtp598mde00kl",
          "cspResourceId": "i-0a922ef6e86d46017",
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
          "createdTime": "2026-07-21 04:14:38",
          "label": {
            "Name": "tbneuo0rtp598mde00kl",
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-07-21 04:14:38",
            "sys.cspResourceId": "i-0a922ef6e86d46017",
            "sys.cspResourceName": "tbneuo0rtp598mde00kl",
            "sys.id": "my-ng-influxdb-back-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbneuo0rtp598mde00kl",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "3.38.153.7",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.233",
          "privateDNS": "ip-10-0-1-233.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 30,
          "RootDeviceName": "/dev/sda1",
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
          "specId": "aws+ap-northeast-2+t3a.xlarge",
          "cspSpecName": "t3a.xlarge",
          "spec": {
            "cspSpecName": "t3a.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1872
          },
          "imageId": "ami-0afe1fd15675c3f15",
          "cspImageName": "ami-0afe1fd15675c3f15",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-0afe1fd15675c3f15",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "vpc-0201fe7dacc5b3501",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "subnet-0b401ce3e440d6a88",
          "networkInterface": "eni-073ba4bba97dcc800",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tbh9rdeut22if8j5cv6k",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBX1UeZ8Cvxy2X5GEJ5wGahwzx0xQQqAJ3ISprs4x9V6c841b1MM6mJD2Q1yL69L1G/xEh4fHhZ7vgk8vfGz5vA=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:cCWmDt/ho5e8VTO6vLsMygOlkc2X6wSbSRb21STGLcA",
            "firstUsedAt": "2026-07-21T04:14:49Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-21T04:14:49Z",
              "completedTime": "2026-07-21T04:15:10Z",
              "elapsedTime": 21,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-233 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "AmiLaunchIndex",
              "value": "0"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "BlockDeviceMappings",
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-07-21T04:14:18Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0eb13b9837198cd82}}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CapacityReservationSpecification",
              "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
            },
            {
              "key": "ClientToken",
              "value": "49FA7C53-59BB-4609-A57B-F19D3761CE5C"
            },
            {
              "key": "CpuOptions",
              "value": "{CoreCount:2,ThreadsPerCore:2}"
            },
            {
              "key": "EbsOptimized",
              "value": "false"
            },
            {
              "key": "EnaSupport",
              "value": "true"
            },
            {
              "key": "EnclaveOptions",
              "value": "{Enabled:false}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "ImageId",
              "value": "ami-0afe1fd15675c3f15"
            },
            {
              "key": "InstanceId",
              "value": "i-0a922ef6e86d46017"
            },
            {
              "key": "InstanceType",
              "value": "t3a.xlarge"
            },
            {
              "key": "KeyName",
              "value": "tbh9rdeut22if8j5cv6k"
            },
            {
              "key": "LaunchTime",
              "value": "2026-07-21T04:14:17Z"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
            },
            {
              "key": "Monitoring",
              "value": "{State:disabled}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.38.153.7},Attachment:{AttachTime:2026-07-21T04:14:17Z,AttachmentId:eni-attach-0446610409a40d148,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:9b:bd:e3:d6:59,NetworkInterfaceId:eni-073ba4bba97dcc800,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.233,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.38.153.7},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.233}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0b401ce3e440d6a88,VpcId:vpc-0201fe7dacc5b3501}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-233.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.233"
            },
            {
              "key": "PublicIpAddress",
              "value": "3.38.153.7"
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
              "key": "SecurityGroups",
              "value": "{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}"
            },
            {
              "key": "SourceDestCheck",
              "value": "true"
            },
            {
              "key": "State",
              "value": "{Code:16,Name:running}"
            },
            {
              "key": "SubnetId",
              "value": "subnet-0b401ce3e440d6a88"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbneuo0rtp598mde00kl}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0201fe7dacc5b3501"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-2",
          "uid": "tbfcjl18emj99142cipm",
          "cspResourceName": "tbfcjl18emj99142cipm",
          "cspResourceId": "i-06f3cc5c469f82f40",
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
          "createdTime": "2026-07-21 04:14:41",
          "label": {
            "Name": "tbfcjl18emj99142cipm",
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-07-21 04:14:41",
            "sys.cspResourceId": "i-06f3cc5c469f82f40",
            "sys.cspResourceName": "tbfcjl18emj99142cipm",
            "sys.id": "my-ng-influxdb-back-2",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbfcjl18emj99142cipm",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "3.39.251.226",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.145",
          "privateDNS": "ip-10-0-1-145.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 30,
          "RootDeviceName": "/dev/sda1",
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
          "specId": "aws+ap-northeast-2+t3a.xlarge",
          "cspSpecName": "t3a.xlarge",
          "spec": {
            "cspSpecName": "t3a.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1872
          },
          "imageId": "ami-0afe1fd15675c3f15",
          "cspImageName": "ami-0afe1fd15675c3f15",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-0afe1fd15675c3f15",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "vpc-0201fe7dacc5b3501",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "subnet-0b401ce3e440d6a88",
          "networkInterface": "eni-00e59e53d2438bedf",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "tbh9rdeut22if8j5cv6k",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEUlFwV+LidXdVFjp33gpGja29R5Zllc2VcCP7KZq1ZCI9XqDBp/pOh+VeM7fpiSGPpJZOtVZc+kdv497A6ja3I=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:jF0XS+JJjHfi4QK/kgIBKgBL79pWqPan4ncwAbaPubE",
            "firstUsedAt": "2026-07-21T04:14:50Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-21T04:14:49Z",
              "completedTime": "2026-07-21T04:14:52Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-145 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "AmiLaunchIndex",
              "value": "0"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "BlockDeviceMappings",
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-07-21T04:14:21Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0a446b7810b00d3d5}}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CapacityReservationSpecification",
              "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
            },
            {
              "key": "ClientToken",
              "value": "C3F8023D-4CF6-4E4F-B17A-08FB56D4EDE6"
            },
            {
              "key": "CpuOptions",
              "value": "{CoreCount:2,ThreadsPerCore:2}"
            },
            {
              "key": "EbsOptimized",
              "value": "false"
            },
            {
              "key": "EnaSupport",
              "value": "true"
            },
            {
              "key": "EnclaveOptions",
              "value": "{Enabled:false}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "ImageId",
              "value": "ami-0afe1fd15675c3f15"
            },
            {
              "key": "InstanceId",
              "value": "i-06f3cc5c469f82f40"
            },
            {
              "key": "InstanceType",
              "value": "t3a.xlarge"
            },
            {
              "key": "KeyName",
              "value": "tbh9rdeut22if8j5cv6k"
            },
            {
              "key": "LaunchTime",
              "value": "2026-07-21T04:14:20Z"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
            },
            {
              "key": "Monitoring",
              "value": "{State:disabled}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.251.226},Attachment:{AttachTime:2026-07-21T04:14:20Z,AttachmentId:eni-attach-06a41c8ec19fbdce5,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:43:e9:6b:c4:b1,NetworkInterfaceId:eni-00e59e53d2438bedf,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.145,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.251.226},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.145}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0b401ce3e440d6a88,VpcId:vpc-0201fe7dacc5b3501}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-145.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.145"
            },
            {
              "key": "PublicIpAddress",
              "value": "3.39.251.226"
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
              "key": "SecurityGroups",
              "value": "{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}"
            },
            {
              "key": "SourceDestCheck",
              "value": "true"
            },
            {
              "key": "State",
              "value": "{Code:16,Name:running}"
            },
            {
              "key": "SubnetId",
              "value": "subnet-0b401ce3e440d6a88"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbfcjl18emj99142cipm}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0201fe7dacc5b3501"
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
            "nodeId": "my-ng-influxdb-back-2",
            "nodeIp": "3.39.251.226",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-145 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "43.203.180.137",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-211 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-1",
            "nodeIp": "3.38.153.7",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-233 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "tb168engtfup19aher7k",
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
    "sys.uid": "tb168engtfup19aher7k"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb7u6fnsqi1moudbbif4",
      "cspResourceName": "tb7u6fnsqi1moudbbif4",
      "cspResourceId": "i-0054a217572ac064b",
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
      "createdTime": "2026-07-21 04:14:43",
      "label": {
        "Name": "tb7u6fnsqi1moudbbif4",
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2026-07-21 04:14:43",
        "sys.cspResourceId": "i-0054a217572ac064b",
        "sys.cspResourceName": "tb7u6fnsqi1moudbbif4",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb7u6fnsqi1moudbbif4",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "43.203.180.137",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.211",
      "privateDNS": "ip-10-0-1-211.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": 10,
      "RootDeviceName": "/dev/sda1",
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
      "specId": "aws+ap-northeast-2+t3a.small",
      "cspSpecName": "t3a.small",
      "spec": {
        "cspSpecName": "t3a.small",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.0234
      },
      "imageId": "ami-0afe1fd15675c3f15",
      "cspImageName": "ami-0afe1fd15675c3f15",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-0afe1fd15675c3f15",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-0201fe7dacc5b3501",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "subnet-0b401ce3e440d6a88",
      "networkInterface": "eni-00965fb7b340143ce",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbh9rdeut22if8j5cv6k",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJUyoe5mfCp1erfPjl3PyvVKByna3hSBCZfBexUKqGSCWPF+HIOnA4SGSO6StoOOpBbIFOOX2/Ey3TYWVaWJ7QI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:F8fQbYegCXDO7uswMxBAHK0gZ6O80QfpEksjXhvJFLw",
        "firstUsedAt": "2026-07-21T04:14:50Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-21T04:14:49Z",
          "completedTime": "2026-07-21T04:15:10Z",
          "elapsedTime": 21,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-211 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-07-21T04:14:22Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-02afa74102200f386}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "E4154BA3-C8B5-404A-BDA3-BA6B2EBDBBDB"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:1,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-0afe1fd15675c3f15"
        },
        {
          "key": "InstanceId",
          "value": "i-0054a217572ac064b"
        },
        {
          "key": "InstanceType",
          "value": "t3a.small"
        },
        {
          "key": "KeyName",
          "value": "tbh9rdeut22if8j5cv6k"
        },
        {
          "key": "LaunchTime",
          "value": "2026-07-21T04:14:21Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.180.137},Attachment:{AttachTime:2026-07-21T04:14:21Z,AttachmentId:eni-attach-043bbf6a874585d2d,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b08a1f2cc70c3840,GroupName:tb2sp42g38j3600jpagj}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:94:76:f8:09:4f,NetworkInterfaceId:eni-00965fb7b340143ce,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.211,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.180.137},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.211}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0b401ce3e440d6a88,VpcId:vpc-0201fe7dacc5b3501}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-211.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.211"
        },
        {
          "key": "PublicIpAddress",
          "value": "43.203.180.137"
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
          "key": "SecurityGroups",
          "value": "{GroupId:sg-0b08a1f2cc70c3840,GroupName:tb2sp42g38j3600jpagj}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0b401ce3e440d6a88"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tb7u6fnsqi1moudbbif4}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0201fe7dacc5b3501"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tbneuo0rtp598mde00kl",
      "cspResourceName": "tbneuo0rtp598mde00kl",
      "cspResourceId": "i-0a922ef6e86d46017",
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
      "createdTime": "2026-07-21 04:14:38",
      "label": {
        "Name": "tbneuo0rtp598mde00kl",
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2026-07-21 04:14:38",
        "sys.cspResourceId": "i-0a922ef6e86d46017",
        "sys.cspResourceName": "tbneuo0rtp598mde00kl",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbneuo0rtp598mde00kl",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "3.38.153.7",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.233",
      "privateDNS": "ip-10-0-1-233.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": 30,
      "RootDeviceName": "/dev/sda1",
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
      "specId": "aws+ap-northeast-2+t3a.xlarge",
      "cspSpecName": "t3a.xlarge",
      "spec": {
        "cspSpecName": "t3a.xlarge",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1872
      },
      "imageId": "ami-0afe1fd15675c3f15",
      "cspImageName": "ami-0afe1fd15675c3f15",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-0afe1fd15675c3f15",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-0201fe7dacc5b3501",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "subnet-0b401ce3e440d6a88",
      "networkInterface": "eni-073ba4bba97dcc800",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbh9rdeut22if8j5cv6k",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBX1UeZ8Cvxy2X5GEJ5wGahwzx0xQQqAJ3ISprs4x9V6c841b1MM6mJD2Q1yL69L1G/xEh4fHhZ7vgk8vfGz5vA=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:cCWmDt/ho5e8VTO6vLsMygOlkc2X6wSbSRb21STGLcA",
        "firstUsedAt": "2026-07-21T04:14:49Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-21T04:14:49Z",
          "completedTime": "2026-07-21T04:15:10Z",
          "elapsedTime": 21,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-233 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-07-21T04:14:18Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0eb13b9837198cd82}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "49FA7C53-59BB-4609-A57B-F19D3761CE5C"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:2,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-0afe1fd15675c3f15"
        },
        {
          "key": "InstanceId",
          "value": "i-0a922ef6e86d46017"
        },
        {
          "key": "InstanceType",
          "value": "t3a.xlarge"
        },
        {
          "key": "KeyName",
          "value": "tbh9rdeut22if8j5cv6k"
        },
        {
          "key": "LaunchTime",
          "value": "2026-07-21T04:14:17Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.38.153.7},Attachment:{AttachTime:2026-07-21T04:14:17Z,AttachmentId:eni-attach-0446610409a40d148,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:9b:bd:e3:d6:59,NetworkInterfaceId:eni-073ba4bba97dcc800,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.233,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.38.153.7},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.233}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0b401ce3e440d6a88,VpcId:vpc-0201fe7dacc5b3501}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-233.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.233"
        },
        {
          "key": "PublicIpAddress",
          "value": "3.38.153.7"
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
          "key": "SecurityGroups",
          "value": "{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0b401ce3e440d6a88"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbneuo0rtp598mde00kl}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0201fe7dacc5b3501"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tbfcjl18emj99142cipm",
      "cspResourceName": "tbfcjl18emj99142cipm",
      "cspResourceId": "i-06f3cc5c469f82f40",
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
      "createdTime": "2026-07-21 04:14:41",
      "label": {
        "Name": "tbfcjl18emj99142cipm",
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2026-07-21 04:14:41",
        "sys.cspResourceId": "i-06f3cc5c469f82f40",
        "sys.cspResourceName": "tbfcjl18emj99142cipm",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbfcjl18emj99142cipm",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "3.39.251.226",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.145",
      "privateDNS": "ip-10-0-1-145.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": 30,
      "RootDeviceName": "/dev/sda1",
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
      "specId": "aws+ap-northeast-2+t3a.xlarge",
      "cspSpecName": "t3a.xlarge",
      "spec": {
        "cspSpecName": "t3a.xlarge",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1872
      },
      "imageId": "ami-0afe1fd15675c3f15",
      "cspImageName": "ami-0afe1fd15675c3f15",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-0afe1fd15675c3f15",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "vpc-0201fe7dacc5b3501",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "subnet-0b401ce3e440d6a88",
      "networkInterface": "eni-00e59e53d2438bedf",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "tbh9rdeut22if8j5cv6k",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEUlFwV+LidXdVFjp33gpGja29R5Zllc2VcCP7KZq1ZCI9XqDBp/pOh+VeM7fpiSGPpJZOtVZc+kdv497A6ja3I=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:jF0XS+JJjHfi4QK/kgIBKgBL79pWqPan4ncwAbaPubE",
        "firstUsedAt": "2026-07-21T04:14:50Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-21T04:14:49Z",
          "completedTime": "2026-07-21T04:14:52Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-145 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-07-21T04:14:21Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0a446b7810b00d3d5}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "C3F8023D-4CF6-4E4F-B17A-08FB56D4EDE6"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:2,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-0afe1fd15675c3f15"
        },
        {
          "key": "InstanceId",
          "value": "i-06f3cc5c469f82f40"
        },
        {
          "key": "InstanceType",
          "value": "t3a.xlarge"
        },
        {
          "key": "KeyName",
          "value": "tbh9rdeut22if8j5cv6k"
        },
        {
          "key": "LaunchTime",
          "value": "2026-07-21T04:14:20Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.251.226},Attachment:{AttachTime:2026-07-21T04:14:20Z,AttachmentId:eni-attach-06a41c8ec19fbdce5,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:43:e9:6b:c4:b1,NetworkInterfaceId:eni-00e59e53d2438bedf,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.145,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.251.226},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.145}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0b401ce3e440d6a88,VpcId:vpc-0201fe7dacc5b3501}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-145.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.145"
        },
        {
          "key": "PublicIpAddress",
          "value": "3.39.251.226"
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
          "key": "SecurityGroups",
          "value": "{GroupId:sg-0b193070b1e7f5d29,GroupName:tb9je04sf7uqsdimkv49}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0b401ce3e440d6a88"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbfcjl18emj99142cipm}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0201fe7dacc5b3501"
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
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "3.39.251.226",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-145 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "43.203.180.137",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-211 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "3.38.153.7",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-233 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "cspResourceName": "tbm6ig3fnsnpgomdcvi5",
      "cspResourceId": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c",
      "name": "my-ng-influxdb-back",
      "connectionName": "aws-ap-northeast-2",
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
        "port": "9999",
        "dnsName": "tbm6ig3fnsnpgomdcvi5-b41f7d221abf517c.elb.ap-northeast-2.amazonaws.com",
        "keyValueList": [
          {
            "key": "DefaultActions",
            "value": "{AuthenticateCognitoConfig:null,AuthenticateOidcConfig:null,FixedResponseConfig:null,ForwardConfig:{TargetGroupStickinessConfig:{DurationSeconds:null,Enabled:false},TargetGroups:[{TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483,Weight:null}]},Order:null,RedirectConfig:null,TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483,Type:forward}"
          },
          {
            "key": "ListenerArn",
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:listener/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c/47c9428e6a5a3e63"
          },
          {
            "key": "LoadBalancerArn",
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
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
        "nodeGroupId": "my-ng-influxdb-back",
        "nodes": [
          "my-ng-influxdb-back-2",
          "my-ng-influxdb-back-1"
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
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
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
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483"
          },
          {
            "key": "TargetGroupName",
            "value": "tbm6ig3fnsnpgomdcvi5"
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
            "value": "vpc-0201fe7dacc5b3501"
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
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
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
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483"
          },
          {
            "key": "TargetGroupName",
            "value": "tbm6ig3fnsnpgomdcvi5"
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
            "value": "vpc-0201fe7dacc5b3501"
          }
        ]
      },
      "createdTime": "2026-07-21T04:15:36.764Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "AvailabilityZones",
          "value": "{LoadBalancerAddresses:null,OutpostId:null,SubnetId:subnet-0b401ce3e440d6a88,ZoneName:ap-northeast-2a}"
        },
        {
          "key": "CanonicalHostedZoneId",
          "value": "ZIBE1TIR4HY56"
        },
        {
          "key": "CreatedTime",
          "value": "2026-07-21T04:15:36.764Z"
        },
        {
          "key": "DNSName",
          "value": "tbm6ig3fnsnpgomdcvi5-b41f7d221abf517c.elb.ap-northeast-2.amazonaws.com"
        },
        {
          "key": "IpAddressType",
          "value": "ipv4"
        },
        {
          "key": "LoadBalancerArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
        },
        {
          "key": "LoadBalancerName",
          "value": "tbm6ig3fnsnpgomdcvi5"
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
          "value": "vpc-0201fe7dacc5b3501"
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
    "cspResourceName": "tbm6ig3fnsnpgomdcvi5",
    "cspResourceId": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c",
    "name": "my-ng-influxdb-back",
    "connectionName": "aws-ap-northeast-2",
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
      "port": "9999",
      "dnsName": "tbm6ig3fnsnpgomdcvi5-b41f7d221abf517c.elb.ap-northeast-2.amazonaws.com",
      "keyValueList": [
        {
          "key": "DefaultActions",
          "value": "{AuthenticateCognitoConfig:null,AuthenticateOidcConfig:null,FixedResponseConfig:null,ForwardConfig:{TargetGroupStickinessConfig:{DurationSeconds:null,Enabled:false},TargetGroups:[{TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483,Weight:null}]},Order:null,RedirectConfig:null,TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483,Type:forward}"
        },
        {
          "key": "ListenerArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:listener/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c/47c9428e6a5a3e63"
        },
        {
          "key": "LoadBalancerArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
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
      "nodeGroupId": "my-ng-influxdb-back",
      "nodes": [
        "my-ng-influxdb-back-2",
        "my-ng-influxdb-back-1"
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
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
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
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483"
        },
        {
          "key": "TargetGroupName",
          "value": "tbm6ig3fnsnpgomdcvi5"
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
          "value": "vpc-0201fe7dacc5b3501"
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
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
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
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483"
        },
        {
          "key": "TargetGroupName",
          "value": "tbm6ig3fnsnpgomdcvi5"
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
          "value": "vpc-0201fe7dacc5b3501"
        }
      ]
    },
    "createdTime": "2026-07-21T04:15:36.764Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "AvailabilityZones",
        "value": "{LoadBalancerAddresses:null,OutpostId:null,SubnetId:subnet-0b401ce3e440d6a88,ZoneName:ap-northeast-2a}"
      },
      {
        "key": "CanonicalHostedZoneId",
        "value": "ZIBE1TIR4HY56"
      },
      {
        "key": "CreatedTime",
        "value": "2026-07-21T04:15:36.764Z"
      },
      {
        "key": "DNSName",
        "value": "tbm6ig3fnsnpgomdcvi5-b41f7d221abf517c.elb.ap-northeast-2.amazonaws.com"
      },
      {
        "key": "IpAddressType",
        "value": "ipv4"
      },
      {
        "key": "LoadBalancerArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
      },
      {
        "key": "LoadBalancerName",
        "value": "tbm6ig3fnsnpgomdcvi5"
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
        "value": "vpc-0201fe7dacc5b3501"
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
  "cspResourceName": "tbm6ig3fnsnpgomdcvi5",
  "cspResourceId": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c",
  "name": "my-ng-influxdb-back",
  "connectionName": "aws-ap-northeast-2",
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
    "port": "9999",
    "dnsName": "tbm6ig3fnsnpgomdcvi5-b41f7d221abf517c.elb.ap-northeast-2.amazonaws.com",
    "keyValueList": [
      {
        "key": "DefaultActions",
        "value": "{AuthenticateCognitoConfig:null,AuthenticateOidcConfig:null,FixedResponseConfig:null,ForwardConfig:{TargetGroupStickinessConfig:{DurationSeconds:null,Enabled:false},TargetGroups:[{TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483,Weight:null}]},Order:null,RedirectConfig:null,TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483,Type:forward}"
      },
      {
        "key": "ListenerArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:listener/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c/47c9428e6a5a3e63"
      },
      {
        "key": "LoadBalancerArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
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
    "nodeGroupId": "my-ng-influxdb-back",
    "nodes": [
      "my-ng-influxdb-back-2",
      "my-ng-influxdb-back-1"
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
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
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
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483"
      },
      {
        "key": "TargetGroupName",
        "value": "tbm6ig3fnsnpgomdcvi5"
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
        "value": "vpc-0201fe7dacc5b3501"
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
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
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
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbm6ig3fnsnpgomdcvi5/20e273ac5a5e5483"
      },
      {
        "key": "TargetGroupName",
        "value": "tbm6ig3fnsnpgomdcvi5"
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
        "value": "vpc-0201fe7dacc5b3501"
      }
    ]
  },
  "createdTime": "2026-07-21T04:15:36.764Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "AvailabilityZones",
      "value": "{LoadBalancerAddresses:null,OutpostId:null,SubnetId:subnet-0b401ce3e440d6a88,ZoneName:ap-northeast-2a}"
    },
    {
      "key": "CanonicalHostedZoneId",
      "value": "ZIBE1TIR4HY56"
    },
    {
      "key": "CreatedTime",
      "value": "2026-07-21T04:15:36.764Z"
    },
    {
      "key": "DNSName",
      "value": "tbm6ig3fnsnpgomdcvi5-b41f7d221abf517c.elb.ap-northeast-2.amazonaws.com"
    },
    {
      "key": "IpAddressType",
      "value": "ipv4"
    },
    {
      "key": "LoadBalancerArn",
      "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbm6ig3fnsnpgomdcvi5/b41f7d221abf517c"
    },
    {
      "key": "LoadBalancerName",
      "value": "tbm6ig3fnsnpgomdcvi5"
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
      "value": "vpc-0201fe7dacc5b3501"
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

**Generated At:** 2026-07-21 04:19:39

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
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
| t3a.small | 2 | 2.0 | - | x86_64 |  | $0.0234 | 1 |
| t3a.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1872 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ami-0afe1fd15675c3f15 | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | i-0054a217572ac064b | Running | 2 vCPU, 2.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 43.203.180.137<br>**Private IP:** 10.0.1.211<br>**SGs:** my-mig-sg-02<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-1 | i-0a922ef6e86d46017 | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 3.38.153.7<br>**Private IP:** 10.0.1.233<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-2 | i-06f3cc5c469f82f40 | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 3.39.251.226<br>**Private IP:** 10.0.1.145<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-vnet-01 |
| **CSP VNet ID** | vpc-0201fe7dacc5b3501 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-mig-subnet-01 | subnet-0b401ce3e440d6a88 | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-mig-sshkey-01 | tbh9rdeut22if8j5cv6k |  | 27:de:16:94:80:a4:a8:73:30:d0:f3:da:83:37:7d:43:54:bd:7c:1e |

### Security Groups

#### Security Group: my-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-01 |
| **CSP Security Group ID** | sg-0b193070b1e7f5d29 |
| **VNet** | my-mig-vnet-01 |
| **Rule Count** | 5 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 8086 | 10.0.0.0/16 |
| inbound | TCP | 8086 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-02 |
| **CSP Security Group ID** | sg-0b08a1f2cc70c3840 |
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
| **Per Hour** | $0.3978 |
| **Per Day** | $9.55 |
| **Per Month (30 days)** | $286.42 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AWS | ap-northeast-2 | 3 | $0.3978 | $286.42 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | t3a.small | $0.0234 | $16.85 |
| my-ng-influxdb-back-1 | t3a.xlarge | $0.1872 | $134.78 |
| my-ng-influxdb-back-2 | t3a.xlarge | $0.1872 | $134.78 |




### Test Case 12: Migration Report

#### 12.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-07-21 04:19:44*

---

## 📊 Migration Summary

**Target Cloud:** AWS

**Target Region:** ap-northeast-2

**Namespace:** mig01 | **Infra ID:** my-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $286.42 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | t3a.small, t3a.xlarge |
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
| 1 | **VM Name:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** i-0054a217572ac064b<br>**Label(sourceMachineId):** ec268ed7-821e-9d73-e79f-961262161624 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 |
| 2 | **VM Name:** my-ng-influxdb-back-1<br>**VM ID:** i-0a922ef6e86d46017<br>**Label(sourceMachineId):** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 |
| 3 | **VM Name:** my-ng-influxdb-back-2<br>**VM ID:** i-06f3cc5c469f82f40<br>**Label(sourceMachineId):** ec288dd0-c6fa-8a49-2f60-bc898311febf | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** t3a.small<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 8 GB |
| 2 | my-ng-influxdb-back-1 | **Spec ID:** t3a.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 30 GB |
| 3 | my-ng-influxdb-back-2 | **Spec ID:** t3a.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 8 GB<br>**Root Disk:** 30 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** ami-0afe1fd15675c3f15<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | my-ng-influxdb-back-1 | **Image ID:** ami-0afe1fd15675c3f15<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | my-ng-influxdb-back-2 | **Image ID:** ami-0afe1fd15675c3f15<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260610 | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 9 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my-mig-sg-01

**CSP ID:** sg-0b193070b1e7f5d29 | **VNet:** my-mig-vnet-01 | **Rules:** 5

**Assigned VMs:**

- **VM:** my-ng-influxdb-back-1
  - **Source Server:** **Hostname:** ip-10-0-1-221, **Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932
- **VM:** my-ng-influxdb-back-2
  - **Source Server:** **Hostname:** ip-10-0-1-138, **Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ALL |  | 10.0.0.0/16 | inbound * * from 10.0.0.0/16 | Migrated from source |
| 2 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 3 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 5 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my-mig-sg-02

**CSP ID:** sg-0b08a1f2cc70c3840 | **VNet:** my-mig-vnet-01 | **Rules:** 4

**Assigned VMs:**

- **VM:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1
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
| 1 | **Name:** my-mig-vnet-01<br>**ID:** vpc-0201fe7dacc5b3501 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my-mig-subnet-01<br>**ID:** subnet-0b401ce3e440d6a88 | 10.0.1.0/24 | my-mig-vnet-01 |

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

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my-mig-sshkey-01 | tbh9rdeut22if8j5cv6k | 27:de:16:94:80:a4:a8:73:30:d0:f3:da:83:37:7d:43:54:bd:7c:1e | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.3978 |
| Daily | $9.55 |
| Monthly | $286.42 |
| Yearly | $3436.99 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | t3a.small | $16.85 | 5.9% |
| ip-10-0-1-221 (migrated) | t3a.xlarge | $134.78 | 47.1% |
| ip-10-0-1-138 (migrated) | t3a.xlarge | $134.78 | 47.1% |

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

