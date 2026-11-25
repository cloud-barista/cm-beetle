# Managed RDBMS (MariaDB) Test Report: OPENSTACK (RegionOne)

- **Test Case:** OpenStack RegionOne MariaDB Test
- **Access Mode:** Public (publicAccess=true)
- **Date & Time:** 2026-11-25 19:27:54
- **Namespace:** `default`
- **Total Duration:** 10m48.978s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** OPENSTACK
- **Target Region:** `RegionOne`
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
- **Duration:** 1.163s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "78d90dae-d21d-4606-a9dd-c1268e321864",
  "specId": "openstack+regionone+m1.xlarge"
}
```
```json
// Response Body
{
  "connectionName": "openstack-regionone",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "openstack-regionone",
    "creationDate": "",
    "cspImageName": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "description": "",
    "details": [
      {
        "key": "ID",
        "value": "78d90dae-d21d-4606-a9dd-c1268e321864"
      },
      {
        "key": "Name",
        "value": "ubuntu-24.04"
      },
      {
        "key": "Status",
        "value": "active"
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
        "value": "0"
      },
      {
        "key": "MinRAMMegabytes",
        "value": "0"
      },
      {
        "key": "Owner",
        "value": "1b513759820b4fb497fac0661d81c347"
      },
      {
        "key": "Protected",
        "value": "false"
      },
      {
        "key": "Visibility",
        "value": "shared"
      },
      {
        "key": "Hidden",
        "value": "false"
      },
      {
        "key": "Checksum",
        "value": "77a77cff0d62968d5166322ea2b78008"
      },
      {
        "key": "CreatedAt",
        "value": "2025-09-16T05:28:43Z"
      },
      {
        "key": "UpdatedAt",
        "value": "2025-09-16T05:28:55Z"
      },
      {
        "key": "File",
        "value": "/v2/images/78d90dae-d21d-4606-a9dd-c1268e321864/file"
      },
      {
        "key": "Schema",
        "value": "/v2/schemas/image"
      },
      {
        "key": "VirtualSize",
        "value": "3758096384"
      }
    ],
    "fetchedTime": "2026.10.29 04:24:54 Thu",
    "id": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 0,
    "osDiskType": "NA",
    "osDistribution": "ubuntu-24.04",
    "osPlatform": "NA",
    "osType": "Ubuntu 24.04",
    "providerName": "openstack",
    "regionList": [
      "regionone"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tbde4jmfid1fbfeg61gi"
  },
  "imageId": "78d90dae-d21d-4606-a9dd-c1268e321864",
  "imageValidation": {
    "cspResourceId": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "isAvailable": true,
    "resourceId": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "resourceName": "78d90dae-d21d-4606-a9dd-c1268e321864",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid but has warnings",
  "providerName": "openstack",
  "regionName": "regionone",
  "specDetails": {
    "architecture": "x86_64",
    "availabilityStatus": "Unverified",
    "connectionName": "openstack-regionone",
    "costPerHour": -1,
    "cspSpecName": "m1.xlarge",
    "details": [
      {
        "key": "ID",
        "value": "61b7f2ff-c21c-4216-b5ca-dbda45b551da"
      },
      {
        "key": "Disk",
        "value": "160"
      },
      {
        "key": "RAM",
        "value": "16384"
      },
      {
        "key": "Name",
        "value": "m1.xlarge"
      },
      {
        "key": "RxTxFactor",
        "value": "1.00"
      },
      {
        "key": "VCPUs",
        "value": "8"
      },
      {
        "key": "IsPublic",
        "value": "true"
      },
      {
        "key": "Ephemeral",
        "value": "0"
      }
    ],
    "diskSizeGB": 160,
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
    "id": "openstack+regionone+m1.xlarge",
    "infraType": "node",
    "memoryGiB": 16,
    "name": "openstack+regionone+m1.xlarge",
    "namespace": "system",
    "providerName": "openstack",
    "regionLatitude": 36.3804,
    "regionLongitude": 127.365,
    "regionName": "regionone",
    "rootDiskSize": 160,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbmeg4q8vljrvg05vr35",
    "vCPU": 8
  },
  "specId": "openstack+regionone+m1.xlarge",
  "specValidation": {
    "cspResourceId": "m1.xlarge",
    "isAvailable": true,
    "message": "no availability checker registered for provider; assumed available",
    "resourceId": "openstack+regionone+m1.xlarge",
    "resourceName": "m1.xlarge",
    "status": "Available"
  },
  "status": "Warning",
  "warnings": [
    "Zone availability notice for 'm1.xlarge': no availability checker registered for provider; assumed available"
  ]
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 4.538s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.6.0.0/16",
  "connectionName": "openstack-regionone",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-openstack",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.6.1.0/24",
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
  "cidrBlock": "OpenStack VPC does not support IPv4_CIDR",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T10:27:59Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T10:27:59Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-11-25T10:27:59Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "54d61f70-86f8-44c3-ac0e-6b7ad2ec93b7",
  "cspResourceName": "tb32rr079fkjibbii4r5",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-openstack",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "Network",
      "value": "{id:54d61f70-86f8-44c3-ac0e-6b7ad2ec93b7,name:tb32rr079fkjibbii4r5,description:,admin_state_up:true,status:ACTIVE,subnets:[7fa67170-ae05-4108-bf47-1cbfcc864346],tenant_id:1b513759820b4fb497fac0661d81c347,project_id:1b513759820b4fb497fac0661d81c347,shared:false,availability_zone_hints:[nova],tags:[],revision_number:2}"
    },
    {
      "key": "NetworkExternalExt",
      "value": "{router:external:false}"
    }
  ],
  "name": "test-rdbms-vnet-openstack",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:27:59Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:27:59Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "7fa67170-ae05-4108-bf47-1cbfcc864346",
      "cspResourceName": "tbg344iv8s8oo8l9q4o7",
      "cspVNetId": "54d61f70-86f8-44c3-ac0e-6b7ad2ec93b7",
      "cspVNetName": "tb32rr079fkjibbii4r5",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.6.1.0/24",
      "keyValueList": [
        {
          "key": "ID",
          "value": "7fa67170-ae05-4108-bf47-1cbfcc864346"
        },
        {
          "key": "NetworkID",
          "value": "54d61f70-86f8-44c3-ac0e-6b7ad2ec93b7"
        },
        {
          "key": "Name",
          "value": "tbg344iv8s8oo8l9q4o7"
        },
        {
          "key": "IPVersion",
          "value": "4"
        },
        {
          "key": "CIDR",
          "value": "10.6.1.0/24"
        },
        {
          "key": "GatewayIP",
          "value": "10.6.1.1"
        },
        {
          "key": "DNSNameservers",
          "value": "8.8.8.8"
        },
        {
          "key": "DNSPublishFixedIP",
          "value": "false"
        },
        {
          "key": "AllocationPools",
          "value": "{start:10.6.1.2,end:10.6.1.254}"
        },
        {
          "key": "EnableDHCP",
          "value": "true"
        },
        {
          "key": "TenantID",
          "value": "1b513759820b4fb497fac0661d81c347"
        },
        {
          "key": "ProjectID",
          "value": "1b513759820b4fb497fac0661d81c347"
        },
        {
          "key": "RevisionNumber",
          "value": "0"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbg344iv8s8oo8l9q4o7"
    }
  ],
  "systemLabel": "",
  "uid": "tb32rr079fkjibbii4r5"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 4.552s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "openstack-regionone",
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
  "name": "test-rdbms-sg-openstack",
  "vNetId": "test-rdbms-vnet-openstack"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "8ebd3c41-e333-402d-86ea-f43ae9cfd37a",
  "cspResourceName": "tbl5dr9kg1n59e77tm0n",
  "description": "Pre-requisite SecurityGroup for CM-Beetle RDBMS test",
  "firewallRules": [
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "outbound",
      "Port": "",
      "Protocol": "ALL"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "3306",
      "Protocol": "TCP"
    },
    {
      "CIDR": "::/0",
      "Direction": "outbound",
      "Port": "",
      "Protocol": "ALL"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "22",
      "Protocol": "TCP"
    }
  ],
  "id": "test-rdbms-sg-openstack",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "IId",
      "value": "{NameId:tbl5dr9kg1n59e77tm0n,SystemId:8ebd3c41-e333-402d-86ea-f43ae9cfd37a}"
    },
    {
      "key": "VpcIID",
      "value": "{NameId:,SystemId:}"
    },
    {
      "key": "SecurityRules",
      "value": "[{Direction:outbound,IPProtocol:all,FromPort:-1,ToPort:-1,CIDR:0.0.0.0/0},{Direction:inbound,IPProtocol:tcp,FromPort:3306,ToPort:3306,CIDR:0.0.0.0/0},{Direction:outbound,IPProtocol:all,FromPort:-1,ToPort:-1,CIDR:::/0},{Direction:inbound,IPProtocol:tcp,FromPort:22,ToPort:22,CIDR:0.0.0.0/0}]"
    }
  ],
  "name": "test-rdbms-sg-openstack",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbl5dr9kg1n59e77tm0n",
  "vNetId": "test-rdbms-vnet-openstack"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 6ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=openstack`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "openstack": {
      "dbOperationMethod": "cspNativeApi",
      "note": "CB-Spider v0.13.15 (#1873, fixing #1866) retries admin privilege GRANT with exponential backoff after Trove reports ACTIVE to prevent Error 1044; in private mode (PublicAccess=false), CB-Spider (#1890) skips this direct-SQL GRANT step. OpenStack Trove 2025.2 (Flamingo) active datastore versions are MySQL 8.4 and MariaDB 11.4.",
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
- **Duration:** 1.132s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=openstack-regionone&dbEngine=mariadb`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "NA",
    "connectionName": "openstack-regionone",
    "dbEngine": "mariadb",
    "dbInstanceSpecOptions": [
      "m1.amphora",
      "m1.large",
      "m1.medium",
      "m1.small",
      "m1.tiny",
      "m1.xlarge"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "1024",
        "name": "m1.amphora",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "8192",
        "name": "m1.large",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "4096",
        "name": "m1.medium",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "2048",
        "name": "m1.small",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "512",
        "name": "m1.tiny",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "16384",
        "name": "m1.xlarge",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
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
          "constraints": "Minimum 1GB storage.",
          "description": "Ceph RADOS Block Device storage (if available in deployment).",
          "displayName": "Ceph RBD",
          "minSizeGB": 1,
          "recommendationLevel": "standard",
          "storageType": "RBD"
        },
        {
          "constraints": "Minimum 1GB storage.",
          "description": "Default storage type provided by OpenStack deployment. Actual backend depends on deployment configuration.",
          "displayName": "Default Storage",
          "minSizeGB": 1,
          "recommendationLevel": "standard",
          "storageType": "__DEFAULT__"
        }
      ]
    },
    "providerName": "openstack",
    "regionName": "RegionOne",
    "requiresSecurityGroup": false,
    "requiresSubnet": false,
    "storageSizeRangeGB": {
      "max": -1,
      "min": 1
    },
    "storageTypeOptions": [
      "RBD",
      "__DEFAULT__"
    ],
    "supportedVersions": [
      "10.4",
      "11.4"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": false,
    "supportsEncryption": false,
    "supportsHighAvailability": false,
    "supportsPublicAccess": true,
    "supportsStorageSizeConfiguration": true,
    "supportsStorageTypeSelection": true,
    "supportsTag": false
  }
}
```

### 6. Beetle POST Recommend RDBMS [✅ SUCCESS]
- **Duration:** 8ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "openstack",
    "region": "RegionOne"
  },
  "sourceRDBMSInstances": [
    {
      "dbEngine": {
        "engine": "mariadb",
        "engineVersion": "11.4",
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
    "adminUserName": "admin",
    "backupRetentionDays": 7,
    "highAvailability": false,
    "publicAccess": true
  }
}
```
```json
// Response Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for openstack (regionone)",
  "status": "recommended",
  "targetCloud": {
    "csp": "openstack",
    "region": "regionone"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "admin",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "charset": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mariadb",
      "dbEngineVersion": "11.4",
      "dbInstanceSpec": "m1.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-mariadb-openstack",
      "securityGroupIds": [
        "test-rdbms-sg-openstack"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "RBD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-openstack"
    }
  ],
  "warnings": [
    "Storage type 'SSD' requested for instance 'Source MySQL 01' is not directly offered on openstack; recommended closest available storage type 'RBD' (available: [RBD __DEFAULT__])."
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 22ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "admin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "openstack-regionone",
  "dbEngine": "mariadb",
  "dbEngineVersion": "11.4",
  "dbInstanceSpec": "m1.medium",
  "name": "rdbms-mariadb-openstack",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-openstack"
  ],
  "storageSize": 100,
  "storageType": "RBD",
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-openstack"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "admin",
    "adminUserPassword": "******",
    "connectionName": "openstack-regionone",
    "dbEngine": "mariadb",
    "dbEngineVersion": "11.4",
    "dbInstanceSpec": "m1.medium",
    "name": "rdbms-mariadb-openstack",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-openstack"
    ],
    "storageSize": 100,
    "storageType": "RBD",
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-openstack"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 3m45.728s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for openstack (regionone)",
  "status": "recommended",
  "targetCloud": {
    "csp": "openstack",
    "region": "regionone"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "admin",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "charset": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mariadb",
      "dbEngineVersion": "11.4",
      "dbInstanceSpec": "m1.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-mariadb-openstack",
      "securityGroupIds": [
        "test-rdbms-sg-openstack"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "RBD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-openstack"
    }
  ],
  "warnings": [
    "Storage type 'SSD' requested for instance 'Source MySQL 01' is not directly offered on openstack; recommended closest available storage type 'RBD' (available: [RBD __DEFAULT__])."
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
- **Duration:** 1.233s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack`
```json
// Response Body
{
  "backupTime": "NA",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T10:31:29Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T10:31:29Z",
      "reason": "Available",
      "status": "True",
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
  "cspResourceId": "66740741-d1a0-41c9-b770-a78e8b885198",
  "cspResourceName": "tb0lg7ct0lnh5v9libpg",
  "dbEngine": "mariadb",
  "dbEngineVersion": "11.4",
  "dbInstanceSpec": "m1.medium",
  "dbInstanceType": "NA",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "183.111.177.152:3306",
  "highAvailability": false,
  "id": "test-rdbms-mariadb-openstack",
  "keyValueList": [
    {
      "key": "Hostname",
      "value": ""
    },
    {
      "key": "VolumeUsed",
      "value": "0.10"
    }
  ],
  "name": "test-rdbms-mariadb-openstack",
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-openstack"
  ],
  "status": "Available",
  "storageSize": 100,
  "storageType": "NA",
  "subnetIds": [
    "subnet-1"
  ],
  "uid": "tb0lg7ct0lnh5v9libpg",
  "vNetId": "test-rdbms-vnet-openstack"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 32ms
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
      "backupTime": "NA",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:31:29Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:31:29Z",
          "reason": "Available",
          "status": "True",
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
      "cspResourceId": "66740741-d1a0-41c9-b770-a78e8b885198",
      "cspResourceName": "tb0lg7ct0lnh5v9libpg",
      "dbEngine": "mariadb",
      "dbEngineVersion": "11.4",
      "dbInstanceSpec": "m1.medium",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "183.111.177.152:3306",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-openstack",
      "keyValueList": [
        {
          "key": "Hostname",
          "value": ""
        },
        {
          "key": "VolumeUsed",
          "value": "0.10"
        }
      ],
      "name": "test-rdbms-mariadb-openstack",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-openstack"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "NA",
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
- **Duration:** 8.038s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack/database`
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
- **Duration:** 3.536s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack/database`
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
- **Duration:** 2.287s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": false,
    "issuer": "CN=MariaDB Server",
    "notAfter": "2036-11-22T10:31:03Z",
    "pem": "-----BEGIN CERTIFICATE-----\nMIIEqzCCApOgAwIBAgIBADANBgkqhkiG9w0BAQsFADAZMRcwFQYDVQQDDA5NYXJp\nYURCIFNlcnZlcjAeFw0yNjExMjUxMDMxMDNaFw0zNjExMjIxMDMxMDNaMBkxFzAV\nBgNVBAMMDk1hcmlhREIgU2VydmVyMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIIC\nCgKCAgEAiC3AlT4BN/dZntOD2VI8bFjjyqV6zNfsBb3AiMdBdJUcqRaBLtTFt+T/\niMRF+r4qSXbhI9oEcGlaMsLHTdGDIEu8YM3GygXZrlUP5+M6H5ISiXBWvA1Ewkwi\nAmxDjNuTKKSgipbaWiIeSFpUfw8PWZmrk6GgqpGtAZd39ubABzdgfxandEplk32w\nHAapmch7IvWDD6iHUE1YKMctV+HG0xsexsOczuzWyRwlJMZQt1SnXhfb1Bej3jTy\nFUBR28C8edjzzjJ/2s61SA8sYcqnFTu959DckOSzFpBYajo2bYZF5in0F2j0JOWr\np23asDpYogdvb3gQG4PPEOel+QXFghq9K5M9VWyBu1UjrC2Q5T1Udm6P9FV3UmXd\nt4/IGcrxody3c2d+2v+cobKunyxw5Ml7SEg+OqC8EHLf0g2/SZoIP51AFZFfGrbd\nOQpoErlPQcLEXd3oOgHq5XeY6lwprkOo2/1SC1mshVfavh00alSvrgOO+caEsysI\n8nAJmyFjsx8Ow5e0Xwe2MiByKTaAQsRF0iqzlVTD36xjVXnSkBeSOndZNbcegfyp\nvJ8ayqaj1+caykjsb18te/3ch2Vc81l/LL/xraXtjhixM7AwggjreFpnKBVlns5O\nmBgsWReRdrXtf44rj4ikWSlGQYjDhWfVRxEDwv48efWW2xSmKukCAwEAATANBgkq\nhkiG9w0BAQsFAAOCAgEAInMqvkpFnnfIoj7xOrXBI4UtVNIv+6CuCxuNAD4buv79\nW0iEgwKOI/LFZyGHMuOAbpEaIqqtY8lAN456Kcc4L5jEEsldcJDdLgINlFRpeGZA\nW+OSvralA/wgdnfMOyEWW07R5TakXHNsFuO9zf5H2tUBkOKxOcgUUaZIztdqWItS\nT88g1hbtuOjNjU3sQbiAJ0ak+UjGQUPPKNaCkUXCqHjgSBNvogm9jKzYDVPbQw2R\nu7TdcTh6h7NBFFYV7vMtDHQss8P+TOM+KT+zTcnZDqw6QlKKTxlffa+MLkxzo1As\nWCQcP0ZgPO4bQIBOQaVZgJcq4wBU/stBlRMASZE24WmaIqfKL8rMhbtTNvW1Jr/v\n1SuuaODS+9Ks8j9wuDe8TDhtTZxry9Lu7msnCEDb0Seb2E2rMR+CpYoDEAypcj8G\nClHIc89PZz/huh8tA2VydLMAEVcPpAvUPHYmwPRDsj11pDpG8oWC16VQHeeHud8i\nsbtMn+k1SMo+oSMK5yB2h/bxY8Kvw0DmZ7AvdNc+TUl+j3DnHEbsleoN0Pc+UYmB\n+dfXxZQJqH2zZ+lAXGEVI6HDSoo3dP8fC41NnvRZPtygnWSQ19DXVxHPgeyJSNKa\nLsmUTKu+o0DMK43zX7EIjfQpbRcIXHVJUKsraUEiRm74HcLrTw1ObamAdPCyVwY=\n-----END CERTIFICATE-----\n",
    "subject": "CN=MariaDB Server"
  },
  "enforced": false,
  "engine": "mariadb",
  "recommendedSSLMode": "VERIFY_CA",
  "requireSecureTransport": "OFF",
  "rules": "",
  "tlsCipher": "TLS_AES_128_GCM_SHA256",
  "tlsInUse": true
}
```

### 14. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 36ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 5m19.15s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 19.419s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 51.645s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-openstack?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 2.301s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 4.154s

