# CM-Beetle test results for AZURE

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.5.0+ (cc39642)
- cm-model: v0.0.20
- CB-Tumblebug: v0.12.3
- CB-Spider: v0.12.11
- CB-MapUI: v0.12.16
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: March 25, 2026
- Test Time: 18:22:58 KST
- Test Execution: 2026-03-25 18:22:58 KST

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
| 1 | `POST /beetle/recommendation/vmInfra` | ✅ **PASS** | 3.238s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 2m39.161s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 499ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 7ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 634ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 2m2.175s | Pass |

**Overall Result**: 7/7 tests passed ✅

**Total Duration**: 6m6.726189825s

*Test executed on March 25, 2026 at 18:22:58 KST (2026-03-25 18:22:58 KST) using CM-Beetle automated test CLI*

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
  "nameSeed": "",
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
  "data": [
    {
      "nameSeed": "",
      "status": "partially-matched",
      "description": "Candidate #1 | partially-matched | Overall Match Rate: Min=51.2% Max=100.0% Avg=94.1% | VMs: 3 total, 2 matched, 1 acceptable",
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
            "name": "vm-ec268ed7-821e-9d73-e79f-961262161624",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2als_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
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
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b4as_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
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
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2as_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
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
      "targetVmSpecList": [
        {
          "id": "azure+koreasouth+standard_b2als_v2",
          "uid": "d66sfh5di7idhnupohb0",
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
          "uid": "d66sfh5di7idhnupohd0",
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
          "uid": "d66sfh5di7idhnupohbg",
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
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
          "uid": "d66t2jldi7idhnv2pe2g",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.02.12 13:10:38 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2",
          "osDiskType": "NA",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": null,
          "systemLabel": "",
          "description": "",
          "commandHistory": null
        },
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
          "uid": "d66t2jldi7idhnv2pe3g",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.02.12 13:10:38 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts",
          "osDiskType": "NA",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": null,
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
      "nameSeed": "",
      "status": "partially-matched",
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=51.2% Max=100.0% Avg=94.1% | VMs: 3 total, 2 matched, 1 acceptable",
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
            "name": "vm-ec268ed7-821e-9d73-e79f-961262161624",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2ls_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
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
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b4s_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
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
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
            "connectionName": "azure-koreasouth",
            "specId": "azure+koreasouth+standard_b2s_v2",
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
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
      "targetVmSpecList": [
        {
          "id": "azure+koreasouth+standard_b2als_v2",
          "uid": "d66sfh5di7idhnupohb0",
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
          "uid": "d66sfh5di7idhnupohd0",
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
          "uid": "d66sfh5di7idhnupohbg",
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
        },
        {
          "id": "azure+koreasouth+standard_b2ls_v2",
          "uid": "d66sfh5di7idhnupoes0",
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
          "uid": "d66sfh5di7idhnupoeu0",
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
          "uid": "d66sfh5di7idhnupoesg",
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
      "targetVmOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
          "uid": "d66t2jldi7idhnv2pe2g",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.02.12 13:10:38 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2",
          "osDiskType": "NA",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": null,
          "systemLabel": "",
          "description": "",
          "commandHistory": null
        },
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "azure",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
          "regionList": [
            "common"
          ],
          "id": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
          "uid": "d66t2jldi7idhnv2pe3g",
          "name": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "azure-australiacentral",
          "infraType": "",
          "fetchedTime": "2026.02.12 13:10:38 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": false,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts",
          "osDiskType": "NA",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": null,
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
  "uid": "d71qjd7693a119t3b8ig",
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
    "countRegistering": 0,
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
    "sys.uid": "d71qjd7693a119t3b8ig"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "d71qjd7693a119t3b8lg",
      "cspResourceName": "d71qjd7693a119t3b8lg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8lg",
      "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "subGroupId": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-03-25 09:25:02",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.184.197",
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts"
      },
      "vNetId": "vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg",
      "subnetId": "subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg/subnets/d71qj0f693a119t3b8g0",
      "networkInterface": "d71qjd7693a119t3b8lg-25225-VNic",
      "securityGroupIds": [
        "sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d71qj2f693a119t3b8gg",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBNweexaFlGZ1XB4c/3JIDrIdYpNjFazLQ0/SVDUgiUQGf0OaNnFbkuAJMOQK1aYqFtd4P9XtWYM5vP6guOHCrxI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:+h5v7NQYScWOmTbLSYLYmdZk0w80HyP3xsDAjijiB6g",
        "firstUsedAt": "2026-03-25T09:25:47Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T09:25:45Z",
          "completedTime": "2026-03-25T09:25:48Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d71qjd7693a119t3b8lg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d71qjd7693a119t3b8lg-25225-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d71qjd7693a119t3b8lg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQChJekscK7nwTaA8bC72oRqtAov0ToWZuHKBSTw4g6+zV4qyTIAH5VBfEVrVB4HMW2/xFH6lbC7F23s1fW8txyWnqe65kqeiTff164M65oBqeuRa57EyuQR9qZey4eoTfEYwjlpZtmLEsBTfgdhuwVZ3tArUdY8/45G+9SzFY+/5CBJLMphgaVPhE/u3lh5njMfDTljmxri7YETf7YwJZBkrX+mnKDrka4NCqSE+6YTmlhllPwO7uOlk85GC86mKNKQVdk3i7pNJcjl7AVXhghs5DV4eSrRAIdoOhjcsxFs8D9FasnTltJtU09pZfDFuJsYZjIQJileMSDt/hJd3kPmp6qDvYydUHc9pEZ5seXEWyHXibU5kwvLAFWlvPh66MNVAhzs7OVXkiZT9ktTzZUPIwQp1RA+o3GNnmpHzQEQZULBLkttl19T3pJb4mQaw/Jr3oOvdhR5I2KscmLeFS9LtSH02VHxPTHmJN4f6hN3zOWfAOHyXnpPWM6oM4RtGqOPwtjqTDd4HQkl1zYV8nsCO6RV+RMRHTMG6vL+XzhwYioNRuJauoQppx6BiJbUk0O2W9WP2uhm1D3/ZgTHfyXHOlVDKMW+63qF9TonY5aYHWmU+r4TpP5baM27Ze38LiR/pDzhhsk9eKFP3PeGUxKwRMpbEbwPit/vUg/LE2cf8w==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202601310,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202601310},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d71qjd7693a119t3b8lg_OsDisk_1_f8aee427cf8c4372bc8d02941a81f0ed,storageAccountType:Premium_LRS},name:d71qjd7693a119t3b8lg_OsDisk_1_f8aee427cf8c4372bc8d02941a81f0ed,osType:Linux}},timeCreated:2026-03-25T09:24:13.116244Z,vmId:6700becc-d787-4d20-af36-73d14dabf8d8}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d71qjd7693a119t3b8lg,keypair:d71qj2f693a119t3b8gg,publicip:d71qjd7693a119t3b8lg-4940-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8lg"
        },
        {
          "key": "Name",
          "value": "d71qjd7693a119t3b8lg"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "d71qjd7693a119t3b8kg",
      "cspResourceName": "d71qjd7693a119t3b8kg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8kg",
      "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "subGroupId": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-03-25 09:25:02",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.169.139",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "spec": {
        "cspSpecName": "Standard_B4as_v2",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.173
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2"
      },
      "vNetId": "vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg",
      "subnetId": "subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg/subnets/d71qj0f693a119t3b8g0",
      "networkInterface": "d71qjd7693a119t3b8kg-43320-VNic",
      "securityGroupIds": [
        "sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d71qj2f693a119t3b8gg",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBK1G/EPTKlJGA4doAZU6XyIIhzbzDibzPZA0lPaBhA78Y276e9O7bC4xrjt0kI5aZLNDWRdE0+wUO7snenZF0FA=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:Mwo0ByOzc/Jv/9gJ1ZP+3xUP9m3+gUFzVRrBzlzq+vk",
        "firstUsedAt": "2026-03-25T09:25:45Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T09:25:45Z",
          "completedTime": "2026-03-25T09:25:47Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d71qjd7693a119t3b8kg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d71qjd7693a119t3b8kg-43320-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d71qjd7693a119t3b8kg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQChJekscK7nwTaA8bC72oRqtAov0ToWZuHKBSTw4g6+zV4qyTIAH5VBfEVrVB4HMW2/xFH6lbC7F23s1fW8txyWnqe65kqeiTff164M65oBqeuRa57EyuQR9qZey4eoTfEYwjlpZtmLEsBTfgdhuwVZ3tArUdY8/45G+9SzFY+/5CBJLMphgaVPhE/u3lh5njMfDTljmxri7YETf7YwJZBkrX+mnKDrka4NCqSE+6YTmlhllPwO7uOlk85GC86mKNKQVdk3i7pNJcjl7AVXhghs5DV4eSrRAIdoOhjcsxFs8D9FasnTltJtU09pZfDFuJsYZjIQJileMSDt/hJd3kPmp6qDvYydUHc9pEZ5seXEWyHXibU5kwvLAFWlvPh66MNVAhzs7OVXkiZT9ktTzZUPIwQp1RA+o3GNnmpHzQEQZULBLkttl19T3pJb4mQaw/Jr3oOvdhR5I2KscmLeFS9LtSH02VHxPTHmJN4f6hN3zOWfAOHyXnpPWM6oM4RtGqOPwtjqTDd4HQkl1zYV8nsCO6RV+RMRHTMG6vL+XzhwYioNRuJauoQppx6BiJbUk0O2W9WP2uhm1D3/ZgTHfyXHOlVDKMW+63qF9TonY5aYHWmU+r4TpP5baM27Ze38LiR/pDzhhsk9eKFP3PeGUxKwRMpbEbwPit/vUg/LE2cf8w==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202601310,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202601310},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d71qjd7693a119t3b8kg_OsDisk_1_114ad1115c59441d956b29335b8e826d,storageAccountType:Premium_LRS},name:d71qjd7693a119t3b8kg_OsDisk_1_114ad1115c59441d956b29335b8e826d,osType:Linux}},timeCreated:2026-03-25T09:24:12.0977482Z,vmId:06fe8dad-90e6-4c19-bd37-cdb92def1b65}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d71qjd7693a119t3b8kg,keypair:d71qj2f693a119t3b8gg,publicip:d71qjd7693a119t3b8kg-10120-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8kg"
        },
        {
          "key": "Name",
          "value": "d71qjd7693a119t3b8kg"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "d71qjd7693a119t3b8jg",
      "cspResourceName": "d71qjd7693a119t3b8jg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8jg",
      "name": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "subGroupId": "vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-03-25 09:25:02",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.214.57.165",
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
      "specId": "azure+koreasouth+standard_b2als_v2",
      "cspSpecName": "Standard_B2als_v2",
      "spec": {
        "cspSpecName": "Standard_B2als_v2",
        "vCPU": 2,
        "memoryGiB": 3.90625,
        "costPerHour": 0.0432
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2"
      },
      "vNetId": "vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg",
      "subnetId": "subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg/subnets/d71qj0f693a119t3b8g0",
      "networkInterface": "d71qjd7693a119t3b8jg-13266-VNic",
      "securityGroupIds": [
        "sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d71qj2f693a119t3b8gg",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAaCHkSouaVM2JRH0A/WdA9IZoxc9BzMyci5iIN/EzwkVs+i9k2XGiSvzkYJLPR2flQlWyc3BZTo4yg8faHNAPg=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:q7FsLm+wFn70wol+Xs1+iamarU+xvkF8Uiei31aCuHQ",
        "firstUsedAt": "2026-03-25T09:25:47Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T09:25:45Z",
          "completedTime": "2026-03-25T09:25:50Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d71qjd7693a119t3b8jg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d71qjd7693a119t3b8jg-13266-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d71qjd7693a119t3b8jg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQChJekscK7nwTaA8bC72oRqtAov0ToWZuHKBSTw4g6+zV4qyTIAH5VBfEVrVB4HMW2/xFH6lbC7F23s1fW8txyWnqe65kqeiTff164M65oBqeuRa57EyuQR9qZey4eoTfEYwjlpZtmLEsBTfgdhuwVZ3tArUdY8/45G+9SzFY+/5CBJLMphgaVPhE/u3lh5njMfDTljmxri7YETf7YwJZBkrX+mnKDrka4NCqSE+6YTmlhllPwO7uOlk85GC86mKNKQVdk3i7pNJcjl7AVXhghs5DV4eSrRAIdoOhjcsxFs8D9FasnTltJtU09pZfDFuJsYZjIQJileMSDt/hJd3kPmp6qDvYydUHc9pEZ5seXEWyHXibU5kwvLAFWlvPh66MNVAhzs7OVXkiZT9ktTzZUPIwQp1RA+o3GNnmpHzQEQZULBLkttl19T3pJb4mQaw/Jr3oOvdhR5I2KscmLeFS9LtSH02VHxPTHmJN4f6hN3zOWfAOHyXnpPWM6oM4RtGqOPwtjqTDd4HQkl1zYV8nsCO6RV+RMRHTMG6vL+XzhwYioNRuJauoQppx6BiJbUk0O2W9WP2uhm1D3/ZgTHfyXHOlVDKMW+63qF9TonY5aYHWmU+r4TpP5baM27Ze38LiR/pDzhhsk9eKFP3PeGUxKwRMpbEbwPit/vUg/LE2cf8w==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202601310,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202601310},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d71qjd7693a119t3b8jg_OsDisk_1_8cc5ec47f72d4e4782d2e6dd82a81b05,storageAccountType:Premium_LRS},name:d71qjd7693a119t3b8jg_OsDisk_1_8cc5ec47f72d4e4782d2e6dd82a81b05,osType:Linux}},timeCreated:2026-03-25T09:24:12.0977482Z,vmId:c69d53c6-1b0f-4349-8426-768049d435fa}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d71qjd7693a119t3b8jg,keypair:d71qj2f693a119t3b8gg,publicip:d71qjd7693a119t3b8jg-20136-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8jg"
        },
        {
          "key": "Name",
          "value": "d71qjd7693a119t3b8jg"
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
        "vmId": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "vmIp": "20.200.169.139",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d71qjd7693a119t3b8kg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "vmIp": "20.200.184.197",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d71qjd7693a119t3b8lg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vmIp": "20.214.57.165",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d71qjd7693a119t3b8jg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "d71qjd7693a119t3b8ig",
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
        "countRegistering": 0,
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
        "sys.uid": "d71qjd7693a119t3b8ig"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "d71qjd7693a119t3b8jg",
          "cspResourceName": "d71qjd7693a119t3b8jg",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8jg",
          "name": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "subGroupId": "vm-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-03-25 09:25:02",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-03-25 09:25:02",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8jg",
            "sys.cspResourceName": "d71qjd7693a119t3b8jg",
            "sys.id": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "subnet-01",
            "sys.uid": "d71qjd7693a119t3b8jg",
            "sys.vNetId": "vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.57.165",
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
          "specId": "azure+koreasouth+standard_b2als_v2",
          "cspSpecName": "Standard_B2als_v2",
          "spec": {
            "cspSpecName": "Standard_B2als_v2",
            "vCPU": 2,
            "memoryGiB": 3.90625,
            "costPerHour": 0.0432
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2"
          },
          "vNetId": "vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg",
          "subnetId": "subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg/subnets/d71qj0f693a119t3b8g0",
          "networkInterface": "d71qjd7693a119t3b8jg-13266-VNic",
          "securityGroupIds": [
            "sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d71qj2f693a119t3b8gg",
          "vmUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAaCHkSouaVM2JRH0A/WdA9IZoxc9BzMyci5iIN/EzwkVs+i9k2XGiSvzkYJLPR2flQlWyc3BZTo4yg8faHNAPg=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:q7FsLm+wFn70wol+Xs1+iamarU+xvkF8Uiei31aCuHQ",
            "firstUsedAt": "2026-03-25T09:25:47Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-03-25T09:25:45Z",
              "completedTime": "2026-03-25T09:25:50Z",
              "elapsedTime": 5,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d71qjd7693a119t3b8jg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d71qjd7693a119t3b8jg-13266-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d71qjd7693a119t3b8jg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQChJekscK7nwTaA8bC72oRqtAov0ToWZuHKBSTw4g6+zV4qyTIAH5VBfEVrVB4HMW2/xFH6lbC7F23s1fW8txyWnqe65kqeiTff164M65oBqeuRa57EyuQR9qZey4eoTfEYwjlpZtmLEsBTfgdhuwVZ3tArUdY8/45G+9SzFY+/5CBJLMphgaVPhE/u3lh5njMfDTljmxri7YETf7YwJZBkrX+mnKDrka4NCqSE+6YTmlhllPwO7uOlk85GC86mKNKQVdk3i7pNJcjl7AVXhghs5DV4eSrRAIdoOhjcsxFs8D9FasnTltJtU09pZfDFuJsYZjIQJileMSDt/hJd3kPmp6qDvYydUHc9pEZ5seXEWyHXibU5kwvLAFWlvPh66MNVAhzs7OVXkiZT9ktTzZUPIwQp1RA+o3GNnmpHzQEQZULBLkttl19T3pJb4mQaw/Jr3oOvdhR5I2KscmLeFS9LtSH02VHxPTHmJN4f6hN3zOWfAOHyXnpPWM6oM4RtGqOPwtjqTDd4HQkl1zYV8nsCO6RV+RMRHTMG6vL+XzhwYioNRuJauoQppx6BiJbUk0O2W9WP2uhm1D3/ZgTHfyXHOlVDKMW+63qF9TonY5aYHWmU+r4TpP5baM27Ze38LiR/pDzhhsk9eKFP3PeGUxKwRMpbEbwPit/vUg/LE2cf8w==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202601310,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202601310},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d71qjd7693a119t3b8jg_OsDisk_1_8cc5ec47f72d4e4782d2e6dd82a81b05,storageAccountType:Premium_LRS},name:d71qjd7693a119t3b8jg_OsDisk_1_8cc5ec47f72d4e4782d2e6dd82a81b05,osType:Linux}},timeCreated:2026-03-25T09:24:12.0977482Z,vmId:c69d53c6-1b0f-4349-8426-768049d435fa}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d71qjd7693a119t3b8jg,keypair:d71qj2f693a119t3b8gg,publicip:d71qjd7693a119t3b8jg-20136-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8jg"
            },
            {
              "key": "Name",
              "value": "d71qjd7693a119t3b8jg"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "d71qjd7693a119t3b8lg",
          "cspResourceName": "d71qjd7693a119t3b8lg",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8lg",
          "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "subGroupId": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
          "createdTime": "2026-03-25 09:25:02",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-03-25 09:25:02",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8lg",
            "sys.cspResourceName": "d71qjd7693a119t3b8lg",
            "sys.id": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "subnet-01",
            "sys.uid": "d71qjd7693a119t3b8lg",
            "sys.vNetId": "vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.184.197",
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts"
          },
          "vNetId": "vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg",
          "subnetId": "subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg/subnets/d71qj0f693a119t3b8g0",
          "networkInterface": "d71qjd7693a119t3b8lg-25225-VNic",
          "securityGroupIds": [
            "sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d71qj2f693a119t3b8gg",
          "vmUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBNweexaFlGZ1XB4c/3JIDrIdYpNjFazLQ0/SVDUgiUQGf0OaNnFbkuAJMOQK1aYqFtd4P9XtWYM5vP6guOHCrxI=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:+h5v7NQYScWOmTbLSYLYmdZk0w80HyP3xsDAjijiB6g",
            "firstUsedAt": "2026-03-25T09:25:47Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-03-25T09:25:45Z",
              "completedTime": "2026-03-25T09:25:48Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d71qjd7693a119t3b8lg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d71qjd7693a119t3b8lg-25225-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d71qjd7693a119t3b8lg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQChJekscK7nwTaA8bC72oRqtAov0ToWZuHKBSTw4g6+zV4qyTIAH5VBfEVrVB4HMW2/xFH6lbC7F23s1fW8txyWnqe65kqeiTff164M65oBqeuRa57EyuQR9qZey4eoTfEYwjlpZtmLEsBTfgdhuwVZ3tArUdY8/45G+9SzFY+/5CBJLMphgaVPhE/u3lh5njMfDTljmxri7YETf7YwJZBkrX+mnKDrka4NCqSE+6YTmlhllPwO7uOlk85GC86mKNKQVdk3i7pNJcjl7AVXhghs5DV4eSrRAIdoOhjcsxFs8D9FasnTltJtU09pZfDFuJsYZjIQJileMSDt/hJd3kPmp6qDvYydUHc9pEZ5seXEWyHXibU5kwvLAFWlvPh66MNVAhzs7OVXkiZT9ktTzZUPIwQp1RA+o3GNnmpHzQEQZULBLkttl19T3pJb4mQaw/Jr3oOvdhR5I2KscmLeFS9LtSH02VHxPTHmJN4f6hN3zOWfAOHyXnpPWM6oM4RtGqOPwtjqTDd4HQkl1zYV8nsCO6RV+RMRHTMG6vL+XzhwYioNRuJauoQppx6BiJbUk0O2W9WP2uhm1D3/ZgTHfyXHOlVDKMW+63qF9TonY5aYHWmU+r4TpP5baM27Ze38LiR/pDzhhsk9eKFP3PeGUxKwRMpbEbwPit/vUg/LE2cf8w==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202601310,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202601310},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d71qjd7693a119t3b8lg_OsDisk_1_f8aee427cf8c4372bc8d02941a81f0ed,storageAccountType:Premium_LRS},name:d71qjd7693a119t3b8lg_OsDisk_1_f8aee427cf8c4372bc8d02941a81f0ed,osType:Linux}},timeCreated:2026-03-25T09:24:13.116244Z,vmId:6700becc-d787-4d20-af36-73d14dabf8d8}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d71qjd7693a119t3b8lg,keypair:d71qj2f693a119t3b8gg,publicip:d71qjd7693a119t3b8lg-4940-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8lg"
            },
            {
              "key": "Name",
              "value": "d71qjd7693a119t3b8lg"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "d71qjd7693a119t3b8kg",
          "cspResourceName": "d71qjd7693a119t3b8kg",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8kg",
          "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "subGroupId": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
          "createdTime": "2026-03-25 09:25:02",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-03-25 09:25:02",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8kg",
            "sys.cspResourceName": "d71qjd7693a119t3b8kg",
            "sys.id": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "subnet-01",
            "sys.uid": "d71qjd7693a119t3b8kg",
            "sys.vNetId": "vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.200.169.139",
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
          "specId": "azure+koreasouth+standard_b4as_v2",
          "cspSpecName": "Standard_B4as_v2",
          "spec": {
            "cspSpecName": "Standard_B4as_v2",
            "vCPU": 4,
            "memoryGiB": 15.625,
            "costPerHour": 0.173
          },
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2"
          },
          "vNetId": "vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg",
          "subnetId": "subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg/subnets/d71qj0f693a119t3b8g0",
          "networkInterface": "d71qjd7693a119t3b8kg-43320-VNic",
          "securityGroupIds": [
            "sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d71qj2f693a119t3b8gg",
          "vmUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBK1G/EPTKlJGA4doAZU6XyIIhzbzDibzPZA0lPaBhA78Y276e9O7bC4xrjt0kI5aZLNDWRdE0+wUO7snenZF0FA=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:Mwo0ByOzc/Jv/9gJ1ZP+3xUP9m3+gUFzVRrBzlzq+vk",
            "firstUsedAt": "2026-03-25T09:25:45Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-03-25T09:25:45Z",
              "completedTime": "2026-03-25T09:25:47Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d71qjd7693a119t3b8kg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d71qjd7693a119t3b8kg-43320-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d71qjd7693a119t3b8kg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQChJekscK7nwTaA8bC72oRqtAov0ToWZuHKBSTw4g6+zV4qyTIAH5VBfEVrVB4HMW2/xFH6lbC7F23s1fW8txyWnqe65kqeiTff164M65oBqeuRa57EyuQR9qZey4eoTfEYwjlpZtmLEsBTfgdhuwVZ3tArUdY8/45G+9SzFY+/5CBJLMphgaVPhE/u3lh5njMfDTljmxri7YETf7YwJZBkrX+mnKDrka4NCqSE+6YTmlhllPwO7uOlk85GC86mKNKQVdk3i7pNJcjl7AVXhghs5DV4eSrRAIdoOhjcsxFs8D9FasnTltJtU09pZfDFuJsYZjIQJileMSDt/hJd3kPmp6qDvYydUHc9pEZ5seXEWyHXibU5kwvLAFWlvPh66MNVAhzs7OVXkiZT9ktTzZUPIwQp1RA+o3GNnmpHzQEQZULBLkttl19T3pJb4mQaw/Jr3oOvdhR5I2KscmLeFS9LtSH02VHxPTHmJN4f6hN3zOWfAOHyXnpPWM6oM4RtGqOPwtjqTDd4HQkl1zYV8nsCO6RV+RMRHTMG6vL+XzhwYioNRuJauoQppx6BiJbUk0O2W9WP2uhm1D3/ZgTHfyXHOlVDKMW+63qF9TonY5aYHWmU+r4TpP5baM27Ze38LiR/pDzhhsk9eKFP3PeGUxKwRMpbEbwPit/vUg/LE2cf8w==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202601310,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202601310},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d71qjd7693a119t3b8kg_OsDisk_1_114ad1115c59441d956b29335b8e826d,storageAccountType:Premium_LRS},name:d71qjd7693a119t3b8kg_OsDisk_1_114ad1115c59441d956b29335b8e826d,osType:Linux}},timeCreated:2026-03-25T09:24:12.0977482Z,vmId:06fe8dad-90e6-4c19-bd37-cdb92def1b65}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d71qjd7693a119t3b8kg,keypair:d71qj2f693a119t3b8gg,publicip:d71qjd7693a119t3b8kg-10120-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8kg"
            },
            {
              "key": "Name",
              "value": "d71qjd7693a119t3b8kg"
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
            "vmId": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "vmIp": "20.200.169.139",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d71qjd7693a119t3b8kg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "vmIp": "20.200.184.197",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d71qjd7693a119t3b8lg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "vmIp": "20.214.57.165",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d71qjd7693a119t3b8jg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "d71qjd7693a119t3b8ig",
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
    "countRegistering": 0,
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
    "sys.uid": "d71qjd7693a119t3b8ig"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "d71qjd7693a119t3b8jg",
      "cspResourceName": "d71qjd7693a119t3b8jg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8jg",
      "name": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "subGroupId": "vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-03-25 09:25:02",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-03-25 09:25:02",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8jg",
        "sys.cspResourceName": "d71qjd7693a119t3b8jg",
        "sys.id": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "subnet-01",
        "sys.uid": "d71qjd7693a119t3b8jg",
        "sys.vNetId": "vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.214.57.165",
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
      "specId": "azure+koreasouth+standard_b2als_v2",
      "cspSpecName": "Standard_B2als_v2",
      "spec": {
        "cspSpecName": "Standard_B2als_v2",
        "vCPU": 2,
        "memoryGiB": 3.90625,
        "costPerHour": 0.0432
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2"
      },
      "vNetId": "vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg",
      "subnetId": "subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg/subnets/d71qj0f693a119t3b8g0",
      "networkInterface": "d71qjd7693a119t3b8jg-13266-VNic",
      "securityGroupIds": [
        "sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d71qj2f693a119t3b8gg",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBAaCHkSouaVM2JRH0A/WdA9IZoxc9BzMyci5iIN/EzwkVs+i9k2XGiSvzkYJLPR2flQlWyc3BZTo4yg8faHNAPg=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:q7FsLm+wFn70wol+Xs1+iamarU+xvkF8Uiei31aCuHQ",
        "firstUsedAt": "2026-03-25T09:25:47Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T09:25:45Z",
          "completedTime": "2026-03-25T09:25:50Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d71qjd7693a119t3b8jg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d71qjd7693a119t3b8jg-13266-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d71qjd7693a119t3b8jg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQChJekscK7nwTaA8bC72oRqtAov0ToWZuHKBSTw4g6+zV4qyTIAH5VBfEVrVB4HMW2/xFH6lbC7F23s1fW8txyWnqe65kqeiTff164M65oBqeuRa57EyuQR9qZey4eoTfEYwjlpZtmLEsBTfgdhuwVZ3tArUdY8/45G+9SzFY+/5CBJLMphgaVPhE/u3lh5njMfDTljmxri7YETf7YwJZBkrX+mnKDrka4NCqSE+6YTmlhllPwO7uOlk85GC86mKNKQVdk3i7pNJcjl7AVXhghs5DV4eSrRAIdoOhjcsxFs8D9FasnTltJtU09pZfDFuJsYZjIQJileMSDt/hJd3kPmp6qDvYydUHc9pEZ5seXEWyHXibU5kwvLAFWlvPh66MNVAhzs7OVXkiZT9ktTzZUPIwQp1RA+o3GNnmpHzQEQZULBLkttl19T3pJb4mQaw/Jr3oOvdhR5I2KscmLeFS9LtSH02VHxPTHmJN4f6hN3zOWfAOHyXnpPWM6oM4RtGqOPwtjqTDd4HQkl1zYV8nsCO6RV+RMRHTMG6vL+XzhwYioNRuJauoQppx6BiJbUk0O2W9WP2uhm1D3/ZgTHfyXHOlVDKMW+63qF9TonY5aYHWmU+r4TpP5baM27Ze38LiR/pDzhhsk9eKFP3PeGUxKwRMpbEbwPit/vUg/LE2cf8w==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202601310,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202601310},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d71qjd7693a119t3b8jg_OsDisk_1_8cc5ec47f72d4e4782d2e6dd82a81b05,storageAccountType:Premium_LRS},name:d71qjd7693a119t3b8jg_OsDisk_1_8cc5ec47f72d4e4782d2e6dd82a81b05,osType:Linux}},timeCreated:2026-03-25T09:24:12.0977482Z,vmId:c69d53c6-1b0f-4349-8426-768049d435fa}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d71qjd7693a119t3b8jg,keypair:d71qj2f693a119t3b8gg,publicip:d71qjd7693a119t3b8jg-20136-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8jg"
        },
        {
          "key": "Name",
          "value": "d71qjd7693a119t3b8jg"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "d71qjd7693a119t3b8lg",
      "cspResourceName": "d71qjd7693a119t3b8lg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8lg",
      "name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "subGroupId": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-03-25 09:25:02",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-03-25 09:25:02",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8lg",
        "sys.cspResourceName": "d71qjd7693a119t3b8lg",
        "sys.id": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "subnet-01",
        "sys.uid": "d71qjd7693a119t3b8lg",
        "sys.vNetId": "vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.184.197",
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202601310",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts"
      },
      "vNetId": "vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg",
      "subnetId": "subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg/subnets/d71qj0f693a119t3b8g0",
      "networkInterface": "d71qjd7693a119t3b8lg-25225-VNic",
      "securityGroupIds": [
        "sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d71qj2f693a119t3b8gg",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBNweexaFlGZ1XB4c/3JIDrIdYpNjFazLQ0/SVDUgiUQGf0OaNnFbkuAJMOQK1aYqFtd4P9XtWYM5vP6guOHCrxI=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:+h5v7NQYScWOmTbLSYLYmdZk0w80HyP3xsDAjijiB6g",
        "firstUsedAt": "2026-03-25T09:25:47Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T09:25:45Z",
          "completedTime": "2026-03-25T09:25:48Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d71qjd7693a119t3b8lg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d71qjd7693a119t3b8lg-25225-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d71qjd7693a119t3b8lg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQChJekscK7nwTaA8bC72oRqtAov0ToWZuHKBSTw4g6+zV4qyTIAH5VBfEVrVB4HMW2/xFH6lbC7F23s1fW8txyWnqe65kqeiTff164M65oBqeuRa57EyuQR9qZey4eoTfEYwjlpZtmLEsBTfgdhuwVZ3tArUdY8/45G+9SzFY+/5CBJLMphgaVPhE/u3lh5njMfDTljmxri7YETf7YwJZBkrX+mnKDrka4NCqSE+6YTmlhllPwO7uOlk85GC86mKNKQVdk3i7pNJcjl7AVXhghs5DV4eSrRAIdoOhjcsxFs8D9FasnTltJtU09pZfDFuJsYZjIQJileMSDt/hJd3kPmp6qDvYydUHc9pEZ5seXEWyHXibU5kwvLAFWlvPh66MNVAhzs7OVXkiZT9ktTzZUPIwQp1RA+o3GNnmpHzQEQZULBLkttl19T3pJb4mQaw/Jr3oOvdhR5I2KscmLeFS9LtSH02VHxPTHmJN4f6hN3zOWfAOHyXnpPWM6oM4RtGqOPwtjqTDd4HQkl1zYV8nsCO6RV+RMRHTMG6vL+XzhwYioNRuJauoQppx6BiJbUk0O2W9WP2uhm1D3/ZgTHfyXHOlVDKMW+63qF9TonY5aYHWmU+r4TpP5baM27Ze38LiR/pDzhhsk9eKFP3PeGUxKwRMpbEbwPit/vUg/LE2cf8w==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202601310,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202601310},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d71qjd7693a119t3b8lg_OsDisk_1_f8aee427cf8c4372bc8d02941a81f0ed,storageAccountType:Premium_LRS},name:d71qjd7693a119t3b8lg_OsDisk_1_f8aee427cf8c4372bc8d02941a81f0ed,osType:Linux}},timeCreated:2026-03-25T09:24:13.116244Z,vmId:6700becc-d787-4d20-af36-73d14dabf8d8}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d71qjd7693a119t3b8lg,keypair:d71qj2f693a119t3b8gg,publicip:d71qjd7693a119t3b8lg-4940-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8lg"
        },
        {
          "key": "Name",
          "value": "d71qjd7693a119t3b8lg"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "d71qjd7693a119t3b8kg",
      "cspResourceName": "d71qjd7693a119t3b8kg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8kg",
      "name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "subGroupId": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-03-25 09:25:02",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-03-25 09:25:02",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8kg",
        "sys.cspResourceName": "d71qjd7693a119t3b8kg",
        "sys.id": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "subnet-01",
        "sys.uid": "d71qjd7693a119t3b8kg",
        "sys.vNetId": "vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.200.169.139",
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
      "specId": "azure+koreasouth+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "spec": {
        "cspSpecName": "Standard_B4as_v2",
        "vCPU": 4,
        "memoryGiB": 15.625,
        "costPerHour": 0.173
      },
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202601310",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2"
      },
      "vNetId": "vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg",
      "subnetId": "subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/d71qj0f693a119t3b8fg/subnets/d71qj0f693a119t3b8g0",
      "networkInterface": "d71qjd7693a119t3b8kg-43320-VNic",
      "securityGroupIds": [
        "sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/d71qj2f693a119t3b8gg",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBK1G/EPTKlJGA4doAZU6XyIIhzbzDibzPZA0lPaBhA78Y276e9O7bC4xrjt0kI5aZLNDWRdE0+wUO7snenZF0FA=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:Mwo0ByOzc/Jv/9gJ1ZP+3xUP9m3+gUFzVRrBzlzq+vk",
        "firstUsedAt": "2026-03-25T09:25:45Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T09:25:45Z",
          "completedTime": "2026-03-25T09:25:47Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d71qjd7693a119t3b8kg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/d71qjd7693a119t3b8kg-43320-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d71qjd7693a119t3b8kg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQChJekscK7nwTaA8bC72oRqtAov0ToWZuHKBSTw4g6+zV4qyTIAH5VBfEVrVB4HMW2/xFH6lbC7F23s1fW8txyWnqe65kqeiTff164M65oBqeuRa57EyuQR9qZey4eoTfEYwjlpZtmLEsBTfgdhuwVZ3tArUdY8/45G+9SzFY+/5CBJLMphgaVPhE/u3lh5njMfDTljmxri7YETf7YwJZBkrX+mnKDrka4NCqSE+6YTmlhllPwO7uOlk85GC86mKNKQVdk3i7pNJcjl7AVXhghs5DV4eSrRAIdoOhjcsxFs8D9FasnTltJtU09pZfDFuJsYZjIQJileMSDt/hJd3kPmp6qDvYydUHc9pEZ5seXEWyHXibU5kwvLAFWlvPh66MNVAhzs7OVXkiZT9ktTzZUPIwQp1RA+o3GNnmpHzQEQZULBLkttl19T3pJb4mQaw/Jr3oOvdhR5I2KscmLeFS9LtSH02VHxPTHmJN4f6hN3zOWfAOHyXnpPWM6oM4RtGqOPwtjqTDd4HQkl1zYV8nsCO6RV+RMRHTMG6vL+XzhwYioNRuJauoQppx6BiJbUk0O2W9WP2uhm1D3/ZgTHfyXHOlVDKMW+63qF9TonY5aYHWmU+r4TpP5baM27Ze38LiR/pDzhhsk9eKFP3PeGUxKwRMpbEbwPit/vUg/LE2cf8w==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202601310,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202601310},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:50,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/d71qjd7693a119t3b8kg_OsDisk_1_114ad1115c59441d956b29335b8e826d,storageAccountType:Premium_LRS},name:d71qjd7693a119t3b8kg_OsDisk_1_114ad1115c59441d956b29335b8e826d,osType:Linux}},timeCreated:2026-03-25T09:24:12.0977482Z,vmId:06fe8dad-90e6-4c19-bd37-cdb92def1b65}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d71qjd7693a119t3b8kg,keypair:d71qj2f693a119t3b8gg,publicip:d71qjd7693a119t3b8kg-10120-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/d71qjd7693a119t3b8kg"
        },
        {
          "key": "Name",
          "value": "d71qjd7693a119t3b8kg"
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
        "vmId": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "vmIp": "20.200.169.139",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d71qjd7693a119t3b8kg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "vmIp": "20.200.184.197",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d71qjd7693a119t3b8lg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vmIp": "20.214.57.165",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d71qjd7693a119t3b8jg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "output": "Linux d71qjd7693a119t3b8jg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "20.214.57.165",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "vm-ec268ed7-821e-9d73-e79f-961262161624",
      "testOrder": 1,
      "userName": "cb-user",
      "vmId": "vm-ec268ed7-821e-9d73-e79f-961262161624-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux d71qjd7693a119t3b8lg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "20.200.184.197",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "testOrder": 2,
      "userName": "cb-user",
      "vmId": "vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux d71qjd7693a119t3b8kg 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "20.200.169.139",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "testOrder": 3,
      "userName": "cb-user",
      "vmId": "vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1"
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
  "message": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: mmci01)",
  "success": true
}
```

