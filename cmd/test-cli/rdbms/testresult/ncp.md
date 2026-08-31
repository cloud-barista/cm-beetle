# Managed RDBMS Test Report: NCP (kr)

- **Test Case:** NCP Korea MySQL Test
- **Date & Time:** 2026-08-31 16:38:03
- **Namespace:** `default`
- **Total Duration:** 26m24.68s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** NCP
- **Target Region:** `kr`
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
- **Duration:** 18.328s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "23214590",
  "specId": "s2-g3"
}
```
```json
// Response Body
{
  "connectionName": "ncp-jpn",
  "errors": [
    "Image '23214590' not available: image '23214590' (CSP id: 23214590) is registered for region(s) [kr] but connection 'ncp-jpn' targets region 'jpn'; pick an image registered for region 'jpn'"
  ],
  "estimatedCost": "$0.0922/hour",
  "imageId": "23214590",
  "imageValidation": {
    "cspResourceId": "23214590",
    "isAvailable": false,
    "message": "image '23214590' (CSP id: 23214590) is registered for region(s) [kr] but connection 'ncp-jpn' targets region 'jpn'; pick an image registered for region 'jpn'",
    "resourceId": "23214590",
    "status": "Unavailable"
  },
  "isValid": false,
  "message": "Image '23214590' is not available",
  "providerName": "ncp",
  "regionName": "jpn",
  "specDetails": {
    "architecture": "x86_64",
    "connectionName": "ncp-jpn",
    "costPerHour": 0.0922,
    "cspSpecName": "s2-g3",
    "details": [
      {
        "key": "ServerSpecCode",
        "value": "s2-g3"
      },
      {
        "key": "GenerationCode",
        "value": "G3"
      },
      {
        "key": "CpuCount",
        "value": "2"
      },
      {
        "key": "MemorySize",
        "value": "8589934592"
      },
      {
        "key": "HypervisorType",
        "value": "{code:KVM,codeName:KVM}"
      },
      {
        "key": "CpuArchitectureType",
        "value": "{code:X86_64,codeName:x86 64bit}"
      },
      {
        "key": "BlockStorageMaxCount",
        "value": "20"
      },
      {
        "key": "BlockStorageMaxIops",
        "value": "4725"
      },
      {
        "key": "BlockStorageMaxThroughput",
        "value": "84934656"
      },
      {
        "key": "NetworkPerformance",
        "value": "1000000000"
      },
      {
        "key": "NetworkInterfaceMaxCount",
        "value": "3"
      },
      {
        "key": "ServerProductCode",
        "value": "SVR.VSVR.STAND.C002.M008.G003"
      },
      {
        "key": "ServerSpecDescription",
        "value": "vCPU 2EA, Memory 8GB"
      },
      {
        "key": "ServerSpecNo",
        "value": "283"
      },
      {
        "key": "CorrespondingImageIds",
        "value": "142902717,108645046,106703137,104027588,26905053,24075285,22224744,22224742,22224738,22224737"
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
    "id": "ncp+jpn+s2-g3",
    "infraType": "node",
    "memoryGiB": 8,
    "name": "ncp+jpn+s2-g3",
    "namespace": "system",
    "providerName": "ncp",
    "regionLatitude": 35.51742,
    "regionLongitude": 136.80291,
    "regionName": "jpn",
    "rootDiskSize": 0,
    "rootDiskType": "default",
    "systemLabel": "from-assets",
    "uid": "tb0tuig9ihijn201tn1c",
    "vCPU": 2
  },
  "specId": "s2-g3",
  "specValidation": {
    "cspResourceId": "s2-g3",
    "isAvailable": true,
    "resourceId": "s2-g3",
    "resourceName": "s2-g3",
    "status": "Available"
  },
  "status": "Error"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 31.109s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.7.0.0/16",
  "connectionName": "ncp-kr",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-ncp",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.7.1.0/24",
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
  "cidrBlock": "10.7.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-08-31T07:38:53Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:53Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:53Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "146849",
  "cspResourceName": "tb86td01f4ldlspttsm5",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-ncp",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "VpcNo",
      "value": "146849"
    },
    {
      "key": "VpcName",
      "value": "tb86td01f4ldlspttsm5"
    },
    {
      "key": "Ipv4CidrBlock",
      "value": "10.7.0.0/16"
    },
    {
      "key": "VpcStatus",
      "value": "{code:RUN,codeName:운영중}"
    },
    {
      "key": "RegionCode",
      "value": "KR"
    },
    {
      "key": "CreateDate",
      "value": "2026-08-31T16:38:24+0900"
    }
  ],
  "name": "test-rdbms-vnet-ncp",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-08-31T07:38:53Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-08-31T07:38:53Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "320305",
      "cspResourceName": "tb6jnhabbcjnj9tql5md",
      "cspVNetId": "146849",
      "cspVNetName": "tb86td01f4ldlspttsm5",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.7.1.0/24",
      "keyValueList": [
        {
          "key": "SubnetNo",
          "value": "320305"
        },
        {
          "key": "VpcNo",
          "value": "146849"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "SubnetName",
          "value": "tb6jnhabbcjnj9tql5md"
        },
        {
          "key": "Subnet",
          "value": "10.7.1.0/24"
        },
        {
          "key": "SubnetStatus",
          "value": "{code:RUN,codeName:운영중}"
        },
        {
          "key": "CreateDate",
          "value": "2026-08-31T16:38:40+0900"
        },
        {
          "key": "SubnetType",
          "value": "{code:PUBLIC,codeName:Public}"
        },
        {
          "key": "UsageType",
          "value": "{code:GEN,codeName:General}"
        },
        {
          "key": "NetworkAclNo",
          "value": "197697"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb6jnhabbcjnj9tql5md",
      "zone": "KR-1"
    }
  ],
  "systemLabel": "",
  "uid": "tb86td01f4ldlspttsm5"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 26.982s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "ncp-kr",
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
  "name": "test-rdbms-sg-ncp",
  "vNetId": "test-rdbms-vnet-ncp"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "389898",
  "cspResourceName": "tbemqeq2qdhuc9j00rvh",
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
      "Protocol": "ICMP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "outbound",
      "Port": "1-65535",
      "Protocol": "UDP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "outbound",
      "Port": "1-65535",
      "Protocol": "TCP"
    }
  ],
  "id": "test-rdbms-sg-ncp",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "AccessControlGroupNo",
      "value": "389898"
    },
    {
      "key": "AccessControlGroupName",
      "value": "tbemqeq2qdhuc9j00rvh"
    },
    {
      "key": "IsDefault",
      "value": "false"
    },
    {
      "key": "VpcNo",
      "value": "146849"
    },
    {
      "key": "AccessControlGroupStatus",
      "value": "{code:RUN,codeName:운영중}"
    }
  ],
  "name": "test-rdbms-sg-ncp",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbemqeq2qdhuc9j00rvh",
  "vNetId": "test-rdbms-vnet-ncp"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=ncp`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "ncp": {
      "dbOperationMethod": "cspNativeApi",
      "note": "Storage type and storage size configuration not supported. NCP G3 generation automatically applies SSD storage, starts at 10GB, and auto-scales by 10GB increments up to 6000GB.",
      "storageTypeSelectable": false,
      "supported": true,
      "supportedDBEngines": [
        "mysql"
      ]
    }
  }
}
```

### 5. Beetle GET RDBMS Capability [✅ SUCCESS]
- **Duration:** 2.625s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=ncp-kr`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "connectionName": "ncp-kr",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C002.M008.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C004.M008.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C004.M016.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C008.M016.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C002.M016.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C008.M032.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C016.M032.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C004.M032.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C016.M064.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C032.M064.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C008.M064.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C048.M096.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C032.M128.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HICPU.C064.M128.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C016.M128.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C048.M192.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.STAND.C064.M256.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C032.M256.NET.SSD.B050.G003",
      "SVR.VDBAS.AMD.HIMEM.C048.M384.NET.SSD.B050.G003"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "4096",
        "name": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "8192",
        "name": "SVR.VDBAS.AMD.HICPU.C004.M008.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "SVR.VDBAS.AMD.HICPU.C008.M016.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "32768",
        "name": "SVR.VDBAS.AMD.HICPU.C016.M032.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "65536",
        "name": "SVR.VDBAS.AMD.HICPU.C032.M064.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "98304",
        "name": "SVR.VDBAS.AMD.HICPU.C048.M096.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "131072",
        "name": "SVR.VDBAS.AMD.HICPU.C064.M128.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "16384",
        "name": "SVR.VDBAS.AMD.HIMEM.C002.M016.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "SVR.VDBAS.AMD.HIMEM.C004.M032.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "65536",
        "name": "SVR.VDBAS.AMD.HIMEM.C008.M064.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "SVR.VDBAS.AMD.HIMEM.C016.M128.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "SVR.VDBAS.AMD.HIMEM.C032.M256.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "393216",
        "name": "SVR.VDBAS.AMD.HIMEM.C048.M384.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "8192",
        "name": "SVR.VDBAS.AMD.STAND.C002.M008.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "SVR.VDBAS.AMD.STAND.C004.M016.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "SVR.VDBAS.AMD.STAND.C008.M032.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "SVR.VDBAS.AMD.STAND.C016.M064.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "SVR.VDBAS.AMD.STAND.C032.M128.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "196608",
        "name": "SVR.VDBAS.AMD.STAND.C048.M192.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "SVR.VDBAS.AMD.STAND.C064.M256.NET.SSD.B050.G003",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
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
          "description": "NCP G3 automatically applies SSD storage (starts at 10GB and auto-scales by 10GB increments up to 6000GB). Storage size and type are not user-configurable.",
          "displayName": "SSD (automatic)",
          "recommendationLevel": "standard",
          "storageType": "NA"
        }
      ]
    },
    "providerName": "ncp",
    "regionName": "KR",
    "requiresSecurityGroup": false,
    "requiresSubnet": true,
    "storageSizeRange": {
      "max": 6000,
      "min": 10
    },
    "storageTypeOptions": [
      "NA"
    ],
    "supportedVersions": [
      "8.4.8",
      "8.4.6",
      "8.0.45",
      "8.0.42",
      "8.0.40",
      "8.0.36"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": true,
    "supportsEncryption": false,
    "supportsHighAvailability": true,
    "supportsPublicAccess": false,
    "supportsStorageSizeConfiguration": false,
    "supportsStorageTypeSelection": false,
    "supportsTag": false
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
    "csp": "ncp",
    "region": "kr"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for ncp (kr)",
  "warnings": [
    "NCP Cloud DB does not provide external public IP by default; instance 'source-mysql-01' will be created within private VPC."
  ],
  "targetCloud": {
    "csp": "ncp",
    "region": "kr"
  },
  "targetRDBMSInstances": [
    {
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "rdbmsName": "test-rdbms-ncp",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.45",
      "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
      "storageSize": 100,
      "adminUserName": "dbadmin",
      "adminUserPassword": "******",
      "vNetId": "test-rdbms-vnet-ncp",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-ncp"
      ],
      "publicAccess": false,
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
- **Duration:** 13ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "dbadmin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "ncp-kr",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0.45",
  "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
  "name": "test-rdbms-ncp",
  "securityGroupIds": [
    "test-rdbms-sg-ncp"
  ],
  "storageSize": 100,
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-ncp"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "dbadmin",
    "adminUserPassword": "******",
    "connectionName": "ncp-kr",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0.45",
    "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
    "name": "test-rdbms-ncp",
    "securityGroupIds": [
      "test-rdbms-sg-ncp"
    ],
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-ncp"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 11m36.202s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for ncp (kr)",
  "status": "recommended",
  "targetCloud": {
    "csp": "ncp",
    "region": "kr"
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
      "dbEngineVersion": "8.0.45",
      "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
      "highAvailability": false,
      "publicAccess": false,
      "rdbmsName": "test-rdbms-ncp",
      "securityGroupIds": [
        "test-rdbms-sg-ncp"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-ncp"
    }
  ],
  "warnings": [
    "NCP Cloud DB does not provide external public IP by default; instance 'source-mysql-01' will be created within private VPC."
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
- **Duration:** 1.318s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-ncp`
```json
// Response Body
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
  "status": "Available",
  "conditions": [
    {
      "type": "Ready",
      "status": "True",
      "reason": "Available",
      "lastTransitionTime": "2026-08-31T07:50:49Z"
    },
    {
      "type": "Synced",
      "status": "True",
      "reason": "Available",
      "lastTransitionTime": "2026-08-31T07:50:49Z"
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
  "adminUserName": "dbadmin",
  "highAvailability": false,
  "backupRetentionDays": 7,
  "backupTime": "02:30",
  "publicAccess": false,
  "deletionProtection": false,
  "endpoint": "db-4a3o7k.vpc-cdb.ntruss.com:3306"
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
      "status": "Deleting",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Deleting",
          "message": "RDBMS deletion in progress",
          "lastTransitionTime": "2026-08-31T07:49:52Z"
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
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:49:43Z"
        },
        {
          "type": "Synced",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:49:43Z"
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
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:50:49Z"
        },
        {
          "type": "Synced",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:50:49Z"
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
      "adminUserName": "dbadmin",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "backupTime": "02:30",
      "publicAccess": false,
      "deletionProtection": false,
      "endpoint": "db-4a3o7k.vpc-cdb.ntruss.com:3306"
    },
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-nhn",
      "uid": "tbna7alivb64t5mbface",
      "cspResourceName": "tbna7alivb64t5mbface",
      "cspResourceId": "291e605f-dde1-44a8-bf27-e32d56dfdb6b",
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
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:47:35Z"
        },
        {
          "type": "Synced",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:47:35Z"
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
      "dbInstanceType": "NA",
      "storageType": "General SSD",
      "storageSize": 100,
      "adminUserName": "myadmin",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "backupTime": "03:00",
      "publicAccess": true,
      "nhnDBSGToAllowAllInbound": true,
      "deletionProtection": false,
      "endpoint": "8c0353cf-3a53-416d-b08f-129934927e69.external.kr1.mysql.rds.nhncloudservice.com:3306"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 7.954s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-ncp/database`
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
- **Duration:** 4.191s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-ncp/database`
```json
// Response Body
{
  "databases": [
    "mydb",
    "sampledb",
    "sampledb_dyn"
  ]
}
```

### 13. Data I/O Test (External Remote) [⚪ SKIPPED (Skipped: Public Access is not supported/configured for this CSP (VPC private endpoint only))]
- **Duration:** 0s

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 7m59.643s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 26.4s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-ncp/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 4m15.47s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-ncp?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 7.275s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 27.158s

