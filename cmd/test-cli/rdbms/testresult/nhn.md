# Managed RDBMS Test Report: NHN (kr1)

- **Test Case:** NHN KR1 MySQL Test
- **Date & Time:** 2026-08-31 16:38:03
- **Namespace:** `default`
- **Total Duration:** 17m53.817s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** NHN
- **Target Region:** `kr1`
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
- **Duration:** 3.283s
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
  "message": "Spec and image pair is valid for provisioning",
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
  "status": "OK"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 31.687s
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
      "lastTransitionTime": "2026-08-31T07:38:38Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:38Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:38Z",
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
  "cspResourceId": "6199ad3d-aa49-4ad5-a021-19a92611b184",
  "cspResourceName": "tbk74g2ehs5eu21hah3p",
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
      "value": "2026-08-31 07:38:08"
    }
  ],
  "name": "test-rdbms-vnet-nhn",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-08-31T07:38:38Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-08-31T07:38:38Z",
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
      "cspResourceId": "8dfb9df9-e876-4426-8e99-63f3fee7ec9c",
      "cspResourceName": "tbv1fh0pc8rv6h44bblq",
      "cspVNetId": "6199ad3d-aa49-4ad5-a021-19a92611b184",
      "cspVNetName": "tbk74g2ehs5eu21hah3p",
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
          "value": "tbv1fh0pc8rv6h44bblq"
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
          "value": "8dfb9df9-e876-4426-8e99-63f3fee7ec9c"
        },
        {
          "key": "RoutingTable",
          "value": "{gateway_id:,default_table:false,explicit:false,id:,name:}"
        },
        {
          "key": "CreateTime",
          "value": "2026-08-31 07:38:33"
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
          "value": "6199ad3d-aa49-4ad5-a021-19a92611b184"
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
      "uid": "tbv1fh0pc8rv6h44bblq"
    }
  ],
  "systemLabel": "",
  "uid": "tbk74g2ehs5eu21hah3p"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 5.59s
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
  "cspResourceId": "a9b8bdd7-9836-4542-a4a4-acbebbed23df",
  "cspResourceName": "tb44thhuvu23m6tk23ue",
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
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "22",
      "Protocol": "TCP"
    }
  ],
  "id": "test-rdbms-sg-nhn",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "Name",
      "value": "tb44thhuvu23m6tk23ue"
    },
    {
      "key": "Description",
      "value": "tb44thhuvu23m6tk23ue"
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
  "uid": "tb44thhuvu23m6tk23ue",
  "vNetId": "test-rdbms-vnet-nhn"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 5ms
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
- **Duration:** 2.345s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=nhn-kr1`
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
          "minSize": 20,
          "recommendationLevel": "recommended",
          "recommended": true,
          "storageType": "General SSD"
        },
        {
          "constraints": "Minimum 20GB storage.",
          "description": "Cost-effective HDD storage for less demanding workloads. Minimum 20 GB.",
          "displayName": "General HDD",
          "minSize": 20,
          "recommendationLevel": "standard",
          "storageType": "General HDD"
        }
      ]
    },
    "providerName": "nhn",
    "regionName": "KR1",
    "requiresSecurityGroup": false,
    "requiresSubnet": true,
    "storageSizeRange": {
      "max": 2048,
      "min": 20
    },
    "storageTypeOptions": [
      "General SSD",
      "General HDD"
    ],
    "supportedVersions": [
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
  "desiredCloud": {
    "csp": "nhn",
    "region": "kr1"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for nhn (kr1)",
  "warnings": [
    "Requested storage type 'SSD' is not supported on target cloud. Replaced with capability recommended storage for instance 'source-mysql-01'."
  ],
  "targetCloud": {
    "csp": "nhn",
    "region": "kr1"
  },
  "targetRDBMSInstances": [
    {
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "rdbmsName": "test-rdbms-nhn",
      "dbEngine": "mysql",
      "dbEngineVersion": "MYSQL_V8046",
      "dbInstanceSpec": "m2.c2m4",
      "storageType": "General SSD",
      "storageSize": 100,
      "adminUserName": "myadmin",
      "adminUserPassword": "******",
      "vNetId": "test-rdbms-vnet-nhn",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-nhn"
      ],
      "publicAccess": true,
      "highAvailability": false,
      "backupRetentionDays": 7,
      "nhnDBSGToAllowAllInbound": true,
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
- **Duration:** 15ms
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
  "name": "test-rdbms-nhn",
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
    "name": "test-rdbms-nhn",
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
- **Duration:** 9m3.893s
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
      "rdbmsName": "test-rdbms-nhn",
      "securityGroupIds": [
        "test-rdbms-sg-nhn"
      ],
      "sourceInstanceName": "source-mysql-01",
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
- **Duration:** 2.474s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-nhn`
```json
// Response Body
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
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 12ms
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
      "status": "Deleting",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Deleting",
          "message": "RDBMS deletion in progress",
          "lastTransitionTime": "2026-08-31T07:47:26Z"
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
      "status": "Deleting",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Deleting",
          "message": "RDBMS deletion in progress",
          "lastTransitionTime": "2026-08-31T07:47:31Z"
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
- **Duration:** 12.279s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-nhn/database`
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
- **Duration:** 4.116s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-nhn/database`
```json
// Response Body
{
  "databases": [
    "sampledb_dyn",
    "sampledb"
  ]
}
```

### 13. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 313ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 5m5.802s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 35.728s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-nhn/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 1m22.184s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-nhn?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 3.942s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 40.145s

