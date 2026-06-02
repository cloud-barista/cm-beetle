# CM-Beetle test results for IBM

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with IBM cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: imdl/v0.1.5+ (4987539)
- imdl: v0.1.5+ (4987539)
- CB-Tumblebug: v0.12.13
- CB-Spider: v0.12.26
- CB-MapUI: v0.12.34
- Target CSP: IBM
- Target Region: au-syd
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: June 2, 2026
- Test Time: 20:57:10 KST
- Test Execution: 2026-06-02 20:57:10 KST

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

## Test result for IBM

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 15.53s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 2m54.786s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 38ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 7ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 26ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.345s | Pass |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.3s | Pass |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m22.808s | Pass |

**Overall Result**: 9/9 tests passed ✅

**Total Duration**: 6m8.262026084s

*Test executed on June 2, 2026 at 20:57:10 KST (2026-06-02 20:57:10 KST) using CM-Beetle automated test CLI*

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
    "csp": "ibm",
    "region": "au-syd"
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
      "status": "highly-matched",
      "description": "Candidate #1 | highly-matched | Overall Match Rate: Min=100.0% Max=100.0% Avg=100.0% | VMs: 3 total, 3 matched, 0 acceptable",
      "targetCloud": {
        "csp": "ibm",
        "region": "au-syd"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+nxf-2x2",
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bxf-4x16",
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bxf-2x8",
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
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
        "connectionName": "ibm-au-syd",
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
        "connectionName": "ibm-au-syd",
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
          "id": "ibm+au-syd+nxf-2x2",
          "uid": "tb7jggfvq658ffi7hanf",
          "cspSpecName": "nxf-2x2",
          "name": "ibm+au-syd+nxf-2x2",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.094,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard,spot]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:2000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "nano"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x2"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "Name",
              "value": "nxf-2x2"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:fixed,value:1}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:500,max:1500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:200}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:dependent}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[10,25,50,100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        },
        {
          "id": "ibm+au-syd+bxf-4x16",
          "uid": "tbtp8i585ail9mh3jnod",
          "cspSpecName": "bxf-4x16",
          "name": "ibm+au-syd+bxf-4x16",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.235,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard,spot]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:8000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:16}"
            },
            {
              "key": "Name",
              "value": "bxf-4x16"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:fixed,value:1}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:2000,max:7500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:200}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:4}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:dependent}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[50,100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        },
        {
          "id": "ibm+au-syd+bxf-2x8",
          "uid": "tbu5hc2ke8pnt5ssgcbo",
          "cspSpecName": "bxf-2x8",
          "name": "ibm+au-syd+bxf-2x8",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.117,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard,spot]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:4000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:8}"
            },
            {
              "key": "Name",
              "value": "bxf-2x8"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:fixed,value:1}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:1000,max:3500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:200}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:dependent}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[50,100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "ibm",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "regionList": [
            "au-syd"
          ],
          "id": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "uid": "tbtrt2lnuipiuftrap37",
          "name": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ibm-au-syd",
          "infraType": "",
          "fetchedTime": "2026.05.27 06:32:55 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)",
          "osDiskType": "NA",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "AllowedUse",
              "value": "{api_version:2024-11-28,bare_metal_server:true,instance:true}"
            },
            {
              "key": "CatalogOffering",
              "value": "{managed:false}"
            },
            {
              "key": "CreatedAt",
              "value": "2026-05-19T23:51:58.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "Encryption",
              "value": "none"
            },
            {
              "key": "File",
              "value": "{checksums:{sha256:05ab42b9bb4881c3944fc5452b46a275bb94a6366831b1a874fa708585c4ecb1},size:1}"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "ID",
              "value": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "MinimumProvisionedSize",
              "value": "10"
            },
            {
              "key": "Name",
              "value": "ibm-ubuntu-22-04-5-minimal-amd64-15"
            },
            {
              "key": "OperatingSystem",
              "value": "{allow_user_image_creation:true,architecture:amd64,dedicated_host_only:false,display_name:Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64),family:Ubuntu Linux,href:https://au-syd.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-22-04-amd64,name:ubuntu-22-04-amd64,user_data_format:cloud_init,vendor:Canonical,version:22.04 LTS Jammy Jellyfish Minimal Install}"
            },
            {
              "key": "Remote",
              "value": "{account:{id:811f8abfbd32425597dc7ba40da98fa6,resource_type:account}}"
            },
            {
              "key": "ResourceGroup",
              "value": "{href:https://resource-controller.cloud.ibm.com/v1/resource_groups/5807b5832a8741179b2e06ca2d2b3b96,id:5807b5832a8741179b2e06ca2d2b3b96,name:Default}"
            },
            {
              "key": "ResourceType",
              "value": "image"
            },
            {
              "key": "Status",
              "value": "available"
            },
            {
              "key": "UserDataFormat",
              "value": "cloud_init"
            },
            {
              "key": "Visibility",
              "value": "public"
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
          "connectionName": "ibm-au-syd",
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
          "connectionName": "ibm-au-syd",
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
          "connectionName": "ibm-au-syd",
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
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=50.0% Max=100.0% Avg=94.4% | VMs: 3 total, 2 matched, 1 acceptable",
      "targetCloud": {
        "csp": "ibm",
        "region": "au-syd"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+nxf-2x1",
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bx2-4x16",
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "ibm-au-syd",
            "specId": "ibm+au-syd+bx2-2x8",
            "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 100,
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
        "connectionName": "ibm-au-syd",
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
        "connectionName": "ibm-au-syd",
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
          "id": "ibm+au-syd+nxf-2x2",
          "uid": "tb7jggfvq658ffi7hanf",
          "cspSpecName": "nxf-2x2",
          "name": "ibm+au-syd+nxf-2x2",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.094,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard,spot]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:2000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "nano"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x2"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "Name",
              "value": "nxf-2x2"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:fixed,value:1}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:500,max:1500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:200}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:dependent}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[10,25,50,100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        },
        {
          "id": "ibm+au-syd+bxf-4x16",
          "uid": "tbtp8i585ail9mh3jnod",
          "cspSpecName": "bxf-4x16",
          "name": "ibm+au-syd+bxf-4x16",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.235,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard,spot]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:8000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:16}"
            },
            {
              "key": "Name",
              "value": "bxf-4x16"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:fixed,value:1}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:2000,max:7500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:200}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:4}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:dependent}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[50,100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        },
        {
          "id": "ibm+au-syd+bxf-2x8",
          "uid": "tbu5hc2ke8pnt5ssgcbo",
          "cspSpecName": "bxf-2x8",
          "name": "ibm+au-syd+bxf-2x8",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.117,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard,spot]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:4000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:8}"
            },
            {
              "key": "Name",
              "value": "bxf-2x8"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:fixed,value:1}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:1000,max:3500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:200}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:dependent}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[50,100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        },
        {
          "id": "ibm+au-syd+nxf-2x1",
          "uid": "tbm961b7uek2fdn6h3to",
          "cspSpecName": "nxf-2x1",
          "name": "ibm+au-syd+nxf-2x1",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1,
          "diskSizeGB": -1,
          "costPerHour": 0.062,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard,spot]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:2000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "nano"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x1"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:1}"
            },
            {
              "key": "Name",
              "value": "nxf-2x1"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:1,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:fixed,value:1}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:500,max:1500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:200}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:dependent}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[10,25,50,100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        },
        {
          "id": "ibm+au-syd+bx2-4x16",
          "uid": "tbgdlsg7fs0qf1nr53pr",
          "cspSpecName": "bx2-4x16",
          "name": "ibm+au-syd+bx2-4x16",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.241,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:8000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bx2-4x16"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:16}"
            },
            {
              "key": "Name",
              "value": "bx2-4x16"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:dependent}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:2000,max:7500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:100}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:4}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:fixed,value:intel}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:weighted,type:enum,values:[weighted]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        },
        {
          "id": "ibm+au-syd+bx2-2x8",
          "uid": "tb4036nluqjoeocg0vnh",
          "cspSpecName": "bx2-2x8",
          "name": "ibm+au-syd+bx2-2x8",
          "namespace": "system",
          "connectionName": "ibm-au-syd",
          "providerName": "ibm",
          "regionName": "au-syd",
          "regionLatitude": -33.86882,
          "regionLongitude": 151.209296,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.12,
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
              "key": "AvailabilityClass",
              "value": "{default:standard,type:enum,values:[standard]}"
            },
            {
              "key": "Bandwidth",
              "value": "{type:fixed,value:4000}"
            },
            {
              "key": "ClusterNetworkAttachmentCount",
              "value": "{type:enum,values:[0]}"
            },
            {
              "key": "ConfidentialComputeModes",
              "value": "{default:disabled,type:enum,values:[disabled]}"
            },
            {
              "key": "Family",
              "value": "balanced"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bx2-2x8"
            },
            {
              "key": "Memory",
              "value": "{type:fixed,value:8}"
            },
            {
              "key": "Name",
              "value": "bx2-2x8"
            },
            {
              "key": "NetworkAttachmentCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NetworkBandwidthMode",
              "value": "{type:fixed,value:divided}"
            },
            {
              "key": "NetworkInterfaceCount",
              "value": "{max:5,min:1,type:range}"
            },
            {
              "key": "NumaCount",
              "value": "{type:dependent}"
            },
            {
              "key": "OsArchitecture",
              "value": "{default:amd64,type:enum,values:[amd64]}"
            },
            {
              "key": "PortSpeed",
              "value": "{type:fixed,value:25000}"
            },
            {
              "key": "ReservationTerms",
              "value": "{type:enum,values:[one_year,three_year]}"
            },
            {
              "key": "ResourceType",
              "value": "instance_profile"
            },
            {
              "key": "SecureBootModes",
              "value": "{default:false,type:enum,values:[false]}"
            },
            {
              "key": "Status",
              "value": "current"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "{type:range,default:1000,max:3500,min:500,step:1}"
            },
            {
              "key": "VcpuArchitecture",
              "value": "{type:fixed,value:amd64}"
            },
            {
              "key": "VcpuBurstLimit",
              "value": "{type:fixed,value:100}"
            },
            {
              "key": "VcpuCount",
              "value": "{type:fixed,value:2}"
            },
            {
              "key": "VcpuManufacturer",
              "value": "{type:fixed,value:intel}"
            },
            {
              "key": "VcpuPercentage",
              "value": "{default:100,type:enum,values:[100]}"
            },
            {
              "key": "VolumeBandwidthQosModes",
              "value": "{default:weighted,type:enum,values:[weighted]}"
            },
            {
              "key": "Zones",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-1,name:au-syd-1}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-2,name:au-syd-2}; {href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/au-syd-3,name:au-syd-3}"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "ibm",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "regionList": [
            "au-syd"
          ],
          "id": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "uid": "tbtrt2lnuipiuftrap37",
          "name": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "ibm-au-syd",
          "infraType": "",
          "fetchedTime": "2026.05.27 06:32:55 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)",
          "osDiskType": "NA",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "AllowedUse",
              "value": "{api_version:2024-11-28,bare_metal_server:true,instance:true}"
            },
            {
              "key": "CatalogOffering",
              "value": "{managed:false}"
            },
            {
              "key": "CreatedAt",
              "value": "2026-05-19T23:51:58.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "Encryption",
              "value": "none"
            },
            {
              "key": "File",
              "value": "{checksums:{sha256:05ab42b9bb4881c3944fc5452b46a275bb94a6366831b1a874fa708585c4ecb1},size:1}"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "ID",
              "value": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d"
            },
            {
              "key": "MinimumProvisionedSize",
              "value": "10"
            },
            {
              "key": "Name",
              "value": "ibm-ubuntu-22-04-5-minimal-amd64-15"
            },
            {
              "key": "OperatingSystem",
              "value": "{allow_user_image_creation:true,architecture:amd64,dedicated_host_only:false,display_name:Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64),family:Ubuntu Linux,href:https://au-syd.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-22-04-amd64,name:ubuntu-22-04-amd64,user_data_format:cloud_init,vendor:Canonical,version:22.04 LTS Jammy Jellyfish Minimal Install}"
            },
            {
              "key": "Remote",
              "value": "{account:{id:811f8abfbd32425597dc7ba40da98fa6,resource_type:account}}"
            },
            {
              "key": "ResourceGroup",
              "value": "{href:https://resource-controller.cloud.ibm.com/v1/resource_groups/5807b5832a8741179b2e06ca2d2b3b96,id:5807b5832a8741179b2e06ca2d2b3b96,name:Default}"
            },
            {
              "key": "ResourceType",
              "value": "image"
            },
            {
              "key": "Status",
              "value": "available"
            },
            {
              "key": "UserDataFormat",
              "value": "cloud_init"
            },
            {
              "key": "Visibility",
              "value": "public"
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
          "connectionName": "ibm-au-syd",
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
          "connectionName": "ibm-au-syd",
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
          "connectionName": "ibm-au-syd",
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
  "id": "my06-infra101",
  "uid": "tbb155e34rt94il69ph3",
  "name": "my06-infra101",
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
    "sys.id": "my06-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my06-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbb155e34rt94il69ph3"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbt7k7116g5o02ekkgfm",
      "cspResourceName": "tbt7k7116g5o02ekkgfm",
      "cspResourceId": "02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457",
      "name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "Australia (Sydney)",
        "latitude": -33.86882,
        "longitude": 151.209296
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-06-02 11:59:56",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-02 11:59:56",
        "sys.cspResourceId": "02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457",
        "sys.cspResourceName": "tbt7k7116g5o02ekkgfm",
        "sys.id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tbt7k7116g5o02ekkgfm",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.102.221",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
      "connectionName": "ibm-au-syd",
      "connectionConfig": {
        "configName": "ibm-au-syd",
        "providerName": "ibm",
        "driverName": "ibm-driver-v1.0.so",
        "credentialName": "ibm",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ibm-au-syd",
        "regionZoneInfo": {
          "assignedRegion": "au-syd",
          "assignedZone": "au-syd-1"
        },
        "regionDetail": {
          "regionId": "au-syd",
          "regionName": "au-syd",
          "description": "Sydney (Australia)",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "zones": [
            "au-syd-1",
            "au-syd-2",
            "au-syd-3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ibm+au-syd+nxf-2x2",
      "cspSpecName": "nxf-2x2",
      "spec": {
        "cspSpecName": "nxf-2x2",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.094
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-5b6a8f03-385e-44ef-b1bc-825c662ed931",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d",
      "networkInterface": "flask-hardship-unfitting-uncapped",
      "securityGroupIds": [
        "my06-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAC8wdkeZ9gSMvTRBE8uUz6TGh+0YMlgOLO+Rh0C626WTd32ms8sDNvgMbaGZpFyUQtaomZR53J9kPjY7csXJAk=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:rhFh2QMNhzSTwiu1H3k4UCkoFE4F6j+Xp8OeEDSWowM",
        "firstUsedAt": "2026-06-02T12:00:09Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-06-02T12:00:09Z",
          "completedTime": "2026-06-02T12:00:13Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbt7k7116g5o02ekkgfm 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "2000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662-jwbn8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/volume_attachments/02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,name:garnet-magically-wifi-feisty,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,id:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,name:fondue-saturate-kettle-utility,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-02T11:59:23.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457"
        },
        {
          "key": "ID",
          "value": "02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "2"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbt7k7116g5o02ekkgfm"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/network_interfaces/02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,id:02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,name:flask-hardship-unfitting-uncapped,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,id:02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,name:dig-backlight-sincerity-roulette,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/network_interfaces/02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,id:02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,name:flask-hardship-unfitting-uncapped,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,id:02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,name:dig-backlight-sincerity-roulette,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x2,name:nxf-2x2,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "1500"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "500"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662-jwbn8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/volume_attachments/02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,name:garnet-magically-wifi-feisty,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,id:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,name:fondue-saturate-kettle-utility,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,id:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,name:tbans8fa21l4pkrjo0a0,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tb1h7ijtv4sdjjt74cq4",
      "cspResourceName": "tb1h7ijtv4sdjjt74cq4",
      "cspResourceId": "02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4",
      "name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "location": {
        "display": "Australia (Sydney)",
        "latitude": -33.86882,
        "longitude": 151.209296
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-06-02 12:00:04",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-02 12:00:04",
        "sys.cspResourceId": "02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4",
        "sys.cspResourceName": "tb1h7ijtv4sdjjt74cq4",
        "sys.id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tb1h7ijtv4sdjjt74cq4",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.98.218",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
      "connectionName": "ibm-au-syd",
      "connectionConfig": {
        "configName": "ibm-au-syd",
        "providerName": "ibm",
        "driverName": "ibm-driver-v1.0.so",
        "credentialName": "ibm",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ibm-au-syd",
        "regionZoneInfo": {
          "assignedRegion": "au-syd",
          "assignedZone": "au-syd-1"
        },
        "regionDetail": {
          "regionId": "au-syd",
          "regionName": "au-syd",
          "description": "Sydney (Australia)",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "zones": [
            "au-syd-1",
            "au-syd-2",
            "au-syd-3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ibm+au-syd+bxf-2x8",
      "cspSpecName": "bxf-2x8",
      "spec": {
        "cspSpecName": "bxf-2x8",
        "vCPU": 2,
        "memoryGiB": 8,
        "costPerHour": 0.117
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-5b6a8f03-385e-44ef-b1bc-825c662ed931",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d",
      "networkInterface": "underling-headband-jawline-dandy",
      "securityGroupIds": [
        "my06-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGPaje4p0cHi3xSR/tidAHCnzFneiJIw0aQiK/EwDvTEXEZR+qsC0LF+tFSFLRUT1u8A6v80/OKlulbvQMXzI6E=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:UUfEs/wSunSjtxc257xEkGRgYJVZi09b3FC0zZFK+oU",
        "firstUsedAt": "2026-06-02T12:00:12Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-02T12:00:09Z",
          "completedTime": "2026-06-02T12:00:30Z",
          "elapsedTime": 21,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to establish SSH connection to target host: ssh: handshake failed: EOF"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "4000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69-lc6sz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/volume_attachments/02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,name:mocker-steersman-mushroom-stock,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1d869fd7-f427-491a-a998-230be98d0f8f,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1d869fd7-f427-491a-a998-230be98d0f8f,id:r026-1d869fd7-f427-491a-a998-230be98d0f8f,name:secluded-hatbox-payphone-lecturer,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-02T11:59:25.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4"
        },
        {
          "key": "ID",
          "value": "02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "8"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tb1h7ijtv4sdjjt74cq4"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/network_interfaces/02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,id:02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,name:underling-headband-jawline-dandy,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-29641c82-c582-42aa-8330-7dd5b55583c2,id:02h7-29641c82-c582-42aa-8330-7dd5b55583c2,name:cleat-overstate-dislike-settling,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/network_interfaces/02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,id:02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,name:underling-headband-jawline-dandy,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-29641c82-c582-42aa-8330-7dd5b55583c2,id:02h7-29641c82-c582-42aa-8330-7dd5b55583c2,name:cleat-overstate-dislike-settling,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8,name:bxf-2x8,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "3000"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "1000"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69-lc6sz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/volume_attachments/02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,name:mocker-steersman-mushroom-stock,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1d869fd7-f427-491a-a998-230be98d0f8f,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1d869fd7-f427-491a-a998-230be98d0f8f,id:r026-1d869fd7-f427-491a-a998-230be98d0f8f,name:secluded-hatbox-payphone-lecturer,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,id:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,name:tbans8fa21l4pkrjo0a0,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tbfb0c8hrtma9tjabtm6",
      "cspResourceName": "tbfb0c8hrtma9tjabtm6",
      "cspResourceId": "02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6",
      "name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "location": {
        "display": "Australia (Sydney)",
        "latitude": -33.86882,
        "longitude": 151.209296
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-06-02 12:00:00",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-02 12:00:00",
        "sys.cspResourceId": "02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6",
        "sys.cspResourceName": "tbfb0c8hrtma9tjabtm6",
        "sys.id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tbfb0c8hrtma9tjabtm6",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.94.110",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
      "connectionName": "ibm-au-syd",
      "connectionConfig": {
        "configName": "ibm-au-syd",
        "providerName": "ibm",
        "driverName": "ibm-driver-v1.0.so",
        "credentialName": "ibm",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ibm-au-syd",
        "regionZoneInfo": {
          "assignedRegion": "au-syd",
          "assignedZone": "au-syd-1"
        },
        "regionDetail": {
          "regionId": "au-syd",
          "regionName": "au-syd",
          "description": "Sydney (Australia)",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "zones": [
            "au-syd-1",
            "au-syd-2",
            "au-syd-3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ibm+au-syd+bxf-4x16",
      "cspSpecName": "bxf-4x16",
      "spec": {
        "cspSpecName": "bxf-4x16",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.235
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-5b6a8f03-385e-44ef-b1bc-825c662ed931",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d",
      "networkInterface": "cast-handcraft-hulk-scalded",
      "securityGroupIds": [
        "my06-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMZia3tsPEXHAdMl3XqD20BJ7FgY29w9Zh5t4KDDweL4FDYix1xEB8GeagoakJ/0EqA0Pot7pz6z0SxH5pXZxFw=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:AxjNblckKR1i97qPsF9JhDs5lyqT3XRq+W7bi6GEgpc",
        "firstUsedAt": "2026-06-02T12:00:12Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-02T12:00:09Z",
          "completedTime": "2026-06-02T12:00:30Z",
          "elapsedTime": 21,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to establish SSH connection to target host: ssh: handshake failed: EOF"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "8000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199-dpdgt},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/volume_attachments/02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,name:riverside-fang-shampoo-affirm,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,id:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,name:runway-preshow-deed-truce,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-02T11:59:22.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6"
        },
        {
          "key": "ID",
          "value": "02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "16"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbfb0c8hrtma9tjabtm6"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/network_interfaces/02h7-2d45942c-5b73-4894-9f98-736c8d83526f,id:02h7-2d45942c-5b73-4894-9f98-736c8d83526f,name:cast-handcraft-hulk-scalded,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,id:02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,name:festoonery-cardigan-magically-obstinate,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/network_interfaces/02h7-2d45942c-5b73-4894-9f98-736c8d83526f,id:02h7-2d45942c-5b73-4894-9f98-736c8d83526f,name:cast-handcraft-hulk-scalded,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,id:02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,name:festoonery-cardigan-magically-obstinate,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16,name:bxf-4x16,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "6000"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "2000"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:4,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199-dpdgt},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/volume_attachments/02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,name:riverside-fang-shampoo-affirm,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,id:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,name:runway-preshow-deed-truce,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,id:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,name:tbans8fa21l4pkrjo0a0,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
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
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "159.23.102.221",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbt7k7116g5o02ekkgfm 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "159.23.98.218",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "nodeIp": "159.23.94.110",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
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
      "id": "my02-infra101",
      "uid": "tbv4r5dg0r6ctsgj7tfe",
      "name": "my02-infra101",
      "status": "Terminated:3 (R:0/3)",
      "statusCount": {
        "countTotal": 3,
        "countCreating": 0,
        "countRunning": 0,
        "countFailed": 0,
        "countSuspended": 0,
        "countRebooting": 0,
        "countTerminated": 3,
        "countSuspending": 0,
        "countResuming": 0,
        "countTerminating": 0,
        "countRegistering": 0,
        "countUndefined": 0
      },
      "targetStatus": "Terminated",
      "targetAction": "Terminate",
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
        "sys.id": "my02-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbv4r5dg0r6ctsgj7tfe"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb81h514blbmih6th0ra",
          "cspResourceName": "tb81h514blbmih6th0ra",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb81h514blbmih6th0ra",
          "name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Terminated",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "terminated VM. No action is acceptable except deletion",
          "createdTime": "2026-06-02 11:58:51",
          "label": {
            "createdBy": "tb81h514blbmih6th0ra",
            "keypair": "tbmeu2jlibrbjhpqt6ks",
            "publicip": "tb81h514blbmih6th0ra-51215-PublicIP",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-06-02 11:58:51",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb81h514blbmih6th0ra",
            "sys.cspResourceName": "tb81h514blbmih6th0ra",
            "sys.id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tb81h514blbmih6th0ra",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "40.89.209.210",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 30,
          "RootDeviceName": "Not visible in Azure",
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
          "cspSpecName": "Standard_B2als_v2",
          "spec": {
            "cspSpecName": "Standard_B2als_v2",
            "vCPU": 2,
            "memoryGiB": 3.90625,
            "costPerHour": 0.0432
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605280",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbaa7lm1qpjpqcm1blta",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbaa7lm1qpjpqcm1blta/subnets/tbdk04p2j6rjaddfeilr",
          "networkInterface": "tb81h514blbmih6th0ra-74785-VNic",
          "securityGroupIds": [
            "my02-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbmeu2jlibrbjhpqt6ks",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFW7tvj2HzsClA8X1H26zUY4x8P6NQYpxpab+6eweSK3WB1Z7hjR9loeEmmH+DcvmhMv+MY0Rswn2avB6fUySUc=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:iARQFNvURLH907WgwzHgMJPVSWgoyBhMXtYvYWM/db0",
            "firstUsedAt": "2026-06-02T11:58:58Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-02T11:58:57Z",
              "completedTime": "2026-06-02T11:59:01Z",
              "elapsedTime": 4,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tb81h514blbmih6th0ra 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb81h514blbmih6th0ra-74785-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb81h514blbmih6th0ra,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCpktoKLh1JaY7GzYG/FsvMS7Gxwj0F8FZzI9Q5KMffkjE6BTyUFuxHPgtyWbGKkZfk48XNdc0Yfiy11Oaol4Z/zUtnZkql9x1p8ktqwvK6RPaZjlxWjG2n4dvidBMsHCdpQjsmj3Zao8xsYG2fovFDgx2CUOeUEpCNp9uY2J54UumH6YLfVXrV2kQfeJkRXRQBlUUw6MObttcVhhNsVlaLxzeqb/sLaHEb5Pa6tO2Fok+VE++/OPAf5nubPirQp+Uo9ld7OhLolUnAfocho0frZSwaGv0VcU1iwpBWl66Kza8lKofk73hQrC3Ppr0eNb4VauY5Rh1sWuHC28C+Ikduc0JzfQ5WOczPDsMvFpylLvxHlMcxYSeEHtnO4/ykKs4duNp5SBIByiJsImpWaDNO44GJP3QdacdghiiLrb++6j7vAA2d/yRJpRrDjQyrMqJLTCYR9PLVmu/jNGx3mBdSareS2TIA0NG70tEvIzuCgNsAhGHOXstLqVbwsf8N6BP0bXx7l+lOTAGLu6/0V25cmutW0ZDkGwty+dUBKx7ocv6IVGZT+G6EUvnpcbGqtcfgG8o+ECZcuEnJGNKrQONYdEDC/WkfcekgpuqD58aS3pxWxb8bdrtQJYT2NQHSfLbtnNqCatVg11nSYvusxPZw13BXmo/5XKRFUCtnxN15qQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202605280,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202605280},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb81h514blbmih6th0ra_OsDisk_1_f68013d858ac4524824a70a9db72638e,storageAccountType:Premium_LRS},name:tb81h514blbmih6th0ra_OsDisk_1_f68013d858ac4524824a70a9db72638e,osType:Linux}},timeCreated:2026-06-02T11:58:00.0230074Z,vmId:a194131f-a308-4f9e-bd5f-1954226d1a73}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tb81h514blbmih6th0ra,keypair:tbmeu2jlibrbjhpqt6ks,publicip:tb81h514blbmih6th0ra-51215-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb81h514blbmih6th0ra"
            },
            {
              "key": "Name",
              "value": "tb81h514blbmih6th0ra"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tbj3l6ut11ch7mvkq2p2",
          "cspResourceName": "tbj3l6ut11ch7mvkq2p2",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbj3l6ut11ch7mvkq2p2",
          "name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Terminated",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "terminated VM. No action is acceptable except deletion",
          "createdTime": "2026-06-02 11:58:51",
          "label": {
            "createdBy": "tbj3l6ut11ch7mvkq2p2",
            "keypair": "tbmeu2jlibrbjhpqt6ks",
            "publicip": "tbj3l6ut11ch7mvkq2p2-63146-PublicIP",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-06-02 11:58:51",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbj3l6ut11ch7mvkq2p2",
            "sys.cspResourceName": "tbj3l6ut11ch7mvkq2p2",
            "sys.id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbj3l6ut11ch7mvkq2p2",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "52.231.195.8",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 30,
          "RootDeviceName": "Not visible in Azure",
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
          "cspSpecName": "Standard_B2as_v2",
          "spec": {
            "cspSpecName": "Standard_B2as_v2",
            "vCPU": 2,
            "memoryGiB": 7.8125,
            "costPerHour": 0.0865
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605280",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbaa7lm1qpjpqcm1blta",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbaa7lm1qpjpqcm1blta/subnets/tbdk04p2j6rjaddfeilr",
          "networkInterface": "tbj3l6ut11ch7mvkq2p2-3298-VNic",
          "securityGroupIds": [
            "my02-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbmeu2jlibrbjhpqt6ks",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJtUMh1cIUNrdvxDyS/sX2gg1UlgprQImyTmtt571zIcfmNpvML9//EZjbJhwyEPPk1DSV9P+GsLYIsancoLI2Q=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:vHAecS7HpoCmHLAn96798w3BOcO14c/PF8E2f32n6W4",
            "firstUsedAt": "2026-06-02T11:59:00Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-02T11:58:57Z",
              "completedTime": "2026-06-02T11:59:03Z",
              "elapsedTime": 6,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbj3l6ut11ch7mvkq2p2 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbj3l6ut11ch7mvkq2p2-3298-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbj3l6ut11ch7mvkq2p2,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCpktoKLh1JaY7GzYG/FsvMS7Gxwj0F8FZzI9Q5KMffkjE6BTyUFuxHPgtyWbGKkZfk48XNdc0Yfiy11Oaol4Z/zUtnZkql9x1p8ktqwvK6RPaZjlxWjG2n4dvidBMsHCdpQjsmj3Zao8xsYG2fovFDgx2CUOeUEpCNp9uY2J54UumH6YLfVXrV2kQfeJkRXRQBlUUw6MObttcVhhNsVlaLxzeqb/sLaHEb5Pa6tO2Fok+VE++/OPAf5nubPirQp+Uo9ld7OhLolUnAfocho0frZSwaGv0VcU1iwpBWl66Kza8lKofk73hQrC3Ppr0eNb4VauY5Rh1sWuHC28C+Ikduc0JzfQ5WOczPDsMvFpylLvxHlMcxYSeEHtnO4/ykKs4duNp5SBIByiJsImpWaDNO44GJP3QdacdghiiLrb++6j7vAA2d/yRJpRrDjQyrMqJLTCYR9PLVmu/jNGx3mBdSareS2TIA0NG70tEvIzuCgNsAhGHOXstLqVbwsf8N6BP0bXx7l+lOTAGLu6/0V25cmutW0ZDkGwty+dUBKx7ocv6IVGZT+G6EUvnpcbGqtcfgG8o+ECZcuEnJGNKrQONYdEDC/WkfcekgpuqD58aS3pxWxb8bdrtQJYT2NQHSfLbtnNqCatVg11nSYvusxPZw13BXmo/5XKRFUCtnxN15qQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202605280,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202605280},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbj3l6ut11ch7mvkq2p2_OsDisk_1_5e8befb085954c5198c00c1288c77c74,storageAccountType:Premium_LRS},name:tbj3l6ut11ch7mvkq2p2_OsDisk_1_5e8befb085954c5198c00c1288c77c74,osType:Linux}},timeCreated:2026-06-02T11:58:02.5923143Z,vmId:26a400ec-766d-4066-a986-31dee1af0a49}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbj3l6ut11ch7mvkq2p2,keypair:tbmeu2jlibrbjhpqt6ks,publicip:tbj3l6ut11ch7mvkq2p2-63146-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbj3l6ut11ch7mvkq2p2"
            },
            {
              "key": "Name",
              "value": "tbj3l6ut11ch7mvkq2p2"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tblumcks3qbld4l0rqdu",
          "cspResourceName": "tblumcks3qbld4l0rqdu",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tblumcks3qbld4l0rqdu",
          "name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "location": {
            "display": "Korea South",
            "latitude": 35.1796,
            "longitude": 129.0756
          },
          "status": "Terminated",
          "targetStatus": "Terminated",
          "targetAction": "Terminate",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "terminated VM. No action is acceptable except deletion",
          "createdTime": "2026-06-02 11:58:51",
          "label": {
            "createdBy": "tblumcks3qbld4l0rqdu",
            "keypair": "tbmeu2jlibrbjhpqt6ks",
            "publicip": "tblumcks3qbld4l0rqdu-15505-PublicIP",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-06-02 11:58:51",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tblumcks3qbld4l0rqdu",
            "sys.cspResourceName": "tblumcks3qbld4l0rqdu",
            "sys.id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tblumcks3qbld4l0rqdu",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.40.12",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.5",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 30,
          "RootDeviceName": "Not visible in Azure",
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
          "cspSpecName": "Standard_B4as_v2",
          "spec": {
            "cspSpecName": "Standard_B4as_v2",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.173
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605280",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbaa7lm1qpjpqcm1blta",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbaa7lm1qpjpqcm1blta/subnets/tbdk04p2j6rjaddfeilr",
          "networkInterface": "tblumcks3qbld4l0rqdu-93021-VNic",
          "securityGroupIds": [
            "my02-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbmeu2jlibrbjhpqt6ks",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBPTL9uzNQSU1yXWyoHY8FKtgjUxNkaKpGjhn6dtQjVMVJC3cpROCI06y1MwW8elX7OtijS4pmgWxkK29xA0kJ80=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:lVsWJmS92R46BkqahsTpl8eAW8rA6D39KJOZaJA3sj8",
            "firstUsedAt": "2026-06-02T11:59:00Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-02T11:58:57Z",
              "completedTime": "2026-06-02T11:59:02Z",
              "elapsedTime": 5,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tblumcks3qbld4l0rqdu 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tblumcks3qbld4l0rqdu-93021-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tblumcks3qbld4l0rqdu,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCpktoKLh1JaY7GzYG/FsvMS7Gxwj0F8FZzI9Q5KMffkjE6BTyUFuxHPgtyWbGKkZfk48XNdc0Yfiy11Oaol4Z/zUtnZkql9x1p8ktqwvK6RPaZjlxWjG2n4dvidBMsHCdpQjsmj3Zao8xsYG2fovFDgx2CUOeUEpCNp9uY2J54UumH6YLfVXrV2kQfeJkRXRQBlUUw6MObttcVhhNsVlaLxzeqb/sLaHEb5Pa6tO2Fok+VE++/OPAf5nubPirQp+Uo9ld7OhLolUnAfocho0frZSwaGv0VcU1iwpBWl66Kza8lKofk73hQrC3Ppr0eNb4VauY5Rh1sWuHC28C+Ikduc0JzfQ5WOczPDsMvFpylLvxHlMcxYSeEHtnO4/ykKs4duNp5SBIByiJsImpWaDNO44GJP3QdacdghiiLrb++6j7vAA2d/yRJpRrDjQyrMqJLTCYR9PLVmu/jNGx3mBdSareS2TIA0NG70tEvIzuCgNsAhGHOXstLqVbwsf8N6BP0bXx7l+lOTAGLu6/0V25cmutW0ZDkGwty+dUBKx7ocv6IVGZT+G6EUvnpcbGqtcfgG8o+ECZcuEnJGNKrQONYdEDC/WkfcekgpuqD58aS3pxWxb8bdrtQJYT2NQHSfLbtnNqCatVg11nSYvusxPZw13BXmo/5XKRFUCtnxN15qQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202605280,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202605280},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tblumcks3qbld4l0rqdu_OsDisk_1_9a7cccbd80604b8199a546a33c83bb41,storageAccountType:Premium_LRS},name:tblumcks3qbld4l0rqdu_OsDisk_1_9a7cccbd80604b8199a546a33c83bb41,osType:Linux}},timeCreated:2026-06-02T11:58:00.3944037Z,vmId:bb9a97ea-bbc0-4b18-b394-cb1accf01389}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tblumcks3qbld4l0rqdu,keypair:tbmeu2jlibrbjhpqt6ks,publicip:tblumcks3qbld4l0rqdu-15505-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tblumcks3qbld4l0rqdu"
            },
            {
              "key": "Name",
              "value": "tblumcks3qbld4l0rqdu"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
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
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "40.89.209.210",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tb81h514blbmih6th0ra 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "20.214.40.12",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tblumcks3qbld4l0rqdu 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "52.231.195.8",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbj3l6ut11ch7mvkq2p2 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "id": "my06-infra101",
      "uid": "tbb155e34rt94il69ph3",
      "name": "my06-infra101",
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
        "sys.id": "my06-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbb155e34rt94il69ph3"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbt7k7116g5o02ekkgfm",
          "cspResourceName": "tbt7k7116g5o02ekkgfm",
          "cspResourceId": "02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457",
          "name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-06-02 11:59:56",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-06-02 11:59:56",
            "sys.cspResourceId": "02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457",
            "sys.cspResourceName": "tbt7k7116g5o02ekkgfm",
            "sys.id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my06-subnet-01",
            "sys.uid": "tbt7k7116g5o02ekkgfm",
            "sys.vNetId": "my06-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.102.221",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.5",
          "privateDNS": "",
          "rootDiskType": "general-purpose",
          "rootDiskSize": 100,
          "RootDeviceName": "Not visible in IBM",
          "connectionName": "ibm-au-syd",
          "connectionConfig": {
            "configName": "ibm-au-syd",
            "providerName": "ibm",
            "driverName": "ibm-driver-v1.0.so",
            "credentialName": "ibm",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ibm-au-syd",
            "regionZoneInfo": {
              "assignedRegion": "au-syd",
              "assignedZone": "au-syd-1"
            },
            "regionDetail": {
              "regionId": "au-syd",
              "regionName": "au-syd",
              "description": "Sydney (Australia)",
              "location": {
                "display": "Australia (Sydney)",
                "latitude": -33.86882,
                "longitude": 151.209296
              },
              "zones": [
                "au-syd-1",
                "au-syd-2",
                "au-syd-3"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ibm+au-syd+nxf-2x2",
          "cspSpecName": "nxf-2x2",
          "spec": {
            "cspSpecName": "nxf-2x2",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.094
          },
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-vnet-01",
          "cspVNetId": "r026-5b6a8f03-385e-44ef-b1bc-825c662ed931",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d",
          "networkInterface": "flask-hardship-unfitting-uncapped",
          "securityGroupIds": [
            "my06-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": "r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAC8wdkeZ9gSMvTRBE8uUz6TGh+0YMlgOLO+Rh0C626WTd32ms8sDNvgMbaGZpFyUQtaomZR53J9kPjY7csXJAk=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:rhFh2QMNhzSTwiu1H3k4UCkoFE4F6j+Xp8OeEDSWowM",
            "firstUsedAt": "2026-06-02T12:00:09Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-02T12:00:09Z",
              "completedTime": "2026-06-02T12:00:13Z",
              "elapsedTime": 4,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbt7k7116g5o02ekkgfm 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Availability",
              "value": "{class:standard}"
            },
            {
              "key": "AvailabilityPolicy",
              "value": "{host_failure:restart,preemption:stop}"
            },
            {
              "key": "Bandwidth",
              "value": "2000"
            },
            {
              "key": "BootVolumeAttachment",
              "value": "{device:{id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662-jwbn8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/volume_attachments/02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,name:garnet-magically-wifi-feisty,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,id:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,name:fondue-saturate-kettle-utility,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-06-02T11:59:23.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457"
            },
            {
              "key": "EnableSecureBoot",
              "value": "false"
            },
            {
              "key": "HealthState",
              "value": "ok"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457"
            },
            {
              "key": "ID",
              "value": "02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457"
            },
            {
              "key": "Image",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
            },
            {
              "key": "LifecycleState",
              "value": "stable"
            },
            {
              "key": "Memory",
              "value": "2"
            },
            {
              "key": "MetadataService",
              "value": "{enabled:false,protocol:http,response_hop_limit:1}"
            },
            {
              "key": "Name",
              "value": "tbt7k7116g5o02ekkgfm"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/network_interfaces/02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,id:02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,name:flask-hardship-unfitting-uncapped,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,id:02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,name:dig-backlight-sincerity-roulette,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/network_interfaces/02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,id:02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,name:flask-hardship-unfitting-uncapped,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,id:02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,name:dig-backlight-sincerity-roulette,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
            },
            {
              "key": "Profile",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x2,name:nxf-2x2,resource_type:instance_profile}"
            },
            {
              "key": "ReservationAffinity",
              "value": "{policy:automatic,pool:[]}"
            },
            {
              "key": "ResourceGroup",
              "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
            },
            {
              "key": "ResourceType",
              "value": "instance"
            },
            {
              "key": "Startable",
              "value": "true"
            },
            {
              "key": "Status",
              "value": "running"
            },
            {
              "key": "TotalNetworkBandwidth",
              "value": "1500"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "500"
            },
            {
              "key": "Vcpu",
              "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
            },
            {
              "key": "VolumeAttachments",
              "value": "{device:{id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662-jwbn8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/volume_attachments/02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,name:garnet-magically-wifi-feisty,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,id:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,name:fondue-saturate-kettle-utility,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,id:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,name:tbans8fa21l4pkrjo0a0,resource_type:vpc}"
            },
            {
              "key": "Zone",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tb1h7ijtv4sdjjt74cq4",
          "cspResourceName": "tb1h7ijtv4sdjjt74cq4",
          "cspResourceId": "02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4",
          "name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-06-02 12:00:04",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-06-02 12:00:04",
            "sys.cspResourceId": "02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4",
            "sys.cspResourceName": "tb1h7ijtv4sdjjt74cq4",
            "sys.id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my06-subnet-01",
            "sys.uid": "tb1h7ijtv4sdjjt74cq4",
            "sys.vNetId": "my06-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.98.218",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "general-purpose",
          "rootDiskSize": 100,
          "RootDeviceName": "Not visible in IBM",
          "connectionName": "ibm-au-syd",
          "connectionConfig": {
            "configName": "ibm-au-syd",
            "providerName": "ibm",
            "driverName": "ibm-driver-v1.0.so",
            "credentialName": "ibm",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ibm-au-syd",
            "regionZoneInfo": {
              "assignedRegion": "au-syd",
              "assignedZone": "au-syd-1"
            },
            "regionDetail": {
              "regionId": "au-syd",
              "regionName": "au-syd",
              "description": "Sydney (Australia)",
              "location": {
                "display": "Australia (Sydney)",
                "latitude": -33.86882,
                "longitude": 151.209296
              },
              "zones": [
                "au-syd-1",
                "au-syd-2",
                "au-syd-3"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ibm+au-syd+bxf-2x8",
          "cspSpecName": "bxf-2x8",
          "spec": {
            "cspSpecName": "bxf-2x8",
            "vCPU": 2,
            "memoryGiB": 8,
            "costPerHour": 0.117
          },
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-vnet-01",
          "cspVNetId": "r026-5b6a8f03-385e-44ef-b1bc-825c662ed931",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d",
          "networkInterface": "underling-headband-jawline-dandy",
          "securityGroupIds": [
            "my06-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": "r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGPaje4p0cHi3xSR/tidAHCnzFneiJIw0aQiK/EwDvTEXEZR+qsC0LF+tFSFLRUT1u8A6v80/OKlulbvQMXzI6E=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:UUfEs/wSunSjtxc257xEkGRgYJVZi09b3FC0zZFK+oU",
            "firstUsedAt": "2026-06-02T12:00:12Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-06-02T12:00:09Z",
              "completedTime": "2026-06-02T12:00:30Z",
              "elapsedTime": 21,
              "resultSummary": "Command execution failed",
              "errorMessage": "failed to establish SSH connection to target host: ssh: handshake failed: EOF"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Availability",
              "value": "{class:standard}"
            },
            {
              "key": "AvailabilityPolicy",
              "value": "{host_failure:restart,preemption:stop}"
            },
            {
              "key": "Bandwidth",
              "value": "4000"
            },
            {
              "key": "BootVolumeAttachment",
              "value": "{device:{id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69-lc6sz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/volume_attachments/02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,name:mocker-steersman-mushroom-stock,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1d869fd7-f427-491a-a998-230be98d0f8f,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1d869fd7-f427-491a-a998-230be98d0f8f,id:r026-1d869fd7-f427-491a-a998-230be98d0f8f,name:secluded-hatbox-payphone-lecturer,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-06-02T11:59:25.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4"
            },
            {
              "key": "EnableSecureBoot",
              "value": "false"
            },
            {
              "key": "HealthState",
              "value": "ok"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4"
            },
            {
              "key": "ID",
              "value": "02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4"
            },
            {
              "key": "Image",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
            },
            {
              "key": "LifecycleState",
              "value": "stable"
            },
            {
              "key": "Memory",
              "value": "8"
            },
            {
              "key": "MetadataService",
              "value": "{enabled:false,protocol:http,response_hop_limit:1}"
            },
            {
              "key": "Name",
              "value": "tb1h7ijtv4sdjjt74cq4"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/network_interfaces/02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,id:02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,name:underling-headband-jawline-dandy,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-29641c82-c582-42aa-8330-7dd5b55583c2,id:02h7-29641c82-c582-42aa-8330-7dd5b55583c2,name:cleat-overstate-dislike-settling,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/network_interfaces/02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,id:02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,name:underling-headband-jawline-dandy,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-29641c82-c582-42aa-8330-7dd5b55583c2,id:02h7-29641c82-c582-42aa-8330-7dd5b55583c2,name:cleat-overstate-dislike-settling,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
            },
            {
              "key": "Profile",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8,name:bxf-2x8,resource_type:instance_profile}"
            },
            {
              "key": "ReservationAffinity",
              "value": "{policy:automatic,pool:[]}"
            },
            {
              "key": "ResourceGroup",
              "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
            },
            {
              "key": "ResourceType",
              "value": "instance"
            },
            {
              "key": "Startable",
              "value": "true"
            },
            {
              "key": "Status",
              "value": "running"
            },
            {
              "key": "TotalNetworkBandwidth",
              "value": "3000"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "1000"
            },
            {
              "key": "Vcpu",
              "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
            },
            {
              "key": "VolumeAttachments",
              "value": "{device:{id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69-lc6sz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/volume_attachments/02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,name:mocker-steersman-mushroom-stock,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1d869fd7-f427-491a-a998-230be98d0f8f,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1d869fd7-f427-491a-a998-230be98d0f8f,id:r026-1d869fd7-f427-491a-a998-230be98d0f8f,name:secluded-hatbox-payphone-lecturer,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,id:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,name:tbans8fa21l4pkrjo0a0,resource_type:vpc}"
            },
            {
              "key": "Zone",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbfb0c8hrtma9tjabtm6",
          "cspResourceName": "tbfb0c8hrtma9tjabtm6",
          "cspResourceId": "02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6",
          "name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2026-06-02 12:00:00",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "ibm-au-syd",
            "sys.createdTime": "2026-06-02 12:00:00",
            "sys.cspResourceId": "02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6",
            "sys.cspResourceName": "tbfb0c8hrtma9tjabtm6",
            "sys.id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my06-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my06-subnet-01",
            "sys.uid": "tbfb0c8hrtma9tjabtm6",
            "sys.vNetId": "my06-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "au-syd",
            "zone": "au-syd-1"
          },
          "publicIP": "159.23.94.110",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "general-purpose",
          "rootDiskSize": 100,
          "RootDeviceName": "Not visible in IBM",
          "connectionName": "ibm-au-syd",
          "connectionConfig": {
            "configName": "ibm-au-syd",
            "providerName": "ibm",
            "driverName": "ibm-driver-v1.0.so",
            "credentialName": "ibm",
            "credentialHolder": "admin",
            "regionZoneInfoName": "ibm-au-syd",
            "regionZoneInfo": {
              "assignedRegion": "au-syd",
              "assignedZone": "au-syd-1"
            },
            "regionDetail": {
              "regionId": "au-syd",
              "regionName": "au-syd",
              "description": "Sydney (Australia)",
              "location": {
                "display": "Australia (Sydney)",
                "latitude": -33.86882,
                "longitude": 151.209296
              },
              "zones": [
                "au-syd-1",
                "au-syd-2",
                "au-syd-3"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "ibm+au-syd+bxf-4x16",
          "cspSpecName": "bxf-4x16",
          "spec": {
            "cspSpecName": "bxf-4x16",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.235
          },
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "resourceType": "image",
            "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
          },
          "vNetId": "my06-vnet-01",
          "cspVNetId": "r026-5b6a8f03-385e-44ef-b1bc-825c662ed931",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d",
          "networkInterface": "cast-handcraft-hulk-scalded",
          "securityGroupIds": [
            "my06-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": "r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMZia3tsPEXHAdMl3XqD20BJ7FgY29w9Zh5t4KDDweL4FDYix1xEB8GeagoakJ/0EqA0Pot7pz6z0SxH5pXZxFw=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:AxjNblckKR1i97qPsF9JhDs5lyqT3XRq+W7bi6GEgpc",
            "firstUsedAt": "2026-06-02T12:00:12Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-06-02T12:00:09Z",
              "completedTime": "2026-06-02T12:00:30Z",
              "elapsedTime": 21,
              "resultSummary": "Command execution failed",
              "errorMessage": "failed to establish SSH connection to target host: ssh: handshake failed: EOF"
            }
          ],
          "addtionalDetails": [
            {
              "key": "Availability",
              "value": "{class:standard}"
            },
            {
              "key": "AvailabilityPolicy",
              "value": "{host_failure:restart,preemption:stop}"
            },
            {
              "key": "Bandwidth",
              "value": "8000"
            },
            {
              "key": "BootVolumeAttachment",
              "value": "{device:{id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199-dpdgt},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/volume_attachments/02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,name:riverside-fang-shampoo-affirm,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,id:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,name:runway-preshow-deed-truce,resource_type:volume}}"
            },
            {
              "key": "ConfidentialComputeMode",
              "value": "disabled"
            },
            {
              "key": "CreatedAt",
              "value": "2026-06-02T11:59:22.000Z"
            },
            {
              "key": "CRN",
              "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6"
            },
            {
              "key": "EnableSecureBoot",
              "value": "false"
            },
            {
              "key": "HealthState",
              "value": "ok"
            },
            {
              "key": "Href",
              "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6"
            },
            {
              "key": "ID",
              "value": "02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6"
            },
            {
              "key": "Image",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
            },
            {
              "key": "LifecycleState",
              "value": "stable"
            },
            {
              "key": "Memory",
              "value": "16"
            },
            {
              "key": "MetadataService",
              "value": "{enabled:false,protocol:http,response_hop_limit:1}"
            },
            {
              "key": "Name",
              "value": "tbfb0c8hrtma9tjabtm6"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/network_interfaces/02h7-2d45942c-5b73-4894-9f98-736c8d83526f,id:02h7-2d45942c-5b73-4894-9f98-736c8d83526f,name:cast-handcraft-hulk-scalded,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,id:02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,name:festoonery-cardigan-magically-obstinate,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
            },
            {
              "key": "NumaCount",
              "value": "1"
            },
            {
              "key": "PrimaryNetworkInterface",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/network_interfaces/02h7-2d45942c-5b73-4894-9f98-736c8d83526f,id:02h7-2d45942c-5b73-4894-9f98-736c8d83526f,name:cast-handcraft-hulk-scalded,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,id:02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,name:festoonery-cardigan-magically-obstinate,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
            },
            {
              "key": "Profile",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16,name:bxf-4x16,resource_type:instance_profile}"
            },
            {
              "key": "ReservationAffinity",
              "value": "{policy:automatic,pool:[]}"
            },
            {
              "key": "ResourceGroup",
              "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
            },
            {
              "key": "ResourceType",
              "value": "instance"
            },
            {
              "key": "Startable",
              "value": "true"
            },
            {
              "key": "Status",
              "value": "running"
            },
            {
              "key": "TotalNetworkBandwidth",
              "value": "6000"
            },
            {
              "key": "TotalVolumeBandwidth",
              "value": "2000"
            },
            {
              "key": "Vcpu",
              "value": "{architecture:amd64,count:4,manufacturer:intel,percentage:100}"
            },
            {
              "key": "VolumeAttachments",
              "value": "{device:{id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199-dpdgt},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/volume_attachments/02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,name:riverside-fang-shampoo-affirm,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,id:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,name:runway-preshow-deed-truce,resource_type:volume}}"
            },
            {
              "key": "VolumeBandwidthQosMode",
              "value": "pooled"
            },
            {
              "key": "VPC",
              "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,id:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,name:tbans8fa21l4pkrjo0a0,resource_type:vpc}"
            },
            {
              "key": "Zone",
              "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
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
            "infraId": "my06-infra101",
            "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "159.23.102.221",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbt7k7116g5o02ekkgfm 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my06-infra101",
            "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "159.23.98.218",
            "command": {
              "0": "uname -a"
            },
            "stdout": {},
            "stderr": {},
            "err": null
          },
          {
            "infraId": "my06-infra101",
            "nodeId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "159.23.94.110",
            "command": {
              "0": "uname -a"
            },
            "stdout": {},
            "stderr": {},
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
    "my02-infra101",
    "my06-infra101"
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
  "id": "my06-infra101",
  "uid": "tbb155e34rt94il69ph3",
  "name": "my06-infra101",
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
    "sys.id": "my06-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my06-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbb155e34rt94il69ph3"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbt7k7116g5o02ekkgfm",
      "cspResourceName": "tbt7k7116g5o02ekkgfm",
      "cspResourceId": "02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457",
      "name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "location": {
        "display": "Australia (Sydney)",
        "latitude": -33.86882,
        "longitude": 151.209296
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-06-02 11:59:56",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-02 11:59:56",
        "sys.cspResourceId": "02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457",
        "sys.cspResourceName": "tbt7k7116g5o02ekkgfm",
        "sys.id": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tbt7k7116g5o02ekkgfm",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.102.221",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
      "connectionName": "ibm-au-syd",
      "connectionConfig": {
        "configName": "ibm-au-syd",
        "providerName": "ibm",
        "driverName": "ibm-driver-v1.0.so",
        "credentialName": "ibm",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ibm-au-syd",
        "regionZoneInfo": {
          "assignedRegion": "au-syd",
          "assignedZone": "au-syd-1"
        },
        "regionDetail": {
          "regionId": "au-syd",
          "regionName": "au-syd",
          "description": "Sydney (Australia)",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "zones": [
            "au-syd-1",
            "au-syd-2",
            "au-syd-3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ibm+au-syd+nxf-2x2",
      "cspSpecName": "nxf-2x2",
      "spec": {
        "cspSpecName": "nxf-2x2",
        "vCPU": 2,
        "memoryGiB": 2,
        "costPerHour": 0.094
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-5b6a8f03-385e-44ef-b1bc-825c662ed931",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d",
      "networkInterface": "flask-hardship-unfitting-uncapped",
      "securityGroupIds": [
        "my06-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAC8wdkeZ9gSMvTRBE8uUz6TGh+0YMlgOLO+Rh0C626WTd32ms8sDNvgMbaGZpFyUQtaomZR53J9kPjY7csXJAk=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:rhFh2QMNhzSTwiu1H3k4UCkoFE4F6j+Xp8OeEDSWowM",
        "firstUsedAt": "2026-06-02T12:00:09Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-06-02T12:00:09Z",
          "completedTime": "2026-06-02T12:00:13Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbt7k7116g5o02ekkgfm 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "2000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662-jwbn8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/volume_attachments/02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,name:garnet-magically-wifi-feisty,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,id:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,name:fondue-saturate-kettle-utility,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-02T11:59:23.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457"
        },
        {
          "key": "ID",
          "value": "02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "2"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbt7k7116g5o02ekkgfm"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/network_interfaces/02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,id:02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,name:flask-hardship-unfitting-uncapped,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,id:02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,name:dig-backlight-sincerity-roulette,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/network_interfaces/02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,id:02h7-6b3534ae-f81a-45db-9555-7e3a76415ba3,name:flask-hardship-unfitting-uncapped,primary_ip:{address:10.0.1.5,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,id:02h7-2d482142-2179-4b13-baa5-44bc5c76b85d,name:dig-backlight-sincerity-roulette,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/nxf-2x2,name:nxf-2x2,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "1500"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "500"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662-jwbn8},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457/volume_attachments/02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,id:02h7-fb51cdd2-60f4-4e13-92ae-ceeb810b6662,name:garnet-magically-wifi-feisty,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,id:r026-34a55d8f-c6ca-4737-abb3-f7fca0efdcb5,name:fondue-saturate-kettle-utility,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,id:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,name:tbans8fa21l4pkrjo0a0,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tb1h7ijtv4sdjjt74cq4",
      "cspResourceName": "tb1h7ijtv4sdjjt74cq4",
      "cspResourceId": "02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4",
      "name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "location": {
        "display": "Australia (Sydney)",
        "latitude": -33.86882,
        "longitude": 151.209296
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-06-02 12:00:04",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-02 12:00:04",
        "sys.cspResourceId": "02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4",
        "sys.cspResourceName": "tb1h7ijtv4sdjjt74cq4",
        "sys.id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tb1h7ijtv4sdjjt74cq4",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.98.218",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
      "connectionName": "ibm-au-syd",
      "connectionConfig": {
        "configName": "ibm-au-syd",
        "providerName": "ibm",
        "driverName": "ibm-driver-v1.0.so",
        "credentialName": "ibm",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ibm-au-syd",
        "regionZoneInfo": {
          "assignedRegion": "au-syd",
          "assignedZone": "au-syd-1"
        },
        "regionDetail": {
          "regionId": "au-syd",
          "regionName": "au-syd",
          "description": "Sydney (Australia)",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "zones": [
            "au-syd-1",
            "au-syd-2",
            "au-syd-3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ibm+au-syd+bxf-2x8",
      "cspSpecName": "bxf-2x8",
      "spec": {
        "cspSpecName": "bxf-2x8",
        "vCPU": 2,
        "memoryGiB": 8,
        "costPerHour": 0.117
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-5b6a8f03-385e-44ef-b1bc-825c662ed931",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d",
      "networkInterface": "underling-headband-jawline-dandy",
      "securityGroupIds": [
        "my06-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGPaje4p0cHi3xSR/tidAHCnzFneiJIw0aQiK/EwDvTEXEZR+qsC0LF+tFSFLRUT1u8A6v80/OKlulbvQMXzI6E=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:UUfEs/wSunSjtxc257xEkGRgYJVZi09b3FC0zZFK+oU",
        "firstUsedAt": "2026-06-02T12:00:12Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-02T12:00:09Z",
          "completedTime": "2026-06-02T12:00:30Z",
          "elapsedTime": 21,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to establish SSH connection to target host: ssh: handshake failed: EOF"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "4000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69-lc6sz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/volume_attachments/02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,name:mocker-steersman-mushroom-stock,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1d869fd7-f427-491a-a998-230be98d0f8f,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1d869fd7-f427-491a-a998-230be98d0f8f,id:r026-1d869fd7-f427-491a-a998-230be98d0f8f,name:secluded-hatbox-payphone-lecturer,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-02T11:59:25.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4"
        },
        {
          "key": "ID",
          "value": "02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "8"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tb1h7ijtv4sdjjt74cq4"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/network_interfaces/02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,id:02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,name:underling-headband-jawline-dandy,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-29641c82-c582-42aa-8330-7dd5b55583c2,id:02h7-29641c82-c582-42aa-8330-7dd5b55583c2,name:cleat-overstate-dislike-settling,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/network_interfaces/02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,id:02h7-f67cddfa-e676-4345-9ddb-f30c4ffa67af,name:underling-headband-jawline-dandy,primary_ip:{address:10.0.1.6,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-29641c82-c582-42aa-8330-7dd5b55583c2,id:02h7-29641c82-c582-42aa-8330-7dd5b55583c2,name:cleat-overstate-dislike-settling,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-2x8,name:bxf-2x8,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "3000"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "1000"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:2,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69-lc6sz},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4/volume_attachments/02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,id:02h7-18c2d556-6054-4a2c-aa66-cd5c81867f69,name:mocker-steersman-mushroom-stock,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-1d869fd7-f427-491a-a998-230be98d0f8f,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-1d869fd7-f427-491a-a998-230be98d0f8f,id:r026-1d869fd7-f427-491a-a998-230be98d0f8f,name:secluded-hatbox-payphone-lecturer,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,id:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,name:tbans8fa21l4pkrjo0a0,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tbfb0c8hrtma9tjabtm6",
      "cspResourceName": "tbfb0c8hrtma9tjabtm6",
      "cspResourceId": "02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6",
      "name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "location": {
        "display": "Australia (Sydney)",
        "latitude": -33.86882,
        "longitude": 151.209296
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2026-06-02 12:00:00",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "ibm-au-syd",
        "sys.createdTime": "2026-06-02 12:00:00",
        "sys.cspResourceId": "02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6",
        "sys.cspResourceName": "tbfb0c8hrtma9tjabtm6",
        "sys.id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my06-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my06-subnet-01",
        "sys.uid": "tbfb0c8hrtma9tjabtm6",
        "sys.vNetId": "my06-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "au-syd",
        "zone": "au-syd-1"
      },
      "publicIP": "159.23.94.110",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "general-purpose",
      "rootDiskSize": 100,
      "RootDeviceName": "Not visible in IBM",
      "connectionName": "ibm-au-syd",
      "connectionConfig": {
        "configName": "ibm-au-syd",
        "providerName": "ibm",
        "driverName": "ibm-driver-v1.0.so",
        "credentialName": "ibm",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ibm-au-syd",
        "regionZoneInfo": {
          "assignedRegion": "au-syd",
          "assignedZone": "au-syd-1"
        },
        "regionDetail": {
          "regionId": "au-syd",
          "regionName": "au-syd",
          "description": "Sydney (Australia)",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "zones": [
            "au-syd-1",
            "au-syd-2",
            "au-syd-3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "ibm+au-syd+bxf-4x16",
      "cspSpecName": "bxf-4x16",
      "spec": {
        "cspSpecName": "bxf-4x16",
        "vCPU": 4,
        "memoryGiB": 16,
        "costPerHour": 0.235
      },
      "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
      "image": {
        "resourceType": "image",
        "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)"
      },
      "vNetId": "my06-vnet-01",
      "cspVNetId": "r026-5b6a8f03-385e-44ef-b1bc-825c662ed931",
      "subnetId": "my06-subnet-01",
      "cspSubnetId": "02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d",
      "networkInterface": "cast-handcraft-hulk-scalded",
      "securityGroupIds": [
        "my06-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my06-sshkey-01",
      "cspSshKeyId": "r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMZia3tsPEXHAdMl3XqD20BJ7FgY29w9Zh5t4KDDweL4FDYix1xEB8GeagoakJ/0EqA0Pot7pz6z0SxH5pXZxFw=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:AxjNblckKR1i97qPsF9JhDs5lyqT3XRq+W7bi6GEgpc",
        "firstUsedAt": "2026-06-02T12:00:12Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Failed",
          "startedTime": "2026-06-02T12:00:09Z",
          "completedTime": "2026-06-02T12:00:30Z",
          "elapsedTime": 21,
          "resultSummary": "Command execution failed",
          "errorMessage": "failed to establish SSH connection to target host: ssh: handshake failed: EOF"
        }
      ],
      "addtionalDetails": [
        {
          "key": "Availability",
          "value": "{class:standard}"
        },
        {
          "key": "AvailabilityPolicy",
          "value": "{host_failure:restart,preemption:stop}"
        },
        {
          "key": "Bandwidth",
          "value": "8000"
        },
        {
          "key": "BootVolumeAttachment",
          "value": "{device:{id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199-dpdgt},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/volume_attachments/02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,name:riverside-fang-shampoo-affirm,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,id:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,name:runway-preshow-deed-truce,resource_type:volume}}"
        },
        {
          "key": "ConfidentialComputeMode",
          "value": "disabled"
        },
        {
          "key": "CreatedAt",
          "value": "2026-06-02T11:59:22.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::instance:02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6"
        },
        {
          "key": "EnableSecureBoot",
          "value": "false"
        },
        {
          "key": "HealthState",
          "value": "ok"
        },
        {
          "key": "Href",
          "value": "https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6"
        },
        {
          "key": "ID",
          "value": "02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6"
        },
        {
          "key": "Image",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/811f8abfbd32425597dc7ba40da98fa6::image:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,href:https://au-syd.iaas.cloud.ibm.com/v1/images/r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,id:r026-c8e249d4-f148-4416-a3c6-555b7a02f67d,name:ibm-ubuntu-22-04-5-minimal-amd64-15,resource_type:image}"
        },
        {
          "key": "LifecycleState",
          "value": "stable"
        },
        {
          "key": "Memory",
          "value": "16"
        },
        {
          "key": "MetadataService",
          "value": "{enabled:false,protocol:http,response_hop_limit:1}"
        },
        {
          "key": "Name",
          "value": "tbfb0c8hrtma9tjabtm6"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/network_interfaces/02h7-2d45942c-5b73-4894-9f98-736c8d83526f,id:02h7-2d45942c-5b73-4894-9f98-736c8d83526f,name:cast-handcraft-hulk-scalded,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,id:02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,name:festoonery-cardigan-magically-obstinate,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "NumaCount",
          "value": "1"
        },
        {
          "key": "PrimaryNetworkInterface",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/network_interfaces/02h7-2d45942c-5b73-4894-9f98-736c8d83526f,id:02h7-2d45942c-5b73-4894-9f98-736c8d83526f,name:cast-handcraft-hulk-scalded,primary_ip:{address:10.0.1.4,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d/reserved_ips/02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,id:02h7-eb5f7396-9c2b-42b7-9fb6-cc21304b10bf,name:festoonery-cardigan-magically-obstinate,resource_type:subnet_reserved_ip},resource_type:network_interface,subnet:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,href:https://au-syd.iaas.cloud.ibm.com/v1/subnets/02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,id:02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d,name:tbif7mdf8fhjkp0ck6ut,resource_type:subnet}}"
        },
        {
          "key": "Profile",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/instance/profiles/bxf-4x16,name:bxf-4x16,resource_type:instance_profile}"
        },
        {
          "key": "ReservationAffinity",
          "value": "{policy:automatic,pool:[]}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "instance"
        },
        {
          "key": "Startable",
          "value": "true"
        },
        {
          "key": "Status",
          "value": "running"
        },
        {
          "key": "TotalNetworkBandwidth",
          "value": "6000"
        },
        {
          "key": "TotalVolumeBandwidth",
          "value": "2000"
        },
        {
          "key": "Vcpu",
          "value": "{architecture:amd64,count:4,manufacturer:intel,percentage:100}"
        },
        {
          "key": "VolumeAttachments",
          "value": "{device:{id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199-dpdgt},href:https://au-syd.iaas.cloud.ibm.com/v1/instances/02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6/volume_attachments/02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,id:02h7-fc17086e-3af1-45a6-af02-459cc2b8b199,name:riverside-fang-shampoo-affirm,volume:{crn:crn:v1:bluemix:public:is:au-syd-1:a/ab205347a7c3b57f09dabb32df178bcf::volume:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,href:https://au-syd.iaas.cloud.ibm.com/v1/volumes/r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,id:r026-01f48f90-c0ec-4014-9ac9-cb4b2db80737,name:runway-preshow-deed-truce,resource_type:volume}}"
        },
        {
          "key": "VolumeBandwidthQosMode",
          "value": "pooled"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:au-syd:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,href:https://au-syd.iaas.cloud.ibm.com/v1/vpcs/r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,id:r026-5b6a8f03-385e-44ef-b1bc-825c662ed931,name:tbans8fa21l4pkrjo0a0,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://au-syd.iaas.cloud.ibm.com/v1/regions/au-syd/zones/au-syd-1,name:au-syd-1}"
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
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "nodeIp": "159.23.102.221",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbt7k7116g5o02ekkgfm 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "159.23.98.218",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
        "err": null
      },
      {
        "infraId": "my06-infra101",
        "nodeId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "nodeIp": "159.23.94.110",
        "command": {
          "0": "uname -a"
        },
        "stdout": {},
        "stderr": {},
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
      "nodeGroup": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "nodeId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "output": "Linux tbt7k7116g5o02ekkgfm 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "159.23.102.221",
      "sshTest": "successful",
      "status": "success",
      "testOrder": 1,
      "userName": "cb-user"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "nodeGroup": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "nodeId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "output": "Linux tb1h7ijtv4sdjjt74cq4 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "159.23.98.218",
      "sshTest": "successful",
      "status": "success",
      "testOrder": 2,
      "userName": "cb-user"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "nodeGroup": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "nodeId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "output": "Linux tbfb0c8hrtma9tjabtm6 5.15.0-1100-ibm #103-Ubuntu SMP Mon Apr 20 14:53:14 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "159.23.94.110",
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

**Generated At:** 2026-06-02 12:01:20

**Namespace:** mig01

**Infra Name:** my06-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my06-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | IBM |
| **Target Region** | au-syd |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| nxf-2x2 | 2 | 2.0 | - | x86_64 |  | $0.0940 | 1 |
| bxf-2x8 | 2 | 8.0 | - | x86_64 |  | $0.1170 | 1 |
| bxf-4x16 | 4 | 16.0 | - | x86_64 |  | $0.2350 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| r026-c8e249d4-f148-4416-a3c6-555b7a02f67d | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | 02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457 | Running | 2 vCPU, 2.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-vnet-01<br>**Subnet:** my06-subnet-01<br>**Public IP:** 159.23.102.221<br>**Private IP:** 10.0.1.5<br>**SGs:** my06-sg-01<br>**SSH:** my06-sshkey-01 |
| my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | 02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4 | Running | 2 vCPU, 8.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-vnet-01<br>**Subnet:** my06-subnet-01<br>**Public IP:** 159.23.98.218<br>**Private IP:** 10.0.1.6<br>**SGs:** my06-sg-03<br>**SSH:** my06-sshkey-01 |
| my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | 02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6 | Running | 4 vCPU, 16.0 GiB | Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) (Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64)) | **VNet:** my06-vnet-01<br>**Subnet:** my06-subnet-01<br>**Public IP:** 159.23.94.110<br>**Private IP:** 10.0.1.4<br>**SGs:** my06-sg-02<br>**SSH:** my06-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my06-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my06-vnet-01 |
| **CSP VNet ID** | r026-5b6a8f03-385e-44ef-b1bc-825c662ed931 |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | ibm-au-syd |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my06-subnet-01 | 02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d | 10.0.1.0/24 | au-syd-1 |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my06-sshkey-01 | r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb |  | SHA256:TvBOM01LhEtmxXCF8izGk3MZUyQr9cHun/EEuzIgBLY |

### Security Groups

#### Security Group: my06-sg-01

| Property | Value |
|----------|-------|
| **Name** | my06-sg-01 |
| **CSP Security Group ID** | r026-cc7a1860-d3e6-4d83-a31f-753910673746 |
| **VNet** | my06-vnet-01 |
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
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my06-sg-02

| Property | Value |
|----------|-------|
| **Name** | my06-sg-02 |
| **CSP Security Group ID** | r026-09e47208-047d-4485-8bd2-19d7a50feb71 |
| **VNet** | my06-vnet-01 |
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
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my06-sg-03

| Property | Value |
|----------|-------|
| **Name** | my06-sg-03 |
| **CSP Security Group ID** | r026-6e6f8825-c77d-4255-af1a-3eeeaad54cbc |
| **VNet** | my06-vnet-01 |
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
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.4460 |
| **Per Day** | $10.70 |
| **Per Month (30 days)** | $321.12 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| IBM | au-syd | 3 | $0.4460 | $321.12 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | nxf-2x2 | $0.0940 | $67.68 |
| my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | bxf-2x8 | $0.1170 | $84.24 |
| my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | bxf-4x16 | $0.2350 | $169.20 |




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

*Report generated: 2026-06-02 12:01:25*

---

## 📊 Migration Summary

**Target Cloud:** IBM

**Target Region:** au-syd

**Namespace:** mig01 | **Infra ID:** my06-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $321.12 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 3 | ✅ Selected | bxf-2x8, bxf-4x16, nxf-2x2 |
| 3 | **VM OS Image** | 1 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my06-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my06-vnet-01) |
| 6 | **Security Group** | 3 security groups | ✅ Created | Total 52 rules in 3 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** 02h7_19d7fc00-e7df-4a3c-b2cc-fa3322e13457<br>**Label(sourceMachineId):** vm-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** 02h7_95a5c89f-123a-4b3a-844f-7df4b78274c4<br>**Label(sourceMachineId):** vm-ec288dd0-c6fa-8a49-2f60 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 |
| 3 | **VM Name:** my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** 02h7_9472bbc0-6de5-4bbe-af5f-cb0b60ba15c6<br>**Label(sourceMachineId):** vm-ec2d32b5-98fb-5a96-7913 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 |

---

## ⚙️ VM Specs

**Summary:** 3 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** nxf-2x2<br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2 | my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Spec ID:** bxf-2x8<br>**vCPUs:** 2<br>**Memory:** 8.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3 | my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Spec ID:** bxf-4x16<br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** r026-c8e249d4-f148-4416-a3c6-555b7a02f67d<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu Linux 22.04 LTS Jammy Jellyfish Minimal Install (amd64) | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my06-sg-01

**CSP ID:** r026-cc7a1860-d3e6-4d83-a31f-753910673746 | **VNet:** my06-vnet-01 | **Rules:** 14

**Assigned VMs:**

- **VM:** my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec268ed7-821e-9d73-e79f

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 2 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 3 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 4 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | inbound | TCP | 80 | 0.0.0.0/0 | inbound tcp 80 | Migrated from source |
| 7 | inbound | TCP | 443 | 0.0.0.0/0 | inbound tcp 443 | Migrated from source |
| 8 | inbound | TCP | 8080 | 0.0.0.0/0 | inbound tcp 8080 | Migrated from source |
| 9 | inbound | TCP | 9113 | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 10 | inbound | UDP | 9113 | 10.0.0.0/16 | inbound udp 9113 from 10.0.0.0/16 | Migrated from source |
| 11 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 12 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 13 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 14 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my06-sg-02

**CSP ID:** r026-09e47208-047d-4485-8bd2-19d7a50feb71 | **VNet:** my06-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec2d32b5-98fb-5a96-7913

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 2 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 3 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 4 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | inbound | TCP | 2049 | 0.0.0.0/0 | inbound tcp 2049 | Migrated from source |
| 7 | inbound | UDP | 2049 | 0.0.0.0/0 | inbound udp 2049 | Migrated from source |
| 8 | inbound | TCP | 111 | 0.0.0.0/0 | inbound tcp 111 | Migrated from source |
| 9 | inbound | UDP | 111 | 0.0.0.0/0 | inbound udp 111 | Migrated from source |
| 10 | inbound | TCP | 20048 | 10.0.0.0/16 | inbound tcp 20048 from 10.0.0.0/16 | Migrated from source |
| 11 | inbound | UDP | 20048 | 10.0.0.0/16 | inbound udp 20048 from 10.0.0.0/16 | Migrated from source |
| 12 | inbound | TCP | 32803 | 10.0.0.0/16 | inbound tcp 32803 from 10.0.0.0/16 | Migrated from source |
| 13 | inbound | UDP | 32803 | 10.0.0.0/16 | inbound udp 32803 from 10.0.0.0/16 | Migrated from source |
| 14 | inbound | TCP | 9100 | 10.0.0.0/16 | inbound tcp 9100 from 10.0.0.0/16 | Migrated from source |
| 15 | inbound | UDP | 9100 | 10.0.0.0/16 | inbound udp 9100 from 10.0.0.0/16 | Migrated from source |
| 16 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 17 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 18 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 19 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

### Security Group: my06-sg-03

**CSP ID:** r026-6e6f8825-c77d-4255-af1a-3eeeaad54cbc | **VNet:** my06-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** vm-ec288dd0-c6fa-8a49-2f60

**Security Rules:**

| No. | Direction | Protocol | Port | CIDR | Source Firewall Rule | Note |
|-----|-----------|----------|------|------|----------------------|------|
| 1 | inbound | ICMP |  | 0.0.0.0/0 | inbound icmp * | Migrated from source |
| 2 | inbound | UDP | 68 | 0.0.0.0/0 | inbound udp 68 | Migrated from source |
| 3 | inbound | UDP | 5353 | 0.0.0.0/0 | inbound udp 5353 | Migrated from source |
| 4 | inbound | UDP | 1900 | 0.0.0.0/0 | inbound udp 1900 | Migrated from source |
| 5 | inbound | TCP | 22 | 0.0.0.0/0 | inbound tcp 22 | Migrated from source |
| 6 | inbound | TCP | 3306 | 10.0.0.0/16 | inbound tcp 3306 from 10.0.0.0/16 | Migrated from source |
| 7 | inbound | UDP | 3306 | 10.0.0.0/16 | inbound udp 3306 from 10.0.0.0/16 | Migrated from source |
| 8 | inbound | TCP | 4567 | 10.0.0.0/16 | inbound tcp 4567 from 10.0.0.0/16 | Migrated from source |
| 9 | inbound | UDP | 4567 | 10.0.0.0/16 | inbound udp 4567 from 10.0.0.0/16 | Migrated from source |
| 10 | inbound | TCP | 4568 | 10.0.0.0/16 | inbound tcp 4568 from 10.0.0.0/16 | Migrated from source |
| 11 | inbound | UDP | 4568 | 10.0.0.0/16 | inbound udp 4568 from 10.0.0.0/16 | Migrated from source |
| 12 | inbound | TCP | 4444 | 10.0.0.0/16 | inbound tcp 4444 from 10.0.0.0/16 | Migrated from source |
| 13 | inbound | UDP | 4444 | 10.0.0.0/16 | inbound udp 4444 from 10.0.0.0/16 | Migrated from source |
| 14 | inbound | TCP | 9104 | 10.0.0.0/16 | inbound tcp 9104 from 10.0.0.0/16 | Migrated from source |
| 15 | inbound | UDP | 9104 | 10.0.0.0/16 | inbound udp 9104 from 10.0.0.0/16 | Migrated from source |
| 16 | inbound | ALL |  | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 17 | outbound | TCP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 18 | outbound | UDP | 1-65535 | 0.0.0.0/0 | - | Created by system |
| 19 | outbound | ALL |  | 0.0.0.0/0 | outbound * * | Migrated from source |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source node network information.

### VPC(VNet)

| No. | VPC(VNet) | CIDR Block |
|-----|-----------|------------|
| 1 | **Name:** my06-vnet-01<br>**ID:** r026-5b6a8f03-385e-44ef-b1bc-825c662ed931 | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my06-subnet-01<br>**ID:** 02h7-3a33398e-c6ca-4655-863a-99dbdad0a91d | 10.0.1.0/24 | my06-vnet-01 |

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
| 1 | my06-sshkey-01 | r026-19980fa4-1edc-4a64-9b7d-4a3d029ed2cb | SHA256:TvBOM01LhEtmxXCF8izGk3MZUyQr9cHun/EEuzIgBLY | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.4460 |
| Daily | $10.70 |
| Monthly | $321.12 |
| Yearly | $3853.44 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | nxf-2x2 | $67.68 | 21.1% |
| ip-10-0-1-221 (migrated) | bxf-4x16 | $169.20 | 52.7% |
| ip-10-0-1-138 (migrated) | bxf-2x8 | $84.24 | 26.2% |

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
  "message": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: my06-infra101)",
  "success": true
}
```

