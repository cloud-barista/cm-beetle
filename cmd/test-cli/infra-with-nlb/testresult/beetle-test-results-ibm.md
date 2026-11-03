# CM-Beetle test results for IBM (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with IBM cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.3+ (d67a8e3)
- imdl: v0.1.19+ (d67a8e3)
- CB-Tumblebug: v0.13.8
- CB-Spider: v0.13.20
- CB-MapUI: v0.13.14
- Target CSP: IBM
- Target Region: au-syd
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: November 3, 2026
- Test Time: 16:42:03 KST
- Test Execution: 2026-11-03 16:42:03 KST

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

## Test result for IBM

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 4.092s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 401ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 3m31.66s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 73ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 7ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 20ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 14.066s | Pass |
| 8 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 7m46.33s | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 9ms | Pass |
| 10 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 7ms | Pass |
| 11 | NLB Load Balancing Verification | ✅ **PASS** | 2m6.077s | Pass |
| 12 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.209s | Pass |
| 13 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.206s | Pass |
| 14 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 3m35.552s | Pass |
| 15 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m15.12s | Pass |

**Overall Result**: 15/15 tests passed ✅

**Total Duration**: 19m34.185366961s

*Test executed on November 3, 2026 at 16:42:03 KST (2026-11-03 16:42:03 KST) using CM-Beetle automated test CLI*

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
  "desiredCsp": "ibm",
  "desiredRegion": "au-syd",
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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=86.1% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "ibm",
        "region": "au-syd"
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bx2-4x16",
            "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          },
          {
            "name": "ng-ec268ed7-821e-9d73-e79f-961262161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=83.3%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+cx2-2x4",
            "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "mig-vnet-01",
        "connectionName": "ibm-au-syd",
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
        "connectionName": "ibm-au-syd",
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
          "id": "ibm+au-syd+bx2-4x16",
          "uid": "tbjjb50rgj6kg43q3quq",
          "cspSpecName": "bx2-4x16",
          "name": "ibm+au-syd+bx2-4x16",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.247,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:8000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bx2-4x16"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:16}"
            },
            {
              "key": "Name",
              "value": "bx2-4x16"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:dependent}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:2000,max:7500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:100}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:4}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:fixed,value:intel}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:weighted,type:enum,values:[weighted]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "au-syd-1",
            "au-syd-2",
            "au-syd-3"
          ],
          "availabilityCheckedAt": "2026-11-03T07:42:11.132146085Z",
          "availabilityRefreshAfter": "2026-11-03T08:42:11.132146085Z"
        },
        {
          "id": "ibm+au-syd+cx2-2x4",
          "uid": "tbsqv9htrvso0ldtlmlv",
          "cspSpecName": "cx2-2x4",
          "name": "ibm+au-syd+cx2-2x4",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 4,
          "diskSizeGB": -1,
          "costPerHour": 0.109,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:4000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "compute"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/cx2-2x4"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:4}"
            },
            {
              "key": "Name",
              "value": "cx2-2x4"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:dependent}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:1000,max:3500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:100}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:fixed,value:intel}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:weighted,type:enum,values:[weighted]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "au-syd-1",
            "au-syd-2",
            "au-syd-3"
          ],
          "availabilityCheckedAt": "2026-11-03T07:42:25.039010893Z",
          "availabilityRefreshAfter": "2026-11-03T08:42:25.039010893Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "ibm",
          "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "regionList": [
            "au-syd"
          ],
          "id": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "uid": "tb2menjfr1ukaehvmfj0",
          "name": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ibm-au-syd",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:25:05 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)",
          "osDiskType": "NA",
          "osDiskSizeGB": 10,
          "imageStatus": "Available",
          "details": [
            {
              "key": "AllowedUse",
              "value": "{api_version:2024-11-28,bare_metal_server:true,instance:true}"
            },
            {
              "key": "CatalogOffering",
              "value": "{managed:false}"
            },
            {
              "key": "CreatedAt",
              "value": "2026-10-27T05:54:41.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-6245bce7-1353-421f-9eb8-47aa0ed18c50"
            },
            {
              "key": "Encryption",
              "value": "none"
            },
            {
              "key": "File",
              "value": "{checksums:{sha256:2ed37d7923d8b81d6f9b26f2a43dad36238c7a71105356458e9e7cbf121c6c93},size:1}"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/images/r026-6245bce7-1353-421f-9eb8-47aa0ed18c50"
            },
            {
              "key": "ID",
              "value": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50"
            },
            {
              "key": "MinimumProvisionedSize",
              "value": "10"
            },
            {
              "key": "Name",
              "value": "ibm-ubuntu-22-04-5-minimal-amd64-20"
            },
            {
              "key": "OperatingSystem",
              "value": "{allow_user_image_creation:true,architecture:amd64,dedicated_host_only:false,display_name:Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64),family:Ubuntu Linux,href:https://au-syd.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-22-04-amd64,name:ubuntu-22-04-amd64,user_data_format:cloud_init,vendor:Canonical,version:22.04 LTS Jammy Jellyfish Minimal Install}"
            },
            {
              "key": "Remote",
              "value": "{account:{id:811f8abfbd32425597dc7ba40da98fa6,resource_type:account}}"
            },
            {
              "key": "ResourceGroup",
              "value": "{href:https://resource-controller.cloud.ibm.com/v1/resource_groups/5807b5832a8741179b2e06ca2d2b3b96,id:5807b5832a8741179b2e06ca2d2b3b96,name:Default}"
            },
            {
              "key": "ResourceType",
              "value": "image"
            },
            {
              "key": "Status",
              "value": "available"
            },
            {
              "key": "UserDataFormat",
              "value": "cloud_init"
            },
            {
              "key": "Visibility",
              "value": "public"
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
          "connectionName": "ibm-au-syd",
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
          "connectionName": "ibm-au-syd",
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
            "timeout": 5
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
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=86.1% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
      "targetCloud": {
        "csp": "ibm",
        "region": "au-syd"
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
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bx2d-4x16",
            "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          },
          {
            "name": "ng-ec268ed7-821e-9d73-e79f-961262161624",
            "nodeGroupSize": 1,
            "label": {
              "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=83.3%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+cx2d-2x4",
            "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "mig-vnet-01",
        "connectionName": "ibm-au-syd",
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
        "connectionName": "ibm-au-syd",
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
          "id": "ibm+au-syd+bx2d-4x16",
          "uid": "tb9jbn59qlud9ku1n9cu",
          "cspSpecName": "bx2d-4x16",
          "name": "ibm+au-syd+bx2d-4x16",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.268,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:8000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Disks",
              "value": "{quantity:{type:fixed,value:1},size:{type:fixed,value:150},supported_interface_types:{default:virtio_blk,type:enum,values:[virtio_blk]}}"
            },
            {
              "key": "Family",
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bx2d-4x16"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:16}"
            },
            {
              "key": "Name",
              "value": "bx2d-4x16"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:dependent}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:2000,max:7500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:100}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:4}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:fixed,value:intel}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:weighted,type:enum,values:[weighted]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "au-syd-1",
            "au-syd-2",
            "au-syd-3"
          ],
          "availabilityCheckedAt": "2026-11-03T07:42:10.276047516Z",
          "availabilityRefreshAfter": "2026-11-03T08:42:10.276047516Z"
        },
        {
          "id": "ibm+au-syd+cx2d-2x4",
          "uid": "tbja065dk3j3jlgu59mo",
          "cspSpecName": "cx2d-2x4",
          "name": "ibm+au-syd+cx2d-2x4",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 4,
          "diskSizeGB": -1,
          "costPerHour": 0.12,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:4000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Disks",
              "value": "{quantity:{type:fixed,value:1},size:{type:fixed,value:75},supported_interface_types:{default:virtio_blk,type:enum,values:[virtio_blk]}}"
            },
            {
              "key": "Family",
              "value": "compute"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/cx2d-2x4"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:4}"
            },
            {
              "key": "Name",
              "value": "cx2d-2x4"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:dependent}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:1000,max:3500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:100}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:fixed,value:intel}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:weighted,type:enum,values:[weighted]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ],
          "availabilityStatus": "Available",
          "availableZones": [
            "au-syd-1",
            "au-syd-2",
            "au-syd-3"
          ],
          "availabilityCheckedAt": "2026-11-03T07:42:25.048686593Z",
          "availabilityRefreshAfter": "2026-11-03T08:42:25.048686593Z"
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "ibm",
          "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "regionList": [
            "au-syd"
          ],
          "id": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "uid": "tb2menjfr1ukaehvmfj0",
          "name": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ibm-au-syd",
          "infraType": "",
          "fetchedTime": "2026.10.29 04:25:05 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)",
          "osDiskType": "NA",
          "osDiskSizeGB": 10,
          "imageStatus": "Available",
          "details": [
            {
              "key": "AllowedUse",
              "value": "{api_version:2024-11-28,bare_metal_server:true,instance:true}"
            },
            {
              "key": "CatalogOffering",
              "value": "{managed:false}"
            },
            {
              "key": "CreatedAt",
              "value": "2026-10-27T05:54:41.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-6245bce7-1353-421f-9eb8-47aa0ed18c50"
            },
            {
              "key": "Encryption",
              "value": "none"
            },
            {
              "key": "File",
              "value": "{checksums:{sha256:2ed37d7923d8b81d6f9b26f2a43dad36238c7a71105356458e9e7cbf121c6c93},size:1}"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/images/r026-6245bce7-1353-421f-9eb8-47aa0ed18c50"
            },
            {
              "key": "ID",
              "value": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50"
            },
            {
              "key": "MinimumProvisionedSize",
              "value": "10"
            },
            {
              "key": "Name",
              "value": "ibm-ubuntu-22-04-5-minimal-amd64-20"
            },
            {
              "key": "OperatingSystem",
              "value": "{allow_user_image_creation:true,architecture:amd64,dedicated_host_only:false,display_name:Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64),family:Ubuntu Linux,href:https://au-syd.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-22-04-amd64,name:ubuntu-22-04-amd64,user_data_format:cloud_init,vendor:Canonical,version:22.04 LTS Jammy Jellyfish Minimal Install}"
            },
            {
              "key": "Remote",
              "value": "{account:{id:811f8abfbd32425597dc7ba40da98fa6,resource_type:account}}"
            },
            {
              "key": "ResourceGroup",
              "value": "{href:https://resource-controller.cloud.ibm.com/v1/resource_groups/5807b5832a8741179b2e06ca2d2b3b96,id:5807b5832a8741179b2e06ca2d2b3b96,name:Default}"
            },
            {
              "key": "ResourceType",
              "value": "image"
            },
            {
              "key": "Status",
              "value": "available"
            },
            {
              "key": "UserDataFormat",
              "value": "cloud_init"
            },
            {
              "key": "Visibility",
              "value": "public"
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
          "connectionName": "ibm-au-syd",
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
          "connectionName": "ibm-au-syd",
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
            "timeout": 5
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
  "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=86.1% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
  "targetCloud": {
    "csp": "ibm",
    "region": "au-syd"
  },
  "targetInfra": {
    "name": "my06-infra101",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "NLB-aware recommended infrastructure for cloud migration",
    "nodeGroups": [
      {
        "name": "my06-ng-influxdb-back",
        "nodeGroupSize": 2,
        "label": {
          "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "nlbBackend": "influxdb_back"
        },
        "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
        "connectionName": "ibm-au-syd",
        "specId": "ibm+au-syd+bx2-4x16",
        "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
        "vNetId": "my06-mig-vnet-01",
        "subnetId": "my06-mig-subnet-01",
        "securityGroupIds": [
          "my06-mig-sg-01"
        ],
        "sshKeyId": "my06-mig-sshkey-01",
        "rootDiskSize": 100,
        "dataDiskIds": null
      },
      {
        "name": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeGroupSize": 1,
        "label": {
          "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=83.3%",
        "connectionName": "ibm-au-syd",
        "specId": "ibm+au-syd+cx2-2x4",
        "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
        "vNetId": "my06-mig-vnet-01",
        "subnetId": "my06-mig-subnet-01",
        "securityGroupIds": [
          "my06-mig-sg-02"
        ],
        "sshKeyId": "my06-mig-sshkey-01",
        "rootDiskSize": 100,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "my06-mig-vnet-01",
    "connectionName": "ibm-au-syd",
    "cidrBlock": "10.0.0.0/21",
    "subnetInfoList": [
      {
        "name": "my06-mig-subnet-01",
        "ipv4_CIDR": "10.0.1.0/24",
        "description": "a recommended subnet for migration"
      }
    ],
    "description": "a recommended vNet for migration"
  },
  "targetSshKey": {
    "name": "my06-mig-sshkey-01",
    "connectionName": "ibm-au-syd",
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
      "id": "ibm+au-syd+bx2-4x16",
      "uid": "tbjjb50rgj6kg43q3quq",
      "cspSpecName": "bx2-4x16",
      "name": "ibm+au-syd+bx2-4x16",
      "namespace": "system",
      "connectionName": "ibm-au-syd",
      "providerName": "ibm",
      "regionName": "au-syd",
      "regionLatitude": -33.86882,
      "regionLongitude": 151.209296,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 16,
      "diskSizeGB": -1,
      "costPerHour": 0.247,
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
          "key": "AvailabilityClass",
          "value": "{default:standard,type:enum,values:[standard]}"
        },
        {
          "key": "Bandwidth",
          "value": "{type:fixed,value:8000}"
        },
        {
          "key": "ClusterNetworkAttachmentCount",
          "value": "{type:enum,values:[0]}"
        },
        {
          "key": "ConfidentialComputeModes",
          "value": "{default:disabled,type:enum,values:[disabled]}"
        },
        {
          "key": "Family",
          "value": "balanced"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bx2-4x16"
        },
        {
          "key": "Memory",
          "value": "{type:fixed,value:16}"
        },
        {
          "key": "Name",
          "value": "bx2-4x16"
        },
        {
          "key": "NetworkAttachmentCount",
          "value": "{max:5,min:1,type:range}"
        },
        {
          "key": "NetworkBandwidthMode",
          "value": "{type:fixed,value:divided}"
        },
        {
          "key": "NetworkInterfaceCount",
          "value": "{max:5,min:1,type:range}"
        },
        {
          "key": "NumaCount",
          "value": "{type:dependent}"
        },
        {
          "key": "OsArchitecture",
          "value": "{default:amd64,type:enum,values:[amd64]}"
        },
        {
          "key": "PortSpeed",
          "value": "{type:fixed,value:25000}"
        },
        {
          "key": "ReservationTerms",
          "value": "{type:enum,values:[one_year,three_year]}"
        },
        {
          "key": "ResourceType",
          "value": "instance_profile"
        },
        {
          "key": "SecureBootModes",
          "value": "{default:false,type:enum,values:[false]}"
        },
        {
          "key": "Status",
          "value": "current"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "{type:range,default:2000,max:7500,min:500,step:1}"
        },
        {
          "key": "VcpuArchitecture",
          "value": "{type:fixed,value:amd64}"
        },
        {
          "key": "VcpuBurstLimit",
          "value": "{type:fixed,value:100}"
        },
        {
          "key": "VcpuCount",
          "value": "{type:fixed,value:4}"
        },
        {
          "key": "VcpuManufacturer",
          "value": "{type:fixed,value:intel}"
        },
        {
          "key": "VcpuPercentage",
          "value": "{default:100,type:enum,values:[100]}"
        },
        {
          "key": "VolumeBandwidthQosModes",
          "value": "{default:weighted,type:enum,values:[weighted]}"
        },
        {
          "key": "Zones",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
        }
      ],
      "availabilityStatus": "Available",
      "availableZones": [
        "au-syd-1",
        "au-syd-2",
        "au-syd-3"
      ],
      "availabilityCheckedAt": "2026-11-03T07:42:11.132146085Z",
      "availabilityRefreshAfter": "2026-11-03T08:42:11.132146085Z"
    },
    {
      "id": "ibm+au-syd+cx2-2x4",
      "uid": "tbsqv9htrvso0ldtlmlv",
      "cspSpecName": "cx2-2x4",
      "name": "ibm+au-syd+cx2-2x4",
      "namespace": "system",
      "connectionName": "ibm-au-syd",
      "providerName": "ibm",
      "regionName": "au-syd",
      "regionLatitude": -33.86882,
      "regionLongitude": 151.209296,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 4,
      "diskSizeGB": -1,
      "costPerHour": 0.109,
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
          "key": "AvailabilityClass",
          "value": "{default:standard,type:enum,values:[standard]}"
        },
        {
          "key": "Bandwidth",
          "value": "{type:fixed,value:4000}"
        },
        {
          "key": "ClusterNetworkAttachmentCount",
          "value": "{type:enum,values:[0]}"
        },
        {
          "key": "ConfidentialComputeModes",
          "value": "{default:disabled,type:enum,values:[disabled]}"
        },
        {
          "key": "Family",
          "value": "compute"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/cx2-2x4"
        },
        {
          "key": "Memory",
          "value": "{type:fixed,value:4}"
        },
        {
          "key": "Name",
          "value": "cx2-2x4"
        },
        {
          "key": "NetworkAttachmentCount",
          "value": "{max:5,min:1,type:range}"
        },
        {
          "key": "NetworkBandwidthMode",
          "value": "{type:fixed,value:divided}"
        },
        {
          "key": "NetworkInterfaceCount",
          "value": "{max:5,min:1,type:range}"
        },
        {
          "key": "NumaCount",
          "value": "{type:dependent}"
        },
        {
          "key": "OsArchitecture",
          "value": "{default:amd64,type:enum,values:[amd64]}"
        },
        {
          "key": "PortSpeed",
          "value": "{type:fixed,value:25000}"
        },
        {
          "key": "ReservationTerms",
          "value": "{type:enum,values:[one_year,three_year]}"
        },
        {
          "key": "ResourceType",
          "value": "instance_profile"
        },
        {
          "key": "SecureBootModes",
          "value": "{default:false,type:enum,values:[false]}"
        },
        {
          "key": "Status",
          "value": "current"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "{type:range,default:1000,max:3500,min:500,step:1}"
        },
        {
          "key": "VcpuArchitecture",
          "value": "{type:fixed,value:amd64}"
        },
        {
          "key": "VcpuBurstLimit",
          "value": "{type:fixed,value:100}"
        },
        {
          "key": "VcpuCount",
          "value": "{type:fixed,value:2}"
        },
        {
          "key": "VcpuManufacturer",
          "value": "{type:fixed,value:intel}"
        },
        {
          "key": "VcpuPercentage",
          "value": "{default:100,type:enum,values:[100]}"
        },
        {
          "key": "VolumeBandwidthQosModes",
          "value": "{default:weighted,type:enum,values:[weighted]}"
        },
        {
          "key": "Zones",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
        }
      ],
      "availabilityStatus": "Available",
      "availableZones": [
        "au-syd-1",
        "au-syd-2",
        "au-syd-3"
      ],
      "availabilityCheckedAt": "2026-11-03T07:42:25.039010893Z",
      "availabilityRefreshAfter": "2026-11-03T08:42:25.039010893Z"
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "ibm",
      "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "regionList": [
        "au-syd"
      ],
      "id": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "uid": "tb2menjfr1ukaehvmfj0",
      "name": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "ibm-au-syd",
      "infraType": "",
      "fetchedTime": "2026.10.29 04:25:05 Thu",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)",
      "osDiskType": "NA",
      "osDiskSizeGB": 10,
      "imageStatus": "Available",
      "details": [
        {
          "key": "AllowedUse",
          "value": "{api_version:2024-11-28,bare_metal_server:true,instance:true}"
        },
        {
          "key": "CatalogOffering",
          "value": "{managed:false}"
        },
        {
          "key": "CreatedAt",
          "value": "2026-10-27T05:54:41.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-6245bce7-1353-421f-9eb8-47aa0ed18c50"
        },
        {
          "key": "Encryption",
          "value": "none"
        },
        {
          "key": "File",
          "value": "{checksums:{sha256:2ed37d7923d8b81d6f9b26f2a43dad36238c7a71105356458e9e7cbf121c6c93},size:1}"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/images/r026-6245bce7-1353-421f-9eb8-47aa0ed18c50"
        },
        {
          "key": "ID",
          "value": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50"
        },
        {
          "key": "MinimumProvisionedSize",
          "value": "10"
        },
        {
          "key": "Name",
          "value": "ibm-ubuntu-22-04-5-minimal-amd64-20"
        },
        {
          "key": "OperatingSystem",
          "value": "{allow_user_image_creation:true,architecture:amd64,dedicated_host_only:false,display_name:Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64),family:Ubuntu Linux,href:https://au-syd.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-22-04-amd64,name:ubuntu-22-04-amd64,user_data_format:cloud_init,vendor:Canonical,version:22.04 LTS Jammy Jellyfish Minimal Install}"
        },
        {
          "key": "Remote",
          "value": "{account:{id:811f8abfbd32425597dc7ba40da98fa6,resource_type:account}}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v1/resource_groups/5807b5832a8741179b2e06ca2d2b3b96,id:5807b5832a8741179b2e06ca2d2b3b96,name:Default}"
        },
        {
          "key": "ResourceType",
          "value": "image"
        },
        {
          "key": "Status",
          "value": "available"
        },
        {
          "key": "UserDataFormat",
          "value": "cloud_init"
        },
        {
          "key": "Visibility",
          "value": "public"
        }
      ],
      "systemLabel": "",
      "description": "",
      "commandHistory": null
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "my06-mig-sg-01",
      "connectionName": "ibm-au-syd",
      "vNetId": "my06-mig-vnet-01",
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
      "name": "my06-mig-sg-02",
      "connectionName": "ibm-au-syd",
      "vNetId": "my06-mig-vnet-01",
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
        "nodeGroupId": "my06-ng-influxdb-back"
      },
      "healthChecker": {
        "interval": 10,
        "threshold": 3,
        "timeout": 5
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
- **MCI ID**: `my06-infra101`
- **MCI Name**: `my06-infra101`
- **Status**: `Running:3 (R:3/3)`

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my06-infra101",
  "uid": "tb7905t31ev4a823krab",
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
  "label": {
    "cm-migrated-at": "2026-11-03T07:43:29Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my06-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my06-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tb7905t31ev4a823krab"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbgvp7m5uevfqkgni54c",
      "cspResourceName": "tbgvp7m5uevfqkgni54c",
      "cspResourceId": "02h7_7521e730-981f-4a86-b15c-b99730b8df4a",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-03 07:43:48",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=83.3%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.99.65",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "rootDeviceName": "Not visible in IBM",
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
        "costPerHour": 0.109
      },
      "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-mig-vnet-01",
      "cspVNetId": "r026-677e7ee6-aabe-4622-b1df-343f2a039427",
      "subnetId": "my06-mig-subnet-01",
      "cspSubnetId": "02h7-55b72f6f-ae69-452f-81d3-581350706304",
      "networkInterface": "smithy-untried-caviar-unkind",
      "securityGroupIds": [
        "my06-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-mig-sshkey-01",
      "cspSshKeyId": "r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBLjbFOBnQYvDd9AeLV98Cq8kK/h/bLuJMMO60OvN2todDJeaaUvcSb3919ye177a8ZTqasWxfJzWWWEsvjfCy5A=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2EfhZ3D2m67za5cwaJaotP3F2xjipuTEsZwRvIe7A2I",
        "firstUsedAt": "2026-11-03T07:44:20Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-03T07:43:49Z",
          "completedTime": "2026-11-03T07:44:03Z",
          "elapsedTime": 14,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.4:22 (as \"cb-user\") via bastion Node \"my06-ng-influxdb-back-2\" at 159.23.101.135:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 159.23.101.135:22 as user \"cb-user\" (bastionNodeId=my06-ng-influxdb-back-2): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T07:44:17Z",
          "completedTime": "2026-11-03T07:44:23Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 4,
          "lastOccurredTime": "2026-11-03T07:45:20Z"
        },
        {
          "index": 3,
          "xRequestId": "pc-my06-infra101-tbj8imi7tg9h5suuoskt",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:24Z",
          "completedTime": "2026-11-03T07:45:29Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbgvp7m5uevfqkgni54c 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-ng-influxdb-back-1",
      "uid": "tbqkumucf4rj1qk9e5k8",
      "cspResourceName": "tbqkumucf4rj1qk9e5k8",
      "cspResourceId": "02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-03 07:43:49",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.99.66",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "rootDeviceName": "Not visible in IBM",
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
        "costPerHour": 0.247
      },
      "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-mig-vnet-01",
      "cspVNetId": "r026-677e7ee6-aabe-4622-b1df-343f2a039427",
      "subnetId": "my06-mig-subnet-01",
      "cspSubnetId": "02h7-55b72f6f-ae69-452f-81d3-581350706304",
      "networkInterface": "hummus-imaging-undertone-bleach",
      "securityGroupIds": [
        "my06-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-mig-sshkey-01",
      "cspSshKeyId": "r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOswVLwNLoCdcG7HYbFZSh0Zesdj4QUhSgmtNGjPrKwnexqANZJ3fJsq2CzsEu1map4OvC7l6CKkyqRgYd25E8Y=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:grQuuOQS3b/NQSPQaVI9ge8JDP+CTHTwz0+zuE4Kh78",
        "firstUsedAt": "2026-11-03T07:44:35Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-03T07:43:49Z",
          "completedTime": "2026-11-03T07:44:04Z",
          "elapsedTime": 15,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my06-ng-influxdb-back-1\" at 10.0.1.5:22 (as \"cb-user\") via bastion Node \"my06-ng-influxdb-back-2\" at 159.23.101.135:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 159.23.101.135:22 as user \"cb-user\" (bastionNodeId=my06-ng-influxdb-back-2): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-03T07:44:17Z",
          "completedTime": "2026-11-03T07:44:22Z",
          "elapsedTime": 5,
          "resultSummary": "Command execution failed",
          "errorMessage": "SSH authentication failed. Please check: 1) private key is valid 2) user 'cb-user' exists on target 3) authorized_keys is properly configured",
          "repeatCount": 3,
          "lastOccurredTime": "2026-11-03T07:45:00Z"
        },
        {
          "index": 3,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:15Z",
          "completedTime": "2026-11-03T07:45:20Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 4,
          "xRequestId": "pc-my06-infra101-tbj8imi7tg9h5suuoskt",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:24Z",
          "completedTime": "2026-11-03T07:45:29Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbqkumucf4rj1qk9e5k8 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-ng-influxdb-back-2",
      "uid": "tb8oblp2ov349k4l2amn",
      "cspResourceName": "tb8oblp2ov349k4l2amn",
      "cspResourceId": "02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-03 07:43:48",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.101.135",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "rootDeviceName": "Not visible in IBM",
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
        "costPerHour": 0.247
      },
      "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-mig-vnet-01",
      "cspVNetId": "r026-677e7ee6-aabe-4622-b1df-343f2a039427",
      "subnetId": "my06-mig-subnet-01",
      "cspSubnetId": "02h7-55b72f6f-ae69-452f-81d3-581350706304",
      "networkInterface": "audience-rebuild-simmering-swinger",
      "securityGroupIds": [
        "my06-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-mig-sshkey-01",
      "cspSshKeyId": "r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGCf8zuYKtkcHiveVGONpuL4Hcxz0qZjk5vxlB5nUYYVZsgrffB1pvZA42COIpKQMGHszfnRbLuvwyC8ix8M7qc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:AMVe7TBr7d1i8edEM5LDXoAbIABeVQUXcZFFnes8Dt8",
        "firstUsedAt": "2026-11-03T07:44:00Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T07:43:49Z",
          "completedTime": "2026-11-03T07:44:07Z",
          "elapsedTime": 9,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 5,
          "lastOccurredTime": "2026-11-03T07:45:24Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my06-infra101-tbj8imi7tg9h5suuoskt",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:24Z",
          "completedTime": "2026-11-03T07:45:33Z",
          "elapsedTime": 9,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb8oblp2ov349k4l2amn 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my06-mig-vnet-01",
      "name": "my06-mig-vnet-01",
      "infraId": "my06-infra101",
      "vNetId": "my06-mig-vnet-01",
      "connectionNames": [
        "ibm-au-syd"
      ],
      "providerNames": [
        "ibm"
      ],
      "regionNames": [
        "au-syd"
      ],
      "nodeGroupIds": [
        "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my06-ng-influxdb-back"
      ],
      "nodeIds": [
        "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my06-ng-influxdb-back-1",
        "my06-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my06-infra101",
            "nodeId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "159.23.99.65",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbgvp7m5uevfqkgni54c 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my06-infra101",
            "nodeId": "my06-ng-influxdb-back-1",
            "nodeIp": "159.23.99.66",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbqkumucf4rj1qk9e5k8 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my06-infra101",
            "nodeId": "my06-ng-influxdb-back-2",
            "nodeIp": "159.23.101.135",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb8oblp2ov349k4l2amn 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my06-infra101-tbj8imi7tg9h5suuoskt"
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
      "uid": "tbblq27aodhadn6mkv7m",
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
        "cm-migrated-at": "2026-11-03T07:41:38Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my01-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my01-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbblq27aodhadn6mkv7m"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb58m0p4olu7cv9gg9eu",
          "cspResourceName": "tb58m0p4olu7cv9gg9eu",
          "cspResourceId": "i-0be90f3aabdd99191",
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
          "createdTime": "2026-11-03 07:42:01",
          "label": {
            "Name": "tb58m0p4olu7cv9gg9eu",
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-11-03 07:42:01",
            "sys.cspResourceId": "i-0be90f3aabdd99191",
            "sys.cspResourceName": "tb58m0p4olu7cv9gg9eu",
            "sys.id": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tb58m0p4olu7cv9gg9eu",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "43.203.237.20",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.74",
          "privateDNS": "ip-10-0-1-74.ap-northeast-2.compute.internal",
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
          "cspVNetId": "vpc-0530bd0f6f0f28bf1",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-0f0d9639c2fdd86cf",
          "networkInterface": "eni-0bbbcd9dd237ec0c0",
          "securityGroupIds": [
            "my01-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tbisjldj8j6botepff4r"
        },
        {
          "resourceType": "node",
          "id": "my01-ng-influxdb-back-1",
          "uid": "tbnrapi25m57aeninult",
          "cspResourceName": "tbnrapi25m57aeninult",
          "cspResourceId": "i-017eb01b702b8df89",
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
          "createdTime": "2026-11-03 07:42:01",
          "label": {
            "Name": "tbnrapi25m57aeninult",
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-11-03 07:42:01",
            "sys.cspResourceId": "i-017eb01b702b8df89",
            "sys.cspResourceName": "tbnrapi25m57aeninult",
            "sys.id": "my01-ng-influxdb-back-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-influxdb-back",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tbnrapi25m57aeninult",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "3.34.142.234",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.24",
          "privateDNS": "ip-10-0-1-24.ap-northeast-2.compute.internal",
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
          "cspVNetId": "vpc-0530bd0f6f0f28bf1",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-0f0d9639c2fdd86cf",
          "networkInterface": "eni-01b30ec8915ee51bd",
          "securityGroupIds": [
            "my01-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tbisjldj8j6botepff4r"
        },
        {
          "resourceType": "node",
          "id": "my01-ng-influxdb-back-2",
          "uid": "tbjhsqdr0iiqpf7ndh7o",
          "cspResourceName": "tbjhsqdr0iiqpf7ndh7o",
          "cspResourceId": "i-0eca101eb0ca896bc",
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
          "createdTime": "2026-11-03 07:41:58",
          "label": {
            "Name": "tbjhsqdr0iiqpf7ndh7o",
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-11-03 07:41:58",
            "sys.cspResourceId": "i-0eca101eb0ca896bc",
            "sys.cspResourceName": "tbjhsqdr0iiqpf7ndh7o",
            "sys.id": "my01-ng-influxdb-back-2",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-ng-influxdb-back",
            "sys.subnetId": "my01-mig-subnet-01",
            "sys.uid": "tbjhsqdr0iiqpf7ndh7o",
            "sys.vNetId": "my01-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "52.78.237.145",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.65",
          "privateDNS": "ip-10-0-1-65.ap-northeast-2.compute.internal",
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
          "cspVNetId": "vpc-0530bd0f6f0f28bf1",
          "subnetId": "my01-mig-subnet-01",
          "cspSubnetId": "subnet-0f0d9639c2fdd86cf",
          "networkInterface": "eni-0bf539f94503c3534",
          "securityGroupIds": [
            "my01-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-mig-sshkey-01",
          "cspSshKeyId": "tbisjldj8j6botepff4r"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my02-infra101",
      "uid": "tbf9brepuratknbi2e7p",
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
        "cm-migrated-at": "2026-11-03T07:42:01Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my02-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbf9brepuratknbi2e7p"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbeemen420e4opep2k80",
          "cspResourceName": "tbeemen420e4opep2k80",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbeemen420e4opep2k80",
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
          "createdTime": "2026-11-03 07:43:01",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "createdBy": "tbeemen420e4opep2k80",
            "keypair": "tbs8fotlhr5f5oae19tk",
            "publicip": "tbeemen420e4opep2k80-75388-PublicIP",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-11-03 07:43:01",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbeemen420e4opep2k80",
            "sys.cspResourceName": "tbeemen420e4opep2k80",
            "sys.id": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tbeemen420e4opep2k80",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.173.91",
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbkn8m7ljnebl78733mk",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbkn8m7ljnebl78733mk/subnets/tbs6n1eoeupnosinejsu",
          "networkInterface": "tbeemen420e4opep2k80-50437-VNic",
          "securityGroupIds": [
            "my02-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbs8fotlhr5f5oae19tk"
        },
        {
          "resourceType": "node",
          "id": "my02-ng-influxdb-back-1",
          "uid": "tbf1c7vc17ipondln9oo",
          "cspResourceName": "tbf1c7vc17ipondln9oo",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbf1c7vc17ipondln9oo",
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
          "createdTime": "2026-11-03 07:43:01",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "createdBy": "tbf1c7vc17ipondln9oo",
            "keypair": "tbs8fotlhr5f5oae19tk",
            "nlbBackend": "influxdb_back",
            "publicip": "tbf1c7vc17ipondln9oo-7140-PublicIP",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-11-03 07:43:01",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbf1c7vc17ipondln9oo",
            "sys.cspResourceName": "tbf1c7vc17ipondln9oo",
            "sys.id": "my02-ng-influxdb-back-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-influxdb-back",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tbf1c7vc17ipondln9oo",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.168.210",
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbkn8m7ljnebl78733mk",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbkn8m7ljnebl78733mk/subnets/tbs6n1eoeupnosinejsu",
          "networkInterface": "tbf1c7vc17ipondln9oo-2865-VNic",
          "securityGroupIds": [
            "my02-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbs8fotlhr5f5oae19tk"
        },
        {
          "resourceType": "node",
          "id": "my02-ng-influxdb-back-2",
          "uid": "tbqtipogut7qkrsde66e",
          "cspResourceName": "tbqtipogut7qkrsde66e",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqtipogut7qkrsde66e",
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
          "createdTime": "2026-11-03 07:43:01",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "createdBy": "tbqtipogut7qkrsde66e",
            "keypair": "tbs8fotlhr5f5oae19tk",
            "nlbBackend": "influxdb_back",
            "publicip": "tbqtipogut7qkrsde66e-72765-PublicIP",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-11-03 07:43:01",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbqtipogut7qkrsde66e",
            "sys.cspResourceName": "tbqtipogut7qkrsde66e",
            "sys.id": "my02-ng-influxdb-back-2",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-ng-influxdb-back",
            "sys.subnetId": "my02-mig-subnet-01",
            "sys.uid": "tbqtipogut7qkrsde66e",
            "sys.vNetId": "my02-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.6.139",
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202610220"
          },
          "vNetId": "my02-mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbkn8m7ljnebl78733mk",
          "subnetId": "my02-mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbkn8m7ljnebl78733mk/subnets/tbs6n1eoeupnosinejsu",
          "networkInterface": "tbqtipogut7qkrsde66e-79241-VNic",
          "securityGroupIds": [
            "my02-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbs8fotlhr5f5oae19tk"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my03-infra101",
      "uid": "tb0lf8kkffh6fdj4lnf8",
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
        "cm-migrated-at": "2026-11-03T07:44:07Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my03-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my03-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tb0lf8kkffh6fdj4lnf8"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbadk2iui8mcqcq39okd",
          "cspResourceName": "tbadk2iui8mcqcq39okd",
          "cspResourceId": "tbadk2iui8mcqcq39okd",
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
          "createdTime": "2026-11-03 07:44:29",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "keypair": "tbsslk45sts0rjvmkip5",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-11-03 07:44:29",
            "sys.cspResourceId": "tbadk2iui8mcqcq39okd",
            "sys.cspResourceName": "tbadk2iui8mcqcq39okd",
            "sys.id": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tbadk2iui8mcqcq39okd",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.64.54.69",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
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
          "specId": "gcp+asia-northeast3+n2-highcpu-2",
          "cspSpecName": "n2-highcpu-2",
          "spec": {
            "cspSpecName": "n2-highcpu-2",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.092072
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21"
          },
          "vNetId": "my03-mig-vnet-01",
          "cspVNetId": "tbr3nelub33c1khtsmap",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tb7mjg75dtuc0r013vhc",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tbsslk45sts0rjvmkip5"
        },
        {
          "resourceType": "node",
          "id": "my03-ng-influxdb-back-1",
          "uid": "tbosjonktf0ju67ngpa3",
          "cspResourceName": "tbosjonktf0ju67ngpa3",
          "cspResourceId": "tbosjonktf0ju67ngpa3",
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
          "createdTime": "2026-11-03 07:44:29",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "keypair": "tbsslk45sts0rjvmkip5",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-11-03 07:44:29",
            "sys.cspResourceId": "tbosjonktf0ju67ngpa3",
            "sys.cspResourceName": "tbosjonktf0ju67ngpa3",
            "sys.id": "my03-ng-influxdb-back-1",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-influxdb-back",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tbosjonktf0ju67ngpa3",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "8.230.27.190",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.3",
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
            "memoryGiB": 16,
            "costPerHour": 0.232954
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21"
          },
          "vNetId": "my03-mig-vnet-01",
          "cspVNetId": "tbr3nelub33c1khtsmap",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tb7mjg75dtuc0r013vhc",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tbsslk45sts0rjvmkip5"
        },
        {
          "resourceType": "node",
          "id": "my03-ng-influxdb-back-2",
          "uid": "tbck4bkd5ptgfbl199pj",
          "cspResourceName": "tbck4bkd5ptgfbl199pj",
          "cspResourceId": "tbck4bkd5ptgfbl199pj",
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
          "createdTime": "2026-11-03 07:44:28",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "keypair": "tbsslk45sts0rjvmkip5",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-11-03 07:44:28",
            "sys.cspResourceId": "tbck4bkd5ptgfbl199pj",
            "sys.cspResourceName": "tbck4bkd5ptgfbl199pj",
            "sys.id": "my03-ng-influxdb-back-2",
            "sys.infraId": "my03-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my03-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my03-ng-influxdb-back",
            "sys.subnetId": "my03-mig-subnet-01",
            "sys.uid": "tbck4bkd5ptgfbl199pj",
            "sys.vNetId": "my03-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.22.67.165",
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
            "memoryGiB": 16,
            "costPerHour": 0.232954
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20261021",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-10-21"
          },
          "vNetId": "my03-mig-vnet-01",
          "cspVNetId": "tbr3nelub33c1khtsmap",
          "subnetId": "my03-mig-subnet-01",
          "cspSubnetId": "tb7mjg75dtuc0r013vhc",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my03-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my03-mig-sshkey-01",
          "cspSshKeyId": "tbsslk45sts0rjvmkip5"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my05-infra101",
      "uid": "tbv6k2784a7j9vj57htr",
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
      "label": null,
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbp5na95t5jqnos3sgte",
          "cspResourceName": "tbp5na95t5jqnos3sgte",
          "cspResourceId": "ins-dv1l9439",
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
          "createdTime": "2026-11-03 07:42:57",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-11-03 07:42:57",
            "sys.cspResourceId": "ins-dv1l9439",
            "sys.cspResourceName": "tbp5na95t5jqnos3sgte",
            "sys.id": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tbp5na95t5jqnos3sgte",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.133.91.35",
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
          "cspVNetId": "vpc-rpd7j9gm",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-5vgo1hrx",
          "networkInterface": "eni-9nj9eopa",
          "securityGroupIds": [
            "my05-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-9u8pzrmj"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-1",
          "uid": "tbkc296sj72nvubhmdnl",
          "cspResourceName": "tbkc296sj72nvubhmdnl",
          "cspResourceId": "ins-891evur7",
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
          "createdTime": "2026-11-03 07:42:50",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-11-03 07:42:50",
            "sys.cspResourceId": "ins-891evur7",
            "sys.cspResourceName": "tbkc296sj72nvubhmdnl",
            "sys.id": "my05-ng-influxdb-back-1",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tbkc296sj72nvubhmdnl",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.155.144.238",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.5",
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
          "cspVNetId": "vpc-rpd7j9gm",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-5vgo1hrx",
          "networkInterface": "eni-1wcalxpi",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-9u8pzrmj"
        },
        {
          "resourceType": "node",
          "id": "my05-ng-influxdb-back-2",
          "uid": "tbtb55pnojspqcpk365b",
          "cspResourceName": "tbtb55pnojspqcpk365b",
          "cspResourceId": "ins-bc9r5ck9",
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
          "createdTime": "2026-11-03 07:43:05",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "tencent-ap-seoul",
            "sys.createdTime": "2026-11-03 07:43:05",
            "sys.cspResourceId": "ins-bc9r5ck9",
            "sys.cspResourceName": "tbtb55pnojspqcpk365b",
            "sys.id": "my05-ng-influxdb-back-2",
            "sys.infraId": "my05-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my05-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my05-ng-influxdb-back",
            "sys.subnetId": "my05-mig-subnet-01",
            "sys.uid": "tbtb55pnojspqcpk365b",
            "sys.vNetId": "my05-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.155.236.183",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.15",
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
          "cspVNetId": "vpc-rpd7j9gm",
          "subnetId": "my05-mig-subnet-01",
          "cspSubnetId": "subnet-5vgo1hrx",
          "networkInterface": "eni-pvkg790i",
          "securityGroupIds": [
            "my05-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-mig-sshkey-01",
          "cspSshKeyId": "skey-9u8pzrmj"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my06-infra101",
      "uid": "tb7905t31ev4a823krab",
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
      "label": {
        "cm-migrated-at": "2026-11-03T07:43:29Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my06-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tb7905t31ev4a823krab"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbgvp7m5uevfqkgni54c",
          "cspResourceName": "tbgvp7m5uevfqkgni54c",
          "cspResourceId": "02h7_7521e730-981f-4a86-b15c-b99730b8df4a",
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
          "createdTime": "2026-11-03 07:43:48",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-11-03 07:43:48",
            "sys.cspResourceId": "02h7_7521e730-981f-4a86-b15c-b99730b8df4a",
            "sys.cspResourceName": "tbgvp7m5uevfqkgni54c",
            "sys.id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbgvp7m5uevfqkgni54c",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.99.65",
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
          "specId": "ibm+au-syd+cx2-2x4",
          "cspSpecName": "cx2-2x4",
          "spec": {
            "cspSpecName": "cx2-2x4",
            "vCPU": 2,
            "memoryGiB": 4,
            "costPerHour": 0.109
          },
          "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-mig-vnet-01",
          "cspVNetId": "r026-677e7ee6-aabe-4622-b1df-343f2a039427",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-55b72f6f-ae69-452f-81d3-581350706304",
          "networkInterface": "smithy-untried-caviar-unkind",
          "securityGroupIds": [
            "my06-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe"
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-1",
          "uid": "tbqkumucf4rj1qk9e5k8",
          "cspResourceName": "tbqkumucf4rj1qk9e5k8",
          "cspResourceId": "02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667",
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
          "createdTime": "2026-11-03 07:43:49",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-11-03 07:43:49",
            "sys.cspResourceId": "02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667",
            "sys.cspResourceName": "tbqkumucf4rj1qk9e5k8",
            "sys.id": "my06-ng-influxdb-back-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-influxdb-back",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tbqkumucf4rj1qk9e5k8",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.99.66",
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
          "specId": "ibm+au-syd+bx2-4x16",
          "cspSpecName": "bx2-4x16",
          "spec": {
            "cspSpecName": "bx2-4x16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.247
          },
          "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-mig-vnet-01",
          "cspVNetId": "r026-677e7ee6-aabe-4622-b1df-343f2a039427",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-55b72f6f-ae69-452f-81d3-581350706304",
          "networkInterface": "hummus-imaging-undertone-bleach",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe"
        },
        {
          "resourceType": "node",
          "id": "my06-ng-influxdb-back-2",
          "uid": "tb8oblp2ov349k4l2amn",
          "cspResourceName": "tb8oblp2ov349k4l2amn",
          "cspResourceId": "02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f",
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
          "createdTime": "2026-11-03 07:43:48",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-11-03 07:43:48",
            "sys.cspResourceId": "02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f",
            "sys.cspResourceName": "tb8oblp2ov349k4l2amn",
            "sys.id": "my06-ng-influxdb-back-2",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-ng-influxdb-back",
            "sys.subnetId": "my06-mig-subnet-01",
            "sys.uid": "tb8oblp2ov349k4l2amn",
            "sys.vNetId": "my06-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.101.135",
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
            "costPerHour": 0.247
          },
          "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-mig-vnet-01",
          "cspVNetId": "r026-677e7ee6-aabe-4622-b1df-343f2a039427",
          "subnetId": "my06-mig-subnet-01",
          "cspSubnetId": "02h7-55b72f6f-ae69-452f-81d3-581350706304",
          "networkInterface": "audience-rebuild-simmering-swinger",
          "securityGroupIds": [
            "my06-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-mig-sshkey-01",
          "cspSshKeyId": "r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe"
        }
      ],
      "newNodeList": null
    },
    {
      "resourceType": "infra",
      "id": "my09-infra101",
      "uid": "tbgortbktp1vcdl67a7c",
      "name": "my09-infra101",
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
        "cm-migrated-at": "2026-11-03T07:43:14Z",
        "sys.description": "NLB-aware recommended infrastructure for cloud migration",
        "sys.id": "my09-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my09-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbgortbktp1vcdl67a7c"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb0tll5bpoeoo1h6o59k",
          "cspResourceName": "tb0tll5bpoeoo1h6o59k",
          "cspResourceId": "538b3c77-3777-4e37-be29-1004fcda8048",
          "name": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-03 07:45:11",
          "label": {
            "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "nhn-kr1",
            "sys.createdTime": "2026-11-03 07:45:11",
            "sys.cspResourceId": "538b3c77-3777-4e37-be29-1004fcda8048",
            "sys.cspResourceName": "tb0tll5bpoeoo1h6o59k",
            "sys.id": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my09-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my09-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my09-mig-subnet-01",
            "sys.uid": "tb0tll5bpoeoo1h6o59k",
            "sys.vNetId": "my09-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "103.218.158.79",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "General_HDD",
          "rootDiskSize": 20,
          "rootDeviceName": "",
          "connectionName": "nhn-kr1",
          "connectionConfig": {
            "configName": "nhn-kr1",
            "providerName": "nhn",
            "driverName": "nhn-driver-v1.0.so",
            "credentialName": "nhn",
            "credentialHolder": "admin",
            "regionZoneInfoName": "nhn-kr1",
            "regionZoneInfo": {
              "assignedRegion": "KR1",
              "assignedZone": "kr-pub-a"
            },
            "regionDetail": {
              "regionId": "KR1",
              "regionName": "kr1",
              "description": "Pangyo (South Korea)",
              "location": {
                "display": "Pangyo (South Korea)",
                "latitude": 37.390889,
                "longitude": 127.096792
              },
              "zones": [
                "kr-pub-a",
                "kr-pub-b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "nhn+kr1+c2.c2m2",
          "cspSpecName": "c2.c2m2",
          "spec": {
            "cspSpecName": "c2.c2m2",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.042179998
          },
          "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "image": {
            "resourceType": "image",
            "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
          },
          "vNetId": "my09-mig-vnet-01",
          "cspVNetId": "5607e2b9-d31d-445b-8235-826821764519",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "e9cf711b-2444-44d6-9963-14630f58b051",
          "networkInterface": "379a9c81-fd5b-4465-88ad-10c17bdc9f7b",
          "securityGroupIds": [
            "my09-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": "tbm9mtmk1plihtn1bjvn"
        },
        {
          "resourceType": "node",
          "id": "my09-ng-influxdb-back-1",
          "uid": "tbktqnuctg3m6jtqn2c2",
          "cspResourceName": "tbktqnuctg3m6jtqn2c2",
          "cspResourceId": "cb04341b-8725-47fd-b5e4-804c24dd8bba",
          "name": "my09-ng-influxdb-back-1",
          "nodeGroupId": "my09-ng-influxdb-back",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-03 07:45:00",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "nhn-kr1",
            "sys.createdTime": "2026-11-03 07:45:00",
            "sys.cspResourceId": "cb04341b-8725-47fd-b5e4-804c24dd8bba",
            "sys.cspResourceName": "tbktqnuctg3m6jtqn2c2",
            "sys.id": "my09-ng-influxdb-back-1",
            "sys.infraId": "my09-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my09-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my09-ng-influxdb-back",
            "sys.subnetId": "my09-mig-subnet-01",
            "sys.uid": "tbktqnuctg3m6jtqn2c2",
            "sys.vNetId": "my09-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "133.186.222.44",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.101",
          "privateDNS": "",
          "rootDiskType": "General_HDD",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "nhn-kr1",
          "connectionConfig": {
            "configName": "nhn-kr1",
            "providerName": "nhn",
            "driverName": "nhn-driver-v1.0.so",
            "credentialName": "nhn",
            "credentialHolder": "admin",
            "regionZoneInfoName": "nhn-kr1",
            "regionZoneInfo": {
              "assignedRegion": "KR1",
              "assignedZone": "kr-pub-a"
            },
            "regionDetail": {
              "regionId": "KR1",
              "regionName": "kr1",
              "description": "Pangyo (South Korea)",
              "location": {
                "display": "Pangyo (South Korea)",
                "latitude": 37.390889,
                "longitude": 127.096792
              },
              "zones": [
                "kr-pub-a",
                "kr-pub-b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "nhn+kr1+r2.c4m16",
          "cspSpecName": "r2.c4m16",
          "spec": {
            "cspSpecName": "r2.c4m16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.19535999
          },
          "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "image": {
            "resourceType": "image",
            "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
          },
          "vNetId": "my09-mig-vnet-01",
          "cspVNetId": "5607e2b9-d31d-445b-8235-826821764519",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "e9cf711b-2444-44d6-9963-14630f58b051",
          "networkInterface": "b480bf36-cb9a-42ce-9b60-410f1a75b017",
          "securityGroupIds": [
            "my09-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": "tbm9mtmk1plihtn1bjvn"
        },
        {
          "resourceType": "node",
          "id": "my09-ng-influxdb-back-2",
          "uid": "tbkkkgtf064hoa3nkg6b",
          "cspResourceName": "tbkkkgtf064hoa3nkg6b",
          "cspResourceId": "f3c5c55a-727c-450f-92c9-e1f6263e3ef7",
          "name": "my09-ng-influxdb-back-2",
          "nodeGroupId": "my09-ng-influxdb-back",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "2026-11-03 07:44:58",
          "label": {
            "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nlbBackend": "influxdb_back",
            "sys.connectionName": "nhn-kr1",
            "sys.createdTime": "2026-11-03 07:44:58",
            "sys.cspResourceId": "f3c5c55a-727c-450f-92c9-e1f6263e3ef7",
            "sys.cspResourceName": "tbkkkgtf064hoa3nkg6b",
            "sys.id": "my09-ng-influxdb-back-2",
            "sys.infraId": "my09-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my09-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my09-ng-influxdb-back",
            "sys.subnetId": "my09-mig-subnet-01",
            "sys.uid": "tbkkkgtf064hoa3nkg6b",
            "sys.vNetId": "my09-mig-vnet-01"
          },
          "description": "",
          "region": {
            "region": "KR1",
            "zone": "kr-pub-a"
          },
          "publicIP": "133.186.213.214",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.100",
          "privateDNS": "",
          "rootDiskType": "General_HDD",
          "rootDiskSize": 30,
          "rootDeviceName": "",
          "connectionName": "nhn-kr1",
          "connectionConfig": {
            "configName": "nhn-kr1",
            "providerName": "nhn",
            "driverName": "nhn-driver-v1.0.so",
            "credentialName": "nhn",
            "credentialHolder": "admin",
            "regionZoneInfoName": "nhn-kr1",
            "regionZoneInfo": {
              "assignedRegion": "KR1",
              "assignedZone": "kr-pub-a"
            },
            "regionDetail": {
              "regionId": "KR1",
              "regionName": "kr1",
              "description": "Pangyo (South Korea)",
              "location": {
                "display": "Pangyo (South Korea)",
                "latitude": 37.390889,
                "longitude": 127.096792
              },
              "zones": [
                "kr-pub-a",
                "kr-pub-b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "nhn+kr1+r2.c4m16",
          "cspSpecName": "r2.c4m16",
          "spec": {
            "cspSpecName": "r2.c4m16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.19535999
          },
          "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "image": {
            "resourceType": "image",
            "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
          },
          "vNetId": "my09-mig-vnet-01",
          "cspVNetId": "5607e2b9-d31d-445b-8235-826821764519",
          "subnetId": "my09-mig-subnet-01",
          "cspSubnetId": "e9cf711b-2444-44d6-9963-14630f58b051",
          "networkInterface": "c66a388a-2338-4c34-a3e8-b84e8ff4c70a",
          "securityGroupIds": [
            "my09-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my09-mig-sshkey-01",
          "cspSshKeyId": "tbm9mtmk1plihtn1bjvn"
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
- **IDs**: [my01-infra101 my02-infra101 my03-infra101 my05-infra101 my06-infra101 my09-infra101]

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
  "id": "my06-infra101",
  "uid": "tb7905t31ev4a823krab",
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
  "label": {
    "cm-migrated-at": "2026-11-03T07:43:29Z",
    "sys.description": "NLB-aware recommended infrastructure for cloud migration",
    "sys.id": "my06-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my06-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tb7905t31ev4a823krab"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbgvp7m5uevfqkgni54c",
      "cspResourceName": "tbgvp7m5uevfqkgni54c",
      "cspResourceId": "02h7_7521e730-981f-4a86-b15c-b99730b8df4a",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-03 07:43:48",
      "label": {
        "cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=83.3%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.99.65",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "rootDeviceName": "Not visible in IBM",
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
        "costPerHour": 0.109
      },
      "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-mig-vnet-01",
      "cspVNetId": "r026-677e7ee6-aabe-4622-b1df-343f2a039427",
      "subnetId": "my06-mig-subnet-01",
      "cspSubnetId": "02h7-55b72f6f-ae69-452f-81d3-581350706304",
      "networkInterface": "smithy-untried-caviar-unkind",
      "securityGroupIds": [
        "my06-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-mig-sshkey-01",
      "cspSshKeyId": "r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBLjbFOBnQYvDd9AeLV98Cq8kK/h/bLuJMMO60OvN2todDJeaaUvcSb3919ye177a8ZTqasWxfJzWWWEsvjfCy5A=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:2EfhZ3D2m67za5cwaJaotP3F2xjipuTEsZwRvIe7A2I",
        "firstUsedAt": "2026-11-03T07:44:20Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-03T07:43:49Z",
          "completedTime": "2026-11-03T07:44:03Z",
          "elapsedTime": 14,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.4:22 (as \"cb-user\") via bastion Node \"my06-ng-influxdb-back-2\" at 159.23.101.135:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 159.23.101.135:22 as user \"cb-user\" (bastionNodeId=my06-ng-influxdb-back-2): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T07:44:17Z",
          "completedTime": "2026-11-03T07:44:23Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 4,
          "lastOccurredTime": "2026-11-03T07:45:20Z"
        },
        {
          "index": 3,
          "xRequestId": "pc-my06-infra101-tbj8imi7tg9h5suuoskt",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:24Z",
          "completedTime": "2026-11-03T07:45:29Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbgvp7m5uevfqkgni54c 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        },
        {
          "index": 4,
          "xRequestId": "1793691938764030488",
          "commandRequested": "echo ready",
          "commandExecuted": "echo ready",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:38Z",
          "completedTime": "2026-11-03T07:45:45Z",
          "elapsedTime": 7,
          "resultSummary": "Command executed successfully",
          "stdout": "ready\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-ng-influxdb-back-1",
      "uid": "tbqkumucf4rj1qk9e5k8",
      "cspResourceName": "tbqkumucf4rj1qk9e5k8",
      "cspResourceId": "02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-03 07:43:49",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.99.66",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "rootDeviceName": "Not visible in IBM",
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
        "costPerHour": 0.247
      },
      "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-mig-vnet-01",
      "cspVNetId": "r026-677e7ee6-aabe-4622-b1df-343f2a039427",
      "subnetId": "my06-mig-subnet-01",
      "cspSubnetId": "02h7-55b72f6f-ae69-452f-81d3-581350706304",
      "networkInterface": "hummus-imaging-undertone-bleach",
      "securityGroupIds": [
        "my06-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-mig-sshkey-01",
      "cspSshKeyId": "r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOswVLwNLoCdcG7HYbFZSh0Zesdj4QUhSgmtNGjPrKwnexqANZJ3fJsq2CzsEu1map4OvC7l6CKkyqRgYd25E8Y=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:grQuuOQS3b/NQSPQaVI9ge8JDP+CTHTwz0+zuE4Kh78",
        "firstUsedAt": "2026-11-03T07:44:35Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-03T07:43:49Z",
          "completedTime": "2026-11-03T07:44:04Z",
          "elapsedTime": 15,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my06-ng-influxdb-back-1\" at 10.0.1.5:22 (as \"cb-user\") via bastion Node \"my06-ng-influxdb-back-2\" at 159.23.101.135:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 159.23.101.135:22 as user \"cb-user\" (bastionNodeId=my06-ng-influxdb-back-2): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Failed",
          "startedTime": "2026-11-03T07:44:17Z",
          "completedTime": "2026-11-03T07:44:22Z",
          "elapsedTime": 5,
          "resultSummary": "Command execution failed",
          "errorMessage": "SSH authentication failed. Please check: 1) private key is valid 2) user 'cb-user' exists on target 3) authorized_keys is properly configured",
          "repeatCount": 3,
          "lastOccurredTime": "2026-11-03T07:45:00Z"
        },
        {
          "index": 3,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:15Z",
          "completedTime": "2026-11-03T07:45:20Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 4,
          "xRequestId": "pc-my06-infra101-tbj8imi7tg9h5suuoskt",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:24Z",
          "completedTime": "2026-11-03T07:45:29Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbqkumucf4rj1qk9e5k8 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        },
        {
          "index": 5,
          "xRequestId": "1793691945144834437",
          "commandRequested": "echo ready",
          "commandExecuted": "echo ready",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:45Z",
          "completedTime": "2026-11-03T07:45:51Z",
          "elapsedTime": 6,
          "resultSummary": "Command executed successfully",
          "stdout": "ready\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-ng-influxdb-back-2",
      "uid": "tb8oblp2ov349k4l2amn",
      "cspResourceName": "tb8oblp2ov349k4l2amn",
      "cspResourceId": "02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f",
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
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-11-03 07:43:48",
      "label": {
        "cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nlbBackend": "influxdb_back"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=83.3%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.101.135",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "rootDeviceName": "Not visible in IBM",
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
        "costPerHour": 0.247
      },
      "imageId": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-6245bce7-1353-421f-9eb8-47aa0ed18c50",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-mig-vnet-01",
      "cspVNetId": "r026-677e7ee6-aabe-4622-b1df-343f2a039427",
      "subnetId": "my06-mig-subnet-01",
      "cspSubnetId": "02h7-55b72f6f-ae69-452f-81d3-581350706304",
      "networkInterface": "audience-rebuild-simmering-swinger",
      "securityGroupIds": [
        "my06-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-mig-sshkey-01",
      "cspSshKeyId": "r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGCf8zuYKtkcHiveVGONpuL4Hcxz0qZjk5vxlB5nUYYVZsgrffB1pvZA42COIpKQMGHszfnRbLuvwyC8ix8M7qc=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:AMVe7TBr7d1i8edEM5LDXoAbIABeVQUXcZFFnes8Dt8",
        "firstUsedAt": "2026-11-03T07:44:00Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-11-03T07:43:49Z",
          "completedTime": "2026-11-03T07:44:07Z",
          "elapsedTime": 9,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n",
          "repeatCount": 5,
          "lastOccurredTime": "2026-11-03T07:45:24Z"
        },
        {
          "index": 2,
          "xRequestId": "pc-my06-infra101-tbj8imi7tg9h5suuoskt",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:24Z",
          "completedTime": "2026-11-03T07:45:33Z",
          "elapsedTime": 9,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb8oblp2ov349k4l2amn 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        },
        {
          "index": 3,
          "xRequestId": "1793691951271332390",
          "commandRequested": "echo ready",
          "commandExecuted": "echo ready",
          "status": "Completed",
          "startedTime": "2026-11-03T07:45:51Z",
          "completedTime": "2026-11-03T07:45:54Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "ready\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my06-mig-vnet-01",
      "name": "my06-mig-vnet-01",
      "infraId": "my06-infra101",
      "vNetId": "my06-mig-vnet-01",
      "connectionNames": [
        "ibm-au-syd"
      ],
      "providerNames": [
        "ibm"
      ],
      "regionNames": [
        "au-syd"
      ],
      "nodeGroupIds": [
        "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "my06-ng-influxdb-back"
      ],
      "nodeIds": [
        "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my06-ng-influxdb-back-1",
        "my06-ng-influxdb-back-2"
      ],
      "nodeGroupCount": 2,
      "nodeCount": 3,
      "representativeNodeGroupId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my06-infra101",
            "nodeId": "my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "159.23.99.65",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbgvp7m5uevfqkgni54c 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my06-infra101",
            "nodeId": "my06-ng-influxdb-back-1",
            "nodeIp": "159.23.99.66",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbqkumucf4rj1qk9e5k8 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my06-infra101",
            "nodeId": "my06-ng-influxdb-back-2",
            "nodeIp": "159.23.101.135",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb8oblp2ov349k4l2amn 5.15.0-1111-ibm #115-Ubuntu SMP Tue Sep 22 16:54:27 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my06-infra101-tbj8imi7tg9h5suuoskt"
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
      "id": "my06-ng-influxdb-back",
      "cspResourceName": "tb5fm18tiu7apkud8814",
      "cspResourceId": "r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39",
      "name": "my06-ng-influxdb-back",
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
      "type": "PUBLIC",
      "scope": "REGION",
      "listener": {
        "protocol": "TCP",
        "ip": "159.23.101.153",
        "port": "9999",
        "dnsName": "3ea8b9b5-au-syd.lb.appdomain.cloud",
        "keyValueList": [
          {
            "key": "Protocol",
            "value": "TCP"
          },
          {
            "key": "IP",
            "value": "159.23.101.153"
          },
          {
            "key": "Port",
            "value": "9999"
          },
          {
            "key": "DNSName",
            "value": "3ea8b9b5-au-syd.lb.appdomain.cloud"
          },
          {
            "key": "CspID",
            "value": "r026-d937a0b0-8b29-4afa-86f0-224ddb68e880"
          }
        ]
      },
      "targetGroup": {
        "protocol": "TCP",
        "port": "8086",
        "nodeGroupId": "my06-ng-influxdb-back",
        "nodes": [
          "my06-ng-influxdb-back-1",
          "my06-ng-influxdb-back-2"
        ],
        "keyValueList": [
          {
            "key": "Protocol",
            "value": "TCP"
          },
          {
            "key": "Port",
            "value": "8086"
          },
          {
            "key": "VMs",
            "value": "[{NameId:tbqkumucf4rj1qk9e5k8,SystemId:02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667},{NameId:tb8oblp2ov349k4l2amn,SystemId:02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f}]"
          },
          {
            "key": "CspID",
            "value": "r026-dc236939-f84d-479d-a02c-2ab1d6284787"
          }
        ]
      },
      "healthChecker": {
        "protocol": "TCP",
        "port": "8086",
        "interval": 10,
        "threshold": 3,
        "timeout": 5,
        "keyValueList": [
          {
            "key": "Protocol",
            "value": "TCP"
          },
          {
            "key": "Port",
            "value": "8086"
          },
          {
            "key": "Interval",
            "value": "10"
          },
          {
            "key": "Timeout",
            "value": "5"
          },
          {
            "key": "Threshold",
            "value": "3"
          }
        ]
      },
      "createdTime": "2026-11-03T07:46:53Z",
      "description": "Migrated from HAProxy backend: influxdb_back",
      "status": "",
      "keyValueList": [
        {
          "key": "AccessMode",
          "value": "public"
        },
        {
          "key": "Availability",
          "value": "subnet"
        },
        {
          "key": "CreatedAt",
          "value": "2026-11-03T07:46:53.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::load-balancer:r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39"
        },
        {
          "key": "FailsafePolicyActions",
          "value": "drop; forward"
        },
        {
          "key": "Hostname",
          "value": "3ea8b9b5-au-syd.lb.appdomain.cloud"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39"
        },
        {
          "key": "ID",
          "value": "r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39"
        },
        {
          "key": "InstanceGroupsSupported",
          "value": "true"
        },
        {
          "key": "IsPrivatePath",
          "value": "false"
        },
        {
          "key": "IsPublic",
          "value": "true"
        },
        {
          "key": "Listeners",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39/listeners/r026-d937a0b0-8b29-4afa-86f0-224ddb68e880,id:r026-d937a0b0-8b29-4afa-86f0-224ddb68e880}"
        },
        {
          "key": "Logging",
          "value": "{datapath:{active:false}}"
        },
        {
          "key": "Name",
          "value": "tb5fm18tiu7apkud8814"
        },
        {
          "key": "OperatingStatus",
          "value": "online"
        },
        {
          "key": "Pools",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39/pools/r026-dc236939-f84d-479d-a02c-2ab1d6284787,id:r026-dc236939-f84d-479d-a02c-2ab1d6284787,name:backend-8086-18657}"
        },
        {
          "key": "PrivateIps",
          "value": "{address:10.0.1.7,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-55b72f6f-ae69-452f-81d3-581350706304/reserved_ips/02h7-492ba809-6a43-4d5b-87b1-6388b1950da9,id:02h7-492ba809-6a43-4d5b-87b1-6388b1950da9,name:deferred-snap-viewing-calorie,resource_type:subnet_reserved_ip}; {address:10.0.1.8,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-55b72f6f-ae69-452f-81d3-581350706304/reserved_ips/02h7-e62e63ba-2d87-4263-8211-a6a80204fa3d,id:02h7-e62e63ba-2d87-4263-8211-a6a80204fa3d,name:choice-hurry-putdown-tamper,resource_type:subnet_reserved_ip}"
        },
        {
          "key": "Profile",
          "value": "{family:network,href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancer/profiles/network-fixed,name:network-fixed}"
        },
        {
          "key": "ProvisioningStatus",
          "value": "active"
        },
        {
          "key": "PublicIps",
          "value": "{address:159.23.101.153}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "load_balancer"
        },
        {
          "key": "RouteMode",
          "value": "false"
        },
        {
          "key": "SecurityGroups",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r026-3cc96667-cbb2-4438-bfdc-d7588f69c382,href:https://au-syd.iaas.cloud.ibm.com/v1/security_groups/r026-3cc96667-cbb2-4438-bfdc-d7588f69c382,id:r026-3cc96667-cbb2-4438-bfdc-d7588f69c382,name:sg-tb5fm18tiu7apkud8814}"
        },
        {
          "key": "SecurityGroupsSupported",
          "value": "true"
        },
        {
          "key": "SourceIPSessionPersistenceSupported",
          "value": "true"
        },
        {
          "key": "Subnets",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-55b72f6f-ae69-452f-81d3-581350706304,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-55b72f6f-ae69-452f-81d3-581350706304,id:02h7-55b72f6f-ae69-452f-81d3-581350706304,name:tbg698apfsnhube3soa7,resource_type:subnet}"
        },
        {
          "key": "UDPSupported",
          "value": "true"
        }
      ],
      "isAutoGenerated": false,
      "location": {
        "display": "Australia (Sydney)",
        "latitude": -33.86882,
        "longitude": 151.209296
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
    "id": "my06-ng-influxdb-back",
    "cspResourceName": "tb5fm18tiu7apkud8814",
    "cspResourceId": "r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39",
    "name": "my06-ng-influxdb-back",
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
    "type": "PUBLIC",
    "scope": "REGION",
    "listener": {
      "protocol": "TCP",
      "ip": "159.23.101.153",
      "port": "9999",
      "dnsName": "3ea8b9b5-au-syd.lb.appdomain.cloud",
      "keyValueList": [
        {
          "key": "Protocol",
          "value": "TCP"
        },
        {
          "key": "IP",
          "value": "159.23.101.153"
        },
        {
          "key": "Port",
          "value": "9999"
        },
        {
          "key": "DNSName",
          "value": "3ea8b9b5-au-syd.lb.appdomain.cloud"
        },
        {
          "key": "CspID",
          "value": "r026-d937a0b0-8b29-4afa-86f0-224ddb68e880"
        }
      ]
    },
    "targetGroup": {
      "protocol": "TCP",
      "port": "8086",
      "nodeGroupId": "my06-ng-influxdb-back",
      "nodes": [
        "my06-ng-influxdb-back-1",
        "my06-ng-influxdb-back-2"
      ],
      "keyValueList": [
        {
          "key": "Protocol",
          "value": "TCP"
        },
        {
          "key": "Port",
          "value": "8086"
        },
        {
          "key": "VMs",
          "value": "[{NameId:tbqkumucf4rj1qk9e5k8,SystemId:02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667},{NameId:tb8oblp2ov349k4l2amn,SystemId:02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f}]"
        },
        {
          "key": "CspID",
          "value": "r026-dc236939-f84d-479d-a02c-2ab1d6284787"
        }
      ]
    },
    "healthChecker": {
      "protocol": "TCP",
      "port": "8086",
      "interval": 10,
      "threshold": 3,
      "timeout": 5,
      "keyValueList": [
        {
          "key": "Protocol",
          "value": "TCP"
        },
        {
          "key": "Port",
          "value": "8086"
        },
        {
          "key": "Interval",
          "value": "10"
        },
        {
          "key": "Timeout",
          "value": "5"
        },
        {
          "key": "Threshold",
          "value": "3"
        }
      ]
    },
    "createdTime": "2026-11-03T07:46:53Z",
    "description": "Migrated from HAProxy backend: influxdb_back",
    "status": "",
    "keyValueList": [
      {
        "key": "AccessMode",
        "value": "public"
      },
      {
        "key": "Availability",
        "value": "subnet"
      },
      {
        "key": "CreatedAt",
        "value": "2026-11-03T07:46:53.000Z"
      },
      {
        "key": "CRN",
        "value": "crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::load-balancer:r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39"
      },
      {
        "key": "FailsafePolicyActions",
        "value": "drop; forward"
      },
      {
        "key": "Hostname",
        "value": "3ea8b9b5-au-syd.lb.appdomain.cloud"
      },
      {
        "key": "Href",
        "value": "https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39"
      },
      {
        "key": "ID",
        "value": "r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39"
      },
      {
        "key": "InstanceGroupsSupported",
        "value": "true"
      },
      {
        "key": "IsPrivatePath",
        "value": "false"
      },
      {
        "key": "IsPublic",
        "value": "true"
      },
      {
        "key": "Listeners",
        "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39/listeners/r026-d937a0b0-8b29-4afa-86f0-224ddb68e880,id:r026-d937a0b0-8b29-4afa-86f0-224ddb68e880}"
      },
      {
        "key": "Logging",
        "value": "{datapath:{active:false}}"
      },
      {
        "key": "Name",
        "value": "tb5fm18tiu7apkud8814"
      },
      {
        "key": "OperatingStatus",
        "value": "online"
      },
      {
        "key": "Pools",
        "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39/pools/r026-dc236939-f84d-479d-a02c-2ab1d6284787,id:r026-dc236939-f84d-479d-a02c-2ab1d6284787,name:backend-8086-18657}"
      },
      {
        "key": "PrivateIps",
        "value": "{address:10.0.1.7,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-55b72f6f-ae69-452f-81d3-581350706304/reserved_ips/02h7-492ba809-6a43-4d5b-87b1-6388b1950da9,id:02h7-492ba809-6a43-4d5b-87b1-6388b1950da9,name:deferred-snap-viewing-calorie,resource_type:subnet_reserved_ip}; {address:10.0.1.8,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-55b72f6f-ae69-452f-81d3-581350706304/reserved_ips/02h7-e62e63ba-2d87-4263-8211-a6a80204fa3d,id:02h7-e62e63ba-2d87-4263-8211-a6a80204fa3d,name:choice-hurry-putdown-tamper,resource_type:subnet_reserved_ip}"
      },
      {
        "key": "Profile",
        "value": "{family:network,href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancer/profiles/network-fixed,name:network-fixed}"
      },
      {
        "key": "ProvisioningStatus",
        "value": "active"
      },
      {
        "key": "PublicIps",
        "value": "{address:159.23.101.153}"
      },
      {
        "key": "ResourceGroup",
        "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
      },
      {
        "key": "ResourceType",
        "value": "load_balancer"
      },
      {
        "key": "RouteMode",
        "value": "false"
      },
      {
        "key": "SecurityGroups",
        "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r026-3cc96667-cbb2-4438-bfdc-d7588f69c382,href:https://au-syd.iaas.cloud.ibm.com/v1/security_groups/r026-3cc96667-cbb2-4438-bfdc-d7588f69c382,id:r026-3cc96667-cbb2-4438-bfdc-d7588f69c382,name:sg-tb5fm18tiu7apkud8814}"
      },
      {
        "key": "SecurityGroupsSupported",
        "value": "true"
      },
      {
        "key": "SourceIPSessionPersistenceSupported",
        "value": "true"
      },
      {
        "key": "Subnets",
        "value": "{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-55b72f6f-ae69-452f-81d3-581350706304,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-55b72f6f-ae69-452f-81d3-581350706304,id:02h7-55b72f6f-ae69-452f-81d3-581350706304,name:tbg698apfsnhube3soa7,resource_type:subnet}"
      },
      {
        "key": "UDPSupported",
        "value": "true"
      }
    ],
    "isAutoGenerated": false,
    "location": {
      "display": "Australia (Sydney)",
      "latitude": -33.86882,
      "longitude": 151.209296
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
  "id": "my06-ng-influxdb-back",
  "cspResourceName": "tb5fm18tiu7apkud8814",
  "cspResourceId": "r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39",
  "name": "my06-ng-influxdb-back",
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
  "type": "PUBLIC",
  "scope": "REGION",
  "listener": {
    "protocol": "TCP",
    "ip": "159.23.101.153",
    "port": "9999",
    "dnsName": "3ea8b9b5-au-syd.lb.appdomain.cloud",
    "keyValueList": [
      {
        "key": "Protocol",
        "value": "TCP"
      },
      {
        "key": "IP",
        "value": "159.23.101.153"
      },
      {
        "key": "Port",
        "value": "9999"
      },
      {
        "key": "DNSName",
        "value": "3ea8b9b5-au-syd.lb.appdomain.cloud"
      },
      {
        "key": "CspID",
        "value": "r026-d937a0b0-8b29-4afa-86f0-224ddb68e880"
      }
    ]
  },
  "targetGroup": {
    "protocol": "TCP",
    "port": "8086",
    "nodeGroupId": "my06-ng-influxdb-back",
    "nodes": [
      "my06-ng-influxdb-back-1",
      "my06-ng-influxdb-back-2"
    ],
    "keyValueList": [
      {
        "key": "Protocol",
        "value": "TCP"
      },
      {
        "key": "Port",
        "value": "8086"
      },
      {
        "key": "VMs",
        "value": "[{NameId:tbqkumucf4rj1qk9e5k8,SystemId:02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667},{NameId:tb8oblp2ov349k4l2amn,SystemId:02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f}]"
      },
      {
        "key": "CspID",
        "value": "r026-dc236939-f84d-479d-a02c-2ab1d6284787"
      }
    ]
  },
  "healthChecker": {
    "protocol": "TCP",
    "port": "8086",
    "interval": 10,
    "threshold": 3,
    "timeout": 5,
    "keyValueList": [
      {
        "key": "Protocol",
        "value": "TCP"
      },
      {
        "key": "Port",
        "value": "8086"
      },
      {
        "key": "Interval",
        "value": "10"
      },
      {
        "key": "Timeout",
        "value": "5"
      },
      {
        "key": "Threshold",
        "value": "3"
      }
    ]
  },
  "createdTime": "2026-11-03T07:46:53Z",
  "description": "Migrated from HAProxy backend: influxdb_back",
  "status": "",
  "keyValueList": [
    {
      "key": "AccessMode",
      "value": "public"
    },
    {
      "key": "Availability",
      "value": "subnet"
    },
    {
      "key": "CreatedAt",
      "value": "2026-11-03T07:46:53.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::load-balancer:r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39"
    },
    {
      "key": "FailsafePolicyActions",
      "value": "drop; forward"
    },
    {
      "key": "Hostname",
      "value": "3ea8b9b5-au-syd.lb.appdomain.cloud"
    },
    {
      "key": "Href",
      "value": "https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39"
    },
    {
      "key": "ID",
      "value": "r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39"
    },
    {
      "key": "InstanceGroupsSupported",
      "value": "true"
    },
    {
      "key": "IsPrivatePath",
      "value": "false"
    },
    {
      "key": "IsPublic",
      "value": "true"
    },
    {
      "key": "Listeners",
      "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39/listeners/r026-d937a0b0-8b29-4afa-86f0-224ddb68e880,id:r026-d937a0b0-8b29-4afa-86f0-224ddb68e880}"
    },
    {
      "key": "Logging",
      "value": "{datapath:{active:false}}"
    },
    {
      "key": "Name",
      "value": "tb5fm18tiu7apkud8814"
    },
    {
      "key": "OperatingStatus",
      "value": "online"
    },
    {
      "key": "Pools",
      "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39/pools/r026-dc236939-f84d-479d-a02c-2ab1d6284787,id:r026-dc236939-f84d-479d-a02c-2ab1d6284787,name:backend-8086-18657}"
    },
    {
      "key": "PrivateIps",
      "value": "{address:10.0.1.7,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-55b72f6f-ae69-452f-81d3-581350706304/reserved_ips/02h7-492ba809-6a43-4d5b-87b1-6388b1950da9,id:02h7-492ba809-6a43-4d5b-87b1-6388b1950da9,name:deferred-snap-viewing-calorie,resource_type:subnet_reserved_ip}; {address:10.0.1.8,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-55b72f6f-ae69-452f-81d3-581350706304/reserved_ips/02h7-e62e63ba-2d87-4263-8211-a6a80204fa3d,id:02h7-e62e63ba-2d87-4263-8211-a6a80204fa3d,name:choice-hurry-putdown-tamper,resource_type:subnet_reserved_ip}"
    },
    {
      "key": "Profile",
      "value": "{family:network,href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancer/profiles/network-fixed,name:network-fixed}"
    },
    {
      "key": "ProvisioningStatus",
      "value": "active"
    },
    {
      "key": "PublicIps",
      "value": "{address:159.23.101.153}"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
    },
    {
      "key": "ResourceType",
      "value": "load_balancer"
    },
    {
      "key": "RouteMode",
      "value": "false"
    },
    {
      "key": "SecurityGroups",
      "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r026-3cc96667-cbb2-4438-bfdc-d7588f69c382,href:https://au-syd.iaas.cloud.ibm.com/v1/security_groups/r026-3cc96667-cbb2-4438-bfdc-d7588f69c382,id:r026-3cc96667-cbb2-4438-bfdc-d7588f69c382,name:sg-tb5fm18tiu7apkud8814}"
    },
    {
      "key": "SecurityGroupsSupported",
      "value": "true"
    },
    {
      "key": "SourceIPSessionPersistenceSupported",
      "value": "true"
    },
    {
      "key": "Subnets",
      "value": "{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-55b72f6f-ae69-452f-81d3-581350706304,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-55b72f6f-ae69-452f-81d3-581350706304,id:02h7-55b72f6f-ae69-452f-81d3-581350706304,name:tbg698apfsnhube3soa7,resource_type:subnet}"
    },
    {
      "key": "UDPSupported",
      "value": "true"
    }
  ],
  "isAutoGenerated": false,
  "location": {
    "display": "Australia (Sydney)",
    "latitude": -33.86882,
    "longitude": 151.209296
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

**Generated At:** 2026-11-03 07:56:41

**Namespace:** mig01

**Infra Name:** my06-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my06-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | IBM |
| **Target Region** | au-syd |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Total NLBs** | 1 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| cx2-2x4 | 2 | 4.0 | - | x86_64 |  | $0.1090 | 1 |
| bx2-4x16 | 4 | 16.0 | - | x86_64 |  | $0.2470 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| r026-6245bce7-1353-421f-9eb8-47aa0ed18c50 | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | 02h7_7521e730-981f-4a86-b15c-b99730b8df4a | Running | 2 vCPU, 4.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-mig-vnet-01<br>**Subnet:** my06-mig-subnet-01<br>**Public IP:** 159.23.99.65<br>**Private IP:** 10.0.1.4<br>**SGs:** my06-mig-sg-02<br>**SSH:** my06-mig-sshkey-01 |
| my06-ng-influxdb-back-1 | 02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-mig-vnet-01<br>**Subnet:** my06-mig-subnet-01<br>**Public IP:** 159.23.99.66<br>**Private IP:** 10.0.1.5<br>**SGs:** my06-mig-sg-01<br>**SSH:** my06-mig-sshkey-01 |
| my06-ng-influxdb-back-2 | 02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-mig-vnet-01<br>**Subnet:** my06-mig-subnet-01<br>**Public IP:** 159.23.101.135<br>**Private IP:** 10.0.1.6<br>**SGs:** my06-mig-sg-01<br>**SSH:** my06-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my06-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my06-mig-vnet-01 |
| **CSP VNet ID** | r026-677e7ee6-aabe-4622-b1df-343f2a039427 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my06-mig-subnet-01 | 02h7-55b72f6f-ae69-452f-81d3-581350706304 | 10.0.1.0/24 | au-syd-1 |

### Network Load Balancers (NLB)

#### NLB: my06-ng-influxdb-back

| Property | Value |
|----------|-------|
| **Name** | my06-ng-influxdb-back |
| **CSP NLB Name** | tb5fm18tiu7apkud8814 |
| **CSP NLB ID** | r026-3ea8b9b5-1c4c-42ae-b756-7c4177394d39 |
| **Type / Scope** | PUBLIC / REGION |
| **Connection** | ibm-au-syd |
| **Listener** | TCP:9999 (DNS: 3ea8b9b5-au-syd.lb.appdomain.cloud, IP: 159.23.101.153) |
| **Target Group** | TCP:8086 (Node Group: my06-ng-influxdb-back, 2 nodes) |
| **Health Checker** | TCP:8086 (Interval: 10s, Timeout: 5s, Threshold: 3) |

**Target Group Nodes:**

| Node Name | CSP VM ID | Private IP | Public IP | Status |
|-----------|-----------|------------|-----------|--------|
| my06-ng-influxdb-back-1 | 02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667 | 10.0.1.5 | 159.23.99.66 | Running |
| my06-ng-influxdb-back-2 | 02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f | 10.0.1.6 | 159.23.101.135 | Running |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my06-mig-sshkey-01 | r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe |  | SHA256:KvZKNP5h8vKILwvbWQc6elnTN+XzEYavsUP7MsvpdvE |

### Security Groups

#### Security Group: my06-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my06-mig-sg-01 |
| **CSP Security Group ID** | r026-ed33697b-fd84-409d-abc3-2a363245935b |
| **VNet** | my06-mig-vnet-01 |
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

#### Security Group: my06-mig-sg-02

| Property | Value |
|----------|-------|
| **Name** | my06-mig-sg-02 |
| **CSP Security Group ID** | r026-08d64d86-ec84-4987-af9a-51fd5293b7cf |
| **VNet** | my06-mig-vnet-01 |
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
| **Per Hour** | $0.6030 |
| **Per Day** | $14.47 |
| **Per Month (30 days)** | $434.16 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| IBM | au-syd | 3 | $0.6030 | $434.16 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | cx2-2x4 | $0.1090 | $78.48 |
| my06-ng-influxdb-back-1 | bx2-4x16 | $0.2470 | $177.84 |
| my06-ng-influxdb-back-2 | bx2-4x16 | $0.2470 | $177.84 |




### Test Case 13: Migration Report

#### 13.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 13.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-11-03 07:56:46*

---

## 📊 Migration Summary

**Target Cloud:** IBM

**Target Region:** au-syd

**Namespace:** mig01 | **Infra ID:** my06-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**NLB Migration Status:** Completed

**Total Load Balancers:** 1

**Migrated Load Balancers:** 1

**💰 Estimated Monthly Cost:** $434.16 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | bx2-4x16, cx2-2x4 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my06-mig-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my06-mig-vnet-01) |
| 6 | **Security Group** | 2 security groups | ✅ Created | Total 10 rules in 2 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |
| 8 | **Network Load Balancer** | 1 | ✅ Created | my06-ng-influxdb-back (TCP:9999 -> TCP:8086) |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** 02h7_7521e730-981f-4a86-b15c-b99730b8df4a<br>**Label(sourceMachineId):** ec268ed7-821e-9d73-e79f-961262161624 | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 |
| 2 | **VM Name:** my06-ng-influxdb-back-1<br>**VM ID:** 02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667<br>**Label(sourceMachineId):** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 |
| 3 | **VM Name:** my06-ng-influxdb-back-2<br>**VM ID:** 02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f<br>**Label(sourceMachineId):** ec288dd0-c6fa-8a49-2f60-bc898311febf | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** cx2-2x4<br>**vCPUs:** 2<br>**Memory:** 4.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **CPUs:** 1<br>**Threads:** 2<br>**Memory:** 2 GB<br>**Root Disk:** 8 GB |
| 2 | my06-ng-influxdb-back-1 | **Spec ID:** bx2-4x16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 16 GB<br>**Root Disk:** 30 GB |
| 3 | my06-ng-influxdb-back-2 | **Spec ID:** bx2-4x16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **CPUs:** 1<br>**Threads:** 4<br>**Memory:** 8 GB<br>**Root Disk:** 30 GB |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** r026-6245bce7-1353-421f-9eb8-47aa0ed18c50<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** ip-10-0-1-30<br>**Machine ID:** ec268ed7-821e-9d73-e79f-961262161624 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 2 | my06-ng-influxdb-back-1 | **Image ID:** r026-6245bce7-1353-421f-9eb8-47aa0ed18c50<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** ip-10-0-1-221<br>**Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932 | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |
| 3 | my06-ng-influxdb-back-2 | **Image ID:** r026-6245bce7-1353-421f-9eb8-47aa0ed18c50<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** ip-10-0-1-138<br>**Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf | **PrettyName:** Ubuntu 22.04.3 LTS<br>**Name:** Ubuntu<br>**Version:** 22.04.3 LTS (Jammy Jellyfish) |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 10 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my06-mig-sg-01

**CSP ID:** r026-ed33697b-fd84-409d-abc3-2a363245935b | **VNet:** my06-mig-vnet-01 | **Rules:** 6

**Assigned VMs:**

- **VM:** my06-ng-influxdb-back-1
  - **Source Server:** **Hostname:** ip-10-0-1-221, **Machine ID:** ec2d32b5-98fb-5a96-7913-d3db1ec18932
- **VM:** my06-ng-influxdb-back-2
  - **Source Server:** **Hostname:** ip-10-0-1-138, **Machine ID:** ec288dd0-c6fa-8a49-2f60-bc898311febf

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 2 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | TCP | 9999 | 0.0.0.0/0 | - | Created by system |
| 5 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |
| 6 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my06-mig-sg-02

**CSP ID:** r026-08d64d86-ec84-4987-af9a-51fd5293b7cf | **VNet:** my06-mig-vnet-01 | **Rules:** 4

**Assigned VMs:**

- **VM:** my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1
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
| 1 | **Name:** my06-mig-vnet-01<br>**ID:** r026-677e7ee6-aabe-4622-b1df-343f2a039427 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my06-mig-subnet-01<br>**ID:** 02h7-55b72f6f-ae69-452f-81d3-581350706304 | 10.0.1.0/24 | my06-mig-vnet-01 |

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

### 1. influxdb_back (haproxy) → my06-ng-influxdb-back

**Configuration Comparison:**

| Property | Source (On-Premise) | Target (Cloud NLB) | Note |
|----------|---------------------|--------------------|------|
| **Load Balancer / Name** | haproxy (Host: ip-10-0-1-30) | my06-ng-influxdb-back (CSP ID: tb5fm18tiu7apkud8814) | Replaced by managed NLB |
| **Type / Scope** | Software LB | PUBLIC / REGION | Cloud-native NLB |
| **Frontend Listener** | *:9999 (TCP) | TCP:9999 | DNS: 3ea8b9b5-au-syd.lb.appdomain.cloud; IP: 159.23.101.153 |
| **Backend / Target Group** | influxdb_back (TCP:8086, roundrobin) | my06-ng-influxdb-back (TCP:8086, 2 nodes) | Target Group attached to NodeGroup |
| **Health Check** | Enabled (Interval: 10s, Timeout: 10s, Threshold: 3) | TCP:8086 (Interval: 10s, Timeout: 5s, Threshold: 3) | Timeout adjusted (10s -> 5s) for CSP constraint |

**Backend Node Mappings:**

| No. | Source Backend Server | Source Endpoint | Target Node | Target CSP VM ID | Target Private IP | Target Public IP | Target Port | Status |
|-----|-----------------------|-----------------|-------------|------------------|-------------------|------------------|-------------|--------|
| 1 | influx1 (ip-10-0-1-221) | 10.0.1.221:8086 | my06-ng-influxdb-back-1 | 02h7_d6c5b46b-a06e-46a2-8518-dda1af54a667 | 10.0.1.5 | 159.23.99.66 | 8086 | Running |
| 2 | influx2 (ip-10-0-1-138) | 10.0.1.138:8086 | my06-ng-influxdb-back-2 | 02h7_7d7e0a38-12f9-4906-b1cc-ea8450b5e38f | 10.0.1.6 | 159.23.101.135 | 8086 | Running |

**Migration Notes:**
- On-premise haproxy on ip-10-0-1-30 (migrated to my06-ng-ec268ed7-821e-9d73-e79f-961262161624-1) is replaced by cloud-managed NLB.
- Health check timeout was adjusted from 10s to 5s to satisfy IBM health check constraints (timeout < interval).

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name | CSP Key ID | Fingerprint | Usage |
|-----|--------------|------------|-------------|-------|
| 1 | my06-mig-sshkey-01 | r026-80f3ea90-fed3-4fed-b8cb-2dffbd4937fe | SHA256:KvZKNP5h8vKILwvbWQc6elnTN+XzEYavsUP7MsvpdvE | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.6030 |
| Daily | $14.47 |
| Monthly | $434.16 |
| Yearly | $5209.92 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | cx2-2x4 | $78.48 | 18.1% |
| ip-10-0-1-221 (migrated) | bx2-4x16 | $177.84 | 41.0% |
| ip-10-0-1-138 (migrated) | bx2-4x16 | $177.84 | 41.0% |

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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my06-infra101)",
  "success": true
}
```

