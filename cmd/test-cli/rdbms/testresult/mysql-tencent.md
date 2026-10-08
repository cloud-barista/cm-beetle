# Managed RDBMS (MySQL) Test Report: TENCENT (ap-seoul)

- **Test Case:** Tencent AP-Seoul MySQL Test
- **Date & Time:** 2026-10-08 13:20:42
- **Namespace:** `default`
- **Total Duration:** 10m15.921s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** TENCENT
- **Target Region:** `ap-seoul`
- **Namespace:** `default`
- **Test Date:** 2026-10-08 13:20:42

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
- **Duration:** 4.376s
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
    "instanceType": "SA2.MEDIUM4",
    "provider": "tencent",
    "queriedAt": "2026-10-08T04:20:45.00205451Z",
    "region": "ap-seoul",
    "source": "tencent:DescribeZoneInstanceConfigInfos",
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
    "fetchedTime": "2026.08.21 13:57:36 Fri",
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
    "uid": "tbm5u74berjs1hfbpgol"
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
        "value": "SELL"
      },
      {
        "key": "Price",
        "value": "{UnitPrice:0.04,ChargeUnit:HOUR,Discount:20,UnitPriceDiscount:0.008,UnitPriceSecondStep:0.04,UnitPriceDiscountSecondStep:0.008,UnitPriceThirdStep:0.04,UnitPriceDiscountThirdStep:0.008}"
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
        "value": "UnderStock"
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
    "uid": "tbus5den09dtgg280ngk",
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
- **Duration:** 3.652s
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
      "lastTransitionTime": "2026-10-08T04:20:50Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-08T04:20:50Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-10-08T04:20:50Z",
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
  "cspResourceId": "vpc-9qvk2ywy",
  "cspResourceName": "tbd05s66av3agkno30hc",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-tencent",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "VpcName",
      "value": "tbd05s66av3agkno30hc"
    },
    {
      "key": "VpcId",
      "value": "vpc-9qvk2ywy"
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
      "value": "2026-10-08 12:20:47"
    },
    {
      "key": "DnsServerSet",
      "value": "183.60.83.19; 183.60.82.98"
    },
    {
      "key": "DhcpOptionsId",
      "value": "dopt-m5gbrklz"
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
          "lastTransitionTime": "2026-10-08T04:20:50Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:20:50Z",
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
      "cspResourceId": "subnet-g6de115l",
      "cspResourceName": "tb0of6vpv0h3tc2rskjh",
      "cspVNetId": "vpc-9qvk2ywy",
      "cspVNetName": "tbd05s66av3agkno30hc",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.4.1.0/24",
      "keyValueList": [
        {
          "key": "VpcId",
          "value": "vpc-9qvk2ywy"
        },
        {
          "key": "SubnetId",
          "value": "subnet-g6de115l"
        },
        {
          "key": "SubnetName",
          "value": "tb0of6vpv0h3tc2rskjh"
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
          "value": "rtb-h3m2r24n"
        },
        {
          "key": "CreatedTime",
          "value": "2026-10-08 12:20:49"
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
      "uid": "tb0of6vpv0h3tc2rskjh"
    }
  ],
  "systemLabel": "",
  "uid": "tbd05s66av3agkno30hc"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 12.354s
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
  "cspResourceId": "sg-jfnyscoh",
  "cspResourceName": "tbjvbv4cpnilisofbaol",
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
  "uid": "tbjvbv4cpnilisofbaol",
  "vNetId": "test-rdbms-vnet-tencent"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 4ms
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
- **Duration:** 3.725s
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
          "max": 3000,
          "min": 25
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
          "max": 30000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
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
          "max": 6000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "49152",
        "name": "48000",
        "storageSizeRangeGB": {
          "max": 6000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "64000",
        "storageSizeRangeGB": {
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
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
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "131072",
        "name": "128000",
        "storageSizeRangeGB": {
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
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
          "max": 12000,
          "min": 25
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "24"
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
        "vCpuCount": "64"
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
- **Duration:** 15ms
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
- **Duration:** 3m47.522s
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
- **Duration:** 5ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "00:00-12:00",
  "conditions": [
    {
      "lastTransitionTime": "2026-10-08T04:24:52Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-08T04:24:52Z",
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
  "cspResourceId": "cdb-5d1t4gzi",
  "cspResourceName": "tbgg0091uj25fuq62c78",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "4000",
  "dbInstanceType": "NA",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "kr-cdb-5d1t4gzi.sql.tencentcdb.com:29206",
  "highAvailability": false,
  "id": "test-rdbms-tencent",
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
  "uid": "tbgg0091uj25fuq62c78",
  "vNetId": "test-rdbms-vnet-tencent"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 6ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "backupTime": "18:34-19:04",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:30Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:30Z",
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
      "cspResourceId": "tb9i05t58qdva0pqii7q",
      "cspResourceName": "tb9i05t58qdva0pqii7q",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.46",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-aws",
      "iops": "3000",
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
          "value": "tb9i05t58qdva0pqii7q"
        }
      ],
      "uid": "tb9i05t58qdva0pqii7q",
      "vNetId": "test-rdbms-vnet-aws"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:02Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:02Z",
          "reason": "Creating",
          "status": "False",
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-azure",
      "name": "test-rdbms-azure",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "status": "Creating",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbncr09n28klnsbdp033",
      "vNetId": "test-rdbms-vnet-azure"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:57Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:57Z",
          "reason": "Creating",
          "status": "False",
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-gcp",
      "name": "test-rdbms-gcp",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "PD_SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tb5ivnnbbt2sn5eqjole",
      "vNetId": "test-rdbms-vnet-gcp"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:12Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:12Z",
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
      "dbInstanceSpec": "b3c.4x16.encrypted",
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
      "uid": "tbjg8pdr8lgovsqj2clo",
      "vNetId": "test-rdbms-vnet-ibm"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "16:00Z-17:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:24:46Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:24:46Z",
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
      "cspResourceId": "rm-mj75n3v6blq0q6gef",
      "cspResourceName": "tb4himg14kr1h819depl",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "mariadb.n2.medium.2c",
      "dbInstanceType": "HighAvailability",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "43.108.66.234:3306",
      "highAvailability": true,
      "id": "test-rdbms-mariadb-alibaba",
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
      "uid": "tb4himg14kr1h819depl",
      "vNetId": "test-rdbms-vnet-mariadb-alibaba"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "16:33-17:03",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:22:18Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:22:18Z",
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
      "cspResourceId": "tbnon77lqh7fau59hm9p",
      "cspResourceName": "tbnon77lqh7fau59hm9p",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6.27",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-aws",
      "iops": "3000",
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
          "value": "tbnon77lqh7fau59hm9p"
        }
      ],
      "uid": "tbnon77lqh7fau59hm9p",
      "vNetId": "test-rdbms-vnet-mariadb-aws"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:25Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:25Z",
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
      "uid": "tbfcn32o0g0dc0nn4ge4",
      "vNetId": "test-rdbms-vnet-mariadb-nhn"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:23Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:23Z",
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
      "uid": "tbgv5t06psvjaa5r2i0d",
      "vNetId": "test-rdbms-vnet-nhn"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "00:00-12:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:24:52Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:24:52Z",
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
      "cspResourceId": "cdb-5d1t4gzi",
      "cspResourceName": "tbgg0091uj25fuq62c78",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "kr-cdb-5d1t4gzi.sql.tencentcdb.com:29206",
      "highAvailability": false,
      "id": "test-rdbms-tencent",
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
      "uid": "tbgg0091uj25fuq62c78",
      "vNetId": "test-rdbms-vnet-tencent"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 924ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent/database`
```json
// Request Body
{
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
- **Duration:** 1.051s
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
- **Duration:** 1.604s
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
- **Duration:** 10ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m39.547s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 16.832s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 1m59.364s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-tencent?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 1.791s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 3.136s

