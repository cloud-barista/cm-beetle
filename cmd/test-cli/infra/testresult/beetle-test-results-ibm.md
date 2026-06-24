# CM-Beetle test results for IBM

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with IBM cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.5.2+ (1c7e6cd)
- imdl: v0.1.7
- CB-Tumblebug: v0.12.19
- CB-Spider: v0.12.32
- CB-MapUI: v0.12.43
- Target CSP: IBM
- Target Region: au-syd
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: June 24, 2026
- Test Time: 18:52:36 KST
- Test Execution: 2026-06-24 18:52:36 KST

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
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 19.297s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 2m45.66s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 46ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 8ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 30ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.335s | Pass |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.306s | Pass |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m51.29s | Pass |

**Overall Result**: 9/9 tests passed ✅

**Total Duration**: 6m35.558532336s

*Test executed on June 24, 2026 at 18:52:36 KST (2026-06-24 18:52:36 KST) using CM-Beetle automated test CLI*

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
        },
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
          "id": "ibm+au-syd+bxf-2x8",
          "uid": "tbn4ocprk802pj3dhb6o",
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
        },
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
          "id": "ibm+au-syd+bxf-2x8",
          "uid": "tbn4ocprk802pj3dhb6o",
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
        },
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
          "id": "ibm+au-syd+bx2-2x8",
          "uid": "tbf8amtdh3136gc0ketg",
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
  "uid": "tb908utvko9e3icakcle",
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
    "sys.uid": "tb908utvko9e3icakcle"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbe5t9gonop9uve0oeio",
      "cspResourceName": "tbe5t9gonop9uve0oeio",
      "cspResourceId": "02h7_f8178175-8693-4c3c-bb11-296961b82139",
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
      "createdTime": "2026-06-24 09:55:29",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-24 09:55:29",
        "sys.cspResourceId": "02h7_f8178175-8693-4c3c-bb11-296961b82139",
        "sys.cspResourceName": "tbe5t9gonop9uve0oeio",
        "sys.id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tbe5t9gonop9uve0oeio",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.94.241",
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
      "cspVNetId": "r026-4650a745-f71b-4cfc-9c2f-bb503634ac18",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179",
      "networkInterface": "navigator-vantage-visibly-canned",
      "securityGroupIds": [
        "my06-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-f149af7c-f2ec-41c2-937f-f4f592fe4602",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-24T09:55:34Z",
          "completedTime": "2026-06-24T09:55:46Z",
          "elapsedTime": 12,
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
          "value": "{device:{id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb-bfwzv},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/volume_attachments/02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,name:specks-sling-brownie-tactless,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-70df89b5-e1a2-498d-8850-932a697c37ea,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-70df89b5-e1a2-498d-8850-932a697c37ea,id:r026-70df89b5-e1a2-498d-8850-932a697c37ea,name:gravel-pulp-demanding-unstitch,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-24T09:54:54.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_f8178175-8693-4c3c-bb11-296961b82139"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139"
        },
        {
          "key": "ID",
          "value": "02h7_f8178175-8693-4c3c-bb11-296961b82139"
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
          "value": "tbe5t9gonop9uve0oeio"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/network_interfaces/02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,id:02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,name:navigator-vantage-visibly-canned,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,id:02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,name:multiple-spur-estate-decline,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/network_interfaces/02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,id:02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,name:navigator-vantage-visibly-canned,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,id:02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,name:multiple-spur-estate-decline,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb-bfwzv},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/volume_attachments/02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,name:specks-sling-brownie-tactless,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-70df89b5-e1a2-498d-8850-932a697c37ea,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-70df89b5-e1a2-498d-8850-932a697c37ea,id:r026-70df89b5-e1a2-498d-8850-932a697c37ea,name:gravel-pulp-demanding-unstitch,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,id:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,name:tb7n59qcn7llmdce5u08,resource_type:vpc}"
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
      "uid": "tbdkfure8m8ec6o89pjk",
      "cspResourceName": "tbdkfure8m8ec6o89pjk",
      "cspResourceId": "02h7_faa52016-7f78-46a7-bd3e-550531ee6879",
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
      "createdTime": "2026-06-24 09:55:26",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-24 09:55:26",
        "sys.cspResourceId": "02h7_faa52016-7f78-46a7-bd3e-550531ee6879",
        "sys.cspResourceName": "tbdkfure8m8ec6o89pjk",
        "sys.id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tbdkfure8m8ec6o89pjk",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.91.18",
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
      "cspVNetId": "r026-4650a745-f71b-4cfc-9c2f-bb503634ac18",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179",
      "networkInterface": "clicker-crave-koru-fragrance",
      "securityGroupIds": [
        "my06-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-f149af7c-f2ec-41c2-937f-f4f592fe4602",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-24T09:55:34Z",
          "completedTime": "2026-06-24T09:55:46Z",
          "elapsedTime": 12,
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
          "value": "{device:{id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e-gl4pw},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/volume_attachments/02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,name:ample-phony-handyman-premiere,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-726acdf4-ffaf-478b-9423-9725f3d5caba,id:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,name:flatness-donation-eternity-class,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-24T09:54:53.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_faa52016-7f78-46a7-bd3e-550531ee6879"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879"
        },
        {
          "key": "ID",
          "value": "02h7_faa52016-7f78-46a7-bd3e-550531ee6879"
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
          "value": "tbdkfure8m8ec6o89pjk"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/network_interfaces/02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,id:02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,name:clicker-crave-koru-fragrance,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,id:02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,name:rubble-hunchback-affluence-variety,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/network_interfaces/02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,id:02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,name:clicker-crave-koru-fragrance,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,id:02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,name:rubble-hunchback-affluence-variety,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e-gl4pw},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/volume_attachments/02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,name:ample-phony-handyman-premiere,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-726acdf4-ffaf-478b-9423-9725f3d5caba,id:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,name:flatness-donation-eternity-class,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,id:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,name:tb7n59qcn7llmdce5u08,resource_type:vpc}"
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
      "uid": "tb66pqlq5hgalbjrknj8",
      "cspResourceName": "tb66pqlq5hgalbjrknj8",
      "cspResourceId": "02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58",
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
      "createdTime": "2026-06-24 09:55:23",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-24 09:55:23",
        "sys.cspResourceId": "02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58",
        "sys.cspResourceName": "tb66pqlq5hgalbjrknj8",
        "sys.id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tb66pqlq5hgalbjrknj8",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.102.203",
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
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-4650a745-f71b-4cfc-9c2f-bb503634ac18",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179",
      "networkInterface": "simmering-sketch-backsight-figure",
      "securityGroupIds": [
        "my06-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-f149af7c-f2ec-41c2-937f-f4f592fe4602",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAwlOoUOidVuZgbw1kXcc8l6eBTpMrDfyZXEetHruvdDwBgYCaFYPbPv5r9tbM+n2kwuv3eJBm6Te9j8ouzNvIE=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:p+1KX05RjVc+tYECE7aLoUL6Hje9YQ4CPo91sDRTVtU",
        "firstUsedAt": "2026-06-24T09:55:34Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-24T09:55:34Z",
          "completedTime": "2026-06-24T09:55:46Z",
          "elapsedTime": 12,
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
          "value": "{device:{id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468-srpq4},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/volume_attachments/02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,name:smelting-sulfide-culture-sandbank,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-c85719db-d67e-45fc-b3ef-b1872ce83515,id:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,name:tanned-lifting-muster-strategy,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-24T09:54:53.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58"
        },
        {
          "key": "ID",
          "value": "02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58"
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
          "value": "tb66pqlq5hgalbjrknj8"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/network_interfaces/02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,id:02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,name:simmering-sketch-backsight-figure,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,id:02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,name:quarters-grading-chapter-starship,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/network_interfaces/02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,id:02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,name:simmering-sketch-backsight-figure,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,id:02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,name:quarters-grading-chapter-starship,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468-srpq4},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/volume_attachments/02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,name:smelting-sulfide-culture-sandbank,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-c85719db-d67e-45fc-b3ef-b1872ce83515,id:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,name:tanned-lifting-muster-strategy,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,id:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,name:tb7n59qcn7llmdce5u08,resource_type:vpc}"
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
        "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "159.23.94.241",
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
        "nodeIp": "159.23.102.203",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "159.23.91.18",
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
      "uid": "tb2j37lp9dabjvq7s4ka",
      "name": "my02-infra101",
      "status": "Terminated:3 (R:0/3)",
      "statusCount": {
        "countTotal": 3,
        "countCreating": 0,
        "countRunning": 0,
        "countFailed": 0,
        "countSuspended": 0,
        "countRebooting": 0,
        "countTerminated": 3,
        "countSuspending": 0,
        "countResuming": 0,
        "countTerminating": 0,
        "countRegistering": 0,
        "countUndefined": 0
      },
      "targetStatus": "Terminated",
      "targetAction": "Terminate",
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
        "sys.id": "my02-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tb2j37lp9dabjvq7s4ka"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbla363cn90obkb6v8e1",
          "cspResourceName": "tbla363cn90obkb6v8e1",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbla363cn90obkb6v8e1",
          "name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Terminated",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "terminated VM. No action is acceptable except deletion",
          "createdTime": "2026-06-24 09:54:07",
          "label": {
            "createdBy": "tbla363cn90obkb6v8e1",
            "keypair": "tb5j2cqd1vmi1jbtvugo",
            "publicip": "tbla363cn90obkb6v8e1-3573-PublicIP",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-06-24 09:54:07",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbla363cn90obkb6v8e1",
            "sys.cspResourceName": "tbla363cn90obkb6v8e1",
            "sys.id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbla363cn90obkb6v8e1",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.32.18",
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
          "specId": "azure+koreasouth+standard_b2als_v2",
          "cspSpecName": "Standard_B2als_v2",
          "spec": {
            "cspSpecName": "Standard_B2als_v2",
            "vCPU": 2,
            "memoryGiB": 3.90625,
            "costPerHour": 0.0432
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606070",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606190",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606070",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606070"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb1nia6ir5el3hht3btk",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb1nia6ir5el3hht3btk/subnets/tb1d56qhsikvcgt8p4vp",
          "networkInterface": "tbla363cn90obkb6v8e1-95321-VNic",
          "securityGroupIds": [
            "my02-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb5j2cqd1vmi1jbtvugo",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEjszZNn7JXUGz7fLZySLNpqZgls2rC4ipGwAZbSq5hiphyAZoo0v6LjqyPC6x7qYMxg/0u7Ri2d5Bw96wSNcdA=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:TRtYY+36//b6+JbVrY3Ifw6DudCFsLmXG2ySZVfVoGw",
            "firstUsedAt": "2026-06-24T09:54:14Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-24T09:54:14Z",
              "completedTime": "2026-06-24T09:54:18Z",
              "elapsedTime": 4,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbla363cn90obkb6v8e1 6.8.0-1059-azure #65~22.04.1-Ubuntu SMP Thu May 28 16:59:19 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbla363cn90obkb6v8e1-95321-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbla363cn90obkb6v8e1,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCuE1Spr6vVfAGIzgievjaRUrZdew0/Z3S2XZaiyFDPeTGli9aUF8lCDii3FulvvUJUiw6xqYAZ6Iv4tpzm+8VwlmWl8k0eHVVrGIcm2XRLKyC+vxtLoeTZwVQ1MzoL1QQ0UxqA0DFtrpcY66bR6m4wXCg5Kc6wUchkgpl4LeSy0US7+sLhuZly12OUYL8PPRdCF3V8wAjWCG35JGVM6C97R+bzmVi5vFzW7pHuyytdi6S+N+bJ7BjLS+SSeTbP/8zmZyVNkl6G5ihMsjJNKUziR1Bt9N682UvSGhkN1QOY4No5mRH+Ux66Do28AuPx5naHaNEZINRwqkOXy4kfOjtuzQASkzV3AS5Vu0x1AiU4zO4Dj3MlGmH34jGM+kzDQNK6cuBHjSJdUg8dq7Xo3+FAAO4zkV2SrWLFryBlIpjiceBWL0lWHRGmSnbuD7r3Hv7yw2L19Eh+mhx2S44M45GIlihJvj2CK8+84FyDjm7m+Icf8IS3XzjR0VrfQtT1eQ9nVtX8QXtGNdTHAjSBdTeCFG/Q3nw2Qexko5dV6VTwCBYeAqSSq5lNqDkmz46h7HgF5LgPAbhNgaldl/4SEGR3MY1vHPqZGCFPWitfBy/ZR3JuIV8bIdlOHYDDmPgyuGcFQQ3GyACvR/3ua6NlVx1Ur8fQjNska0KXRWJKDDVajQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202606190,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202606190},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbla363cn90obkb6v8e1_OsDisk_1_b8bd4da2853743fbabf428b432ea702f,storageAccountType:Premium_LRS},name:tbla363cn90obkb6v8e1_OsDisk_1_b8bd4da2853743fbabf428b432ea702f,osType:Linux}},timeCreated:2026-06-24T09:53:17.2939699Z,vmId:29170e71-6ec7-499a-ac95-8ae15099663c}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbla363cn90obkb6v8e1,keypair:tb5j2cqd1vmi1jbtvugo,publicip:tbla363cn90obkb6v8e1-3573-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbla363cn90obkb6v8e1"
            },
            {
              "key": "Name",
              "value": "tbla363cn90obkb6v8e1"
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
          "uid": "tb97ipkml8mk89v330pj",
          "cspResourceName": "tb97ipkml8mk89v330pj",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb97ipkml8mk89v330pj",
          "name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Terminated",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "terminated VM. No action is acceptable except deletion",
          "createdTime": "2026-06-24 09:54:07",
          "label": {
            "createdBy": "tb97ipkml8mk89v330pj",
            "keypair": "tb5j2cqd1vmi1jbtvugo",
            "publicip": "tb97ipkml8mk89v330pj-94356-PublicIP",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-06-24 09:54:07",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb97ipkml8mk89v330pj",
            "sys.cspResourceName": "tb97ipkml8mk89v330pj",
            "sys.id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tb97ipkml8mk89v330pj",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.40.2",
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
          "specId": "azure+koreasouth+standard_b2as_v2",
          "cspSpecName": "Standard_B2as_v2",
          "spec": {
            "cspSpecName": "Standard_B2as_v2",
            "vCPU": 2,
            "memoryGiB": 7.8125,
            "costPerHour": 0.0865
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606070",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606190",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606070",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606070"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb1nia6ir5el3hht3btk",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb1nia6ir5el3hht3btk/subnets/tb1d56qhsikvcgt8p4vp",
          "networkInterface": "tb97ipkml8mk89v330pj-10447-VNic",
          "securityGroupIds": [
            "my02-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb5j2cqd1vmi1jbtvugo",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOd/mqkj+VZLlZ+JgHa4LDMDo1V4+JasGtLtUnZG/FuM/uLXqi6mkDGH3Lum3d8Hi7hlAfAWvIMkd0NvuBzXGzY=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:JVsePcnFu0GRcpSf90mugir5+ZYO330Sdk1WvD3iklI",
            "firstUsedAt": "2026-06-24T09:54:17Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-24T09:54:14Z",
              "completedTime": "2026-06-24T09:54:20Z",
              "elapsedTime": 6,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tb97ipkml8mk89v330pj 6.8.0-1059-azure #65~22.04.1-Ubuntu SMP Thu May 28 16:59:19 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb97ipkml8mk89v330pj-10447-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb97ipkml8mk89v330pj,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCuE1Spr6vVfAGIzgievjaRUrZdew0/Z3S2XZaiyFDPeTGli9aUF8lCDii3FulvvUJUiw6xqYAZ6Iv4tpzm+8VwlmWl8k0eHVVrGIcm2XRLKyC+vxtLoeTZwVQ1MzoL1QQ0UxqA0DFtrpcY66bR6m4wXCg5Kc6wUchkgpl4LeSy0US7+sLhuZly12OUYL8PPRdCF3V8wAjWCG35JGVM6C97R+bzmVi5vFzW7pHuyytdi6S+N+bJ7BjLS+SSeTbP/8zmZyVNkl6G5ihMsjJNKUziR1Bt9N682UvSGhkN1QOY4No5mRH+Ux66Do28AuPx5naHaNEZINRwqkOXy4kfOjtuzQASkzV3AS5Vu0x1AiU4zO4Dj3MlGmH34jGM+kzDQNK6cuBHjSJdUg8dq7Xo3+FAAO4zkV2SrWLFryBlIpjiceBWL0lWHRGmSnbuD7r3Hv7yw2L19Eh+mhx2S44M45GIlihJvj2CK8+84FyDjm7m+Icf8IS3XzjR0VrfQtT1eQ9nVtX8QXtGNdTHAjSBdTeCFG/Q3nw2Qexko5dV6VTwCBYeAqSSq5lNqDkmz46h7HgF5LgPAbhNgaldl/4SEGR3MY1vHPqZGCFPWitfBy/ZR3JuIV8bIdlOHYDDmPgyuGcFQQ3GyACvR/3ua6NlVx1Ur8fQjNska0KXRWJKDDVajQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202606190,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202606190},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb97ipkml8mk89v330pj_OsDisk_1_acedce0cae9d4d70901217e8aa96db99,storageAccountType:Premium_LRS},name:tb97ipkml8mk89v330pj_OsDisk_1_acedce0cae9d4d70901217e8aa96db99,osType:Linux}},timeCreated:2026-06-24T09:53:17.923193Z,vmId:c8694b72-db80-4487-9921-f6914f13b86a}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tb97ipkml8mk89v330pj,keypair:tb5j2cqd1vmi1jbtvugo,publicip:tb97ipkml8mk89v330pj-94356-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb97ipkml8mk89v330pj"
            },
            {
              "key": "Name",
              "value": "tb97ipkml8mk89v330pj"
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
          "uid": "tbon10b0iqmi223vf4nc",
          "cspResourceName": "tbon10b0iqmi223vf4nc",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbon10b0iqmi223vf4nc",
          "name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Terminated",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "terminated VM. No action is acceptable except deletion",
          "createdTime": "2026-06-24 09:54:08",
          "label": {
            "createdBy": "tbon10b0iqmi223vf4nc",
            "keypair": "tb5j2cqd1vmi1jbtvugo",
            "publicip": "tbon10b0iqmi223vf4nc-52517-PublicIP",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-06-24 09:54:08",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbon10b0iqmi223vf4nc",
            "sys.cspResourceName": "tbon10b0iqmi223vf4nc",
            "sys.id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbon10b0iqmi223vf4nc",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.145.167",
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606070",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606190",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606070",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606070"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb1nia6ir5el3hht3btk",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb1nia6ir5el3hht3btk/subnets/tb1d56qhsikvcgt8p4vp",
          "networkInterface": "tbon10b0iqmi223vf4nc-73114-VNic",
          "securityGroupIds": [
            "my02-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tb5j2cqd1vmi1jbtvugo",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBIJgPkur/t0ShHCO/MRadf6s/CtzLSq0m9fwL20n74ZvDxWdJ8uWOz0FuYd6xepyklHVPue3oPRKCwxEp8fUVqg=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:nxytgz6BgRMUAiqYKjgFsqsFz+zj4f9rwb09pVoRHiw",
            "firstUsedAt": "2026-06-24T09:54:17Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-24T09:54:14Z",
              "completedTime": "2026-06-24T09:54:20Z",
              "elapsedTime": 6,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbon10b0iqmi223vf4nc 6.8.0-1059-azure #65~22.04.1-Ubuntu SMP Thu May 28 16:59:19 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbon10b0iqmi223vf4nc-73114-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbon10b0iqmi223vf4nc,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCuE1Spr6vVfAGIzgievjaRUrZdew0/Z3S2XZaiyFDPeTGli9aUF8lCDii3FulvvUJUiw6xqYAZ6Iv4tpzm+8VwlmWl8k0eHVVrGIcm2XRLKyC+vxtLoeTZwVQ1MzoL1QQ0UxqA0DFtrpcY66bR6m4wXCg5Kc6wUchkgpl4LeSy0US7+sLhuZly12OUYL8PPRdCF3V8wAjWCG35JGVM6C97R+bzmVi5vFzW7pHuyytdi6S+N+bJ7BjLS+SSeTbP/8zmZyVNkl6G5ihMsjJNKUziR1Bt9N682UvSGhkN1QOY4No5mRH+Ux66Do28AuPx5naHaNEZINRwqkOXy4kfOjtuzQASkzV3AS5Vu0x1AiU4zO4Dj3MlGmH34jGM+kzDQNK6cuBHjSJdUg8dq7Xo3+FAAO4zkV2SrWLFryBlIpjiceBWL0lWHRGmSnbuD7r3Hv7yw2L19Eh+mhx2S44M45GIlihJvj2CK8+84FyDjm7m+Icf8IS3XzjR0VrfQtT1eQ9nVtX8QXtGNdTHAjSBdTeCFG/Q3nw2Qexko5dV6VTwCBYeAqSSq5lNqDkmz46h7HgF5LgPAbhNgaldl/4SEGR3MY1vHPqZGCFPWitfBy/ZR3JuIV8bIdlOHYDDmPgyuGcFQQ3GyACvR/3ua6NlVx1Ur8fQjNska0KXRWJKDDVajQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202606190,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202606190},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbon10b0iqmi223vf4nc_OsDisk_1_a8827e71f5ca41a1873bf34e9a1d9c64,storageAccountType:Premium_LRS},name:tbon10b0iqmi223vf4nc_OsDisk_1_a8827e71f5ca41a1873bf34e9a1d9c64,osType:Linux}},timeCreated:2026-06-24T09:53:19.9085858Z,vmId:20935233-304e-4c82-b99c-ed02f1a21cf6}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbon10b0iqmi223vf4nc,keypair:tb5j2cqd1vmi1jbtvugo,publicip:tbon10b0iqmi223vf4nc-52517-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbon10b0iqmi223vf4nc"
            },
            {
              "key": "Name",
              "value": "tbon10b0iqmi223vf4nc"
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
            "nodeId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "20.214.32.18",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbla363cn90obkb6v8e1 6.8.0-1059-azure #65~22.04.1-Ubuntu SMP Thu May 28 16:59:19 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "20.214.40.2",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb97ipkml8mk89v330pj 6.8.0-1059-azure #65~22.04.1-Ubuntu SMP Thu May 28 16:59:19 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "20.200.145.167",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbon10b0iqmi223vf4nc 6.8.0-1059-azure #65~22.04.1-Ubuntu SMP Thu May 28 16:59:19 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "tb908utvko9e3icakcle",
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
        "sys.uid": "tb908utvko9e3icakcle"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbe5t9gonop9uve0oeio",
          "cspResourceName": "tbe5t9gonop9uve0oeio",
          "cspResourceId": "02h7_f8178175-8693-4c3c-bb11-296961b82139",
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
          "createdTime": "2026-06-24 09:55:29",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-06-24 09:55:29",
            "sys.cspResourceId": "02h7_f8178175-8693-4c3c-bb11-296961b82139",
            "sys.cspResourceName": "tbe5t9gonop9uve0oeio",
            "sys.id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my06-subnet-01",
            "sys.uid": "tbe5t9gonop9uve0oeio",
            "sys.vNetId": "my06-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.94.241",
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
          "cspVNetId": "r026-4650a745-f71b-4cfc-9c2f-bb503634ac18",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179",
          "networkInterface": "navigator-vantage-visibly-canned",
          "securityGroupIds": [
            "my06-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": "r026-f149af7c-f2ec-41c2-937f-f4f592fe4602",
          "nodeUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-06-24T09:55:34Z",
              "completedTime": "2026-06-24T09:55:46Z",
              "elapsedTime": 12,
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
              "value": "{device:{id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb-bfwzv},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/volume_attachments/02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,name:specks-sling-brownie-tactless,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-70df89b5-e1a2-498d-8850-932a697c37ea,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-70df89b5-e1a2-498d-8850-932a697c37ea,id:r026-70df89b5-e1a2-498d-8850-932a697c37ea,name:gravel-pulp-demanding-unstitch,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-06-24T09:54:54.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_f8178175-8693-4c3c-bb11-296961b82139"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139"
            },
            {
              "key": "ID",
              "value": "02h7_f8178175-8693-4c3c-bb11-296961b82139"
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
              "value": "tbe5t9gonop9uve0oeio"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/network_interfaces/02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,id:02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,name:navigator-vantage-visibly-canned,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,id:02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,name:multiple-spur-estate-decline,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/network_interfaces/02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,id:02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,name:navigator-vantage-visibly-canned,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,id:02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,name:multiple-spur-estate-decline,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb-bfwzv},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/volume_attachments/02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,name:specks-sling-brownie-tactless,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-70df89b5-e1a2-498d-8850-932a697c37ea,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-70df89b5-e1a2-498d-8850-932a697c37ea,id:r026-70df89b5-e1a2-498d-8850-932a697c37ea,name:gravel-pulp-demanding-unstitch,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,id:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,name:tb7n59qcn7llmdce5u08,resource_type:vpc}"
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
          "uid": "tbdkfure8m8ec6o89pjk",
          "cspResourceName": "tbdkfure8m8ec6o89pjk",
          "cspResourceId": "02h7_faa52016-7f78-46a7-bd3e-550531ee6879",
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
          "createdTime": "2026-06-24 09:55:26",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-06-24 09:55:26",
            "sys.cspResourceId": "02h7_faa52016-7f78-46a7-bd3e-550531ee6879",
            "sys.cspResourceName": "tbdkfure8m8ec6o89pjk",
            "sys.id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my06-subnet-01",
            "sys.uid": "tbdkfure8m8ec6o89pjk",
            "sys.vNetId": "my06-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.91.18",
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
          "cspVNetId": "r026-4650a745-f71b-4cfc-9c2f-bb503634ac18",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179",
          "networkInterface": "clicker-crave-koru-fragrance",
          "securityGroupIds": [
            "my06-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": "r026-f149af7c-f2ec-41c2-937f-f4f592fe4602",
          "nodeUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-06-24T09:55:34Z",
              "completedTime": "2026-06-24T09:55:46Z",
              "elapsedTime": 12,
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
              "value": "{device:{id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e-gl4pw},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/volume_attachments/02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,name:ample-phony-handyman-premiere,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-726acdf4-ffaf-478b-9423-9725f3d5caba,id:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,name:flatness-donation-eternity-class,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-06-24T09:54:53.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_faa52016-7f78-46a7-bd3e-550531ee6879"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879"
            },
            {
              "key": "ID",
              "value": "02h7_faa52016-7f78-46a7-bd3e-550531ee6879"
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
              "value": "tbdkfure8m8ec6o89pjk"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/network_interfaces/02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,id:02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,name:clicker-crave-koru-fragrance,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,id:02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,name:rubble-hunchback-affluence-variety,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/network_interfaces/02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,id:02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,name:clicker-crave-koru-fragrance,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,id:02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,name:rubble-hunchback-affluence-variety,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e-gl4pw},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/volume_attachments/02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,name:ample-phony-handyman-premiere,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-726acdf4-ffaf-478b-9423-9725f3d5caba,id:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,name:flatness-donation-eternity-class,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,id:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,name:tb7n59qcn7llmdce5u08,resource_type:vpc}"
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
          "uid": "tb66pqlq5hgalbjrknj8",
          "cspResourceName": "tb66pqlq5hgalbjrknj8",
          "cspResourceId": "02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58",
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
          "createdTime": "2026-06-24 09:55:23",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-06-24 09:55:23",
            "sys.cspResourceId": "02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58",
            "sys.cspResourceName": "tb66pqlq5hgalbjrknj8",
            "sys.id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my06-subnet-01",
            "sys.uid": "tb66pqlq5hgalbjrknj8",
            "sys.vNetId": "my06-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.102.203",
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
          "vNetId": "my06-vnet-01",
          "cspVNetId": "r026-4650a745-f71b-4cfc-9c2f-bb503634ac18",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179",
          "networkInterface": "simmering-sketch-backsight-figure",
          "securityGroupIds": [
            "my06-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": "r026-f149af7c-f2ec-41c2-937f-f4f592fe4602",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAwlOoUOidVuZgbw1kXcc8l6eBTpMrDfyZXEetHruvdDwBgYCaFYPbPv5r9tbM+n2kwuv3eJBm6Te9j8ouzNvIE=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:p+1KX05RjVc+tYECE7aLoUL6Hje9YQ4CPo91sDRTVtU",
            "firstUsedAt": "2026-06-24T09:55:34Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-06-24T09:55:34Z",
              "completedTime": "2026-06-24T09:55:46Z",
              "elapsedTime": 12,
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
              "value": "{device:{id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468-srpq4},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/volume_attachments/02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,name:smelting-sulfide-culture-sandbank,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-c85719db-d67e-45fc-b3ef-b1872ce83515,id:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,name:tanned-lifting-muster-strategy,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-06-24T09:54:53.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58"
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
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58"
            },
            {
              "key": "ID",
              "value": "02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58"
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
              "value": "tb66pqlq5hgalbjrknj8"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/network_interfaces/02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,id:02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,name:simmering-sketch-backsight-figure,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,id:02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,name:quarters-grading-chapter-starship,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/network_interfaces/02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,id:02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,name:simmering-sketch-backsight-figure,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,id:02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,name:quarters-grading-chapter-starship,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
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
              "value": "{device:{id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468-srpq4},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/volume_attachments/02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,name:smelting-sulfide-culture-sandbank,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-c85719db-d67e-45fc-b3ef-b1872ce83515,id:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,name:tanned-lifting-muster-strategy,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,id:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,name:tb7n59qcn7llmdce5u08,resource_type:vpc}"
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
            "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "159.23.94.241",
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
            "nodeIp": "159.23.102.203",
            "command": {
              "0": "uname -a"
            },
            "stdout": {},
            "stderr": {},
            "err": null
          },
          {
            "infraId": "my06-infra101",
            "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "159.23.91.18",
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
  "uid": "tb908utvko9e3icakcle",
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
    "sys.uid": "tb908utvko9e3icakcle"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbe5t9gonop9uve0oeio",
      "cspResourceName": "tbe5t9gonop9uve0oeio",
      "cspResourceId": "02h7_f8178175-8693-4c3c-bb11-296961b82139",
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
      "createdTime": "2026-06-24 09:55:29",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-24 09:55:29",
        "sys.cspResourceId": "02h7_f8178175-8693-4c3c-bb11-296961b82139",
        "sys.cspResourceName": "tbe5t9gonop9uve0oeio",
        "sys.id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tbe5t9gonop9uve0oeio",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.94.241",
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
      "cspVNetId": "r026-4650a745-f71b-4cfc-9c2f-bb503634ac18",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179",
      "networkInterface": "navigator-vantage-visibly-canned",
      "securityGroupIds": [
        "my06-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-f149af7c-f2ec-41c2-937f-f4f592fe4602",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-24T09:55:34Z",
          "completedTime": "2026-06-24T09:55:46Z",
          "elapsedTime": 12,
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
          "value": "{device:{id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb-bfwzv},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/volume_attachments/02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,name:specks-sling-brownie-tactless,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-70df89b5-e1a2-498d-8850-932a697c37ea,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-70df89b5-e1a2-498d-8850-932a697c37ea,id:r026-70df89b5-e1a2-498d-8850-932a697c37ea,name:gravel-pulp-demanding-unstitch,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-24T09:54:54.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_f8178175-8693-4c3c-bb11-296961b82139"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139"
        },
        {
          "key": "ID",
          "value": "02h7_f8178175-8693-4c3c-bb11-296961b82139"
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
          "value": "tbe5t9gonop9uve0oeio"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/network_interfaces/02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,id:02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,name:navigator-vantage-visibly-canned,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,id:02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,name:multiple-spur-estate-decline,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/network_interfaces/02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,id:02h7-3156ec48-9271-4221-a4e9-dfa6f6af246e,name:navigator-vantage-visibly-canned,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,id:02h7-9be166ec-be71-4fe4-95a2-919bf46932ee,name:multiple-spur-estate-decline,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb-bfwzv},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_f8178175-8693-4c3c-bb11-296961b82139/volume_attachments/02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,id:02h7-8d6b5e6b-af4a-4784-9e87-dbc0df4630fb,name:specks-sling-brownie-tactless,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-70df89b5-e1a2-498d-8850-932a697c37ea,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-70df89b5-e1a2-498d-8850-932a697c37ea,id:r026-70df89b5-e1a2-498d-8850-932a697c37ea,name:gravel-pulp-demanding-unstitch,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,id:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,name:tb7n59qcn7llmdce5u08,resource_type:vpc}"
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
      "uid": "tbdkfure8m8ec6o89pjk",
      "cspResourceName": "tbdkfure8m8ec6o89pjk",
      "cspResourceId": "02h7_faa52016-7f78-46a7-bd3e-550531ee6879",
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
      "createdTime": "2026-06-24 09:55:26",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-24 09:55:26",
        "sys.cspResourceId": "02h7_faa52016-7f78-46a7-bd3e-550531ee6879",
        "sys.cspResourceName": "tbdkfure8m8ec6o89pjk",
        "sys.id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tbdkfure8m8ec6o89pjk",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.91.18",
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
      "cspVNetId": "r026-4650a745-f71b-4cfc-9c2f-bb503634ac18",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179",
      "networkInterface": "clicker-crave-koru-fragrance",
      "securityGroupIds": [
        "my06-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-f149af7c-f2ec-41c2-937f-f4f592fe4602",
      "nodeUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-24T09:55:34Z",
          "completedTime": "2026-06-24T09:55:46Z",
          "elapsedTime": 12,
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
          "value": "{device:{id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e-gl4pw},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/volume_attachments/02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,name:ample-phony-handyman-premiere,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-726acdf4-ffaf-478b-9423-9725f3d5caba,id:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,name:flatness-donation-eternity-class,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-24T09:54:53.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_faa52016-7f78-46a7-bd3e-550531ee6879"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879"
        },
        {
          "key": "ID",
          "value": "02h7_faa52016-7f78-46a7-bd3e-550531ee6879"
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
          "value": "tbdkfure8m8ec6o89pjk"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/network_interfaces/02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,id:02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,name:clicker-crave-koru-fragrance,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,id:02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,name:rubble-hunchback-affluence-variety,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/network_interfaces/02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,id:02h7-e9d1ee5c-ac08-4d71-9e5b-fcbd769a5175,name:clicker-crave-koru-fragrance,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,id:02h7-ec803ace-2b1d-4dee-80e7-009474cdcdc8,name:rubble-hunchback-affluence-variety,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e-gl4pw},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_faa52016-7f78-46a7-bd3e-550531ee6879/volume_attachments/02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,id:02h7-0a0a26f0-d5fb-4f8f-8cfb-e474b86a290e,name:ample-phony-handyman-premiere,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-726acdf4-ffaf-478b-9423-9725f3d5caba,id:r026-726acdf4-ffaf-478b-9423-9725f3d5caba,name:flatness-donation-eternity-class,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,id:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,name:tb7n59qcn7llmdce5u08,resource_type:vpc}"
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
      "uid": "tb66pqlq5hgalbjrknj8",
      "cspResourceName": "tb66pqlq5hgalbjrknj8",
      "cspResourceId": "02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58",
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
      "createdTime": "2026-06-24 09:55:23",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-24 09:55:23",
        "sys.cspResourceId": "02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58",
        "sys.cspResourceName": "tb66pqlq5hgalbjrknj8",
        "sys.id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tb66pqlq5hgalbjrknj8",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.102.203",
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
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-4650a745-f71b-4cfc-9c2f-bb503634ac18",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179",
      "networkInterface": "simmering-sketch-backsight-figure",
      "securityGroupIds": [
        "my06-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-f149af7c-f2ec-41c2-937f-f4f592fe4602",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAwlOoUOidVuZgbw1kXcc8l6eBTpMrDfyZXEetHruvdDwBgYCaFYPbPv5r9tbM+n2kwuv3eJBm6Te9j8ouzNvIE=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:p+1KX05RjVc+tYECE7aLoUL6Hje9YQ4CPo91sDRTVtU",
        "firstUsedAt": "2026-06-24T09:55:34Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-24T09:55:34Z",
          "completedTime": "2026-06-24T09:55:46Z",
          "elapsedTime": 12,
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
          "value": "{device:{id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468-srpq4},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/volume_attachments/02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,name:smelting-sulfide-culture-sandbank,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-c85719db-d67e-45fc-b3ef-b1872ce83515,id:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,name:tanned-lifting-muster-strategy,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-24T09:54:53.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58"
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
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58"
        },
        {
          "key": "ID",
          "value": "02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58"
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
          "value": "tb66pqlq5hgalbjrknj8"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/network_interfaces/02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,id:02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,name:simmering-sketch-backsight-figure,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,id:02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,name:quarters-grading-chapter-starship,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/network_interfaces/02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,id:02h7-0d732162-5a2b-4ef9-8f9c-67d5cc075b06,name:simmering-sketch-backsight-figure,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179/reserved_ips/02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,id:02h7-58b6aa83-9ef1-4ad0-96b5-66c9e94deee8,name:quarters-grading-chapter-starship,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,id:02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179,name:tba5or5tgsmm97udbla0,resource_type:subnet}}"
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
          "value": "{device:{id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468-srpq4},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58/volume_attachments/02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,id:02h7-34fbd0a7-d91d-4129-8b81-cf92c6118468,name:smelting-sulfide-culture-sandbank,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-c85719db-d67e-45fc-b3ef-b1872ce83515,id:r026-c85719db-d67e-45fc-b3ef-b1872ce83515,name:tanned-lifting-muster-strategy,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,id:r026-4650a745-f71b-4cfc-9c2f-bb503634ac18,name:tb7n59qcn7llmdce5u08,resource_type:vpc}"
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
        "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "159.23.94.241",
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
        "nodeIp": "159.23.102.203",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "159.23.91.18",
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
      "output": "Linux tbe5t9gonop9uve0oeio 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "159.23.94.241",
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
      "output": "Linux tbdkfure8m8ec6o89pjk 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "159.23.91.18",
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
      "output": "Linux tb66pqlq5hgalbjrknj8 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "159.23.102.203",
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

**Generated At:** 2026-06-24 09:56:39

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
| nxf-2x2 | 2 | 2.0 | - | x86_64 |  | $0.0940 | 1 |
| bxf-2x8 | 2 | 8.0 | - | x86_64 |  | $0.1170 | 1 |
| bxf-4x16 | 4 | 16.0 | - | x86_64 |  | $0.2350 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| r026-c8e249d4-f148-4416-a3c6-555b7a02f67d | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | 02h7_f8178175-8693-4c3c-bb11-296961b82139 | Running | 2 vCPU, 2.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-vnet-01<br>**Subnet:** my06-subnet-01<br>**Public IP:** 159.23.94.241<br>**Private IP:** 10.0.1.6<br>**SGs:** my06-sg-01<br>**SSH:** my06-sshkey-01 |
| my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | 02h7_faa52016-7f78-46a7-bd3e-550531ee6879 | Running | 2 vCPU, 8.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-vnet-01<br>**Subnet:** my06-subnet-01<br>**Public IP:** 159.23.91.18<br>**Private IP:** 10.0.1.5<br>**SGs:** my06-sg-03<br>**SSH:** my06-sshkey-01 |
| my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | 02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-vnet-01<br>**Subnet:** my06-subnet-01<br>**Public IP:** 159.23.102.203<br>**Private IP:** 10.0.1.4<br>**SGs:** my06-sg-02<br>**SSH:** my06-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my06-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my06-vnet-01 |
| **CSP VNet ID** | r026-4650a745-f71b-4cfc-9c2f-bb503634ac18 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my06-subnet-01 | 02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179 | 10.0.1.0/24 | au-syd-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my06-sshkey-01 | r026-f149af7c-f2ec-41c2-937f-f4f592fe4602 |  | SHA256:hpm7J8C7RXdoPYHRNRKt9tB0B04Zuncu03GzqKEshzM |

### Security Groups

#### Security Group: my06-sg-01

| Property | Value |
|----------|-------|
| **Name** | my06-sg-01 |
| **CSP Security Group ID** | r026-4a4bcf87-47c6-4e4a-aef9-14c9b0081370 |
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
| **CSP Security Group ID** | r026-51dd6208-ab6c-486a-b2c9-b1e68dcf7791 |
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
| **CSP Security Group ID** | r026-fca6d682-9c93-4321-8801-d4436cf67c14 |
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

*Report generated: 2026-06-24 09:56:44*

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
| 1 | **VM Name:** my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** 02h7_f8178175-8693-4c3c-bb11-296961b82139<br>**Label(sourceMachineId):** vm-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** 02h7_faa52016-7f78-46a7-bd3e-550531ee6879<br>**Label(sourceMachineId):** vm-ec288dd0-c6fa-8a49-2f60 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 |
| 3 | **VM Name:** my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** 02h7_15d76d40-be51-4a99-8c5d-7fe9d8939e58<br>**Label(sourceMachineId):** vm-ec2d32b5-98fb-5a96-7913 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 |

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

**CSP ID:** r026-4a4bcf87-47c6-4e4a-aef9-14c9b0081370 | **VNet:** my06-vnet-01 | **Rules:** 14

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

**CSP ID:** r026-51dd6208-ab6c-486a-b2c9-b1e68dcf7791 | **VNet:** my06-vnet-01 | **Rules:** 19

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

**CSP ID:** r026-fca6d682-9c93-4321-8801-d4436cf67c14 | **VNet:** my06-vnet-01 | **Rules:** 19

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
| 1 | **Name:** my06-vnet-01<br>**ID:** r026-4650a745-f71b-4cfc-9c2f-bb503634ac18 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my06-subnet-01<br>**ID:** 02h7-7b0aa5a3-b655-489f-bc3c-2146af5f1179 | 10.0.1.0/24 | my06-vnet-01 |

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
| 1 | my06-sshkey-01 | r026-f149af7c-f2ec-41c2-937f-f4f592fe4602 | SHA256:hpm7J8C7RXdoPYHRNRKt9tB0B04Zuncu03GzqKEshzM | Used by all 3 VMs |

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

