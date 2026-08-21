# CM-Beetle K8s Infra Migration Test Results — AWS-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → migrate → list → get (verified against
> the recommendation) → delete → residual resource check.

## Environment

- CSP / Region: aws / ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.0+ (1cfe558)
- Git Commit: 1cfe558
- Namespace: mig01
- Test Date: 2026-08-21 11:40:25 KST
- Cluster ID: mig01-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 23ms |
| 2 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 9m0.231s |
| 3 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 1ms |
| 4 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 1.108s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 1m28.876s |
| 6 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 9m5.627s |
| 7 | Residual resource check (Tumblebug) | ✅ **PASS** | 3ms |

**Overall Result**: 7/7 steps passed ✅

**Total Duration**: 19m35s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 23ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.33)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge nodes=2

### Step 2 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 9m0.231s
- **Status Code**: 202

- ℹ️  nameSeed: mig01
- ℹ️  async reqId: 1787280025463497959
- ℹ️  cluster id: mig01-on-prem-k8s-cluster
- ℹ️  elapsed: 9m0s
- ✅ status: Active

### Step 3 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 1ms
- **Status Code**: 200

- ✅ migrated cluster present in list (1 total)

### Step 4 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 1.108s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=aws+ap-northeast-2+c5a.xlarge, nodes=2)
- ✅ version: 1.33 (recommended 1.33)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 1m28.876s

- ✅ kubeconfig obtained (server: https://A42C0404A862C06B0E9998C1EC42D7E1.gr7.ap-northeast-2.eks.amazonaws.com)
- ℹ️  auth method: exec credential plugin
- ✅ cluster token obtained from Tumblebug
- ✅ API server reachable (v1.33.13-eks-bca9cf6)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 1)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: a3916bbf5c64b4340b952386c3f11f25-1313336685.ap-northeast-2.elb.amazonaws.com
- ✅ nginx served over the LoadBalancer at http://a3916bbf5c64b4340b952386c3f11f25-1313336685.ap-northeast-2.elb.amazonaws.com/ (attempt 5)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 6 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 9m5.627s
- **Status Code**: 200

- ✅ deleted on attempt 1 (9m5s)

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
  "uid": "tbqajlaplk30cjpsahc6",
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
    "Name": "tbqajlaplk30cjpsahc6",
    "sys.connectionName": "aws-ap-northeast-2",
    "sys.createdTime": "2026-08-21 02:40:28.745 +0000 UTC",
    "sys.cspResourceId": "tbqajlaplk30cjpsahc6",
    "sys.cspResourceName": "tbqajlaplk30cjpsahc6",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig01-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig01-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbqajlaplk30cjpsahc6",
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
        "value": "sg-0467761160db4a2ff"
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
          "cspResourceName": "i-03dc94bf5dbd38c86",
          "cspResourceId": "i-03dc94bf5dbd38c86"
        },
        {
          "cspResourceName": "i-0bf3d6c84a78defb4",
          "cspResourceId": "i-0bf3d6c84a78defb4"
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
          "value": "{NameId:,SystemId:i-03dc94bf5dbd38c86}; {NameId:,SystemId:i-0bf3d6c84a78defb4}"
        },
        {
          "key": "KeyValueList",
          "value": "{Key:AmiType,Value:AL2023_x86_64_STANDARD}; {Key:CapacityType,Value:ON_DEMAND}; {Key:ClusterName,Value:tbqajlaplk30cjpsahc6}; {Key:CreatedAt,Value:2026-08-21T02:47:11.123Z}; {Key:DiskSize,Value:100}; {Key:Health,Value:{Issues:[]}}; {Key:InstanceTypes,Value:c5a.xlarge}; {Key:ModifiedAt,Value:2026-08-21T02:48:53.478Z}; {Key:NodeRole,Value:arn:aws:iam::635484366616:role/cloud-barista-eks-nodegroup-role}; {Key:NodegroupArn,Value:arn:aws:eks:ap-northeast-2:635484366616:nodegroup/tbqajlaplk30cjpsahc6/workers1/8ed0111b-86ff-c035-79d7-ce32357dd95a}; {Key:NodegroupName,Value:workers1}; {Key:ReleaseVersion,Value:1.33.13-20260818}; {Key:RemoteAccess,Value:{Ec2SshKey:tbdn1fr85uqtviuhtgle,SourceSecurityGroups:[sg-00b0fa9c3c57c2cc5]}}; {Key:Resources,Value:{AutoScalingGroups:[{Name:eks-workers1-8ed0111b-86ff-c035-79d7-ce32357dd95a}],RemoteAccessSecurityGroup:sg-08bb5f8b9279adcc9}}; {Key:ScalingConfig,Value:{DesiredSize:2,MaxSize:2,MinSize:2}}; {Key:Status,Value:ACTIVE}; {Key:Subnets,Value:subnet-05247863bd834f2df; subnet-0b020a2249bb6475b}; {Key:Tags,Value:{key:nodegroup,value:workers1}}; {Key:UpdateConfig,Value:{MaxUnavailable:1,MaxUnavailablePercentage:null}}; {Key:Version,Value:1.33}"
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
            "NameId": "i-03dc94bf5dbd38c86",
            "SystemId": "i-03dc94bf5dbd38c86"
          },
          {
            "NameId": "i-0bf3d6c84a78defb4",
            "SystemId": "i-0bf3d6c84a78defb4"
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
            "value": "{NameId:,SystemId:i-03dc94bf5dbd38c86}; {NameId:,SystemId:i-0bf3d6c84a78defb4}"
          },
          {
            "key": "KeyValueList",
            "value": "{Key:AmiType,Value:AL2023_x86_64_STANDARD}; {Key:CapacityType,Value:ON_DEMAND}; {Key:ClusterName,Value:tbqajlaplk30cjpsahc6}; {Key:CreatedAt,Value:2026-08-21T02:47:11.123Z}; {Key:DiskSize,Value:100}; {Key:Health,Value:{Issues:[]}}; {Key:InstanceTypes,Value:c5a.xlarge}; {Key:ModifiedAt,Value:2026-08-21T02:48:53.478Z}; {Key:NodeRole,Value:arn:aws:iam::635484366616:role/cloud-barista-eks-nodegroup-role}; {Key:NodegroupArn,Value:arn:aws:eks:ap-northeast-2:635484366616:nodegroup/tbqajlaplk30cjpsahc6/workers1/8ed0111b-86ff-c035-79d7-ce32357dd95a}; {Key:NodegroupName,Value:workers1}; {Key:ReleaseVersion,Value:1.33.13-20260818}; {Key:RemoteAccess,Value:{Ec2SshKey:tbdn1fr85uqtviuhtgle,SourceSecurityGroups:[sg-00b0fa9c3c57c2cc5]}}; {Key:Resources,Value:{AutoScalingGroups:[{Name:eks-workers1-8ed0111b-86ff-c035-79d7-ce32357dd95a}],RemoteAccessSecurityGroup:sg-08bb5f8b9279adcc9}}; {Key:ScalingConfig,Value:{DesiredSize:2,MaxSize:2,MinSize:2}}; {Key:Status,Value:ACTIVE}; {Key:Subnets,Value:subnet-05247863bd834f2df; subnet-0b020a2249bb6475b}; {Key:Tags,Value:{key:nodegroup,value:workers1}}; {Key:UpdateConfig,Value:{MaxUnavailable:1,MaxUnavailablePercentage:null}}; {Key:Version,Value:1.33}"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://A42C0404A862C06B0E9998C1EC42D7E1.gr7.ap-northeast-2.eks.amazonaws.com",
    "kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://A42C0404A862C06B0E9998C1EC42D7E1.gr7.ap-northeast-2.eks.amazonaws.com\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREVENDQWZXZ0F3SUJBZ0lRYk9HOENhQ05pQkU2UGVpTEFaYUE5VEFOQmdrcWhraUc5dzBCQVFzRkFEQVYKTVJNd0VRWURWUVFERXdwcmRXSmxjbTVsZEdWek1CNFhEVEkyTURneU1UQXlOREF6T1ZvWERUTXhNRGd5TURBeQpOREF6T1Zvd0ZURVRNQkVHQTFVRUF4TUthM1ZpWlhKdVpYUmxjekNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFECmdnRVBBRENDQVFvQ2dnRUJBTWpoc1puYnMrK1NoMEVvd2ROVzdJL2x2RHdIdTNjREIvTkY0NHJMeFJkZWlMTHMKQ2d6MnZkdWY0aHR4cWhJRU1lb0ZkTk9VMkZiR2kyeE9rNWlGdkhYQWtFV1VPTHgxdldxNzU2bUczTk9RRTBTaAorMzlqOGZJOUh1Yk9Qc2lLZlhJN1lxM3o4YWF3TzRFeE1lSFdxNlhIZlBvdnIzR29ZaFFyL2txK3NJUWFpYnQ2CncxL2N1bTJzejJXMFo3bXR1Vjlub2ZTWGxucEtIOTl1Q1ZXalIvamVEWUNMR0VudExkWnQ0THhxc0h5dWtpeEcKRFk2T2tXQ29jMFRRTlpSSWRzeVZ0MUZmRnV4dFdoVDQvRkhhMFlkL09OK2FGek1KZmhtV2ZIWkRaN1NzV1FwMgplR3ZpNHhtS0VCRzd4TUpyL0xtM2w4RFNjUEt4NEJjV2VuUi9JLzBDQXdFQUFhTlpNRmN3RGdZRFZSMFBBUUgvCkJBUURBZ0trTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRkNjaWxHTGFCREo3MjYxM3ZDRloKdUtjQkJpMXNNQlVHQTFVZEVRUU9NQXlDQ210MVltVnlibVYwWlhNd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBSHIzUGZEcFZ0THplWDV1emJKRlFkUkNNa3ZJckx2WHg0UlhTa1lNc3B4Nld0U2RRMGE5Q2cyMlZrR1dtaGdNCnRjMXo2dmxNcHVZUGVhcS8wQ25RQTFhVlRQZWxndFRBakpFaVY3VzZlQWZiZWgvSGxLZW8vVEJnVmdudTRXaVIKTEZOWFl1L1RkUUJjanZ2WFE2MklUMzZkekNGb2pLN0E2U29hUnkydFFld0E2ZWxHekJOaUlEcmcyT0VOQldJagpDZXJoYkxIdDBSN0U2Y2t0Z3dZVHprV082QXdzRTAySjhqWWlRaEVQZWxJWTAzMGhKRHkrbzNQbDZTcC9FSldECjdDa1Z5KzVCTFlBbWszSUMyVjk0M0RJN1RPQUxYRUtIOXFYVHB6Q0Jld1lZOFpidURwM2N2OVc1VlBJdGtGUnAKL3p5VG9FTlZqVTVYSGJOcS9kZlhEMFk9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: tbqajlaplk30cjpsahc6\ncontexts:\n- context:\n    cluster: tbqajlaplk30cjpsahc6\n    user: aws-dynamic-token\n  name: tbqajlaplk30cjpsahc6\ncurrent-context: tbqajlaplk30cjpsahc6\nusers:\n- name: aws-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tbqajlaplk30cjpsahc6/token?ConnectionName=aws-ap-northeast-2\\\"\"\n"
  },
  "addons": {
    "keyValueList": [
      {
        "key": "AddonArn",
        "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tbqajlaplk30cjpsahc6/aws-ebs-csi-driver/32d01118-7bee-835a-b742-59eb902f372c"
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
        "value": "tbqajlaplk30cjpsahc6"
      },
      {
        "key": "CreatedAt",
        "value": "2026-08-21T02:40:29.969Z"
      },
      {
        "key": "Health",
        "value": "{Issues:[]}"
      },
      {
        "key": "ModifiedAt",
        "value": "2026-08-21T02:40:30.04Z"
      },
      {
        "key": "Status",
        "value": "CREATING"
      },
      {
        "key": "AddonArn",
        "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tbqajlaplk30cjpsahc6/eks-pod-identity-agent/68d01118-7d04-d95c-159d-92ab6f64aff9"
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
        "value": "tbqajlaplk30cjpsahc6"
      },
      {
        "key": "CreatedAt",
        "value": "2026-08-21T02:40:30.382Z"
      },
      {
        "key": "Health",
        "value": "{Issues:[]}"
      },
      {
        "key": "ModifiedAt",
        "value": "2026-08-21T02:48:13.224Z"
      },
      {
        "key": "Status",
        "value": "ACTIVE"
      }
    ]
  },
  "status": "Active",
  "createdTime": "2026-08-21T02:40:28.745Z",
  "keyValueList": [
    {
      "key": "Arn",
      "value": "arn:aws:eks:ap-northeast-2:635484366616:cluster/tbqajlaplk30cjpsahc6"
    },
    {
      "key": "CertificateAuthority",
      "value": "{Data:LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREVENDQWZXZ0F3SUJBZ0lRYk9HOENhQ05pQkU2UGVpTEFaYUE5VEFOQmdrcWhraUc5dzBCQVFzRkFEQVYKTVJNd0VRWURWUVFERXdwcmRXSmxjbTVsZEdWek1CNFhEVEkyTURneU1UQXlOREF6T1ZvWERUTXhNRGd5TURBeQpOREF6T1Zvd0ZURVRNQkVHQTFVRUF4TUthM1ZpWlhKdVpYUmxjekNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFECmdnRVBBRENDQVFvQ2dnRUJBTWpoc1puYnMrK1NoMEVvd2ROVzdJL2x2RHdIdTNjREIvTkY0NHJMeFJkZWlMTHMKQ2d6MnZkdWY0aHR4cWhJRU1lb0ZkTk9VMkZiR2kyeE9rNWlGdkhYQWtFV1VPTHgxdldxNzU2bUczTk9RRTBTaAorMzlqOGZJOUh1Yk9Qc2lLZlhJN1lxM3o4YWF3TzRFeE1lSFdxNlhIZlBvdnIzR29ZaFFyL2txK3NJUWFpYnQ2CncxL2N1bTJzejJXMFo3bXR1Vjlub2ZTWGxucEtIOTl1Q1ZXalIvamVEWUNMR0VudExkWnQ0THhxc0h5dWtpeEcKRFk2T2tXQ29jMFRRTlpSSWRzeVZ0MUZmRnV4dFdoVDQvRkhhMFlkL09OK2FGek1KZmhtV2ZIWkRaN1NzV1FwMgplR3ZpNHhtS0VCRzd4TUpyL0xtM2w4RFNjUEt4NEJjV2VuUi9JLzBDQXdFQUFhTlpNRmN3RGdZRFZSMFBBUUgvCkJBUURBZ0trTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRkNjaWxHTGFCREo3MjYxM3ZDRloKdUtjQkJpMXNNQlVHQTFVZEVRUU9NQXlDQ210MVltVnlibVYwWlhNd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBSHIzUGZEcFZ0THplWDV1emJKRlFkUkNNa3ZJckx2WHg0UlhTa1lNc3B4Nld0U2RRMGE5Q2cyMlZrR1dtaGdNCnRjMXo2dmxNcHVZUGVhcS8wQ25RQTFhVlRQZWxndFRBakpFaVY3VzZlQWZiZWgvSGxLZW8vVEJnVmdudTRXaVIKTEZOWFl1L1RkUUJjanZ2WFE2MklUMzZkekNGb2pLN0E2U29hUnkydFFld0E2ZWxHekJOaUlEcmcyT0VOQldJagpDZXJoYkxIdDBSN0U2Y2t0Z3dZVHprV082QXdzRTAySjhqWWlRaEVQZWxJWTAzMGhKRHkrbzNQbDZTcC9FSldECjdDa1Z5KzVCTFlBbWszSUMyVjk0M0RJN1RPQUxYRUtIOXFYVHB6Q0Jld1lZOFpidURwM2N2OVc1VlBJdGtGUnAKL3p5VG9FTlZqVTVYSGJOcS9kZlhEMFk9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K}"
    },
    {
      "key": "CreatedAt",
      "value": "2026-08-21T02:40:28.745Z"
    },
    {
      "key": "Endpoint",
      "value": "https://A42C0404A862C06B0E9998C1EC42D7E1.gr7.ap-northeast-2.eks.amazonaws.com"
    },
    {
      "key": "Identity",
      "value": "{Oidc:{Issuer:https://oidc.eks.ap-northeast-2.amazonaws.com/id/A42C0404A862C06B0E9998C1EC42D7E1}}"
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
      "value": "tbqajlaplk30cjpsahc6"
    },
    {
      "key": "PlatformVersion",
      "value": "eks.45"
    },
    {
      "key": "ResourcesVpcConfig",
      "value": "{ClusterSecurityGroupId:sg-0467761160db4a2ff,EndpointPrivateAccess:false,EndpointPublicAccess:true,PublicAccessCidrs:[0.0.0.0/0],SecurityGroupIds:[sg-00b0fa9c3c57c2cc5],SubnetIds:[subnet-05247863bd834f2df,subnet-0b020a2249bb6475b],VpcId:vpc-01af9f247d39edce6}"
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
      "value": "{Name:tbqajlaplk30cjpsahc6,sys.cspResourceId:tbqajlaplk30cjpsahc6,sys.cspResourceName:tbqajlaplk30cjpsahc6,sys.labelType:k8s,sys.namespace:mig01}"
    },
    {
      "key": "Version",
      "value": "1.33"
    }
  ],
  "cspResourceName": "tbqajlaplk30cjpsahc6",
  "cspResourceId": "tbqajlaplk30cjpsahc6",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbqajlaplk30cjpsahc6",
      "SystemId": "tbqajlaplk30cjpsahc6"
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
          "value": "sg-0467761160db4a2ff"
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
            "NameId": "i-03dc94bf5dbd38c86",
            "SystemId": "i-03dc94bf5dbd38c86"
          },
          {
            "NameId": "i-0bf3d6c84a78defb4",
            "SystemId": "i-0bf3d6c84a78defb4"
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
            "value": "{NameId:,SystemId:i-03dc94bf5dbd38c86}; {NameId:,SystemId:i-0bf3d6c84a78defb4}"
          },
          {
            "key": "KeyValueList",
            "value": "{Key:AmiType,Value:AL2023_x86_64_STANDARD}; {Key:CapacityType,Value:ON_DEMAND}; {Key:ClusterName,Value:tbqajlaplk30cjpsahc6}; {Key:CreatedAt,Value:2026-08-21T02:47:11.123Z}; {Key:DiskSize,Value:100}; {Key:Health,Value:{Issues:[]}}; {Key:InstanceTypes,Value:c5a.xlarge}; {Key:ModifiedAt,Value:2026-08-21T02:48:53.478Z}; {Key:NodeRole,Value:arn:aws:iam::635484366616:role/cloud-barista-eks-nodegroup-role}; {Key:NodegroupArn,Value:arn:aws:eks:ap-northeast-2:635484366616:nodegroup/tbqajlaplk30cjpsahc6/workers1/8ed0111b-86ff-c035-79d7-ce32357dd95a}; {Key:NodegroupName,Value:workers1}; {Key:ReleaseVersion,Value:1.33.13-20260818}; {Key:RemoteAccess,Value:{Ec2SshKey:tbdn1fr85uqtviuhtgle,SourceSecurityGroups:[sg-00b0fa9c3c57c2cc5]}}; {Key:Resources,Value:{AutoScalingGroups:[{Name:eks-workers1-8ed0111b-86ff-c035-79d7-ce32357dd95a}],RemoteAccessSecurityGroup:sg-08bb5f8b9279adcc9}}; {Key:ScalingConfig,Value:{DesiredSize:2,MaxSize:2,MinSize:2}}; {Key:Status,Value:ACTIVE}; {Key:Subnets,Value:subnet-05247863bd834f2df; subnet-0b020a2249bb6475b}; {Key:Tags,Value:{key:nodegroup,value:workers1}}; {Key:UpdateConfig,Value:{MaxUnavailable:1,MaxUnavailablePercentage:null}}; {Key:Version,Value:1.33}"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://A42C0404A862C06B0E9998C1EC42D7E1.gr7.ap-northeast-2.eks.amazonaws.com",
      "Kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://A42C0404A862C06B0E9998C1EC42D7E1.gr7.ap-northeast-2.eks.amazonaws.com\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREVENDQWZXZ0F3SUJBZ0lRYk9HOENhQ05pQkU2UGVpTEFaYUE5VEFOQmdrcWhraUc5dzBCQVFzRkFEQVYKTVJNd0VRWURWUVFERXdwcmRXSmxjbTVsZEdWek1CNFhEVEkyTURneU1UQXlOREF6T1ZvWERUTXhNRGd5TURBeQpOREF6T1Zvd0ZURVRNQkVHQTFVRUF4TUthM1ZpWlhKdVpYUmxjekNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFECmdnRVBBRENDQVFvQ2dnRUJBTWpoc1puYnMrK1NoMEVvd2ROVzdJL2x2RHdIdTNjREIvTkY0NHJMeFJkZWlMTHMKQ2d6MnZkdWY0aHR4cWhJRU1lb0ZkTk9VMkZiR2kyeE9rNWlGdkhYQWtFV1VPTHgxdldxNzU2bUczTk9RRTBTaAorMzlqOGZJOUh1Yk9Qc2lLZlhJN1lxM3o4YWF3TzRFeE1lSFdxNlhIZlBvdnIzR29ZaFFyL2txK3NJUWFpYnQ2CncxL2N1bTJzejJXMFo3bXR1Vjlub2ZTWGxucEtIOTl1Q1ZXalIvamVEWUNMR0VudExkWnQ0THhxc0h5dWtpeEcKRFk2T2tXQ29jMFRRTlpSSWRzeVZ0MUZmRnV4dFdoVDQvRkhhMFlkL09OK2FGek1KZmhtV2ZIWkRaN1NzV1FwMgplR3ZpNHhtS0VCRzd4TUpyL0xtM2w4RFNjUEt4NEJjV2VuUi9JLzBDQXdFQUFhTlpNRmN3RGdZRFZSMFBBUUgvCkJBUURBZ0trTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRkNjaWxHTGFCREo3MjYxM3ZDRloKdUtjQkJpMXNNQlVHQTFVZEVRUU9NQXlDQ210MVltVnlibVYwWlhNd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBSHIzUGZEcFZ0THplWDV1emJKRlFkUkNNa3ZJckx2WHg0UlhTa1lNc3B4Nld0U2RRMGE5Q2cyMlZrR1dtaGdNCnRjMXo2dmxNcHVZUGVhcS8wQ25RQTFhVlRQZWxndFRBakpFaVY3VzZlQWZiZWgvSGxLZW8vVEJnVmdudTRXaVIKTEZOWFl1L1RkUUJjanZ2WFE2MklUMzZkekNGb2pLN0E2U29hUnkydFFld0E2ZWxHekJOaUlEcmcyT0VOQldJagpDZXJoYkxIdDBSN0U2Y2t0Z3dZVHprV082QXdzRTAySjhqWWlRaEVQZWxJWTAzMGhKRHkrbzNQbDZTcC9FSldECjdDa1Z5KzVCTFlBbWszSUMyVjk0M0RJN1RPQUxYRUtIOXFYVHB6Q0Jld1lZOFpidURwM2N2OVc1VlBJdGtGUnAKL3p5VG9FTlZqVTVYSGJOcS9kZlhEMFk9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: tbqajlaplk30cjpsahc6\ncontexts:\n- context:\n    cluster: tbqajlaplk30cjpsahc6\n    user: aws-dynamic-token\n  name: tbqajlaplk30cjpsahc6\ncurrent-context: tbqajlaplk30cjpsahc6\nusers:\n- name: aws-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tbqajlaplk30cjpsahc6/token?ConnectionName=aws-ap-northeast-2\\\"\"\n"
    },
    "Addons": {
      "KeyValueList": [
        {
          "key": "AddonArn",
          "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tbqajlaplk30cjpsahc6/aws-ebs-csi-driver/32d01118-7bee-835a-b742-59eb902f372c"
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
          "value": "tbqajlaplk30cjpsahc6"
        },
        {
          "key": "CreatedAt",
          "value": "2026-08-21T02:40:29.969Z"
        },
        {
          "key": "Health",
          "value": "{Issues:[]}"
        },
        {
          "key": "ModifiedAt",
          "value": "2026-08-21T02:40:30.04Z"
        },
        {
          "key": "Status",
          "value": "CREATING"
        },
        {
          "key": "AddonArn",
          "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tbqajlaplk30cjpsahc6/eks-pod-identity-agent/68d01118-7d04-d95c-159d-92ab6f64aff9"
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
          "value": "tbqajlaplk30cjpsahc6"
        },
        {
          "key": "CreatedAt",
          "value": "2026-08-21T02:40:30.382Z"
        },
        {
          "key": "Health",
          "value": "{Issues:[]}"
        },
        {
          "key": "ModifiedAt",
          "value": "2026-08-21T02:48:13.224Z"
        },
        {
          "key": "Status",
          "value": "ACTIVE"
        }
      ]
    },
    "Status": "Active",
    "CreatedTime": "2026-08-21T02:40:28.745Z",
    "KeyValueList": [
      {
        "key": "Arn",
        "value": "arn:aws:eks:ap-northeast-2:635484366616:cluster/tbqajlaplk30cjpsahc6"
      },
      {
        "key": "CertificateAuthority",
        "value": "{Data:LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREVENDQWZXZ0F3SUJBZ0lRYk9HOENhQ05pQkU2UGVpTEFaYUE5VEFOQmdrcWhraUc5dzBCQVFzRkFEQVYKTVJNd0VRWURWUVFERXdwcmRXSmxjbTVsZEdWek1CNFhEVEkyTURneU1UQXlOREF6T1ZvWERUTXhNRGd5TURBeQpOREF6T1Zvd0ZURVRNQkVHQTFVRUF4TUthM1ZpWlhKdVpYUmxjekNDQVNJd0RRWUpLb1pJaHZjTkFRRUJCUUFECmdnRVBBRENDQVFvQ2dnRUJBTWpoc1puYnMrK1NoMEVvd2ROVzdJL2x2RHdIdTNjREIvTkY0NHJMeFJkZWlMTHMKQ2d6MnZkdWY0aHR4cWhJRU1lb0ZkTk9VMkZiR2kyeE9rNWlGdkhYQWtFV1VPTHgxdldxNzU2bUczTk9RRTBTaAorMzlqOGZJOUh1Yk9Qc2lLZlhJN1lxM3o4YWF3TzRFeE1lSFdxNlhIZlBvdnIzR29ZaFFyL2txK3NJUWFpYnQ2CncxL2N1bTJzejJXMFo3bXR1Vjlub2ZTWGxucEtIOTl1Q1ZXalIvamVEWUNMR0VudExkWnQ0THhxc0h5dWtpeEcKRFk2T2tXQ29jMFRRTlpSSWRzeVZ0MUZmRnV4dFdoVDQvRkhhMFlkL09OK2FGek1KZmhtV2ZIWkRaN1NzV1FwMgplR3ZpNHhtS0VCRzd4TUpyL0xtM2w4RFNjUEt4NEJjV2VuUi9JLzBDQXdFQUFhTlpNRmN3RGdZRFZSMFBBUUgvCkJBUURBZ0trTUE4R0ExVWRFd0VCL3dRRk1BTUJBZjh3SFFZRFZSME9CQllFRkNjaWxHTGFCREo3MjYxM3ZDRloKdUtjQkJpMXNNQlVHQTFVZEVRUU9NQXlDQ210MVltVnlibVYwWlhNd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBSHIzUGZEcFZ0THplWDV1emJKRlFkUkNNa3ZJckx2WHg0UlhTa1lNc3B4Nld0U2RRMGE5Q2cyMlZrR1dtaGdNCnRjMXo2dmxNcHVZUGVhcS8wQ25RQTFhVlRQZWxndFRBakpFaVY3VzZlQWZiZWgvSGxLZW8vVEJnVmdudTRXaVIKTEZOWFl1L1RkUUJjanZ2WFE2MklUMzZkekNGb2pLN0E2U29hUnkydFFld0E2ZWxHekJOaUlEcmcyT0VOQldJagpDZXJoYkxIdDBSN0U2Y2t0Z3dZVHprV082QXdzRTAySjhqWWlRaEVQZWxJWTAzMGhKRHkrbzNQbDZTcC9FSldECjdDa1Z5KzVCTFlBbWszSUMyVjk0M0RJN1RPQUxYRUtIOXFYVHB6Q0Jld1lZOFpidURwM2N2OVc1VlBJdGtGUnAKL3p5VG9FTlZqVTVYSGJOcS9kZlhEMFk9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K}"
      },
      {
        "key": "CreatedAt",
        "value": "2026-08-21T02:40:28.745Z"
      },
      {
        "key": "Endpoint",
        "value": "https://A42C0404A862C06B0E9998C1EC42D7E1.gr7.ap-northeast-2.eks.amazonaws.com"
      },
      {
        "key": "Identity",
        "value": "{Oidc:{Issuer:https://oidc.eks.ap-northeast-2.amazonaws.com/id/A42C0404A862C06B0E9998C1EC42D7E1}}"
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
        "value": "tbqajlaplk30cjpsahc6"
      },
      {
        "key": "PlatformVersion",
        "value": "eks.45"
      },
      {
        "key": "ResourcesVpcConfig",
        "value": "{ClusterSecurityGroupId:sg-0467761160db4a2ff,EndpointPrivateAccess:false,EndpointPublicAccess:true,PublicAccessCidrs:[0.0.0.0/0],SecurityGroupIds:[sg-00b0fa9c3c57c2cc5],SubnetIds:[subnet-05247863bd834f2df,subnet-0b020a2249bb6475b],VpcId:vpc-01af9f247d39edce6}"
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
        "value": "{Name:tbqajlaplk30cjpsahc6,sys.cspResourceId:tbqajlaplk30cjpsahc6,sys.cspResourceName:tbqajlaplk30cjpsahc6,sys.labelType:k8s,sys.namespace:mig01}"
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

