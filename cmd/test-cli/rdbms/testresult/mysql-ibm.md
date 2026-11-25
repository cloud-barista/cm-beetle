# Managed RDBMS (MySQL) Test Report: IBM (us-south)

- **Test Case:** IBM US-South MySQL Test
- **Access Mode:** Public (publicAccess=true)
- **Date & Time:** 2026-11-25 18:54:32
- **Namespace:** `default`
- **Total Duration:** 17m30.868s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** IBM
- **Target Region:** `us-south`
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
- **Duration:** 3.525s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22",
  "specId": "ibm+us-south+cxf-2x4"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "cached": true,
    "instanceType": "cxf-2x4",
    "provider": "ibm",
    "queriedAt": "2026-11-25T09:54:35.20444583Z",
    "refreshAfter": "2026-11-25T10:54:35.20444583Z",
    "region": "us-south",
    "source": "ibm:GetInstanceProfile",
    "status": "Available",
    "ttlSeconds": 3600,
    "zones": [
      {
        "available": true,
        "status": "current",
        "zoneId": "us-south-1"
      },
      {
        "available": true,
        "status": "current",
        "zoneId": "us-south-2"
      },
      {
        "available": true,
        "status": "current",
        "zoneId": "us-south-3"
      }
    ]
  },
  "connectionName": "ibm-us-south",
  "estimatedCost": "$0.0850/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "ibm-us-south",
    "creationDate": "",
    "cspImageName": "r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22",
    "description": "",
    "details": [
      {
        "key": "AllowedUse",
        "value": "{api_version:2024-11-28,bare_metal_server:true,instance:true}"
      },
      {
        "key": "CatalogOffering",
        "value": "{managed:false}"
      },
      {
        "key": "CreatedAt",
        "value": "2026-10-27T05:55:45.000Z"
      },
      {
        "key": "CRN",
        "value": "crn:v1:bluemix:public:is:us-south:a/811f8abfbd32425597dc7ba40da98fa6::image:r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22"
      },
      {
        "key": "Encryption",
        "value": "none"
      },
      {
        "key": "File",
        "value": "{checksums:{sha256:3f1c4200aaf0e4f4dd8ac4eac7e132a23e6a8b627022e1e54ea6e9d1c1ce7275},size:2}"
      },
      {
        "key": "Href",
        "value": "https://us-south.iaas.cloud.ibm.com/v1/images/r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22"
      },
      {
        "key": "ID",
        "value": "r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22"
      },
      {
        "key": "MinimumProvisionedSize",
        "value": "10"
      },
      {
        "key": "Name",
        "value": "ibm-ubuntu-24-04-5-minimal-amd64-2"
      },
      {
        "key": "OperatingSystem",
        "value": "{allow_user_image_creation:true,architecture:amd64,dedicated_host_only:false,display_name:Ubuntu Linux 24.04 LTS Noble Numbat Minimal Install (amd64),family:Ubuntu Linux,href:https://us-south.iaas.cloud.ibm.com/v1/operating_systems/ubuntu-24-04-amd64,name:ubuntu-24-04-amd64,user_data_format:cloud_init,vendor:Canonical,version:24.04 LTS Noble Numbat Minimal Install}"
      },
      {
        "key": "Remote",
        "value": "{account:{id:811f8abfbd32425597dc7ba40da98fa6,resource_type:account}}"
      },
      {
        "key": "ResourceGroup",
        "value": "{href:https://resource-controller.cloud.ibm.com/v1/resource_groups/5807b5832a8741179b2e06ca2d2b3b96,id:5807b5832a8741179b2e06ca2d2b3b96,name:Default}"
      },
      {
        "key": "ResourceType",
        "value": "image"
      },
      {
        "key": "Status",
        "value": "available"
      },
      {
        "key": "UserDataFormat",
        "value": "cloud_init"
      },
      {
        "key": "Visibility",
        "value": "public"
      }
    ],
    "fetchedTime": "2026.10.29 04:27:42 Thu",
    "id": "r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 10,
    "osDiskType": "NA",
    "osDistribution": "Ubuntu Linux 24.04 LTS Noble Numbat Minimal Install (amd64)",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "ibm",
    "regionList": [
      "us-south"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tbbs95ribpl755c3jtd0"
  },
  "imageId": "r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22",
  "imageValidation": {
    "cspResourceId": "r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22",
    "isAvailable": true,
    "resourceId": "r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22",
    "resourceName": "r006-d0870d62-2c9e-449d-a7b8-bcd41e5fef22",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "ibm",
  "regionName": "us-south",
  "specDetails": {
    "architecture": "x86_64",
    "availabilityCheckedAt": "2026-11-25T09:54:35.20444583Z",
    "availabilityRefreshAfter": "2026-11-25T10:54:35.20444583Z",
    "availabilityStatus": "Available",
    "availableZones": [
      "us-south-1",
      "us-south-2",
      "us-south-3"
    ],
    "connectionName": "ibm-us-south",
    "costPerHour": 0.085,
    "cspSpecName": "cxf-2x4",
    "details": [
      {
        "key": "AvailabilityClass",
        "value": "{default:standard,type:enum,values:[standard,spot]}"
      },
      {
        "key": "Bandwidth",
        "value": "{type:fixed,value:4000}"
      },
      {
        "key": "ClusterNetworkAttachmentCount",
        "value": "{type:enum,values:[0]}"
      },
      {
        "key": "ConfidentialComputeModes",
        "value": "{default:disabled,type:enum,values:[disabled]}"
      },
      {
        "key": "Family",
        "value": "compute"
      },
      {
        "key": "Href",
        "value": "https://us-south.iaas.cloud.ibm.com/v1/instance/profiles/cxf-2x4"
      },
      {
        "key": "Memory",
        "value": "{type:fixed,value:4}"
      },
      {
        "key": "Name",
        "value": "cxf-2x4"
      },
      {
        "key": "NetworkAttachmentCount",
        "value": "{max:5,min:1,type:range}"
      },
      {
        "key": "NetworkBandwidthMode",
        "value": "{type:fixed,value:divided}"
      },
      {
        "key": "NetworkInterfaceCount",
        "value": "{max:5,min:1,type:range}"
      },
      {
        "key": "NumaCount",
        "value": "{type:fixed,value:1}"
      },
      {
        "key": "OsArchitecture",
        "value": "{default:amd64,type:enum,values:[amd64]}"
      },
      {
        "key": "PortSpeed",
        "value": "{type:fixed,value:25000}"
      },
      {
        "key": "ReservationTerms",
        "value": "{type:enum,values:[one_year,three_year]}"
      },
      {
        "key": "ResourceType",
        "value": "instance_profile"
      },
      {
        "key": "SecureBootModes",
        "value": "{default:false,type:enum,values:[true,false]}"
      },
      {
        "key": "Status",
        "value": "current"
      },
      {
        "key": "TotalVolumeBandwidth",
        "value": "{type:range,default:1000,max:3500,min:500,step:1}"
      },
      {
        "key": "VcpuArchitecture",
        "value": "{type:fixed,value:amd64}"
      },
      {
        "key": "VcpuBurstLimit",
        "value": "{type:fixed,value:200}"
      },
      {
        "key": "VcpuCount",
        "value": "{type:fixed,value:2}"
      },
      {
        "key": "VcpuManufacturer",
        "value": "{type:dependent}"
      },
      {
        "key": "VcpuPercentage",
        "value": "{default:100,type:enum,values:[25,50,100]}"
      },
      {
        "key": "VolumeBandwidthQosModes",
        "value": "{default:pooled,type:enum,values:[weighted,pooled]}"
      },
      {
        "key": "Zones",
        "value": "{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/us-south-1,name:us-south-1}; {href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/us-south-2,name:us-south-2}; {href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/us-south-3,name:us-south-3}"
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
    "id": "ibm+us-south+cxf-2x4",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "ibm+us-south+cxf-2x4",
    "namespace": "system",
    "providerName": "ibm",
    "regionLatitude": 32.81248,
    "regionLongitude": -96.77619,
    "regionName": "us-south",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbvcikbk74fniqbkppd8",
    "vCPU": 2
  },
  "specId": "ibm+us-south+cxf-2x4",
  "specValidation": {
    "cspResourceId": "cxf-2x4",
    "isAvailable": true,
    "resourceId": "ibm+us-south+cxf-2x4",
    "resourceName": "cxf-2x4",
    "status": "Available"
  },
  "status": "OK",
  "suggestedZone": "us-south-1"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 22.217s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.5.0.0/16",
  "connectionName": "ibm-us-south",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-ibm",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.5.1.0/24",
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
  "cidrBlock": "10.5.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T09:54:56Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:56Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:56Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "r006-04491d9f-e847-49be-9944-9448e3f8b10a",
  "cspResourceName": "tbm6nu8dndcaujb0scli",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-ibm",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "ClassicAccess",
      "value": "false"
    },
    {
      "key": "CreatedAt",
      "value": "2026-11-25T09:54:42.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-04491d9f-e847-49be-9944-9448e3f8b10a"
    },
    {
      "key": "CseSourceIps",
      "value": "{ip:{address:10.12.123.100},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}}; {ip:{address:10.12.156.226},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-2,name:us-south-2}}; {ip:{address:10.12.168.127},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-3,name:us-south-3}}"
    },
    {
      "key": "DefaultNetworkACL",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-d19cf538-5cfd-4de1-ae88-c016b335ee97,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-d19cf538-5cfd-4de1-ae88-c016b335ee97,id:r006-d19cf538-5cfd-4de1-ae88-c016b335ee97,name:squatter-proofing-emphases-unthawed}"
    },
    {
      "key": "DefaultRoutingTable",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-04491d9f-e847-49be-9944-9448e3f8b10a/r006-105c240f-62a9-4b28-bbf2-ec151a44461b,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-04491d9f-e847-49be-9944-9448e3f8b10a/routing_tables/r006-105c240f-62a9-4b28-bbf2-ec151a44461b,id:r006-105c240f-62a9-4b28-bbf2-ec151a44461b,name:chili-exclusion-conceded-overhand,resource_type:routing_table}"
    },
    {
      "key": "DefaultSecurityGroup",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r006-0eccd748-a3ee-42c3-b2e3-f6d50c83a711,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-0eccd748-a3ee-42c3-b2e3-f6d50c83a711,id:r006-0eccd748-a3ee-42c3-b2e3-f6d50c83a711,name:manifesto-slur-headstand-bottle}"
    },
    {
      "key": "Dns",
      "value": "{enable_hub:false,resolution_binding_count:0,resolver:{servers:[{address:161.26.0.10},{address:161.26.0.11}],type:system,configuration:default}}"
    },
    {
      "key": "HealthState",
      "value": "ok"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-04491d9f-e847-49be-9944-9448e3f8b10a"
    },
    {
      "key": "ID",
      "value": "r006-04491d9f-e847-49be-9944-9448e3f8b10a"
    },
    {
      "key": "Name",
      "value": "tbm6nu8dndcaujb0scli"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
    },
    {
      "key": "ResourceType",
      "value": "vpc"
    },
    {
      "key": "Status",
      "value": "available"
    },
    {
      "key": "AvailableIpv4AddressCount",
      "value": "251"
    },
    {
      "key": "CreatedAt",
      "value": "2026-11-25T09:54:51.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:0717-77963d09-dcf1-43e0-9201-04f89da02022"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-77963d09-dcf1-43e0-9201-04f89da02022"
    },
    {
      "key": "ID",
      "value": "0717-77963d09-dcf1-43e0-9201-04f89da02022"
    },
    {
      "key": "IPVersion",
      "value": "ipv4"
    },
    {
      "key": "Ipv4CIDRBlock",
      "value": "10.5.1.0/24"
    },
    {
      "key": "Name",
      "value": "tb57pq87p48ek35mlrq7"
    },
    {
      "key": "NetworkACL",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-d19cf538-5cfd-4de1-ae88-c016b335ee97,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-d19cf538-5cfd-4de1-ae88-c016b335ee97,id:r006-d19cf538-5cfd-4de1-ae88-c016b335ee97,name:squatter-proofing-emphases-unthawed}"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
    },
    {
      "key": "ResourceType",
      "value": "subnet"
    },
    {
      "key": "RoutingTable",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-04491d9f-e847-49be-9944-9448e3f8b10a/r006-105c240f-62a9-4b28-bbf2-ec151a44461b,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-04491d9f-e847-49be-9944-9448e3f8b10a/routing_tables/r006-105c240f-62a9-4b28-bbf2-ec151a44461b,id:r006-105c240f-62a9-4b28-bbf2-ec151a44461b,name:chili-exclusion-conceded-overhand,resource_type:routing_table}"
    },
    {
      "key": "Status",
      "value": "available"
    },
    {
      "key": "TotalIpv4AddressCount",
      "value": "256"
    },
    {
      "key": "VPC",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-04491d9f-e847-49be-9944-9448e3f8b10a,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-04491d9f-e847-49be-9944-9448e3f8b10a,id:r006-04491d9f-e847-49be-9944-9448e3f8b10a,name:tbm6nu8dndcaujb0scli,resource_type:vpc}"
    },
    {
      "key": "Zone",
      "value": "{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}"
    }
  ],
  "name": "test-rdbms-vnet-ibm",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:54:56Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:54:56Z",
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
      "cspResourceId": "0717-77963d09-dcf1-43e0-9201-04f89da02022",
      "cspResourceName": "tb57pq87p48ek35mlrq7",
      "cspVNetId": "r006-04491d9f-e847-49be-9944-9448e3f8b10a",
      "cspVNetName": "tbm6nu8dndcaujb0scli",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.5.1.0/24",
      "keyValueList": [
        {
          "key": "AvailableIpv4AddressCount",
          "value": "251"
        },
        {
          "key": "CreatedAt",
          "value": "2026-11-25T09:54:51.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:us-south-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:0717-77963d09-dcf1-43e0-9201-04f89da02022"
        },
        {
          "key": "Href",
          "value": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-77963d09-dcf1-43e0-9201-04f89da02022"
        },
        {
          "key": "ID",
          "value": "0717-77963d09-dcf1-43e0-9201-04f89da02022"
        },
        {
          "key": "IPVersion",
          "value": "ipv4"
        },
        {
          "key": "Ipv4CIDRBlock",
          "value": "10.5.1.0/24"
        },
        {
          "key": "Name",
          "value": "tb57pq87p48ek35mlrq7"
        },
        {
          "key": "NetworkACL",
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-d19cf538-5cfd-4de1-ae88-c016b335ee97,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-d19cf538-5cfd-4de1-ae88-c016b335ee97,id:r006-d19cf538-5cfd-4de1-ae88-c016b335ee97,name:squatter-proofing-emphases-unthawed}"
        },
        {
          "key": "ResourceGroup",
          "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
        },
        {
          "key": "ResourceType",
          "value": "subnet"
        },
        {
          "key": "RoutingTable",
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-04491d9f-e847-49be-9944-9448e3f8b10a/r006-105c240f-62a9-4b28-bbf2-ec151a44461b,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-04491d9f-e847-49be-9944-9448e3f8b10a/routing_tables/r006-105c240f-62a9-4b28-bbf2-ec151a44461b,id:r006-105c240f-62a9-4b28-bbf2-ec151a44461b,name:chili-exclusion-conceded-overhand,resource_type:routing_table}"
        },
        {
          "key": "Status",
          "value": "available"
        },
        {
          "key": "TotalIpv4AddressCount",
          "value": "256"
        },
        {
          "key": "VPC",
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-04491d9f-e847-49be-9944-9448e3f8b10a,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-04491d9f-e847-49be-9944-9448e3f8b10a,id:r006-04491d9f-e847-49be-9944-9448e3f8b10a,name:tbm6nu8dndcaujb0scli,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb57pq87p48ek35mlrq7",
      "zone": "us-south-1"
    }
  ],
  "systemLabel": "",
  "uid": "tbm6nu8dndcaujb0scli"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 7.156s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "ibm-us-south",
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
  "name": "test-rdbms-sg-ibm",
  "vNetId": "test-rdbms-vnet-ibm"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "r006-e9386ee2-dc76-4a9c-88d4-7e1df3e0b7df",
  "cspResourceName": "tbbijd8qlo7rudnt5vi8",
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
  "id": "test-rdbms-sg-ibm",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "CreatedAt",
      "value": "2026-11-25T09:54:58.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r006-e9386ee2-dc76-4a9c-88d4-7e1df3e0b7df"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-e9386ee2-dc76-4a9c-88d4-7e1df3e0b7df"
    },
    {
      "key": "ID",
      "value": "r006-e9386ee2-dc76-4a9c-88d4-7e1df3e0b7df"
    },
    {
      "key": "Name",
      "value": "tbbijd8qlo7rudnt5vi8"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
    },
    {
      "key": "Rules",
      "value": "{direction:inbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-e9386ee2-dc76-4a9c-88d4-7e1df3e0b7df/rules/r006-5630009d-e0fa-4246-9903-e806a3b41c92,id:r006-5630009d-e0fa-4246-9903-e806a3b41c92,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:mossy-join-architecture-tipoff,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,port_max:3306,port_min:3306,protocol:tcp}; {direction:inbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-e9386ee2-dc76-4a9c-88d4-7e1df3e0b7df/rules/r006-e2131fdc-d2ac-47dc-9c40-0f3abeba92ec,id:r006-e2131fdc-d2ac-47dc-9c40-0f3abeba92ec,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:dispersed-engraved-mouse-grinning,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,port_max:22,port_min:22,protocol:tcp}; {direction:outbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-e9386ee2-dc76-4a9c-88d4-7e1df3e0b7df/rules/r006-18bca974-c5f4-4a91-8e1d-dfc05ef62209,id:r006-18bca974-c5f4-4a91-8e1d-dfc05ef62209,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:rinse-sarcastic-quit-rubble,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,protocol:any}"
    },
    {
      "key": "VPC",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-04491d9f-e847-49be-9944-9448e3f8b10a,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-04491d9f-e847-49be-9944-9448e3f8b10a,id:r006-04491d9f-e847-49be-9944-9448e3f8b10a,name:tbm6nu8dndcaujb0scli,resource_type:vpc}"
    }
  ],
  "name": "test-rdbms-sg-ibm",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbbijd8qlo7rudnt5vi8",
  "vNetId": "test-rdbms-vnet-ibm"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 7ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=ibm`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "ibm": {
      "dbOperationMethod": "sqlFallback",
      "note": "Storage type selection not supported. IBM Cloud Databases manages storage automatically. Hosting Models \u0026 Provisioning Times: - 'multitenant' (shared): Recommended default. Fast provisioning in ~9-10 minutes. - Dedicated host flavors ('b3c.*', 'm3c.*'): Require dedicated bare-metal/VPC hardware slice allocation and disk encryption, taking 35-45+ minutes. If selected, ensure client and reverse-proxy timeouts are configured for at least 50 minutes, or invoke migration asynchronously (Prefer: respond-async).",
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
- **Duration:** 6.014s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=ibm-us-south&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "NA",
    "connectionName": "ibm-us-south",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "b3c.16x64.encrypted",
      "b3c.32x128.encrypted",
      "b3c.4x16.encrypted",
      "b3c.8x32.encrypted",
      "m3c.30x240.encrypted",
      "m3c.8x64.encrypted",
      "multitenant"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "65536",
        "name": "b3c.16x64.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "b3c.32x128.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "b3c.4x16.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "b3c.8x32.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "245760",
        "name": "m3c.30x240.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "30"
      },
      {
        "memSizeMiB": "65536",
        "name": "m3c.8x64.encrypted",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "12288",
        "name": "multitenant",
        "storageSizeRangeGB": {
          "max": 13194,
          "min": 32
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "0"
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
          "constraints": "Minimum 32GB storage.",
          "description": "Storage is managed automatically by IBM Cloud Databases. User cannot specify storage type.",
          "displayName": "Automatic (IBM-managed)",
          "maxSizeGB": 13194,
          "minSizeGB": 32,
          "recommendationLevel": "standard",
          "storageType": "NA"
        }
      ]
    },
    "providerName": "ibm",
    "regionName": "us-south",
    "requiresSecurityGroup": false,
    "requiresSubnet": false,
    "storageSizeRangeGB": {
      "max": 13194,
      "min": 32
    },
    "storageTypeOptions": [
      "NA"
    ],
    "supportedVersions": [
      "8.4"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": true,
    "supportsEncryption": true,
    "supportsHighAvailability": true,
    "supportsPublicAccess": true,
    "supportsStorageSizeConfiguration": true,
    "supportsStorageTypeSelection": false,
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
    "csp": "ibm",
    "region": "us-south"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for ibm (us-south)",
  "status": "recommended",
  "targetCloud": {
    "csp": "ibm",
    "region": "us-south"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "admin",
      "adminUserPassword": "******",
      "databases": [
        {
          "charset": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "multitenant",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-ibm",
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-ibm"
    }
  ],
  "warnings": [
    "Requested mysql version '8.0' is not directly supported on target cloud; recommended closest available version '8.4' (supported versions: [8.4]).",
    "IBM Cloud Databases 'multitenant' hosting model was recommended for instance 'Source MySQL 01' for fast provisioning (~10 minutes). Dedicated host flavors (e.g., 'b3c.4x16.encrypted') take 35~45+ minutes to provision.",
    "Storage type selection is not configurable on target cloud (ibm); requested storage type 'SSD' for instance 'Source MySQL 01' will be managed automatically by the provider."
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 28ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "admin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "ibm-us-south",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.4",
  "dbInstanceSpec": "multitenant",
  "name": "rdbms-ibm",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-ibm"
  ],
  "storageSize": 100,
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-ibm"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "admin",
    "adminUserPassword": "******",
    "connectionName": "ibm-us-south",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.4",
    "dbInstanceSpec": "multitenant",
    "name": "rdbms-ibm",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-ibm"
    ],
    "storageSize": 100,
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-ibm"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 11m13.997s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for ibm (us-south)",
  "status": "recommended",
  "targetCloud": {
    "csp": "ibm",
    "region": "us-south"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "admin",
      "adminUserPassword": "******",
      "databases": [
        {
          "charset": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "multitenant",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-ibm",
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-ibm"
    }
  ],
  "warnings": [
    "Requested mysql version '8.0' is not directly supported on target cloud; recommended closest available version '8.4' (supported versions: [8.4]).",
    "IBM Cloud Databases 'multitenant' hosting model was recommended for instance 'Source MySQL 01' for fast provisioning (~10 minutes). Dedicated host flavors (e.g., 'b3c.4x16.encrypted') take 35~45+ minutes to provision.",
    "Storage type selection is not configurable on target cloud (ibm); requested storage type 'SSD' for instance 'Source MySQL 01' will be managed automatically by the provider."
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
- **Duration:** 4.14s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm`
```json
// Response Body
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
      "backupRetentionDays": 7,
      "backupTime": "AUTO",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:05:41Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
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
      "status": "Deleting",
      "storageSize": 100,
      "storageType": "Premium_LRS",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbqt11ul0jr723gfs0im",
      "vNetId": "test-rdbms-vnet-azure"
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
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 4.342s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm/database`
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
- **Duration:** 7.537s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm/database`
```json
// Response Body
{
  "databases": [
    "ibmclouddb",
    "information_schema",
    "meta",
    "mysql",
    "performance_schema",
    "sampledb",
    "sampledb_dyn",
    "sys"
  ]
}
```

### 13. Beetle GET Secure Transport Info [✅ SUCCESS]
- **Duration:** 8.559s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": true,
    "issuer": "CN=IBM Cloud Databases",
    "notAfter": "2028-06-22T14:29:00Z",
    "pem": "-----BEGIN CERTIFICATE-----\nMIIDDzCCAfegAwIBAgIJANEH58y2/kzHMA0GCSqGSIb3DQEBCwUAMB4xHDAaBgNV\nBAMME0lCTSBDbG91ZCBEYXRhYmFzZXMwHhcNMTgwNjI1MTQyOTAwWhcNMjgwNjIy\nMTQyOTAwWjAeMRwwGgYDVQQDDBNJQk0gQ2xvdWQgRGF0YWJhc2VzMIIBIjANBgkq\nhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8lpaQGzcFdGqeMlmqjffMPpIQhqpd8qJ\nPr3bIkrXJbTcJJ9uIckSUcCjw4Z/rSg8nnT13SCcOl+1to+7kdMiU8qOWKiceYZ5\ny+yZYfCkGaiZVfazQBm45zBtFWv+AB/8hfCTdNF7VY4spaA3oBE2aS7OANNSRZSK\npwy24IUgUcILJW+mcvW80Vx+GXRfD9Ytt6PRJgBhYuUBpgzvngmCMGBn+l2KNiSf\nweovYDCD6Vngl2+6W9QFAFtWXWgF3iDQD5nl/n4mripMSX6UG/n6657u7TDdgkvA\n1eKI2FLzYKpoKBe5rcnrM7nHgNc/nCdEs5JecHb1dHv1QfPm6pzIxwIDAQABo1Aw\nTjAdBgNVHQ4EFgQUK3+XZo1wyKs+DEoYXbHruwSpXjgwHwYDVR0jBBgwFoAUK3+X\nZo1wyKs+DEoYXbHruwSpXjgwDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOC\nAQEAJf5dvlzUpqaix26qJEuqFG0IP57QQI5TCRJ6Xt/supRHo63eDvKw8zR7tlWQ\nlV5P0N2xwuSl9ZqAJt7/k/3ZeB+nYwPoyO3KvKvATunRvlPBn4FWVXeaPsG+7fhS\nqsejmkyonYw77HRzGOzJH4Zg8UN6mfpbaWSsyaExvqknCp9SoTQP3D67AzWqb1zY\ndoqqgGIZ2nxCkp5/FXxF/TMb55vteTQwfgBy60jVVkbF7eVOWCv0KaNHPF5hrqbN\ni+3XjJ7/peF3xMvTMoy35DcT3E2ZeSVjouZs15O90kI3k2daS2OHJABW0vSj4nLz\n+PQzp/B9cQmOO8dCe049Q3oaUA==\n-----END CERTIFICATE-----\n",
    "subject": "CN=IBM Cloud Databases"
  },
  "enforced": false,
  "engine": "mysql",
  "recommendedSSLMode": "VERIFY_IDENTITY",
  "requireSecureTransport": "ON",
  "rules": "",
  "tlsCipher": "TLS_AES_128_GCM_SHA256",
  "tlsInUse": true
}
```

### 14. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 909ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m6.948s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 31.515s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 24.028s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 2.604s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 7.314s

