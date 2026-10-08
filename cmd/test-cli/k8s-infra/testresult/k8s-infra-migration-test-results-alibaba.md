# CM-Beetle K8s Infra Migration Test Results — Alibaba-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → validate → migrate → list → get → workload → delete → residual.

## Environment

- CSP / Region: alibaba / ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (cd1f3a2)
- Git Commit: cd1f3a2
- Namespace: mig01
- Test Date: 2026-10-07 16:41:35 KST
- Cluster ID: mig04-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 20ms |
| 2 | POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation) | ✅ **PASS** | 31ms |
| 3 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 6m0.102s |
| 4 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 2ms |
| 5 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 1ms |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 21.392s |
| 7 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 13m43.715s |
| 8 | Residual resource check (Tumblebug) | ✅ **PASS** | 4ms |

**Overall Result**: 8/8 steps passed ✅

**Total Duration**: 20m5s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 20ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.35.7-aliyun.1)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=alibaba+ap-northeast-2+ecs.e-c1m2.xlarge nodes=2

### Step 2 — POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)

- **Duration**: 31ms
- **Status Code**: 200

- ✅ Target K8s infra model is valid for migration (0 issues)

### Step 3 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 6m0.102s
- **Status Code**: 202

- ℹ️  nameSeed: mig04
- ℹ️  async reqId: 1791358895700471574
- ℹ️  cluster id: mig04-on-prem-k8s-cluster
- ℹ️  elapsed: 6m0s
- ✅ status: Active

### Step 4 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 2ms
- **Status Code**: 200

- ✅ migrated cluster present in list (8 total)

### Step 5 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 1ms
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=alibaba+ap-northeast-2+ecs.e-c1m2.xlarge, nodes=2)
- ✅ version: 1.35.7-aliyun.1 (recommended 1.35.7-aliyun.1)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 21.392s

- ✅ kubeconfig obtained (server: https://43.108.104.153:6443)
- ℹ️  auth method: client certificate in kubeconfig
- ✅ API server reachable (v1.35.7-aliyun.1)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 2)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: 43.108.32.92
- ✅ nginx served over the LoadBalancer at http://43.108.32.92/ (attempt 1)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 7 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 13m43.715s
- **Status Code**: 200

- ✅ deleted on attempt 1 (13m43s)

### Step 8 — Residual resource check (Tumblebug)

- **Duration**: 4ms

- ℹ️  VNet mig04-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig04-k8s-sg still exists (known gap)
- ℹ️  SshKey mig04-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for alibaba ap-northeast-2 (source: v1.32.3 → target: v1.35.7-aliyun.1)",
  "targetCloud": {
    "csp": "alibaba",
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
    "connectionName": "alibaba-ap-northeast-2",
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
    "connectionName": "alibaba-ap-northeast-2",
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
      "connectionName": "alibaba-ap-northeast-2",
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
    "connectionName": "alibaba-ap-northeast-2",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "1.35.7-aliyun.1",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": [
      {
        "name": "workers1",
        "imageId": "default",
        "specId": "alibaba+ap-northeast-2+ecs.e-c1m2.xlarge",
        "rootDiskType": "cloud_essd",
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
  "id": "mig04-on-prem-k8s-cluster",
  "uid": "tbv3okbhruku5snohcm7",
  "name": "mig04-on-prem-k8s-cluster",
  "connectionName": "alibaba-ap-northeast-2",
  "connectionConfig": {
    "configName": "alibaba-ap-northeast-2",
    "providerName": "alibaba",
    "driverName": "alibaba-driver-v1.0.so",
    "credentialName": "alibaba",
    "credentialHolder": "admin",
    "regionZoneInfoName": "alibaba-ap-northeast-2",
    "regionZoneInfo": {
      "assignedRegion": "ap-northeast-2",
      "assignedZone": "ap-northeast-2a"
    },
    "regionDetail": {
      "regionId": "ap-northeast-2",
      "regionName": "ap-northeast-2",
      "description": "South Korea (Seoul)",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "zones": [
        "ap-northeast-2a",
        "ap-northeast-2b"
      ]
    },
    "regionRepresentative": true,
    "verified": true
  },
  "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
  "systemMessage": "",
  "label": {
    "sys.connectionName": "alibaba-ap-northeast-2",
    "sys.createdTime": "2026-10-07 15:42:17 +0800 +0800",
    "sys.cspResourceId": "c6e077638479f4f34aebd09db488bac55",
    "sys.cspResourceName": "tbv3okbhruku5snohcm7",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig04-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig04-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbv3okbhruku5snohcm7",
    "sys.version": "1.35.7-aliyun.1"
  },
  "systemLabel": "",
  "version": "1.35.7-aliyun.1",
  "network": {
    "vNetId": "mig04-k8s-vpc",
    "subnetIds": [
      "mig04-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "mig04-k8s-sg"
    ],
    "keyValueList": null
  },
  "k8sNodeGroupList": [
    {
      "id": "workers1",
      "name": "workers1",
      "imageId": "default",
      "specId": "alibaba+ap-northeast-2+ecs.e-c1m2.xlarge",
      "rootDiskType": "cloud_essd",
      "rootDiskSize": 100,
      "sshKeyId": "mig04-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 0,
      "maxNodeSize": 0,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "worker-k8s-for-cs-c6e077638479f4f34aebd09db488bac55",
          "cspResourceId": "i-mj703gvdu16h7jhky1qd"
        },
        {
          "cspResourceName": "worker-k8s-for-cs-c6e077638479f4f34aebd09db488bac55",
          "cspResourceId": "i-mj703gvdu16h7jhky1qe"
        }
      ],
      "keyValueList": [
        {
          "key": "AutoScaling",
          "value": "{eip_bandwidth:0,eip_internet_charge_type:,enable:false,max_instances:0,min_instances:0,type:}"
        },
        {
          "key": "KubernetesConfig",
          "value": "{cms_enabled:false,cpu_policy:none,node_name_mode:nodeip,runtime:containerd,runtime_version:2.3.4,unschedulable:false,user_data:}"
        },
        {
          "key": "Management",
          "value": "{enable:false,upgrade_config:{auto_upgrade:false,max_unavailable:0,surge:0}}"
        },
        {
          "key": "NodeConfig",
          "value": "{}"
        },
        {
          "key": "NodepoolInfo",
          "value": "{created:2026-10-07T15:45:51.807275905+08:00,is_default:false,name:workers1,nodepool_id:npa6c7b2a751134408bae8605e1a54c0c7,region_id:ap-northeast-2,resource_group_id:,type:ess,updated:2026-10-07T15:47:19.653+08:00}"
        },
        {
          "key": "ScalingGroup",
          "value": "{auto_renew:false,auto_renew_period:0,deploymentset_id:,desired_size:2,image_id:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,image_type:AliyunLinux3ContainerOptimized,instance_charge_type:PostPaid,instance_types:[ecs.e-c1m2.xlarge],internet_charge_type:,internet_max_bandwidth_out:0,key_pair:tbaf9fb55i8vvji1ar6g,login_password:,multi_az_policy:BALANCE,period:0,period_unit:,platform:AliyunLinux,private_pool_options:{},ram_policy:KubernetesWorkerRole-f5cc0480-7fab-4de5-8b79-10abe998047f,scaling_group_id:asg-mj76fokzdbwhbk9dipbr,scaling_policy:release,security_group_id:sg-mj7hydw78w20s8gutrve,security_group_ids:[sg-mj7hydw78w20s8gutrve],spot_strategy:NoSpot,system_disk_categories:[cloud_essd],system_disk_category:cloud_essd,system_disk_encrypt_algorithm:,system_disk_encrypted:false,system_disk_kms_key_id:,system_disk_performance_level:,system_disk_size:100,vswitch_ids:[vsw-mj7mkdtynodmq5f1ov0ec]}"
        },
        {
          "key": "Status",
          "value": "{failed_nodes:0,healthy_nodes:2,initial_nodes:2,offline_nodes:0,removing_nodes:0,serving_nodes:2,state:active,total_nodes:2}"
        },
        {
          "key": "TeeConfig",
          "value": "{tee_enable:false}"
        }
      ],
      "cspResourceName": "workers1",
      "cspResourceId": "npa6c7b2a751134408bae8605e1a54c0c7",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "npa6c7b2a751134408bae8605e1a54c0c7"
        },
        "ImageIID": {
          "NameId": "aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd",
          "SystemId": "aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd"
        },
        "VMSpecName": "ecs.e-c1m2.xlarge",
        "RootDiskType": "cloud_essd",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tbaf9fb55i8vvji1ar6g",
          "SystemId": "tbaf9fb55i8vvji1ar6g"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "worker-k8s-for-cs-c6e077638479f4f34aebd09db488bac55",
            "SystemId": "i-mj703gvdu16h7jhky1qd"
          },
          {
            "NameId": "worker-k8s-for-cs-c6e077638479f4f34aebd09db488bac55",
            "SystemId": "i-mj703gvdu16h7jhky1qe"
          }
        ],
        "KeyValueList": [
          {
            "key": "AutoScaling",
            "value": "{eip_bandwidth:0,eip_internet_charge_type:,enable:false,max_instances:0,min_instances:0,type:}"
          },
          {
            "key": "KubernetesConfig",
            "value": "{cms_enabled:false,cpu_policy:none,node_name_mode:nodeip,runtime:containerd,runtime_version:2.3.4,unschedulable:false,user_data:}"
          },
          {
            "key": "Management",
            "value": "{enable:false,upgrade_config:{auto_upgrade:false,max_unavailable:0,surge:0}}"
          },
          {
            "key": "NodeConfig",
            "value": "{}"
          },
          {
            "key": "NodepoolInfo",
            "value": "{created:2026-10-07T15:45:51.807275905+08:00,is_default:false,name:workers1,nodepool_id:npa6c7b2a751134408bae8605e1a54c0c7,region_id:ap-northeast-2,resource_group_id:,type:ess,updated:2026-10-07T15:47:19.653+08:00}"
          },
          {
            "key": "ScalingGroup",
            "value": "{auto_renew:false,auto_renew_period:0,deploymentset_id:,desired_size:2,image_id:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,image_type:AliyunLinux3ContainerOptimized,instance_charge_type:PostPaid,instance_types:[ecs.e-c1m2.xlarge],internet_charge_type:,internet_max_bandwidth_out:0,key_pair:tbaf9fb55i8vvji1ar6g,login_password:,multi_az_policy:BALANCE,period:0,period_unit:,platform:AliyunLinux,private_pool_options:{},ram_policy:KubernetesWorkerRole-f5cc0480-7fab-4de5-8b79-10abe998047f,scaling_group_id:asg-mj76fokzdbwhbk9dipbr,scaling_policy:release,security_group_id:sg-mj7hydw78w20s8gutrve,security_group_ids:[sg-mj7hydw78w20s8gutrve],spot_strategy:NoSpot,system_disk_categories:[cloud_essd],system_disk_category:cloud_essd,system_disk_encrypt_algorithm:,system_disk_encrypted:false,system_disk_kms_key_id:,system_disk_performance_level:,system_disk_size:100,vswitch_ids:[vsw-mj7mkdtynodmq5f1ov0ec]}"
          },
          {
            "key": "Status",
            "value": "{failed_nodes:0,healthy_nodes:2,initial_nodes:2,offline_nodes:0,removing_nodes:0,serving_nodes:2,state:active,total_nodes:2}"
          },
          {
            "key": "TeeConfig",
            "value": "{tee_enable:false}"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://43.108.104.153:6443",
    "kubeconfig": "\napiVersion: v1\nclusters:\n- cluster:\n    server: https://43.108.104.153:6443\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUR2VENDQXFXZ0F3SUJBZ0lJQWh4aU9SSUFIb1F3RFFZSktvWklodmNOQVFFTEJRQXdRREVSTUE4R0ExVUUKQ2hNSWFHRnVaM3BvYjNVeEZqQVVCZ05WQkFzVERXRnNhV0poWW1FZ1kyeHZkV1F4RXpBUkJnTlZCQU1UQ210MQpZbVZ5Ym1WMFpYTXdJQmNOTWpZeE1EQTNNRGN6TnpBd1doZ1BNakExTmpBNU1qa3dOelF5TVRkYU1FQXhFVEFQCkJnTlZCQW9UQ0doaGJtZDZhRzkxTVJZd0ZBWURWUVFMRXcxaGJHbGlZV0poSUdOc2IzVmtNUk13RVFZRFZRUUQKRXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBeExJdwphc0E0bUVlajMrTkhCMG1Pa2hWMnduaWVlTzk5ZVovVHp3eXRjVFU4R055ejJhSHlPOWY2MXNROXV0TkRqcGxXCjBET1ZWbzQwL2RqZlVNd09mOTE5VFFzQk13czc0bk1DRU9NZDJOekExYUJZR2JNb0FhWUNNUSthQUQ1eTN2YmMKcWsxdFVOM2htSEFvRjN2L1FqV0F1QWw3VDNiWUR6UEpBOHdFYmY3YUVtMFMxZC9XOTBHTkZCNnllTFRBNy9nMwpoN3hUSkFXUDdiZHRHeGxpZUhnSFNEQVNXTHgvVmFOR0orQlZVb1ZqMy9za2ZNdzA4Tlhub284djNQZkN3R0VvCmdCbXlISTc4TllZTlRKbFg4Z2VidUhEWTJYYkl1azJsUHdFNGdIb1ZkdGhFdm1OeVd4a21RVnRWWjNLb1M5eEMKRDFDVDRZeWJ5ZFhGS1J5cHFRSURBUUFCbzRHNE1JRzFNQTRHQTFVZER3RUIvd1FFQXdJQ3JEQVBCZ05WSFJNQgpBZjhFQlRBREFRSC9NQjBHQTFVZERnUVdCQlNac0Q5aVlxNk50a21zZTY3Z3M0Q0tVeDFUQ1RBOEJnZ3JCZ0VGCkJRY0JBUVF3TUM0d0xBWUlLd1lCQlFVSE1BR0dJR2gwZEhBNkx5OWpaWEowY3k1aFkzTXVZV3hwZVhWdUxtTnYKYlM5dlkzTndNRFVHQTFVZEh3UXVNQ3d3S3FBb29DYUdKR2gwZEhBNkx5OWpaWEowY3k1aFkzTXVZV3hwZVhWdQpMbU52YlM5eWIyOTBMbU55YkRBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQXNXN0ZGOW53d3BieHgxeCt4QktrCnFxOUw1WlFUUHB4SzZyczh3QTFER1VjTFJSRGE0aHlISGpCQkQ4SDk3Q0tzY0ZvTUNoeFM0SVlFNFQ4UFdBa2EKTDFla1NDZUw0akJEKzQ3TWlEZkR0eGw1bmNBb2xJZWJOeGgvVENZOC83MkJ5aUpOSW8weEVWOEQxTHVZcEVtcApUOEtOT1VXMkhnQVR0enpndXNuSUtnN2NodDhGRFh6ZmFzUE80T0RsNGdicFNDZW1rQTR3bWZTemQ1Sm9BTDhUCk9EWWQ5ZUNyZWN4Q1RNb3NpcWF6RWVqZWtnNDFOVHBjeHJuSnh5UXhtaVJXQXZ3ZHYraDlSKzNUaUhQalBtYmIKSVVxVEF3SklybmlXL2RUMUxMWk0xcU1Ya2EzUnJTV0w2WTBlRHRqVGhSZFZZQzl5bXpnL0cyMUtmTloyTW8vUQpRUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: kubernetes\ncontexts:\n- context:\n    cluster: kubernetes\n    user: \"213652363353305376\"\n  name: 213652363353305376-c6e077638479f4f34aebd09db488bac55\ncurrent-context: 213652363353305376-c6e077638479f4f34aebd09db488bac55\nkind: Config\npreferences: {}\nusers:\n- name: \"213652363353305376\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUQyVENDQXNHZ0F3SUJBZ0lJQWh4aWlvb0FIb1F3RFFZSktvWklodmNOQVFFTEJRQXdRREVSTUE4R0ExVUUKQ2hNSWFHRnVaM3BvYjNVeEZqQVVCZ05WQkFzVERXRnNhV0poWW1FZ1kyeHZkV1F4RXpBUkJnTlZCQU1UQ210MQpZbVZ5Ym1WMFpYTXdIaGNOTWpZeE1EQTNNRGMwTURBd1doY05Namt4TURBMk1EYzBOVFEyV2pCS01SVXdFd1lEClZRUUtFd3h6ZVhOMFpXMDZkWE5sY25NeENUQUhCZ05WQkFzVEFERW1NQ1FHQTFVRUF4TWRNakV6TmpVeU16WXoKTXpVek16QTFNemMyTFRFM09URXpOVGt4TkRZd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFSwpBb0lCQVFEcmszQ29QL3lJdmRYSzAxSmw1c1ZnUE1VME5WTkxWalhjWVc5ZTh3N21vUFVQOGExS0dFdHUwUGsyCmpHZFhWRDJyRGFWRlVENWRGMWw0Q291SmhpZmszaGVFMFpPVG5ucVFMMFlFdExYem5LSTVyNVBOcFk4c1FrQkUKWGJFNlBaeEdFbWtVRVg0RE9MTEZSc3BHTFFqUEd3bml5cmh1dDA0cXdkTGJGZVZKUFJFWDIwd3Y1M2dHVHIvNApuRmZzNTBIeWNwWHIwR1VjNW5KL203L0lvU29Db3pRK1lrTWVFWUJJUzBSMmZMU0pzRFA3UjFTcTFYMFRZTGpjCkVDVUJneS9ZMFZ4VDQ4S25mbHhNdm9PVFhDWEZGUVcxbGRpd2NnWnY0bG5CUDkrdlJFbHcwMDI4aUJUT1Q2dFcKd3pSd0RoS3FYYmpWTXFkcjRuSXAvOWNEZUV0L0FnTUJBQUdqZ2N3d2dja3dEZ1lEVlIwUEFRSC9CQVFEQWdlQQpNQk1HQTFVZEpRUU1NQW9HQ0NzR0FRVUZCd01DTUF3R0ExVWRFd0VCL3dRQ01BQXdId1lEVlIwakJCZ3dGb0FVCm1iQS9ZbUt1amJaSnJIdXU0TE9BaWxNZFV3a3dQQVlJS3dZQkJRVUhBUUVFTURBdU1Dd0dDQ3NHQVFVRkJ6QUIKaGlCb2RIUndPaTh2WTJWeWRITXVZV056TG1Gc2FYbDFiaTVqYjIwdmIyTnpjREExQmdOVkhSOEVMakFzTUNxZwpLS0FtaGlSb2RIUndPaTh2WTJWeWRITXVZV056TG1Gc2FYbDFiaTVqYjIwdmNtOXZkQzVqY213d0RRWUpLb1pJCmh2Y05BUUVMQlFBRGdnRUJBR2NMc2k0T09jajMxczA2ajZ4bEg4T2hoMEpzSjFRcDlkdENRbU5TU3QvRmF6TUYKZm1nV2FIblk2Q3lpMTlVMmFaa1hlZzJmUmxDVG1TMG53MitYNkNSZ21HNkV0RmdhZVF3dDdTVnozMEFZbnpHWApzOS9Xd25EZDVzakJETXRuc3VHMlZwZUszSVMzQ0s3Ym03Wmo3VG5aV3ArVTZ5UzV4NjNJSUNYMU4zNGd3d1FUCkVBdU16aWN4eWVtNjEwdzRQbkhUc0ViZkNZSi9IL25zbFFHT0VzNkh2MjVHdTJEVzBySVppMWFZOGJsT3NNK0YKZUxVNzhici9JWEJ4SHIxclZ5QTNOa29oT0xCa3lVcExmS2g1UzlJbUo0STlwUGNhWTFWRVduWUNrSXA1NXZuTgpqZlU4bDFlQ2N0VXJPKzg5ZkFPZHlnbFhPUHlXUThHYnVpZlZzbFE9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBNjVOd3FELzhpTDNWeXROU1plYkZZRHpGTkRWVFMxWTEzR0Z2WHZNTzVxRDFEL0d0ClNoaExidEQ1Tm94blYxUTlxdzJsUlZBK1hSZFplQXFMaVlZbjVONFhoTkdUazU1NmtDOUdCTFMxODV5aU9hK1QKemFXUExFSkFSRjJ4T2oyY1JoSnBGQkYrQXppeXhVYktSaTBJenhzSjRzcTRicmRPS3NIUzJ4WGxTVDBSRjl0TQpMK2Q0Qms2LytKeFg3T2RCOG5LVjY5QmxIT1p5ZjV1L3lLRXFBcU0wUG1KREhoR0FTRXRFZG55MGliQXorMGRVCnF0VjlFMkM0M0JBbEFZTXYyTkZjVStQQ3AzNWNUTDZEazF3bHhSVUZ0WlhZc0hJR2IrSlp3VC9mcjBSSmNOTk4KdklnVXprK3JWc00wY0E0U3FsMjQxVEtuYStKeUtmL1hBM2hMZndJREFRQUJBb0lCQVFDRVBsS0J3Tm5ORVhSUQoxZXh0ayt5OUo2QTB2TEt6bWdQR1lWUVo0eXc4UDZNU2ZrTWZVeUFWWjI1Zk50WlVhYy9zbEU1dzZLajVyVC9tCmFWVUhzSzM0aHN5QkhQMVJZeWUraFNzelBSYmZXTkNndlhXbGZna3ZlVW1HbDJvRUhjMzZjQjlZeXJFSXNlMTUKWFZIenJ4aEEyeGJqbjVXRllaV1ozeEMxT0Vkc29xMlhqalRlZlVDdWFEaE9KS0d1WUxZTjNlY1RNbEJPVVhKTgoweEtpOG5hY1prc3NOUGozYjZYWC9oZ2srcW1WcDhWaEtaU3RmN1NiZ28xQ3VkT0tqem1uTTBwZ1c2K29rWTU3CmRrdkp2eGtBOTFFWDUzTm9SSkcyWVRUS1NkYkpOeFcxUzNKRFI1Z3crbHJYR3Zod0g5eVE0UEJEM0c3UjY2TlAKQjVsQWMzVlJBb0dCQVBoVjRkallHTnhlQ1MrZkdnS1pJRHVYNEJIYzRqQ1N4THBxdjZBbVpmU0dBNU1mNmVkSApEQnhVTGZiVGFMMDNpbnMvY1BYY3pianB2MnhTdXprcCtwQXh6UEtqTjMwbVlSTXBBWEUrMUUyajVWZTRzQjJCCnUwUFdkSEpZQTkwM2JaMVFJU2dxVFY3VGxhdnArc1RLV2FoYThXQ1dRWHVpbVlBWWlJNTBPekN0QW9HQkFQTFkKdm1ZRkRvaW1PRDJSNFRFTmtIZkZtMS9UUHRmQWI5Z0RBa2ZCZXhjWnNrdDRyWkdLeDhraG93ZHlidDNyby9mRgpQUXkrWXBuRGRXWmlrTmhpTG9LK2hDVDRwbnRqZG80STR6a0w2QTlaNVNRRm13UTczWTkyV2o3SXo0L3djVkxaClFMYTdpUlNHdld3c3ZUSExGVkEyQjBsa0crczFzWjZRQ0gxWWY3WmJBb0dBVnZXVDZWLzZqS2d0SlV6Y1NjNmYKRjk4ZDZvTmpmVWpYdE1PT3FLRHBrTStnenRNZEVBeXo4L09TU04xTXp2MTA4NldLZzcyM0dDcGFDcStKdWdHMgpLT29YYjF1eUlaUGY2RnF5azVwQnM1SlJ5Lzd6Uk5IWjVtNWhSbTBGcFdBMGRTTExDWXFPbjBjT0lTNEV6d1pnCmtQQ1BsNWZtQkVveTRFVTNCRk0xS3QwQ2dZRUFueFdGaE9lRElkbGh0bE8rSlpneEw0VHZ4OUptdEllSHZRWWIKbEU0WENJYjQyWi93ZEF0cDNVUi91LzBteTVIMkUvWE5qRytid0FiZS9YZE1VN1BkckNDS0NINHE4V0d5NUZERwpLVFMzakhiak9MbkRWdjZ3b2E1eVovYThvaHBzNGswWHE1MG1xNStvcnhpUHgvSzF2NW5sSkJyRWYwenBVbW9nClpaeDM3VWtDZ1lFQThGWnRDMUE0eHdSeGdjWGRSTUNCYmdkM00yWUVibFNhV3lUb2wyVFRKYkxCbHRzVGpvUmYKNjZDdUFSa3BPaEh2ZWREY2pYOGcyWGh0MTZIaWtPTkMycTZZbkUrVGQzM085Vkh2QVBsK1NLaGlJTnVWWGdwdgpkY3VDZGpDMi9IYlQrUU5UbUNIWGloUllLeThZZ3U3eVE1c2E0aDM3ZG9GbCtHNlowdFdBZmxzPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\n"
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-10-07T15:42:17+08:00",
  "keyValueList": [
    {
      "key": "ClusterId",
      "value": "c6e077638479f4f34aebd09db488bac55"
    },
    {
      "key": "ClusterSpec",
      "value": "ack.pro.small"
    },
    {
      "key": "ClusterType",
      "value": "ManagedKubernetes"
    },
    {
      "key": "Created",
      "value": "2026-10-07T15:42:17+08:00"
    },
    {
      "key": "CurrentVersion",
      "value": "1.35.7-aliyun.1"
    },
    {
      "key": "DeletionProtection",
      "value": "false"
    },
    {
      "key": "ExternalLoadbalancerId",
      "value": "lb-mj7d0gxerspz8pjrrif6h"
    },
    {
      "key": "InitVersion",
      "value": "1.35.7-aliyun.1"
    },
    {
      "key": "MaintenanceWindow",
      "value": "{enable:false,weekly_period:}"
    },
    {
      "key": "MasterUrl",
      "value": "{\\api_server_endpoint\\:\\https://43.108.104.153:6443\\,\\intranet_api_server_endpoint\\:\\https://10.0.1.70:6443\\}"
    },
    {
      "key": "MetaData",
      "value": "{\\Addons\\:[{\\name\\:\\storage-operator\\,\\version\\:\\v1.35.4\\},{\\name\\:\\csi-provisioner\\,\\version\\:\\v1.37.2\\},{\\name\\:\\gateway-api\\,\\version\\:\\1.3.0\\},{\\name\\:\\cloud-controller-manager\\,\\version\\:\\v2.15.0\\},{\\name\\:\\alicloud-monitor-controller\\,\\version\\:\\v1.8.13\\},{\\name\\:\\metrics-server\\,\\version\\:\\v0.3.16-ab96b23-aliyun\\},{\\name\\:\\metrics-aggregator\\,\\version\\:\\v1.9.6\\},{\\name\\:\\ack-scheduler\\,\\version\\:\\v1.35.2-apsara.6.12.3.a6eed53c\\},{\\name\\:\\pod-readinessgate-webhook\\,\\version\\:\\v1.0.0\\},{\\name\\:\\coredns\\,\\version\\:\\v1.13.2.3\\},{\\name\\:\\kube-flannel-ds\\,\\version\\:\\v0.28.0.6\\},{\\name\\:\\ack-nvidia-device-plugin\\,\\version\\:\\0.8.1\\},{\\name\\:\\managed-kube-proxy\\,\\version\\:\\v1.35.7-aliyun.1\\},{\\name\\:\\ack-ram-authenticator\\,\\version\\:\\0.5.1\\},{\\name\\:\\csi-plugin\\,\\version\\:\\v1.37.2\\},{\\name\\:\\cnfs-controller\\,\\version\\:\\v1.2.5\\,\\config\\:\\{\\\\\\CreateDefaultCNFS\\\\\\:false}\\},{\\name\\:\\kube-apiserver\\,\\version\\:\\v1.35.7-aliyun.1\\},{\\name\\:\\kube-controller-manager\\,\\version\\:\\v1.35.7-aliyun.1\\}],\\AuditProjectName\\:\\\\,\\Capabilities\\:{\\AnyAZ\\:true,\\CSI\\:true,\\CpuPolicy\\:true,\\DeploymentSet\\:true,\\DisableEncryption\\:true,\\EncryptionKMSKeyId\\:\\\\,\\EnterpriseSecurityGroup\\:true,\\HpcCluster\\:true,\\IntelSGX\\:false,\\Knative\\:true,\\Network\\:\\Flannel\\,\\NgwPayByLcu\\:true,\\NodeCIDRMask\\:\\25\\,\\NodeNameMode\\:true,\\ProxyMode\\:\\\\,\\PublicSLB\\:true,\\RamRoleType\\:\\restricted\\,\\SLSProjectName\\:true,\\SandboxRuntime\\:true,\\SnapshotPolicy\\:true,\\Taint\\:true,\\TerwayEniip\\:true,\\UserData\\:true},\\CloudMonitorVersion\\:\\\\,\\ClusterDomain\\:\\\\,\\ControlPlaneLogConfig\\:{\\components\\:null},\\DockerVersion\\:\\\\,\\EtcdVersion\\:\\\\,\\ExtraCertSAN\\:null,\\FreeTier\\:false,\\HasSandboxRuntime\\:false,\\IPStack\\:\\ipv4\\,\\ImageType\\:\\AliyunLinux3ContainerOptimized\\,\\KubernetesVersion\\:\\1.35.7-aliyun.1\\,\\MultiAZ\\:false,\\NameMode\\:\\\\,\\NextVersion\\:\\\\,\\OSType\\:\\Linux\\,\\Platform\\:\\AliyunLinux\\,\\PodVswitchId\\:\\\\,\\Provider\\:\\\\,\\RRSAConfig\\:{\\enabled\\:false},\\RamRoleType\\:\\\\,\\ResourceGroupId\\:\\rg-acfnvekhilw5kmy\\,\\Runtime\\:\\containerd\\,\\RuntimeVersion\\:\\2.3.4\\,\\ServiceCIDR\\:\\172.20.0.0/16\\,\\SubClass\\:\\default\\,\\SupportPlatforms\\:[\\CentOS\\,\\AliyunLinux\\,\\Windows\\,\\WindowsCore\\],\\Timezone\\:\\\\,\\VSwitchIds\\:[\\vsw-mj7mkdtynodmq5f1ov0ec\\],\\VersionSpec\\:null,\\VpcCidr\\:\\10.0.0.0/22\\,\\ack-nvidia-device-pluginVersion\\:\\0.8.1\\,\\ack-ram-authenticatorVersion\\:\\0.5.1\\,\\ack-schedulerVersion\\:\\v1.35.2-apsara.6.12.3.a6eed53c\\,\\alicloud-monitor-controllerVersion\\:\\v1.8.13\\,\\cloud-controller-managerVersion\\:\\v2.15.0\\,\\cnfs-controllerVersion\\:\\v1.2.5\\,\\corednsVersion\\:\\v1.13.2.3\\,\\csi-pluginVersion\\:\\v1.37.2\\,\\csi-provisionerVersion\\:\\v1.37.2\\,\\gateway-apiVersion\\:\\1.3.0\\,\\kube-apiserverVersion\\:\\v1.35.7-aliyun.1\\,\\kube-controller-managerVersion\\:\\v1.35.7-aliyun.1\\,\\kube-flannel-dsVersion\\:\\v0.28.0.6\\,\\metrics-aggregatorVersion\\:\\v1.9.6\\,\\metrics-serverVersion\\:\\v0.3.16-ab96b23-aliyun\\,\\pod-readinessgate-webhookVersion\\:\\v1.0.0\\,\\storage-operatorVersion\\:\\v1.35.4\\}"
    },
    {
      "key": "Name",
      "value": "tbv3okbhruku5snohcm7"
    },
    {
      "key": "NetworkMode",
      "value": "vpc"
    },
    {
      "key": "Parameters",
      "value": "{ALIYUN::AccountId:5469257408566579,ALIYUN::NoValue:None,ALIYUN::Region:ap-northeast-2,ALIYUN::ResourceGroupId:rg-acfnvekhilw5kmy,ALIYUN::StackId:f5cc0480-7fab-4de5-8b79-10abe998047f,ALIYUN::StackName:k8s-for-cs-c6e077638479f4f34aebd09db488bac55,ALIYUN::TenantId:5469257408566579,AdjustmentType:TotalCapacity,BetaVersion:,CloudMonitorFlags:False,CloudMonitorVersion:1.3.7,ClusterDns:172.20.0.10,ClusterId:c6e077638479f4f34aebd09db488bac55,ContainerCIDR:172.19.0.0/16,CustomK8sWorkerRole:,DisableAddons:True,DisableAutoCreateK8sWorkerRole:False,DisableAutoCreateK8sWorkerRolePolicy:True,DockerVersion:17.06.2-ce-3,ESSDeletionProtection:True,Eip:True,EipAddress:,EtcdVersion:v3.1.11,ExecuteVersion:63084271,HealthCheckType:NONE,IPStack:ipv4,ImageId:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,K8sWorkerPolicyDocument:{\\Version\\: \\1\\, \\Statement\\: [{\\Action\\: [\\ecs:DescribeInstanceAttribute\\, \\ecs:DescribeInstances\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}, {\\Action\\: [\\log:GetProject\\, \\log:GetLogStore\\, \\log:GetConfig\\, \\log:GetMachineGroup\\, \\log:GetAppliedMachineGroups\\, \\log:GetAppliedConfigs\\, \\log:GetIndex\\, \\log:GetSavedSearch\\, \\log:GetDashboard\\, \\log:GetJob\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}, {\\Action\\: [\\cr:GetAuthorizationToken\\, \\cr:ListInstanceEndpoint\\, \\cr:PullRepository\\, \\cr:GetInstanceVpcEndpoint\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}]},KeyPair:,KubernetesVersion:1.35.7-aliyun.1,MasterSLBPrivateIP:10.0.1.70,NatGateway:False,NatGatewayId:ngw-mj7r0n0opzyvi2m4q0v9v,NatGatewayType:Enhanced,NatGatewayVswitchId:,Network:Flannel,NodeNameMode:nodeip,NumOfNodes:0,OSType:Linux,Password:******,PodVswitchIds:[],ProtectedInstances:,ProxyMode:iptables,RemoveInstanceIds:,ResourceGroupId:rg-acfnvekhilw5kmy,SNatEntry:True,ScaleOutToken:None,SecurityGroupId:sg-mj7hydw78w20s8gutrve,ServiceCIDR:172.20.0.0/16,SetUpArgs:--addon-names kube-flannel-ds,csi-plugin,csi-provisioner --labels CB-SPIDER:PMKS:CLUSTER=owned --node-cidr-mask 25 --registry-url registry-ap-northeast-2-vpc.ack.aliyuncs.com --runtime containerd --runtime-version 2.3.4,SnatTableId:stb-mj7dos9owfy0vkja6ibfb,Tags:[{\\Key\\: \\CB-SPIDER:PMKS:CLUSTER\\, \\Value\\: \\owned\\}, {\\Key\\: \\ack.aliyun.com\\, \\Value\\: \\c6e077638479f4f34aebd09db488bac55\\}],UserData:,VpcCidrWithSecondaryCidrs:[\\10.0.0.0/22\\],VpcId:vpc-mj7s33x3t7vhkkospgg4f,WorkerAutoRenew:False,WorkerAutoRenewPeriod:1,WorkerDataDisk:False,WorkerDataDisks:[],WorkerDeletionProtection:True,WorkerDeploymentSetId:,WorkerHpcClusterId:,WorkerImageId:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,WorkerInstanceChargeType:PostPaid,WorkerInstanceTypes:ecs.n4.large,WorkerKeyPair:,WorkerLoginPassword:******,WorkerPeriod:3,WorkerPeriodUnit:Month,WorkerSnapshotPolicyId:******,WorkerSystemDiskCategory:cloud_ssd,WorkerSystemDiskPerformanceLevel:null,WorkerSystemDiskSize:40,WorkerVSwitchIds:vsw-mj7mkdtynodmq5f1ov0ec,ZoneId:}"
    },
    {
      "key": "Profile",
      "value": "Default"
    },
    {
      "key": "RegionId",
      "value": "ap-northeast-2"
    },
    {
      "key": "ResourceGroupId",
      "value": "rg-acfnvekhilw5kmy"
    },
    {
      "key": "SecurityGroupId",
      "value": "sg-mj7hydw78w20s8gutrve"
    },
    {
      "key": "Size",
      "value": "0"
    },
    {
      "key": "State",
      "value": "running"
    },
    {
      "key": "SubnetCidr",
      "value": "172.19.0.0/16"
    },
    {
      "key": "Tags",
      "value": "{key:CB-SPIDER:PMKS:CLUSTER,value:owned}; {key:ack.aliyun.com,value:c6e077638479f4f34aebd09db488bac55}"
    },
    {
      "key": "Updated",
      "value": "2026-10-07T15:47:00+08:00"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7s33x3t7vhkkospgg4f"
    },
    {
      "key": "VswitchId",
      "value": "vsw-mj7mkdtynodmq5f1ov0ec"
    },
    {
      "key": "WorkerRamRoleName",
      "value": "KubernetesWorkerRole-f5cc0480-7fab-4de5-8b79-10abe998047f"
    },
    {
      "key": "ZoneId",
      "value": "ap-northeast-2a"
    }
  ],
  "cspResourceName": "tbv3okbhruku5snohcm7",
  "cspResourceId": "c6e077638479f4f34aebd09db488bac55",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbv3okbhruku5snohcm7",
      "SystemId": "c6e077638479f4f34aebd09db488bac55"
    },
    "Version": "1.35.7-aliyun.1",
    "Network": {
      "VpcIID": {
        "NameId": "tb8jt8phjo7jc3imnuf2",
        "SystemId": "vpc-mj7s33x3t7vhkkospgg4f"
      },
      "SubnetIIDs": [
        {
          "NameId": "tbhk4ilncoianpemtie9",
          "SystemId": "vsw-mj7mkdtynodmq5f1ov0ec"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "tbbcqnl7o1q6aj35661t",
          "SystemId": "sg-mj7hydw78w20s8gutrve"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "npa6c7b2a751134408bae8605e1a54c0c7"
        },
        "ImageIID": {
          "NameId": "aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd",
          "SystemId": "aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd"
        },
        "VMSpecName": "ecs.e-c1m2.xlarge",
        "RootDiskType": "cloud_essd",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tbaf9fb55i8vvji1ar6g",
          "SystemId": "tbaf9fb55i8vvji1ar6g"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "worker-k8s-for-cs-c6e077638479f4f34aebd09db488bac55",
            "SystemId": "i-mj703gvdu16h7jhky1qd"
          },
          {
            "NameId": "worker-k8s-for-cs-c6e077638479f4f34aebd09db488bac55",
            "SystemId": "i-mj703gvdu16h7jhky1qe"
          }
        ],
        "KeyValueList": [
          {
            "key": "AutoScaling",
            "value": "{eip_bandwidth:0,eip_internet_charge_type:,enable:false,max_instances:0,min_instances:0,type:}"
          },
          {
            "key": "KubernetesConfig",
            "value": "{cms_enabled:false,cpu_policy:none,node_name_mode:nodeip,runtime:containerd,runtime_version:2.3.4,unschedulable:false,user_data:}"
          },
          {
            "key": "Management",
            "value": "{enable:false,upgrade_config:{auto_upgrade:false,max_unavailable:0,surge:0}}"
          },
          {
            "key": "NodeConfig",
            "value": "{}"
          },
          {
            "key": "NodepoolInfo",
            "value": "{created:2026-10-07T15:45:51.807275905+08:00,is_default:false,name:workers1,nodepool_id:npa6c7b2a751134408bae8605e1a54c0c7,region_id:ap-northeast-2,resource_group_id:,type:ess,updated:2026-10-07T15:47:19.653+08:00}"
          },
          {
            "key": "ScalingGroup",
            "value": "{auto_renew:false,auto_renew_period:0,deploymentset_id:,desired_size:2,image_id:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,image_type:AliyunLinux3ContainerOptimized,instance_charge_type:PostPaid,instance_types:[ecs.e-c1m2.xlarge],internet_charge_type:,internet_max_bandwidth_out:0,key_pair:tbaf9fb55i8vvji1ar6g,login_password:,multi_az_policy:BALANCE,period:0,period_unit:,platform:AliyunLinux,private_pool_options:{},ram_policy:KubernetesWorkerRole-f5cc0480-7fab-4de5-8b79-10abe998047f,scaling_group_id:asg-mj76fokzdbwhbk9dipbr,scaling_policy:release,security_group_id:sg-mj7hydw78w20s8gutrve,security_group_ids:[sg-mj7hydw78w20s8gutrve],spot_strategy:NoSpot,system_disk_categories:[cloud_essd],system_disk_category:cloud_essd,system_disk_encrypt_algorithm:,system_disk_encrypted:false,system_disk_kms_key_id:,system_disk_performance_level:,system_disk_size:100,vswitch_ids:[vsw-mj7mkdtynodmq5f1ov0ec]}"
          },
          {
            "key": "Status",
            "value": "{failed_nodes:0,healthy_nodes:2,initial_nodes:2,offline_nodes:0,removing_nodes:0,serving_nodes:2,state:active,total_nodes:2}"
          },
          {
            "key": "TeeConfig",
            "value": "{tee_enable:false}"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://43.108.104.153:6443",
      "Kubeconfig": "\napiVersion: v1\nclusters:\n- cluster:\n    server: https://43.108.104.153:6443\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUR2VENDQXFXZ0F3SUJBZ0lJQWh4aU9SSUFIb1F3RFFZSktvWklodmNOQVFFTEJRQXdRREVSTUE4R0ExVUUKQ2hNSWFHRnVaM3BvYjNVeEZqQVVCZ05WQkFzVERXRnNhV0poWW1FZ1kyeHZkV1F4RXpBUkJnTlZCQU1UQ210MQpZbVZ5Ym1WMFpYTXdJQmNOTWpZeE1EQTNNRGN6TnpBd1doZ1BNakExTmpBNU1qa3dOelF5TVRkYU1FQXhFVEFQCkJnTlZCQW9UQ0doaGJtZDZhRzkxTVJZd0ZBWURWUVFMRXcxaGJHbGlZV0poSUdOc2IzVmtNUk13RVFZRFZRUUQKRXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBeExJdwphc0E0bUVlajMrTkhCMG1Pa2hWMnduaWVlTzk5ZVovVHp3eXRjVFU4R055ejJhSHlPOWY2MXNROXV0TkRqcGxXCjBET1ZWbzQwL2RqZlVNd09mOTE5VFFzQk13czc0bk1DRU9NZDJOekExYUJZR2JNb0FhWUNNUSthQUQ1eTN2YmMKcWsxdFVOM2htSEFvRjN2L1FqV0F1QWw3VDNiWUR6UEpBOHdFYmY3YUVtMFMxZC9XOTBHTkZCNnllTFRBNy9nMwpoN3hUSkFXUDdiZHRHeGxpZUhnSFNEQVNXTHgvVmFOR0orQlZVb1ZqMy9za2ZNdzA4Tlhub284djNQZkN3R0VvCmdCbXlISTc4TllZTlRKbFg4Z2VidUhEWTJYYkl1azJsUHdFNGdIb1ZkdGhFdm1OeVd4a21RVnRWWjNLb1M5eEMKRDFDVDRZeWJ5ZFhGS1J5cHFRSURBUUFCbzRHNE1JRzFNQTRHQTFVZER3RUIvd1FFQXdJQ3JEQVBCZ05WSFJNQgpBZjhFQlRBREFRSC9NQjBHQTFVZERnUVdCQlNac0Q5aVlxNk50a21zZTY3Z3M0Q0tVeDFUQ1RBOEJnZ3JCZ0VGCkJRY0JBUVF3TUM0d0xBWUlLd1lCQlFVSE1BR0dJR2gwZEhBNkx5OWpaWEowY3k1aFkzTXVZV3hwZVhWdUxtTnYKYlM5dlkzTndNRFVHQTFVZEh3UXVNQ3d3S3FBb29DYUdKR2gwZEhBNkx5OWpaWEowY3k1aFkzTXVZV3hwZVhWdQpMbU52YlM5eWIyOTBMbU55YkRBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQXNXN0ZGOW53d3BieHgxeCt4QktrCnFxOUw1WlFUUHB4SzZyczh3QTFER1VjTFJSRGE0aHlISGpCQkQ4SDk3Q0tzY0ZvTUNoeFM0SVlFNFQ4UFdBa2EKTDFla1NDZUw0akJEKzQ3TWlEZkR0eGw1bmNBb2xJZWJOeGgvVENZOC83MkJ5aUpOSW8weEVWOEQxTHVZcEVtcApUOEtOT1VXMkhnQVR0enpndXNuSUtnN2NodDhGRFh6ZmFzUE80T0RsNGdicFNDZW1rQTR3bWZTemQ1Sm9BTDhUCk9EWWQ5ZUNyZWN4Q1RNb3NpcWF6RWVqZWtnNDFOVHBjeHJuSnh5UXhtaVJXQXZ3ZHYraDlSKzNUaUhQalBtYmIKSVVxVEF3SklybmlXL2RUMUxMWk0xcU1Ya2EzUnJTV0w2WTBlRHRqVGhSZFZZQzl5bXpnL0cyMUtmTloyTW8vUQpRUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: kubernetes\ncontexts:\n- context:\n    cluster: kubernetes\n    user: \"213652363353305376\"\n  name: 213652363353305376-c6e077638479f4f34aebd09db488bac55\ncurrent-context: 213652363353305376-c6e077638479f4f34aebd09db488bac55\nkind: Config\npreferences: {}\nusers:\n- name: \"213652363353305376\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUQyVENDQXNHZ0F3SUJBZ0lJQWh4aWlvb0FIb1F3RFFZSktvWklodmNOQVFFTEJRQXdRREVSTUE4R0ExVUUKQ2hNSWFHRnVaM3BvYjNVeEZqQVVCZ05WQkFzVERXRnNhV0poWW1FZ1kyeHZkV1F4RXpBUkJnTlZCQU1UQ210MQpZbVZ5Ym1WMFpYTXdIaGNOTWpZeE1EQTNNRGMwTURBd1doY05Namt4TURBMk1EYzBOVFEyV2pCS01SVXdFd1lEClZRUUtFd3h6ZVhOMFpXMDZkWE5sY25NeENUQUhCZ05WQkFzVEFERW1NQ1FHQTFVRUF4TWRNakV6TmpVeU16WXoKTXpVek16QTFNemMyTFRFM09URXpOVGt4TkRZd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFSwpBb0lCQVFEcmszQ29QL3lJdmRYSzAxSmw1c1ZnUE1VME5WTkxWalhjWVc5ZTh3N21vUFVQOGExS0dFdHUwUGsyCmpHZFhWRDJyRGFWRlVENWRGMWw0Q291SmhpZmszaGVFMFpPVG5ucVFMMFlFdExYem5LSTVyNVBOcFk4c1FrQkUKWGJFNlBaeEdFbWtVRVg0RE9MTEZSc3BHTFFqUEd3bml5cmh1dDA0cXdkTGJGZVZKUFJFWDIwd3Y1M2dHVHIvNApuRmZzNTBIeWNwWHIwR1VjNW5KL203L0lvU29Db3pRK1lrTWVFWUJJUzBSMmZMU0pzRFA3UjFTcTFYMFRZTGpjCkVDVUJneS9ZMFZ4VDQ4S25mbHhNdm9PVFhDWEZGUVcxbGRpd2NnWnY0bG5CUDkrdlJFbHcwMDI4aUJUT1Q2dFcKd3pSd0RoS3FYYmpWTXFkcjRuSXAvOWNEZUV0L0FnTUJBQUdqZ2N3d2dja3dEZ1lEVlIwUEFRSC9CQVFEQWdlQQpNQk1HQTFVZEpRUU1NQW9HQ0NzR0FRVUZCd01DTUF3R0ExVWRFd0VCL3dRQ01BQXdId1lEVlIwakJCZ3dGb0FVCm1iQS9ZbUt1amJaSnJIdXU0TE9BaWxNZFV3a3dQQVlJS3dZQkJRVUhBUUVFTURBdU1Dd0dDQ3NHQVFVRkJ6QUIKaGlCb2RIUndPaTh2WTJWeWRITXVZV056TG1Gc2FYbDFiaTVqYjIwdmIyTnpjREExQmdOVkhSOEVMakFzTUNxZwpLS0FtaGlSb2RIUndPaTh2WTJWeWRITXVZV056TG1Gc2FYbDFiaTVqYjIwdmNtOXZkQzVqY213d0RRWUpLb1pJCmh2Y05BUUVMQlFBRGdnRUJBR2NMc2k0T09jajMxczA2ajZ4bEg4T2hoMEpzSjFRcDlkdENRbU5TU3QvRmF6TUYKZm1nV2FIblk2Q3lpMTlVMmFaa1hlZzJmUmxDVG1TMG53MitYNkNSZ21HNkV0RmdhZVF3dDdTVnozMEFZbnpHWApzOS9Xd25EZDVzakJETXRuc3VHMlZwZUszSVMzQ0s3Ym03Wmo3VG5aV3ArVTZ5UzV4NjNJSUNYMU4zNGd3d1FUCkVBdU16aWN4eWVtNjEwdzRQbkhUc0ViZkNZSi9IL25zbFFHT0VzNkh2MjVHdTJEVzBySVppMWFZOGJsT3NNK0YKZUxVNzhici9JWEJ4SHIxclZ5QTNOa29oT0xCa3lVcExmS2g1UzlJbUo0STlwUGNhWTFWRVduWUNrSXA1NXZuTgpqZlU4bDFlQ2N0VXJPKzg5ZkFPZHlnbFhPUHlXUThHYnVpZlZzbFE9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBNjVOd3FELzhpTDNWeXROU1plYkZZRHpGTkRWVFMxWTEzR0Z2WHZNTzVxRDFEL0d0ClNoaExidEQ1Tm94blYxUTlxdzJsUlZBK1hSZFplQXFMaVlZbjVONFhoTkdUazU1NmtDOUdCTFMxODV5aU9hK1QKemFXUExFSkFSRjJ4T2oyY1JoSnBGQkYrQXppeXhVYktSaTBJenhzSjRzcTRicmRPS3NIUzJ4WGxTVDBSRjl0TQpMK2Q0Qms2LytKeFg3T2RCOG5LVjY5QmxIT1p5ZjV1L3lLRXFBcU0wUG1KREhoR0FTRXRFZG55MGliQXorMGRVCnF0VjlFMkM0M0JBbEFZTXYyTkZjVStQQ3AzNWNUTDZEazF3bHhSVUZ0WlhZc0hJR2IrSlp3VC9mcjBSSmNOTk4KdklnVXprK3JWc00wY0E0U3FsMjQxVEtuYStKeUtmL1hBM2hMZndJREFRQUJBb0lCQVFDRVBsS0J3Tm5ORVhSUQoxZXh0ayt5OUo2QTB2TEt6bWdQR1lWUVo0eXc4UDZNU2ZrTWZVeUFWWjI1Zk50WlVhYy9zbEU1dzZLajVyVC9tCmFWVUhzSzM0aHN5QkhQMVJZeWUraFNzelBSYmZXTkNndlhXbGZna3ZlVW1HbDJvRUhjMzZjQjlZeXJFSXNlMTUKWFZIenJ4aEEyeGJqbjVXRllaV1ozeEMxT0Vkc29xMlhqalRlZlVDdWFEaE9KS0d1WUxZTjNlY1RNbEJPVVhKTgoweEtpOG5hY1prc3NOUGozYjZYWC9oZ2srcW1WcDhWaEtaU3RmN1NiZ28xQ3VkT0tqem1uTTBwZ1c2K29rWTU3CmRrdkp2eGtBOTFFWDUzTm9SSkcyWVRUS1NkYkpOeFcxUzNKRFI1Z3crbHJYR3Zod0g5eVE0UEJEM0c3UjY2TlAKQjVsQWMzVlJBb0dCQVBoVjRkallHTnhlQ1MrZkdnS1pJRHVYNEJIYzRqQ1N4THBxdjZBbVpmU0dBNU1mNmVkSApEQnhVTGZiVGFMMDNpbnMvY1BYY3pianB2MnhTdXprcCtwQXh6UEtqTjMwbVlSTXBBWEUrMUUyajVWZTRzQjJCCnUwUFdkSEpZQTkwM2JaMVFJU2dxVFY3VGxhdnArc1RLV2FoYThXQ1dRWHVpbVlBWWlJNTBPekN0QW9HQkFQTFkKdm1ZRkRvaW1PRDJSNFRFTmtIZkZtMS9UUHRmQWI5Z0RBa2ZCZXhjWnNrdDRyWkdLeDhraG93ZHlidDNyby9mRgpQUXkrWXBuRGRXWmlrTmhpTG9LK2hDVDRwbnRqZG80STR6a0w2QTlaNVNRRm13UTczWTkyV2o3SXo0L3djVkxaClFMYTdpUlNHdld3c3ZUSExGVkEyQjBsa0crczFzWjZRQ0gxWWY3WmJBb0dBVnZXVDZWLzZqS2d0SlV6Y1NjNmYKRjk4ZDZvTmpmVWpYdE1PT3FLRHBrTStnenRNZEVBeXo4L09TU04xTXp2MTA4NldLZzcyM0dDcGFDcStKdWdHMgpLT29YYjF1eUlaUGY2RnF5azVwQnM1SlJ5Lzd6Uk5IWjVtNWhSbTBGcFdBMGRTTExDWXFPbjBjT0lTNEV6d1pnCmtQQ1BsNWZtQkVveTRFVTNCRk0xS3QwQ2dZRUFueFdGaE9lRElkbGh0bE8rSlpneEw0VHZ4OUptdEllSHZRWWIKbEU0WENJYjQyWi93ZEF0cDNVUi91LzBteTVIMkUvWE5qRytid0FiZS9YZE1VN1BkckNDS0NINHE4V0d5NUZERwpLVFMzakhiak9MbkRWdjZ3b2E1eVovYThvaHBzNGswWHE1MG1xNStvcnhpUHgvSzF2NW5sSkJyRWYwenBVbW9nClpaeDM3VWtDZ1lFQThGWnRDMUE0eHdSeGdjWGRSTUNCYmdkM00yWUVibFNhV3lUb2wyVFRKYkxCbHRzVGpvUmYKNjZDdUFSa3BPaEh2ZWREY2pYOGcyWGh0MTZIaWtPTkMycTZZbkUrVGQzM085Vkh2QVBsK1NLaGlJTnVWWGdwdgpkY3VDZGpDMi9IYlQrUU5UbUNIWGloUllLeThZZ3U3eVE1c2E0aDM3ZG9GbCtHNlowdFdBZmxzPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\n"
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-10-07T15:42:17+08:00",
    "KeyValueList": [
      {
        "key": "ClusterId",
        "value": "c6e077638479f4f34aebd09db488bac55"
      },
      {
        "key": "ClusterSpec",
        "value": "ack.pro.small"
      },
      {
        "key": "ClusterType",
        "value": "ManagedKubernetes"
      },
      {
        "key": "Created",
        "value": "2026-10-07T15:42:17+08:00"
      },
      {
        "key": "CurrentVersion",
        "value": "1.35.7-aliyun.1"
      },
      {
        "key": "DeletionProtection",
        "value": "false"
      },
      {
        "key": "ExternalLoadbalancerId",
        "value": "lb-mj7d0gxerspz8pjrrif6h"
      },
      {
        "key": "InitVersion",
        "value": "1.35.7-aliyun.1"
      },
      {
        "key": "MaintenanceWindow",
        "value": "{enable:false,weekly_period:}"
      },
      {
        "key": "MasterUrl",
        "value": "{\\api_server_endpoint\\:\\https://43.108.104.153:6443\\,\\intranet_api_server_endpoint\\:\\https://10.0.1.70:6443\\}"
      },
      {
        "key": "MetaData",
        "value": "{\\Addons\\:[{\\name\\:\\storage-operator\\,\\version\\:\\v1.35.4\\},{\\name\\:\\csi-provisioner\\,\\version\\:\\v1.37.2\\},{\\name\\:\\gateway-api\\,\\version\\:\\1.3.0\\},{\\name\\:\\cloud-controller-manager\\,\\version\\:\\v2.15.0\\},{\\name\\:\\alicloud-monitor-controller\\,\\version\\:\\v1.8.13\\},{\\name\\:\\metrics-server\\,\\version\\:\\v0.3.16-ab96b23-aliyun\\},{\\name\\:\\metrics-aggregator\\,\\version\\:\\v1.9.6\\},{\\name\\:\\ack-scheduler\\,\\version\\:\\v1.35.2-apsara.6.12.3.a6eed53c\\},{\\name\\:\\pod-readinessgate-webhook\\,\\version\\:\\v1.0.0\\},{\\name\\:\\coredns\\,\\version\\:\\v1.13.2.3\\},{\\name\\:\\kube-flannel-ds\\,\\version\\:\\v0.28.0.6\\},{\\name\\:\\ack-nvidia-device-plugin\\,\\version\\:\\0.8.1\\},{\\name\\:\\managed-kube-proxy\\,\\version\\:\\v1.35.7-aliyun.1\\},{\\name\\:\\ack-ram-authenticator\\,\\version\\:\\0.5.1\\},{\\name\\:\\csi-plugin\\,\\version\\:\\v1.37.2\\},{\\name\\:\\cnfs-controller\\,\\version\\:\\v1.2.5\\,\\config\\:\\{\\\\\\CreateDefaultCNFS\\\\\\:false}\\},{\\name\\:\\kube-apiserver\\,\\version\\:\\v1.35.7-aliyun.1\\},{\\name\\:\\kube-controller-manager\\,\\version\\:\\v1.35.7-aliyun.1\\}],\\AuditProjectName\\:\\\\,\\Capabilities\\:{\\AnyAZ\\:true,\\CSI\\:true,\\CpuPolicy\\:true,\\DeploymentSet\\:true,\\DisableEncryption\\:true,\\EncryptionKMSKeyId\\:\\\\,\\EnterpriseSecurityGroup\\:true,\\HpcCluster\\:true,\\IntelSGX\\:false,\\Knative\\:true,\\Network\\:\\Flannel\\,\\NgwPayByLcu\\:true,\\NodeCIDRMask\\:\\25\\,\\NodeNameMode\\:true,\\ProxyMode\\:\\\\,\\PublicSLB\\:true,\\RamRoleType\\:\\restricted\\,\\SLSProjectName\\:true,\\SandboxRuntime\\:true,\\SnapshotPolicy\\:true,\\Taint\\:true,\\TerwayEniip\\:true,\\UserData\\:true},\\CloudMonitorVersion\\:\\\\,\\ClusterDomain\\:\\\\,\\ControlPlaneLogConfig\\:{\\components\\:null},\\DockerVersion\\:\\\\,\\EtcdVersion\\:\\\\,\\ExtraCertSAN\\:null,\\FreeTier\\:false,\\HasSandboxRuntime\\:false,\\IPStack\\:\\ipv4\\,\\ImageType\\:\\AliyunLinux3ContainerOptimized\\,\\KubernetesVersion\\:\\1.35.7-aliyun.1\\,\\MultiAZ\\:false,\\NameMode\\:\\\\,\\NextVersion\\:\\\\,\\OSType\\:\\Linux\\,\\Platform\\:\\AliyunLinux\\,\\PodVswitchId\\:\\\\,\\Provider\\:\\\\,\\RRSAConfig\\:{\\enabled\\:false},\\RamRoleType\\:\\\\,\\ResourceGroupId\\:\\rg-acfnvekhilw5kmy\\,\\Runtime\\:\\containerd\\,\\RuntimeVersion\\:\\2.3.4\\,\\ServiceCIDR\\:\\172.20.0.0/16\\,\\SubClass\\:\\default\\,\\SupportPlatforms\\:[\\CentOS\\,\\AliyunLinux\\,\\Windows\\,\\WindowsCore\\],\\Timezone\\:\\\\,\\VSwitchIds\\:[\\vsw-mj7mkdtynodmq5f1ov0ec\\],\\VersionSpec\\:null,\\VpcCidr\\:\\10.0.0.0/22\\,\\ack-nvidia-device-pluginVersion\\:\\0.8.1\\,\\ack-ram-authenticatorVersion\\:\\0.5.1\\,\\ack-schedulerVersion\\:\\v1.35.2-apsara.6.12.3.a6eed53c\\,\\alicloud-monitor-controllerVersion\\:\\v1.8.13\\,\\cloud-controller-managerVersion\\:\\v2.15.0\\,\\cnfs-controllerVersion\\:\\v1.2.5\\,\\corednsVersion\\:\\v1.13.2.3\\,\\csi-pluginVersion\\:\\v1.37.2\\,\\csi-provisionerVersion\\:\\v1.37.2\\,\\gateway-apiVersion\\:\\1.3.0\\,\\kube-apiserverVersion\\:\\v1.35.7-aliyun.1\\,\\kube-controller-managerVersion\\:\\v1.35.7-aliyun.1\\,\\kube-flannel-dsVersion\\:\\v0.28.0.6\\,\\metrics-aggregatorVersion\\:\\v1.9.6\\,\\metrics-serverVersion\\:\\v0.3.16-ab96b23-aliyun\\,\\pod-readinessgate-webhookVersion\\:\\v1.0.0\\,\\storage-operatorVersion\\:\\v1.35.4\\}"
      },
      {
        "key": "Name",
        "value": "tbv3okbhruku5snohcm7"
      },
      {
        "key": "NetworkMode",
        "value": "vpc"
      },
      {
        "key": "Parameters",
        "value": "{ALIYUN::AccountId:5469257408566579,ALIYUN::NoValue:None,ALIYUN::Region:ap-northeast-2,ALIYUN::ResourceGroupId:rg-acfnvekhilw5kmy,ALIYUN::StackId:f5cc0480-7fab-4de5-8b79-10abe998047f,ALIYUN::StackName:k8s-for-cs-c6e077638479f4f34aebd09db488bac55,ALIYUN::TenantId:5469257408566579,AdjustmentType:TotalCapacity,BetaVersion:,CloudMonitorFlags:False,CloudMonitorVersion:1.3.7,ClusterDns:172.20.0.10,ClusterId:c6e077638479f4f34aebd09db488bac55,ContainerCIDR:172.19.0.0/16,CustomK8sWorkerRole:,DisableAddons:True,DisableAutoCreateK8sWorkerRole:False,DisableAutoCreateK8sWorkerRolePolicy:True,DockerVersion:17.06.2-ce-3,ESSDeletionProtection:True,Eip:True,EipAddress:,EtcdVersion:v3.1.11,ExecuteVersion:63084271,HealthCheckType:NONE,IPStack:ipv4,ImageId:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,K8sWorkerPolicyDocument:{\\Version\\: \\1\\, \\Statement\\: [{\\Action\\: [\\ecs:DescribeInstanceAttribute\\, \\ecs:DescribeInstances\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}, {\\Action\\: [\\log:GetProject\\, \\log:GetLogStore\\, \\log:GetConfig\\, \\log:GetMachineGroup\\, \\log:GetAppliedMachineGroups\\, \\log:GetAppliedConfigs\\, \\log:GetIndex\\, \\log:GetSavedSearch\\, \\log:GetDashboard\\, \\log:GetJob\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}, {\\Action\\: [\\cr:GetAuthorizationToken\\, \\cr:ListInstanceEndpoint\\, \\cr:PullRepository\\, \\cr:GetInstanceVpcEndpoint\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}]},KeyPair:,KubernetesVersion:1.35.7-aliyun.1,MasterSLBPrivateIP:10.0.1.70,NatGateway:False,NatGatewayId:ngw-mj7r0n0opzyvi2m4q0v9v,NatGatewayType:Enhanced,NatGatewayVswitchId:,Network:Flannel,NodeNameMode:nodeip,NumOfNodes:0,OSType:Linux,Password:******,PodVswitchIds:[],ProtectedInstances:,ProxyMode:iptables,RemoveInstanceIds:,ResourceGroupId:rg-acfnvekhilw5kmy,SNatEntry:True,ScaleOutToken:None,SecurityGroupId:sg-mj7hydw78w20s8gutrve,ServiceCIDR:172.20.0.0/16,SetUpArgs:--addon-names kube-flannel-ds,csi-plugin,csi-provisioner --labels CB-SPIDER:PMKS:CLUSTER=owned --node-cidr-mask 25 --registry-url registry-ap-northeast-2-vpc.ack.aliyuncs.com --runtime containerd --runtime-version 2.3.4,SnatTableId:stb-mj7dos9owfy0vkja6ibfb,Tags:[{\\Key\\: \\CB-SPIDER:PMKS:CLUSTER\\, \\Value\\: \\owned\\}, {\\Key\\: \\ack.aliyun.com\\, \\Value\\: \\c6e077638479f4f34aebd09db488bac55\\}],UserData:,VpcCidrWithSecondaryCidrs:[\\10.0.0.0/22\\],VpcId:vpc-mj7s33x3t7vhkkospgg4f,WorkerAutoRenew:False,WorkerAutoRenewPeriod:1,WorkerDataDisk:False,WorkerDataDisks:[],WorkerDeletionProtection:True,WorkerDeploymentSetId:,WorkerHpcClusterId:,WorkerImageId:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,WorkerInstanceChargeType:PostPaid,WorkerInstanceTypes:ecs.n4.large,WorkerKeyPair:,WorkerLoginPassword:******,WorkerPeriod:3,WorkerPeriodUnit:Month,WorkerSnapshotPolicyId:******,WorkerSystemDiskCategory:cloud_ssd,WorkerSystemDiskPerformanceLevel:null,WorkerSystemDiskSize:40,WorkerVSwitchIds:vsw-mj7mkdtynodmq5f1ov0ec,ZoneId:}"
      },
      {
        "key": "Profile",
        "value": "Default"
      },
      {
        "key": "RegionId",
        "value": "ap-northeast-2"
      },
      {
        "key": "ResourceGroupId",
        "value": "rg-acfnvekhilw5kmy"
      },
      {
        "key": "SecurityGroupId",
        "value": "sg-mj7hydw78w20s8gutrve"
      },
      {
        "key": "Size",
        "value": "0"
      },
      {
        "key": "State",
        "value": "running"
      },
      {
        "key": "SubnetCidr",
        "value": "172.19.0.0/16"
      },
      {
        "key": "Tags",
        "value": "{key:CB-SPIDER:PMKS:CLUSTER,value:owned}; {key:ack.aliyun.com,value:c6e077638479f4f34aebd09db488bac55}"
      },
      {
        "key": "Updated",
        "value": "2026-10-07T15:47:00+08:00"
      },
      {
        "key": "VpcId",
        "value": "vpc-mj7s33x3t7vhkkospgg4f"
      },
      {
        "key": "VswitchId",
        "value": "vsw-mj7mkdtynodmq5f1ov0ec"
      },
      {
        "key": "WorkerRamRoleName",
        "value": "KubernetesWorkerRole-f5cc0480-7fab-4de5-8b79-10abe998047f"
      },
      {
        "key": "ZoneId",
        "value": "ap-northeast-2a"
      }
    ]
  }
}
```

</details>

