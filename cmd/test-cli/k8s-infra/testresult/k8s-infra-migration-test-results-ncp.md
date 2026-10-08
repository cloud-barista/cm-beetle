# CM-Beetle K8s Infra Migration Test Results — NCP-Seoul

> [!NOTE]
> Full lifecycle against a real CSP: recommend → validate → migrate → list → get → workload → delete → residual.

## Environment

- CSP / Region: ncp / kr
- CM-Beetle URL: http://localhost:8056
- CM-Beetle Version: v0.6.1+ (cd1f3a2)
- Git Commit: cd1f3a2
- Namespace: mig01
- Test Date: 2026-10-07 16:42:05 KST
- Cluster ID: mig07-on-prem-k8s-cluster

## Test Results Summary

| Step | Description | Status | Duration |
|------|-------------|--------|----------|
| 1 | POST /recommendation/k8sCluster | ✅ **PASS** | 19ms |
| 2 | POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation) | ✅ **PASS** | 28ms |
| 3 | POST /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 12m15.221s |
| 4 | GET /migration/ns/{nsId}/k8sCluster | ✅ **PASS** | 3ms |
| 5 | GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation | ✅ **PASS** | 2ms |
| 5 | Workload verification (kubeconfig -> K8s API -> nginx) | ✅ **PASS** | 1m12.398s |
| 7 | DELETE /migration/ns/{nsId}/k8sCluster/{id} | ✅ **PASS** | 9m37.164s |
| 8 | Residual resource check (Tumblebug) | ✅ **PASS** | 4ms |

**Overall Result**: 8/8 steps passed ✅

**Total Duration**: 23m4s

---

## Step Details

### Step 1 — POST /recommendation/k8sCluster

- **Duration**: 19ms
- **Status Code**: 200

- ℹ️  cluster: on-prem-k8s-cluster (version 1.34.3-nks.2)
- ℹ️  node groups: 1
- ℹ️  node group[0] "workers1" spec=ncp+kr+s4-g3a nodes=2

### Step 2 — POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)

- **Duration**: 28ms
- **Status Code**: 200

- ✅ Target K8s infra model is valid for migration (0 issues)

### Step 3 — POST /migration/ns/{nsId}/k8sCluster

- **Duration**: 12m15.221s
- **Status Code**: 202

- ℹ️  nameSeed: mig07
- ℹ️  async reqId: 1791358925722624156
- ℹ️  cluster id: mig07-on-prem-k8s-cluster
- ℹ️  elapsed: 12m15s
- ✅ status: Active

### Step 4 — GET /migration/ns/{nsId}/k8sCluster

- **Duration**: 3ms
- **Status Code**: 200

- ✅ migrated cluster present in list (6 total)

### Step 5 — GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation

- **Duration**: 2ms
- **Status Code**: 200

- ✅ status: Active
- ✅ node group count matches recommendation: 1
- ✅ node group "workers1" matches (spec=ncp+kr+s4-g3a, nodes=2)
- ✅ version: 1.34.3-nks.2 (recommended 1.34.3-nks.2)

### Step 5 — Workload verification (kubeconfig -> K8s API -> nginx)

- **Duration**: 1m12.398s

- ✅ kubeconfig obtained (server: https://5864c959-453e-41c2-a909-bad33c7b0949.kr.vnks.ntruss.com)
- ℹ️  auth method: exec credential plugin
- ✅ cluster token obtained from Tumblebug
- ✅ API server reachable (v1.34.3)
- ✅ 2 node(s) Ready, matching the recommendation
- ✅ nginx Deployment created
- ✅ nginx pod Running (attempt 2)
- ✅ LoadBalancer Service created
- ✅ LoadBalancer address assigned: default-beetle-test-ngin-2d64e-146081322-5f08732b7bd8.kr.lb.naverncp.com
- ✅ nginx served over the LoadBalancer at http://default-beetle-test-ngin-2d64e-146081322-5f08732b7bd8.kr.lb.naverncp.com/ (attempt 1)
- ✅ LoadBalancer Service removed
- ✅ nginx Deployment removed

### Step 7 — DELETE /migration/ns/{nsId}/k8sCluster/{id}

- **Duration**: 9m37.164s
- **Status Code**: 200

- ✅ deleted on attempt 1 (9m37s)

### Step 8 — Residual resource check (Tumblebug)

- **Duration**: 4ms

- ℹ️  VNet mig07-k8s-vpc still exists (known gap)
- ℹ️  SecurityGroup mig07-k8s-sg still exists (known gap)
- ℹ️  SshKey mig07-k8s-sshkey still exists (known gap)

## Recommendation (input to migration)

<details>
  <summary> <ins>Click to see the recommendation</ins> </summary>

```json
{
  "status": "recommended",
  "description": "K8s cluster recommendation for ncp kr (source: v1.32.3 → target: v1.34.3-nks.2)",
  "targetCloud": {
    "csp": "ncp",
    "region": "kr"
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
    "connectionName": "ncp-kr",
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
    "connectionName": "ncp-kr",
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
      "connectionName": "ncp-kr",
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
    "connectionName": "ncp-kr",
    "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "name": "on-prem-k8s-cluster",
    "version": "1.34.3-nks.2",
    "vNetId": "",
    "subnetIds": null,
    "securityGroupIds": null,
    "k8sNodeGroupList": [
      {
        "name": "workers1",
        "imageId": "default",
        "specId": "ncp+kr+s4-g3a",
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
  "id": "mig07-on-prem-k8s-cluster",
  "uid": "tbvr2bql5thcrlb3ki68",
  "name": "mig07-on-prem-k8s-cluster",
  "connectionName": "ncp-kr",
  "connectionConfig": {
    "configName": "ncp-kr",
    "providerName": "ncp",
    "driverName": "ncp-driver-v1.0.so",
    "credentialName": "ncp",
    "credentialHolder": "admin",
    "regionZoneInfoName": "ncp-kr",
    "regionZoneInfo": {
      "assignedRegion": "KR",
      "assignedZone": "KR-1"
    },
    "regionDetail": {
      "regionId": "KR",
      "regionName": "kr",
      "description": "Korea 1",
      "location": {
        "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
        "latitude": 37.4754,
        "longitude": 126.8831
      },
      "zones": [
        "KR-1",
        "KR-2"
      ]
    },
    "regionRepresentative": true,
    "verified": true
  },
  "description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
  "systemMessage": "",
  "label": {
    "sys.connectionName": "ncp-kr",
    "sys.createdTime": "2026-10-07 07:43:08 +0000 UTC",
    "sys.cspResourceId": "5864c959-453e-41c2-a909-bad33c7b0949",
    "sys.cspResourceName": "tbvr2bql5thcrlb3ki68",
    "sys.description": "Migrated from on-premise K8s cluster (v1.32.3, 2 workers)",
    "sys.id": "mig07-on-prem-k8s-cluster",
    "sys.labelType": "k8s",
    "sys.manager": "cb-tumblebug",
    "sys.name": "mig07-on-prem-k8s-cluster",
    "sys.namespace": "mig01",
    "sys.uid": "tbvr2bql5thcrlb3ki68",
    "sys.version": "1.34.3-nks.2"
  },
  "systemLabel": "",
  "version": "1.34.3-nks.2",
  "network": {
    "vNetId": "mig07-k8s-vpc",
    "subnetIds": [
      "mig07-k8s-subnet-a"
    ],
    "securityGroupIds": [
      "mig07-k8s-sg"
    ],
    "keyValueList": null
  },
  "k8sNodeGroupList": [
    {
      "id": "workers1",
      "name": "workers1",
      "imageId": "default",
      "specId": "ncp+kr+s4-g3a",
      "rootDiskType": "",
      "rootDiskSize": 100,
      "sshKeyId": "mig07-k8s-sshkey",
      "onAutoScaling": false,
      "desiredNodeSize": 2,
      "minNodeSize": 0,
      "maxNodeSize": 0,
      "status": "Active",
      "k8sNodes": [
        {
          "cspResourceName": "workers1-w-f4bb",
          "cspResourceId": "146081105"
        },
        {
          "cspResourceName": "workers1-w-8e09",
          "cspResourceId": "146081108"
        }
      ],
      "keyValueList": [
        {
          "key": "InstanceNo",
          "value": "146081042"
        },
        {
          "key": "Status",
          "value": "RUN"
        },
        {
          "key": "ServerSpecCode",
          "value": "s4-g3a"
        },
        {
          "key": "SoftwareCode",
          "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.WRKND.G003"
        },
        {
          "key": "AutoScalingEnabled",
          "value": "false"
        },
        {
          "key": "AutoScalingMin",
          "value": "0"
        },
        {
          "key": "AutoScalingMax",
          "value": "0"
        }
      ],
      "cspResourceName": "workers1",
      "cspResourceId": "146081042",
      "spiderViewK8sNodeGroupDetail": {
        "IId": {
          "NameId": "workers1",
          "SystemId": "146081042"
        },
        "ImageIID": {
          "NameId": "",
          "SystemId": ""
        },
        "VMSpecName": "s4-g3a",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tbv1vv58f0quobrsetla",
          "SystemId": ""
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "workers1-w-f4bb",
            "SystemId": "146081105"
          },
          {
            "NameId": "workers1-w-8e09",
            "SystemId": "146081108"
          }
        ],
        "KeyValueList": [
          {
            "key": "InstanceNo",
            "value": "146081042"
          },
          {
            "key": "Status",
            "value": "RUN"
          },
          {
            "key": "ServerSpecCode",
            "value": "s4-g3a"
          },
          {
            "key": "SoftwareCode",
            "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.WRKND.G003"
          },
          {
            "key": "AutoScalingEnabled",
            "value": "false"
          },
          {
            "key": "AutoScalingMin",
            "value": "0"
          },
          {
            "key": "AutoScalingMax",
            "value": "0"
          }
        ]
      }
    }
  ],
  "accessInfo": {
    "endpoint": "https://5864c959-453e-41c2-a909-bad33c7b0949.kr.vnks.ntruss.com",
    "kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://5864c959-453e-41c2-a909-bad33c7b0949.kr.vnks.ntruss.com\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM2VENDQWRHZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQ0FYRFRJMk1UQXdOekEzTkRNeU0xb1lEekl4TWpZd09URXpNRGMwTXpJeldqQVZNUk13RVFZRApWUVFERXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBCm8rOERhNjExVitQY3U4RXJiSHA5OHBhODFBaUdTcGZlZ3RZaGc3azIydkRjVEFNNFAxcDhRUHIrSE0yR2NkM0wKZmpXTnVESGEvM1RGb01NaFdYNlhmOUdDNmhiVzViSVJhOVViZ2Z1d0dLZHhkLzd4Yk5nRUI2N1ppOFlDbndJQwpUb2d4YXlNazdVZkFXWW9jQm9KQW9oellxTzVTOFNwQVU4dzQyejhvUmV2MHZLdFNGblFsMGlSTU40d012azRECmQvT2l5UjRWK2lzR3hQL0ZxYW1haFRUS1Q2cFZFUU8yeWYvUEMzbmwzOUJyaEY0eWppNTNxWVRNUzRLbDVmdTAKMGtlRDFSU1B6M0NqV2FNZUhOTFJHT0lLaTYyeUk0b3ZiMytmSEwrNW16Z3pWWldrMGxmamN4d2xjVVVEU2JLVAo2TGVRZXBzZVg0WU13WksyYU9IK1p3SURBUUFCbzBJd1FEQU9CZ05WSFE4QkFmOEVCQU1DQXFRd0R3WURWUjBUCkFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVuREJVR1RtWU54ZWM1SDk1M016b3E0Qk43d0V3RFFZSktvWkkKaHZjTkFRRUxCUUFEZ2dFQkFBNE5Zdmg1bGxmZ3NxdC82WG90WWJaY0djRll1ZXBUME9QQVlKNEljN3BnSk44bgoyWUhtbk9VRmxkL0d0WmNUQUpzeGpRelZ4cTJSQlhxbnNiRjVLaXBWcTNCMmRvZENEUFdSU0NVRG5SK3lnTVIyCktZY1dSSDZwMjBOR2RWUWV4TjNNYVBxcExSYmMvV2ticlkxelppZklaTUcycVZHQ2FNUTNDT1daVjFESUF0TWUKZStiSUhYVm4vb2dka0xHT0ZXWDhqdjJqYVhFanNTUlhmNjBIYURNWkZsdHdrRFRieTd2eHhUWC9HVE5HbjhmdApzZGc3SzNLblI1SmhMMFRnYzdyZ3hQdW1KWjRBL1RvV3VjNDlXYXoxQTBHZEQwQ0F4M3NNTlZWNUE3TDBBdVFpCi80Zk5CWG9QR3lUU0t0aEFKUlJHeUswMUQzTG1tdlBxUk9HbzlVST0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\n  name: tbvr2bql5thcrlb3ki68\ncontexts:\n- context:\n    cluster: tbvr2bql5thcrlb3ki68\n    user: ncp-dynamic-token\n  name: tbvr2bql5thcrlb3ki68\ncurrent-context: tbvr2bql5thcrlb3ki68\nusers:\n- name: ncp-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tbvr2bql5thcrlb3ki68/token?ConnectionName=ncp-kr\\\"\"\n"
  },
  "addons": {
    "keyValueList": null
  },
  "status": "Active",
  "createdTime": "2026-10-07T07:43:08Z",
  "keyValueList": [
    {
      "key": "Status",
      "value": "RUNNING"
    },
    {
      "key": "Uuid",
      "value": "5864c959-453e-41c2-a909-bad33c7b0949"
    },
    {
      "key": "VpcNo",
      "value": "150011"
    },
    {
      "key": "Endpoint",
      "value": "https://5864c959-453e-41c2-a909-bad33c7b0949.kr.vnks.ntruss.com"
    },
    {
      "key": "K8sVersion",
      "value": "1.34.3-nks.2"
    },
    {
      "key": "HypervisorCode",
      "value": "KVM"
    },
    {
      "key": "ClusterType",
      "value": "SVR.VNKS.STAND.C004.M016.G003"
    },
    {
      "key": "AcgName",
      "value": "nks-36624-2ez0ul"
    },
    {
      "key": "AcgNo",
      "value": "404644"
    }
  ],
  "cspResourceName": "tbvr2bql5thcrlb3ki68",
  "cspResourceId": "5864c959-453e-41c2-a909-bad33c7b0949",
  "spiderViewK8sClusterDetail": {
    "IId": {
      "NameId": "tbvr2bql5thcrlb3ki68",
      "SystemId": "5864c959-453e-41c2-a909-bad33c7b0949"
    },
    "Version": "1.34.3-nks.2",
    "Network": {
      "VpcIID": {
        "NameId": "tbn3t6ja32ecsj7bbghl",
        "SystemId": "150011"
      },
      "SubnetIIDs": [
        {
          "NameId": "tbsm5vemad8va3o8v2u9",
          "SystemId": "328830"
        }
      ],
      "SecurityGroupIIDs": [
        {
          "NameId": "nks-36624-2ez0ul",
          "SystemId": "404644"
        }
      ],
      "KeyValueList": null
    },
    "NodeGroupList": [
      {
        "IId": {
          "NameId": "workers1",
          "SystemId": "146081042"
        },
        "ImageIID": {
          "NameId": "",
          "SystemId": ""
        },
        "VMSpecName": "s4-g3a",
        "RootDiskSize": "100",
        "KeyPairIID": {
          "NameId": "tbv1vv58f0quobrsetla",
          "SystemId": ""
        },
        "OnAutoScaling": false,
        "DesiredNodeSize": 2,
        "MinNodeSize": 0,
        "MaxNodeSize": 0,
        "Status": "Active",
        "Nodes": [
          {
            "NameId": "workers1-w-f4bb",
            "SystemId": "146081105"
          },
          {
            "NameId": "workers1-w-8e09",
            "SystemId": "146081108"
          }
        ],
        "KeyValueList": [
          {
            "key": "InstanceNo",
            "value": "146081042"
          },
          {
            "key": "Status",
            "value": "RUN"
          },
          {
            "key": "ServerSpecCode",
            "value": "s4-g3a"
          },
          {
            "key": "SoftwareCode",
            "value": "SW.VSVR.OS.LNX64.UBNTU.SVR22.WRKND.G003"
          },
          {
            "key": "AutoScalingEnabled",
            "value": "false"
          },
          {
            "key": "AutoScalingMin",
            "value": "0"
          },
          {
            "key": "AutoScalingMax",
            "value": "0"
          }
        ]
      }
    ],
    "AccessInfo": {
      "Endpoint": "https://5864c959-453e-41c2-a909-bad33c7b0949.kr.vnks.ntruss.com",
      "Kubeconfig": "apiVersion: v1\nkind: Config\nclusters:\n- cluster:\n    server: https://5864c959-453e-41c2-a909-bad33c7b0949.kr.vnks.ntruss.com\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM2VENDQWRHZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQ0FYRFRJMk1UQXdOekEzTkRNeU0xb1lEekl4TWpZd09URXpNRGMwTXpJeldqQVZNUk13RVFZRApWUVFERXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBCm8rOERhNjExVitQY3U4RXJiSHA5OHBhODFBaUdTcGZlZ3RZaGc3azIydkRjVEFNNFAxcDhRUHIrSE0yR2NkM0wKZmpXTnVESGEvM1RGb01NaFdYNlhmOUdDNmhiVzViSVJhOVViZ2Z1d0dLZHhkLzd4Yk5nRUI2N1ppOFlDbndJQwpUb2d4YXlNazdVZkFXWW9jQm9KQW9oellxTzVTOFNwQVU4dzQyejhvUmV2MHZLdFNGblFsMGlSTU40d012azRECmQvT2l5UjRWK2lzR3hQL0ZxYW1haFRUS1Q2cFZFUU8yeWYvUEMzbmwzOUJyaEY0eWppNTNxWVRNUzRLbDVmdTAKMGtlRDFSU1B6M0NqV2FNZUhOTFJHT0lLaTYyeUk0b3ZiMytmSEwrNW16Z3pWWldrMGxmamN4d2xjVVVEU2JLVAo2TGVRZXBzZVg0WU13WksyYU9IK1p3SURBUUFCbzBJd1FEQU9CZ05WSFE4QkFmOEVCQU1DQXFRd0R3WURWUjBUCkFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVuREJVR1RtWU54ZWM1SDk1M016b3E0Qk43d0V3RFFZSktvWkkKaHZjTkFRRUxCUUFEZ2dFQkFBNE5Zdmg1bGxmZ3NxdC82WG90WWJaY0djRll1ZXBUME9QQVlKNEljN3BnSk44bgoyWUhtbk9VRmxkL0d0WmNUQUpzeGpRelZ4cTJSQlhxbnNiRjVLaXBWcTNCMmRvZENEUFdSU0NVRG5SK3lnTVIyCktZY1dSSDZwMjBOR2RWUWV4TjNNYVBxcExSYmMvV2ticlkxelppZklaTUcycVZHQ2FNUTNDT1daVjFESUF0TWUKZStiSUhYVm4vb2dka0xHT0ZXWDhqdjJqYVhFanNTUlhmNjBIYURNWkZsdHdrRFRieTd2eHhUWC9HVE5HbjhmdApzZGc3SzNLblI1SmhMMFRnYzdyZ3hQdW1KWjRBL1RvV3VjNDlXYXoxQTBHZEQwQ0F4M3NNTlZWNUE3TDBBdVFpCi80Zk5CWG9QR3lUU0t0aEFKUlJHeUswMUQzTG1tdlBxUk9HbzlVST0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\n  name: tbvr2bql5thcrlb3ki68\ncontexts:\n- context:\n    cluster: tbvr2bql5thcrlb3ki68\n    user: ncp-dynamic-token\n  name: tbvr2bql5thcrlb3ki68\ncurrent-context: tbvr2bql5thcrlb3ki68\nusers:\n- name: ncp-dynamic-token\n  user:\n    exec:\n      apiVersion: client.authentication.k8s.io/v1\n      interactiveMode: Never\n      command: sh\n      args:\n      - -c\n      - \". ~/.cb-spider/.spider-credential \u0026\u0026 curl -s -u \\\"$SPIDER_USERNAME:$SPIDER_PASSWORD\\\" \\\"http://0.0.0.0:1024/spider/cluster/tbvr2bql5thcrlb3ki68/token?ConnectionName=ncp-kr\\\"\"\n"
    },
    "Addons": {
      "KeyValueList": null
    },
    "Status": "Active",
    "CreatedTime": "2026-10-07T07:43:08Z",
    "KeyValueList": [
      {
        "key": "Status",
        "value": "RUNNING"
      },
      {
        "key": "Uuid",
        "value": "5864c959-453e-41c2-a909-bad33c7b0949"
      },
      {
        "key": "VpcNo",
        "value": "150011"
      },
      {
        "key": "Endpoint",
        "value": "https://5864c959-453e-41c2-a909-bad33c7b0949.kr.vnks.ntruss.com"
      },
      {
        "key": "K8sVersion",
        "value": "1.34.3-nks.2"
      },
      {
        "key": "HypervisorCode",
        "value": "KVM"
      },
      {
        "key": "ClusterType",
        "value": "SVR.VNKS.STAND.C004.M016.G003"
      },
      {
        "key": "AcgName",
        "value": "nks-36624-2ez0ul"
      },
      {
        "key": "AcgNo",
        "value": "404644"
      }
    ]
  }
}
```

</details>

