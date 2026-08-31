# Managed RDBMS Test Report: IBM (us-south)

- **Test Case:** IBM US-South MySQL Test
- **Date & Time:** 2026-08-31 16:38:03
- **Namespace:** `default`
- **Total Duration:** 24m22.434s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** IBM
- **Target Region:** `us-south`
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
- **Duration:** 4.091s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
  "specId": "ibm+us-south+cxf-2x4"
}
```
```json
// Response Body
{
  "connectionName": "ibm-us-south",
  "estimatedCost": "$0.0850/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "ibm-us-south",
    "creationDate": "",
    "cspImageName": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
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
        "value": "2026-07-21T05:06:10.000Z"
      },
      {
        "key": "CRN",
        "value": "crn:v1:bluemix:public:is:us-south:a/811f8abfbd32425597dc7ba40da98fa6::image:r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da"
      },
      {
        "key": "Encryption",
        "value": "none"
      },
      {
        "key": "File",
        "value": "{checksums:{sha256:576fcfb94804e51dd910b51dba7925c078f00a2a268912ae1120af5bca05e4a1},size:2}"
      },
      {
        "key": "Href",
        "value": "https://us-south.iaas.cloud.ibm.com/v1/images/r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da"
      },
      {
        "key": "ID",
        "value": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da"
      },
      {
        "key": "MinimumProvisionedSize",
        "value": "10"
      },
      {
        "key": "Name",
        "value": "ibm-ubuntu-24-04-4-minimal-amd64-6"
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
    "fetchedTime": "2026.08.21 14:00:57 Fri",
    "id": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": -1,
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
    "uid": "tbfri1sv9qegh0ro7vmr"
  },
  "imageId": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
  "imageValidation": {
    "cspResourceId": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "isAvailable": true,
    "resourceId": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "resourceName": "r006-36c4e271-037a-4ad0-94f0-f0ae11cc79da",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "ibm",
  "regionName": "us-south",
  "specDetails": {
    "architecture": "x86_64",
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
        "value": "{max:1,min:1,type:range}"
      },
      {
        "key": "NetworkBandwidthMode",
        "value": "{type:fixed,value:divided}"
      },
      {
        "key": "NetworkInterfaceCount",
        "value": "{max:1,min:1,type:range}"
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
        "value": "{default:false,type:enum,values:[false]}"
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
    "uid": "tbfa0ks787p6sdeq4b30",
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
  "status": "OK"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 12.932s
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
      "lastTransitionTime": "2026-08-31T07:38:20Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:20Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:20Z",
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
  "cspResourceId": "r006-54aac5d5-2459-4804-9ca3-cc92b302dda2",
  "cspResourceName": "tbb9ptjau06mujh1jkjk",
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
      "value": "2026-08-31T07:38:10.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-54aac5d5-2459-4804-9ca3-cc92b302dda2"
    },
    {
      "key": "CseSourceIps",
      "value": "{ip:{address:10.12.125.20},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}}; {ip:{address:10.12.156.162},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-2,name:us-south-2}}; {ip:{address:10.249.214.73},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-3,name:us-south-3}}"
    },
    {
      "key": "DefaultNetworkACL",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-55859005-a10a-48b9-bd8c-2f2a0bab3fa3,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-55859005-a10a-48b9-bd8c-2f2a0bab3fa3,id:r006-55859005-a10a-48b9-bd8c-2f2a0bab3fa3,name:stallion-resample-tamper-battle}"
    },
    {
      "key": "DefaultRoutingTable",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-54aac5d5-2459-4804-9ca3-cc92b302dda2/r006-43f88f22-083f-4b13-adc5-609a008df263,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-54aac5d5-2459-4804-9ca3-cc92b302dda2/routing_tables/r006-43f88f22-083f-4b13-adc5-609a008df263,id:r006-43f88f22-083f-4b13-adc5-609a008df263,name:citation-exhale-washroom-occupant,resource_type:routing_table}"
    },
    {
      "key": "DefaultSecurityGroup",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r006-e1e532d6-274d-41a8-8f58-135b8ba3d5ae,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-e1e532d6-274d-41a8-8f58-135b8ba3d5ae,id:r006-e1e532d6-274d-41a8-8f58-135b8ba3d5ae,name:ninth-panoramic-cane-unify}"
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
      "value": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-54aac5d5-2459-4804-9ca3-cc92b302dda2"
    },
    {
      "key": "ID",
      "value": "r006-54aac5d5-2459-4804-9ca3-cc92b302dda2"
    },
    {
      "key": "Name",
      "value": "tbb9ptjau06mujh1jkjk"
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
      "value": "2026-08-31T07:38:17.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:0717-5da098da-da63-4e5f-8072-11df4a6e62c9"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-5da098da-da63-4e5f-8072-11df4a6e62c9"
    },
    {
      "key": "ID",
      "value": "0717-5da098da-da63-4e5f-8072-11df4a6e62c9"
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
      "value": "tb1mphvtp8karrqmpme7"
    },
    {
      "key": "NetworkACL",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-55859005-a10a-48b9-bd8c-2f2a0bab3fa3,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-55859005-a10a-48b9-bd8c-2f2a0bab3fa3,id:r006-55859005-a10a-48b9-bd8c-2f2a0bab3fa3,name:stallion-resample-tamper-battle}"
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
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-54aac5d5-2459-4804-9ca3-cc92b302dda2/r006-43f88f22-083f-4b13-adc5-609a008df263,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-54aac5d5-2459-4804-9ca3-cc92b302dda2/routing_tables/r006-43f88f22-083f-4b13-adc5-609a008df263,id:r006-43f88f22-083f-4b13-adc5-609a008df263,name:citation-exhale-washroom-occupant,resource_type:routing_table}"
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
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-54aac5d5-2459-4804-9ca3-cc92b302dda2,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-54aac5d5-2459-4804-9ca3-cc92b302dda2,id:r006-54aac5d5-2459-4804-9ca3-cc92b302dda2,name:tbb9ptjau06mujh1jkjk,resource_type:vpc}"
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
          "lastTransitionTime": "2026-08-31T07:38:20Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-08-31T07:38:20Z",
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
      "cspResourceId": "0717-5da098da-da63-4e5f-8072-11df4a6e62c9",
      "cspResourceName": "tb1mphvtp8karrqmpme7",
      "cspVNetId": "r006-54aac5d5-2459-4804-9ca3-cc92b302dda2",
      "cspVNetName": "tbb9ptjau06mujh1jkjk",
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
          "value": "2026-08-31T07:38:17.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:us-south-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:0717-5da098da-da63-4e5f-8072-11df4a6e62c9"
        },
        {
          "key": "Href",
          "value": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-5da098da-da63-4e5f-8072-11df4a6e62c9"
        },
        {
          "key": "ID",
          "value": "0717-5da098da-da63-4e5f-8072-11df4a6e62c9"
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
          "value": "tb1mphvtp8karrqmpme7"
        },
        {
          "key": "NetworkACL",
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-55859005-a10a-48b9-bd8c-2f2a0bab3fa3,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-55859005-a10a-48b9-bd8c-2f2a0bab3fa3,id:r006-55859005-a10a-48b9-bd8c-2f2a0bab3fa3,name:stallion-resample-tamper-battle}"
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
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-54aac5d5-2459-4804-9ca3-cc92b302dda2/r006-43f88f22-083f-4b13-adc5-609a008df263,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-54aac5d5-2459-4804-9ca3-cc92b302dda2/routing_tables/r006-43f88f22-083f-4b13-adc5-609a008df263,id:r006-43f88f22-083f-4b13-adc5-609a008df263,name:citation-exhale-washroom-occupant,resource_type:routing_table}"
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
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-54aac5d5-2459-4804-9ca3-cc92b302dda2,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-54aac5d5-2459-4804-9ca3-cc92b302dda2,id:r006-54aac5d5-2459-4804-9ca3-cc92b302dda2,name:tbb9ptjau06mujh1jkjk,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb1mphvtp8karrqmpme7",
      "zone": "us-south-1"
    }
  ],
  "systemLabel": "",
  "uid": "tbb9ptjau06mujh1jkjk"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 5.505s
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
  "cspResourceId": "r006-0c8424b8-bd1c-4f40-aa8b-22676b3b683c",
  "cspResourceName": "tblv0hd5v44mf2simt7u",
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
      "value": "2026-08-31T07:38:23.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r006-0c8424b8-bd1c-4f40-aa8b-22676b3b683c"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-0c8424b8-bd1c-4f40-aa8b-22676b3b683c"
    },
    {
      "key": "ID",
      "value": "r006-0c8424b8-bd1c-4f40-aa8b-22676b3b683c"
    },
    {
      "key": "Name",
      "value": "tblv0hd5v44mf2simt7u"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
    },
    {
      "key": "Rules",
      "value": "{direction:inbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-0c8424b8-bd1c-4f40-aa8b-22676b3b683c/rules/r006-412fa327-64d8-42a7-b9cc-b6fa8e2c8c0c,id:r006-412fa327-64d8-42a7-b9cc-b6fa8e2c8c0c,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:county-safeguard-pamperer-trash,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,port_max:3306,port_min:3306,protocol:tcp}; {direction:inbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-0c8424b8-bd1c-4f40-aa8b-22676b3b683c/rules/r006-aa49e808-6302-4757-bbcf-7f1eb0ef7e51,id:r006-aa49e808-6302-4757-bbcf-7f1eb0ef7e51,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:frisk-bootlace-salvage-doormat,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,port_max:22,port_min:22,protocol:tcp}; {direction:outbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-0c8424b8-bd1c-4f40-aa8b-22676b3b683c/rules/r006-8c86fc7a-0457-44b9-aa8f-99dc91ee0dc3,id:r006-8c86fc7a-0457-44b9-aa8f-99dc91ee0dc3,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:prudence-garden-given-duckling,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,protocol:any}"
    },
    {
      "key": "VPC",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-54aac5d5-2459-4804-9ca3-cc92b302dda2,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-54aac5d5-2459-4804-9ca3-cc92b302dda2,id:r006-54aac5d5-2459-4804-9ca3-cc92b302dda2,name:tbb9ptjau06mujh1jkjk,resource_type:vpc}"
    }
  ],
  "name": "test-rdbms-sg-ibm",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tblv0hd5v44mf2simt7u",
  "vNetId": "test-rdbms-vnet-ibm"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 6ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=ibm`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "ibm": {
      "dbOperationMethod": "sqlFallback",
      "note": "Storage type selection not supported. IBM Cloud Databases manages storage automatically.",
      "storageTypeSelectable": false,
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
- **Duration:** 7.247s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=ibm-us-south`
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
          "description": "Storage is managed automatically by IBM Cloud Databases. User cannot specify storage type.",
          "displayName": "Automatic (IBM-managed)",
          "recommendationLevel": "standard",
          "storageType": "NA"
        }
      ]
    },
    "providerName": "ibm",
    "regionName": "us-south",
    "requiresSecurityGroup": false,
    "requiresSubnet": false,
    "storageSizeRange": {
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
  "desiredCloud": {
    "csp": "ibm",
    "region": "us-south"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for ibm (us-south)",
  "warnings": [
    "Requested mysql version '8.0' could not be strictly matched. Selected supported version '8.4' for instance 'source-mysql-01'."
  ],
  "targetCloud": {
    "csp": "ibm",
    "region": "us-south"
  },
  "targetRDBMSInstances": [
    {
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "rdbmsName": "test-rdbms-ibm",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "multitenant",
      "storageSize": 100,
      "adminUserName": "admin",
      "adminUserPassword": "******",
      "vNetId": "test-rdbms-vnet-ibm",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "publicAccess": true,
      "highAvailability": false,
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
- **Duration:** 14ms
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
  "name": "test-rdbms-ibm",
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
    "name": "test-rdbms-ibm",
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
- **Duration:** 18m17.373s
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
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "multitenant",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "test-rdbms-ibm",
      "securityGroupIds": [
        "test-rdbms-sg-ibm"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-ibm"
    }
  ],
  "warnings": [
    "Requested mysql version '8.0' could not be strictly matched. Selected supported version '8.4' for instance 'source-mysql-01'."
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
- **Duration:** 4.16s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-ibm`
```json
// Response Body
{
  "resourceType": "rdbms",
  "id": "test-test-rdbms-ibm",
  "uid": "tbi6vdt2vrd1ikjkk8ap",
  "cspResourceName": "tbi6vdt2vrd1ikjkk8ap",
  "cspResourceId": "33b7ce28-70eb-47f0-a4d2-6c62ff91b582",
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
  "status": "Available",
  "conditions": [
    {
      "type": "Ready",
      "status": "True",
      "reason": "Available",
      "lastTransitionTime": "2026-08-31T07:56:43Z"
    },
    {
      "type": "Synced",
      "status": "True",
      "reason": "Available",
      "lastTransitionTime": "2026-08-31T07:56:43Z"
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
  "dbInstanceType": "NA",
  "storageType": "NA",
  "storageSize": 100,
  "adminUserName": "admin",
  "highAvailability": false,
  "backupRetentionDays": 30,
  "backupTime": "AUTO",
  "publicAccess": true,
  "deletionProtection": false,
  "encryption": true,
  "endpoint": "33b7ce28-70eb-47f0-a4d2-6c62ff91b582.blijtlfd05jdimoomdig.databases.appdomain.cloud:30665"
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
      "cspResourceName": "tbi6vdt2vrd1ikjkk8ap",
      "cspResourceId": "33b7ce28-70eb-47f0-a4d2-6c62ff91b582",
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
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:56:43Z"
        },
        {
          "type": "Synced",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:56:43Z"
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
      "dbInstanceType": "NA",
      "storageType": "NA",
      "storageSize": 100,
      "adminUserName": "admin",
      "highAvailability": false,
      "backupRetentionDays": 30,
      "backupTime": "AUTO",
      "publicAccess": true,
      "deletionProtection": false,
      "encryption": true,
      "endpoint": "33b7ce28-70eb-47f0-a4d2-6c62ff91b582.blijtlfd05jdimoomdig.databases.appdomain.cloud:30665"
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
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 7.49s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-ibm/database`
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
- **Duration:** 11.019s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-ibm/database`
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

### 13. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 481ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m55.408s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 39.699s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-ibm/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 22.895s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-ibm?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 3.864s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 10.24s

