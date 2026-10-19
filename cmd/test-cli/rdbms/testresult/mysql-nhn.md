# Managed RDBMS (MySQL) Test Report: NHN (kr1)

- **Test Case:** NHN KR1 MySQL Test
- **Date & Time:** 2026-10-19 09:49:54
- **Namespace:** `default`
- **Total Duration:** 18m22.884s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** NHN
- **Target Region:** `kr1`
- **Namespace:** `default`
- **Test Date:** 2026-10-19 09:49:54

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
- **Duration:** 3.177s
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
    "fetchedTime": "2026.08.21 13:57:30 Fri",
    "id": "5c38715f-0375-4167-af4e-56f75ba8b252",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": true,
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
    "uid": "tb713bp8hrirhkg81rh8"
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
    "uid": "tbpnvei7jqqgj4d9678f",
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
    "Moderate provisioning failure risk: Spec 'nhn+kr1+m2.c2m4' has 50% failure rate (1 failures, 1 successes) | Image '5c38715f-0375-4167-af4e-56f75ba8b252' with spec 'nhn+kr1+m2.c2m4' has succeeded before, but also had past failures. Recent failures: VM provisioning failed (status: Terminating) - check CSP console for details"
  ]
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 35.551s
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
      "lastTransitionTime": "2026-10-19T00:50:30Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-19T00:50:30Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-10-19T00:50:30Z",
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
  "cspResourceId": "0d7e79eb-beb6-483f-9c51-5caea37e071a",
  "cspResourceName": "tbgrflds24l9g5iu65pm",
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
      "value": "2026-10-19 00:49:58"
    }
  ],
  "name": "test-rdbms-vnet-nhn",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T00:50:30Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T00:50:30Z",
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
      "cspResourceId": "a0998a30-4827-4989-9fc4-db6a78bc934e",
      "cspResourceName": "tb3qthqj4f65cu27bmeb",
      "cspVNetId": "0d7e79eb-beb6-483f-9c51-5caea37e071a",
      "cspVNetName": "tbgrflds24l9g5iu65pm",
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
          "value": "tb3qthqj4f65cu27bmeb"
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
          "value": "a0998a30-4827-4989-9fc4-db6a78bc934e"
        },
        {
          "key": "RoutingTable",
          "value": "{gateway_id:,default_table:false,explicit:false,id:,name:}"
        },
        {
          "key": "CreateTime",
          "value": "2026-10-19 00:50:25"
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
          "value": "0d7e79eb-beb6-483f-9c51-5caea37e071a"
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
      "uid": "tb3qthqj4f65cu27bmeb"
    }
  ],
  "systemLabel": "",
  "uid": "tbgrflds24l9g5iu65pm"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 7.139s
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
  "cspResourceId": "9d57c1e5-c800-40cc-94c7-6550c442552a",
  "cspResourceName": "tbq7m29lpra4gjo57r1o",
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
  "id": "test-rdbms-sg-nhn",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "Name",
      "value": "tbq7m29lpra4gjo57r1o"
    },
    {
      "key": "Description",
      "value": "tbq7m29lpra4gjo57r1o"
    },
    {
      "key": "Rules",
      "value": "{from_port:3306,to_port:3306,ip_protocol:tcp,ip_range:{CIDR:0.0.0.0/0},Group:{tenant_id:,Name:}}; {from_port:22,to_port:22,ip_protocol:tcp,ip_range:{CIDR:0.0.0.0/0},Group:{tenant_id:,Name:}}"
    },
    {
      "key": "TenantID",
      "value": "6fc6c2ef568f45f7a9bdc6be211f52f9"
    }
  ],
  "name": "test-rdbms-sg-nhn",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbq7m29lpra4gjo57r1o",
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
- **Duration:** 1.917s
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
- **Duration:** 3ms
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
- **Duration:** 12ms
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
- **Duration:** 9m3.003s
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
- **Duration:** 2.316s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-nhn`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "03:00",
  "conditions": [
    {
      "lastTransitionTime": "2026-10-19T00:58:48Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-19T00:58:48Z",
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
  "cspResourceId": "e0ea4718-41a9-490a-b400-640d6022ed39",
  "cspResourceName": "tb2s1h0bch71jggm9aqn",
  "dbEngine": "mysql",
  "dbEngineVersion": "MYSQL_V8046",
  "dbInstanceSpec": "m2.c2m4",
  "dbInstanceType": "NA",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "8ede9688-9fbb-4350-bd7c-2ec5371cef17.external.kr1.mysql.rds.nhncloudservice.com:3306",
  "highAvailability": false,
  "id": "test-rdbms-nhn",
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
  "uid": "tb2s1h0bch71jggm9aqn",
  "vNetId": "test-rdbms-vnet-nhn"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 8ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "backupTime": "16:00Z-17:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T00:56:55Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T00:56:55Z",
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
      "cspResourceId": "rm-mj77hqffq9u8lfd98",
      "cspResourceName": "tb4canurhcvccfe1cd7q",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "dbInstanceType": "Basic",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "43.108.67.176:3306",
      "highAvailability": false,
      "id": "test-rdbms-alibaba",
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
      "uid": "tb4canurhcvccfe1cd7q",
      "vNetId": "test-rdbms-vnet-alibaba"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "17:34-18:04",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T00:50:49Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T00:50:49Z",
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
      "cspResourceId": "tbmo5d06mbrtcdfjposi",
      "cspResourceName": "tbmo5d06mbrtcdfjposi",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.46-rds.20260908",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "tbmo5d06mbrtcdfjposi.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com:3306",
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
          "value": "tbmo5d06mbrtcdfjposi"
        }
      ],
      "uid": "tbmo5d06mbrtcdfjposi",
      "vNetId": "test-rdbms-vnet-aws"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "AUTO",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T00:54:15Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T00:54:15Z",
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
      "cspResourceId": "tb30d3dnjrhkt7rajm80",
      "cspResourceName": "tb30d3dnjrhkt7rajm80",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "dbInstanceType": "Burstable",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "tb30d3dnjrhkt7rajm80.mysql.database.azure.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-azure",
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
      "uid": "tb30d3dnjrhkt7rajm80",
      "vNetId": "test-rdbms-vnet-azure"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "20:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T00:58:33Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T00:54:00Z",
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
      "cspResourceId": "tbj7djjk245lg59evd6k",
      "cspResourceName": "tbj7djjk245lg59evd6k",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "dbInstanceType": "ZONAL",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "encryption": true,
      "endpoint": "34.9.134.166:3306",
      "highAvailability": false,
      "id": "test-rdbms-gcp",
      "name": "test-rdbms-gcp",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "status": "Deleting",
      "storageSize": 100,
      "storageType": "PD_SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbj7djjk245lg59evd6k",
      "vNetId": "test-rdbms-vnet-gcp"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T00:50:27Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T00:50:27Z",
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
      "uid": "tblp8rdv0nfnboaibv38",
      "vNetId": "test-rdbms-vnet-ibm"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T00:50:44Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T00:50:44Z",
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
      "uid": "tbs39bb5093bjft6118j",
      "vNetId": "test-rdbms-vnet-ncp"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "03:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T00:58:48Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T00:58:48Z",
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
      "cspResourceId": "e0ea4718-41a9-490a-b400-640d6022ed39",
      "cspResourceName": "tb2s1h0bch71jggm9aqn",
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL_V8046",
      "dbInstanceSpec": "m2.c2m4",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "8ede9688-9fbb-4350-bd7c-2ec5371cef17.external.kr1.mysql.rds.nhncloudservice.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-nhn",
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
      "uid": "tb2s1h0bch71jggm9aqn",
      "vNetId": "test-rdbms-vnet-nhn"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "00:00-12:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T00:57:46Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T00:54:07Z",
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
      "cspResourceId": "cdb-d1hzepv8",
      "cspResourceName": "tbvl2gccnqjborsmdd2u",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "kr-cdb-d1hzepv8.sql.tencentcdb.com:21946",
      "highAvailability": false,
      "id": "test-rdbms-tencent",
      "name": "test-rdbms-tencent",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "status": "Deleting",
      "storageSize": 100,
      "storageType": "CLOUD_HSSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbvl2gccnqjborsmdd2u",
      "vNetId": "test-rdbms-vnet-tencent"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 12.506s
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
- **Duration:** 3.638s
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
- **Duration:** 2.209s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-nhn/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": true,
    "issuer": "CN=MySQL_Server_8.0.46_Auto_Generated_CA_Certificate",
    "notAfter": "2036-10-16T00:57:47Z",
    "pem": "-----BEGIN CERTIFICATE-----\nMIIDBjCCAe6gAwIBAgIBATANBgkqhkiG9w0BAQsFADA8MTowOAYDVQQDDDFNeVNR\nTF9TZXJ2ZXJfOC4wLjQ2X0F1dG9fR2VuZXJhdGVkX0NBX0NlcnRpZmljYXRlMB4X\nDTI2MTAxOTAwNTc0N1oXDTM2MTAxNjAwNTc0N1owPDE6MDgGA1UEAwwxTXlTUUxf\nU2VydmVyXzguMC40Nl9BdXRvX0dlbmVyYXRlZF9DQV9DZXJ0aWZpY2F0ZTCCASIw\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALzJmGKnViSYwfZ1aIfJJ3hVdnR5\nQYRb+jw7RK/Xf36WDab34+ifKxHxpoTBQ04aPICpfzOkyYR5Zg9evdL5jJvcM0w2\nspxIcnKfC3mJ6OOl9aK/j0BjSAmIyLXjGp5QHuI796V1zWdi+oGcpI64xO3AoACv\njSxeI9KljlgnhiEpxtMnCD8NhcIRnR75GAtukopX5YIdaNq2xYR9o9EHOMv06yxu\njDgU/NhLCWHcPCq+3c5yX90lnrjpMxnvfoPG5pLTvSeuRmShgItjezJ7aSrFDqlA\nfLpXBjdkAnV7/pDi5x3wq40YMmJkRYFWRXfqOSTlv2uZTQ0qj8NrrOVnpy0CAwEA\nAaMTMBEwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAJGuMeROU\nKfuGnR2DUvUh/tC4Hml4K4HCdAhGttKeXqXNE68uejBr29A0jpBN7DSzBQ2CqNrk\nxF2VnwhTm9vSVK+lgJ09w+5vODoGD2M0XxUAgiFHx2IGDTv61YvPwWTc4PNStvcw\nkFMGVrwcrkObKnEl9ASRM7hv2jaXz0qYzg5W61mDFu6Y0r4wak+RDaC9kQwm/HzK\nGcjQH2jD7fMD7a6GAc2sxrVP64JbkkN4IZGBWz/kd0PxYLV+A/NnBwBcdyqdhd2Z\nWW1OnwAVOUsgKP0rknmpDEcR0PjLhuImIWEbkHR0MJQQSh7sRPzBmvEI/63q7m2Z\nr/hYLq3xq5JqcA==\n-----END CERTIFICATE-----\n",
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
- **Duration:** 30ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 5m21.965s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 40.018s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-nhn/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 1m21.206s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-nhn?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 4.148s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 44.032s

