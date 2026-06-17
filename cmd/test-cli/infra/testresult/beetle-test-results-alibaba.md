# CM-Beetle test results for ALIBABA

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with ALIBABA cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.5.2+ (8c6611e)
- imdl: v0.1.6+ (8c6611e)
- CB-Tumblebug: v0.12.15
- CB-Spider: v0.12.30
- CB-MapUI: v0.12.39
- Target CSP: ALIBABA
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: June 17, 2026
- Test Time: 19:45:15 KST
- Test Execution: 2026-06-17 19:45:15 KST

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

## Test result for ALIBABA

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 4.549s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m8.859s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 33ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 4ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 14ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.426s | Pass |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.516s | Pass |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 27.862s | Pass |

**Overall Result**: 9/9 tests passed ✅

**Total Duration**: 2m54.30791844s

*Test executed on June 17, 2026 at 19:45:15 KST (2026-06-17 19:45:15 KST) using CM-Beetle automated test CLI*

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
        "csp": "alibaba",
        "region": "ap-northeast-2"
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
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskType": "cloud_essd_entry",
            "rootDiskSize": 40,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskType": "cloud_auto",
            "rootDiskSize": 40,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskType": "cloud_auto",
            "rootDiskSize": 40,
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
        "connectionName": "alibaba-ap-northeast-2",
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
        "connectionName": "alibaba-ap-northeast-2",
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
          "id": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "uid": "tbai41dc0rp9b7gpcck5",
          "cspSpecName": "ecs.e-c1m1.large",
          "name": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.0178,
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
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "uid": "tbbghi471c985rbqe33v",
          "cspSpecName": "ecs.e-c1m4.xlarge",
          "name": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1582,
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
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
          "uid": "tb123jkbkdc07ohb6ab2",
          "cspSpecName": "ecs.e-c1m4.large",
          "name": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.0791,
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
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
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
          "id": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "uid": "tb15cl4rjlgtva8v38op",
          "name": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.06.08 09:13:03 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
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
              "value": "Kernel version is 5.15.0-179-generic, 2026.5.29"
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
              "value": "v2026.5.29"
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
              "value": "2026-05-29T07:00:58Z"
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
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
          "description": "Kernel version is 5.15.0-179-generic, 2026.5.29",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "sg-01",
          "connectionName": "alibaba-ap-northeast-2",
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
          "connectionName": "alibaba-ap-northeast-2",
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
          "connectionName": "alibaba-ap-northeast-2",
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
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=75.0% Max=100.0% Avg=91.7% | VMs: 3 total, 0 matched, 3 acceptable",
      "targetCloud": {
        "csp": "alibaba",
        "region": "ap-northeast-2"
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
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskType": "cloud_auto",
            "rootDiskSize": 40,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskType": "cloud_auto",
            "rootDiskSize": 40,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskType": "cloud_auto",
            "rootDiskSize": 40,
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
        "connectionName": "alibaba-ap-northeast-2",
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
        "connectionName": "alibaba-ap-northeast-2",
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
          "id": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "uid": "tbai41dc0rp9b7gpcck5",
          "cspSpecName": "ecs.e-c1m1.large",
          "name": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.0178,
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
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "uid": "tbbghi471c985rbqe33v",
          "cspSpecName": "ecs.e-c1m4.xlarge",
          "name": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1582,
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
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
          "uid": "tb123jkbkdc07ohb6ab2",
          "cspSpecName": "ecs.e-c1m4.large",
          "name": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.0791,
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
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "uid": "tbc3l7erbrt3lpp0obgq",
          "cspSpecName": "ecs.t6-c1m1.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
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
          "rootDiskSize": -1,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "uid": "tbj7kinu2tgei794cq0i",
          "cspSpecName": "ecs.t6-c1m4.xlarge",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
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
          "rootDiskSize": -1,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "uid": "tb7cjk2h5m01r9d0rp38",
          "cspSpecName": "ecs.t6-c1m4.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "node",
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
          "rootDiskSize": -1,
          "systemLabel": "auto-gen",
          "details": [
            {
              "key": "CpuArchitecture",
              "value": "X86"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
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
          "id": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "uid": "tb15cl4rjlgtva8v38op",
          "name": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.06.08 09:13:03 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
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
              "value": "Kernel version is 5.15.0-179-generic, 2026.5.29"
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
              "value": "v2026.5.29"
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
              "value": "2026-05-29T07:00:58Z"
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
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
          "description": "Kernel version is 5.15.0-179-generic, 2026.5.29",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "sg-01",
          "connectionName": "alibaba-ap-northeast-2",
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
          "connectionName": "alibaba-ap-northeast-2",
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
          "connectionName": "alibaba-ap-northeast-2",
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
  "id": "my04-infra101",
  "uid": "tblj8nd2qghk41q6vomi",
  "name": "my04-infra101",
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
    "sys.id": "my04-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my04-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tblj8nd2qghk41q6vomi"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbh3rqdsuu1gqqnps914",
      "cspResourceName": "tbh3rqdsuu1gqqnps914",
      "cspResourceId": "i-mj795jd01tss2h2hgudd",
      "name": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-06-17 10:46:16",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-06-17 10:46:16",
        "sys.cspResourceId": "i-mj795jd01tss2h2hgudd",
        "sys.cspResourceName": "tbh3rqdsuu1gqqnps914",
        "sys.id": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my04-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my04-subnet-01",
        "sys.uid": "tbh3rqdsuu1gqqnps914",
        "sys.vNetId": "my04-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "8.213.136.209",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.226",
      "privateDNS": "",
      "rootDiskType": "cloud_essd_entry",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
      "cspSpecName": "ecs.e-c1m1.large",
      "spec": {
        "cspSpecName": "ecs.e-c1m1.large",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.0178
      },
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-vnet-01",
      "cspVNetId": "vpc-mj7x77ndgl64ijgqiqlga",
      "subnetId": "my04-subnet-01",
      "cspSubnetId": "vsw-mj7n9mzainfnzpl4c4x01",
      "networkInterface": "eni-mj795jd01tss2h2d7t9p",
      "securityGroupIds": [
        "my04-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-sshkey-01",
      "cspSshKeyId": "tb9olhb6136go2ho5il5",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "$vqrbi26!t1n4A",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFCfpmI/JmHVxTl+PvH8rxTDtuWJJpJXqUO3Qw91CrPcA7Hdr/P6JaGkhgQSiWbFUvztGjA57sZRWnQ9ul+AoRw=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:O3SrzlWylgU7CHdnJPBgqhYsmWFxUoYLa2KkUhm4cUs",
        "firstUsedAt": "2026-06-17T10:46:32Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-06-17T10:46:31Z",
          "completedTime": "2026-06-17T10:46:33Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj795jd01tss2h2hguddZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m1.large"
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
          "value": "tbh3rqdsuu1gqqnps914"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "200"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj795jd01tss2h2hguddZ"
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
          "value": "d8c905b7-763b-4646-8361-0b6be828d109"
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
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj795jd01tss2h2hgudd"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "KeyPairName",
          "value": "tb9olhb6136go2ho5il5"
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
          "value": "{SecurityGroupId:[sg-mj7a6pez667rl7u7fi58]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.213.136.209]}"
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
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:ON,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
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
          "value": "{VSwitchId:vsw-mj7n9mzainfnzpl4c4x01,VpcId:vpc-mj7x77ndgl64ijgqiqlga,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.226]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:11:00,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj795jd01tss2h2d7t9p,PrimaryIpAddress:10.0.1.226,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.226,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tbh88aba9dn0m6cfqna9",
      "cspResourceName": "tbh88aba9dn0m6cfqna9",
      "cspResourceId": "i-mj717btj3i68709mku4t",
      "name": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-06-17 10:46:22",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-06-17 10:46:22",
        "sys.cspResourceId": "i-mj717btj3i68709mku4t",
        "sys.cspResourceName": "tbh88aba9dn0m6cfqna9",
        "sys.id": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my04-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my04-subnet-01",
        "sys.uid": "tbh88aba9dn0m6cfqna9",
        "sys.vNetId": "my04-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "8.220.200.128",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.227",
      "privateDNS": "",
      "rootDiskType": "cloud_auto",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
      "cspSpecName": "ecs.e-c1m4.large",
      "spec": {
        "cspSpecName": "ecs.e-c1m4.large",
        "vCPU": 2,
        "memoryGiB": 8,
        "costPerHour": 0.0791
      },
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-vnet-01",
      "cspVNetId": "vpc-mj7x77ndgl64ijgqiqlga",
      "subnetId": "my04-subnet-01",
      "cspSubnetId": "vsw-mj7n9mzainfnzpl4c4x01",
      "networkInterface": "eni-mj717btj3i68709jcw1q",
      "securityGroupIds": [
        "my04-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-sshkey-01",
      "cspSshKeyId": "tb9olhb6136go2ho5il5",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "s$1!404a9Ate1b",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCTmmvDc00p2KEUqVDIWukRpv/4pXgMX2xBZ00DxAh9t3UdM5mX1s36LV8gsO6X8ZX0I6AMU9bQM/85pfwShKn4=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:c9tUt+7EmaJAGqR1CN/HROK98bSoR8LD4qCNAdCEx8w",
        "firstUsedAt": "2026-06-17T10:46:32Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-06-17T10:46:31Z",
          "completedTime": "2026-06-17T10:46:33Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj717btj3i68709mku4tZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m4.large"
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
          "value": "tbh88aba9dn0m6cfqna9"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "400"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj717btj3i68709mku4tZ"
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
          "value": "959d402a-c829-43ec-be6a-fac669797357"
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
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj717btj3i68709mku4t"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "KeyPairName",
          "value": "tb9olhb6136go2ho5il5"
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
          "value": "{SecurityGroupId:[sg-mj79ry1emh74fyd3ve1h]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.220.200.128]}"
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
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:ON,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
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
          "value": "{VSwitchId:vsw-mj7n9mzainfnzpl4c4x01,VpcId:vpc-mj7x77ndgl64ijgqiqlga,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.227]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:11:01,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj717btj3i68709jcw1q,PrimaryIpAddress:10.0.1.227,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.227,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tbid32rjajb5vjhlbbdt",
      "cspResourceName": "tbid32rjajb5vjhlbbdt",
      "cspResourceId": "i-mj779tn22erksm9zzha7",
      "name": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-06-17 10:46:15",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-06-17 10:46:15",
        "sys.cspResourceId": "i-mj779tn22erksm9zzha7",
        "sys.cspResourceName": "tbid32rjajb5vjhlbbdt",
        "sys.id": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my04-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my04-subnet-01",
        "sys.uid": "tbid32rjajb5vjhlbbdt",
        "sys.vNetId": "my04-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "47.80.241.163",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.228",
      "privateDNS": "",
      "rootDiskType": "cloud_auto",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
      "cspSpecName": "ecs.e-c1m4.xlarge",
      "spec": {
        "cspSpecName": "ecs.e-c1m4.xlarge",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1582
      },
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-vnet-01",
      "cspVNetId": "vpc-mj7x77ndgl64ijgqiqlga",
      "subnetId": "my04-subnet-01",
      "cspSubnetId": "vsw-mj7n9mzainfnzpl4c4x01",
      "networkInterface": "eni-mj779tn22erksma9bom6",
      "securityGroupIds": [
        "my04-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-sshkey-01",
      "cspSshKeyId": "tb9olhb6136go2ho5il5",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "tbAhfm5r$e!1nt",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCKkMu0c7nOYSI4zR2CeZ5fRkVFE23j2TU0BKFAHls7SkGbGVQJsIQzuHlqJooPwSDeAU+/mcfd9V2ovEU8qM18=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:meiti/YWavWpS0xXx7lK6egtB/tudpB4hf0eUpLTHDI",
        "firstUsedAt": "2026-06-17T10:46:31Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-06-17T10:46:31Z",
          "completedTime": "2026-06-17T10:46:32Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj779tn22erksm9zzha7Z 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m4.xlarge"
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
          "value": "tbid32rjajb5vjhlbbdt"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "800"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj779tn22erksm9zzha7Z"
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
          "value": "74a537fc-37bd-458d-a607-540a1730f068"
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
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj779tn22erksm9zzha7"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "KeyPairName",
          "value": "tb9olhb6136go2ho5il5"
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
          "value": "{SecurityGroupId:[sg-mj7fnyc1ybuyk2bmusev]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[47.80.241.163]}"
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
          "value": "{VSwitchId:vsw-mj7n9mzainfnzpl4c4x01,VpcId:vpc-mj7x77ndgl64ijgqiqlga,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.228]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:11:02,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj779tn22erksma9bom6,PrimaryIpAddress:10.0.1.228,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.228,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
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
        "infraId": "my04-infra101",
        "nodeId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "nodeIp": "47.80.241.163",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj779tn22erksm9zzha7Z 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my04-infra101",
        "nodeId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "8.213.136.209",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj795jd01tss2h2hguddZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my04-infra101",
        "nodeId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "8.220.200.128",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj717btj3i68709mku4tZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "id": "my01-infra101",
      "uid": "tbnl4tdm1b7hl3vdq2v7",
      "name": "my01-infra101",
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
        "sys.id": "my01-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my01-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbnl4tdm1b7hl3vdq2v7"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbs2vo5n48rk3a27qi7l",
          "cspResourceName": "tbs2vo5n48rk3a27qi7l",
          "cspResourceId": "i-0425ea303a3ec34ba",
          "name": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-06-17 10:45:58",
          "label": {
            "Name": "tbs2vo5n48rk3a27qi7l",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-06-17 10:45:58",
            "sys.cspResourceId": "i-0425ea303a3ec34ba",
            "sys.cspResourceName": "tbs2vo5n48rk3a27qi7l",
            "sys.id": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tbs2vo5n48rk3a27qi7l",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "13.124.194.20",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.171",
          "privateDNS": "ip-10-0-1-171.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 10,
          "RootDeviceName": "/dev/sda1",
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
          "spec": {
            "cspSpecName": "t3a.small",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.0234
          },
          "imageId": "ami-09a72717a566d88fa",
          "cspImageName": "ami-09a72717a566d88fa",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-09a72717a566d88fa",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260602"
          },
          "vNetId": "my01-vnet-01",
          "cspVNetId": "vpc-0fe4533820f162636",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-0d327d4ea137a0502",
          "networkInterface": "eni-attach-09e5856fa6eb773a1",
          "securityGroupIds": [
            "my01-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbojleiabbl4rakk9qhs",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFEI9z8f2p3AamCRz/dZ7expmaS9bilNCBqtFGYGaD+9OXR65v9/phpjBATFRQKjPd6am9NR5GIVFN3lIoRMhQY=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:Uocd7pMAzhGZCDuN5ettpog/FB4rWSbIWO9XG4Ftoe0",
            "firstUsedAt": "2026-06-17T10:46:07Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-17T10:46:05Z",
              "completedTime": "2026-06-17T10:46:08Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-171 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-06-17T10:45:38Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-01189266ec7e16e28}}"
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
              "value": "3F6DBBE4-FF3B-43AE-9CD2-6EB349F65C43"
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
              "value": "ami-09a72717a566d88fa"
            },
            {
              "key": "InstanceId",
              "value": "i-0425ea303a3ec34ba"
            },
            {
              "key": "InstanceType",
              "value": "t3a.small"
            },
            {
              "key": "KeyName",
              "value": "tbojleiabbl4rakk9qhs"
            },
            {
              "key": "LaunchTime",
              "value": "2026-06-17T10:45:37Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.124.194.20},Attachment:{AttachTime:2026-06-17T10:45:37Z,AttachmentId:eni-attach-09e5856fa6eb773a1,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0453baa45805db6fb,GroupName:tboo9tl54g7njkr0j3un}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:46:34:be:23:71,NetworkInterfaceId:eni-02fd5f227eabf04a9,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.171,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.124.194.20},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.171}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0d327d4ea137a0502,VpcId:vpc-0fe4533820f162636}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-171.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.171"
            },
            {
              "key": "PublicIpAddress",
              "value": "13.124.194.20"
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
              "value": "{GroupId:sg-0453baa45805db6fb,GroupName:tboo9tl54g7njkr0j3un}"
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
              "value": "subnet-0d327d4ea137a0502"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbs2vo5n48rk3a27qi7l}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0fe4533820f162636"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tb3rbnn8r9hum5f8q2hk",
          "cspResourceName": "tb3rbnn8r9hum5f8q2hk",
          "cspResourceId": "i-00a515dadf434b01a",
          "name": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
          "createdTime": "2026-06-17 10:45:58",
          "label": {
            "Name": "tb3rbnn8r9hum5f8q2hk",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-06-17 10:45:58",
            "sys.cspResourceId": "i-00a515dadf434b01a",
            "sys.cspResourceName": "tb3rbnn8r9hum5f8q2hk",
            "sys.id": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tb3rbnn8r9hum5f8q2hk",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "13.125.176.19",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.10",
          "privateDNS": "ip-10-0-1-10.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 10,
          "RootDeviceName": "/dev/sda1",
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
          "spec": {
            "cspSpecName": "t3a.large",
            "vCPU": 2,
            "memoryGiB": 8,
            "costPerHour": 0.0936
          },
          "imageId": "ami-09a72717a566d88fa",
          "cspImageName": "ami-09a72717a566d88fa",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-09a72717a566d88fa",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260602"
          },
          "vNetId": "my01-vnet-01",
          "cspVNetId": "vpc-0fe4533820f162636",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-0d327d4ea137a0502",
          "networkInterface": "eni-attach-0c9ec535ca7060b15",
          "securityGroupIds": [
            "my01-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbojleiabbl4rakk9qhs",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJzJv2R9A1MTg0deRNeJpYVGaKuWEk/lXm/+PEm5uMdNgg7fjHUvVyQO1tCtaQaGD3Pz8TGTBEZjIgTZFawbsTc=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:gLwmdHn9yOEV5lNknMylvrTAY1zzgthZulFjZUGsPoI",
            "firstUsedAt": "2026-06-17T10:46:05Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-17T10:46:05Z",
              "completedTime": "2026-06-17T10:46:08Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-10 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-06-17T10:45:36Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-034ef444a2c04dead}}"
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
              "value": "7D98FA17-E7F9-4D55-8D83-4D647D751B0A"
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
              "value": "ami-09a72717a566d88fa"
            },
            {
              "key": "InstanceId",
              "value": "i-00a515dadf434b01a"
            },
            {
              "key": "InstanceType",
              "value": "t3a.large"
            },
            {
              "key": "KeyName",
              "value": "tbojleiabbl4rakk9qhs"
            },
            {
              "key": "LaunchTime",
              "value": "2026-06-17T10:45:35Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.125.176.19},Attachment:{AttachTime:2026-06-17T10:45:35Z,AttachmentId:eni-attach-0c9ec535ca7060b15,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-06c3ad0e17b06483b,GroupName:tb85r2mi88lep4nc2t9i}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:11:39:ec:34:d9,NetworkInterfaceId:eni-094543d4ffb54d886,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.10,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.125.176.19},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.10}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0d327d4ea137a0502,VpcId:vpc-0fe4533820f162636}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-10.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.10"
            },
            {
              "key": "PublicIpAddress",
              "value": "13.125.176.19"
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
              "value": "{GroupId:sg-06c3ad0e17b06483b,GroupName:tb85r2mi88lep4nc2t9i}"
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
              "value": "subnet-0d327d4ea137a0502"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tb3rbnn8r9hum5f8q2hk}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0fe4533820f162636"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbi6nauamr7o5nkpahga",
          "cspResourceName": "tbi6nauamr7o5nkpahga",
          "cspResourceId": "i-0e8b6794681b19e56",
          "name": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
          "createdTime": "2026-06-17 10:46:00",
          "label": {
            "Name": "tbi6nauamr7o5nkpahga",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-06-17 10:46:00",
            "sys.cspResourceId": "i-0e8b6794681b19e56",
            "sys.cspResourceName": "tbi6nauamr7o5nkpahga",
            "sys.id": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tbi6nauamr7o5nkpahga",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "52.78.66.65",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.208",
          "privateDNS": "ip-10-0-1-208.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 10,
          "RootDeviceName": "/dev/sda1",
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
          "spec": {
            "cspSpecName": "t3a.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1872
          },
          "imageId": "ami-09a72717a566d88fa",
          "cspImageName": "ami-09a72717a566d88fa",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-09a72717a566d88fa",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260602"
          },
          "vNetId": "my01-vnet-01",
          "cspVNetId": "vpc-0fe4533820f162636",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-0d327d4ea137a0502",
          "networkInterface": "eni-attach-00cc39f503414e726",
          "securityGroupIds": [
            "my01-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbojleiabbl4rakk9qhs",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBm9pNxT0pMhyN9FtLQf43CVFgn9T9lsndQobYmjtlAIGFVxnzu+RYAWa/o9Vr69UqR4jwsPc8LUbJQhBb36YFs=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:ZpLGIKD2/3JH0f0Mj/PkbA9p9HlQxQlpOzZBCtLTDPg",
            "firstUsedAt": "2026-06-17T10:46:07Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-17T10:46:05Z",
              "completedTime": "2026-06-17T10:46:08Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-208 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-06-17T10:45:37Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-016b9a5f4bd9dd00e}}"
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
              "value": "06AD8265-2901-4A94-B2C2-DF161A337C69"
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
              "value": "ami-09a72717a566d88fa"
            },
            {
              "key": "InstanceId",
              "value": "i-0e8b6794681b19e56"
            },
            {
              "key": "InstanceType",
              "value": "t3a.xlarge"
            },
            {
              "key": "KeyName",
              "value": "tbojleiabbl4rakk9qhs"
            },
            {
              "key": "LaunchTime",
              "value": "2026-06-17T10:45:37Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:52.78.66.65},Attachment:{AttachTime:2026-06-17T10:45:37Z,AttachmentId:eni-attach-00cc39f503414e726,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-033d323238fcb3132,GroupName:tbjc7r8u2iq57pmuouud}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:fc:ba:f3:19:73,NetworkInterfaceId:eni-028c65e6e6889dc45,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.208,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:52.78.66.65},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.208}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0d327d4ea137a0502,VpcId:vpc-0fe4533820f162636}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-208.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.208"
            },
            {
              "key": "PublicIpAddress",
              "value": "52.78.66.65"
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
              "value": "{GroupId:sg-033d323238fcb3132,GroupName:tbjc7r8u2iq57pmuouud}"
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
              "value": "subnet-0d327d4ea137a0502"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbi6nauamr7o5nkpahga}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0fe4533820f162636"
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
            "infraId": "my01-infra101",
            "nodeId": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "13.125.176.19",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-10 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "52.78.66.65",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-208 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "13.124.194.20",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-171 6.8.0-1057-aws #60~22.04.1-Ubuntu SMP Wed May 27 08:16:59 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "id": "my02-infra101",
      "uid": "tb5rdln409vo6rm7e7ng",
      "name": "my02-infra101",
      "status": "Creating:3 (R:0/3)",
      "statusCount": {
        "countTotal": 3,
        "countCreating": 3,
        "countRunning": 0,
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
      "targetStatus": "Running",
      "targetAction": "Create",
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
        "sys.id": "my02-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tb5rdln409vo6rm7e7ng"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbrc1c5ene0o6mcb6c64",
          "name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": null,
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
          "region": {
            "region": ""
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
          "rootDiskSize": 30,
          "RootDeviceName": "",
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
          "cspSpecName": "",
          "spec": {},
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202606070",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202606110",
          "image": {
            "osType": ""
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my02-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tbj82pstbuudh0ekclti",
          "name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": null,
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": ""
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
          "rootDiskSize": 30,
          "RootDeviceName": "",
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
          "cspSpecName": "",
          "spec": {},
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606070",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202606110",
          "image": {
            "osType": ""
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my02-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbft43rqs51uj0konsm8",
          "name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": null,
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": ""
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
          "rootDiskSize": 30,
          "RootDeviceName": "",
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
          "cspSpecName": "",
          "spec": {},
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202606070",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202606110",
          "image": {
            "osType": ""
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my02-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": ""
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
        "results": null
      }
    },
    {
      "resourceType": "infra",
      "id": "my04-infra101",
      "uid": "tblj8nd2qghk41q6vomi",
      "name": "my04-infra101",
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
        "sys.id": "my04-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my04-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tblj8nd2qghk41q6vomi"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbh3rqdsuu1gqqnps914",
          "cspResourceName": "tbh3rqdsuu1gqqnps914",
          "cspResourceId": "i-mj795jd01tss2h2hgudd",
          "name": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-06-17 10:46:16",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-06-17 10:46:16",
            "sys.cspResourceId": "i-mj795jd01tss2h2hgudd",
            "sys.cspResourceName": "tbh3rqdsuu1gqqnps914",
            "sys.id": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my04-subnet-01",
            "sys.uid": "tbh3rqdsuu1gqqnps914",
            "sys.vNetId": "my04-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.213.136.209",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.226",
          "privateDNS": "",
          "rootDiskType": "cloud_essd_entry",
          "rootDiskSize": 40,
          "RootDeviceName": "/dev/xvda",
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
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "cspSpecName": "ecs.e-c1m1.large",
          "spec": {
            "cspSpecName": "ecs.e-c1m1.large",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.0178
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-vnet-01",
          "cspVNetId": "vpc-mj7x77ndgl64ijgqiqlga",
          "subnetId": "my04-subnet-01",
          "cspSubnetId": "vsw-mj7n9mzainfnzpl4c4x01",
          "networkInterface": "eni-mj795jd01tss2h2d7t9p",
          "securityGroupIds": [
            "my04-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-sshkey-01",
          "cspSshKeyId": "tb9olhb6136go2ho5il5",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "$vqrbi26!t1n4A",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFCfpmI/JmHVxTl+PvH8rxTDtuWJJpJXqUO3Qw91CrPcA7Hdr/P6JaGkhgQSiWbFUvztGjA57sZRWnQ9ul+AoRw=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:O3SrzlWylgU7CHdnJPBgqhYsmWFxUoYLa2KkUhm4cUs",
            "firstUsedAt": "2026-06-17T10:46:32Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-17T10:46:31Z",
              "completedTime": "2026-06-17T10:46:33Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj795jd01tss2h2hguddZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
            },
            {
              "key": "InstanceType",
              "value": "ecs.e-c1m1.large"
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
              "value": "tbh3rqdsuu1gqqnps914"
            },
            {
              "key": "DeploymentSetGroupNo",
              "value": "0"
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
              "value": "2026-06-17T10:45Z"
            },
            {
              "key": "ZoneId",
              "value": "ap-northeast-2a"
            },
            {
              "key": "InternetMaxBandwidthIn",
              "value": "200"
            },
            {
              "key": "InternetChargeType",
              "value": "PayByBandwidth"
            },
            {
              "key": "HostName",
              "value": "iZmj795jd01tss2h2hguddZ"
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
              "value": "d8c905b7-763b-4646-8361-0b6be828d109"
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
              "value": "ecs.e"
            },
            {
              "key": "InstanceId",
              "value": "i-mj795jd01tss2h2hgudd"
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
              "value": "2026-06-17T10:45Z"
            },
            {
              "key": "KeyPairName",
              "value": "tb9olhb6136go2ho5il5"
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
              "value": "{SecurityGroupId:[sg-mj7a6pez667rl7u7fi58]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[8.213.136.209]}"
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
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:ON,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
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
              "value": "{VSwitchId:vsw-mj7n9mzainfnzpl4c4x01,VpcId:vpc-mj7x77ndgl64ijgqiqlga,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.226]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:11:00,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj795jd01tss2h2d7t9p,PrimaryIpAddress:10.0.1.226,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.226,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
            },
            {
              "key": "OperationLocks",
              "value": "{LockReason:[]}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tbh88aba9dn0m6cfqna9",
          "cspResourceName": "tbh88aba9dn0m6cfqna9",
          "cspResourceId": "i-mj717btj3i68709mku4t",
          "name": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
          "createdTime": "2026-06-17 10:46:22",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-06-17 10:46:22",
            "sys.cspResourceId": "i-mj717btj3i68709mku4t",
            "sys.cspResourceName": "tbh88aba9dn0m6cfqna9",
            "sys.id": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my04-subnet-01",
            "sys.uid": "tbh88aba9dn0m6cfqna9",
            "sys.vNetId": "my04-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.220.200.128",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.227",
          "privateDNS": "",
          "rootDiskType": "cloud_auto",
          "rootDiskSize": 40,
          "RootDeviceName": "/dev/xvda",
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
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
          "cspSpecName": "ecs.e-c1m4.large",
          "spec": {
            "cspSpecName": "ecs.e-c1m4.large",
            "vCPU": 2,
            "memoryGiB": 8,
            "costPerHour": 0.0791
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-vnet-01",
          "cspVNetId": "vpc-mj7x77ndgl64ijgqiqlga",
          "subnetId": "my04-subnet-01",
          "cspSubnetId": "vsw-mj7n9mzainfnzpl4c4x01",
          "networkInterface": "eni-mj717btj3i68709jcw1q",
          "securityGroupIds": [
            "my04-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-sshkey-01",
          "cspSshKeyId": "tb9olhb6136go2ho5il5",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "s$1!404a9Ate1b",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCTmmvDc00p2KEUqVDIWukRpv/4pXgMX2xBZ00DxAh9t3UdM5mX1s36LV8gsO6X8ZX0I6AMU9bQM/85pfwShKn4=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:c9tUt+7EmaJAGqR1CN/HROK98bSoR8LD4qCNAdCEx8w",
            "firstUsedAt": "2026-06-17T10:46:32Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-17T10:46:31Z",
              "completedTime": "2026-06-17T10:46:33Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj717btj3i68709mku4tZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
            },
            {
              "key": "InstanceType",
              "value": "ecs.e-c1m4.large"
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
              "value": "tbh88aba9dn0m6cfqna9"
            },
            {
              "key": "DeploymentSetGroupNo",
              "value": "0"
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
              "value": "2026-06-17T10:45Z"
            },
            {
              "key": "ZoneId",
              "value": "ap-northeast-2a"
            },
            {
              "key": "InternetMaxBandwidthIn",
              "value": "400"
            },
            {
              "key": "InternetChargeType",
              "value": "PayByBandwidth"
            },
            {
              "key": "HostName",
              "value": "iZmj717btj3i68709mku4tZ"
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
              "value": "959d402a-c829-43ec-be6a-fac669797357"
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
              "value": "ecs.e"
            },
            {
              "key": "InstanceId",
              "value": "i-mj717btj3i68709mku4t"
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
              "value": "2026-06-17T10:45Z"
            },
            {
              "key": "KeyPairName",
              "value": "tb9olhb6136go2ho5il5"
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
              "value": "{SecurityGroupId:[sg-mj79ry1emh74fyd3ve1h]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[8.220.200.128]}"
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
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:ON,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
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
              "value": "{VSwitchId:vsw-mj7n9mzainfnzpl4c4x01,VpcId:vpc-mj7x77ndgl64ijgqiqlga,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.227]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:11:01,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj717btj3i68709jcw1q,PrimaryIpAddress:10.0.1.227,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.227,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
            },
            {
              "key": "OperationLocks",
              "value": "{LockReason:[]}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbid32rjajb5vjhlbbdt",
          "cspResourceName": "tbid32rjajb5vjhlbbdt",
          "cspResourceId": "i-mj779tn22erksm9zzha7",
          "name": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
          "createdTime": "2026-06-17 10:46:15",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-06-17 10:46:15",
            "sys.cspResourceId": "i-mj779tn22erksm9zzha7",
            "sys.cspResourceName": "tbid32rjajb5vjhlbbdt",
            "sys.id": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my04-subnet-01",
            "sys.uid": "tbid32rjajb5vjhlbbdt",
            "sys.vNetId": "my04-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.241.163",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.228",
          "privateDNS": "",
          "rootDiskType": "cloud_auto",
          "rootDiskSize": 40,
          "RootDeviceName": "/dev/xvda",
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
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "cspSpecName": "ecs.e-c1m4.xlarge",
          "spec": {
            "cspSpecName": "ecs.e-c1m4.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1582
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-vnet-01",
          "cspVNetId": "vpc-mj7x77ndgl64ijgqiqlga",
          "subnetId": "my04-subnet-01",
          "cspSubnetId": "vsw-mj7n9mzainfnzpl4c4x01",
          "networkInterface": "eni-mj779tn22erksma9bom6",
          "securityGroupIds": [
            "my04-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-sshkey-01",
          "cspSshKeyId": "tb9olhb6136go2ho5il5",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "tbAhfm5r$e!1nt",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCKkMu0c7nOYSI4zR2CeZ5fRkVFE23j2TU0BKFAHls7SkGbGVQJsIQzuHlqJooPwSDeAU+/mcfd9V2ovEU8qM18=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:meiti/YWavWpS0xXx7lK6egtB/tudpB4hf0eUpLTHDI",
            "firstUsedAt": "2026-06-17T10:46:31Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-17T10:46:31Z",
              "completedTime": "2026-06-17T10:46:32Z",
              "elapsedTime": 1,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj779tn22erksm9zzha7Z 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
            },
            {
              "key": "InstanceType",
              "value": "ecs.e-c1m4.xlarge"
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
              "value": "tbid32rjajb5vjhlbbdt"
            },
            {
              "key": "DeploymentSetGroupNo",
              "value": "0"
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
              "value": "2026-06-17T10:45Z"
            },
            {
              "key": "ZoneId",
              "value": "ap-northeast-2a"
            },
            {
              "key": "InternetMaxBandwidthIn",
              "value": "800"
            },
            {
              "key": "InternetChargeType",
              "value": "PayByBandwidth"
            },
            {
              "key": "HostName",
              "value": "iZmj779tn22erksm9zzha7Z"
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
              "value": "74a537fc-37bd-458d-a607-540a1730f068"
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
              "value": "ecs.e"
            },
            {
              "key": "InstanceId",
              "value": "i-mj779tn22erksm9zzha7"
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
              "value": "2026-06-17T10:45Z"
            },
            {
              "key": "KeyPairName",
              "value": "tb9olhb6136go2ho5il5"
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
              "value": "{SecurityGroupId:[sg-mj7fnyc1ybuyk2bmusev]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[47.80.241.163]}"
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
              "value": "{VSwitchId:vsw-mj7n9mzainfnzpl4c4x01,VpcId:vpc-mj7x77ndgl64ijgqiqlga,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.228]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:11:02,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj779tn22erksma9bom6,PrimaryIpAddress:10.0.1.228,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.228,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
            },
            {
              "key": "OperationLocks",
              "value": "{LockReason:[]}"
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
            "infraId": "my04-infra101",
            "nodeId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "47.80.241.163",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj779tn22erksm9zzha7Z 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "8.213.136.209",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj795jd01tss2h2hguddZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "8.220.200.128",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj717btj3i68709mku4tZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
    "my01-infra101",
    "my02-infra101",
    "my04-infra101"
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
  "id": "my04-infra101",
  "uid": "tblj8nd2qghk41q6vomi",
  "name": "my04-infra101",
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
    "sys.id": "my04-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my04-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tblj8nd2qghk41q6vomi"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbh3rqdsuu1gqqnps914",
      "cspResourceName": "tbh3rqdsuu1gqqnps914",
      "cspResourceId": "i-mj795jd01tss2h2hgudd",
      "name": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-06-17 10:46:16",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-06-17 10:46:16",
        "sys.cspResourceId": "i-mj795jd01tss2h2hgudd",
        "sys.cspResourceName": "tbh3rqdsuu1gqqnps914",
        "sys.id": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my04-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my04-subnet-01",
        "sys.uid": "tbh3rqdsuu1gqqnps914",
        "sys.vNetId": "my04-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "8.213.136.209",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.226",
      "privateDNS": "",
      "rootDiskType": "cloud_essd_entry",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
      "cspSpecName": "ecs.e-c1m1.large",
      "spec": {
        "cspSpecName": "ecs.e-c1m1.large",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.0178
      },
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-vnet-01",
      "cspVNetId": "vpc-mj7x77ndgl64ijgqiqlga",
      "subnetId": "my04-subnet-01",
      "cspSubnetId": "vsw-mj7n9mzainfnzpl4c4x01",
      "networkInterface": "eni-mj795jd01tss2h2d7t9p",
      "securityGroupIds": [
        "my04-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-sshkey-01",
      "cspSshKeyId": "tb9olhb6136go2ho5il5",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "$vqrbi26!t1n4A",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFCfpmI/JmHVxTl+PvH8rxTDtuWJJpJXqUO3Qw91CrPcA7Hdr/P6JaGkhgQSiWbFUvztGjA57sZRWnQ9ul+AoRw=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:O3SrzlWylgU7CHdnJPBgqhYsmWFxUoYLa2KkUhm4cUs",
        "firstUsedAt": "2026-06-17T10:46:32Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-06-17T10:46:31Z",
          "completedTime": "2026-06-17T10:46:33Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj795jd01tss2h2hguddZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m1.large"
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
          "value": "tbh3rqdsuu1gqqnps914"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "200"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj795jd01tss2h2hguddZ"
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
          "value": "d8c905b7-763b-4646-8361-0b6be828d109"
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
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj795jd01tss2h2hgudd"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "KeyPairName",
          "value": "tb9olhb6136go2ho5il5"
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
          "value": "{SecurityGroupId:[sg-mj7a6pez667rl7u7fi58]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.213.136.209]}"
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
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:ON,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
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
          "value": "{VSwitchId:vsw-mj7n9mzainfnzpl4c4x01,VpcId:vpc-mj7x77ndgl64ijgqiqlga,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.226]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:11:00,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj795jd01tss2h2d7t9p,PrimaryIpAddress:10.0.1.226,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.226,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tbh88aba9dn0m6cfqna9",
      "cspResourceName": "tbh88aba9dn0m6cfqna9",
      "cspResourceId": "i-mj717btj3i68709mku4t",
      "name": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-06-17 10:46:22",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-06-17 10:46:22",
        "sys.cspResourceId": "i-mj717btj3i68709mku4t",
        "sys.cspResourceName": "tbh88aba9dn0m6cfqna9",
        "sys.id": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my04-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my04-subnet-01",
        "sys.uid": "tbh88aba9dn0m6cfqna9",
        "sys.vNetId": "my04-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "8.220.200.128",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.227",
      "privateDNS": "",
      "rootDiskType": "cloud_auto",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
      "cspSpecName": "ecs.e-c1m4.large",
      "spec": {
        "cspSpecName": "ecs.e-c1m4.large",
        "vCPU": 2,
        "memoryGiB": 8,
        "costPerHour": 0.0791
      },
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-vnet-01",
      "cspVNetId": "vpc-mj7x77ndgl64ijgqiqlga",
      "subnetId": "my04-subnet-01",
      "cspSubnetId": "vsw-mj7n9mzainfnzpl4c4x01",
      "networkInterface": "eni-mj717btj3i68709jcw1q",
      "securityGroupIds": [
        "my04-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-sshkey-01",
      "cspSshKeyId": "tb9olhb6136go2ho5il5",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "s$1!404a9Ate1b",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCTmmvDc00p2KEUqVDIWukRpv/4pXgMX2xBZ00DxAh9t3UdM5mX1s36LV8gsO6X8ZX0I6AMU9bQM/85pfwShKn4=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:c9tUt+7EmaJAGqR1CN/HROK98bSoR8LD4qCNAdCEx8w",
        "firstUsedAt": "2026-06-17T10:46:32Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-06-17T10:46:31Z",
          "completedTime": "2026-06-17T10:46:33Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj717btj3i68709mku4tZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m4.large"
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
          "value": "tbh88aba9dn0m6cfqna9"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "400"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj717btj3i68709mku4tZ"
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
          "value": "959d402a-c829-43ec-be6a-fac669797357"
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
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj717btj3i68709mku4t"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "KeyPairName",
          "value": "tb9olhb6136go2ho5il5"
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
          "value": "{SecurityGroupId:[sg-mj79ry1emh74fyd3ve1h]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.220.200.128]}"
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
          "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:1,CoreFactor:0,Numa:ON,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
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
          "value": "{VSwitchId:vsw-mj7n9mzainfnzpl4c4x01,VpcId:vpc-mj7x77ndgl64ijgqiqlga,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.227]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:11:01,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj717btj3i68709jcw1q,PrimaryIpAddress:10.0.1.227,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.227,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tbid32rjajb5vjhlbbdt",
      "cspResourceName": "tbid32rjajb5vjhlbbdt",
      "cspResourceId": "i-mj779tn22erksm9zzha7",
      "name": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-06-17 10:46:15",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2026-06-17 10:46:15",
        "sys.cspResourceId": "i-mj779tn22erksm9zzha7",
        "sys.cspResourceName": "tbid32rjajb5vjhlbbdt",
        "sys.id": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my04-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my04-subnet-01",
        "sys.uid": "tbid32rjajb5vjhlbbdt",
        "sys.vNetId": "my04-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "47.80.241.163",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.228",
      "privateDNS": "",
      "rootDiskType": "cloud_auto",
      "rootDiskSize": 40,
      "RootDeviceName": "/dev/xvda",
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
      "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
      "cspSpecName": "ecs.e-c1m4.xlarge",
      "spec": {
        "cspSpecName": "ecs.e-c1m4.xlarge",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.1582
      },
      "imageId": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
      "image": {
        "resourceType": "image",
        "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260522.vhd",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu  22.04 64 bit"
      },
      "vNetId": "my04-vnet-01",
      "cspVNetId": "vpc-mj7x77ndgl64ijgqiqlga",
      "subnetId": "my04-subnet-01",
      "cspSubnetId": "vsw-mj7n9mzainfnzpl4c4x01",
      "networkInterface": "eni-mj779tn22erksma9bom6",
      "securityGroupIds": [
        "my04-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my04-sshkey-01",
      "cspSshKeyId": "tb9olhb6136go2ho5il5",
      "nodeUserName": "cb-user",
      "nodeUserPassword": "tbAhfm5r$e!1nt",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBCKkMu0c7nOYSI4zR2CeZ5fRkVFE23j2TU0BKFAHls7SkGbGVQJsIQzuHlqJooPwSDeAU+/mcfd9V2ovEU8qM18=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:meiti/YWavWpS0xXx7lK6egtB/tudpB4hf0eUpLTHDI",
        "firstUsedAt": "2026-06-17T10:46:31Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-06-17T10:46:31Z",
          "completedTime": "2026-06-17T10:46:32Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux iZmj779tn22erksm9zzha7Z 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20260522.vhd"
        },
        {
          "key": "InstanceType",
          "value": "ecs.e-c1m4.xlarge"
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
          "value": "tbid32rjajb5vjhlbbdt"
        },
        {
          "key": "DeploymentSetGroupNo",
          "value": "0"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "InternetMaxBandwidthIn",
          "value": "800"
        },
        {
          "key": "InternetChargeType",
          "value": "PayByBandwidth"
        },
        {
          "key": "HostName",
          "value": "iZmj779tn22erksm9zzha7Z"
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
          "value": "74a537fc-37bd-458d-a607-540a1730f068"
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
          "value": "ecs.e"
        },
        {
          "key": "InstanceId",
          "value": "i-mj779tn22erksm9zzha7"
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
          "value": "2026-06-17T10:45Z"
        },
        {
          "key": "KeyPairName",
          "value": "tb9olhb6136go2ho5il5"
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
          "value": "{SecurityGroupId:[sg-mj7fnyc1ybuyk2bmusev]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[47.80.241.163]}"
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
          "value": "{VSwitchId:vsw-mj7n9mzainfnzpl4c4x01,VpcId:vpc-mj7x77ndgl64ijgqiqlga,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.228]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:11:02,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj779tn22erksma9bom6,PrimaryIpAddress:10.0.1.228,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.228,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
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
        "infraId": "my04-infra101",
        "nodeId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "nodeIp": "47.80.241.163",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj779tn22erksm9zzha7Z 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my04-infra101",
        "nodeId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "8.213.136.209",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj795jd01tss2h2hguddZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my04-infra101",
        "nodeId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "8.220.200.128",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux iZmj717btj3i68709mku4tZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "nodeGroup": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "nodeId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "output": "Linux iZmj795jd01tss2h2hguddZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "8.213.136.209",
      "sshTest": "successful",
      "status": "success",
      "testOrder": 1,
      "userName": "cb-user"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "nodeGroup": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "nodeId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "output": "Linux iZmj717btj3i68709mku4tZ 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "8.220.200.128",
      "sshTest": "successful",
      "status": "success",
      "testOrder": 2,
      "userName": "cb-user"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "nodeGroup": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "nodeId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "output": "Linux iZmj779tn22erksm9zzha7Z 5.15.0-179-generic #189-Ubuntu SMP Tue May 5 18:20:56 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "47.80.241.163",
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

**Generated At:** 2026-06-17 10:46:58

**Namespace:** mig01

**Infra Name:** my04-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my04-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | ALIBABA |
| **Target Region** | ap-northeast-2 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| ecs.e-c1m4.large | 2 | 8.0 | - | x86_64 |  | $0.0791 | 1 |
| ecs.e-c1m4.xlarge | 4 | 16.0 | - | x86_64 |  | $0.1582 | 1 |
| ecs.e-c1m1.large | 2 | 2.0 | - | x86_64 |  | $0.0178 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| ubuntu_22_04_x64_20G_alibase_20260522.vhd | Ubuntu  22.04 64 bit | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 20 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | i-mj795jd01tss2h2hgudd | Running | 2 vCPU, 2.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-vnet-01<br>**Subnet:** my04-subnet-01<br>**Public IP:** 8.213.136.209<br>**Private IP:** 10.0.1.226<br>**SGs:** my04-sg-01<br>**SSH:** my04-sshkey-01 |
| my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | i-mj717btj3i68709mku4t | Running | 2 vCPU, 8.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-vnet-01<br>**Subnet:** my04-subnet-01<br>**Public IP:** 8.220.200.128<br>**Private IP:** 10.0.1.227<br>**SGs:** my04-sg-03<br>**SSH:** my04-sshkey-01 |
| my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | i-mj779tn22erksm9zzha7 | Running | 4 vCPU, 16.0 GiB | Ubuntu  22.04 64 bit (Ubuntu  22.04 64 bit) | **VNet:** my04-vnet-01<br>**Subnet:** my04-subnet-01<br>**Public IP:** 47.80.241.163<br>**Private IP:** 10.0.1.228<br>**SGs:** my04-sg-02<br>**SSH:** my04-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my04-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my04-vnet-01 |
| **CSP VNet ID** | vpc-mj7x77ndgl64ijgqiqlga |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | alibaba-ap-northeast-2 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my04-subnet-01 | vsw-mj7n9mzainfnzpl4c4x01 | 10.0.1.0/24 | ap-northeast-2a |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my04-sshkey-01 | tb9olhb6136go2ho5il5 |  | da1323e457ca703332e8f0ea2530007f |

### Security Groups

#### Security Group: my04-sg-01

| Property | Value |
|----------|-------|
| **Name** | my04-sg-01 |
| **CSP Security Group ID** | sg-mj7a6pez667rl7u7fi58 |
| **VNet** | my04-vnet-01 |
| **Rule Count** | 14 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | UDP | 9113 | 10.0.0.0/16 |
| inbound | TCP | 9113 | 10.0.0.0/16 |
| inbound | TCP | 8080 | 0.0.0.0/0 |
| inbound | TCP | 443 | 0.0.0.0/0 |
| inbound | TCP | 80 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my04-sg-02

| Property | Value |
|----------|-------|
| **Name** | my04-sg-02 |
| **CSP Security Group ID** | sg-mj7fnyc1ybuyk2bmusev |
| **VNet** | my04-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | UDP | 9100 | 10.0.0.0/16 |
| inbound | TCP | 9100 | 10.0.0.0/16 |
| inbound | UDP | 32803 | 10.0.0.0/16 |
| inbound | TCP | 32803 | 10.0.0.0/16 |
| inbound | UDP | 20048 | 10.0.0.0/16 |
| inbound | TCP | 20048 | 10.0.0.0/16 |
| inbound | UDP | 111 | 0.0.0.0/0 |
| inbound | TCP | 111 | 0.0.0.0/0 |
| inbound | UDP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 2049 | 0.0.0.0/0 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my04-sg-03

| Property | Value |
|----------|-------|
| **Name** | my04-sg-03 |
| **CSP Security Group ID** | sg-mj79ry1emh74fyd3ve1h |
| **VNet** | my04-vnet-01 |
| **Rule Count** | 19 rules |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR |
|-----------|----------|------------|------|
| inbound | ALL |  | 10.0.0.0/16 |
| inbound | UDP | 9104 | 10.0.0.0/16 |
| inbound | TCP | 9104 | 10.0.0.0/16 |
| inbound | UDP | 4444 | 10.0.0.0/16 |
| inbound | TCP | 4444 | 10.0.0.0/16 |
| inbound | UDP | 4568 | 10.0.0.0/16 |
| inbound | TCP | 4568 | 10.0.0.0/16 |
| inbound | UDP | 4567 | 10.0.0.0/16 |
| inbound | TCP | 4567 | 10.0.0.0/16 |
| inbound | UDP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 3306 | 10.0.0.0/16 |
| inbound | TCP | 22 | 0.0.0.0/0 |
| inbound | UDP | 1900 | 0.0.0.0/0 |
| inbound | UDP | 5353 | 0.0.0.0/0 |
| inbound | UDP | 68 | 0.0.0.0/0 |
| inbound | ICMP |  | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.2551 |
| **Per Day** | $6.12 |
| **Per Month (30 days)** | $183.67 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| ALIBABA | ap-northeast-2 | 3 | $0.2551 | $183.67 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | ecs.e-c1m1.large | $0.0178 | $12.82 |
| my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | ecs.e-c1m4.large | $0.0791 | $56.95 |
| my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | ecs.e-c1m4.xlarge | $0.1582 | $113.90 |




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

*Report generated: 2026-06-17 10:47:04*

---

## 📊 Migration Summary

**Target Cloud:** ALIBABA

**Target Region:** ap-northeast-2

**Namespace:** mig01 | **Infra ID:** my04-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $183.67 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 3 | ✅ Selected | ecs.e-c1m1.large, ecs.e-c1m4.large, ecs.e-c1m4.xlarge |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my04-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my04-vnet-01) |
| 6 | **Security Group** | 3 security groups | ✅ Created | Total 52 rules in 3 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** i-mj795jd01tss2h2hgudd<br>**Label(sourceMachineId):** vm-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** i-mj717btj3i68709mku4t<br>**Label(sourceMachineId):** vm-ec288dd0-c6fa-8a49-2f60 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 |
| 3 | **VM Name:** my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** i-mj779tn22erksm9zzha7<br>**Label(sourceMachineId):** vm-ec2d32b5-98fb-5a96-7913 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 |

---

## ⚙️ VM Specs

**Summary:** 3 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** ecs.e-c1m1.large<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2 | my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Spec ID:** ecs.e-c1m4.large<br>**vCPUs:** 2<br>**Memory:** 8.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3 | my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Spec ID:** ecs.e-c1m4.xlarge<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260522.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260522.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260522.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu  22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my04-sg-01

**CSP ID:** sg-mj7a6pez667rl7u7fi58 | **VNet:** my04-vnet-01 | **Rules:** 14

**Assigned VMs:**

- **VM:** my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec268ed7-821e-9d73-e79f

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 2 | inbound | UDP | 9113 | 10.0.0.0/16 | inbound udp 9113 from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | TCP | 9113 | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | TCP | 8080 | 0.0.0.0/0 | inbound tcp 8080 | Migrated from source |
| 5 | inbound | TCP | 443 | 0.0.0.0/0 | inbound tcp 443 | Migrated from source |
| 6 | inbound | TCP | 80 | 0.0.0.0/0 | inbound tcp 80 | Migrated from source |
| 7 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 8 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 9 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 10 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 11 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 12 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 13 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 14 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my04-sg-02

**CSP ID:** sg-mj7fnyc1ybuyk2bmusev | **VNet:** my04-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec2d32b5-98fb-5a96-7913

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 2 | inbound | UDP | 9100 | 10.0.0.0/16 | inbound udp 9100 from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | TCP | 9100 | 10.0.0.0/16 | inbound tcp 9100 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | UDP | 32803 | 10.0.0.0/16 | inbound udp 32803 from 10.0.0.0/16 | Migrated from source |
| 5 | inbound | TCP | 32803 | 10.0.0.0/16 | inbound tcp 32803 from 10.0.0.0/16 | Migrated from source |
| 6 | inbound | UDP | 20048 | 10.0.0.0/16 | inbound udp 20048 from 10.0.0.0/16 | Migrated from source |
| 7 | inbound | TCP | 20048 | 10.0.0.0/16 | inbound tcp 20048 from 10.0.0.0/16 | Migrated from source |
| 8 | inbound | UDP | 111 | 0.0.0.0/0 | inbound udp 111 | Migrated from source |
| 9 | inbound | TCP | 111 | 0.0.0.0/0 | inbound tcp 111 | Migrated from source |
| 10 | inbound | UDP | 2049 | 0.0.0.0/0 | inbound udp 2049 | Migrated from source |
| 11 | inbound | TCP | 2049 | 0.0.0.0/0 | inbound tcp 2049 | Migrated from source |
| 12 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 13 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 14 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 15 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 16 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 17 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 18 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 19 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my04-sg-03

**CSP ID:** sg-mj79ry1emh74fyd3ve1h | **VNet:** my04-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec288dd0-c6fa-8a49-2f60

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 2 | inbound | UDP | 9104 | 10.0.0.0/16 | inbound udp 9104 from 10.0.0.0/16 | Migrated from source |
| 3 | inbound | TCP | 9104 | 10.0.0.0/16 | inbound tcp 9104 from 10.0.0.0/16 | Migrated from source |
| 4 | inbound | UDP | 4444 | 10.0.0.0/16 | inbound udp 4444 from 10.0.0.0/16 | Migrated from source |
| 5 | inbound | TCP | 4444 | 10.0.0.0/16 | inbound tcp 4444 from 10.0.0.0/16 | Migrated from source |
| 6 | inbound | UDP | 4568 | 10.0.0.0/16 | inbound udp 4568 from 10.0.0.0/16 | Migrated from source |
| 7 | inbound | TCP | 4568 | 10.0.0.0/16 | inbound tcp 4568 from 10.0.0.0/16 | Migrated from source |
| 8 | inbound | UDP | 4567 | 10.0.0.0/16 | inbound udp 4567 from 10.0.0.0/16 | Migrated from source |
| 9 | inbound | TCP | 4567 | 10.0.0.0/16 | inbound tcp 4567 from 10.0.0.0/16 | Migrated from source |
| 10 | inbound | UDP | 3306 | 10.0.0.0/16 | inbound udp 3306 from 10.0.0.0/16 | Migrated from source |
| 11 | inbound | TCP | 3306 | 10.0.0.0/16 | inbound tcp 3306 from 10.0.0.0/16 | Migrated from source |
| 12 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 13 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 14 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 15 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 16 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 17 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 18 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 19 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source node network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** my04-vnet-01<br>**ID:** vpc-mj7x77ndgl64ijgqiqlga | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my04-subnet-01<br>**ID:** vsw-mj7n9mzainfnzpl4c4x01 | 10.0.1.0/24 | my04-vnet-01 |

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
| 1 | my04-sshkey-01 | tb9olhb6136go2ho5il5 | da1323e457ca703332e8f0ea2530007f | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.2551 |
| Daily | $6.12 |
| Monthly | $183.67 |
| Yearly | $2204.06 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | ecs.e-c1m1.large | $12.82 | 7.0% |
| ip-10-0-1-221 (migrated) | ecs.e-c1m4.xlarge | $113.90 | 62.0% |
| ip-10-0-1-138 (migrated) | ecs.e-c1m4.large | $56.95 | 31.0% |

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
  "message": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: my04-infra101)",
  "success": true
}
```

