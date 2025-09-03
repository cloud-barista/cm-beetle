# CM-Beetle test results for AZURE

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AZURE cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.0 (b250b72)
- cm-model: v0.0.13
- CB-Tumblebug: v0.11.9
- CB-Spider: v0.11.5
- CB-MapUI: v0.11.12
- Target CSP: AZURE
- Target Region: koreacentral
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: September 3, 2025
- Test Time: 12:18:04 KST
- Test Execution: 2025-09-03 12:18:04 KST

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
| 1 | `POST /beetle/recommendation/mci` | ✅ **PASS** | 434ms | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 4m57.195s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 3.498s | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 62ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 2.813s | Pass |
| 6 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 2m19.173s | Pass |

**Overall Result**: 6/6 tests passed ✅

**Total Duration**: 7m53.294494697s

*Test executed on September 3, 2025 at 12:18:04 KST (2025-09-03 12:18:04 KST) using CM-Beetle automated test CLI*

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
    "csp": "azure",
    "region": "koreacentral"
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
    "csp": "azure",
    "region": "koreacentral"
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
        "connectionName": "azure-koreacentral",
        "specId": "azure+koreacentral+standard_b4as_v2",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
        "vNetId": "mig-vnet-01",
        "subnetId": "mig-subnet-01",
        "securityGroupIds": [
          "mig-sg-01"
        ],
        "sshKeyId": "mig-sshkey-01",
        "dataDiskIds": null
      },
      {
        "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "subGroupSize": "",
        "label": {
          "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c"
        },
        "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "connectionName": "azure-koreacentral",
        "specId": "azure+koreacentral+standard_b4as_v2",
        "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
        "vNetId": "mig-vnet-01",
        "subnetId": "mig-subnet-01",
        "securityGroupIds": [
          "mig-sg-02"
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
    "connectionName": "azure-koreacentral",
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
    "connectionName": "azure-koreacentral",
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
      "id": "azure+koreacentral+standard_b4as_v2",
      "uid": "d2rpbhlr41dql6d5ft60",
      "cspSpecName": "Standard_B4as_v2",
      "name": "azure+koreacentral+standard_b4as_v2",
      "namespace": "system",
      "connectionName": "azure-koreacentral",
      "providerName": "azure",
      "regionName": "koreacentral",
      "regionLatitude": 37.5665,
      "regionLongitude": 126.978,
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 4,
      "memoryGiB": 15.625,
      "costPerHour": 0.187,
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
    }
  ],
  "targetVmOsImageList": [
    {
      "namespace": "system",
      "providerName": "azure",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "regionList": [
        "common"
      ],
      "id": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "uid": "d2rpdndr41dql6dcvi00",
      "name": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "connectionName": "azure-australiacentral",
      "infraType": "vm",
      "fetchedTime": "2025.09.03 01:16:13 Wed",
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
          "value": "22.04.202507300"
        },
        {
          "key": "ID",
          "value": "/Subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/0001-com-ubuntu-server-jammy/Skus/22_04-lts-gen2/Versions/22.04.202507300"
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
      "connectionName": "azure-koreacentral",
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
      "connectionName": "azure-koreacentral",
      "vNetId": "mig-vnet-01",
      "description": "Recommended security group for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "firewallRules": [
        {
          "Ports": "22",
          "Protocol": "tcp",
          "Direction": "inbound",
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
  "uid": "d2rr7l5r41dql6dh7180",
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
    "sys.uid": "d2rr7l5r41dql6dh7180"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "uid": "d2rr7l5r41dql6dh71a0",
      "cspResourceName": "d2rr7l5r41dql6dh71a0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh71a0",
      "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "location": {
        "display": "Korea Central",
        "latitude": 37.5665,
        "longitude": 126.978
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-09-03 03:20:26",
      "label": {
        "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "sys.connectionName": "azure-koreacentral",
        "sys.createdTime": "2025-09-03 03:20:26",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh71a0",
        "sys.cspResourceName": "d2rr7l5r41dql6dh71a0",
        "sys.id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "sys.uid": "d2rr7l5r41dql6dh71a0"
      },
      "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "region": {
        "Region": "koreacentral",
        "Zone": "1"
      },
      "publicIP": "40.82.155.99",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.4",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "30",
      "rootDiskName": "",
      "connectionName": "azure-koreacentral",
      "connectionConfig": {
        "configName": "azure-koreacentral",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreacentral",
        "regionZoneInfo": {
          "assignedRegion": "koreacentral",
          "assignedZone": "1"
        },
        "regionDetail": {
          "regionId": "koreacentral",
          "regionName": "koreacentral",
          "description": "Korea Central",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          },
          "zones": [
            "1",
            "2",
            "3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreacentral+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g/subnets/d2rr6u5r41dql6dh7160",
      "networkInterface": "d2rr7l5r41dql6dh71a0-40940-VNic",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/sshPublicKeys/d2rr70tr41dql6dh716g",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreacentral"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/networkInterfaces/d2rr7l5r41dql6dh71a0-40940-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d2rr7l5r41dql6dh71a0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDGDZI7e62jeivTxmz8zP4z94StF9ulkhCaxAS0aTFYi7fGZPvgsCKdacy0HBQCHBzbRb8iOckSuzFERmzwMBa6GbQorqA1iSZhNAocYfOFKQUrskudUvL4AWRRwvG0rqS2PmzUnQD8PcFIstCQINFsFzaKBTOF79ZF/Qi0ovc3Gm+Ji2XWX/0d21hdSbSvXp4zClxXAXzAcuZsinRYsE1Rb0VvkMvjcuv858BOq2bLVDFs2Eve9u4HDimFCute2sfpN9HPIO+Ku3BkZhCtRRtsDKD9TW8glEvuUjonL209LZtDbN+Gg0MRd89vgHtqwWCQKAvJLi9yX5S9A5rlO0vKvxXS9jDvEWeNHJm8JkwOtHbpsPIH8RKQJK6T9HJdDGiZiHH8b5O5O31w+8gM1BObjnBKHDl1W7c2cmiCTB990LJSs//4tcxZv3vVh+keNmpt+bMFGRtOktK7dMxN/SIF3r4SVSLMbW1T3HwYV8dF1qdIScqKSfXglkpK+Up+HTcLsDgGyGaCnWEx6SairBbYkir2zZdeLQh/AumE7yTF3tclKuhKw4c4oOPL1YZCWgSz7N2drTLoZzVHyzebyj5AYRx8vn39gFyDzupIMMz+8hbh1J31Ijbz4pPk/tVTwjC6vEcVNk8/MCWIvz9s+n0dId0bw+pnrpoLsmKhPaqDvw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202507300,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202507300},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/disks/d2rr7l5r41dql6dh71a0_disk1_c998e530ef1143fbb56aecc53d9d884e,storageAccountType:Premium_LRS},name:d2rr7l5r41dql6dh71a0_disk1_c998e530ef1143fbb56aecc53d9d884e,osType:Linux}},timeCreated:2025-09-03T03:20:05.7312456Z,vmId:d5f412cd-8d0f-4d60-9ff5-25eb14c700e6}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d2rr7l5r41dql6dh71a0,keypair:d2rr70tr41dql6dh716g,publicip:d2rr7l5r41dql6dh71a0-22216-PublicIP}"
        },
        {
          "key": "Zones",
          "value": "1"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh71a0"
        },
        {
          "key": "Name",
          "value": "d2rr7l5r41dql6dh71a0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "uid": "d2rr7l5r41dql6dh7190",
      "cspResourceName": "d2rr7l5r41dql6dh7190",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh7190",
      "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "location": {
        "display": "Korea Central",
        "latitude": 37.5665,
        "longitude": 126.978
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-09-03 03:22:34",
      "label": {
        "sourceMachineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.connectionName": "azure-koreacentral",
        "sys.createdTime": "2025-09-03 03:22:34",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh7190",
        "sys.cspResourceName": "d2rr7l5r41dql6dh7190",
        "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.uid": "d2rr7l5r41dql6dh7190"
      },
      "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "region": {
        "Region": "koreacentral",
        "Zone": "1"
      },
      "publicIP": "20.196.112.189",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "30",
      "rootDiskName": "",
      "connectionName": "azure-koreacentral",
      "connectionConfig": {
        "configName": "azure-koreacentral",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreacentral",
        "regionZoneInfo": {
          "assignedRegion": "koreacentral",
          "assignedZone": "1"
        },
        "regionDetail": {
          "regionId": "koreacentral",
          "regionName": "koreacentral",
          "description": "Korea Central",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          },
          "zones": [
            "1",
            "2",
            "3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreacentral+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g/subnets/d2rr6u5r41dql6dh7160",
      "networkInterface": "d2rr7l5r41dql6dh7190-82978-VNic",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/sshPublicKeys/d2rr70tr41dql6dh716g",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreacentral"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/networkInterfaces/d2rr7l5r41dql6dh7190-82978-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d2rr7l5r41dql6dh7190,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDGDZI7e62jeivTxmz8zP4z94StF9ulkhCaxAS0aTFYi7fGZPvgsCKdacy0HBQCHBzbRb8iOckSuzFERmzwMBa6GbQorqA1iSZhNAocYfOFKQUrskudUvL4AWRRwvG0rqS2PmzUnQD8PcFIstCQINFsFzaKBTOF79ZF/Qi0ovc3Gm+Ji2XWX/0d21hdSbSvXp4zClxXAXzAcuZsinRYsE1Rb0VvkMvjcuv858BOq2bLVDFs2Eve9u4HDimFCute2sfpN9HPIO+Ku3BkZhCtRRtsDKD9TW8glEvuUjonL209LZtDbN+Gg0MRd89vgHtqwWCQKAvJLi9yX5S9A5rlO0vKvxXS9jDvEWeNHJm8JkwOtHbpsPIH8RKQJK6T9HJdDGiZiHH8b5O5O31w+8gM1BObjnBKHDl1W7c2cmiCTB990LJSs//4tcxZv3vVh+keNmpt+bMFGRtOktK7dMxN/SIF3r4SVSLMbW1T3HwYV8dF1qdIScqKSfXglkpK+Up+HTcLsDgGyGaCnWEx6SairBbYkir2zZdeLQh/AumE7yTF3tclKuhKw4c4oOPL1YZCWgSz7N2drTLoZzVHyzebyj5AYRx8vn39gFyDzupIMMz+8hbh1J31Ijbz4pPk/tVTwjC6vEcVNk8/MCWIvz9s+n0dId0bw+pnrpoLsmKhPaqDvw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202507300,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202507300},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/disks/d2rr7l5r41dql6dh7190_disk1_a8a9277249864b818b6d062213acd543,storageAccountType:Premium_LRS},name:d2rr7l5r41dql6dh7190_disk1_a8a9277249864b818b6d062213acd543,osType:Linux}},timeCreated:2025-09-03T03:20:10.371995Z,vmId:fc00fa2c-ad99-4ba8-a57f-850dad3b1de7}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d2rr7l5r41dql6dh7190,keypair:d2rr70tr41dql6dh716g,publicip:d2rr7l5r41dql6dh7190-83889-PublicIP}"
        },
        {
          "key": "Zones",
          "value": "1"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh7190"
        },
        {
          "key": "Name",
          "value": "d2rr7l5r41dql6dh7190"
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
      "uid": "d2rr7l5r41dql6dh7180",
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
        "sys.uid": "d2rr7l5r41dql6dh7180"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "a recommended multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
          "uid": "d2rr7l5r41dql6dh71a0",
          "cspResourceName": "d2rr7l5r41dql6dh71a0",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh71a0",
          "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
          "subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-09-03 03:20:26",
          "label": {
            "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c",
            "sys.connectionName": "azure-koreacentral",
            "sys.createdTime": "2025-09-03 03:20:26",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh71a0",
            "sys.cspResourceName": "d2rr7l5r41dql6dh71a0",
            "sys.id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
            "sys.uid": "d2rr7l5r41dql6dh71a0"
          },
          "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
          "region": {
            "Region": "koreacentral",
            "Zone": "1"
          },
          "publicIP": "40.82.155.99",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "192.168.110.4",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": "30",
          "rootDiskName": "",
          "connectionName": "azure-koreacentral",
          "connectionConfig": {
            "configName": "azure-koreacentral",
            "providerName": "azure",
            "driverName": "azure-driver-v1.0.so",
            "credentialName": "azure",
            "credentialHolder": "admin",
            "regionZoneInfoName": "azure-koreacentral",
            "regionZoneInfo": {
              "assignedRegion": "koreacentral",
              "assignedZone": "1"
            },
            "regionDetail": {
              "regionId": "koreacentral",
              "regionName": "koreacentral",
              "description": "Korea Central",
              "location": {
                "display": "Korea Central",
                "latitude": 37.5665,
                "longitude": 126.978
              },
              "zones": [
                "1",
                "2",
                "3"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "azure+koreacentral+standard_b4as_v2",
          "cspSpecName": "Standard_B4as_v2",
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g/subnets/d2rr6u5r41dql6dh7160",
          "networkInterface": "d2rr7l5r41dql6dh71a0-40940-VNic",
          "securityGroupIds": [
            "mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/sshPublicKeys/d2rr70tr41dql6dh716g",
          "vmUserName": "cb-user",
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreacentral"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/networkInterfaces/d2rr7l5r41dql6dh71a0-40940-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d2rr7l5r41dql6dh71a0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDGDZI7e62jeivTxmz8zP4z94StF9ulkhCaxAS0aTFYi7fGZPvgsCKdacy0HBQCHBzbRb8iOckSuzFERmzwMBa6GbQorqA1iSZhNAocYfOFKQUrskudUvL4AWRRwvG0rqS2PmzUnQD8PcFIstCQINFsFzaKBTOF79ZF/Qi0ovc3Gm+Ji2XWX/0d21hdSbSvXp4zClxXAXzAcuZsinRYsE1Rb0VvkMvjcuv858BOq2bLVDFs2Eve9u4HDimFCute2sfpN9HPIO+Ku3BkZhCtRRtsDKD9TW8glEvuUjonL209LZtDbN+Gg0MRd89vgHtqwWCQKAvJLi9yX5S9A5rlO0vKvxXS9jDvEWeNHJm8JkwOtHbpsPIH8RKQJK6T9HJdDGiZiHH8b5O5O31w+8gM1BObjnBKHDl1W7c2cmiCTB990LJSs//4tcxZv3vVh+keNmpt+bMFGRtOktK7dMxN/SIF3r4SVSLMbW1T3HwYV8dF1qdIScqKSfXglkpK+Up+HTcLsDgGyGaCnWEx6SairBbYkir2zZdeLQh/AumE7yTF3tclKuhKw4c4oOPL1YZCWgSz7N2drTLoZzVHyzebyj5AYRx8vn39gFyDzupIMMz+8hbh1J31Ijbz4pPk/tVTwjC6vEcVNk8/MCWIvz9s+n0dId0bw+pnrpoLsmKhPaqDvw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202507300,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202507300},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/disks/d2rr7l5r41dql6dh71a0_disk1_c998e530ef1143fbb56aecc53d9d884e,storageAccountType:Premium_LRS},name:d2rr7l5r41dql6dh71a0_disk1_c998e530ef1143fbb56aecc53d9d884e,osType:Linux}},timeCreated:2025-09-03T03:20:05.7312456Z,vmId:d5f412cd-8d0f-4d60-9ff5-25eb14c700e6}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d2rr7l5r41dql6dh71a0,keypair:d2rr70tr41dql6dh716g,publicip:d2rr7l5r41dql6dh71a0-22216-PublicIP}"
            },
            {
              "key": "Zones",
              "value": "1"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh71a0"
            },
            {
              "key": "Name",
              "value": "d2rr7l5r41dql6dh71a0"
            },
            {
              "key": "Type",
              "value": "Microsoft.Compute/virtualMachines"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
          "uid": "d2rr7l5r41dql6dh7190",
          "cspResourceName": "d2rr7l5r41dql6dh7190",
          "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh7190",
          "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
          "subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          },
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "monAgentStatus": "notInstalled",
          "networkAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2025-09-03 03:22:34",
          "label": {
            "sourceMachineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c",
            "sys.connectionName": "azure-koreacentral",
            "sys.createdTime": "2025-09-03 03:22:34",
            "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh7190",
            "sys.cspResourceName": "d2rr7l5r41dql6dh7190",
            "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
            "sys.uid": "d2rr7l5r41dql6dh7190"
          },
          "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
          "region": {
            "Region": "koreacentral",
            "Zone": "1"
          },
          "publicIP": "20.196.112.189",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "192.168.110.5",
          "privateDNS": "",
          "rootDiskType": "PremiumSSD",
          "rootDiskSize": "30",
          "rootDiskName": "",
          "connectionName": "azure-koreacentral",
          "connectionConfig": {
            "configName": "azure-koreacentral",
            "providerName": "azure",
            "driverName": "azure-driver-v1.0.so",
            "credentialName": "azure",
            "credentialHolder": "admin",
            "regionZoneInfoName": "azure-koreacentral",
            "regionZoneInfo": {
              "assignedRegion": "koreacentral",
              "assignedZone": "1"
            },
            "regionDetail": {
              "regionId": "koreacentral",
              "regionName": "koreacentral",
              "description": "Korea Central",
              "location": {
                "display": "Korea Central",
                "latitude": 37.5665,
                "longitude": 126.978
              },
              "zones": [
                "1",
                "2",
                "3"
              ]
            },
            "regionRepresentative": true,
            "verified": true
          },
          "specId": "azure+koreacentral+standard_b4as_v2",
          "cspSpecName": "Standard_B4as_v2",
          "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
          "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g/subnets/d2rr6u5r41dql6dh7160",
          "networkInterface": "d2rr7l5r41dql6dh7190-82978-VNic",
          "securityGroupIds": [
            "mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/sshPublicKeys/d2rr70tr41dql6dh716g",
          "vmUserName": "cb-user",
          "addtionalDetails": [
            {
              "key": "Location",
              "value": "koreacentral"
            },
            {
              "key": "Properties",
              "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/networkInterfaces/d2rr7l5r41dql6dh7190-82978-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d2rr7l5r41dql6dh7190,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDGDZI7e62jeivTxmz8zP4z94StF9ulkhCaxAS0aTFYi7fGZPvgsCKdacy0HBQCHBzbRb8iOckSuzFERmzwMBa6GbQorqA1iSZhNAocYfOFKQUrskudUvL4AWRRwvG0rqS2PmzUnQD8PcFIstCQINFsFzaKBTOF79ZF/Qi0ovc3Gm+Ji2XWX/0d21hdSbSvXp4zClxXAXzAcuZsinRYsE1Rb0VvkMvjcuv858BOq2bLVDFs2Eve9u4HDimFCute2sfpN9HPIO+Ku3BkZhCtRRtsDKD9TW8glEvuUjonL209LZtDbN+Gg0MRd89vgHtqwWCQKAvJLi9yX5S9A5rlO0vKvxXS9jDvEWeNHJm8JkwOtHbpsPIH8RKQJK6T9HJdDGiZiHH8b5O5O31w+8gM1BObjnBKHDl1W7c2cmiCTB990LJSs//4tcxZv3vVh+keNmpt+bMFGRtOktK7dMxN/SIF3r4SVSLMbW1T3HwYV8dF1qdIScqKSfXglkpK+Up+HTcLsDgGyGaCnWEx6SairBbYkir2zZdeLQh/AumE7yTF3tclKuhKw4c4oOPL1YZCWgSz7N2drTLoZzVHyzebyj5AYRx8vn39gFyDzupIMMz+8hbh1J31Ijbz4pPk/tVTwjC6vEcVNk8/MCWIvz9s+n0dId0bw+pnrpoLsmKhPaqDvw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202507300,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202507300},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/disks/d2rr7l5r41dql6dh7190_disk1_a8a9277249864b818b6d062213acd543,storageAccountType:Premium_LRS},name:d2rr7l5r41dql6dh7190_disk1_a8a9277249864b818b6d062213acd543,osType:Linux}},timeCreated:2025-09-03T03:20:10.371995Z,vmId:fc00fa2c-ad99-4ba8-a57f-850dad3b1de7}"
            },
            {
              "key": "Tags",
              "value": "{createdBy:d2rr7l5r41dql6dh7190,keypair:d2rr70tr41dql6dh716g,publicip:d2rr7l5r41dql6dh7190-83889-PublicIP}"
            },
            {
              "key": "Zones",
              "value": "1"
            },
            {
              "key": "Etag",
              "value": "\\1\\"
            },
            {
              "key": "ID",
              "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh7190"
            },
            {
              "key": "Name",
              "value": "d2rr7l5r41dql6dh7190"
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
  "uid": "d2rr7l5r41dql6dh7180",
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
    "sys.uid": "d2rr7l5r41dql6dh7180"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "uid": "d2rr7l5r41dql6dh71a0",
      "cspResourceName": "d2rr7l5r41dql6dh71a0",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh71a0",
      "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "location": {
        "display": "Korea Central",
        "latitude": 37.5665,
        "longitude": 126.978
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-09-03 03:20:26",
      "label": {
        "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "sys.connectionName": "azure-koreacentral",
        "sys.createdTime": "2025-09-03 03:20:26",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh71a0",
        "sys.cspResourceName": "d2rr7l5r41dql6dh71a0",
        "sys.id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "sys.uid": "d2rr7l5r41dql6dh71a0"
      },
      "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "region": {
        "Region": "koreacentral",
        "Zone": "1"
      },
      "publicIP": "40.82.155.99",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.4",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "30",
      "rootDiskName": "",
      "connectionName": "azure-koreacentral",
      "connectionConfig": {
        "configName": "azure-koreacentral",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreacentral",
        "regionZoneInfo": {
          "assignedRegion": "koreacentral",
          "assignedZone": "1"
        },
        "regionDetail": {
          "regionId": "koreacentral",
          "regionName": "koreacentral",
          "description": "Korea Central",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          },
          "zones": [
            "1",
            "2",
            "3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreacentral+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g/subnets/d2rr6u5r41dql6dh7160",
      "networkInterface": "d2rr7l5r41dql6dh71a0-40940-VNic",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/sshPublicKeys/d2rr70tr41dql6dh716g",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreacentral"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/networkInterfaces/d2rr7l5r41dql6dh71a0-40940-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d2rr7l5r41dql6dh71a0,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDGDZI7e62jeivTxmz8zP4z94StF9ulkhCaxAS0aTFYi7fGZPvgsCKdacy0HBQCHBzbRb8iOckSuzFERmzwMBa6GbQorqA1iSZhNAocYfOFKQUrskudUvL4AWRRwvG0rqS2PmzUnQD8PcFIstCQINFsFzaKBTOF79ZF/Qi0ovc3Gm+Ji2XWX/0d21hdSbSvXp4zClxXAXzAcuZsinRYsE1Rb0VvkMvjcuv858BOq2bLVDFs2Eve9u4HDimFCute2sfpN9HPIO+Ku3BkZhCtRRtsDKD9TW8glEvuUjonL209LZtDbN+Gg0MRd89vgHtqwWCQKAvJLi9yX5S9A5rlO0vKvxXS9jDvEWeNHJm8JkwOtHbpsPIH8RKQJK6T9HJdDGiZiHH8b5O5O31w+8gM1BObjnBKHDl1W7c2cmiCTB990LJSs//4tcxZv3vVh+keNmpt+bMFGRtOktK7dMxN/SIF3r4SVSLMbW1T3HwYV8dF1qdIScqKSfXglkpK+Up+HTcLsDgGyGaCnWEx6SairBbYkir2zZdeLQh/AumE7yTF3tclKuhKw4c4oOPL1YZCWgSz7N2drTLoZzVHyzebyj5AYRx8vn39gFyDzupIMMz+8hbh1J31Ijbz4pPk/tVTwjC6vEcVNk8/MCWIvz9s+n0dId0bw+pnrpoLsmKhPaqDvw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202507300,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202507300},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/disks/d2rr7l5r41dql6dh71a0_disk1_c998e530ef1143fbb56aecc53d9d884e,storageAccountType:Premium_LRS},name:d2rr7l5r41dql6dh71a0_disk1_c998e530ef1143fbb56aecc53d9d884e,osType:Linux}},timeCreated:2025-09-03T03:20:05.7312456Z,vmId:d5f412cd-8d0f-4d60-9ff5-25eb14c700e6}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d2rr7l5r41dql6dh71a0,keypair:d2rr70tr41dql6dh716g,publicip:d2rr7l5r41dql6dh71a0-22216-PublicIP}"
        },
        {
          "key": "Zones",
          "value": "1"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh71a0"
        },
        {
          "key": "Name",
          "value": "d2rr7l5r41dql6dh71a0"
        },
        {
          "key": "Type",
          "value": "Microsoft.Compute/virtualMachines"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "uid": "d2rr7l5r41dql6dh7190",
      "cspResourceName": "d2rr7l5r41dql6dh7190",
      "cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh7190",
      "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "location": {
        "display": "Korea Central",
        "latitude": 37.5665,
        "longitude": 126.978
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-09-03 03:22:34",
      "label": {
        "sourceMachineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.connectionName": "azure-koreacentral",
        "sys.createdTime": "2025-09-03 03:22:34",
        "sys.cspResourceId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh7190",
        "sys.cspResourceName": "d2rr7l5r41dql6dh7190",
        "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.uid": "d2rr7l5r41dql6dh7190"
      },
      "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "region": {
        "Region": "koreacentral",
        "Zone": "1"
      },
      "publicIP": "20.196.112.189",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "30",
      "rootDiskName": "",
      "connectionName": "azure-koreacentral",
      "connectionConfig": {
        "configName": "azure-koreacentral",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreacentral",
        "regionZoneInfo": {
          "assignedRegion": "koreacentral",
          "assignedZone": "1"
        },
        "regionDetail": {
          "regionId": "koreacentral",
          "regionName": "koreacentral",
          "description": "Korea Central",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          },
          "zones": [
            "1",
            "2",
            "3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreacentral+standard_b4as_v2",
      "cspSpecName": "Standard_B4as_v2",
      "imageId": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202507300",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/d2rr6u5r41dql6dh715g/subnets/d2rr6u5r41dql6dh7160",
      "networkInterface": "d2rr7l5r41dql6dh7190-82978-VNic",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/sshPublicKeys/d2rr70tr41dql6dh716g",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "Location",
          "value": "koreacentral"
        },
        {
          "key": "Properties",
          "value": "{hardwareProfile:{vmSize:Standard_B4as_v2},networkProfile:{networkInterfaces:[{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Network/networkInterfaces/d2rr7l5r41dql6dh7190-82978-VNic,properties:{primary:true}}]},osProfile:{adminUsername:cb-user,allowExtensionOperations:true,computerName:d2rr7l5r41dql6dh7190,linuxConfiguration:{disablePasswordAuthentication:true,patchSettings:{assessmentMode:ImageDefault,patchMode:ImageDefault},provisionVMAgent:true,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDGDZI7e62jeivTxmz8zP4z94StF9ulkhCaxAS0aTFYi7fGZPvgsCKdacy0HBQCHBzbRb8iOckSuzFERmzwMBa6GbQorqA1iSZhNAocYfOFKQUrskudUvL4AWRRwvG0rqS2PmzUnQD8PcFIstCQINFsFzaKBTOF79ZF/Qi0ovc3Gm+Ji2XWX/0d21hdSbSvXp4zClxXAXzAcuZsinRYsE1Rb0VvkMvjcuv858BOq2bLVDFs2Eve9u4HDimFCute2sfpN9HPIO+Ku3BkZhCtRRtsDKD9TW8glEvuUjonL209LZtDbN+Gg0MRd89vgHtqwWCQKAvJLi9yX5S9A5rlO0vKvxXS9jDvEWeNHJm8JkwOtHbpsPIH8RKQJK6T9HJdDGiZiHH8b5O5O31w+8gM1BObjnBKHDl1W7c2cmiCTB990LJSs//4tcxZv3vVh+keNmpt+bMFGRtOktK7dMxN/SIF3r4SVSLMbW1T3HwYV8dF1qdIScqKSfXglkpK+Up+HTcLsDgGyGaCnWEx6SairBbYkir2zZdeLQh/AumE7yTF3tclKuhKw4c4oOPL1YZCWgSz7N2drTLoZzVHyzebyj5AYRx8vn39gFyDzupIMMz+8hbh1J31Ijbz4pPk/tVTwjC6vEcVNk8/MCWIvz9s+n0dId0bw+pnrpoLsmKhPaqDvw==\\n,path:/home/cb-user/.ssh/authorized_keys}]}},requireGuestProvisionSignal:true,secrets:[]},provisioningState:Succeeded,storageProfile:{dataDisks:[],diskControllerType:SCSI,imageReference:{exactVersion:22.04.202507300,offer:0001-com-ubuntu-server-jammy,publisher:Canonical,sku:22_04-lts-gen2,version:22.04.202507300},osDisk:{caching:ReadWrite,createOption:FromImage,deleteOption:Delete,diskSizeGB:30,managedDisk:{id:/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/disks/d2rr7l5r41dql6dh7190_disk1_a8a9277249864b818b6d062213acd543,storageAccountType:Premium_LRS},name:d2rr7l5r41dql6dh7190_disk1_a8a9277249864b818b6d062213acd543,osType:Linux}},timeCreated:2025-09-03T03:20:10.371995Z,vmId:fc00fa2c-ad99-4ba8-a57f-850dad3b1de7}"
        },
        {
          "key": "Tags",
          "value": "{createdBy:d2rr7l5r41dql6dh7190,keypair:d2rr70tr41dql6dh716g,publicip:d2rr7l5r41dql6dh7190-83889-PublicIP}"
        },
        {
          "key": "Zones",
          "value": "1"
        },
        {
          "key": "Etag",
          "value": "\\1\\"
        },
        {
          "key": "ID",
          "value": "/subscriptions/a20fed83-96bd-4480-92a9-140b8e3b7c3a/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/d2rr7l5r41dql6dh7190"
        },
        {
          "key": "Name",
          "value": "d2rr7l5r41dql6dh7190"
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

