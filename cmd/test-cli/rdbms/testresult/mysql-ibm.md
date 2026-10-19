# Managed RDBMS (MySQL) Test Report: IBM (us-south)

- **Test Case:** IBM US-South MySQL Test
- **Date & Time:** 2026-10-19 10:24:36
- **Namespace:** `default`
- **Total Duration:** 19m19.108s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** IBM
- **Target Region:** `us-south`
- **Namespace:** `default`
- **Test Date:** 2026-10-19 10:24:36

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
- **Duration:** 4.077s
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
- **Duration:** 29.13s
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
      "lastTransitionTime": "2026-10-19T01:25:07Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-19T01:25:07Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-10-19T01:25:07Z",
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
  "cspResourceId": "r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077",
  "cspResourceName": "tbq29193jp4s2kdrc7v5",
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
      "value": "2026-10-19T01:24:46.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077"
    },
    {
      "key": "CseSourceIps",
      "value": "{ip:{address:10.22.22.127},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}}; {ip:{address:10.16.240.240},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-2,name:us-south-2}}; {ip:{address:10.12.168.171},zone:{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-3,name:us-south-3}}"
    },
    {
      "key": "DefaultNetworkACL",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-affb6216-8d6e-426a-a595-4cab4b139803,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-affb6216-8d6e-426a-a595-4cab4b139803,id:r006-affb6216-8d6e-426a-a595-4cab4b139803,name:aware-void-angles-liability}"
    },
    {
      "key": "DefaultRoutingTable",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077/r006-c0747a4f-dd37-4828-bfeb-b54e558bcf94,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077/routing_tables/r006-c0747a4f-dd37-4828-bfeb-b54e558bcf94,id:r006-c0747a4f-dd37-4828-bfeb-b54e558bcf94,name:caboose-mustard-antiquely-tremor,resource_type:routing_table}"
    },
    {
      "key": "DefaultSecurityGroup",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r006-6dd162fd-b46d-4783-a26a-f76e1108c962,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-6dd162fd-b46d-4783-a26a-f76e1108c962,id:r006-6dd162fd-b46d-4783-a26a-f76e1108c962,name:sculpture-camera-uncured-elitism}"
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
      "value": "https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077"
    },
    {
      "key": "ID",
      "value": "r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077"
    },
    {
      "key": "Name",
      "value": "tbq29193jp4s2kdrc7v5"
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
      "value": "2026-10-19T01:24:55.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:0717-e5d831d0-6d55-44c8-ae67-52bd321608a7"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-e5d831d0-6d55-44c8-ae67-52bd321608a7"
    },
    {
      "key": "ID",
      "value": "0717-e5d831d0-6d55-44c8-ae67-52bd321608a7"
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
      "value": "tb28p69t6c7tdgjbaeu9"
    },
    {
      "key": "NetworkACL",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-affb6216-8d6e-426a-a595-4cab4b139803,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-affb6216-8d6e-426a-a595-4cab4b139803,id:r006-affb6216-8d6e-426a-a595-4cab4b139803,name:aware-void-angles-liability}"
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
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077/r006-c0747a4f-dd37-4828-bfeb-b54e558bcf94,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077/routing_tables/r006-c0747a4f-dd37-4828-bfeb-b54e558bcf94,id:r006-c0747a4f-dd37-4828-bfeb-b54e558bcf94,name:caboose-mustard-antiquely-tremor,resource_type:routing_table}"
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
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077,id:r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077,name:tbq29193jp4s2kdrc7v5,resource_type:vpc}"
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
          "lastTransitionTime": "2026-10-19T01:25:07Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T01:25:07Z",
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
      "cspResourceId": "0717-e5d831d0-6d55-44c8-ae67-52bd321608a7",
      "cspResourceName": "tb28p69t6c7tdgjbaeu9",
      "cspVNetId": "r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077",
      "cspVNetName": "tbq29193jp4s2kdrc7v5",
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
          "value": "2026-10-19T01:24:55.000Z"
        },
        {
          "key": "CRN",
          "value": "crn:v1:bluemix:public:is:us-south-1:a/ab205347a7c3b57f09dabb32df178bcf::subnet:0717-e5d831d0-6d55-44c8-ae67-52bd321608a7"
        },
        {
          "key": "Href",
          "value": "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-e5d831d0-6d55-44c8-ae67-52bd321608a7"
        },
        {
          "key": "ID",
          "value": "0717-e5d831d0-6d55-44c8-ae67-52bd321608a7"
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
          "value": "tb28p69t6c7tdgjbaeu9"
        },
        {
          "key": "NetworkACL",
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::network-acl:r006-affb6216-8d6e-426a-a595-4cab4b139803,href:https://us-south.iaas.cloud.ibm.com/v1/network_acls/r006-affb6216-8d6e-426a-a595-4cab4b139803,id:r006-affb6216-8d6e-426a-a595-4cab4b139803,name:aware-void-angles-liability}"
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
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc-routing-table:r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077/r006-c0747a4f-dd37-4828-bfeb-b54e558bcf94,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077/routing_tables/r006-c0747a4f-dd37-4828-bfeb-b54e558bcf94,id:r006-c0747a4f-dd37-4828-bfeb-b54e558bcf94,name:caboose-mustard-antiquely-tremor,resource_type:routing_table}"
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
          "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077,id:r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077,name:tbq29193jp4s2kdrc7v5,resource_type:vpc}"
        },
        {
          "key": "Zone",
          "value": "{href:https://us-south.iaas.cloud.ibm.com/v1/regions/us-south/zones/us-south-1,name:us-south-1}"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb28p69t6c7tdgjbaeu9",
      "zone": "us-south-1"
    }
  ],
  "systemLabel": "",
  "uid": "tbq29193jp4s2kdrc7v5"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 11.621s
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
  "cspResourceId": "r006-09ca2922-1341-4d70-942f-ea647fa5cacf",
  "cspResourceName": "tbbscs7f0l4g8konngf3",
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
      "value": "2026-10-19T01:25:12.000Z"
    },
    {
      "key": "CRN",
      "value": "crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::security-group:r006-09ca2922-1341-4d70-942f-ea647fa5cacf"
    },
    {
      "key": "Href",
      "value": "https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-09ca2922-1341-4d70-942f-ea647fa5cacf"
    },
    {
      "key": "ID",
      "value": "r006-09ca2922-1341-4d70-942f-ea647fa5cacf"
    },
    {
      "key": "Name",
      "value": "tbbscs7f0l4g8konngf3"
    },
    {
      "key": "ResourceGroup",
      "value": "{href:https://resource-controller.cloud.ibm.com/v2/resource_groups/e7c20a4f7ee64603b1c06d46b0c2385c,id:e7c20a4f7ee64603b1c06d46b0c2385c,name:default}"
    },
    {
      "key": "Rules",
      "value": "{direction:inbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-09ca2922-1341-4d70-942f-ea647fa5cacf/rules/r006-889fccb7-6275-470d-8cad-c3112f9ff902,id:r006-889fccb7-6275-470d-8cad-c3112f9ff902,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:grouped-tattered-untracked-billion,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,port_max:3306,port_min:3306,protocol:tcp}; {direction:inbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-09ca2922-1341-4d70-942f-ea647fa5cacf/rules/r006-6799b146-3cac-4b6b-927b-15624d018d30,id:r006-6799b146-3cac-4b6b-927b-15624d018d30,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:mandatory-delirium-glamorous-unclamp,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,port_max:22,port_min:22,protocol:tcp}; {direction:outbound,href:https://us-south.iaas.cloud.ibm.com/v1/security_groups/r006-09ca2922-1341-4d70-942f-ea647fa5cacf/rules/r006-21579daa-0e22-4ae3-8fe3-4fa45a39538a,id:r006-21579daa-0e22-4ae3-8fe3-4fa45a39538a,ip_version:ipv4,local:{cidr_block:0.0.0.0/0},name:paradox-undesirable-upward-judge,remote:{cidr_block:0.0.0.0/0},resource_type:security_group_rule,protocol:any}"
    },
    {
      "key": "VPC",
      "value": "{crn:crn:v1:bluemix:public:is:us-south:a/ab205347a7c3b57f09dabb32df178bcf::vpc:r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077,href:https://us-south.iaas.cloud.ibm.com/v1/vpcs/r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077,id:r006-d6dbf0ed-f955-415a-87b8-be1a70cd7077,name:tbq29193jp4s2kdrc7v5,resource_type:vpc}"
    }
  ],
  "name": "test-rdbms-sg-ibm",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbbscs7f0l4g8konngf3",
  "vNetId": "test-rdbms-vnet-ibm"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 15ms
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
- **Duration:** 6.595s
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
- **Duration:** 4ms
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
- **Duration:** 20ms
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
- **Duration:** 12m27.606s
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
- **Duration:** 5.478s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm`
```json
// Response Body
{
  "backupRetentionDays": 30,
  "backupTime": "AUTO",
  "conditions": [
    {
      "lastTransitionTime": "2026-10-19T01:36:53Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-19T01:36:53Z",
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
  "cspResourceId": "ed0b7186-1ef5-44fc-9e58-3ed59cdb6093",
  "cspResourceName": "tbkn302rtmvu4dtvqsu5",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.4",
  "dbInstanceSpec": "multitenant",
  "dbInstanceType": "NA",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "encryption": true,
  "endpoint": "ed0b7186-1ef5-44fc-9e58-3ed59cdb6093.c7dvrhud08vgdqo60090.databases.appdomain.cloud:30386",
  "highAvailability": false,
  "id": "test-rdbms-ibm",
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
  "uid": "tbkn302rtmvu4dtvqsu5",
  "vNetId": "test-rdbms-vnet-ibm"
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
      "backupRetentionDays": 30,
      "backupTime": "AUTO",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T01:36:53Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T01:36:53Z",
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
      "cspResourceId": "ed0b7186-1ef5-44fc-9e58-3ed59cdb6093",
      "cspResourceName": "tbkn302rtmvu4dtvqsu5",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "multitenant",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "encryption": true,
      "endpoint": "ed0b7186-1ef5-44fc-9e58-3ed59cdb6093.c7dvrhud08vgdqo60090.databases.appdomain.cloud:30386",
      "highAvailability": false,
      "id": "test-rdbms-ibm",
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
      "uid": "tbkn302rtmvu4dtvqsu5",
      "vNetId": "test-rdbms-vnet-ibm"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 6.963s
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
- **Duration:** 10.206s
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
- **Duration:** 12.102s
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
- **Duration:** 1.398s
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m51.486s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 47.791s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 29.689s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-ibm?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 3.813s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 11.112s

