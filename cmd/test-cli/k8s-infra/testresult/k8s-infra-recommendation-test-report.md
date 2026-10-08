# CM-Beetle K8s Infra Recommendation Test Results

> [!NOTE]
> Verifies `POST /recommendation/k8sCluster` against on-premise scenario fixtures.
> No cloud resources are provisioned by this test.

## Environment

- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (0f10b3e)
- Git Commit: 0f10b3e
- Test Date: 2026-10-06 14:53:22 KST
- Targets (3): aws/ap-northeast-2, gcp/asia-northeast3, ncp/kr
- Scenarios: 20

## Summary Matrix

| Scenario | AWS-Seoul | GCP-Seoul | NCP-Seoul |
|---|---|---|---|
| `baseline` | ✅ 200 | ✅ 200 | ✅ 200 |
| `workers0` | ⚠️  500 | ⚠️  500 | ⚠️  500 |
| `workers1` | ✅ 200 | ✅ 200 | ✅ 200 |
| `workers5` | ✅ 200 | ✅ 200 | ✅ 200 |
| `tiny-upscale` | ✅ 200 | ✅ 200 | ✅ 200 |
| `three-groups` | ✅ 200 | ❌ 200 | ❌ 200 |
| `hetero-spec` | ✅ 200 | ✅ 200 | ✅ 200 |
| `converge-on-target-spec` | ✅ 200 | ✅ 200 | ✅ 200 |
| `k8s-api-collected-cpus0` | ✅ 200 | ✅ 200 | ✅ 200 |
| `mixed-arch` | ❌ 200 | ❌ 200 | ❌ 200 |
| `arm64` | ❌ 500 | ❌ 500 | ❌ 500 |
| `samecpu-diffdisk` | ✅ 200 | ✅ 200 | ✅ 200 |
| `spec-small` | ✅ 200 | ✅ 200 | ✅ 200 |
| `spec-large` | ✅ 200 | ✅ 200 | ✅ 200 |
| `version-1.34` | ✅ 200 | ✅ 200 | ✅ 200 |
| `version-1.99-fallback` | ✅ 200 | ✅ 200 | ✅ 200 |
| `no-worker-role` | ⚠️  500 | ⚠️  500 | ⚠️  500 |
| `neg-raw-source-group` | ✅ 400 | ✅ 400 | ✅ 400 |
| `neg-raw-connection-info` | ⚠️  500 | ⚠️  500 | ⚠️  500 |
| `neg-unwrapped-servers` | ⚠️  500 | ⚠️  500 | ⚠️  500 |

| Legend | Meaning |
|---|---|
| ✅ | Behaved as expected, with the status code the API declares |
| ⚠️ | Behaved as expected (input correctly accepted or rejected), but status code differs |
| ❌ | Did not behave as expected |

**Behaved as expected**: 52/60 ❌

- Fully conforming: 40
- Status code differs: 12
- Unexpected behaviour: 8

---

## Case Details

### baseline — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-refined-infra.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 41ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### workers0 — AWS-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers0.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 500
- **Duration**: 1ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "no worker nodes found in source K8s cluster",
  "success": false
}
```

</details>

---

### workers1 — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers1.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 22ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=1
- ✅ total node size: 1

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 1 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### workers5 — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers5.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 23ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=5
- ✅ total node size: 5

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 5 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (5 node(s))",
          "desiredNodeSize": 5,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 5,
          "minNodeSize": 5,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### tiny-upscale — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-tiny.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 21ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+t3a.medium image=default nodes=1
- ✅ total node size: 1

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 1 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s)) (worker spec upscaled from source 1vCPU/2GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+t3a.medium",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### three-groups — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-3groups.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 40ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ✅ node group count: 2
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.2xlarge image=default nodes=2
- ℹ️  node group[1] "workers2" spec=aws+ap-northeast-2+m5a.4xlarge image=default nodes=1
- ✅ total node size: 3

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 3 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s)) (consolidated into c5a.2xlarge: c5a.xlarge (4vCPU/8GiB) merged in, so these workers share one node group instead of 2. They are provisioned larger than their source, which is the cost of the smaller node group count.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.2xlarge",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers2",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+m5a.4xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### hetero-spec — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-hetero.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 27ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ✅ node group count: 2
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=1
- ℹ️  node group[1] "workers2" spec=aws+ap-northeast-2+m5a.4xlarge image=default nodes=1
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers2",
          "onAutoScaling": "true",
          "rootDiskSize": 300,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+m5a.4xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### converge-on-target-spec — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-converge.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 21ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+t3a.medium image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s)) (worker spec upscaled from source 1vCPU/2GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 300,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+t3a.medium",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "converge-k8s",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### k8s-api-collected-cpus0 — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-k8sapi-cpus0.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 23ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.2xlarge image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s)) (consolidated into c5a.2xlarge: t3a.xlarge (4vCPU/16GiB) merged in, so these workers share one node group instead of 2. They are provisioned larger than their source, which is the cost of the smaller node group count.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 300,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.2xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "k8sapi-cpus0-k8s",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### mixed-arch — AWS-Seoul (❌ UNEXPECTED)

- **Fixture**: `testconf/scenarios/honeybee-k8s-mixed-arch.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 35ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=1
- ❌ total node size 1, want 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34). 1 source worker(s) excluded: node image selection failed: no arm64 K8s node image is available for provider \"aws\" region \"ap-northeast-2\". Use x86_64 worker nodes, or check the available node images with searchImage(isKubernetesImage=true, provider=\"aws\") (worker00000000000000000000000000000002)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 1 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### arm64 — AWS-Seoul (❌ UNEXPECTED)

- **Fixture**: `testconf/scenarios/honeybee-k8s-arm64.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 500
- **Duration**: 20ms

**Checks**:

- ❌ input was rejected (HTTP 500) but a recommendation was expected

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "no K8s worker node group could be recommended: 2 source worker(s) excluded: node image selection failed: no arm64 K8s node image is available for provider \"aws\" region \"ap-northeast-2\". Use x86_64 worker nodes, or check the available node images with searchImage(isKubernetesImage=true, provider=\"aws\") (a1b2c3d4e5f647809abcdef012345678, c3d4e5f6a7b849012cdef345678901ab)",
  "success": false
}
```

</details>

---

### samecpu-diffdisk — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-samecpu-diffdisk.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 21ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 500,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### spec-small — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-spec-small.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 227ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+t3a.medium image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s)) (worker spec upscaled from source 1vCPU/1GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+t3a.medium",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### spec-large — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-spec-large.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 15ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+m5a.4xlarge image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+m5a.4xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### version-1.34 — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-v134.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 22ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.34.2 → target: v1.34)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.34.2, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### version-1.99-fallback — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-v199.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 21ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.36
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.99.0 → target: v1.36)",
    "status": "recommended",
    "targetCloud": {
      "csp": "aws",
      "region": "ap-northeast-2"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.99.0, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "true",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.36"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "aws-ap-northeast-2",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "aws-ap-northeast-2",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a",
          "zone": "ap-northeast-2a"
        },
        {
          "ipv4_CIDR": "10.0.2.0/24",
          "name": "k8s-subnet-b",
          "zone": "ap-northeast-2b"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### no-worker-role — AWS-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/honeybee-k8s-norole.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 500
- **Duration**: 1ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "no worker nodes found in source K8s cluster",
  "success": false
}
```

</details>

---

### neg-raw-source-group — AWS-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/negative/honeybee-k8s-raw-source-group.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 400
- **Duration**: 1ms

**Checks**:

- ✅ status code 400 as expected

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "Invalid request format",
  "success": false
}
```

</details>

---

### neg-raw-connection-info — AWS-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/negative/honeybee-k8s-raw-connection-info.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 500
- **Duration**: 1ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "source infra has no K8s cluster information",
  "success": false
}
```

</details>

---

### neg-unwrapped-servers — AWS-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/negative/beetle-k8s-recommendation-request.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 500
- **Duration**: 0s

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "source infra has no K8s cluster information",
  "success": false
}
```

</details>

---

### baseline — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-refined-infra.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 40ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-4 image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-4",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### workers0 — GCP-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers0.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 500
- **Duration**: 1ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "no worker nodes found in source K8s cluster",
  "success": false
}
```

</details>

---

### workers1 — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers1.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 20ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-4 image=default nodes=1
- ✅ total node size: 1

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 1 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-4",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### workers5 — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers5.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 20ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-4 image=default nodes=5
- ✅ total node size: 5

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 5 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (5 node(s))",
          "desiredNodeSize": 5,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-4",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### tiny-upscale — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-tiny.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 21ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+c4-standard-2 image=default nodes=1
- ✅ total node size: 1

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 1 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s)) (worker spec upscaled from source 1vCPU/2GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+c4-standard-2",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### three-groups — GCP-Seoul (❌ UNEXPECTED)

- **Fixture**: `testconf/scenarios/honeybee-k8s-3groups.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 61ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ❌ node group count 3, want 2
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-4 image=default nodes=1
- ℹ️  node group[1] "workers2" spec=gcp+asia-northeast3+n1-standard-8 image=default nodes=1
- ℹ️  node group[2] "workers3" spec=gcp+asia-northeast3+c4-standard-24 image=default nodes=1
- ✅ total node size: 3

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 3 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-4",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers2",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-8",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers3",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+c4-standard-24",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### hetero-spec — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-hetero.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 38ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ✅ node group count: 2
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-4 image=default nodes=1
- ℹ️  node group[1] "workers2" spec=gcp+asia-northeast3+c4-standard-24 image=default nodes=1
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-4",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers2",
          "onAutoScaling": "false",
          "rootDiskSize": 300,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+c4-standard-24",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### converge-on-target-spec — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-converge.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 19ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+c4-standard-2 image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s)) (worker spec upscaled from source 1vCPU/2GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 300,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+c4-standard-2",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "converge-k8s",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### k8s-api-collected-cpus0 — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-k8sapi-cpus0.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 25ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-8 image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s)) (consolidated into n1-standard-8: n1-highmem-4 (4vCPU/25GiB) merged in, so these workers share one node group instead of 2. They are provisioned larger than their source, which is the cost of the smaller node group count.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 300,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-8",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "k8sapi-cpus0-k8s",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### mixed-arch — GCP-Seoul (❌ UNEXPECTED)

- **Fixture**: `testconf/scenarios/honeybee-k8s-mixed-arch.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 412ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-4 image=default nodes=1
- ❌ total node size 1, want 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000). 1 source worker(s) excluded: spec recommendation failed: no spec found for K8s worker node (vCPU\u003e=4, memory\u003e=8GiB) in gcp asia-northeast3 after 5 attempts (worker00000000000000000000000000000002)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 1 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-4",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### arm64 — GCP-Seoul (❌ UNEXPECTED)

- **Fixture**: `testconf/scenarios/honeybee-k8s-arm64.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 500
- **Duration**: 195ms

**Checks**:

- ❌ input was rejected (HTTP 500) but a recommendation was expected

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "no K8s worker node group could be recommended: 2 source worker(s) excluded: spec recommendation failed: no spec found for K8s worker node (vCPU\u003e=4, memory\u003e=8GiB) in gcp asia-northeast3 after 5 attempts (a1b2c3d4e5f647809abcdef012345678, c3d4e5f6a7b849012cdef345678901ab)",
  "success": false
}
```

</details>

---

### samecpu-diffdisk — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-samecpu-diffdisk.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 80ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-4 image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 500,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-4",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### spec-small — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-spec-small.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 27ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+c4-standard-2 image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s)) (worker spec upscaled from source 1vCPU/1GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+c4-standard-2",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### spec-large — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-spec-large.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 19ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+c4-standard-24 image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+c4-standard-24",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### version-1.34 — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-v134.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 18ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.11-gke.1044000
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-4 image=default nodes=2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.34.2 → target: v1.34.11-gke.1044000)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.34.2, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-4",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.11-gke.1044000"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### version-1.99-fallback — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-v199.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 200
- **Duration**: 21ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.35.6-gke.1250001
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-4 image=default nodes=2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.99.0 → target: v1.35.6-gke.1250001)",
    "status": "recommended",
    "targetCloud": {
      "csp": "gcp",
      "region": "asia-northeast3"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.99.0, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "gcp+asia-northeast3+n1-standard-4",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.35.6-gke.1250001"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "gcp-asia-northeast3",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "gcp-asia-northeast3",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "gcp-asia-northeast3",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### no-worker-role — GCP-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/honeybee-k8s-norole.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 500
- **Duration**: 2ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "no worker nodes found in source K8s cluster",
  "success": false
}
```

</details>

---

### neg-raw-source-group — GCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/negative/honeybee-k8s-raw-source-group.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 400
- **Duration**: 1ms

**Checks**:

- ✅ status code 400 as expected

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "Invalid request format",
  "success": false
}
```

</details>

---

### neg-raw-connection-info — GCP-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/negative/honeybee-k8s-raw-connection-info.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 500
- **Duration**: 1ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "source infra has no K8s cluster information",
  "success": false
}
```

</details>

---

### neg-unwrapped-servers — GCP-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/negative/beetle-k8s-recommendation-request.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=gcp&desiredRegion=asia-northeast3`
- **Status Code**: 500
- **Duration**: 1ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "source infra has no K8s cluster information",
  "success": false
}
```

</details>

---

### baseline — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-refined-infra.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 30ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s4-g3a",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### workers0 — NCP-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers0.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 500
- **Duration**: 1ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "no worker nodes found in source K8s cluster",
  "success": false
}
```

</details>

---

### workers1 — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers1.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 20ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a image=default nodes=1
- ✅ total node size: 1

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 1 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s4-g3a",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### workers5 — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers5.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 20ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a image=default nodes=5
- ✅ total node size: 5

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 5 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (5 node(s))",
          "desiredNodeSize": 5,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s4-g3a",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### tiny-upscale — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-tiny.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 22ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s2-g3a image=default nodes=1
- ✅ total node size: 1

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 1 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s)) (worker spec upscaled from source 1vCPU/2GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s2-g3a",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### three-groups — NCP-Seoul (❌ UNEXPECTED)

- **Fixture**: `testconf/scenarios/honeybee-k8s-3groups.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 41ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ❌ node group count 3, want 2
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a image=default nodes=1
- ℹ️  node group[1] "workers2" spec=ncp+kr+s8-g3 image=default nodes=1
- ℹ️  node group[2] "workers3" spec=ncp+kr+s16-g3 image=default nodes=1
- ✅ total node size: 3

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 3 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s4-g3a",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers2",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s8-g3",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers3",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s16-g3",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### hetero-spec — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-hetero.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 25ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ✅ node group count: 2
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a image=default nodes=1
- ℹ️  node group[1] "workers2" spec=ncp+kr+s16-g3 image=default nodes=1
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s4-g3a",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers2",
          "onAutoScaling": "false",
          "rootDiskSize": 300,
          "rootDiskType": "default",
          "specId": "ncp+kr+s16-g3",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### converge-on-target-spec — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-converge.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 17ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s2-g3a image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s)) (worker spec upscaled from source 1vCPU/2GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 300,
          "rootDiskType": "default",
          "specId": "ncp+kr+s2-g3a",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "converge-k8s",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### k8s-api-collected-cpus0 — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-k8sapi-cpus0.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 20ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s8-g3 image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s)) (consolidated into s8-g3: s4-g3 (4vCPU/16GiB) merged in, so these workers share one node group instead of 2. They are provisioned larger than their source, which is the cost of the smaller node group count.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 300,
          "rootDiskType": "default",
          "specId": "ncp+kr+s8-g3",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "k8sapi-cpus0-k8s",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### mixed-arch — NCP-Seoul (❌ UNEXPECTED)

- **Fixture**: `testconf/scenarios/honeybee-k8s-mixed-arch.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 460ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a image=default nodes=1
- ❌ total node size 1, want 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2). 1 source worker(s) excluded: spec recommendation failed: no spec found for K8s worker node (vCPU\u003e=4, memory\u003e=8GiB) in ncp kr after 5 attempts (worker00000000000000000000000000000002)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 1 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (1 node(s))",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s4-g3a",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### arm64 — NCP-Seoul (❌ UNEXPECTED)

- **Fixture**: `testconf/scenarios/honeybee-k8s-arm64.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 500
- **Duration**: 195ms

**Checks**:

- ❌ input was rejected (HTTP 500) but a recommendation was expected

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "no K8s worker node group could be recommended: 2 source worker(s) excluded: spec recommendation failed: no spec found for K8s worker node (vCPU\u003e=4, memory\u003e=8GiB) in ncp kr after 5 attempts (a1b2c3d4e5f647809abcdef012345678, c3d4e5f6a7b849012cdef345678901ab)",
  "success": false
}
```

</details>

---

### samecpu-diffdisk — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-samecpu-diffdisk.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 79ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 500,
          "rootDiskType": "default",
          "specId": "ncp+kr+s4-g3a",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for worker00000000000000000000000000000001",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### spec-small — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-spec-small.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 26ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s2-g3a image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s)) (worker spec upscaled from source 1vCPU/1GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s2-g3a",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### spec-large — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-spec-large.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 7ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s16-g3 image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s16-g3",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### version-1.34 — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-v134.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 18ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.34.3-nks.2
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a image=default nodes=2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.34.2 → target: v1.34.3-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.34.2, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s4-g3a",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.34.3-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### version-1.99-fallback — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/honeybee-k8s-v199.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 200
- **Duration**: 19ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.36.2-nks.2
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a image=default nodes=2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for ncp kr (source: v1.99.0 → target: v1.36.2-nks.2)",
    "status": "recommended",
    "targetCloud": {
      "csp": "ncp",
      "region": "kr"
    },
    "targetInfra": {
      "description": "",
      "installMonAgent": "",
      "label": null,
      "name": "",
      "nodeGroups": null,
      "policyOnPartialFailure": "",
      "systemLabel": ""
    },
    "targetK8sCluster": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "Migrated from on-premise K8s cluster (v1.99.0, 2 workers)",
      "k8sNodeGroupList": [
        {
          "description": "Worker node group migrated from on-premise (2 node(s))",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 0,
          "minNodeSize": 0,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "ncp+kr+s4-g3a",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.36.2-nks.2"
    },
    "targetOsImageList": null,
    "targetSecurityGroupList": [
      {
        "connectionName": "ncp-kr",
        "cspResourceId": "",
        "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
        "firewallRules": [
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "22",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/24",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "30000-32767",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "6443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "443",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "inbound",
            "Ports": "10250",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "10.0.0.0/22",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "172.16.0.0/12",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "0.0.0.0/0",
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
          },
          {
            "CIDR": "192.168.0.0/16",
            "Direction": "inbound",
            "Ports": "1-65535",
            "Protocol": "UDP"
          }
        ],
        "name": "k8s-sg",
        "vNetId": "INSERT_YOUR_VNET_ID"
      }
    ],
    "targetSpecList": null,
    "targetSshKey": {
      "connectionName": "ncp-kr",
      "cspResourceId": "",
      "description": "SSH key for K8s worker nodes",
      "fingerprint": "",
      "name": "k8s-sshkey",
      "privateKey": "",
      "publicKey": "",
      "username": "",
      "verifiedUsername": ""
    },
    "targetVNet": {
      "cidrBlock": "10.0.0.0/22",
      "connectionName": "ncp-kr",
      "description": "VPC for migrated K8s cluster",
      "name": "k8s-vpc",
      "subnetInfoList": [
        {
          "ipv4_CIDR": "10.0.1.0/24",
          "name": "k8s-subnet-a"
        }
      ]
    }
  },
  "success": true
}
```

</details>

---

### no-worker-role — NCP-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/honeybee-k8s-norole.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 500
- **Duration**: 2ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "no worker nodes found in source K8s cluster",
  "success": false
}
```

</details>

---

### neg-raw-source-group — NCP-Seoul (✅ as expected)

- **Fixture**: `testconf/scenarios/negative/honeybee-k8s-raw-source-group.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 400
- **Duration**: 1ms

**Checks**:

- ✅ status code 400 as expected

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "Invalid request format",
  "success": false
}
```

</details>

---

### neg-raw-connection-info — NCP-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/negative/honeybee-k8s-raw-connection-info.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 500
- **Duration**: 1ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "source infra has no K8s cluster information",
  "success": false
}
```

</details>

---

### neg-unwrapped-servers — NCP-Seoul (⚠️  as expected, non-conforming status code)

- **Fixture**: `testconf/scenarios/negative/beetle-k8s-recommendation-request.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=ncp&desiredRegion=kr`
- **Status Code**: 500
- **Duration**: 1ms

**Checks**:

- ✅ input rejected as expected
- ⚠️  status code 500, but the API declares 400 for this case

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "error": "source infra has no K8s cluster information",
  "success": false
}
```

</details>

---

