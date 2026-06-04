# CM-Beetle test results for IBM

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with IBM cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.5.0+ (2aeaf75)
- imdl: v0.1.6+ (2aeaf75)
- CB-Tumblebug: v0.12.13
- CB-Spider: v0.12.26
- CB-MapUI: v0.12.34
- Target CSP: IBM
- Target Region: au-syd
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: June 4, 2026
- Test Time: 15:43:50 KST
- Test Execution: 2026-06-04 15:43:50 KST

### Scenario

1. Recommend a target model for computing infra via Beetle
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
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 23.241s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 3m8.867s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 37ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 4ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 46ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.327s | Pass |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.303s | Pass |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m43.539s | Pass |

**Overall Result**: 9/9 tests passed ✅

**Total Duration**: 6m54.197963798s

*Test executed on June 4, 2026 at 15:43:50 KST (2026-06-04 15:43:50 KST) using CM-Beetle automated test CLI*

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
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
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
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
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
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
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
        "postCommand": {
          "userName": "",
          "command": null
        },
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
          "uid": "tb7jggfvq658ffi7hanf",
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
          "uid": "tbtp8i585ail9mh3jnod",
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
          "uid": "tbu5hc2ke8pnt5ssgcbo",
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
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "regionList": [
            "au-syd"
          ],
          "id": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "uid": "tbtrt2lnuipiuftrap37",
          "name": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ibm-au-syd",
          "infraType": "",
          "fetchedTime": "2026.05.27 06:32:55 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
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
      ]
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
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
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
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
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
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
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
        "postCommand": {
          "userName": "",
          "command": null
        },
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
          "uid": "tb7jggfvq658ffi7hanf",
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
          "uid": "tbtp8i585ail9mh3jnod",
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
          "uid": "tbu5hc2ke8pnt5ssgcbo",
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
          "uid": "tbm961b7uek2fdn6h3to",
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
          "uid": "tbgdlsg7fs0qf1nr53pr",
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
          "uid": "tb4036nluqjoeocg0vnh",
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
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "regionList": [
            "au-syd"
          ],
          "id": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "uid": "tbtrt2lnuipiuftrap37",
          "name": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ibm-au-syd",
          "infraType": "",
          "fetchedTime": "2026.05.27 06:32:55 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
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
      ]
    }
  ]
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
- **Response**: Infrastructure migration completed successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my06-infra101",
  "uid": "tbjcl30st67erasq8qib",
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
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
    "sys.id": "my06-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my06-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbjcl30st67erasq8qib"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbkh25vfquo719hfs9mk",
      "cspResourceName": "tbkh25vfquo719hfs9mk",
      "cspResourceId": "02h7_a8b28819-79bf-46f5-9513-850d56fe677e",
      "name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-06-04 06:47:04",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-04 06:47:04",
        "sys.cspResourceId": "02h7_a8b28819-79bf-46f5-9513-850d56fe677e",
        "sys.cspResourceName": "tbkh25vfquo719hfs9mk",
        "sys.id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tbkh25vfquo719hfs9mk",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.102.142",
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
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-58d4abe8-3239-4895-a179-9b215f0cb59c",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed",
      "networkInterface": "saddled-eloquent-appraiser-wasting",
      "securityGroupIds": [
        "my06-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-04T06:47:09Z",
          "completedTime": "2026-06-04T06:47:20Z",
          "elapsedTime": 11,
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
          "value": "{device:{id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc-rtvhz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/volume_attachments/02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,name:imperfect-stucco-undusted-paltry,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-a9f6cd56-33a2-439b-959a-653bba19cc45,id:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,name:clammy-surgical-ahead-deodorize,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-04T06:46:21.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_a8b28819-79bf-46f5-9513-850d56fe677e"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e"
        },
        {
          "key": "ID",
          "value": "02h7_a8b28819-79bf-46f5-9513-850d56fe677e"
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
          "value": "tbkh25vfquo719hfs9mk"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/network_interfaces/02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,id:02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,name:saddled-eloquent-appraiser-wasting,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,id:02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,name:myth-skating-cage-safeness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/network_interfaces/02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,id:02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,name:saddled-eloquent-appraiser-wasting,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,id:02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,name:myth-skating-cage-safeness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc-rtvhz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/volume_attachments/02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,name:imperfect-stucco-undusted-paltry,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-a9f6cd56-33a2-439b-959a-653bba19cc45,id:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,name:clammy-surgical-ahead-deodorize,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-58d4abe8-3239-4895-a179-9b215f0cb59c,id:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,name:tb7f7olgfakvdpgt6ehn,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tb38rvud7c2vr8j0dak0",
      "cspResourceName": "tb38rvud7c2vr8j0dak0",
      "cspResourceId": "02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e",
      "name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-06-04 06:47:02",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-04 06:47:02",
        "sys.cspResourceId": "02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e",
        "sys.cspResourceName": "tb38rvud7c2vr8j0dak0",
        "sys.id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tb38rvud7c2vr8j0dak0",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.98.133",
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
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-58d4abe8-3239-4895-a179-9b215f0cb59c",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed",
      "networkInterface": "sadden-endnote-stint-elixir",
      "securityGroupIds": [
        "my06-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-04T06:47:09Z",
          "completedTime": "2026-06-04T06:47:20Z",
          "elapsedTime": 11,
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
          "value": "4000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51-rpj79},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/volume_attachments/02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,name:smartly-exquisite-overgrown-ovary,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,id:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,name:juice-unpleased-overcast-nationality,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-04T06:46:22.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e"
        },
        {
          "key": "ID",
          "value": "02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e"
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
          "value": "8"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tb38rvud7c2vr8j0dak0"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/network_interfaces/02h7-8117ed2e-dcda-450b-900e-177c2ad34825,id:02h7-8117ed2e-dcda-450b-900e-177c2ad34825,name:sadden-endnote-stint-elixir,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,id:02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,name:displace-flicker-cabdriver-proven,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/network_interfaces/02h7-8117ed2e-dcda-450b-900e-177c2ad34825,id:02h7-8117ed2e-dcda-450b-900e-177c2ad34825,name:sadden-endnote-stint-elixir,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,id:02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,name:displace-flicker-cabdriver-proven,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51-rpj79},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/volume_attachments/02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,name:smartly-exquisite-overgrown-ovary,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,id:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,name:juice-unpleased-overcast-nationality,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-58d4abe8-3239-4895-a179-9b215f0cb59c,id:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,name:tb7f7olgfakvdpgt6ehn,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tb2r7jm4b2t8u5kq6pck",
      "cspResourceName": "tb2r7jm4b2t8u5kq6pck",
      "cspResourceId": "02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c",
      "name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-06-04 06:47:00",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-04 06:47:00",
        "sys.cspResourceId": "02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c",
        "sys.cspResourceName": "tb2r7jm4b2t8u5kq6pck",
        "sys.id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tb2r7jm4b2t8u5kq6pck",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.93.185",
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
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-58d4abe8-3239-4895-a179-9b215f0cb59c",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed",
      "networkInterface": "rage-armored-disloyal-lunchroom",
      "securityGroupIds": [
        "my06-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBH0zxSfYzkuJ3tu4y/hKqTeBrH1l/URoK+U8BD8CIoJ+pDi2zSUGDXz1WO86V4o2CL0oCj+OhhGyazwx2wQmJuE=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:EMqpfvrng1Hm0Q/n5bBkKoDck8pMRw/7Z3VfNbl44j0",
        "firstUsedAt": "2026-06-04T06:47:09Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-04T06:47:09Z",
          "completedTime": "2026-06-04T06:47:20Z",
          "elapsedTime": 11,
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
          "value": "{device:{id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb-7jmc8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/volume_attachments/02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,name:captivate-shorter-stock-captivate,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-af2499cb-8605-4905-9dd3-93908cc58635,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-af2499cb-8605-4905-9dd3-93908cc58635,id:r026-af2499cb-8605-4905-9dd3-93908cc58635,name:ranging-trough-finished-decode,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-04T06:46:22.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c"
        },
        {
          "key": "ID",
          "value": "02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c"
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
          "value": "tb2r7jm4b2t8u5kq6pck"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/network_interfaces/02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,id:02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,name:rage-armored-disloyal-lunchroom,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,id:02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,name:sprint-wing-upstairs-brittle,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/network_interfaces/02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,id:02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,name:rage-armored-disloyal-lunchroom,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,id:02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,name:sprint-wing-upstairs-brittle,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb-7jmc8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/volume_attachments/02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,name:captivate-shorter-stock-captivate,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-af2499cb-8605-4905-9dd3-93908cc58635,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-af2499cb-8605-4905-9dd3-93908cc58635,id:r026-af2499cb-8605-4905-9dd3-93908cc58635,name:ranging-trough-finished-decode,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-58d4abe8-3239-4895-a179-9b215f0cb59c,id:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,name:tb7f7olgfakvdpgt6ehn,resource_type:vpc}"
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
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "159.23.98.133",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "nodeIp": "159.23.93.185",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "159.23.102.142",
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
- **Purpose**: Retrieve all migrated cloud infrastructure instances
- **Namespace ID**: `mig01`
- **Request Body**: None (GET request)

#### 3.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infra list retrieved successfully

**Response Body**:

```json
{
  "infra": [
    {
      "resourceType": "infra",
      "id": "my02-infra101",
      "uid": "tb9ogd8ev5vg3eo0g473",
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
        "countUndefined": 0
      },
      "targetStatus": "None",
      "targetAction": "None",
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
        "sys.id": "my02-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tb9ogd8ev5vg3eo0g473"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbljfa7jgombrnu3e5n6",
          "cspResourceName": "tbljfa7jgombrnu3e5n6",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbljfa7jgombrnu3e5n6",
          "name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-06-04 06:46:10",
          "label": {
            "createdBy": "tbljfa7jgombrnu3e5n6",
            "keypair": "tbquel73pdc2l911ctd6",
            "publicip": "tbljfa7jgombrnu3e5n6-97352-PublicIP",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-06-04 06:46:10",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbljfa7jgombrnu3e5n6",
            "sys.cspResourceName": "tbljfa7jgombrnu3e5n6",
            "sys.id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbljfa7jgombrnu3e5n6",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "40.89.198.4",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.5",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 30,
          "RootDeviceName": "Not visible in Azure",
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
          "specId": "azure+koreasouth+standard_b2als_v2",
          "cspSpecName": "Standard_B2als_v2",
          "spec": {
            "cspSpecName": "Standard_B2als_v2",
            "vCPU": 2,
            "memoryGiB": 3.90625,
            "costPerHour": 0.0432
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202606030",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb3s7i1nefdgi56dgcqq",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb3s7i1nefdgi56dgcqq/subnets/tb3k2ch9kloatsolje7h",
          "networkInterface": "tbljfa7jgombrnu3e5n6-63454-VNic",
          "securityGroupIds": [
            "my02-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbquel73pdc2l911ctd6",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBPjJKcfrB5xvkOQD6xpbWbj8HK+sP4tcCk058nYCdJJskpt/HYfVhqCjo8dvevN18Oj7TADIrED1sBDl6h+6ths=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:Dsu82oygRAX6uykMsdZeAbJfJ4KESZiJDKKAXCmrj24",
            "firstUsedAt": "2026-06-04T06:46:21Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-04T06:46:16Z",
              "completedTime": "2026-06-04T06:46:23Z",
              "elapsedTime": 7,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbljfa7jgombrnu3e5n6 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreasouth"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbljfa7jgombrnu3e5n6-63454-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbljfa7jgombrnu3e5n6,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDFDWFA/U9Ej/p+2P1B5KOprbQemNCcsaRKqDTukdMOhKkghonIuLClGb5gWkxWWpSS91GEtTEeK9faNbxiADdLy4oPXHLE727CWzLCbyevpnWHXt/eYu9W4DhBgHQm2xIUzAtZqHsWxsNCFIZ82BNPfvsryx+aVT5rOzjPy3KU0g8R0uZUUzaVp4fcI5rYLXlEUcaVUrUBU1Av7LoB2XcCve1pxiY+mUZQEcZoaTm0ak8+X4Bo89lGYMT9nEOk4Wld2lqGCe2HjQ4brAct+m9C3MkixlF4bCfdrkL6vYZwYMDfmLt5Ppl8snNZ4ZKg60hoSpDTfRiaBW84EZs0owrGmKArE2chQFpIdm9vii3WsecdhuglK9MkNLWA++WTpgXpbDdSyRO196bBBxUTDutHcfdBOoNYcwvRvjXu7qVGdFz0U2Yh/F7p01nXklGdWM2f8duSmNpcLcr68dCppMWJdLSkUoahuME2BEXZ+HEo3ugn693q2v8zAUIks3mvRxho/imsHnEqOKb3rhBE2K0W5fDwsaUxXNGkPXRO0kNHBxxZH+0D9mFMGUMolLKpjc8Ex25rzpGMpx+vaIjEP1/mLOh5vwVLwzFiGfcf/fHahcHu54TacehPV/WLza/A8RWjHnsqhPr+93i6WJlwUCmjFdadzvqWj0LC1f/frezgnw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202606030,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202606030},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbljfa7jgombrnu3e5n6_OsDisk_1_071b42db73fd4288974bb04492164da9,storageAccountType:Premium_LRS},name:tbljfa7jgombrnu3e5n6_OsDisk_1_071b42db73fd4288974bb04492164da9,osType:Linux}},timeCreated:2026-06-04T06:45:16.6959963Z,vmId:c50fd0c4-261e-4d12-b77a-393217024b08}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbljfa7jgombrnu3e5n6,keypair:tbquel73pdc2l911ctd6,publicip:tbljfa7jgombrnu3e5n6-97352-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbljfa7jgombrnu3e5n6"
            },
            {
              "key": "Name",
              "value": "tbljfa7jgombrnu3e5n6"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tbs0d4i0m204b65angu1",
          "cspResourceName": "tbs0d4i0m204b65angu1",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbs0d4i0m204b65angu1",
          "name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-06-04 06:46:10",
          "label": {
            "createdBy": "tbs0d4i0m204b65angu1",
            "keypair": "tbquel73pdc2l911ctd6",
            "publicip": "tbs0d4i0m204b65angu1-57116-PublicIP",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-06-04 06:46:10",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbs0d4i0m204b65angu1",
            "sys.cspResourceName": "tbs0d4i0m204b65angu1",
            "sys.id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbs0d4i0m204b65angu1",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.2.21",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 30,
          "RootDeviceName": "Not visible in Azure",
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
          "specId": "azure+koreasouth+standard_b2as_v2",
          "cspSpecName": "Standard_B2as_v2",
          "spec": {
            "cspSpecName": "Standard_B2as_v2",
            "vCPU": 2,
            "memoryGiB": 7.8125,
            "costPerHour": 0.0865
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606030",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb3s7i1nefdgi56dgcqq",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb3s7i1nefdgi56dgcqq/subnets/tb3k2ch9kloatsolje7h",
          "networkInterface": "tbs0d4i0m204b65angu1-55506-VNic",
          "securityGroupIds": [
            "my02-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbquel73pdc2l911ctd6",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGzd9ZW71uFnX7xo5HyL9AaVHZ+2YjpDU0+7Iq4nmndR0lSRSOqQPq4ROka2PrvLGKlZEZG3CFYMoWCx0JtyfZg=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:qvCaanMfWZPFRNRAZzN7CfF+0VQxIsldrX8i+eLst3g",
            "firstUsedAt": "2026-06-04T06:46:16Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-04T06:46:16Z",
              "completedTime": "2026-06-04T06:46:21Z",
              "elapsedTime": 5,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbs0d4i0m204b65angu1 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreasouth"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbs0d4i0m204b65angu1-55506-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbs0d4i0m204b65angu1,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDFDWFA/U9Ej/p+2P1B5KOprbQemNCcsaRKqDTukdMOhKkghonIuLClGb5gWkxWWpSS91GEtTEeK9faNbxiADdLy4oPXHLE727CWzLCbyevpnWHXt/eYu9W4DhBgHQm2xIUzAtZqHsWxsNCFIZ82BNPfvsryx+aVT5rOzjPy3KU0g8R0uZUUzaVp4fcI5rYLXlEUcaVUrUBU1Av7LoB2XcCve1pxiY+mUZQEcZoaTm0ak8+X4Bo89lGYMT9nEOk4Wld2lqGCe2HjQ4brAct+m9C3MkixlF4bCfdrkL6vYZwYMDfmLt5Ppl8snNZ4ZKg60hoSpDTfRiaBW84EZs0owrGmKArE2chQFpIdm9vii3WsecdhuglK9MkNLWA++WTpgXpbDdSyRO196bBBxUTDutHcfdBOoNYcwvRvjXu7qVGdFz0U2Yh/F7p01nXklGdWM2f8duSmNpcLcr68dCppMWJdLSkUoahuME2BEXZ+HEo3ugn693q2v8zAUIks3mvRxho/imsHnEqOKb3rhBE2K0W5fDwsaUxXNGkPXRO0kNHBxxZH+0D9mFMGUMolLKpjc8Ex25rzpGMpx+vaIjEP1/mLOh5vwVLwzFiGfcf/fHahcHu54TacehPV/WLza/A8RWjHnsqhPr+93i6WJlwUCmjFdadzvqWj0LC1f/frezgnw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202606030,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202606030},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbs0d4i0m204b65angu1_OsDisk_1_32ffe1d1e99b45a8a6277f6e90b09234,storageAccountType:Premium_LRS},name:tbs0d4i0m204b65angu1_OsDisk_1_32ffe1d1e99b45a8a6277f6e90b09234,osType:Linux}},timeCreated:2026-06-04T06:45:15.9301804Z,vmId:4a5850e6-ca6b-4ffc-a759-856cdfebc7b8}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbs0d4i0m204b65angu1,keypair:tbquel73pdc2l911ctd6,publicip:tbs0d4i0m204b65angu1-57116-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbs0d4i0m204b65angu1"
            },
            {
              "key": "Name",
              "value": "tbs0d4i0m204b65angu1"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbu1m5ooq1dmgvvj25vg",
          "cspResourceName": "tbu1m5ooq1dmgvvj25vg",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbu1m5ooq1dmgvvj25vg",
          "name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-06-04 06:46:11",
          "label": {
            "createdBy": "tbu1m5ooq1dmgvvj25vg",
            "keypair": "tbquel73pdc2l911ctd6",
            "publicip": "tbu1m5ooq1dmgvvj25vg-31309-PublicIP",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-06-04 06:46:11",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbu1m5ooq1dmgvvj25vg",
            "sys.cspResourceName": "tbu1m5ooq1dmgvvj25vg",
            "sys.id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbu1m5ooq1dmgvvj25vg",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.153.165",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 30,
          "RootDeviceName": "Not visible in Azure",
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
          "specId": "azure+koreasouth+standard_b4as_v2",
          "cspSpecName": "Standard_B4as_v2",
          "spec": {
            "cspSpecName": "Standard_B4as_v2",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.173
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606030",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb3s7i1nefdgi56dgcqq",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb3s7i1nefdgi56dgcqq/subnets/tb3k2ch9kloatsolje7h",
          "networkInterface": "tbu1m5ooq1dmgvvj25vg-61466-VNic",
          "securityGroupIds": [
            "my02-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbquel73pdc2l911ctd6",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDSCW51BfVuSaYNBOTsXrsZXRvBVbFfOBsu3VH/mI7tUKnoM365zi31fdpUmm+iUFmWcw23//9b6fGf9IvCtKiU=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:nhnHfbKB6KxThXm+tZuJGIJjIKyB5oY2swLATCvYqAA",
            "firstUsedAt": "2026-06-04T06:46:21Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-04T06:46:16Z",
              "completedTime": "2026-06-04T06:46:24Z",
              "elapsedTime": 8,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbu1m5ooq1dmgvvj25vg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreasouth"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbu1m5ooq1dmgvvj25vg-61466-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbu1m5ooq1dmgvvj25vg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDFDWFA/U9Ej/p+2P1B5KOprbQemNCcsaRKqDTukdMOhKkghonIuLClGb5gWkxWWpSS91GEtTEeK9faNbxiADdLy4oPXHLE727CWzLCbyevpnWHXt/eYu9W4DhBgHQm2xIUzAtZqHsWxsNCFIZ82BNPfvsryx+aVT5rOzjPy3KU0g8R0uZUUzaVp4fcI5rYLXlEUcaVUrUBU1Av7LoB2XcCve1pxiY+mUZQEcZoaTm0ak8+X4Bo89lGYMT9nEOk4Wld2lqGCe2HjQ4brAct+m9C3MkixlF4bCfdrkL6vYZwYMDfmLt5Ppl8snNZ4ZKg60hoSpDTfRiaBW84EZs0owrGmKArE2chQFpIdm9vii3WsecdhuglK9MkNLWA++WTpgXpbDdSyRO196bBBxUTDutHcfdBOoNYcwvRvjXu7qVGdFz0U2Yh/F7p01nXklGdWM2f8duSmNpcLcr68dCppMWJdLSkUoahuME2BEXZ+HEo3ugn693q2v8zAUIks3mvRxho/imsHnEqOKb3rhBE2K0W5fDwsaUxXNGkPXRO0kNHBxxZH+0D9mFMGUMolLKpjc8Ex25rzpGMpx+vaIjEP1/mLOh5vwVLwzFiGfcf/fHahcHu54TacehPV/WLza/A8RWjHnsqhPr+93i6WJlwUCmjFdadzvqWj0LC1f/frezgnw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202606030,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202606030},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbu1m5ooq1dmgvvj25vg_OsDisk_1_67cc033dbedc4224abcaa67401292f41,storageAccountType:Premium_LRS},name:tbu1m5ooq1dmgvvj25vg_OsDisk_1_67cc033dbedc4224abcaa67401292f41,osType:Linux}},timeCreated:2026-06-04T06:45:17.1737878Z,vmId:785aa425-98b5-44d5-b6db-83b3a6a8d4ba}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbu1m5ooq1dmgvvj25vg,keypair:tbquel73pdc2l911ctd6,publicip:tbu1m5ooq1dmgvvj25vg-31309-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbu1m5ooq1dmgvvj25vg"
            },
            {
              "key": "Name",
              "value": "tbu1m5ooq1dmgvvj25vg"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
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
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "20.214.2.21",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbs0d4i0m204b65angu1 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "40.89.198.4",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbljfa7jgombrnu3e5n6 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "20.200.153.165",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbu1m5ooq1dmgvvj25vg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          }
        ]
      }
    },
    {
      "resourceType": "infra",
      "id": "my06-infra101",
      "uid": "tbjcl30st67erasq8qib",
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
        "countUndefined": 0
      },
      "targetStatus": "None",
      "targetAction": "None",
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
        "sys.id": "my06-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbjcl30st67erasq8qib"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbkh25vfquo719hfs9mk",
          "cspResourceName": "tbkh25vfquo719hfs9mk",
          "cspResourceId": "02h7_a8b28819-79bf-46f5-9513-850d56fe677e",
          "name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-06-04 06:47:04",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-06-04 06:47:04",
            "sys.cspResourceId": "02h7_a8b28819-79bf-46f5-9513-850d56fe677e",
            "sys.cspResourceName": "tbkh25vfquo719hfs9mk",
            "sys.id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my06-subnet-01",
            "sys.uid": "tbkh25vfquo719hfs9mk",
            "sys.vNetId": "my06-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.102.142",
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
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-vnet-01",
          "cspVNetId": "r026-58d4abe8-3239-4895-a179-9b215f0cb59c",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed",
          "networkInterface": "saddled-eloquent-appraiser-wasting",
          "securityGroupIds": [
            "my06-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": "r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0",
          "nodeUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-06-04T06:47:09Z",
              "completedTime": "2026-06-04T06:47:20Z",
              "elapsedTime": 11,
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
              "value": "{device:{id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc-rtvhz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/volume_attachments/02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,name:imperfect-stucco-undusted-paltry,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-a9f6cd56-33a2-439b-959a-653bba19cc45,id:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,name:clammy-surgical-ahead-deodorize,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-06-04T06:46:21.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_a8b28819-79bf-46f5-9513-850d56fe677e"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e"
            },
            {
              "key": "ID",
              "value": "02h7_a8b28819-79bf-46f5-9513-850d56fe677e"
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
              "value": "tbkh25vfquo719hfs9mk"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/network_interfaces/02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,id:02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,name:saddled-eloquent-appraiser-wasting,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,id:02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,name:myth-skating-cage-safeness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/network_interfaces/02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,id:02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,name:saddled-eloquent-appraiser-wasting,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,id:02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,name:myth-skating-cage-safeness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc-rtvhz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/volume_attachments/02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,name:imperfect-stucco-undusted-paltry,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-a9f6cd56-33a2-439b-959a-653bba19cc45,id:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,name:clammy-surgical-ahead-deodorize,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-58d4abe8-3239-4895-a179-9b215f0cb59c,id:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,name:tb7f7olgfakvdpgt6ehn,resource_type:vpc}"
            },
            {
              "key": "Zone",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tb38rvud7c2vr8j0dak0",
          "cspResourceName": "tb38rvud7c2vr8j0dak0",
          "cspResourceId": "02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e",
          "name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
          "createdTime": "2026-06-04 06:47:02",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-06-04 06:47:02",
            "sys.cspResourceId": "02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e",
            "sys.cspResourceName": "tb38rvud7c2vr8j0dak0",
            "sys.id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my06-subnet-01",
            "sys.uid": "tb38rvud7c2vr8j0dak0",
            "sys.vNetId": "my06-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.98.133",
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
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-vnet-01",
          "cspVNetId": "r026-58d4abe8-3239-4895-a179-9b215f0cb59c",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed",
          "networkInterface": "sadden-endnote-stint-elixir",
          "securityGroupIds": [
            "my06-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": "r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0",
          "nodeUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-06-04T06:47:09Z",
              "completedTime": "2026-06-04T06:47:20Z",
              "elapsedTime": 11,
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
              "value": "4000"
            },
            {
              "key": "BootVolumeAttachment",
              "value": "{device:{id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51-rpj79},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/volume_attachments/02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,name:smartly-exquisite-overgrown-ovary,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,id:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,name:juice-unpleased-overcast-nationality,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-06-04T06:46:22.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e"
            },
            {
              "key": "ID",
              "value": "02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e"
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
              "value": "8"
            },
            {
              "key": "MetadataService",
              "value": "{enabled:false,protocol:http,response_hop_limit:1}"
            },
            {
              "key": "Name",
              "value": "tb38rvud7c2vr8j0dak0"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/network_interfaces/02h7-8117ed2e-dcda-450b-900e-177c2ad34825,id:02h7-8117ed2e-dcda-450b-900e-177c2ad34825,name:sadden-endnote-stint-elixir,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,id:02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,name:displace-flicker-cabdriver-proven,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/network_interfaces/02h7-8117ed2e-dcda-450b-900e-177c2ad34825,id:02h7-8117ed2e-dcda-450b-900e-177c2ad34825,name:sadden-endnote-stint-elixir,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,id:02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,name:displace-flicker-cabdriver-proven,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51-rpj79},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/volume_attachments/02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,name:smartly-exquisite-overgrown-ovary,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,id:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,name:juice-unpleased-overcast-nationality,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-58d4abe8-3239-4895-a179-9b215f0cb59c,id:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,name:tb7f7olgfakvdpgt6ehn,resource_type:vpc}"
            },
            {
              "key": "Zone",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tb2r7jm4b2t8u5kq6pck",
          "cspResourceName": "tb2r7jm4b2t8u5kq6pck",
          "cspResourceId": "02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c",
          "name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
          "createdTime": "2026-06-04 06:47:00",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-06-04 06:47:00",
            "sys.cspResourceId": "02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c",
            "sys.cspResourceName": "tb2r7jm4b2t8u5kq6pck",
            "sys.id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my06-subnet-01",
            "sys.uid": "tb2r7jm4b2t8u5kq6pck",
            "sys.vNetId": "my06-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.93.185",
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
          "vNetId": "my06-vnet-01",
          "cspVNetId": "r026-58d4abe8-3239-4895-a179-9b215f0cb59c",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed",
          "networkInterface": "rage-armored-disloyal-lunchroom",
          "securityGroupIds": [
            "my06-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": "r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBH0zxSfYzkuJ3tu4y/hKqTeBrH1l/URoK+U8BD8CIoJ+pDi2zSUGDXz1WO86V4o2CL0oCj+OhhGyazwx2wQmJuE=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:EMqpfvrng1Hm0Q/n5bBkKoDck8pMRw/7Z3VfNbl44j0",
            "firstUsedAt": "2026-06-04T06:47:09Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-06-04T06:47:09Z",
              "completedTime": "2026-06-04T06:47:20Z",
              "elapsedTime": 11,
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
              "value": "{device:{id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb-7jmc8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/volume_attachments/02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,name:captivate-shorter-stock-captivate,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-af2499cb-8605-4905-9dd3-93908cc58635,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-af2499cb-8605-4905-9dd3-93908cc58635,id:r026-af2499cb-8605-4905-9dd3-93908cc58635,name:ranging-trough-finished-decode,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-06-04T06:46:22.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c"
            },
            {
              "key": "ID",
              "value": "02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c"
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
              "value": "tb2r7jm4b2t8u5kq6pck"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/network_interfaces/02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,id:02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,name:rage-armored-disloyal-lunchroom,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,id:02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,name:sprint-wing-upstairs-brittle,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/network_interfaces/02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,id:02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,name:rage-armored-disloyal-lunchroom,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,id:02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,name:sprint-wing-upstairs-brittle,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb-7jmc8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/volume_attachments/02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,name:captivate-shorter-stock-captivate,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-af2499cb-8605-4905-9dd3-93908cc58635,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-af2499cb-8605-4905-9dd3-93908cc58635,id:r026-af2499cb-8605-4905-9dd3-93908cc58635,name:ranging-trough-finished-decode,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-58d4abe8-3239-4895-a179-9b215f0cb59c,id:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,name:tb7f7olgfakvdpgt6ehn,resource_type:vpc}"
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
            "infraId": "my06-infra101",
            "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "159.23.98.133",
            "command": {
              "0": "uname -a"
            },
            "stdout": {},
            "stderr": {},
            "err": null
          },
          {
            "infraId": "my06-infra101",
            "nodeId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "159.23.93.185",
            "command": {
              "0": "uname -a"
            },
            "stdout": {},
            "stderr": {},
            "err": null
          },
          {
            "infraId": "my06-infra101",
            "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "159.23.102.142",
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

### Test Case 4: Get a list of infra IDs

#### 4.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra?option=id`
- **Purpose**: Retrieve infra IDs only (lightweight response)
- **Namespace ID**: `mig01`
- **Query Parameter**: `option=id`
- **Request Body**: None (GET request)

#### 4.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infra IDs retrieved successfully

**Response Body**:

```json
{
  "idList": [
    "my02-infra101",
    "my06-infra101"
  ]
}
```

### Test Case 5: Get a specific infra

#### 5.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Retrieve detailed information for a specific infra
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The specific infra identifier
- **Request Body**: None (GET request)

#### 5.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infra details retrieved successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my06-infra101",
  "uid": "tbjcl30st67erasq8qib",
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
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
    "sys.id": "my06-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my06-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbjcl30st67erasq8qib"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbkh25vfquo719hfs9mk",
      "cspResourceName": "tbkh25vfquo719hfs9mk",
      "cspResourceId": "02h7_a8b28819-79bf-46f5-9513-850d56fe677e",
      "name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-06-04 06:47:04",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-04 06:47:04",
        "sys.cspResourceId": "02h7_a8b28819-79bf-46f5-9513-850d56fe677e",
        "sys.cspResourceName": "tbkh25vfquo719hfs9mk",
        "sys.id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tbkh25vfquo719hfs9mk",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.102.142",
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
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-58d4abe8-3239-4895-a179-9b215f0cb59c",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed",
      "networkInterface": "saddled-eloquent-appraiser-wasting",
      "securityGroupIds": [
        "my06-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-04T06:47:09Z",
          "completedTime": "2026-06-04T06:47:20Z",
          "elapsedTime": 11,
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
          "value": "{device:{id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc-rtvhz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/volume_attachments/02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,name:imperfect-stucco-undusted-paltry,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-a9f6cd56-33a2-439b-959a-653bba19cc45,id:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,name:clammy-surgical-ahead-deodorize,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-04T06:46:21.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_a8b28819-79bf-46f5-9513-850d56fe677e"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e"
        },
        {
          "key": "ID",
          "value": "02h7_a8b28819-79bf-46f5-9513-850d56fe677e"
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
          "value": "tbkh25vfquo719hfs9mk"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/network_interfaces/02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,id:02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,name:saddled-eloquent-appraiser-wasting,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,id:02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,name:myth-skating-cage-safeness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/network_interfaces/02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,id:02h7-57d9b3a8-6678-48c1-a951-a57e336a6471,name:saddled-eloquent-appraiser-wasting,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,id:02h7-e8063f2f-87fc-4480-88b9-f67c7f248516,name:myth-skating-cage-safeness,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc-rtvhz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a8b28819-79bf-46f5-9513-850d56fe677e/volume_attachments/02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,id:02h7-c7a5d84e-0b33-4cd1-bb6d-853baaaf13fc,name:imperfect-stucco-undusted-paltry,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-a9f6cd56-33a2-439b-959a-653bba19cc45,id:r026-a9f6cd56-33a2-439b-959a-653bba19cc45,name:clammy-surgical-ahead-deodorize,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-58d4abe8-3239-4895-a179-9b215f0cb59c,id:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,name:tb7f7olgfakvdpgt6ehn,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tb38rvud7c2vr8j0dak0",
      "cspResourceName": "tb38rvud7c2vr8j0dak0",
      "cspResourceId": "02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e",
      "name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-06-04 06:47:02",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-04 06:47:02",
        "sys.cspResourceId": "02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e",
        "sys.cspResourceName": "tb38rvud7c2vr8j0dak0",
        "sys.id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tb38rvud7c2vr8j0dak0",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.98.133",
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
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-58d4abe8-3239-4895-a179-9b215f0cb59c",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed",
      "networkInterface": "sadden-endnote-stint-elixir",
      "securityGroupIds": [
        "my06-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-04T06:47:09Z",
          "completedTime": "2026-06-04T06:47:20Z",
          "elapsedTime": 11,
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
          "value": "4000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51-rpj79},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/volume_attachments/02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,name:smartly-exquisite-overgrown-ovary,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,id:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,name:juice-unpleased-overcast-nationality,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-04T06:46:22.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e"
        },
        {
          "key": "ID",
          "value": "02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e"
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
          "value": "8"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tb38rvud7c2vr8j0dak0"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/network_interfaces/02h7-8117ed2e-dcda-450b-900e-177c2ad34825,id:02h7-8117ed2e-dcda-450b-900e-177c2ad34825,name:sadden-endnote-stint-elixir,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,id:02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,name:displace-flicker-cabdriver-proven,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/network_interfaces/02h7-8117ed2e-dcda-450b-900e-177c2ad34825,id:02h7-8117ed2e-dcda-450b-900e-177c2ad34825,name:sadden-endnote-stint-elixir,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,id:02h7-6543ae57-7a15-47f1-a82a-c2f1104a4840,name:displace-flicker-cabdriver-proven,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51-rpj79},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e/volume_attachments/02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,id:02h7-94aa4121-a478-4f84-abba-2f04fa0e6d51,name:smartly-exquisite-overgrown-ovary,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,id:r026-8a4e8ef5-0b5a-4526-91ea-95a0724f03e8,name:juice-unpleased-overcast-nationality,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-58d4abe8-3239-4895-a179-9b215f0cb59c,id:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,name:tb7f7olgfakvdpgt6ehn,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tb2r7jm4b2t8u5kq6pck",
      "cspResourceName": "tb2r7jm4b2t8u5kq6pck",
      "cspResourceId": "02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c",
      "name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-06-04 06:47:00",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-04 06:47:00",
        "sys.cspResourceId": "02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c",
        "sys.cspResourceName": "tb2r7jm4b2t8u5kq6pck",
        "sys.id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tb2r7jm4b2t8u5kq6pck",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.93.185",
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
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-58d4abe8-3239-4895-a179-9b215f0cb59c",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed",
      "networkInterface": "rage-armored-disloyal-lunchroom",
      "securityGroupIds": [
        "my06-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBH0zxSfYzkuJ3tu4y/hKqTeBrH1l/URoK+U8BD8CIoJ+pDi2zSUGDXz1WO86V4o2CL0oCj+OhhGyazwx2wQmJuE=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:EMqpfvrng1Hm0Q/n5bBkKoDck8pMRw/7Z3VfNbl44j0",
        "firstUsedAt": "2026-06-04T06:47:09Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-04T06:47:09Z",
          "completedTime": "2026-06-04T06:47:20Z",
          "elapsedTime": 11,
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
          "value": "{device:{id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb-7jmc8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/volume_attachments/02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,name:captivate-shorter-stock-captivate,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-af2499cb-8605-4905-9dd3-93908cc58635,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-af2499cb-8605-4905-9dd3-93908cc58635,id:r026-af2499cb-8605-4905-9dd3-93908cc58635,name:ranging-trough-finished-decode,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-04T06:46:22.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c"
        },
        {
          "key": "ID",
          "value": "02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c"
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
          "value": "tb2r7jm4b2t8u5kq6pck"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/network_interfaces/02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,id:02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,name:rage-armored-disloyal-lunchroom,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,id:02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,name:sprint-wing-upstairs-brittle,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/network_interfaces/02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,id:02h7-5673211e-cce7-4b22-a86b-7419eb6c29ff,name:rage-armored-disloyal-lunchroom,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed/reserved_ips/02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,id:02h7-5f4e3a6f-bf70-427a-bb7a-8befe98c9c07,name:sprint-wing-upstairs-brittle,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,id:02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed,name:tbmuqa27qjtqibjhmtde,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb-7jmc8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c/volume_attachments/02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,id:02h7-f1d549c9-49e3-4915-ae3c-595e41f201bb,name:captivate-shorter-stock-captivate,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-af2499cb-8605-4905-9dd3-93908cc58635,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-af2499cb-8605-4905-9dd3-93908cc58635,id:r026-af2499cb-8605-4905-9dd3-93908cc58635,name:ranging-trough-finished-decode,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-58d4abe8-3239-4895-a179-9b215f0cb59c,id:r026-58d4abe8-3239-4895-a179-9b215f0cb59c,name:tb7f7olgfakvdpgt6ehn,resource_type:vpc}"
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
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "159.23.98.133",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "nodeIp": "159.23.93.185",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "159.23.102.142",
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
- **Purpose**: Verify that all migrated VMs are accessible via SSH
- **Method**: Extract public IP and SSH key from MCI access info for each VM, then execute remote command
- **Command Executed**: `uname -a` (to verify system information)
- **Authentication**: SSH key-based authentication
- **Scope**: Tests all VMs across all subgroups in the MCI

#### 6.2 Test Result Information

- **Status**: ✅ **SUCCESS**
- **Result**: All VMs are accessible via SSH

**Summary**: 3/3 VMs accessible via SSH

**Test Statistics**:

- Total VMs: 3
- Successful Tests: 3
- Failed Tests: 0

**Complete Test Details**:

<details>
  <summary> <ins>Click to see detailed test results </ins> </summary>

```json
{
  "failedTests": 0,
  "overallStatus": {
    "message": "3/3 VMs accessible via SSH",
    "success": true
  },
  "successfulTests": 3,
  "totalVMs": 3,
  "vmResults": [
    {
      "attempts": 1,
      "command": "uname -a",
      "nodeGroup": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "output": "Linux tbkh25vfquo719hfs9mk 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "159.23.102.142",
      "sshTest": "successful",
      "status": "success",
      "testOrder": 1,
      "userName": "cb-user"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "nodeGroup": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "output": "Linux tb38rvud7c2vr8j0dak0 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "159.23.98.133",
      "sshTest": "successful",
      "status": "success",
      "testOrder": 2,
      "userName": "cb-user"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "nodeGroup": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "nodeId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "output": "Linux tb2r7jm4b2t8u5kq6pck 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "159.23.93.185",
      "sshTest": "successful",
      "status": "success",
      "testOrder": 3,
      "userName": "cb-user"
    }
  ]
}
```

</details>

### Test Case 7: Target Infrastructure Summary

#### 7.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`
- **Purpose**: Get a summary of the migrated target infrastructure in Markdown format
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier
- **Query Parameter**: `format=md`

#### 7.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Target Infrastructure Summary**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-06-04 06:48:12

**Namespace:** mig01

**Infra Name:** my06-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my06-infra101 |
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
| bxf-2x8 | 2 | 8.0 | - | x86_64 |  | $0.1170 | 1 |
| bxf-4x16 | 4 | 16.0 | - | x86_64 |  | $0.2350 | 1 |
| nxf-2x2 | 2 | 2.0 | - | x86_64 |  | $0.0940 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| r026-c8e249d4-f148-4416-a3c6-555b7a02f67d | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | 02h7_a8b28819-79bf-46f5-9513-850d56fe677e | Running | 2 vCPU, 2.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-vnet-01<br>**Subnet:** my06-subnet-01<br>**Public IP:** 159.23.102.142<br>**Private IP:** 10.0.1.4<br>**SGs:** my06-sg-01<br>**SSH:** my06-sshkey-01 |
| my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | 02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e | Running | 2 vCPU, 8.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-vnet-01<br>**Subnet:** my06-subnet-01<br>**Public IP:** 159.23.98.133<br>**Private IP:** 10.0.1.5<br>**SGs:** my06-sg-03<br>**SSH:** my06-sshkey-01 |
| my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | 02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-vnet-01<br>**Subnet:** my06-subnet-01<br>**Public IP:** 159.23.93.185<br>**Private IP:** 10.0.1.6<br>**SGs:** my06-sg-02<br>**SSH:** my06-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my06-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my06-vnet-01 |
| **CSP VNet ID** | r026-58d4abe8-3239-4895-a179-9b215f0cb59c |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my06-subnet-01 | 02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed | 10.0.1.0/24 | au-syd-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my06-sshkey-01 | r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0 |  | SHA256:fTxm1Ki8SJ21qAriyv0yb/0v2jVppUxgujxryII5Tv0 |

### Security Groups

#### Security Group: my06-sg-01

| Property | Value |
|----------|-------|
| **Name** | my06-sg-01 |
| **CSP Security Group ID** | r026-55b8fbc9-d619-44ce-a591-f887c3ce3a76 |
| **VNet** | my06-vnet-01 |
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

#### Security Group: my06-sg-02

| Property | Value |
|----------|-------|
| **Name** | my06-sg-02 |
| **CSP Security Group ID** | r026-86b117ed-a044-49c4-b9ce-6c70e80338b5 |
| **VNet** | my06-vnet-01 |
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

#### Security Group: my06-sg-03

| Property | Value |
|----------|-------|
| **Name** | my06-sg-03 |
| **CSP Security Group ID** | r026-d404d077-e547-45dd-835c-2c14a86ec7b7 |
| **VNet** | my06-vnet-01 |
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
| my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | nxf-2x2 | $0.0940 | $67.68 |
| my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | bxf-2x8 | $0.1170 | $84.24 |
| my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | bxf-4x16 | $0.2350 | $169.20 |




### Test Case 8: Migration Report

#### 8.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Generate a comprehensive migration report matching source to target
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier

#### 8.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

*Report generated: 2026-06-04 06:48:16*

---

## 📊 Migration Summary

**Target Cloud:** IBM

**Target Region:** au-syd

**Namespace:** mig01 | **Infra ID:** my06-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $321.12 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 3 | ✅ Selected | nxf-2x2, bxf-2x8, bxf-4x16 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my06-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my06-vnet-01) |
| 6 | **Security Group** | 3 security groups | ✅ Created | Total 52 rules in 3 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** 02h7_a8b28819-79bf-46f5-9513-850d56fe677e<br>**Label(sourceMachineId):** vm-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** 02h7_b7d4d83c-1193-4ed4-b18f-43d0ce199b9e<br>**Label(sourceMachineId):** vm-ec288dd0-c6fa-8a49-2f60 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 |
| 3 | **VM Name:** my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** 02h7_a1f6e466-600d-413c-bbfa-48fbf2885a9c<br>**Label(sourceMachineId):** vm-ec2d32b5-98fb-5a96-7913 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 |

---

## ⚙️ VM Specs

**Summary:** 3 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** nxf-2x2<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2 | my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Spec ID:** bxf-2x8<br>**vCPUs:** 2<br>**Memory:** 8.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3 | my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Spec ID:** bxf-4x16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my06-sg-01

**CSP ID:** r026-55b8fbc9-d619-44ce-a591-f887c3ce3a76 | **VNet:** my06-vnet-01 | **Rules:** 14

**Assigned VMs:**

- **VM:** my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec268ed7-821e-9d73-e79f

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 2 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 3 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 4 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | inbound | TCP | 80 | 0.0.0.0/0 | inbound tcp 80 | Migrated from source |
| 7 | inbound | TCP | 443 | 0.0.0.0/0 | inbound tcp 443 | Migrated from source |
| 8 | inbound | TCP | 8080 | 0.0.0.0/0 | inbound tcp 8080 | Migrated from source |
| 9 | inbound | TCP | 9113 | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 10 | inbound | UDP | 9113 | 10.0.0.0/16 | inbound udp 9113 from 10.0.0.0/16 | Migrated from source |
| 11 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 12 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 13 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 14 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my06-sg-02

**CSP ID:** r026-86b117ed-a044-49c4-b9ce-6c70e80338b5 | **VNet:** my06-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec2d32b5-98fb-5a96-7913

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 2 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 3 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 4 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | inbound | TCP | 2049 | 0.0.0.0/0 | inbound tcp 2049 | Migrated from source |
| 7 | inbound | UDP | 2049 | 0.0.0.0/0 | inbound udp 2049 | Migrated from source |
| 8 | inbound | TCP | 111 | 0.0.0.0/0 | inbound tcp 111 | Migrated from source |
| 9 | inbound | UDP | 111 | 0.0.0.0/0 | inbound udp 111 | Migrated from source |
| 10 | inbound | TCP | 20048 | 10.0.0.0/16 | inbound tcp 20048 from 10.0.0.0/16 | Migrated from source |
| 11 | inbound | UDP | 20048 | 10.0.0.0/16 | inbound udp 20048 from 10.0.0.0/16 | Migrated from source |
| 12 | inbound | TCP | 32803 | 10.0.0.0/16 | inbound tcp 32803 from 10.0.0.0/16 | Migrated from source |
| 13 | inbound | UDP | 32803 | 10.0.0.0/16 | inbound udp 32803 from 10.0.0.0/16 | Migrated from source |
| 14 | inbound | TCP | 9100 | 10.0.0.0/16 | inbound tcp 9100 from 10.0.0.0/16 | Migrated from source |
| 15 | inbound | UDP | 9100 | 10.0.0.0/16 | inbound udp 9100 from 10.0.0.0/16 | Migrated from source |
| 16 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 17 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 18 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 19 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my06-sg-03

**CSP ID:** r026-d404d077-e547-45dd-835c-2c14a86ec7b7 | **VNet:** my06-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec288dd0-c6fa-8a49-2f60

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 2 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 3 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 4 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | inbound | TCP | 3306 | 10.0.0.0/16 | inbound tcp 3306 from 10.0.0.0/16 | Migrated from source |
| 7 | inbound | UDP | 3306 | 10.0.0.0/16 | inbound udp 3306 from 10.0.0.0/16 | Migrated from source |
| 8 | inbound | TCP | 4567 | 10.0.0.0/16 | inbound tcp 4567 from 10.0.0.0/16 | Migrated from source |
| 9 | inbound | UDP | 4567 | 10.0.0.0/16 | inbound udp 4567 from 10.0.0.0/16 | Migrated from source |
| 10 | inbound | TCP | 4568 | 10.0.0.0/16 | inbound tcp 4568 from 10.0.0.0/16 | Migrated from source |
| 11 | inbound | UDP | 4568 | 10.0.0.0/16 | inbound udp 4568 from 10.0.0.0/16 | Migrated from source |
| 12 | inbound | TCP | 4444 | 10.0.0.0/16 | inbound tcp 4444 from 10.0.0.0/16 | Migrated from source |
| 13 | inbound | UDP | 4444 | 10.0.0.0/16 | inbound udp 4444 from 10.0.0.0/16 | Migrated from source |
| 14 | inbound | TCP | 9104 | 10.0.0.0/16 | inbound tcp 9104 from 10.0.0.0/16 | Migrated from source |
| 15 | inbound | UDP | 9104 | 10.0.0.0/16 | inbound udp 9104 from 10.0.0.0/16 | Migrated from source |
| 16 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 17 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 18 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 19 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source node network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** my06-vnet-01<br>**ID:** r026-58d4abe8-3239-4895-a179-9b215f0cb59c | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my06-subnet-01<br>**ID:** 02h7-aa0d21f0-1d37-4dbf-9b09-2f500d7360ed | 10.0.1.0/24 | my06-vnet-01 |

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
| 1 | my06-sshkey-01 | r026-d843bfc0-288c-4a69-8ffe-dfd6afd504d0 | SHA256:fTxm1Ki8SJ21qAriyv0yb/0v2jVppUxgujxryII5Tv0 | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.4460 |
| Daily | $10.70 |
| Monthly | $321.12 |
| Yearly | $3853.44 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | nxf-2x2 | $67.68 | 21.1% |
| ip-10-0-1-221 (migrated) | bxf-4x16 | $169.20 | 52.7% |
| ip-10-0-1-138 (migrated) | bxf-2x8 | $84.24 | 26.2% |

---


---

*Report generated by CM-Beetle*


### Test Case 9: Delete the migrated computing infra

#### 9.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Delete the migrated infrastructure and clean up resources
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier to delete
- **Query Parameter**: `option=terminate` (terminates all resources)
- **Request Body**: None (DELETE request)

#### 9.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infrastructure deletion completed successfully

**Response Body**:

```json
{
  "message": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: my06-infra101)",
  "success": true
}
```

