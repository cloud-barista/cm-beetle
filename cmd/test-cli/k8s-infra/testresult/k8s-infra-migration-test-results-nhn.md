# CM-Beetle K8s Infra Migration Test Results — NHNCloud-Pangyo

> [!NOTE]
> Full lifecycle against a real CSP: recommend → migrate → list → get (verified against
> the recommendation) → delete → residual resource check.

## Environment

- CSP / Region: nhn / kr1
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (e7e0c24)
- Git Commit: e7e0c24
- Namespace: mig01
- Test Date: 2026-09-29 10:39:15 KST
- Cluster ID: k8sver02-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 20ms |
| 2 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 24m30.384s |
| 3 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 2ms |
| 4 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 3ms |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 1m57.137s |
| 6 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 2m7.486s |
| 7 | Residual resource check (Tumblebug) | ✅ **PASS** | 5ms |

**Overall Result**: 7/7 steps passed ✅

**Total Duration**: 28m35s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 20ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version v1.33.4)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=nhn+kr1+m2.c4m8 nodes=2

### Step 2 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 24m30.384s
- **Status Code**: 202

- ℹ️  nameSeed: k8sver02
- ℹ️  async reqId: 1790645955801228113
- ℹ️  cluster id: k8sver02-on-prem-k8s-cluster
- ℹ️  elapsed: 24m30s
- ✅ status: Active

### Step 3 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 2ms
- **Status Code**: 200

- ✅ migrated cluster present in list (2 total)

### Step 4 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 3ms
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=nhn+kr1+m2.c4m8, nodes=2)
- ✅ version: v1.33.4 (recommended v1.33.4)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 1m57.137s

- ✅ kubeconfig obtained (server: https://f4044cc2-nks-kr1.container.nhncloud.com:6443)
- ℹ️  auth method: client certificate in kubeconfig
- ✅ API server reachable (v1.33.4)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 1)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: 125.6.38.16
- ✅ nginx served over the LoadBalancer at http://125.6.38.16/ (attempt 1)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 6 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 2m7.486s
- **Status Code**: 200

- ✅ deleted on attempt 1 (2m7s)

### Step 7 — Residual resource check (Tumblebug)

- **Duration**: 5ms

- ℹ️  VNet k8sver02-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup k8sver02-k8s-sg still exists (known gap)
- ℹ️  SshKey k8sver02-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for nhn kr1 (source: v1.32.3 → target: vv1.33.4)",
  "targetCloud": {
    "csp": "nhn",
    "region": "kr1"
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
    "connectionName": "nhn-kr1",
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
    "connectionName": "nhn-kr1",
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
      "connectionName": "nhn-kr1",
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
    "connectionName": "nhn-kr1",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "v1.33.4",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": [
      {
        "name": "workers1",
        "imageId": "default",
        "specId": "nhn+kr1+m2.c4m8",
        "rootDiskType": "default",
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
  "id": "k8sver02-on-prem-k8s-cluster",
  "uid": "tb5mmn6e88tsf1etu0v2",
  "name": "k8sver02-on-prem-k8s-cluster",
  "connectionName": "nhn-kr1",
  "connectionConfig": {
    "configName": "nhn-kr1",
    "providerName": "nhn",
    "driverName": "nhn-driver-v1.0.so",
    "credentialName": "nhn",
    "credentialHolder": "admin",
    "regionZoneInfoName": "nhn-kr1",
    "regionZoneInfo": {
      "assignedRegion": "KR1",
      "assignedZone": "kr-pub-a"
    },
    "regionDetail": {
      "regionId": "KR1",
      "regionName": "kr1",
      "description": "Pangyo (South Korea)",
      "location": {
        "display": "Pangyo (South Korea)",
        "latitude": 37.390889,
        "longitude": 127.096792
      },
      "zones": [
        "kr-pub-a",
        "kr-pub-b"
      ]
    },
    "regionRepresentative": true,
    "verified": true
  },
  "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
  "systemMessage": "",
  "label": {
    "sys.connectionName": "nhn-kr1",
    "sys.createdTime": "2026-09-29 01:40:19 +0000 UTC",
    "sys.cspResourceId": "f4044cc2-8b24-4e55-b15b-33a7637547bd",
    "sys.cspResourceName": "tb5mmn6e88tsf1etu0v2",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "k8sver02-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "k8sver02-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tb5mmn6e88tsf1etu0v2",
    "sys.version": "v1.33.4"
  },
  "systemLabel": "",
  "version": "v1.33.4",
  "network": {
    "vNetId": "k8sver02-k8s-vpc",
    "subnetIds": [
      "k8sver02-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "k8sver02-k8s-sg"
    ],
    "keyValueList": null
  },
  "k8sNodeGroupList": [
    {
      "id": "workers1",
      "name": "workers1",
      "imageId": "default",
      "specId": "nhn+kr1+m2.c4m8",
      "rootDiskType": "General HDD",
      "rootDiskSize": 100,
      "sshKeyId": "k8sver02-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 0,
      "maxNodeSize": 0,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "tb5mmn6e88tsf1etu0v2-default-worker-node-0",
          "cspResourceId": "c60fe51b-55d3-44bc-9b71-ea9056dfad4e"
        },
        {
          "cspResourceName": "tb5mmn6e88tsf1etu0v2-default-worker-node-1",
          "cspResourceId": "129061c5-e3e4-49d8-9cfc-e643b7356c09"
        }
      ],
      "keyValueList": [
        {
          "key": "ID",
          "value": "108312"
        },
        {
          "key": "UUID",
          "value": "b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d"
        },
        {
          "key": "Name",
          "value": "default-worker"
        },
        {
          "key": "ClusterID",
          "value": "f4044cc2-8b24-4e55-b15b-33a7637547bd"
        },
        {
          "key": "ProjectID",
          "value": "01d5dbda7c204a7c9872a61b5a46dd40"
        },
        {
          "key": "Labels",
          "value": "{additional_network_id_list:,additional_subnet_id_list:,allow_signal_api:false,availability_zone:kr-pub-a,boot_volume_size:100,boot_volume_type:General HDD,ca_enable:false,ca_max_node_count:0,ca_min_node_count:0,ca_pod_replicas:1,ca_scale_down_delay_after_add:10,ca_scale_down_enable:true,ca_scale_down_unneeded_time:10,ca_scale_down_util_thresh:50,cert_manager_api:True,cgroup:v2,clusterautoscale:nodegroupfeature,cni_driver:calico,cni_tag:v3.31.4,deploy_action:,external_network_id:a858742a-245b-41d3-9a05-617e1b069eb9,external_subnet_id_list:679217b7-1e03-4ce0-8007-7e5cf4f1c47c:05f97c52-a6c3-4422-8b24-31de607fed1c:de20ab10-e4dc-448d-a9a1-cc081fffed9e:5641871e-1d78-4978-a89a-63661e851cba:5cdf279e-1335-4b75-b802-0f5c75337b54:22011c23-653a-425b-b093-7bf3fb25574d:7c5aabaa-63ea-410d-8644-20185b7e31c8:be54270f-4858-413c-b448-135dc37521da:fc06aecc-b0f9-48ad-9034-62a411c8cc9a:f7038b14-a310-4eb6-92ed-a7d9551b82d7:a107a4f8-0e05-44ca-8574-c311f36a9c96:cb12a663-9938-4aa5-85aa-0945d74c4462:995add64-8bda-468e-af0f-533c00888f5a:85558a25-94b7-459d-b8ba-53740c9310db:cf714aed-b3c6-45eb-8867-08dce9f27f92:6d8e37b1-4c4b-4ed6-b34c-b2e12957d00a:d832f723-5fe6-4f55-9493-5fe9467cc259:91ab5eda-69a2-4312-8968-ec48bf7eaef0:7f6c2ade-1b11-433b-ad6c-3edefbfc038b:714c8fe2-d064-47f6-8451-9fbcc388817d:3e4a8dea-4287-4039-9f52-7317a1065723:0af1bb82-f31c-4d74-9108-f93001bc21f2:26e80586-ce00-4b45-9f23-a5f8a27485b1:4b4ed6ce-5f41-4d3c-8363-9e9928cdf51e:697edd38-a62e-4ed2-a225-e27d0a40ce89:e70a4a7f-d58d-497e-9161-ec75415c0724:170c7e54-3710-4889-a726-f8fa868b191d:e4375237-baba-4a41-a928-b4693ba0e7a3:d1e765e5-3f2e-49cf-bf9c-c6cbe5a2e496:d3e22864-b278-4923-9af5-d4f07d2ee0fb:7e486769-a377-46aa-8ec3-f0e4c2dd6c30:11133113-1a2c-4a67-b06f-0b4cf9809e92:8c4e3353-266b-4d0d-b35f-997659a154bc:62733ccb-1ab7-4913-b407-62e0381b5a6b:161335a4-c65b-47a6-bd87-4205b75ce7ca:d73e0bcc-4bce-4e3d-ade8-8b041cbf8232:0c844317-a014-4965-9bc8-c5588d668698:27de04d3-8984-4430-843d-bcfd6f821fdf:24d9fcc8-5725-4da5-9aa4-ad7707b1cfed:6fe5086b-ef56-4454-8aee-fe9d169bfdbf:0bce294a-b772-4c6c-97c9-25e3c2f2ce72:e8951dff-aa5a-4619-bebe-fc919a92866b:a6f27f47-166a-4a71-b09d-ae2d3ff0aedf:42ae5884-70bf-4232-8baa-f514ac280775:a2551567-b464-424c-ab13-8dcaf1871348:0641f8ac-c7e9-43a8-9eb5-ba63d08b83e0:9217260d-35df-4c69-af35-340a721533a8:eeb4ec59-4886-4283-94d9-7192ee41bd65:716da39f-19cb-4adb-a533-8a040232e4fd:fb3d25ea-3f60-4928-b56a-ca152f8eceee,extra_security_groups:[],extra_volumes:[],kube_tag:v1.33.4,kube_version_status:LATEST,master_lb_floating_ip_enabled:True,mba_scale_in:{\\enable\\: false, \\min_node_count\\: 1, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},mba_scale_out:{\\enable\\: false, \\max_node_count\\: 10, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},node_image:107cc02d-02d8-44fd-84d6-6c316045b817,node_list:c60fe51b-55d3-44bc-9b71-ea9056dfad4e:129061c5-e3e4-49d8-9cfc-e643b7356c09,node_resize_complete_time:2026-09-29T11:00:07,platform_version:1.202608.0,pods_network_cidr:10.100.0.0/16,pods_network_subnet:24,project_domain:NORMAL,requested_node_count:2,restore_status_on_failure:,service_cluster_ip_range:10.254.0.0/16,station_id:default,strict_sg_rules:False,upgradable_kube_versions:[],upgradable_platform_version:}"
        },
        {
          "key": "Links",
          "value": "{href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/v1/clusters/f4044cc2-8b24-4e55-b15b-33a7637547bd/nodegroups/b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d,rel:self}; {href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/clusters/f4044cc2-8b24-4e55-b15b-33a7637547bd/nodegroups/b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d,rel:bookmark}"
        },
        {
          "key": "FlavorID",
          "value": "edc79d63-98c3-4b77-a2d4-482d70e6b554"
        },
        {
          "key": "ImageID",
          "value": "107cc02d-02d8-44fd-84d6-6c316045b817"
        },
        {
          "key": "NodeAddresses",
          "value": "10.0.1.103; 10.0.1.58"
        },
        {
          "key": "NodeCount",
          "value": "2"
        },
        {
          "key": "Role",
          "value": "worker"
        },
        {
          "key": "MinNodeCount",
          "value": "1"
        },
        {
          "key": "IsDefault",
          "value": "false"
        },
        {
          "key": "StackID",
          "value": "c303476f-6250-494b-ab7b-7ebaf4b5e38e"
        },
        {
          "key": "Status",
          "value": "CREATE_COMPLETE"
        },
        {
          "key": "StatusReason",
          "value": "Deploy completed successfully"
        },
        {
          "key": "CreatedAt",
          "value": "2026-09-29T01:40:20Z"
        },
        {
          "key": "UpdatedAt",
          "value": "2026-09-29T02:01:09Z"
        }
      ],
      "cspResourceName": "workers1",
      "cspResourceId": "b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d"
        },
        "ImageIID": {
          "NameId": "107cc02d-02d8-44fd-84d6-6c316045b817",
          "SystemId": "107cc02d-02d8-44fd-84d6-6c316045b817"
        },
        "VMSpecName": "m2.c4m8",
        "RootDiskType": "General HDD",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tbmb4dvq13rrcm5o9af8",
          "SystemId": "tbmb4dvq13rrcm5o9af8"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "tb5mmn6e88tsf1etu0v2-default-worker-node-0",
            "SystemId": "c60fe51b-55d3-44bc-9b71-ea9056dfad4e"
          },
          {
            "NameId": "tb5mmn6e88tsf1etu0v2-default-worker-node-1",
            "SystemId": "129061c5-e3e4-49d8-9cfc-e643b7356c09"
          }
        ],
        "KeyValueList": [
          {
            "key": "ID",
            "value": "108312"
          },
          {
            "key": "UUID",
            "value": "b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d"
          },
          {
            "key": "Name",
            "value": "default-worker"
          },
          {
            "key": "ClusterID",
            "value": "f4044cc2-8b24-4e55-b15b-33a7637547bd"
          },
          {
            "key": "ProjectID",
            "value": "01d5dbda7c204a7c9872a61b5a46dd40"
          },
          {
            "key": "Labels",
            "value": "{additional_network_id_list:,additional_subnet_id_list:,allow_signal_api:false,availability_zone:kr-pub-a,boot_volume_size:100,boot_volume_type:General HDD,ca_enable:false,ca_max_node_count:0,ca_min_node_count:0,ca_pod_replicas:1,ca_scale_down_delay_after_add:10,ca_scale_down_enable:true,ca_scale_down_unneeded_time:10,ca_scale_down_util_thresh:50,cert_manager_api:True,cgroup:v2,clusterautoscale:nodegroupfeature,cni_driver:calico,cni_tag:v3.31.4,deploy_action:,external_network_id:a858742a-245b-41d3-9a05-617e1b069eb9,external_subnet_id_list:679217b7-1e03-4ce0-8007-7e5cf4f1c47c:05f97c52-a6c3-4422-8b24-31de607fed1c:de20ab10-e4dc-448d-a9a1-cc081fffed9e:5641871e-1d78-4978-a89a-63661e851cba:5cdf279e-1335-4b75-b802-0f5c75337b54:22011c23-653a-425b-b093-7bf3fb25574d:7c5aabaa-63ea-410d-8644-20185b7e31c8:be54270f-4858-413c-b448-135dc37521da:fc06aecc-b0f9-48ad-9034-62a411c8cc9a:f7038b14-a310-4eb6-92ed-a7d9551b82d7:a107a4f8-0e05-44ca-8574-c311f36a9c96:cb12a663-9938-4aa5-85aa-0945d74c4462:995add64-8bda-468e-af0f-533c00888f5a:85558a25-94b7-459d-b8ba-53740c9310db:cf714aed-b3c6-45eb-8867-08dce9f27f92:6d8e37b1-4c4b-4ed6-b34c-b2e12957d00a:d832f723-5fe6-4f55-9493-5fe9467cc259:91ab5eda-69a2-4312-8968-ec48bf7eaef0:7f6c2ade-1b11-433b-ad6c-3edefbfc038b:714c8fe2-d064-47f6-8451-9fbcc388817d:3e4a8dea-4287-4039-9f52-7317a1065723:0af1bb82-f31c-4d74-9108-f93001bc21f2:26e80586-ce00-4b45-9f23-a5f8a27485b1:4b4ed6ce-5f41-4d3c-8363-9e9928cdf51e:697edd38-a62e-4ed2-a225-e27d0a40ce89:e70a4a7f-d58d-497e-9161-ec75415c0724:170c7e54-3710-4889-a726-f8fa868b191d:e4375237-baba-4a41-a928-b4693ba0e7a3:d1e765e5-3f2e-49cf-bf9c-c6cbe5a2e496:d3e22864-b278-4923-9af5-d4f07d2ee0fb:7e486769-a377-46aa-8ec3-f0e4c2dd6c30:11133113-1a2c-4a67-b06f-0b4cf9809e92:8c4e3353-266b-4d0d-b35f-997659a154bc:62733ccb-1ab7-4913-b407-62e0381b5a6b:161335a4-c65b-47a6-bd87-4205b75ce7ca:d73e0bcc-4bce-4e3d-ade8-8b041cbf8232:0c844317-a014-4965-9bc8-c5588d668698:27de04d3-8984-4430-843d-bcfd6f821fdf:24d9fcc8-5725-4da5-9aa4-ad7707b1cfed:6fe5086b-ef56-4454-8aee-fe9d169bfdbf:0bce294a-b772-4c6c-97c9-25e3c2f2ce72:e8951dff-aa5a-4619-bebe-fc919a92866b:a6f27f47-166a-4a71-b09d-ae2d3ff0aedf:42ae5884-70bf-4232-8baa-f514ac280775:a2551567-b464-424c-ab13-8dcaf1871348:0641f8ac-c7e9-43a8-9eb5-ba63d08b83e0:9217260d-35df-4c69-af35-340a721533a8:eeb4ec59-4886-4283-94d9-7192ee41bd65:716da39f-19cb-4adb-a533-8a040232e4fd:fb3d25ea-3f60-4928-b56a-ca152f8eceee,extra_security_groups:[],extra_volumes:[],kube_tag:v1.33.4,kube_version_status:LATEST,master_lb_floating_ip_enabled:True,mba_scale_in:{\\enable\\: false, \\min_node_count\\: 1, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},mba_scale_out:{\\enable\\: false, \\max_node_count\\: 10, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},node_image:107cc02d-02d8-44fd-84d6-6c316045b817,node_list:c60fe51b-55d3-44bc-9b71-ea9056dfad4e:129061c5-e3e4-49d8-9cfc-e643b7356c09,node_resize_complete_time:2026-09-29T11:00:07,platform_version:1.202608.0,pods_network_cidr:10.100.0.0/16,pods_network_subnet:24,project_domain:NORMAL,requested_node_count:2,restore_status_on_failure:,service_cluster_ip_range:10.254.0.0/16,station_id:default,strict_sg_rules:False,upgradable_kube_versions:[],upgradable_platform_version:}"
          },
          {
            "key": "Links",
            "value": "{href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/v1/clusters/f4044cc2-8b24-4e55-b15b-33a7637547bd/nodegroups/b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d,rel:self}; {href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/clusters/f4044cc2-8b24-4e55-b15b-33a7637547bd/nodegroups/b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d,rel:bookmark}"
          },
          {
            "key": "FlavorID",
            "value": "edc79d63-98c3-4b77-a2d4-482d70e6b554"
          },
          {
            "key": "ImageID",
            "value": "107cc02d-02d8-44fd-84d6-6c316045b817"
          },
          {
            "key": "NodeAddresses",
            "value": "10.0.1.103; 10.0.1.58"
          },
          {
            "key": "NodeCount",
            "value": "2"
          },
          {
            "key": "Role",
            "value": "worker"
          },
          {
            "key": "MinNodeCount",
            "value": "1"
          },
          {
            "key": "IsDefault",
            "value": "false"
          },
          {
            "key": "StackID",
            "value": "c303476f-6250-494b-ab7b-7ebaf4b5e38e"
          },
          {
            "key": "Status",
            "value": "CREATE_COMPLETE"
          },
          {
            "key": "StatusReason",
            "value": "Deploy completed successfully"
          },
          {
            "key": "CreatedAt",
            "value": "2026-09-29T01:40:20Z"
          },
          {
            "key": "UpdatedAt",
            "value": "2026-09-29T02:01:09Z"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://f4044cc2-nks-kr1.container.nhncloud.com:6443",
    "kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURMekNDQWhlZ0F3SUJBZ0lSQU41RXArK3JPRTdpbWorSnRXVDM0VkV3RFFZSktvWklodmNOQVFFTEJRQXcKSHpFZE1Cc0dBMVVFQXd3VWRHSTFiVzF1Tm1VNE9IUnpaakZsZEhVd2RqSXdIaGNOTWpZd09USTRNREUwTURJdwpXaGNOTXpFd09USTRNREUwTURJd1dqQWZNUjB3R3dZRFZRUUREQlIwWWpWdGJXNDJaVGc0ZEhObU1XVjBkVEIyCk1qQ0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQU1ZSlNYZ1hDaUd4ZzJ2QjhwbHgKRHRuS0EzMzRjYXlvVHU5OWJXMDNSUm5FcUsxc3MwZWIwTFJvRHdUVDJwT0tvV1pxMzRoemVQNUpGV3BJbUg5NQpjMzNPdXloT0RXME80NHMrOGpzQkJvSi81TlJ4Y3IzVkRkdi9oaEZyUWZTK1lTWVF0QVFWTkl4c0VnT3N2ZUlNCmIzQW1Vd0xjS1ludUtDaytKOVBOY3dFZlhlakcvWjA5MTR4UFQ1ck84b0Jab3dqd2ZUVEdEV3JqRmdhRmNHNngKMVNKZFgzSTZRc0Q0L2NsMFY2bFllYXdXYkhwaUNnUzNSbVFkbDdHUlFFUkQ0ZjZ0WUFQTllKL0tPTnVSRnhwMgo2R0RqZ1NHOG8yVXZZZzFwSXc3K0h3Q2dvWndWK2s0MWJiOTRFUUdBdmgveFMvYWYremxxdWdtMjBucFBZSTF2CmtGRUNBd0VBQWFObU1HUXdFZ1lEVlIwVEFRSC9CQWd3QmdFQi93SUJBREFPQmdOVkhROEJBZjhFQkFNQ0FnUXcKSFFZRFZSME9CQllFRkdsK1Q1MmRLMHd5MzVVZ0FVSmVUUG9LT2RqZ01COEdBMVVkSXdRWU1CYUFGR2wrVDUyZApLMHd5MzVVZ0FVSmVUUG9LT2RqZ01BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQkkzQUs5bFNRS0JkU3paUkJuCmVmRGVpclI4UE40eVR1TWFYUkdKZE5xanVCQWtvR1Z0RmJ0aGpHbGp1QUI2dERoalp4dkVJcHhudVlrYWZ1UkwKTXJkc2V2QzN1RUtML01Cejc1MW9OVDhwM2NsSGtob0IydGpWWmRyOWgrUXR0WGRNc1ZaUGJXWVhWTHlDbDJHRwpaRnRiZXVZQ1c0SDVqb2tIMGRUT3hVSjB1TzBWUEU4cWlYQndrTWVRYUliV3hVMHZDZDdhQ1p3VGFPK0Z1QncxCiszaUF3V2ZqTVJLUDVreHFnNkwxdit3cXFob3lYK0ZxL0lZRFRJK2JPTzlrelJQV2JGSEdBc3o1WkphdnV2NDEKQWxXdldVdVkrYXdEWUVyRnhrdi9UYyt5a1FhY2hMVU9McWplcDBaR3BMQ2p1WGpuSUx6bUV0bjVta0pIRHdaawpkNHMrCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=\n    server: https://f4044cc2-nks-kr1.container.nhncloud.com:6443\n  name: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\ncontexts:\n- context:\n    cluster: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\n    user: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\n  name: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\ncurrent-context: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\nkind: Config\npreferences: {}\nusers:\n- name: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGVENDQWYyZ0F3SUJBZ0lSQUpmdHd3d3B5MHZJdDQ1UGFyQlJONEF3RFFZSktvWklodmNOQVFFTEJRQXcKSHpFZE1Cc0dBMVVFQXd3VWRHSTFiVzF1Tm1VNE9IUnpaakZsZEhVd2RqSXdIaGNOTWpZd09USTRNREl3TVRJMwpXaGNOTXpFd09USTRNREl3TVRJM1dqQXBNUTR3REFZRFZRUUREQVZoWkcxcGJqRVhNQlVHQTFVRUNnd09jM2x6CmRHVnRPbTFoYzNSbGNuTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFEYXBia08KNnpJTEs4RTFLcVg1SEkvRUI1WjV5RnZBNWw1TVMzK25qa0FFcXRPSXpvOWN4RUk4ZTBmMjcrSkFpRWdseG1CawpnRkFUZzMwYm9sYW5Ea3kyM05OeUc4OFB5dmovZVMvNDhxRHVCUVNQdzI2WWdEdjFCU3dpS2JzUm1LR1kzaTI2CkFHNTBVa1Y5R0ZzZW9GTHppRTFoSHlJdTlYeHpiNDVITjlHbkQ5VWg3Wm1hNzNUZVFoWU1HZStpVWo3M1lFWVgKZktDaUtTTHBkSnYwZ1N6cG1TbHJ2ZCtIdFVSSDZqMHdjNnY2eFRzVEo1VU5HL3kvaDB4MFl4cjdlRVQ1bVhZMAp4QnNXNVRUOStaRlVmdVJWRU41TWRjbnB3RzJrbVQ4Z0R3bis2aEhHLzQ5dEl2Slo0WXBCUzF5S2U3ZDEwYUpFCnlrZEtueTJNYUp1SW1YZ2RBZ01CQUFHalFqQkFNQjBHQTFVZERnUVdCQlJORldHaUhZVWg0cG8yQ21ZVFYxR3MKWGhNakVqQWZCZ05WSFNNRUdEQVdnQlJwZmsrZG5TdE1NdCtWSUFGQ1hrejZDam5ZNERBTkJna3Foa2lHOXcwQgpBUXNGQUFPQ0FRRUFMaXMvclRCUzdNa0cyYUpQVzN3RHo3Yi8xQmNqcFpyeGJKRzVvTEJEZFJ1RFd1OGxWT3RDCkhEUUNPbVJpQ1hPeGVFcmNacXFLTnBONE5QTmpvazNSSEU3ZTFJREhDTUtuR05ZdEVMVmkxbkZGQWNMMWdMUncKemU2cVJ0TS95OFczVGxCVndlQkRPaDcydjhkY0NWRy8wdlJndElrdVNrUXprSVd1UDJJOEhJNVRIZlU3R3dKTwo5YVROcjZlcGp0NUhKeG1lY1I2OG9ycTFPVDdpampwdERYOW0rOUU3VGwzWlVCSVYvcXBLVGpYRnBVZk9mcmN3CnJtaDgwODNRdFZTUUhCdDcxeGxhZ3hUVEl4bm95U0NxQzd6ODJFeTJEUnV5ZTJrQ3VWVUdkV2JxRFJQQ0NFV3gKZm9nV0UwRXVoKzJ0cUdqSnlkN2RKQkFOZFl6eGZYTUIzQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBMnFXNUR1c3lDeXZCTlNxbCtSeVB4QWVXZWNoYndPWmVURXQvcDQ1QUJLclRpTTZQClhNUkNQSHRIOXUvaVFJaElKY1pnWklCUUU0TjlHNkpXcHc1TXR0elRjaHZQRDhyNC8za3YrUEtnN2dVRWo4TnUKbUlBNzlRVXNJaW03RVppaG1ONHR1Z0J1ZEZKRmZSaGJIcUJTODRoTllSOGlMdlY4YzIrT1J6ZlJwdy9WSWUyWgptdTkwM2tJV0RCbnZvbEkrOTJCR0YzeWdvaWtpNlhTYjlJRXM2WmtwYTczZmg3VkVSK285TUhPcitzVTdFeWVWCkRSdjh2NGRNZEdNYSszaEUrWmwyTk1RYkZ1VTAvZm1SVkg3a1ZSRGVUSFhKNmNCdHBKay9JQThKL3VvUnh2K1AKYlNMeVdlR0tRVXRjaW51M2RkR2lSTXBIU3A4dGpHaWJpSmw0SFFJREFRQUJBb0lCQVFDd3pRVWhmU25RUXFkTwoySmV4SWxsV1NGUnpVWUp3TDFmZEZjZTVzNXNzcXYyMlNHRkF3Q3BYUWREbGF3Qm04a3gweno2dXhkcjZqSDZqCjA5ZUI2bHc2R2NLWktNZDhtOEpRd3F2NkFDZ0ZqK1VxWXZ1Uk1WQktSczV6S1k5dElTQzZ2aDMzbzlXdEZjRysKNyt6dWpQSEduMWNDeSt3V1VNYzdpTjloMDA4aWFIcEwzZmNnZVc2RjdJZ2ZPWVduOGNXOE80RWQxZ3lPdDhZSgpuZ2JsZE1OSUFURTNXMlZLVnp6WnVacU9ZSWVUL1ZiWjliT3RRVS9hWXMrZ3hkd0ZERjJUU2NhZHZMdzl4V3g0CjFzUG5RZEpCV0lrWWpBeW1idnVFc1VVcFRNc1Rjcm9XYkt6akpXZWtnTG1XNEZJTjU4NHMxWTRTVVN6NTRFRDMKY2pXWDVhMUJBb0dCQVBMNXZiU2R2MHNQM2RYcEZHd0JNeWxESU9MZE5zaVpqTWwzclVteFZJL28yTFk1ODVJdgowS2FyZDMwSXJmNXhrN2VOZ1NFSm9FRkFRaVFNQXoxMWRlRWV4SXptTElWakxTUFMvWTBiUGxmdXprUzcyOHo4CmFYL1hBN0lXZmVxeC9PSHRSY013L3hCR2FPM1MzRFVDbFNvTVRGNjAveDhMMkF2UE5CMHF2dHhKQW9HQkFPWmUKSXE2MWNIbVREaStZK0hLVzBMM1FUR3lRVk9LSUxlQmhKcUVYeEZWZWZDMFdsS0VvYWZ5TXZJc25Fa0ZCM2pWZgpzb0Mydk13TWVDYXBOdytranBxZUliMWw0U0RFVExvWTh6MkVIRTBMSmc3SVFYYThrZnJsVzNVTGtTaTZtcjl4CmdQanJpdktuMUxhaW5jSEJma2tRV0xOcHhRa0dJSEE0ZlVtc0pQVTFBb0dBUUU4OW1NcVAwUXc0Q09BU0dhd1AKb0lJMStCWFk1Q3RRQ2hyMDhLWlEzVzRodmNtRTRGSnJoVkdvNUowaGdGRUxhZSs0RjhoMmRBN1A4cjZETlFjYgoxaVBRbmdKbUVqLzN1SjJsb20xdGlOU2FIN01oTUJZMnpqRll0eEFnNzdlQVdVUDF6UDN3NUp2ZU5lUXppSXhRCmNycWlsQWFQNStXNG54ZU9rWkc0eHBFQ2dZQW1YdCtnQWhDdDcxU1prUDB3K1BYajUrSVM0eWVBWS9aZ1BVNVYKM3NPUkJKL2lVclNHODFoVC9JMGJFSEwxODZhemRURWlSMDNESHdDVVQvTWY0K1RzMUJJQ25nbVZqNXpJRW9mUgpZMFBqZ1V2aGduR0UrWHZITXBTOU5pUURpTEZsMmQ0Rm1CWVl2T090V0FDMjJTZlR1NmxLbVA5OHRVeUo1SjdaCnVwYWRVUUtCZ1FEYS9EVm5ORFBVWmlxYzBZSmg2MHdPakZrcjE2WkllRVBjc0JrRTMxS0pkUXNDRE55VC9VUUcKMytwUjhVRm9rU0ZQZ3I4bGZRU0xqRWZDV2lNZ2tYT21UZjluQXRJKzZyR2dNMXhvTDNML1FjRlhCUGt5cWJqZApPWmhDQnVJUGNjeWtyeUZwUnhYZVNkU0RoTWJDQ3R5VXZGajByazlKOEo2ZjJ2WkFJOFpicVE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-09-29T01:40:19Z",
  "keyValueList": [
    {
      "key": "APIAddress",
      "value": "https://f4044cc2-nks-kr1.container.nhncloud.com:6443"
    },
    {
      "key": "COEVersion",
      "value": "v1.33.4"
    },
    {
      "key": "ContainerVersion",
      "value": "1.12.6"
    },
    {
      "key": "CreateTimeout",
      "value": "60"
    },
    {
      "key": "CreatedAt",
      "value": "2026-09-29T01:40:19Z"
    },
    {
      "key": "DockerVolumeSize",
      "value": "0"
    },
    {
      "key": "FlavorID",
      "value": "edc79d63-98c3-4b77-a2d4-482d70e6b554"
    },
    {
      "key": "KeyPair",
      "value": "tbmb4dvq13rrcm5o9af8"
    },
    {
      "key": "Labels",
      "value": "{additional_network_id_list:,additional_subnet_id_list:,api_ep_ipacl_enable:false,availability_zone:kr-pub-a,cert_manager_api:True,certificate_expiry:2031-09-28T01:40:19+00:00,cni_driver:calico,cni_tag:v3.31.4,default_security_group:270fbcc1-d439-48ab-8f1a-17b989ef0564,external_network_id:a858742a-245b-41d3-9a05-617e1b069eb9,external_subnet_id_list:679217b7-1e03-4ce0-8007-7e5cf4f1c47c:05f97c52-a6c3-4422-8b24-31de607fed1c:de20ab10-e4dc-448d-a9a1-cc081fffed9e:5641871e-1d78-4978-a89a-63661e851cba:5cdf279e-1335-4b75-b802-0f5c75337b54:22011c23-653a-425b-b093-7bf3fb25574d:7c5aabaa-63ea-410d-8644-20185b7e31c8:be54270f-4858-413c-b448-135dc37521da:fc06aecc-b0f9-48ad-9034-62a411c8cc9a:f7038b14-a310-4eb6-92ed-a7d9551b82d7:a107a4f8-0e05-44ca-8574-c311f36a9c96:cb12a663-9938-4aa5-85aa-0945d74c4462:995add64-8bda-468e-af0f-533c00888f5a:85558a25-94b7-459d-b8ba-53740c9310db:cf714aed-b3c6-45eb-8867-08dce9f27f92:6d8e37b1-4c4b-4ed6-b34c-b2e12957d00a:d832f723-5fe6-4f55-9493-5fe9467cc259:91ab5eda-69a2-4312-8968-ec48bf7eaef0:7f6c2ade-1b11-433b-ad6c-3edefbfc038b:714c8fe2-d064-47f6-8451-9fbcc388817d:3e4a8dea-4287-4039-9f52-7317a1065723:0af1bb82-f31c-4d74-9108-f93001bc21f2:26e80586-ce00-4b45-9f23-a5f8a27485b1:4b4ed6ce-5f41-4d3c-8363-9e9928cdf51e:697edd38-a62e-4ed2-a225-e27d0a40ce89:e70a4a7f-d58d-497e-9161-ec75415c0724:170c7e54-3710-4889-a726-f8fa868b191d:e4375237-baba-4a41-a928-b4693ba0e7a3:d1e765e5-3f2e-49cf-bf9c-c6cbe5a2e496:d3e22864-b278-4923-9af5-d4f07d2ee0fb:7e486769-a377-46aa-8ec3-f0e4c2dd6c30:11133113-1a2c-4a67-b06f-0b4cf9809e92:8c4e3353-266b-4d0d-b35f-997659a154bc:62733ccb-1ab7-4913-b407-62e0381b5a6b:161335a4-c65b-47a6-bd87-4205b75ce7ca:d73e0bcc-4bce-4e3d-ade8-8b041cbf8232:0c844317-a014-4965-9bc8-c5588d668698:27de04d3-8984-4430-843d-bcfd6f821fdf:24d9fcc8-5725-4da5-9aa4-ad7707b1cfed:6fe5086b-ef56-4454-8aee-fe9d169bfdbf:0bce294a-b772-4c6c-97c9-25e3c2f2ce72:e8951dff-aa5a-4619-bebe-fc919a92866b:a6f27f47-166a-4a71-b09d-ae2d3ff0aedf:42ae5884-70bf-4232-8baa-f514ac280775:a2551567-b464-424c-ab13-8dcaf1871348:0641f8ac-c7e9-43a8-9eb5-ba63d08b83e0:9217260d-35df-4c69-af35-340a721533a8:eeb4ec59-4886-4283-94d9-7192ee41bd65:716da39f-19cb-4adb-a533-8a040232e4fd:fb3d25ea-3f60-4928-b56a-ca152f8eceee,k8s_args:{\\kube-apiserver/default-not-ready-toleration-seconds\\: 300, \\kube-apiserver/default-unreachable-toleration-seconds\\: 300, \\kube-controller-manager/node-monitor-grace-period\\: 40, \\kube-controller-manager/unhealthy-zone-threshold\\: 55},kube_tag:v1.33.4,kube_version_status:NEED_K8S_UPGRADE,master_lb_floating_ip_enabled:True,master_metric_agent:telegraf,nks_registry_url:dfe965c3-kr1-registry.container.nhncloud.com/container_service,node_image:107cc02d-02d8-44fd-84d6-6c316045b817,platform_version:1.202608.0,pods_network_cidr:10.100.0.0/16,pods_network_subnet:24,project_domain:NORMAL,service_cluster_ip_range:10.254.0.0/16,service_user_enabled:True,station_id:default,strict_sg_rules:False,term_of_validity:5,upgradable_kube_versions:[\\v1.34.3\\],upgradable_platform_version:}"
    },
    {
      "key": "Links",
      "value": "{href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/v1/clusters/f4044cc2-8b24-4e55-b15b-33a7637547bd,rel:self}; {href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/clusters/f4044cc2-8b24-4e55-b15b-33a7637547bd,rel:bookmark}"
    },
    {
      "key": "MasterCount",
      "value": "3"
    },
    {
      "key": "Name",
      "value": "tb5mmn6e88tsf1etu0v2"
    },
    {
      "key": "NodeAddresses",
      "value": "10.0.1.103; 10.0.1.58"
    },
    {
      "key": "NodeCount",
      "value": "2"
    },
    {
      "key": "ProjectID",
      "value": "01d5dbda7c204a7c9872a61b5a46dd40"
    },
    {
      "key": "StackID",
      "value": "4b47e8ef-1ca9-4e6a-968a-ab4f44b3a3c7"
    },
    {
      "key": "Status",
      "value": "CREATE_COMPLETE"
    },
    {
      "key": "UUID",
      "value": "f4044cc2-8b24-4e55-b15b-33a7637547bd"
    },
    {
      "key": "UpdatedAt",
      "value": "2026-09-29T02:01:24Z"
    },
    {
      "key": "UserID",
      "value": "3d980026a565497fa944b7a1d0466713"
    },
    {
      "key": "FloatingIPEnabled",
      "value": "false"
    },
    {
      "key": "FixedNetwork",
      "value": "f664f619-77a2-4aa4-abbd-c3692306b3c8"
    },
    {
      "key": "FixedSubnet",
      "value": "366b79b0-8d6d-4717-8b34-511c5779f6c6"
    },
    {
      "key": "HealthStatus",
      "value": "FRESH"
    },
    {
      "key": "HealthStatusReason",
      "value": "{api:OK,cluster.api_status:NORMAL,cluster.node_status:NORMAL,nodegroup-stats.default-worker:2:0,nodegroup.node_status.default-worker:NORMAL,timestamp:2026-09-29T11:01:23.622472}"
    }
  ],
  "cspResourceName": "tb5mmn6e88tsf1etu0v2",
  "cspResourceId": "f4044cc2-8b24-4e55-b15b-33a7637547bd",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tb5mmn6e88tsf1etu0v2",
      "SystemId": "f4044cc2-8b24-4e55-b15b-33a7637547bd"
    },
    "Version": "v1.33.4",
    "Network": {
      "VpcIID": {
        "NameId": "tb72m59o3e5sn4slhsra",
        "SystemId": "f664f619-77a2-4aa4-abbd-c3692306b3c8"
      },
      "SubnetIIDs": [
        {
          "NameId": "tbe6dr1qd27kh20pr7ch",
          "SystemId": "366b79b0-8d6d-4717-8b34-511c5779f6c6"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "#270fbcc1-d439-48ab-8f1a-17b989ef0564",
          "SystemId": "270fbcc1-d439-48ab-8f1a-17b989ef0564"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d"
        },
        "ImageIID": {
          "NameId": "107cc02d-02d8-44fd-84d6-6c316045b817",
          "SystemId": "107cc02d-02d8-44fd-84d6-6c316045b817"
        },
        "VMSpecName": "m2.c4m8",
        "RootDiskType": "General HDD",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tbmb4dvq13rrcm5o9af8",
          "SystemId": "tbmb4dvq13rrcm5o9af8"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "tb5mmn6e88tsf1etu0v2-default-worker-node-0",
            "SystemId": "c60fe51b-55d3-44bc-9b71-ea9056dfad4e"
          },
          {
            "NameId": "tb5mmn6e88tsf1etu0v2-default-worker-node-1",
            "SystemId": "129061c5-e3e4-49d8-9cfc-e643b7356c09"
          }
        ],
        "KeyValueList": [
          {
            "key": "ID",
            "value": "108312"
          },
          {
            "key": "UUID",
            "value": "b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d"
          },
          {
            "key": "Name",
            "value": "default-worker"
          },
          {
            "key": "ClusterID",
            "value": "f4044cc2-8b24-4e55-b15b-33a7637547bd"
          },
          {
            "key": "ProjectID",
            "value": "01d5dbda7c204a7c9872a61b5a46dd40"
          },
          {
            "key": "Labels",
            "value": "{additional_network_id_list:,additional_subnet_id_list:,allow_signal_api:false,availability_zone:kr-pub-a,boot_volume_size:100,boot_volume_type:General HDD,ca_enable:false,ca_max_node_count:0,ca_min_node_count:0,ca_pod_replicas:1,ca_scale_down_delay_after_add:10,ca_scale_down_enable:true,ca_scale_down_unneeded_time:10,ca_scale_down_util_thresh:50,cert_manager_api:True,cgroup:v2,clusterautoscale:nodegroupfeature,cni_driver:calico,cni_tag:v3.31.4,deploy_action:,external_network_id:a858742a-245b-41d3-9a05-617e1b069eb9,external_subnet_id_list:679217b7-1e03-4ce0-8007-7e5cf4f1c47c:05f97c52-a6c3-4422-8b24-31de607fed1c:de20ab10-e4dc-448d-a9a1-cc081fffed9e:5641871e-1d78-4978-a89a-63661e851cba:5cdf279e-1335-4b75-b802-0f5c75337b54:22011c23-653a-425b-b093-7bf3fb25574d:7c5aabaa-63ea-410d-8644-20185b7e31c8:be54270f-4858-413c-b448-135dc37521da:fc06aecc-b0f9-48ad-9034-62a411c8cc9a:f7038b14-a310-4eb6-92ed-a7d9551b82d7:a107a4f8-0e05-44ca-8574-c311f36a9c96:cb12a663-9938-4aa5-85aa-0945d74c4462:995add64-8bda-468e-af0f-533c00888f5a:85558a25-94b7-459d-b8ba-53740c9310db:cf714aed-b3c6-45eb-8867-08dce9f27f92:6d8e37b1-4c4b-4ed6-b34c-b2e12957d00a:d832f723-5fe6-4f55-9493-5fe9467cc259:91ab5eda-69a2-4312-8968-ec48bf7eaef0:7f6c2ade-1b11-433b-ad6c-3edefbfc038b:714c8fe2-d064-47f6-8451-9fbcc388817d:3e4a8dea-4287-4039-9f52-7317a1065723:0af1bb82-f31c-4d74-9108-f93001bc21f2:26e80586-ce00-4b45-9f23-a5f8a27485b1:4b4ed6ce-5f41-4d3c-8363-9e9928cdf51e:697edd38-a62e-4ed2-a225-e27d0a40ce89:e70a4a7f-d58d-497e-9161-ec75415c0724:170c7e54-3710-4889-a726-f8fa868b191d:e4375237-baba-4a41-a928-b4693ba0e7a3:d1e765e5-3f2e-49cf-bf9c-c6cbe5a2e496:d3e22864-b278-4923-9af5-d4f07d2ee0fb:7e486769-a377-46aa-8ec3-f0e4c2dd6c30:11133113-1a2c-4a67-b06f-0b4cf9809e92:8c4e3353-266b-4d0d-b35f-997659a154bc:62733ccb-1ab7-4913-b407-62e0381b5a6b:161335a4-c65b-47a6-bd87-4205b75ce7ca:d73e0bcc-4bce-4e3d-ade8-8b041cbf8232:0c844317-a014-4965-9bc8-c5588d668698:27de04d3-8984-4430-843d-bcfd6f821fdf:24d9fcc8-5725-4da5-9aa4-ad7707b1cfed:6fe5086b-ef56-4454-8aee-fe9d169bfdbf:0bce294a-b772-4c6c-97c9-25e3c2f2ce72:e8951dff-aa5a-4619-bebe-fc919a92866b:a6f27f47-166a-4a71-b09d-ae2d3ff0aedf:42ae5884-70bf-4232-8baa-f514ac280775:a2551567-b464-424c-ab13-8dcaf1871348:0641f8ac-c7e9-43a8-9eb5-ba63d08b83e0:9217260d-35df-4c69-af35-340a721533a8:eeb4ec59-4886-4283-94d9-7192ee41bd65:716da39f-19cb-4adb-a533-8a040232e4fd:fb3d25ea-3f60-4928-b56a-ca152f8eceee,extra_security_groups:[],extra_volumes:[],kube_tag:v1.33.4,kube_version_status:LATEST,master_lb_floating_ip_enabled:True,mba_scale_in:{\\enable\\: false, \\min_node_count\\: 1, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},mba_scale_out:{\\enable\\: false, \\max_node_count\\: 10, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},node_image:107cc02d-02d8-44fd-84d6-6c316045b817,node_list:c60fe51b-55d3-44bc-9b71-ea9056dfad4e:129061c5-e3e4-49d8-9cfc-e643b7356c09,node_resize_complete_time:2026-09-29T11:00:07,platform_version:1.202608.0,pods_network_cidr:10.100.0.0/16,pods_network_subnet:24,project_domain:NORMAL,requested_node_count:2,restore_status_on_failure:,service_cluster_ip_range:10.254.0.0/16,station_id:default,strict_sg_rules:False,upgradable_kube_versions:[],upgradable_platform_version:}"
          },
          {
            "key": "Links",
            "value": "{href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/v1/clusters/f4044cc2-8b24-4e55-b15b-33a7637547bd/nodegroups/b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d,rel:self}; {href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/clusters/f4044cc2-8b24-4e55-b15b-33a7637547bd/nodegroups/b6b82757-13c3-48a3-8fa6-37fb5b2e4f6d,rel:bookmark}"
          },
          {
            "key": "FlavorID",
            "value": "edc79d63-98c3-4b77-a2d4-482d70e6b554"
          },
          {
            "key": "ImageID",
            "value": "107cc02d-02d8-44fd-84d6-6c316045b817"
          },
          {
            "key": "NodeAddresses",
            "value": "10.0.1.103; 10.0.1.58"
          },
          {
            "key": "NodeCount",
            "value": "2"
          },
          {
            "key": "Role",
            "value": "worker"
          },
          {
            "key": "MinNodeCount",
            "value": "1"
          },
          {
            "key": "IsDefault",
            "value": "false"
          },
          {
            "key": "StackID",
            "value": "c303476f-6250-494b-ab7b-7ebaf4b5e38e"
          },
          {
            "key": "Status",
            "value": "CREATE_COMPLETE"
          },
          {
            "key": "StatusReason",
            "value": "Deploy completed successfully"
          },
          {
            "key": "CreatedAt",
            "value": "2026-09-29T01:40:20Z"
          },
          {
            "key": "UpdatedAt",
            "value": "2026-09-29T02:01:09Z"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://f4044cc2-nks-kr1.container.nhncloud.com:6443",
      "Kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURMekNDQWhlZ0F3SUJBZ0lSQU41RXArK3JPRTdpbWorSnRXVDM0VkV3RFFZSktvWklodmNOQVFFTEJRQXcKSHpFZE1Cc0dBMVVFQXd3VWRHSTFiVzF1Tm1VNE9IUnpaakZsZEhVd2RqSXdIaGNOTWpZd09USTRNREUwTURJdwpXaGNOTXpFd09USTRNREUwTURJd1dqQWZNUjB3R3dZRFZRUUREQlIwWWpWdGJXNDJaVGc0ZEhObU1XVjBkVEIyCk1qQ0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQU1ZSlNYZ1hDaUd4ZzJ2QjhwbHgKRHRuS0EzMzRjYXlvVHU5OWJXMDNSUm5FcUsxc3MwZWIwTFJvRHdUVDJwT0tvV1pxMzRoemVQNUpGV3BJbUg5NQpjMzNPdXloT0RXME80NHMrOGpzQkJvSi81TlJ4Y3IzVkRkdi9oaEZyUWZTK1lTWVF0QVFWTkl4c0VnT3N2ZUlNCmIzQW1Vd0xjS1ludUtDaytKOVBOY3dFZlhlakcvWjA5MTR4UFQ1ck84b0Jab3dqd2ZUVEdEV3JqRmdhRmNHNngKMVNKZFgzSTZRc0Q0L2NsMFY2bFllYXdXYkhwaUNnUzNSbVFkbDdHUlFFUkQ0ZjZ0WUFQTllKL0tPTnVSRnhwMgo2R0RqZ1NHOG8yVXZZZzFwSXc3K0h3Q2dvWndWK2s0MWJiOTRFUUdBdmgveFMvYWYremxxdWdtMjBucFBZSTF2CmtGRUNBd0VBQWFObU1HUXdFZ1lEVlIwVEFRSC9CQWd3QmdFQi93SUJBREFPQmdOVkhROEJBZjhFQkFNQ0FnUXcKSFFZRFZSME9CQllFRkdsK1Q1MmRLMHd5MzVVZ0FVSmVUUG9LT2RqZ01COEdBMVVkSXdRWU1CYUFGR2wrVDUyZApLMHd5MzVVZ0FVSmVUUG9LT2RqZ01BMEdDU3FHU0liM0RRRUJDd1VBQTRJQkFRQkkzQUs5bFNRS0JkU3paUkJuCmVmRGVpclI4UE40eVR1TWFYUkdKZE5xanVCQWtvR1Z0RmJ0aGpHbGp1QUI2dERoalp4dkVJcHhudVlrYWZ1UkwKTXJkc2V2QzN1RUtML01Cejc1MW9OVDhwM2NsSGtob0IydGpWWmRyOWgrUXR0WGRNc1ZaUGJXWVhWTHlDbDJHRwpaRnRiZXVZQ1c0SDVqb2tIMGRUT3hVSjB1TzBWUEU4cWlYQndrTWVRYUliV3hVMHZDZDdhQ1p3VGFPK0Z1QncxCiszaUF3V2ZqTVJLUDVreHFnNkwxdit3cXFob3lYK0ZxL0lZRFRJK2JPTzlrelJQV2JGSEdBc3o1WkphdnV2NDEKQWxXdldVdVkrYXdEWUVyRnhrdi9UYyt5a1FhY2hMVU9McWplcDBaR3BMQ2p1WGpuSUx6bUV0bjVta0pIRHdaawpkNHMrCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=\n    server: https://f4044cc2-nks-kr1.container.nhncloud.com:6443\n  name: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\ncontexts:\n- context:\n    cluster: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\n    user: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\n  name: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\ncurrent-context: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\nkind: Config\npreferences: {}\nusers:\n- name: \"nks_tb5mmn6e88tsf1etu0v2_f4044cc2-8b24-4e55-b15b-33a7637547bd\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGVENDQWYyZ0F3SUJBZ0lSQUpmdHd3d3B5MHZJdDQ1UGFyQlJONEF3RFFZSktvWklodmNOQVFFTEJRQXcKSHpFZE1Cc0dBMVVFQXd3VWRHSTFiVzF1Tm1VNE9IUnpaakZsZEhVd2RqSXdIaGNOTWpZd09USTRNREl3TVRJMwpXaGNOTXpFd09USTRNREl3TVRJM1dqQXBNUTR3REFZRFZRUUREQVZoWkcxcGJqRVhNQlVHQTFVRUNnd09jM2x6CmRHVnRPbTFoYzNSbGNuTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFEYXBia08KNnpJTEs4RTFLcVg1SEkvRUI1WjV5RnZBNWw1TVMzK25qa0FFcXRPSXpvOWN4RUk4ZTBmMjcrSkFpRWdseG1CawpnRkFUZzMwYm9sYW5Ea3kyM05OeUc4OFB5dmovZVMvNDhxRHVCUVNQdzI2WWdEdjFCU3dpS2JzUm1LR1kzaTI2CkFHNTBVa1Y5R0ZzZW9GTHppRTFoSHlJdTlYeHpiNDVITjlHbkQ5VWg3Wm1hNzNUZVFoWU1HZStpVWo3M1lFWVgKZktDaUtTTHBkSnYwZ1N6cG1TbHJ2ZCtIdFVSSDZqMHdjNnY2eFRzVEo1VU5HL3kvaDB4MFl4cjdlRVQ1bVhZMAp4QnNXNVRUOStaRlVmdVJWRU41TWRjbnB3RzJrbVQ4Z0R3bis2aEhHLzQ5dEl2Slo0WXBCUzF5S2U3ZDEwYUpFCnlrZEtueTJNYUp1SW1YZ2RBZ01CQUFHalFqQkFNQjBHQTFVZERnUVdCQlJORldHaUhZVWg0cG8yQ21ZVFYxR3MKWGhNakVqQWZCZ05WSFNNRUdEQVdnQlJwZmsrZG5TdE1NdCtWSUFGQ1hrejZDam5ZNERBTkJna3Foa2lHOXcwQgpBUXNGQUFPQ0FRRUFMaXMvclRCUzdNa0cyYUpQVzN3RHo3Yi8xQmNqcFpyeGJKRzVvTEJEZFJ1RFd1OGxWT3RDCkhEUUNPbVJpQ1hPeGVFcmNacXFLTnBONE5QTmpvazNSSEU3ZTFJREhDTUtuR05ZdEVMVmkxbkZGQWNMMWdMUncKemU2cVJ0TS95OFczVGxCVndlQkRPaDcydjhkY0NWRy8wdlJndElrdVNrUXprSVd1UDJJOEhJNVRIZlU3R3dKTwo5YVROcjZlcGp0NUhKeG1lY1I2OG9ycTFPVDdpampwdERYOW0rOUU3VGwzWlVCSVYvcXBLVGpYRnBVZk9mcmN3CnJtaDgwODNRdFZTUUhCdDcxeGxhZ3hUVEl4bm95U0NxQzd6ODJFeTJEUnV5ZTJrQ3VWVUdkV2JxRFJQQ0NFV3gKZm9nV0UwRXVoKzJ0cUdqSnlkN2RKQkFOZFl6eGZYTUIzQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBMnFXNUR1c3lDeXZCTlNxbCtSeVB4QWVXZWNoYndPWmVURXQvcDQ1QUJLclRpTTZQClhNUkNQSHRIOXUvaVFJaElKY1pnWklCUUU0TjlHNkpXcHc1TXR0elRjaHZQRDhyNC8za3YrUEtnN2dVRWo4TnUKbUlBNzlRVXNJaW03RVppaG1ONHR1Z0J1ZEZKRmZSaGJIcUJTODRoTllSOGlMdlY4YzIrT1J6ZlJwdy9WSWUyWgptdTkwM2tJV0RCbnZvbEkrOTJCR0YzeWdvaWtpNlhTYjlJRXM2WmtwYTczZmg3VkVSK285TUhPcitzVTdFeWVWCkRSdjh2NGRNZEdNYSszaEUrWmwyTk1RYkZ1VTAvZm1SVkg3a1ZSRGVUSFhKNmNCdHBKay9JQThKL3VvUnh2K1AKYlNMeVdlR0tRVXRjaW51M2RkR2lSTXBIU3A4dGpHaWJpSmw0SFFJREFRQUJBb0lCQVFDd3pRVWhmU25RUXFkTwoySmV4SWxsV1NGUnpVWUp3TDFmZEZjZTVzNXNzcXYyMlNHRkF3Q3BYUWREbGF3Qm04a3gweno2dXhkcjZqSDZqCjA5ZUI2bHc2R2NLWktNZDhtOEpRd3F2NkFDZ0ZqK1VxWXZ1Uk1WQktSczV6S1k5dElTQzZ2aDMzbzlXdEZjRysKNyt6dWpQSEduMWNDeSt3V1VNYzdpTjloMDA4aWFIcEwzZmNnZVc2RjdJZ2ZPWVduOGNXOE80RWQxZ3lPdDhZSgpuZ2JsZE1OSUFURTNXMlZLVnp6WnVacU9ZSWVUL1ZiWjliT3RRVS9hWXMrZ3hkd0ZERjJUU2NhZHZMdzl4V3g0CjFzUG5RZEpCV0lrWWpBeW1idnVFc1VVcFRNc1Rjcm9XYkt6akpXZWtnTG1XNEZJTjU4NHMxWTRTVVN6NTRFRDMKY2pXWDVhMUJBb0dCQVBMNXZiU2R2MHNQM2RYcEZHd0JNeWxESU9MZE5zaVpqTWwzclVteFZJL28yTFk1ODVJdgowS2FyZDMwSXJmNXhrN2VOZ1NFSm9FRkFRaVFNQXoxMWRlRWV4SXptTElWakxTUFMvWTBiUGxmdXprUzcyOHo4CmFYL1hBN0lXZmVxeC9PSHRSY013L3hCR2FPM1MzRFVDbFNvTVRGNjAveDhMMkF2UE5CMHF2dHhKQW9HQkFPWmUKSXE2MWNIbVREaStZK0hLVzBMM1FUR3lRVk9LSUxlQmhKcUVYeEZWZWZDMFdsS0VvYWZ5TXZJc25Fa0ZCM2pWZgpzb0Mydk13TWVDYXBOdytranBxZUliMWw0U0RFVExvWTh6MkVIRTBMSmc3SVFYYThrZnJsVzNVTGtTaTZtcjl4CmdQanJpdktuMUxhaW5jSEJma2tRV0xOcHhRa0dJSEE0ZlVtc0pQVTFBb0dBUUU4OW1NcVAwUXc0Q09BU0dhd1AKb0lJMStCWFk1Q3RRQ2hyMDhLWlEzVzRodmNtRTRGSnJoVkdvNUowaGdGRUxhZSs0RjhoMmRBN1A4cjZETlFjYgoxaVBRbmdKbUVqLzN1SjJsb20xdGlOU2FIN01oTUJZMnpqRll0eEFnNzdlQVdVUDF6UDN3NUp2ZU5lUXppSXhRCmNycWlsQWFQNStXNG54ZU9rWkc0eHBFQ2dZQW1YdCtnQWhDdDcxU1prUDB3K1BYajUrSVM0eWVBWS9aZ1BVNVYKM3NPUkJKL2lVclNHODFoVC9JMGJFSEwxODZhemRURWlSMDNESHdDVVQvTWY0K1RzMUJJQ25nbVZqNXpJRW9mUgpZMFBqZ1V2aGduR0UrWHZITXBTOU5pUURpTEZsMmQ0Rm1CWVl2T090V0FDMjJTZlR1NmxLbVA5OHRVeUo1SjdaCnVwYWRVUUtCZ1FEYS9EVm5ORFBVWmlxYzBZSmg2MHdPakZrcjE2WkllRVBjc0JrRTMxS0pkUXNDRE55VC9VUUcKMytwUjhVRm9rU0ZQZ3I4bGZRU0xqRWZDV2lNZ2tYT21UZjluQXRJKzZyR2dNMXhvTDNML1FjRlhCUGt5cWJqZApPWmhDQnVJUGNjeWtyeUZwUnhYZVNkU0RoTWJDQ3R5VXZGajByazlKOEo2ZjJ2WkFJOFpicVE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-09-29T01:40:19Z",
    "KeyValueList": [
      {
        "key": "APIAddress",
        "value": "https://f4044cc2-nks-kr1.container.nhncloud.com:6443"
      },
      {
        "key": "COEVersion",
        "value": "v1.33.4"
      },
      {
        "key": "ContainerVersion",
        "value": "1.12.6"
      },
      {
        "key": "CreateTimeout",
        "value": "60"
      },
      {
        "key": "CreatedAt",
        "value": "2026-09-29T01:40:19Z"
      },
      {
        "key": "DockerVolumeSize",
        "value": "0"
      },
      {
        "key": "FlavorID",
        "value": "edc79d63-98c3-4b77-a2d4-482d70e6b554"
      },
      {
        "key": "KeyPair",
        "value": "tbmb4dvq13rrcm5o9af8"
      },
      {
        "key": "Labels",
        "value": "{additional_network_id_list:,additional_subnet_id_list:,api_ep_ipacl_enable:false,availability_zone:kr-pub-a,cert_manager_api:True,certificate_expiry:2031-09-28T01:40:19+00:00,cni_driver:calico,cni_tag:v3.31.4,default_security_group:270fbcc1-d439-48ab-8f1a-17b989ef0564,external_network_id:a858742a-245b-41d3-9a05-617e1b069eb9,external_subnet_id_list:679217b7-1e03-4ce0-8007-7e5cf4f1c47c:05f97c52-a6c3-4422-8b24-31de607fed1c:de20ab10-e4dc-448d-a9a1-cc081fffed9e:5641871e-1d78-4978-a89a-63661e851cba:5cdf279e-1335-4b75-b802-0f5c75337b54:22011c23-653a-425b-b093-7bf3fb25574d:7c5aabaa-63ea-410d-8644-20185b7e31c8:be54270f-4858-413c-b448-135dc37521da:fc06aecc-b0f9-48ad-9034-62a411c8cc9a:f7038b14-a310-4eb6-92ed-a7d9551b82d7:a107a4f8-0e05-44ca-8574-c311f36a9c96:cb12a663-9938-4aa5-85aa-0945d74c4462:995add64-8bda-468e-af0f-533c00888f5a:85558a25-94b7-459d-b8ba-53740c9310db:cf714aed-b3c6-45eb-8867-08dce9f27f92:6d8e37b1-4c4b-4ed6-b34c-b2e12957d00a:d832f723-5fe6-4f55-9493-5fe9467cc259:91ab5eda-69a2-4312-8968-ec48bf7eaef0:7f6c2ade-1b11-433b-ad6c-3edefbfc038b:714c8fe2-d064-47f6-8451-9fbcc388817d:3e4a8dea-4287-4039-9f52-7317a1065723:0af1bb82-f31c-4d74-9108-f93001bc21f2:26e80586-ce00-4b45-9f23-a5f8a27485b1:4b4ed6ce-5f41-4d3c-8363-9e9928cdf51e:697edd38-a62e-4ed2-a225-e27d0a40ce89:e70a4a7f-d58d-497e-9161-ec75415c0724:170c7e54-3710-4889-a726-f8fa868b191d:e4375237-baba-4a41-a928-b4693ba0e7a3:d1e765e5-3f2e-49cf-bf9c-c6cbe5a2e496:d3e22864-b278-4923-9af5-d4f07d2ee0fb:7e486769-a377-46aa-8ec3-f0e4c2dd6c30:11133113-1a2c-4a67-b06f-0b4cf9809e92:8c4e3353-266b-4d0d-b35f-997659a154bc:62733ccb-1ab7-4913-b407-62e0381b5a6b:161335a4-c65b-47a6-bd87-4205b75ce7ca:d73e0bcc-4bce-4e3d-ade8-8b041cbf8232:0c844317-a014-4965-9bc8-c5588d668698:27de04d3-8984-4430-843d-bcfd6f821fdf:24d9fcc8-5725-4da5-9aa4-ad7707b1cfed:6fe5086b-ef56-4454-8aee-fe9d169bfdbf:0bce294a-b772-4c6c-97c9-25e3c2f2ce72:e8951dff-aa5a-4619-bebe-fc919a92866b:a6f27f47-166a-4a71-b09d-ae2d3ff0aedf:42ae5884-70bf-4232-8baa-f514ac280775:a2551567-b464-424c-ab13-8dcaf1871348:0641f8ac-c7e9-43a8-9eb5-ba63d08b83e0:9217260d-35df-4c69-af35-340a721533a8:eeb4ec59-4886-4283-94d9-7192ee41bd65:716da39f-19cb-4adb-a533-8a040232e4fd:fb3d25ea-3f60-4928-b56a-ca152f8eceee,k8s_args:{\\kube-apiserver/default-not-ready-toleration-seconds\\: 300, \\kube-apiserver/default-unreachable-toleration-seconds\\: 300, \\kube-controller-manager/node-monitor-grace-period\\: 40, \\kube-controller-manager/unhealthy-zone-threshold\\: 55},kube_tag:v1.33.4,kube_version_status:NEED_K8S_UPGRADE,master_lb_floating_ip_enabled:True,master_metric_agent:telegraf,nks_registry_url:dfe965c3-kr1-registry.container.nhncloud.com/container_service,node_image:107cc02d-02d8-44fd-84d6-6c316045b817,platform_version:1.202608.0,pods_network_cidr:10.100.0.0/16,pods_network_subnet:24,project_domain:NORMAL,service_cluster_ip_range:10.254.0.0/16,service_user_enabled:True,station_id:default,strict_sg_rules:False,term_of_validity:5,upgradable_kube_versions:[\\v1.34.3\\],upgradable_platform_version:}"
      },
      {
        "key": "Links",
        "value": "{href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/v1/clusters/f4044cc2-8b24-4e55-b15b-33a7637547bd,rel:self}; {href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/clusters/f4044cc2-8b24-4e55-b15b-33a7637547bd,rel:bookmark}"
      },
      {
        "key": "MasterCount",
        "value": "3"
      },
      {
        "key": "Name",
        "value": "tb5mmn6e88tsf1etu0v2"
      },
      {
        "key": "NodeAddresses",
        "value": "10.0.1.103; 10.0.1.58"
      },
      {
        "key": "NodeCount",
        "value": "2"
      },
      {
        "key": "ProjectID",
        "value": "01d5dbda7c204a7c9872a61b5a46dd40"
      },
      {
        "key": "StackID",
        "value": "4b47e8ef-1ca9-4e6a-968a-ab4f44b3a3c7"
      },
      {
        "key": "Status",
        "value": "CREATE_COMPLETE"
      },
      {
        "key": "UUID",
        "value": "f4044cc2-8b24-4e55-b15b-33a7637547bd"
      },
      {
        "key": "UpdatedAt",
        "value": "2026-09-29T02:01:24Z"
      },
      {
        "key": "UserID",
        "value": "3d980026a565497fa944b7a1d0466713"
      },
      {
        "key": "FloatingIPEnabled",
        "value": "false"
      },
      {
        "key": "FixedNetwork",
        "value": "f664f619-77a2-4aa4-abbd-c3692306b3c8"
      },
      {
        "key": "FixedSubnet",
        "value": "366b79b0-8d6d-4717-8b34-511c5779f6c6"
      },
      {
        "key": "HealthStatus",
        "value": "FRESH"
      },
      {
        "key": "HealthStatusReason",
        "value": "{api:OK,cluster.api_status:NORMAL,cluster.node_status:NORMAL,nodegroup-stats.default-worker:2:0,nodegroup.node_status.default-worker:NORMAL,timestamp:2026-09-29T11:01:23.622472}"
      }
    ]
  }
}
```

</details>

