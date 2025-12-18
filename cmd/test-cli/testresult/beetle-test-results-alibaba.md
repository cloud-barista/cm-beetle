# CM-Beetle test results for ALIBABA

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with ALIBABA cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.7+ (f13f82f)
- cm-model: v0.0.15
- CB-Tumblebug: v0.12.1
- CB-Spider: v0.12.0
- CB-MapUI: v0.12.1
- Target CSP: ALIBABA
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: December 19, 2025
- Test Time: 00:11:41 KST
- Test Execution: 2025-12-19 00:11:41 KST

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
| 1 | `POST /beetle/recommendation/vmInfra` | ✅ **PASS** | 647ms | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 1m30.208s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 496ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 77ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 637ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 32.579s | Pass |

**Overall Result**: 7/7 tests passed ✅

**Total Duration**: 3m27.159021758s

*Test executed on December 19, 2025 at 00:11:41 KST (2025-12-19 00:11:41 KST) using CM-Beetle automated test CLI*

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
  "data": [
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
            "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
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
            "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
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
            "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
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
          "uid": "d520j49iir5gq777jdc0",
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
          "uid": "d520j49iir5gq777jdd0",
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
          "uid": "d520j49iir5gq777jde0",
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
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
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
          "id": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
          "uid": "d520jjhiir5gq778da90",
          "name": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-ap-northeast-1",
          "infraType": "",
          "fetchedTime": "2025.12.18 14:03:58 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
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
              "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
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
              "value": "Kernel version is 5.15.0-161-generic, 2025.11.27"
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
              "value": "v2025.11.27"
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
              "value": "2025-11-27T06:55:55Z"
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
              "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
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
          "systemLabel": "",
          "description": "Kernel version is 5.15.0-161-generic, 2025.11.27",
          "commandHistory": null
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
            "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
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
          "uid": "d520j49iir5gq777jdc0",
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
          "uid": "d520j49iir5gq777jdd0",
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
          "uid": "d520j49iir5gq777jde0",
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
          "uid": "d520j49iir5gq777jdbg",
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
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
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
          "id": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
          "uid": "d520jjhiir5gq778da90",
          "name": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-ap-northeast-1",
          "infraType": "",
          "fetchedTime": "2025.12.18 14:03:58 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
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
              "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
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
              "value": "Kernel version is 5.15.0-161-generic, 2025.11.27"
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
              "value": "v2025.11.27"
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
              "value": "2025-11-27T06:55:55Z"
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
              "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
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
          "systemLabel": "",
          "description": "Kernel version is 5.15.0-161-generic, 2025.11.27",
          "commandHistory": null
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

- **Status**: ✅ **SUCCESS**
- **Response**: Infrastructure migration completed successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "d521jmpiir5gq77gftqg",
  "name": "mmci01",
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
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "d521jmpiir5gq77gftqg"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "d521jmpiir5gq77gfttg",
      "cspResourceName": "d521jmpiir5gq77gfttg",
      "cspResourceId": "i-mj741hxlfexfauttybba",
      "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-18 15:13:05",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "8.213.128.181",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.109",
      "privateDNS": "",
      "rootDiskType": "cloud_essd",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "alibaba-ap-northeast-2",
      "connectionConfig": {
        "configName": "alibaba-ap-northeast-2",
        "providerName": "alibaba",
        "driverName": "alibaba-driver-v1.0.so",
        "credentialName": "alibaba",
        "credentialHolder": "admin",
        "regionZoneInfoName": "alibaba-ap-northeast-2",
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionDetail": {
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "description": "South Korea (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
      "cspSpecName": "ecs.t6-c1m4.large",
      "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-mj7m4azyh8kfhz7cls3z1",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "vsw-mj76dkvovy3jskgyqehcq",
      "networkInterface": "eni-mj741hxlfexfautvsgph",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d521jf9iir5gq77gftog",
      "vmUserName": "cb-user",
      "vmUserPassword": "1d5!A1n1ij2ir$",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-18T15:13:17Z",
          "completedTime": "2025-12-18T15:13:20Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj741hxlfexfauttybbaZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.t6-c1m4.large"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "d521jmpiir5gq77gfttg"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "CreditSpecification",
          "value": "Standard"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "80"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj741hxlfexfauttybbaZ"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "e4d29241-72ba-4caf-9ee3-dfb544db1304"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.t6"
        },
        {
          "key": "InstanceId",
          "value": "i-mj741hxlfexfauttybba"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "8192"
        },
        {
          "key": "CreationTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "KeyPairName",
          "value": "d521jf9iir5gq77gftog"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj766f3apmievxq4y5uq]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.213.128.181]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj76dkvovy3jskgyqehcq,VpcId:vpc-mj7m4azyh8kfhz7cls3z1,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.109]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:0a:17:c0,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj741hxlfexfautvsgph,PrimaryIpAddress:10.0.1.109,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.109,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "d521jmpiir5gq77gftrg",
      "cspResourceName": "d521jmpiir5gq77gftrg",
      "cspResourceId": "i-mj75ubrbvaiwdzbkis6p",
      "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-18 15:13:07",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "8.220.218.17",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.110",
      "privateDNS": "",
      "rootDiskType": "cloud_essd",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "alibaba-ap-northeast-2",
      "connectionConfig": {
        "configName": "alibaba-ap-northeast-2",
        "providerName": "alibaba",
        "driverName": "alibaba-driver-v1.0.so",
        "credentialName": "alibaba",
        "credentialHolder": "admin",
        "regionZoneInfoName": "alibaba-ap-northeast-2",
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionDetail": {
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "description": "South Korea (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
      "cspSpecName": "ecs.t6-c1m1.large",
      "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-mj7m4azyh8kfhz7cls3z1",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "vsw-mj76dkvovy3jskgyqehcq",
      "networkInterface": "eni-mj75ubrbvaiwdzbqhilj",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d521jf9iir5gq77gftog",
      "vmUserName": "cb-user",
      "vmUserPassword": "An$1jd1i15ri2!",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-18T15:13:17Z",
          "completedTime": "2025-12-18T15:13:21Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj75ubrbvaiwdzbkis6pZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.t6-c1m1.large"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "d521jmpiir5gq77gftrg"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "CreditSpecification",
          "value": "Standard"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "80"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj75ubrbvaiwdzbkis6pZ"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "13823586-da91-4577-b9a9-2087cef645fb"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.t6"
        },
        {
          "key": "InstanceId",
          "value": "i-mj75ubrbvaiwdzbkis6p"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "2048"
        },
        {
          "key": "CreationTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "KeyPairName",
          "value": "d521jf9iir5gq77gftog"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj71gu0v2ghsfyf7fpvj]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.220.218.17]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj76dkvovy3jskgyqehcq,VpcId:vpc-mj7m4azyh8kfhz7cls3z1,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.110]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:0a:17:c1,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj75ubrbvaiwdzbqhilj,PrimaryIpAddress:10.0.1.110,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.110,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "d521jmpiir5gq77gftsg",
      "cspResourceName": "d521jmpiir5gq77gftsg",
      "cspResourceId": "i-mj7e6duzlirvgpg3s54d",
      "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-18 15:13:02",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "47.80.11.73",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.108",
      "privateDNS": "",
      "rootDiskType": "cloud_essd",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "alibaba-ap-northeast-2",
      "connectionConfig": {
        "configName": "alibaba-ap-northeast-2",
        "providerName": "alibaba",
        "driverName": "alibaba-driver-v1.0.so",
        "credentialName": "alibaba",
        "credentialHolder": "admin",
        "regionZoneInfoName": "alibaba-ap-northeast-2",
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionDetail": {
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "description": "South Korea (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
      "cspSpecName": "ecs.t6-c1m4.xlarge",
      "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-mj7m4azyh8kfhz7cls3z1",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "vsw-mj76dkvovy3jskgyqehcq",
      "networkInterface": "eni-mj7e6duzlirvgpg33v6y",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d521jf9iir5gq77gftog",
      "vmUserName": "cb-user",
      "vmUserPassword": "i1!d1iA1jn25r$",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-18T15:13:17Z",
          "completedTime": "2025-12-18T15:13:20Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7e6duzlirvgpg3s54dZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.t6-c1m4.xlarge"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "d521jmpiir5gq77gftsg"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "CreditSpecification",
          "value": "Standard"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "160"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj7e6duzlirvgpg3s54dZ"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "4"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "1a9816a2-8c59-4a96-af80-a94324cc097c"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.t6"
        },
        {
          "key": "InstanceId",
          "value": "i-mj7e6duzlirvgpg3s54d"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "16384"
        },
        {
          "key": "CreationTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "KeyPairName",
          "value": "d521jf9iir5gq77gftog"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj7ev57i7l0tgzgcptod]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[47.80.11.73]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:2,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj76dkvovy3jskgyqehcq,VpcId:vpc-mj7m4azyh8kfhz7cls3z1,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.108]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:0a:17:bf,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7e6duzlirvgpg33v6y,PrimaryIpAddress:10.0.1.108,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.108,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    }
  ],
  "newVmList": null,
  "postCommand": {
    "userName": "cb-user",
    "command": [
      "uname -a"
    ]
  },
  "postCommandResult": {
    "results": [
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "vmIp": "47.80.11.73",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj7e6duzlirvgpg3s54dZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "vmIp": "8.213.128.181",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj741hxlfexfauttybbaZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vmIp": "8.220.218.17",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj75ubrbvaiwdzbkis6pZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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

### Test Case 3: Get a list of MCIs

#### 3.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/mci`
- **Purpose**: Retrieve all Multi-Cloud Infrastructure instances
- **Namespace ID**: `mig01`
- **Request Body**: None (GET request)

#### 3.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: MCI list retrieved successfully

**Response Body**:

```json
{
  "mci": [
    {
      "resourceType": "mci",
      "id": "mmci01",
      "uid": "d521jmpiir5gq77gftqg",
      "name": "mmci01",
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
        "countUndefined": 0
      },
      "targetStatus": "None",
      "targetAction": "None",
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
        "sys.id": "mmci01",
        "sys.labelType": "mci",
        "sys.manager": "cb-tumblebug",
        "sys.name": "mmci01",
        "sys.namespace": "mig01",
        "sys.uid": "d521jmpiir5gq77gftqg"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "d521jmpiir5gq77gftrg",
          "cspResourceName": "d521jmpiir5gq77gftrg",
          "cspResourceId": "i-mj75ubrbvaiwdzbkis6p",
          "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
          "subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-12-18 15:13:07",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2025-12-18 15:13:07",
            "sys.cspResourceId": "i-mj75ubrbvaiwdzbkis6p",
            "sys.cspResourceName": "d521jmpiir5gq77gftrg",
            "sys.id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.uid": "d521jmpiir5gq77gftrg"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "Region": "ap-northeast-2",
            "Zone": "ap-northeast-2a"
          },
          "publicIP": "8.220.218.17",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.110",
          "privateDNS": "",
          "rootDiskType": "cloud_essd",
          "rootDiskSize": "50",
          "rootDiskName": "",
          "connectionName": "alibaba-ap-northeast-2",
          "connectionConfig": {
            "configName": "alibaba-ap-northeast-2",
            "providerName": "alibaba",
            "driverName": "alibaba-driver-v1.0.so",
            "credentialName": "alibaba",
            "credentialHolder": "admin",
            "regionZoneInfoName": "alibaba-ap-northeast-2",
            "regionZoneInfo": {
              "assignedRegion": "ap-northeast-2",
              "assignedZone": "ap-northeast-2a"
            },
            "regionDetail": {
              "regionId": "ap-northeast-2",
              "regionName": "ap-northeast-2",
              "description": "South Korea (Seoul)",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.36,
                "longitude": 126.78
              },
              "zones": [
                "ap-northeast-2a",
                "ap-northeast-2b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "cspSpecName": "ecs.t6-c1m1.large",
          "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "vpc-mj7m4azyh8kfhz7cls3z1",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "vsw-mj76dkvovy3jskgyqehcq",
          "networkInterface": "eni-mj75ubrbvaiwdzbqhilj",
          "securityGroupIds": [
            "mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d521jf9iir5gq77gftog",
          "vmUserName": "cb-user",
          "vmUserPassword": "An$1jd1i15ri2!",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-12-18T15:13:17Z",
              "completedTime": "2025-12-18T15:13:21Z",
              "elapsedTime": 4,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj75ubrbvaiwdzbkis6pZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
            },
            {
              "key": "InstanceType",
              "value": "ecs.t6-c1m1.large"
            },
            {
              "key": "DeviceAvailable",
              "value": "true"
            },
            {
              "key": "InstanceNetworkType",
              "value": "vpc"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "IsSpot",
              "value": "false"
            },
            {
              "key": "InstanceChargeType",
              "value": "PostPaid"
            },
            {
              "key": "InstanceName",
              "value": "d521jmpiir5gq77gftrg"
            },
            {
              "key": "DeploymentSetGroupNo",
              "value": "0"
            },
            {
              "key": "CreditSpecification",
              "value": "Standard"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "Connected",
              "value": "false"
            },
            {
              "key": "InvocationCount",
              "value": "0"
            },
            {
              "key": "StartTime",
              "value": "2025-12-18T15:12Z"
            },
            {
              "key": "ZoneId",
              "value": "ap-northeast-2a"
            },
            {
              "key": "InternetMaxBandwidthIn",
              "value": "80"
            },
            {
              "key": "InternetChargeType",
              "value": "PayByBandwidth"
            },
            {
              "key": "HostName",
              "value": "iZmj75ubrbvaiwdzbkis6pZ"
            },
            {
              "key": "Status",
              "value": "Running"
            },
            {
              "key": "CPU",
              "value": "0"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "SpotPriceLimit",
              "value": "0.00"
            },
            {
              "key": "OSName",
              "value": "Ubuntu  22.04 64位"
            },
            {
              "key": "InstanceOwnerId",
              "value": "0"
            },
            {
              "key": "OSNameEn",
              "value": "Ubuntu  22.04 64 bit"
            },
            {
              "key": "SerialNumber",
              "value": "13823586-da91-4577-b9a9-2087cef645fb"
            },
            {
              "key": "RegionId",
              "value": "ap-northeast-2"
            },
            {
              "key": "IoOptimized",
              "value": "true"
            },
            {
              "key": "InternetMaxBandwidthOut",
              "value": "5"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InstanceId",
              "value": "i-mj75ubrbvaiwdzbkis6p"
            },
            {
              "key": "Recyclable",
              "value": "false"
            },
            {
              "key": "ExpiredTime",
              "value": "2099-12-31T15:59Z"
            },
            {
              "key": "OSType",
              "value": "linux"
            },
            {
              "key": "Memory",
              "value": "2048"
            },
            {
              "key": "CreationTime",
              "value": "2025-12-18T15:12Z"
            },
            {
              "key": "KeyPairName",
              "value": "d521jf9iir5gq77gftog"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "StoppedMode",
              "value": "Not-applicable"
            },
            {
              "key": "SpotStrategy",
              "value": "NoSpot"
            },
            {
              "key": "SpotDuration",
              "value": "0"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "SecurityGroupIds",
              "value": "{SecurityGroupId:[sg-mj71gu0v2ghsfyf7fpvj]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[8.220.218.17]}"
            },
            {
              "key": "RdmaIpAddress",
              "value": "{IpAddress:null}"
            },
            {
              "key": "DedicatedHostAttribute",
              "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
            },
            {
              "key": "EcsCapacityReservationAttr",
              "value": "{CapacityReservationPreference:,CapacityReservationId:}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "DedicatedInstanceAttribute",
              "value": "{Affinity:,Tenancy:}"
            },
            {
              "key": "PrivateDnsNameOptions",
              "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
            },
            {
              "key": "AdditionalInfo",
              "value": "{EnableHighDensityMode:false}"
            },
            {
              "key": "ImageOptions",
              "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
            },
            {
              "key": "EipAddress",
              "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
            },
            {
              "key": "VpcAttributes",
              "value": "{VSwitchId:vsw-mj76dkvovy3jskgyqehcq,VpcId:vpc-mj7m4azyh8kfhz7cls3z1,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.110]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:0a:17:c1,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj75ubrbvaiwdzbqhilj,PrimaryIpAddress:10.0.1.110,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.110,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
            },
            {
              "key": "OperationLocks",
              "value": "{LockReason:[]}"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "d521jmpiir5gq77gfttg",
          "cspResourceName": "d521jmpiir5gq77gfttg",
          "cspResourceId": "i-mj741hxlfexfauttybba",
          "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-12-18 15:13:05",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2025-12-18 15:13:05",
            "sys.cspResourceId": "i-mj741hxlfexfauttybba",
            "sys.cspResourceName": "d521jmpiir5gq77gfttg",
            "sys.id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.uid": "d521jmpiir5gq77gfttg"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "Region": "ap-northeast-2",
            "Zone": "ap-northeast-2a"
          },
          "publicIP": "8.213.128.181",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.109",
          "privateDNS": "",
          "rootDiskType": "cloud_essd",
          "rootDiskSize": "50",
          "rootDiskName": "",
          "connectionName": "alibaba-ap-northeast-2",
          "connectionConfig": {
            "configName": "alibaba-ap-northeast-2",
            "providerName": "alibaba",
            "driverName": "alibaba-driver-v1.0.so",
            "credentialName": "alibaba",
            "credentialHolder": "admin",
            "regionZoneInfoName": "alibaba-ap-northeast-2",
            "regionZoneInfo": {
              "assignedRegion": "ap-northeast-2",
              "assignedZone": "ap-northeast-2a"
            },
            "regionDetail": {
              "regionId": "ap-northeast-2",
              "regionName": "ap-northeast-2",
              "description": "South Korea (Seoul)",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.36,
                "longitude": 126.78
              },
              "zones": [
                "ap-northeast-2a",
                "ap-northeast-2b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "cspSpecName": "ecs.t6-c1m4.large",
          "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "vpc-mj7m4azyh8kfhz7cls3z1",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "vsw-mj76dkvovy3jskgyqehcq",
          "networkInterface": "eni-mj741hxlfexfautvsgph",
          "securityGroupIds": [
            "mig-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d521jf9iir5gq77gftog",
          "vmUserName": "cb-user",
          "vmUserPassword": "1d5!A1n1ij2ir$",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-12-18T15:13:17Z",
              "completedTime": "2025-12-18T15:13:20Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj741hxlfexfauttybbaZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
            },
            {
              "key": "InstanceType",
              "value": "ecs.t6-c1m4.large"
            },
            {
              "key": "DeviceAvailable",
              "value": "true"
            },
            {
              "key": "InstanceNetworkType",
              "value": "vpc"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "IsSpot",
              "value": "false"
            },
            {
              "key": "InstanceChargeType",
              "value": "PostPaid"
            },
            {
              "key": "InstanceName",
              "value": "d521jmpiir5gq77gfttg"
            },
            {
              "key": "DeploymentSetGroupNo",
              "value": "0"
            },
            {
              "key": "CreditSpecification",
              "value": "Standard"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "Connected",
              "value": "false"
            },
            {
              "key": "InvocationCount",
              "value": "0"
            },
            {
              "key": "StartTime",
              "value": "2025-12-18T15:12Z"
            },
            {
              "key": "ZoneId",
              "value": "ap-northeast-2a"
            },
            {
              "key": "InternetMaxBandwidthIn",
              "value": "80"
            },
            {
              "key": "InternetChargeType",
              "value": "PayByBandwidth"
            },
            {
              "key": "HostName",
              "value": "iZmj741hxlfexfauttybbaZ"
            },
            {
              "key": "Status",
              "value": "Running"
            },
            {
              "key": "CPU",
              "value": "0"
            },
            {
              "key": "Cpu",
              "value": "2"
            },
            {
              "key": "SpotPriceLimit",
              "value": "0.00"
            },
            {
              "key": "OSName",
              "value": "Ubuntu  22.04 64位"
            },
            {
              "key": "InstanceOwnerId",
              "value": "0"
            },
            {
              "key": "OSNameEn",
              "value": "Ubuntu  22.04 64 bit"
            },
            {
              "key": "SerialNumber",
              "value": "e4d29241-72ba-4caf-9ee3-dfb544db1304"
            },
            {
              "key": "RegionId",
              "value": "ap-northeast-2"
            },
            {
              "key": "IoOptimized",
              "value": "true"
            },
            {
              "key": "InternetMaxBandwidthOut",
              "value": "5"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InstanceId",
              "value": "i-mj741hxlfexfauttybba"
            },
            {
              "key": "Recyclable",
              "value": "false"
            },
            {
              "key": "ExpiredTime",
              "value": "2099-12-31T15:59Z"
            },
            {
              "key": "OSType",
              "value": "linux"
            },
            {
              "key": "Memory",
              "value": "8192"
            },
            {
              "key": "CreationTime",
              "value": "2025-12-18T15:12Z"
            },
            {
              "key": "KeyPairName",
              "value": "d521jf9iir5gq77gftog"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "StoppedMode",
              "value": "Not-applicable"
            },
            {
              "key": "SpotStrategy",
              "value": "NoSpot"
            },
            {
              "key": "SpotDuration",
              "value": "0"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "SecurityGroupIds",
              "value": "{SecurityGroupId:[sg-mj766f3apmievxq4y5uq]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[8.213.128.181]}"
            },
            {
              "key": "RdmaIpAddress",
              "value": "{IpAddress:null}"
            },
            {
              "key": "DedicatedHostAttribute",
              "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
            },
            {
              "key": "EcsCapacityReservationAttr",
              "value": "{CapacityReservationPreference:,CapacityReservationId:}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "DedicatedInstanceAttribute",
              "value": "{Affinity:,Tenancy:}"
            },
            {
              "key": "PrivateDnsNameOptions",
              "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
            },
            {
              "key": "AdditionalInfo",
              "value": "{EnableHighDensityMode:false}"
            },
            {
              "key": "ImageOptions",
              "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
            },
            {
              "key": "EipAddress",
              "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
            },
            {
              "key": "VpcAttributes",
              "value": "{VSwitchId:vsw-mj76dkvovy3jskgyqehcq,VpcId:vpc-mj7m4azyh8kfhz7cls3z1,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.109]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:0a:17:c0,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj741hxlfexfautvsgph,PrimaryIpAddress:10.0.1.109,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.109,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
            },
            {
              "key": "OperationLocks",
              "value": "{LockReason:[]}"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "d521jmpiir5gq77gftsg",
          "cspResourceName": "d521jmpiir5gq77gftsg",
          "cspResourceId": "i-mj7e6duzlirvgpg3s54d",
          "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-12-18 15:13:02",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2025-12-18 15:13:02",
            "sys.cspResourceId": "i-mj7e6duzlirvgpg3s54d",
            "sys.cspResourceName": "d521jmpiir5gq77gftsg",
            "sys.id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.uid": "d521jmpiir5gq77gftsg"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "Region": "ap-northeast-2",
            "Zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.11.73",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.108",
          "privateDNS": "",
          "rootDiskType": "cloud_essd",
          "rootDiskSize": "50",
          "rootDiskName": "",
          "connectionName": "alibaba-ap-northeast-2",
          "connectionConfig": {
            "configName": "alibaba-ap-northeast-2",
            "providerName": "alibaba",
            "driverName": "alibaba-driver-v1.0.so",
            "credentialName": "alibaba",
            "credentialHolder": "admin",
            "regionZoneInfoName": "alibaba-ap-northeast-2",
            "regionZoneInfo": {
              "assignedRegion": "ap-northeast-2",
              "assignedZone": "ap-northeast-2a"
            },
            "regionDetail": {
              "regionId": "ap-northeast-2",
              "regionName": "ap-northeast-2",
              "description": "South Korea (Seoul)",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.36,
                "longitude": 126.78
              },
              "zones": [
                "ap-northeast-2a",
                "ap-northeast-2b"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "cspSpecName": "ecs.t6-c1m4.xlarge",
          "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "vpc-mj7m4azyh8kfhz7cls3z1",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "vsw-mj76dkvovy3jskgyqehcq",
          "networkInterface": "eni-mj7e6duzlirvgpg33v6y",
          "securityGroupIds": [
            "mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d521jf9iir5gq77gftog",
          "vmUserName": "cb-user",
          "vmUserPassword": "i1!d1iA1jn25r$",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-12-18T15:13:17Z",
              "completedTime": "2025-12-18T15:13:20Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj7e6duzlirvgpg3s54dZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
            },
            {
              "key": "InstanceType",
              "value": "ecs.t6-c1m4.xlarge"
            },
            {
              "key": "DeviceAvailable",
              "value": "true"
            },
            {
              "key": "InstanceNetworkType",
              "value": "vpc"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "IsSpot",
              "value": "false"
            },
            {
              "key": "InstanceChargeType",
              "value": "PostPaid"
            },
            {
              "key": "InstanceName",
              "value": "d521jmpiir5gq77gftsg"
            },
            {
              "key": "DeploymentSetGroupNo",
              "value": "0"
            },
            {
              "key": "CreditSpecification",
              "value": "Standard"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "Connected",
              "value": "false"
            },
            {
              "key": "InvocationCount",
              "value": "0"
            },
            {
              "key": "StartTime",
              "value": "2025-12-18T15:12Z"
            },
            {
              "key": "ZoneId",
              "value": "ap-northeast-2a"
            },
            {
              "key": "InternetMaxBandwidthIn",
              "value": "160"
            },
            {
              "key": "InternetChargeType",
              "value": "PayByBandwidth"
            },
            {
              "key": "HostName",
              "value": "iZmj7e6duzlirvgpg3s54dZ"
            },
            {
              "key": "Status",
              "value": "Running"
            },
            {
              "key": "CPU",
              "value": "0"
            },
            {
              "key": "Cpu",
              "value": "4"
            },
            {
              "key": "SpotPriceLimit",
              "value": "0.00"
            },
            {
              "key": "OSName",
              "value": "Ubuntu  22.04 64位"
            },
            {
              "key": "InstanceOwnerId",
              "value": "0"
            },
            {
              "key": "OSNameEn",
              "value": "Ubuntu  22.04 64 bit"
            },
            {
              "key": "SerialNumber",
              "value": "1a9816a2-8c59-4a96-af80-a94324cc097c"
            },
            {
              "key": "RegionId",
              "value": "ap-northeast-2"
            },
            {
              "key": "IoOptimized",
              "value": "true"
            },
            {
              "key": "InternetMaxBandwidthOut",
              "value": "5"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InstanceId",
              "value": "i-mj7e6duzlirvgpg3s54d"
            },
            {
              "key": "Recyclable",
              "value": "false"
            },
            {
              "key": "ExpiredTime",
              "value": "2099-12-31T15:59Z"
            },
            {
              "key": "OSType",
              "value": "linux"
            },
            {
              "key": "Memory",
              "value": "16384"
            },
            {
              "key": "CreationTime",
              "value": "2025-12-18T15:12Z"
            },
            {
              "key": "KeyPairName",
              "value": "d521jf9iir5gq77gftog"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "StoppedMode",
              "value": "Not-applicable"
            },
            {
              "key": "SpotStrategy",
              "value": "NoSpot"
            },
            {
              "key": "SpotDuration",
              "value": "0"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "SecurityGroupIds",
              "value": "{SecurityGroupId:[sg-mj7ev57i7l0tgzgcptod]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[47.80.11.73]}"
            },
            {
              "key": "RdmaIpAddress",
              "value": "{IpAddress:null}"
            },
            {
              "key": "DedicatedHostAttribute",
              "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
            },
            {
              "key": "EcsCapacityReservationAttr",
              "value": "{CapacityReservationPreference:,CapacityReservationId:}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:2,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "DedicatedInstanceAttribute",
              "value": "{Affinity:,Tenancy:}"
            },
            {
              "key": "PrivateDnsNameOptions",
              "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
            },
            {
              "key": "AdditionalInfo",
              "value": "{EnableHighDensityMode:false}"
            },
            {
              "key": "ImageOptions",
              "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
            },
            {
              "key": "EipAddress",
              "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
            },
            {
              "key": "VpcAttributes",
              "value": "{VSwitchId:vsw-mj76dkvovy3jskgyqehcq,VpcId:vpc-mj7m4azyh8kfhz7cls3z1,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.108]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:0a:17:bf,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7e6duzlirvgpg33v6y,PrimaryIpAddress:10.0.1.108,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.108,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
            },
            {
              "key": "OperationLocks",
              "value": "{LockReason:[]}"
            }
          ]
        }
      ],
      "newVmList": null,
      "postCommand": {
        "userName": "cb-user",
        "command": [
          "uname -a"
        ]
      },
      "postCommandResult": {
        "results": [
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "vmIp": "47.80.11.73",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7e6duzlirvgpg3s54dZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "vmIp": "8.213.128.181",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj741hxlfexfauttybbaZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
            "vmIp": "8.220.218.17",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj75ubrbvaiwdzbkis6pZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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

### Test Case 4: Get a list of MCI IDs

#### 4.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/mci?option=id`
- **Purpose**: Retrieve MCI IDs only (lightweight response)
- **Namespace ID**: `mig01`
- **Query Parameter**: `option=id`
- **Request Body**: None (GET request)

#### 4.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: MCI IDs retrieved successfully

**Response Body**:

```json
{
  "idList": [
    "mmci01"
  ]
}
```

### Test Case 5: Get a specific MCI

#### 5.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/mci/{{mciId}}`
- **Purpose**: Retrieve detailed information for a specific MCI
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{mciId}}` - The specific MCI identifier
- **Request Body**: None (GET request)

#### 5.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: MCI details retrieved successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "d521jmpiir5gq77gftqg",
  "name": "mmci01",
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
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "d521jmpiir5gq77gftqg"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "d521jmpiir5gq77gftrg",
      "cspResourceName": "d521jmpiir5gq77gftrg",
      "cspResourceId": "i-mj75ubrbvaiwdzbkis6p",
      "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-18 15:13:07",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2025-12-18 15:13:07",
        "sys.cspResourceId": "i-mj75ubrbvaiwdzbkis6p",
        "sys.cspResourceName": "d521jmpiir5gq77gftrg",
        "sys.id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.uid": "d521jmpiir5gq77gftrg"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "8.220.218.17",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.110",
      "privateDNS": "",
      "rootDiskType": "cloud_essd",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "alibaba-ap-northeast-2",
      "connectionConfig": {
        "configName": "alibaba-ap-northeast-2",
        "providerName": "alibaba",
        "driverName": "alibaba-driver-v1.0.so",
        "credentialName": "alibaba",
        "credentialHolder": "admin",
        "regionZoneInfoName": "alibaba-ap-northeast-2",
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionDetail": {
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "description": "South Korea (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
      "cspSpecName": "ecs.t6-c1m1.large",
      "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-mj7m4azyh8kfhz7cls3z1",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "vsw-mj76dkvovy3jskgyqehcq",
      "networkInterface": "eni-mj75ubrbvaiwdzbqhilj",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d521jf9iir5gq77gftog",
      "vmUserName": "cb-user",
      "vmUserPassword": "An$1jd1i15ri2!",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-18T15:13:17Z",
          "completedTime": "2025-12-18T15:13:21Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj75ubrbvaiwdzbkis6pZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.t6-c1m1.large"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "d521jmpiir5gq77gftrg"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "CreditSpecification",
          "value": "Standard"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "80"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj75ubrbvaiwdzbkis6pZ"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "13823586-da91-4577-b9a9-2087cef645fb"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.t6"
        },
        {
          "key": "InstanceId",
          "value": "i-mj75ubrbvaiwdzbkis6p"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "2048"
        },
        {
          "key": "CreationTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "KeyPairName",
          "value": "d521jf9iir5gq77gftog"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj71gu0v2ghsfyf7fpvj]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.220.218.17]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj76dkvovy3jskgyqehcq,VpcId:vpc-mj7m4azyh8kfhz7cls3z1,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.110]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:0a:17:c1,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj75ubrbvaiwdzbqhilj,PrimaryIpAddress:10.0.1.110,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.110,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "d521jmpiir5gq77gfttg",
      "cspResourceName": "d521jmpiir5gq77gfttg",
      "cspResourceId": "i-mj741hxlfexfauttybba",
      "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-18 15:13:05",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2025-12-18 15:13:05",
        "sys.cspResourceId": "i-mj741hxlfexfauttybba",
        "sys.cspResourceName": "d521jmpiir5gq77gfttg",
        "sys.id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.uid": "d521jmpiir5gq77gfttg"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "8.213.128.181",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.109",
      "privateDNS": "",
      "rootDiskType": "cloud_essd",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "alibaba-ap-northeast-2",
      "connectionConfig": {
        "configName": "alibaba-ap-northeast-2",
        "providerName": "alibaba",
        "driverName": "alibaba-driver-v1.0.so",
        "credentialName": "alibaba",
        "credentialHolder": "admin",
        "regionZoneInfoName": "alibaba-ap-northeast-2",
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionDetail": {
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "description": "South Korea (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
      "cspSpecName": "ecs.t6-c1m4.large",
      "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-mj7m4azyh8kfhz7cls3z1",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "vsw-mj76dkvovy3jskgyqehcq",
      "networkInterface": "eni-mj741hxlfexfautvsgph",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d521jf9iir5gq77gftog",
      "vmUserName": "cb-user",
      "vmUserPassword": "1d5!A1n1ij2ir$",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-18T15:13:17Z",
          "completedTime": "2025-12-18T15:13:20Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj741hxlfexfauttybbaZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.t6-c1m4.large"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "d521jmpiir5gq77gfttg"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "CreditSpecification",
          "value": "Standard"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "80"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj741hxlfexfauttybbaZ"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "e4d29241-72ba-4caf-9ee3-dfb544db1304"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.t6"
        },
        {
          "key": "InstanceId",
          "value": "i-mj741hxlfexfauttybba"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "8192"
        },
        {
          "key": "CreationTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "KeyPairName",
          "value": "d521jf9iir5gq77gftog"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj766f3apmievxq4y5uq]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.213.128.181]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj76dkvovy3jskgyqehcq,VpcId:vpc-mj7m4azyh8kfhz7cls3z1,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.109]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:0a:17:c0,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj741hxlfexfautvsgph,PrimaryIpAddress:10.0.1.109,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.109,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "d521jmpiir5gq77gftsg",
      "cspResourceName": "d521jmpiir5gq77gftsg",
      "cspResourceId": "i-mj7e6duzlirvgpg3s54d",
      "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-18 15:13:02",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2025-12-18 15:13:02",
        "sys.cspResourceId": "i-mj7e6duzlirvgpg3s54d",
        "sys.cspResourceName": "d521jmpiir5gq77gftsg",
        "sys.id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.uid": "d521jmpiir5gq77gftsg"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "47.80.11.73",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.108",
      "privateDNS": "",
      "rootDiskType": "cloud_essd",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "alibaba-ap-northeast-2",
      "connectionConfig": {
        "configName": "alibaba-ap-northeast-2",
        "providerName": "alibaba",
        "driverName": "alibaba-driver-v1.0.so",
        "credentialName": "alibaba",
        "credentialHolder": "admin",
        "regionZoneInfoName": "alibaba-ap-northeast-2",
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionDetail": {
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "description": "South Korea (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
      "cspSpecName": "ecs.t6-c1m4.xlarge",
      "imageId": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20251126.vhd",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-mj7m4azyh8kfhz7cls3z1",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "vsw-mj76dkvovy3jskgyqehcq",
      "networkInterface": "eni-mj7e6duzlirvgpg33v6y",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d521jf9iir5gq77gftog",
      "vmUserName": "cb-user",
      "vmUserPassword": "i1!d1iA1jn25r$",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-18T15:13:17Z",
          "completedTime": "2025-12-18T15:13:20Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj7e6duzlirvgpg3s54dZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20251126.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.t6-c1m4.xlarge"
        },
        {
          "key": "DeviceAvailable",
          "value": "true"
        },
        {
          "key": "InstanceNetworkType",
          "value": "vpc"
        },
        {
          "key": "LocalStorageAmount",
          "value": "0"
        },
        {
          "key": "IsSpot",
          "value": "false"
        },
        {
          "key": "InstanceChargeType",
          "value": "PostPaid"
        },
        {
          "key": "InstanceName",
          "value": "d521jmpiir5gq77gftsg"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
        },
        {
          "key": "CreditSpecification",
          "value": "Standard"
        },
        {
          "key": "GPUAmount",
          "value": "0"
        },
        {
          "key": "Connected",
          "value": "false"
        },
        {
          "key": "InvocationCount",
          "value": "0"
        },
        {
          "key": "StartTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "160"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj7e6duzlirvgpg3s54dZ"
        },
        {
          "key": "Status",
          "value": "Running"
        },
        {
          "key": "CPU",
          "value": "0"
        },
        {
          "key": "Cpu",
          "value": "4"
        },
        {
          "key": "SpotPriceLimit",
          "value": "0.00"
        },
        {
          "key": "OSName",
          "value": "Ubuntu  22.04 64位"
        },
        {
          "key": "InstanceOwnerId",
          "value": "0"
        },
        {
          "key": "OSNameEn",
          "value": "Ubuntu  22.04 64 bit"
        },
        {
          "key": "SerialNumber",
          "value": "1a9816a2-8c59-4a96-af80-a94324cc097c"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "IoOptimized",
          "value": "true"
        },
        {
          "key": "InternetMaxBandwidthOut",
          "value": "5"
        },
        {
          "key": "InstanceTypeFamily",
          "value": "ecs.t6"
        },
        {
          "key": "InstanceId",
          "value": "i-mj7e6duzlirvgpg3s54d"
        },
        {
          "key": "Recyclable",
          "value": "false"
        },
        {
          "key": "ExpiredTime",
          "value": "2099-12-31T15:59Z"
        },
        {
          "key": "OSType",
          "value": "linux"
        },
        {
          "key": "Memory",
          "value": "16384"
        },
        {
          "key": "CreationTime",
          "value": "2025-12-18T15:12Z"
        },
        {
          "key": "KeyPairName",
          "value": "d521jf9iir5gq77gftog"
        },
        {
          "key": "LocalStorageCapacity",
          "value": "0"
        },
        {
          "key": "StoppedMode",
          "value": "Not-applicable"
        },
        {
          "key": "SpotStrategy",
          "value": "NoSpot"
        },
        {
          "key": "SpotDuration",
          "value": "0"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "SecurityGroupIds",
          "value": "{SecurityGroupId:[sg-mj7ev57i7l0tgzgcptod]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[47.80.11.73]}"
        },
        {
          "key": "RdmaIpAddress",
          "value": "{IpAddress:null}"
        },
        {
          "key": "DedicatedHostAttribute",
          "value": "{DedicatedHostName:,DedicatedHostClusterId:,DedicatedHostId:}"
        },
        {
          "key": "EcsCapacityReservationAttr",
          "value": "{CapacityReservationPreference:,CapacityReservationId:}"
        },
        {
          "key": "CpuOptions",
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:2,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "DedicatedInstanceAttribute",
          "value": "{Affinity:,Tenancy:}"
        },
        {
          "key": "PrivateDnsNameOptions",
          "value": "{EnableInstanceIdDnsARecord:false,EnableInstanceIdDnsAAAARecord:false,EnableIpDnsARecord:false,EnableIpDnsPtrRecord:false,HostnameType:}"
        },
        {
          "key": "AdditionalInfo",
          "value": "{EnableHighDensityMode:false}"
        },
        {
          "key": "ImageOptions",
          "value": "{ImageFamily:,LoginAsNonRoot:false,ImageName:,Description:,CurrentOSNVMeSupported:false,ImageFeatures:{NvmeSupport:},ImageTags:{ImageTag:null}}"
        },
        {
          "key": "EipAddress",
          "value": "{IsSupportUnassociate:false,InternetChargeType:,IpAddress:,Bandwidth:0,AllocationId:}"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:,HttpPutResponseHopLimit:0,HttpTokens:}"
        },
        {
          "key": "VpcAttributes",
          "value": "{VSwitchId:vsw-mj76dkvovy3jskgyqehcq,VpcId:vpc-mj7m4azyh8kfhz7cls3z1,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.108]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:0a:17:bf,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7e6duzlirvgpg33v6y,PrimaryIpAddress:10.0.1.108,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.108,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    }
  ],
  "newVmList": null,
  "postCommand": {
    "userName": "cb-user",
    "command": [
      "uname -a"
    ]
  },
  "postCommandResult": {
    "results": [
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "vmIp": "47.80.11.73",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj7e6duzlirvgpg3s54dZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "vmIp": "8.213.128.181",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj741hxlfexfauttybbaZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vmIp": "8.220.218.17",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj75ubrbvaiwdzbkis6pZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "output": "Linux iZmj75ubrbvaiwdzbkis6pZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "8.220.218.17",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
      "testOrder": 1,
      "userName": "cb-user",
      "vmId": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux iZmj741hxlfexfauttybbaZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "8.213.128.181",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "testOrder": 2,
      "userName": "cb-user",
      "vmId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux iZmj7e6duzlirvgpg3s54dZ 5.15.0-161-generic #171-Ubuntu SMP Sat Oct 11 08:17:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "47.80.11.73",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "testOrder": 3,
      "userName": "cb-user",
      "vmId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1"
    }
  ]
}
```

</details>

### Test Case 7: Delete the migrated computing infra

#### 7.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}`
- **Purpose**: Delete the migrated infrastructure and clean up resources
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{mciId}}` - The MCI identifier to delete
- **Query Parameter**: `option=terminate` (terminates all resources)
- **Request Body**: None (DELETE request)

#### 7.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infrastructure deletion completed successfully

**Response Body**:

```json
{
  "success": true,
  "text": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: mmci01)"
}
```

