# Managed RDBMS (MariaDB) Test Report: ALIBABA (ap-northeast-2)

- **Test Case:** Alibaba AP-Northeast-2 (Seoul) MariaDB Test
- **Access Mode:** Public (publicAccess=true)
- **Date & Time:** 2026-11-25 19:27:54
- **Namespace:** `default`
- **Total Duration:** 15m35.363s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** ALIBABA
- **Target Region:** `ap-northeast-2`
- **Namespace:** `default`
- **Test Date:** 2026-11-25 19:27:54

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
- **Duration:** 1.862s
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
    "queriedAt": "2026-11-25T10:27:55.738422136Z",
    "refreshAfter": "2026-11-25T11:27:55.738422136Z",
    "region": "ap-northeast-2",
    "source": "alibaba:DescribeAvailableResource",
    "status": "Available",
    "ttlSeconds": 3600,
    "zones": [
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
      },
      {
        "available": true,
        "status": "Available",
        "supportedDisks": [
          "cloud_auto",
          "cloud_essd"
        ],
        "zoneId": "ap-northeast-2b"
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
    "availabilityCheckedAt": "2026-11-25T10:27:55.738422136Z",
    "availabilityRefreshAfter": "2026-11-25T11:27:55.738422136Z",
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
  "suggestedSystemDisk": "cloud_essd_entry",
  "suggestedZone": "ap-northeast-2a"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 7.4s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.13.0.0/16",
  "connectionName": "alibaba-ap-northeast-2",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-mariadb-alibaba",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.13.1.0/24",
      "name": "subnet-1",
      "zone": "ap-northeast-2a"
    }
  ]
}
```
```json
// Response Body
{
  "associatedObjectList": null,
  "cidrBlock": "10.13.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T10:28:03Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T10:28:03Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-11-25T10:28:03Z",
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
  "cspResourceId": "vpc-mj7fosavcqa9tajq8qsyy",
  "cspResourceName": "tbmn0752el49bmc647rn",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-mariadb-alibaba",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "CreationTime",
      "value": "2026-11-25T10:27:55Z"
    },
    {
      "key": "Status",
      "value": "Available"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7fosavcqa9tajq8qsyy"
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
      "value": "tbmn0752el49bmc647rn"
    },
    {
      "key": "VRouterId",
      "value": "vrt-mj7geazw9m9822uysc7jf"
    },
    {
      "key": "CidrBlock",
      "value": "10.13.0.0/16"
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
      "value": "{VSwitchId:[vsw-mj7mlhi0t2mhxca0s54s7]}"
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
      "value": "{RouterTableIds:[vtb-mj7osfow5k73c8wrvhsfg]}"
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
  "name": "test-rdbms-vnet-mariadb-alibaba",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:28:03Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:28:03Z",
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
      "cspResourceId": "vsw-mj7mlhi0t2mhxca0s54s7",
      "cspResourceName": "tbothpirh3am43b8ta16",
      "cspVNetId": "vpc-mj7fosavcqa9tajq8qsyy",
      "cspVNetName": "tbmn0752el49bmc647rn",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.13.1.0/24",
      "keyValueList": [
        {
          "key": "VpcId",
          "value": "vpc-mj7fosavcqa9tajq8qsyy"
        },
        {
          "key": "Status",
          "value": "Available"
        },
        {
          "key": "CreationTime",
          "value": "2026-11-25T10:27:58Z"
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
          "value": "vsw-mj7mlhi0t2mhxca0s54s7"
        },
        {
          "key": "CidrBlock",
          "value": "10.13.1.0/24"
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
          "value": "tbothpirh3am43b8ta16"
        },
        {
          "key": "EnabledIpv6",
          "value": "false"
        },
        {
          "key": "RouteTable",
          "value": "{ResourceGroupId:,CreationTime:,Status:,RouteTableType:System,VRouterId:,RouteTableId:vtb-mj7osfow5k73c8wrvhsfg,VSwitchIds:{VSwitchId:null},RouteEntrys:{RouteEntry:null}}"
        },
        {
          "key": "Tags",
          "value": "{Tag:null}"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbothpirh3am43b8ta16",
      "zone": "ap-northeast-2a"
    }
  ],
  "systemLabel": "",
  "uid": "tbmn0752el49bmc647rn"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 871ms
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
  "name": "test-rdbms-sg-mariadb-alibaba",
  "vNetId": "test-rdbms-vnet-mariadb-alibaba"
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
  "cspResourceId": "sg-mj7e5z2uqzhddk96jr8c",
  "cspResourceName": "tb9ehgnl877annadvmc6",
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
  "id": "test-rdbms-sg-mariadb-alibaba",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "SecurityGroupId",
      "value": "sg-mj7e5z2uqzhddk96jr8c"
    },
    {
      "key": "SecurityGroupName",
      "value": "tb9ehgnl877annadvmc6"
    },
    {
      "key": "Description",
      "value": "tb9ehgnl877annadvmc6"
    },
    {
      "key": "SecurityGroupType",
      "value": "enterprise"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7fosavcqa9tajq8qsyy"
    },
    {
      "key": "CreationTime",
      "value": "2026-11-25T10:28:02Z"
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
  "name": "test-rdbms-sg-mariadb-alibaba",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tb9ehgnl877annadvmc6",
  "vNetId": "test-rdbms-vnet-mariadb-alibaba"
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
- **Duration:** 40.459s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=alibaba-ap-northeast-2&dbEngine=mariadb`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "7-730",
    "connectionName": "alibaba-ap-northeast-2",
    "dbEngine": "mariadb",
    "dbInstanceSpecOptions": [
      "mariadb.n2.medium.2c",
      "mariadb.n2.small.2c",
      "mariadb.x2.2xlarge.2c",
      "mariadb.x2.large.2c",
      "mariadb.x2.xlarge.2c",
      "mariadb.x4.2xlarge.2c",
      "mariadb.x4.4xlarge.2c",
      "mariadb.x4.8xlarge.2c",
      "mariadb.x4.large.2c",
      "mariadb.x4.xlarge.2c",
      "mariadb.x8.2xlarge.2c",
      "mariadb.x8.4xlarge.2c",
      "mariadb.x8.8xlarge.2c"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "4096",
        "name": "mariadb.n2.medium.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "2048",
        "name": "mariadb.n2.small.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "32768",
        "name": "mariadb.x2.2xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "8192",
        "name": "mariadb.x2.large.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "mariadb.x2.xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "mariadb.x4.2xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "mariadb.x4.4xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "229376",
        "name": "mariadb.x4.8xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "56 "
      },
      {
        "memSizeMiB": "16384",
        "name": "mariadb.x4.large.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "mariadb.x4.xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "mariadb.x8.2xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "mariadb.x8.4xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "491520",
        "name": "mariadb.x8.8xlarge.2c",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "56 "
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
      "max": 32000,
      "min": 20
    },
    "storageTypeOptions": [
      "cloud_essd",
      "cloud_essd2",
      "cloud_essd3",
      "cloud_ssd"
    ],
    "supportedVersions": [
      "10.3",
      "10.6"
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
- **Duration:** 6ms
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
        "engine": "mariadb",
        "engineVersion": "10.6",
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
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "mariadb.n2.medium.2c",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-mariadb-alibaba",
      "securityGroupIds": [
        "test-rdbms-sg-mariadb-alibaba"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "cloud_essd",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-mariadb-alibaba"
    }
  ],
  "warnings": [
    "instance 'Source MySQL 01' database 'sampledb': omitted collation 'utf8mb4_unicode_ci' because Alibaba Cloud MariaDB OpenAPI accepts charset only"
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 20.983s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "dbadmin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "alibaba-ap-northeast-2",
  "dbEngine": "mariadb",
  "dbEngineVersion": "10.6",
  "dbInstanceSpec": "mariadb.n2.medium.2c",
  "name": "rdbms-mariadb-alibaba",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-mariadb-alibaba"
  ],
  "storageSize": 100,
  "storageType": "cloud_essd",
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-mariadb-alibaba"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "dbadmin",
    "adminUserPassword": "******",
    "connectionName": "alibaba-ap-northeast-2",
    "dbEngine": "mariadb",
    "dbEngineVersion": "10.6",
    "dbInstanceSpec": "mariadb.n2.medium.2c",
    "name": "rdbms-mariadb-alibaba",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-mariadb-alibaba"
    ],
    "storageSize": 100,
    "storageType": "cloud_essd",
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-mariadb-alibaba"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 2m30.505s
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
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "mariadb.n2.medium.2c",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-mariadb-alibaba",
      "securityGroupIds": [
        "test-rdbms-sg-mariadb-alibaba"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "cloud_essd",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-mariadb-alibaba"
    }
  ],
  "warnings": [
    "instance 'Source MySQL 01' database 'sampledb': omitted collation 'utf8mb4_unicode_ci' because Alibaba Cloud MariaDB OpenAPI accepts charset only"
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
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "09:00Z-10:00Z",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T10:31:22Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T10:31:22Z",
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
  "cspResourceId": "rm-mj7yv170t22mnj1bz",
  "cspResourceName": "tbg504bm99cilsq7lbb2",
  "dbEngine": "mariadb",
  "dbEngineVersion": "10.6",
  "dbInstanceSpec": "mariadb.n2.medium.2c",
  "dbInstanceType": "HighAvailability",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "43.108.66.115:3306",
  "highAvailability": true,
  "id": "test-rdbms-mariadb-alibaba",
  "keyValueList": [
    {
      "key": "Engine",
      "value": "MariaDB"
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
      "value": "rm-mj7yv170t22mnj1bz-20261125182909"
    },
    {
      "key": "DBInstanceNetType",
      "value": "Intranet"
    },
    {
      "key": "DBInstanceClass",
      "value": "mariadb.n2.medium.2c"
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
      "value": "rm-mj7yv170t22mnj1bz"
    },
    {
      "key": "VpcId",
      "value": "vpc-mj7fosavcqa9tajq8qsyy"
    },
    {
      "key": "DBInstanceMemory",
      "value": "4096"
    },
    {
      "key": "DBInstanceStorageType",
      "value": "cloud_essd"
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
      "value": "10.6"
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
      "value": "mariadb106_20260531"
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
      "key": "MasterZone",
      "value": "ap-northeast-2a"
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
      "value": "HighAvailability"
    },
    {
      "key": "InsId",
      "value": "1"
    },
    {
      "key": "DBInstanceDescription",
      "value": "tbg504bm99cilsq7lbb2"
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
      "value": "0"
    },
    {
      "key": "ResourceGroupId",
      "value": "rg-acfnvekhilw5kmy"
    },
    {
      "key": "DispenseMode",
      "value": "MultiAVZDispenseMode"
    },
    {
      "key": "TempUpgradeRecoveryMemory",
      "value": "0"
    },
    {
      "key": "ProxyType",
      "value": "0"
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
      "value": "vsw-mj7mlhi0t2mhxca0s54s7"
    },
    {
      "key": "CreationTime",
      "value": "2026-11-25T10:29:13Z"
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
      "value": "rm-mj7yv170t22mnj1bz.mariadb.ap-northeast-2.rds.aliyuncs.com"
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
      "value": "{SlaveZone:[{ZoneId:ap-northeast-2a}]}"
    }
  ],
  "name": "test-rdbms-mariadb-alibaba",
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-mariadb-alibaba"
  ],
  "status": "Available",
  "storageSize": 100,
  "storageType": "cloud_essd",
  "subnetIds": [
    "subnet-1"
  ],
  "uid": "tbg504bm99cilsq7lbb2",
  "vNetId": "test-rdbms-vnet-mariadb-alibaba"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 21ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "backupTime": "09:00Z-10:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:31:22Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:31:22Z",
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
      "cspResourceId": "rm-mj7yv170t22mnj1bz",
      "cspResourceName": "tbg504bm99cilsq7lbb2",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "mariadb.n2.medium.2c",
      "dbInstanceType": "HighAvailability",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "43.108.66.115:3306",
      "highAvailability": true,
      "id": "test-rdbms-mariadb-alibaba",
      "keyValueList": [
        {
          "key": "Engine",
          "value": "MariaDB"
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
          "value": "rm-mj7yv170t22mnj1bz-20261125182909"
        },
        {
          "key": "DBInstanceNetType",
          "value": "Intranet"
        },
        {
          "key": "DBInstanceClass",
          "value": "mariadb.n2.medium.2c"
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
          "value": "rm-mj7yv170t22mnj1bz"
        },
        {
          "key": "VpcId",
          "value": "vpc-mj7fosavcqa9tajq8qsyy"
        },
        {
          "key": "DBInstanceMemory",
          "value": "4096"
        },
        {
          "key": "DBInstanceStorageType",
          "value": "cloud_essd"
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
          "value": "10.6"
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
          "value": "mariadb106_20260531"
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
          "key": "MasterZone",
          "value": "ap-northeast-2a"
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
          "value": "HighAvailability"
        },
        {
          "key": "InsId",
          "value": "1"
        },
        {
          "key": "DBInstanceDescription",
          "value": "tbg504bm99cilsq7lbb2"
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
          "value": "0"
        },
        {
          "key": "ResourceGroupId",
          "value": "rg-acfnvekhilw5kmy"
        },
        {
          "key": "DispenseMode",
          "value": "MultiAVZDispenseMode"
        },
        {
          "key": "TempUpgradeRecoveryMemory",
          "value": "0"
        },
        {
          "key": "ProxyType",
          "value": "0"
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
          "value": "vsw-mj7mlhi0t2mhxca0s54s7"
        },
        {
          "key": "CreationTime",
          "value": "2026-11-25T10:29:13Z"
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
          "value": "rm-mj7yv170t22mnj1bz.mariadb.ap-northeast-2.rds.aliyuncs.com"
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
          "value": "{SlaveZone:[{ZoneId:ap-northeast-2a}]}"
        }
      ],
      "name": "test-rdbms-mariadb-alibaba",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-mariadb-alibaba"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "cloud_essd",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbg504bm99cilsq7lbb2",
      "vNetId": "test-rdbms-vnet-mariadb-alibaba"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "15:06-15:36",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:29:09Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:29:09Z",
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
      "cspResourceId": "tb55rm4aqne7bvtmh109",
      "cspResourceName": "tb55rm4aqne7bvtmh109",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6.28",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-aws",
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
          "value": "arn:aws:rds:ap-northeast-2:635484366616:db:tb55rm4aqne7bvtmh109"
        },
        {
          "key": "DBInstanceClass",
          "value": "db.t3.medium"
        },
        {
          "key": "DBInstanceIdentifier",
          "value": "tb55rm4aqne7bvtmh109"
        },
        {
          "key": "DBInstanceStatus",
          "value": "creating"
        },
        {
          "key": "DBParameterGroups",
          "value": "{DBParameterGroupName:default.mariadb10.6,ParameterApplyStatus:in-sync}"
        },
        {
          "key": "DBSubnetGroup",
          "value": "{DBSubnetGroupArn:null,DBSubnetGroupDescription:CB-Spider RDBMS subnet group for cb-spider-tb55rm4aqne7bvtmh109,DBSubnetGroupName:cb-spider-tb55rm4aqne7bvtmh109,SubnetGroupStatus:Complete,Subnets:[{SubnetAvailabilityZone:{Name:ap-northeast-2a},SubnetIdentifier:subnet-0332b3b067d7b2be3,SubnetOutpost:{Arn:null},SubnetStatus:Active},{SubnetAvailabilityZone:{Name:ap-northeast-2c},SubnetIdentifier:subnet-0fcd6322d6b4d1720,SubnetOutpost:{Arn:null},SubnetStatus:Active}],VpcId:vpc-0df32b14dbcdda989}"
        },
        {
          "key": "DbInstancePort",
          "value": "0"
        },
        {
          "key": "DbiResourceId",
          "value": "db-6S54HILSY73L2WFC6GXJJLPUJM"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Engine",
          "value": "mariadb"
        },
        {
          "key": "EngineVersion",
          "value": "10.6.28"
        },
        {
          "key": "IAMDatabaseAuthenticationEnabled",
          "value": "false"
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
          "value": "root"
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
          "value": "{OptionGroupName:default:mariadb-10-6,Status:in-sync}"
        },
        {
          "key": "PendingModifiedValues",
          "value": "{AllocatedStorage:null,BackupRetentionPeriod:null,CACertificateIdentifier:null,DBInstanceClass:null,DBInstanceIdentifier:null,DBSubnetGroupName:null,EngineVersion:null,IAMDatabaseAuthenticationEnabled:null,Iops:null,LicenseModel:null,MasterUserPassword:****,MultiAZ:null,PendingCloudwatchLogsExports:null,Port:null,ProcessorFeatures:null,StorageType:null}"
        },
        {
          "key": "PerformanceInsightsEnabled",
          "value": "false"
        },
        {
          "key": "PreferredBackupWindow",
          "value": "15:06-15:36"
        },
        {
          "key": "PreferredMaintenanceWindow",
          "value": "sun:16:40-sun:17:10"
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
          "value": "{Key:Name,Value:tb55rm4aqne7bvtmh109}"
        },
        {
          "key": "VpcSecurityGroups",
          "value": "{Status:active,VpcSecurityGroupId:sg-0dec51a4099d5aa2c}"
        }
      ],
      "name": "test-rdbms-mariadb-aws",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-mariadb-aws"
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
          "value": "tb55rm4aqne7bvtmh109"
        }
      ],
      "uid": "tb55rm4aqne7bvtmh109",
      "vNetId": "test-rdbms-vnet-mariadb-aws"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:28:37Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:28:37Z",
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
      "dbEngine": "mariadb",
      "dbEngineVersion": "MARIADB_V101118",
      "dbInstanceSpec": "m2.c2m4",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-nhn",
      "name": "test-rdbms-mariadb-nhn",
      "nhnDBSGToAllowAllInbound": true,
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-mariadb-nhn"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "General SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tb5jnso3laaus5n58f8i",
      "vNetId": "test-rdbms-vnet-mariadb-nhn"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:28:03Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:28:03Z",
          "reason": "Creating",
          "status": "False",
          "type": "Synced"
        }
      ],
      "connectionConfig": {
        "configName": "openstack-regionone",
        "credentialHolder": "admin",
        "credentialName": "openstack",
        "driverName": "openstack-driver-v1.0.so",
        "providerName": "openstack",
        "regionDetail": {
          "description": "Korea Daejeon (Internal)",
          "location": {
            "display": "South Korea (Daejeon)",
            "latitude": 36.3804,
            "longitude": 127.365
          },
          "regionId": "RegionOne",
          "regionName": "regionone",
          "zones": [
            "nova"
          ]
        },
        "regionRepresentative": true,
        "regionZoneInfo": {
          "assignedRegion": "RegionOne",
          "assignedZone": "nova"
        },
        "regionZoneInfoName": "openstack-regionone",
        "verified": true
      },
      "connectionName": "openstack-regionone",
      "dbEngine": "mariadb",
      "dbEngineVersion": "11.4",
      "dbInstanceSpec": "m1.medium",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-openstack",
      "name": "test-rdbms-mariadb-openstack",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-openstack"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "RBD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tb0lg7ct0lnh5v9libpg",
      "vNetId": "test-rdbms-vnet-openstack"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 1.049s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba/database`
```json
// Request Body
{
  "charset": "utf8mb4",
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
- **Duration:** 1.611s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba/database`
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
- **Duration:** 806ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba/secure-transport`
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
  "engine": "mariadb",
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
- **Duration:** 2m42.692s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 17.153s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 8m41.532s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-alibaba?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 917ms

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 7.469s

