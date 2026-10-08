# CM-Beetle K8s Infra Migration Test Results — IBMCloud-Sydney

> [!NOTE]
> Full lifecycle against a real CSP: recommend → validate → migrate → list → get → workload → delete → residual.

## Environment

- CSP / Region: ibm / au-syd
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (cd1f3a2)
- Git Commit: cd1f3a2
- Namespace: mig01
- Test Date: 2026-10-07 20:05:55 KST
- Cluster ID: mig06-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 867ms |
| 2 | POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation) | ✅ **PASS** | 160ms |
| 3 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 22m16.034s |
| 4 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 2ms |
| 5 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 5ms |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 3m59.009s |
| 7 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 5m34.67s |
| 8 | Residual resource check (Tumblebug) | ✅ **PASS** | 6ms |

**Overall Result**: 8/8 steps passed ✅

**Total Duration**: 31m50s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 867ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.35.9)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=ibm+au-syd+cxf-4x8 nodes=2

### Step 2 — POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)

- **Duration**: 160ms
- **Status Code**: 200

- ✅ Target K8s infra model is valid for migration (0 issues)

### Step 3 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 22m16.034s
- **Status Code**: 202

- ℹ️  nameSeed: mig06
- ℹ️  async reqId: 1791371156924617100
- ℹ️  cluster id: mig06-on-prem-k8s-cluster
- ℹ️  elapsed: 22m16s
- ✅ status: Active

### Step 4 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 2ms
- **Status Code**: 200

- ✅ migrated cluster present in list (1 total)

### Step 5 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 5ms
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=ibm+au-syd+cxf-4x8, nodes=2)
- ✅ version: 1.35.9_1546 (recommended 1.35.9)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 3m59.009s

- ✅ kubeconfig obtained (server: https://c105.au-syd.containers.cloud.ibm.com:32343)
- ℹ️  auth method: client certificate in kubeconfig
- ✅ API server reachable (v1.35.9+IKS)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 3)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: 9c279c4b-au-syd.lb.appdomain.cloud
- ✅ nginx served over the LoadBalancer at 9c279c4b-au-syd.lb.appdomain.cloud (attempt 1)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 7 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 5m34.67s
- **Status Code**: 200

- ✅ deleted on attempt 1 (5m34s)

### Step 8 — Residual resource check (Tumblebug)

- **Duration**: 6ms

- ℹ️  VNet mig06-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig06-k8s-sg still exists (known gap)
- ℹ️  SshKey mig06-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for ibm au-syd (source: v1.32.3 → target: v1.35.9)",
  "targetCloud": {
    "csp": "ibm",
    "region": "au-syd"
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
    "connectionName": "ibm-au-syd",
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
    "connectionName": "ibm-au-syd",
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
      "connectionName": "ibm-au-syd",
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
    "connectionName": "ibm-au-syd",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "1.35.9",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": [
      {
        "name": "workers1",
        "imageId": "default",
        "specId": "ibm+au-syd+cxf-4x8",
        "rootDiskType": "default",
        "rootDiskSize": 100,
        "sshKeyId": "",
        "onAutoScaling": "false",
        "desiredNodeSize": 2,
        "minNodeSize": 2,
        "maxNodeSize": 2,
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
  "id": "mig06-on-prem-k8s-cluster",
  "uid": "tbn9e2fqrrhumthk6j20",
  "name": "mig06-on-prem-k8s-cluster",
  "connectionName": "ibm-au-syd",
  "connectionConfig": {
    "configName": "ibm-au-syd",
    "providerName": "ibm",
    "driverName": "ibm-driver-v1.0.so",
    "credentialName": "ibm",
    "credentialHolder": "admin",
    "regionZoneInfoName": "ibm-au-syd",
    "regionZoneInfo": {
      "assignedRegion": "au-syd",
      "assignedZone": "au-syd-1"
    },
    "regionDetail": {
      "regionId": "au-syd",
      "regionName": "au-syd",
      "description": "Sydney (Australia)",
      "location": {
        "display": "Australia (Sydney)",
        "latitude": -33.86882,
        "longitude": 151.209296
      },
      "zones": [
        "au-syd-1",
        "au-syd-2",
        "au-syd-3"
      ]
    },
    "regionRepresentative": true,
    "verified": true
  },
  "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
  "systemMessage": "",
  "label": {
    "sys.connectionName": "ibm-au-syd",
    "sys.createdTime": "2026-10-07 11:06:11 +0000 UTC",
    "sys.cspResourceId": "db32f95s0learmn856i0",
    "sys.cspResourceName": "tbn9e2fqrrhumthk6j20",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig06-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig06-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbn9e2fqrrhumthk6j20",
    "sys.version": "1.35.9_1546"
  },
  "systemLabel": "",
  "version": "1.35.9_1546",
  "network": {
    "vNetId": "mig06-k8s-vpc",
    "subnetIds": [
      "mig06-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "mig06-k8s-sg"
    ],
    "keyValueList": [
      {
        "key": "VpcIID",
        "value": "{NameId:tblo85b3uaik2dvvkjg2,SystemId:r026-1f28c17f-150a-434e-ae0f-a2a321fd4536}"
      },
      {
        "key": "SubnetIIDs",
        "value": "{NameId:tbp1ba7rn56rirasj2pe,SystemId:02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147}"
      },
      {
        "key": "SecurityGroupIIDs",
        "value": "{NameId:kube-db32f95s0learmn856i0,SystemId:r026-02b8e88f-2be0-4f01-8592-ce88b2a66999}"
      },
      {
        "key": "KeyValueList",
        "value": "{Key:VpcIID,Value:{NameId:tblo85b3uaik2dvvkjg2,SystemId:r026-1f28c17f-150a-434e-ae0f-a2a321fd4536}}; {Key:SubnetIIDs,Value:{NameId:tbp1ba7rn56rirasj2pe,SystemId:02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147}}; {Key:SecurityGroupIIDs,Value:{NameId:kube-db32f95s0learmn856i0,SystemId:r026-02b8e88f-2be0-4f01-8592-ce88b2a66999}}"
      }
    ]
  },
  "k8sNodeGroupList": [
    {
      "id": "workers1",
      "name": "workers1",
      "imageId": "default",
      "specId": "ibm+au-syd+cxf-4x8",
      "rootDiskType": "Not visible in IBM",
      "rootDiskSize": 0,
      "sshKeyId": "mig06-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 0,
      "maxNodeSize": 0,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "Not visible in IBM",
          "cspResourceId": "kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-00000168"
        },
        {
          "cspResourceName": "Not visible in IBM",
          "cspResourceId": "kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-000002da"
        }
      ],
      "keyValueList": [
        {
          "key": "IId",
          "value": "{NameId:workers1,SystemId:db32f95s0learmn856i0-c67c62f}"
        },
        {
          "key": "ImageIID",
          "value": "{NameId:Not visible in IBM,SystemId:Not visible in IBM}"
        },
        {
          "key": "VMSpecName",
          "value": "cxf.4x8"
        },
        {
          "key": "RootDiskType",
          "value": "Not visible in IBM"
        },
        {
          "key": "RootDiskSize",
          "value": "Not visible in IBM"
        },
        {
          "key": "KeyPairIID",
          "value": "{NameId:Not visible in IBM,SystemId:Not visible in IBM}"
        },
        {
          "key": "OnAutoScaling",
          "value": "false"
        },
        {
          "key": "DesiredNodeSize",
          "value": "2"
        },
        {
          "key": "MinNodeSize",
          "value": "0"
        },
        {
          "key": "MaxNodeSize",
          "value": "0"
        },
        {
          "key": "Status",
          "value": "Active"
        },
        {
          "key": "Nodes",
          "value": "{NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-00000168}; {NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-000002da}"
        }
      ],
      "cspResourceName": "workers1",
      "cspResourceId": "db32f95s0learmn856i0-c67c62f",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "db32f95s0learmn856i0-c67c62f"
        },
        "ImageIID": {
          "NameId": "Not visible in IBM",
          "SystemId": "Not visible in IBM"
        },
        "VMSpecName": "cxf.4x8",
        "RootDiskType": "Not visible in IBM",
        "RootDiskSize": "Not visible in IBM",
        "KeyPairIID": {
          "NameId": "",
          "SystemId": "Not visible in IBM"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "Not visible in IBM",
            "SystemId": "kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-00000168"
          },
          {
            "NameId": "Not visible in IBM",
            "SystemId": "kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-000002da"
          }
        ],
        "KeyValueList": [
          {
            "key": "IId",
            "value": "{NameId:workers1,SystemId:db32f95s0learmn856i0-c67c62f}"
          },
          {
            "key": "ImageIID",
            "value": "{NameId:Not visible in IBM,SystemId:Not visible in IBM}"
          },
          {
            "key": "VMSpecName",
            "value": "cxf.4x8"
          },
          {
            "key": "RootDiskType",
            "value": "Not visible in IBM"
          },
          {
            "key": "RootDiskSize",
            "value": "Not visible in IBM"
          },
          {
            "key": "KeyPairIID",
            "value": "{NameId:Not visible in IBM,SystemId:Not visible in IBM}"
          },
          {
            "key": "OnAutoScaling",
            "value": "false"
          },
          {
            "key": "DesiredNodeSize",
            "value": "2"
          },
          {
            "key": "MinNodeSize",
            "value": "0"
          },
          {
            "key": "MaxNodeSize",
            "value": "0"
          },
          {
            "key": "Status",
            "value": "Active"
          },
          {
            "key": "Nodes",
            "value": "{NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-00000168}; {NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-000002da}"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://c105.au-syd.containers.cloud.ibm.com:32343",
    "kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURRRENDQWlpZ0F3SUJBZ0lVZnp2VEoxdkpXQVVoOU93cnk0T0J2VHdqNExFd0RRWUpLb1pJaHZjTkFRRUwKQlFBd09ERTJNRFFHQTFVRUF4TXRaR0l6TW1ZNU5YTXdiR1ZoY20xdU9EVTJhVEF0YTNWaVpYSnVaWFJsY3kxagpZUzB4TnpreE16Y3hNalF5TUI0WERUSTJNVEF3TnpFeE1ETXdNRm9YRFRNMk1UQXdOREV4TURNd01Gb3dPREUyCk1EUUdBMVVFQXhNdFpHSXpNbVk1TlhNd2JHVmhjbTF1T0RVMmFUQXRhM1ZpWlhKdVpYUmxjeTFqWVMweE56a3gKTXpjeE1qUXlNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQThubGY3V3FQeFlhNgphRG0rc0VCWS9qSVBiNlFXb2RBY09zTEdTNDdRT24xTEN6TWlpRTJSZzdtemtoQmNnamNkUWRTQ1NCVTF3RWFQCkN1N01oNCtJTWlsUGxYdzRmaG13RHIwbVlpR01qdFVhN3UrcWtNYmlsSTU5cUpoYXVFMFBVYXV4ZWpEZUxzb3cKaXpuellXaVZ4MDFlcG9PQ3RHUDEzU3oxWmFzSjNBbXl2Z1B3T1BmTnk2U2hESGRtWUUrWGJEZW5qR2tOMWwxUwovVGZEcjk3bU5WREdDNHhUSU1Kcjl3S3htdmIyYnV0bksrb3NaUlNFRkpoRmFDYmRlQXNkb05mVEdlemFiLzNxCmQ0S2w0d1hkdk1wcnZGT3gyWS91UzlXSDlJZ21GSUc4N254RzE4cXNkV2E0OUxPSGFBT1Rrczl4Tmg4ZEhUNk0KWDFXRlNMRU1IUUlEQVFBQm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FZWXdEd1lEVlIwVEFRSC9CQVV3QXdFQgovekFkQmdOVkhRNEVGZ1FVWGtCVlU5QmZpczc2aHNtZ3hVeHRERVU0ZjZrd0RRWUpLb1pJaHZjTkFRRUxCUUFECmdnRUJBS3RrcnppcVdSNVVpTk5uOW1kRmN2L1haRFJlYkEraDMwVHFGTXA5UFRSM2dFVWhmclJqUVlJMGJWWUIKOEgxZUdQdStoV0lvUElzVUJvYTJDM3g2c1dJQWN6c0tUWm9lQ0xJdXJ4YWZlVXVjdEpmaHFVRHhiN3FxSFl1SAoxanp1TmdzZ2pDZGliSFpUeGFLSURWbTF4TW0vVWpQMW05TXc1eWUvcXhETElEaTRFaGtSTGQyc0xsdDZoTTQxCks2YlRYMjUydFFIbDA0Y05ISStPY1kyOVhSVnpvVW82VzNVdDEzZnEzMTBGbGNJZm0zQndUcWNDL2VYZFFCdkQKRFhJbjh1RXFBVGdHanpGM1BocDVVVHlSd2ZJNnUrWnFSdXlPNXFveW9jVC9uTUQxbUgyQW85UUNRQmhOL1JHUQp5WnQrcjJXSXdiVjFGaE5NcUl5MWk2SXBIL009Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://c105.au-syd.containers.cloud.ibm.com:32343\n  name: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0\ncontexts:\n- context:\n    cluster: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0\n    namespace: default\n    user: admin/db32f95s0learmn856i0\n  name: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0/admin\ncurrent-context: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0/admin\nkind: Config\nusers:\n- name: admin/db32f95s0learmn856i0\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURzRENDQXBpZ0F3SUJBZ0lSQVBDUk1tUGZnS1hMLytyRG5QU1NhY1l3RFFZSktvWklodmNOQVFFTEJRQXcKT0RFMk1EUUdBMVVFQXhNdFpHSXpNbVk1TlhNd2JHVmhjbTF1T0RVMmFUQXRhM1ZpWlhKdVpYUmxjeTFqWVMweApOemt4TXpjeE1qUXlNQjRYRFRJMk1UQXdOekV4TVRrMU4xb1hEVEk0TVRBd056RXhNVGsxTjFvd2F6RUxNQWtHCkExVUVCaE1DVlZNeEZqQVVCZ05WQkFnVERWTmhiaUJHY21GdVkybHpZMjh4Q3pBSkJnTlZCQWNUQWtOQk1STXcKRVFZRFZRUUtFd3BwWW0wdFlXUnRhVzV6TVNJd0lBWURWUVFEREJsTGRXSmxRMlZ5ZENOSlFrMXBaQzAyT1RRdwpNREZDVHpkWE1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBdC8vc0J3cFA1bDhpClNzM3RiSkFZZi9jN0U5QlM2N0s1YXRGRFJlcEYzUVZJUzBUNTBoY285WElqSW5IREl2Umo1VFMyYzlnQi9FekoKT3BoM3lJOFNVWTgzanVzZUlBalVLaHBadFUwNEMvZWFFcEhYVk9hdzQ3Qi9iTXlhY3Era3NLUHZ5WnN2eS9qeAo2bW1UOUxEZ1E2dTFHbTVtUkpnYzFUb2NvVDFacW1XSXN6WWFRY25VVnpDVGdFa3V3dkpyQkJjQ09tdllqWitrCnRLc3czM3pIQ01aNElkS2hZN3N4c0dEVTk4VHpLaG9ybEhZRGlMRWR3RThmeCtRUkREbVByRDZqbDlYTDFESUYKVUZ4Y2N4ZVZTS0FKcm9vRVRIN09nTmxTVWFZZ2I4UTZoZ1RJUVRaM1ZEYlZlQmR0OWdUMksxQXhxNkNnSzZ2ZQpEZEpuRzVmaTN3SURBUUFCbzRHQk1IOHdEZ1lEVlIwUEFRSC9CQVFEQWdLRU1CMEdBMVVkSlFRV01CUUdDQ3NHCkFRVUZCd01DQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRkY1QVZWUFEKWDRyTytvYkpvTVZNYlF4Rk9IK3BNQjhHQTFVZEVRUVlNQmFCRkdoaGJtbDZZVzVuTnpkQVoyMWhhV3d1WTI5dApNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUF3OUVkQkJWWmhTUHMrWXJqUDhoNUJVVWxWTzBzZnZhR1djMlVKCi9VTlVHVUJTZm1FTXFSanZHYythV1JtbjdnV1VBMW92azJ4NE10NU5WTWpNYkVOVEMwWk1Ja1NTdWttR3hOYW4KSkE2RGdXMXpOdlBzazIvMDNnQU1DbUJtd1RheC93V0lwWC9QQTU2RXk4bEhIRVZpdWttNS9uU0lGYUIxZ0ZNbQpvczB2OE1WdFc4YVVlRGEySTVrajRaQlZicytOc3hXU0RCK2RoL2NBcUlGVS9sQzgwakRjeFMxbkl0K2o3UnN3CjVQeFJHcEllbUM1NjdDNEMxdUpZa2tLazJ2WStvZ05GczlQVzlGTXBjdjdPUEtMc3ZNb0lyNER4L1lvK3V4dTEKS3BuRm40M21tUVk5TERMelliTCtyK2R2Nm9ZVVFEZDZqS0VQSS9CK293d3BzcGxxCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBdC8vc0J3cFA1bDhpU3MzdGJKQVlmL2M3RTlCUzY3SzVhdEZEUmVwRjNRVklTMFQ1CjBoY285WElqSW5IREl2Umo1VFMyYzlnQi9FekpPcGgzeUk4U1VZODNqdXNlSUFqVUtocFp0VTA0Qy9lYUVwSFgKVk9hdzQ3Qi9iTXlhY3Era3NLUHZ5WnN2eS9qeDZtbVQ5TERnUTZ1MUdtNW1SSmdjMVRvY29UMVpxbVdJc3pZYQpRY25VVnpDVGdFa3V3dkpyQkJjQ09tdllqWitrdEtzdzMzekhDTVo0SWRLaFk3c3hzR0RVOThUektob3JsSFlECmlMRWR3RThmeCtRUkREbVByRDZqbDlYTDFESUZVRnhjY3hlVlNLQUpyb29FVEg3T2dObFNVYVlnYjhRNmhnVEkKUVRaM1ZEYlZlQmR0OWdUMksxQXhxNkNnSzZ2ZURkSm5HNWZpM3dJREFRQUJBb0lCQUFhRkpoc0FiWXU1OXFBZgpUZmpaTTZwaWhoQm5XRzQyWmhyMzVyZGYzREUvSjZ4Vnp3M05JNkZzU3Z4dkpaSzVNRXBiVDMvVURyWDBZL08vClJXSUdWOTJBclYrb3l6elNzMFNrQWJ2S29iVjRMNjNWY2RqMkFuSC9FTW8zM0lmbDIzaDY4VnRSTkNTNU4wYXUKNDc3NXFaTnhvK0xzNHo0dVdibmdsa3AzRkduL29iRDFpcGdoTzVTbkxQOWhpUDRjM3psNnE0T3dibXk0d0M0MApRcHlKTk9LcVVRa2hoVzhlREV4aDlGNW9GYUtMWmxLb2RaejN0c3llUHlPa3dFdHNBSVRoMG13S1RhZlNxdWlOCjBWUDlBMVZWTzlUMXpzZURMK0FZdFBNN2lZekRUN0tSbmNsUE43NmNrTEp5aytpQUlsejFQTzB2QWFEb2p4NmoKUGVpd0d0a0NnWUVBM3dkY1BRcHNSTVdnSWlMVU1xTVdEb0E0RlFxb2lOQmc1SVAxb1RYa1Rhc01ZSVM4NUQ3ZwpxakZXR0RLSnNSZ3I4eCtvSFoxa1BuWThPdlVBTWw1REwxWnZTcC9RRHpCYktlc1lXeU96MFNRWjZsMHZmbjNUClFNNllOektGWm5ZcXlGTXIxTUpNQStMV3hsbUxVNTRhbHdXQUUrWEErMWsvOFk1LzY1UkdIbHNDZ1lFQTB6TjgKbFBuTTdkVE1DMSsvK2V2K0JPQ05iWFdWY1F5OXdFVEw5WUVmZ0VCaDVyc0xQYmR5eWtBazJuNXBTMmhHMDVYNAowcDRzaHAvbmZRaEJaRXovVDlBWVppemtPUGhRL1RhRkY4aFF4VDVkeTVlVFc4OURBa3Jtb3EzbWpBTFdiWjVtCnRhbW1QelgyTTArMzA1bnc0L2d1N2lSalpPVkpvMUVxOVdHdy9NMENnWUVBcUt2TVdtMnpqQjlhQi9jSFBIU1MKamN6eW5SYytkcG9CYlZGUFJ1aVhEUlk3ZWhOcE44VkY2L0Q4QjdqUTRacENRdERDT0FOOGVMQzZ4R3ZlQmptNwozZVVrcmU5SFR2Vm5QTUNMM3dHVlVLcFkzeUl5ZFJ0NzFSdHBpdlV1WmxzZjUvamV3VDFnZDkvcVJvQWFHdGNMCmpRT2Y1V3h4RXFaZzhiS3ZrOTdEV05rQ2dZRUFoaHMxU2l0c0VzQ0NaOTUrdWlVOWdMOU5UbW5SWUoxa2g0ZW4Kc3RZd3VIRXBPU2NmdGlxY093eUwyaWxXbHNrNTMvUmtzT2c3QWFqYmhxc05TckVSbFE1Zno5RkZnVjg3bmUxVwppWWxxc0RRdnZxMGFwcnR1b3pBSVR0ZjVnb0h5d2x4SWY2V2ZxSmVOSTN2RkVCbTV1aWZITlBQcUlSRHV0ME04CjhkNzhVU2tDZ1lFQXpIMWlpL0xWYUJxOFducFBaUFVBOHhtRjZkUEtESGR1T0JVQURNNklFUGx0TXdjWm03NmoKQ3V4RXdkVTE1MVhkNUNDRUxzYVU3d1d0Tjd2TzEvRk9kZTljUStSR2NOd3JCRTR1MWxEMkpseUltOUROZ1BiaApoNm8rOWJQMHlqdGZzNHAreEZVVm5ab0NGd0hqYks4eUM2Q3pqYVZjK2M2NFhuSk1VUmlRNnZJPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\n"
  },
  "addons": {
    "keyValueList": [
      {
        "key": "cluster-autoscaler",
        "value": "{\"installOptionsTemplate\":{},\"name\":\"cluster-autoscaler\",\"targetVersion\":\"2.0.0\",\"version\":\"2.0.0\"}"
      },
      {
        "key": "ibm-storage-operator",
        "value": "{\"healthStatus\":\"Enabling\",\"installOptionsTemplate\":{},\"name\":\"ibm-storage-operator\",\"targetVersion\":\"1.0\",\"version\":\"1.0\"}"
      },
      {
        "key": "vpc-block-csi-driver",
        "value": "{\"allowed_upgrade_versions\":[\"5.2\"],\"healthStatus\":\"Enabling\",\"installOptionsTemplate\":{},\"name\":\"vpc-block-csi-driver\",\"targetVersion\":\"5.1\",\"version\":\"5.1\"}"
      }
    ]
  },
  "status": "Active",
  "createdTime": "2026-10-07T11:06:11Z",
  "keyValueList": [
    {
      "key": "IId",
      "value": "{NameId:tbn9e2fqrrhumthk6j20,SystemId:db32f95s0learmn856i0}"
    },
    {
      "key": "Version",
      "value": "1.35.9_1546"
    },
    {
      "key": "Network",
      "value": "{VpcIID:{NameId:tblo85b3uaik2dvvkjg2,SystemId:r026-1f28c17f-150a-434e-ae0f-a2a321fd4536},SubnetIIDs:[{NameId:tbp1ba7rn56rirasj2pe,SystemId:02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147}],SecurityGroupIIDs:[{NameId:kube-db32f95s0learmn856i0,SystemId:r026-02b8e88f-2be0-4f01-8592-ce88b2a66999}],KeyValueList:[{Key:VpcIID,Value:{NameId:tblo85b3uaik2dvvkjg2,SystemId:r026-1f28c17f-150a-434e-ae0f-a2a321fd4536}},{Key:SubnetIIDs,Value:{NameId:tbp1ba7rn56rirasj2pe,SystemId:02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147}},{Key:SecurityGroupIIDs,Value:{NameId:kube-db32f95s0learmn856i0,SystemId:r026-02b8e88f-2be0-4f01-8592-ce88b2a66999}},{Key:KeyValueList,Value:{Key:VpcIID,Value:{NameId:tblo85b3uaik2dvvkjg2,SystemId:r026-1f28c17f-150a-434e-ae0f-a2a321fd4536}}; {Key:SubnetIIDs,Value:{NameId:tbp1ba7rn56rirasj2pe,SystemId:02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147}}; {Key:SecurityGroupIIDs,Value:{NameId:kube-db32f95s0learmn856i0,SystemId:r026-02b8e88f-2be0-4f01-8592-ce88b2a66999}}}]}"
    },
    {
      "key": "NodeGroupList",
      "value": "{IId:{NameId:workers1,SystemId:db32f95s0learmn856i0-c67c62f},ImageIID:{NameId:Not visible in IBM,SystemId:Not visible in IBM},VMSpecName:cxf.4x8,RootDiskType:Not visible in IBM,RootDiskSize:Not visible in IBM,KeyPairIID:{NameId:Not visible in IBM,SystemId:Not visible in IBM},OnAutoScaling:false,DesiredNodeSize:2,MinNodeSize:0,MaxNodeSize:0,AutoScalingInfoAvailable:false,Status:Active,Nodes:[{NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-00000168},{NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-000002da}],KeyValueList:[{Key:IId,Value:{NameId:workers1,SystemId:db32f95s0learmn856i0-c67c62f}},{Key:ImageIID,Value:{NameId:Not visible in IBM,SystemId:Not visible in IBM}},{Key:VMSpecName,Value:cxf.4x8},{Key:RootDiskType,Value:Not visible in IBM},{Key:RootDiskSize,Value:Not visible in IBM},{Key:KeyPairIID,Value:{NameId:Not visible in IBM,SystemId:Not visible in IBM}},{Key:OnAutoScaling,Value:false},{Key:DesiredNodeSize,Value:2},{Key:MinNodeSize,Value:0},{Key:MaxNodeSize,Value:0},{Key:Status,Value:Active},{Key:Nodes,Value:{NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-00000168}; {NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-000002da}}]}"
    },
    {
      "key": "AccessInfo",
      "value": "{Endpoint:https://c105.au-syd.containers.cloud.ibm.com:32343,Kubeconfig:apiVersion: v1\\nclusters:\\n- cluster:\\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURRRENDQWlpZ0F3SUJBZ0lVZnp2VEoxdkpXQVVoOU93cnk0T0J2VHdqNExFd0RRWUpLb1pJaHZjTkFRRUwKQlFBd09ERTJNRFFHQTFVRUF4TXRaR0l6TW1ZNU5YTXdiR1ZoY20xdU9EVTJhVEF0YTNWaVpYSnVaWFJsY3kxagpZUzB4TnpreE16Y3hNalF5TUI0WERUSTJNVEF3TnpFeE1ETXdNRm9YRFRNMk1UQXdOREV4TURNd01Gb3dPREUyCk1EUUdBMVVFQXhNdFpHSXpNbVk1TlhNd2JHVmhjbTF1T0RVMmFUQXRhM1ZpWlhKdVpYUmxjeTFqWVMweE56a3gKTXpjeE1qUXlNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQThubGY3V3FQeFlhNgphRG0rc0VCWS9qSVBiNlFXb2RBY09zTEdTNDdRT24xTEN6TWlpRTJSZzdtemtoQmNnamNkUWRTQ1NCVTF3RWFQCkN1N01oNCtJTWlsUGxYdzRmaG13RHIwbVlpR01qdFVhN3UrcWtNYmlsSTU5cUpoYXVFMFBVYXV4ZWpEZUxzb3cKaXpuellXaVZ4MDFlcG9PQ3RHUDEzU3oxWmFzSjNBbXl2Z1B3T1BmTnk2U2hESGRtWUUrWGJEZW5qR2tOMWwxUwovVGZEcjk3bU5WREdDNHhUSU1Kcjl3S3htdmIyYnV0bksrb3NaUlNFRkpoRmFDYmRlQXNkb05mVEdlemFiLzNxCmQ0S2w0d1hkdk1wcnZGT3gyWS91UzlXSDlJZ21GSUc4N254RzE4cXNkV2E0OUxPSGFBT1Rrczl4Tmg4ZEhUNk0KWDFXRlNMRU1IUUlEQVFBQm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FZWXdEd1lEVlIwVEFRSC9CQVV3QXdFQgovekFkQmdOVkhRNEVGZ1FVWGtCVlU5QmZpczc2aHNtZ3hVeHRERVU0ZjZrd0RRWUpLb1pJaHZjTkFRRUxCUUFECmdnRUJBS3RrcnppcVdSNVVpTk5uOW1kRmN2L1haRFJlYkEraDMwVHFGTXA5UFRSM2dFVWhmclJqUVlJMGJWWUIKOEgxZUdQdStoV0lvUElzVUJvYTJDM3g2c1dJQWN6c0tUWm9lQ0xJdXJ4YWZlVXVjdEpmaHFVRHhiN3FxSFl1SAoxanp1TmdzZ2pDZGliSFpUeGFLSURWbTF4TW0vVWpQMW05TXc1eWUvcXhETElEaTRFaGtSTGQyc0xsdDZoTTQxCks2YlRYMjUydFFIbDA0Y05ISStPY1kyOVhSVnpvVW82VzNVdDEzZnEzMTBGbGNJZm0zQndUcWNDL2VYZFFCdkQKRFhJbjh1RXFBVGdHanpGM1BocDVVVHlSd2ZJNnUrWnFSdXlPNXFveW9jVC9uTUQxbUgyQW85UUNRQmhOL1JHUQp5WnQrcjJXSXdiVjFGaE5NcUl5MWk2SXBIL009Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\\n    server: https://c105.au-syd.containers.cloud.ibm.com:32343\\n  name: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0\\ncontexts:\\n- context:\\n    cluster: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0\\n    namespace: default\\n    user: admin/db32f95s0learmn856i0\\n  name: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0/admin\\ncurrent-context: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0/admin\\nkind: Config\\nusers:\\n- name: admin/db32f95s0learmn856i0\\n  user:\\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURzRENDQXBpZ0F3SUJBZ0lSQVBDUk1tUGZnS1hMLytyRG5QU1NhY1l3RFFZSktvWklodmNOQVFFTEJRQXcKT0RFMk1EUUdBMVVFQXhNdFpHSXpNbVk1TlhNd2JHVmhjbTF1T0RVMmFUQXRhM1ZpWlhKdVpYUmxjeTFqWVMweApOemt4TXpjeE1qUXlNQjRYRFRJMk1UQXdOekV4TVRrMU4xb1hEVEk0TVRBd056RXhNVGsxTjFvd2F6RUxNQWtHCkExVUVCaE1DVlZNeEZqQVVCZ05WQkFnVERWTmhiaUJHY21GdVkybHpZMjh4Q3pBSkJnTlZCQWNUQWtOQk1STXcKRVFZRFZRUUtFd3BwWW0wdFlXUnRhVzV6TVNJd0lBWURWUVFEREJsTGRXSmxRMlZ5ZENOSlFrMXBaQzAyT1RRdwpNREZDVHpkWE1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBdC8vc0J3cFA1bDhpClNzM3RiSkFZZi9jN0U5QlM2N0s1YXRGRFJlcEYzUVZJUzBUNTBoY285WElqSW5IREl2Umo1VFMyYzlnQi9FekoKT3BoM3lJOFNVWTgzanVzZUlBalVLaHBadFUwNEMvZWFFcEhYVk9hdzQ3Qi9iTXlhY3Era3NLUHZ5WnN2eS9qeAo2bW1UOUxEZ1E2dTFHbTVtUkpnYzFUb2NvVDFacW1XSXN6WWFRY25VVnpDVGdFa3V3dkpyQkJjQ09tdllqWitrCnRLc3czM3pIQ01aNElkS2hZN3N4c0dEVTk4VHpLaG9ybEhZRGlMRWR3RThmeCtRUkREbVByRDZqbDlYTDFESUYKVUZ4Y2N4ZVZTS0FKcm9vRVRIN09nTmxTVWFZZ2I4UTZoZ1RJUVRaM1ZEYlZlQmR0OWdUMksxQXhxNkNnSzZ2ZQpEZEpuRzVmaTN3SURBUUFCbzRHQk1IOHdEZ1lEVlIwUEFRSC9CQVFEQWdLRU1CMEdBMVVkSlFRV01CUUdDQ3NHCkFRVUZCd01DQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRkY1QVZWUFEKWDRyTytvYkpvTVZNYlF4Rk9IK3BNQjhHQTFVZEVRUVlNQmFCRkdoaGJtbDZZVzVuTnpkQVoyMWhhV3d1WTI5dApNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUF3OUVkQkJWWmhTUHMrWXJqUDhoNUJVVWxWTzBzZnZhR1djMlVKCi9VTlVHVUJTZm1FTXFSanZHYythV1JtbjdnV1VBMW92azJ4NE10NU5WTWpNYkVOVEMwWk1Ja1NTdWttR3hOYW4KSkE2RGdXMXpOdlBzazIvMDNnQU1DbUJtd1RheC93V0lwWC9QQTU2RXk4bEhIRVZpdWttNS9uU0lGYUIxZ0ZNbQpvczB2OE1WdFc4YVVlRGEySTVrajRaQlZicytOc3hXU0RCK2RoL2NBcUlGVS9sQzgwakRjeFMxbkl0K2o3UnN3CjVQeFJHcEllbUM1NjdDNEMxdUpZa2tLazJ2WStvZ05GczlQVzlGTXBjdjdPUEtMc3ZNb0lyNER4L1lvK3V4dTEKS3BuRm40M21tUVk5TERMelliTCtyK2R2Nm9ZVVFEZDZqS0VQSS9CK293d3BzcGxxCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBdC8vc0J3cFA1bDhpU3MzdGJKQVlmL2M3RTlCUzY3SzVhdEZEUmVwRjNRVklTMFQ1CjBoY285WElqSW5IREl2Umo1VFMyYzlnQi9FekpPcGgzeUk4U1VZODNqdXNlSUFqVUtocFp0VTA0Qy9lYUVwSFgKVk9hdzQ3Qi9iTXlhY3Era3NLUHZ5WnN2eS9qeDZtbVQ5TERnUTZ1MUdtNW1SSmdjMVRvY29UMVpxbVdJc3pZYQpRY25VVnpDVGdFa3V3dkpyQkJjQ09tdllqWitrdEtzdzMzekhDTVo0SWRLaFk3c3hzR0RVOThUektob3JsSFlECmlMRWR3RThmeCtRUkREbVByRDZqbDlYTDFESUZVRnhjY3hlVlNLQUpyb29FVEg3T2dObFNVYVlnYjhRNmhnVEkKUVRaM1ZEYlZlQmR0OWdUMksxQXhxNkNnSzZ2ZURkSm5HNWZpM3dJREFRQUJBb0lCQUFhRkpoc0FiWXU1OXFBZgpUZmpaTTZwaWhoQm5XRzQyWmhyMzVyZGYzREUvSjZ4Vnp3M05JNkZzU3Z4dkpaSzVNRXBiVDMvVURyWDBZL08vClJXSUdWOTJBclYrb3l6elNzMFNrQWJ2S29iVjRMNjNWY2RqMkFuSC9FTW8zM0lmbDIzaDY4VnRSTkNTNU4wYXUKNDc3NXFaTnhvK0xzNHo0dVdibmdsa3AzRkduL29iRDFpcGdoTzVTbkxQOWhpUDRjM3psNnE0T3dibXk0d0M0MApRcHlKTk9LcVVRa2hoVzhlREV4aDlGNW9GYUtMWmxLb2RaejN0c3llUHlPa3dFdHNBSVRoMG13S1RhZlNxdWlOCjBWUDlBMVZWTzlUMXpzZURMK0FZdFBNN2lZekRUN0tSbmNsUE43NmNrTEp5aytpQUlsejFQTzB2QWFEb2p4NmoKUGVpd0d0a0NnWUVBM3dkY1BRcHNSTVdnSWlMVU1xTVdEb0E0RlFxb2lOQmc1SVAxb1RYa1Rhc01ZSVM4NUQ3ZwpxakZXR0RLSnNSZ3I4eCtvSFoxa1BuWThPdlVBTWw1REwxWnZTcC9RRHpCYktlc1lXeU96MFNRWjZsMHZmbjNUClFNNllOektGWm5ZcXlGTXIxTUpNQStMV3hsbUxVNTRhbHdXQUUrWEErMWsvOFk1LzY1UkdIbHNDZ1lFQTB6TjgKbFBuTTdkVE1DMSsvK2V2K0JPQ05iWFdWY1F5OXdFVEw5WUVmZ0VCaDVyc0xQYmR5eWtBazJuNXBTMmhHMDVYNAowcDRzaHAvbmZRaEJaRXovVDlBWVppemtPUGhRL1RhRkY4aFF4VDVkeTVlVFc4OURBa3Jtb3EzbWpBTFdiWjVtCnRhbW1QelgyTTArMzA1bnc0L2d1N2lSalpPVkpvMUVxOVdHdy9NMENnWUVBcUt2TVdtMnpqQjlhQi9jSFBIU1MKamN6eW5SYytkcG9CYlZGUFJ1aVhEUlk3ZWhOcE44VkY2L0Q4QjdqUTRacENRdERDT0FOOGVMQzZ4R3ZlQmptNwozZVVrcmU5SFR2Vm5QTUNMM3dHVlVLcFkzeUl5ZFJ0NzFSdHBpdlV1WmxzZjUvamV3VDFnZDkvcVJvQWFHdGNMCmpRT2Y1V3h4RXFaZzhiS3ZrOTdEV05rQ2dZRUFoaHMxU2l0c0VzQ0NaOTUrdWlVOWdMOU5UbW5SWUoxa2g0ZW4Kc3RZd3VIRXBPU2NmdGlxY093eUwyaWxXbHNrNTMvUmtzT2c3QWFqYmhxc05TckVSbFE1Zno5RkZnVjg3bmUxVwppWWxxc0RRdnZxMGFwcnR1b3pBSVR0ZjVnb0h5d2x4SWY2V2ZxSmVOSTN2RkVCbTV1aWZITlBQcUlSRHV0ME04CjhkNzhVU2tDZ1lFQXpIMWlpL0xWYUJxOFducFBaUFVBOHhtRjZkUEtESGR1T0JVQURNNklFUGx0TXdjWm03NmoKQ3V4RXdkVTE1MVhkNUNDRUxzYVU3d1d0Tjd2TzEvRk9kZTljUStSR2NOd3JCRTR1MWxEMkpseUltOUROZ1BiaApoNm8rOWJQMHlqdGZzNHAreEZVVm5ab0NGd0hqYks4eUM2Q3pqYVZjK2M2NFhuSk1VUmlRNnZJPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\\n}"
    },
    {
      "key": "Addons",
      "value": "{KeyValueList:[{Key:cluster-autoscaler,Value:{\\installOptionsTemplate\\:{},\\name\\:\\cluster-autoscaler\\,\\targetVersion\\:\\2.0.0\\,\\version\\:\\2.0.0\\}},{Key:ibm-storage-operator,Value:{\\healthStatus\\:\\Enabling\\,\\installOptionsTemplate\\:{},\\name\\:\\ibm-storage-operator\\,\\targetVersion\\:\\1.0\\,\\version\\:\\1.0\\}},{Key:vpc-block-csi-driver,Value:{\\allowed_upgrade_versions\\:[\\5.2\\],\\healthStatus\\:\\Enabling\\,\\installOptionsTemplate\\:{},\\name\\:\\vpc-block-csi-driver\\,\\targetVersion\\:\\5.1\\,\\version\\:\\5.1\\}}]}"
    },
    {
      "key": "Status",
      "value": "Active"
    },
    {
      "key": "CreatedTime",
      "value": "2026-10-07T11:06:11Z"
    },
    {
      "key": "TagList",
      "value": "{Key:cb-spider-pmks-autoscaler-status,Value:deploying}"
    }
  ],
  "cspResourceName": "tbn9e2fqrrhumthk6j20",
  "cspResourceId": "db32f95s0learmn856i0",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbn9e2fqrrhumthk6j20",
      "SystemId": "db32f95s0learmn856i0"
    },
    "Version": "1.35.9_1546",
    "Network": {
      "VpcIID": {
        "NameId": "tblo85b3uaik2dvvkjg2",
        "SystemId": "r026-1f28c17f-150a-434e-ae0f-a2a321fd4536"
      },
      "SubnetIIDs": [
        {
          "NameId": "tbp1ba7rn56rirasj2pe",
          "SystemId": "02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "kube-db32f95s0learmn856i0",
          "SystemId": "r026-02b8e88f-2be0-4f01-8592-ce88b2a66999"
        }
      ],
      "KeyValueList": [
        {
          "key": "VpcIID",
          "value": "{NameId:tblo85b3uaik2dvvkjg2,SystemId:r026-1f28c17f-150a-434e-ae0f-a2a321fd4536}"
        },
        {
          "key": "SubnetIIDs",
          "value": "{NameId:tbp1ba7rn56rirasj2pe,SystemId:02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147}"
        },
        {
          "key": "SecurityGroupIIDs",
          "value": "{NameId:kube-db32f95s0learmn856i0,SystemId:r026-02b8e88f-2be0-4f01-8592-ce88b2a66999}"
        },
        {
          "key": "KeyValueList",
          "value": "{Key:VpcIID,Value:{NameId:tblo85b3uaik2dvvkjg2,SystemId:r026-1f28c17f-150a-434e-ae0f-a2a321fd4536}}; {Key:SubnetIIDs,Value:{NameId:tbp1ba7rn56rirasj2pe,SystemId:02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147}}; {Key:SecurityGroupIIDs,Value:{NameId:kube-db32f95s0learmn856i0,SystemId:r026-02b8e88f-2be0-4f01-8592-ce88b2a66999}}"
        }
      ]
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "db32f95s0learmn856i0-c67c62f"
        },
        "ImageIID": {
          "NameId": "Not visible in IBM",
          "SystemId": "Not visible in IBM"
        },
        "VMSpecName": "cxf.4x8",
        "RootDiskType": "Not visible in IBM",
        "RootDiskSize": "Not visible in IBM",
        "KeyPairIID": {
          "NameId": "",
          "SystemId": "Not visible in IBM"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "Not visible in IBM",
            "SystemId": "kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-00000168"
          },
          {
            "NameId": "Not visible in IBM",
            "SystemId": "kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-000002da"
          }
        ],
        "KeyValueList": [
          {
            "key": "IId",
            "value": "{NameId:workers1,SystemId:db32f95s0learmn856i0-c67c62f}"
          },
          {
            "key": "ImageIID",
            "value": "{NameId:Not visible in IBM,SystemId:Not visible in IBM}"
          },
          {
            "key": "VMSpecName",
            "value": "cxf.4x8"
          },
          {
            "key": "RootDiskType",
            "value": "Not visible in IBM"
          },
          {
            "key": "RootDiskSize",
            "value": "Not visible in IBM"
          },
          {
            "key": "KeyPairIID",
            "value": "{NameId:Not visible in IBM,SystemId:Not visible in IBM}"
          },
          {
            "key": "OnAutoScaling",
            "value": "false"
          },
          {
            "key": "DesiredNodeSize",
            "value": "2"
          },
          {
            "key": "MinNodeSize",
            "value": "0"
          },
          {
            "key": "MaxNodeSize",
            "value": "0"
          },
          {
            "key": "Status",
            "value": "Active"
          },
          {
            "key": "Nodes",
            "value": "{NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-00000168}; {NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-000002da}"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://c105.au-syd.containers.cloud.ibm.com:32343",
      "Kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURRRENDQWlpZ0F3SUJBZ0lVZnp2VEoxdkpXQVVoOU93cnk0T0J2VHdqNExFd0RRWUpLb1pJaHZjTkFRRUwKQlFBd09ERTJNRFFHQTFVRUF4TXRaR0l6TW1ZNU5YTXdiR1ZoY20xdU9EVTJhVEF0YTNWaVpYSnVaWFJsY3kxagpZUzB4TnpreE16Y3hNalF5TUI0WERUSTJNVEF3TnpFeE1ETXdNRm9YRFRNMk1UQXdOREV4TURNd01Gb3dPREUyCk1EUUdBMVVFQXhNdFpHSXpNbVk1TlhNd2JHVmhjbTF1T0RVMmFUQXRhM1ZpWlhKdVpYUmxjeTFqWVMweE56a3gKTXpjeE1qUXlNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQThubGY3V3FQeFlhNgphRG0rc0VCWS9qSVBiNlFXb2RBY09zTEdTNDdRT24xTEN6TWlpRTJSZzdtemtoQmNnamNkUWRTQ1NCVTF3RWFQCkN1N01oNCtJTWlsUGxYdzRmaG13RHIwbVlpR01qdFVhN3UrcWtNYmlsSTU5cUpoYXVFMFBVYXV4ZWpEZUxzb3cKaXpuellXaVZ4MDFlcG9PQ3RHUDEzU3oxWmFzSjNBbXl2Z1B3T1BmTnk2U2hESGRtWUUrWGJEZW5qR2tOMWwxUwovVGZEcjk3bU5WREdDNHhUSU1Kcjl3S3htdmIyYnV0bksrb3NaUlNFRkpoRmFDYmRlQXNkb05mVEdlemFiLzNxCmQ0S2w0d1hkdk1wcnZGT3gyWS91UzlXSDlJZ21GSUc4N254RzE4cXNkV2E0OUxPSGFBT1Rrczl4Tmg4ZEhUNk0KWDFXRlNMRU1IUUlEQVFBQm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FZWXdEd1lEVlIwVEFRSC9CQVV3QXdFQgovekFkQmdOVkhRNEVGZ1FVWGtCVlU5QmZpczc2aHNtZ3hVeHRERVU0ZjZrd0RRWUpLb1pJaHZjTkFRRUxCUUFECmdnRUJBS3RrcnppcVdSNVVpTk5uOW1kRmN2L1haRFJlYkEraDMwVHFGTXA5UFRSM2dFVWhmclJqUVlJMGJWWUIKOEgxZUdQdStoV0lvUElzVUJvYTJDM3g2c1dJQWN6c0tUWm9lQ0xJdXJ4YWZlVXVjdEpmaHFVRHhiN3FxSFl1SAoxanp1TmdzZ2pDZGliSFpUeGFLSURWbTF4TW0vVWpQMW05TXc1eWUvcXhETElEaTRFaGtSTGQyc0xsdDZoTTQxCks2YlRYMjUydFFIbDA0Y05ISStPY1kyOVhSVnpvVW82VzNVdDEzZnEzMTBGbGNJZm0zQndUcWNDL2VYZFFCdkQKRFhJbjh1RXFBVGdHanpGM1BocDVVVHlSd2ZJNnUrWnFSdXlPNXFveW9jVC9uTUQxbUgyQW85UUNRQmhOL1JHUQp5WnQrcjJXSXdiVjFGaE5NcUl5MWk2SXBIL009Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://c105.au-syd.containers.cloud.ibm.com:32343\n  name: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0\ncontexts:\n- context:\n    cluster: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0\n    namespace: default\n    user: admin/db32f95s0learmn856i0\n  name: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0/admin\ncurrent-context: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0/admin\nkind: Config\nusers:\n- name: admin/db32f95s0learmn856i0\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURzRENDQXBpZ0F3SUJBZ0lSQVBDUk1tUGZnS1hMLytyRG5QU1NhY1l3RFFZSktvWklodmNOQVFFTEJRQXcKT0RFMk1EUUdBMVVFQXhNdFpHSXpNbVk1TlhNd2JHVmhjbTF1T0RVMmFUQXRhM1ZpWlhKdVpYUmxjeTFqWVMweApOemt4TXpjeE1qUXlNQjRYRFRJMk1UQXdOekV4TVRrMU4xb1hEVEk0TVRBd056RXhNVGsxTjFvd2F6RUxNQWtHCkExVUVCaE1DVlZNeEZqQVVCZ05WQkFnVERWTmhiaUJHY21GdVkybHpZMjh4Q3pBSkJnTlZCQWNUQWtOQk1STXcKRVFZRFZRUUtFd3BwWW0wdFlXUnRhVzV6TVNJd0lBWURWUVFEREJsTGRXSmxRMlZ5ZENOSlFrMXBaQzAyT1RRdwpNREZDVHpkWE1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBdC8vc0J3cFA1bDhpClNzM3RiSkFZZi9jN0U5QlM2N0s1YXRGRFJlcEYzUVZJUzBUNTBoY285WElqSW5IREl2Umo1VFMyYzlnQi9FekoKT3BoM3lJOFNVWTgzanVzZUlBalVLaHBadFUwNEMvZWFFcEhYVk9hdzQ3Qi9iTXlhY3Era3NLUHZ5WnN2eS9qeAo2bW1UOUxEZ1E2dTFHbTVtUkpnYzFUb2NvVDFacW1XSXN6WWFRY25VVnpDVGdFa3V3dkpyQkJjQ09tdllqWitrCnRLc3czM3pIQ01aNElkS2hZN3N4c0dEVTk4VHpLaG9ybEhZRGlMRWR3RThmeCtRUkREbVByRDZqbDlYTDFESUYKVUZ4Y2N4ZVZTS0FKcm9vRVRIN09nTmxTVWFZZ2I4UTZoZ1RJUVRaM1ZEYlZlQmR0OWdUMksxQXhxNkNnSzZ2ZQpEZEpuRzVmaTN3SURBUUFCbzRHQk1IOHdEZ1lEVlIwUEFRSC9CQVFEQWdLRU1CMEdBMVVkSlFRV01CUUdDQ3NHCkFRVUZCd01DQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRkY1QVZWUFEKWDRyTytvYkpvTVZNYlF4Rk9IK3BNQjhHQTFVZEVRUVlNQmFCRkdoaGJtbDZZVzVuTnpkQVoyMWhhV3d1WTI5dApNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUF3OUVkQkJWWmhTUHMrWXJqUDhoNUJVVWxWTzBzZnZhR1djMlVKCi9VTlVHVUJTZm1FTXFSanZHYythV1JtbjdnV1VBMW92azJ4NE10NU5WTWpNYkVOVEMwWk1Ja1NTdWttR3hOYW4KSkE2RGdXMXpOdlBzazIvMDNnQU1DbUJtd1RheC93V0lwWC9QQTU2RXk4bEhIRVZpdWttNS9uU0lGYUIxZ0ZNbQpvczB2OE1WdFc4YVVlRGEySTVrajRaQlZicytOc3hXU0RCK2RoL2NBcUlGVS9sQzgwakRjeFMxbkl0K2o3UnN3CjVQeFJHcEllbUM1NjdDNEMxdUpZa2tLazJ2WStvZ05GczlQVzlGTXBjdjdPUEtMc3ZNb0lyNER4L1lvK3V4dTEKS3BuRm40M21tUVk5TERMelliTCtyK2R2Nm9ZVVFEZDZqS0VQSS9CK293d3BzcGxxCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBdC8vc0J3cFA1bDhpU3MzdGJKQVlmL2M3RTlCUzY3SzVhdEZEUmVwRjNRVklTMFQ1CjBoY285WElqSW5IREl2Umo1VFMyYzlnQi9FekpPcGgzeUk4U1VZODNqdXNlSUFqVUtocFp0VTA0Qy9lYUVwSFgKVk9hdzQ3Qi9iTXlhY3Era3NLUHZ5WnN2eS9qeDZtbVQ5TERnUTZ1MUdtNW1SSmdjMVRvY29UMVpxbVdJc3pZYQpRY25VVnpDVGdFa3V3dkpyQkJjQ09tdllqWitrdEtzdzMzekhDTVo0SWRLaFk3c3hzR0RVOThUektob3JsSFlECmlMRWR3RThmeCtRUkREbVByRDZqbDlYTDFESUZVRnhjY3hlVlNLQUpyb29FVEg3T2dObFNVYVlnYjhRNmhnVEkKUVRaM1ZEYlZlQmR0OWdUMksxQXhxNkNnSzZ2ZURkSm5HNWZpM3dJREFRQUJBb0lCQUFhRkpoc0FiWXU1OXFBZgpUZmpaTTZwaWhoQm5XRzQyWmhyMzVyZGYzREUvSjZ4Vnp3M05JNkZzU3Z4dkpaSzVNRXBiVDMvVURyWDBZL08vClJXSUdWOTJBclYrb3l6elNzMFNrQWJ2S29iVjRMNjNWY2RqMkFuSC9FTW8zM0lmbDIzaDY4VnRSTkNTNU4wYXUKNDc3NXFaTnhvK0xzNHo0dVdibmdsa3AzRkduL29iRDFpcGdoTzVTbkxQOWhpUDRjM3psNnE0T3dibXk0d0M0MApRcHlKTk9LcVVRa2hoVzhlREV4aDlGNW9GYUtMWmxLb2RaejN0c3llUHlPa3dFdHNBSVRoMG13S1RhZlNxdWlOCjBWUDlBMVZWTzlUMXpzZURMK0FZdFBNN2lZekRUN0tSbmNsUE43NmNrTEp5aytpQUlsejFQTzB2QWFEb2p4NmoKUGVpd0d0a0NnWUVBM3dkY1BRcHNSTVdnSWlMVU1xTVdEb0E0RlFxb2lOQmc1SVAxb1RYa1Rhc01ZSVM4NUQ3ZwpxakZXR0RLSnNSZ3I4eCtvSFoxa1BuWThPdlVBTWw1REwxWnZTcC9RRHpCYktlc1lXeU96MFNRWjZsMHZmbjNUClFNNllOektGWm5ZcXlGTXIxTUpNQStMV3hsbUxVNTRhbHdXQUUrWEErMWsvOFk1LzY1UkdIbHNDZ1lFQTB6TjgKbFBuTTdkVE1DMSsvK2V2K0JPQ05iWFdWY1F5OXdFVEw5WUVmZ0VCaDVyc0xQYmR5eWtBazJuNXBTMmhHMDVYNAowcDRzaHAvbmZRaEJaRXovVDlBWVppemtPUGhRL1RhRkY4aFF4VDVkeTVlVFc4OURBa3Jtb3EzbWpBTFdiWjVtCnRhbW1QelgyTTArMzA1bnc0L2d1N2lSalpPVkpvMUVxOVdHdy9NMENnWUVBcUt2TVdtMnpqQjlhQi9jSFBIU1MKamN6eW5SYytkcG9CYlZGUFJ1aVhEUlk3ZWhOcE44VkY2L0Q4QjdqUTRacENRdERDT0FOOGVMQzZ4R3ZlQmptNwozZVVrcmU5SFR2Vm5QTUNMM3dHVlVLcFkzeUl5ZFJ0NzFSdHBpdlV1WmxzZjUvamV3VDFnZDkvcVJvQWFHdGNMCmpRT2Y1V3h4RXFaZzhiS3ZrOTdEV05rQ2dZRUFoaHMxU2l0c0VzQ0NaOTUrdWlVOWdMOU5UbW5SWUoxa2g0ZW4Kc3RZd3VIRXBPU2NmdGlxY093eUwyaWxXbHNrNTMvUmtzT2c3QWFqYmhxc05TckVSbFE1Zno5RkZnVjg3bmUxVwppWWxxc0RRdnZxMGFwcnR1b3pBSVR0ZjVnb0h5d2x4SWY2V2ZxSmVOSTN2RkVCbTV1aWZITlBQcUlSRHV0ME04CjhkNzhVU2tDZ1lFQXpIMWlpL0xWYUJxOFducFBaUFVBOHhtRjZkUEtESGR1T0JVQURNNklFUGx0TXdjWm03NmoKQ3V4RXdkVTE1MVhkNUNDRUxzYVU3d1d0Tjd2TzEvRk9kZTljUStSR2NOd3JCRTR1MWxEMkpseUltOUROZ1BiaApoNm8rOWJQMHlqdGZzNHAreEZVVm5ab0NGd0hqYks4eUM2Q3pqYVZjK2M2NFhuSk1VUmlRNnZJPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\n"
    },
    "Addons": {
      "KeyValueList": [
        {
          "key": "cluster-autoscaler",
          "value": "{\"installOptionsTemplate\":{},\"name\":\"cluster-autoscaler\",\"targetVersion\":\"2.0.0\",\"version\":\"2.0.0\"}"
        },
        {
          "key": "ibm-storage-operator",
          "value": "{\"healthStatus\":\"Enabling\",\"installOptionsTemplate\":{},\"name\":\"ibm-storage-operator\",\"targetVersion\":\"1.0\",\"version\":\"1.0\"}"
        },
        {
          "key": "vpc-block-csi-driver",
          "value": "{\"allowed_upgrade_versions\":[\"5.2\"],\"healthStatus\":\"Enabling\",\"installOptionsTemplate\":{},\"name\":\"vpc-block-csi-driver\",\"targetVersion\":\"5.1\",\"version\":\"5.1\"}"
        }
      ]
    },
    "Status": "Active",
    "CreatedTime": "2026-10-07T11:06:11Z",
    "KeyValueList": [
      {
        "key": "IId",
        "value": "{NameId:tbn9e2fqrrhumthk6j20,SystemId:db32f95s0learmn856i0}"
      },
      {
        "key": "Version",
        "value": "1.35.9_1546"
      },
      {
        "key": "Network",
        "value": "{VpcIID:{NameId:tblo85b3uaik2dvvkjg2,SystemId:r026-1f28c17f-150a-434e-ae0f-a2a321fd4536},SubnetIIDs:[{NameId:tbp1ba7rn56rirasj2pe,SystemId:02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147}],SecurityGroupIIDs:[{NameId:kube-db32f95s0learmn856i0,SystemId:r026-02b8e88f-2be0-4f01-8592-ce88b2a66999}],KeyValueList:[{Key:VpcIID,Value:{NameId:tblo85b3uaik2dvvkjg2,SystemId:r026-1f28c17f-150a-434e-ae0f-a2a321fd4536}},{Key:SubnetIIDs,Value:{NameId:tbp1ba7rn56rirasj2pe,SystemId:02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147}},{Key:SecurityGroupIIDs,Value:{NameId:kube-db32f95s0learmn856i0,SystemId:r026-02b8e88f-2be0-4f01-8592-ce88b2a66999}},{Key:KeyValueList,Value:{Key:VpcIID,Value:{NameId:tblo85b3uaik2dvvkjg2,SystemId:r026-1f28c17f-150a-434e-ae0f-a2a321fd4536}}; {Key:SubnetIIDs,Value:{NameId:tbp1ba7rn56rirasj2pe,SystemId:02h7-f7ea2bcc-1ff6-41d8-a8f8-e2acdb019147}}; {Key:SecurityGroupIIDs,Value:{NameId:kube-db32f95s0learmn856i0,SystemId:r026-02b8e88f-2be0-4f01-8592-ce88b2a66999}}}]}"
      },
      {
        "key": "NodeGroupList",
        "value": "{IId:{NameId:workers1,SystemId:db32f95s0learmn856i0-c67c62f},ImageIID:{NameId:Not visible in IBM,SystemId:Not visible in IBM},VMSpecName:cxf.4x8,RootDiskType:Not visible in IBM,RootDiskSize:Not visible in IBM,KeyPairIID:{NameId:Not visible in IBM,SystemId:Not visible in IBM},OnAutoScaling:false,DesiredNodeSize:2,MinNodeSize:0,MaxNodeSize:0,AutoScalingInfoAvailable:false,Status:Active,Nodes:[{NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-00000168},{NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-000002da}],KeyValueList:[{Key:IId,Value:{NameId:workers1,SystemId:db32f95s0learmn856i0-c67c62f}},{Key:ImageIID,Value:{NameId:Not visible in IBM,SystemId:Not visible in IBM}},{Key:VMSpecName,Value:cxf.4x8},{Key:RootDiskType,Value:Not visible in IBM},{Key:RootDiskSize,Value:Not visible in IBM},{Key:KeyPairIID,Value:{NameId:Not visible in IBM,SystemId:Not visible in IBM}},{Key:OnAutoScaling,Value:false},{Key:DesiredNodeSize,Value:2},{Key:MinNodeSize,Value:0},{Key:MaxNodeSize,Value:0},{Key:Status,Value:Active},{Key:Nodes,Value:{NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-00000168}; {NameId:Not visible in IBM,SystemId:kube-db32f95s0learmn856i0-tbn9e2fqrrh-workers-000002da}}]}"
      },
      {
        "key": "AccessInfo",
        "value": "{Endpoint:https://c105.au-syd.containers.cloud.ibm.com:32343,Kubeconfig:apiVersion: v1\\nclusters:\\n- cluster:\\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURRRENDQWlpZ0F3SUJBZ0lVZnp2VEoxdkpXQVVoOU93cnk0T0J2VHdqNExFd0RRWUpLb1pJaHZjTkFRRUwKQlFBd09ERTJNRFFHQTFVRUF4TXRaR0l6TW1ZNU5YTXdiR1ZoY20xdU9EVTJhVEF0YTNWaVpYSnVaWFJsY3kxagpZUzB4TnpreE16Y3hNalF5TUI0WERUSTJNVEF3TnpFeE1ETXdNRm9YRFRNMk1UQXdOREV4TURNd01Gb3dPREUyCk1EUUdBMVVFQXhNdFpHSXpNbVk1TlhNd2JHVmhjbTF1T0RVMmFUQXRhM1ZpWlhKdVpYUmxjeTFqWVMweE56a3gKTXpjeE1qUXlNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQThubGY3V3FQeFlhNgphRG0rc0VCWS9qSVBiNlFXb2RBY09zTEdTNDdRT24xTEN6TWlpRTJSZzdtemtoQmNnamNkUWRTQ1NCVTF3RWFQCkN1N01oNCtJTWlsUGxYdzRmaG13RHIwbVlpR01qdFVhN3UrcWtNYmlsSTU5cUpoYXVFMFBVYXV4ZWpEZUxzb3cKaXpuellXaVZ4MDFlcG9PQ3RHUDEzU3oxWmFzSjNBbXl2Z1B3T1BmTnk2U2hESGRtWUUrWGJEZW5qR2tOMWwxUwovVGZEcjk3bU5WREdDNHhUSU1Kcjl3S3htdmIyYnV0bksrb3NaUlNFRkpoRmFDYmRlQXNkb05mVEdlemFiLzNxCmQ0S2w0d1hkdk1wcnZGT3gyWS91UzlXSDlJZ21GSUc4N254RzE4cXNkV2E0OUxPSGFBT1Rrczl4Tmg4ZEhUNk0KWDFXRlNMRU1IUUlEQVFBQm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FZWXdEd1lEVlIwVEFRSC9CQVV3QXdFQgovekFkQmdOVkhRNEVGZ1FVWGtCVlU5QmZpczc2aHNtZ3hVeHRERVU0ZjZrd0RRWUpLb1pJaHZjTkFRRUxCUUFECmdnRUJBS3RrcnppcVdSNVVpTk5uOW1kRmN2L1haRFJlYkEraDMwVHFGTXA5UFRSM2dFVWhmclJqUVlJMGJWWUIKOEgxZUdQdStoV0lvUElzVUJvYTJDM3g2c1dJQWN6c0tUWm9lQ0xJdXJ4YWZlVXVjdEpmaHFVRHhiN3FxSFl1SAoxanp1TmdzZ2pDZGliSFpUeGFLSURWbTF4TW0vVWpQMW05TXc1eWUvcXhETElEaTRFaGtSTGQyc0xsdDZoTTQxCks2YlRYMjUydFFIbDA0Y05ISStPY1kyOVhSVnpvVW82VzNVdDEzZnEzMTBGbGNJZm0zQndUcWNDL2VYZFFCdkQKRFhJbjh1RXFBVGdHanpGM1BocDVVVHlSd2ZJNnUrWnFSdXlPNXFveW9jVC9uTUQxbUgyQW85UUNRQmhOL1JHUQp5WnQrcjJXSXdiVjFGaE5NcUl5MWk2SXBIL009Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\\n    server: https://c105.au-syd.containers.cloud.ibm.com:32343\\n  name: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0\\ncontexts:\\n- context:\\n    cluster: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0\\n    namespace: default\\n    user: admin/db32f95s0learmn856i0\\n  name: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0/admin\\ncurrent-context: tbn9e2fqrrhumthk6j20/db32f95s0learmn856i0/admin\\nkind: Config\\nusers:\\n- name: admin/db32f95s0learmn856i0\\n  user:\\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURzRENDQXBpZ0F3SUJBZ0lSQVBDUk1tUGZnS1hMLytyRG5QU1NhY1l3RFFZSktvWklodmNOQVFFTEJRQXcKT0RFMk1EUUdBMVVFQXhNdFpHSXpNbVk1TlhNd2JHVmhjbTF1T0RVMmFUQXRhM1ZpWlhKdVpYUmxjeTFqWVMweApOemt4TXpjeE1qUXlNQjRYRFRJMk1UQXdOekV4TVRrMU4xb1hEVEk0TVRBd056RXhNVGsxTjFvd2F6RUxNQWtHCkExVUVCaE1DVlZNeEZqQVVCZ05WQkFnVERWTmhiaUJHY21GdVkybHpZMjh4Q3pBSkJnTlZCQWNUQWtOQk1STXcKRVFZRFZRUUtFd3BwWW0wdFlXUnRhVzV6TVNJd0lBWURWUVFEREJsTGRXSmxRMlZ5ZENOSlFrMXBaQzAyT1RRdwpNREZDVHpkWE1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBdC8vc0J3cFA1bDhpClNzM3RiSkFZZi9jN0U5QlM2N0s1YXRGRFJlcEYzUVZJUzBUNTBoY285WElqSW5IREl2Umo1VFMyYzlnQi9FekoKT3BoM3lJOFNVWTgzanVzZUlBalVLaHBadFUwNEMvZWFFcEhYVk9hdzQ3Qi9iTXlhY3Era3NLUHZ5WnN2eS9qeAo2bW1UOUxEZ1E2dTFHbTVtUkpnYzFUb2NvVDFacW1XSXN6WWFRY25VVnpDVGdFa3V3dkpyQkJjQ09tdllqWitrCnRLc3czM3pIQ01aNElkS2hZN3N4c0dEVTk4VHpLaG9ybEhZRGlMRWR3RThmeCtRUkREbVByRDZqbDlYTDFESUYKVUZ4Y2N4ZVZTS0FKcm9vRVRIN09nTmxTVWFZZ2I4UTZoZ1RJUVRaM1ZEYlZlQmR0OWdUMksxQXhxNkNnSzZ2ZQpEZEpuRzVmaTN3SURBUUFCbzRHQk1IOHdEZ1lEVlIwUEFRSC9CQVFEQWdLRU1CMEdBMVVkSlFRV01CUUdDQ3NHCkFRVUZCd01DQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRkY1QVZWUFEKWDRyTytvYkpvTVZNYlF4Rk9IK3BNQjhHQTFVZEVRUVlNQmFCRkdoaGJtbDZZVzVuTnpkQVoyMWhhV3d1WTI5dApNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUF3OUVkQkJWWmhTUHMrWXJqUDhoNUJVVWxWTzBzZnZhR1djMlVKCi9VTlVHVUJTZm1FTXFSanZHYythV1JtbjdnV1VBMW92azJ4NE10NU5WTWpNYkVOVEMwWk1Ja1NTdWttR3hOYW4KSkE2RGdXMXpOdlBzazIvMDNnQU1DbUJtd1RheC93V0lwWC9QQTU2RXk4bEhIRVZpdWttNS9uU0lGYUIxZ0ZNbQpvczB2OE1WdFc4YVVlRGEySTVrajRaQlZicytOc3hXU0RCK2RoL2NBcUlGVS9sQzgwakRjeFMxbkl0K2o3UnN3CjVQeFJHcEllbUM1NjdDNEMxdUpZa2tLazJ2WStvZ05GczlQVzlGTXBjdjdPUEtMc3ZNb0lyNER4L1lvK3V4dTEKS3BuRm40M21tUVk5TERMelliTCtyK2R2Nm9ZVVFEZDZqS0VQSS9CK293d3BzcGxxCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBdC8vc0J3cFA1bDhpU3MzdGJKQVlmL2M3RTlCUzY3SzVhdEZEUmVwRjNRVklTMFQ1CjBoY285WElqSW5IREl2Umo1VFMyYzlnQi9FekpPcGgzeUk4U1VZODNqdXNlSUFqVUtocFp0VTA0Qy9lYUVwSFgKVk9hdzQ3Qi9iTXlhY3Era3NLUHZ5WnN2eS9qeDZtbVQ5TERnUTZ1MUdtNW1SSmdjMVRvY29UMVpxbVdJc3pZYQpRY25VVnpDVGdFa3V3dkpyQkJjQ09tdllqWitrdEtzdzMzekhDTVo0SWRLaFk3c3hzR0RVOThUektob3JsSFlECmlMRWR3RThmeCtRUkREbVByRDZqbDlYTDFESUZVRnhjY3hlVlNLQUpyb29FVEg3T2dObFNVYVlnYjhRNmhnVEkKUVRaM1ZEYlZlQmR0OWdUMksxQXhxNkNnSzZ2ZURkSm5HNWZpM3dJREFRQUJBb0lCQUFhRkpoc0FiWXU1OXFBZgpUZmpaTTZwaWhoQm5XRzQyWmhyMzVyZGYzREUvSjZ4Vnp3M05JNkZzU3Z4dkpaSzVNRXBiVDMvVURyWDBZL08vClJXSUdWOTJBclYrb3l6elNzMFNrQWJ2S29iVjRMNjNWY2RqMkFuSC9FTW8zM0lmbDIzaDY4VnRSTkNTNU4wYXUKNDc3NXFaTnhvK0xzNHo0dVdibmdsa3AzRkduL29iRDFpcGdoTzVTbkxQOWhpUDRjM3psNnE0T3dibXk0d0M0MApRcHlKTk9LcVVRa2hoVzhlREV4aDlGNW9GYUtMWmxLb2RaejN0c3llUHlPa3dFdHNBSVRoMG13S1RhZlNxdWlOCjBWUDlBMVZWTzlUMXpzZURMK0FZdFBNN2lZekRUN0tSbmNsUE43NmNrTEp5aytpQUlsejFQTzB2QWFEb2p4NmoKUGVpd0d0a0NnWUVBM3dkY1BRcHNSTVdnSWlMVU1xTVdEb0E0RlFxb2lOQmc1SVAxb1RYa1Rhc01ZSVM4NUQ3ZwpxakZXR0RLSnNSZ3I4eCtvSFoxa1BuWThPdlVBTWw1REwxWnZTcC9RRHpCYktlc1lXeU96MFNRWjZsMHZmbjNUClFNNllOektGWm5ZcXlGTXIxTUpNQStMV3hsbUxVNTRhbHdXQUUrWEErMWsvOFk1LzY1UkdIbHNDZ1lFQTB6TjgKbFBuTTdkVE1DMSsvK2V2K0JPQ05iWFdWY1F5OXdFVEw5WUVmZ0VCaDVyc0xQYmR5eWtBazJuNXBTMmhHMDVYNAowcDRzaHAvbmZRaEJaRXovVDlBWVppemtPUGhRL1RhRkY4aFF4VDVkeTVlVFc4OURBa3Jtb3EzbWpBTFdiWjVtCnRhbW1QelgyTTArMzA1bnc0L2d1N2lSalpPVkpvMUVxOVdHdy9NMENnWUVBcUt2TVdtMnpqQjlhQi9jSFBIU1MKamN6eW5SYytkcG9CYlZGUFJ1aVhEUlk3ZWhOcE44VkY2L0Q4QjdqUTRacENRdERDT0FOOGVMQzZ4R3ZlQmptNwozZVVrcmU5SFR2Vm5QTUNMM3dHVlVLcFkzeUl5ZFJ0NzFSdHBpdlV1WmxzZjUvamV3VDFnZDkvcVJvQWFHdGNMCmpRT2Y1V3h4RXFaZzhiS3ZrOTdEV05rQ2dZRUFoaHMxU2l0c0VzQ0NaOTUrdWlVOWdMOU5UbW5SWUoxa2g0ZW4Kc3RZd3VIRXBPU2NmdGlxY093eUwyaWxXbHNrNTMvUmtzT2c3QWFqYmhxc05TckVSbFE1Zno5RkZnVjg3bmUxVwppWWxxc0RRdnZxMGFwcnR1b3pBSVR0ZjVnb0h5d2x4SWY2V2ZxSmVOSTN2RkVCbTV1aWZITlBQcUlSRHV0ME04CjhkNzhVU2tDZ1lFQXpIMWlpL0xWYUJxOFducFBaUFVBOHhtRjZkUEtESGR1T0JVQURNNklFUGx0TXdjWm03NmoKQ3V4RXdkVTE1MVhkNUNDRUxzYVU3d1d0Tjd2TzEvRk9kZTljUStSR2NOd3JCRTR1MWxEMkpseUltOUROZ1BiaApoNm8rOWJQMHlqdGZzNHAreEZVVm5ab0NGd0hqYks4eUM2Q3pqYVZjK2M2NFhuSk1VUmlRNnZJPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\\n}"
      },
      {
        "key": "Addons",
        "value": "{KeyValueList:[{Key:cluster-autoscaler,Value:{\\installOptionsTemplate\\:{},\\name\\:\\cluster-autoscaler\\,\\targetVersion\\:\\2.0.0\\,\\version\\:\\2.0.0\\}},{Key:ibm-storage-operator,Value:{\\healthStatus\\:\\Enabling\\,\\installOptionsTemplate\\:{},\\name\\:\\ibm-storage-operator\\,\\targetVersion\\:\\1.0\\,\\version\\:\\1.0\\}},{Key:vpc-block-csi-driver,Value:{\\allowed_upgrade_versions\\:[\\5.2\\],\\healthStatus\\:\\Enabling\\,\\installOptionsTemplate\\:{},\\name\\:\\vpc-block-csi-driver\\,\\targetVersion\\:\\5.1\\,\\version\\:\\5.1\\}}]}"
      },
      {
        "key": "Status",
        "value": "Active"
      },
      {
        "key": "CreatedTime",
        "value": "2026-10-07T11:06:11Z"
      },
      {
        "key": "TagList",
        "value": "{Key:cb-spider-pmks-autoscaler-status,Value:deploying}"
      }
    ]
  }
}
```

</details>

