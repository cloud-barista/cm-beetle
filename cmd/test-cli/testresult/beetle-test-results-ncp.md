# CM-Beetle test results for NCP

> [!NOTE]
> This document presents comprehensive test results for CM-Beetle integration with NCP cloud infrastructure.

## Environment and scenario

### Environment

- CM-Beetle: v0.4.0 (220af0b)
- cm-model: v0.0.11
- CB-Tumblebug: v0.11.3
- CB-Spider: v0.11.1
- CB-MapUI: v0.11.4
- Target CSP: NCP
- Target Region: kr
- CM-Beetle URL: http://localhost:8056
- Namespace: mig01
- Test CLI: Custom automated testing tool
- Test Date: August 13, 2025
- Test Time: 21:52:28 KST
- Test Execution: 2025-08-13 21:52:28 KST

### Scenario

1. Recommend a target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. List all MCIs via Beetle
1. List MCI IDs via Beetle
1. Get specific MCI details via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Test result for NCP

### Test Results Summary

| Test | Endpoint                                          | Status      | Duration  | Details |
| ---- | ------------------------------------------------- | ----------- | --------- | ------- |
| 1    | `POST /beetle/recommendation/mci`                 | ✅ **PASS** | 1.529s    | Success |
| 2    | `POST /beetle/migration/ns/mig01/mci`             | ✅ **PASS** | 3m19.045s | Success |
| 3    | `GET /beetle/migration/ns/mig01/mci`              | ✅ **PASS** | 133ms     | Success |
| 4    | `GET /beetle/migration/ns/mig01/mci?option=id`    | ✅ **PASS** | 58ms      | Success |
| 5    | `GET /beetle/migration/ns/mig01/mci/{{mciId}}`    | ✅ **PASS** | 139ms     | Success |
| 6    | `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}` | ✅ **PASS** | 13m4.461s | Success |

**Overall Result**: 6/6 tests passed ✅

**Total Duration**: 16m55.509589159s

_Test executed on August 13, 2025 at 21:52:28 KST (2025-08-13 21:52:28 KST) using CM-Beetle automated test CLI_

### Recommend a target model for computing infra

> [!Note] > `desiredCsp` and `desiredRegion` are required in the request body.

- API: `POST /beetle/recommendation/mci`
- Request body:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "desiredCspAndRegionPair": {
    "csp": "ncp",
    "region": "kr"
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
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
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
            "ipv4CidrBlocks": ["172.29.0.102/24", "172.29.0.200/32"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b003/64"],
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
            "ipv4CidrBlocks": ["192.168.110.200/32"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b004/64"],
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
            "ipv4CidrBlocks": ["172.20.0.1/16"],
            "mtu": 1500,
            "state": "down"
          },
          {
            "name": "br-f67138586d47",
            "macAddress": "02:42:6e:92:df:03",
            "ipv4CidrBlocks": ["172.19.0.1/16"],
            "mtu": 1500,
            "state": "down"
          },
          {
            "name": "br-068801a3f047",
            "macAddress": "02:42:cc:24:25:30",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:ccff:fe24:2530/64"],
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
            "ipv4CidrBlocks": ["10.1.0.106/24"],
            "ipv6CidrBlocks": ["fe80::f816:3eff:fe9d:89c5/64"],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "genev_sys_6081",
            "macAddress": "de:4b:8c:92:4c:db",
            "ipv6CidrBlocks": ["fe80::2852:51ff:fe36:258b/64"],
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
            "ipv4CidrBlocks": ["192.168.110.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b004/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap0481d752-40",
            "macAddress": "6a:2a:78:65:42:32",
            "ipv6CidrBlocks": ["fe80::682a:78ff:fe65:4232/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap935cb764-41",
            "macAddress": "fe:16:3e:4c:39:2b",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fe4c:392b/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap19d6d4d9-a4",
            "macAddress": "fe:16:3e:d5:6f:85",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fed5:6f85/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap7422e216-ff",
            "macAddress": "fe:16:3e:4d:31:9e",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fe4d:319e/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapa53b173c-e4",
            "macAddress": "fe:16:3e:52:91:4b",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fe52:914b/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapabb5f299-74",
            "macAddress": "fe:16:3e:46:9b:72",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fe46:9b72/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapf6929430-67",
            "macAddress": "fe:16:3e:3e:15:10",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fe3e:1510/64"],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "tap3968711d-8a",
            "macAddress": "fe:16:3e:65:ad:39",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fe65:ad39/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap49d44128-d0",
            "macAddress": "fe:16:3e:1e:c7:fc",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fe1e:c7fc/64"],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "tap708d34b6-e0",
            "macAddress": "fe:16:3e:19:8c:71",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fe19:8c71/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap1479d90f-c0",
            "macAddress": "7a:0f:53:ad:50:84",
            "ipv6CidrBlocks": ["fe80::780f:53ff:fead:5084/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap1a03c4f4-e8",
            "macAddress": "fa:16:3e:c9:ea:1c",
            "ipv4CidrBlocks": ["10.254.0.27/28", "10.254.0.3/28"],
            "ipv6CidrBlocks": ["fe80::f816:3eff:fec9:ea1c/64"],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "veth0b8a5f4",
            "macAddress": "be:22:36:27:01:d2",
            "ipv6CidrBlocks": ["fe80::bc22:36ff:fe27:1d2/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "veth87e839e",
            "macAddress": "32:de:9f:d7:cd:24",
            "ipv6CidrBlocks": ["fe80::38f0:78ff:fef7:358/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "veth089f03a",
            "macAddress": "2a:8f:e3:66:fd:99",
            "ipv6CidrBlocks": ["fe80::5c87:18ff:fe73:d0dd/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapaf1a281f-c0",
            "macAddress": "32:3c:e7:79:ee:ef",
            "ipv6CidrBlocks": ["fe80::303c:e7ff:fe79:eeef/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tap0e0c519d-d0",
            "macAddress": "fe:16:3e:8a:c2:22",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fe8a:c222/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "tapd801f01d-d6",
            "macAddress": "fe:16:3e:09:e9:f5",
            "ipv6CidrBlocks": ["fe80::fc16:3eff:fe09:e9f5/64"],
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

- Response body:

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
        "connectionName": "ncpvpc-kr",
        "specId": "s2-g3",
        "imageId": "23789321",
        "vNetId": "mig-vnet-01",
        "subnetId": "mig-subnet-01",
        "securityGroupIds": ["mig-sg-01"],
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
    "connectionName": "ncpvpc-kr",
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
    "connectionName": "ncpvpc-kr",
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
      "id": "ncpvpc+kr+s2-g3",
      "cspSpecName": "s2-g3",
      "name": "ncpvpc+kr+s2-g3",
      "namespace": "system",
      "connectionName": "ncpvpc-kr",
      "providerName": "ncpvpc",
      "regionName": "kr",
      "infraType": "vm",
      "architecture": "x86_64",
      "vCPU": 2,
      "memoryGiB": 8,
      "diskSizeGB": -1,
      "costPerHour": 0.0848,
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
      "rootDiskType": "default",
      "rootDiskSize": "-1",
      "systemLabel": "from-assets",
      "details": [
        {
          "key": "ServerSpecCode",
          "value": "s2-g3"
        },
        {
          "key": "GenerationCode",
          "value": "G3"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize",
          "value": "8589934592"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "CpuArchitectureType",
          "value": "{code:X86_64,codeName:x86 64bit}"
        },
        {
          "key": "BlockStorageMaxCount",
          "value": "20"
        },
        {
          "key": "BlockStorageMaxIops",
          "value": "4725"
        },
        {
          "key": "BlockStorageMaxThroughput",
          "value": "81"
        },
        {
          "key": "NetworkPerformance",
          "value": "1000000000"
        },
        {
          "key": "NetworkInterfaceMaxCount",
          "value": "3"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.STAND.C002.M008.G003"
        },
        {
          "key": "ServerSpecDescription",
          "value": "vCPU 2EA, Memory 8GB"
        },
        {
          "key": "CorrespondingImageIds",
          "value": "107029409,104630229,100524418,25495367,23221307,23221289,23214590,19463675,17552318,16946033"
        }
      ]
    }
  ],
  "targetVmOsImageList": [
    {
      "namespace": "system",
      "providerName": "ncpvpc",
      "cspImageName": "23789321",
      "regionList": ["kr"],
      "id": "ncpvpc+23789321",
      "name": "ncpvpc+23789321",
      "connectionName": "ncpvpc-kr",
      "fetchedTime": "2025.08.12 07:58:40 Tue",
      "isGPUImage": true,
      "osType": "Ubuntu 22.04",
      "osArchitecture": "x86_64",
      "osPlatform": "Linux/UNIX",
      "osDistribution": "ubuntu-22.04-gpu (Hypervisor:KVM)",
      "osDiskType": "Common BlockStorage 1",
      "osDiskSizeGB": 107,
      "imageStatus": "Available",
      "details": [
        {
          "key": "ServerImageNo",
          "value": "23789321"
        },
        {
          "key": "ServerImageName",
          "value": "ubuntu-22.04-gpu"
        },
        {
          "key": "ServerImageDescription",
          "value": "kernel version : 5.15.0-140-generic"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerImageType",
          "value": "{code:NCP,codeName:NCP 서버이미지}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "CpuArchitectureType",
          "value": "{code:X86_64,codeName:x86 64bit}"
        },
        {
          "key": "OsCategoryType",
          "value": "{code:LINUX,codeName:LINUX}"
        },
        {
          "key": "OsType",
          "value": "{code:UBUNTU,codeName:UBUNTU}"
        },
        {
          "key": "ServerImageStatus",
          "value": "{code:CREAT,codeName:NSI CREATED state}"
        },
        {
          "key": "ServerImageOperation",
          "value": "{code:NULL,codeName:NSI NULL OP}"
        },
        {
          "key": "ServerImageStatusName",
          "value": "created"
        },
        {
          "key": "CreateDate",
          "value": "2024-04-25T18:39:01+0900"
        },
        {
          "key": "ShareStatus",
          "value": "{code:NULL,codeName:NSI Share NULL State}"
        },
        {
          "key": "BlockStorageMappingList",
          "value": "{order:0,blockStorageSnapshotInstanceNo:23752581,blockStorageSnapshotName:snapshot of ubuntu-22.04-gpu,blockStorageSize:107374182400,blockStorageVolumeType:{code:CB1,codeName:Common BlockStorage 1},iops:300,throughput:104857600,isEncryptedVolume:false}"
        }
      ]
    }
  ],
  "targetSecurityGroupList": [
    {
      "name": "mig-sg-01",
      "connectionName": "ncpvpc-kr",
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

### Migrate the computing infra as defined in the target model

- API: `POST /beetle/migration/ns/mig01/mci`
- nsId: `mig01`
- Request body: **same as the response from the previous step**
- Response body:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "d2e8lhkv38occbiu5oig",
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
    "sys.uid": "d2e8lhkv38occbiu5oig"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "uid": "d2e8lhkv38occbiu5ojg",
      "cspResourceName": "d2e8lhkv38occbiu5ojg",
      "cspResourceId": "108085622",
      "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-08-13 12:55:59",
      "label": {
        "sys.connectionName": "ncpvpc-kr",
        "sys.createdTime": "2025-08-13 12:55:59",
        "sys.cspResourceId": "108085622",
        "sys.cspResourceName": "d2e8lhkv38occbiu5ojg",
        "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.uid": "d2e8lhkv38occbiu5ojg"
      },
      "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "211.188.60.25",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.6",
      "privateDNS": "",
      "rootDiskType": "SSD",
      "rootDiskSize": "100",
      "rootDiskName": "",
      "connectionName": "ncpvpc-kr",
      "connectionConfig": {
        "configName": "ncpvpc-kr",
        "providerName": "ncpvpc",
        "driverName": "ncpvpc-driver-v1.0.so",
        "credentialName": "ncpvpc",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncpvpc-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": ["KR-1", "KR-2"]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "s2-g3",
      "cspSpecName": "s2-g3",
      "imageId": "23789321",
      "cspImageName": "23789321",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "118605",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "250374",
      "networkInterface": "eth0",
      "securityGroupIds": ["mig-sg-01"],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d2e8lbsv38occbiu5ohg",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "108085622"
        },
        {
          "key": "ServerName",
          "value": "d2e8lhkv38occbiu5ojg"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize",
          "value": "8589934592"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "d2e8lbsv38occbiu5ohg"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2025-08-13T21:53:45+0900"
        },
        {
          "key": "Uptime",
          "value": "2025-08-13T21:55:51+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.STAND.C002.M008.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "118605"
        },
        {
          "key": "SubnetNo",
          "value": "250374"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "4892970"
        },
        {
          "key": "InitScriptNo",
          "value": "134609"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:STAND,codeName:Standard}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23789321"
        },
        {
          "key": "ServerSpecCode",
          "value": "s2-g3"
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

### Get a list of MCIs

- API: `GET /beetle/migration/ns/mig01/mci`
- nsId: `mig01`
- Request body: None
- Response body:

```json
{
  "mci": [
    {
      "resourceType": "mci",
      "id": "mmci01",
      "uid": "d2e8lhkv38occbiu5oig",
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
          "uid": "d2e8lhkv38occbiu5oig",
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

### Get a list of MCI IDs

- API: `GET /beetle/migration/ns/mig01/mci?option=id`
- nsId: `mig01`
- Request body: None
- Response body:

```json
{
  "idList": ["mmci01"]
}
```

### Get a specific MCI

- API: `GET /beetle/migration/ns/mig01/mci/{{mciId}}`
- nsId: `mig01`
- Request body: None
- Response body:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "d2e8lhkv38occbiu5oig",
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
    "sys.uid": "d2e8lhkv38occbiu5oig"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "a recommended multi-cloud infrastructure",
  "vm": [
    {
      "resourceType": "vm",
      "id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "uid": "d2e8lhkv38occbiu5ojg",
      "cspResourceName": "d2e8lhkv38occbiu5ojg",
      "cspResourceId": "108085622",
      "name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
      "subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "status": "Running",
      "targetStatus": "None",
      "targetAction": "None",
      "monAgentStatus": "notInstalled",
      "networkAgentStatus": "notInstalled",
      "systemMessage": "",
      "createdTime": "2025-08-13 12:55:59",
      "label": {
        "sys.connectionName": "ncpvpc-kr",
        "sys.createdTime": "2025-08-13 12:55:59",
        "sys.cspResourceId": "108085622",
        "sys.cspResourceName": "d2e8lhkv38occbiu5ojg",
        "sys.id": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "migrated-00a9f3d4-74b6-e811-906e-000ffee02d5c",
        "sys.uid": "d2e8lhkv38occbiu5ojg"
      },
      "description": "a recommended virtual machine 01 for 00a9f3d4-74b6-e811-906e-000ffee02d5c",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "211.188.60.25",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "192.168.110.6",
      "privateDNS": "",
      "rootDiskType": "SSD",
      "rootDiskSize": "100",
      "rootDiskName": "",
      "connectionName": "ncpvpc-kr",
      "connectionConfig": {
        "configName": "ncpvpc-kr",
        "providerName": "ncpvpc",
        "driverName": "ncpvpc-driver-v1.0.so",
        "credentialName": "ncpvpc",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ncpvpc-kr",
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionDetail": {
          "regionId": "KR",
          "regionName": "kr",
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "zones": ["KR-1", "KR-2"]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "s2-g3",
      "cspSpecName": "s2-g3",
      "imageId": "23789321",
      "cspImageName": "23789321",
      "vNetId": "mig-vnet-01",
      "cspVNetId": "118605",
      "subnetId": "mig-subnet-01",
      "cspSubnetId": "250374",
      "networkInterface": "eth0",
      "securityGroupIds": ["mig-sg-01"],
      "dataDiskIds": null,
      "sshKeyId": "mig-sshkey-01",
      "cspSshKeyId": "d2e8lbsv38occbiu5ohg",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "ServerInstanceNo",
          "value": "108085622"
        },
        {
          "key": "ServerName",
          "value": "d2e8lhkv38occbiu5ojg"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize",
          "value": "8589934592"
        },
        {
          "key": "PlatformType",
          "value": "{code:UBD64,codeName:Ubuntu Desktop 64 Bit}"
        },
        {
          "key": "LoginKeyName",
          "value": "d2e8lbsv38occbiu5ohg"
        },
        {
          "key": "ServerInstanceStatus",
          "value": "{code:RUN,codeName:서버 RUN 상태}"
        },
        {
          "key": "ServerInstanceOperation",
          "value": "{code:NULL,codeName:서버 NULL OP}"
        },
        {
          "key": "ServerInstanceStatusName",
          "value": "running"
        },
        {
          "key": "CreateDate",
          "value": "2025-08-13T21:53:45+0900"
        },
        {
          "key": "Uptime",
          "value": "2025-08-13T21:55:51+0900"
        },
        {
          "key": "ServerImageProductCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.G003"
        },
        {
          "key": "ServerProductCode",
          "value": "SVR.VSVR.STAND.C002.M008.G003"
        },
        {
          "key": "IsProtectServerTermination",
          "value": "false"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "RegionCode",
          "value": "KR"
        },
        {
          "key": "VpcNo",
          "value": "118605"
        },
        {
          "key": "SubnetNo",
          "value": "250374"
        },
        {
          "key": "NetworkInterfaceNoList",
          "value": "4892970"
        },
        {
          "key": "InitScriptNo",
          "value": "134609"
        },
        {
          "key": "ServerInstanceType",
          "value": "{code:STAND,codeName:Standard}"
        },
        {
          "key": "BaseBlockStorageDiskType",
          "value": "{code:NET,codeName:네트웍 스토리지}"
        },
        {
          "key": "BaseBlockStorageDiskDetailType",
          "value": "{code:SSD,codeName:SSD}"
        },
        {
          "key": "HypervisorType",
          "value": "{code:KVM,codeName:KVM}"
        },
        {
          "key": "ServerImageNo",
          "value": "23789321"
        },
        {
          "key": "ServerSpecCode",
          "value": "s2-g3"
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

### Delete the migrated computing infra

- API: `DELETE /beetle/migration/ns/mig01/mci/{{mciId}}`
- nsId: `mig01`
- Request body: None
- Response body:

```json
{
  "success": true,
  "text": "Successfully deleted the infrastructure and resources (nsId: mig01, infraId: mmci01)"
}
```
