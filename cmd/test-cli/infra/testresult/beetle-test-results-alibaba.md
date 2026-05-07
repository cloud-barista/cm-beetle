# CM-Beetle test results for ALIBABA

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with ALIBABA cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.5.0+ (1655425)
- cm-model: v0.0.20
- CB-Tumblebug: v0.12.3
- CB-Spider: v0.12.11
- CB-MapUI: v0.12.16
- Target CSP: ALIBABA
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: March 25, 2026
- Test Time: 22:30:49 KST
- Test Execution: 2026-03-25 22:30:49 KST

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

| Test | Step (Endpoint / Description)                          | Status      | Duration | Details |
| ---- | ------------------------------------------------------ | ----------- | -------- | ------- |
| 1    | `POST /beetle/recommendation/vmInfra`                  | ✅ **PASS** | 407ms    | Pass    |
| 2    | `POST /beetle/migration/ns/mig01/mci`                  | ❌ **FAIL** | 27.775s  | Fail    |
| 3    | `GET /beetle/migration/ns/mig01/mci`                   | ⏭️ **SKIP** | 0s       | Skip    |
| 4    | `GET /beetle/migration/ns/mig01/mci?option=id`         | ⏭️ **SKIP** | 0s       | Skip    |
| 5    | `GET /beetle/migration/ns/mig01/mci/{{mciId}}`         | ⏭️ **SKIP** | 0s       | Skip    |
| 6    | Remote Command Accessibility Check                     | ⏭️ **SKIP** | 0s       | Skip    |
| 7    | `GET /beetle/summary/target/ns/mig01/mci/{{mciId}}`    | ✅ **PASS** | 6.133s   | Pass    |
| 8    | `POST /beetle/report/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 6.14s    | Pass    |
| 9    | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}`      | ✅ **PASS** | 28.901s  | Pass    |

**Overall Result**: 4/9 tests passed, 4 skipped ❌

**Total Duration**: 1m49.468521764s

_Test executed on March 25, 2026 at 22:30:49 KST (2026-03-25 22:30:49 KST) using CM-Beetle automated test CLI_

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
  "nameSeed": "mig-4",
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
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536,
            "state": "up"
          },
          {
            "name": "ens5",
            "macAddress": "02:6f:de:fc:71:b1",
            "ipv4CidrBlocks": ["10.0.1.30/24"],
            "ipv6CidrBlocks": ["fe80::6f:deff:fefc:71b1/64"],
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
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536,
            "state": "up"
          },
          {
            "name": "ens5",
            "macAddress": "02:08:96:7d:f4:17",
            "ipv4CidrBlocks": ["10.0.1.221/24"],
            "ipv6CidrBlocks": ["fe80::8:96ff:fe7d:f417/64"],
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
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536,
            "state": "up"
          },
          {
            "name": "ens5",
            "macAddress": "02:bf:6e:6c:6e:31",
            "ipv4CidrBlocks": ["10.0.1.138/24"],
            "ipv6CidrBlocks": ["fe80::bf:6eff:fe6c:6e31/64"],
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
      "nameSeed": "mig-4",
      "status": "partially-matched",
      "description": "Candidate #1 | partially-matched | Overall Match Rate: Min=75.0% Max=100.0% Avg=91.7% | VMs: 3 total, 0 matched, 3 acceptable",
      "targetCloud": {
        "csp": "alibaba",
        "region": "ap-northeast-2"
      },
      "targetVmInfra": {
        "name": "mci101",
        "installMonAgent": "",
        "label": null,
        "systemLabel": "",
        "description": "Recommended VMs comprising multi-cloud infrastructure",
        "subGroups": [
          {
            "name": "mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
            "vNetId": "mig-4-vnet-01",
            "subnetId": "mig-4-subnet-01",
            "securityGroupIds": ["mig-4-sg-01"],
            "sshKeyId": "mig-4-sshkey-01",
            "rootDiskType": "TYPE1",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
            "vNetId": "mig-4-vnet-01",
            "subnetId": "mig-4-subnet-01",
            "securityGroupIds": ["mig-4-sg-02"],
            "sshKeyId": "mig-4-sshkey-01",
            "rootDiskType": "TYPE1",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
            "vNetId": "mig-4-vnet-01",
            "subnetId": "mig-4-subnet-01",
            "securityGroupIds": ["mig-4-sg-03"],
            "sshKeyId": "mig-4-sshkey-01",
            "rootDiskType": "TYPE1",
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
        "connectionName": "alibaba-ap-northeast-2",
        "cidrBlock": "10.0.0.0/21",
        "subnetInfoList": [
          {
            "name": "mig-4-subnet-01",
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
      "targetVmSpecList": [
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "uid": "d66sfhddi7idhnupq6fg",
          "cspSpecName": "ecs.t6-c1m1.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
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
              "key": "MemorySize",
              "value": "2.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "100000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "2"
            },
            {
              "key": "CpuCoreCount",
              "value": "2"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m1.large"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "81920"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "60"
            },
            {
              "key": "InstancePpsTx",
              "value": "100000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "81920"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "40"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "uid": "d66sfhddi7idhnupq6gg",
          "cspSpecName": "ecs.t6-c1m4.xlarge",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
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
              "key": "MemorySize",
              "value": "16.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "200000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "6"
            },
            {
              "key": "CpuCoreCount",
              "value": "4"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m4.xlarge"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "163840"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "120"
            },
            {
              "key": "InstancePpsTx",
              "value": "200000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "163840"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "160"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "uid": "d66sfhddi7idhnupq6hg",
          "cspSpecName": "ecs.t6-c1m4.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
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
              "key": "MemorySize",
              "value": "8.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "100000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "2"
            },
            {
              "key": "CpuCoreCount",
              "value": "2"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m4.large"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "81920"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "60"
            },
            {
              "key": "InstancePpsTx",
              "value": "100000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "81920"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "60"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        }
      ],
      "targetVmOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
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
            "cn-huhehaote",
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
          "id": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
          "uid": "d66sfo5di7idhnuqkh60",
          "name": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-ap-northeast-1",
          "infraType": "",
          "fetchedTime": "2026.02.12 12:30:24 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260119.vhd"
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
              "value": "Kernel version is 5.15.0-164-generic, 2026.1.22"
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
              "value": "v2026.1.22"
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
              "value": "2026-01-22T10:59:29Z"
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260119.vhd"
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
          "description": "Kernel version is 5.15.0-164-generic, 2026.1.22",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-4-sg-01",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-4-vnet-01",
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
          "name": "mig-4-sg-02",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-4-vnet-01",
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
          "name": "mig-4-sg-03",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-4-vnet-01",
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
      "nameSeed": "mig-4",
      "status": "partially-matched",
      "description": "Candidate #2 | partially-matched | Overall Match Rate: Min=50.0% Max=100.0% Avg=75.0% | VMs: 1 total, 0 matched, 1 acceptable",
      "targetCloud": {
        "csp": "alibaba",
        "region": "ap-northeast-2"
      },
      "targetVmInfra": {
        "name": "mci101",
        "installMonAgent": "",
        "label": null,
        "systemLabel": "",
        "description": "Recommended VMs comprising multi-cloud infrastructure",
        "subGroups": [
          {
            "name": "mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=50.0% Image=75.0%",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "alibaba+ap-northeast-2+ecs.t6-c2m1.large",
            "imageId": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
            "vNetId": "mig-4-vnet-01",
            "subnetId": "mig-4-subnet-01",
            "securityGroupIds": ["mig-4-sg-01"],
            "sshKeyId": "mig-4-sshkey-01",
            "rootDiskType": "TYPE1",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "a recommended virtual machine 02 for ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "",
            "imageId": "",
            "vNetId": "mig-4-vnet-01",
            "subnetId": "mig-4-subnet-01",
            "securityGroupIds": ["mig-4-sg-02"],
            "sshKeyId": "mig-4-sshkey-01",
            "rootDiskType": "TYPE1",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "a recommended virtual machine 03 for ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "connectionName": "alibaba-ap-northeast-2",
            "specId": "",
            "imageId": "",
            "vNetId": "mig-4-vnet-01",
            "subnetId": "mig-4-subnet-01",
            "securityGroupIds": ["mig-4-sg-03"],
            "sshKeyId": "mig-4-sshkey-01",
            "rootDiskType": "TYPE1",
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
        "connectionName": "alibaba-ap-northeast-2",
        "cidrBlock": "10.0.0.0/21",
        "subnetInfoList": [
          {
            "name": "mig-4-subnet-01",
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
      "targetVmSpecList": [
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "uid": "d66sfhddi7idhnupq6fg",
          "cspSpecName": "ecs.t6-c1m1.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
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
              "key": "MemorySize",
              "value": "2.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "100000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "2"
            },
            {
              "key": "CpuCoreCount",
              "value": "2"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m1.large"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "81920"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "60"
            },
            {
              "key": "InstancePpsTx",
              "value": "100000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "81920"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "40"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "uid": "d66sfhddi7idhnupq6gg",
          "cspSpecName": "ecs.t6-c1m4.xlarge",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
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
              "key": "MemorySize",
              "value": "16.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "200000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "6"
            },
            {
              "key": "CpuCoreCount",
              "value": "4"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m4.xlarge"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "163840"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "120"
            },
            {
              "key": "InstancePpsTx",
              "value": "200000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "163840"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "160"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "uid": "d66sfhddi7idhnupq6hg",
          "cspSpecName": "ecs.t6-c1m4.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
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
              "key": "MemorySize",
              "value": "8.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "100000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "2"
            },
            {
              "key": "CpuCoreCount",
              "value": "2"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c1m4.large"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "81920"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "60"
            },
            {
              "key": "InstancePpsTx",
              "value": "100000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "81920"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "60"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        },
        {
          "id": "alibaba+ap-northeast-2+ecs.t6-c2m1.large",
          "uid": "d66sfhddi7idhnupq6f0",
          "cspSpecName": "ecs.t6-c2m1.large",
          "name": "alibaba+ap-northeast-2+ecs.t6-c2m1.large",
          "namespace": "system",
          "connectionName": "alibaba-ap-northeast-2",
          "providerName": "alibaba",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 1,
          "diskSizeGB": -1,
          "costPerHour": 0.01116,
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
              "key": "MemorySize",
              "value": "1.00"
            },
            {
              "key": "InstancePpsRx",
              "value": "60000"
            },
            {
              "key": "EriQuantity",
              "value": "0"
            },
            {
              "key": "EniPrivateIpAddressQuantity",
              "value": "2"
            },
            {
              "key": "CpuCoreCount",
              "value": "2"
            },
            {
              "key": "EniTotalQuantity",
              "value": "2"
            },
            {
              "key": "NetworkEncryptionSupport",
              "value": "false"
            },
            {
              "key": "Cores",
              "value": "0"
            },
            {
              "key": "NetworkCardQuantity",
              "value": "0"
            },
            {
              "key": "JumboFrameSupport",
              "value": "false"
            },
            {
              "key": "InstanceTypeId",
              "value": "ecs.t6-c2m1.large"
            },
            {
              "key": "InstanceBandwidthRx",
              "value": "81920"
            },
            {
              "key": "QueuePairNumber",
              "value": "0"
            },
            {
              "key": "EniQuantity",
              "value": "2"
            },
            {
              "key": "InstanceTypeFamily",
              "value": "ecs.t6"
            },
            {
              "key": "InitialCredit",
              "value": "60"
            },
            {
              "key": "InstancePpsTx",
              "value": "60000"
            },
            {
              "key": "InstanceFamilyLevel",
              "value": "CreditEntryLevel"
            },
            {
              "key": "LocalStorageAmount",
              "value": "0"
            },
            {
              "key": "TotalEniQueueQuantity",
              "value": "2"
            },
            {
              "key": "CpuArchitecture",
              "value": "X86"
            },
            {
              "key": "SecondaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "InstanceBandwidthTx",
              "value": "81920"
            },
            {
              "key": "MaximumQueueNumberPerEni",
              "value": "0"
            },
            {
              "key": "DiskQuantity",
              "value": "17"
            },
            {
              "key": "PrimaryEniQueueNumber",
              "value": "1"
            },
            {
              "key": "Memory",
              "value": "0"
            },
            {
              "key": "CpuTurboFrequency",
              "value": "3.20"
            },
            {
              "key": "BaselineCredit",
              "value": "20"
            },
            {
              "key": "EniTrunkSupported",
              "value": "false"
            },
            {
              "key": "GPUAmount",
              "value": "0"
            },
            {
              "key": "GPUMemorySize",
              "value": "0.00"
            },
            {
              "key": "NvmeSupport",
              "value": "unsupported"
            },
            {
              "key": "InstanceCategory",
              "value": "Shared"
            },
            {
              "key": "EniIpv6AddressQuantity",
              "value": "1"
            },
            {
              "key": "LocalStorageCapacity",
              "value": "0"
            },
            {
              "key": "CpuSpeedFrequency",
              "value": "2.50"
            },
            {
              "key": "PhysicalProcessorModel",
              "value": "Intel Xeon (Cascade Lake) Platinum 8269CY"
            },
            {
              "key": "SupportedBootModes",
              "value": "{SupportedBootMode:[BIOS,UEFI]}"
            },
            {
              "key": "EnhancedNetwork",
              "value": "{EnableSriov:false,SriovSupport:false,RssSupport:false,VfQueueNumberPerEni:0,EnableRss:false}"
            },
            {
              "key": "CpuOptions",
              "value": "{Core:0,HyperThreadingAdjustable:false,CoreCount:0,CoreFactor:0,Numa:,TopologyType:,ThreadsPerCore:0,SupportedTopologyTypes:{SupportedTopologyType:null}}"
            },
            {
              "key": "NetworkCards",
              "value": "{NetworkCardInfo:null}"
            }
          ]
        }
      ],
      "targetVmOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "alibaba",
          "cspImageName": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
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
            "cn-huhehaote",
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
          "id": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
          "uid": "d66sfo5di7idhnuqkh60",
          "name": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "alibaba-ap-northeast-1",
          "infraType": "",
          "fetchedTime": "2026.02.12 12:30:24 Thu",
          "creationDate": "",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260119.vhd"
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
              "value": "Kernel version is 5.15.0-164-generic, 2026.1.22"
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
              "value": "v2026.1.22"
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
              "value": "2026-01-22T10:59:29Z"
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
              "value": "ubuntu_22_04_x64_20G_alibase_20260119.vhd"
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
          "description": "Kernel version is 5.15.0-164-generic, 2026.1.22",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-4-sg-01",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-4-vnet-01",
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
          "name": "mig-4-sg-02",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-4-vnet-01",
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
          "name": "mig-4-sg-03",
          "connectionName": "alibaba-ap-northeast-2",
          "vNetId": "mig-4-vnet-01",
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
  "resourceType": "",
  "id": "",
  "name": "",
  "status": "",
  "statusCount": {
    "countTotal": 0,
    "countCreating": 0,
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
  "targetStatus": "",
  "targetAction": "",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": null,
  "systemLabel": "",
  "systemMessage": null,
  "description": "",
  "vm": null,
  "newVmList": null,
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

### Test Case 3: Get a list of MCIs

#### 3.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/mci`
- **Purpose**: Retrieve all Multi-Cloud Infrastructure instances
- **Namespace ID**: `mig01`
- **Request Body**: None (GET request)

#### 3.2 API Response Information

- **Status**: ❌ **FAILED**
- **Error**: No response received

**Error Message**:

```
Test skipped due to previous test failure
```

### Test Case 4: Get a list of MCI IDs

#### 4.1 API Request Information

- **API Endpoint**: `GET /beetle/migration/ns/mig01/mci?option=id`
- **Purpose**: Retrieve MCI IDs only (lightweight response)
- **Namespace ID**: `mig01`
- **Query Parameter**: `option=id`
- **Request Body**: None (GET request)

#### 4.2 API Response Information

- **Status**: ❌ **FAILED**
- **Error**: No response received

**Error Message**:

```
Test skipped due to previous test failure
```

### Test Case 6: Remote Command Accessibility Check

#### 6.1 Test Information

- **Test Type**: SSH Connectivity Test for All VMs
- **Purpose**: Verify that all migrated VMs are accessible via SSH
- **Method**: Extract public IP and SSH key from MCI access info for each VM, then execute remote command
- **Command Executed**: `uname -a` (to verify system information)
- **Authentication**: SSH key-based authentication
- **Scope**: Tests all VMs across all subgroups in the MCI

#### 6.2 Test Result Information

- **Status**: ⏭️ **SKIPPED**
- **Reason**:

### Test Case 7: Target Infrastructure Summary

#### 7.1 API Request Information

- **API Endpoint**: `GET /beetle/summary/target/ns/mig01/mci/{{mciId}}?format=md`
- **Purpose**: Get a summary of the migrated target infrastructure in Markdown format
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{mciId}}` - The MCI identifier
- **Query Parameter**: `format=md`

#### 7.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Target Infrastructure Summary**:

# Target Cloud Infrastructure Summary

**Generated At:** 2026-03-25 13:31:33

**Namespace:** mig01

**MCI Name:** mig-4-mci101

---

## Overview

| Property             | Value                                                 |
| -------------------- | ----------------------------------------------------- |
| **MCI Name**         | mig-4-mci101                                          |
| **Description**      | Recommended VMs comprising multi-cloud infrastructure |
| **Status**           | Failed:3 (R:0/3)                                      |
| **Target Cloud**     | ALIBABA                                               |
| **Target Region**    |                                                       |
| **Total VMs**        | 3                                                     |
| **Running VMs**      | 0                                                     |
| **Stopped VMs**      | 0                                                     |
| **Monitoring Agent** |                                                       |

## Compute Resources

### VM Specifications

| Name               | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | VMs Using This Spec |
| ------------------ | ----- | ------------ | --- | ------------ | --------- | --------------- | ------------------- |
| ecs.t6-c1m4.xlarge | 4     | 16.0         | -   | x86_64       |           | $0.1693         | 1                   |
| ecs.t6-c1m1.large  | 2     | 2.0          | -   | x86_64       |           | $0.0214         | 1                   |
| ecs.t6-c1m4.large  | 2     | 8.0          | -   | x86_64       |           | $0.0846         | 1                   |

### VM Images

| Name                                      | Distribution        | OS Type      | OS Platform | Architecture | Root Disk Type | Root Disk Size | VMs Using This Image |
| ----------------------------------------- | ------------------- | ------------ | ----------- | ------------ | -------------- | -------------- | -------------------- |
| ubuntu_22_04_x64_20G_alibase_20260119.vhd | Ubuntu 22.04 64 bit | Ubuntu 22.04 | Linux/UNIX  | x86_64       | NA             | 20 GB          | 3                    |

### Virtual Machines

| VM Name                                         | CSP VM ID | Status | Spec (vCPU, Memory GiB) | Image                                     | Misc                                                                                                                                              |
| ----------------------------------------------- | --------- | ------ | ----------------------- | ----------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1 |           | Failed | 2 vCPU, 2.0 GiB         | Ubuntu 22.04 64 bit (Ubuntu 22.04 64 bit) | **VNet:** mig-4-vnet-01<br>**Subnet:** mig-4-subnet-01<br>**Public IP:** <br>**Private IP:** <br>**SGs:** mig-4-sg-01<br>**SSH:** mig-4-sshkey-01 |
| mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 |           | Failed | 2 vCPU, 8.0 GiB         | Ubuntu 22.04 64 bit (Ubuntu 22.04 64 bit) | **VNet:** mig-4-vnet-01<br>**Subnet:** mig-4-subnet-01<br>**Public IP:** <br>**Private IP:** <br>**SGs:** mig-4-sg-03<br>**SSH:** mig-4-sshkey-01 |
| mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 |           | Failed | 4 vCPU, 16.0 GiB        | Ubuntu 22.04 64 bit (Ubuntu 22.04 64 bit) | **VNet:** mig-4-vnet-01<br>**Subnet:** mig-4-subnet-01<br>**Public IP:** <br>**Private IP:** <br>**SGs:** mig-4-sg-02<br>**SSH:** mig-4-sshkey-01 |

## Network Resources

### Virtual Networks (VPC/VNet)

#### VNet: mig-4-vnet-01

| Property         | Value                     |
| ---------------- | ------------------------- |
| **Name**         | mig-4-vnet-01             |
| **CSP VNet ID**  | vpc-mj72oyy2s6tly01dcchqs |
| **CIDR Block**   | 10.0.0.0/21               |
| **Connection**   | alibaba-ap-northeast-2    |
| **Subnet Count** | 1                         |

**Subnets:**

| Name            | CSP Subnet ID             | CIDR Block  | Zone            |
| --------------- | ------------------------- | ----------- | --------------- |
| mig-4-subnet-01 | vsw-mj782snkv4i7r0see4na0 | 10.0.1.0/24 | ap-northeast-2a |

## Security Resources

### SSH Keys

| Name            | CSP SSH Key ID       | Username | Fingerprint                      |
| --------------- | -------------------- | -------- | -------------------------------- |
| mig-4-sshkey-01 | d71u75v693a119sjo6g0 |          | e88593506421268e6551072903407d33 |

### Security Groups

#### Security Group: mig-4-sg-01

| Property                  | Value                   |
| ------------------------- | ----------------------- |
| **Name**                  | mig-4-sg-01             |
| **CSP Security Group ID** | sg-mj7bmek530yys7apfwdi |
| **VNet**                  | mig-4-vnet-01           |
| **Rule Count**            | 14 rules                |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR        |
| --------- | -------- | ---------- | ----------- |
| inbound   | ALL      |            | 10.0.0.0/16 |
| inbound   | UDP      | 9113       | 10.0.0.0/16 |
| inbound   | TCP      | 9113       | 10.0.0.0/16 |
| inbound   | TCP      | 8080       | 0.0.0.0/0   |
| inbound   | TCP      | 443        | 0.0.0.0/0   |
| inbound   | TCP      | 80         | 0.0.0.0/0   |
| inbound   | TCP      | 22         | 0.0.0.0/0   |
| inbound   | UDP      | 1900       | 0.0.0.0/0   |
| inbound   | UDP      | 5353       | 0.0.0.0/0   |
| inbound   | UDP      | 68         | 0.0.0.0/0   |
| inbound   | ICMP     |            | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | ALL      |            | 0.0.0.0/0   |

#### Security Group: mig-4-sg-02

| Property                  | Value                   |
| ------------------------- | ----------------------- |
| **Name**                  | mig-4-sg-02             |
| **CSP Security Group ID** | sg-mj7eu0g9c9oo47ounzk6 |
| **VNet**                  | mig-4-vnet-01           |
| **Rule Count**            | 19 rules                |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR        |
| --------- | -------- | ---------- | ----------- |
| inbound   | ALL      |            | 10.0.0.0/16 |
| inbound   | UDP      | 9100       | 10.0.0.0/16 |
| inbound   | TCP      | 9100       | 10.0.0.0/16 |
| inbound   | UDP      | 32803      | 10.0.0.0/16 |
| inbound   | TCP      | 32803      | 10.0.0.0/16 |
| inbound   | UDP      | 20048      | 10.0.0.0/16 |
| inbound   | TCP      | 20048      | 10.0.0.0/16 |
| inbound   | UDP      | 111        | 0.0.0.0/0   |
| inbound   | TCP      | 111        | 0.0.0.0/0   |
| inbound   | UDP      | 2049       | 0.0.0.0/0   |
| inbound   | TCP      | 2049       | 0.0.0.0/0   |
| inbound   | TCP      | 22         | 0.0.0.0/0   |
| inbound   | UDP      | 1900       | 0.0.0.0/0   |
| inbound   | UDP      | 5353       | 0.0.0.0/0   |
| inbound   | UDP      | 68         | 0.0.0.0/0   |
| inbound   | ICMP     |            | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | ALL      |            | 0.0.0.0/0   |

#### Security Group: mig-4-sg-03

| Property                  | Value                   |
| ------------------------- | ----------------------- |
| **Name**                  | mig-4-sg-03             |
| **CSP Security Group ID** | sg-mj7i5crahq8i5rzcer4u |
| **VNet**                  | mig-4-vnet-01           |
| **Rule Count**            | 19 rules                |

**Security Group Rules:**

| Direction | Protocol | Port Range | CIDR        |
| --------- | -------- | ---------- | ----------- |
| inbound   | ALL      |            | 10.0.0.0/16 |
| inbound   | UDP      | 9104       | 10.0.0.0/16 |
| inbound   | TCP      | 9104       | 10.0.0.0/16 |
| inbound   | UDP      | 4444       | 10.0.0.0/16 |
| inbound   | TCP      | 4444       | 10.0.0.0/16 |
| inbound   | UDP      | 4568       | 10.0.0.0/16 |
| inbound   | TCP      | 4568       | 10.0.0.0/16 |
| inbound   | UDP      | 4567       | 10.0.0.0/16 |
| inbound   | TCP      | 4567       | 10.0.0.0/16 |
| inbound   | UDP      | 3306       | 10.0.0.0/16 |
| inbound   | TCP      | 3306       | 10.0.0.0/16 |
| inbound   | TCP      | 22         | 0.0.0.0/0   |
| inbound   | UDP      | 1900       | 0.0.0.0/0   |
| inbound   | UDP      | 5353       | 0.0.0.0/0   |
| inbound   | UDP      | 68         | 0.0.0.0/0   |
| inbound   | ICMP     |            | 0.0.0.0/0   |
| outbound  | UDP      | 1-65535    | 0.0.0.0/0   |
| outbound  | TCP      | 1-65535    | 0.0.0.0/0   |
| outbound  | ALL      |            | 0.0.0.0/0   |

## Cost Estimation

### Total Cost Summary

| Period                  | Cost (USD) |
| ----------------------- | ---------- |
| **Per Hour**            | $0.0000    |
| **Per Day**             | $0.00      |
| **Per Month (30 days)** | $0.00      |

### Cost by Region

| CSP     | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |
| ------- | ------ | -------- | --------------- | ---------------- |
| ALIBABA |        | 3        | $0.0000         | $0.00            |

### Cost by Virtual Machine

| VM Name                                         | Spec | Cost/Hour (USD) | Cost/Month (USD) |
| ----------------------------------------------- | ---- | --------------- | ---------------- |
| mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1 |      | $0.0000         | $0.00            |
| mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 |      | $0.0000         | $0.00            |
| mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 |      | $0.0000         | $0.00            |

### Test Case 8: Migration Report

#### 8.1 API Request Information

- **API Endpoint**: `POST /beetle/report/migration/ns/mig01/mci/{{mciId}}`
- **Purpose**: Generate a comprehensive migration report matching source to target
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{mciId}}` - The MCI identifier

#### 8.2 API Response Information

- **Status**: ✅ **SUCCESS**

**Migration Report**:

# 🚀 Infrastructure Migration Report

This report provides a comprehensive summary of the infrastructure migration from on-premise to cloud environment, including detailed information about migrated resources, costs, and configurations.

_Report generated: 2026-03-25 13:31:38_

---

## 📊 Migration Summary

**Target Cloud:** ALIBABA

**Target Region:**

**Namespace:** mig01 | **MCI ID:** mig-4-mci101

**Migration Status:** Completed

**Total Servers:** 3

**Migrated Servers:** 3

---

## 📦 Migrated Resources Overview

Summary of key infrastructure resources created or configured in the target cloud:

| #   | Resource Type       | Count             | Status          | Details                          |
| --- | ------------------- | ----------------- | --------------- | -------------------------------- |
| 1   | **Virtual Machine** | 3                 | ✅ Created      | 0 running, 3 total               |
| 2   | **VM Spec**         | 0                 | ⚠️ Not selected | No specs used                    |
| 3   | **VM OS Image**     | 1                 | ✅ Selected     | Ubuntu 22.04                     |
| 4   | **VNet (VPC)**      | 1                 | ✅ Created      | mig-4-vnet-01, CIDR: 10.0.0.0/21 |
| 5   | **Subnet**          | 1                 | ✅ Created      | 10.0.1.0/24 (in mig-4-vnet-01)   |
| 6   | **Security Group**  | 3 security groups | ✅ Created      | Total 52 rules in 3 sgs          |
| 7   | **SSH Key**         | 1 keys            | ✅ Created      | For VM access control            |

---

## 💻 Virtual Machines (VMs)

**Summary:** 3 VM(s) have been successfully created in the target cloud, migrated from 3 source server(s) in the on-premise infrastructure.

| No. | Migrated VM                                                                                                                        | Source Server                                                |
| --- | ---------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| 1   | **VM Name:** mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1<br>**VM ID:** <br>**Label(sourceMachineId):** 4-vm-ec268ed7-821e-9d73 | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec268ed7-821e-9d73 |
| 2   | **VM Name:** mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1<br>**VM ID:** <br>**Label(sourceMachineId):** 4-vm-ec288dd0-c6fa-8a49 | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec288dd0-c6fa-8a49 |
| 3   | **VM Name:** mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1<br>**VM ID:** <br>**Label(sourceMachineId):** 4-vm-ec2d32b5-98fb-5a96 | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec2d32b5-98fb-5a96 |

---

## ⚙️ VM Specs

**Summary:** 0 VM specification(s) have been selected and used for the migrated VMs.

| No. | Migrated VM                                     | VM Spec                                                                      | Source Server                                                | Source Server Spec                                                         |
| --- | ----------------------------------------------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------ | -------------------------------------------------------------------------- |
| 1   | mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Spec ID:** <br>**vCPUs:** 2<br>**Memory:** 2.0 GB<br>**Root Disk:** 50 GB  | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec268ed7-821e-9d73 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 2   | mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Spec ID:** <br>**vCPUs:** 2<br>**Memory:** 8.0 GB<br>**Root Disk:** 50 GB  | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec288dd0-c6fa-8a49 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |
| 3   | mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Spec ID:** <br>**vCPUs:** 4<br>**Memory:** 16.0 GB<br>**Root Disk:** 50 GB | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec2d32b5-98fb-5a96 | **CPUs:** N/A<br>**Threads:** N/A<br>**Memory:** N/A<br>**Root Disk:** N/A |

---

## 💿 VM OS Images

**Summary:** 1 OS image(s) have been selected and used for the migrated VMs.

| No. | Migrated VM                                     | VM OS Image Info                                                                                                                 | Source Server                                                | Source OS                                                |
| --- | ----------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ | -------------------------------------------------------- |
| 1   | mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260119.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu 22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec268ed7-821e-9d73 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 2   | mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260119.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu 22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec288dd0-c6fa-8a49 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |
| 3   | mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1 | **Image ID:** ubuntu_22_04_x64_20G_alibase_20260119.vhd<br>**OS Type:** Ubuntu 22.04<br>**OS Distribution:** Ubuntu 22.04 64 bit | **Hostname:** N/A<br>**Machine ID:** 4-vm-ec2d32b5-98fb-5a96 | **PrettyName:** N/A<br>**Name:** N/A<br>**Version:** N/A |

---

## 🔒 Security Groups

**Summary:** 3 security group(s) with 52 security rule(s) have been created and configured for the migrated VMs.

### Security Group: mig-4-sg-01

**CSP ID:** sg-mj7bmek530yys7apfwdi | **VNet:** mig-4-vnet-01 | **Rules:** 14

**Assigned VMs:**

- **VM:** mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 4-vm-ec268ed7-821e-9d73

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule              | Note                 |
| --- | --------- | -------- | ------- | ----------- | --------------------------------- | -------------------- |
| 1   | inbound   | ALL      |         | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 2   | inbound   | UDP      | 9113    | 10.0.0.0/16 | inbound udp 9113 from 10.0.0.0/16 | Migrated from source |
| 3   | inbound   | TCP      | 9113    | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 4   | inbound   | TCP      | 8080    | 0.0.0.0/0   | inbound tcp 8080                  | Migrated from source |
| 5   | inbound   | TCP      | 443     | 0.0.0.0/0   | inbound tcp 443                   | Migrated from source |
| 6   | inbound   | TCP      | 80      | 0.0.0.0/0   | inbound tcp 80                    | Migrated from source |
| 7   | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                    | Migrated from source |
| 8   | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                  | Migrated from source |
| 9   | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                  | Migrated from source |
| 10  | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                    | Migrated from source |
| 11  | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                   | Migrated from source |
| 12  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 13  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 14  | outbound  | ALL      |         | 0.0.0.0/0   | outbound \* \*                    | Migrated from source |

### Security Group: mig-4-sg-02

**CSP ID:** sg-mj7eu0g9c9oo47ounzk6 | **VNet:** mig-4-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 4-vm-ec2d32b5-98fb-5a96

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule               | Note                 |
| --- | --------- | -------- | ------- | ----------- | ---------------------------------- | -------------------- |
| 1   | inbound   | ALL      |         | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16  | Migrated from source |
| 2   | inbound   | UDP      | 9100    | 10.0.0.0/16 | inbound udp 9100 from 10.0.0.0/16  | Migrated from source |
| 3   | inbound   | TCP      | 9100    | 10.0.0.0/16 | inbound tcp 9100 from 10.0.0.0/16  | Migrated from source |
| 4   | inbound   | UDP      | 32803   | 10.0.0.0/16 | inbound udp 32803 from 10.0.0.0/16 | Migrated from source |
| 5   | inbound   | TCP      | 32803   | 10.0.0.0/16 | inbound tcp 32803 from 10.0.0.0/16 | Migrated from source |
| 6   | inbound   | UDP      | 20048   | 10.0.0.0/16 | inbound udp 20048 from 10.0.0.0/16 | Migrated from source |
| 7   | inbound   | TCP      | 20048   | 10.0.0.0/16 | inbound tcp 20048 from 10.0.0.0/16 | Migrated from source |
| 8   | inbound   | UDP      | 111     | 0.0.0.0/0   | inbound udp 111                    | Migrated from source |
| 9   | inbound   | TCP      | 111     | 0.0.0.0/0   | inbound tcp 111                    | Migrated from source |
| 10  | inbound   | UDP      | 2049    | 0.0.0.0/0   | inbound udp 2049                   | Migrated from source |
| 11  | inbound   | TCP      | 2049    | 0.0.0.0/0   | inbound tcp 2049                   | Migrated from source |
| 12  | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                     | Migrated from source |
| 13  | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                   | Migrated from source |
| 14  | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                   | Migrated from source |
| 15  | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                     | Migrated from source |
| 16  | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                    | Migrated from source |
| 17  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                  | Created by system    |
| 18  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                  | Created by system    |
| 19  | outbound  | ALL      |         | 0.0.0.0/0   | outbound \* \*                     | Migrated from source |

### Security Group: mig-4-sg-03

**CSP ID:** sg-mj7i5crahq8i5rzcer4u | **VNet:** mig-4-vnet-01 | **Rules:** 19

**Assigned VMs:**

- **VM:** mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1
  - **Source Server:** **Hostname:** N/A, **Machine ID:** 4-vm-ec288dd0-c6fa-8a49

**Security Rules:**

| No. | Direction | Protocol | Port    | CIDR        | Source Firewall Rule              | Note                 |
| --- | --------- | -------- | ------- | ----------- | --------------------------------- | -------------------- |
| 1   | inbound   | ALL      |         | 10.0.0.0/16 | inbound tcp 9113 from 10.0.0.0/16 | Migrated from source |
| 2   | inbound   | UDP      | 9104    | 10.0.0.0/16 | inbound udp 9104 from 10.0.0.0/16 | Migrated from source |
| 3   | inbound   | TCP      | 9104    | 10.0.0.0/16 | inbound tcp 9104 from 10.0.0.0/16 | Migrated from source |
| 4   | inbound   | UDP      | 4444    | 10.0.0.0/16 | inbound udp 4444 from 10.0.0.0/16 | Migrated from source |
| 5   | inbound   | TCP      | 4444    | 10.0.0.0/16 | inbound tcp 4444 from 10.0.0.0/16 | Migrated from source |
| 6   | inbound   | UDP      | 4568    | 10.0.0.0/16 | inbound udp 4568 from 10.0.0.0/16 | Migrated from source |
| 7   | inbound   | TCP      | 4568    | 10.0.0.0/16 | inbound tcp 4568 from 10.0.0.0/16 | Migrated from source |
| 8   | inbound   | UDP      | 4567    | 10.0.0.0/16 | inbound udp 4567 from 10.0.0.0/16 | Migrated from source |
| 9   | inbound   | TCP      | 4567    | 10.0.0.0/16 | inbound tcp 4567 from 10.0.0.0/16 | Migrated from source |
| 10  | inbound   | UDP      | 3306    | 10.0.0.0/16 | inbound udp 3306 from 10.0.0.0/16 | Migrated from source |
| 11  | inbound   | TCP      | 3306    | 10.0.0.0/16 | inbound tcp 3306 from 10.0.0.0/16 | Migrated from source |
| 12  | inbound   | TCP      | 22      | 0.0.0.0/0   | inbound tcp 22                    | Migrated from source |
| 13  | inbound   | UDP      | 1900    | 0.0.0.0/0   | inbound udp 1900                  | Migrated from source |
| 14  | inbound   | UDP      | 5353    | 0.0.0.0/0   | inbound udp 5353                  | Migrated from source |
| 15  | inbound   | UDP      | 68      | 0.0.0.0/0   | inbound udp 68                    | Migrated from source |
| 16  | inbound   | ICMP     |         | 0.0.0.0/0   | inbound icmp \*                   | Migrated from source |
| 17  | outbound  | UDP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 18  | outbound  | TCP      | 1-65535 | 0.0.0.0/0   | -                                 | Created by system    |
| 19  | outbound  | ALL      |         | 0.0.0.0/0   | outbound \* \*                    | Migrated from source |

---

## 🌐 VPC(VNet) and Subnets

**Summary:** Virtual Private Cloud (VPC) and subnet infrastructure have been created based on the source server network information.

### VPC(VNet)

| No. | VPC(VNet)                                                    | CIDR Block  |
| --- | ------------------------------------------------------------ | ----------- |
| 1   | **Name:** mig-4-vnet-01<br>**ID:** vpc-mj72oyy2s6tly01dcchqs | 10.0.0.0/21 |

### Subnets

| No. | Subnet                                                         | CIDR Block  | Associated VPC(VNet) |
| --- | -------------------------------------------------------------- | ----------- | -------------------- |
| 1   | **Name:** mig-4-subnet-01<br>**ID:** vsw-mj782snkv4i7r0see4na0 | 10.0.1.0/24 | mig-4-vnet-01        |

### Source Network Information

**CIDR:** 10.0.1.0/24 | **Gateway:** 10.0.1.1 | **Connected Servers:** 3

### Network Details by Server (3 servers)

#### 1. ip-10-0-1-30

**Active Interfaces:**

| Interface | IP Address  | State |
| --------- | ----------- | ----- |
| lo        | 127.0.0.1/8 | up    |

#### 2. ip-10-0-1-221

**Active Interfaces:**

| Interface | IP Address  | State |
| --------- | ----------- | ----- |
| lo        | 127.0.0.1/8 | up    |

#### 3. ip-10-0-1-138

**Active Interfaces:**

| Interface | IP Address  | State |
| --------- | ----------- | ----- |
| lo        | 127.0.0.1/8 | up    |

---

## 🔑 SSH Keys

**Summary:** 1 SSH key(s) have been created for secure access to the migrated VMs.

> **Note:** Due to security constraints and operational efficiency, it is challenging to transfer existing SSH keys from the source infrastructure. Therefore, new SSH key(s) have been generated and are commonly used across all migrated VMs. This approach ensures secure access while simplifying key management in the cloud environment.

| No. | SSH Key Name    | CSP Key ID           | Fingerprint                      | Usage             |
| --- | --------------- | -------------------- | -------------------------------- | ----------------- |
| 1   | mig-4-sshkey-01 | d71u75v693a119sjo6g0 | e88593506421268e6551072903407d33 | Used by all 3 VMs |

---

## 💰 Cost Summary

### Total Estimated Costs

| Period  | Cost (USD) |
| ------- | ---------- |
| Hourly  | $0.0000    |
| Daily   | $0.00      |
| Monthly | $0.00      |
| Yearly  | $0.00      |

### Cost Breakdown by Component

| Component                | Spec | Monthly Cost | Percentage |
| ------------------------ | ---- | ------------ | ---------- |
| ip-10-0-1-30 (migrated)  | N/A  | $0.00        | 0.0%       |
| ip-10-0-1-221 (migrated) | N/A  | $0.00        | 0.0%       |
| ip-10-0-1-138 (migrated) | N/A  | $0.00        | 0.0%       |

---

---

_Report generated by CM-Beetle_

### Test Case 9: Delete the migrated computing infra

#### 9.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}`
- **Purpose**: Delete the migrated infrastructure and clean up resources
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{mciId}}` - The MCI identifier to delete
- **Query Parameter**: `option=terminate` (terminates all resources)
- **Request Body**: None (DELETE request)

#### 9.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infrastructure deletion completed successfully

**Response Body**:

```json
{
  "message": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: mig-4-mci101)",
  "success": true
}
```
