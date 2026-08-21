# CM-Beetle K8s Infra Migration Test Results — Azure-Busan

> [!NOTE]
> Full lifecycle against a real CSP: recommend → migrate → list → get (verified against
> the recommendation) → delete → residual resource check.

## Environment

- CSP / Region: azure / koreasouth
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.0
- Git Commit: 803afb4
- Namespace: mig01
- Test Date: 2026-08-20 15:15:45 KST
- Cluster ID: mig02-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 19ms |
| 2 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 5m0.055s |
| 3 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 2ms |
| 4 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 5.781s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 38.004s |
| 6 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 5m32.187s |
| 7 | Residual resource check (Tumblebug) | ✅ **PASS** | 3ms |

**Overall Result**: 7/7 steps passed ✅

**Total Duration**: 11m16s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 19ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.34.8)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=azure+koreasouth+standard_b4as_v2 nodes=2

### Step 2 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 5m0.055s
- **Status Code**: 202

- ℹ️  nameSeed: mig02
- ℹ️  async reqId: 1787206545868968035
- ℹ️  cluster id: mig02-on-prem-k8s-cluster
- ℹ️  elapsed: 5m0s
- ✅ status: Active

### Step 3 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 2ms
- **Status Code**: 200

- ✅ migrated cluster present in list (3 total)

### Step 4 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 5.781s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=azure+koreasouth+standard_b4as_v2, nodes=2)
- ✅ version: 1.34.8 (recommended 1.34.8)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 38.004s

- ✅ kubeconfig obtained (server: https://dns-1787206565697750064-yka2fsje.hcp.koreasouth.azmk8s.io:443)
- ℹ️  auth method: static token in kubeconfig
- ✅ API server reachable (v1.34.8)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 1)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: 52.147.121.158
- ✅ nginx served over the LoadBalancer at http://52.147.121.158/ (attempt 1)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 6 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 5m32.187s
- **Status Code**: 200

- ✅ deleted on attempt 1 (5m32s)

### Step 7 — Residual resource check (Tumblebug)

- **Duration**: 3ms

- ℹ️  VNet mig02-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig02-k8s-sg still exists (known gap)
- ℹ️  SshKey mig02-k8s-sshkey still exists (known gap)

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
  "id": "mig02-on-prem-k8s-cluster",
  "uid": "tbkpeno23hnamhgbrptu",
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
    "createdAt": "1787206557",
    "ownerCluster": "tbkpeno23hnamhgbrptu",
    "sshkey": "tbh9oq28l3i7cb1062kn",
    "sys.connectionName": "azure-koreasouth",
    "sys.createdTime": "2026-08-20 06:15:57 +0000 UTC",
    "sys.cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tbkpeno23hnamhgbrptu",
    "sys.cspResourceName": "tbkpeno23hnamhgbrptu",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig02-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig02-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbkpeno23hnamhgbrptu",
    "sys.version": "1.34.8"
  },
  "systemLabel": "",
  "version": "1.34.8",
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
          "cspResourceName": "aks-workers1-34776577-vmss_0",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-34776577-vmss/virtualMachines/0"
        },
        {
          "cspResourceName": "aks-workers1-34776577-vmss_1",
          "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-34776577-vmss/virtualMachines/1"
        }
      ],
      "keyValueList": null,
      "cspResourceName": "workers1",
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tbkpeno23hnamhgbrptu/agentPools/workers1",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tbkpeno23hnamhgbrptu/agentPools/workers1"
        },
        "ImageIID": {
          "NameId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd/versions/202608.06.1",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd/versions/202608.06.1"
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
            "NameId": "aks-workers1-34776577-vmss_0",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-34776577-vmss/virtualMachines/0"
          },
          {
            "NameId": "aks-workers1-34776577-vmss_1",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-34776577-vmss/virtualMachines/1"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://dns-1787206565697750064-yka2fsje.hcp.koreasouth.azmk8s.io:443",
    "kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUU2RENDQXRDZ0F3SUJBZ0lRZDB5Tm1qRmJwcEZ4TEFMcEV1cVprekFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWdGdzB5TmpBNE1qQXdOakEyTkRsYUdBOHlNRFUyTURneU1EQTJNVFkwT1ZvdwpEVEVMTUFrR0ExVUVBeE1DWTJFd2dnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUUMzCmZtMGR6YUJvdnVWb3FoWlJTUlcxS0x1RHM4ZFk5M05qUXdHaGdhNUV6R1YzZ0UwcHptbHZMdlRqRVh2aFhjaEQKUDYxV2VPb1d6N3laSFpraUsyaEZGV09PbmV5NGdTMDdUTDQ3bTFPT1pzTzF4R09sN3JoM005VCt0TGdmVjhobApMblJOMWJmMEliM3N3cHhRd1MxVlh4K203Z2MxaEZkaFh6QW11NWwzWkhWY0pCVWlkVjZKdjFMU3lMc2NSL3V5CnFVcTQxYXFITllXQ0xKZFc5OFg1K1BBdUtteWRkM0w0R0w4SXhKaDZORXdvWmloYXhFTkhaRnlnK0JJMC9OTXEKU0psb0s1U1RFUWdPendiT2h5eFEzRUMyZnhZcG5VUm5rR05lUW1VckdWUk92UmphSEJrcUpid2ptOGo1ZVhmTQpSVGpnMG5rVTYwL1NBejRUazdFdjErakJhN3B1emdCMm8vQ0lnY3lpV3Y4aldvYXVsNnNUcEhLQit3eEFZaUdVCmtPR0Fsbldmc3RyTGRDTmFOMUs2Y2p3R24wajhvaXhxUXNkRjNPTWRqT3MyVHM3eUJRV2VZbFNhblNjYk5uYisKenZpeFZ5TE1DT1J4dWdYNi8yS0p3ZC8vUEhnVUJHS20xQzlXZDdHYm9uOW1MK21ObGZzb004dURYbmdoRnlhWgp2ZVVGMEpNZjNHTWJCNi9Cc29oK1V0bUlTb3Y0ejI0eGl4d25ub1Z2OTRoQjFjOFA5UTZ3MDNKRE91Wms3QjRBCjJ2NzNWYWRYRzYxZjluTXFkUlFGcXUyU2RmRDdmVVhPeG9lVnYzeUxmbllyOFdCTmJKNTdYdE9lVmVLS3g0Z1oKQU95MzJEb2d4ZFYwam9lMTZmWVdNa2NFNEt3NXlUQTdwMTF0cSsxSXJ3SURBUUFCbzBJd1FEQU9CZ05WSFE4QgpBZjhFQkFNQ0FxUXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVqYTRuTEJzVzIvUTk3NU5qCkhmQnArUkttVkRzd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dJQkFBRGpuZWVreWQ4SVlranlWSWc0ZXpkSmVHemEKclo1RlJteDFNZVlsVW02c3ViN3ZrRmFYM3pPSHZ0anZBeEp1TytrNjFVVmw4eUUvQXVVSG11c2NlaFY4WlpQRAozRjcvSU4zeUxkSUx6L0lsdkJrS09xUk5sWW5wZmYvRmxFWS9uY2Jkajg4UytwVzJaMDY4U2o0VnkvTGJrczFzCmMwTis5RVcxdjlMekZEOUtaNG16ckVKRXV2all0YUZOOUo5cTdJNGtnRVJEcEdwa1ZCSlQvNUN0NGlCSG50V3MKWVAwU2J2QmgyQ0hJTVZSVEVoQ2l6ZmN3UmlWekpCYTNVNXRlTDk5T1NaRDJ0Q0paSUlQREU0bjBXcWsrK3JSdgowU25QUzdiTHUrWFQ0eExVb3pyUlRsUVN0bzB6RjJHMXRpNlUyVTNjU3Nod2xYSHpXbkpxd0ZBZDFGcHdncU9uCnFDU25TRGx2dnlKRmF6WnBBUi8yM01iRGlUQ2tLckpoM3BhOUpieHVnaUVkcXJVWTd2SGtPTFdsaXNuLzJHMTYKcExXRk5DYVBveWxWc2QvNTQxQWpUS3ExSDduK0llM09RTTJ6TmFZbzZTMnc5SnlWNE1mMm5CNUsyNFN1UXNOegpTTnJpRVlrNGE4UW9OV0xxakluNkpLTGZ0c2ZNZTIxTCtpVDFKSnJEQVMyZVlOU28yOUJPd2tQSGx0VFpNUmtoCnBwYU5CZy82Y1l6MkkwZjlOYU02WHZ5Y2ovcGhldzdTbDllVkp4OW84N205cWo4ZjREYXBSRlp1ekwxdWVPc24KVTNJL0tTUUZxTkV1TXlFanY1Z0IyMng4RFh6OVRLSWNTSDJiTWhkemd3eG4wcW9PdTdlV0FURUt0NEFMQ01KMgpJeVVVLzlnb0FyZnR1aVNtCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://dns-1787206565697750064-yka2fsje.hcp.koreasouth.azmk8s.io:443\n  name: tbkpeno23hnamhgbrptu\ncontexts:\n- context:\n    cluster: tbkpeno23hnamhgbrptu\n    user: clusterAdmin_koreasouth_tbkpeno23hnamhgbrptu\n  name: tbkpeno23hnamhgbrptu\ncurrent-context: tbkpeno23hnamhgbrptu\nkind: Config\nusers:\n- name: clusterAdmin_koreasouth_tbkpeno23hnamhgbrptu\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZIakNDQXdhZ0F3SUJBZ0lSQU1OeXpjWWs4dVBWWWxQSGZiUXkzSUV3RFFZSktvWklodmNOQVFFTEJRQXcKRFRFTE1Ba0dBMVVFQXhNQ1kyRXdIaGNOTWpZd09ESXdNRFl3TmpVd1doY05Namd3T0RJd01EWXhOalV3V2pBdwpNUmN3RlFZRFZRUUtFdzV6ZVhOMFpXMDZiV0Z6ZEdWeWN6RVZNQk1HQTFVRUF4TU1iV0Z6ZEdWeVkyeHBaVzUwCk1JSUNJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBZzhBTUlJQ0NnS0NBZ0VBMk9NaGRNM2loTEJFdGx6bzNkbC8KWHNHUEd4UEJLRWg1VDZxcEhoVXZCV1praVdTK1FoRm9KOHpzU3JRTmgvZjR2a3JnK0dnQURKTGNVcEhmQXl5RApBRjdCOG93bDN6TEVyTkt1R2h3ajBDUUFRYkZ5czl0YjVqRVE1dm1QRDd3MkZpWE0xUWJySnBmNVE4UmluT3A0ClpIVXpiRGR0bnZLdEpobmJiM1doSHBDUm9iSEdZZVNxeW0zaWZ1SGRseUIwbTVlcHptcU41N0pxbWFXT2cyRUEKcEVZY2xxNnJnV2sycXhoNjBicFRnOXV2NjFKZ0ZGQStVcTQ1THcvdDdmM1VLRFpxVnBVMDlIdGdURnZZNGVRZQorZFdnOFNLMEE0M082YU1jM1E3Ynhkcmw4aEZRSkRML1RWZDdwaS9VMmk5MUdYMHhMYW9lNkpyQmhERnZad3hJCm1xTFoyVDZPbElTV2NaN3ByY3VqTUUrbE1IQXBIVGRIWkpoRVFnRjc1RVlxMXUwRkRZQUVDUll2OGlKUUtVQ0gKTkxWZVlGNlRwOW92YnlGUWdpM3hJcXNiVWpaNGgzdmU0ei9Nc2JYVWVYYTh4WUlQakkremJpb0JrWjV3T1BpMwpQaE5jLzVpRXB6dEkzOUdLeERGUldYNXlHb0ZYd0VXbloxWVRlL2JJZ21vMHRuK1lmaldCUWRXVFhGYmN4bFllCnpCY0NoMFRIN1llUHd1MWZqYUZQK05IanBaZHNQSERYbkVPU2ZFSHN4cUhDc3NtbzNQWWRUaFNzdEJ1SmlnZDMKZlJKK3VMUDcvT3o0ZkRKUjQ1ek9UTHhncnBZQlZVeUlJZmwrM2xUcEpaek5KR3dtRFZLQkFlZ1FkY3BkMElxUgpFblh2dmdnZzdIckIvSmFmSnhISVZDTUNBd0VBQWFOV01GUXdEZ1lEVlIwUEFRSC9CQVFEQWdXZ01CTUdBMVVkCkpRUU1NQW9HQ0NzR0FRVUZCd01DTUF3R0ExVWRFd0VCL3dRQ01BQXdId1lEVlIwakJCZ3dGb0FVamE0bkxCc1cKMi9ROTc1TmpIZkJwK1JLbVZEc3dEUVlKS29aSWh2Y05BUUVMQlFBRGdnSUJBSFh4T2s5WHI3K2U3T01TclhMTwptZFNVQWdPaFpDTXFSL014b1FiWHkvOUYwZHBsSkZINjlJVzh6TEo0ZDhWdnphUGNNZTQ4YmdHUDdIbko0SVJRCktZYjk2Y3cvZENuUjdRQ0xkM2VPNEd2dG1xV2dWQmIwdzN3WDYzSnFHL1ZQUkdqd0d4aFEvQmJ2clJpa1BESTMKeEpDdWRLNmdYSmJYN1NFUkFIemhiTnR0dnRwYkFFSWdjWVVLSjVaaDdvZTFYZHlqSzhJYmtjVkg1aXlPb2dTbApjaktBUExDbWVSVWVsdlZMWFVHVC9LaDk0ZGVIMzJkdGF0aUdoY0N3d1pYK2pvVXlEY2RSdWdDUHJsRVE1VzNzClg0bU90aUNNa2Y3MzB4ck9CaGNtWWNPM2tEWkNsWDdsYml3V2hpV20rTVZ2ODdXWnZrcS9QcWVzaVFVOEkzc3AKMTJJSUw4N0h2UkhUZEw3RGJTN2JCY2FlTksySjQvSitNaStpVkRoOXNpY2VzMnd1Znd0RWdOZlY5NDlPbklweApQQlkwUC94c1lZN3Naa2JaLzNKRzBqQWp4TDRjaU53NmlSSmhqVEQ2cFQ3K0JXNUlKQzAxbkU0L0tJaXhOclkyCkduL0pyVFJRWVRzRTNUcWFyeHVlMTJRQ3phZEN4M1pKV0R4bUNzSWpTdTVKV1dKaHBtUDl5amowaEFnUS9KSkoKNWY2Y3QvV05mR2pQVysxYllmNjhpUDQraWVpK2FReVNEaEtiOGhhTW4zcXE2K0FFbzVRbXh6ZmhxVytROHQzNgo5UnRKNElSOThqVzJqemYwR2VVQko3d1hyYnZyRFJTNXJOTjdYOUNVRlI1UWJjcWNsV21XWjhPQnNEZEdzTGVSCmNnN05OTkwzVkpvSEZlelJMdVZZK3JwNQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS1FJQkFBS0NBZ0VBMk9NaGRNM2loTEJFdGx6bzNkbC9Yc0dQR3hQQktFaDVUNnFwSGhVdkJXWmtpV1MrClFoRm9KOHpzU3JRTmgvZjR2a3JnK0dnQURKTGNVcEhmQXl5REFGN0I4b3dsM3pMRXJOS3VHaHdqMENRQVFiRnkKczl0YjVqRVE1dm1QRDd3MkZpWE0xUWJySnBmNVE4UmluT3A0WkhVemJEZHRudkt0SmhuYmIzV2hIcENSb2JIRwpZZVNxeW0zaWZ1SGRseUIwbTVlcHptcU41N0pxbWFXT2cyRUFwRVljbHE2cmdXazJxeGg2MGJwVGc5dXY2MUpnCkZGQStVcTQ1THcvdDdmM1VLRFpxVnBVMDlIdGdURnZZNGVRZStkV2c4U0swQTQzTzZhTWMzUTdieGRybDhoRlEKSkRML1RWZDdwaS9VMmk5MUdYMHhMYW9lNkpyQmhERnZad3hJbXFMWjJUNk9sSVNXY1o3cHJjdWpNRStsTUhBcApIVGRIWkpoRVFnRjc1RVlxMXUwRkRZQUVDUll2OGlKUUtVQ0hOTFZlWUY2VHA5b3ZieUZRZ2kzeElxc2JValo0CmgzdmU0ei9Nc2JYVWVYYTh4WUlQakkremJpb0JrWjV3T1BpM1BoTmMvNWlFcHp0STM5R0t4REZSV1g1eUdvRlgKd0VXbloxWVRlL2JJZ21vMHRuK1lmaldCUWRXVFhGYmN4bFllekJjQ2gwVEg3WWVQd3UxZmphRlArTkhqcFpkcwpQSERYbkVPU2ZFSHN4cUhDc3NtbzNQWWRUaFNzdEJ1SmlnZDNmUkordUxQNy9PejRmREpSNDV6T1RMeGdycFlCClZVeUlJZmwrM2xUcEpaek5KR3dtRFZLQkFlZ1FkY3BkMElxUkVuWHZ2Z2dnN0hyQi9KYWZKeEhJVkNNQ0F3RUEKQVFLQ0FnQkI3RWV2Q1NWZ3ozTVRPd3BNNUY4aW5oS3hXRC9OenJtUXpYNjU5aFprdmNxeE9EM2NOdzVCaXJnSAp2TktnRVc4NTUraVptSUxyVDNoSVlLNDRlTDhZemJTRjFMTnVOREF6bDVYenVibm8rZ2haNzJXOTVWNzVpTkJxClpGQm5wLzJJbmRTMHEzV3VOV00raGVLemIxRkl0NWI1dlo5RVFOOEFSYnU5RlRQejVsMWRtSHVFSmMwRDJvS04Kcm5sOEJoRnJlWjNUYisvU0RSajV1cWltcGtWYnFUUG5XUkFvTmFLNFBxaVdOdHhMcCtyQXpEa0g4NXY5NVpiYwpCeXQ2dXp4UlBMajF1RVJ3UzAvcDVjRDJhREJDSC96YlRvRUkwNEdnNGtOVHJjQi9VeG14aWpHaHp4NXFrN3l4CnRyZ3IyV0R1Ym04VVFqRkM0a2NQdHpiMVMzYUZldHZlaks1ZkJubTIzcjcvNlpUenhsek9Vd0VBcmtZejVJZkoKYTE3Wkt0b2wyOWNrRmsvaGNyb0R5b0ZCUTlGTUZ5MGxYM1VmRkh3T1FvazVrVjNFQmhJK2NMbUR4YXZaN3cweQptdXlpWnZLYjlEQ0lEdlIxQ2hxTGVuVzBTd1M1dTV0WHhiYWFTVlQzV3h2VDZYSDFxYTc1alpKVEJjNG1kdzkvCjNGbjlQQ2RZUVdsNHZKWEJsdUtFSVJOYktlNWhPSnVwVmMrK0hJSnd5clg3cm1aS3lkS1RpeEFYZUlleitnNmIKWFVmZFZQdEJTMW5ZUXhTaVg1RFhTUjdBNUJ0TE02VUR4K0xUVlBXNjIvYXJKcUNqcjlSM2F4eHhHUnVOVTJvcApyMGp6UDJtdElKMVVVcUlDQ3dGTTFjeEFrUHBOUTMzSzBneVJmeU1INTBCMkVDaEhnUUtDQVFFQTYxQjhJalJFClJqOWsyQVA0TEh6VGN6WFRtdGVISGFQRXdraXJxMUdPTXljc0xyWDZUNWpMU05JK1p4bWprdnNFZ1o5SFFnQUIKYVAxSDNsSndHZWdONmg2UnkyN1hwUlFuTlJTK1o5VTlET3VzNEtRMDhMc2pOOUErRWlMYXlnVkt0V21BMExuWAozc05VQ2Y0dkZvNGhyRHdtUDlnMkhYM0FpNXN1TUswaDBhMytpUEYxNXdsWUFGeUxCMzZKRmdxK2ZyUjA3bGdYClVHak1YN1k4d0JubUFVaFBQSVFENDZlK1gwNEpuWE12VC9tZEw2Z1I2MDhBQmt4RmRMVkprcmM4SWJlN1h2dUYKckQzMnBpMmNUZm5mM0prdm5YclZhUlYwendjbGJyaGRsK0JVZ0ZwUE9SbmI5K3NEOTlzUGE1UHYzNEZOV2NHVAp1RjVCWnBBZXdEWDhad0tDQVFFQTYvUDErK2RzOStlaSttN2dmSXJ5TkFYUmJIVzZnUzgyMUpsVGQveWhOUnJrCnB1Z1BIQUxrSGdkQndpT294akI3V2QraGU2VWxtVnJVMHo1R1lEb24wOUx6eVJ4WE9lZ2t4T3UrK09FTE9UbTQKaEhqQmhlM0JNbzdRRjBtdXh2SUxOd1gxTkpaZ1UzdmZkNFZsOHdDQlNNdWJEMkVaazdBdE0zZGtjaFlkRXB0dQpDRkJWdjVTRXRiTjlLZzNqb2o1RW5LbkJHUXR4VFk3V1RxZGlHSCtSZWhMV2ZQNWxSRE9jNUhDV1JKMEhBMVU4Ci93YTFzbTJYblhMWnB5bXBmZzJUblVxeXV4eDk4RGE0VUgzV3FWdVhveml6YVBPSE9EVWMxUW5BWXNPUU1DVXkKZk5TNS9MMXpGYnR5d2xyOHlZUUx5cmthVWwwK3MxSXByalFVUEo2VTVRS0NBUUFZd2NQOW1VQWhuK1BOTWtXMgo4SDhTblBRaFUxR2MxYkVLdTdpTDhxMmlSaG5JNUU1c2QyZlR4b0xZT0FOVW9HSXQvUUx6TjZydVQ4OXkzWHQ3CnprVkFmMnpaV1ZVSXdpRUozWi9XcnNHWWpXY0h6MTdlZ09ISXFua05VV3R4VzdNcmVPa2JqS0hnaHU1ZGlzZUwKZVBLaiswUU83WUZzQXVIeURpYUM2b1FuV2tYd1JHOGlHb0tPcnkzVllRT3ROUDRydUhLZzdOV3ZHUWQvZmwzUAozQ210c3R6YlFneGl0REE4T0txY1RSVUtOZm5Lbk1VZDI1Ym1Fcm92K0M3QVo5VEV1MTdVTkdRdzVlZ0FQY1kzCkVmWHljSTlvNHhaMjB0SVNRZTgzUWVCZTdUUVd1T21pMlV5aVBiQ1NNQkxrUDVFNkU1Rit3dlgyckx2MnZXenUKemY4N0FvSUJBUUNFbUhOcW5WSUtPbHpIT1Zua0F6MDY2TzRZY2t4ZDNvZUVqNmx0YTBXNGp5VmhlbFZMVzRDUQpNMm5MekxoQ3Irb1J4bTk4Q1lHSW5aZXVJbmZ3Q1o1cUZra3pnajZ1WnZ1S3dpUnV2aUROaHRkZmNuRG1iNGE3CmY3QUc5anhHeHF4d3ZtTmVxd2IwdzA4QVhyRzlEbEtZOHZwdmVSU2pmMFRYZ0VldEtTb3JVN2RRNnJ4VlRnUUsKREJUUmRqNnU1U2t2bE9IVHppOWM4MkVSa0ZTN0NhMWFHWTM1YmdqQWUvUzJGMk1LcWVmUUFxMmxiMExhUTJZSgpjQXBLTzBwcGNQMjhUY2NGQ1d6b2VnZTREQTkrMnQ3ck5hajAySzNyYzBXQm50cERaano0SVY4dThXaVhWR3VCCkVmYmFxOEVWQ2FTS3h0eTQzbmVtMUF4aVBoZ0ZQT1RWQW9JQkFRRGNwUlpXWkRvMEVnU2haWG5HQ0NaYUQxNlQKVXA0eFRZRmg2NmZsQThJdHBKUDVWdWJ4eER5YURRanJnbDFZYXovNTdyUXdLUk92czdDQ3YvUUZVN2hkcEZOTAoxVHdkTjgrS2Q5VmtnKzR0YjBJQ2h5cDJacW9VMmpsRkVMMjBlSUdRWTdLUGhJREtYMisvcVQ2bEZOWGdxQVZyClBWbmJuWUl0SGd3ZWtBZjVPbHFreDV2SjMvaHVERDRnRVg3NElGemMwVnVqS3pZZmxoZ2xxSVd1TDdQd1V3M1kKSTBQZVdtVjRZejJSVnI1amd3T3M5eFd1QWxXNUVUdHQrUysreEpZVVpQYytlL0FpSGg3SElaTHFrMU5KSk1VUAo4YklnaTNIS2Eza21HajVMbHQ1Kzl4WG4vTzgvY0Z3ak9nSHFqNTJRL09teDdnWTcvNkRwVmdaOUxVZm8KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K\n    token: o3f9fxe7v3k0fgtdxt5wa0bi7gcwv1jmqtje43z1wk8qh8385a7yw8ccnai5gm7ee7o6b95mvsdvw2skjatcn1rfptoevq2j5urzwv03klst997cpo7y7dei6s961qam\n"
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-08-20T06:15:57Z",
  "keyValueList": [
    {
      "key": "Location",
      "value": "koreasouth"
    },
    {
      "key": "Identity",
      "value": "{principalId:af8b295b-80f2-4e28-b34b-11a597b8a6d4,tenantId:fb98dda1-32ff-48eb-a489-62777cd9ccd8,type:SystemAssigned}"
    },
    {
      "key": "Kind",
      "value": "Base"
    },
    {
      "key": "Properties",
      "value": "{agentPoolProfiles:[{count:2,currentOrchestratorVersion:1.34.8,eTag:9235cebc-da3f-4363-9ea7-3b0097a130f9,enableAutoScaling:false,enableFIPS:false,enableNodePublicIP:true,kubeletDiskType:OS,maxPods:110,mode:System,name:workers1,nodeImageVersion:AKSUbuntu-2204gen2containerd-202608.06.1,orchestratorVersion:1.34.8,osDiskSizeGB:100,osDiskType:Managed,osSKU:Ubuntu,osType:Linux,powerState:{code:Running},provisioningState:Succeeded,scaleDownMode:Delete,securityProfile:{enableSecureBoot:false,enableVTPM:false,sshAccess:LocalUser},type:VirtualMachineScaleSets,upgradeSettings:{maxSurge:10%,maxUnavailable:0},vmSize:Standard_B4as_v2,vnetSubnetID:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbjvs627fsvq9c8qjdmm/subnets/tbi5s3q7etm8cahfjubo}],autoUpgradeProfile:{nodeOSUpgradeChannel:NodeImage},azurePortalFQDN:dns-1787206565697750064-yka2fsje.portal.hcp.koreasouth.azmk8s.io,bootstrapProfile:{artifactSource:Direct},currentKubernetesVersion:1.34.8,dnsPrefix:dns-1787206565697750064,enableRBAC:true,fqdn:dns-1787206565697750064-yka2fsje.hcp.koreasouth.azmk8s.io,identityProfile:{kubeletidentity:{clientId:abd41f07-a701-4feb-ba6b-1285f8ab81da,objectId:7a46f1bd-30fb-4ee8-bf99-b27c0e72e658,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/tbkpeno23hnamhgbrptu-agentpool}},ingressProfile:{webAppRouting:{dnsZoneResourceIds:[/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/dnszones/tbkpeno23hnamhgbrptu.com],enabled:true,gatewayAPIImplementations:{appRoutingIstio:{mode:Disabled}},identity:{clientId:75ca6f0a-bac5-4065-99b8-62477be438c6,objectId:383c4ad1-bd05-49bd-a902-ecb08d20c4e3,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/webapprouting-tbkpeno23hnamhgbrptu},nginx:{defaultIngressControllerType:AnnotationControlled}}},kubernetesVersion:1.34.8,linuxProfile:{adminUsername:cb-user,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDA7EqTWJv+hdZCkyjquj4TaYf5mYZQ4lsvvjLwuWaFV6BKS0Owt/B+rMl4PYS6zuK2hBjBFFYVY/hmPyhUiSzyX93wWSrHk2axbJyL++kJSVQ0dmO8Tcnr47K8Tp6WI+Ex72s5vIeKnltDbjLIsYHcisuB0bPNfDREFO3aUkKjub3l40Nowq2SmlOV4pB1MpUigKUPrh4QVrRhF0vdcXSzZ1wYCfWiPzF/fU0d+7f7raTwc8bDgHRSYsOQ7r4Je8XV2HPyG4kMU1z55l1RZkWE4Nhn2Ojy+H2nq/PGt+LB9M10U7f1iykdIG7N40SoyGnTekyPDZ9/o8E8yOjaMMAhBLFy9xs1VTARi2ZhJUKF/8MQRV/hZYMEZclYH8CGCiFvDBsU0xCgKjtsso8WO2zoGcmdZTDQb4+gy/wvvHRFGQPnTY79MfKPLtUu3B7c0uDVGtS5reeDlDsU2Is2e7pFX29sMZ3r9C8w61raTlUUagTLNbiHkTPiXgU4MFHDIeDpqKP5dBF6FW/qq4whuwQJSpjVg7H/m529RCo+5WgtnoZNEbvVZbXi5R69AEemoEvanua+ThwtytrE2yOsu7hVo0kvMjezxgodoyZLxwxYi8trvvNJxen3Ck2c6r/AzVx0slcjguz3uNi0RQf0hQQ30gva4VzDVjZwLWQx5Ngz6Q==\\n}]}},maxAgentPools:100,metricsProfile:{costAnalysis:{enabled:false}},networkProfile:{dnsServiceIP:10.1.0.10,ipFamilies:[IPv4],loadBalancerProfile:{backendPoolType:nodeIPConfiguration,effectiveOutboundIPs:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.Network/publicIPAddresses/3db7c26c-f4e9-4809-8c72-397081fa2272}],managedOutboundIPs:{count:1}},loadBalancerSku:standard,networkDataplane:azure,networkPlugin:azure,networkPolicy:azure,outboundType:loadBalancer,serviceCidr:10.1.0.0/16,serviceCidrs:[10.1.0.0/16]},nodeProvisioningProfile:{mode:Manual},nodeResourceGroup:CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth,oidcIssuerProfile:{enabled:true,issuerURL:https://koreasouth.oic.prod-aks.azure.com/fb98dda1-32ff-48eb-a489-62777cd9ccd8/56872a0b-9f7c-48ba-9126-59931441f480/},powerState:{code:Running},provisioningState:Succeeded,resourceUID:6a869bac954cc2000164bc4f,securityProfile:{},servicePrincipalProfile:{clientId:msi},storageProfile:{diskCSIDriver:{enabled:true},fileCSIDriver:{enabled:true},snapshotController:{enabled:true}},supportPlan:KubernetesOfficial,windowsProfile:{adminUsername:azureuser,enableCSIProxy:true},workloadAutoScalerProfile:{}}"
    },
    {
      "key": "SKU",
      "value": "{name:Base,tier:Standard}"
    },
    {
      "key": "Tags",
      "value": "{createdAt:1787206557,ownerCluster:tbkpeno23hnamhgbrptu,sshkey:tbh9oq28l3i7cb1062kn,sys.connectionName:azure-koreasouth,sys.createdTime:2026-08-20 06:15:57 +0000 UTC,sys.cspResourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tbkpeno23hnamhgbrptu,sys.cspResourceName:tbkpeno23hnamhgbrptu,sys.description:Migrated from on-premise K8s cluster (v1.32.3, 2 workers),sys.id:mig02-on-prem-k8s-cluster,sys.labelType:k8s,sys.manager:cb-tumblebug,sys.name:mig02-on-prem-k8s-cluster,sys.namespace:mig01,sys.uid:tbkpeno23hnamhgbrptu,sys.version:1.34.8}"
    },
    {
      "key": "ETag",
      "value": "b754aaf9-c1eb-4055-8256-4a1942a7e032"
    },
    {
      "key": "ID",
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tbkpeno23hnamhgbrptu"
    },
    {
      "key": "Name",
      "value": "tbkpeno23hnamhgbrptu"
    },
    {
      "key": "Type",
      "value": "Microsoft.ContainerService/ManagedClusters"
    }
  ],
  "cspResourceName": "tbkpeno23hnamhgbrptu",
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tbkpeno23hnamhgbrptu",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbkpeno23hnamhgbrptu",
      "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tbkpeno23hnamhgbrptu"
    },
    "Version": "1.34.8",
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
          "NameId": "#aks-agentpool-15876628-nsg",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/cb_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.Network/networkSecurityGroups/aks-agentpool-15876628-nsg"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tbkpeno23hnamhgbrptu/agentPools/workers1"
        },
        "ImageIID": {
          "NameId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd/versions/202608.06.1",
          "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/AKS-Ubuntu/providers/Microsoft.Compute/galleries/AKSUbuntu/images/2204gen2containerd/versions/202608.06.1"
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
            "NameId": "aks-workers1-34776577-vmss_0",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-34776577-vmss/virtualMachines/0"
          },
          {
            "NameId": "aks-workers1-34776577-vmss_1",
            "SystemId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.Compute/virtualMachineScaleSets/aks-workers1-34776577-vmss/virtualMachines/1"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://dns-1787206565697750064-yka2fsje.hcp.koreasouth.azmk8s.io:443",
      "Kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUU2RENDQXRDZ0F3SUJBZ0lRZDB5Tm1qRmJwcEZ4TEFMcEV1cVprekFOQmdrcWhraUc5dzBCQVFzRkFEQU4KTVFzd0NRWURWUVFERXdKallUQWdGdzB5TmpBNE1qQXdOakEyTkRsYUdBOHlNRFUyTURneU1EQTJNVFkwT1ZvdwpEVEVMTUFrR0ExVUVBeE1DWTJFd2dnSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJS0FvSUNBUUMzCmZtMGR6YUJvdnVWb3FoWlJTUlcxS0x1RHM4ZFk5M05qUXdHaGdhNUV6R1YzZ0UwcHptbHZMdlRqRVh2aFhjaEQKUDYxV2VPb1d6N3laSFpraUsyaEZGV09PbmV5NGdTMDdUTDQ3bTFPT1pzTzF4R09sN3JoM005VCt0TGdmVjhobApMblJOMWJmMEliM3N3cHhRd1MxVlh4K203Z2MxaEZkaFh6QW11NWwzWkhWY0pCVWlkVjZKdjFMU3lMc2NSL3V5CnFVcTQxYXFITllXQ0xKZFc5OFg1K1BBdUtteWRkM0w0R0w4SXhKaDZORXdvWmloYXhFTkhaRnlnK0JJMC9OTXEKU0psb0s1U1RFUWdPendiT2h5eFEzRUMyZnhZcG5VUm5rR05lUW1VckdWUk92UmphSEJrcUpid2ptOGo1ZVhmTQpSVGpnMG5rVTYwL1NBejRUazdFdjErakJhN3B1emdCMm8vQ0lnY3lpV3Y4aldvYXVsNnNUcEhLQit3eEFZaUdVCmtPR0Fsbldmc3RyTGRDTmFOMUs2Y2p3R24wajhvaXhxUXNkRjNPTWRqT3MyVHM3eUJRV2VZbFNhblNjYk5uYisKenZpeFZ5TE1DT1J4dWdYNi8yS0p3ZC8vUEhnVUJHS20xQzlXZDdHYm9uOW1MK21ObGZzb004dURYbmdoRnlhWgp2ZVVGMEpNZjNHTWJCNi9Cc29oK1V0bUlTb3Y0ejI0eGl4d25ub1Z2OTRoQjFjOFA5UTZ3MDNKRE91Wms3QjRBCjJ2NzNWYWRYRzYxZjluTXFkUlFGcXUyU2RmRDdmVVhPeG9lVnYzeUxmbllyOFdCTmJKNTdYdE9lVmVLS3g0Z1oKQU95MzJEb2d4ZFYwam9lMTZmWVdNa2NFNEt3NXlUQTdwMTF0cSsxSXJ3SURBUUFCbzBJd1FEQU9CZ05WSFE4QgpBZjhFQkFNQ0FxUXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVqYTRuTEJzVzIvUTk3NU5qCkhmQnArUkttVkRzd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dJQkFBRGpuZWVreWQ4SVlranlWSWc0ZXpkSmVHemEKclo1RlJteDFNZVlsVW02c3ViN3ZrRmFYM3pPSHZ0anZBeEp1TytrNjFVVmw4eUUvQXVVSG11c2NlaFY4WlpQRAozRjcvSU4zeUxkSUx6L0lsdkJrS09xUk5sWW5wZmYvRmxFWS9uY2Jkajg4UytwVzJaMDY4U2o0VnkvTGJrczFzCmMwTis5RVcxdjlMekZEOUtaNG16ckVKRXV2all0YUZOOUo5cTdJNGtnRVJEcEdwa1ZCSlQvNUN0NGlCSG50V3MKWVAwU2J2QmgyQ0hJTVZSVEVoQ2l6ZmN3UmlWekpCYTNVNXRlTDk5T1NaRDJ0Q0paSUlQREU0bjBXcWsrK3JSdgowU25QUzdiTHUrWFQ0eExVb3pyUlRsUVN0bzB6RjJHMXRpNlUyVTNjU3Nod2xYSHpXbkpxd0ZBZDFGcHdncU9uCnFDU25TRGx2dnlKRmF6WnBBUi8yM01iRGlUQ2tLckpoM3BhOUpieHVnaUVkcXJVWTd2SGtPTFdsaXNuLzJHMTYKcExXRk5DYVBveWxWc2QvNTQxQWpUS3ExSDduK0llM09RTTJ6TmFZbzZTMnc5SnlWNE1mMm5CNUsyNFN1UXNOegpTTnJpRVlrNGE4UW9OV0xxakluNkpLTGZ0c2ZNZTIxTCtpVDFKSnJEQVMyZVlOU28yOUJPd2tQSGx0VFpNUmtoCnBwYU5CZy82Y1l6MkkwZjlOYU02WHZ5Y2ovcGhldzdTbDllVkp4OW84N205cWo4ZjREYXBSRlp1ekwxdWVPc24KVTNJL0tTUUZxTkV1TXlFanY1Z0IyMng4RFh6OVRLSWNTSDJiTWhkemd3eG4wcW9PdTdlV0FURUt0NEFMQ01KMgpJeVVVLzlnb0FyZnR1aVNtCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://dns-1787206565697750064-yka2fsje.hcp.koreasouth.azmk8s.io:443\n  name: tbkpeno23hnamhgbrptu\ncontexts:\n- context:\n    cluster: tbkpeno23hnamhgbrptu\n    user: clusterAdmin_koreasouth_tbkpeno23hnamhgbrptu\n  name: tbkpeno23hnamhgbrptu\ncurrent-context: tbkpeno23hnamhgbrptu\nkind: Config\nusers:\n- name: clusterAdmin_koreasouth_tbkpeno23hnamhgbrptu\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZIakNDQXdhZ0F3SUJBZ0lSQU1OeXpjWWs4dVBWWWxQSGZiUXkzSUV3RFFZSktvWklodmNOQVFFTEJRQXcKRFRFTE1Ba0dBMVVFQXhNQ1kyRXdIaGNOTWpZd09ESXdNRFl3TmpVd1doY05Namd3T0RJd01EWXhOalV3V2pBdwpNUmN3RlFZRFZRUUtFdzV6ZVhOMFpXMDZiV0Z6ZEdWeWN6RVZNQk1HQTFVRUF4TU1iV0Z6ZEdWeVkyeHBaVzUwCk1JSUNJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBZzhBTUlJQ0NnS0NBZ0VBMk9NaGRNM2loTEJFdGx6bzNkbC8KWHNHUEd4UEJLRWg1VDZxcEhoVXZCV1praVdTK1FoRm9KOHpzU3JRTmgvZjR2a3JnK0dnQURKTGNVcEhmQXl5RApBRjdCOG93bDN6TEVyTkt1R2h3ajBDUUFRYkZ5czl0YjVqRVE1dm1QRDd3MkZpWE0xUWJySnBmNVE4UmluT3A0ClpIVXpiRGR0bnZLdEpobmJiM1doSHBDUm9iSEdZZVNxeW0zaWZ1SGRseUIwbTVlcHptcU41N0pxbWFXT2cyRUEKcEVZY2xxNnJnV2sycXhoNjBicFRnOXV2NjFKZ0ZGQStVcTQ1THcvdDdmM1VLRFpxVnBVMDlIdGdURnZZNGVRZQorZFdnOFNLMEE0M082YU1jM1E3Ynhkcmw4aEZRSkRML1RWZDdwaS9VMmk5MUdYMHhMYW9lNkpyQmhERnZad3hJCm1xTFoyVDZPbElTV2NaN3ByY3VqTUUrbE1IQXBIVGRIWkpoRVFnRjc1RVlxMXUwRkRZQUVDUll2OGlKUUtVQ0gKTkxWZVlGNlRwOW92YnlGUWdpM3hJcXNiVWpaNGgzdmU0ei9Nc2JYVWVYYTh4WUlQakkremJpb0JrWjV3T1BpMwpQaE5jLzVpRXB6dEkzOUdLeERGUldYNXlHb0ZYd0VXbloxWVRlL2JJZ21vMHRuK1lmaldCUWRXVFhGYmN4bFllCnpCY0NoMFRIN1llUHd1MWZqYUZQK05IanBaZHNQSERYbkVPU2ZFSHN4cUhDc3NtbzNQWWRUaFNzdEJ1SmlnZDMKZlJKK3VMUDcvT3o0ZkRKUjQ1ek9UTHhncnBZQlZVeUlJZmwrM2xUcEpaek5KR3dtRFZLQkFlZ1FkY3BkMElxUgpFblh2dmdnZzdIckIvSmFmSnhISVZDTUNBd0VBQWFOV01GUXdEZ1lEVlIwUEFRSC9CQVFEQWdXZ01CTUdBMVVkCkpRUU1NQW9HQ0NzR0FRVUZCd01DTUF3R0ExVWRFd0VCL3dRQ01BQXdId1lEVlIwakJCZ3dGb0FVamE0bkxCc1cKMi9ROTc1TmpIZkJwK1JLbVZEc3dEUVlKS29aSWh2Y05BUUVMQlFBRGdnSUJBSFh4T2s5WHI3K2U3T01TclhMTwptZFNVQWdPaFpDTXFSL014b1FiWHkvOUYwZHBsSkZINjlJVzh6TEo0ZDhWdnphUGNNZTQ4YmdHUDdIbko0SVJRCktZYjk2Y3cvZENuUjdRQ0xkM2VPNEd2dG1xV2dWQmIwdzN3WDYzSnFHL1ZQUkdqd0d4aFEvQmJ2clJpa1BESTMKeEpDdWRLNmdYSmJYN1NFUkFIemhiTnR0dnRwYkFFSWdjWVVLSjVaaDdvZTFYZHlqSzhJYmtjVkg1aXlPb2dTbApjaktBUExDbWVSVWVsdlZMWFVHVC9LaDk0ZGVIMzJkdGF0aUdoY0N3d1pYK2pvVXlEY2RSdWdDUHJsRVE1VzNzClg0bU90aUNNa2Y3MzB4ck9CaGNtWWNPM2tEWkNsWDdsYml3V2hpV20rTVZ2ODdXWnZrcS9QcWVzaVFVOEkzc3AKMTJJSUw4N0h2UkhUZEw3RGJTN2JCY2FlTksySjQvSitNaStpVkRoOXNpY2VzMnd1Znd0RWdOZlY5NDlPbklweApQQlkwUC94c1lZN3Naa2JaLzNKRzBqQWp4TDRjaU53NmlSSmhqVEQ2cFQ3K0JXNUlKQzAxbkU0L0tJaXhOclkyCkduL0pyVFJRWVRzRTNUcWFyeHVlMTJRQ3phZEN4M1pKV0R4bUNzSWpTdTVKV1dKaHBtUDl5amowaEFnUS9KSkoKNWY2Y3QvV05mR2pQVysxYllmNjhpUDQraWVpK2FReVNEaEtiOGhhTW4zcXE2K0FFbzVRbXh6ZmhxVytROHQzNgo5UnRKNElSOThqVzJqemYwR2VVQko3d1hyYnZyRFJTNXJOTjdYOUNVRlI1UWJjcWNsV21XWjhPQnNEZEdzTGVSCmNnN05OTkwzVkpvSEZlelJMdVZZK3JwNQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS1FJQkFBS0NBZ0VBMk9NaGRNM2loTEJFdGx6bzNkbC9Yc0dQR3hQQktFaDVUNnFwSGhVdkJXWmtpV1MrClFoRm9KOHpzU3JRTmgvZjR2a3JnK0dnQURKTGNVcEhmQXl5REFGN0I4b3dsM3pMRXJOS3VHaHdqMENRQVFiRnkKczl0YjVqRVE1dm1QRDd3MkZpWE0xUWJySnBmNVE4UmluT3A0WkhVemJEZHRudkt0SmhuYmIzV2hIcENSb2JIRwpZZVNxeW0zaWZ1SGRseUIwbTVlcHptcU41N0pxbWFXT2cyRUFwRVljbHE2cmdXazJxeGg2MGJwVGc5dXY2MUpnCkZGQStVcTQ1THcvdDdmM1VLRFpxVnBVMDlIdGdURnZZNGVRZStkV2c4U0swQTQzTzZhTWMzUTdieGRybDhoRlEKSkRML1RWZDdwaS9VMmk5MUdYMHhMYW9lNkpyQmhERnZad3hJbXFMWjJUNk9sSVNXY1o3cHJjdWpNRStsTUhBcApIVGRIWkpoRVFnRjc1RVlxMXUwRkRZQUVDUll2OGlKUUtVQ0hOTFZlWUY2VHA5b3ZieUZRZ2kzeElxc2JValo0CmgzdmU0ei9Nc2JYVWVYYTh4WUlQakkremJpb0JrWjV3T1BpM1BoTmMvNWlFcHp0STM5R0t4REZSV1g1eUdvRlgKd0VXbloxWVRlL2JJZ21vMHRuK1lmaldCUWRXVFhGYmN4bFllekJjQ2gwVEg3WWVQd3UxZmphRlArTkhqcFpkcwpQSERYbkVPU2ZFSHN4cUhDc3NtbzNQWWRUaFNzdEJ1SmlnZDNmUkordUxQNy9PejRmREpSNDV6T1RMeGdycFlCClZVeUlJZmwrM2xUcEpaek5KR3dtRFZLQkFlZ1FkY3BkMElxUkVuWHZ2Z2dnN0hyQi9KYWZKeEhJVkNNQ0F3RUEKQVFLQ0FnQkI3RWV2Q1NWZ3ozTVRPd3BNNUY4aW5oS3hXRC9OenJtUXpYNjU5aFprdmNxeE9EM2NOdzVCaXJnSAp2TktnRVc4NTUraVptSUxyVDNoSVlLNDRlTDhZemJTRjFMTnVOREF6bDVYenVibm8rZ2haNzJXOTVWNzVpTkJxClpGQm5wLzJJbmRTMHEzV3VOV00raGVLemIxRkl0NWI1dlo5RVFOOEFSYnU5RlRQejVsMWRtSHVFSmMwRDJvS04Kcm5sOEJoRnJlWjNUYisvU0RSajV1cWltcGtWYnFUUG5XUkFvTmFLNFBxaVdOdHhMcCtyQXpEa0g4NXY5NVpiYwpCeXQ2dXp4UlBMajF1RVJ3UzAvcDVjRDJhREJDSC96YlRvRUkwNEdnNGtOVHJjQi9VeG14aWpHaHp4NXFrN3l4CnRyZ3IyV0R1Ym04VVFqRkM0a2NQdHpiMVMzYUZldHZlaks1ZkJubTIzcjcvNlpUenhsek9Vd0VBcmtZejVJZkoKYTE3Wkt0b2wyOWNrRmsvaGNyb0R5b0ZCUTlGTUZ5MGxYM1VmRkh3T1FvazVrVjNFQmhJK2NMbUR4YXZaN3cweQptdXlpWnZLYjlEQ0lEdlIxQ2hxTGVuVzBTd1M1dTV0WHhiYWFTVlQzV3h2VDZYSDFxYTc1alpKVEJjNG1kdzkvCjNGbjlQQ2RZUVdsNHZKWEJsdUtFSVJOYktlNWhPSnVwVmMrK0hJSnd5clg3cm1aS3lkS1RpeEFYZUlleitnNmIKWFVmZFZQdEJTMW5ZUXhTaVg1RFhTUjdBNUJ0TE02VUR4K0xUVlBXNjIvYXJKcUNqcjlSM2F4eHhHUnVOVTJvcApyMGp6UDJtdElKMVVVcUlDQ3dGTTFjeEFrUHBOUTMzSzBneVJmeU1INTBCMkVDaEhnUUtDQVFFQTYxQjhJalJFClJqOWsyQVA0TEh6VGN6WFRtdGVISGFQRXdraXJxMUdPTXljc0xyWDZUNWpMU05JK1p4bWprdnNFZ1o5SFFnQUIKYVAxSDNsSndHZWdONmg2UnkyN1hwUlFuTlJTK1o5VTlET3VzNEtRMDhMc2pOOUErRWlMYXlnVkt0V21BMExuWAozc05VQ2Y0dkZvNGhyRHdtUDlnMkhYM0FpNXN1TUswaDBhMytpUEYxNXdsWUFGeUxCMzZKRmdxK2ZyUjA3bGdYClVHak1YN1k4d0JubUFVaFBQSVFENDZlK1gwNEpuWE12VC9tZEw2Z1I2MDhBQmt4RmRMVkprcmM4SWJlN1h2dUYKckQzMnBpMmNUZm5mM0prdm5YclZhUlYwendjbGJyaGRsK0JVZ0ZwUE9SbmI5K3NEOTlzUGE1UHYzNEZOV2NHVAp1RjVCWnBBZXdEWDhad0tDQVFFQTYvUDErK2RzOStlaSttN2dmSXJ5TkFYUmJIVzZnUzgyMUpsVGQveWhOUnJrCnB1Z1BIQUxrSGdkQndpT294akI3V2QraGU2VWxtVnJVMHo1R1lEb24wOUx6eVJ4WE9lZ2t4T3UrK09FTE9UbTQKaEhqQmhlM0JNbzdRRjBtdXh2SUxOd1gxTkpaZ1UzdmZkNFZsOHdDQlNNdWJEMkVaazdBdE0zZGtjaFlkRXB0dQpDRkJWdjVTRXRiTjlLZzNqb2o1RW5LbkJHUXR4VFk3V1RxZGlHSCtSZWhMV2ZQNWxSRE9jNUhDV1JKMEhBMVU4Ci93YTFzbTJYblhMWnB5bXBmZzJUblVxeXV4eDk4RGE0VUgzV3FWdVhveml6YVBPSE9EVWMxUW5BWXNPUU1DVXkKZk5TNS9MMXpGYnR5d2xyOHlZUUx5cmthVWwwK3MxSXByalFVUEo2VTVRS0NBUUFZd2NQOW1VQWhuK1BOTWtXMgo4SDhTblBRaFUxR2MxYkVLdTdpTDhxMmlSaG5JNUU1c2QyZlR4b0xZT0FOVW9HSXQvUUx6TjZydVQ4OXkzWHQ3CnprVkFmMnpaV1ZVSXdpRUozWi9XcnNHWWpXY0h6MTdlZ09ISXFua05VV3R4VzdNcmVPa2JqS0hnaHU1ZGlzZUwKZVBLaiswUU83WUZzQXVIeURpYUM2b1FuV2tYd1JHOGlHb0tPcnkzVllRT3ROUDRydUhLZzdOV3ZHUWQvZmwzUAozQ210c3R6YlFneGl0REE4T0txY1RSVUtOZm5Lbk1VZDI1Ym1Fcm92K0M3QVo5VEV1MTdVTkdRdzVlZ0FQY1kzCkVmWHljSTlvNHhaMjB0SVNRZTgzUWVCZTdUUVd1T21pMlV5aVBiQ1NNQkxrUDVFNkU1Rit3dlgyckx2MnZXenUKemY4N0FvSUJBUUNFbUhOcW5WSUtPbHpIT1Zua0F6MDY2TzRZY2t4ZDNvZUVqNmx0YTBXNGp5VmhlbFZMVzRDUQpNMm5MekxoQ3Irb1J4bTk4Q1lHSW5aZXVJbmZ3Q1o1cUZra3pnajZ1WnZ1S3dpUnV2aUROaHRkZmNuRG1iNGE3CmY3QUc5anhHeHF4d3ZtTmVxd2IwdzA4QVhyRzlEbEtZOHZwdmVSU2pmMFRYZ0VldEtTb3JVN2RRNnJ4VlRnUUsKREJUUmRqNnU1U2t2bE9IVHppOWM4MkVSa0ZTN0NhMWFHWTM1YmdqQWUvUzJGMk1LcWVmUUFxMmxiMExhUTJZSgpjQXBLTzBwcGNQMjhUY2NGQ1d6b2VnZTREQTkrMnQ3ck5hajAySzNyYzBXQm50cERaano0SVY4dThXaVhWR3VCCkVmYmFxOEVWQ2FTS3h0eTQzbmVtMUF4aVBoZ0ZQT1RWQW9JQkFRRGNwUlpXWkRvMEVnU2haWG5HQ0NaYUQxNlQKVXA0eFRZRmg2NmZsQThJdHBKUDVWdWJ4eER5YURRanJnbDFZYXovNTdyUXdLUk92czdDQ3YvUUZVN2hkcEZOTAoxVHdkTjgrS2Q5VmtnKzR0YjBJQ2h5cDJacW9VMmpsRkVMMjBlSUdRWTdLUGhJREtYMisvcVQ2bEZOWGdxQVZyClBWbmJuWUl0SGd3ZWtBZjVPbHFreDV2SjMvaHVERDRnRVg3NElGemMwVnVqS3pZZmxoZ2xxSVd1TDdQd1V3M1kKSTBQZVdtVjRZejJSVnI1amd3T3M5eFd1QWxXNUVUdHQrUysreEpZVVpQYytlL0FpSGg3SElaTHFrMU5KSk1VUAo4YklnaTNIS2Eza21HajVMbHQ1Kzl4WG4vTzgvY0Z3ak9nSHFqNTJRL09teDdnWTcvNkRwVmdaOUxVZm8KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K\n    token: o3f9fxe7v3k0fgtdxt5wa0bi7gcwv1jmqtje43z1wk8qh8385a7yw8ccnai5gm7ee7o6b95mvsdvw2skjatcn1rfptoevq2j5urzwv03klst997cpo7y7dei6s961qam\n"
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-08-20T06:15:57Z",
    "KeyValueList": [
      {
        "key": "Location",
        "value": "koreasouth"
      },
      {
        "key": "Identity",
        "value": "{principalId:af8b295b-80f2-4e28-b34b-11a597b8a6d4,tenantId:fb98dda1-32ff-48eb-a489-62777cd9ccd8,type:SystemAssigned}"
      },
      {
        "key": "Kind",
        "value": "Base"
      },
      {
        "key": "Properties",
        "value": "{agentPoolProfiles:[{count:2,currentOrchestratorVersion:1.34.8,eTag:9235cebc-da3f-4363-9ea7-3b0097a130f9,enableAutoScaling:false,enableFIPS:false,enableNodePublicIP:true,kubeletDiskType:OS,maxPods:110,mode:System,name:workers1,nodeImageVersion:AKSUbuntu-2204gen2containerd-202608.06.1,orchestratorVersion:1.34.8,osDiskSizeGB:100,osDiskType:Managed,osSKU:Ubuntu,osType:Linux,powerState:{code:Running},provisioningState:Succeeded,scaleDownMode:Delete,securityProfile:{enableSecureBoot:false,enableVTPM:false,sshAccess:LocalUser},type:VirtualMachineScaleSets,upgradeSettings:{maxSurge:10%,maxUnavailable:0},vmSize:Standard_B4as_v2,vnetSubnetID:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/virtualNetworks/tbjvs627fsvq9c8qjdmm/subnets/tbi5s3q7etm8cahfjubo}],autoUpgradeProfile:{nodeOSUpgradeChannel:NodeImage},azurePortalFQDN:dns-1787206565697750064-yka2fsje.portal.hcp.koreasouth.azmk8s.io,bootstrapProfile:{artifactSource:Direct},currentKubernetesVersion:1.34.8,dnsPrefix:dns-1787206565697750064,enableRBAC:true,fqdn:dns-1787206565697750064-yka2fsje.hcp.koreasouth.azmk8s.io,identityProfile:{kubeletidentity:{clientId:abd41f07-a701-4feb-ba6b-1285f8ab81da,objectId:7a46f1bd-30fb-4ee8-bf99-b27c0e72e658,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/tbkpeno23hnamhgbrptu-agentpool}},ingressProfile:{webAppRouting:{dnsZoneResourceIds:[/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreasouth/providers/Microsoft.Network/dnszones/tbkpeno23hnamhgbrptu.com],enabled:true,gatewayAPIImplementations:{appRoutingIstio:{mode:Disabled}},identity:{clientId:75ca6f0a-bac5-4065-99b8-62477be438c6,objectId:383c4ad1-bd05-49bd-a902-ecb08d20c4e3,resourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.ManagedIdentity/userAssignedIdentities/webapprouting-tbkpeno23hnamhgbrptu},nginx:{defaultIngressControllerType:AnnotationControlled}}},kubernetesVersion:1.34.8,linuxProfile:{adminUsername:cb-user,ssh:{publicKeys:[{keyData:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDA7EqTWJv+hdZCkyjquj4TaYf5mYZQ4lsvvjLwuWaFV6BKS0Owt/B+rMl4PYS6zuK2hBjBFFYVY/hmPyhUiSzyX93wWSrHk2axbJyL++kJSVQ0dmO8Tcnr47K8Tp6WI+Ex72s5vIeKnltDbjLIsYHcisuB0bPNfDREFO3aUkKjub3l40Nowq2SmlOV4pB1MpUigKUPrh4QVrRhF0vdcXSzZ1wYCfWiPzF/fU0d+7f7raTwc8bDgHRSYsOQ7r4Je8XV2HPyG4kMU1z55l1RZkWE4Nhn2Ojy+H2nq/PGt+LB9M10U7f1iykdIG7N40SoyGnTekyPDZ9/o8E8yOjaMMAhBLFy9xs1VTARi2ZhJUKF/8MQRV/hZYMEZclYH8CGCiFvDBsU0xCgKjtsso8WO2zoGcmdZTDQb4+gy/wvvHRFGQPnTY79MfKPLtUu3B7c0uDVGtS5reeDlDsU2Is2e7pFX29sMZ3r9C8w61raTlUUagTLNbiHkTPiXgU4MFHDIeDpqKP5dBF6FW/qq4whuwQJSpjVg7H/m529RCo+5WgtnoZNEbvVZbXi5R69AEemoEvanua+ThwtytrE2yOsu7hVo0kvMjezxgodoyZLxwxYi8trvvNJxen3Ck2c6r/AzVx0slcjguz3uNi0RQf0hQQ30gva4VzDVjZwLWQx5Ngz6Q==\\n}]}},maxAgentPools:100,metricsProfile:{costAnalysis:{enabled:false}},networkProfile:{dnsServiceIP:10.1.0.10,ipFamilies:[IPv4],loadBalancerProfile:{backendPoolType:nodeIPConfiguration,effectiveOutboundIPs:[{id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth/providers/Microsoft.Network/publicIPAddresses/3db7c26c-f4e9-4809-8c72-397081fa2272}],managedOutboundIPs:{count:1}},loadBalancerSku:standard,networkDataplane:azure,networkPlugin:azure,networkPolicy:azure,outboundType:loadBalancer,serviceCidr:10.1.0.0/16,serviceCidrs:[10.1.0.0/16]},nodeProvisioningProfile:{mode:Manual},nodeResourceGroup:CB_koreasouth_tbkpeno23hnamhgbrptu_koreasouth,oidcIssuerProfile:{enabled:true,issuerURL:https://koreasouth.oic.prod-aks.azure.com/fb98dda1-32ff-48eb-a489-62777cd9ccd8/56872a0b-9f7c-48ba-9126-59931441f480/},powerState:{code:Running},provisioningState:Succeeded,resourceUID:6a869bac954cc2000164bc4f,securityProfile:{},servicePrincipalProfile:{clientId:msi},storageProfile:{diskCSIDriver:{enabled:true},fileCSIDriver:{enabled:true},snapshotController:{enabled:true}},supportPlan:KubernetesOfficial,windowsProfile:{adminUsername:azureuser,enableCSIProxy:true},workloadAutoScalerProfile:{}}"
      },
      {
        "key": "SKU",
        "value": "{name:Base,tier:Standard}"
      },
      {
        "key": "Tags",
        "value": "{createdAt:1787206557,ownerCluster:tbkpeno23hnamhgbrptu,sshkey:tbh9oq28l3i7cb1062kn,sys.connectionName:azure-koreasouth,sys.createdTime:2026-08-20 06:15:57 +0000 UTC,sys.cspResourceId:/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tbkpeno23hnamhgbrptu,sys.cspResourceName:tbkpeno23hnamhgbrptu,sys.description:Migrated from on-premise K8s cluster (v1.32.3, 2 workers),sys.id:mig02-on-prem-k8s-cluster,sys.labelType:k8s,sys.manager:cb-tumblebug,sys.name:mig02-on-prem-k8s-cluster,sys.namespace:mig01,sys.uid:tbkpeno23hnamhgbrptu,sys.version:1.34.8}"
      },
      {
        "key": "ETag",
        "value": "b754aaf9-c1eb-4055-8256-4a1942a7e032"
      },
      {
        "key": "ID",
        "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourcegroups/koreasouth/providers/Microsoft.ContainerService/managedClusters/tbkpeno23hnamhgbrptu"
      },
      {
        "key": "Name",
        "value": "tbkpeno23hnamhgbrptu"
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

