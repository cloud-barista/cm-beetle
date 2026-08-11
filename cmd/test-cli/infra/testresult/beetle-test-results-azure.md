# CM-Beetle test results for AZURE

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.5.9+ (83b6bb2)
- imdl: v0.1.10+ (83b6bb2)
- CB-Tumblebug: v0.12.30
- CB-Spider: v0.12.42
- CB-MapUI: v0.12.56
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: August 11, 2026
- Test Time: 13:13:06 KST
- Test Execution: 2026-08-11 13:13:06 KST

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

## Test result for AZURE

### Test Results Summary

| Test | Step (Endpoint / Description) | Status | Duration | Details |
|------|-------------------------------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 3.507s | Pass |
| 2 | `POST /beetle/validation/ns/mig01/infra` | ✅ **PASS** | 446ms | Pass |
| 3 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m38.968s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 22ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 5ms | Pass |
| 6 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 13ms | Pass |
| 7 | Remote Command Accessibility Check | ✅ **PASS** | 6.902s | Pass |
| 8 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.355s | Pass |
| 9 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.333s | Pass |
| 10 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 2m10.478s | Pass |

**Overall Result**: 10/10 tests passed ✅

**Total Duration**: 5m16.29495601s

*Test executed on August 11, 2026 at 13:13:06 KST (2026-08-11 13:13:06 KST) using CM-Beetle automated test CLI*

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
      "description": "Candidate #1 | partially-matched | Overall Match Rate: Min=51.2% Max=100.0% Avg=90.9% | VMs: 3 total, 2 matched, 1 acceptable",
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2als_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b4as_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2as_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
          "uid": "tbp2modtn0bgras2hitq",
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
          "uid": "tbsncjt92f4d837mfelh",
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
          "uid": "tbe881brd6rbogviosto",
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "uid": "tbbh7g5gaht53run916b",
          "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:09:29 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
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
              "value": "0001-com-ubuntu-server-jammy"
            },
            {
              "key": "SKU",
              "value": "22_04-lts-gen2"
            },
            {
              "key": "Version",
              "value": "22.04.202606110"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202606110"
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
      "status": "partially-matched",
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=51.2% Max=100.0% Avg=90.9% | VMs: 3 total, 2 matched, 1 acceptable",
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2ls_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b4s_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2s_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
          "uid": "tbp2modtn0bgras2hitq",
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
          "uid": "tbsncjt92f4d837mfelh",
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
          "uid": "tbe881brd6rbogviosto",
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
          "uid": "tb8sm6o1g1ch6hjctbrc",
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
          "uid": "tb53tmibahilpks21ooj",
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
          "uid": "tbutu60t6pfqnc21a9nv",
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
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "uid": "tbbh7g5gaht53run916b",
          "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "sourceNodeUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.06.29 18:09:29 Mon",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "isBasicGpuImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
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
              "value": "0001-com-ubuntu-server-jammy"
            },
            {
              "key": "SKU",
              "value": "22_04-lts-gen2"
            },
            {
              "key": "Version",
              "value": "22.04.202606110"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202606110"
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
  "status": "partially-matched",
  "description": "Candidate #1 | partially-matched | Overall Match Rate: Min=51.2% Max=100.0% Avg=90.9% | VMs: 3 total, 2 matched, 1 acceptable",
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
        "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
        "connectionName": "azure-koreasouth",
        "specId": "azure+koreasouth+standard_b2als_v2",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
        "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
        "connectionName": "azure-koreasouth",
        "specId": "azure+koreasouth+standard_b4as_v2",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
        "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
        "connectionName": "azure-koreasouth",
        "specId": "azure+koreasouth+standard_b2as_v2",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
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
      "uid": "tbp2modtn0bgras2hitq",
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
      "uid": "tbsncjt92f4d837mfelh",
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
      "uid": "tbe881brd6rbogviosto",
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
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
      "regionList": [
        "common"
      ],
      "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
      "uid": "tbbh7g5gaht53run916b",
      "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
      "sourceNodeUid": "",
      "sourceCspImageName": "",
      "connectionName": "azure-australiacentral",
      "infraType": "",
      "fetchedTime": "2026.06.29 18:09:29 Mon",
      "creationDate": "",
      "isGPUImage": false,
      "isKubernetesImage": false,
      "isBasicImage": true,
      "isBasicGpuImage": false,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
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
          "value": "0001-com-ubuntu-server-jammy"
        },
        {
          "key": "SKU",
          "value": "22_04-lts-gen2"
        },
        {
          "key": "Version",
          "value": "22.04.202606110"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202606110"
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
  "uid": "tbtr7hpofcun7f9dda9v",
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
    "sys.uid": "tbtr7hpofcun7f9dda9v"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb47rdbchrml8ed3m04o",
      "cspResourceName": "tb47rdbchrml8ed3m04o",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o",
      "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-08-11 04:14:45",
      "label": {
        "createdBy": "tb47rdbchrml8ed3m04o",
        "keypair": "tbcl499f2quf7f8rrenr",
        "publicip": "tb47rdbchrml8ed3m04o-8703-PublicIP",
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-08-11 04:14:45",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o",
        "sys.cspResourceName": "tb47rdbchrml8ed3m04o",
        "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tb47rdbchrml8ed3m04o",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "52.231.152.188",
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
      "specId": "azure+koreasouth+standard_b2als_v2",
      "cspSpecName": "Standard_B2als_v2",
      "spec": {
        "cspSpecName": "Standard_B2als_v2",
        "vCPU": 2,
        "memoryGiB": 3.90625,
        "costPerHour": 0.0432
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
      "networkInterface": "tb47rdbchrml8ed3m04o-22528-VNic",
      "securityGroupIds": [
        "my-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
      "nodeUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb47rdbchrml8ed3m04o-22528-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb47rdbchrml8ed3m04o,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb47rdbchrml8ed3m04o_OsDisk_1_6efc4706b54a47b28cd7054122f62ae5,storageAccountType:Premium_LRS},name:tb47rdbchrml8ed3m04o_OsDisk_1_6efc4706b54a47b28cd7054122f62ae5,osType:Linux}},timeCreated:2026-08-11T04:13:58.0507992Z,vmId:365f0112-f31a-47a1-849a-13b78c4ac4dd}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tb47rdbchrml8ed3m04o,keypair:tbcl499f2quf7f8rrenr,publicip:tb47rdbchrml8ed3m04o-8703-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o"
        },
        {
          "key": "Name",
          "value": "tb47rdbchrml8ed3m04o"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tb764v6q5ovm4g4o6v0k",
      "cspResourceName": "tb764v6q5ovm4g4o6v0k",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k",
      "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-08-11 04:14:47",
      "label": {
        "createdBy": "tb764v6q5ovm4g4o6v0k",
        "keypair": "tbcl499f2quf7f8rrenr",
        "publicip": "tb764v6q5ovm4g4o6v0k-97332-PublicIP",
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-08-11 04:14:47",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k",
        "sys.cspResourceName": "tb764v6q5ovm4g4o6v0k",
        "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tb764v6q5ovm4g4o6v0k",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.131.21",
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
      "networkInterface": "tb764v6q5ovm4g4o6v0k-73764-VNic",
      "securityGroupIds": [
        "my-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
      "nodeUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb764v6q5ovm4g4o6v0k-73764-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb764v6q5ovm4g4o6v0k,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb764v6q5ovm4g4o6v0k_OsDisk_1_bfae7790b75a4f31a491c09e7add2c4e,storageAccountType:Premium_LRS},name:tb764v6q5ovm4g4o6v0k_OsDisk_1_bfae7790b75a4f31a491c09e7add2c4e,osType:Linux}},timeCreated:2026-08-11T04:14:00.8807571Z,vmId:0821cb1a-5223-4e21-a69e-ddf260e5afa8}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tb764v6q5ovm4g4o6v0k,keypair:tbcl499f2quf7f8rrenr,publicip:tb764v6q5ovm4g4o6v0k-97332-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k"
        },
        {
          "key": "Name",
          "value": "tb764v6q5ovm4g4o6v0k"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tb9adug1btujrrpl8oc6",
      "cspResourceName": "tb9adug1btujrrpl8oc6",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6",
      "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-08-11 04:14:44",
      "label": {
        "createdBy": "tb9adug1btujrrpl8oc6",
        "keypair": "tbcl499f2quf7f8rrenr",
        "publicip": "tb9adug1btujrrpl8oc6-12823-PublicIP",
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-08-11 04:14:44",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6",
        "sys.cspResourceName": "tb9adug1btujrrpl8oc6",
        "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tb9adug1btujrrpl8oc6",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.137.188",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "spec": {
        "cspSpecName": "Standard_B4as_v2",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.173
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
      "networkInterface": "tb9adug1btujrrpl8oc6-1194-VNic",
      "securityGroupIds": [
        "my-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
      "nodeUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb9adug1btujrrpl8oc6-1194-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb9adug1btujrrpl8oc6,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb9adug1btujrrpl8oc6_OsDisk_1_f956c3ea9c914d4b8b0a5fd2a9162097,storageAccountType:Premium_LRS},name:tb9adug1btujrrpl8oc6_OsDisk_1_f956c3ea9c914d4b8b0a5fd2a9162097,osType:Linux}},timeCreated:2026-08-11T04:13:56.6565532Z,vmId:2330d5c4-1e20-4f33-9c9a-468097897209}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tb9adug1btujrrpl8oc6,keypair:tbcl499f2quf7f8rrenr,publicip:tb9adug1btujrrpl8oc6-12823-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6"
        },
        {
          "key": "Name",
          "value": "tb9adug1btujrrpl8oc6"
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
      "uid": "tbtr7hpofcun7f9dda9v",
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
        "sys.uid": "tbtr7hpofcun7f9dda9v"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tb47rdbchrml8ed3m04o",
          "cspResourceName": "tb47rdbchrml8ed3m04o",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o",
          "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-08-11 04:14:45",
          "label": {
            "createdBy": "tb47rdbchrml8ed3m04o",
            "keypair": "tbcl499f2quf7f8rrenr",
            "publicip": "tb47rdbchrml8ed3m04o-8703-PublicIP",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-08-11 04:14:45",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o",
            "sys.cspResourceName": "tb47rdbchrml8ed3m04o",
            "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my-subnet-01",
            "sys.uid": "tb47rdbchrml8ed3m04o",
            "sys.vNetId": "my-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "52.231.152.188",
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
          "specId": "azure+koreasouth+standard_b2als_v2",
          "cspSpecName": "Standard_B2als_v2",
          "spec": {
            "cspSpecName": "Standard_B2als_v2",
            "vCPU": 2,
            "memoryGiB": 3.90625,
            "costPerHour": 0.0432
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110"
          },
          "vNetId": "my-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
          "subnetId": "my-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
          "networkInterface": "tb47rdbchrml8ed3m04o-22528-VNic",
          "securityGroupIds": [
            "my-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
          "nodeUserName": "cb-user",
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreasouth"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb47rdbchrml8ed3m04o-22528-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb47rdbchrml8ed3m04o,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb47rdbchrml8ed3m04o_OsDisk_1_6efc4706b54a47b28cd7054122f62ae5,storageAccountType:Premium_LRS},name:tb47rdbchrml8ed3m04o_OsDisk_1_6efc4706b54a47b28cd7054122f62ae5,osType:Linux}},timeCreated:2026-08-11T04:13:58.0507992Z,vmId:365f0112-f31a-47a1-849a-13b78c4ac4dd}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tb47rdbchrml8ed3m04o,keypair:tbcl499f2quf7f8rrenr,publicip:tb47rdbchrml8ed3m04o-8703-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o"
            },
            {
              "key": "Name",
              "value": "tb47rdbchrml8ed3m04o"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tb764v6q5ovm4g4o6v0k",
          "cspResourceName": "tb764v6q5ovm4g4o6v0k",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k",
          "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
          "createdTime": "2026-08-11 04:14:47",
          "label": {
            "createdBy": "tb764v6q5ovm4g4o6v0k",
            "keypair": "tbcl499f2quf7f8rrenr",
            "publicip": "tb764v6q5ovm4g4o6v0k-97332-PublicIP",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-08-11 04:14:47",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k",
            "sys.cspResourceName": "tb764v6q5ovm4g4o6v0k",
            "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my-subnet-01",
            "sys.uid": "tb764v6q5ovm4g4o6v0k",
            "sys.vNetId": "my-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.131.21",
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110"
          },
          "vNetId": "my-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
          "subnetId": "my-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
          "networkInterface": "tb764v6q5ovm4g4o6v0k-73764-VNic",
          "securityGroupIds": [
            "my-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
          "nodeUserName": "cb-user",
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreasouth"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb764v6q5ovm4g4o6v0k-73764-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb764v6q5ovm4g4o6v0k,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb764v6q5ovm4g4o6v0k_OsDisk_1_bfae7790b75a4f31a491c09e7add2c4e,storageAccountType:Premium_LRS},name:tb764v6q5ovm4g4o6v0k_OsDisk_1_bfae7790b75a4f31a491c09e7add2c4e,osType:Linux}},timeCreated:2026-08-11T04:14:00.8807571Z,vmId:0821cb1a-5223-4e21-a69e-ddf260e5afa8}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tb764v6q5ovm4g4o6v0k,keypair:tbcl499f2quf7f8rrenr,publicip:tb764v6q5ovm4g4o6v0k-97332-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k"
            },
            {
              "key": "Name",
              "value": "tb764v6q5ovm4g4o6v0k"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tb9adug1btujrrpl8oc6",
          "cspResourceName": "tb9adug1btujrrpl8oc6",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6",
          "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
          "createdTime": "2026-08-11 04:14:44",
          "label": {
            "createdBy": "tb9adug1btujrrpl8oc6",
            "keypair": "tbcl499f2quf7f8rrenr",
            "publicip": "tb9adug1btujrrpl8oc6-12823-PublicIP",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-08-11 04:14:44",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6",
            "sys.cspResourceName": "tb9adug1btujrrpl8oc6",
            "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my-subnet-01",
            "sys.uid": "tb9adug1btujrrpl8oc6",
            "sys.vNetId": "my-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.137.188",
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
          "specId": "azure+koreasouth+standard_b4as_v2",
          "cspSpecName": "Standard_B4as_v2",
          "spec": {
            "cspSpecName": "Standard_B4as_v2",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.173
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110"
          },
          "vNetId": "my-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
          "subnetId": "my-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
          "networkInterface": "tb9adug1btujrrpl8oc6-1194-VNic",
          "securityGroupIds": [
            "my-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
          "nodeUserName": "cb-user",
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreasouth"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb9adug1btujrrpl8oc6-1194-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb9adug1btujrrpl8oc6,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb9adug1btujrrpl8oc6_OsDisk_1_f956c3ea9c914d4b8b0a5fd2a9162097,storageAccountType:Premium_LRS},name:tb9adug1btujrrpl8oc6_OsDisk_1_f956c3ea9c914d4b8b0a5fd2a9162097,osType:Linux}},timeCreated:2026-08-11T04:13:56.6565532Z,vmId:2330d5c4-1e20-4f33-9c9a-468097897209}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tb9adug1btujrrpl8oc6,keypair:tbcl499f2quf7f8rrenr,publicip:tb9adug1btujrrpl8oc6-12823-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6"
            },
            {
              "key": "Name",
              "value": "tb9adug1btujrrpl8oc6"
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
  "uid": "tbtr7hpofcun7f9dda9v",
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
    "sys.uid": "tbtr7hpofcun7f9dda9v"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tb47rdbchrml8ed3m04o",
      "cspResourceName": "tb47rdbchrml8ed3m04o",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o",
      "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-08-11 04:14:45",
      "label": {
        "createdBy": "tb47rdbchrml8ed3m04o",
        "keypair": "tbcl499f2quf7f8rrenr",
        "publicip": "tb47rdbchrml8ed3m04o-8703-PublicIP",
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-08-11 04:14:45",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o",
        "sys.cspResourceName": "tb47rdbchrml8ed3m04o",
        "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tb47rdbchrml8ed3m04o",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "52.231.152.188",
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
      "specId": "azure+koreasouth+standard_b2als_v2",
      "cspSpecName": "Standard_B2als_v2",
      "spec": {
        "cspSpecName": "Standard_B2als_v2",
        "vCPU": 2,
        "memoryGiB": 3.90625,
        "costPerHour": 0.0432
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
      "networkInterface": "tb47rdbchrml8ed3m04o-22528-VNic",
      "securityGroupIds": [
        "my-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
      "nodeUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb47rdbchrml8ed3m04o-22528-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb47rdbchrml8ed3m04o,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb47rdbchrml8ed3m04o_OsDisk_1_6efc4706b54a47b28cd7054122f62ae5,storageAccountType:Premium_LRS},name:tb47rdbchrml8ed3m04o_OsDisk_1_6efc4706b54a47b28cd7054122f62ae5,osType:Linux}},timeCreated:2026-08-11T04:13:58.0507992Z,vmId:365f0112-f31a-47a1-849a-13b78c4ac4dd}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tb47rdbchrml8ed3m04o,keypair:tbcl499f2quf7f8rrenr,publicip:tb47rdbchrml8ed3m04o-8703-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o"
        },
        {
          "key": "Name",
          "value": "tb47rdbchrml8ed3m04o"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "tb764v6q5ovm4g4o6v0k",
      "cspResourceName": "tb764v6q5ovm4g4o6v0k",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k",
      "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-08-11 04:14:47",
      "label": {
        "createdBy": "tb764v6q5ovm4g4o6v0k",
        "keypair": "tbcl499f2quf7f8rrenr",
        "publicip": "tb764v6q5ovm4g4o6v0k-97332-PublicIP",
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-08-11 04:14:47",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k",
        "sys.cspResourceName": "tb764v6q5ovm4g4o6v0k",
        "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tb764v6q5ovm4g4o6v0k",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.131.21",
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
      "networkInterface": "tb764v6q5ovm4g4o6v0k-73764-VNic",
      "securityGroupIds": [
        "my-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
      "nodeUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb764v6q5ovm4g4o6v0k-73764-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb764v6q5ovm4g4o6v0k,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb764v6q5ovm4g4o6v0k_OsDisk_1_bfae7790b75a4f31a491c09e7add2c4e,storageAccountType:Premium_LRS},name:tb764v6q5ovm4g4o6v0k_OsDisk_1_bfae7790b75a4f31a491c09e7add2c4e,osType:Linux}},timeCreated:2026-08-11T04:14:00.8807571Z,vmId:0821cb1a-5223-4e21-a69e-ddf260e5afa8}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tb764v6q5ovm4g4o6v0k,keypair:tbcl499f2quf7f8rrenr,publicip:tb764v6q5ovm4g4o6v0k-97332-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k"
        },
        {
          "key": "Name",
          "value": "tb764v6q5ovm4g4o6v0k"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "node",
      "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "tb9adug1btujrrpl8oc6",
      "cspResourceName": "tb9adug1btujrrpl8oc6",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6",
      "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-08-11 04:14:44",
      "label": {
        "createdBy": "tb9adug1btujrrpl8oc6",
        "keypair": "tbcl499f2quf7f8rrenr",
        "publicip": "tb9adug1btujrrpl8oc6-12823-PublicIP",
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-08-11 04:14:44",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6",
        "sys.cspResourceName": "tb9adug1btujrrpl8oc6",
        "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my-subnet-01",
        "sys.uid": "tb9adug1btujrrpl8oc6",
        "sys.vNetId": "my-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.137.188",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "spec": {
        "cspSpecName": "Standard_B4as_v2",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.173
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110"
      },
      "vNetId": "my-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
      "subnetId": "my-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
      "networkInterface": "tb9adug1btujrrpl8oc6-1194-VNic",
      "securityGroupIds": [
        "my-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
      "nodeUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreasouth"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb9adug1btujrrpl8oc6-1194-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb9adug1btujrrpl8oc6,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb9adug1btujrrpl8oc6_OsDisk_1_f956c3ea9c914d4b8b0a5fd2a9162097,storageAccountType:Premium_LRS},name:tb9adug1btujrrpl8oc6_OsDisk_1_f956c3ea9c914d4b8b0a5fd2a9162097,osType:Linux}},timeCreated:2026-08-11T04:13:56.6565532Z,vmId:2330d5c4-1e20-4f33-9c9a-468097897209}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tb9adug1btujrrpl8oc6,keypair:tbcl499f2quf7f8rrenr,publicip:tb9adug1btujrrpl8oc6-12823-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6"
        },
        {
          "key": "Name",
          "value": "tb9adug1btujrrpl8oc6"
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
      "sys.uid": "tbtr7hpofcun7f9dda9v"
    },
    "name": "my-infra101",
    "newNodeList": null,
    "node": [
      {
        "RootDeviceName": "Not visible in Azure",
        "addtionalDetails": [
          {
            "key": "Location",
            "value": "koreasouth"
          },
          {
            "key": "Properties",
            "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb47rdbchrml8ed3m04o-22528-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb47rdbchrml8ed3m04o,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb47rdbchrml8ed3m04o_OsDisk_1_6efc4706b54a47b28cd7054122f62ae5,storageAccountType:Premium_LRS},name:tb47rdbchrml8ed3m04o_OsDisk_1_6efc4706b54a47b28cd7054122f62ae5,osType:Linux}},timeCreated:2026-08-11T04:13:58.0507992Z,vmId:365f0112-f31a-47a1-849a-13b78c4ac4dd}"
          },
          {
            "key": "Tags",
            "value": "{createdBy:tb47rdbchrml8ed3m04o,keypair:tbcl499f2quf7f8rrenr,publicip:tb47rdbchrml8ed3m04o-8703-PublicIP}"
          },
          {
            "key": "Etag",
            "value": "\\1\\"
          },
          {
            "key": "ID",
            "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o"
          },
          {
            "key": "Name",
            "value": "tb47rdbchrml8ed3m04o"
          },
          {
            "key": "Type",
            "value": "Microsoft.Compute/virtualMachines"
          }
        ],
        "connectionConfig": {
          "configName": "azure-koreasouth",
          "credentialHolder": "admin",
          "credentialName": "azure",
          "driverName": "azure-driver-v1.0.so",
          "providerName": "azure",
          "regionDetail": {
            "description": "Korea South",
            "location": {
              "display": "Korea South",
              "latitude": 35.1796,
              "longitude": 129.0756
            },
            "regionId": "koreasouth",
            "regionName": "koreasouth",
            "zones": []
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "koreasouth",
            "assignedZone": ""
          },
          "regionZoneInfoName": "azure-koreasouth",
          "verified": true
        },
        "connectionName": "azure-koreasouth",
        "createdTime": "2026-08-11 04:14:45",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
        "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o",
        "cspResourceName": "tb47rdbchrml8ed3m04o",
        "cspSpecName": "Standard_B2als_v2",
        "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
        "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
        "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
        "dataDiskIds": null,
        "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
        "id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "image": {
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "osArchitecture": "x86_64",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "label": {
          "createdBy": "tb47rdbchrml8ed3m04o",
          "keypair": "tbcl499f2quf7f8rrenr",
          "publicip": "tb47rdbchrml8ed3m04o-8703-PublicIP",
          "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
          "sys.connectionName": "azure-koreasouth",
          "sys.createdTime": "2026-08-11 04:14:45",
          "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o",
          "sys.cspResourceName": "tb47rdbchrml8ed3m04o",
          "sys.id": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "sys.infraId": "my-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "sys.subnetId": "my-subnet-01",
          "sys.uid": "tb47rdbchrml8ed3m04o",
          "sys.vNetId": "my-vnet-01"
        },
        "location": {
          "display": "Korea South",
          "latitude": 35.1796,
          "longitude": 129.0756
        },
        "monAgentStatus": "notInstalled",
        "name": "my-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "tb47rdbchrml8ed3m04o-22528-VNic",
        "nodeGroupId": "my-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.5",
        "publicDNS": "",
        "publicIP": "52.231.152.188",
        "region": {
          "region": "koreasouth"
        },
        "resourceType": "node",
        "rootDiskSize": 30,
        "rootDiskType": "PremiumSSD",
        "securityGroupIds": [
          "my-sg-01"
        ],
        "spec": {
          "costPerHour": 0.0432,
          "cspSpecName": "Standard_B2als_v2",
          "memoryGiB": 3.90625,
          "vCPU": 2
        },
        "specId": "azure+koreasouth+standard_b2als_v2",
        "sshKeyId": "my-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tb47rdbchrml8ed3m04o",
        "vNetId": "my-vnet-01"
      },
      {
        "RootDeviceName": "Not visible in Azure",
        "addtionalDetails": [
          {
            "key": "Location",
            "value": "koreasouth"
          },
          {
            "key": "Properties",
            "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb764v6q5ovm4g4o6v0k-73764-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb764v6q5ovm4g4o6v0k,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb764v6q5ovm4g4o6v0k_OsDisk_1_bfae7790b75a4f31a491c09e7add2c4e,storageAccountType:Premium_LRS},name:tb764v6q5ovm4g4o6v0k_OsDisk_1_bfae7790b75a4f31a491c09e7add2c4e,osType:Linux}},timeCreated:2026-08-11T04:14:00.8807571Z,vmId:0821cb1a-5223-4e21-a69e-ddf260e5afa8}"
          },
          {
            "key": "Tags",
            "value": "{createdBy:tb764v6q5ovm4g4o6v0k,keypair:tbcl499f2quf7f8rrenr,publicip:tb764v6q5ovm4g4o6v0k-97332-PublicIP}"
          },
          {
            "key": "Etag",
            "value": "\\1\\"
          },
          {
            "key": "ID",
            "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k"
          },
          {
            "key": "Name",
            "value": "tb764v6q5ovm4g4o6v0k"
          },
          {
            "key": "Type",
            "value": "Microsoft.Compute/virtualMachines"
          }
        ],
        "connectionConfig": {
          "configName": "azure-koreasouth",
          "credentialHolder": "admin",
          "credentialName": "azure",
          "driverName": "azure-driver-v1.0.so",
          "providerName": "azure",
          "regionDetail": {
            "description": "Korea South",
            "location": {
              "display": "Korea South",
              "latitude": 35.1796,
              "longitude": 129.0756
            },
            "regionId": "koreasouth",
            "regionName": "koreasouth",
            "zones": []
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "koreasouth",
            "assignedZone": ""
          },
          "regionZoneInfoName": "azure-koreasouth",
          "verified": true
        },
        "connectionName": "azure-koreasouth",
        "createdTime": "2026-08-11 04:14:47",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
        "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k",
        "cspResourceName": "tb764v6q5ovm4g4o6v0k",
        "cspSpecName": "Standard_B2as_v2",
        "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
        "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
        "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
        "dataDiskIds": null,
        "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
        "id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "image": {
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "osArchitecture": "x86_64",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "label": {
          "createdBy": "tb764v6q5ovm4g4o6v0k",
          "keypair": "tbcl499f2quf7f8rrenr",
          "publicip": "tb764v6q5ovm4g4o6v0k-97332-PublicIP",
          "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "sys.connectionName": "azure-koreasouth",
          "sys.createdTime": "2026-08-11 04:14:47",
          "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k",
          "sys.cspResourceName": "tb764v6q5ovm4g4o6v0k",
          "sys.id": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "sys.infraId": "my-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "sys.subnetId": "my-subnet-01",
          "sys.uid": "tb764v6q5ovm4g4o6v0k",
          "sys.vNetId": "my-vnet-01"
        },
        "location": {
          "display": "Korea South",
          "latitude": 35.1796,
          "longitude": 129.0756
        },
        "monAgentStatus": "notInstalled",
        "name": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "tb764v6q5ovm4g4o6v0k-73764-VNic",
        "nodeGroupId": "my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.6",
        "publicDNS": "",
        "publicIP": "20.200.131.21",
        "region": {
          "region": "koreasouth"
        },
        "resourceType": "node",
        "rootDiskSize": 30,
        "rootDiskType": "PremiumSSD",
        "securityGroupIds": [
          "my-sg-03"
        ],
        "spec": {
          "costPerHour": 0.0865,
          "cspSpecName": "Standard_B2as_v2",
          "memoryGiB": 7.8125,
          "vCPU": 2
        },
        "specId": "azure+koreasouth+standard_b2as_v2",
        "sshKeyId": "my-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tb764v6q5ovm4g4o6v0k",
        "vNetId": "my-vnet-01"
      },
      {
        "RootDeviceName": "Not visible in Azure",
        "addtionalDetails": [
          {
            "key": "Location",
            "value": "koreasouth"
          },
          {
            "key": "Properties",
            "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tb9adug1btujrrpl8oc6-1194-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tb9adug1btujrrpl8oc6,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDPp2diz8l9C6/IO9G4+sChs4NMA6wXJEAZzXTH0kqBNKkMemmNDytBAa8/YXGQMuB8MjCy5HoFz3ZDnoWpP8ZMzSjjLor5x6Mlxu2XObO0wcIGWSJHpPKFD0zGgCUaqshk1m7cQnRQthOitKL0Yy0kbL6AebRXydCMEsuvVrA5kvReBTqoxzdWQd7myz7DxRmRx1XQS18L34XUEGmBYi/+A9bLv39HIdSrAxPYBH4Ppo/9uSMfUkIA76m7/euLXMecQm1XyJbgJlLBBSxo8dssLSymNNph0dYBxGLnzghHl1mGHWDV9LSrQjXoQhlzIW7BHpVeeTmOajZsrn8qIdH/XWnO3jbp++QwpQE7zERvvOqzYA2tV0YSIuGn3NlwJoAP+li/ag2QlBunxV05NISpUfFiw/Et1/xbc1qbf4Ihrpad/wTwoREWBNv2uTWKTAiKCuPinGatP6txBzhyJYDCbHqxk4mGxd03sWEd+DghflrHzELdrVdfuCiRPPoFR2b+R1XuDw6TInHTabxSM6z3L7zY5QedaRoMNtqKStyT25z3QMc6p/wmhPHAZxTvtALpQFke4gHtfgBS63sQ9nn/xPJquV/uHE+IaL6GeQy7szAfMDxb0QxA1g4bbJeuC6PvE58SEo1jqCAupBjDBaXTIrp0DCxiI8dXtcTTXLZljQ==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,securityProfile:{securityType:Standard},storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202608060,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202608060},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tb9adug1btujrrpl8oc6_OsDisk_1_f956c3ea9c914d4b8b0a5fd2a9162097,storageAccountType:Premium_LRS},name:tb9adug1btujrrpl8oc6_OsDisk_1_f956c3ea9c914d4b8b0a5fd2a9162097,osType:Linux}},timeCreated:2026-08-11T04:13:56.6565532Z,vmId:2330d5c4-1e20-4f33-9c9a-468097897209}"
          },
          {
            "key": "Tags",
            "value": "{createdBy:tb9adug1btujrrpl8oc6,keypair:tbcl499f2quf7f8rrenr,publicip:tb9adug1btujrrpl8oc6-12823-PublicIP}"
          },
          {
            "key": "Etag",
            "value": "\\1\\"
          },
          {
            "key": "ID",
            "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6"
          },
          {
            "key": "Name",
            "value": "tb9adug1btujrrpl8oc6"
          },
          {
            "key": "Type",
            "value": "Microsoft.Compute/virtualMachines"
          }
        ],
        "connectionConfig": {
          "configName": "azure-koreasouth",
          "credentialHolder": "admin",
          "credentialName": "azure",
          "driverName": "azure-driver-v1.0.so",
          "providerName": "azure",
          "regionDetail": {
            "description": "Korea South",
            "location": {
              "display": "Korea South",
              "latitude": 35.1796,
              "longitude": 129.0756
            },
            "regionId": "koreasouth",
            "regionName": "koreasouth",
            "zones": []
          },
          "regionRepresentative": true,
          "regionZoneInfo": {
            "assignedRegion": "koreasouth",
            "assignedZone": ""
          },
          "regionZoneInfoName": "azure-koreasouth",
          "verified": true
        },
        "connectionName": "azure-koreasouth",
        "createdTime": "2026-08-11 04:14:44",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060",
        "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6",
        "cspResourceName": "tb9adug1btujrrpl8oc6",
        "cspSpecName": "Standard_B4as_v2",
        "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr",
        "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0",
        "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e",
        "dataDiskIds": null,
        "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
        "id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "image": {
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "osArchitecture": "x86_64",
          "osDistribution": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
          "osType": "Ubuntu 22.04",
          "resourceType": "image"
        },
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110",
        "label": {
          "createdBy": "tb9adug1btujrrpl8oc6",
          "keypair": "tbcl499f2quf7f8rrenr",
          "publicip": "tb9adug1btujrrpl8oc6-12823-PublicIP",
          "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "sys.connectionName": "azure-koreasouth",
          "sys.createdTime": "2026-08-11 04:14:44",
          "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6",
          "sys.cspResourceName": "tb9adug1btujrrpl8oc6",
          "sys.id": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "sys.infraId": "my-infra101",
          "sys.labelType": "node",
          "sys.manager": "cb-tumblebug",
          "sys.name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "sys.namespace": "mig01",
          "sys.nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "sys.subnetId": "my-subnet-01",
          "sys.uid": "tb9adug1btujrrpl8oc6",
          "sys.vNetId": "my-vnet-01"
        },
        "location": {
          "display": "Korea South",
          "latitude": 35.1796,
          "longitude": 129.0756
        },
        "monAgentStatus": "notInstalled",
        "name": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "networkAgentStatus": "notInstalled",
        "networkInterface": "tb9adug1btujrrpl8oc6-1194-VNic",
        "nodeGroupId": "my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "nodeUserName": "cb-user",
        "privateDNS": "",
        "privateIP": "10.0.1.4",
        "publicDNS": "",
        "publicIP": "20.200.137.188",
        "region": {
          "region": "koreasouth"
        },
        "resourceType": "node",
        "rootDiskSize": 30,
        "rootDiskType": "PremiumSSD",
        "securityGroupIds": [
          "my-sg-02"
        ],
        "spec": {
          "costPerHour": 0.173,
          "cspSpecName": "Standard_B4as_v2",
          "memoryGiB": 15.625,
          "vCPU": 4
        },
        "specId": "azure+koreasouth+standard_b4as_v2",
        "sshKeyId": "my-sshkey-01",
        "sshPort": 22,
        "status": "Running",
        "subnetId": "my-subnet-01",
        "systemMessage": "",
        "targetAction": "None",
        "targetStatus": "None",
        "uid": "tb9adug1btujrrpl8oc6",
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
    "uid": "tbtr7hpofcun7f9dda9v"
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

**Generated At:** 2026-08-11 04:15:23

**Namespace:** mig01

**Infra Name:** my-infra101

---

## Overview

| Property | Value |
|----------|-------|
| **Infra Name** | my-infra101 |
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
| Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110 | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 3 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb47rdbchrml8ed3m04o | Running | 2 vCPU, 3.9 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 52.231.152.188<br>**Private IP:** 10.0.1.5<br>**SGs:** my-sg-01<br>**SSH:** my-sshkey-01 |
| my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb764v6q5ovm4g4o6v0k | Running | 2 vCPU, 7.8 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 20.200.131.21<br>**Private IP:** 10.0.1.6<br>**SGs:** my-sg-03<br>**SSH:** my-sshkey-01 |
| my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tb9adug1btujrrpl8oc6 | Running | 4 vCPU, 15.6 GiB | Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110 (Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202606110) | **VNet:** my-vnet-01<br>**Subnet:** my-subnet-01<br>**Public IP:** 20.200.137.188<br>**Private IP:** 10.0.1.4<br>**SGs:** my-sg-02<br>**SSH:** my-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my-vnet-01 |
| **CSP VNet ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | azure-koreasouth |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my-subnet-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb71ak3s50q2mhtvng8e/subnets/tblvt7dndv3p4ij1t7p0 | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my-sshkey-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tbcl499f2quf7f8rrenr |  |  |

### Security Groups

#### Security Group: my-sg-01

| Property | Value |
|----------|-------|
| **Name** | my-sg-01 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tb4hht9gnvg1e6p3ubo9 |
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
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-sg-02

| Property | Value |
|----------|-------|
| **Name** | my-sg-02 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbhehlg26acmcjei88hm |
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
| outbound | TCP | 1-65535 | 0.0.0.0/0 |
| outbound | UDP | 1-65535 | 0.0.0.0/0 |
| outbound | ALL |  | 0.0.0.0/0 |

#### Security Group: my-sg-03

| Property | Value |
|----------|-------|
| **Name** | my-sg-03 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbu13bbjhkuuk2e5phft |
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
| my-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | Standard_B2als_v2 | $0.0432 | $31.10 |
| my-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | Standard_B2as_v2 | $0.0865 | $62.28 |
| my-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | Standard_B4as_v2 | $0.1730 | $124.56 |




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

