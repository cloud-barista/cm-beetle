# CM-Beetle K8s Infra Migration Test Results — AWS-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → validate → migrate → list → get → workload → delete → residual.

## Environment

- CSP / Region: aws / ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (cd1f3a2)
- Git Commit: cd1f3a2
- Namespace: mig01
- Test Date: 2026-10-07 16:41:05 KST
- Cluster ID: mig01-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 52ms |
| 2 | POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation) | ✅ **PASS** | 335ms |
| 3 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 12m15.102s |
| 4 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 3ms |
| 5 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 1.012s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 1m23.054s |
| 7 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 8m58.608s |
| 8 | Residual resource check (Tumblebug) | ✅ **PASS** | 5ms |

**Overall Result**: 8/8 steps passed ✅

**Total Duration**: 22m38s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 52ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.34)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=aws+ap-northeast-2+c5a.xlarge nodes=2

### Step 2 — POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)

- **Duration**: 335ms
- **Status Code**: 200

- ✅ Target K8s infra model is valid for migration (0 issues)

### Step 3 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 12m15.102s
- **Status Code**: 202

- ℹ️  nameSeed: mig01
- ℹ️  async reqId: 1791358865988222999
- ℹ️  cluster id: mig01-on-prem-k8s-cluster
- ℹ️  elapsed: 12m15s
- ✅ status: Active

### Step 4 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 3ms
- **Status Code**: 200

- ✅ migrated cluster present in list (6 total)

### Step 5 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 1.012s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=aws+ap-northeast-2+c5a.xlarge, nodes=2)
- ✅ version: 1.34 (recommended 1.34)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 1m23.054s

- ✅ kubeconfig obtained (server: https://2E128CC0E45344BF3744FDA996EEB348.gr7.ap-northeast-2.eks.amazonaws.com)
- ℹ️  auth method: exec credential plugin
- ✅ cluster token obtained from Tumblebug
- ✅ API server reachable (v1.34.11-eks-cfb47f5)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 2)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: ab3c8e23e840644338351da7c306ba7c-2028199555.ap-northeast-2.elb.amazonaws.com
- ✅ nginx served over the LoadBalancer at http://ab3c8e23e840644338351da7c306ba7c-2028199555.ap-northeast-2.elb.amazonaws.com/ (attempt 7)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 7 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 8m58.608s
- **Status Code**: 200

- ✅ deleted on attempt 1 (8m58s)

### Step 8 — Residual resource check (Tumblebug)

- **Duration**: 5ms

- ℹ️  VNet mig01-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig01-k8s-sg still exists (known gap)
- ℹ️  SshKey mig01-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for aws ap-northeast-2 (source: v1.32.3 → target: v1.34)",
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
    "connectionName": "aws-ap-northeast-2",
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
        "specId": "aws+ap-northeast-2+c5a.xlarge",
        "rootDiskType": "default",
        "rootDiskSize": 100,
        "sshKeyId": "",
        "onAutoScaling": "true",
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
  "id": "mig01-on-prem-k8s-cluster",
  "uid": "tb56lfotn8retqg9nve4",
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
    "sys.connectionName": "aws-ap-northeast-2",
    "sys.createdTime": "2026-10-07 07:41:09.979 +0000 UTC",
    "sys.cspResourceId": "tb56lfotn8retqg9nve4",
    "sys.cspResourceName": "tb56lfotn8retqg9nve4",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig01-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig01-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tb56lfotn8retqg9nve4",
    "sys.version": "1.34"
  },
  "systemLabel": "",
  "version": "1.34",
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
        "value": "sg-0150e18b69f6edc2f"
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
      "onAutoScaling": true,
      "desiredNodeSize": 2,
      "minNodeSize": 2,
      "maxNodeSize": 2,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "i-00b909ff903e1e18a",
          "cspResourceId": "i-00b909ff903e1e18a"
        },
        {
          "cspResourceName": "i-0d2b1715eb6cb2b15",
          "cspResourceId": "i-0d2b1715eb6cb2b15"
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
          "value": "true"
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
          "value": "{NameId:,SystemId:i-00b909ff903e1e18a}; {NameId:,SystemId:i-0d2b1715eb6cb2b15}"
        },
        {
          "key": "KeyValueList",
          "value": "{Key:AmiType,Value:AL2023_x86_64_STANDARD}; {Key:CapacityType,Value:ON_DEMAND}; {Key:ClusterName,Value:tb56lfotn8retqg9nve4}; {Key:CreatedAt,Value:2026-10-07T07:51:07.089Z}; {Key:DiskSize,Value:100}; {Key:Health,Value:{Issues:[]}}; {Key:InstanceTypes,Value:c5a.xlarge}; {Key:ModifiedAt,Value:2026-10-07T07:53:06.971Z}; {Key:NodeRole,Value:arn:aws:iam::635484366616:role/cloud-barista-eks-nodegroup-role}; {Key:NodegroupArn,Value:arn:aws:eks:ap-northeast-2:635484366616:nodegroup/tb56lfotn8retqg9nve4/workers1/bed08aac-1a4c-835b-6b12-589af7129a52}; {Key:NodegroupName,Value:workers1}; {Key:ReleaseVersion,Value:1.34.11-20260930}; {Key:RemoteAccess,Value:{Ec2SshKey:tbdn1fr85uqtviuhtgle,SourceSecurityGroups:[sg-00b0fa9c3c57c2cc5]}}; {Key:Resources,Value:{AutoScalingGroups:[{Name:eks-workers1-bed08aac-1a4c-835b-6b12-589af7129a52}],RemoteAccessSecurityGroup:sg-0e427ac95f147e0cf}}; {Key:ScalingConfig,Value:{DesiredSize:2,MaxSize:2,MinSize:2}}; {Key:Status,Value:ACTIVE}; {Key:Subnets,Value:subnet-05247863bd834f2df; subnet-0b020a2249bb6475b}; {Key:Tags,Value:{key:nodegroup,value:workers1}}; {Key:UpdateConfig,Value:{MaxUnavailable:1,MaxUnavailablePercentage:null}}; {Key:Version,Value:1.34}"
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
        "OnAutoScaling": true,
        "DesiredNodeSize": 2,
        "MinNodeSize": 2,
        "MaxNodeSize": 2,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "i-00b909ff903e1e18a",
            "SystemId": "i-00b909ff903e1e18a"
          },
          {
            "NameId": "i-0d2b1715eb6cb2b15",
            "SystemId": "i-0d2b1715eb6cb2b15"
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
            "value": "true"
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
            "value": "{NameId:,SystemId:i-00b909ff903e1e18a}; {NameId:,SystemId:i-0d2b1715eb6cb2b15}"
          },
          {
            "key": "KeyValueList",
            "value": "{Key:AmiType,Value:AL2023_x86_64_STANDARD}; {Key:CapacityType,Value:ON_DEMAND}; {Key:ClusterName,Value:tb56lfotn8retqg9nve4}; {Key:CreatedAt,Value:2026-10-07T07:51:07.089Z}; {Key:DiskSize,Value:100}; {Key:Health,Value:{Issues:[]}}; {Key:InstanceTypes,Value:c5a.xlarge}; {Key:ModifiedAt,Value:2026-10-07T07:53:06.971Z}; {Key:NodeRole,Value:arn:aws:iam::635484366616:role/cloud-barista-eks-nodegroup-role}; {Key:NodegroupArn,Value:arn:aws:eks:ap-northeast-2:635484366616:nodegroup/tb56lfotn8retqg9nve4/workers1/bed08aac-1a4c-835b-6b12-589af7129a52}; {Key:NodegroupName,Value:workers1}; {Key:ReleaseVersion,Value:1.34.11-20260930}; {Key:RemoteAccess,Value:{Ec2SshKey:tbdn1fr85uqtviuhtgle,SourceSecurityGroups:[sg-00b0fa9c3c57c2cc5]}}; {Key:Resources,Value:{AutoScalingGroups:[{Name:eks-workers1-bed08aac-1a4c-835b-6b12-589af7129a52}],RemoteAccessSecurityGroup:sg-0e427ac95f147e0cf}}; {Key:ScalingConfig,Value:{DesiredSize:2,MaxSize:2,MinSize:2}}; {Key:Status,Value:ACTIVE}; {Key:Subnets,Value:subnet-05247863bd834f2df; subnet-0b020a2249bb6475b}; {Key:Tags,Value:{key:nodegroup,value:workers1}}; {Key:UpdateConfig,Value:{MaxUnavailable:1,MaxUnavailablePercentage:null}}; {Key:Version,Value:1.34}"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://2E128CC0E45344BF3744FDA996EEB348.gr7.ap-northeast-2.eks.amazonaws.com",
    "kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://2E128CC0E45344BF3744FDA996EEB348.gr7.ap-northeast-2.eks.amazonaws.com\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURNakNDQWhxZ0F3SUJBZ0lSQUoxMU5GOUhhYURCdndqVmdHNGlGb3d3RFFZSktvWklodmNOQVFFTEJRQXcKSnpFUU1BNEdBMVVFQ2hNSFFWZFRJRVZMVXpFVE1CRUdBMVVFQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TmpFdwpNRGN3TnpReE1UWmFGdzB6TVRFd01EWXdOelF4TVRaYU1DY3hFREFPQmdOVkJBb1RCMEZYVXlCRlMxTXhFekFSCkJnTlZCQU1UQ210MVltVnlibVYwWlhNd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUIKQVFEVGJDc2QwRFBDdC90QVUrM1V5d3NxRHlHbDlHc0JNU2JyeXZLUi9HcGZ3Nmh5UVdjckZzQS9hTkZobWxVawo4dTVQcEpHeHNub3lSOFEyYUt6VnE3V0YyNUI4NUhmUXpsTmFtNCtXNGFrYXozdUdpSHUzZTI2aUdBRVE0eWVlCkhVZ09UUmpFRTc2V0crcUIzcjJrK3V1Uk8ydzhTMDMyMStpMkZvRUROcm4vanNRdVVOSW9XZUVtRDdHUkhXVFIKVzBuNkc1TU1hVzQyUWJEQ3l3Z2VLcUlmQStUTVZRUldlV28vL2JaNW0zZVZ6R3NuQkRLbFB5Q0Uydm5ETTZ5NQp2dVB2SW9iK3JWQWYvMGROZWsxRUh3ZUtYYzdTUm84N1M5TkRXcTZmUU9JcmxtNzJuU3l1K2doZVhKTWhWcFJQCmJEYi8yWHBXTS9ZQzVYZUlPd25uYTZnZkFnTUJBQUdqV1RCWE1BNEdBMVVkRHdFQi93UUVBd0lDcERBUEJnTlYKSFJNQkFmOEVCVEFEQVFIL01CMEdBMVVkRGdRV0JCUm54YUdXU2ZFdDZvTFdnUE9ENVh2OW5YZjJpREFWQmdOVgpIUkVFRGpBTWdncHJkV0psY201bGRHVnpNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUIydytvaE9KajlNSmlqCkdiSWhXYjZLZ2MrODZZdmR2MFk1b2dlekNrdHFicmJITTRzQjdxWmhhd2MxUURTdmZFVG9HS21jblNXeVlxaDQKV0VEZFRhODNxU1dTV3E1TnFQYndWSDJSd1hWWUF0cFBpVi9HSnF6VWhaSlMzTit0aFhMSGYxOXYwaVo4VHhJeQpIOG5uWUo3UEFLakh6c1VSVVNQY2l0YlZYOWxadFozUk9mMUdTQ25yQXVWTFVNTHU3MGNrUG1QYXhvclRkeFVICi90Q0lrR1pwckl5QWpqYXAzQ2gvMlBOZ2dRVzFnemIyTTNWcjVTbHZwdy9iVkNsc0pPS2t0aHJkbHcycEc4UDYKdk5DODJmRmI2dEVoY2FBZ1VTemlTWVRVN1IrUHZ3aTdHdElhYm81NCtxTURGSGhRTGJjMW90a2o4OElmdFRsRQpnYmVqd0lTOQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n  name: tb56lfotn8retqg9nve4\ncontexts:\n- context:\n    cluster: tb56lfotn8retqg9nve4\n    user: aws-dynamic-token\n  name: tb56lfotn8retqg9nve4\ncurrent-context: tb56lfotn8retqg9nve4\nusers:\n- name: aws-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tb56lfotn8retqg9nve4/token?ConnectionName=aws-ap-northeast-2\\\"\"\n"
  },
  "addons": {
    "keyValueList": [
      {
        "key": "AddonArn",
        "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tb56lfotn8retqg9nve4/aws-ebs-csi-driver/32d08aa7-9317-e79d-865a-c768c8525784"
      },
      {
        "key": "AddonName",
        "value": "aws-ebs-csi-driver"
      },
      {
        "key": "AddonVersion",
        "value": "v1.66.0-eksbuild.1"
      },
      {
        "key": "ClusterName",
        "value": "tb56lfotn8retqg9nve4"
      },
      {
        "key": "CreatedAt",
        "value": "2026-10-07T07:41:11.253Z"
      },
      {
        "key": "Health",
        "value": "{Issues:[]}"
      },
      {
        "key": "ModifiedAt",
        "value": "2026-10-07T07:41:11.264Z"
      },
      {
        "key": "Status",
        "value": "CREATING"
      },
      {
        "key": "AddonArn",
        "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tb56lfotn8retqg9nve4/eks-pod-identity-agent/6ad08aa7-93b1-d073-af49-b4207e0bef9f"
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
        "value": "tb56lfotn8retqg9nve4"
      },
      {
        "key": "CreatedAt",
        "value": "2026-10-07T07:41:11.581Z"
      },
      {
        "key": "Health",
        "value": "{Issues:[]}"
      },
      {
        "key": "ModifiedAt",
        "value": "2026-10-07T07:41:11.594Z"
      },
      {
        "key": "Status",
        "value": "CREATING"
      }
    ]
  },
  "status": "Active",
  "createdTime": "2026-10-07T07:41:09.979Z",
  "keyValueList": [
    {
      "key": "Arn",
      "value": "arn:aws:eks:ap-northeast-2:635484366616:cluster/tb56lfotn8retqg9nve4"
    },
    {
      "key": "CertificateAuthority",
      "value": "{Data:LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURNakNDQWhxZ0F3SUJBZ0lSQUoxMU5GOUhhYURCdndqVmdHNGlGb3d3RFFZSktvWklodmNOQVFFTEJRQXcKSnpFUU1BNEdBMVVFQ2hNSFFWZFRJRVZMVXpFVE1CRUdBMVVFQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TmpFdwpNRGN3TnpReE1UWmFGdzB6TVRFd01EWXdOelF4TVRaYU1DY3hFREFPQmdOVkJBb1RCMEZYVXlCRlMxTXhFekFSCkJnTlZCQU1UQ210MVltVnlibVYwWlhNd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUIKQVFEVGJDc2QwRFBDdC90QVUrM1V5d3NxRHlHbDlHc0JNU2JyeXZLUi9HcGZ3Nmh5UVdjckZzQS9hTkZobWxVawo4dTVQcEpHeHNub3lSOFEyYUt6VnE3V0YyNUI4NUhmUXpsTmFtNCtXNGFrYXozdUdpSHUzZTI2aUdBRVE0eWVlCkhVZ09UUmpFRTc2V0crcUIzcjJrK3V1Uk8ydzhTMDMyMStpMkZvRUROcm4vanNRdVVOSW9XZUVtRDdHUkhXVFIKVzBuNkc1TU1hVzQyUWJEQ3l3Z2VLcUlmQStUTVZRUldlV28vL2JaNW0zZVZ6R3NuQkRLbFB5Q0Uydm5ETTZ5NQp2dVB2SW9iK3JWQWYvMGROZWsxRUh3ZUtYYzdTUm84N1M5TkRXcTZmUU9JcmxtNzJuU3l1K2doZVhKTWhWcFJQCmJEYi8yWHBXTS9ZQzVYZUlPd25uYTZnZkFnTUJBQUdqV1RCWE1BNEdBMVVkRHdFQi93UUVBd0lDcERBUEJnTlYKSFJNQkFmOEVCVEFEQVFIL01CMEdBMVVkRGdRV0JCUm54YUdXU2ZFdDZvTFdnUE9ENVh2OW5YZjJpREFWQmdOVgpIUkVFRGpBTWdncHJkV0psY201bGRHVnpNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUIydytvaE9KajlNSmlqCkdiSWhXYjZLZ2MrODZZdmR2MFk1b2dlekNrdHFicmJITTRzQjdxWmhhd2MxUURTdmZFVG9HS21jblNXeVlxaDQKV0VEZFRhODNxU1dTV3E1TnFQYndWSDJSd1hWWUF0cFBpVi9HSnF6VWhaSlMzTit0aFhMSGYxOXYwaVo4VHhJeQpIOG5uWUo3UEFLakh6c1VSVVNQY2l0YlZYOWxadFozUk9mMUdTQ25yQXVWTFVNTHU3MGNrUG1QYXhvclRkeFVICi90Q0lrR1pwckl5QWpqYXAzQ2gvMlBOZ2dRVzFnemIyTTNWcjVTbHZwdy9iVkNsc0pPS2t0aHJkbHcycEc4UDYKdk5DODJmRmI2dEVoY2FBZ1VTemlTWVRVN1IrUHZ3aTdHdElhYm81NCtxTURGSGhRTGJjMW90a2o4OElmdFRsRQpnYmVqd0lTOQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==}"
    },
    {
      "key": "CreatedAt",
      "value": "2026-10-07T07:41:09.979Z"
    },
    {
      "key": "Endpoint",
      "value": "https://2E128CC0E45344BF3744FDA996EEB348.gr7.ap-northeast-2.eks.amazonaws.com"
    },
    {
      "key": "Identity",
      "value": "{Oidc:{Issuer:https://oidc.eks.ap-northeast-2.amazonaws.com/id/2E128CC0E45344BF3744FDA996EEB348}}"
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
      "value": "tb56lfotn8retqg9nve4"
    },
    {
      "key": "PlatformVersion",
      "value": "eks.34"
    },
    {
      "key": "ResourcesVpcConfig",
      "value": "{ClusterSecurityGroupId:sg-0150e18b69f6edc2f,EndpointPrivateAccess:false,EndpointPublicAccess:true,PublicAccessCidrs:[0.0.0.0/0],SecurityGroupIds:[sg-00b0fa9c3c57c2cc5],SubnetIds:[subnet-05247863bd834f2df,subnet-0b020a2249bb6475b],VpcId:vpc-01af9f247d39edce6}"
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
      "value": "{Name:tb56lfotn8retqg9nve4}"
    },
    {
      "key": "Version",
      "value": "1.34"
    }
  ],
  "cspResourceName": "tb56lfotn8retqg9nve4",
  "cspResourceId": "tb56lfotn8retqg9nve4",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tb56lfotn8retqg9nve4",
      "SystemId": "tb56lfotn8retqg9nve4"
    },
    "Version": "1.34",
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
          "value": "sg-0150e18b69f6edc2f"
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
        "OnAutoScaling": true,
        "DesiredNodeSize": 2,
        "MinNodeSize": 2,
        "MaxNodeSize": 2,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "i-00b909ff903e1e18a",
            "SystemId": "i-00b909ff903e1e18a"
          },
          {
            "NameId": "i-0d2b1715eb6cb2b15",
            "SystemId": "i-0d2b1715eb6cb2b15"
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
            "value": "true"
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
            "value": "{NameId:,SystemId:i-00b909ff903e1e18a}; {NameId:,SystemId:i-0d2b1715eb6cb2b15}"
          },
          {
            "key": "KeyValueList",
            "value": "{Key:AmiType,Value:AL2023_x86_64_STANDARD}; {Key:CapacityType,Value:ON_DEMAND}; {Key:ClusterName,Value:tb56lfotn8retqg9nve4}; {Key:CreatedAt,Value:2026-10-07T07:51:07.089Z}; {Key:DiskSize,Value:100}; {Key:Health,Value:{Issues:[]}}; {Key:InstanceTypes,Value:c5a.xlarge}; {Key:ModifiedAt,Value:2026-10-07T07:53:06.971Z}; {Key:NodeRole,Value:arn:aws:iam::635484366616:role/cloud-barista-eks-nodegroup-role}; {Key:NodegroupArn,Value:arn:aws:eks:ap-northeast-2:635484366616:nodegroup/tb56lfotn8retqg9nve4/workers1/bed08aac-1a4c-835b-6b12-589af7129a52}; {Key:NodegroupName,Value:workers1}; {Key:ReleaseVersion,Value:1.34.11-20260930}; {Key:RemoteAccess,Value:{Ec2SshKey:tbdn1fr85uqtviuhtgle,SourceSecurityGroups:[sg-00b0fa9c3c57c2cc5]}}; {Key:Resources,Value:{AutoScalingGroups:[{Name:eks-workers1-bed08aac-1a4c-835b-6b12-589af7129a52}],RemoteAccessSecurityGroup:sg-0e427ac95f147e0cf}}; {Key:ScalingConfig,Value:{DesiredSize:2,MaxSize:2,MinSize:2}}; {Key:Status,Value:ACTIVE}; {Key:Subnets,Value:subnet-05247863bd834f2df; subnet-0b020a2249bb6475b}; {Key:Tags,Value:{key:nodegroup,value:workers1}}; {Key:UpdateConfig,Value:{MaxUnavailable:1,MaxUnavailablePercentage:null}}; {Key:Version,Value:1.34}"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://2E128CC0E45344BF3744FDA996EEB348.gr7.ap-northeast-2.eks.amazonaws.com",
      "Kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://2E128CC0E45344BF3744FDA996EEB348.gr7.ap-northeast-2.eks.amazonaws.com\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURNakNDQWhxZ0F3SUJBZ0lSQUoxMU5GOUhhYURCdndqVmdHNGlGb3d3RFFZSktvWklodmNOQVFFTEJRQXcKSnpFUU1BNEdBMVVFQ2hNSFFWZFRJRVZMVXpFVE1CRUdBMVVFQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TmpFdwpNRGN3TnpReE1UWmFGdzB6TVRFd01EWXdOelF4TVRaYU1DY3hFREFPQmdOVkJBb1RCMEZYVXlCRlMxTXhFekFSCkJnTlZCQU1UQ210MVltVnlibVYwWlhNd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUIKQVFEVGJDc2QwRFBDdC90QVUrM1V5d3NxRHlHbDlHc0JNU2JyeXZLUi9HcGZ3Nmh5UVdjckZzQS9hTkZobWxVawo4dTVQcEpHeHNub3lSOFEyYUt6VnE3V0YyNUI4NUhmUXpsTmFtNCtXNGFrYXozdUdpSHUzZTI2aUdBRVE0eWVlCkhVZ09UUmpFRTc2V0crcUIzcjJrK3V1Uk8ydzhTMDMyMStpMkZvRUROcm4vanNRdVVOSW9XZUVtRDdHUkhXVFIKVzBuNkc1TU1hVzQyUWJEQ3l3Z2VLcUlmQStUTVZRUldlV28vL2JaNW0zZVZ6R3NuQkRLbFB5Q0Uydm5ETTZ5NQp2dVB2SW9iK3JWQWYvMGROZWsxRUh3ZUtYYzdTUm84N1M5TkRXcTZmUU9JcmxtNzJuU3l1K2doZVhKTWhWcFJQCmJEYi8yWHBXTS9ZQzVYZUlPd25uYTZnZkFnTUJBQUdqV1RCWE1BNEdBMVVkRHdFQi93UUVBd0lDcERBUEJnTlYKSFJNQkFmOEVCVEFEQVFIL01CMEdBMVVkRGdRV0JCUm54YUdXU2ZFdDZvTFdnUE9ENVh2OW5YZjJpREFWQmdOVgpIUkVFRGpBTWdncHJkV0psY201bGRHVnpNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUIydytvaE9KajlNSmlqCkdiSWhXYjZLZ2MrODZZdmR2MFk1b2dlekNrdHFicmJITTRzQjdxWmhhd2MxUURTdmZFVG9HS21jblNXeVlxaDQKV0VEZFRhODNxU1dTV3E1TnFQYndWSDJSd1hWWUF0cFBpVi9HSnF6VWhaSlMzTit0aFhMSGYxOXYwaVo4VHhJeQpIOG5uWUo3UEFLakh6c1VSVVNQY2l0YlZYOWxadFozUk9mMUdTQ25yQXVWTFVNTHU3MGNrUG1QYXhvclRkeFVICi90Q0lrR1pwckl5QWpqYXAzQ2gvMlBOZ2dRVzFnemIyTTNWcjVTbHZwdy9iVkNsc0pPS2t0aHJkbHcycEc4UDYKdk5DODJmRmI2dEVoY2FBZ1VTemlTWVRVN1IrUHZ3aTdHdElhYm81NCtxTURGSGhRTGJjMW90a2o4OElmdFRsRQpnYmVqd0lTOQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n  name: tb56lfotn8retqg9nve4\ncontexts:\n- context:\n    cluster: tb56lfotn8retqg9nve4\n    user: aws-dynamic-token\n  name: tb56lfotn8retqg9nve4\ncurrent-context: tb56lfotn8retqg9nve4\nusers:\n- name: aws-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tb56lfotn8retqg9nve4/token?ConnectionName=aws-ap-northeast-2\\\"\"\n"
    },
    "Addons": {
      "KeyValueList": [
        {
          "key": "AddonArn",
          "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tb56lfotn8retqg9nve4/aws-ebs-csi-driver/32d08aa7-9317-e79d-865a-c768c8525784"
        },
        {
          "key": "AddonName",
          "value": "aws-ebs-csi-driver"
        },
        {
          "key": "AddonVersion",
          "value": "v1.66.0-eksbuild.1"
        },
        {
          "key": "ClusterName",
          "value": "tb56lfotn8retqg9nve4"
        },
        {
          "key": "CreatedAt",
          "value": "2026-10-07T07:41:11.253Z"
        },
        {
          "key": "Health",
          "value": "{Issues:[]}"
        },
        {
          "key": "ModifiedAt",
          "value": "2026-10-07T07:41:11.264Z"
        },
        {
          "key": "Status",
          "value": "CREATING"
        },
        {
          "key": "AddonArn",
          "value": "arn:aws:eks:ap-northeast-2:635484366616:addon/tb56lfotn8retqg9nve4/eks-pod-identity-agent/6ad08aa7-93b1-d073-af49-b4207e0bef9f"
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
          "value": "tb56lfotn8retqg9nve4"
        },
        {
          "key": "CreatedAt",
          "value": "2026-10-07T07:41:11.581Z"
        },
        {
          "key": "Health",
          "value": "{Issues:[]}"
        },
        {
          "key": "ModifiedAt",
          "value": "2026-10-07T07:41:11.594Z"
        },
        {
          "key": "Status",
          "value": "CREATING"
        }
      ]
    },
    "Status": "Active",
    "CreatedTime": "2026-10-07T07:41:09.979Z",
    "KeyValueList": [
      {
        "key": "Arn",
        "value": "arn:aws:eks:ap-northeast-2:635484366616:cluster/tb56lfotn8retqg9nve4"
      },
      {
        "key": "CertificateAuthority",
        "value": "{Data:LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURNakNDQWhxZ0F3SUJBZ0lSQUoxMU5GOUhhYURCdndqVmdHNGlGb3d3RFFZSktvWklodmNOQVFFTEJRQXcKSnpFUU1BNEdBMVVFQ2hNSFFWZFRJRVZMVXpFVE1CRUdBMVVFQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TmpFdwpNRGN3TnpReE1UWmFGdzB6TVRFd01EWXdOelF4TVRaYU1DY3hFREFPQmdOVkJBb1RCMEZYVXlCRlMxTXhFekFSCkJnTlZCQU1UQ210MVltVnlibVYwWlhNd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUIKQVFEVGJDc2QwRFBDdC90QVUrM1V5d3NxRHlHbDlHc0JNU2JyeXZLUi9HcGZ3Nmh5UVdjckZzQS9hTkZobWxVawo4dTVQcEpHeHNub3lSOFEyYUt6VnE3V0YyNUI4NUhmUXpsTmFtNCtXNGFrYXozdUdpSHUzZTI2aUdBRVE0eWVlCkhVZ09UUmpFRTc2V0crcUIzcjJrK3V1Uk8ydzhTMDMyMStpMkZvRUROcm4vanNRdVVOSW9XZUVtRDdHUkhXVFIKVzBuNkc1TU1hVzQyUWJEQ3l3Z2VLcUlmQStUTVZRUldlV28vL2JaNW0zZVZ6R3NuQkRLbFB5Q0Uydm5ETTZ5NQp2dVB2SW9iK3JWQWYvMGROZWsxRUh3ZUtYYzdTUm84N1M5TkRXcTZmUU9JcmxtNzJuU3l1K2doZVhKTWhWcFJQCmJEYi8yWHBXTS9ZQzVYZUlPd25uYTZnZkFnTUJBQUdqV1RCWE1BNEdBMVVkRHdFQi93UUVBd0lDcERBUEJnTlYKSFJNQkFmOEVCVEFEQVFIL01CMEdBMVVkRGdRV0JCUm54YUdXU2ZFdDZvTFdnUE9ENVh2OW5YZjJpREFWQmdOVgpIUkVFRGpBTWdncHJkV0psY201bGRHVnpNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUIydytvaE9KajlNSmlqCkdiSWhXYjZLZ2MrODZZdmR2MFk1b2dlekNrdHFicmJITTRzQjdxWmhhd2MxUURTdmZFVG9HS21jblNXeVlxaDQKV0VEZFRhODNxU1dTV3E1TnFQYndWSDJSd1hWWUF0cFBpVi9HSnF6VWhaSlMzTit0aFhMSGYxOXYwaVo4VHhJeQpIOG5uWUo3UEFLakh6c1VSVVNQY2l0YlZYOWxadFozUk9mMUdTQ25yQXVWTFVNTHU3MGNrUG1QYXhvclRkeFVICi90Q0lrR1pwckl5QWpqYXAzQ2gvMlBOZ2dRVzFnemIyTTNWcjVTbHZwdy9iVkNsc0pPS2t0aHJkbHcycEc4UDYKdk5DODJmRmI2dEVoY2FBZ1VTemlTWVRVN1IrUHZ3aTdHdElhYm81NCtxTURGSGhRTGJjMW90a2o4OElmdFRsRQpnYmVqd0lTOQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==}"
      },
      {
        "key": "CreatedAt",
        "value": "2026-10-07T07:41:09.979Z"
      },
      {
        "key": "Endpoint",
        "value": "https://2E128CC0E45344BF3744FDA996EEB348.gr7.ap-northeast-2.eks.amazonaws.com"
      },
      {
        "key": "Identity",
        "value": "{Oidc:{Issuer:https://oidc.eks.ap-northeast-2.amazonaws.com/id/2E128CC0E45344BF3744FDA996EEB348}}"
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
        "value": "tb56lfotn8retqg9nve4"
      },
      {
        "key": "PlatformVersion",
        "value": "eks.34"
      },
      {
        "key": "ResourcesVpcConfig",
        "value": "{ClusterSecurityGroupId:sg-0150e18b69f6edc2f,EndpointPrivateAccess:false,EndpointPublicAccess:true,PublicAccessCidrs:[0.0.0.0/0],SecurityGroupIds:[sg-00b0fa9c3c57c2cc5],SubnetIds:[subnet-05247863bd834f2df,subnet-0b020a2249bb6475b],VpcId:vpc-01af9f247d39edce6}"
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
        "value": "{Name:tb56lfotn8retqg9nve4}"
      },
      {
        "key": "Version",
        "value": "1.34"
      }
    ]
  }
}
```

</details>

