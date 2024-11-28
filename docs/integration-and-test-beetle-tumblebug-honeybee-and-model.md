# Integration and Test: Beetle, Tumblebug, Honeybee and model

## Environment and scenario

### Envrionment

- Beetle v0.2.5
- cm-model v0.0.3 (Damselfly v0.2.1)
- Honeybee v0.2.10
- Tumblebug v0.10.3 (Spider v0.10.0, CB-MapUI v0.10.0)

### Scenario

1. Get a source group list via Honeybee
1. Get the refined source computing infra info via Honeybee

- Refined source computing infra info = on-premise model (a.k.a computing infra source model)
- **Used 2 servers info** as dicussed (i.e., web, nfs)

1. Recommend a target model for computing infra via Beetle
1. Migrate the computing infra as defined in the target model via Beetle
1. Delete the migrated computing infra via Beetle

> [!NOTE]
> Some long request/response bodies are in the collapsible section for better readability.

## Honeybee section

> [!Note]
> The Honeybee has been providing the always-running servers. The server has been used.

### Get a list of source group

- API: `GET /source_group`
- Request body: None
- Response body:

```json
{
  "source_group": [
    {
      "id": "db652288-047b-480b-ac86-3ef7ed57f68e",
      "name": "cm-test",
      "description": "Cloud-Migrator Test Source Group",
      "connection_info_status_count": {
        "count_connection_success": 5,
        "count_connection_failed": 0,
        "count_agent_success": 5,
        "count_agent_failed": 0,
        "connection_info_total": 5
      }
    },
    {
      "id": "ddcfa917-17e0-4718-a878-f1e99f97ed6d",
      "name": "migration-test",
      "description": "Migration Test Group",
      "connection_info_status_count": {
        "count_connection_success": 2,
        "count_connection_failed": 0,
        "count_agent_success": 2,
        "count_agent_failed": 0,
        "connection_info_total": 2
      }
    }
  ],
  "connection_info_status_count": {
    "count_connection_success": 7,
    "count_connection_failed": 0,
    "count_agent_success": 7,
    "count_agent_failed": 0,
    "connection_info_total": 7
  }
}
```

### Get the refined computing infra info

- API: `GET /source_group/{sgId}/infra/refined`
- sgId: `db652288-047b-480b-ac86-3ef7ed57f68e`
- Request body: None
- Response body:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "onpremiseInfraModel": {
    "network": {
      "ipv4Networks": [
        "10.0.0.0/24",
        "192.168.110.0/24",
        "172.17.0.0/16",
        "172.16.0.128/32",
        "172.16.0.80/32",
        "172.29.0.0/24",
        "172.18.0.0/16",
        "172.21.0.0/16",
        "172.22.0.0/16"
      ],
      "ipv6Networks": ["fe80::/64"]
    },
    "servers": [
      {
        "hostname": "cm-docker",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 4,
          "cores": 1,
          "threads": 1,
          "maxSpeed": 2.294,
          "vendor": "GenuineIntel",
          "model": "Intel Xeon Processor (Skylake, IBRS)"
        },
        "memory": {
          "type": "RAM",
          "totalSize": 8,
          "available": 7
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 57,
          "available": 55,
          "used": 2
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536
          },
          {
            "name": "eth0",
            "macAddress": "fa:16:3e:ce:95:d6",
            "ipv4CidrBlocks": ["10.0.0.201/24"],
            "ipv6CidrBlocks": ["fe80::f816:3eff:fece:95d6/64"],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "eth1",
            "macAddress": "fa:16:3e:b2:bd:57",
            "ipv4CidrBlocks": ["192.168.110.227/24"],
            "ipv6CidrBlocks": ["fe80::f816:3eff:feb2:bd57/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "docker0",
            "macAddress": "02:42:98:a5:74:2b",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "1.0.0.1",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "1.0.0.1",
            "interface": "eth0",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "1.1.1.1",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "1.1.1.1",
            "interface": "eth0",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "10.0.0.0",
            "interface": "eth0",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "10.0.0.2",
            "interface": "eth0",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "169.254.169.254",
            "interface": "eth0",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "169.254.169.254",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.226",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.254",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "metric": 100,
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "k8s-master.example.com",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 4,
          "cores": 1,
          "threads": 1,
          "maxSpeed": 2.294,
          "vendor": "GenuineIntel",
          "model": "Intel Xeon Processor (Skylake, IBRS)"
        },
        "memory": {
          "type": "RAM",
          "totalSize": 8,
          "available": 6,
          "used": 1
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 57,
          "available": 42,
          "used": 15
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536
          },
          {
            "name": "eth0",
            "macAddress": "fa:16:3e:fb:3b:98",
            "ipv4CidrBlocks": ["10.0.0.146/24"],
            "ipv6CidrBlocks": ["fe80::f816:3eff:fefb:3b98/64"],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "eth1",
            "macAddress": "fa:16:3e:c9:28:82",
            "ipv4CidrBlocks": ["192.168.110.213/24"],
            "ipv6CidrBlocks": ["fe80::f816:3eff:fec9:2882/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "cali7827be5483d",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali8c3392366fd",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali352921d3328",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "calia99136011ef",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali0a47f80037a",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "tunl0",
            "ipv4CidrBlocks": ["172.16.0.128/32"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali7d7a5dc7de4",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali52c7dca0137",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali54f0ce29898",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "calic43e969f430",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "calibd2ccf7ccb0",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "10.0.0.0",
            "interface": "eth0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.64",
            "interface": "tunl0",
            "protocol": "bird",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.128",
            "interface": "*"
          },
          {
            "destination": "172.16.0.179",
            "interface": "cali7827be5483d",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.180",
            "interface": "cali8c3392366fd",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.181",
            "interface": "cali352921d3328",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.182",
            "interface": "calia99136011ef",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.183",
            "interface": "cali0a47f80037a",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.184",
            "interface": "cali7d7a5dc7de4",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.185",
            "interface": "cali52c7dca0137",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.186",
            "interface": "cali54f0ce29898",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.187",
            "interface": "calic43e969f430",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.188",
            "interface": "calibd2ccf7ccb0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "k8s-worker.example.com",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 4,
          "cores": 1,
          "threads": 1,
          "maxSpeed": 2.294,
          "vendor": "GenuineIntel",
          "model": "Intel Xeon Processor (Skylake, IBRS)"
        },
        "memory": {
          "type": "RAM",
          "totalSize": 8,
          "available": 4,
          "used": 3
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 57,
          "available": 47,
          "used": 10
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536
          },
          {
            "name": "eth0",
            "macAddress": "fa:16:3e:03:5a:2f",
            "ipv4CidrBlocks": ["10.0.0.24/24"],
            "ipv6CidrBlocks": ["fe80::f816:3eff:fe03:5a2f/64"],
            "mtu": 1442,
            "state": "up"
          },
          {
            "name": "eth1",
            "macAddress": "fa:16:3e:fc:18:6b",
            "ipv4CidrBlocks": ["192.168.110.212/24"],
            "ipv6CidrBlocks": ["fe80::f816:3eff:fefc:186b/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "cali2cf4bfaf835",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "tunl0",
            "ipv4CidrBlocks": ["172.16.0.80/32"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "calib8b0370b004",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali9613c8416ed",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "calib68ee0153a2",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali45a163c7b6d",
            "macAddress": "26:a1:67:af:b4:85",
            "ipv6CidrBlocks": ["fe80::24a1:67ff:feaf:b485/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "calif7e05d18904",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali5168c2b7230",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali6912d996d95",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "califaaa3bfab3c",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "calie50055bacbc",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "calicc0f8f9a0e0",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "cali9221d5a29df",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          },
          {
            "name": "calic5ed947e64b",
            "macAddress": "ee:ee:ee:ee:ee:ee",
            "ipv6CidrBlocks": ["fe80::ecee:eeff:feee:eeee/64"],
            "mtu": 1422,
            "state": "up"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "10.0.0.0",
            "interface": "eth0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.64",
            "interface": "califaaa3bfab3c",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.64",
            "interface": "*"
          },
          {
            "destination": "172.16.0.65",
            "interface": "calie50055bacbc",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.66",
            "interface": "calicc0f8f9a0e0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.67",
            "interface": "cali9221d5a29df",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.69",
            "interface": "calic5ed947e64b",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.120",
            "interface": "cali2cf4bfaf835",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.121",
            "interface": "calib8b0370b004",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.122",
            "interface": "cali9613c8416ed",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.123",
            "interface": "calib68ee0153a2",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.124",
            "interface": "cali45a163c7b6d",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.125",
            "interface": "calif7e05d18904",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.126",
            "interface": "cali5168c2b7230",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.127",
            "interface": "cali6912d996d95",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.16.0.128",
            "interface": "tunl0",
            "protocol": "bird",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eth1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "cm-nfs",
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
          "totalSize": 255,
          "available": 248,
          "used": 7
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1093,
          "available": 998,
          "used": 39
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
            "mtu": 65536
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:52:fa:e8",
            "mtu": 1500
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b0:03",
            "ipv4CidrBlocks": ["172.29.0.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b003/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:52:fa:e9",
            "mtu": 1500
          },
          {
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:47:70:f0",
            "mtu": 1500
          },
          {
            "name": "enp26s0f0",
            "macAddress": "b4:96:91:53:01:54",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b0:04",
            "ipv4CidrBlocks": ["192.168.110.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b004/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp26s0f1",
            "macAddress": "b4:96:91:53:01:55",
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
            "name": "docker0",
            "macAddress": "02:42:e3:e3:d6:97",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          },
          {
            "name": "br-58dab7383fc5",
            "macAddress": "02:42:03:69:20:4f",
            "ipv4CidrBlocks": ["172.18.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:3ff:fe69:204f/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "veth3308dac",
            "macAddress": "56:ed:7d:2a:96:7c",
            "ipv6CidrBlocks": ["fe80::54ed:7dff:fe2a:967c/64"],
            "mtu": 1500
          },
          {
            "name": "veth1d4c5c1",
            "macAddress": "76:8f:5d:78:d7:0b",
            "ipv6CidrBlocks": ["fe80::748f:5dff:fe78:d70b/64"],
            "mtu": 1500
          },
          {
            "name": "vethcdc65b0",
            "macAddress": "b6:cf:0c:4a:a0:18",
            "ipv6CidrBlocks": ["fe80::b4cf:cff:fe4a:a018/64"],
            "mtu": 1500
          },
          {
            "name": "br-b236d7c30f50",
            "macAddress": "02:42:19:72:d8:1f",
            "ipv4CidrBlocks": ["172.21.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:19ff:fe72:d81f/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "br-0d09f6c0e547",
            "macAddress": "02:42:d7:47:50:92",
            "ipv4CidrBlocks": ["172.22.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:d7ff:fe47:5092/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "vetha41e00b",
            "macAddress": "8e:0e:4b:1b:b4:f7",
            "ipv6CidrBlocks": ["fe80::8c0e:4bff:fe1b:b4f7/64"],
            "mtu": 1500
          },
          {
            "name": "veth908225d",
            "macAddress": "da:0d:32:ea:45:07",
            "ipv6CidrBlocks": ["fe80::d80d:32ff:feea:4507/64"],
            "mtu": 1500
          },
          {
            "name": "veth4a8cc8d",
            "macAddress": "f2:a4:00:66:45:85",
            "ipv6CidrBlocks": ["fe80::f0a4:ff:fe66:4585/64"],
            "mtu": 1500
          },
          {
            "name": "veth3c415c0",
            "macAddress": "be:13:49:9b:84:e2",
            "ipv6CidrBlocks": ["fe80::bc13:49ff:fe9b:84e2/64"],
            "mtu": 1500
          },
          {
            "name": "veth442f928",
            "macAddress": "be:66:6f:73:e1:63",
            "ipv6CidrBlocks": ["fe80::bc66:6fff:fe73:e163/64"],
            "mtu": 1500
          },
          {
            "name": "veth644fb0c",
            "macAddress": "46:d7:01:9f:6f:3a",
            "ipv6CidrBlocks": ["fe80::44d7:1ff:fe9f:6f3a/64"],
            "mtu": 1500
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.18.0.0",
            "interface": "br-58dab7383fc5",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.21.0.0",
            "interface": "br-b236d7c30f50",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.22.0.0",
            "interface": "br-0d09f6c0e547",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "cm-web",
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
          "totalSize": 255,
          "available": 251,
          "used": 4
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1312,
          "available": 1229,
          "used": 16
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b1:1b",
            "ipv4CidrBlocks": ["172.29.0.103/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b11b/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:53:01:58",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b1:1c",
            "ipv4CidrBlocks": ["192.168.110.103/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b11c/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:53:01:59",
            "mtu": 1500
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
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:55:23:8c",
            "mtu": 1500
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
            "name": "docker0",
            "macAddress": "02:42:55:3d:6b:90",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
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

**Used 2 servers info (i.e., nfs, web) as follows**

<details>
  <summary> <ins>Click to see the information </ins> </summary>
  
```json
{
  "onpremiseInfraModel": {
    "network": {
      "ipv4Networks": [
        "10.0.0.0/24",
        "192.168.110.0/24",
        "172.17.0.0/16",
        "172.16.0.128/32",
        "172.16.0.80/32",
        "172.29.0.0/24",
        "172.18.0.0/16",
        "172.21.0.0/16",
        "172.22.0.0/16"
      ],
      "ipv6Networks": [
        "fe80::/64"
      ]
    },
    "servers": [
      {
        "hostname": "cm-nfs",
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
          "totalSize": 255,
          "available": 248,
          "used": 7
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1093,
          "available": 998,
          "used": 39
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
            "mtu": 65536
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:52:fa:e8",
            "mtu": 1500
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b0:03",
            "ipv4CidrBlocks": [
              "172.29.0.102/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::a6bf:1ff:fe5a:b003/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:52:fa:e9",
            "mtu": 1500
          },
          {
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:47:70:f0",
            "mtu": 1500
          },
          {
            "name": "enp26s0f0",
            "macAddress": "b4:96:91:53:01:54",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
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
            "name": "enp26s0f1",
            "macAddress": "b4:96:91:53:01:55",
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
            "name": "docker0",
            "macAddress": "02:42:e3:e3:d6:97",
            "ipv4CidrBlocks": [
              "172.17.0.1/16"
            ],
            "mtu": 1500,
            "state": "down"
          },
          {
            "name": "br-58dab7383fc5",
            "macAddress": "02:42:03:69:20:4f",
            "ipv4CidrBlocks": [
              "172.18.0.1/16"
            ],
            "ipv6CidrBlocks": [
              "fe80::42:3ff:fe69:204f/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "veth3308dac",
            "macAddress": "56:ed:7d:2a:96:7c",
            "ipv6CidrBlocks": [
              "fe80::54ed:7dff:fe2a:967c/64"
            ],
            "mtu": 1500
          },
          {
            "name": "veth1d4c5c1",
            "macAddress": "76:8f:5d:78:d7:0b",
            "ipv6CidrBlocks": [
              "fe80::748f:5dff:fe78:d70b/64"
            ],
            "mtu": 1500
          },
          {
            "name": "vethcdc65b0",
            "macAddress": "b6:cf:0c:4a:a0:18",
            "ipv6CidrBlocks": [
              "fe80::b4cf:cff:fe4a:a018/64"
            ],
            "mtu": 1500
          },
          {
            "name": "br-b236d7c30f50",
            "macAddress": "02:42:19:72:d8:1f",
            "ipv4CidrBlocks": [
              "172.21.0.1/16"
            ],
            "ipv6CidrBlocks": [
              "fe80::42:19ff:fe72:d81f/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "br-0d09f6c0e547",
            "macAddress": "02:42:d7:47:50:92",
            "ipv4CidrBlocks": [
              "172.22.0.1/16"
            ],
            "ipv6CidrBlocks": [
              "fe80::42:d7ff:fe47:5092/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "vetha41e00b",
            "macAddress": "8e:0e:4b:1b:b4:f7",
            "ipv6CidrBlocks": [
              "fe80::8c0e:4bff:fe1b:b4f7/64"
            ],
            "mtu": 1500
          },
          {
            "name": "veth908225d",
            "macAddress": "da:0d:32:ea:45:07",
            "ipv6CidrBlocks": [
              "fe80::d80d:32ff:feea:4507/64"
            ],
            "mtu": 1500
          },
          {
            "name": "veth4a8cc8d",
            "macAddress": "f2:a4:00:66:45:85",
            "ipv6CidrBlocks": [
              "fe80::f0a4:ff:fe66:4585/64"
            ],
            "mtu": 1500
          },
          {
            "name": "veth3c415c0",
            "macAddress": "be:13:49:9b:84:e2",
            "ipv6CidrBlocks": [
              "fe80::bc13:49ff:fe9b:84e2/64"
            ],
            "mtu": 1500
          },
          {
            "name": "veth442f928",
            "macAddress": "be:66:6f:73:e1:63",
            "ipv6CidrBlocks": [
              "fe80::bc66:6fff:fe73:e163/64"
            ],
            "mtu": 1500
          },
          {
            "name": "veth644fb0c",
            "macAddress": "46:d7:01:9f:6f:3a",
            "ipv6CidrBlocks": [
              "fe80::44d7:1ff:fe9f:6f3a/64"
            ],
            "mtu": 1500
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.18.0.0",
            "interface": "br-58dab7383fc5",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.21.0.0",
            "interface": "br-b236d7c30f50",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.22.0.0",
            "interface": "br-0d09f6c0e547",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "cm-web",
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
          "totalSize": 255,
          "available": 251,
          "used": 4
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1312,
          "available": 1229,
          "used": 16
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
            "mtu": 65536
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
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:53:01:58",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b1:1c",
            "ipv4CidrBlocks": [
              "192.168.110.103/24"
            ],
            "ipv6CidrBlocks": [
              "fe80::a6bf:1ff:fe5a:b11c/64"
            ],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:53:01:59",
            "mtu": 1500
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
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:55:23:8c",
            "mtu": 1500
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
            "name": "docker0",
            "macAddress": "02:42:55:3d:6b:90",
            "ipv4CidrBlocks": [
              "172.17.0.1/16"
            ],
            "mtu": 1500,
            "state": "down"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
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

## Beetle section

### AWS

#### Recommend a target model for computing infra

> [!Note] > `desiredProvider` and `desiredRegion` are required.
>
> - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
> - If `desiredProvider` and `desiredRegion` are set on request body, the values in the query parameter will be ignored.

- API: `POST /recommendation/mci`
- Query params: `desiredProvider=aws`, `desiredRegion=ap-northeast-2`
  - Used query param for the later Cicada test
- Request body:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "desiredProvider": "",
  "desiredRegion": "",
  "onpremiseInfraModel": {
    "network": {
      "ipv4Networks": [
        "10.0.0.0/24",
        "192.168.110.0/24",
        "172.17.0.0/16",
        "172.16.0.128/32",
        "172.16.0.80/32",
        "172.29.0.0/24",
        "172.18.0.0/16",
        "172.21.0.0/16",
        "172.22.0.0/16"
      ],
      "ipv6Networks": ["fe80::/64"]
    },
    "servers": [
      {
        "hostname": "cm-nfs",
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
          "totalSize": 255,
          "available": 248,
          "used": 7
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1093,
          "available": 998,
          "used": 39
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
            "mtu": 65536
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:52:fa:e8",
            "mtu": 1500
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b0:03",
            "ipv4CidrBlocks": ["172.29.0.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b003/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:52:fa:e9",
            "mtu": 1500
          },
          {
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:47:70:f0",
            "mtu": 1500
          },
          {
            "name": "enp26s0f0",
            "macAddress": "b4:96:91:53:01:54",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b0:04",
            "ipv4CidrBlocks": ["192.168.110.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b004/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp26s0f1",
            "macAddress": "b4:96:91:53:01:55",
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
            "name": "docker0",
            "macAddress": "02:42:e3:e3:d6:97",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          },
          {
            "name": "br-58dab7383fc5",
            "macAddress": "02:42:03:69:20:4f",
            "ipv4CidrBlocks": ["172.18.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:3ff:fe69:204f/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "veth3308dac",
            "macAddress": "56:ed:7d:2a:96:7c",
            "ipv6CidrBlocks": ["fe80::54ed:7dff:fe2a:967c/64"],
            "mtu": 1500
          },
          {
            "name": "veth1d4c5c1",
            "macAddress": "76:8f:5d:78:d7:0b",
            "ipv6CidrBlocks": ["fe80::748f:5dff:fe78:d70b/64"],
            "mtu": 1500
          },
          {
            "name": "vethcdc65b0",
            "macAddress": "b6:cf:0c:4a:a0:18",
            "ipv6CidrBlocks": ["fe80::b4cf:cff:fe4a:a018/64"],
            "mtu": 1500
          },
          {
            "name": "br-b236d7c30f50",
            "macAddress": "02:42:19:72:d8:1f",
            "ipv4CidrBlocks": ["172.21.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:19ff:fe72:d81f/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "br-0d09f6c0e547",
            "macAddress": "02:42:d7:47:50:92",
            "ipv4CidrBlocks": ["172.22.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:d7ff:fe47:5092/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "vetha41e00b",
            "macAddress": "8e:0e:4b:1b:b4:f7",
            "ipv6CidrBlocks": ["fe80::8c0e:4bff:fe1b:b4f7/64"],
            "mtu": 1500
          },
          {
            "name": "veth908225d",
            "macAddress": "da:0d:32:ea:45:07",
            "ipv6CidrBlocks": ["fe80::d80d:32ff:feea:4507/64"],
            "mtu": 1500
          },
          {
            "name": "veth4a8cc8d",
            "macAddress": "f2:a4:00:66:45:85",
            "ipv6CidrBlocks": ["fe80::f0a4:ff:fe66:4585/64"],
            "mtu": 1500
          },
          {
            "name": "veth3c415c0",
            "macAddress": "be:13:49:9b:84:e2",
            "ipv6CidrBlocks": ["fe80::bc13:49ff:fe9b:84e2/64"],
            "mtu": 1500
          },
          {
            "name": "veth442f928",
            "macAddress": "be:66:6f:73:e1:63",
            "ipv6CidrBlocks": ["fe80::bc66:6fff:fe73:e163/64"],
            "mtu": 1500
          },
          {
            "name": "veth644fb0c",
            "macAddress": "46:d7:01:9f:6f:3a",
            "ipv6CidrBlocks": ["fe80::44d7:1ff:fe9f:6f3a/64"],
            "mtu": 1500
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.18.0.0",
            "interface": "br-58dab7383fc5",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.21.0.0",
            "interface": "br-b236d7c30f50",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.22.0.0",
            "interface": "br-0d09f6c0e547",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "cm-web",
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
          "totalSize": 255,
          "available": 251,
          "used": 4
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1312,
          "available": 1229,
          "used": 16
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b1:1b",
            "ipv4CidrBlocks": ["172.29.0.103/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b11b/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:53:01:58",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b1:1c",
            "ipv4CidrBlocks": ["192.168.110.103/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b11c/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:53:01:59",
            "mtu": 1500
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
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:55:23:8c",
            "mtu": 1500
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
            "name": "docker0",
            "macAddress": "02:42:55:3d:6b:90",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
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

```json
{
  "status": "ok",
  "description": "Target infra is recommended.",
  "targetInfra": {
    "name": "mmci01",
    "installMonAgent": "no",
    "label": null,
    "systemLabel": "",
    "description": "A cloud infra recommended by CM-Beetle",
    "vm": [
      {
        "name": "rehosted-cm-nfs",
        "subGroupSize": "",
        "label": null,
        "description": "a recommended virtual machine",
        "commonSpec": "aws+ap-northeast-2+i3.8xlarge",
        "commonImage": "aws+ap-northeast-2+ubuntu22.04"
      },
      {
        "name": "rehosted-cm-web",
        "subGroupSize": "",
        "label": null,
        "description": "a recommended virtual machine",
        "commonSpec": "aws+ap-northeast-2+i3.8xlarge",
        "commonImage": "aws+ap-northeast-2+ubuntu22.04"
      }
    ]
  }
}
```

#### Migrate the computing infra as defined in the target model

- API: `POST /migration/ns/{nsId}/mci`
- nsId: `mig01` (default)
- Request body:

> [!NOTE]
> As you can see, the partial info of previous result is used.
> `mmci01` is used for the name of migrated computing infra.

```json
{
  "name": "mmci01",
  "installMonAgent": "no",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "name": "rehosted-cm-nfs",
      "description": "a recommended virtual machine",
      "commonSpec": "aws+ap-northeast-2+i3.8xlarge",
      "commonImage": "aws+ap-northeast-2+ubuntu22.04"
    },
    {
      "name": "rehosted-cm-web",
      "description": "a recommended virtual machine",
      "commonSpec": "aws+ap-northeast-2+i3.8xlarge",
      "commonImage": "aws+ap-northeast-2+ubuntu22.04"
    }
  ]
}
```

- Response body:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "cstctumqujj5l373fr6g",
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
  "installMonAgent": "no",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "A cloud infra recommended by CM-Beetle",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "cstctumqujj5l373fr6g"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "resourceType": "vm",
      "id": "rehosted-cm-nfs-1",
      "uid": "cstctumqujj5l373fr7g",
      "cspResourceName": "cstctumqujj5l373fr7g",
      "cspResourceId": "i-0dea362991f4cd798",
      "name": "rehosted-cm-nfs-1",
      "subGroupId": "rehosted-cm-nfs",
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
      "createdTime": "2024-11-18 05:23:00",
      "label": {
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2024-11-18 05:23:00",
        "sys.cspResourceId": "i-0dea362991f4cd798",
        "sys.cspResourceName": "cstctumqujj5l373fr7g",
        "sys.id": "rehosted-cm-nfs-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-nfs-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-nfs",
        "sys.uid": "cstctumqujj5l373fr7g"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "13.125.167.165",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.4.18.246",
      "privateDNS": "ip-10-4-18-246.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "8",
      "rootDeviceName": "/dev/sda1",
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
      "specId": "aws+ap-northeast-2+i3.8xlarge",
      "cspSpecName": "i3.8xlarge",
      "imageId": "aws+ap-northeast-2+ubuntu22.04",
      "cspImageName": "ami-058165de3b7202099",
      "vNetId": "mig01-shared-aws-ap-northeast-2",
      "cspVNetId": "vpc-060e63bdbad0b74ec",
      "subnetId": "mig01-shared-aws-ap-northeast-2",
      "cspSubnetId": "subnet-048c467a1419095dc",
      "networkInterface": "eni-attach-0076e5bf7315db831",
      "securityGroupIds": ["mig01-shared-aws-ap-northeast-2"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-aws-ap-northeast-2",
      "cspSshKeyId": "cstcttuqujj5l373fr50",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "State",
          "value": "running"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "VpcId",
          "value": "vpc-060e63bdbad0b74ec"
        },
        {
          "key": "SubnetId",
          "value": "subnet-048c467a1419095dc"
        },
        {
          "key": "KeyName",
          "value": "cstcttuqujj5l373fr50"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "rehosted-cm-web-1",
      "uid": "cstctumqujj5l373fr8g",
      "cspResourceName": "cstctumqujj5l373fr8g",
      "cspResourceId": "i-0dd7bc31a24013186",
      "name": "rehosted-cm-web-1",
      "subGroupId": "rehosted-cm-web",
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
      "createdTime": "2024-11-18 05:22:33",
      "label": {
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2024-11-18 05:22:33",
        "sys.cspResourceId": "i-0dd7bc31a24013186",
        "sys.cspResourceName": "cstctumqujj5l373fr8g",
        "sys.id": "rehosted-cm-web-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-web-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-web",
        "sys.uid": "cstctumqujj5l373fr8g"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "43.201.5.95",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.4.26.112",
      "privateDNS": "ip-10-4-26-112.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "8",
      "rootDeviceName": "/dev/sda1",
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
      "specId": "aws+ap-northeast-2+i3.8xlarge",
      "cspSpecName": "i3.8xlarge",
      "imageId": "aws+ap-northeast-2+ubuntu22.04",
      "cspImageName": "ami-058165de3b7202099",
      "vNetId": "mig01-shared-aws-ap-northeast-2",
      "cspVNetId": "vpc-060e63bdbad0b74ec",
      "subnetId": "mig01-shared-aws-ap-northeast-2",
      "cspSubnetId": "subnet-048c467a1419095dc",
      "networkInterface": "eni-attach-02c7dd97d29c5ad6a",
      "securityGroupIds": ["mig01-shared-aws-ap-northeast-2"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-aws-ap-northeast-2",
      "cspSshKeyId": "cstcttuqujj5l373fr50",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "State",
          "value": "running"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "VpcId",
          "value": "vpc-060e63bdbad0b74ec"
        },
        {
          "key": "SubnetId",
          "value": "subnet-048c467a1419095dc"
        },
        {
          "key": "KeyName",
          "value": "cstcttuqujj5l373fr50"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        }
      ]
    }
  ],
  "newVmList": null
}
```

</details>

#### List the migrated computing infrastructures

- API: `GET /migration/ns/{nsId}/mci`
- nsId `mig01` (default)
- option: `id`
- Request body: none
- Response body:

```json
{
  "idList": ["mmci01"]
}
```

#### List the migrated computing infrastructures

- API: `GET /migration/ns/{nsId}/mci`
- nsId `mig01` (default)
- option: `info`
- Request body: none
- Response body:

<details>
	<summary> <ins>Click to see response body </ins> </summary>

```json
{
  "mci": [
    {
      "id": "mmci01",
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
      "installMonAgent": "no",
      "masterVmId": "rehosted-cm-nfs-1",
      "masterIp": "13.125.167.165",
      "masterSSHPort": "22",
      "label": null,
      "systemLabel": "",
      "vm": [
        {
          "id": "rehosted-cm-nfs-1",
          "cspResourceName": "cstctumqujj5l373fr7g",
          "name": "rehosted-cm-nfs-1",
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "nativeStatus": "Running",
          "monAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2024-11-18 05:23:00",
          "publicIp": "13.125.167.165",
          "privateIp": "10.4.18.246",
          "sshPort": "22",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          }
        },
        {
          "id": "rehosted-cm-web-1",
          "cspResourceName": "cstctumqujj5l373fr8g",
          "name": "rehosted-cm-web-1",
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "nativeStatus": "Running",
          "monAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2024-11-18 05:22:33",
          "publicIp": "43.201.5.95",
          "privateIp": "10.4.26.112",
          "sshPort": "22",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          }
        }
      ]
    }
  ]
}
```

</details>

#### Get the migrated computing infra

- API: `GET /migration/ns/{nsId}/mci/{mciId}`
- nsId `mig01` (default)
- mciId `mmci01`(default)
- Request body: None

<details>
	<summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "cstctumqujj5l373fr6g",
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
  "installMonAgent": "no",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "A cloud infra recommended by CM-Beetle",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "cstctumqujj5l373fr6g"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "resourceType": "vm",
      "id": "rehosted-cm-nfs-1",
      "uid": "cstctumqujj5l373fr7g",
      "cspResourceName": "cstctumqujj5l373fr7g",
      "cspResourceId": "i-0dea362991f4cd798",
      "name": "rehosted-cm-nfs-1",
      "subGroupId": "rehosted-cm-nfs",
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
      "createdTime": "2024-11-18 05:23:00",
      "label": {
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2024-11-18 05:23:00",
        "sys.cspResourceId": "i-0dea362991f4cd798",
        "sys.cspResourceName": "cstctumqujj5l373fr7g",
        "sys.id": "rehosted-cm-nfs-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-nfs-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-nfs",
        "sys.uid": "cstctumqujj5l373fr7g"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "13.125.167.165",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.4.18.246",
      "privateDNS": "ip-10-4-18-246.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "8",
      "rootDeviceName": "/dev/sda1",
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
      "specId": "aws+ap-northeast-2+i3.8xlarge",
      "cspSpecName": "i3.8xlarge",
      "imageId": "aws+ap-northeast-2+ubuntu22.04",
      "cspImageName": "ami-058165de3b7202099",
      "vNetId": "mig01-shared-aws-ap-northeast-2",
      "cspVNetId": "vpc-060e63bdbad0b74ec",
      "subnetId": "mig01-shared-aws-ap-northeast-2",
      "cspSubnetId": "subnet-048c467a1419095dc",
      "networkInterface": "eni-attach-0076e5bf7315db831",
      "securityGroupIds": ["mig01-shared-aws-ap-northeast-2"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-aws-ap-northeast-2",
      "cspSshKeyId": "cstcttuqujj5l373fr50",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "State",
          "value": "running"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "VpcId",
          "value": "vpc-060e63bdbad0b74ec"
        },
        {
          "key": "SubnetId",
          "value": "subnet-048c467a1419095dc"
        },
        {
          "key": "KeyName",
          "value": "cstcttuqujj5l373fr50"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "rehosted-cm-web-1",
      "uid": "cstctumqujj5l373fr8g",
      "cspResourceName": "cstctumqujj5l373fr8g",
      "cspResourceId": "i-0dd7bc31a24013186",
      "name": "rehosted-cm-web-1",
      "subGroupId": "rehosted-cm-web",
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
      "createdTime": "2024-11-18 05:22:33",
      "label": {
        "sys.connectionName": "aws-ap-northeast-2",
        "sys.createdTime": "2024-11-18 05:22:33",
        "sys.cspResourceId": "i-0dd7bc31a24013186",
        "sys.cspResourceName": "cstctumqujj5l373fr8g",
        "sys.id": "rehosted-cm-web-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-web-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-web",
        "sys.uid": "cstctumqujj5l373fr8g"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "ap-northeast-2",
        "Zone": "ap-northeast-2a"
      },
      "publicIP": "43.201.5.95",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.4.26.112",
      "privateDNS": "ip-10-4-26-112.ap-northeast-2.compute.internal",
      "rootDiskType": "gp2",
      "rootDiskSize": "8",
      "rootDeviceName": "/dev/sda1",
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
      "specId": "aws+ap-northeast-2+i3.8xlarge",
      "cspSpecName": "i3.8xlarge",
      "imageId": "aws+ap-northeast-2+ubuntu22.04",
      "cspImageName": "ami-058165de3b7202099",
      "vNetId": "mig01-shared-aws-ap-northeast-2",
      "cspVNetId": "vpc-060e63bdbad0b74ec",
      "subnetId": "mig01-shared-aws-ap-northeast-2",
      "cspSubnetId": "subnet-048c467a1419095dc",
      "networkInterface": "eni-attach-02c7dd97d29c5ad6a",
      "securityGroupIds": ["mig01-shared-aws-ap-northeast-2"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-aws-ap-northeast-2",
      "cspSshKeyId": "cstcttuqujj5l373fr50",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "State",
          "value": "running"
        },
        {
          "key": "Architecture",
          "value": "x86_64"
        },
        {
          "key": "VpcId",
          "value": "vpc-060e63bdbad0b74ec"
        },
        {
          "key": "SubnetId",
          "value": "subnet-048c467a1419095dc"
        },
        {
          "key": "KeyName",
          "value": "cstcttuqujj5l373fr50"
        },
        {
          "key": "VirtualizationType",
          "value": "hvm"
        }
      ]
    }
  ],
  "newVmList": null
}
```

</details>

#### Delete the migrated computing infra

- API: `DELETE /migration/ns/{nsId}/mci/{mciId}`
- nsId: `mig01`
- mciId: `mmci01`
- Query param: `?action=terminate` (default)
- Request body: None
- Response body:

```json
{
  "success": true,
  "text": ""
}
```

### Azure

> [!WARNING]
> Due to Total Regional Cores quota in Azure, we have **downgraded the server specifications** to the source computing environment.

#### Recommend a target model for computing infra

> [!Note] > `desiredProvider` and `desiredRegion` are required.
>
> - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
> - If `desiredProvider` and `desiredRegion` are set on request body, the values in the query parameter will be ignored.

- API: `POST /recommendation/mci`
- Query params: `desiredProvider=azure`, `desiredRegion=koreacentral`
  - Used query param for the later Cicada test
- Request body:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "desiredProvider": "",
  "desiredRegion": "",
  "onpremiseInfraModel": {
    "network": {
      "ipv4Networks": [
        "10.0.0.0/24",
        "192.168.110.0/24",
        "172.17.0.0/16",
        "172.16.0.128/32",
        "172.16.0.80/32",
        "172.29.0.0/24",
        "172.18.0.0/16",
        "172.21.0.0/16",
        "172.22.0.0/16"
      ],
      "ipv6Networks": ["fe80::/64"]
    },
    "servers": [
      {
        "hostname": "cm-nfs",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 2,
          "cores": 8,
          "threads": 16,
          "maxSpeed": 3.7,
          "vendor": "GenuineIntel",
          "model": "Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 32,
          "available": 25,
          "used": 7
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1093,
          "available": 998,
          "used": 39
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
            "mtu": 65536
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:52:fa:e8",
            "mtu": 1500
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b0:03",
            "ipv4CidrBlocks": ["172.29.0.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b003/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:52:fa:e9",
            "mtu": 1500
          },
          {
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:47:70:f0",
            "mtu": 1500
          },
          {
            "name": "enp26s0f0",
            "macAddress": "b4:96:91:53:01:54",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b0:04",
            "ipv4CidrBlocks": ["192.168.110.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b004/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp26s0f1",
            "macAddress": "b4:96:91:53:01:55",
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
            "name": "docker0",
            "macAddress": "02:42:e3:e3:d6:97",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          },
          {
            "name": "br-58dab7383fc5",
            "macAddress": "02:42:03:69:20:4f",
            "ipv4CidrBlocks": ["172.18.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:3ff:fe69:204f/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "veth3308dac",
            "macAddress": "56:ed:7d:2a:96:7c",
            "ipv6CidrBlocks": ["fe80::54ed:7dff:fe2a:967c/64"],
            "mtu": 1500
          },
          {
            "name": "veth1d4c5c1",
            "macAddress": "76:8f:5d:78:d7:0b",
            "ipv6CidrBlocks": ["fe80::748f:5dff:fe78:d70b/64"],
            "mtu": 1500
          },
          {
            "name": "vethcdc65b0",
            "macAddress": "b6:cf:0c:4a:a0:18",
            "ipv6CidrBlocks": ["fe80::b4cf:cff:fe4a:a018/64"],
            "mtu": 1500
          },
          {
            "name": "br-b236d7c30f50",
            "macAddress": "02:42:19:72:d8:1f",
            "ipv4CidrBlocks": ["172.21.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:19ff:fe72:d81f/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "br-0d09f6c0e547",
            "macAddress": "02:42:d7:47:50:92",
            "ipv4CidrBlocks": ["172.22.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:d7ff:fe47:5092/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "vetha41e00b",
            "macAddress": "8e:0e:4b:1b:b4:f7",
            "ipv6CidrBlocks": ["fe80::8c0e:4bff:fe1b:b4f7/64"],
            "mtu": 1500
          },
          {
            "name": "veth908225d",
            "macAddress": "da:0d:32:ea:45:07",
            "ipv6CidrBlocks": ["fe80::d80d:32ff:feea:4507/64"],
            "mtu": 1500
          },
          {
            "name": "veth4a8cc8d",
            "macAddress": "f2:a4:00:66:45:85",
            "ipv6CidrBlocks": ["fe80::f0a4:ff:fe66:4585/64"],
            "mtu": 1500
          },
          {
            "name": "veth3c415c0",
            "macAddress": "be:13:49:9b:84:e2",
            "ipv6CidrBlocks": ["fe80::bc13:49ff:fe9b:84e2/64"],
            "mtu": 1500
          },
          {
            "name": "veth442f928",
            "macAddress": "be:66:6f:73:e1:63",
            "ipv6CidrBlocks": ["fe80::bc66:6fff:fe73:e163/64"],
            "mtu": 1500
          },
          {
            "name": "veth644fb0c",
            "macAddress": "46:d7:01:9f:6f:3a",
            "ipv6CidrBlocks": ["fe80::44d7:1ff:fe9f:6f3a/64"],
            "mtu": 1500
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.18.0.0",
            "interface": "br-58dab7383fc5",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.21.0.0",
            "interface": "br-b236d7c30f50",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.22.0.0",
            "interface": "br-0d09f6c0e547",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "cm-web",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 2,
          "cores": 8,
          "threads": 16,
          "maxSpeed": 3.7,
          "vendor": "GenuineIntel",
          "model": "Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 32,
          "available": 25,
          "used": 7
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1312,
          "available": 1229,
          "used": 16
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b1:1b",
            "ipv4CidrBlocks": ["172.29.0.103/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b11b/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:53:01:58",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b1:1c",
            "ipv4CidrBlocks": ["192.168.110.103/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b11c/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:53:01:59",
            "mtu": 1500
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
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:55:23:8c",
            "mtu": 1500
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
            "name": "docker0",
            "macAddress": "02:42:55:3d:6b:90",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
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

```json
{
  "status": "ok",
  "description": "Target infra is recommended.",
  "targetInfra": {
    "name": "mmci01",
    "installMonAgent": "no",
    "label": null,
    "systemLabel": "",
    "description": "A cloud infra recommended by CM-Beetle",
    "vm": [
      {
        "name": "rehosted-cm-nfs",
        "subGroupSize": "",
        "label": null,
        "description": "a recommended virtual machine",
        "commonSpec": "azure+koreacentral+standard_b4ms",
        "commonImage": "azure+koreacentral+ubuntu22.04"
      },
      {
        "name": "rehosted-cm-web",
        "subGroupSize": "",
        "label": null,
        "description": "a recommended virtual machine",
        "commonSpec": "azure+koreacentral+standard_b4ms",
        "commonImage": "azure+koreacentral+ubuntu22.04"
      }
    ]
  }
}
```

### Migrate the computing infra as defined in the target model

- API: `POST /migration/ns/{nsId}/mci`
- nsId: `mig01` (default)
- Request body:

> [!NOTE]
> As you can see, the partial info of previous result is used.
> `mmci01` is used for the name of migrated computing infra.

```json
{
  "name": "mmci01",
  "installMonAgent": "no",
  "label": null,
  "systemLabel": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "name": "rehosted-cm-nfs",
      "subGroupSize": "",
      "label": null,
      "description": "a recommended virtual machine",
      "commonSpec": "azure+koreacentral+standard_b4ms",
      "commonImage": "azure+koreacentral+ubuntu22.04"
    },
    {
      "name": "rehosted-cm-web",
      "subGroupSize": "",
      "label": null,
      "description": "a recommended virtual machine",
      "commonSpec": "azure+koreacentral+standard_b4ms",
      "commonImage": "azure+koreacentral+ubuntu22.04"
    }
  ]
}
```

- Response body:
<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "cstf4uuqujj5l314leqg",
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
  "installMonAgent": "no",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "A cloud infra recommended by CM-Beetle",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "cstf4uuqujj5l314leqg"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "resourceType": "vm",
      "id": "rehosted-cm-nfs-1",
      "uid": "cstf4uuqujj5l314lesg",
      "cspResourceName": "cstf4uuqujj5l314lesg",
      "cspResourceId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/cstf4uuqujj5l314lesg",
      "name": "rehosted-cm-nfs-1",
      "subGroupId": "rehosted-cm-nfs",
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
      "createdTime": "2024-11-18 07:54:34",
      "label": {
        "sys.connectionName": "azure-koreacentral",
        "sys.createdTime": "2024-11-18 07:54:34",
        "sys.cspResourceId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/cstf4uuqujj5l314lesg",
        "sys.cspResourceName": "cstf4uuqujj5l314lesg",
        "sys.id": "rehosted-cm-nfs-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-nfs-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-nfs",
        "sys.uid": "cstf4uuqujj5l314lesg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "koreacentral",
        "Zone": "1"
      },
      "publicIP": "52.231.107.88",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.48.0.4",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "30",
      "rootDeviceName": "Not visible in Azure",
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
          "zones": ["1", "2", "3"]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreacentral+standard_b4ms",
      "cspSpecName": "Standard_B4ms",
      "imageId": "azure+koreacentral+ubuntu22.04",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts:22.04.202404090",
      "vNetId": "mig01-shared-azure-koreacentral",
      "cspVNetId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/cstf4puqujj5l314leo0",
      "subnetId": "mig01-shared-azure-koreacentral",
      "cspSubnetId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/cstf4puqujj5l314leo0/subnets/cstf4puqujj5l314leog",
      "networkInterface": "cstf4uuqujj5l314lesg-96478-VNic",
      "securityGroupIds": ["mig01-shared-azure-koreacentral"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-azure-koreacentral",
      "cspSshKeyId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/sshPublicKeys/cstf4tequjj5l314lepg",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "publicip",
          "value": "cstf4uuqujj5l314lesg-2651-PublicIP"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "rehosted-cm-web-1",
      "uid": "cstf4uuqujj5l314lerg",
      "cspResourceName": "cstf4uuqujj5l314lerg",
      "cspResourceId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/cstf4uuqujj5l314lerg",
      "name": "rehosted-cm-web-1",
      "subGroupId": "rehosted-cm-web",
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
      "createdTime": "2024-11-18 07:54:37",
      "label": {
        "sys.connectionName": "azure-koreacentral",
        "sys.createdTime": "2024-11-18 07:54:37",
        "sys.cspResourceId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/cstf4uuqujj5l314lerg",
        "sys.cspResourceName": "cstf4uuqujj5l314lerg",
        "sys.id": "rehosted-cm-web-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-web-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-web",
        "sys.uid": "cstf4uuqujj5l314lerg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "koreacentral",
        "Zone": "1"
      },
      "publicIP": "52.231.111.187",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.48.0.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "30",
      "rootDeviceName": "Not visible in Azure",
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
          "zones": ["1", "2", "3"]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreacentral+standard_b4ms",
      "cspSpecName": "Standard_B4ms",
      "imageId": "azure+koreacentral+ubuntu22.04",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts:22.04.202404090",
      "vNetId": "mig01-shared-azure-koreacentral",
      "cspVNetId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/cstf4puqujj5l314leo0",
      "subnetId": "mig01-shared-azure-koreacentral",
      "cspSubnetId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/cstf4puqujj5l314leo0/subnets/cstf4puqujj5l314leog",
      "networkInterface": "cstf4uuqujj5l314lerg-83701-VNic",
      "securityGroupIds": ["mig01-shared-azure-koreacentral"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-azure-koreacentral",
      "cspSshKeyId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/sshPublicKeys/cstf4tequjj5l314lepg",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "publicip",
          "value": "cstf4uuqujj5l314lerg-71346-PublicIP"
        }
      ]
    }
  ],
  "newVmList": null
}
```

</details>

#### List the migrated computing infrastructures

- API: `GET /migration/ns/{nsId}/mci`
- nsId `mig01` (default)
- option: `id`
- Request body: none
- Response body:

```json
{
  "idList": ["mmci01"]
}
```

#### List the migrated computing infrastructures

- API: `GET /migration/ns/{nsId}/mci`
- nsId `mig01` (default)
- option: `info`
- Request body: none
- Response body:

<details>
	<summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "mci": [
    {
      "id": "mmci01",
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
      "installMonAgent": "no",
      "masterVmId": "rehosted-cm-nfs-1",
      "masterIp": "52.231.107.88",
      "masterSSHPort": "22",
      "label": null,
      "systemLabel": "",
      "vm": [
        {
          "id": "rehosted-cm-nfs-1",
          "cspResourceName": "cstf4uuqujj5l314lesg",
          "name": "rehosted-cm-nfs-1",
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "nativeStatus": "Running",
          "monAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2024-11-18 07:54:34",
          "publicIp": "52.231.107.88",
          "privateIp": "10.48.0.4",
          "sshPort": "22",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          }
        },
        {
          "id": "rehosted-cm-web-1",
          "cspResourceName": "cstf4uuqujj5l314lerg",
          "name": "rehosted-cm-web-1",
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "nativeStatus": "Running",
          "monAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2024-11-18 07:54:37",
          "publicIp": "52.231.111.187",
          "privateIp": "10.48.0.5",
          "sshPort": "22",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          }
        }
      ]
    }
  ]
}
```

</details>

### Get the migrated computing infra

- API: `GET /migration/ns/{nsId}/mci/{mciId}`
- nsId `mig01` (default)
- mciId `mmci01`(default)

- Request body: None

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "cstf4uuqujj5l314leqg",
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
  "installMonAgent": "no",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "A cloud infra recommended by CM-Beetle",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "cstf4uuqujj5l314leqg"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "resourceType": "vm",
      "id": "rehosted-cm-nfs-1",
      "uid": "cstf4uuqujj5l314lesg",
      "cspResourceName": "cstf4uuqujj5l314lesg",
      "cspResourceId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/cstf4uuqujj5l314lesg",
      "name": "rehosted-cm-nfs-1",
      "subGroupId": "rehosted-cm-nfs",
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
      "createdTime": "2024-11-18 07:54:34",
      "label": {
        "sys.connectionName": "azure-koreacentral",
        "sys.createdTime": "2024-11-18 07:54:34",
        "sys.cspResourceId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/cstf4uuqujj5l314lesg",
        "sys.cspResourceName": "cstf4uuqujj5l314lesg",
        "sys.id": "rehosted-cm-nfs-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-nfs-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-nfs",
        "sys.uid": "cstf4uuqujj5l314lesg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "koreacentral",
        "Zone": "1"
      },
      "publicIP": "52.231.107.88",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.48.0.4",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "30",
      "rootDeviceName": "Not visible in Azure",
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
          "zones": ["1", "2", "3"]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreacentral+standard_b4ms",
      "cspSpecName": "Standard_B4ms",
      "imageId": "azure+koreacentral+ubuntu22.04",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts:22.04.202404090",
      "vNetId": "mig01-shared-azure-koreacentral",
      "cspVNetId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/cstf4puqujj5l314leo0",
      "subnetId": "mig01-shared-azure-koreacentral",
      "cspSubnetId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/cstf4puqujj5l314leo0/subnets/cstf4puqujj5l314leog",
      "networkInterface": "cstf4uuqujj5l314lesg-96478-VNic",
      "securityGroupIds": ["mig01-shared-azure-koreacentral"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-azure-koreacentral",
      "cspSshKeyId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/sshPublicKeys/cstf4tequjj5l314lepg",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "publicip",
          "value": "cstf4uuqujj5l314lesg-2651-PublicIP"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "rehosted-cm-web-1",
      "uid": "cstf4uuqujj5l314lerg",
      "cspResourceName": "cstf4uuqujj5l314lerg",
      "cspResourceId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/cstf4uuqujj5l314lerg",
      "name": "rehosted-cm-web-1",
      "subGroupId": "rehosted-cm-web",
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
      "createdTime": "2024-11-18 07:54:37",
      "label": {
        "sys.connectionName": "azure-koreacentral",
        "sys.createdTime": "2024-11-18 07:54:37",
        "sys.cspResourceId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/virtualMachines/cstf4uuqujj5l314lerg",
        "sys.cspResourceName": "cstf4uuqujj5l314lerg",
        "sys.id": "rehosted-cm-web-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-web-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-web",
        "sys.uid": "cstf4uuqujj5l314lerg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "koreacentral",
        "Zone": "1"
      },
      "publicIP": "52.231.111.187",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.48.0.5",
      "privateDNS": "",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": "30",
      "rootDeviceName": "Not visible in Azure",
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
          "zones": ["1", "2", "3"]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "specId": "azure+koreacentral+standard_b4ms",
      "cspSpecName": "Standard_B4ms",
      "imageId": "azure+koreacentral+ubuntu22.04",
      "cspImageName": "Canonical:0001-com-ubuntu-server-jammy:22_04-lts:22.04.202404090",
      "vNetId": "mig01-shared-azure-koreacentral",
      "cspVNetId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/cstf4puqujj5l314leo0",
      "subnetId": "mig01-shared-azure-koreacentral",
      "cspSubnetId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/cstf4puqujj5l314leo0/subnets/cstf4puqujj5l314leog",
      "networkInterface": "cstf4uuqujj5l314lerg-83701-VNic",
      "securityGroupIds": ["mig01-shared-azure-koreacentral"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-azure-koreacentral",
      "cspSshKeyId": "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxxx/resourceGroups/koreacentral/providers/Microsoft.Compute/sshPublicKeys/cstf4tequjj5l314lepg",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "publicip",
          "value": "cstf4uuqujj5l314lerg-71346-PublicIP"
        }
      ]
    }
  ],
  "newVmList": null
}
```

</details>

### Delete the migrated computing infra

- API: `DELETE /migration/ns/{nsId}/mci/{mciId}`
- nsId: `mig01`
- mciId: `mmci01`
- Query param: `?action=terminate` (default)
- Request body: None
- Response body:

```json
{
  "success": true,
  "text": ""
}
```

### GCP

#### Recommend a target model for computing infra

> [!Note] > `desiredProvider` and `desiredRegion` are required.
>
> - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
> - If `desiredProvider` and `desiredRegion` are set on request body, the values in the query parameter will be ignored.

- API: `POST /recommendation/mci`
- Query params: `desiredProvider=gcp`, `desiredRegion=asia-northeast3`
  - Used query param for the later Cicada test
- Request body:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "desiredProvider": "",
  "desiredRegion": "",
  "onpremiseInfraModel": {
    "network": {
      "ipv4Networks": [
        "10.0.0.0/24",
        "192.168.110.0/24",
        "172.17.0.0/16",
        "172.16.0.128/32",
        "172.16.0.80/32",
        "172.29.0.0/24",
        "172.18.0.0/16",
        "172.21.0.0/16",
        "172.22.0.0/16"
      ],
      "ipv6Networks": ["fe80::/64"]
    },
    "servers": [
      {
        "hostname": "cm-nfs",
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
          "totalSize": 255,
          "available": 248,
          "used": 7
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1093,
          "available": 998,
          "used": 39
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
            "mtu": 65536
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:52:fa:e8",
            "mtu": 1500
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b0:03",
            "ipv4CidrBlocks": ["172.29.0.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b003/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:52:fa:e9",
            "mtu": 1500
          },
          {
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:47:70:f0",
            "mtu": 1500
          },
          {
            "name": "enp26s0f0",
            "macAddress": "b4:96:91:53:01:54",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b0:04",
            "ipv4CidrBlocks": ["192.168.110.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b004/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp26s0f1",
            "macAddress": "b4:96:91:53:01:55",
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
            "name": "docker0",
            "macAddress": "02:42:e3:e3:d6:97",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          },
          {
            "name": "br-58dab7383fc5",
            "macAddress": "02:42:03:69:20:4f",
            "ipv4CidrBlocks": ["172.18.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:3ff:fe69:204f/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "veth3308dac",
            "macAddress": "56:ed:7d:2a:96:7c",
            "ipv6CidrBlocks": ["fe80::54ed:7dff:fe2a:967c/64"],
            "mtu": 1500
          },
          {
            "name": "veth1d4c5c1",
            "macAddress": "76:8f:5d:78:d7:0b",
            "ipv6CidrBlocks": ["fe80::748f:5dff:fe78:d70b/64"],
            "mtu": 1500
          },
          {
            "name": "vethcdc65b0",
            "macAddress": "b6:cf:0c:4a:a0:18",
            "ipv6CidrBlocks": ["fe80::b4cf:cff:fe4a:a018/64"],
            "mtu": 1500
          },
          {
            "name": "br-b236d7c30f50",
            "macAddress": "02:42:19:72:d8:1f",
            "ipv4CidrBlocks": ["172.21.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:19ff:fe72:d81f/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "br-0d09f6c0e547",
            "macAddress": "02:42:d7:47:50:92",
            "ipv4CidrBlocks": ["172.22.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:d7ff:fe47:5092/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "vetha41e00b",
            "macAddress": "8e:0e:4b:1b:b4:f7",
            "ipv6CidrBlocks": ["fe80::8c0e:4bff:fe1b:b4f7/64"],
            "mtu": 1500
          },
          {
            "name": "veth908225d",
            "macAddress": "da:0d:32:ea:45:07",
            "ipv6CidrBlocks": ["fe80::d80d:32ff:feea:4507/64"],
            "mtu": 1500
          },
          {
            "name": "veth4a8cc8d",
            "macAddress": "f2:a4:00:66:45:85",
            "ipv6CidrBlocks": ["fe80::f0a4:ff:fe66:4585/64"],
            "mtu": 1500
          },
          {
            "name": "veth3c415c0",
            "macAddress": "be:13:49:9b:84:e2",
            "ipv6CidrBlocks": ["fe80::bc13:49ff:fe9b:84e2/64"],
            "mtu": 1500
          },
          {
            "name": "veth442f928",
            "macAddress": "be:66:6f:73:e1:63",
            "ipv6CidrBlocks": ["fe80::bc66:6fff:fe73:e163/64"],
            "mtu": 1500
          },
          {
            "name": "veth644fb0c",
            "macAddress": "46:d7:01:9f:6f:3a",
            "ipv6CidrBlocks": ["fe80::44d7:1ff:fe9f:6f3a/64"],
            "mtu": 1500
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.18.0.0",
            "interface": "br-58dab7383fc5",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.21.0.0",
            "interface": "br-b236d7c30f50",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.22.0.0",
            "interface": "br-0d09f6c0e547",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "cm-web",
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
          "totalSize": 255,
          "available": 251,
          "used": 4
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1312,
          "available": 1229,
          "used": 16
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b1:1b",
            "ipv4CidrBlocks": ["172.29.0.103/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b11b/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:53:01:58",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b1:1c",
            "ipv4CidrBlocks": ["192.168.110.103/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b11c/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:53:01:59",
            "mtu": 1500
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
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:55:23:8c",
            "mtu": 1500
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
            "name": "docker0",
            "macAddress": "02:42:55:3d:6b:90",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
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

```json
{
  "status": "ok",
  "description": "Target infra is recommended.",
  "targetInfra": {
    "name": "mmci01",
    "installMonAgent": "no",
    "label": null,
    "systemLabel": "",
    "description": "A cloud infra recommended by CM-Beetle",
    "vm": [
      {
        "name": "rehosted-cm-nfs",
        "subGroupSize": "",
        "label": null,
        "description": "a recommended virtual machine",
        "commonSpec": "gcp+asia-northeast3+g2-standard-32",
        "commonImage": "gcp+asia-northeast3+ubuntu22.04"
      },
      {
        "name": "rehosted-cm-web",
        "subGroupSize": "",
        "label": null,
        "description": "a recommended virtual machine",
        "commonSpec": "gcp+asia-northeast3+g2-standard-32",
        "commonImage": "gcp+asia-northeast3+ubuntu22.04"
      }
    ]
  }
}
```

### Migrate the computing infra as defined in the target model

- API: `POST /migration/ns/{nsId}/mci`
- nsId: `mig01` (default)
- Request body:

> [!NOTE]
> As you can see, the partial info of previous result is used.
> `mmci01` is used for the name of migrated computing infra.

```json
{
  "name": "mmci01",
  "installMonAgent": "no",
  "label": null,
  "systemLabel": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "name": "rehosted-cm-nfs",
      "subGroupSize": "",
      "label": null,
      "description": "a recommended virtual machine",
      "commonSpec": "gcp+asia-northeast3+g2-standard-32",
      "commonImage": "gcp+asia-northeast3+ubuntu22.04"
    },
    {
      "name": "rehosted-cm-web",
      "subGroupSize": "",
      "label": null,
      "description": "a recommended virtual machine",
      "commonSpec": "gcp+asia-northeast3+g2-standard-32",
      "commonImage": "gcp+asia-northeast3+ubuntu22.04"
    }
  ]
}
```

- Response body:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "cstfbbequjj5l31al3jg",
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
  "installMonAgent": "no",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "A cloud infra recommended by CM-Beetle",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "cstfbbequjj5l31al3jg"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "resourceType": "vm",
      "id": "rehosted-cm-nfs-1",
      "uid": "cstfbbequjj5l31al3kg",
      "cspResourceName": "cstfbbequjj5l31al3kg",
      "cspResourceId": "cstfbbequjj5l31al3kg",
      "name": "rehosted-cm-nfs-1",
      "subGroupId": "rehosted-cm-nfs",
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
      "createdTime": "2024-11-18 08:08:03",
      "label": {
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2024-11-18 08:08:03",
        "sys.cspResourceId": "cstfbbequjj5l31al3kg",
        "sys.cspResourceName": "cstfbbequjj5l31al3kg",
        "sys.id": "rehosted-cm-nfs-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-nfs-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-nfs",
        "sys.uid": "cstfbbequjj5l31al3kg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.22.91.51",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.77.0.2",
      "privateDNS": "",
      "rootDiskType": "pd-balanced",
      "rootDiskSize": "10",
      "rootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+g2-standard-32",
      "cspSpecName": "g2-standard-32",
      "imageId": "gcp+asia-northeast3+ubuntu22.04",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20240319",
      "vNetId": "mig01-shared-gcp-asia-northeast3",
      "cspVNetId": "cstfaduqujj5l31al3g0",
      "subnetId": "mig01-shared-gcp-asia-northeast3",
      "cspSubnetId": "cstfaduqujj5l31al3gg",
      "networkInterface": "nic0",
      "securityGroupIds": ["mig01-shared-gcp-asia-northeast3"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-gcp-asia-northeast3",
      "cspSshKeyId": "cstfasuqujj5l31al3hg",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "SubNetwork",
          "value": "https://www.googleapis.com/compute/v1/projects/xxxxxxx/regions/asia-northeast3/subnetworks/cstfaduqujj5l31al3gg"
        },
        {
          "key": "AccessConfigName",
          "value": "External NAT"
        },
        {
          "key": "NetworkTier",
          "value": "PREMIUM"
        },
        {
          "key": "DiskDeviceName",
          "value": "persistent-disk-0"
        },
        {
          "key": "DiskName",
          "value": "https://www.googleapis.com/compute/v1/projects/xxxxxxx/zones/asia-northeast3-a/disks/cstfbbequjj5l31al3kg"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "ZoneUrl",
          "value": "https://www.googleapis.com/compute/v1/projects/xxxxxxx/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "rehosted-cm-web-1",
      "uid": "cstfbbequjj5l31al3lg",
      "cspResourceName": "cstfbbequjj5l31al3lg",
      "cspResourceId": "cstfbbequjj5l31al3lg",
      "name": "rehosted-cm-web-1",
      "subGroupId": "rehosted-cm-web",
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
      "createdTime": "2024-11-18 08:08:06",
      "label": {
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2024-11-18 08:08:06",
        "sys.cspResourceId": "cstfbbequjj5l31al3lg",
        "sys.cspResourceName": "cstfbbequjj5l31al3lg",
        "sys.id": "rehosted-cm-web-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-web-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-web",
        "sys.uid": "cstfbbequjj5l31al3lg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.64.54.80",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.77.0.3",
      "privateDNS": "",
      "rootDiskType": "pd-balanced",
      "rootDiskSize": "10",
      "rootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+g2-standard-32",
      "cspSpecName": "g2-standard-32",
      "imageId": "gcp+asia-northeast3+ubuntu22.04",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20240319",
      "vNetId": "mig01-shared-gcp-asia-northeast3",
      "cspVNetId": "cstfaduqujj5l31al3g0",
      "subnetId": "mig01-shared-gcp-asia-northeast3",
      "cspSubnetId": "cstfaduqujj5l31al3gg",
      "networkInterface": "nic0",
      "securityGroupIds": ["mig01-shared-gcp-asia-northeast3"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-gcp-asia-northeast3",
      "cspSshKeyId": "cstfasuqujj5l31al3hg",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "SubNetwork",
          "value": "https://www.googleapis.com/compute/v1/projects/xxxxxxx/regions/asia-northeast3/subnetworks/cstfaduqujj5l31al3gg"
        },
        {
          "key": "AccessConfigName",
          "value": "External NAT"
        },
        {
          "key": "NetworkTier",
          "value": "PREMIUM"
        },
        {
          "key": "DiskDeviceName",
          "value": "persistent-disk-0"
        },
        {
          "key": "DiskName",
          "value": "https://www.googleapis.com/compute/v1/projects/xxxxxxx/zones/asia-northeast3-a/disks/cstfbbequjj5l31al3lg"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "ZoneUrl",
          "value": "https://www.googleapis.com/compute/v1/projects/xxxxxxx/zones/asia-northeast3-a"
        }
      ]
    }
  ],
  "newVmList": null
}
```

</details>

#### List the migrated computing infrastructures

- API: `GET /migration/ns/{nsId}/mci`
- nsId `mig01` (default)
- option: `id`
- Request body: none
- Response body:

```json
{
  "idList": ["mmci01"]
}
```

#### List the migrated computing infrastructures

- API: `GET /migration/ns/{nsId}/mci`
- nsId `mig01` (default)
- option: `info`
- Request body: none
- Response body:

<details>
	<summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "mci": [
    {
      "id": "mmci01",
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
      "installMonAgent": "no",
      "masterVmId": "rehosted-cm-nfs-1",
      "masterIp": "34.22.91.51",
      "masterSSHPort": "22",
      "label": null,
      "systemLabel": "",
      "vm": [
        {
          "id": "rehosted-cm-nfs-1",
          "cspResourceName": "cstfbbequjj5l31al3kg",
          "name": "rehosted-cm-nfs-1",
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "nativeStatus": "Running",
          "monAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2024-11-18 08:08:03",
          "publicIp": "34.22.91.51",
          "privateIp": "10.77.0.2",
          "sshPort": "22",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          }
        },
        {
          "id": "rehosted-cm-web-1",
          "cspResourceName": "cstfbbequjj5l31al3lg",
          "name": "rehosted-cm-web-1",
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "nativeStatus": "Running",
          "monAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2024-11-18 08:08:06",
          "publicIp": "34.64.54.80",
          "privateIp": "10.77.0.3",
          "sshPort": "22",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.2,
            "longitude": 127
          }
        }
      ]
    }
  ]
}
```

</details>

### Get the migrated computing infra

- API: `GET /migration/ns/{nsId}/mci/{mciId}`
- nsId `mig01` (default)
- mciId `mmci01`(default)
- Request body: None
- Response body:
<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "cstfbbequjj5l31al3jg",
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
  "installMonAgent": "no",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "A cloud infra recommended by CM-Beetle",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "cstfbbequjj5l31al3jg"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "resourceType": "vm",
      "id": "rehosted-cm-nfs-1",
      "uid": "cstfbbequjj5l31al3kg",
      "cspResourceName": "cstfbbequjj5l31al3kg",
      "cspResourceId": "cstfbbequjj5l31al3kg",
      "name": "rehosted-cm-nfs-1",
      "subGroupId": "rehosted-cm-nfs",
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
      "createdTime": "2024-11-18 08:08:03",
      "label": {
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2024-11-18 08:08:03",
        "sys.cspResourceId": "cstfbbequjj5l31al3kg",
        "sys.cspResourceName": "cstfbbequjj5l31al3kg",
        "sys.id": "rehosted-cm-nfs-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-nfs-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-nfs",
        "sys.uid": "cstfbbequjj5l31al3kg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.22.91.51",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.77.0.2",
      "privateDNS": "",
      "rootDiskType": "pd-balanced",
      "rootDiskSize": "10",
      "rootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+g2-standard-32",
      "cspSpecName": "g2-standard-32",
      "imageId": "gcp+asia-northeast3+ubuntu22.04",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20240319",
      "vNetId": "mig01-shared-gcp-asia-northeast3",
      "cspVNetId": "cstfaduqujj5l31al3g0",
      "subnetId": "mig01-shared-gcp-asia-northeast3",
      "cspSubnetId": "cstfaduqujj5l31al3gg",
      "networkInterface": "nic0",
      "securityGroupIds": ["mig01-shared-gcp-asia-northeast3"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-gcp-asia-northeast3",
      "cspSshKeyId": "cstfasuqujj5l31al3hg",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "SubNetwork",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/cstfaduqujj5l31al3gg"
        },
        {
          "key": "AccessConfigName",
          "value": "External NAT"
        },
        {
          "key": "NetworkTier",
          "value": "PREMIUM"
        },
        {
          "key": "DiskDeviceName",
          "value": "persistent-disk-0"
        },
        {
          "key": "DiskName",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/cstfbbequjj5l31al3kg"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "ZoneUrl",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "rehosted-cm-web-1",
      "uid": "cstfbbequjj5l31al3lg",
      "cspResourceName": "cstfbbequjj5l31al3lg",
      "cspResourceId": "cstfbbequjj5l31al3lg",
      "name": "rehosted-cm-web-1",
      "subGroupId": "rehosted-cm-web",
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
      "createdTime": "2024-11-18 08:08:06",
      "label": {
        "sys.connectionName": "gcp-asia-northeast3",
        "sys.createdTime": "2024-11-18 08:08:06",
        "sys.cspResourceId": "cstfbbequjj5l31al3lg",
        "sys.cspResourceName": "cstfbbequjj5l31al3lg",
        "sys.id": "rehosted-cm-web-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-web-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-web",
        "sys.uid": "cstfbbequjj5l31al3lg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "asia-northeast3",
        "Zone": "asia-northeast3-a"
      },
      "publicIP": "34.64.54.80",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.77.0.3",
      "privateDNS": "",
      "rootDiskType": "pd-balanced",
      "rootDiskSize": "10",
      "rootDeviceName": "persistent-disk-0",
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
      "specId": "gcp+asia-northeast3+g2-standard-32",
      "cspSpecName": "g2-standard-32",
      "imageId": "gcp+asia-northeast3+ubuntu22.04",
      "cspImageName": "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-2204-jammy-v20240319",
      "vNetId": "mig01-shared-gcp-asia-northeast3",
      "cspVNetId": "cstfaduqujj5l31al3g0",
      "subnetId": "mig01-shared-gcp-asia-northeast3",
      "cspSubnetId": "cstfaduqujj5l31al3gg",
      "networkInterface": "nic0",
      "securityGroupIds": ["mig01-shared-gcp-asia-northeast3"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-gcp-asia-northeast3",
      "cspSshKeyId": "cstfasuqujj5l31al3hg",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "SubNetwork",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/regions/asia-northeast3/subnetworks/cstfaduqujj5l31al3gg"
        },
        {
          "key": "AccessConfigName",
          "value": "External NAT"
        },
        {
          "key": "NetworkTier",
          "value": "PREMIUM"
        },
        {
          "key": "DiskDeviceName",
          "value": "persistent-disk-0"
        },
        {
          "key": "DiskName",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a/disks/cstfbbequjj5l31al3lg"
        },
        {
          "key": "Kind",
          "value": "compute#instance"
        },
        {
          "key": "ZoneUrl",
          "value": "https://www.googleapis.com/compute/v1/projects/ykkim-etri/zones/asia-northeast3-a"
        }
      ]
    }
  ],
  "newVmList": null
}
```

</details>

### Delete the migrated computing infra

- API: `DELETE /migration/ns/{nsId}/mci/{mciId}`
- nsId: `mig01`
- mciId: `mmci01`
- Query param: `?action=terminate` (default)
- Request body: None
- Response body:

```json
{
  "success": true,
  "text": ""
}
```

### NCP

> [!WARNING]
> Due to NCP resource constraints, we have **downgraded the server specifications** to the source computing environment.

#### Recommend a target model for computing infra

> [!Note] > `desiredProvider` and `desiredRegion` are required.
>
> - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
> - If `desiredProvider` and `desiredRegion` are set on request body, the values in the query parameter will be ignored.

- API: `POST /recommendation/mci`
- Query params: `desiredProvider=ncp`, `desiredRegion=kr`
  - Used query param for the later Cicada test
- Request body:

<details>
  <summary> <ins>Click to see the request body </ins> </summary>

```json
{
  "desiredProvider": "",
  "desiredRegion": "",
  "onpremiseInfraModel": {
    "network": {
      "ipv4Networks": [
        "10.0.0.0/24",
        "192.168.110.0/24",
        "172.17.0.0/16",
        "172.16.0.128/32",
        "172.16.0.80/32",
        "172.29.0.0/24",
        "172.18.0.0/16",
        "172.21.0.0/16",
        "172.22.0.0/16"
      ],
      "ipv6Networks": ["fe80::/64"]
    },
    "servers": [
      {
        "hostname": "cm-nfs",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 2,
          "cores": 2,
          "threads": 2,
          "maxSpeed": 3.7,
          "vendor": "GenuineIntel",
          "model": "Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 4,
          "available": 2,
          "used": 2
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1093,
          "available": 998,
          "used": 39
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
            "mtu": 65536
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:52:fa:e8",
            "mtu": 1500
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b0:03",
            "ipv4CidrBlocks": ["172.29.0.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b003/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:52:fa:e9",
            "mtu": 1500
          },
          {
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:47:70:f0",
            "mtu": 1500
          },
          {
            "name": "enp26s0f0",
            "macAddress": "b4:96:91:53:01:54",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b0:04",
            "ipv4CidrBlocks": ["192.168.110.102/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b004/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp26s0f1",
            "macAddress": "b4:96:91:53:01:55",
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
            "name": "docker0",
            "macAddress": "02:42:e3:e3:d6:97",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          },
          {
            "name": "br-58dab7383fc5",
            "macAddress": "02:42:03:69:20:4f",
            "ipv4CidrBlocks": ["172.18.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:3ff:fe69:204f/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "veth3308dac",
            "macAddress": "56:ed:7d:2a:96:7c",
            "ipv6CidrBlocks": ["fe80::54ed:7dff:fe2a:967c/64"],
            "mtu": 1500
          },
          {
            "name": "veth1d4c5c1",
            "macAddress": "76:8f:5d:78:d7:0b",
            "ipv6CidrBlocks": ["fe80::748f:5dff:fe78:d70b/64"],
            "mtu": 1500
          },
          {
            "name": "vethcdc65b0",
            "macAddress": "b6:cf:0c:4a:a0:18",
            "ipv6CidrBlocks": ["fe80::b4cf:cff:fe4a:a018/64"],
            "mtu": 1500
          },
          {
            "name": "br-b236d7c30f50",
            "macAddress": "02:42:19:72:d8:1f",
            "ipv4CidrBlocks": ["172.21.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:19ff:fe72:d81f/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "br-0d09f6c0e547",
            "macAddress": "02:42:d7:47:50:92",
            "ipv4CidrBlocks": ["172.22.0.1/16"],
            "ipv6CidrBlocks": ["fe80::42:d7ff:fe47:5092/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "vetha41e00b",
            "macAddress": "8e:0e:4b:1b:b4:f7",
            "ipv6CidrBlocks": ["fe80::8c0e:4bff:fe1b:b4f7/64"],
            "mtu": 1500
          },
          {
            "name": "veth908225d",
            "macAddress": "da:0d:32:ea:45:07",
            "ipv6CidrBlocks": ["fe80::d80d:32ff:feea:4507/64"],
            "mtu": 1500
          },
          {
            "name": "veth4a8cc8d",
            "macAddress": "f2:a4:00:66:45:85",
            "ipv6CidrBlocks": ["fe80::f0a4:ff:fe66:4585/64"],
            "mtu": 1500
          },
          {
            "name": "veth3c415c0",
            "macAddress": "be:13:49:9b:84:e2",
            "ipv6CidrBlocks": ["fe80::bc13:49ff:fe9b:84e2/64"],
            "mtu": 1500
          },
          {
            "name": "veth442f928",
            "macAddress": "be:66:6f:73:e1:63",
            "ipv6CidrBlocks": ["fe80::bc66:6fff:fe73:e163/64"],
            "mtu": 1500
          },
          {
            "name": "veth644fb0c",
            "macAddress": "46:d7:01:9f:6f:3a",
            "ipv6CidrBlocks": ["fe80::44d7:1ff:fe9f:6f3a/64"],
            "mtu": 1500
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.18.0.0",
            "interface": "br-58dab7383fc5",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.21.0.0",
            "interface": "br-b236d7c30f50",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.22.0.0",
            "interface": "br-0d09f6c0e547",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
          "name": "Ubuntu",
          "versionId": "22.04",
          "versionCodename": "jammy",
          "id": "ubuntu",
          "idLike": "debian"
        }
      },
      {
        "hostname": "cm-web",
        "cpu": {
          "architecture": "x86_64",
          "cpus": 2,
          "cores": 2,
          "threads": 2,
          "maxSpeed": 3.7,
          "vendor": "GenuineIntel",
          "model": "Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"
        },
        "memory": {
          "type": "DDR4",
          "totalSize": 4,
          "available": 2,
          "used": 2
        },
        "rootDisk": {
          "label": "unknown",
          "type": "HDD",
          "totalSize": 1312,
          "available": 1229,
          "used": 16
        },
        "interfaces": [
          {
            "name": "lo",
            "ipv4CidrBlocks": ["127.0.0.1/8"],
            "ipv6CidrBlocks": ["::1/128"],
            "mtu": 65536
          },
          {
            "name": "eno1np0",
            "macAddress": "a4:bf:01:5a:b1:1b",
            "ipv4CidrBlocks": ["172.29.0.103/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b11b/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f0",
            "macAddress": "b4:96:91:53:01:58",
            "mtu": 1500
          },
          {
            "name": "eno2np1",
            "macAddress": "a4:bf:01:5a:b1:1c",
            "ipv4CidrBlocks": ["192.168.110.103/24"],
            "ipv6CidrBlocks": ["fe80::a6bf:1ff:fe5a:b11c/64"],
            "mtu": 1500,
            "state": "up"
          },
          {
            "name": "enp24s0f1",
            "macAddress": "b4:96:91:53:01:59",
            "mtu": 1500
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
            "name": "enp175s0f0",
            "macAddress": "b4:96:91:55:23:8c",
            "mtu": 1500
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
            "name": "docker0",
            "macAddress": "02:42:55:3d:6b:90",
            "ipv4CidrBlocks": ["172.17.0.1/16"],
            "mtu": 1500,
            "state": "down"
          }
        ],
        "routingTable": [
          {
            "destination": "0.0.0.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "172.17.0.0",
            "interface": "docker0",
            "protocol": "kernel",
            "scope": "link",
            "source": "172.17.0.1",
            "linkState": "down"
          },
          {
            "destination": "172.29.0.0",
            "interface": "eno1np0",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          },
          {
            "destination": "192.168.110.0",
            "gateway": "192.168.110.254",
            "interface": "eno2np1",
            "protocol": "kernel",
            "scope": "universe",
            "linkState": "up"
          }
        ],
        "os": {
          "prettyName": "Ubuntu 22.04.4 LTS",
          "version": "22.04.4 LTS (Jammy Jellyfish)",
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

```json
{
  "status": "ok",
  "description": "Target infra is recommended.",
  "targetInfra": {
    "name": "mmci01",
    "installMonAgent": "no",
    "label": null,
    "systemLabel": "",
    "description": "A cloud infra recommended by CM-Beetle",
    "vm": [
      {
        "name": "rehosted-cm-nfs",
        "subGroupSize": "",
        "label": null,
        "description": "a recommended virtual machine",
        "commonSpec": "ncpvpc+kr+svr.vsvr.hicpu.c002.m004.net.hdd.b050.g002",
        "commonImage": "ncpvpc+kr+ubuntu20.04"
      },
      {
        "name": "rehosted-cm-web",
        "subGroupSize": "",
        "label": null,
        "description": "a recommended virtual machine",
        "commonSpec": "ncpvpc+kr+svr.vsvr.hicpu.c002.m004.net.hdd.b050.g002",
        "commonImage": "ncpvpc+kr+ubuntu20.04"
      }
    ]
  }
}
```

### Migrate the computing infra as defined in the target model

- API: `POST /migration/ns/{nsId}/mci`
- nsId: `mig01` (default)
- Request body:

> [!NOTE]
> As you can see, the partial info of previous result is used.
> `mmci01` is used for the name of migrated computing infra.

```json
{
  "name": "mmci01",
  "installMonAgent": "no",
  "label": null,
  "systemLabel": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "name": "rehosted-cm-nfs",
      "subGroupSize": "",
      "label": null,
      "description": "a recommended virtual machine",
      "commonSpec": "ncpvpc+kr+svr.vsvr.hicpu.c002.m004.net.hdd.b050.g002",
      "commonImage": "ncpvpc+kr+ubuntu20.04"
    },
    {
      "name": "rehosted-cm-web",
      "subGroupSize": "",
      "label": null,
      "description": "a recommended virtual machine",
      "commonSpec": "ncpvpc+kr+svr.vsvr.hicpu.c002.m004.net.hdd.b050.g002",
      "commonImage": "ncpvpc+kr+ubuntu20.04"
    }
  ]
}
```

- Response body:

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "cstfo6equjj5l31al3pg",
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
  "installMonAgent": "no",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "A cloud infra recommended by CM-Beetle",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "cstfo6equjj5l31al3pg"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "resourceType": "vm",
      "id": "rehosted-cm-nfs-1",
      "uid": "cstfo6equjj5l31al3rg",
      "cspResourceName": "cstfo6equjj5l31al3rg",
      "cspResourceId": "100510302",
      "name": "rehosted-cm-nfs-1",
      "subGroupId": "rehosted-cm-nfs",
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
      "createdTime": "2024-11-18 08:39:54",
      "label": {
        "sys.connectionName": "ncpvpc-kr",
        "sys.createdTime": "2024-11-18 08:39:54",
        "sys.cspResourceId": "100510302",
        "sys.cspResourceName": "cstfo6equjj5l31al3rg",
        "sys.id": "rehosted-cm-nfs-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-nfs-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-nfs",
        "sys.uid": "cstfo6equjj5l31al3rg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "211.188.59.28",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.113.0.7",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": "50",
      "rootDeviceName": "/dev/xvda",
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
      "specId": "ncpvpc+kr+svr.vsvr.hicpu.c002.m004.net.hdd.b050.g002",
      "cspSpecName": "SVR.VSVR.HICPU.C002.M004.NET.HDD.B050.G002",
      "imageId": "ncpvpc+kr+ubuntu20.04",
      "cspImageName": "SW.VSVR.OS.LNX64.UBNTU.SVR2004.B050",
      "vNetId": "mig01-shared-ncpvpc-kr",
      "cspVNetId": "83297",
      "subnetId": "mig01-shared-ncpvpc-kr",
      "cspSubnetId": "186653",
      "networkInterface": "eth0",
      "securityGroupIds": ["mig01-shared-ncpvpc-kr"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-ncpvpc-kr",
      "cspSshKeyId": "cstc0l6qujj5l31ovfug",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "ServerInstanceType",
          "value": "High CPU"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize(GB)",
          "value": "4"
        },
        {
          "key": "PlatformType",
          "value": "Ubuntu Server 64 Bit"
        },
        {
          "key": "PublicIpID",
          "value": "100510390"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "rehosted-cm-web-1",
      "uid": "cstfo6equjj5l31al3qg",
      "cspResourceName": "cstfo6equjj5l31al3qg",
      "cspResourceId": "100510299",
      "name": "rehosted-cm-web-1",
      "subGroupId": "rehosted-cm-web",
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
      "createdTime": "2024-11-18 08:40:03",
      "label": {
        "sys.connectionName": "ncpvpc-kr",
        "sys.createdTime": "2024-11-18 08:40:03",
        "sys.cspResourceId": "100510299",
        "sys.cspResourceName": "cstfo6equjj5l31al3qg",
        "sys.id": "rehosted-cm-web-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-web-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-web",
        "sys.uid": "cstfo6equjj5l31al3qg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "211.188.59.29",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.113.0.6",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": "50",
      "rootDeviceName": "/dev/xvda",
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
      "specId": "ncpvpc+kr+svr.vsvr.hicpu.c002.m004.net.hdd.b050.g002",
      "cspSpecName": "SVR.VSVR.HICPU.C002.M004.NET.HDD.B050.G002",
      "imageId": "ncpvpc+kr+ubuntu20.04",
      "cspImageName": "SW.VSVR.OS.LNX64.UBNTU.SVR2004.B050",
      "vNetId": "mig01-shared-ncpvpc-kr",
      "cspVNetId": "83297",
      "subnetId": "mig01-shared-ncpvpc-kr",
      "cspSubnetId": "186653",
      "networkInterface": "eth0",
      "securityGroupIds": ["mig01-shared-ncpvpc-kr"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-ncpvpc-kr",
      "cspSshKeyId": "cstc0l6qujj5l31ovfug",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "ServerInstanceType",
          "value": "High CPU"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize(GB)",
          "value": "4"
        },
        {
          "key": "PlatformType",
          "value": "Ubuntu Server 64 Bit"
        },
        {
          "key": "PublicIpID",
          "value": "100510394"
        }
      ]
    }
  ],
  "newVmList": null
}
```

</details>

#### List the migrated computing infrastructures

- API: `GET /migration/ns/{nsId}/mci`
- nsId `mig01` (default)
- option: `id`
- Request body: none
- Response body:

```json
{
  "idList": ["mmci01"]
}
```

#### List the migrated computing infrastructures

- API: `GET /migration/ns/{nsId}/mci`
- nsId `mig01` (default)
- option: `info`
- Request body: none
- Response body:

<details>
	<summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "mci": [
    {
      "id": "mmci01",
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
      "installMonAgent": "no",
      "masterVmId": "rehosted-cm-nfs-1",
      "masterIp": "211.188.59.28",
      "masterSSHPort": "22",
      "label": null,
      "systemLabel": "",
      "vm": [
        {
          "id": "rehosted-cm-nfs-1",
          "cspResourceName": "cstfo6equjj5l31al3rg",
          "name": "rehosted-cm-nfs-1",
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "nativeStatus": "Running",
          "monAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2024-11-18 08:39:54",
          "publicIp": "211.188.59.28",
          "privateIp": "10.113.0.7",
          "sshPort": "22",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          }
        },
        {
          "id": "rehosted-cm-web-1",
          "cspResourceName": "cstfo6equjj5l31al3qg",
          "name": "rehosted-cm-web-1",
          "status": "Running",
          "targetStatus": "None",
          "targetAction": "None",
          "nativeStatus": "Running",
          "monAgentStatus": "notInstalled",
          "systemMessage": "",
          "createdTime": "2024-11-18 08:40:03",
          "publicIp": "211.188.59.29",
          "privateIp": "10.113.0.6",
          "sshPort": "22",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          }
        }
      ]
    }
  ]
}
```

</details>

### Get the migrated computing infra

- API: `GET /migration/ns/{nsId}/mci/{mciId}`
- nsId `mig01` (default)
- mciId `mmci01`(default)
- Request body: None

<details>
  <summary> <ins>Click to see the response body </ins> </summary>

```json
{
  "resourceType": "mci",
  "id": "mmci01",
  "uid": "cstfo6equjj5l31al3pg",
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
  "installMonAgent": "no",
  "configureCloudAdaptiveNetwork": "",
  "label": {
    "sys.description": "A cloud infra recommended by CM-Beetle",
    "sys.id": "mmci01",
    "sys.labelType": "mci",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mmci01",
    "sys.namespace": "mig01",
    "sys.uid": "cstfo6equjj5l31al3pg"
  },
  "systemLabel": "",
  "systemMessage": "",
  "description": "A cloud infra recommended by CM-Beetle",
  "vm": [
    {
      "resourceType": "vm",
      "id": "rehosted-cm-nfs-1",
      "uid": "cstfo6equjj5l31al3rg",
      "cspResourceName": "cstfo6equjj5l31al3rg",
      "cspResourceId": "100510302",
      "name": "rehosted-cm-nfs-1",
      "subGroupId": "rehosted-cm-nfs",
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
      "createdTime": "2024-11-18 08:39:54",
      "label": {
        "sys.connectionName": "ncpvpc-kr",
        "sys.createdTime": "2024-11-18 08:39:54",
        "sys.cspResourceId": "100510302",
        "sys.cspResourceName": "cstfo6equjj5l31al3rg",
        "sys.id": "rehosted-cm-nfs-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-nfs-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-nfs",
        "sys.uid": "cstfo6equjj5l31al3rg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "211.188.59.28",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.113.0.7",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": "50",
      "rootDeviceName": "/dev/xvda",
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
      "specId": "ncpvpc+kr+svr.vsvr.hicpu.c002.m004.net.hdd.b050.g002",
      "cspSpecName": "SVR.VSVR.HICPU.C002.M004.NET.HDD.B050.G002",
      "imageId": "ncpvpc+kr+ubuntu20.04",
      "cspImageName": "SW.VSVR.OS.LNX64.UBNTU.SVR2004.B050",
      "vNetId": "mig01-shared-ncpvpc-kr",
      "cspVNetId": "83297",
      "subnetId": "mig01-shared-ncpvpc-kr",
      "cspSubnetId": "186653",
      "networkInterface": "eth0",
      "securityGroupIds": ["mig01-shared-ncpvpc-kr"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-ncpvpc-kr",
      "cspSshKeyId": "cstc0l6qujj5l31ovfug",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "ServerInstanceType",
          "value": "High CPU"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize(GB)",
          "value": "4"
        },
        {
          "key": "PlatformType",
          "value": "Ubuntu Server 64 Bit"
        },
        {
          "key": "PublicIpID",
          "value": "100510390"
        }
      ]
    },
    {
      "resourceType": "vm",
      "id": "rehosted-cm-web-1",
      "uid": "cstfo6equjj5l31al3qg",
      "cspResourceName": "cstfo6equjj5l31al3qg",
      "cspResourceId": "100510299",
      "name": "rehosted-cm-web-1",
      "subGroupId": "rehosted-cm-web",
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
      "createdTime": "2024-11-18 08:40:03",
      "label": {
        "sys.connectionName": "ncpvpc-kr",
        "sys.createdTime": "2024-11-18 08:40:03",
        "sys.cspResourceId": "100510299",
        "sys.cspResourceName": "cstfo6equjj5l31al3qg",
        "sys.id": "rehosted-cm-web-1",
        "sys.labelType": "vm",
        "sys.manager": "cb-tumblebug",
        "sys.mciId": "mmci01",
        "sys.name": "rehosted-cm-web-1",
        "sys.namespace": "mig01",
        "sys.subGroupId": "rehosted-cm-web",
        "sys.uid": "cstfo6equjj5l31al3qg"
      },
      "description": "a recommended virtual machine",
      "region": {
        "Region": "KR",
        "Zone": "KR-1"
      },
      "publicIP": "211.188.59.29",
      "sshPort": "22",
      "publicDNS": "",
      "privateIP": "10.113.0.6",
      "privateDNS": "",
      "rootDiskType": "HDD",
      "rootDiskSize": "50",
      "rootDeviceName": "/dev/xvda",
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
      "specId": "ncpvpc+kr+svr.vsvr.hicpu.c002.m004.net.hdd.b050.g002",
      "cspSpecName": "SVR.VSVR.HICPU.C002.M004.NET.HDD.B050.G002",
      "imageId": "ncpvpc+kr+ubuntu20.04",
      "cspImageName": "SW.VSVR.OS.LNX64.UBNTU.SVR2004.B050",
      "vNetId": "mig01-shared-ncpvpc-kr",
      "cspVNetId": "83297",
      "subnetId": "mig01-shared-ncpvpc-kr",
      "cspSubnetId": "186653",
      "networkInterface": "eth0",
      "securityGroupIds": ["mig01-shared-ncpvpc-kr"],
      "dataDiskIds": null,
      "sshKeyId": "mig01-shared-ncpvpc-kr",
      "cspSshKeyId": "cstc0l6qujj5l31ovfug",
      "vmUserName": "cb-user",
      "addtionalDetails": [
        {
          "key": "ServerInstanceType",
          "value": "High CPU"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySize(GB)",
          "value": "4"
        },
        {
          "key": "PlatformType",
          "value": "Ubuntu Server 64 Bit"
        },
        {
          "key": "PublicIpID",
          "value": "100510394"
        }
      ]
    }
  ],
  "newVmList": null
}
```

</details>

### Delete the migrated computing infra

- API: `DELETE /migration/ns/{nsId}/mci/{mciId}`
- nsId: `mig01`
- mciId: `mmci01`
- Query param: `?action=terminate` (default)
- Request body: None
- Response body:

```json
{
  "success": true,
  "text": ""
}
```
