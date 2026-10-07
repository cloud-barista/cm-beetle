# CM-Beetle K8s Infra Migration Test Results — NCP-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → validate → migrate → list → get → workload → delete → residual.

## Environment

- CSP / Region: ncp / kr
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (108faf9)
- Git Commit: 108faf9
- Namespace: mig01
- Test Date: 2026-10-07 14:17:27 KST

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 18ms |
| 2 | POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation) | ✅ **PASS** | 24ms |

**Overall Result**: 2/2 steps passed ✅

**Total Duration**: 0s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 18ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.34.3-nks.2)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a nodes=2

### Step 2 — POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)

- **Duration**: 24ms
- **Status Code**: 200

- ✅ Target K8s infra model is valid for migration (0 issues)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
  "targetCloud": {
    "csp": "ncp",
    "region": "kr"
  },
  "targetInfra": {
    "name": "",
    "installMonAgent": "",
    "label": null,
    "systemLabel": "",
    "description": "",
    "nodeGroups": null,
    "policyOnPartialFailure": ""
  },
  "targetVNet": {
    "name": "k8s-vpc",
    "connectionName": "ncp-kr",
    "cidrBlock": "10.0.0.0/22",
    "subnetInfoList": [
      {
        "name": "k8s-subnet-a",
        "ipv4_CIDR": "10.0.1.0/24"
      }
    ],
    "description": "VPC for migrated K8s cluster"
  },
  "targetSshKey": {
    "name": "k8s-sshkey",
    "connectionName": "ncp-kr",
    "description": "SSH key for K8s worker nodes",
    "cspResourceId": "",
    "fingerprint": "",
    "username": "",
    "verifiedUsername": "",
    "publicKey": "",
    "privateKey": ""
  },
  "targetSpecList": null,
  "targetOsImageList": null,
  "targetSecurityGroupList": [
    {
      "name": "k8s-sg",
      "connectionName": "ncp-kr",
      "vNetId": "INSERT_YOUR_VNET_ID",
      "description": "Recommended security group for a1b2c3d4e5f647809abcdef012345678",
      "firewallRules": [
        {
          "Ports": "22",
          "Protocol": "TCP",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "10250",
          "Protocol": "TCP",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/24"
        },
        {
          "Ports": "30000-32767",
          "Protocol": "TCP",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "6443",
          "Protocol": "TCP",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "443",
          "Protocol": "TCP",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "10250",
          "Protocol": "TCP",
          "Direction": "inbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "1-65535",
          "Protocol": "TCP",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/22"
        },
        {
          "Ports": "1-65535",
          "Protocol": "UDP",
          "Direction": "inbound",
          "CIDR": "10.0.0.0/22"
        },
        {
          "Ports": "1-65535",
          "Protocol": "TCP",
          "Direction": "inbound",
          "CIDR": "172.16.0.0/12"
        },
        {
          "Ports": "1-65535",
          "Protocol": "UDP",
          "Direction": "inbound",
          "CIDR": "172.16.0.0/12"
        },
        {
          "Ports": "1-65535",
          "Protocol": "TCP",
          "Direction": "outbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "1-65535",
          "Protocol": "UDP",
          "Direction": "outbound",
          "CIDR": "0.0.0.0/0"
        },
        {
          "Ports": "1-65535",
          "Protocol": "TCP",
          "Direction": "inbound",
          "CIDR": "192.168.0.0/16"
        },
        {
          "Ports": "1-65535",
          "Protocol": "UDP",
          "Direction": "inbound",
          "CIDR": "192.168.0.0/16"
        }
      ],
      "cspResourceId": ""
    }
  ],
  "targetK8sCluster": {
    "connectionName": "ncp-kr",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "1.34.3-nks.2",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": [
      {
        "name": "workers1",
        "imageId": "default",
        "specId": "ncp+kr+s4-g3a",
        "rootDiskType": "default",
        "rootDiskSize": 100,
        "sshKeyId": "",
        "onAutoScaling": "false",
        "desiredNodeSize": 2,
        "minNodeSize": 0,
        "maxNodeSize": 0,
        "label": null,
        "description": "Worker node group migrated from on-premise (2 node(s))"
      }
    ],
    "cspResourceId": "",
    "label": null,
    "systemLabel": ""
  }
}
```

</details>

