# CM-Beetle K8s Infra Migration Test Results — IBMCloud-Sydney

> [!NOTE]
> Full lifecycle against a real CSP: recommend → migrate → list → get (verified against
> the recommendation) → delete → residual resource check.

## Environment

- CSP / Region: ibm / au-syd
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (8928ba9)
- Git Commit: 8928ba9
- Namespace: mig01
- Test Date: 2026-09-16 19:10:41 KST
- Cluster ID: k8s02-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 21ms |
| 2 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 23m30.564s |
| 3 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 1ms |
| 4 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 12.522s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 4m5.219s |
| 6 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 5m42.478s |
| 7 | Residual resource check (Tumblebug) | ✅ **PASS** | 3ms |

**Overall Result**: 7/7 steps passed ✅

**Total Duration**: 33m30s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 21ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.33.13)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=ibm+au-syd+cxf-4x8 nodes=2

### Step 2 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 23m30.564s
- **Status Code**: 202

- ℹ️  nameSeed: k8s02
- ℹ️  async reqId: 1789553441808820764
- ℹ️  cluster id: k8s02-on-prem-k8s-cluster
- ℹ️  elapsed: 23m30s
- ✅ status: Active

### Step 3 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 1ms
- **Status Code**: 200

- ✅ migrated cluster present in list (1 total)

### Step 4 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 12.522s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=ibm+au-syd+cxf-4x8, nodes=2)
- ✅ version: 1.33.13_1580 (recommended 1.33.13)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 4m5.219s

- ✅ kubeconfig obtained (server: https://c104.au-syd.containers.cloud.ibm.com:31606)
- ℹ️  auth method: client certificate in kubeconfig
- ✅ API server reachable (v1.33.13+IKS)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 3)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: c2927af0-au-syd.lb.appdomain.cloud
- ✅ nginx served over the LoadBalancer at http://c2927af0-au-syd.lb.appdomain.cloud/ (attempt 13)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 6 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 5m42.478s
- **Status Code**: 200

- ✅ deleted on attempt 1 (5m42s)

### Step 7 — Residual resource check (Tumblebug)

- **Duration**: 3ms

- ℹ️  VNet k8s02-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup k8s02-k8s-sg still exists (known gap)
- ℹ️  SshKey k8s02-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for ibm au-syd (source: v1.32.3 → target: v1.33.13)",
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
        }
      ],
      "cspResourceId": ""
    }
  ],
  "targetK8sCluster": {
    "connectionName": "ibm-au-syd",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "1.33.13",
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
  "id": "k8s02-on-prem-k8s-cluster",
  "uid": "tbj7i7ch33krh023o26f",
  "name": "k8s02-on-prem-k8s-cluster",
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
    "sys.createdTime": "2026-09-16 10:12:03 +0000 UTC",
    "sys.cspResourceId": "dal6mt0s03rqu2pe3mrg",
    "sys.cspResourceName": "tbj7i7ch33krh023o26f",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "k8s02-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "k8s02-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbj7i7ch33krh023o26f",
    "sys.version": "1.33.13_1580"
  },
  "systemLabel": "",
  "version": "1.33.13_1580",
  "network": {
    "vNetId": "k8s02-k8s-vpc",
    "subnetIds": [
      "k8s02-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "k8s02-k8s-sg"
    ],
    "keyValueList": [
      {
        "key": "VpcIID",
        "value": "{NameId:tbjitcfui5iupk1ohrn2,SystemId:r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653}"
      },
      {
        "key": "SubnetIIDs",
        "value": "{NameId:tbr5pgg6n3mkoagtfkgc,SystemId:02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06}"
      },
      {
        "key": "SecurityGroupIIDs",
        "value": "{NameId:kube-dal6mt0s03rqu2pe3mrg,SystemId:r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de}"
      },
      {
        "key": "KeyValueList",
        "value": "{Key:VpcIID,Value:{NameId:tbjitcfui5iupk1ohrn2,SystemId:r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653}}; {Key:SubnetIIDs,Value:{NameId:tbr5pgg6n3mkoagtfkgc,SystemId:02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06}}; {Key:SecurityGroupIIDs,Value:{NameId:kube-dal6mt0s03rqu2pe3mrg,SystemId:r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de}}"
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
      "sshKeyId": "k8s02-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": -1,
      "maxNodeSize": -1,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "Not visible in IBM",
          "cspResourceId": "kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-0000019b"
        },
        {
          "cspResourceName": "Not visible in IBM",
          "cspResourceId": "kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-000002b1"
        }
      ],
      "keyValueList": [
        {
          "key": "IId",
          "value": "{NameId:workers1,SystemId:dal6mt0s03rqu2pe3mrg-f09c430}"
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
          "value": "{NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-0000019b}; {NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-000002b1}"
        }
      ],
      "cspResourceName": "workers1",
      "cspResourceId": "dal6mt0s03rqu2pe3mrg-f09c430",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "dal6mt0s03rqu2pe3mrg-f09c430"
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
        "MinNodeSize": -1,
        "MaxNodeSize": -1,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "Not visible in IBM",
            "SystemId": "kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-0000019b"
          },
          {
            "NameId": "Not visible in IBM",
            "SystemId": "kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-000002b1"
          }
        ],
        "KeyValueList": [
          {
            "key": "IId",
            "value": "{NameId:workers1,SystemId:dal6mt0s03rqu2pe3mrg-f09c430}"
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
            "value": "{NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-0000019b}; {NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-000002b1}"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://c104.au-syd.containers.cloud.ibm.com:31606",
    "kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURRRENDQWlpZ0F3SUJBZ0lVYzEwNit1NWxCMGtLUE5xR2hJS0svcDhOeFZZd0RRWUpLb1pJaHZjTkFRRUwKQlFBd09ERTJNRFFHQTFVRUF4TXRaR0ZzTm0xME1ITXdNM0p4ZFRKd1pUTnRjbWN0YTNWaVpYSnVaWFJsY3kxagpZUzB4TnpnNU5UVXpOVGt3TUI0WERUSTJNRGt4TmpFd01EZ3dNRm9YRFRNMk1Ea3hNekV3TURnd01Gb3dPREUyCk1EUUdBMVVFQXhNdFpHRnNObTEwTUhNd00zSnhkVEp3WlROdGNtY3RhM1ZpWlhKdVpYUmxjeTFqWVMweE56ZzUKTlRVek5Ua3dNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXNWL20rQ1VNTlpiZQpIdFdReTRXdmZBS0d4eVZmR2h3TEZ5R0hyZVBwRWZCSGdwbnlHMlNydVhSbGlqbVE4L0JBblA1a2VtUm5ScGp4CnR5TGZmYmxTeEwzRHpNSXhNclFqd01MaFp2MWVkM3lJeXAzVzQ0TE5aTVd6WitrUndyWTNKN0cxWTU5S2orNGwKVVpjajZETTVORXFocFhEelhtTEtVWHN2dUwxNTVoQXZHV0ljY0NuK0NFR2F6L245NlROYjJhV3A5VTVBVUtUWQphZVE5TFRTTEdMbm41bGVKLzQ5YzVsMmxzR0FCaWc3amJuM2lLQmNMdlJ3UFhpM0c4U2E1czh0ZXlLWUFDTjI2CnB6cGh2T00ySlltWkNCbzRxV1JNRUt2ZE04bkk1REZpOHI1M3RoVjdyRzVJb091Q1NCaHYyMllrN3Jad1FPNlIKMy9xcnhBczlFd0lEQVFBQm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FZWXdEd1lEVlIwVEFRSC9CQVV3QXdFQgovekFkQmdOVkhRNEVGZ1FVMUFhZTE2NllnelppdEVzM3V4ck5wUFFBdW9vd0RRWUpLb1pJaHZjTkFRRUxCUUFECmdnRUJBSWoxMW5LbFd1cDI0UEJMSDUvN3RFNDNWNGFtQXgyR3JOaGUycFk3NVN6Nm05MzJCMGdYT2R6N3dLZkQKUVhCUFJ1SmV3S0tnZUpudmFXNU9oS1kzeG1XWVhQZ2xvL202L1R4MjV1MXR6Q2lkZE8weVpwenQzZlRsVE0yVQp2WERyVlRvd1JnazVrR1FlcUQwOHFKbG5WVFJGcVRiOW9TRURkOTc5OHRVTHlDZzUzN0hlWjZ4M1BEcTMrZWJkCmdzMGcxZUdrb1BEUGRUVVk5WlphaGRIdDltV0EwaHNOQ3lielptU2h4VHVINmFrQTVobEQrR1ZNS1hUZ3ZMZUEKY0IxWDFnYzJOY0tKZTRmeW5teTI3S0VTcU1mL0tRaXJTR2tjNEJ4Y1o5eGNiUlUyQ002YjVlY0c5Q2JTM2dsegowY3FpanFvUjNuUEhreUN1c3l3RFYzOTJmOTQ9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://c104.au-syd.containers.cloud.ibm.com:31606\n  name: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg\ncontexts:\n- context:\n    cluster: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg\n    namespace: default\n    user: admin/dal6mt0s03rqu2pe3mrg\n  name: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg/admin\ncurrent-context: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg/admin\nkind: Config\nusers:\n- name: admin/dal6mt0s03rqu2pe3mrg\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURzRENDQXBpZ0F3SUJBZ0lSQUt0em9vaWZ5dXVZMW94YmlGdjYxc2d3RFFZSktvWklodmNOQVFFTEJRQXcKT0RFMk1EUUdBMVVFQXhNdFpHRnNObTEwTUhNd00zSnhkVEp3WlROdGNtY3RhM1ZpWlhKdVpYUmxjeTFqWVMweApOemc1TlRVek5Ua3dNQjRYRFRJMk1Ea3hOakV3TWpZd05Gb1hEVEk0TURreE5qRXdNall3TkZvd2F6RUxNQWtHCkExVUVCaE1DVlZNeEZqQVVCZ05WQkFnVERWTmhiaUJHY21GdVkybHpZMjh4Q3pBSkJnTlZCQWNUQWtOQk1STXcKRVFZRFZRUUtFd3BwWW0wdFlXUnRhVzV6TVNJd0lBWURWUVFEREJsTGRXSmxRMlZ5ZENOSlFrMXBaQzAyT1RRdwpNREZDVHpkWE1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBOFpYbHpaeDlrd0xOCjRVQlBTRTcvaVFHVWFZQmducDFMZlNmVWRTNTlHQmRUMm9XQXNrWTdLV2VPZlFaT3UycitFRXJnaEZKR0JoQlgKc0pYTjVQVDVyMWZyWit1TStFdGhBRHBBMXpVZU5GRldNall5Qjdta1BHNFVhZ2N6b3BUblM3NFBuc2VRRmVtRwptOHI2NE1CYjRnVHpDNzNCTWVFd01wbXpPRVlQWkkzcmFoeks3VWFvMVUvVTg1UlRUYnkyTkZCSjdtY1l1SWNGClp3c0RWQ0dnNWVqbW45eFVMajA3c3RNSFVCZ2RIU2pCcGgzbmVIekZIZVFkYm9acUpCZFFJcG9QUHNZcVZrdUMKWE5NajFDaEJ2YXJxR2lZeVRlTkdZbllZbzRER2VvdGJWWUt3U0I4Y1hWbUMrWHJuRHhNb3RjYzE3TVZTR2xYeAp5YU84NFpWS0J3SURBUUFCbzRHQk1IOHdEZ1lEVlIwUEFRSC9CQVFEQWdLRU1CMEdBMVVkSlFRV01CUUdDQ3NHCkFRVUZCd01DQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRk5RR250ZXUKbUlNMllyUkxON3NhemFUMEFMcUtNQjhHQTFVZEVRUVlNQmFCRkdoaGJtbDZZVzVuTnpkQVoyMWhhV3d1WTI5dApNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUJDMTVuVDBIWDZpclF4Tm1KYXE5ZWNOUVZXU1BkV2NRTmt1WUpKCkpGd1Zsb2tGSVExbSs3RWY0RmRhMERDbUE2aDJyQ0FPVE1HeXZrS0YyQzJTN05qN0RKcjM1ZitEcmRid2VHUUkKUnFWQTA2Zm1Va1BCWERMNDBsUDVFeXNaQW1CbGNnNUxvNFliYlRqNk5DMnZWTEYvbEhvWTFlcURoV0R3Qi91cAo5dHJnbGlzWXdXUC9kQUtEb2hhcmI0ZnNYb0pWWHU0WXhtVjB0c21DS041dGhlYkVJb1FEa0RqM1djWmxaOXJICkI0SEhKclh1LzlQUm8yTGdINUMvMytTY3V2Qk0yUXpJSDNnanZJdThKbWRKdWpEbGZ0NjVNdllMMzZIc2lGeGcKWkhPQUFGM3NON2h2bVpIZ1JWcjlRTTlyRWpTWXJweEI2N1lBOGRmbzRhNmhabmJ4Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBOFpYbHpaeDlrd0xONFVCUFNFNy9pUUdVYVlCZ25wMUxmU2ZVZFM1OUdCZFQyb1dBCnNrWTdLV2VPZlFaT3UycitFRXJnaEZKR0JoQlhzSlhONVBUNXIxZnJaK3VNK0V0aEFEcEExelVlTkZGV01qWXkKQjdta1BHNFVhZ2N6b3BUblM3NFBuc2VRRmVtR204cjY0TUJiNGdUekM3M0JNZUV3TXBtek9FWVBaSTNyYWh6Swo3VWFvMVUvVTg1UlRUYnkyTkZCSjdtY1l1SWNGWndzRFZDR2c1ZWptbjl4VUxqMDdzdE1IVUJnZEhTakJwaDNuCmVIekZIZVFkYm9acUpCZFFJcG9QUHNZcVZrdUNYTk1qMUNoQnZhcnFHaVl5VGVOR1luWVlvNERHZW90YlZZS3cKU0I4Y1hWbUMrWHJuRHhNb3RjYzE3TVZTR2xYeHlhTzg0WlZLQndJREFRQUJBb0lCQUE5MXlra0lOTGtOdXVNcApYR21DTkxRdDE3T1F0WjR3N3IzSnFMei9CcDVlRDgyeU1YUTNMbDROOUg1bnd1NFhnTDdHSyt3TDM5TlBoRzBXCmlTQ1gxTXExMDZqSTJES2prRWVWY3NUUzcyWGx0cUJyKzNPbkc2MktWZUFiS2VERHFyR2NMaCs1SWExbFRtbjEKNld3c296U3BvR3dsN3BFa21oTUM1d2M0NUk5SXJlMGw3YmZQeS9UaDhzbWs1T2sxYVBoZ0xvZHhBMGIzTkRPSQp0ZGU5czlKRGRTOU1zY3pBUGxza3I0bUx3N2w0WEYzWW5tUXc5UjY5RnRJQXlCMDJOa05Zbi9zTHA1d1cxdHlDCkRhL1FYM0lSdlBqOUZaa2NZb1ZHOGQxQ2xuRWlWTnJzRE5BK1FrdXN3SjF5TC9QamZvSnlsY1NydTZxUWNBLysKU0NaZmwzRUNnWUVBKzNHaXE1ZDhZKzBSL0VrNnUzZm5ZNDRHRGhpMjBwdE04M3oyM09xWURaekY3MFBNaHJxQwpLOXZzYlJ1MHcvQXBTY2RVRzV5ekJ3NDdpdXJVeXBIY1JjTDlFYkdXU0Z2NTllRGZ1bzBLQTlBMCtldnN1WWFxCnpvdXIwem5WVjBhNUI5Y2F4OW04NnF4SmJmQXhMNS9oNkhqQnRFVmZ0elZWcjhZR0l2WTJRaE1DZ1lFQTlmYUkKV2xDeFFkQlI0ekJwalRWSGsvbVpoNnM3cFBpUzMxUWNINHhibEFVdXVETE9iRStYdHFNVFhTYWl3SUJmYnJPaQpXYXBEZjFTK0ZjNFFlTUFjekFTY0xVMzQ2RGd2YmorV2YwS1krZ1pLNXhvek5iNnRFNEsxMEhMTVNyZTk4SWg3CnMyNVJiN1kvamFOc3lqcmFyQ2p2QjkxRHBYOU5mdWZDdEY3ZnRyMENnWUVBMG9xSVIxNEZEamNJQkZQZEZmU0UKajl2d1BnVjdzRVhSM1dBWjVVbWFJR3ZSWVZOSUF0aFEveUNiaTVEVGYrMnM1TlkvR3cvTzZHMkdkZi9FUmdwMApnd1dPbWk0MVJFbWZ0NzZnRjdqWlZmQVZLOS9jekV4eTRaZ2FQRGdFNTV1VWUzZ21PSW1kb25LNDJaRngzZ3JtClFwNDZ0QlFTM1htUFVpdGlJQXhCeW5rQ2dZQWxnL0RRTmJhVG56NmVOR2dsRFpkWlRweklRS25jUTczREtvVVAKbXN6dENzMVJjdzVoSHRLNUhLNTdhc1V3TDJSZThpODFGZTh0b0xOTmlCeWpEa3BXSSszZVN5Skg2U255MnVnTgppUTdrTThtQTdsSVpSSGdKbmNvMWZRMEQ2SHFrRVcwc3RRcmV5eUZ1YlJyT3phTkUxd2wrWFpWUHpOYjVJRWhtClVvSTAwUUtCZ1FEWkhTTllUbjZETWM5R0N2S3ZKNTRGME43d3hXYzcyak9JYkVXUWNXUUw0Tk5hNW5OVkRmRlkKMWQ0aktSOFg2UmZ3Tzg3QW5qaTBvdWc5YUs4ZitFL2J2cUdiK0lubjR3N0pHTlJ1dWJOM0JwQ0xEbjR3cmg3KwpTNk5keHBWMTFBWVpqRkN1TmhJL3k5WDBoZk9USHNiZ1IzeGpBaGJxcFUrK3dMUXY1ZWR6SFE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\n"
  },
  "addons": {
    "keyValueList": [
      {
        "key": "cluster-autoscaler",
        "value": "{\"allowed_upgrade_versions\":[\"2.0.0\"],\"installOptionsTemplate\":{},\"name\":\"cluster-autoscaler\",\"targetVersion\":\"1.2.4\",\"version\":\"1.2.4\"}"
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
  "createdTime": "2026-09-16T10:12:03Z",
  "keyValueList": [
    {
      "key": "IId",
      "value": "{NameId:tbj7i7ch33krh023o26f,SystemId:dal6mt0s03rqu2pe3mrg}"
    },
    {
      "key": "Version",
      "value": "1.33.13_1580"
    },
    {
      "key": "Network",
      "value": "{VpcIID:{NameId:tbjitcfui5iupk1ohrn2,SystemId:r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653},SubnetIIDs:[{NameId:tbr5pgg6n3mkoagtfkgc,SystemId:02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06}],SecurityGroupIIDs:[{NameId:kube-dal6mt0s03rqu2pe3mrg,SystemId:r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de}],KeyValueList:[{Key:VpcIID,Value:{NameId:tbjitcfui5iupk1ohrn2,SystemId:r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653}},{Key:SubnetIIDs,Value:{NameId:tbr5pgg6n3mkoagtfkgc,SystemId:02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06}},{Key:SecurityGroupIIDs,Value:{NameId:kube-dal6mt0s03rqu2pe3mrg,SystemId:r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de}},{Key:KeyValueList,Value:{Key:VpcIID,Value:{NameId:tbjitcfui5iupk1ohrn2,SystemId:r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653}}; {Key:SubnetIIDs,Value:{NameId:tbr5pgg6n3mkoagtfkgc,SystemId:02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06}}; {Key:SecurityGroupIIDs,Value:{NameId:kube-dal6mt0s03rqu2pe3mrg,SystemId:r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de}}}]}"
    },
    {
      "key": "NodeGroupList",
      "value": "{IId:{NameId:workers1,SystemId:dal6mt0s03rqu2pe3mrg-f09c430},ImageIID:{NameId:Not visible in IBM,SystemId:Not visible in IBM},VMSpecName:cxf.4x8,RootDiskType:Not visible in IBM,RootDiskSize:Not visible in IBM,KeyPairIID:{NameId:Not visible in IBM,SystemId:Not visible in IBM},OnAutoScaling:false,DesiredNodeSize:2,MinNodeSize:-1,MaxNodeSize:-1,Status:Active,Nodes:[{NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-0000019b},{NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-000002b1}],KeyValueList:[{Key:IId,Value:{NameId:workers1,SystemId:dal6mt0s03rqu2pe3mrg-f09c430}},{Key:ImageIID,Value:{NameId:Not visible in IBM,SystemId:Not visible in IBM}},{Key:VMSpecName,Value:cxf.4x8},{Key:RootDiskType,Value:Not visible in IBM},{Key:RootDiskSize,Value:Not visible in IBM},{Key:KeyPairIID,Value:{NameId:Not visible in IBM,SystemId:Not visible in IBM}},{Key:OnAutoScaling,Value:false},{Key:DesiredNodeSize,Value:2},{Key:MinNodeSize,Value:0},{Key:MaxNodeSize,Value:0},{Key:Status,Value:Active},{Key:Nodes,Value:{NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-0000019b}; {NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-000002b1}}]}"
    },
    {
      "key": "AccessInfo",
      "value": "{Endpoint:https://c104.au-syd.containers.cloud.ibm.com:31606,Kubeconfig:apiVersion: v1\\nclusters:\\n- cluster:\\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURRRENDQWlpZ0F3SUJBZ0lVYzEwNit1NWxCMGtLUE5xR2hJS0svcDhOeFZZd0RRWUpLb1pJaHZjTkFRRUwKQlFBd09ERTJNRFFHQTFVRUF4TXRaR0ZzTm0xME1ITXdNM0p4ZFRKd1pUTnRjbWN0YTNWaVpYSnVaWFJsY3kxagpZUzB4TnpnNU5UVXpOVGt3TUI0WERUSTJNRGt4TmpFd01EZ3dNRm9YRFRNMk1Ea3hNekV3TURnd01Gb3dPREUyCk1EUUdBMVVFQXhNdFpHRnNObTEwTUhNd00zSnhkVEp3WlROdGNtY3RhM1ZpWlhKdVpYUmxjeTFqWVMweE56ZzUKTlRVek5Ua3dNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXNWL20rQ1VNTlpiZQpIdFdReTRXdmZBS0d4eVZmR2h3TEZ5R0hyZVBwRWZCSGdwbnlHMlNydVhSbGlqbVE4L0JBblA1a2VtUm5ScGp4CnR5TGZmYmxTeEwzRHpNSXhNclFqd01MaFp2MWVkM3lJeXAzVzQ0TE5aTVd6WitrUndyWTNKN0cxWTU5S2orNGwKVVpjajZETTVORXFocFhEelhtTEtVWHN2dUwxNTVoQXZHV0ljY0NuK0NFR2F6L245NlROYjJhV3A5VTVBVUtUWQphZVE5TFRTTEdMbm41bGVKLzQ5YzVsMmxzR0FCaWc3amJuM2lLQmNMdlJ3UFhpM0c4U2E1czh0ZXlLWUFDTjI2CnB6cGh2T00ySlltWkNCbzRxV1JNRUt2ZE04bkk1REZpOHI1M3RoVjdyRzVJb091Q1NCaHYyMllrN3Jad1FPNlIKMy9xcnhBczlFd0lEQVFBQm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FZWXdEd1lEVlIwVEFRSC9CQVV3QXdFQgovekFkQmdOVkhRNEVGZ1FVMUFhZTE2NllnelppdEVzM3V4ck5wUFFBdW9vd0RRWUpLb1pJaHZjTkFRRUxCUUFECmdnRUJBSWoxMW5LbFd1cDI0UEJMSDUvN3RFNDNWNGFtQXgyR3JOaGUycFk3NVN6Nm05MzJCMGdYT2R6N3dLZkQKUVhCUFJ1SmV3S0tnZUpudmFXNU9oS1kzeG1XWVhQZ2xvL202L1R4MjV1MXR6Q2lkZE8weVpwenQzZlRsVE0yVQp2WERyVlRvd1JnazVrR1FlcUQwOHFKbG5WVFJGcVRiOW9TRURkOTc5OHRVTHlDZzUzN0hlWjZ4M1BEcTMrZWJkCmdzMGcxZUdrb1BEUGRUVVk5WlphaGRIdDltV0EwaHNOQ3lielptU2h4VHVINmFrQTVobEQrR1ZNS1hUZ3ZMZUEKY0IxWDFnYzJOY0tKZTRmeW5teTI3S0VTcU1mL0tRaXJTR2tjNEJ4Y1o5eGNiUlUyQ002YjVlY0c5Q2JTM2dsegowY3FpanFvUjNuUEhreUN1c3l3RFYzOTJmOTQ9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\\n    server: https://c104.au-syd.containers.cloud.ibm.com:31606\\n  name: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg\\ncontexts:\\n- context:\\n    cluster: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg\\n    namespace: default\\n    user: admin/dal6mt0s03rqu2pe3mrg\\n  name: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg/admin\\ncurrent-context: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg/admin\\nkind: Config\\nusers:\\n- name: admin/dal6mt0s03rqu2pe3mrg\\n  user:\\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURzRENDQXBpZ0F3SUJBZ0lSQUt0em9vaWZ5dXVZMW94YmlGdjYxc2d3RFFZSktvWklodmNOQVFFTEJRQXcKT0RFMk1EUUdBMVVFQXhNdFpHRnNObTEwTUhNd00zSnhkVEp3WlROdGNtY3RhM1ZpWlhKdVpYUmxjeTFqWVMweApOemc1TlRVek5Ua3dNQjRYRFRJMk1Ea3hOakV3TWpZd05Gb1hEVEk0TURreE5qRXdNall3TkZvd2F6RUxNQWtHCkExVUVCaE1DVlZNeEZqQVVCZ05WQkFnVERWTmhiaUJHY21GdVkybHpZMjh4Q3pBSkJnTlZCQWNUQWtOQk1STXcKRVFZRFZRUUtFd3BwWW0wdFlXUnRhVzV6TVNJd0lBWURWUVFEREJsTGRXSmxRMlZ5ZENOSlFrMXBaQzAyT1RRdwpNREZDVHpkWE1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBOFpYbHpaeDlrd0xOCjRVQlBTRTcvaVFHVWFZQmducDFMZlNmVWRTNTlHQmRUMm9XQXNrWTdLV2VPZlFaT3UycitFRXJnaEZKR0JoQlgKc0pYTjVQVDVyMWZyWit1TStFdGhBRHBBMXpVZU5GRldNall5Qjdta1BHNFVhZ2N6b3BUblM3NFBuc2VRRmVtRwptOHI2NE1CYjRnVHpDNzNCTWVFd01wbXpPRVlQWkkzcmFoeks3VWFvMVUvVTg1UlRUYnkyTkZCSjdtY1l1SWNGClp3c0RWQ0dnNWVqbW45eFVMajA3c3RNSFVCZ2RIU2pCcGgzbmVIekZIZVFkYm9acUpCZFFJcG9QUHNZcVZrdUMKWE5NajFDaEJ2YXJxR2lZeVRlTkdZbllZbzRER2VvdGJWWUt3U0I4Y1hWbUMrWHJuRHhNb3RjYzE3TVZTR2xYeAp5YU84NFpWS0J3SURBUUFCbzRHQk1IOHdEZ1lEVlIwUEFRSC9CQVFEQWdLRU1CMEdBMVVkSlFRV01CUUdDQ3NHCkFRVUZCd01DQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRk5RR250ZXUKbUlNMllyUkxON3NhemFUMEFMcUtNQjhHQTFVZEVRUVlNQmFCRkdoaGJtbDZZVzVuTnpkQVoyMWhhV3d1WTI5dApNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUJDMTVuVDBIWDZpclF4Tm1KYXE5ZWNOUVZXU1BkV2NRTmt1WUpKCkpGd1Zsb2tGSVExbSs3RWY0RmRhMERDbUE2aDJyQ0FPVE1HeXZrS0YyQzJTN05qN0RKcjM1ZitEcmRid2VHUUkKUnFWQTA2Zm1Va1BCWERMNDBsUDVFeXNaQW1CbGNnNUxvNFliYlRqNk5DMnZWTEYvbEhvWTFlcURoV0R3Qi91cAo5dHJnbGlzWXdXUC9kQUtEb2hhcmI0ZnNYb0pWWHU0WXhtVjB0c21DS041dGhlYkVJb1FEa0RqM1djWmxaOXJICkI0SEhKclh1LzlQUm8yTGdINUMvMytTY3V2Qk0yUXpJSDNnanZJdThKbWRKdWpEbGZ0NjVNdllMMzZIc2lGeGcKWkhPQUFGM3NON2h2bVpIZ1JWcjlRTTlyRWpTWXJweEI2N1lBOGRmbzRhNmhabmJ4Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBOFpYbHpaeDlrd0xONFVCUFNFNy9pUUdVYVlCZ25wMUxmU2ZVZFM1OUdCZFQyb1dBCnNrWTdLV2VPZlFaT3UycitFRXJnaEZKR0JoQlhzSlhONVBUNXIxZnJaK3VNK0V0aEFEcEExelVlTkZGV01qWXkKQjdta1BHNFVhZ2N6b3BUblM3NFBuc2VRRmVtR204cjY0TUJiNGdUekM3M0JNZUV3TXBtek9FWVBaSTNyYWh6Swo3VWFvMVUvVTg1UlRUYnkyTkZCSjdtY1l1SWNGWndzRFZDR2c1ZWptbjl4VUxqMDdzdE1IVUJnZEhTakJwaDNuCmVIekZIZVFkYm9acUpCZFFJcG9QUHNZcVZrdUNYTk1qMUNoQnZhcnFHaVl5VGVOR1luWVlvNERHZW90YlZZS3cKU0I4Y1hWbUMrWHJuRHhNb3RjYzE3TVZTR2xYeHlhTzg0WlZLQndJREFRQUJBb0lCQUE5MXlra0lOTGtOdXVNcApYR21DTkxRdDE3T1F0WjR3N3IzSnFMei9CcDVlRDgyeU1YUTNMbDROOUg1bnd1NFhnTDdHSyt3TDM5TlBoRzBXCmlTQ1gxTXExMDZqSTJES2prRWVWY3NUUzcyWGx0cUJyKzNPbkc2MktWZUFiS2VERHFyR2NMaCs1SWExbFRtbjEKNld3c296U3BvR3dsN3BFa21oTUM1d2M0NUk5SXJlMGw3YmZQeS9UaDhzbWs1T2sxYVBoZ0xvZHhBMGIzTkRPSQp0ZGU5czlKRGRTOU1zY3pBUGxza3I0bUx3N2w0WEYzWW5tUXc5UjY5RnRJQXlCMDJOa05Zbi9zTHA1d1cxdHlDCkRhL1FYM0lSdlBqOUZaa2NZb1ZHOGQxQ2xuRWlWTnJzRE5BK1FrdXN3SjF5TC9QamZvSnlsY1NydTZxUWNBLysKU0NaZmwzRUNnWUVBKzNHaXE1ZDhZKzBSL0VrNnUzZm5ZNDRHRGhpMjBwdE04M3oyM09xWURaekY3MFBNaHJxQwpLOXZzYlJ1MHcvQXBTY2RVRzV5ekJ3NDdpdXJVeXBIY1JjTDlFYkdXU0Z2NTllRGZ1bzBLQTlBMCtldnN1WWFxCnpvdXIwem5WVjBhNUI5Y2F4OW04NnF4SmJmQXhMNS9oNkhqQnRFVmZ0elZWcjhZR0l2WTJRaE1DZ1lFQTlmYUkKV2xDeFFkQlI0ekJwalRWSGsvbVpoNnM3cFBpUzMxUWNINHhibEFVdXVETE9iRStYdHFNVFhTYWl3SUJmYnJPaQpXYXBEZjFTK0ZjNFFlTUFjekFTY0xVMzQ2RGd2YmorV2YwS1krZ1pLNXhvek5iNnRFNEsxMEhMTVNyZTk4SWg3CnMyNVJiN1kvamFOc3lqcmFyQ2p2QjkxRHBYOU5mdWZDdEY3ZnRyMENnWUVBMG9xSVIxNEZEamNJQkZQZEZmU0UKajl2d1BnVjdzRVhSM1dBWjVVbWFJR3ZSWVZOSUF0aFEveUNiaTVEVGYrMnM1TlkvR3cvTzZHMkdkZi9FUmdwMApnd1dPbWk0MVJFbWZ0NzZnRjdqWlZmQVZLOS9jekV4eTRaZ2FQRGdFNTV1VWUzZ21PSW1kb25LNDJaRngzZ3JtClFwNDZ0QlFTM1htUFVpdGlJQXhCeW5rQ2dZQWxnL0RRTmJhVG56NmVOR2dsRFpkWlRweklRS25jUTczREtvVVAKbXN6dENzMVJjdzVoSHRLNUhLNTdhc1V3TDJSZThpODFGZTh0b0xOTmlCeWpEa3BXSSszZVN5Skg2U255MnVnTgppUTdrTThtQTdsSVpSSGdKbmNvMWZRMEQ2SHFrRVcwc3RRcmV5eUZ1YlJyT3phTkUxd2wrWFpWUHpOYjVJRWhtClVvSTAwUUtCZ1FEWkhTTllUbjZETWM5R0N2S3ZKNTRGME43d3hXYzcyak9JYkVXUWNXUUw0Tk5hNW5OVkRmRlkKMWQ0aktSOFg2UmZ3Tzg3QW5qaTBvdWc5YUs4ZitFL2J2cUdiK0lubjR3N0pHTlJ1dWJOM0JwQ0xEbjR3cmg3KwpTNk5keHBWMTFBWVpqRkN1TmhJL3k5WDBoZk9USHNiZ1IzeGpBaGJxcFUrK3dMUXY1ZWR6SFE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\\n}"
    },
    {
      "key": "Addons",
      "value": "{KeyValueList:[{Key:cluster-autoscaler,Value:{\\allowed_upgrade_versions\\:[\\2.0.0\\],\\installOptionsTemplate\\:{},\\name\\:\\cluster-autoscaler\\,\\targetVersion\\:\\1.2.4\\,\\version\\:\\1.2.4\\}},{Key:ibm-storage-operator,Value:{\\healthStatus\\:\\Enabling\\,\\installOptionsTemplate\\:{},\\name\\:\\ibm-storage-operator\\,\\targetVersion\\:\\1.0\\,\\version\\:\\1.0\\}},{Key:vpc-block-csi-driver,Value:{\\allowed_upgrade_versions\\:[\\5.2\\],\\healthStatus\\:\\Enabling\\,\\installOptionsTemplate\\:{},\\name\\:\\vpc-block-csi-driver\\,\\targetVersion\\:\\5.1\\,\\version\\:\\5.1\\}}]}"
    },
    {
      "key": "Status",
      "value": "Active"
    },
    {
      "key": "CreatedTime",
      "value": "2026-09-16T10:12:03Z"
    },
    {
      "key": "TagList",
      "value": "{Key:cb-spider-pmks-autoscaler-status,Value:deploying}"
    }
  ],
  "cspResourceName": "tbj7i7ch33krh023o26f",
  "cspResourceId": "dal6mt0s03rqu2pe3mrg",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbj7i7ch33krh023o26f",
      "SystemId": "dal6mt0s03rqu2pe3mrg"
    },
    "Version": "1.33.13_1580",
    "Network": {
      "VpcIID": {
        "NameId": "tbjitcfui5iupk1ohrn2",
        "SystemId": "r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653"
      },
      "SubnetIIDs": [
        {
          "NameId": "tbr5pgg6n3mkoagtfkgc",
          "SystemId": "02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "kube-dal6mt0s03rqu2pe3mrg",
          "SystemId": "r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de"
        }
      ],
      "KeyValueList": [
        {
          "key": "VpcIID",
          "value": "{NameId:tbjitcfui5iupk1ohrn2,SystemId:r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653}"
        },
        {
          "key": "SubnetIIDs",
          "value": "{NameId:tbr5pgg6n3mkoagtfkgc,SystemId:02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06}"
        },
        {
          "key": "SecurityGroupIIDs",
          "value": "{NameId:kube-dal6mt0s03rqu2pe3mrg,SystemId:r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de}"
        },
        {
          "key": "KeyValueList",
          "value": "{Key:VpcIID,Value:{NameId:tbjitcfui5iupk1ohrn2,SystemId:r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653}}; {Key:SubnetIIDs,Value:{NameId:tbr5pgg6n3mkoagtfkgc,SystemId:02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06}}; {Key:SecurityGroupIIDs,Value:{NameId:kube-dal6mt0s03rqu2pe3mrg,SystemId:r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de}}"
        }
      ]
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "dal6mt0s03rqu2pe3mrg-f09c430"
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
        "MinNodeSize": -1,
        "MaxNodeSize": -1,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "Not visible in IBM",
            "SystemId": "kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-0000019b"
          },
          {
            "NameId": "Not visible in IBM",
            "SystemId": "kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-000002b1"
          }
        ],
        "KeyValueList": [
          {
            "key": "IId",
            "value": "{NameId:workers1,SystemId:dal6mt0s03rqu2pe3mrg-f09c430}"
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
            "value": "{NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-0000019b}; {NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-000002b1}"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://c104.au-syd.containers.cloud.ibm.com:31606",
      "Kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURRRENDQWlpZ0F3SUJBZ0lVYzEwNit1NWxCMGtLUE5xR2hJS0svcDhOeFZZd0RRWUpLb1pJaHZjTkFRRUwKQlFBd09ERTJNRFFHQTFVRUF4TXRaR0ZzTm0xME1ITXdNM0p4ZFRKd1pUTnRjbWN0YTNWaVpYSnVaWFJsY3kxagpZUzB4TnpnNU5UVXpOVGt3TUI0WERUSTJNRGt4TmpFd01EZ3dNRm9YRFRNMk1Ea3hNekV3TURnd01Gb3dPREUyCk1EUUdBMVVFQXhNdFpHRnNObTEwTUhNd00zSnhkVEp3WlROdGNtY3RhM1ZpWlhKdVpYUmxjeTFqWVMweE56ZzUKTlRVek5Ua3dNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXNWL20rQ1VNTlpiZQpIdFdReTRXdmZBS0d4eVZmR2h3TEZ5R0hyZVBwRWZCSGdwbnlHMlNydVhSbGlqbVE4L0JBblA1a2VtUm5ScGp4CnR5TGZmYmxTeEwzRHpNSXhNclFqd01MaFp2MWVkM3lJeXAzVzQ0TE5aTVd6WitrUndyWTNKN0cxWTU5S2orNGwKVVpjajZETTVORXFocFhEelhtTEtVWHN2dUwxNTVoQXZHV0ljY0NuK0NFR2F6L245NlROYjJhV3A5VTVBVUtUWQphZVE5TFRTTEdMbm41bGVKLzQ5YzVsMmxzR0FCaWc3amJuM2lLQmNMdlJ3UFhpM0c4U2E1czh0ZXlLWUFDTjI2CnB6cGh2T00ySlltWkNCbzRxV1JNRUt2ZE04bkk1REZpOHI1M3RoVjdyRzVJb091Q1NCaHYyMllrN3Jad1FPNlIKMy9xcnhBczlFd0lEQVFBQm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FZWXdEd1lEVlIwVEFRSC9CQVV3QXdFQgovekFkQmdOVkhRNEVGZ1FVMUFhZTE2NllnelppdEVzM3V4ck5wUFFBdW9vd0RRWUpLb1pJaHZjTkFRRUxCUUFECmdnRUJBSWoxMW5LbFd1cDI0UEJMSDUvN3RFNDNWNGFtQXgyR3JOaGUycFk3NVN6Nm05MzJCMGdYT2R6N3dLZkQKUVhCUFJ1SmV3S0tnZUpudmFXNU9oS1kzeG1XWVhQZ2xvL202L1R4MjV1MXR6Q2lkZE8weVpwenQzZlRsVE0yVQp2WERyVlRvd1JnazVrR1FlcUQwOHFKbG5WVFJGcVRiOW9TRURkOTc5OHRVTHlDZzUzN0hlWjZ4M1BEcTMrZWJkCmdzMGcxZUdrb1BEUGRUVVk5WlphaGRIdDltV0EwaHNOQ3lielptU2h4VHVINmFrQTVobEQrR1ZNS1hUZ3ZMZUEKY0IxWDFnYzJOY0tKZTRmeW5teTI3S0VTcU1mL0tRaXJTR2tjNEJ4Y1o5eGNiUlUyQ002YjVlY0c5Q2JTM2dsegowY3FpanFvUjNuUEhreUN1c3l3RFYzOTJmOTQ9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    server: https://c104.au-syd.containers.cloud.ibm.com:31606\n  name: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg\ncontexts:\n- context:\n    cluster: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg\n    namespace: default\n    user: admin/dal6mt0s03rqu2pe3mrg\n  name: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg/admin\ncurrent-context: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg/admin\nkind: Config\nusers:\n- name: admin/dal6mt0s03rqu2pe3mrg\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURzRENDQXBpZ0F3SUJBZ0lSQUt0em9vaWZ5dXVZMW94YmlGdjYxc2d3RFFZSktvWklodmNOQVFFTEJRQXcKT0RFMk1EUUdBMVVFQXhNdFpHRnNObTEwTUhNd00zSnhkVEp3WlROdGNtY3RhM1ZpWlhKdVpYUmxjeTFqWVMweApOemc1TlRVek5Ua3dNQjRYRFRJMk1Ea3hOakV3TWpZd05Gb1hEVEk0TURreE5qRXdNall3TkZvd2F6RUxNQWtHCkExVUVCaE1DVlZNeEZqQVVCZ05WQkFnVERWTmhiaUJHY21GdVkybHpZMjh4Q3pBSkJnTlZCQWNUQWtOQk1STXcKRVFZRFZRUUtFd3BwWW0wdFlXUnRhVzV6TVNJd0lBWURWUVFEREJsTGRXSmxRMlZ5ZENOSlFrMXBaQzAyT1RRdwpNREZDVHpkWE1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBOFpYbHpaeDlrd0xOCjRVQlBTRTcvaVFHVWFZQmducDFMZlNmVWRTNTlHQmRUMm9XQXNrWTdLV2VPZlFaT3UycitFRXJnaEZKR0JoQlgKc0pYTjVQVDVyMWZyWit1TStFdGhBRHBBMXpVZU5GRldNall5Qjdta1BHNFVhZ2N6b3BUblM3NFBuc2VRRmVtRwptOHI2NE1CYjRnVHpDNzNCTWVFd01wbXpPRVlQWkkzcmFoeks3VWFvMVUvVTg1UlRUYnkyTkZCSjdtY1l1SWNGClp3c0RWQ0dnNWVqbW45eFVMajA3c3RNSFVCZ2RIU2pCcGgzbmVIekZIZVFkYm9acUpCZFFJcG9QUHNZcVZrdUMKWE5NajFDaEJ2YXJxR2lZeVRlTkdZbllZbzRER2VvdGJWWUt3U0I4Y1hWbUMrWHJuRHhNb3RjYzE3TVZTR2xYeAp5YU84NFpWS0J3SURBUUFCbzRHQk1IOHdEZ1lEVlIwUEFRSC9CQVFEQWdLRU1CMEdBMVVkSlFRV01CUUdDQ3NHCkFRVUZCd01DQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRk5RR250ZXUKbUlNMllyUkxON3NhemFUMEFMcUtNQjhHQTFVZEVRUVlNQmFCRkdoaGJtbDZZVzVuTnpkQVoyMWhhV3d1WTI5dApNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUJDMTVuVDBIWDZpclF4Tm1KYXE5ZWNOUVZXU1BkV2NRTmt1WUpKCkpGd1Zsb2tGSVExbSs3RWY0RmRhMERDbUE2aDJyQ0FPVE1HeXZrS0YyQzJTN05qN0RKcjM1ZitEcmRid2VHUUkKUnFWQTA2Zm1Va1BCWERMNDBsUDVFeXNaQW1CbGNnNUxvNFliYlRqNk5DMnZWTEYvbEhvWTFlcURoV0R3Qi91cAo5dHJnbGlzWXdXUC9kQUtEb2hhcmI0ZnNYb0pWWHU0WXhtVjB0c21DS041dGhlYkVJb1FEa0RqM1djWmxaOXJICkI0SEhKclh1LzlQUm8yTGdINUMvMytTY3V2Qk0yUXpJSDNnanZJdThKbWRKdWpEbGZ0NjVNdllMMzZIc2lGeGcKWkhPQUFGM3NON2h2bVpIZ1JWcjlRTTlyRWpTWXJweEI2N1lBOGRmbzRhNmhabmJ4Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBOFpYbHpaeDlrd0xONFVCUFNFNy9pUUdVYVlCZ25wMUxmU2ZVZFM1OUdCZFQyb1dBCnNrWTdLV2VPZlFaT3UycitFRXJnaEZKR0JoQlhzSlhONVBUNXIxZnJaK3VNK0V0aEFEcEExelVlTkZGV01qWXkKQjdta1BHNFVhZ2N6b3BUblM3NFBuc2VRRmVtR204cjY0TUJiNGdUekM3M0JNZUV3TXBtek9FWVBaSTNyYWh6Swo3VWFvMVUvVTg1UlRUYnkyTkZCSjdtY1l1SWNGWndzRFZDR2c1ZWptbjl4VUxqMDdzdE1IVUJnZEhTakJwaDNuCmVIekZIZVFkYm9acUpCZFFJcG9QUHNZcVZrdUNYTk1qMUNoQnZhcnFHaVl5VGVOR1luWVlvNERHZW90YlZZS3cKU0I4Y1hWbUMrWHJuRHhNb3RjYzE3TVZTR2xYeHlhTzg0WlZLQndJREFRQUJBb0lCQUE5MXlra0lOTGtOdXVNcApYR21DTkxRdDE3T1F0WjR3N3IzSnFMei9CcDVlRDgyeU1YUTNMbDROOUg1bnd1NFhnTDdHSyt3TDM5TlBoRzBXCmlTQ1gxTXExMDZqSTJES2prRWVWY3NUUzcyWGx0cUJyKzNPbkc2MktWZUFiS2VERHFyR2NMaCs1SWExbFRtbjEKNld3c296U3BvR3dsN3BFa21oTUM1d2M0NUk5SXJlMGw3YmZQeS9UaDhzbWs1T2sxYVBoZ0xvZHhBMGIzTkRPSQp0ZGU5czlKRGRTOU1zY3pBUGxza3I0bUx3N2w0WEYzWW5tUXc5UjY5RnRJQXlCMDJOa05Zbi9zTHA1d1cxdHlDCkRhL1FYM0lSdlBqOUZaa2NZb1ZHOGQxQ2xuRWlWTnJzRE5BK1FrdXN3SjF5TC9QamZvSnlsY1NydTZxUWNBLysKU0NaZmwzRUNnWUVBKzNHaXE1ZDhZKzBSL0VrNnUzZm5ZNDRHRGhpMjBwdE04M3oyM09xWURaekY3MFBNaHJxQwpLOXZzYlJ1MHcvQXBTY2RVRzV5ekJ3NDdpdXJVeXBIY1JjTDlFYkdXU0Z2NTllRGZ1bzBLQTlBMCtldnN1WWFxCnpvdXIwem5WVjBhNUI5Y2F4OW04NnF4SmJmQXhMNS9oNkhqQnRFVmZ0elZWcjhZR0l2WTJRaE1DZ1lFQTlmYUkKV2xDeFFkQlI0ekJwalRWSGsvbVpoNnM3cFBpUzMxUWNINHhibEFVdXVETE9iRStYdHFNVFhTYWl3SUJmYnJPaQpXYXBEZjFTK0ZjNFFlTUFjekFTY0xVMzQ2RGd2YmorV2YwS1krZ1pLNXhvek5iNnRFNEsxMEhMTVNyZTk4SWg3CnMyNVJiN1kvamFOc3lqcmFyQ2p2QjkxRHBYOU5mdWZDdEY3ZnRyMENnWUVBMG9xSVIxNEZEamNJQkZQZEZmU0UKajl2d1BnVjdzRVhSM1dBWjVVbWFJR3ZSWVZOSUF0aFEveUNiaTVEVGYrMnM1TlkvR3cvTzZHMkdkZi9FUmdwMApnd1dPbWk0MVJFbWZ0NzZnRjdqWlZmQVZLOS9jekV4eTRaZ2FQRGdFNTV1VWUzZ21PSW1kb25LNDJaRngzZ3JtClFwNDZ0QlFTM1htUFVpdGlJQXhCeW5rQ2dZQWxnL0RRTmJhVG56NmVOR2dsRFpkWlRweklRS25jUTczREtvVVAKbXN6dENzMVJjdzVoSHRLNUhLNTdhc1V3TDJSZThpODFGZTh0b0xOTmlCeWpEa3BXSSszZVN5Skg2U255MnVnTgppUTdrTThtQTdsSVpSSGdKbmNvMWZRMEQ2SHFrRVcwc3RRcmV5eUZ1YlJyT3phTkUxd2wrWFpWUHpOYjVJRWhtClVvSTAwUUtCZ1FEWkhTTllUbjZETWM5R0N2S3ZKNTRGME43d3hXYzcyak9JYkVXUWNXUUw0Tk5hNW5OVkRmRlkKMWQ0aktSOFg2UmZ3Tzg3QW5qaTBvdWc5YUs4ZitFL2J2cUdiK0lubjR3N0pHTlJ1dWJOM0JwQ0xEbjR3cmg3KwpTNk5keHBWMTFBWVpqRkN1TmhJL3k5WDBoZk9USHNiZ1IzeGpBaGJxcFUrK3dMUXY1ZWR6SFE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\n"
    },
    "Addons": {
      "KeyValueList": [
        {
          "key": "cluster-autoscaler",
          "value": "{\"allowed_upgrade_versions\":[\"2.0.0\"],\"installOptionsTemplate\":{},\"name\":\"cluster-autoscaler\",\"targetVersion\":\"1.2.4\",\"version\":\"1.2.4\"}"
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
    "CreatedTime": "2026-09-16T10:12:03Z",
    "KeyValueList": [
      {
        "key": "IId",
        "value": "{NameId:tbj7i7ch33krh023o26f,SystemId:dal6mt0s03rqu2pe3mrg}"
      },
      {
        "key": "Version",
        "value": "1.33.13_1580"
      },
      {
        "key": "Network",
        "value": "{VpcIID:{NameId:tbjitcfui5iupk1ohrn2,SystemId:r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653},SubnetIIDs:[{NameId:tbr5pgg6n3mkoagtfkgc,SystemId:02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06}],SecurityGroupIIDs:[{NameId:kube-dal6mt0s03rqu2pe3mrg,SystemId:r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de}],KeyValueList:[{Key:VpcIID,Value:{NameId:tbjitcfui5iupk1ohrn2,SystemId:r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653}},{Key:SubnetIIDs,Value:{NameId:tbr5pgg6n3mkoagtfkgc,SystemId:02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06}},{Key:SecurityGroupIIDs,Value:{NameId:kube-dal6mt0s03rqu2pe3mrg,SystemId:r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de}},{Key:KeyValueList,Value:{Key:VpcIID,Value:{NameId:tbjitcfui5iupk1ohrn2,SystemId:r026-ea7b8de8-38c6-4d7a-bd4d-1b9f2be79653}}; {Key:SubnetIIDs,Value:{NameId:tbr5pgg6n3mkoagtfkgc,SystemId:02h7-18fae5ee-0b08-4e95-b9dd-10bea99baa06}}; {Key:SecurityGroupIIDs,Value:{NameId:kube-dal6mt0s03rqu2pe3mrg,SystemId:r026-fcd88d9c-5f39-4f96-819e-05a8fc86a7de}}}]}"
      },
      {
        "key": "NodeGroupList",
        "value": "{IId:{NameId:workers1,SystemId:dal6mt0s03rqu2pe3mrg-f09c430},ImageIID:{NameId:Not visible in IBM,SystemId:Not visible in IBM},VMSpecName:cxf.4x8,RootDiskType:Not visible in IBM,RootDiskSize:Not visible in IBM,KeyPairIID:{NameId:Not visible in IBM,SystemId:Not visible in IBM},OnAutoScaling:false,DesiredNodeSize:2,MinNodeSize:-1,MaxNodeSize:-1,Status:Active,Nodes:[{NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-0000019b},{NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-000002b1}],KeyValueList:[{Key:IId,Value:{NameId:workers1,SystemId:dal6mt0s03rqu2pe3mrg-f09c430}},{Key:ImageIID,Value:{NameId:Not visible in IBM,SystemId:Not visible in IBM}},{Key:VMSpecName,Value:cxf.4x8},{Key:RootDiskType,Value:Not visible in IBM},{Key:RootDiskSize,Value:Not visible in IBM},{Key:KeyPairIID,Value:{NameId:Not visible in IBM,SystemId:Not visible in IBM}},{Key:OnAutoScaling,Value:false},{Key:DesiredNodeSize,Value:2},{Key:MinNodeSize,Value:0},{Key:MaxNodeSize,Value:0},{Key:Status,Value:Active},{Key:Nodes,Value:{NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-0000019b}; {NameId:Not visible in IBM,SystemId:kube-dal6mt0s03rqu2pe3mrg-tbj7i7ch33k-workers-000002b1}}]}"
      },
      {
        "key": "AccessInfo",
        "value": "{Endpoint:https://c104.au-syd.containers.cloud.ibm.com:31606,Kubeconfig:apiVersion: v1\\nclusters:\\n- cluster:\\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURRRENDQWlpZ0F3SUJBZ0lVYzEwNit1NWxCMGtLUE5xR2hJS0svcDhOeFZZd0RRWUpLb1pJaHZjTkFRRUwKQlFBd09ERTJNRFFHQTFVRUF4TXRaR0ZzTm0xME1ITXdNM0p4ZFRKd1pUTnRjbWN0YTNWaVpYSnVaWFJsY3kxagpZUzB4TnpnNU5UVXpOVGt3TUI0WERUSTJNRGt4TmpFd01EZ3dNRm9YRFRNMk1Ea3hNekV3TURnd01Gb3dPREUyCk1EUUdBMVVFQXhNdFpHRnNObTEwTUhNd00zSnhkVEp3WlROdGNtY3RhM1ZpWlhKdVpYUmxjeTFqWVMweE56ZzUKTlRVek5Ua3dNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXNWL20rQ1VNTlpiZQpIdFdReTRXdmZBS0d4eVZmR2h3TEZ5R0hyZVBwRWZCSGdwbnlHMlNydVhSbGlqbVE4L0JBblA1a2VtUm5ScGp4CnR5TGZmYmxTeEwzRHpNSXhNclFqd01MaFp2MWVkM3lJeXAzVzQ0TE5aTVd6WitrUndyWTNKN0cxWTU5S2orNGwKVVpjajZETTVORXFocFhEelhtTEtVWHN2dUwxNTVoQXZHV0ljY0NuK0NFR2F6L245NlROYjJhV3A5VTVBVUtUWQphZVE5TFRTTEdMbm41bGVKLzQ5YzVsMmxzR0FCaWc3amJuM2lLQmNMdlJ3UFhpM0c4U2E1czh0ZXlLWUFDTjI2CnB6cGh2T00ySlltWkNCbzRxV1JNRUt2ZE04bkk1REZpOHI1M3RoVjdyRzVJb091Q1NCaHYyMllrN3Jad1FPNlIKMy9xcnhBczlFd0lEQVFBQm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FZWXdEd1lEVlIwVEFRSC9CQVV3QXdFQgovekFkQmdOVkhRNEVGZ1FVMUFhZTE2NllnelppdEVzM3V4ck5wUFFBdW9vd0RRWUpLb1pJaHZjTkFRRUxCUUFECmdnRUJBSWoxMW5LbFd1cDI0UEJMSDUvN3RFNDNWNGFtQXgyR3JOaGUycFk3NVN6Nm05MzJCMGdYT2R6N3dLZkQKUVhCUFJ1SmV3S0tnZUpudmFXNU9oS1kzeG1XWVhQZ2xvL202L1R4MjV1MXR6Q2lkZE8weVpwenQzZlRsVE0yVQp2WERyVlRvd1JnazVrR1FlcUQwOHFKbG5WVFJGcVRiOW9TRURkOTc5OHRVTHlDZzUzN0hlWjZ4M1BEcTMrZWJkCmdzMGcxZUdrb1BEUGRUVVk5WlphaGRIdDltV0EwaHNOQ3lielptU2h4VHVINmFrQTVobEQrR1ZNS1hUZ3ZMZUEKY0IxWDFnYzJOY0tKZTRmeW5teTI3S0VTcU1mL0tRaXJTR2tjNEJ4Y1o5eGNiUlUyQ002YjVlY0c5Q2JTM2dsegowY3FpanFvUjNuUEhreUN1c3l3RFYzOTJmOTQ9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\\n    server: https://c104.au-syd.containers.cloud.ibm.com:31606\\n  name: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg\\ncontexts:\\n- context:\\n    cluster: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg\\n    namespace: default\\n    user: admin/dal6mt0s03rqu2pe3mrg\\n  name: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg/admin\\ncurrent-context: tbj7i7ch33krh023o26f/dal6mt0s03rqu2pe3mrg/admin\\nkind: Config\\nusers:\\n- name: admin/dal6mt0s03rqu2pe3mrg\\n  user:\\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURzRENDQXBpZ0F3SUJBZ0lSQUt0em9vaWZ5dXVZMW94YmlGdjYxc2d3RFFZSktvWklodmNOQVFFTEJRQXcKT0RFMk1EUUdBMVVFQXhNdFpHRnNObTEwTUhNd00zSnhkVEp3WlROdGNtY3RhM1ZpWlhKdVpYUmxjeTFqWVMweApOemc1TlRVek5Ua3dNQjRYRFRJMk1Ea3hOakV3TWpZd05Gb1hEVEk0TURreE5qRXdNall3TkZvd2F6RUxNQWtHCkExVUVCaE1DVlZNeEZqQVVCZ05WQkFnVERWTmhiaUJHY21GdVkybHpZMjh4Q3pBSkJnTlZCQWNUQWtOQk1STXcKRVFZRFZRUUtFd3BwWW0wdFlXUnRhVzV6TVNJd0lBWURWUVFEREJsTGRXSmxRMlZ5ZENOSlFrMXBaQzAyT1RRdwpNREZDVHpkWE1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBOFpYbHpaeDlrd0xOCjRVQlBTRTcvaVFHVWFZQmducDFMZlNmVWRTNTlHQmRUMm9XQXNrWTdLV2VPZlFaT3UycitFRXJnaEZKR0JoQlgKc0pYTjVQVDVyMWZyWit1TStFdGhBRHBBMXpVZU5GRldNall5Qjdta1BHNFVhZ2N6b3BUblM3NFBuc2VRRmVtRwptOHI2NE1CYjRnVHpDNzNCTWVFd01wbXpPRVlQWkkzcmFoeks3VWFvMVUvVTg1UlRUYnkyTkZCSjdtY1l1SWNGClp3c0RWQ0dnNWVqbW45eFVMajA3c3RNSFVCZ2RIU2pCcGgzbmVIekZIZVFkYm9acUpCZFFJcG9QUHNZcVZrdUMKWE5NajFDaEJ2YXJxR2lZeVRlTkdZbllZbzRER2VvdGJWWUt3U0I4Y1hWbUMrWHJuRHhNb3RjYzE3TVZTR2xYeAp5YU84NFpWS0J3SURBUUFCbzRHQk1IOHdEZ1lEVlIwUEFRSC9CQVFEQWdLRU1CMEdBMVVkSlFRV01CUUdDQ3NHCkFRVUZCd01DQmdnckJnRUZCUWNEQVRBTUJnTlZIUk1CQWY4RUFqQUFNQjhHQTFVZEl3UVlNQmFBRk5RR250ZXUKbUlNMllyUkxON3NhemFUMEFMcUtNQjhHQTFVZEVRUVlNQmFCRkdoaGJtbDZZVzVuTnpkQVoyMWhhV3d1WTI5dApNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUJDMTVuVDBIWDZpclF4Tm1KYXE5ZWNOUVZXU1BkV2NRTmt1WUpKCkpGd1Zsb2tGSVExbSs3RWY0RmRhMERDbUE2aDJyQ0FPVE1HeXZrS0YyQzJTN05qN0RKcjM1ZitEcmRid2VHUUkKUnFWQTA2Zm1Va1BCWERMNDBsUDVFeXNaQW1CbGNnNUxvNFliYlRqNk5DMnZWTEYvbEhvWTFlcURoV0R3Qi91cAo5dHJnbGlzWXdXUC9kQUtEb2hhcmI0ZnNYb0pWWHU0WXhtVjB0c21DS041dGhlYkVJb1FEa0RqM1djWmxaOXJICkI0SEhKclh1LzlQUm8yTGdINUMvMytTY3V2Qk0yUXpJSDNnanZJdThKbWRKdWpEbGZ0NjVNdllMMzZIc2lGeGcKWkhPQUFGM3NON2h2bVpIZ1JWcjlRTTlyRWpTWXJweEI2N1lBOGRmbzRhNmhabmJ4Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBOFpYbHpaeDlrd0xONFVCUFNFNy9pUUdVYVlCZ25wMUxmU2ZVZFM1OUdCZFQyb1dBCnNrWTdLV2VPZlFaT3UycitFRXJnaEZKR0JoQlhzSlhONVBUNXIxZnJaK3VNK0V0aEFEcEExelVlTkZGV01qWXkKQjdta1BHNFVhZ2N6b3BUblM3NFBuc2VRRmVtR204cjY0TUJiNGdUekM3M0JNZUV3TXBtek9FWVBaSTNyYWh6Swo3VWFvMVUvVTg1UlRUYnkyTkZCSjdtY1l1SWNGWndzRFZDR2c1ZWptbjl4VUxqMDdzdE1IVUJnZEhTakJwaDNuCmVIekZIZVFkYm9acUpCZFFJcG9QUHNZcVZrdUNYTk1qMUNoQnZhcnFHaVl5VGVOR1luWVlvNERHZW90YlZZS3cKU0I4Y1hWbUMrWHJuRHhNb3RjYzE3TVZTR2xYeHlhTzg0WlZLQndJREFRQUJBb0lCQUE5MXlra0lOTGtOdXVNcApYR21DTkxRdDE3T1F0WjR3N3IzSnFMei9CcDVlRDgyeU1YUTNMbDROOUg1bnd1NFhnTDdHSyt3TDM5TlBoRzBXCmlTQ1gxTXExMDZqSTJES2prRWVWY3NUUzcyWGx0cUJyKzNPbkc2MktWZUFiS2VERHFyR2NMaCs1SWExbFRtbjEKNld3c296U3BvR3dsN3BFa21oTUM1d2M0NUk5SXJlMGw3YmZQeS9UaDhzbWs1T2sxYVBoZ0xvZHhBMGIzTkRPSQp0ZGU5czlKRGRTOU1zY3pBUGxza3I0bUx3N2w0WEYzWW5tUXc5UjY5RnRJQXlCMDJOa05Zbi9zTHA1d1cxdHlDCkRhL1FYM0lSdlBqOUZaa2NZb1ZHOGQxQ2xuRWlWTnJzRE5BK1FrdXN3SjF5TC9QamZvSnlsY1NydTZxUWNBLysKU0NaZmwzRUNnWUVBKzNHaXE1ZDhZKzBSL0VrNnUzZm5ZNDRHRGhpMjBwdE04M3oyM09xWURaekY3MFBNaHJxQwpLOXZzYlJ1MHcvQXBTY2RVRzV5ekJ3NDdpdXJVeXBIY1JjTDlFYkdXU0Z2NTllRGZ1bzBLQTlBMCtldnN1WWFxCnpvdXIwem5WVjBhNUI5Y2F4OW04NnF4SmJmQXhMNS9oNkhqQnRFVmZ0elZWcjhZR0l2WTJRaE1DZ1lFQTlmYUkKV2xDeFFkQlI0ekJwalRWSGsvbVpoNnM3cFBpUzMxUWNINHhibEFVdXVETE9iRStYdHFNVFhTYWl3SUJmYnJPaQpXYXBEZjFTK0ZjNFFlTUFjekFTY0xVMzQ2RGd2YmorV2YwS1krZ1pLNXhvek5iNnRFNEsxMEhMTVNyZTk4SWg3CnMyNVJiN1kvamFOc3lqcmFyQ2p2QjkxRHBYOU5mdWZDdEY3ZnRyMENnWUVBMG9xSVIxNEZEamNJQkZQZEZmU0UKajl2d1BnVjdzRVhSM1dBWjVVbWFJR3ZSWVZOSUF0aFEveUNiaTVEVGYrMnM1TlkvR3cvTzZHMkdkZi9FUmdwMApnd1dPbWk0MVJFbWZ0NzZnRjdqWlZmQVZLOS9jekV4eTRaZ2FQRGdFNTV1VWUzZ21PSW1kb25LNDJaRngzZ3JtClFwNDZ0QlFTM1htUFVpdGlJQXhCeW5rQ2dZQWxnL0RRTmJhVG56NmVOR2dsRFpkWlRweklRS25jUTczREtvVVAKbXN6dENzMVJjdzVoSHRLNUhLNTdhc1V3TDJSZThpODFGZTh0b0xOTmlCeWpEa3BXSSszZVN5Skg2U255MnVnTgppUTdrTThtQTdsSVpSSGdKbmNvMWZRMEQ2SHFrRVcwc3RRcmV5eUZ1YlJyT3phTkUxd2wrWFpWUHpOYjVJRWhtClVvSTAwUUtCZ1FEWkhTTllUbjZETWM5R0N2S3ZKNTRGME43d3hXYzcyak9JYkVXUWNXUUw0Tk5hNW5OVkRmRlkKMWQ0aktSOFg2UmZ3Tzg3QW5qaTBvdWc5YUs4ZitFL2J2cUdiK0lubjR3N0pHTlJ1dWJOM0JwQ0xEbjR3cmg3KwpTNk5keHBWMTFBWVpqRkN1TmhJL3k5WDBoZk9USHNiZ1IzeGpBaGJxcFUrK3dMUXY1ZWR6SFE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\\n}"
      },
      {
        "key": "Addons",
        "value": "{KeyValueList:[{Key:cluster-autoscaler,Value:{\\allowed_upgrade_versions\\:[\\2.0.0\\],\\installOptionsTemplate\\:{},\\name\\:\\cluster-autoscaler\\,\\targetVersion\\:\\1.2.4\\,\\version\\:\\1.2.4\\}},{Key:ibm-storage-operator,Value:{\\healthStatus\\:\\Enabling\\,\\installOptionsTemplate\\:{},\\name\\:\\ibm-storage-operator\\,\\targetVersion\\:\\1.0\\,\\version\\:\\1.0\\}},{Key:vpc-block-csi-driver,Value:{\\allowed_upgrade_versions\\:[\\5.2\\],\\healthStatus\\:\\Enabling\\,\\installOptionsTemplate\\:{},\\name\\:\\vpc-block-csi-driver\\,\\targetVersion\\:\\5.1\\,\\version\\:\\5.1\\}}]}"
      },
      {
        "key": "Status",
        "value": "Active"
      },
      {
        "key": "CreatedTime",
        "value": "2026-09-16T10:12:03Z"
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

