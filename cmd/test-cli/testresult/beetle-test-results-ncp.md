# CM-Beetle test results for NCP

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with NCP cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.6+ (d76d44e)
- cm-model: v0.0.15
- CB-Tumblebug: v0.11.19
- CB-Spider: v0.11.16
- CB-MapUI: v0.11.19
- Target CSP: NCP
- Target Region: kr
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: December 10, 2025
- Test Time: 16:45:46 KST
- Test Execution: 2025-12-10 16:45:46 KST

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for NCP

### Test Results Summary

| Test | Endpoint | Status | Duration | Details |
|------|----------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/vmInfra` | ✅ **PASS** | 781ms | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 8m26.8s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 856ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 150ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 885ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 2m25.769s | Pass |

**Overall Result**: 7/7 tests passed ✅

**Total Duration**: 12m26.446076104s

*Test executed on December 10, 2025 at 16:45:46 KST (2025-12-10 16:45:46 KST) using CM-Beetle automated test CLI*

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
    "csp": "ncp",
    "region": "kr"
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
      "description": "Candidate #1 | partially-matched | Overall Match Rate: Min=50.0% Max=100.0% Avg=86.1% | VMs: 3 total, 0 matched, 3 acceptable",
      "targetCloud": {
        "csp": "ncp",
        "region": "kr"
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
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+ci2-g3",
            "imageId": "23214590",
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
            "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+s4-g3a",
            "imageId": "23214590",
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
            "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+s2-g3a",
            "imageId": "23214590",
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
        "connectionName": "ncp-kr",
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
        "connectionName": "ncp-kr",
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
          "id": "ncp+kr+ci2-g3",
          "uid": "d4sgb4a5npi2mb9tbabg",
          "cspSpecName": "ci2-g3",
          "name": "ncp+kr+ci2-g3",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 4,
          "diskSizeGB": -1,
          "costPerHour": 0.073,
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
          "rootDiskSize": "-1",
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "ci2-g3"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "2"
            },
            {
              "key": "MemorySize",
              "value": "4294967296"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "5250"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "84934656"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.CPU.C002.M004.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 2EA, Memory 4GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318,16946033"
            }
          ]
        },
        {
          "id": "ncp+kr+s4-g3a",
          "uid": "d4sgb4a5npi2mb9tba6g",
          "cspSpecName": "s4-g3a",
          "name": "ncp+kr+s4-g3a",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1747,
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
          "rootDiskSize": "-1",
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "s4-g3a"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "4"
            },
            {
              "key": "MemorySize",
              "value": "17179869184"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "7875"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "150994944"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.AMD.STAND.C004.M016.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 4EA, Memory 16GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318"
            }
          ]
        },
        {
          "id": "ncp+kr+s2-g3a",
          "uid": "d4sgb4a5npi2mb9tbahg",
          "cspSpecName": "s2-g3a",
          "name": "ncp+kr+s2-g3a",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.0848,
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
          "rootDiskSize": "-1",
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "s2-g3a"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "2"
            },
            {
              "key": "MemorySize",
              "value": "8589934592"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "4725"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "84934656"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.AMD.STAND.C002.M008.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 2EA, Memory 8GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318"
            }
          ]
        }
      ],
      "targetVmOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "ncp",
          "cspImageName": "23214590",
          "regionList": [
            "kr"
          ],
          "id": "23214590",
          "uid": "d4sgb625npi2mb9tbh70",
          "name": "23214590",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "ncp-kr",
          "infraType": "",
          "fetchedTime": "2025.12.10 05:31:04 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)",
          "osDiskType": "Common BlockStorage 1",
          "osDiskSizeGB": 10,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ServerImageNo",
              "value": "23214590"
            },
            {
              "key": "ServerImageName",
              "value": "ubuntu-22.04-base"
            },
            {
              "key": "ServerImageDescription",
              "value": "kernel version : 5.15.0-140-generic"
            },
            {
              "key": "ServerImageProductCode",
              "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
            },
            {
              "key": "ServerImageType",
              "value": "{code:NCP,codeName:NCP 서버이미지}"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "OsCategoryType",
              "value": "{code:LINUX,codeName:LINUX}"
            },
            {
              "key": "OsType",
              "value": "{code:UBUNTU,codeName:UBUNTU}"
            },
            {
              "key": "ServerImageStatus",
              "value": "{code:CREAT,codeName:NSI CREATED state}"
            },
            {
              "key": "ServerImageOperation",
              "value": "{code:NULL,codeName:NSI NULL OP}"
            },
            {
              "key": "ServerImageStatusName",
              "value": "created"
            },
            {
              "key": "CreateDate",
              "value": "2024-03-21T18:22:55+0900"
            },
            {
              "key": "ShareStatus",
              "value": "{code:NULL,codeName:NSI Share NULL State}"
            },
            {
              "key": "BlockStorageMappingList",
              "value": "{order:0,blockStorageSnapshotInstanceNo:23214591,blockStorageSnapshotName:snapshot of ubuntu-22.04-base,blockStorageSize:10737418240,blockStorageVolumeType:{code:CB1,codeName:Common BlockStorage 1},iops:100,throughput:104857600,isEncryptedVolume:false}"
            }
          ],
          "systemLabel": "",
          "description": "",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "ncp-kr",
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
          "connectionName": "ncp-kr",
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
          "connectionName": "ncp-kr",
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
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=25.0% Max=100.0% Avg=83.3% | VMs: 3 total, 0 matched, 3 acceptable",
      "targetCloud": {
        "csp": "ncp",
        "region": "kr"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=25.0% Image=75.0%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+s2-g3a",
            "imageId": "23214590",
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
            "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+s4-g3",
            "imageId": "23214590",
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
            "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "ncp-kr",
            "specId": "ncp+kr+s2-g3",
            "imageId": "23214590",
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
        "connectionName": "ncp-kr",
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
        "connectionName": "ncp-kr",
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
          "id": "ncp+kr+ci2-g3",
          "uid": "d4sgb4a5npi2mb9tbabg",
          "cspSpecName": "ci2-g3",
          "name": "ncp+kr+ci2-g3",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 4,
          "diskSizeGB": -1,
          "costPerHour": 0.073,
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
          "rootDiskSize": "-1",
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "ci2-g3"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "2"
            },
            {
              "key": "MemorySize",
              "value": "4294967296"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "5250"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "84934656"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.CPU.C002.M004.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 2EA, Memory 4GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318,16946033"
            }
          ]
        },
        {
          "id": "ncp+kr+s4-g3a",
          "uid": "d4sgb4a5npi2mb9tba6g",
          "cspSpecName": "s4-g3a",
          "name": "ncp+kr+s4-g3a",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1747,
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
          "rootDiskSize": "-1",
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "s4-g3a"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "4"
            },
            {
              "key": "MemorySize",
              "value": "17179869184"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "7875"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "150994944"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.AMD.STAND.C004.M016.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 4EA, Memory 16GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318"
            }
          ]
        },
        {
          "id": "ncp+kr+s2-g3a",
          "uid": "d4sgb4a5npi2mb9tbahg",
          "cspSpecName": "s2-g3a",
          "name": "ncp+kr+s2-g3a",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.0848,
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
          "rootDiskSize": "-1",
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "s2-g3a"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "2"
            },
            {
              "key": "MemorySize",
              "value": "8589934592"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "4725"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "84934656"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.AMD.STAND.C002.M008.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 2EA, Memory 8GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318"
            }
          ]
        },
        {
          "id": "ncp+kr+s4-g3",
          "uid": "d4sgb4a5npi2mb9tbal0",
          "cspSpecName": "s4-g3",
          "name": "ncp+kr+s4-g3",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1747,
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
          "rootDiskSize": "-1",
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "s4-g3"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "4"
            },
            {
              "key": "MemorySize",
              "value": "17179869184"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "7875"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "150994944"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.STAND.C004.M016.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 4EA, Memory 16GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318,16946033"
            }
          ]
        },
        {
          "id": "ncp+kr+s2-g3",
          "uid": "d4sgb4a5npi2mb9tbb1g",
          "cspSpecName": "s2-g3",
          "name": "ncp+kr+s2-g3",
          "namespace": "system",
          "connectionName": "ncp-kr",
          "providerName": "ncp",
          "regionName": "kr",
          "regionLatitude": 37.4754,
          "regionLongitude": 126.8831,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.0848,
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
          "rootDiskSize": "-1",
          "systemLabel": "from-assets",
          "details": [
            {
              "key": "ServerSpecCode",
              "value": "s2-g3"
            },
            {
              "key": "GenerationCode",
              "value": "G3"
            },
            {
              "key": "CpuCount",
              "value": "2"
            },
            {
              "key": "MemorySize",
              "value": "8589934592"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "BlockStorageMaxCount",
              "value": "20"
            },
            {
              "key": "BlockStorageMaxIops",
              "value": "4725"
            },
            {
              "key": "BlockStorageMaxThroughput",
              "value": "84934656"
            },
            {
              "key": "NetworkPerformance",
              "value": "1000000000"
            },
            {
              "key": "NetworkInterfaceMaxCount",
              "value": "3"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.STAND.C002.M008.G003"
            },
            {
              "key": "ServerSpecDescription",
              "value": "vCPU 2EA, Memory 8GB"
            },
            {
              "key": "CorrespondingImageIds",
              "value": "107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318,16946033"
            }
          ]
        }
      ],
      "targetVmOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "ncp",
          "cspImageName": "23214590",
          "regionList": [
            "kr"
          ],
          "id": "23214590",
          "uid": "d4sgb625npi2mb9tbh70",
          "name": "23214590",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "ncp-kr",
          "infraType": "",
          "fetchedTime": "2025.12.10 05:31:04 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu-22.04-base (Hypervisor:KVM)",
          "osDiskType": "Common BlockStorage 1",
          "osDiskSizeGB": 10,
          "imageStatus": "Available",
          "details": [
            {
              "key": "ServerImageNo",
              "value": "23214590"
            },
            {
              "key": "ServerImageName",
              "value": "ubuntu-22.04-base"
            },
            {
              "key": "ServerImageDescription",
              "value": "kernel version : 5.15.0-140-generic"
            },
            {
              "key": "ServerImageProductCode",
              "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
            },
            {
              "key": "ServerImageType",
              "value": "{code:NCP,codeName:NCP 서버이미지}"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "CpuArchitectureType",
              "value": "{code:X86_64,codeName:x86 64bit}"
            },
            {
              "key": "OsCategoryType",
              "value": "{code:LINUX,codeName:LINUX}"
            },
            {
              "key": "OsType",
              "value": "{code:UBUNTU,codeName:UBUNTU}"
            },
            {
              "key": "ServerImageStatus",
              "value": "{code:CREAT,codeName:NSI CREATED state}"
            },
            {
              "key": "ServerImageOperation",
              "value": "{code:NULL,codeName:NSI NULL OP}"
            },
            {
              "key": "ServerImageStatusName",
              "value": "created"
            },
            {
              "key": "CreateDate",
              "value": "2024-03-21T18:22:55+0900"
            },
            {
              "key": "ShareStatus",
              "value": "{code:NULL,codeName:NSI Share NULL State}"
            },
            {
              "key": "BlockStorageMappingList",
              "value": "{order:0,blockStorageSnapshotInstanceNo:23214591,blockStorageSnapshotName:snapshot of ubuntu-22.04-base,blockStorageSize:10737418240,blockStorageVolumeType:{code:CB1,codeName:Common BlockStorage 1},iops:100,throughput:104857600,isEncryptedVolume:false}"
            }
          ],
          "systemLabel": "",
          "description": "",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-sg-01",
          "connectionName": "ncp-kr",
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
          "connectionName": "ncp-kr",
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
          "connectionName": "ncp-kr",
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
  "uid": "d4siboa5npi2mbdfo0n0",
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
    "sys.uid": "d4siboa5npi2mbdfo0n0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "d4siboa5npi2mbdfo0o0",
      "cspResourceName": "d4siboa5npi2mbdfo0o0",
      "cspResourceId": "115562572",
      "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-10 07:53:49",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=75.0%",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "223.130.141.142",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.8",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+ci2-g3",
      "cspSpecName": "ci2-g3",
      "imageId": "23214590",
      "cspImageName": "23214590",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "130336",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "278279",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4siala5npi2mbdfo0l0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-10T07:53:55Z",
          "completedTime": "2025-12-10T07:53:56Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4siboa5npi2mbdfo0o0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "115562572"
        },
        {
          "key": "ServerName",
          "value": "d4siboa5npi2mbdfo0o0"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize",
          "value": "4294967296"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "d4siala5npi2mbdfo0l0"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2025-12-10T16:49:00+0900"
        },
        {
          "key": "Uptime",
          "value": "2025-12-10T16:51:39+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.CPU.C002.M004.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "130336"
        },
        {
          "key": "SubnetNo",
          "value": "278279"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5122440"
        },
        {
          "key": "InitScriptNo",
          "value": "153611"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:CPU,codeName:CPU-Intensive}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "ci2-g3"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "d4siboa5npi2mbdfo0q0",
      "cspResourceName": "d4siboa5npi2mbdfo0q0",
      "cspResourceId": "115562565",
      "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-10 07:53:49",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "175.45.203.195",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.7",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+s2-g3a",
      "cspSpecName": "s2-g3a",
      "imageId": "23214590",
      "cspImageName": "23214590",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "130336",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "278279",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4siala5npi2mbdfo0l0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-10T07:53:55Z",
          "completedTime": "2025-12-10T07:53:56Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4siboa5npi2mbdfo0q0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "115562565"
        },
        {
          "key": "ServerName",
          "value": "d4siboa5npi2mbdfo0q0"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize",
          "value": "8589934592"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "d4siala5npi2mbdfo0l0"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2025-12-10T16:48:59+0900"
        },
        {
          "key": "Uptime",
          "value": "2025-12-10T16:51:38+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.AMD.STAND.C002.M008.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "130336"
        },
        {
          "key": "SubnetNo",
          "value": "278279"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5122439"
        },
        {
          "key": "InitScriptNo",
          "value": "153610"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:STAND,codeName:Standard}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "s2-g3a"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "d4siboa5npi2mbdfo0p0",
      "cspResourceName": "d4siboa5npi2mbdfo0p0",
      "cspResourceId": "115562559",
      "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-10 07:53:26",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "223.130.152.254",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+s4-g3a",
      "cspSpecName": "s4-g3a",
      "imageId": "23214590",
      "cspImageName": "23214590",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "130336",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "278279",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4siala5npi2mbdfo0l0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-10T07:53:55Z",
          "completedTime": "2025-12-10T07:53:56Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4siboa5npi2mbdfo0p0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "115562559"
        },
        {
          "key": "ServerName",
          "value": "d4siboa5npi2mbdfo0p0"
        },
        {
          "key": "CpuCount",
          "value": "4"
        },
        {
          "key": "MemorySize",
          "value": "17179869184"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "d4siala5npi2mbdfo0l0"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2025-12-10T16:48:59+0900"
        },
        {
          "key": "Uptime",
          "value": "2025-12-10T16:51:18+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.AMD.STAND.C004.M016.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "130336"
        },
        {
          "key": "SubnetNo",
          "value": "278279"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5122438"
        },
        {
          "key": "InitScriptNo",
          "value": "153609"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:STAND,codeName:Standard}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "s4-g3a"
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
        "vmIp": "223.130.152.254",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4siboa5npi2mbdfo0p0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "vmIp": "175.45.203.195",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4siboa5npi2mbdfo0q0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vmIp": "223.130.141.142",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4siboa5npi2mbdfo0o0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "d4siboa5npi2mbdfo0n0",
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
        "sys.uid": "d4siboa5npi2mbdfo0n0"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "d4siboa5npi2mbdfo0o0",
          "cspResourceName": "d4siboa5npi2mbdfo0o0",
          "cspResourceId": "115562572",
          "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
          "subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-12-10 07:53:49",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2025-12-10 07:53:49",
            "sys.cspResourceId": "115562572",
            "sys.cspResourceName": "d4siboa5npi2mbdfo0o0",
            "sys.id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.uid": "d4siboa5npi2mbdfo0o0"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=75.0%",
          "region": {
            "Region": "KR",
            "Zone": "KR-1"
          },
          "publicIP": "223.130.141.142",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.8",
          "privateDNS": "",
          "rootDiskType": "HDD",
          "rootDiskSize": "50",
          "rootDiskName": "",
          "connectionName": "ncp-kr",
          "connectionConfig": {
            "configName": "ncp-kr",
            "providerName": "ncp",
            "driverName": "ncp-driver-v1.0.so",
            "credentialName": "ncp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ncp-kr",
            "regionZoneInfo": {
              "assignedRegion": "KR",
              "assignedZone": "KR-1"
            },
            "regionDetail": {
              "regionId": "KR",
              "regionName": "kr",
              "description": "Korea 1",
              "location": {
                "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
                "latitude": 37.4754,
                "longitude": 126.8831
              },
              "zones": [
                "KR-1",
                "KR-2"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ncp+kr+ci2-g3",
          "cspSpecName": "ci2-g3",
          "imageId": "23214590",
          "cspImageName": "23214590",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "130336",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "278279",
          "networkInterface": "eth0",
          "securityGroupIds": [
            "mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d4siala5npi2mbdfo0l0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-12-10T07:53:55Z",
              "completedTime": "2025-12-10T07:53:56Z",
              "elapsedTime": 1,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4siboa5npi2mbdfo0o0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ServerInstanceNo",
              "value": "115562572"
            },
            {
              "key": "ServerName",
              "value": "d4siboa5npi2mbdfo0o0"
            },
            {
              "key": "CpuCount",
              "value": "2"
            },
            {
              "key": "MemorySize",
              "value": "4294967296"
            },
            {
              "key": "PlatformType",
              "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
            },
            {
              "key": "LoginKeyName",
              "value": "d4siala5npi2mbdfo0l0"
            },
            {
              "key": "ServerInstanceStatus",
              "value": "{code:RUN,codeName:서버 RUN 상태}"
            },
            {
              "key": "ServerInstanceOperation",
              "value": "{code:NULL,codeName:서버 NULL OP}"
            },
            {
              "key": "ServerInstanceStatusName",
              "value": "running"
            },
            {
              "key": "CreateDate",
              "value": "2025-12-10T16:49:00+0900"
            },
            {
              "key": "Uptime",
              "value": "2025-12-10T16:51:39+0900"
            },
            {
              "key": "ServerImageProductCode",
              "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.CPU.C002.M004.G003"
            },
            {
              "key": "IsProtectServerTermination",
              "value": "false"
            },
            {
              "key": "ZoneCode",
              "value": "KR-1"
            },
            {
              "key": "RegionCode",
              "value": "KR"
            },
            {
              "key": "VpcNo",
              "value": "130336"
            },
            {
              "key": "SubnetNo",
              "value": "278279"
            },
            {
              "key": "NetworkInterfaceNoList",
              "value": "5122440"
            },
            {
              "key": "InitScriptNo",
              "value": "153611"
            },
            {
              "key": "ServerInstanceType",
              "value": "{code:CPU,codeName:CPU-Intensive}"
            },
            {
              "key": "BaseBlockStorageDiskType",
              "value": "{code:NET,codeName:네트웍 스토리지}"
            },
            {
              "key": "BaseBlockStorageDiskDetailType",
              "value": "{code:SSD,codeName:SSD}"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "ServerImageNo",
              "value": "23214590"
            },
            {
              "key": "ServerSpecCode",
              "value": "ci2-g3"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "d4siboa5npi2mbdfo0q0",
          "cspResourceName": "d4siboa5npi2mbdfo0q0",
          "cspResourceId": "115562565",
          "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-12-10 07:53:49",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2025-12-10 07:53:49",
            "sys.cspResourceId": "115562565",
            "sys.cspResourceName": "d4siboa5npi2mbdfo0q0",
            "sys.id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.uid": "d4siboa5npi2mbdfo0q0"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "Region": "KR",
            "Zone": "KR-1"
          },
          "publicIP": "175.45.203.195",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.7",
          "privateDNS": "",
          "rootDiskType": "HDD",
          "rootDiskSize": "50",
          "rootDiskName": "",
          "connectionName": "ncp-kr",
          "connectionConfig": {
            "configName": "ncp-kr",
            "providerName": "ncp",
            "driverName": "ncp-driver-v1.0.so",
            "credentialName": "ncp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ncp-kr",
            "regionZoneInfo": {
              "assignedRegion": "KR",
              "assignedZone": "KR-1"
            },
            "regionDetail": {
              "regionId": "KR",
              "regionName": "kr",
              "description": "Korea 1",
              "location": {
                "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
                "latitude": 37.4754,
                "longitude": 126.8831
              },
              "zones": [
                "KR-1",
                "KR-2"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ncp+kr+s2-g3a",
          "cspSpecName": "s2-g3a",
          "imageId": "23214590",
          "cspImageName": "23214590",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "130336",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "278279",
          "networkInterface": "eth0",
          "securityGroupIds": [
            "mig-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d4siala5npi2mbdfo0l0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-12-10T07:53:55Z",
              "completedTime": "2025-12-10T07:53:56Z",
              "elapsedTime": 1,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4siboa5npi2mbdfo0q0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ServerInstanceNo",
              "value": "115562565"
            },
            {
              "key": "ServerName",
              "value": "d4siboa5npi2mbdfo0q0"
            },
            {
              "key": "CpuCount",
              "value": "2"
            },
            {
              "key": "MemorySize",
              "value": "8589934592"
            },
            {
              "key": "PlatformType",
              "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
            },
            {
              "key": "LoginKeyName",
              "value": "d4siala5npi2mbdfo0l0"
            },
            {
              "key": "ServerInstanceStatus",
              "value": "{code:RUN,codeName:서버 RUN 상태}"
            },
            {
              "key": "ServerInstanceOperation",
              "value": "{code:NULL,codeName:서버 NULL OP}"
            },
            {
              "key": "ServerInstanceStatusName",
              "value": "running"
            },
            {
              "key": "CreateDate",
              "value": "2025-12-10T16:48:59+0900"
            },
            {
              "key": "Uptime",
              "value": "2025-12-10T16:51:38+0900"
            },
            {
              "key": "ServerImageProductCode",
              "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.AMD.STAND.C002.M008.G003"
            },
            {
              "key": "IsProtectServerTermination",
              "value": "false"
            },
            {
              "key": "ZoneCode",
              "value": "KR-1"
            },
            {
              "key": "RegionCode",
              "value": "KR"
            },
            {
              "key": "VpcNo",
              "value": "130336"
            },
            {
              "key": "SubnetNo",
              "value": "278279"
            },
            {
              "key": "NetworkInterfaceNoList",
              "value": "5122439"
            },
            {
              "key": "InitScriptNo",
              "value": "153610"
            },
            {
              "key": "ServerInstanceType",
              "value": "{code:STAND,codeName:Standard}"
            },
            {
              "key": "BaseBlockStorageDiskType",
              "value": "{code:NET,codeName:네트웍 스토리지}"
            },
            {
              "key": "BaseBlockStorageDiskDetailType",
              "value": "{code:SSD,codeName:SSD}"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "ServerImageNo",
              "value": "23214590"
            },
            {
              "key": "ServerSpecCode",
              "value": "s2-g3a"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "d4siboa5npi2mbdfo0p0",
          "cspResourceName": "d4siboa5npi2mbdfo0p0",
          "cspResourceId": "115562559",
          "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-12-10 07:53:26",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "ncp-kr",
            "sys.createdTime": "2025-12-10 07:53:26",
            "sys.cspResourceId": "115562559",
            "sys.cspResourceName": "d4siboa5npi2mbdfo0p0",
            "sys.id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.uid": "d4siboa5npi2mbdfo0p0"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "Region": "KR",
            "Zone": "KR-1"
          },
          "publicIP": "223.130.152.254",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "HDD",
          "rootDiskSize": "50",
          "rootDiskName": "",
          "connectionName": "ncp-kr",
          "connectionConfig": {
            "configName": "ncp-kr",
            "providerName": "ncp",
            "driverName": "ncp-driver-v1.0.so",
            "credentialName": "ncp",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ncp-kr",
            "regionZoneInfo": {
              "assignedRegion": "KR",
              "assignedZone": "KR-1"
            },
            "regionDetail": {
              "regionId": "KR",
              "regionName": "kr",
              "description": "Korea 1",
              "location": {
                "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
                "latitude": 37.4754,
                "longitude": 126.8831
              },
              "zones": [
                "KR-1",
                "KR-2"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ncp+kr+s4-g3a",
          "cspSpecName": "s4-g3a",
          "imageId": "23214590",
          "cspImageName": "23214590",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "130336",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "278279",
          "networkInterface": "eth0",
          "securityGroupIds": [
            "mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d4siala5npi2mbdfo0l0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-12-10T07:53:55Z",
              "completedTime": "2025-12-10T07:53:56Z",
              "elapsedTime": 1,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4siboa5npi2mbdfo0p0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ServerInstanceNo",
              "value": "115562559"
            },
            {
              "key": "ServerName",
              "value": "d4siboa5npi2mbdfo0p0"
            },
            {
              "key": "CpuCount",
              "value": "4"
            },
            {
              "key": "MemorySize",
              "value": "17179869184"
            },
            {
              "key": "PlatformType",
              "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
            },
            {
              "key": "LoginKeyName",
              "value": "d4siala5npi2mbdfo0l0"
            },
            {
              "key": "ServerInstanceStatus",
              "value": "{code:RUN,codeName:서버 RUN 상태}"
            },
            {
              "key": "ServerInstanceOperation",
              "value": "{code:NULL,codeName:서버 NULL OP}"
            },
            {
              "key": "ServerInstanceStatusName",
              "value": "running"
            },
            {
              "key": "CreateDate",
              "value": "2025-12-10T16:48:59+0900"
            },
            {
              "key": "Uptime",
              "value": "2025-12-10T16:51:18+0900"
            },
            {
              "key": "ServerImageProductCode",
              "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
            },
            {
              "key": "ServerProductCode",
              "value": "SVR.VSVR.AMD.STAND.C004.M016.G003"
            },
            {
              "key": "IsProtectServerTermination",
              "value": "false"
            },
            {
              "key": "ZoneCode",
              "value": "KR-1"
            },
            {
              "key": "RegionCode",
              "value": "KR"
            },
            {
              "key": "VpcNo",
              "value": "130336"
            },
            {
              "key": "SubnetNo",
              "value": "278279"
            },
            {
              "key": "NetworkInterfaceNoList",
              "value": "5122438"
            },
            {
              "key": "InitScriptNo",
              "value": "153609"
            },
            {
              "key": "ServerInstanceType",
              "value": "{code:STAND,codeName:Standard}"
            },
            {
              "key": "BaseBlockStorageDiskType",
              "value": "{code:NET,codeName:네트웍 스토리지}"
            },
            {
              "key": "BaseBlockStorageDiskDetailType",
              "value": "{code:SSD,codeName:SSD}"
            },
            {
              "key": "HypervisorType",
              "value": "{code:KVM,codeName:KVM}"
            },
            {
              "key": "ServerImageNo",
              "value": "23214590"
            },
            {
              "key": "ServerSpecCode",
              "value": "s4-g3a"
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
            "vmIp": "223.130.152.254",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4siboa5npi2mbdfo0p0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "vmIp": "175.45.203.195",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4siboa5npi2mbdfo0q0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
            "vmIp": "223.130.141.142",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4siboa5npi2mbdfo0o0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "d4siboa5npi2mbdfo0n0",
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
    "sys.uid": "d4siboa5npi2mbdfo0n0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "d4siboa5npi2mbdfo0o0",
      "cspResourceName": "d4siboa5npi2mbdfo0o0",
      "cspResourceId": "115562572",
      "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-10 07:53:49",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ncp-kr",
        "sys.createdTime": "2025-12-10 07:53:49",
        "sys.cspResourceId": "115562572",
        "sys.cspResourceName": "d4siboa5npi2mbdfo0o0",
        "sys.id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.uid": "d4siboa5npi2mbdfo0o0"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=75.0%",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "223.130.141.142",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.8",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+ci2-g3",
      "cspSpecName": "ci2-g3",
      "imageId": "23214590",
      "cspImageName": "23214590",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "130336",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "278279",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4siala5npi2mbdfo0l0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-10T07:53:55Z",
          "completedTime": "2025-12-10T07:53:56Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4siboa5npi2mbdfo0o0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "115562572"
        },
        {
          "key": "ServerName",
          "value": "d4siboa5npi2mbdfo0o0"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize",
          "value": "4294967296"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "d4siala5npi2mbdfo0l0"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2025-12-10T16:49:00+0900"
        },
        {
          "key": "Uptime",
          "value": "2025-12-10T16:51:39+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.CPU.C002.M004.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "130336"
        },
        {
          "key": "SubnetNo",
          "value": "278279"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5122440"
        },
        {
          "key": "InitScriptNo",
          "value": "153611"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:CPU,codeName:CPU-Intensive}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "ci2-g3"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "d4siboa5npi2mbdfo0q0",
      "cspResourceName": "d4siboa5npi2mbdfo0q0",
      "cspResourceId": "115562565",
      "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-10 07:53:49",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ncp-kr",
        "sys.createdTime": "2025-12-10 07:53:49",
        "sys.cspResourceId": "115562565",
        "sys.cspResourceName": "d4siboa5npi2mbdfo0q0",
        "sys.id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.uid": "d4siboa5npi2mbdfo0q0"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "175.45.203.195",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.7",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+s2-g3a",
      "cspSpecName": "s2-g3a",
      "imageId": "23214590",
      "cspImageName": "23214590",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "130336",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "278279",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4siala5npi2mbdfo0l0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-10T07:53:55Z",
          "completedTime": "2025-12-10T07:53:56Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4siboa5npi2mbdfo0q0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "115562565"
        },
        {
          "key": "ServerName",
          "value": "d4siboa5npi2mbdfo0q0"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize",
          "value": "8589934592"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "d4siala5npi2mbdfo0l0"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2025-12-10T16:48:59+0900"
        },
        {
          "key": "Uptime",
          "value": "2025-12-10T16:51:38+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.AMD.STAND.C002.M008.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "130336"
        },
        {
          "key": "SubnetNo",
          "value": "278279"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5122439"
        },
        {
          "key": "InitScriptNo",
          "value": "153610"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:STAND,codeName:Standard}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "s2-g3a"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "d4siboa5npi2mbdfo0p0",
      "cspResourceName": "d4siboa5npi2mbdfo0p0",
      "cspResourceId": "115562559",
      "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-12-10 07:53:26",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "ncp-kr",
        "sys.createdTime": "2025-12-10 07:53:26",
        "sys.cspResourceId": "115562559",
        "sys.cspResourceName": "d4siboa5npi2mbdfo0p0",
        "sys.id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.uid": "d4siboa5npi2mbdfo0p0"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "223.130.152.254",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": "50",
      "rootDiskName": "",
      "connectionName": "ncp-kr",
      "connectionConfig": {
        "configName": "ncp-kr",
        "providerName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "credentialName": "ncp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncp-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ncp+kr+s4-g3a",
      "cspSpecName": "s4-g3a",
      "imageId": "23214590",
      "cspImageName": "23214590",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "130336",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "278279",
      "networkInterface": "eth0",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d4siala5npi2mbdfo0l0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-10T07:53:55Z",
          "completedTime": "2025-12-10T07:53:56Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4siboa5npi2mbdfo0p0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "115562559"
        },
        {
          "key": "ServerName",
          "value": "d4siboa5npi2mbdfo0p0"
        },
        {
          "key": "CpuCount",
          "value": "4"
        },
        {
          "key": "MemorySize",
          "value": "17179869184"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "d4siala5npi2mbdfo0l0"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2025-12-10T16:48:59+0900"
        },
        {
          "key": "Uptime",
          "value": "2025-12-10T16:51:18+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.AMD.STAND.C004.M016.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "130336"
        },
        {
          "key": "SubnetNo",
          "value": "278279"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "5122438"
        },
        {
          "key": "InitScriptNo",
          "value": "153609"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:STAND,codeName:Standard}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23214590"
        },
        {
          "key": "ServerSpecCode",
          "value": "s4-g3a"
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
        "vmIp": "223.130.152.254",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4siboa5npi2mbdfo0p0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "vmIp": "175.45.203.195",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4siboa5npi2mbdfo0q0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vmIp": "223.130.141.142",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4siboa5npi2mbdfo0o0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "output": "Linux d4siboa5npi2mbdfo0o0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "223.130.141.142",
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
      "output": "Linux d4siboa5npi2mbdfo0q0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "175.45.203.195",
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
      "output": "Linux d4siboa5npi2mbdfo0p0 5.15.0-140-generic #150-Ubuntu SMP Sat Apr 12 06:00:09 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "223.130.152.254",
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

