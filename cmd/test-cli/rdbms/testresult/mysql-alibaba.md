# Managed RDBMS (MySQL) Test Report: ALIBABA (ap-northeast-2)

- **Test Case:** Alibaba AP-Northeast-2 (Seoul) MySQL Test
- **Access Mode:** Public (publicAccess=true)
- **Date & Time:** 2026-11-25 18:54:32
- **Namespace:** `default`
- **Total Duration:** 19m24.735s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** ALIBABA
- **Target Region:** `ap-northeast-2`
- **Namespace:** `default`
- **Test Date:** 2026-11-25 18:54:32

### Scenario & Tested APIs
1. **Pre-flight Spec & Image Review**: `POST /tumblebug/specImagePairReview`
2. **Create Pre-requisite Infra (VNet/SG)**: `POST /tumblebug/ns/{nsId}/resources/vNet`, `POST /tumblebug/ns/{nsId}/resources/securityGroup`
3. **Get RDBMS Support Matrix**: `GET /beetle/recommendation/middleware/rdbms/support`
4. **Get Real-time Capability**: `GET /beetle/recommendation/middleware/rdbms/capability`
5. **Recommend Managed RDBMS**: `POST /beetle/recommendation/middleware/rdbms`
6. **Validate Recommendation**: `POST /beetle/recommendation/middleware/rdbms/validate`
7. **Migrate RDBMS (Provisioning)**: `POST /beetle/migration/middleware/ns/{nsId}/rdbms`
8. **Get RDBMS Info & List**: `GET /beetle/migration/middleware/ns/{nsId}/rdbms`
9. **Create Logical Database**: `POST /beetle/migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database`
10. **Get Secure Transport Info**: `GET /beetle/migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/secure-transport`
11. **External Data I/O**: Direct TCP/SQL connectivity & TLS handshake test
12. **Internal Data I/O**: SQL execution via internal Runner VM (`POST /tumblebug/ns/{nsId}/infra`)
13. **Delete Logical Database**: `DELETE /beetle/migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database/{databaseName}`
14. **Delete RDBMS**: `DELETE /beetle/migration/middleware/ns/{nsId}/rdbms/{rdbmsId}`
15. **Delete Pre-requisite SG & VNet**: `DELETE /tumblebug/ns/{nsId}/resources/securityGroup/{sgId}`, `DELETE /tumblebug/ns/{nsId}/resources/vNet/{vNetId}`

## Execution Steps & API Traces

### 1. Tumblebug POST /specImagePairReview (Pre-flight Spec & Image Review) [✅ SUCCESS]
- **Duration:** 4.037s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "ubuntu_24_04_x64_20G_alibase_20260916.vhd",
  "specId": "alibaba+ap-northeast-2+ecs.e-c1m2.large"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "cached": true,
    "instanceType": "ecs.e-c1m2.large",
    "provider": "alibaba",
    "queriedAt": "2026-11-25T09:15:43.113602083Z",
    "refreshAfter": "2026-11-25T10:15:43.113602083Z",
    "region": "ap-northeast-2",
    "source": "alibaba:DescribeAvailableResource",
    "status": "Available",
    "ttlSeconds": 3600,
    "zones": [
      {
        "available": true,
        "status": "Available",
        "supportedDisks": [
          "cloud_auto",
          "cloud_essd"
        ],
        "zoneId": "ap-northeast-2b"
      },
      {
        "available": true,
        "status": "Available",
        "supportedDisks": [
          "cloud_essd_entry",
          "cloud_auto",
          "cloud_essd"
        ],
        "zoneId": "ap-northeast-2a"
      },
      {
        "available": true,
        "status": "Available",
        "supportedDisks": [
          "cloud_auto",
          "cloud_essd"
        ],
        "zoneId": "ap-northeast-2c"
      }
    ]
  },
  "connectionName": "alibaba-ap-northeast-2",
  "estimatedCost": "$0.0356/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "alibaba-us-west-1",
    "creationDate": "",
    "cspImageName": "ubuntu_24_04_x64_20G_alibase_20261016.vhd",
    "description": "Kernel version is 6.8.0-139-generic, 2026.9.17",
    "details": [
      {
        "key": "BootMode",
        "value": "UEFI-Preferred"
      },
      {
        "key": "ImageId",
        "value": "ubuntu_24_04_x64_20G_alibase_20260916.vhd"
      },
      {
        "key": "ImageOwnerAlias",
        "value": "system"
      },
      {
        "key": "OSName",
        "value": "Ubuntu  24.04 64位"
      },
      {
        "key": "OSNameEn",
        "value": "Ubuntu  24.04 64 bit"
      },
      {
        "key": "ImageFamily",
        "value": "acs:ubuntu_24_04_x64"
      },
      {
        "key": "Architecture",
        "value": "x86_64"
      },
      {
        "key": "IsSupportIoOptimized",
        "value": "true"
      },
      {
        "key": "Size",
        "value": "20"
      },
      {
        "key": "Description",
        "value": "Kernel version is 6.8.0-139-generic, 2026.9.17"
      },
      {
        "key": "Usage",
        "value": "instance"
      },
      {
        "key": "IsCopied",
        "value": "false"
      },
      {
        "key": "LoginAsNonRootSupported",
        "value": "true"
      },
      {
        "key": "ImageVersion",
        "value": "v2026.9.17"
      },
      {
        "key": "OSType",
        "value": "linux"
      },
      {
        "key": "IsSubscribed",
        "value": "false"
      },
      {
        "key": "IsSupportCloudinit",
        "value": "true"
      },
      {
        "key": "CreationTime",
        "value": "2026-09-17T05:15:31Z"
      },
      {
        "key": "Progress",
        "value": "100%"
      },
      {
        "key": "Platform",
        "value": "Ubuntu"
      },
      {
        "key": "ImageName",
        "value": "ubuntu_24_04_x64_20G_alibase_20260916.vhd"
      },
      {
        "key": "Status",
        "value": "Available"
      },
      {
        "key": "ImageOwnerId",
        "value": "0"
      },
      {
        "key": "IsPublic",
        "value": "true"
      },
      {
        "key": "DetectionOptions",
        "value": "{Status:,Items:{Item:null}}"
      },
      {
        "key": "Features",
        "value": "{MemoryOnlineUpgrade:unsupported,NvmeSupport:supported,CpuOnlineDowngrade:unsupported,ImdsSupport:v2,MemoryOnlineDowngrade:unsupported,CpuOnlineUpgrade:unsupported}"
      },
      {
        "key": "Tags",
        "value": "{Tag:[]}"
      },
      {
        "key": "DiskDeviceMappings",
        "value": "{DiskDeviceMapping:[]}"
      }
    ],
    "fetchedTime": "2026.10.29 04:25:32 Thu",
    "id": "ubuntu_24_04_x64_20G_alibase_20260916.vhd",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "ubuntu_24_04_x64_20G_alibase_20260916.vhd",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 20,
    "osDiskType": "NA",
    "osDistribution": "Ubuntu  24.04 64 bit",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "alibaba",
    "regionList": [
      "ap-northeast-1",
      "ap-northeast-2",
      "ap-southeast-3",
      "ap-southeast-5",
      "ap-southeast-6",
      "ap-southeast-7",
      "ap-southeast-8",
      "cn-beijing",
      "cn-fuzhou",
      "cn-guangzhou",
      "cn-heyuan",
      "cn-huhehaote",
      "cn-nanjing",
      "cn-shanghai",
      "cn-shenzhen",
      "cn-wuhan-lr",
      "cn-zhangjiakou",
      "cn-zhongwei",
      "eu-central-1",
      "eu-west-1",
      "eu-west-2",
      "me-central-1",
      "me-east-1",
      "na-south-1",
      "us-east-1",
      "us-west-1"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tbovbcl1usnmuihjr6f5"
  },
  "imageId": "ubuntu_24_04_x64_20G_alibase_20260916.vhd",
  "imageValidation": {
    "cspResourceId": "ubuntu_24_04_x64_20G_alibase_20261016.vhd",
    "isAvailable": true,
    "resourceId": "ubuntu_24_04_x64_20G_alibase_20260916.vhd",
    "resourceName": "ubuntu_24_04_x64_20G_alibase_20260916.vhd",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "alibaba",
  "regionName": "ap-northeast-2",
  "specDetails": {
    "architecture": "x86_64",
    "availabilityCheckedAt": "2026-11-25T09:15:43.113602083Z",
    "availabilityRefreshAfter": "2026-11-25T10:15:43.113602083Z",
    "availabilityStatus": "Available",
    "availableZones": [
      "ap-northeast-2a",
      "ap-northeast-2b",
      "ap-northeast-2c"
    ],
    "connectionName": "alibaba-ap-northeast-2",
    "costPerHour": 0.0356,
    "cspSpecName": "ecs.e-c1m2.large",
    "details": [
      {
        "key": "CpuArchitecture",
        "value": "X86"
      }
    ],
    "diskSizeGB": -1,
    "evaluationScore01": -1,
    "evaluationScore02": -1,
    "evaluationScore03": -1,
    "evaluationScore04": -1,
    "evaluationScore05": -1,
    "evaluationScore06": -1,
    "evaluationScore07": -1,
    "evaluationScore08": -1,
    "evaluationScore09": -1,
    "evaluationScore10": -1,
    "id": "alibaba+ap-northeast-2+ecs.e-c1m2.large",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "alibaba+ap-northeast-2+ecs.e-c1m2.large",
    "namespace": "system",
    "providerName": "alibaba",
    "regionLatitude": 37.36,
    "regionLongitude": 126.78,
    "regionName": "ap-northeast-2",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbo04el2f2u3pljoa4hq",
    "vCPU": 2
  },
  "specId": "alibaba+ap-northeast-2+ecs.e-c1m2.large",
  "specValidation": {
    "cspResourceId": "ecs.e-c1m2.large",
    "isAvailable": true,
    "resourceId": "alibaba+ap-northeast-2+ecs.e-c1m2.large",
    "resourceName": "ecs.e-c1m2.large",
    "status": "Available"
  },
  "status": "OK",
  "suggestedSystemDisk": "cloud_auto",
  "suggestedZone": "ap-northeast-2b"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 7.722s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.3.0.0/16",
  "connectionName": "alibaba-ap-northeast-2",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-alibaba",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.3.1.0/24",
      "name": "subnet-1",
      "zone": ""
    }
  ]
}
```
```json
// Response Body
{
  "associatedObjectList": null,
  "cidrBlock": "10.3.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T09:54:42Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:42Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:42Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
    }
  ],
  "connectionConfig": {
    "configName": "alibaba-ap-northeast-2",
    "credentialHolder": "admin",
    "credentialName": "alibaba",
    "driverName": "alibaba-driver-v1.0.so",
    "providerName": "alibaba",
    "regionDetail": {
      "description": "South Korea (Seoul)",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "regionId": "ap-northeast-2",
      "regionName": "ap-northeast-2",
      "zones": [
        "ap-northeast-2a",
        "ap-northeast-2b"
      ]
    },
    "regionRepresentative": true,
    "regionZoneInfo": {
      "assignedRegion": "ap-northeast-2",
      "assignedZone": "ap-northeast-2a"
    },
    "regionZoneInfoName": "alibaba-ap-northeast-2",
    "verified": true
  },
  "connectionName": "alibaba-ap-northeast-2",
  "cspResourceId": "vpc-mj7t8uchc521yb50jec8v",
  "cspResourceName": "tb5n4g9nsp6sk04krouf",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-alibaba",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "CreationTime",
      "value": "2026-11-25T09:54:35Z"
    },
    {
      "key": "Status",
      "value": "Available"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7t8uchc521yb50jec8v"
    },
    {
      "key": "IsDefault",
      "value": "false"
    },
    {
      "key": "AdvancedResource",
      "value": "false"
    },
    {
      "key": "OwnerId",
      "value": "5469257408566579"
    },
    {
      "key": "RegionId",
      "value": "ap-northeast-2"
    },
    {
      "key": "VpcName",
      "value": "tb5n4g9nsp6sk04krouf"
    },
    {
      "key": "VRouterId",
      "value": "vrt-mj7j7gytgzr6w12uosubk"
    },
    {
      "key": "CidrBlock",
      "value": "10.3.0.0/16"
    },
    {
      "key": "NetworkAclNum",
      "value": "0"
    },
    {
      "key": "SupportAdvancedFeature",
      "value": "false"
    },
    {
      "key": "ResourceGroupId",
      "value": "rg-acfnvekhilw5kmy"
    },
    {
      "key": "CenStatus",
      "value": "Detached"
    },
    {
      "key": "EnabledIpv6",
      "value": "false"
    },
    {
      "key": "DnsHostnameStatus",
      "value": "DISABLED"
    },
    {
      "key": "VSwitchIds",
      "value": "{VSwitchId:[vsw-mj7yfvd5kk99kumix6i1u]}"
    },
    {
      "key": "SecondaryCidrBlocks",
      "value": "{SecondaryCidrBlock:[]}"
    },
    {
      "key": "UserCidrs",
      "value": "{UserCidr:[]}"
    },
    {
      "key": "NatGatewayIds",
      "value": "{NatGatewayIds:[]}"
    },
    {
      "key": "RouterTableIds",
      "value": "{RouterTableIds:[vtb-mj7x3djrftgxkajouege9]}"
    },
    {
      "key": "Tags",
      "value": "{Tag:null}"
    },
    {
      "key": "Ipv6CidrBlocks",
      "value": "{Ipv6CidrBlock:null}"
    }
  ],
  "name": "test-rdbms-vnet-alibaba",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:54:42Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:54:42Z",
          "reason": "Available",
          "status": "True",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "alibaba-ap-northeast-2",
        "credentialHolder": "admin",
        "credentialName": "alibaba",
        "driverName": "alibaba-driver-v1.0.so",
        "providerName": "alibaba",
        "regionDetail": {
          "description": "South Korea (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionZoneInfoName": "alibaba-ap-northeast-2",
        "verified": true
      },
      "connectionName": "alibaba-ap-northeast-2",
      "cspResourceId": "vsw-mj7yfvd5kk99kumix6i1u",
      "cspResourceName": "tb9et6arq5nb9rmrhj6j",
      "cspVNetId": "vpc-mj7t8uchc521yb50jec8v",
      "cspVNetName": "tb5n4g9nsp6sk04krouf",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.3.1.0/24",
      "keyValueList": [
        {
          "key": "VpcId",
          "value": "vpc-mj7t8uchc521yb50jec8v"
        },
        {
          "key": "Status",
          "value": "Available"
        },
        {
          "key": "CreationTime",
          "value": "2026-11-25T09:54:38Z"
        },
        {
          "key": "IsDefault",
          "value": "false"
        },
        {
          "key": "AvailableIpAddressCount",
          "value": "252"
        },
        {
          "key": "OwnerId",
          "value": "5469257408566579"
        },
        {
          "key": "VSwitchId",
          "value": "vsw-mj7yfvd5kk99kumix6i1u"
        },
        {
          "key": "CidrBlock",
          "value": "10.3.1.0/24"
        },
        {
          "key": "ResourceGroupId",
          "value": "rg-acfnvekhilw5kmy"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "VSwitchName",
          "value": "tb9et6arq5nb9rmrhj6j"
        },
        {
          "key": "EnabledIpv6",
          "value": "false"
        },
        {
          "key": "RouteTable",
          "value": "{ResourceGroupId:,CreationTime:,Status:,RouteTableType:System,VRouterId:,RouteTableId:vtb-mj7x3djrftgxkajouege9,VSwitchIds:{VSwitchId:null},RouteEntrys:{RouteEntry:null}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb9et6arq5nb9rmrhj6j",
      "zone": "ap-northeast-2a"
    }
  ],
  "systemLabel": "",
  "uid": "tb5n4g9nsp6sk04krouf"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 5.979s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "alibaba-ap-northeast-2",
  "description": "Pre-requisite SecurityGroup for CM-Beetle RDBMS test",
  "firewallRules": [
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Ports": "3306",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Ports": "22",
      "Protocol": "TCP"
    }
  ],
  "name": "test-rdbms-sg-alibaba",
  "vNetId": "test-rdbms-vnet-alibaba"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
  "connectionConfig": {
    "configName": "alibaba-ap-northeast-2",
    "credentialHolder": "admin",
    "credentialName": "alibaba",
    "driverName": "alibaba-driver-v1.0.so",
    "providerName": "alibaba",
    "regionDetail": {
      "description": "South Korea (Seoul)",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "regionId": "ap-northeast-2",
      "regionName": "ap-northeast-2",
      "zones": [
        "ap-northeast-2a",
        "ap-northeast-2b"
      ]
    },
    "regionRepresentative": true,
    "regionZoneInfo": {
      "assignedRegion": "ap-northeast-2",
      "assignedZone": "ap-northeast-2a"
    },
    "regionZoneInfoName": "alibaba-ap-northeast-2",
    "verified": true
  },
  "connectionName": "alibaba-ap-northeast-2",
  "cspResourceId": "sg-mj77a670vw20ux8j15li",
  "cspResourceName": "tb7ptjq8kj8lthhdfrs0",
  "description": "Pre-requisite SecurityGroup for CM-Beetle RDBMS test",
  "firewallRules": [
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "22",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "3306",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "outbound",
      "Port": "",
      "Protocol": "ALL"
    }
  ],
  "id": "test-rdbms-sg-alibaba",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "SecurityGroupId",
      "value": "sg-mj77a670vw20ux8j15li"
    },
    {
      "key": "SecurityGroupName",
      "value": "tb7ptjq8kj8lthhdfrs0"
    },
    {
      "key": "Description",
      "value": "tb7ptjq8kj8lthhdfrs0"
    },
    {
      "key": "SecurityGroupType",
      "value": "enterprise"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7t8uchc521yb50jec8v"
    },
    {
      "key": "CreationTime",
      "value": "2026-11-25T09:54:47Z"
    },
    {
      "key": "EcsCount",
      "value": "0"
    },
    {
      "key": "AvailableInstanceAmount",
      "value": "0"
    },
    {
      "key": "ServiceManaged",
      "value": "false"
    },
    {
      "key": "ServiceID",
      "value": "0"
    },
    {
      "key": "RuleCount",
      "value": "3"
    },
    {
      "key": "GroupToGroupRuleCount",
      "value": "0"
    },
    {
      "key": "Tags",
      "value": "{Tag:[]}"
    }
  ],
  "name": "test-rdbms-sg-alibaba",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tb7ptjq8kj8lthhdfrs0",
  "vNetId": "test-rdbms-vnet-alibaba"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 9ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=alibaba`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "alibaba": {
      "dbOperationMethod": "cspNativeApi",
      "storageTypeSelectable": true,
      "supported": true,
      "supportedDBEngines": [
        "mysql",
        "mariadb"
      ],
      "supportsTag": true
    }
  }
}
```

### 5. Beetle GET RDBMS Capability [✅ SUCCESS]
- **Duration:** 3m4.993s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=alibaba-ap-northeast-2&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "7-730",
    "connectionName": "alibaba-ap-northeast-2",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "myduck.x2.2xlarge.xc",
      "myduck.x2.4xlarge.xc",
      "myduck.x2.8xlarge.xc",
      "myduck.x2.large.xc",
      "myduck.x2.xlarge.xc",
      "myduck.x4.2xlarge.xc",
      "myduck.x4.4xlarge.xc",
      "myduck.x4.8xlarge.xc",
      "myduck.x4.large.xc",
      "myduck.x4.xlarge.xc",
      "myduck.x8.2xlarge.xc",
      "myduck.x8.4xlarge.xc",
      "myduck.x8.8xlarge.xc",
      "myduck.x8.large.xc",
      "myduck.x8.xlarge.xc",
      "mysql.n2.large.xc",
      "mysql.n2.medium.2c",
      "mysql.n2.medium.xc",
      "mysql.n2.small.2c",
      "mysql.n2.xlarge.xc",
      "mysql.n2e.large.xc",
      "mysql.n2e.medium.xc",
      "mysql.n2e.small.xc",
      "mysql.n2e.xlarge.xc",
      "mysql.n2m.large.2c",
      "mysql.n2m.medium.2c",
      "mysql.n2m.small.2c",
      "mysql.n2m.xlarge.2c",
      "mysql.n4.large.xc",
      "mysql.n4.medium.xc",
      "mysql.n4.xlarge.xc",
      "mysql.n4e.large.xc",
      "mysql.n4e.medium.xc",
      "mysql.n4e.xlarge.xc",
      "mysql.n4m.large.2c",
      "mysql.n4m.medium.2c",
      "mysql.n4m.xlarge.2c",
      "mysql.n8.large.xc",
      "mysql.n8.medium.xc",
      "mysql.n8.xlarge.xc",
      "mysql.n8e.large.xc",
      "mysql.n8e.medium.xc",
      "mysql.n8e.xlarge.xc",
      "mysql.n8m.large.2c",
      "mysql.n8m.medium.2c",
      "mysql.n8m.xlarge.2c",
      "mysql.x2.13large.2c",
      "mysql.x2.13large.xc",
      "mysql.x2.13xlarge.2c",
      "mysql.x2.13xlarge.xc",
      "mysql.x2.2xlarge.2c",
      "mysql.x2.2xlarge.xc",
      "mysql.x2.3large.2c",
      "mysql.x2.3large.xc",
      "mysql.x2.3xlarge.2c",
      "mysql.x2.3xlarge.xc",
      "mysql.x2.4xlarge.2c",
      "mysql.x2.4xlarge.xc",
      "mysql.x2.6xlarge.2c",
      "mysql.x2.8xlarge.2c",
      "mysql.x2.8xlarge.xc",
      "mysql.x2.large.2c",
      "mysql.x2.large.xc",
      "mysql.x2.medium.2c",
      "mysql.x2.medium.xc",
      "mysql.x2.xlarge.2c",
      "mysql.x2.xlarge.xc",
      "mysql.x2e.2xlarge.xc",
      "mysql.x2e.4xlarge.xc",
      "mysql.x2e.8xlarge.xc",
      "mysql.x2e.large.xc",
      "mysql.x2e.medium.xc",
      "mysql.x2e.xlarge.xc",
      "mysql.x2m.2xlarge.2c",
      "mysql.x2m.4xlarge.2c",
      "mysql.x2m.large.2c",
      "mysql.x2m.medium.2c",
      "mysql.x2m.xlarge.2c",
      "mysql.x4.13large.2c",
      "mysql.x4.13large.xc",
      "mysql.x4.13xlarge.2c",
      "mysql.x4.13xlarge.xc",
      "mysql.x4.16xlarge.2c",
      "mysql.x4.2xlarge.2c",
      "mysql.x4.2xlarge.xc",
      "mysql.x4.3large.2c",
      "mysql.x4.3large.xc",
      "mysql.x4.3xlarge.2c",
      "mysql.x4.3xlarge.xc",
      "mysql.x4.4xlarge.2c",
      "mysql.x4.4xlarge.xc",
      "mysql.x4.6xlarge.2c",
      "mysql.x4.8xlarge.2c",
      "mysql.x4.8xlarge.xc",
      "mysql.x4.large.2c",
      "mysql.x4.large.xc",
      "mysql.x4.medium.2c",
      "mysql.x4.medium.xc",
      "mysql.x4.xlarge.2c",
      "mysql.x4.xlarge.xc",
      "mysql.x4e.2xlarge.xc",
      "mysql.x4e.4xlarge.xc",
      "mysql.x4e.8xlarge.xc",
      "mysql.x4e.large.xc",
      "mysql.x4e.medium.xc",
      "mysql.x4e.xlarge.xc",
      "mysql.x4m.2xlarge.2c",
      "mysql.x4m.4xlarge.2c",
      "mysql.x4m.8xlarge.2c",
      "mysql.x4m.large.2c",
      "mysql.x4m.medium.2c",
      "mysql.x4m.xlarge.2c",
      "mysql.x8.13large.2c",
      "mysql.x8.13large.xc",
      "mysql.x8.13xlarge.2c",
      "mysql.x8.13xlarge.xc",
      "mysql.x8.2xlarge.2c",
      "mysql.x8.2xlarge.xc",
      "mysql.x8.3large.2c",
      "mysql.x8.3large.xc",
      "mysql.x8.3xlarge.2c",
      "mysql.x8.3xlarge.xc",
      "mysql.x8.4xlarge.2c",
      "mysql.x8.4xlarge.xc",
      "mysql.x8.6xlarge.2c",
      "mysql.x8.8xlarge.2c",
      "mysql.x8.8xlarge.xc",
      "mysql.x8.large.2c",
      "mysql.x8.large.xc",
      "mysql.x8.medium.2c",
      "mysql.x8.medium.xc",
      "mysql.x8.xlarge.2c",
      "mysql.x8.xlarge.xc",
      "mysql.x8e.2xlarge.xc",
      "mysql.x8e.4xlarge.xc",
      "mysql.x8e.8xlarge.xc",
      "mysql.x8e.large.xc",
      "mysql.x8e.medium.xc",
      "mysql.x8e.xlarge.xc",
      "mysql.x8m.2xlarge.2c",
      "mysql.x8m.4xlarge.2c",
      "mysql.x8m.8xlarge.2c",
      "mysql.x8m.large.2c",
      "mysql.x8m.medium.2c",
      "mysql.x8m.xlarge.2c",
      "myduck.n2.2xlarge.1",
      "myduck.n2.large.1",
      "myduck.n2.xlarge.1",
      "myduck.n4.2xlarge.1",
      "myduck.n4.large.1",
      "myduck.n4.xlarge.1",
      "myduck.n8.2xlarge.1",
      "myduck.n8.large.1",
      "myduck.n8.xlarge.1",
      "mysql.n1.micro.1",
      "mysql.n1e.medium.1",
      "mysql.n1e.small.1",
      "mysql.n2.large.1",
      "mysql.n2.medium.1",
      "mysql.n2.small.1",
      "mysql.n2.xlarge.1",
      "mysql.n2e.medium.1",
      "mysql.n2e.small.1",
      "mysql.n4.large.1",
      "mysql.n4.medium.1",
      "mysql.n4.xlarge.1"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "32768",
        "name": "myduck.n2.2xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16 "
      },
      {
        "memSizeMiB": "8192",
        "name": "myduck.n2.large.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "myduck.n2.xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "myduck.n4.2xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16 "
      },
      {
        "memSizeMiB": "16384",
        "name": "myduck.n4.large.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "myduck.n4.xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "myduck.n8.2xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16 "
      },
      {
        "memSizeMiB": "32768",
        "name": "myduck.n8.large.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "65536",
        "name": "myduck.n8.xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "2048",
        "name": "mysql.n1e.medium.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "1024",
        "name": "mysql.n1e.small.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "8192",
        "name": "mysql.n2.large.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "4096",
        "name": "mysql.n2.medium.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "mysql.n2.xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "4096",
        "name": "mysql.n2e.medium.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "2048",
        "name": "mysql.n2e.small.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "16384",
        "name": "mysql.n4.large.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "8192",
        "name": "mysql.n4.medium.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "mysql.n4.xlarge.1",
        "storageSizeRangeGB": {
          "max": 64000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      }
    ],
    "dbOperationMethod": "",
    "defaultStorageType": "",
    "liveSupportedEngines": [
      "mysql",
      "mariadb"
    ],
    "notes": {
      "storageTypes": [
        {
          "constraints": "Minimum 20GB storage.",
          "description": "Alibaba Cloud automatically selects optimal SSD storage (General ESSD / PL1) based on region and instance specification. Recommended for simplicity.",
          "displayName": "Auto-selected Storage Type (SSD)",
          "maxSizeGB": 32768,
          "minSizeGB": 20,
          "recommendationLevel": "recommended",
          "recommended": true,
          "storageType": "cloud_auto"
        },
        {
          "constraints": "Minimum 20GB storage.",
          "description": "Standard enhanced SSD with good balance of performance and cost. Performance Level 1.",
          "displayName": "Enhanced SSD (ESSD PL1)",
          "maxSizeGB": 32768,
          "minSizeGB": 20,
          "recommendationLevel": "standard",
          "storageType": "cloud_essd"
        },
        {
          "constraints": "Minimum 500GB storage. Not compatible with dbInstanceSpec(s): mysql.n4.*.",
          "description": "High-performance enhanced SSD. Performance Level 2. Minimum 500 GB storage. Not compatible with mysql.n4.* instance specifications.",
          "displayName": "Enhanced SSD (ESSD PL2)",
          "incompatibleSpecs": [
            "mysql.n4.*"
          ],
          "maxSizeGB": 32768,
          "minSizeGB": 500,
          "recommendationLevel": "premium",
          "storageType": "cloud_essd2"
        },
        {
          "constraints": "Minimum 1500GB storage. Not compatible with dbInstanceSpec(s): mysql.n4.*.",
          "description": "Ultra-high-performance enhanced SSD. Performance Level 3. Minimum 1,500 GB storage. Not compatible with mysql.n4.* instance specifications.",
          "displayName": "Enhanced SSD (ESSD PL3)",
          "incompatibleSpecs": [
            "mysql.n4.*"
          ],
          "maxSizeGB": 32768,
          "minSizeGB": 1500,
          "recommendationLevel": "premium",
          "storageType": "cloud_essd3"
        },
        {
          "description": "Storage type details not yet documented.",
          "displayName": "cloud_ssd",
          "storageType": "cloud_ssd"
        }
      ]
    },
    "providerName": "alibaba",
    "regionName": "ap-northeast-2",
    "requiresSecurityGroup": false,
    "requiresSubnet": true,
    "storageSizeRangeGB": {
      "max": 64000,
      "min": 20
    },
    "storageTypeOptions": [
      "cloud_auto",
      "cloud_essd",
      "cloud_essd2",
      "cloud_essd3",
      "cloud_ssd"
    ],
    "supportedVersions": [
      "5.7",
      "8.0",
      "8.4"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": true,
    "supportsEncryption": true,
    "supportsHighAvailability": true,
    "supportsPublicAccess": true,
    "supportsStorageSizeConfiguration": true,
    "supportsStorageTypeSelection": true,
    "supportsTag": true
  }
}
```

### 6. Beetle POST Recommend RDBMS [✅ SUCCESS]
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "alibaba",
    "region": "ap-northeast-2"
  },
  "sourceRDBMSInstances": [
    {
      "dbEngine": {
        "engine": "mysql",
        "engineVersion": "8.0",
        "port": 3306,
        "role": "primary"
      },
      "dbNode": {
        "cpu": {
          "cores": 2,
          "cpus": 1,
          "maxSpeed": 2.4,
          "threads": 2
        },
        "dataDisks": [
          {
            "label": "/data",
            "totalSize": 50,
            "type": "SSD"
          }
        ],
        "hostname": "db-server-01",
        "machineId": "node-550e8400-e29b-41d4-a716-446655440000",
        "memory": {
          "totalSize": 4
        },
        "rootDisk": {
          "label": "/",
          "totalSize": 50,
          "type": "SSD"
        }
      },
      "description": "Production database instance discovered from on-premise node",
      "displayName": "Source MySQL 01",
      "innerDatabases": [
        {
          "characterSet": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ]
    }
  ],
  "targetPreferences": {
    "adminUserName": "dbadmin",
    "backupRetentionDays": 7,
    "highAvailability": false,
    "publicAccess": true
  }
}
```
```json
// Response Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for alibaba (ap-northeast-2)",
  "status": "recommended",
  "targetCloud": {
    "csp": "alibaba",
    "region": "ap-northeast-2"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "dbadmin",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "charset": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-alibaba",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "cloud_auto",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-alibaba"
    }
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 1m24.96s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "dbadmin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "alibaba-ap-northeast-2",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "mysql.n2.medium.1",
  "name": "rdbms-alibaba",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-alibaba"
  ],
  "storageSize": 100,
  "storageType": "cloud_auto",
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-alibaba"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "dbadmin",
    "adminUserPassword": "******",
    "connectionName": "alibaba-ap-northeast-2",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0",
    "dbInstanceSpec": "mysql.n2.medium.1",
    "name": "rdbms-alibaba",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-alibaba"
    ],
    "storageSize": 100,
    "storageType": "cloud_auto",
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-alibaba"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 2m48.514s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for alibaba (ap-northeast-2)",
  "status": "recommended",
  "targetCloud": {
    "csp": "alibaba",
    "region": "ap-northeast-2"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "dbadmin",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "charset": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-alibaba",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "cloud_auto",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-alibaba"
    }
  ]
}
```
```json
// Response Body
{
  "message": "Managed RDBMS instances created successfully",
  "success": true
}
```

### 9. Beetle GET RDBMS Info [✅ SUCCESS]
- **Duration:** 12ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "07:00Z-08:00Z",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T10:01:41Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T10:01:41Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    }
  ],
  "connectionConfig": {
    "configName": "alibaba-ap-northeast-2",
    "credentialHolder": "admin",
    "credentialName": "alibaba",
    "driverName": "alibaba-driver-v1.0.so",
    "providerName": "alibaba",
    "regionDetail": {
      "description": "South Korea (Seoul)",
      "location": {
        "display": "South Korea (Seoul)",
        "latitude": 37.36,
        "longitude": 126.78
      },
      "regionId": "ap-northeast-2",
      "regionName": "ap-northeast-2",
      "zones": [
        "ap-northeast-2a",
        "ap-northeast-2b"
      ]
    },
    "regionRepresentative": true,
    "regionZoneInfo": {
      "assignedRegion": "ap-northeast-2",
      "assignedZone": "ap-northeast-2a"
    },
    "regionZoneInfoName": "alibaba-ap-northeast-2",
    "verified": true
  },
  "connectionName": "alibaba-ap-northeast-2",
  "cspResourceId": "rm-mj7uar0nh1r8n558k",
  "cspResourceName": "tbc066ftkd9ijporkvrq",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "mysql.n2.medium.1",
  "dbInstanceType": "Basic",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "43.108.66.129:3306",
  "highAvailability": false,
  "id": "test-rdbms-alibaba",
  "keyValueList": [
    {
      "key": "Engine",
      "value": "MySQL"
    },
    {
      "key": "DBInstanceDiskUsed",
      "value": "0"
    },
    {
      "key": "AdvancedFeatures",
      "value": "LinkedServer,DistributeTransaction"
    },
    {
      "key": "VpcCloudInstanceId",
      "value": "rm-mj7uar0nh1r8n558k-20261125175911"
    },
    {
      "key": "DBInstanceNetType",
      "value": "Intranet"
    },
    {
      "key": "DBInstanceClass",
      "value": "mysql.n2.medium.1"
    },
    {
      "key": "MaxIOMBPS",
      "value": "128"
    },
    {
      "key": "DBMaxQuantity",
      "value": "99999"
    },
    {
      "key": "DBInstanceCPU",
      "value": "2"
    },
    {
      "key": "MaxConnections",
      "value": "4000"
    },
    {
      "key": "InstanceNetworkType",
      "value": "VPC"
    },
    {
      "key": "MultipleTempUpgrade",
      "value": "false"
    },
    {
      "key": "DBInstanceType",
      "value": "Primary"
    },
    {
      "key": "DBInstanceId",
      "value": "rm-mj7uar0nh1r8n558k"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7t8uchc521yb50jec8v"
    },
    {
      "key": "DBInstanceMemory",
      "value": "4096"
    },
    {
      "key": "DBInstanceStorageType",
      "value": "general_essd"
    },
    {
      "key": "LatestKernelVersion",
      "value": "rds_20260620"
    },
    {
      "key": "SecurityIPList",
      "value": "0.0.0.0/0"
    },
    {
      "key": "SupportUpgradeAccountType",
      "value": "No"
    },
    {
      "key": "MaxIOPS",
      "value": "7550"
    },
    {
      "key": "TipsLevel",
      "value": "1"
    },
    {
      "key": "MaintainTime",
      "value": "18:00Z-22:00Z"
    },
    {
      "key": "EngineVersion",
      "value": "8.0"
    },
    {
      "key": "InstructionSetArch",
      "value": "x86"
    },
    {
      "key": "PayType",
      "value": "Postpaid"
    },
    {
      "key": "BurstingEnabled",
      "value": "false"
    },
    {
      "key": "Tips",
      "value": "ok"
    },
    {
      "key": "DBInstanceStorage",
      "value": "100"
    },
    {
      "key": "KindCode",
      "value": "18"
    },
    {
      "key": "ColdDataEnabled",
      "value": "false"
    },
    {
      "key": "SupportCreateSuperAccount",
      "value": "No"
    },
    {
      "key": "CurrentKernelVersion",
      "value": "rds_20260620"
    },
    {
      "key": "ZoneId",
      "value": "ap-northeast-2a"
    },
    {
      "key": "IPType",
      "value": "IPv4"
    },
    {
      "key": "ConnectionMode",
      "value": "Standard"
    },
    {
      "key": "LockMode",
      "value": "Unlock"
    },
    {
      "key": "CanTempUpgrade",
      "value": "false"
    },
    {
      "key": "Category",
      "value": "Basic"
    },
    {
      "key": "InsId",
      "value": "1"
    },
    {
      "key": "DBInstanceDescription",
      "value": "tbc066ftkd9ijporkvrq"
    },
    {
      "key": "AccountType",
      "value": "Mix"
    },
    {
      "key": "RegionId",
      "value": "ap-northeast-2"
    },
    {
      "key": "ConsoleVersion",
      "value": "2"
    },
    {
      "key": "ResourceGroupId",
      "value": "rg-acfnvekhilw5kmy"
    },
    {
      "key": "DispenseMode",
      "value": "ClassicDispenseMode"
    },
    {
      "key": "TempUpgradeRecoveryMemory",
      "value": "0"
    },
    {
      "key": "ProxyType",
      "value": "2"
    },
    {
      "key": "AccountMaxQuantity",
      "value": "99999"
    },
    {
      "key": "Port",
      "value": "3306"
    },
    {
      "key": "VSwitchId",
      "value": "vsw-mj7yfvd5kk99kumix6i1u"
    },
    {
      "key": "CreationTime",
      "value": "2026-11-25T09:59:16Z"
    },
    {
      "key": "SecurityIPMode",
      "value": "normal"
    },
    {
      "key": "DBInstanceClassType",
      "value": "x"
    },
    {
      "key": "DBInstanceStatus",
      "value": "Running"
    },
    {
      "key": "ConnectionString",
      "value": "rm-mj7uar0nh1r8n558k.mysql.ap-northeast-2.rds.aliyuncs.com"
    },
    {
      "key": "TempUpgradeRecoveryCpu",
      "value": "0"
    },
    {
      "key": "AvailabilityValue",
      "value": "100.0%"
    },
    {
      "key": "AutoUpgradeMinorVersion",
      "value": "Auto"
    },
    {
      "key": "DeletionProtection",
      "value": "false"
    },
    {
      "key": "Extra",
      "value": "{ReplicaGroupStatus:,ActiveReplicaDBInstanceID:,ReplicaGroupID:,RecoveryModel:,DBInstanceIds:{DBInstanceId:[]}}"
    },
    {
      "key": "ServerlessConfig",
      "value": "{ScaleMin:,ScaleMax:,AutoPause:false,SwitchForce:false}"
    },
    {
      "key": "BabelfishConfig",
      "value": "{BabelfishEnabled:,MigrationMode:}"
    },
    {
      "key": "DBClusterNodes",
      "value": "{DBClusterNode:[]}"
    },
    {
      "key": "ReadOnlyDBInstanceIds",
      "value": "{ReadOnlyDBInstanceId:[]}"
    },
    {
      "key": "SlaveZones",
      "value": "{SlaveZone:[]}"
    }
  ],
  "name": "test-rdbms-alibaba",
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-alibaba"
  ],
  "status": "Available",
  "storageSize": 100,
  "storageType": "general_essd",
  "subnetIds": [
    "subnet-1"
  ],
  "uid": "tbc066ftkd9ijporkvrq",
  "vNetId": "test-rdbms-vnet-alibaba"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 66ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "backupTime": "07:00Z-08:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:01:41Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:01:41Z",
          "reason": "Available",
          "status": "True",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "alibaba-ap-northeast-2",
        "credentialHolder": "admin",
        "credentialName": "alibaba",
        "driverName": "alibaba-driver-v1.0.so",
        "providerName": "alibaba",
        "regionDetail": {
          "description": "South Korea (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionZoneInfoName": "alibaba-ap-northeast-2",
        "verified": true
      },
      "connectionName": "alibaba-ap-northeast-2",
      "cspResourceId": "rm-mj7uar0nh1r8n558k",
      "cspResourceName": "tbc066ftkd9ijporkvrq",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "dbInstanceType": "Basic",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "43.108.66.129:3306",
      "highAvailability": false,
      "id": "test-rdbms-alibaba",
      "keyValueList": [
        {
          "key": "Engine",
          "value": "MySQL"
        },
        {
          "key": "DBInstanceDiskUsed",
          "value": "0"
        },
        {
          "key": "AdvancedFeatures",
          "value": "LinkedServer,DistributeTransaction"
        },
        {
          "key": "VpcCloudInstanceId",
          "value": "rm-mj7uar0nh1r8n558k-20261125175911"
        },
        {
          "key": "DBInstanceNetType",
          "value": "Intranet"
        },
        {
          "key": "DBInstanceClass",
          "value": "mysql.n2.medium.1"
        },
        {
          "key": "MaxIOMBPS",
          "value": "128"
        },
        {
          "key": "DBMaxQuantity",
          "value": "99999"
        },
        {
          "key": "DBInstanceCPU",
          "value": "2"
        },
        {
          "key": "MaxConnections",
          "value": "4000"
        },
        {
          "key": "InstanceNetworkType",
          "value": "VPC"
        },
        {
          "key": "MultipleTempUpgrade",
          "value": "false"
        },
        {
          "key": "DBInstanceType",
          "value": "Primary"
        },
        {
          "key": "DBInstanceId",
          "value": "rm-mj7uar0nh1r8n558k"
        },
        {
          "key": "VpcId",
          "value": "vpc-mj7t8uchc521yb50jec8v"
        },
        {
          "key": "DBInstanceMemory",
          "value": "4096"
        },
        {
          "key": "DBInstanceStorageType",
          "value": "general_essd"
        },
        {
          "key": "LatestKernelVersion",
          "value": "rds_20260620"
        },
        {
          "key": "SecurityIPList",
          "value": "0.0.0.0/0"
        },
        {
          "key": "SupportUpgradeAccountType",
          "value": "No"
        },
        {
          "key": "MaxIOPS",
          "value": "7550"
        },
        {
          "key": "TipsLevel",
          "value": "1"
        },
        {
          "key": "MaintainTime",
          "value": "18:00Z-22:00Z"
        },
        {
          "key": "EngineVersion",
          "value": "8.0"
        },
        {
          "key": "InstructionSetArch",
          "value": "x86"
        },
        {
          "key": "PayType",
          "value": "Postpaid"
        },
        {
          "key": "BurstingEnabled",
          "value": "false"
        },
        {
          "key": "Tips",
          "value": "ok"
        },
        {
          "key": "DBInstanceStorage",
          "value": "100"
        },
        {
          "key": "KindCode",
          "value": "18"
        },
        {
          "key": "ColdDataEnabled",
          "value": "false"
        },
        {
          "key": "SupportCreateSuperAccount",
          "value": "No"
        },
        {
          "key": "CurrentKernelVersion",
          "value": "rds_20260620"
        },
        {
          "key": "ZoneId",
          "value": "ap-northeast-2a"
        },
        {
          "key": "IPType",
          "value": "IPv4"
        },
        {
          "key": "ConnectionMode",
          "value": "Standard"
        },
        {
          "key": "LockMode",
          "value": "Unlock"
        },
        {
          "key": "CanTempUpgrade",
          "value": "false"
        },
        {
          "key": "Category",
          "value": "Basic"
        },
        {
          "key": "InsId",
          "value": "1"
        },
        {
          "key": "DBInstanceDescription",
          "value": "tbc066ftkd9ijporkvrq"
        },
        {
          "key": "AccountType",
          "value": "Mix"
        },
        {
          "key": "RegionId",
          "value": "ap-northeast-2"
        },
        {
          "key": "ConsoleVersion",
          "value": "2"
        },
        {
          "key": "ResourceGroupId",
          "value": "rg-acfnvekhilw5kmy"
        },
        {
          "key": "DispenseMode",
          "value": "ClassicDispenseMode"
        },
        {
          "key": "TempUpgradeRecoveryMemory",
          "value": "0"
        },
        {
          "key": "ProxyType",
          "value": "2"
        },
        {
          "key": "AccountMaxQuantity",
          "value": "99999"
        },
        {
          "key": "Port",
          "value": "3306"
        },
        {
          "key": "VSwitchId",
          "value": "vsw-mj7yfvd5kk99kumix6i1u"
        },
        {
          "key": "CreationTime",
          "value": "2026-11-25T09:59:16Z"
        },
        {
          "key": "SecurityIPMode",
          "value": "normal"
        },
        {
          "key": "DBInstanceClassType",
          "value": "x"
        },
        {
          "key": "DBInstanceStatus",
          "value": "Running"
        },
        {
          "key": "ConnectionString",
          "value": "rm-mj7uar0nh1r8n558k.mysql.ap-northeast-2.rds.aliyuncs.com"
        },
        {
          "key": "TempUpgradeRecoveryCpu",
          "value": "0"
        },
        {
          "key": "AvailabilityValue",
          "value": "100.0%"
        },
        {
          "key": "AutoUpgradeMinorVersion",
          "value": "Auto"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Extra",
          "value": "{ReplicaGroupStatus:,ActiveReplicaDBInstanceID:,ReplicaGroupID:,RecoveryModel:,DBInstanceIds:{DBInstanceId:[]}}"
        },
        {
          "key": "ServerlessConfig",
          "value": "{ScaleMin:,ScaleMax:,AutoPause:false,SwitchForce:false}"
        },
        {
          "key": "BabelfishConfig",
          "value": "{BabelfishEnabled:,MigrationMode:}"
        },
        {
          "key": "DBClusterNodes",
          "value": "{DBClusterNode:[]}"
        },
        {
          "key": "ReadOnlyDBInstanceIds",
          "value": "{ReadOnlyDBInstanceId:[]}"
        },
        {
          "key": "SlaveZones",
          "value": "{SlaveZone:[]}"
        }
      ],
      "name": "test-rdbms-alibaba",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "general_essd",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbc066ftkd9ijporkvrq",
      "vNetId": "test-rdbms-vnet-alibaba"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "16:23-16:53",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:55:24Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:55:24Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "aws-ap-northeast-2",
        "credentialHolder": "admin",
        "credentialName": "aws",
        "driverName": "aws-driver-v1.0.so",
        "providerName": "aws",
        "regionDetail": {
          "description": "Asia Pacific (Seoul)",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.36,
            "longitude": 126.78
          },
          "regionId": "ap-northeast-2",
          "regionName": "ap-northeast-2",
          "zones": [
            "ap-northeast-2a",
            "ap-northeast-2b",
            "ap-northeast-2c",
            "ap-northeast-2d"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "ap-northeast-2",
          "assignedZone": "ap-northeast-2a"
        },
        "regionZoneInfoName": "aws-ap-northeast-2",
        "verified": true
      },
      "connectionName": "aws-ap-northeast-2",
      "cspResourceId": "tbvtoteol9rtlis035nc",
      "cspResourceName": "tbvtoteol9rtlis035nc",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.46-rds.20260908",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "tbvtoteol9rtlis035nc.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-aws",
      "iops": "3000",
      "keyValueList": [
        {
          "key": "ActivityStreamStatus",
          "value": "stopped"
        },
        {
          "key": "AllocatedStorage",
          "value": "100"
        },
        {
          "key": "AutoMinorVersionUpgrade",
          "value": "true"
        },
        {
          "key": "AvailabilityZone",
          "value": "ap-northeast-2c"
        },
        {
          "key": "BackupRetentionPeriod",
          "value": "7"
        },
        {
          "key": "CACertificateIdentifier",
          "value": "rds-ca-rsa2048-g1"
        },
        {
          "key": "CopyTagsToSnapshot",
          "value": "false"
        },
        {
          "key": "CustomerOwnedIpEnabled",
          "value": "false"
        },
        {
          "key": "DBInstanceArn",
          "value": "arn:aws:rds:ap-northeast-2:635484366616:db:tbvtoteol9rtlis035nc"
        },
        {
          "key": "DBInstanceClass",
          "value": "db.t3.medium"
        },
        {
          "key": "DBInstanceIdentifier",
          "value": "tbvtoteol9rtlis035nc"
        },
        {
          "key": "DBInstanceStatus",
          "value": "backing-up"
        },
        {
          "key": "DBParameterGroups",
          "value": "{DBParameterGroupName:default.mysql8.0,ParameterApplyStatus:in-sync}"
        },
        {
          "key": "DBSubnetGroup",
          "value": "{DBSubnetGroupArn:null,DBSubnetGroupDescription:CB-Spider RDBMS subnet group for cb-spider-tbvtoteol9rtlis035nc,DBSubnetGroupName:cb-spider-tbvtoteol9rtlis035nc,SubnetGroupStatus:Complete,Subnets:[{SubnetAvailabilityZone:{Name:ap-northeast-2a},SubnetIdentifier:subnet-0cc10b17b6ec493e1,SubnetOutpost:{Arn:null},SubnetStatus:Active},{SubnetAvailabilityZone:{Name:ap-northeast-2c},SubnetIdentifier:subnet-0d95aacbfb1b0dbeb,SubnetOutpost:{Arn:null},SubnetStatus:Active}],VpcId:vpc-014472542cc87bfa5}"
        },
        {
          "key": "DbInstancePort",
          "value": "0"
        },
        {
          "key": "DbiResourceId",
          "value": "db-YIQ2PUCPDSN5UZ6AP34UZKPPJI"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Endpoint",
          "value": "{Address:tbvtoteol9rtlis035nc.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com,HostedZoneId:ZLA2NUCOLGUUR,Port:3306}"
        },
        {
          "key": "Engine",
          "value": "mysql"
        },
        {
          "key": "EngineVersion",
          "value": "8.0.46-rds.20260908"
        },
        {
          "key": "IAMDatabaseAuthenticationEnabled",
          "value": "false"
        },
        {
          "key": "InstanceCreateTime",
          "value": "2026-11-25T09:58:58.611Z"
        },
        {
          "key": "Iops",
          "value": "3000"
        },
        {
          "key": "LicenseModel",
          "value": "general-public-license"
        },
        {
          "key": "MasterUsername",
          "value": "admin"
        },
        {
          "key": "MonitoringInterval",
          "value": "0"
        },
        {
          "key": "MultiAZ",
          "value": "false"
        },
        {
          "key": "OptionGroupMemberships",
          "value": "{OptionGroupName:default:mysql-8-0,Status:in-sync}"
        },
        {
          "key": "PendingModifiedValues",
          "value": "{AllocatedStorage:null,BackupRetentionPeriod:null,CACertificateIdentifier:null,DBInstanceClass:null,DBInstanceIdentifier:null,DBSubnetGroupName:null,EngineVersion:null,IAMDatabaseAuthenticationEnabled:null,Iops:null,LicenseModel:null,MasterUserPassword:null,MultiAZ:null,PendingCloudwatchLogsExports:null,Port:null,ProcessorFeatures:null,StorageType:null}"
        },
        {
          "key": "PerformanceInsightsEnabled",
          "value": "false"
        },
        {
          "key": "PreferredBackupWindow",
          "value": "16:23-16:53"
        },
        {
          "key": "PreferredMaintenanceWindow",
          "value": "mon:14:06-mon:14:36"
        },
        {
          "key": "PubliclyAccessible",
          "value": "true"
        },
        {
          "key": "StorageEncrypted",
          "value": "false"
        },
        {
          "key": "StorageType",
          "value": "gp3"
        },
        {
          "key": "TagList",
          "value": "{Key:Name,Value:tbvtoteol9rtlis035nc}"
        },
        {
          "key": "VpcSecurityGroups",
          "value": "{Status:active,VpcSecurityGroupId:sg-00fbbe9cd2de98ea2}"
        }
      ],
      "name": "test-rdbms-aws",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-aws"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "gp3",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "tagList": [
        {
          "key": "Name",
          "value": "tbvtoteol9rtlis035nc"
        }
      ],
      "uid": "tbvtoteol9rtlis035nc",
      "vNetId": "test-rdbms-vnet-aws"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "AUTO",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:59:42Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:59:42Z",
          "reason": "Available",
          "status": "True",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "azure-koreacentral",
        "credentialHolder": "admin",
        "credentialName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "providerName": "azure",
        "regionDetail": {
          "description": "Korea Central",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          },
          "regionId": "koreacentral",
          "regionName": "koreacentral",
          "zones": [
            "1",
            "2",
            "3"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "koreacentral",
          "assignedZone": ""
        },
        "regionZoneInfoName": "azure-koreacentral",
        "verified": true
      },
      "connectionName": "azure-koreacentral",
      "cspResourceId": "tbqt11ul0jr723gfs0im",
      "cspResourceName": "tbqt11ul0jr723gfs0im",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "dbInstanceType": "Burstable",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "tbqt11ul0jr723gfs0im.mysql.database.azure.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-azure",
      "keyValueList": [
        {
          "key": "Location",
          "value": "Korea Central"
        },
        {
          "key": "Properties",
          "value": "{administratorLogin:dbadmin,availabilityZone:1,backup:{backupRetentionDays:7,earliestRestoreDate:2026-11-25T10:05:13.3889866Z,geoRedundantBackup:Disabled},fullyQualifiedDomainName:tbqt11ul0jr723gfs0im.mysql.database.azure.com,highAvailability:{mode:Disabled,standbyAvailabilityZone:,state:NotEnabled},maintenanceWindow:{customWindow:Disabled,dayOfWeek:0,startHour:0,startMinute:0},network:{publicNetworkAccess:Enabled},replicaCapacity:10,replicationRole:None,state:Ready,storage:{autoGrow:Enabled,iops:600,storageSizeGB:100,storageSku:Premium_LRS},version:8.0.21}"
        },
        {
          "key": "SKU",
          "value": "{name:Standard_B2s,tier:Burstable}"
        },
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.DBforMySQL/flexibleServers/tbqt11ul0jr723gfs0im"
        },
        {
          "key": "Name",
          "value": "tbqt11ul0jr723gfs0im"
        },
        {
          "key": "SystemData",
          "value": "{createdAt:2026-11-25T09:55:13.3889866Z}"
        },
        {
          "key": "Type",
          "value": "Microsoft.DBforMySQL/flexibleServers"
        }
      ],
      "name": "test-rdbms-azure",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "Premium_LRS",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbqt11ul0jr723gfs0im",
      "vNetId": "test-rdbms-vnet-azure"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "16:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:58:37Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:58:37Z",
          "reason": "Available",
          "status": "True",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "gcp-us-central1",
        "credentialHolder": "admin",
        "credentialName": "gcp",
        "driverName": "gcp-driver-v1.0.so",
        "providerName": "gcp",
        "regionDetail": {
          "description": "Council Bluffs Iowa  USA",
          "location": {
            "display": "Council Bluffs Iowa USA",
            "latitude": 41.2522,
            "longitude": -95.8575
          },
          "regionId": "us-central1",
          "regionName": "us-central1",
          "zones": [
            "us-central1-a",
            "us-central1-b",
            "us-central1-c",
            "us-central1-f"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "us-central1",
          "assignedZone": "us-central1-a"
        },
        "regionZoneInfoName": "gcp-us-central1",
        "verified": true
      },
      "connectionName": "gcp-us-central1",
      "cspResourceId": "tbc0aj584ko18830vb1j",
      "cspResourceName": "tbc0aj584ko18830vb1j",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "dbInstanceType": "ZONAL",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "encryption": true,
      "endpoint": "136.80.5.81:3306",
      "highAvailability": false,
      "id": "test-rdbms-gcp",
      "keyValueList": [
        {
          "key": "BackendType",
          "value": "SECOND_GEN"
        },
        {
          "key": "ConnectionName",
          "value": "ykkim-etri:us-central1:tbc0aj584ko18830vb1j"
        },
        {
          "key": "CreateTime",
          "value": "2026-11-25T09:55:37.790Z"
        },
        {
          "key": "CurrentDiskSize",
          "value": "0"
        },
        {
          "key": "DatabaseInstalledVersion",
          "value": "MYSQL_8_0_46"
        },
        {
          "key": "DatabaseVersion",
          "value": "MYSQL_8_0"
        },
        {
          "key": "Etag",
          "value": "74c240921252e7aa3276e6b5dc8adb3ebefa361edcf0d4a331a9eee4d0012ea1"
        },
        {
          "key": "GceZone",
          "value": "us-central1-c"
        },
        {
          "key": "GeminiConfig",
          "value": "{flagRecommenderEnabled:true}"
        },
        {
          "key": "IncludeReplicasForMajorVersionUpgrade",
          "value": "false"
        },
        {
          "key": "InstanceType",
          "value": "CLOUD_SQL_INSTANCE"
        },
        {
          "key": "IpAddresses",
          "value": "{ipAddress:136.80.5.81,type:PRIMARY}"
        },
        {
          "key": "Kind",
          "value": "sql#instance"
        },
        {
          "key": "MaintenanceVersion",
          "value": "MYSQL_8_0_46.R20261013.02_15"
        },
        {
          "key": "MaxDiskSize",
          "value": "0"
        },
        {
          "key": "Name",
          "value": "tbc0aj584ko18830vb1j"
        },
        {
          "key": "NodeCount",
          "value": "0"
        },
        {
          "key": "Project",
          "value": "ykkim-etri"
        },
        {
          "key": "Region",
          "value": "us-central1"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "SelfLink",
          "value": "https://sqladmin.googleapis.com/sql/v1beta4/projects/GCP_PROJECT_ID/instances/tbc0aj584ko18830vb1j"
        },
        {
          "key": "ServerCaCert",
          "value": "{cert:-----BEGIN CERTIFICATE-----\\nMIIDcTCCAlmgAwIBAgIBADANBgkqhkiG9w0BAQsFADBwMS0wKwYDVQQuEyRhODg2\\nMWU3Ni0xYjRmLTRlZmItYmM4My0zOTZiYmZlZGEwNWIxHDAaBgNVBAMTE0Nsb3Vk\\nIFNRTCBTZXJ2ZXIgQ0ExFDASBgNVBAoTC0dvb2dsZSwgSW5jMQswCQYDVQQGEwJV\\nUzAeFw0yNjExMjUwOTU1MjZaFw0zNjExMjIwOTU2MjZaMHAxLTArBgNVBC4TJGE4\\nODYxZTc2LTFiNGYtNGVmYi1iYzgzLTM5NmJiZmVkYTA1YjEcMBoGA1UEAxMTQ2xv\\ndWQgU1FMIFNlcnZlciBDQTEUMBIGA1UEChMLR29vZ2xlLCBJbmMxCzAJBgNVBAYT\\nAlVTMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAilpdp2ZUIpAlxNob\\nmpLtJjEGEMvs0V77gI+27HmvrWkNOT3ppFhxLEnvsqdzKbVzSvD3ptPYm7zqiAoX\\nthlJ0LFmSUVw5ViP6EiTtywMiG0QNbt7UUSvU849j96MYN3h2wTJZgKm0co26Vxw\\ndZ2W4TK8WKOD5BF9vCGSkJVkpF1xPzOFl231DYm0KLrBRM4rMThSwB3ICifQiUso\\nlnIgiLZtdcEd1dXm1E8SMaAva6S7J3nd1SKpypnaCr4B3STY5wUEu8OXhDvYfzId\\nlRfEH5r0xptyIZjAR6ZrfcAUhTtOHSvSakaf9aSwzZnjEaYiO6dSHx0UhwH+gIxF\\n4jWwvQIDAQABoxYwFDASBgNVHRMBAf8ECDAGAQH/AgEAMA0GCSqGSIb3DQEBCwUA\\nA4IBAQByQWdUNDPEyiGz5HCpboeKwSmCrYxuM5JzORuaj6yJY3G7TPKnYb4ltM+6\\nU5d99PArl09QgNt6RNWyVdqp+/QycPwadArcS5G+1z/Md6yfp6/KlMD1cuxtxxsS\\nl4vGbaTAwUT5ldwQ3HfyZFnQfWyMTr7f4fwsrJkyjTPetPIqYHQgg63qbyucFiut\\nT+XkL+oQU41uSH9SHZ9iqJwTK+UFYi3d0cu2ex9ouf8YDihNn60sLKZiwHkYM5aj\\n0919+nHI2xXAwZi9+Qjo8VL5qxJK4U2tUUAu06NRnWTc+khUOI9qy3D5YjglqmV1\\nMb5SVhJd+pwMe1BK/zMlEOpk2Z3D\\n-----END CERTIFICATE-----,certSerialNumber:0,commonName:C=US,O=Google\\\\, Inc,CN=Cloud SQL Server CA,dnQualifier=a8861e76-1b4f-4efb-bc83-396bbfeda05b,createTime:2026-11-25T09:55:26.377Z,expirationTime:2036-11-22T09:56:26.377Z,instance:tbc0aj584ko18830vb1j,kind:sql#sslCert,sha1Fingerprint:6a34e29353b85234b2b6769082c46082deded54a}"
        },
        {
          "key": "ServiceAccountEmailAddress",
          "value": "MASKED_EMAIL"
        },
        {
          "key": "Settings",
          "value": "{activationPolicy:ALWAYS,availabilityType:ZONAL,backupConfiguration:{backupRetentionSettings:{retainedBackups:7,retentionUnit:COUNT},backupTier:STANDARD,enabled:true,kind:sql#backupConfiguration,startTime:16:00,transactionLogRetentionDays:7,transactionalLogStorageState:TRANSACTIONAL_LOG_STORAGE_STATE_UNSPECIFIED},connectorEnforcement:NOT_REQUIRED,dataDiskSizeGb:100,dataDiskType:PD_SSD,edition:ENTERPRISE,enableDataplexIntegration:true,ipConfiguration:{authorizedNetworks:[{kind:sql#aclEntry,name:AllowAll,value:0.0.0.0/0}],ipv4Enabled:true,serverCaMode:GOOGLE_MANAGED_INTERNAL_CA,serverCertificateRotationMode:SERVER_CERTIFICATE_ROTATION_MODE_UNSPECIFIED,sslMode:ALLOW_UNENCRYPTED_AND_ENCRYPTED},kind:sql#settings,locationPreference:{kind:sql#locationPreference,zone:us-central1-c},pricingPlan:PER_USE,replicationLagMaxSeconds:31536000,replicationType:SYNCHRONOUS,settingsVersion:1,storageAutoResize:true,tier:db-n1-standard-2}"
        },
        {
          "key": "SqlNetworkArchitecture",
          "value": "NEW_NETWORK_ARCHITECTURE"
        },
        {
          "key": "State",
          "value": "RUNNABLE"
        },
        {
          "key": "SwitchTransactionLogsToCloudStorageEnabled",
          "value": "false"
        },
        {
          "key": "UpgradableDatabaseVersions",
          "value": "{displayName:MySQL 8.0.18,majorVersion:MYSQL_8_0,name:MYSQL_8_0_18}; {displayName:MySQL 8.0.26,majorVersion:MYSQL_8_0,name:MYSQL_8_0_26}; {displayName:MySQL 8.0.27,majorVersion:MYSQL_8_0,name:MYSQL_8_0_27}; {displayName:MySQL 8.0.28,majorVersion:MYSQL_8_0,name:MYSQL_8_0_28}; {displayName:MySQL 8.0.29,majorVersion:MYSQL_8_0,name:MYSQL_8_0_29}; {displayName:MySQL 8.0.30,majorVersion:MYSQL_8_0,name:MYSQL_8_0_30}; {displayName:MySQL 8.0.31,majorVersion:MYSQL_8_0,name:MYSQL_8_0_31}; {displayName:MySQL 8.0.32,majorVersion:MYSQL_8_0,name:MYSQL_8_0_32}; {displayName:MySQL 8.0.33,majorVersion:MYSQL_8_0,name:MYSQL_8_0_33}; {displayName:MySQL 8.0.34,majorVersion:MYSQL_8_0,name:MYSQL_8_0_34}; {displayName:MySQL 8.0.35,majorVersion:MYSQL_8_0,name:MYSQL_8_0_35}; {displayName:MySQL 8.0.36,majorVersion:MYSQL_8_0,name:MYSQL_8_0_36}; {displayName:MySQL 8.0.37,majorVersion:MYSQL_8_0,name:MYSQL_8_0_37}; {displayName:MySQL 8.0.39,majorVersion:MYSQL_8_0,name:MYSQL_8_0_39}; {displayName:MySQL 8.0.40,majorVersion:MYSQL_8_0,name:MYSQL_8_0_40}; {displayName:MySQL 8.0.41,majorVersion:MYSQL_8_0,name:MYSQL_8_0_41}; {displayName:MySQL 8.0.42,majorVersion:MYSQL_8_0,name:MYSQL_8_0_42}; {displayName:MySQL 8.0.43,majorVersion:MYSQL_8_0,name:MYSQL_8_0_43}; {displayName:MySQL 8.0.44,majorVersion:MYSQL_8_0,name:MYSQL_8_0_44}; {displayName:MySQL 8.0.45,majorVersion:MYSQL_8_0,name:MYSQL_8_0_45}; {displayName:MySQL 8.4,majorVersion:MYSQL_8_4,name:MYSQL_8_4}"
        }
      ],
      "name": "test-rdbms-gcp",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "PD_SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbc0aj584ko18830vb1j",
      "vNetId": "test-rdbms-vnet-gcp"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:55:07Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:55:07Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "ibm-us-south",
        "credentialHolder": "admin",
        "credentialName": "ibm",
        "driverName": "ibm-driver-v1.0.so",
        "providerName": "ibm",
        "regionDetail": {
          "description": "us-south",
          "location": {
            "display": "Dallas USA",
            "latitude": 32.81248,
            "longitude": -96.77619
          },
          "regionId": "us-south",
          "regionName": "us-south",
          "zones": [
            "us-south-1",
            "us-south-2",
            "us-south-3"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "us-south",
          "assignedZone": "us-south-1"
        },
        "regionZoneInfoName": "ibm-us-south",
        "verified": true
      },
      "connectionName": "ibm-us-south",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "multitenant",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-ibm",
      "name": "test-rdbms-ibm",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "status": "Creating",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbun4h3ej4rgndeesotl",
      "vNetId": "test-rdbms-vnet-ibm"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:55:13Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:55:13Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "ncp-kr",
        "credentialHolder": "admin",
        "credentialName": "ncp",
        "driverName": "ncp-driver-v1.0.so",
        "providerName": "ncp",
        "regionDetail": {
          "description": "Korea 1",
          "location": {
            "display": "Seoul(Gasan) / Pyeongchon (South Korea)",
            "latitude": 37.4754,
            "longitude": 126.8831
          },
          "regionId": "KR",
          "regionName": "kr",
          "zones": [
            "KR-1",
            "KR-2"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "KR",
          "assignedZone": "KR-1"
        },
        "regionZoneInfoName": "ncp-kr",
        "verified": true
      },
      "connectionName": "ncp-kr",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.45",
      "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-ncp",
      "name": "test-rdbms-ncp",
      "ncpDBACGToAllowAllInbound": true,
      "publicAccess": false,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-ncp"
      ],
      "status": "Creating",
      "storageSize": 0,
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbpth30tovigfplksoib",
      "vNetId": "test-rdbms-vnet-ncp"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:55:13Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:55:13Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "nhn-kr1",
        "credentialHolder": "admin",
        "credentialName": "nhn",
        "driverName": "nhn-driver-v1.0.so",
        "providerName": "nhn",
        "regionDetail": {
          "description": "Pangyo (South Korea)",
          "location": {
            "display": "Pangyo (South Korea)",
            "latitude": 37.390889,
            "longitude": 127.096792
          },
          "regionId": "KR1",
          "regionName": "kr1",
          "zones": [
            "kr-pub-a",
            "kr-pub-b"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "KR1",
          "assignedZone": "kr-pub-a"
        },
        "regionZoneInfoName": "nhn-kr1",
        "verified": true
      },
      "connectionName": "nhn-kr1",
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL_V8046",
      "dbInstanceSpec": "m2.c2m4",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-nhn",
      "name": "test-rdbms-nhn",
      "nhnDBSGToAllowAllInbound": true,
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-nhn"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "General SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbsn6ljiu7npo3re8iqr",
      "vNetId": "test-rdbms-vnet-nhn"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "00:00-12:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:59:46Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:59:46Z",
          "reason": "Available",
          "status": "True",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "tencent-ap-seoul",
        "credentialHolder": "admin",
        "credentialName": "tencent",
        "driverName": "tencent-driver-v1.0.so",
        "providerName": "tencent",
        "regionDetail": {
          "description": "Seoul",
          "location": {
            "display": "South Korea (Seoul)",
            "latitude": 37.566536,
            "longitude": 126.977966
          },
          "regionId": "ap-seoul",
          "regionName": "ap-seoul",
          "zones": [
            "ap-seoul-1",
            "ap-seoul-2"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "ap-seoul",
          "assignedZone": "ap-seoul-1"
        },
        "regionZoneInfoName": "tencent-ap-seoul",
        "verified": true
      },
      "connectionName": "tencent-ap-seoul",
      "cspResourceId": "cdb-q1mk2sbg",
      "cspResourceName": "tbjgc7odeet6a3manird",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "kr-cdb-q1mk2sbg.sql.tencentcdb.com:24489",
      "highAvailability": false,
      "id": "test-rdbms-tencent",
      "keyValueList": [
        {
          "key": "WanStatus",
          "value": "1"
        },
        {
          "key": "Zone",
          "value": "ap-seoul-1"
        },
        {
          "key": "InitFlag",
          "value": "1"
        },
        {
          "key": "Memory",
          "value": "4000"
        },
        {
          "key": "Status",
          "value": "1"
        },
        {
          "key": "VpcId",
          "value": "9730283"
        },
        {
          "key": "InstanceId",
          "value": "cdb-q1mk2sbg"
        },
        {
          "key": "Volume",
          "value": "100"
        },
        {
          "key": "AutoRenew",
          "value": "0"
        },
        {
          "key": "ProtectMode",
          "value": "0"
        },
        {
          "key": "SubnetId",
          "value": "820587"
        },
        {
          "key": "InstanceType",
          "value": "1"
        },
        {
          "key": "ProjectId",
          "value": "0"
        },
        {
          "key": "Region",
          "value": "ap-seoul"
        },
        {
          "key": "DeadlineTime",
          "value": "0000-00-00 00:00:00"
        },
        {
          "key": "DeployMode",
          "value": "0"
        },
        {
          "key": "TaskStatus",
          "value": "0"
        },
        {
          "key": "DeviceType",
          "value": "CLOUD_NATIVE_CLUSTER"
        },
        {
          "key": "EngineVersion",
          "value": "8.0"
        },
        {
          "key": "InstanceName",
          "value": "tbjgc7odeet6a3manird"
        },
        {
          "key": "WanDomain",
          "value": "kr-cdb-q1mk2sbg.sql.tencentcdb.com"
        },
        {
          "key": "WanPort",
          "value": "24489"
        },
        {
          "key": "PayType",
          "value": "1"
        },
        {
          "key": "CreateTime",
          "value": "2026-11-25 17:55:12"
        },
        {
          "key": "Vip",
          "value": "10.4.1.13"
        },
        {
          "key": "Vport",
          "value": "3306"
        },
        {
          "key": "CdbError",
          "value": "0"
        },
        {
          "key": "UniqVpcId",
          "value": "vpc-7yqgk5bc"
        },
        {
          "key": "UniqSubnetId",
          "value": "subnet-f4cf8dsh"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "Qps",
          "value": "4400"
        },
        {
          "key": "ZoneId",
          "value": "180001"
        },
        {
          "key": "InstanceNodes",
          "value": "2"
        },
        {
          "key": "EngineType",
          "value": "InnoDB"
        },
        {
          "key": "MaxDelayTime",
          "value": "0"
        },
        {
          "key": "DiskType",
          "value": "CLOUD_HSSD"
        },
        {
          "key": "ExpandCpu",
          "value": "0"
        },
        {
          "key": "DeviceBandwidth",
          "value": "0"
        },
        {
          "key": "DestroyProtect",
          "value": "off"
        }
      ],
      "name": "test-rdbms-tencent",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "CLOUD_HSSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbjgc7odeet6a3manird",
      "vNetId": "test-rdbms-vnet-tencent"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 1.302s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba/database`
```json
// Request Body
{
  "charset": "utf8mb4",
  "collation": "utf8mb4_unicode_ci",
  "databaseName": "sampledb_dyn"
}
```
```json
// Response Body
{
  "message": "Logical database 'sampledb_dyn' created successfully",
  "success": true
}
```

### 12. Beetle GET List Logical Databases [✅ SUCCESS]
- **Duration:** 2.035s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba/database`
```json
// Response Body
{
  "databases": [
    "sampledb",
    "sampledb_dyn"
  ]
}
```

### 13. Beetle GET Secure Transport Info [✅ SUCCESS]
- **Duration:** 1.071s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": false,
    "issuer": "",
    "notAfter": "0001-01-01T00:00:00Z",
    "pem": "",
    "subject": ""
  },
  "enforced": false,
  "engine": "mysql",
  "recommendedSSLMode": "DISABLED",
  "requireSecureTransport": "OFF",
  "rules": "",
  "tlsCipher": "",
  "tlsInUse": false
}
```

### 14. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 7ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 2m45.802s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 7.109s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 8m42.038s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-alibaba?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 1.375s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 7.7s

