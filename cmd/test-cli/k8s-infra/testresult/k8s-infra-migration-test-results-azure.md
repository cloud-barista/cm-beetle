# CM-Beetle K8s Infra Migration Test Results — Azure-Busan

> [!NOTE]
> Full lifecycle against a real CSP: recommend → validate → migrate → list → get → workload → delete → residual.

## Environment

- CSP / Region: azure / koreasouth
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (cd1f3a2)
- Git Commit: cd1f3a2
- Namespace: mig01
- Test Date: 2026-10-07 16:41:15 KST
- Cluster ID: mig02-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 22ms |
| 2 | POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation) | ✅ **PASS** | 36ms |
| 3 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 4m45.047s |
| 4 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 3ms |
| 5 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 4.55s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 30.549s |
| 7 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 5m27.617s |
| 8 | Residual resource check (Tumblebug) | ✅ **PASS** | 5ms |

**Overall Result**: 8/8 steps passed ✅

**Total Duration**: 10m47s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 22ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.35.8)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=azure+koreasouth+standard_b4as_v2 nodes=2

### Step 2 — POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)

- **Duration**: 36ms
- **Status Code**: 200

- ✅ Target K8s infra model is valid for migration (0 issues)

### Step 3 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 4m45.047s
- **Status Code**: 202

- ℹ️  nameSeed: mig02
- ℹ️  async reqId: 1791358875687531656
- ℹ️  cluster id: mig02-on-prem-k8s-cluster
- ℹ️  elapsed: 4m45s
- ✅ status: Active

### Step 4 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 3ms
- **Status Code**: 200

- ✅ migrated cluster present in list (8 total)

### Step 5 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 4.55s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=azure+koreasouth+standard_b4as_v2, nodes=2)
- ✅ version: 1.35.8 (recommended 1.35.8)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 30.549s

- ✅ kubeconfig obtained (server: https://dns-1791358885035685337-n9757uld.hcp.koreasouth.azmk8s.io:443)
- ℹ️  auth method: static token in kubeconfig
- ✅ API server reachable (v1.35.8)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 2)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: 20.214.7.115
- ✅ nginx served over the LoadBalancer at http://20.214.7.115/ (attempt 1)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 7 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 5m27.617s
- **Status Code**: 200

- ✅ deleted on attempt 1 (5m27s)

### Step 8 — Residual resource check (Tumblebug)

- **Duration**: 5ms

- ℹ️  VNet mig02-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig02-k8s-sg still exists (known gap)
- ℹ️  SshKey mig02-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for azure koreasouth (source: v1.32.3 → target: v1.35.8)",
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
    "version": "1.35.8",
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
  "id": "mig02-on-prem-k8s-cluster",
  "uid": "tb07qlatru1cgnjdlo69",
  "name": "mig02-on-prem-k8s-cluster",
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
    "sys.connectionName": "azure-koreasouth",
    "sys.createdTime": "2026-10-07 07:41:17 +0000 UTC",
    "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tb07qlatru1cgnjdlo69",
    "sys.cspResourceName": "tb07qlatru1cgnjdlo69",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig02-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig02-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tb07qlatru1cgnjdlo69",
    "sys.version": "1.35.8"
  },
  "systemLabel": "",
  "version": "1.35.8",
  "network": {
    "vNetId": "mig02-k8s-vpc",
    "subnetIds": [
      "mig02-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "mig02-k8s-sg"
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
      "sshKeyId": "mig02-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 0,
      "maxNodeSize": 0,
      "status": "Updating",
      "k8sNodes": [
        {
          "cspResourceName": "aks-workers1-10476153-vmss_0",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-10476153-vmss/virtualMachines/0"
        },
        {
          "cspResourceName": "aks-workers1-10476153-vmss_1",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-10476153-vmss/virtualMachines/1"
        }
      ],
      "keyValueList": null,
      "cspResourceName": "workers1",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tb07qlatru1cgnjdlo69/agentPools/workers1",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tb07qlatru1cgnjdlo69/agentPools/workers1"
        },
        "ImageIID": {
          "NameId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2404gen2containerd/versions/202609.15.0",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2404gen2containerd/versions/202609.15.0"
        },
        "VMSpecName": "Standard_B4as_v2",
        "RootDiskType": "PremiumSSD",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbh9oq28l3i7cb1062kn"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Updating",
        "Nodes": [
          {
            "NameId": "aks-workers1-10476153-vmss_0",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-10476153-vmss/virtualMachines/0"
          },
          {
            "NameId": "aks-workers1-10476153-vmss_1",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-10476153-vmss/virtualMachines/1"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://dns-1791358885035685337-n9757uld.hcp.koreasouth.azmk8s.io:443",
    "kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUU2RENDQXRDZ0F3SUJBZ0lRTHI1UVAwTmRxZW9wTUNoQ25SY2hlakFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWdGdzB5TmpFd01EY3dOek15TURKYUdBOHlNRFUyTVRBd056QTNOREl3TWxvdwpEVEVMTUFrR0ExVUVBeE1DWTJFd2dnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUURwCnJVc2pVaUJTQ2dsOFVhSHZ2S0xyYTVhNFBGdlNvTk1OcGU3dllXQTVra0RQd2pIcCtkR2QvRjIvazd6V3A0YWQKZSsxazFTd1ZaK0tKNEdCMGpITzZiMXlXZkIyd1JvQW9oTStMT3NyRHRZbUNSVGpQVjdudUd0ME9xMVRzaFY4bApkZFpWeWxJMndtc3pYMTNHMzErQkZYUzB5QTc3R2ozNzlNRUdFM3g2UG5pdENxd1V0V0Z3b0s5VmVDMURickVsCmQyYnFnRU9uNXpDUDU2M05XU05GMFdFN3pUTkovYTdTQXVXNEdjV3ZSdnNDK2FNYjZBd0RKM0NqQjdMTkVHdTkKTGlqWG85Q2pIZm5GY2lyalRpQWZOQWlzUXR5UisvekpKM21RYmZvWkx2MHJtcm1PcmgwbW9Ddm5UclhwQTE3TgpzS3Z1SlZRWmNaaXcwR3NPcklFZEJ2STZiUCtsZ1FUVi9jUzUzUkFUcTR1YjFWcFQxcHg2Y3JWOHRTcHAyTk9tCmpIaTBaK2RSdU01S3A1c0k1T2RxRjRkWFhxaDNEb0diektiSFNCU3NwQ3dIeVRDdEY2RlBQLzdkdVFoYXE5Y0kKempLbWNqVGFhaXpaLzExTVBKZjRlUFkrQW82NnY1VzFSRWZ5OGZlREtGc1BsQzJteWxLVEtEc0VBamFrbVV6ZgpuTkRVSDY0VitzR2NuOXZsS0RSRFdrdGNZRkVDNkhCay9PZVk5aGpCSzI0RFRQVXZoU3hUbnhFOFhjd3k3aTZNCnNCVmY4ZlZXUmk5UkZROU11amxMWGFzcWY2VFdCdHFiU2dWS1orZTlvS29haHJPL0xrRHFwVFNWeGJ2SS8rY3EKK0o5Y1pGNUY0dnp6Wk12ZHI4Wmc0dHFVZk1aajRyWHFvM2xPT3liYUN3SURBUUFCbzBJd1FEQU9CZ05WSFE4QgpBZjhFQkFNQ0FxUXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVha0NpTklDelBkcVlJTlRICm9kajJ4aE9QSG13d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dJQkFFVXorNDY0L0xpcnBlckRkcDFQZGhXQS9HNmUKSXJEeE43N0hVNkE2N1FiMmE4NzlkVnJ1TzB3MG0xbS9vNHhhL05manhZZWlqT3ZKUUgxVTlpQlJZK1I3M1R0WQo3ZFRMdUNrS0kyRzhhVW1MZUVGelZXMzJJT3ZJUVdxamM4V2swMHREdjRYVWZndHp4aXBYWEdCRXJmbnlZM1piClBCTTN4YnlUVkRCbmRvQ2pxOVRvYkVBSHlHczJMYXl5TnhHYTg2UnFoTjlhcVFGQ1k5a1IvK050Z3FZbnVtbHkKeGxkTXg5ejAyUjltem5sN1FSZlM0emNnNFpIcFMrUnhiSjV2LzlGNElaZmk1djE5OHFQTWRPVjA4OU5qVXc4VwpnM09nWTFXbFl3NWVVNVpBSGt2QkdVNEUzZnNwaGRoWnlGNXYyUmMxWU5ub3pZdEpHRUNaRnBwdU94T0JESUQ4CmtKRG1zaVB3L1c5SlIwZWM5WDJ0SVhuS3d6YTJrak5YYjF2eE9QampoNWR4TjRidjZ5Z3NFTFlzSTZnb1NtTzUKVjFPOWRQVWQxdjdCbUVieVhEaERWT0pmL2Jia0YvTDJsbDNDa1NOOGE2SWxJQWs4ellta2xvWGIzbnpvS1pYRwpBd1VFOCt5YloxdHBUTisrRXp5a3NKMFRUSG9MaTRwTUtSU2JYYWRKK2dON3ZqSHNYZ0VwNXNEd2VLdG1YVnpCCjd0QVNMWlZHeFBvdW1mLzdDMnlPNUVuRnFabFRKSzc5UXUzQjJ3VERsYTMxSURjT0MwSEpXZWNNbHJFZlUrRUoKMDRkaHZnbmVtczdUOEkvbHEzM1RBSFBOemhXbFkreFQ2cEZaTUlkaU11NDVTeHFVUDRLNjdva0p1UlZ0OXhoSApsaTNTYkJjdmswbzhZalBxCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://dns-1791358885035685337-n9757uld.hcp.koreasouth.azmk8s.io:443\n  name: tb07qlatru1cgnjdlo69\ncontexts:\n- context:\n    cluster: tb07qlatru1cgnjdlo69\n    user: clusterAdmin_koreasouth_tb07qlatru1cgnjdlo69\n  name: tb07qlatru1cgnjdlo69\ncurrent-context: tb07qlatru1cgnjdlo69\nkind: Config\nusers:\n- name: clusterAdmin_koreasouth_tb07qlatru1cgnjdlo69\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZIVENDQXdXZ0F3SUJBZ0lRUW54NTZXeEFTalVKOGVsVzAyUDhJakFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWVGdzB5TmpFd01EY3dOek15TURKYUZ3MHlPREV3TURjd056UXlNREphTURBeApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1SVXdFd1lEVlFRREV3eHRZWE4wWlhKamJHbGxiblF3CmdnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUUM2dlRPWnc4UklJS3ErWVpmRy81RlEKWENicnlJaVJTUDFEUHB3dmdDR1U1T3BIS05aem1FZDRqQ2QxR3dGL2gxZHozR3A3UWtlRXB6Qk1hdlJ0TEtyeApvanRaTzF0WHdFR2tRWDN2anJsRGlMUUNIUGwvRGwvem96Wm95c2VPKzlWUVB0VGUyYnZHUWVncmEzZmJLNHNtCkJWbG1lOWIrSnVjczJyYTdsZTIwVS9HUFgvLy9UaytFd1EzcGRXU09XVkFaV3dMZTUwd0xOaTdmYkFhWHN3M1oKbGpoTXZHSXdlaUtNTXdnWGFzS3lHbkJkYStIdk1yaCtEOHExVit1S2grandURW5TamNTSERQVHk1L3Faa3pGWQozbXNDaXRXWTFxQmZNbmFGc3hhVjRYZnpkY1d5dGdObjhrcSswbTkveHpyWDQ4S3ZmT0MxdjlGS2xhUkpsMGJtClhPWXMvN2FOeHVIQ2hMTWFvem9FamovSkduR1hpYjFVL2ZOaExSWUo0Rk1LQjliNTlpdm83Tm1NaGcwdmhWZm8KWEg3T0RqUzlKQmVGYnZMQ1ZZUEFKRTZERUpBNEw3VUVKajFtaEJFUm5GTXpPMEVUTFdvVGI3d0JQUVlMSjZLWApqQzdBUTd4S2VlSm1XWDZwYjB4R2Z3blRTUEk4NEtzbFRieFdmNTJyTUxIcVE5SzZ0UEduR25wVkgvN2VITGdSCjdua05qa3Q0YjJ2allqK3RlYW55TitrU0tnNFhYU2laaGgvaWhIbFJrZ1E2NHBCS0NYYTNYMmN5N0hPRHdMaVkKVHIxK1oxcHVMcUpzUkplV0lVYThhNy9GY1ZGY1Rkc3RwSHJmQWR4Zzhub0tCQnlFVmpzR3FRVjZFWHNqZzFNQwpkUkJtTUxSUEROL1lKalpma3ZEdWhRSURBUUFCbzFZd1ZEQU9CZ05WSFE4QkFmOEVCQU1DQmFBd0V3WURWUjBsCkJBd3dDZ1lJS3dZQkJRVUhBd0l3REFZRFZSMFRBUUgvQkFJd0FEQWZCZ05WSFNNRUdEQVdnQlJxUUtJMGdMTTkKMnBnZzFNZWgyUGJHRTQ4ZWJEQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FnRUFsOVMrdStDSytYczY1TlBabmlaOApxLy9rRzdzRkxIa0daMXFoaWRQNmtoVE41R2RBUGswZzNXbWdsZ2d3UUtNSHhIYjdSRmRwbDE0SzcvK2QrZmlyClAvd2N0N1NxaE9SNmR4TEl4alkwbXZQcCs5dStJcVFVcHlpOVF3bWt6TGg2TkM4cFZQRVhteWpURHVTMHh2Q1UKVitiQ2dhU0N1VUZWY2VVUHJ4azg0QVN3LzRZN2hGTDhOanlrTUJ6OXNHUlR3WTllU29LdVlDekNRQmpubHYrVwo2UEx5d1ViQmhzZXNaU0FFUzN4RnJtNjk0UlZQQVowR1U1SE5kQVI3Rk4ybzVyc1V2V0RSckZRZG56NzRNQmlQCmw2YjMvZGRHVkIxcnBIZVJTN2tJbG0wR0ZqR3lqSjVEdjQzdnpScmV6c0R3MVRJa1hmVk4yOFJ4MlpFeG8zekgKZkZET05GWTBncGNyb3FNanEvMlo2bG9WME5YSGMxcjZDem8wRzdpbEpnOGRyUERlcjRVZTdwWExqRUpUNEorRApxY1Zsb1ZoVUhLaDdBd0hrNktIbitRd056WVNjTUNRMisvS214b25kNWJMMDBqSFprRGk1SGp4dVBIVHJuRnFOCm0vWHNlOGNlQ2lrUmtsSmVCM3NGdXNGeXZzTHRPeE5pRlBYQ3NjTmNSTjVOdG9KUzBiNlBFeE5hY01XcjBMaC8KVGZIaWducnNQV0hyeU9RUU1PVmhMak5FbWtwb05yN2IyOWpMSVg0VVRRS3pOeTFKQlpsT0FjMlBhTDhyZ0h1dAp5cC9rd0dWZGY0a21rckQvcGprcFpwRlFQeENJUDZjN3VHb2ZZaWl6ZkFzbk84TlpKaE9GODRzNTZSREVpSVRCCkZVQzIxV0lUcFpDUU56eGlPbDdnNHJ3PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS0FJQkFBS0NBZ0VBdXIwem1jUEVTQ0Nxdm1HWHh2K1JVRndtNjhpSWtVajlRejZjTDRBaGxPVHFSeWpXCmM1aEhlSXduZFJzQmY0ZFhjOXhxZTBKSGhLY3dUR3IwYlN5cThhSTdXVHRiVjhCQnBFRjk3NDY1UTRpMEFoejUKZnc1Zjg2TTJhTXJIanZ2VlVEN1UzdG03eGtIb0sydDMyeXVMSmdWWlpudlcvaWJuTE5xMnU1WHR0RlB4ajEvLwovMDVQaE1FTjZYVmtqbGxRR1ZzQzN1ZE1Dell1MzJ3R2w3TU4yWlk0VEx4aU1Ib2lqRE1JRjJyQ3NocHdYV3ZoCjd6SzRmZy9LdFZmcmlvZm84RXhKMG8zRWh3ejA4dWY2bVpNeFdONXJBb3JWbU5hZ1h6SjJoYk1XbGVGMzgzWEYKc3JZRFovSkt2dEp2ZjhjNjErUENyM3pndGIvUlNwV2tTWmRHNWx6bUxQKzJqY2Jod29TekdxTTZCSTQveVJweApsNG05VlAzellTMFdDZUJUQ2dmVytmWXI2T3paaklZTkw0Vlg2Rngremc0MHZTUVhoVzd5d2xXRHdDUk9neENRCk9DKzFCQ1k5Wm9RUkVaeFRNenRCRXkxcUUyKzhBVDBHQ3llaWw0d3V3RU84U25uaVpsbCtxVzlNUm44SjAwankKUE9DckpVMjhWbitkcXpDeDZrUFN1clR4cHhwNlZSLyszaHk0RWU1NURZNUxlRzlyNDJJL3JYbXA4amZwRWlvTwpGMTBvbVlZZjRvUjVVWklFT3VLUVNnbDJ0MTluTXV4emc4QzRtRTY5Zm1kYWJpNmliRVNYbGlGR3ZHdS94WEZSClhFM2JMYVI2M3dIY1lQSjZDZ1FjaEZZN0Jxa0ZlaEY3STROVEFuVVFaakMwVHd6ZjJDWTJYNUx3N29VQ0F3RUEKQVFLQ0FnQXpBYTBQaWc2YUdQb3FhR1Bad2tNQjdUbmdXM1VhSWhLemhHMks5L2UwUU5iUk94RmduNjZkK2NKcApWVWdTZW5oanVFZ0poUEFlQnNERmpzMVc2TVdFbk9pVEZnaDhMcEovZURnYThDUWdrejEvK2tRWDlXZzJGMVdzClIvODhTZ05aR09CeVFvenN1V0hlUWt1SnZSb1Q5NUFVMHl6RlhRRDJHb09oTjNHbk9PcVpYQUtEelBNaGNzNEcKLzc3SXR4S082bklkU3BaVjFhWDE5djdSL3VHQWxteW9ZU2g4ejArU2RmbjIwMzE4d29ZQXJwQVhFd2Fid0ZxQgpNQk94VWpCSk9yMUxXdGZjZjZpRWxPSzVnaUFQZ3lrOWw0aUNaOVdFU3pkZWk2MVdvVEJ2ZXJMV2ZxcnRnSnYwCkJXelpQU1VRaFREUFdFL1V0U0ZJdXhrVjFORlRjUDlJNCs1dVJBU2o4REFPdUNIUHhWSU1LcjZTU3JRc1p6MmoKTGpCM3I3UE1tTHRHZE8zbUJwUXZyTWtrbE9qMDNiZXZ4L0ZTbTlJOTVyejVINWFjckZwTGNjRUtVSERvUlhqRQpERU0rTEloMUZXZGtDT2xpMWIzd2RhOHh6eFpsYlFhUjhFTmRqUExpMDErM005WENNb2VLM2xDak84VzJiSTBpCkdpa2NxemhqUjVRZGVPRTJRbG8wSllGWEVwSUxRQjZkdXZ0L0Rxd0FZajNrWTNKOWp5N2xmdE4vNk1VYjB5Wm4KZmY2OG1ud0thQ21ZWXkxeVl6Y3NzN2VIcmFVSXdnTGxoSk95K2s4VUc1aDg2UVBpUUk5MWthQ0JJUi9yR0dKQgp0OGUyR01jVlVubFJQYW9YemhXcjhUL2JMREQyM1RWVlFybXFEVEJ3WlZoempCbXNQUUtDQVFFQXh0eUNQNUU1ClJyNEt3M21HOC9TTWZibW9uR2xsTVhwejNFbnVqODhsMGk2OWlFbSszT29HM3dXZndZTW9yd3BLNXVPYU4xQ2MKTGtheHQrVUdFbWtvWlh3SUlrUkRWMmVaNERtZGI2UEV4dXhzdlR2T1pCZklsQnBVTXZZMlpKTnk1bEJuYTgyMAp3YlgvY2ZqVjVBL1dvS3o2TDk0c1BtSTIvTmdKN0VQRHV1RUcraDE5UW5sZ3pLeEVFc1FRYXV1RjlGWjZWdmwrCnpKKzNUUDd6NlZodE5nNE1mMDhIMGlLWStaS0l2MU9uUUpGbC9lTUkxamhERGkwRlVydVFXN0VSdTlHSWVsOG4KR1RTc0pNNDlmYkJJZnBzOFJ0d2tTRjByM0NuUmUyYmR5bWp6Zy9rSmk3TktFckJKTUNyY09MdHhsZmhGN1huRQpwbjlHL1RBNU10bVQxd0tDQVFFQThHVUZ5Vy9TRGtxbE9TcWE3cmQ5eE1vQmwxSFlmUWxVeHRuSy93cnhmdUlKCjRJNXQ3M05vV01McTRXTDY1U2R6QVRmQzZTTTNZY1p5MDYrUTdDVGZFSHNNSzN1MHdYMUZIVWtOTGg4WXdNVkUKU09aOHA5WkxDc0dVQy9nRFdhVDYrOEo1bzBXVWFkNnh1UTB0Qitsa0xrQklrMGxUOFRuNzA1Z0Y4SG0yTkFmMwp1VzloS0JFZTAwK3lhYUtQK1NUNHYwbnlrTmNvbktXM1ltYnhMTXprcEZKQXpoa2hwT1M2MUJ4SFcvSlV0YjdtClI3MHA2RVJQUkwrMDQyaGNXS1loaGZTdE5wbHYrKzN3QmNBWncrSE00c0FIQXloMGlLVVdqcTNWNFVjYmxTTVEKQ3ZvR2huZXkzQytNOXZCOE5EV1BxdStsMktVekkxZU1rcCsyMldjRkF3S0NBUUFkTVgvNWZ3Tk9qRldlTHJnTwpGa2VOcVhURVNZQ3VpOUI0ZFc0Tk5KYlR6VUxMd0MvQjFLVmZsajdHdE0rREt0cW1IS0dtSHpKZVpNaHdPN0h4CmYvaUxOWE1vUEtjNkxKNWRXWGZ3VExWMWtuM0FKZ2g5anpSRjhidU1vN0tHZjFMdDFyM29DSkhSb2pzMjZ4WFQKWGpQZThLWGw0eExSc2E5cTNQaFA4LzdHNzRRTHNjcVN1S2pxUXh2WE9XdERsd0hhUmR1OTZ2Q2ZiNDhFUWFWagpDelUxSzBqUVk3UzlONUVIaW1pQzFmQUd2WkdnZlBUdUplSktNWGZIbG50eHlvUFU4OFM2V2ordUpwcGk2TmdFCkhKMzM0Q2d3S3Q4MHRHRGMrRktsY0F0OVRIejdVZE1CN0RjaW1UaVZWcVZ2dXF6SkhHSi9vMUdvTTZValJESVIKY2ZtUEFvSUJBUUN3aGI0bUdOSnFaQzNIT0oza1JTRTViQXlYRFNhblNqN1NkZ00ySE5jUWtQTW56ZTYwTWs2aApQeUF5dVBmN1pXaHhzSHlUcmFSeFM4UHQ1YkhKZFpuSEJGUC9haXRsR1pPeTJYMndMRzJFd0ZaUUljL1BmdkJECnlibm93QW5WdmV2L0N1TS9IVXpLSVJqb0JlRFhPbTZ4OThwVFBUbENuWTFwNXV3VmxZcFIwcmY1bHV3RSsrMWkKeGRjQThybzkvUVFia3pWblhsY1lFQ2dGUGdwREY0R2RtRGkySG9ZeUN6T1ZwZDVRaWpYaDczZ3huWm8wVUZUVQo2dUR0VENqamY3cUdIeEZDVXBHRXNVZmJNU0M2VWdpZFNOemRXTXVadTRCQTdTMXM2MFk1MFRGcW5nWkVuNGpiClhwOW9oVmJ2RDVXYnM3Wll5NVl0a1pCTDZyUGVHR1A1QW9JQkFEelNMUjF1K3I4WnErNFVhZkU0SVZGMmVTNG8KU2RYbkNHRmpkWVp3WHlBeWZKc09LaTVDQlNub2NTKytuUCtVbXBoMFZNWExkVzhJZGRucXhnQlQxWVlubWxEcgpwMmRVTVlxdlpnb2FTdEQ2SWIzQm4wd3VuNFdnU2RrOE8reE5admJXa3gwRGgxdU1lbWlsYWNraGNVbGpadmFNCmpXYTc1Z25zbW1BNk44ZkJSY3VJeXJvOXl2S3B6a0xrVEtSdmlEaG1MUGJveGhJMnZUV3FqbDJLL2FWeEY3T3IKK3hYZytQQ1NWeHNNV2xvQ3E5SzkzZFNhOGtxRHc1V0o2VUdqMnpSMWl0am4vSzNsb1R2dVkxb2VLUzlGcHFPNApOa3dvSE9mU2lyejBBbHd4c1VIV2xhMktiNU9KdURSb1hkNjZ0NlpINktyRTBjd1ZvR3lpWENBenZ3VT0KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K\n    token: f2orrx81vh474tzdz0a5da1mvet8u6i6ui2vgzavwmmq87hzfwbjqlimmd8ilaq9dw6bhicdqimec9zj8hvsq2qzj2statwv6i7vy53t84qkkz2g4vrbnc8xt4vwh0ok\n"
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-10-07T07:41:17Z",
  "keyValueList": [
    {
      "key": "Location",
      "value": "koreasouth"
    },
    {
      "key": "Identity",
      "value": "{principalId:e99c6063-3753-47f3-80ea-e33b557c5965,tenantId:fb98dda1-32ff-48eb-a489-62777cd9ccd8,type:SystemAssigned}"
    },
    {
      "key": "Kind",
      "value": "Base"
    },
    {
      "key": "Properties",
      "value": "{agentPoolProfiles:[{count:2,currentOrchestratorVersion:1.35.8,eTag:b7a178a9-f728-4f33-8f4f-fd15ed9a6704,enableAutoScaling:false,enableFIPS:false,enableNodePublicIP:true,kubeletDiskType:OS,maxPods:110,mode:System,name:workers1,nodeImageVersion:AKSUbuntu-2404gen2containerd-202609.15.0,orchestratorVersion:1.35.8,osDiskSizeGB:100,osDiskType:Managed,osSKU:Ubuntu,osType:Linux,powerState:{code:Running},provisioningState:Succeeded,scaleDownMode:Delete,securityProfile:{enableSecureBoot:false,enableVTPM:false,sshAccess:LocalUser},type:VirtualMachineScaleSets,upgradeSettings:{maxSurge:10%,maxUnavailable:0},vmSize:Standard_B4as_v2,vnetSubnetID:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbjvs627fsvq9c8qjdmm/subnets/tbi5s3q7etm8cahfjubo}],autoUpgradeProfile:{nodeOSUpgradeChannel:NodeImage},azurePortalFQDN:dns-1791358885035685337-n9757uld.portal.hcp.koreasouth.azmk8s.io,bootstrapProfile:{artifactSource:Direct},currentKubernetesVersion:1.35.8,dnsPrefix:dns-1791358885035685337,enableRBAC:true,fqdn:dns-1791358885035685337-n9757uld.hcp.koreasouth.azmk8s.io,identityProfile:{kubeletidentity:{clientId:aa26e04b-b027-4a58-9065-6cc825160177,objectId:e2ef4ebc-4130-4afb-8cda-cc9a3825c0ee,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/tb07qlatru1cgnjdlo69-agentpool}},ingressProfile:{webAppRouting:{dnsZoneResourceIds:[/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/dnszones/tb07qlatru1cgnjdlo69.com],enabled:true,gatewayAPIImplementations:{appRoutingIstio:{mode:Disabled}},identity:{clientId:7768eff2-560d-4c61-a3bb-4740fccf8536,objectId:845c8887-a41b-4b29-b4a3-c77e8020237f,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/webapprouting-tb07qlatru1cgnjdlo69},nginx:{defaultIngressControllerType:AnnotationControlled}}},kubernetesVersion:1.35.8,linuxProfile:{adminUsername:cb-user,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDA7EqTWJv+hdZCkyjquj4TaYf5mYZQ4lsvvjLwuWaFV6BKS0Owt/B+rMl4PYS6zuK2hBjBFFYVY/hmPyhUiSzyX93wWSrHk2axbJyL++kJSVQ0dmO8Tcnr47K8Tp6WI+Ex72s5vIeKnltDbjLIsYHcisuB0bPNfDREFO3aUkKjub3l40Nowq2SmlOV4pB1MpUigKUPrh4QVrRhF0vdcXSzZ1wYCfWiPzF/fU0d+7f7raTwc8bDgHRSYsOQ7r4Je8XV2HPyG4kMU1z55l1RZkWE4Nhn2Ojy+H2nq/PGt+LB9M10U7f1iykdIG7N40SoyGnTekyPDZ9/o8E8yOjaMMAhBLFy9xs1VTARi2ZhJUKF/8MQRV/hZYMEZclYH8CGCiFvDBsU0xCgKjtsso8WO2zoGcmdZTDQb4+gy/wvvHRFGQPnTY79MfKPLtUu3B7c0uDVGtS5reeDlDsU2Is2e7pFX29sMZ3r9C8w61raTlUUagTLNbiHkTPiXgU4MFHDIeDpqKP5dBF6FW/qq4whuwQJSpjVg7H/m529RCo+5WgtnoZNEbvVZbXi5R69AEemoEvanua+ThwtytrE2yOsu7hVo0kvMjezxgodoyZLxwxYi8trvvNJxen3Ck2c6r/AzVx0slcjguz3uNi0RQf0hQQ30gva4VzDVjZwLWQx5Ngz6Q==\\n}]}},maxAgentPools:100,metricsProfile:{costAnalysis:{enabled:false}},networkProfile:{dnsServiceIP:10.1.0.10,ipFamilies:[IPv4],loadBalancerProfile:{backendPoolType:nodeIPConfiguration,effectiveOutboundIPs:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.Network/publicIPAddresses/2f1608e8-04bb-43b6-a4cf-d880526113d3}],managedOutboundIPs:{count:1}},loadBalancerSku:standard,networkDataplane:azure,networkPlugin:azure,networkPolicy:azure,outboundType:loadBalancer,serviceCidr:10.1.0.0/16,serviceCidrs:[10.1.0.0/16]},nodeProvisioningProfile:{mode:Manual},nodeResourceGroup:CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth,oidcIssuerProfile:{enabled:true,issuerURL:https://koreasouth.oic.prod-aks.azure.com/fb98dda1-32ff-48eb-a489-62777cd9ccd8/7410c9fc-0ac2-4592-89ff-28f5bb32a74c/},powerState:{code:Running},provisioningState:Succeeded,resourceUID:6ac5f7a8eb7c6b00013d1377,securityProfile:{},servicePrincipalProfile:{clientId:msi},storageProfile:{diskCSIDriver:{enabled:true},fileCSIDriver:{enabled:true},snapshotController:{enabled:true}},supportPlan:KubernetesOfficial,windowsProfile:{adminUsername:azureuser,enableCSIProxy:true},workloadAutoScalerProfile:{}}"
    },
    {
      "key": "SKU",
      "value": "{name:Base,tier:Standard}"
    },
    {
      "key": "Tags",
      "value": "{createdAt:1791358877,ownerCluster:tb07qlatru1cgnjdlo69,sshkey:tbh9oq28l3i7cb1062kn}"
    },
    {
      "key": "ETag",
      "value": "be9b0b92-0c64-471d-8cb8-d17b110ccb9e"
    },
    {
      "key": "ID",
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tb07qlatru1cgnjdlo69"
    },
    {
      "key": "Name",
      "value": "tb07qlatru1cgnjdlo69"
    },
    {
      "key": "Type",
      "value": "Microsoft.ContainerService/ManagedClusters"
    }
  ],
  "cspResourceName": "tb07qlatru1cgnjdlo69",
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tb07qlatru1cgnjdlo69",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tb07qlatru1cgnjdlo69",
      "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tb07qlatru1cgnjdlo69"
    },
    "Version": "1.35.8",
    "Network": {
      "VpcIID": {
        "NameId": "tbjvs627fsvq9c8qjdmm",
        "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbjvs627fsvq9c8qjdmm"
      },
      "SubnetIIDs": [
        {
          "NameId": "tbi5s3q7etm8cahfjubo",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbjvs627fsvq9c8qjdmm/subnets/tbi5s3q7etm8cahfjubo"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "#aks-agentpool-21615671-nsg",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/cb_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.Network/networkSecurityGroups/aks-agentpool-21615671-nsg"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tb07qlatru1cgnjdlo69/agentPools/workers1"
        },
        "ImageIID": {
          "NameId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2404gen2containerd/versions/202609.15.0",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2404gen2containerd/versions/202609.15.0"
        },
        "VMSpecName": "Standard_B4as_v2",
        "RootDiskType": "PremiumSSD",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Compute/sshPublicKeys/tbh9oq28l3i7cb1062kn"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Updating",
        "Nodes": [
          {
            "NameId": "aks-workers1-10476153-vmss_0",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-10476153-vmss/virtualMachines/0"
          },
          {
            "NameId": "aks-workers1-10476153-vmss_1",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-10476153-vmss/virtualMachines/1"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://dns-1791358885035685337-n9757uld.hcp.koreasouth.azmk8s.io:443",
      "Kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUU2RENDQXRDZ0F3SUJBZ0lRTHI1UVAwTmRxZW9wTUNoQ25SY2hlakFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWdGdzB5TmpFd01EY3dOek15TURKYUdBOHlNRFUyTVRBd056QTNOREl3TWxvdwpEVEVMTUFrR0ExVUVBeE1DWTJFd2dnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUURwCnJVc2pVaUJTQ2dsOFVhSHZ2S0xyYTVhNFBGdlNvTk1OcGU3dllXQTVra0RQd2pIcCtkR2QvRjIvazd6V3A0YWQKZSsxazFTd1ZaK0tKNEdCMGpITzZiMXlXZkIyd1JvQW9oTStMT3NyRHRZbUNSVGpQVjdudUd0ME9xMVRzaFY4bApkZFpWeWxJMndtc3pYMTNHMzErQkZYUzB5QTc3R2ozNzlNRUdFM3g2UG5pdENxd1V0V0Z3b0s5VmVDMURickVsCmQyYnFnRU9uNXpDUDU2M05XU05GMFdFN3pUTkovYTdTQXVXNEdjV3ZSdnNDK2FNYjZBd0RKM0NqQjdMTkVHdTkKTGlqWG85Q2pIZm5GY2lyalRpQWZOQWlzUXR5UisvekpKM21RYmZvWkx2MHJtcm1PcmgwbW9Ddm5UclhwQTE3TgpzS3Z1SlZRWmNaaXcwR3NPcklFZEJ2STZiUCtsZ1FUVi9jUzUzUkFUcTR1YjFWcFQxcHg2Y3JWOHRTcHAyTk9tCmpIaTBaK2RSdU01S3A1c0k1T2RxRjRkWFhxaDNEb0diektiSFNCU3NwQ3dIeVRDdEY2RlBQLzdkdVFoYXE5Y0kKempLbWNqVGFhaXpaLzExTVBKZjRlUFkrQW82NnY1VzFSRWZ5OGZlREtGc1BsQzJteWxLVEtEc0VBamFrbVV6ZgpuTkRVSDY0VitzR2NuOXZsS0RSRFdrdGNZRkVDNkhCay9PZVk5aGpCSzI0RFRQVXZoU3hUbnhFOFhjd3k3aTZNCnNCVmY4ZlZXUmk5UkZROU11amxMWGFzcWY2VFdCdHFiU2dWS1orZTlvS29haHJPL0xrRHFwVFNWeGJ2SS8rY3EKK0o5Y1pGNUY0dnp6Wk12ZHI4Wmc0dHFVZk1aajRyWHFvM2xPT3liYUN3SURBUUFCbzBJd1FEQU9CZ05WSFE4QgpBZjhFQkFNQ0FxUXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVha0NpTklDelBkcVlJTlRICm9kajJ4aE9QSG13d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dJQkFFVXorNDY0L0xpcnBlckRkcDFQZGhXQS9HNmUKSXJEeE43N0hVNkE2N1FiMmE4NzlkVnJ1TzB3MG0xbS9vNHhhL05manhZZWlqT3ZKUUgxVTlpQlJZK1I3M1R0WQo3ZFRMdUNrS0kyRzhhVW1MZUVGelZXMzJJT3ZJUVdxamM4V2swMHREdjRYVWZndHp4aXBYWEdCRXJmbnlZM1piClBCTTN4YnlUVkRCbmRvQ2pxOVRvYkVBSHlHczJMYXl5TnhHYTg2UnFoTjlhcVFGQ1k5a1IvK050Z3FZbnVtbHkKeGxkTXg5ejAyUjltem5sN1FSZlM0emNnNFpIcFMrUnhiSjV2LzlGNElaZmk1djE5OHFQTWRPVjA4OU5qVXc4VwpnM09nWTFXbFl3NWVVNVpBSGt2QkdVNEUzZnNwaGRoWnlGNXYyUmMxWU5ub3pZdEpHRUNaRnBwdU94T0JESUQ4CmtKRG1zaVB3L1c5SlIwZWM5WDJ0SVhuS3d6YTJrak5YYjF2eE9QampoNWR4TjRidjZ5Z3NFTFlzSTZnb1NtTzUKVjFPOWRQVWQxdjdCbUVieVhEaERWT0pmL2Jia0YvTDJsbDNDa1NOOGE2SWxJQWs4ellta2xvWGIzbnpvS1pYRwpBd1VFOCt5YloxdHBUTisrRXp5a3NKMFRUSG9MaTRwTUtSU2JYYWRKK2dON3ZqSHNYZ0VwNXNEd2VLdG1YVnpCCjd0QVNMWlZHeFBvdW1mLzdDMnlPNUVuRnFabFRKSzc5UXUzQjJ3VERsYTMxSURjT0MwSEpXZWNNbHJFZlUrRUoKMDRkaHZnbmVtczdUOEkvbHEzM1RBSFBOemhXbFkreFQ2cEZaTUlkaU11NDVTeHFVUDRLNjdva0p1UlZ0OXhoSApsaTNTYkJjdmswbzhZalBxCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://dns-1791358885035685337-n9757uld.hcp.koreasouth.azmk8s.io:443\n  name: tb07qlatru1cgnjdlo69\ncontexts:\n- context:\n    cluster: tb07qlatru1cgnjdlo69\n    user: clusterAdmin_koreasouth_tb07qlatru1cgnjdlo69\n  name: tb07qlatru1cgnjdlo69\ncurrent-context: tb07qlatru1cgnjdlo69\nkind: Config\nusers:\n- name: clusterAdmin_koreasouth_tb07qlatru1cgnjdlo69\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZIVENDQXdXZ0F3SUJBZ0lRUW54NTZXeEFTalVKOGVsVzAyUDhJakFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWVGdzB5TmpFd01EY3dOek15TURKYUZ3MHlPREV3TURjd056UXlNREphTURBeApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1SVXdFd1lEVlFRREV3eHRZWE4wWlhKamJHbGxiblF3CmdnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUUM2dlRPWnc4UklJS3ErWVpmRy81RlEKWENicnlJaVJTUDFEUHB3dmdDR1U1T3BIS05aem1FZDRqQ2QxR3dGL2gxZHozR3A3UWtlRXB6Qk1hdlJ0TEtyeApvanRaTzF0WHdFR2tRWDN2anJsRGlMUUNIUGwvRGwvem96Wm95c2VPKzlWUVB0VGUyYnZHUWVncmEzZmJLNHNtCkJWbG1lOWIrSnVjczJyYTdsZTIwVS9HUFgvLy9UaytFd1EzcGRXU09XVkFaV3dMZTUwd0xOaTdmYkFhWHN3M1oKbGpoTXZHSXdlaUtNTXdnWGFzS3lHbkJkYStIdk1yaCtEOHExVit1S2grandURW5TamNTSERQVHk1L3Faa3pGWQozbXNDaXRXWTFxQmZNbmFGc3hhVjRYZnpkY1d5dGdObjhrcSswbTkveHpyWDQ4S3ZmT0MxdjlGS2xhUkpsMGJtClhPWXMvN2FOeHVIQ2hMTWFvem9FamovSkduR1hpYjFVL2ZOaExSWUo0Rk1LQjliNTlpdm83Tm1NaGcwdmhWZm8KWEg3T0RqUzlKQmVGYnZMQ1ZZUEFKRTZERUpBNEw3VUVKajFtaEJFUm5GTXpPMEVUTFdvVGI3d0JQUVlMSjZLWApqQzdBUTd4S2VlSm1XWDZwYjB4R2Z3blRTUEk4NEtzbFRieFdmNTJyTUxIcVE5SzZ0UEduR25wVkgvN2VITGdSCjdua05qa3Q0YjJ2allqK3RlYW55TitrU0tnNFhYU2laaGgvaWhIbFJrZ1E2NHBCS0NYYTNYMmN5N0hPRHdMaVkKVHIxK1oxcHVMcUpzUkplV0lVYThhNy9GY1ZGY1Rkc3RwSHJmQWR4Zzhub0tCQnlFVmpzR3FRVjZFWHNqZzFNQwpkUkJtTUxSUEROL1lKalpma3ZEdWhRSURBUUFCbzFZd1ZEQU9CZ05WSFE4QkFmOEVCQU1DQmFBd0V3WURWUjBsCkJBd3dDZ1lJS3dZQkJRVUhBd0l3REFZRFZSMFRBUUgvQkFJd0FEQWZCZ05WSFNNRUdEQVdnQlJxUUtJMGdMTTkKMnBnZzFNZWgyUGJHRTQ4ZWJEQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FnRUFsOVMrdStDSytYczY1TlBabmlaOApxLy9rRzdzRkxIa0daMXFoaWRQNmtoVE41R2RBUGswZzNXbWdsZ2d3UUtNSHhIYjdSRmRwbDE0SzcvK2QrZmlyClAvd2N0N1NxaE9SNmR4TEl4alkwbXZQcCs5dStJcVFVcHlpOVF3bWt6TGg2TkM4cFZQRVhteWpURHVTMHh2Q1UKVitiQ2dhU0N1VUZWY2VVUHJ4azg0QVN3LzRZN2hGTDhOanlrTUJ6OXNHUlR3WTllU29LdVlDekNRQmpubHYrVwo2UEx5d1ViQmhzZXNaU0FFUzN4RnJtNjk0UlZQQVowR1U1SE5kQVI3Rk4ybzVyc1V2V0RSckZRZG56NzRNQmlQCmw2YjMvZGRHVkIxcnBIZVJTN2tJbG0wR0ZqR3lqSjVEdjQzdnpScmV6c0R3MVRJa1hmVk4yOFJ4MlpFeG8zekgKZkZET05GWTBncGNyb3FNanEvMlo2bG9WME5YSGMxcjZDem8wRzdpbEpnOGRyUERlcjRVZTdwWExqRUpUNEorRApxY1Zsb1ZoVUhLaDdBd0hrNktIbitRd056WVNjTUNRMisvS214b25kNWJMMDBqSFprRGk1SGp4dVBIVHJuRnFOCm0vWHNlOGNlQ2lrUmtsSmVCM3NGdXNGeXZzTHRPeE5pRlBYQ3NjTmNSTjVOdG9KUzBiNlBFeE5hY01XcjBMaC8KVGZIaWducnNQV0hyeU9RUU1PVmhMak5FbWtwb05yN2IyOWpMSVg0VVRRS3pOeTFKQlpsT0FjMlBhTDhyZ0h1dAp5cC9rd0dWZGY0a21rckQvcGprcFpwRlFQeENJUDZjN3VHb2ZZaWl6ZkFzbk84TlpKaE9GODRzNTZSREVpSVRCCkZVQzIxV0lUcFpDUU56eGlPbDdnNHJ3PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS0FJQkFBS0NBZ0VBdXIwem1jUEVTQ0Nxdm1HWHh2K1JVRndtNjhpSWtVajlRejZjTDRBaGxPVHFSeWpXCmM1aEhlSXduZFJzQmY0ZFhjOXhxZTBKSGhLY3dUR3IwYlN5cThhSTdXVHRiVjhCQnBFRjk3NDY1UTRpMEFoejUKZnc1Zjg2TTJhTXJIanZ2VlVEN1UzdG03eGtIb0sydDMyeXVMSmdWWlpudlcvaWJuTE5xMnU1WHR0RlB4ajEvLwovMDVQaE1FTjZYVmtqbGxRR1ZzQzN1ZE1Dell1MzJ3R2w3TU4yWlk0VEx4aU1Ib2lqRE1JRjJyQ3NocHdYV3ZoCjd6SzRmZy9LdFZmcmlvZm84RXhKMG8zRWh3ejA4dWY2bVpNeFdONXJBb3JWbU5hZ1h6SjJoYk1XbGVGMzgzWEYKc3JZRFovSkt2dEp2ZjhjNjErUENyM3pndGIvUlNwV2tTWmRHNWx6bUxQKzJqY2Jod29TekdxTTZCSTQveVJweApsNG05VlAzellTMFdDZUJUQ2dmVytmWXI2T3paaklZTkw0Vlg2Rngremc0MHZTUVhoVzd5d2xXRHdDUk9neENRCk9DKzFCQ1k5Wm9RUkVaeFRNenRCRXkxcUUyKzhBVDBHQ3llaWw0d3V3RU84U25uaVpsbCtxVzlNUm44SjAwankKUE9DckpVMjhWbitkcXpDeDZrUFN1clR4cHhwNlZSLyszaHk0RWU1NURZNUxlRzlyNDJJL3JYbXA4amZwRWlvTwpGMTBvbVlZZjRvUjVVWklFT3VLUVNnbDJ0MTluTXV4emc4QzRtRTY5Zm1kYWJpNmliRVNYbGlGR3ZHdS94WEZSClhFM2JMYVI2M3dIY1lQSjZDZ1FjaEZZN0Jxa0ZlaEY3STROVEFuVVFaakMwVHd6ZjJDWTJYNUx3N29VQ0F3RUEKQVFLQ0FnQXpBYTBQaWc2YUdQb3FhR1Bad2tNQjdUbmdXM1VhSWhLemhHMks5L2UwUU5iUk94RmduNjZkK2NKcApWVWdTZW5oanVFZ0poUEFlQnNERmpzMVc2TVdFbk9pVEZnaDhMcEovZURnYThDUWdrejEvK2tRWDlXZzJGMVdzClIvODhTZ05aR09CeVFvenN1V0hlUWt1SnZSb1Q5NUFVMHl6RlhRRDJHb09oTjNHbk9PcVpYQUtEelBNaGNzNEcKLzc3SXR4S082bklkU3BaVjFhWDE5djdSL3VHQWxteW9ZU2g4ejArU2RmbjIwMzE4d29ZQXJwQVhFd2Fid0ZxQgpNQk94VWpCSk9yMUxXdGZjZjZpRWxPSzVnaUFQZ3lrOWw0aUNaOVdFU3pkZWk2MVdvVEJ2ZXJMV2ZxcnRnSnYwCkJXelpQU1VRaFREUFdFL1V0U0ZJdXhrVjFORlRjUDlJNCs1dVJBU2o4REFPdUNIUHhWSU1LcjZTU3JRc1p6MmoKTGpCM3I3UE1tTHRHZE8zbUJwUXZyTWtrbE9qMDNiZXZ4L0ZTbTlJOTVyejVINWFjckZwTGNjRUtVSERvUlhqRQpERU0rTEloMUZXZGtDT2xpMWIzd2RhOHh6eFpsYlFhUjhFTmRqUExpMDErM005WENNb2VLM2xDak84VzJiSTBpCkdpa2NxemhqUjVRZGVPRTJRbG8wSllGWEVwSUxRQjZkdXZ0L0Rxd0FZajNrWTNKOWp5N2xmdE4vNk1VYjB5Wm4KZmY2OG1ud0thQ21ZWXkxeVl6Y3NzN2VIcmFVSXdnTGxoSk95K2s4VUc1aDg2UVBpUUk5MWthQ0JJUi9yR0dKQgp0OGUyR01jVlVubFJQYW9YemhXcjhUL2JMREQyM1RWVlFybXFEVEJ3WlZoempCbXNQUUtDQVFFQXh0eUNQNUU1ClJyNEt3M21HOC9TTWZibW9uR2xsTVhwejNFbnVqODhsMGk2OWlFbSszT29HM3dXZndZTW9yd3BLNXVPYU4xQ2MKTGtheHQrVUdFbWtvWlh3SUlrUkRWMmVaNERtZGI2UEV4dXhzdlR2T1pCZklsQnBVTXZZMlpKTnk1bEJuYTgyMAp3YlgvY2ZqVjVBL1dvS3o2TDk0c1BtSTIvTmdKN0VQRHV1RUcraDE5UW5sZ3pLeEVFc1FRYXV1RjlGWjZWdmwrCnpKKzNUUDd6NlZodE5nNE1mMDhIMGlLWStaS0l2MU9uUUpGbC9lTUkxamhERGkwRlVydVFXN0VSdTlHSWVsOG4KR1RTc0pNNDlmYkJJZnBzOFJ0d2tTRjByM0NuUmUyYmR5bWp6Zy9rSmk3TktFckJKTUNyY09MdHhsZmhGN1huRQpwbjlHL1RBNU10bVQxd0tDQVFFQThHVUZ5Vy9TRGtxbE9TcWE3cmQ5eE1vQmwxSFlmUWxVeHRuSy93cnhmdUlKCjRJNXQ3M05vV01McTRXTDY1U2R6QVRmQzZTTTNZY1p5MDYrUTdDVGZFSHNNSzN1MHdYMUZIVWtOTGg4WXdNVkUKU09aOHA5WkxDc0dVQy9nRFdhVDYrOEo1bzBXVWFkNnh1UTB0Qitsa0xrQklrMGxUOFRuNzA1Z0Y4SG0yTkFmMwp1VzloS0JFZTAwK3lhYUtQK1NUNHYwbnlrTmNvbktXM1ltYnhMTXprcEZKQXpoa2hwT1M2MUJ4SFcvSlV0YjdtClI3MHA2RVJQUkwrMDQyaGNXS1loaGZTdE5wbHYrKzN3QmNBWncrSE00c0FIQXloMGlLVVdqcTNWNFVjYmxTTVEKQ3ZvR2huZXkzQytNOXZCOE5EV1BxdStsMktVekkxZU1rcCsyMldjRkF3S0NBUUFkTVgvNWZ3Tk9qRldlTHJnTwpGa2VOcVhURVNZQ3VpOUI0ZFc0Tk5KYlR6VUxMd0MvQjFLVmZsajdHdE0rREt0cW1IS0dtSHpKZVpNaHdPN0h4CmYvaUxOWE1vUEtjNkxKNWRXWGZ3VExWMWtuM0FKZ2g5anpSRjhidU1vN0tHZjFMdDFyM29DSkhSb2pzMjZ4WFQKWGpQZThLWGw0eExSc2E5cTNQaFA4LzdHNzRRTHNjcVN1S2pxUXh2WE9XdERsd0hhUmR1OTZ2Q2ZiNDhFUWFWagpDelUxSzBqUVk3UzlONUVIaW1pQzFmQUd2WkdnZlBUdUplSktNWGZIbG50eHlvUFU4OFM2V2ordUpwcGk2TmdFCkhKMzM0Q2d3S3Q4MHRHRGMrRktsY0F0OVRIejdVZE1CN0RjaW1UaVZWcVZ2dXF6SkhHSi9vMUdvTTZValJESVIKY2ZtUEFvSUJBUUN3aGI0bUdOSnFaQzNIT0oza1JTRTViQXlYRFNhblNqN1NkZ00ySE5jUWtQTW56ZTYwTWs2aApQeUF5dVBmN1pXaHhzSHlUcmFSeFM4UHQ1YkhKZFpuSEJGUC9haXRsR1pPeTJYMndMRzJFd0ZaUUljL1BmdkJECnlibm93QW5WdmV2L0N1TS9IVXpLSVJqb0JlRFhPbTZ4OThwVFBUbENuWTFwNXV3VmxZcFIwcmY1bHV3RSsrMWkKeGRjQThybzkvUVFia3pWblhsY1lFQ2dGUGdwREY0R2RtRGkySG9ZeUN6T1ZwZDVRaWpYaDczZ3huWm8wVUZUVQo2dUR0VENqamY3cUdIeEZDVXBHRXNVZmJNU0M2VWdpZFNOemRXTXVadTRCQTdTMXM2MFk1MFRGcW5nWkVuNGpiClhwOW9oVmJ2RDVXYnM3Wll5NVl0a1pCTDZyUGVHR1A1QW9JQkFEelNMUjF1K3I4WnErNFVhZkU0SVZGMmVTNG8KU2RYbkNHRmpkWVp3WHlBeWZKc09LaTVDQlNub2NTKytuUCtVbXBoMFZNWExkVzhJZGRucXhnQlQxWVlubWxEcgpwMmRVTVlxdlpnb2FTdEQ2SWIzQm4wd3VuNFdnU2RrOE8reE5admJXa3gwRGgxdU1lbWlsYWNraGNVbGpadmFNCmpXYTc1Z25zbW1BNk44ZkJSY3VJeXJvOXl2S3B6a0xrVEtSdmlEaG1MUGJveGhJMnZUV3FqbDJLL2FWeEY3T3IKK3hYZytQQ1NWeHNNV2xvQ3E5SzkzZFNhOGtxRHc1V0o2VUdqMnpSMWl0am4vSzNsb1R2dVkxb2VLUzlGcHFPNApOa3dvSE9mU2lyejBBbHd4c1VIV2xhMktiNU9KdURSb1hkNjZ0NlpINktyRTBjd1ZvR3lpWENBenZ3VT0KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K\n    token: f2orrx81vh474tzdz0a5da1mvet8u6i6ui2vgzavwmmq87hzfwbjqlimmd8ilaq9dw6bhicdqimec9zj8hvsq2qzj2statwv6i7vy53t84qkkz2g4vrbnc8xt4vwh0ok\n"
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-10-07T07:41:17Z",
    "KeyValueList": [
      {
        "key": "Location",
        "value": "koreasouth"
      },
      {
        "key": "Identity",
        "value": "{principalId:e99c6063-3753-47f3-80ea-e33b557c5965,tenantId:fb98dda1-32ff-48eb-a489-62777cd9ccd8,type:SystemAssigned}"
      },
      {
        "key": "Kind",
        "value": "Base"
      },
      {
        "key": "Properties",
        "value": "{agentPoolProfiles:[{count:2,currentOrchestratorVersion:1.35.8,eTag:b7a178a9-f728-4f33-8f4f-fd15ed9a6704,enableAutoScaling:false,enableFIPS:false,enableNodePublicIP:true,kubeletDiskType:OS,maxPods:110,mode:System,name:workers1,nodeImageVersion:AKSUbuntu-2404gen2containerd-202609.15.0,orchestratorVersion:1.35.8,osDiskSizeGB:100,osDiskType:Managed,osSKU:Ubuntu,osType:Linux,powerState:{code:Running},provisioningState:Succeeded,scaleDownMode:Delete,securityProfile:{enableSecureBoot:false,enableVTPM:false,sshAccess:LocalUser},type:VirtualMachineScaleSets,upgradeSettings:{maxSurge:10%,maxUnavailable:0},vmSize:Standard_B4as_v2,vnetSubnetID:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbjvs627fsvq9c8qjdmm/subnets/tbi5s3q7etm8cahfjubo}],autoUpgradeProfile:{nodeOSUpgradeChannel:NodeImage},azurePortalFQDN:dns-1791358885035685337-n9757uld.portal.hcp.koreasouth.azmk8s.io,bootstrapProfile:{artifactSource:Direct},currentKubernetesVersion:1.35.8,dnsPrefix:dns-1791358885035685337,enableRBAC:true,fqdn:dns-1791358885035685337-n9757uld.hcp.koreasouth.azmk8s.io,identityProfile:{kubeletidentity:{clientId:aa26e04b-b027-4a58-9065-6cc825160177,objectId:e2ef4ebc-4130-4afb-8cda-cc9a3825c0ee,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/tb07qlatru1cgnjdlo69-agentpool}},ingressProfile:{webAppRouting:{dnsZoneResourceIds:[/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/dnszones/tb07qlatru1cgnjdlo69.com],enabled:true,gatewayAPIImplementations:{appRoutingIstio:{mode:Disabled}},identity:{clientId:7768eff2-560d-4c61-a3bb-4740fccf8536,objectId:845c8887-a41b-4b29-b4a3-c77e8020237f,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/webapprouting-tb07qlatru1cgnjdlo69},nginx:{defaultIngressControllerType:AnnotationControlled}}},kubernetesVersion:1.35.8,linuxProfile:{adminUsername:cb-user,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDA7EqTWJv+hdZCkyjquj4TaYf5mYZQ4lsvvjLwuWaFV6BKS0Owt/B+rMl4PYS6zuK2hBjBFFYVY/hmPyhUiSzyX93wWSrHk2axbJyL++kJSVQ0dmO8Tcnr47K8Tp6WI+Ex72s5vIeKnltDbjLIsYHcisuB0bPNfDREFO3aUkKjub3l40Nowq2SmlOV4pB1MpUigKUPrh4QVrRhF0vdcXSzZ1wYCfWiPzF/fU0d+7f7raTwc8bDgHRSYsOQ7r4Je8XV2HPyG4kMU1z55l1RZkWE4Nhn2Ojy+H2nq/PGt+LB9M10U7f1iykdIG7N40SoyGnTekyPDZ9/o8E8yOjaMMAhBLFy9xs1VTARi2ZhJUKF/8MQRV/hZYMEZclYH8CGCiFvDBsU0xCgKjtsso8WO2zoGcmdZTDQb4+gy/wvvHRFGQPnTY79MfKPLtUu3B7c0uDVGtS5reeDlDsU2Is2e7pFX29sMZ3r9C8w61raTlUUagTLNbiHkTPiXgU4MFHDIeDpqKP5dBF6FW/qq4whuwQJSpjVg7H/m529RCo+5WgtnoZNEbvVZbXi5R69AEemoEvanua+ThwtytrE2yOsu7hVo0kvMjezxgodoyZLxwxYi8trvvNJxen3Ck2c6r/AzVx0slcjguz3uNi0RQf0hQQ30gva4VzDVjZwLWQx5Ngz6Q==\\n}]}},maxAgentPools:100,metricsProfile:{costAnalysis:{enabled:false}},networkProfile:{dnsServiceIP:10.1.0.10,ipFamilies:[IPv4],loadBalancerProfile:{backendPoolType:nodeIPConfiguration,effectiveOutboundIPs:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth/providers/Microsoft.Network/publicIPAddresses/2f1608e8-04bb-43b6-a4cf-d880526113d3}],managedOutboundIPs:{count:1}},loadBalancerSku:standard,networkDataplane:azure,networkPlugin:azure,networkPolicy:azure,outboundType:loadBalancer,serviceCidr:10.1.0.0/16,serviceCidrs:[10.1.0.0/16]},nodeProvisioningProfile:{mode:Manual},nodeResourceGroup:CB_koreasouth_tb07qlatru1cgnjdlo69_koreasouth,oidcIssuerProfile:{enabled:true,issuerURL:https://koreasouth.oic.prod-aks.azure.com/fb98dda1-32ff-48eb-a489-62777cd9ccd8/7410c9fc-0ac2-4592-89ff-28f5bb32a74c/},powerState:{code:Running},provisioningState:Succeeded,resourceUID:6ac5f7a8eb7c6b00013d1377,securityProfile:{},servicePrincipalProfile:{clientId:msi},storageProfile:{diskCSIDriver:{enabled:true},fileCSIDriver:{enabled:true},snapshotController:{enabled:true}},supportPlan:KubernetesOfficial,windowsProfile:{adminUsername:azureuser,enableCSIProxy:true},workloadAutoScalerProfile:{}}"
      },
      {
        "key": "SKU",
        "value": "{name:Base,tier:Standard}"
      },
      {
        "key": "Tags",
        "value": "{createdAt:1791358877,ownerCluster:tb07qlatru1cgnjdlo69,sshkey:tbh9oq28l3i7cb1062kn}"
      },
      {
        "key": "ETag",
        "value": "be9b0b92-0c64-471d-8cb8-d17b110ccb9e"
      },
      {
        "key": "ID",
        "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tb07qlatru1cgnjdlo69"
      },
      {
        "key": "Name",
        "value": "tb07qlatru1cgnjdlo69"
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

