# CM-Beetle K8s Infra Migration Test Results — Tencent-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → migrate → list → get (verified against
> the recommendation) → delete → residual resource check.

## Environment

- CSP / Region: tencent / ap-seoul
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (e7e0c24)
- Git Commit: e7e0c24
- Namespace: mig01
- Test Date: 2026-09-29 10:09:13 KST
- Cluster ID: k8s4csp02-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 21ms |
| 2 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 6m0.141s |
| 3 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 2ms |
| 4 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 3.23s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ❌ **FAIL** | 10m34.117s |
| 6 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 38.919s |
| 7 | Residual resource check (Tumblebug) | ✅ **PASS** | 4ms |

**Overall Result**: 6/7 steps passed ❌

**Total Duration**: 17m16s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 21ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.32.2)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=tencent+ap-seoul+bf1.large8 nodes=2

### Step 2 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 6m0.141s
- **Status Code**: 202

- ℹ️  nameSeed: k8s4csp02
- ℹ️  async reqId: 1790644153748123553
- ℹ️  cluster id: k8s4csp02-on-prem-k8s-cluster
- ℹ️  elapsed: 6m0s
- ✅ status: Active

### Step 3 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 2ms
- **Status Code**: 200

- ✅ migrated cluster present in list (1 total)

### Step 4 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 3.23s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=tencent+ap-seoul+bf1.large8, nodes=2)
- ✅ version: 1.32.2 (recommended 1.32.2)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 10m34.117s
- **Error**: kubeconfig not ready within 600s


### Step 6 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 38.919s
- **Status Code**: 200

- ✅ deleted on attempt 1 (38s)

### Step 7 — Residual resource check (Tumblebug)

- **Duration**: 4ms

- ℹ️  VNet k8s4csp02-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup k8s4csp02-k8s-sg still exists (known gap)
- ℹ️  SshKey k8s4csp02-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for tencent ap-seoul (source: v1.32.3 → target: v1.32.2)",
  "targetCloud": {
    "csp": "tencent",
    "region": "ap-seoul"
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
    "connectionName": "tencent-ap-seoul",
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
    "connectionName": "tencent-ap-seoul",
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
      "connectionName": "tencent-ap-seoul",
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
    "connectionName": "tencent-ap-seoul",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "1.32.2",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": [
      {
        "name": "workers1",
        "imageId": "default",
        "specId": "tencent+ap-seoul+bf1.large8",
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
  "id": "k8s4csp02-on-prem-k8s-cluster",
  "uid": "tbttntecb61l03n9c2hh",
  "name": "k8s4csp02-on-prem-k8s-cluster",
  "connectionName": "tencent-ap-seoul",
  "connectionConfig": {
    "configName": "tencent-ap-seoul",
    "providerName": "tencent",
    "driverName": "tencent-driver-v1.0.so",
    "credentialName": "tencent",
    "credentialHolder": "admin",
    "regionZoneInfoName": "tencent-ap-seoul",
    "regionZoneInfo": {
      "assignedRegion": "ap-seoul",
      "assignedZone": "ap-seoul-1"
    },
    "regionDetail": {
      "regionId": "ap-seoul",
      "regionName": "ap-seoul",
      "description": "Seoul",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.566536,
        "longitude": 126.977966
      },
      "zones": [
        "ap-seoul-1",
        "ap-seoul-2"
      ]
    },
    "regionRepresentative": true,
    "verified": true
  },
  "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
  "systemMessage": "",
  "label": {
    "CB-SPIDER-PMKS-SECURITYGROUP-ID": "sg-l7hdrwah",
    "CB-SPIDER-PMKS-SUBNET-ID": "subnet-cby6bpod",
    "sys.connectionName": "tencent-ap-seoul",
    "sys.createdTime": "2026-09-29 01:10:10 +0000 UTC",
    "sys.cspResourceId": "cls-ooqqxk9z",
    "sys.cspResourceName": "tbttntecb61l03n9c2hh",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "k8s4csp02-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "k8s4csp02-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbttntecb61l03n9c2hh",
    "sys.version": "1.32.2"
  },
  "systemLabel": "",
  "version": "1.32.2",
  "network": {
    "vNetId": "k8s4csp02-k8s-vpc",
    "subnetIds": [
      "k8s4csp02-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "k8s4csp02-k8s-sg"
    ],
    "keyValueList": null
  },
  "k8sNodeGroupList": [
    {
      "id": "workers1",
      "name": "workers1",
      "imageId": "default",
      "specId": "tencent+ap-seoul+bf1.large8",
      "rootDiskType": "CLOUD_PREMIUM",
      "rootDiskSize": 100,
      "sshKeyId": "k8s4csp02-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 2,
      "maxNodeSize": 2,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "ins-2etqzrrz",
          "cspResourceId": "ins-2etqzrrz"
        },
        {
          "cspResourceName": "ins-5liwpcgn",
          "cspResourceId": "ins-5liwpcgn"
        }
      ],
      "keyValueList": [
        {
          "key": "NodePoolId",
          "value": "np-8t8clgdj"
        },
        {
          "key": "Name",
          "value": "workers1"
        },
        {
          "key": "ClusterInstanceId",
          "value": "cls-ooqqxk9z"
        },
        {
          "key": "LifeState",
          "value": "normal"
        },
        {
          "key": "LaunchConfigurationId",
          "value": "asc-e85gaze7"
        },
        {
          "key": "AutoscalingGroupId",
          "value": "asg-11ojbivl"
        },
        {
          "key": "NodeCountSummary",
          "value": "{ManuallyAdded:{Joining:0,Initializing:0,Normal:0,Total:0},AutoscalingAdded:{Joining:0,Initializing:2,Normal:0,Total:2}}"
        },
        {
          "key": "AutoscalingGroupStatus",
          "value": "disabled"
        },
        {
          "key": "MaxNodesNum",
          "value": "2"
        },
        {
          "key": "MinNodesNum",
          "value": "2"
        },
        {
          "key": "DesiredNodesNum",
          "value": "2"
        },
        {
          "key": "RuntimeConfig",
          "value": "{RuntimeType:containerd,RuntimeVersion:1.6.9}"
        },
        {
          "key": "NodePoolOs",
          "value": "ubuntu16.04.1 LTSx86_64"
        },
        {
          "key": "OsCustomizeType",
          "value": "GENERAL"
        },
        {
          "key": "DesiredPodNum",
          "value": "256"
        },
        {
          "key": "Tags",
          "value": "{Key:sys.connectionName,Value:tencent-ap-seoul}; {Key:sys.createdTime,Value:2026-09-29 01:10:10 +0000 UTC}; {Key:sys.description,Value:Migrated from on-premise K8s cluster (v1.32.3, 2 workers)}; {Key:sys.cspResourceName,Value:tbttntecb61l03n9c2hh}; {Key:sys.version,Value:1.32.2}; {Key:sys.uid,Value:tbttntecb61l03n9c2hh}; {Key:sys.name,Value:k8s4csp02-on-prem-k8s-cluster}; {Key:sys.cspResourceId,Value:cls-ooqqxk9z}; {Key:sys.id,Value:k8s4csp02-on-prem-k8s-cluster}; {Key:sys.namespace,Value:mig01}; {Key:sys.labelType,Value:k8s}; {Key:sys.manager,Value:cb-tumblebug}; {Key:CB-SPIDER-PMKS-SUBNET-ID,Value:subnet-cby6bpod}; {Key:CB-SPIDER-PMKS-SECURITYGROUP-ID,Value:sg-l7hdrwah}"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "ExtraArgs",
          "value": "{}"
        },
        {
          "key": "GPUArgs",
          "value": "{CUDA:{Name:,Version:},CUDNN:{Name:,Version:,DevName:,DocName:},CustomDriver:{Address:},Driver:{Name:,Version:},MIGEnable:false}"
        },
        {
          "key": "Unschedulable",
          "value": "0"
        }
      ],
      "cspResourceName": "workers1",
      "cspResourceId": "np-8t8clgdj",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "np-8t8clgdj"
        },
        "ImageIID": {
          "NameId": "img-4wpaazux",
          "SystemId": "img-4wpaazux"
        },
        "VMSpecName": "BF1.LARGE8",
        "RootDiskType": "CLOUD_PREMIUM",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tb2gnq0hjcbhh1j7iih1",
          "SystemId": "skey-jtm9tnln"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 2,
        "MaxNodeSize": 2,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "ins-2etqzrrz",
            "SystemId": "ins-2etqzrrz"
          },
          {
            "NameId": "ins-5liwpcgn",
            "SystemId": "ins-5liwpcgn"
          }
        ],
        "KeyValueList": [
          {
            "key": "NodePoolId",
            "value": "np-8t8clgdj"
          },
          {
            "key": "Name",
            "value": "workers1"
          },
          {
            "key": "ClusterInstanceId",
            "value": "cls-ooqqxk9z"
          },
          {
            "key": "LifeState",
            "value": "normal"
          },
          {
            "key": "LaunchConfigurationId",
            "value": "asc-e85gaze7"
          },
          {
            "key": "AutoscalingGroupId",
            "value": "asg-11ojbivl"
          },
          {
            "key": "NodeCountSummary",
            "value": "{ManuallyAdded:{Joining:0,Initializing:0,Normal:0,Total:0},AutoscalingAdded:{Joining:0,Initializing:2,Normal:0,Total:2}}"
          },
          {
            "key": "AutoscalingGroupStatus",
            "value": "disabled"
          },
          {
            "key": "MaxNodesNum",
            "value": "2"
          },
          {
            "key": "MinNodesNum",
            "value": "2"
          },
          {
            "key": "DesiredNodesNum",
            "value": "2"
          },
          {
            "key": "RuntimeConfig",
            "value": "{RuntimeType:containerd,RuntimeVersion:1.6.9}"
          },
          {
            "key": "NodePoolOs",
            "value": "ubuntu16.04.1 LTSx86_64"
          },
          {
            "key": "OsCustomizeType",
            "value": "GENERAL"
          },
          {
            "key": "DesiredPodNum",
            "value": "256"
          },
          {
            "key": "Tags",
            "value": "{Key:sys.connectionName,Value:tencent-ap-seoul}; {Key:sys.createdTime,Value:2026-09-29 01:10:10 +0000 UTC}; {Key:sys.description,Value:Migrated from on-premise K8s cluster (v1.32.3, 2 workers)}; {Key:sys.cspResourceName,Value:tbttntecb61l03n9c2hh}; {Key:sys.version,Value:1.32.2}; {Key:sys.uid,Value:tbttntecb61l03n9c2hh}; {Key:sys.name,Value:k8s4csp02-on-prem-k8s-cluster}; {Key:sys.cspResourceId,Value:cls-ooqqxk9z}; {Key:sys.id,Value:k8s4csp02-on-prem-k8s-cluster}; {Key:sys.namespace,Value:mig01}; {Key:sys.labelType,Value:k8s}; {Key:sys.manager,Value:cb-tumblebug}; {Key:CB-SPIDER-PMKS-SUBNET-ID,Value:subnet-cby6bpod}; {Key:CB-SPIDER-PMKS-SECURITYGROUP-ID,Value:sg-l7hdrwah}"
          },
          {
            "key": "DeletionProtection",
            "value": "false"
          },
          {
            "key": "ExtraArgs",
            "value": "{}"
          },
          {
            "key": "GPUArgs",
            "value": "{CUDA:{Name:,Version:},CUDNN:{Name:,Version:,DevName:,DocName:},CustomDriver:{Address:},Driver:{Name:,Version:},MIGEnable:false}"
          },
          {
            "key": "Unschedulable",
            "value": "0"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "Preparing....",
    "kubeconfig": "Kubeconfig is not ready yet!"
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-09-29T01:10:10Z",
  "keyValueList": [
    {
      "key": "ClusterId",
      "value": "cls-ooqqxk9z"
    },
    {
      "key": "ClusterName",
      "value": "tbttntecb61l03n9c2hh"
    },
    {
      "key": "ClusterVersion",
      "value": "1.32.2"
    },
    {
      "key": "ClusterOs",
      "value": "ubuntu16.04.1 LTSx86_64"
    },
    {
      "key": "ClusterType",
      "value": "MANAGED_CLUSTER"
    },
    {
      "key": "ClusterNetworkSettings",
      "value": "{ClusterCIDR:172.17.0.0/16,IgnoreClusterCIDRConflict:false,MaxNodePodNum:256,MaxClusterServiceNum:4096,Ipvs:false,VpcId:vpc-ih4b782u,Cni:true,KubeProxyMode:,ServiceCIDR:10.200.16.0/20,IgnoreServiceCIDRConflict:false,IsDualStack:false,Ipv6ServiceCIDR:,CiliumMode:,SubnetId:,DataPlaneV2:false}"
    },
    {
      "key": "ClusterNodeNum",
      "value": "2"
    },
    {
      "key": "ProjectId",
      "value": "0"
    },
    {
      "key": "TagSpecification",
      "value": "{ResourceType:cluster,Tags:[{Key:sys.connectionName,Value:tencent-ap-seoul},{Key:sys.createdTime,Value:2026-09-29 01:10:10 +0000 UTC},{Key:sys.description,Value:Migrated from on-premise K8s cluster (v1.32.3, 2 workers)},{Key:sys.cspResourceName,Value:tbttntecb61l03n9c2hh},{Key:sys.version,Value:1.32.2},{Key:sys.uid,Value:tbttntecb61l03n9c2hh},{Key:sys.name,Value:k8s4csp02-on-prem-k8s-cluster},{Key:sys.cspResourceId,Value:cls-ooqqxk9z},{Key:sys.id,Value:k8s4csp02-on-prem-k8s-cluster},{Key:sys.namespace,Value:mig01},{Key:sys.labelType,Value:k8s},{Key:sys.manager,Value:cb-tumblebug},{Key:CB-SPIDER-PMKS-SUBNET-ID,Value:subnet-cby6bpod},{Key:CB-SPIDER-PMKS-SECURITYGROUP-ID,Value:sg-l7hdrwah}]}"
    },
    {
      "key": "ClusterStatus",
      "value": "Running"
    },
    {
      "key": "Property",
      "value": "{\\NodeNameType\\:\\lan-ip\\,\\NetworkType\\:\\GR\\,\\IsNetworkWithApp\\:true}"
    },
    {
      "key": "ClusterMaterNodeNum",
      "value": "1"
    },
    {
      "key": "ContainerRuntime",
      "value": "containerd"
    },
    {
      "key": "CreatedTime",
      "value": "2026-09-29T01:10:10Z"
    },
    {
      "key": "DeletionProtection",
      "value": "false"
    },
    {
      "key": "EnableExternalNode",
      "value": "false"
    },
    {
      "key": "ClusterLevel",
      "value": "L5"
    },
    {
      "key": "AutoUpgradeClusterLevel",
      "value": "true"
    },
    {
      "key": "QGPUShareEnable",
      "value": "false"
    },
    {
      "key": "RuntimeVersion",
      "value": "1.6.9"
    },
    {
      "key": "ClusterEtcdNodeNum",
      "value": "0"
    },
    {
      "key": "IsHighAvailability",
      "value": "true"
    },
    {
      "key": "SecurityModeConfig",
      "value": "{Enabled:false}"
    }
  ],
  "cspResourceName": "tbttntecb61l03n9c2hh",
  "cspResourceId": "cls-ooqqxk9z",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbttntecb61l03n9c2hh",
      "SystemId": "cls-ooqqxk9z"
    },
    "Version": "1.32.2",
    "Network": {
      "VpcIID": {
        "NameId": "tbsrck7t004347r1bqkn",
        "SystemId": "vpc-ih4b782u"
      },
      "SubnetIIDs": [
        {
          "NameId": "tb844l56oqt725adu087",
          "SystemId": "subnet-cby6bpod"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "tb80l6s2jmor9p8b1a8o",
          "SystemId": "sg-l7hdrwah"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "np-8t8clgdj"
        },
        "ImageIID": {
          "NameId": "img-4wpaazux",
          "SystemId": "img-4wpaazux"
        },
        "VMSpecName": "BF1.LARGE8",
        "RootDiskType": "CLOUD_PREMIUM",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tb2gnq0hjcbhh1j7iih1",
          "SystemId": "skey-jtm9tnln"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 2,
        "MaxNodeSize": 2,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "ins-2etqzrrz",
            "SystemId": "ins-2etqzrrz"
          },
          {
            "NameId": "ins-5liwpcgn",
            "SystemId": "ins-5liwpcgn"
          }
        ],
        "KeyValueList": [
          {
            "key": "NodePoolId",
            "value": "np-8t8clgdj"
          },
          {
            "key": "Name",
            "value": "workers1"
          },
          {
            "key": "ClusterInstanceId",
            "value": "cls-ooqqxk9z"
          },
          {
            "key": "LifeState",
            "value": "normal"
          },
          {
            "key": "LaunchConfigurationId",
            "value": "asc-e85gaze7"
          },
          {
            "key": "AutoscalingGroupId",
            "value": "asg-11ojbivl"
          },
          {
            "key": "NodeCountSummary",
            "value": "{ManuallyAdded:{Joining:0,Initializing:0,Normal:0,Total:0},AutoscalingAdded:{Joining:0,Initializing:2,Normal:0,Total:2}}"
          },
          {
            "key": "AutoscalingGroupStatus",
            "value": "disabled"
          },
          {
            "key": "MaxNodesNum",
            "value": "2"
          },
          {
            "key": "MinNodesNum",
            "value": "2"
          },
          {
            "key": "DesiredNodesNum",
            "value": "2"
          },
          {
            "key": "RuntimeConfig",
            "value": "{RuntimeType:containerd,RuntimeVersion:1.6.9}"
          },
          {
            "key": "NodePoolOs",
            "value": "ubuntu16.04.1 LTSx86_64"
          },
          {
            "key": "OsCustomizeType",
            "value": "GENERAL"
          },
          {
            "key": "DesiredPodNum",
            "value": "256"
          },
          {
            "key": "Tags",
            "value": "{Key:sys.connectionName,Value:tencent-ap-seoul}; {Key:sys.createdTime,Value:2026-09-29 01:10:10 +0000 UTC}; {Key:sys.description,Value:Migrated from on-premise K8s cluster (v1.32.3, 2 workers)}; {Key:sys.cspResourceName,Value:tbttntecb61l03n9c2hh}; {Key:sys.version,Value:1.32.2}; {Key:sys.uid,Value:tbttntecb61l03n9c2hh}; {Key:sys.name,Value:k8s4csp02-on-prem-k8s-cluster}; {Key:sys.cspResourceId,Value:cls-ooqqxk9z}; {Key:sys.id,Value:k8s4csp02-on-prem-k8s-cluster}; {Key:sys.namespace,Value:mig01}; {Key:sys.labelType,Value:k8s}; {Key:sys.manager,Value:cb-tumblebug}; {Key:CB-SPIDER-PMKS-SUBNET-ID,Value:subnet-cby6bpod}; {Key:CB-SPIDER-PMKS-SECURITYGROUP-ID,Value:sg-l7hdrwah}"
          },
          {
            "key": "DeletionProtection",
            "value": "false"
          },
          {
            "key": "ExtraArgs",
            "value": "{}"
          },
          {
            "key": "GPUArgs",
            "value": "{CUDA:{Name:,Version:},CUDNN:{Name:,Version:,DevName:,DocName:},CustomDriver:{Address:},Driver:{Name:,Version:},MIGEnable:false}"
          },
          {
            "key": "Unschedulable",
            "value": "0"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "Preparing....",
      "Kubeconfig": "Kubeconfig is not ready yet!"
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-09-29T01:10:10Z",
    "KeyValueList": [
      {
        "key": "ClusterId",
        "value": "cls-ooqqxk9z"
      },
      {
        "key": "ClusterName",
        "value": "tbttntecb61l03n9c2hh"
      },
      {
        "key": "ClusterVersion",
        "value": "1.32.2"
      },
      {
        "key": "ClusterOs",
        "value": "ubuntu16.04.1 LTSx86_64"
      },
      {
        "key": "ClusterType",
        "value": "MANAGED_CLUSTER"
      },
      {
        "key": "ClusterNetworkSettings",
        "value": "{ClusterCIDR:172.17.0.0/16,IgnoreClusterCIDRConflict:false,MaxNodePodNum:256,MaxClusterServiceNum:4096,Ipvs:false,VpcId:vpc-ih4b782u,Cni:true,KubeProxyMode:,ServiceCIDR:10.200.16.0/20,IgnoreServiceCIDRConflict:false,IsDualStack:false,Ipv6ServiceCIDR:,CiliumMode:,SubnetId:,DataPlaneV2:false}"
      },
      {
        "key": "ClusterNodeNum",
        "value": "2"
      },
      {
        "key": "ProjectId",
        "value": "0"
      },
      {
        "key": "TagSpecification",
        "value": "{ResourceType:cluster,Tags:[{Key:sys.connectionName,Value:tencent-ap-seoul},{Key:sys.createdTime,Value:2026-09-29 01:10:10 +0000 UTC},{Key:sys.description,Value:Migrated from on-premise K8s cluster (v1.32.3, 2 workers)},{Key:sys.cspResourceName,Value:tbttntecb61l03n9c2hh},{Key:sys.version,Value:1.32.2},{Key:sys.uid,Value:tbttntecb61l03n9c2hh},{Key:sys.name,Value:k8s4csp02-on-prem-k8s-cluster},{Key:sys.cspResourceId,Value:cls-ooqqxk9z},{Key:sys.id,Value:k8s4csp02-on-prem-k8s-cluster},{Key:sys.namespace,Value:mig01},{Key:sys.labelType,Value:k8s},{Key:sys.manager,Value:cb-tumblebug},{Key:CB-SPIDER-PMKS-SUBNET-ID,Value:subnet-cby6bpod},{Key:CB-SPIDER-PMKS-SECURITYGROUP-ID,Value:sg-l7hdrwah}]}"
      },
      {
        "key": "ClusterStatus",
        "value": "Running"
      },
      {
        "key": "Property",
        "value": "{\\NodeNameType\\:\\lan-ip\\,\\NetworkType\\:\\GR\\,\\IsNetworkWithApp\\:true}"
      },
      {
        "key": "ClusterMaterNodeNum",
        "value": "1"
      },
      {
        "key": "ContainerRuntime",
        "value": "containerd"
      },
      {
        "key": "CreatedTime",
        "value": "2026-09-29T01:10:10Z"
      },
      {
        "key": "DeletionProtection",
        "value": "false"
      },
      {
        "key": "EnableExternalNode",
        "value": "false"
      },
      {
        "key": "ClusterLevel",
        "value": "L5"
      },
      {
        "key": "AutoUpgradeClusterLevel",
        "value": "true"
      },
      {
        "key": "QGPUShareEnable",
        "value": "false"
      },
      {
        "key": "RuntimeVersion",
        "value": "1.6.9"
      },
      {
        "key": "ClusterEtcdNodeNum",
        "value": "0"
      },
      {
        "key": "IsHighAvailability",
        "value": "true"
      },
      {
        "key": "SecurityModeConfig",
        "value": "{Enabled:false}"
      }
    ]
  }
}
```

</details>

