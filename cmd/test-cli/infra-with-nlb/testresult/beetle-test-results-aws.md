# CM-Beetle test results for AWS (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AWS cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.5+ (83361e8)
- imdl: v0.1.22+ (83361e8)
- CB-Tumblebug: v0.14.0
- CB-Spider: v0.14.0
- CB-MapUI: v0.14.0
- Target CSP: AWS
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: November 25, 2026
- Test Time: 18:14:11 KST
- Test Execution: 2026-11-25 18:14:11 KST

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
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 5.919s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 601ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 53.177s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 35ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 4ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 39ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 3.09s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 3.24s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 17ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 14ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 3m45.646s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.616s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.253s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 15.579s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m0.245s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 7m8.794490939s

*Test executed on November 25, 2026 at 18:14:11 KST (2026-11-25 18:14:11 KST) using CM-Beetle automated test CLI*

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
            "imageId": "ami-046333f9ecafdf406",
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
            "name": "ng-ip-10-0-1-30-62161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3.small",
            "imageId": "ami-046333f9ecafdf406",
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
            "zone": "ap-northeast-2a",
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
          "uid": "tbgkte3l83n06d4u71ap",
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
              "value": "{AttachmentLimitType:shared,EbsCards:null,EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,MaximumEbsAttachments:27,MaximumEbsCards:null,NvmeSupport:required}"
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
              "value": "{BandwidthWeightings:null,ConnectionTrackingConfiguration:{DefaultTcpEstablishedTimeout:432000,DefaultUdpStreamTimeout:180,DefaultUdpTimeout:30},DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSrdSupported:false,EnaSupport:required,EncryptionInTransitSupported:false,FlexibleEnaQueuesSupport:unsupported,Ipv4AddressesPerInterface:15,Ipv4AddressesPerSecondaryInterface:null,Ipv6AddressesPerInterface:15,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:4,MaximumSecondaryNetworkInterfaces:null,NetworkCards:[{AdditionalFlexibleNetworkInterfaces:null,BaselineBandwidthInGbps:1.024,DefaultEnaQueueCountPerInterface:4,MaximumEnaQueueCount:null,MaximumEnaQueueCountPerInterface:null,MaximumNetworkInterfaces:4,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit,PeakBandwidthInGbps:5}],NetworkPerformance:Up to 5 Gigabit,SecondaryNetworkSupported:false}"
            },
            {
              "key": "NitroEnclavesSupport",
              "value": "unsupported"
            },
            {
              "key": "NitroTpmInfo",
              "value": "{SupportedVersions:[2.0]}"
            },
            {
              "key": "NitroTpmSupport",
              "value": "supported"
            },
            {
              "key": "PhcSupport",
              "value": "unsupported"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{Manufacturer:Intel,SupportedArchitectures:[x86_64],SupportedFeatures:null,SustainedClockSpeedInGhz:2.5}"
            },
            {
              "key": "RebootMigrationSupport",
              "value": "supported"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedInRegion",
              "value": "true"
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
          "uid": "tbpb0jh7d27ttp8iu4i4",
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
              "value": "{AttachmentLimitType:shared,EbsCards:null,EbsOptimizedInfo:{BaselineBandwidthInMbps:174,BaselineIops:1000,BaselineThroughputInMBps:21.75,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,MaximumEbsAttachments:27,MaximumEbsCards:null,NvmeSupport:required}"
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
              "value": "{BandwidthWeightings:null,ConnectionTrackingConfiguration:{DefaultTcpEstablishedTimeout:432000,DefaultUdpStreamTimeout:180,DefaultUdpTimeout:30},DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSrdSupported:false,EnaSupport:required,EncryptionInTransitSupported:false,FlexibleEnaQueuesSupport:unsupported,Ipv4AddressesPerInterface:4,Ipv4AddressesPerSecondaryInterface:null,Ipv6AddressesPerInterface:4,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,MaximumSecondaryNetworkInterfaces:null,NetworkCards:[{AdditionalFlexibleNetworkInterfaces:null,BaselineBandwidthInGbps:0.128,DefaultEnaQueueCountPerInterface:2,MaximumEnaQueueCount:null,MaximumEnaQueueCountPerInterface:null,MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit,PeakBandwidthInGbps:5}],NetworkPerformance:Up to 5 Gigabit,SecondaryNetworkSupported:false}"
            },
            {
              "key": "NitroEnclavesSupport",
              "value": "unsupported"
            },
            {
              "key": "NitroTpmInfo",
              "value": "{SupportedVersions:[2.0]}"
            },
            {
              "key": "NitroTpmSupport",
              "value": "supported"
            },
            {
              "key": "PhcSupport",
              "value": "unsupported"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{Manufacturer:Intel,SupportedArchitectures:[x86_64],SupportedFeatures:null,SustainedClockSpeedInGhz:2.5}"
            },
            {
              "key": "RebootMigrationSupport",
              "value": "supported"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedInRegion",
              "value": "true"
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
          "availabilityQuota": {
            "name": "Running On-Demand Standard (A, C, D, H, I, M, R, T, Z) instances",
            "scope": "region",
            "unit": "vCPU",
            "limit": 2421,
            "used": 12,
            "perInstance": 2,
            "fitCount": 1204
          },
          "availabilityCheckedAt": "2026-11-25T09:14:21.347811168Z",
          "availabilityRefreshAfter": "2026-11-25T09:19:21.347811168Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "aws",
          "cspImageName": "ami-046333f9ecafdf406",
          "regionList": [
            "ap-northeast-2"
          ],
          "id": "ami-046333f9ecafdf406",
          "uid": "tbs0b7t23lhm1gte3fgt",
          "name": "ami-046333f9ecafdf406",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "aws-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:25:17 Thu",
          "creationDate": "2026-10-21T06:42:02.000Z",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-00e57b44018849846,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CreationDate",
              "value": "2026-10-21T06:42:02.000Z"
            },
            {
              "key": "DeprecationTime",
              "value": "2028-10-21T06:42:02.000Z"
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
              "value": "ami-046333f9ecafdf406"
            },
            {
              "key": "ImageLocation",
              "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021"
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
              "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021"
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
      "status": "highly-matched",
      "description": "Candidate #2 | highly-matched | 1 NLB(s) | Overall Match Rate: Min=100.0% Max=100.0% Avg=100.0% | VMs: 2 total, 2 matched, 0 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "aws",
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t2.xlarge",
            "imageId": "ami-046333f9ecafdf406",
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
            "name": "ng-ip-10-0-1-30-62161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t8i.small",
            "imageId": "ami-046333f9ecafdf406",
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
          "uid": "tbpisfcrphiqaa3mms09",
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
              "value": "{AttachmentLimitType:shared,EbsCards:null,EbsOptimizedInfo:null,EbsOptimizedSupport:unsupported,EncryptionSupport:supported,MaximumEbsAttachments:40,MaximumEbsCards:null,NvmeSupport:unsupported}"
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
              "value": "{BandwidthWeightings:null,ConnectionTrackingConfiguration:{DefaultTcpEstablishedTimeout:432000,DefaultUdpStreamTimeout:180,DefaultUdpTimeout:30},DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSrdSupported:false,EnaSupport:unsupported,EncryptionInTransitSupported:false,FlexibleEnaQueuesSupport:unsupported,Ipv4AddressesPerInterface:15,Ipv4AddressesPerSecondaryInterface:null,Ipv6AddressesPerInterface:15,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,MaximumSecondaryNetworkInterfaces:null,NetworkCards:[{AdditionalFlexibleNetworkInterfaces:null,BaselineBandwidthInGbps:0.75,DefaultEnaQueueCountPerInterface:1,MaximumEnaQueueCount:null,MaximumEnaQueueCountPerInterface:null,MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Moderate,PeakBandwidthInGbps:1.024}],NetworkPerformance:Moderate,SecondaryNetworkSupported:false}"
            },
            {
              "key": "NitroEnclavesSupport",
              "value": "unsupported"
            },
            {
              "key": "NitroTpmSupport",
              "value": "unsupported"
            },
            {
              "key": "PhcSupport",
              "value": "unsupported"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{Manufacturer:Intel,SupportedArchitectures:[x86_64],SupportedFeatures:null,SustainedClockSpeedInGhz:2.3}"
            },
            {
              "key": "RebootMigrationSupport",
              "value": "supported"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios"
            },
            {
              "key": "SupportedInRegion",
              "value": "true"
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
          "id": "aws+ap-northeast-2+t8i.small",
          "uid": "tbi5qnd02g2u28qp33t7",
          "cspSpecName": "t8i.small",
          "name": "aws+ap-northeast-2+t8i.small",
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
          "costPerHour": 0.0312,
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
              "value": "{AttachmentLimitType:dedicated,EbsCards:null,EbsOptimizedInfo:{BaselineBandwidthInMbps:174,BaselineIops:1500,BaselineThroughputInMBps:21.75,MaximumBandwidthInMbps:5000,MaximumIops:20000,MaximumThroughputInMBps:625},EbsOptimizedSupport:default,EncryptionSupport:supported,MaximumEbsAttachments:32,MaximumEbsCards:null,NvmeSupport:required}"
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
              "value": "t8i.small"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:2048}"
            },
            {
              "key": "NetworkInfo",
              "value": "{BandwidthWeightings:null,ConnectionTrackingConfiguration:{DefaultTcpEstablishedTimeout:350,DefaultUdpStreamTimeout:180,DefaultUdpTimeout:30},DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSrdSupported:false,EnaSupport:required,EncryptionInTransitSupported:true,FlexibleEnaQueuesSupport:unsupported,Ipv4AddressesPerInterface:4,Ipv4AddressesPerSecondaryInterface:0,Ipv6AddressesPerInterface:4,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,MaximumSecondaryNetworkInterfaces:null,NetworkCards:[{AdditionalFlexibleNetworkInterfaces:null,BaselineBandwidthInGbps:0.128,DefaultEnaQueueCountPerInterface:1,MaximumEnaQueueCount:2,MaximumEnaQueueCountPerInterface:1,MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 6.25 Gigabit,PeakBandwidthInGbps:6.25}],NetworkPerformance:Up to 6.25 Gigabit,SecondaryNetworkSupported:false}"
            },
            {
              "key": "NitroEnclavesSupport",
              "value": "unsupported"
            },
            {
              "key": "NitroTpmInfo",
              "value": "{SupportedVersions:[2.0]}"
            },
            {
              "key": "NitroTpmSupport",
              "value": "supported"
            },
            {
              "key": "PhcSupport",
              "value": "supported"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{Manufacturer:Intel,SupportedArchitectures:[x86_64],SupportedFeatures:null,SustainedClockSpeedInGhz:3.9}"
            },
            {
              "key": "RebootMigrationSupport",
              "value": "supported"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedInRegion",
              "value": "true"
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
          "availabilityStatus": "Unverified"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "aws",
          "cspImageName": "ami-046333f9ecafdf406",
          "regionList": [
            "ap-northeast-2"
          ],
          "id": "ami-046333f9ecafdf406",
          "uid": "tbs0b7t23lhm1gte3fgt",
          "name": "ami-046333f9ecafdf406",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "aws-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:25:17 Thu",
          "creationDate": "2026-10-21T06:42:02.000Z",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-00e57b44018849846,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CreationDate",
              "value": "2026-10-21T06:42:02.000Z"
            },
            {
              "key": "DeprecationTime",
              "value": "2028-10-21T06:42:02.000Z"
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
              "value": "ami-046333f9ecafdf406"
            },
            {
              "key": "ImageLocation",
              "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021"
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
              "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021"
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
        "imageId": "ami-046333f9ecafdf406",
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
        "name": "my01-ng-ip-10-0-1-30-62161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
        "connectionName": "aws-ap-northeast-2",
        "specId": "aws+ap-northeast-2+t3.small",
        "imageId": "ami-046333f9ecafdf406",
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
        "zone": "ap-northeast-2a",
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
      "uid": "tbgkte3l83n06d4u71ap",
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
          "value": "{AttachmentLimitType:shared,EbsCards:null,EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,MaximumEbsAttachments:27,MaximumEbsCards:null,NvmeSupport:required}"
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
          "value": "{BandwidthWeightings:null,ConnectionTrackingConfiguration:{DefaultTcpEstablishedTimeout:432000,DefaultUdpStreamTimeout:180,DefaultUdpTimeout:30},DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSrdSupported:false,EnaSupport:required,EncryptionInTransitSupported:false,FlexibleEnaQueuesSupport:unsupported,Ipv4AddressesPerInterface:15,Ipv4AddressesPerSecondaryInterface:null,Ipv6AddressesPerInterface:15,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:4,MaximumSecondaryNetworkInterfaces:null,NetworkCards:[{AdditionalFlexibleNetworkInterfaces:null,BaselineBandwidthInGbps:1.024,DefaultEnaQueueCountPerInterface:4,MaximumEnaQueueCount:null,MaximumEnaQueueCountPerInterface:null,MaximumNetworkInterfaces:4,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit,PeakBandwidthInGbps:5}],NetworkPerformance:Up to 5 Gigabit,SecondaryNetworkSupported:false}"
        },
        {
          "key": "NitroEnclavesSupport",
          "value": "unsupported"
        },
        {
          "key": "NitroTpmInfo",
          "value": "{SupportedVersions:[2.0]}"
        },
        {
          "key": "NitroTpmSupport",
          "value": "supported"
        },
        {
          "key": "PhcSupport",
          "value": "unsupported"
        },
        {
          "key": "PlacementGroupInfo",
          "value": "{SupportedStrategies:[partition,spread]}"
        },
        {
          "key": "ProcessorInfo",
          "value": "{Manufacturer:Intel,SupportedArchitectures:[x86_64],SupportedFeatures:null,SustainedClockSpeedInGhz:2.5}"
        },
        {
          "key": "RebootMigrationSupport",
          "value": "supported"
        },
        {
          "key": "SupportedBootModes",
          "value": "legacy-bios; uefi"
        },
        {
          "key": "SupportedInRegion",
          "value": "true"
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
      "uid": "tbpb0jh7d27ttp8iu4i4",
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
          "value": "{AttachmentLimitType:shared,EbsCards:null,EbsOptimizedInfo:{BaselineBandwidthInMbps:174,BaselineIops:1000,BaselineThroughputInMBps:21.75,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,MaximumEbsAttachments:27,MaximumEbsCards:null,NvmeSupport:required}"
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
          "value": "{BandwidthWeightings:null,ConnectionTrackingConfiguration:{DefaultTcpEstablishedTimeout:432000,DefaultUdpStreamTimeout:180,DefaultUdpTimeout:30},DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSrdSupported:false,EnaSupport:required,EncryptionInTransitSupported:false,FlexibleEnaQueuesSupport:unsupported,Ipv4AddressesPerInterface:4,Ipv4AddressesPerSecondaryInterface:null,Ipv6AddressesPerInterface:4,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,MaximumSecondaryNetworkInterfaces:null,NetworkCards:[{AdditionalFlexibleNetworkInterfaces:null,BaselineBandwidthInGbps:0.128,DefaultEnaQueueCountPerInterface:2,MaximumEnaQueueCount:null,MaximumEnaQueueCountPerInterface:null,MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit,PeakBandwidthInGbps:5}],NetworkPerformance:Up to 5 Gigabit,SecondaryNetworkSupported:false}"
        },
        {
          "key": "NitroEnclavesSupport",
          "value": "unsupported"
        },
        {
          "key": "NitroTpmInfo",
          "value": "{SupportedVersions:[2.0]}"
        },
        {
          "key": "NitroTpmSupport",
          "value": "supported"
        },
        {
          "key": "PhcSupport",
          "value": "unsupported"
        },
        {
          "key": "PlacementGroupInfo",
          "value": "{SupportedStrategies:[partition,spread]}"
        },
        {
          "key": "ProcessorInfo",
          "value": "{Manufacturer:Intel,SupportedArchitectures:[x86_64],SupportedFeatures:null,SustainedClockSpeedInGhz:2.5}"
        },
        {
          "key": "RebootMigrationSupport",
          "value": "supported"
        },
        {
          "key": "SupportedBootModes",
          "value": "legacy-bios; uefi"
        },
        {
          "key": "SupportedInRegion",
          "value": "true"
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
      "availabilityQuota": {
        "name": "Running On-Demand Standard (A, C, D, H, I, M, R, T, Z) instances",
        "scope": "region",
        "unit": "vCPU",
        "limit": 2421,
        "used": 12,
        "perInstance": 2,
        "fitCount": 1204
      },
      "availabilityCheckedAt": "2026-11-25T09:14:21.347811168Z",
      "availabilityRefreshAfter": "2026-11-25T09:19:21.347811168Z"
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "aws",
      "cspImageName": "ami-046333f9ecafdf406",
      "regionList": [
        "ap-northeast-2"
      ],
      "id": "ami-046333f9ecafdf406",
      "uid": "tbs0b7t23lhm1gte3fgt",
      "name": "ami-046333f9ecafdf406",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "aws-ap-northeast-2",
      "infraType": "",
      "fetchedTime": "2026.10.29 04:25:17 Thu",
      "creationDate": "2026-10-21T06:42:02.000Z",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-00e57b44018849846,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CreationDate",
          "value": "2026-10-21T06:42:02.000Z"
        },
        {
          "key": "DeprecationTime",
          "value": "2028-10-21T06:42:02.000Z"
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
          "value": "ami-046333f9ecafdf406"
        },
        {
          "key": "ImageLocation",
          "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021"
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
          "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021"
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
        "nodeGroupId": "my01-ng-influxdb-back"
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:14:55",
      "label": {
        "Name": "tb9kb21aeetl91i0qfdj",
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "cspSshKeyId": "tb0tqsk2ejcg4plsgf7r",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBNboWsTnVyED09nTw9fmdaZMma74gC9pBbxYbP+j7SUJoaUPYZYOrkBxpUFASlmT7TgenIvyORxfdLcKAsZLCy0=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:HWvuPCu5EG6XvkC+builcJdiexIS5oa95B39AnbXTwI",
        "firstUsedAt": "2026-11-25T09:14:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:14:59Z",
          "completedTime": "2026-11-25T09:15:05Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-11-25T09:15:15Z"
        },
        {
          "index": 2,
          "xRequestId": "ps-my01-infra101-tbs7ejcp5lc80ealgd0t",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:15:15Z",
          "completedTime": "2026-11-25T09:15:17Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-123 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:14:59",
      "label": {
        "Name": "tbn8mcdp8a5j8e8918um",
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "cspSshKeyId": "tb0tqsk2ejcg4plsgf7r",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHplfUQJRdN9MBkJVyKqoe738ZbuqE4jMTZggEM5E7v8RIOJ+L2Y9ghlmNZrqJl7XkgUbsv9F9DMX5GfqaCepW4=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:+WDzMpXJ9er2NdtB5q0PeWMSHKCH/WvxjinQ7GUXJmA",
        "firstUsedAt": "2026-11-25T09:15:04Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-25T09:14:59Z",
          "completedTime": "2026-11-25T09:15:04Z",
          "elapsedTime": 5,
          "resultSummary": "Command execution failed",
          "errorMessage": "SSH authentication failed. Please check: 1) private key is valid 2) user 'cb-user' exists on target 3) authorized_keys is properly configured"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:15:13Z",
          "completedTime": "2026-11-25T09:15:14Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "ps-my01-infra101-tbs7ejcp5lc80ealgd0t",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:15:15Z",
          "completedTime": "2026-11-25T09:15:16Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-193 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:14:56",
      "label": {
        "Name": "tbqfgi5tbnlnpkk7ec6e",
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "cspSshKeyId": "tb0tqsk2ejcg4plsgf7r",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDjFWQwFi6ZZQN28mgqCh6yc9/vM5f1bCDukhgA6um773Zo+Ca2RBrtUTkQDzYLBwMbXcoTJhM2MbkWKKbq3xWI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wDGom4hhZTClKY4/tPceD5hA8xmE1K5pMrWDAqeaXNY",
        "firstUsedAt": "2026-11-25T09:15:04Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-25T09:14:59Z",
          "completedTime": "2026-11-25T09:15:04Z",
          "elapsedTime": 5,
          "resultSummary": "Command execution failed",
          "errorMessage": "SSH authentication failed. Please check: 1) private key is valid 2) user 'cb-user' exists on target 3) authorized_keys is properly configured"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:15:13Z",
          "completedTime": "2026-11-25T09:15:15Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "ps-my01-infra101-tbs7ejcp5lc80ealgd0t",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:15:15Z",
          "completedTime": "2026-11-25T09:15:16Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-194 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
        "my01-ng-influxdb-back",
        "my01-ng-ip-10-0-1-30-62161624"
      ],
      "nodeIds": [
        "my01-ng-influxdb-back-1",
        "my01-ng-influxdb-back-2",
        "my01-ng-ip-10-0-1-30-62161624-1"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my01-ng-influxdb-back",
      "representativeNodeId": "my01-ng-influxdb-back-1"
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
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-influxdb-back-2",
            "nodeIp": "52.79.141.27",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-193 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-ip-10-0-1-30-62161624-1",
            "nodeIp": "15.165.15.17",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-194 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-influxdb-back-1",
            "nodeIp": "3.38.252.17",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-123 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postScriptRequestId": "ps-my01-infra101-tbs7ejcp5lc80ealgd0t"
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
          "id": "my02-ng-influxdb-back-1",
          "uid": "tb2mefhc7qac6sg6gmoh",
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
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
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
          "uid": "tb6lhvd6gifs72gf64ok",
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
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
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
          "id": "my02-ng-ip-10-0-1-30-62161624-1",
          "uid": "tbd32ihs5kcb30u4ok9r",
          "name": "my02-ng-ip-10-0-1-30-62161624-1",
          "nodeGroupId": "my02-ng-ip-10-0-1-30-62161624",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
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
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202611110",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my02-mig-sg-02"
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:14:55",
      "label": {
        "Name": "tb9kb21aeetl91i0qfdj",
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "cspSshKeyId": "tb0tqsk2ejcg4plsgf7r",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBNboWsTnVyED09nTw9fmdaZMma74gC9pBbxYbP+j7SUJoaUPYZYOrkBxpUFASlmT7TgenIvyORxfdLcKAsZLCy0=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:HWvuPCu5EG6XvkC+builcJdiexIS5oa95B39AnbXTwI",
        "firstUsedAt": "2026-11-25T09:14:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:14:59Z",
          "completedTime": "2026-11-25T09:15:05Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 2,
          "lastOccurredTime": "2026-11-25T09:15:15Z"
        },
        {
          "index": 2,
          "xRequestId": "ps-my01-infra101-tbs7ejcp5lc80ealgd0t",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:15:15Z",
          "completedTime": "2026-11-25T09:15:17Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-123 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:14:59",
      "label": {
        "Name": "tbn8mcdp8a5j8e8918um",
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "cspSshKeyId": "tb0tqsk2ejcg4plsgf7r",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHplfUQJRdN9MBkJVyKqoe738ZbuqE4jMTZggEM5E7v8RIOJ+L2Y9ghlmNZrqJl7XkgUbsv9F9DMX5GfqaCepW4=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:+WDzMpXJ9er2NdtB5q0PeWMSHKCH/WvxjinQ7GUXJmA",
        "firstUsedAt": "2026-11-25T09:15:04Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-25T09:14:59Z",
          "completedTime": "2026-11-25T09:15:04Z",
          "elapsedTime": 5,
          "resultSummary": "Command execution failed",
          "errorMessage": "SSH authentication failed. Please check: 1) private key is valid 2) user 'cb-user' exists on target 3) authorized_keys is properly configured"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:15:13Z",
          "completedTime": "2026-11-25T09:15:14Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "ps-my01-infra101-tbs7ejcp5lc80ealgd0t",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:15:15Z",
          "completedTime": "2026-11-25T09:15:16Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-193 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-25 09:14:56",
      "label": {
        "Name": "tbqfgi5tbnlnpkk7ec6e",
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
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
      "cspSshKeyId": "tb0tqsk2ejcg4plsgf7r",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDjFWQwFi6ZZQN28mgqCh6yc9/vM5f1bCDukhgA6um773Zo+Ca2RBrtUTkQDzYLBwMbXcoTJhM2MbkWKKbq3xWI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:wDGom4hhZTClKY4/tPceD5hA8xmE1K5pMrWDAqeaXNY",
        "firstUsedAt": "2026-11-25T09:15:04Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-25T09:14:59Z",
          "completedTime": "2026-11-25T09:15:04Z",
          "elapsedTime": 5,
          "resultSummary": "Command execution failed",
          "errorMessage": "SSH authentication failed. Please check: 1) private key is valid 2) user 'cb-user' exists on target 3) authorized_keys is properly configured"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-25T09:15:13Z",
          "completedTime": "2026-11-25T09:15:15Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "ps-my01-infra101-tbs7ejcp5lc80ealgd0t",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-25T09:15:15Z",
          "completedTime": "2026-11-25T09:15:16Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-194 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
        "my01-ng-influxdb-back",
        "my01-ng-ip-10-0-1-30-62161624"
      ],
      "nodeIds": [
        "my01-ng-influxdb-back-1",
        "my01-ng-influxdb-back-2",
        "my01-ng-ip-10-0-1-30-62161624-1"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my01-ng-influxdb-back",
      "representativeNodeId": "my01-ng-influxdb-back-1"
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
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-influxdb-back-2",
            "nodeIp": "52.79.141.27",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-193 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-ip-10-0-1-30-62161624-1",
            "nodeIp": "15.165.15.17",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-194 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-ng-influxdb-back-1",
            "nodeIp": "3.38.252.17",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-123 6.8.0-1066-aws #69~22.04.1-Ubuntu SMP Mon Sep 21 16:20:15 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postScriptRequestId": "ps-my01-infra101-tbs7ejcp5lc80ealgd0t"
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
      "id": "my01-ng-influxdb-back",
      "cspResourceName": "tbflau73msv0h73tfh71",
      "cspResourceId": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de",
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
        "dnsName": "tbflau73msv0h73tfh71-cd5319d52fc061de.elb.ap-northeast-2.amazonaws.com",
        "keyValueList": [
          {
            "key": "DefaultActions",
            "value": "{AuthenticateCognitoConfig:null,AuthenticateOidcConfig:null,FixedResponseConfig:null,ForwardConfig:{TargetGroupStickinessConfig:{DurationSeconds:null,Enabled:false},TargetGroups:[{TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3,Weight:null}]},Order:null,RedirectConfig:null,TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3,Type:forward}"
          },
          {
            "key": "ListenerArn",
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:listener/net/tbflau73msv0h73tfh71/cd5319d52fc061de/c6ff5047cd192180"
          },
          {
            "key": "LoadBalancerArn",
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
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
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
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
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3"
          },
          {
            "key": "TargetGroupName",
            "value": "tbflau73msv0h73tfh71"
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
            "value": "vpc-09aa8042d1c874768"
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
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
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
            "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3"
          },
          {
            "key": "TargetGroupName",
            "value": "tbflau73msv0h73tfh71"
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
            "value": "vpc-09aa8042d1c874768"
          }
        ]
      },
      "createdTime": "2026-11-25T09:15:50.354Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "AvailabilityZones",
          "value": "{LoadBalancerAddresses:null,OutpostId:null,SubnetId:subnet-063b3e54670c7a6f5,ZoneName:ap-northeast-2a}"
        },
        {
          "key": "CanonicalHostedZoneId",
          "value": "ZIBE1TIR4HY56"
        },
        {
          "key": "CreatedTime",
          "value": "2026-11-25T09:15:50.354Z"
        },
        {
          "key": "DNSName",
          "value": "tbflau73msv0h73tfh71-cd5319d52fc061de.elb.ap-northeast-2.amazonaws.com"
        },
        {
          "key": "IpAddressType",
          "value": "ipv4"
        },
        {
          "key": "LoadBalancerArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
        },
        {
          "key": "LoadBalancerName",
          "value": "tbflau73msv0h73tfh71"
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
          "value": "vpc-09aa8042d1c874768"
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
    "cspResourceName": "tbflau73msv0h73tfh71",
    "cspResourceId": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de",
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
      "dnsName": "tbflau73msv0h73tfh71-cd5319d52fc061de.elb.ap-northeast-2.amazonaws.com",
      "keyValueList": [
        {
          "key": "DefaultActions",
          "value": "{AuthenticateCognitoConfig:null,AuthenticateOidcConfig:null,FixedResponseConfig:null,ForwardConfig:{TargetGroupStickinessConfig:{DurationSeconds:null,Enabled:false},TargetGroups:[{TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3,Weight:null}]},Order:null,RedirectConfig:null,TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3,Type:forward}"
        },
        {
          "key": "ListenerArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:listener/net/tbflau73msv0h73tfh71/cd5319d52fc061de/c6ff5047cd192180"
        },
        {
          "key": "LoadBalancerArn",
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
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
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
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
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3"
        },
        {
          "key": "TargetGroupName",
          "value": "tbflau73msv0h73tfh71"
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
          "value": "vpc-09aa8042d1c874768"
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
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
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
          "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3"
        },
        {
          "key": "TargetGroupName",
          "value": "tbflau73msv0h73tfh71"
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
          "value": "vpc-09aa8042d1c874768"
        }
      ]
    },
    "createdTime": "2026-11-25T09:15:50.354Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "AvailabilityZones",
        "value": "{LoadBalancerAddresses:null,OutpostId:null,SubnetId:subnet-063b3e54670c7a6f5,ZoneName:ap-northeast-2a}"
      },
      {
        "key": "CanonicalHostedZoneId",
        "value": "ZIBE1TIR4HY56"
      },
      {
        "key": "CreatedTime",
        "value": "2026-11-25T09:15:50.354Z"
      },
      {
        "key": "DNSName",
        "value": "tbflau73msv0h73tfh71-cd5319d52fc061de.elb.ap-northeast-2.amazonaws.com"
      },
      {
        "key": "IpAddressType",
        "value": "ipv4"
      },
      {
        "key": "LoadBalancerArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
      },
      {
        "key": "LoadBalancerName",
        "value": "tbflau73msv0h73tfh71"
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
        "value": "vpc-09aa8042d1c874768"
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
  "cspResourceName": "tbflau73msv0h73tfh71",
  "cspResourceId": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de",
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
    "dnsName": "tbflau73msv0h73tfh71-cd5319d52fc061de.elb.ap-northeast-2.amazonaws.com",
    "keyValueList": [
      {
        "key": "DefaultActions",
        "value": "{AuthenticateCognitoConfig:null,AuthenticateOidcConfig:null,FixedResponseConfig:null,ForwardConfig:{TargetGroupStickinessConfig:{DurationSeconds:null,Enabled:false},TargetGroups:[{TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3,Weight:null}]},Order:null,RedirectConfig:null,TargetGroupArn:arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3,Type:forward}"
      },
      {
        "key": "ListenerArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:listener/net/tbflau73msv0h73tfh71/cd5319d52fc061de/c6ff5047cd192180"
      },
      {
        "key": "LoadBalancerArn",
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
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
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
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
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3"
      },
      {
        "key": "TargetGroupName",
        "value": "tbflau73msv0h73tfh71"
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
        "value": "vpc-09aa8042d1c874768"
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
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
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
        "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:targetgroup/tbflau73msv0h73tfh71/a13294ea57c528d3"
      },
      {
        "key": "TargetGroupName",
        "value": "tbflau73msv0h73tfh71"
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
        "value": "vpc-09aa8042d1c874768"
      }
    ]
  },
  "createdTime": "2026-11-25T09:15:50.354Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "AvailabilityZones",
      "value": "{LoadBalancerAddresses:null,OutpostId:null,SubnetId:subnet-063b3e54670c7a6f5,ZoneName:ap-northeast-2a}"
    },
    {
      "key": "CanonicalHostedZoneId",
      "value": "ZIBE1TIR4HY56"
    },
    {
      "key": "CreatedTime",
      "value": "2026-11-25T09:15:50.354Z"
    },
    {
      "key": "DNSName",
      "value": "tbflau73msv0h73tfh71-cd5319d52fc061de.elb.ap-northeast-2.amazonaws.com"
    },
    {
      "key": "IpAddressType",
      "value": "ipv4"
    },
    {
      "key": "LoadBalancerArn",
      "value": "arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de"
    },
    {
      "key": "LoadBalancerName",
      "value": "tbflau73msv0h73tfh71"
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
      "value": "vpc-09aa8042d1c874768"
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

**Generated At:** 2026-11-25 09:19:41

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
| **Total NLBs** | 1 |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| t3.xlarge | 4 | 16.0 | - | x86_64 |  | $0.2080 | 2 |
| t3.small | 2 | 2.0 | - | x86_64 |  | $0.0260 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ami-046333f9ecafdf406 | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 | Ubuntu 22.04 | Linux/UNIX | x86_64 | ebs | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my01-ng-influxdb-back-1 | i-0bbb8c166a68d3f41 | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 3.38.252.17<br>**Private IP:** 10.0.1.123<br>**SGs:** my01-mig-sg-01<br>**SSH:** my01-mig-sshkey-01 |
| my01-ng-influxdb-back-2 | i-0f6efd5a8c05d59e1 | Running | 4 vCPU, 16.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 52.79.141.27<br>**Private IP:** 10.0.1.193<br>**SGs:** my01-mig-sg-01<br>**SSH:** my01-mig-sshkey-01 |
| my01-ng-ip-10-0-1-30-62161624-1 | i-07196fca4d742d8ca | Running | 2 vCPU, 2.0 GiB | ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 (ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021) | **VNet:** my01-mig-vnet-01<br>**Subnet:** my01-mig-subnet-01<br>**Public IP:** 15.165.15.17<br>**Private IP:** 10.0.1.194<br>**SGs:** my01-mig-sg-02<br>**SSH:** my01-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my01-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my01-mig-vnet-01 |
| **CSP VNet ID** | vpc-09aa8042d1c874768 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | aws-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my01-mig-subnet-01 | subnet-063b3e54670c7a6f5 | 10.0.1.0/24 | ap-northeast-2a |

### Network Load Balancers (NLB)

#### NLB: my01-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my01-ng-influxdb-back |
| **CSP NLB Name** | tbflau73msv0h73tfh71 |
| **CSP NLB ID** | arn:aws:elasticloadbalancing:ap-northeast-2:635484366616:loadbalancer/net/tbflau73msv0h73tfh71/cd5319d52fc061de |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | aws-ap-northeast-2 |
| **Listener** | TCP:9999 (DNS: tbflau73msv0h73tfh71-cd5319d52fc061de.elb.ap-northeast-2.amazonaws.com) |
| **Target Group** | TCP:8086 (Node Group: my01-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my01-ng-influxdb-back-1 | i-0bbb8c166a68d3f41 | 10.0.1.123 | 3.38.252.17 | Running |
| my01-ng-influxdb-back-2 | i-0f6efd5a8c05d59e1 | 10.0.1.193 | 52.79.141.27 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my01-mig-sshkey-01 | tb0tqsk2ejcg4plsgf7r |  | 5c:d1:dc:58:54:a1:aa:ed:e2:66:41:59:4b:c6:e7:b4:a3:99:86:2a |

### Security Groups

#### Security Group: my01-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my01-mig-sg-01 |
| **CSP Security Group ID** | sg-04c56c75d85c525e4 |
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
| **CSP Security Group ID** | sg-0c81f2afb6f7d9926 |
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
| my01-ng-influxdb-back-1 | t3.xlarge | $0.2080 | $149.76 |
| my01-ng-influxdb-back-2 | t3.xlarge | $0.2080 | $149.76 |
| my01-ng-ip-10-0-1-30-62161624-1 | t3.small | $0.0260 | $18.72 |




### Test Case 13: Migration Report

#### 13.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 13.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-11-25 09:19:46*

---

## 📊 Migration Summary

**Target Cloud:** AWS

**Target Region:** ap-northeast-2

**Namespace:** mig01 | **Infra ID:** my01-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**NLB Migration Status:** Completed

**Total Load Balancers:** 1

**Migrated Load Balancers:** 1

**💰 Estimated Monthly Cost:** $318.24 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | t3.xlarge, t3.small |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my01-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my01-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 10 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |
| 8 | **Network Load Balancer** | 1 | ✅ Created | my01-ng-influxdb-back (TCP:9999 -> TCP:8086) |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my01-ng-influxdb-back-1<br>**VM ID:** i-0bbb8c166a68d3f41<br>**Label(sourceMachineId):** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 |
| 2 | **VM Name:** my01-ng-influxdb-back-2<br>**VM ID:** i-0f6efd5a8c05d59e1<br>**Label(sourceMachineId):** ec288dd0-c6fa-8a49-2f60-bc898311febf | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf |
| 3 | **VM Name:** my01-ng-ip-10-0-1-30-62161624-1<br>**VM ID:** i-07196fca4d742d8ca<br>**Label(sourceMachineId):** ec268ed7-821e-9d73-e79f-961262161624 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my01-ng-influxdb-back-1 | **Spec ID:** t3.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 30 GB |
| 2 | my01-ng-influxdb-back-2 | **Spec ID:** t3.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 8 GB<br>**Root Disk:** 30 GB |
| 3 | my01-ng-ip-10-0-1-30-62161624-1 | **Spec ID:** t3.small<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 8 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my01-ng-influxdb-back-1 | **Image ID:** ami-046333f9ecafdf406<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | my01-ng-influxdb-back-2 | **Image ID:** ami-046333f9ecafdf406<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | my01-ng-ip-10-0-1-30-62161624-1 | **Image ID:** ami-046333f9ecafdf406<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20261021 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 10 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my01-mig-sg-01

**CSP ID:** sg-04c56c75d85c525e4 | **VNet:** my01-mig-vnet-01 | **Rules:** 6

**Assigned VMs:**

- **VM:** my01-ng-influxdb-back-1
  - **Source Server:** **Hostname:** ip-10-0-1-221, **Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932
- **VM:** my01-ng-influxdb-back-2
  - **Source Server:** **Hostname:** ip-10-0-1-138, **Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 9999 | 0.0.0.0/0 | - | Created by system |
| 2 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 4 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 5 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 6 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my01-mig-sg-02

**CSP ID:** sg-0c81f2afb6f7d9926 | **VNet:** my01-mig-vnet-01 | **Rules:** 4

**Assigned VMs:**

- **VM:** my01-ng-ip-10-0-1-30-62161624-1
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
| 1 | **Name:** my01-mig-vnet-01<br>**ID:** vpc-09aa8042d1c874768 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my01-mig-subnet-01<br>**ID:** subnet-063b3e54670c7a6f5 | 10.0.1.0/24 | my01-mig-vnet-01 |

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

### 1. influxdb_back (haproxy) → my01-ng-influxdb-back

**Configuration Comparison:**

| Property | Source (On-Premise) | Target (Cloud NLB) | Note |
|----------|---------------------|--------------------|------|
| **Load Balancer / Name** | haproxy (Host: ip-10-0-1-30) | my01-ng-influxdb-back (CSP ID: tbflau73msv0h73tfh71) | Replaced by managed NLB |
| **Type / Scope** | Software LB | PUBLIC / REGION | Cloud-native NLB |
| **Frontend Listener** | *:9999 (TCP) | TCP:9999 | DNS: tbflau73msv0h73tfh71-cd5319d52fc061de.elb.ap-northeast-2.amazonaws.com |
| **Backend / Target Group** | influxdb_back (TCP:8086, roundrobin) | my01-ng-influxdb-back (TCP:8086, 2 nodes) | Target Group attached to NodeGroup |
| **Health Check** | Enabled (Interval: 10s, Timeout: 10s, Threshold: 3) | TCP:8086 (Interval: 10s, Timeout: 10s, Threshold: 3) | Preserved |

**Backend Node Mappings:**

| No. | Source Backend Server | Source Endpoint | Target Node | Target CSP VM ID | Target Private IP | Target Public IP | Target Port | Status |
|-----|-----------------------|-----------------|-------------|------------------|-------------------|------------------|-------------|--------|
| 1 | influx1 (ip-10-0-1-221) | 10.0.1.221:8086 | my01-ng-influxdb-back-1 | i-0bbb8c166a68d3f41 | 10.0.1.123 | 3.38.252.17 | 8086 | Running |
| 2 | influx2 (ip-10-0-1-138) | 10.0.1.138:8086 | my01-ng-influxdb-back-2 | i-0f6efd5a8c05d59e1 | 10.0.1.193 | 52.79.141.27 | 8086 | Running |

**Migration Notes:**
- On-premise haproxy on ip-10-0-1-30 (migrated to my01-ng-ip-10-0-1-30-62161624-1) is replaced by cloud-managed NLB.

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my01-mig-sshkey-01 | tb0tqsk2ejcg4plsgf7r | 5c:d1:dc:58:54:a1:aa:ed:e2:66:41:59:4b:c6:e7:b4:a3:99:86:2a | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.4420 |
| Daily | $10.61 |
| Monthly | $318.24 |
| Yearly | $3818.88 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | t3.small | $18.72 | 5.9% |
| ip-10-0-1-221 (migrated) | t3.xlarge | $149.76 | 47.1% |
| ip-10-0-1-138 (migrated) | t3.xlarge | $149.76 | 47.1% |

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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my01-infra101)",
  "success": true
}
```

