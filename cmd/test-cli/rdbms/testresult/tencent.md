# Managed RDBMS Test Report: TENCENT (ap-seoul)

- **Test Case:** Tencent AP-Seoul MySQL Test
- **Date & Time:** 2026-08-31 16:38:03
- **Namespace:** `default`
- **Total Duration:** 11m31.895s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** TENCENT
- **Target Region:** `ap-seoul`
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
- **Duration:** 3.713s
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
    "queriedAt": "2026-08-31T07:38:06.954398001Z",
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
- **Duration:** 30.446s
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
      "lastTransitionTime": "2026-08-31T07:38:11Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:11Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:11Z",
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
  "cspResourceId": "vpc-iq9q8guc",
  "cspResourceName": "tb3ud3ftgm7n4678t1k0",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-tencent",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "VpcName",
      "value": "tb3ud3ftgm7n4678t1k0"
    },
    {
      "key": "VpcId",
      "value": "vpc-iq9q8guc"
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
      "value": "2026-08-31 15:38:08"
    },
    {
      "key": "DnsServerSet",
      "value": "183.60.83.19; 183.60.82.98"
    },
    {
      "key": "DhcpOptionsId",
      "value": "dopt-cxnyvfy5"
    },
    {
      "key": "EnableDhcp",
      "value": "true"
    }
  ],
  "name": "test-rdbms-vnet-tencent",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-08-31T07:38:11Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-08-31T07:38:11Z",
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
      "cspResourceId": "subnet-s0s91fbx",
      "cspResourceName": "tbler7141c312f7ioj75",
      "cspVNetId": "vpc-iq9q8guc",
      "cspVNetName": "tb3ud3ftgm7n4678t1k0",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.4.1.0/24",
      "keyValueList": [
        {
          "key": "VpcId",
          "value": "vpc-iq9q8guc"
        },
        {
          "key": "SubnetId",
          "value": "subnet-s0s91fbx"
        },
        {
          "key": "SubnetName",
          "value": "tbler7141c312f7ioj75"
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
          "value": "rtb-0colmqzp"
        },
        {
          "key": "CreatedTime",
          "value": "2026-08-31 15:38:10"
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
      "uid": "tbler7141c312f7ioj75"
    }
  ],
  "systemLabel": "",
  "uid": "tb3ud3ftgm7n4678t1k0"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 11.628s
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
  "cspResourceId": "sg-k3u5wp01",
  "cspResourceName": "tbibe79ti3pqn2f8p2ds",
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
  "uid": "tbibe79ti3pqn2f8p2ds",
  "vNetId": "test-rdbms-vnet-tencent"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 5ms
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
- **Duration:** 3.052s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=tencent-ap-seoul`
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
          "max": 3000,
          "min": 25
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
          "max": 30000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
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
        "vCpuCount": "32"
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
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
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
        "vCpuCount": "32"
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
          "max": 32000,
          "min": 20
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "128"
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
          "minSize": 50,
          "recommendationLevel": "standard",
          "storageType": "CLOUD_HSSD"
        },
        {
          "constraints": "Minimum 50GB storage.",
          "description": "Premium cloud storage with good balance. Minimum 50 GB.",
          "displayName": "Cloud Premium",
          "minSize": 50,
          "recommendationLevel": "standard",
          "storageType": "CLOUD_PREMIUM"
        },
        {
          "constraints": "Minimum 50GB storage.",
          "description": "Standard cloud SSD storage. Minimum 50 GB.",
          "displayName": "Cloud SSD",
          "minSize": 50,
          "recommendationLevel": "standard",
          "storageType": "CLOUD_SSD"
        },
        {
          "constraints": "Minimum 50GB storage.",
          "description": "High-performance local SSD storage. Minimum 50 GB.",
          "displayName": "Local SSD",
          "minSize": 50,
          "recommendationLevel": "premium",
          "storageType": "local_ssd"
        }
      ]
    },
    "providerName": "tencent",
    "regionName": "ap-seoul",
    "requiresSecurityGroup": false,
    "requiresSubnet": true,
    "storageSizeRange": {
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
      "5.6",
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
  "desiredCloud": {
    "csp": "tencent",
    "region": "ap-seoul"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for tencent (ap-seoul)",
  "warnings": [
    "Requested storage type 'SSD' is not supported on target cloud. Replaced with capability recommended storage for instance 'source-mysql-01'."
  ],
  "targetCloud": {
    "csp": "tencent",
    "region": "ap-seoul"
  },
  "targetRDBMSInstances": [
    {
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "rdbmsName": "test-rdbms-tencent",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "storageType": "CLOUD_HSSD",
      "storageSize": 100,
      "adminUserName": "root",
      "adminUserPassword": "******",
      "vNetId": "test-rdbms-vnet-tencent",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
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
- **Duration:** 12ms
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
  "name": "test-rdbms-tencent",
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
    "name": "test-rdbms-tencent",
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
- **Duration:** 3m59.379s
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
      "rdbmsName": "test-rdbms-tencent",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "CLOUD_HSSD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-tencent"
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
- **Duration:** 3ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-tencent`
```json
// Response Body
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
  "endpoint": "kr-cdb-l97oei3e.sql.tencentcdb.com:27450",
  "tagList": [
    {
      "key": "sys.labelType",
      "value": "rdbms"
    },
    {
      "key": "sys.manager",
      "value": "cb-tumblebug"
    },
    {
      "key": "sys.connectionName",
      "value": "tencent-ap-seoul"
    },
    {
      "key": "sys.cspResourceName",
      "value": "tbhpvqgrnhe0qd4b99h2"
    },
    {
      "key": "sys.cspResourceId",
      "value": "cdb-l97oei3e"
    },
    {
      "key": "sys.namespace",
      "value": "default"
    },
    {
      "key": "sys.description",
      "value": "Migrated by CM-Beetle from source instance source-mysql-01"
    },
    {
      "key": "sys.uid",
      "value": "tbhpvqgrnhe0qd4b99h2"
    },
    {
      "key": "sys.name",
      "value": "test-test-rdbms-tencent"
    },
    {
      "key": "sys.id",
      "value": "test-test-rdbms-tencent"
    }
  ]
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 5ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-alibaba",
      "uid": "tbn49stuu7oo1htp85o8",
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
      "status": "Creating",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Creating",
          "message": "RDBMS creation in progress",
          "lastTransitionTime": "2026-08-31T07:42:36Z"
        },
        {
          "type": "Synced",
          "status": "False",
          "reason": "Creating",
          "lastTransitionTime": "2026-08-31T07:42:36Z"
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
      "storageType": "cloud_auto",
      "storageSize": 100,
      "adminUserName": "dbadmin",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "publicAccess": true,
      "deletionProtection": false
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
      "endpoint": ":3306"
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
- **Duration:** 887ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-tencent/database`
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
- **Duration:** 1.277s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-tencent/database`
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
- **Duration:** 332ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m29.811s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 7.083s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-tencent/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 1m59.246s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-tencent?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 1.672s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 3.343s

