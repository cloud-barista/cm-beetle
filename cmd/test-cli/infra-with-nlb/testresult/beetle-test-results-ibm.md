# CM-Beetle test results for IBM (with NLB)

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with IBM cloud infrastructure with NLBs.

## Environment and scenario

### Environment

- CM-Beetle: b418c24
- imdl: unknown
- CB-Tumblebug: Unknown (Fallback to Latest)
- CB-Spider: Unknown (Fallback to Latest)
- CB-MapUI: Unknown (Fallback to Latest)
- Target CSP: IBM
- Target Region: au-syd
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: July 13, 2026
- Test Time: 18:28:53 KST
- Test Execution: 2026-07-13 18:28:53 KST

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
| 1 | `POST /beetle/recommendation/infraWithNlb` | ✅ **PASS** | 12.3s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 3m19.568s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 36ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 4ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 14ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 21.563s | Pass |
| 7 | `POST /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 7m59.407s | Pass |
| 8 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb` | ✅ **PASS** | 4ms | Pass |
| 9 | `GET /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 8ms | Pass |
| 10 | NLB Load Balancing Verification | ✅ **PASS** | 2m16.981s | Pass |
| 11 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.232s | Pass |
| 12 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.278s | Pass |
| 13 | `DELETE /beetle/migration/middleware/ns/mig01/infra/{{infraId}}/nlb/{{nlbId}}` | ✅ **PASS** | 4m0.518s | Pass |
| 14 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m46.506s | Pass |

**Overall Result**: 14/14 tests passed ✅

**Total Duration**: 20m52.7592688s

*Test executed on July 13, 2026 at 18:28:53 KST (2026-07-13 18:28:53 KST) using CM-Beetle automated test CLI*

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
            "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
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
            "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
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
          "uid": "tbf87akpialuj0p3grka",
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
          "uid": "tbkqg2m0c8k80478q369",
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
          "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "regionList": [
            "au-syd"
          ],
          "id": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "uid": "tb2di1r42i2gj8q8kjov",
          "name": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ibm-au-syd",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:05:11 Mon",
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
              "value": "2026-06-17T03:40:05.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599"
            },
            {
              "key": "Encryption",
              "value": "none"
            },
            {
              "key": "File",
              "value": "{checksums:{sha256:a8dd44d0b27814db6ad5a0b845de74da2983d7dce60dd5a9014e87a94da7d250},size:1}"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599"
            },
            {
              "key": "ID",
              "value": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599"
            },
            {
              "key": "MinimumProvisionedSize",
              "value": "10"
            },
            {
              "key": "Name",
              "value": "ibm-ubuntu-22-04-5-minimal-amd64-16"
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
            "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
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
            "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
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
          "uid": "tbaum1bjl8ftjvlt35lv",
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
          "uid": "tb48j2ctufldbd268l58",
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
          "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "regionList": [
            "au-syd"
          ],
          "id": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "uid": "tb2di1r42i2gj8q8kjov",
          "name": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ibm-au-syd",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:05:11 Mon",
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
              "value": "2026-06-17T03:40:05.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599"
            },
            {
              "key": "Encryption",
              "value": "none"
            },
            {
              "key": "File",
              "value": "{checksums:{sha256:a8dd44d0b27814db6ad5a0b845de74da2983d7dce60dd5a9014e87a94da7d250},size:1}"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599"
            },
            {
              "key": "ID",
              "value": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599"
            },
            {
              "key": "MinimumProvisionedSize",
              "value": "10"
            },
            {
              "key": "Name",
              "value": "ibm-ubuntu-22-04-5-minimal-amd64-16"
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
  "uid": "tbkq84m33m8at0vs8vpq",
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
    "sys.uid": "tbkq84m33m8at0vs8vpq"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbr23d5skatofpm90u1n",
      "cspResourceName": "tbr23d5skatofpm90u1n",
      "cspResourceId": "02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3",
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
      "createdTime": "2026-07-13 09:31:18",
      "label": {
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-13 09:31:18",
        "sys.cspResourceId": "02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3",
        "sys.cspResourceName": "tbr23d5skatofpm90u1n",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbr23d5skatofpm90u1n",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.93.219",
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
      "specId": "ibm+au-syd+nxf-2x2",
      "cspSpecName": "nxf-2x2",
      "spec": {
        "cspSpecName": "nxf-2x2",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.094
      },
      "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-973a751e-d6bc-4105-a2b1-e76438bea613",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a",
      "networkInterface": "gaffe-hardship-courier-zookeeper",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGHVSvoLBvbbahHvw0Q+To1SmMT5yd1UL0/DdKeuFceKoRtAbQRtLbS7j/aGsDhbcruHnm19jrn+wuvKAyDILQU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:EDJ6KCd04b6QTkjbHntZYHgZcqzWVEc6nhN2ow7asng",
        "firstUsedAt": "2026-07-13T09:31:26Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:31:23Z",
          "completedTime": "2026-07-13T09:31:55Z",
          "elapsedTime": 32,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbr23d5skatofpm90u1n 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
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
          "value": "{device:{id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a-847qk},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/volume_attachments/02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,name:unstuffed-prolonged-phony-shady,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,id:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,name:spouse-average-goggles-evolve,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-13T09:30:41.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3"
        },
        {
          "key": "ID",
          "value": "02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,id:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,name:ibm-ubuntu-22-04-5-minimal-amd64-16,resource_type:image}"
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
          "value": "tbr23d5skatofpm90u1n"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_attachments/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,id:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,name:gaffe-hardship-courier-zookeeper,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_interfaces/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_attachments/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,id:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,name:gaffe-hardship-courier-zookeeper,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_interfaces/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a-847qk},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/volume_attachments/02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,name:unstuffed-prolonged-phony-shady,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,id:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,name:spouse-average-goggles-evolve,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-973a751e-d6bc-4105-a2b1-e76438bea613,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-973a751e-d6bc-4105-a2b1-e76438bea613,id:r026-973a751e-d6bc-4105-a2b1-e76438bea613,name:tb1docicnup03ohquors,resource_type:vpc}"
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
      "uid": "tb82d7cujrpndsnck3q4",
      "cspResourceName": "tb82d7cujrpndsnck3q4",
      "cspResourceId": "02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61",
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
      "createdTime": "2026-07-13 09:31:17",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-13 09:31:17",
        "sys.cspResourceId": "02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61",
        "sys.cspResourceName": "tb82d7cujrpndsnck3q4",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb82d7cujrpndsnck3q4",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.98.228",
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
      "specId": "ibm+au-syd+bxf-4x16",
      "cspSpecName": "bxf-4x16",
      "spec": {
        "cspSpecName": "bxf-4x16",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.235
      },
      "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-973a751e-d6bc-4105-a2b1-e76438bea613",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a",
      "networkInterface": "florist-recycled-upheld-athlete",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGY3ytTfPObqV/zkgpJMa/Vn8yI4lneqR+HAcvPFNZ9blgklXv9VLPUnMBeXKhfTKk1c84oe/rzhnCAXPpTDZ9w=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:uJvQp0hfNlsftX4qbKLyoYCuAS6vklGcCV/arXGLI64",
        "firstUsedAt": "2026-07-13T09:31:24Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:31:23Z",
          "completedTime": "2026-07-13T09:32:17Z",
          "elapsedTime": 54,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb82d7cujrpndsnck3q4 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
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
          "value": "{device:{id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596-z6w77},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/volume_attachments/02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,name:clothing-economic-snowplow-disloyal,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-5f07f60a-2545-4761-8bc6-db1dc2492800,id:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,name:salvage-antelope-canal-baritone,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-13T09:30:43.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61"
        },
        {
          "key": "ID",
          "value": "02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,id:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,name:ibm-ubuntu-22-04-5-minimal-amd64-16,resource_type:image}"
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
          "value": "tb82d7cujrpndsnck3q4"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_attachments/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-d3246a49-5d99-4598-a75d-109a09b6441a,id:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,name:florist-recycled-upheld-athlete,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_interfaces/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_attachments/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-d3246a49-5d99-4598-a75d-109a09b6441a,id:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,name:florist-recycled-upheld-athlete,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_interfaces/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596-z6w77},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/volume_attachments/02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,name:clothing-economic-snowplow-disloyal,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-5f07f60a-2545-4761-8bc6-db1dc2492800,id:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,name:salvage-antelope-canal-baritone,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-973a751e-d6bc-4105-a2b1-e76438bea613,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-973a751e-d6bc-4105-a2b1-e76438bea613,id:r026-973a751e-d6bc-4105-a2b1-e76438bea613,name:tb1docicnup03ohquors,resource_type:vpc}"
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
      "uid": "tbgr28uost9hr5phd913",
      "cspResourceName": "tbgr28uost9hr5phd913",
      "cspResourceId": "02h7_57bb1846-1e20-4101-b3a3-351f585f49d7",
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
      "createdTime": "2026-07-13 09:31:17",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-13 09:31:17",
        "sys.cspResourceId": "02h7_57bb1846-1e20-4101-b3a3-351f585f49d7",
        "sys.cspResourceName": "tbgr28uost9hr5phd913",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbgr28uost9hr5phd913",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.93.220",
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
      "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-973a751e-d6bc-4105-a2b1-e76438bea613",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a",
      "networkInterface": "swore-coastline-delicatessen-cinnamon",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBP1Kdp3VRf4Uw05dE5ZQBDj/0G50OwMEu5rhMtkVD2ZXcEKdy9VM+yu2WfOYnlN6ykoChsDS9/du+Njg2Gbgcxs=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:MgstMjyaUDe21tFjf2H4K3OWgbv07T2XkkbzL82yTnA",
        "firstUsedAt": "2026-07-13T09:31:26Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:31:23Z",
          "completedTime": "2026-07-13T09:32:15Z",
          "elapsedTime": 52,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbgr28uost9hr5phd913 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
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
          "value": "{device:{id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c-pxnf2},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/volume_attachments/02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,name:chicken-stinging-aspect-subduing,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-306512fd-be7b-44ef-8e71-111f1e4fde70,id:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,name:flier-purify-sixth-character,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-13T09:30:44.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_57bb1846-1e20-4101-b3a3-351f585f49d7"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7"
        },
        {
          "key": "ID",
          "value": "02h7_57bb1846-1e20-4101-b3a3-351f585f49d7"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,id:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,name:ibm-ubuntu-22-04-5-minimal-amd64-16,resource_type:image}"
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
          "value": "tbgr28uost9hr5phd913"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_attachments/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,id:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,name:swore-coastline-delicatessen-cinnamon,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_interfaces/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_attachments/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,id:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,name:swore-coastline-delicatessen-cinnamon,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_interfaces/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c-pxnf2},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/volume_attachments/02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,name:chicken-stinging-aspect-subduing,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-306512fd-be7b-44ef-8e71-111f1e4fde70,id:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,name:flier-purify-sixth-character,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-973a751e-d6bc-4105-a2b1-e76438bea613,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-973a751e-d6bc-4105-a2b1-e76438bea613,id:r026-973a751e-d6bc-4105-a2b1-e76438bea613,name:tb1docicnup03ohquors,resource_type:vpc}"
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
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "159.23.93.219",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbr23d5skatofpm90u1n 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "159.23.93.220",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbgr28uost9hr5phd913 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "159.23.98.228",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tb82d7cujrpndsnck3q4 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "tbkq84m33m8at0vs8vpq",
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
        "sys.uid": "tbkq84m33m8at0vs8vpq"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "NLB-aware recommended infrastructure for cloud migration",
      "node": [
        {
          "resourceType": "node",
          "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbr23d5skatofpm90u1n",
          "cspResourceName": "tbr23d5skatofpm90u1n",
          "cspResourceId": "02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3",
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
          "createdTime": "2026-07-13 09:31:18",
          "label": {
            "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-07-13 09:31:18",
            "sys.cspResourceId": "02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3",
            "sys.cspResourceName": "tbr23d5skatofpm90u1n",
            "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbr23d5skatofpm90u1n",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.93.219",
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
          "specId": "ibm+au-syd+nxf-2x2",
          "cspSpecName": "nxf-2x2",
          "spec": {
            "cspSpecName": "nxf-2x2",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.094
          },
          "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "r026-973a751e-d6bc-4105-a2b1-e76438bea613",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a",
          "networkInterface": "gaffe-hardship-courier-zookeeper",
          "securityGroupIds": [
            "my-mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGHVSvoLBvbbahHvw0Q+To1SmMT5yd1UL0/DdKeuFceKoRtAbQRtLbS7j/aGsDhbcruHnm19jrn+wuvKAyDILQU=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:EDJ6KCd04b6QTkjbHntZYHgZcqzWVEc6nhN2ow7asng",
            "firstUsedAt": "2026-07-13T09:31:26Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-13T09:31:23Z",
              "completedTime": "2026-07-13T09:31:55Z",
              "elapsedTime": 32,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbr23d5skatofpm90u1n 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
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
              "value": "{device:{id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a-847qk},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/volume_attachments/02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,name:unstuffed-prolonged-phony-shady,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,id:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,name:spouse-average-goggles-evolve,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-07-13T09:30:41.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3"
            },
            {
              "key": "ID",
              "value": "02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3"
            },
            {
              "key": "Image",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,id:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,name:ibm-ubuntu-22-04-5-minimal-amd64-16,resource_type:image}"
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
              "value": "tbr23d5skatofpm90u1n"
            },
            {
              "key": "NetworkAttachments",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_attachments/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,id:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,name:gaffe-hardship-courier-zookeeper,resource_type:virtual_network_interface}}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_interfaces/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkAttachment",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_attachments/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,id:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,name:gaffe-hardship-courier-zookeeper,resource_type:virtual_network_interface}}"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_interfaces/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a-847qk},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/volume_attachments/02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,name:unstuffed-prolonged-phony-shady,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,id:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,name:spouse-average-goggles-evolve,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-973a751e-d6bc-4105-a2b1-e76438bea613,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-973a751e-d6bc-4105-a2b1-e76438bea613,id:r026-973a751e-d6bc-4105-a2b1-e76438bea613,name:tb1docicnup03ohquors,resource_type:vpc}"
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
          "uid": "tb82d7cujrpndsnck3q4",
          "cspResourceName": "tb82d7cujrpndsnck3q4",
          "cspResourceId": "02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61",
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
          "createdTime": "2026-07-13 09:31:17",
          "label": {
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-07-13 09:31:17",
            "sys.cspResourceId": "02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61",
            "sys.cspResourceName": "tb82d7cujrpndsnck3q4",
            "sys.id": "my-ng-influxdb-back-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tb82d7cujrpndsnck3q4",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.98.228",
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
          "specId": "ibm+au-syd+bxf-4x16",
          "cspSpecName": "bxf-4x16",
          "spec": {
            "cspSpecName": "bxf-4x16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.235
          },
          "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "r026-973a751e-d6bc-4105-a2b1-e76438bea613",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a",
          "networkInterface": "florist-recycled-upheld-athlete",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGY3ytTfPObqV/zkgpJMa/Vn8yI4lneqR+HAcvPFNZ9blgklXv9VLPUnMBeXKhfTKk1c84oe/rzhnCAXPpTDZ9w=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:uJvQp0hfNlsftX4qbKLyoYCuAS6vklGcCV/arXGLI64",
            "firstUsedAt": "2026-07-13T09:31:24Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-13T09:31:23Z",
              "completedTime": "2026-07-13T09:32:17Z",
              "elapsedTime": 54,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tb82d7cujrpndsnck3q4 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
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
              "value": "{device:{id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596-z6w77},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/volume_attachments/02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,name:clothing-economic-snowplow-disloyal,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-5f07f60a-2545-4761-8bc6-db1dc2492800,id:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,name:salvage-antelope-canal-baritone,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-07-13T09:30:43.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61"
            },
            {
              "key": "ID",
              "value": "02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61"
            },
            {
              "key": "Image",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,id:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,name:ibm-ubuntu-22-04-5-minimal-amd64-16,resource_type:image}"
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
              "value": "tb82d7cujrpndsnck3q4"
            },
            {
              "key": "NetworkAttachments",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_attachments/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-d3246a49-5d99-4598-a75d-109a09b6441a,id:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,name:florist-recycled-upheld-athlete,resource_type:virtual_network_interface}}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_interfaces/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkAttachment",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_attachments/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-d3246a49-5d99-4598-a75d-109a09b6441a,id:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,name:florist-recycled-upheld-athlete,resource_type:virtual_network_interface}}"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_interfaces/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596-z6w77},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/volume_attachments/02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,name:clothing-economic-snowplow-disloyal,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-5f07f60a-2545-4761-8bc6-db1dc2492800,id:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,name:salvage-antelope-canal-baritone,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-973a751e-d6bc-4105-a2b1-e76438bea613,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-973a751e-d6bc-4105-a2b1-e76438bea613,id:r026-973a751e-d6bc-4105-a2b1-e76438bea613,name:tb1docicnup03ohquors,resource_type:vpc}"
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
          "uid": "tbgr28uost9hr5phd913",
          "cspResourceName": "tbgr28uost9hr5phd913",
          "cspResourceId": "02h7_57bb1846-1e20-4101-b3a3-351f585f49d7",
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
          "createdTime": "2026-07-13 09:31:17",
          "label": {
            "nlbBackend": "influxdb_back",
            "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-07-13 09:31:17",
            "sys.cspResourceId": "02h7_57bb1846-1e20-4101-b3a3-351f585f49d7",
            "sys.cspResourceName": "tbgr28uost9hr5phd913",
            "sys.id": "my-ng-influxdb-back-2",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-ng-influxdb-back-2",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-ng-influxdb-back",
            "sys.subnetId": "my-mig-subnet-01",
            "sys.uid": "tbgr28uost9hr5phd913",
            "sys.vNetId": "my-mig-vnet-01"
          },
          "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.93.220",
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
          "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my-mig-vnet-01",
          "cspVNetId": "r026-973a751e-d6bc-4105-a2b1-e76438bea613",
          "subnetId": "my-mig-subnet-01",
          "cspSubnetId": "02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a",
          "networkInterface": "swore-coastline-delicatessen-cinnamon",
          "securityGroupIds": [
            "my-mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-mig-sshkey-01",
          "cspSshKeyId": "r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBP1Kdp3VRf4Uw05dE5ZQBDj/0G50OwMEu5rhMtkVD2ZXcEKdy9VM+yu2WfOYnlN6ykoChsDS9/du+Njg2Gbgcxs=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:MgstMjyaUDe21tFjf2H4K3OWgbv07T2XkkbzL82yTnA",
            "firstUsedAt": "2026-07-13T09:31:26Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-07-13T09:31:23Z",
              "completedTime": "2026-07-13T09:32:15Z",
              "elapsedTime": 52,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbgr28uost9hr5phd913 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
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
              "value": "{device:{id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c-pxnf2},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/volume_attachments/02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,name:chicken-stinging-aspect-subduing,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-306512fd-be7b-44ef-8e71-111f1e4fde70,id:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,name:flier-purify-sixth-character,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-07-13T09:30:44.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_57bb1846-1e20-4101-b3a3-351f585f49d7"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7"
            },
            {
              "key": "ID",
              "value": "02h7_57bb1846-1e20-4101-b3a3-351f585f49d7"
            },
            {
              "key": "Image",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,id:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,name:ibm-ubuntu-22-04-5-minimal-amd64-16,resource_type:image}"
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
              "value": "tbgr28uost9hr5phd913"
            },
            {
              "key": "NetworkAttachments",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_attachments/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,id:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,name:swore-coastline-delicatessen-cinnamon,resource_type:virtual_network_interface}}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_interfaces/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkAttachment",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_attachments/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,id:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,name:swore-coastline-delicatessen-cinnamon,resource_type:virtual_network_interface}}"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_interfaces/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c-pxnf2},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/volume_attachments/02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,name:chicken-stinging-aspect-subduing,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-306512fd-be7b-44ef-8e71-111f1e4fde70,id:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,name:flier-purify-sixth-character,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-973a751e-d6bc-4105-a2b1-e76438bea613,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-973a751e-d6bc-4105-a2b1-e76438bea613,id:r026-973a751e-d6bc-4105-a2b1-e76438bea613,name:tb1docicnup03ohquors,resource_type:vpc}"
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
            "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "159.23.93.219",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbr23d5skatofpm90u1n 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-2",
            "nodeIp": "159.23.93.220",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbgr28uost9hr5phd913 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my-infra101",
            "nodeId": "my-ng-influxdb-back-1",
            "nodeIp": "159.23.98.228",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb82d7cujrpndsnck3q4 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "tbkq84m33m8at0vs8vpq",
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
    "sys.uid": "tbkq84m33m8at0vs8vpq"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "NLB-aware recommended infrastructure for cloud migration",
  "node": [
    {
      "resourceType": "node",
      "id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbr23d5skatofpm90u1n",
      "cspResourceName": "tbr23d5skatofpm90u1n",
      "cspResourceId": "02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3",
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
      "createdTime": "2026-07-13 09:31:18",
      "label": {
        "sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-13 09:31:18",
        "sys.cspResourceId": "02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3",
        "sys.cspResourceName": "tbr23d5skatofpm90u1n",
        "sys.id": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbr23d5skatofpm90u1n",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM 01 for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.93.219",
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
      "specId": "ibm+au-syd+nxf-2x2",
      "cspSpecName": "nxf-2x2",
      "spec": {
        "cspSpecName": "nxf-2x2",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.094
      },
      "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-973a751e-d6bc-4105-a2b1-e76438bea613",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a",
      "networkInterface": "gaffe-hardship-courier-zookeeper",
      "securityGroupIds": [
        "my-mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGHVSvoLBvbbahHvw0Q+To1SmMT5yd1UL0/DdKeuFceKoRtAbQRtLbS7j/aGsDhbcruHnm19jrn+wuvKAyDILQU=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:EDJ6KCd04b6QTkjbHntZYHgZcqzWVEc6nhN2ow7asng",
        "firstUsedAt": "2026-07-13T09:31:26Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:31:23Z",
          "completedTime": "2026-07-13T09:31:55Z",
          "elapsedTime": 32,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbr23d5skatofpm90u1n 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
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
          "value": "{device:{id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a-847qk},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/volume_attachments/02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,name:unstuffed-prolonged-phony-shady,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,id:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,name:spouse-average-goggles-evolve,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-13T09:30:41.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3"
        },
        {
          "key": "ID",
          "value": "02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,id:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,name:ibm-ubuntu-22-04-5-minimal-amd64-16,resource_type:image}"
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
          "value": "tbr23d5skatofpm90u1n"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_attachments/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,id:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,name:gaffe-hardship-courier-zookeeper,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_interfaces/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_attachments/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,id:02h7-aa44cc84-eda7-4030-9dbf-d0aa33481579,name:gaffe-hardship-courier-zookeeper,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/network_interfaces/02h7-d79c0886-7920-4900-b10b-558ac1811d62,id:02h7-d79c0886-7920-4900-b10b-558ac1811d62,name:trash-gloomily-enjoyer-arrival,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,id:02h7-3640a1af-b63b-46c3-aaa1-4768ba47bb64,name:popeyed-unpadded-moonscape-smugness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a-847qk},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3/volume_attachments/02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,id:02h7-64571f85-b57e-49ca-8f34-ac6a9729533a,name:unstuffed-prolonged-phony-shady,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,id:r026-fc7abd72-f10e-4fba-ae73-c12f18d21054,name:spouse-average-goggles-evolve,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-973a751e-d6bc-4105-a2b1-e76438bea613,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-973a751e-d6bc-4105-a2b1-e76438bea613,id:r026-973a751e-d6bc-4105-a2b1-e76438bea613,name:tb1docicnup03ohquors,resource_type:vpc}"
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
      "uid": "tb82d7cujrpndsnck3q4",
      "cspResourceName": "tb82d7cujrpndsnck3q4",
      "cspResourceId": "02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61",
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
      "createdTime": "2026-07-13 09:31:17",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-13 09:31:17",
        "sys.cspResourceId": "02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61",
        "sys.cspResourceName": "tb82d7cujrpndsnck3q4",
        "sys.id": "my-ng-influxdb-back-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tb82d7cujrpndsnck3q4",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.98.228",
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
      "specId": "ibm+au-syd+bxf-4x16",
      "cspSpecName": "bxf-4x16",
      "spec": {
        "cspSpecName": "bxf-4x16",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.235
      },
      "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-973a751e-d6bc-4105-a2b1-e76438bea613",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a",
      "networkInterface": "florist-recycled-upheld-athlete",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGY3ytTfPObqV/zkgpJMa/Vn8yI4lneqR+HAcvPFNZ9blgklXv9VLPUnMBeXKhfTKk1c84oe/rzhnCAXPpTDZ9w=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:uJvQp0hfNlsftX4qbKLyoYCuAS6vklGcCV/arXGLI64",
        "firstUsedAt": "2026-07-13T09:31:24Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:31:23Z",
          "completedTime": "2026-07-13T09:32:17Z",
          "elapsedTime": 54,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tb82d7cujrpndsnck3q4 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
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
          "value": "{device:{id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596-z6w77},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/volume_attachments/02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,name:clothing-economic-snowplow-disloyal,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-5f07f60a-2545-4761-8bc6-db1dc2492800,id:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,name:salvage-antelope-canal-baritone,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-13T09:30:43.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61"
        },
        {
          "key": "ID",
          "value": "02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,id:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,name:ibm-ubuntu-22-04-5-minimal-amd64-16,resource_type:image}"
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
          "value": "tb82d7cujrpndsnck3q4"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_attachments/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-d3246a49-5d99-4598-a75d-109a09b6441a,id:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,name:florist-recycled-upheld-athlete,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_interfaces/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_attachments/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-d3246a49-5d99-4598-a75d-109a09b6441a,id:02h7-d3246a49-5d99-4598-a75d-109a09b6441a,name:florist-recycled-upheld-athlete,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/network_interfaces/02h7-04dd1640-f90e-463f-b862-376b8d324af7,id:02h7-04dd1640-f90e-463f-b862-376b8d324af7,name:assure-ridden-decade-gravity,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,id:02h7-c1c49d92-4913-45c9-a88f-02bf941aea32,name:directed-anemone-polyester-virtuous,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596-z6w77},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61/volume_attachments/02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,id:02h7-c74bb95a-9c39-4635-b44b-8fa6b88b1596,name:clothing-economic-snowplow-disloyal,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-5f07f60a-2545-4761-8bc6-db1dc2492800,id:r026-5f07f60a-2545-4761-8bc6-db1dc2492800,name:salvage-antelope-canal-baritone,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-973a751e-d6bc-4105-a2b1-e76438bea613,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-973a751e-d6bc-4105-a2b1-e76438bea613,id:r026-973a751e-d6bc-4105-a2b1-e76438bea613,name:tb1docicnup03ohquors,resource_type:vpc}"
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
      "uid": "tbgr28uost9hr5phd913",
      "cspResourceName": "tbgr28uost9hr5phd913",
      "cspResourceId": "02h7_57bb1846-1e20-4101-b3a3-351f585f49d7",
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
      "createdTime": "2026-07-13 09:31:17",
      "label": {
        "nlbBackend": "influxdb_back",
        "sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-07-13 09:31:17",
        "sys.cspResourceId": "02h7_57bb1846-1e20-4101-b3a3-351f585f49d7",
        "sys.cspResourceName": "tbgr28uost9hr5phd913",
        "sys.id": "my-ng-influxdb-back-2",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-ng-influxdb-back-2",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-ng-influxdb-back",
        "sys.subnetId": "my-mig-subnet-01",
        "sys.uid": "tbgr28uost9hr5phd913",
        "sys.vNetId": "my-mig-vnet-01"
      },
      "description": "Recommended VM for NLB backend influxdb_back (2 nodes) | Match Rate: CPU=100.0% Memory=100.0% Image=80.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.93.220",
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
      "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my-mig-vnet-01",
      "cspVNetId": "r026-973a751e-d6bc-4105-a2b1-e76438bea613",
      "subnetId": "my-mig-subnet-01",
      "cspSubnetId": "02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a",
      "networkInterface": "swore-coastline-delicatessen-cinnamon",
      "securityGroupIds": [
        "my-mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-mig-sshkey-01",
      "cspSshKeyId": "r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBP1Kdp3VRf4Uw05dE5ZQBDj/0G50OwMEu5rhMtkVD2ZXcEKdy9VM+yu2WfOYnlN6ykoChsDS9/du+Njg2Gbgcxs=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:MgstMjyaUDe21tFjf2H4K3OWgbv07T2XkkbzL82yTnA",
        "firstUsedAt": "2026-07-13T09:31:26Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-07-13T09:31:23Z",
          "completedTime": "2026-07-13T09:32:15Z",
          "elapsedTime": 52,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbgr28uost9hr5phd913 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
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
          "value": "{device:{id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c-pxnf2},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/volume_attachments/02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,name:chicken-stinging-aspect-subduing,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-306512fd-be7b-44ef-8e71-111f1e4fde70,id:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,name:flier-purify-sixth-character,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-07-13T09:30:44.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_57bb1846-1e20-4101-b3a3-351f585f49d7"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7"
        },
        {
          "key": "ID",
          "value": "02h7_57bb1846-1e20-4101-b3a3-351f585f49d7"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,id:r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599,name:ibm-ubuntu-22-04-5-minimal-amd64-16,resource_type:image}"
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
          "value": "tbgr28uost9hr5phd913"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_attachments/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,id:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,name:swore-coastline-delicatessen-cinnamon,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_interfaces/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_attachments/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,id:02h7-057af767-3498-4ad0-87a9-cf26cd6dfc40,name:swore-coastline-delicatessen-cinnamon,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/network_interfaces/02h7-60775056-1082-464a-9e4b-fb4d6691fe89,id:02h7-60775056-1082-464a-9e4b-fb4d6691fe89,name:corse-pulsate-shakable-diffused,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,id:02h7-da1c0501-8db5-40ab-9ea3-ddeec89ccb62,name:join-untainted-distract-grain,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c-pxnf2},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_57bb1846-1e20-4101-b3a3-351f585f49d7/volume_attachments/02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,id:02h7-b239c258-a4d6-4acc-a1bc-c03f9507140c,name:chicken-stinging-aspect-subduing,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-306512fd-be7b-44ef-8e71-111f1e4fde70,id:r026-306512fd-be7b-44ef-8e71-111f1e4fde70,name:flier-purify-sixth-character,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-973a751e-d6bc-4105-a2b1-e76438bea613,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-973a751e-d6bc-4105-a2b1-e76438bea613,id:r026-973a751e-d6bc-4105-a2b1-e76438bea613,name:tb1docicnup03ohquors,resource_type:vpc}"
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
        "nodeId": "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "159.23.93.219",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbr23d5skatofpm90u1n 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-2",
        "nodeIp": "159.23.93.220",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbgr28uost9hr5phd913 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my-infra101",
        "nodeId": "my-ng-influxdb-back-1",
        "nodeIp": "159.23.98.228",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tb82d7cujrpndsnck3q4 5.15.0-1103-ibm #106-Ubuntu SMP Mon May 25 12:14:22 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "cspResourceName": "tbt6tq5tir95ggq2845m",
      "cspResourceId": "r026-ef36ca07-bf7b-408d-9986-86a5c72505a3",
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
        "ip": "159.23.102.198",
        "port": "9999",
        "dnsName": "ef36ca07-au-syd.lb.appdomain.cloud",
        "keyValueList": [
          {
            "key": "Protocol",
            "value": "TCP"
          },
          {
            "key": "IP",
            "value": "159.23.102.198"
          },
          {
            "key": "Port",
            "value": "9999"
          },
          {
            "key": "DNSName",
            "value": "ef36ca07-au-syd.lb.appdomain.cloud"
          },
          {
            "key": "CspID",
            "value": "r026-5eb04d42-56cd-4983-84d0-2adff594e988"
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
            "key": "Protocol",
            "value": "TCP"
          },
          {
            "key": "Port",
            "value": "8086"
          },
          {
            "key": "VMs",
            "value": "[{NameId:tbgr28uost9hr5phd913,SystemId:02h7_57bb1846-1e20-4101-b3a3-351f585f49d7},{NameId:tb82d7cujrpndsnck3q4,SystemId:02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61}]"
          },
          {
            "key": "CspID",
            "value": "r026-d2e8ae87-c680-4cf6-8bf3-7d02e1af55f4"
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
      "createdTime": "2026-07-13T09:33:21Z",
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
          "value": "2026-07-13T09:33:21.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::load-balancer:r026-ef36ca07-bf7b-408d-9986-86a5c72505a3"
        },
        {
          "key": "FailsafePolicyActions",
          "value": "drop; forward"
        },
        {
          "key": "Hostname",
          "value": "ef36ca07-au-syd.lb.appdomain.cloud"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-ef36ca07-bf7b-408d-9986-86a5c72505a3"
        },
        {
          "key": "ID",
          "value": "r026-ef36ca07-bf7b-408d-9986-86a5c72505a3"
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
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-ef36ca07-bf7b-408d-9986-86a5c72505a3/listeners/r026-5eb04d42-56cd-4983-84d0-2adff594e988,id:r026-5eb04d42-56cd-4983-84d0-2adff594e988}"
        },
        {
          "key": "Logging",
          "value": "{datapath:{active:false}}"
        },
        {
          "key": "Name",
          "value": "tbt6tq5tir95ggq2845m"
        },
        {
          "key": "OperatingStatus",
          "value": "online"
        },
        {
          "key": "Pools",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-ef36ca07-bf7b-408d-9986-86a5c72505a3/pools/r026-d2e8ae87-c680-4cf6-8bf3-7d02e1af55f4,id:r026-d2e8ae87-c680-4cf6-8bf3-7d02e1af55f4,name:backend-8086-816605}"
        },
        {
          "key": "PrivateIps",
          "value": "{address:10.0.1.7,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-8e604922-e070-46d4-87f6-a9c376aa6aab,id:02h7-8e604922-e070-46d4-87f6-a9c376aa6aab,name:everyday-flashcard-console-unfairly,resource_type:subnet_reserved_ip}; {address:10.0.1.8,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-74336c26-76e8-4808-a836-0776e7e80b0f,id:02h7-74336c26-76e8-4808-a836-0776e7e80b0f,name:unchanged-gowl-upbeat-disregard,resource_type:subnet_reserved_ip}"
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
          "value": "{address:159.23.102.198}"
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
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r026-a67d64e8-97fb-45e1-91d2-852b3ea5ee1f,href:https://au-syd.iaas.cloud.ibm.com/v1/security_groups/r026-a67d64e8-97fb-45e1-91d2-852b3ea5ee1f,id:r026-a67d64e8-97fb-45e1-91d2-852b3ea5ee1f,name:sg-tbt6tq5tir95ggq2845m}"
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
          "value": "{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}"
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
    "cspResourceName": "tbt6tq5tir95ggq2845m",
    "cspResourceId": "r026-ef36ca07-bf7b-408d-9986-86a5c72505a3",
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
      "ip": "159.23.102.198",
      "port": "9999",
      "dnsName": "ef36ca07-au-syd.lb.appdomain.cloud",
      "keyValueList": [
        {
          "key": "Protocol",
          "value": "TCP"
        },
        {
          "key": "IP",
          "value": "159.23.102.198"
        },
        {
          "key": "Port",
          "value": "9999"
        },
        {
          "key": "DNSName",
          "value": "ef36ca07-au-syd.lb.appdomain.cloud"
        },
        {
          "key": "CspID",
          "value": "r026-5eb04d42-56cd-4983-84d0-2adff594e988"
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
          "key": "Protocol",
          "value": "TCP"
        },
        {
          "key": "Port",
          "value": "8086"
        },
        {
          "key": "VMs",
          "value": "[{NameId:tbgr28uost9hr5phd913,SystemId:02h7_57bb1846-1e20-4101-b3a3-351f585f49d7},{NameId:tb82d7cujrpndsnck3q4,SystemId:02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61}]"
        },
        {
          "key": "CspID",
          "value": "r026-d2e8ae87-c680-4cf6-8bf3-7d02e1af55f4"
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
    "createdTime": "2026-07-13T09:33:21Z",
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
        "value": "2026-07-13T09:33:21.000Z"
      },
      {
        "key": "CRN",
        "value": "crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::load-balancer:r026-ef36ca07-bf7b-408d-9986-86a5c72505a3"
      },
      {
        "key": "FailsafePolicyActions",
        "value": "drop; forward"
      },
      {
        "key": "Hostname",
        "value": "ef36ca07-au-syd.lb.appdomain.cloud"
      },
      {
        "key": "Href",
        "value": "https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-ef36ca07-bf7b-408d-9986-86a5c72505a3"
      },
      {
        "key": "ID",
        "value": "r026-ef36ca07-bf7b-408d-9986-86a5c72505a3"
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
        "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-ef36ca07-bf7b-408d-9986-86a5c72505a3/listeners/r026-5eb04d42-56cd-4983-84d0-2adff594e988,id:r026-5eb04d42-56cd-4983-84d0-2adff594e988}"
      },
      {
        "key": "Logging",
        "value": "{datapath:{active:false}}"
      },
      {
        "key": "Name",
        "value": "tbt6tq5tir95ggq2845m"
      },
      {
        "key": "OperatingStatus",
        "value": "online"
      },
      {
        "key": "Pools",
        "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-ef36ca07-bf7b-408d-9986-86a5c72505a3/pools/r026-d2e8ae87-c680-4cf6-8bf3-7d02e1af55f4,id:r026-d2e8ae87-c680-4cf6-8bf3-7d02e1af55f4,name:backend-8086-816605}"
      },
      {
        "key": "PrivateIps",
        "value": "{address:10.0.1.7,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-8e604922-e070-46d4-87f6-a9c376aa6aab,id:02h7-8e604922-e070-46d4-87f6-a9c376aa6aab,name:everyday-flashcard-console-unfairly,resource_type:subnet_reserved_ip}; {address:10.0.1.8,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-74336c26-76e8-4808-a836-0776e7e80b0f,id:02h7-74336c26-76e8-4808-a836-0776e7e80b0f,name:unchanged-gowl-upbeat-disregard,resource_type:subnet_reserved_ip}"
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
        "value": "{address:159.23.102.198}"
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
        "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r026-a67d64e8-97fb-45e1-91d2-852b3ea5ee1f,href:https://au-syd.iaas.cloud.ibm.com/v1/security_groups/r026-a67d64e8-97fb-45e1-91d2-852b3ea5ee1f,id:r026-a67d64e8-97fb-45e1-91d2-852b3ea5ee1f,name:sg-tbt6tq5tir95ggq2845m}"
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
        "value": "{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}"
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
  "cspResourceName": "tbt6tq5tir95ggq2845m",
  "cspResourceId": "r026-ef36ca07-bf7b-408d-9986-86a5c72505a3",
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
    "ip": "159.23.102.198",
    "port": "9999",
    "dnsName": "ef36ca07-au-syd.lb.appdomain.cloud",
    "keyValueList": [
      {
        "key": "Protocol",
        "value": "TCP"
      },
      {
        "key": "IP",
        "value": "159.23.102.198"
      },
      {
        "key": "Port",
        "value": "9999"
      },
      {
        "key": "DNSName",
        "value": "ef36ca07-au-syd.lb.appdomain.cloud"
      },
      {
        "key": "CspID",
        "value": "r026-5eb04d42-56cd-4983-84d0-2adff594e988"
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
        "key": "Protocol",
        "value": "TCP"
      },
      {
        "key": "Port",
        "value": "8086"
      },
      {
        "key": "VMs",
        "value": "[{NameId:tbgr28uost9hr5phd913,SystemId:02h7_57bb1846-1e20-4101-b3a3-351f585f49d7},{NameId:tb82d7cujrpndsnck3q4,SystemId:02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61}]"
      },
      {
        "key": "CspID",
        "value": "r026-d2e8ae87-c680-4cf6-8bf3-7d02e1af55f4"
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
  "createdTime": "2026-07-13T09:33:21Z",
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
      "value": "2026-07-13T09:33:21.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::load-balancer:r026-ef36ca07-bf7b-408d-9986-86a5c72505a3"
    },
    {
      "key": "FailsafePolicyActions",
      "value": "drop; forward"
    },
    {
      "key": "Hostname",
      "value": "ef36ca07-au-syd.lb.appdomain.cloud"
    },
    {
      "key": "Href",
      "value": "https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-ef36ca07-bf7b-408d-9986-86a5c72505a3"
    },
    {
      "key": "ID",
      "value": "r026-ef36ca07-bf7b-408d-9986-86a5c72505a3"
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
      "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-ef36ca07-bf7b-408d-9986-86a5c72505a3/listeners/r026-5eb04d42-56cd-4983-84d0-2adff594e988,id:r026-5eb04d42-56cd-4983-84d0-2adff594e988}"
    },
    {
      "key": "Logging",
      "value": "{datapath:{active:false}}"
    },
    {
      "key": "Name",
      "value": "tbt6tq5tir95ggq2845m"
    },
    {
      "key": "OperatingStatus",
      "value": "online"
    },
    {
      "key": "Pools",
      "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/load_balancers/r026-ef36ca07-bf7b-408d-9986-86a5c72505a3/pools/r026-d2e8ae87-c680-4cf6-8bf3-7d02e1af55f4,id:r026-d2e8ae87-c680-4cf6-8bf3-7d02e1af55f4,name:backend-8086-816605}"
    },
    {
      "key": "PrivateIps",
      "value": "{address:10.0.1.7,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-8e604922-e070-46d4-87f6-a9c376aa6aab,id:02h7-8e604922-e070-46d4-87f6-a9c376aa6aab,name:everyday-flashcard-console-unfairly,resource_type:subnet_reserved_ip}; {address:10.0.1.8,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a/reserved_ips/02h7-74336c26-76e8-4808-a836-0776e7e80b0f,id:02h7-74336c26-76e8-4808-a836-0776e7e80b0f,name:unchanged-gowl-upbeat-disregard,resource_type:subnet_reserved_ip}"
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
      "value": "{address:159.23.102.198}"
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
      "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r026-a67d64e8-97fb-45e1-91d2-852b3ea5ee1f,href:https://au-syd.iaas.cloud.ibm.com/v1/security_groups/r026-a67d64e8-97fb-45e1-91d2-852b3ea5ee1f,id:r026-a67d64e8-97fb-45e1-91d2-852b3ea5ee1f,name:sg-tbt6tq5tir95ggq2845m}"
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
      "value": "{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,id:02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a,name:tbf2ljcg349itnii06m3,resource_type:subnet}"
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

**Generated At:** 2026-07-13 09:42:43

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
| r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599 | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | 02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3 | Running | 2 vCPU, 2.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 159.23.93.219<br>**Private IP:** 10.0.1.4<br>**SGs:** my-mig-sg-02<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-1 | 02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 159.23.98.228<br>**Private IP:** 10.0.1.5<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |
| my-ng-influxdb-back-2 | 02h7_57bb1846-1e20-4101-b3a3-351f585f49d7 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-mig-vnet-01<br>**Subnet:** my-mig-subnet-01<br>**Public IP:** 159.23.93.220<br>**Private IP:** 10.0.1.6<br>**SGs:** my-mig-sg-01<br>**SSH:** my-mig-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-mig-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-vnet-01 |
| **CSP VNet ID** | r026-973a751e-d6bc-4105-a2b1-e76438bea613 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-mig-subnet-01 | 02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a | 10.0.1.0/24 | au-syd-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-mig-sshkey-01 | r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0 |  | SHA256:W9+bkBcpsLeBdDiDOPUXfFQJ/q1cPg31WqKwqVudX/E |

### Security Groups

#### Security Group: my-mig-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-mig-sg-01 |
| **CSP Security Group ID** | r026-e9472d3e-f8b2-4c3e-a8e5-be6a646610fa |
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
| **CSP Security Group ID** | r026-95fb8ca5-af0a-4f5c-b195-cee7dfe0f2d8 |
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

*Report generated: 2026-07-13 09:42:49*

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
| 1 | **VM Name:** my-ng-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** 02h7_86aaa57b-7093-43cd-8b00-111ae902a1e3<br>**Label(sourceMachineId):** ng-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my-ng-influxdb-back-1<br>**VM ID:** 02h7_35a2336c-2360-41bc-9d1a-c0329c7aea61<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |
| 3 | **VM Name:** my-ng-influxdb-back-2<br>**VM ID:** 02h7_57bb1846-1e20-4101-b3a3-351f585f49d7<br>**Label(sourceMachineId):** ng | **Hostname:** N/A<br>**Machine ID:** ng |

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
| 1 | my-ng-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** ng-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my-ng-influxdb-back-1 | **Image ID:** r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my-ng-influxdb-back-2 | **Image ID:** r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** ng | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 2 security group(s) with 9 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my-mig-sg-01

**CSP ID:** r026-e9472d3e-f8b2-4c3e-a8e5-be6a646610fa | **VNet:** my-mig-vnet-01 | **Rules:** 5

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

**CSP ID:** r026-95fb8ca5-af0a-4f5c-b195-cee7dfe0f2d8 | **VNet:** my-mig-vnet-01 | **Rules:** 4

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
| 1 | **Name:** my-mig-vnet-01<br>**ID:** r026-973a751e-d6bc-4105-a2b1-e76438bea613 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my-mig-subnet-01<br>**ID:** 02h7-4890848c-4a5a-414f-9ca3-3e72ab12c84a | 10.0.1.0/24 | my-mig-vnet-01 |

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
| 1 | my-mig-sshkey-01 | r026-b3b9388b-6054-4ef4-9d25-d4a0b45ffac0 | SHA256:W9+bkBcpsLeBdDiDOPUXfFQJ/q1cPg31WqKwqVudX/E | Used by all 3 VMs |

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

