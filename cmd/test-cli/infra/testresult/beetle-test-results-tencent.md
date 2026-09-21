# CM-Beetle test results for TENCENT

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with TENCENT cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.6.1+ (39d1774)
- imdl: v0.1.15+ (39d1774)
- CB-Tumblebug: v0.13.3
- CB-Spider: v0.13.4
- CB-MapUI: v0.13.7
- Target CSP: TENCENT
- Target Region: ap-seoul
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: September 21, 2026
- Test Time: 13:17:06 KST
- Test Execution: 2026-09-21 13:17:06 KST

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

## Test result for TENCENT

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 7.233s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 396ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 2m39.059s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 9ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 4ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 16ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 56.752s | Pass |
| 8 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.277s | Pass |
| 9 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.287s | Pass |
| 10 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 53.725s | Pass |

**Overall Result**: 10/10 tests passed ✅

**Total Duration**: 5m53.045630695s

*Test executed on September 21, 2026 at 13:17:06 KST (2026-09-21 13:17:06 KST) using CM-Beetle automated test CLI*

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
    "csp": "tencent",
    "region": "ap-seoul"
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
        "csp": "tencent",
        "region": "ap-seoul"
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
            "connectionName": "tencent-ap-seoul",
            "specId": "tencent+ap-seoul+bf1.medium2",
            "imageId": "img-7rotv4ux",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "tencent-ap-seoul",
            "specId": "tencent+ap-seoul+bf1.large16",
            "imageId": "img-7rotv4ux",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "tencent-ap-seoul",
            "specId": "tencent+ap-seoul+bf1.medium8",
            "imageId": "img-7rotv4ux",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "tencent-ap-seoul",
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
        "connectionName": "tencent-ap-seoul",
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
          "id": "tencent+ap-seoul+bf1.medium2",
          "uid": "tb71icnf2ipgc6r7vlil",
          "cspSpecName": "BF1.MEDIUM2",
          "name": "tencent+ap-seoul+bf1.medium2",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.02,
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
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "BF1.MEDIUM2"
            },
            {
              "key": "InstanceChargeType",
              "value": "POSTPAID_BY_HOUR"
            },
            {
              "key": "NetworkCard",
              "value": "100"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "Memory",
              "value": "2"
            },
            {
              "key": "InstanceFamily",
              "value": "BF1"
            },
            {
              "key": "TypeName",
              "value": "BF1"
            },
            {
              "key": "Status",
              "value": "SELL"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.02,ChargeUnit:HOUR,Discount:100,UnitPriceDiscount:0.02,UnitPriceSecondStep:0.02,UnitPriceDiscountSecondStep:0.02,UnitPriceThirdStep:0.02,UnitPriceDiscountThirdStep:0.02}"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "25"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "-"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "-"
            },
            {
              "key": "StatusCategory",
              "value": "UnderStock"
            }
          ]
        },
        {
          "id": "tencent+ap-seoul+bf1.large16",
          "uid": "tb1cmris4eu56fa89t2k",
          "cspSpecName": "BF1.LARGE16",
          "name": "tencent+ap-seoul+bf1.large16",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.15,
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
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "BF1.LARGE16"
            },
            {
              "key": "InstanceChargeType",
              "value": "POSTPAID_BY_HOUR"
            },
            {
              "key": "NetworkCard",
              "value": "100"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "4"
            },
            {
              "key": "Memory",
              "value": "16"
            },
            {
              "key": "InstanceFamily",
              "value": "BF1"
            },
            {
              "key": "TypeName",
              "value": "BF1"
            },
            {
              "key": "Status",
              "value": "SOLD_OUT"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.15,ChargeUnit:HOUR,Discount:100,UnitPriceDiscount:0.15,UnitPriceSecondStep:0.15,UnitPriceDiscountSecondStep:0.15,UnitPriceThirdStep:0.15,UnitPriceDiscountThirdStep:0.15}"
            },
            {
              "key": "SoldOutReason",
              "value": "ResourcesSoldOut.SpecifiedInstanceType"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "25"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "-"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "-"
            },
            {
              "key": "StatusCategory",
              "value": "WithoutStock"
            }
          ]
        },
        {
          "id": "tencent+ap-seoul+bf1.medium8",
          "uid": "tb7duhhnsa5d7414bjdn",
          "cspSpecName": "BF1.MEDIUM8",
          "name": "tencent+ap-seoul+bf1.medium8",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.07,
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
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "BF1.MEDIUM8"
            },
            {
              "key": "InstanceChargeType",
              "value": "POSTPAID_BY_HOUR"
            },
            {
              "key": "NetworkCard",
              "value": "100"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "Memory",
              "value": "8"
            },
            {
              "key": "InstanceFamily",
              "value": "BF1"
            },
            {
              "key": "TypeName",
              "value": "BF1"
            },
            {
              "key": "Status",
              "value": "SOLD_OUT"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.07,ChargeUnit:HOUR,Discount:100,UnitPriceDiscount:0.07,UnitPriceSecondStep:0.07,UnitPriceDiscountSecondStep:0.07,UnitPriceThirdStep:0.07,UnitPriceDiscountThirdStep:0.07}"
            },
            {
              "key": "SoldOutReason",
              "value": "ResourcesSoldOut.SpecifiedInstanceType"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "25"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "-"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "-"
            },
            {
              "key": "StatusCategory",
              "value": "WithoutStock"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "tencent",
          "cspImageName": "img-7rotv4ux",
          "regionList": [
            "ap-bangkok",
            "ap-beijing",
            "ap-chengdu",
            "ap-chongqing",
            "ap-guangzhou",
            "ap-hongkong",
            "ap-jakarta",
            "ap-nanjing",
            "ap-seoul",
            "ap-shanghai",
            "ap-singapore",
            "ap-tokyo",
            "eu-frankfurt",
            "me-saudi-arabia",
            "na-ashburn",
            "na-siliconvalley",
            "sa-saopaulo"
          ],
          "id": "img-7rotv4ux",
          "uid": "tbm5u74berjs1hfbpgol",
          "name": "img-7rotv4ux",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "tencent-sa-saopaulo",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:36 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ImageId",
              "value": "img-7rotv4ux"
            },
            {
              "key": "OsName",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageType",
              "value": "PUBLIC_IMAGE"
            },
            {
              "key": "ImageName",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageDescription",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageSize",
              "value": "20"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "ImageState",
              "value": "NORMAL"
            },
            {
              "key": "Platform",
              "value": "Ubuntu"
            },
            {
              "key": "ImageSource",
              "value": "OFFICIAL"
            },
            {
              "key": "IsSupportCloudinit",
              "value": "true"
            },
            {
              "key": "ImageDeprecated",
              "value": "false"
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
          "connectionName": "tencent-ap-seoul",
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
          "connectionName": "tencent-ap-seoul",
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
          "connectionName": "tencent-ap-seoul",
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
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=75.0% Max=100.0% Avg=91.7% | VMs: 3 total, 0 matched, 3 acceptable",
      "targetCloud": {
        "csp": "tencent",
        "region": "ap-seoul"
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
            "connectionName": "tencent-ap-seoul",
            "specId": "tencent+ap-seoul+sa2.medium2",
            "imageId": "img-7rotv4ux",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "tencent-ap-seoul",
            "specId": "tencent+ap-seoul+sa2.large16",
            "imageId": "img-7rotv4ux",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "tencent-ap-seoul",
            "specId": "tencent+ap-seoul+sa2.medium8",
            "imageId": "img-7rotv4ux",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "tencent-ap-seoul",
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
        "connectionName": "tencent-ap-seoul",
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
          "id": "tencent+ap-seoul+bf1.medium2",
          "uid": "tb71icnf2ipgc6r7vlil",
          "cspSpecName": "BF1.MEDIUM2",
          "name": "tencent+ap-seoul+bf1.medium2",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.02,
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
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "BF1.MEDIUM2"
            },
            {
              "key": "InstanceChargeType",
              "value": "POSTPAID_BY_HOUR"
            },
            {
              "key": "NetworkCard",
              "value": "100"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "Memory",
              "value": "2"
            },
            {
              "key": "InstanceFamily",
              "value": "BF1"
            },
            {
              "key": "TypeName",
              "value": "BF1"
            },
            {
              "key": "Status",
              "value": "SELL"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.02,ChargeUnit:HOUR,Discount:100,UnitPriceDiscount:0.02,UnitPriceSecondStep:0.02,UnitPriceDiscountSecondStep:0.02,UnitPriceThirdStep:0.02,UnitPriceDiscountThirdStep:0.02}"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "25"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "-"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "-"
            },
            {
              "key": "StatusCategory",
              "value": "UnderStock"
            }
          ]
        },
        {
          "id": "tencent+ap-seoul+bf1.large16",
          "uid": "tb1cmris4eu56fa89t2k",
          "cspSpecName": "BF1.LARGE16",
          "name": "tencent+ap-seoul+bf1.large16",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.15,
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
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "BF1.LARGE16"
            },
            {
              "key": "InstanceChargeType",
              "value": "POSTPAID_BY_HOUR"
            },
            {
              "key": "NetworkCard",
              "value": "100"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "4"
            },
            {
              "key": "Memory",
              "value": "16"
            },
            {
              "key": "InstanceFamily",
              "value": "BF1"
            },
            {
              "key": "TypeName",
              "value": "BF1"
            },
            {
              "key": "Status",
              "value": "SOLD_OUT"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.15,ChargeUnit:HOUR,Discount:100,UnitPriceDiscount:0.15,UnitPriceSecondStep:0.15,UnitPriceDiscountSecondStep:0.15,UnitPriceThirdStep:0.15,UnitPriceDiscountThirdStep:0.15}"
            },
            {
              "key": "SoldOutReason",
              "value": "ResourcesSoldOut.SpecifiedInstanceType"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "25"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "-"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "-"
            },
            {
              "key": "StatusCategory",
              "value": "WithoutStock"
            }
          ]
        },
        {
          "id": "tencent+ap-seoul+bf1.medium8",
          "uid": "tb7duhhnsa5d7414bjdn",
          "cspSpecName": "BF1.MEDIUM8",
          "name": "tencent+ap-seoul+bf1.medium8",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.07,
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
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "BF1.MEDIUM8"
            },
            {
              "key": "InstanceChargeType",
              "value": "POSTPAID_BY_HOUR"
            },
            {
              "key": "NetworkCard",
              "value": "100"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "Memory",
              "value": "8"
            },
            {
              "key": "InstanceFamily",
              "value": "BF1"
            },
            {
              "key": "TypeName",
              "value": "BF1"
            },
            {
              "key": "Status",
              "value": "SOLD_OUT"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.07,ChargeUnit:HOUR,Discount:100,UnitPriceDiscount:0.07,UnitPriceSecondStep:0.07,UnitPriceDiscountSecondStep:0.07,UnitPriceThirdStep:0.07,UnitPriceDiscountThirdStep:0.07}"
            },
            {
              "key": "SoldOutReason",
              "value": "ResourcesSoldOut.SpecifiedInstanceType"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "25"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "-"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "-"
            },
            {
              "key": "StatusCategory",
              "value": "WithoutStock"
            }
          ]
        },
        {
          "id": "tencent+ap-seoul+sa2.medium2",
          "uid": "tb7bjn4bv9f9aajc8511",
          "cspSpecName": "SA2.MEDIUM2",
          "name": "tencent+ap-seoul+sa2.medium2",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.02,
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
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "SA2.MEDIUM2"
            },
            {
              "key": "InstanceChargeType",
              "value": "SPOTPAID"
            },
            {
              "key": "NetworkCard",
              "value": "25"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "Memory",
              "value": "2"
            },
            {
              "key": "InstanceFamily",
              "value": "SA2"
            },
            {
              "key": "TypeName",
              "value": "SA2"
            },
            {
              "key": "Status",
              "value": "SELL"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.02,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.004,UnitPriceSecondStep:0.02,UnitPriceDiscountSecondStep:0.004,UnitPriceThirdStep:0.02,UnitPriceDiscountThirdStep:0.004}"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "30"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "AMD EPYC™ Rome"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "2.6GHz/3.3GHz"
            },
            {
              "key": "StatusCategory",
              "value": "UnderStock"
            }
          ]
        },
        {
          "id": "tencent+ap-seoul+sa2.large16",
          "uid": "tbmvl0oj6d85smbvccef",
          "cspSpecName": "SA2.LARGE16",
          "name": "tencent+ap-seoul+sa2.large16",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.16,
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
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "SA2.LARGE16"
            },
            {
              "key": "InstanceChargeType",
              "value": "SPOTPAID"
            },
            {
              "key": "NetworkCard",
              "value": "25"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "4"
            },
            {
              "key": "Memory",
              "value": "16"
            },
            {
              "key": "InstanceFamily",
              "value": "SA2"
            },
            {
              "key": "TypeName",
              "value": "SA2"
            },
            {
              "key": "Status",
              "value": "SOLD_OUT"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.16,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.032,UnitPriceSecondStep:0.16,UnitPriceDiscountSecondStep:0.032,UnitPriceThirdStep:0.16,UnitPriceDiscountThirdStep:0.032}"
            },
            {
              "key": "SoldOutReason",
              "value": "ResourcesSoldOut.SpecifiedInstanceType"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "50"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "AMD EPYC™ Rome"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "2.6GHz/3.3GHz"
            },
            {
              "key": "StatusCategory",
              "value": "WithoutStock"
            }
          ]
        },
        {
          "id": "tencent+ap-seoul+sa2.medium8",
          "uid": "tbg8tdksbg0e0o52dcqb",
          "cspSpecName": "SA2.MEDIUM8",
          "name": "tencent+ap-seoul+sa2.medium8",
          "namespace": "system",
          "connectionName": "tencent-ap-seoul",
          "providerName": "tencent",
          "regionName": "ap-seoul",
          "regionLatitude": 37.566536,
          "regionLongitude": 126.977966,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.08,
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
              "key": "Zone",
              "value": "ap-seoul-1"
            },
            {
              "key": "InstanceType",
              "value": "SA2.MEDIUM8"
            },
            {
              "key": "InstanceChargeType",
              "value": "SPOTPAID"
            },
            {
              "key": "NetworkCard",
              "value": "25"
            },
            {
              "key": "Externals",
              "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "Memory",
              "value": "8"
            },
            {
              "key": "InstanceFamily",
              "value": "SA2"
            },
            {
              "key": "TypeName",
              "value": "SA2"
            },
            {
              "key": "Status",
              "value": "SOLD_OUT"
            },
            {
              "key": "Price",
              "value": "{UnitPrice:0.08,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.016,UnitPriceSecondStep:0.08,UnitPriceDiscountSecondStep:0.016,UnitPriceThirdStep:0.08,UnitPriceDiscountThirdStep:0.016}"
            },
            {
              "key": "SoldOutReason",
              "value": "ResourcesSoldOut.SpecifiedInstanceType"
            },
            {
              "key": "InstanceBandwidth",
              "value": "1.5"
            },
            {
              "key": "InstancePps",
              "value": "30"
            },
            {
              "key": "StorageBlockAmount",
              "value": "0"
            },
            {
              "key": "CpuType",
              "value": "AMD EPYC™ Rome"
            },
            {
              "key": "Gpu",
              "value": "0"
            },
            {
              "key": "Fpga",
              "value": "0"
            },
            {
              "key": "GpuCount",
              "value": "0"
            },
            {
              "key": "Frequency",
              "value": "2.6GHz/3.3GHz"
            },
            {
              "key": "StatusCategory",
              "value": "WithoutStock"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "tencent",
          "cspImageName": "img-7rotv4ux",
          "regionList": [
            "ap-bangkok",
            "ap-beijing",
            "ap-chengdu",
            "ap-chongqing",
            "ap-guangzhou",
            "ap-hongkong",
            "ap-jakarta",
            "ap-nanjing",
            "ap-seoul",
            "ap-shanghai",
            "ap-singapore",
            "ap-tokyo",
            "eu-frankfurt",
            "me-saudi-arabia",
            "na-ashburn",
            "na-siliconvalley",
            "sa-saopaulo"
          ],
          "id": "img-7rotv4ux",
          "uid": "tbm5u74berjs1hfbpgol",
          "name": "img-7rotv4ux",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "tencent-sa-saopaulo",
          "infraType": "",
          "fetchedTime": "2026.08.21 13:57:36 Fri",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ImageId",
              "value": "img-7rotv4ux"
            },
            {
              "key": "OsName",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageType",
              "value": "PUBLIC_IMAGE"
            },
            {
              "key": "ImageName",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageDescription",
              "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
            },
            {
              "key": "ImageSize",
              "value": "20"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "ImageState",
              "value": "NORMAL"
            },
            {
              "key": "Platform",
              "value": "Ubuntu"
            },
            {
              "key": "ImageSource",
              "value": "OFFICIAL"
            },
            {
              "key": "IsSupportCloudinit",
              "value": "true"
            },
            {
              "key": "ImageDeprecated",
              "value": "false"
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
          "connectionName": "tencent-ap-seoul",
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
          "connectionName": "tencent-ap-seoul",
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
          "connectionName": "tencent-ap-seoul",
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
    "csp": "tencent",
    "region": "ap-seoul"
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
        "connectionName": "tencent-ap-seoul",
        "specId": "tencent+ap-seoul+bf1.medium2",
        "imageId": "img-7rotv4ux",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-01"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 50,
        "dataDiskIds": null
      },
      {
        "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "nodeGroupSize": 1,
        "label": {
          "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
        },
        "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
        "connectionName": "tencent-ap-seoul",
        "specId": "tencent+ap-seoul+bf1.large16",
        "imageId": "img-7rotv4ux",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-02"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 50,
        "dataDiskIds": null
      },
      {
        "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nodeGroupSize": 1,
        "label": {
          "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
        },
        "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
        "connectionName": "tencent-ap-seoul",
        "specId": "tencent+ap-seoul+bf1.medium8",
        "imageId": "img-7rotv4ux",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-03"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 50,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "vnet-01",
    "connectionName": "tencent-ap-seoul",
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
    "connectionName": "tencent-ap-seoul",
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
      "id": "tencent+ap-seoul+bf1.medium2",
      "uid": "tb71icnf2ipgc6r7vlil",
      "cspSpecName": "BF1.MEDIUM2",
      "name": "tencent+ap-seoul+bf1.medium2",
      "namespace": "system",
      "connectionName": "tencent-ap-seoul",
      "providerName": "tencent",
      "regionName": "ap-seoul",
      "regionLatitude": 37.566536,
      "regionLongitude": 126.977966,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 2,
      "diskSizeGB": -1,
      "costPerHour": 0.02,
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
          "key": "Zone",
          "value": "ap-seoul-1"
        },
        {
          "key": "InstanceType",
          "value": "BF1.MEDIUM2"
        },
        {
          "key": "InstanceChargeType",
          "value": "POSTPAID_BY_HOUR"
        },
        {
          "key": "NetworkCard",
          "value": "100"
        },
        {
          "key": "Externals",
          "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "Memory",
          "value": "2"
        },
        {
          "key": "InstanceFamily",
          "value": "BF1"
        },
        {
          "key": "TypeName",
          "value": "BF1"
        },
        {
          "key": "Status",
          "value": "SELL"
        },
        {
          "key": "Price",
          "value": "{UnitPrice:0.02,ChargeUnit:HOUR,Discount:100,UnitPriceDiscount:0.02,UnitPriceSecondStep:0.02,UnitPriceDiscountSecondStep:0.02,UnitPriceThirdStep:0.02,UnitPriceDiscountThirdStep:0.02}"
        },
        {
          "key": "InstanceBandwidth",
          "value": "1.5"
        },
        {
          "key": "InstancePps",
          "value": "25"
        },
        {
          "key": "StorageBlockAmount",
          "value": "0"
        },
        {
          "key": "CpuType",
          "value": "-"
        },
        {
          "key": "Gpu",
          "value": "0"
        },
        {
          "key": "Fpga",
          "value": "0"
        },
        {
          "key": "GpuCount",
          "value": "0"
        },
        {
          "key": "Frequency",
          "value": "-"
        },
        {
          "key": "StatusCategory",
          "value": "UnderStock"
        }
      ]
    },
    {
      "id": "tencent+ap-seoul+bf1.large16",
      "uid": "tb1cmris4eu56fa89t2k",
      "cspSpecName": "BF1.LARGE16",
      "name": "tencent+ap-seoul+bf1.large16",
      "namespace": "system",
      "connectionName": "tencent-ap-seoul",
      "providerName": "tencent",
      "regionName": "ap-seoul",
      "regionLatitude": 37.566536,
      "regionLongitude": 126.977966,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 16,
      "diskSizeGB": -1,
      "costPerHour": 0.15,
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
          "key": "Zone",
          "value": "ap-seoul-1"
        },
        {
          "key": "InstanceType",
          "value": "BF1.LARGE16"
        },
        {
          "key": "InstanceChargeType",
          "value": "POSTPAID_BY_HOUR"
        },
        {
          "key": "NetworkCard",
          "value": "100"
        },
        {
          "key": "Externals",
          "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
        },
        {
          "key": "Cpu",
          "value": "4"
        },
        {
          "key": "Memory",
          "value": "16"
        },
        {
          "key": "InstanceFamily",
          "value": "BF1"
        },
        {
          "key": "TypeName",
          "value": "BF1"
        },
        {
          "key": "Status",
          "value": "SOLD_OUT"
        },
        {
          "key": "Price",
          "value": "{UnitPrice:0.15,ChargeUnit:HOUR,Discount:100,UnitPriceDiscount:0.15,UnitPriceSecondStep:0.15,UnitPriceDiscountSecondStep:0.15,UnitPriceThirdStep:0.15,UnitPriceDiscountThirdStep:0.15}"
        },
        {
          "key": "SoldOutReason",
          "value": "ResourcesSoldOut.SpecifiedInstanceType"
        },
        {
          "key": "InstanceBandwidth",
          "value": "1.5"
        },
        {
          "key": "InstancePps",
          "value": "25"
        },
        {
          "key": "StorageBlockAmount",
          "value": "0"
        },
        {
          "key": "CpuType",
          "value": "-"
        },
        {
          "key": "Gpu",
          "value": "0"
        },
        {
          "key": "Fpga",
          "value": "0"
        },
        {
          "key": "GpuCount",
          "value": "0"
        },
        {
          "key": "Frequency",
          "value": "-"
        },
        {
          "key": "StatusCategory",
          "value": "WithoutStock"
        }
      ]
    },
    {
      "id": "tencent+ap-seoul+bf1.medium8",
      "uid": "tb7duhhnsa5d7414bjdn",
      "cspSpecName": "BF1.MEDIUM8",
      "name": "tencent+ap-seoul+bf1.medium8",
      "namespace": "system",
      "connectionName": "tencent-ap-seoul",
      "providerName": "tencent",
      "regionName": "ap-seoul",
      "regionLatitude": 37.566536,
      "regionLongitude": 126.977966,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 8,
      "diskSizeGB": -1,
      "costPerHour": 0.07,
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
          "key": "Zone",
          "value": "ap-seoul-1"
        },
        {
          "key": "InstanceType",
          "value": "BF1.MEDIUM8"
        },
        {
          "key": "InstanceChargeType",
          "value": "POSTPAID_BY_HOUR"
        },
        {
          "key": "NetworkCard",
          "value": "100"
        },
        {
          "key": "Externals",
          "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "Memory",
          "value": "8"
        },
        {
          "key": "InstanceFamily",
          "value": "BF1"
        },
        {
          "key": "TypeName",
          "value": "BF1"
        },
        {
          "key": "Status",
          "value": "SOLD_OUT"
        },
        {
          "key": "Price",
          "value": "{UnitPrice:0.07,ChargeUnit:HOUR,Discount:100,UnitPriceDiscount:0.07,UnitPriceSecondStep:0.07,UnitPriceDiscountSecondStep:0.07,UnitPriceThirdStep:0.07,UnitPriceDiscountThirdStep:0.07}"
        },
        {
          "key": "SoldOutReason",
          "value": "ResourcesSoldOut.SpecifiedInstanceType"
        },
        {
          "key": "InstanceBandwidth",
          "value": "1.5"
        },
        {
          "key": "InstancePps",
          "value": "25"
        },
        {
          "key": "StorageBlockAmount",
          "value": "0"
        },
        {
          "key": "CpuType",
          "value": "-"
        },
        {
          "key": "Gpu",
          "value": "0"
        },
        {
          "key": "Fpga",
          "value": "0"
        },
        {
          "key": "GpuCount",
          "value": "0"
        },
        {
          "key": "Frequency",
          "value": "-"
        },
        {
          "key": "StatusCategory",
          "value": "WithoutStock"
        }
      ]
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "tencent",
      "cspImageName": "img-7rotv4ux",
      "regionList": [
        "ap-bangkok",
        "ap-beijing",
        "ap-chengdu",
        "ap-chongqing",
        "ap-guangzhou",
        "ap-hongkong",
        "ap-jakarta",
        "ap-nanjing",
        "ap-seoul",
        "ap-shanghai",
        "ap-singapore",
        "ap-tokyo",
        "eu-frankfurt",
        "me-saudi-arabia",
        "na-ashburn",
        "na-siliconvalley",
        "sa-saopaulo"
      ],
      "id": "img-7rotv4ux",
      "uid": "tbm5u74berjs1hfbpgol",
      "name": "img-7rotv4ux",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "tencent-sa-saopaulo",
      "infraType": "",
      "fetchedTime": "2026.08.21 13:57:36 Fri",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI",
      "osDiskType": "NA",
      "osDiskSizeGB": 20,
      "imageStatus": "Available",
      "details": [
        {
          "key": "ImageId",
          "value": "img-7rotv4ux"
        },
        {
          "key": "OsName",
          "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
        },
        {
          "key": "ImageType",
          "value": "PUBLIC_IMAGE"
        },
        {
          "key": "ImageName",
          "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
        },
        {
          "key": "ImageDescription",
          "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
        },
        {
          "key": "ImageSize",
          "value": "20"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "ImageState",
          "value": "NORMAL"
        },
        {
          "key": "Platform",
          "value": "Ubuntu"
        },
        {
          "key": "ImageSource",
          "value": "OFFICIAL"
        },
        {
          "key": "IsSupportCloudinit",
          "value": "true"
        },
        {
          "key": "ImageDeprecated",
          "value": "false"
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
      "connectionName": "tencent-ap-seoul",
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
      "connectionName": "tencent-ap-seoul",
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
      "connectionName": "tencent-ap-seoul",
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
  "id": "my05-infra101",
  "uid": "tb77o0p6kep6u9dv9qk2",
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
  "label": {
    "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
    "sys.id": "my05-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my05-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tb77o0p6kep6u9dv9qk2"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbrl352l031jcihidb83",
      "cspResourceName": "tbrl352l031jcihidb83",
      "cspResourceId": "ins-pjjkbdwz",
      "name": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.566536,
        "longitude": 126.977966
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-09-21 04:19:24",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "tencent-ap-seoul",
        "sys.createdTime": "2026-09-21 04:19:24",
        "sys.cspResourceId": "ins-pjjkbdwz",
        "sys.cspResourceName": "tbrl352l031jcihidb83",
        "sys.id": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my05-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my05-subnet-01",
        "sys.uid": "tbrl352l031jcihidb83",
        "sys.vNetId": "my05-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.155.208.156",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.13",
      "privateDNS": "",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 50,
      "RootDeviceName": "disk-54u70mk7",
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
      "specId": "tencent+ap-seoul+bf1.medium2",
      "cspSpecName": "BF1.MEDIUM2",
      "spec": {
        "cspSpecName": "BF1.MEDIUM2",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.02
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
      "vNetId": "my05-vnet-01",
      "cspVNetId": "vpc-ktohe25c",
      "subnetId": "my05-subnet-01",
      "cspSubnetId": "subnet-0ilffklb",
      "networkInterface": "eni-d7379z5y",
      "securityGroupIds": [
        "my05-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-sshkey-01",
      "cspSshKeyId": "skey-c8que8al",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI4o6uxQe4mvKQflwn1ZIy/flZYOuxFpx0lHZNQFfKR0DBy8o9bt9EF9EbWUtSKGzx3eRaqYoeYo0d+mrf0OSrk=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:SSV7O87ajeoDGV3eWELkIXD5DSSXj9lvQ1i+HSS1RkQ",
        "firstUsedAt": "2026-09-21T04:19:42Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:42Z",
          "completedTime": "2026-09-21T04:19:46Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:46Z",
          "completedTime": "2026-09-21T04:19:47Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-13-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tbh30hjqjpieg0073p53",
      "cspResourceName": "tbh30hjqjpieg0073p53",
      "cspResourceId": "ins-pl7luz01",
      "name": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.566536,
        "longitude": 126.977966
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-09-21 04:19:29",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "tencent-ap-seoul",
        "sys.createdTime": "2026-09-21 04:19:29",
        "sys.cspResourceId": "ins-pl7luz01",
        "sys.cspResourceName": "tbh30hjqjpieg0073p53",
        "sys.id": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my05-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my05-subnet-01",
        "sys.uid": "tbh30hjqjpieg0073p53",
        "sys.vNetId": "my05-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.155.222.67",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.3",
      "privateDNS": "",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 50,
      "RootDeviceName": "disk-9i43ufjl",
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
      "specId": "tencent+ap-seoul+bf1.medium8",
      "cspSpecName": "BF1.MEDIUM8",
      "spec": {
        "cspSpecName": "BF1.MEDIUM8",
        "vCPU": 2,
        "memoryGiB": 8,
        "costPerHour": 0.07
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
      "vNetId": "my05-vnet-01",
      "cspVNetId": "vpc-ktohe25c",
      "subnetId": "my05-subnet-01",
      "cspSubnetId": "subnet-0ilffklb",
      "networkInterface": "eni-c2nu0hpi",
      "securityGroupIds": [
        "my05-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-sshkey-01",
      "cspSshKeyId": "skey-c8que8al",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAVpofUAPI/rXLYNg8O5zsSVVhjWp6WRw6i+xtXIU8TiVT/E1vDWxYzWyEN+vdzkb+PneKiiFDCYpl7+QiSgdzI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:M4JVB/FnZ5W/EbZfkxvCEKTqtuVlIsOWNcrNzn7TEDA",
        "firstUsedAt": "2026-09-21T04:19:44Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:42Z",
          "completedTime": "2026-09-21T04:19:45Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:46Z",
          "completedTime": "2026-09-21T04:19:47Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-3-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tb5q8f0aukaedjc42fbi",
      "cspResourceName": "tb5q8f0aukaedjc42fbi",
      "cspResourceId": "ins-olowx4jd",
      "name": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.566536,
        "longitude": 126.977966
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-09-21 04:19:29",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "tencent-ap-seoul",
        "sys.createdTime": "2026-09-21 04:19:29",
        "sys.cspResourceId": "ins-olowx4jd",
        "sys.cspResourceName": "tb5q8f0aukaedjc42fbi",
        "sys.id": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my05-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my05-subnet-01",
        "sys.uid": "tb5q8f0aukaedjc42fbi",
        "sys.vNetId": "my05-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "119.28.150.173",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.12",
      "privateDNS": "",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 50,
      "RootDeviceName": "disk-nz94tgkf",
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
      "specId": "tencent+ap-seoul+bf1.large16",
      "cspSpecName": "BF1.LARGE16",
      "spec": {
        "cspSpecName": "BF1.LARGE16",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.15
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
      "vNetId": "my05-vnet-01",
      "cspVNetId": "vpc-ktohe25c",
      "subnetId": "my05-subnet-01",
      "cspSubnetId": "subnet-0ilffklb",
      "networkInterface": "eni-d6t78qnc",
      "securityGroupIds": [
        "my05-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-sshkey-01",
      "cspSshKeyId": "skey-c8que8al",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBONF4QzfiQtnXjCjJK8dVCfisxuF5quDUCul3SPpSxInXGNITh+5uX3BCQeJovhEzXxw3cKG+zKvMBt8A6gdabA=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:jYL32RRu0DRJgb7qMEbLKo+uatySuzBEEW2EhGXHndU",
        "firstUsedAt": "2026-09-21T04:19:44Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:42Z",
          "completedTime": "2026-09-21T04:19:45Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:46Z",
          "completedTime": "2026-09-21T04:19:47Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-12-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my05-vnet-01",
      "name": "my05-vnet-01",
      "infraId": "my05-infra101",
      "vNetId": "my05-vnet-01",
      "connectionNames": [
        "tencent-ap-seoul"
      ],
      "providerNames": [
        "tencent"
      ],
      "regionNames": [
        "ap-seoul"
      ],
      "nodeGroupIds": [
        "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932"
      ],
      "nodeIds": [
        "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1"
      ],
      "nodeGroupCount": 3,
      "nodeCount": 3,
      "representativeNodeGroupId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my05-infra101",
            "nodeId": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "119.28.150.173",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-12-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "43.155.222.67",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-3-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "43.155.208.156",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-13-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r"
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
      "id": "my05-infra101",
      "uid": "tb77o0p6kep6u9dv9qk2",
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
      "label": {
        "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
        "sys.id": "my05-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tb77o0p6kep6u9dv9qk2"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbrl352l031jcihidb83",
          "cspResourceName": "tbrl352l031jcihidb83",
          "cspResourceId": "ins-pjjkbdwz",
          "name": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.155.208.156",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.13",
          "privateDNS": "",
          "rootDiskType": "CLOUD_PREMIUM",
          "rootDiskSize": 50,
          "RootDeviceName": "",
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
          "specId": "tencent+ap-seoul+bf1.medium2",
          "cspSpecName": "BF1.MEDIUM2",
          "spec": {
            "cspSpecName": "BF1.MEDIUM2",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.02
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
          "vNetId": "my05-vnet-01",
          "cspVNetId": "vpc-ktohe25c",
          "subnetId": "my05-subnet-01",
          "cspSubnetId": "subnet-0ilffklb",
          "networkInterface": "eni-d7379z5y",
          "securityGroupIds": [
            "my05-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-sshkey-01",
          "cspSshKeyId": "skey-c8que8al"
        },
        {
          "resourceType": "node",
          "id": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tbh30hjqjpieg0073p53",
          "cspResourceName": "tbh30hjqjpieg0073p53",
          "cspResourceId": "ins-pl7luz01",
          "name": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "43.155.222.67",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.3",
          "privateDNS": "",
          "rootDiskType": "CLOUD_PREMIUM",
          "rootDiskSize": 50,
          "RootDeviceName": "",
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
          "specId": "tencent+ap-seoul+bf1.medium8",
          "cspSpecName": "BF1.MEDIUM8",
          "spec": {
            "cspSpecName": "BF1.MEDIUM8",
            "vCPU": 2,
            "memoryGiB": 8,
            "costPerHour": 0.07
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
          "vNetId": "my05-vnet-01",
          "cspVNetId": "vpc-ktohe25c",
          "subnetId": "my05-subnet-01",
          "cspSubnetId": "subnet-0ilffklb",
          "networkInterface": "eni-c2nu0hpi",
          "securityGroupIds": [
            "my05-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-sshkey-01",
          "cspSshKeyId": "skey-c8que8al"
        },
        {
          "resourceType": "node",
          "id": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tb5q8f0aukaedjc42fbi",
          "cspResourceName": "tb5q8f0aukaedjc42fbi",
          "cspResourceId": "ins-olowx4jd",
          "name": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "Running==Running",
          "createdTime": "",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
          },
          "description": "",
          "region": {
            "region": "ap-seoul",
            "zone": "ap-seoul-1"
          },
          "publicIP": "119.28.150.173",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.12",
          "privateDNS": "",
          "rootDiskType": "CLOUD_PREMIUM",
          "rootDiskSize": 50,
          "RootDeviceName": "",
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
          "specId": "tencent+ap-seoul+bf1.large16",
          "cspSpecName": "BF1.LARGE16",
          "spec": {
            "cspSpecName": "BF1.LARGE16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.15
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
          "vNetId": "my05-vnet-01",
          "cspVNetId": "vpc-ktohe25c",
          "subnetId": "my05-subnet-01",
          "cspSubnetId": "subnet-0ilffklb",
          "networkInterface": "eni-d6t78qnc",
          "securityGroupIds": [
            "my05-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my05-sshkey-01",
          "cspSshKeyId": "skey-c8que8al"
        }
      ],
      "newNodeList": null
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
    "my05-infra101"
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
  "id": "my05-infra101",
  "uid": "tb77o0p6kep6u9dv9qk2",
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
  "label": {
    "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
    "sys.id": "my05-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my05-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tb77o0p6kep6u9dv9qk2"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbrl352l031jcihidb83",
      "cspResourceName": "tbrl352l031jcihidb83",
      "cspResourceId": "ins-pjjkbdwz",
      "name": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.566536,
        "longitude": 126.977966
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-09-21 04:19:24",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "tencent-ap-seoul",
        "sys.createdTime": "2026-09-21 04:19:24",
        "sys.cspResourceId": "ins-pjjkbdwz",
        "sys.cspResourceName": "tbrl352l031jcihidb83",
        "sys.id": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my05-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my05-subnet-01",
        "sys.uid": "tbrl352l031jcihidb83",
        "sys.vNetId": "my05-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.155.208.156",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.13",
      "privateDNS": "",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 50,
      "RootDeviceName": "disk-54u70mk7",
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
      "specId": "tencent+ap-seoul+bf1.medium2",
      "cspSpecName": "BF1.MEDIUM2",
      "spec": {
        "cspSpecName": "BF1.MEDIUM2",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.02
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
      "vNetId": "my05-vnet-01",
      "cspVNetId": "vpc-ktohe25c",
      "subnetId": "my05-subnet-01",
      "cspSubnetId": "subnet-0ilffklb",
      "networkInterface": "eni-d7379z5y",
      "securityGroupIds": [
        "my05-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-sshkey-01",
      "cspSshKeyId": "skey-c8que8al",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI4o6uxQe4mvKQflwn1ZIy/flZYOuxFpx0lHZNQFfKR0DBy8o9bt9EF9EbWUtSKGzx3eRaqYoeYo0d+mrf0OSrk=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:SSV7O87ajeoDGV3eWELkIXD5DSSXj9lvQ1i+HSS1RkQ",
        "firstUsedAt": "2026-09-21T04:19:42Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:42Z",
          "completedTime": "2026-09-21T04:19:46Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:46Z",
          "completedTime": "2026-09-21T04:19:47Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-13-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tbh30hjqjpieg0073p53",
      "cspResourceName": "tbh30hjqjpieg0073p53",
      "cspResourceId": "ins-pl7luz01",
      "name": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.566536,
        "longitude": 126.977966
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-09-21 04:19:29",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "tencent-ap-seoul",
        "sys.createdTime": "2026-09-21 04:19:29",
        "sys.cspResourceId": "ins-pl7luz01",
        "sys.cspResourceName": "tbh30hjqjpieg0073p53",
        "sys.id": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my05-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my05-subnet-01",
        "sys.uid": "tbh30hjqjpieg0073p53",
        "sys.vNetId": "my05-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "43.155.222.67",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.3",
      "privateDNS": "",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 50,
      "RootDeviceName": "disk-9i43ufjl",
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
      "specId": "tencent+ap-seoul+bf1.medium8",
      "cspSpecName": "BF1.MEDIUM8",
      "spec": {
        "cspSpecName": "BF1.MEDIUM8",
        "vCPU": 2,
        "memoryGiB": 8,
        "costPerHour": 0.07
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
      "vNetId": "my05-vnet-01",
      "cspVNetId": "vpc-ktohe25c",
      "subnetId": "my05-subnet-01",
      "cspSubnetId": "subnet-0ilffklb",
      "networkInterface": "eni-c2nu0hpi",
      "securityGroupIds": [
        "my05-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-sshkey-01",
      "cspSshKeyId": "skey-c8que8al",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAVpofUAPI/rXLYNg8O5zsSVVhjWp6WRw6i+xtXIU8TiVT/E1vDWxYzWyEN+vdzkb+PneKiiFDCYpl7+QiSgdzI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:M4JVB/FnZ5W/EbZfkxvCEKTqtuVlIsOWNcrNzn7TEDA",
        "firstUsedAt": "2026-09-21T04:19:44Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:42Z",
          "completedTime": "2026-09-21T04:19:45Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:46Z",
          "completedTime": "2026-09-21T04:19:47Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-3-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tb5q8f0aukaedjc42fbi",
      "cspResourceName": "tb5q8f0aukaedjc42fbi",
      "cspResourceId": "ins-olowx4jd",
      "name": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.566536,
        "longitude": 126.977966
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-09-21 04:19:29",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "tencent-ap-seoul",
        "sys.createdTime": "2026-09-21 04:19:29",
        "sys.cspResourceId": "ins-olowx4jd",
        "sys.cspResourceName": "tb5q8f0aukaedjc42fbi",
        "sys.id": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my05-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my05-subnet-01",
        "sys.uid": "tb5q8f0aukaedjc42fbi",
        "sys.vNetId": "my05-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-seoul",
        "zone": "ap-seoul-1"
      },
      "publicIP": "119.28.150.173",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.12",
      "privateDNS": "",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 50,
      "RootDeviceName": "disk-nz94tgkf",
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
      "specId": "tencent+ap-seoul+bf1.large16",
      "cspSpecName": "BF1.LARGE16",
      "spec": {
        "cspSpecName": "BF1.LARGE16",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.15
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
      "vNetId": "my05-vnet-01",
      "cspVNetId": "vpc-ktohe25c",
      "subnetId": "my05-subnet-01",
      "cspSubnetId": "subnet-0ilffklb",
      "networkInterface": "eni-d6t78qnc",
      "securityGroupIds": [
        "my05-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my05-sshkey-01",
      "cspSshKeyId": "skey-c8que8al",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBONF4QzfiQtnXjCjJK8dVCfisxuF5quDUCul3SPpSxInXGNITh+5uX3BCQeJovhEzXxw3cKG+zKvMBt8A6gdabA=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:jYL32RRu0DRJgb7qMEbLKo+uatySuzBEEW2EhGXHndU",
        "firstUsedAt": "2026-09-21T04:19:44Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "true",
          "commandExecuted": "true",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:42Z",
          "completedTime": "2026-09-21T04:19:45Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "\n",
          "stderr": "\n"
        },
        {
          "index": 2,
          "xRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r",
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-09-21T04:19:46Z",
          "completedTime": "2026-09-21T04:19:47Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux VM-1-12-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ]
    }
  ],
  "cluster": [
    {
      "id": "my05-vnet-01",
      "name": "my05-vnet-01",
      "infraId": "my05-infra101",
      "vNetId": "my05-vnet-01",
      "connectionNames": [
        "tencent-ap-seoul"
      ],
      "providerNames": [
        "tencent"
      ],
      "regionNames": [
        "ap-seoul"
      ],
      "nodeGroupIds": [
        "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932"
      ],
      "nodeIds": [
        "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1"
      ],
      "nodeGroupCount": 3,
      "nodeCount": 3,
      "representativeNodeGroupId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "representativeNodeId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1"
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
            "infraId": "my05-infra101",
            "nodeId": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "119.28.150.173",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-12-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "43.155.222.67",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-3-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "error": ""
          },
          {
            "infraId": "my05-infra101",
            "nodeId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "43.155.208.156",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux VM-1-13-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "postCommandRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r"
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
    "cluster": [
      {
        "connectionNames": [
          "tencent-ap-seoul"
        ],
        "id": "my05-vnet-01",
        "infraId": "my05-infra101",
        "name": "my05-vnet-01",
        "nodeCount": 3,
        "nodeGroupCount": 3,
        "nodeGroupIds": [
          "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932"
        ],
        "nodeIds": [
          "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1"
        ],
        "providerNames": [
          "tencent"
        ],
        "regionNames": [
          "ap-seoul"
        ],
        "representativeNodeGroupId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "representativeNodeId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vNetId": "my05-vnet-01"
      }
    ],
    "configureCloudAdaptiveNetwork": "",
    "description": "Recommended VMs comprising multi-cloud infrastructure",
    "id": "my05-infra101",
    "installMonAgent": "",
    "label": {
      "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
      "sys.id": "my05-infra101",
      "sys.labelType": "infra",
      "sys.manager": "cb-tumblebug",
      "sys.name": "my05-infra101",
      "sys.namespace": "mig01",
      "sys.uid": "tb77o0p6kep6u9dv9qk2"
    },
    "name": "my05-infra101",
    "newNodeList": null,
    "node": [
      {
        "RootDeviceName": "disk-54u70mk7",
        "commandStatus": [
          {
            "commandExecuted": "true",
            "commandRequested": "true",
            "completedTime": "2026-09-21T04:19:46Z",
            "elapsedTime": 4,
            "index": 1,
            "resultSummary": "Command executed successfully",
            "startedTime": "2026-09-21T04:19:42Z",
            "status": "Completed",
            "stderr": "\n",
            "stdout": "\n"
          },
          {
            "commandExecuted": "uname -a",
            "commandRequested": "uname -a",
            "completedTime": "2026-09-21T04:19:47Z",
            "elapsedTime": 1,
            "index": 2,
            "resultSummary": "Command executed successfully",
            "startedTime": "2026-09-21T04:19:46Z",
            "status": "Completed",
            "stderr": "\n",
            "stdout": "Linux VM-1-13-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
            "xRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r"
          }
        ],
        "connectionConfig": {
          "configName": "tencent-ap-seoul",
          "credentialHolder": "admin",
          "credentialName": "tencent",
          "driverName": "tencent-driver-v1.0.so",
          "providerName": "tencent",
          "regionDetail": {
            "description": "Seoul",
            "location": {
              "display": "South Korea (Seoul)",
              "latitude": 37.566536,
              "longitude": 126.977966
            },
            "regionId": "ap-seoul",
            "regionName": "ap-seoul",
            "zones": [
              "ap-seoul-1",
              "ap-seoul-2"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "ap-seoul",
            "assignedZone": "ap-seoul-1"
          },
          "regionZoneInfoName": "tencent-ap-seoul",
          "verified": true
        },
        "connectionName": "tencent-ap-seoul",
        "createdTime": "2026-09-21 04:19:24",
        "cspImageName": "img-7rotv4ux",
        "cspResourceId": "ins-pjjkbdwz",
        "cspResourceName": "tbrl352l031jcihidb83",
        "cspSpecName": "BF1.MEDIUM2",
        "cspSshKeyId": "skey-c8que8al",
        "cspSubnetId": "subnet-0ilffklb",
        "cspVNetId": "vpc-ktohe25c",
        "dataDiskIds": null,
        "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
        "id": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "image": {
          "cspImageName": "img-7rotv4ux",
          "osArchitecture": "x86_64",
          "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "img-7rotv4ux",
        "label": {
          "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
          "sys.connectionName": "tencent-ap-seoul",
          "sys.createdTime": "2026-09-21 04:19:24",
          "sys.cspResourceId": "ins-pjjkbdwz",
          "sys.cspResourceName": "tbrl352l031jcihidb83",
          "sys.id": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "sys.infraId": "my05-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "sys.subnetId": "my05-subnet-01",
          "sys.uid": "tbrl352l031jcihidb83",
          "sys.vNetId": "my05-vnet-01"
        },
        "location": {
          "display": "South Korea (Seoul)",
          "latitude": 37.566536,
          "longitude": 126.977966
        },
        "monAgentStatus": "notInstalled",
        "name": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "eni-d7379z5y",
        "nodeGroupId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.13",
        "publicDNS": "",
        "publicIP": "43.155.208.156",
        "region": {
          "region": "ap-seoul",
          "zone": "ap-seoul-1"
        },
        "resourceType": "node",
        "rootDiskSize": 50,
        "rootDiskType": "CLOUD_PREMIUM",
        "securityGroupIds": [
          "my05-sg-01"
        ],
        "spec": {
          "costPerHour": 0.02,
          "cspSpecName": "BF1.MEDIUM2",
          "memoryGiB": 2,
          "vCPU": 2
        },
        "specId": "tencent+ap-seoul+bf1.medium2",
        "sshHostKeyInfo": {
          "fingerprint": "SHA256:SSV7O87ajeoDGV3eWELkIXD5DSSXj9lvQ1i+HSS1RkQ",
          "firstUsedAt": "2026-09-21T04:19:42Z",
          "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI4o6uxQe4mvKQflwn1ZIy/flZYOuxFpx0lHZNQFfKR0DBy8o9bt9EF9EbWUtSKGzx3eRaqYoeYo0d+mrf0OSrk=",
          "keyType": "ecdsa-sha2-nistp256"
        },
        "sshKeyId": "my05-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my05-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tbrl352l031jcihidb83",
        "vNetId": "my05-vnet-01"
      },
      {
        "RootDeviceName": "disk-9i43ufjl",
        "commandStatus": [
          {
            "commandExecuted": "true",
            "commandRequested": "true",
            "completedTime": "2026-09-21T04:19:45Z",
            "elapsedTime": 3,
            "index": 1,
            "resultSummary": "Command executed successfully",
            "startedTime": "2026-09-21T04:19:42Z",
            "status": "Completed",
            "stderr": "\n",
            "stdout": "\n"
          },
          {
            "commandExecuted": "uname -a",
            "commandRequested": "uname -a",
            "completedTime": "2026-09-21T04:19:47Z",
            "elapsedTime": 1,
            "index": 2,
            "resultSummary": "Command executed successfully",
            "startedTime": "2026-09-21T04:19:46Z",
            "status": "Completed",
            "stderr": "\n",
            "stdout": "Linux VM-1-3-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
            "xRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r"
          }
        ],
        "connectionConfig": {
          "configName": "tencent-ap-seoul",
          "credentialHolder": "admin",
          "credentialName": "tencent",
          "driverName": "tencent-driver-v1.0.so",
          "providerName": "tencent",
          "regionDetail": {
            "description": "Seoul",
            "location": {
              "display": "South Korea (Seoul)",
              "latitude": 37.566536,
              "longitude": 126.977966
            },
            "regionId": "ap-seoul",
            "regionName": "ap-seoul",
            "zones": [
              "ap-seoul-1",
              "ap-seoul-2"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "ap-seoul",
            "assignedZone": "ap-seoul-1"
          },
          "regionZoneInfoName": "tencent-ap-seoul",
          "verified": true
        },
        "connectionName": "tencent-ap-seoul",
        "createdTime": "2026-09-21 04:19:29",
        "cspImageName": "img-7rotv4ux",
        "cspResourceId": "ins-pl7luz01",
        "cspResourceName": "tbh30hjqjpieg0073p53",
        "cspSpecName": "BF1.MEDIUM8",
        "cspSshKeyId": "skey-c8que8al",
        "cspSubnetId": "subnet-0ilffklb",
        "cspVNetId": "vpc-ktohe25c",
        "dataDiskIds": null,
        "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
        "id": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "image": {
          "cspImageName": "img-7rotv4ux",
          "osArchitecture": "x86_64",
          "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "img-7rotv4ux",
        "label": {
          "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "sys.connectionName": "tencent-ap-seoul",
          "sys.createdTime": "2026-09-21 04:19:29",
          "sys.cspResourceId": "ins-pl7luz01",
          "sys.cspResourceName": "tbh30hjqjpieg0073p53",
          "sys.id": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "sys.infraId": "my05-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "sys.subnetId": "my05-subnet-01",
          "sys.uid": "tbh30hjqjpieg0073p53",
          "sys.vNetId": "my05-vnet-01"
        },
        "location": {
          "display": "South Korea (Seoul)",
          "latitude": 37.566536,
          "longitude": 126.977966
        },
        "monAgentStatus": "notInstalled",
        "name": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "eni-c2nu0hpi",
        "nodeGroupId": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.3",
        "publicDNS": "",
        "publicIP": "43.155.222.67",
        "region": {
          "region": "ap-seoul",
          "zone": "ap-seoul-1"
        },
        "resourceType": "node",
        "rootDiskSize": 50,
        "rootDiskType": "CLOUD_PREMIUM",
        "securityGroupIds": [
          "my05-sg-03"
        ],
        "spec": {
          "costPerHour": 0.07,
          "cspSpecName": "BF1.MEDIUM8",
          "memoryGiB": 8,
          "vCPU": 2
        },
        "specId": "tencent+ap-seoul+bf1.medium8",
        "sshHostKeyInfo": {
          "fingerprint": "SHA256:M4JVB/FnZ5W/EbZfkxvCEKTqtuVlIsOWNcrNzn7TEDA",
          "firstUsedAt": "2026-09-21T04:19:44Z",
          "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAVpofUAPI/rXLYNg8O5zsSVVhjWp6WRw6i+xtXIU8TiVT/E1vDWxYzWyEN+vdzkb+PneKiiFDCYpl7+QiSgdzI=",
          "keyType": "ecdsa-sha2-nistp256"
        },
        "sshKeyId": "my05-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my05-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tbh30hjqjpieg0073p53",
        "vNetId": "my05-vnet-01"
      },
      {
        "RootDeviceName": "disk-nz94tgkf",
        "commandStatus": [
          {
            "commandExecuted": "true",
            "commandRequested": "true",
            "completedTime": "2026-09-21T04:19:45Z",
            "elapsedTime": 3,
            "index": 1,
            "resultSummary": "Command executed successfully",
            "startedTime": "2026-09-21T04:19:42Z",
            "status": "Completed",
            "stderr": "\n",
            "stdout": "\n"
          },
          {
            "commandExecuted": "uname -a",
            "commandRequested": "uname -a",
            "completedTime": "2026-09-21T04:19:47Z",
            "elapsedTime": 1,
            "index": 2,
            "resultSummary": "Command executed successfully",
            "startedTime": "2026-09-21T04:19:46Z",
            "status": "Completed",
            "stderr": "\n",
            "stdout": "Linux VM-1-12-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n\n",
            "xRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r"
          }
        ],
        "connectionConfig": {
          "configName": "tencent-ap-seoul",
          "credentialHolder": "admin",
          "credentialName": "tencent",
          "driverName": "tencent-driver-v1.0.so",
          "providerName": "tencent",
          "regionDetail": {
            "description": "Seoul",
            "location": {
              "display": "South Korea (Seoul)",
              "latitude": 37.566536,
              "longitude": 126.977966
            },
            "regionId": "ap-seoul",
            "regionName": "ap-seoul",
            "zones": [
              "ap-seoul-1",
              "ap-seoul-2"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "ap-seoul",
            "assignedZone": "ap-seoul-1"
          },
          "regionZoneInfoName": "tencent-ap-seoul",
          "verified": true
        },
        "connectionName": "tencent-ap-seoul",
        "createdTime": "2026-09-21 04:19:29",
        "cspImageName": "img-7rotv4ux",
        "cspResourceId": "ins-olowx4jd",
        "cspResourceName": "tb5q8f0aukaedjc42fbi",
        "cspSpecName": "BF1.LARGE16",
        "cspSshKeyId": "skey-c8que8al",
        "cspSubnetId": "subnet-0ilffklb",
        "cspVNetId": "vpc-ktohe25c",
        "dataDiskIds": null,
        "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
        "id": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "image": {
          "cspImageName": "img-7rotv4ux",
          "osArchitecture": "x86_64",
          "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "img-7rotv4ux",
        "label": {
          "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "sys.connectionName": "tencent-ap-seoul",
          "sys.createdTime": "2026-09-21 04:19:29",
          "sys.cspResourceId": "ins-olowx4jd",
          "sys.cspResourceName": "tb5q8f0aukaedjc42fbi",
          "sys.id": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "sys.infraId": "my05-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "sys.subnetId": "my05-subnet-01",
          "sys.uid": "tb5q8f0aukaedjc42fbi",
          "sys.vNetId": "my05-vnet-01"
        },
        "location": {
          "display": "South Korea (Seoul)",
          "latitude": 37.566536,
          "longitude": 126.977966
        },
        "monAgentStatus": "notInstalled",
        "name": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "eni-d6t78qnc",
        "nodeGroupId": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.12",
        "publicDNS": "",
        "publicIP": "119.28.150.173",
        "region": {
          "region": "ap-seoul",
          "zone": "ap-seoul-1"
        },
        "resourceType": "node",
        "rootDiskSize": 50,
        "rootDiskType": "CLOUD_PREMIUM",
        "securityGroupIds": [
          "my05-sg-02"
        ],
        "spec": {
          "costPerHour": 0.15,
          "cspSpecName": "BF1.LARGE16",
          "memoryGiB": 16,
          "vCPU": 4
        },
        "specId": "tencent+ap-seoul+bf1.large16",
        "sshHostKeyInfo": {
          "fingerprint": "SHA256:jYL32RRu0DRJgb7qMEbLKo+uatySuzBEEW2EhGXHndU",
          "firstUsedAt": "2026-09-21T04:19:44Z",
          "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBONF4QzfiQtnXjCjJK8dVCfisxuF5quDUCul3SPpSxInXGNITh+5uX3BCQeJovhEzXxw3cKG+zKvMBt8A6gdabA=",
          "keyType": "ecdsa-sha2-nistp256"
        },
        "sshKeyId": "my05-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my05-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tb5q8f0aukaedjc42fbi",
        "vNetId": "my05-vnet-01"
      }
    ],
    "postCommandRequestId": "pc-my05-infra101-tb2a44lh6devo243ip5r",
    "postCommandResults": [
      {
        "phase": 1,
        "results": {
          "results": [
            {
              "command": {
                "0": "uname -a"
              },
              "error": "",
              "infraId": "my05-infra101",
              "nodeId": "my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
              "nodeIp": "119.28.150.173",
              "stderr": {
                "0": ""
              },
              "stdout": {
                "0": "Linux VM-1-12-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
              }
            },
            {
              "command": {
                "0": "uname -a"
              },
              "error": "",
              "infraId": "my05-infra101",
              "nodeId": "my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
              "nodeIp": "43.155.222.67",
              "stderr": {
                "0": ""
              },
              "stdout": {
                "0": "Linux VM-1-3-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
              }
            },
            {
              "command": {
                "0": "uname -a"
              },
              "error": "",
              "infraId": "my05-infra101",
              "nodeId": "my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
              "nodeIp": "43.155.208.156",
              "stderr": {
                "0": ""
              },
              "stdout": {
                "0": "Linux VM-1-13-ubuntu 5.15.0-119-generic #129-Ubuntu SMP Fri Aug 2 19:25:20 UTC 2024 x86_64 x86_64 x86_64 GNU/Linux\n"
              }
            }
          ]
        },
        "status": "Completed",
        "target": "all nodes"
      }
    ],
    "postCommandStatus": "Completed",
    "postCommands": [
      {
        "command": [
          "uname -a"
        ],
        "userName": "cb-user"
      }
    ],
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
    "uid": "tb77o0p6kep6u9dv9qk2"
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

**Generated At:** 2026-09-21 04:21:08

**Namespace:** mig01

**Infra Name:** my05-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my05-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | TENCENT |
| **Target Region** | ap-seoul |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| BF1.MEDIUM2 | 2 | 2.0 | - | x86_64 |  | $0.0200 | 1 |
| BF1.MEDIUM8 | 2 | 8.0 | - | x86_64 |  | $0.0700 | 1 |
| BF1.LARGE16 | 4 | 16.0 | - | x86_64 |  | $0.1500 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| img-7rotv4ux | Ubuntu Server 22.04 LTS 64bit UEFI | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | ins-pjjkbdwz | Running | 2 vCPU, 2.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-vnet-01<br>**Subnet:** my05-subnet-01<br>**Public IP:** 43.155.208.156<br>**Private IP:** 10.0.1.13<br>**SGs:** my05-sg-01<br>**SSH:** my05-sshkey-01 |
| my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | ins-pl7luz01 | Running | 2 vCPU, 8.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-vnet-01<br>**Subnet:** my05-subnet-01<br>**Public IP:** 43.155.222.67<br>**Private IP:** 10.0.1.3<br>**SGs:** my05-sg-03<br>**SSH:** my05-sshkey-01 |
| my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | ins-olowx4jd | Running | 4 vCPU, 16.0 GiB | Ubuntu Server 22.04 LTS 64bit UEFI (Ubuntu Server 22.04 LTS 64bit UEFI) | **VNet:** my05-vnet-01<br>**Subnet:** my05-subnet-01<br>**Public IP:** 119.28.150.173<br>**Private IP:** 10.0.1.12<br>**SGs:** my05-sg-02<br>**SSH:** my05-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my05-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my05-vnet-01 |
| **CSP VNet ID** | vpc-ktohe25c |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | tencent-ap-seoul |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my05-subnet-01 | subnet-0ilffklb | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my05-sshkey-01 | skey-c8que8al |  |  |

### Security Groups

#### Security Group: my05-sg-01

| Property | Value |
|----------|-------|
| **Name** | my05-sg-01 |
| **CSP Security Group ID** | sg-m9z8xo0l |
| **VNet** | my05-vnet-01 |
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
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my05-sg-02

| Property | Value |
|----------|-------|
| **Name** | my05-sg-02 |
| **CSP Security Group ID** | sg-3z1hp8d1 |
| **VNet** | my05-vnet-01 |
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
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my05-sg-03

| Property | Value |
|----------|-------|
| **Name** | my05-sg-03 |
| **CSP Security Group ID** | sg-4au7mzat |
| **VNet** | my05-vnet-01 |
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
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.2400 |
| **Per Day** | $5.76 |
| **Per Month (30 days)** | $172.80 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| TENCENT | ap-seoul | 3 | $0.2400 | $172.80 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my05-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | BF1.MEDIUM2 | $0.0200 | $14.40 |
| my05-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | BF1.MEDIUM8 | $0.0700 | $50.40 |
| my05-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | BF1.LARGE16 | $0.1500 | $108.00 |




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
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my05-infra101)",
  "success": true
}
```

