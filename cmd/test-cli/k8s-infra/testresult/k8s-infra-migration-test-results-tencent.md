# CM-Beetle K8s Infra Migration Test Results — Tencent-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → validate → migrate → list → get → workload → delete → residual.

## Environment

- CSP / Region: tencent / ap-seoul
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (cd1f3a2)
- Git Commit: cd1f3a2
- Namespace: mig01
- Test Date: 2026-10-07 16:41:45 KST
- Cluster ID: mig05-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 21ms |
| 2 | POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation) | ✅ **PASS** | 30ms |
| 3 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 4m45.083s |
| 4 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 2ms |
| 5 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 3.003s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ❌ **FAIL** | 2m6.621s |
| 7 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 36.335s |
| 8 | Residual resource check (Tumblebug) | ✅ **PASS** | 4ms |

**Overall Result**: 7/8 steps passed ❌

**Total Duration**: 7m31s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 21ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.32.2)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=tencent+ap-seoul+bf1.large8 nodes=2

### Step 2 — POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)

- **Duration**: 30ms
- **Status Code**: 200

- ✅ Target K8s infra model is valid for migration (0 issues)

### Step 3 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 4m45.083s
- **Status Code**: 202

- ℹ️  nameSeed: mig05
- ℹ️  async reqId: 1791358905694941639
- ℹ️  cluster id: mig05-on-prem-k8s-cluster
- ℹ️  elapsed: 4m45s
- ✅ status: Active

### Step 4 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 2ms
- **Status Code**: 200

- ✅ migrated cluster present in list (8 total)

### Step 5 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 3.003s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=tencent+ap-seoul+bf1.large8, nodes=2)
- ✅ version: 1.32.2 (recommended 1.32.2)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 2m6.621s
- **Error**: kubeconfig not ready within 120s


### Step 7 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 36.335s
- **Status Code**: 200

- ✅ deleted on attempt 1 (36s)

### Step 8 — Residual resource check (Tumblebug)

- **Duration**: 4ms

- ℹ️  VNet mig05-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig05-k8s-sg still exists (known gap)
- ℹ️  SshKey mig05-k8s-sshkey still exists (known gap)

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
  "id": "mig05-on-prem-k8s-cluster",
  "uid": "tbqnipqvm7kp78l34k04",
  "name": "mig05-on-prem-k8s-cluster",
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
    "sys.connectionName": "tencent-ap-seoul",
    "sys.createdTime": "2026-10-07 07:41:57 +0000 UTC",
    "sys.cspResourceId": "cls-lj87sx9d",
    "sys.cspResourceName": "tbqnipqvm7kp78l34k04",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig05-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig05-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbqnipqvm7kp78l34k04",
    "sys.version": "1.32.2"
  },
  "systemLabel": "",
  "version": "1.32.2",
  "network": {
    "vNetId": "mig05-k8s-vpc",
    "subnetIds": [
      "mig05-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "mig05-k8s-sg"
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
      "sshKeyId": "mig05-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 2,
      "maxNodeSize": 2,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "ins-9nawhbof",
          "cspResourceId": "ins-9nawhbof"
        },
        {
          "cspResourceName": "ins-rcfembsh",
          "cspResourceId": "ins-rcfembsh"
        }
      ],
      "keyValueList": [
        {
          "key": "NodePoolId",
          "value": "np-3n8n34un"
        },
        {
          "key": "Name",
          "value": "workers1"
        },
        {
          "key": "ClusterInstanceId",
          "value": "cls-lj87sx9d"
        },
        {
          "key": "LifeState",
          "value": "normal"
        },
        {
          "key": "LaunchConfigurationId",
          "value": "asc-9dlteodp"
        },
        {
          "key": "AutoscalingGroupId",
          "value": "asg-b2gc3rcl"
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
          "value": "{Key:CB-SPIDER-PMKS-SUBNET-ID,Value:subnet-1u5rqx0f}; {Key:CB-SPIDER-PMKS-SECURITYGROUP-ID,Value:sg-rsc1c5rn}"
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
      "cspResourceId": "np-3n8n34un",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "np-3n8n34un"
        },
        "ImageIID": {
          "NameId": "img-4wpaazux",
          "SystemId": "img-4wpaazux"
        },
        "VMSpecName": "BF1.LARGE8",
        "RootDiskType": "CLOUD_PREMIUM",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tb8n9jfjjiascnfn0qpb",
          "SystemId": "skey-3dm9bi4f"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 2,
        "MaxNodeSize": 2,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "ins-9nawhbof",
            "SystemId": "ins-9nawhbof"
          },
          {
            "NameId": "ins-rcfembsh",
            "SystemId": "ins-rcfembsh"
          }
        ],
        "KeyValueList": [
          {
            "key": "NodePoolId",
            "value": "np-3n8n34un"
          },
          {
            "key": "Name",
            "value": "workers1"
          },
          {
            "key": "ClusterInstanceId",
            "value": "cls-lj87sx9d"
          },
          {
            "key": "LifeState",
            "value": "normal"
          },
          {
            "key": "LaunchConfigurationId",
            "value": "asc-9dlteodp"
          },
          {
            "key": "AutoscalingGroupId",
            "value": "asg-b2gc3rcl"
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
            "value": "{Key:CB-SPIDER-PMKS-SUBNET-ID,Value:subnet-1u5rqx0f}; {Key:CB-SPIDER-PMKS-SECURITYGROUP-ID,Value:sg-rsc1c5rn}"
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
    "endpoint": "First, add a nodegroup.",
    "kubeconfig": "Kubeconfig is not ready yet!"
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-10-07T07:41:57Z",
  "keyValueList": [
    {
      "key": "ClusterId",
      "value": "cls-lj87sx9d"
    },
    {
      "key": "ClusterName",
      "value": "tbqnipqvm7kp78l34k04"
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
      "value": "{ClusterCIDR:172.17.0.0/16,IgnoreClusterCIDRConflict:false,MaxNodePodNum:256,MaxClusterServiceNum:4096,Ipvs:false,VpcId:vpc-j2l404me,Cni:true,KubeProxyMode:,ServiceCIDR:10.200.16.0/20,IgnoreServiceCIDRConflict:false,IsDualStack:false,Ipv6ServiceCIDR:,CiliumMode:,SubnetId:,DataPlaneV2:false}"
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
      "value": "{ResourceType:cluster,Tags:[{Key:CB-SPIDER-PMKS-SUBNET-ID,Value:subnet-1u5rqx0f},{Key:CB-SPIDER-PMKS-SECURITYGROUP-ID,Value:sg-rsc1c5rn}]}"
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
      "value": "2026-10-07T07:41:57Z"
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
  "cspResourceName": "tbqnipqvm7kp78l34k04",
  "cspResourceId": "cls-lj87sx9d",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbqnipqvm7kp78l34k04",
      "SystemId": "cls-lj87sx9d"
    },
    "Version": "1.32.2",
    "Network": {
      "VpcIID": {
        "NameId": "tbcphgtktgjeolhbollq",
        "SystemId": "vpc-j2l404me"
      },
      "SubnetIIDs": [
        {
          "NameId": "tb4lvsba1mqmeuur32l4",
          "SystemId": "subnet-1u5rqx0f"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "tbjfd5e3olk92e8s0aou",
          "SystemId": "sg-rsc1c5rn"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "np-3n8n34un"
        },
        "ImageIID": {
          "NameId": "img-4wpaazux",
          "SystemId": "img-4wpaazux"
        },
        "VMSpecName": "BF1.LARGE8",
        "RootDiskType": "CLOUD_PREMIUM",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tb8n9jfjjiascnfn0qpb",
          "SystemId": "skey-3dm9bi4f"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 2,
        "MaxNodeSize": 2,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "ins-9nawhbof",
            "SystemId": "ins-9nawhbof"
          },
          {
            "NameId": "ins-rcfembsh",
            "SystemId": "ins-rcfembsh"
          }
        ],
        "KeyValueList": [
          {
            "key": "NodePoolId",
            "value": "np-3n8n34un"
          },
          {
            "key": "Name",
            "value": "workers1"
          },
          {
            "key": "ClusterInstanceId",
            "value": "cls-lj87sx9d"
          },
          {
            "key": "LifeState",
            "value": "normal"
          },
          {
            "key": "LaunchConfigurationId",
            "value": "asc-9dlteodp"
          },
          {
            "key": "AutoscalingGroupId",
            "value": "asg-b2gc3rcl"
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
            "value": "{Key:CB-SPIDER-PMKS-SUBNET-ID,Value:subnet-1u5rqx0f}; {Key:CB-SPIDER-PMKS-SECURITYGROUP-ID,Value:sg-rsc1c5rn}"
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
      "Endpoint": "First, add a nodegroup.",
      "Kubeconfig": "Kubeconfig is not ready yet!"
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-10-07T07:41:57Z",
    "KeyValueList": [
      {
        "key": "ClusterId",
        "value": "cls-lj87sx9d"
      },
      {
        "key": "ClusterName",
        "value": "tbqnipqvm7kp78l34k04"
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
        "value": "{ClusterCIDR:172.17.0.0/16,IgnoreClusterCIDRConflict:false,MaxNodePodNum:256,MaxClusterServiceNum:4096,Ipvs:false,VpcId:vpc-j2l404me,Cni:true,KubeProxyMode:,ServiceCIDR:10.200.16.0/20,IgnoreServiceCIDRConflict:false,IsDualStack:false,Ipv6ServiceCIDR:,CiliumMode:,SubnetId:,DataPlaneV2:false}"
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
        "value": "{ResourceType:cluster,Tags:[{Key:CB-SPIDER-PMKS-SUBNET-ID,Value:subnet-1u5rqx0f},{Key:CB-SPIDER-PMKS-SECURITYGROUP-ID,Value:sg-rsc1c5rn}]}"
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
        "value": "2026-10-07T07:41:57Z"
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

