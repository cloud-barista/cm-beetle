# CM-Beetle K8s Infra Recommendation Test Results

> [!NOTE]
> Verifies `POST /recommendation/k8sCluster` against on-premise scenario fixtures.
> No resources are provisioned by this test.

## Environment

- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.5.11+ (1131ca5)
- Git Commit: 1131ca5
- Test Date: 2026-08-19 11:36:41 KST
- Targets (1): aws/ap-northeast-2
- Scenarios: 18

## Summary Matrix

| Scenario | AWS-Seoul |
|---|---|
| `baseline` | ✅ 200 |
| `workers0` | ❌ 500 |
| `workers1` | ✅ 200 |
| `workers5` | ✅ 200 |
| `tiny-upscale` | ✅ 200 |
| `three-groups` | ✅ 200 |
| `hetero-spec` | ✅ 200 |
| `mixed-arch` | ✅ 200 |
| `arm64` | ✅ 200 |
| `samecpu-diffdisk` | ✅ 200 |
| `spec-small` | ✅ 200 |
| `spec-large` | ✅ 200 |
| `version-1.34` | ✅ 200 |
| `version-1.99-fallback` | ✅ 200 |
| `no-worker-role` | ❌ 500 |
| `neg-raw-source-group` | ✅ 400 |
| `neg-raw-connection-info` | ❌ 500 |
| `neg-unwrapped-servers` | ❌ 500 |

**Overall Result**: 14/18 cases passed ❌

---

## Case Details

### baseline — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-refined-infra.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 810ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (2 nodes)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "false",
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
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### workers0 — AWS-Seoul (❌ FAIL)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers0.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 500
- **Duration**: 0s

**Checks**:

- ❌ status code 500, want 400

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

### workers1 — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers1.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 12ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=1
- ✅ total node size: 1

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (1 nodes)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers1",
          "onAutoScaling": "false",
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
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### workers5 — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-workers5.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 17ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=5
- ✅ total node size: 5

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (5 nodes)",
          "desiredNodeSize": 5,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 5,
          "minNodeSize": 5,
          "name": "workers1",
          "onAutoScaling": "false",
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
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### tiny-upscale — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-tiny.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 419ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+t3a.medium image=default nodes=1
- ✅ total node size: 1

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (1 nodes) (worker spec upscaled from source 1vCPU/2GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers1",
          "onAutoScaling": "false",
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
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### three-groups — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-3groups.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 83ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ✅ node group count: 3
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=1
- ℹ️  node group[1] "workers2" spec=aws+ap-northeast-2+c5a.2xlarge image=default nodes=1
- ℹ️  node group[2] "workers3" spec=aws+ap-northeast-2+m5a.4xlarge image=default nodes=1
- ✅ total node size: 3

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (1 nodes)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 nodes)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers2",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.2xlarge",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 nodes)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers3",
          "onAutoScaling": "false",
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
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### hetero-spec — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-hetero.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 16ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ✅ node group count: 2
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=1
- ℹ️  node group[1] "workers2" spec=aws+ap-northeast-2+m5a.4xlarge image=default nodes=1
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (1 nodes)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 nodes)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers2",
          "onAutoScaling": "false",
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
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### mixed-arch — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-mixed-arch.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 23ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ℹ️  node group count: 2
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=1
- ℹ️  node group[1] "workers2" spec=aws+ap-northeast-2+c6g.xlarge image=AL2023_ARM_64_STANDARD nodes=1
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (1 nodes)",
          "desiredNodeSize": 1,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c5a.xlarge",
          "sshKeyId": ""
        },
        {
          "description": "Worker node group migrated from on-premise (1 nodes)",
          "desiredNodeSize": 1,
          "imageId": "AL2023_ARM_64_STANDARD",
          "label": null,
          "maxNodeSize": 1,
          "minNodeSize": 1,
          "name": "workers2",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c6g.xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### arm64 — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-arm64.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 13ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c6g.xlarge image=AL2023_ARM_64_STANDARD nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (2 nodes)",
          "desiredNodeSize": 2,
          "imageId": "AL2023_ARM_64_STANDARD",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "false",
          "rootDiskSize": 100,
          "rootDiskType": "default",
          "specId": "aws+ap-northeast-2+c6g.xlarge",
          "sshKeyId": ""
        }
      ],
      "label": null,
      "name": "on-prem-k8s-cluster",
      "securityGroupIds": null,
      "subnetIds": null,
      "systemLabel": "",
      "vNetId": "",
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### samecpu-diffdisk — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-samecpu-diffdisk.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 14ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (2 nodes)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "false",
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
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### spec-small — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-spec-small.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 16ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+t3a.medium image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (2 nodes) (worker spec upscaled from source 1vCPU/1GiB to 2vCPU/4GiB — the minimum node size accepted by the target K8s node recommendation. A node this small leaves little allocatable capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods may fail to schedule at the source size.)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "false",
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
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### spec-large — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-spec-large.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 6ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.33
- ✅ node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+m5a.4xlarge image=default nodes=2
- ✅ total node size: 2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
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
          "description": "Worker node group migrated from on-premise (2 nodes)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "false",
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
      "version": "1.33"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### version-1.34 — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-v134.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 13ms

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
          "description": "Worker node group migrated from on-premise (2 nodes)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "false",
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### version-1.99-fallback — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/honeybee-k8s-v199.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 200
- **Duration**: 19ms

**Checks**:

- ✅ status code 200 as expected
- ✅ version: 1.35
- ℹ️  node group count: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge image=default nodes=2

<details>
  <summary> <ins>Click to see the response body</ins> </summary>

```json
{
  "data": {
    "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.99.0 → target: v1.35)",
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
          "description": "Worker node group migrated from on-premise (2 nodes)",
          "desiredNodeSize": 2,
          "imageId": "default",
          "label": null,
          "maxNodeSize": 2,
          "minNodeSize": 2,
          "name": "workers1",
          "onAutoScaling": "false",
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
      "version": "1.35"
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
            "Direction": "outbound",
            "Ports": "1-65535",
            "Protocol": "TCP"
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

### no-worker-role — AWS-Seoul (❌ FAIL)

- **Fixture**: `testconf/scenarios/honeybee-k8s-norole.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 500
- **Duration**: 1ms

**Checks**:

- ❌ status code 500, want 400

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

### neg-raw-source-group — AWS-Seoul (✅ PASS)

- **Fixture**: `testconf/scenarios/negative/honeybee-k8s-raw-source-group.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 400
- **Duration**: 0s

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

### neg-raw-connection-info — AWS-Seoul (❌ FAIL)

- **Fixture**: `testconf/scenarios/negative/honeybee-k8s-raw-connection-info.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 500
- **Duration**: 0s

**Checks**:

- ❌ status code 500, want 400

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

### neg-unwrapped-servers — AWS-Seoul (❌ FAIL)

- **Fixture**: `testconf/scenarios/negative/beetle-k8s-recommendation-request.json`
- **Request**: `POST http://localhost:8056/beetle/recommendation/k8sCluster?desiredProvider=aws&desiredRegion=ap-northeast-2`
- **Status Code**: 500
- **Duration**: 0s

**Checks**:

- ❌ status code 500, want 400

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

