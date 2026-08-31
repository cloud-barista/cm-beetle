# Managed RDBMS Test Report: ALIBABA (ap-northeast-2)

- **Test Case:** Alibaba AP-Northeast-2 (Seoul) MySQL Test
- **Date & Time:** 2026-08-31 16:38:03
- **Namespace:** `default`
- **Total Duration:** 20m39.492s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** ALIBABA
- **Target Region:** `ap-northeast-2`
- **Namespace:** `default`
- **Test Date:** 2026-08-31 16:38:03

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
10. **External Data I/O**: Direct TCP/SQL connectivity test
11. **Internal Data I/O**: SQL execution via internal Runner VM (`POST /tumblebug/ns/{nsId}/infra`)
12. **Delete Logical Database**: `DELETE /beetle/migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database/{databaseName}`
13. **Delete RDBMS**: `DELETE /beetle/migration/middleware/ns/{nsId}/rdbms/{rdbmsId}`
14. **Delete Pre-requisite SG & VNet**: `DELETE /tumblebug/ns/{nsId}/resources/securityGroup/{sgId}`, `DELETE /tumblebug/ns/{nsId}/resources/vNet/{vNetId}`

## Execution Steps & API Traces

### 1. Tumblebug POST /specImagePairReview (Pre-flight Spec & Image Review) [✅ SUCCESS]
- **Duration:** 3.417s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
  "specId": "alibaba+ap-northeast-2+ecs.e-c1m2.large"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "instanceType": "ecs.e-c1m2.large",
    "provider": "alibaba",
    "queriedAt": "2026-08-31T07:38:05.969572165Z",
    "region": "ap-northeast-2",
    "source": "alibaba:DescribeAvailableResource",
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
      }
    ]
  },
  "connectionName": "alibaba-ap-northeast-2",
  "estimatedCost": "$0.0356/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "alibaba-us-west-1",
    "creationDate": "",
    "cspImageName": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "description": "Kernel version is 6.8.0-137-generic, 2026.8.13",
    "details": [
      {
        "key": "BootMode",
        "value": "UEFI-Preferred"
      },
      {
        "key": "ImageId",
        "value": "ubuntu_24_04_x64_20G_alibase_20260810.vhd"
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
        "value": "Kernel version is 6.8.0-137-generic, 2026.8.13"
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
        "value": "v2026.8.13"
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
        "value": "2026-08-13T01:49:37Z"
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
        "value": "ubuntu_24_04_x64_20G_alibase_20260810.vhd"
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
    "fetchedTime": "2026.08.21 13:58:28 Fri",
    "id": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
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
      "ap-southeast-1",
      "ap-southeast-3",
      "ap-southeast-5",
      "ap-southeast-6",
      "ap-southeast-7",
      "ap-southeast-8",
      "cn-beijing",
      "cn-chengdu",
      "cn-fuzhou",
      "cn-guangzhou",
      "cn-hangzhou",
      "cn-heyuan",
      "cn-hongkong",
      "cn-huhehaote",
      "cn-nanjing",
      "cn-qingdao",
      "cn-shanghai",
      "cn-shenzhen",
      "cn-wuhan-lr",
      "cn-wulanchabu",
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
    "uid": "tbm30neucqi9cs1ci7ut"
  },
  "imageId": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
  "imageValidation": {
    "cspResourceId": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "isAvailable": true,
    "resourceId": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "resourceName": "ubuntu_24_04_x64_20G_alibase_20260810.vhd",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "alibaba",
  "regionName": "ap-northeast-2",
  "specDetails": {
    "architecture": "x86_64",
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
    "uid": "tba6uee340r6ln51e0hd",
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
- **Duration:** 7.395s
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
      "lastTransitionTime": "2026-08-31T07:38:14Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:14Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:14Z",
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
  "cspResourceId": "vpc-mj72w81iaxwgvlooalyiz",
  "cspResourceName": "tb4r2jghd6kgo2kkai31",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-alibaba",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "CreationTime",
      "value": "2026-08-31T07:38:07Z"
    },
    {
      "key": "Status",
      "value": "Available"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj72w81iaxwgvlooalyiz"
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
      "value": "tb4r2jghd6kgo2kkai31"
    },
    {
      "key": "VRouterId",
      "value": "vrt-mj7s7s1ukurksqpiil98s"
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
      "value": "{VSwitchId:[vsw-mj7htj8tsmlt6gilfwwjy]}"
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
      "value": "{RouterTableIds:[vtb-mj7x0y1y7ga6b7yuwcefh]}"
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
          "lastTransitionTime": "2026-08-31T07:38:14Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-08-31T07:38:14Z",
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
      "cspResourceId": "vsw-mj7htj8tsmlt6gilfwwjy",
      "cspResourceName": "tb938q7g73nj4tvt2mqb",
      "cspVNetId": "vpc-mj72w81iaxwgvlooalyiz",
      "cspVNetName": "tb4r2jghd6kgo2kkai31",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.3.1.0/24",
      "keyValueList": [
        {
          "key": "VpcId",
          "value": "vpc-mj72w81iaxwgvlooalyiz"
        },
        {
          "key": "Status",
          "value": "Available"
        },
        {
          "key": "CreationTime",
          "value": "2026-08-31T07:38:11Z"
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
          "value": "vsw-mj7htj8tsmlt6gilfwwjy"
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
          "value": "tb938q7g73nj4tvt2mqb"
        },
        {
          "key": "EnabledIpv6",
          "value": "false"
        },
        {
          "key": "RouteTable",
          "value": "{ResourceGroupId:,CreationTime:,Status:,RouteTableType:System,VRouterId:,RouteTableId:vtb-mj7x0y1y7ga6b7yuwcefh,VSwitchIds:{VSwitchId:null},RouteEntrys:{RouteEntry:null}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb938q7g73nj4tvt2mqb",
      "zone": "ap-northeast-2a"
    }
  ],
  "systemLabel": "",
  "uid": "tb4r2jghd6kgo2kkai31"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 8.573s
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
  "cspResourceId": "sg-mj70c752ldpcl7ah5l28",
  "cspResourceName": "tbfqgmcistmlnnm8jb38",
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
      "value": "sg-mj70c752ldpcl7ah5l28"
    },
    {
      "key": "SecurityGroupName",
      "value": "tbfqgmcistmlnnm8jb38"
    },
    {
      "key": "Description",
      "value": "tbfqgmcistmlnnm8jb38"
    },
    {
      "key": "SecurityGroupType",
      "value": "enterprise"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj72w81iaxwgvlooalyiz"
    },
    {
      "key": "CreationTime",
      "value": "2026-08-31T07:38:19Z"
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
  "uid": "tbfqgmcistmlnnm8jb38",
  "vNetId": "test-rdbms-vnet-alibaba"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 6ms
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
- **Duration:** 2m52.985s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=alibaba-ap-northeast-2`
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
          "description": "Alibaba Cloud automatically selects the optimal storage type based on region and instance specification. Recommended for simplicity.",
          "displayName": "Auto-selected Storage Type",
          "maxSize": 32768,
          "minSize": 20,
          "recommendationLevel": "recommended",
          "recommended": true,
          "storageType": "cloud_auto"
        },
        {
          "constraints": "Minimum 20GB storage.",
          "description": "Standard enhanced SSD with good balance of performance and cost. Performance Level 1.",
          "displayName": "Enhanced SSD (ESSD PL1)",
          "maxSize": 32768,
          "minSize": 20,
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
          "maxSize": 32768,
          "minSize": 500,
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
          "maxSize": 32768,
          "minSize": 1500,
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
    "storageSizeRange": {
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
- **Duration:** 14ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "desiredCloud": {
    "csp": "alibaba",
    "region": "ap-northeast-2"
  },
  "sourceRDBMSInstances": [
    {
      "backupRetentionDays": 7,
      "databases": [
        {
          "characterSet": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ],
      "engine": "mysql",
      "engineVersion": "8.0",
      "instanceName": "source-mysql-01",
      "iops": 3000,
      "machineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "memoryMb": 4096,
      "port": 3306,
      "publicAccess": true,
      "storageSizeGb": 100,
      "storageType": "SSD",
      "vcpu": 2
    }
  ]
}
```
```json
// Response Body
{
  "status": "recommended",
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for alibaba (ap-northeast-2)",
  "warnings": [
    "Requested storage type 'SSD' is not supported on target cloud. Replaced with capability recommended storage for instance 'source-mysql-01'."
  ],
  "targetCloud": {
    "csp": "alibaba",
    "region": "ap-northeast-2"
  },
  "targetRDBMSInstances": [
    {
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "rdbmsName": "test-rdbms-alibaba",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n4.large.1",
      "storageType": "cloud_auto",
      "storageSize": 100,
      "adminUserName": "dbadmin",
      "adminUserPassword": "******",
      "vNetId": "test-rdbms-vnet-alibaba",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "publicAccess": true,
      "highAvailability": false,
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ]
    }
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 1m20.067s
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
  "dbInstanceSpec": "mysql.n4.large.1",
  "name": "test-rdbms-alibaba",
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
    "dbInstanceSpec": "mysql.n4.large.1",
    "name": "test-rdbms-alibaba",
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
- **Duration:** 3m39.732s
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
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n4.large.1",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "test-rdbms-alibaba",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "cloud_auto",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-alibaba"
    }
  ],
  "warnings": [
    "Requested storage type 'SSD' is not supported on target cloud. Replaced with capability recommended storage for instance 'source-mysql-01'."
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
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-alibaba`
```json
// Response Body
{
  "resourceType": "rdbms",
  "id": "test-test-rdbms-alibaba",
  "uid": "tbn49stuu7oo1htp85o8",
  "cspResourceName": "tbn49stuu7oo1htp85o8",
  "cspResourceId": "rm-mj71ftn8t67593chy",
  "name": "test-test-rdbms-alibaba",
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
  "description": "Migrated by CM-Beetle from source instance source-mysql-01",
  "status": "Available",
  "conditions": [
    {
      "type": "Ready",
      "status": "True",
      "reason": "Available",
      "lastTransitionTime": "2026-08-31T07:46:10Z"
    },
    {
      "type": "Synced",
      "status": "True",
      "reason": "Available",
      "lastTransitionTime": "2026-08-31T07:46:10Z"
    }
  ],
  "vNetId": "test-rdbms-vnet-alibaba",
  "subnetIds": [
    "subnet-1"
  ],
  "securityGroupIds": [
    "test-rdbms-sg-alibaba"
  ],
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "mysql.n4.large.1",
  "dbInstanceType": "Basic",
  "storageType": "general_essd",
  "storageSize": 100,
  "adminUserName": "dbadmin",
  "highAvailability": false,
  "backupRetentionDays": 7,
  "backupTime": "09:00Z-10:00Z",
  "publicAccess": true,
  "deletionProtection": false,
  "endpoint": "43.108.66.120:3306",
  "tagList": [
    {
      "key": "sys.labelType",
      "value": "rdbms"
    },
    {
      "key": "sys.namespace",
      "value": "default"
    },
    {
      "key": "sys.id",
      "value": "test-test-rdbms-alibaba"
    },
    {
      "key": "sys.name",
      "value": "test-test-rdbms-alibaba"
    },
    {
      "key": "sys.description",
      "value": "Migrated by CM-Beetle from source instance source-mysql-01"
    },
    {
      "key": "sys.connectionName",
      "value": "alibaba-ap-northeast-2"
    },
    {
      "key": "sys.uid",
      "value": "tbn49stuu7oo1htp85o8"
    },
    {
      "key": "sys.cspResourceId",
      "value": "rm-mj71ftn8t67593chy"
    },
    {
      "key": "sys.cspResourceName",
      "value": "tbn49stuu7oo1htp85o8"
    },
    {
      "key": "sys.manager",
      "value": "cb-tumblebug"
    }
  ]
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 7ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-alibaba",
      "uid": "tbn49stuu7oo1htp85o8",
      "cspResourceName": "tbn49stuu7oo1htp85o8",
      "cspResourceId": "rm-mj71ftn8t67593chy",
      "name": "test-test-rdbms-alibaba",
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
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:46:10Z"
        },
        {
          "type": "Synced",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:46:10Z"
        }
      ],
      "vNetId": "test-rdbms-vnet-alibaba",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n4.large.1",
      "dbInstanceType": "Basic",
      "storageType": "general_essd",
      "storageSize": 100,
      "adminUserName": "dbadmin",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "backupTime": "09:00Z-10:00Z",
      "publicAccess": true,
      "deletionProtection": false,
      "endpoint": "43.108.66.120:3306"
    },
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-aws",
      "uid": "tb1k45ucvlpo2hrndcpf",
      "cspResourceName": "tb1k45ucvlpo2hrndcpf",
      "cspResourceId": "tb1k45ucvlpo2hrndcpf",
      "name": "test-test-rdbms-aws",
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
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "status": "Creating",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Creating",
          "message": "RDBMS creation in progress",
          "lastTransitionTime": "2026-08-31T07:38:47Z"
        },
        {
          "type": "Synced",
          "status": "False",
          "reason": "Creating",
          "lastTransitionTime": "2026-08-31T07:38:47Z"
        }
      ],
      "vNetId": "test-rdbms-vnet-aws",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-aws"
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.46",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "storageType": "gp3",
      "storageSize": 100,
      "iops": "3000",
      "adminUserName": "root",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "backupTime": "16:59-17:29",
      "publicAccess": true,
      "deletionProtection": false,
      "endpoint": "tb1k45ucvlpo2hrndcpf.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com:3306",
      "tagList": [
        {
          "key": "Name",
          "value": "tb1k45ucvlpo2hrndcpf"
        }
      ]
    },
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-azure",
      "uid": "tbsn9vgep8ehset1mj4o",
      "cspResourceName": "tbsn9vgep8ehset1mj4o",
      "cspResourceId": "tbsn9vgep8ehset1mj4o",
      "name": "test-test-rdbms-azure",
      "connectionName": "azure-koreacentral",
      "connectionConfig": {
        "configName": "azure-koreacentral",
        "providerName": "azure",
        "driverName": "azure-driver-v1.0.so",
        "credentialName": "azure",
        "credentialHolder": "admin",
        "regionZoneInfoName": "azure-koreacentral",
        "regionZoneInfo": {
          "assignedRegion": "koreacentral",
          "assignedZone": ""
        },
        "regionDetail": {
          "regionId": "koreacentral",
          "regionName": "koreacentral",
          "description": "Korea Central",
          "location": {
            "display": "Korea Central",
            "latitude": 37.5665,
            "longitude": 126.978
          },
          "zones": [
            "1",
            "2",
            "3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:42:41Z"
        },
        {
          "type": "Synced",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:42:41Z"
        }
      ],
      "vNetId": "test-rdbms-vnet-azure",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "dbInstanceType": "Burstable",
      "storageType": "Premium_LRS",
      "storageSize": 32,
      "adminUserName": "azureuser",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "backupTime": "AUTO",
      "publicAccess": true,
      "deletionProtection": false,
      "endpoint": "tbsn9vgep8ehset1mj4o.mysql.database.azure.com:3306"
    },
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-gcp",
      "uid": "tbdg0s34l9kanu2tth22",
      "cspResourceName": "tbdg0s34l9kanu2tth22",
      "cspResourceId": "tbdg0s34l9kanu2tth22",
      "name": "test-test-rdbms-gcp",
      "connectionName": "gcp-us-central1",
      "connectionConfig": {
        "configName": "gcp-us-central1",
        "providerName": "gcp",
        "driverName": "gcp-driver-v1.0.so",
        "credentialName": "gcp",
        "credentialHolder": "admin",
        "regionZoneInfoName": "gcp-us-central1",
        "regionZoneInfo": {
          "assignedRegion": "us-central1",
          "assignedZone": "us-central1-a"
        },
        "regionDetail": {
          "regionId": "us-central1",
          "regionName": "us-central1",
          "description": "Council Bluffs Iowa  USA",
          "location": {
            "display": "Council Bluffs Iowa USA",
            "latitude": 41.2522,
            "longitude": -95.8575
          },
          "zones": [
            "us-central1-a",
            "us-central1-b",
            "us-central1-c",
            "us-central1-f"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:42:34Z"
        },
        {
          "type": "Synced",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:42:34Z"
        }
      ],
      "vNetId": "test-rdbms-vnet-gcp",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "dbInstanceType": "ZONAL",
      "storageType": "PD_SSD",
      "storageSize": 100,
      "adminUserName": "admin",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "backupTime": "03:00",
      "publicAccess": true,
      "deletionProtection": false,
      "encryption": true,
      "endpoint": "34.29.113.205:3306"
    },
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-ibm",
      "uid": "tbi6vdt2vrd1ikjkk8ap",
      "name": "test-test-rdbms-ibm",
      "connectionName": "ibm-us-south",
      "connectionConfig": {
        "configName": "ibm-us-south",
        "providerName": "ibm",
        "driverName": "ibm-driver-v1.0.so",
        "credentialName": "ibm",
        "credentialHolder": "admin",
        "regionZoneInfoName": "ibm-us-south",
        "regionZoneInfo": {
          "assignedRegion": "us-south",
          "assignedZone": "us-south-1"
        },
        "regionDetail": {
          "regionId": "us-south",
          "regionName": "us-south",
          "description": "us-south",
          "location": {
            "display": "Dallas USA",
            "latitude": 32.81248,
            "longitude": -96.77619
          },
          "zones": [
            "us-south-1",
            "us-south-2",
            "us-south-3"
          ]
        },
        "regionRepresentative": true,
        "verified": true
      },
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "status": "Creating",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Creating",
          "message": "RDBMS creation in progress",
          "lastTransitionTime": "2026-08-31T07:38:33Z"
        },
        {
          "type": "Synced",
          "status": "False",
          "reason": "Creating",
          "lastTransitionTime": "2026-08-31T07:38:33Z"
        }
      ],
      "vNetId": "test-rdbms-vnet-ibm",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "multitenant",
      "storageSize": 100,
      "adminUserName": "admin",
      "highAvailability": false,
      "publicAccess": true,
      "deletionProtection": false
    },
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-ncp",
      "uid": "tbvh83qejdrc1o89ai6e",
      "cspResourceName": "tbvh83qejdrc1o89ai6e",
      "cspResourceId": "144826483",
      "name": "test-test-rdbms-ncp",
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
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "status": "Creating",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Creating",
          "message": "RDBMS creation in progress",
          "lastTransitionTime": "2026-08-31T07:39:22Z"
        },
        {
          "type": "Synced",
          "status": "False",
          "reason": "Creating",
          "lastTransitionTime": "2026-08-31T07:39:22Z"
        }
      ],
      "vNetId": "test-rdbms-vnet-ncp",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-ncp"
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL8.0.45",
      "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
      "dbInstanceType": "Stand Alone",
      "storageType": "SSD",
      "storageSize": 10,
      "adminUserName": "",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "backupTime": "00:00",
      "publicAccess": false,
      "deletionProtection": false,
      "endpoint": "db-4a3o7k.vpc-cdb.ntruss.com:3306"
    },
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-nhn",
      "uid": "tbna7alivb64t5mbface",
      "name": "test-test-rdbms-nhn",
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
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "status": "Creating",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Creating",
          "message": "RDBMS creation in progress",
          "lastTransitionTime": "2026-08-31T07:38:46Z"
        },
        {
          "type": "Synced",
          "status": "False",
          "reason": "Creating",
          "lastTransitionTime": "2026-08-31T07:38:46Z"
        }
      ],
      "vNetId": "test-rdbms-vnet-nhn",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-nhn"
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL_V8046",
      "dbInstanceSpec": "m2.c2m4",
      "storageType": "General SSD",
      "storageSize": 100,
      "adminUserName": "myadmin",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "publicAccess": true,
      "nhnDBSGToAllowAllInbound": true,
      "deletionProtection": false
    },
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-tencent",
      "uid": "tbhpvqgrnhe0qd4b99h2",
      "cspResourceName": "tbhpvqgrnhe0qd4b99h2",
      "cspResourceId": "cdb-l97oei3e",
      "name": "test-test-rdbms-tencent",
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
      "description": "Migrated by CM-Beetle from source instance source-mysql-01",
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:42:38Z"
        },
        {
          "type": "Synced",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:42:38Z"
        }
      ],
      "vNetId": "test-rdbms-vnet-tencent",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "dbInstanceType": "NA",
      "storageType": "CLOUD_HSSD",
      "storageSize": 100,
      "adminUserName": "root",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "backupTime": "00:00-12:00",
      "publicAccess": true,
      "deletionProtection": false,
      "endpoint": "kr-cdb-l97oei3e.sql.tencentcdb.com:27450"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 1.546s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-alibaba/database`
```json
// Request Body
{
  "adminUserPassword": "******",
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
- **Duration:** 2.287s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-alibaba/database`
```json
// Response Body
{
  "databases": [
    "sampledb",
    "sampledb_dyn"
  ]
}
```

### 13. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 5ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m25.094s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 7.527s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-alibaba/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 8m42.826s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-alibaba?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 702ms

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 7.306s

