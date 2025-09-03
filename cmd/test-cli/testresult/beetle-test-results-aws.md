# CM-Beetle test results for AWS

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with AWS cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.0 (b250b72)
- cm-model: v0.0.13
- CB-Tumblebug: v0.11.9
- CB-Spider: v0.11.5
- CB-MapUI: v0.11.12
- Target CSP: AWS
- Target Region: ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: September 3, 2025
- Test Time: 12:08:49 KST
- Test Execution: 2025-09-03 12:08:49 KST

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for AWS

### Test Results Summary

| Test | Endpoint | Status | Duration | Details |
|------|----------|--------|----------|----------|
| 1 | `POST /beetle/recommendation/mci` | ✅ **PASS** | 3.663s | Pass |
| 2 | `POST /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 2m51.536s | Pass |
| 3 | `GET /beetle/migration/ns/mig01/mci` | ✅ **PASS** | 195ms | Pass |
| 4 | `GET /beetle/migration/ns/mig01/mci?option=id` | ✅ **PASS** | 62ms | Pass |
| 5 | `GET /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 174ms | Pass |
| 6 | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 5m31.008s | Pass |

**Overall Result**: 6/6 tests passed ✅

**Total Duration**: 8m56.745004425s

*Test executed on September 3, 2025 at 12:08:49 KST (2025-09-03 12:08:49 KST) using CM-Beetle automated test CLI*

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
    "csp": "aws",
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
    "csp": "aws",
    "region": "ap-northeast-2"
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
        "connectionName": "aws-ap-northeast-2",
        "specId": "aws+ap-northeast-2+g4dn.xlarge",
        "imageId": "ami-08943a151bd468f4e",
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
        "connectionName": "aws-ap-northeast-2",
        "specId": "aws+ap-northeast-2+g4dn.xlarge",
        "imageId": "ami-08943a151bd468f4e",
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
    "connectionName": "aws-ap-northeast-2",
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
      "id": "aws+ap-northeast-2+g4dn.xlarge",
      "uid": "d2rpbfdr41dql6d5971g",
      "cspSpecName": "g4dn.xlarge",
      "name": "aws+ap-northeast-2+g4dn.xlarge",
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
      "acceleratorModel": "NVIDIA T4",
      "acceleratorCount": 1,
      "acceleratorMemoryGB": 17,
      "acceleratorType": "gpu",
      "costPerHour": 0.647,
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
          "key": "AutoRecoverySupported",
          "value": "false"
        },
        {
          "key": "BareMetal",
          "value": "false"
        },
        {
          "key": "BurstablePerformanceSupported",
          "value": "false"
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
          "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:950,BaselineIops:3000,BaselineThroughputInMBps:118.75,MaximumBandwidthInMbps:3500,MaximumIops:20000,MaximumThroughputInMBps:437.5},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
        },
        {
          "key": "FreeTierEligible",
          "value": "false"
        },
        {
          "key": "GpuInfo",
          "value": "{Gpus:[{Count:1,Manufacturer:NVIDIA,MemoryInfo:{SizeInMiB:16384},Name:T4}],TotalGpuMemoryInMiB:16384}"
        },
        {
          "key": "HibernationSupported",
          "value": "false"
        },
        {
          "key": "Hypervisor",
          "value": "nitro"
        },
        {
          "key": "InstanceStorageInfo",
          "value": "{Disks:[{Count:1,SizeInGB:125,Type:ssd}],NvmeSupport:required,TotalSizeInGB:125}"
        },
        {
          "key": "InstanceStorageSupported",
          "value": "true"
        },
        {
          "key": "InstanceType",
          "value": "g4dn.xlarge"
        },
        {
          "key": "MemoryInfo",
          "value": "{SizeInMiB:16384}"
        },
        {
          "key": "NetworkInfo",
          "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:10,Ipv6AddressesPerInterface:10,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 25 Gigabit}],NetworkPerformance:Up to 25 Gigabit}"
        },
        {
          "key": "PlacementGroupInfo",
          "value": "{SupportedStrategies:[cluster,partition,spread]}"
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
    }
  ],
  "targetVmOsImageList": [
    {
      "namespace": "system",
      "providerName": "aws",
      "cspImageName": "ami-08943a151bd468f4e",
      "regionList": [
        "ap-northeast-2"
      ],
      "id": "ami-08943a151bd468f4e",
      "uid": "d2rpbulr41dql6d73q1g",
      "name": "ami-08943a151bd468f4e",
      "connectionName": "aws-ap-northeast-2",
      "fetchedTime": "2025.09.03 01:12:26 Wed",
      "creationDate": "2025-05-16T06:42:23.000Z",
      "isBasicImage": true,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20250516",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-0aa76065655cac4d5,Throughput:null,VolumeSize:8,VolumeType:gp2},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
        },
        {
          "key": "BootMode",
          "value": "uefi-preferred"
        },
        {
          "key": "CreationDate",
          "value": "2025-05-16T06:42:23.000Z"
        },
        {
          "key": "DeprecationTime",
          "value": "2027-05-16T06:42:23.000Z"
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
          "value": "ami-08943a151bd468f4e"
        },
        {
          "key": "ImageLocation",
          "value": "amazon/ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20250516"
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
          "value": "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20250516"
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
      "description": "Canonical, Ubuntu, 22.04, amd64 jammy image"
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "mig-sg-01",
      "connectionName": "aws-ap-northeast-2",
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
      "connectionName": "aws-ap-northeast-2",
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
  "uid": "d2rr2n5r41dql6dh7120",
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
    "sys.uid": "d2rr2n5r41dql6dh7120"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "uid": "d2rr2n5r41dql6dh7140",
      "cspResourceName": "d2rr2n5r41dql6dh7140",
      "cspResourceId": "i-0076fe4936eb3f70e",
      "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
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
      "createdTime": "2025-09-03 03:09:49",
      "label": {
        "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2025-09-03 03:09:49",
        "sys.cspResourceId": "i-0076fe4936eb3f70e",
        "sys.cspResourceName": "d2rr2n5r41dql6dh7140",
        "sys.id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "sys.uid": "d2rr2n5r41dql6dh7140"
      },
      "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "43.202.62.178",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.117",
      "privateDNS": "ip-192-168-110-117.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "8",
      "rootDiskName": "",
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
      "specId": "aws+ap-northeast-2+g4dn.xlarge",
      "cspSpecName": "g4dn.xlarge",
      "imageId": "ami-08943a151bd468f4e",
      "cspImageName": "ami-08943a151bd468f4e",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-0ecd4eb6d98a69004",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "subnet-02ee9d54886666a0e",
      "networkInterface": "eni-attach-06f9192516cbf46ed",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d2rr2ldr41dql6dh710g",
      "vmUserName": "cb-user",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-09-03T03:09:34Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0b2ed52d3a9d83895}}"
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
          "value": "D47446CB-9622-487F-B738-6DE59C71B395"
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
          "value": "ami-08943a151bd468f4e"
        },
        {
          "key": "InstanceId",
          "value": "i-0076fe4936eb3f70e"
        },
        {
          "key": "InstanceType",
          "value": "g4dn.xlarge"
        },
        {
          "key": "KeyName",
          "value": "d2rr2ldr41dql6dh710g"
        },
        {
          "key": "LaunchTime",
          "value": "2025-09-03T03:09:33Z"
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
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.202.62.178},Attachment:{AttachTime:2025-09-03T03:09:33Z,AttachmentId:eni-attach-06f9192516cbf46ed,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b4d6930e1c7d30bc,GroupName:d2rr2mlr41dql6dh711g}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:f8:88:3f:3f:b1,NetworkInterfaceId:eni-06face27cb5f62cfb,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:192.168.110.117,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.202.62.178},Primary:true,PrivateDnsName:null,PrivateIpAddress:192.168.110.117}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-02ee9d54886666a0e,VpcId:vpc-0ecd4eb6d98a69004}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-192-168-110-117.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "192.168.110.117"
        },
        {
          "key": "PublicIpAddress",
          "value": "43.202.62.178"
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
          "value": "{GroupId:sg-0b4d6930e1c7d30bc,GroupName:d2rr2mlr41dql6dh711g}"
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
          "value": "subnet-02ee9d54886666a0e"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d2rr2n5r41dql6dh7140}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0ecd4eb6d98a69004"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "uid": "d2rr2n5r41dql6dh7130",
      "cspResourceName": "d2rr2n5r41dql6dh7130",
      "cspResourceId": "i-0d86a4c26af0079a1",
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
      "createdTime": "2025-09-03 03:11:54",
      "label": {
        "sourceMachineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2025-09-03 03:11:54",
        "sys.cspResourceId": "i-0d86a4c26af0079a1",
        "sys.cspResourceName": "d2rr2n5r41dql6dh7130",
        "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.uid": "d2rr2n5r41dql6dh7130"
      },
      "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "43.203.192.10",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.39",
      "privateDNS": "ip-192-168-110-39.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "8",
      "rootDiskName": "",
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
      "specId": "aws+ap-northeast-2+g4dn.xlarge",
      "cspSpecName": "g4dn.xlarge",
      "imageId": "ami-08943a151bd468f4e",
      "cspImageName": "ami-08943a151bd468f4e",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-0ecd4eb6d98a69004",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "subnet-02ee9d54886666a0e",
      "networkInterface": "eni-attach-08366ede1c8353527",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d2rr2ldr41dql6dh710g",
      "vmUserName": "cb-user",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-09-03T03:09:34Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0bdf2d1913588a947}}"
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
          "value": "45EB701E-E271-4E43-AC8E-016B12A5FF08"
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
          "value": "ami-08943a151bd468f4e"
        },
        {
          "key": "InstanceId",
          "value": "i-0d86a4c26af0079a1"
        },
        {
          "key": "InstanceType",
          "value": "g4dn.xlarge"
        },
        {
          "key": "KeyName",
          "value": "d2rr2ldr41dql6dh710g"
        },
        {
          "key": "LaunchTime",
          "value": "2025-09-03T03:09:33Z"
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
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.192.10},Attachment:{AttachTime:2025-09-03T03:09:33Z,AttachmentId:eni-attach-08366ede1c8353527,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0f98ae5cab23576c8,GroupName:d2rr2m5r41dql6dh7110}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:0e:e7:af:84:55,NetworkInterfaceId:eni-0997a49c0ba0fff8f,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:192.168.110.39,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.192.10},Primary:true,PrivateDnsName:null,PrivateIpAddress:192.168.110.39}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-02ee9d54886666a0e,VpcId:vpc-0ecd4eb6d98a69004}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-192-168-110-39.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "192.168.110.39"
        },
        {
          "key": "PublicIpAddress",
          "value": "43.203.192.10"
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
          "value": "{GroupId:sg-0f98ae5cab23576c8,GroupName:d2rr2m5r41dql6dh7110}"
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
          "value": "subnet-02ee9d54886666a0e"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d2rr2n5r41dql6dh7130}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0ecd4eb6d98a69004"
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
      "uid": "d2rr2n5r41dql6dh7120",
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
        "sys.uid": "d2rr2n5r41dql6dh7120"
      },
      "systemLabel": "",
      "systemMessage": null,
      "description": "a recommended multi-cloud infrastructure",
      "vm": [
        {
          "resourceType": "vm",
          "id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
          "uid": "d2rr2n5r41dql6dh7140",
          "cspResourceName": "d2rr2n5r41dql6dh7140",
          "cspResourceId": "i-0076fe4936eb3f70e",
          "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
          "subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
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
          "createdTime": "2025-09-03 03:09:49",
          "label": {
            "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2025-09-03 03:09:49",
            "sys.cspResourceId": "i-0076fe4936eb3f70e",
            "sys.cspResourceName": "d2rr2n5r41dql6dh7140",
            "sys.id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
            "sys.uid": "d2rr2n5r41dql6dh7140"
          },
          "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
          "region": {
            "Region": "ap-northeast-2",
            "Zone": "ap-northeast-2a"
          },
          "publicIP": "43.202.62.178",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "192.168.110.117",
          "privateDNS": "ip-192-168-110-117.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": "8",
          "rootDiskName": "",
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
          "specId": "aws+ap-northeast-2+g4dn.xlarge",
          "cspSpecName": "g4dn.xlarge",
          "imageId": "ami-08943a151bd468f4e",
          "cspImageName": "ami-08943a151bd468f4e",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "vpc-0ecd4eb6d98a69004",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "subnet-02ee9d54886666a0e",
          "networkInterface": "eni-attach-06f9192516cbf46ed",
          "securityGroupIds": [
            "mig-sg-02"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d2rr2ldr41dql6dh710g",
          "vmUserName": "cb-user",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-09-03T03:09:34Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0b2ed52d3a9d83895}}"
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
              "value": "D47446CB-9622-487F-B738-6DE59C71B395"
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
              "value": "ami-08943a151bd468f4e"
            },
            {
              "key": "InstanceId",
              "value": "i-0076fe4936eb3f70e"
            },
            {
              "key": "InstanceType",
              "value": "g4dn.xlarge"
            },
            {
              "key": "KeyName",
              "value": "d2rr2ldr41dql6dh710g"
            },
            {
              "key": "LaunchTime",
              "value": "2025-09-03T03:09:33Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.202.62.178},Attachment:{AttachTime:2025-09-03T03:09:33Z,AttachmentId:eni-attach-06f9192516cbf46ed,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b4d6930e1c7d30bc,GroupName:d2rr2mlr41dql6dh711g}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:f8:88:3f:3f:b1,NetworkInterfaceId:eni-06face27cb5f62cfb,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:192.168.110.117,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.202.62.178},Primary:true,PrivateDnsName:null,PrivateIpAddress:192.168.110.117}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-02ee9d54886666a0e,VpcId:vpc-0ecd4eb6d98a69004}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-192-168-110-117.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "192.168.110.117"
            },
            {
              "key": "PublicIpAddress",
              "value": "43.202.62.178"
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
              "value": "{GroupId:sg-0b4d6930e1c7d30bc,GroupName:d2rr2mlr41dql6dh711g}"
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
              "value": "subnet-02ee9d54886666a0e"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:d2rr2n5r41dql6dh7140}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0ecd4eb6d98a69004"
            }
          ]
        },
        {
          "resourceType": "vm",
          "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
          "uid": "d2rr2n5r41dql6dh7130",
          "cspResourceName": "d2rr2n5r41dql6dh7130",
          "cspResourceId": "i-0d86a4c26af0079a1",
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
          "createdTime": "2025-09-03 03:11:54",
          "label": {
            "sourceMachineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c",
            "sys.connectionName": "aws-ap-northeast-2",
            "sys.createdTime": "2025-09-03 03:11:54",
            "sys.cspResourceId": "i-0d86a4c26af0079a1",
            "sys.cspResourceName": "d2rr2n5r41dql6dh7130",
            "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
            "sys.labelType": "vm",
            "sys.manager": "cb-tumblebug",
            "sys.mciId": "mmci01",
            "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
            "sys.namespace": "mig01",
            "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
            "sys.uid": "d2rr2n5r41dql6dh7130"
          },
          "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
          "region": {
            "Region": "ap-northeast-2",
            "Zone": "ap-northeast-2a"
          },
          "publicIP": "43.203.192.10",
          "sshPort": "22",
          "publicDNS": "",
          "privateIP": "192.168.110.39",
          "privateDNS": "ip-192-168-110-39.ap-northeast-2.compute.internal",
          "rootDiskType": "gp2",
          "rootDiskSize": "8",
          "rootDiskName": "",
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
          "specId": "aws+ap-northeast-2+g4dn.xlarge",
          "cspSpecName": "g4dn.xlarge",
          "imageId": "ami-08943a151bd468f4e",
          "cspImageName": "ami-08943a151bd468f4e",
          "vNetId": "mig-vnet-01",
          "cspVNetId": "vpc-0ecd4eb6d98a69004",
          "subnetId": "mig-subnet-01",
          "cspSubnetId": "subnet-02ee9d54886666a0e",
          "networkInterface": "eni-attach-08366ede1c8353527",
          "securityGroupIds": [
            "mig-sg-01"
          ],
          "dataDiskIds": null,
          "sshKeyId": "mig-sshkey-01",
          "cspSshKeyId": "d2rr2ldr41dql6dh710g",
          "vmUserName": "cb-user",
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
              "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-09-03T03:09:34Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0bdf2d1913588a947}}"
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
              "value": "45EB701E-E271-4E43-AC8E-016B12A5FF08"
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
              "value": "ami-08943a151bd468f4e"
            },
            {
              "key": "InstanceId",
              "value": "i-0d86a4c26af0079a1"
            },
            {
              "key": "InstanceType",
              "value": "g4dn.xlarge"
            },
            {
              "key": "KeyName",
              "value": "d2rr2ldr41dql6dh710g"
            },
            {
              "key": "LaunchTime",
              "value": "2025-09-03T03:09:33Z"
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
              "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.192.10},Attachment:{AttachTime:2025-09-03T03:09:33Z,AttachmentId:eni-attach-08366ede1c8353527,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0f98ae5cab23576c8,GroupName:d2rr2m5r41dql6dh7110}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:0e:e7:af:84:55,NetworkInterfaceId:eni-0997a49c0ba0fff8f,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:192.168.110.39,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.192.10},Primary:true,PrivateDnsName:null,PrivateIpAddress:192.168.110.39}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-02ee9d54886666a0e,VpcId:vpc-0ecd4eb6d98a69004}"
            },
            {
              "key": "Placement",
              "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
            },
            {
              "key": "PrivateDnsName",
              "value": "ip-192-168-110-39.ap-northeast-2.compute.internal"
            },
            {
              "key": "PrivateIpAddress",
              "value": "192.168.110.39"
            },
            {
              "key": "PublicIpAddress",
              "value": "43.203.192.10"
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
              "value": "{GroupId:sg-0f98ae5cab23576c8,GroupName:d2rr2m5r41dql6dh7110}"
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
              "value": "subnet-02ee9d54886666a0e"
            },
            {
              "key": "Tags",
              "value": "{Key:Name,Value:d2rr2n5r41dql6dh7130}"
            },
            {
              "key": "VirtualizationType",
              "value": "hvm"
            },
            {
              "key": "VpcId",
              "value": "vpc-0ecd4eb6d98a69004"
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
  "uid": "d2rr2n5r41dql6dh7120",
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
    "sys.uid": "d2rr2n5r41dql6dh7120"
  },
  "systemLabel": "",
  "systemMessage": null,
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "uid": "d2rr2n5r41dql6dh7140",
      "cspResourceName": "d2rr2n5r41dql6dh7140",
      "cspResourceId": "i-0076fe4936eb3f70e",
      "name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
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
      "createdTime": "2025-09-03 03:09:49",
      "label": {
        "sourceMachineId": "0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2025-09-03 03:09:49",
        "sys.cspResourceId": "i-0076fe4936eb3f70e",
        "sys.cspResourceName": "d2rr2n5r41dql6dh7140",
        "sys.id": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c",
        "sys.uid": "d2rr2n5r41dql6dh7140"
      },
      "description": "a recommended virtual machine 02 for 0036e4b9-c8b4-e811-906e-000ffee02d5c",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "43.202.62.178",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.117",
      "privateDNS": "ip-192-168-110-117.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "8",
      "rootDiskName": "",
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
      "specId": "aws+ap-northeast-2+g4dn.xlarge",
      "cspSpecName": "g4dn.xlarge",
      "imageId": "ami-08943a151bd468f4e",
      "cspImageName": "ami-08943a151bd468f4e",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-0ecd4eb6d98a69004",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "subnet-02ee9d54886666a0e",
      "networkInterface": "eni-attach-06f9192516cbf46ed",
      "securityGroupIds": [
        "mig-sg-02"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d2rr2ldr41dql6dh710g",
      "vmUserName": "cb-user",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-09-03T03:09:34Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0b2ed52d3a9d83895}}"
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
          "value": "D47446CB-9622-487F-B738-6DE59C71B395"
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
          "value": "ami-08943a151bd468f4e"
        },
        {
          "key": "InstanceId",
          "value": "i-0076fe4936eb3f70e"
        },
        {
          "key": "InstanceType",
          "value": "g4dn.xlarge"
        },
        {
          "key": "KeyName",
          "value": "d2rr2ldr41dql6dh710g"
        },
        {
          "key": "LaunchTime",
          "value": "2025-09-03T03:09:33Z"
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
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.202.62.178},Attachment:{AttachTime:2025-09-03T03:09:33Z,AttachmentId:eni-attach-06f9192516cbf46ed,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0b4d6930e1c7d30bc,GroupName:d2rr2mlr41dql6dh711g}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:f8:88:3f:3f:b1,NetworkInterfaceId:eni-06face27cb5f62cfb,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:192.168.110.117,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.202.62.178},Primary:true,PrivateDnsName:null,PrivateIpAddress:192.168.110.117}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-02ee9d54886666a0e,VpcId:vpc-0ecd4eb6d98a69004}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-192-168-110-117.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "192.168.110.117"
        },
        {
          "key": "PublicIpAddress",
          "value": "43.202.62.178"
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
          "value": "{GroupId:sg-0b4d6930e1c7d30bc,GroupName:d2rr2mlr41dql6dh711g}"
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
          "value": "subnet-02ee9d54886666a0e"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d2rr2n5r41dql6dh7140}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0ecd4eb6d98a69004"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "uid": "d2rr2n5r41dql6dh7130",
      "cspResourceName": "d2rr2n5r41dql6dh7130",
      "cspResourceId": "i-0d86a4c26af0079a1",
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
      "createdTime": "2025-09-03 03:11:54",
      "label": {
        "sourceMachineId": "00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2025-09-03 03:11:54",
        "sys.cspResourceId": "i-0d86a4c26af0079a1",
        "sys.cspResourceName": "d2rr2n5r41dql6dh7130",
        "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.uid": "d2rr2n5r41dql6dh7130"
      },
      "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "43.203.192.10",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.39",
      "privateDNS": "ip-192-168-110-39.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "8",
      "rootDiskName": "",
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
      "specId": "aws+ap-northeast-2+g4dn.xlarge",
      "cspSpecName": "g4dn.xlarge",
      "imageId": "ami-08943a151bd468f4e",
      "cspImageName": "ami-08943a151bd468f4e",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "vpc-0ecd4eb6d98a69004",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "subnet-02ee9d54886666a0e",
      "networkInterface": "eni-attach-08366ede1c8353527",
      "securityGroupIds": [
        "mig-sg-01"
      ],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d2rr2ldr41dql6dh710g",
      "vmUserName": "cb-user",
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
          "value": "{DeviceName:/dev/sda1,Ebs:{AttachTime:2025-09-03T03:09:34Z,DeleteOnTermination:true,Status:attached,VolumeId:vol-0bdf2d1913588a947}}"
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
          "value": "45EB701E-E271-4E43-AC8E-016B12A5FF08"
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
          "value": "ami-08943a151bd468f4e"
        },
        {
          "key": "InstanceId",
          "value": "i-0d86a4c26af0079a1"
        },
        {
          "key": "InstanceType",
          "value": "g4dn.xlarge"
        },
        {
          "key": "KeyName",
          "value": "d2rr2ldr41dql6dh710g"
        },
        {
          "key": "LaunchTime",
          "value": "2025-09-03T03:09:33Z"
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
          "value": "{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.192.10},Attachment:{AttachTime:2025-09-03T03:09:33Z,AttachmentId:eni-attach-08366ede1c8353527,DeleteOnTermination:true,DeviceIndex:0,NetworkCardIndex:0,Status:attached},Description:,Groups:[{GroupId:sg-0f98ae5cab23576c8,GroupName:d2rr2m5r41dql6dh7110}],InterfaceType:interface,Ipv6Addresses:null,MacAddress:02:0e:e7:af:84:55,NetworkInterfaceId:eni-0997a49c0ba0fff8f,OwnerId:635484366616,PrivateDnsName:null,PrivateIpAddress:192.168.110.39,PrivateIpAddresses:[{Association:{CarrierIp:null,IpOwnerId:amazon,PublicDnsName:,PublicIp:43.203.192.10},Primary:true,PrivateDnsName:null,PrivateIpAddress:192.168.110.39}],SourceDestCheck:true,Status:in-use,SubnetId:subnet-02ee9d54886666a0e,VpcId:vpc-0ecd4eb6d98a69004}"
        },
        {
          "key": "Placement",
          "value": "{Affinity:null,AvailabilityZone:ap-northeast-2a,GroupName:,HostId:null,HostResourceGroupArn:null,PartitionNumber:null,SpreadDomain:null,Tenancy:default}"
        },
        {
          "key": "PrivateDnsName",
          "value": "ip-192-168-110-39.ap-northeast-2.compute.internal"
        },
        {
          "key": "PrivateIpAddress",
          "value": "192.168.110.39"
        },
        {
          "key": "PublicIpAddress",
          "value": "43.203.192.10"
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
          "value": "{GroupId:sg-0f98ae5cab23576c8,GroupName:d2rr2m5r41dql6dh7110}"
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
          "value": "subnet-02ee9d54886666a0e"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:d2rr2n5r41dql6dh7130}"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        },
        {
          "key": "VpcId",
          "value": "vpc-0ecd4eb6d98a69004"
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

