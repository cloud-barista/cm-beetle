# CM-Beetle test results for GCP

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with GCP cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.4 (b04128d)
- cm-model: v0.0.14
- CB-Tumblebug: v0.11.19
- CB-Spider: v0.11.16
- CB-MapUI: v0.11.19
- Target CSP: GCP
- Target Region: asia-northeast3
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: November 10, 2025
- Test Time: 20:39:17 KST
- Test Execution: 2025-11-10 20:39:17 KST

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for GCP

### Test Results Summary

| Test | Endpoint | Status | Duration | Details |
|------|----------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/mci` | ✅ **PASS** | 1.268s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 7m33.693s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 881ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 87ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 410ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 8m19.233s | Pass |

**Overall Result**: 7/7 tests passed ✅

**Total Duration**: 16m42.610524582s

*Test executed on November 10, 2025 at 20:39:17 KST (2025-11-10 20:39:17 KST) using CM-Beetle automated test CLI*

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
    "csp": "gcp",
    "region": "asia-northeast3"
  },
  "OnpremiseInfraModel": {
    "network": {
      "ipv4Networks": {
        "defaultGateways": [
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c"
          },
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939"
          },
          {
            "ip": "10.0.1.1",
            "interfaceName": "ens5",
            "machineId": "ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4"
          }
        ]
      },
      "ipv6Networks": {}
    },
    "servers": [
      {
        "hostname": "ip-10-0-1-25",
        "machineId": "ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
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
            "macAddress": "02:1d:0c:42:f1:51",
            "ipv4CidrBlocks": [
              "10.0.1.25/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::1d:cff:fe42:f151/64"
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
            "destination": "fe80::1d:cff:fe42:f151/128",
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
        "hostname": "ip-10-0-1-220",
        "machineId": "ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
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
            "macAddress": "02:4c:5f:01:55:b3",
            "ipv4CidrBlocks": [
              "10.0.1.220/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::4c:5fff:fe01:55b3/64"
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
            "destination": "fe80::4c:5fff:fe01:55b3/128",
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
        "hostname": "ip-10-0-1-131",
        "machineId": "ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
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
            "macAddress": "02:c0:4b:06:72:c7",
            "ipv4CidrBlocks": [
              "10.0.1.131/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::c0:4bff:fe06:72c7/64"
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
            "destination": "fe80::c0:4bff:fe06:72c7/128",
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
    "csp": "gcp",
    "region": "asia-northeast3"
  },
  "targetVmInfra": {
    "name": "mmci01",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "a recommended multi-cloud infrastructure",
    "subGroups": [
      {
        "name": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c"
        },
        "description": "a recommended virtual machine 01 for ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
        "connectionName": "gcp-asia-northeast3",
        "specId": "gcp+asia-northeast3+e2-small",
        "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
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
        "name": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939"
        },
        "description": "a recommended virtual machine 02 for ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
        "connectionName": "gcp-asia-northeast3",
        "specId": "gcp+asia-northeast3+e2-standard-4",
        "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
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
        "name": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4"
        },
        "description": "a recommended virtual machine 03 for ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
        "connectionName": "gcp-asia-northeast3",
        "specId": "gcp+asia-northeast3+e2-standard-2",
        "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
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
    "connectionName": "gcp-asia-northeast3",
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
    "connectionName": "gcp-asia-northeast3",
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
      "id": "gcp+asia-northeast3+e2-small",
      "uid": "d3vjcqeqjs728pq3b1q0",
      "cspSpecName": "e2-small",
      "name": "gcp+asia-northeast3+e2-small",
      "namespace": "system",
      "connectionName": "gcp-asia-northeast3",
      "providerName": "gcp",
      "regionName": "asia-northeast3",
      "regionLatitude": 37.2,
      "regionLongitude": 127,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 1.953125,
      "diskSizeGB": -1,
      "costPerHour": 0.02149143,
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
          "key": "CreationTimestamp",
          "value": "1969-12-31T16:00:00.000-08:00"
        },
        {
          "key": "Description",
          "value": "Efficient Instance, 2 vCPU (1/4 shared physical core) and 2 GB RAM"
        },
        {
          "key": "GuestCpus",
          "value": "2"
        },
        {
          "key": "Id",
          "value": "334003"
        },
        {
          "key": "ImageSpaceGb",
          "value": "0"
        },
        {
          "key": "IsSharedCpu",
          "value": "true"
        },
        {
          "key": "Kind",
          "value": "compute#machineType"
        },
        {
          "key": "MaximumPersistentDisks",
          "value": "16"
        },
        {
          "key": "MaximumPersistentDisksSizeGb",
          "value": "3072"
        },
        {
          "key": "MemoryMb",
          "value": "2048"
        },
        {
          "key": "Name",
          "value": "e2-small"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-small"
        },
        {
          "key": "Zone",
          "value": "asia-northeast3-a"
        }
      ]
    },
    {
      "id": "gcp+asia-northeast3+e2-standard-4",
      "uid": "d3vjcqeqjs728pq3b1s0",
      "cspSpecName": "e2-standard-4",
      "name": "gcp+asia-northeast3+e2-standard-4",
      "namespace": "system",
      "connectionName": "gcp-asia-northeast3",
      "providerName": "gcp",
      "regionName": "asia-northeast3",
      "regionLatitude": 37.2,
      "regionLongitude": 127,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 15.625,
      "diskSizeGB": -1,
      "costPerHour": 0.17193145,
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
          "key": "CreationTimestamp",
          "value": "1969-12-31T16:00:00.000-08:00"
        },
        {
          "key": "Description",
          "value": "Efficient Instance, 4 vCPUs, 16 GB RAM"
        },
        {
          "key": "GuestCpus",
          "value": "4"
        },
        {
          "key": "Id",
          "value": "335004"
        },
        {
          "key": "ImageSpaceGb",
          "value": "0"
        },
        {
          "key": "IsSharedCpu",
          "value": "false"
        },
        {
          "key": "Kind",
          "value": "compute#machineType"
        },
        {
          "key": "MaximumPersistentDisks",
          "value": "128"
        },
        {
          "key": "MaximumPersistentDisksSizeGb",
          "value": "263168"
        },
        {
          "key": "MemoryMb",
          "value": "16384"
        },
        {
          "key": "Name",
          "value": "e2-standard-4"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Zone",
          "value": "asia-northeast3-a"
        }
      ]
    },
    {
      "id": "gcp+asia-northeast3+e2-standard-2",
      "uid": "d3vjcqeqjs728pq3b1r0",
      "cspSpecName": "e2-standard-2",
      "name": "gcp+asia-northeast3+e2-standard-2",
      "namespace": "system",
      "connectionName": "gcp-asia-northeast3",
      "providerName": "gcp",
      "regionName": "asia-northeast3",
      "regionLatitude": 37.2,
      "regionLongitude": 127,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 7.8125,
      "diskSizeGB": -1,
      "costPerHour": 0.08596572,
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
          "key": "CreationTimestamp",
          "value": "1969-12-31T16:00:00.000-08:00"
        },
        {
          "key": "Description",
          "value": "Efficient Instance, 2 vCPUs, 8 GB RAM"
        },
        {
          "key": "GuestCpus",
          "value": "2"
        },
        {
          "key": "Id",
          "value": "335002"
        },
        {
          "key": "ImageSpaceGb",
          "value": "0"
        },
        {
          "key": "IsSharedCpu",
          "value": "false"
        },
        {
          "key": "Kind",
          "value": "compute#machineType"
        },
        {
          "key": "MaximumPersistentDisks",
          "value": "128"
        },
        {
          "key": "MaximumPersistentDisksSizeGb",
          "value": "263168"
        },
        {
          "key": "MemoryMb",
          "value": "8192"
        },
        {
          "key": "Name",
          "value": "e2-standard-2"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-2"
        },
        {
          "key": "Zone",
          "value": "asia-northeast3-a"
        }
      ]
    }
  ],
  "targetVmOsImageList": [
    {
      "namespace": "system",
      "providerName": "gcp",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "regionList": [
        "common"
      ],
      "id": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "uid": "d3vjd8eqjs728pq56b5g",
      "name": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "connectionName": "gcp-africa-south1",
      "infraType": "k8s|kubernetes|container",
      "fetchedTime": "2025.10.27 09:08:49 Mon",
      "isKubernetesImage": true,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2025-10-23",
      "osDiskType": "NA",
      "osDiskSizeGB": 10,
      "imageStatus": "Available",
      "details": [
        {
          "key": "Architecture",
          "value": "X86_64"
        },
        {
          "key": "ArchiveSizeBytes",
          "value": "3928290624"
        },
        {
          "key": "CreationTimestamp",
          "value": "2025-10-23T09:59:33.029-07:00"
        },
        {
          "key": "Description",
          "value": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2025-10-23"
        },
        {
          "key": "DiskSizeGb",
          "value": "10"
        },
        {
          "key": "EnableConfidentialCompute",
          "value": "false"
        },
        {
          "key": "Family",
          "value": "ubuntu-2204-lts"
        },
        {
          "key": "GuestOsFeatures",
          "value": "{type:VIRTIO_SCSI_MULTIQUEUE}; {type:SEV_CAPABLE}; {type:SEV_SNP_CAPABLE}; {type:SEV_LIVE_MIGRATABLE}; {type:SEV_LIVE_MIGRATABLE_V2}; {type:IDPF}; {type:TDX_CAPABLE}; {type:UEFI_COMPATIBLE}; {type:GVNIC}"
        },
        {
          "key": "Id",
          "value": "6937401861558755355"
        },
        {
          "key": "Kind",
          "value": "compute#image"
        },
        {
          "key": "LabelFingerprint",
          "value": "iNBmVNCFF9w="
        },
        {
          "key": "Labels",
          "value": "{public-image:true}"
        },
        {
          "key": "LicenseCodes",
          "value": "5511465778777431107"
        },
        {
          "key": "Licenses",
          "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts"
        },
        {
          "key": "Name",
          "value": "ubuntu-2204-jammy-v20251023"
        },
        {
          "key": "RawDisk",
          "value": "{containerType:TAR}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "false"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023"
        },
        {
          "key": "SourceType",
          "value": "RAW"
        },
        {
          "key": "Status",
          "value": "READY"
        },
        {
          "key": "StorageLocations",
          "value": "asia; eu; us"
        }
      ],
      "description": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2025-10-23"
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "mig-sg-01",
      "connectionName": "gcp-asia-northeast3",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
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
      "connectionName": "gcp-asia-northeast3",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
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
      "connectionName": "gcp-asia-northeast3",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
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
  "uid": "d48t0k8t49avt161vnmg",
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
    "sys.uid": "d48t0k8t49avt161vnmg"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
      "uid": "d48t0k8t49avt161vnog",
      "cspResourceName": "d48t0k8t49avt161vnog",
      "cspResourceId": "d48t0k8t49avt161vnog",
      "name": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
      "subGroupId": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-11-10 11:46:08",
      "label": {
        "sourceMachineId": "ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939"
      },
      "description": "a recommended virtual machine 02 for ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.50.23.84",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.2",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "gcp-asia-northeast3",
      "connectionConfig": {
        "configName": "gcp-asia-northeast3",
        "providerName": "gcp",
        "driverName": "gcp-driver-v1.0.so",
        "credentialName": "gcp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "gcp-asia-northeast3",
        "regionZoneInfo": {
          "assignedRegion": "asia-northeast3",
          "assignedZone": "asia-northeast3-a"
        },
        "regionDetail": {
          "regionId": "asia-northeast3",
          "regionName": "asia-northeast3",
          "description": "Seoul South Korea",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "zones": [
            "asia-northeast3-a",
            "asia-northeast3-b",
            "asia-northeast3-c"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "gcp+asia-northeast3+e2-standard-4",
      "cspSpecName": "e2-standard-4",
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "d48strot49avt161vnjg",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "d48strot49avt161vnk0",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d48su38t49avt161vnkg",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-10T11:46:21Z",
          "completedTime": "2025-11-10T11:46:20Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d48t0k8t49avt161vnog 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2025-11-10T03:45:28.644-08:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d48t0k8t49avt161vnog,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "-mtIceZL_8Y="
        },
        {
          "key": "Id",
          "value": "2038569319180037815"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "GcPBmaOCA1o="
        },
        {
          "key": "Labels",
          "value": "{keypair:d48su38t49avt161vnkg}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2025-11-10T03:45:39.688-08:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:fQXYV2AGZR8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDOYl9nmr2wlGkRH8IJPjNNZYy+1QnLxudDBQe65h/csl8teCvQU9YpDd24XUrqk2Vj/rlE5MlbQrXX5RByUjolasCVGy7/QLXH3BwmqNuoHvYyFTpnQlFBBDg7Dq+U3T1FdGRdSY/ZwRuz6CQC1Vg+kIRweajbtxOyMIKGHRrHAI8Xw1GafDod/Wz83sLbm+F3ZB0UnVnFeu8xNdaqfdupoI3Rqgw0sNgnBahvOXhHGsnX8WgxrhDeRqsBO5SYLcqi614dNbXS3LwXG6XH6GBYB/M2TJLNTvd/mH7cEUxTsTttiFreF8WZE64oGHGPQgAoXSIzaS0/F+W9k6rgFCqkcQwaecqYvcxQp8u0WOChBcjECt1cHtHvQmPdgdI5URK1DxlJoTEionAslQBP6McxRE3hXyDH5RTM9pdpaFxRXY5AHUIjpKCuRZeRY4Xa4mnB0X9IMk7U1j6S05A4plKeCJtnJsuw6hnnDNVfQ24J0OsywdBmDrC5GRN7OPTZtqoelCqMHJ6h9ii7WB5qUz7ITLUQ2KSl4cmXqJ0dfOCi/vw5/vmH7iLgXfZwbzfLCAXMSGp4KopRxTvoorYxGLLcKhs1CrWHEsusF1QmXRIICrSKkq7P6yOn3Iq4OmwEZ1L9Z9RXABRZ+RPaMVwymnCtDJ54UoJYkYlVvxu0uOzwSQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "d48t0k8t49avt161vnog"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.50.23.84,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:f9JJzLXhNnA=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d48strot49avt161vnjg,networkIP:10.0.1.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d48strot49avt161vnk0}"
        },
        {
          "key": "ResourceStatus",
          "value": "{}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d48t0k8t49avt161vnog"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:TnO5353i2ig=,items:[d48supgt49avt161vnlg]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
      "uid": "d48t0k8t49avt161vnpg",
      "cspResourceName": "d48t0k8t49avt161vnpg",
      "cspResourceId": "d48t0k8t49avt161vnpg",
      "name": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
      "subGroupId": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-11-10 11:46:15",
      "label": {
        "sourceMachineId": "ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4"
      },
      "description": "a recommended virtual machine 03 for ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.64.193.4",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "gcp-asia-northeast3",
      "connectionConfig": {
        "configName": "gcp-asia-northeast3",
        "providerName": "gcp",
        "driverName": "gcp-driver-v1.0.so",
        "credentialName": "gcp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "gcp-asia-northeast3",
        "regionZoneInfo": {
          "assignedRegion": "asia-northeast3",
          "assignedZone": "asia-northeast3-a"
        },
        "regionDetail": {
          "regionId": "asia-northeast3",
          "regionName": "asia-northeast3",
          "description": "Seoul South Korea",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "zones": [
            "asia-northeast3-a",
            "asia-northeast3-b",
            "asia-northeast3-c"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "gcp+asia-northeast3+e2-standard-2",
      "cspSpecName": "e2-standard-2",
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "d48strot49avt161vnjg",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "d48strot49avt161vnk0",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d48su38t49avt161vnkg",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-10T11:46:21Z",
          "completedTime": "2025-11-10T11:46:24Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d48t0k8t49avt161vnpg 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2025-11-10T03:45:30.352-08:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d48t0k8t49avt161vnpg,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "FeW8WkYqYmM="
        },
        {
          "key": "Id",
          "value": "6036202776671579829"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "GcPBmaOCA1o="
        },
        {
          "key": "Labels",
          "value": "{keypair:d48su38t49avt161vnkg}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2025-11-10T03:45:45.274-08:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-2"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:fQXYV2AGZR8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDOYl9nmr2wlGkRH8IJPjNNZYy+1QnLxudDBQe65h/csl8teCvQU9YpDd24XUrqk2Vj/rlE5MlbQrXX5RByUjolasCVGy7/QLXH3BwmqNuoHvYyFTpnQlFBBDg7Dq+U3T1FdGRdSY/ZwRuz6CQC1Vg+kIRweajbtxOyMIKGHRrHAI8Xw1GafDod/Wz83sLbm+F3ZB0UnVnFeu8xNdaqfdupoI3Rqgw0sNgnBahvOXhHGsnX8WgxrhDeRqsBO5SYLcqi614dNbXS3LwXG6XH6GBYB/M2TJLNTvd/mH7cEUxTsTttiFreF8WZE64oGHGPQgAoXSIzaS0/F+W9k6rgFCqkcQwaecqYvcxQp8u0WOChBcjECt1cHtHvQmPdgdI5URK1DxlJoTEionAslQBP6McxRE3hXyDH5RTM9pdpaFxRXY5AHUIjpKCuRZeRY4Xa4mnB0X9IMk7U1j6S05A4plKeCJtnJsuw6hnnDNVfQ24J0OsywdBmDrC5GRN7OPTZtqoelCqMHJ6h9ii7WB5qUz7ITLUQ2KSl4cmXqJ0dfOCi/vw5/vmH7iLgXfZwbzfLCAXMSGp4KopRxTvoorYxGLLcKhs1CrWHEsusF1QmXRIICrSKkq7P6yOn3Iq4OmwEZ1L9Z9RXABRZ+RPaMVwymnCtDJ54UoJYkYlVvxu0uOzwSQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "d48t0k8t49avt161vnpg"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.193.4,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:kIYgYOib8Xo=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d48strot49avt161vnjg,networkIP:10.0.1.4,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d48strot49avt161vnk0}"
        },
        {
          "key": "ResourceStatus",
          "value": "{}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d48t0k8t49avt161vnpg"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:5KCAhqva_iI=,items:[d48svkgt49avt161vnm0]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
      "uid": "d48t0k8t49avt161vnng",
      "cspResourceName": "d48t0k8t49avt161vnng",
      "cspResourceId": "d48t0k8t49avt161vnng",
      "name": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
      "subGroupId": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-11-10 11:46:09",
      "label": {
        "sourceMachineId": "ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c"
      },
      "description": "a recommended virtual machine 01 for ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.50.33.86",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.3",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "gcp-asia-northeast3",
      "connectionConfig": {
        "configName": "gcp-asia-northeast3",
        "providerName": "gcp",
        "driverName": "gcp-driver-v1.0.so",
        "credentialName": "gcp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "gcp-asia-northeast3",
        "regionZoneInfo": {
          "assignedRegion": "asia-northeast3",
          "assignedZone": "asia-northeast3-a"
        },
        "regionDetail": {
          "regionId": "asia-northeast3",
          "regionName": "asia-northeast3",
          "description": "Seoul South Korea",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "zones": [
            "asia-northeast3-a",
            "asia-northeast3-b",
            "asia-northeast3-c"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "gcp+asia-northeast3+e2-small",
      "cspSpecName": "e2-small",
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "d48strot49avt161vnjg",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "d48strot49avt161vnk0",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d48su38t49avt161vnkg",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-10T11:46:21Z",
          "completedTime": "2025-11-10T11:46:21Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d48t0k8t49avt161vnng 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2025-11-10T03:45:29.652-08:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d48t0k8t49avt161vnng,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "sSXEs3PF3tk="
        },
        {
          "key": "Id",
          "value": "7899419285426826934"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "GcPBmaOCA1o="
        },
        {
          "key": "Labels",
          "value": "{keypair:d48su38t49avt161vnkg}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2025-11-10T03:45:40.781-08:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-small"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:fQXYV2AGZR8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDOYl9nmr2wlGkRH8IJPjNNZYy+1QnLxudDBQe65h/csl8teCvQU9YpDd24XUrqk2Vj/rlE5MlbQrXX5RByUjolasCVGy7/QLXH3BwmqNuoHvYyFTpnQlFBBDg7Dq+U3T1FdGRdSY/ZwRuz6CQC1Vg+kIRweajbtxOyMIKGHRrHAI8Xw1GafDod/Wz83sLbm+F3ZB0UnVnFeu8xNdaqfdupoI3Rqgw0sNgnBahvOXhHGsnX8WgxrhDeRqsBO5SYLcqi614dNbXS3LwXG6XH6GBYB/M2TJLNTvd/mH7cEUxTsTttiFreF8WZE64oGHGPQgAoXSIzaS0/F+W9k6rgFCqkcQwaecqYvcxQp8u0WOChBcjECt1cHtHvQmPdgdI5URK1DxlJoTEionAslQBP6McxRE3hXyDH5RTM9pdpaFxRXY5AHUIjpKCuRZeRY4Xa4mnB0X9IMk7U1j6S05A4plKeCJtnJsuw6hnnDNVfQ24J0OsywdBmDrC5GRN7OPTZtqoelCqMHJ6h9ii7WB5qUz7ITLUQ2KSl4cmXqJ0dfOCi/vw5/vmH7iLgXfZwbzfLCAXMSGp4KopRxTvoorYxGLLcKhs1CrWHEsusF1QmXRIICrSKkq7P6yOn3Iq4OmwEZ1L9Z9RXABRZ+RPaMVwymnCtDJ54UoJYkYlVvxu0uOzwSQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "d48t0k8t49avt161vnng"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.50.33.86,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:j1JWJtpbWkg=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d48strot49avt161vnjg,networkIP:10.0.1.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d48strot49avt161vnk0}"
        },
        {
          "key": "ResourceStatus",
          "value": "{}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d48t0k8t49avt161vnng"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:Z0ZQ4oN1kFM=,items:[d48su3gt49avt161vnl0]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
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
        "vmId": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
        "vmIp": "34.50.23.84",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d48t0k8t49avt161vnog 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
        "vmIp": "34.50.33.86",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d48t0k8t49avt161vnng 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
        "vmIp": "34.64.193.4",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d48t0k8t49avt161vnpg 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "d48t0k8t49avt161vnmg",
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
        "sys.uid": "d48t0k8t49avt161vnmg"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "a recommended multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
          "uid": "d48t0k8t49avt161vnng",
          "cspResourceName": "d48t0k8t49avt161vnng",
          "cspResourceId": "d48t0k8t49avt161vnng",
          "name": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
          "subGroupId": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-11-10 11:46:09",
          "label": {
            "sourceMachineId": "ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2025-11-10 11:46:09",
            "sys.cspResourceId": "d48t0k8t49avt161vnng",
            "sys.cspResourceName": "d48t0k8t49avt161vnng",
            "sys.id": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
            "sys.uid": "d48t0k8t49avt161vnng"
          },
          "description": "a recommended virtual machine 01 for ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
          "region": {
            "Region": "asia-northeast3",
            "Zone": "asia-northeast3-a"
          },
          "publicIP": "34.50.33.86",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.3",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": "50",
          "rootDiskName": "",
          "connectionName": "gcp-asia-northeast3",
          "connectionConfig": {
            "configName": "gcp-asia-northeast3",
            "providerName": "gcp",
            "driverName": "gcp-driver-v1.0.so",
            "credentialName": "gcp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "gcp-asia-northeast3",
            "regionZoneInfo": {
              "assignedRegion": "asia-northeast3",
              "assignedZone": "asia-northeast3-a"
            },
            "regionDetail": {
              "regionId": "asia-northeast3",
              "regionName": "asia-northeast3",
              "description": "Seoul South Korea",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.2,
                "longitude": 127
              },
              "zones": [
                "asia-northeast3-a",
                "asia-northeast3-b",
                "asia-northeast3-c"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "gcp+asia-northeast3+e2-small",
          "cspSpecName": "e2-small",
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "d48strot49avt161vnjg",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "d48strot49avt161vnk0",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d48su38t49avt161vnkg",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-10T11:46:21Z",
              "completedTime": "2025-11-10T11:46:21Z",
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d48t0k8t49avt161vnng 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "CanIpForward",
              "value": "false"
            },
            {
              "key": "CpuPlatform",
              "value": "Intel Broadwell"
            },
            {
              "key": "CreationTimestamp",
              "value": "2025-11-10T03:45:29.652-08:00"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "Description",
              "value": "compute sample instance"
            },
            {
              "key": "Disks",
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d48t0k8t49avt161vnng,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "sSXEs3PF3tk="
            },
            {
              "key": "Id",
              "value": "7899419285426826934"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "GcPBmaOCA1o="
            },
            {
              "key": "Labels",
              "value": "{keypair:d48su38t49avt161vnkg}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2025-11-10T03:45:40.781-08:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-small"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:fQXYV2AGZR8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDOYl9nmr2wlGkRH8IJPjNNZYy+1QnLxudDBQe65h/csl8teCvQU9YpDd24XUrqk2Vj/rlE5MlbQrXX5RByUjolasCVGy7/QLXH3BwmqNuoHvYyFTpnQlFBBDg7Dq+U3T1FdGRdSY/ZwRuz6CQC1Vg+kIRweajbtxOyMIKGHRrHAI8Xw1GafDod/Wz83sLbm+F3ZB0UnVnFeu8xNdaqfdupoI3Rqgw0sNgnBahvOXhHGsnX8WgxrhDeRqsBO5SYLcqi614dNbXS3LwXG6XH6GBYB/M2TJLNTvd/mH7cEUxTsTttiFreF8WZE64oGHGPQgAoXSIzaS0/F+W9k6rgFCqkcQwaecqYvcxQp8u0WOChBcjECt1cHtHvQmPdgdI5URK1DxlJoTEionAslQBP6McxRE3hXyDH5RTM9pdpaFxRXY5AHUIjpKCuRZeRY4Xa4mnB0X9IMk7U1j6S05A4plKeCJtnJsuw6hnnDNVfQ24J0OsywdBmDrC5GRN7OPTZtqoelCqMHJ6h9ii7WB5qUz7ITLUQ2KSl4cmXqJ0dfOCi/vw5/vmH7iLgXfZwbzfLCAXMSGp4KopRxTvoorYxGLLcKhs1CrWHEsusF1QmXRIICrSKkq7P6yOn3Iq4OmwEZ1L9Z9RXABRZ+RPaMVwymnCtDJ54UoJYkYlVvxu0uOzwSQ== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "d48t0k8t49avt161vnng"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.50.33.86,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:j1JWJtpbWkg=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d48strot49avt161vnjg,networkIP:10.0.1.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d48strot49avt161vnk0}"
            },
            {
              "key": "ResourceStatus",
              "value": "{}"
            },
            {
              "key": "SatisfiesPzi",
              "value": "true"
            },
            {
              "key": "SatisfiesPzs",
              "value": "false"
            },
            {
              "key": "Scheduling",
              "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d48t0k8t49avt161vnng"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
            },
            {
              "key": "ShieldedInstanceConfig",
              "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
            },
            {
              "key": "ShieldedInstanceIntegrityPolicy",
              "value": "{updateAutoLearnPolicy:true}"
            },
            {
              "key": "StartRestricted",
              "value": "false"
            },
            {
              "key": "Status",
              "value": "RUNNING"
            },
            {
              "key": "Tags",
              "value": "{fingerprint:Z0ZQ4oN1kFM=,items:[d48su3gt49avt161vnl0]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
          "uid": "d48t0k8t49avt161vnpg",
          "cspResourceName": "d48t0k8t49avt161vnpg",
          "cspResourceId": "d48t0k8t49avt161vnpg",
          "name": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
          "subGroupId": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-11-10 11:46:15",
          "label": {
            "sourceMachineId": "ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2025-11-10 11:46:15",
            "sys.cspResourceId": "d48t0k8t49avt161vnpg",
            "sys.cspResourceName": "d48t0k8t49avt161vnpg",
            "sys.id": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
            "sys.uid": "d48t0k8t49avt161vnpg"
          },
          "description": "a recommended virtual machine 03 for ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
          "region": {
            "Region": "asia-northeast3",
            "Zone": "asia-northeast3-a"
          },
          "publicIP": "34.64.193.4",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": "50",
          "rootDiskName": "",
          "connectionName": "gcp-asia-northeast3",
          "connectionConfig": {
            "configName": "gcp-asia-northeast3",
            "providerName": "gcp",
            "driverName": "gcp-driver-v1.0.so",
            "credentialName": "gcp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "gcp-asia-northeast3",
            "regionZoneInfo": {
              "assignedRegion": "asia-northeast3",
              "assignedZone": "asia-northeast3-a"
            },
            "regionDetail": {
              "regionId": "asia-northeast3",
              "regionName": "asia-northeast3",
              "description": "Seoul South Korea",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.2,
                "longitude": 127
              },
              "zones": [
                "asia-northeast3-a",
                "asia-northeast3-b",
                "asia-northeast3-c"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "gcp+asia-northeast3+e2-standard-2",
          "cspSpecName": "e2-standard-2",
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "d48strot49avt161vnjg",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "d48strot49avt161vnk0",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "mig-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d48su38t49avt161vnkg",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-10T11:46:21Z",
              "completedTime": "2025-11-10T11:46:24Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d48t0k8t49avt161vnpg 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "CanIpForward",
              "value": "false"
            },
            {
              "key": "CpuPlatform",
              "value": "Intel Broadwell"
            },
            {
              "key": "CreationTimestamp",
              "value": "2025-11-10T03:45:30.352-08:00"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "Description",
              "value": "compute sample instance"
            },
            {
              "key": "Disks",
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d48t0k8t49avt161vnpg,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "FeW8WkYqYmM="
            },
            {
              "key": "Id",
              "value": "6036202776671579829"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "GcPBmaOCA1o="
            },
            {
              "key": "Labels",
              "value": "{keypair:d48su38t49avt161vnkg}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2025-11-10T03:45:45.274-08:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-2"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:fQXYV2AGZR8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDOYl9nmr2wlGkRH8IJPjNNZYy+1QnLxudDBQe65h/csl8teCvQU9YpDd24XUrqk2Vj/rlE5MlbQrXX5RByUjolasCVGy7/QLXH3BwmqNuoHvYyFTpnQlFBBDg7Dq+U3T1FdGRdSY/ZwRuz6CQC1Vg+kIRweajbtxOyMIKGHRrHAI8Xw1GafDod/Wz83sLbm+F3ZB0UnVnFeu8xNdaqfdupoI3Rqgw0sNgnBahvOXhHGsnX8WgxrhDeRqsBO5SYLcqi614dNbXS3LwXG6XH6GBYB/M2TJLNTvd/mH7cEUxTsTttiFreF8WZE64oGHGPQgAoXSIzaS0/F+W9k6rgFCqkcQwaecqYvcxQp8u0WOChBcjECt1cHtHvQmPdgdI5URK1DxlJoTEionAslQBP6McxRE3hXyDH5RTM9pdpaFxRXY5AHUIjpKCuRZeRY4Xa4mnB0X9IMk7U1j6S05A4plKeCJtnJsuw6hnnDNVfQ24J0OsywdBmDrC5GRN7OPTZtqoelCqMHJ6h9ii7WB5qUz7ITLUQ2KSl4cmXqJ0dfOCi/vw5/vmH7iLgXfZwbzfLCAXMSGp4KopRxTvoorYxGLLcKhs1CrWHEsusF1QmXRIICrSKkq7P6yOn3Iq4OmwEZ1L9Z9RXABRZ+RPaMVwymnCtDJ54UoJYkYlVvxu0uOzwSQ== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "d48t0k8t49avt161vnpg"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.193.4,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:kIYgYOib8Xo=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d48strot49avt161vnjg,networkIP:10.0.1.4,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d48strot49avt161vnk0}"
            },
            {
              "key": "ResourceStatus",
              "value": "{}"
            },
            {
              "key": "SatisfiesPzi",
              "value": "true"
            },
            {
              "key": "SatisfiesPzs",
              "value": "false"
            },
            {
              "key": "Scheduling",
              "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d48t0k8t49avt161vnpg"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
            },
            {
              "key": "ShieldedInstanceConfig",
              "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
            },
            {
              "key": "ShieldedInstanceIntegrityPolicy",
              "value": "{updateAutoLearnPolicy:true}"
            },
            {
              "key": "StartRestricted",
              "value": "false"
            },
            {
              "key": "Status",
              "value": "RUNNING"
            },
            {
              "key": "Tags",
              "value": "{fingerprint:5KCAhqva_iI=,items:[d48svkgt49avt161vnm0]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
          "uid": "d48t0k8t49avt161vnog",
          "cspResourceName": "d48t0k8t49avt161vnog",
          "cspResourceId": "d48t0k8t49avt161vnog",
          "name": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
          "subGroupId": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-11-10 11:46:08",
          "label": {
            "sourceMachineId": "ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2025-11-10 11:46:08",
            "sys.cspResourceId": "d48t0k8t49avt161vnog",
            "sys.cspResourceName": "d48t0k8t49avt161vnog",
            "sys.id": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
            "sys.uid": "d48t0k8t49avt161vnog"
          },
          "description": "a recommended virtual machine 02 for ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
          "region": {
            "Region": "asia-northeast3",
            "Zone": "asia-northeast3-a"
          },
          "publicIP": "34.50.23.84",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.2",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": "50",
          "rootDiskName": "",
          "connectionName": "gcp-asia-northeast3",
          "connectionConfig": {
            "configName": "gcp-asia-northeast3",
            "providerName": "gcp",
            "driverName": "gcp-driver-v1.0.so",
            "credentialName": "gcp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "gcp-asia-northeast3",
            "regionZoneInfo": {
              "assignedRegion": "asia-northeast3",
              "assignedZone": "asia-northeast3-a"
            },
            "regionDetail": {
              "regionId": "asia-northeast3",
              "regionName": "asia-northeast3",
              "description": "Seoul South Korea",
              "location": {
                "display": "South Korea (Seoul)",
                "latitude": 37.2,
                "longitude": 127
              },
              "zones": [
                "asia-northeast3-a",
                "asia-northeast3-b",
                "asia-northeast3-c"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "gcp+asia-northeast3+e2-standard-4",
          "cspSpecName": "e2-standard-4",
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "d48strot49avt161vnjg",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "d48strot49avt161vnk0",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d48su38t49avt161vnkg",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-11-10T11:46:21Z",
              "completedTime": "2025-11-10T11:46:20Z",
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d48t0k8t49avt161vnog 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "CanIpForward",
              "value": "false"
            },
            {
              "key": "CpuPlatform",
              "value": "Intel Broadwell"
            },
            {
              "key": "CreationTimestamp",
              "value": "2025-11-10T03:45:28.644-08:00"
            },
            {
              "key": "DeletionProtection",
              "value": "false"
            },
            {
              "key": "Description",
              "value": "compute sample instance"
            },
            {
              "key": "Disks",
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d48t0k8t49avt161vnog,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "-mtIceZL_8Y="
            },
            {
              "key": "Id",
              "value": "2038569319180037815"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "GcPBmaOCA1o="
            },
            {
              "key": "Labels",
              "value": "{keypair:d48su38t49avt161vnkg}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2025-11-10T03:45:39.688-08:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:fQXYV2AGZR8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDOYl9nmr2wlGkRH8IJPjNNZYy+1QnLxudDBQe65h/csl8teCvQU9YpDd24XUrqk2Vj/rlE5MlbQrXX5RByUjolasCVGy7/QLXH3BwmqNuoHvYyFTpnQlFBBDg7Dq+U3T1FdGRdSY/ZwRuz6CQC1Vg+kIRweajbtxOyMIKGHRrHAI8Xw1GafDod/Wz83sLbm+F3ZB0UnVnFeu8xNdaqfdupoI3Rqgw0sNgnBahvOXhHGsnX8WgxrhDeRqsBO5SYLcqi614dNbXS3LwXG6XH6GBYB/M2TJLNTvd/mH7cEUxTsTttiFreF8WZE64oGHGPQgAoXSIzaS0/F+W9k6rgFCqkcQwaecqYvcxQp8u0WOChBcjECt1cHtHvQmPdgdI5URK1DxlJoTEionAslQBP6McxRE3hXyDH5RTM9pdpaFxRXY5AHUIjpKCuRZeRY4Xa4mnB0X9IMk7U1j6S05A4plKeCJtnJsuw6hnnDNVfQ24J0OsywdBmDrC5GRN7OPTZtqoelCqMHJ6h9ii7WB5qUz7ITLUQ2KSl4cmXqJ0dfOCi/vw5/vmH7iLgXfZwbzfLCAXMSGp4KopRxTvoorYxGLLcKhs1CrWHEsusF1QmXRIICrSKkq7P6yOn3Iq4OmwEZ1L9Z9RXABRZ+RPaMVwymnCtDJ54UoJYkYlVvxu0uOzwSQ== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "d48t0k8t49avt161vnog"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.50.23.84,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:f9JJzLXhNnA=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d48strot49avt161vnjg,networkIP:10.0.1.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d48strot49avt161vnk0}"
            },
            {
              "key": "ResourceStatus",
              "value": "{}"
            },
            {
              "key": "SatisfiesPzi",
              "value": "true"
            },
            {
              "key": "SatisfiesPzs",
              "value": "false"
            },
            {
              "key": "Scheduling",
              "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d48t0k8t49avt161vnog"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
            },
            {
              "key": "ShieldedInstanceConfig",
              "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
            },
            {
              "key": "ShieldedInstanceIntegrityPolicy",
              "value": "{updateAutoLearnPolicy:true}"
            },
            {
              "key": "StartRestricted",
              "value": "false"
            },
            {
              "key": "Status",
              "value": "RUNNING"
            },
            {
              "key": "Tags",
              "value": "{fingerprint:TnO5353i2ig=,items:[d48supgt49avt161vnlg]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
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
            "vmId": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
            "vmIp": "34.50.23.84",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d48t0k8t49avt161vnog 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
            "vmIp": "34.50.33.86",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d48t0k8t49avt161vnng 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
            "vmIp": "34.64.193.4",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d48t0k8t49avt161vnpg 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "d48t0k8t49avt161vnmg",
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
    "sys.uid": "d48t0k8t49avt161vnmg"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
      "uid": "d48t0k8t49avt161vnng",
      "cspResourceName": "d48t0k8t49avt161vnng",
      "cspResourceId": "d48t0k8t49avt161vnng",
      "name": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
      "subGroupId": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-11-10 11:46:09",
      "label": {
        "sourceMachineId": "ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2025-11-10 11:46:09",
        "sys.cspResourceId": "d48t0k8t49avt161vnng",
        "sys.cspResourceName": "d48t0k8t49avt161vnng",
        "sys.id": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
        "sys.uid": "d48t0k8t49avt161vnng"
      },
      "description": "a recommended virtual machine 01 for ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.50.33.86",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.3",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "gcp-asia-northeast3",
      "connectionConfig": {
        "configName": "gcp-asia-northeast3",
        "providerName": "gcp",
        "driverName": "gcp-driver-v1.0.so",
        "credentialName": "gcp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "gcp-asia-northeast3",
        "regionZoneInfo": {
          "assignedRegion": "asia-northeast3",
          "assignedZone": "asia-northeast3-a"
        },
        "regionDetail": {
          "regionId": "asia-northeast3",
          "regionName": "asia-northeast3",
          "description": "Seoul South Korea",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "zones": [
            "asia-northeast3-a",
            "asia-northeast3-b",
            "asia-northeast3-c"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "gcp+asia-northeast3+e2-small",
      "cspSpecName": "e2-small",
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "d48strot49avt161vnjg",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "d48strot49avt161vnk0",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d48su38t49avt161vnkg",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-10T11:46:21Z",
          "completedTime": "2025-11-10T11:46:21Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d48t0k8t49avt161vnng 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2025-11-10T03:45:29.652-08:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d48t0k8t49avt161vnng,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "sSXEs3PF3tk="
        },
        {
          "key": "Id",
          "value": "7899419285426826934"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "GcPBmaOCA1o="
        },
        {
          "key": "Labels",
          "value": "{keypair:d48su38t49avt161vnkg}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2025-11-10T03:45:40.781-08:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-small"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:fQXYV2AGZR8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDOYl9nmr2wlGkRH8IJPjNNZYy+1QnLxudDBQe65h/csl8teCvQU9YpDd24XUrqk2Vj/rlE5MlbQrXX5RByUjolasCVGy7/QLXH3BwmqNuoHvYyFTpnQlFBBDg7Dq+U3T1FdGRdSY/ZwRuz6CQC1Vg+kIRweajbtxOyMIKGHRrHAI8Xw1GafDod/Wz83sLbm+F3ZB0UnVnFeu8xNdaqfdupoI3Rqgw0sNgnBahvOXhHGsnX8WgxrhDeRqsBO5SYLcqi614dNbXS3LwXG6XH6GBYB/M2TJLNTvd/mH7cEUxTsTttiFreF8WZE64oGHGPQgAoXSIzaS0/F+W9k6rgFCqkcQwaecqYvcxQp8u0WOChBcjECt1cHtHvQmPdgdI5URK1DxlJoTEionAslQBP6McxRE3hXyDH5RTM9pdpaFxRXY5AHUIjpKCuRZeRY4Xa4mnB0X9IMk7U1j6S05A4plKeCJtnJsuw6hnnDNVfQ24J0OsywdBmDrC5GRN7OPTZtqoelCqMHJ6h9ii7WB5qUz7ITLUQ2KSl4cmXqJ0dfOCi/vw5/vmH7iLgXfZwbzfLCAXMSGp4KopRxTvoorYxGLLcKhs1CrWHEsusF1QmXRIICrSKkq7P6yOn3Iq4OmwEZ1L9Z9RXABRZ+RPaMVwymnCtDJ54UoJYkYlVvxu0uOzwSQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "d48t0k8t49avt161vnng"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.50.33.86,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:j1JWJtpbWkg=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d48strot49avt161vnjg,networkIP:10.0.1.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d48strot49avt161vnk0}"
        },
        {
          "key": "ResourceStatus",
          "value": "{}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d48t0k8t49avt161vnng"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:Z0ZQ4oN1kFM=,items:[d48su3gt49avt161vnl0]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
      "uid": "d48t0k8t49avt161vnpg",
      "cspResourceName": "d48t0k8t49avt161vnpg",
      "cspResourceId": "d48t0k8t49avt161vnpg",
      "name": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
      "subGroupId": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-11-10 11:46:15",
      "label": {
        "sourceMachineId": "ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2025-11-10 11:46:15",
        "sys.cspResourceId": "d48t0k8t49avt161vnpg",
        "sys.cspResourceName": "d48t0k8t49avt161vnpg",
        "sys.id": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
        "sys.uid": "d48t0k8t49avt161vnpg"
      },
      "description": "a recommended virtual machine 03 for ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.64.193.4",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "gcp-asia-northeast3",
      "connectionConfig": {
        "configName": "gcp-asia-northeast3",
        "providerName": "gcp",
        "driverName": "gcp-driver-v1.0.so",
        "credentialName": "gcp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "gcp-asia-northeast3",
        "regionZoneInfo": {
          "assignedRegion": "asia-northeast3",
          "assignedZone": "asia-northeast3-a"
        },
        "regionDetail": {
          "regionId": "asia-northeast3",
          "regionName": "asia-northeast3",
          "description": "Seoul South Korea",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "zones": [
            "asia-northeast3-a",
            "asia-northeast3-b",
            "asia-northeast3-c"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "gcp+asia-northeast3+e2-standard-2",
      "cspSpecName": "e2-standard-2",
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "d48strot49avt161vnjg",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "d48strot49avt161vnk0",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d48su38t49avt161vnkg",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-10T11:46:21Z",
          "completedTime": "2025-11-10T11:46:24Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d48t0k8t49avt161vnpg 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2025-11-10T03:45:30.352-08:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d48t0k8t49avt161vnpg,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "FeW8WkYqYmM="
        },
        {
          "key": "Id",
          "value": "6036202776671579829"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "GcPBmaOCA1o="
        },
        {
          "key": "Labels",
          "value": "{keypair:d48su38t49avt161vnkg}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2025-11-10T03:45:45.274-08:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-2"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:fQXYV2AGZR8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDOYl9nmr2wlGkRH8IJPjNNZYy+1QnLxudDBQe65h/csl8teCvQU9YpDd24XUrqk2Vj/rlE5MlbQrXX5RByUjolasCVGy7/QLXH3BwmqNuoHvYyFTpnQlFBBDg7Dq+U3T1FdGRdSY/ZwRuz6CQC1Vg+kIRweajbtxOyMIKGHRrHAI8Xw1GafDod/Wz83sLbm+F3ZB0UnVnFeu8xNdaqfdupoI3Rqgw0sNgnBahvOXhHGsnX8WgxrhDeRqsBO5SYLcqi614dNbXS3LwXG6XH6GBYB/M2TJLNTvd/mH7cEUxTsTttiFreF8WZE64oGHGPQgAoXSIzaS0/F+W9k6rgFCqkcQwaecqYvcxQp8u0WOChBcjECt1cHtHvQmPdgdI5URK1DxlJoTEionAslQBP6McxRE3hXyDH5RTM9pdpaFxRXY5AHUIjpKCuRZeRY4Xa4mnB0X9IMk7U1j6S05A4plKeCJtnJsuw6hnnDNVfQ24J0OsywdBmDrC5GRN7OPTZtqoelCqMHJ6h9ii7WB5qUz7ITLUQ2KSl4cmXqJ0dfOCi/vw5/vmH7iLgXfZwbzfLCAXMSGp4KopRxTvoorYxGLLcKhs1CrWHEsusF1QmXRIICrSKkq7P6yOn3Iq4OmwEZ1L9Z9RXABRZ+RPaMVwymnCtDJ54UoJYkYlVvxu0uOzwSQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "d48t0k8t49avt161vnpg"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.193.4,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:kIYgYOib8Xo=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d48strot49avt161vnjg,networkIP:10.0.1.4,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d48strot49avt161vnk0}"
        },
        {
          "key": "ResourceStatus",
          "value": "{}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d48t0k8t49avt161vnpg"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:5KCAhqva_iI=,items:[d48svkgt49avt161vnm0]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
      "uid": "d48t0k8t49avt161vnog",
      "cspResourceName": "d48t0k8t49avt161vnog",
      "cspResourceId": "d48t0k8t49avt161vnog",
      "name": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
      "subGroupId": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.2,
        "longitude": 127
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-11-10 11:46:08",
      "label": {
        "sourceMachineId": "ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2025-11-10 11:46:08",
        "sys.cspResourceId": "d48t0k8t49avt161vnog",
        "sys.cspResourceName": "d48t0k8t49avt161vnog",
        "sys.id": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
        "sys.uid": "d48t0k8t49avt161vnog"
      },
      "description": "a recommended virtual machine 02 for ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.50.23.84",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.2",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "gcp-asia-northeast3",
      "connectionConfig": {
        "configName": "gcp-asia-northeast3",
        "providerName": "gcp",
        "driverName": "gcp-driver-v1.0.so",
        "credentialName": "gcp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "gcp-asia-northeast3",
        "regionZoneInfo": {
          "assignedRegion": "asia-northeast3",
          "assignedZone": "asia-northeast3-a"
        },
        "regionDetail": {
          "regionId": "asia-northeast3",
          "regionName": "asia-northeast3",
          "description": "Seoul South Korea",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          },
          "zones": [
            "asia-northeast3-a",
            "asia-northeast3-b",
            "asia-northeast3-c"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "gcp+asia-northeast3+e2-standard-4",
      "cspSpecName": "e2-standard-4",
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20251023",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "d48strot49avt161vnjg",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "d48strot49avt161vnk0",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d48su38t49avt161vnkg",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-11-10T11:46:21Z",
          "completedTime": "2025-11-10T11:46:20Z",
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d48t0k8t49avt161vnog 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "CanIpForward",
          "value": "false"
        },
        {
          "key": "CpuPlatform",
          "value": "Intel Broadwell"
        },
        {
          "key": "CreationTimestamp",
          "value": "2025-11-10T03:45:28.644-08:00"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Description",
          "value": "compute sample instance"
        },
        {
          "key": "Disks",
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d48t0k8t49avt161vnog,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "-mtIceZL_8Y="
        },
        {
          "key": "Id",
          "value": "2038569319180037815"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "GcPBmaOCA1o="
        },
        {
          "key": "Labels",
          "value": "{keypair:d48su38t49avt161vnkg}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2025-11-10T03:45:39.688-08:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:fQXYV2AGZR8=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDOYl9nmr2wlGkRH8IJPjNNZYy+1QnLxudDBQe65h/csl8teCvQU9YpDd24XUrqk2Vj/rlE5MlbQrXX5RByUjolasCVGy7/QLXH3BwmqNuoHvYyFTpnQlFBBDg7Dq+U3T1FdGRdSY/ZwRuz6CQC1Vg+kIRweajbtxOyMIKGHRrHAI8Xw1GafDod/Wz83sLbm+F3ZB0UnVnFeu8xNdaqfdupoI3Rqgw0sNgnBahvOXhHGsnX8WgxrhDeRqsBO5SYLcqi614dNbXS3LwXG6XH6GBYB/M2TJLNTvd/mH7cEUxTsTttiFreF8WZE64oGHGPQgAoXSIzaS0/F+W9k6rgFCqkcQwaecqYvcxQp8u0WOChBcjECt1cHtHvQmPdgdI5URK1DxlJoTEionAslQBP6McxRE3hXyDH5RTM9pdpaFxRXY5AHUIjpKCuRZeRY4Xa4mnB0X9IMk7U1j6S05A4plKeCJtnJsuw6hnnDNVfQ24J0OsywdBmDrC5GRN7OPTZtqoelCqMHJ6h9ii7WB5qUz7ITLUQ2KSl4cmXqJ0dfOCi/vw5/vmH7iLgXfZwbzfLCAXMSGp4KopRxTvoorYxGLLcKhs1CrWHEsusF1QmXRIICrSKkq7P6yOn3Iq4OmwEZ1L9Z9RXABRZ+RPaMVwymnCtDJ54UoJYkYlVvxu0uOzwSQ== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "d48t0k8t49avt161vnog"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.50.23.84,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:f9JJzLXhNnA=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d48strot49avt161vnjg,networkIP:10.0.1.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d48strot49avt161vnk0}"
        },
        {
          "key": "ResourceStatus",
          "value": "{}"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "Scheduling",
          "value": "{automaticRestart:true,onHostMaintenance:MIGRATE,provisioningModel:STANDARD}"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d48t0k8t49avt161vnog"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
        },
        {
          "key": "ShieldedInstanceConfig",
          "value": "{enableIntegrityMonitoring:true,enableVtpm:true}"
        },
        {
          "key": "ShieldedInstanceIntegrityPolicy",
          "value": "{updateAutoLearnPolicy:true}"
        },
        {
          "key": "StartRestricted",
          "value": "false"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "Tags",
          "value": "{fingerprint:TnO5353i2ig=,items:[d48supgt49avt161vnlg]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
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
        "vmId": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1",
        "vmIp": "34.50.23.84",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d48t0k8t49avt161vnog 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1",
        "vmIp": "34.50.33.86",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d48t0k8t49avt161vnng 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1",
        "vmIp": "34.64.193.4",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d48t0k8t49avt161vnpg 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "output": "Linux d48t0k8t49avt161vnng 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "34.50.33.86",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c",
      "testOrder": 1,
      "userName": "cb-user",
      "vmId": "migrated-ec2241d9-56e0-d4d7-5d88-07af0e2b7a9c-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux d48t0k8t49avt161vnpg 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "34.64.193.4",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4",
      "testOrder": 2,
      "userName": "cb-user",
      "vmId": "migrated-ec26f82e-8ce4-1bb1-7a14-c379e4d50ec4-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux d48t0k8t49avt161vnog 6.8.0-1042-gcp #45~22.04.1-Ubuntu SMP Tue Oct  7 19:06:40 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "34.50.23.84",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939",
      "testOrder": 3,
      "userName": "cb-user",
      "vmId": "migrated-ec2db2cd-7f2b-f86d-97f8-6b5ea91ea939-1"
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

