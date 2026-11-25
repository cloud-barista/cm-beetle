# Managed RDBMS (MySQL) Test Report: NHN (kr1)

- **Test Case:** NHN KR1 MySQL Test
- **Access Mode:** Public (publicAccess=true)
- **Date & Time:** 2026-11-25 18:54:32
- **Namespace:** `default`
- **Total Duration:** 24m39.255s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** NHN
- **Target Region:** `kr1`
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
- **Duration:** 2.283s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "5c38715f-0375-4167-af4e-56f75ba8b252",
  "specId": "nhn+kr1+m2.c2m4"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "cached": true,
    "instanceType": "m2.c2m4",
    "provider": "nhn",
    "queriedAt": "2026-11-25T09:17:37.696737458Z",
    "refreshAfter": "2026-11-25T10:17:37.696737458Z",
    "region": "kr1",
    "source": "nhn:flavors+limits",
    "status": "Available",
    "ttlSeconds": 3600
  },
  "connectionName": "nhn-kr1",
  "estimatedCost": "$0.0688/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "nhn-kr1",
    "creationDate": "",
    "cspImageName": "5c38715f-0375-4167-af4e-56f75ba8b252",
    "description": "",
    "details": [
      {
        "key": "ID",
        "value": "5c38715f-0375-4167-af4e-56f75ba8b252"
      },
      {
        "key": "Name",
        "value": "Ubuntu Server 24.04.3 LTS (2026.03.10)"
      },
      {
        "key": "Status",
        "value": "active"
      },
      {
        "key": "Tags",
        "value": "BASE"
      },
      {
        "key": "ContainerFormat",
        "value": "bare"
      },
      {
        "key": "DiskFormat",
        "value": "qcow2"
      },
      {
        "key": "MinDiskGigabytes",
        "value": "20"
      },
      {
        "key": "MinRAMMegabytes",
        "value": "0"
      },
      {
        "key": "Owner",
        "value": "c289b99209ca4e189095cdecebbd092d"
      },
      {
        "key": "Protected",
        "value": "true"
      },
      {
        "key": "Visibility",
        "value": "public"
      },
      {
        "key": "Hidden",
        "value": "false"
      },
      {
        "key": "Checksum",
        "value": "034b1571183de81ad9cc385338181184"
      },
      {
        "key": "Properties",
        "value": "{deprecate_date:null,description:Ubuntu Server 24.04.3 LTS (2026.03.10),hw_cpu_sockets:1,hw_qemu_guest_agent:yes,hw_vif_multiqueue_enabled:true,login_username:ubuntu,max_cpu:,min_cpu:0,monitoring_agent:sysmon,nhncloud_allow_autoscale:true,nhncloud_allow_cgroup:v2,nhncloud_allow_compute_flavor:true,nhncloud_allow_cpu_flavor:true,nhncloud_allow_download:false,nhncloud_allow_gpu_flavor:false,nhncloud_allow_image_create:true,nhncloud_allow_imagebuilder:true,nhncloud_allow_instance_template:true,nhncloud_allow_local_bootdisk_flavor:true,nhncloud_allow_nks_cpu_flavor:false,nhncloud_allow_nks_gpu_flavor:false,nhncloud_allow_user_script:true,nhncloud_category:OS,nhncloud_product:compute,os_architecture:amd64,os_distro:ubuntu,os_type:linux,os_version:Server 24.04 LTS,project_domain:WDI;NORMAL,release_date:2026.03.10,tc_env:cloudmon,sysmon}"
      },
      {
        "key": "CreatedAt",
        "value": "2026-03-09T01:34:35Z"
      },
      {
        "key": "UpdatedAt",
        "value": "2026-03-09T21:14:17Z"
      },
      {
        "key": "File",
        "value": "/v2/images/5c38715f-0375-4167-af4e-56f75ba8b252/file"
      },
      {
        "key": "Schema",
        "value": "/v2/schemas/image"
      },
      {
        "key": "VirtualSize",
        "value": "0"
      }
    ],
    "fetchedTime": "2026.10.29 07:15:53 Thu",
    "id": "5c38715f-0375-4167-af4e-56f75ba8b252",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "5c38715f-0375-4167-af4e-56f75ba8b252",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 20,
    "osDiskType": "NA",
    "osDistribution": "Ubuntu Server 24.04.3 LTS (2026.03.10)",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "nhn",
    "regionList": [
      "kr1"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tb7nvga93fgrh1btgd3e"
  },
  "imageId": "5c38715f-0375-4167-af4e-56f75ba8b252",
  "imageValidation": {
    "cspResourceId": "5c38715f-0375-4167-af4e-56f75ba8b252",
    "isAvailable": true,
    "resourceId": "5c38715f-0375-4167-af4e-56f75ba8b252",
    "resourceName": "5c38715f-0375-4167-af4e-56f75ba8b252",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid but has warnings",
  "providerName": "nhn",
  "regionName": "kr1",
  "specDetails": {
    "architecture": "x86_64",
    "availabilityCheckedAt": "2026-11-25T09:17:37.696737458Z",
    "availabilityRefreshAfter": "2026-11-25T10:17:37.696737458Z",
    "availabilityStatus": "Available",
    "connectionName": "nhn-kr1",
    "costPerHour": 0.06882,
    "cspSpecName": "m2.c2m4",
    "details": [
      {
        "key": "ID",
        "value": "35a73b57-58a7-434d-aa08-5249aaa95b3e"
      },
      {
        "key": "Name",
        "value": "m2.c2m4"
      },
      {
        "key": "Links",
        "value": "{href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/v2.1/flavors/35a73b57-58a7-434d-aa08-5249aaa95b3e,rel:self}; {href:http://nova.iaas.tcc1.cloud.toastoven.net:8774/flavors/35a73b57-58a7-434d-aa08-5249aaa95b3e,rel:bookmark}"
      },
      {
        "key": "RAM",
        "value": "4096"
      },
      {
        "key": "Disabled",
        "value": "false"
      },
      {
        "key": "VCPUs",
        "value": "2"
      },
      {
        "key": "ExtraSpecs",
        "value": "{flavor_type:general}"
      },
      {
        "key": "IsPublic",
        "value": "true"
      },
      {
        "key": "RxTxFactor",
        "value": "1.00"
      },
      {
        "key": "Ephemeral",
        "value": "0"
      },
      {
        "key": "Disk",
        "value": "0"
      },
      {
        "key": "Notice!!",
        "value": "Specify 'RootDiskType' and 'RootDiskSize' when VM Creation to Boot from the Attached Volume!!"
      }
    ],
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
    "id": "nhn+kr1+m2.c2m4",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "nhn+kr1+m2.c2m4",
    "namespace": "system",
    "providerName": "nhn",
    "regionLatitude": 37.390889,
    "regionLongitude": 127.096792,
    "regionName": "kr1",
    "rootDiskSize": 0,
    "rootDiskType": "default",
    "systemLabel": "from-assets",
    "uid": "tbjd179s4cm2g5ij21tt",
    "vCPU": 2
  },
  "specId": "nhn+kr1+m2.c2m4",
  "specValidation": {
    "cspResourceId": "m2.c2m4",
    "isAvailable": true,
    "resourceId": "nhn+kr1+m2.c2m4",
    "resourceName": "m2.c2m4",
    "status": "Available"
  },
  "status": "Warning",
  "warnings": [
    "Moderate provisioning failure risk: Image '5c38715f-0375-4167-af4e-56f75ba8b252' with spec 'nhn+kr1+m2.c2m4' has succeeded before, but also had past failures. Recent failures: VM provisioning failed (status: Terminating) - check CSP console for details"
  ]
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 33.948s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.8.0.0/16",
  "connectionName": "nhn-kr1",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-nhn",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.8.1.0/24",
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
  "cidrBlock": "10.8.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T09:55:05Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T09:55:05Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-11-25T09:55:05Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "5f5b3c87-f70d-4928-be10-bb7a5df0293f",
  "cspResourceName": "tb7bsim34621n59sijmu",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-nhn",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "Status",
      "value": "available"
    },
    {
      "key": "RouterExternal",
      "value": "No"
    },
    {
      "key": "CreatedTime",
      "value": "2026-11-25 09:54:33"
    }
  ],
  "name": "test-rdbms-vnet-nhn",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:55:05Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:55:05Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "200d6e94-75e0-4809-ac92-bb11535fcfe1",
      "cspResourceName": "tb5h7hmm0smgjga023rk",
      "cspVNetId": "5f5b3c87-f70d-4928-be10-bb7a5df0293f",
      "cspVNetName": "tb7bsim34621n59sijmu",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.8.1.0/24",
      "keyValueList": [
        {
          "key": "RouterExternal",
          "value": "false"
        },
        {
          "key": "Name",
          "value": "tb5h7hmm0smgjga023rk"
        },
        {
          "key": "TenantID",
          "value": "6fc6c2ef568f45f7a9bdc6be211f52f9"
        },
        {
          "key": "State",
          "value": "available"
        },
        {
          "key": "ID",
          "value": "200d6e94-75e0-4809-ac92-bb11535fcfe1"
        },
        {
          "key": "RoutingTable",
          "value": "{gateway_id:,default_table:false,explicit:false,id:,name:}"
        },
        {
          "key": "CreateTime",
          "value": "2026-11-25 09:54:58"
        },
        {
          "key": "AvailableIPCount",
          "value": "0"
        },
        {
          "key": "VPC",
          "value": "{shared:false,state:,id:,cidrv4:,name:}"
        },
        {
          "key": "VPCID",
          "value": "5f5b3c87-f70d-4928-be10-bb7a5df0293f"
        },
        {
          "key": "Shared",
          "value": "false"
        },
        {
          "key": "CIDR",
          "value": "10.8.1.0/24"
        },
        {
          "key": "Gateway",
          "value": "10.8.1.1"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb5h7hmm0smgjga023rk"
    }
  ],
  "systemLabel": "",
  "uid": "tb7bsim34621n59sijmu"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 6.053s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "nhn-kr1",
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
  "name": "test-rdbms-sg-nhn",
  "vNetId": "test-rdbms-vnet-nhn"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "c38510e4-ea35-4422-82fd-da5ef07315ad",
  "cspResourceName": "tbmvjok7micdf568m4oa",
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
  "id": "test-rdbms-sg-nhn",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "Name",
      "value": "tbmvjok7micdf568m4oa"
    },
    {
      "key": "Description",
      "value": "tbmvjok7micdf568m4oa"
    },
    {
      "key": "Rules",
      "value": "{from_port:22,to_port:22,ip_protocol:tcp,ip_range:{CIDR:0.0.0.0/0},Group:{tenant_id:,Name:}}; {from_port:3306,to_port:3306,ip_protocol:tcp,ip_range:{CIDR:0.0.0.0/0},Group:{tenant_id:,Name:}}"
    },
    {
      "key": "TenantID",
      "value": "6fc6c2ef568f45f7a9bdc6be211f52f9"
    }
  ],
  "name": "test-rdbms-sg-nhn",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbmvjok7micdf568m4oa",
  "vNetId": "test-rdbms-vnet-nhn"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 6ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=nhn`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "nhn": {
      "dbOperationMethod": "cspNativeApi",
      "storageTypeSelectable": true,
      "supported": true,
      "supportedDBEngines": [
        "mysql",
        "mariadb"
      ]
    }
  }
}
```

### 5. Beetle GET RDBMS Capability [✅ SUCCESS]
- **Duration:** 2.405s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=nhn-kr1&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "1-730",
    "connectionName": "nhn-kr1",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "m2.c1m2",
      "m2.c2m4",
      "m2.c4m8",
      "m2.c8m16",
      "m2.c16m32",
      "c2.c2m2",
      "c2.c4m4",
      "c2.c8m8",
      "c2.c16m16",
      "r2.c2m8",
      "r2.c4m16",
      "r2.c8m32",
      "r2.c8m64",
      "x1.c16m64",
      "x1.c16m128",
      "x1.c32m128",
      "x1.c32m256"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "16384",
        "name": "c2.c16m16",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "2048",
        "name": "c2.c2m2",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "4096",
        "name": "c2.c4m4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "8192",
        "name": "c2.c8m8",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "32768",
        "name": "m2.c16m32",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "2048",
        "name": "m2.c1m2",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "4096",
        "name": "m2.c2m4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "8192",
        "name": "m2.c4m8",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "m2.c8m16",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "8192",
        "name": "r2.c2m8",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "r2.c4m16",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "r2.c8m32",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "r2.c8m64",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "x1.c16m128",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "65536",
        "name": "x1.c16m64",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "x1.c32m128",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "262144",
        "name": "x1.c32m256",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
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
          "description": "Standard SSD storage for better performance. Recommended for most workloads. Minimum 20 GB.",
          "displayName": "General SSD",
          "maxSizeGB": 2048,
          "minSizeGB": 20,
          "recommendationLevel": "recommended",
          "recommended": true,
          "storageType": "General SSD"
        },
        {
          "constraints": "Minimum 20GB storage.",
          "description": "Cost-effective HDD storage for less demanding workloads. Minimum 20 GB.",
          "displayName": "General HDD",
          "maxSizeGB": 2048,
          "minSizeGB": 20,
          "recommendationLevel": "standard",
          "storageType": "General HDD"
        }
      ]
    },
    "providerName": "nhn",
    "regionName": "KR1",
    "requiresSecurityGroup": false,
    "requiresSubnet": true,
    "storageSizeRangeGB": {
      "max": 2048,
      "min": 20
    },
    "storageTypeOptions": [
      "General SSD",
      "General HDD"
    ],
    "supportedVersions": [
      "MYSQL_V8413",
      "MYSQL_V8411",
      "MYSQL_V8409",
      "MYSQL_V8408",
      "MYSQL_V8407",
      "MYSQL_V8406",
      "MYSQL_V8405",
      "MYSQL_V8046",
      "MYSQL_V8045",
      "MYSQL_V8044",
      "MYSQL_V8043",
      "MYSQL_V8042",
      "MYSQL_V8041",
      "MYSQL_V8040",
      "MYSQL_V8036",
      "MYSQL_V8035",
      "MYSQL_V8034",
      "MYSQL_V8033",
      "MYSQL_V8032",
      "MYSQL_V8028",
      "MYSQL_V8023",
      "MYSQL_V8018",
      "MYSQL_V5737",
      "MYSQL_V5733",
      "MYSQL_V5726",
      "MYSQL_V5719",
      "MYSQL_V5715"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": true,
    "supportsEncryption": false,
    "supportsHighAvailability": true,
    "supportsPublicAccess": true,
    "supportsStorageSizeConfiguration": true,
    "supportsStorageTypeSelection": true,
    "supportsTag": false
  }
}
```

### 6. Beetle POST Recommend RDBMS [✅ SUCCESS]
- **Duration:** 114ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "nhn",
    "region": "kr1"
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
    "adminUserName": "myadmin",
    "backupRetentionDays": 7,
    "highAvailability": false,
    "nhnDBSGToAllowAllInbound": true,
    "publicAccess": true
  }
}
```
```json
// Response Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for nhn (kr1)",
  "status": "recommended",
  "targetCloud": {
    "csp": "nhn",
    "region": "kr1"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "myadmin",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL_V8046",
      "dbInstanceSpec": "m2.c2m4",
      "highAvailability": false,
      "nhnDBSGToAllowAllInbound": true,
      "publicAccess": true,
      "rdbmsName": "rdbms-nhn",
      "securityGroupIds": [
        "test-rdbms-sg-nhn"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "General SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-nhn"
    }
  ],
  "warnings": [
    "instance 'Source MySQL 01' database 'sampledb': omitted charset/collation on NHN Cloud so database creation uses the native NHN REST API without requiring console 'Direct Control' pre-activation"
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 81ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "myadmin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "nhn-kr1",
  "dbEngine": "mysql",
  "dbEngineVersion": "MYSQL_V8046",
  "dbInstanceSpec": "m2.c2m4",
  "name": "rdbms-nhn",
  "nhnDBSGToAllowAllInbound": true,
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-nhn"
  ],
  "storageSize": 100,
  "storageType": "General SSD",
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-nhn"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "myadmin",
    "adminUserPassword": "******",
    "connectionName": "nhn-kr1",
    "dbEngine": "mysql",
    "dbEngineVersion": "MYSQL_V8046",
    "dbInstanceSpec": "m2.c2m4",
    "name": "rdbms-nhn",
    "nhnDBSGToAllowAllInbound": true,
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-nhn"
    ],
    "storageSize": 100,
    "storageType": "General SSD",
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-nhn"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 13m5.171s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for nhn (kr1)",
  "status": "recommended",
  "targetCloud": {
    "csp": "nhn",
    "region": "kr1"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "myadmin",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL_V8046",
      "dbInstanceSpec": "m2.c2m4",
      "highAvailability": false,
      "nhnDBSGToAllowAllInbound": true,
      "publicAccess": true,
      "rdbmsName": "rdbms-nhn",
      "securityGroupIds": [
        "test-rdbms-sg-nhn"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "General SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-nhn"
    }
  ],
  "warnings": [
    "instance 'Source MySQL 01' database 'sampledb': omitted charset/collation on NHN Cloud so database creation uses the native NHN REST API without requiring console 'Direct Control' pre-activation"
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
- **Duration:** 1.687s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-nhn`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "03:00",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T10:07:21Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T10:07:21Z",
      "reason": "Available",
      "status": "True",
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
  "cspResourceId": "14d502e5-fa5d-4412-8d93-d1ff9f7f71ad",
  "cspResourceName": "tbsn6ljiu7npo3re8iqr",
  "dbEngine": "mysql",
  "dbEngineVersion": "MYSQL_V8046",
  "dbInstanceSpec": "m2.c2m4",
  "dbInstanceType": "NA",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "9d20cc3f-d6dd-4b6f-a1a4-1ea553fd42ae.external.kr1.mysql.rds.nhncloudservice.com:3306",
  "highAvailability": false,
  "id": "test-rdbms-nhn",
  "keyValueList": [
    {
      "key": "DBInstanceStatus",
      "value": "AVAILABLE"
    },
    {
      "key": "DBFlavorId",
      "value": "71f69bf9-3c01-4c1a-b135-bb75e93f6268"
    },
    {
      "key": "SubnetId",
      "value": "200d6e94-75e0-4809-ac92-bb11535fcfe1"
    }
  ],
  "name": "test-rdbms-nhn",
  "nhnDBSGToAllowAllInbound": true,
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-nhn"
  ],
  "status": "Available",
  "storageSize": 100,
  "storageType": "General SSD",
  "subnetIds": [
    "subnet-1"
  ],
  "uid": "tbsn6ljiu7npo3re8iqr",
  "vNetId": "test-rdbms-vnet-nhn"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 25ms
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
      "backupTime": "03:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:07:21Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:07:21Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "14d502e5-fa5d-4412-8d93-d1ff9f7f71ad",
      "cspResourceName": "tbsn6ljiu7npo3re8iqr",
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL_V8046",
      "dbInstanceSpec": "m2.c2m4",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "9d20cc3f-d6dd-4b6f-a1a4-1ea553fd42ae.external.kr1.mysql.rds.nhncloudservice.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-nhn",
      "keyValueList": [
        {
          "key": "DBInstanceStatus",
          "value": "AVAILABLE"
        },
        {
          "key": "DBFlavorId",
          "value": "71f69bf9-3c01-4c1a-b135-bb75e93f6268"
        },
        {
          "key": "SubnetId",
          "value": "200d6e94-75e0-4809-ac92-bb11535fcfe1"
        }
      ],
      "name": "test-rdbms-nhn",
      "nhnDBSGToAllowAllInbound": true,
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-nhn"
      ],
      "status": "Available",
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
- **Duration:** 12.152s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-nhn/database`
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
- **Duration:** 3.536s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-nhn/database`
```json
// Response Body
{
  "databases": [
    "sampledb_dyn",
    "sampledb"
  ]
}
```

### 13. Beetle GET Secure Transport Info [✅ SUCCESS]
- **Duration:** 2.15s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-nhn/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": true,
    "issuer": "CN=MySQL_Server_8.0.46_Auto_Generated_CA_Certificate",
    "notAfter": "2036-11-22T10:06:14Z",
    "pem": "-----BEGIN CERTIFICATE-----\nMIIDBjCCAe6gAwIBAgIBATANBgkqhkiG9w0BAQsFADA8MTowOAYDVQQDDDFNeVNR\nTF9TZXJ2ZXJfOC4wLjQ2X0F1dG9fR2VuZXJhdGVkX0NBX0NlcnRpZmljYXRlMB4X\nDTI2MTEyNTEwMDYxNFoXDTM2MTEyMjEwMDYxNFowPDE6MDgGA1UEAwwxTXlTUUxf\nU2VydmVyXzguMC40Nl9BdXRvX0dlbmVyYXRlZF9DQV9DZXJ0aWZpY2F0ZTCCASIw\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAK3FI2soxHnmhF8euinQJ0WHTrWU\nB4Kcyjr1jbpB7eM9915NNuSFc8iwf6DkUy2lSPOc8wUwI31QLa9yZ/3qRqIAmVbg\nI97wJl+kd6xeLiYkFVGt98dAh5ujQvnRl+bulFhWqE0GDrgQPaoK2kCBf4Nx07tm\nu7FjsvKmbwU4muoW4YFMeR05h5e3Iwo1lmDOfp3AbCUXvMjRJ5/+RdZjV9V+nhGz\nZj5Gjs2UdEvXwlj03majSHbLah8uvC4NXbpxzSiwBK5Xg9TrvquctIctEJftAOB7\nlYugqQG+T6HrbF9/c/Bn03zyM93Owvej4P+pm+ZtrNhSN1oC4qzrn1CqZ1cCAwEA\nAaMTMBEwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAJFAk70Yh\nhZVkKhesd+oSbugRtSkzdEH/gFSeBD57+E9u6bV1C9dhlbUk3d51klmk2KkPbzq9\nI24PnD9D9pa8FaRxFemFQUOY9e3sYWKFly9QVj1ShFU5GvDxojoFV7cAvnG1xBXp\n7GgNgJEiusBeYfvQd7pUoMe/CJxCW4QED/3Gy6Qo2SWJ9TXMhTm41HXkhvGKuUkh\n8bv91NliVcSB9Bmb73AOiQ//03eQHAjFvXHvztqwiBKx4fLL0D23BIxhMsPLP++t\nqyYnWwO8HHQlPRWAQIET6tB+qXxn5flkHGLdpCVY9fAUqSLVEmKDtFlP48mlKx8n\nOlI+oAXywb7+MQ==\n-----END CERTIFICATE-----\n",
    "subject": "CN=MySQL_Server_8.0.46_Auto_Generated_CA_Certificate"
  },
  "enforced": false,
  "engine": "mysql",
  "recommendedSSLMode": "VERIFY_CA",
  "requireSecureTransport": "OFF",
  "rules": "",
  "tlsCipher": "TLS_AES_128_GCM_SHA256",
  "tlsInUse": true
}
```

### 14. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 35ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 7m40.455s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 36.754s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-nhn/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 1m21.078s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-nhn?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 2.535s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 48.787s

