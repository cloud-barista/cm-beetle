# CM-Beetle test results for AWS

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AWS cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.4+ (3f15f14)
- cm-model: v0.0.14
- CB-Tumblebug: v0.11.19
- CB-Spider: v0.11.16
- CB-MapUI: v0.11.19
- Target CSP: AWS
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: November 17, 2025
- Test Time: 04:46:40 UTC
- Test Execution: 2025-11-17 04:46:40 UTC

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for AWS

### Test Results Summary

| Test | Endpoint | Status | Duration | Details |
|------|----------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/mci` | ✅ **PASS** | 1.549s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 50.039s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 166ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 71ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 85ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 1m43.924s | Pass |

**Overall Result**: 7/7 tests passed ✅

**Total Duration**: 3m51.390743166s

*Test executed on November 17, 2025 at 04:46:40 UTC (2025-11-17 04:46:40 UTC) using CM-Beetle automated test CLI*

---

## Detailed Test Case Results

> [!INFO]
> This section provides detailed information for each test case, including API request information and response details.

### Test Case 1: Recommend a target model for computing infra

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/recommendation/mci`
- **Purpose**: Get infrastructure recommendations for migration
- **Required Parameters**: `desiredCsp` and `desiredRegion` in request body

**Request Body**:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "desiredCspAndRegionPair": {
    "csp": "aws",
    "region": "ap-northeast-2"
  },
  "OnpremiseInfraModel": {
    "network": {
      "ipv4Networks": {
        "defaultGateways": [
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec2643f0-9388-3d97-f3a4-f387cd52696c"
          },
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec21fd51-16bb-7e10-5e23-12ef283b2204"
          },
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec2876a6-3c84-7e62-aaf9-c3203f12e0b8"
          }
        ]
      },
      "ipv6Networks": {}
    },
    "servers": [
      {
        "hostname": "ip-10-0-1-66",
        "machineId": "ec2643f0-9388-3d97-f3a4-f387cd52696c",
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
            "macAddress": "02:41:71:df:76:2f",
            "ipv4CidrBlocks": [
              "10.0.1.66/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::41:71ff:fedf:762f/64"
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
            "destination": "fe80::41:71ff:fedf:762f/128",
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
        "hostname": "ip-10-0-1-85",
        "machineId": "ec21fd51-16bb-7e10-5e23-12ef283b2204",
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
            "macAddress": "02:d0:77:0d:fe:c9",
            "ipv4CidrBlocks": [
              "10.0.1.85/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::d0:77ff:fe0d:fec9/64"
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
            "destination": "fe80::d0:77ff:fe0d:fec9/128",
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
        "hostname": "ip-10-0-1-9",
        "machineId": "ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
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
            "macAddress": "02:f2:b6:d6:9b:75",
            "ipv4CidrBlocks": [
              "10.0.1.9/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::f2:b6ff:fed6:9b75/64"
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
            "destination": "fe80::f2:b6ff:fed6:9b75/128",
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
  "status": "",
  "description": "This is a list of recommended target infrastructures. Please review and use them.",
  "targetCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
  },
  "targetVmInfra": {
    "name": "mmci01",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "a recommended multi-cloud infrastructure",
    "subGroups": [
      {
        "name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec2643f0-9388-3d97-f3a4-f387cd52696c"
        },
        "description": "a recommended virtual machine 01 for ec2643f0-9388-3d97-f3a4-f387cd52696c",
        "connectionName": "aws-ap-northeast-2",
        "specId": "aws+ap-northeast-2+t3a.small",
        "imageId": "ami-010be25c3775061c9",
        "vNetId": "mig-vnet-01",
        "subnetId": "mig-subnet-01",
        "securityGroupIds": [
          "mig-sg-01"
        ],
        "sshKeyId": "mig-sshkey-01",
        "rootDiskSize": "50",
        "dataDiskIds": null
      },
      {
        "name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec21fd51-16bb-7e10-5e23-12ef283b2204"
        },
        "description": "a recommended virtual machine 02 for ec21fd51-16bb-7e10-5e23-12ef283b2204",
        "connectionName": "aws-ap-northeast-2",
        "specId": "aws+ap-northeast-2+t3a.xlarge",
        "imageId": "ami-010be25c3775061c9",
        "vNetId": "mig-vnet-01",
        "subnetId": "mig-subnet-01",
        "securityGroupIds": [
          "mig-sg-02"
        ],
        "sshKeyId": "mig-sshkey-01",
        "rootDiskSize": "50",
        "dataDiskIds": null
      },
      {
        "name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec2876a6-3c84-7e62-aaf9-c3203f12e0b8"
        },
        "description": "a recommended virtual machine 03 for ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
        "connectionName": "aws-ap-northeast-2",
        "specId": "aws+ap-northeast-2+t3a.large",
        "imageId": "ami-010be25c3775061c9",
        "vNetId": "mig-vnet-01",
        "subnetId": "mig-subnet-01",
        "securityGroupIds": [
          "mig-sg-03"
        ],
        "sshKeyId": "mig-sshkey-01",
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
    "connectionName": "aws-ap-northeast-2",
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
    "connectionName": "aws-ap-northeast-2",
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
      "id": "aws+ap-northeast-2+t3a.small",
      "uid": "d4d80kqp2foc73e81dl0",
      "cspSpecName": "t3a.small",
      "name": "aws+ap-northeast-2+t3a.small",
      "namespace": "system",
      "connectionName": "aws-ap-northeast-2",
      "providerName": "aws",
      "regionName": "ap-northeast-2",
      "regionLatitude": 37.36,
      "regionLongitude": 126.78,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 2,
      "diskSizeGB": -1,
      "costPerHour": 0.0234,
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
          "key": "AutoRecoverySupported",
          "value": "true"
        },
        {
          "key": "BareMetal",
          "value": "false"
        },
        {
          "key": "BurstablePerformanceSupported",
          "value": "true"
        },
        {
          "key": "CurrentGeneration",
          "value": "true"
        },
        {
          "key": "DedicatedHostsSupported",
          "value": "false"
        },
        {
          "key": "EbsInfo",
          "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:175,BaselineIops:1000,BaselineThroughputInMBps:21.875,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
        },
        {
          "key": "FreeTierEligible",
          "value": "false"
        },
        {
          "key": "HibernationSupported",
          "value": "true"
        },
        {
          "key": "Hypervisor",
          "value": "nitro"
        },
        {
          "key": "InstanceStorageSupported",
          "value": "false"
        },
        {
          "key": "InstanceType",
          "value": "t3a.small"
        },
        {
          "key": "MemoryInfo",
          "value": "{SizeInMiB:2048}"
        },
        {
          "key": "NetworkInfo",
          "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:4,Ipv6AddressesPerInterface:4,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:2,NetworkCards:[{MaximumNetworkInterfaces:2,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
        },
        {
          "key": "PlacementGroupInfo",
          "value": "{SupportedStrategies:[partition,spread]}"
        },
        {
          "key": "ProcessorInfo",
          "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
        },
        {
          "key": "SupportedBootModes",
          "value": "legacy-bios; uefi"
        },
        {
          "key": "SupportedRootDeviceTypes",
          "value": "ebs"
        },
        {
          "key": "SupportedUsageClasses",
          "value": "on-demand; spot"
        },
        {
          "key": "SupportedVirtualizationTypes",
          "value": "hvm"
        },
        {
          "key": "VCpuInfo",
          "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
        }
      ]
    },
    {
      "id": "aws+ap-northeast-2+t3a.xlarge",
      "uid": "d4d80kqp2foc73e81ju0",
      "cspSpecName": "t3a.xlarge",
      "name": "aws+ap-northeast-2+t3a.xlarge",
      "namespace": "system",
      "connectionName": "aws-ap-northeast-2",
      "providerName": "aws",
      "regionName": "ap-northeast-2",
      "regionLatitude": 37.36,
      "regionLongitude": 126.78,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 16,
      "diskSizeGB": -1,
      "costPerHour": 0.1872,
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
          "key": "AutoRecoverySupported",
          "value": "true"
        },
        {
          "key": "BareMetal",
          "value": "false"
        },
        {
          "key": "BurstablePerformanceSupported",
          "value": "true"
        },
        {
          "key": "CurrentGeneration",
          "value": "true"
        },
        {
          "key": "DedicatedHostsSupported",
          "value": "false"
        },
        {
          "key": "EbsInfo",
          "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
        },
        {
          "key": "FreeTierEligible",
          "value": "false"
        },
        {
          "key": "HibernationSupported",
          "value": "true"
        },
        {
          "key": "Hypervisor",
          "value": "nitro"
        },
        {
          "key": "InstanceStorageSupported",
          "value": "false"
        },
        {
          "key": "InstanceType",
          "value": "t3a.xlarge"
        },
        {
          "key": "MemoryInfo",
          "value": "{SizeInMiB:16384}"
        },
        {
          "key": "NetworkInfo",
          "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:15,Ipv6AddressesPerInterface:15,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:4,NetworkCards:[{MaximumNetworkInterfaces:4,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
        },
        {
          "key": "PlacementGroupInfo",
          "value": "{SupportedStrategies:[partition,spread]}"
        },
        {
          "key": "ProcessorInfo",
          "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
        },
        {
          "key": "SupportedBootModes",
          "value": "legacy-bios; uefi"
        },
        {
          "key": "SupportedRootDeviceTypes",
          "value": "ebs"
        },
        {
          "key": "SupportedUsageClasses",
          "value": "on-demand; spot"
        },
        {
          "key": "SupportedVirtualizationTypes",
          "value": "hvm"
        },
        {
          "key": "VCpuInfo",
          "value": "{DefaultCores:2,DefaultThreadsPerCore:2,DefaultVCpus:4,ValidCores:[2],ValidThreadsPerCore:[1,2]}"
        }
      ]
    },
    {
      "id": "aws+ap-northeast-2+t3a.large",
      "uid": "d4d80kqp2foc73e81fqg",
      "cspSpecName": "t3a.large",
      "name": "aws+ap-northeast-2+t3a.large",
      "namespace": "system",
      "connectionName": "aws-ap-northeast-2",
      "providerName": "aws",
      "regionName": "ap-northeast-2",
      "regionLatitude": 37.36,
      "regionLongitude": 126.78,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 8,
      "diskSizeGB": -1,
      "costPerHour": 0.0936,
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
          "key": "AutoRecoverySupported",
          "value": "true"
        },
        {
          "key": "BareMetal",
          "value": "false"
        },
        {
          "key": "BurstablePerformanceSupported",
          "value": "true"
        },
        {
          "key": "CurrentGeneration",
          "value": "true"
        },
        {
          "key": "DedicatedHostsSupported",
          "value": "false"
        },
        {
          "key": "EbsInfo",
          "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
        },
        {
          "key": "FreeTierEligible",
          "value": "false"
        },
        {
          "key": "HibernationSupported",
          "value": "true"
        },
        {
          "key": "Hypervisor",
          "value": "nitro"
        },
        {
          "key": "InstanceStorageSupported",
          "value": "false"
        },
        {
          "key": "InstanceType",
          "value": "t3a.large"
        },
        {
          "key": "MemoryInfo",
          "value": "{SizeInMiB:8192}"
        },
        {
          "key": "NetworkInfo",
          "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:12,Ipv6AddressesPerInterface:12,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
        },
        {
          "key": "PlacementGroupInfo",
          "value": "{SupportedStrategies:[partition,spread]}"
        },
        {
          "key": "ProcessorInfo",
          "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
        },
        {
          "key": "SupportedBootModes",
          "value": "legacy-bios; uefi"
        },
        {
          "key": "SupportedRootDeviceTypes",
          "value": "ebs"
        },
        {
          "key": "SupportedUsageClasses",
          "value": "on-demand; spot"
        },
        {
          "key": "SupportedVirtualizationTypes",
          "value": "hvm"
        },
        {
          "key": "VCpuInfo",
          "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
        }
      ]
    }
  ],
  "targetVmOsImageList": [
    {
      "namespace": "system",
      "providerName": "aws",
      "cspImageName": "ami-010be25c3775061c9",
      "regionList": [
        "ap-northeast-2"
      ],
      "id": "ami-010be25c3775061c9",
      "uid": "d4d80v2p2foc73e9roc0",
      "name": "ami-010be25c3775061c9",
      "connectionName": "aws-ap-northeast-2",
      "fetchedTime": "2025.11.17 01:55:08 Mon",
      "creationDate": "2025-10-15T08:05:34.000Z",
      "isBasicImage": true,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015",
      "osDiskType": "ebs",
      "osDiskSizeGB": -1,
      "imageStatus": "Available",
      "details": [
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-0bad9aec71fa3bdcb,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CreationDate",
          "value": "2025-10-15T08:05:34.000Z"
        },
        {
          "key": "DeprecationTime",
          "value": "2027-10-15T08:05:34.000Z"
        },
        {
          "key": "Description",
          "value": "Canonical, Ubuntu, 22.04, amd64 jammy image"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-010be25c3775061c9"
        },
        {
          "key": "ImageLocation",
          "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015"
        },
        {
          "key": "ImageOwnerAlias",
          "value": "amazon"
        },
        {
          "key": "ImageType",
          "value": "machine"
        },
        {
          "key": "Name",
          "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251015"
        },
        {
          "key": "OwnerId",
          "value": "099720109477"
        },
        {
          "key": "PlatformDetails",
          "value": "Linux/UNIX"
        },
        {
          "key": "Public",
          "value": "true"
        },
        {
          "key": "RootDeviceName",
          "value": "/dev/sda1"
        },
        {
          "key": "RootDeviceType",
          "value": "ebs"
        },
        {
          "key": "SriovNetSupport",
          "value": "simple"
        },
        {
          "key": "State",
          "value": "available"
        },
        {
          "key": "UsageOperation",
          "value": "RunInstances"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        }
      ],
      "description": "Canonical, Ubuntu, 22.04, amd64 jammy image"
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "mig-sg-01",
      "connectionName": "aws-ap-northeast-2",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for ec2643f0-9388-3d97-f3a4-f387cd52696c",
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
      "connectionName": "aws-ap-northeast-2",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for ec21fd51-16bb-7e10-5e23-12ef283b2204",
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
      "connectionName": "aws-ap-northeast-2",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
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
  "uid": "d4dahj2p2foc73bcejf0",
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
    "sys.description": "a recommended multi-cloud infrastructure",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "d4dahj2p2foc73bcejf0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
      "uid": "d4dahj2p2foc73bcejh0",
      "cspResourceName": "d4dahj2p2foc73bcejh0",
      "cspResourceId": "i-0dfa1890fcefba4be",
      "name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
      "subGroupId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
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
      "createdTime": "2025-11-17 04:47:29",
      "label": {
        "sourceMachineId": "ec21fd51-16bb-7e10-5e23-12ef283b2204"
      },
      "description": "a recommended virtual machine 02 for ec21fd51-16bb-7e10-5e23-12ef283b2204",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "43.203.121.210",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.62",
      "privateDNS": "ip-10-0-1-62.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
      "specId": "aws+ap-northeast-2+t3a.xlarge",
      "cspSpecName": "t3a.xlarge",
      "imageId": "ami-010be25c3775061c9",
      "cspImageName": "ami-010be25c3775061c9",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-0330b15879aede8d2",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "subnet-00440b2d1ddf6b379",
      "networkInterface": "eni-attach-0c159cce7fb3b5358",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4dahgap2foc73bcejd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:47:38Z",
          "completedTime": "2025-11-17T04:47:41Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-62 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-11-17T04:47:12Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-020d3659c57d50830}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "2A25D144-CAA6-4CFF-B264-3AFE27484140"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:2,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-010be25c3775061c9"
        },
        {
          "key": "InstanceId",
          "value": "i-0dfa1890fcefba4be"
        },
        {
          "key": "InstanceType",
          "value": "t3a.xlarge"
        },
        {
          "key": "KeyName",
          "value": "d4dahgap2foc73bcejd0"
        },
        {
          "key": "LaunchTime",
          "value": "2025-11-17T04:47:11Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.121.210},Attachment:{AttachTime:2025-11-17T04:47:11Z,AttachmentId:eni-attach-0c159cce7fb3b5358,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0893e1a7e93bcf006,GroupName:d4dahhip2foc73bceje0}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:81:c5:56:58:31,NetworkInterfaceId:eni-0f0d37224ebe972af,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.62,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.121.210},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.62}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-00440b2d1ddf6b379,VpcId:vpc-0330b15879aede8d2}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-62.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.62"
        },
        {
          "key": "PublicIpAddress",
          "value": "43.203.121.210"
        },
        {
          "key": "RootDeviceName",
          "value": "/dev/sda1"
        },
        {
          "key": "RootDeviceType",
          "value": "ebs"
        },
        {
          "key": "SecurityGroups",
          "value": "{GroupId:sg-0893e1a7e93bcf006,GroupName:d4dahhip2foc73bceje0}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-00440b2d1ddf6b379"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d4dahj2p2foc73bcejh0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0330b15879aede8d2"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
      "uid": "d4dahj2p2foc73bcejg0",
      "cspResourceName": "d4dahj2p2foc73bcejg0",
      "cspResourceId": "i-00c25957c7fdb5cfc",
      "name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
      "subGroupId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
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
      "createdTime": "2025-11-17 04:47:29",
      "label": {
        "sourceMachineId": "ec2643f0-9388-3d97-f3a4-f387cd52696c"
      },
      "description": "a recommended virtual machine 01 for ec2643f0-9388-3d97-f3a4-f387cd52696c",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "13.124.191.35",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.241",
      "privateDNS": "ip-10-0-1-241.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
      "specId": "aws+ap-northeast-2+t3a.small",
      "cspSpecName": "t3a.small",
      "imageId": "ami-010be25c3775061c9",
      "cspImageName": "ami-010be25c3775061c9",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-0330b15879aede8d2",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "subnet-00440b2d1ddf6b379",
      "networkInterface": "eni-attach-0b65eb9ad52fe76c6",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4dahgap2foc73bcejd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:47:38Z",
          "completedTime": "2025-11-17T04:47:40Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-241 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-11-17T04:47:12Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0d40b4ecdf3a18873}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "3CE7E5E0-AF45-43E8-8A97-11FDDDFE8593"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:1,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-010be25c3775061c9"
        },
        {
          "key": "InstanceId",
          "value": "i-00c25957c7fdb5cfc"
        },
        {
          "key": "InstanceType",
          "value": "t3a.small"
        },
        {
          "key": "KeyName",
          "value": "d4dahgap2foc73bcejd0"
        },
        {
          "key": "LaunchTime",
          "value": "2025-11-17T04:47:11Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.124.191.35},Attachment:{AttachTime:2025-11-17T04:47:11Z,AttachmentId:eni-attach-0b65eb9ad52fe76c6,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-08d51712b18745856,GroupName:d4dahh2p2foc73bcejdg}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:ba:38:58:8c:11,NetworkInterfaceId:eni-0a3d64de6fe9e9fed,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.241,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.124.191.35},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.241}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-00440b2d1ddf6b379,VpcId:vpc-0330b15879aede8d2}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-241.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.241"
        },
        {
          "key": "PublicIpAddress",
          "value": "13.124.191.35"
        },
        {
          "key": "RootDeviceName",
          "value": "/dev/sda1"
        },
        {
          "key": "RootDeviceType",
          "value": "ebs"
        },
        {
          "key": "SecurityGroups",
          "value": "{GroupId:sg-08d51712b18745856,GroupName:d4dahh2p2foc73bcejdg}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-00440b2d1ddf6b379"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d4dahj2p2foc73bcejg0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0330b15879aede8d2"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
      "uid": "d4dahj2p2foc73bceji0",
      "cspResourceName": "d4dahj2p2foc73bceji0",
      "cspResourceId": "i-0a5882b217e54f74f",
      "name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
      "subGroupId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
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
      "createdTime": "2025-11-17 04:47:31",
      "label": {
        "sourceMachineId": "ec2876a6-3c84-7e62-aaf9-c3203f12e0b8"
      },
      "description": "a recommended virtual machine 03 for ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "3.39.6.146",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.185",
      "privateDNS": "ip-10-0-1-185.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
      "specId": "aws+ap-northeast-2+t3a.large",
      "cspSpecName": "t3a.large",
      "imageId": "ami-010be25c3775061c9",
      "cspImageName": "ami-010be25c3775061c9",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-0330b15879aede8d2",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "subnet-00440b2d1ddf6b379",
      "networkInterface": "eni-attach-09ef8a8839fd87e58",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4dahgap2foc73bcejd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:47:38Z",
          "completedTime": "2025-11-17T04:47:41Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-185 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-11-17T04:47:11Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0f5ce9be98cb95562}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "C81902E7-4F92-4804-A85D-0B2979D2493D"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:1,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-010be25c3775061c9"
        },
        {
          "key": "InstanceId",
          "value": "i-0a5882b217e54f74f"
        },
        {
          "key": "InstanceType",
          "value": "t3a.large"
        },
        {
          "key": "KeyName",
          "value": "d4dahgap2foc73bcejd0"
        },
        {
          "key": "LaunchTime",
          "value": "2025-11-17T04:47:10Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.6.146},Attachment:{AttachTime:2025-11-17T04:47:10Z,AttachmentId:eni-attach-09ef8a8839fd87e58,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-023e0f0b2c789b8f8,GroupName:d4dahiap2foc73bcejeg}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:95:ec:03:1e:e5,NetworkInterfaceId:eni-0f94e50ba0cd6b0d5,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.185,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.6.146},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.185}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-00440b2d1ddf6b379,VpcId:vpc-0330b15879aede8d2}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-185.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.185"
        },
        {
          "key": "PublicIpAddress",
          "value": "3.39.6.146"
        },
        {
          "key": "RootDeviceName",
          "value": "/dev/sda1"
        },
        {
          "key": "RootDeviceType",
          "value": "ebs"
        },
        {
          "key": "SecurityGroups",
          "value": "{GroupId:sg-023e0f0b2c789b8f8,GroupName:d4dahiap2foc73bcejeg}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-00440b2d1ddf6b379"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d4dahj2p2foc73bceji0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0330b15879aede8d2"
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
        "vmId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
        "vmIp": "13.124.191.35",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-241 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
        "vmIp": "43.203.121.210",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-62 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
        "vmIp": "3.39.6.146",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-185 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "d4dahj2p2foc73bcejf0",
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
        "sys.description": "a recommended multi-cloud infrastructure",
        "sys.id": "mmci01",
        "sys.labelType": "mci",
        "sys.manager": "cb-tumblebug",
        "sys.name": "mmci01",
        "sys.namespace": "mig01",
        "sys.uid": "d4dahj2p2foc73bcejf0"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "a recommended multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
          "uid": "d4dahj2p2foc73bcejh0",
          "cspResourceName": "d4dahj2p2foc73bcejh0",
          "cspResourceId": "i-0dfa1890fcefba4be",
          "name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
          "subGroupId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
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
          "createdTime": "2025-11-17 04:47:29",
          "label": {
            "sourceMachineId": "ec21fd51-16bb-7e10-5e23-12ef283b2204",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2025-11-17 04:47:29",
            "sys.cspResourceId": "i-0dfa1890fcefba4be",
            "sys.cspResourceName": "d4dahj2p2foc73bcejh0",
            "sys.id": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
            "sys.uid": "d4dahj2p2foc73bcejh0"
          },
          "description": "a recommended virtual machine 02 for ec21fd51-16bb-7e10-5e23-12ef283b2204",
          "region": {
            "Region": "ap-northeast-2",
            "Zone": "ap-northeast-2a"
          },
          "publicIP": "43.203.121.210",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.62",
          "privateDNS": "ip-10-0-1-62.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": "50",
          "rootDiskName": "",
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
          "specId": "aws+ap-northeast-2+t3a.xlarge",
          "cspSpecName": "t3a.xlarge",
          "imageId": "ami-010be25c3775061c9",
          "cspImageName": "ami-010be25c3775061c9",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "vpc-0330b15879aede8d2",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "subnet-00440b2d1ddf6b379",
          "networkInterface": "eni-attach-0c159cce7fb3b5358",
          "securityGroupIds": [
            "mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d4dahgap2foc73bcejd0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-17T04:47:38Z",
              "completedTime": "2025-11-17T04:47:41Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-62 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "AmiLaunchIndex",
              "value": "0"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "BlockDeviceMappings",
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-11-17T04:47:12Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-020d3659c57d50830}}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CapacityReservationSpecification",
              "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
            },
            {
              "key": "ClientToken",
              "value": "2A25D144-CAA6-4CFF-B264-3AFE27484140"
            },
            {
              "key": "CpuOptions",
              "value": "{CoreCount:2,ThreadsPerCore:2}"
            },
            {
              "key": "EbsOptimized",
              "value": "false"
            },
            {
              "key": "EnaSupport",
              "value": "true"
            },
            {
              "key": "EnclaveOptions",
              "value": "{Enabled:false}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "ImageId",
              "value": "ami-010be25c3775061c9"
            },
            {
              "key": "InstanceId",
              "value": "i-0dfa1890fcefba4be"
            },
            {
              "key": "InstanceType",
              "value": "t3a.xlarge"
            },
            {
              "key": "KeyName",
              "value": "d4dahgap2foc73bcejd0"
            },
            {
              "key": "LaunchTime",
              "value": "2025-11-17T04:47:11Z"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
            },
            {
              "key": "Monitoring",
              "value": "{State:disabled}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.121.210},Attachment:{AttachTime:2025-11-17T04:47:11Z,AttachmentId:eni-attach-0c159cce7fb3b5358,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0893e1a7e93bcf006,GroupName:d4dahhip2foc73bceje0}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:81:c5:56:58:31,NetworkInterfaceId:eni-0f0d37224ebe972af,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.62,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.121.210},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.62}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-00440b2d1ddf6b379,VpcId:vpc-0330b15879aede8d2}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-62.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.62"
            },
            {
              "key": "PublicIpAddress",
              "value": "43.203.121.210"
            },
            {
              "key": "RootDeviceName",
              "value": "/dev/sda1"
            },
            {
              "key": "RootDeviceType",
              "value": "ebs"
            },
            {
              "key": "SecurityGroups",
              "value": "{GroupId:sg-0893e1a7e93bcf006,GroupName:d4dahhip2foc73bceje0}"
            },
            {
              "key": "SourceDestCheck",
              "value": "true"
            },
            {
              "key": "State",
              "value": "{Code:16,Name:running}"
            },
            {
              "key": "SubnetId",
              "value": "subnet-00440b2d1ddf6b379"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:d4dahj2p2foc73bcejh0}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0330b15879aede8d2"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
          "uid": "d4dahj2p2foc73bcejg0",
          "cspResourceName": "d4dahj2p2foc73bcejg0",
          "cspResourceId": "i-00c25957c7fdb5cfc",
          "name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
          "subGroupId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
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
          "createdTime": "2025-11-17 04:47:29",
          "label": {
            "sourceMachineId": "ec2643f0-9388-3d97-f3a4-f387cd52696c",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2025-11-17 04:47:29",
            "sys.cspResourceId": "i-00c25957c7fdb5cfc",
            "sys.cspResourceName": "d4dahj2p2foc73bcejg0",
            "sys.id": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
            "sys.uid": "d4dahj2p2foc73bcejg0"
          },
          "description": "a recommended virtual machine 01 for ec2643f0-9388-3d97-f3a4-f387cd52696c",
          "region": {
            "Region": "ap-northeast-2",
            "Zone": "ap-northeast-2a"
          },
          "publicIP": "13.124.191.35",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.241",
          "privateDNS": "ip-10-0-1-241.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": "50",
          "rootDiskName": "",
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
          "specId": "aws+ap-northeast-2+t3a.small",
          "cspSpecName": "t3a.small",
          "imageId": "ami-010be25c3775061c9",
          "cspImageName": "ami-010be25c3775061c9",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "vpc-0330b15879aede8d2",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "subnet-00440b2d1ddf6b379",
          "networkInterface": "eni-attach-0b65eb9ad52fe76c6",
          "securityGroupIds": [
            "mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d4dahgap2foc73bcejd0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-17T04:47:38Z",
              "completedTime": "2025-11-17T04:47:40Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-241 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "AmiLaunchIndex",
              "value": "0"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "BlockDeviceMappings",
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-11-17T04:47:12Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0d40b4ecdf3a18873}}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CapacityReservationSpecification",
              "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
            },
            {
              "key": "ClientToken",
              "value": "3CE7E5E0-AF45-43E8-8A97-11FDDDFE8593"
            },
            {
              "key": "CpuOptions",
              "value": "{CoreCount:1,ThreadsPerCore:2}"
            },
            {
              "key": "EbsOptimized",
              "value": "false"
            },
            {
              "key": "EnaSupport",
              "value": "true"
            },
            {
              "key": "EnclaveOptions",
              "value": "{Enabled:false}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "ImageId",
              "value": "ami-010be25c3775061c9"
            },
            {
              "key": "InstanceId",
              "value": "i-00c25957c7fdb5cfc"
            },
            {
              "key": "InstanceType",
              "value": "t3a.small"
            },
            {
              "key": "KeyName",
              "value": "d4dahgap2foc73bcejd0"
            },
            {
              "key": "LaunchTime",
              "value": "2025-11-17T04:47:11Z"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
            },
            {
              "key": "Monitoring",
              "value": "{State:disabled}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.124.191.35},Attachment:{AttachTime:2025-11-17T04:47:11Z,AttachmentId:eni-attach-0b65eb9ad52fe76c6,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-08d51712b18745856,GroupName:d4dahh2p2foc73bcejdg}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:ba:38:58:8c:11,NetworkInterfaceId:eni-0a3d64de6fe9e9fed,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.241,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.124.191.35},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.241}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-00440b2d1ddf6b379,VpcId:vpc-0330b15879aede8d2}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-241.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.241"
            },
            {
              "key": "PublicIpAddress",
              "value": "13.124.191.35"
            },
            {
              "key": "RootDeviceName",
              "value": "/dev/sda1"
            },
            {
              "key": "RootDeviceType",
              "value": "ebs"
            },
            {
              "key": "SecurityGroups",
              "value": "{GroupId:sg-08d51712b18745856,GroupName:d4dahh2p2foc73bcejdg}"
            },
            {
              "key": "SourceDestCheck",
              "value": "true"
            },
            {
              "key": "State",
              "value": "{Code:16,Name:running}"
            },
            {
              "key": "SubnetId",
              "value": "subnet-00440b2d1ddf6b379"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:d4dahj2p2foc73bcejg0}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0330b15879aede8d2"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
          "uid": "d4dahj2p2foc73bceji0",
          "cspResourceName": "d4dahj2p2foc73bceji0",
          "cspResourceId": "i-0a5882b217e54f74f",
          "name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
          "subGroupId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
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
          "createdTime": "2025-11-17 04:47:31",
          "label": {
            "sourceMachineId": "ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2025-11-17 04:47:31",
            "sys.cspResourceId": "i-0a5882b217e54f74f",
            "sys.cspResourceName": "d4dahj2p2foc73bceji0",
            "sys.id": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
            "sys.uid": "d4dahj2p2foc73bceji0"
          },
          "description": "a recommended virtual machine 03 for ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
          "region": {
            "Region": "ap-northeast-2",
            "Zone": "ap-northeast-2a"
          },
          "publicIP": "3.39.6.146",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.185",
          "privateDNS": "ip-10-0-1-185.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": "50",
          "rootDiskName": "",
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
          "specId": "aws+ap-northeast-2+t3a.large",
          "cspSpecName": "t3a.large",
          "imageId": "ami-010be25c3775061c9",
          "cspImageName": "ami-010be25c3775061c9",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "vpc-0330b15879aede8d2",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "subnet-00440b2d1ddf6b379",
          "networkInterface": "eni-attach-09ef8a8839fd87e58",
          "securityGroupIds": [
            "mig-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d4dahgap2foc73bcejd0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-17T04:47:38Z",
              "completedTime": "2025-11-17T04:47:41Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-185 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "AmiLaunchIndex",
              "value": "0"
            },
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "BlockDeviceMappings",
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-11-17T04:47:11Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0f5ce9be98cb95562}}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CapacityReservationSpecification",
              "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
            },
            {
              "key": "ClientToken",
              "value": "C81902E7-4F92-4804-A85D-0B2979D2493D"
            },
            {
              "key": "CpuOptions",
              "value": "{CoreCount:1,ThreadsPerCore:2}"
            },
            {
              "key": "EbsOptimized",
              "value": "false"
            },
            {
              "key": "EnaSupport",
              "value": "true"
            },
            {
              "key": "EnclaveOptions",
              "value": "{Enabled:false}"
            },
            {
              "key": "HibernationOptions",
              "value": "{Configured:false}"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "ImageId",
              "value": "ami-010be25c3775061c9"
            },
            {
              "key": "InstanceId",
              "value": "i-0a5882b217e54f74f"
            },
            {
              "key": "InstanceType",
              "value": "t3a.large"
            },
            {
              "key": "KeyName",
              "value": "d4dahgap2foc73bcejd0"
            },
            {
              "key": "LaunchTime",
              "value": "2025-11-17T04:47:10Z"
            },
            {
              "key": "MetadataOptions",
              "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
            },
            {
              "key": "Monitoring",
              "value": "{State:disabled}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.6.146},Attachment:{AttachTime:2025-11-17T04:47:10Z,AttachmentId:eni-attach-09ef8a8839fd87e58,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-023e0f0b2c789b8f8,GroupName:d4dahiap2foc73bcejeg}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:95:ec:03:1e:e5,NetworkInterfaceId:eni-0f94e50ba0cd6b0d5,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.185,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.6.146},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.185}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-00440b2d1ddf6b379,VpcId:vpc-0330b15879aede8d2}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-185.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.185"
            },
            {
              "key": "PublicIpAddress",
              "value": "3.39.6.146"
            },
            {
              "key": "RootDeviceName",
              "value": "/dev/sda1"
            },
            {
              "key": "RootDeviceType",
              "value": "ebs"
            },
            {
              "key": "SecurityGroups",
              "value": "{GroupId:sg-023e0f0b2c789b8f8,GroupName:d4dahiap2foc73bcejeg}"
            },
            {
              "key": "SourceDestCheck",
              "value": "true"
            },
            {
              "key": "State",
              "value": "{Code:16,Name:running}"
            },
            {
              "key": "SubnetId",
              "value": "subnet-00440b2d1ddf6b379"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:d4dahj2p2foc73bceji0}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0330b15879aede8d2"
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
            "vmId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
            "vmIp": "13.124.191.35",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-241 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
            "vmIp": "43.203.121.210",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-62 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
            "vmIp": "3.39.6.146",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-185 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "d4dahj2p2foc73bcejf0",
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
    "sys.description": "a recommended multi-cloud infrastructure",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "d4dahj2p2foc73bcejf0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
      "uid": "d4dahj2p2foc73bcejh0",
      "cspResourceName": "d4dahj2p2foc73bcejh0",
      "cspResourceId": "i-0dfa1890fcefba4be",
      "name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
      "subGroupId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
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
      "createdTime": "2025-11-17 04:47:29",
      "label": {
        "sourceMachineId": "ec21fd51-16bb-7e10-5e23-12ef283b2204",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2025-11-17 04:47:29",
        "sys.cspResourceId": "i-0dfa1890fcefba4be",
        "sys.cspResourceName": "d4dahj2p2foc73bcejh0",
        "sys.id": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
        "sys.uid": "d4dahj2p2foc73bcejh0"
      },
      "description": "a recommended virtual machine 02 for ec21fd51-16bb-7e10-5e23-12ef283b2204",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "43.203.121.210",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.62",
      "privateDNS": "ip-10-0-1-62.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
      "specId": "aws+ap-northeast-2+t3a.xlarge",
      "cspSpecName": "t3a.xlarge",
      "imageId": "ami-010be25c3775061c9",
      "cspImageName": "ami-010be25c3775061c9",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-0330b15879aede8d2",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "subnet-00440b2d1ddf6b379",
      "networkInterface": "eni-attach-0c159cce7fb3b5358",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4dahgap2foc73bcejd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:47:38Z",
          "completedTime": "2025-11-17T04:47:41Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-62 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-11-17T04:47:12Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-020d3659c57d50830}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "2A25D144-CAA6-4CFF-B264-3AFE27484140"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:2,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-010be25c3775061c9"
        },
        {
          "key": "InstanceId",
          "value": "i-0dfa1890fcefba4be"
        },
        {
          "key": "InstanceType",
          "value": "t3a.xlarge"
        },
        {
          "key": "KeyName",
          "value": "d4dahgap2foc73bcejd0"
        },
        {
          "key": "LaunchTime",
          "value": "2025-11-17T04:47:11Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.121.210},Attachment:{AttachTime:2025-11-17T04:47:11Z,AttachmentId:eni-attach-0c159cce7fb3b5358,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0893e1a7e93bcf006,GroupName:d4dahhip2foc73bceje0}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:81:c5:56:58:31,NetworkInterfaceId:eni-0f0d37224ebe972af,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.62,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.121.210},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.62}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-00440b2d1ddf6b379,VpcId:vpc-0330b15879aede8d2}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-62.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.62"
        },
        {
          "key": "PublicIpAddress",
          "value": "43.203.121.210"
        },
        {
          "key": "RootDeviceName",
          "value": "/dev/sda1"
        },
        {
          "key": "RootDeviceType",
          "value": "ebs"
        },
        {
          "key": "SecurityGroups",
          "value": "{GroupId:sg-0893e1a7e93bcf006,GroupName:d4dahhip2foc73bceje0}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-00440b2d1ddf6b379"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d4dahj2p2foc73bcejh0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0330b15879aede8d2"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
      "uid": "d4dahj2p2foc73bcejg0",
      "cspResourceName": "d4dahj2p2foc73bcejg0",
      "cspResourceId": "i-00c25957c7fdb5cfc",
      "name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
      "subGroupId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
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
      "createdTime": "2025-11-17 04:47:29",
      "label": {
        "sourceMachineId": "ec2643f0-9388-3d97-f3a4-f387cd52696c",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2025-11-17 04:47:29",
        "sys.cspResourceId": "i-00c25957c7fdb5cfc",
        "sys.cspResourceName": "d4dahj2p2foc73bcejg0",
        "sys.id": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
        "sys.uid": "d4dahj2p2foc73bcejg0"
      },
      "description": "a recommended virtual machine 01 for ec2643f0-9388-3d97-f3a4-f387cd52696c",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "13.124.191.35",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.241",
      "privateDNS": "ip-10-0-1-241.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
      "specId": "aws+ap-northeast-2+t3a.small",
      "cspSpecName": "t3a.small",
      "imageId": "ami-010be25c3775061c9",
      "cspImageName": "ami-010be25c3775061c9",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-0330b15879aede8d2",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "subnet-00440b2d1ddf6b379",
      "networkInterface": "eni-attach-0b65eb9ad52fe76c6",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4dahgap2foc73bcejd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:47:38Z",
          "completedTime": "2025-11-17T04:47:40Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-241 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-11-17T04:47:12Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0d40b4ecdf3a18873}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "3CE7E5E0-AF45-43E8-8A97-11FDDDFE8593"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:1,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-010be25c3775061c9"
        },
        {
          "key": "InstanceId",
          "value": "i-00c25957c7fdb5cfc"
        },
        {
          "key": "InstanceType",
          "value": "t3a.small"
        },
        {
          "key": "KeyName",
          "value": "d4dahgap2foc73bcejd0"
        },
        {
          "key": "LaunchTime",
          "value": "2025-11-17T04:47:11Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.124.191.35},Attachment:{AttachTime:2025-11-17T04:47:11Z,AttachmentId:eni-attach-0b65eb9ad52fe76c6,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-08d51712b18745856,GroupName:d4dahh2p2foc73bcejdg}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:ba:38:58:8c:11,NetworkInterfaceId:eni-0a3d64de6fe9e9fed,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.241,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.124.191.35},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.241}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-00440b2d1ddf6b379,VpcId:vpc-0330b15879aede8d2}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-241.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.241"
        },
        {
          "key": "PublicIpAddress",
          "value": "13.124.191.35"
        },
        {
          "key": "RootDeviceName",
          "value": "/dev/sda1"
        },
        {
          "key": "RootDeviceType",
          "value": "ebs"
        },
        {
          "key": "SecurityGroups",
          "value": "{GroupId:sg-08d51712b18745856,GroupName:d4dahh2p2foc73bcejdg}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-00440b2d1ddf6b379"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d4dahj2p2foc73bcejg0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0330b15879aede8d2"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
      "uid": "d4dahj2p2foc73bceji0",
      "cspResourceName": "d4dahj2p2foc73bceji0",
      "cspResourceId": "i-0a5882b217e54f74f",
      "name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
      "subGroupId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
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
      "createdTime": "2025-11-17 04:47:31",
      "label": {
        "sourceMachineId": "ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2025-11-17 04:47:31",
        "sys.cspResourceId": "i-0a5882b217e54f74f",
        "sys.cspResourceName": "d4dahj2p2foc73bceji0",
        "sys.id": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
        "sys.uid": "d4dahj2p2foc73bceji0"
      },
      "description": "a recommended virtual machine 03 for ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "3.39.6.146",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.185",
      "privateDNS": "ip-10-0-1-185.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
      "specId": "aws+ap-northeast-2+t3a.large",
      "cspSpecName": "t3a.large",
      "imageId": "ami-010be25c3775061c9",
      "cspImageName": "ami-010be25c3775061c9",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-0330b15879aede8d2",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "subnet-00440b2d1ddf6b379",
      "networkInterface": "eni-attach-09ef8a8839fd87e58",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4dahgap2foc73bcejd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:47:38Z",
          "completedTime": "2025-11-17T04:47:41Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-185 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "AmiLaunchIndex",
          "value": "0"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "BlockDeviceMappings",
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-11-17T04:47:11Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0f5ce9be98cb95562}}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CapacityReservationSpecification",
          "value": "{CapacityReservationPreference:open,CapacityReservationTarget:null}"
        },
        {
          "key": "ClientToken",
          "value": "C81902E7-4F92-4804-A85D-0B2979D2493D"
        },
        {
          "key": "CpuOptions",
          "value": "{CoreCount:1,ThreadsPerCore:2}"
        },
        {
          "key": "EbsOptimized",
          "value": "false"
        },
        {
          "key": "EnaSupport",
          "value": "true"
        },
        {
          "key": "EnclaveOptions",
          "value": "{Enabled:false}"
        },
        {
          "key": "HibernationOptions",
          "value": "{Configured:false}"
        },
        {
          "key": "Hypervisor",
          "value": "xen"
        },
        {
          "key": "ImageId",
          "value": "ami-010be25c3775061c9"
        },
        {
          "key": "InstanceId",
          "value": "i-0a5882b217e54f74f"
        },
        {
          "key": "InstanceType",
          "value": "t3a.large"
        },
        {
          "key": "KeyName",
          "value": "d4dahgap2foc73bcejd0"
        },
        {
          "key": "LaunchTime",
          "value": "2025-11-17T04:47:10Z"
        },
        {
          "key": "MetadataOptions",
          "value": "{HttpEndpoint:enabled,HttpPutResponseHopLimit:1,HttpTokens:optional,State:applied}"
        },
        {
          "key": "Monitoring",
          "value": "{State:disabled}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.6.146},Attachment:{AttachTime:2025-11-17T04:47:10Z,AttachmentId:eni-attach-09ef8a8839fd87e58,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-023e0f0b2c789b8f8,GroupName:d4dahiap2foc73bcejeg}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:95:ec:03:1e:e5,NetworkInterfaceId:eni-0f94e50ba0cd6b0d5,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.185,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.39.6.146},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.185}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-00440b2d1ddf6b379,VpcId:vpc-0330b15879aede8d2}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-185.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.185"
        },
        {
          "key": "PublicIpAddress",
          "value": "3.39.6.146"
        },
        {
          "key": "RootDeviceName",
          "value": "/dev/sda1"
        },
        {
          "key": "RootDeviceType",
          "value": "ebs"
        },
        {
          "key": "SecurityGroups",
          "value": "{GroupId:sg-023e0f0b2c789b8f8,GroupName:d4dahiap2foc73bcejeg}"
        },
        {
          "key": "SourceDestCheck",
          "value": "true"
        },
        {
          "key": "State",
          "value": "{Code:16,Name:running}"
        },
        {
          "key": "SubnetId",
          "value": "subnet-00440b2d1ddf6b379"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d4dahj2p2foc73bceji0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0330b15879aede8d2"
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
        "vmId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
        "vmIp": "13.124.191.35",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-241 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
        "vmIp": "43.203.121.210",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-62 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
        "vmIp": "3.39.6.146",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-185 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "output": "Linux ip-10-0-1-62 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "43.203.121.210",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
      "testOrder": 1,
      "userName": "cb-user",
      "vmId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux ip-10-0-1-241 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "13.124.191.35",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
      "testOrder": 2,
      "userName": "cb-user",
      "vmId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux ip-10-0-1-185 6.8.0-1040-aws #42~22.04.1-Ubuntu SMP Wed Sep 24 10:26:57 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "3.39.6.146",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
      "testOrder": 3,
      "userName": "cb-user",
      "vmId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1"
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

