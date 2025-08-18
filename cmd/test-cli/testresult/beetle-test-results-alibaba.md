# CM-Beetle test results for ALIBABA

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with ALIBABA cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.0 (8261244)
- cm-model: v0.0.11
- CB-Tumblebug: v0.11.3
- CB-Spider: v0.11.1
- CB-MapUI: v0.11.4
- Target CSP: ALIBABA
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: August 18, 2025
- Test Time: 19:24:28 KST
- Test Execution: 2025-08-18 19:24:28 KST

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for ALIBABA

### Test Results Summary

| Test | Endpoint | Status | Duration | Details |
|------|----------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/mci` | ✅ **PASS** | 257ms | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 1m0.903s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 109ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 64ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 115ms | Pass |
| 6 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 45.73s | Pass |

**Overall Result**: 6/6 tests passed ✅

**Total Duration**: 2m17.435222369s

*Test executed on August 18, 2025 at 19:24:28 KST (2025-08-18 19:24:28 KST) using CM-Beetle automated test CLI*

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
    "csp": "alibaba",
    "region": "ap-northeast-2"
  },
  "OnpremiseInfraModel": {
    "network": {
      "ipv4Networks": {
        "defaultGateways": [
          {
            "ip": "192.168.110.254",
            "interfaceName": "br-ex",
            "machineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c"
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
          "cpus": 2,
          "cores": 18,
          "threads": 36,
          "maxSpeed": 3.7,
          "vendor": "GenuineIntel",
          "model": "Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 8,
          "available": 3,
          "used": 5
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
          },
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
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8086",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8888",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9201",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9202",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9203",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9204",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9206",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "3100",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "3000",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8443",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9000",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9001",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "18080",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "13000",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9101",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9100",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9106",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9105",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8080",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9102",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9103",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "9104",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "5672",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "1883",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "4369",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "15672",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "15675",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "25672",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8883",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "16567",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "192.168.110.0/24",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "8000",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "0.0.0.0/0",
            "srcPorts": "*",
            "dstCIDR": "0.0.0.0/0",
            "dstPorts": "*",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
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
            "dstPorts": "*",
            "protocol": "tcp",
            "direction": "inbound",
            "action": "allow"
          },
          {
            "srcCIDR": "::/0",
            "srcPorts": "*",
            "dstCIDR": "::/0",
            "dstPorts": "*",
            "protocol": "udp",
            "direction": "inbound",
            "action": "allow"
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
  "targetVmInfra": {
    "name": "mmci01",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "a recommended multi-cloud infrastructure",
    "vm": [
      {
        "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "subGroupSize": "",
        "label": null,
        "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "connectionName": "alibaba-ap-northeast-2",
        "specId": "ecs.e-c1m4.large",
        "imageId": "ubuntu_22_04_x64_20G_alibase_20250722.vhd",
        "vNetId": "mig-vnet-01",
        "subnetId": "mig-subnet-01",
        "securityGroupIds": [
          "mig-sg-01"
        ],
        "sshKeyId": "mig-sshkey-01",
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
    "connectionName": "alibaba-ap-northeast-2",
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
      "id": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
      "cspSpecName": "ecs.e-c1m4.large",
      "name": "alibaba+ap-northeast-2+ecs.e-c1m4.large",
      "namespace": "system",
      "connectionName": "alibaba-ap-northeast-2",
      "providerName": "alibaba",
      "regionName": "ap-northeast-2",
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 8,
      "diskSizeGB": -1,
      "costPerHour": 0.0791,
      "orderInFilteredResult": 1,
      "evaluationScore01": -1,
      "evaluationScore02": -1,
      "evaluationScore03": -1,
      "evaluationScore04": -1,
      "evaluationScore05": -1,
      "evaluationScore06": -1,
      "evaluationScore07": -1,
      "evaluationScore08": -1,
      "evaluationScore09": 1.0000001,
      "evaluationScore10": -1,
      "rootDiskType": "",
      "rootDiskSize": "-1",
      "systemLabel": "auto-gen",
      "details": [
        {
          "key": "MemorySize",
          "value": "8.00"
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
          "value": "ecs.e-c1m4.large"
        },
        {
          "key": "InstanceBandwidthRx",
          "value": "409600"
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
          "value": "ecs.e"
        },
        {
          "key": "InitialCredit",
          "value": "0"
        },
        {
          "key": "InstancePpsTx",
          "value": "200000"
        },
        {
          "key": "InstanceFamilyLevel",
          "value": "EntryLevel"
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
          "value": "409600"
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
          "value": "0.00"
        },
        {
          "key": "BaselineCredit",
          "value": "0"
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
          "value": "Intel(R) Xeon(R) Platinum"
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
          "value": "{Core:1,HyperThreadingAdjustable:true,CoreCount:0,CoreFactor:1,Numa:,TopologyType:,ThreadsPerCore:2,SupportedTopologyTypes:{SupportedTopologyType:null}}"
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
      "namespace": "system",
      "providerName": "alibaba",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20250722.vhd",
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
        "me-east-1",
        "na-south-1",
        "us-east-1",
        "us-west-1"
      ],
      "id": "alibaba+ubuntu_22_04_x64_20g_alibase_20250722.vhd",
      "name": "alibaba+ubuntu_22_04_x64_20g_alibase_20250722.vhd",
      "connectionName": "alibaba-ap-northeast-1",
      "fetchedTime": "2025.08.12 07:58:40 Tue",
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
          "value": "ubuntu_22_04_x64_20G_alibase_20250722.vhd"
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
          "value": "Kernel version is 5.15.0-144-generic, 2025.7.31"
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
          "value": "v2025.7.31"
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
          "value": "2025-07-31T09:12:38Z"
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
          "value": "ubuntu_22_04_x64_20G_alibase_20250722.vhd"
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
      "description": "Kernel version is 5.15.0-144-generic, 2025.7.31"
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "mig-sg-01",
      "connectionName": "alibaba-ap-northeast-2",
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
        },
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
          "Ports": "8086",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "8888",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9201",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9202",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9203",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9204",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9206",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "3100",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "3000",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "8443",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9000",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9001",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "18080",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "13000",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9101",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9100",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9106",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9105",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "8080",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9102",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9103",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "9104",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "5672",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "1883",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "4369",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "15672",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "15675",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "25672",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "8883",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "16567",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "8000",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "192.168.110.0/24"
        },
        {
          "Ports": "1-65535",
          "Protocol": "tcp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "1-65535",
          "Protocol": "udp",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
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
  "uid": "d2hfus4djuna1t2ovmj0",
  "name": "mmci01",
  "status": "Running:1 (R:1/1)",
  "statusCount": {
    "countTotal": 1,
    "countCreating": 0,
    "countRunning": 1,
    "countFailed": 0,
    "countSuspended": 0,
    "countRebooting": 0,
    "countTerminated": 0,
    "countSuspending": 0,
    "countResuming": 0,
    "countTerminating": 0,
    "countUndefined": 0
  },
  "targetStatus": "Running",
  "targetAction": "Create",
  "installMonAgent": "",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "a recommended multi-cloud infrastructure",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "d2hfus4djuna1t2ovmj0"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "uid": "d2hfus4djuna1t2ovmk0",
      "cspResourceName": "d2hfus4djuna1t2ovmk0",
      "cspResourceId": "i-mj7cxi74lacgci8l7cl9",
      "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
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
      "createdTime": "2025-08-18 10:25:33",
      "label": {
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2025-08-18 10:25:33",
        "sys.cspResourceId": "i-mj7cxi74lacgci8l7cl9",
        "sys.cspResourceName": "d2hfus4djuna1t2ovmk0",
        "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.uid": "d2hfus4djuna1t2ovmk0"
      },
      "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "8.220.214.219",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.20",
      "privateDNS": "",
      "rootDiskType": "cloud_essd",
      "rootDiskSize": "40",
      "rootDiskName": "",
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
      "specId": "ecs.e-c1m4.large",
      "cspSpecName": "ecs.e-c1m4.large",
      "imageId": "ubuntu_22_04_x64_20G_alibase_20250722.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20250722.vhd",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-mj7rn5qgal81ddfdf1fq4",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "vsw-mj703pc1l4txdkwjfirxc",
      "networkInterface": "eni-mj7cxi74lacgci8lewrb",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d2hfuqcdjuna1t2ovmi0",
      "vmUserName": "cb-user",
      "vmUserPassword": "$dd!uj4fs2h1uA",
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20250722.vhd"
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
          "value": "d2hfus4djuna1t2ovmk0"
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
          "value": "2025-08-18T10:25Z"
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
          "value": "iZmj7cxi74lacgci8l7cl9Z"
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
          "value": "32bffe3d-7bcc-4156-a53f-9821406ed1e0"
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
          "value": "i-mj7cxi74lacgci8l7cl9"
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
          "value": "2025-08-18T10:25Z"
        },
        {
          "key": "KeyPairName",
          "value": "d2hfuqcdjuna1t2ovmi0"
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
          "value": "{SecurityGroupId:[sg-mj7huz3swfzzhrto8waq]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.220.214.219]}"
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
          "value": "{VSwitchId:vsw-mj703pc1l4txdkwjfirxc,VpcId:vpc-mj7rn5qgal81ddfdf1fq4,NatIpAddress:,PrivateIpAddress:{IpAddress:[192.168.110.20]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:01:b6:d5,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7cxi74lacgci8lewrb,PrimaryIpAddress:192.168.110.20,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:192.168.110.20,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    }
  ],
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

- **Status**: ✅ **SUCCESS**
- **Response**: MCI list retrieved successfully

**Response Body**:

```json
{
  "mci": [
    {
      "resourceType": "mci",
      "id": "mmci01",
      "uid": "d2hfus4djuna1t2ovmj0",
      "name": "mmci01",
      "status": "Running:1 (R:1/1)",
      "statusCount": {
        "countTotal": 1,
        "countCreating": 0,
        "countRunning": 1,
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
      "label": null,
      "systemLabel": "",
      "systemMessage": "",
      "description": "a recommended multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "mci",
          "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
          "uid": "d2hfus4djuna1t2ovmj0",
          "name": "mmci01",
          "subGroupId": "",
          "location": {
            "display": "",
            "latitude": 0,
            "longitude": 0
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "",
          "networkAgentStatus": "",
          "systemMessage": "",
          "createdTime": "",
          "label": null,
          "description": "a recommended multi-cloud infrastructure",
          "region": {
            "Region": "",
            "Zone": ""
          },
          "publicIP": "",
          "sshPort": "",
          "publicDNS": "",
          "privateIP": "",
          "privateDNS": "",
          "rootDiskType": "",
          "rootDiskSize": "",
          "rootDiskName": "",
          "connectionName": "",
          "connectionConfig": {
            "configName": "",
            "providerName": "",
            "driverName": "",
            "credentialName": "",
            "credentialHolder": "",
            "regionZoneInfoName": "",
            "regionZoneInfo": {
              "assignedRegion": "",
              "assignedZone": ""
            },
            "regionDetail": {
              "regionId": "",
              "regionName": "",
              "description": "",
              "location": {
                "display": "",
                "latitude": 0,
                "longitude": 0
              },
              "zones": null
            },
            "regionRepresentative": false,
            "verified": false
          },
          "specId": "",
          "cspSpecName": "",
          "imageId": "",
          "cspImageName": "",
          "vNetId": "",
          "cspVNetId": "",
          "subnetId": "",
          "cspSubnetId": "",
          "networkInterface": "",
          "securityGroupIds": null,
          "dataDiskIds": null,
          "sshKeyId": "",
          "cspSshKeyId": ""
        }
      ],
      "newVmList": null,
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
  "uid": "d2hfus4djuna1t2ovmj0",
  "name": "mmci01",
  "status": "Running:1 (R:1/1)",
  "statusCount": {
    "countTotal": 1,
    "countCreating": 0,
    "countRunning": 1,
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
    "sys.uid": "d2hfus4djuna1t2ovmj0"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "uid": "d2hfus4djuna1t2ovmk0",
      "cspResourceName": "d2hfus4djuna1t2ovmk0",
      "cspResourceId": "i-mj7cxi74lacgci8l7cl9",
      "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
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
      "createdTime": "2025-08-18 10:25:33",
      "label": {
        "sys.connectionName": "alibaba-ap-northeast-2",
        "sys.createdTime": "2025-08-18 10:25:33",
        "sys.cspResourceId": "i-mj7cxi74lacgci8l7cl9",
        "sys.cspResourceName": "d2hfus4djuna1t2ovmk0",
        "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.uid": "d2hfus4djuna1t2ovmk0"
      },
      "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "8.220.214.219",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.20",
      "privateDNS": "",
      "rootDiskType": "cloud_essd",
      "rootDiskSize": "40",
      "rootDiskName": "",
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
      "specId": "ecs.e-c1m4.large",
      "cspSpecName": "ecs.e-c1m4.large",
      "imageId": "ubuntu_22_04_x64_20G_alibase_20250722.vhd",
      "cspImageName": "ubuntu_22_04_x64_20G_alibase_20250722.vhd",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-mj7rn5qgal81ddfdf1fq4",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "vsw-mj703pc1l4txdkwjfirxc",
      "networkInterface": "eni-mj7cxi74lacgci8lewrb",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d2hfuqcdjuna1t2ovmi0",
      "vmUserName": "cb-user",
      "vmUserPassword": "$dd!uj4fs2h1uA",
      "addtionalDetails": [
        {
          "key": "ImageId",
          "value": "ubuntu_22_04_x64_20G_alibase_20250722.vhd"
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
          "value": "d2hfus4djuna1t2ovmk0"
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
          "value": "2025-08-18T10:25Z"
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
          "value": "iZmj7cxi74lacgci8l7cl9Z"
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
          "value": "32bffe3d-7bcc-4156-a53f-9821406ed1e0"
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
          "value": "i-mj7cxi74lacgci8l7cl9"
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
          "value": "2025-08-18T10:25Z"
        },
        {
          "key": "KeyPairName",
          "value": "d2hfuqcdjuna1t2ovmi0"
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
          "value": "{SecurityGroupId:[sg-mj7huz3swfzzhrto8waq]}"
        },
        {
          "key": "InnerIpAddress",
          "value": "{IpAddress:[]}"
        },
        {
          "key": "PublicIpAddress",
          "value": "{IpAddress:[8.220.214.219]}"
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
          "value": "{VSwitchId:vsw-mj703pc1l4txdkwjfirxc,VpcId:vpc-mj7rn5qgal81ddfdf1fq4,NatIpAddress:,PrivateIpAddress:{IpAddress:[192.168.110.20]}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        },
        {
          "key": "NetworkInterfaces",
          "value": "{NetworkInterface:[{SecurityGroupId:,VSwitchId:,DeleteOnRelease:false,InstanceType:,MacAddress:00:16:3e:01:b6:d5,NetworkInterfaceTrafficMode:,NetworkInterfaceName:,NetworkInterfaceId:eni-mj7cxi74lacgci8lewrb,PrimaryIpAddress:192.168.110.20,Description:,Type:Primary,SecurityGroupIds:{SecurityGroupId:null},Ipv6PrefixSets:{Ipv6PrefixSet:null},Ipv4PrefixSets:{Ipv4PrefixSet:null},Ipv6Sets:{Ipv6Set:null},PrivateIpSets:{PrivateIpSet:[{PrivateIpAddress:192.168.110.20,PrivateDnsName:,Primary:true,AssociatedPublicIp:{PublicIpAddress:,AllocationId:}}]}}]}"
        },
        {
          "key": "OperationLocks",
          "value": "{LockReason:[]}"
        }
      ]
    }
  ],
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

### Test Case 6: Delete the migrated computing infra

#### 6.1 API Request Information

- **API Endpoint**: `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}`
- **Purpose**: Delete the migrated infrastructure and clean up resources
- **Namespace ID**: `mig01`
- **Path Parameter**: `{{mciId}}` - The MCI identifier to delete
- **Query Parameter**: `option=terminate` (terminates all resources)
- **Request Body**: None (DELETE request)

#### 6.2 API Response Information

- **Status**: ✅ **SUCCESS**
- **Response**: Infrastructure deletion completed successfully

**Response Body**:

```json
{
  "success": true,
  "text": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: mmci01)"
}
```

