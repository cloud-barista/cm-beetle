# CM-Beetle test results for AWS

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AWS cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.5.0+ (1655425)
- cm-model: v0.0.20
- CB-Tumblebug: v0.12.3
- CB-Spider: v0.12.11
- CB-MapUI: v0.12.16
- Target CSP: AWS
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: March 25, 2026
- Test Time: 21:39:45 KST
- Test Execution: 2026-03-25 21:39:45 KST

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

## Test result for AWS

### Test Results Summary

| Test | Step (Endpoint / Description)                          | Status      | Duration | Details |
| ---- | ------------------------------------------------------ | ----------- | -------- | ------- |
| 1    | `POST /beetle/recommendation/vmInfra`                  | ✅ **PASS** | 4.652s   | Pass    |
| 2    | `POST /beetle/migration/ns/mig01/mci`                  | ✅ **PASS** | 1m7.068s | Pass    |
| 3    | `GET /beetle/migration/ns/mig01/mci`                   | ✅ **PASS** | 225ms    | Pass    |
| 4    | `GET /beetle/migration/ns/mig01/mci?option=id`         | ✅ **PASS** | 6ms      | Pass    |
| 5    | `GET /beetle/migration/ns/mig01/mci/{{mciId}}`         | ✅ **PASS** | 269ms    | Pass    |
| 6    | Remote Command Accessibility Check                     | ✅ **PASS** | 0s       | Pass    |
| 7    | `GET /beetle/summary/target/ns/mig01/mci/{{mciId}}`    | ✅ **PASS** | 6.823s   | Pass    |
| 8    | `POST /beetle/report/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 6.405s   | Pass    |
| 9    | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}`      | ✅ **PASS** | 1m0.853s | Pass    |

**Overall Result**: 9/9 tests passed ✅

**Total Duration**: 3m30.777583817s

_Test executed on March 25, 2026 at 21:39:45 KST (2026-03-25 21:39:45 KST) using CM-Beetle automated test CLI_

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
  "nameSeed": "mig-0",
  "desiredCspAndRegionPair": {
    "csp": "aws",
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
      "nameSeed": "mig-0",
      "status": "highly-matched",
      "description": "Candidate #1 | highly-matched | Overall Match Rate: Min=100.0% Max=100.0% Avg=100.0% | VMs: 3 total, 3 matched, 0 acceptable",
      "targetCloud": {
        "csp": "aws",
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
            "name": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3a.small",
            "imageId": "ami-08a4fd517a4872931",
            "vNetId": "mig-0-vnet-01",
            "subnetId": "mig-0-subnet-01",
            "securityGroupIds": ["mig-0-sg-01"],
            "sshKeyId": "mig-0-sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3a.xlarge",
            "imageId": "ami-08a4fd517a4872931",
            "vNetId": "mig-0-vnet-01",
            "subnetId": "mig-0-subnet-01",
            "securityGroupIds": ["mig-0-sg-02"],
            "sshKeyId": "mig-0-sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3a.large",
            "imageId": "ami-08a4fd517a4872931",
            "vNetId": "mig-0-vnet-01",
            "subnetId": "mig-0-subnet-01",
            "securityGroupIds": ["mig-0-sg-03"],
            "sshKeyId": "mig-0-sshkey-01",
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
        "connectionName": "aws-ap-northeast-2",
        "cidrBlock": "10.0.0.0/21",
        "subnetInfoList": [
          {
            "name": "mig-0-subnet-01",
            "ipv4_CIDR": "10.0.1.0/24",
            "description": "a recommended subnet for migration"
          }
        ],
        "description": "a recommended vNet for migration"
      },
      "targetSshKey": {
        "name": "sshkey-01",
        "connectionName": "aws-ap-northeast-2",
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
          "id": "aws+ap-northeast-2+t3a.small",
          "uid": "d66sfgldi7idhnupe7pg",
          "cspSpecName": "t3a.small",
          "name": "aws+ap-northeast-2+t3a.small",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.0234,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "false"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:175,BaselineIops:1000,BaselineThroughputInMBps:21.875,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "false"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3a.small"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:2048}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:4,Ipv6AddressesPerInterface:4,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:2,NetworkCards:[{MaximumNetworkInterfaces:2,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
            }
          ]
        },
        {
          "id": "aws+ap-northeast-2+t3a.xlarge",
          "uid": "d66sfgldi7idhnupee1g",
          "cspSpecName": "t3a.xlarge",
          "name": "aws+ap-northeast-2+t3a.xlarge",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1872,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "false"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "false"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3a.xlarge"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:16384}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:15,Ipv6AddressesPerInterface:15,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:4,NetworkCards:[{MaximumNetworkInterfaces:4,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:2,DefaultThreadsPerCore:2,DefaultVCpus:4,ValidCores:[2],ValidThreadsPerCore:[1,2]}"
            }
          ]
        },
        {
          "id": "aws+ap-northeast-2+t3a.large",
          "uid": "d66sfgldi7idhnupebjg",
          "cspSpecName": "t3a.large",
          "name": "aws+ap-northeast-2+t3a.large",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.0936,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "false"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "false"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3a.large"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:8192}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:12,Ipv6AddressesPerInterface:12,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
            }
          ]
        }
      ],
      "targetVmOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "aws",
          "cspImageName": "ami-08a4fd517a4872931",
          "regionList": ["ap-northeast-2"],
          "id": "ami-08a4fd517a4872931",
          "uid": "d66sft5di7idhnurmtjg",
          "name": "ami-08a4fd517a4872931",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "aws-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.02.12 12:30:44 Thu",
          "creationDate": "2025-12-12T06:54:19.000Z",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212",
          "osDiskType": "ebs",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "BlockDeviceMappings",
              "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-012789365ea74b6e9,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CreationDate",
              "value": "2025-12-12T06:54:19.000Z"
            },
            {
              "key": "DeprecationTime",
              "value": "2027-12-12T06:54:19.000Z"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04, amd64 jammy image"
            },
            {
              "key": "EnaSupport",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "ImageId",
              "value": "ami-08a4fd517a4872931"
            },
            {
              "key": "ImageLocation",
              "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
            },
            {
              "key": "ImageOwnerAlias",
              "value": "amazon"
            },
            {
              "key": "ImageType",
              "value": "machine"
            },
            {
              "key": "Name",
              "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
            },
            {
              "key": "OwnerId",
              "value": "099720109477"
            },
            {
              "key": "PlatformDetails",
              "value": "Linux/UNIX"
            },
            {
              "key": "Public",
              "value": "true"
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
              "key": "SriovNetSupport",
              "value": "simple"
            },
            {
              "key": "State",
              "value": "available"
            },
            {
              "key": "UsageOperation",
              "value": "RunInstances"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04, amd64 jammy image",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-0-sg-01",
          "connectionName": "aws-ap-northeast-2",
          "vNetId": "mig-0-vnet-01",
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
          "name": "mig-0-sg-02",
          "connectionName": "aws-ap-northeast-2",
          "vNetId": "mig-0-vnet-01",
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
          "name": "mig-0-sg-03",
          "connectionName": "aws-ap-northeast-2",
          "vNetId": "mig-0-vnet-01",
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
      "nameSeed": "mig-0",
      "status": "highly-matched",
      "description": "Candidate #2 | highly-matched | Overall Match Rate: Min=100.0% Max=100.0% Avg=100.0% | VMs: 3 total, 3 matched, 0 acceptable",
      "targetCloud": {
        "csp": "aws",
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
            "name": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
            },
            "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3.small",
            "imageId": "ami-08a4fd517a4872931",
            "vNetId": "mig-0-vnet-01",
            "subnetId": "mig-0-subnet-01",
            "securityGroupIds": ["mig-0-sg-01"],
            "sshKeyId": "mig-0-sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
            },
            "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3.xlarge",
            "imageId": "ami-08a4fd517a4872931",
            "vNetId": "mig-0-vnet-01",
            "subnetId": "mig-0-subnet-01",
            "securityGroupIds": ["mig-0-sg-02"],
            "sshKeyId": "mig-0-sshkey-01",
            "rootDiskSize": 50,
            "dataDiskIds": null
          },
          {
            "name": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "subGroupSize": 1,
            "label": {
              "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
            },
            "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
            "connectionName": "aws-ap-northeast-2",
            "specId": "aws+ap-northeast-2+t3.large",
            "imageId": "ami-08a4fd517a4872931",
            "vNetId": "mig-0-vnet-01",
            "subnetId": "mig-0-subnet-01",
            "securityGroupIds": ["mig-0-sg-03"],
            "sshKeyId": "mig-0-sshkey-01",
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
        "connectionName": "aws-ap-northeast-2",
        "cidrBlock": "10.0.0.0/21",
        "subnetInfoList": [
          {
            "name": "mig-0-subnet-01",
            "ipv4_CIDR": "10.0.1.0/24",
            "description": "a recommended subnet for migration"
          }
        ],
        "description": "a recommended vNet for migration"
      },
      "targetSshKey": {
        "name": "sshkey-01",
        "connectionName": "aws-ap-northeast-2",
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
          "id": "aws+ap-northeast-2+t3a.small",
          "uid": "d66sfgldi7idhnupe7pg",
          "cspSpecName": "t3a.small",
          "name": "aws+ap-northeast-2+t3a.small",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.0234,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "false"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:175,BaselineIops:1000,BaselineThroughputInMBps:21.875,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "false"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3a.small"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:2048}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:4,Ipv6AddressesPerInterface:4,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:2,NetworkCards:[{MaximumNetworkInterfaces:2,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
            }
          ]
        },
        {
          "id": "aws+ap-northeast-2+t3a.xlarge",
          "uid": "d66sfgldi7idhnupee1g",
          "cspSpecName": "t3a.xlarge",
          "name": "aws+ap-northeast-2+t3a.xlarge",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.1872,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "false"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "false"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3a.xlarge"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:16384}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:15,Ipv6AddressesPerInterface:15,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:4,NetworkCards:[{MaximumNetworkInterfaces:4,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:2,DefaultThreadsPerCore:2,DefaultVCpus:4,ValidCores:[2],ValidThreadsPerCore:[1,2]}"
            }
          ]
        },
        {
          "id": "aws+ap-northeast-2+t3a.large",
          "uid": "d66sfgldi7idhnupebjg",
          "cspSpecName": "t3a.large",
          "name": "aws+ap-northeast-2+t3a.large",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.0936,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "false"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "false"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3a.large"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:8192}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:12,Ipv6AddressesPerInterface:12,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.2}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
            }
          ]
        },
        {
          "id": "aws+ap-northeast-2+t3.small",
          "uid": "d66sfgldi7idhnupecj0",
          "cspSpecName": "t3.small",
          "name": "aws+ap-northeast-2+t3.small",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 2,
          "diskSizeGB": -1,
          "costPerHour": 0.026,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "true"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:174,BaselineIops:1000,BaselineThroughputInMBps:21.75,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "true"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3.small"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:2048}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:4,Ipv6AddressesPerInterface:4,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.5}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
            }
          ]
        },
        {
          "id": "aws+ap-northeast-2+t3.xlarge",
          "uid": "d66sfgldi7idhnupe8v0",
          "cspSpecName": "t3.xlarge",
          "name": "aws+ap-northeast-2+t3.xlarge",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 4,
          "memoryGiB": 16,
          "diskSizeGB": -1,
          "costPerHour": 0.208,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "true"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "false"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3.xlarge"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:16384}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:15,Ipv6AddressesPerInterface:15,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:4,NetworkCards:[{MaximumNetworkInterfaces:4,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.5}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:2,DefaultThreadsPerCore:2,DefaultVCpus:4,ValidCores:[2],ValidThreadsPerCore:[1,2]}"
            }
          ]
        },
        {
          "id": "aws+ap-northeast-2+t3.large",
          "uid": "d66sfgldi7idhnupe8pg",
          "cspSpecName": "t3.large",
          "name": "aws+ap-northeast-2+t3.large",
          "namespace": "system",
          "connectionName": "aws-ap-northeast-2",
          "providerName": "aws",
          "regionName": "ap-northeast-2",
          "regionLatitude": 37.36,
          "regionLongitude": 126.78,
          "infraType": "vm",
          "architecture": "x86_64",
          "vCPU": 2,
          "memoryGiB": 8,
          "diskSizeGB": -1,
          "costPerHour": 0.104,
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
              "key": "AutoRecoverySupported",
              "value": "true"
            },
            {
              "key": "BareMetal",
              "value": "false"
            },
            {
              "key": "BurstablePerformanceSupported",
              "value": "true"
            },
            {
              "key": "CurrentGeneration",
              "value": "true"
            },
            {
              "key": "DedicatedHostsSupported",
              "value": "true"
            },
            {
              "key": "EbsInfo",
              "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:695,BaselineIops:4000,BaselineThroughputInMBps:86.875,MaximumBandwidthInMbps:2780,MaximumIops:15700,MaximumThroughputInMBps:347.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
            },
            {
              "key": "FreeTierEligible",
              "value": "false"
            },
            {
              "key": "HibernationSupported",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "nitro"
            },
            {
              "key": "InstanceStorageSupported",
              "value": "false"
            },
            {
              "key": "InstanceType",
              "value": "t3.large"
            },
            {
              "key": "MemoryInfo",
              "value": "{SizeInMiB:8192}"
            },
            {
              "key": "NetworkInfo",
              "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:12,Ipv6AddressesPerInterface:12,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
            },
            {
              "key": "PlacementGroupInfo",
              "value": "{SupportedStrategies:[partition,spread]}"
            },
            {
              "key": "ProcessorInfo",
              "value": "{SupportedArchitectures:[x86_64],SustainedClockSpeedInGhz:2.5}"
            },
            {
              "key": "SupportedBootModes",
              "value": "legacy-bios; uefi"
            },
            {
              "key": "SupportedRootDeviceTypes",
              "value": "ebs"
            },
            {
              "key": "SupportedUsageClasses",
              "value": "on-demand; spot"
            },
            {
              "key": "SupportedVirtualizationTypes",
              "value": "hvm"
            },
            {
              "key": "VCpuInfo",
              "value": "{DefaultCores:1,DefaultThreadsPerCore:2,DefaultVCpus:2,ValidCores:[1],ValidThreadsPerCore:[1,2]}"
            }
          ]
        }
      ],
      "targetVmOsImageList": [
        {
          "resourceType": "image",
          "namespace": "system",
          "providerName": "aws",
          "cspImageName": "ami-08a4fd517a4872931",
          "regionList": ["ap-northeast-2"],
          "id": "ami-08a4fd517a4872931",
          "uid": "d66sft5di7idhnurmtjg",
          "name": "ami-08a4fd517a4872931",
          "sourceVmUid": "",
          "sourceCspImageName": "",
          "connectionName": "aws-ap-northeast-2",
          "infraType": "",
          "fetchedTime": "2026.02.12 12:30:44 Thu",
          "creationDate": "2025-12-12T06:54:19.000Z",
          "isGPUImage": false,
          "isKubernetesImage": false,
          "isBasicImage": true,
          "osType": "Ubuntu 22.04",
          "osArchitecture": "x86_64",
          "osPlatform": "Linux/UNIX",
          "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212",
          "osDiskType": "ebs",
          "osDiskSizeGB": -1,
          "imageStatus": "Available",
          "details": [
            {
              "key": "Architecture",
              "value": "x86_64"
            },
            {
              "key": "BlockDeviceMappings",
              "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-012789365ea74b6e9,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
            },
            {
              "key": "BootMode",
              "value": "uefi-preferred"
            },
            {
              "key": "CreationDate",
              "value": "2025-12-12T06:54:19.000Z"
            },
            {
              "key": "DeprecationTime",
              "value": "2027-12-12T06:54:19.000Z"
            },
            {
              "key": "Description",
              "value": "Canonical, Ubuntu, 22.04, amd64 jammy image"
            },
            {
              "key": "EnaSupport",
              "value": "true"
            },
            {
              "key": "Hypervisor",
              "value": "xen"
            },
            {
              "key": "ImageId",
              "value": "ami-08a4fd517a4872931"
            },
            {
              "key": "ImageLocation",
              "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
            },
            {
              "key": "ImageOwnerAlias",
              "value": "amazon"
            },
            {
              "key": "ImageType",
              "value": "machine"
            },
            {
              "key": "Name",
              "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
            },
            {
              "key": "OwnerId",
              "value": "099720109477"
            },
            {
              "key": "PlatformDetails",
              "value": "Linux/UNIX"
            },
            {
              "key": "Public",
              "value": "true"
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
              "key": "SriovNetSupport",
              "value": "simple"
            },
            {
              "key": "State",
              "value": "available"
            },
            {
              "key": "UsageOperation",
              "value": "RunInstances"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            }
          ],
          "systemLabel": "",
          "description": "Canonical, Ubuntu, 22.04, amd64 jammy image",
          "commandHistory": null
        }
      ],
      "targetSecurityGroupList": [
        {
          "name": "mig-0-sg-01",
          "connectionName": "aws-ap-northeast-2",
          "vNetId": "mig-0-vnet-01",
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
          "name": "mig-0-sg-02",
          "connectionName": "aws-ap-northeast-2",
          "vNetId": "mig-0-vnet-01",
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
          "name": "mig-0-sg-03",
          "connectionName": "aws-ap-northeast-2",
          "vNetId": "mig-0-vnet-01",
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
  "id": "mig-0-mci101",
  "uid": "d71tfdv693a119qk81c0",
  "name": "mig-0-mci101",
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
    "sys.id": "mig-0-mci101",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig-0-mci101",
    "sys.namespace": "mig01",
    "sys.uid": "d71tfdv693a119qk81c0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "d71tfdv693a119qk81f0",
      "cspResourceName": "d71tfdv693a119qk81f0",
      "cspResourceId": "i-0782207638dbbc799",
      "name": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "subGroupId": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-03-25 12:40:51",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "3.35.236.241",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.127",
      "privateDNS": "ip-10-0-1-127.ap-northeast-2.compute.internal",
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
      "imageId": "ami-08a4fd517a4872931",
      "cspImageName": "ami-08a4fd517a4872931",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-08a4fd517a4872931",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
      },
      "vNetId": "mig-0-vnet-01",
      "cspVNetId": "vpc-05ceab61c5363f94e",
      "subnetId": "mig-0-subnet-01",
      "cspSubnetId": "subnet-0580779a990b87623",
      "networkInterface": "eni-attach-0fad4537a9a612e3a",
      "securityGroupIds": ["mig-0-sg-03"],
      "dataDiskIds": null,
      "sshKeyId": "mig-0-sshkey-01",
      "cspSshKeyId": "d71tfaf693a119qk819g",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBM4xTBvHfYDt7Fpc52KcKFuTlDGSj1qPqWp9LOAiAiyNonhEFzUb99xvjko9tBGgoRwYzWSrKKjC8uMAEoaNxNw=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:NGErlYlgKujoxb73GU3Qke3bV6Fqap7x6b0ElStKPrc",
        "firstUsedAt": "2026-03-25T12:41:03Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T12:41:01Z",
          "completedTime": "2026-03-25T12:41:04Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-127 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-03-25T12:40:31Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-00f37daa67dbe0a4a}}"
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
          "value": "A2B0B065-9069-44B4-BDBD-DE2A322040CA"
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
          "value": "ami-08a4fd517a4872931"
        },
        {
          "key": "InstanceId",
          "value": "i-0782207638dbbc799"
        },
        {
          "key": "InstanceType",
          "value": "t3a.large"
        },
        {
          "key": "KeyName",
          "value": "d71tfaf693a119qk819g"
        },
        {
          "key": "LaunchTime",
          "value": "2026-03-25T12:40:31Z"
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
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.35.236.241},Attachment:{AttachTime:2026-03-25T12:40:31Z,AttachmentId:eni-attach-0fad4537a9a612e3a,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0d4d1ba1d981b61ec,GroupName:d71tfd7693a119qk81bg}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:3a:e1:aa:cb:65,NetworkInterfaceId:eni-0faf4a794626eaf77,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.127,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.35.236.241},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.127}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0580779a990b87623,VpcId:vpc-05ceab61c5363f94e}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-127.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.127"
        },
        {
          "key": "PublicIpAddress",
          "value": "3.35.236.241"
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
          "value": "{GroupId:sg-0d4d1ba1d981b61ec,GroupName:d71tfd7693a119qk81bg}"
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
          "value": "subnet-0580779a990b87623"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d71tfdv693a119qk81f0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-05ceab61c5363f94e"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "d71tfdv693a119qk81e0",
      "cspResourceName": "d71tfdv693a119qk81e0",
      "cspResourceId": "i-0f22e32b590e94308",
      "name": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "subGroupId": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-03-25 12:40:52",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "3.36.67.247",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.216",
      "privateDNS": "ip-10-0-1-216.ap-northeast-2.compute.internal",
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
      "imageId": "ami-08a4fd517a4872931",
      "cspImageName": "ami-08a4fd517a4872931",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-08a4fd517a4872931",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
      },
      "vNetId": "mig-0-vnet-01",
      "cspVNetId": "vpc-05ceab61c5363f94e",
      "subnetId": "mig-0-subnet-01",
      "cspSubnetId": "subnet-0580779a990b87623",
      "networkInterface": "eni-attach-0a8b7507f9edad524",
      "securityGroupIds": ["mig-0-sg-02"],
      "dataDiskIds": null,
      "sshKeyId": "mig-0-sshkey-01",
      "cspSshKeyId": "d71tfaf693a119qk819g",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFEGKQwm9K85ti0yNWSpAJWGYAsnjqWj7HIHjSMS/dHLHHvxS5g43BFH1vPmtBsA2EI1d6YBG+qK8ECjJaw+xwo=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:VOP3/itINbkhczFeDqXCaoBVJoB20lv+zkwKswlv82U",
        "firstUsedAt": "2026-03-25T12:41:04Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T12:41:01Z",
          "completedTime": "2026-03-25T12:41:04Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-216 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-03-25T12:40:32Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-086d5e1a7e424921d}}"
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
          "value": "2FF9CCBB-5A7C-4E6E-A3E3-CA7325B13412"
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
          "value": "ami-08a4fd517a4872931"
        },
        {
          "key": "InstanceId",
          "value": "i-0f22e32b590e94308"
        },
        {
          "key": "InstanceType",
          "value": "t3a.xlarge"
        },
        {
          "key": "KeyName",
          "value": "d71tfaf693a119qk819g"
        },
        {
          "key": "LaunchTime",
          "value": "2026-03-25T12:40:31Z"
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
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.36.67.247},Attachment:{AttachTime:2026-03-25T12:40:31Z,AttachmentId:eni-attach-0a8b7507f9edad524,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0e6eb672171533064,GroupName:d71tfcf693a119qk81ag}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:d3:ae:91:b6:81,NetworkInterfaceId:eni-02d530a64aee7259c,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.216,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.36.67.247},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.216}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0580779a990b87623,VpcId:vpc-05ceab61c5363f94e}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-216.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.216"
        },
        {
          "key": "PublicIpAddress",
          "value": "3.36.67.247"
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
          "value": "{GroupId:sg-0e6eb672171533064,GroupName:d71tfcf693a119qk81ag}"
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
          "value": "subnet-0580779a990b87623"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d71tfdv693a119qk81e0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-05ceab61c5363f94e"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "d71tfdv693a119qk81d0",
      "cspResourceName": "d71tfdv693a119qk81d0",
      "cspResourceId": "i-0fe44e8407a4c2c47",
      "name": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "subGroupId": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-03-25 12:40:51",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "43.201.70.207",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.161",
      "privateDNS": "ip-10-0-1-161.ap-northeast-2.compute.internal",
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
      "imageId": "ami-08a4fd517a4872931",
      "cspImageName": "ami-08a4fd517a4872931",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-08a4fd517a4872931",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
      },
      "vNetId": "mig-0-vnet-01",
      "cspVNetId": "vpc-05ceab61c5363f94e",
      "subnetId": "mig-0-subnet-01",
      "cspSubnetId": "subnet-0580779a990b87623",
      "networkInterface": "eni-attach-0b000aff64a496c87",
      "securityGroupIds": ["mig-0-sg-01"],
      "dataDiskIds": null,
      "sshKeyId": "mig-0-sshkey-01",
      "cspSshKeyId": "d71tfaf693a119qk819g",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHLxvUXgCzq6Z/kqrNBoJTgVVrrv0bFFVkSM0HSK+c4Ab3ZcIwBrBKmnCHKldjBHjp4x4vvWi1JefyXPty5/L8Y=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:66QBu7klGw74bMlMh/HMbLKK4wMpDSUPUS+413ztqrc",
        "firstUsedAt": "2026-03-25T12:41:01Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T12:41:01Z",
          "completedTime": "2026-03-25T12:41:03Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-161 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-03-25T12:40:29Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0bd0e63c3dd3f5f47}}"
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
          "value": "0C440F0B-8072-41A4-A624-8E0221CEA2A1"
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
          "value": "ami-08a4fd517a4872931"
        },
        {
          "key": "InstanceId",
          "value": "i-0fe44e8407a4c2c47"
        },
        {
          "key": "InstanceType",
          "value": "t3a.small"
        },
        {
          "key": "KeyName",
          "value": "d71tfaf693a119qk819g"
        },
        {
          "key": "LaunchTime",
          "value": "2026-03-25T12:40:28Z"
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
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.201.70.207},Attachment:{AttachTime:2026-03-25T12:40:28Z,AttachmentId:eni-attach-0b000aff64a496c87,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0219bbbd3555ef292,GroupName:d71tfbn693a119qk81a0}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:9e:e3:10:42:6d,NetworkInterfaceId:eni-0e143fa98360e4ba3,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.161,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.201.70.207},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.161}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0580779a990b87623,VpcId:vpc-05ceab61c5363f94e}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-161.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.161"
        },
        {
          "key": "PublicIpAddress",
          "value": "43.201.70.207"
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
          "value": "{GroupId:sg-0219bbbd3555ef292,GroupName:d71tfbn693a119qk81a0}"
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
          "value": "subnet-0580779a990b87623"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d71tfdv693a119qk81d0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-05ceab61c5363f94e"
        }
      ]
    }
  ],
  "newVmList": null,
  "postCommand": {
    "userName": "cb-user",
    "command": ["uname -a"]
  },
  "postCommandResult": {
    "results": [
      {
        "mciId": "mig-0-mci101",
        "vmId": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vmIp": "43.201.70.207",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-161 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mig-0-mci101",
        "vmId": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "vmIp": "3.36.67.247",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-216 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mig-0-mci101",
        "vmId": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "vmIp": "3.35.236.241",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-127 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "id": "mig-0-mci101",
      "uid": "d71tfdv693a119qk81c0",
      "name": "mig-0-mci101",
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
        "sys.id": "mig-0-mci101",
        "sys.labelType": "mci",
        "sys.manager": "cb-tumblebug",
        "sys.name": "mig-0-mci101",
        "sys.namespace": "mig01",
        "sys.uid": "d71tfdv693a119qk81c0"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "d71tfdv693a119qk81d0",
          "cspResourceName": "d71tfdv693a119qk81d0",
          "cspResourceId": "i-0fe44e8407a4c2c47",
          "name": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "subGroupId": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
          "createdTime": "2026-03-25 12:40:51",
          "label": {
            "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-03-25 12:40:51",
            "sys.cspResourceId": "i-0fe44e8407a4c2c47",
            "sys.cspResourceName": "d71tfdv693a119qk81d0",
            "sys.id": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mig-0-mci101",
            "sys.name": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624",
            "sys.subnetId": "mig-0-subnet-01",
            "sys.uid": "d71tfdv693a119qk81d0",
            "sys.vNetId": "mig-0-vnet-01"
          },
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "43.201.70.207",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.161",
          "privateDNS": "ip-10-0-1-161.ap-northeast-2.compute.internal",
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
          "imageId": "ami-08a4fd517a4872931",
          "cspImageName": "ami-08a4fd517a4872931",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-08a4fd517a4872931",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
          },
          "vNetId": "mig-0-vnet-01",
          "cspVNetId": "vpc-05ceab61c5363f94e",
          "subnetId": "mig-0-subnet-01",
          "cspSubnetId": "subnet-0580779a990b87623",
          "networkInterface": "eni-attach-0b000aff64a496c87",
          "securityGroupIds": ["mig-0-sg-01"],
          "dataDiskIds": null,
          "sshKeyId": "mig-0-sshkey-01",
          "cspSshKeyId": "d71tfaf693a119qk819g",
          "vmUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHLxvUXgCzq6Z/kqrNBoJTgVVrrv0bFFVkSM0HSK+c4Ab3ZcIwBrBKmnCHKldjBHjp4x4vvWi1JefyXPty5/L8Y=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:66QBu7klGw74bMlMh/HMbLKK4wMpDSUPUS+413ztqrc",
            "firstUsedAt": "2026-03-25T12:41:01Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-03-25T12:41:01Z",
              "completedTime": "2026-03-25T12:41:03Z",
              "elapsedTime": 2,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-161 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-03-25T12:40:29Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0bd0e63c3dd3f5f47}}"
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
              "value": "0C440F0B-8072-41A4-A624-8E0221CEA2A1"
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
              "value": "ami-08a4fd517a4872931"
            },
            {
              "key": "InstanceId",
              "value": "i-0fe44e8407a4c2c47"
            },
            {
              "key": "InstanceType",
              "value": "t3a.small"
            },
            {
              "key": "KeyName",
              "value": "d71tfaf693a119qk819g"
            },
            {
              "key": "LaunchTime",
              "value": "2026-03-25T12:40:28Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.201.70.207},Attachment:{AttachTime:2026-03-25T12:40:28Z,AttachmentId:eni-attach-0b000aff64a496c87,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0219bbbd3555ef292,GroupName:d71tfbn693a119qk81a0}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:9e:e3:10:42:6d,NetworkInterfaceId:eni-0e143fa98360e4ba3,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.161,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.201.70.207},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.161}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0580779a990b87623,VpcId:vpc-05ceab61c5363f94e}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-161.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.161"
            },
            {
              "key": "PublicIpAddress",
              "value": "43.201.70.207"
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
              "value": "{GroupId:sg-0219bbbd3555ef292,GroupName:d71tfbn693a119qk81a0}"
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
              "value": "subnet-0580779a990b87623"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:d71tfdv693a119qk81d0}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-05ceab61c5363f94e"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "d71tfdv693a119qk81f0",
          "cspResourceName": "d71tfdv693a119qk81f0",
          "cspResourceId": "i-0782207638dbbc799",
          "name": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "subGroupId": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
          "createdTime": "2026-03-25 12:40:51",
          "label": {
            "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-03-25 12:40:51",
            "sys.cspResourceId": "i-0782207638dbbc799",
            "sys.cspResourceName": "d71tfdv693a119qk81f0",
            "sys.id": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mig-0-mci101",
            "sys.name": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
            "sys.subnetId": "mig-0-subnet-01",
            "sys.uid": "d71tfdv693a119qk81f0",
            "sys.vNetId": "mig-0-vnet-01"
          },
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "3.35.236.241",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.127",
          "privateDNS": "ip-10-0-1-127.ap-northeast-2.compute.internal",
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
          "imageId": "ami-08a4fd517a4872931",
          "cspImageName": "ami-08a4fd517a4872931",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-08a4fd517a4872931",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
          },
          "vNetId": "mig-0-vnet-01",
          "cspVNetId": "vpc-05ceab61c5363f94e",
          "subnetId": "mig-0-subnet-01",
          "cspSubnetId": "subnet-0580779a990b87623",
          "networkInterface": "eni-attach-0fad4537a9a612e3a",
          "securityGroupIds": ["mig-0-sg-03"],
          "dataDiskIds": null,
          "sshKeyId": "mig-0-sshkey-01",
          "cspSshKeyId": "d71tfaf693a119qk819g",
          "vmUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBM4xTBvHfYDt7Fpc52KcKFuTlDGSj1qPqWp9LOAiAiyNonhEFzUb99xvjko9tBGgoRwYzWSrKKjC8uMAEoaNxNw=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:NGErlYlgKujoxb73GU3Qke3bV6Fqap7x6b0ElStKPrc",
            "firstUsedAt": "2026-03-25T12:41:03Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-03-25T12:41:01Z",
              "completedTime": "2026-03-25T12:41:04Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-127 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-03-25T12:40:31Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-00f37daa67dbe0a4a}}"
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
              "value": "A2B0B065-9069-44B4-BDBD-DE2A322040CA"
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
              "value": "ami-08a4fd517a4872931"
            },
            {
              "key": "InstanceId",
              "value": "i-0782207638dbbc799"
            },
            {
              "key": "InstanceType",
              "value": "t3a.large"
            },
            {
              "key": "KeyName",
              "value": "d71tfaf693a119qk819g"
            },
            {
              "key": "LaunchTime",
              "value": "2026-03-25T12:40:31Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.35.236.241},Attachment:{AttachTime:2026-03-25T12:40:31Z,AttachmentId:eni-attach-0fad4537a9a612e3a,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0d4d1ba1d981b61ec,GroupName:d71tfd7693a119qk81bg}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:3a:e1:aa:cb:65,NetworkInterfaceId:eni-0faf4a794626eaf77,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.127,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.35.236.241},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.127}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0580779a990b87623,VpcId:vpc-05ceab61c5363f94e}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-127.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.127"
            },
            {
              "key": "PublicIpAddress",
              "value": "3.35.236.241"
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
              "value": "{GroupId:sg-0d4d1ba1d981b61ec,GroupName:d71tfd7693a119qk81bg}"
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
              "value": "subnet-0580779a990b87623"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:d71tfdv693a119qk81f0}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-05ceab61c5363f94e"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "d71tfdv693a119qk81e0",
          "cspResourceName": "d71tfdv693a119qk81e0",
          "cspResourceId": "i-0f22e32b590e94308",
          "name": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "subGroupId": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
          "createdTime": "2026-03-25 12:40:52",
          "label": {
            "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2026-03-25 12:40:52",
            "sys.cspResourceId": "i-0f22e32b590e94308",
            "sys.cspResourceName": "d71tfdv693a119qk81e0",
            "sys.id": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mig-0-mci101",
            "sys.name": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
            "sys.subnetId": "mig-0-subnet-01",
            "sys.uid": "d71tfdv693a119qk81e0",
            "sys.vNetId": "mig-0-vnet-01"
          },
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
          "region": {
            "region": "ap-northeast-2",
            "zone": "ap-northeast-2a"
          },
          "publicIP": "3.36.67.247",
          "sshPort": 22,
          "publicDNS": "",
          "privateIP": "10.0.1.216",
          "privateDNS": "ip-10-0-1-216.ap-northeast-2.compute.internal",
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
          "imageId": "ami-08a4fd517a4872931",
          "cspImageName": "ami-08a4fd517a4872931",
          "image": {
            "resourceType": "image",
            "cspImageName": "ami-08a4fd517a4872931",
            "osType": "Ubuntu 22.04",
            "osArchitecture": "x86_64",
            "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
          },
          "vNetId": "mig-0-vnet-01",
          "cspVNetId": "vpc-05ceab61c5363f94e",
          "subnetId": "mig-0-subnet-01",
          "cspSubnetId": "subnet-0580779a990b87623",
          "networkInterface": "eni-attach-0a8b7507f9edad524",
          "securityGroupIds": ["mig-0-sg-02"],
          "dataDiskIds": null,
          "sshKeyId": "mig-0-sshkey-01",
          "cspSshKeyId": "d71tfaf693a119qk819g",
          "vmUserName": "cb-user",
          "sshHostKeyInfo": {
            "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFEGKQwm9K85ti0yNWSpAJWGYAsnjqWj7HIHjSMS/dHLHHvxS5g43BFH1vPmtBsA2EI1d6YBG+qK8ECjJaw+xwo=",
            "keyType": "ecdsa-sha2-nistp256",
            "fingerprint": "SHA256:VOP3/itINbkhczFeDqXCaoBVJoB20lv+zkwKswlv82U",
            "firstUsedAt": "2026-03-25T12:41:04Z"
          },
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2026-03-25T12:41:01Z",
              "completedTime": "2026-03-25T12:41:04Z",
              "elapsedTime": 3,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux ip-10-0-1-216 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-03-25T12:40:32Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-086d5e1a7e424921d}}"
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
              "value": "2FF9CCBB-5A7C-4E6E-A3E3-CA7325B13412"
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
              "value": "ami-08a4fd517a4872931"
            },
            {
              "key": "InstanceId",
              "value": "i-0f22e32b590e94308"
            },
            {
              "key": "InstanceType",
              "value": "t3a.xlarge"
            },
            {
              "key": "KeyName",
              "value": "d71tfaf693a119qk819g"
            },
            {
              "key": "LaunchTime",
              "value": "2026-03-25T12:40:31Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.36.67.247},Attachment:{AttachTime:2026-03-25T12:40:31Z,AttachmentId:eni-attach-0a8b7507f9edad524,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0e6eb672171533064,GroupName:d71tfcf693a119qk81ag}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:d3:ae:91:b6:81,NetworkInterfaceId:eni-02d530a64aee7259c,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.216,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.36.67.247},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.216}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0580779a990b87623,VpcId:vpc-05ceab61c5363f94e}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-10-0-1-216.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "10.0.1.216"
            },
            {
              "key": "PublicIpAddress",
              "value": "3.36.67.247"
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
              "value": "{GroupId:sg-0e6eb672171533064,GroupName:d71tfcf693a119qk81ag}"
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
              "value": "subnet-0580779a990b87623"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:d71tfdv693a119qk81e0}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-05ceab61c5363f94e"
            }
          ]
        }
      ],
      "newVmList": null,
      "postCommand": {
        "userName": "cb-user",
        "command": ["uname -a"]
      },
      "postCommandResult": {
        "results": [
          {
            "mciId": "mig-0-mci101",
            "vmId": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
            "vmIp": "43.201.70.207",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-161 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mig-0-mci101",
            "vmId": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
            "vmIp": "3.36.67.247",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-216 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mig-0-mci101",
            "vmId": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
            "vmIp": "3.35.236.241",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux ip-10-0-1-127 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "resourceType": "mci",
      "id": "mig-4-mci101",
      "uid": "d71t5hv693a119p8fcag",
      "name": "mig-4-mci101",
      "status": "Failed:3 (R:0/3)",
      "statusCount": {
        "countTotal": 3,
        "countCreating": 0,
        "countRunning": 0,
        "countFailed": 3,
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
        "sys.id": "mig-4-mci101",
        "sys.labelType": "mci",
        "sys.manager": "cb-tumblebug",
        "sys.name": "mig-4-mci101",
        "sys.namespace": "mig01",
        "sys.uid": "d71t5hv693a119p8fcag"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "Recommended VMs comprising multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "uid": "d71t5hv693a119p8fcbg",
          "name": "mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
          "subGroupId": "mig-4-vm-ec268ed7-821e-9d73-e79f-961262161624",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Failed",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "no result with request image IID(NameId/SystemId) : ubuntu_22_04_x64_20G_alibase_20260119.vhd/ubuntu_22_04_x64_20G_alibase_20260119.vhd (from cb-spider:1024/spider/vm (500 Internal Server Error))",
          "createdTime": "",
          "label": null,
          "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": ""
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "TYPE1",
          "rootDiskSize": 50,
          "RootDeviceName": "",
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
              "zones": ["ap-northeast-2a", "ap-northeast-2b"]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "alibaba+ap-northeast-2+ecs.t6-c1m1.large",
          "cspSpecName": "",
          "spec": {},
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
          "cspImageName": "",
          "image": {
            "osType": ""
          },
          "vNetId": "mig-4-vnet-01",
          "cspVNetId": "",
          "subnetId": "mig-4-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": ["mig-4-sg-01"],
          "dataDiskIds": null,
          "sshKeyId": "mig-4-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "vm",
          "id": "mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "uid": "d71t5hv693a119p8fcdg",
          "name": "mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
          "subGroupId": "mig-4-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Failed",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "no result with request image IID(NameId/SystemId) : ubuntu_22_04_x64_20G_alibase_20260119.vhd/ubuntu_22_04_x64_20G_alibase_20260119.vhd (from cb-spider:1024/spider/vm (500 Internal Server Error))",
          "createdTime": "",
          "label": null,
          "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": ""
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "TYPE1",
          "rootDiskSize": 50,
          "RootDeviceName": "",
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
              "zones": ["ap-northeast-2a", "ap-northeast-2b"]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.large",
          "cspSpecName": "",
          "spec": {},
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
          "cspImageName": "",
          "image": {
            "osType": ""
          },
          "vNetId": "mig-4-vnet-01",
          "cspVNetId": "",
          "subnetId": "mig-4-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": ["mig-4-sg-03"],
          "dataDiskIds": null,
          "sshKeyId": "mig-4-sshkey-01",
          "cspSshKeyId": ""
        },
        {
          "resourceType": "vm",
          "id": "mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "uid": "d71t5hv693a119p8fccg",
          "name": "mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
          "subGroupId": "mig-4-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "status": "Failed",
          "targetStatus": "Running",
          "targetAction": "Create",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "no result with request image IID(NameId/SystemId) : ubuntu_22_04_x64_20G_alibase_20260119.vhd/ubuntu_22_04_x64_20G_alibase_20260119.vhd (from cb-spider:1024/spider/vm (500 Internal Server Error))",
          "createdTime": "",
          "label": null,
          "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=75.0%",
          "region": {
            "region": ""
          },
          "publicIP": "",
          "sshPort": 0,
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "TYPE1",
          "rootDiskSize": 50,
          "RootDeviceName": "",
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
              "zones": ["ap-northeast-2a", "ap-northeast-2b"]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "alibaba+ap-northeast-2+ecs.t6-c1m4.xlarge",
          "cspSpecName": "",
          "spec": {},
          "imageId": "ubuntu_22_04_x64_20G_alibase_20260119.vhd",
          "cspImageName": "",
          "image": {
            "osType": ""
          },
          "vNetId": "mig-4-vnet-01",
          "cspVNetId": "",
          "subnetId": "mig-4-subnet-01",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": ["mig-4-sg-02"],
          "dataDiskIds": null,
          "sshKeyId": "mig-4-sshkey-01",
          "cspSshKeyId": ""
        }
      ],
      "newVmList": null,
      "postCommand": {
        "userName": "cb-user",
        "command": ["uname -a"]
      },
      "postCommandResult": {
        "results": null
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
  "idList": ["mig-0-mci101", "mig-4-mci101"]
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
  "id": "mig-0-mci101",
  "uid": "d71tfdv693a119qk81c0",
  "name": "mig-0-mci101",
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
    "sys.id": "mig-0-mci101",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig-0-mci101",
    "sys.namespace": "mig01",
    "sys.uid": "d71tfdv693a119qk81c0"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "Recommended VMs comprising multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "uid": "d71tfdv693a119qk81d0",
      "cspResourceName": "d71tfdv693a119qk81d0",
      "cspResourceId": "i-0fe44e8407a4c2c47",
      "name": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
      "subGroupId": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624",
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
      "createdTime": "2026-03-25 12:40:51",
      "label": {
        "sourceMachineId": "ec268ed7-821e-9d73-e79f-961262161624",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2026-03-25 12:40:51",
        "sys.cspResourceId": "i-0fe44e8407a4c2c47",
        "sys.cspResourceName": "d71tfdv693a119qk81d0",
        "sys.id": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mig-0-mci101",
        "sys.name": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624",
        "sys.subnetId": "mig-0-subnet-01",
        "sys.uid": "d71tfdv693a119qk81d0",
        "sys.vNetId": "mig-0-vnet-01"
      },
      "description": "Recommended VM for ec268ed7-821e-9d73-e79f-961262161624 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "43.201.70.207",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.161",
      "privateDNS": "ip-10-0-1-161.ap-northeast-2.compute.internal",
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
      "imageId": "ami-08a4fd517a4872931",
      "cspImageName": "ami-08a4fd517a4872931",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-08a4fd517a4872931",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
      },
      "vNetId": "mig-0-vnet-01",
      "cspVNetId": "vpc-05ceab61c5363f94e",
      "subnetId": "mig-0-subnet-01",
      "cspSubnetId": "subnet-0580779a990b87623",
      "networkInterface": "eni-attach-0b000aff64a496c87",
      "securityGroupIds": ["mig-0-sg-01"],
      "dataDiskIds": null,
      "sshKeyId": "mig-0-sshkey-01",
      "cspSshKeyId": "d71tfaf693a119qk819g",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBHLxvUXgCzq6Z/kqrNBoJTgVVrrv0bFFVkSM0HSK+c4Ab3ZcIwBrBKmnCHKldjBHjp4x4vvWi1JefyXPty5/L8Y=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:66QBu7klGw74bMlMh/HMbLKK4wMpDSUPUS+413ztqrc",
        "firstUsedAt": "2026-03-25T12:41:01Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T12:41:01Z",
          "completedTime": "2026-03-25T12:41:03Z",
          "elapsedTime": 2,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-161 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-03-25T12:40:29Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0bd0e63c3dd3f5f47}}"
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
          "value": "0C440F0B-8072-41A4-A624-8E0221CEA2A1"
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
          "value": "ami-08a4fd517a4872931"
        },
        {
          "key": "InstanceId",
          "value": "i-0fe44e8407a4c2c47"
        },
        {
          "key": "InstanceType",
          "value": "t3a.small"
        },
        {
          "key": "KeyName",
          "value": "d71tfaf693a119qk819g"
        },
        {
          "key": "LaunchTime",
          "value": "2026-03-25T12:40:28Z"
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
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.201.70.207},Attachment:{AttachTime:2026-03-25T12:40:28Z,AttachmentId:eni-attach-0b000aff64a496c87,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0219bbbd3555ef292,GroupName:d71tfbn693a119qk81a0}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:9e:e3:10:42:6d,NetworkInterfaceId:eni-0e143fa98360e4ba3,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.161,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.201.70.207},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.161}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0580779a990b87623,VpcId:vpc-05ceab61c5363f94e}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-161.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.161"
        },
        {
          "key": "PublicIpAddress",
          "value": "43.201.70.207"
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
          "value": "{GroupId:sg-0219bbbd3555ef292,GroupName:d71tfbn693a119qk81a0}"
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
          "value": "subnet-0580779a990b87623"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d71tfdv693a119qk81d0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-05ceab61c5363f94e"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "uid": "d71tfdv693a119qk81f0",
      "cspResourceName": "d71tfdv693a119qk81f0",
      "cspResourceId": "i-0782207638dbbc799",
      "name": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
      "subGroupId": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
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
      "createdTime": "2026-03-25 12:40:51",
      "label": {
        "sourceMachineId": "ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2026-03-25 12:40:51",
        "sys.cspResourceId": "i-0782207638dbbc799",
        "sys.cspResourceName": "d71tfdv693a119qk81f0",
        "sys.id": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mig-0-mci101",
        "sys.name": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
        "sys.subnetId": "mig-0-subnet-01",
        "sys.uid": "d71tfdv693a119qk81f0",
        "sys.vNetId": "mig-0-vnet-01"
      },
      "description": "Recommended VM for ec288dd0-c6fa-8a49-2f60-bc898311febf | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "3.35.236.241",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.127",
      "privateDNS": "ip-10-0-1-127.ap-northeast-2.compute.internal",
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
      "imageId": "ami-08a4fd517a4872931",
      "cspImageName": "ami-08a4fd517a4872931",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-08a4fd517a4872931",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
      },
      "vNetId": "mig-0-vnet-01",
      "cspVNetId": "vpc-05ceab61c5363f94e",
      "subnetId": "mig-0-subnet-01",
      "cspSubnetId": "subnet-0580779a990b87623",
      "networkInterface": "eni-attach-0fad4537a9a612e3a",
      "securityGroupIds": ["mig-0-sg-03"],
      "dataDiskIds": null,
      "sshKeyId": "mig-0-sshkey-01",
      "cspSshKeyId": "d71tfaf693a119qk819g",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBM4xTBvHfYDt7Fpc52KcKFuTlDGSj1qPqWp9LOAiAiyNonhEFzUb99xvjko9tBGgoRwYzWSrKKjC8uMAEoaNxNw=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:NGErlYlgKujoxb73GU3Qke3bV6Fqap7x6b0ElStKPrc",
        "firstUsedAt": "2026-03-25T12:41:03Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T12:41:01Z",
          "completedTime": "2026-03-25T12:41:04Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-127 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-03-25T12:40:31Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-00f37daa67dbe0a4a}}"
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
          "value": "A2B0B065-9069-44B4-BDBD-DE2A322040CA"
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
          "value": "ami-08a4fd517a4872931"
        },
        {
          "key": "InstanceId",
          "value": "i-0782207638dbbc799"
        },
        {
          "key": "InstanceType",
          "value": "t3a.large"
        },
        {
          "key": "KeyName",
          "value": "d71tfaf693a119qk819g"
        },
        {
          "key": "LaunchTime",
          "value": "2026-03-25T12:40:31Z"
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
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.35.236.241},Attachment:{AttachTime:2026-03-25T12:40:31Z,AttachmentId:eni-attach-0fad4537a9a612e3a,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0d4d1ba1d981b61ec,GroupName:d71tfd7693a119qk81bg}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:3a:e1:aa:cb:65,NetworkInterfaceId:eni-0faf4a794626eaf77,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.127,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.35.236.241},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.127}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0580779a990b87623,VpcId:vpc-05ceab61c5363f94e}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-127.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.127"
        },
        {
          "key": "PublicIpAddress",
          "value": "3.35.236.241"
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
          "value": "{GroupId:sg-0d4d1ba1d981b61ec,GroupName:d71tfd7693a119qk81bg}"
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
          "value": "subnet-0580779a990b87623"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d71tfdv693a119qk81f0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-05ceab61c5363f94e"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "uid": "d71tfdv693a119qk81e0",
      "cspResourceName": "d71tfdv693a119qk81e0",
      "cspResourceId": "i-0f22e32b590e94308",
      "name": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
      "subGroupId": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
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
      "createdTime": "2026-03-25 12:40:52",
      "label": {
        "sourceMachineId": "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2026-03-25 12:40:52",
        "sys.cspResourceId": "i-0f22e32b590e94308",
        "sys.cspResourceName": "d71tfdv693a119qk81e0",
        "sys.id": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mig-0-mci101",
        "sys.name": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
        "sys.subnetId": "mig-0-subnet-01",
        "sys.uid": "d71tfdv693a119qk81e0",
        "sys.vNetId": "mig-0-vnet-01"
      },
      "description": "Recommended VM for ec2d32b5-98fb-5a96-7913-d3db1ec18932 | Match Rate: CPU=100.0% Memory=100.0% Image=100.0%",
      "region": {
        "region": "ap-northeast-2",
        "zone": "ap-northeast-2a"
      },
      "publicIP": "3.36.67.247",
      "sshPort": 22,
      "publicDNS": "",
      "privateIP": "10.0.1.216",
      "privateDNS": "ip-10-0-1-216.ap-northeast-2.compute.internal",
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
      "imageId": "ami-08a4fd517a4872931",
      "cspImageName": "ami-08a4fd517a4872931",
      "image": {
        "resourceType": "image",
        "cspImageName": "ami-08a4fd517a4872931",
        "osType": "Ubuntu 22.04",
        "osArchitecture": "x86_64",
        "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20251212"
      },
      "vNetId": "mig-0-vnet-01",
      "cspVNetId": "vpc-05ceab61c5363f94e",
      "subnetId": "mig-0-subnet-01",
      "cspSubnetId": "subnet-0580779a990b87623",
      "networkInterface": "eni-attach-0a8b7507f9edad524",
      "securityGroupIds": ["mig-0-sg-02"],
      "dataDiskIds": null,
      "sshKeyId": "mig-0-sshkey-01",
      "cspSshKeyId": "d71tfaf693a119qk819g",
      "vmUserName": "cb-user",
      "sshHostKeyInfo": {
        "hostKey": "AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBFEGKQwm9K85ti0yNWSpAJWGYAsnjqWj7HIHjSMS/dHLHHvxS5g43BFH1vPmtBsA2EI1d6YBG+qK8ECjJaw+xwo=",
        "keyType": "ecdsa-sha2-nistp256",
        "fingerprint": "SHA256:VOP3/itINbkhczFeDqXCaoBVJoB20lv+zkwKswlv82U",
        "firstUsedAt": "2026-03-25T12:41:04Z"
      },
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2026-03-25T12:41:01Z",
          "completedTime": "2026-03-25T12:41:04Z",
          "elapsedTime": 3,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux ip-10-0-1-216 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2026-03-25T12:40:32Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-086d5e1a7e424921d}}"
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
          "value": "2FF9CCBB-5A7C-4E6E-A3E3-CA7325B13412"
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
          "value": "ami-08a4fd517a4872931"
        },
        {
          "key": "InstanceId",
          "value": "i-0f22e32b590e94308"
        },
        {
          "key": "InstanceType",
          "value": "t3a.xlarge"
        },
        {
          "key": "KeyName",
          "value": "d71tfaf693a119qk819g"
        },
        {
          "key": "LaunchTime",
          "value": "2026-03-25T12:40:31Z"
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
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.36.67.247},Attachment:{AttachTime:2026-03-25T12:40:31Z,AttachmentId:eni-attach-0a8b7507f9edad524,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0e6eb672171533064,GroupName:d71tfcf693a119qk81ag}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:d3:ae:91:b6:81,NetworkInterfaceId:eni-02d530a64aee7259c,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:10.0.1.216,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:3.36.67.247},Primary:true,PrivateDnsName:null,PrivateIpAddress:10.0.1.216}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-0580779a990b87623,VpcId:vpc-05ceab61c5363f94e}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-10-0-1-216.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "10.0.1.216"
        },
        {
          "key": "PublicIpAddress",
          "value": "3.36.67.247"
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
          "value": "{GroupId:sg-0e6eb672171533064,GroupName:d71tfcf693a119qk81ag}"
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
          "value": "subnet-0580779a990b87623"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d71tfdv693a119qk81e0}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-05ceab61c5363f94e"
        }
      ]
    }
  ],
  "newVmList": null,
  "postCommand": {
    "userName": "cb-user",
    "command": ["uname -a"]
  },
  "postCommandResult": {
    "results": [
      {
        "mciId": "mig-0-mci101",
        "vmId": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1",
        "vmIp": "43.201.70.207",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-161 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mig-0-mci101",
        "vmId": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1",
        "vmIp": "3.36.67.247",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-216 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mig-0-mci101",
        "vmId": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1",
        "vmIp": "3.35.236.241",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux ip-10-0-1-127 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "output": "Linux ip-10-0-1-161 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "43.201.70.207",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624",
      "testOrder": 1,
      "userName": "cb-user",
      "vmId": "mig-0-vm-ec268ed7-821e-9d73-e79f-961262161624-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux ip-10-0-1-127 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "3.35.236.241",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf",
      "testOrder": 2,
      "userName": "cb-user",
      "vmId": "mig-0-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1"
    },
    {
      "attempts": 1,
      "command": "uname -a",
      "output": "Linux ip-10-0-1-216 6.8.0-1044-aws #46~22.04.1-Ubuntu SMP Tue Dec  2 12:52:18 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "3.36.67.247",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932",
      "testOrder": 3,
      "userName": "cb-user",
      "vmId": "mig-0-vm-ec2d32b5-98fb-5a96-7913-d3db1ec18932-1"
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
  "message": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: mig-0-mci101)",
  "success": true
}
```
