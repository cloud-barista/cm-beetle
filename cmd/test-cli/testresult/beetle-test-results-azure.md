# CM-Beetle test results for AZURE

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.4+ (3f15f14)
- cm-model: v0.0.14
- CB-Tumblebug: v0.11.19
- CB-Spider: v0.11.16
- CB-MapUI: v0.11.19
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: November 12, 2025
- Test Time: 11:01:16 UTC
- Test Execution: 2025-11-12 11:01:16 UTC

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for AZURE

### Test Results Summary

| Test | Endpoint | Status | Duration | Details |
|------|----------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/mci` | ✅ **PASS** | 644ms | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 4m13.417s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 7.906s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 73ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 84ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 2m21.364s | Pass |

**Overall Result**: 7/7 tests passed ✅

**Total Duration**: 8m22.943061209s

*Test executed on November 12, 2025 at 11:01:16 UTC (2025-11-12 11:01:16 UTC) using CM-Beetle automated test CLI*

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
    "csp": "azure",
    "region": "koreasouth"
  },
  "OnpremiseInfraModel": {
    "network": {
      "ipv4Networks": {
        "defaultGateways": [
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec2a4cef-a613-1856-a953-0b12211163ab"
          },
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec266012-92f5-d3bc-99a9-2a49201f5158"
          },
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec2cd540-09af-4961-c40d-c5336d4cb7e8"
          }
        ]
      },
      "ipv6Networks": {}
    },
    "servers": [
      {
        "hostname": "ip-10-0-1-184",
        "machineId": "ec2a4cef-a613-1856-a953-0b12211163ab",
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
            "macAddress": "02:63:1e:12:7b:7b",
            "ipv4CidrBlocks": [
              "10.0.1.184/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::63:1eff:fe12:7b7b/64"
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
            "destination": "fe80::63:1eff:fe12:7b7b/128",
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
        "hostname": "ip-10-0-1-5",
        "machineId": "ec266012-92f5-d3bc-99a9-2a49201f5158",
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
            "macAddress": "02:9e:62:f5:36:9f",
            "ipv4CidrBlocks": [
              "10.0.1.5/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::9e:62ff:fef5:369f/64"
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
            "destination": "fe80::9e:62ff:fef5:369f/128",
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
        "hostname": "ip-10-0-1-187",
        "machineId": "ec2cd540-09af-4961-c40d-c5336d4cb7e8",
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
            "macAddress": "02:24:9f:f8:ed:2b",
            "ipv4CidrBlocks": [
              "10.0.1.187/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::24:9fff:fef8:ed2b/64"
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
            "destination": "fe80::24:9fff:fef8:ed2b/128",
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
    "csp": "azure",
    "region": "koreasouth"
  },
  "targetVmInfra": {
    "name": "mmci01",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "a recommended multi-cloud infrastructure",
    "subGroups": [
      {
        "name": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec2a4cef-a613-1856-a953-0b12211163ab"
        },
        "description": "a recommended virtual machine 01 for ec2a4cef-a613-1856-a953-0b12211163ab",
        "connectionName": "azure-koreasouth",
        "specId": "azure+koreasouth+standard_b2als_v2",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
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
        "name": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec266012-92f5-d3bc-99a9-2a49201f5158"
        },
        "description": "a recommended virtual machine 02 for ec266012-92f5-d3bc-99a9-2a49201f5158",
        "connectionName": "azure-koreasouth",
        "specId": "azure+koreasouth+standard_b4as_v2",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
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
        "name": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec2cd540-09af-4961-c40d-c5336d4cb7e8"
        },
        "description": "a recommended virtual machine 03 for ec2cd540-09af-4961-c40d-c5336d4cb7e8",
        "connectionName": "azure-koreasouth",
        "specId": "azure+koreasouth+standard_b2as_v2",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
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
    "connectionName": "azure-koreasouth",
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
    "connectionName": "azure-koreasouth",
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
      "id": "azure+koreasouth+standard_b2als_v2",
      "uid": "d49f8pvo5uas7383uuj0",
      "cspSpecName": "Standard_B2als_v2",
      "name": "azure+koreasouth+standard_b2als_v2",
      "namespace": "system",
      "connectionName": "azure-koreasouth",
      "providerName": "azure",
      "regionName": "koreasouth",
      "regionLatitude": 35.1796,
      "regionLongitude": 129.0756,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 3.90625,
      "costPerHour": 0.0432,
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
      "rootDiskSize": "0",
      "systemLabel": "auto-gen",
      "details": [
        {
          "key": "MaxDataDiskCount",
          "value": "4"
        },
        {
          "key": "MemoryInMB",
          "value": "4096"
        },
        {
          "key": "Name",
          "value": "Standard_B2als_v2"
        },
        {
          "key": "NumberOfCores",
          "value": "2"
        },
        {
          "key": "OSDiskSizeInMB",
          "value": "1047552"
        },
        {
          "key": "ResourceDiskSizeInMB",
          "value": "0"
        },
        {
          "key": "MaxResourceVolumeMB",
          "value": "0"
        },
        {
          "key": "OSVhdSizeMB",
          "value": "1047552"
        },
        {
          "key": "vCPUs",
          "value": "2"
        },
        {
          "key": "MemoryPreservingMaintenanceSupported",
          "value": "True"
        },
        {
          "key": "HyperVGenerations",
          "value": "V1,V2"
        },
        {
          "key": "SupportedCapacityReservationTypes",
          "value": "Targeted"
        },
        {
          "key": "MemoryGB",
          "value": "4"
        },
        {
          "key": "MaxDataDiskCount",
          "value": "4"
        },
        {
          "key": "CpuArchitectureType",
          "value": "x64"
        },
        {
          "key": "LowPriorityCapable",
          "value": "True"
        },
        {
          "key": "HibernationSupported",
          "value": "True"
        },
        {
          "key": "PremiumIO",
          "value": "True"
        },
        {
          "key": "VMDeploymentTypes",
          "value": "IaaS"
        },
        {
          "key": "vCPUsConstraintsAllowed",
          "value": "1, 2"
        },
        {
          "key": "vCPUsAvailable",
          "value": "2"
        },
        {
          "key": "vCPUsPerCore",
          "value": "2"
        },
        {
          "key": "CombinedTempDiskAndCachedIOPS",
          "value": "9000"
        },
        {
          "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
          "value": "125000000"
        },
        {
          "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
          "value": "125000000"
        },
        {
          "key": "UncachedDiskIOPS",
          "value": "3750"
        },
        {
          "key": "UncachedDiskBytesPerSecond",
          "value": "85000000"
        },
        {
          "key": "EphemeralOSDiskSupported",
          "value": "False"
        },
        {
          "key": "EncryptionAtHostSupported",
          "value": "True"
        },
        {
          "key": "CapacityReservationSupported",
          "value": "True"
        },
        {
          "key": "AcceleratedNetworkingEnabled",
          "value": "True"
        },
        {
          "key": "RdmaEnabled",
          "value": "False"
        },
        {
          "key": "MaxNetworkInterfaces",
          "value": "2"
        },
        {
          "key": "UltraSSDAvailable",
          "value": "False"
        },
        {
          "key": "LocationInfo_0_Location",
          "value": "KoreaSouth"
        },
        {
          "key": "Family",
          "value": "standardBasv2Family"
        },
        {
          "key": "Tier",
          "value": "Standard"
        },
        {
          "key": "Size",
          "value": "B2als_v2"
        },
        {
          "key": "ResourceType",
          "value": "virtualMachines"
        }
      ]
    },
    {
      "id": "azure+koreasouth+standard_b4as_v2",
      "uid": "d49f8pvo5uas7383uul0",
      "cspSpecName": "Standard_B4as_v2",
      "name": "azure+koreasouth+standard_b4as_v2",
      "namespace": "system",
      "connectionName": "azure-koreasouth",
      "providerName": "azure",
      "regionName": "koreasouth",
      "regionLatitude": 35.1796,
      "regionLongitude": 129.0756,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 15.625,
      "costPerHour": 0.173,
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
      "rootDiskSize": "0",
      "systemLabel": "auto-gen",
      "details": [
        {
          "key": "MaxDataDiskCount",
          "value": "8"
        },
        {
          "key": "MemoryInMB",
          "value": "16384"
        },
        {
          "key": "Name",
          "value": "Standard_B4as_v2"
        },
        {
          "key": "NumberOfCores",
          "value": "4"
        },
        {
          "key": "OSDiskSizeInMB",
          "value": "1047552"
        },
        {
          "key": "ResourceDiskSizeInMB",
          "value": "0"
        },
        {
          "key": "MaxResourceVolumeMB",
          "value": "0"
        },
        {
          "key": "OSVhdSizeMB",
          "value": "1047552"
        },
        {
          "key": "vCPUs",
          "value": "4"
        },
        {
          "key": "MemoryPreservingMaintenanceSupported",
          "value": "True"
        },
        {
          "key": "HyperVGenerations",
          "value": "V1,V2"
        },
        {
          "key": "SupportedCapacityReservationTypes",
          "value": "Targeted"
        },
        {
          "key": "MemoryGB",
          "value": "16"
        },
        {
          "key": "MaxDataDiskCount",
          "value": "8"
        },
        {
          "key": "CpuArchitectureType",
          "value": "x64"
        },
        {
          "key": "LowPriorityCapable",
          "value": "True"
        },
        {
          "key": "HibernationSupported",
          "value": "True"
        },
        {
          "key": "PremiumIO",
          "value": "True"
        },
        {
          "key": "VMDeploymentTypes",
          "value": "IaaS"
        },
        {
          "key": "vCPUsConstraintsAllowed",
          "value": "1, 2, 4"
        },
        {
          "key": "vCPUsAvailable",
          "value": "4"
        },
        {
          "key": "vCPUsPerCore",
          "value": "2"
        },
        {
          "key": "CombinedTempDiskAndCachedIOPS",
          "value": "19000"
        },
        {
          "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
          "value": "250000000"
        },
        {
          "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
          "value": "250000000"
        },
        {
          "key": "UncachedDiskIOPS",
          "value": "6400"
        },
        {
          "key": "UncachedDiskBytesPerSecond",
          "value": "145000000"
        },
        {
          "key": "EphemeralOSDiskSupported",
          "value": "False"
        },
        {
          "key": "EncryptionAtHostSupported",
          "value": "True"
        },
        {
          "key": "CapacityReservationSupported",
          "value": "True"
        },
        {
          "key": "AcceleratedNetworkingEnabled",
          "value": "True"
        },
        {
          "key": "RdmaEnabled",
          "value": "False"
        },
        {
          "key": "MaxNetworkInterfaces",
          "value": "3"
        },
        {
          "key": "UltraSSDAvailable",
          "value": "False"
        },
        {
          "key": "LocationInfo_0_Location",
          "value": "KoreaSouth"
        },
        {
          "key": "Family",
          "value": "standardBasv2Family"
        },
        {
          "key": "Tier",
          "value": "Standard"
        },
        {
          "key": "Size",
          "value": "B4as_v2"
        },
        {
          "key": "ResourceType",
          "value": "virtualMachines"
        }
      ]
    },
    {
      "id": "azure+koreasouth+standard_b2as_v2",
      "uid": "d49f8pvo5uas7383uujg",
      "cspSpecName": "Standard_B2as_v2",
      "name": "azure+koreasouth+standard_b2as_v2",
      "namespace": "system",
      "connectionName": "azure-koreasouth",
      "providerName": "azure",
      "regionName": "koreasouth",
      "regionLatitude": 35.1796,
      "regionLongitude": 129.0756,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 7.8125,
      "costPerHour": 0.0865,
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
      "rootDiskSize": "0",
      "systemLabel": "auto-gen",
      "details": [
        {
          "key": "MaxDataDiskCount",
          "value": "4"
        },
        {
          "key": "MemoryInMB",
          "value": "8192"
        },
        {
          "key": "Name",
          "value": "Standard_B2as_v2"
        },
        {
          "key": "NumberOfCores",
          "value": "2"
        },
        {
          "key": "OSDiskSizeInMB",
          "value": "1047552"
        },
        {
          "key": "ResourceDiskSizeInMB",
          "value": "0"
        },
        {
          "key": "MaxResourceVolumeMB",
          "value": "0"
        },
        {
          "key": "OSVhdSizeMB",
          "value": "1047552"
        },
        {
          "key": "vCPUs",
          "value": "2"
        },
        {
          "key": "MemoryPreservingMaintenanceSupported",
          "value": "True"
        },
        {
          "key": "HyperVGenerations",
          "value": "V1,V2"
        },
        {
          "key": "SupportedCapacityReservationTypes",
          "value": "Targeted"
        },
        {
          "key": "MemoryGB",
          "value": "8"
        },
        {
          "key": "MaxDataDiskCount",
          "value": "4"
        },
        {
          "key": "CpuArchitectureType",
          "value": "x64"
        },
        {
          "key": "LowPriorityCapable",
          "value": "True"
        },
        {
          "key": "HibernationSupported",
          "value": "True"
        },
        {
          "key": "PremiumIO",
          "value": "True"
        },
        {
          "key": "VMDeploymentTypes",
          "value": "IaaS"
        },
        {
          "key": "vCPUsConstraintsAllowed",
          "value": "1, 2"
        },
        {
          "key": "vCPUsAvailable",
          "value": "2"
        },
        {
          "key": "vCPUsPerCore",
          "value": "2"
        },
        {
          "key": "CombinedTempDiskAndCachedIOPS",
          "value": "9000"
        },
        {
          "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
          "value": "125000000"
        },
        {
          "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
          "value": "125000000"
        },
        {
          "key": "UncachedDiskIOPS",
          "value": "3750"
        },
        {
          "key": "UncachedDiskBytesPerSecond",
          "value": "85000000"
        },
        {
          "key": "EphemeralOSDiskSupported",
          "value": "False"
        },
        {
          "key": "EncryptionAtHostSupported",
          "value": "True"
        },
        {
          "key": "CapacityReservationSupported",
          "value": "True"
        },
        {
          "key": "AcceleratedNetworkingEnabled",
          "value": "True"
        },
        {
          "key": "RdmaEnabled",
          "value": "False"
        },
        {
          "key": "MaxNetworkInterfaces",
          "value": "2"
        },
        {
          "key": "UltraSSDAvailable",
          "value": "False"
        },
        {
          "key": "LocationInfo_0_Location",
          "value": "KoreaSouth"
        },
        {
          "key": "Family",
          "value": "standardBasv2Family"
        },
        {
          "key": "Tier",
          "value": "Standard"
        },
        {
          "key": "Size",
          "value": "B2as_v2"
        },
        {
          "key": "ResourceType",
          "value": "virtualMachines"
        }
      ]
    }
  ],
  "targetVmOsImageList": [
    {
      "namespace": "system",
      "providerName": "azure",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "regionList": [
        "common"
      ],
      "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "uid": "d49faofo5uas738c8ogg",
      "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "connectionName": "azure-all",
      "infraType": "vm",
      "fetchedTime": "2025.11.11 08:35:45 Tue",
      "isKubernetesImage": true,
      "isBasicImage": true,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "0001-com-ubuntu-server-jammy:22_04-lts-gen2",
      "osDiskType": "NA",
      "osDiskSizeGB": -1,
      "imageStatus": "Available",
      "details": [
        {
          "key": "Location",
          "value": "AustraliaCentral"
        },
        {
          "key": "Name",
          "value": "22.04.202510230"
        },
        {
          "key": "ID",
          "value": "/Subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202510230"
        },
        {
          "key": "Properties",
          "value": "{architecture:x64,automaticOSUpgradeProperties:{automaticOSUpgradeSupported:false},dataDiskImages:[],disallowed:{vmDiskType:Unmanaged},features:[{name:SecurityType,value:TrustedLaunchSupported},{name:IsAcceleratedNetworkSupported,value:True},{name:DiskControllerTypes,value:SCSI, NVMe},{name:IsHibernateSupported,value:True}],hyperVGeneration:V2,imageDeprecationStatus:{imageState:Active},osDiskImage:{operatingSystem:Linux}}"
        }
      ],
      "systemLabel": "from-assets"
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "mig-sg-01",
      "connectionName": "azure-koreasouth",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for ec2a4cef-a613-1856-a953-0b12211163ab",
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
      "connectionName": "azure-koreasouth",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for ec266012-92f5-d3bc-99a9-2a49201f5158",
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
      "connectionName": "azure-koreasouth",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for ec2cd540-09af-4961-c40d-c5336d4cb7e8",
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
  "uid": "d4a6jbfo5uas73f10ld0",
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
    "sys.uid": "d4a6jbfo5uas73f10ld0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
      "uid": "d4a6jbfo5uas73f10le0",
      "cspResourceName": "d4a6jbfo5uas73f10le0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10le0",
      "name": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
      "subGroupId": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab",
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
      "createdTime": "2025-11-12 11:04:44",
      "label": {
        "sourceMachineId": "ec2a4cef-a613-1856-a953-0b12211163ab"
      },
      "description": "a recommended virtual machine 01 for ec2a4cef-a613-1856-a953-0b12211163ab",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.230.72",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
          "assignedZone": "N/A"
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0/subnets/d4a6i27o5uas73f10lag",
      "networkInterface": "d4a6jbfo5uas73f10le0-1543-VNic",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4a6i4vo5uas73f10lb0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-12T11:05:33Z",
          "completedTime": "2025-11-12T11:05:36Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4a6jbfo5uas73f10le0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4a6jbfo5uas73f10le0-1543-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4a6jbfo5uas73f10le0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8g3Dez3c16SnSB1aqhqMSIKxoh/B716TdwnZOz39iX7bhwhV5GYN0DPvc7QTqltUBqpxGFoiu7HIPPvIAY0XjVSQlSrlxrnS74F4r+yDAOq2EQRFv/DnAFtXIPl8GNC0SaCdZ6Q2T9CsUQAcWT1WyhlAT+TVntX+92tDGWG8zSpljVSyfZWcQz2rY6qZd8iCzQX4CJ0Qejf+MtoYJ/eD5Q7PFNuIw356TSXNYptaw9rUSBbaRSY8aS7FfQSNN5xx3zQXpFduQdPsycIc2pDvVrkHstnhhj2JFv8QgZXQ6l8L33KrAedi/PWGsJQ7rnvtLf+Kj2oZdRl0+RBdhn56+NcyQ7+x4tYzQjasUZCoEbl8wpNUWlEzG0E25hlxhrpu+SJKbp9NSqECtZKIQ9klZ68xJKjm+mNdyGy456xncl+QhcHoyB/5DtUPFwq69QvJZ7cMIoz72p/2QIADcbWj8I8FLheShcSx5+aTlitD1/pPsiSauBKi1SrfEPFr46rEpGkCCifDX1RRGNv6CWTSa9x/2OJOlhzxjpt5SU7iSZ57DenHLUkfoCkclGY0L8P5t16W8yOJbIPtLW9wlJgdFrD+Ursno01UIhIMVvZ/gj59ER9PpppqKBIn7Cae5VVvNEhUQRF27kmPX74P4/Tb1UA9wauDlPkFfBdikd8a7kw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4a6jbfo5uas73f10le0_disk1_27f61ad31acb453d8059fad34b213d48,storageAccountType:Premium_LRS},name:d4a6jbfo5uas73f10le0_disk1_27f61ad31acb453d8059fad34b213d48,osType:Linux}},timeCreated:2025-11-12T11:04:27.7000813Z,vmId:946a6be1-2fdc-4217-86ef-df1d887b6212}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4a6jbfo5uas73f10le0,keypair:d4a6i4vo5uas73f10lb0,publicip:d4a6jbfo5uas73f10le0-26898-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10le0"
        },
        {
          "key": "Name",
          "value": "d4a6jbfo5uas73f10le0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
      "uid": "d4a6jbfo5uas73f10lg0",
      "cspResourceName": "d4a6jbfo5uas73f10lg0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lg0",
      "name": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
      "subGroupId": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8",
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
      "createdTime": "2025-11-12 11:04:47",
      "label": {
        "sourceMachineId": "ec2cd540-09af-4961-c40d-c5336d4cb7e8"
      },
      "description": "a recommended virtual machine 03 for ec2cd540-09af-4961-c40d-c5336d4cb7e8",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.222.176",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
          "assignedZone": "N/A"
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0/subnets/d4a6i27o5uas73f10lag",
      "networkInterface": "d4a6jbfo5uas73f10lg0-51653-VNic",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4a6i4vo5uas73f10lb0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-12T11:05:33Z",
          "completedTime": "2025-11-12T11:05:37Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4a6jbfo5uas73f10lg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4a6jbfo5uas73f10lg0-51653-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4a6jbfo5uas73f10lg0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8g3Dez3c16SnSB1aqhqMSIKxoh/B716TdwnZOz39iX7bhwhV5GYN0DPvc7QTqltUBqpxGFoiu7HIPPvIAY0XjVSQlSrlxrnS74F4r+yDAOq2EQRFv/DnAFtXIPl8GNC0SaCdZ6Q2T9CsUQAcWT1WyhlAT+TVntX+92tDGWG8zSpljVSyfZWcQz2rY6qZd8iCzQX4CJ0Qejf+MtoYJ/eD5Q7PFNuIw356TSXNYptaw9rUSBbaRSY8aS7FfQSNN5xx3zQXpFduQdPsycIc2pDvVrkHstnhhj2JFv8QgZXQ6l8L33KrAedi/PWGsJQ7rnvtLf+Kj2oZdRl0+RBdhn56+NcyQ7+x4tYzQjasUZCoEbl8wpNUWlEzG0E25hlxhrpu+SJKbp9NSqECtZKIQ9klZ68xJKjm+mNdyGy456xncl+QhcHoyB/5DtUPFwq69QvJZ7cMIoz72p/2QIADcbWj8I8FLheShcSx5+aTlitD1/pPsiSauBKi1SrfEPFr46rEpGkCCifDX1RRGNv6CWTSa9x/2OJOlhzxjpt5SU7iSZ57DenHLUkfoCkclGY0L8P5t16W8yOJbIPtLW9wlJgdFrD+Ursno01UIhIMVvZ/gj59ER9PpppqKBIn7Cae5VVvNEhUQRF27kmPX74P4/Tb1UA9wauDlPkFfBdikd8a7kw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4a6jbfo5uas73f10lg0_disk1_044368e912c54a8394923cfcb5ba8c33,storageAccountType:Premium_LRS},name:d4a6jbfo5uas73f10lg0_disk1_044368e912c54a8394923cfcb5ba8c33,osType:Linux}},timeCreated:2025-11-12T11:04:29.3094535Z,vmId:c6855625-e0a7-4151-8396-5784c1b7114d}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4a6jbfo5uas73f10lg0,keypair:d4a6i4vo5uas73f10lb0,publicip:d4a6jbfo5uas73f10lg0-93048-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lg0"
        },
        {
          "key": "Name",
          "value": "d4a6jbfo5uas73f10lg0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
      "uid": "d4a6jbfo5uas73f10lf0",
      "cspResourceName": "d4a6jbfo5uas73f10lf0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lf0",
      "name": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
      "subGroupId": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158",
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
      "createdTime": "2025-11-12 11:04:44",
      "label": {
        "sourceMachineId": "ec266012-92f5-d3bc-99a9-2a49201f5158"
      },
      "description": "a recommended virtual machine 02 for ec266012-92f5-d3bc-99a9-2a49201f5158",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.222.172",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
          "assignedZone": "N/A"
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0/subnets/d4a6i27o5uas73f10lag",
      "networkInterface": "d4a6jbfo5uas73f10lf0-4074-VNic",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4a6i4vo5uas73f10lb0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-12T11:05:33Z",
          "completedTime": "2025-11-12T11:05:35Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4a6jbfo5uas73f10lf0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4a6jbfo5uas73f10lf0-4074-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4a6jbfo5uas73f10lf0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8g3Dez3c16SnSB1aqhqMSIKxoh/B716TdwnZOz39iX7bhwhV5GYN0DPvc7QTqltUBqpxGFoiu7HIPPvIAY0XjVSQlSrlxrnS74F4r+yDAOq2EQRFv/DnAFtXIPl8GNC0SaCdZ6Q2T9CsUQAcWT1WyhlAT+TVntX+92tDGWG8zSpljVSyfZWcQz2rY6qZd8iCzQX4CJ0Qejf+MtoYJ/eD5Q7PFNuIw356TSXNYptaw9rUSBbaRSY8aS7FfQSNN5xx3zQXpFduQdPsycIc2pDvVrkHstnhhj2JFv8QgZXQ6l8L33KrAedi/PWGsJQ7rnvtLf+Kj2oZdRl0+RBdhn56+NcyQ7+x4tYzQjasUZCoEbl8wpNUWlEzG0E25hlxhrpu+SJKbp9NSqECtZKIQ9klZ68xJKjm+mNdyGy456xncl+QhcHoyB/5DtUPFwq69QvJZ7cMIoz72p/2QIADcbWj8I8FLheShcSx5+aTlitD1/pPsiSauBKi1SrfEPFr46rEpGkCCifDX1RRGNv6CWTSa9x/2OJOlhzxjpt5SU7iSZ57DenHLUkfoCkclGY0L8P5t16W8yOJbIPtLW9wlJgdFrD+Ursno01UIhIMVvZ/gj59ER9PpppqKBIn7Cae5VVvNEhUQRF27kmPX74P4/Tb1UA9wauDlPkFfBdikd8a7kw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4a6jbfo5uas73f10lf0_disk1_54349fa8449f4c28af6edfbaf5dd4343,storageAccountType:Premium_LRS},name:d4a6jbfo5uas73f10lf0_disk1_54349fa8449f4c28af6edfbaf5dd4343,osType:Linux}},timeCreated:2025-11-12T11:04:28.059458Z,vmId:f6b0d8f3-c09a-43eb-86d8-732ebcbfc366}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4a6jbfo5uas73f10lf0,keypair:d4a6i4vo5uas73f10lb0,publicip:d4a6jbfo5uas73f10lf0-30399-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lf0"
        },
        {
          "key": "Name",
          "value": "d4a6jbfo5uas73f10lf0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
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
        "vmId": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
        "vmIp": "52.231.222.172",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4a6jbfo5uas73f10lf0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
        "vmIp": "52.231.230.72",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4a6jbfo5uas73f10le0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
        "vmIp": "52.231.222.176",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4a6jbfo5uas73f10lg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "d4a6jbfo5uas73f10ld0",
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
        "sys.uid": "d4a6jbfo5uas73f10ld0"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "a recommended multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
          "uid": "d4a6jbfo5uas73f10lf0",
          "cspResourceName": "d4a6jbfo5uas73f10lf0",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lf0",
          "name": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
          "subGroupId": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158",
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
          "createdTime": "2025-11-12 11:04:44",
          "label": {
            "sourceMachineId": "ec266012-92f5-d3bc-99a9-2a49201f5158",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2025-11-12 11:04:44",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lf0",
            "sys.cspResourceName": "d4a6jbfo5uas73f10lf0",
            "sys.id": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158",
            "sys.uid": "d4a6jbfo5uas73f10lf0"
          },
          "description": "a recommended virtual machine 02 for ec266012-92f5-d3bc-99a9-2a49201f5158",
          "region": {
            "Region": "koreasouth",
            "Zone": ""
          },
          "publicIP": "52.231.222.172",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.5",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": "50",
          "rootDiskName": "",
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
              "assignedZone": "N/A"
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0/subnets/d4a6i27o5uas73f10lag",
          "networkInterface": "d4a6jbfo5uas73f10lf0-4074-VNic",
          "securityGroupIds": [
            "mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4a6i4vo5uas73f10lb0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-12T11:05:33Z",
              "completedTime": "2025-11-12T11:05:35Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4a6jbfo5uas73f10lf0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4a6jbfo5uas73f10lf0-4074-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4a6jbfo5uas73f10lf0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8g3Dez3c16SnSB1aqhqMSIKxoh/B716TdwnZOz39iX7bhwhV5GYN0DPvc7QTqltUBqpxGFoiu7HIPPvIAY0XjVSQlSrlxrnS74F4r+yDAOq2EQRFv/DnAFtXIPl8GNC0SaCdZ6Q2T9CsUQAcWT1WyhlAT+TVntX+92tDGWG8zSpljVSyfZWcQz2rY6qZd8iCzQX4CJ0Qejf+MtoYJ/eD5Q7PFNuIw356TSXNYptaw9rUSBbaRSY8aS7FfQSNN5xx3zQXpFduQdPsycIc2pDvVrkHstnhhj2JFv8QgZXQ6l8L33KrAedi/PWGsJQ7rnvtLf+Kj2oZdRl0+RBdhn56+NcyQ7+x4tYzQjasUZCoEbl8wpNUWlEzG0E25hlxhrpu+SJKbp9NSqECtZKIQ9klZ68xJKjm+mNdyGy456xncl+QhcHoyB/5DtUPFwq69QvJZ7cMIoz72p/2QIADcbWj8I8FLheShcSx5+aTlitD1/pPsiSauBKi1SrfEPFr46rEpGkCCifDX1RRGNv6CWTSa9x/2OJOlhzxjpt5SU7iSZ57DenHLUkfoCkclGY0L8P5t16W8yOJbIPtLW9wlJgdFrD+Ursno01UIhIMVvZ/gj59ER9PpppqKBIn7Cae5VVvNEhUQRF27kmPX74P4/Tb1UA9wauDlPkFfBdikd8a7kw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4a6jbfo5uas73f10lf0_disk1_54349fa8449f4c28af6edfbaf5dd4343,storageAccountType:Premium_LRS},name:d4a6jbfo5uas73f10lf0_disk1_54349fa8449f4c28af6edfbaf5dd4343,osType:Linux}},timeCreated:2025-11-12T11:04:28.059458Z,vmId:f6b0d8f3-c09a-43eb-86d8-732ebcbfc366}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d4a6jbfo5uas73f10lf0,keypair:d4a6i4vo5uas73f10lb0,publicip:d4a6jbfo5uas73f10lf0-30399-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lf0"
            },
            {
              "key": "Name",
              "value": "d4a6jbfo5uas73f10lf0"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
          "uid": "d4a6jbfo5uas73f10le0",
          "cspResourceName": "d4a6jbfo5uas73f10le0",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10le0",
          "name": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
          "subGroupId": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab",
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
          "createdTime": "2025-11-12 11:04:44",
          "label": {
            "sourceMachineId": "ec2a4cef-a613-1856-a953-0b12211163ab",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2025-11-12 11:04:44",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10le0",
            "sys.cspResourceName": "d4a6jbfo5uas73f10le0",
            "sys.id": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab",
            "sys.uid": "d4a6jbfo5uas73f10le0"
          },
          "description": "a recommended virtual machine 01 for ec2a4cef-a613-1856-a953-0b12211163ab",
          "region": {
            "Region": "koreasouth",
            "Zone": ""
          },
          "publicIP": "52.231.230.72",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": "50",
          "rootDiskName": "",
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
              "assignedZone": "N/A"
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0/subnets/d4a6i27o5uas73f10lag",
          "networkInterface": "d4a6jbfo5uas73f10le0-1543-VNic",
          "securityGroupIds": [
            "mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4a6i4vo5uas73f10lb0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-12T11:05:33Z",
              "completedTime": "2025-11-12T11:05:36Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4a6jbfo5uas73f10le0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4a6jbfo5uas73f10le0-1543-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4a6jbfo5uas73f10le0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8g3Dez3c16SnSB1aqhqMSIKxoh/B716TdwnZOz39iX7bhwhV5GYN0DPvc7QTqltUBqpxGFoiu7HIPPvIAY0XjVSQlSrlxrnS74F4r+yDAOq2EQRFv/DnAFtXIPl8GNC0SaCdZ6Q2T9CsUQAcWT1WyhlAT+TVntX+92tDGWG8zSpljVSyfZWcQz2rY6qZd8iCzQX4CJ0Qejf+MtoYJ/eD5Q7PFNuIw356TSXNYptaw9rUSBbaRSY8aS7FfQSNN5xx3zQXpFduQdPsycIc2pDvVrkHstnhhj2JFv8QgZXQ6l8L33KrAedi/PWGsJQ7rnvtLf+Kj2oZdRl0+RBdhn56+NcyQ7+x4tYzQjasUZCoEbl8wpNUWlEzG0E25hlxhrpu+SJKbp9NSqECtZKIQ9klZ68xJKjm+mNdyGy456xncl+QhcHoyB/5DtUPFwq69QvJZ7cMIoz72p/2QIADcbWj8I8FLheShcSx5+aTlitD1/pPsiSauBKi1SrfEPFr46rEpGkCCifDX1RRGNv6CWTSa9x/2OJOlhzxjpt5SU7iSZ57DenHLUkfoCkclGY0L8P5t16W8yOJbIPtLW9wlJgdFrD+Ursno01UIhIMVvZ/gj59ER9PpppqKBIn7Cae5VVvNEhUQRF27kmPX74P4/Tb1UA9wauDlPkFfBdikd8a7kw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4a6jbfo5uas73f10le0_disk1_27f61ad31acb453d8059fad34b213d48,storageAccountType:Premium_LRS},name:d4a6jbfo5uas73f10le0_disk1_27f61ad31acb453d8059fad34b213d48,osType:Linux}},timeCreated:2025-11-12T11:04:27.7000813Z,vmId:946a6be1-2fdc-4217-86ef-df1d887b6212}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d4a6jbfo5uas73f10le0,keypair:d4a6i4vo5uas73f10lb0,publicip:d4a6jbfo5uas73f10le0-26898-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10le0"
            },
            {
              "key": "Name",
              "value": "d4a6jbfo5uas73f10le0"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
          "uid": "d4a6jbfo5uas73f10lg0",
          "cspResourceName": "d4a6jbfo5uas73f10lg0",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lg0",
          "name": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
          "subGroupId": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8",
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
          "createdTime": "2025-11-12 11:04:47",
          "label": {
            "sourceMachineId": "ec2cd540-09af-4961-c40d-c5336d4cb7e8",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2025-11-12 11:04:47",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lg0",
            "sys.cspResourceName": "d4a6jbfo5uas73f10lg0",
            "sys.id": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8",
            "sys.uid": "d4a6jbfo5uas73f10lg0"
          },
          "description": "a recommended virtual machine 03 for ec2cd540-09af-4961-c40d-c5336d4cb7e8",
          "region": {
            "Region": "koreasouth",
            "Zone": ""
          },
          "publicIP": "52.231.222.176",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": "50",
          "rootDiskName": "",
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
              "assignedZone": "N/A"
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0/subnets/d4a6i27o5uas73f10lag",
          "networkInterface": "d4a6jbfo5uas73f10lg0-51653-VNic",
          "securityGroupIds": [
            "mig-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4a6i4vo5uas73f10lb0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-12T11:05:33Z",
              "completedTime": "2025-11-12T11:05:37Z",
              "elapsedTime": 4,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4a6jbfo5uas73f10lg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4a6jbfo5uas73f10lg0-51653-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4a6jbfo5uas73f10lg0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8g3Dez3c16SnSB1aqhqMSIKxoh/B716TdwnZOz39iX7bhwhV5GYN0DPvc7QTqltUBqpxGFoiu7HIPPvIAY0XjVSQlSrlxrnS74F4r+yDAOq2EQRFv/DnAFtXIPl8GNC0SaCdZ6Q2T9CsUQAcWT1WyhlAT+TVntX+92tDGWG8zSpljVSyfZWcQz2rY6qZd8iCzQX4CJ0Qejf+MtoYJ/eD5Q7PFNuIw356TSXNYptaw9rUSBbaRSY8aS7FfQSNN5xx3zQXpFduQdPsycIc2pDvVrkHstnhhj2JFv8QgZXQ6l8L33KrAedi/PWGsJQ7rnvtLf+Kj2oZdRl0+RBdhn56+NcyQ7+x4tYzQjasUZCoEbl8wpNUWlEzG0E25hlxhrpu+SJKbp9NSqECtZKIQ9klZ68xJKjm+mNdyGy456xncl+QhcHoyB/5DtUPFwq69QvJZ7cMIoz72p/2QIADcbWj8I8FLheShcSx5+aTlitD1/pPsiSauBKi1SrfEPFr46rEpGkCCifDX1RRGNv6CWTSa9x/2OJOlhzxjpt5SU7iSZ57DenHLUkfoCkclGY0L8P5t16W8yOJbIPtLW9wlJgdFrD+Ursno01UIhIMVvZ/gj59ER9PpppqKBIn7Cae5VVvNEhUQRF27kmPX74P4/Tb1UA9wauDlPkFfBdikd8a7kw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4a6jbfo5uas73f10lg0_disk1_044368e912c54a8394923cfcb5ba8c33,storageAccountType:Premium_LRS},name:d4a6jbfo5uas73f10lg0_disk1_044368e912c54a8394923cfcb5ba8c33,osType:Linux}},timeCreated:2025-11-12T11:04:29.3094535Z,vmId:c6855625-e0a7-4151-8396-5784c1b7114d}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d4a6jbfo5uas73f10lg0,keypair:d4a6i4vo5uas73f10lb0,publicip:d4a6jbfo5uas73f10lg0-93048-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lg0"
            },
            {
              "key": "Name",
              "value": "d4a6jbfo5uas73f10lg0"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
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
            "vmId": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
            "vmIp": "52.231.222.172",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4a6jbfo5uas73f10lf0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
            "vmIp": "52.231.230.72",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4a6jbfo5uas73f10le0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
            "vmIp": "52.231.222.176",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4a6jbfo5uas73f10lg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "d4a6jbfo5uas73f10ld0",
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
    "sys.uid": "d4a6jbfo5uas73f10ld0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
      "uid": "d4a6jbfo5uas73f10lf0",
      "cspResourceName": "d4a6jbfo5uas73f10lf0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lf0",
      "name": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
      "subGroupId": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158",
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
      "createdTime": "2025-11-12 11:04:44",
      "label": {
        "sourceMachineId": "ec266012-92f5-d3bc-99a9-2a49201f5158",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2025-11-12 11:04:44",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lf0",
        "sys.cspResourceName": "d4a6jbfo5uas73f10lf0",
        "sys.id": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158",
        "sys.uid": "d4a6jbfo5uas73f10lf0"
      },
      "description": "a recommended virtual machine 02 for ec266012-92f5-d3bc-99a9-2a49201f5158",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.222.172",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
          "assignedZone": "N/A"
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0/subnets/d4a6i27o5uas73f10lag",
      "networkInterface": "d4a6jbfo5uas73f10lf0-4074-VNic",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4a6i4vo5uas73f10lb0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-12T11:05:33Z",
          "completedTime": "2025-11-12T11:05:35Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4a6jbfo5uas73f10lf0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4a6jbfo5uas73f10lf0-4074-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4a6jbfo5uas73f10lf0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8g3Dez3c16SnSB1aqhqMSIKxoh/B716TdwnZOz39iX7bhwhV5GYN0DPvc7QTqltUBqpxGFoiu7HIPPvIAY0XjVSQlSrlxrnS74F4r+yDAOq2EQRFv/DnAFtXIPl8GNC0SaCdZ6Q2T9CsUQAcWT1WyhlAT+TVntX+92tDGWG8zSpljVSyfZWcQz2rY6qZd8iCzQX4CJ0Qejf+MtoYJ/eD5Q7PFNuIw356TSXNYptaw9rUSBbaRSY8aS7FfQSNN5xx3zQXpFduQdPsycIc2pDvVrkHstnhhj2JFv8QgZXQ6l8L33KrAedi/PWGsJQ7rnvtLf+Kj2oZdRl0+RBdhn56+NcyQ7+x4tYzQjasUZCoEbl8wpNUWlEzG0E25hlxhrpu+SJKbp9NSqECtZKIQ9klZ68xJKjm+mNdyGy456xncl+QhcHoyB/5DtUPFwq69QvJZ7cMIoz72p/2QIADcbWj8I8FLheShcSx5+aTlitD1/pPsiSauBKi1SrfEPFr46rEpGkCCifDX1RRGNv6CWTSa9x/2OJOlhzxjpt5SU7iSZ57DenHLUkfoCkclGY0L8P5t16W8yOJbIPtLW9wlJgdFrD+Ursno01UIhIMVvZ/gj59ER9PpppqKBIn7Cae5VVvNEhUQRF27kmPX74P4/Tb1UA9wauDlPkFfBdikd8a7kw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4a6jbfo5uas73f10lf0_disk1_54349fa8449f4c28af6edfbaf5dd4343,storageAccountType:Premium_LRS},name:d4a6jbfo5uas73f10lf0_disk1_54349fa8449f4c28af6edfbaf5dd4343,osType:Linux}},timeCreated:2025-11-12T11:04:28.059458Z,vmId:f6b0d8f3-c09a-43eb-86d8-732ebcbfc366}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4a6jbfo5uas73f10lf0,keypair:d4a6i4vo5uas73f10lb0,publicip:d4a6jbfo5uas73f10lf0-30399-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lf0"
        },
        {
          "key": "Name",
          "value": "d4a6jbfo5uas73f10lf0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
      "uid": "d4a6jbfo5uas73f10le0",
      "cspResourceName": "d4a6jbfo5uas73f10le0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10le0",
      "name": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
      "subGroupId": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab",
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
      "createdTime": "2025-11-12 11:04:44",
      "label": {
        "sourceMachineId": "ec2a4cef-a613-1856-a953-0b12211163ab",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2025-11-12 11:04:44",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10le0",
        "sys.cspResourceName": "d4a6jbfo5uas73f10le0",
        "sys.id": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab",
        "sys.uid": "d4a6jbfo5uas73f10le0"
      },
      "description": "a recommended virtual machine 01 for ec2a4cef-a613-1856-a953-0b12211163ab",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.230.72",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
          "assignedZone": "N/A"
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0/subnets/d4a6i27o5uas73f10lag",
      "networkInterface": "d4a6jbfo5uas73f10le0-1543-VNic",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4a6i4vo5uas73f10lb0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-12T11:05:33Z",
          "completedTime": "2025-11-12T11:05:36Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4a6jbfo5uas73f10le0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4a6jbfo5uas73f10le0-1543-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4a6jbfo5uas73f10le0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8g3Dez3c16SnSB1aqhqMSIKxoh/B716TdwnZOz39iX7bhwhV5GYN0DPvc7QTqltUBqpxGFoiu7HIPPvIAY0XjVSQlSrlxrnS74F4r+yDAOq2EQRFv/DnAFtXIPl8GNC0SaCdZ6Q2T9CsUQAcWT1WyhlAT+TVntX+92tDGWG8zSpljVSyfZWcQz2rY6qZd8iCzQX4CJ0Qejf+MtoYJ/eD5Q7PFNuIw356TSXNYptaw9rUSBbaRSY8aS7FfQSNN5xx3zQXpFduQdPsycIc2pDvVrkHstnhhj2JFv8QgZXQ6l8L33KrAedi/PWGsJQ7rnvtLf+Kj2oZdRl0+RBdhn56+NcyQ7+x4tYzQjasUZCoEbl8wpNUWlEzG0E25hlxhrpu+SJKbp9NSqECtZKIQ9klZ68xJKjm+mNdyGy456xncl+QhcHoyB/5DtUPFwq69QvJZ7cMIoz72p/2QIADcbWj8I8FLheShcSx5+aTlitD1/pPsiSauBKi1SrfEPFr46rEpGkCCifDX1RRGNv6CWTSa9x/2OJOlhzxjpt5SU7iSZ57DenHLUkfoCkclGY0L8P5t16W8yOJbIPtLW9wlJgdFrD+Ursno01UIhIMVvZ/gj59ER9PpppqKBIn7Cae5VVvNEhUQRF27kmPX74P4/Tb1UA9wauDlPkFfBdikd8a7kw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4a6jbfo5uas73f10le0_disk1_27f61ad31acb453d8059fad34b213d48,storageAccountType:Premium_LRS},name:d4a6jbfo5uas73f10le0_disk1_27f61ad31acb453d8059fad34b213d48,osType:Linux}},timeCreated:2025-11-12T11:04:27.7000813Z,vmId:946a6be1-2fdc-4217-86ef-df1d887b6212}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4a6jbfo5uas73f10le0,keypair:d4a6i4vo5uas73f10lb0,publicip:d4a6jbfo5uas73f10le0-26898-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10le0"
        },
        {
          "key": "Name",
          "value": "d4a6jbfo5uas73f10le0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
      "uid": "d4a6jbfo5uas73f10lg0",
      "cspResourceName": "d4a6jbfo5uas73f10lg0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lg0",
      "name": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
      "subGroupId": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8",
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
      "createdTime": "2025-11-12 11:04:47",
      "label": {
        "sourceMachineId": "ec2cd540-09af-4961-c40d-c5336d4cb7e8",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2025-11-12 11:04:47",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lg0",
        "sys.cspResourceName": "d4a6jbfo5uas73f10lg0",
        "sys.id": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8",
        "sys.uid": "d4a6jbfo5uas73f10lg0"
      },
      "description": "a recommended virtual machine 03 for ec2cd540-09af-4961-c40d-c5336d4cb7e8",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.222.176",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
          "assignedZone": "N/A"
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4a6i27o5uas73f10la0/subnets/d4a6i27o5uas73f10lag",
      "networkInterface": "d4a6jbfo5uas73f10lg0-51653-VNic",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4a6i4vo5uas73f10lb0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-12T11:05:33Z",
          "completedTime": "2025-11-12T11:05:37Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4a6jbfo5uas73f10lg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4a6jbfo5uas73f10lg0-51653-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4a6jbfo5uas73f10lg0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC8g3Dez3c16SnSB1aqhqMSIKxoh/B716TdwnZOz39iX7bhwhV5GYN0DPvc7QTqltUBqpxGFoiu7HIPPvIAY0XjVSQlSrlxrnS74F4r+yDAOq2EQRFv/DnAFtXIPl8GNC0SaCdZ6Q2T9CsUQAcWT1WyhlAT+TVntX+92tDGWG8zSpljVSyfZWcQz2rY6qZd8iCzQX4CJ0Qejf+MtoYJ/eD5Q7PFNuIw356TSXNYptaw9rUSBbaRSY8aS7FfQSNN5xx3zQXpFduQdPsycIc2pDvVrkHstnhhj2JFv8QgZXQ6l8L33KrAedi/PWGsJQ7rnvtLf+Kj2oZdRl0+RBdhn56+NcyQ7+x4tYzQjasUZCoEbl8wpNUWlEzG0E25hlxhrpu+SJKbp9NSqECtZKIQ9klZ68xJKjm+mNdyGy456xncl+QhcHoyB/5DtUPFwq69QvJZ7cMIoz72p/2QIADcbWj8I8FLheShcSx5+aTlitD1/pPsiSauBKi1SrfEPFr46rEpGkCCifDX1RRGNv6CWTSa9x/2OJOlhzxjpt5SU7iSZ57DenHLUkfoCkclGY0L8P5t16W8yOJbIPtLW9wlJgdFrD+Ursno01UIhIMVvZ/gj59ER9PpppqKBIn7Cae5VVvNEhUQRF27kmPX74P4/Tb1UA9wauDlPkFfBdikd8a7kw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4a6jbfo5uas73f10lg0_disk1_044368e912c54a8394923cfcb5ba8c33,storageAccountType:Premium_LRS},name:d4a6jbfo5uas73f10lg0_disk1_044368e912c54a8394923cfcb5ba8c33,osType:Linux}},timeCreated:2025-11-12T11:04:29.3094535Z,vmId:c6855625-e0a7-4151-8396-5784c1b7114d}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4a6jbfo5uas73f10lg0,keypair:d4a6i4vo5uas73f10lb0,publicip:d4a6jbfo5uas73f10lg0-93048-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4a6jbfo5uas73f10lg0"
        },
        {
          "key": "Name",
          "value": "d4a6jbfo5uas73f10lg0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
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
        "vmId": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1",
        "vmIp": "52.231.222.172",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4a6jbfo5uas73f10lf0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1",
        "vmIp": "52.231.230.72",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4a6jbfo5uas73f10le0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1",
        "vmIp": "52.231.222.176",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4a6jbfo5uas73f10lg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "output": "Linux d4a6jbfo5uas73f10lf0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.222.172",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158",
      "testOrder": 1,
      "userName": "cb-user",
      "vmId": "migrated-ec266012-92f5-d3bc-99a9-2a49201f5158-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux d4a6jbfo5uas73f10le0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.230.72",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab",
      "testOrder": 2,
      "userName": "cb-user",
      "vmId": "migrated-ec2a4cef-a613-1856-a953-0b12211163ab-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux d4a6jbfo5uas73f10lg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.222.176",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8",
      "testOrder": 3,
      "userName": "cb-user",
      "vmId": "migrated-ec2cd540-09af-4961-c40d-c5336d4cb7e8-1"
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

