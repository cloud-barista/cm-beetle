# CM-Beetle test results for AZURE

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: imdl/v0.1.5+ (c574040)
- cm-model: unknown
- CB-Tumblebug: vunknown
- CB-Spider: vunknown
- CB-MapUI: vunknown
- Target CSP: AZURE
- Target Region: koreasouth
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: May 18, 2026
- Test Time: 17:01:38 KST
- Test Execution: 2026-05-18 17:01:38 KST

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
| 1 | `POST /beetle/recommendation/infra` | ✅ **PASS** | 7.581s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 1m39.898s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ **PASS** | 42ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ **PASS** | 4ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 28ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.269s | Pass |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 5.399s | Pass |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ **PASS** | 1m48.41s | Pass |

**Overall Result**: 9/9 tests passed ✅

**Total Duration**: 4m56.692755828s

*Test executed on May 18, 2026 at 17:01:38 KST (2026-05-18 17:01:38 KST) using CM-Beetle automated test CLI*

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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605150",
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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605150",
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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605150",
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
        },
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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605150",
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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605150",
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
            "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605150",
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
        },
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
  "uid": "tbd85cf3epr9ha6omvqjbg",
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
    "sys.uid": "tbd85cf3epr9ha6omvqjbg"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbd85cf3epr9ha6omvqjcg",
      "cspResourceName": "tbd85cf3epr9ha6omvqjcg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg",
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
      "createdTime": "2026-05-18 08:03:14",
      "label": {
        "createdBy": "tbd85cf3epr9ha6omvqjcg",
        "keypair": "tbd85cf06pr9ha6omvqj3g",
        "publicip": "tbd85cf3epr9ha6omvqjcg-42454-PublicIP",
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-18 08:03:14",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg",
        "sys.cspResourceName": "tbd85cf3epr9ha6omvqjcg",
        "sys.id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd85cf3epr9ha6omvqjcg",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.214.25.229",
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605150",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0",
      "networkInterface": "tbd85cf3epr9ha6omvqjcg-39198-VNic",
      "securityGroupIds": [
        "my02-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBL0VTwP0mj+c6VjV8msgi3Q9eibNRewPMT54DUhTIpAg1Sht7ha22G1ufqkYptq1cFE80rR711dScHPLLhmJ+q8=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:L+I1/Qhho+5MTJJQUC536mZqODIvIydVP6vJjVbdai0",
        "firstUsedAt": "2026-05-18T08:03:23Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-18T08:03:23Z",
          "completedTime": "2026-05-18T08:03:26Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd85cf3epr9ha6omvqjcg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd85cf3epr9ha6omvqjcg-39198-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd85cf3epr9ha6omvqjcg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDpWcbrhtQvDFRYHly+CcVE71id44biB3QrFVXkBa7y8vHQa3Kj2v668qxkIfYEMUzInhuyqssv3AM5dKNvdl2cKz2BIBRNzpNxjGKEoHGNikkeKpFuEZxMhR5iIv1E36MU5mv6iAXDScum56Tp/97kKRSBMeZZo0Dq6b+SYUvFRUkBCBf84aIlQFD4KqiZ+Fe8j7V+vfOirroTVwZEaIUOSzfBvkZvKbpk8b86GzrxhnSgJyedkl7VoG+xpVg9iC7FjmIzyg9FK/8tk/FIje3BBmx2p83xhMyzquSopaJ4irYgDHRk5NWucds+nlJht5Zv9oxNsVXvo0eXu5Gg6EIi1S1zpPi6pEu1B9wx+0K5I93ZCMTrS2kbjnsPC/hubwXwGSAOdrm2uZAo8XYNH2LY8AyKHBAmBGtUwol7X7Dis4PLFkbApL2PLNovAHtLwpb0/+Zbqif4jMNV1ijaNsmN8n65MxuZ09SubadWql6JMh4glfS/8Rt7b0TJ/B9HtFSF1e+hGkt6nfxFtjT9ucyJvzQ+s+H6RZFUQYltLPFmqX0LSef1/xEVooKDofYtT5yFy/cE5/kFMbkKg2uWe1MkOlzwsBOXlmWDWX1U5pyZ9vg9wbtLM38qh+2Q3yBaroftUBxXy/nt9lDdPj/KnZV0NxGw4gQ82BBPYD80pfCZsw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202605150,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202605150},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd85cf3epr9ha6omvqjcg_OsDisk_1_15c2d99efd01471b8724e152bfe481ca,storageAccountType:Premium_LRS},name:tbd85cf3epr9ha6omvqjcg_OsDisk_1_15c2d99efd01471b8724e152bfe481ca,osType:Linux}},timeCreated:2026-05-18T08:02:26.9394222Z,vmId:66b1ceb1-9da8-497a-ab77-8e1d40b4f26c}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd85cf3epr9ha6omvqjcg,keypair:tbd85cf06pr9ha6omvqj3g,publicip:tbd85cf3epr9ha6omvqjcg-42454-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg"
        },
        {
          "key": "Name",
          "value": "tbd85cf3epr9ha6omvqjcg"
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
      "uid": "tbd85cf3epr9ha6omvqjeg",
      "cspResourceName": "tbd85cf3epr9ha6omvqjeg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg",
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
      "createdTime": "2026-05-18 08:03:17",
      "label": {
        "createdBy": "tbd85cf3epr9ha6omvqjeg",
        "keypair": "tbd85cf06pr9ha6omvqj3g",
        "publicip": "tbd85cf3epr9ha6omvqjeg-83151-PublicIP",
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-18 08:03:17",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg",
        "sys.cspResourceName": "tbd85cf3epr9ha6omvqjeg",
        "sys.id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd85cf3epr9ha6omvqjeg",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.214.26.6",
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605150",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0",
      "networkInterface": "tbd85cf3epr9ha6omvqjeg-47871-VNic",
      "securityGroupIds": [
        "my02-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEjxGVDmlxPviCew7MPsxQCg/SbRH0g1LPINkzdzqnLKnu1r1BVcN2sOxPrrm4d3ektW4i1wjM5H3IyYVNdYIXM=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:kXaggWkouCHVkByZEeZ7WbredHOydMu4PhHCk95sdGA",
        "firstUsedAt": "2026-05-18T08:03:25Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-18T08:03:23Z",
          "completedTime": "2026-05-18T08:03:27Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd85cf3epr9ha6omvqjeg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd85cf3epr9ha6omvqjeg-47871-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd85cf3epr9ha6omvqjeg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDpWcbrhtQvDFRYHly+CcVE71id44biB3QrFVXkBa7y8vHQa3Kj2v668qxkIfYEMUzInhuyqssv3AM5dKNvdl2cKz2BIBRNzpNxjGKEoHGNikkeKpFuEZxMhR5iIv1E36MU5mv6iAXDScum56Tp/97kKRSBMeZZo0Dq6b+SYUvFRUkBCBf84aIlQFD4KqiZ+Fe8j7V+vfOirroTVwZEaIUOSzfBvkZvKbpk8b86GzrxhnSgJyedkl7VoG+xpVg9iC7FjmIzyg9FK/8tk/FIje3BBmx2p83xhMyzquSopaJ4irYgDHRk5NWucds+nlJht5Zv9oxNsVXvo0eXu5Gg6EIi1S1zpPi6pEu1B9wx+0K5I93ZCMTrS2kbjnsPC/hubwXwGSAOdrm2uZAo8XYNH2LY8AyKHBAmBGtUwol7X7Dis4PLFkbApL2PLNovAHtLwpb0/+Zbqif4jMNV1ijaNsmN8n65MxuZ09SubadWql6JMh4glfS/8Rt7b0TJ/B9HtFSF1e+hGkt6nfxFtjT9ucyJvzQ+s+H6RZFUQYltLPFmqX0LSef1/xEVooKDofYtT5yFy/cE5/kFMbkKg2uWe1MkOlzwsBOXlmWDWX1U5pyZ9vg9wbtLM38qh+2Q3yBaroftUBxXy/nt9lDdPj/KnZV0NxGw4gQ82BBPYD80pfCZsw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202605150,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202605150},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd85cf3epr9ha6omvqjeg_OsDisk_1_8aaddc61d3c5476e8e7f50a0be8384c8,storageAccountType:Premium_LRS},name:tbd85cf3epr9ha6omvqjeg_OsDisk_1_8aaddc61d3c5476e8e7f50a0be8384c8,osType:Linux}},timeCreated:2026-05-18T08:02:28.9784295Z,vmId:85838c4c-38c3-4641-9333-9b5b9ab4769a}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd85cf3epr9ha6omvqjeg,keypair:tbd85cf06pr9ha6omvqj3g,publicip:tbd85cf3epr9ha6omvqjeg-83151-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg"
        },
        {
          "key": "Name",
          "value": "tbd85cf3epr9ha6omvqjeg"
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
      "uid": "tbd85cf3epr9ha6omvqjdg",
      "cspResourceName": "tbd85cf3epr9ha6omvqjdg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg",
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
      "createdTime": "2026-05-18 08:03:15",
      "label": {
        "createdBy": "tbd85cf3epr9ha6omvqjdg",
        "keypair": "tbd85cf06pr9ha6omvqj3g",
        "publicip": "tbd85cf3epr9ha6omvqjdg-86268-PublicIP",
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-18 08:03:15",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg",
        "sys.cspResourceName": "tbd85cf3epr9ha6omvqjdg",
        "sys.id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd85cf3epr9ha6omvqjdg",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "52.231.195.68",
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605150",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0",
      "networkInterface": "tbd85cf3epr9ha6omvqjdg-3065-VNic",
      "securityGroupIds": [
        "my02-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBPTNwiwA06iMy3QoeSkT4pJXpOsH22Jn6+LnZ+ak239UDz03i5xCcIOmCYDASX7JMLHv1NfehWWjviZjXQOCm0A=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:zSgb3/NtuvOcCa3S2XoWo/Kj3Tpb4B15i3wLbr896JA",
        "firstUsedAt": "2026-05-18T08:03:25Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-18T08:03:23Z",
          "completedTime": "2026-05-18T08:03:28Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd85cf3epr9ha6omvqjdg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd85cf3epr9ha6omvqjdg-3065-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd85cf3epr9ha6omvqjdg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDpWcbrhtQvDFRYHly+CcVE71id44biB3QrFVXkBa7y8vHQa3Kj2v668qxkIfYEMUzInhuyqssv3AM5dKNvdl2cKz2BIBRNzpNxjGKEoHGNikkeKpFuEZxMhR5iIv1E36MU5mv6iAXDScum56Tp/97kKRSBMeZZo0Dq6b+SYUvFRUkBCBf84aIlQFD4KqiZ+Fe8j7V+vfOirroTVwZEaIUOSzfBvkZvKbpk8b86GzrxhnSgJyedkl7VoG+xpVg9iC7FjmIzyg9FK/8tk/FIje3BBmx2p83xhMyzquSopaJ4irYgDHRk5NWucds+nlJht5Zv9oxNsVXvo0eXu5Gg6EIi1S1zpPi6pEu1B9wx+0K5I93ZCMTrS2kbjnsPC/hubwXwGSAOdrm2uZAo8XYNH2LY8AyKHBAmBGtUwol7X7Dis4PLFkbApL2PLNovAHtLwpb0/+Zbqif4jMNV1ijaNsmN8n65MxuZ09SubadWql6JMh4glfS/8Rt7b0TJ/B9HtFSF1e+hGkt6nfxFtjT9ucyJvzQ+s+H6RZFUQYltLPFmqX0LSef1/xEVooKDofYtT5yFy/cE5/kFMbkKg2uWe1MkOlzwsBOXlmWDWX1U5pyZ9vg9wbtLM38qh+2Q3yBaroftUBxXy/nt9lDdPj/KnZV0NxGw4gQ82BBPYD80pfCZsw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202605150,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202605150},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd85cf3epr9ha6omvqjdg_OsDisk_1_c531642e3a2143bc9b21c3a62be80a3d,storageAccountType:Premium_LRS},name:tbd85cf3epr9ha6omvqjdg_OsDisk_1_c531642e3a2143bc9b21c3a62be80a3d,osType:Linux}},timeCreated:2026-05-18T08:02:27.2996702Z,vmId:9f82564d-bc46-434c-955e-4e8a0c6ad147}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd85cf3epr9ha6omvqjdg,keypair:tbd85cf06pr9ha6omvqj3g,publicip:tbd85cf3epr9ha6omvqjdg-86268-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg"
        },
        {
          "key": "Name",
          "value": "tbd85cf3epr9ha6omvqjdg"
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
        "nodeIp": "20.214.25.229",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd85cf3epr9ha6omvqjcg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my02-infra101",
        "nodeId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "20.214.26.6",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd85cf3epr9ha6omvqjeg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my02-infra101",
        "nodeId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "nodeIp": "52.231.195.68",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd85cf3epr9ha6omvqjdg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "tbd85cevepr9ha6omvqit0",
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
        "sys.uid": "tbd85cevepr9ha6omvqit0"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbd85cevepr9ha6omvqiu0",
          "cspResourceName": "tbd85cevepr9ha6omvqiu0",
          "cspResourceId": "i-05cc436e2832ffd93",
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
          "createdTime": "2026-05-18 08:02:29",
          "label": {
            "Name": "tbd85cevepr9ha6omvqiu0",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-05-18 08:02:29",
            "sys.cspResourceId": "i-05cc436e2832ffd93",
            "sys.cspResourceName": "tbd85cevepr9ha6omvqiu0",
            "sys.id": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tbd85cevepr9ha6omvqiu0",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "43.203.239.151",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.41",
          "privateDNS": "ip-10-0-1-41.ap-northeast-2.compute.internal",
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
          "cspVNetId": "vpc-07e218262146e7003",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-0695096ac3c17916f",
          "networkInterface": "eni-attach-04753091ce5855ac2",
          "securityGroupIds": [
            "my01-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbd85cetmpr9ha6omvqiog",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBL9WpZDZ14fXmSWHWCNYrwkApOqX/SoVpoE9eR/XjTGoct1jqM8ZiZGyiZe2PnPcV5WkwUfFRJGmIJ72lSHwRyU=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:L/BdBg26WaqVC1vQRLLOePkyYT/RR3AJ6iJqZ8JNevs",
            "firstUsedAt": "2026-05-18T08:02:39Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-18T08:02:37Z",
              "completedTime": "2026-05-18T08:02:40Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-41 6.8.0-1051-aws #54~22.04.1-Ubuntu SMP Wed Mar 25 15:41:00 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-05-18T08:02:08Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0c33945aafaedb72d}}"
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
              "value": "368C7E62-3F66-4741-941F-5D64BEAB1E56"
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
              "value": "i-05cc436e2832ffd93"
            },
            {
              "key": "InstanceType",
              "value": "t3a.small"
            },
            {
              "key": "KeyName",
              "value": "tbd85cetmpr9ha6omvqiog"
            },
            {
              "key": "LaunchTime",
              "value": "2026-05-18T08:02:08Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.239.151},Attachment:{AttachTime:2026-05-18T08:02:08Z,AttachmentId:eni-attach-04753091ce5855ac2,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-094d187b192dfdf77,GroupName:tbd85cetupr9ha6omvqip0}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:20:f5:c8:48:cb,NetworkInterfaceId:eni-095815b03a754c42c,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.41,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.239.151},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.41}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0695096ac3c17916f,VpcId:vpc-07e218262146e7003}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-41.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.41"
            },
            {
              "key": "PublicIpAddress",
              "value": "43.203.239.151"
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
              "value": "{GroupId:sg-094d187b192dfdf77,GroupName:tbd85cetupr9ha6omvqip0}"
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
              "value": "subnet-0695096ac3c17916f"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbd85cevepr9ha6omvqiu0}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-07e218262146e7003"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "tbd85cevepr9ha6omvqj00",
          "cspResourceName": "tbd85cevepr9ha6omvqj00",
          "cspResourceId": "i-055c3be58ba3b717a",
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
          "createdTime": "2026-05-18 08:02:31",
          "label": {
            "Name": "tbd85cevepr9ha6omvqj00",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-05-18 08:02:31",
            "sys.cspResourceId": "i-055c3be58ba3b717a",
            "sys.cspResourceName": "tbd85cevepr9ha6omvqj00",
            "sys.id": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tbd85cevepr9ha6omvqj00",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "13.124.168.73",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.46",
          "privateDNS": "ip-10-0-1-46.ap-northeast-2.compute.internal",
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
          "cspVNetId": "vpc-07e218262146e7003",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-0695096ac3c17916f",
          "networkInterface": "eni-attach-0512107afc0efe575",
          "securityGroupIds": [
            "my01-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbd85cetmpr9ha6omvqiog",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBF+I10zmH+ojLGsV4az5CwdBc9/041SKt5GPUc/YqKVNPxdyQITwRs/cLqkZHw97eOVFlKJAHO8SmQpMotHG9ek=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:JIQ/foOGdNVKJH+QeoKeGnpkqZ050cJXHJqCN63pMLY",
            "firstUsedAt": "2026-05-18T08:02:39Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-18T08:02:37Z",
              "completedTime": "2026-05-18T08:02:40Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-46 6.8.0-1051-aws #54~22.04.1-Ubuntu SMP Wed Mar 25 15:41:00 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-05-18T08:02:09Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0574dd1a09d5c632d}}"
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
              "value": "1FF0C9D1-BAAF-4015-A570-54CE3F1704A3"
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
              "value": "i-055c3be58ba3b717a"
            },
            {
              "key": "InstanceType",
              "value": "t3a.large"
            },
            {
              "key": "KeyName",
              "value": "tbd85cetmpr9ha6omvqiog"
            },
            {
              "key": "LaunchTime",
              "value": "2026-05-18T08:02:08Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.124.168.73},Attachment:{AttachTime:2026-05-18T08:02:08Z,AttachmentId:eni-attach-0512107afc0efe575,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b2020b2fa704162e,GroupName:tbd85ceuupr9ha6omvqis0}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:f4:96:e2:c0:8b,NetworkInterfaceId:eni-0c3a984205d7629b3,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.46,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:13.124.168.73},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.46}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0695096ac3c17916f,VpcId:vpc-07e218262146e7003}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-46.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.46"
            },
            {
              "key": "PublicIpAddress",
              "value": "13.124.168.73"
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
              "value": "{GroupId:sg-0b2020b2fa704162e,GroupName:tbd85ceuupr9ha6omvqis0}"
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
              "value": "subnet-0695096ac3c17916f"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbd85cevepr9ha6omvqj00}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-07e218262146e7003"
            }
          ]
        },
        {
          "resourceType": "node",
          "id": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "tbd85cevepr9ha6omvqiv0",
          "cspResourceName": "tbd85cevepr9ha6omvqiv0",
          "cspResourceId": "i-0fa2d84cd0c6e08bb",
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
          "createdTime": "2026-05-18 08:02:28",
          "label": {
            "Name": "tbd85cevepr9ha6omvqiv0",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-05-18 08:02:28",
            "sys.cspResourceId": "i-0fa2d84cd0c6e08bb",
            "sys.cspResourceName": "tbd85cevepr9ha6omvqiv0",
            "sys.id": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my01-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my01-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my01-subnet-01",
            "sys.uid": "tbd85cevepr9ha6omvqiv0",
            "sys.vNetId": "my01-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "52.78.166.174",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.204",
          "privateDNS": "ip-10-0-1-204.ap-northeast-2.compute.internal",
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
          "cspVNetId": "vpc-07e218262146e7003",
          "subnetId": "my01-subnet-01",
          "cspSubnetId": "subnet-0695096ac3c17916f",
          "networkInterface": "eni-attach-0dacfd07e29c82162",
          "securityGroupIds": [
            "my01-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my01-sshkey-01",
          "cspSshKeyId": "tbd85cetmpr9ha6omvqiog",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJyKArhSkC+xcVNnvzj9CwMwKnXva6HvRVhVO1h0wfBpcTDP/ViBWoe96OAkh9DZCT1QWtAcmhnEDgJ5jAGXLnE=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:PhaxYolMCyF8JSnMZ5JzXZ8GWc+VS+x3YRIZQ/caJN4",
            "firstUsedAt": "2026-05-18T08:02:37Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-18T08:02:37Z",
              "completedTime": "2026-05-18T08:02:40Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-204 6.8.0-1051-aws #54~22.04.1-Ubuntu SMP Wed Mar 25 15:41:00 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-05-18T08:02:07Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-072521d142190aba5}}"
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
              "value": "FDE8AD79-2C70-4CD5-B35E-3DBC181BEDD7"
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
              "value": "i-0fa2d84cd0c6e08bb"
            },
            {
              "key": "InstanceType",
              "value": "t3a.xlarge"
            },
            {
              "key": "KeyName",
              "value": "tbd85cetmpr9ha6omvqiog"
            },
            {
              "key": "LaunchTime",
              "value": "2026-05-18T08:02:07Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:52.78.166.174},Attachment:{AttachTime:2026-05-18T08:02:07Z,AttachmentId:eni-attach-0dacfd07e29c82162,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-014a3bc57edc785c7,GroupName:tbd85ceuepr9ha6omvqirg}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:06:f2:cb:48:2b,NetworkInterfaceId:eni-071a58e90add5f477,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.204,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:52.78.166.174},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.204}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0695096ac3c17916f,VpcId:vpc-07e218262146e7003}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-204.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.204"
            },
            {
              "key": "PublicIpAddress",
              "value": "52.78.166.174"
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
              "value": "{GroupId:sg-014a3bc57edc785c7,GroupName:tbd85ceuepr9ha6omvqirg}"
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
              "value": "subnet-0695096ac3c17916f"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:tbd85cevepr9ha6omvqiv0}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-07e218262146e7003"
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
            "nodeIp": "52.78.166.174",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-204 6.8.0-1051-aws #54~22.04.1-Ubuntu SMP Wed Mar 25 15:41:00 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "43.203.239.151",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-41 6.8.0-1051-aws #54~22.04.1-Ubuntu SMP Wed Mar 25 15:41:00 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my01-infra101",
            "nodeId": "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "13.124.168.73",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-46 6.8.0-1051-aws #54~22.04.1-Ubuntu SMP Wed Mar 25 15:41:00 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "tbd85cf3epr9ha6omvqjbg",
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
        "sys.uid": "tbd85cf3epr9ha6omvqjbg"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbd85cf3epr9ha6omvqjcg",
          "cspResourceName": "tbd85cf3epr9ha6omvqjcg",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg",
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
          "createdTime": "2026-05-18 08:03:14",
          "label": {
            "createdBy": "tbd85cf3epr9ha6omvqjcg",
            "keypair": "tbd85cf06pr9ha6omvqj3g",
            "publicip": "tbd85cf3epr9ha6omvqjcg-42454-PublicIP",
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-05-18 08:03:14",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg",
            "sys.cspResourceName": "tbd85cf3epr9ha6omvqjcg",
            "sys.id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbd85cf3epr9ha6omvqjcg",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.25.229",
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605150",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0",
          "networkInterface": "tbd85cf3epr9ha6omvqjcg-39198-VNic",
          "securityGroupIds": [
            "my02-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBL0VTwP0mj+c6VjV8msgi3Q9eibNRewPMT54DUhTIpAg1Sht7ha22G1ufqkYptq1cFE80rR711dScHPLLhmJ+q8=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:L+I1/Qhho+5MTJJQUC536mZqODIvIydVP6vJjVbdai0",
            "firstUsedAt": "2026-05-18T08:03:23Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-18T08:03:23Z",
              "completedTime": "2026-05-18T08:03:26Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbd85cf3epr9ha6omvqjcg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd85cf3epr9ha6omvqjcg-39198-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd85cf3epr9ha6omvqjcg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDpWcbrhtQvDFRYHly+CcVE71id44biB3QrFVXkBa7y8vHQa3Kj2v668qxkIfYEMUzInhuyqssv3AM5dKNvdl2cKz2BIBRNzpNxjGKEoHGNikkeKpFuEZxMhR5iIv1E36MU5mv6iAXDScum56Tp/97kKRSBMeZZo0Dq6b+SYUvFRUkBCBf84aIlQFD4KqiZ+Fe8j7V+vfOirroTVwZEaIUOSzfBvkZvKbpk8b86GzrxhnSgJyedkl7VoG+xpVg9iC7FjmIzyg9FK/8tk/FIje3BBmx2p83xhMyzquSopaJ4irYgDHRk5NWucds+nlJht5Zv9oxNsVXvo0eXu5Gg6EIi1S1zpPi6pEu1B9wx+0K5I93ZCMTrS2kbjnsPC/hubwXwGSAOdrm2uZAo8XYNH2LY8AyKHBAmBGtUwol7X7Dis4PLFkbApL2PLNovAHtLwpb0/+Zbqif4jMNV1ijaNsmN8n65MxuZ09SubadWql6JMh4glfS/8Rt7b0TJ/B9HtFSF1e+hGkt6nfxFtjT9ucyJvzQ+s+H6RZFUQYltLPFmqX0LSef1/xEVooKDofYtT5yFy/cE5/kFMbkKg2uWe1MkOlzwsBOXlmWDWX1U5pyZ9vg9wbtLM38qh+2Q3yBaroftUBxXy/nt9lDdPj/KnZV0NxGw4gQ82BBPYD80pfCZsw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202605150,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202605150},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd85cf3epr9ha6omvqjcg_OsDisk_1_15c2d99efd01471b8724e152bfe481ca,storageAccountType:Premium_LRS},name:tbd85cf3epr9ha6omvqjcg_OsDisk_1_15c2d99efd01471b8724e152bfe481ca,osType:Linux}},timeCreated:2026-05-18T08:02:26.9394222Z,vmId:66b1ceb1-9da8-497a-ab77-8e1d40b4f26c}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbd85cf3epr9ha6omvqjcg,keypair:tbd85cf06pr9ha6omvqj3g,publicip:tbd85cf3epr9ha6omvqjcg-42454-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg"
            },
            {
              "key": "Name",
              "value": "tbd85cf3epr9ha6omvqjcg"
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
          "uid": "tbd85cf3epr9ha6omvqjeg",
          "cspResourceName": "tbd85cf3epr9ha6omvqjeg",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg",
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
          "createdTime": "2026-05-18 08:03:17",
          "label": {
            "createdBy": "tbd85cf3epr9ha6omvqjeg",
            "keypair": "tbd85cf06pr9ha6omvqj3g",
            "publicip": "tbd85cf3epr9ha6omvqjeg-83151-PublicIP",
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-05-18 08:03:17",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg",
            "sys.cspResourceName": "tbd85cf3epr9ha6omvqjeg",
            "sys.id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbd85cf3epr9ha6omvqjeg",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "20.214.26.6",
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605150",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0",
          "networkInterface": "tbd85cf3epr9ha6omvqjeg-47871-VNic",
          "securityGroupIds": [
            "my02-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEjxGVDmlxPviCew7MPsxQCg/SbRH0g1LPINkzdzqnLKnu1r1BVcN2sOxPrrm4d3ektW4i1wjM5H3IyYVNdYIXM=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:kXaggWkouCHVkByZEeZ7WbredHOydMu4PhHCk95sdGA",
            "firstUsedAt": "2026-05-18T08:03:25Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-18T08:03:23Z",
              "completedTime": "2026-05-18T08:03:27Z",
              "elapsedTime": 4,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbd85cf3epr9ha6omvqjeg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd85cf3epr9ha6omvqjeg-47871-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd85cf3epr9ha6omvqjeg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDpWcbrhtQvDFRYHly+CcVE71id44biB3QrFVXkBa7y8vHQa3Kj2v668qxkIfYEMUzInhuyqssv3AM5dKNvdl2cKz2BIBRNzpNxjGKEoHGNikkeKpFuEZxMhR5iIv1E36MU5mv6iAXDScum56Tp/97kKRSBMeZZo0Dq6b+SYUvFRUkBCBf84aIlQFD4KqiZ+Fe8j7V+vfOirroTVwZEaIUOSzfBvkZvKbpk8b86GzrxhnSgJyedkl7VoG+xpVg9iC7FjmIzyg9FK/8tk/FIje3BBmx2p83xhMyzquSopaJ4irYgDHRk5NWucds+nlJht5Zv9oxNsVXvo0eXu5Gg6EIi1S1zpPi6pEu1B9wx+0K5I93ZCMTrS2kbjnsPC/hubwXwGSAOdrm2uZAo8XYNH2LY8AyKHBAmBGtUwol7X7Dis4PLFkbApL2PLNovAHtLwpb0/+Zbqif4jMNV1ijaNsmN8n65MxuZ09SubadWql6JMh4glfS/8Rt7b0TJ/B9HtFSF1e+hGkt6nfxFtjT9ucyJvzQ+s+H6RZFUQYltLPFmqX0LSef1/xEVooKDofYtT5yFy/cE5/kFMbkKg2uWe1MkOlzwsBOXlmWDWX1U5pyZ9vg9wbtLM38qh+2Q3yBaroftUBxXy/nt9lDdPj/KnZV0NxGw4gQ82BBPYD80pfCZsw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202605150,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202605150},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd85cf3epr9ha6omvqjeg_OsDisk_1_8aaddc61d3c5476e8e7f50a0be8384c8,storageAccountType:Premium_LRS},name:tbd85cf3epr9ha6omvqjeg_OsDisk_1_8aaddc61d3c5476e8e7f50a0be8384c8,osType:Linux}},timeCreated:2026-05-18T08:02:28.9784295Z,vmId:85838c4c-38c3-4641-9333-9b5b9ab4769a}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbd85cf3epr9ha6omvqjeg,keypair:tbd85cf06pr9ha6omvqj3g,publicip:tbd85cf3epr9ha6omvqjeg-83151-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg"
            },
            {
              "key": "Name",
              "value": "tbd85cf3epr9ha6omvqjeg"
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
          "uid": "tbd85cf3epr9ha6omvqjdg",
          "cspResourceName": "tbd85cf3epr9ha6omvqjdg",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg",
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
          "createdTime": "2026-05-18 08:03:15",
          "label": {
            "createdBy": "tbd85cf3epr9ha6omvqjdg",
            "keypair": "tbd85cf06pr9ha6omvqj3g",
            "publicip": "tbd85cf3epr9ha6omvqjdg-86268-PublicIP",
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "azure-koreasouth",
            "sys.createdTime": "2026-05-18 08:03:15",
            "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg",
            "sys.cspResourceName": "tbd85cf3epr9ha6omvqjdg",
            "sys.id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my02-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my02-subnet-01",
            "sys.uid": "tbd85cf3epr9ha6omvqjdg",
            "sys.vNetId": "my02-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
          "region": {
            "region": "koreasouth"
          },
          "publicIP": "52.231.195.68",
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
          "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605150",
          "image": {
            "resourceType": "image",
            "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160"
          },
          "vNetId": "my02-vnet-01",
          "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg",
          "subnetId": "my02-subnet-01",
          "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0",
          "networkInterface": "tbd85cf3epr9ha6omvqjdg-3065-VNic",
          "securityGroupIds": [
            "my02-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my02-sshkey-01",
          "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g",
          "nodeUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBPTNwiwA06iMy3QoeSkT4pJXpOsH22Jn6+LnZ+ak239UDz03i5xCcIOmCYDASX7JMLHv1NfehWWjviZjXQOCm0A=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:zSgb3/NtuvOcCa3S2XoWo/Kj3Tpb4B15i3wLbr896JA",
            "firstUsedAt": "2026-05-18T08:03:25Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-18T08:03:23Z",
              "completedTime": "2026-05-18T08:03:28Z",
              "elapsedTime": 5,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux tbd85cf3epr9ha6omvqjdg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd85cf3epr9ha6omvqjdg-3065-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd85cf3epr9ha6omvqjdg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDpWcbrhtQvDFRYHly+CcVE71id44biB3QrFVXkBa7y8vHQa3Kj2v668qxkIfYEMUzInhuyqssv3AM5dKNvdl2cKz2BIBRNzpNxjGKEoHGNikkeKpFuEZxMhR5iIv1E36MU5mv6iAXDScum56Tp/97kKRSBMeZZo0Dq6b+SYUvFRUkBCBf84aIlQFD4KqiZ+Fe8j7V+vfOirroTVwZEaIUOSzfBvkZvKbpk8b86GzrxhnSgJyedkl7VoG+xpVg9iC7FjmIzyg9FK/8tk/FIje3BBmx2p83xhMyzquSopaJ4irYgDHRk5NWucds+nlJht5Zv9oxNsVXvo0eXu5Gg6EIi1S1zpPi6pEu1B9wx+0K5I93ZCMTrS2kbjnsPC/hubwXwGSAOdrm2uZAo8XYNH2LY8AyKHBAmBGtUwol7X7Dis4PLFkbApL2PLNovAHtLwpb0/+Zbqif4jMNV1ijaNsmN8n65MxuZ09SubadWql6JMh4glfS/8Rt7b0TJ/B9HtFSF1e+hGkt6nfxFtjT9ucyJvzQ+s+H6RZFUQYltLPFmqX0LSef1/xEVooKDofYtT5yFy/cE5/kFMbkKg2uWe1MkOlzwsBOXlmWDWX1U5pyZ9vg9wbtLM38qh+2Q3yBaroftUBxXy/nt9lDdPj/KnZV0NxGw4gQ82BBPYD80pfCZsw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202605150,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202605150},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd85cf3epr9ha6omvqjdg_OsDisk_1_c531642e3a2143bc9b21c3a62be80a3d,storageAccountType:Premium_LRS},name:tbd85cf3epr9ha6omvqjdg_OsDisk_1_c531642e3a2143bc9b21c3a62be80a3d,osType:Linux}},timeCreated:2026-05-18T08:02:27.2996702Z,vmId:9f82564d-bc46-434c-955e-4e8a0c6ad147}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:tbd85cf3epr9ha6omvqjdg,keypair:tbd85cf06pr9ha6omvqj3g,publicip:tbd85cf3epr9ha6omvqjdg-86268-PublicIP}"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg"
            },
            {
              "key": "Name",
              "value": "tbd85cf3epr9ha6omvqjdg"
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
            "nodeIp": "20.214.25.229",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbd85cf3epr9ha6omvqjcg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "20.214.26.6",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbd85cf3epr9ha6omvqjeg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my02-infra101",
            "nodeId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "nodeIp": "52.231.195.68",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux tbd85cf3epr9ha6omvqjdg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "tbd85cf1upr9ha6omvqj60",
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
        "sys.uid": "tbd85cf1upr9ha6omvqj60"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "node": [
        {
          "resourceType": "node",
          "id": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "tbd85cf1upr9ha6omvqj70",
          "cspResourceName": "tbd85cf1upr9ha6omvqj70",
          "cspResourceId": "i-mj7ctm2infbmnn6vq0ex",
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
          "createdTime": "2026-05-18 08:02:47",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-05-18 08:02:47",
            "sys.cspResourceId": "i-mj7ctm2infbmnn6vq0ex",
            "sys.cspResourceName": "tbd85cf1upr9ha6omvqj70",
            "sys.id": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "my04-subnet-01",
            "sys.uid": "tbd85cf1upr9ha6omvqj70",
            "sys.vNetId": "my04-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.240.250",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.15",
          "privateDNS": "",
          "rootDiskType": "cloud_essd",
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
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260316.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260316.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-vnet-01",
          "cspVNetId": "vpc-mj7kupnhl8vcko5nzmtjs",
          "subnetId": "my04-subnet-01",
          "cspSubnetId": "vsw-mj7noc100izqdsv9ei3ge",
          "networkInterface": "eni-mj7ctm2infbmnn6uucox",
          "securityGroupIds": [
            "my04-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-sshkey-01",
          "cspSshKeyId": "tbd85ceuupr9ha6omvqisg",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "upb!t1$1fd8c5A",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBPIxjgydAuYGlPu0chuI8q9j7hzoY0Rq4M/6ULrhSpcsPA78KSE4g2MEgG2KnC/Y9kaOYioaIMVk355QCS+/2JM=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:ElR8y8aTDUmfc8Kf9iCIhXjocBPKZg5jjX79lSxer8U",
            "firstUsedAt": "2026-05-18T08:02:56Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-18T08:02:55Z",
              "completedTime": "2026-05-18T08:02:57Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj7ctm2infbmnn6vq0exZ 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "tbd85cf1upr9ha6omvqj70"
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
              "value": "2026-05-18T08:02Z"
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
              "value": "iZmj7ctm2infbmnn6vq0exZ"
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
              "value": "13344c44-f69b-4fdb-bc04-e9e4756fde66"
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
              "value": "i-mj7ctm2infbmnn6vq0ex"
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
              "value": "2026-05-18T08:02Z"
            },
            {
              "key": "KeyPairName",
              "value": "tbd85ceuupr9ha6omvqisg"
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
              "value": "{SecurityGroupId:[sg-mj75hai0hn9b26o6u0lu]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[47.80.240.250]}"
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
              "value": "{VSwitchId:vsw-mj7noc100izqdsv9ei3ge,VpcId:vpc-mj7kupnhl8vcko5nzmtjs,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.15]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:08:0e:30,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7ctm2infbmnn6uucox,PrimaryIpAddress:10.0.1.15,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.15,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
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
          "uid": "tbd85cf1upr9ha6omvqj90",
          "cspResourceName": "tbd85cf1upr9ha6omvqj90",
          "cspResourceId": "i-mj70qfjkj53k1lfmwem5",
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
          "createdTime": "2026-05-18 08:02:47",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-05-18 08:02:47",
            "sys.cspResourceId": "i-mj70qfjkj53k1lfmwem5",
            "sys.cspResourceName": "tbd85cf1upr9ha6omvqj90",
            "sys.id": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "my04-subnet-01",
            "sys.uid": "tbd85cf1upr9ha6omvqj90",
            "sys.vNetId": "my04-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "8.220.222.227",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.17",
          "privateDNS": "",
          "rootDiskType": "cloud_essd",
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
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260316.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260316.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-vnet-01",
          "cspVNetId": "vpc-mj7kupnhl8vcko5nzmtjs",
          "subnetId": "my04-subnet-01",
          "cspSubnetId": "vsw-mj7noc100izqdsv9ei3ge",
          "networkInterface": "eni-mj70qfjkj53k1lfk3g4v",
          "securityGroupIds": [
            "my04-sg-03"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-sshkey-01",
          "cspSshKeyId": "tbd85ceuupr9ha6omvqisg",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "1bc15$udf8pAt!",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBBjpJO+iCEU9UWWxdK8tLAyQm+wo3x/MrYpIbM/i6AEC/aXL7gIArd+gJWLDbG4U7beYUDeqNCYoNaQ84LkntQ4=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:CN3Vr5UeDqGvqf+08il0YBKW/Q92T5KR3JBAFnULA/I",
            "firstUsedAt": "2026-05-18T08:02:56Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-18T08:02:55Z",
              "completedTime": "2026-05-18T08:02:57Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj70qfjkj53k1lfmwem5Z 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "tbd85cf1upr9ha6omvqj90"
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
              "value": "2026-05-18T08:02Z"
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
              "value": "iZmj70qfjkj53k1lfmwem5Z"
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
              "value": "59a11a5d-1a8e-4b85-8c4f-4e84016912bd"
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
              "value": "i-mj70qfjkj53k1lfmwem5"
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
              "value": "2026-05-18T08:02Z"
            },
            {
              "key": "KeyPairName",
              "value": "tbd85ceuupr9ha6omvqisg"
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
              "value": "{SecurityGroupId:[sg-mj78vtpurhkjrhyd2zt7]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[8.220.222.227]}"
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
              "value": "{VSwitchId:vsw-mj7noc100izqdsv9ei3ge,VpcId:vpc-mj7kupnhl8vcko5nzmtjs,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.17]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:08:0e:32,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj70qfjkj53k1lfk3g4v,PrimaryIpAddress:10.0.1.17,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.17,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
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
          "uid": "tbd85cf1upr9ha6omvqj80",
          "cspResourceName": "tbd85cf1upr9ha6omvqj80",
          "cspResourceId": "i-mj7bmodw2wgq13tm2g8j",
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
          "createdTime": "2026-05-18 08:02:47",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "alibaba-ap-northeast-2",
            "sys.createdTime": "2026-05-18 08:02:47",
            "sys.cspResourceId": "i-mj7bmodw2wgq13tm2g8j",
            "sys.cspResourceName": "tbd85cf1upr9ha6omvqj80",
            "sys.id": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.infraId": "my04-infra101",
            "sys.labelType": "node",
            "sys.manager": "cb-tumblebug",
            "sys.name": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.nodeGroupId": "my04-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "my04-subnet-01",
            "sys.uid": "tbd85cf1upr9ha6omvqj80",
            "sys.vNetId": "my04-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "47.80.57.53",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.16",
          "privateDNS": "",
          "rootDiskType": "cloud_essd",
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
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260316.vhd",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260506.vhd",
          "image": {
            "resourceType": "image",
            "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260316.vhd",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "Ubuntu  22.04 64 bit"
          },
          "vNetId": "my04-vnet-01",
          "cspVNetId": "vpc-mj7kupnhl8vcko5nzmtjs",
          "subnetId": "my04-subnet-01",
          "cspSubnetId": "vsw-mj7noc100izqdsv9ei3ge",
          "networkInterface": "eni-mj7bmodw2wgq13tslyzg",
          "securityGroupIds": [
            "my04-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "my04-sshkey-01",
          "cspSshKeyId": "tbd85ceuupr9ha6omvqisg",
          "nodeUserName": "cb-user",
          "nodeUserPassword": "$dA5bucp!f1t81",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBIoOQ2gvfcNkTlNvuqkusRt8H2qhBA89swB2tyZNTyPQI6IG5Rq30BzABy4uD/mVN5T1FFu4vaI+o5kEgApc87Q=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:Y6GQw48WwR78o9JdeZKtnxMOFr9e5MIESK+ZhUdpvPM",
            "firstUsedAt": "2026-05-18T08:02:55Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-05-18T08:02:55Z",
              "completedTime": "2026-05-18T08:02:56Z",
              "elapsedTime": 1,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux iZmj7bmodw2wgq13tm2g8jZ 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "tbd85cf1upr9ha6omvqj80"
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
              "value": "2026-05-18T08:02Z"
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
              "value": "iZmj7bmodw2wgq13tm2g8jZ"
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
              "value": "3e72e546-fbd2-4921-a2ff-d588c6c712fc"
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
              "value": "i-mj7bmodw2wgq13tm2g8j"
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
              "value": "2026-05-18T08:02Z"
            },
            {
              "key": "KeyPairName",
              "value": "tbd85ceuupr9ha6omvqisg"
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
              "value": "{SecurityGroupId:[sg-mj71whg4s0wud8qlexce]}"
            },
            {
              "key": "InnerIpAddress",
              "value": "{IpAddress:[]}"
            },
            {
              "key": "PublicIpAddress",
              "value": "{IpAddress:[47.80.57.53]}"
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
              "value": "{VSwitchId:vsw-mj7noc100izqdsv9ei3ge,VpcId:vpc-mj7kupnhl8vcko5nzmtjs,NatIpAddress:,PrivateIpAddress:{IpAddress:[10.0.1.16]}}"
            },
            {
              "key": "Tags",
              "value": "{Tag:null}"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:08:0e:31,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7bmodw2wgq13tslyzg,PrimaryIpAddress:10.0.1.16,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:10.0.1.16,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
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
            "nodeIp": "47.80.57.53",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7bmodw2wgq13tm2g8jZ 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "nodeIp": "47.80.240.250",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj7ctm2infbmnn6vq0exZ 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "infraId": "my04-infra101",
            "nodeId": "my04-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "nodeIp": "8.220.222.227",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux iZmj70qfjkj53k1lfmwem5Z 5.15.0-177-generic #187-Ubuntu SMP Sat Apr 11 22:54:33 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "id": "my02-infra101",
  "uid": "tbd85cf3epr9ha6omvqjbg",
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
    "sys.uid": "tbd85cf3epr9ha6omvqjbg"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "node": [
    {
      "resourceType": "node",
      "id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "tbd85cf3epr9ha6omvqjcg",
      "cspResourceName": "tbd85cf3epr9ha6omvqjcg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg",
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
      "createdTime": "2026-05-18 08:03:14",
      "label": {
        "createdBy": "tbd85cf3epr9ha6omvqjcg",
        "keypair": "tbd85cf06pr9ha6omvqj3g",
        "publicip": "tbd85cf3epr9ha6omvqjcg-42454-PublicIP",
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-18 08:03:14",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg",
        "sys.cspResourceName": "tbd85cf3epr9ha6omvqjcg",
        "sys.id": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd85cf3epr9ha6omvqjcg",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=51.2% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.214.25.229",
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605150",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0",
      "networkInterface": "tbd85cf3epr9ha6omvqjcg-39198-VNic",
      "securityGroupIds": [
        "my02-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBL0VTwP0mj+c6VjV8msgi3Q9eibNRewPMT54DUhTIpAg1Sht7ha22G1ufqkYptq1cFE80rR711dScHPLLhmJ+q8=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:L+I1/Qhho+5MTJJQUC536mZqODIvIydVP6vJjVbdai0",
        "firstUsedAt": "2026-05-18T08:03:23Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-18T08:03:23Z",
          "completedTime": "2026-05-18T08:03:26Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd85cf3epr9ha6omvqjcg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2als_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd85cf3epr9ha6omvqjcg-39198-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd85cf3epr9ha6omvqjcg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDpWcbrhtQvDFRYHly+CcVE71id44biB3QrFVXkBa7y8vHQa3Kj2v668qxkIfYEMUzInhuyqssv3AM5dKNvdl2cKz2BIBRNzpNxjGKEoHGNikkeKpFuEZxMhR5iIv1E36MU5mv6iAXDScum56Tp/97kKRSBMeZZo0Dq6b+SYUvFRUkBCBf84aIlQFD4KqiZ+Fe8j7V+vfOirroTVwZEaIUOSzfBvkZvKbpk8b86GzrxhnSgJyedkl7VoG+xpVg9iC7FjmIzyg9FK/8tk/FIje3BBmx2p83xhMyzquSopaJ4irYgDHRk5NWucds+nlJht5Zv9oxNsVXvo0eXu5Gg6EIi1S1zpPi6pEu1B9wx+0K5I93ZCMTrS2kbjnsPC/hubwXwGSAOdrm2uZAo8XYNH2LY8AyKHBAmBGtUwol7X7Dis4PLFkbApL2PLNovAHtLwpb0/+Zbqif4jMNV1ijaNsmN8n65MxuZ09SubadWql6JMh4glfS/8Rt7b0TJ/B9HtFSF1e+hGkt6nfxFtjT9ucyJvzQ+s+H6RZFUQYltLPFmqX0LSef1/xEVooKDofYtT5yFy/cE5/kFMbkKg2uWe1MkOlzwsBOXlmWDWX1U5pyZ9vg9wbtLM38qh+2Q3yBaroftUBxXy/nt9lDdPj/KnZV0NxGw4gQ82BBPYD80pfCZsw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202605150,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202605150},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd85cf3epr9ha6omvqjcg_OsDisk_1_15c2d99efd01471b8724e152bfe481ca,storageAccountType:Premium_LRS},name:tbd85cf3epr9ha6omvqjcg_OsDisk_1_15c2d99efd01471b8724e152bfe481ca,osType:Linux}},timeCreated:2026-05-18T08:02:26.9394222Z,vmId:66b1ceb1-9da8-497a-ab77-8e1d40b4f26c}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd85cf3epr9ha6omvqjcg,keypair:tbd85cf06pr9ha6omvqj3g,publicip:tbd85cf3epr9ha6omvqjcg-42454-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg"
        },
        {
          "key": "Name",
          "value": "tbd85cf3epr9ha6omvqjcg"
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
      "uid": "tbd85cf3epr9ha6omvqjeg",
      "cspResourceName": "tbd85cf3epr9ha6omvqjeg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg",
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
      "createdTime": "2026-05-18 08:03:17",
      "label": {
        "createdBy": "tbd85cf3epr9ha6omvqjeg",
        "keypair": "tbd85cf06pr9ha6omvqj3g",
        "publicip": "tbd85cf3epr9ha6omvqjeg-83151-PublicIP",
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-18 08:03:17",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg",
        "sys.cspResourceName": "tbd85cf3epr9ha6omvqjeg",
        "sys.id": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd85cf3epr9ha6omvqjeg",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "20.214.26.6",
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202605150",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0",
      "networkInterface": "tbd85cf3epr9ha6omvqjeg-47871-VNic",
      "securityGroupIds": [
        "my02-sg-03"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEjxGVDmlxPviCew7MPsxQCg/SbRH0g1LPINkzdzqnLKnu1r1BVcN2sOxPrrm4d3ektW4i1wjM5H3IyYVNdYIXM=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:kXaggWkouCHVkByZEeZ7WbredHOydMu4PhHCk95sdGA",
        "firstUsedAt": "2026-05-18T08:03:25Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-18T08:03:23Z",
          "completedTime": "2026-05-18T08:03:27Z",
          "elapsedTime": 4,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd85cf3epr9ha6omvqjeg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B2as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd85cf3epr9ha6omvqjeg-47871-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd85cf3epr9ha6omvqjeg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDpWcbrhtQvDFRYHly+CcVE71id44biB3QrFVXkBa7y8vHQa3Kj2v668qxkIfYEMUzInhuyqssv3AM5dKNvdl2cKz2BIBRNzpNxjGKEoHGNikkeKpFuEZxMhR5iIv1E36MU5mv6iAXDScum56Tp/97kKRSBMeZZo0Dq6b+SYUvFRUkBCBf84aIlQFD4KqiZ+Fe8j7V+vfOirroTVwZEaIUOSzfBvkZvKbpk8b86GzrxhnSgJyedkl7VoG+xpVg9iC7FjmIzyg9FK/8tk/FIje3BBmx2p83xhMyzquSopaJ4irYgDHRk5NWucds+nlJht5Zv9oxNsVXvo0eXu5Gg6EIi1S1zpPi6pEu1B9wx+0K5I93ZCMTrS2kbjnsPC/hubwXwGSAOdrm2uZAo8XYNH2LY8AyKHBAmBGtUwol7X7Dis4PLFkbApL2PLNovAHtLwpb0/+Zbqif4jMNV1ijaNsmN8n65MxuZ09SubadWql6JMh4glfS/8Rt7b0TJ/B9HtFSF1e+hGkt6nfxFtjT9ucyJvzQ+s+H6RZFUQYltLPFmqX0LSef1/xEVooKDofYtT5yFy/cE5/kFMbkKg2uWe1MkOlzwsBOXlmWDWX1U5pyZ9vg9wbtLM38qh+2Q3yBaroftUBxXy/nt9lDdPj/KnZV0NxGw4gQ82BBPYD80pfCZsw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],imageReference:{exactVersion:22.04.202605150,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts,version:22.04.202605150},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd85cf3epr9ha6omvqjeg_OsDisk_1_8aaddc61d3c5476e8e7f50a0be8384c8,storageAccountType:Premium_LRS},name:tbd85cf3epr9ha6omvqjeg_OsDisk_1_8aaddc61d3c5476e8e7f50a0be8384c8,osType:Linux}},timeCreated:2026-05-18T08:02:28.9784295Z,vmId:85838c4c-38c3-4641-9333-9b5b9ab4769a}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd85cf3epr9ha6omvqjeg,keypair:tbd85cf06pr9ha6omvqj3g,publicip:tbd85cf3epr9ha6omvqjeg-83151-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg"
        },
        {
          "key": "Name",
          "value": "tbd85cf3epr9ha6omvqjeg"
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
      "uid": "tbd85cf3epr9ha6omvqjdg",
      "cspResourceName": "tbd85cf3epr9ha6omvqjdg",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg",
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
      "createdTime": "2026-05-18 08:03:15",
      "label": {
        "createdBy": "tbd85cf3epr9ha6omvqjdg",
        "keypair": "tbd85cf06pr9ha6omvqj3g",
        "publicip": "tbd85cf3epr9ha6omvqjdg-86268-PublicIP",
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "azure-koreasouth",
        "sys.createdTime": "2026-05-18 08:03:15",
        "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg",
        "sys.cspResourceName": "tbd85cf3epr9ha6omvqjdg",
        "sys.id": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.infraId": "my02-infra101",
        "sys.labelType": "node",
        "sys.manager": "cb-tumblebug",
        "sys.name": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.nodeGroupId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "my02-subnet-01",
        "sys.uid": "tbd85cf3epr9ha6omvqjdg",
        "sys.vNetId": "my02-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=97.7% Image=100.0%",
      "region": {
        "region": "koreasouth"
      },
      "publicIP": "52.231.195.68",
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
      "imageId": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202605150",
      "image": {
        "resourceType": "image",
        "cspImageName": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160"
      },
      "vNetId": "my02-vnet-01",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg",
      "subnetId": "my02-subnet-01",
      "cspSubnetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0",
      "networkInterface": "tbd85cf3epr9ha6omvqjdg-3065-VNic",
      "securityGroupIds": [
        "my02-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "my02-sshkey-01",
      "cspSshKeyId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g",
      "nodeUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBPTNwiwA06iMy3QoeSkT4pJXpOsH22Jn6+LnZ+ak239UDz03i5xCcIOmCYDASX7JMLHv1NfehWWjviZjXQOCm0A=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:zSgb3/NtuvOcCa3S2XoWo/Kj3Tpb4B15i3wLbr896JA",
        "firstUsedAt": "2026-05-18T08:03:25Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-05-18T08:03:23Z",
          "completedTime": "2026-05-18T08:03:28Z",
          "elapsedTime": 5,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux tbd85cf3epr9ha6omvqjdg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkInterfaces/tbd85cf3epr9ha6omvqjdg-3065-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:tbd85cf3epr9ha6omvqjdg,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDpWcbrhtQvDFRYHly+CcVE71id44biB3QrFVXkBa7y8vHQa3Kj2v668qxkIfYEMUzInhuyqssv3AM5dKNvdl2cKz2BIBRNzpNxjGKEoHGNikkeKpFuEZxMhR5iIv1E36MU5mv6iAXDScum56Tp/97kKRSBMeZZo0Dq6b+SYUvFRUkBCBf84aIlQFD4KqiZ+Fe8j7V+vfOirroTVwZEaIUOSzfBvkZvKbpk8b86GzrxhnSgJyedkl7VoG+xpVg9iC7FjmIzyg9FK/8tk/FIje3BBmx2p83xhMyzquSopaJ4irYgDHRk5NWucds+nlJht5Zv9oxNsVXvo0eXu5Gg6EIi1S1zpPi6pEu1B9wx+0K5I93ZCMTrS2kbjnsPC/hubwXwGSAOdrm2uZAo8XYNH2LY8AyKHBAmBGtUwol7X7Dis4PLFkbApL2PLNovAHtLwpb0/+Zbqif4jMNV1ijaNsmN8n65MxuZ09SubadWql6JMh4glfS/8Rt7b0TJ/B9HtFSF1e+hGkt6nfxFtjT9ucyJvzQ+s+H6RZFUQYltLPFmqX0LSef1/xEVooKDofYtT5yFy/cE5/kFMbkKg2uWe1MkOlzwsBOXlmWDWX1U5pyZ9vg9wbtLM38qh+2Q3yBaroftUBxXy/nt9lDdPj/KnZV0NxGw4gQ82BBPYD80pfCZsw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202605150,offer:0001-com-ubuntu-server-jammy-daily,publisher:Canonical,sku:22_04-daily-lts-gen2,version:22.04.202605150},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/disks/tbd85cf3epr9ha6omvqjdg_OsDisk_1_c531642e3a2143bc9b21c3a62be80a3d,storageAccountType:Premium_LRS},name:tbd85cf3epr9ha6omvqjdg_OsDisk_1_c531642e3a2143bc9b21c3a62be80a3d,osType:Linux}},timeCreated:2026-05-18T08:02:27.2996702Z,vmId:9f82564d-bc46-434c-955e-4e8a0c6ad147}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:tbd85cf3epr9ha6omvqjdg,keypair:tbd85cf06pr9ha6omvqj3g,publicip:tbd85cf3epr9ha6omvqjdg-86268-PublicIP}"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg"
        },
        {
          "key": "Name",
          "value": "tbd85cf3epr9ha6omvqjdg"
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
        "nodeIp": "20.214.25.229",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd85cf3epr9ha6omvqjcg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my02-infra101",
        "nodeId": "my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "nodeIp": "20.214.26.6",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd85cf3epr9ha6omvqjeg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "infraId": "my02-infra101",
        "nodeId": "my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "nodeIp": "52.231.195.68",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux tbd85cf3epr9ha6omvqjdg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "output": "Linux tbd85cf3epr9ha6omvqjcg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "20.214.25.229",
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
      "output": "Linux tbd85cf3epr9ha6omvqjeg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "20.214.26.6",
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
      "output": "Linux tbd85cf3epr9ha6omvqjdg 6.8.0-1052-azure #58~22.04.1-Ubuntu SMP Thu Mar 26 05:02:21 UTC 2026 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "52.231.195.68",
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

**Generated At:** 2026-05-18 08:04:01

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
| Standard_B4as_v2 | 4 | 15.6 | - | x86_64 |  | $0.1730 | 1 |
| Standard_B2als_v2 | 2 | 3.9 | - | x86_64 |  | $0.0432 | 1 |
| Standard_B2as_v2 | 2 | 7.8 | - | x86_64 |  | $0.0865 | 1 |

### VM Images

| Name | Distribution | OS Type | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
|------|--------------|---------|-------------|--------------|----------------|----------------|----------------------|
| Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 2 |
| Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160 | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160 | Ubuntu 22.04 | Linux/UNIX | x86_64 | default | - | 1 |

### Virtual Machines

| VM Name | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |
|---------|-----------|--------|-------------------------|-------|------|
| my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg | Running | 2 vCPU, 3.9 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160) | **VNet:** my02-vnet-01<br>**Subnet:** my02-subnet-01<br>**Public IP:** 20.214.25.229<br>**Private IP:** 10.0.1.5<br>**SGs:** my02-sg-01<br>**SSH:** my02-sshkey-01 |
| my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg | Running | 2 vCPU, 7.8 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160) | **VNet:** my02-vnet-01<br>**Subnet:** my02-subnet-01<br>**Public IP:** 20.214.26.6<br>**Private IP:** 10.0.1.6<br>**SGs:** my02-sg-03<br>**SSH:** my02-sshkey-01 |
| my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg | Running | 4 vCPU, 15.6 GiB | Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160 (Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160) | **VNet:** my02-vnet-01<br>**Subnet:** my02-subnet-01<br>**Public IP:** 52.231.195.68<br>**Private IP:** 10.0.1.4<br>**SGs:** my02-sg-02<br>**SSH:** my02-sshkey-01 |


## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: my02-vnet-01

| Property | Value |
|----------|-------|
| **Name** | my02-vnet-01 |
| **CSP VNet ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg |
| **CIDR Block** | 10.0.0.0/21 |
| **Connection** | azure-koreasouth |
| **Subnet Count** | 1 |

**Subnets:**

| Name | CSP Subnet ID | CIDR Block | Zone |
|------|---------------|------------|------|
| my02-subnet-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0 | 10.0.1.0/24 |  |


## Security Resources

### SSH Keys

| Name | CSP SSH Key ID | Username | Fingerprint |
|------|----------------|----------|-------------|
| my02-sshkey-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g |  |  |

### Security Groups

#### Security Group: my02-sg-01

| Property | Value |
|----------|-------|
| **Name** | my02-sg-01 |
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd85cf0mpr9ha6omvqj4g |
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
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd85cf1mpr9ha6omvqj5g |
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
| **CSP Security Group ID** | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd85cf2mpr9ha6omvqjb0 |
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

*Report generated: 2026-05-18 08:04:07*

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
| 2 | **VM Spec** | 3 | ✅ Selected | Standard_B2as_v2, Standard_B4as_v2, Standard_B2als_v2 |
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
| 1 | **VM Name:** my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjcg<br>**Label(sourceMachineId):** vm-ec268ed7-821e-9d73-e79f | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f |
| 2 | **VM Name:** my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjeg<br>**Label(sourceMachineId):** vm-ec288dd0-c6fa-8a49-2f60 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 |
| 3 | **VM Name:** my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/virtualMachines/tbd85cf3epr9ha6omvqjdg<br>**Label(sourceMachineId):** vm-ec2d32b5-98fb-5a96-7913 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 |

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
| 1 | my02-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 | **Hostname:** N/A<br>**Machine ID:** vm-ec268ed7-821e-9d73-e79f | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2 | my02-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts:22.04.202604160 | **Hostname:** N/A<br>**Machine ID:** vm-ec288dd0-c6fa-8a49-2f60 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3 | my02-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Canonical:0001-com-ubuntu-server-jammy-daily:22_04-daily-lts-gen2:22.04.202604160 | **Hostname:** N/A<br>**Machine ID:** vm-ec2d32b5-98fb-5a96-7913 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: my02-sg-01

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd85cf0mpr9ha6omvqj4g | **VNet:** my02-vnet-01 | **Rules:** 14

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

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd85cf1mpr9ha6omvqj5g | **VNet:** my02-vnet-01 | **Rules:** 19

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

**CSP ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/networkSecurityGroups/tbd85cf2mpr9ha6omvqjb0 | **VNet:** my02-vnet-01 | **Rules:** 19

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
| 1 | **Name:** my02-vnet-01<br>**ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg | 10.0.0.0/21 |

### Subnets

| No. | Subnet | CIDR Block | Associated VPC(VNet) |
|-----|--------|------------|----------------------|
| 1 | **Name:** my02-subnet-01<br>**ID:** /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbd85cetupr9ha6omvqiqg/subnets/tbd85cetupr9ha6omvqir0 | 10.0.1.0/24 | my02-vnet-01 |

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
| 1 | my02-sshkey-01 | /subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/KOREASOUTH/providers/Microsoft.Compute/sshPublicKeys/tbd85cf06pr9ha6omvqj3g |  | Used by all 3 VMs |

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

