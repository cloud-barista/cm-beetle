# Managed RDBMS (MySQL) Test Report: GCP (us-central1)

- **Test Case:** GCP US-Central1 MySQL Test
- **Date & Time:** 2026-10-19 10:24:36
- **Namespace:** `default`
- **Total Duration:** 14m45.843s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** GCP
- **Target Region:** `us-central1`
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
- **Duration:** 5.037s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2404-noble-amd64-v20260820",
  "specId": "gcp+us-central1+e2-standard-2"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "instanceType": "e2-standard-2",
    "provider": "gcp",
    "queriedAt": "2026-10-19T01:24:41.723394236Z",
    "region": "us-central1",
    "source": "gcp:machineTypes.aggregatedList",
    "zones": [
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-b"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-c"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-a"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-f"
      }
    ]
  },
  "connectionName": "gcp-us-central1",
  "estimatedCost": "$0.0670/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "gcp-africa-south1",
    "creationDate": "",
    "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2404-noble-amd64-v20260820",
    "description": "Canonical, Ubuntu, 24.04 LTS, amd64 noble image built on 2026-08-20",
    "details": [
      {
        "key": "Architecture",
        "value": "X86_64"
      },
      {
        "key": "ArchiveSizeBytes",
        "value": "60138870400"
      },
      {
        "key": "CreationTimestamp",
        "value": "2026-08-20T05:48:46.989-07:00"
      },
      {
        "key": "Description",
        "value": "Canonical, Ubuntu, 24.04 LTS, amd64 noble image built on 2026-08-20"
      },
      {
        "key": "DiskSizeGb",
        "value": "10"
      },
      {
        "key": "EnableConfidentialCompute",
        "value": "false"
      },
      {
        "key": "Family",
        "value": "ubuntu-2404-lts-amd64"
      },
      {
        "key": "GuestOsFeatures",
        "value": "{type:VIRTIO_SCSI_MULTIQUEUE}; {type:SEV_CAPABLE}; {type:SEV_SNP_CAPABLE}; {type:SEV_LIVE_MIGRATABLE}; {type:SEV_LIVE_MIGRATABLE_V2}; {type:SNP_SVSM_CAPABLE}; {type:IDPF}; {type:TDX_CAPABLE}; {type:UEFI_COMPATIBLE}; {type:GVNIC}"
      },
      {
        "key": "Id",
        "value": "2634297659785693505"
      },
      {
        "key": "Kind",
        "value": "compute#image"
      },
      {
        "key": "LabelFingerprint",
        "value": "iNBmVNCFF9w="
      },
      {
        "key": "Labels",
        "value": "{public-image:true}"
      },
      {
        "key": "LicenseCodes",
        "value": "3242930272766215801"
      },
      {
        "key": "Licenses",
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-2404-lts"
      },
      {
        "key": "Name",
        "value": "ubuntu-2404-noble-amd64-v20260820"
      },
      {
        "key": "RawDisk",
        "value": "{containerType:TAR}"
      },
      {
        "key": "SatisfiesPzi",
        "value": "false"
      },
      {
        "key": "SatisfiesPzs",
        "value": "false"
      },
      {
        "key": "SelfLink",
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2404-noble-amd64-v20260820"
      },
      {
        "key": "SourceType",
        "value": "RAW"
      },
      {
        "key": "Status",
        "value": "READY"
      },
      {
        "key": "StorageLocations",
        "value": "us; asia-east1; europe-west1; asia-northeast1; asia-northeast2; asia-southeast2; asia-southeast3; us-central1; us-west3; australia-southeast2; northamerica-northeast1; us-central2; australia-southeast1; asia; europe-west8; northamerica-south1; northamerica-northeast2; asia-south1; me-central2; us-west2; me-central1; asia-northeast3; europe-north2; us-east5; us-east7; europe-southwest1; asia-east2; europe-west6; europe-west4; europe-west3; us-east4; us-east1; asia-southeast1; europe-west12; europe-north1; southamerica-west1; europe-west9; europe-west10; us-west1; us-west8; us-south1; southamerica-east1; europe-west15; us-west4; asia-south2; europe-west2; europe-central2; africa-south1; me-west1; eu"
      }
    ],
    "fetchedTime": "2026.08.21 13:58:59 Fri",
    "id": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2404-noble-amd64-v20260820",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2404-noble-amd64-v20260820",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 10,
    "osDiskType": "NA",
    "osDistribution": "Canonical, Ubuntu, 24.04 LTS, amd64 noble image built on 2026-08-20",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "gcp",
    "regionList": [
      "common"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tbs76r4cga1lqjhvc013"
  },
  "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2404-noble-amd64-v20260820",
  "imageValidation": {
    "cspResourceId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2404-noble-amd64-v20260820",
    "isAvailable": true,
    "resourceId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2404-noble-amd64-v20260820",
    "resourceName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-2404-noble-amd64-v20260820",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "gcp",
  "regionName": "us-central1",
  "specDetails": {
    "architecture": "x86_64",
    "connectionName": "gcp-us-central1",
    "costPerHour": 0.067011,
    "cspSpecName": "e2-standard-2",
    "details": [
      {
        "key": "CreationTimestamp",
        "value": "1969-12-31T16:00:00.000-08:00"
      },
      {
        "key": "Description",
        "value": "Efficient Instance, 2 vCPUs, 8 GB RAM"
      },
      {
        "key": "GuestCpus",
        "value": "2"
      },
      {
        "key": "Id",
        "value": "335002"
      },
      {
        "key": "ImageSpaceGb",
        "value": "0"
      },
      {
        "key": "IsSharedCpu",
        "value": "false"
      },
      {
        "key": "Kind",
        "value": "compute#machineType"
      },
      {
        "key": "MaximumPersistentDisks",
        "value": "128"
      },
      {
        "key": "MaximumPersistentDisksSizeGb",
        "value": "263168"
      },
      {
        "key": "MemoryMb",
        "value": "8192"
      },
      {
        "key": "Name",
        "value": "e2-standard-2"
      },
      {
        "key": "SelfLink",
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/us-central1-a/machineTypes/e2-standard-2"
      },
      {
        "key": "Zone",
        "value": "us-central1-a"
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
    "id": "gcp+us-central1+e2-standard-2",
    "infraType": "node",
    "memoryGiB": 7.8125,
    "name": "gcp+us-central1+e2-standard-2",
    "namespace": "system",
    "providerName": "gcp",
    "regionLatitude": 41.2522,
    "regionLongitude": -95.8575,
    "regionName": "us-central1",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbe0e9dnnuunhi69h1tq",
    "vCPU": 2
  },
  "specId": "gcp+us-central1+e2-standard-2",
  "specValidation": {
    "cspResourceId": "e2-standard-2",
    "isAvailable": true,
    "resourceId": "gcp+us-central1+e2-standard-2",
    "resourceName": "e2-standard-2",
    "status": "Available"
  },
  "status": "OK",
  "suggestedZone": "us-central1-b"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 31.72s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.2.0.0/16",
  "connectionName": "gcp-us-central1",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-gcp",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.2.1.0/24",
      "name": "subnet-1",
      "zone": "us-central1-a"
    }
  ]
}
```
```json
// Response Body
{
  "associatedObjectList": null,
  "cidrBlock": "10.2.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-10-19T01:25:11Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-19T01:25:11Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-10-19T01:25:11Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "tbvhmd58u08s9bt95b6i",
  "cspResourceName": "tbvhmd58u08s9bt95b6i",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-gcp",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "AutoCreateSubnetworks",
      "value": "false"
    },
    {
      "key": "CreationTimestamp",
      "value": "2026-10-18T18:24:41.871-07:00"
    },
    {
      "key": "EnableUlaInternalIpv6",
      "value": "false"
    },
    {
      "key": "Id",
      "value": "7187400942776012726"
    },
    {
      "key": "Kind",
      "value": "compute#network"
    },
    {
      "key": "Mtu",
      "value": "0"
    },
    {
      "key": "Name",
      "value": "tbvhmd58u08s9bt95b6i"
    },
    {
      "key": "NetworkFirewallPolicyEnforcementOrder",
      "value": "AFTER_CLASSIC_FIREWALL"
    },
    {
      "key": "RoutingConfig",
      "value": "{bgpBestPathSelectionMode:LEGACY,routingMode:REGIONAL}"
    },
    {
      "key": "SelfLink",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbvhmd58u08s9bt95b6i"
    },
    {
      "key": "SelfLinkWithId",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/7187400942776012726"
    },
    {
      "key": "Subnetworks",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1/subnetworks/tbn40jpit6tig8keji81"
    }
  ],
  "name": "test-rdbms-vnet-gcp",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T01:25:11Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T01:25:11Z",
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
      "cspResourceId": "tbn40jpit6tig8keji81",
      "cspResourceName": "tbn40jpit6tig8keji81",
      "cspVNetId": "tbvhmd58u08s9bt95b6i",
      "cspVNetName": "tbvhmd58u08s9bt95b6i",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.2.1.0/24",
      "keyValueList": [
        {
          "key": "AllowSubnetCidrRoutesOverlap",
          "value": "false"
        },
        {
          "key": "CreationTimestamp",
          "value": "2026-10-18T18:24:58.368-07:00"
        },
        {
          "key": "EnableFlowLogs",
          "value": "false"
        },
        {
          "key": "Fingerprint",
          "value": "-X7tJ8bMdcU="
        },
        {
          "key": "GatewayAddress",
          "value": "10.2.1.1"
        },
        {
          "key": "Id",
          "value": "6823187256513712005"
        },
        {
          "key": "IpCidrRange",
          "value": "10.2.1.0/24"
        },
        {
          "key": "Kind",
          "value": "compute#subnetwork"
        },
        {
          "key": "Name",
          "value": "tbn40jpit6tig8keji81"
        },
        {
          "key": "Network",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbvhmd58u08s9bt95b6i"
        },
        {
          "key": "PrivateIpGoogleAccess",
          "value": "false"
        },
        {
          "key": "PrivateIpv6GoogleAccess",
          "value": "DISABLE_GOOGLE_ACCESS"
        },
        {
          "key": "Purpose",
          "value": "PRIVATE"
        },
        {
          "key": "Region",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1"
        },
        {
          "key": "SelfLink",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1/subnetworks/tbn40jpit6tig8keji81"
        },
        {
          "key": "StackType",
          "value": "IPV4_ONLY"
        },
        {
          "key": "region",
          "value": "us-central1"
        },
        {
          "key": "subnet",
          "value": "tbn40jpit6tig8keji81"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbn40jpit6tig8keji81",
      "zone": "us-central1-a"
    }
  ],
  "systemLabel": "",
  "uid": "tbvhmd58u08s9bt95b6i"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 34.968s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "gcp-us-central1",
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
  "name": "test-rdbms-sg-gcp",
  "vNetId": "test-rdbms-vnet-gcp"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "tbeu3spdsseb2ir2ff7r",
  "cspResourceName": "tbeu3spdsseb2ir2ff7r",
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
  "id": "test-rdbms-sg-gcp",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "Items",
      "value": "{allowed:[{IPProtocol:tcp,ports:[3306]}],creationTimestamp:2026-10-18T18:25:28.854-07:00,direction:INGRESS,id:2611272614994825063,kind:compute#firewall,logConfig:{},name:tbeu3spdsseb2ir2ff7r-i-001,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbvhmd58u08s9bt95b6i,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tbeu3spdsseb2ir2ff7r-i-001,sourceRanges:[0.0.0.0/0],targetTags:[tbeu3spdsseb2ir2ff7r]}; {allowed:[{IPProtocol:tcp,ports:[22]}],creationTimestamp:2026-10-18T18:25:35.435-07:00,direction:INGRESS,id:4704689453736185696,kind:compute#firewall,logConfig:{},name:tbeu3spdsseb2ir2ff7r-i-002,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbvhmd58u08s9bt95b6i,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tbeu3spdsseb2ir2ff7r-i-002,sourceRanges:[0.0.0.0/0],targetTags:[tbeu3spdsseb2ir2ff7r]}; {creationTimestamp:2026-10-18T18:25:11.838-07:00,denied:[{IPProtocol:all}],destinationRanges:[0.0.0.0/0],direction:EGRESS,id:7832559593899741080,kind:compute#firewall,logConfig:{},name:tbeu3spdsseb2ir2ff7r-o-001,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbvhmd58u08s9bt95b6i,priority:65535,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tbeu3spdsseb2ir2ff7r-o-001,targetTags:[tbeu3spdsseb2ir2ff7r]}; {allowed:[{IPProtocol:all}],creationTimestamp:2026-10-18T18:25:22.178-07:00,destinationRanges:[0.0.0.0/0],direction:EGRESS,id:3037403760977039214,kind:compute#firewall,logConfig:{},name:tbeu3spdsseb2ir2ff7r-o-002,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbvhmd58u08s9bt95b6i,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tbeu3spdsseb2ir2ff7r-o-002,targetTags:[tbeu3spdsseb2ir2ff7r]}"
    }
  ],
  "name": "test-rdbms-sg-gcp",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbeu3spdsseb2ir2ff7r",
  "vNetId": "test-rdbms-vnet-gcp"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=gcp`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "gcp": {
      "dbOperationMethod": "cspNativeApi",
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
- **Duration:** 2.424s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=gcp-us-central1&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "1-7",
    "connectionName": "gcp-us-central1",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "db-c4a-highmem-16",
      "db-c4a-highmem-2",
      "db-c4a-highmem-32",
      "db-c4a-highmem-4",
      "db-c4a-highmem-48",
      "db-c4a-highmem-64",
      "db-c4a-highmem-72",
      "db-c4a-highmem-8",
      "db-f1-micro",
      "db-g1-small",
      "db-memory-optimized-N-16",
      "db-memory-optimized-N-4",
      "db-memory-optimized-N-8",
      "db-n1-highmem-16",
      "db-n1-highmem-2",
      "db-n1-highmem-32",
      "db-n1-highmem-4",
      "db-n1-highmem-64",
      "db-n1-highmem-8",
      "db-n1-highmem-96",
      "db-n1-standard-1",
      "db-n1-standard-16",
      "db-n1-standard-2",
      "db-n1-standard-32",
      "db-n1-standard-4",
      "db-n1-standard-64",
      "db-n1-standard-8",
      "db-n1-standard-96",
      "db-perf-optimized-N-128",
      "db-perf-optimized-N-16",
      "db-perf-optimized-N-2",
      "db-perf-optimized-N-32",
      "db-perf-optimized-N-4",
      "db-perf-optimized-N-48",
      "db-perf-optimized-N-64",
      "db-perf-optimized-N-8",
      "db-perf-optimized-N-80",
      "db-perf-optimized-N-96"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "131072",
        "name": "db-c4a-highmem-16",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "16384",
        "name": "db-c4a-highmem-2",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "262144",
        "name": "db-c4a-highmem-32",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "32768",
        "name": "db-c4a-highmem-4",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db-c4a-highmem-48",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db-c4a-highmem-64",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "589824",
        "name": "db-c4a-highmem-72",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "72"
      },
      {
        "memSizeMiB": "65536",
        "name": "db-c4a-highmem-8",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "614",
        "name": "db-f1-micro",
        "storageSizeRangeGB": {
          "max": 3279,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "1741",
        "name": "db-g1-small",
        "storageSizeRangeGB": {
          "max": 3279,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "524288",
        "name": "db-memory-optimized-N-16",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "131072",
        "name": "db-memory-optimized-N-4",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "262144",
        "name": "db-memory-optimized-N-8",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "106496",
        "name": "db-n1-highmem-16",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "13312",
        "name": "db-n1-highmem-2",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "212992",
        "name": "db-n1-highmem-32",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "26624",
        "name": "db-n1-highmem-4",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "425984",
        "name": "db-n1-highmem-64",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "53248",
        "name": "db-n1-highmem-8",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "638976",
        "name": "db-n1-highmem-96",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "3840",
        "name": "db-n1-standard-1",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "61440",
        "name": "db-n1-standard-16",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "7680",
        "name": "db-n1-standard-2",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "122880",
        "name": "db-n1-standard-32",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "15360",
        "name": "db-n1-standard-4",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "245760",
        "name": "db-n1-standard-64",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "30720",
        "name": "db-n1-standard-8",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "368640",
        "name": "db-n1-standard-96",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "884736",
        "name": "db-perf-optimized-N-128",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "131072",
        "name": "db-perf-optimized-N-16",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "16384",
        "name": "db-perf-optimized-N-2",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "262144",
        "name": "db-perf-optimized-N-32",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "32768",
        "name": "db-perf-optimized-N-4",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "393216",
        "name": "db-perf-optimized-N-48",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "524288",
        "name": "db-perf-optimized-N-64",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "65536",
        "name": "db-perf-optimized-N-8",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "655360",
        "name": "db-perf-optimized-N-80",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
      },
      {
        "memSizeMiB": "786432",
        "name": "db-perf-optimized-N-96",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": -1
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "-1"
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
          "constraints": "Minimum 20GB storage. Only available on machine series: C4A, N4.",
          "description": "High-performance next-generation storage. Minimum 20 GB. Only available on C4A (db-c4a-highmem-*) and N4 (db-custom-N4-*) machine series.",
          "displayName": "Hyperdisk Balanced",
          "maxSizeGB": 70369,
          "minSizeGB": 20,
          "recommendationLevel": "premium",
          "storageType": "HYPERDISK_BALANCED"
        },
        {
          "constraints": "Minimum 10GB storage.",
          "description": "Cost-effective HDD storage for less I/O-intensive workloads. Available on Shared/Dedicated core instances.",
          "displayName": "Persistent Disk HDD",
          "maxSizeGB": 70369,
          "minSizeGB": 10,
          "recommendationLevel": "standard",
          "storageType": "PD_HDD"
        },
        {
          "constraints": "Minimum 10GB storage.",
          "description": "Standard SSD storage. Automatically selected for N2 machine series (db-perf-optimized-N-*) and Shared/Dedicated core instances.",
          "displayName": "Persistent Disk SSD",
          "maxSizeGB": 70369,
          "minSizeGB": 10,
          "recommendationLevel": "standard",
          "storageType": "PD_SSD"
        }
      ]
    },
    "providerName": "gcp",
    "regionName": "us-central1",
    "requiresSecurityGroup": false,
    "requiresSubnet": false,
    "storageSizeRangeGB": {
      "max": 70369,
      "min": 10
    },
    "storageTypeOptions": [
      "HYPERDISK_BALANCED",
      "PD_HDD",
      "PD_SSD"
    ],
    "supportedVersions": [
      "5.1",
      "5.5",
      "5.6",
      "5.7",
      "8.0",
      "8.0.18",
      "8.0.26",
      "8.0.27",
      "8.0.28",
      "8.0.29",
      "8.0.30",
      "8.0.31",
      "8.0.32",
      "8.0.33",
      "8.0.34",
      "8.0.35",
      "8.0.36",
      "8.0.37",
      "8.0.39",
      "8.0.40",
      "8.0.41",
      "8.0.42",
      "8.0.43",
      "8.0.44",
      "8.0.45",
      "8.0.46",
      "8.4",
      "9.7"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": true,
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
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "gcp",
    "region": "us-central1"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for gcp (us-central1)",
  "status": "recommended",
  "targetCloud": {
    "csp": "gcp",
    "region": "us-central1"
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-gcp",
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "PD_SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-gcp"
    }
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 11ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "admin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "gcp-us-central1",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "db-n1-standard-2",
  "name": "rdbms-gcp",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-gcp"
  ],
  "storageSize": 100,
  "storageType": "PD_SSD",
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-gcp"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "admin",
    "adminUserPassword": "******",
    "connectionName": "gcp-us-central1",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0",
    "dbInstanceSpec": "db-n1-standard-2",
    "name": "rdbms-gcp",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-gcp"
    ],
    "storageSize": 100,
    "storageType": "PD_SSD",
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-gcp"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 3m15.011s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for gcp (us-central1)",
  "status": "recommended",
  "targetCloud": {
    "csp": "gcp",
    "region": "us-central1"
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-gcp",
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "PD_SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-gcp"
    }
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
- **Duration:** 16ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "19:00",
  "conditions": [
    {
      "lastTransitionTime": "2026-10-19T01:28:43Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-19T01:28:43Z",
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
  "cspResourceId": "tb252llfvohcv95drpgr",
  "cspResourceName": "tb252llfvohcv95drpgr",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "db-n1-standard-2",
  "dbInstanceType": "ZONAL",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "encryption": true,
  "endpoint": "34.71.200.102:3306",
  "highAvailability": false,
  "id": "test-rdbms-gcp",
  "name": "test-rdbms-gcp",
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-gcp"
  ],
  "status": "Available",
  "storageSize": 100,
  "storageType": "PD_SSD",
  "subnetIds": [
    "subnet-1"
  ],
  "uid": "tb252llfvohcv95drpgr",
  "vNetId": "test-rdbms-vnet-gcp"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 11ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "backupTime": "19:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T01:28:43Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T01:28:43Z",
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
      "cspResourceId": "tb252llfvohcv95drpgr",
      "cspResourceName": "tb252llfvohcv95drpgr",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "dbInstanceType": "ZONAL",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "encryption": true,
      "endpoint": "34.71.200.102:3306",
      "highAvailability": false,
      "id": "test-rdbms-gcp",
      "name": "test-rdbms-gcp",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "PD_SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tb252llfvohcv95drpgr",
      "vNetId": "test-rdbms-vnet-gcp"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-10-19T01:25:25Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-19T01:25:25Z",
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
      "uid": "tbkn302rtmvu4dtvqsu5",
      "vNetId": "test-rdbms-vnet-ibm"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 2.087s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/database`
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
- **Duration:** 1.772s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/database`
```json
// Response Body
{
  "databases": [
    "mysql",
    "information_schema",
    "performance_schema",
    "sys",
    "sampledb",
    "sampledb_dyn"
  ]
}
```

### 13. Beetle GET Secure Transport Info [✅ SUCCESS]
- **Duration:** 4.467s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": false,
    "issuer": "CN=Cloud SQL Server CA,O=Google\\, Inc,C=US,2.5.4.46=#132437303265366336622d316537662d343839352d386261612d613666326435323339333433",
    "notAfter": "2036-10-16T01:26:30Z",
    "pem": "-----BEGIN CERTIFICATE-----\nMIIDkDCCAnigAwIBAgIEe5SO+zANBgkqhkiG9w0BAQsFADBwMS0wKwYDVQQuEyQ3\nMDJlNmM2Yi0xZTdmLTQ4OTUtOGJhYS1hNmYyZDUyMzkzNDMxHDAaBgNVBAMTE0Ns\nb3VkIFNRTCBTZXJ2ZXIgQ0ExFDASBgNVBAoTC0dvb2dsZSwgSW5jMQswCQYDVQQG\nEwJVUzAeFw0yNjEwMTkwMTI1MzBaFw0zNjEwMTYwMTI2MzBaME0xKDAmBgNVBAMT\nH3lra2ltLWV0cmk6dGIyNTJsbGZ2b2hjdjk1ZHJwZ3IxFDASBgNVBAoTC0dvb2ds\nZSwgSW5jMQswCQYDVQQGEwJVUzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBAKIU+yP7NLKRcAOqD9RucaMY6agE8O10wrif1A5KUxE0/sLrfTZEGvAQqr9w\n1/XtNlVMz5cO8KnkKpl08SvzoHAYi8ZwMAZ2TyUnQQstJ/yRgPCeiGJrstPKvjAF\nRjXDBBQMqwNsfyARETmiovry/5maIqn11UT9yPQ0IaNdFmnC8TgU+e9/qls87VSy\n1PO4+cSFjJk5dxDGthaZTAwrMiUbHD6LAVd8vyok3R4fj3tEU7bKbsjrboilT99G\n4YoWQkp8V6jh4BnV0qqak4v1OONMTXBW1jkE4lmnD81aCPKOPMYSx22xdtTSlTR/\nSmGmvNb5tgjbNF0nuQ2j/FQQzU0CAwEAAaNVMFMwCQYDVR0TBAIwADBGBgNVHREE\nPzA9gjsxLWM5Y2NhNGMwLTZkMmQtNDMwNi1hNWQxLWYzMGYwZDg1NmE3NS51cy1j\nZW50cmFsMS5zcWwuZ29vZzANBgkqhkiG9w0BAQsFAAOCAQEADupsLlC5C5DsENDG\nwBaq//3vzTaBbocKx3FMf5YnuWsxlxRFYqYBVOw5sw91rNqbIINlR7snbDD5cwo/\n3EwcdPJXASndFRwQ6N6u12JCTQrc7QjGgPjv/gFUkJdpMUqWQA+OdcDrpbNNG8dv\nB/DXPpbrTAo0bNBLbUw9ftj9VRxY93PM8hB8TJZnMiCN3xdjC6ButjbzY/KaU2r9\nKDh0BV/HMoZ+wVKTRg6RPYvkLVaUQZ4UtXcRf8+r1/D3QMigpDKgHJ5l+8LohF1O\nBu8qmuxE25BOZTjPpPcgIJhQ2VtzWTqwtmDqvjJTmkGaXuc6Lfnr1nPJwGWNl1ZW\nNr0PEA==\n-----END CERTIFICATE-----\n",
    "subject": "CN=ykkim-etri:tb252llfvohcv95drpgr,O=Google\\, Inc,C=US"
  },
  "enforced": false,
  "engine": "mysql",
  "recommendedSSLMode": "VERIFY_IDENTITY",
  "requireSecureTransport": "OFF",
  "rules": "",
  "tlsCipher": "ECDHE-RSA-AES128-GCM-SHA256",
  "tlsInUse": true
}
```

### 14. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 2.278s
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m32.571s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 12.665s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 3m0.326s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 35.069s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 1m45.401s

