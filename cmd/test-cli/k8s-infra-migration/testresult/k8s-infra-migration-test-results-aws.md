# CM-Beetle K8s Infra Migration Test Results — AWS-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → migrate → list → get (verified against
> the recommendation) → delete → residual resource check.

## Environment

- CSP / Region: aws / ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.0
- Git Commit: 803afb4
- Namespace: mig01
- Test Date: 2026-08-20 15:15:35 KST
- Cluster ID: mig01-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 931ms |
| 2 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 7m30.093s |
| 3 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 1ms |
| 4 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 1.515s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 1m43.493s |
| 6 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 9m16.536s |
| 7 | Residual resource check (Tumblebug) | ✅ **PASS** | 3ms |

**Overall Result**: 7/7 steps passed ✅

**Total Duration**: 18m32s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 931ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.33)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge nodes=2

### Step 2 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 7m30.093s
- **Status Code**: 202

- ℹ️  nameSeed: mig01
- ℹ️  async reqId: 1787206536789845713
- ℹ️  cluster id: mig01-on-prem-k8s-cluster
- ℹ️  elapsed: 7m30s
- ✅ status: Active

### Step 3 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 1ms
- **Status Code**: 200

- ✅ migrated cluster present in list (3 total)

### Step 4 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 1.515s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=aws+ap-northeast-2+c5a.xlarge, nodes=2)
- ✅ version: 1.33 (recommended 1.33)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 1m43.493s

- ✅ kubeconfig obtained (server: https://A1E5A99B85D7AEE154F241C6E2E2D464.yl4.ap-northeast-2.eks.amazonaws.com)
- ℹ️  auth method: exec credential plugin
- ✅ cluster token obtained from Tumblebug
- ✅ API server reachable (v1.33.13-eks-bca9cf6)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 1)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: a744452d05ccb4780ba3bb83a48a7239-1969825577.ap-northeast-2.elb.amazonaws.com
- ✅ nginx served over the LoadBalancer at http://a744452d05ccb4780ba3bb83a48a7239-1969825577.ap-northeast-2.elb.amazonaws.com/ (attempt 6)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 6 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 9m16.536s
- **Status Code**: 200

- ✅ deleted on attempt 1 (9m16s)

### Step 7 — Residual resource check (Tumblebug)

- **Duration**: 3ms

- ℹ️  VNet mig01-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig01-k8s-sg still exists (known gap)
- ℹ️  SshKey mig01-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.33)",
  "targetCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
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
    "connectionName": "aws-ap-northeast-2",
    "cidrBlock": "10.0.0.0/22",
    "subnetInfoList": [
      {
        "name": "k8s-subnet-a",
        "ipv4_CIDR": "10.0.1.0/24",
        "zone": "ap-northeast-2a"
      },
      {
        "name": "k8s-subnet-b",
        "ipv4_CIDR": "10.0.2.0/24",
        "zone": "ap-northeast-2b"
      }
    ],
    "description": "VPC for migrated K8s cluster"
  },
  "targetSshKey": {
    "name": "k8s-sshkey",
    "connectionName": "aws-ap-northeast-2",
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
      "connectionName": "aws-ap-northeast-2",
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
    "connectionName": "aws-ap-northeast-2",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "1.33",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": [
      {
        "name": "workers1",
        "imageId": "default",
        "specId": "aws+ap-northeast-2+c5a.xlarge",
        "rootDiskType": "default",
        "rootDiskSize": 100,
        "sshKeyId": "",
        "onAutoScaling": "false",
        "desiredNodeSize": 2,
        "minNodeSize": 2,
        "maxNodeSize": 2,
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
  "id": "mig01-on-prem-k8s-cluster",
  "uid": "tbdlsc18uvrh6drbdv70",
  "name": "mig01-on-prem-k8s-cluster",
  "connectionName": "aws-ap-northeast-2",
  "connectionConfig": {
    "configName": "aws-ap-northeast-2",
    "providerName": "aws",
    "driverName": "aws-driver-v1.0.so",
    "credentialName": "aws",
    "credentialHolder": "admin",
    "regionZoneInfoName": "aws-ap-northeast-2",
    "regionZoneInfo": {
      "assignedRegion": "ap-northeast-2",
      "assignedZone": "ap-northeast-2a"
    },
    "regionDetail": {
      "regionId": "ap-northeast-2",
      "regionName": "ap-northeast-2",
      "description": "Asia Pacific (Seoul)",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "zones": [
        "ap-northeast-2a",
        "ap-northeast-2b",
        "ap-northeast-2c",
        "ap-northeast-2d"
      ]
    },
    "regionRepresentative": true,
    "verified": true
  },
  "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
  "systemMessage": "",
  "label": {
    "Name": "tbdlsc18uvrh6drbdv70",
    "sys.connectionName": "aws-ap-northeast-2",
    "sys.createdTime": "2026-08-20 06:15:41.284 +0000 UTC",
    "sys.cspResourceId": "tbdlsc18uvrh6drbdv70",
    "sys.cspResourceName": "tbdlsc18uvrh6drbdv70",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig01-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig01-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbdlsc18uvrh6drbdv70",
    "sys.version": "1.33"
  },
  "systemLabel": "",
  "version": "1.33",
  "network": {
    "vNetId": "mig01-k8s-vpc",
    "subnetIds": [
      "mig01-k8s-subnet-a",
      "mig01-k8s-subnet-b"
    ],
    "securityGroupIds": [
      "mig01-k8s-sg"
    ],
    "keyValueList": [
      {
        "key": "ClusterSecurityGroupId",
        "value": "sg-0335b04180c79d7ca"
      },
      {
        "key": "EndpointPrivateAccess",
        "value": "false"
      },
      {
        "key": "EndpointPublicAccess",
        "value": "true"
      },
      {
        "key": "PublicAccessCidrs",
        "value": "0.0.0.0/0"
      },
      {
        "key": "SecurityGroupIds",
        "value": "sg-00b0fa9c3c57c2cc5"
      },
      {
        "key": "SubnetIds",
        "value": "subnet-05247863bd834f2df; subnet-0b020a2249bb6475b"
      },
      {
        "key": "VpcId",
        "value": "vpc-01af9f247d39edce6"
      }
    ]
  },
  "k8sNodeGroupList": [
    {
      "id": "workers1",
      "name": "workers1",
      "imageId": "default",
      "specId": "aws+ap-northeast-2+c5a.xlarge",
      "rootDiskType": "",
      "rootDiskSize": 100,
      "sshKeyId": "mig01-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 2,
      "maxNodeSize": 2,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "i-031447656bc5f8abf",
          "cspResourceId": "i-031447656bc5f8abf"
        },
        {
          "cspResourceName": "i-0e3c3fffd635b89bb",
          "cspResourceId": "i-0e3c3fffd635b89bb"
        }
      ],
      "keyValueList": [
        {
          "key": "IId",
          "value": "{NameId:workers1,SystemId:workers1}"
        },
        {
          "key": "ImageIID",
          "value": "{NameId:AL2023_x86_64_STANDARD,SystemId:}"
        },
        {
          "key": "VMSpecName",
          "value": "c5a.xlarge"
        },
        {
          "key": "RootDiskSize",
          "value": "100"
        },
        {
          "key": "KeyPairIID",
          "value": "{NameId:,SystemId:tbdn1fr85uqtviuhtgle}"
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
          "value": "2"
        },
        {
          "key": "MaxNodeSize",
          "value": "2"
        },
        {
          "key": "Status",
          "value": "Active"
        },
        {
          "key": "Nodes",
          "value": "{NameId:,SystemId:i-031447656bc5f8abf}; {NameId:,SystemId:i-0e3c3fffd635b89bb}"
        },
        {
          "key": "KeyValueList",
          "value": "{Key:AmiType,Value:AL2023_x86_64_STANDARD}; {Key:CapacityType,Value:ON_DEMAND}; {Key:ClusterName,Value:tbdlsc18uvrh6drbdv70}; {Key:CreatedAt,Value:2026-08-20T06:21:32.328Z}; {Key:DiskSize,Value:100}; {Key:Health,Value:{Issues:[]}}; {Key:InstanceTypes,Value:c5a.xlarge}; {Key:ModifiedAt,Value:2026-08-20T06:22:43.889Z}; {Key:NodeRole,Value:arn:aws:iam::635484366616:role/cloud-barista-eks-nodegroup-role}; {Key:NodegroupArn,Value:arn:aws:eks:ap-northeast-2:635484366616:nodegroup/tbdlsc18uvrh6drbdv70/workers1/5cd00eea-7860-d7b8-7846-0ab62961a4b5}; {Key:NodegroupName,Value:workers1}; {Key:ReleaseVersion,Value:1.33.13-20260818}; {Key:RemoteAccess,Value:{Ec2SshKey:tbdn1fr85uqtviuhtgle,SourceSecurityGroups:[sg-00b0fa9c3c57c2cc5]}}; {Key:Resources,Value:{AutoScalingGroups:[{Name:eks-workers1-5cd00eea-7860-d7b8-7846-0ab62961a4b5}],RemoteAccessSecurityGroup:sg-074607780cfdd6b92}}; {Key:ScalingConfig,Value:{DesiredSize:2,MaxSize:2,MinSize:2}}; {Key:Status,Value:ACTIVE}; {Key:Subnets,Value:subnet-05247863bd834f2df; subnet-0b020a2249bb6475b}; {Key:Tags,Value:{key:nodegroup,value:workers1}}; {Key:UpdateConfig,Value:{MaxUnavailable:1,MaxUnavailablePercentage:null}}; {Key:Version,Value:1.33}"
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
          "NameId": "",
          "SystemId": ""
        },
        "VMSpecName": "c5a.xlarge",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tbdn1fr85uqtviuhtgle",
          "SystemId": "tbdn1fr85uqtviuhtgle"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 2,
        "MaxNodeSize": 2,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "i-031447656bc5f8abf",
            "SystemId": "i-031447656bc5f8abf"
          },
          {
            "NameId": "i-0e3c3fffd635b89bb",
            "SystemId": "i-0e3c3fffd635b89bb"
          }
        ],
        "KeyValueList": [
          {
            "key": "IId",
            "value": "{NameId:workers1,SystemId:workers1}"
          },
          {
            "key": "ImageIID",
            "value": "{NameId:AL2023_x86_64_STANDARD,SystemId:}"
          },
          {
            "key": "VMSpecName",
            "value": "c5a.xlarge"
          },
          {
            "key": "RootDiskSize",
            "value": "100"
          },
          {
            "key": "KeyPairIID",
            "value": "{NameId:,SystemId:tbdn1fr85uqtviuhtgle}"
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
            "value": "2"
          },
          {
            "key": "MaxNodeSize",
            "value": "2"
          },
          {
            "key": "Status",
            "value": "Active"
          },
          {
            "key": "Nodes",
            "value": "{NameId:,SystemId:i-031447656bc5f8abf}; {NameId:,SystemId:i-0e3c3fffd635b89bb}"
          },
          {
            "key": "KeyValueList",
            "value": "{Key:AmiType,Value:AL2023_x86_64_STANDARD}; {Key:CapacityType,Value:ON_DEMAND}; {Key:ClusterName,Value:tbdlsc18uvrh6drbdv70}; {Key:CreatedAt,Value:2026-08-20T06:21:32.328Z}; {Key:DiskSize,Value:100}; {Key:Health,Value:{Issues:[]}}; {Key:InstanceTypes,Value:c5a.xlarge}; {Key:ModifiedAt,Value:2026-08-20T06:22:43.889Z}; {Key:NodeRole,Value:arn:aws:iam::635484366616:role/cloud-barista-eks-nodegroup-role}; {Key:NodegroupArn,Value:arn:aws:eks:ap-northeast-2:635484366616:nodegroup/tbdlsc18uvrh6drbdv70/workers1/5cd00eea-7860-d7b8-7846-0ab62961a4b5}; {Key:NodegroupName,Value:workers1}; {Key:ReleaseVersion,Value:1.33.13-20260818}; {Key:RemoteAccess,Value:{Ec2SshKey:tbdn1fr85uqtviuhtgle,SourceSecurityGroups:[sg-00b0fa9c3c57c2cc5]}}; {Key:Resources,Value:{AutoScalingGroups:[{Name:eks-workers1-5cd00eea-7860-d7b8-7846-0ab62961a4b5}],RemoteAccessSecurityGroup:sg-074607780cfdd6b92}}; {Key:ScalingConfig,Value:{DesiredSize:2,MaxSize:2,MinSize:2}}; {Key:Status,Value:ACTIVE}; {Key:Subnets,Value:subnet-05247863bd834f2df; subnet-0b020a2249bb6475b}; {Key:Tags,Value:{key:nodegroup,value:workers1}}; {Key:UpdateConfig,Value:{MaxUnavailable:1,MaxUnavailablePercentage:null}}; {Key:Version,Value:1.33}"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://A1E5A99B85D7AEE154F241C6E2E2D464.yl4.ap-northeast-2.eks.amazonaws.com",
    "kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://A1E5A99B85D7AEE154F241C6E2E2D464.yl4.ap-northeast-2.eks.amazonaws.com\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREVENDQWZXZ0F3SUJBZ0lRS3pTV1A3ODE1VTFzdUpJOXFzWE1TakFOQmdrcWhraUc5dzBCQVFzRkFEQVYKTVJNd0VRWURWUVFERXdwcmRXSmxjbTVsZEdWek1CNFhEVEkyTURneU1EQTJNVFUxTTFvWERUTXhNRGd4T1RBMgpNVFUxTTFvd0ZURVRNQkVHQTFVRUF4TUthM1ZpWlhKdVpYUmxjekNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFECmdnRVBBRENDQVFvQ2dnRUJBT3czVjd0Zk44WU1FR1E1REE2Q0dvNEhMZGJnc3hJRzZNa29qV3RycWZsZGYwb3IKOWRZc1VDc3NISGhwRklWcmNOdmkwTmJ1UVlqQlBnMmFaZUwrNXlUU05SV1pOdXJFcHZiZ1h1MDFLVlJwdVBIdwpSQjdZQUh3cllRUSt4b0JtU1dVL3ZUaVh0SzRQRWhydjU5cjlhT0l3S0p5OTFiRGV6QXhVdFlVTld3TUNEU3hOClhSTk5OdC91SC9CaDN1YnN4Um9jVUlzN3VOWnViVjBhcDNpc0NJbExRdENYeTgyRVRQRmE5OU5xa0g2Rlhwc1IKRXhGTzhWSnQ2eW1KR1hmdTBvTVFBTjYxc1RzUkwxd1BUWmxEZjJhSFdJRmxwL2dDbSt4Z2xjUnkwNVRCZkV2dwphV1BUQ0FCdE4wVTR6Y3crYU5sUmFqU29NZTlGUlFVbjhlSzVzVVVDQXdFQUFhTlpNRmN3RGdZRFZSMFBBUUgvCkJBUURBZ0trTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRkRZVU1KdzBYNGNmUTBpV2I0encKc21GSzYwcHVNQlVHQTFVZEVRUU9NQXlDQ210MVltVnlibVYwWlhNd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBS3NPSmVseWxsMlJ2TmNRTDgxbnRLMjhSRGxHSWxqNm85L3V0clhnU2YzNlFaRENiZzdIWjA3VUx4aDRxcUNvCnFTK3ZvMWVSMDZIbkZTOWJlTGFKSjJhMjhSUUhLT2NzdUFBWVNsVTcxQ0taUjcxMWlRSlQxMTRRM2VkSW9qWGYKYVlHWCs0Q1lvc0lpazdmd2hNSkxzRWM5aTRIdnRlOHNYWGtHZDQ4RncxZHlOdk5lK0J2YW1jaXk1bHd2cXRtQQpJVmlVMVdnYXl6YzYraEkrTnZyTFpVSFc0a00ybXZwa0xCWXNVdElyL05kaXBmYTZjMEsyYjg3RHJtUU9nYjNMCk5QSlZxbGQvZnByNkExT3lMT1UwdlhCNjJEWktLdDJpcWdWcUFXRG9tY2FNKzZZUDhMMmJkc3dITlVIbEczQzUKSFNZVWhtUVR5akg1TkNEWXAxRjc4eXc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: tbdlsc18uvrh6drbdv70\ncontexts:\n- context:\n    cluster: tbdlsc18uvrh6drbdv70\n    user: aws-dynamic-token\n  name: tbdlsc18uvrh6drbdv70\ncurrent-context: tbdlsc18uvrh6drbdv70\nusers:\n- name: aws-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tbdlsc18uvrh6drbdv70/token?ConnectionName=aws-ap-northeast-2\\\"\"\n"
  },
  "addons": {
    "keyValueList": [
      {
        "key": "AddonArn",
        "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tbdlsc18uvrh6drbdv70/aws-ebs-csi-driver/c4d00ee7-d2c7-dc8b-57d9-c1e0e7d04dc5"
      },
      {
        "key": "AddonName",
        "value": "aws-ebs-csi-driver"
      },
      {
        "key": "AddonVersion",
        "value": "v1.64.0-eksbuild.1"
      },
      {
        "key": "ClusterName",
        "value": "tbdlsc18uvrh6drbdv70"
      },
      {
        "key": "CreatedAt",
        "value": "2026-08-20T06:15:42.923Z"
      },
      {
        "key": "Health",
        "value": "{Issues:[]}"
      },
      {
        "key": "ModifiedAt",
        "value": "2026-08-20T06:15:42.94Z"
      },
      {
        "key": "Status",
        "value": "CREATING"
      },
      {
        "key": "AddonArn",
        "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tbdlsc18uvrh6drbdv70/eks-pod-identity-agent/aed00ee7-d395-8f71-c953-1ddd85bae1aa"
      },
      {
        "key": "AddonName",
        "value": "eks-pod-identity-agent"
      },
      {
        "key": "AddonVersion",
        "value": "v1.3.10-eksbuild.3"
      },
      {
        "key": "ClusterName",
        "value": "tbdlsc18uvrh6drbdv70"
      },
      {
        "key": "CreatedAt",
        "value": "2026-08-20T06:15:43.325Z"
      },
      {
        "key": "Health",
        "value": "{Issues:[]}"
      },
      {
        "key": "ModifiedAt",
        "value": "2026-08-20T06:15:43.34Z"
      },
      {
        "key": "Status",
        "value": "CREATING"
      }
    ]
  },
  "status": "Active",
  "createdTime": "2026-08-20T06:15:41.284Z",
  "keyValueList": [
    {
      "key": "Arn",
      "value": "arn:aws:eks:ap-northeast-2:635484366616:cluster/tbdlsc18uvrh6drbdv70"
    },
    {
      "key": "CertificateAuthority",
      "value": "{Data:LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREVENDQWZXZ0F3SUJBZ0lRS3pTV1A3ODE1VTFzdUpJOXFzWE1TakFOQmdrcWhraUc5dzBCQVFzRkFEQVYKTVJNd0VRWURWUVFERXdwcmRXSmxjbTVsZEdWek1CNFhEVEkyTURneU1EQTJNVFUxTTFvWERUTXhNRGd4T1RBMgpNVFUxTTFvd0ZURVRNQkVHQTFVRUF4TUthM1ZpWlhKdVpYUmxjekNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFECmdnRVBBRENDQVFvQ2dnRUJBT3czVjd0Zk44WU1FR1E1REE2Q0dvNEhMZGJnc3hJRzZNa29qV3RycWZsZGYwb3IKOWRZc1VDc3NISGhwRklWcmNOdmkwTmJ1UVlqQlBnMmFaZUwrNXlUU05SV1pOdXJFcHZiZ1h1MDFLVlJwdVBIdwpSQjdZQUh3cllRUSt4b0JtU1dVL3ZUaVh0SzRQRWhydjU5cjlhT0l3S0p5OTFiRGV6QXhVdFlVTld3TUNEU3hOClhSTk5OdC91SC9CaDN1YnN4Um9jVUlzN3VOWnViVjBhcDNpc0NJbExRdENYeTgyRVRQRmE5OU5xa0g2Rlhwc1IKRXhGTzhWSnQ2eW1KR1hmdTBvTVFBTjYxc1RzUkwxd1BUWmxEZjJhSFdJRmxwL2dDbSt4Z2xjUnkwNVRCZkV2dwphV1BUQ0FCdE4wVTR6Y3crYU5sUmFqU29NZTlGUlFVbjhlSzVzVVVDQXdFQUFhTlpNRmN3RGdZRFZSMFBBUUgvCkJBUURBZ0trTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRkRZVU1KdzBYNGNmUTBpV2I0encKc21GSzYwcHVNQlVHQTFVZEVRUU9NQXlDQ210MVltVnlibVYwWlhNd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBS3NPSmVseWxsMlJ2TmNRTDgxbnRLMjhSRGxHSWxqNm85L3V0clhnU2YzNlFaRENiZzdIWjA3VUx4aDRxcUNvCnFTK3ZvMWVSMDZIbkZTOWJlTGFKSjJhMjhSUUhLT2NzdUFBWVNsVTcxQ0taUjcxMWlRSlQxMTRRM2VkSW9qWGYKYVlHWCs0Q1lvc0lpazdmd2hNSkxzRWM5aTRIdnRlOHNYWGtHZDQ4RncxZHlOdk5lK0J2YW1jaXk1bHd2cXRtQQpJVmlVMVdnYXl6YzYraEkrTnZyTFpVSFc0a00ybXZwa0xCWXNVdElyL05kaXBmYTZjMEsyYjg3RHJtUU9nYjNMCk5QSlZxbGQvZnByNkExT3lMT1UwdlhCNjJEWktLdDJpcWdWcUFXRG9tY2FNKzZZUDhMMmJkc3dITlVIbEczQzUKSFNZVWhtUVR5akg1TkNEWXAxRjc4eXc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K}"
    },
    {
      "key": "CreatedAt",
      "value": "2026-08-20T06:15:41.284Z"
    },
    {
      "key": "Endpoint",
      "value": "https://A1E5A99B85D7AEE154F241C6E2E2D464.yl4.ap-northeast-2.eks.amazonaws.com"
    },
    {
      "key": "Identity",
      "value": "{Oidc:{Issuer:https://oidc.eks.ap-northeast-2.amazonaws.com/id/A1E5A99B85D7AEE154F241C6E2E2D464}}"
    },
    {
      "key": "KubernetesNetworkConfig",
      "value": "{ServiceIpv4Cidr:172.20.0.0/16}"
    },
    {
      "key": "Logging",
      "value": "{ClusterLogging:[{Enabled:false,Types:[api,audit,authenticator,controllerManager,scheduler]}]}"
    },
    {
      "key": "Name",
      "value": "tbdlsc18uvrh6drbdv70"
    },
    {
      "key": "PlatformVersion",
      "value": "eks.45"
    },
    {
      "key": "ResourcesVpcConfig",
      "value": "{ClusterSecurityGroupId:sg-0335b04180c79d7ca,EndpointPrivateAccess:false,EndpointPublicAccess:true,PublicAccessCidrs:[0.0.0.0/0],SecurityGroupIds:[sg-00b0fa9c3c57c2cc5],SubnetIds:[subnet-05247863bd834f2df,subnet-0b020a2249bb6475b],VpcId:vpc-01af9f247d39edce6}"
    },
    {
      "key": "RoleArn",
      "value": "arn:aws:iam::635484366616:role/cloud-barista-eks-cluster-role"
    },
    {
      "key": "Status",
      "value": "ACTIVE"
    },
    {
      "key": "Tags",
      "value": "{Name:tbdlsc18uvrh6drbdv70,sys.id:mig01-on-prem-k8s-cluster,sys.uid:tbdlsc18uvrh6drbdv70}"
    },
    {
      "key": "Version",
      "value": "1.33"
    }
  ],
  "cspResourceName": "tbdlsc18uvrh6drbdv70",
  "cspResourceId": "tbdlsc18uvrh6drbdv70",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbdlsc18uvrh6drbdv70",
      "SystemId": "tbdlsc18uvrh6drbdv70"
    },
    "Version": "1.33",
    "Network": {
      "VpcIID": {
        "NameId": "tbdmgt1mbasrbltfs1e6",
        "SystemId": "vpc-01af9f247d39edce6"
      },
      "SubnetIIDs": [
        {
          "NameId": "tbdoke7m88osrtd6dju7",
          "SystemId": "subnet-05247863bd834f2df"
        },
        {
          "NameId": "tbv8r1cvca834kh0g8lj",
          "SystemId": "subnet-0b020a2249bb6475b"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "tb73eajo8k7k082am4s1",
          "SystemId": "sg-00b0fa9c3c57c2cc5"
        }
      ],
      "KeyValueList": [
        {
          "key": "ClusterSecurityGroupId",
          "value": "sg-0335b04180c79d7ca"
        },
        {
          "key": "EndpointPrivateAccess",
          "value": "false"
        },
        {
          "key": "EndpointPublicAccess",
          "value": "true"
        },
        {
          "key": "PublicAccessCidrs",
          "value": "0.0.0.0/0"
        },
        {
          "key": "SecurityGroupIds",
          "value": "sg-00b0fa9c3c57c2cc5"
        },
        {
          "key": "SubnetIds",
          "value": "subnet-05247863bd834f2df; subnet-0b020a2249bb6475b"
        },
        {
          "key": "VpcId",
          "value": "vpc-01af9f247d39edce6"
        }
      ]
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "workers1"
        },
        "ImageIID": {
          "NameId": "",
          "SystemId": ""
        },
        "VMSpecName": "c5a.xlarge",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tbdn1fr85uqtviuhtgle",
          "SystemId": "tbdn1fr85uqtviuhtgle"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 2,
        "MaxNodeSize": 2,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "i-031447656bc5f8abf",
            "SystemId": "i-031447656bc5f8abf"
          },
          {
            "NameId": "i-0e3c3fffd635b89bb",
            "SystemId": "i-0e3c3fffd635b89bb"
          }
        ],
        "KeyValueList": [
          {
            "key": "IId",
            "value": "{NameId:workers1,SystemId:workers1}"
          },
          {
            "key": "ImageIID",
            "value": "{NameId:AL2023_x86_64_STANDARD,SystemId:}"
          },
          {
            "key": "VMSpecName",
            "value": "c5a.xlarge"
          },
          {
            "key": "RootDiskSize",
            "value": "100"
          },
          {
            "key": "KeyPairIID",
            "value": "{NameId:,SystemId:tbdn1fr85uqtviuhtgle}"
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
            "value": "2"
          },
          {
            "key": "MaxNodeSize",
            "value": "2"
          },
          {
            "key": "Status",
            "value": "Active"
          },
          {
            "key": "Nodes",
            "value": "{NameId:,SystemId:i-031447656bc5f8abf}; {NameId:,SystemId:i-0e3c3fffd635b89bb}"
          },
          {
            "key": "KeyValueList",
            "value": "{Key:AmiType,Value:AL2023_x86_64_STANDARD}; {Key:CapacityType,Value:ON_DEMAND}; {Key:ClusterName,Value:tbdlsc18uvrh6drbdv70}; {Key:CreatedAt,Value:2026-08-20T06:21:32.328Z}; {Key:DiskSize,Value:100}; {Key:Health,Value:{Issues:[]}}; {Key:InstanceTypes,Value:c5a.xlarge}; {Key:ModifiedAt,Value:2026-08-20T06:22:43.889Z}; {Key:NodeRole,Value:arn:aws:iam::635484366616:role/cloud-barista-eks-nodegroup-role}; {Key:NodegroupArn,Value:arn:aws:eks:ap-northeast-2:635484366616:nodegroup/tbdlsc18uvrh6drbdv70/workers1/5cd00eea-7860-d7b8-7846-0ab62961a4b5}; {Key:NodegroupName,Value:workers1}; {Key:ReleaseVersion,Value:1.33.13-20260818}; {Key:RemoteAccess,Value:{Ec2SshKey:tbdn1fr85uqtviuhtgle,SourceSecurityGroups:[sg-00b0fa9c3c57c2cc5]}}; {Key:Resources,Value:{AutoScalingGroups:[{Name:eks-workers1-5cd00eea-7860-d7b8-7846-0ab62961a4b5}],RemoteAccessSecurityGroup:sg-074607780cfdd6b92}}; {Key:ScalingConfig,Value:{DesiredSize:2,MaxSize:2,MinSize:2}}; {Key:Status,Value:ACTIVE}; {Key:Subnets,Value:subnet-05247863bd834f2df; subnet-0b020a2249bb6475b}; {Key:Tags,Value:{key:nodegroup,value:workers1}}; {Key:UpdateConfig,Value:{MaxUnavailable:1,MaxUnavailablePercentage:null}}; {Key:Version,Value:1.33}"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://A1E5A99B85D7AEE154F241C6E2E2D464.yl4.ap-northeast-2.eks.amazonaws.com",
      "Kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://A1E5A99B85D7AEE154F241C6E2E2D464.yl4.ap-northeast-2.eks.amazonaws.com\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREVENDQWZXZ0F3SUJBZ0lRS3pTV1A3ODE1VTFzdUpJOXFzWE1TakFOQmdrcWhraUc5dzBCQVFzRkFEQVYKTVJNd0VRWURWUVFERXdwcmRXSmxjbTVsZEdWek1CNFhEVEkyTURneU1EQTJNVFUxTTFvWERUTXhNRGd4T1RBMgpNVFUxTTFvd0ZURVRNQkVHQTFVRUF4TUthM1ZpWlhKdVpYUmxjekNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFECmdnRVBBRENDQVFvQ2dnRUJBT3czVjd0Zk44WU1FR1E1REE2Q0dvNEhMZGJnc3hJRzZNa29qV3RycWZsZGYwb3IKOWRZc1VDc3NISGhwRklWcmNOdmkwTmJ1UVlqQlBnMmFaZUwrNXlUU05SV1pOdXJFcHZiZ1h1MDFLVlJwdVBIdwpSQjdZQUh3cllRUSt4b0JtU1dVL3ZUaVh0SzRQRWhydjU5cjlhT0l3S0p5OTFiRGV6QXhVdFlVTld3TUNEU3hOClhSTk5OdC91SC9CaDN1YnN4Um9jVUlzN3VOWnViVjBhcDNpc0NJbExRdENYeTgyRVRQRmE5OU5xa0g2Rlhwc1IKRXhGTzhWSnQ2eW1KR1hmdTBvTVFBTjYxc1RzUkwxd1BUWmxEZjJhSFdJRmxwL2dDbSt4Z2xjUnkwNVRCZkV2dwphV1BUQ0FCdE4wVTR6Y3crYU5sUmFqU29NZTlGUlFVbjhlSzVzVVVDQXdFQUFhTlpNRmN3RGdZRFZSMFBBUUgvCkJBUURBZ0trTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRkRZVU1KdzBYNGNmUTBpV2I0encKc21GSzYwcHVNQlVHQTFVZEVRUU9NQXlDQ210MVltVnlibVYwWlhNd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBS3NPSmVseWxsMlJ2TmNRTDgxbnRLMjhSRGxHSWxqNm85L3V0clhnU2YzNlFaRENiZzdIWjA3VUx4aDRxcUNvCnFTK3ZvMWVSMDZIbkZTOWJlTGFKSjJhMjhSUUhLT2NzdUFBWVNsVTcxQ0taUjcxMWlRSlQxMTRRM2VkSW9qWGYKYVlHWCs0Q1lvc0lpazdmd2hNSkxzRWM5aTRIdnRlOHNYWGtHZDQ4RncxZHlOdk5lK0J2YW1jaXk1bHd2cXRtQQpJVmlVMVdnYXl6YzYraEkrTnZyTFpVSFc0a00ybXZwa0xCWXNVdElyL05kaXBmYTZjMEsyYjg3RHJtUU9nYjNMCk5QSlZxbGQvZnByNkExT3lMT1UwdlhCNjJEWktLdDJpcWdWcUFXRG9tY2FNKzZZUDhMMmJkc3dITlVIbEczQzUKSFNZVWhtUVR5akg1TkNEWXAxRjc4eXc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: tbdlsc18uvrh6drbdv70\ncontexts:\n- context:\n    cluster: tbdlsc18uvrh6drbdv70\n    user: aws-dynamic-token\n  name: tbdlsc18uvrh6drbdv70\ncurrent-context: tbdlsc18uvrh6drbdv70\nusers:\n- name: aws-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tbdlsc18uvrh6drbdv70/token?ConnectionName=aws-ap-northeast-2\\\"\"\n"
    },
    "Addons": {
      "KeyValueList": [
        {
          "key": "AddonArn",
          "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tbdlsc18uvrh6drbdv70/aws-ebs-csi-driver/c4d00ee7-d2c7-dc8b-57d9-c1e0e7d04dc5"
        },
        {
          "key": "AddonName",
          "value": "aws-ebs-csi-driver"
        },
        {
          "key": "AddonVersion",
          "value": "v1.64.0-eksbuild.1"
        },
        {
          "key": "ClusterName",
          "value": "tbdlsc18uvrh6drbdv70"
        },
        {
          "key": "CreatedAt",
          "value": "2026-08-20T06:15:42.923Z"
        },
        {
          "key": "Health",
          "value": "{Issues:[]}"
        },
        {
          "key": "ModifiedAt",
          "value": "2026-08-20T06:15:42.94Z"
        },
        {
          "key": "Status",
          "value": "CREATING"
        },
        {
          "key": "AddonArn",
          "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tbdlsc18uvrh6drbdv70/eks-pod-identity-agent/aed00ee7-d395-8f71-c953-1ddd85bae1aa"
        },
        {
          "key": "AddonName",
          "value": "eks-pod-identity-agent"
        },
        {
          "key": "AddonVersion",
          "value": "v1.3.10-eksbuild.3"
        },
        {
          "key": "ClusterName",
          "value": "tbdlsc18uvrh6drbdv70"
        },
        {
          "key": "CreatedAt",
          "value": "2026-08-20T06:15:43.325Z"
        },
        {
          "key": "Health",
          "value": "{Issues:[]}"
        },
        {
          "key": "ModifiedAt",
          "value": "2026-08-20T06:15:43.34Z"
        },
        {
          "key": "Status",
          "value": "CREATING"
        }
      ]
    },
    "Status": "Active",
    "CreatedTime": "2026-08-20T06:15:41.284Z",
    "KeyValueList": [
      {
        "key": "Arn",
        "value": "arn:aws:eks:ap-northeast-2:635484366616:cluster/tbdlsc18uvrh6drbdv70"
      },
      {
        "key": "CertificateAuthority",
        "value": "{Data:LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREVENDQWZXZ0F3SUJBZ0lRS3pTV1A3ODE1VTFzdUpJOXFzWE1TakFOQmdrcWhraUc5dzBCQVFzRkFEQVYKTVJNd0VRWURWUVFERXdwcmRXSmxjbTVsZEdWek1CNFhEVEkyTURneU1EQTJNVFUxTTFvWERUTXhNRGd4T1RBMgpNVFUxTTFvd0ZURVRNQkVHQTFVRUF4TUthM1ZpWlhKdVpYUmxjekNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFECmdnRVBBRENDQVFvQ2dnRUJBT3czVjd0Zk44WU1FR1E1REE2Q0dvNEhMZGJnc3hJRzZNa29qV3RycWZsZGYwb3IKOWRZc1VDc3NISGhwRklWcmNOdmkwTmJ1UVlqQlBnMmFaZUwrNXlUU05SV1pOdXJFcHZiZ1h1MDFLVlJwdVBIdwpSQjdZQUh3cllRUSt4b0JtU1dVL3ZUaVh0SzRQRWhydjU5cjlhT0l3S0p5OTFiRGV6QXhVdFlVTld3TUNEU3hOClhSTk5OdC91SC9CaDN1YnN4Um9jVUlzN3VOWnViVjBhcDNpc0NJbExRdENYeTgyRVRQRmE5OU5xa0g2Rlhwc1IKRXhGTzhWSnQ2eW1KR1hmdTBvTVFBTjYxc1RzUkwxd1BUWmxEZjJhSFdJRmxwL2dDbSt4Z2xjUnkwNVRCZkV2dwphV1BUQ0FCdE4wVTR6Y3crYU5sUmFqU29NZTlGUlFVbjhlSzVzVVVDQXdFQUFhTlpNRmN3RGdZRFZSMFBBUUgvCkJBUURBZ0trTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRkRZVU1KdzBYNGNmUTBpV2I0encKc21GSzYwcHVNQlVHQTFVZEVRUU9NQXlDQ210MVltVnlibVYwWlhNd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBS3NPSmVseWxsMlJ2TmNRTDgxbnRLMjhSRGxHSWxqNm85L3V0clhnU2YzNlFaRENiZzdIWjA3VUx4aDRxcUNvCnFTK3ZvMWVSMDZIbkZTOWJlTGFKSjJhMjhSUUhLT2NzdUFBWVNsVTcxQ0taUjcxMWlRSlQxMTRRM2VkSW9qWGYKYVlHWCs0Q1lvc0lpazdmd2hNSkxzRWM5aTRIdnRlOHNYWGtHZDQ4RncxZHlOdk5lK0J2YW1jaXk1bHd2cXRtQQpJVmlVMVdnYXl6YzYraEkrTnZyTFpVSFc0a00ybXZwa0xCWXNVdElyL05kaXBmYTZjMEsyYjg3RHJtUU9nYjNMCk5QSlZxbGQvZnByNkExT3lMT1UwdlhCNjJEWktLdDJpcWdWcUFXRG9tY2FNKzZZUDhMMmJkc3dITlVIbEczQzUKSFNZVWhtUVR5akg1TkNEWXAxRjc4eXc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K}"
      },
      {
        "key": "CreatedAt",
        "value": "2026-08-20T06:15:41.284Z"
      },
      {
        "key": "Endpoint",
        "value": "https://A1E5A99B85D7AEE154F241C6E2E2D464.yl4.ap-northeast-2.eks.amazonaws.com"
      },
      {
        "key": "Identity",
        "value": "{Oidc:{Issuer:https://oidc.eks.ap-northeast-2.amazonaws.com/id/A1E5A99B85D7AEE154F241C6E2E2D464}}"
      },
      {
        "key": "KubernetesNetworkConfig",
        "value": "{ServiceIpv4Cidr:172.20.0.0/16}"
      },
      {
        "key": "Logging",
        "value": "{ClusterLogging:[{Enabled:false,Types:[api,audit,authenticator,controllerManager,scheduler]}]}"
      },
      {
        "key": "Name",
        "value": "tbdlsc18uvrh6drbdv70"
      },
      {
        "key": "PlatformVersion",
        "value": "eks.45"
      },
      {
        "key": "ResourcesVpcConfig",
        "value": "{ClusterSecurityGroupId:sg-0335b04180c79d7ca,EndpointPrivateAccess:false,EndpointPublicAccess:true,PublicAccessCidrs:[0.0.0.0/0],SecurityGroupIds:[sg-00b0fa9c3c57c2cc5],SubnetIds:[subnet-05247863bd834f2df,subnet-0b020a2249bb6475b],VpcId:vpc-01af9f247d39edce6}"
      },
      {
        "key": "RoleArn",
        "value": "arn:aws:iam::635484366616:role/cloud-barista-eks-cluster-role"
      },
      {
        "key": "Status",
        "value": "ACTIVE"
      },
      {
        "key": "Tags",
        "value": "{Name:tbdlsc18uvrh6drbdv70,sys.id:mig01-on-prem-k8s-cluster,sys.uid:tbdlsc18uvrh6drbdv70}"
      },
      {
        "key": "Version",
        "value": "1.33"
      }
    ]
  }
}
```

</details>

