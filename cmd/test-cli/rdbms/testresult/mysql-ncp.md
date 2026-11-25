# Managed RDBMS (MySQL) Test Report: NCP (kr)

- **Test Case:** NCP Korea MySQL Test
- **Access Mode:** Private (publicAccess=false)
- **Date & Time:** 2026-11-25 18:54:32
- **Namespace:** `default`
- **Total Duration:** 23m52.465s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** NCP
- **Target Region:** `kr`
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
- **Duration:** 2.274s
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
  "availability": {
    "available": true,
    "instanceType": "s2-g3",
    "provider": "ncp",
    "queriedAt": "2026-11-25T09:54:34.582147085Z",
    "refreshAfter": "2026-11-25T10:54:34.582147085Z",
    "region": "jpn",
    "source": "ncp:GetServerSpecList",
    "status": "Available",
    "ttlSeconds": 3600
  },
  "connectionName": "ncp-jpn",
  "errors": [
    "Image '23214590' not available: image '23214590' (CSP id: 23214590) is registered for region(s) [kr] but connection 'ncp-jpn' targets region 'jpn'; pick an image registered for region 'jpn'"
  ],
  "estimatedCost": "$0.0920/hour",
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
    "availabilityStatus": "Unverified",
    "connectionName": "ncp-jpn",
    "costPerHour": 0.092,
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
        "value": "145005335,142902717,108645046,106703137,104027588,24075285,22224744,22224742,22224738,22224737"
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
    "uid": "tbboq5torqfds0j60qtq",
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
- **Duration:** 24.795s
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
      "lastTransitionTime": "2026-11-25T09:54:57Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:57Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:57Z",
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
  "cspResourceId": "153793",
  "cspResourceName": "tbduitifgo74onfebkgm",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-ncp",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "VpcNo",
      "value": "153793"
    },
    {
      "key": "VpcName",
      "value": "tbduitifgo74onfebkgm"
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
      "value": "2026-11-25T18:54:33+0900"
    }
  ],
  "name": "test-rdbms-vnet-ncp",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:54:57Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:54:57Z",
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
      "cspResourceId": "336715",
      "cspResourceName": "tbvjpsapkq6l5itoovl8",
      "cspVNetId": "153793",
      "cspVNetName": "tbduitifgo74onfebkgm",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.7.1.0/24",
      "keyValueList": [
        {
          "key": "SubnetNo",
          "value": "336715"
        },
        {
          "key": "VpcNo",
          "value": "153793"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "SubnetName",
          "value": "tbvjpsapkq6l5itoovl8"
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
          "value": "2026-11-25T18:54:44+0900"
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
          "value": "207172"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbvjpsapkq6l5itoovl8",
      "zone": "KR-1"
    }
  ],
  "systemLabel": "",
  "uid": "tbduitifgo74onfebkgm"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 16.586s
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
  "cspResourceId": "423545",
  "cspResourceName": "tbm14odffhkjcjt61ror",
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
      "value": "423545"
    },
    {
      "key": "AccessControlGroupName",
      "value": "tbm14odffhkjcjt61ror"
    },
    {
      "key": "IsDefault",
      "value": "false"
    },
    {
      "key": "VpcNo",
      "value": "153793"
    },
    {
      "key": "AccessControlGroupStatus",
      "value": "{code:RUN,codeName:운영중}"
    }
  ],
  "name": "test-rdbms-sg-ncp",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbm14odffhkjcjt61ror",
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
      "note": "PublicAccess=true is officially not supported because NCP Cloud DB provides only private VPC endpoints (*.vpc-cdb.ntruss.com) and lacks an Open API/SDK for public domain allocation (enforced in CB-Spider v0.13.16 #1879). Storage type and storage size configuration not supported. NCP G3 generation automatically applies SSD storage, starts at 10GB, and auto-scales by 10GB increments up to 6000GB. NCP Cloud DB ACGs apply a default-deny policy (0 inbound rules), blocking port 3306 even for internal VPC VMs until an ACG inbound permit rule is added. For test/dev, setting ncpDBACGToAllowAllInbound=true (NCPAutoOpenACG=true in CB-Spider v0.13.16+) auto-configures 0.0.0.0/0 on the ACG (TEST ONLY; do not use in production).",
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
- **Duration:** 898ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=ncp-kr&dbEngine=mysql`
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
          "constraints": "Minimum 10GB storage.",
          "description": "NCP G3 automatically applies SSD storage (starts at 10GB and auto-scales by 10GB increments up to 6000GB). Storage size and type are not user-configurable.",
          "displayName": "SSD (automatic)",
          "maxSizeGB": 6000,
          "minSizeGB": 10,
          "recommendationLevel": "standard",
          "storageType": "NA"
        }
      ]
    },
    "providerName": "ncp",
    "regionName": "KR",
    "requiresSecurityGroup": false,
    "requiresSubnet": true,
    "storageSizeRangeGB": {
      "max": 6000,
      "min": 10
    },
    "storageTypeOptions": [
      "NA"
    ],
    "supportedVersions": [
      "8.4.11",
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
- **Duration:** 6ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "ncp",
    "region": "kr"
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
    "ncpDBACGToAllowAllInbound": true,
    "publicAccess": false
  }
}
```
```json
// Response Body
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
      "ncpDBACGToAllowAllInbound": true,
      "publicAccess": false,
      "rdbmsName": "rdbms-ncp",
      "securityGroupIds": [
        "test-rdbms-sg-ncp"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-ncp"
    }
  ],
  "warnings": [
    "Storage type selection is not configurable on target cloud (ncp); requested storage type 'SSD' for instance 'Source MySQL 01' will be managed automatically by the provider.",
    "instance 'Source MySQL 01' database 'sampledb': omitted charset/collation on NCP so database creation uses the native NCP REST API without requiring direct VPC SQL connectivity"
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 44ms
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
  "name": "rdbms-ncp",
  "ncpDBACGToAllowAllInbound": true,
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
    "name": "rdbms-ncp",
    "ncpDBACGToAllowAllInbound": true,
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
- **Duration:** 11m48.855s
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
      "ncpDBACGToAllowAllInbound": true,
      "publicAccess": false,
      "rdbmsName": "rdbms-ncp",
      "securityGroupIds": [
        "test-rdbms-sg-ncp"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-ncp"
    }
  ],
  "warnings": [
    "Storage type selection is not configurable on target cloud (ncp); requested storage type 'SSD' for instance 'Source MySQL 01' will be managed automatically by the provider.",
    "instance 'Source MySQL 01' database 'sampledb': omitted charset/collation on NCP so database creation uses the native NCP REST API without requiring direct VPC SQL connectivity"
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
- **Duration:** 331ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "07:45",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T10:06:12Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T10:06:12Z",
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
  "cspResourceId": "147164850",
  "cspResourceName": "tbpth30tovigfplksoib",
  "dbEngine": "mysql",
  "dbEngineVersion": "MYSQL8.0.45",
  "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
  "dbInstanceType": "Stand Alone",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "db-4cb3lv.vpc-cdb.ntruss.com:3306",
  "highAvailability": false,
  "id": "test-rdbms-ncp",
  "keyValueList": [
    {
      "key": "GenerationCode",
      "value": "G3"
    },
    {
      "key": "ImageProductCode",
      "value": "SW.VMYSL.OS.LNX64.ROCKY.08.MYSQL.B050.G003"
    },
    {
      "key": "IsMultiZone",
      "value": "false"
    },
    {
      "key": "License",
      "value": "라이센스코드-CLOUD DATABASE(VPC)"
    },
    {
      "key": "AccessControlGroupNos",
      "value": "423546"
    },
    {
      "key": "ServerName",
      "value": "tbpth30tovigfpl-001-a40j"
    },
    {
      "key": "ServerInstanceNo",
      "value": "147164863"
    },
    {
      "key": "ZoneCode",
      "value": "KR-1"
    },
    {
      "key": "PrivateIp",
      "value": "10.7.1.6"
    },
    {
      "key": "CpuCount",
      "value": "2"
    },
    {
      "key": "MemorySizeGB",
      "value": "4"
    },
    {
      "key": "UsedStorageSizeGB",
      "value": "0"
    },
    {
      "key": "Uptime",
      "value": "2026-11-25T18:57:56+0900"
    }
  ],
  "name": "test-rdbms-ncp",
  "ncpDBACGToAllowAllInbound": true,
  "publicAccess": false,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-ncp-acg"
  ],
  "status": "Available",
  "storageSize": 10,
  "storageType": "SSD",
  "subnetIds": [
    "subnet-1"
  ],
  "uid": "tbpth30tovigfplksoib",
  "vNetId": "test-rdbms-vnet-ncp"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 20ms
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
          "lastTransitionTime": "2026-11-25T10:04:29Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
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
      "status": "Deleting",
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
          "lastTransitionTime": "2026-11-25T10:04:51Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:04:51Z",
          "reason": "Available",
          "status": "True",
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
          "value": "available"
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
          "key": "LatestRestorableTime",
          "value": "2026-11-25T10:00:09.563Z"
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
      "status": "Available",
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
      "backupRetentionDays": 30,
      "backupTime": "AUTO",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:05:39Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:05:39Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "2fe2e2f8-4d98-44a7-be74-b2092b2ff2fe",
      "cspResourceName": "tbun4h3ej4rgndeesotl",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "multitenant",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "encryption": true,
      "endpoint": "2fe2e2f8-4d98-44a7-be74-b2092b2ff2fe.c7dvrhud08vgdqo60090.databases.appdomain.cloud:32473",
      "highAvailability": false,
      "id": "test-rdbms-ibm",
      "keyValueList": [
        {
          "key": "ID",
          "value": "crn:v1:bluemix:public:databases-for-mysql:us-south:a/ab205347a7c3b57f09dabb32df178bcf:2fe2e2f8-4d98-44a7-be74-b2092b2ff2fe::"
        },
        {
          "key": "GUID",
          "value": "2fe2e2f8-4d98-44a7-be74-b2092b2ff2fe"
        },
        {
          "key": "URL",
          "value": "/v2/resource_instances/2fe2e2f8-4d98-44a7-be74-b2092b2ff2fe"
        },
        {
          "key": "CreatedAt",
          "value": "2026-11-25T09:55:15.299Z"
        },
        {
          "key": "UpdatedAt",
          "value": "2026-11-25T10:05:17.058Z"
        },
        {
          "key": "CreatedBy",
          "value": "IBMid-692000RYMF"
        },
        {
          "key": "Name",
          "value": "tbun4h3ej4rgndeesotl"
        },
        {
          "key": "RegionID",
          "value": "us-south"
        },
        {
          "key": "AccountID",
          "value": "ab205347a7c3b57f09dabb32df178bcf"
        },
        {
          "key": "ResourcePlanID",
          "value": "databases-for-mysql-standard"
        },
        {
          "key": "ResourceGroupID",
          "value": "e7c20a4f7ee64603b1c06d46b0c2385c"
        },
        {
          "key": "ResourceGroupCRN",
          "value": "crn:v1:bluemix:public:resource-controller::a/ab205347a7c3b57f09dabb32df178bcf::resource-group:e7c20a4f7ee64603b1c06d46b0c2385c"
        },
        {
          "key": "TargetCRN",
          "value": "crn:v1:bluemix:public:globalcatalog::::deployment:databases-for-mysql-standard%3Aus-south"
        },
        {
          "key": "OnetimeCredentials",
          "value": "false"
        },
        {
          "key": "Parameters",
          "value": "{members_disk_allocation_mb:307200,members_host_flavor:multitenant,service-endpoints:public,version:8.4}"
        },
        {
          "key": "AllowCleanup",
          "value": "false"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:databases-for-mysql:us-south:a/ab205347a7c3b57f09dabb32df178bcf:2fe2e2f8-4d98-44a7-be74-b2092b2ff2fe::"
        },
        {
          "key": "State",
          "value": "active"
        },
        {
          "key": "Type",
          "value": "service_instance"
        },
        {
          "key": "SubType",
          "value": "Public"
        },
        {
          "key": "ResourceID",
          "value": "databases-for-mysql"
        },
        {
          "key": "DashboardURL",
          "value": "https://cloud.ibm.com/?instance_id=crn%3Av1%3Abluemix%3Apublic%3Adatabases-for-mysql%3Aus-south%3Aa%2Fab205347a7c3b57f09dabb32df178bcf%3A2fe2e2f8-4d98-44a7-be74-b2092b2ff2fe%3A%3A\\u0026platform=ibm"
        },
        {
          "key": "LastOperation",
          "value": "{async:true,cancelable:false,description:Provisioning mysql with version 8.4 (100%),poll:true,poll_after:30,state:succeeded,type:create}"
        },
        {
          "key": "ResourceKeysURL",
          "value": "/v2/resource_instances/2fe2e2f8-4d98-44a7-be74-b2092b2ff2fe/resource_keys"
        },
        {
          "key": "PlanHistory",
          "value": "{resource_plan_id:databases-for-mysql-standard,start_date:2026-11-25T09:55:15.299Z,requestor_id:IBMid-692000RYMF}"
        },
        {
          "key": "Migrated",
          "value": "false"
        },
        {
          "key": "Locked",
          "value": "false"
        }
      ],
      "name": "test-rdbms-ibm",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "NA",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbun4h3ej4rgndeesotl",
      "vNetId": "test-rdbms-vnet-ibm"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "07:45",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:06:12Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:06:12Z",
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
      "cspResourceId": "147164850",
      "cspResourceName": "tbpth30tovigfplksoib",
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL8.0.45",
      "dbInstanceSpec": "SVR.VDBAS.AMD.HICPU.C002.M004.NET.SSD.B050.G003",
      "dbInstanceType": "Stand Alone",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "db-4cb3lv.vpc-cdb.ntruss.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-ncp",
      "keyValueList": [
        {
          "key": "GenerationCode",
          "value": "G3"
        },
        {
          "key": "ImageProductCode",
          "value": "SW.VMYSL.OS.LNX64.ROCKY.08.MYSQL.B050.G003"
        },
        {
          "key": "IsMultiZone",
          "value": "false"
        },
        {
          "key": "License",
          "value": "라이센스코드-CLOUD DATABASE(VPC)"
        },
        {
          "key": "AccessControlGroupNos",
          "value": "423546"
        },
        {
          "key": "ServerName",
          "value": "tbpth30tovigfpl-001-a40j"
        },
        {
          "key": "ServerInstanceNo",
          "value": "147164863"
        },
        {
          "key": "ZoneCode",
          "value": "KR-1"
        },
        {
          "key": "PrivateIp",
          "value": "10.7.1.6"
        },
        {
          "key": "CpuCount",
          "value": "2"
        },
        {
          "key": "MemorySizeGB",
          "value": "4"
        },
        {
          "key": "UsedStorageSizeGB",
          "value": "0"
        },
        {
          "key": "Uptime",
          "value": "2026-11-25T18:57:56+0900"
        }
      ],
      "name": "test-rdbms-ncp",
      "ncpDBACGToAllowAllInbound": true,
      "publicAccess": false,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-ncp-acg"
      ],
      "status": "Available",
      "storageSize": 10,
      "storageType": "SSD",
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
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 7.454s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp/database`
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
- **Duration:** 2.092s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp/database`
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

### 13. Beetle GET Secure Transport Info [⚪ SKIPPED (Expected HTTP 422 fast-fail on private RDBMS (inspect TLS inside VPC))]
- **Duration:** 22ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp/secure-transport`

### 14. Data I/O Test (External Remote) [⚪ SKIPPED (Skipped: Public Access is not supported/configured for this CSP (VPC private endpoint only))]
- **Duration:** 0s

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 5m56.158s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 18.677s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 4m22.238s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ncp?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 1.122s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 30.889s

