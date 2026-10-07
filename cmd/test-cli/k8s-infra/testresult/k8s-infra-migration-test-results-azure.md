# CM-Beetle K8s Infra Migration Test Results — Azure-Busan

> [!NOTE]
> Full lifecycle against a real CSP: recommend → migrate → list → get (verified against
> the recommendation) → delete → residual resource check.

## Environment

- CSP / Region: azure / koreasouth
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (8928ba9)
- Git Commit: 8928ba9
- Namespace: mig01
- Test Date: 2026-09-16 20:37:48 KST
- Cluster ID: k8sazgcp01-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 23ms |
| 2 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 8m30.146s |
| 3 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 2ms |
| 4 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 11.433s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 38.949s |
| 6 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 6m17.954s |
| 7 | Residual resource check (Tumblebug) | ✅ **PASS** | 4ms |

**Overall Result**: 7/7 steps passed ✅

**Total Duration**: 15m38s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 23ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.34.8)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=azure+koreasouth+standard_b4as_v2 nodes=2

### Step 2 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 8m30.146s
- **Status Code**: 202

- ℹ️  nameSeed: k8sazgcp01
- ℹ️  async reqId: 1789558668508842438
- ℹ️  cluster id: k8sazgcp01-on-prem-k8s-cluster
- ℹ️  elapsed: 8m30s
- ✅ status: Active

### Step 3 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 2ms
- **Status Code**: 200

- ✅ migrated cluster present in list (1 total)

### Step 4 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 11.433s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=azure+koreasouth+standard_b4as_v2, nodes=2)
- ✅ version: 1.34.8 (recommended 1.34.8)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 38.949s

- ✅ kubeconfig obtained (server: https://dns-1789558793456303805-pyi9o5lf.hcp.koreasouth.azmk8s.io:443)
- ℹ️  auth method: static token in kubeconfig
- ✅ API server reachable (v1.34.8)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 1)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: 20.214.58.130
- ✅ nginx served over the LoadBalancer at http://20.214.58.130/ (attempt 1)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 6 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 6m17.954s
- **Status Code**: 200

- ✅ deleted on attempt 1 (6m17s)

### Step 7 — Residual resource check (Tumblebug)

- **Duration**: 4ms

- ℹ️  VNet k8sazgcp01-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup k8sazgcp01-k8s-sg still exists (known gap)
- ℹ️  SshKey k8sazgcp01-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for azure koreasouth (source: v1.32.3 → target: v1.34.8)",
  "targetCloud": {
    "csp": "azure",
    "region": "koreasouth"
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
    "connectionName": "azure-koreasouth",
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
    "connectionName": "azure-koreasouth",
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
      "connectionName": "azure-koreasouth",
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
    "connectionName": "azure-koreasouth",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "1.34.8",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": [
      {
        "name": "workers1",
        "imageId": "default",
        "specId": "azure+koreasouth+standard_b4as_v2",
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
  "id": "k8sazgcp01-on-prem-k8s-cluster",
  "uid": "tba3u1gnvc7mh37jtcck",
  "name": "k8sazgcp01-on-prem-k8s-cluster",
  "connectionName": "azure-koreasouth",
  "connectionConfig": {
    "configName": "azure-koreasouth",
    "providerName": "azure",
    "driverName": "azure-driver-v1.0.so",
    "credentialName": "azure",
    "credentialHolder": "admin",
    "regionZoneInfoName": "azure-koreasouth",
    "regionZoneInfo": {
      "assignedRegion": "koreasouth",
      "assignedZone": ""
    },
    "regionDetail": {
      "regionId": "koreasouth",
      "regionName": "koreasouth",
      "description": "Korea South",
      "location": {
        "display": "Korea South",
        "latitude": 35.1796,
        "longitude": 129.0756
      },
      "zones": []
    },
    "regionRepresentative": true,
    "verified": true
  },
  "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
  "systemMessage": "",
  "label": {
    "createdAt": "1789558784",
    "ownerCluster": "tba3u1gnvc7mh37jtcck",
    "sshkey": "tbk0pp1fi356dlisll7m",
    "sys.connectionName": "azure-koreasouth",
    "sys.createdTime": "2026-09-16 11:39:44 +0000 UTC",
    "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tba3u1gnvc7mh37jtcck",
    "sys.cspResourceName": "tba3u1gnvc7mh37jtcck",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "k8sazgcp01-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "k8sazgcp01-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tba3u1gnvc7mh37jtcck",
    "sys.version": "1.34.8"
  },
  "systemLabel": "",
  "version": "1.34.8",
  "network": {
    "vNetId": "k8sazgcp01-k8s-vpc",
    "subnetIds": [
      "k8sazgcp01-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "k8sazgcp01-k8s-sg"
    ],
    "keyValueList": null
  },
  "k8sNodeGroupList": [
    {
      "id": "workers1",
      "name": "workers1",
      "imageId": "default",
      "specId": "azure+koreasouth+standard_b4as_v2",
      "rootDiskType": "PremiumSSD",
      "rootDiskSize": 100,
      "sshKeyId": "k8sazgcp01-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 0,
      "maxNodeSize": 0,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "aks-workers1-74257999-vmss_0",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-74257999-vmss/virtualMachines/0"
        },
        {
          "cspResourceName": "aks-workers1-74257999-vmss_1",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-74257999-vmss/virtualMachines/1"
        }
      ],
      "keyValueList": null,
      "cspResourceName": "workers1",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tba3u1gnvc7mh37jtcck/agentPools/workers1",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tba3u1gnvc7mh37jtcck/agentPools/workers1"
        },
        "ImageIID": {
          "NameId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd/versions/202608.26.0",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd/versions/202608.26.0"
        },
        "VMSpecName": "Standard_B4as_v2",
        "RootDiskType": "PremiumSSD",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbk0pp1fi356dlisll7m"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "aks-workers1-74257999-vmss_0",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-74257999-vmss/virtualMachines/0"
          },
          {
            "NameId": "aks-workers1-74257999-vmss_1",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-74257999-vmss/virtualMachines/1"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://dns-1789558793456303805-pyi9o5lf.hcp.koreasouth.azmk8s.io:443",
    "kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUU2RENDQXRDZ0F3SUJBZ0lRSUZnZ0txbDBlcldOdmFuS2x5bFJOREFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWdGdzB5TmpBNU1UWXhNVE13TXpGYUdBOHlNRFUyTURreE5qRXhOREF6TVZvdwpEVEVMTUFrR0ExVUVBeE1DWTJFd2dnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUUNyClFIczRPZmdxczBCdkZiR3dVL1RXMjQ5OEZRRzV4OEs4NTJTTC9OaVpFY2I3NDVWOVN4V0NGalI1Y3dlNFNVNjAKeXBIbVdHbEdhY1RoR2NzWVBvQzdsR3pDRkV4dno4dGxRNnlsRkxMbzYrRElwdzVpd2czQ1BWY2hqRnBRdlVtRQpUemlCTGpIMlJ2SkRLTE1mVURzNkhuaG0xYkpuRExwYXZ4YlhFWE9IMTZFQzJ2VHpMTVkvMzZqNEhhVEpIemFSCkx3QlBXZDAwYUloeXZmWDFhd1lWV1RzRWIvR0RGU2dLSGl5UXVNVExGSkJYZU9sUVdmTStPeTBFbk1CbjJURkwKbWZvbW1wY3MzMTR2NnlMYlJzb2RzemxSbVAweFJVcFZvU3oycTdMRXU0bUlUaXNSby9QbmNYVllRMUN0eUlDNApJVmdVVy9jZG9ITjVLYUE3ZkQydVlObUtLYnNnb0ZYc0Rqc25QZ1g4cE4yNFpONjYxS2ROVEl4ejZ4YWRPZTlUCjBDNEhITEJMRGcyUXBoZExuSG5BY2xZNVVPWThQSmhLZzc3ZjZzNUNHSjRtSytZL2xTNElkMHlJcmhxb3hXNjUKTWdlOWg3aFJWUVRYWWNQSkI5ZEhTeGRhaUNmMXc3YStwSks2b0xnRnNEbENvaHBTRmdReTBSc25yUzN2ZWJjUwpGYmprOUNWYTlyZFdzYnkyMTRXK1ZVV1gyY1Jzcm1KRkpHOEIwSWZPSjlkaDZqeUdRMnJCUURYbDF4MGFqcVFNCjdpS3owVHdhVEUzeHh3Y2JJNTIvcWloeU5LdkFkSUlZelViU2VrWUwyRTNaa1hZVlNMcE1rQ0ZodkI1ckl3Z0UKdWtnOHc0RjNLeGVlYWp5czZTalhDUXBrNExKT2pFaUNTckV4d2VFa2R3SURBUUFCbzBJd1FEQU9CZ05WSFE4QgpBZjhFQkFNQ0FxUXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVBd2kvYXNQTy84VlI5UEhsCmh2T0tZUkhROHFZd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dJQkFBVTFWRkxJVExlM05aZG1WU0NsckQ3UUowTnkKQXR4VGVDeis5eHJXenczUzBGQi9XYWNGVGlFWkxDc2Y0ZVlQWTNpOUdOVUZiZGhCd2RrQmYwcjF2Ujgva3BDSgpGOHhudTMzcTB2US81NDVoaFdMdjlmaUQ1d1BrUW9tY3hSVCs2cnhDOENselJaVWZMc1RwcVZ6VWtNMkJSZXY3CkNEazNPcGJQdktVYWtEdVVlaXFwWXJoTmhuWmFUTm5YbXdNWEtJSUlsaktwK292dUJmOTZjY3dmQWNBMjVIV3IKNFc0amNYQVlGbGpOcFdGckE0MjZQbjgzNGR5N2hWcHhZOHlWU201ckxrQjNCeWx2ZlJ5dFkya1lGWW9YRElhZgpaSUxjL3VSSnF4OEdMNE55L05iUExzRWs3a1FTbHZ6ZHlhNUR0QU02M1ptV0lDcnNnTk5oOGtSOHRicytZaTJxCi95dWR2Ni9YUzUrRzArbkQ4QXA0Y05jbjdmRmJpbXE5MVEyeHc5MkpQS3hBVy9aNkh6ODErWUlkWGZOWkt4R0gKK3htcWtXajV3TEh3dC9CdjhIQXZxMkU0U3dvYUE1RXNSSzIwWXVQOTh4WXAvZ21Fcks3V1ZXd1ZGTmFGS0hPdwoxWlNFNmNXd1loejUvNUhmcVVlWkUvMVlJQ1BMRStMRFZXU0NKYVh4bFlkL044TXFIMDM1KzBSTFU3WTM5WU9QCk1lWWhoMWFkaFUvN29Yb2h4dnhjdzYyL0t4UVJ3REk5VnR2U2dOb3Jhdk1KdzV1R290NVBTTmhEQk5pZDBMcWYKa3VjOUU3ODAzdXhEWUhMOTNsdEJPZHRQZDZGRlBockdEbmw0MWJJZjlZYVQ5eXozc0E4WWlQSmMwLzF6SU4xLwo1UmlGdDkvK2NwR0IzWUJOCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://dns-1789558793456303805-pyi9o5lf.hcp.koreasouth.azmk8s.io:443\n  name: tba3u1gnvc7mh37jtcck\ncontexts:\n- context:\n    cluster: tba3u1gnvc7mh37jtcck\n    user: clusterAdmin_koreasouth_tba3u1gnvc7mh37jtcck\n  name: tba3u1gnvc7mh37jtcck\ncurrent-context: tba3u1gnvc7mh37jtcck\nkind: Config\nusers:\n- name: clusterAdmin_koreasouth_tba3u1gnvc7mh37jtcck\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZIVENDQXdXZ0F3SUJBZ0lRT2hsa1pLOTJLWnoxTzJnNjA0Z0hRakFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWVGdzB5TmpBNU1UWXhNVE13TXpGYUZ3MHlPREE1TVRZeE1UUXdNekZhTURBeApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1SVXdFd1lEVlFRREV3eHRZWE4wWlhKamJHbGxiblF3CmdnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUURRQ0V6SXM1N1hkd0hGQTE4M1hCdHUKSnZ0emF4ZUZrQ05qeXhqbGdXN2RHRytjR2N5UHNnMnhKQmpHVkdJcXY5MWJnWVNQeFRqRFBjTENqVldsd1RXdQpCaTNUclRHS2oxS3dQc2lNb3daVXdjSmU3aHpmbkNZZ0hEMStQOVFlWUxYWWlJbjFpeXR5TitxbDVPS1FLSXVMCjFLdHpDL2EwajIxRTF2bEtTcTVtOVJOYmVQd0lpS2JEZWpERjAraWk3SzdtV1lKbDZqc3N3ZS8yc3BnZ3Y0VDMKNVdmL1JFL0ZhOUxuRk1GRUZ0VitodnFVMHhwRDZ1d05GUE5kY0Q1THVUeDJJYVRLUHFwb2JYUGhiYUhHODkwSwptRnArMWNIRkRIMGZRK0ZodzFya0JMcmdDaGprVTdJSWJTYXR0Uk9IQkhSbm0xZG13TFp3TGJiQk9mMXNhTkQ5CkVxUWM1aU1wVVlMMWxiWEJCYjhZOHpoajN0VXhlZ2NlWHhnMDRBaFhVKzdWdTArdXE0aWJMT3VHQ2UzUjF5OGUKMzhoVkJTWXBUTDBsUTlQc1UvWmh4VzM0blpMdUtzWGxxbi9ITnQ3YzJwdmV1RUd5c0k3enlFWDFqRWIwck9NSgorR2p3UDFjT2tqWE5uQnF4MVFEVjBQRVJWYXdickh1YVNlQko4NWlENTNwREFZeVVkRmsrc1RKa1FCVjVrYm4yClk5RDkwZ3NJZ29LZXQxNSthUjlPMnRCR1c3ZzdaNlV5a0xEZ0s0TjZpRFR4bW5zUm45VFZNMkJkL1J5bzVXeFcKamZmWjd4d2s1eERaWUN1N1ZueFV1VmxVY2RnN3ZpTmJKTWNhejJDZ0QwK3JYdWpRTnVtaWFHTS83ZUd0a1lScAo5RUxpL0VNUFdiYnRTWFRUYmlsY2l3SURBUUFCbzFZd1ZEQU9CZ05WSFE4QkFmOEVCQU1DQmFBd0V3WURWUjBsCkJBd3dDZ1lJS3dZQkJRVUhBd0l3REFZRFZSMFRBUUgvQkFJd0FEQWZCZ05WSFNNRUdEQVdnQlFEQ0w5cXc4Ny8KeFZIMDhlV0c4NHBoRWREeXBqQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FnRUFXUGNqRUdSd25VSE5WVy9JaVlTeQpzZkxvL2JrTkN0dTI4NEdPTDV6RWZnbFBtVlZVelZIMkRLQ0xjd2xkM0tUTFhpaENhOTIvenNmcGJPZlJxZXRJCnlCRzVtcEs2V3U2N0ZBdmQ5TVN6TXhTOFU2cnZkQXBrUlVXcU0vVVowallWbzNjN21FeTlyUzVXelNwTWllVXMKblJoRTcybHMxbjMwbG1wcHJMRUhtMGtSWW05OHRXNm9IKzhNOHdISjJJUExIcHJBQkQyUWZFczhXUVJjSDR0RgpBaVgvdzQ3cGU0WkpZY29UMUZLVG9zVTM2SzROV3pOU3QyKzRFN0h3aGxLdTQxTFFOejF5NDdPNHJ4U05EbDJDCmlkdjBqTjZnMEQ5U3ZvMFlBTm1CMkprYXRXbGk3bXNXbFpCUTlwaFBkRWNxNnk5ZjRpTnNYd2N4OFlOdTMreDEKaWhlb3l5ZHZpV3FQZkx5d1JndUFtQ1E0TExMTEZ0SnMreWUvaEhWUmhhdm5BNXA1TlFWWm0yUGR0emVSV2JZMgpOOFpLc01obEt3UWE1MFJYdXRtazlOZGdBZGM4ZlVFanAyV01nUGdtNnA4RGM1UEFzcVh3Qk0yZlNPQnU5YWhzClpOVVdnajdZVTJwNmkyekxZenArTVh5VjQ2VTZOeGN4RXY0aEloOFJqN3N1bmVkZG9ucHNYeEtwUWxhYzFhUHUKOEtmMFpaTEZYZEQ1dEJvRlVPR2o1bU54QlpLMU1SZEJyK0hVMy9BZG1LMkR1VkNvc1I0M2ZkVlZTK0VmQWpzZQpnM2tCWldMN2w2WjJWaDlyWGpLRERjUlBPbDVGZk9nV3F4Tlhoa3NyaE5jZnowL0YyQ1ZMT254SlRDMjlUenQyCm5zOURzak1TY1E3REdrNjN4WCttZkRzPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS1FJQkFBS0NBZ0VBMEFoTXlMT2UxM2NCeFFOZk4xd2JiaWI3YzJzWGhaQWpZOHNZNVlGdTNSaHZuQm5NCmo3SU5zU1FZeGxSaUtyL2RXNEdFajhVNHd6M0N3bzFWcGNFMXJnWXQwNjB4aW85U3NEN0lqS01HVk1IQ1h1NGMKMzV3bUlCdzlmai9VSG1DMTJJaUo5WXNyY2pmcXBlVGlrQ2lMaTlTcmN3djJ0STl0Uk5iNVNrcXVadlVUVzNqOApDSWltdzNvd3hkUG9vdXl1NWxtQ1plbzdMTUh2OXJLWUlMK0U5K1ZuLzBSUHhXdlM1eFRCUkJiVmZvYjZsTk1hClErcnNEUlR6WFhBK1M3azhkaUdreWo2cWFHMXo0VzJoeHZQZENwaGFmdFhCeFF4OUgwUGhZY05hNUFTNjRBb1kKNUZPeUNHMG1yYlVUaHdSMFo1dFhac0MyY0MyMndUbjliR2pRL1JLa0hPWWpLVkdDOVpXMXdRVy9HUE00WTk3VgpNWG9ISGw4WU5PQUlWMVB1MWJ0UHJxdUlteXpyaGdudDBkY3ZIdC9JVlFVbUtVeTlKVVBUN0ZQMlljVnQrSjJTCjdpckY1YXAveHpiZTNOcWIzcmhCc3JDTzg4aEY5WXhHOUt6akNmaG84RDlYRHBJMXpad2FzZFVBMWREeEVWV3MKRzZ4N21rbmdTZk9ZZytkNlF3R01sSFJaUHJFeVpFQVZlWkc1OW1QUS9kSUxDSUtDbnJkZWZta2ZUdHJRUmx1NApPMmVsTXBDdzRDdURlb2cwOFpwN0VaL1UxVE5nWGYwY3FPVnNWbzMzMmU4Y0pPY1EyV0FydTFaOFZMbFpWSEhZCk83NGpXeVRIR3M5Z29BOVBxMTdvMERicG9taGpQKzNoclpHRWFmUkM0dnhERDFtMjdVbDAwMjRwWElzQ0F3RUEKQVFLQ0FnQUJycStBU0FPVzFuMkxMRlhPeXMzbC9DYTRianRJZHp2eUNLaHc0clVVMEtmR2FXY0FHbjZGMmpiaApFN21mZ3VHMVpieSt0T2Vhbkp0QW00Zi95U28zK0JEU3oybkJKeHVTRlUvbDQwT2YxOWxNanp4b2lvaThaYjRSCmtVNlQyRkJnS2VxRHM5WnNIQXVudjh3ZFFsYXVrTSs2SkhTZ1RUQ2pCK0lJT2NaalNzUVhUMGtxZ0lCb3dFbzQKcjFnSVNVVzQweXg4eW9Ja0FJV0NSenQzdUZUSTlHLzgzbjZPcUNxU2Q5YUFTSGI5aVBYcDBqTTZMV0l2VU9ZTwo4V052bFZYcDFxZlFndVU1NXZpeThBSUk2OXZ6dy91blh1OHNPc2VnUXhwRmdpRHdDeUcvd2hVbGM0L1RCWjcwCmRLeHR3UmwzNVFlMTZ3VU9yeS8xTEJUaWxZNXJHY1hXcG5scVVjNjQyM1laaFljVXNocnF1UFI4Y1lmOHNZL0gKU0ZHNzJhYUhsV2trcWt5OFJJV0JaMTFEMzZPRDNFZlpBV3JaVDNuSGw5b2toWmwxa2pWZEordmlQS1lGbFVObwoyaC9vdGp1ekg3RFNORG1kdXBaV2pYWWp0Y1Q0cXd3dnUrZlR1bkJIUjRPWm5wTlNKdmU0TlEwMFBkYnFzVmJqCjVhQzBtdW12QVpRV1R5NFNUQ0Ztem1uYTBaZEtTRDBadUFiWUNJTzZ5aWsxYXRuWWRReWJmV3RuUklqUjZnK0UKOXBvaThKYlFxRzhDSTZjOUllcndwZWIrRXpYMEoyclYrcVZ0Z2tVQnpvZVh3RFVIWXgyb1FIYW5TRGFSenN6dgp6MVZVMUt4eVZrMjJ4MC9HeUpTT3g3VHRJYTRqMzRtdDIrS05kRzFmdy9MeklETHhBUUtDQVFFQTBsZmdWYytzClJFcHJ1Z2xzOVh1T210OUFrRngyVXFhTlYvREZ1T3lPRndRQWppWkMzRGJzUzBweEFBaGFUcUhUekNTZFdQSTIKYWFaWlAvUlNINlJFeElJMTZBVTh5L01sbWZocXI2cVpSRmMzMVJBWG8wZ204ZG55ckxQTTVBRDdxVWhVZE8yTwoyY3Axcnhyck54bTAzMjBCMHgvNlZEWHZ2eGxjYjE4dXdxcnRubFlHa0ZiaTkwVllDeVJpVWxwRmlVc3dMSVk4CjZuMWVjcFFmN1lRWEJrVjNhUjFlVkdBam1ucHhBZ2c2M0ZyN0xrT0hEQmhTdTVKNER6Q2tKUmQxajZqcitGY3kKUDRLc0Vmb2V2akJtZjducVdCQ1NUdUoxQzdESEpNZnQrek9XLzlSQkM5ckVNY0RVb1A0WmxjK1VYcGJrR3o4YwpHK0tjbW5tQzV1bEt5d0tDQVFFQS9UQUVZZXkwZmI1L0pLRTZLUmIrcWE2M0xCT1diNTlBR1VmZktudFhpdnVnCkVycEtSSC9IWkNkbHhBUy9qTFliMXV1TWpsTjUrYjlXVmgwL2hYb2ZHSjdJMmNLL21wZTF5SkFBQTNTVm9hbHUKTDRtUTR0dVhtM1lXU25pQnE1S2hJejR0bmkrMnAvNVJRNFNLb3o2dnh4SjNlRFNQNzJGc2w0bDRZRENHbUpoTworMUdvVStLN1JEYjY2RTR6aDhZbThCTGRobFZ1TXJONVVWbUwyNW1ldzVyNGVXc3lIZHFUMllGS1pOZzFxUkpiClFPZ01yREFOL0IzT1ZtMTZnbjNDWWt4U3JwcGVCK2p5bWRXejRGOXlnSnJqdzEyRlJja1ptaTAzaFZYMXN1T1AKcVhER3MvZWpITjZmVnJEd2JnejFucFVhSTNrbzNjcVB6WXlxeC9FOVFRS0NBUUJsSDNBdmNQYTlvaFdtUzYvQgpXTlJYS210c3U3YjZ4eTEra0xkTnp4UUVocDBKdWVVODMzNjhONTZaeUdvNkVWeTBjUW5nWUJtK3N3V0hWSnRDCnNRT0tnWnNPMzYyNTB6eEppSDhwMHRkNlhuL3VBNTJKbHo5NTJERHR6RWI5ZW1lQ204NVdwSHFmdzlET0RSLzEKem5zQlN2T3NuMXdHcmlPRGVOclNoQzBEMDQ4SEp3NWl5Qkszay9QZTczQUJiWHF1cGFJVVZiamxkcjQ0VVhRRQpDUlBVcFFMaVd3Y2xnMDY2anBEVHpFY1g5dmw0NVdnQWJaVWdyaXJnQ3A5dllaYllLUHJBMnBMQ3E3eXpxODZwCjJyVDhSTklmNGwra1l3U29TU1dFVGtYZ0ZZNUxrYXh2MlkybTZiQ1BjWG0xWnlWS1VEcDd5dHhsbjQxd1NtaWYKUUpyL0FvSUJBUUNUQkZuWEh1S2pqWVU0bU9JenorVnFWRWlLc0lrUElkTFBtenRMNkxrcHUvajBSdll2RmwrSwptWkh2STY2eW4yQkZDUnZoM1RrYnUyMy9yUllnaEl3ZitMdTZMdXFoY3V5Y25IbFZpRklHd0dKNStoQ0dtbXBMClhHT1BOTEJmRjFLNEt6ZkQvZ0s3UStLZUtRMXp3MGZBZWNtanBDbmNINTgxMHg1eUJGdHpxaVZhcTh6cGdPT3EKdFo0MlhJcHhrYTgwZ2svZDNDZVVDMEVyNnFwYWhyWjQ4TGpOa0dCV2s0QjhzU1dvcng5aG9JWWFoMlFzYSs4Qwo4YS9KNGVKY2VYZnhLazVza3JoOU1WZ0YyZUNNTGdCSUN5aXNYZGF3Y2hpS3ZTemJJTkg5NFVPZmFSd0lqb2hKCmtEL0gwQkNjWnF3TlBKa2o5Z3V6MWhad2pmanhCdlRCQW9JQkFRRE1OcmJzRTlDREM0NDRqT093cHdab1l3Ri8KS2xpM0J5NmRWMDBQZmNQajdVd0p4dGcyNnZIRWg5TGpHSExuVldqMzl6bkFqYUpuejZLNG5VU2dHSFJzT05GRwordG1SZlRERWxOalNqbzk4WnJ6RzNUZUdlYU1iNERCeE51S2F1TzdYZmptS0N3cnZrSC9hdHY5N2NrV0ZqRkVwCllWc2xUMm14WkNIRCtkdGVFT1FKdmZ3MXU4Q2twMVg1dEJkWGlDbjJ2Mml0OHRSaXdsWDZ5RGVvSDVubG9IT3YKeFU0UEdJczN6MlYrM01pMjJxaVRIKytLV0NacWtERG10WmFjOE53VWVaTVlkcG5wWDc3Z0dTMjBrS0hZY1QyMApNZ2tpU21uSVNsdTAyeXpSSmxpRHlrdkdycHJFcG5SVittZWVwUjlXVVcyd2oza1k2a0NkYy9yVE40cXkKLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K\n    token: 642bmwx1u4bni65xv2uqiaj6495zi0hv49uhw317fvlq69ucgoo9zjhy0deymay8433ymmbaftwzr5p3o3ozxcbs4p2ows703ufl0f9opw8a19m80xw0kmnpx9ptqpgq\n"
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-09-16T11:39:44Z",
  "keyValueList": [
    {
      "key": "Location",
      "value": "koreasouth"
    },
    {
      "key": "Identity",
      "value": "{principalId:2b189827-f0a5-4940-9067-25503b9848ee,tenantId:fb98dda1-32ff-48eb-a489-62777cd9ccd8,type:SystemAssigned}"
    },
    {
      "key": "Kind",
      "value": "Base"
    },
    {
      "key": "Properties",
      "value": "{agentPoolProfiles:[{count:2,currentOrchestratorVersion:1.34.8,eTag:db29d286-5d41-4be0-b26c-472200a89b23,enableAutoScaling:false,enableFIPS:false,enableNodePublicIP:true,kubeletDiskType:OS,maxPods:110,mode:System,name:workers1,nodeImageVersion:AKSUbuntu-2204gen2containerd-202608.26.0,orchestratorVersion:1.34.8,osDiskSizeGB:100,osDiskType:Managed,osSKU:Ubuntu,osType:Linux,powerState:{code:Running},provisioningState:Succeeded,scaleDownMode:Delete,securityProfile:{enableSecureBoot:false,enableVTPM:false,sshAccess:LocalUser},type:VirtualMachineScaleSets,upgradeSettings:{maxSurge:10%,maxUnavailable:0},vmSize:Standard_B4as_v2,vnetSubnetID:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb3re4cev3tpmj6onblt/subnets/tblfcbvuf4dqbh9smm5l}],autoUpgradeProfile:{nodeOSUpgradeChannel:NodeImage},azurePortalFQDN:dns-1789558793456303805-pyi9o5lf.portal.hcp.koreasouth.azmk8s.io,bootstrapProfile:{artifactSource:Direct},currentKubernetesVersion:1.34.8,dnsPrefix:dns-1789558793456303805,enableRBAC:true,fqdn:dns-1789558793456303805-pyi9o5lf.hcp.koreasouth.azmk8s.io,identityProfile:{kubeletidentity:{clientId:b75814df-61e9-4f6d-8476-a2637710c8b8,objectId:6c70b044-5ba8-4b55-9d9d-041ec01df227,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/tba3u1gnvc7mh37jtcck-agentpool}},ingressProfile:{webAppRouting:{dnsZoneResourceIds:[/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/dnszones/tba3u1gnvc7mh37jtcck.com],enabled:true,gatewayAPIImplementations:{appRoutingIstio:{mode:Disabled}},identity:{clientId:a7aaba24-aec8-438c-8203-d49b45af4874,objectId:ef2aa35c-6376-4c31-a85e-3d08276a4001,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/webapprouting-tba3u1gnvc7mh37jtcck},nginx:{defaultIngressControllerType:AnnotationControlled}}},kubernetesVersion:1.34.8,linuxProfile:{adminUsername:cb-user,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDLxAxSdnehI5glyZeeDMtrLnqQnMbcC9j20KbF2Cfqinio09bWfVS0DF/hsi2TkcsTCWi4Z17/eHO3jYfVrIdg5aAn+1PAjrEPN9LChSeTqe2nI2bZrvW4L0ebzYWpjzXNF2c4KU0x181o+A8YQengU5C65hB6nBBKFg+mnefDM/BjHSrSDgRbFlXcUzmScJ62VCZZwG0YQzCrg6vdFDeWT5M34kN7pz5NAwJjri5TtRj4fWYine3IHOlJ1Q0Lqqd3N0NqcDvVm29zP4JsOGB1obvpnuU0IXdGs6e2K1Jbs+zHMUw88c10m8SSjuJPnXOAYwkbhl6IiQDtCuok4uLOpi5nyXGOXRO/oxhV1X4/7O5ttrQGnEksvUAdDTzQfMXpfKjPeX1toU9Iwq80XenBJ2vB5FVFV7cvVj40LATeKh+dMEjcRe2lePmBohdy3oQJhgJq/kPYMHtGW7hzVPe/xbwH3NEjgGiqDoI0meRt9oa5FVNqyI+KHlS+Y2kma/yga+HiBAGGcl7/oDIVAqiPpw0gs2r8YjbmsNDlqMJcQw8f5G3jRzpTcWcXP3nOzlS3wjzoX75tfUj5d1ghBFw2FKmqtEOLlAzReOfRfr9qVuv//4v1+MphjOfnzSfV9cQOmxLS+jfdt33ezdUlby7LAgeW9enlmcUDHkvVOW++zQ==\\n}]}},maxAgentPools:100,metricsProfile:{costAnalysis:{enabled:false}},networkProfile:{dnsServiceIP:10.1.0.10,ipFamilies:[IPv4],loadBalancerProfile:{backendPoolType:nodeIPConfiguration,effectiveOutboundIPs:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.Network/publicIPAddresses/4d404521-548b-4bdd-ab61-e14ea4a43c7a}],managedOutboundIPs:{count:1}},loadBalancerSku:standard,networkDataplane:azure,networkPlugin:azure,networkPolicy:azure,outboundType:loadBalancer,serviceCidr:10.1.0.0/16,serviceCidrs:[10.1.0.0/16]},nodeProvisioningProfile:{mode:Manual},nodeResourceGroup:CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth,oidcIssuerProfile:{enabled:true,issuerURL:https://koreasouth.oic.prod-aks.azure.com/fb98dda1-32ff-48eb-a489-62777cd9ccd8/c0271972-08cc-4988-b847-0d0a9fb97b8e/},powerState:{code:Running},provisioningState:Succeeded,resourceUID:6aaa800ec7aec500016591ee,securityProfile:{},servicePrincipalProfile:{clientId:msi},storageProfile:{diskCSIDriver:{enabled:true},fileCSIDriver:{enabled:true},snapshotController:{enabled:true}},supportPlan:KubernetesOfficial,windowsProfile:{adminUsername:azureuser,enableCSIProxy:true},workloadAutoScalerProfile:{}}"
    },
    {
      "key": "SKU",
      "value": "{name:Base,tier:Standard}"
    },
    {
      "key": "Tags",
      "value": "{createdAt:1789558784,ownerCluster:tba3u1gnvc7mh37jtcck,sshkey:tbk0pp1fi356dlisll7m,sys.connectionName:azure-koreasouth,sys.createdTime:2026-09-16 11:39:44 +0000 UTC,sys.cspResourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tba3u1gnvc7mh37jtcck,sys.cspResourceName:tba3u1gnvc7mh37jtcck,sys.description:Migrated from on-premise K8s cluster (v1.32.3, 2 workers),sys.id:k8sazgcp01-on-prem-k8s-cluster,sys.labelType:k8s,sys.manager:cb-tumblebug,sys.name:k8sazgcp01-on-prem-k8s-cluster,sys.namespace:mig01,sys.uid:tba3u1gnvc7mh37jtcck,sys.version:1.34.8}"
    },
    {
      "key": "ETag",
      "value": "36b817bf-0b40-4150-b353-9675a97f66ee"
    },
    {
      "key": "ID",
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tba3u1gnvc7mh37jtcck"
    },
    {
      "key": "Name",
      "value": "tba3u1gnvc7mh37jtcck"
    },
    {
      "key": "Type",
      "value": "Microsoft.ContainerService/ManagedClusters"
    }
  ],
  "cspResourceName": "tba3u1gnvc7mh37jtcck",
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tba3u1gnvc7mh37jtcck",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tba3u1gnvc7mh37jtcck",
      "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tba3u1gnvc7mh37jtcck"
    },
    "Version": "1.34.8",
    "Network": {
      "VpcIID": {
        "NameId": "tb3re4cev3tpmj6onblt",
        "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb3re4cev3tpmj6onblt"
      },
      "SubnetIIDs": [
        {
          "NameId": "tblfcbvuf4dqbh9smm5l",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb3re4cev3tpmj6onblt/subnets/tblfcbvuf4dqbh9smm5l"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "#aks-agentpool-16996668-nsg",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/cb_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.Network/networkSecurityGroups/aks-agentpool-16996668-nsg"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tba3u1gnvc7mh37jtcck/agentPools/workers1"
        },
        "ImageIID": {
          "NameId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd/versions/202608.26.0",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd/versions/202608.26.0"
        },
        "VMSpecName": "Standard_B4as_v2",
        "RootDiskType": "PremiumSSD",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbk0pp1fi356dlisll7m"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "aks-workers1-74257999-vmss_0",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-74257999-vmss/virtualMachines/0"
          },
          {
            "NameId": "aks-workers1-74257999-vmss_1",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-74257999-vmss/virtualMachines/1"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://dns-1789558793456303805-pyi9o5lf.hcp.koreasouth.azmk8s.io:443",
      "Kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUU2RENDQXRDZ0F3SUJBZ0lRSUZnZ0txbDBlcldOdmFuS2x5bFJOREFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWdGdzB5TmpBNU1UWXhNVE13TXpGYUdBOHlNRFUyTURreE5qRXhOREF6TVZvdwpEVEVMTUFrR0ExVUVBeE1DWTJFd2dnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUUNyClFIczRPZmdxczBCdkZiR3dVL1RXMjQ5OEZRRzV4OEs4NTJTTC9OaVpFY2I3NDVWOVN4V0NGalI1Y3dlNFNVNjAKeXBIbVdHbEdhY1RoR2NzWVBvQzdsR3pDRkV4dno4dGxRNnlsRkxMbzYrRElwdzVpd2czQ1BWY2hqRnBRdlVtRQpUemlCTGpIMlJ2SkRLTE1mVURzNkhuaG0xYkpuRExwYXZ4YlhFWE9IMTZFQzJ2VHpMTVkvMzZqNEhhVEpIemFSCkx3QlBXZDAwYUloeXZmWDFhd1lWV1RzRWIvR0RGU2dLSGl5UXVNVExGSkJYZU9sUVdmTStPeTBFbk1CbjJURkwKbWZvbW1wY3MzMTR2NnlMYlJzb2RzemxSbVAweFJVcFZvU3oycTdMRXU0bUlUaXNSby9QbmNYVllRMUN0eUlDNApJVmdVVy9jZG9ITjVLYUE3ZkQydVlObUtLYnNnb0ZYc0Rqc25QZ1g4cE4yNFpONjYxS2ROVEl4ejZ4YWRPZTlUCjBDNEhITEJMRGcyUXBoZExuSG5BY2xZNVVPWThQSmhLZzc3ZjZzNUNHSjRtSytZL2xTNElkMHlJcmhxb3hXNjUKTWdlOWg3aFJWUVRYWWNQSkI5ZEhTeGRhaUNmMXc3YStwSks2b0xnRnNEbENvaHBTRmdReTBSc25yUzN2ZWJjUwpGYmprOUNWYTlyZFdzYnkyMTRXK1ZVV1gyY1Jzcm1KRkpHOEIwSWZPSjlkaDZqeUdRMnJCUURYbDF4MGFqcVFNCjdpS3owVHdhVEUzeHh3Y2JJNTIvcWloeU5LdkFkSUlZelViU2VrWUwyRTNaa1hZVlNMcE1rQ0ZodkI1ckl3Z0UKdWtnOHc0RjNLeGVlYWp5czZTalhDUXBrNExKT2pFaUNTckV4d2VFa2R3SURBUUFCbzBJd1FEQU9CZ05WSFE4QgpBZjhFQkFNQ0FxUXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVBd2kvYXNQTy84VlI5UEhsCmh2T0tZUkhROHFZd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dJQkFBVTFWRkxJVExlM05aZG1WU0NsckQ3UUowTnkKQXR4VGVDeis5eHJXenczUzBGQi9XYWNGVGlFWkxDc2Y0ZVlQWTNpOUdOVUZiZGhCd2RrQmYwcjF2Ujgva3BDSgpGOHhudTMzcTB2US81NDVoaFdMdjlmaUQ1d1BrUW9tY3hSVCs2cnhDOENselJaVWZMc1RwcVZ6VWtNMkJSZXY3CkNEazNPcGJQdktVYWtEdVVlaXFwWXJoTmhuWmFUTm5YbXdNWEtJSUlsaktwK292dUJmOTZjY3dmQWNBMjVIV3IKNFc0amNYQVlGbGpOcFdGckE0MjZQbjgzNGR5N2hWcHhZOHlWU201ckxrQjNCeWx2ZlJ5dFkya1lGWW9YRElhZgpaSUxjL3VSSnF4OEdMNE55L05iUExzRWs3a1FTbHZ6ZHlhNUR0QU02M1ptV0lDcnNnTk5oOGtSOHRicytZaTJxCi95dWR2Ni9YUzUrRzArbkQ4QXA0Y05jbjdmRmJpbXE5MVEyeHc5MkpQS3hBVy9aNkh6ODErWUlkWGZOWkt4R0gKK3htcWtXajV3TEh3dC9CdjhIQXZxMkU0U3dvYUE1RXNSSzIwWXVQOTh4WXAvZ21Fcks3V1ZXd1ZGTmFGS0hPdwoxWlNFNmNXd1loejUvNUhmcVVlWkUvMVlJQ1BMRStMRFZXU0NKYVh4bFlkL044TXFIMDM1KzBSTFU3WTM5WU9QCk1lWWhoMWFkaFUvN29Yb2h4dnhjdzYyL0t4UVJ3REk5VnR2U2dOb3Jhdk1KdzV1R290NVBTTmhEQk5pZDBMcWYKa3VjOUU3ODAzdXhEWUhMOTNsdEJPZHRQZDZGRlBockdEbmw0MWJJZjlZYVQ5eXozc0E4WWlQSmMwLzF6SU4xLwo1UmlGdDkvK2NwR0IzWUJOCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://dns-1789558793456303805-pyi9o5lf.hcp.koreasouth.azmk8s.io:443\n  name: tba3u1gnvc7mh37jtcck\ncontexts:\n- context:\n    cluster: tba3u1gnvc7mh37jtcck\n    user: clusterAdmin_koreasouth_tba3u1gnvc7mh37jtcck\n  name: tba3u1gnvc7mh37jtcck\ncurrent-context: tba3u1gnvc7mh37jtcck\nkind: Config\nusers:\n- name: clusterAdmin_koreasouth_tba3u1gnvc7mh37jtcck\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZIVENDQXdXZ0F3SUJBZ0lRT2hsa1pLOTJLWnoxTzJnNjA0Z0hRakFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWVGdzB5TmpBNU1UWXhNVE13TXpGYUZ3MHlPREE1TVRZeE1UUXdNekZhTURBeApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1SVXdFd1lEVlFRREV3eHRZWE4wWlhKamJHbGxiblF3CmdnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUURRQ0V6SXM1N1hkd0hGQTE4M1hCdHUKSnZ0emF4ZUZrQ05qeXhqbGdXN2RHRytjR2N5UHNnMnhKQmpHVkdJcXY5MWJnWVNQeFRqRFBjTENqVldsd1RXdQpCaTNUclRHS2oxS3dQc2lNb3daVXdjSmU3aHpmbkNZZ0hEMStQOVFlWUxYWWlJbjFpeXR5TitxbDVPS1FLSXVMCjFLdHpDL2EwajIxRTF2bEtTcTVtOVJOYmVQd0lpS2JEZWpERjAraWk3SzdtV1lKbDZqc3N3ZS8yc3BnZ3Y0VDMKNVdmL1JFL0ZhOUxuRk1GRUZ0VitodnFVMHhwRDZ1d05GUE5kY0Q1THVUeDJJYVRLUHFwb2JYUGhiYUhHODkwSwptRnArMWNIRkRIMGZRK0ZodzFya0JMcmdDaGprVTdJSWJTYXR0Uk9IQkhSbm0xZG13TFp3TGJiQk9mMXNhTkQ5CkVxUWM1aU1wVVlMMWxiWEJCYjhZOHpoajN0VXhlZ2NlWHhnMDRBaFhVKzdWdTArdXE0aWJMT3VHQ2UzUjF5OGUKMzhoVkJTWXBUTDBsUTlQc1UvWmh4VzM0blpMdUtzWGxxbi9ITnQ3YzJwdmV1RUd5c0k3enlFWDFqRWIwck9NSgorR2p3UDFjT2tqWE5uQnF4MVFEVjBQRVJWYXdickh1YVNlQko4NWlENTNwREFZeVVkRmsrc1RKa1FCVjVrYm4yClk5RDkwZ3NJZ29LZXQxNSthUjlPMnRCR1c3ZzdaNlV5a0xEZ0s0TjZpRFR4bW5zUm45VFZNMkJkL1J5bzVXeFcKamZmWjd4d2s1eERaWUN1N1ZueFV1VmxVY2RnN3ZpTmJKTWNhejJDZ0QwK3JYdWpRTnVtaWFHTS83ZUd0a1lScAo5RUxpL0VNUFdiYnRTWFRUYmlsY2l3SURBUUFCbzFZd1ZEQU9CZ05WSFE4QkFmOEVCQU1DQmFBd0V3WURWUjBsCkJBd3dDZ1lJS3dZQkJRVUhBd0l3REFZRFZSMFRBUUgvQkFJd0FEQWZCZ05WSFNNRUdEQVdnQlFEQ0w5cXc4Ny8KeFZIMDhlV0c4NHBoRWREeXBqQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FnRUFXUGNqRUdSd25VSE5WVy9JaVlTeQpzZkxvL2JrTkN0dTI4NEdPTDV6RWZnbFBtVlZVelZIMkRLQ0xjd2xkM0tUTFhpaENhOTIvenNmcGJPZlJxZXRJCnlCRzVtcEs2V3U2N0ZBdmQ5TVN6TXhTOFU2cnZkQXBrUlVXcU0vVVowallWbzNjN21FeTlyUzVXelNwTWllVXMKblJoRTcybHMxbjMwbG1wcHJMRUhtMGtSWW05OHRXNm9IKzhNOHdISjJJUExIcHJBQkQyUWZFczhXUVJjSDR0RgpBaVgvdzQ3cGU0WkpZY29UMUZLVG9zVTM2SzROV3pOU3QyKzRFN0h3aGxLdTQxTFFOejF5NDdPNHJ4U05EbDJDCmlkdjBqTjZnMEQ5U3ZvMFlBTm1CMkprYXRXbGk3bXNXbFpCUTlwaFBkRWNxNnk5ZjRpTnNYd2N4OFlOdTMreDEKaWhlb3l5ZHZpV3FQZkx5d1JndUFtQ1E0TExMTEZ0SnMreWUvaEhWUmhhdm5BNXA1TlFWWm0yUGR0emVSV2JZMgpOOFpLc01obEt3UWE1MFJYdXRtazlOZGdBZGM4ZlVFanAyV01nUGdtNnA4RGM1UEFzcVh3Qk0yZlNPQnU5YWhzClpOVVdnajdZVTJwNmkyekxZenArTVh5VjQ2VTZOeGN4RXY0aEloOFJqN3N1bmVkZG9ucHNYeEtwUWxhYzFhUHUKOEtmMFpaTEZYZEQ1dEJvRlVPR2o1bU54QlpLMU1SZEJyK0hVMy9BZG1LMkR1VkNvc1I0M2ZkVlZTK0VmQWpzZQpnM2tCWldMN2w2WjJWaDlyWGpLRERjUlBPbDVGZk9nV3F4Tlhoa3NyaE5jZnowL0YyQ1ZMT254SlRDMjlUenQyCm5zOURzak1TY1E3REdrNjN4WCttZkRzPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS1FJQkFBS0NBZ0VBMEFoTXlMT2UxM2NCeFFOZk4xd2JiaWI3YzJzWGhaQWpZOHNZNVlGdTNSaHZuQm5NCmo3SU5zU1FZeGxSaUtyL2RXNEdFajhVNHd6M0N3bzFWcGNFMXJnWXQwNjB4aW85U3NEN0lqS01HVk1IQ1h1NGMKMzV3bUlCdzlmai9VSG1DMTJJaUo5WXNyY2pmcXBlVGlrQ2lMaTlTcmN3djJ0STl0Uk5iNVNrcXVadlVUVzNqOApDSWltdzNvd3hkUG9vdXl1NWxtQ1plbzdMTUh2OXJLWUlMK0U5K1ZuLzBSUHhXdlM1eFRCUkJiVmZvYjZsTk1hClErcnNEUlR6WFhBK1M3azhkaUdreWo2cWFHMXo0VzJoeHZQZENwaGFmdFhCeFF4OUgwUGhZY05hNUFTNjRBb1kKNUZPeUNHMG1yYlVUaHdSMFo1dFhac0MyY0MyMndUbjliR2pRL1JLa0hPWWpLVkdDOVpXMXdRVy9HUE00WTk3VgpNWG9ISGw4WU5PQUlWMVB1MWJ0UHJxdUlteXpyaGdudDBkY3ZIdC9JVlFVbUtVeTlKVVBUN0ZQMlljVnQrSjJTCjdpckY1YXAveHpiZTNOcWIzcmhCc3JDTzg4aEY5WXhHOUt6akNmaG84RDlYRHBJMXpad2FzZFVBMWREeEVWV3MKRzZ4N21rbmdTZk9ZZytkNlF3R01sSFJaUHJFeVpFQVZlWkc1OW1QUS9kSUxDSUtDbnJkZWZta2ZUdHJRUmx1NApPMmVsTXBDdzRDdURlb2cwOFpwN0VaL1UxVE5nWGYwY3FPVnNWbzMzMmU4Y0pPY1EyV0FydTFaOFZMbFpWSEhZCk83NGpXeVRIR3M5Z29BOVBxMTdvMERicG9taGpQKzNoclpHRWFmUkM0dnhERDFtMjdVbDAwMjRwWElzQ0F3RUEKQVFLQ0FnQUJycStBU0FPVzFuMkxMRlhPeXMzbC9DYTRianRJZHp2eUNLaHc0clVVMEtmR2FXY0FHbjZGMmpiaApFN21mZ3VHMVpieSt0T2Vhbkp0QW00Zi95U28zK0JEU3oybkJKeHVTRlUvbDQwT2YxOWxNanp4b2lvaThaYjRSCmtVNlQyRkJnS2VxRHM5WnNIQXVudjh3ZFFsYXVrTSs2SkhTZ1RUQ2pCK0lJT2NaalNzUVhUMGtxZ0lCb3dFbzQKcjFnSVNVVzQweXg4eW9Ja0FJV0NSenQzdUZUSTlHLzgzbjZPcUNxU2Q5YUFTSGI5aVBYcDBqTTZMV0l2VU9ZTwo4V052bFZYcDFxZlFndVU1NXZpeThBSUk2OXZ6dy91blh1OHNPc2VnUXhwRmdpRHdDeUcvd2hVbGM0L1RCWjcwCmRLeHR3UmwzNVFlMTZ3VU9yeS8xTEJUaWxZNXJHY1hXcG5scVVjNjQyM1laaFljVXNocnF1UFI4Y1lmOHNZL0gKU0ZHNzJhYUhsV2trcWt5OFJJV0JaMTFEMzZPRDNFZlpBV3JaVDNuSGw5b2toWmwxa2pWZEordmlQS1lGbFVObwoyaC9vdGp1ekg3RFNORG1kdXBaV2pYWWp0Y1Q0cXd3dnUrZlR1bkJIUjRPWm5wTlNKdmU0TlEwMFBkYnFzVmJqCjVhQzBtdW12QVpRV1R5NFNUQ0Ztem1uYTBaZEtTRDBadUFiWUNJTzZ5aWsxYXRuWWRReWJmV3RuUklqUjZnK0UKOXBvaThKYlFxRzhDSTZjOUllcndwZWIrRXpYMEoyclYrcVZ0Z2tVQnpvZVh3RFVIWXgyb1FIYW5TRGFSenN6dgp6MVZVMUt4eVZrMjJ4MC9HeUpTT3g3VHRJYTRqMzRtdDIrS05kRzFmdy9MeklETHhBUUtDQVFFQTBsZmdWYytzClJFcHJ1Z2xzOVh1T210OUFrRngyVXFhTlYvREZ1T3lPRndRQWppWkMzRGJzUzBweEFBaGFUcUhUekNTZFdQSTIKYWFaWlAvUlNINlJFeElJMTZBVTh5L01sbWZocXI2cVpSRmMzMVJBWG8wZ204ZG55ckxQTTVBRDdxVWhVZE8yTwoyY3Axcnhyck54bTAzMjBCMHgvNlZEWHZ2eGxjYjE4dXdxcnRubFlHa0ZiaTkwVllDeVJpVWxwRmlVc3dMSVk4CjZuMWVjcFFmN1lRWEJrVjNhUjFlVkdBam1ucHhBZ2c2M0ZyN0xrT0hEQmhTdTVKNER6Q2tKUmQxajZqcitGY3kKUDRLc0Vmb2V2akJtZjducVdCQ1NUdUoxQzdESEpNZnQrek9XLzlSQkM5ckVNY0RVb1A0WmxjK1VYcGJrR3o4YwpHK0tjbW5tQzV1bEt5d0tDQVFFQS9UQUVZZXkwZmI1L0pLRTZLUmIrcWE2M0xCT1diNTlBR1VmZktudFhpdnVnCkVycEtSSC9IWkNkbHhBUy9qTFliMXV1TWpsTjUrYjlXVmgwL2hYb2ZHSjdJMmNLL21wZTF5SkFBQTNTVm9hbHUKTDRtUTR0dVhtM1lXU25pQnE1S2hJejR0bmkrMnAvNVJRNFNLb3o2dnh4SjNlRFNQNzJGc2w0bDRZRENHbUpoTworMUdvVStLN1JEYjY2RTR6aDhZbThCTGRobFZ1TXJONVVWbUwyNW1ldzVyNGVXc3lIZHFUMllGS1pOZzFxUkpiClFPZ01yREFOL0IzT1ZtMTZnbjNDWWt4U3JwcGVCK2p5bWRXejRGOXlnSnJqdzEyRlJja1ptaTAzaFZYMXN1T1AKcVhER3MvZWpITjZmVnJEd2JnejFucFVhSTNrbzNjcVB6WXlxeC9FOVFRS0NBUUJsSDNBdmNQYTlvaFdtUzYvQgpXTlJYS210c3U3YjZ4eTEra0xkTnp4UUVocDBKdWVVODMzNjhONTZaeUdvNkVWeTBjUW5nWUJtK3N3V0hWSnRDCnNRT0tnWnNPMzYyNTB6eEppSDhwMHRkNlhuL3VBNTJKbHo5NTJERHR6RWI5ZW1lQ204NVdwSHFmdzlET0RSLzEKem5zQlN2T3NuMXdHcmlPRGVOclNoQzBEMDQ4SEp3NWl5Qkszay9QZTczQUJiWHF1cGFJVVZiamxkcjQ0VVhRRQpDUlBVcFFMaVd3Y2xnMDY2anBEVHpFY1g5dmw0NVdnQWJaVWdyaXJnQ3A5dllaYllLUHJBMnBMQ3E3eXpxODZwCjJyVDhSTklmNGwra1l3U29TU1dFVGtYZ0ZZNUxrYXh2MlkybTZiQ1BjWG0xWnlWS1VEcDd5dHhsbjQxd1NtaWYKUUpyL0FvSUJBUUNUQkZuWEh1S2pqWVU0bU9JenorVnFWRWlLc0lrUElkTFBtenRMNkxrcHUvajBSdll2RmwrSwptWkh2STY2eW4yQkZDUnZoM1RrYnUyMy9yUllnaEl3ZitMdTZMdXFoY3V5Y25IbFZpRklHd0dKNStoQ0dtbXBMClhHT1BOTEJmRjFLNEt6ZkQvZ0s3UStLZUtRMXp3MGZBZWNtanBDbmNINTgxMHg1eUJGdHpxaVZhcTh6cGdPT3EKdFo0MlhJcHhrYTgwZ2svZDNDZVVDMEVyNnFwYWhyWjQ4TGpOa0dCV2s0QjhzU1dvcng5aG9JWWFoMlFzYSs4Qwo4YS9KNGVKY2VYZnhLazVza3JoOU1WZ0YyZUNNTGdCSUN5aXNYZGF3Y2hpS3ZTemJJTkg5NFVPZmFSd0lqb2hKCmtEL0gwQkNjWnF3TlBKa2o5Z3V6MWhad2pmanhCdlRCQW9JQkFRRE1OcmJzRTlDREM0NDRqT093cHdab1l3Ri8KS2xpM0J5NmRWMDBQZmNQajdVd0p4dGcyNnZIRWg5TGpHSExuVldqMzl6bkFqYUpuejZLNG5VU2dHSFJzT05GRwordG1SZlRERWxOalNqbzk4WnJ6RzNUZUdlYU1iNERCeE51S2F1TzdYZmptS0N3cnZrSC9hdHY5N2NrV0ZqRkVwCllWc2xUMm14WkNIRCtkdGVFT1FKdmZ3MXU4Q2twMVg1dEJkWGlDbjJ2Mml0OHRSaXdsWDZ5RGVvSDVubG9IT3YKeFU0UEdJczN6MlYrM01pMjJxaVRIKytLV0NacWtERG10WmFjOE53VWVaTVlkcG5wWDc3Z0dTMjBrS0hZY1QyMApNZ2tpU21uSVNsdTAyeXpSSmxpRHlrdkdycHJFcG5SVittZWVwUjlXVVcyd2oza1k2a0NkYy9yVE40cXkKLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K\n    token: 642bmwx1u4bni65xv2uqiaj6495zi0hv49uhw317fvlq69ucgoo9zjhy0deymay8433ymmbaftwzr5p3o3ozxcbs4p2ows703ufl0f9opw8a19m80xw0kmnpx9ptqpgq\n"
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-09-16T11:39:44Z",
    "KeyValueList": [
      {
        "key": "Location",
        "value": "koreasouth"
      },
      {
        "key": "Identity",
        "value": "{principalId:2b189827-f0a5-4940-9067-25503b9848ee,tenantId:fb98dda1-32ff-48eb-a489-62777cd9ccd8,type:SystemAssigned}"
      },
      {
        "key": "Kind",
        "value": "Base"
      },
      {
        "key": "Properties",
        "value": "{agentPoolProfiles:[{count:2,currentOrchestratorVersion:1.34.8,eTag:db29d286-5d41-4be0-b26c-472200a89b23,enableAutoScaling:false,enableFIPS:false,enableNodePublicIP:true,kubeletDiskType:OS,maxPods:110,mode:System,name:workers1,nodeImageVersion:AKSUbuntu-2204gen2containerd-202608.26.0,orchestratorVersion:1.34.8,osDiskSizeGB:100,osDiskType:Managed,osSKU:Ubuntu,osType:Linux,powerState:{code:Running},provisioningState:Succeeded,scaleDownMode:Delete,securityProfile:{enableSecureBoot:false,enableVTPM:false,sshAccess:LocalUser},type:VirtualMachineScaleSets,upgradeSettings:{maxSurge:10%,maxUnavailable:0},vmSize:Standard_B4as_v2,vnetSubnetID:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tb3re4cev3tpmj6onblt/subnets/tblfcbvuf4dqbh9smm5l}],autoUpgradeProfile:{nodeOSUpgradeChannel:NodeImage},azurePortalFQDN:dns-1789558793456303805-pyi9o5lf.portal.hcp.koreasouth.azmk8s.io,bootstrapProfile:{artifactSource:Direct},currentKubernetesVersion:1.34.8,dnsPrefix:dns-1789558793456303805,enableRBAC:true,fqdn:dns-1789558793456303805-pyi9o5lf.hcp.koreasouth.azmk8s.io,identityProfile:{kubeletidentity:{clientId:b75814df-61e9-4f6d-8476-a2637710c8b8,objectId:6c70b044-5ba8-4b55-9d9d-041ec01df227,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/tba3u1gnvc7mh37jtcck-agentpool}},ingressProfile:{webAppRouting:{dnsZoneResourceIds:[/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/dnszones/tba3u1gnvc7mh37jtcck.com],enabled:true,gatewayAPIImplementations:{appRoutingIstio:{mode:Disabled}},identity:{clientId:a7aaba24-aec8-438c-8203-d49b45af4874,objectId:ef2aa35c-6376-4c31-a85e-3d08276a4001,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/webapprouting-tba3u1gnvc7mh37jtcck},nginx:{defaultIngressControllerType:AnnotationControlled}}},kubernetesVersion:1.34.8,linuxProfile:{adminUsername:cb-user,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDLxAxSdnehI5glyZeeDMtrLnqQnMbcC9j20KbF2Cfqinio09bWfVS0DF/hsi2TkcsTCWi4Z17/eHO3jYfVrIdg5aAn+1PAjrEPN9LChSeTqe2nI2bZrvW4L0ebzYWpjzXNF2c4KU0x181o+A8YQengU5C65hB6nBBKFg+mnefDM/BjHSrSDgRbFlXcUzmScJ62VCZZwG0YQzCrg6vdFDeWT5M34kN7pz5NAwJjri5TtRj4fWYine3IHOlJ1Q0Lqqd3N0NqcDvVm29zP4JsOGB1obvpnuU0IXdGs6e2K1Jbs+zHMUw88c10m8SSjuJPnXOAYwkbhl6IiQDtCuok4uLOpi5nyXGOXRO/oxhV1X4/7O5ttrQGnEksvUAdDTzQfMXpfKjPeX1toU9Iwq80XenBJ2vB5FVFV7cvVj40LATeKh+dMEjcRe2lePmBohdy3oQJhgJq/kPYMHtGW7hzVPe/xbwH3NEjgGiqDoI0meRt9oa5FVNqyI+KHlS+Y2kma/yga+HiBAGGcl7/oDIVAqiPpw0gs2r8YjbmsNDlqMJcQw8f5G3jRzpTcWcXP3nOzlS3wjzoX75tfUj5d1ghBFw2FKmqtEOLlAzReOfRfr9qVuv//4v1+MphjOfnzSfV9cQOmxLS+jfdt33ezdUlby7LAgeW9enlmcUDHkvVOW++zQ==\\n}]}},maxAgentPools:100,metricsProfile:{costAnalysis:{enabled:false}},networkProfile:{dnsServiceIP:10.1.0.10,ipFamilies:[IPv4],loadBalancerProfile:{backendPoolType:nodeIPConfiguration,effectiveOutboundIPs:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth/providers/Microsoft.Network/publicIPAddresses/4d404521-548b-4bdd-ab61-e14ea4a43c7a}],managedOutboundIPs:{count:1}},loadBalancerSku:standard,networkDataplane:azure,networkPlugin:azure,networkPolicy:azure,outboundType:loadBalancer,serviceCidr:10.1.0.0/16,serviceCidrs:[10.1.0.0/16]},nodeProvisioningProfile:{mode:Manual},nodeResourceGroup:CB_koreasouth_tba3u1gnvc7mh37jtcck_koreasouth,oidcIssuerProfile:{enabled:true,issuerURL:https://koreasouth.oic.prod-aks.azure.com/fb98dda1-32ff-48eb-a489-62777cd9ccd8/c0271972-08cc-4988-b847-0d0a9fb97b8e/},powerState:{code:Running},provisioningState:Succeeded,resourceUID:6aaa800ec7aec500016591ee,securityProfile:{},servicePrincipalProfile:{clientId:msi},storageProfile:{diskCSIDriver:{enabled:true},fileCSIDriver:{enabled:true},snapshotController:{enabled:true}},supportPlan:KubernetesOfficial,windowsProfile:{adminUsername:azureuser,enableCSIProxy:true},workloadAutoScalerProfile:{}}"
      },
      {
        "key": "SKU",
        "value": "{name:Base,tier:Standard}"
      },
      {
        "key": "Tags",
        "value": "{createdAt:1789558784,ownerCluster:tba3u1gnvc7mh37jtcck,sshkey:tbk0pp1fi356dlisll7m,sys.connectionName:azure-koreasouth,sys.createdTime:2026-09-16 11:39:44 +0000 UTC,sys.cspResourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tba3u1gnvc7mh37jtcck,sys.cspResourceName:tba3u1gnvc7mh37jtcck,sys.description:Migrated from on-premise K8s cluster (v1.32.3, 2 workers),sys.id:k8sazgcp01-on-prem-k8s-cluster,sys.labelType:k8s,sys.manager:cb-tumblebug,sys.name:k8sazgcp01-on-prem-k8s-cluster,sys.namespace:mig01,sys.uid:tba3u1gnvc7mh37jtcck,sys.version:1.34.8}"
      },
      {
        "key": "ETag",
        "value": "36b817bf-0b40-4150-b353-9675a97f66ee"
      },
      {
        "key": "ID",
        "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tba3u1gnvc7mh37jtcck"
      },
      {
        "key": "Name",
        "value": "tba3u1gnvc7mh37jtcck"
      },
      {
        "key": "Type",
        "value": "Microsoft.ContainerService/ManagedClusters"
      }
    ]
  }
}
```

</details>

