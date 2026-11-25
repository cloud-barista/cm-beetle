# Managed RDBMS (MySQL) Test Report: TENCENT (ap-seoul)

- **Test Case:** Tencent AP-Seoul MySQL Test
- **Access Mode:** Public (publicAccess=true)
- **Date & Time:** 2026-11-25 18:54:32
- **Namespace:** `default`
- **Total Duration:** 12m29.896s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** TENCENT
- **Target Region:** `ap-seoul`
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
- **Duration:** 2.858s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "img-7rotv4ux",
  "specId": "tencent+ap-seoul+sa2.medium4"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "cached": true,
    "instanceType": "SA2.MEDIUM4",
    "provider": "tencent",
    "queriedAt": "2026-11-25T09:54:33.927733455Z",
    "quota": {
      "fitCount": 60,
      "limit": 60,
      "name": "postpaid instances",
      "perInstance": 1,
      "scope": "zone ap-seoul-1",
      "unit": "instances",
      "used": 0
    },
    "refreshAfter": "2026-11-25T09:59:33.927733455Z",
    "region": "ap-seoul",
    "source": "tencent:DescribeZoneInstanceConfigInfos",
    "status": "Available",
    "ttlSeconds": 300,
    "zones": [
      {
        "available": true,
        "status": "SELL",
        "zoneId": "ap-seoul-1"
      },
      {
        "available": true,
        "status": "SELL",
        "zoneId": "ap-seoul-2"
      }
    ]
  },
  "connectionName": "tencent-ap-seoul",
  "estimatedCost": "$0.0400/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "tencent-sa-saopaulo",
    "creationDate": "",
    "cspImageName": "img-7rotv4ux",
    "description": "",
    "details": [
      {
        "key": "ImageId",
        "value": "img-7rotv4ux"
      },
      {
        "key": "OsName",
        "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
      },
      {
        "key": "ImageType",
        "value": "PUBLIC_IMAGE"
      },
      {
        "key": "ImageName",
        "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
      },
      {
        "key": "ImageDescription",
        "value": "Ubuntu Server 22.04 LTS 64bit UEFI"
      },
      {
        "key": "ImageSize",
        "value": "20"
      },
      {
        "key": "Architecture",
        "value": "x86_64"
      },
      {
        "key": "ImageState",
        "value": "NORMAL"
      },
      {
        "key": "Platform",
        "value": "Ubuntu"
      },
      {
        "key": "ImageSource",
        "value": "OFFICIAL"
      },
      {
        "key": "IsSupportCloudinit",
        "value": "true"
      },
      {
        "key": "ImageDeprecated",
        "value": "false"
      }
    ],
    "fetchedTime": "2026.10.29 04:25:24 Thu",
    "id": "img-7rotv4ux",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "img-7rotv4ux",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 20,
    "osDiskType": "NA",
    "osDistribution": "Ubuntu Server 22.04 LTS 64bit UEFI",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 22.04",
    "providerName": "tencent",
    "regionList": [
      "ap-bangkok",
      "ap-beijing",
      "ap-chengdu",
      "ap-chongqing",
      "ap-guangzhou",
      "ap-hongkong",
      "ap-jakarta",
      "ap-nanjing",
      "ap-seoul",
      "ap-shanghai",
      "ap-singapore",
      "ap-tokyo",
      "eu-frankfurt",
      "me-saudi-arabia",
      "na-ashburn",
      "na-siliconvalley",
      "sa-saopaulo"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tbrtv38m36npv5tcj4qk"
  },
  "imageId": "img-7rotv4ux",
  "imageValidation": {
    "cspResourceId": "img-7rotv4ux",
    "isAvailable": true,
    "resourceId": "img-7rotv4ux",
    "resourceName": "img-7rotv4ux",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "tencent",
  "regionName": "ap-seoul",
  "specDetails": {
    "architecture": "x86_64",
    "availabilityCheckedAt": "2026-11-25T09:54:33.927733455Z",
    "availabilityQuota": {
      "fitCount": 60,
      "limit": 60,
      "name": "postpaid instances",
      "perInstance": 1,
      "scope": "zone ap-seoul-1",
      "unit": "instances",
      "used": 0
    },
    "availabilityRefreshAfter": "2026-11-25T09:59:33.927733455Z",
    "availabilityStatus": "Available",
    "availableZones": [
      "ap-seoul-1",
      "ap-seoul-2"
    ],
    "connectionName": "tencent-ap-seoul",
    "costPerHour": 0.04,
    "cspSpecName": "SA2.MEDIUM4",
    "details": [
      {
        "key": "Zone",
        "value": "ap-seoul-1"
      },
      {
        "key": "InstanceType",
        "value": "SA2.MEDIUM4"
      },
      {
        "key": "InstanceChargeType",
        "value": "SPOTPAID"
      },
      {
        "key": "NetworkCard",
        "value": "25"
      },
      {
        "key": "Externals",
        "value": "{UnsupportNetworks:[BASIC,VPC1.0]}"
      },
      {
        "key": "Cpu",
        "value": "2"
      },
      {
        "key": "Memory",
        "value": "4"
      },
      {
        "key": "InstanceFamily",
        "value": "SA2"
      },
      {
        "key": "TypeName",
        "value": "SA2"
      },
      {
        "key": "Status",
        "value": "SOLD_OUT"
      },
      {
        "key": "Price",
        "value": "{UnitPrice:0.04,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.008,UnitPriceSecondStep:0.04,UnitPriceDiscountSecondStep:0.008,UnitPriceThirdStep:0.04,UnitPriceDiscountThirdStep:0.008}"
      },
      {
        "key": "SoldOutReason",
        "value": "ResourcesSoldOut.SpecifiedInstanceType"
      },
      {
        "key": "InstanceBandwidth",
        "value": "1.5"
      },
      {
        "key": "InstancePps",
        "value": "30"
      },
      {
        "key": "StorageBlockAmount",
        "value": "0"
      },
      {
        "key": "CpuType",
        "value": "AMD EPYC™ Rome"
      },
      {
        "key": "Gpu",
        "value": "0"
      },
      {
        "key": "Fpga",
        "value": "0"
      },
      {
        "key": "GpuCount",
        "value": "0"
      },
      {
        "key": "Frequency",
        "value": "2.6GHz/3.3GHz"
      },
      {
        "key": "StatusCategory",
        "value": "WithoutStock"
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
    "id": "tencent+ap-seoul+sa2.medium4",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "tencent+ap-seoul+sa2.medium4",
    "namespace": "system",
    "providerName": "tencent",
    "regionLatitude": 37.566536,
    "regionLongitude": 126.977966,
    "regionName": "ap-seoul",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbn11n2ns4e170ammcfs",
    "vCPU": 2
  },
  "specId": "tencent+ap-seoul+sa2.medium4",
  "specValidation": {
    "cspResourceId": "SA2.MEDIUM4",
    "isAvailable": true,
    "resourceId": "tencent+ap-seoul+sa2.medium4",
    "resourceName": "SA2.MEDIUM4",
    "status": "Available"
  },
  "status": "OK",
  "suggestedZone": "ap-seoul-1"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 2.811s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.4.0.0/16",
  "connectionName": "tencent-ap-seoul",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-tencent",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.4.1.0/24",
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
  "cidrBlock": "10.4.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T09:54:36Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:36Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:36Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "vpc-7yqgk5bc",
  "cspResourceName": "tbtk14i64e4uqhcbtrgi",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-tencent",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "VpcName",
      "value": "tbtk14i64e4uqhcbtrgi"
    },
    {
      "key": "VpcId",
      "value": "vpc-7yqgk5bc"
    },
    {
      "key": "CidrBlock",
      "value": "10.4.0.0/16"
    },
    {
      "key": "IsDefault",
      "value": "false"
    },
    {
      "key": "EnableMulticast",
      "value": "false"
    },
    {
      "key": "CreatedTime",
      "value": "2026-11-25 17:54:34"
    },
    {
      "key": "DnsServerSet",
      "value": "183.60.83.19; 183.60.82.98"
    },
    {
      "key": "DhcpOptionsId",
      "value": "dopt-p7yzs2wn"
    },
    {
      "key": "EnableDhcp",
      "value": "true"
    },
    {
      "key": "EnableRouteVpcPublish",
      "value": "false"
    },
    {
      "key": "EnableRouteVpcPublishIpv6",
      "value": "false"
    }
  ],
  "name": "test-rdbms-vnet-tencent",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:54:36Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:54:36Z",
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
      "cspResourceId": "subnet-f4cf8dsh",
      "cspResourceName": "tbu4rb2gqbn6jnmi557d",
      "cspVNetId": "vpc-7yqgk5bc",
      "cspVNetName": "tbtk14i64e4uqhcbtrgi",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.4.1.0/24",
      "keyValueList": [
        {
          "key": "VpcId",
          "value": "vpc-7yqgk5bc"
        },
        {
          "key": "SubnetId",
          "value": "subnet-f4cf8dsh"
        },
        {
          "key": "SubnetName",
          "value": "tbu4rb2gqbn6jnmi557d"
        },
        {
          "key": "CidrBlock",
          "value": "10.4.1.0/24"
        },
        {
          "key": "IsDefault",
          "value": "false"
        },
        {
          "key": "EnableBroadcast",
          "value": "false"
        },
        {
          "key": "Zone",
          "value": "ap-seoul-1"
        },
        {
          "key": "RouteTableId",
          "value": "rtb-cu09d0xv"
        },
        {
          "key": "CreatedTime",
          "value": "2026-11-25 17:54:34"
        },
        {
          "key": "AvailableIpAddressCount",
          "value": "253"
        },
        {
          "key": "IsRemoteVpcSnat",
          "value": "false"
        },
        {
          "key": "TotalIpAddressCount",
          "value": "253"
        },
        {
          "key": "IsCdcSubnet",
          "value": "0"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbu4rb2gqbn6jnmi557d"
    }
  ],
  "systemLabel": "",
  "uid": "tbtk14i64e4uqhcbtrgi"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 22.393s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "tencent-ap-seoul",
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
  "name": "test-rdbms-sg-tencent",
  "vNetId": "test-rdbms-vnet-tencent"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "sg-mpikmp9v",
  "cspResourceName": "tbl39govn1etj0eo22tq",
  "description": "Pre-requisite SecurityGroup for CM-Beetle RDBMS test",
  "firewallRules": [
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "3306",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "22",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "outbound",
      "Port": "",
      "Protocol": "ALL"
    }
  ],
  "id": "test-rdbms-sg-tencent",
  "isAutoGenerated": false,
  "keyValueList": null,
  "name": "test-rdbms-sg-tencent",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbl39govn1etj0eo22tq",
  "vNetId": "test-rdbms-vnet-tencent"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 7ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=tencent`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "tencent": {
      "dbOperationMethod": "cspNativeApi",
      "note": "SecurityGroup is optional (shares the VM's security group if provided). Other CSPs except AWS ignore securityGroupIds.",
      "storageTypeSelectable": true,
      "supported": true,
      "supportedDBEngines": [
        "mysql"
      ],
      "supportsTag": true
    }
  }
}
```

### 5. Beetle GET RDBMS Capability [✅ SUCCESS]
- **Duration:** 3.231s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=tencent-ap-seoul&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "7-1830",
    "connectionName": "tencent-ap-seoul",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "1000",
      "2000",
      "4000",
      "8000",
      "12000",
      "16000",
      "24000",
      "32000",
      "48000",
      "64000",
      "72000",
      "96000",
      "128000",
      "144000",
      "160000",
      "192000",
      "224000",
      "244000",
      "256000",
      "288000",
      "320000",
      "384000",
      "448000",
      "488000",
      "512000",
      "690000",
      "720000"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "1024",
        "name": "1000",
        "storageSizeRangeGB": {
          "max": 3000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "2048",
        "name": "2000",
        "storageSizeRangeGB": {
          "max": 30000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "4096",
        "name": "4000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "8192",
        "name": "8000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "12288",
        "name": "12000",
        "storageSizeRangeGB": {
          "max": 4000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "16000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "24576",
        "name": "24000",
        "storageSizeRangeGB": {
          "max": 6000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "32000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "49152",
        "name": "48000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "24"
      },
      {
        "memSizeMiB": "65536",
        "name": "64000",
        "storageSizeRangeGB": {
          "max": 6000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "73728",
        "name": "72000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "12"
      },
      {
        "memSizeMiB": "98304",
        "name": "96000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "128000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "147456",
        "name": "144000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "24"
      },
      {
        "memSizeMiB": "163840",
        "name": "160000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "80"
      },
      {
        "memSizeMiB": "196608",
        "name": "192000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "229376",
        "name": "224000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "112"
      },
      {
        "memSizeMiB": "249856",
        "name": "244000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "24"
      },
      {
        "memSizeMiB": "262144",
        "name": "256000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "294912",
        "name": "288000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "327680",
        "name": "320000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "80"
      },
      {
        "memSizeMiB": "393216",
        "name": "384000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "458752",
        "name": "448000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "112"
      },
      {
        "memSizeMiB": "499712",
        "name": "488000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "512000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "706560",
        "name": "690000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "80"
      },
      {
        "memSizeMiB": "737280",
        "name": "720000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "90"
      }
    ],
    "dbOperationMethod": "",
    "defaultStorageType": "",
    "liveSupportedEngines": [
      "mysql"
    ],
    "notes": {
      "storageTypes": [
        {
          "constraints": "Minimum 50GB storage.",
          "description": "Cloud-based enhanced SSD with high performance. Minimum 50 GB.",
          "displayName": "Cloud Enhanced SSD",
          "minSizeGB": 50,
          "recommendationLevel": "standard",
          "storageType": "CLOUD_HSSD"
        },
        {
          "constraints": "Minimum 50GB storage.",
          "description": "Premium cloud storage with good balance. Minimum 50 GB.",
          "displayName": "Cloud Premium",
          "minSizeGB": 50,
          "recommendationLevel": "standard",
          "storageType": "CLOUD_PREMIUM"
        },
        {
          "constraints": "Minimum 50GB storage.",
          "description": "Standard cloud SSD storage. Minimum 50 GB.",
          "displayName": "Cloud SSD",
          "minSizeGB": 50,
          "recommendationLevel": "standard",
          "storageType": "CLOUD_SSD"
        },
        {
          "constraints": "Minimum 50GB storage.",
          "description": "High-performance local SSD storage. Minimum 50 GB.",
          "displayName": "Local SSD",
          "minSizeGB": 50,
          "recommendationLevel": "premium",
          "storageType": "local_ssd"
        }
      ]
    },
    "providerName": "tencent",
    "regionName": "ap-seoul",
    "requiresSecurityGroup": false,
    "requiresSubnet": true,
    "storageSizeRangeGB": {
      "max": 32000,
      "min": 20
    },
    "storageTypeOptions": [
      "CLOUD_HSSD",
      "CLOUD_PREMIUM",
      "CLOUD_SSD",
      "local_ssd"
    ],
    "supportedVersions": [
      "5.7",
      "8.0",
      "8.4"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": false,
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
- **Duration:** 3ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "tencent",
    "region": "ap-seoul"
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
    "adminUserName": "root",
    "backupRetentionDays": 7,
    "highAvailability": false,
    "publicAccess": true
  }
}
```
```json
// Response Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for tencent (ap-seoul)",
  "status": "recommended",
  "targetCloud": {
    "csp": "tencent",
    "region": "ap-seoul"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "root",
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
      "dbInstanceSpec": "4000",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-tencent",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "CLOUD_HSSD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-tencent"
    }
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 13ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "root",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "tencent-ap-seoul",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "4000",
  "name": "rdbms-tencent",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-tencent"
  ],
  "storageSize": 100,
  "storageType": "CLOUD_HSSD",
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-tencent"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "root",
    "adminUserPassword": "******",
    "connectionName": "tencent-ap-seoul",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0",
    "dbInstanceSpec": "4000",
    "name": "rdbms-tencent",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-tencent"
    ],
    "storageSize": 100,
    "storageType": "CLOUD_HSSD",
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-tencent"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 5m5.05s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for tencent (ap-seoul)",
  "status": "recommended",
  "targetCloud": {
    "csp": "tencent",
    "region": "ap-seoul"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "root",
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
      "dbInstanceSpec": "4000",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-tencent",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "CLOUD_HSSD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-tencent"
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
- **Duration:** 17ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent`
```json
// Response Body
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
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 34ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:59:03Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:59:03Z",
          "reason": "Creating",
          "status": "False",
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-alibaba",
      "name": "test-rdbms-alibaba",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "cloud_auto",
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
          "value": "configuring-enhanced-monitoring"
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
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:54:55Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:54:55Z",
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "m1.medium",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-openstack",
      "name": "test-rdbms-openstack",
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
      "uid": "tbvg10m2pcf9d7qe09hr",
      "vNetId": "test-rdbms-vnet-openstack"
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
- **Duration:** 1.066s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent/database`
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
- **Duration:** 1.855s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent/database`
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
- **Duration:** 1.11s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent/secure-transport`
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
- **Duration:** 8ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m35.206s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 8.346s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m30.196s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 32.22s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 3.471s

