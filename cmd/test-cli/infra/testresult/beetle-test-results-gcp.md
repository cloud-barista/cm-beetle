# CM-Beetle test results for GCP

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with GCP cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.5.9+ (83b6bb2)
- imdl: v0.1.10+ (83b6bb2)
- CB-Tumblebug: v0.12.30
- CB-Spider: v0.12.42
- CB-MapUI: v0.12.56
- Target CSP: GCP
- Target Region: asia-northeast3
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: August 11, 2026
- Test Time: 13:17:54 KST
- Test Execution: 2026-08-11 13:17:54 KST

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Validate the target model for computing infra via Beetle
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

## Test result for GCP

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 20.119s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 552ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 8m19.818s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 21ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 3ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 32ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 4.433s | Pass |
| 8 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.74s | Pass |
| 9 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.327s | Pass |
| 10 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 11m7.776s | Pass |

**Overall Result**: 10/10 tests passed ✅

**Total Duration**: 21m9.16825079s

*Test executed on August 11, 2026 at 13:17:54 KST (2026-08-11 13:17:54 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | highly-matched | Overall Match Rate: Min=97.7% Max=100.0% Avg=99.2% | VMs: 3 total, 3 matched, 0 acceptable",
      "targetCloud": {
        "csp": "gcp",
        "region": "asia-northeast3"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+e2-highcpu-2",
            "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 10,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+e2-standard-4",
            "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 10,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+e2-standard-2",
            "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 10,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "gcp-asia-northeast3",
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
        "connectionName": "gcp-asia-northeast3",
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
          "id": "gcp+asia-northeast3+e2-highcpu-2",
          "uid": "tb5vdhg2g76dc1jdvku8",
          "cspSpecName": "e2-highcpu-2",
          "name": "gcp+asia-northeast3+e2-highcpu-2",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1.953125,
          "diskSizeGB": -1,
          "costPerHour": 0.063531,
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
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "Efficient Instance, 2 vCPUs, 2 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "2"
            },
            {
              "key": "Id",
              "value": "337002"
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
              "value": "2048"
            },
            {
              "key": "Name",
              "value": "e2-highcpu-2"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        },
        {
          "id": "gcp+asia-northeast3+e2-standard-4",
          "uid": "tbjkou3d0g0gqutn1g4u",
          "cspSpecName": "e2-standard-4",
          "name": "gcp+asia-northeast3+e2-standard-4",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 15.625,
          "diskSizeGB": -1,
          "costPerHour": 0.171931,
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-4"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        },
        {
          "id": "gcp+asia-northeast3+e2-standard-2",
          "uid": "tbc2buo2dcu87tc74o4s",
          "cspSpecName": "e2-standard-2",
          "name": "gcp+asia-northeast3+e2-standard-2",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 7.8125,
          "diskSizeGB": -1,
          "costPerHour": 0.085966,
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-2"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "gcp",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "regionList": [
            "common"
          ],
          "id": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "uid": "tbgi22itnhj6oho5semi",
          "name": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "gcp-africa-south1",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:06:43 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
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
              "value": "3633714624"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-06-23T08:11:14.769-07:00"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
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
              "value": "768974430924119293"
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts"
            },
            {
              "key": "Name",
              "value": "ubuntu-2204-jammy-v20260623"
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623"
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
              "value": "eu; us; asia"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "sg-01",
          "connectionName": "gcp-asia-northeast3",
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
          "connectionName": "gcp-asia-northeast3",
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
          "connectionName": "gcp-asia-northeast3",
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
      ],
      "targetK8sCluster": {
        "connectionName": "",
        "description": "",
        "name": "",
        "version": "",
        "vNetId": "",
        "subnetIds": null,
        "securityGroupIds": null,
        "k8sNodeGroupList": null,
        "cspResourceId": "",
        "label": null,
        "systemLabel": ""
      }
    },
    {
      "status": "highly-matched",
      "description": "Candidate #2 | highly-matched | Overall Match Rate: Min=97.7% Max=100.0% Avg=99.2% | VMs: 3 total, 3 matched, 0 acceptable",
      "targetCloud": {
        "csp": "gcp",
        "region": "asia-northeast3"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+n2d-highcpu-2",
            "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-01"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 10,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+n2d-standard-4",
            "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-02"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 10,
            "dataDiskIds": null
          },
          {
            "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "nodeGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "gcp-asia-northeast3",
            "specId": "gcp+asia-northeast3+n2d-standard-2",
            "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
            "vNetId": "vnet-01",
            "subnetId": "subnet-01",
            "securityGroupIds": [
              "sg-03"
            ],
            "sshKeyId": "sshkey-01",
            "rootDiskSize": 10,
            "dataDiskIds": null
          }
        ],
        "policyOnPartialFailure": ""
      },
      "targetVNet": {
        "name": "vnet-01",
        "connectionName": "gcp-asia-northeast3",
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
        "connectionName": "gcp-asia-northeast3",
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
          "id": "gcp+asia-northeast3+e2-highcpu-2",
          "uid": "tb5vdhg2g76dc1jdvku8",
          "cspSpecName": "e2-highcpu-2",
          "name": "gcp+asia-northeast3+e2-highcpu-2",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1.953125,
          "diskSizeGB": -1,
          "costPerHour": 0.063531,
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
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "Efficient Instance, 2 vCPUs, 2 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "2"
            },
            {
              "key": "Id",
              "value": "337002"
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
              "value": "2048"
            },
            {
              "key": "Name",
              "value": "e2-highcpu-2"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        },
        {
          "id": "gcp+asia-northeast3+e2-standard-4",
          "uid": "tbjkou3d0g0gqutn1g4u",
          "cspSpecName": "e2-standard-4",
          "name": "gcp+asia-northeast3+e2-standard-4",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 15.625,
          "diskSizeGB": -1,
          "costPerHour": 0.171931,
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-4"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        },
        {
          "id": "gcp+asia-northeast3+e2-standard-2",
          "uid": "tbc2buo2dcu87tc74o4s",
          "cspSpecName": "e2-standard-2",
          "name": "gcp+asia-northeast3+e2-standard-2",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 7.8125,
          "diskSizeGB": -1,
          "costPerHour": 0.085966,
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-2"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        },
        {
          "id": "gcp+asia-northeast3+n2d-highcpu-2",
          "uid": "tbsbe60fntfghnkvihur",
          "cspSpecName": "n2d-highcpu-2",
          "name": "gcp+asia-northeast3+n2d-highcpu-2",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1.953125,
          "diskSizeGB": -1,
          "costPerHour": 0.080366,
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
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "2 vCPUs 2 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "2"
            },
            {
              "key": "Id",
              "value": "910002"
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
              "value": "2048"
            },
            {
              "key": "Name",
              "value": "n2d-highcpu-2"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/n2d-highcpu-2"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        },
        {
          "id": "gcp+asia-northeast3+n2d-standard-4",
          "uid": "tbg6v2c6adbnbld6nc05",
          "cspSpecName": "n2d-standard-4",
          "name": "gcp+asia-northeast3+n2d-standard-4",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 15.625,
          "diskSizeGB": -1,
          "costPerHour": 0.217708,
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
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "4 vCPUs 16 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "4"
            },
            {
              "key": "Id",
              "value": "911004"
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
              "value": "n2d-standard-4"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/n2d-standard-4"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        },
        {
          "id": "gcp+asia-northeast3+n2d-standard-2",
          "uid": "tbav7keo61nuc30mlnij",
          "cspSpecName": "n2d-standard-2",
          "name": "gcp+asia-northeast3+n2d-standard-2",
          "namespace": "system",
          "connectionName": "gcp-asia-northeast3",
          "providerName": "gcp",
          "regionName": "asia-northeast3",
          "regionLatitude": 37.2,
          "regionLongitude": 127,
          "infraType": "node",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 7.8125,
          "diskSizeGB": -1,
          "costPerHour": 0.108854,
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
              "key": "CreationTimestamp",
              "value": "1969-12-31T16:00:00.000-08:00"
            },
            {
              "key": "Description",
              "value": "2 vCPUs 8 GB RAM"
            },
            {
              "key": "GuestCpus",
              "value": "2"
            },
            {
              "key": "Id",
              "value": "911002"
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
              "value": "n2d-standard-2"
            },
            {
              "key": "SelfLink",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/n2d-standard-2"
            },
            {
              "key": "Zone",
              "value": "asia-northeast3-a"
            }
          ]
        }
      ],
      "targetOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "gcp",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "regionList": [
            "common"
          ],
          "id": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "uid": "tbgi22itnhj6oho5semi",
          "name": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "gcp-africa-south1",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:06:43 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
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
              "value": "3633714624"
            },
            {
              "key": "CreationTimestamp",
              "value": "2026-06-23T08:11:14.769-07:00"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
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
              "value": "768974430924119293"
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts"
            },
            {
              "key": "Name",
              "value": "ubuntu-2204-jammy-v20260623"
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623"
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
              "value": "eu; us; asia"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "sg-01",
          "connectionName": "gcp-asia-northeast3",
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
          "connectionName": "gcp-asia-northeast3",
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
          "connectionName": "gcp-asia-northeast3",
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
      ],
      "targetK8sCluster": {
        "connectionName": "",
        "description": "",
        "name": "",
        "version": "",
        "vNetId": "",
        "subnetIds": null,
        "securityGroupIds": null,
        "k8sNodeGroupList": null,
        "cspResourceId": "",
        "label": null,
        "systemLabel": ""
      }
    }
  ]
}
```

</details>

### Test Case 2: Validate the target model for computing infra

#### 2.1 API Request Information

- **API Endpoint**: `POST /beetle/validation/ns/mig01/infra`
- **Purpose**: Validate the recommended target model before migration (name collisions, spec/image compatibility, resource availability)

**Request Body**:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "status": "highly-matched",
  "description": "Candidate #1 | highly-matched | Overall Match Rate: Min=97.7% Max=100.0% Avg=99.2% | VMs: 3 total, 3 matched, 0 acceptable",
  "targetCloud": {
    "csp": "gcp",
    "region": "asia-northeast3"
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
        "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
        "connectionName": "gcp-asia-northeast3",
        "specId": "gcp+asia-northeast3+e2-highcpu-2",
        "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-01"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 10,
        "dataDiskIds": null
      },
      {
        "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "nodeGroupSize": 1,
        "label": {
          "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
        },
        "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
        "connectionName": "gcp-asia-northeast3",
        "specId": "gcp+asia-northeast3+e2-standard-4",
        "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-02"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 10,
        "dataDiskIds": null
      },
      {
        "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nodeGroupSize": 1,
        "label": {
          "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
        },
        "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
        "connectionName": "gcp-asia-northeast3",
        "specId": "gcp+asia-northeast3+e2-standard-2",
        "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "securityGroupIds": [
          "sg-03"
        ],
        "sshKeyId": "sshkey-01",
        "rootDiskSize": 10,
        "dataDiskIds": null
      }
    ],
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "vnet-01",
    "connectionName": "gcp-asia-northeast3",
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
    "connectionName": "gcp-asia-northeast3",
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
      "id": "gcp+asia-northeast3+e2-highcpu-2",
      "uid": "tb5vdhg2g76dc1jdvku8",
      "cspSpecName": "e2-highcpu-2",
      "name": "gcp+asia-northeast3+e2-highcpu-2",
      "namespace": "system",
      "connectionName": "gcp-asia-northeast3",
      "providerName": "gcp",
      "regionName": "asia-northeast3",
      "regionLatitude": 37.2,
      "regionLongitude": 127,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 1.953125,
      "diskSizeGB": -1,
      "costPerHour": 0.063531,
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
          "key": "CreationTimestamp",
          "value": "1969-12-31T16:00:00.000-08:00"
        },
        {
          "key": "Description",
          "value": "Efficient Instance, 2 vCPUs, 2 GB RAM"
        },
        {
          "key": "GuestCpus",
          "value": "2"
        },
        {
          "key": "Id",
          "value": "337002"
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
          "value": "2048"
        },
        {
          "key": "Name",
          "value": "e2-highcpu-2"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
        },
        {
          "key": "Zone",
          "value": "asia-northeast3-a"
        }
      ]
    },
    {
      "id": "gcp+asia-northeast3+e2-standard-4",
      "uid": "tbjkou3d0g0gqutn1g4u",
      "cspSpecName": "e2-standard-4",
      "name": "gcp+asia-northeast3+e2-standard-4",
      "namespace": "system",
      "connectionName": "gcp-asia-northeast3",
      "providerName": "gcp",
      "regionName": "asia-northeast3",
      "regionLatitude": 37.2,
      "regionLongitude": 127,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 15.625,
      "diskSizeGB": -1,
      "costPerHour": 0.171931,
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Zone",
          "value": "asia-northeast3-a"
        }
      ]
    },
    {
      "id": "gcp+asia-northeast3+e2-standard-2",
      "uid": "tbc2buo2dcu87tc74o4s",
      "cspSpecName": "e2-standard-2",
      "name": "gcp+asia-northeast3+e2-standard-2",
      "namespace": "system",
      "connectionName": "gcp-asia-northeast3",
      "providerName": "gcp",
      "regionName": "asia-northeast3",
      "regionLatitude": 37.2,
      "regionLongitude": 127,
      "infraType": "node",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 7.8125,
      "diskSizeGB": -1,
      "costPerHour": 0.085966,
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-2"
        },
        {
          "key": "Zone",
          "value": "asia-northeast3-a"
        }
      ]
    }
  ],
  "targetOsImageList": [
    {
      "resourceType": "image",
      "namespace": "system",
      "providerName": "gcp",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "regionList": [
        "common"
      ],
      "id": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "uid": "tbgi22itnhj6oho5semi",
      "name": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "gcp-africa-south1",
      "infraType": "",
      "fetchedTime": "2026.06.29 18:06:43 Mon",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
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
          "value": "3633714624"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-06-23T08:11:14.769-07:00"
        },
        {
          "key": "Description",
          "value": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
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
          "value": "768974430924119293"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts"
        },
        {
          "key": "Name",
          "value": "ubuntu-2204-jammy-v20260623"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623"
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
          "value": "eu; us; asia"
        }
      ],
      "systemLabel": "",
      "description": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
      "commandHistory": null
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "sg-01",
      "connectionName": "gcp-asia-northeast3",
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
      "connectionName": "gcp-asia-northeast3",
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
      "connectionName": "gcp-asia-northeast3",
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
  ],
  "targetK8sCluster": {
    "connectionName": "",
    "description": "",
    "name": "",
    "version": "",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": null,
    "cspResourceId": "",
    "label": null,
    "systemLabel": ""
  }
}
```

</details>

#### 2.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Target model is valid, no issues found

**Response Body**:

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "valid": true,
  "issues": []
}
```

</details>

### Test Case 3: Migrate the computing infra as defined in the target model

#### 3.1 API Request Information

- **API Endpoint**: `POST /beetle/migration/ns/mig01/infra`
- **Purpose**: Create and migrate infrastructure based on recommendation
- **Namespace ID**: `mig01`
- **Request Body**: Uses the response from the previous recommendation step

#### 3.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infrastructure migration completed successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my-infra101",
  "uid": "tbst2r01jvvvkqaisr7m",
  "name": "my-infra101",
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
    "countReconciling": 0,
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
    "sys.id": "my-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbst2r01jvvvkqaisr7m"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbb2us79heiktq3ob488",
      "cspResourceName": "tbb2us79heiktq3ob488",
      "cspResourceId": "tbb2us79heiktq3ob488",
      "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-08-11 04:25:44",
      "label": {
        "keypair": "tbv35ke87t7qr6c98ip5",
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-08-11 04:25:44",
        "sys.cspResourceId": "tbb2us79heiktq3ob488",
        "sys.cspResourceName": "tbb2us79heiktq3ob488",
        "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tbb2us79heiktq3ob488",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "34.64.193.50",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.3",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 10,
      "RootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+e2-highcpu-2",
      "cspSpecName": "e2-highcpu-2",
      "spec": {
        "cspSpecName": "e2-highcpu-2",
        "vCPU": 2,
        "memoryGiB": 1.953125,
        "costPerHour": 0.063531
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "tb8cvocsl30p76r85miv",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "tb2doh1m827n7bbd2vep",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
      "nodeUserName": "cb-user",
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
          "value": "2026-08-10T21:25:00.451-07:00"
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
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbb2us79heiktq3ob488,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "vSO2jutUIOM="
        },
        {
          "key": "Id",
          "value": "4480321822488206067"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "UmlDyI5ts9Q="
        },
        {
          "key": "Labels",
          "value": "{keypair:tbv35ke87t7qr6c98ip5}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-08-10T21:25:08.564-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tbb2us79heiktq3ob488"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.193.50,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:iRGpEf8FN3Y=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbb2us79heiktq3ob488"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
          "value": "{fingerprint:hXKI-FDAfv4=,items:[tbg0dgj7fs23ald6d3h1]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tbt3lbatmtt3tl0gja2u",
      "cspResourceName": "tbt3lbatmtt3tl0gja2u",
      "cspResourceId": "tbt3lbatmtt3tl0gja2u",
      "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-08-11 04:25:52",
      "label": {
        "keypair": "tbv35ke87t7qr6c98ip5",
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-08-11 04:25:52",
        "sys.cspResourceId": "tbt3lbatmtt3tl0gja2u",
        "sys.cspResourceName": "tbt3lbatmtt3tl0gja2u",
        "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tbt3lbatmtt3tl0gja2u",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "34.47.86.251",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.2",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 10,
      "RootDeviceName": "persistent-disk-0",
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
      "spec": {
        "cspSpecName": "e2-standard-2",
        "vCPU": 2,
        "memoryGiB": 7.8125,
        "costPerHour": 0.085966
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "tb8cvocsl30p76r85miv",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "tb2doh1m827n7bbd2vep",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
      "nodeUserName": "cb-user",
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
          "value": "2026-08-10T21:25:00.609-07:00"
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
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbt3lbatmtt3tl0gja2u,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "S8Lrz__ggRE="
        },
        {
          "key": "Id",
          "value": "5111726197328780019"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "UmlDyI5ts9Q="
        },
        {
          "key": "Labels",
          "value": "{keypair:tbv35ke87t7qr6c98ip5}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-08-10T21:25:14.599-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-2"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tbt3lbatmtt3tl0gja2u"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.47.86.251,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:HbmVrpu7SiY=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbt3lbatmtt3tl0gja2u"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
          "value": "{fingerprint:tX4Ywt0zhO0=,items:[tbd49adsltghd6rc3is3]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tbu3ah1dbkptorqb9oi5",
      "cspResourceName": "tbu3ah1dbkptorqb9oi5",
      "cspResourceId": "tbu3ah1dbkptorqb9oi5",
      "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-08-11 04:25:48",
      "label": {
        "keypair": "tbv35ke87t7qr6c98ip5",
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-08-11 04:25:48",
        "sys.cspResourceId": "tbu3ah1dbkptorqb9oi5",
        "sys.cspResourceName": "tbu3ah1dbkptorqb9oi5",
        "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tbu3ah1dbkptorqb9oi5",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "8.230.13.164",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 10,
      "RootDeviceName": "persistent-disk-0",
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
      "spec": {
        "cspSpecName": "e2-standard-4",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.171931
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "tb8cvocsl30p76r85miv",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "tb2doh1m827n7bbd2vep",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
      "nodeUserName": "cb-user",
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
          "value": "2026-08-10T21:25:03.087-07:00"
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
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbu3ah1dbkptorqb9oi5,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "QEsgg_icBFs="
        },
        {
          "key": "Id",
          "value": "2682894231203683057"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "UmlDyI5ts9Q="
        },
        {
          "key": "Labels",
          "value": "{keypair:tbv35ke87t7qr6c98ip5}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-08-10T21:25:12.760-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tbu3ah1dbkptorqb9oi5"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:8.230.13.164,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:YMZ74xLl5Ek=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.4,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbu3ah1dbkptorqb9oi5"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
          "value": "{fingerprint:MHudoUWciaM=,items:[tbad4lrkikd7agjvj944]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
        }
      ]
    }
  ],
  "newNodeList": null,
  "postCommand": {
    "userName": "",
    "command": null
  },
  "postCommandResult": {
    "results": null
  }
}
```

</details>

### Test Case 4: Get a list of infras

#### 4.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra`
- **Purpose**: Retrieve all migrated cloud infrastructure instances
- **Namespace ID**: `mig01`
- **Request Body**: None (GET request)

#### 4.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infra list retrieved successfully

**Response Body**:

```json
{
  "infra": [
    {
      "resourceType": "infra",
      "id": "my-infra101",
      "uid": "tbst2r01jvvvkqaisr7m",
      "name": "my-infra101",
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
        "countReconciling": 0,
        "countUndefined": 0
      },
      "targetStatus": "None",
      "targetAction": "None",
      "installMonAgent": "",
      "configureCloudAdaptiveNetwork": "",
      "label": {
        "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
        "sys.id": "my-infra101",
        "sys.labelType": "infra",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-infra101",
        "sys.namespace": "mig01",
        "sys.uid": "tbst2r01jvvvkqaisr7m"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbb2us79heiktq3ob488",
          "cspResourceName": "tbb2us79heiktq3ob488",
          "cspResourceId": "tbb2us79heiktq3ob488",
          "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-08-11 04:25:44",
          "label": {
            "keypair": "tbv35ke87t7qr6c98ip5",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-08-11 04:25:44",
            "sys.cspResourceId": "tbb2us79heiktq3ob488",
            "sys.cspResourceName": "tbb2us79heiktq3ob488",
            "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my-subnet-01",
            "sys.uid": "tbb2us79heiktq3ob488",
            "sys.vNetId": "my-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.64.193.50",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.3",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": 10,
          "RootDeviceName": "persistent-disk-0",
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
          "specId": "gcp+asia-northeast3+e2-highcpu-2",
          "cspSpecName": "e2-highcpu-2",
          "spec": {
            "cspSpecName": "e2-highcpu-2",
            "vCPU": 2,
            "memoryGiB": 1.953125,
            "costPerHour": 0.063531
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
          },
          "vNetId": "my-vnet-01",
          "cspVNetId": "tb8cvocsl30p76r85miv",
          "subnetId": "my-subnet-01",
          "cspSubnetId": "tb2doh1m827n7bbd2vep",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-sshkey-01",
          "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
          "nodeUserName": "cb-user",
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
              "value": "2026-08-10T21:25:00.451-07:00"
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
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbb2us79heiktq3ob488,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "vSO2jutUIOM="
            },
            {
              "key": "Id",
              "value": "4480321822488206067"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "UmlDyI5ts9Q="
            },
            {
              "key": "Labels",
              "value": "{keypair:tbv35ke87t7qr6c98ip5}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2026-08-10T21:25:08.564-07:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "tbb2us79heiktq3ob488"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.193.50,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:iRGpEf8FN3Y=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
            },
            {
              "key": "ResourceStatus",
              "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbb2us79heiktq3ob488"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
              "value": "{fingerprint:hXKI-FDAfv4=,items:[tbg0dgj7fs23ald6d3h1]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tbt3lbatmtt3tl0gja2u",
          "cspResourceName": "tbt3lbatmtt3tl0gja2u",
          "cspResourceId": "tbt3lbatmtt3tl0gja2u",
          "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
          "createdTime": "2026-08-11 04:25:52",
          "label": {
            "keypair": "tbv35ke87t7qr6c98ip5",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-08-11 04:25:52",
            "sys.cspResourceId": "tbt3lbatmtt3tl0gja2u",
            "sys.cspResourceName": "tbt3lbatmtt3tl0gja2u",
            "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my-subnet-01",
            "sys.uid": "tbt3lbatmtt3tl0gja2u",
            "sys.vNetId": "my-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "34.47.86.251",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.2",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": 10,
          "RootDeviceName": "persistent-disk-0",
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
          "spec": {
            "cspSpecName": "e2-standard-2",
            "vCPU": 2,
            "memoryGiB": 7.8125,
            "costPerHour": 0.085966
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
          },
          "vNetId": "my-vnet-01",
          "cspVNetId": "tb8cvocsl30p76r85miv",
          "subnetId": "my-subnet-01",
          "cspSubnetId": "tb2doh1m827n7bbd2vep",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-sshkey-01",
          "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
          "nodeUserName": "cb-user",
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
              "value": "2026-08-10T21:25:00.609-07:00"
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
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbt3lbatmtt3tl0gja2u,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "S8Lrz__ggRE="
            },
            {
              "key": "Id",
              "value": "5111726197328780019"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "UmlDyI5ts9Q="
            },
            {
              "key": "Labels",
              "value": "{keypair:tbv35ke87t7qr6c98ip5}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2026-08-10T21:25:14.599-07:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-2"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "tbt3lbatmtt3tl0gja2u"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.47.86.251,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:HbmVrpu7SiY=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
            },
            {
              "key": "ResourceStatus",
              "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbt3lbatmtt3tl0gja2u"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
              "value": "{fingerprint:tX4Ywt0zhO0=,items:[tbd49adsltghd6rc3is3]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbu3ah1dbkptorqb9oi5",
          "cspResourceName": "tbu3ah1dbkptorqb9oi5",
          "cspResourceId": "tbu3ah1dbkptorqb9oi5",
          "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
          "createdTime": "2026-08-11 04:25:48",
          "label": {
            "keypair": "tbv35ke87t7qr6c98ip5",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2026-08-11 04:25:48",
            "sys.cspResourceId": "tbu3ah1dbkptorqb9oi5",
            "sys.cspResourceName": "tbu3ah1dbkptorqb9oi5",
            "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my-subnet-01",
            "sys.uid": "tbu3ah1dbkptorqb9oi5",
            "sys.vNetId": "my-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "asia-northeast3",
            "zone": "asia-northeast3-a"
          },
          "publicIP": "8.230.13.164",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.4",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": 10,
          "RootDeviceName": "persistent-disk-0",
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
          "spec": {
            "cspSpecName": "e2-standard-4",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.171931
          },
          "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "image": {
            "resourceType": "image",
            "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
          },
          "vNetId": "my-vnet-01",
          "cspVNetId": "tb8cvocsl30p76r85miv",
          "subnetId": "my-subnet-01",
          "cspSubnetId": "tb2doh1m827n7bbd2vep",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "my-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-sshkey-01",
          "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
          "nodeUserName": "cb-user",
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
              "value": "2026-08-10T21:25:03.087-07:00"
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
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbu3ah1dbkptorqb9oi5,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "QEsgg_icBFs="
            },
            {
              "key": "Id",
              "value": "2682894231203683057"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "UmlDyI5ts9Q="
            },
            {
              "key": "Labels",
              "value": "{keypair:tbv35ke87t7qr6c98ip5}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2026-08-10T21:25:12.760-07:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-4"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "tbu3ah1dbkptorqb9oi5"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:8.230.13.164,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:YMZ74xLl5Ek=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.4,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
            },
            {
              "key": "ResourceStatus",
              "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbu3ah1dbkptorqb9oi5"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
              "value": "{fingerprint:MHudoUWciaM=,items:[tbad4lrkikd7agjvj944]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
            }
          ]
        }
      ],
      "newNodeList": null,
      "postCommand": {
        "userName": "",
        "command": null
      },
      "postCommandResult": {
        "results": null
      }
    }
  ]
}
```

### Test Case 5: Get a list of infra IDs

#### 5.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra?option=id`
- **Purpose**: Retrieve infra IDs only (lightweight response)
- **Namespace ID**: `mig01`
- **Query Parameter**: `option=id`
- **Request Body**: None (GET request)

#### 5.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infra IDs retrieved successfully

**Response Body**:

```json
{
  "idList": [
    "my-infra101"
  ]
}
```

### Test Case 6: Get a specific infra

#### 6.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Retrieve detailed information for a specific infra
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The specific infra identifier
- **Request Body**: None (GET request)

#### 6.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infra details retrieved successfully

**Response Body**:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "infra",
  "id": "my-infra101",
  "uid": "tbst2r01jvvvkqaisr7m",
  "name": "my-infra101",
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
    "countReconciling": 0,
    "countUndefined": 0
  },
  "targetStatus": "None",
  "targetAction": "None",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
    "sys.id": "my-infra101",
    "sys.labelType": "infra",
    "sys.manager": "cb-tumblebug",
    "sys.name": "my-infra101",
    "sys.namespace": "mig01",
    "sys.uid": "tbst2r01jvvvkqaisr7m"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbb2us79heiktq3ob488",
      "cspResourceName": "tbb2us79heiktq3ob488",
      "cspResourceId": "tbb2us79heiktq3ob488",
      "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-08-11 04:25:44",
      "label": {
        "keypair": "tbv35ke87t7qr6c98ip5",
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-08-11 04:25:44",
        "sys.cspResourceId": "tbb2us79heiktq3ob488",
        "sys.cspResourceName": "tbb2us79heiktq3ob488",
        "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tbb2us79heiktq3ob488",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "34.64.193.50",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.3",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 10,
      "RootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+e2-highcpu-2",
      "cspSpecName": "e2-highcpu-2",
      "spec": {
        "cspSpecName": "e2-highcpu-2",
        "vCPU": 2,
        "memoryGiB": 1.953125,
        "costPerHour": 0.063531
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "tb8cvocsl30p76r85miv",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "tb2doh1m827n7bbd2vep",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
      "nodeUserName": "cb-user",
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
          "value": "2026-08-10T21:25:00.451-07:00"
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
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbb2us79heiktq3ob488,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "vSO2jutUIOM="
        },
        {
          "key": "Id",
          "value": "4480321822488206067"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "UmlDyI5ts9Q="
        },
        {
          "key": "Labels",
          "value": "{keypair:tbv35ke87t7qr6c98ip5}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-08-10T21:25:08.564-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tbb2us79heiktq3ob488"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.193.50,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:iRGpEf8FN3Y=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbb2us79heiktq3ob488"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
          "value": "{fingerprint:hXKI-FDAfv4=,items:[tbg0dgj7fs23ald6d3h1]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tbt3lbatmtt3tl0gja2u",
      "cspResourceName": "tbt3lbatmtt3tl0gja2u",
      "cspResourceId": "tbt3lbatmtt3tl0gja2u",
      "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-08-11 04:25:52",
      "label": {
        "keypair": "tbv35ke87t7qr6c98ip5",
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-08-11 04:25:52",
        "sys.cspResourceId": "tbt3lbatmtt3tl0gja2u",
        "sys.cspResourceName": "tbt3lbatmtt3tl0gja2u",
        "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tbt3lbatmtt3tl0gja2u",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "34.47.86.251",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.2",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 10,
      "RootDeviceName": "persistent-disk-0",
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
      "spec": {
        "cspSpecName": "e2-standard-2",
        "vCPU": 2,
        "memoryGiB": 7.8125,
        "costPerHour": 0.085966
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "tb8cvocsl30p76r85miv",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "tb2doh1m827n7bbd2vep",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
      "nodeUserName": "cb-user",
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
          "value": "2026-08-10T21:25:00.609-07:00"
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
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbt3lbatmtt3tl0gja2u,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "S8Lrz__ggRE="
        },
        {
          "key": "Id",
          "value": "5111726197328780019"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "UmlDyI5ts9Q="
        },
        {
          "key": "Labels",
          "value": "{keypair:tbv35ke87t7qr6c98ip5}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-08-10T21:25:14.599-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-2"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tbt3lbatmtt3tl0gja2u"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.47.86.251,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:HbmVrpu7SiY=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbt3lbatmtt3tl0gja2u"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
          "value": "{fingerprint:tX4Ywt0zhO0=,items:[tbd49adsltghd6rc3is3]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tbu3ah1dbkptorqb9oi5",
      "cspResourceName": "tbu3ah1dbkptorqb9oi5",
      "cspResourceId": "tbu3ah1dbkptorqb9oi5",
      "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-08-11 04:25:48",
      "label": {
        "keypair": "tbv35ke87t7qr6c98ip5",
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2026-08-11 04:25:48",
        "sys.cspResourceId": "tbu3ah1dbkptorqb9oi5",
        "sys.cspResourceName": "tbu3ah1dbkptorqb9oi5",
        "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tbu3ah1dbkptorqb9oi5",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "asia-northeast3",
        "zone": "asia-northeast3-a"
      },
      "publicIP": "8.230.13.164",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.4",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": 10,
      "RootDeviceName": "persistent-disk-0",
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
      "spec": {
        "cspSpecName": "e2-standard-4",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.171931
      },
      "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
      "image": {
        "resourceType": "image",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "tb8cvocsl30p76r85miv",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "tb2doh1m827n7bbd2vep",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "my-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
      "nodeUserName": "cb-user",
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
          "value": "2026-08-10T21:25:03.087-07:00"
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
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbu3ah1dbkptorqb9oi5,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "QEsgg_icBFs="
        },
        {
          "key": "Id",
          "value": "2682894231203683057"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "UmlDyI5ts9Q="
        },
        {
          "key": "Labels",
          "value": "{keypair:tbv35ke87t7qr6c98ip5}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2026-08-10T21:25:12.760-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "tbu3ah1dbkptorqb9oi5"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:8.230.13.164,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:YMZ74xLl5Ek=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.4,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
        },
        {
          "key": "ResourceStatus",
          "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbu3ah1dbkptorqb9oi5"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
          "value": "{fingerprint:MHudoUWciaM=,items:[tbad4lrkikd7agjvj944]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
        }
      ]
    }
  ],
  "newNodeList": null,
  "postCommand": {
    "userName": "",
    "command": null
  },
  "postCommandResult": {
    "results": null
  }
}
```

</details>

### Test Case 7: Remote Command Accessibility Check

#### 7.1 Test Information

- **Test Type**: SSH Connectivity Test for All VMs
- **Purpose**: Verify that all migrated VMs are accessible via SSH
- **Method**: Extract public IP and SSH key from MCI access info for each VM, then execute remote command
- **Command Executed**: `uname -a` (to verify system information)
- **Authentication**: SSH key-based authentication
- **Scope**: Tests all VMs across all subgroups in the MCI

#### 7.2 Test Result Information

- **Status**: ✅ **SUCCESS**
- **Result**: All VMs are accessible via SSH

**Complete Test Details**:

<details>
  <summary> <ins>Click to see detailed test results </ins> </summary>

```json
{
  "data": {
    "configureCloudAdaptiveNetwork": "",
    "description": "Recommended VMs comprising multi-cloud infrastructure",
    "id": "my-infra101",
    "installMonAgent": "",
    "label": {
      "sys.description": "Recommended VMs comprising multi-cloud infrastructure",
      "sys.id": "my-infra101",
      "sys.labelType": "infra",
      "sys.manager": "cb-tumblebug",
      "sys.name": "my-infra101",
      "sys.namespace": "mig01",
      "sys.uid": "tbst2r01jvvvkqaisr7m"
    },
    "name": "my-infra101",
    "newNodeList": null,
    "node": [
      {
        "RootDeviceName": "persistent-disk-0",
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
            "value": "2026-08-10T21:25:00.451-07:00"
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
            "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbb2us79heiktq3ob488,type:PERSISTENT}"
          },
          {
            "key": "Fingerprint",
            "value": "vSO2jutUIOM="
          },
          {
            "key": "Id",
            "value": "4480321822488206067"
          },
          {
            "key": "Kind",
            "value": "compute#instance"
          },
          {
            "key": "LabelFingerprint",
            "value": "UmlDyI5ts9Q="
          },
          {
            "key": "Labels",
            "value": "{keypair:tbv35ke87t7qr6c98ip5}"
          },
          {
            "key": "LastStartTimestamp",
            "value": "2026-08-10T21:25:08.564-07:00"
          },
          {
            "key": "MachineType",
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-highcpu-2"
          },
          {
            "key": "Metadata",
            "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
          },
          {
            "key": "Name",
            "value": "tbb2us79heiktq3ob488"
          },
          {
            "key": "NetworkInterfaces",
            "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.193.50,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:iRGpEf8FN3Y=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
          },
          {
            "key": "ResourceStatus",
            "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbb2us79heiktq3ob488"
          },
          {
            "key": "ServiceAccounts",
            "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
            "value": "{fingerprint:hXKI-FDAfv4=,items:[tbg0dgj7fs23ald6d3h1]}"
          },
          {
            "key": "Zone",
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
          }
        ],
        "connectionConfig": {
          "configName": "gcp-asia-northeast3",
          "credentialHolder": "admin",
          "credentialName": "gcp",
          "driverName": "gcp-driver-v1.0.so",
          "providerName": "gcp",
          "regionDetail": {
            "description": "Seoul South Korea",
            "location": {
              "display": "South Korea (Seoul)",
              "latitude": 37.2,
              "longitude": 127
            },
            "regionId": "asia-northeast3",
            "regionName": "asia-northeast3",
            "zones": [
              "asia-northeast3-a",
              "asia-northeast3-b",
              "asia-northeast3-c"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "asia-northeast3",
            "assignedZone": "asia-northeast3-a"
          },
          "regionZoneInfoName": "gcp-asia-northeast3",
          "verified": true
        },
        "connectionName": "gcp-asia-northeast3",
        "createdTime": "2026-08-11 04:25:44",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "cspResourceId": "tbb2us79heiktq3ob488",
        "cspResourceName": "tbb2us79heiktq3ob488",
        "cspSpecName": "e2-highcpu-2",
        "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
        "cspSubnetId": "tb2doh1m827n7bbd2vep",
        "cspVNetId": "tb8cvocsl30p76r85miv",
        "dataDiskIds": null,
        "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
        "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "image": {
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "osArchitecture": "x86_64",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "label": {
          "keypair": "tbv35ke87t7qr6c98ip5",
          "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
          "sys.connectionName": "gcp-asia-northeast3",
          "sys.createdTime": "2026-08-11 04:25:44",
          "sys.cspResourceId": "tbb2us79heiktq3ob488",
          "sys.cspResourceName": "tbb2us79heiktq3ob488",
          "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "sys.infraId": "my-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "sys.subnetId": "my-subnet-01",
          "sys.uid": "tbb2us79heiktq3ob488",
          "sys.vNetId": "my-vnet-01"
        },
        "location": {
          "display": "South Korea (Seoul)",
          "latitude": 37.2,
          "longitude": 127
        },
        "monAgentStatus": "notInstalled",
        "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "nic0",
        "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.3",
        "publicDNS": "",
        "publicIP": "34.64.193.50",
        "region": {
          "region": "asia-northeast3",
          "zone": "asia-northeast3-a"
        },
        "resourceType": "node",
        "rootDiskSize": 10,
        "rootDiskType": "pd-standard",
        "securityGroupIds": [
          "my-sg-01"
        ],
        "spec": {
          "costPerHour": 0.063531,
          "cspSpecName": "e2-highcpu-2",
          "memoryGiB": 1.953125,
          "vCPU": 2
        },
        "specId": "gcp+asia-northeast3+e2-highcpu-2",
        "sshKeyId": "my-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tbb2us79heiktq3ob488",
        "vNetId": "my-vnet-01"
      },
      {
        "RootDeviceName": "persistent-disk-0",
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
            "value": "2026-08-10T21:25:00.609-07:00"
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
            "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbt3lbatmtt3tl0gja2u,type:PERSISTENT}"
          },
          {
            "key": "Fingerprint",
            "value": "S8Lrz__ggRE="
          },
          {
            "key": "Id",
            "value": "5111726197328780019"
          },
          {
            "key": "Kind",
            "value": "compute#instance"
          },
          {
            "key": "LabelFingerprint",
            "value": "UmlDyI5ts9Q="
          },
          {
            "key": "Labels",
            "value": "{keypair:tbv35ke87t7qr6c98ip5}"
          },
          {
            "key": "LastStartTimestamp",
            "value": "2026-08-10T21:25:14.599-07:00"
          },
          {
            "key": "MachineType",
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-2"
          },
          {
            "key": "Metadata",
            "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
          },
          {
            "key": "Name",
            "value": "tbt3lbatmtt3tl0gja2u"
          },
          {
            "key": "NetworkInterfaces",
            "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.47.86.251,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:HbmVrpu7SiY=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
          },
          {
            "key": "ResourceStatus",
            "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbt3lbatmtt3tl0gja2u"
          },
          {
            "key": "ServiceAccounts",
            "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
            "value": "{fingerprint:tX4Ywt0zhO0=,items:[tbd49adsltghd6rc3is3]}"
          },
          {
            "key": "Zone",
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
          }
        ],
        "connectionConfig": {
          "configName": "gcp-asia-northeast3",
          "credentialHolder": "admin",
          "credentialName": "gcp",
          "driverName": "gcp-driver-v1.0.so",
          "providerName": "gcp",
          "regionDetail": {
            "description": "Seoul South Korea",
            "location": {
              "display": "South Korea (Seoul)",
              "latitude": 37.2,
              "longitude": 127
            },
            "regionId": "asia-northeast3",
            "regionName": "asia-northeast3",
            "zones": [
              "asia-northeast3-a",
              "asia-northeast3-b",
              "asia-northeast3-c"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "asia-northeast3",
            "assignedZone": "asia-northeast3-a"
          },
          "regionZoneInfoName": "gcp-asia-northeast3",
          "verified": true
        },
        "connectionName": "gcp-asia-northeast3",
        "createdTime": "2026-08-11 04:25:52",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "cspResourceId": "tbt3lbatmtt3tl0gja2u",
        "cspResourceName": "tbt3lbatmtt3tl0gja2u",
        "cspSpecName": "e2-standard-2",
        "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
        "cspSubnetId": "tb2doh1m827n7bbd2vep",
        "cspVNetId": "tb8cvocsl30p76r85miv",
        "dataDiskIds": null,
        "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
        "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "image": {
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "osArchitecture": "x86_64",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "label": {
          "keypair": "tbv35ke87t7qr6c98ip5",
          "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "sys.connectionName": "gcp-asia-northeast3",
          "sys.createdTime": "2026-08-11 04:25:52",
          "sys.cspResourceId": "tbt3lbatmtt3tl0gja2u",
          "sys.cspResourceName": "tbt3lbatmtt3tl0gja2u",
          "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "sys.infraId": "my-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "sys.subnetId": "my-subnet-01",
          "sys.uid": "tbt3lbatmtt3tl0gja2u",
          "sys.vNetId": "my-vnet-01"
        },
        "location": {
          "display": "South Korea (Seoul)",
          "latitude": 37.2,
          "longitude": 127
        },
        "monAgentStatus": "notInstalled",
        "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "nic0",
        "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.2",
        "publicDNS": "",
        "publicIP": "34.47.86.251",
        "region": {
          "region": "asia-northeast3",
          "zone": "asia-northeast3-a"
        },
        "resourceType": "node",
        "rootDiskSize": 10,
        "rootDiskType": "pd-standard",
        "securityGroupIds": [
          "my-sg-03"
        ],
        "spec": {
          "costPerHour": 0.085966,
          "cspSpecName": "e2-standard-2",
          "memoryGiB": 7.8125,
          "vCPU": 2
        },
        "specId": "gcp+asia-northeast3+e2-standard-2",
        "sshKeyId": "my-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tbt3lbatmtt3tl0gja2u",
        "vNetId": "my-vnet-01"
      },
      {
        "RootDeviceName": "persistent-disk-0",
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
            "value": "2026-08-10T21:25:03.087-07:00"
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
            "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:10,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/disks/tbu3ah1dbkptorqb9oi5,type:PERSISTENT}"
          },
          {
            "key": "Fingerprint",
            "value": "QEsgg_icBFs="
          },
          {
            "key": "Id",
            "value": "2682894231203683057"
          },
          {
            "key": "Kind",
            "value": "compute#instance"
          },
          {
            "key": "LabelFingerprint",
            "value": "UmlDyI5ts9Q="
          },
          {
            "key": "Labels",
            "value": "{keypair:tbv35ke87t7qr6c98ip5}"
          },
          {
            "key": "LastStartTimestamp",
            "value": "2026-08-10T21:25:12.760-07:00"
          },
          {
            "key": "MachineType",
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/machineTypes/e2-standard-4"
          },
          {
            "key": "Metadata",
            "value": "{fingerprint:Kfn36fy9Uvs=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC2RkDH28JuHg5mXM/717W6/qkb0o5HTlRQZJ9rLUiWgzoeN71faP2Mxc/y0jFSzGRSYbmXi5bqrxsvlkDu5OytZHSCSjzfvMDhptkPgMV02wRKl+DVXThE75HME+QeZV3fTHKESlLQY4dyEcZlenqoq97eJMiv0KeDilXJFvW5r8U6Ywn1kHsDDvAn+03Utg0ZvMYzd5smOvxrv8HvoXz0g/fKneAA8P4AFxCRyMo93ZwAJZ/u0vCtYbcGnnPfNciYBJqx9nSbft/tJrN9lcFlSVHeX5v86dljG6ZE2X5n/ngdirDA6lNWnGqoi2ah4FYcQtV3uNSP6YsN8bRK+K3LaaUfoqjTWVttoWxAGYlwLeVVRV6u86TZIww9TboyiY4rYufeAKl8SG7eMWPgZ/GAn0KcfSw3b4iBdVnA6HN3M6HMUltEYFKSE2Pmp9uyMmVSPjYVlFytuT8YzyWv4Meb4SNPpxhIrJP/gmOaxxoi2X5Z6R4g/+xtsjQAK5YHWw/TJcB8nUkMTEl58Hganksaju4Mep8+CoYJXuOc0LZtTOIsBtQEVK0d6L6qkybhia2jWSaF2zkCRP3ksCGuj5AlovwVAfPlL4Rhj/2IUdXV3hJOrCchiI4ue33S4kZCXIBgApKeb2JAD/Uckc3gN+N5mXgNn0mPSlBpdyS3bdiUkw== cb-user}],kind:compute#metadata}"
          },
          {
            "key": "Name",
            "value": "tbu3ah1dbkptorqb9oi5"
          },
          {
            "key": "NetworkInterfaces",
            "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:8.230.13.164,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:YMZ74xLl5Ek=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb8cvocsl30p76r85miv,networkIP:10.0.1.4,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tb2doh1m827n7bbd2vep}"
          },
          {
            "key": "ResourceStatus",
            "value": "{effectiveInstanceMetadata:{vmDnsSettingMetadataValue:ZonalOnly}}"
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
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instances/tbu3ah1dbkptorqb9oi5"
          },
          {
            "key": "ServiceAccounts",
            "value": "{email:MASKED_EMAIL,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
            "value": "{fingerprint:MHudoUWciaM=,items:[tbad4lrkikd7agjvj944]}"
          },
          {
            "key": "Zone",
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a"
          }
        ],
        "connectionConfig": {
          "configName": "gcp-asia-northeast3",
          "credentialHolder": "admin",
          "credentialName": "gcp",
          "driverName": "gcp-driver-v1.0.so",
          "providerName": "gcp",
          "regionDetail": {
            "description": "Seoul South Korea",
            "location": {
              "display": "South Korea (Seoul)",
              "latitude": 37.2,
              "longitude": 127
            },
            "regionId": "asia-northeast3",
            "regionName": "asia-northeast3",
            "zones": [
              "asia-northeast3-a",
              "asia-northeast3-b",
              "asia-northeast3-c"
            ]
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "asia-northeast3",
            "assignedZone": "asia-northeast3-a"
          },
          "regionZoneInfoName": "gcp-asia-northeast3",
          "verified": true
        },
        "connectionName": "gcp-asia-northeast3",
        "createdTime": "2026-08-11 04:25:48",
        "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "cspResourceId": "tbu3ah1dbkptorqb9oi5",
        "cspResourceName": "tbu3ah1dbkptorqb9oi5",
        "cspSpecName": "e2-standard-4",
        "cspSshKeyId": "tbv35ke87t7qr6c98ip5",
        "cspSubnetId": "tb2doh1m827n7bbd2vep",
        "cspVNetId": "tb8cvocsl30p76r85miv",
        "dataDiskIds": null,
        "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
        "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "image": {
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
          "osArchitecture": "x86_64",
          "osDistribution": "Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623",
        "label": {
          "keypair": "tbv35ke87t7qr6c98ip5",
          "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "sys.connectionName": "gcp-asia-northeast3",
          "sys.createdTime": "2026-08-11 04:25:48",
          "sys.cspResourceId": "tbu3ah1dbkptorqb9oi5",
          "sys.cspResourceName": "tbu3ah1dbkptorqb9oi5",
          "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "sys.infraId": "my-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "sys.subnetId": "my-subnet-01",
          "sys.uid": "tbu3ah1dbkptorqb9oi5",
          "sys.vNetId": "my-vnet-01"
        },
        "location": {
          "display": "South Korea (Seoul)",
          "latitude": 37.2,
          "longitude": 127
        },
        "monAgentStatus": "notInstalled",
        "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "nic0",
        "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.4",
        "publicDNS": "",
        "publicIP": "8.230.13.164",
        "region": {
          "region": "asia-northeast3",
          "zone": "asia-northeast3-a"
        },
        "resourceType": "node",
        "rootDiskSize": 10,
        "rootDiskType": "pd-standard",
        "securityGroupIds": [
          "my-sg-02"
        ],
        "spec": {
          "costPerHour": 0.171931,
          "cspSpecName": "e2-standard-4",
          "memoryGiB": 15.625,
          "vCPU": 4
        },
        "specId": "gcp+asia-northeast3+e2-standard-4",
        "sshKeyId": "my-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tbu3ah1dbkptorqb9oi5",
        "vNetId": "my-vnet-01"
      }
    ],
    "postCommand": {
      "command": null,
      "userName": ""
    },
    "postCommandResult": {
      "results": null
    },
    "resourceType": "infra",
    "status": "Running:3 (R:3/3)",
    "statusCount": {
      "countCreating": 0,
      "countFailed": 0,
      "countRebooting": 0,
      "countReconciling": 0,
      "countRegistering": 0,
      "countResuming": 0,
      "countRunning": 3,
      "countSuspended": 0,
      "countSuspending": 0,
      "countTerminated": 0,
      "countTerminating": 0,
      "countTotal": 3,
      "countUndefined": 0
    },
    "systemLabel": "",
    "systemMessage": null,
    "targetAction": "None",
    "targetStatus": "None",
    "uid": "tbst2r01jvvvkqaisr7m"
  },
  "success": true
}
```

</details>

### Test Case 8: Target Infrastructure Summary

#### 8.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}?format=md`
- **Purpose**: Get a summary of the migrated target infrastructure in Markdown format
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier
- **Query Parameter**: `format=md`

#### 8.2 API Response Information

- **Status**: ✅ **SUCCESS**

### Test Case 9: Migration Report

#### 9.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Generate a comprehensive migration report matching source to target
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier

#### 9.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-08-11 04:26:29

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
| **Description** | Recommended VMs comprising multi-cloud infrastructure |
| **Status** | Running:3 (R:3/3) |
| **Target Cloud** | GCP |
| **Target Region** | asia-northeast3 |
| **Total VMs** | 3 |
| **Running VMs** | 3 |
| **Stopped VMs** | 0 |
| **Monitoring Agent** |  |

## Compute Resources

### VM Specifications

| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
|------|-------|--------------|-----|--------------|-----------|-----------------|---------------------|
| e2-highcpu-2 | 2 | 2.0 | - | x86_64 |  | $0.0635 | 1 |
| e2-standard-2 | 2 | 7.8 | - | x86_64 |  | $0.0860 | 1 |
| e2-standard-4 | 4 | 15.6 | - | x86_64 |  | $0.1719 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2204-jammy-v20260623 | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23 | Ubuntu 22.04 | Linux/UNIX | x86_64 | NA | 10 GB | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | tbb2us79heiktq3ob488 | Running | 2 vCPU, 2.0 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 34.64.193.50<br>**Private IP:** 10.0.1.3<br>**SGs:** my-sg-01<br>**SSH:** my-sshkey-01 |
| my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | tbt3lbatmtt3tl0gja2u | Running | 2 vCPU, 7.8 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 34.47.86.251<br>**Private IP:** 10.0.1.2<br>**SGs:** my-sg-03<br>**SSH:** my-sshkey-01 |
| my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | tbu3ah1dbkptorqb9oi5 | Running | 4 vCPU, 15.6 GiB | Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23 (Canonical, Ubuntu, 22.04 LTS, amd64 jammy image built on 2026-06-23) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 8.230.13.164<br>**Private IP:** 10.0.1.4<br>**SGs:** my-sg-02<br>**SSH:** my-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-vnet-01 |
| **CSP VNet ID** | tb8cvocsl30p76r85miv |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | gcp-asia-northeast3 |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-subnet-01 | tb2doh1m827n7bbd2vep | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-sshkey-01 | tbv35ke87t7qr6c98ip5 |  |  |

### Security Groups

#### Security Group: my-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-sg-01 |
| **CSP Security Group ID** | tbg0dgj7fs23ald6d3h1 |
| **VNet** | my-vnet-01 |
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
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-sg-02 |
| **CSP Security Group ID** | tbad4lrkikd7agjvj944 |
| **VNet** | my-vnet-01 |
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
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |

#### Security Group: my-sg-03

| Property | Value |
|----------|-------|
| **Name** | my-sg-03 |
| **CSP Security Group ID** | tbd49adsltghd6rc3is3 |
| **VNet** | my-vnet-01 |
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
| outbound | ALL |  | 0.0.0.0/0 |
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |


## Cost Estimation

### Total Cost Summary

| Period | Cost (USD) |
|--------|------------|
| **Per Hour** | $0.3214 |
| **Per Day** | $7.71 |
| **Per Month (30 days)** | $231.43 |

### Cost by Region

| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
|-----|--------|----------|-----------------|------------------|
| GCP | asia-northeast3 | 3 | $0.3214 | $231.43 |

### Cost by Virtual Machine

| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |
|---------|------|-----------------|------------------|
| my-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | e2-highcpu-2 | $0.0635 | $45.74 |
| my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | e2-standard-2 | $0.0860 | $61.90 |
| my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | e2-standard-4 | $0.1719 | $123.79 |




### Test Case 10: Delete the migrated computing infra

#### 10.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}`
- **Purpose**: Delete the migrated infrastructure and clean up resources
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{infraId}}` - The infra identifier to delete
- **Query Parameter**: `option=terminate` (terminates all resources)
- **Request Body**: None (DELETE request)

#### 10.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infrastructure deletion completed successfully

**Response Body**:

```json
{
  "message": "Infrastructure and resources deleted successfully (nsId: mig01, infraId: my-infra101)",
  "success": true
}
```

