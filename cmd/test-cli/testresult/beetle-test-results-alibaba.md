# CM-Beetle test results for ALIBABA

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with ALIBABA cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.5+ (ed85eeb)
- cm-model: v0.0.14
- CB-Tumblebug: v0.11.19
- CB-Spider: v0.11.16
- CB-MapUI: v0.11.19
- Target CSP: ALIBABA
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: December 5, 2025
- Test Time: 16:16:07 KST
- Test Execution: 2025-12-05 16:16:07 KST

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for ALIBABA

### Test Results Summary

| Test | Endpoint | Status | Duration | Details |
|------|----------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/vmInfra` | ✅ **PASS** | 764ms | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ❌ **FAIL** | 26.496s | Fail |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ⏭️ **SKIP** | 0s | Skip |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ⏭️ **SKIP** | 0s | Skip |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ⏭️ **SKIP** | 0s | Skip |
| 6 | Remote Command Accessibility Check | ⏭️ **SKIP** | 0s | Skip |
| 7 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ⏭️ **SKIP** | 0s | Skip |

**Overall Result**: 1/7 tests passed, 5 skipped ❌

**Total Duration**: 37.316146806s

*Test executed on December 5, 2025 at 16:16:07 KST (2025-12-05 16:16:07 KST) using CM-Beetle automated test CLI*

---

## Detailed Test Case Results

> [!INFO]
> This section provides detailed information for each test case, including API request information and response details.

### Test Case 1: Recommend a target model for computing infra

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/recommendation/vmInfra`
- **Purpose**: Get infrastructure recommendations for migration
- **Required Parameters**: `desiredCsp` and `desiredRegion` in request body

**Request Body**:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "desiredCspAndRegionPair": {
    "csp": "alibaba",
    "region": "ap-northeast-2"
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
    "servers": [
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
  "items": [
    {
      "status": "partially-matched",
      "description": "Candidate #1 | partially-matched | Overall Match Rate: Min=75.0% Max=100.0% Avg=91.7% | VMs: 3 total, 0 matched, 3 acceptable",
      "targetCloud": {
        "csp": "alibaba",
        "region": "ap-northeast-2"
      },
      "targetVmInfra": {
        "name": "mmci01",
        "installMonAgent": "",
        "label": null,
        "systemLabel": "",
        "description": "Recommended VMs comprising multi-cloud infrastructure",
        "subGroups": [
          {
            "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20250917.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "TYPE1",
            "rootDiskSize": "50",
            "dataDiskIds": null
          },
          {
            "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20250917.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "TYPE1",
            "rootDiskSize": "50",
            "dataDiskIds": null
          },
          {
            "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20250917.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-03"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "TYPE1",
            "rootDiskSize": "50",
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
        "connectionName": "alibaba-ap-northeast-2",
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
        "connectionName": "alibaba-ap-northeast-2",
        "description": "a SSH Key pair for migration (Note - provided ONLY once, MUST be downloaded",
        "cspResourceId": "",
        "fingerprint": "",
        "username": "",
        "verifiedUsername": "",
        "publicKey": "",
        "privateKey": ""
      },
      "targetVmSpecList": [
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "uid": "d3vjcrmqjs728pq3g4jg",
          "cspSpecName": "ecs.t6-c1m1.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.02139,
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
          "rootDiskSize": "-1",
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MemorySize",
              "value": "2.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "100000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "2"
            },
            {
              "key": "CpuCoreCount",
              "value": "2"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m1.large"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "81920"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "60"
            },
            {
              "key": "InstancePpsTx",
              "value": "100000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "81920"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "40"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "uid": "d3vjcrmqjs728pq3g4kg",
          "cspSpecName": "ecs.t6-c1m4.xlarge",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.16926,
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
          "rootDiskSize": "-1",
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MemorySize",
              "value": "16.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "200000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "6"
            },
            {
              "key": "CpuCoreCount",
              "value": "4"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m4.xlarge"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "163840"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "120"
            },
            {
              "key": "InstancePpsTx",
              "value": "200000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "163840"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "160"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "uid": "d3vjcrmqjs728pq3g4lg",
          "cspSpecName": "ecs.t6-c1m4.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.08463,
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
          "rootDiskSize": "-1",
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MemorySize",
              "value": "8.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "100000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "2"
            },
            {
              "key": "CpuCoreCount",
              "value": "2"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m4.large"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "81920"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "60"
            },
            {
              "key": "InstancePpsTx",
              "value": "100000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "81920"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "60"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        }
      ],
      "targetVmOsImageList": [
        {
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20250917.vhd",
          "regionList": [
            "ap-northeast-1",
            "ap-northeast-2",
            "ap-southeast-1",
            "ap-southeast-3",
            "ap-southeast-5",
            "ap-southeast-6",
            "ap-southeast-7",
            "cn-beijing",
            "cn-chengdu",
            "cn-fuzhou",
            "cn-guangzhou",
            "cn-hangzhou",
            "cn-heyuan",
            "cn-hongkong",
            "cn-huhehaote",
            "cn-nanjing",
            "cn-qingdao",
            "cn-shanghai",
            "cn-shenzhen",
            "cn-wuhan-lr",
            "cn-wulanchabu",
            "cn-zhangjiakou",
            "eu-central-1",
            "eu-west-1",
            "me-central-1",
            "me-east-1",
            "na-south-1",
            "us-east-1",
            "us-west-1"
          ],
          "id": "ubuntu_22_04_x64_20G_alibase_20250917.vhd",
          "uid": "d3vjd0mqjs728pq4adm0",
          "name": "ubuntu_22_04_x64_20G_alibase_20250917.vhd",
          "connectionName": "alibaba-ap-northeast-1",
          "fetchedTime": "2025.10.27 09:08:18 Mon",
          "isBasicImage": true,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu  22.04 64 bit",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "BootMode",
              "value": "UEFI-Preferred"
            },
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20250917.vhd"
            },
            {
              "key": "ImageOwnerAlias",
              "value": "system"
            },
            {
              "key": "OSName",
              "value": "Ubuntu  22.04 64位"
            },
            {
              "key": "OSNameEn",
              "value": "Ubuntu  22.04 64 bit"
            },
            {
              "key": "ImageFamily",
              "value": "acs:ubuntu_22_04_x64"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "IsSupportIoOptimized",
              "value": "true"
            },
            {
              "key": "Size",
              "value": "20"
            },
            {
              "key": "Description",
              "value": "Kernel version is 5.15.0-153-generic, 2025.9.17"
            },
            {
              "key": "Usage",
              "value": "instance"
            },
            {
              "key": "IsCopied",
              "value": "false"
            },
            {
              "key": "LoginAsNonRootSupported",
              "value": "true"
            },
            {
              "key": "ImageVersion",
              "value": "v2025.9.17"
            },
            {
              "key": "OSType",
              "value": "linux"
            },
            {
              "key": "IsSubscribed",
              "value": "false"
            },
            {
              "key": "IsSupportCloudinit",
              "value": "true"
            },
            {
              "key": "CreationTime",
              "value": "2025-09-17T15:45:18Z"
            },
            {
              "key": "Progress",
              "value": "100%"
            },
            {
              "key": "Platform",
              "value": "Ubuntu"
            },
            {
              "key": "ImageName",
              "value": "ubuntu_22_04_x64_20G_alibase_20250917.vhd"
            },
            {
              "key": "Status",
              "value": "Available"
            },
            {
              "key": "ImageOwnerId",
              "value": "0"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "DetectionOptions",
              "value": "{Status:,Items:{Item:null}}"
            },
            {
              "key": "Features",
              "value": "{MemoryOnlineUpgrade:unsupported,NvmeSupport:supported,CpuOnlineDowngrade:unsupported,ImdsSupport:v2,MemoryOnlineDowngrade:unsupported,CpuOnlineUpgrade:unsupported}"
            },
            {
              "key": "Tags",
              "value": "{Tag:[]}"
            },
            {
              "key": "DiskDeviceMappings",
              "value": "{DiskDeviceMapping:[]}"
            }
          ],
          "description": "Kernel version is 5.15.0-153-generic, 2025.9.17"
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-vnet-01",
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
          "name": "mig-sg-02",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-vnet-01",
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
          "name": "mig-sg-03",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-vnet-01",
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
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=50.0% Max=100.0% Avg=75.0% | VMs: 1 total, 0 matched, 1 acceptable",
      "targetCloud": {
        "csp": "alibaba",
        "region": "ap-northeast-2"
      },
      "targetVmInfra": {
        "name": "mmci01",
        "installMonAgent": "",
        "label": null,
        "systemLabel": "",
        "description": "Recommended VMs comprising multi-cloud infrastructure",
        "subGroups": [
          {
            "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c2m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20250917.vhd",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-01"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "TYPE1",
            "rootDiskSize": "50",
            "dataDiskIds": null
          },
          {
            "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "a recommended virtual machine 02 for ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "",
            "imageId": "",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-02"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "TYPE1",
            "rootDiskSize": "50",
            "dataDiskIds": null
          },
          {
            "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "a recommended virtual machine 03 for ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "",
            "imageId": "",
            "vNetId": "mig-vnet-01",
            "subnetId": "mig-subnet-01",
            "securityGroupIds": [
              "mig-sg-03"
            ],
            "sshKeyId": "mig-sshkey-01",
            "rootDiskType": "TYPE1",
            "rootDiskSize": "50",
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
        "connectionName": "alibaba-ap-northeast-2",
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
        "connectionName": "alibaba-ap-northeast-2",
        "description": "a SSH Key pair for migration (Note - provided ONLY once, MUST be downloaded",
        "cspResourceId": "",
        "fingerprint": "",
        "username": "",
        "verifiedUsername": "",
        "publicKey": "",
        "privateKey": ""
      },
      "targetVmSpecList": [
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "uid": "d3vjcrmqjs728pq3g4jg",
          "cspSpecName": "ecs.t6-c1m1.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.02139,
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
          "rootDiskSize": "-1",
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MemorySize",
              "value": "2.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "100000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "2"
            },
            {
              "key": "CpuCoreCount",
              "value": "2"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m1.large"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "81920"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "60"
            },
            {
              "key": "InstancePpsTx",
              "value": "100000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "81920"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "40"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "uid": "d3vjcrmqjs728pq3g4kg",
          "cspSpecName": "ecs.t6-c1m4.xlarge",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.16926,
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
          "rootDiskSize": "-1",
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MemorySize",
              "value": "16.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "200000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "6"
            },
            {
              "key": "CpuCoreCount",
              "value": "4"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m4.xlarge"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "163840"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "120"
            },
            {
              "key": "InstancePpsTx",
              "value": "200000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "163840"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "160"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "uid": "d3vjcrmqjs728pq3g4lg",
          "cspSpecName": "ecs.t6-c1m4.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.08463,
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
          "rootDiskSize": "-1",
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MemorySize",
              "value": "8.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "100000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "2"
            },
            {
              "key": "CpuCoreCount",
              "value": "2"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m4.large"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "81920"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "60"
            },
            {
              "key": "InstancePpsTx",
              "value": "100000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "81920"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "60"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c2m1.large",
          "uid": "d3vjcrmqjs728pq3g4j0",
          "cspSpecName": "ecs.t6-c2m1.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c2m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1,
          "diskSizeGB": -1,
          "costPerHour": 0.01116,
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
          "rootDiskSize": "-1",
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "MemorySize",
              "value": "1.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "60000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "2"
            },
            {
              "key": "CpuCoreCount",
              "value": "2"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c2m1.large"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "81920"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "60"
            },
            {
              "key": "InstancePpsTx",
              "value": "60000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "81920"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "20"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        }
      ],
      "targetVmOsImageList": [
        {
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20250917.vhd",
          "regionList": [
            "ap-northeast-1",
            "ap-northeast-2",
            "ap-southeast-1",
            "ap-southeast-3",
            "ap-southeast-5",
            "ap-southeast-6",
            "ap-southeast-7",
            "cn-beijing",
            "cn-chengdu",
            "cn-fuzhou",
            "cn-guangzhou",
            "cn-hangzhou",
            "cn-heyuan",
            "cn-hongkong",
            "cn-huhehaote",
            "cn-nanjing",
            "cn-qingdao",
            "cn-shanghai",
            "cn-shenzhen",
            "cn-wuhan-lr",
            "cn-wulanchabu",
            "cn-zhangjiakou",
            "eu-central-1",
            "eu-west-1",
            "me-central-1",
            "me-east-1",
            "na-south-1",
            "us-east-1",
            "us-west-1"
          ],
          "id": "ubuntu_22_04_x64_20G_alibase_20250917.vhd",
          "uid": "d3vjd0mqjs728pq4adm0",
          "name": "ubuntu_22_04_x64_20G_alibase_20250917.vhd",
          "connectionName": "alibaba-ap-northeast-1",
          "fetchedTime": "2025.10.27 09:08:18 Mon",
          "isBasicImage": true,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu  22.04 64 bit",
          "osDiskType": "NA",
          "osDiskSizeGB": 20,
          "imageStatus": "Available",
          "details": [
            {
              "key": "BootMode",
              "value": "UEFI-Preferred"
            },
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20250917.vhd"
            },
            {
              "key": "ImageOwnerAlias",
              "value": "system"
            },
            {
              "key": "OSName",
              "value": "Ubuntu  22.04 64位"
            },
            {
              "key": "OSNameEn",
              "value": "Ubuntu  22.04 64 bit"
            },
            {
              "key": "ImageFamily",
              "value": "acs:ubuntu_22_04_x64"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "IsSupportIoOptimized",
              "value": "true"
            },
            {
              "key": "Size",
              "value": "20"
            },
            {
              "key": "Description",
              "value": "Kernel version is 5.15.0-153-generic, 2025.9.17"
            },
            {
              "key": "Usage",
              "value": "instance"
            },
            {
              "key": "IsCopied",
              "value": "false"
            },
            {
              "key": "LoginAsNonRootSupported",
              "value": "true"
            },
            {
              "key": "ImageVersion",
              "value": "v2025.9.17"
            },
            {
              "key": "OSType",
              "value": "linux"
            },
            {
              "key": "IsSubscribed",
              "value": "false"
            },
            {
              "key": "IsSupportCloudinit",
              "value": "true"
            },
            {
              "key": "CreationTime",
              "value": "2025-09-17T15:45:18Z"
            },
            {
              "key": "Progress",
              "value": "100%"
            },
            {
              "key": "Platform",
              "value": "Ubuntu"
            },
            {
              "key": "ImageName",
              "value": "ubuntu_22_04_x64_20G_alibase_20250917.vhd"
            },
            {
              "key": "Status",
              "value": "Available"
            },
            {
              "key": "ImageOwnerId",
              "value": "0"
            },
            {
              "key": "IsPublic",
              "value": "true"
            },
            {
              "key": "DetectionOptions",
              "value": "{Status:,Items:{Item:null}}"
            },
            {
              "key": "Features",
              "value": "{MemoryOnlineUpgrade:unsupported,NvmeSupport:supported,CpuOnlineDowngrade:unsupported,ImdsSupport:v2,MemoryOnlineDowngrade:unsupported,CpuOnlineUpgrade:unsupported}"
            },
            {
              "key": "Tags",
              "value": "{Tag:[]}"
            },
            {
              "key": "DiskDeviceMappings",
              "value": "{DiskDeviceMapping:[]}"
            }
          ],
          "description": "Kernel version is 5.15.0-153-generic, 2025.9.17"
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-vnet-01",
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
          "name": "mig-sg-02",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-vnet-01",
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
          "name": "mig-sg-03",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-vnet-01",
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

- **API Endpoint**: `POST /beetle/migration/ns/mig01/mci`
- **Purpose**: Create and migrate infrastructure based on recommendation
- **Namespace ID**: `mig01`
- **Request Body**: Uses the response from the previous recommendation step

#### 2.2 API Response Information

- **Status**: ❌ **FAILED**
- **Error**: Migration failed

### Test Case 3: Get a list of MCIs

#### 3.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/mci`
- **Purpose**: Retrieve all Multi-Cloud Infrastructure instances
- **Namespace ID**: `mig01`
- **Request Body**: None (GET request)

#### 3.2 API Response Information

- **Status**: ❌ **FAILED**
- **Error**: No response received

**Error Message**:

```
Test skipped due to previous test failure
```

### Test Case 4: Get a list of MCI IDs

#### 4.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/mci?option=id`
- **Purpose**: Retrieve MCI IDs only (lightweight response)
- **Namespace ID**: `mig01`
- **Query Parameter**: `option=id`
- **Request Body**: None (GET request)

#### 4.2 API Response Information

- **Status**: ❌ **FAILED**
- **Error**: No response received

**Error Message**:

```
Test skipped due to previous test failure
```

### Test Case 6: Remote Command Accessibility Check

#### 6.1 Test Information

- **Test Type**: SSH Connectivity Test for All VMs
- **Purpose**: Verify that all migrated VMs are accessible via SSH
- **Method**: Extract public IP and SSH key from MCI access info for each VM, then execute remote command
- **Command Executed**: `uname -a` (to verify system information)
- **Authentication**: SSH key-based authentication
- **Scope**: Tests all VMs across all subgroups in the MCI

#### 6.2 Test Result Information

- **Status**: ⏭️ **SKIPPED**
- **Reason**: Test skipped due to previous test failure

### Test Case 7: Delete the migrated computing infra

#### 7.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}`
- **Purpose**: Delete the migrated infrastructure and clean up resources
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{mciId}}` - The MCI identifier to delete
- **Query Parameter**: `option=terminate` (terminates all resources)
- **Request Body**: None (DELETE request)

#### 7.2 API Response Information

- **Status**: ⏭️ **SKIPPED**
- **Reason**: Test skipped due to previous test failure

