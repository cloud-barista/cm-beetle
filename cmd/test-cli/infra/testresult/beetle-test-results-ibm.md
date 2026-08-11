# CM-Beetle test results for IBM

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with IBM cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.5.9+ (83b6bb2)
- imdl: v0.1.10+ (83b6bb2)
- CB-Tumblebug: v0.12.30
- CB-Spider: v0.12.42
- CB-MapUI: v0.12.56
- Target CSP: IBM
- Target Region: au-syd
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: August 11, 2026
- Test Time: 13:37:08 KST
- Test Execution: 2026-08-11 13:37:08 KST

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Validate the target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Remote Command Accessibility Check
1. Target Infrastructure Summary via Beetle
1. Migration Report via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for IBM

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 10.74s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 446ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 3m54.312s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 64ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 18ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 26ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 15.953s | Pass |
| 8 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.508s | Pass |
| 9 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.67s | Pass |
| 10 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m9.949s | Pass |

**Overall Result**: 10/10 tests passed ✅

**Total Duration**: 6m48.046506935s

*Test executed on August 11, 2026 at 13:37:08 KST (2026-08-11 13:37:08 KST) using CM-Beetle automated test CLI*

---

## Detailed Test Case Results

> [!INFO]
> This section provides detailed information for each test case, including API request information and response details.

### Test Case 1: Recommend a target model for computing infra

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/recommendation/infra`
- **Purpose**: Get infrastructure recommendations for migration
- **Required Parameters**: `desiredCsp` and `desiredRegion` in request body

**Request Body**:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "desiredCspAndRegionPair": {
    "csp": "ibm",
    "region": "au-syd"
  },
  "OnpremiseInfraModel": {
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
          "type": "",
          "totalSize": 0
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
            "destination": "10.0.0.2/32",
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
          },
          {
            "destination": "10.0.1.1/32",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::1/128",
            "gateway": "on-link",
            "interface": "lo",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::/0",
            "gateway": "on-link",
            "interface": "lo",
            "metric": 2147483647,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::1/128",
            "gateway": "on-link",
            "interface": "lo",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::6f:deff:fefc:71b1/128",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::/0",
            "gateway": "on-link",
            "interface": "lo",
            "metric": 2147483647,
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
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "icmp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "67",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "68",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "224.0.0.251/32",
            "dstPorts": "5353",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "239.255.255.250/32",
            "dstPorts": "1900",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
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
            "dstPorts": "80",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "443",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8080",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "3306",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "5432",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9113",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9113",
            "protocol": "udp",
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
            "dstPorts": "23",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "135",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "139",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "445",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "tcp",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "udp",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "fe80::/10",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "fe80::/10",
            "srcPorts": "547",
            "dstCIDR": "fe80::/10",
            "dstPorts": "546",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "ff02::fb/128",
            "dstPorts": "5353",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "ff02::f/128",
            "dstPorts": "1900",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "22",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "80",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "443",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "8080",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "3306",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "5432",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "23",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "135",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "139",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "445",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "fe80::/10",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "tcp",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "udp",
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
          "type": "",
          "totalSize": 0
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
            "destination": "10.0.0.2/32",
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
          },
          {
            "destination": "10.0.1.1/32",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::1/128",
            "gateway": "on-link",
            "interface": "lo",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::/0",
            "gateway": "on-link",
            "interface": "lo",
            "metric": 2147483647,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::1/128",
            "gateway": "on-link",
            "interface": "lo",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::8:96ff:fe7d:f417/128",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::/0",
            "gateway": "on-link",
            "interface": "lo",
            "metric": 2147483647,
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
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "icmp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "67",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "68",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "224.0.0.251/32",
            "dstPorts": "5353",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "239.255.255.250/32",
            "dstPorts": "1900",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
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
            "dstPorts": "2049",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "2049",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "111",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "111",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "20048",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "20048",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "32803",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "32803",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "80",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "443",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9100",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9100",
            "protocol": "udp",
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
            "dstPorts": "23",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "135",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "139",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "445",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "tcp",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "udp",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "fe80::/10",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "fe80::/10",
            "srcPorts": "547",
            "dstCIDR": "fe80::/10",
            "dstPorts": "546",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "ff02::fb/128",
            "dstPorts": "5353",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "ff02::f/128",
            "dstPorts": "1900",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "22",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "2049",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "2049",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "111",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "111",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "80",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "443",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "23",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "135",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "139",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "445",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "fe80::/10",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "tcp",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "udp",
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
          "cores": 1,
          "threads": 2,
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
          "type": "",
          "totalSize": 0
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
            "destination": "10.0.0.2/32",
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
          },
          {
            "destination": "10.0.1.1/32",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::1/128",
            "gateway": "on-link",
            "interface": "lo",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::/0",
            "gateway": "on-link",
            "interface": "lo",
            "metric": 2147483647,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::1/128",
            "gateway": "on-link",
            "interface": "lo",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::bf:6eff:fe6c:6e31/128",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "gateway": "10.0.1.1",
            "interface": "ens5",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "::/0",
            "gateway": "on-link",
            "interface": "lo",
            "metric": 2147483647,
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
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "icmp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "67",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "68",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "224.0.0.251/32",
            "dstPorts": "5353",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "239.255.255.250/32",
            "dstPorts": "1900",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
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
            "dstPorts": "3306",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "3306",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "4567",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "4567",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "4568",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "4568",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "4444",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "4444",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "80",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "443",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8080",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "3306",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "3306",
            "protocol": "udp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9104",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "10.0.0.0/16",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9104",
            "protocol": "udp",
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
            "dstPorts": "23",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "135",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "139",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "445",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "tcp",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "udp",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "fe80::/10",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "fe80::/10",
            "srcPorts": "547",
            "dstCIDR": "fe80::/10",
            "dstPorts": "546",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "ff02::fb/128",
            "dstPorts": "5353",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "ff02::f/128",
            "dstPorts": "1900",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "22",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "80",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "443",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "8080",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "3306",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "3306",
            "protocol": "udp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "23",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "135",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "139",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "445",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "*",
            "direction": "outbound",
            "action": "deny"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "fe80::/10",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "icmpv6",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "tcp",
            "direction": "outbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "udp",
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
      "description": "Candidate #1 | highly-matched | Overall Match Rate: Min=100.0% Max=100.0% Avg=100.0% | VMs: 3 total, 3 matched, 0 acceptable",
      "targetCloud": {
        "csp": "ibm",
        "region": "au-syd"
      },
      "targetInfra": {
        "name": "infra101",
        "installMonAgent": "",
        "label": null,
        "systemLabel": "",
        "description": "Recommended VMs comprising multi-cloud infrastructure",
        "nodeGroups": [
          {
            "name": "vm-ec268ed7-821e-9d73-e79f-961262161624",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+nxf-2x2",
            "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bxf-4x16",
            "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bxf-2x8",
            "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "ibm-au-syd",
        "cidrBlock": "10.0.0.0/21",
        "subnetInfoList": [
          {
            "name": "subnet-01",
            "ipv4_CIDR": "10.0.1.0/24",
            "description": "a recommended subnet for migration"
          }
        ],
        "description": "a recommended vNet for migration"
      },
      "targetSshKey": {
        "name": "sshkey-01",
        "connectionName": "ibm-au-syd",
        "description": "a SSH Key pair for migration (Note - provided ONLY once, MUST be downloaded",
        "cspResourceId": "",
        "fingerprint": "",
        "username": "",
        "verifiedUsername": "",
        "publicKey": "",
        "privateKey": ""
      },
      "targetSpecList": [
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
        },
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
          "id": "ibm+au-syd+bxf-2x8",
          "uid": "tb8moamua469rnvg885i",
          "cspSpecName": "bxf-2x8",
          "name": "ibm+au-syd+bxf-2x8",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.117,
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
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:8}"
            },
            {
              "key": "Name",
              "value": "bxf-2x8"
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
              "value": "{type:range,default:1000,max:3500,min:500,step:1}"
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
          "name": "sg-01",
          "connectionName": "ibm-au-syd",
          "vNetId": "vnet-01",
          "description": "Recommended security group for ec268ed7-821e-9d73-e79f-961262161624",
          "firewallRules": [
            {
              "Ports": "",
              "Protocol": "icmp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "68",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "5353",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1900",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "80",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "443",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "8080",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "9113",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9113",
              "Protocol": "udp",
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
              "Ports": "1-65535",
              "Protocol": "tcp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1-65535",
              "Protocol": "udp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            }
          ],
          "cspResourceId": ""
        },
        {
          "name": "sg-02",
          "connectionName": "ibm-au-syd",
          "vNetId": "vnet-01",
          "description": "Recommended security group for ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "firewallRules": [
            {
              "Ports": "",
              "Protocol": "icmp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "68",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "5353",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1900",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "2049",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "2049",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "111",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "111",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "20048",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "20048",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "32803",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "32803",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9100",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9100",
              "Protocol": "udp",
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
              "Ports": "1-65535",
              "Protocol": "tcp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1-65535",
              "Protocol": "udp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            }
          ],
          "cspResourceId": ""
        },
        {
          "name": "sg-03",
          "connectionName": "ibm-au-syd",
          "vNetId": "vnet-01",
          "description": "Recommended security group for ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "firewallRules": [
            {
              "Ports": "",
              "Protocol": "icmp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "68",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "5353",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1900",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "3306",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "3306",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4567",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4567",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4568",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4568",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4444",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4444",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9104",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9104",
              "Protocol": "udp",
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
              "Ports": "1-65535",
              "Protocol": "tcp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1-65535",
              "Protocol": "udp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            }
          ],
          "cspResourceId": ""
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
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=50.0% Max=100.0% Avg=94.4% | VMs: 3 total, 2 matched, 1 acceptable",
      "targetCloud": {
        "csp": "ibm",
        "region": "au-syd"
      },
      "targetInfra": {
        "name": "infra101",
        "installMonAgent": "",
        "label": null,
        "systemLabel": "",
        "description": "Recommended VMs comprising multi-cloud infrastructure",
        "nodeGroups": [
          {
            "name": "vm-ec268ed7-821e-9d73-e79f-961262161624",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+nxf-2x1",
            "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bx2-4x16",
            "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bx2-2x8",
            "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "ibm-au-syd",
        "cidrBlock": "10.0.0.0/21",
        "subnetInfoList": [
          {
            "name": "subnet-01",
            "ipv4_CIDR": "10.0.1.0/24",
            "description": "a recommended subnet for migration"
          }
        ],
        "description": "a recommended vNet for migration"
      },
      "targetSshKey": {
        "name": "sshkey-01",
        "connectionName": "ibm-au-syd",
        "description": "a SSH Key pair for migration (Note - provided ONLY once, MUST be downloaded",
        "cspResourceId": "",
        "fingerprint": "",
        "username": "",
        "verifiedUsername": "",
        "publicKey": "",
        "privateKey": ""
      },
      "targetSpecList": [
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
        },
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
          "id": "ibm+au-syd+bxf-2x8",
          "uid": "tb8moamua469rnvg885i",
          "cspSpecName": "bxf-2x8",
          "name": "ibm+au-syd+bxf-2x8",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.117,
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
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:8}"
            },
            {
              "key": "Name",
              "value": "bxf-2x8"
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
              "value": "{type:range,default:1000,max:3500,min:500,step:1}"
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
        },
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
          "id": "ibm+au-syd+bx2-2x8",
          "uid": "tbbd8blq4h4h5m6g0ad4",
          "cspSpecName": "bx2-2x8",
          "name": "ibm+au-syd+bx2-2x8",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.12,
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
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bx2-2x8"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:8}"
            },
            {
              "key": "Name",
              "value": "bx2-2x8"
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
          "name": "sg-01",
          "connectionName": "ibm-au-syd",
          "vNetId": "vnet-01",
          "description": "Recommended security group for ec268ed7-821e-9d73-e79f-961262161624",
          "firewallRules": [
            {
              "Ports": "",
              "Protocol": "icmp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "68",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "5353",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1900",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "80",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "443",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "8080",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "9113",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9113",
              "Protocol": "udp",
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
              "Ports": "1-65535",
              "Protocol": "tcp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1-65535",
              "Protocol": "udp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            }
          ],
          "cspResourceId": ""
        },
        {
          "name": "sg-02",
          "connectionName": "ibm-au-syd",
          "vNetId": "vnet-01",
          "description": "Recommended security group for ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "firewallRules": [
            {
              "Ports": "",
              "Protocol": "icmp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "68",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "5353",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1900",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "2049",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "2049",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "111",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "111",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "20048",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "20048",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "32803",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "32803",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9100",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9100",
              "Protocol": "udp",
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
              "Ports": "1-65535",
              "Protocol": "tcp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1-65535",
              "Protocol": "udp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            }
          ],
          "cspResourceId": ""
        },
        {
          "name": "sg-03",
          "connectionName": "ibm-au-syd",
          "vNetId": "vnet-01",
          "description": "Recommended security group for ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "firewallRules": [
            {
              "Ports": "",
              "Protocol": "icmp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "68",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "5353",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1900",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "22",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "3306",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "3306",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4567",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4567",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4568",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4568",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4444",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "4444",
              "Protocol": "udp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9104",
              "Protocol": "tcp",
              "Direction": "inbound",
              "CIDR": "10.0.0.0/16"
            },
            {
              "Ports": "9104",
              "Protocol": "udp",
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
              "Ports": "1-65535",
              "Protocol": "tcp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            },
            {
              "Ports": "1-65535",
              "Protocol": "udp",
              "Direction": "outbound",
              "CIDR": "0.0.0.0/0"
            }
          ],
          "cspResourceId": ""
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
  ]
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
  "description": "Candidate #1 | highly-matched | Overall Match Rate: Min=100.0% Max=100.0% Avg=100.0% | VMs: 3 total, 3 matched, 0 acceptable",
  "targetCloud": {
    "csp": "ibm",
    "region": "au-syd"
  },
  "targetInfra": {
    "name": "infra101",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "Recommended VMs comprising multi-cloud infrastructure",
    "nodeGroups": [
      {
        "name": "vm-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeGroupSize": 1,
        "label": {
          "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
        },
        "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
        "connectionName": "ibm-au-syd",
        "specId": "ibm+au-syd+nxf-2x2",
        "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-01"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 100,
        "dataDiskIds": null
      },
      {
        "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "nodeGroupSize": 1,
        "label": {
          "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
        },
        "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
        "connectionName": "ibm-au-syd",
        "specId": "ibm+au-syd+bxf-4x16",
        "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-02"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 100,
        "dataDiskIds": null
      },
      {
        "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nodeGroupSize": 1,
        "label": {
          "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
        },
        "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
        "connectionName": "ibm-au-syd",
        "specId": "ibm+au-syd+bxf-2x8",
        "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-03"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 100,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "vnet-01",
    "connectionName": "ibm-au-syd",
    "cidrBlock": "10.0.0.0/21",
    "subnetInfoList": [
      {
        "name": "subnet-01",
        "ipv4_CIDR": "10.0.1.0/24",
        "description": "a recommended subnet for migration"
      }
    ],
    "description": "a recommended vNet for migration"
  },
  "targetSshKey": {
    "name": "sshkey-01",
    "connectionName": "ibm-au-syd",
    "description": "a SSH Key pair for migration (Note - provided ONLY once, MUST be downloaded",
    "cspResourceId": "",
    "fingerprint": "",
    "username": "",
    "verifiedUsername": "",
    "publicKey": "",
    "privateKey": ""
  },
  "targetSpecList": [
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
    },
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
      "id": "ibm+au-syd+bxf-2x8",
      "uid": "tb8moamua469rnvg885i",
      "cspSpecName": "bxf-2x8",
      "name": "ibm+au-syd+bxf-2x8",
      "namespace": "system",
      "connectionName": "ibm-au-syd",
      "providerName": "ibm",
      "regionName": "au-syd",
      "regionLatitude": -33.86882,
      "regionLongitude": 151.209296,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 8,
      "diskSizeGB": -1,
      "costPerHour": 0.117,
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
          "value": "balanced"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8"
        },
        {
          "key": "Memory",
          "value": "{type:fixed,value:8}"
        },
        {
          "key": "Name",
          "value": "bxf-2x8"
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
          "value": "{type:range,default:1000,max:3500,min:500,step:1}"
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
      "name": "sg-01",
      "connectionName": "ibm-au-syd",
      "vNetId": "vnet-01",
      "description": "Recommended security group for ec268ed7-821e-9d73-e79f-961262161624",
      "firewallRules": [
        {
          "Ports": "",
          "Protocol": "icmp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "68",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "5353",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "1900",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "22",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "80",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "443",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "8080",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "9113",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "9113",
          "Protocol": "udp",
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
          "Ports": "1-65535",
          "Protocol": "tcp",
          "Direction": "outbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "1-65535",
          "Protocol": "udp",
          "Direction": "outbound",
          "CIDR": "0.0.0.0/0"
        }
      ],
      "cspResourceId": ""
    },
    {
      "name": "sg-02",
      "connectionName": "ibm-au-syd",
      "vNetId": "vnet-01",
      "description": "Recommended security group for ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "firewallRules": [
        {
          "Ports": "",
          "Protocol": "icmp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "68",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "5353",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "1900",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "22",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "2049",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "2049",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "111",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "111",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "20048",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "20048",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "32803",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "32803",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "9100",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "9100",
          "Protocol": "udp",
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
          "Ports": "1-65535",
          "Protocol": "tcp",
          "Direction": "outbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "1-65535",
          "Protocol": "udp",
          "Direction": "outbound",
          "CIDR": "0.0.0.0/0"
        }
      ],
      "cspResourceId": ""
    },
    {
      "name": "sg-03",
      "connectionName": "ibm-au-syd",
      "vNetId": "vnet-01",
      "description": "Recommended security group for ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "firewallRules": [
        {
          "Ports": "",
          "Protocol": "icmp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "68",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "5353",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "1900",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "22",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "3306",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "3306",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "4567",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "4567",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "4568",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "4568",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "4444",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "4444",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "9104",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/16"
        },
        {
          "Ports": "9104",
          "Protocol": "udp",
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
          "Ports": "1-65535",
          "Protocol": "tcp",
          "Direction": "outbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "1-65535",
          "Protocol": "udp",
          "Direction": "outbound",
          "CIDR": "0.0.0.0/0"
        }
      ],
      "cspResourceId": ""
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
- **Response**: Infrastructure migration completed successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my-infra101",
  "uid": "tbaq2rvggg9h9omg1fon",
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
    "countReconciling": 0,
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
    "sys.id": "my-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbaq2rvggg9h9omg1fon"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb4a66h2fav9qceos5f6",
      "cspResourceName": "tb4a66h2fav9qceos5f6",
      "cspResourceId": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b",
      "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-08-11 04:39:52",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-08-11 04:39:52",
        "sys.cspResourceId": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b",
        "sys.cspResourceName": "tb4a66h2fav9qceos5f6",
        "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tb4a66h2fav9qceos5f6",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.93.107",
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
      "vNetId": "my-vnet-01",
      "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
      "networkInterface": "overhear-professed-mummy-gully",
      "securityGroupIds": [
        "my-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
      "nodeUserName": "cb-user",
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
          "value": "{device:{id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4-7zxhl},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/volume_attachments/02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,name:trouble-concierge-unexpired-purging,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1fd787b5-d186-41ff-b524-8973f88d59a7,id:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,name:enlighten-paradox-freeway-cedar,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-08-11T04:39:12.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
        },
        {
          "key": "ID",
          "value": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
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
          "value": "tb4a66h2fav9qceos5f6"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_attachments/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,id:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,name:overhear-professed-mummy-gully,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_interfaces/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_attachments/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,id:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,name:overhear-professed-mummy-gully,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_interfaces/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4-7zxhl},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/volume_attachments/02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,name:trouble-concierge-unexpired-purging,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1fd787b5-d186-41ff-b524-8973f88d59a7,id:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,name:enlighten-paradox-freeway-cedar,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tbgj8u5dhooh582l8oci",
      "cspResourceName": "tbgj8u5dhooh582l8oci",
      "cspResourceId": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1",
      "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-08-11 04:39:48",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-08-11 04:39:48",
        "sys.cspResourceId": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1",
        "sys.cspResourceName": "tbgj8u5dhooh582l8oci",
        "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tbgj8u5dhooh582l8oci",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.101.96",
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
      "specId": "ibm+au-syd+bxf-2x8",
      "cspSpecName": "bxf-2x8",
      "spec": {
        "cspSpecName": "bxf-2x8",
        "vCPU": 2,
        "memoryGiB": 8,
        "costPerHour": 0.117
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
      "vNetId": "my-vnet-01",
      "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
      "networkInterface": "boatyard-boneyard-pointing-trousers",
      "securityGroupIds": [
        "my-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
      "nodeUserName": "cb-user",
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
          "value": "4000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189-dvchd},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/volume_attachments/02h7-0b796c79-8554-4334-b960-5f6a8aba2189,id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189,name:stoning-twirl-monastery-widow,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-38f72105-d482-4603-b6fc-16938806efd0,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-38f72105-d482-4603-b6fc-16938806efd0,id:r026-38f72105-d482-4603-b6fc-16938806efd0,name:undusted-grouch-crunchy-makeshift,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-08-11T04:39:12.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
        },
        {
          "key": "ID",
          "value": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
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
          "value": "8"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbgj8u5dhooh582l8oci"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_attachments/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,id:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,name:boatyard-boneyard-pointing-trousers,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_interfaces/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_attachments/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,id:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,name:boatyard-boneyard-pointing-trousers,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_interfaces/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8,name:bxf-2x8,resource_type:instance_profile}"
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
          "value": "3000"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "1000"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189-dvchd},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/volume_attachments/02h7-0b796c79-8554-4334-b960-5f6a8aba2189,id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189,name:stoning-twirl-monastery-widow,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-38f72105-d482-4603-b6fc-16938806efd0,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-38f72105-d482-4603-b6fc-16938806efd0,id:r026-38f72105-d482-4603-b6fc-16938806efd0,name:undusted-grouch-crunchy-makeshift,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tbsl1sg3d1ia78gcq2j7",
      "cspResourceName": "tbsl1sg3d1ia78gcq2j7",
      "cspResourceId": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250",
      "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-08-11 04:39:51",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-08-11 04:39:51",
        "sys.cspResourceId": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250",
        "sys.cspResourceName": "tbsl1sg3d1ia78gcq2j7",
        "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tbsl1sg3d1ia78gcq2j7",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.97.222",
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
      "vNetId": "my-vnet-01",
      "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
      "networkInterface": "dish-accuracy-astrology-company",
      "securityGroupIds": [
        "my-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
      "nodeUserName": "cb-user",
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
          "value": "{device:{id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8-r22fj},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/volume_attachments/02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,name:unweave-unwind-preppy-striving,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-ec373155-d457-4345-b065-f982da8ff830,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-ec373155-d457-4345-b065-f982da8ff830,id:r026-ec373155-d457-4345-b065-f982da8ff830,name:discolor-trunk-uncivic-sandy,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-08-11T04:39:13.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
        },
        {
          "key": "ID",
          "value": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
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
          "value": "tbsl1sg3d1ia78gcq2j7"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_attachments/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,id:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,name:dish-accuracy-astrology-company,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_interfaces/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_attachments/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,id:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,name:dish-accuracy-astrology-company,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_interfaces/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8-r22fj},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/volume_attachments/02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,name:unweave-unwind-preppy-striving,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-ec373155-d457-4345-b065-f982da8ff830,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-ec373155-d457-4345-b065-f982da8ff830,id:r026-ec373155-d457-4345-b065-f982da8ff830,name:discolor-trunk-uncivic-sandy,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
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
    "userName": "",
    "command": null
  },
  "postCommandResult": {
    "results": null
  }
}
```

</details>

### Test Case 4: Get a list of infras

#### 4.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra`
- **Purpose**: Retrieve all migrated cloud infrastructure instances
- **Namespace ID**: `mig01`
- **Request Body**: None (GET request)

#### 4.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infra list retrieved successfully

**Response Body**:

```json
{
  "infra": [
    {
      "resourceType": "infra",
      "id": "my-infra101",
      "uid": "tbaq2rvggg9h9omg1fon",
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
        "countReconciling": 0,
        "countUndefined": 0
      },
      "targetStatus": "None",
      "targetAction": "None",
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
        "sys.id": "my-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbaq2rvggg9h9omg1fon"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb4a66h2fav9qceos5f6",
          "cspResourceName": "tb4a66h2fav9qceos5f6",
          "cspResourceId": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b",
          "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-08-11 04:39:52",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-08-11 04:39:52",
            "sys.cspResourceId": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b",
            "sys.cspResourceName": "tb4a66h2fav9qceos5f6",
            "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my-subnet-01",
            "sys.uid": "tb4a66h2fav9qceos5f6",
            "sys.vNetId": "my-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.93.107",
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
          "vNetId": "my-vnet-01",
          "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
          "subnetId": "my-subnet-01",
          "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
          "networkInterface": "overhear-professed-mummy-gully",
          "securityGroupIds": [
            "my-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-sshkey-01",
          "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBNZsviSF208ANKUFzQ3IejHRQN9gWLcEhYpQeucXVCi7o8P0dRLIVgTZegubU7jseIkiWZkmjStHEvLX9zNw5w4=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:/9bAgCtV0qHBQJHTL+4DE5da7jgtuULlrcQMEA95dds",
            "firstUsedAt": "2026-08-11T04:40:59Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "xRequestId": "1786423193733877296",
              "commandRequested": "echo ready",
              "commandExecuted": "echo ready",
              "status": "Failed",
              "startedTime": "2026-08-11T04:39:53Z",
              "completedTime": "2026-08-11T04:40:07Z",
              "elapsedTime": 14,
              "resultSummary": "Command execution failed",
              "errorMessage": "failed to connect to target Node \"my-vm-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.4:22 (as \"cb-user\") via bastion Node \"my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1\" at 159.23.97.222:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 159.23.97.222:22 as user \"cb-user\" (bastionNodeId=my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
            },
            {
              "index": 2,
              "xRequestId": "1786423254880607892",
              "commandRequested": "echo ready",
              "commandExecuted": "echo ready",
              "status": "Completed",
              "startedTime": "2026-08-11T04:40:54Z",
              "completedTime": "2026-08-11T04:41:02Z",
              "elapsedTime": 8,
              "resultSummary": "Command executed successfully",
              "stdout": "ready\n\n",
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
              "value": "{device:{id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4-7zxhl},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/volume_attachments/02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,name:trouble-concierge-unexpired-purging,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1fd787b5-d186-41ff-b524-8973f88d59a7,id:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,name:enlighten-paradox-freeway-cedar,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-08-11T04:39:12.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
            },
            {
              "key": "ID",
              "value": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
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
              "value": "tb4a66h2fav9qceos5f6"
            },
            {
              "key": "NetworkAttachments",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_attachments/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,id:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,name:overhear-professed-mummy-gully,resource_type:virtual_network_interface}}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_interfaces/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkAttachment",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_attachments/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,id:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,name:overhear-professed-mummy-gully,resource_type:virtual_network_interface}}"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_interfaces/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4-7zxhl},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/volume_attachments/02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,name:trouble-concierge-unexpired-purging,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1fd787b5-d186-41ff-b524-8973f88d59a7,id:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,name:enlighten-paradox-freeway-cedar,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
            },
            {
              "key": "Zone",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tbgj8u5dhooh582l8oci",
          "cspResourceName": "tbgj8u5dhooh582l8oci",
          "cspResourceId": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1",
          "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
          "createdTime": "2026-08-11 04:39:48",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-08-11 04:39:48",
            "sys.cspResourceId": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1",
            "sys.cspResourceName": "tbgj8u5dhooh582l8oci",
            "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my-subnet-01",
            "sys.uid": "tbgj8u5dhooh582l8oci",
            "sys.vNetId": "my-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.101.96",
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
          "specId": "ibm+au-syd+bxf-2x8",
          "cspSpecName": "bxf-2x8",
          "spec": {
            "cspSpecName": "bxf-2x8",
            "vCPU": 2,
            "memoryGiB": 8,
            "costPerHour": 0.117
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
          "vNetId": "my-vnet-01",
          "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
          "subnetId": "my-subnet-01",
          "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
          "networkInterface": "boatyard-boneyard-pointing-trousers",
          "securityGroupIds": [
            "my-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-sshkey-01",
          "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBQJ6Vt5Y87q+e7YJhJaz2tzKii642TdW/r5QKuq/EHIl9yHftX84CAoy9y3g1WzWrceica+GO52uR0BhOnhG18=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:zNb+hLKgDO4fO7zNmc7WL1ZIX6bqN5E5fkPrstj1Zqk",
            "firstUsedAt": "2026-08-11T04:41:05Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "xRequestId": "1786423207619887356",
              "commandRequested": "echo ready",
              "commandExecuted": "echo ready",
              "status": "Failed",
              "startedTime": "2026-08-11T04:40:07Z",
              "completedTime": "2026-08-11T04:40:21Z",
              "elapsedTime": 14,
              "resultSummary": "Command execution failed",
              "errorMessage": "failed to connect to target Node \"my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1\" at 10.0.1.5:22 (as \"cb-user\") via bastion Node \"my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1\" at 159.23.97.222:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 159.23.97.222:22 as user \"cb-user\" (bastionNodeId=my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
            },
            {
              "index": 2,
              "xRequestId": "1786423262287557098",
              "commandRequested": "echo ready",
              "commandExecuted": "echo ready",
              "status": "Completed",
              "startedTime": "2026-08-11T04:41:02Z",
              "completedTime": "2026-08-11T04:41:14Z",
              "elapsedTime": 12,
              "resultSummary": "Command executed successfully",
              "stdout": "ready\n\n",
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
              "value": "4000"
            },
            {
              "key": "BootVolumeAttachment",
              "value": "{device:{id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189-dvchd},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/volume_attachments/02h7-0b796c79-8554-4334-b960-5f6a8aba2189,id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189,name:stoning-twirl-monastery-widow,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-38f72105-d482-4603-b6fc-16938806efd0,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-38f72105-d482-4603-b6fc-16938806efd0,id:r026-38f72105-d482-4603-b6fc-16938806efd0,name:undusted-grouch-crunchy-makeshift,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-08-11T04:39:12.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
            },
            {
              "key": "ID",
              "value": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
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
              "value": "8"
            },
            {
              "key": "MetadataService",
              "value": "{enabled:false,protocol:http,response_hop_limit:1}"
            },
            {
              "key": "Name",
              "value": "tbgj8u5dhooh582l8oci"
            },
            {
              "key": "NetworkAttachments",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_attachments/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,id:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,name:boatyard-boneyard-pointing-trousers,resource_type:virtual_network_interface}}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_interfaces/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkAttachment",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_attachments/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,id:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,name:boatyard-boneyard-pointing-trousers,resource_type:virtual_network_interface}}"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_interfaces/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
            },
            {
              "key": "Profile",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8,name:bxf-2x8,resource_type:instance_profile}"
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
              "value": "3000"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "1000"
            },
            {
              "key": "Vcpu",
              "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
            },
            {
              "key": "VolumeAttachments",
              "value": "{device:{id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189-dvchd},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/volume_attachments/02h7-0b796c79-8554-4334-b960-5f6a8aba2189,id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189,name:stoning-twirl-monastery-widow,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-38f72105-d482-4603-b6fc-16938806efd0,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-38f72105-d482-4603-b6fc-16938806efd0,id:r026-38f72105-d482-4603-b6fc-16938806efd0,name:undusted-grouch-crunchy-makeshift,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
            },
            {
              "key": "Zone",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbsl1sg3d1ia78gcq2j7",
          "cspResourceName": "tbsl1sg3d1ia78gcq2j7",
          "cspResourceId": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250",
          "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
          "createdTime": "2026-08-11 04:39:51",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-08-11 04:39:51",
            "sys.cspResourceId": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250",
            "sys.cspResourceName": "tbsl1sg3d1ia78gcq2j7",
            "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my-subnet-01",
            "sys.uid": "tbsl1sg3d1ia78gcq2j7",
            "sys.vNetId": "my-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.97.222",
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
          "vNetId": "my-vnet-01",
          "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
          "subnetId": "my-subnet-01",
          "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
          "networkInterface": "dish-accuracy-astrology-company",
          "securityGroupIds": [
            "my-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-sshkey-01",
          "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBL4tOGcSc/c+YGOKQglWRlT0kY4G8A+l7VgkhPrEuoboNkjL4GeLBkV+NKWaFygDtGciiEZeQHp4H3MwCyaxkTM=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:QWe4H2mA1W76IOOpkxyx3huu/ssmz9DIre9ZOLSCBX8",
            "firstUsedAt": "2026-08-11T04:39:55Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "xRequestId": "1786423221699134211",
              "commandRequested": "echo ready",
              "commandExecuted": "echo ready",
              "status": "Completed",
              "startedTime": "2026-08-11T04:40:21Z",
              "completedTime": "2026-08-11T04:40:45Z",
              "elapsedTime": 24,
              "resultSummary": "Command executed successfully",
              "stdout": "ready\n\n",
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
              "value": "{device:{id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8-r22fj},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/volume_attachments/02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,name:unweave-unwind-preppy-striving,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-ec373155-d457-4345-b065-f982da8ff830,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-ec373155-d457-4345-b065-f982da8ff830,id:r026-ec373155-d457-4345-b065-f982da8ff830,name:discolor-trunk-uncivic-sandy,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-08-11T04:39:13.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
            },
            {
              "key": "ID",
              "value": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
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
              "value": "tbsl1sg3d1ia78gcq2j7"
            },
            {
              "key": "NetworkAttachments",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_attachments/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,id:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,name:dish-accuracy-astrology-company,resource_type:virtual_network_interface}}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_interfaces/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkAttachment",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_attachments/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,id:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,name:dish-accuracy-astrology-company,resource_type:virtual_network_interface}}"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_interfaces/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8-r22fj},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/volume_attachments/02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,name:unweave-unwind-preppy-striving,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-ec373155-d457-4345-b065-f982da8ff830,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-ec373155-d457-4345-b065-f982da8ff830,id:r026-ec373155-d457-4345-b065-f982da8ff830,name:discolor-trunk-uncivic-sandy,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
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
        "userName": "",
        "command": null
      },
      "postCommandResult": {
        "results": null
      }
    }
  ]
}
```

### Test Case 5: Get a list of infra IDs

#### 5.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra?option=id`
- **Purpose**: Retrieve infra IDs only (lightweight response)
- **Namespace ID**: `mig01`
- **Query Parameter**: `option=id`
- **Request Body**: None (GET request)

#### 5.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infra IDs retrieved successfully

**Response Body**:

```json
{
  "idList": [
    "my-infra101"
  ]
}
```

### Test Case 6: Get a specific infra

#### 6.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Retrieve detailed information for a specific infra
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The specific infra identifier
- **Request Body**: None (GET request)

#### 6.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infra details retrieved successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my-infra101",
  "uid": "tbaq2rvggg9h9omg1fon",
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
    "countReconciling": 0,
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
    "sys.id": "my-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbaq2rvggg9h9omg1fon"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb4a66h2fav9qceos5f6",
      "cspResourceName": "tb4a66h2fav9qceos5f6",
      "cspResourceId": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b",
      "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-08-11 04:39:52",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-08-11 04:39:52",
        "sys.cspResourceId": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b",
        "sys.cspResourceName": "tb4a66h2fav9qceos5f6",
        "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tb4a66h2fav9qceos5f6",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.93.107",
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
      "vNetId": "my-vnet-01",
      "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
      "networkInterface": "overhear-professed-mummy-gully",
      "securityGroupIds": [
        "my-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBNZsviSF208ANKUFzQ3IejHRQN9gWLcEhYpQeucXVCi7o8P0dRLIVgTZegubU7jseIkiWZkmjStHEvLX9zNw5w4=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:/9bAgCtV0qHBQJHTL+4DE5da7jgtuULlrcQMEA95dds",
        "firstUsedAt": "2026-08-11T04:40:59Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "xRequestId": "1786423193733877296",
          "commandRequested": "echo ready",
          "commandExecuted": "echo ready",
          "status": "Failed",
          "startedTime": "2026-08-11T04:39:53Z",
          "completedTime": "2026-08-11T04:40:07Z",
          "elapsedTime": 14,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my-vm-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.4:22 (as \"cb-user\") via bastion Node \"my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1\" at 159.23.97.222:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 159.23.97.222:22 as user \"cb-user\" (bastionNodeId=my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "xRequestId": "1786423254880607892",
          "commandRequested": "echo ready",
          "commandExecuted": "echo ready",
          "status": "Completed",
          "startedTime": "2026-08-11T04:40:54Z",
          "completedTime": "2026-08-11T04:41:02Z",
          "elapsedTime": 8,
          "resultSummary": "Command executed successfully",
          "stdout": "ready\n\n",
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
          "value": "{device:{id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4-7zxhl},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/volume_attachments/02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,name:trouble-concierge-unexpired-purging,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1fd787b5-d186-41ff-b524-8973f88d59a7,id:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,name:enlighten-paradox-freeway-cedar,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-08-11T04:39:12.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
        },
        {
          "key": "ID",
          "value": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
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
          "value": "tb4a66h2fav9qceos5f6"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_attachments/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,id:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,name:overhear-professed-mummy-gully,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_interfaces/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_attachments/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,id:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,name:overhear-professed-mummy-gully,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_interfaces/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4-7zxhl},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/volume_attachments/02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,name:trouble-concierge-unexpired-purging,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1fd787b5-d186-41ff-b524-8973f88d59a7,id:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,name:enlighten-paradox-freeway-cedar,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tbgj8u5dhooh582l8oci",
      "cspResourceName": "tbgj8u5dhooh582l8oci",
      "cspResourceId": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1",
      "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-08-11 04:39:48",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-08-11 04:39:48",
        "sys.cspResourceId": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1",
        "sys.cspResourceName": "tbgj8u5dhooh582l8oci",
        "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tbgj8u5dhooh582l8oci",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.101.96",
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
      "specId": "ibm+au-syd+bxf-2x8",
      "cspSpecName": "bxf-2x8",
      "spec": {
        "cspSpecName": "bxf-2x8",
        "vCPU": 2,
        "memoryGiB": 8,
        "costPerHour": 0.117
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
      "vNetId": "my-vnet-01",
      "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
      "networkInterface": "boatyard-boneyard-pointing-trousers",
      "securityGroupIds": [
        "my-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBQJ6Vt5Y87q+e7YJhJaz2tzKii642TdW/r5QKuq/EHIl9yHftX84CAoy9y3g1WzWrceica+GO52uR0BhOnhG18=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:zNb+hLKgDO4fO7zNmc7WL1ZIX6bqN5E5fkPrstj1Zqk",
        "firstUsedAt": "2026-08-11T04:41:05Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "xRequestId": "1786423207619887356",
          "commandRequested": "echo ready",
          "commandExecuted": "echo ready",
          "status": "Failed",
          "startedTime": "2026-08-11T04:40:07Z",
          "completedTime": "2026-08-11T04:40:21Z",
          "elapsedTime": 14,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to connect to target Node \"my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1\" at 10.0.1.5:22 (as \"cb-user\") via bastion Node \"my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1\" at 159.23.97.222:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 159.23.97.222:22 as user \"cb-user\" (bastionNodeId=my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain"
        },
        {
          "index": 2,
          "xRequestId": "1786423262287557098",
          "commandRequested": "echo ready",
          "commandExecuted": "echo ready",
          "status": "Completed",
          "startedTime": "2026-08-11T04:41:02Z",
          "completedTime": "2026-08-11T04:41:14Z",
          "elapsedTime": 12,
          "resultSummary": "Command executed successfully",
          "stdout": "ready\n\n",
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
          "value": "4000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189-dvchd},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/volume_attachments/02h7-0b796c79-8554-4334-b960-5f6a8aba2189,id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189,name:stoning-twirl-monastery-widow,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-38f72105-d482-4603-b6fc-16938806efd0,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-38f72105-d482-4603-b6fc-16938806efd0,id:r026-38f72105-d482-4603-b6fc-16938806efd0,name:undusted-grouch-crunchy-makeshift,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-08-11T04:39:12.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
        },
        {
          "key": "ID",
          "value": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
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
          "value": "8"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbgj8u5dhooh582l8oci"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_attachments/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,id:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,name:boatyard-boneyard-pointing-trousers,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_interfaces/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_attachments/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,id:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,name:boatyard-boneyard-pointing-trousers,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_interfaces/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8,name:bxf-2x8,resource_type:instance_profile}"
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
          "value": "3000"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "1000"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189-dvchd},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/volume_attachments/02h7-0b796c79-8554-4334-b960-5f6a8aba2189,id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189,name:stoning-twirl-monastery-widow,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-38f72105-d482-4603-b6fc-16938806efd0,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-38f72105-d482-4603-b6fc-16938806efd0,id:r026-38f72105-d482-4603-b6fc-16938806efd0,name:undusted-grouch-crunchy-makeshift,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tbsl1sg3d1ia78gcq2j7",
      "cspResourceName": "tbsl1sg3d1ia78gcq2j7",
      "cspResourceId": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250",
      "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-08-11 04:39:51",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-08-11 04:39:51",
        "sys.cspResourceId": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250",
        "sys.cspResourceName": "tbsl1sg3d1ia78gcq2j7",
        "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tbsl1sg3d1ia78gcq2j7",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.97.222",
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
      "vNetId": "my-vnet-01",
      "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
      "networkInterface": "dish-accuracy-astrology-company",
      "securityGroupIds": [
        "my-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBL4tOGcSc/c+YGOKQglWRlT0kY4G8A+l7VgkhPrEuoboNkjL4GeLBkV+NKWaFygDtGciiEZeQHp4H3MwCyaxkTM=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:QWe4H2mA1W76IOOpkxyx3huu/ssmz9DIre9ZOLSCBX8",
        "firstUsedAt": "2026-08-11T04:39:55Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "xRequestId": "1786423221699134211",
          "commandRequested": "echo ready",
          "commandExecuted": "echo ready",
          "status": "Completed",
          "startedTime": "2026-08-11T04:40:21Z",
          "completedTime": "2026-08-11T04:40:45Z",
          "elapsedTime": 24,
          "resultSummary": "Command executed successfully",
          "stdout": "ready\n\n",
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
          "value": "{device:{id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8-r22fj},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/volume_attachments/02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,name:unweave-unwind-preppy-striving,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-ec373155-d457-4345-b065-f982da8ff830,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-ec373155-d457-4345-b065-f982da8ff830,id:r026-ec373155-d457-4345-b065-f982da8ff830,name:discolor-trunk-uncivic-sandy,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-08-11T04:39:13.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
        },
        {
          "key": "ID",
          "value": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
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
          "value": "tbsl1sg3d1ia78gcq2j7"
        },
        {
          "key": "NetworkAttachments",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_attachments/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,id:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,name:dish-accuracy-astrology-company,resource_type:virtual_network_interface}}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_interfaces/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkAttachment",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_attachments/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,id:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,name:dish-accuracy-astrology-company,resource_type:virtual_network_interface}}"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_interfaces/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8-r22fj},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/volume_attachments/02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,name:unweave-unwind-preppy-striving,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-ec373155-d457-4345-b065-f982da8ff830,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-ec373155-d457-4345-b065-f982da8ff830,id:r026-ec373155-d457-4345-b065-f982da8ff830,name:discolor-trunk-uncivic-sandy,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
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
    "userName": "",
    "command": null
  },
  "postCommandResult": {
    "results": null
  }
}
```

</details>

### Test Case 7: Remote Command Accessibility Check

#### 7.1 Test Information

- **Test Type**: SSH Connectivity Test for All VMs
- **Purpose**: Verify that all migrated VMs are accessible via SSH
- **Method**: Extract public IP and SSH key from MCI access info for each VM, then execute remote command
- **Command Executed**: `uname -a` (to verify system information)
- **Authentication**: SSH key-based authentication
- **Scope**: Tests all VMs across all subgroups in the MCI

#### 7.2 Test Result Information

- **Status**: ✅ **SUCCESS**
- **Result**: All VMs are accessible via SSH

**Complete Test Details**:

<details>
  <summary> <ins>Click to see detailed test results </ins> </summary>

```json
{
  "data": {
    "configureCloudAdaptiveNetwork": "",
    "description": "Recommended VMs comprising multi-cloud infrastructure",
    "id": "my-infra101",
    "installMonAgent": "",
    "label": {
      "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
      "sys.id": "my-infra101",
      "sys.labelType": "infra",
      "sys.manager": "cb-tumblebug",
      "sys.name": "my-infra101",
      "sys.namespace": "mig01",
      "sys.uid": "tbaq2rvggg9h9omg1fon"
    },
    "name": "my-infra101",
    "newNodeList": null,
    "node": [
      {
        "RootDeviceName": "Not visible in IBM",
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
            "value": "{device:{id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4-7zxhl},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/volume_attachments/02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,name:trouble-concierge-unexpired-purging,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1fd787b5-d186-41ff-b524-8973f88d59a7,id:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,name:enlighten-paradox-freeway-cedar,resource_type:volume}}"
          },
          {
            "key": "ConfidentialComputeMode",
            "value": "disabled"
          },
          {
            "key": "CreatedAt",
            "value": "2026-08-11T04:39:12.000Z"
          },
          {
            "key": "CRN",
            "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
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
            "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
          },
          {
            "key": "ID",
            "value": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b"
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
            "value": "tb4a66h2fav9qceos5f6"
          },
          {
            "key": "NetworkAttachments",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_attachments/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,id:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,name:overhear-professed-mummy-gully,resource_type:virtual_network_interface}}"
          },
          {
            "key": "NetworkInterfaces",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_interfaces/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
          },
          {
            "key": "NumaCount",
            "value": "1"
          },
          {
            "key": "PrimaryNetworkAttachment",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_attachments/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,id:02h7-8330b4af-e457-464a-8491-d1f2ff0bbab2,name:overhear-professed-mummy-gully,resource_type:virtual_network_interface}}"
          },
          {
            "key": "PrimaryNetworkInterface",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/network_interfaces/02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,id:02h7-0e3b5d19-3320-4902-b712-38bbb656bcd5,name:context-collie-pasture-copurify,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,id:02h7-5b3f7727-94d0-4812-ab79-8500d36d2210,name:sarcastic-stupor-framing-pretzel,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
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
            "value": "{device:{id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4-7zxhl},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b/volume_attachments/02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,id:02h7-8d724bfb-3712-409f-bfba-a60c710c46e4,name:trouble-concierge-unexpired-purging,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1fd787b5-d186-41ff-b524-8973f88d59a7,id:r026-1fd787b5-d186-41ff-b524-8973f88d59a7,name:enlighten-paradox-freeway-cedar,resource_type:volume}}"
          },
          {
            "key": "VolumeBandwidthQosMode",
            "value": "pooled"
          },
          {
            "key": "VPC",
            "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
          },
          {
            "key": "Zone",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
          }
        ],
        "commandStatus": [
          {
            "commandExecuted": "echo ready",
            "commandRequested": "echo ready",
            "completedTime": "2026-08-11T04:40:07Z",
            "elapsedTime": 14,
            "errorMessage": "failed to connect to target Node \"my-vm-ec268ed7-821e-9d73-e79f-961262161624-1\" at 10.0.1.4:22 (as \"cb-user\") via bastion Node \"my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1\" at 159.23.97.222:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 159.23.97.222:22 as user \"cb-user\" (bastionNodeId=my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain",
            "index": 1,
            "resultSummary": "Command execution failed",
            "startedTime": "2026-08-11T04:39:53Z",
            "status": "Failed",
            "xRequestId": "1786423193733877296"
          },
          {
            "commandExecuted": "echo ready",
            "commandRequested": "echo ready",
            "completedTime": "2026-08-11T04:41:02Z",
            "elapsedTime": 8,
            "index": 2,
            "resultSummary": "Command executed successfully",
            "startedTime": "2026-08-11T04:40:54Z",
            "status": "Completed",
            "stderr": "\n",
            "stdout": "ready\n\n",
            "xRequestId": "1786423254880607892"
          }
        ],
        "connectionConfig": {
          "configName": "ibm-au-syd",
          "credentialHolder": "admin",
          "credentialName": "ibm",
          "driverName": "ibm-driver-v1.0.so",
          "providerName": "ibm",
          "regionDetail": {
            "description": "Sydney (Australia)",
            "location": {
              "display": "Australia (Sydney)",
              "latitude": -33.86882,
              "longitude": 151.209296
            },
            "regionId": "au-syd",
            "regionName": "au-syd",
            "zones": [
              "au-syd-1",
              "au-syd-2",
              "au-syd-3"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "au-syd",
            "assignedZone": "au-syd-1"
          },
          "regionZoneInfoName": "ibm-au-syd",
          "verified": true
        },
        "connectionName": "ibm-au-syd",
        "createdTime": "2026-08-11 04:39:52",
        "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "cspResourceId": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b",
        "cspResourceName": "tb4a66h2fav9qceos5f6",
        "cspSpecName": "nxf-2x2",
        "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
        "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
        "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
        "dataDiskIds": null,
        "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
        "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "image": {
          "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "osArchitecture": "x86_64",
          "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "label": {
          "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
          "sys.connectionName": "ibm-au-syd",
          "sys.createdTime": "2026-08-11 04:39:52",
          "sys.cspResourceId": "02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b",
          "sys.cspResourceName": "tb4a66h2fav9qceos5f6",
          "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "sys.infraId": "my-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "sys.subnetId": "my-subnet-01",
          "sys.uid": "tb4a66h2fav9qceos5f6",
          "sys.vNetId": "my-vnet-01"
        },
        "location": {
          "display": "Australia (Sydney)",
          "latitude": -33.86882,
          "longitude": 151.209296
        },
        "monAgentStatus": "notInstalled",
        "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "overhear-professed-mummy-gully",
        "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.4",
        "publicDNS": "",
        "publicIP": "159.23.93.107",
        "region": {
          "region": "au-syd",
          "zone": "au-syd-1"
        },
        "resourceType": "node",
        "rootDiskSize": 100,
        "rootDiskType": "general-purpose",
        "securityGroupIds": [
          "my-sg-01"
        ],
        "spec": {
          "costPerHour": 0.094,
          "cspSpecName": "nxf-2x2",
          "memoryGiB": 2,
          "vCPU": 2
        },
        "specId": "ibm+au-syd+nxf-2x2",
        "sshHostKeyInfo": {
          "fingerprint": "SHA256:/9bAgCtV0qHBQJHTL+4DE5da7jgtuULlrcQMEA95dds",
          "firstUsedAt": "2026-08-11T04:40:59Z",
          "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBNZsviSF208ANKUFzQ3IejHRQN9gWLcEhYpQeucXVCi7o8P0dRLIVgTZegubU7jseIkiWZkmjStHEvLX9zNw5w4=",
          "keyType": "ecdsa-sha2-nistp256"
        },
        "sshKeyId": "my-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tb4a66h2fav9qceos5f6",
        "vNetId": "my-vnet-01"
      },
      {
        "RootDeviceName": "Not visible in IBM",
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
            "value": "4000"
          },
          {
            "key": "BootVolumeAttachment",
            "value": "{device:{id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189-dvchd},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/volume_attachments/02h7-0b796c79-8554-4334-b960-5f6a8aba2189,id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189,name:stoning-twirl-monastery-widow,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-38f72105-d482-4603-b6fc-16938806efd0,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-38f72105-d482-4603-b6fc-16938806efd0,id:r026-38f72105-d482-4603-b6fc-16938806efd0,name:undusted-grouch-crunchy-makeshift,resource_type:volume}}"
          },
          {
            "key": "ConfidentialComputeMode",
            "value": "disabled"
          },
          {
            "key": "CreatedAt",
            "value": "2026-08-11T04:39:12.000Z"
          },
          {
            "key": "CRN",
            "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
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
            "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
          },
          {
            "key": "ID",
            "value": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1"
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
            "value": "8"
          },
          {
            "key": "MetadataService",
            "value": "{enabled:false,protocol:http,response_hop_limit:1}"
          },
          {
            "key": "Name",
            "value": "tbgj8u5dhooh582l8oci"
          },
          {
            "key": "NetworkAttachments",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_attachments/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,id:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,name:boatyard-boneyard-pointing-trousers,resource_type:virtual_network_interface}}"
          },
          {
            "key": "NetworkInterfaces",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_interfaces/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
          },
          {
            "key": "NumaCount",
            "value": "1"
          },
          {
            "key": "PrimaryNetworkAttachment",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_attachments/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,id:02h7-ee696aa3-e5c9-46e5-b20f-53250e9966cb,name:boatyard-boneyard-pointing-trousers,resource_type:virtual_network_interface}}"
          },
          {
            "key": "PrimaryNetworkInterface",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/network_interfaces/02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,id:02h7-acc028e7-3f89-46e8-8400-47034cbba1ff,name:unlawful-dictation-tubeless-streak,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,id:02h7-00ff9987-759a-421e-9dff-ea5ae90d4282,name:divisibly-drainpipe-choosing-sarcasm,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
          },
          {
            "key": "Profile",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8,name:bxf-2x8,resource_type:instance_profile}"
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
            "value": "3000"
          },
          {
            "key": "TotalVolumeBandwidth",
            "value": "1000"
          },
          {
            "key": "Vcpu",
            "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
          },
          {
            "key": "VolumeAttachments",
            "value": "{device:{id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189-dvchd},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1/volume_attachments/02h7-0b796c79-8554-4334-b960-5f6a8aba2189,id:02h7-0b796c79-8554-4334-b960-5f6a8aba2189,name:stoning-twirl-monastery-widow,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-38f72105-d482-4603-b6fc-16938806efd0,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-38f72105-d482-4603-b6fc-16938806efd0,id:r026-38f72105-d482-4603-b6fc-16938806efd0,name:undusted-grouch-crunchy-makeshift,resource_type:volume}}"
          },
          {
            "key": "VolumeBandwidthQosMode",
            "value": "pooled"
          },
          {
            "key": "VPC",
            "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
          },
          {
            "key": "Zone",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
          }
        ],
        "commandStatus": [
          {
            "commandExecuted": "echo ready",
            "commandRequested": "echo ready",
            "completedTime": "2026-08-11T04:40:21Z",
            "elapsedTime": 14,
            "errorMessage": "failed to connect to target Node \"my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1\" at 10.0.1.5:22 (as \"cb-user\") via bastion Node \"my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1\" at 159.23.97.222:22 (as \"cb-user\") after 3 attempts: [bastion] failed to establish SSH connection to bastion 159.23.97.222:22 as user \"cb-user\" (bastionNodeId=my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1): ssh: handshake failed: ssh: unable to authenticate, attempted methods [none publickey], no supported methods remain",
            "index": 1,
            "resultSummary": "Command execution failed",
            "startedTime": "2026-08-11T04:40:07Z",
            "status": "Failed",
            "xRequestId": "1786423207619887356"
          },
          {
            "commandExecuted": "echo ready",
            "commandRequested": "echo ready",
            "completedTime": "2026-08-11T04:41:14Z",
            "elapsedTime": 12,
            "index": 2,
            "resultSummary": "Command executed successfully",
            "startedTime": "2026-08-11T04:41:02Z",
            "status": "Completed",
            "stderr": "\n",
            "stdout": "ready\n\n",
            "xRequestId": "1786423262287557098"
          }
        ],
        "connectionConfig": {
          "configName": "ibm-au-syd",
          "credentialHolder": "admin",
          "credentialName": "ibm",
          "driverName": "ibm-driver-v1.0.so",
          "providerName": "ibm",
          "regionDetail": {
            "description": "Sydney (Australia)",
            "location": {
              "display": "Australia (Sydney)",
              "latitude": -33.86882,
              "longitude": 151.209296
            },
            "regionId": "au-syd",
            "regionName": "au-syd",
            "zones": [
              "au-syd-1",
              "au-syd-2",
              "au-syd-3"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "au-syd",
            "assignedZone": "au-syd-1"
          },
          "regionZoneInfoName": "ibm-au-syd",
          "verified": true
        },
        "connectionName": "ibm-au-syd",
        "createdTime": "2026-08-11 04:39:48",
        "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "cspResourceId": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1",
        "cspResourceName": "tbgj8u5dhooh582l8oci",
        "cspSpecName": "bxf-2x8",
        "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
        "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
        "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
        "dataDiskIds": null,
        "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
        "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "image": {
          "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "osArchitecture": "x86_64",
          "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "label": {
          "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "sys.connectionName": "ibm-au-syd",
          "sys.createdTime": "2026-08-11 04:39:48",
          "sys.cspResourceId": "02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1",
          "sys.cspResourceName": "tbgj8u5dhooh582l8oci",
          "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "sys.infraId": "my-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "sys.subnetId": "my-subnet-01",
          "sys.uid": "tbgj8u5dhooh582l8oci",
          "sys.vNetId": "my-vnet-01"
        },
        "location": {
          "display": "Australia (Sydney)",
          "latitude": -33.86882,
          "longitude": 151.209296
        },
        "monAgentStatus": "notInstalled",
        "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "boatyard-boneyard-pointing-trousers",
        "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.5",
        "publicDNS": "",
        "publicIP": "159.23.101.96",
        "region": {
          "region": "au-syd",
          "zone": "au-syd-1"
        },
        "resourceType": "node",
        "rootDiskSize": 100,
        "rootDiskType": "general-purpose",
        "securityGroupIds": [
          "my-sg-03"
        ],
        "spec": {
          "costPerHour": 0.117,
          "cspSpecName": "bxf-2x8",
          "memoryGiB": 8,
          "vCPU": 2
        },
        "specId": "ibm+au-syd+bxf-2x8",
        "sshHostKeyInfo": {
          "fingerprint": "SHA256:zNb+hLKgDO4fO7zNmc7WL1ZIX6bqN5E5fkPrstj1Zqk",
          "firstUsedAt": "2026-08-11T04:41:05Z",
          "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBQJ6Vt5Y87q+e7YJhJaz2tzKii642TdW/r5QKuq/EHIl9yHftX84CAoy9y3g1WzWrceica+GO52uR0BhOnhG18=",
          "keyType": "ecdsa-sha2-nistp256"
        },
        "sshKeyId": "my-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tbgj8u5dhooh582l8oci",
        "vNetId": "my-vnet-01"
      },
      {
        "RootDeviceName": "Not visible in IBM",
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
            "value": "{device:{id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8-r22fj},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/volume_attachments/02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,name:unweave-unwind-preppy-striving,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-ec373155-d457-4345-b065-f982da8ff830,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-ec373155-d457-4345-b065-f982da8ff830,id:r026-ec373155-d457-4345-b065-f982da8ff830,name:discolor-trunk-uncivic-sandy,resource_type:volume}}"
          },
          {
            "key": "ConfidentialComputeMode",
            "value": "disabled"
          },
          {
            "key": "CreatedAt",
            "value": "2026-08-11T04:39:13.000Z"
          },
          {
            "key": "CRN",
            "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
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
            "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
          },
          {
            "key": "ID",
            "value": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250"
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
            "value": "tbsl1sg3d1ia78gcq2j7"
          },
          {
            "key": "NetworkAttachments",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_attachments/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,id:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,name:dish-accuracy-astrology-company,resource_type:virtual_network_interface}}"
          },
          {
            "key": "NetworkInterfaces",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_interfaces/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
          },
          {
            "key": "NumaCount",
            "value": "1"
          },
          {
            "key": "PrimaryNetworkAttachment",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_attachments/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:instance_network_attachment,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet},virtual_network_interface:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::virtual-network-interface:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,href:https://au-syd.iaas.cloud.ibm.com/v1/virtual_network_interfaces/02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,id:02h7-eb81082b-32b9-48df-bdf0-fba1592b9f6d,name:dish-accuracy-astrology-company,resource_type:virtual_network_interface}}"
          },
          {
            "key": "PrimaryNetworkInterface",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/network_interfaces/02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,id:02h7-ab5dedf7-76fc-467d-91c4-a2747105897b,name:derailed-recoil-reunion-manicure,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825/reserved_ips/02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,id:02h7-dfcefd1c-ac46-4c0e-808e-d01ce58c45f1,name:stumble-scarce-thirstily-purveyor,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,id:02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825,name:tbtbqrk0cmcr2i602nh6,resource_type:subnet}}"
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
            "value": "{device:{id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8-r22fj},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_28866bb8-4b30-421d-bcc4-4d7884eca250/volume_attachments/02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,id:02h7-81d60f3c-fe11-4377-b6f4-fd351f0238a8,name:unweave-unwind-preppy-striving,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-ec373155-d457-4345-b065-f982da8ff830,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-ec373155-d457-4345-b065-f982da8ff830,id:r026-ec373155-d457-4345-b065-f982da8ff830,name:discolor-trunk-uncivic-sandy,resource_type:volume}}"
          },
          {
            "key": "VolumeBandwidthQosMode",
            "value": "pooled"
          },
          {
            "key": "VPC",
            "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,id:r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06,name:tbi5p64tmkjt344t0rso,resource_type:vpc}"
          },
          {
            "key": "Zone",
            "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
          }
        ],
        "commandStatus": [
          {
            "commandExecuted": "echo ready",
            "commandRequested": "echo ready",
            "completedTime": "2026-08-11T04:40:45Z",
            "elapsedTime": 24,
            "index": 1,
            "resultSummary": "Command executed successfully",
            "startedTime": "2026-08-11T04:40:21Z",
            "status": "Completed",
            "stderr": "\n",
            "stdout": "ready\n\n",
            "xRequestId": "1786423221699134211"
          }
        ],
        "connectionConfig": {
          "configName": "ibm-au-syd",
          "credentialHolder": "admin",
          "credentialName": "ibm",
          "driverName": "ibm-driver-v1.0.so",
          "providerName": "ibm",
          "regionDetail": {
            "description": "Sydney (Australia)",
            "location": {
              "display": "Australia (Sydney)",
              "latitude": -33.86882,
              "longitude": 151.209296
            },
            "regionId": "au-syd",
            "regionName": "au-syd",
            "zones": [
              "au-syd-1",
              "au-syd-2",
              "au-syd-3"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "au-syd",
            "assignedZone": "au-syd-1"
          },
          "regionZoneInfoName": "ibm-au-syd",
          "verified": true
        },
        "connectionName": "ibm-au-syd",
        "createdTime": "2026-08-11 04:39:51",
        "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "cspResourceId": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250",
        "cspResourceName": "tbsl1sg3d1ia78gcq2j7",
        "cspSpecName": "bxf-4x16",
        "cspSshKeyId": "r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b",
        "cspSubnetId": "02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825",
        "cspVNetId": "r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06",
        "dataDiskIds": null,
        "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
        "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "image": {
          "cspImageName": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
          "osArchitecture": "x86_64",
          "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599",
        "label": {
          "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "sys.connectionName": "ibm-au-syd",
          "sys.createdTime": "2026-08-11 04:39:51",
          "sys.cspResourceId": "02h7_28866bb8-4b30-421d-bcc4-4d7884eca250",
          "sys.cspResourceName": "tbsl1sg3d1ia78gcq2j7",
          "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "sys.infraId": "my-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "sys.subnetId": "my-subnet-01",
          "sys.uid": "tbsl1sg3d1ia78gcq2j7",
          "sys.vNetId": "my-vnet-01"
        },
        "location": {
          "display": "Australia (Sydney)",
          "latitude": -33.86882,
          "longitude": 151.209296
        },
        "monAgentStatus": "notInstalled",
        "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "dish-accuracy-astrology-company",
        "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.6",
        "publicDNS": "",
        "publicIP": "159.23.97.222",
        "region": {
          "region": "au-syd",
          "zone": "au-syd-1"
        },
        "resourceType": "node",
        "rootDiskSize": 100,
        "rootDiskType": "general-purpose",
        "securityGroupIds": [
          "my-sg-02"
        ],
        "spec": {
          "costPerHour": 0.235,
          "cspSpecName": "bxf-4x16",
          "memoryGiB": 16,
          "vCPU": 4
        },
        "specId": "ibm+au-syd+bxf-4x16",
        "sshHostKeyInfo": {
          "fingerprint": "SHA256:QWe4H2mA1W76IOOpkxyx3huu/ssmz9DIre9ZOLSCBX8",
          "firstUsedAt": "2026-08-11T04:39:55Z",
          "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBL4tOGcSc/c+YGOKQglWRlT0kY4G8A+l7VgkhPrEuoboNkjL4GeLBkV+NKWaFygDtGciiEZeQHp4H3MwCyaxkTM=",
          "keyType": "ecdsa-sha2-nistp256"
        },
        "sshKeyId": "my-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tbsl1sg3d1ia78gcq2j7",
        "vNetId": "my-vnet-01"
      }
    ],
    "postCommand": {
      "command": null,
      "userName": ""
    },
    "postCommandResult": {
      "results": null
    },
    "resourceType": "infra",
    "status": "Running:3 (R:3/3)",
    "statusCount": {
      "countCreating": 0,
      "countFailed": 0,
      "countRebooting": 0,
      "countReconciling": 0,
      "countRegistering": 0,
      "countResuming": 0,
      "countRunning": 3,
      "countSuspended": 0,
      "countSuspending": 0,
      "countTerminated": 0,
      "countTerminating": 0,
      "countTotal": 3,
      "countUndefined": 0
    },
    "systemLabel": "",
    "systemMessage": null,
    "targetAction": "None",
    "targetStatus": "None",
    "uid": "tbaq2rvggg9h9omg1fon"
  },
  "success": true
}
```

</details>

### Test Case 8: Target Infrastructure Summary

#### 8.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`
- **Purpose**: Get a summary of the migrated target infrastructure in Markdown format
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier
- **Query Parameter**: `format=md`

#### 8.2 API Response Information

- **Status**: ✅ **SUCCESS**

### Test Case 9: Migration Report

#### 9.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Generate a comprehensive migration report matching source to target
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier

#### 9.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-08-11 04:41:45

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
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
| bxf-4x16 | 4 | 16.0 | - | x86_64 |  | $0.2350 | 1 |
| nxf-2x2 | 2 | 2.0 | - | x86_64 |  | $0.0940 | 1 |
| bxf-2x8 | 2 | 8.0 | - | x86_64 |  | $0.1170 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| r026-7769ffd7-a85c-45e0-a4e8-390fe5ef9599 | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | 02h7_4b5bd22c-c51d-43c3-8a4d-946e949e216b | Running | 2 vCPU, 2.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 159.23.93.107<br>**Private IP:** 10.0.1.4<br>**SGs:** my-sg-01<br>**SSH:** my-sshkey-01 |
| my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | 02h7_e7f3fefd-274f-43b8-b3f2-3094baa5b1c1 | Running | 2 vCPU, 8.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 159.23.101.96<br>**Private IP:** 10.0.1.5<br>**SGs:** my-sg-03<br>**SSH:** my-sshkey-01 |
| my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | 02h7_28866bb8-4b30-421d-bcc4-4d7884eca250 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 159.23.97.222<br>**Private IP:** 10.0.1.6<br>**SGs:** my-sg-02<br>**SSH:** my-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-vnet-01 |
| **CSP VNet ID** | r026-77fe3d8b-4a09-4841-8757-7b1a2f815f06 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-subnet-01 | 02h7-1dab2d1b-0df4-46ce-96c4-ece22b1dc825 | 10.0.1.0/24 | au-syd-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-sshkey-01 | r026-dcd9a530-8911-40bd-9d9c-4d1d2ede1e3b |  | SHA256:LoJrXJ257SVogVWm+HbONfKROwn+uNeeY0SjohTa664 |

### Security Groups

#### Security Group: my-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-sg-01 |
| **CSP Security Group ID** | r026-9bd29978-ac02-4cea-a031-bce1f4f64189 |
| **VNet** | my-vnet-01 |
| **Rule Count** | 14 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 80 | 0.0.0.0/0 |
| inbound | TCP | 443 | 0.0.0.0/0 |
| inbound | TCP | 8080 | 0.0.0.0/0 |
| inbound | TCP | 9113 | 10.0.0.0/16 |
| inbound | UDP | 9113 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-sg-02 |
| **CSP Security Group ID** | r026-e337237b-892e-485c-ab73-15920661d4f9 |
| **VNet** | my-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 2049 | 0.0.0.0/0 |
| inbound | UDP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 111 | 0.0.0.0/0 |
| inbound | UDP | 111 | 0.0.0.0/0 |
| inbound | TCP | 20048 | 10.0.0.0/16 |
| inbound | UDP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 32803 | 10.0.0.0/16 |
| inbound | UDP | 32803 | 10.0.0.0/16 |
| inbound | TCP | 9100 | 10.0.0.0/16 |
| inbound | UDP | 9100 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-sg-03

| Property | Value |
|----------|-------|
| **Name** | my-sg-03 |
| **CSP Security Group ID** | r026-9f907d23-dbfe-4ec2-937d-fa8f558ad960 |
| **VNet** | my-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ICMP |  | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | TCP | 3306 | 10.0.0.0/16 |
| inbound | UDP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 4567 | 10.0.0.0/16 |
| inbound | UDP | 4567 | 10.0.0.0/16 |
| inbound | TCP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4568 | 10.0.0.0/16 |
| inbound | TCP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 4444 | 10.0.0.0/16 |
| inbound | TCP | 9104 | 10.0.0.0/16 |
| inbound | UDP | 9104 | 10.0.0.0/16 |
| inbound | ALL |  | 10.0.0.0/16 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.4460 |
| **Per Day** | $10.70 |
| **Per Month (30 days)** | $321.12 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| IBM | au-syd | 3 | $0.4460 | $321.12 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | nxf-2x2 | $0.0940 | $67.68 |
| my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | bxf-2x8 | $0.1170 | $84.24 |
| my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | bxf-4x16 | $0.2350 | $169.20 |




### Test Case 10: Delete the migrated computing infra

#### 10.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Delete the migrated infrastructure and clean up resources
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier to delete
- **Query Parameter**: `option=terminate` (terminates all resources)
- **Request Body**: None (DELETE request)

#### 10.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infrastructure deletion completed successfully

**Response Body**:

```json
{
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my-infra101)",
  "success": true
}
```

