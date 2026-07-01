# CM-Beetle test results for IBM (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with IBM cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: 3950cc5
- imdl: unknown
- CB-Tumblebug: Unknown (Fallback to Latest)
- CB-Spider: Unknown (Fallback to Latest)
- CB-MapUI: Unknown (Fallback to Latest)
- Target CSP: IBM
- Target Region: au-syd
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: July 1, 2026
- Test Time: 14:31:54 KST
- Test Execution: 2026-07-01 14:31:54 KST

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

## Test result for IBM

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 12.41s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 2m29.33s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 19ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 4ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 21ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 22.279s | Pass |
| 7 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 7m54.088s | Pass |
| 8 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 4ms | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 8ms | Pass |
| 10 | NLB Load Balancing Verification | ✅ **PASS** | 2m15.481s | Pass |
| 11 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.262s | Pass |
| 12 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.25s | Pass |
| 13 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 4m6.043s | Pass |
| 14 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 2m19.166s | Pass |

**Overall Result**: 14/14 tests passed ✅

**Total Duration**: 20m34.677649033s

*Test executed on July 1, 2026 at 14:31:54 KST (2026-07-01 14:31:54 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=80.0% Max=100.0% Avg=93.3% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "nlbBackend": "influxdb_back",
              "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bxf-4x16",
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
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
              "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+nxf-2x2",
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
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
        "postCommand": {
          "userName": "",
          "command": null
        },
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
          "id": "ibm+au-syd+bxf-4x16",
          "uid": "tbjneffpj97nm3tr5hou",
          "cspSpecName": "bxf-4x16",
          "name": "ibm+au-syd+bxf-4x16",
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
          "costPerHour": 0.235,
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
              "value": "{default:standard,type:enum,values:[standard,spot]}"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:16}"
            },
            {
              "key": "Name",
              "value": "bxf-4x16"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:fixed,value:1}"
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
              "value": "{type:fixed,value:200}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:4}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:dependent}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[50,100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        },
        {
          "id": "ibm+au-syd+nxf-2x2",
          "uid": "tb87piitj5ocq6ni7t06",
          "cspSpecName": "nxf-2x2",
          "name": "ibm+au-syd+nxf-2x2",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.094,
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
              "value": "{default:standard,type:enum,values:[standard,spot]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:2000}"
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
              "value": "nano"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x2"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "Name",
              "value": "nxf-2x2"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:fixed,value:1}"
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
              "value": "{type:range,default:500,max:1500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:200}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:dependent}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[10,25,50,100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "ibm",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "regionList": [
            "au-syd"
          ],
          "id": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "uid": "tblhnsks7d8647qm524n",
          "name": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ibm-au-syd",
          "infraType": "",
          "fetchedTime": "2026.06.08 09:13:16 Mon",
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
          "osDiskSizeGB": -1,
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
              "value": "2026-05-19T23:51:58.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "Encryption",
              "value": "none"
            },
            {
              "key": "File",
              "value": "{checksums:{sha256:05ab42b9bb4881c3944fc5452b46a275bb94a6366831b1a874fa708585c4ecb1},size:1}"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "ID",
              "value": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "MinimumProvisionedSize",
              "value": "10"
            },
            {
              "key": "Name",
              "value": "ibm-ubuntu-22-04-5-minimal-amd64-15"
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
      ]
    },
    {
      "status": "partially-matched",
      "description": "Candidate #2 | partially-matched | 1 NLB(s) | Overall Match Rate: Min=50.0% Max=100.0% Avg=85.0% | VMs: 2 total, 0 matched, 2 acceptable | 1 NLB warning(s): NLB backend 'influxdb_back': load-balancing algorithm 'roundrobin' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
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
              "nlbBackend": "influxdb_back",
              "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bx2-4x16",
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
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
              "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=80.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+nxf-2x1",
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
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
        "postCommand": {
          "userName": "",
          "command": null
        },
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
          "uid": "tbedrtkrbc8kfnfbbn30",
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
          "costPerHour": 0.241,
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
          ]
        },
        {
          "id": "ibm+au-syd+nxf-2x1",
          "uid": "tbuf3opfnl35dc45cndv",
          "cspSpecName": "nxf-2x1",
          "name": "ibm+au-syd+nxf-2x1",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1,
          "diskSizeGB": -1,
          "costPerHour": 0.062,
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
              "value": "{default:standard,type:enum,values:[standard,spot]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:2000}"
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
              "value": "nano"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x1"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:1}"
            },
            {
              "key": "Name",
              "value": "nxf-2x1"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:fixed,value:1}"
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
              "value": "{type:range,default:500,max:1500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:200}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:dependent}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[10,25,50,100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "ibm",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "regionList": [
            "au-syd"
          ],
          "id": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "uid": "tblhnsks7d8647qm524n",
          "name": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ibm-au-syd",
          "infraType": "",
          "fetchedTime": "2026.06.08 09:13:16 Mon",
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
          "osDiskSizeGB": -1,
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
              "value": "2026-05-19T23:51:58.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "Encryption",
              "value": "none"
            },
            {
              "key": "File",
              "value": "{checksums:{sha256:05ab42b9bb4881c3944fc5452b46a275bb94a6366831b1a874fa708585c4ecb1},size:1}"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "ID",
              "value": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "MinimumProvisionedSize",
              "value": "10"
            },
            {
              "key": "Name",
              "value": "ibm-ubuntu-22-04-5-minimal-amd64-15"
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
  "uid": "tbphfu6hhvohsnkjlae0",
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
    "sys.uid": "tbphfu6hhvohsnkjlae0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbk5oqcsop91pbhic1u4",
      "cspResourceName": "tbk5oqcsop91pbhic1u4",
      "cspResourceId": "02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87",
      "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-07-01 05:34:15",
      "label": {
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-01 05:34:15",
        "sys.cspResourceId": "02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87",
        "sys.cspResourceName": "tbk5oqcsop91pbhic1u4",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbk5oqcsop91pbhic1u4",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.93.79",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
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
      "specId": "ibm+au-syd+nxf-2x2",
      "cspSpecName": "nxf-2x2",
      "spec": {
        "cspSpecName": "nxf-2x2",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.094
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-7212abef-b5a8-4644-ac0b-d144865621fa",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252",
      "networkInterface": "splashy-sandpit-chatting-identical",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-2088d198-ebc6-40fc-b315-1705d5317a4f",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-07-01T05:34:20Z",
          "completedTime": "2026-07-01T05:34:33Z",
          "elapsedTime": 13,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target host via bastion after 3 attempts: failed to establish SSH connection to bastion host: ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "2000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a-5kbzm},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/volume_attachments/02h7-5b89e52a-26da-4145-bb82-32465b318a2a,id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a,name:hamper-graceful-nimbly-captain,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,id:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,name:calculate-cold-earflap-hurricane,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-01T05:33:35.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87"
        },
        {
          "key": "ID",
          "value": "02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "2"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbk5oqcsop91pbhic1u4"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/network_interfaces/02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,id:02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,name:splashy-sandpit-chatting-identical,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,id:02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,name:entourage-attendant-collard-postcard,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/network_interfaces/02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,id:02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,name:splashy-sandpit-chatting-identical,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,id:02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,name:entourage-attendant-collard-postcard,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x2,name:nxf-2x2,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "1500"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "500"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a-5kbzm},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/volume_attachments/02h7-5b89e52a-26da-4145-bb82-32465b318a2a,id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a,name:hamper-graceful-nimbly-captain,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,id:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,name:calculate-cold-earflap-hurricane,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-7212abef-b5a8-4644-ac0b-d144865621fa,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-7212abef-b5a8-4644-ac0b-d144865621fa,id:r026-7212abef-b5a8-4644-ac0b-d144865621fa,name:tbs8qnakggjn5oera7dh,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tbtddbjda2lllfuho1a9",
      "cspResourceName": "tbtddbjda2lllfuho1a9",
      "cspResourceId": "02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab",
      "name": "my-ng-influxdb-back-1",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-01 05:34:14",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-01 05:34:14",
        "sys.cspResourceId": "02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab",
        "sys.cspResourceName": "tbtddbjda2lllfuho1a9",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbtddbjda2lllfuho1a9",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.102.107",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
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
      "specId": "ibm+au-syd+bxf-4x16",
      "cspSpecName": "bxf-4x16",
      "spec": {
        "cspSpecName": "bxf-4x16",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.235
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-7212abef-b5a8-4644-ac0b-d144865621fa",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252",
      "networkInterface": "nylon-shanty-cultural-motivate",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-2088d198-ebc6-40fc-b315-1705d5317a4f",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-07-01T05:34:20Z",
          "completedTime": "2026-07-01T05:34:33Z",
          "elapsedTime": 13,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target host via bastion after 3 attempts: failed to establish SSH connection to bastion host: ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "8000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030-mzszp},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/volume_attachments/02h7-0de201de-c73d-4894-b0e5-b71215cdd030,id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030,name:failing-vitally-headcount-overreach,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,id:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,name:favored-favorite-pavement-disowned,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-01T05:33:39.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab"
        },
        {
          "key": "ID",
          "value": "02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "16"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbtddbjda2lllfuho1a9"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/network_interfaces/02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,id:02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,name:nylon-shanty-cultural-motivate,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-9398db33-011e-4736-bc8c-b195c7600795,id:02h7-9398db33-011e-4736-bc8c-b195c7600795,name:client-districts-impromptu-plated,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/network_interfaces/02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,id:02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,name:nylon-shanty-cultural-motivate,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-9398db33-011e-4736-bc8c-b195c7600795,id:02h7-9398db33-011e-4736-bc8c-b195c7600795,name:client-districts-impromptu-plated,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16,name:bxf-4x16,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "6000"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "2000"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:4,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030-mzszp},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/volume_attachments/02h7-0de201de-c73d-4894-b0e5-b71215cdd030,id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030,name:failing-vitally-headcount-overreach,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,id:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,name:favored-favorite-pavement-disowned,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-7212abef-b5a8-4644-ac0b-d144865621fa,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-7212abef-b5a8-4644-ac0b-d144865621fa,id:r026-7212abef-b5a8-4644-ac0b-d144865621fa,name:tbs8qnakggjn5oera7dh,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tbt5ckjhndln9g8p23d5",
      "cspResourceName": "tbt5ckjhndln9g8p23d5",
      "cspResourceId": "02h7_ce163b62-c155-4950-802c-ec40da1bbee1",
      "name": "my-ng-influxdb-back-2",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-01 05:34:11",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-01 05:34:11",
        "sys.cspResourceId": "02h7_ce163b62-c155-4950-802c-ec40da1bbee1",
        "sys.cspResourceName": "tbt5ckjhndln9g8p23d5",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbt5ckjhndln9g8p23d5",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.97.167",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
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
      "specId": "ibm+au-syd+bxf-4x16",
      "cspSpecName": "bxf-4x16",
      "spec": {
        "cspSpecName": "bxf-4x16",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.235
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-7212abef-b5a8-4644-ac0b-d144865621fa",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252",
      "networkInterface": "avenue-gallon-ignore-obsessive",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-2088d198-ebc6-40fc-b315-1705d5317a4f",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJOHhJMvNzbWuoNueVZ358kak7qJBzvuvWje+VWWyQ5f9JijOOmFvFLmibjnypVZCmE9YFTQjGXPnQ7Gg1MkHmM=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:aCq7EJqGXuDYwtJ1OceEpVgmTrQ5RDlzebG586o+Sp0",
        "firstUsedAt": "2026-07-01T05:34:21Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-07-01T05:34:20Z",
          "completedTime": "2026-07-01T05:34:33Z",
          "elapsedTime": 13,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target host via bastion after 3 attempts: failed to establish SSH connection to bastion host: ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "8000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-85724725-6251-4f51-8944-c502c2cc4a61-5wtfr},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/volume_attachments/02h7-85724725-6251-4f51-8944-c502c2cc4a61,id:02h7-85724725-6251-4f51-8944-c502c2cc4a61,name:bouquet-ducktail-grime-brilliant,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,id:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,name:composer-doubt-duckling-emptier,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-01T05:33:35.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_ce163b62-c155-4950-802c-ec40da1bbee1"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1"
        },
        {
          "key": "ID",
          "value": "02h7_ce163b62-c155-4950-802c-ec40da1bbee1"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "16"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbt5ckjhndln9g8p23d5"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/network_interfaces/02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,id:02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,name:avenue-gallon-ignore-obsessive,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,id:02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,name:scrubbed-harmonica-getup-germproof,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/network_interfaces/02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,id:02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,name:avenue-gallon-ignore-obsessive,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,id:02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,name:scrubbed-harmonica-getup-germproof,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16,name:bxf-4x16,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "6000"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "2000"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:4,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-85724725-6251-4f51-8944-c502c2cc4a61-5wtfr},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/volume_attachments/02h7-85724725-6251-4f51-8944-c502c2cc4a61,id:02h7-85724725-6251-4f51-8944-c502c2cc4a61,name:bouquet-ducktail-grime-brilliant,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,id:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,name:composer-doubt-duckling-emptier,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-7212abef-b5a8-4644-ac0b-d144865621fa,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-7212abef-b5a8-4644-ac0b-d144865621fa,id:r026-7212abef-b5a8-4644-ac0b-d144865621fa,name:tbs8qnakggjn5oera7dh,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
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
        "nodeIp": "159.23.97.167",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "159.23.102.107",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "159.23.93.79",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
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
      "uid": "tbphfu6hhvohsnkjlae0",
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
        "sys.uid": "tbphfu6hhvohsnkjlae0"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbk5oqcsop91pbhic1u4",
          "cspResourceName": "tbk5oqcsop91pbhic1u4",
          "cspResourceId": "02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87",
          "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-07-01 05:34:15",
          "label": {
            "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-07-01 05:34:15",
            "sys.cspResourceId": "02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87",
            "sys.cspResourceName": "tbk5oqcsop91pbhic1u4",
            "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbk5oqcsop91pbhic1u4",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.93.79",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.5",
          "privateDNS": "",
          "rootDiskType": "general-purpose",
          "rootDiskSize": 100,
          "RootDeviceName": "Not visible in IBM",
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
          "specId": "ibm+au-syd+nxf-2x2",
          "cspSpecName": "nxf-2x2",
          "spec": {
            "cspSpecName": "nxf-2x2",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.094
          },
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "r026-7212abef-b5a8-4644-ac0b-d144865621fa",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252",
          "networkInterface": "splashy-sandpit-chatting-identical",
          "securityGroupIds": [
            "my-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "r026-2088d198-ebc6-40fc-b315-1705d5317a4f",
          "nodeUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-07-01T05:34:20Z",
              "completedTime": "2026-07-01T05:34:33Z",
              "elapsedTime": 13,
              "resultSummary": "Command execution failed",
              "errorMessage": "failed to connect to target host via bastion after 3 attempts: failed to establish SSH connection to bastion host: ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Availability",
              "value": "{class:standard}"
            },
            {
              "key": "AvailabilityPolicy",
              "value": "{host_failure:restart,preemption:stop}"
            },
            {
              "key": "Bandwidth",
              "value": "2000"
            },
            {
              "key": "BootVolumeAttachment",
              "value": "{device:{id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a-5kbzm},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/volume_attachments/02h7-5b89e52a-26da-4145-bb82-32465b318a2a,id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a,name:hamper-graceful-nimbly-captain,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,id:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,name:calculate-cold-earflap-hurricane,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-07-01T05:33:35.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87"
            },
            {
              "key": "EnableSecureBoot",
              "value": "false"
            },
            {
              "key": "HealthState",
              "value": "ok"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87"
            },
            {
              "key": "ID",
              "value": "02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87"
            },
            {
              "key": "Image",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
            },
            {
              "key": "LifecycleState",
              "value": "stable"
            },
            {
              "key": "Memory",
              "value": "2"
            },
            {
              "key": "MetadataService",
              "value": "{enabled:false,protocol:http,response_hop_limit:1}"
            },
            {
              "key": "Name",
              "value": "tbk5oqcsop91pbhic1u4"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/network_interfaces/02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,id:02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,name:splashy-sandpit-chatting-identical,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,id:02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,name:entourage-attendant-collard-postcard,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/network_interfaces/02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,id:02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,name:splashy-sandpit-chatting-identical,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,id:02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,name:entourage-attendant-collard-postcard,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
            },
            {
              "key": "Profile",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x2,name:nxf-2x2,resource_type:instance_profile}"
            },
            {
              "key": "ReservationAffinity",
              "value": "{policy:automatic,pool:[]}"
            },
            {
              "key": "ResourceGroup",
              "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
            },
            {
              "key": "ResourceType",
              "value": "instance"
            },
            {
              "key": "Startable",
              "value": "true"
            },
            {
              "key": "Status",
              "value": "running"
            },
            {
              "key": "TotalNetworkBandwidth",
              "value": "1500"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "500"
            },
            {
              "key": "Vcpu",
              "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
            },
            {
              "key": "VolumeAttachments",
              "value": "{device:{id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a-5kbzm},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/volume_attachments/02h7-5b89e52a-26da-4145-bb82-32465b318a2a,id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a,name:hamper-graceful-nimbly-captain,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,id:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,name:calculate-cold-earflap-hurricane,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-7212abef-b5a8-4644-ac0b-d144865621fa,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-7212abef-b5a8-4644-ac0b-d144865621fa,id:r026-7212abef-b5a8-4644-ac0b-d144865621fa,name:tbs8qnakggjn5oera7dh,resource_type:vpc}"
            },
            {
              "key": "Zone",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-1",
          "uid": "tbtddbjda2lllfuho1a9",
          "cspResourceName": "tbtddbjda2lllfuho1a9",
          "cspResourceId": "02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab",
          "name": "my-ng-influxdb-back-1",
          "nodeGroupId": "my-ng-influxdb-back",
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
          "createdTime": "2026-07-01 05:34:14",
          "label": {
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-07-01 05:34:14",
            "sys.cspResourceId": "02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab",
            "sys.cspResourceName": "tbtddbjda2lllfuho1a9",
            "sys.id": "my-ng-influxdb-back-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbtddbjda2lllfuho1a9",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.102.107",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "general-purpose",
          "rootDiskSize": 100,
          "RootDeviceName": "Not visible in IBM",
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
          "specId": "ibm+au-syd+bxf-4x16",
          "cspSpecName": "bxf-4x16",
          "spec": {
            "cspSpecName": "bxf-4x16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.235
          },
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "r026-7212abef-b5a8-4644-ac0b-d144865621fa",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252",
          "networkInterface": "nylon-shanty-cultural-motivate",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "r026-2088d198-ebc6-40fc-b315-1705d5317a4f",
          "nodeUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-07-01T05:34:20Z",
              "completedTime": "2026-07-01T05:34:33Z",
              "elapsedTime": 13,
              "resultSummary": "Command execution failed",
              "errorMessage": "failed to connect to target host via bastion after 3 attempts: failed to establish SSH connection to bastion host: ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Availability",
              "value": "{class:standard}"
            },
            {
              "key": "AvailabilityPolicy",
              "value": "{host_failure:restart,preemption:stop}"
            },
            {
              "key": "Bandwidth",
              "value": "8000"
            },
            {
              "key": "BootVolumeAttachment",
              "value": "{device:{id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030-mzszp},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/volume_attachments/02h7-0de201de-c73d-4894-b0e5-b71215cdd030,id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030,name:failing-vitally-headcount-overreach,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,id:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,name:favored-favorite-pavement-disowned,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-07-01T05:33:39.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab"
            },
            {
              "key": "EnableSecureBoot",
              "value": "false"
            },
            {
              "key": "HealthState",
              "value": "ok"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab"
            },
            {
              "key": "ID",
              "value": "02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab"
            },
            {
              "key": "Image",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
            },
            {
              "key": "LifecycleState",
              "value": "stable"
            },
            {
              "key": "Memory",
              "value": "16"
            },
            {
              "key": "MetadataService",
              "value": "{enabled:false,protocol:http,response_hop_limit:1}"
            },
            {
              "key": "Name",
              "value": "tbtddbjda2lllfuho1a9"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/network_interfaces/02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,id:02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,name:nylon-shanty-cultural-motivate,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-9398db33-011e-4736-bc8c-b195c7600795,id:02h7-9398db33-011e-4736-bc8c-b195c7600795,name:client-districts-impromptu-plated,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/network_interfaces/02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,id:02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,name:nylon-shanty-cultural-motivate,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-9398db33-011e-4736-bc8c-b195c7600795,id:02h7-9398db33-011e-4736-bc8c-b195c7600795,name:client-districts-impromptu-plated,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
            },
            {
              "key": "Profile",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16,name:bxf-4x16,resource_type:instance_profile}"
            },
            {
              "key": "ReservationAffinity",
              "value": "{policy:automatic,pool:[]}"
            },
            {
              "key": "ResourceGroup",
              "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
            },
            {
              "key": "ResourceType",
              "value": "instance"
            },
            {
              "key": "Startable",
              "value": "true"
            },
            {
              "key": "Status",
              "value": "running"
            },
            {
              "key": "TotalNetworkBandwidth",
              "value": "6000"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "2000"
            },
            {
              "key": "Vcpu",
              "value": "{architecture:amd64,count:4,manufacturer:intel,percentage:100}"
            },
            {
              "key": "VolumeAttachments",
              "value": "{device:{id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030-mzszp},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/volume_attachments/02h7-0de201de-c73d-4894-b0e5-b71215cdd030,id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030,name:failing-vitally-headcount-overreach,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,id:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,name:favored-favorite-pavement-disowned,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-7212abef-b5a8-4644-ac0b-d144865621fa,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-7212abef-b5a8-4644-ac0b-d144865621fa,id:r026-7212abef-b5a8-4644-ac0b-d144865621fa,name:tbs8qnakggjn5oera7dh,resource_type:vpc}"
            },
            {
              "key": "Zone",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-ng-influxdb-back-2",
          "uid": "tbt5ckjhndln9g8p23d5",
          "cspResourceName": "tbt5ckjhndln9g8p23d5",
          "cspResourceId": "02h7_ce163b62-c155-4950-802c-ec40da1bbee1",
          "name": "my-ng-influxdb-back-2",
          "nodeGroupId": "my-ng-influxdb-back",
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
          "createdTime": "2026-07-01 05:34:11",
          "label": {
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-07-01 05:34:11",
            "sys.cspResourceId": "02h7_ce163b62-c155-4950-802c-ec40da1bbee1",
            "sys.cspResourceName": "tbt5ckjhndln9g8p23d5",
            "sys.id": "my-ng-influxdb-back-2",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbt5ckjhndln9g8p23d5",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.97.167",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "general-purpose",
          "rootDiskSize": 100,
          "RootDeviceName": "Not visible in IBM",
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
          "specId": "ibm+au-syd+bxf-4x16",
          "cspSpecName": "bxf-4x16",
          "spec": {
            "cspSpecName": "bxf-4x16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.235
          },
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "r026-7212abef-b5a8-4644-ac0b-d144865621fa",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252",
          "networkInterface": "avenue-gallon-ignore-obsessive",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "r026-2088d198-ebc6-40fc-b315-1705d5317a4f",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJOHhJMvNzbWuoNueVZ358kak7qJBzvuvWje+VWWyQ5f9JijOOmFvFLmibjnypVZCmE9YFTQjGXPnQ7Gg1MkHmM=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:aCq7EJqGXuDYwtJ1OceEpVgmTrQ5RDlzebG586o+Sp0",
            "firstUsedAt": "2026-07-01T05:34:21Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-07-01T05:34:20Z",
              "completedTime": "2026-07-01T05:34:33Z",
              "elapsedTime": 13,
              "resultSummary": "Command execution failed",
              "errorMessage": "failed to connect to target host via bastion after 3 attempts: failed to establish SSH connection to bastion host: ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Availability",
              "value": "{class:standard}"
            },
            {
              "key": "AvailabilityPolicy",
              "value": "{host_failure:restart,preemption:stop}"
            },
            {
              "key": "Bandwidth",
              "value": "8000"
            },
            {
              "key": "BootVolumeAttachment",
              "value": "{device:{id:02h7-85724725-6251-4f51-8944-c502c2cc4a61-5wtfr},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/volume_attachments/02h7-85724725-6251-4f51-8944-c502c2cc4a61,id:02h7-85724725-6251-4f51-8944-c502c2cc4a61,name:bouquet-ducktail-grime-brilliant,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,id:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,name:composer-doubt-duckling-emptier,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-07-01T05:33:35.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_ce163b62-c155-4950-802c-ec40da1bbee1"
            },
            {
              "key": "EnableSecureBoot",
              "value": "false"
            },
            {
              "key": "HealthState",
              "value": "ok"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1"
            },
            {
              "key": "ID",
              "value": "02h7_ce163b62-c155-4950-802c-ec40da1bbee1"
            },
            {
              "key": "Image",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
            },
            {
              "key": "LifecycleState",
              "value": "stable"
            },
            {
              "key": "Memory",
              "value": "16"
            },
            {
              "key": "MetadataService",
              "value": "{enabled:false,protocol:http,response_hop_limit:1}"
            },
            {
              "key": "Name",
              "value": "tbt5ckjhndln9g8p23d5"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/network_interfaces/02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,id:02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,name:avenue-gallon-ignore-obsessive,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,id:02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,name:scrubbed-harmonica-getup-germproof,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/network_interfaces/02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,id:02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,name:avenue-gallon-ignore-obsessive,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,id:02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,name:scrubbed-harmonica-getup-germproof,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
            },
            {
              "key": "Profile",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16,name:bxf-4x16,resource_type:instance_profile}"
            },
            {
              "key": "ReservationAffinity",
              "value": "{policy:automatic,pool:[]}"
            },
            {
              "key": "ResourceGroup",
              "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
            },
            {
              "key": "ResourceType",
              "value": "instance"
            },
            {
              "key": "Startable",
              "value": "true"
            },
            {
              "key": "Status",
              "value": "running"
            },
            {
              "key": "TotalNetworkBandwidth",
              "value": "6000"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "2000"
            },
            {
              "key": "Vcpu",
              "value": "{architecture:amd64,count:4,manufacturer:intel,percentage:100}"
            },
            {
              "key": "VolumeAttachments",
              "value": "{device:{id:02h7-85724725-6251-4f51-8944-c502c2cc4a61-5wtfr},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/volume_attachments/02h7-85724725-6251-4f51-8944-c502c2cc4a61,id:02h7-85724725-6251-4f51-8944-c502c2cc4a61,name:bouquet-ducktail-grime-brilliant,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,id:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,name:composer-doubt-duckling-emptier,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-7212abef-b5a8-4644-ac0b-d144865621fa,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-7212abef-b5a8-4644-ac0b-d144865621fa,id:r026-7212abef-b5a8-4644-ac0b-d144865621fa,name:tbs8qnakggjn5oera7dh,resource_type:vpc}"
            },
            {
              "key": "Zone",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
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
            "nodeIp": "159.23.97.167",
            "command": {
              "0": "uname -a"
            },
            "stdout": {},
            "stderr": {},
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-1",
            "nodeIp": "159.23.102.107",
            "command": {
              "0": "uname -a"
            },
            "stdout": {},
            "stderr": {},
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "159.23.93.79",
            "command": {
              "0": "uname -a"
            },
            "stdout": {},
            "stderr": {},
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
  "uid": "tbphfu6hhvohsnkjlae0",
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
    "sys.uid": "tbphfu6hhvohsnkjlae0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbk5oqcsop91pbhic1u4",
      "cspResourceName": "tbk5oqcsop91pbhic1u4",
      "cspResourceId": "02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87",
      "name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-07-01 05:34:15",
      "label": {
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-01 05:34:15",
        "sys.cspResourceId": "02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87",
        "sys.cspResourceName": "tbk5oqcsop91pbhic1u4",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbk5oqcsop91pbhic1u4",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.93.79",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
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
      "specId": "ibm+au-syd+nxf-2x2",
      "cspSpecName": "nxf-2x2",
      "spec": {
        "cspSpecName": "nxf-2x2",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.094
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-7212abef-b5a8-4644-ac0b-d144865621fa",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252",
      "networkInterface": "splashy-sandpit-chatting-identical",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-2088d198-ebc6-40fc-b315-1705d5317a4f",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-07-01T05:34:20Z",
          "completedTime": "2026-07-01T05:34:33Z",
          "elapsedTime": 13,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target host via bastion after 3 attempts: failed to establish SSH connection to bastion host: ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "2000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a-5kbzm},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/volume_attachments/02h7-5b89e52a-26da-4145-bb82-32465b318a2a,id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a,name:hamper-graceful-nimbly-captain,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,id:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,name:calculate-cold-earflap-hurricane,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-01T05:33:35.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87"
        },
        {
          "key": "ID",
          "value": "02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "2"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbk5oqcsop91pbhic1u4"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/network_interfaces/02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,id:02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,name:splashy-sandpit-chatting-identical,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,id:02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,name:entourage-attendant-collard-postcard,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/network_interfaces/02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,id:02h7-8eabc17b-1423-4a1d-b436-8a3acd697001,name:splashy-sandpit-chatting-identical,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,id:02h7-7e1e606e-be3c-4cb2-9ab5-d6113fd03091,name:entourage-attendant-collard-postcard,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x2,name:nxf-2x2,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "1500"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "500"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a-5kbzm},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87/volume_attachments/02h7-5b89e52a-26da-4145-bb82-32465b318a2a,id:02h7-5b89e52a-26da-4145-bb82-32465b318a2a,name:hamper-graceful-nimbly-captain,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,id:r026-72284e3d-4fba-4375-bdfa-7c0c6a35f3ef,name:calculate-cold-earflap-hurricane,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-7212abef-b5a8-4644-ac0b-d144865621fa,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-7212abef-b5a8-4644-ac0b-d144865621fa,id:r026-7212abef-b5a8-4644-ac0b-d144865621fa,name:tbs8qnakggjn5oera7dh,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-1",
      "uid": "tbtddbjda2lllfuho1a9",
      "cspResourceName": "tbtddbjda2lllfuho1a9",
      "cspResourceId": "02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab",
      "name": "my-ng-influxdb-back-1",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-01 05:34:14",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-01 05:34:14",
        "sys.cspResourceId": "02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab",
        "sys.cspResourceName": "tbtddbjda2lllfuho1a9",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbtddbjda2lllfuho1a9",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.102.107",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
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
      "specId": "ibm+au-syd+bxf-4x16",
      "cspSpecName": "bxf-4x16",
      "spec": {
        "cspSpecName": "bxf-4x16",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.235
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-7212abef-b5a8-4644-ac0b-d144865621fa",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252",
      "networkInterface": "nylon-shanty-cultural-motivate",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-2088d198-ebc6-40fc-b315-1705d5317a4f",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-07-01T05:34:20Z",
          "completedTime": "2026-07-01T05:34:33Z",
          "elapsedTime": 13,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target host via bastion after 3 attempts: failed to establish SSH connection to bastion host: ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "8000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030-mzszp},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/volume_attachments/02h7-0de201de-c73d-4894-b0e5-b71215cdd030,id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030,name:failing-vitally-headcount-overreach,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,id:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,name:favored-favorite-pavement-disowned,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-01T05:33:39.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab"
        },
        {
          "key": "ID",
          "value": "02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "16"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbtddbjda2lllfuho1a9"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/network_interfaces/02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,id:02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,name:nylon-shanty-cultural-motivate,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-9398db33-011e-4736-bc8c-b195c7600795,id:02h7-9398db33-011e-4736-bc8c-b195c7600795,name:client-districts-impromptu-plated,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/network_interfaces/02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,id:02h7-18b7ae91-e286-4a09-a4d1-3f5bbfc9fde5,name:nylon-shanty-cultural-motivate,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-9398db33-011e-4736-bc8c-b195c7600795,id:02h7-9398db33-011e-4736-bc8c-b195c7600795,name:client-districts-impromptu-plated,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16,name:bxf-4x16,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "6000"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "2000"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:4,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030-mzszp},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab/volume_attachments/02h7-0de201de-c73d-4894-b0e5-b71215cdd030,id:02h7-0de201de-c73d-4894-b0e5-b71215cdd030,name:failing-vitally-headcount-overreach,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,id:r026-2c8e82fa-10a5-4bc0-8a6d-f7b1b490ff51,name:favored-favorite-pavement-disowned,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-7212abef-b5a8-4644-ac0b-d144865621fa,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-7212abef-b5a8-4644-ac0b-d144865621fa,id:r026-7212abef-b5a8-4644-ac0b-d144865621fa,name:tbs8qnakggjn5oera7dh,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-ng-influxdb-back-2",
      "uid": "tbt5ckjhndln9g8p23d5",
      "cspResourceName": "tbt5ckjhndln9g8p23d5",
      "cspResourceId": "02h7_ce163b62-c155-4950-802c-ec40da1bbee1",
      "name": "my-ng-influxdb-back-2",
      "nodeGroupId": "my-ng-influxdb-back",
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
      "createdTime": "2026-07-01 05:34:11",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-01 05:34:11",
        "sys.cspResourceId": "02h7_ce163b62-c155-4950-802c-ec40da1bbee1",
        "sys.cspResourceName": "tbt5ckjhndln9g8p23d5",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbt5ckjhndln9g8p23d5",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.97.167",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
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
      "specId": "ibm+au-syd+bxf-4x16",
      "cspSpecName": "bxf-4x16",
      "spec": {
        "cspSpecName": "bxf-4x16",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.235
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-7212abef-b5a8-4644-ac0b-d144865621fa",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252",
      "networkInterface": "avenue-gallon-ignore-obsessive",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-2088d198-ebc6-40fc-b315-1705d5317a4f",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJOHhJMvNzbWuoNueVZ358kak7qJBzvuvWje+VWWyQ5f9JijOOmFvFLmibjnypVZCmE9YFTQjGXPnQ7Gg1MkHmM=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:aCq7EJqGXuDYwtJ1OceEpVgmTrQ5RDlzebG586o+Sp0",
        "firstUsedAt": "2026-07-01T05:34:21Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-07-01T05:34:20Z",
          "completedTime": "2026-07-01T05:34:33Z",
          "elapsedTime": 13,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target host via bastion after 3 attempts: failed to establish SSH connection to bastion host: ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "8000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-85724725-6251-4f51-8944-c502c2cc4a61-5wtfr},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/volume_attachments/02h7-85724725-6251-4f51-8944-c502c2cc4a61,id:02h7-85724725-6251-4f51-8944-c502c2cc4a61,name:bouquet-ducktail-grime-brilliant,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,id:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,name:composer-doubt-duckling-emptier,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-01T05:33:35.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_ce163b62-c155-4950-802c-ec40da1bbee1"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1"
        },
        {
          "key": "ID",
          "value": "02h7_ce163b62-c155-4950-802c-ec40da1bbee1"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "16"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbt5ckjhndln9g8p23d5"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/network_interfaces/02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,id:02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,name:avenue-gallon-ignore-obsessive,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,id:02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,name:scrubbed-harmonica-getup-germproof,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/network_interfaces/02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,id:02h7-52b00b0c-5c7d-43c4-bc73-97ea8618ff17,name:avenue-gallon-ignore-obsessive,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252/reserved_ips/02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,id:02h7-694e7581-cd15-4648-b3d9-6fe89b6da788,name:scrubbed-harmonica-getup-germproof,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16,name:bxf-4x16,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "6000"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "2000"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:4,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-85724725-6251-4f51-8944-c502c2cc4a61-5wtfr},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_ce163b62-c155-4950-802c-ec40da1bbee1/volume_attachments/02h7-85724725-6251-4f51-8944-c502c2cc4a61,id:02h7-85724725-6251-4f51-8944-c502c2cc4a61,name:bouquet-ducktail-grime-brilliant,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,id:r026-9b6a4d43-a11c-4645-9d23-cd25a8bb5700,name:composer-doubt-duckling-emptier,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-7212abef-b5a8-4644-ac0b-d144865621fa,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-7212abef-b5a8-4644-ac0b-d144865621fa,id:r026-7212abef-b5a8-4644-ac0b-d144865621fa,name:tbs8qnakggjn5oera7dh,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
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
        "nodeIp": "159.23.97.167",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "159.23.102.107",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "159.23.93.79",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
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
      "cspResourceName": "tbpsuefvrv0l0vp1ejdr",
      "cspResourceId": "r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb",
      "name": "my-ng-influxdb-back",
      "connectionName": "ibm-au-syd",
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
        "keyValueList": [
          {
            "key": "Protocol",
            "value": "TCP"
          },
          {
            "key": "Port",
            "value": "9999"
          },
          {
            "key": "CspID",
            "value": "r026-66a973a0-55a7-469b-8658-4a1d62e0d347"
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
            "value": "[{NameId:tbtddbjda2lllfuho1a9,SystemId:02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab},{NameId:tbt5ckjhndln9g8p23d5,SystemId:02h7_ce163b62-c155-4950-802c-ec40da1bbee1}]"
          },
          {
            "key": "CspID",
            "value": "r026-1de74355-ea0d-4a5f-a825-71f20f0bf6bf"
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
      "createdTime": "2026-07-01T05:35:37Z",
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
          "value": "2026-07-01T05:35:37.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::load-balancer:r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb"
        },
        {
          "key": "FailsafePolicyActions",
          "value": "drop; forward"
        },
        {
          "key": "Hostname",
          "value": "618a9ae8-au-syd.lb.appdomain.cloud"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb"
        },
        {
          "key": "ID",
          "value": "r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb"
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
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb/listeners/r026-66a973a0-55a7-469b-8658-4a1d62e0d347,id:r026-66a973a0-55a7-469b-8658-4a1d62e0d347}"
        },
        {
          "key": "Logging",
          "value": "{datapath:{active:false}}"
        },
        {
          "key": "Name",
          "value": "tbpsuefvrv0l0vp1ejdr"
        },
        {
          "key": "OperatingStatus",
          "value": "offline"
        },
        {
          "key": "Pools",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb/pools/r026-1de74355-ea0d-4a5f-a825-71f20f0bf6bf,id:r026-1de74355-ea0d-4a5f-a825-71f20f0bf6bf,name:backend-8086-265874}"
        },
        {
          "key": "Profile",
          "value": "{family:network,href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancer/profiles/network-fixed,name:network-fixed}"
        },
        {
          "key": "ProvisioningStatus",
          "value": "create_pending"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:-}"
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
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r026-89d94b9f-bf61-4ce3-91ec-ba7adc4c4131,href:https://au-syd.iaas.cloud.ibm.com/v1/security_groups/r026-89d94b9f-bf61-4ce3-91ec-ba7adc4c4131,id:r026-89d94b9f-bf61-4ce3-91ec-ba7adc4c4131,name:sg-tbpsuefvrv0l0vp1ejdr}"
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
          "value": "{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}"
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
    "cspResourceName": "tbpsuefvrv0l0vp1ejdr",
    "cspResourceId": "r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb",
    "name": "my-ng-influxdb-back",
    "connectionName": "ibm-au-syd",
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
      "keyValueList": [
        {
          "key": "Protocol",
          "value": "TCP"
        },
        {
          "key": "Port",
          "value": "9999"
        },
        {
          "key": "CspID",
          "value": "r026-66a973a0-55a7-469b-8658-4a1d62e0d347"
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
          "value": "[{NameId:tbtddbjda2lllfuho1a9,SystemId:02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab},{NameId:tbt5ckjhndln9g8p23d5,SystemId:02h7_ce163b62-c155-4950-802c-ec40da1bbee1}]"
        },
        {
          "key": "CspID",
          "value": "r026-1de74355-ea0d-4a5f-a825-71f20f0bf6bf"
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
    "createdTime": "2026-07-01T05:35:37Z",
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
        "value": "2026-07-01T05:35:37.000Z"
      },
      {
        "key": "CRN",
        "value": "crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::load-balancer:r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb"
      },
      {
        "key": "FailsafePolicyActions",
        "value": "drop; forward"
      },
      {
        "key": "Hostname",
        "value": "618a9ae8-au-syd.lb.appdomain.cloud"
      },
      {
        "key": "Href",
        "value": "https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb"
      },
      {
        "key": "ID",
        "value": "r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb"
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
        "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb/listeners/r026-66a973a0-55a7-469b-8658-4a1d62e0d347,id:r026-66a973a0-55a7-469b-8658-4a1d62e0d347}"
      },
      {
        "key": "Logging",
        "value": "{datapath:{active:false}}"
      },
      {
        "key": "Name",
        "value": "tbpsuefvrv0l0vp1ejdr"
      },
      {
        "key": "OperatingStatus",
        "value": "offline"
      },
      {
        "key": "Pools",
        "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb/pools/r026-1de74355-ea0d-4a5f-a825-71f20f0bf6bf,id:r026-1de74355-ea0d-4a5f-a825-71f20f0bf6bf,name:backend-8086-265874}"
      },
      {
        "key": "Profile",
        "value": "{family:network,href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancer/profiles/network-fixed,name:network-fixed}"
      },
      {
        "key": "ProvisioningStatus",
        "value": "create_pending"
      },
      {
        "key": "ResourceGroup",
        "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:-}"
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
        "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r026-89d94b9f-bf61-4ce3-91ec-ba7adc4c4131,href:https://au-syd.iaas.cloud.ibm.com/v1/security_groups/r026-89d94b9f-bf61-4ce3-91ec-ba7adc4c4131,id:r026-89d94b9f-bf61-4ce3-91ec-ba7adc4c4131,name:sg-tbpsuefvrv0l0vp1ejdr}"
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
        "value": "{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}"
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
  "cspResourceName": "tbpsuefvrv0l0vp1ejdr",
  "cspResourceId": "r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb",
  "name": "my-ng-influxdb-back",
  "connectionName": "ibm-au-syd",
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
    "keyValueList": [
      {
        "key": "Protocol",
        "value": "TCP"
      },
      {
        "key": "Port",
        "value": "9999"
      },
      {
        "key": "CspID",
        "value": "r026-66a973a0-55a7-469b-8658-4a1d62e0d347"
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
        "value": "[{NameId:tbtddbjda2lllfuho1a9,SystemId:02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab},{NameId:tbt5ckjhndln9g8p23d5,SystemId:02h7_ce163b62-c155-4950-802c-ec40da1bbee1}]"
      },
      {
        "key": "CspID",
        "value": "r026-1de74355-ea0d-4a5f-a825-71f20f0bf6bf"
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
  "createdTime": "2026-07-01T05:35:37Z",
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
      "value": "2026-07-01T05:35:37.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::load-balancer:r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb"
    },
    {
      "key": "FailsafePolicyActions",
      "value": "drop; forward"
    },
    {
      "key": "Hostname",
      "value": "618a9ae8-au-syd.lb.appdomain.cloud"
    },
    {
      "key": "Href",
      "value": "https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb"
    },
    {
      "key": "ID",
      "value": "r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb"
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
      "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb/listeners/r026-66a973a0-55a7-469b-8658-4a1d62e0d347,id:r026-66a973a0-55a7-469b-8658-4a1d62e0d347}"
    },
    {
      "key": "Logging",
      "value": "{datapath:{active:false}}"
    },
    {
      "key": "Name",
      "value": "tbpsuefvrv0l0vp1ejdr"
    },
    {
      "key": "OperatingStatus",
      "value": "offline"
    },
    {
      "key": "Pools",
      "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-618a9ae8-2a53-4b53-80b5-d3d639b9bfbb/pools/r026-1de74355-ea0d-4a5f-a825-71f20f0bf6bf,id:r026-1de74355-ea0d-4a5f-a825-71f20f0bf6bf,name:backend-8086-265874}"
    },
    {
      "key": "Profile",
      "value": "{family:network,href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancer/profiles/network-fixed,name:network-fixed}"
    },
    {
      "key": "ProvisioningStatus",
      "value": "create_pending"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:-}"
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
      "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r026-89d94b9f-bf61-4ce3-91ec-ba7adc4c4131,href:https://au-syd.iaas.cloud.ibm.com/v1/security_groups/r026-89d94b9f-bf61-4ce3-91ec-ba7adc4c4131,id:r026-89d94b9f-bf61-4ce3-91ec-ba7adc4c4131,name:sg-tbpsuefvrv0l0vp1ejdr}"
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
      "value": "{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,id:02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252,name:tbp5tp86t9pcvurnl1bt,resource_type:subnet}"
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

**Generated At:** 2026-07-01 05:44:53

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
| **Description** | NLB-aware recommended infrastructure for cloud migration |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | IBM |
| **Target Region** | au-syd |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| nxf-2x2 | 2 | 2.0 | - | x86_64 |  | $0.0940 | 1 |
| bxf-4x16 | 4 | 16.0 | - | x86_64 |  | $0.2350 | 2 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| r026-c8e249d4-f148-4416-a3c6-555b7a02f67d | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | 02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87 | Running | 2 vCPU, 2.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 159.23.93.79<br>**Private IP:** 10.0.1.5<br>**SGs:** my-mig-sg-02<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-1 | 02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 159.23.102.107<br>**Private IP:** 10.0.1.6<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-2 | 02h7_ce163b62-c155-4950-802c-ec40da1bbee1 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 159.23.97.167<br>**Private IP:** 10.0.1.4<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-vnet-01 |
| **CSP VNet ID** | r026-7212abef-b5a8-4644-ac0b-d144865621fa |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-mig-subnet-01 | 02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252 | 10.0.1.0/24 | au-syd-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-mig-sshkey-01 | r026-2088d198-ebc6-40fc-b315-1705d5317a4f |  | SHA256:FgjlOFRS1MnzOV56AtNKu8a1izvF/6ZBK4S8B+2sHqI |

### Security Groups

#### Security Group: my-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-01 |
| **CSP Security Group ID** | r026-164ce32e-81ac-426a-a167-377f1f23ed5d |
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
| **CSP Security Group ID** | r026-63762c39-f4c5-48f4-8919-4321bb449b04 |
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
| **Per Hour** | $0.5640 |
| **Per Day** | $13.54 |
| **Per Month (30 days)** | $406.08 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| IBM | au-syd | 3 | $0.5640 | $406.08 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | nxf-2x2 | $0.0940 | $67.68 |
| my-ng-influxdb-back-1 | bxf-4x16 | $0.2350 | $169.20 |
| my-ng-influxdb-back-2 | bxf-4x16 | $0.2350 | $169.20 |




### Test Case 12: Migration Report

#### 12.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`

#### 12.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-07-01 05:44:58*

---

## 📊 Migration Summary

**Target Cloud:** IBM

**Target Region:** au-syd

**Namespace:** mig01 | **Infra ID:** my-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $406.08 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 2 | ✅ Selected | nxf-2x2, bxf-4x16 |
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
| 1 | **VM Name:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** 02h7_328ef6c6-2022-4e3f-80e6-8350bd9fbf87<br>**Label(sourceMachineId):** ng-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my-ng-influxdb-back-1<br>**VM ID:** 02h7_2ad476fc-d760-48c9-bc4b-6fab410225ab<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |
| 3 | **VM Name:** my-ng-influxdb-back-2<br>**VM ID:** 02h7_ce163b62-c155-4950-802c-ec40da1bbee1<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |

---

## ⚙️ VM Specs

**Summary:** 2 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** nxf-2x2<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2 | my-ng-influxdb-back-1 | **Spec ID:** bxf-4x16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3 | my-ng-influxdb-back-2 | **Spec ID:** bxf-4x16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** ng | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my-ng-influxdb-back-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my-ng-influxdb-back-2 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 9 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my-mig-sg-01

**CSP ID:** r026-164ce32e-81ac-426a-a167-377f1f23ed5d | **VNet:** my-mig-vnet-01 | **Rules:** 5

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

**CSP ID:** r026-63762c39-f4c5-48f4-8919-4321bb449b04 | **VNet:** my-mig-vnet-01 | **Rules:** 4

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
| 1 | **Name:** my-mig-vnet-01<br>**ID:** r026-7212abef-b5a8-4644-ac0b-d144865621fa | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my-mig-subnet-01<br>**ID:** 02h7-ec0a1ab9-6ef0-492c-a1a3-df8868f54252 | 10.0.1.0/24 | my-mig-vnet-01 |

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
| 1 | my-mig-sshkey-01 | r026-2088d198-ebc6-40fc-b315-1705d5317a4f | SHA256:FgjlOFRS1MnzOV56AtNKu8a1izvF/6ZBK4S8B+2sHqI | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.5640 |
| Daily | $13.54 |
| Monthly | $406.08 |
| Yearly | $4872.96 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | nxf-2x2 | $67.68 | 16.7% |

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

