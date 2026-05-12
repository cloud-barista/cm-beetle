# CM-Beetle test results for AZURE

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: imdl/v0.1.4+ (d99ac81)
- cm-model: unknown
- CB-Tumblebug: vunknown
- CB-Spider: vunknown
- CB-MapUI: vunknown
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: May 12, 2026
- Test Time: 18:43:58 KST
- Test Execution: 2026-05-12 18:43:58 KST

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

## Test result for AZURE

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 2.53s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m49.869s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 37ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 12ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 140ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 7.434s | Pass |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.355s | Pass |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m56.399s | Pass |

**Overall Result**: 9/9 tests passed ✅

**Total Duration**: 5m15.987516747s

*Test executed on May 12, 2026 at 18:43:58 KST (2026-05-12 18:43:58 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | partially-matched | Overall Match Rate: Min=51.2% Max=100.0% Avg=94.1% | VMs: 3 total, 2 matched, 1 acceptable",
      "targetCloud": {
        "csp": "azure",
        "region": "koreasouth"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2als_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
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
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b4as_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
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
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2as_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
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
        "postCommand": {
          "userName": "",
          "command": null
        },
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "azure-koreasouth",
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
        "connectionName": "azure-koreasouth",
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
          "id": "azure+koreasouth+standard_b2als_v2",
          "uid": "d7k8jatq1uu3ogms3t1g",
          "cspSpecName": "Standard_B2als_v2",
          "name": "azure+koreasouth+standard_b2als_v2",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
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
          "rootDiskSize": 0,
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
              "value": "Open,Targeted"
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
          "uid": "d7k8jatq1uu3ogms3t3g",
          "cspSpecName": "Standard_B4as_v2",
          "name": "azure+koreasouth+standard_b4as_v2",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
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
          "rootDiskSize": 0,
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
              "value": "Open,Targeted"
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
          "uid": "d7k8jatq1uu3ogms3t20",
          "cspSpecName": "Standard_B2as_v2",
          "name": "azure+koreasouth+standard_b2as_v2",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
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
          "rootDiskSize": 0,
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
              "value": "Open,Targeted"
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
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "uid": "d7k8lctq1uu3ogn5gg30",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.04.22 08:45:39 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "osDiskType": "default",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Location",
              "value": "australiacentral"
            },
            {
              "key": "Publisher",
              "value": "Canonical"
            },
            {
              "key": "Offer",
              "value": "0001-com-ubuntu-server-jammy-daily"
            },
            {
              "key": "SKU",
              "value": "22_04-daily-lts-gen2"
            },
            {
              "key": "Version",
              "value": "22.04.202604160"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy-daily/Skus/22_04-daily-lts-gen2/Versions/22.04.202604160"
            },
            {
              "key": "HyperVGeneration",
              "value": "V2"
            },
            {
              "key": "Features",
              "value": "SecurityType=TrustedLaunchSupported, IsAcceleratedNetworkSupported=True, DiskControllerTypes=SCSI, NVMe, IsHibernateSupported=True"
            },
            {
              "key": "FeatureCount",
              "value": "4"
            },
            {
              "key": "ImageDeprecationState",
              "value": "Active"
            }
          ],
          "systemLabel": "",
          "description": "",
          "commandHistory": null
        },
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "uid": "d7k8lctq1uu3ogn5gg20",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.04.22 08:45:39 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "osDiskType": "default",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Location",
              "value": "australiacentral"
            },
            {
              "key": "Publisher",
              "value": "Canonical"
            },
            {
              "key": "Offer",
              "value": "0001-com-ubuntu-server-jammy-daily"
            },
            {
              "key": "SKU",
              "value": "22_04-daily-lts"
            },
            {
              "key": "Version",
              "value": "22.04.202604160"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy-daily/Skus/22_04-daily-lts/Versions/22.04.202604160"
            },
            {
              "key": "HyperVGeneration",
              "value": "V1"
            },
            {
              "key": "Features",
              "value": "IsAcceleratedNetworkSupported=True, DiskControllerTypes=SCSI, IsHibernateSupported=True"
            },
            {
              "key": "FeatureCount",
              "value": "3"
            },
            {
              "key": "ImageDeprecationState",
              "value": "Active"
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
          "connectionName": "azure-koreasouth",
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
          "connectionName": "azure-koreasouth",
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
          "connectionName": "azure-koreasouth",
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
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=51.2% Max=100.0% Avg=94.1% | VMs: 3 total, 2 matched, 1 acceptable",
      "targetCloud": {
        "csp": "azure",
        "region": "koreasouth"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2ls_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
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
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b4s_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
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
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2s_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
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
        "postCommand": {
          "userName": "",
          "command": null
        },
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "azure-koreasouth",
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
        "connectionName": "azure-koreasouth",
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
          "id": "azure+koreasouth+standard_b2als_v2",
          "uid": "d7k8jatq1uu3ogms3t1g",
          "cspSpecName": "Standard_B2als_v2",
          "name": "azure+koreasouth+standard_b2als_v2",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
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
          "rootDiskSize": 0,
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
              "value": "Open,Targeted"
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
          "uid": "d7k8jatq1uu3ogms3t3g",
          "cspSpecName": "Standard_B4as_v2",
          "name": "azure+koreasouth+standard_b4as_v2",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
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
          "rootDiskSize": 0,
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
              "value": "Open,Targeted"
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
          "uid": "d7k8jatq1uu3ogms3t20",
          "cspSpecName": "Standard_B2as_v2",
          "name": "azure+koreasouth+standard_b2as_v2",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
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
          "rootDiskSize": 0,
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
              "value": "Open,Targeted"
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
        },
        {
          "id": "azure+koreasouth+standard_b2ls_v2",
          "uid": "d7k8jatq1uu3ogms3qjg",
          "cspSpecName": "Standard_B2ls_v2",
          "name": "azure+koreasouth+standard_b2ls_v2",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 3.90625,
          "costPerHour": 0.0478,
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
          "rootDiskSize": 0,
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
              "value": "Standard_B2ls_v2"
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
              "value": "Open,Targeted"
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
              "value": "True"
            },
            {
              "key": "LocationInfo_0_Location",
              "value": "KoreaSouth"
            },
            {
              "key": "Family",
              "value": "standardBsv2Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "B2ls_v2"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b4s_v2",
          "uid": "d7k8jatq1uu3ogms3qlg",
          "cspSpecName": "Standard_B4s_v2",
          "name": "azure+koreasouth+standard_b4s_v2",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 15.625,
          "costPerHour": 0.191,
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
          "rootDiskSize": 0,
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
              "value": "Standard_B4s_v2"
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
              "value": "Open,Targeted"
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
              "value": "True"
            },
            {
              "key": "LocationInfo_0_Location",
              "value": "KoreaSouth"
            },
            {
              "key": "Family",
              "value": "standardBsv2Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "B4s_v2"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b2s_v2",
          "uid": "d7k8jatq1uu3ogms3qk0",
          "cspSpecName": "Standard_B2s_v2",
          "name": "azure+koreasouth+standard_b2s_v2",
          "namespace": "system",
          "connectionName": "azure-koreasouth",
          "providerName": "azure",
          "regionName": "koreasouth",
          "regionLatitude": 35.1796,
          "regionLongitude": 129.0756,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 7.8125,
          "costPerHour": 0.0957,
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
          "rootDiskSize": 0,
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
              "value": "Standard_B2s_v2"
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
              "value": "Open,Targeted"
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
              "value": "True"
            },
            {
              "key": "LocationInfo_0_Location",
              "value": "KoreaSouth"
            },
            {
              "key": "Family",
              "value": "standardBsv2Family"
            },
            {
              "key": "Tier",
              "value": "Standard"
            },
            {
              "key": "Size",
              "value": "B2s_v2"
            },
            {
              "key": "ResourceType",
              "value": "virtualMachines"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "uid": "d7k8lctq1uu3ogn5gg30",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.04.22 08:45:39 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "osDiskType": "default",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Location",
              "value": "australiacentral"
            },
            {
              "key": "Publisher",
              "value": "Canonical"
            },
            {
              "key": "Offer",
              "value": "0001-com-ubuntu-server-jammy-daily"
            },
            {
              "key": "SKU",
              "value": "22_04-daily-lts-gen2"
            },
            {
              "key": "Version",
              "value": "22.04.202604160"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy-daily/Skus/22_04-daily-lts-gen2/Versions/22.04.202604160"
            },
            {
              "key": "HyperVGeneration",
              "value": "V2"
            },
            {
              "key": "Features",
              "value": "SecurityType=TrustedLaunchSupported, IsAcceleratedNetworkSupported=True, DiskControllerTypes=SCSI, NVMe, IsHibernateSupported=True"
            },
            {
              "key": "FeatureCount",
              "value": "4"
            },
            {
              "key": "ImageDeprecationState",
              "value": "Active"
            }
          ],
          "systemLabel": "",
          "description": "",
          "commandHistory": null
        },
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "uid": "d7k8lctq1uu3ogn5gg20",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.04.22 08:45:39 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "osDiskType": "default",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Location",
              "value": "australiacentral"
            },
            {
              "key": "Publisher",
              "value": "Canonical"
            },
            {
              "key": "Offer",
              "value": "0001-com-ubuntu-server-jammy-daily"
            },
            {
              "key": "SKU",
              "value": "22_04-daily-lts"
            },
            {
              "key": "Version",
              "value": "22.04.202604160"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy-daily/Skus/22_04-daily-lts/Versions/22.04.202604160"
            },
            {
              "key": "HyperVGeneration",
              "value": "V1"
            },
            {
              "key": "Features",
              "value": "IsAcceleratedNetworkSupported=True, DiskControllerTypes=SCSI, IsHibernateSupported=True"
            },
            {
              "key": "FeatureCount",
              "value": "3"
            },
            {
              "key": "ImageDeprecationState",
              "value": "Active"
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
          "connectionName": "azure-koreasouth",
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
          "connectionName": "azure-koreasouth",
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
          "connectionName": "azure-koreasouth",
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
  "id": "my02-infra101",
  "uid": "tbd81fd366aji65qtg7s9g",
  "name": "my02-infra101",
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
    "sys.id": "my02-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my02-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbd81fd366aji65qtg7s9g"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbd81fd3e6aji65qtg7sag",
      "cspResourceName": "tbd81fd3e6aji65qtg7sag",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag",
      "name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-05-12 09:45:45",
      "label": {
        "createdBy": "tbd81fd3e6aji65qtg7sag",
        "keypair": "tbd81fcu66aji65qtg7s7g",
        "publicip": "tbd81fd3e6aji65qtg7sag-6948-PublicIP",
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-12 09:45:45",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag",
        "sys.cspResourceName": "tbd81fd3e6aji65qtg7sag",
        "sys.id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd81fd3e6aji65qtg7sag",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.214.4.247",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 50,
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0",
      "networkInterface": "tbd81fd3e6aji65qtg7sag-97625-VNic",
      "securityGroupIds": [
        "my02-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOlhZnTKot5CpvRf0YzyGcCOMZuXeAMAgdZPJBjMssWEV3mgDw0DIAaZL1vj8YsjarGo8/mGHmt0/jtU0XI2DU0=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:NLSKcyS77IrfjQOPBXWdHjfqLW8dVJipgwWxHgZQYe0",
        "firstUsedAt": "2026-05-12T09:45:55Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-12T09:45:55Z",
          "completedTime": "2026-05-12T09:45:59Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd81fd3e6aji65qtg7sag 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd81fd3e6aji65qtg7sag-97625-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd81fd3e6aji65qtg7sag,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDUkSh+Hm49mSp1aeExqrCuh1wK6S64rPm4ONnlNHe088Sxf6kcgqyR5b+poXYPT2iUjyQBqWdiIqhHCd1iTxz0h0y92l8DmRS00ez9QV6IB3CV1ZKiBdP3s2S5btgEEDIVbnzmKoMg9yU67MtZE3ER9ePK4HzT/Ps0DNMbsyseEILG7q1kGauxvySBoaHb7x+Dai/879pIFLN2tHfYzakdAsnze1iH1NEzbnx+cfJb+X0q2EWU0cnjGJpGxCOsJ2jq4UuS7/mSBxQzPxZE0a34pBf3d6Jut8lrS/NMiWSrnAIX6ZHrYOrx1aJhbdkp2OXaYZyIrxUtlHxS75UlNjPAkTmVBrpySw5oeiRq6Y6Qxj7+tfQ+xrOQr2NtTqlKcYo+bfvYo81P0s0q8+Z9h4a06U+I+lqLvQ8zPbCpqeI2r+49moYdQK3lok2PhbkIDXoBp+NacfhIKQMJW0VZ74kdmolI9edR9ZlTyWvXMIM4jA4WSuDhGU0FDaGJzJzRAHwXj0EZNy37R+M9v54rYdcaZzqmnoFwQ+6hPF3FD4anG1yhI5Z+gIvs3aZ64HEeKrNOZEXR1pLmg17N8o2QNZyKenjm0LPIPpfYxRhJEzYoVte4ZiVqbJj+MSoNd7ReAYhesL8iSUFw85QyDZdq5O1tQokGL4RqU2/nV7fhpiSPcw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202604160,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202604160},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd81fd3e6aji65qtg7sag_OsDisk_1_8475ebb7b2fc4065a10ea9965259015e,storageAccountType:Premium_LRS},name:tbd81fd3e6aji65qtg7sag_OsDisk_1_8475ebb7b2fc4065a10ea9965259015e,osType:Linux}},timeCreated:2026-05-12T09:44:53.6437916Z,vmId:1b543ee0-8c3c-44e9-9d8b-573c0e3fc6c9}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd81fd3e6aji65qtg7sag,keypair:tbd81fcu66aji65qtg7s7g,publicip:tbd81fd3e6aji65qtg7sag-6948-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag"
        },
        {
          "key": "Name",
          "value": "tbd81fd3e6aji65qtg7sag"
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
      "uid": "tbd81fd3e6aji65qtg7scg",
      "cspResourceName": "tbd81fd3e6aji65qtg7scg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg",
      "name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-05-12 09:45:48",
      "label": {
        "createdBy": "tbd81fd3e6aji65qtg7scg",
        "keypair": "tbd81fcu66aji65qtg7s7g",
        "publicip": "tbd81fd3e6aji65qtg7scg-43856-PublicIP",
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-12 09:45:48",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg",
        "sys.cspResourceName": "tbd81fd3e6aji65qtg7scg",
        "sys.id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd81fd3e6aji65qtg7scg",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "52.231.199.161",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 50,
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0",
      "networkInterface": "tbd81fd3e6aji65qtg7scg-38274-VNic",
      "securityGroupIds": [
        "my02-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGNC3nphY/v0+xM4rigF6rS6/BryTA6wUKcYifFO+GrlgOg05dR7/deJP4WKbiryKR++SIuzGvpIHF7XfWkCxbs=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:Y2gaNklAb4QHPPMwKNRhTABm2kr7c8p0/zvEwSdvNKU",
        "firstUsedAt": "2026-05-12T09:45:58Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-12T09:45:55Z",
          "completedTime": "2026-05-12T09:46:00Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd81fd3e6aji65qtg7scg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd81fd3e6aji65qtg7scg-38274-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd81fd3e6aji65qtg7scg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDUkSh+Hm49mSp1aeExqrCuh1wK6S64rPm4ONnlNHe088Sxf6kcgqyR5b+poXYPT2iUjyQBqWdiIqhHCd1iTxz0h0y92l8DmRS00ez9QV6IB3CV1ZKiBdP3s2S5btgEEDIVbnzmKoMg9yU67MtZE3ER9ePK4HzT/Ps0DNMbsyseEILG7q1kGauxvySBoaHb7x+Dai/879pIFLN2tHfYzakdAsnze1iH1NEzbnx+cfJb+X0q2EWU0cnjGJpGxCOsJ2jq4UuS7/mSBxQzPxZE0a34pBf3d6Jut8lrS/NMiWSrnAIX6ZHrYOrx1aJhbdkp2OXaYZyIrxUtlHxS75UlNjPAkTmVBrpySw5oeiRq6Y6Qxj7+tfQ+xrOQr2NtTqlKcYo+bfvYo81P0s0q8+Z9h4a06U+I+lqLvQ8zPbCpqeI2r+49moYdQK3lok2PhbkIDXoBp+NacfhIKQMJW0VZ74kdmolI9edR9ZlTyWvXMIM4jA4WSuDhGU0FDaGJzJzRAHwXj0EZNy37R+M9v54rYdcaZzqmnoFwQ+6hPF3FD4anG1yhI5Z+gIvs3aZ64HEeKrNOZEXR1pLmg17N8o2QNZyKenjm0LPIPpfYxRhJEzYoVte4ZiVqbJj+MSoNd7ReAYhesL8iSUFw85QyDZdq5O1tQokGL4RqU2/nV7fhpiSPcw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202604160,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202604160},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd81fd3e6aji65qtg7scg_OsDisk_1_9fca1cfb0b9a4d3ead6e5403bc3b9dac,storageAccountType:Premium_LRS},name:tbd81fd3e6aji65qtg7scg_OsDisk_1_9fca1cfb0b9a4d3ead6e5403bc3b9dac,osType:Linux}},timeCreated:2026-05-12T09:44:57.9968299Z,vmId:5e905251-d043-46c4-814e-489b8c8fe657}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd81fd3e6aji65qtg7scg,keypair:tbd81fcu66aji65qtg7s7g,publicip:tbd81fd3e6aji65qtg7scg-43856-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg"
        },
        {
          "key": "Name",
          "value": "tbd81fd3e6aji65qtg7scg"
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
      "uid": "tbd81fd3e6aji65qtg7sbg",
      "cspResourceName": "tbd81fd3e6aji65qtg7sbg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg",
      "name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-05-12 09:45:47",
      "label": {
        "createdBy": "tbd81fd3e6aji65qtg7sbg",
        "keypair": "tbd81fcu66aji65qtg7s7g",
        "publicip": "tbd81fd3e6aji65qtg7sbg-88014-PublicIP",
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-12 09:45:47",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg",
        "sys.cspResourceName": "tbd81fd3e6aji65qtg7sbg",
        "sys.id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd81fd3e6aji65qtg7sbg",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "52.231.199.157",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 50,
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0",
      "networkInterface": "tbd81fd3e6aji65qtg7sbg-96100-VNic",
      "securityGroupIds": [
        "my02-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAN6J4eynmqAFw1iIOLTlcDd2Nj3VXVk/ntLQad0HFdZmLHM1XlpHkgbQBQ0ZbT7WdXlCNwjQFUTJUz2Peu6tt8=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:oURuT7phD9rhemFegoCfIhW9JIwJd97nUZcaLxVND60",
        "firstUsedAt": "2026-05-12T09:45:58Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-12T09:45:55Z",
          "completedTime": "2026-05-12T09:46:00Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd81fd3e6aji65qtg7sbg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd81fd3e6aji65qtg7sbg-96100-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd81fd3e6aji65qtg7sbg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDUkSh+Hm49mSp1aeExqrCuh1wK6S64rPm4ONnlNHe088Sxf6kcgqyR5b+poXYPT2iUjyQBqWdiIqhHCd1iTxz0h0y92l8DmRS00ez9QV6IB3CV1ZKiBdP3s2S5btgEEDIVbnzmKoMg9yU67MtZE3ER9ePK4HzT/Ps0DNMbsyseEILG7q1kGauxvySBoaHb7x+Dai/879pIFLN2tHfYzakdAsnze1iH1NEzbnx+cfJb+X0q2EWU0cnjGJpGxCOsJ2jq4UuS7/mSBxQzPxZE0a34pBf3d6Jut8lrS/NMiWSrnAIX6ZHrYOrx1aJhbdkp2OXaYZyIrxUtlHxS75UlNjPAkTmVBrpySw5oeiRq6Y6Qxj7+tfQ+xrOQr2NtTqlKcYo+bfvYo81P0s0q8+Z9h4a06U+I+lqLvQ8zPbCpqeI2r+49moYdQK3lok2PhbkIDXoBp+NacfhIKQMJW0VZ74kdmolI9edR9ZlTyWvXMIM4jA4WSuDhGU0FDaGJzJzRAHwXj0EZNy37R+M9v54rYdcaZzqmnoFwQ+6hPF3FD4anG1yhI5Z+gIvs3aZ64HEeKrNOZEXR1pLmg17N8o2QNZyKenjm0LPIPpfYxRhJEzYoVte4ZiVqbJj+MSoNd7ReAYhesL8iSUFw85QyDZdq5O1tQokGL4RqU2/nV7fhpiSPcw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202604160,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202604160},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd81fd3e6aji65qtg7sbg_OsDisk_1_aef0e02e4b6a4ebca659642181e7ee12,storageAccountType:Premium_LRS},name:tbd81fd3e6aji65qtg7sbg_OsDisk_1_aef0e02e4b6a4ebca659642181e7ee12,osType:Linux}},timeCreated:2026-05-12T09:44:57.0636694Z,vmId:7f4cdcdf-b136-45ae-8898-b261faa81b51}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd81fd3e6aji65qtg7sbg,keypair:tbd81fcu66aji65qtg7s7g,publicip:tbd81fd3e6aji65qtg7sbg-88014-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg"
        },
        {
          "key": "Name",
          "value": "tbd81fd3e6aji65qtg7sbg"
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
        "nodeIp": "20.214.4.247",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd81fd3e6aji65qtg7sag 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my02-infra101",
        "nodeId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "52.231.199.161",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd81fd3e6aji65qtg7scg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my02-infra101",
        "nodeId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "nodeIp": "52.231.199.157",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd81fd3e6aji65qtg7sbg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "tbd81fctu6aji65qtg7s2g",
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
        "sys.uid": "tbd81fctu6aji65qtg7s2g"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbd81fctu6aji65qtg7s3g",
          "cspResourceName": "tbd81fctu6aji65qtg7s3g",
          "cspResourceId": "i-0951db77036789be7",
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
          "createdTime": "2026-05-12 09:44:47",
          "label": {
            "Name": "tbd81fctu6aji65qtg7s3g",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-05-12 09:44:47",
            "sys.cspResourceId": "i-0951db77036789be7",
            "sys.cspResourceName": "tbd81fctu6aji65qtg7s3g",
            "sys.id": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tbd81fctu6aji65qtg7s3g",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "43.203.118.191",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.221",
          "privateDNS": "ip-10-0-1-221.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 50,
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
          "imageId": "ami-08a1c21841a4c7a5f",
          "cspImageName": "ami-08a1c21841a4c7a5f",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-08a1c21841a4c7a5f",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260410"
          },
          "vNetId": "my01-vnet-01",
          "cspVNetId": "vpc-03418b1c97bcb7bb0",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-0e40174384d5575ff",
          "networkInterface": "eni-attach-015d05102e00d5423",
          "securityGroupIds": [
            "my01-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbd81fcs66aji65qtg7s0g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBMTCLbkv5mpN3rODVz7y/91K05G7O+Xdi/lajmxBMXj2gKUBjIq5vvtQL6zkmW76hKyt5/fAYnFnNnaVzTsCL44=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:Q1MQ+W69Z/sEU8y0ct5FT3JPvnyQX4KOzS4GPmo7UMc",
            "firstUsedAt": "2026-05-12T09:44:56Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-12T09:44:55Z",
              "completedTime": "2026-05-12T09:44:58Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-221 6.8.0-1051-aws #54~22.04.1-Ubuntu SMP Wed Mar 25 15:41:00 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-05-12T09:44:25Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-03efa0d6f0567f705}}"
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
              "value": "FFCD8649-19EE-4627-81E5-C849E26E75B9"
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
              "value": "ami-08a1c21841a4c7a5f"
            },
            {
              "key": "InstanceId",
              "value": "i-0951db77036789be7"
            },
            {
              "key": "InstanceType",
              "value": "t3a.small"
            },
            {
              "key": "KeyName",
              "value": "tbd81fcs66aji65qtg7s0g"
            },
            {
              "key": "LaunchTime",
              "value": "2026-05-12T09:44:24Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.118.191},Attachment:{AttachTime:2026-05-12T09:44:24Z,AttachmentId:eni-attach-015d05102e00d5423,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0df249c4f598c135c,GroupName:tbd81fcse6aji65qtg7s10}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:6a:7a:ff:02:23,NetworkInterfaceId:eni-0e23cf5cf082e7929,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.221,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.118.191},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.221}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0e40174384d5575ff,VpcId:vpc-03418b1c97bcb7bb0}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-221.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.221"
            },
            {
              "key": "PublicIpAddress",
              "value": "43.203.118.191"
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
              "value": "{GroupId:sg-0df249c4f598c135c,GroupName:tbd81fcse6aji65qtg7s10}"
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
              "value": "subnet-0e40174384d5575ff"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbd81fctu6aji65qtg7s3g}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-03418b1c97bcb7bb0"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tbd81fctu6aji65qtg7s5g",
          "cspResourceName": "tbd81fctu6aji65qtg7s5g",
          "cspResourceId": "i-0e4676d603cfb7af7",
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
          "createdTime": "2026-05-12 09:44:49",
          "label": {
            "Name": "tbd81fctu6aji65qtg7s5g",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-05-12 09:44:49",
            "sys.cspResourceId": "i-0e4676d603cfb7af7",
            "sys.cspResourceName": "tbd81fctu6aji65qtg7s5g",
            "sys.id": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tbd81fctu6aji65qtg7s5g",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "3.38.94.194",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.193",
          "privateDNS": "ip-10-0-1-193.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 50,
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
          "imageId": "ami-08a1c21841a4c7a5f",
          "cspImageName": "ami-08a1c21841a4c7a5f",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-08a1c21841a4c7a5f",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260410"
          },
          "vNetId": "my01-vnet-01",
          "cspVNetId": "vpc-03418b1c97bcb7bb0",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-0e40174384d5575ff",
          "networkInterface": "eni-attach-07537aa2a55e39358",
          "securityGroupIds": [
            "my01-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbd81fcs66aji65qtg7s0g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBHwzgi4IZgpbMDZPgXLWdcBDd/jtIJrfTBl06XnJd+JVDbQbW7E9T4BN9h/YYZAQT995HFdjkTTuiVVgLDUCnc=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:8DCFGUDqCdW+VXKPkvD16/aaD0h1iupZHGlKlrsphXE",
            "firstUsedAt": "2026-05-12T09:44:56Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-05-12T09:44:55Z",
              "completedTime": "2026-05-12T09:45:14Z",
              "elapsedTime": 19,
              "resultSummary": "Command execution failed",
              "errorMessage": "failed to establish SSH connection to target host: ssh: handshake failed: EOF"
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-05-12T09:44:29Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-04425355d7c4126e6}}"
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
              "value": "1FBED08A-500E-4330-A49E-EB55E9C8010C"
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
              "value": "ami-08a1c21841a4c7a5f"
            },
            {
              "key": "InstanceId",
              "value": "i-0e4676d603cfb7af7"
            },
            {
              "key": "InstanceType",
              "value": "t3a.large"
            },
            {
              "key": "KeyName",
              "value": "tbd81fcs66aji65qtg7s0g"
            },
            {
              "key": "LaunchTime",
              "value": "2026-05-12T09:44:29Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.38.94.194},Attachment:{AttachTime:2026-05-12T09:44:29Z,AttachmentId:eni-attach-07537aa2a55e39358,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-07dddf2bbd13b61bb,GroupName:tbd81fcte6aji65qtg7s20}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:b8:c1:16:63:29,NetworkInterfaceId:eni-020186ae30aa051cd,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.193,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.38.94.194},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.193}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0e40174384d5575ff,VpcId:vpc-03418b1c97bcb7bb0}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-193.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.193"
            },
            {
              "key": "PublicIpAddress",
              "value": "3.38.94.194"
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
              "value": "{GroupId:sg-07dddf2bbd13b61bb,GroupName:tbd81fcte6aji65qtg7s20}"
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
              "value": "subnet-0e40174384d5575ff"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbd81fctu6aji65qtg7s5g}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-03418b1c97bcb7bb0"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbd81fctu6aji65qtg7s4g",
          "cspResourceName": "tbd81fctu6aji65qtg7s4g",
          "cspResourceId": "i-0bfdfe83fdd3bd1d6",
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
          "createdTime": "2026-05-12 09:44:46",
          "label": {
            "Name": "tbd81fctu6aji65qtg7s4g",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-05-12 09:44:46",
            "sys.cspResourceId": "i-0bfdfe83fdd3bd1d6",
            "sys.cspResourceName": "tbd81fctu6aji65qtg7s4g",
            "sys.id": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tbd81fctu6aji65qtg7s4g",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "3.36.99.199",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.174",
          "privateDNS": "ip-10-0-1-174.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": 50,
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
          "imageId": "ami-08a1c21841a4c7a5f",
          "cspImageName": "ami-08a1c21841a4c7a5f",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-08a1c21841a4c7a5f",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260410"
          },
          "vNetId": "my01-vnet-01",
          "cspVNetId": "vpc-03418b1c97bcb7bb0",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-0e40174384d5575ff",
          "networkInterface": "eni-attach-07760d189c761c505",
          "securityGroupIds": [
            "my01-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbd81fcs66aji65qtg7s0g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBH3tBbd6R8+4iT1y1lp+1W8vcxSluVVqO7KG8cTazIsy2v3zEI+9fWte5DLnYyhHS7ZHoXQvgN+qPmfIXCeH8l0=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:Lugdl/PqgyoITxlK6SrXaAe3qwm1yusTEDe28CHQmtY",
            "firstUsedAt": "2026-05-12T09:44:55Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-12T09:44:55Z",
              "completedTime": "2026-05-12T09:44:57Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-174 6.8.0-1051-aws #54~22.04.1-Ubuntu SMP Wed Mar 25 15:41:00 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-05-12T09:44:28Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0772b3222158e12f4}}"
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
              "value": "88EF160F-F24C-439B-8CD6-D4326CC7CCB4"
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
              "value": "ami-08a1c21841a4c7a5f"
            },
            {
              "key": "InstanceId",
              "value": "i-0bfdfe83fdd3bd1d6"
            },
            {
              "key": "InstanceType",
              "value": "t3a.xlarge"
            },
            {
              "key": "KeyName",
              "value": "tbd81fcs66aji65qtg7s0g"
            },
            {
              "key": "LaunchTime",
              "value": "2026-05-12T09:44:27Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.36.99.199},Attachment:{AttachTime:2026-05-12T09:44:27Z,AttachmentId:eni-attach-07760d189c761c505,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-095f9ab3806e7754b,GroupName:tbd81fcsu6aji65qtg7s1g}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:b9:0a:1b:3a:01,NetworkInterfaceId:eni-0b320aca826247e1d,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.174,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.36.99.199},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.174}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0e40174384d5575ff,VpcId:vpc-03418b1c97bcb7bb0}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-174.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.174"
            },
            {
              "key": "PublicIpAddress",
              "value": "3.36.99.199"
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
              "value": "{GroupId:sg-095f9ab3806e7754b,GroupName:tbd81fcsu6aji65qtg7s1g}"
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
              "value": "subnet-0e40174384d5575ff"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbd81fctu6aji65qtg7s4g}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-03418b1c97bcb7bb0"
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
            "nodeId": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "3.36.99.199",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-174 6.8.0-1051-aws #54~22.04.1-Ubuntu SMP Wed Mar 25 15:41:00 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "43.203.118.191",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-221 6.8.0-1051-aws #54~22.04.1-Ubuntu SMP Wed Mar 25 15:41:00 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "3.38.94.194",
            "command": {
              "0": "uname -a"
            },
            "stdout": {},
            "stderr": {},
            "err": null
          }
        ]
      }
    },
    {
      "resourceType": "infra",
      "id": "my02-infra101",
      "uid": "tbd81fd366aji65qtg7s9g",
      "name": "my02-infra101",
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
        "sys.id": "my02-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbd81fd366aji65qtg7s9g"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbd81fd3e6aji65qtg7sag",
          "cspResourceName": "tbd81fd3e6aji65qtg7sag",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag",
          "name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-05-12 09:45:45",
          "label": {
            "createdBy": "tbd81fd3e6aji65qtg7sag",
            "keypair": "tbd81fcu66aji65qtg7s7g",
            "publicip": "tbd81fd3e6aji65qtg7sag-6948-PublicIP",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-05-12 09:45:45",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag",
            "sys.cspResourceName": "tbd81fd3e6aji65qtg7sag",
            "sys.id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbd81fd3e6aji65qtg7sag",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.4.247",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 50,
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0",
          "networkInterface": "tbd81fd3e6aji65qtg7sag-97625-VNic",
          "securityGroupIds": [
            "my02-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOlhZnTKot5CpvRf0YzyGcCOMZuXeAMAgdZPJBjMssWEV3mgDw0DIAaZL1vj8YsjarGo8/mGHmt0/jtU0XI2DU0=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:NLSKcyS77IrfjQOPBXWdHjfqLW8dVJipgwWxHgZQYe0",
            "firstUsedAt": "2026-05-12T09:45:55Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-12T09:45:55Z",
              "completedTime": "2026-05-12T09:45:59Z",
              "elapsedTime": 4,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbd81fd3e6aji65qtg7sag 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd81fd3e6aji65qtg7sag-97625-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd81fd3e6aji65qtg7sag,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDUkSh+Hm49mSp1aeExqrCuh1wK6S64rPm4ONnlNHe088Sxf6kcgqyR5b+poXYPT2iUjyQBqWdiIqhHCd1iTxz0h0y92l8DmRS00ez9QV6IB3CV1ZKiBdP3s2S5btgEEDIVbnzmKoMg9yU67MtZE3ER9ePK4HzT/Ps0DNMbsyseEILG7q1kGauxvySBoaHb7x+Dai/879pIFLN2tHfYzakdAsnze1iH1NEzbnx+cfJb+X0q2EWU0cnjGJpGxCOsJ2jq4UuS7/mSBxQzPxZE0a34pBf3d6Jut8lrS/NMiWSrnAIX6ZHrYOrx1aJhbdkp2OXaYZyIrxUtlHxS75UlNjPAkTmVBrpySw5oeiRq6Y6Qxj7+tfQ+xrOQr2NtTqlKcYo+bfvYo81P0s0q8+Z9h4a06U+I+lqLvQ8zPbCpqeI2r+49moYdQK3lok2PhbkIDXoBp+NacfhIKQMJW0VZ74kdmolI9edR9ZlTyWvXMIM4jA4WSuDhGU0FDaGJzJzRAHwXj0EZNy37R+M9v54rYdcaZzqmnoFwQ+6hPF3FD4anG1yhI5Z+gIvs3aZ64HEeKrNOZEXR1pLmg17N8o2QNZyKenjm0LPIPpfYxRhJEzYoVte4ZiVqbJj+MSoNd7ReAYhesL8iSUFw85QyDZdq5O1tQokGL4RqU2/nV7fhpiSPcw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202604160,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202604160},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd81fd3e6aji65qtg7sag_OsDisk_1_8475ebb7b2fc4065a10ea9965259015e,storageAccountType:Premium_LRS},name:tbd81fd3e6aji65qtg7sag_OsDisk_1_8475ebb7b2fc4065a10ea9965259015e,osType:Linux}},timeCreated:2026-05-12T09:44:53.6437916Z,vmId:1b543ee0-8c3c-44e9-9d8b-573c0e3fc6c9}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbd81fd3e6aji65qtg7sag,keypair:tbd81fcu66aji65qtg7s7g,publicip:tbd81fd3e6aji65qtg7sag-6948-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag"
            },
            {
              "key": "Name",
              "value": "tbd81fd3e6aji65qtg7sag"
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
          "uid": "tbd81fd3e6aji65qtg7scg",
          "cspResourceName": "tbd81fd3e6aji65qtg7scg",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg",
          "name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
          "createdTime": "2026-05-12 09:45:48",
          "label": {
            "createdBy": "tbd81fd3e6aji65qtg7scg",
            "keypair": "tbd81fcu66aji65qtg7s7g",
            "publicip": "tbd81fd3e6aji65qtg7scg-43856-PublicIP",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-05-12 09:45:48",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg",
            "sys.cspResourceName": "tbd81fd3e6aji65qtg7scg",
            "sys.id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbd81fd3e6aji65qtg7scg",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "52.231.199.161",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.6",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 50,
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0",
          "networkInterface": "tbd81fd3e6aji65qtg7scg-38274-VNic",
          "securityGroupIds": [
            "my02-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGNC3nphY/v0+xM4rigF6rS6/BryTA6wUKcYifFO+GrlgOg05dR7/deJP4WKbiryKR++SIuzGvpIHF7XfWkCxbs=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:Y2gaNklAb4QHPPMwKNRhTABm2kr7c8p0/zvEwSdvNKU",
            "firstUsedAt": "2026-05-12T09:45:58Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-12T09:45:55Z",
              "completedTime": "2026-05-12T09:46:00Z",
              "elapsedTime": 5,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbd81fd3e6aji65qtg7scg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd81fd3e6aji65qtg7scg-38274-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd81fd3e6aji65qtg7scg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDUkSh+Hm49mSp1aeExqrCuh1wK6S64rPm4ONnlNHe088Sxf6kcgqyR5b+poXYPT2iUjyQBqWdiIqhHCd1iTxz0h0y92l8DmRS00ez9QV6IB3CV1ZKiBdP3s2S5btgEEDIVbnzmKoMg9yU67MtZE3ER9ePK4HzT/Ps0DNMbsyseEILG7q1kGauxvySBoaHb7x+Dai/879pIFLN2tHfYzakdAsnze1iH1NEzbnx+cfJb+X0q2EWU0cnjGJpGxCOsJ2jq4UuS7/mSBxQzPxZE0a34pBf3d6Jut8lrS/NMiWSrnAIX6ZHrYOrx1aJhbdkp2OXaYZyIrxUtlHxS75UlNjPAkTmVBrpySw5oeiRq6Y6Qxj7+tfQ+xrOQr2NtTqlKcYo+bfvYo81P0s0q8+Z9h4a06U+I+lqLvQ8zPbCpqeI2r+49moYdQK3lok2PhbkIDXoBp+NacfhIKQMJW0VZ74kdmolI9edR9ZlTyWvXMIM4jA4WSuDhGU0FDaGJzJzRAHwXj0EZNy37R+M9v54rYdcaZzqmnoFwQ+6hPF3FD4anG1yhI5Z+gIvs3aZ64HEeKrNOZEXR1pLmg17N8o2QNZyKenjm0LPIPpfYxRhJEzYoVte4ZiVqbJj+MSoNd7ReAYhesL8iSUFw85QyDZdq5O1tQokGL4RqU2/nV7fhpiSPcw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202604160,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202604160},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd81fd3e6aji65qtg7scg_OsDisk_1_9fca1cfb0b9a4d3ead6e5403bc3b9dac,storageAccountType:Premium_LRS},name:tbd81fd3e6aji65qtg7scg_OsDisk_1_9fca1cfb0b9a4d3ead6e5403bc3b9dac,osType:Linux}},timeCreated:2026-05-12T09:44:57.9968299Z,vmId:5e905251-d043-46c4-814e-489b8c8fe657}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbd81fd3e6aji65qtg7scg,keypair:tbd81fcu66aji65qtg7s7g,publicip:tbd81fd3e6aji65qtg7scg-43856-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg"
            },
            {
              "key": "Name",
              "value": "tbd81fd3e6aji65qtg7scg"
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
          "uid": "tbd81fd3e6aji65qtg7sbg",
          "cspResourceName": "tbd81fd3e6aji65qtg7sbg",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg",
          "name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
          "createdTime": "2026-05-12 09:45:47",
          "label": {
            "createdBy": "tbd81fd3e6aji65qtg7sbg",
            "keypair": "tbd81fcu66aji65qtg7s7g",
            "publicip": "tbd81fd3e6aji65qtg7sbg-88014-PublicIP",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-05-12 09:45:47",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg",
            "sys.cspResourceName": "tbd81fd3e6aji65qtg7sbg",
            "sys.id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbd81fd3e6aji65qtg7sbg",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "52.231.199.157",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.5",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": 50,
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0",
          "networkInterface": "tbd81fd3e6aji65qtg7sbg-96100-VNic",
          "securityGroupIds": [
            "my02-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAN6J4eynmqAFw1iIOLTlcDd2Nj3VXVk/ntLQad0HFdZmLHM1XlpHkgbQBQ0ZbT7WdXlCNwjQFUTJUz2Peu6tt8=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:oURuT7phD9rhemFegoCfIhW9JIwJd97nUZcaLxVND60",
            "firstUsedAt": "2026-05-12T09:45:58Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-12T09:45:55Z",
              "completedTime": "2026-05-12T09:46:00Z",
              "elapsedTime": 5,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbd81fd3e6aji65qtg7sbg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd81fd3e6aji65qtg7sbg-96100-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd81fd3e6aji65qtg7sbg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDUkSh+Hm49mSp1aeExqrCuh1wK6S64rPm4ONnlNHe088Sxf6kcgqyR5b+poXYPT2iUjyQBqWdiIqhHCd1iTxz0h0y92l8DmRS00ez9QV6IB3CV1ZKiBdP3s2S5btgEEDIVbnzmKoMg9yU67MtZE3ER9ePK4HzT/Ps0DNMbsyseEILG7q1kGauxvySBoaHb7x+Dai/879pIFLN2tHfYzakdAsnze1iH1NEzbnx+cfJb+X0q2EWU0cnjGJpGxCOsJ2jq4UuS7/mSBxQzPxZE0a34pBf3d6Jut8lrS/NMiWSrnAIX6ZHrYOrx1aJhbdkp2OXaYZyIrxUtlHxS75UlNjPAkTmVBrpySw5oeiRq6Y6Qxj7+tfQ+xrOQr2NtTqlKcYo+bfvYo81P0s0q8+Z9h4a06U+I+lqLvQ8zPbCpqeI2r+49moYdQK3lok2PhbkIDXoBp+NacfhIKQMJW0VZ74kdmolI9edR9ZlTyWvXMIM4jA4WSuDhGU0FDaGJzJzRAHwXj0EZNy37R+M9v54rYdcaZzqmnoFwQ+6hPF3FD4anG1yhI5Z+gIvs3aZ64HEeKrNOZEXR1pLmg17N8o2QNZyKenjm0LPIPpfYxRhJEzYoVte4ZiVqbJj+MSoNd7ReAYhesL8iSUFw85QyDZdq5O1tQokGL4RqU2/nV7fhpiSPcw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202604160,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202604160},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd81fd3e6aji65qtg7sbg_OsDisk_1_aef0e02e4b6a4ebca659642181e7ee12,storageAccountType:Premium_LRS},name:tbd81fd3e6aji65qtg7sbg_OsDisk_1_aef0e02e4b6a4ebca659642181e7ee12,osType:Linux}},timeCreated:2026-05-12T09:44:57.0636694Z,vmId:7f4cdcdf-b136-45ae-8898-b261faa81b51}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbd81fd3e6aji65qtg7sbg,keypair:tbd81fcu66aji65qtg7s7g,publicip:tbd81fd3e6aji65qtg7sbg-88014-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg"
            },
            {
              "key": "Name",
              "value": "tbd81fd3e6aji65qtg7sbg"
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
            "nodeIp": "20.214.4.247",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbd81fd3e6aji65qtg7sag 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "52.231.199.161",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbd81fd3e6aji65qtg7scg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "52.231.199.157",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbd81fd3e6aji65qtg7sbg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
    "my02-infra101"
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
  "id": "my02-infra101",
  "uid": "tbd81fd366aji65qtg7s9g",
  "name": "my02-infra101",
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
    "sys.id": "my02-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my02-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbd81fd366aji65qtg7s9g"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbd81fd3e6aji65qtg7sag",
      "cspResourceName": "tbd81fd3e6aji65qtg7sag",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag",
      "name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-05-12 09:45:45",
      "label": {
        "createdBy": "tbd81fd3e6aji65qtg7sag",
        "keypair": "tbd81fcu66aji65qtg7s7g",
        "publicip": "tbd81fd3e6aji65qtg7sag-6948-PublicIP",
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-12 09:45:45",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag",
        "sys.cspResourceName": "tbd81fd3e6aji65qtg7sag",
        "sys.id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd81fd3e6aji65qtg7sag",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.214.4.247",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 50,
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0",
      "networkInterface": "tbd81fd3e6aji65qtg7sag-97625-VNic",
      "securityGroupIds": [
        "my02-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOlhZnTKot5CpvRf0YzyGcCOMZuXeAMAgdZPJBjMssWEV3mgDw0DIAaZL1vj8YsjarGo8/mGHmt0/jtU0XI2DU0=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:NLSKcyS77IrfjQOPBXWdHjfqLW8dVJipgwWxHgZQYe0",
        "firstUsedAt": "2026-05-12T09:45:55Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-12T09:45:55Z",
          "completedTime": "2026-05-12T09:45:59Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd81fd3e6aji65qtg7sag 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd81fd3e6aji65qtg7sag-97625-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd81fd3e6aji65qtg7sag,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDUkSh+Hm49mSp1aeExqrCuh1wK6S64rPm4ONnlNHe088Sxf6kcgqyR5b+poXYPT2iUjyQBqWdiIqhHCd1iTxz0h0y92l8DmRS00ez9QV6IB3CV1ZKiBdP3s2S5btgEEDIVbnzmKoMg9yU67MtZE3ER9ePK4HzT/Ps0DNMbsyseEILG7q1kGauxvySBoaHb7x+Dai/879pIFLN2tHfYzakdAsnze1iH1NEzbnx+cfJb+X0q2EWU0cnjGJpGxCOsJ2jq4UuS7/mSBxQzPxZE0a34pBf3d6Jut8lrS/NMiWSrnAIX6ZHrYOrx1aJhbdkp2OXaYZyIrxUtlHxS75UlNjPAkTmVBrpySw5oeiRq6Y6Qxj7+tfQ+xrOQr2NtTqlKcYo+bfvYo81P0s0q8+Z9h4a06U+I+lqLvQ8zPbCpqeI2r+49moYdQK3lok2PhbkIDXoBp+NacfhIKQMJW0VZ74kdmolI9edR9ZlTyWvXMIM4jA4WSuDhGU0FDaGJzJzRAHwXj0EZNy37R+M9v54rYdcaZzqmnoFwQ+6hPF3FD4anG1yhI5Z+gIvs3aZ64HEeKrNOZEXR1pLmg17N8o2QNZyKenjm0LPIPpfYxRhJEzYoVte4ZiVqbJj+MSoNd7ReAYhesL8iSUFw85QyDZdq5O1tQokGL4RqU2/nV7fhpiSPcw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202604160,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202604160},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd81fd3e6aji65qtg7sag_OsDisk_1_8475ebb7b2fc4065a10ea9965259015e,storageAccountType:Premium_LRS},name:tbd81fd3e6aji65qtg7sag_OsDisk_1_8475ebb7b2fc4065a10ea9965259015e,osType:Linux}},timeCreated:2026-05-12T09:44:53.6437916Z,vmId:1b543ee0-8c3c-44e9-9d8b-573c0e3fc6c9}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd81fd3e6aji65qtg7sag,keypair:tbd81fcu66aji65qtg7s7g,publicip:tbd81fd3e6aji65qtg7sag-6948-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag"
        },
        {
          "key": "Name",
          "value": "tbd81fd3e6aji65qtg7sag"
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
      "uid": "tbd81fd3e6aji65qtg7scg",
      "cspResourceName": "tbd81fd3e6aji65qtg7scg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg",
      "name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-05-12 09:45:48",
      "label": {
        "createdBy": "tbd81fd3e6aji65qtg7scg",
        "keypair": "tbd81fcu66aji65qtg7s7g",
        "publicip": "tbd81fd3e6aji65qtg7scg-43856-PublicIP",
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-12 09:45:48",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg",
        "sys.cspResourceName": "tbd81fd3e6aji65qtg7scg",
        "sys.id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd81fd3e6aji65qtg7scg",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "52.231.199.161",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.6",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 50,
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0",
      "networkInterface": "tbd81fd3e6aji65qtg7scg-38274-VNic",
      "securityGroupIds": [
        "my02-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBGNC3nphY/v0+xM4rigF6rS6/BryTA6wUKcYifFO+GrlgOg05dR7/deJP4WKbiryKR++SIuzGvpIHF7XfWkCxbs=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:Y2gaNklAb4QHPPMwKNRhTABm2kr7c8p0/zvEwSdvNKU",
        "firstUsedAt": "2026-05-12T09:45:58Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-12T09:45:55Z",
          "completedTime": "2026-05-12T09:46:00Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd81fd3e6aji65qtg7scg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd81fd3e6aji65qtg7scg-38274-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd81fd3e6aji65qtg7scg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDUkSh+Hm49mSp1aeExqrCuh1wK6S64rPm4ONnlNHe088Sxf6kcgqyR5b+poXYPT2iUjyQBqWdiIqhHCd1iTxz0h0y92l8DmRS00ez9QV6IB3CV1ZKiBdP3s2S5btgEEDIVbnzmKoMg9yU67MtZE3ER9ePK4HzT/Ps0DNMbsyseEILG7q1kGauxvySBoaHb7x+Dai/879pIFLN2tHfYzakdAsnze1iH1NEzbnx+cfJb+X0q2EWU0cnjGJpGxCOsJ2jq4UuS7/mSBxQzPxZE0a34pBf3d6Jut8lrS/NMiWSrnAIX6ZHrYOrx1aJhbdkp2OXaYZyIrxUtlHxS75UlNjPAkTmVBrpySw5oeiRq6Y6Qxj7+tfQ+xrOQr2NtTqlKcYo+bfvYo81P0s0q8+Z9h4a06U+I+lqLvQ8zPbCpqeI2r+49moYdQK3lok2PhbkIDXoBp+NacfhIKQMJW0VZ74kdmolI9edR9ZlTyWvXMIM4jA4WSuDhGU0FDaGJzJzRAHwXj0EZNy37R+M9v54rYdcaZzqmnoFwQ+6hPF3FD4anG1yhI5Z+gIvs3aZ64HEeKrNOZEXR1pLmg17N8o2QNZyKenjm0LPIPpfYxRhJEzYoVte4ZiVqbJj+MSoNd7ReAYhesL8iSUFw85QyDZdq5O1tQokGL4RqU2/nV7fhpiSPcw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202604160,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202604160},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd81fd3e6aji65qtg7scg_OsDisk_1_9fca1cfb0b9a4d3ead6e5403bc3b9dac,storageAccountType:Premium_LRS},name:tbd81fd3e6aji65qtg7scg_OsDisk_1_9fca1cfb0b9a4d3ead6e5403bc3b9dac,osType:Linux}},timeCreated:2026-05-12T09:44:57.9968299Z,vmId:5e905251-d043-46c4-814e-489b8c8fe657}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd81fd3e6aji65qtg7scg,keypair:tbd81fcu66aji65qtg7s7g,publicip:tbd81fd3e6aji65qtg7scg-43856-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg"
        },
        {
          "key": "Name",
          "value": "tbd81fd3e6aji65qtg7scg"
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
      "uid": "tbd81fd3e6aji65qtg7sbg",
      "cspResourceName": "tbd81fd3e6aji65qtg7sbg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg",
      "name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-05-12 09:45:47",
      "label": {
        "createdBy": "tbd81fd3e6aji65qtg7sbg",
        "keypair": "tbd81fcu66aji65qtg7s7g",
        "publicip": "tbd81fd3e6aji65qtg7sbg-88014-PublicIP",
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-12 09:45:47",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg",
        "sys.cspResourceName": "tbd81fd3e6aji65qtg7sbg",
        "sys.id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd81fd3e6aji65qtg7sbg",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "52.231.199.157",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 50,
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0",
      "networkInterface": "tbd81fd3e6aji65qtg7sbg-96100-VNic",
      "securityGroupIds": [
        "my02-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAN6J4eynmqAFw1iIOLTlcDd2Nj3VXVk/ntLQad0HFdZmLHM1XlpHkgbQBQ0ZbT7WdXlCNwjQFUTJUz2Peu6tt8=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:oURuT7phD9rhemFegoCfIhW9JIwJd97nUZcaLxVND60",
        "firstUsedAt": "2026-05-12T09:45:58Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-12T09:45:55Z",
          "completedTime": "2026-05-12T09:46:00Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd81fd3e6aji65qtg7sbg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd81fd3e6aji65qtg7sbg-96100-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd81fd3e6aji65qtg7sbg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDUkSh+Hm49mSp1aeExqrCuh1wK6S64rPm4ONnlNHe088Sxf6kcgqyR5b+poXYPT2iUjyQBqWdiIqhHCd1iTxz0h0y92l8DmRS00ez9QV6IB3CV1ZKiBdP3s2S5btgEEDIVbnzmKoMg9yU67MtZE3ER9ePK4HzT/Ps0DNMbsyseEILG7q1kGauxvySBoaHb7x+Dai/879pIFLN2tHfYzakdAsnze1iH1NEzbnx+cfJb+X0q2EWU0cnjGJpGxCOsJ2jq4UuS7/mSBxQzPxZE0a34pBf3d6Jut8lrS/NMiWSrnAIX6ZHrYOrx1aJhbdkp2OXaYZyIrxUtlHxS75UlNjPAkTmVBrpySw5oeiRq6Y6Qxj7+tfQ+xrOQr2NtTqlKcYo+bfvYo81P0s0q8+Z9h4a06U+I+lqLvQ8zPbCpqeI2r+49moYdQK3lok2PhbkIDXoBp+NacfhIKQMJW0VZ74kdmolI9edR9ZlTyWvXMIM4jA4WSuDhGU0FDaGJzJzRAHwXj0EZNy37R+M9v54rYdcaZzqmnoFwQ+6hPF3FD4anG1yhI5Z+gIvs3aZ64HEeKrNOZEXR1pLmg17N8o2QNZyKenjm0LPIPpfYxRhJEzYoVte4ZiVqbJj+MSoNd7ReAYhesL8iSUFw85QyDZdq5O1tQokGL4RqU2/nV7fhpiSPcw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202604160,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202604160},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd81fd3e6aji65qtg7sbg_OsDisk_1_aef0e02e4b6a4ebca659642181e7ee12,storageAccountType:Premium_LRS},name:tbd81fd3e6aji65qtg7sbg_OsDisk_1_aef0e02e4b6a4ebca659642181e7ee12,osType:Linux}},timeCreated:2026-05-12T09:44:57.0636694Z,vmId:7f4cdcdf-b136-45ae-8898-b261faa81b51}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd81fd3e6aji65qtg7sbg,keypair:tbd81fcu66aji65qtg7s7g,publicip:tbd81fd3e6aji65qtg7sbg-88014-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg"
        },
        {
          "key": "Name",
          "value": "tbd81fd3e6aji65qtg7sbg"
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
        "nodeIp": "20.214.4.247",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd81fd3e6aji65qtg7sag 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my02-infra101",
        "nodeId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "52.231.199.161",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd81fd3e6aji65qtg7scg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my02-infra101",
        "nodeId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "nodeIp": "52.231.199.157",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd81fd3e6aji65qtg7sbg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "nodeGroup": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "nodeId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "output": "Linux tbd81fd3e6aji65qtg7sag 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "20.214.4.247",
      "sshTest": "successful",
      "status": "success",
      "testOrder": 1,
      "userName": "cb-user"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "nodeGroup": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "nodeId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "output": "Linux tbd81fd3e6aji65qtg7scg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.199.161",
      "sshTest": "successful",
      "status": "success",
      "testOrder": 2,
      "userName": "cb-user"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "nodeGroup": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "nodeId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "output": "Linux tbd81fd3e6aji65qtg7sbg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.199.157",
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

**Generated At:** 2026-05-12 09:46:42

**Namespace:** mig01

**Infra Name:** my02-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my02-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | AZURE |
| **Target Region** | koreasouth |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| Standard_B2als_v2 | 2 | 3.9 | - | x86_64 |  | $0.0432 | 1 |
| Standard_B2as_v2 | 2 | 7.8 | - | x86_64 |  | $0.0865 | 1 |
| Standard_B4as_v2 | 4 | 15.6 | - | x86_64 |  | $0.1730 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160 | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 1 |
| Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 2 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag | Running | 2 vCPU, 3.9 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160) | **VNet:** my02-vnet-01<br>**Subnet:** my02-subnet-01<br>**Public IP:** 20.214.4.247<br>**Private IP:** 10.0.1.4<br>**SGs:** my02-sg-01<br>**SSH:** my02-sshkey-01 |
| my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg | Running | 2 vCPU, 7.8 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160) | **VNet:** my02-vnet-01<br>**Subnet:** my02-subnet-01<br>**Public IP:** 52.231.199.161<br>**Private IP:** 10.0.1.6<br>**SGs:** my02-sg-03<br>**SSH:** my02-sshkey-01 |
| my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg | Running | 4 vCPU, 15.6 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160) | **VNet:** my02-vnet-01<br>**Subnet:** my02-subnet-01<br>**Public IP:** 52.231.199.157<br>**Private IP:** 10.0.1.5<br>**SGs:** my02-sg-02<br>**SSH:** my02-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my02-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my02-vnet-01 |
| **CSP VNet ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | azure-koreasouth |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my02-subnet-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0 | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my02-sshkey-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g |  |  |

### Security Groups

#### Security Group: my02-sg-01

| Property | Value |
|----------|-------|
| **Name** | my02-sg-01 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd81fcvm6aji65qtg7s80 |
| **VNet** | my02-vnet-01 |
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

#### Security Group: my02-sg-02

| Property | Value |
|----------|-------|
| **Name** | my02-sg-02 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd81fd0u6aji65qtg7s8g |
| **VNet** | my02-vnet-01 |
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

#### Security Group: my02-sg-03

| Property | Value |
|----------|-------|
| **Name** | my02-sg-03 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd81fd1u6aji65qtg7s90 |
| **VNet** | my02-vnet-01 |
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
| **Per Hour** | $0.3027 |
| **Per Day** | $7.26 |
| **Per Month (30 days)** | $217.94 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| AZURE | koreasouth | 3 | $0.3027 | $217.94 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | Standard_B2als_v2 | $0.0432 | $31.10 |
| my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | Standard_B2as_v2 | $0.0865 | $62.28 |
| my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | Standard_B4as_v2 | $0.1730 | $124.56 |




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

*Report generated: 2026-05-12 09:46:48*

---

## 📊 Migration Summary

**Target Cloud:** AZURE

**Target Region:** koreasouth

**Namespace:** mig01 | **Infra ID:** my02-infra101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

**💰 Estimated Monthly Cost:** $217.94 USD

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| # | Resource Type | Count | Status | Details |
|---|---------------|-------|--------|----------|
| 1 | **Virtual Machine** | 3 | ✅ Created | 3 running, 3 total |
| 2 | **VM Spec** | 3 | ✅ Selected | Standard_B4as_v2, Standard_B2als_v2, Standard_B2as_v2 |
| 3 | **VM OS Image** | 2 | ✅ Selected | Ubuntu 22.04 |
| 4 | **VNet (VPC)** | 1 | ✅ Created | my02-vnet-01, CIDR: 10.0.0.0/21 |
| 5 | **Subnet** | 1 | ✅ Created | 10.0.1.0/24 (in my02-vnet-01) |
| 6 | **Security Group** | 3 security groups | ✅ Created | Total 52 rules in 3 sgs |
| 7 | **SSH Key** | 1 keys | ✅ Created | For VM access control |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source node(s) in the on-premise infrastructure.

| No. | Migrated VM | Source Server |
|-----|-------------|---------------|
| 1 | **VM Name:** my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sag<br>**Label(sourceMachineId):** vm-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7scg<br>**Label(sourceMachineId):** vm-ec288dd0-c6fa-8a49-2f60 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 |
| 3 | **VM Name:** my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd81fd3e6aji65qtg7sbg<br>**Label(sourceMachineId):** vm-ec2d32b5-98fb-5a96-7913 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 |

---

## ⚙️ VM Specs

**Summary:** 3 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM Spec | Source Server | Source Server Spec |
|-----|-------------|---------|---------------|--------------------|
| 1 | my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** Standard_B2als_v2<br>**vCPUs:** 2<br>**Memory:** 3.9 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2 | my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Spec ID:** Standard_B2as_v2<br>**vCPUs:** 2<br>**Memory:** 7.8 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3 | my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Spec ID:** Standard_B4as_v2<br>**vCPUs:** 4<br>**Memory:** 15.6 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 2 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM | VM OS Image Info | Source Server | Source OS |
|-----|-------------|------------------|---------------|-----------|
| 1 | my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160 | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my02-sg-01

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd81fcvm6aji65qtg7s80 | **VNet:** my02-vnet-01 | **Rules:** 14

**Assigned VMs:**

- **VM:** my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1
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

### Security Group: my02-sg-02

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd81fd0u6aji65qtg7s8g | **VNet:** my02-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1
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

### Security Group: my02-sg-03

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd81fd1u6aji65qtg7s90 | **VNet:** my02-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1
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
| 1 | **Name:** my02-vnet-01<br>**ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my02-subnet-01<br>**ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd81fcrm6aji65qtg7rug/subnets/tbd81fcrm6aji65qtg7rv0 | 10.0.1.0/24 | my02-vnet-01 |

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
| 1 | my02-sshkey-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tbd81fcu66aji65qtg7s7g |  | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period | Cost (USD) |
|--------|------------|
| Hourly | $0.3027 |
| Daily | $7.26 |
| Monthly | $217.94 |
| Yearly | $2615.33 |

### Cost Breakdown by Component

| Component | Spec | Monthly Cost | Percentage |
|-----------|------|--------------|------------|
| ip-10-0-1-30 (migrated) | Standard_B2als_v2 | $31.10 | 14.3% |
| ip-10-0-1-221 (migrated) | Standard_B4as_v2 | $124.56 | 57.2% |
| ip-10-0-1-138 (migrated) | Standard_B2as_v2 | $62.28 | 28.6% |

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
  "message": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: my02-infra101)",
  "success": true
}
```

