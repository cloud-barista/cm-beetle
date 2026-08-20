# CM-Beetle K8s Infra Migration Test Results — GCP-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → migrate → list → get (verified against
> the recommendation) → delete → residual resource check.

## Environment

- CSP / Region: gcp / asia-northeast3
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.0
- Git Commit: 803afb4
- Namespace: mig01
- Test Date: 2026-08-20 15:15:55 KST
- Cluster ID: mig03-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 17ms |
| 2 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 5m0.057s |
| 3 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 1ms |
| 4 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 6.698s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 1m4.399s |
| 6 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 9m14.565s |
| 7 | Residual resource check (Tumblebug) | ✅ **PASS** | 2ms |

**Overall Result**: 7/7 steps passed ✅

**Total Duration**: 15m25s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 17ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.33.12-gke.1000000)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=gcp+asia-northeast3+e2-standard-4 nodes=2

### Step 2 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 5m0.057s
- **Status Code**: 202

- ℹ️  nameSeed: mig03
- ℹ️  async reqId: 1787206555862907371
- ℹ️  cluster id: mig03-on-prem-k8s-cluster
- ℹ️  elapsed: 5m0s
- ✅ status: Active

### Step 3 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 1ms
- **Status Code**: 200

- ✅ migrated cluster present in list (3 total)

### Step 4 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 6.698s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=gcp+asia-northeast3+e2-standard-4, nodes=2)
- ✅ version: 1.33.12-gke.1000000 (recommended 1.33.12-gke.1000000)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 1m4.399s

- ✅ kubeconfig obtained (server: https://34.47.85.231)
- ℹ️  auth method: exec credential plugin
- ✅ cluster token obtained from Tumblebug
- ✅ API server reachable (v1.33.12-gke.1000000)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 1)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: 34.50.39.250
- ✅ nginx served over the LoadBalancer at http://34.50.39.250/ (attempt 1)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 6 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 9m14.565s
- **Status Code**: 200

- ✅ deleted on attempt 1 (9m14s)

### Step 7 — Residual resource check (Tumblebug)

- **Duration**: 2ms

- ℹ️  VNet mig03-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig03-k8s-sg still exists (known gap)
- ℹ️  SshKey mig03-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for gcp asia-northeast3 (source: v1.32.3 → target: v1.33.12-gke.1000000)",
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
          "Ports": "1-65535",
          "Protocol": "TCP",
          "Direction": "outbound",
          "CIDR": "0.0.0.0/0"
        }
      ],
      "cspResourceId": ""
    }
  ],
  "targetK8sCluster": {
    "connectionName": "gcp-asia-northeast3",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "1.33.12-gke.1000000",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": [
      {
        "name": "workers1",
        "imageId": "default",
        "specId": "gcp+asia-northeast3+e2-standard-4",
        "rootDiskType": "default",
        "rootDiskSize": 100,
        "sshKeyId": "",
        "onAutoScaling": "false",
        "desiredNodeSize": 2,
        "minNodeSize": 0,
        "maxNodeSize": 0,
        "label": null,
        "description": "Worker node group migrated from on-premise (2 nodes)"
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
  "uid": "tb4376joq2qqrv47qau4",
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
    "cb-spider-pmks-securitygroup-0": "tb5vuo31rrnng9587m1n",
    "sys.connectionName": "gcp-asia-northeast3",
    "sys.createdTime": "2026-08-20 06:15:59 +0000 UTC",
    "sys.cspResourceId": "tb4376joq2qqrv47qau4",
    "sys.cspResourceName": "tb4376joq2qqrv47qau4",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig03-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig03-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tb4376joq2qqrv47qau4",
    "sys.version": "1.33.12-gke.1000000"
  },
  "systemLabel": "",
  "version": "1.33.12-gke.1000000",
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
      "specId": "gcp+asia-northeast3+e2-standard-4",
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
          "cspResourceName": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-631x",
          "cspResourceId": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-631x"
        },
        {
          "cspResourceName": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-szqw",
          "cspResourceId": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-szqw"
        }
      ],
      "keyValueList": [
        {
          "key": "Config",
          "value": "{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:e2-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}}"
        },
        {
          "key": "Etag",
          "value": "d9c380dd-9dd4-4d32-8d99-abfb1c62ba68"
        },
        {
          "key": "InitialNodeCount",
          "value": "2"
        },
        {
          "key": "InstanceGroupUrls",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-grp"
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
          "value": "{networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podIpv4CidrBlock:10.8.0.0/14,podIpv4RangeUtilization:0.002,podRange:gke-tb4376joq2qqrv47qau4-pods-331fc160,subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo}"
        },
        {
          "key": "PodIpv4CidrSize",
          "value": "24"
        },
        {
          "key": "SelfLink",
          "value": "https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tb4376joq2qqrv47qau4/nodePools/workers1"
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
          "value": "1.33.12-gke.1000000"
        },
        {
          "key": "InstanceGroup_0",
          "value": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-grp"
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
        "VMSpecName": "e2-standard-4",
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
            "NameId": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-631x",
            "SystemId": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-631x"
          },
          {
            "NameId": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-szqw",
            "SystemId": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-szqw"
          }
        ],
        "KeyValueList": [
          {
            "key": "Config",
            "value": "{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:e2-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}}"
          },
          {
            "key": "Etag",
            "value": "d9c380dd-9dd4-4d32-8d99-abfb1c62ba68"
          },
          {
            "key": "InitialNodeCount",
            "value": "2"
          },
          {
            "key": "InstanceGroupUrls",
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-grp"
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
            "value": "{networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podIpv4CidrBlock:10.8.0.0/14,podIpv4RangeUtilization:0.002,podRange:gke-tb4376joq2qqrv47qau4-pods-331fc160,subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo}"
          },
          {
            "key": "PodIpv4CidrSize",
            "value": "24"
          },
          {
            "key": "SelfLink",
            "value": "https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tb4376joq2qqrv47qau4/nodePools/workers1"
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
            "value": "1.33.12-gke.1000000"
          },
          {
            "key": "InstanceGroup_0",
            "value": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-grp"
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
    "endpoint": "34.47.85.231",
    "kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://34.47.85.231\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVMVENDQXBXZ0F3SUJBZ0lSQUxJZTRTMTBNaTBFcWFRK3ZoZzlqUVV3RFFZSktvWklodmNOQVFFTEJRQXcKTHpFdE1Dc0dBMVVFQXhNa1l6RXdZVFprTldZdFpUTm1PQzAwTVRka0xXRmtPVGd0T1dRNE0yRTROVEV6T1RBegpNQ0FYRFRJMk1EZ3lNREExTVRZd01Gb1lEekl3TlRZd09ERXlNRFl4TmpBd1dqQXZNUzB3S3dZRFZRUURFeVJqCk1UQmhObVExWmkxbE0yWTRMVFF4TjJRdFlXUTVPQzA1WkRnellUZzFNVE01TURNd2dnR2lNQTBHQ1NxR1NJYjMKRFFFQkFRVUFBNElCandBd2dnR0tBb0lCZ1FDaCtzWlRYTFJ0bk5hY284V1h5dmhuUUw3L1NKUmd2T042QnRZTQo0SjdBVSsyNE1IWlhQdEdXdVdYaU96dDV2Njh0c1I3TUZMcnFOaEp4M05uY21IbFRBbmI2TDgzcDBFNXhKeTBOCmZkSDdSSkdEKzJ5dmFCMGRVQTZPTkZQTjQ1YTduY1ZvOEowSnI4aTkxdUVyYzE3N3pxdlo1WWh0c0VxS1dMM2kKY3hrVVZhdkJxS2pNQi9hOFpQUC90eDRaeDBwNWxudDBKVUcyMWZWYlZiKytYUjc4OE5zbkVONkZWdWw1R2NwWgpkSVlGVE1IUjJtbkx1VnlFb3d1eFk2YnBoSVJaendFRnJvdmxaK1Z5Z1F3czlURnlyMWc1VHVhVTd0S1dXbVpoCmx2bnY1amNuVlBFeGZ4MXR1RXdNM0xDZGZCMngyYUIrSDRnelBWTUZlakE1THRDcnhtM3llR2I1aStNZlpGYWYKNHViNWM2R1llUnUwWGZFekQ4dXFyT0JVVm9HTHNtT25CTW16d1hEeFE5MDVmU3g2bGxWTXRpbmoyMFkvYklrRwpmc0pSek44cDBUWTV0cUN5bGJsc3E3WUJiMXA5UHpZZUE0clBjNERvTmJMM3E5SXJqeXdpc0poRHo2ZGZ1dVlwCkZpSHhweWlXaGNKMVlzam1iTmtQcmRjVlhJc0NBd0VBQWFOQ01FQXdEZ1lEVlIwUEFRSC9CQVFEQWdJRU1BOEcKQTFVZEV3RUIvd1FGTUFNQkFmOHdIUVlEVlIwT0JCWUVGUHpKYi9pZis2TzFISnJLeHNaemFaeDhqQTljTUEwRwpDU3FHU0liM0RRRUJDd1VBQTRJQmdRQlNsajJDaXZKa0x5REZHZThHN0ZuY1BvMXdhS3NPaktmZE5pRGpXZjNNCk9Hc1d6QWgxYW5DTVB2d3o0NzNEUHRuYzJQU3pBWDJ6aGxGYWVwcUJSSlgzeDJLOXZPbHhTaUtkRzMyeXA2alcKSVRjd21QcWx4N0NrV0tNSzJGZFBaQ2IyUm9nS3Fwc2lLN2dLbng3UGtMNzdINnpWSGJ2SUlWVHM2QjA0RVNYcQprZEZTLzhOWDhxbFoxd0V4TjNkOVFEYVVrRjBUUk90REVNdDJ0NWowRGtzMzQwcXhlNlVJZDArL0pvSDRXMXk3CnZDM0FuV0lJc3BKTDZ3dDEwN09CbUdsaEJpSC8vUHlrYjNwbzdERzEzdk9TYXhFVGJObjVHZjMzQ2hoVEthVGQKM2p2OC92cU4wQXRDY3ZoRWY0SHI2YmlUY2R2ZXBFenY5Z2dWajVXVzA3eEVPVlRlbko1MGxNZnpNaGtGYUFYRApLaVduckNlRUg1anhNUDZIRVI4MGtOVmd0bk5HZDJyeHdqS21naUFKN25nUnEybGNhQmhwMzI3YVFadmt0T3V4CmFpQ28yT3VHbllLWkV2OUNyMVdpaERMNW1INm4rMHZTbGdwdkZqdVl0ZDkvR25uNnJoTy9CZ3Q1Z2dWcFFuTzkKQTVXZnU1dmZFaTBXcjlrb2xxYWxWdFE9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: gke_asia-northeast3-a_tb4376joq2qqrv47qau4\ncontexts:\n- context:\n    cluster: gke_asia-northeast3-a_tb4376joq2qqrv47qau4\n    user: gcp-dynamic-token\n  name: gke_asia-northeast3-a_tb4376joq2qqrv47qau4\ncurrent-context: gke_asia-northeast3-a_tb4376joq2qqrv47qau4\nusers:\n- name: gcp-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tb4376joq2qqrv47qau4/token?ConnectionName=gcp-asia-northeast3\\\"\"\n"
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-08-20T06:15:59Z",
  "keyValueList": [
    {
      "key": "AddonsConfig",
      "value": "{gcePersistentDiskCsiDriverConfig:{enabled:true},kubernetesDashboard:{disabled:true},networkPolicyConfig:{disabled:true}}"
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
      "value": "10.8.0.0/14"
    },
    {
      "key": "ControlPlaneEndpointsConfig",
      "value": "{dnsEndpointConfig:{endpoint:gke-331fc160950b490f8aa57aa47d944bf5b40b-1064665102650.asia-northeast3-a.gke.goog},ipEndpointsConfig:{authorizedNetworksConfig:{},enablePublicEndpoint:true,enabled:true,privateEndpoint:10.0.1.8,publicEndpoint:34.47.85.231}}"
    },
    {
      "key": "CreateTime",
      "value": "2026-08-20T06:15:59+00:00"
    },
    {
      "key": "CurrentMasterVersion",
      "value": "1.33.12-gke.1000000"
    },
    {
      "key": "CurrentNodeCount",
      "value": "2"
    },
    {
      "key": "CurrentNodeVersion",
      "value": "1.33.12-gke.1000000"
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
      "value": "34.47.85.231"
    },
    {
      "key": "EnterpriseConfig",
      "value": "{clusterTier:STANDARD}"
    },
    {
      "key": "Etag",
      "value": "fb3b6858-b12a-4c44-b14d-d17d2f1e2744"
    },
    {
      "key": "Id",
      "value": "331fc160950b490f8aa57aa47d944bf5b40b2562089242069dca40971400e3ab"
    },
    {
      "key": "InitialClusterVersion",
      "value": "1.33.12-gke.1000000"
    },
    {
      "key": "InitialNodeCount",
      "value": "0"
    },
    {
      "key": "InstanceGroupUrls",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-grp"
    },
    {
      "key": "IpAllocationPolicy",
      "value": "{clusterIpv4Cidr:10.8.0.0/14,clusterIpv4CidrBlock:10.8.0.0/14,clusterSecondaryRangeName:gke-tb4376joq2qqrv47qau4-pods-331fc160,defaultPodIpv4RangeUtilization:0.002,networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podCidrOverprovisionConfig:{},servicesIpv4Cidr:34.118.224.0/20,servicesIpv4CidrBlock:34.118.224.0/20,stackType:IPV4,useIpAliases:true}"
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
      "value": "{clientCertificateConfig:{},clusterCaCertificate:LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVMVENDQXBXZ0F3SUJBZ0lSQUxJZTRTMTBNaTBFcWFRK3ZoZzlqUVV3RFFZSktvWklodmNOQVFFTEJRQXcKTHpFdE1Dc0dBMVVFQXhNa1l6RXdZVFprTldZdFpUTm1PQzAwTVRka0xXRmtPVGd0T1dRNE0yRTROVEV6T1RBegpNQ0FYRFRJMk1EZ3lNREExTVRZd01Gb1lEekl3TlRZd09ERXlNRFl4TmpBd1dqQXZNUzB3S3dZRFZRUURFeVJqCk1UQmhObVExWmkxbE0yWTRMVFF4TjJRdFlXUTVPQzA1WkRnellUZzFNVE01TURNd2dnR2lNQTBHQ1NxR1NJYjMKRFFFQkFRVUFBNElCandBd2dnR0tBb0lCZ1FDaCtzWlRYTFJ0bk5hY284V1h5dmhuUUw3L1NKUmd2T042QnRZTQo0SjdBVSsyNE1IWlhQdEdXdVdYaU96dDV2Njh0c1I3TUZMcnFOaEp4M05uY21IbFRBbmI2TDgzcDBFNXhKeTBOCmZkSDdSSkdEKzJ5dmFCMGRVQTZPTkZQTjQ1YTduY1ZvOEowSnI4aTkxdUVyYzE3N3pxdlo1WWh0c0VxS1dMM2kKY3hrVVZhdkJxS2pNQi9hOFpQUC90eDRaeDBwNWxudDBKVUcyMWZWYlZiKytYUjc4OE5zbkVONkZWdWw1R2NwWgpkSVlGVE1IUjJtbkx1VnlFb3d1eFk2YnBoSVJaendFRnJvdmxaK1Z5Z1F3czlURnlyMWc1VHVhVTd0S1dXbVpoCmx2bnY1amNuVlBFeGZ4MXR1RXdNM0xDZGZCMngyYUIrSDRnelBWTUZlakE1THRDcnhtM3llR2I1aStNZlpGYWYKNHViNWM2R1llUnUwWGZFekQ4dXFyT0JVVm9HTHNtT25CTW16d1hEeFE5MDVmU3g2bGxWTXRpbmoyMFkvYklrRwpmc0pSek44cDBUWTV0cUN5bGJsc3E3WUJiMXA5UHpZZUE0clBjNERvTmJMM3E5SXJqeXdpc0poRHo2ZGZ1dVlwCkZpSHhweWlXaGNKMVlzam1iTmtQcmRjVlhJc0NBd0VBQWFOQ01FQXdEZ1lEVlIwUEFRSC9CQVFEQWdJRU1BOEcKQTFVZEV3RUIvd1FGTUFNQkFmOHdIUVlEVlIwT0JCWUVGUHpKYi9pZis2TzFISnJLeHNaemFaeDhqQTljTUEwRwpDU3FHU0liM0RRRUJDd1VBQTRJQmdRQlNsajJDaXZKa0x5REZHZThHN0ZuY1BvMXdhS3NPaktmZE5pRGpXZjNNCk9Hc1d6QWgxYW5DTVB2d3o0NzNEUHRuYzJQU3pBWDJ6aGxGYWVwcUJSSlgzeDJLOXZPbHhTaUtkRzMyeXA2alcKSVRjd21QcWx4N0NrV0tNSzJGZFBaQ2IyUm9nS3Fwc2lLN2dLbng3UGtMNzdINnpWSGJ2SUlWVHM2QjA0RVNYcQprZEZTLzhOWDhxbFoxd0V4TjNkOVFEYVVrRjBUUk90REVNdDJ0NWowRGtzMzQwcXhlNlVJZDArL0pvSDRXMXk3CnZDM0FuV0lJc3BKTDZ3dDEwN09CbUdsaEJpSC8vUHlrYjNwbzdERzEzdk9TYXhFVGJObjVHZjMzQ2hoVEthVGQKM2p2OC92cU4wQXRDY3ZoRWY0SHI2YmlUY2R2ZXBFenY5Z2dWajVXVzA3eEVPVlRlbko1MGxNZnpNaGtGYUFYRApLaVduckNlRUg1anhNUDZIRVI4MGtOVmd0bk5HZDJyeHdqS21naUFKN25nUnEybGNhQmhwMzI3YVFadmt0T3V4CmFpQ28yT3VHbllLWkV2OUNyMVdpaERMNW1INm4rMHZTbGdwdkZqdVl0ZDkvR25uNnJoTy9CZ3Q1Z2dWcFFuTzkKQTVXZnU1dmZFaTBXcjlrb2xxYWxWdFE9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K}"
    },
    {
      "key": "MasterAuthorizedNetworksConfig",
      "value": "{}"
    },
    {
      "key": "MonitoringConfig",
      "value": "{advancedDatapathObservabilityConfig:{},componentConfig:{enableComponents:[SYSTEM_COMPONENTS,STATEFULSET,JOBSET,STORAGE,HPA,POD,DAEMONSET,DEPLOYMENT,CADVISOR,KUBELET,DCGM]},managedPrometheusConfig:{enabled:true}}"
    },
    {
      "key": "MonitoringService",
      "value": "monitoring.googleapis.com/kubernetes"
    },
    {
      "key": "Name",
      "value": "tb4376joq2qqrv47qau4"
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
      "value": "{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:e2-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}}"
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
      "value": "{config:{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:e2-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}},etag:d9c380dd-9dd4-4d32-8d99-abfb1c62ba68,initialNodeCount:2,instanceGroupUrls:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-grp],locations:[asia-northeast3-a],management:{autoRepair:true,autoUpgrade:true},maxPodsConstraint:{maxPodsPerNode:110},name:workers1,networkConfig:{networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podIpv4CidrBlock:10.8.0.0/14,podIpv4RangeUtilization:0.002,podRange:gke-tb4376joq2qqrv47qau4-pods-331fc160,subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo},podIpv4CidrSize:24,selfLink:https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tb4376joq2qqrv47qau4/nodePools/workers1,status:RUNNING,upgradeSettings:{maxSurge:1,strategy:SURGE},version:1.33.12-gke.1000000}"
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
      "value": "{privateEndpoint:10.0.1.8,publicEndpoint:34.47.85.231}"
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
      "value": "https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tb4376joq2qqrv47qau4"
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
  "cspResourceName": "tb4376joq2qqrv47qau4",
  "cspResourceId": "tb4376joq2qqrv47qau4",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tb4376joq2qqrv47qau4",
      "SystemId": "tb4376joq2qqrv47qau4"
    },
    "Version": "1.33.12-gke.1000000",
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
        "VMSpecName": "e2-standard-4",
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
            "NameId": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-631x",
            "SystemId": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-631x"
          },
          {
            "NameId": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-szqw",
            "SystemId": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-szqw"
          }
        ],
        "KeyValueList": [
          {
            "key": "Config",
            "value": "{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:e2-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}}"
          },
          {
            "key": "Etag",
            "value": "d9c380dd-9dd4-4d32-8d99-abfb1c62ba68"
          },
          {
            "key": "InitialNodeCount",
            "value": "2"
          },
          {
            "key": "InstanceGroupUrls",
            "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-grp"
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
            "value": "{networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podIpv4CidrBlock:10.8.0.0/14,podIpv4RangeUtilization:0.002,podRange:gke-tb4376joq2qqrv47qau4-pods-331fc160,subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo}"
          },
          {
            "key": "PodIpv4CidrSize",
            "value": "24"
          },
          {
            "key": "SelfLink",
            "value": "https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tb4376joq2qqrv47qau4/nodePools/workers1"
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
            "value": "1.33.12-gke.1000000"
          },
          {
            "key": "InstanceGroup_0",
            "value": "gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-grp"
          },
          {
            "key": "keypair",
            "value": "tb4jmnhc0vlt4svljbps"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "34.47.85.231",
      "Kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://34.47.85.231\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVMVENDQXBXZ0F3SUJBZ0lSQUxJZTRTMTBNaTBFcWFRK3ZoZzlqUVV3RFFZSktvWklodmNOQVFFTEJRQXcKTHpFdE1Dc0dBMVVFQXhNa1l6RXdZVFprTldZdFpUTm1PQzAwTVRka0xXRmtPVGd0T1dRNE0yRTROVEV6T1RBegpNQ0FYRFRJMk1EZ3lNREExTVRZd01Gb1lEekl3TlRZd09ERXlNRFl4TmpBd1dqQXZNUzB3S3dZRFZRUURFeVJqCk1UQmhObVExWmkxbE0yWTRMVFF4TjJRdFlXUTVPQzA1WkRnellUZzFNVE01TURNd2dnR2lNQTBHQ1NxR1NJYjMKRFFFQkFRVUFBNElCandBd2dnR0tBb0lCZ1FDaCtzWlRYTFJ0bk5hY284V1h5dmhuUUw3L1NKUmd2T042QnRZTQo0SjdBVSsyNE1IWlhQdEdXdVdYaU96dDV2Njh0c1I3TUZMcnFOaEp4M05uY21IbFRBbmI2TDgzcDBFNXhKeTBOCmZkSDdSSkdEKzJ5dmFCMGRVQTZPTkZQTjQ1YTduY1ZvOEowSnI4aTkxdUVyYzE3N3pxdlo1WWh0c0VxS1dMM2kKY3hrVVZhdkJxS2pNQi9hOFpQUC90eDRaeDBwNWxudDBKVUcyMWZWYlZiKytYUjc4OE5zbkVONkZWdWw1R2NwWgpkSVlGVE1IUjJtbkx1VnlFb3d1eFk2YnBoSVJaendFRnJvdmxaK1Z5Z1F3czlURnlyMWc1VHVhVTd0S1dXbVpoCmx2bnY1amNuVlBFeGZ4MXR1RXdNM0xDZGZCMngyYUIrSDRnelBWTUZlakE1THRDcnhtM3llR2I1aStNZlpGYWYKNHViNWM2R1llUnUwWGZFekQ4dXFyT0JVVm9HTHNtT25CTW16d1hEeFE5MDVmU3g2bGxWTXRpbmoyMFkvYklrRwpmc0pSek44cDBUWTV0cUN5bGJsc3E3WUJiMXA5UHpZZUE0clBjNERvTmJMM3E5SXJqeXdpc0poRHo2ZGZ1dVlwCkZpSHhweWlXaGNKMVlzam1iTmtQcmRjVlhJc0NBd0VBQWFOQ01FQXdEZ1lEVlIwUEFRSC9CQVFEQWdJRU1BOEcKQTFVZEV3RUIvd1FGTUFNQkFmOHdIUVlEVlIwT0JCWUVGUHpKYi9pZis2TzFISnJLeHNaemFaeDhqQTljTUEwRwpDU3FHU0liM0RRRUJDd1VBQTRJQmdRQlNsajJDaXZKa0x5REZHZThHN0ZuY1BvMXdhS3NPaktmZE5pRGpXZjNNCk9Hc1d6QWgxYW5DTVB2d3o0NzNEUHRuYzJQU3pBWDJ6aGxGYWVwcUJSSlgzeDJLOXZPbHhTaUtkRzMyeXA2alcKSVRjd21QcWx4N0NrV0tNSzJGZFBaQ2IyUm9nS3Fwc2lLN2dLbng3UGtMNzdINnpWSGJ2SUlWVHM2QjA0RVNYcQprZEZTLzhOWDhxbFoxd0V4TjNkOVFEYVVrRjBUUk90REVNdDJ0NWowRGtzMzQwcXhlNlVJZDArL0pvSDRXMXk3CnZDM0FuV0lJc3BKTDZ3dDEwN09CbUdsaEJpSC8vUHlrYjNwbzdERzEzdk9TYXhFVGJObjVHZjMzQ2hoVEthVGQKM2p2OC92cU4wQXRDY3ZoRWY0SHI2YmlUY2R2ZXBFenY5Z2dWajVXVzA3eEVPVlRlbko1MGxNZnpNaGtGYUFYRApLaVduckNlRUg1anhNUDZIRVI4MGtOVmd0bk5HZDJyeHdqS21naUFKN25nUnEybGNhQmhwMzI3YVFadmt0T3V4CmFpQ28yT3VHbllLWkV2OUNyMVdpaERMNW1INm4rMHZTbGdwdkZqdVl0ZDkvR25uNnJoTy9CZ3Q1Z2dWcFFuTzkKQTVXZnU1dmZFaTBXcjlrb2xxYWxWdFE9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: gke_asia-northeast3-a_tb4376joq2qqrv47qau4\ncontexts:\n- context:\n    cluster: gke_asia-northeast3-a_tb4376joq2qqrv47qau4\n    user: gcp-dynamic-token\n  name: gke_asia-northeast3-a_tb4376joq2qqrv47qau4\ncurrent-context: gke_asia-northeast3-a_tb4376joq2qqrv47qau4\nusers:\n- name: gcp-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tb4376joq2qqrv47qau4/token?ConnectionName=gcp-asia-northeast3\\\"\"\n"
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-08-20T06:15:59Z",
    "KeyValueList": [
      {
        "key": "AddonsConfig",
        "value": "{gcePersistentDiskCsiDriverConfig:{enabled:true},kubernetesDashboard:{disabled:true},networkPolicyConfig:{disabled:true}}"
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
        "value": "10.8.0.0/14"
      },
      {
        "key": "ControlPlaneEndpointsConfig",
        "value": "{dnsEndpointConfig:{endpoint:gke-331fc160950b490f8aa57aa47d944bf5b40b-1064665102650.asia-northeast3-a.gke.goog},ipEndpointsConfig:{authorizedNetworksConfig:{},enablePublicEndpoint:true,enabled:true,privateEndpoint:10.0.1.8,publicEndpoint:34.47.85.231}}"
      },
      {
        "key": "CreateTime",
        "value": "2026-08-20T06:15:59+00:00"
      },
      {
        "key": "CurrentMasterVersion",
        "value": "1.33.12-gke.1000000"
      },
      {
        "key": "CurrentNodeCount",
        "value": "2"
      },
      {
        "key": "CurrentNodeVersion",
        "value": "1.33.12-gke.1000000"
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
        "value": "34.47.85.231"
      },
      {
        "key": "EnterpriseConfig",
        "value": "{clusterTier:STANDARD}"
      },
      {
        "key": "Etag",
        "value": "fb3b6858-b12a-4c44-b14d-d17d2f1e2744"
      },
      {
        "key": "Id",
        "value": "331fc160950b490f8aa57aa47d944bf5b40b2562089242069dca40971400e3ab"
      },
      {
        "key": "InitialClusterVersion",
        "value": "1.33.12-gke.1000000"
      },
      {
        "key": "InitialNodeCount",
        "value": "0"
      },
      {
        "key": "InstanceGroupUrls",
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-grp"
      },
      {
        "key": "IpAllocationPolicy",
        "value": "{clusterIpv4Cidr:10.8.0.0/14,clusterIpv4CidrBlock:10.8.0.0/14,clusterSecondaryRangeName:gke-tb4376joq2qqrv47qau4-pods-331fc160,defaultPodIpv4RangeUtilization:0.002,networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podCidrOverprovisionConfig:{},servicesIpv4Cidr:34.118.224.0/20,servicesIpv4CidrBlock:34.118.224.0/20,stackType:IPV4,useIpAliases:true}"
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
        "value": "{clientCertificateConfig:{},clusterCaCertificate:LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVMVENDQXBXZ0F3SUJBZ0lSQUxJZTRTMTBNaTBFcWFRK3ZoZzlqUVV3RFFZSktvWklodmNOQVFFTEJRQXcKTHpFdE1Dc0dBMVVFQXhNa1l6RXdZVFprTldZdFpUTm1PQzAwTVRka0xXRmtPVGd0T1dRNE0yRTROVEV6T1RBegpNQ0FYRFRJMk1EZ3lNREExTVRZd01Gb1lEekl3TlRZd09ERXlNRFl4TmpBd1dqQXZNUzB3S3dZRFZRUURFeVJqCk1UQmhObVExWmkxbE0yWTRMVFF4TjJRdFlXUTVPQzA1WkRnellUZzFNVE01TURNd2dnR2lNQTBHQ1NxR1NJYjMKRFFFQkFRVUFBNElCandBd2dnR0tBb0lCZ1FDaCtzWlRYTFJ0bk5hY284V1h5dmhuUUw3L1NKUmd2T042QnRZTQo0SjdBVSsyNE1IWlhQdEdXdVdYaU96dDV2Njh0c1I3TUZMcnFOaEp4M05uY21IbFRBbmI2TDgzcDBFNXhKeTBOCmZkSDdSSkdEKzJ5dmFCMGRVQTZPTkZQTjQ1YTduY1ZvOEowSnI4aTkxdUVyYzE3N3pxdlo1WWh0c0VxS1dMM2kKY3hrVVZhdkJxS2pNQi9hOFpQUC90eDRaeDBwNWxudDBKVUcyMWZWYlZiKytYUjc4OE5zbkVONkZWdWw1R2NwWgpkSVlGVE1IUjJtbkx1VnlFb3d1eFk2YnBoSVJaendFRnJvdmxaK1Z5Z1F3czlURnlyMWc1VHVhVTd0S1dXbVpoCmx2bnY1amNuVlBFeGZ4MXR1RXdNM0xDZGZCMngyYUIrSDRnelBWTUZlakE1THRDcnhtM3llR2I1aStNZlpGYWYKNHViNWM2R1llUnUwWGZFekQ4dXFyT0JVVm9HTHNtT25CTW16d1hEeFE5MDVmU3g2bGxWTXRpbmoyMFkvYklrRwpmc0pSek44cDBUWTV0cUN5bGJsc3E3WUJiMXA5UHpZZUE0clBjNERvTmJMM3E5SXJqeXdpc0poRHo2ZGZ1dVlwCkZpSHhweWlXaGNKMVlzam1iTmtQcmRjVlhJc0NBd0VBQWFOQ01FQXdEZ1lEVlIwUEFRSC9CQVFEQWdJRU1BOEcKQTFVZEV3RUIvd1FGTUFNQkFmOHdIUVlEVlIwT0JCWUVGUHpKYi9pZis2TzFISnJLeHNaemFaeDhqQTljTUEwRwpDU3FHU0liM0RRRUJDd1VBQTRJQmdRQlNsajJDaXZKa0x5REZHZThHN0ZuY1BvMXdhS3NPaktmZE5pRGpXZjNNCk9Hc1d6QWgxYW5DTVB2d3o0NzNEUHRuYzJQU3pBWDJ6aGxGYWVwcUJSSlgzeDJLOXZPbHhTaUtkRzMyeXA2alcKSVRjd21QcWx4N0NrV0tNSzJGZFBaQ2IyUm9nS3Fwc2lLN2dLbng3UGtMNzdINnpWSGJ2SUlWVHM2QjA0RVNYcQprZEZTLzhOWDhxbFoxd0V4TjNkOVFEYVVrRjBUUk90REVNdDJ0NWowRGtzMzQwcXhlNlVJZDArL0pvSDRXMXk3CnZDM0FuV0lJc3BKTDZ3dDEwN09CbUdsaEJpSC8vUHlrYjNwbzdERzEzdk9TYXhFVGJObjVHZjMzQ2hoVEthVGQKM2p2OC92cU4wQXRDY3ZoRWY0SHI2YmlUY2R2ZXBFenY5Z2dWajVXVzA3eEVPVlRlbko1MGxNZnpNaGtGYUFYRApLaVduckNlRUg1anhNUDZIRVI4MGtOVmd0bk5HZDJyeHdqS21naUFKN25nUnEybGNhQmhwMzI3YVFadmt0T3V4CmFpQ28yT3VHbllLWkV2OUNyMVdpaERMNW1INm4rMHZTbGdwdkZqdVl0ZDkvR25uNnJoTy9CZ3Q1Z2dWcFFuTzkKQTVXZnU1dmZFaTBXcjlrb2xxYWxWdFE9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K}"
      },
      {
        "key": "MasterAuthorizedNetworksConfig",
        "value": "{}"
      },
      {
        "key": "MonitoringConfig",
        "value": "{advancedDatapathObservabilityConfig:{},componentConfig:{enableComponents:[SYSTEM_COMPONENTS,STATEFULSET,JOBSET,STORAGE,HPA,POD,DAEMONSET,DEPLOYMENT,CADVISOR,KUBELET,DCGM]},managedPrometheusConfig:{enabled:true}}"
      },
      {
        "key": "MonitoringService",
        "value": "monitoring.googleapis.com/kubernetes"
      },
      {
        "key": "Name",
        "value": "tb4376joq2qqrv47qau4"
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
        "value": "{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:e2-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}}"
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
        "value": "{config:{bootDisk:{diskType:pd-balanced,sizeGb:100},diskSizeGb:100,diskType:pd-balanced,effectiveCgroupMode:EFFECTIVE_CGROUP_MODE_V2,imageType:COS_CONTAINERD,kubeletConfig:{maxParallelImagePulls:2},labels:{keypair:tb4jmnhc0vlt4svljbps},machineType:e2-standard-4,metadata:{disable-legacy-endpoints:true},oauthScopes:[https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/trace.append],resourceLabels:{goog-gke-node-pool-provisioning-model:on-demand},serviceAccount:default,shieldedInstanceConfig:{enableIntegrityMonitoring:true},tags:[tb5vuo31rrnng9587m1n],windowsNodeConfig:{}},etag:d9c380dd-9dd4-4d32-8d99-abfb1c62ba68,initialNodeCount:2,instanceGroupUrls:[https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/instanceGroupManagers/gke-tb4376joq2qqrv47qau4-workers1-21b6f7d2-grp],locations:[asia-northeast3-a],management:{autoRepair:true,autoUpgrade:true},maxPodsConstraint:{maxPodsPerNode:110},name:workers1,networkConfig:{networkTierConfig:{networkTier:NETWORK_TIER_DEFAULT},podIpv4CidrBlock:10.8.0.0/14,podIpv4RangeUtilization:0.002,podRange:gke-tb4376joq2qqrv47qau4-pods-331fc160,subnetwork:projects/GCP_PROJECT_ID/regions/asia-northeast3/subnetworks/tbf74se4bdd1ilmjkavo},podIpv4CidrSize:24,selfLink:https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tb4376joq2qqrv47qau4/nodePools/workers1,status:RUNNING,upgradeSettings:{maxSurge:1,strategy:SURGE},version:1.33.12-gke.1000000}"
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
        "value": "{privateEndpoint:10.0.1.8,publicEndpoint:34.47.85.231}"
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
        "value": "https://container.googleapis.com/v1/projects/GCP_PROJECT_ID/zones/asia-northeast3-a/clusters/tb4376joq2qqrv47qau4"
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

