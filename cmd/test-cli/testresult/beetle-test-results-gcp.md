# CM-Beetle test results for GCP

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with GCP cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.0 (9bc6b19)
- cm-model: v0.0.14
- CB-Tumblebug: v0.11.13
- CB-Spider: v0.11.13
- CB-MapUI: v0.11.16
- Target CSP: GCP
- Target Region: asia-northeast3
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: October 10, 2025
- Test Time: 18:16:15 KST
- Test Execution: 2025-10-10 18:16:15 KST

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for GCP

### Test Results Summary

| Test | Endpoint | Status | Duration | Details |
|------|----------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/mci` | ✅ **PASS** | 712ms | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 3m29.347s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 405ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 79ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 432ms | Pass |
| 6 | Remote Command Accessibility Check | ✅ **PASS** | 0s | Pass |
| 7 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 3m47.83s | Pass |

**Overall Result**: 7/7 tests passed ✅

**Total Duration**: 7m56.54835096s

*Test executed on October 10, 2025 at 18:16:15 KST (2025-10-10 18:16:15 KST) using CM-Beetle automated test CLI*

---

## Detailed Test Case Results

> [!INFO]
> This section provides detailed information for each test case, including API request information and response details.

### Test Case 1: Recommend a target model for computing infra

#### 1.1 API Request Information

- **API Endpoint**: `POST /beetle/recommendation/mci`
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
            "ip": "192.168.110.254",
            "interfaceName": "br-ex",
            "machineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c"
          },
          {
            "ip": "192.168.110.254",
            "interfaceName": "br-ex",
            "machineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c"
          }
        ]
      },
      "ipv6Networks": {}
    },
    "servers": [
      {
        "hostname": "cm-nfs",
        "machineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 4,
          "cores": 18,
          "threads": 36,
          "maxSpeed": 3.7,
          "vendor": "GenuineIntel",
          "model": "Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 16,
          "available": 12,
          "used": 4
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1093,
          "available": 972,
          "used": 65
        },
        "dataDisks": [
          {
            "label": "unknown",
            "type": "HDD",
            "totalSize": 0
          }
        ],
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
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:52:fa:e8",
            "mtu": 1500
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:52:fa:e9",
            "mtu": 1500
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b0:03",
            "ipv4CidrBlocks": [
              "172.29.0.102/24",
              "172.29.0.200/32"
            ],
            "ipv6CidrBlocks": [
              "fe80::a6bf:1ff:fe5a:b003/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp26s0f0",
            "macAddress": "b4:96:91:53:01:54",
            "mtu": 1500
          },
          {
            "name": "enp26s0f1",
            "macAddress": "b4:96:91:53:01:55",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b0:04",
            "ipv4CidrBlocks": [
              "192.168.110.200/32"
            ],
            "ipv6CidrBlocks": [
              "fe80::a6bf:1ff:fe5a:b004/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:47:70:f0",
            "mtu": 1500
          },
          {
            "name": "enp175s0f1",
            "macAddress": "b4:96:91:47:70:f2",
            "mtu": 1500
          },
          {
            "name": "enp177s0f0",
            "macAddress": "b4:96:91:47:80:0c",
            "mtu": 1500
          },
          {
            "name": "enp177s0f1",
            "macAddress": "b4:96:91:47:80:0e",
            "mtu": 1500
          },
          {
            "name": "br-189b10762332",
            "macAddress": "02:42:32:c2:37:0e",
            "ipv4CidrBlocks": [
              "172.20.0.1/16"
            ],
            "mtu": 1500,
            "state": "down"
          },
          {
            "name": "br-f67138586d47",
            "macAddress": "02:42:6e:92:df:03",
            "ipv4CidrBlocks": [
              "172.19.0.1/16"
            ],
            "mtu": 1500,
            "state": "down"
          },
          {
            "name": "br-068801a3f047",
            "macAddress": "02:42:cc:24:25:30",
            "ipv4CidrBlocks": [
              "172.17.0.1/16"
            ],
            "ipv6CidrBlocks": [
              "fe80::42:ccff:fe24:2530/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "ovs-system",
            "macAddress": "f6:db:ff:2d:fa:8a",
            "mtu": 1500
          },
          {
            "name": "octavia-hm0",
            "macAddress": "fa:16:3e:9d:89:c5",
            "ipv4CidrBlocks": [
              "10.1.0.106/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::f816:3eff:fe9d:89c5/64"
            ],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "genev_sys_6081",
            "macAddress": "de:4b:8c:92:4c:db",
            "ipv6CidrBlocks": [
              "fe80::2852:51ff:fe36:258b/64"
            ],
            "mtu": 65000,
            "state": "up"
          },
          {
            "name": "br-int",
            "macAddress": "62:9b:45:53:d2:4f",
            "mtu": 1442
          },
          {
            "name": "br-ex",
            "macAddress": "a4:bf:01:5a:b0:04",
            "ipv4CidrBlocks": [
              "192.168.110.102/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::a6bf:1ff:fe5a:b004/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap0481d752-40",
            "macAddress": "6a:2a:78:65:42:32",
            "ipv6CidrBlocks": [
              "fe80::682a:78ff:fe65:4232/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap935cb764-41",
            "macAddress": "fe:16:3e:4c:39:2b",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe4c:392b/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap19d6d4d9-a4",
            "macAddress": "fe:16:3e:d5:6f:85",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fed5:6f85/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap7422e216-ff",
            "macAddress": "fe:16:3e:4d:31:9e",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe4d:319e/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapa53b173c-e4",
            "macAddress": "fe:16:3e:52:91:4b",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe52:914b/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapabb5f299-74",
            "macAddress": "fe:16:3e:46:9b:72",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe46:9b72/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapf6929430-67",
            "macAddress": "fe:16:3e:3e:15:10",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe3e:1510/64"
            ],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "tap3968711d-8a",
            "macAddress": "fe:16:3e:65:ad:39",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe65:ad39/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap49d44128-d0",
            "macAddress": "fe:16:3e:1e:c7:fc",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe1e:c7fc/64"
            ],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "tap708d34b6-e0",
            "macAddress": "fe:16:3e:19:8c:71",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe19:8c71/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap1479d90f-c0",
            "macAddress": "7a:0f:53:ad:50:84",
            "ipv6CidrBlocks": [
              "fe80::780f:53ff:fead:5084/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap1a03c4f4-e8",
            "macAddress": "fa:16:3e:c9:ea:1c",
            "ipv4CidrBlocks": [
              "10.254.0.27/28",
              "10.254.0.3/28"
            ],
            "ipv6CidrBlocks": [
              "fe80::f816:3eff:fec9:ea1c/64"
            ],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "veth0b8a5f4",
            "macAddress": "be:22:36:27:01:d2",
            "ipv6CidrBlocks": [
              "fe80::bc22:36ff:fe27:1d2/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "veth87e839e",
            "macAddress": "32:de:9f:d7:cd:24",
            "ipv6CidrBlocks": [
              "fe80::38f0:78ff:fef7:358/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "veth089f03a",
            "macAddress": "2a:8f:e3:66:fd:99",
            "ipv6CidrBlocks": [
              "fe80::5c87:18ff:fe73:d0dd/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapaf1a281f-c0",
            "macAddress": "32:3c:e7:79:ee:ef",
            "ipv6CidrBlocks": [
              "fe80::303c:e7ff:fe79:eeef/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap0e0c519d-d0",
            "macAddress": "fe:16:3e:8a:c2:22",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe8a:c222/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapd801f01d-d6",
            "macAddress": "fe:16:3e:09:e9:f5",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe09:e9f5/64"
            ],
            "mtu": 1442,
            "state": "up"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0/0",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "10.1.0.0/24",
            "interface": "octavia-hm0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "10.254.0.0/28",
            "interface": "tap1a03c4f4-e8",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "10.254.0.16/28",
            "interface": "tap1a03c4f4-e8",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0/16",
            "interface": "br-068801a3f047",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.19.0.0/16",
            "interface": "br-f67138586d47",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.19.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.20.0.0/16",
            "interface": "br-189b10762332",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.20.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.29.0.0/24",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0/24",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
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
            "interface": "eno2np1",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "eno1np0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "br-068801a3f047",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "genev_sys_6081",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "octavia-hm0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap0481d752-40",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap935cb764-41",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap19d6d4d9-a4",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap7422e216-ff",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tapa53b173c-e4",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tapabb5f299-74",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tapf6929430-67",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap3968711d-8a",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap49d44128-d0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap708d34b6-e0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap1479d90f-c0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap1a03c4f4-e8",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "veth0b8a5f4",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "veth87e839e",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "veth089f03a",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tapaf1a281f-c0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap0e0c519d-d0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tapd801f01d-d6",
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
            "destination": "fe80::42:ccff:fe24:2530/128",
            "interface": "br-068801a3f047",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::2852:51ff:fe36:258b/128",
            "interface": "genev_sys_6081",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::303c:e7ff:fe79:eeef/128",
            "interface": "tapaf1a281f-c0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::38f0:78ff:fef7:358/128",
            "interface": "veth87e839e",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::5c87:18ff:fe73:d0dd/128",
            "interface": "veth089f03a",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::682a:78ff:fe65:4232/128",
            "interface": "tap0481d752-40",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::780f:53ff:fead:5084/128",
            "interface": "tap1479d90f-c0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::a6bf:1ff:fe5a:b003/128",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::a6bf:1ff:fe5a:b004/128",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::a6bf:1ff:fe5a:b004/128",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::bc22:36ff:fe27:1d2/128",
            "interface": "veth0b8a5f4",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::f816:3eff:fe9d:89c5/128",
            "interface": "octavia-hm0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::f816:3eff:fec9:ea1c/128",
            "interface": "tap1a03c4f4-e8",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe09:e9f5/128",
            "interface": "tapd801f01d-d6",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe19:8c71/128",
            "interface": "tap708d34b6-e0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe1e:c7fc/128",
            "interface": "tap49d44128-d0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe3e:1510/128",
            "interface": "tapf6929430-67",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe46:9b72/128",
            "interface": "tapabb5f299-74",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe4c:392b/128",
            "interface": "tap935cb764-41",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe4d:319e/128",
            "interface": "tap7422e216-ff",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe52:914b/128",
            "interface": "tapa53b173c-e4",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe65:ad39/128",
            "interface": "tap3968711d-8a",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe8a:c222/128",
            "interface": "tap0e0c519d-d0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fed5:6f85/128",
            "interface": "tap19d6d4d9-a4",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "eno2np1",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "eno1np0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "br-068801a3f047",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "genev_sys_6081",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "octavia-hm0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap0481d752-40",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap935cb764-41",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap19d6d4d9-a4",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap7422e216-ff",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tapa53b173c-e4",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tapabb5f299-74",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tapf6929430-67",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap3968711d-8a",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap49d44128-d0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap708d34b6-e0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap1479d90f-c0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap1a03c4f4-e8",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "veth0b8a5f4",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "veth87e839e",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "veth089f03a",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tapaf1a281f-c0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap0e0c519d-d0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tapd801f01d-d6",
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
            "dstPorts": "10022",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8081,8082",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "53",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.5 LTS",
          "version": "22.04.5 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "cm-web",
        "machineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 4,
          "cores": 18,
          "threads": 36,
          "maxSpeed": 3.7,
          "vendor": "GenuineIntel",
          "model": "Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 16,
          "available": 12,
          "used": 4
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1312,
          "available": 1222,
          "used": 23
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
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:53:01:58",
            "mtu": 1500
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:53:01:59",
            "mtu": 1500
          },
          {
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:55:23:8c",
            "mtu": 1500
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b1:1b",
            "ipv4CidrBlocks": [
              "172.29.0.103/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::a6bf:1ff:fe5a:b11b/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp26s0f0",
            "macAddress": "b4:96:91:53:01:6c",
            "mtu": 1500
          },
          {
            "name": "enp26s0f1",
            "macAddress": "b4:96:91:53:01:6d",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b1:1c",
            "ipv6CidrBlocks": [
              "fe80::a6bf:1ff:fe5a:b11c/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp175s0f1",
            "macAddress": "b4:96:91:55:23:8e",
            "mtu": 1500
          },
          {
            "name": "enp177s0f0",
            "macAddress": "b4:96:91:55:1e:04",
            "mtu": 1500
          },
          {
            "name": "enp177s0f1",
            "macAddress": "b4:96:91:55:1e:06",
            "mtu": 1500
          },
          {
            "name": "ovs-system",
            "macAddress": "6e:a8:ca:69:96:82",
            "mtu": 1500
          },
          {
            "name": "br-ex",
            "macAddress": "a4:bf:01:5a:b1:1c",
            "ipv4CidrBlocks": [
              "192.168.110.103/24"
            ],
            "ipv6CidrBlocks": [
              "2001::1000/64",
              "fe80::7824:d2ff:fe2c:7330/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "genev_sys_6081",
            "macAddress": "fa:e3:ea:20:21:0c",
            "ipv6CidrBlocks": [
              "fe80::2caf:1eff:fe7f:f78f/64"
            ],
            "mtu": 65000,
            "state": "up"
          },
          {
            "name": "br-int",
            "macAddress": "ea:d0:e7:43:23:41",
            "mtu": 1442
          },
          {
            "name": "tap334a688a-76",
            "macAddress": "fe:16:3e:52:10:6e",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe52:106e/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapd4e69ee0-72",
            "macAddress": "fe:16:3e:2c:59:7a",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe2c:597a/64"
            ],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "tap50ee370b-a7",
            "macAddress": "fe:16:3e:71:22:43",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe71:2243/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tape46b2f03-d0",
            "macAddress": "fe:16:3e:e2:ea:0f",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fee2:ea0f/64"
            ],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "tap78b1ab69-36",
            "macAddress": "fe:16:3e:14:65:fb",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe14:65fb/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tape545ab95-ab",
            "macAddress": "fe:16:3e:73:24:90",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe73:2490/64"
            ],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "tapadd1bc06-e8",
            "macAddress": "fe:16:3e:26:ea:51",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe26:ea51/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapd7ff0608-f0",
            "macAddress": "fe:16:3e:2b:75:d6",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe2b:75d6/64"
            ],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "tap26b9063f-c8",
            "macAddress": "fe:16:3e:6c:c9:90",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe6c:c990/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap0d667e3a-e2",
            "macAddress": "fe:16:3e:9f:7a:65",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe9f:7a65/64"
            ],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "tap1df68fb5-f9",
            "macAddress": "fe:16:3e:13:66:6f",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe13:666f/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap3b66c516-59",
            "macAddress": "fe:16:3e:77:e4:da",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe77:e4da/64"
            ],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "tap6f653485-7f",
            "macAddress": "fe:16:3e:7d:85:5b",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fe7d:855b/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap1479d90f-c0",
            "macAddress": "02:b9:31:31:0d:fe",
            "ipv6CidrBlocks": [
              "fe80::b9:31ff:fe31:dfe/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap87daf3f9-0f",
            "macAddress": "fe:16:3e:cb:97:79",
            "ipv6CidrBlocks": [
              "fe80::fc16:3eff:fecb:9779/64"
            ],
            "mtu": 1500,
            "state": "up"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0/0",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.29.0.0/24",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0/24",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
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
            "destination": "2001::/64",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "eno2np1",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "eno1np0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "genev_sys_6081",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap334a688a-76",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tapd4e69ee0-72",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap50ee370b-a7",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tape46b2f03-d0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap78b1ab69-36",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tape545ab95-ab",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tapadd1bc06-e8",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tapd7ff0608-f0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap26b9063f-c8",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap0d667e3a-e2",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap1df68fb5-f9",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap3b66c516-59",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap6f653485-7f",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap1479d90f-c0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::/64",
            "interface": "tap87daf3f9-0f",
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
            "destination": "2001::1000/128",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::b9:31ff:fe31:dfe/128",
            "interface": "tap1479d90f-c0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::2caf:1eff:fe7f:f78f/128",
            "interface": "genev_sys_6081",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::7824:d2ff:fe2c:7330/128",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::a6bf:1ff:fe5a:b11b/128",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::a6bf:1ff:fe5a:b11c/128",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe13:666f/128",
            "interface": "tap1df68fb5-f9",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe14:65fb/128",
            "interface": "tap78b1ab69-36",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe26:ea51/128",
            "interface": "tapadd1bc06-e8",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe2b:75d6/128",
            "interface": "tapd7ff0608-f0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe2c:597a/128",
            "interface": "tapd4e69ee0-72",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe52:106e/128",
            "interface": "tap334a688a-76",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe6c:c990/128",
            "interface": "tap26b9063f-c8",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe71:2243/128",
            "interface": "tap50ee370b-a7",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe73:2490/128",
            "interface": "tape545ab95-ab",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe77:e4da/128",
            "interface": "tap3b66c516-59",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe7d:855b/128",
            "interface": "tap6f653485-7f",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fe9f:7a65/128",
            "interface": "tap0d667e3a-e2",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fecb:9779/128",
            "interface": "tap87daf3f9-0f",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "fe80::fc16:3eff:fee2:ea0f/128",
            "interface": "tape46b2f03-d0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "eno2np1",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "eno1np0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "gateway": "192.168.110.254",
            "interface": "br-ex",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "genev_sys_6081",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap334a688a-76",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tapd4e69ee0-72",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap50ee370b-a7",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tape46b2f03-d0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap78b1ab69-36",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tape545ab95-ab",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tapadd1bc06-e8",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tapd7ff0608-f0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap26b9063f-c8",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap0d667e3a-e2",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap1df68fb5-f9",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap3b66c516-59",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap6f653485-7f",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap1479d90f-c0",
            "metric": 256,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "ff00::/8",
            "interface": "tap87daf3f9-0f",
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
        "os": {
          "prettyName": "Ubuntu 22.04.5 LTS",
          "version": "22.04.5 LTS (Jammy Jellyfish)",
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
  "status": "",
  "description": "This is a list of recommended target infrastructures. Please review and use them.",
  "targetCloud": {
    "csp": "gcp",
    "region": "asia-northeast3"
  },
  "targetVmInfra": {
    "name": "mmci01",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "a recommended multi-cloud infrastructure",
    "subGroups": [
      {
        "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c"
        },
        "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "connectionName": "gcp-asia-northeast3",
        "specId": "gcp+asia-northeast3+e2-standard-4",
        "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
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
        "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c"
        },
        "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "connectionName": "gcp-asia-northeast3",
        "specId": "gcp+asia-northeast3+e2-standard-4",
        "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
        "vNetId": "mig-vnet-01",
        "subnetId": "mig-subnet-01",
        "securityGroupIds": [
          "mig-sg-02"
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
    "connectionName": "gcp-asia-northeast3",
    "cidrBlock": "192.168.96.0/19",
    "subnetInfoList": [
      {
        "name": "mig-subnet-01",
        "ipv4_CIDR": "192.168.110.0/24",
        "description": "a recommended subnet for migration"
      }
    ],
    "description": "a recommended vNet for migration"
  },
  "targetSshKey": {
    "name": "mig-sshkey-01",
    "connectionName": "gcp-asia-northeast3",
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
      "id": "gcp+asia-northeast3+e2-standard-4",
      "uid": "d3kbqjt9bnj9363dgrk0",
      "cspSpecName": "e2-standard-4",
      "name": "gcp+asia-northeast3+e2-standard-4",
      "namespace": "system",
      "connectionName": "gcp-asia-northeast3",
      "providerName": "gcp",
      "regionName": "asia-northeast3",
      "regionLatitude": 37.2,
      "regionLongitude": 127,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 15.625,
      "diskSizeGB": -1,
      "costPerHour": 0.17193145,
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
      "rootDiskSize": "-1",
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
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Zone",
          "value": "asia-northeast3-a"
        }
      ]
    }
  ],
  "targetVmOsImageList": [
    {
      "namespace": "system",
      "providerName": "gcp",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "regionList": [
        "common"
      ],
      "id": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "uid": "d3kbr3l9bnj9363fg3fg",
      "name": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "connectionName": "gcp-africa-south1",
      "infraType": "k8s|kubernetes|container",
      "fetchedTime": "2025.10.10 08:02:22 Fri",
      "isKubernetesImage": true,
      "isBasicImage": true,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "Canonical, Ubuntu, 22.04 LTS Minimal, amd64 jammy minimal image built on 2025-10-01",
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
          "value": "2979044160"
        },
        {
          "key": "CreationTimestamp",
          "value": "2025-10-01T16:52:45.602-07:00"
        },
        {
          "key": "Description",
          "value": "Canonical, Ubuntu, 22.04 LTS Minimal, amd64 jammy minimal image built on 2025-10-01"
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
          "value": "ubuntu-minimal-2204-lts"
        },
        {
          "key": "GuestOsFeatures",
          "value": "{type:VIRTIO_SCSI_MULTIQUEUE}; {type:SEV_CAPABLE}; {type:SEV_SNP_CAPABLE}; {type:SEV_LIVE_MIGRATABLE}; {type:SEV_LIVE_MIGRATABLE_V2}; {type:IDPF}; {type:TDX_CAPABLE}; {type:UEFI_COMPATIBLE}; {type:GVNIC}"
        },
        {
          "key": "Id",
          "value": "8985639578537820194"
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
          "value": "591260547102547226"
        },
        {
          "key": "Licenses",
          "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-minimal-2204-lts"
        },
        {
          "key": "Name",
          "value": "ubuntu-minimal-2204-jammy-v20251001"
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
          "value": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001"
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
          "value": "asia; eu; us"
        }
      ],
      "description": "Canonical, Ubuntu, 22.04 LTS Minimal, amd64 jammy minimal image built on 2025-10-01"
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "mig-sg-01",
      "connectionName": "gcp-asia-northeast3",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "firewallRules": [
        {
          "Ports": "10022",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "8081,8082",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "53",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        }
      ],
      "cspResourceId": ""
    },
    {
      "name": "mig-sg-02",
      "connectionName": "gcp-asia-northeast3",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "firewallRules": null,
      "cspResourceId": ""
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
  "uid": "d3kcutt9bnj9366k9b30",
  "name": "mmci01",
  "status": "Running:2 (R:2/2)",
  "statusCount": {
    "countTotal": 2,
    "countCreating": 0,
    "countRunning": 2,
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
    "sys.description": "a recommended multi-cloud infrastructure",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "d3kcutt9bnj9366k9b30"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "uid": "d3kcutt9bnj9366k9b40",
      "cspResourceName": "d3kcutt9bnj9366k9b40",
      "cspResourceId": "d3kcutt9bnj9366k9b40",
      "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
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
      "createdTime": "2025-10-10 09:19:26",
      "label": {
        "sourceMachineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c"
      },
      "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.50.53.190",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.2",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "d3kctql9bnj9366k9b0g",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "d3kctql9bnj9366k9b10",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d3kcuct9bnj9366k9b1g",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-10-10T09:19:33Z",
          "completedTime": "2025-10-10T09:19:34Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d3kcutt9bnj9366k9b40 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
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
          "value": "2025-10-10T02:18:50.738-07:00"
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
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-minimal-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d3kcutt9bnj9366k9b40,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "6mnY6zxQjys="
        },
        {
          "key": "Id",
          "value": "8409623114828575125"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "lNS-KFYjk30="
        },
        {
          "key": "Labels",
          "value": "{keypair:d3kcuct9bnj9366k9b1g}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2025-10-10T02:18:59.959-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:uhohwutlC1A=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDBANwDm8XBRCSZOdqRQ2nHeslnNTqaGxmi57F+XS4S9w+jTOfFiNT9G+EPLapnyJ0SqGYi5S8ey2QfQaxCJsiZZrPYHG8UhX1FsuIsDNGOFnywBpnbEM/HueeFDl5hi+vJHxUHrVAtk5NW09RzMfC4600zQEH4UGF9ZZ4OuuFFQdHnnleUUuG1KyQsSwgoK62KXx7NXxOR+Qputhrsy59hhMnOi5TeyF+WVPmpbXjZ45kPhGIv/6SBDeMirzjmYToVJBzrCTZzKmgPyOH8/vSYt2OvM40Q6GHssOXaLSquCflhI5+MRSaGEmZG9iLshLntAr4y07L9Pu4ci7sxGtoa4o9gbY8w29r+52W1rRAtU3vaY+HqS79TQ6W6AAzb5Mj6hH/roB1KDToxUSjTWW1GPZvYZIuZX/o/5gFaq/ZjVRTCm1vlTKLJkPzDijWvGbyUx19LYHoWnuHC15+dq1fhAeeJywsN4hSEcPoTj05zF7Q9558qC4be6lIrpj6EjQuxvuCQtYE0qceGt5dknImFQC10e/9W0Ssg6sOLekbEM2O+j0n6O3qSezIf6S6HQa3Q894DPvzVXDTZh7MRIDjeZwhs4911wyVqZae3E0voUmef4AosM0UjFyvozga/xtVnvySJ7wXmk+a0YWtzrI7IwUNnaa1JNxNoSCJQf/EmNw== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "d3kcutt9bnj9366k9b40"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.50.53.190,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:AQu7lt00r8A=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d3kctql9bnj9366k9b0g,networkIP:192.168.110.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d3kctql9bnj9366k9b10}"
        },
        {
          "key": "ResourceStatus",
          "value": "{}"
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
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d3kcutt9bnj9366k9b40"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
          "value": "{fingerprint:gzD-E3husyM=,items:[d3kcudl9bnj9366k9b20]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "uid": "d3kcutt9bnj9366k9b50",
      "cspResourceName": "d3kcutt9bnj9366k9b50",
      "cspResourceId": "d3kcutt9bnj9366k9b50",
      "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
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
      "createdTime": "2025-10-10 09:19:28",
      "label": {
        "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c"
      },
      "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.64.42.196",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.3",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "d3kctql9bnj9366k9b0g",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "d3kctql9bnj9366k9b10",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d3kcuct9bnj9366k9b1g",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-10-10T09:19:33Z",
          "completedTime": "2025-10-10T09:19:34Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d3kcutt9bnj9366k9b50 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
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
          "value": "2025-10-10T02:18:53.293-07:00"
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
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-minimal-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d3kcutt9bnj9366k9b50,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "kJptbUfWa3Y="
        },
        {
          "key": "Id",
          "value": "4418389652529312146"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "lNS-KFYjk30="
        },
        {
          "key": "Labels",
          "value": "{keypair:d3kcuct9bnj9366k9b1g}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2025-10-10T02:19:01.110-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:uhohwutlC1A=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDBANwDm8XBRCSZOdqRQ2nHeslnNTqaGxmi57F+XS4S9w+jTOfFiNT9G+EPLapnyJ0SqGYi5S8ey2QfQaxCJsiZZrPYHG8UhX1FsuIsDNGOFnywBpnbEM/HueeFDl5hi+vJHxUHrVAtk5NW09RzMfC4600zQEH4UGF9ZZ4OuuFFQdHnnleUUuG1KyQsSwgoK62KXx7NXxOR+Qputhrsy59hhMnOi5TeyF+WVPmpbXjZ45kPhGIv/6SBDeMirzjmYToVJBzrCTZzKmgPyOH8/vSYt2OvM40Q6GHssOXaLSquCflhI5+MRSaGEmZG9iLshLntAr4y07L9Pu4ci7sxGtoa4o9gbY8w29r+52W1rRAtU3vaY+HqS79TQ6W6AAzb5Mj6hH/roB1KDToxUSjTWW1GPZvYZIuZX/o/5gFaq/ZjVRTCm1vlTKLJkPzDijWvGbyUx19LYHoWnuHC15+dq1fhAeeJywsN4hSEcPoTj05zF7Q9558qC4be6lIrpj6EjQuxvuCQtYE0qceGt5dknImFQC10e/9W0Ssg6sOLekbEM2O+j0n6O3qSezIf6S6HQa3Q894DPvzVXDTZh7MRIDjeZwhs4911wyVqZae3E0voUmef4AosM0UjFyvozga/xtVnvySJ7wXmk+a0YWtzrI7IwUNnaa1JNxNoSCJQf/EmNw== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "d3kcutt9bnj9366k9b50"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.42.196,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:c7WAAfUJreA=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d3kctql9bnj9366k9b0g,networkIP:192.168.110.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d3kctql9bnj9366k9b10}"
        },
        {
          "key": "ResourceStatus",
          "value": "{}"
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
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d3kcutt9bnj9366k9b50"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
          "value": "{fingerprint:wB1lgrY7TGI=,items:[d3kcuot9bnj9366k9b2g]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
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
        "vmId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "vmIp": "34.50.53.190",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d3kcutt9bnj9366k9b40 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "vmIp": "34.64.42.196",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d3kcutt9bnj9366k9b50 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
      "uid": "d3kcutt9bnj9366k9b30",
      "name": "mmci01",
      "status": "Running:2 (R:2/2)",
      "statusCount": {
        "countTotal": 2,
        "countCreating": 0,
        "countRunning": 2,
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
        "sys.description": "a recommended multi-cloud infrastructure",
        "sys.id": "mmci01",
        "sys.labelType": "mci",
        "sys.manager": "cb-tumblebug",
        "sys.name": "mmci01",
        "sys.namespace": "mig01",
        "sys.uid": "d3kcutt9bnj9366k9b30"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "a recommended multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
          "uid": "d3kcutt9bnj9366k9b50",
          "cspResourceName": "d3kcutt9bnj9366k9b50",
          "cspResourceId": "d3kcutt9bnj9366k9b50",
          "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
          "subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
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
          "createdTime": "2025-10-10 09:19:28",
          "label": {
            "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2025-10-10 09:19:28",
            "sys.cspResourceId": "d3kcutt9bnj9366k9b50",
            "sys.cspResourceName": "d3kcutt9bnj9366k9b50",
            "sys.id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
            "sys.uid": "d3kcutt9bnj9366k9b50"
          },
          "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
          "region": {
            "Region": "asia-northeast3",
            "Zone": "asia-northeast3-a"
          },
          "publicIP": "34.64.42.196",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "192.168.110.3",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": "50",
          "rootDiskName": "",
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
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "d3kctql9bnj9366k9b0g",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "d3kctql9bnj9366k9b10",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d3kcuct9bnj9366k9b1g",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-10-10T09:19:33Z",
              "completedTime": "2025-10-10T09:19:34Z",
              "elapsedTime": 1,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d3kcutt9bnj9366k9b50 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
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
              "value": "2025-10-10T02:18:53.293-07:00"
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
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-minimal-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d3kcutt9bnj9366k9b50,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "kJptbUfWa3Y="
            },
            {
              "key": "Id",
              "value": "4418389652529312146"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "lNS-KFYjk30="
            },
            {
              "key": "Labels",
              "value": "{keypair:d3kcuct9bnj9366k9b1g}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2025-10-10T02:19:01.110-07:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:uhohwutlC1A=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDBANwDm8XBRCSZOdqRQ2nHeslnNTqaGxmi57F+XS4S9w+jTOfFiNT9G+EPLapnyJ0SqGYi5S8ey2QfQaxCJsiZZrPYHG8UhX1FsuIsDNGOFnywBpnbEM/HueeFDl5hi+vJHxUHrVAtk5NW09RzMfC4600zQEH4UGF9ZZ4OuuFFQdHnnleUUuG1KyQsSwgoK62KXx7NXxOR+Qputhrsy59hhMnOi5TeyF+WVPmpbXjZ45kPhGIv/6SBDeMirzjmYToVJBzrCTZzKmgPyOH8/vSYt2OvM40Q6GHssOXaLSquCflhI5+MRSaGEmZG9iLshLntAr4y07L9Pu4ci7sxGtoa4o9gbY8w29r+52W1rRAtU3vaY+HqS79TQ6W6AAzb5Mj6hH/roB1KDToxUSjTWW1GPZvYZIuZX/o/5gFaq/ZjVRTCm1vlTKLJkPzDijWvGbyUx19LYHoWnuHC15+dq1fhAeeJywsN4hSEcPoTj05zF7Q9558qC4be6lIrpj6EjQuxvuCQtYE0qceGt5dknImFQC10e/9W0Ssg6sOLekbEM2O+j0n6O3qSezIf6S6HQa3Q894DPvzVXDTZh7MRIDjeZwhs4911wyVqZae3E0voUmef4AosM0UjFyvozga/xtVnvySJ7wXmk+a0YWtzrI7IwUNnaa1JNxNoSCJQf/EmNw== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "d3kcutt9bnj9366k9b50"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.42.196,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:c7WAAfUJreA=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d3kctql9bnj9366k9b0g,networkIP:192.168.110.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d3kctql9bnj9366k9b10}"
            },
            {
              "key": "ResourceStatus",
              "value": "{}"
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
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d3kcutt9bnj9366k9b50"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
              "value": "{fingerprint:wB1lgrY7TGI=,items:[d3kcuot9bnj9366k9b2g]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
          "uid": "d3kcutt9bnj9366k9b40",
          "cspResourceName": "d3kcutt9bnj9366k9b40",
          "cspResourceId": "d3kcutt9bnj9366k9b40",
          "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
          "subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
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
          "createdTime": "2025-10-10 09:19:26",
          "label": {
            "sourceMachineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c",
            "sys.connectionName": "gcp-asia-northeast3",
            "sys.createdTime": "2025-10-10 09:19:26",
            "sys.cspResourceId": "d3kcutt9bnj9366k9b40",
            "sys.cspResourceName": "d3kcutt9bnj9366k9b40",
            "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
            "sys.uid": "d3kcutt9bnj9366k9b40"
          },
          "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
          "region": {
            "Region": "asia-northeast3",
            "Zone": "asia-northeast3-a"
          },
          "publicIP": "34.50.53.190",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "192.168.110.2",
          "privateDNS": "",
          "rootDiskType": "pd-standard",
          "rootDiskSize": "50",
          "rootDiskName": "",
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
          "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
          "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "d3kctql9bnj9366k9b0g",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "d3kctql9bnj9366k9b10",
          "networkInterface": "nic0",
          "securityGroupIds": [
            "mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d3kcuct9bnj9366k9b1g",
          "vmUserName": "cb-user",
          "commandStatus": [
            {
              "index": 1,
              "commandRequested": "uname -a",
              "commandExecuted": "uname -a",
              "status": "Completed",
              "startedTime": "2025-10-10T09:19:33Z",
              "completedTime": "2025-10-10T09:19:34Z",
              "elapsedTime": 1,
              "resultSummary": "Command executed successfully",
              "stdout": "Linux d3kcutt9bnj9366k9b40 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
              "stderr": "\n"
            }
          ],
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
              "value": "2025-10-10T02:18:50.738-07:00"
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
              "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-minimal-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d3kcutt9bnj9366k9b40,type:PERSISTENT}"
            },
            {
              "key": "Fingerprint",
              "value": "6mnY6zxQjys="
            },
            {
              "key": "Id",
              "value": "8409623114828575125"
            },
            {
              "key": "Kind",
              "value": "compute#instance"
            },
            {
              "key": "LabelFingerprint",
              "value": "lNS-KFYjk30="
            },
            {
              "key": "Labels",
              "value": "{keypair:d3kcuct9bnj9366k9b1g}"
            },
            {
              "key": "LastStartTimestamp",
              "value": "2025-10-10T02:18:59.959-07:00"
            },
            {
              "key": "MachineType",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
            },
            {
              "key": "Metadata",
              "value": "{fingerprint:uhohwutlC1A=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDBANwDm8XBRCSZOdqRQ2nHeslnNTqaGxmi57F+XS4S9w+jTOfFiNT9G+EPLapnyJ0SqGYi5S8ey2QfQaxCJsiZZrPYHG8UhX1FsuIsDNGOFnywBpnbEM/HueeFDl5hi+vJHxUHrVAtk5NW09RzMfC4600zQEH4UGF9ZZ4OuuFFQdHnnleUUuG1KyQsSwgoK62KXx7NXxOR+Qputhrsy59hhMnOi5TeyF+WVPmpbXjZ45kPhGIv/6SBDeMirzjmYToVJBzrCTZzKmgPyOH8/vSYt2OvM40Q6GHssOXaLSquCflhI5+MRSaGEmZG9iLshLntAr4y07L9Pu4ci7sxGtoa4o9gbY8w29r+52W1rRAtU3vaY+HqS79TQ6W6AAzb5Mj6hH/roB1KDToxUSjTWW1GPZvYZIuZX/o/5gFaq/ZjVRTCm1vlTKLJkPzDijWvGbyUx19LYHoWnuHC15+dq1fhAeeJywsN4hSEcPoTj05zF7Q9558qC4be6lIrpj6EjQuxvuCQtYE0qceGt5dknImFQC10e/9W0Ssg6sOLekbEM2O+j0n6O3qSezIf6S6HQa3Q894DPvzVXDTZh7MRIDjeZwhs4911wyVqZae3E0voUmef4AosM0UjFyvozga/xtVnvySJ7wXmk+a0YWtzrI7IwUNnaa1JNxNoSCJQf/EmNw== cb-user}],kind:compute#metadata}"
            },
            {
              "key": "Name",
              "value": "d3kcutt9bnj9366k9b40"
            },
            {
              "key": "NetworkInterfaces",
              "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.50.53.190,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:AQu7lt00r8A=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d3kctql9bnj9366k9b0g,networkIP:192.168.110.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d3kctql9bnj9366k9b10}"
            },
            {
              "key": "ResourceStatus",
              "value": "{}"
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
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d3kcutt9bnj9366k9b40"
            },
            {
              "key": "ServiceAccounts",
              "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
              "value": "{fingerprint:gzD-E3husyM=,items:[d3kcudl9bnj9366k9b20]}"
            },
            {
              "key": "Zone",
              "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
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
            "vmId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
            "vmIp": "34.50.53.190",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d3kcutt9bnj9366k9b40 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
            },
            "stderr": {
              "0": ""
            },
            "err": null
          },
          {
            "mciId": "mmci01",
            "vmId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
            "vmIp": "34.64.42.196",
            "command": {
              "0": "uname -a"
            },
            "stdout": {
              "0": "Linux d3kcutt9bnj9366k9b50 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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
  "uid": "d3kcutt9bnj9366k9b30",
  "name": "mmci01",
  "status": "Running:2 (R:2/2)",
  "statusCount": {
    "countTotal": 2,
    "countCreating": 0,
    "countRunning": 2,
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
    "sys.description": "a recommended multi-cloud infrastructure",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "d3kcutt9bnj9366k9b30"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "uid": "d3kcutt9bnj9366k9b50",
      "cspResourceName": "d3kcutt9bnj9366k9b50",
      "cspResourceId": "d3kcutt9bnj9366k9b50",
      "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
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
      "createdTime": "2025-10-10 09:19:28",
      "label": {
        "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2025-10-10 09:19:28",
        "sys.cspResourceId": "d3kcutt9bnj9366k9b50",
        "sys.cspResourceName": "d3kcutt9bnj9366k9b50",
        "sys.id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "sys.uid": "d3kcutt9bnj9366k9b50"
      },
      "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.64.42.196",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.3",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "d3kctql9bnj9366k9b0g",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "d3kctql9bnj9366k9b10",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d3kcuct9bnj9366k9b1g",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-10-10T09:19:33Z",
          "completedTime": "2025-10-10T09:19:34Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d3kcutt9bnj9366k9b50 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
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
          "value": "2025-10-10T02:18:53.293-07:00"
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
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-minimal-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d3kcutt9bnj9366k9b50,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "kJptbUfWa3Y="
        },
        {
          "key": "Id",
          "value": "4418389652529312146"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "lNS-KFYjk30="
        },
        {
          "key": "Labels",
          "value": "{keypair:d3kcuct9bnj9366k9b1g}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2025-10-10T02:19:01.110-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:uhohwutlC1A=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDBANwDm8XBRCSZOdqRQ2nHeslnNTqaGxmi57F+XS4S9w+jTOfFiNT9G+EPLapnyJ0SqGYi5S8ey2QfQaxCJsiZZrPYHG8UhX1FsuIsDNGOFnywBpnbEM/HueeFDl5hi+vJHxUHrVAtk5NW09RzMfC4600zQEH4UGF9ZZ4OuuFFQdHnnleUUuG1KyQsSwgoK62KXx7NXxOR+Qputhrsy59hhMnOi5TeyF+WVPmpbXjZ45kPhGIv/6SBDeMirzjmYToVJBzrCTZzKmgPyOH8/vSYt2OvM40Q6GHssOXaLSquCflhI5+MRSaGEmZG9iLshLntAr4y07L9Pu4ci7sxGtoa4o9gbY8w29r+52W1rRAtU3vaY+HqS79TQ6W6AAzb5Mj6hH/roB1KDToxUSjTWW1GPZvYZIuZX/o/5gFaq/ZjVRTCm1vlTKLJkPzDijWvGbyUx19LYHoWnuHC15+dq1fhAeeJywsN4hSEcPoTj05zF7Q9558qC4be6lIrpj6EjQuxvuCQtYE0qceGt5dknImFQC10e/9W0Ssg6sOLekbEM2O+j0n6O3qSezIf6S6HQa3Q894DPvzVXDTZh7MRIDjeZwhs4911wyVqZae3E0voUmef4AosM0UjFyvozga/xtVnvySJ7wXmk+a0YWtzrI7IwUNnaa1JNxNoSCJQf/EmNw== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "d3kcutt9bnj9366k9b50"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.64.42.196,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:c7WAAfUJreA=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d3kctql9bnj9366k9b0g,networkIP:192.168.110.3,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d3kctql9bnj9366k9b10}"
        },
        {
          "key": "ResourceStatus",
          "value": "{}"
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
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d3kcutt9bnj9366k9b50"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
          "value": "{fingerprint:wB1lgrY7TGI=,items:[d3kcuot9bnj9366k9b2g]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "uid": "d3kcutt9bnj9366k9b40",
      "cspResourceName": "d3kcutt9bnj9366k9b40",
      "cspResourceId": "d3kcutt9bnj9366k9b40",
      "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
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
      "createdTime": "2025-10-10 09:19:26",
      "label": {
        "sourceMachineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2025-10-10 09:19:26",
        "sys.cspResourceId": "d3kcutt9bnj9366k9b40",
        "sys.cspResourceName": "d3kcutt9bnj9366k9b40",
        "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.uid": "d3kcutt9bnj9366k9b40"
      },
      "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.50.53.190",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.2",
      "privateDNS": "",
      "rootDiskType": "pd-standard",
      "rootDiskSize": "50",
      "rootDiskName": "",
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
      "imageId": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20251001",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "d3kctql9bnj9366k9b0g",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "d3kctql9bnj9366k9b10",
      "networkInterface": "nic0",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d3kcuct9bnj9366k9b1g",
      "vmUserName": "cb-user",
      "commandStatus": [
        {
          "index": 1,
          "commandRequested": "uname -a",
          "commandExecuted": "uname -a",
          "status": "Completed",
          "startedTime": "2025-10-10T09:19:33Z",
          "completedTime": "2025-10-10T09:19:34Z",
          "elapsedTime": 1,
          "resultSummary": "Command executed successfully",
          "stdout": "Linux d3kcutt9bnj9366k9b40 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n\n",
          "stderr": "\n"
        }
      ],
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
          "value": "2025-10-10T02:18:50.738-07:00"
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
          "value": "{architecture:X86_64,autoDelete:true,boot:true,deviceName:persistent-disk-0,diskSizeGb:50,guestOsFeatures:[{type:VIRTIO_SCSI_MULTIQUEUE},{type:SEV_CAPABLE},{type:SEV_SNP_CAPABLE},{type:SEV_LIVE_MIGRATABLE},{type:SEV_LIVE_MIGRATABLE_V2},{type:IDPF},{type:TDX_CAPABLE},{type:UEFI_COMPATIBLE},{type:GVNIC}],interface:SCSI,kind:compute#attachedDisk,licenses:[https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/licenses/ubuntu-minimal-2204-lts],mode:READ_WRITE,shieldedInstanceInitialState:{dbxs:[{content:2gcDBhMRFQAAAAAAAAAAABENAAAAAvEOndKvSt9o7kmKqTR9N1ZlpzCCDPUCAQExDzANBglghkgBZQMEAgEFADALBgkqhkiG9w0BBwGgggsIMIIFGDCCBACgAwIBAgITMwAAABNryScg3e1ZiAAAAAAAEzANBgkqhkiG9w0BAQsFADCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMB4XDTE2MDEwNjE4MzQxNVoXDTE3MDQwNjE4MzQxNVowgZUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdSZWRtb25kMR4wHAYDVQQKExVNaWNyb3NvZnQgQ29ycG9yYXRpb24xDTALBgNVBAsTBE1PUFIxMDAuBgNVBAMTJ01pY3Jvc29mdCBXaW5kb3dzIFVFRkkgS2V5IEV4Y2hhbmdlIEtleTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKXiCkZgbboTnVZnS1h_JbnlcVst9wtFK8NQjTpeB9wirml3h-fzi8vzki0hSNBD2Dg49lGEvs4egyowmTsLu1TnBUH1f_Hi8Noa7fKXV6F93qYrTPajx5v9L7NedplWnMEPsRvJrQdrysTZwtoXMLYDhc8bQHI5nlJDfgqrB8JiC4A3vL9i19lkQOTq4PZb5AcVcE0wlG7lR_btoQN0g5B4_7pI2S_9mU1PXr1NBSEl48Kl4cJwO2GyvOVvxQ6wUSFTExmCBKrT3LnPU5lZY68n3MpZ5VY4skhrEt2dyf5bZNzkYTTouxC0n37OrMbGGq3tpv7JDD6E_Rfqua3dXYECAwEAAaOCAXIwggFuMBQGA1UdJQQNMAsGCSsGAQQBgjdPATAdBgNVHQ4EFgQUVsJIppTfox2XYoAJRIlnxAUOy2owUQYDVR0RBEowSKRGMEQxDTALBgNVBAsTBE1PUFIxMzAxBgNVBAUTKjMxNjMxKzJjNDU2Y2JjLTA1NDItNDdkOS05OWU1LWQzOWI4MTVjNTczZTAfBgNVHSMEGDAWgBRi_EPNoD6ky2cS0lvZVax7zLaKXzBTBgNVHR8ETDBKMEigRqBEhkJodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NybC9NaWNDb3JLRUtDQTIwMTFfMjAxMS0wNi0yNC5jcmwwYAYIKwYBBQUHAQEEVDBSMFAGCCsGAQUFBzAChkRodHRwOi8vd3d3Lm1pY3Jvc29mdC5jb20vcGtpb3BzL2NlcnRzL01pY0NvcktFS0NBMjAxMV8yMDExLTA2LTI0LmNydDAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQCGjTFLjxsKmyLESJueg0S2Cp8N7MOq2IALsitZHwfYw2jMhY9b9kmKvIdSqVna1moZ6_zJSOS_JY6HkWZr6dDJe9Lj7xiW_e4qPP-KDrCVb02vBnK4EktVjTdJpyMhxBMdXUcq1eGl6518oCkQ27tu0-WZjaWEVsEY_gpQj0ye2UA4HYUYgJlpT24oJRi7TeQ03Nebb-ZrUkbf9uxl0OVV_mg2R5FDwOc3REoRAgv5jnw6X7ha5hlRCl2cLF27TFrFIRQQT4eSM33eDiitXXpYmD13jqKeHhLVXr07QSwqvKe1o1UYokJngP0pTwoDnt2qRuLnZ71jw732dSPN9B57MIIF6DCCA9CgAwIBAgIKYQrRiAAAAAAAAzANBgkqhkiG9w0BAQsFADCBkTELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjE7MDkGA1UEAxMyTWljcm9zb2Z0IENvcnBvcmF0aW9uIFRoaXJkIFBhcnR5IE1hcmtldHBsYWNlIFJvb3QwHhcNMTEwNjI0MjA0MTI5WhcNMjYwNjI0MjA1MTI5WjCBgDELMAkGA1UEBhMCVVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1JlZG1vbmQxHjAcBgNVBAoTFU1pY3Jvc29mdCBDb3Jwb3JhdGlvbjEqMCgGA1UEAxMhTWljcm9zb2Z0IENvcnBvcmF0aW9uIEtFSyBDQSAyMDExMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxOi1ir-tVyawJsPq5_tXekQCXQcN2krldCrmsA_sbevsf7njWmMyfBEXTw7jC6c4FZOOxvXghLGamyzn9beR1gnh4sAEqKwwHN9I8wZQmmSnUX_IhU-PIIbO_i_hn_-CwO3pzc70U2piOgtDueIl_f4F-dTEFKsR4iOJjXC3pB1N7K7lnPoWwtfBy9ToxC_lme4kiwPsjfKL6sNK-0MREgt-tUeSbNzmBInr9TME6xABKnHl-YMTPP8lCS9odkb_uk--3K1xKliq-w7SeT3km2U7zCkqn_xyWaLrrpLv9jUTgMYC7ORfzJ12ze9jksGveUCEeYd_41Ko6J17B2mPFQIDAQABo4IBTzCCAUswEAYJKwYBBAGCNxUBBAMCAQAwHQYDVR0OBBYEFGL8Q82gPqTLZxLSW9lVrHvMtopfMBkGCSsGAQQBgjcUAgQMHgoAUwB1AGIAQwBBMAsGA1UdDwQEAwIBhjAPBgNVHRMBAf8EBTADAQH_MB8GA1UdIwQYMBaAFEVmUkPhflgRv9ZOniNVCDs6ImqoMFwGA1UdHwRVMFMwUaBPoE2GS2h0dHA6Ly9jcmwubWljcm9zb2Z0LmNvbS9wa2kvY3JsL3Byb2R1Y3RzL01pY0NvclRoaVBhck1hclJvb18yMDEwLTEwLTA1LmNybDBgBggrBgEFBQcBAQRUMFIwUAYIKwYBBQUHMAKGRGh0dHA6Ly93d3cubWljcm9zb2Z0LmNvbS9wa2kvY2VydHMvTWljQ29yVGhpUGFyTWFyUm9vXzIwMTAtMTAtMDUuY3J0MA0GCSqGSIb3DQEBCwUAA4ICAQDUhIj1FJQYAsoqPPsqkhwM16DR8ehSZqjuorV1epAAqi2kdlrqebe5N2pRexBk9uFk8gJnvveoG3i9us6IWGQM1lfIGaNfBdbbxtBpzkhLMrfrXdIw9cD1uLp4B6Mr_pvbNFaE7ILKrkElcJxr6f6QD9eWH-XnlB-yKgyNS_8oKRB799d8pdF2uQXIee0PkJKcwv7fb35sD3vUwUXdNFGWOQ_lXlbYGAWW9AemQrOgd_0IGfJxVsyfhiOkh8um_Vh-1GlnFZF-gfJ_E-UNi4o8h4Tr4869Q-WtLYSTjmorWnxE-lKqgcgtHLvgUt8AEfiaPcFgsOEztaOI0WUZChrnrHykwYKHTjixLw3FFIdv_Y0uvDm25-bD4OTNJ4TvlELvKYuQRkE7gRtn2PlDWWXLDbz9AJJP9HU7p6kk_FBBQHngLU8Kaid2blLtlml7rw_3hwXQRcKtUxSBH_swBKo3NmHaSmkbNNho7dYCz2yUDNPPbCJ5rbHwvAOiRmCpxAfCIYLx_fLoeTJgv9ispSIUS8rB2EvrfT9XNbLmT3W0sGADIlOukXkd1ptBHxWGVHCy3g01D3ywNHK6l2A78HnrorIcXaIWuIfF6Rv2tZclbzif45H6inmYw2kOt6McIAWX-MoUrgDXxPPAFBB1azSgG7WZYPNcsMVXTjbSMoS_njGCAcQwggHAAgEBMIGYMIGAMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEeMBwGA1UEChMVTWljcm9zb2Z0IENvcnBvcmF0aW9uMSowKAYDVQQDEyFNaWNyb3NvZnQgQ29ycG9yYXRpb24gS0VLIENBIDIwMTECEzMAAAATa8knIN3tWYgAAAAAABMwDQYJYIZIAWUDBAIBBQAwDQYJKoZIhvcNAQEBBQAEggEAhabaxRIJ7nUZ-m__mIG0lII6yD-lxoeI8S83ZKTP8Qx5h5asySWl7420eGhna7zyaVRvVVIhkjOMIfcKr29LgzQpYDqPUc8aYAdGCsZKZGmHCMjEulnq5TDK79GKinzZfb2sAWXEJ68N8oNnY7faBKjHjmmJbAEz8ufE4DijgJ_NBov2xmhTZyNHQ7pB1iCdrEUGObzdJc0Qtmh3CNOEcmH0ukd8sTHE9acBBTFHS8dvreR_sP7dXClZJbJiWAFKvQn3EjCTiYizkZ4I_5xiqjHELht_ORQKN-Hnoqnl4kcRINhZRV7JlgAQDlBJLv3OTjShRO_ZWCdcu7PtwhweiSYWxMFMUJJArKlB-TaTQyiMDgAAAAAAADAAAAC9mvp3WQMyTb1gKPTnj3hLgLTZaTG_DQL9kaYeGdFPHaRS5m2yQIyoYE1BH5Jlnwq9mvp3WQMyTb1gKPTnj3hL9S-Do_qc-9aSD3IoJNvkA0U00luFByRrO5V9rG4bznq9mvp3WQMyTb1gKPTnj3hLxdnYoYbiyC0Jr6oqb38uc4cNPmT3LE4I72d5aoQPD729mvp3WQMyTb1gKPTnj3hLNjOE0U0fLgt4FWJkhMRZrVejGO9DliZgSNBYxaGbv3a9mvp3WQMyTb1gKPTnj3hLGuyEuEtsZaUSIKm-cYGWUjAhDWLW0zxImZxrKVorCga9mvp3WQMyTb1gKPTnj3hL5spo6UFGYprwP2nC-G5r72L5MLN8b7zIeLeN-YwDNOW9mvp3WQMyTb1gKPTnj3hLw6maRg2kZKBXw1htg8719K4ItxA5ee2JMnQt8O1TDGa9mvp3WQMyTb1gKPTnj3hLWPuUGu-VollDs_tfJRCg3z_kTFjJXgq4BIcpdWirl3G9mvp3WQMyTb1gKPTnj3hLU5HDovsRIQKmqh7cJa534Z9dbwnNCe6yUJkiv81Zkuq9mvp3WQMyTb1gKPTnj3hL1iYVfh1qcYvBJKuNony7ZQcsoDp7ayV9vcu9YPZe89G9mvp3WQMyTb1gKPTnj3hL0GPsKPZ-ulPxZC2_ff8zxqMq3YafYBP-Fi4sMvHL5W29mvp3WQMyTb1gKPTnj3hLKcbrUrQ8OqGLLNjtbqhgfO88-uG6_hFldVzy5hSESkS9mvp3WQMyTb1gKPTnj3hLkPvnDmnWM0CNPhcMaDLbstIJ4CclJ9-2PUnSlXKm9Ey9mvp3WQMyTb1gKPTnj3hLB17qBgWJVIugYLL-7RDaPCDH_psXzQJrlOimg7gRUji9mvp3WQMyTb1gKPTnj3hLB-bGqFhkb7HvxnkD_iixFgEfI2f-kua-KzaZnv850J69mvp3WQMyTb1gKPTnj3hLCd9fTlESCOx4uW0S0IEl_bYDho3jn29yknhSWZtlnCa9mvp3WQMyTb1gKPTnj3hLC7tDktqseribMKSsZXUxuXv6qwT5Cw2v5fm265CgY3S9mvp3WQMyTb1gKPTnj3hLDBiTOXYt8zarPdAGpGPfcVo5z7D0kkZcYA5sa9e9iYy9mvp3WQMyTb1gKPTnj3hLDQ2-ym8p7KBvMxp9cuSISxIJf7NImDoqFKDXP08QFA-9mvp3WQMyTb1gKPTnj3hLDcnz-5mWIUjDyoM2MnWNPtT8jQsAB7lbMeZSjyrNW_y9mvp3WQMyTb1gKPTnj3hLEG-s6s_s_U4wO3T0gKCAmOLQgCuTb47HdM4h8xaGaJy9mvp3WQMyTb1gKPTnj3hLF046C1tDxqYHu9NATwU0Hj3POWJnzpT4tQ4uI6nakgy9mvp3WQMyTb1gKPTnj3hLGDM0Kf8FYu2flwM-EUjc7uUtvi5JbVQQtc_WyGTS0Q-9mvp3WQMyTb1gKPTnj3hLK5nPJkIukv42X79Lww0nCGye4Ut6b_9E-y9rkAFpmTm9mvp3WQMyTb1gKPTnj3hLK78sp7jx2R8n7lK2-ypd0Em4WiubUpxdZmIGgQSwVfi9mvp3WQMyTb1gKPTnj3hLLHPZMyW6bcvlidSkxjxbk1VZ75L78FDtUMTiCFIG8X29mvp3WQMyTb1gKPTnj3hLLnCRZ4am93NRH6cYH6sPHXC1V8YyLqkjsqjTuStRr329mvp3WQMyTb1gKPTnj3hLMGYo-lR3MFcoukpGfefQOHpU9WnTdp_OXnXsidKNFZO9mvp3WQMyTb1gKPTnj3hLNgjtuvWtD0GkFKF3er8vr15nAzRnXsOZXmk1gp4MqtK9mvp3WQMyTb1gKPTnj3hLOEHSITaNFYPXXAoC5iFgOU1sTgpnYLb2B7kDYryFWwK9mvp3WQMyTb1gKPTnj3hLP86bn98-8J1UUrD5XuSBwrfwbXQ6c3lxVY5wE2rOPnO9mvp3WQMyTb1gKPTnj3hLQ5fayoOef2MHfLUMkt9DvC0vsqj1nyb8eg5L1Nl1FpK9mvp3WQMyTb1gKPTnj3hLR8wIYSfiBpqG4Dpr7yzUEPjFWm1r2zYhaMMbLOMqWt-9mvp3WQMyTb1gKPTnj3hLUYgx_nOCtRTQPhXGISKLirZUeb0Mv6PFwdD0jZwwYTW9mvp3WQMyTb1gKPTnj3hLWulJ6ohV65PkOdvGW9ouQoUsL99nifoUZzbjw0EPK1y9mvp3WQMyTb1gKPTnj3hLax0TgHjkQYqmjet7s14GYJLPR57rjOTNEufQcsy0L2a9mvp3WQMyTb1gKPTnj3hLbIhUR43VWeKTUbgmwGy4v-8rlK01ODWHctGT-C7RyhG9mvp3WQMyTb1gKPTnj3hLbxQo_3HJ2w7Vrx8ue7_Lq2R8wmXd9bKTzbYm9Qo6eF69mvp3WQMyTb1gKPTnj3hLcfKQb9IiSX5Uo0ZiqySX_MgQIHcP9RNo6ePZv8v9Y3W9mvp3WQMyTb1gKPTnj3hLcms-tlQEajDz-D2bls4D9nDpqAbRcIoDceYtxJ0sI8G9mvp3WQMyTb1gKPTnj3hLcuC9GGfPXZ1WqxWK3zvdvIK_MqjYqh2MXi9t8pQo1ti9mvp3WQMyTb1gKPTnj3hLeCevmTYs-vBxfa3ksb_gQ4rRccFa3cJIt1v4yqRLssW9mvp3WQMyTb1gKPTnj3hLgai5ZbuE04drlCmpVIHMlVMYz6oUEtgIyKM7_TP_8OS9mvp3WQMyTb1gKPTnj3hLgts7zrT2CEPOnZfD0YfNm1lBzT3oEA5YbyvaVjdXX2e9mvp3WQMyTb1gKPTnj3hLiVqXhfYXyh1-1E_BoUcLcfPxIjhi2f-dzDri35IWPa-9mvp3WQMyTb1gKPTnj3hLitZIWfGVtfWNr6qUC2phZ6zWeohuj0aTZBdyIcVZRbm9mvp3WQMyTb1gKPTnj3hLi_Q0tJ4AzPcVAqLNkAhlywHsOz2gPDW-UF_fe9Vj9SG9mvp3WQMyTb1gKPTnj3hLjY6iic_nChwHq3NlyyjuUe3TPPJQbeiI-63WDr-ASBy9mvp3WQMyTb1gKPTnj3hLmZjTY8SRvha9dLoQuU2SkQAWEXNv3KZDo2ZkvA8xWkK9mvp3WQMyTb1gKPTnj3hLnkppFzFhaC5V_ej-9WDriOwf_tyvBAAfZsDK9weytzS9mvp3WQMyTb1gKPTnj3hLprUVHzZV06KvDUcnWXlr5KQgDlSVp9hpdUxISIV0CKe9mvp3WQMyTb1gKPTnj3hLp_MvUI1OsP6tmgh--U7RugrsXeb372_wpiuTvt9dRY29mvp3WQMyTb1gKPTnj3hLrWgm4ZRtJtPq82hciNl9hd47Tcs9DuKugccFYNE8VyC9mvp3WQMyTb1gKPTnj3hLruuuMVEnEnPtlaouZxE57TGphWcwOjMimPg3CanVWqG9mvp3WQMyTb1gKPTnj3hLr-IDCvt9LNoT-fozOgLjT2dRr-wRsBDbzUQf30xAArO9mvp3WQMyTb1gKPTnj3hLtU8e5jZjH61oBY07CTcDGsG5DMsXBio5HMpor9vkDVW9mvp3WQMyTb1gKPTnj3hLuPB42YOiSsQzIWOTiDUUzZMsM68Y591wiEyCNfQnVza9mvp3WQMyTb1gKPTnj3hLuXoIiQWcA1_x1UtttTsRuXZmaNn5VSR8AosoN9egTNm9mvp3WQMyTb1gKPTnj3hLvIemaOgZZkictQjugFGDwZ5qzSTPF3mcoGLS44TaDqe9mvp3WQMyTb1gKPTnj3hLxAm9rEd1rdjbkqoitbcY-4yUoUYsH-mkFrldijOIwvy9mvp3WQMyTb1gKPTnj3hLxhfBqLHuKoEcKLWoG0yD18mLWwwnKB1hAgfr5pLCln-9mvp3WQMyTb1gKPTnj3hLyQ8zZhe45_mDl1QTyZfxC3PrJn_YoQy5472_xmer24u9mvp3WQMyTb1gKPTnj3hLy2uFi0DToJh2WBW1ksFRSklgT6_WCBnaiNenbpd4_ve9mvp3WQMyTb1gKPTnj3hLzjv6vlnWfOisjf1KFvfEPvnCJFE_vGVZV9c1-in1QM69mvp3WQMyTb1gKPTnj3hL2MvrlzX1Zys2fk-WzcdJaWFdFwdK6WxyTULOAhb48_q9mvp3WQMyTb1gKPTnj3hL6Swi6ztWQtZcHsLK8kfSWUc47rt_s4QaRJVvWeKw0fq9mvp3WQMyTb1gKPTnj3hL_d1uPSnqhMd0Pa1KG9vHALX-wbOR-TJAkIasxx3W29i9mvp3WQMyTb1gKPTnj3hL_mOoT3gsydP88sz5_BH70Ddgh4dY0mKF7RJmm9xubQG9mvp3WQMyTb1gKPTnj3hL_s-yMtEumUttSF0scWdyiqVSWYStXKYedRYiHweaFDa9mvp3WQMyTb1gKPTnj3hLyhcdYUqNfhIck5SM0P5V05mB-dEaqW4DRQpBUifCxlu9mvp3WQMyTb1gKPTnj3hLVbmbDeU9vP5IWqnHN88_thbvPZH6tZmqfKsZ7adjtbq9mvp3WQMyTb1gKPTnj3hLd90ZD6MNiP9eOwEaCuYeYgl4DBMLU17Lh-bwiIoLay-9mvp3WQMyTb1gKPTnj3hLyDyxOSKtmfVgdEZ13TfMlNytWh_Lpkcv7jQRcdk56IS9mvp3WQMyTb1gKPTnj3hLOwKHUz4Mw9DsGqgjy_CpQarYchV50cSZgC3Rw6Y2uKm9mvp3WQMyTb1gKPTnj3hLk5ru9PX6UeIzQMPy5JBIzohyUmr991LDp_Oj8ryfYEm9mvp3WQMyTb1gKPTnj3hLZFdb2RJ4mi4UrVb2NB9Sr2v4DPlEAHhZdenwTi1k10W9mvp3WQMyTb1gKPTnj3hLRcfIrnUKz7tI_DdSfWQS3WRNrtiRPM2KJMlNhWln344=,fileType:BIN}]},source:https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/d3kcutt9bnj9366k9b40,type:PERSISTENT}"
        },
        {
          "key": "Fingerprint",
          "value": "6mnY6zxQjys="
        },
        {
          "key": "Id",
          "value": "8409623114828575125"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "LabelFingerprint",
          "value": "lNS-KFYjk30="
        },
        {
          "key": "Labels",
          "value": "{keypair:d3kcuct9bnj9366k9b1g}"
        },
        {
          "key": "LastStartTimestamp",
          "value": "2025-10-10T02:18:59.959-07:00"
        },
        {
          "key": "MachineType",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/machineTypes/e2-standard-4"
        },
        {
          "key": "Metadata",
          "value": "{fingerprint:uhohwutlC1A=,items:[{key:ssh-keys,value:cb-user:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDBANwDm8XBRCSZOdqRQ2nHeslnNTqaGxmi57F+XS4S9w+jTOfFiNT9G+EPLapnyJ0SqGYi5S8ey2QfQaxCJsiZZrPYHG8UhX1FsuIsDNGOFnywBpnbEM/HueeFDl5hi+vJHxUHrVAtk5NW09RzMfC4600zQEH4UGF9ZZ4OuuFFQdHnnleUUuG1KyQsSwgoK62KXx7NXxOR+Qputhrsy59hhMnOi5TeyF+WVPmpbXjZ45kPhGIv/6SBDeMirzjmYToVJBzrCTZzKmgPyOH8/vSYt2OvM40Q6GHssOXaLSquCflhI5+MRSaGEmZG9iLshLntAr4y07L9Pu4ci7sxGtoa4o9gbY8w29r+52W1rRAtU3vaY+HqS79TQ6W6AAzb5Mj6hH/roB1KDToxUSjTWW1GPZvYZIuZX/o/5gFaq/ZjVRTCm1vlTKLJkPzDijWvGbyUx19LYHoWnuHC15+dq1fhAeeJywsN4hSEcPoTj05zF7Q9558qC4be6lIrpj6EjQuxvuCQtYE0qceGt5dknImFQC10e/9W0Ssg6sOLekbEM2O+j0n6O3qSezIf6S6HQa3Q894DPvzVXDTZh7MRIDjeZwhs4911wyVqZae3E0voUmef4AosM0UjFyvozga/xtVnvySJ7wXmk+a0YWtzrI7IwUNnaa1JNxNoSCJQf/EmNw== cb-user}],kind:compute#metadata}"
        },
        {
          "key": "Name",
          "value": "d3kcutt9bnj9366k9b40"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{accessConfigs:[{kind:compute#accessConfig,name:External NAT,natIP:34.50.53.190,networkTier:PREMIUM,type:ONE_TO_ONE_NAT}],fingerprint:AQu7lt00r8A=,kind:compute#networkInterface,name:nic0,network:https://www.googleapis.com/compute/v1/projects/ykkim-etri/global/networks/d3kctql9bnj9366k9b0g,networkIP:192.168.110.2,stackType:IPV4_ONLY,subnetwork:https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/d3kctql9bnj9366k9b10}"
        },
        {
          "key": "ResourceStatus",
          "value": "{}"
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
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/instances/d3kcutt9bnj9366k9b40"
        },
        {
          "key": "ServiceAccounts",
          "value": "{email:kimy-service-account@ykkim-etri.iam.gserviceaccount.com,scopes:[https://www.googleapis.com/auth/devstorage.full_control,https://www.googleapis.com/auth/compute]}"
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
          "value": "{fingerprint:gzD-E3husyM=,items:[d3kcudl9bnj9366k9b20]}"
        },
        {
          "key": "Zone",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
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
        "vmId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "vmIp": "34.50.53.190",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d3kcutt9bnj9366k9b40 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
        },
        "stderr": {
          "0": ""
        },
        "err": null
      },
      {
        "mciId": "mmci01",
        "vmId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "vmIp": "34.64.42.196",
        "command": {
          "0": "uname -a"
        },
        "stdout": {
          "0": "Linux d3kcutt9bnj9366k9b50 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux\n"
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

**Summary**: 2/2 VMs accessible via SSH

**Test Statistics**:

- Total VMs: 2
- Successful Tests: 2
- Failed Tests: 0

**Complete Test Details**:

<details>
  <summary> <ins>Click to see detailed test results </ins> </summary>

```json
{
  "failedTests": 0,
  "overallStatus": {
    "message": "2/2 VMs accessible via SSH",
    "success": true
  },
  "successfulTests": 2,
  "totalVMs": 2,
  "vmResults": [
    {
      "command": "uname -a",
      "output": "Linux d3kcutt9bnj9366k9b50 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "34.64.42.196",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "testOrder": 1,
      "userName": "cb-user",
      "vmId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1"
    },
    {
      "command": "uname -a",
      "output": "Linux d3kcutt9bnj9366k9b40 6.8.0-1041-gcp #43~22.04.1-Ubuntu SMP Wed Sep 24 23:11:19 UTC 2025 x86_64 x86_64 x86_64 GNU/Linux",
      "publicIP": "34.50.53.190",
      "sshTest": "successful",
      "status": "success",
      "subGroup": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "testOrder": 2,
      "userName": "cb-user",
      "vmId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1"
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

