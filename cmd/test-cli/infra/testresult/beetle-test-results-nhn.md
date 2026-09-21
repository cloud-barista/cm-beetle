# CM-Beetle test results for NHN

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with NHN cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.1+ (39d1774)
- imdl: v0.1.15+ (39d1774)
- CB-Tumblebug: v0.13.3
- CB-Spider: v0.13.4
- CB-MapUI: v0.13.7
- Target CSP: NHN
- Target Region: kr1
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: September 21, 2026
- Test Time: 13:22:32 KST
- Test Execution: 2026-09-21 13:22:32 KST

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

## Test result for NHN

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 6.306s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 412ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ❌ **FAIL** | 49.176s | Fail |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ⏭️ **SKIP** | 0s | Skip |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ⏭️ **SKIP** | 0s | Skip |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ⏭️ **SKIP** | 0s | Skip |
| 7 | Remote Command Accessibility Check | ⏭️ **SKIP** | 0s | Skip |
| 8 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ⏭️ **SKIP** | 0s | Skip |
| 9 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ⏭️ **SKIP** | 0s | Skip |
| 10 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 3ms | Pass |

**Overall Result**: 3/10 tests passed, 6 skipped ❌

**Total Duration**: 1m41.04110752s

*Test executed on September 21, 2026 at 13:22:32 KST (2026-09-21 13:22:32 KST) using CM-Beetle automated test CLI*

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
    "csp": "nhn",
    "region": "kr1"
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
      "status": "partially-matched",
      "description": "Candidate #1 | partially-matched | Overall Match Rate: Min=75.0% Max=100.0% Avg=91.7% | VMs: 3 total, 0 matched, 3 acceptable",
      "targetCloud": {
        "csp": "nhn",
        "region": "kr1"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "nhn-kr1",
            "specId": "nhn+kr1+c2.c2m2",
            "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 20,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "nhn-kr1",
            "specId": "nhn+kr1+r2.c4m16",
            "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 20,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "nhn-kr1",
            "specId": "nhn+kr1+r2.c2m8",
            "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 20,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "nhn-kr1",
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
        "connectionName": "nhn-kr1",
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
          "id": "nhn+kr1+c2.c2m2",
          "uid": "tbcqsc0oqtknhnnqaunv",
          "cspSpecName": "c2.c2m2",
          "name": "nhn+kr1+c2.c2m2",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "costPerHour": 0.042179998,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ID",
              "value": "f5ebd3a7-f978-477d-8dcf-2a1578e29df3"
            },
            {
              "key": "Name",
              "value": "c2.c2m2"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/f5ebd3a7-f978-477d-8dcf-2a1578e29df3,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/f5ebd3a7-f978-477d-8dcf-2a1578e29df3,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "2048"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "2"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:core}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        },
        {
          "id": "nhn+kr1+r2.c4m16",
          "uid": "tb9ngrgm4pd16v667jgg",
          "cspSpecName": "r2.c4m16",
          "name": "nhn+kr1+r2.c4m16",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "costPerHour": 0.19535999,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ID",
              "value": "8339613a-2837-4c72-906b-a8a47aa17c3f"
            },
            {
              "key": "Name",
              "value": "r2.c4m16"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/8339613a-2837-4c72-906b-a8a47aa17c3f,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/8339613a-2837-4c72-906b-a8a47aa17c3f,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "16384"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "4"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:memory}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        },
        {
          "id": "nhn+kr1+r2.c2m8",
          "uid": "tbv81tfmshja5ltn83iu",
          "cspSpecName": "r2.c2m8",
          "name": "nhn+kr1+r2.c2m8",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "costPerHour": 0.09472,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ID",
              "value": "db51400e-85a1-4d65-930a-52d1d40941a9"
            },
            {
              "key": "Name",
              "value": "r2.c2m8"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/db51400e-85a1-4d65-930a-52d1d40941a9,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/db51400e-85a1-4d65-930a-52d1d40941a9,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "8192"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "2"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:memory}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "nhn",
          "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "regionList": [
            "kr1"
          ],
          "id": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "uid": "tb76dofaj88ahqu4op2v",
          "name": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "nhn-kr1",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:30 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": true,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ID",
              "value": "0f07c795-2a46-44fc-a61b-fa0d96763ce2"
            },
            {
              "key": "Name",
              "value": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
            },
            {
              "key": "Status",
              "value": "active"
            },
            {
              "key": "Tags",
              "value": "BASE"
            },
            {
              "key": "ContainerFormat",
              "value": "bare"
            },
            {
              "key": "DiskFormat",
              "value": "qcow2"
            },
            {
              "key": "MinDiskGigabytes",
              "value": "20"
            },
            {
              "key": "MinRAMMegabytes",
              "value": "0"
            },
            {
              "key": "Owner",
              "value": "c289b99209ca4e189095cdecebbd092d"
            },
            {
              "key": "Protected",
              "value": "true"
            },
            {
              "key": "Visibility",
              "value": "public"
            },
            {
              "key": "Hidden",
              "value": "false"
            },
            {
              "key": "Checksum",
              "value": "4c563c67af5d73b15e7003af724d6b55"
            },
            {
              "key": "Properties",
              "value": "{deprecate_date:null,description:Ubuntu Server 22.04.5 LTS (2026.03.10),hw_cpu_sockets:1,hw_qemu_guest_agent:yes,hw_vif_multiqueue_enabled:true,login_username:ubuntu,max_cpu:,min_cpu:0,monitoring_agent:sysmon,nhncloud_allow_autoscale:true,nhncloud_allow_cgroup:v2,nhncloud_allow_compute_flavor:true,nhncloud_allow_cpu_flavor:true,nhncloud_allow_download:false,nhncloud_allow_gpu_flavor:false,nhncloud_allow_image_create:true,nhncloud_allow_imagebuilder:true,nhncloud_allow_instance_template:true,nhncloud_allow_local_bootdisk_flavor:true,nhncloud_allow_nks_cpu_flavor:false,nhncloud_allow_nks_gpu_flavor:false,nhncloud_allow_user_script:true,nhncloud_category:OS,nhncloud_product:compute,os_architecture:amd64,os_distro:ubuntu,os_type:linux,os_version:Server 22.04 LTS,project_domain:WDI;NORMAL,release_date:2026.03.10,tc_env:cloudmon,sysmon}"
            },
            {
              "key": "CreatedAt",
              "value": "2026-03-09T01:34:35Z"
            },
            {
              "key": "UpdatedAt",
              "value": "2026-03-09T21:14:13Z"
            },
            {
              "key": "File",
              "value": "/v2/images/0f07c795-2a46-44fc-a61b-fa0d96763ce2/file"
            },
            {
              "key": "Schema",
              "value": "/v2/schemas/image"
            },
            {
              "key": "VirtualSize",
              "value": "0"
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
          "connectionName": "nhn-kr1",
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
          "connectionName": "nhn-kr1",
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
          "connectionName": "nhn-kr1",
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
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=50.0% Max=100.0% Avg=75.0% | VMs: 3 total, 0 matched, 3 acceptable",
      "targetCloud": {
        "csp": "nhn",
        "region": "kr1"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=75.0%",
            "connectionName": "nhn-kr1",
            "specId": "nhn+kr1+m2.c2m4",
            "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 20,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=50.0% Memory=100.0% Image=75.0%",
            "connectionName": "nhn-kr1",
            "specId": "nhn+kr1+m2.c8m16",
            "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 20,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=50.0% Memory=100.0% Image=75.0%",
            "connectionName": "nhn-kr1",
            "specId": "nhn+kr1+m2.c4m8",
            "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 20,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "nhn-kr1",
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
        "connectionName": "nhn-kr1",
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
          "id": "nhn+kr1+c2.c2m2",
          "uid": "tbcqsc0oqtknhnnqaunv",
          "cspSpecName": "c2.c2m2",
          "name": "nhn+kr1+c2.c2m2",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "costPerHour": 0.042179998,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ID",
              "value": "f5ebd3a7-f978-477d-8dcf-2a1578e29df3"
            },
            {
              "key": "Name",
              "value": "c2.c2m2"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/f5ebd3a7-f978-477d-8dcf-2a1578e29df3,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/f5ebd3a7-f978-477d-8dcf-2a1578e29df3,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "2048"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "2"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:core}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        },
        {
          "id": "nhn+kr1+r2.c4m16",
          "uid": "tb9ngrgm4pd16v667jgg",
          "cspSpecName": "r2.c4m16",
          "name": "nhn+kr1+r2.c4m16",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "costPerHour": 0.19535999,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ID",
              "value": "8339613a-2837-4c72-906b-a8a47aa17c3f"
            },
            {
              "key": "Name",
              "value": "r2.c4m16"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/8339613a-2837-4c72-906b-a8a47aa17c3f,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/8339613a-2837-4c72-906b-a8a47aa17c3f,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "16384"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "4"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:memory}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        },
        {
          "id": "nhn+kr1+r2.c2m8",
          "uid": "tbv81tfmshja5ltn83iu",
          "cspSpecName": "r2.c2m8",
          "name": "nhn+kr1+r2.c2m8",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "costPerHour": 0.09472,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ID",
              "value": "db51400e-85a1-4d65-930a-52d1d40941a9"
            },
            {
              "key": "Name",
              "value": "r2.c2m8"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/db51400e-85a1-4d65-930a-52d1d40941a9,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/db51400e-85a1-4d65-930a-52d1d40941a9,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "8192"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "2"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:memory}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        },
        {
          "id": "nhn+kr1+m2.c2m4",
          "uid": "tbpnvei7jqqgj4d9678f",
          "cspSpecName": "m2.c2m4",
          "name": "nhn+kr1+m2.c2m4",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 4,
          "costPerHour": 0.06882,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ID",
              "value": "35a73b57-58a7-434d-aa08-5249aaa95b3e"
            },
            {
              "key": "Name",
              "value": "m2.c2m4"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/35a73b57-58a7-434d-aa08-5249aaa95b3e,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/35a73b57-58a7-434d-aa08-5249aaa95b3e,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "4096"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "2"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:general}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        },
        {
          "id": "nhn+kr1+m2.c8m16",
          "uid": "tbstcl8ocf1p2kvit3nv",
          "cspSpecName": "m2.c8m16",
          "name": "nhn+kr1+m2.c8m16",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 8,
          "memoryGiB": 16,
          "costPerHour": 0.29525998,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ID",
              "value": "eaddf5bd-2725-409a-ac57-044935747183"
            },
            {
              "key": "Name",
              "value": "m2.c8m16"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/eaddf5bd-2725-409a-ac57-044935747183,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/eaddf5bd-2725-409a-ac57-044935747183,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "16384"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "8"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:general}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        },
        {
          "id": "nhn+kr1+m2.c4m8",
          "uid": "tbj7sfnortak7tbokrgk",
          "cspSpecName": "m2.c4m8",
          "name": "nhn+kr1+m2.c4m8",
          "namespace": "system",
          "connectionName": "nhn-kr1",
          "providerName": "nhn",
          "regionName": "kr1",
          "regionLatitude": 37.390889,
          "regionLongitude": 127.096792,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 8,
          "costPerHour": 0.14503999,
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
          "rootDiskType": "default",
          "rootDiskSize": 0,
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ID",
              "value": "edc79d63-98c3-4b77-a2d4-482d70e6b554"
            },
            {
              "key": "Name",
              "value": "m2.c4m8"
            },
            {
              "key": "Links",
              "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/edc79d63-98c3-4b77-a2d4-482d70e6b554,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/edc79d63-98c3-4b77-a2d4-482d70e6b554,rel:bookmark}"
            },
            {
              "key": "RAM",
              "value": "8192"
            },
            {
              "key": "Disabled",
              "value": "false"
            },
            {
              "key": "VCPUs",
              "value": "4"
            },
            {
              "key": "ExtraSpecs",
              "value": "{flavor_type:general}"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "RxTxFactor",
              "value": "1.00"
            },
            {
              "key": "Ephemeral",
              "value": "0"
            },
            {
              "key": "Disk",
              "value": "0"
            },
            {
              "key": "Notice!!",
              "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "nhn",
          "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "regionList": [
            "kr1"
          ],
          "id": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "uid": "tb76dofaj88ahqu4op2v",
          "name": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "nhn-kr1",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:30 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": true,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ID",
              "value": "0f07c795-2a46-44fc-a61b-fa0d96763ce2"
            },
            {
              "key": "Name",
              "value": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
            },
            {
              "key": "Status",
              "value": "active"
            },
            {
              "key": "Tags",
              "value": "BASE"
            },
            {
              "key": "ContainerFormat",
              "value": "bare"
            },
            {
              "key": "DiskFormat",
              "value": "qcow2"
            },
            {
              "key": "MinDiskGigabytes",
              "value": "20"
            },
            {
              "key": "MinRAMMegabytes",
              "value": "0"
            },
            {
              "key": "Owner",
              "value": "c289b99209ca4e189095cdecebbd092d"
            },
            {
              "key": "Protected",
              "value": "true"
            },
            {
              "key": "Visibility",
              "value": "public"
            },
            {
              "key": "Hidden",
              "value": "false"
            },
            {
              "key": "Checksum",
              "value": "4c563c67af5d73b15e7003af724d6b55"
            },
            {
              "key": "Properties",
              "value": "{deprecate_date:null,description:Ubuntu Server 22.04.5 LTS (2026.03.10),hw_cpu_sockets:1,hw_qemu_guest_agent:yes,hw_vif_multiqueue_enabled:true,login_username:ubuntu,max_cpu:,min_cpu:0,monitoring_agent:sysmon,nhncloud_allow_autoscale:true,nhncloud_allow_cgroup:v2,nhncloud_allow_compute_flavor:true,nhncloud_allow_cpu_flavor:true,nhncloud_allow_download:false,nhncloud_allow_gpu_flavor:false,nhncloud_allow_image_create:true,nhncloud_allow_imagebuilder:true,nhncloud_allow_instance_template:true,nhncloud_allow_local_bootdisk_flavor:true,nhncloud_allow_nks_cpu_flavor:false,nhncloud_allow_nks_gpu_flavor:false,nhncloud_allow_user_script:true,nhncloud_category:OS,nhncloud_product:compute,os_architecture:amd64,os_distro:ubuntu,os_type:linux,os_version:Server 22.04 LTS,project_domain:WDI;NORMAL,release_date:2026.03.10,tc_env:cloudmon,sysmon}"
            },
            {
              "key": "CreatedAt",
              "value": "2026-03-09T01:34:35Z"
            },
            {
              "key": "UpdatedAt",
              "value": "2026-03-09T21:14:13Z"
            },
            {
              "key": "File",
              "value": "/v2/images/0f07c795-2a46-44fc-a61b-fa0d96763ce2/file"
            },
            {
              "key": "Schema",
              "value": "/v2/schemas/image"
            },
            {
              "key": "VirtualSize",
              "value": "0"
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
          "connectionName": "nhn-kr1",
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
          "connectionName": "nhn-kr1",
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
          "connectionName": "nhn-kr1",
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
  "status": "partially-matched",
  "description": "Candidate #1 | partially-matched | Overall Match Rate: Min=75.0% Max=100.0% Avg=91.7% | VMs: 3 total, 0 matched, 3 acceptable",
  "targetCloud": {
    "csp": "nhn",
    "region": "kr1"
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
        "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
        "connectionName": "nhn-kr1",
        "specId": "nhn+kr1+c2.c2m2",
        "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-01"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 20,
        "dataDiskIds": null
      },
      {
        "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "nodeGroupSize": 1,
        "label": {
          "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
        },
        "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
        "connectionName": "nhn-kr1",
        "specId": "nhn+kr1+r2.c4m16",
        "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-02"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 20,
        "dataDiskIds": null
      },
      {
        "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nodeGroupSize": 1,
        "label": {
          "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
        },
        "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
        "connectionName": "nhn-kr1",
        "specId": "nhn+kr1+r2.c2m8",
        "imageId": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-03"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 20,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "vnet-01",
    "connectionName": "nhn-kr1",
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
    "connectionName": "nhn-kr1",
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
      "id": "nhn+kr1+c2.c2m2",
      "uid": "tbcqsc0oqtknhnnqaunv",
      "cspSpecName": "c2.c2m2",
      "name": "nhn+kr1+c2.c2m2",
      "namespace": "system",
      "connectionName": "nhn-kr1",
      "providerName": "nhn",
      "regionName": "kr1",
      "regionLatitude": 37.390889,
      "regionLongitude": 127.096792,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 2,
      "costPerHour": 0.042179998,
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
      "rootDiskType": "default",
      "rootDiskSize": 0,
      "systemLabel": "from-assets",
      "details": [
        {
          "key": "ID",
          "value": "f5ebd3a7-f978-477d-8dcf-2a1578e29df3"
        },
        {
          "key": "Name",
          "value": "c2.c2m2"
        },
        {
          "key": "Links",
          "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/f5ebd3a7-f978-477d-8dcf-2a1578e29df3,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/f5ebd3a7-f978-477d-8dcf-2a1578e29df3,rel:bookmark}"
        },
        {
          "key": "RAM",
          "value": "2048"
        },
        {
          "key": "Disabled",
          "value": "false"
        },
        {
          "key": "VCPUs",
          "value": "2"
        },
        {
          "key": "ExtraSpecs",
          "value": "{flavor_type:core}"
        },
        {
          "key": "IsPublic",
          "value": "true"
        },
        {
          "key": "RxTxFactor",
          "value": "1.00"
        },
        {
          "key": "Ephemeral",
          "value": "0"
        },
        {
          "key": "Disk",
          "value": "0"
        },
        {
          "key": "Notice!!",
          "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
        }
      ]
    },
    {
      "id": "nhn+kr1+r2.c4m16",
      "uid": "tb9ngrgm4pd16v667jgg",
      "cspSpecName": "r2.c4m16",
      "name": "nhn+kr1+r2.c4m16",
      "namespace": "system",
      "connectionName": "nhn-kr1",
      "providerName": "nhn",
      "regionName": "kr1",
      "regionLatitude": 37.390889,
      "regionLongitude": 127.096792,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 16,
      "costPerHour": 0.19535999,
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
      "rootDiskType": "default",
      "rootDiskSize": 0,
      "systemLabel": "from-assets",
      "details": [
        {
          "key": "ID",
          "value": "8339613a-2837-4c72-906b-a8a47aa17c3f"
        },
        {
          "key": "Name",
          "value": "r2.c4m16"
        },
        {
          "key": "Links",
          "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/8339613a-2837-4c72-906b-a8a47aa17c3f,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/8339613a-2837-4c72-906b-a8a47aa17c3f,rel:bookmark}"
        },
        {
          "key": "RAM",
          "value": "16384"
        },
        {
          "key": "Disabled",
          "value": "false"
        },
        {
          "key": "VCPUs",
          "value": "4"
        },
        {
          "key": "ExtraSpecs",
          "value": "{flavor_type:memory}"
        },
        {
          "key": "IsPublic",
          "value": "true"
        },
        {
          "key": "RxTxFactor",
          "value": "1.00"
        },
        {
          "key": "Ephemeral",
          "value": "0"
        },
        {
          "key": "Disk",
          "value": "0"
        },
        {
          "key": "Notice!!",
          "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
        }
      ]
    },
    {
      "id": "nhn+kr1+r2.c2m8",
      "uid": "tbv81tfmshja5ltn83iu",
      "cspSpecName": "r2.c2m8",
      "name": "nhn+kr1+r2.c2m8",
      "namespace": "system",
      "connectionName": "nhn-kr1",
      "providerName": "nhn",
      "regionName": "kr1",
      "regionLatitude": 37.390889,
      "regionLongitude": 127.096792,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 8,
      "costPerHour": 0.09472,
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
      "rootDiskType": "default",
      "rootDiskSize": 0,
      "systemLabel": "from-assets",
      "details": [
        {
          "key": "ID",
          "value": "db51400e-85a1-4d65-930a-52d1d40941a9"
        },
        {
          "key": "Name",
          "value": "r2.c2m8"
        },
        {
          "key": "Links",
          "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/db51400e-85a1-4d65-930a-52d1d40941a9,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/db51400e-85a1-4d65-930a-52d1d40941a9,rel:bookmark}"
        },
        {
          "key": "RAM",
          "value": "8192"
        },
        {
          "key": "Disabled",
          "value": "false"
        },
        {
          "key": "VCPUs",
          "value": "2"
        },
        {
          "key": "ExtraSpecs",
          "value": "{flavor_type:memory}"
        },
        {
          "key": "IsPublic",
          "value": "true"
        },
        {
          "key": "RxTxFactor",
          "value": "1.00"
        },
        {
          "key": "Ephemeral",
          "value": "0"
        },
        {
          "key": "Disk",
          "value": "0"
        },
        {
          "key": "Notice!!",
          "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
        }
      ]
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "nhn",
      "cspImageName": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
      "regionList": [
        "kr1"
      ],
      "id": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
      "uid": "tb76dofaj88ahqu4op2v",
      "name": "0f07c795-2a46-44fc-a61b-fa0d96763ce2",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "nhn-kr1",
      "infraType": "",
      "fetchedTime": "2026.08.21 13:57:30 Fri",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": true,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Ubuntu Server 22.04.5 LTS (2026.03.10)",
      "osDiskType": "NA",
      "osDiskSizeGB": 20,
      "imageStatus": "Available",
      "details": [
        {
          "key": "ID",
          "value": "0f07c795-2a46-44fc-a61b-fa0d96763ce2"
        },
        {
          "key": "Name",
          "value": "Ubuntu Server 22.04.5 LTS (2026.03.10)"
        },
        {
          "key": "Status",
          "value": "active"
        },
        {
          "key": "Tags",
          "value": "BASE"
        },
        {
          "key": "ContainerFormat",
          "value": "bare"
        },
        {
          "key": "DiskFormat",
          "value": "qcow2"
        },
        {
          "key": "MinDiskGigabytes",
          "value": "20"
        },
        {
          "key": "MinRAMMegabytes",
          "value": "0"
        },
        {
          "key": "Owner",
          "value": "c289b99209ca4e189095cdecebbd092d"
        },
        {
          "key": "Protected",
          "value": "true"
        },
        {
          "key": "Visibility",
          "value": "public"
        },
        {
          "key": "Hidden",
          "value": "false"
        },
        {
          "key": "Checksum",
          "value": "4c563c67af5d73b15e7003af724d6b55"
        },
        {
          "key": "Properties",
          "value": "{deprecate_date:null,description:Ubuntu Server 22.04.5 LTS (2026.03.10),hw_cpu_sockets:1,hw_qemu_guest_agent:yes,hw_vif_multiqueue_enabled:true,login_username:ubuntu,max_cpu:,min_cpu:0,monitoring_agent:sysmon,nhncloud_allow_autoscale:true,nhncloud_allow_cgroup:v2,nhncloud_allow_compute_flavor:true,nhncloud_allow_cpu_flavor:true,nhncloud_allow_download:false,nhncloud_allow_gpu_flavor:false,nhncloud_allow_image_create:true,nhncloud_allow_imagebuilder:true,nhncloud_allow_instance_template:true,nhncloud_allow_local_bootdisk_flavor:true,nhncloud_allow_nks_cpu_flavor:false,nhncloud_allow_nks_gpu_flavor:false,nhncloud_allow_user_script:true,nhncloud_category:OS,nhncloud_product:compute,os_architecture:amd64,os_distro:ubuntu,os_type:linux,os_version:Server 22.04 LTS,project_domain:WDI;NORMAL,release_date:2026.03.10,tc_env:cloudmon,sysmon}"
        },
        {
          "key": "CreatedAt",
          "value": "2026-03-09T01:34:35Z"
        },
        {
          "key": "UpdatedAt",
          "value": "2026-03-09T21:14:13Z"
        },
        {
          "key": "File",
          "value": "/v2/images/0f07c795-2a46-44fc-a61b-fa0d96763ce2/file"
        },
        {
          "key": "Schema",
          "value": "/v2/schemas/image"
        },
        {
          "key": "VirtualSize",
          "value": "0"
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
      "connectionName": "nhn-kr1",
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
      "connectionName": "nhn-kr1",
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
      "connectionName": "nhn-kr1",
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
  "resourceType": "",
  "id": "",
  "name": "",
  "status": "",
  "statusCount": {
    "countTotal": 0,
    "countCreating": 0,
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
  "targetStatus": "",
  "targetAction": "",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": null,
  "systemLabel": "",
  "systemMessage": null,
  "description": "",
  "node": null,
  "newNodeList": null
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

- **Status**: ❌ **FAILED**
- **Error**: No response received

### Test Case 5: Get a list of infra IDs

#### 5.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra?option=id`
- **Purpose**: Retrieve infra IDs only (lightweight response)
- **Namespace ID**: `mig01`
- **Query Parameter**: `option=id`
- **Request Body**: None (GET request)

#### 5.2 API Response Information

- **Status**: ❌ **FAILED**
- **Error**: No response received

**Error Message**:

```
Test skipped due to previous test failure
```

### Test Case 7: Remote Command Accessibility Check

#### 7.1 Test Information

- **Test Type**: SSH Connectivity Test for All VMs
- **Purpose**: Verify that all migrated VMs are accessible via SSH
- **Method**: Extract public IP and SSH key from MCI access info for each VM, then execute remote command
- **Command Executed**: `uname -a` (to verify system information)
- **Authentication**: SSH key-based authentication
- **Scope**: Tests all VMs across all subgroups in the MCI

#### 7.2 Test Result Information

- **Status**: ⏭️ **SKIPPED**
- **Reason**: 

### Test Case 8: Target Infrastructure Summary

#### 8.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`
- **Purpose**: Get a summary of the migrated target infrastructure in Markdown format
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier
- **Query Parameter**: `format=md`

#### 8.2 API Response Information

- **Status**: ⏭️ **SKIPPED**
- **Reason**: 

### Test Case 9: Migration Report

#### 9.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Generate a comprehensive migration report matching source to target
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier

#### 9.2 API Response Information

- **Status**: ⏭️ **SKIPPED**
- **Reason**: Test skipped: MCI ID is required but not available

### Test Case 10: Delete the migrated computing infra

#### 10.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Delete the migrated infrastructure and clean up resources
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier to delete
- **Query Parameter**: `option=terminate` (terminates all resources)
- **Request Body**: None (DELETE request)

#### 10.2 API Response Information

- **Status**: ⏭️ **SKIPPED**
- **Reason**: Test skipped: MCI ID is required but not available

