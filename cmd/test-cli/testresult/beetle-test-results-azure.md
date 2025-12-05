# CM-Beetle test results for AZURE

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.5+ (ed85eeb)
- cm-model: v0.0.14
- CB-Tumblebug: v0.11.19
- CB-Spider: v0.11.16
- CB-MapUI: v0.11.19
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: December 5, 2025
- Test Time: 15:32:57 KST
- Test Execution: 2025-12-05 15:32:57 KST

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
| 1 | `POST /beetle/recommendation/vmInfra` | ✅ **PASS** | 1.204s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 4m45.985s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 3.191s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 118ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 2.998s | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 2m39.942s | Pass |

**Overall Result**: 7/7 tests passed ✅

**Total Duration**: 9m17.281511641s

*Test executed on December 5, 2025 at 15:32:57 KST (2025-12-05 15:32:57 KST) using CM-Beetle automated test CLI*

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
  "items": [
    {
      "status": "partially-matched",
      "description": "Candidate #1 | partially-matched | Overall Match Rate: Min=51.2% Max=100.0% Avg=90.9% | VMs: 3 total, 2 matched, 1 acceptable",
      "targetCloud": {
        "csp": "azure",
        "region": "koreasouth"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
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
            "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
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
            "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
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
          "uid": "d3vjcu6qjs728pq3oru0",
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
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b4as_v2",
          "uid": "d3vjcu6qjs728pq3os00",
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
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b2as_v2",
          "uid": "d3vjcu6qjs728pq3orug",
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
          "uid": "d3vjf8uqjs728pqbmtf0",
          "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "connectionName": "azure-all",
          "infraType": "vm",
          "fetchedTime": "2025.10.27 09:13:07 Mon",
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
          "connectionName": "azure-koreasouth",
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
          "connectionName": "azure-koreasouth",
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
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=51.2% Max=100.0% Avg=90.9% | VMs: 3 total, 2 matched, 1 acceptable",
      "targetCloud": {
        "csp": "azure",
        "region": "koreasouth"
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
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2ls_v2",
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
            "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b4s_v2",
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
            "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "subGroupSize": "",
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2s_v2",
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
          "uid": "d3vjcu6qjs728pq3oru0",
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
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b4as_v2",
          "uid": "d3vjcu6qjs728pq3os00",
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
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b2as_v2",
          "uid": "d3vjcu6qjs728pq3orug",
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
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b2ls_v2",
          "uid": "d3vjcu6qjs728pq3opo0",
          "cspSpecName": "Standard_B2ls_v2",
          "name": "azure+koreasouth+standard_b2ls_v2",
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
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b4s_v2",
          "uid": "d3vjcu6qjs728pq3opq0",
          "cspSpecName": "Standard_B4s_v2",
          "name": "azure+koreasouth+standard_b4s_v2",
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
            }
          ]
        },
        {
          "id": "azure+koreasouth+standard_b2s_v2",
          "uid": "d3vjcu6qjs728pq3opog",
          "cspSpecName": "Standard_B2s_v2",
          "name": "azure+koreasouth+standard_b2s_v2",
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
          "uid": "d3vjf8uqjs728pqbmtf0",
          "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "connectionName": "azure-all",
          "infraType": "vm",
          "fetchedTime": "2025.10.27 09:13:07 Mon",
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
          "connectionName": "azure-koreasouth",
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
          "connectionName": "azure-koreasouth",
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
  "uid": "d4p7qi6chgjmd93omsf0",
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
    "sys.uid": "d4p7qi6chgjmd93omsf0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "d4p7qi6chgjmd93omsg0",
      "cspResourceName": "d4p7qi6chgjmd93omsg0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsg0",
      "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2025-12-05 06:36:51",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.227.120",
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
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0/subnets/d4p7p9echgjmd93omscg",
      "networkInterface": "d4p7qi6chgjmd93omsg0-42088-VNic",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4p7pc6chgjmd93omsd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-05T06:37:40Z",
          "completedTime": "2025-12-05T06:37:42Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4p7qi6chgjmd93omsg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4p7qi6chgjmd93omsg0-42088-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4p7qi6chgjmd93omsg0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDEiCMyDccYZ1X3UulBkc2nfkVarhrKm3MlRwWiVztCSBQ1WEZv1cH4/FNVKJy+ZFyg3Fw9Slp1NEPk49DD5iiYBLHqbYsD6VNy0Wv3lOYuNlAEuXpAmnJitkZ2uSrCXP/ldWVsRinhRjlsHRO7hbj/6HPIrB01lxisrygDrCqYCniAobpUbXVQFc4QeN2KcT+tndAWJX9I6JEM/3s3PYYW3fxyckWkODtb++YW3e7HgTX/jxv31XY25mM8XlJE0qvNnsta71Rmek7yAx8cbr5u2xSUFrK6pyGDSZcnqiv4Dm4rg8UxKus86I/QY1CdDfJWSRVtCK9xlHz5IFflX9jwTU8zqshvfGu8PZBHpzfTnj5kyw2Vnpcs1j9ZnlUSF8iAIs2PwGAMprZkUgYpXEJqmfqr6AsEP3KJKhGcHQy4Wud/OZkm9UmjkH7FdeXuhSkO36w4JQieLq1Pv0azXll5fTqkwj+JazB1QA/GZRHlw+rnFLAxDV3NnoLuRQV1/P5U625CHYYqOtE0iTvdYDyotSNG0OlXzzPlFXra/5RlmTvHx5DdjeQ92tgUeNaaQf8ncELAlFOl0GswEoRV8/aR7VvGes7STnsWkXbYdTlg4JP7C/IYKS9EpDnMmhIKSTwlApIsqwStELoP0e9TpBAIRActPsv8u3KnKddUkJIelw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4p7qi6chgjmd93omsg0_OsDisk_1_55300a34e5554fb9847153b633a84b38,storageAccountType:Premium_LRS},name:d4p7qi6chgjmd93omsg0_OsDisk_1_55300a34e5554fb9847153b633a84b38,osType:Linux}},timeCreated:2025-12-05T06:36:04.2776726Z,vmId:e52bd9af-9c0e-4208-9dcd-9e112cef4943}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4p7qi6chgjmd93omsg0,keypair:d4p7pc6chgjmd93omsd0,publicip:d4p7qi6chgjmd93omsg0-28815-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsg0"
        },
        {
          "key": "Name",
          "value": "d4p7qi6chgjmd93omsg0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "d4p7qi6chgjmd93omsi0",
      "cspResourceName": "d4p7qi6chgjmd93omsi0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsi0",
      "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2025-12-05 06:36:52",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.228.238",
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
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0/subnets/d4p7p9echgjmd93omscg",
      "networkInterface": "d4p7qi6chgjmd93omsi0-76862-VNic",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4p7pc6chgjmd93omsd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-05T06:37:40Z",
          "completedTime": "2025-12-05T06:37:44Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4p7qi6chgjmd93omsi0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4p7qi6chgjmd93omsi0-76862-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4p7qi6chgjmd93omsi0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDEiCMyDccYZ1X3UulBkc2nfkVarhrKm3MlRwWiVztCSBQ1WEZv1cH4/FNVKJy+ZFyg3Fw9Slp1NEPk49DD5iiYBLHqbYsD6VNy0Wv3lOYuNlAEuXpAmnJitkZ2uSrCXP/ldWVsRinhRjlsHRO7hbj/6HPIrB01lxisrygDrCqYCniAobpUbXVQFc4QeN2KcT+tndAWJX9I6JEM/3s3PYYW3fxyckWkODtb++YW3e7HgTX/jxv31XY25mM8XlJE0qvNnsta71Rmek7yAx8cbr5u2xSUFrK6pyGDSZcnqiv4Dm4rg8UxKus86I/QY1CdDfJWSRVtCK9xlHz5IFflX9jwTU8zqshvfGu8PZBHpzfTnj5kyw2Vnpcs1j9ZnlUSF8iAIs2PwGAMprZkUgYpXEJqmfqr6AsEP3KJKhGcHQy4Wud/OZkm9UmjkH7FdeXuhSkO36w4JQieLq1Pv0azXll5fTqkwj+JazB1QA/GZRHlw+rnFLAxDV3NnoLuRQV1/P5U625CHYYqOtE0iTvdYDyotSNG0OlXzzPlFXra/5RlmTvHx5DdjeQ92tgUeNaaQf8ncELAlFOl0GswEoRV8/aR7VvGes7STnsWkXbYdTlg4JP7C/IYKS9EpDnMmhIKSTwlApIsqwStELoP0e9TpBAIRActPsv8u3KnKddUkJIelw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4p7qi6chgjmd93omsi0_OsDisk_1_471d1d5af9b44b81a7ec416082ae0cd0,storageAccountType:Premium_LRS},name:d4p7qi6chgjmd93omsi0_OsDisk_1_471d1d5af9b44b81a7ec416082ae0cd0,osType:Linux}},timeCreated:2025-12-05T06:36:06.0433117Z,vmId:8e86e3c7-1263-44af-b37b-a794358b0b9b}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4p7qi6chgjmd93omsi0,keypair:d4p7pc6chgjmd93omsd0,publicip:d4p7qi6chgjmd93omsi0-20011-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsi0"
        },
        {
          "key": "Name",
          "value": "d4p7qi6chgjmd93omsi0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "d4p7qi6chgjmd93omsh0",
      "cspResourceName": "d4p7qi6chgjmd93omsh0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsh0",
      "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2025-12-05 06:36:54",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.229.55",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0/subnets/d4p7p9echgjmd93omscg",
      "networkInterface": "d4p7qi6chgjmd93omsh0-52129-VNic",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4p7pc6chgjmd93omsd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-05T06:37:40Z",
          "completedTime": "2025-12-05T06:37:44Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4p7qi6chgjmd93omsh0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4p7qi6chgjmd93omsh0-52129-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4p7qi6chgjmd93omsh0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDEiCMyDccYZ1X3UulBkc2nfkVarhrKm3MlRwWiVztCSBQ1WEZv1cH4/FNVKJy+ZFyg3Fw9Slp1NEPk49DD5iiYBLHqbYsD6VNy0Wv3lOYuNlAEuXpAmnJitkZ2uSrCXP/ldWVsRinhRjlsHRO7hbj/6HPIrB01lxisrygDrCqYCniAobpUbXVQFc4QeN2KcT+tndAWJX9I6JEM/3s3PYYW3fxyckWkODtb++YW3e7HgTX/jxv31XY25mM8XlJE0qvNnsta71Rmek7yAx8cbr5u2xSUFrK6pyGDSZcnqiv4Dm4rg8UxKus86I/QY1CdDfJWSRVtCK9xlHz5IFflX9jwTU8zqshvfGu8PZBHpzfTnj5kyw2Vnpcs1j9ZnlUSF8iAIs2PwGAMprZkUgYpXEJqmfqr6AsEP3KJKhGcHQy4Wud/OZkm9UmjkH7FdeXuhSkO36w4JQieLq1Pv0azXll5fTqkwj+JazB1QA/GZRHlw+rnFLAxDV3NnoLuRQV1/P5U625CHYYqOtE0iTvdYDyotSNG0OlXzzPlFXra/5RlmTvHx5DdjeQ92tgUeNaaQf8ncELAlFOl0GswEoRV8/aR7VvGes7STnsWkXbYdTlg4JP7C/IYKS9EpDnMmhIKSTwlApIsqwStELoP0e9TpBAIRActPsv8u3KnKddUkJIelw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4p7qi6chgjmd93omsh0_OsDisk_1_cbdb45dda2934f4996a6029aa29ed597,storageAccountType:Premium_LRS},name:d4p7qi6chgjmd93omsh0_OsDisk_1_cbdb45dda2934f4996a6029aa29ed597,osType:Linux}},timeCreated:2025-12-05T06:36:07.3402841Z,vmId:34fe32d6-32fb-42e2-8ede-08696c6fec8b}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4p7qi6chgjmd93omsh0,keypair:d4p7pc6chgjmd93omsd0,publicip:d4p7qi6chgjmd93omsh0-25478-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsh0"
        },
        {
          "key": "Name",
          "value": "d4p7qi6chgjmd93omsh0"
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
        "vmId": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vmIp": "52.231.227.120",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4p7qi6chgjmd93omsg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "vmIp": "52.231.229.55",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4p7qi6chgjmd93omsh0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "vmIp": "52.231.228.238",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4p7qi6chgjmd93omsi0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "d4p7qi6chgjmd93omsf0",
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
        "sys.uid": "d4p7qi6chgjmd93omsf0"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "d4p7qi6chgjmd93omsg0",
          "cspResourceName": "d4p7qi6chgjmd93omsg0",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsg0",
          "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
          "subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2025-12-05 06:36:51",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2025-12-05 06:36:51",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsg0",
            "sys.cspResourceName": "d4p7qi6chgjmd93omsg0",
            "sys.id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.uid": "d4p7qi6chgjmd93omsg0"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
          "region": {
            "Region": "koreasouth",
            "Zone": ""
          },
          "publicIP": "52.231.227.120",
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
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0/subnets/d4p7p9echgjmd93omscg",
          "networkInterface": "d4p7qi6chgjmd93omsg0-42088-VNic",
          "securityGroupIds": [
            "mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4p7pc6chgjmd93omsd0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-12-05T06:37:40Z",
              "completedTime": "2025-12-05T06:37:42Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4p7qi6chgjmd93omsg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4p7qi6chgjmd93omsg0-42088-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4p7qi6chgjmd93omsg0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDEiCMyDccYZ1X3UulBkc2nfkVarhrKm3MlRwWiVztCSBQ1WEZv1cH4/FNVKJy+ZFyg3Fw9Slp1NEPk49DD5iiYBLHqbYsD6VNy0Wv3lOYuNlAEuXpAmnJitkZ2uSrCXP/ldWVsRinhRjlsHRO7hbj/6HPIrB01lxisrygDrCqYCniAobpUbXVQFc4QeN2KcT+tndAWJX9I6JEM/3s3PYYW3fxyckWkODtb++YW3e7HgTX/jxv31XY25mM8XlJE0qvNnsta71Rmek7yAx8cbr5u2xSUFrK6pyGDSZcnqiv4Dm4rg8UxKus86I/QY1CdDfJWSRVtCK9xlHz5IFflX9jwTU8zqshvfGu8PZBHpzfTnj5kyw2Vnpcs1j9ZnlUSF8iAIs2PwGAMprZkUgYpXEJqmfqr6AsEP3KJKhGcHQy4Wud/OZkm9UmjkH7FdeXuhSkO36w4JQieLq1Pv0azXll5fTqkwj+JazB1QA/GZRHlw+rnFLAxDV3NnoLuRQV1/P5U625CHYYqOtE0iTvdYDyotSNG0OlXzzPlFXra/5RlmTvHx5DdjeQ92tgUeNaaQf8ncELAlFOl0GswEoRV8/aR7VvGes7STnsWkXbYdTlg4JP7C/IYKS9EpDnMmhIKSTwlApIsqwStELoP0e9TpBAIRActPsv8u3KnKddUkJIelw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4p7qi6chgjmd93omsg0_OsDisk_1_55300a34e5554fb9847153b633a84b38,storageAccountType:Premium_LRS},name:d4p7qi6chgjmd93omsg0_OsDisk_1_55300a34e5554fb9847153b633a84b38,osType:Linux}},timeCreated:2025-12-05T06:36:04.2776726Z,vmId:e52bd9af-9c0e-4208-9dcd-9e112cef4943}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d4p7qi6chgjmd93omsg0,keypair:d4p7pc6chgjmd93omsd0,publicip:d4p7qi6chgjmd93omsg0-28815-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsg0"
            },
            {
              "key": "Name",
              "value": "d4p7qi6chgjmd93omsg0"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "d4p7qi6chgjmd93omsi0",
          "cspResourceName": "d4p7qi6chgjmd93omsi0",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsi0",
          "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
          "createdTime": "2025-12-05 06:36:52",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2025-12-05 06:36:52",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsi0",
            "sys.cspResourceName": "d4p7qi6chgjmd93omsi0",
            "sys.id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.uid": "d4p7qi6chgjmd93omsi0"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
          "region": {
            "Region": "koreasouth",
            "Zone": ""
          },
          "publicIP": "52.231.228.238",
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
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0/subnets/d4p7p9echgjmd93omscg",
          "networkInterface": "d4p7qi6chgjmd93omsi0-76862-VNic",
          "securityGroupIds": [
            "mig-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4p7pc6chgjmd93omsd0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-12-05T06:37:40Z",
              "completedTime": "2025-12-05T06:37:44Z",
              "elapsedTime": 4,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4p7qi6chgjmd93omsi0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4p7qi6chgjmd93omsi0-76862-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4p7qi6chgjmd93omsi0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDEiCMyDccYZ1X3UulBkc2nfkVarhrKm3MlRwWiVztCSBQ1WEZv1cH4/FNVKJy+ZFyg3Fw9Slp1NEPk49DD5iiYBLHqbYsD6VNy0Wv3lOYuNlAEuXpAmnJitkZ2uSrCXP/ldWVsRinhRjlsHRO7hbj/6HPIrB01lxisrygDrCqYCniAobpUbXVQFc4QeN2KcT+tndAWJX9I6JEM/3s3PYYW3fxyckWkODtb++YW3e7HgTX/jxv31XY25mM8XlJE0qvNnsta71Rmek7yAx8cbr5u2xSUFrK6pyGDSZcnqiv4Dm4rg8UxKus86I/QY1CdDfJWSRVtCK9xlHz5IFflX9jwTU8zqshvfGu8PZBHpzfTnj5kyw2Vnpcs1j9ZnlUSF8iAIs2PwGAMprZkUgYpXEJqmfqr6AsEP3KJKhGcHQy4Wud/OZkm9UmjkH7FdeXuhSkO36w4JQieLq1Pv0azXll5fTqkwj+JazB1QA/GZRHlw+rnFLAxDV3NnoLuRQV1/P5U625CHYYqOtE0iTvdYDyotSNG0OlXzzPlFXra/5RlmTvHx5DdjeQ92tgUeNaaQf8ncELAlFOl0GswEoRV8/aR7VvGes7STnsWkXbYdTlg4JP7C/IYKS9EpDnMmhIKSTwlApIsqwStELoP0e9TpBAIRActPsv8u3KnKddUkJIelw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4p7qi6chgjmd93omsi0_OsDisk_1_471d1d5af9b44b81a7ec416082ae0cd0,storageAccountType:Premium_LRS},name:d4p7qi6chgjmd93omsi0_OsDisk_1_471d1d5af9b44b81a7ec416082ae0cd0,osType:Linux}},timeCreated:2025-12-05T06:36:06.0433117Z,vmId:8e86e3c7-1263-44af-b37b-a794358b0b9b}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d4p7qi6chgjmd93omsi0,keypair:d4p7pc6chgjmd93omsd0,publicip:d4p7qi6chgjmd93omsi0-20011-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsi0"
            },
            {
              "key": "Name",
              "value": "d4p7qi6chgjmd93omsi0"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "d4p7qi6chgjmd93omsh0",
          "cspResourceName": "d4p7qi6chgjmd93omsh0",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsh0",
          "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
          "createdTime": "2025-12-05 06:36:54",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2025-12-05 06:36:54",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsh0",
            "sys.cspResourceName": "d4p7qi6chgjmd93omsh0",
            "sys.id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.uid": "d4p7qi6chgjmd93omsh0"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
          "region": {
            "Region": "koreasouth",
            "Zone": ""
          },
          "publicIP": "52.231.229.55",
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
          "specId": "azure+koreasouth+standard_b4as_v2",
          "cspSpecName": "Standard_B4as_v2",
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0/subnets/d4p7p9echgjmd93omscg",
          "networkInterface": "d4p7qi6chgjmd93omsh0-52129-VNic",
          "securityGroupIds": [
            "mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4p7pc6chgjmd93omsd0",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-12-05T06:37:40Z",
              "completedTime": "2025-12-05T06:37:44Z",
              "elapsedTime": 4,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d4p7qi6chgjmd93omsh0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4p7qi6chgjmd93omsh0-52129-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4p7qi6chgjmd93omsh0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDEiCMyDccYZ1X3UulBkc2nfkVarhrKm3MlRwWiVztCSBQ1WEZv1cH4/FNVKJy+ZFyg3Fw9Slp1NEPk49DD5iiYBLHqbYsD6VNy0Wv3lOYuNlAEuXpAmnJitkZ2uSrCXP/ldWVsRinhRjlsHRO7hbj/6HPIrB01lxisrygDrCqYCniAobpUbXVQFc4QeN2KcT+tndAWJX9I6JEM/3s3PYYW3fxyckWkODtb++YW3e7HgTX/jxv31XY25mM8XlJE0qvNnsta71Rmek7yAx8cbr5u2xSUFrK6pyGDSZcnqiv4Dm4rg8UxKus86I/QY1CdDfJWSRVtCK9xlHz5IFflX9jwTU8zqshvfGu8PZBHpzfTnj5kyw2Vnpcs1j9ZnlUSF8iAIs2PwGAMprZkUgYpXEJqmfqr6AsEP3KJKhGcHQy4Wud/OZkm9UmjkH7FdeXuhSkO36w4JQieLq1Pv0azXll5fTqkwj+JazB1QA/GZRHlw+rnFLAxDV3NnoLuRQV1/P5U625CHYYqOtE0iTvdYDyotSNG0OlXzzPlFXra/5RlmTvHx5DdjeQ92tgUeNaaQf8ncELAlFOl0GswEoRV8/aR7VvGes7STnsWkXbYdTlg4JP7C/IYKS9EpDnMmhIKSTwlApIsqwStELoP0e9TpBAIRActPsv8u3KnKddUkJIelw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4p7qi6chgjmd93omsh0_OsDisk_1_cbdb45dda2934f4996a6029aa29ed597,storageAccountType:Premium_LRS},name:d4p7qi6chgjmd93omsh0_OsDisk_1_cbdb45dda2934f4996a6029aa29ed597,osType:Linux}},timeCreated:2025-12-05T06:36:07.3402841Z,vmId:34fe32d6-32fb-42e2-8ede-08696c6fec8b}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d4p7qi6chgjmd93omsh0,keypair:d4p7pc6chgjmd93omsd0,publicip:d4p7qi6chgjmd93omsh0-25478-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsh0"
            },
            {
              "key": "Name",
              "value": "d4p7qi6chgjmd93omsh0"
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
            "vmId": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
            "vmIp": "52.231.227.120",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4p7qi6chgjmd93omsg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "vmIp": "52.231.229.55",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4p7qi6chgjmd93omsh0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "vmIp": "52.231.228.238",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d4p7qi6chgjmd93omsi0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "d4p7qi6chgjmd93omsf0",
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
    "sys.uid": "d4p7qi6chgjmd93omsf0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "d4p7qi6chgjmd93omsg0",
      "cspResourceName": "d4p7qi6chgjmd93omsg0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsg0",
      "name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
      "subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2025-12-05 06:36:51",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2025-12-05 06:36:51",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsg0",
        "sys.cspResourceName": "d4p7qi6chgjmd93omsg0",
        "sys.id": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.uid": "d4p7qi6chgjmd93omsg0"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=90.6%",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.227.120",
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
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0/subnets/d4p7p9echgjmd93omscg",
      "networkInterface": "d4p7qi6chgjmd93omsg0-42088-VNic",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4p7pc6chgjmd93omsd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-05T06:37:40Z",
          "completedTime": "2025-12-05T06:37:42Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4p7qi6chgjmd93omsg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4p7qi6chgjmd93omsg0-42088-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4p7qi6chgjmd93omsg0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDEiCMyDccYZ1X3UulBkc2nfkVarhrKm3MlRwWiVztCSBQ1WEZv1cH4/FNVKJy+ZFyg3Fw9Slp1NEPk49DD5iiYBLHqbYsD6VNy0Wv3lOYuNlAEuXpAmnJitkZ2uSrCXP/ldWVsRinhRjlsHRO7hbj/6HPIrB01lxisrygDrCqYCniAobpUbXVQFc4QeN2KcT+tndAWJX9I6JEM/3s3PYYW3fxyckWkODtb++YW3e7HgTX/jxv31XY25mM8XlJE0qvNnsta71Rmek7yAx8cbr5u2xSUFrK6pyGDSZcnqiv4Dm4rg8UxKus86I/QY1CdDfJWSRVtCK9xlHz5IFflX9jwTU8zqshvfGu8PZBHpzfTnj5kyw2Vnpcs1j9ZnlUSF8iAIs2PwGAMprZkUgYpXEJqmfqr6AsEP3KJKhGcHQy4Wud/OZkm9UmjkH7FdeXuhSkO36w4JQieLq1Pv0azXll5fTqkwj+JazB1QA/GZRHlw+rnFLAxDV3NnoLuRQV1/P5U625CHYYqOtE0iTvdYDyotSNG0OlXzzPlFXra/5RlmTvHx5DdjeQ92tgUeNaaQf8ncELAlFOl0GswEoRV8/aR7VvGes7STnsWkXbYdTlg4JP7C/IYKS9EpDnMmhIKSTwlApIsqwStELoP0e9TpBAIRActPsv8u3KnKddUkJIelw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4p7qi6chgjmd93omsg0_OsDisk_1_55300a34e5554fb9847153b633a84b38,storageAccountType:Premium_LRS},name:d4p7qi6chgjmd93omsg0_OsDisk_1_55300a34e5554fb9847153b633a84b38,osType:Linux}},timeCreated:2025-12-05T06:36:04.2776726Z,vmId:e52bd9af-9c0e-4208-9dcd-9e112cef4943}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4p7qi6chgjmd93omsg0,keypair:d4p7pc6chgjmd93omsd0,publicip:d4p7qi6chgjmd93omsg0-28815-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsg0"
        },
        {
          "key": "Name",
          "value": "d4p7qi6chgjmd93omsg0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "d4p7qi6chgjmd93omsi0",
      "cspResourceName": "d4p7qi6chgjmd93omsi0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsi0",
      "name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2025-12-05 06:36:52",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2025-12-05 06:36:52",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsi0",
        "sys.cspResourceName": "d4p7qi6chgjmd93omsi0",
        "sys.id": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.uid": "d4p7qi6chgjmd93omsi0"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.228.238",
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
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0/subnets/d4p7p9echgjmd93omscg",
      "networkInterface": "d4p7qi6chgjmd93omsi0-76862-VNic",
      "securityGroupIds": [
        "mig-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4p7pc6chgjmd93omsd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-05T06:37:40Z",
          "completedTime": "2025-12-05T06:37:44Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4p7qi6chgjmd93omsi0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4p7qi6chgjmd93omsi0-76862-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4p7qi6chgjmd93omsi0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDEiCMyDccYZ1X3UulBkc2nfkVarhrKm3MlRwWiVztCSBQ1WEZv1cH4/FNVKJy+ZFyg3Fw9Slp1NEPk49DD5iiYBLHqbYsD6VNy0Wv3lOYuNlAEuXpAmnJitkZ2uSrCXP/ldWVsRinhRjlsHRO7hbj/6HPIrB01lxisrygDrCqYCniAobpUbXVQFc4QeN2KcT+tndAWJX9I6JEM/3s3PYYW3fxyckWkODtb++YW3e7HgTX/jxv31XY25mM8XlJE0qvNnsta71Rmek7yAx8cbr5u2xSUFrK6pyGDSZcnqiv4Dm4rg8UxKus86I/QY1CdDfJWSRVtCK9xlHz5IFflX9jwTU8zqshvfGu8PZBHpzfTnj5kyw2Vnpcs1j9ZnlUSF8iAIs2PwGAMprZkUgYpXEJqmfqr6AsEP3KJKhGcHQy4Wud/OZkm9UmjkH7FdeXuhSkO36w4JQieLq1Pv0azXll5fTqkwj+JazB1QA/GZRHlw+rnFLAxDV3NnoLuRQV1/P5U625CHYYqOtE0iTvdYDyotSNG0OlXzzPlFXra/5RlmTvHx5DdjeQ92tgUeNaaQf8ncELAlFOl0GswEoRV8/aR7VvGes7STnsWkXbYdTlg4JP7C/IYKS9EpDnMmhIKSTwlApIsqwStELoP0e9TpBAIRActPsv8u3KnKddUkJIelw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4p7qi6chgjmd93omsi0_OsDisk_1_471d1d5af9b44b81a7ec416082ae0cd0,storageAccountType:Premium_LRS},name:d4p7qi6chgjmd93omsi0_OsDisk_1_471d1d5af9b44b81a7ec416082ae0cd0,osType:Linux}},timeCreated:2025-12-05T06:36:06.0433117Z,vmId:8e86e3c7-1263-44af-b37b-a794358b0b9b}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4p7qi6chgjmd93omsi0,keypair:d4p7pc6chgjmd93omsd0,publicip:d4p7qi6chgjmd93omsi0-20011-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsi0"
        },
        {
          "key": "Name",
          "value": "d4p7qi6chgjmd93omsi0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "d4p7qi6chgjmd93omsh0",
      "cspResourceName": "d4p7qi6chgjmd93omsh0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsh0",
      "name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2025-12-05 06:36:54",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2025-12-05 06:36:54",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsh0",
        "sys.cspResourceName": "d4p7qi6chgjmd93omsh0",
        "sys.id": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.uid": "d4p7qi6chgjmd93omsh0"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=90.6%",
      "region": {
        "Region": "koreasouth",
        "Zone": ""
      },
      "publicIP": "52.231.229.55",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202510230",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d4p7p9echgjmd93omsc0/subnets/d4p7p9echgjmd93omscg",
      "networkInterface": "d4p7qi6chgjmd93omsh0-52129-VNic",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d4p7pc6chgjmd93omsd0",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-12-05T06:37:40Z",
          "completedTime": "2025-12-05T06:37:44Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d4p7qi6chgjmd93omsh0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d4p7qi6chgjmd93omsh0-52129-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d4p7qi6chgjmd93omsh0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDEiCMyDccYZ1X3UulBkc2nfkVarhrKm3MlRwWiVztCSBQ1WEZv1cH4/FNVKJy+ZFyg3Fw9Slp1NEPk49DD5iiYBLHqbYsD6VNy0Wv3lOYuNlAEuXpAmnJitkZ2uSrCXP/ldWVsRinhRjlsHRO7hbj/6HPIrB01lxisrygDrCqYCniAobpUbXVQFc4QeN2KcT+tndAWJX9I6JEM/3s3PYYW3fxyckWkODtb++YW3e7HgTX/jxv31XY25mM8XlJE0qvNnsta71Rmek7yAx8cbr5u2xSUFrK6pyGDSZcnqiv4Dm4rg8UxKus86I/QY1CdDfJWSRVtCK9xlHz5IFflX9jwTU8zqshvfGu8PZBHpzfTnj5kyw2Vnpcs1j9ZnlUSF8iAIs2PwGAMprZkUgYpXEJqmfqr6AsEP3KJKhGcHQy4Wud/OZkm9UmjkH7FdeXuhSkO36w4JQieLq1Pv0azXll5fTqkwj+JazB1QA/GZRHlw+rnFLAxDV3NnoLuRQV1/P5U625CHYYqOtE0iTvdYDyotSNG0OlXzzPlFXra/5RlmTvHx5DdjeQ92tgUeNaaQf8ncELAlFOl0GswEoRV8/aR7VvGes7STnsWkXbYdTlg4JP7C/IYKS9EpDnMmhIKSTwlApIsqwStELoP0e9TpBAIRActPsv8u3KnKddUkJIelw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202510230,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202510230},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d4p7qi6chgjmd93omsh0_OsDisk_1_cbdb45dda2934f4996a6029aa29ed597,storageAccountType:Premium_LRS},name:d4p7qi6chgjmd93omsh0_OsDisk_1_cbdb45dda2934f4996a6029aa29ed597,osType:Linux}},timeCreated:2025-12-05T06:36:07.3402841Z,vmId:34fe32d6-32fb-42e2-8ede-08696c6fec8b}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d4p7qi6chgjmd93omsh0,keypair:d4p7pc6chgjmd93omsd0,publicip:d4p7qi6chgjmd93omsh0-25478-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d4p7qi6chgjmd93omsh0"
        },
        {
          "key": "Name",
          "value": "d4p7qi6chgjmd93omsh0"
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
        "vmId": "migrated-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vmIp": "52.231.227.120",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4p7qi6chgjmd93omsg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "vmIp": "52.231.229.55",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4p7qi6chgjmd93omsh0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "vmIp": "52.231.228.238",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d4p7qi6chgjmd93omsi0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "output": "Linux d4p7qi6chgjmd93omsg0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.227.120",
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
      "output": "Linux d4p7qi6chgjmd93omsi0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.228.238",
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
      "output": "Linux d4p7qi6chgjmd93omsh0 6.8.0-1041-azure #47~22.04.1-Ubuntu SMP Fri Oct  3 20:43:01 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.229.55",
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

