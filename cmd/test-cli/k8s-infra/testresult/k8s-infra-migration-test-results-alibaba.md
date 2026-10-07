# CM-Beetle K8s Infra Migration Test Results — Alibaba-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → migrate → list → get (verified against
> the recommendation) → delete → residual resource check.

## Environment

- CSP / Region: alibaba / ap-northeast-2
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (8928ba9)
- Git Commit: 8928ba9
- Namespace: mig01
- Test Date: 2026-09-16 20:13:53 KST
- Cluster ID: k8sfix02-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 20ms |
| 2 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 6m0.085s |
| 3 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 2ms |
| 4 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 1ms |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 26.226s |
| 6 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 14m18.944s |
| 7 | Residual resource check (Tumblebug) | ✅ **PASS** | 3ms |

**Overall Result**: 7/7 steps passed ✅

**Total Duration**: 20m45s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 20ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.34.10-aliyun.1)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=alibaba+ap-northeast-2+ecs.e-c1m2.xlarge nodes=2

### Step 2 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 6m0.085s
- **Status Code**: 202

- ℹ️  nameSeed: k8sfix02
- ℹ️  async reqId: 1789557233061360649
- ℹ️  cluster id: k8sfix02-on-prem-k8s-cluster
- ℹ️  elapsed: 6m0s
- ✅ status: Active

### Step 3 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 2ms
- **Status Code**: 200

- ✅ migrated cluster present in list (2 total)

### Step 4 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 1ms
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=alibaba+ap-northeast-2+ecs.e-c1m2.xlarge, nodes=2)
- ✅ version: 1.34.10-aliyun.1 (recommended 1.34.10-aliyun.1)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 26.226s

- ✅ kubeconfig obtained (server: https://43.108.35.249:6443)
- ℹ️  auth method: client certificate in kubeconfig
- ✅ API server reachable (v1.34.10-aliyun.1)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 1)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: 8.213.133.141
- ✅ nginx served over the LoadBalancer at http://8.213.133.141/ (attempt 1)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 6 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 14m18.944s
- **Status Code**: 200

- ✅ deleted on attempt 1 (14m18s)

### Step 7 — Residual resource check (Tumblebug)

- **Duration**: 3ms

- ℹ️  VNet k8sfix02-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup k8sfix02-k8s-sg still exists (known gap)
- ℹ️  SshKey k8sfix02-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for alibaba ap-northeast-2 (source: v1.32.3 → target: v1.34.10-aliyun.1)",
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
    "version": "1.34.10-aliyun.1",
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
  "id": "k8sfix02-on-prem-k8s-cluster",
  "uid": "tbqfbinr8jv2lc7g6aeq",
  "name": "k8sfix02-on-prem-k8s-cluster",
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
    "sys.createdTime": "2026-09-16 19:14:47 +0800 +0800",
    "sys.cspResourceId": "c7bb68aecc47b4ac38f7c0c690d5fa579",
    "sys.cspResourceName": "tbqfbinr8jv2lc7g6aeq",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "k8sfix02-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "k8sfix02-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbqfbinr8jv2lc7g6aeq",
    "sys.version": "1.34.10-aliyun.1"
  },
  "systemLabel": "",
  "version": "1.34.10-aliyun.1",
  "network": {
    "vNetId": "k8sfix02-k8s-vpc",
    "subnetIds": [
      "k8sfix02-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "k8sfix02-k8s-sg"
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
      "sshKeyId": "k8sfix02-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 0,
      "maxNodeSize": 0,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "worker-k8s-for-cs-c7bb68aecc47b4ac38f7c0c690d5fa579",
          "cspResourceId": "i-mj7012hrh739bepeuwrp"
        },
        {
          "cspResourceName": "worker-k8s-for-cs-c7bb68aecc47b4ac38f7c0c690d5fa579",
          "cspResourceId": "i-mj7012hrh739bepeuwro"
        }
      ],
      "keyValueList": [
        {
          "key": "AutoScaling",
          "value": "{eip_bandwidth:0,eip_internet_charge_type:,enable:false,max_instances:0,min_instances:0,type:}"
        },
        {
          "key": "KubernetesConfig",
          "value": "{cms_enabled:false,cpu_policy:none,node_name_mode:nodeip,runtime:containerd,runtime_version:2.1.9,unschedulable:false,user_data:}"
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
          "value": "{created:2026-09-16T19:17:49.109695065+08:00,is_default:false,name:workers1,nodepool_id:np30f252a975da4f39a1ed043e9f0e41a6,region_id:ap-northeast-2,resource_group_id:,type:ess,updated:2026-09-16T19:19:37.827+08:00}"
        },
        {
          "key": "ScalingGroup",
          "value": "{auto_renew:false,auto_renew_period:0,deploymentset_id:,desired_size:2,image_id:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,image_type:AliyunLinux3ContainerOptimized,instance_charge_type:PostPaid,instance_types:[ecs.e-c1m2.xlarge],internet_charge_type:,internet_max_bandwidth_out:0,key_pair:tb9e95hm088kd3tge6mq,login_password:,multi_az_policy:BALANCE,period:0,period_unit:,platform:AliyunLinux,private_pool_options:{},ram_policy:KubernetesWorkerRole-00623c79-cd5b-4325-bcd2-b0f0b9bf8396,scaling_group_id:asg-mj73ravctixal7xont26,scaling_policy:release,security_group_id:sg-mj78e5i8z3rmauatexhd,security_group_ids:[sg-mj78e5i8z3rmauatexhd],spot_strategy:NoSpot,system_disk_categories:[cloud_essd],system_disk_category:cloud_essd,system_disk_encrypt_algorithm:,system_disk_encrypted:false,system_disk_kms_key_id:,system_disk_performance_level:,system_disk_size:100,vswitch_ids:[vsw-mj7qh726wq5bds252lzmj]}"
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
      "cspResourceId": "np30f252a975da4f39a1ed043e9f0e41a6",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "np30f252a975da4f39a1ed043e9f0e41a6"
        },
        "ImageIID": {
          "NameId": "aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd",
          "SystemId": "aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd"
        },
        "VMSpecName": "ecs.e-c1m2.xlarge",
        "RootDiskType": "cloud_essd",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tb9e95hm088kd3tge6mq",
          "SystemId": "tb9e95hm088kd3tge6mq"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "worker-k8s-for-cs-c7bb68aecc47b4ac38f7c0c690d5fa579",
            "SystemId": "i-mj7012hrh739bepeuwrp"
          },
          {
            "NameId": "worker-k8s-for-cs-c7bb68aecc47b4ac38f7c0c690d5fa579",
            "SystemId": "i-mj7012hrh739bepeuwro"
          }
        ],
        "KeyValueList": [
          {
            "key": "AutoScaling",
            "value": "{eip_bandwidth:0,eip_internet_charge_type:,enable:false,max_instances:0,min_instances:0,type:}"
          },
          {
            "key": "KubernetesConfig",
            "value": "{cms_enabled:false,cpu_policy:none,node_name_mode:nodeip,runtime:containerd,runtime_version:2.1.9,unschedulable:false,user_data:}"
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
            "value": "{created:2026-09-16T19:17:49.109695065+08:00,is_default:false,name:workers1,nodepool_id:np30f252a975da4f39a1ed043e9f0e41a6,region_id:ap-northeast-2,resource_group_id:,type:ess,updated:2026-09-16T19:19:37.827+08:00}"
          },
          {
            "key": "ScalingGroup",
            "value": "{auto_renew:false,auto_renew_period:0,deploymentset_id:,desired_size:2,image_id:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,image_type:AliyunLinux3ContainerOptimized,instance_charge_type:PostPaid,instance_types:[ecs.e-c1m2.xlarge],internet_charge_type:,internet_max_bandwidth_out:0,key_pair:tb9e95hm088kd3tge6mq,login_password:,multi_az_policy:BALANCE,period:0,period_unit:,platform:AliyunLinux,private_pool_options:{},ram_policy:KubernetesWorkerRole-00623c79-cd5b-4325-bcd2-b0f0b9bf8396,scaling_group_id:asg-mj73ravctixal7xont26,scaling_policy:release,security_group_id:sg-mj78e5i8z3rmauatexhd,security_group_ids:[sg-mj78e5i8z3rmauatexhd],spot_strategy:NoSpot,system_disk_categories:[cloud_essd],system_disk_category:cloud_essd,system_disk_encrypt_algorithm:,system_disk_encrypted:false,system_disk_kms_key_id:,system_disk_performance_level:,system_disk_size:100,vswitch_ids:[vsw-mj7qh726wq5bds252lzmj]}"
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
    "endpoint": "https://43.108.35.249:6443",
    "kubeconfig": "\napiVersion: v1\nclusters:\n- cluster:\n    server: https://43.108.35.249:6443\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUR2VENDQXFXZ0F3SUJBZ0lJQWhHbEg2SUFIb013RFFZSktvWklodmNOQVFFTEJRQXdRREVSTUE4R0ExVUUKQ2hNSWFHRnVaM3BvYjNVeEZqQVVCZ05WQkFzVERXRnNhV0poWW1FZ1kyeHZkV1F4RXpBUkJnTlZCQU1UQ210MQpZbVZ5Ym1WMFpYTXdJQmNOTWpZd09URTJNVEV3T1RBd1doZ1BNakExTmpBNU1EZ3hNVEUwTkRkYU1FQXhFVEFQCkJnTlZCQW9UQ0doaGJtZDZhRzkxTVJZd0ZBWURWUVFMRXcxaGJHbGlZV0poSUdOc2IzVmtNUk13RVFZRFZRUUQKRXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBeURNbgp2ZFRacjNXZStnU29nZy9iUXcrY3UzUlRyRU5uUEJ2cGNnQVRyMVpVeERLWHQrZ1JWc2NNSzJCQnpwQXF6cEp4CkZqS2RFb3BkWjZQWkhGeTlLRUc5TDJtQy9tR3QydkszdVV4TElnUHF6clhhTEg4d0cra3lzNUtTUFIxb3YxS3EKdHJ5Wmw3MXd4TzV6RFYyMGdZbHkyVmxPQWEyT0h0L3ZzM2ZqSXVBcWREZ3BFUVE4TVFiRHBMNVlHbUQvZGZXSgpTYjNMcW9yNDVob3hrYkNQUUd6c1Z3ZUtKcVVHVXYvMHhoMHBKamtjYkYxRGNJT1U5N0tFL09yLzJoSHNWdk1RCmpRMDF4QWF6RUpiOWZEZE0wT2JrWTdRQjRsaU8xSUZITVJ0UWZPUGViYmhqZlR3RHBCblpYYkp3SEdYV1cwR1IKdUdOci9sOUZKUFJEcm9aZzF3SURBUUFCbzRHNE1JRzFNQTRHQTFVZER3RUIvd1FFQXdJQ3JEQVBCZ05WSFJNQgpBZjhFQlRBREFRSC9NQjBHQTFVZERnUVdCQlMvdWJKMVR5OEJTRVI2b1ZWbS9UWEpVTndUdFRBOEJnZ3JCZ0VGCkJRY0JBUVF3TUM0d0xBWUlLd1lCQlFVSE1BR0dJR2gwZEhBNkx5OWpaWEowY3k1aFkzTXVZV3hwZVhWdUxtTnYKYlM5dlkzTndNRFVHQTFVZEh3UXVNQ3d3S3FBb29DYUdKR2gwZEhBNkx5OWpaWEowY3k1aFkzTXVZV3hwZVhWdQpMbU52YlM5eWIyOTBMbU55YkRBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQWFiTnpxMFVTcmZnRzFRbS9rMTVVCk5lcUptTmVJWlU1d1VTTWljUkVmTTV0MmJJNDJCZDY3dTM0OFkzMU94WjA5QXNrT1pVSGd4TU0xNWRxbFdUK0UKVUpQbUJrbklyajFTd3NuaEJLNnU4eDV3QkhZOTkvZ3IrSzNVenhNYUExU2M5bXFEVzlNby9PdmR0dVphWEE0dQpIdElmakpXV3ZUQVBBR1BXQi85Q3pjOTMxNjNwRjM3TE9tNnY2WkpNYkJIbWVaNDh4cWFFTlhENHB0QXcybElCClVZbEI1OHlTZW5Ka0lZZlNZSjVBZGZkdlNOWitwaEhKSnRWM05aUDlHbXBWVDNBVStCcVlnNFJla1pyK3VpY3MKZTcrNE5neTljOWVWUDZQd3FmY2tpT2MxRDZCQWdxc3lBd2dGSStwUkJzUVB6NXJ5TFhHUUpodlhYZlVrMXc1cgpvUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: kubernetes\ncontexts:\n- context:\n    cluster: kubernetes\n    user: \"213652363353305376\"\n  name: 213652363353305376-c7bb68aecc47b4ac38f7c0c690d5fa579\ncurrent-context: 213652363353305376-c7bb68aecc47b4ac38f7c0c690d5fa579\nkind: Config\npreferences: {}\nusers:\n- name: \"213652363353305376\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUQyVENDQXNHZ0F3SUJBZ0lJQWhHbFlZUUFIb013RFFZSktvWklodmNOQVFFTEJRQXdRREVSTUE4R0ExVUUKQ2hNSWFHRnVaM3BvYjNVeEZqQVVCZ05WQkFzVERXRnNhV0poWW1FZ1kyeHZkV1F4RXpBUkJnTlZCQU1UQ210MQpZbVZ5Ym1WMFpYTXdIaGNOTWpZd09URTJNVEV4TWpBd1doY05Namt3T1RFMU1URXhOek0yV2pCS01SVXdFd1lEClZRUUtFd3h6ZVhOMFpXMDZkWE5sY25NeENUQUhCZ05WQkFzVEFERW1NQ1FHQTFVRUF4TWRNakV6TmpVeU16WXoKTXpVek16QTFNemMyTFRFM09EazFOVGMwTlRZd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFSwpBb0lCQVFDNUs0VTR2NFp3Z3ZUYW5xSXEvbnFPSGI4NFhuU3VJZDEvak9jb1BseE9jWTJ2YmVkOE1TM2FROE1TCnh2ZVZPZlJDNTZlNGhFQVVjR294KzR1NFI3Zk0vdVFwQ2puemhLaXp1ekxMaGcrRFFXNHpYVTY1Wm5XbUtkcHgKb1ZRRnRRbkdsZGR6dVh2b1NFM3QzTzZOZCtCM21MQzJieUd5cnNJdVdIVXMyN2xwREpBdDkzd25qdlNkTkl5Rgo3VEoyOEtjeWFVeEdWRDlGNXQrNXRyV0pXMGhqL3ZqOVprWG1oODd2YUJIeGY0MDJZekhaYWV5RExQZmp4ek1JCjFibklONWpFQ2IzdGFKdXk5NEpCakdqTG5IK0FQUThKNDV5VGFYMWhDck55YVFyVjZ3TnJKdDJ3NEllU3Iwd0IKYVFUcElHZUhVdFErdzBjLzBvWndpQ3JHMWNPTkFnTUJBQUdqZ2N3d2dja3dEZ1lEVlIwUEFRSC9CQVFEQWdlQQpNQk1HQTFVZEpRUU1NQW9HQ0NzR0FRVUZCd01DTUF3R0ExVWRFd0VCL3dRQ01BQXdId1lEVlIwakJCZ3dGb0FVCnY3bXlkVTh2QVVoRWVxRlZadjAxeVZEY0U3VXdQQVlJS3dZQkJRVUhBUUVFTURBdU1Dd0dDQ3NHQVFVRkJ6QUIKaGlCb2RIUndPaTh2WTJWeWRITXVZV056TG1Gc2FYbDFiaTVqYjIwdmIyTnpjREExQmdOVkhSOEVMakFzTUNxZwpLS0FtaGlSb2RIUndPaTh2WTJWeWRITXVZV056TG1Gc2FYbDFiaTVqYjIwdmNtOXZkQzVqY213d0RRWUpLb1pJCmh2Y05BUUVMQlFBRGdnRUJBRmlmRE8rZ3huVm8zQzlDKzVFa0F3TVV5OTNZdU81NytQRUN1YjFEYWFRSnJUc04KTHlxRmc5cHpVMytRYVBhYVNVVElXeFhGdDVSOEplVkRwRzRIZHRhRVFEY2NaWkQxOUhNdm9HMDNUa256dGFoNQpXZ01xTUVMTDNyTU41TXYrMWpKbXdUMDBTVlFJeE1oRlpZTmNEc214UXhUT3JXc0grclRVUlJMY0JVV3NtZlJLCnhDQXExRGVIdmY3ZmVCUVhaWjlweEZJU0NxUjhCeWM2dnFIZ3ZmVEVDaklBZW9Hb3FWRWdXbnE2TTZGdDJaWncKNmlFeE5BRWZEcEpieE1qMmZ2SVEyV2tBQXlGVDZrVUEvRmJFNmFNeFBGR3pnMWZvVytBbUdKblhqeWdHdmhDMwpBUEE5a2F2VVIvcy8ra0NzWjNod3VrWDUyR05RRWhldGVlR1c3anc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBdVN1Rk9MK0djSUwwMnA2aUt2NTZqaDIvT0Y1MHJpSGRmNHpuS0Q1Y1RuR05yMjNuCmZERXQya1BERXNiM2xUbjBRdWVudUlSQUZIQnFNZnVMdUVlM3pQN2tLUW81ODRTb3M3c3l5NFlQZzBGdU0xMU8KdVdaMXBpbmFjYUZVQmJVSnhwWFhjN2w3NkVoTjdkenVqWGZnZDVpd3RtOGhzcTdDTGxoMUxOdTVhUXlRTGZkOApKNDcwblRTTWhlMHlkdkNuTW1sTVJsUS9SZWJmdWJhMWlWdElZLzc0L1daRjVvZk83MmdSOFgrTk5tTXgyV25zCmd5ejM0OGN6Q05XNXlEZVl4QW05N1dpYnN2ZUNRWXhveTV4L2dEMFBDZU9jazJsOVlRcXpjbWtLMWVzRGF5YmQKc09DSGtxOU1BV2tFNlNCbmgxTFVQc05IUDlLR2NJZ3F4dFhEalFJREFRQUJBb0lCQUhOWjB5SG8zZDBWRjJkaApUdkN0bXJjUmZOK21wOVVhTDV1WHNTQjJ5Slh0VXpBbnJQN0d2Q090OXNXcHdPM3JMbWpIV3NvdFNLWHk2WlM5CmVPcVJjc1IzUExiOE1lR3JrRlIybDB3RjlYLzBzS3U2d3FIb1cvM3BNTFY4cEpUeGxHZHJUTTVBakJuRmdSWmUKZlJVUHludDJXUThzNTdjaCtzRllSWlp1UW5CWFhjeHB1UTNPSW9OVHRSQXJvN28vRXlsTlBoa01lTGUxaEdRWgpLa1RQSWFmLzNpV0lzb2hNbG40M2JNa0JYeHVISEJsL1ZXMG03am5mam5OUjVKZjJtR0tUSW9GVVM4TTMwZU1WCkt0ZW9oeUpXeTV2Wk9rNGkxSEE2aVl2Zk1FbzNUeWQvZDd6dWhRUEdFbFVPNUFzanZsSXd4Z1ZJVDdhRHVVT08KTnN6cEVGa0NnWUVBNEtIL1JkWmNVWWVzemJWY1ZaVkk1eVZrWlc5NnRvRTFHV0dzL3RrS1dDK3FWU0E5ZkdPVwp6ZmdMcFRCYkVuY0ZYSC9STndmeUJzQTluZnBUYk5TMUwrbmtpSmtMV2o0S3pRYzRUMEZFaWJJZithUjFrWEF1CjlkdnB1d2RDOU91U0ZOTUpPbFVsa3FEUGFEOWtaK3preDNyTG01a2ZwYXd4QmxHTG5KNlNwTjhDZ1lFQTB3YlcKUjQwdnhwUWVOYzRFQUt5ZFBwbVova09hR21zd0NZZFNCZHhEZ1FJeDBBL3V3bWd5YWEzRlV3Rko4NHl6RFZFWQo1Y0xqTHdsS1dTNHBEeVpoZ3lTeE5hY1hxSjlOcXhCZlJqZ0tjR25QbDhmL1FVNXByaGJ2N0lDRG1tQ2pDbmZFCndJbnVDV2p4MzNNQWV2WVcwaGFLcC94bFc5U3lmclZiVTF1ZVdSTUNnWUVBdjFINlkzYmhoaWN5VExvcy9FOTcKbWpORXdRZ2owaEZXTWxuNHg4My9lNENOYUpkZkJ2U2pXcVhxOVRTc1BKdldteVBhQXk3bk9reTlyakdvb0ViQwpVeDY0b0ljSkhrRTlYY2JLZE9ZNEE3Y01lTWxUd1IxMVFiMmQ0c2VhaFpPbUJjcUFUNGg1eWRyaEMvOEIrMm5PCnFQK3pMc0ZLTFFidGNsNDl0SlZ5ZkswQ2dZQndJMi8yOGRYemhuNVBSVHpuUnRNQWt3czhESngxY1lSRGEvOXcKVWM0bFhnOHVhMmtMTWVlb05NbHBCSVJSd2ZEY0lMNUVTajREbnNJOWhjUVg0dU5xbHpMOE9lRGVvRmpia0lXdwpnTk03VFY4ZFh4QUxtaW1hYzJIbm9adE5qQkRYM1RGV3gzVVExdzNCR2hHbFJlUjJtN2Z1OExiRUI5RWFlREVoCnFNd0t0UUtCZ0I5czJWcm1uTWdnVVpRdlZ4SHRoeGJScEZodGtMYVpyY0NKV1hIcHRwbGFONmFzZU5nQlF1WUkKMFl5aHV4c2lsNGl2cUlKdEpxMVJwbkpFYnIvTk1yNjlTSXVuZnlLdFdKTUJycUhmSEFwQWlaOGdDOGliKzBzVAp3bWJjdUM4cURuMncwcmFGalVqRGt0T0NhRnA0YVhZQ2V3Sk5UdThrSkZhS3NjNzQ2Qk8wCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==\n"
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-09-16T19:14:47+08:00",
  "keyValueList": [
    {
      "key": "ClusterId",
      "value": "c7bb68aecc47b4ac38f7c0c690d5fa579"
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
      "value": "2026-09-16T19:14:47+08:00"
    },
    {
      "key": "CurrentVersion",
      "value": "1.34.10-aliyun.1"
    },
    {
      "key": "DeletionProtection",
      "value": "false"
    },
    {
      "key": "ExternalLoadbalancerId",
      "value": "lb-mj7fbk5ltg8uaqv90lx88"
    },
    {
      "key": "InitVersion",
      "value": "1.34.10-aliyun.1"
    },
    {
      "key": "MaintenanceWindow",
      "value": "{enable:false,weekly_period:}"
    },
    {
      "key": "MasterUrl",
      "value": "{\\api_server_endpoint\\:\\https://43.108.35.249:6443\\,\\intranet_api_server_endpoint\\:\\https://10.0.1.3:6443\\}"
    },
    {
      "key": "MetaData",
      "value": "{\\Addons\\:[{\\name\\:\\alicloud-monitor-controller\\,\\version\\:\\v1.8.13\\},{\\name\\:\\metrics-server\\,\\version\\:\\v0.3.16-ab96b23-aliyun\\},{\\name\\:\\csi-provisioner\\,\\version\\:\\v1.37.2\\},{\\name\\:\\cloud-controller-manager\\,\\version\\:\\v2.15.0\\},{\\name\\:\\csi-plugin\\,\\version\\:\\v1.37.2\\},{\\name\\:\\ack-scheduler\\,\\version\\:\\v1.34.0-apsara.6.11.14.80067fc3\\},{\\name\\:\\coredns\\,\\version\\:\\v1.13.2.2\\},{\\name\\:\\gateway-api\\,\\version\\:\\1.3.0\\},{\\name\\:\\storage-operator\\,\\version\\:\\v1.35.3\\},{\\name\\:\\managed-kube-proxy\\,\\version\\:\\v1.34.10-aliyun.1\\},{\\name\\:\\kube-flannel-ds\\,\\version\\:\\v0.15.1.23-33d25c1-aliyun\\},{\\name\\:\\ack-ram-authenticator\\,\\version\\:\\0.5.1\\},{\\name\\:\\cnfs-controller\\,\\version\\:\\v1.2.5\\,\\config\\:\\{\\\\\\CreateDefaultCNFS\\\\\\:false}\\},{\\name\\:\\ack-nvidia-device-plugin\\,\\version\\:\\0.8.1\\},{\\name\\:\\kube-apiserver\\,\\version\\:\\v1.34.10-aliyun.1\\},{\\name\\:\\pod-readinessgate-webhook\\,\\version\\:\\v1.0.0\\},{\\name\\:\\metrics-aggregator\\,\\version\\:\\v1.9.6\\},{\\name\\:\\kube-controller-manager\\,\\version\\:\\v1.34.10-aliyun.1\\}],\\AuditProjectName\\:\\\\,\\Capabilities\\:{\\AnyAZ\\:true,\\CSI\\:true,\\CpuPolicy\\:true,\\DeploymentSet\\:true,\\DisableEncryption\\:true,\\EncryptionKMSKeyId\\:\\\\,\\EnterpriseSecurityGroup\\:true,\\HpcCluster\\:true,\\IntelSGX\\:false,\\Knative\\:true,\\Network\\:\\Flannel\\,\\NgwPayByLcu\\:true,\\NodeCIDRMask\\:\\25\\,\\NodeNameMode\\:true,\\ProxyMode\\:\\\\,\\PublicSLB\\:true,\\RamRoleType\\:\\restricted\\,\\SLSProjectName\\:true,\\SandboxRuntime\\:true,\\SnapshotPolicy\\:true,\\Taint\\:true,\\TerwayEniip\\:true,\\UserData\\:true},\\CloudMonitorVersion\\:\\\\,\\ClusterDomain\\:\\\\,\\ControlPlaneLogConfig\\:{\\components\\:null},\\DockerVersion\\:\\\\,\\EtcdVersion\\:\\\\,\\ExtraCertSAN\\:null,\\FreeTier\\:false,\\HasSandboxRuntime\\:false,\\IPStack\\:\\ipv4\\,\\ImageType\\:\\AliyunLinux3ContainerOptimized\\,\\KubernetesVersion\\:\\1.34.10-aliyun.1\\,\\MultiAZ\\:false,\\NameMode\\:\\\\,\\NextVersion\\:\\\\,\\OSType\\:\\Linux\\,\\Platform\\:\\AliyunLinux\\,\\PodVswitchId\\:\\\\,\\Provider\\:\\\\,\\RRSAConfig\\:{\\enabled\\:false},\\RamRoleType\\:\\\\,\\ResourceGroupId\\:\\rg-acfnvekhilw5kmy\\,\\Runtime\\:\\containerd\\,\\RuntimeVersion\\:\\2.1.9\\,\\ServiceCIDR\\:\\172.23.0.0/16\\,\\SubClass\\:\\default\\,\\SupportPlatforms\\:[\\CentOS\\,\\AliyunLinux\\,\\Windows\\,\\WindowsCore\\],\\Timezone\\:\\\\,\\VSwitchIds\\:[\\vsw-mj7qh726wq5bds252lzmj\\],\\VersionSpec\\:null,\\VpcCidr\\:\\10.0.0.0/22\\,\\ack-nvidia-device-pluginVersion\\:\\0.8.1\\,\\ack-ram-authenticatorVersion\\:\\0.5.1\\,\\ack-schedulerVersion\\:\\v1.34.0-apsara.6.11.14.80067fc3\\,\\alicloud-monitor-controllerVersion\\:\\v1.8.13\\,\\cloud-controller-managerVersion\\:\\v2.15.0\\,\\cnfs-controllerVersion\\:\\v1.2.5\\,\\corednsVersion\\:\\v1.13.2.2\\,\\csi-pluginVersion\\:\\v1.37.2\\,\\csi-provisionerVersion\\:\\v1.37.2\\,\\gateway-apiVersion\\:\\1.3.0\\,\\kube-apiserverVersion\\:\\v1.34.10-aliyun.1\\,\\kube-controller-managerVersion\\:\\v1.34.10-aliyun.1\\,\\kube-flannel-dsVersion\\:\\v0.15.1.23-33d25c1-aliyun\\,\\metrics-aggregatorVersion\\:\\v1.9.6\\,\\metrics-serverVersion\\:\\v0.3.16-ab96b23-aliyun\\,\\pod-readinessgate-webhookVersion\\:\\v1.0.0\\,\\storage-operatorVersion\\:\\v1.35.3\\}"
    },
    {
      "key": "Name",
      "value": "tbqfbinr8jv2lc7g6aeq"
    },
    {
      "key": "NetworkMode",
      "value": "vpc"
    },
    {
      "key": "Parameters",
      "value": "{ALIYUN::AccountId:5469257408566579,ALIYUN::NoValue:None,ALIYUN::Region:ap-northeast-2,ALIYUN::ResourceGroupId:rg-acfnvekhilw5kmy,ALIYUN::StackId:00623c79-cd5b-4325-bcd2-b0f0b9bf8396,ALIYUN::StackName:k8s-for-cs-c7bb68aecc47b4ac38f7c0c690d5fa579,ALIYUN::TenantId:5469257408566579,AdjustmentType:TotalCapacity,BetaVersion:,CloudMonitorFlags:False,CloudMonitorVersion:1.3.7,ClusterDns:172.23.0.10,ClusterId:c7bb68aecc47b4ac38f7c0c690d5fa579,ContainerCIDR:172.22.0.0/16,CustomK8sWorkerRole:,DisableAddons:True,DisableAutoCreateK8sWorkerRole:False,DisableAutoCreateK8sWorkerRolePolicy:True,DockerVersion:17.06.2-ce-3,ESSDeletionProtection:True,Eip:True,EipAddress:,EtcdVersion:v3.1.11,ExecuteVersion:445807839,HealthCheckType:NONE,IPStack:ipv4,ImageId:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,K8sWorkerPolicyDocument:{\\Version\\: \\1\\, \\Statement\\: [{\\Action\\: [\\ecs:DescribeInstanceAttribute\\, \\ecs:DescribeInstances\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}, {\\Action\\: [\\log:GetProject\\, \\log:GetLogStore\\, \\log:GetConfig\\, \\log:GetMachineGroup\\, \\log:GetAppliedMachineGroups\\, \\log:GetAppliedConfigs\\, \\log:GetIndex\\, \\log:GetSavedSearch\\, \\log:GetDashboard\\, \\log:GetJob\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}, {\\Action\\: [\\cr:GetAuthorizationToken\\, \\cr:ListInstanceEndpoint\\, \\cr:PullRepository\\, \\cr:GetInstanceVpcEndpoint\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}]},KeyPair:,KubernetesVersion:1.34.10-aliyun.1,MasterSLBPrivateIP:10.0.1.3,NatGateway:False,NatGatewayId:ngw-mj7mpup1ekxdr1xyxm5gj,NatGatewayType:Enhanced,NatGatewayVswitchId:,Network:Flannel,NodeNameMode:nodeip,NumOfNodes:0,OSType:Linux,Password:******,PodVswitchIds:[],ProtectedInstances:,ProxyMode:iptables,RemoveInstanceIds:,ResourceGroupId:rg-acfnvekhilw5kmy,SNatEntry:True,ScaleOutToken:None,SecurityGroupId:sg-mj78e5i8z3rmauatexhd,ServiceCIDR:172.23.0.0/16,SetUpArgs:--addon-names kube-flannel-ds,csi-plugin,csi-provisioner --labels CB-SPIDER:PMKS:CLUSTER=owned --node-cidr-mask 25 --registry-url registry-ap-northeast-2-vpc.ack.aliyuncs.com --runtime containerd --runtime-version 2.1.9,SnatTableId:stb-mj7ojrzorjio7r304ea07,Tags:[{\\Key\\: \\CB-SPIDER:PMKS:CLUSTER\\, \\Value\\: \\owned\\}, {\\Key\\: \\ack.aliyun.com\\, \\Value\\: \\c7bb68aecc47b4ac38f7c0c690d5fa579\\}],UserData:,VpcCidrWithSecondaryCidrs:[\\10.0.0.0/22\\],VpcId:vpc-mj74969ho2twcr6xzw557,WorkerAutoRenew:False,WorkerAutoRenewPeriod:1,WorkerDataDisk:False,WorkerDataDisks:[],WorkerDeletionProtection:True,WorkerDeploymentSetId:,WorkerHpcClusterId:,WorkerImageId:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,WorkerInstanceChargeType:PostPaid,WorkerInstanceTypes:ecs.n4.large,WorkerKeyPair:,WorkerLoginPassword:******,WorkerPeriod:3,WorkerPeriodUnit:Month,WorkerSnapshotPolicyId:******,WorkerSystemDiskCategory:cloud_ssd,WorkerSystemDiskPerformanceLevel:null,WorkerSystemDiskSize:40,WorkerVSwitchIds:vsw-mj7qh726wq5bds252lzmj,ZoneId:}"
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
      "value": "sg-mj78e5i8z3rmauatexhd"
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
      "value": "172.22.0.0/16"
    },
    {
      "key": "Tags",
      "value": "{key:CB-SPIDER:PMKS:CLUSTER,value:owned}; {key:ack.aliyun.com,value:c7bb68aecc47b4ac38f7c0c690d5fa579}"
    },
    {
      "key": "Updated",
      "value": "2026-09-16T19:18:32+08:00"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj74969ho2twcr6xzw557"
    },
    {
      "key": "VswitchId",
      "value": "vsw-mj7qh726wq5bds252lzmj"
    },
    {
      "key": "WorkerRamRoleName",
      "value": "KubernetesWorkerRole-00623c79-cd5b-4325-bcd2-b0f0b9bf8396"
    },
    {
      "key": "ZoneId",
      "value": "ap-northeast-2a"
    }
  ],
  "cspResourceName": "tbqfbinr8jv2lc7g6aeq",
  "cspResourceId": "c7bb68aecc47b4ac38f7c0c690d5fa579",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbqfbinr8jv2lc7g6aeq",
      "SystemId": "c7bb68aecc47b4ac38f7c0c690d5fa579"
    },
    "Version": "1.34.10-aliyun.1",
    "Network": {
      "VpcIID": {
        "NameId": "tbuiesqmm9596ealr5in",
        "SystemId": "vpc-mj74969ho2twcr6xzw557"
      },
      "SubnetIIDs": [
        {
          "NameId": "tbfponsdmhoqvn3j50pt",
          "SystemId": "vsw-mj7qh726wq5bds252lzmj"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "tblbq0l1qioj9n922n8v",
          "SystemId": "sg-mj78e5i8z3rmauatexhd"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "np30f252a975da4f39a1ed043e9f0e41a6"
        },
        "ImageIID": {
          "NameId": "aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd",
          "SystemId": "aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd"
        },
        "VMSpecName": "ecs.e-c1m2.xlarge",
        "RootDiskType": "cloud_essd",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tb9e95hm088kd3tge6mq",
          "SystemId": "tb9e95hm088kd3tge6mq"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "worker-k8s-for-cs-c7bb68aecc47b4ac38f7c0c690d5fa579",
            "SystemId": "i-mj7012hrh739bepeuwrp"
          },
          {
            "NameId": "worker-k8s-for-cs-c7bb68aecc47b4ac38f7c0c690d5fa579",
            "SystemId": "i-mj7012hrh739bepeuwro"
          }
        ],
        "KeyValueList": [
          {
            "key": "AutoScaling",
            "value": "{eip_bandwidth:0,eip_internet_charge_type:,enable:false,max_instances:0,min_instances:0,type:}"
          },
          {
            "key": "KubernetesConfig",
            "value": "{cms_enabled:false,cpu_policy:none,node_name_mode:nodeip,runtime:containerd,runtime_version:2.1.9,unschedulable:false,user_data:}"
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
            "value": "{created:2026-09-16T19:17:49.109695065+08:00,is_default:false,name:workers1,nodepool_id:np30f252a975da4f39a1ed043e9f0e41a6,region_id:ap-northeast-2,resource_group_id:,type:ess,updated:2026-09-16T19:19:37.827+08:00}"
          },
          {
            "key": "ScalingGroup",
            "value": "{auto_renew:false,auto_renew_period:0,deploymentset_id:,desired_size:2,image_id:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,image_type:AliyunLinux3ContainerOptimized,instance_charge_type:PostPaid,instance_types:[ecs.e-c1m2.xlarge],internet_charge_type:,internet_max_bandwidth_out:0,key_pair:tb9e95hm088kd3tge6mq,login_password:,multi_az_policy:BALANCE,period:0,period_unit:,platform:AliyunLinux,private_pool_options:{},ram_policy:KubernetesWorkerRole-00623c79-cd5b-4325-bcd2-b0f0b9bf8396,scaling_group_id:asg-mj73ravctixal7xont26,scaling_policy:release,security_group_id:sg-mj78e5i8z3rmauatexhd,security_group_ids:[sg-mj78e5i8z3rmauatexhd],spot_strategy:NoSpot,system_disk_categories:[cloud_essd],system_disk_category:cloud_essd,system_disk_encrypt_algorithm:,system_disk_encrypted:false,system_disk_kms_key_id:,system_disk_performance_level:,system_disk_size:100,vswitch_ids:[vsw-mj7qh726wq5bds252lzmj]}"
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
      "Endpoint": "https://43.108.35.249:6443",
      "Kubeconfig": "\napiVersion: v1\nclusters:\n- cluster:\n    server: https://43.108.35.249:6443\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUR2VENDQXFXZ0F3SUJBZ0lJQWhHbEg2SUFIb013RFFZSktvWklodmNOQVFFTEJRQXdRREVSTUE4R0ExVUUKQ2hNSWFHRnVaM3BvYjNVeEZqQVVCZ05WQkFzVERXRnNhV0poWW1FZ1kyeHZkV1F4RXpBUkJnTlZCQU1UQ210MQpZbVZ5Ym1WMFpYTXdJQmNOTWpZd09URTJNVEV3T1RBd1doZ1BNakExTmpBNU1EZ3hNVEUwTkRkYU1FQXhFVEFQCkJnTlZCQW9UQ0doaGJtZDZhRzkxTVJZd0ZBWURWUVFMRXcxaGJHbGlZV0poSUdOc2IzVmtNUk13RVFZRFZRUUQKRXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBeURNbgp2ZFRacjNXZStnU29nZy9iUXcrY3UzUlRyRU5uUEJ2cGNnQVRyMVpVeERLWHQrZ1JWc2NNSzJCQnpwQXF6cEp4CkZqS2RFb3BkWjZQWkhGeTlLRUc5TDJtQy9tR3QydkszdVV4TElnUHF6clhhTEg4d0cra3lzNUtTUFIxb3YxS3EKdHJ5Wmw3MXd4TzV6RFYyMGdZbHkyVmxPQWEyT0h0L3ZzM2ZqSXVBcWREZ3BFUVE4TVFiRHBMNVlHbUQvZGZXSgpTYjNMcW9yNDVob3hrYkNQUUd6c1Z3ZUtKcVVHVXYvMHhoMHBKamtjYkYxRGNJT1U5N0tFL09yLzJoSHNWdk1RCmpRMDF4QWF6RUpiOWZEZE0wT2JrWTdRQjRsaU8xSUZITVJ0UWZPUGViYmhqZlR3RHBCblpYYkp3SEdYV1cwR1IKdUdOci9sOUZKUFJEcm9aZzF3SURBUUFCbzRHNE1JRzFNQTRHQTFVZER3RUIvd1FFQXdJQ3JEQVBCZ05WSFJNQgpBZjhFQlRBREFRSC9NQjBHQTFVZERnUVdCQlMvdWJKMVR5OEJTRVI2b1ZWbS9UWEpVTndUdFRBOEJnZ3JCZ0VGCkJRY0JBUVF3TUM0d0xBWUlLd1lCQlFVSE1BR0dJR2gwZEhBNkx5OWpaWEowY3k1aFkzTXVZV3hwZVhWdUxtTnYKYlM5dlkzTndNRFVHQTFVZEh3UXVNQ3d3S3FBb29DYUdKR2gwZEhBNkx5OWpaWEowY3k1aFkzTXVZV3hwZVhWdQpMbU52YlM5eWIyOTBMbU55YkRBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQWFiTnpxMFVTcmZnRzFRbS9rMTVVCk5lcUptTmVJWlU1d1VTTWljUkVmTTV0MmJJNDJCZDY3dTM0OFkzMU94WjA5QXNrT1pVSGd4TU0xNWRxbFdUK0UKVUpQbUJrbklyajFTd3NuaEJLNnU4eDV3QkhZOTkvZ3IrSzNVenhNYUExU2M5bXFEVzlNby9PdmR0dVphWEE0dQpIdElmakpXV3ZUQVBBR1BXQi85Q3pjOTMxNjNwRjM3TE9tNnY2WkpNYkJIbWVaNDh4cWFFTlhENHB0QXcybElCClVZbEI1OHlTZW5Ka0lZZlNZSjVBZGZkdlNOWitwaEhKSnRWM05aUDlHbXBWVDNBVStCcVlnNFJla1pyK3VpY3MKZTcrNE5neTljOWVWUDZQd3FmY2tpT2MxRDZCQWdxc3lBd2dGSStwUkJzUVB6NXJ5TFhHUUpodlhYZlVrMXc1cgpvUT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n  name: kubernetes\ncontexts:\n- context:\n    cluster: kubernetes\n    user: \"213652363353305376\"\n  name: 213652363353305376-c7bb68aecc47b4ac38f7c0c690d5fa579\ncurrent-context: 213652363353305376-c7bb68aecc47b4ac38f7c0c690d5fa579\nkind: Config\npreferences: {}\nusers:\n- name: \"213652363353305376\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUQyVENDQXNHZ0F3SUJBZ0lJQWhHbFlZUUFIb013RFFZSktvWklodmNOQVFFTEJRQXdRREVSTUE4R0ExVUUKQ2hNSWFHRnVaM3BvYjNVeEZqQVVCZ05WQkFzVERXRnNhV0poWW1FZ1kyeHZkV1F4RXpBUkJnTlZCQU1UQ210MQpZbVZ5Ym1WMFpYTXdIaGNOTWpZd09URTJNVEV4TWpBd1doY05Namt3T1RFMU1URXhOek0yV2pCS01SVXdFd1lEClZRUUtFd3h6ZVhOMFpXMDZkWE5sY25NeENUQUhCZ05WQkFzVEFERW1NQ1FHQTFVRUF4TWRNakV6TmpVeU16WXoKTXpVek16QTFNemMyTFRFM09EazFOVGMwTlRZd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFSwpBb0lCQVFDNUs0VTR2NFp3Z3ZUYW5xSXEvbnFPSGI4NFhuU3VJZDEvak9jb1BseE9jWTJ2YmVkOE1TM2FROE1TCnh2ZVZPZlJDNTZlNGhFQVVjR294KzR1NFI3Zk0vdVFwQ2puemhLaXp1ekxMaGcrRFFXNHpYVTY1Wm5XbUtkcHgKb1ZRRnRRbkdsZGR6dVh2b1NFM3QzTzZOZCtCM21MQzJieUd5cnNJdVdIVXMyN2xwREpBdDkzd25qdlNkTkl5Rgo3VEoyOEtjeWFVeEdWRDlGNXQrNXRyV0pXMGhqL3ZqOVprWG1oODd2YUJIeGY0MDJZekhaYWV5RExQZmp4ek1JCjFibklONWpFQ2IzdGFKdXk5NEpCakdqTG5IK0FQUThKNDV5VGFYMWhDck55YVFyVjZ3TnJKdDJ3NEllU3Iwd0IKYVFUcElHZUhVdFErdzBjLzBvWndpQ3JHMWNPTkFnTUJBQUdqZ2N3d2dja3dEZ1lEVlIwUEFRSC9CQVFEQWdlQQpNQk1HQTFVZEpRUU1NQW9HQ0NzR0FRVUZCd01DTUF3R0ExVWRFd0VCL3dRQ01BQXdId1lEVlIwakJCZ3dGb0FVCnY3bXlkVTh2QVVoRWVxRlZadjAxeVZEY0U3VXdQQVlJS3dZQkJRVUhBUUVFTURBdU1Dd0dDQ3NHQVFVRkJ6QUIKaGlCb2RIUndPaTh2WTJWeWRITXVZV056TG1Gc2FYbDFiaTVqYjIwdmIyTnpjREExQmdOVkhSOEVMakFzTUNxZwpLS0FtaGlSb2RIUndPaTh2WTJWeWRITXVZV056TG1Gc2FYbDFiaTVqYjIwdmNtOXZkQzVqY213d0RRWUpLb1pJCmh2Y05BUUVMQlFBRGdnRUJBRmlmRE8rZ3huVm8zQzlDKzVFa0F3TVV5OTNZdU81NytQRUN1YjFEYWFRSnJUc04KTHlxRmc5cHpVMytRYVBhYVNVVElXeFhGdDVSOEplVkRwRzRIZHRhRVFEY2NaWkQxOUhNdm9HMDNUa256dGFoNQpXZ01xTUVMTDNyTU41TXYrMWpKbXdUMDBTVlFJeE1oRlpZTmNEc214UXhUT3JXc0grclRVUlJMY0JVV3NtZlJLCnhDQXExRGVIdmY3ZmVCUVhaWjlweEZJU0NxUjhCeWM2dnFIZ3ZmVEVDaklBZW9Hb3FWRWdXbnE2TTZGdDJaWncKNmlFeE5BRWZEcEpieE1qMmZ2SVEyV2tBQXlGVDZrVUEvRmJFNmFNeFBGR3pnMWZvVytBbUdKblhqeWdHdmhDMwpBUEE5a2F2VVIvcy8ra0NzWjNod3VrWDUyR05RRWhldGVlR1c3anc9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBdVN1Rk9MK0djSUwwMnA2aUt2NTZqaDIvT0Y1MHJpSGRmNHpuS0Q1Y1RuR05yMjNuCmZERXQya1BERXNiM2xUbjBRdWVudUlSQUZIQnFNZnVMdUVlM3pQN2tLUW81ODRTb3M3c3l5NFlQZzBGdU0xMU8KdVdaMXBpbmFjYUZVQmJVSnhwWFhjN2w3NkVoTjdkenVqWGZnZDVpd3RtOGhzcTdDTGxoMUxOdTVhUXlRTGZkOApKNDcwblRTTWhlMHlkdkNuTW1sTVJsUS9SZWJmdWJhMWlWdElZLzc0L1daRjVvZk83MmdSOFgrTk5tTXgyV25zCmd5ejM0OGN6Q05XNXlEZVl4QW05N1dpYnN2ZUNRWXhveTV4L2dEMFBDZU9jazJsOVlRcXpjbWtLMWVzRGF5YmQKc09DSGtxOU1BV2tFNlNCbmgxTFVQc05IUDlLR2NJZ3F4dFhEalFJREFRQUJBb0lCQUhOWjB5SG8zZDBWRjJkaApUdkN0bXJjUmZOK21wOVVhTDV1WHNTQjJ5Slh0VXpBbnJQN0d2Q090OXNXcHdPM3JMbWpIV3NvdFNLWHk2WlM5CmVPcVJjc1IzUExiOE1lR3JrRlIybDB3RjlYLzBzS3U2d3FIb1cvM3BNTFY4cEpUeGxHZHJUTTVBakJuRmdSWmUKZlJVUHludDJXUThzNTdjaCtzRllSWlp1UW5CWFhjeHB1UTNPSW9OVHRSQXJvN28vRXlsTlBoa01lTGUxaEdRWgpLa1RQSWFmLzNpV0lzb2hNbG40M2JNa0JYeHVISEJsL1ZXMG03am5mam5OUjVKZjJtR0tUSW9GVVM4TTMwZU1WCkt0ZW9oeUpXeTV2Wk9rNGkxSEE2aVl2Zk1FbzNUeWQvZDd6dWhRUEdFbFVPNUFzanZsSXd4Z1ZJVDdhRHVVT08KTnN6cEVGa0NnWUVBNEtIL1JkWmNVWWVzemJWY1ZaVkk1eVZrWlc5NnRvRTFHV0dzL3RrS1dDK3FWU0E5ZkdPVwp6ZmdMcFRCYkVuY0ZYSC9STndmeUJzQTluZnBUYk5TMUwrbmtpSmtMV2o0S3pRYzRUMEZFaWJJZithUjFrWEF1CjlkdnB1d2RDOU91U0ZOTUpPbFVsa3FEUGFEOWtaK3preDNyTG01a2ZwYXd4QmxHTG5KNlNwTjhDZ1lFQTB3YlcKUjQwdnhwUWVOYzRFQUt5ZFBwbVova09hR21zd0NZZFNCZHhEZ1FJeDBBL3V3bWd5YWEzRlV3Rko4NHl6RFZFWQo1Y0xqTHdsS1dTNHBEeVpoZ3lTeE5hY1hxSjlOcXhCZlJqZ0tjR25QbDhmL1FVNXByaGJ2N0lDRG1tQ2pDbmZFCndJbnVDV2p4MzNNQWV2WVcwaGFLcC94bFc5U3lmclZiVTF1ZVdSTUNnWUVBdjFINlkzYmhoaWN5VExvcy9FOTcKbWpORXdRZ2owaEZXTWxuNHg4My9lNENOYUpkZkJ2U2pXcVhxOVRTc1BKdldteVBhQXk3bk9reTlyakdvb0ViQwpVeDY0b0ljSkhrRTlYY2JLZE9ZNEE3Y01lTWxUd1IxMVFiMmQ0c2VhaFpPbUJjcUFUNGg1eWRyaEMvOEIrMm5PCnFQK3pMc0ZLTFFidGNsNDl0SlZ5ZkswQ2dZQndJMi8yOGRYemhuNVBSVHpuUnRNQWt3czhESngxY1lSRGEvOXcKVWM0bFhnOHVhMmtMTWVlb05NbHBCSVJSd2ZEY0lMNUVTajREbnNJOWhjUVg0dU5xbHpMOE9lRGVvRmpia0lXdwpnTk03VFY4ZFh4QUxtaW1hYzJIbm9adE5qQkRYM1RGV3gzVVExdzNCR2hHbFJlUjJtN2Z1OExiRUI5RWFlREVoCnFNd0t0UUtCZ0I5czJWcm1uTWdnVVpRdlZ4SHRoeGJScEZodGtMYVpyY0NKV1hIcHRwbGFONmFzZU5nQlF1WUkKMFl5aHV4c2lsNGl2cUlKdEpxMVJwbkpFYnIvTk1yNjlTSXVuZnlLdFdKTUJycUhmSEFwQWlaOGdDOGliKzBzVAp3bWJjdUM4cURuMncwcmFGalVqRGt0T0NhRnA0YVhZQ2V3Sk5UdThrSkZhS3NjNzQ2Qk8wCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==\n"
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-09-16T19:14:47+08:00",
    "KeyValueList": [
      {
        "key": "ClusterId",
        "value": "c7bb68aecc47b4ac38f7c0c690d5fa579"
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
        "value": "2026-09-16T19:14:47+08:00"
      },
      {
        "key": "CurrentVersion",
        "value": "1.34.10-aliyun.1"
      },
      {
        "key": "DeletionProtection",
        "value": "false"
      },
      {
        "key": "ExternalLoadbalancerId",
        "value": "lb-mj7fbk5ltg8uaqv90lx88"
      },
      {
        "key": "InitVersion",
        "value": "1.34.10-aliyun.1"
      },
      {
        "key": "MaintenanceWindow",
        "value": "{enable:false,weekly_period:}"
      },
      {
        "key": "MasterUrl",
        "value": "{\\api_server_endpoint\\:\\https://43.108.35.249:6443\\,\\intranet_api_server_endpoint\\:\\https://10.0.1.3:6443\\}"
      },
      {
        "key": "MetaData",
        "value": "{\\Addons\\:[{\\name\\:\\alicloud-monitor-controller\\,\\version\\:\\v1.8.13\\},{\\name\\:\\metrics-server\\,\\version\\:\\v0.3.16-ab96b23-aliyun\\},{\\name\\:\\csi-provisioner\\,\\version\\:\\v1.37.2\\},{\\name\\:\\cloud-controller-manager\\,\\version\\:\\v2.15.0\\},{\\name\\:\\csi-plugin\\,\\version\\:\\v1.37.2\\},{\\name\\:\\ack-scheduler\\,\\version\\:\\v1.34.0-apsara.6.11.14.80067fc3\\},{\\name\\:\\coredns\\,\\version\\:\\v1.13.2.2\\},{\\name\\:\\gateway-api\\,\\version\\:\\1.3.0\\},{\\name\\:\\storage-operator\\,\\version\\:\\v1.35.3\\},{\\name\\:\\managed-kube-proxy\\,\\version\\:\\v1.34.10-aliyun.1\\},{\\name\\:\\kube-flannel-ds\\,\\version\\:\\v0.15.1.23-33d25c1-aliyun\\},{\\name\\:\\ack-ram-authenticator\\,\\version\\:\\0.5.1\\},{\\name\\:\\cnfs-controller\\,\\version\\:\\v1.2.5\\,\\config\\:\\{\\\\\\CreateDefaultCNFS\\\\\\:false}\\},{\\name\\:\\ack-nvidia-device-plugin\\,\\version\\:\\0.8.1\\},{\\name\\:\\kube-apiserver\\,\\version\\:\\v1.34.10-aliyun.1\\},{\\name\\:\\pod-readinessgate-webhook\\,\\version\\:\\v1.0.0\\},{\\name\\:\\metrics-aggregator\\,\\version\\:\\v1.9.6\\},{\\name\\:\\kube-controller-manager\\,\\version\\:\\v1.34.10-aliyun.1\\}],\\AuditProjectName\\:\\\\,\\Capabilities\\:{\\AnyAZ\\:true,\\CSI\\:true,\\CpuPolicy\\:true,\\DeploymentSet\\:true,\\DisableEncryption\\:true,\\EncryptionKMSKeyId\\:\\\\,\\EnterpriseSecurityGroup\\:true,\\HpcCluster\\:true,\\IntelSGX\\:false,\\Knative\\:true,\\Network\\:\\Flannel\\,\\NgwPayByLcu\\:true,\\NodeCIDRMask\\:\\25\\,\\NodeNameMode\\:true,\\ProxyMode\\:\\\\,\\PublicSLB\\:true,\\RamRoleType\\:\\restricted\\,\\SLSProjectName\\:true,\\SandboxRuntime\\:true,\\SnapshotPolicy\\:true,\\Taint\\:true,\\TerwayEniip\\:true,\\UserData\\:true},\\CloudMonitorVersion\\:\\\\,\\ClusterDomain\\:\\\\,\\ControlPlaneLogConfig\\:{\\components\\:null},\\DockerVersion\\:\\\\,\\EtcdVersion\\:\\\\,\\ExtraCertSAN\\:null,\\FreeTier\\:false,\\HasSandboxRuntime\\:false,\\IPStack\\:\\ipv4\\,\\ImageType\\:\\AliyunLinux3ContainerOptimized\\,\\KubernetesVersion\\:\\1.34.10-aliyun.1\\,\\MultiAZ\\:false,\\NameMode\\:\\\\,\\NextVersion\\:\\\\,\\OSType\\:\\Linux\\,\\Platform\\:\\AliyunLinux\\,\\PodVswitchId\\:\\\\,\\Provider\\:\\\\,\\RRSAConfig\\:{\\enabled\\:false},\\RamRoleType\\:\\\\,\\ResourceGroupId\\:\\rg-acfnvekhilw5kmy\\,\\Runtime\\:\\containerd\\,\\RuntimeVersion\\:\\2.1.9\\,\\ServiceCIDR\\:\\172.23.0.0/16\\,\\SubClass\\:\\default\\,\\SupportPlatforms\\:[\\CentOS\\,\\AliyunLinux\\,\\Windows\\,\\WindowsCore\\],\\Timezone\\:\\\\,\\VSwitchIds\\:[\\vsw-mj7qh726wq5bds252lzmj\\],\\VersionSpec\\:null,\\VpcCidr\\:\\10.0.0.0/22\\,\\ack-nvidia-device-pluginVersion\\:\\0.8.1\\,\\ack-ram-authenticatorVersion\\:\\0.5.1\\,\\ack-schedulerVersion\\:\\v1.34.0-apsara.6.11.14.80067fc3\\,\\alicloud-monitor-controllerVersion\\:\\v1.8.13\\,\\cloud-controller-managerVersion\\:\\v2.15.0\\,\\cnfs-controllerVersion\\:\\v1.2.5\\,\\corednsVersion\\:\\v1.13.2.2\\,\\csi-pluginVersion\\:\\v1.37.2\\,\\csi-provisionerVersion\\:\\v1.37.2\\,\\gateway-apiVersion\\:\\1.3.0\\,\\kube-apiserverVersion\\:\\v1.34.10-aliyun.1\\,\\kube-controller-managerVersion\\:\\v1.34.10-aliyun.1\\,\\kube-flannel-dsVersion\\:\\v0.15.1.23-33d25c1-aliyun\\,\\metrics-aggregatorVersion\\:\\v1.9.6\\,\\metrics-serverVersion\\:\\v0.3.16-ab96b23-aliyun\\,\\pod-readinessgate-webhookVersion\\:\\v1.0.0\\,\\storage-operatorVersion\\:\\v1.35.3\\}"
      },
      {
        "key": "Name",
        "value": "tbqfbinr8jv2lc7g6aeq"
      },
      {
        "key": "NetworkMode",
        "value": "vpc"
      },
      {
        "key": "Parameters",
        "value": "{ALIYUN::AccountId:5469257408566579,ALIYUN::NoValue:None,ALIYUN::Region:ap-northeast-2,ALIYUN::ResourceGroupId:rg-acfnvekhilw5kmy,ALIYUN::StackId:00623c79-cd5b-4325-bcd2-b0f0b9bf8396,ALIYUN::StackName:k8s-for-cs-c7bb68aecc47b4ac38f7c0c690d5fa579,ALIYUN::TenantId:5469257408566579,AdjustmentType:TotalCapacity,BetaVersion:,CloudMonitorFlags:False,CloudMonitorVersion:1.3.7,ClusterDns:172.23.0.10,ClusterId:c7bb68aecc47b4ac38f7c0c690d5fa579,ContainerCIDR:172.22.0.0/16,CustomK8sWorkerRole:,DisableAddons:True,DisableAutoCreateK8sWorkerRole:False,DisableAutoCreateK8sWorkerRolePolicy:True,DockerVersion:17.06.2-ce-3,ESSDeletionProtection:True,Eip:True,EipAddress:,EtcdVersion:v3.1.11,ExecuteVersion:445807839,HealthCheckType:NONE,IPStack:ipv4,ImageId:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,K8sWorkerPolicyDocument:{\\Version\\: \\1\\, \\Statement\\: [{\\Action\\: [\\ecs:DescribeInstanceAttribute\\, \\ecs:DescribeInstances\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}, {\\Action\\: [\\log:GetProject\\, \\log:GetLogStore\\, \\log:GetConfig\\, \\log:GetMachineGroup\\, \\log:GetAppliedMachineGroups\\, \\log:GetAppliedConfigs\\, \\log:GetIndex\\, \\log:GetSavedSearch\\, \\log:GetDashboard\\, \\log:GetJob\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}, {\\Action\\: [\\cr:GetAuthorizationToken\\, \\cr:ListInstanceEndpoint\\, \\cr:PullRepository\\, \\cr:GetInstanceVpcEndpoint\\], \\Resource\\: [\\*\\], \\Effect\\: \\Allow\\}]},KeyPair:,KubernetesVersion:1.34.10-aliyun.1,MasterSLBPrivateIP:10.0.1.3,NatGateway:False,NatGatewayId:ngw-mj7mpup1ekxdr1xyxm5gj,NatGatewayType:Enhanced,NatGatewayVswitchId:,Network:Flannel,NodeNameMode:nodeip,NumOfNodes:0,OSType:Linux,Password:******,PodVswitchIds:[],ProtectedInstances:,ProxyMode:iptables,RemoveInstanceIds:,ResourceGroupId:rg-acfnvekhilw5kmy,SNatEntry:True,ScaleOutToken:None,SecurityGroupId:sg-mj78e5i8z3rmauatexhd,ServiceCIDR:172.23.0.0/16,SetUpArgs:--addon-names kube-flannel-ds,csi-plugin,csi-provisioner --labels CB-SPIDER:PMKS:CLUSTER=owned --node-cidr-mask 25 --registry-url registry-ap-northeast-2-vpc.ack.aliyuncs.com --runtime containerd --runtime-version 2.1.9,SnatTableId:stb-mj7ojrzorjio7r304ea07,Tags:[{\\Key\\: \\CB-SPIDER:PMKS:CLUSTER\\, \\Value\\: \\owned\\}, {\\Key\\: \\ack.aliyun.com\\, \\Value\\: \\c7bb68aecc47b4ac38f7c0c690d5fa579\\}],UserData:,VpcCidrWithSecondaryCidrs:[\\10.0.0.0/22\\],VpcId:vpc-mj74969ho2twcr6xzw557,WorkerAutoRenew:False,WorkerAutoRenewPeriod:1,WorkerDataDisk:False,WorkerDataDisks:[],WorkerDeletionProtection:True,WorkerDeploymentSetId:,WorkerHpcClusterId:,WorkerImageId:aliyun_3_x64_20G_container_optimized_alibase_20260720.vhd,WorkerInstanceChargeType:PostPaid,WorkerInstanceTypes:ecs.n4.large,WorkerKeyPair:,WorkerLoginPassword:******,WorkerPeriod:3,WorkerPeriodUnit:Month,WorkerSnapshotPolicyId:******,WorkerSystemDiskCategory:cloud_ssd,WorkerSystemDiskPerformanceLevel:null,WorkerSystemDiskSize:40,WorkerVSwitchIds:vsw-mj7qh726wq5bds252lzmj,ZoneId:}"
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
        "value": "sg-mj78e5i8z3rmauatexhd"
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
        "value": "172.22.0.0/16"
      },
      {
        "key": "Tags",
        "value": "{key:CB-SPIDER:PMKS:CLUSTER,value:owned}; {key:ack.aliyun.com,value:c7bb68aecc47b4ac38f7c0c690d5fa579}"
      },
      {
        "key": "Updated",
        "value": "2026-09-16T19:18:32+08:00"
      },
      {
        "key": "VpcId",
        "value": "vpc-mj74969ho2twcr6xzw557"
      },
      {
        "key": "VswitchId",
        "value": "vsw-mj7qh726wq5bds252lzmj"
      },
      {
        "key": "WorkerRamRoleName",
        "value": "KubernetesWorkerRole-00623c79-cd5b-4325-bcd2-b0f0b9bf8396"
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

