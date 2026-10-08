# CM-Beetle K8s Infra Migration Test Results — GCP-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → validate → migrate → list → get → workload → delete → residual.

## Environment

- CSP / Region: gcp / asia-northeast3
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (cd1f3a2)
- Git Commit: cd1f3a2
- Namespace: mig01
- Test Date: 2026-10-07 16:41:25 KST
- Cluster ID: mig03-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 22ms |
| 2 | POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation) | ✅ **PASS** | 36ms |
| 3 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 4m0.069s |
| 4 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 3ms |
| 5 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 1.169s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 1m1.786s |
| 7 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 8m34.543s |
| 8 | Residual resource check (Tumblebug) | ✅ **PASS** | 4ms |

**Overall Result**: 8/8 steps passed ✅

**Total Duration**: 13m37s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 22ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.34)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+n1-standard-4 nodes=2

### Step 2 — POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)

- **Duration**: 36ms
- **Status Code**: 200

- ✅ Target K8s infra model is valid for migration (0 issues)

### Step 3 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 4m0.069s
- **Status Code**: 202

- ℹ️  nameSeed: mig03
- ℹ️  async reqId: 1791358885692729312
- ℹ️  cluster id: mig03-on-prem-k8s-cluster
- ℹ️  elapsed: 4m0s
- ✅ status: Active

### Step 4 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 3ms
- **Status Code**: 200

- ✅ migrated cluster present in list (8 total)

### Step 5 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 1.169s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=gcp+asia-northeast3+n1-standard-4, nodes=2)
- ✅ version: 1.34.11-gke.1044000 (recommended 1.34)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 1m1.786s

- ✅ kubeconfig obtained (server: https://34.64.51.188)
- ℹ️  auth method: exec credential plugin
- ✅ cluster token obtained from Tumblebug
- ✅ API server reachable (v1.34.11-gke.1044000)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 2)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: 34.47.81.95
- ✅ nginx served over the LoadBalancer at http://34.47.81.95/ (attempt 1)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 7 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 8m34.543s
- **Status Code**: 200

- ✅ deleted on attempt 1 (8m34s)

### Step 8 — Residual resource check (Tumblebug)

- **Duration**: 4ms

- ℹ️  VNet mig03-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig03-k8s-sg still exists (known gap)
- ℹ️  SshKey mig03-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.34)",
  "targetCloud": {
    "csp": "gcp",
    "region": "asia-northeast3"
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
    "connectionName": "gcp-asia-northeast3",
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
    "connectionName": "gcp-asia-northeast3",
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
      "connectionName": "gcp-asia-northeast3",
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
    "connectionName": "gcp-asia-northeast3",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "1.34",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": [
      {
        "name": "workers1",
        "imageId": "default",
        "specId": "gcp+asia-northeast3+n1-standard-4",
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

## Created Cluster

<details>
  <summary> <ins>Click to see the cluster info</ins> </summary>

```json
{
  "resourceType": "k8s",
  "id": "mig03-on-prem-k8s-cluster",
  "uid": "tbnr2ov8dm5t5gbqr6fo",
  "name": "mig03-on-prem-k8s-cluster",
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
  "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
  "systemMessage": "",
  "label": {
    "sys.connectionName": "gcp-asia-northeast3",
    "sys.createdTime": "2026-10-07 07:41:26 +0000 UTC",
    "sys.cspResourceId": "tbnr2ov8dm5t5gbqr6fo",
    "sys.cspResourceName": "tbnr2ov8dm5t5gbqr6fo",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig03-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig03-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbnr2ov8dm5t5gbqr6fo",
    "sys.version": "1.34.11-gke.1044000"
  },
  "systemLabel": "",
  "version": "1.34.11-gke.1044000",
  "network": {
    "vNetId": "mig03-k8s-vpc",
    "subnetIds": [
      "mig03-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "mig03-k8s-sg"
    ],
    "keyValueList": null
  },
  "k8sNodeGroupList": [
    {
      "id": "workers1",
      "name": "workers1",
      "imageId": "default",
      "specId": "gcp+asia-northeast3+n1-standard-4",
      "rootDiskType": "pd-balanced",
      "rootDiskSize": 100,
      "sshKeyId": "mig03-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 0,
      "maxNodeSize": 0,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-c7t0",
          "cspResourceId": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-c7t0"
        },
        {
          "cspResourceName": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-nwcg",
          "cspResourceId": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-nwcg"
        }
      ],
      "keyValueList": [
        {
          "key": "Config",
          "value": "{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:n1-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}}"
        },
        {
          "key": "Etag",
          "value": "4d9b274e-12d0-43bd-9eec-463d29afdd49"
        },
        {
          "key": "InitialNodeCount",
          "value": "2"
        },
        {
          "key": "InstanceGroupUrls",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-grp"
        },
        {
          "key": "Locations",
          "value": "asia-northeast3-a"
        },
        {
          "key": "Management",
          "value": "{autoRepair:true,autoUpgrade:true}"
        },
        {
          "key": "MaxPodsConstraint",
          "value": "{maxPodsPerNode:110}"
        },
        {
          "key": "Name",
          "value": "workers1"
        },
        {
          "key": "NetworkConfig",
          "value": "{networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podIpv4CidrBlock:10.180.0.0/14,podIpv4RangeUtilization:0.002,podRange:gke-tbnr2ov8dm5t5gbqr6fo-pods-b6e4e373,subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo}"
        },
        {
          "key": "PodIpv4CidrSize",
          "value": "24"
        },
        {
          "key": "SelfLink",
          "value": "https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tbnr2ov8dm5t5gbqr6fo/nodePools/workers1"
        },
        {
          "key": "Status",
          "value": "RUNNING"
        },
        {
          "key": "UpgradeSettings",
          "value": "{maxSurge:1,strategy:SURGE}"
        },
        {
          "key": "Version",
          "value": "1.34.11-gke.1044000"
        },
        {
          "key": "InstanceGroup_0",
          "value": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-grp"
        },
        {
          "key": "keypair",
          "value": "tb4jmnhc0vlt4svljbps"
        }
      ],
      "cspResourceName": "workers1",
      "cspResourceId": "workers1",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "workers1"
        },
        "ImageIID": {
          "NameId": "COS_CONTAINERD",
          "SystemId": "COS_CONTAINERD"
        },
        "VMSpecName": "n1-standard-4",
        "RootDiskType": "pd-balanced",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tb4jmnhc0vlt4svljbps",
          "SystemId": "tb4jmnhc0vlt4svljbps"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-c7t0",
            "SystemId": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-c7t0"
          },
          {
            "NameId": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-nwcg",
            "SystemId": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-nwcg"
          }
        ],
        "KeyValueList": [
          {
            "key": "Config",
            "value": "{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:n1-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}}"
          },
          {
            "key": "Etag",
            "value": "4d9b274e-12d0-43bd-9eec-463d29afdd49"
          },
          {
            "key": "InitialNodeCount",
            "value": "2"
          },
          {
            "key": "InstanceGroupUrls",
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-grp"
          },
          {
            "key": "Locations",
            "value": "asia-northeast3-a"
          },
          {
            "key": "Management",
            "value": "{autoRepair:true,autoUpgrade:true}"
          },
          {
            "key": "MaxPodsConstraint",
            "value": "{maxPodsPerNode:110}"
          },
          {
            "key": "Name",
            "value": "workers1"
          },
          {
            "key": "NetworkConfig",
            "value": "{networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podIpv4CidrBlock:10.180.0.0/14,podIpv4RangeUtilization:0.002,podRange:gke-tbnr2ov8dm5t5gbqr6fo-pods-b6e4e373,subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo}"
          },
          {
            "key": "PodIpv4CidrSize",
            "value": "24"
          },
          {
            "key": "SelfLink",
            "value": "https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tbnr2ov8dm5t5gbqr6fo/nodePools/workers1"
          },
          {
            "key": "Status",
            "value": "RUNNING"
          },
          {
            "key": "UpgradeSettings",
            "value": "{maxSurge:1,strategy:SURGE}"
          },
          {
            "key": "Version",
            "value": "1.34.11-gke.1044000"
          },
          {
            "key": "InstanceGroup_0",
            "value": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-grp"
          },
          {
            "key": "keypair",
            "value": "tb4jmnhc0vlt4svljbps"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "34.64.51.188",
    "kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://34.64.51.188\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVMRENDQXBTZ0F3SUJBZ0lRUXBwc0hTRnl0V2czYUNLUHYwaXIxekFOQmdrcWhraUc5dzBCQVFzRkFEQXYKTVMwd0t3WURWUVFERXlRek1qazBaR1ZtTkMwM1pEVmpMVFF6WWpjdE9EaGhZUzFsTmpZMVpEZGtOVGhoT1RBdwpJQmNOTWpZeE1EQTNNRFkwTVRJM1doZ1BNakExTmpBNU1qa3dOelF4TWpkYU1DOHhMVEFyQmdOVkJBTVRKRE15Ck9UUmtaV1kwTFRka05XTXRORE5pTnkwNE9HRmhMV1UyTmpWa04yUTFPR0U1TURDQ0FhSXdEUVlKS29aSWh2Y04KQVFFQkJRQURnZ0dQQURDQ0FZb0NnZ0dCQU94ZHdnaVNLNlVmQlBWdTJpcUZZRzVHbTlWN21rNEtYTzA4UElaWQpBQjYzYU11bnpkbVZXK2VBNGJ1c1FKYXk0Q0ZCbnJNdE9hdHBpQjdMWUsvN3BuN3ZrSWFHVGdCN1h5ZkNNSEoxCjlUVlMrQm1mUE1lLzNpenFhODdIbEhDM2NHK1BaeFU0aXhlWjJwdWZtQ1hSOUgvYzNmU20vejJkbHpYazZQK0gKRmR6SmhmWVF2U0NHV240c0k0TGhiT0lhQ2Y0cXRldTZWQzNudExjZTY0WFpqcmtPT0ZrRjhOaWFRaXV2bVN3YQo4Q21WL3gvYnNveVdabS9rdHZWaXZOaERwNDZiNkxWd0JpeWpzcWdUdEdEWWE3T1R6eWg0WDRRSmlYU0ljQWdxClJjd09DOEQrckNrbzhOcFdvVUVhTUJCb3I1UkxVZjI2VmdkU0FQSFoxLzlFOUR0Y1FDRWR2YWorUFRiZWJVa2gKMGVrK0dzai9Jem02cUVMWGxvbDhiYzcvaWhuQlF3NUtqNkdSdUF4NFVqNGZIcUl4OUVOWUx4bjVqenZhV0JvUAphL09aVnBVVjFRL2JXMkNicmd6bk5XY0MydlJTZkF0V25obTc2Mk1ST2J4cWRJVElVRDNYam1Ha3RQM29FKzFlCnNTL3hNLzN5dTFoTHpNdm5IcUZHY0JhbGxRSURBUUFCbzBJd1FEQU9CZ05WSFE4QkFmOEVCQU1DQWdRd0R3WUQKVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVwTzUwMmZML3dRdk9GQ0MxUjNRaW1YemxMczR3RFFZSgpLb1pJaHZjTkFRRUxCUUFEZ2dHQkFGN0xUQWZURUhmUjBwak5BRktwbHlTK0RZRVg3aExLM3JpcHE5cURibGVWClNvQXUrK2tzZi85SmtRZWJFS25vNC8rS0FJVEdzcHlBK3ZqbTZlNTZkdlpDWW5PdUpTc0pWSWNlYVQxOXJYM0cKaUVEQkx0Z3BnMFZseW1UNUdHN3hWWk1pcHBzb0JBY2phRFRVd0dza01ablFNdVdoMjQ1cW04cmNQT2dUeEdZNwpiOGVlREtQS2ZRb1pTQXJDYzViL1RWckowRE42NVJIWUQ0czB6bXBSU0Vndk5hbVlINXpvVG5TQ2xwcnB2WFM5CmhUQ1o4NFVPMHR5NGIremFOUHdQdzMzYzVkUVVWL2Y2aWlvSGJiTEVPdkZtMlFrcDlRQW50UHUxMEtJZ3VRbE8KcTBacVU2SUpKUWxsbnVEOUhzQjZ6S2JlZDNoY0FjSnlOQktsTVZxZVZIamxnY092WFQ0cVN3MkZvK0xMMXhMUgp3bGszdmlmOHg3akVXOXV5V0RiQis2WjVuZ2ZMWG1RM1FPRGJKa21jczRodzZ4OVRiWFRydTFORTR4N1ArV1EzCi96ZkNsT2tSdC8yZE1wRWorL2w5WWtWc1Z4VER3NVVmWTU5VHVqb0NGeDc3d09oQkdMSmdBdXF4c3VudmQ0TTEKNE9GUHdMSkhkL2JjbTQ1TjROc21wQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: gke_asia-northeast3-a_tbnr2ov8dm5t5gbqr6fo\ncontexts:\n- context:\n    cluster: gke_asia-northeast3-a_tbnr2ov8dm5t5gbqr6fo\n    user: gcp-dynamic-token\n  name: gke_asia-northeast3-a_tbnr2ov8dm5t5gbqr6fo\ncurrent-context: gke_asia-northeast3-a_tbnr2ov8dm5t5gbqr6fo\nusers:\n- name: gcp-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tbnr2ov8dm5t5gbqr6fo/token?ConnectionName=gcp-asia-northeast3\\\"\"\n"
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-10-07T07:41:26Z",
  "keyValueList": [
    {
      "key": "AddonsConfig",
      "value": "{dnsCacheConfig:{enabled:true},gcePersistentDiskCsiDriverConfig:{enabled:true},kubernetesDashboard:{disabled:true},networkPolicyConfig:{disabled:true}}"
    },
    {
      "key": "AnonymousAuthenticationConfig",
      "value": "{mode:ENABLED}"
    },
    {
      "key": "Autopilot",
      "value": "{}"
    },
    {
      "key": "Autoscaling",
      "value": "{autoprovisioningNodePoolDefaults:{imageType:COS_CONTAINERD,management:{autoRepair:true,autoUpgrade:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],serviceAccount:default},autoscalingProfile:BALANCED}"
    },
    {
      "key": "ClusterIpv4Cidr",
      "value": "10.180.0.0/14"
    },
    {
      "key": "ControlPlaneEndpointsConfig",
      "value": "{dnsEndpointConfig:{endpoint:gke-b6e4e373f0564f21b70fb4d5e71d63e42d4a-1064665102650.asia-northeast3-a.gke.goog},ipEndpointsConfig:{authorizedNetworksConfig:{},enablePublicEndpoint:true,enabled:true,privateEndpoint:10.0.1.11,publicEndpoint:34.64.51.188}}"
    },
    {
      "key": "CreateTime",
      "value": "2026-10-07T07:41:26+00:00"
    },
    {
      "key": "CurrentMasterVersion",
      "value": "1.34.11-gke.1044000"
    },
    {
      "key": "CurrentNodeCount",
      "value": "2"
    },
    {
      "key": "CurrentNodeVersion",
      "value": "1.34.11-gke.1044000"
    },
    {
      "key": "DatabaseEncryption",
      "value": "{currentState:CURRENT_STATE_DECRYPTED,state:DECRYPTED}"
    },
    {
      "key": "DefaultMaxPodsConstraint",
      "value": "{maxPodsPerNode:110}"
    },
    {
      "key": "EnableKubernetesAlpha",
      "value": "false"
    },
    {
      "key": "EnableTpu",
      "value": "false"
    },
    {
      "key": "Endpoint",
      "value": "34.64.51.188"
    },
    {
      "key": "EnterpriseConfig",
      "value": "{clusterTier:STANDARD}"
    },
    {
      "key": "Etag",
      "value": "7be10119-5698-4923-b82e-811e17532a16"
    },
    {
      "key": "Id",
      "value": "b6e4e373f0564f21b70fb4d5e71d63e42d4aea8938854933bf1ca3a7c6f5b039"
    },
    {
      "key": "InitialClusterVersion",
      "value": "1.34.11-gke.1044000"
    },
    {
      "key": "InitialNodeCount",
      "value": "0"
    },
    {
      "key": "InstanceGroupUrls",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-grp"
    },
    {
      "key": "IpAllocationPolicy",
      "value": "{clusterIpv4Cidr:10.180.0.0/14,clusterIpv4CidrBlock:10.180.0.0/14,clusterSecondaryRangeName:gke-tbnr2ov8dm5t5gbqr6fo-pods-b6e4e373,defaultPodIpv4RangeUtilization:0.002,networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podCidrOverprovisionConfig:{},servicesIpv4Cidr:34.118.224.0/20,servicesIpv4CidrBlock:34.118.224.0/20,stackType:IPV4,useIpAliases:true}"
    },
    {
      "key": "LabelFingerprint",
      "value": "17c2404d"
    },
    {
      "key": "LegacyAbac",
      "value": "{}"
    },
    {
      "key": "Location",
      "value": "asia-northeast3-a"
    },
    {
      "key": "Locations",
      "value": "asia-northeast3-a"
    },
    {
      "key": "LoggingConfig",
      "value": "{componentConfig:{enableComponents:[SYSTEM_COMPONENTS,WORKLOADS]}}"
    },
    {
      "key": "LoggingService",
      "value": "logging.googleapis.com/kubernetes"
    },
    {
      "key": "MaintenancePolicy",
      "value": "{resourceVersion:e3b0c442}"
    },
    {
      "key": "MasterAuth",
      "value": "{clientCertificateConfig:{},clusterCaCertificate:LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVMRENDQXBTZ0F3SUJBZ0lRUXBwc0hTRnl0V2czYUNLUHYwaXIxekFOQmdrcWhraUc5dzBCQVFzRkFEQXYKTVMwd0t3WURWUVFERXlRek1qazBaR1ZtTkMwM1pEVmpMVFF6WWpjdE9EaGhZUzFsTmpZMVpEZGtOVGhoT1RBdwpJQmNOTWpZeE1EQTNNRFkwTVRJM1doZ1BNakExTmpBNU1qa3dOelF4TWpkYU1DOHhMVEFyQmdOVkJBTVRKRE15Ck9UUmtaV1kwTFRka05XTXRORE5pTnkwNE9HRmhMV1UyTmpWa04yUTFPR0U1TURDQ0FhSXdEUVlKS29aSWh2Y04KQVFFQkJRQURnZ0dQQURDQ0FZb0NnZ0dCQU94ZHdnaVNLNlVmQlBWdTJpcUZZRzVHbTlWN21rNEtYTzA4UElaWQpBQjYzYU11bnpkbVZXK2VBNGJ1c1FKYXk0Q0ZCbnJNdE9hdHBpQjdMWUsvN3BuN3ZrSWFHVGdCN1h5ZkNNSEoxCjlUVlMrQm1mUE1lLzNpenFhODdIbEhDM2NHK1BaeFU0aXhlWjJwdWZtQ1hSOUgvYzNmU20vejJkbHpYazZQK0gKRmR6SmhmWVF2U0NHV240c0k0TGhiT0lhQ2Y0cXRldTZWQzNudExjZTY0WFpqcmtPT0ZrRjhOaWFRaXV2bVN3YQo4Q21WL3gvYnNveVdabS9rdHZWaXZOaERwNDZiNkxWd0JpeWpzcWdUdEdEWWE3T1R6eWg0WDRRSmlYU0ljQWdxClJjd09DOEQrckNrbzhOcFdvVUVhTUJCb3I1UkxVZjI2VmdkU0FQSFoxLzlFOUR0Y1FDRWR2YWorUFRiZWJVa2gKMGVrK0dzai9Jem02cUVMWGxvbDhiYzcvaWhuQlF3NUtqNkdSdUF4NFVqNGZIcUl4OUVOWUx4bjVqenZhV0JvUAphL09aVnBVVjFRL2JXMkNicmd6bk5XY0MydlJTZkF0V25obTc2Mk1ST2J4cWRJVElVRDNYam1Ha3RQM29FKzFlCnNTL3hNLzN5dTFoTHpNdm5IcUZHY0JhbGxRSURBUUFCbzBJd1FEQU9CZ05WSFE4QkFmOEVCQU1DQWdRd0R3WUQKVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVwTzUwMmZML3dRdk9GQ0MxUjNRaW1YemxMczR3RFFZSgpLb1pJaHZjTkFRRUxCUUFEZ2dHQkFGN0xUQWZURUhmUjBwak5BRktwbHlTK0RZRVg3aExLM3JpcHE5cURibGVWClNvQXUrK2tzZi85SmtRZWJFS25vNC8rS0FJVEdzcHlBK3ZqbTZlNTZkdlpDWW5PdUpTc0pWSWNlYVQxOXJYM0cKaUVEQkx0Z3BnMFZseW1UNUdHN3hWWk1pcHBzb0JBY2phRFRVd0dza01ablFNdVdoMjQ1cW04cmNQT2dUeEdZNwpiOGVlREtQS2ZRb1pTQXJDYzViL1RWckowRE42NVJIWUQ0czB6bXBSU0Vndk5hbVlINXpvVG5TQ2xwcnB2WFM5CmhUQ1o4NFVPMHR5NGIremFOUHdQdzMzYzVkUVVWL2Y2aWlvSGJiTEVPdkZtMlFrcDlRQW50UHUxMEtJZ3VRbE8KcTBacVU2SUpKUWxsbnVEOUhzQjZ6S2JlZDNoY0FjSnlOQktsTVZxZVZIamxnY092WFQ0cVN3MkZvK0xMMXhMUgp3bGszdmlmOHg3akVXOXV5V0RiQis2WjVuZ2ZMWG1RM1FPRGJKa21jczRodzZ4OVRiWFRydTFORTR4N1ArV1EzCi96ZkNsT2tSdC8yZE1wRWorL2w5WWtWc1Z4VER3NVVmWTU5VHVqb0NGeDc3d09oQkdMSmdBdXF4c3VudmQ0TTEKNE9GUHdMSkhkL2JjbTQ1TjROc21wQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K}"
    },
    {
      "key": "MasterAuthorizedNetworksConfig",
      "value": "{}"
    },
    {
      "key": "MonitoringConfig",
      "value": "{advancedDatapathObservabilityConfig:{},componentConfig:{enableComponents:[SYSTEM_COMPONENTS,STORAGE,HPA,POD,DAEMONSET,DEPLOYMENT,STATEFULSET,CADVISOR,KUBELET,DCGM,JOBSET]},managedPrometheusConfig:{enabled:true}}"
    },
    {
      "key": "MonitoringService",
      "value": "monitoring.googleapis.com/kubernetes"
    },
    {
      "key": "Name",
      "value": "tbnr2ov8dm5t5gbqr6fo"
    },
    {
      "key": "Network",
      "value": "tb8e4ivjhph8tj8eq3c5"
    },
    {
      "key": "NetworkConfig",
      "value": "{network:projects/GCP_PROJECT_ID/global/networks/tb8e4ivjhph8tj8eq3c5,serviceExternalIpsConfig:{},subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo}"
    },
    {
      "key": "NodeConfig",
      "value": "{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:n1-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}}"
    },
    {
      "key": "NodeIpv4CidrSize",
      "value": "0"
    },
    {
      "key": "NodePoolAutoConfig",
      "value": "{nodeKubeletConfig:{}}"
    },
    {
      "key": "NodePoolDefaults",
      "value": "{nodeConfigDefaults:{loggingConfig:{variantConfig:{variant:DEFAULT}},nodeKubeletConfig:{}}}"
    },
    {
      "key": "NodePools",
      "value": "{config:{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:n1-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}},etag:4d9b274e-12d0-43bd-9eec-463d29afdd49,initialNodeCount:2,instanceGroupUrls:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-grp],locations:[asia-northeast3-a],management:{autoRepair:true,autoUpgrade:true},maxPodsConstraint:{maxPodsPerNode:110},name:workers1,networkConfig:{networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podIpv4CidrBlock:10.180.0.0/14,podIpv4RangeUtilization:0.002,podRange:gke-tbnr2ov8dm5t5gbqr6fo-pods-b6e4e373,subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo},podIpv4CidrSize:24,selfLink:https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tbnr2ov8dm5t5gbqr6fo/nodePools/workers1,status:RUNNING,upgradeSettings:{maxSurge:1,strategy:SURGE},version:1.34.11-gke.1044000}"
    },
    {
      "key": "NotificationConfig",
      "value": "{pubsub:{}}"
    },
    {
      "key": "PodAutoscaling",
      "value": "{hpaProfile:PERFORMANCE}"
    },
    {
      "key": "PrivateClusterConfig",
      "value": "{privateEndpoint:10.0.1.11,publicEndpoint:34.64.51.188}"
    },
    {
      "key": "RbacBindingConfig",
      "value": "{enableInsecureBindingSystemAuthenticated:true,enableInsecureBindingSystemUnauthenticated:true}"
    },
    {
      "key": "ReleaseChannel",
      "value": "{channel:STABLE}"
    },
    {
      "key": "ResourceLabels",
      "value": "{cb-spider-pmks-securitygroup-0:tb5vuo31rrnng9587m1n}"
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
      "key": "SecurityPostureConfig",
      "value": "{mode:BASIC,vulnerabilityMode:VULNERABILITY_MODE_UNSPECIFIED}"
    },
    {
      "key": "SelfLink",
      "value": "https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tbnr2ov8dm5t5gbqr6fo"
    },
    {
      "key": "ServicesIpv4Cidr",
      "value": "34.118.224.0/20"
    },
    {
      "key": "ShieldedNodes",
      "value": "{enabled:true}"
    },
    {
      "key": "Status",
      "value": "RUNNING"
    },
    {
      "key": "Subnetwork",
      "value": "tbf74se4bdd1ilmjkavo"
    },
    {
      "key": "UserManagedKeysConfig",
      "value": "{}"
    },
    {
      "key": "Zone",
      "value": "asia-northeast3-a"
    }
  ],
  "cspResourceName": "tbnr2ov8dm5t5gbqr6fo",
  "cspResourceId": "tbnr2ov8dm5t5gbqr6fo",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbnr2ov8dm5t5gbqr6fo",
      "SystemId": "tbnr2ov8dm5t5gbqr6fo"
    },
    "Version": "1.34.11-gke.1044000",
    "Network": {
      "VpcIID": {
        "NameId": "tb8e4ivjhph8tj8eq3c5",
        "SystemId": "tb8e4ivjhph8tj8eq3c5"
      },
      "SubnetIIDs": [
        {
          "NameId": "tbf74se4bdd1ilmjkavo",
          "SystemId": "tbf74se4bdd1ilmjkavo"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "tb5vuo31rrnng9587m1n",
          "SystemId": "tb5vuo31rrnng9587m1n"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "workers1"
        },
        "ImageIID": {
          "NameId": "COS_CONTAINERD",
          "SystemId": "COS_CONTAINERD"
        },
        "VMSpecName": "n1-standard-4",
        "RootDiskType": "pd-balanced",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tb4jmnhc0vlt4svljbps",
          "SystemId": "tb4jmnhc0vlt4svljbps"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-c7t0",
            "SystemId": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-c7t0"
          },
          {
            "NameId": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-nwcg",
            "SystemId": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-nwcg"
          }
        ],
        "KeyValueList": [
          {
            "key": "Config",
            "value": "{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:n1-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}}"
          },
          {
            "key": "Etag",
            "value": "4d9b274e-12d0-43bd-9eec-463d29afdd49"
          },
          {
            "key": "InitialNodeCount",
            "value": "2"
          },
          {
            "key": "InstanceGroupUrls",
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-grp"
          },
          {
            "key": "Locations",
            "value": "asia-northeast3-a"
          },
          {
            "key": "Management",
            "value": "{autoRepair:true,autoUpgrade:true}"
          },
          {
            "key": "MaxPodsConstraint",
            "value": "{maxPodsPerNode:110}"
          },
          {
            "key": "Name",
            "value": "workers1"
          },
          {
            "key": "NetworkConfig",
            "value": "{networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podIpv4CidrBlock:10.180.0.0/14,podIpv4RangeUtilization:0.002,podRange:gke-tbnr2ov8dm5t5gbqr6fo-pods-b6e4e373,subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo}"
          },
          {
            "key": "PodIpv4CidrSize",
            "value": "24"
          },
          {
            "key": "SelfLink",
            "value": "https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tbnr2ov8dm5t5gbqr6fo/nodePools/workers1"
          },
          {
            "key": "Status",
            "value": "RUNNING"
          },
          {
            "key": "UpgradeSettings",
            "value": "{maxSurge:1,strategy:SURGE}"
          },
          {
            "key": "Version",
            "value": "1.34.11-gke.1044000"
          },
          {
            "key": "InstanceGroup_0",
            "value": "gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-grp"
          },
          {
            "key": "keypair",
            "value": "tb4jmnhc0vlt4svljbps"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "34.64.51.188",
      "Kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://34.64.51.188\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVMRENDQXBTZ0F3SUJBZ0lRUXBwc0hTRnl0V2czYUNLUHYwaXIxekFOQmdrcWhraUc5dzBCQVFzRkFEQXYKTVMwd0t3WURWUVFERXlRek1qazBaR1ZtTkMwM1pEVmpMVFF6WWpjdE9EaGhZUzFsTmpZMVpEZGtOVGhoT1RBdwpJQmNOTWpZeE1EQTNNRFkwTVRJM1doZ1BNakExTmpBNU1qa3dOelF4TWpkYU1DOHhMVEFyQmdOVkJBTVRKRE15Ck9UUmtaV1kwTFRka05XTXRORE5pTnkwNE9HRmhMV1UyTmpWa04yUTFPR0U1TURDQ0FhSXdEUVlKS29aSWh2Y04KQVFFQkJRQURnZ0dQQURDQ0FZb0NnZ0dCQU94ZHdnaVNLNlVmQlBWdTJpcUZZRzVHbTlWN21rNEtYTzA4UElaWQpBQjYzYU11bnpkbVZXK2VBNGJ1c1FKYXk0Q0ZCbnJNdE9hdHBpQjdMWUsvN3BuN3ZrSWFHVGdCN1h5ZkNNSEoxCjlUVlMrQm1mUE1lLzNpenFhODdIbEhDM2NHK1BaeFU0aXhlWjJwdWZtQ1hSOUgvYzNmU20vejJkbHpYazZQK0gKRmR6SmhmWVF2U0NHV240c0k0TGhiT0lhQ2Y0cXRldTZWQzNudExjZTY0WFpqcmtPT0ZrRjhOaWFRaXV2bVN3YQo4Q21WL3gvYnNveVdabS9rdHZWaXZOaERwNDZiNkxWd0JpeWpzcWdUdEdEWWE3T1R6eWg0WDRRSmlYU0ljQWdxClJjd09DOEQrckNrbzhOcFdvVUVhTUJCb3I1UkxVZjI2VmdkU0FQSFoxLzlFOUR0Y1FDRWR2YWorUFRiZWJVa2gKMGVrK0dzai9Jem02cUVMWGxvbDhiYzcvaWhuQlF3NUtqNkdSdUF4NFVqNGZIcUl4OUVOWUx4bjVqenZhV0JvUAphL09aVnBVVjFRL2JXMkNicmd6bk5XY0MydlJTZkF0V25obTc2Mk1ST2J4cWRJVElVRDNYam1Ha3RQM29FKzFlCnNTL3hNLzN5dTFoTHpNdm5IcUZHY0JhbGxRSURBUUFCbzBJd1FEQU9CZ05WSFE4QkFmOEVCQU1DQWdRd0R3WUQKVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVwTzUwMmZML3dRdk9GQ0MxUjNRaW1YemxMczR3RFFZSgpLb1pJaHZjTkFRRUxCUUFEZ2dHQkFGN0xUQWZURUhmUjBwak5BRktwbHlTK0RZRVg3aExLM3JpcHE5cURibGVWClNvQXUrK2tzZi85SmtRZWJFS25vNC8rS0FJVEdzcHlBK3ZqbTZlNTZkdlpDWW5PdUpTc0pWSWNlYVQxOXJYM0cKaUVEQkx0Z3BnMFZseW1UNUdHN3hWWk1pcHBzb0JBY2phRFRVd0dza01ablFNdVdoMjQ1cW04cmNQT2dUeEdZNwpiOGVlREtQS2ZRb1pTQXJDYzViL1RWckowRE42NVJIWUQ0czB6bXBSU0Vndk5hbVlINXpvVG5TQ2xwcnB2WFM5CmhUQ1o4NFVPMHR5NGIremFOUHdQdzMzYzVkUVVWL2Y2aWlvSGJiTEVPdkZtMlFrcDlRQW50UHUxMEtJZ3VRbE8KcTBacVU2SUpKUWxsbnVEOUhzQjZ6S2JlZDNoY0FjSnlOQktsTVZxZVZIamxnY092WFQ0cVN3MkZvK0xMMXhMUgp3bGszdmlmOHg3akVXOXV5V0RiQis2WjVuZ2ZMWG1RM1FPRGJKa21jczRodzZ4OVRiWFRydTFORTR4N1ArV1EzCi96ZkNsT2tSdC8yZE1wRWorL2w5WWtWc1Z4VER3NVVmWTU5VHVqb0NGeDc3d09oQkdMSmdBdXF4c3VudmQ0TTEKNE9GUHdMSkhkL2JjbTQ1TjROc21wQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: gke_asia-northeast3-a_tbnr2ov8dm5t5gbqr6fo\ncontexts:\n- context:\n    cluster: gke_asia-northeast3-a_tbnr2ov8dm5t5gbqr6fo\n    user: gcp-dynamic-token\n  name: gke_asia-northeast3-a_tbnr2ov8dm5t5gbqr6fo\ncurrent-context: gke_asia-northeast3-a_tbnr2ov8dm5t5gbqr6fo\nusers:\n- name: gcp-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tbnr2ov8dm5t5gbqr6fo/token?ConnectionName=gcp-asia-northeast3\\\"\"\n"
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-10-07T07:41:26Z",
    "KeyValueList": [
      {
        "key": "AddonsConfig",
        "value": "{dnsCacheConfig:{enabled:true},gcePersistentDiskCsiDriverConfig:{enabled:true},kubernetesDashboard:{disabled:true},networkPolicyConfig:{disabled:true}}"
      },
      {
        "key": "AnonymousAuthenticationConfig",
        "value": "{mode:ENABLED}"
      },
      {
        "key": "Autopilot",
        "value": "{}"
      },
      {
        "key": "Autoscaling",
        "value": "{autoprovisioningNodePoolDefaults:{imageType:COS_CONTAINERD,management:{autoRepair:true,autoUpgrade:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],serviceAccount:default},autoscalingProfile:BALANCED}"
      },
      {
        "key": "ClusterIpv4Cidr",
        "value": "10.180.0.0/14"
      },
      {
        "key": "ControlPlaneEndpointsConfig",
        "value": "{dnsEndpointConfig:{endpoint:gke-b6e4e373f0564f21b70fb4d5e71d63e42d4a-1064665102650.asia-northeast3-a.gke.goog},ipEndpointsConfig:{authorizedNetworksConfig:{},enablePublicEndpoint:true,enabled:true,privateEndpoint:10.0.1.11,publicEndpoint:34.64.51.188}}"
      },
      {
        "key": "CreateTime",
        "value": "2026-10-07T07:41:26+00:00"
      },
      {
        "key": "CurrentMasterVersion",
        "value": "1.34.11-gke.1044000"
      },
      {
        "key": "CurrentNodeCount",
        "value": "2"
      },
      {
        "key": "CurrentNodeVersion",
        "value": "1.34.11-gke.1044000"
      },
      {
        "key": "DatabaseEncryption",
        "value": "{currentState:CURRENT_STATE_DECRYPTED,state:DECRYPTED}"
      },
      {
        "key": "DefaultMaxPodsConstraint",
        "value": "{maxPodsPerNode:110}"
      },
      {
        "key": "EnableKubernetesAlpha",
        "value": "false"
      },
      {
        "key": "EnableTpu",
        "value": "false"
      },
      {
        "key": "Endpoint",
        "value": "34.64.51.188"
      },
      {
        "key": "EnterpriseConfig",
        "value": "{clusterTier:STANDARD}"
      },
      {
        "key": "Etag",
        "value": "7be10119-5698-4923-b82e-811e17532a16"
      },
      {
        "key": "Id",
        "value": "b6e4e373f0564f21b70fb4d5e71d63e42d4aea8938854933bf1ca3a7c6f5b039"
      },
      {
        "key": "InitialClusterVersion",
        "value": "1.34.11-gke.1044000"
      },
      {
        "key": "InitialNodeCount",
        "value": "0"
      },
      {
        "key": "InstanceGroupUrls",
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-grp"
      },
      {
        "key": "IpAllocationPolicy",
        "value": "{clusterIpv4Cidr:10.180.0.0/14,clusterIpv4CidrBlock:10.180.0.0/14,clusterSecondaryRangeName:gke-tbnr2ov8dm5t5gbqr6fo-pods-b6e4e373,defaultPodIpv4RangeUtilization:0.002,networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podCidrOverprovisionConfig:{},servicesIpv4Cidr:34.118.224.0/20,servicesIpv4CidrBlock:34.118.224.0/20,stackType:IPV4,useIpAliases:true}"
      },
      {
        "key": "LabelFingerprint",
        "value": "17c2404d"
      },
      {
        "key": "LegacyAbac",
        "value": "{}"
      },
      {
        "key": "Location",
        "value": "asia-northeast3-a"
      },
      {
        "key": "Locations",
        "value": "asia-northeast3-a"
      },
      {
        "key": "LoggingConfig",
        "value": "{componentConfig:{enableComponents:[SYSTEM_COMPONENTS,WORKLOADS]}}"
      },
      {
        "key": "LoggingService",
        "value": "logging.googleapis.com/kubernetes"
      },
      {
        "key": "MaintenancePolicy",
        "value": "{resourceVersion:e3b0c442}"
      },
      {
        "key": "MasterAuth",
        "value": "{clientCertificateConfig:{},clusterCaCertificate:LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVMRENDQXBTZ0F3SUJBZ0lRUXBwc0hTRnl0V2czYUNLUHYwaXIxekFOQmdrcWhraUc5dzBCQVFzRkFEQXYKTVMwd0t3WURWUVFERXlRek1qazBaR1ZtTkMwM1pEVmpMVFF6WWpjdE9EaGhZUzFsTmpZMVpEZGtOVGhoT1RBdwpJQmNOTWpZeE1EQTNNRFkwTVRJM1doZ1BNakExTmpBNU1qa3dOelF4TWpkYU1DOHhMVEFyQmdOVkJBTVRKRE15Ck9UUmtaV1kwTFRka05XTXRORE5pTnkwNE9HRmhMV1UyTmpWa04yUTFPR0U1TURDQ0FhSXdEUVlKS29aSWh2Y04KQVFFQkJRQURnZ0dQQURDQ0FZb0NnZ0dCQU94ZHdnaVNLNlVmQlBWdTJpcUZZRzVHbTlWN21rNEtYTzA4UElaWQpBQjYzYU11bnpkbVZXK2VBNGJ1c1FKYXk0Q0ZCbnJNdE9hdHBpQjdMWUsvN3BuN3ZrSWFHVGdCN1h5ZkNNSEoxCjlUVlMrQm1mUE1lLzNpenFhODdIbEhDM2NHK1BaeFU0aXhlWjJwdWZtQ1hSOUgvYzNmU20vejJkbHpYazZQK0gKRmR6SmhmWVF2U0NHV240c0k0TGhiT0lhQ2Y0cXRldTZWQzNudExjZTY0WFpqcmtPT0ZrRjhOaWFRaXV2bVN3YQo4Q21WL3gvYnNveVdabS9rdHZWaXZOaERwNDZiNkxWd0JpeWpzcWdUdEdEWWE3T1R6eWg0WDRRSmlYU0ljQWdxClJjd09DOEQrckNrbzhOcFdvVUVhTUJCb3I1UkxVZjI2VmdkU0FQSFoxLzlFOUR0Y1FDRWR2YWorUFRiZWJVa2gKMGVrK0dzai9Jem02cUVMWGxvbDhiYzcvaWhuQlF3NUtqNkdSdUF4NFVqNGZIcUl4OUVOWUx4bjVqenZhV0JvUAphL09aVnBVVjFRL2JXMkNicmd6bk5XY0MydlJTZkF0V25obTc2Mk1ST2J4cWRJVElVRDNYam1Ha3RQM29FKzFlCnNTL3hNLzN5dTFoTHpNdm5IcUZHY0JhbGxRSURBUUFCbzBJd1FEQU9CZ05WSFE4QkFmOEVCQU1DQWdRd0R3WUQKVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVwTzUwMmZML3dRdk9GQ0MxUjNRaW1YemxMczR3RFFZSgpLb1pJaHZjTkFRRUxCUUFEZ2dHQkFGN0xUQWZURUhmUjBwak5BRktwbHlTK0RZRVg3aExLM3JpcHE5cURibGVWClNvQXUrK2tzZi85SmtRZWJFS25vNC8rS0FJVEdzcHlBK3ZqbTZlNTZkdlpDWW5PdUpTc0pWSWNlYVQxOXJYM0cKaUVEQkx0Z3BnMFZseW1UNUdHN3hWWk1pcHBzb0JBY2phRFRVd0dza01ablFNdVdoMjQ1cW04cmNQT2dUeEdZNwpiOGVlREtQS2ZRb1pTQXJDYzViL1RWckowRE42NVJIWUQ0czB6bXBSU0Vndk5hbVlINXpvVG5TQ2xwcnB2WFM5CmhUQ1o4NFVPMHR5NGIremFOUHdQdzMzYzVkUVVWL2Y2aWlvSGJiTEVPdkZtMlFrcDlRQW50UHUxMEtJZ3VRbE8KcTBacVU2SUpKUWxsbnVEOUhzQjZ6S2JlZDNoY0FjSnlOQktsTVZxZVZIamxnY092WFQ0cVN3MkZvK0xMMXhMUgp3bGszdmlmOHg3akVXOXV5V0RiQis2WjVuZ2ZMWG1RM1FPRGJKa21jczRodzZ4OVRiWFRydTFORTR4N1ArV1EzCi96ZkNsT2tSdC8yZE1wRWorL2w5WWtWc1Z4VER3NVVmWTU5VHVqb0NGeDc3d09oQkdMSmdBdXF4c3VudmQ0TTEKNE9GUHdMSkhkL2JjbTQ1TjROc21wQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K}"
      },
      {
        "key": "MasterAuthorizedNetworksConfig",
        "value": "{}"
      },
      {
        "key": "MonitoringConfig",
        "value": "{advancedDatapathObservabilityConfig:{},componentConfig:{enableComponents:[SYSTEM_COMPONENTS,STORAGE,HPA,POD,DAEMONSET,DEPLOYMENT,STATEFULSET,CADVISOR,KUBELET,DCGM,JOBSET]},managedPrometheusConfig:{enabled:true}}"
      },
      {
        "key": "MonitoringService",
        "value": "monitoring.googleapis.com/kubernetes"
      },
      {
        "key": "Name",
        "value": "tbnr2ov8dm5t5gbqr6fo"
      },
      {
        "key": "Network",
        "value": "tb8e4ivjhph8tj8eq3c5"
      },
      {
        "key": "NetworkConfig",
        "value": "{network:projects/GCP_PROJECT_ID/global/networks/tb8e4ivjhph8tj8eq3c5,serviceExternalIpsConfig:{},subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo}"
      },
      {
        "key": "NodeConfig",
        "value": "{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:n1-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}}"
      },
      {
        "key": "NodeIpv4CidrSize",
        "value": "0"
      },
      {
        "key": "NodePoolAutoConfig",
        "value": "{nodeKubeletConfig:{}}"
      },
      {
        "key": "NodePoolDefaults",
        "value": "{nodeConfigDefaults:{loggingConfig:{variantConfig:{variant:DEFAULT}},nodeKubeletConfig:{}}}"
      },
      {
        "key": "NodePools",
        "value": "{config:{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:n1-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}},etag:4d9b274e-12d0-43bd-9eec-463d29afdd49,initialNodeCount:2,instanceGroupUrls:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tbnr2ov8dm5t5gbqr6fo-workers1-3d98af22-grp],locations:[asia-northeast3-a],management:{autoRepair:true,autoUpgrade:true},maxPodsConstraint:{maxPodsPerNode:110},name:workers1,networkConfig:{networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podIpv4CidrBlock:10.180.0.0/14,podIpv4RangeUtilization:0.002,podRange:gke-tbnr2ov8dm5t5gbqr6fo-pods-b6e4e373,subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo},podIpv4CidrSize:24,selfLink:https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tbnr2ov8dm5t5gbqr6fo/nodePools/workers1,status:RUNNING,upgradeSettings:{maxSurge:1,strategy:SURGE},version:1.34.11-gke.1044000}"
      },
      {
        "key": "NotificationConfig",
        "value": "{pubsub:{}}"
      },
      {
        "key": "PodAutoscaling",
        "value": "{hpaProfile:PERFORMANCE}"
      },
      {
        "key": "PrivateClusterConfig",
        "value": "{privateEndpoint:10.0.1.11,publicEndpoint:34.64.51.188}"
      },
      {
        "key": "RbacBindingConfig",
        "value": "{enableInsecureBindingSystemAuthenticated:true,enableInsecureBindingSystemUnauthenticated:true}"
      },
      {
        "key": "ReleaseChannel",
        "value": "{channel:STABLE}"
      },
      {
        "key": "ResourceLabels",
        "value": "{cb-spider-pmks-securitygroup-0:tb5vuo31rrnng9587m1n}"
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
        "key": "SecurityPostureConfig",
        "value": "{mode:BASIC,vulnerabilityMode:VULNERABILITY_MODE_UNSPECIFIED}"
      },
      {
        "key": "SelfLink",
        "value": "https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tbnr2ov8dm5t5gbqr6fo"
      },
      {
        "key": "ServicesIpv4Cidr",
        "value": "34.118.224.0/20"
      },
      {
        "key": "ShieldedNodes",
        "value": "{enabled:true}"
      },
      {
        "key": "Status",
        "value": "RUNNING"
      },
      {
        "key": "Subnetwork",
        "value": "tbf74se4bdd1ilmjkavo"
      },
      {
        "key": "UserManagedKeysConfig",
        "value": "{}"
      },
      {
        "key": "Zone",
        "value": "asia-northeast3-a"
      }
    ]
  }
}
```

</details>

