# CM-Beetle K8s Infra Migration Test Results — NHNCloud-Pangyo

> [!NOTE]
> Full lifecycle against a real CSP: recommend → validate → migrate → list → get → workload → delete → residual.

## Environment

- CSP / Region: nhn / kr1
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (cd1f3a2)
- Git Commit: cd1f3a2
- Namespace: mig01
- Test Date: 2026-10-07 16:42:15 KST
- Cluster ID: mig08-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 18ms |
| 2 | POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation) | ✅ **PASS** | 29ms |
| 3 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 20m15.361s |
| 4 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 2ms |
| 5 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 4.242s |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 1m29.741s |
| 7 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 3m11.438s |
| 8 | Residual resource check (Tumblebug) | ✅ **PASS** | 5ms |

**Overall Result**: 8/8 steps passed ✅

**Total Duration**: 25m0s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 18ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version v1.33.4)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=nhn+kr1+m2.c4m8 nodes=2

### Step 2 — POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)

- **Duration**: 29ms
- **Status Code**: 200

- ✅ Target K8s infra model is valid for migration (0 issues)

### Step 3 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 20m15.361s
- **Status Code**: 202

- ℹ️  nameSeed: mig08
- ℹ️  async reqId: 1791358935731092515
- ℹ️  cluster id: mig08-on-prem-k8s-cluster
- ℹ️  elapsed: 20m15s
- ✅ status: Active

### Step 4 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 2ms
- **Status Code**: 200

- ✅ migrated cluster present in list (4 total)

### Step 5 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 4.242s
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=nhn+kr1+m2.c4m8, nodes=2)
- ✅ version: v1.33.4 (recommended v1.33.4)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 1m29.741s

- ✅ kubeconfig obtained (server: https://3e1392d2-nks-kr1.container.nhncloud.com:6443)
- ℹ️  auth method: client certificate in kubeconfig
- ✅ API server reachable (v1.33.4)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 2)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: 133.186.218.251
- ✅ nginx served over the LoadBalancer at http://133.186.218.251/ (attempt 2)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 7 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 3m11.438s
- **Status Code**: 200

- ✅ deleted on attempt 1 (3m11s)

### Step 8 — Residual resource check (Tumblebug)

- **Duration**: 5ms

- ℹ️  VNet mig08-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig08-k8s-sg still exists (known gap)
- ℹ️  SshKey mig08-k8s-sshkey still exists (known gap)

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
  "id": "mig08-on-prem-k8s-cluster",
  "uid": "tbqarbc38u4jdhjfp8ho",
  "name": "mig08-on-prem-k8s-cluster",
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
    "sys.createdTime": "2026-10-07 07:43:17 +0000 UTC",
    "sys.cspResourceId": "3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7",
    "sys.cspResourceName": "tbqarbc38u4jdhjfp8ho",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig08-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig08-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbqarbc38u4jdhjfp8ho",
    "sys.version": "v1.33.4"
  },
  "systemLabel": "",
  "version": "v1.33.4",
  "network": {
    "vNetId": "mig08-k8s-vpc",
    "subnetIds": [
      "mig08-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "mig08-k8s-sg"
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
      "sshKeyId": "mig08-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 0,
      "maxNodeSize": 0,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "tbqarbc38u4jdhjfp8ho-default-worker-node-0",
          "cspResourceId": "8f4add36-6a8b-4db8-8854-c83515a035d3"
        },
        {
          "cspResourceName": "tbqarbc38u4jdhjfp8ho-default-worker-node-1",
          "cspResourceId": "07afb0fc-0226-449d-a7d4-bfb7ba6e0bca"
        }
      ],
      "keyValueList": [
        {
          "key": "ID",
          "value": "109030"
        },
        {
          "key": "UUID",
          "value": "0af3aadb-40b4-47ea-93f0-bd63ab43abd4"
        },
        {
          "key": "Name",
          "value": "default-worker"
        },
        {
          "key": "ClusterID",
          "value": "3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7"
        },
        {
          "key": "ProjectID",
          "value": "01d5dbda7c204a7c9872a61b5a46dd40"
        },
        {
          "key": "Labels",
          "value": "{additional_network_id_list:,additional_subnet_id_list:,allow_signal_api:false,availability_zone:kr-pub-a,boot_volume_size:100,boot_volume_type:General HDD,ca_enable:false,ca_max_node_count:0,ca_min_node_count:0,ca_pod_replicas:1,ca_scale_down_delay_after_add:10,ca_scale_down_enable:true,ca_scale_down_unneeded_time:10,ca_scale_down_util_thresh:50,cert_manager_api:True,cgroup:v2,clusterautoscale:nodegroupfeature,cni_driver:calico,cni_tag:v3.31.4,deploy_action:,external_network_id:a858742a-245b-41d3-9a05-617e1b069eb9,external_subnet_id_list:679217b7-1e03-4ce0-8007-7e5cf4f1c47c:05f97c52-a6c3-4422-8b24-31de607fed1c:de20ab10-e4dc-448d-a9a1-cc081fffed9e:5641871e-1d78-4978-a89a-63661e851cba:5cdf279e-1335-4b75-b802-0f5c75337b54:7c5aabaa-63ea-410d-8644-20185b7e31c8:f7038b14-a310-4eb6-92ed-a7d9551b82d7:22011c23-653a-425b-b093-7bf3fb25574d:fc06aecc-b0f9-48ad-9034-62a411c8cc9a:be54270f-4858-413c-b448-135dc37521da:a107a4f8-0e05-44ca-8574-c311f36a9c96:cb12a663-9938-4aa5-85aa-0945d74c4462:995add64-8bda-468e-af0f-533c00888f5a:85558a25-94b7-459d-b8ba-53740c9310db:cf714aed-b3c6-45eb-8867-08dce9f27f92:6d8e37b1-4c4b-4ed6-b34c-b2e12957d00a:d832f723-5fe6-4f55-9493-5fe9467cc259:7f6c2ade-1b11-433b-ad6c-3edefbfc038b:91ab5eda-69a2-4312-8968-ec48bf7eaef0:714c8fe2-d064-47f6-8451-9fbcc388817d:3e4a8dea-4287-4039-9f52-7317a1065723:0af1bb82-f31c-4d74-9108-f93001bc21f2:26e80586-ce00-4b45-9f23-a5f8a27485b1:4b4ed6ce-5f41-4d3c-8363-9e9928cdf51e:697edd38-a62e-4ed2-a225-e27d0a40ce89:e70a4a7f-d58d-497e-9161-ec75415c0724:170c7e54-3710-4889-a726-f8fa868b191d:e4375237-baba-4a41-a928-b4693ba0e7a3:d1e765e5-3f2e-49cf-bf9c-c6cbe5a2e496:d3e22864-b278-4923-9af5-d4f07d2ee0fb:7e486769-a377-46aa-8ec3-f0e4c2dd6c30:11133113-1a2c-4a67-b06f-0b4cf9809e92:8c4e3353-266b-4d0d-b35f-997659a154bc:62733ccb-1ab7-4913-b407-62e0381b5a6b:161335a4-c65b-47a6-bd87-4205b75ce7ca:d73e0bcc-4bce-4e3d-ade8-8b041cbf8232:0c844317-a014-4965-9bc8-c5588d668698:27de04d3-8984-4430-843d-bcfd6f821fdf:24d9fcc8-5725-4da5-9aa4-ad7707b1cfed:6fe5086b-ef56-4454-8aee-fe9d169bfdbf:0bce294a-b772-4c6c-97c9-25e3c2f2ce72:a6f27f47-166a-4a71-b09d-ae2d3ff0aedf:e8951dff-aa5a-4619-bebe-fc919a92866b:42ae5884-70bf-4232-8baa-f514ac280775:a2551567-b464-424c-ab13-8dcaf1871348:0641f8ac-c7e9-43a8-9eb5-ba63d08b83e0:9217260d-35df-4c69-af35-340a721533a8:eeb4ec59-4886-4283-94d9-7192ee41bd65:716da39f-19cb-4adb-a533-8a040232e4fd:fb3d25ea-3f60-4928-b56a-ca152f8eceee,extra_security_groups:[],extra_volumes:[],kube_tag:v1.33.4,kube_version_status:LATEST,master_lb_floating_ip_enabled:True,mba_scale_in:{\\enable\\: false, \\min_node_count\\: 1, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},mba_scale_out:{\\enable\\: false, \\max_node_count\\: 10, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},node_image:107cc02d-02d8-44fd-84d6-6c316045b817,node_list:8f4add36-6a8b-4db8-8854-c83515a035d3:07afb0fc-0226-449d-a7d4-bfb7ba6e0bca,node_resize_complete_time:2026-10-07T17:01:29,platform_version:1.202608.0,pods_network_cidr:10.100.0.0/16,pods_network_subnet:24,project_domain:NORMAL,requested_node_count:2,restore_status_on_failure:,service_cluster_ip_range:10.254.0.0/16,station_id:default,strict_sg_rules:False,upgradable_kube_versions:[],upgradable_platform_version:}"
        },
        {
          "key": "Links",
          "value": "{href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/v1/clusters/3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7/nodegroups/0af3aadb-40b4-47ea-93f0-bd63ab43abd4,rel:self}; {href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/clusters/3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7/nodegroups/0af3aadb-40b4-47ea-93f0-bd63ab43abd4,rel:bookmark}"
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
          "value": "10.0.1.104; 10.0.1.6"
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
          "value": "1a993d78-21ec-4fbd-9bae-81f045e5c101"
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
          "value": "2026-10-07T07:43:17Z"
        },
        {
          "key": "UpdatedAt",
          "value": "2026-10-07T08:01:50Z"
        }
      ],
      "cspResourceName": "workers1",
      "cspResourceId": "0af3aadb-40b4-47ea-93f0-bd63ab43abd4",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "0af3aadb-40b4-47ea-93f0-bd63ab43abd4"
        },
        "ImageIID": {
          "NameId": "107cc02d-02d8-44fd-84d6-6c316045b817",
          "SystemId": "107cc02d-02d8-44fd-84d6-6c316045b817"
        },
        "VMSpecName": "m2.c4m8",
        "RootDiskType": "General HDD",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tbtj1kicavi8btlht22d",
          "SystemId": "tbtj1kicavi8btlht22d"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "tbqarbc38u4jdhjfp8ho-default-worker-node-0",
            "SystemId": "8f4add36-6a8b-4db8-8854-c83515a035d3"
          },
          {
            "NameId": "tbqarbc38u4jdhjfp8ho-default-worker-node-1",
            "SystemId": "07afb0fc-0226-449d-a7d4-bfb7ba6e0bca"
          }
        ],
        "KeyValueList": [
          {
            "key": "ID",
            "value": "109030"
          },
          {
            "key": "UUID",
            "value": "0af3aadb-40b4-47ea-93f0-bd63ab43abd4"
          },
          {
            "key": "Name",
            "value": "default-worker"
          },
          {
            "key": "ClusterID",
            "value": "3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7"
          },
          {
            "key": "ProjectID",
            "value": "01d5dbda7c204a7c9872a61b5a46dd40"
          },
          {
            "key": "Labels",
            "value": "{additional_network_id_list:,additional_subnet_id_list:,allow_signal_api:false,availability_zone:kr-pub-a,boot_volume_size:100,boot_volume_type:General HDD,ca_enable:false,ca_max_node_count:0,ca_min_node_count:0,ca_pod_replicas:1,ca_scale_down_delay_after_add:10,ca_scale_down_enable:true,ca_scale_down_unneeded_time:10,ca_scale_down_util_thresh:50,cert_manager_api:True,cgroup:v2,clusterautoscale:nodegroupfeature,cni_driver:calico,cni_tag:v3.31.4,deploy_action:,external_network_id:a858742a-245b-41d3-9a05-617e1b069eb9,external_subnet_id_list:679217b7-1e03-4ce0-8007-7e5cf4f1c47c:05f97c52-a6c3-4422-8b24-31de607fed1c:de20ab10-e4dc-448d-a9a1-cc081fffed9e:5641871e-1d78-4978-a89a-63661e851cba:5cdf279e-1335-4b75-b802-0f5c75337b54:7c5aabaa-63ea-410d-8644-20185b7e31c8:f7038b14-a310-4eb6-92ed-a7d9551b82d7:22011c23-653a-425b-b093-7bf3fb25574d:fc06aecc-b0f9-48ad-9034-62a411c8cc9a:be54270f-4858-413c-b448-135dc37521da:a107a4f8-0e05-44ca-8574-c311f36a9c96:cb12a663-9938-4aa5-85aa-0945d74c4462:995add64-8bda-468e-af0f-533c00888f5a:85558a25-94b7-459d-b8ba-53740c9310db:cf714aed-b3c6-45eb-8867-08dce9f27f92:6d8e37b1-4c4b-4ed6-b34c-b2e12957d00a:d832f723-5fe6-4f55-9493-5fe9467cc259:7f6c2ade-1b11-433b-ad6c-3edefbfc038b:91ab5eda-69a2-4312-8968-ec48bf7eaef0:714c8fe2-d064-47f6-8451-9fbcc388817d:3e4a8dea-4287-4039-9f52-7317a1065723:0af1bb82-f31c-4d74-9108-f93001bc21f2:26e80586-ce00-4b45-9f23-a5f8a27485b1:4b4ed6ce-5f41-4d3c-8363-9e9928cdf51e:697edd38-a62e-4ed2-a225-e27d0a40ce89:e70a4a7f-d58d-497e-9161-ec75415c0724:170c7e54-3710-4889-a726-f8fa868b191d:e4375237-baba-4a41-a928-b4693ba0e7a3:d1e765e5-3f2e-49cf-bf9c-c6cbe5a2e496:d3e22864-b278-4923-9af5-d4f07d2ee0fb:7e486769-a377-46aa-8ec3-f0e4c2dd6c30:11133113-1a2c-4a67-b06f-0b4cf9809e92:8c4e3353-266b-4d0d-b35f-997659a154bc:62733ccb-1ab7-4913-b407-62e0381b5a6b:161335a4-c65b-47a6-bd87-4205b75ce7ca:d73e0bcc-4bce-4e3d-ade8-8b041cbf8232:0c844317-a014-4965-9bc8-c5588d668698:27de04d3-8984-4430-843d-bcfd6f821fdf:24d9fcc8-5725-4da5-9aa4-ad7707b1cfed:6fe5086b-ef56-4454-8aee-fe9d169bfdbf:0bce294a-b772-4c6c-97c9-25e3c2f2ce72:a6f27f47-166a-4a71-b09d-ae2d3ff0aedf:e8951dff-aa5a-4619-bebe-fc919a92866b:42ae5884-70bf-4232-8baa-f514ac280775:a2551567-b464-424c-ab13-8dcaf1871348:0641f8ac-c7e9-43a8-9eb5-ba63d08b83e0:9217260d-35df-4c69-af35-340a721533a8:eeb4ec59-4886-4283-94d9-7192ee41bd65:716da39f-19cb-4adb-a533-8a040232e4fd:fb3d25ea-3f60-4928-b56a-ca152f8eceee,extra_security_groups:[],extra_volumes:[],kube_tag:v1.33.4,kube_version_status:LATEST,master_lb_floating_ip_enabled:True,mba_scale_in:{\\enable\\: false, \\min_node_count\\: 1, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},mba_scale_out:{\\enable\\: false, \\max_node_count\\: 10, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},node_image:107cc02d-02d8-44fd-84d6-6c316045b817,node_list:8f4add36-6a8b-4db8-8854-c83515a035d3:07afb0fc-0226-449d-a7d4-bfb7ba6e0bca,node_resize_complete_time:2026-10-07T17:01:29,platform_version:1.202608.0,pods_network_cidr:10.100.0.0/16,pods_network_subnet:24,project_domain:NORMAL,requested_node_count:2,restore_status_on_failure:,service_cluster_ip_range:10.254.0.0/16,station_id:default,strict_sg_rules:False,upgradable_kube_versions:[],upgradable_platform_version:}"
          },
          {
            "key": "Links",
            "value": "{href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/v1/clusters/3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7/nodegroups/0af3aadb-40b4-47ea-93f0-bd63ab43abd4,rel:self}; {href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/clusters/3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7/nodegroups/0af3aadb-40b4-47ea-93f0-bd63ab43abd4,rel:bookmark}"
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
            "value": "10.0.1.104; 10.0.1.6"
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
            "value": "1a993d78-21ec-4fbd-9bae-81f045e5c101"
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
            "value": "2026-10-07T07:43:17Z"
          },
          {
            "key": "UpdatedAt",
            "value": "2026-10-07T08:01:50Z"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://3e1392d2-nks-kr1.container.nhncloud.com:6443",
    "kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURMakNDQWhhZ0F3SUJBZ0lRQmd5Z3JoSTdUREcyMWZveHBDdU1IVEFOQmdrcWhraUc5dzBCQVFzRkFEQWYKTVIwd0d3WURWUVFEREJSMFluRmhjbUpqTXpoMU5HcGthR3BtY0Rob2J6QWVGdzB5TmpFd01EWXdOelF6TVRoYQpGdzB6TVRFd01EWXdOelF6TVRoYU1COHhIVEFiQmdOVkJBTU1GSFJpY1dGeVltTXpPSFUwYW1Sb2FtWndPR2h2Ck1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBbmtmb3drNjF6MlJZaEdiV2x5NDkKSjFES3daZFhKRzBZUVp3NnBmYnZuYWpSZUl0SG5Ba3JJOExRUHBaOFVSSUg0c3FsQmZHUWJFMERkRHB6cHZiOQpqd3E4MFdoajJHQWJLcERJWXNvZFJrZjUyZWtrdTRWOXhnY0FLMVc0c2NxNVhrTzNXZVQrYURCSGs4dTBxY2tzCmVhdEhJV0xCNno1S0RmOEZwYzNCTWdjLzhlUC9VMW1rTk5va3A0dGpZU0E2ODBLaHpEWVBwOWdRdXFXWFBHajcKelAvVEdwVnBFN0swazFNQWw2WU1CWmJaaGREalBSWWd6WEZpM1Q3ZDhhNjRhVnplQ1pLeUUxOExpK2p5S0w5YgppTitIT0orWFczWkd5TnNWM3A4a1pJSHhwRG9SdEhUWi96UStGNmMwYk5IOGtrdWhtdE9SV05vZVV6MVYza1Q4CnJ3SURBUUFCbzJZd1pEQVNCZ05WSFJNQkFmOEVDREFHQVFIL0FnRUFNQTRHQTFVZER3RUIvd1FFQXdJQ0JEQWQKQmdOVkhRNEVGZ1FVd252ZUtTMEFERzVyWHF1dmhtc3ZaeHExRms0d0h3WURWUjBqQkJnd0ZvQVV3bnZlS1MwQQpERzVyWHF1dmhtc3ZaeHExRms0d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFBNGxlUys5cmFndzdYNXdXVkVLCmtkNWxvSzlCQTUwaGVhK2ZSVzZESmZZMmN3bEJtem9QWkVsc0pJbXFNSE56SlBqM1QvZ0lxQWViNWFPSjdIZTYKa2I1NENCOG13Q0dvMXRYdndDTkpYL1hldXV3d0xzSjN4SmdaM2RuVVVEeVJzemdNTEEzNzBURFNvUDVzTTd4NApCWWV3YzI3dEE1U0d2bDhaRWM3cUJRTnk5aHg2RUVTSE5sM0Y2R1d6MFNxRFI2MjQzZCtPRkRIdlBUWFViaEd0Cld1QTR1K2VQZkpDTHZ5MDI1SmpZZDY5ZFg0a1FWOTY5L3pIQjFMNWkyQ2RtbStYdzJFbEk2Tmd1M1VVY05vQjIKODBIVlZVY2FBSW85N1NiR2hiZ25SVmJVeTJ3eERaeFE4bHNSa1ZMN0xDajVxWHhEdUl5YTJXNk1SWDhGOUNmRQpPbUE9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=\n    server: https://3e1392d2-nks-kr1.container.nhncloud.com:6443\n  name: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\ncontexts:\n- context:\n    cluster: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\n    user: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\n  name: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\ncurrent-context: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\nkind: Config\npreferences: {}\nusers:\n- name: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGVENDQWYyZ0F3SUJBZ0lSQVA4MkE0WVlaRUt5bVZPekt2cS9uYU13RFFZSktvWklodmNOQVFFTEJRQXcKSHpFZE1Cc0dBMVVFQXd3VWRHSnhZWEppWXpNNGRUUnFaR2hxWm5BNGFHOHdIaGNOTWpZeE1EQTJNRGd3TWpJMApXaGNOTXpFeE1EQTJNRGd3TWpJMFdqQXBNUTR3REFZRFZRUUREQVZoWkcxcGJqRVhNQlVHQTFVRUNnd09jM2x6CmRHVnRPbTFoYzNSbGNuTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFDbFhYUjUKTE90WExOL0JNS2R0ajFZaEdyeTJ2NTgwMFVjcXZGWVYybFlBcGhXUVNRb2kyTS9XOTdlSXQ4VDhncTNKNW1LcgpNNVkxR0lEQ0xJSnl4R3owRWR2dFVISzlzZE0yMnhzY292b2tKM2ZIYWs0bGlWTGxMcTFIdEJaQ3pIOFpKa2N6Cmh2OU9vakNoaWxPbXpSeXRySjVKRHkxRnNxbTB6aVVzbFJJWm1QdGtCZEU0eVZUQzV5d3pMaTJiKzRzRElINHAKUlNBQysrU2NNQWJTaGVEd2ROOW9VVHYrc1FCRjgzZlgrUHpibklNU0R0VGh1TlA2bWpTdDF0WUhsUWN3bU1ZNgpUV0wwL1FjNzYyWW5HWWlKNGhLb2syaUprWTNyVWxNVWl3dEp5dlZWU0xVSFNxU09mQ1pUUVFPYVZLemJpeEVrCk82aEl3cTVIMU9Fc0JjQS9BZ01CQUFHalFqQkFNQjBHQTFVZERnUVdCQlNnbzJBeGF6VmlETEVnSE5tQjZReDAKWDkwWkNEQWZCZ05WSFNNRUdEQVdnQlRDZTk0cExRQU1ibXRlcTYrR2F5OW5HclVXVGpBTkJna3Foa2lHOXcwQgpBUXNGQUFPQ0FRRUFGL3ZZNGN5cEFGWng1Rkl1N3hCSytkYzQ2R0ZTOXljT1VTZFhyWm1hZlZLMzJ3cSttWUZDClJ0MDIyTGhTZ0NSSHFyNGNlTEFNTWZqQXB6dHpvWEJCdGE4a1EvaDFaVDhwU1kvaHFZQ3k5QjBqV3NkckQ4blIKSHpNL1grOFE5aVpPZXFhMTAvdnYxd3czMjNyeVFlM1RuZnRqaHhlWlVyQ0hrbUllOVZlMWhFdHE5L1M3ODg5TQpqUUF3Z0ZSVGNFR2FsTGJMQmxFcGhlUGhzbUFFVTJPd3ViYVpGTHA1MlJGdXNlU1BOcktGcFFHOXluUDlpY0JWClRyVm9XcUlQRW9BLzBMQ0VyQkI1N01US0poNTNFK2VQVk5SRTZoUFNNdnBSeUhaclBQSmVpWkt3WHFkZHhqVVUKbHV4TStEQUpHU2EwZU5iOXFnWnNmelRpb1dIZytyWVRtQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBcFYxMGVTenJWeXpmd1RDbmJZOVdJUnE4dHIrZk5ORkhLcnhXRmRwV0FLWVZrRWtLCkl0alAxdmUzaUxmRS9JS3R5ZVppcXpPV05SaUF3aXlDY3NSczlCSGI3VkJ5dmJIVE50c2JIS0w2SkNkM3gycE8KSllsUzVTNnRSN1FXUXN4L0dTWkhNNGIvVHFJd29ZcFRwczBjcmF5ZVNROHRSYktwdE00bExKVVNHWmo3WkFYUgpPTWxVd3Vjc015NHRtL3VMQXlCK0tVVWdBdnZrbkRBRzBvWGc4SFRmYUZFNy9yRUFSZk4zMS9qODI1eURFZzdVCjRialQrcG8wcmRiV0I1VUhNSmpHT2sxaTlQMEhPK3RtSnhtSWllSVNxSk5vaVpHTjYxSlRGSXNMU2NyMVZVaTEKQjBxa2pud21VMEVEbWxTczI0c1JKRHVvU01LdVI5VGhMQVhBUHdJREFRQUJBb0lCQVFDUEhzaGN6QXpQZ1U0VApOVm5qRCtoNmRGcWFURlN0Y09lSVRFS3hYU3VNR2pLVmt6R0xvVnlZOEFMUCtFVXNTcW5CRVgyYU1PYXpBNGNGCkIyTExrOGU4b2V4WGc0WWJPZG5WSTJOZXNJdVFXUHhwQWsrcVd5VDhxQjVlN2JWS2tSVDduazZIZ1RKY1N6R04Kemx2Q0JNbnFZbmtJUUhOdFhFbFc3QWZ1Z3dLTVZhQWxDWmZ6Y2VhbjdFajFXcEozYW1VYnRvaFBLZmVvNHUyOAppVUczZmdoQVhVa2xCNjdzNWRLWkwrbG1hdHBkRGFaeTdpaGRDT3BMWXBneW9RWVdKN3ZrUm5WUTN5QmdQVW5JCmpkMzZUWjhtV2VyaVVSQXlGZU1zL051Rk1WQVBLN2h4U1hBT2lPTVlGS3J3Qm5qQkc2S0x6dTRlak1OelJ0OFEKRFNXMGJrMTVBb0dCQU5EVklNQ1BXT0FtVmtpZjFaOWJjMFZUM3p1VGliTWZpRmtFaEdTVkhvdXMzdTVMOVFPYgpUcmlVT2ZxY2FGY2NVNkR5OGlublUzVDNIbWo0Wml5TDZWQlltWnNwNmsxTW5uZzhVMXhpM3ZzYzl1Ym1IeEJNCkt5VVlPOFhtUHNUbUJ2OWREc1Q1c1B3RzRFNTgwUlFGU3Q3SkZ0N05nQnBLRXpzMmluTC9pamR0QW9HQkFNcTMKQUZiUkZiOFJrS1ZTQ2xGdVBteERrc0NqaGdtLzQyUTlIWC9JSThQcmR4bWJmTmpsdFFlUFZ2emo4NG92VGlPWAorSysyTXM2VHRqTXhrc3d0R1dSblJUK1ZwNE94L2tmRkdTemtsaVh4dHJkUW5Ga3pWSmdleFpTS3RxT2tTTXh3Ckl3RUUwVXV5QllFZGw2ZHhkQnF2QmZmWk9PWUpOODNxbjhxNHBlN2JBb0dBUHhVeWlKaUV1Mms1UlVUckRmRmQKRjhNbk91THRoS3R3ejVzYXR4dUlsNTdIMU9vakFFem84YjdzNkxSWW5IL1ZEMWgwS0k1VldvN1BvOXZ0V1dXRwpQQndYZ3BTZHoyM21yT0ZrT1JNdzBtQkVnU2xnbHRhN2JjQXRSakd3SW1CdUdyT3NvM25kcWJRV0MvK09WT0xvClkvWVFyNHNhT0ZiS0NtZ0dXSWl5NHQwQ2dZRUFpSW1HVXZOMWdFSm1lTk92VzNXejJkd1J4bXJkNTBjMTg5Z3UKUEZranJkeTFWYXdqQlR5REdGcDFydFFpdjFwcGVSUHRUdnliY0FTUjNoMXYwTkkzbmlib3Y1RWZTVmJqL0pkSApBN1BiWmlkT2VGNTFVS2VBUFEzTTZ1WWJhbEZITDF3QVY3bFU5M1VxdS9LZ3FRbUR1RTFXNmIrSDBYazMreHdqCkVMb2FHYTBDZ1lBTWwzVWVNM2dNTHFsTXNWRmVMeEhoeEhOK2dCdXZtTWZ0OHZSekxCOHh0dlZKVDNoeW82dXkKNmlTTkllUDI2SVlqNis5cjN4cWRKcUlIODc2Q3ozaWhNSitSWnI5d2VncEFiV080UE9qRkUrOE1qeVAzQ09tQQptek9udEZEd3VhVzRwTXhGZkg1Q3gxcnAvS0QwaG9rWXdTb21UVy8wZDJMWnVDV0ZRL21pMWc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-10-07T07:43:17Z",
  "keyValueList": [
    {
      "key": "APIAddress",
      "value": "https://3e1392d2-nks-kr1.container.nhncloud.com:6443"
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
      "value": "2026-10-07T07:43:17Z"
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
      "value": "tbtj1kicavi8btlht22d"
    },
    {
      "key": "Labels",
      "value": "{additional_network_id_list:,additional_subnet_id_list:,api_ep_ipacl_enable:false,availability_zone:kr-pub-a,cert_manager_api:True,certificate_expiry:2031-10-06T07:43:17+00:00,cni_driver:calico,cni_tag:v3.31.4,default_security_group:eb2724f0-84af-413f-8a3a-2ff7cfa5e9c3,external_network_id:a858742a-245b-41d3-9a05-617e1b069eb9,external_subnet_id_list:679217b7-1e03-4ce0-8007-7e5cf4f1c47c:05f97c52-a6c3-4422-8b24-31de607fed1c:de20ab10-e4dc-448d-a9a1-cc081fffed9e:5641871e-1d78-4978-a89a-63661e851cba:5cdf279e-1335-4b75-b802-0f5c75337b54:7c5aabaa-63ea-410d-8644-20185b7e31c8:f7038b14-a310-4eb6-92ed-a7d9551b82d7:22011c23-653a-425b-b093-7bf3fb25574d:fc06aecc-b0f9-48ad-9034-62a411c8cc9a:be54270f-4858-413c-b448-135dc37521da:a107a4f8-0e05-44ca-8574-c311f36a9c96:cb12a663-9938-4aa5-85aa-0945d74c4462:995add64-8bda-468e-af0f-533c00888f5a:85558a25-94b7-459d-b8ba-53740c9310db:cf714aed-b3c6-45eb-8867-08dce9f27f92:6d8e37b1-4c4b-4ed6-b34c-b2e12957d00a:d832f723-5fe6-4f55-9493-5fe9467cc259:7f6c2ade-1b11-433b-ad6c-3edefbfc038b:91ab5eda-69a2-4312-8968-ec48bf7eaef0:714c8fe2-d064-47f6-8451-9fbcc388817d:3e4a8dea-4287-4039-9f52-7317a1065723:0af1bb82-f31c-4d74-9108-f93001bc21f2:26e80586-ce00-4b45-9f23-a5f8a27485b1:4b4ed6ce-5f41-4d3c-8363-9e9928cdf51e:697edd38-a62e-4ed2-a225-e27d0a40ce89:e70a4a7f-d58d-497e-9161-ec75415c0724:170c7e54-3710-4889-a726-f8fa868b191d:e4375237-baba-4a41-a928-b4693ba0e7a3:d1e765e5-3f2e-49cf-bf9c-c6cbe5a2e496:d3e22864-b278-4923-9af5-d4f07d2ee0fb:7e486769-a377-46aa-8ec3-f0e4c2dd6c30:11133113-1a2c-4a67-b06f-0b4cf9809e92:8c4e3353-266b-4d0d-b35f-997659a154bc:62733ccb-1ab7-4913-b407-62e0381b5a6b:161335a4-c65b-47a6-bd87-4205b75ce7ca:d73e0bcc-4bce-4e3d-ade8-8b041cbf8232:0c844317-a014-4965-9bc8-c5588d668698:27de04d3-8984-4430-843d-bcfd6f821fdf:24d9fcc8-5725-4da5-9aa4-ad7707b1cfed:6fe5086b-ef56-4454-8aee-fe9d169bfdbf:0bce294a-b772-4c6c-97c9-25e3c2f2ce72:a6f27f47-166a-4a71-b09d-ae2d3ff0aedf:e8951dff-aa5a-4619-bebe-fc919a92866b:42ae5884-70bf-4232-8baa-f514ac280775:a2551567-b464-424c-ab13-8dcaf1871348:0641f8ac-c7e9-43a8-9eb5-ba63d08b83e0:9217260d-35df-4c69-af35-340a721533a8:eeb4ec59-4886-4283-94d9-7192ee41bd65:716da39f-19cb-4adb-a533-8a040232e4fd:fb3d25ea-3f60-4928-b56a-ca152f8eceee,k8s_args:{\\kube-apiserver/default-not-ready-toleration-seconds\\: 300, \\kube-apiserver/default-unreachable-toleration-seconds\\: 300, \\kube-controller-manager/node-monitor-grace-period\\: 40, \\kube-controller-manager/unhealthy-zone-threshold\\: 55},kube_tag:v1.33.4,kube_version_status:NEED_K8S_UPGRADE,master_lb_floating_ip_enabled:True,master_metric_agent:telegraf,nks_registry_url:dfe965c3-kr1-registry.container.nhncloud.com/container_service,node_image:107cc02d-02d8-44fd-84d6-6c316045b817,platform_version:1.202608.0,pods_network_cidr:10.100.0.0/16,pods_network_subnet:24,project_domain:NORMAL,service_cluster_ip_range:10.254.0.0/16,service_user_enabled:True,station_id:default,strict_sg_rules:False,term_of_validity:5,upgradable_kube_versions:[\\v1.34.3\\],upgradable_platform_version:}"
    },
    {
      "key": "Links",
      "value": "{href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/v1/clusters/3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7,rel:self}; {href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/clusters/3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7,rel:bookmark}"
    },
    {
      "key": "MasterCount",
      "value": "3"
    },
    {
      "key": "Name",
      "value": "tbqarbc38u4jdhjfp8ho"
    },
    {
      "key": "NodeAddresses",
      "value": "10.0.1.104; 10.0.1.6"
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
      "value": "3cb524cc-8d6c-4e4e-99f4-893569fbac7e"
    },
    {
      "key": "Status",
      "value": "CREATE_COMPLETE"
    },
    {
      "key": "UUID",
      "value": "3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7"
    },
    {
      "key": "UpdatedAt",
      "value": "2026-10-07T08:02:03Z"
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
      "value": "d3c022a0-6f2c-42c3-9b2f-c15766f4ca04"
    },
    {
      "key": "FixedSubnet",
      "value": "78c23b29-ad0d-4621-8095-87ec9f253010"
    },
    {
      "key": "HealthStatus",
      "value": "FRESH"
    },
    {
      "key": "HealthStatusReason",
      "value": "{api:OK,cluster.api_status:NORMAL,cluster.node_status:NORMAL,nodegroup-stats.default-worker:2:0,nodegroup.node_status.default-worker:NORMAL,timestamp:2026-10-07T17:02:02.663710}"
    }
  ],
  "cspResourceName": "tbqarbc38u4jdhjfp8ho",
  "cspResourceId": "3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbqarbc38u4jdhjfp8ho",
      "SystemId": "3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7"
    },
    "Version": "v1.33.4",
    "Network": {
      "VpcIID": {
        "NameId": "tb49fg3iamnab5vmms1h",
        "SystemId": "d3c022a0-6f2c-42c3-9b2f-c15766f4ca04"
      },
      "SubnetIIDs": [
        {
          "NameId": "tbq2g1outbpm8831q09o",
          "SystemId": "78c23b29-ad0d-4621-8095-87ec9f253010"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "#eb2724f0-84af-413f-8a3a-2ff7cfa5e9c3",
          "SystemId": "eb2724f0-84af-413f-8a3a-2ff7cfa5e9c3"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "0af3aadb-40b4-47ea-93f0-bd63ab43abd4"
        },
        "ImageIID": {
          "NameId": "107cc02d-02d8-44fd-84d6-6c316045b817",
          "SystemId": "107cc02d-02d8-44fd-84d6-6c316045b817"
        },
        "VMSpecName": "m2.c4m8",
        "RootDiskType": "General HDD",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tbtj1kicavi8btlht22d",
          "SystemId": "tbtj1kicavi8btlht22d"
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "tbqarbc38u4jdhjfp8ho-default-worker-node-0",
            "SystemId": "8f4add36-6a8b-4db8-8854-c83515a035d3"
          },
          {
            "NameId": "tbqarbc38u4jdhjfp8ho-default-worker-node-1",
            "SystemId": "07afb0fc-0226-449d-a7d4-bfb7ba6e0bca"
          }
        ],
        "KeyValueList": [
          {
            "key": "ID",
            "value": "109030"
          },
          {
            "key": "UUID",
            "value": "0af3aadb-40b4-47ea-93f0-bd63ab43abd4"
          },
          {
            "key": "Name",
            "value": "default-worker"
          },
          {
            "key": "ClusterID",
            "value": "3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7"
          },
          {
            "key": "ProjectID",
            "value": "01d5dbda7c204a7c9872a61b5a46dd40"
          },
          {
            "key": "Labels",
            "value": "{additional_network_id_list:,additional_subnet_id_list:,allow_signal_api:false,availability_zone:kr-pub-a,boot_volume_size:100,boot_volume_type:General HDD,ca_enable:false,ca_max_node_count:0,ca_min_node_count:0,ca_pod_replicas:1,ca_scale_down_delay_after_add:10,ca_scale_down_enable:true,ca_scale_down_unneeded_time:10,ca_scale_down_util_thresh:50,cert_manager_api:True,cgroup:v2,clusterautoscale:nodegroupfeature,cni_driver:calico,cni_tag:v3.31.4,deploy_action:,external_network_id:a858742a-245b-41d3-9a05-617e1b069eb9,external_subnet_id_list:679217b7-1e03-4ce0-8007-7e5cf4f1c47c:05f97c52-a6c3-4422-8b24-31de607fed1c:de20ab10-e4dc-448d-a9a1-cc081fffed9e:5641871e-1d78-4978-a89a-63661e851cba:5cdf279e-1335-4b75-b802-0f5c75337b54:7c5aabaa-63ea-410d-8644-20185b7e31c8:f7038b14-a310-4eb6-92ed-a7d9551b82d7:22011c23-653a-425b-b093-7bf3fb25574d:fc06aecc-b0f9-48ad-9034-62a411c8cc9a:be54270f-4858-413c-b448-135dc37521da:a107a4f8-0e05-44ca-8574-c311f36a9c96:cb12a663-9938-4aa5-85aa-0945d74c4462:995add64-8bda-468e-af0f-533c00888f5a:85558a25-94b7-459d-b8ba-53740c9310db:cf714aed-b3c6-45eb-8867-08dce9f27f92:6d8e37b1-4c4b-4ed6-b34c-b2e12957d00a:d832f723-5fe6-4f55-9493-5fe9467cc259:7f6c2ade-1b11-433b-ad6c-3edefbfc038b:91ab5eda-69a2-4312-8968-ec48bf7eaef0:714c8fe2-d064-47f6-8451-9fbcc388817d:3e4a8dea-4287-4039-9f52-7317a1065723:0af1bb82-f31c-4d74-9108-f93001bc21f2:26e80586-ce00-4b45-9f23-a5f8a27485b1:4b4ed6ce-5f41-4d3c-8363-9e9928cdf51e:697edd38-a62e-4ed2-a225-e27d0a40ce89:e70a4a7f-d58d-497e-9161-ec75415c0724:170c7e54-3710-4889-a726-f8fa868b191d:e4375237-baba-4a41-a928-b4693ba0e7a3:d1e765e5-3f2e-49cf-bf9c-c6cbe5a2e496:d3e22864-b278-4923-9af5-d4f07d2ee0fb:7e486769-a377-46aa-8ec3-f0e4c2dd6c30:11133113-1a2c-4a67-b06f-0b4cf9809e92:8c4e3353-266b-4d0d-b35f-997659a154bc:62733ccb-1ab7-4913-b407-62e0381b5a6b:161335a4-c65b-47a6-bd87-4205b75ce7ca:d73e0bcc-4bce-4e3d-ade8-8b041cbf8232:0c844317-a014-4965-9bc8-c5588d668698:27de04d3-8984-4430-843d-bcfd6f821fdf:24d9fcc8-5725-4da5-9aa4-ad7707b1cfed:6fe5086b-ef56-4454-8aee-fe9d169bfdbf:0bce294a-b772-4c6c-97c9-25e3c2f2ce72:a6f27f47-166a-4a71-b09d-ae2d3ff0aedf:e8951dff-aa5a-4619-bebe-fc919a92866b:42ae5884-70bf-4232-8baa-f514ac280775:a2551567-b464-424c-ab13-8dcaf1871348:0641f8ac-c7e9-43a8-9eb5-ba63d08b83e0:9217260d-35df-4c69-af35-340a721533a8:eeb4ec59-4886-4283-94d9-7192ee41bd65:716da39f-19cb-4adb-a533-8a040232e4fd:fb3d25ea-3f60-4928-b56a-ca152f8eceee,extra_security_groups:[],extra_volumes:[],kube_tag:v1.33.4,kube_version_status:LATEST,master_lb_floating_ip_enabled:True,mba_scale_in:{\\enable\\: false, \\min_node_count\\: 1, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},mba_scale_out:{\\enable\\: false, \\max_node_count\\: 10, \\rules_operator\\: \\OR\\, \\delay\\: 10, \\adjustment_count\\: 1, \\rules\\: []},node_image:107cc02d-02d8-44fd-84d6-6c316045b817,node_list:8f4add36-6a8b-4db8-8854-c83515a035d3:07afb0fc-0226-449d-a7d4-bfb7ba6e0bca,node_resize_complete_time:2026-10-07T17:01:29,platform_version:1.202608.0,pods_network_cidr:10.100.0.0/16,pods_network_subnet:24,project_domain:NORMAL,requested_node_count:2,restore_status_on_failure:,service_cluster_ip_range:10.254.0.0/16,station_id:default,strict_sg_rules:False,upgradable_kube_versions:[],upgradable_platform_version:}"
          },
          {
            "key": "Links",
            "value": "{href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/v1/clusters/3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7/nodegroups/0af3aadb-40b4-47ea-93f0-bd63ab43abd4,rel:self}; {href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/clusters/3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7/nodegroups/0af3aadb-40b4-47ea-93f0-bd63ab43abd4,rel:bookmark}"
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
            "value": "10.0.1.104; 10.0.1.6"
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
            "value": "1a993d78-21ec-4fbd-9bae-81f045e5c101"
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
            "value": "2026-10-07T07:43:17Z"
          },
          {
            "key": "UpdatedAt",
            "value": "2026-10-07T08:01:50Z"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://3e1392d2-nks-kr1.container.nhncloud.com:6443",
      "Kubeconfig": "apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURMakNDQWhhZ0F3SUJBZ0lRQmd5Z3JoSTdUREcyMWZveHBDdU1IVEFOQmdrcWhraUc5dzBCQVFzRkFEQWYKTVIwd0d3WURWUVFEREJSMFluRmhjbUpqTXpoMU5HcGthR3BtY0Rob2J6QWVGdzB5TmpFd01EWXdOelF6TVRoYQpGdzB6TVRFd01EWXdOelF6TVRoYU1COHhIVEFiQmdOVkJBTU1GSFJpY1dGeVltTXpPSFUwYW1Sb2FtWndPR2h2Ck1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBbmtmb3drNjF6MlJZaEdiV2x5NDkKSjFES3daZFhKRzBZUVp3NnBmYnZuYWpSZUl0SG5Ba3JJOExRUHBaOFVSSUg0c3FsQmZHUWJFMERkRHB6cHZiOQpqd3E4MFdoajJHQWJLcERJWXNvZFJrZjUyZWtrdTRWOXhnY0FLMVc0c2NxNVhrTzNXZVQrYURCSGs4dTBxY2tzCmVhdEhJV0xCNno1S0RmOEZwYzNCTWdjLzhlUC9VMW1rTk5va3A0dGpZU0E2ODBLaHpEWVBwOWdRdXFXWFBHajcKelAvVEdwVnBFN0swazFNQWw2WU1CWmJaaGREalBSWWd6WEZpM1Q3ZDhhNjRhVnplQ1pLeUUxOExpK2p5S0w5YgppTitIT0orWFczWkd5TnNWM3A4a1pJSHhwRG9SdEhUWi96UStGNmMwYk5IOGtrdWhtdE9SV05vZVV6MVYza1Q4CnJ3SURBUUFCbzJZd1pEQVNCZ05WSFJNQkFmOEVDREFHQVFIL0FnRUFNQTRHQTFVZER3RUIvd1FFQXdJQ0JEQWQKQmdOVkhRNEVGZ1FVd252ZUtTMEFERzVyWHF1dmhtc3ZaeHExRms0d0h3WURWUjBqQkJnd0ZvQVV3bnZlS1MwQQpERzVyWHF1dmhtc3ZaeHExRms0d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFBNGxlUys5cmFndzdYNXdXVkVLCmtkNWxvSzlCQTUwaGVhK2ZSVzZESmZZMmN3bEJtem9QWkVsc0pJbXFNSE56SlBqM1QvZ0lxQWViNWFPSjdIZTYKa2I1NENCOG13Q0dvMXRYdndDTkpYL1hldXV3d0xzSjN4SmdaM2RuVVVEeVJzemdNTEEzNzBURFNvUDVzTTd4NApCWWV3YzI3dEE1U0d2bDhaRWM3cUJRTnk5aHg2RUVTSE5sM0Y2R1d6MFNxRFI2MjQzZCtPRkRIdlBUWFViaEd0Cld1QTR1K2VQZkpDTHZ5MDI1SmpZZDY5ZFg0a1FWOTY5L3pIQjFMNWkyQ2RtbStYdzJFbEk2Tmd1M1VVY05vQjIKODBIVlZVY2FBSW85N1NiR2hiZ25SVmJVeTJ3eERaeFE4bHNSa1ZMN0xDajVxWHhEdUl5YTJXNk1SWDhGOUNmRQpPbUE9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=\n    server: https://3e1392d2-nks-kr1.container.nhncloud.com:6443\n  name: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\ncontexts:\n- context:\n    cluster: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\n    user: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\n  name: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\ncurrent-context: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\nkind: Config\npreferences: {}\nusers:\n- name: \"nks_tbqarbc38u4jdhjfp8ho_3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7\"\n  user:\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGVENDQWYyZ0F3SUJBZ0lSQVA4MkE0WVlaRUt5bVZPekt2cS9uYU13RFFZSktvWklodmNOQVFFTEJRQXcKSHpFZE1Cc0dBMVVFQXd3VWRHSnhZWEppWXpNNGRUUnFaR2hxWm5BNGFHOHdIaGNOTWpZeE1EQTJNRGd3TWpJMApXaGNOTXpFeE1EQTJNRGd3TWpJMFdqQXBNUTR3REFZRFZRUUREQVZoWkcxcGJqRVhNQlVHQTFVRUNnd09jM2x6CmRHVnRPbTFoYzNSbGNuTXdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFDbFhYUjUKTE90WExOL0JNS2R0ajFZaEdyeTJ2NTgwMFVjcXZGWVYybFlBcGhXUVNRb2kyTS9XOTdlSXQ4VDhncTNKNW1LcgpNNVkxR0lEQ0xJSnl4R3owRWR2dFVISzlzZE0yMnhzY292b2tKM2ZIYWs0bGlWTGxMcTFIdEJaQ3pIOFpKa2N6Cmh2OU9vakNoaWxPbXpSeXRySjVKRHkxRnNxbTB6aVVzbFJJWm1QdGtCZEU0eVZUQzV5d3pMaTJiKzRzRElINHAKUlNBQysrU2NNQWJTaGVEd2ROOW9VVHYrc1FCRjgzZlgrUHpibklNU0R0VGh1TlA2bWpTdDF0WUhsUWN3bU1ZNgpUV0wwL1FjNzYyWW5HWWlKNGhLb2syaUprWTNyVWxNVWl3dEp5dlZWU0xVSFNxU09mQ1pUUVFPYVZLemJpeEVrCk82aEl3cTVIMU9Fc0JjQS9BZ01CQUFHalFqQkFNQjBHQTFVZERnUVdCQlNnbzJBeGF6VmlETEVnSE5tQjZReDAKWDkwWkNEQWZCZ05WSFNNRUdEQVdnQlRDZTk0cExRQU1ibXRlcTYrR2F5OW5HclVXVGpBTkJna3Foa2lHOXcwQgpBUXNGQUFPQ0FRRUFGL3ZZNGN5cEFGWng1Rkl1N3hCSytkYzQ2R0ZTOXljT1VTZFhyWm1hZlZLMzJ3cSttWUZDClJ0MDIyTGhTZ0NSSHFyNGNlTEFNTWZqQXB6dHpvWEJCdGE4a1EvaDFaVDhwU1kvaHFZQ3k5QjBqV3NkckQ4blIKSHpNL1grOFE5aVpPZXFhMTAvdnYxd3czMjNyeVFlM1RuZnRqaHhlWlVyQ0hrbUllOVZlMWhFdHE5L1M3ODg5TQpqUUF3Z0ZSVGNFR2FsTGJMQmxFcGhlUGhzbUFFVTJPd3ViYVpGTHA1MlJGdXNlU1BOcktGcFFHOXluUDlpY0JWClRyVm9XcUlQRW9BLzBMQ0VyQkI1N01US0poNTNFK2VQVk5SRTZoUFNNdnBSeUhaclBQSmVpWkt3WHFkZHhqVVUKbHV4TStEQUpHU2EwZU5iOXFnWnNmelRpb1dIZytyWVRtQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0=\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBcFYxMGVTenJWeXpmd1RDbmJZOVdJUnE4dHIrZk5ORkhLcnhXRmRwV0FLWVZrRWtLCkl0alAxdmUzaUxmRS9JS3R5ZVppcXpPV05SaUF3aXlDY3NSczlCSGI3VkJ5dmJIVE50c2JIS0w2SkNkM3gycE8KSllsUzVTNnRSN1FXUXN4L0dTWkhNNGIvVHFJd29ZcFRwczBjcmF5ZVNROHRSYktwdE00bExKVVNHWmo3WkFYUgpPTWxVd3Vjc015NHRtL3VMQXlCK0tVVWdBdnZrbkRBRzBvWGc4SFRmYUZFNy9yRUFSZk4zMS9qODI1eURFZzdVCjRialQrcG8wcmRiV0I1VUhNSmpHT2sxaTlQMEhPK3RtSnhtSWllSVNxSk5vaVpHTjYxSlRGSXNMU2NyMVZVaTEKQjBxa2pud21VMEVEbWxTczI0c1JKRHVvU01LdVI5VGhMQVhBUHdJREFRQUJBb0lCQVFDUEhzaGN6QXpQZ1U0VApOVm5qRCtoNmRGcWFURlN0Y09lSVRFS3hYU3VNR2pLVmt6R0xvVnlZOEFMUCtFVXNTcW5CRVgyYU1PYXpBNGNGCkIyTExrOGU4b2V4WGc0WWJPZG5WSTJOZXNJdVFXUHhwQWsrcVd5VDhxQjVlN2JWS2tSVDduazZIZ1RKY1N6R04Kemx2Q0JNbnFZbmtJUUhOdFhFbFc3QWZ1Z3dLTVZhQWxDWmZ6Y2VhbjdFajFXcEozYW1VYnRvaFBLZmVvNHUyOAppVUczZmdoQVhVa2xCNjdzNWRLWkwrbG1hdHBkRGFaeTdpaGRDT3BMWXBneW9RWVdKN3ZrUm5WUTN5QmdQVW5JCmpkMzZUWjhtV2VyaVVSQXlGZU1zL051Rk1WQVBLN2h4U1hBT2lPTVlGS3J3Qm5qQkc2S0x6dTRlak1OelJ0OFEKRFNXMGJrMTVBb0dCQU5EVklNQ1BXT0FtVmtpZjFaOWJjMFZUM3p1VGliTWZpRmtFaEdTVkhvdXMzdTVMOVFPYgpUcmlVT2ZxY2FGY2NVNkR5OGlublUzVDNIbWo0Wml5TDZWQlltWnNwNmsxTW5uZzhVMXhpM3ZzYzl1Ym1IeEJNCkt5VVlPOFhtUHNUbUJ2OWREc1Q1c1B3RzRFNTgwUlFGU3Q3SkZ0N05nQnBLRXpzMmluTC9pamR0QW9HQkFNcTMKQUZiUkZiOFJrS1ZTQ2xGdVBteERrc0NqaGdtLzQyUTlIWC9JSThQcmR4bWJmTmpsdFFlUFZ2emo4NG92VGlPWAorSysyTXM2VHRqTXhrc3d0R1dSblJUK1ZwNE94L2tmRkdTemtsaVh4dHJkUW5Ga3pWSmdleFpTS3RxT2tTTXh3Ckl3RUUwVXV5QllFZGw2ZHhkQnF2QmZmWk9PWUpOODNxbjhxNHBlN2JBb0dBUHhVeWlKaUV1Mms1UlVUckRmRmQKRjhNbk91THRoS3R3ejVzYXR4dUlsNTdIMU9vakFFem84YjdzNkxSWW5IL1ZEMWgwS0k1VldvN1BvOXZ0V1dXRwpQQndYZ3BTZHoyM21yT0ZrT1JNdzBtQkVnU2xnbHRhN2JjQXRSakd3SW1CdUdyT3NvM25kcWJRV0MvK09WT0xvClkvWVFyNHNhT0ZiS0NtZ0dXSWl5NHQwQ2dZRUFpSW1HVXZOMWdFSm1lTk92VzNXejJkd1J4bXJkNTBjMTg5Z3UKUEZranJkeTFWYXdqQlR5REdGcDFydFFpdjFwcGVSUHRUdnliY0FTUjNoMXYwTkkzbmlib3Y1RWZTVmJqL0pkSApBN1BiWmlkT2VGNTFVS2VBUFEzTTZ1WWJhbEZITDF3QVY3bFU5M1VxdS9LZ3FRbUR1RTFXNmIrSDBYazMreHdqCkVMb2FHYTBDZ1lBTWwzVWVNM2dNTHFsTXNWRmVMeEhoeEhOK2dCdXZtTWZ0OHZSekxCOHh0dlZKVDNoeW82dXkKNmlTTkllUDI2SVlqNis5cjN4cWRKcUlIODc2Q3ozaWhNSitSWnI5d2VncEFiV080UE9qRkUrOE1qeVAzQ09tQQptek9udEZEd3VhVzRwTXhGZkg1Q3gxcnAvS0QwaG9rWXdTb21UVy8wZDJMWnVDV0ZRL21pMWc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-10-07T07:43:17Z",
    "KeyValueList": [
      {
        "key": "APIAddress",
        "value": "https://3e1392d2-nks-kr1.container.nhncloud.com:6443"
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
        "value": "2026-10-07T07:43:17Z"
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
        "value": "tbtj1kicavi8btlht22d"
      },
      {
        "key": "Labels",
        "value": "{additional_network_id_list:,additional_subnet_id_list:,api_ep_ipacl_enable:false,availability_zone:kr-pub-a,cert_manager_api:True,certificate_expiry:2031-10-06T07:43:17+00:00,cni_driver:calico,cni_tag:v3.31.4,default_security_group:eb2724f0-84af-413f-8a3a-2ff7cfa5e9c3,external_network_id:a858742a-245b-41d3-9a05-617e1b069eb9,external_subnet_id_list:679217b7-1e03-4ce0-8007-7e5cf4f1c47c:05f97c52-a6c3-4422-8b24-31de607fed1c:de20ab10-e4dc-448d-a9a1-cc081fffed9e:5641871e-1d78-4978-a89a-63661e851cba:5cdf279e-1335-4b75-b802-0f5c75337b54:7c5aabaa-63ea-410d-8644-20185b7e31c8:f7038b14-a310-4eb6-92ed-a7d9551b82d7:22011c23-653a-425b-b093-7bf3fb25574d:fc06aecc-b0f9-48ad-9034-62a411c8cc9a:be54270f-4858-413c-b448-135dc37521da:a107a4f8-0e05-44ca-8574-c311f36a9c96:cb12a663-9938-4aa5-85aa-0945d74c4462:995add64-8bda-468e-af0f-533c00888f5a:85558a25-94b7-459d-b8ba-53740c9310db:cf714aed-b3c6-45eb-8867-08dce9f27f92:6d8e37b1-4c4b-4ed6-b34c-b2e12957d00a:d832f723-5fe6-4f55-9493-5fe9467cc259:7f6c2ade-1b11-433b-ad6c-3edefbfc038b:91ab5eda-69a2-4312-8968-ec48bf7eaef0:714c8fe2-d064-47f6-8451-9fbcc388817d:3e4a8dea-4287-4039-9f52-7317a1065723:0af1bb82-f31c-4d74-9108-f93001bc21f2:26e80586-ce00-4b45-9f23-a5f8a27485b1:4b4ed6ce-5f41-4d3c-8363-9e9928cdf51e:697edd38-a62e-4ed2-a225-e27d0a40ce89:e70a4a7f-d58d-497e-9161-ec75415c0724:170c7e54-3710-4889-a726-f8fa868b191d:e4375237-baba-4a41-a928-b4693ba0e7a3:d1e765e5-3f2e-49cf-bf9c-c6cbe5a2e496:d3e22864-b278-4923-9af5-d4f07d2ee0fb:7e486769-a377-46aa-8ec3-f0e4c2dd6c30:11133113-1a2c-4a67-b06f-0b4cf9809e92:8c4e3353-266b-4d0d-b35f-997659a154bc:62733ccb-1ab7-4913-b407-62e0381b5a6b:161335a4-c65b-47a6-bd87-4205b75ce7ca:d73e0bcc-4bce-4e3d-ade8-8b041cbf8232:0c844317-a014-4965-9bc8-c5588d668698:27de04d3-8984-4430-843d-bcfd6f821fdf:24d9fcc8-5725-4da5-9aa4-ad7707b1cfed:6fe5086b-ef56-4454-8aee-fe9d169bfdbf:0bce294a-b772-4c6c-97c9-25e3c2f2ce72:a6f27f47-166a-4a71-b09d-ae2d3ff0aedf:e8951dff-aa5a-4619-bebe-fc919a92866b:42ae5884-70bf-4232-8baa-f514ac280775:a2551567-b464-424c-ab13-8dcaf1871348:0641f8ac-c7e9-43a8-9eb5-ba63d08b83e0:9217260d-35df-4c69-af35-340a721533a8:eeb4ec59-4886-4283-94d9-7192ee41bd65:716da39f-19cb-4adb-a533-8a040232e4fd:fb3d25ea-3f60-4928-b56a-ca152f8eceee,k8s_args:{\\kube-apiserver/default-not-ready-toleration-seconds\\: 300, \\kube-apiserver/default-unreachable-toleration-seconds\\: 300, \\kube-controller-manager/node-monitor-grace-period\\: 40, \\kube-controller-manager/unhealthy-zone-threshold\\: 55},kube_tag:v1.33.4,kube_version_status:NEED_K8S_UPGRADE,master_lb_floating_ip_enabled:True,master_metric_agent:telegraf,nks_registry_url:dfe965c3-kr1-registry.container.nhncloud.com/container_service,node_image:107cc02d-02d8-44fd-84d6-6c316045b817,platform_version:1.202608.0,pods_network_cidr:10.100.0.0/16,pods_network_subnet:24,project_domain:NORMAL,service_cluster_ip_range:10.254.0.0/16,service_user_enabled:True,station_id:default,strict_sg_rules:False,term_of_validity:5,upgradable_kube_versions:[\\v1.34.3\\],upgradable_platform_version:}"
      },
      {
        "key": "Links",
        "value": "{href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/v1/clusters/3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7,rel:self}; {href:https://kr1-api-kubernetes-infrastructure.nhncloudservice.com/clusters/3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7,rel:bookmark}"
      },
      {
        "key": "MasterCount",
        "value": "3"
      },
      {
        "key": "Name",
        "value": "tbqarbc38u4jdhjfp8ho"
      },
      {
        "key": "NodeAddresses",
        "value": "10.0.1.104; 10.0.1.6"
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
        "value": "3cb524cc-8d6c-4e4e-99f4-893569fbac7e"
      },
      {
        "key": "Status",
        "value": "CREATE_COMPLETE"
      },
      {
        "key": "UUID",
        "value": "3e1392d2-b4b2-43eb-ba42-5cc6c868c1d7"
      },
      {
        "key": "UpdatedAt",
        "value": "2026-10-07T08:02:03Z"
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
        "value": "d3c022a0-6f2c-42c3-9b2f-c15766f4ca04"
      },
      {
        "key": "FixedSubnet",
        "value": "78c23b29-ad0d-4621-8095-87ec9f253010"
      },
      {
        "key": "HealthStatus",
        "value": "FRESH"
      },
      {
        "key": "HealthStatusReason",
        "value": "{api:OK,cluster.api_status:NORMAL,cluster.node_status:NORMAL,nodegroup-stats.default-worker:2:0,nodegroup.node_status.default-worker:NORMAL,timestamp:2026-10-07T17:02:02.663710}"
      }
    ]
  }
}
```

</details>

