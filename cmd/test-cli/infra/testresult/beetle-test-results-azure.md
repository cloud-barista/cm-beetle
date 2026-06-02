# CM-Beetle test results for AZURE

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: imdl/v0.1.5+ (4987539)
- imdl: v0.1.5+ (4987539)
- CB-Tumblebug: v0.12.13
- CB-Spider: v0.12.26
- CB-MapUI: v0.12.34
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: June 2, 2026
- Test Time: 20:57:03 KST
- Test Execution: 2026-06-02 20:57:03 KST

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
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 13.084s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m36.82s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 126ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 16ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 28ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.584s | Pass |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.604s | Pass |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m47.036s | Pass |

**Overall Result**: 9/9 tests passed ✅

**Total Duration**: 4m57.44799377s

*Test executed on June 2, 2026 at 20:57:03 KST (2026-06-02 20:57:03 KST) using CM-Beetle automated test CLI*

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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605280",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 30,
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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605280",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 30,
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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605280",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 30,
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
          "uid": "tbu3qnmmi15svg3cnahm",
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
          "uid": "tbclu29d1tthdaa1orf0",
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
          "uid": "tb3h0eijnsghu8e1q6t8",
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
          "uid": "tbpgo13bkvl7mci3fmno",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.05.27 06:36:46 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
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
              "value": "22.04.202605260"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy-daily/Skus/22_04-daily-lts-gen2/Versions/22.04.202605260"
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
          "uid": "tbgtm1agmi0me5fnfl35",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.05.27 06:36:46 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
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
              "value": "22.04.202605260"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy-daily/Skus/22_04-daily-lts/Versions/22.04.202605260"
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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605280",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 30,
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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605280",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 30,
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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605280",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 30,
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
          "uid": "tbu3qnmmi15svg3cnahm",
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
          "uid": "tbclu29d1tthdaa1orf0",
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
          "uid": "tb3h0eijnsghu8e1q6t8",
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
          "uid": "tbe2su80r3g6ipn1o5c4",
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
          "uid": "tbafcljckm3ds559b5ti",
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
          "uid": "tb6a6lihtts92j0q40ol",
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
          "uid": "tbpgo13bkvl7mci3fmno",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.05.27 06:36:46 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260",
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
              "value": "22.04.202605260"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy-daily/Skus/22_04-daily-lts-gen2/Versions/22.04.202605260"
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
          "uid": "tbgtm1agmi0me5fnfl35",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.05.27 06:36:46 Wed",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260",
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
              "value": "22.04.202605260"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy-daily/Skus/22_04-daily-lts/Versions/22.04.202605260"
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
  "uid": "tbv4r5dg0r6ctsgj7tfe",
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
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
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
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
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
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
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
      "uid": "tb2drt09mfdcqasro93p",
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
        "sys.uid": "tb2drt09mfdcqasro93p"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb9lrr7qljeqacqgv2u5",
          "cspResourceName": "tb9lrr7qljeqacqgv2u5",
          "cspResourceId": "i-06534a82f97650b25",
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
          "createdTime": "2026-06-02 11:57:58",
          "label": {
            "Name": "tb9lrr7qljeqacqgv2u5",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-06-02 11:57:58",
            "sys.cspResourceId": "i-06534a82f97650b25",
            "sys.cspResourceName": "tb9lrr7qljeqacqgv2u5",
            "sys.id": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tb9lrr7qljeqacqgv2u5",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "52.78.244.82",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.140",
          "privateDNS": "ip-10-0-1-140.ap-northeast-2.compute.internal",
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
          "imageId": "ami-0596f7562954deb8e",
          "cspImageName": "ami-0596f7562954deb8e",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-0596f7562954deb8e",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260521"
          },
          "vNetId": "my01-vnet-01",
          "cspVNetId": "vpc-00b2b1e59e8d2653d",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-005c104e63421b753",
          "networkInterface": "eni-attach-05d9a8c5dc9658780",
          "securityGroupIds": [
            "my01-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbgkl68qoqbt2irs6nn1",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJTQ3SSJJ7tSaj0WZyFGnlovTpmAJ3neLicuOZcZ8+5894+2Fg7gxHO3jPuVAXf7mjQZiXmXhr3ebTh7z3C5fNk=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:3SwXDiiyzPRZcm7FKBhxVlokHTSSsW+3aNJln+hby+Y",
            "firstUsedAt": "2026-06-02T11:58:07Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-02T11:58:06Z",
              "completedTime": "2026-06-02T11:58:08Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-140 6.8.0-1055-aws #58~22.04.1-Ubuntu SMP Thu May  7 22:16:53 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-06-02T11:57:39Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0e1622ebc7e0ae0be}}"
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
              "value": "2A6DB7BA-D37C-45E3-8AFE-B062DEC1496A"
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
              "value": "ami-0596f7562954deb8e"
            },
            {
              "key": "InstanceId",
              "value": "i-06534a82f97650b25"
            },
            {
              "key": "InstanceType",
              "value": "t3a.small"
            },
            {
              "key": "KeyName",
              "value": "tbgkl68qoqbt2irs6nn1"
            },
            {
              "key": "LaunchTime",
              "value": "2026-06-02T11:57:38Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:52.78.244.82},Attachment:{AttachTime:2026-06-02T11:57:38Z,AttachmentId:eni-attach-05d9a8c5dc9658780,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-073146fb5cc82a36a,GroupName:tbqf93ju7qe7u9ri3rs4}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:55:94:dc:34:3b,NetworkInterfaceId:eni-04ea4fac87e5ad333,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.140,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:52.78.244.82},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.140}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-005c104e63421b753,VpcId:vpc-00b2b1e59e8d2653d}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-140.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.140"
            },
            {
              "key": "PublicIpAddress",
              "value": "52.78.244.82"
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
              "value": "{GroupId:sg-073146fb5cc82a36a,GroupName:tbqf93ju7qe7u9ri3rs4}"
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
              "value": "subnet-005c104e63421b753"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tb9lrr7qljeqacqgv2u5}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-00b2b1e59e8d2653d"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tbd9vt5h7b3mi7hlnsrl",
          "cspResourceName": "tbd9vt5h7b3mi7hlnsrl",
          "cspResourceId": "i-0b8bf7f9b3b9876c4",
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
          "createdTime": "2026-06-02 11:58:01",
          "label": {
            "Name": "tbd9vt5h7b3mi7hlnsrl",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-06-02 11:58:01",
            "sys.cspResourceId": "i-0b8bf7f9b3b9876c4",
            "sys.cspResourceName": "tbd9vt5h7b3mi7hlnsrl",
            "sys.id": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tbd9vt5h7b3mi7hlnsrl",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "43.201.83.130",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.136",
          "privateDNS": "ip-10-0-1-136.ap-northeast-2.compute.internal",
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
          "imageId": "ami-0596f7562954deb8e",
          "cspImageName": "ami-0596f7562954deb8e",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-0596f7562954deb8e",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260521"
          },
          "vNetId": "my01-vnet-01",
          "cspVNetId": "vpc-00b2b1e59e8d2653d",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-005c104e63421b753",
          "networkInterface": "eni-attach-00d6d6972528706c2",
          "securityGroupIds": [
            "my01-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbgkl68qoqbt2irs6nn1",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBNtx2r8Ou4fXFF0/0OPQ0Rk+x8nxvnqwJeTL+9tBaIzgNXkZJMZ89BpJskE2U2XGlPGjmzheQYG3idrFrwV6U98=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:kledbHEJr8DxkYCKFS+YqOfbAGwfb7bYstah9IAFp4U",
            "firstUsedAt": "2026-06-02T11:58:08Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-06-02T11:58:06Z",
              "completedTime": "2026-06-02T11:58:26Z",
              "elapsedTime": 20,
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-06-02T11:57:41Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0cd90939abf365ee0}}"
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
              "value": "8A15C17A-F15C-460F-A71C-A57EF5F8E78E"
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
              "value": "ami-0596f7562954deb8e"
            },
            {
              "key": "InstanceId",
              "value": "i-0b8bf7f9b3b9876c4"
            },
            {
              "key": "InstanceType",
              "value": "t3a.large"
            },
            {
              "key": "KeyName",
              "value": "tbgkl68qoqbt2irs6nn1"
            },
            {
              "key": "LaunchTime",
              "value": "2026-06-02T11:57:40Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.201.83.130},Attachment:{AttachTime:2026-06-02T11:57:40Z,AttachmentId:eni-attach-00d6d6972528706c2,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0199f16c43ae8cc63,GroupName:tb6qo7797mvvc05iul83}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:67:60:31:ee:45,NetworkInterfaceId:eni-010ee71d81f863c52,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.136,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.201.83.130},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.136}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-005c104e63421b753,VpcId:vpc-00b2b1e59e8d2653d}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-136.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.136"
            },
            {
              "key": "PublicIpAddress",
              "value": "43.201.83.130"
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
              "value": "{GroupId:sg-0199f16c43ae8cc63,GroupName:tb6qo7797mvvc05iul83}"
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
              "value": "subnet-005c104e63421b753"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbd9vt5h7b3mi7hlnsrl}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-00b2b1e59e8d2653d"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbrh382k22c02gj6or2q",
          "cspResourceName": "tbrh382k22c02gj6or2q",
          "cspResourceId": "i-0273cb7e68fdf8fac",
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
          "createdTime": "2026-06-02 11:58:00",
          "label": {
            "Name": "tbrh382k22c02gj6or2q",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-06-02 11:58:00",
            "sys.cspResourceId": "i-0273cb7e68fdf8fac",
            "sys.cspResourceName": "tbrh382k22c02gj6or2q",
            "sys.id": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tbrh382k22c02gj6or2q",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "52.78.4.197",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.239",
          "privateDNS": "ip-10-0-1-239.ap-northeast-2.compute.internal",
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
          "imageId": "ami-0596f7562954deb8e",
          "cspImageName": "ami-0596f7562954deb8e",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-0596f7562954deb8e",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20260521"
          },
          "vNetId": "my01-vnet-01",
          "cspVNetId": "vpc-00b2b1e59e8d2653d",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-005c104e63421b753",
          "networkInterface": "eni-attach-0638cbee4cdabd7bd",
          "securityGroupIds": [
            "my01-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbgkl68qoqbt2irs6nn1",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBH97D952EztvAI/plpOdoCRtd21dx035GA7gRqjte2Nms1wIDFCgt9jQFauHyJRz8HXg/43acRo6d4xSZzc2ACw=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:Fk9INcBJUzgOLcr6k2etg2P8ucfdRb5QkjAk/dNW5Bo",
            "firstUsedAt": "2026-06-02T11:58:08Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Failed",
              "startedTime": "2026-06-02T11:58:06Z",
              "completedTime": "2026-06-02T11:58:26Z",
              "elapsedTime": 20,
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-06-02T11:57:41Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0387f952743f64ec0}}"
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
              "value": "9F574316-EB84-4652-B3FF-92E90421877B"
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
              "value": "ami-0596f7562954deb8e"
            },
            {
              "key": "InstanceId",
              "value": "i-0273cb7e68fdf8fac"
            },
            {
              "key": "InstanceType",
              "value": "t3a.xlarge"
            },
            {
              "key": "KeyName",
              "value": "tbgkl68qoqbt2irs6nn1"
            },
            {
              "key": "LaunchTime",
              "value": "2026-06-02T11:57:40Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:52.78.4.197},Attachment:{AttachTime:2026-06-02T11:57:40Z,AttachmentId:eni-attach-0638cbee4cdabd7bd,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0ce7c1c68b49dfdfd,GroupName:tbgd82m18cv28pblnuad}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:c1:be:34:65:23,NetworkInterfaceId:eni-0dff89b4de2d2ff27,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.239,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:52.78.4.197},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.239}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-005c104e63421b753,VpcId:vpc-00b2b1e59e8d2653d}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-239.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.239"
            },
            {
              "key": "PublicIpAddress",
              "value": "52.78.4.197"
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
              "value": "{GroupId:sg-0ce7c1c68b49dfdfd,GroupName:tbgd82m18cv28pblnuad}"
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
              "value": "subnet-005c104e63421b753"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbrh382k22c02gj6or2q}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-00b2b1e59e8d2653d"
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
            "nodeId": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "52.78.244.82",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-140 6.8.0-1055-aws #58~22.04.1-Ubuntu SMP Thu May  7 22:16:53 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "52.78.4.197",
            "command": {
              "0": "uname -a"
            },
            "stdout": {},
            "stderr": {},
            "err": null
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "43.201.83.130",
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
      "uid": "tbv4r5dg0r6ctsgj7tfe",
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
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
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
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
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
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
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
      "id": "my04-infra101",
      "uid": "tbsb9sq51v4fnsipkeap",
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
        "sys.uid": "tbsb9sq51v4fnsipkeap"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbhs2tl9c3te0mksjp02",
          "cspResourceName": "tbhs2tl9c3te0mksjp02",
          "cspResourceId": "i-mj7e35tls2uwpys3780q",
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
          "createdTime": "2026-06-02 11:58:17",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-06-02 11:58:17",
            "sys.cspResourceId": "i-mj7e35tls2uwpys3780q",
            "sys.cspResourceName": "tbhs2tl9c3te0mksjp02",
            "sys.id": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my04-subnet-01",
            "sys.uid": "tbhs2tl9c3te0mksjp02",
            "sys.vNetId": "my04-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.213.148.90",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.130",
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
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m1.large",
          "cspSpecName": "ecs.e-c1m1.large",
          "spec": {
            "cspSpecName": "ecs.e-c1m1.large",
            "vCPU": 2,
            "memoryGiB": 2,
            "costPerHour": 0.0178
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-vnet-01",
          "cspVNetId": "vpc-mj77rilp5f4fn312w79hj",
          "subnetId": "my04-subnet-01",
          "cspSubnetId": "vsw-mj713tfhuu6jxzc6pn93z",
          "networkInterface": "eni-mj7e35tls2uwpys2lvv7",
          "securityGroupIds": [
            "my04-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-sshkey-01",
          "cspSshKeyId": "tbkc0tqjtqulqspk3vmj",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "$cpn1fbmnt!Abl",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBNhMNxzCV4xDgqQZg5+XGyt06ozb7LlWixrfbFJWox/iiHeJdkvx+xu/LmaVx6kjg/KPkOVil3M8LlcpqqmEpdk=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:wyhhKV/Me4xRzt4oodcl1P0bOWE1botx6CeeFuzBOW0",
            "firstUsedAt": "2026-06-02T11:58:31Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-02T11:58:31Z",
              "completedTime": "2026-06-02T11:58:33Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj7e35tls2uwpys3780qZ 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260506.vhd"
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
              "value": "tbhs2tl9c3te0mksjp02"
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
              "value": "2026-06-02T11:57Z"
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
              "value": "iZmj7e35tls2uwpys3780qZ"
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
              "value": "aaf96940-aed7-409b-939f-aea783b600b6"
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
              "value": "i-mj7e35tls2uwpys3780q"
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
              "value": "2026-06-02T11:57Z"
            },
            {
              "key": "KeyPairName",
              "value": "tbkc0tqjtqulqspk3vmj"
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
              "value": "{SecurityGroupId:[sg-mj71w6ok3fzhi53tij5j]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[8.213.148.90]}"
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
              "value": "{VSwitchId:vsw-mj713tfhuu6jxzc6pn93z,VpcId:vpc-mj77rilp5f4fn312w79hj,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.130]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:b8:a8,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7e35tls2uwpys2lvv7,PrimaryIpAddress:10.0.1.130,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.130,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
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
          "uid": "tbcpprl9juav33e9q3c2",
          "cspResourceName": "tbcpprl9juav33e9q3c2",
          "cspResourceId": "i-mj7e35tls2uwpys3780r",
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
          "createdTime": "2026-06-02 11:58:24",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-06-02 11:58:24",
            "sys.cspResourceId": "i-mj7e35tls2uwpys3780r",
            "sys.cspResourceName": "tbcpprl9juav33e9q3c2",
            "sys.id": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my04-subnet-01",
            "sys.uid": "tbcpprl9juav33e9q3c2",
            "sys.vNetId": "my04-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.241.105",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.131",
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
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
          "cspSpecName": "ecs.e-c1m4.large",
          "spec": {
            "cspSpecName": "ecs.e-c1m4.large",
            "vCPU": 2,
            "memoryGiB": 8,
            "costPerHour": 0.0791
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-vnet-01",
          "cspVNetId": "vpc-mj77rilp5f4fn312w79hj",
          "subnetId": "my04-subnet-01",
          "cspSubnetId": "vsw-mj713tfhuu6jxzc6pn93z",
          "networkInterface": "eni-mj7e35tls2uwpys2lvva",
          "securityGroupIds": [
            "my04-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-sshkey-01",
          "cspSshKeyId": "tbkc0tqjtqulqspk3vmj",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "og1p!0lt8bfAm$",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBD1gzSk9e926lbr/5G5aVaQiCrn/ll3UEMwqHa/mQNWN3remyRmCxaVjOBpRTG6Oq8982X2IJ403G4BSAEzHBmg=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:riHiQusw8+6WASDQXHPWeQ9s9W79YdIKy6bv+E597Kc",
            "firstUsedAt": "2026-06-02T11:58:32Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-02T11:58:31Z",
              "completedTime": "2026-06-02T11:58:33Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj7e35tls2uwpys3780rZ 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260506.vhd"
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
              "value": "tbcpprl9juav33e9q3c2"
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
              "value": "2026-06-02T11:57Z"
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
              "value": "iZmj7e35tls2uwpys3780rZ"
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
              "value": "769359a1-ff5c-47f6-aa50-705ede2c8b58"
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
              "value": "i-mj7e35tls2uwpys3780r"
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
              "value": "2026-06-02T11:57Z"
            },
            {
              "key": "KeyPairName",
              "value": "tbkc0tqjtqulqspk3vmj"
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
              "value": "{SecurityGroupId:[sg-mj780ni10nbsxrfvdoxa]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[47.80.241.105]}"
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
              "value": "{VSwitchId:vsw-mj713tfhuu6jxzc6pn93z,VpcId:vpc-mj77rilp5f4fn312w79hj,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.131]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:b8:a9,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7e35tls2uwpys2lvva,PrimaryIpAddress:10.0.1.131,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.131,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
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
          "uid": "tb0gc9fc2ec0v7te0en2",
          "cspResourceName": "tb0gc9fc2ec0v7te0en2",
          "cspResourceId": "i-mj7caxyslyfjiirlhjj0",
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
          "createdTime": "2026-06-02 11:58:19",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-06-02 11:58:19",
            "sys.cspResourceId": "i-mj7caxyslyfjiirlhjj0",
            "sys.cspResourceName": "tb0gc9fc2ec0v7te0en2",
            "sys.id": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my04-subnet-01",
            "sys.uid": "tb0gc9fc2ec0v7te0en2",
            "sys.vNetId": "my04-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.241.200",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.129",
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
          "specId": "alibaba+ap-northeast-2+ecs.e-c1m4.xlarge",
          "cspSpecName": "ecs.e-c1m4.xlarge",
          "spec": {
            "cspSpecName": "ecs.e-c1m4.xlarge",
            "vCPU": 4,
            "memoryGiB": 16,
            "costPerHour": 0.1582
          },
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-vnet-01",
          "cspVNetId": "vpc-mj77rilp5f4fn312w79hj",
          "subnetId": "my04-subnet-01",
          "cspSubnetId": "vsw-mj713tfhuu6jxzc6pn93z",
          "networkInterface": "eni-mj7caxyslyfjiirmuk6r",
          "securityGroupIds": [
            "my04-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-sshkey-01",
          "cspSshKeyId": "tbkc0tqjtqulqspk3vmj",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "07qr$1j!1Ab9ot",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBA36aUcHKHQ+RZ14yQy69ODrhFbF5gSjK7aKhCsDtxNKCAhSEkZ/x4ee4OO+z34nzoF7nlslcVVpwf2C1pOA02k=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:Il27QVw0TAS24BNoLHeyuwBFqvo13kGaDO+NVrf6Roo",
            "firstUsedAt": "2026-06-02T11:58:32Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-06-02T11:58:31Z",
              "completedTime": "2026-06-02T11:58:33Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj7caxyslyfjiirlhjj0Z 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
          "addtionalDetails": [
            {
              "key": "ImageId",
              "value": "ubuntu_22_04_x64_20G_alibase_20260506.vhd"
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
              "value": "tb0gc9fc2ec0v7te0en2"
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
              "value": "2026-06-02T11:57Z"
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
              "value": "iZmj7caxyslyfjiirlhjj0Z"
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
              "value": "6ea8094f-30c3-4bf8-b920-8f69ab5460e0"
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
              "value": "i-mj7caxyslyfjiirlhjj0"
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
              "value": "2026-06-02T11:57Z"
            },
            {
              "key": "KeyPairName",
              "value": "tbkc0tqjtqulqspk3vmj"
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
              "value": "{SecurityGroupId:[sg-mj7d6hrwmdy4yo1s8y96]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[47.80.241.200]}"
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
              "value": "{VSwitchId:vsw-mj713tfhuu6jxzc6pn93z,VpcId:vpc-mj77rilp5f4fn312w79hj,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.129]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:06:b8:a7,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7caxyslyfjiirmuk6r,PrimaryIpAddress:10.0.1.129,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.129,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
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
            "nodeId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "8.213.148.90",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7e35tls2uwpys3780qZ 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "47.80.241.200",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7caxyslyfjiirlhjj0Z 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "47.80.241.105",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7e35tls2uwpys3780rZ 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
          "name": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my06-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": null,
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": ""
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
          "rootDiskSize": 100,
          "RootDeviceName": "",
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
          "cspSpecName": "",
          "spec": {},
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "osType": ""
          },
          "vNetId": "my06-vnet-01",
          "cspVNetId": "",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my06-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tb1h7ijtv4sdjjt74cq4",
          "name": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my06-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": null,
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": ""
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
          "rootDiskSize": 100,
          "RootDeviceName": "",
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
          "cspSpecName": "",
          "spec": {},
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "osType": ""
          },
          "vNetId": "my06-vnet-01",
          "cspVNetId": "",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my06-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "node",
          "id": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbfb0c8hrtma9tjabtm6",
          "name": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my06-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "location": {
            "display": "Australia (Sydney)",
            "latitude": -33.86882,
            "longitude": 151.209296
          },
          "status": "Creating",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": null,
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": ""
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
          "rootDiskSize": 100,
          "RootDeviceName": "",
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
          "cspSpecName": "",
          "spec": {},
          "imageId": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "cspImageName": "r026-c8e249d4-f148-4416-a3c6-555b7a02f67d",
          "image": {
            "osType": ""
          },
          "vNetId": "my06-vnet-01",
          "cspVNetId": "",
          "subnetId": "my06-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": [
            "my06-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my06-sshkey-01",
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
    "my04-infra101",
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
  "id": "my02-infra101",
  "uid": "tbv4r5dg0r6ctsgj7tfe",
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
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
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
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
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
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
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
      "output": "Linux tb81h514blbmih6th0ra 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "40.89.209.210",
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
      "output": "Linux tbj3l6ut11ch7mvkq2p2 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.195.8",
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
      "output": "Linux tblumcks3qbld4l0rqdu 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "20.214.40.12",
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

**Generated At:** 2026-06-02 11:59:38

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
| Standard_B2as_v2 | 2 | 7.8 | - | x86_64 |  | $0.0865 | 1 |
| Standard_B4as_v2 | 4 | 15.6 | - | x86_64 |  | $0.1730 | 1 |
| Standard_B2als_v2 | 2 | 3.9 | - | x86_64 |  | $0.0432 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260 | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 2 |
| Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260 | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 1 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb81h514blbmih6th0ra | Running | 2 vCPU, 3.9 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260) | **VNet:** my02-vnet-01<br>**Subnet:** my02-subnet-01<br>**Public IP:** 40.89.209.210<br>**Private IP:** 10.0.1.4<br>**SGs:** my02-sg-01<br>**SSH:** my02-sshkey-01 |
| my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbj3l6ut11ch7mvkq2p2 | Running | 2 vCPU, 7.8 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260) | **VNet:** my02-vnet-01<br>**Subnet:** my02-subnet-01<br>**Public IP:** 52.231.195.8<br>**Private IP:** 10.0.1.6<br>**SGs:** my02-sg-03<br>**SSH:** my02-sshkey-01 |
| my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tblumcks3qbld4l0rqdu | Running | 4 vCPU, 15.6 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260) | **VNet:** my02-vnet-01<br>**Subnet:** my02-subnet-01<br>**Public IP:** 20.214.40.12<br>**Private IP:** 10.0.1.5<br>**SGs:** my02-sg-02<br>**SSH:** my02-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my02-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my02-vnet-01 |
| **CSP VNet ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbaa7lm1qpjpqcm1blta |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | azure-koreasouth |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my02-subnet-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbaa7lm1qpjpqcm1blta/subnets/tbdk04p2j6rjaddfeilr | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my02-sshkey-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tbmeu2jlibrbjhpqt6ks |  |  |

### Security Groups

#### Security Group: my02-sg-01

| Property | Value |
|----------|-------|
| **Name** | my02-sg-01 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tba5q2q197u97gaalfpr |
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
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbfrrlj3eu12v1kqh8rp |
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
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tb5vuraeoacia4ihdkd6 |
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

*Report generated: 2026-06-02 11:59:44*

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
| 2 | **VM Spec** | 3 | ✅ Selected | Standard_B2als_v2, Standard_B2as_v2, Standard_B4as_v2 |
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
| 1 | **VM Name:** my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb81h514blbmih6th0ra<br>**Label(sourceMachineId):** vm-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbj3l6ut11ch7mvkq2p2<br>**Label(sourceMachineId):** vm-ec288dd0-c6fa-8a49-2f60 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 |
| 3 | **VM Name:** my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tblumcks3qbld4l0rqdu<br>**Label(sourceMachineId):** vm-ec2d32b5-98fb-5a96-7913 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 |

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
| 1 | my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260 | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605260 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605260 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my02-sg-01

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tba5q2q197u97gaalfpr | **VNet:** my02-vnet-01 | **Rules:** 14

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

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbfrrlj3eu12v1kqh8rp | **VNet:** my02-vnet-01 | **Rules:** 19

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

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tb5vuraeoacia4ihdkd6 | **VNet:** my02-vnet-01 | **Rules:** 19

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
| 1 | **Name:** my02-vnet-01<br>**ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbaa7lm1qpjpqcm1blta | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my02-subnet-01<br>**ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbaa7lm1qpjpqcm1blta/subnets/tbdk04p2j6rjaddfeilr | 10.0.1.0/24 | my02-vnet-01 |

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
| 1 | my02-sshkey-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tbmeu2jlibrbjhpqt6ks |  | Used by all 3 VMs |

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

