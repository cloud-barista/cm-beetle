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
- Test Date: November 17, 2025
- Test Time: 04:50:41 UTC
- Test Execution: 2025-11-17 04:50:41 UTC

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
| 1 | `POST /beetle/recommendation/mci` | ✅ **PASS** | 649ms | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 4m28.23s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 2.638s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 72ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 84ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 2m8.72s | Pass |

**Overall Result**: 7/7 tests passed ✅

**Total Duration**: 8m19.221685514s

*Test executed on November 17, 2025 at 04:50:41 UTC (2025-11-17 04:50:41 UTC) using CM-Beetle automated test CLI*

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
        "name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec2643f0-9388-3d97-f3a4-f387cd52696c"
        },
        "description": "a recommended virtual machine 01 for ec2643f0-9388-3d97-f3a4-f387cd52696c",
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
        "name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec21fd51-16bb-7e10-5e23-12ef283b2204"
        },
        "description": "a recommended virtual machine 02 for ec21fd51-16bb-7e10-5e23-12ef283b2204",
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
        "name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec2876a6-3c84-7e62-aaf9-c3203f12e0b8"
        },
        "description": "a recommended virtual machine 03 for ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
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
      "uid": "d4d80m2p2foc73e8b3v0",
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
      "uid": "d4d80m2p2foc73e8b410",
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
      "uid": "d4d80m2p2foc73e8b3vg",
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
      "uid": "d4d82kap2foc73egkejg",
      "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "connectionName": "azure-all",
      "infraType": "vm",
      "fetchedTime": "2025.11.17 01:58:41 Mon",
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
      "connectionName": "azure-koreasouth",
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
      "connectionName": "azure-koreasouth",
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
  "uid": "d4dakiip2foc73bcejn0",
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
    "sys.uid": "d4dakiip2foc73bcejn0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
      "uid": "d4dakiip2foc73bcejp0",
      "cspResourceName": "d4dakiip2foc73bcejp0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejp0",
      "name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
      "subGroupId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
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
      "createdTime": "2025-11-17 04:54:31",
      "label": {
        "sourceMachineId": "ec21fd51-16bb-7e10-5e23-12ef283b2204"
      },
      "description": "a recommended virtual machine 02 for ec21fd51-16bb-7e10-5e23-12ef283b2204",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.224.138",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0/subnets/d4dajbap2foc73bcejkg",
      "networkInterface": "d4dakiip2foc73bcejp0-3040-VNic",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4daje2p2foc73bcejl0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:55:16Z",
          "completedTime": "2025-11-17T04:55:18Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4dakiip2foc73bcejp0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4dakiip2foc73bcejp0-3040-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4dakiip2foc73bcejp0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCVOfrQgxYonLFyITWLcesa5oe5/61JY3WEAK8xFzmdg0NgDnM7ujyxXPKtQwpgaP+Yk7eWtLomfPkQhD77pjB5kOG1SnPXN/Q4sEcsx7ESh0lD123BN4F05WHM/HCHMQixf+3p8m32X3MDEMmkb0wtaEQ+AKfLFevCbkc8QTi5vb7oAwFuPWfMRerAHcFHkbFMbLYHdp3YANIzLHSyXfTtuIIk2qs9aFBQLUQ0JrxrxVn1c496i3jE8up4GoWVfupFypNmEfAWuRU951QvBCbH50lEKaHhin0GJG22UOjvM+4w1xp6Jx09vhiqqJR2pnbRDltP7k9nuwa3uZ9oSdHsvzyJxnHqJhWnVTkbLbTxbLB8dChDr98jeja3VOYvGB7ikAPTtxpuGrU/SiDHruUSseDSk0xJZ8+6utItex3vvac3YsglJqktHyPVhx5ErI8WekNImq008iJ75CZ5BV/3OggfLeaAa10jyuoU0IUAsVRUqXKdwXOiG8B66FWXB3Zch4Y086niqb+im7I04d2/Y3SsEnHASzw6r2vEV2nkJlQ1xgas8EijRb7LFxrIVnYVza2STSY3h+O0UJwl9f3EFXsOjZOtiEuvk2yX8nRRYnvceies5Jw+XVhwOfIA+2LGlBhLfeSz6pVNzGiTEfucRVm3aFlNOOhITbt7A1pQ4Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4dakiip2foc73bcejp0_OsDisk_1_f83e6cd3d4764fca89965f193fb27a3c,storageAccountType:Premium_LRS},name:d4dakiip2foc73bcejp0_OsDisk_1_f83e6cd3d4764fca89965f193fb27a3c,osType:Linux}},timeCreated:2025-11-17T04:53:43.2838588Z,vmId:b23d5c8c-70da-474d-a348-3223b667c494}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4dakiip2foc73bcejp0,keypair:d4daje2p2foc73bcejl0,publicip:d4dakiip2foc73bcejp0-13587-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejp0"
        },
        {
          "key": "Name",
          "value": "d4dakiip2foc73bcejp0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
      "uid": "d4dakiip2foc73bcejq0",
      "cspResourceName": "d4dakiip2foc73bcejq0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejq0",
      "name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
      "subGroupId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
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
      "createdTime": "2025-11-17 04:54:32",
      "label": {
        "sourceMachineId": "ec2876a6-3c84-7e62-aaf9-c3203f12e0b8"
      },
      "description": "a recommended virtual machine 03 for ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.224.164",
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
      "specId": "azure+koreasouth+standard_b2as_v2",
      "cspSpecName": "Standard_B2as_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0/subnets/d4dajbap2foc73bcejkg",
      "networkInterface": "d4dakiip2foc73bcejq0-14462-VNic",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4daje2p2foc73bcejl0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:55:16Z",
          "completedTime": "2025-11-17T04:55:20Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4dakiip2foc73bcejq0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4dakiip2foc73bcejq0-14462-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4dakiip2foc73bcejq0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCVOfrQgxYonLFyITWLcesa5oe5/61JY3WEAK8xFzmdg0NgDnM7ujyxXPKtQwpgaP+Yk7eWtLomfPkQhD77pjB5kOG1SnPXN/Q4sEcsx7ESh0lD123BN4F05WHM/HCHMQixf+3p8m32X3MDEMmkb0wtaEQ+AKfLFevCbkc8QTi5vb7oAwFuPWfMRerAHcFHkbFMbLYHdp3YANIzLHSyXfTtuIIk2qs9aFBQLUQ0JrxrxVn1c496i3jE8up4GoWVfupFypNmEfAWuRU951QvBCbH50lEKaHhin0GJG22UOjvM+4w1xp6Jx09vhiqqJR2pnbRDltP7k9nuwa3uZ9oSdHsvzyJxnHqJhWnVTkbLbTxbLB8dChDr98jeja3VOYvGB7ikAPTtxpuGrU/SiDHruUSseDSk0xJZ8+6utItex3vvac3YsglJqktHyPVhx5ErI8WekNImq008iJ75CZ5BV/3OggfLeaAa10jyuoU0IUAsVRUqXKdwXOiG8B66FWXB3Zch4Y086niqb+im7I04d2/Y3SsEnHASzw6r2vEV2nkJlQ1xgas8EijRb7LFxrIVnYVza2STSY3h+O0UJwl9f3EFXsOjZOtiEuvk2yX8nRRYnvceies5Jw+XVhwOfIA+2LGlBhLfeSz6pVNzGiTEfucRVm3aFlNOOhITbt7A1pQ4Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4dakiip2foc73bcejq0_OsDisk_1_dfdf78a08cc24954b9e1d1d4e24b5ce5,storageAccountType:Premium_LRS},name:d4dakiip2foc73bcejq0_OsDisk_1_dfdf78a08cc24954b9e1d1d4e24b5ce5,osType:Linux}},timeCreated:2025-11-17T04:53:45.7839512Z,vmId:dac8bf16-7235-45df-b652-4c723ba9fc6f}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4dakiip2foc73bcejq0,keypair:d4daje2p2foc73bcejl0,publicip:d4dakiip2foc73bcejq0-49659-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejq0"
        },
        {
          "key": "Name",
          "value": "d4dakiip2foc73bcejq0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
      "uid": "d4dakiip2foc73bcejo0",
      "cspResourceName": "d4dakiip2foc73bcejo0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejo0",
      "name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
      "subGroupId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
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
      "createdTime": "2025-11-17 04:54:33",
      "label": {
        "sourceMachineId": "ec2643f0-9388-3d97-f3a4-f387cd52696c"
      },
      "description": "a recommended virtual machine 01 for ec2643f0-9388-3d97-f3a4-f387cd52696c",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.224.191",
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
      "specId": "azure+koreasouth+standard_b2als_v2",
      "cspSpecName": "Standard_B2als_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0/subnets/d4dajbap2foc73bcejkg",
      "networkInterface": "d4dakiip2foc73bcejo0-33099-VNic",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4daje2p2foc73bcejl0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:55:16Z",
          "completedTime": "2025-11-17T04:55:19Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4dakiip2foc73bcejo0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4dakiip2foc73bcejo0-33099-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4dakiip2foc73bcejo0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCVOfrQgxYonLFyITWLcesa5oe5/61JY3WEAK8xFzmdg0NgDnM7ujyxXPKtQwpgaP+Yk7eWtLomfPkQhD77pjB5kOG1SnPXN/Q4sEcsx7ESh0lD123BN4F05WHM/HCHMQixf+3p8m32X3MDEMmkb0wtaEQ+AKfLFevCbkc8QTi5vb7oAwFuPWfMRerAHcFHkbFMbLYHdp3YANIzLHSyXfTtuIIk2qs9aFBQLUQ0JrxrxVn1c496i3jE8up4GoWVfupFypNmEfAWuRU951QvBCbH50lEKaHhin0GJG22UOjvM+4w1xp6Jx09vhiqqJR2pnbRDltP7k9nuwa3uZ9oSdHsvzyJxnHqJhWnVTkbLbTxbLB8dChDr98jeja3VOYvGB7ikAPTtxpuGrU/SiDHruUSseDSk0xJZ8+6utItex3vvac3YsglJqktHyPVhx5ErI8WekNImq008iJ75CZ5BV/3OggfLeaAa10jyuoU0IUAsVRUqXKdwXOiG8B66FWXB3Zch4Y086niqb+im7I04d2/Y3SsEnHASzw6r2vEV2nkJlQ1xgas8EijRb7LFxrIVnYVza2STSY3h+O0UJwl9f3EFXsOjZOtiEuvk2yX8nRRYnvceies5Jw+XVhwOfIA+2LGlBhLfeSz6pVNzGiTEfucRVm3aFlNOOhITbt7A1pQ4Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4dakiip2foc73bcejo0_OsDisk_1_23108f5f194142beace6b397d428798a,storageAccountType:Premium_LRS},name:d4dakiip2foc73bcejo0_OsDisk_1_23108f5f194142beace6b397d428798a,osType:Linux}},timeCreated:2025-11-17T04:53:46.2683236Z,vmId:91adf7d7-0973-4f01-b9ad-542b400abac8}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4dakiip2foc73bcejo0,keypair:d4daje2p2foc73bcejl0,publicip:d4dakiip2foc73bcejo0-10794-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejo0"
        },
        {
          "key": "Name",
          "value": "d4dakiip2foc73bcejo0"
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
        "vmId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
        "vmIp": "52.231.224.138",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4dakiip2foc73bcejp0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
        "vmIp": "52.231.224.191",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4dakiip2foc73bcejo0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
        "vmIp": "52.231.224.164",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4dakiip2foc73bcejq0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "d4dakiip2foc73bcejn0",
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
        "sys.uid": "d4dakiip2foc73bcejn0"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "a recommended multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
          "uid": "d4dakiip2foc73bcejp0",
          "cspResourceName": "d4dakiip2foc73bcejp0",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejp0",
          "name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
          "subGroupId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
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
          "createdTime": "2025-11-17 04:54:31",
          "label": {
            "sourceMachineId": "ec21fd51-16bb-7e10-5e23-12ef283b2204",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2025-11-17 04:54:31",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejp0",
            "sys.cspResourceName": "d4dakiip2foc73bcejp0",
            "sys.id": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
            "sys.uid": "d4dakiip2foc73bcejp0"
          },
          "description": "a recommended virtual machine 02 for ec21fd51-16bb-7e10-5e23-12ef283b2204",
          "region": {
            "Region": "koreasouth",
            "Zone": ""
          },
          "publicIP": "52.231.224.138",
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
          "specId": "azure+koreasouth+standard_b4as_v2",
          "cspSpecName": "Standard_B4as_v2",
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0/subnets/d4dajbap2foc73bcejkg",
          "networkInterface": "d4dakiip2foc73bcejp0-3040-VNic",
          "securityGroupIds": [
            "mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4daje2p2foc73bcejl0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-17T04:55:16Z",
              "completedTime": "2025-11-17T04:55:18Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4dakiip2foc73bcejp0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4dakiip2foc73bcejp0-3040-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4dakiip2foc73bcejp0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCVOfrQgxYonLFyITWLcesa5oe5/61JY3WEAK8xFzmdg0NgDnM7ujyxXPKtQwpgaP+Yk7eWtLomfPkQhD77pjB5kOG1SnPXN/Q4sEcsx7ESh0lD123BN4F05WHM/HCHMQixf+3p8m32X3MDEMmkb0wtaEQ+AKfLFevCbkc8QTi5vb7oAwFuPWfMRerAHcFHkbFMbLYHdp3YANIzLHSyXfTtuIIk2qs9aFBQLUQ0JrxrxVn1c496i3jE8up4GoWVfupFypNmEfAWuRU951QvBCbH50lEKaHhin0GJG22UOjvM+4w1xp6Jx09vhiqqJR2pnbRDltP7k9nuwa3uZ9oSdHsvzyJxnHqJhWnVTkbLbTxbLB8dChDr98jeja3VOYvGB7ikAPTtxpuGrU/SiDHruUSseDSk0xJZ8+6utItex3vvac3YsglJqktHyPVhx5ErI8WekNImq008iJ75CZ5BV/3OggfLeaAa10jyuoU0IUAsVRUqXKdwXOiG8B66FWXB3Zch4Y086niqb+im7I04d2/Y3SsEnHASzw6r2vEV2nkJlQ1xgas8EijRb7LFxrIVnYVza2STSY3h+O0UJwl9f3EFXsOjZOtiEuvk2yX8nRRYnvceies5Jw+XVhwOfIA+2LGlBhLfeSz6pVNzGiTEfucRVm3aFlNOOhITbt7A1pQ4Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4dakiip2foc73bcejp0_OsDisk_1_f83e6cd3d4764fca89965f193fb27a3c,storageAccountType:Premium_LRS},name:d4dakiip2foc73bcejp0_OsDisk_1_f83e6cd3d4764fca89965f193fb27a3c,osType:Linux}},timeCreated:2025-11-17T04:53:43.2838588Z,vmId:b23d5c8c-70da-474d-a348-3223b667c494}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d4dakiip2foc73bcejp0,keypair:d4daje2p2foc73bcejl0,publicip:d4dakiip2foc73bcejp0-13587-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejp0"
            },
            {
              "key": "Name",
              "value": "d4dakiip2foc73bcejp0"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
          "uid": "d4dakiip2foc73bcejo0",
          "cspResourceName": "d4dakiip2foc73bcejo0",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejo0",
          "name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
          "subGroupId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
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
          "createdTime": "2025-11-17 04:54:33",
          "label": {
            "sourceMachineId": "ec2643f0-9388-3d97-f3a4-f387cd52696c",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2025-11-17 04:54:33",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejo0",
            "sys.cspResourceName": "d4dakiip2foc73bcejo0",
            "sys.id": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
            "sys.uid": "d4dakiip2foc73bcejo0"
          },
          "description": "a recommended virtual machine 01 for ec2643f0-9388-3d97-f3a4-f387cd52696c",
          "region": {
            "Region": "koreasouth",
            "Zone": ""
          },
          "publicIP": "52.231.224.191",
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
          "specId": "azure+koreasouth+standard_b2als_v2",
          "cspSpecName": "Standard_B2als_v2",
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0/subnets/d4dajbap2foc73bcejkg",
          "networkInterface": "d4dakiip2foc73bcejo0-33099-VNic",
          "securityGroupIds": [
            "mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4daje2p2foc73bcejl0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-17T04:55:16Z",
              "completedTime": "2025-11-17T04:55:19Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4dakiip2foc73bcejo0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4dakiip2foc73bcejo0-33099-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4dakiip2foc73bcejo0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCVOfrQgxYonLFyITWLcesa5oe5/61JY3WEAK8xFzmdg0NgDnM7ujyxXPKtQwpgaP+Yk7eWtLomfPkQhD77pjB5kOG1SnPXN/Q4sEcsx7ESh0lD123BN4F05WHM/HCHMQixf+3p8m32X3MDEMmkb0wtaEQ+AKfLFevCbkc8QTi5vb7oAwFuPWfMRerAHcFHkbFMbLYHdp3YANIzLHSyXfTtuIIk2qs9aFBQLUQ0JrxrxVn1c496i3jE8up4GoWVfupFypNmEfAWuRU951QvBCbH50lEKaHhin0GJG22UOjvM+4w1xp6Jx09vhiqqJR2pnbRDltP7k9nuwa3uZ9oSdHsvzyJxnHqJhWnVTkbLbTxbLB8dChDr98jeja3VOYvGB7ikAPTtxpuGrU/SiDHruUSseDSk0xJZ8+6utItex3vvac3YsglJqktHyPVhx5ErI8WekNImq008iJ75CZ5BV/3OggfLeaAa10jyuoU0IUAsVRUqXKdwXOiG8B66FWXB3Zch4Y086niqb+im7I04d2/Y3SsEnHASzw6r2vEV2nkJlQ1xgas8EijRb7LFxrIVnYVza2STSY3h+O0UJwl9f3EFXsOjZOtiEuvk2yX8nRRYnvceies5Jw+XVhwOfIA+2LGlBhLfeSz6pVNzGiTEfucRVm3aFlNOOhITbt7A1pQ4Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4dakiip2foc73bcejo0_OsDisk_1_23108f5f194142beace6b397d428798a,storageAccountType:Premium_LRS},name:d4dakiip2foc73bcejo0_OsDisk_1_23108f5f194142beace6b397d428798a,osType:Linux}},timeCreated:2025-11-17T04:53:46.2683236Z,vmId:91adf7d7-0973-4f01-b9ad-542b400abac8}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d4dakiip2foc73bcejo0,keypair:d4daje2p2foc73bcejl0,publicip:d4dakiip2foc73bcejo0-10794-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejo0"
            },
            {
              "key": "Name",
              "value": "d4dakiip2foc73bcejo0"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
          "uid": "d4dakiip2foc73bcejq0",
          "cspResourceName": "d4dakiip2foc73bcejq0",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejq0",
          "name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
          "subGroupId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
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
          "createdTime": "2025-11-17 04:54:32",
          "label": {
            "sourceMachineId": "ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2025-11-17 04:54:32",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejq0",
            "sys.cspResourceName": "d4dakiip2foc73bcejq0",
            "sys.id": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
            "sys.uid": "d4dakiip2foc73bcejq0"
          },
          "description": "a recommended virtual machine 03 for ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
          "region": {
            "Region": "koreasouth",
            "Zone": ""
          },
          "publicIP": "52.231.224.164",
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
          "specId": "azure+koreasouth+standard_b2as_v2",
          "cspSpecName": "Standard_B2as_v2",
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0/subnets/d4dajbap2foc73bcejkg",
          "networkInterface": "d4dakiip2foc73bcejq0-14462-VNic",
          "securityGroupIds": [
            "mig-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4daje2p2foc73bcejl0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-17T04:55:16Z",
              "completedTime": "2025-11-17T04:55:20Z",
              "elapsedTime": 4,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4dakiip2foc73bcejq0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4dakiip2foc73bcejq0-14462-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4dakiip2foc73bcejq0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCVOfrQgxYonLFyITWLcesa5oe5/61JY3WEAK8xFzmdg0NgDnM7ujyxXPKtQwpgaP+Yk7eWtLomfPkQhD77pjB5kOG1SnPXN/Q4sEcsx7ESh0lD123BN4F05WHM/HCHMQixf+3p8m32X3MDEMmkb0wtaEQ+AKfLFevCbkc8QTi5vb7oAwFuPWfMRerAHcFHkbFMbLYHdp3YANIzLHSyXfTtuIIk2qs9aFBQLUQ0JrxrxVn1c496i3jE8up4GoWVfupFypNmEfAWuRU951QvBCbH50lEKaHhin0GJG22UOjvM+4w1xp6Jx09vhiqqJR2pnbRDltP7k9nuwa3uZ9oSdHsvzyJxnHqJhWnVTkbLbTxbLB8dChDr98jeja3VOYvGB7ikAPTtxpuGrU/SiDHruUSseDSk0xJZ8+6utItex3vvac3YsglJqktHyPVhx5ErI8WekNImq008iJ75CZ5BV/3OggfLeaAa10jyuoU0IUAsVRUqXKdwXOiG8B66FWXB3Zch4Y086niqb+im7I04d2/Y3SsEnHASzw6r2vEV2nkJlQ1xgas8EijRb7LFxrIVnYVza2STSY3h+O0UJwl9f3EFXsOjZOtiEuvk2yX8nRRYnvceies5Jw+XVhwOfIA+2LGlBhLfeSz6pVNzGiTEfucRVm3aFlNOOhITbt7A1pQ4Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4dakiip2foc73bcejq0_OsDisk_1_dfdf78a08cc24954b9e1d1d4e24b5ce5,storageAccountType:Premium_LRS},name:d4dakiip2foc73bcejq0_OsDisk_1_dfdf78a08cc24954b9e1d1d4e24b5ce5,osType:Linux}},timeCreated:2025-11-17T04:53:45.7839512Z,vmId:dac8bf16-7235-45df-b652-4c723ba9fc6f}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d4dakiip2foc73bcejq0,keypair:d4daje2p2foc73bcejl0,publicip:d4dakiip2foc73bcejq0-49659-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejq0"
            },
            {
              "key": "Name",
              "value": "d4dakiip2foc73bcejq0"
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
            "vmId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
            "vmIp": "52.231.224.138",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4dakiip2foc73bcejp0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
            "vmIp": "52.231.224.191",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4dakiip2foc73bcejo0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
            "vmIp": "52.231.224.164",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4dakiip2foc73bcejq0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "d4dakiip2foc73bcejn0",
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
    "sys.uid": "d4dakiip2foc73bcejn0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
      "uid": "d4dakiip2foc73bcejp0",
      "cspResourceName": "d4dakiip2foc73bcejp0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejp0",
      "name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
      "subGroupId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
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
      "createdTime": "2025-11-17 04:54:31",
      "label": {
        "sourceMachineId": "ec21fd51-16bb-7e10-5e23-12ef283b2204",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2025-11-17 04:54:31",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejp0",
        "sys.cspResourceName": "d4dakiip2foc73bcejp0",
        "sys.id": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204",
        "sys.uid": "d4dakiip2foc73bcejp0"
      },
      "description": "a recommended virtual machine 02 for ec21fd51-16bb-7e10-5e23-12ef283b2204",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.224.138",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0/subnets/d4dajbap2foc73bcejkg",
      "networkInterface": "d4dakiip2foc73bcejp0-3040-VNic",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4daje2p2foc73bcejl0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:55:16Z",
          "completedTime": "2025-11-17T04:55:18Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4dakiip2foc73bcejp0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4dakiip2foc73bcejp0-3040-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4dakiip2foc73bcejp0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCVOfrQgxYonLFyITWLcesa5oe5/61JY3WEAK8xFzmdg0NgDnM7ujyxXPKtQwpgaP+Yk7eWtLomfPkQhD77pjB5kOG1SnPXN/Q4sEcsx7ESh0lD123BN4F05WHM/HCHMQixf+3p8m32X3MDEMmkb0wtaEQ+AKfLFevCbkc8QTi5vb7oAwFuPWfMRerAHcFHkbFMbLYHdp3YANIzLHSyXfTtuIIk2qs9aFBQLUQ0JrxrxVn1c496i3jE8up4GoWVfupFypNmEfAWuRU951QvBCbH50lEKaHhin0GJG22UOjvM+4w1xp6Jx09vhiqqJR2pnbRDltP7k9nuwa3uZ9oSdHsvzyJxnHqJhWnVTkbLbTxbLB8dChDr98jeja3VOYvGB7ikAPTtxpuGrU/SiDHruUSseDSk0xJZ8+6utItex3vvac3YsglJqktHyPVhx5ErI8WekNImq008iJ75CZ5BV/3OggfLeaAa10jyuoU0IUAsVRUqXKdwXOiG8B66FWXB3Zch4Y086niqb+im7I04d2/Y3SsEnHASzw6r2vEV2nkJlQ1xgas8EijRb7LFxrIVnYVza2STSY3h+O0UJwl9f3EFXsOjZOtiEuvk2yX8nRRYnvceies5Jw+XVhwOfIA+2LGlBhLfeSz6pVNzGiTEfucRVm3aFlNOOhITbt7A1pQ4Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4dakiip2foc73bcejp0_OsDisk_1_f83e6cd3d4764fca89965f193fb27a3c,storageAccountType:Premium_LRS},name:d4dakiip2foc73bcejp0_OsDisk_1_f83e6cd3d4764fca89965f193fb27a3c,osType:Linux}},timeCreated:2025-11-17T04:53:43.2838588Z,vmId:b23d5c8c-70da-474d-a348-3223b667c494}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4dakiip2foc73bcejp0,keypair:d4daje2p2foc73bcejl0,publicip:d4dakiip2foc73bcejp0-13587-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejp0"
        },
        {
          "key": "Name",
          "value": "d4dakiip2foc73bcejp0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
      "uid": "d4dakiip2foc73bcejo0",
      "cspResourceName": "d4dakiip2foc73bcejo0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejo0",
      "name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
      "subGroupId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
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
      "createdTime": "2025-11-17 04:54:33",
      "label": {
        "sourceMachineId": "ec2643f0-9388-3d97-f3a4-f387cd52696c",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2025-11-17 04:54:33",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejo0",
        "sys.cspResourceName": "d4dakiip2foc73bcejo0",
        "sys.id": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c",
        "sys.uid": "d4dakiip2foc73bcejo0"
      },
      "description": "a recommended virtual machine 01 for ec2643f0-9388-3d97-f3a4-f387cd52696c",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.224.191",
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
      "specId": "azure+koreasouth+standard_b2als_v2",
      "cspSpecName": "Standard_B2als_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0/subnets/d4dajbap2foc73bcejkg",
      "networkInterface": "d4dakiip2foc73bcejo0-33099-VNic",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4daje2p2foc73bcejl0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:55:16Z",
          "completedTime": "2025-11-17T04:55:19Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4dakiip2foc73bcejo0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4dakiip2foc73bcejo0-33099-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4dakiip2foc73bcejo0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCVOfrQgxYonLFyITWLcesa5oe5/61JY3WEAK8xFzmdg0NgDnM7ujyxXPKtQwpgaP+Yk7eWtLomfPkQhD77pjB5kOG1SnPXN/Q4sEcsx7ESh0lD123BN4F05WHM/HCHMQixf+3p8m32X3MDEMmkb0wtaEQ+AKfLFevCbkc8QTi5vb7oAwFuPWfMRerAHcFHkbFMbLYHdp3YANIzLHSyXfTtuIIk2qs9aFBQLUQ0JrxrxVn1c496i3jE8up4GoWVfupFypNmEfAWuRU951QvBCbH50lEKaHhin0GJG22UOjvM+4w1xp6Jx09vhiqqJR2pnbRDltP7k9nuwa3uZ9oSdHsvzyJxnHqJhWnVTkbLbTxbLB8dChDr98jeja3VOYvGB7ikAPTtxpuGrU/SiDHruUSseDSk0xJZ8+6utItex3vvac3YsglJqktHyPVhx5ErI8WekNImq008iJ75CZ5BV/3OggfLeaAa10jyuoU0IUAsVRUqXKdwXOiG8B66FWXB3Zch4Y086niqb+im7I04d2/Y3SsEnHASzw6r2vEV2nkJlQ1xgas8EijRb7LFxrIVnYVza2STSY3h+O0UJwl9f3EFXsOjZOtiEuvk2yX8nRRYnvceies5Jw+XVhwOfIA+2LGlBhLfeSz6pVNzGiTEfucRVm3aFlNOOhITbt7A1pQ4Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4dakiip2foc73bcejo0_OsDisk_1_23108f5f194142beace6b397d428798a,storageAccountType:Premium_LRS},name:d4dakiip2foc73bcejo0_OsDisk_1_23108f5f194142beace6b397d428798a,osType:Linux}},timeCreated:2025-11-17T04:53:46.2683236Z,vmId:91adf7d7-0973-4f01-b9ad-542b400abac8}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4dakiip2foc73bcejo0,keypair:d4daje2p2foc73bcejl0,publicip:d4dakiip2foc73bcejo0-10794-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejo0"
        },
        {
          "key": "Name",
          "value": "d4dakiip2foc73bcejo0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
      "uid": "d4dakiip2foc73bcejq0",
      "cspResourceName": "d4dakiip2foc73bcejq0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejq0",
      "name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
      "subGroupId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
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
      "createdTime": "2025-11-17 04:54:32",
      "label": {
        "sourceMachineId": "ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2025-11-17 04:54:32",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejq0",
        "sys.cspResourceName": "d4dakiip2foc73bcejq0",
        "sys.id": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
        "sys.uid": "d4dakiip2foc73bcejq0"
      },
      "description": "a recommended virtual machine 03 for ec2876a6-3c84-7e62-aaf9-c3203f12e0b8",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.224.164",
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
      "specId": "azure+koreasouth+standard_b2as_v2",
      "cspSpecName": "Standard_B2as_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4dajbap2foc73bcejk0/subnets/d4dajbap2foc73bcejkg",
      "networkInterface": "d4dakiip2foc73bcejq0-14462-VNic",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4daje2p2foc73bcejl0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-17T04:55:16Z",
          "completedTime": "2025-11-17T04:55:20Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4dakiip2foc73bcejq0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4dakiip2foc73bcejq0-14462-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4dakiip2foc73bcejq0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCVOfrQgxYonLFyITWLcesa5oe5/61JY3WEAK8xFzmdg0NgDnM7ujyxXPKtQwpgaP+Yk7eWtLomfPkQhD77pjB5kOG1SnPXN/Q4sEcsx7ESh0lD123BN4F05WHM/HCHMQixf+3p8m32X3MDEMmkb0wtaEQ+AKfLFevCbkc8QTi5vb7oAwFuPWfMRerAHcFHkbFMbLYHdp3YANIzLHSyXfTtuIIk2qs9aFBQLUQ0JrxrxVn1c496i3jE8up4GoWVfupFypNmEfAWuRU951QvBCbH50lEKaHhin0GJG22UOjvM+4w1xp6Jx09vhiqqJR2pnbRDltP7k9nuwa3uZ9oSdHsvzyJxnHqJhWnVTkbLbTxbLB8dChDr98jeja3VOYvGB7ikAPTtxpuGrU/SiDHruUSseDSk0xJZ8+6utItex3vvac3YsglJqktHyPVhx5ErI8WekNImq008iJ75CZ5BV/3OggfLeaAa10jyuoU0IUAsVRUqXKdwXOiG8B66FWXB3Zch4Y086niqb+im7I04d2/Y3SsEnHASzw6r2vEV2nkJlQ1xgas8EijRb7LFxrIVnYVza2STSY3h+O0UJwl9f3EFXsOjZOtiEuvk2yX8nRRYnvceies5Jw+XVhwOfIA+2LGlBhLfeSz6pVNzGiTEfucRVm3aFlNOOhITbt7A1pQ4Q==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4dakiip2foc73bcejq0_OsDisk_1_dfdf78a08cc24954b9e1d1d4e24b5ce5,storageAccountType:Premium_LRS},name:d4dakiip2foc73bcejq0_OsDisk_1_dfdf78a08cc24954b9e1d1d4e24b5ce5,osType:Linux}},timeCreated:2025-11-17T04:53:45.7839512Z,vmId:dac8bf16-7235-45df-b652-4c723ba9fc6f}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4dakiip2foc73bcejq0,keypair:d4daje2p2foc73bcejl0,publicip:d4dakiip2foc73bcejq0-49659-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4dakiip2foc73bcejq0"
        },
        {
          "key": "Name",
          "value": "d4dakiip2foc73bcejq0"
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
        "vmId": "migrated-ec21fd51-16bb-7e10-5e23-12ef283b2204-1",
        "vmIp": "52.231.224.138",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4dakiip2foc73bcejp0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2643f0-9388-3d97-f3a4-f387cd52696c-1",
        "vmIp": "52.231.224.191",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4dakiip2foc73bcejo0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2876a6-3c84-7e62-aaf9-c3203f12e0b8-1",
        "vmIp": "52.231.224.164",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4dakiip2foc73bcejq0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "output": "Linux d4dakiip2foc73bcejp0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.224.138",
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
      "output": "Linux d4dakiip2foc73bcejo0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.224.191",
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
      "output": "Linux d4dakiip2foc73bcejq0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.224.164",
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

