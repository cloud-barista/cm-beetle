# Managed RDBMS Test Report: GCP (us-central1)

- **Test Case:** GCP US-Central1 MySQL Test
- **Date & Time:** 2026-08-31 16:38:03
- **Namespace:** `default`
- **Total Duration:** 13m56.546s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** GCP
- **Target Region:** `us-central1`
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
- **Duration:** 6.264s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
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
    "queriedAt": "2026-08-31T07:38:09.901446109Z",
    "region": "us-central1",
    "source": "gcp:machineTypes.aggregatedList",
    "zones": [
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-a"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-f"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-b"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-c"
      }
    ]
  },
  "connectionName": "gcp-us-central1",
  "estimatedCost": "$0.0670/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "gcp-africa-south1",
    "creationDate": "",
    "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "description": "Canonical, Ubuntu, 24.04 LTS Minimal, amd64 noble minimal image built on 2026-08-17",
    "details": [
      {
        "key": "Architecture",
        "value": "X86_64"
      },
      {
        "key": "ArchiveSizeBytes",
        "value": "50232915200"
      },
      {
        "key": "CreationTimestamp",
        "value": "2026-08-18T02:42:50.563-07:00"
      },
      {
        "key": "Description",
        "value": "Canonical, Ubuntu, 24.04 LTS Minimal, amd64 noble minimal image built on 2026-08-17"
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
        "value": "ubuntu-minimal-2404-lts-amd64"
      },
      {
        "key": "GuestOsFeatures",
        "value": "{type:VIRTIO_SCSI_MULTIQUEUE}; {type:SEV_CAPABLE}; {type:SEV_SNP_CAPABLE}; {type:SEV_LIVE_MIGRATABLE}; {type:SEV_LIVE_MIGRATABLE_V2}; {type:SNP_SVSM_CAPABLE}; {type:IDPF}; {type:TDX_CAPABLE}; {type:UEFI_COMPATIBLE}; {type:GVNIC}"
      },
      {
        "key": "Id",
        "value": "5467262061071931381"
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
        "value": "6508311393003325021"
      },
      {
        "key": "Licenses",
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/licenses/ubuntu-minimal-2404-lts"
      },
      {
        "key": "Name",
        "value": "ubuntu-minimal-2404-noble-amd64-v20260817"
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
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817"
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
        "value": "me-central2; asia-northeast3; asia-east1; us-central2; europe-north1; europe-west15; europe-southwest1; asia-northeast1; northamerica-northeast2; me-west1; australia-southeast2; us-west1; europe-west10; australia-southeast1; africa-south1; europe-west3; us-west2; asia-southeast3; europe-central2; asia-northeast2; us-central1; northamerica-south1; us-west8; us-east7; asia-east2; us-west3; southamerica-west1; europe-west12; us-west4; europe-north2; asia; us; me-central1; us-east1; europe-west2; us-south1; asia-south1; asia-south2; europe-west4; europe-west1; asia-southeast1; us-east4; europe-west9; europe-west6; southamerica-east1; northamerica-northeast1; europe-west8; us-east5; asia-southeast2; eu"
      }
    ],
    "fetchedTime": "2026.08.21 13:58:59 Fri",
    "id": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 10,
    "osDiskType": "NA",
    "osDistribution": "Canonical, Ubuntu, 24.04 LTS Minimal, amd64 noble minimal image built on 2026-08-17",
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
    "uid": "tbljac0fg9pr918114et"
  },
  "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
  "imageValidation": {
    "cspResourceId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "isAvailable": true,
    "resourceId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
    "resourceName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20260817",
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
  "suggestedZone": "us-central1-a"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 24.656s
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
      "lastTransitionTime": "2026-08-31T07:38:34Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:34Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:34Z",
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
  "cspResourceId": "tbkmnp1vc4np1p7c8ir2",
  "cspResourceName": "tbkmnp1vc4np1p7c8ir2",
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
      "value": "2026-08-31T00:38:12.031-07:00"
    },
    {
      "key": "EnableUlaInternalIpv6",
      "value": "false"
    },
    {
      "key": "Id",
      "value": "8312895456241291660"
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
      "value": "tbkmnp1vc4np1p7c8ir2"
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
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbkmnp1vc4np1p7c8ir2"
    },
    {
      "key": "SelfLinkWithId",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/8312895456241291660"
    },
    {
      "key": "Subnetworks",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1/subnetworks/tbrbup8j3vd48h2vhsoq"
    }
  ],
  "name": "test-rdbms-vnet-gcp",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-08-31T07:38:34Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-08-31T07:38:34Z",
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
      "cspResourceId": "tbrbup8j3vd48h2vhsoq",
      "cspResourceName": "tbrbup8j3vd48h2vhsoq",
      "cspVNetId": "tbkmnp1vc4np1p7c8ir2",
      "cspVNetName": "tbkmnp1vc4np1p7c8ir2",
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
          "value": "2026-08-31T00:38:23.434-07:00"
        },
        {
          "key": "EnableFlowLogs",
          "value": "false"
        },
        {
          "key": "Fingerprint",
          "value": "f-RWLJDxYnY="
        },
        {
          "key": "GatewayAddress",
          "value": "10.2.1.1"
        },
        {
          "key": "Id",
          "value": "7082648631747580288"
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
          "value": "tbrbup8j3vd48h2vhsoq"
        },
        {
          "key": "Network",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbkmnp1vc4np1p7c8ir2"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1/subnetworks/tbrbup8j3vd48h2vhsoq"
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
          "value": "tbrbup8j3vd48h2vhsoq"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbrbup8j3vd48h2vhsoq",
      "zone": "us-central1-a"
    }
  ],
  "systemLabel": "",
  "uid": "tbkmnp1vc4np1p7c8ir2"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 35.147s
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
  "cspResourceId": "tb8oauj7phftkq6gt0c5",
  "cspResourceName": "tb8oauj7phftkq6gt0c5",
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
      "value": "{allowed:[{IPProtocol:tcp,ports:[3306]}],creationTimestamp:2026-08-31T00:38:52.926-07:00,direction:INGRESS,id:6576889131399942499,kind:compute#firewall,logConfig:{},name:tb8oauj7phftkq6gt0c5-i-001,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbkmnp1vc4np1p7c8ir2,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tb8oauj7phftkq6gt0c5-i-001,sourceRanges:[0.0.0.0/0],targetTags:[tb8oauj7phftkq6gt0c5]}; {allowed:[{IPProtocol:tcp,ports:[22]}],creationTimestamp:2026-08-31T00:39:02.580-07:00,direction:INGRESS,id:8834581889323405689,kind:compute#firewall,logConfig:{},name:tb8oauj7phftkq6gt0c5-i-002,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbkmnp1vc4np1p7c8ir2,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tb8oauj7phftkq6gt0c5-i-002,sourceRanges:[0.0.0.0/0],targetTags:[tb8oauj7phftkq6gt0c5]}; {creationTimestamp:2026-08-31T00:38:36.302-07:00,denied:[{IPProtocol:all}],destinationRanges:[0.0.0.0/0],direction:EGRESS,id:1423687644813496723,kind:compute#firewall,logConfig:{},name:tb8oauj7phftkq6gt0c5-o-001,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbkmnp1vc4np1p7c8ir2,priority:65535,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tb8oauj7phftkq6gt0c5-o-001,targetTags:[tb8oauj7phftkq6gt0c5]}; {allowed:[{IPProtocol:all}],creationTimestamp:2026-08-31T00:38:44.730-07:00,destinationRanges:[0.0.0.0/0],direction:EGRESS,id:7189339759728693611,kind:compute#firewall,logConfig:{},name:tb8oauj7phftkq6gt0c5-o-002,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tbkmnp1vc4np1p7c8ir2,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tb8oauj7phftkq6gt0c5-o-002,targetTags:[tb8oauj7phftkq6gt0c5]}"
    }
  ],
  "name": "test-rdbms-sg-gcp",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tb8oauj7phftkq6gt0c5",
  "vNetId": "test-rdbms-vnet-gcp"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 3ms
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
- **Duration:** 915ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=gcp-us-central1`
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
          "maxSize": 65536,
          "minSize": 20,
          "recommendationLevel": "premium",
          "storageType": "HYPERDISK_BALANCED"
        },
        {
          "constraints": "Minimum 10GB storage.",
          "description": "Cost-effective HDD storage for less I/O-intensive workloads. Available on Shared/Dedicated core instances.",
          "displayName": "Persistent Disk HDD",
          "maxSize": 65536,
          "minSize": 10,
          "recommendationLevel": "standard",
          "storageType": "PD_HDD"
        },
        {
          "constraints": "Minimum 10GB storage.",
          "description": "Standard SSD storage. Automatically selected for N2 machine series (db-perf-optimized-N-*) and Shared/Dedicated core instances.",
          "displayName": "Persistent Disk SSD",
          "maxSize": 65536,
          "minSize": 10,
          "recommendationLevel": "standard",
          "storageType": "PD_SSD"
        }
      ]
    },
    "providerName": "gcp",
    "regionName": "us-central1",
    "requiresSecurityGroup": false,
    "requiresSubnet": false,
    "storageSizeRange": {
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
- **Duration:** 5ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "desiredCloud": {
    "csp": "gcp",
    "region": "us-central1"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for gcp (us-central1)",
  "warnings": [
    "Requested storage type 'SSD' is not supported on target cloud. Replaced with capability recommended storage for instance 'source-mysql-01'."
  ],
  "targetCloud": {
    "csp": "gcp",
    "region": "us-central1"
  },
  "targetRDBMSInstances": [
    {
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "rdbmsName": "test-rdbms-gcp",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "storageType": "PD_SSD",
      "storageSize": 100,
      "adminUserName": "admin",
      "adminUserPassword": "******",
      "vNetId": "test-rdbms-vnet-gcp",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "publicAccess": true,
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
- **Duration:** 16ms
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
  "name": "test-rdbms-gcp",
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
    "name": "test-rdbms-gcp",
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
- **Duration:** 3m29.191s
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
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "test-rdbms-gcp",
      "securityGroupIds": [
        "test-rdbms-sg-gcp"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "PD_SSD",
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-gcp"
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
- **Duration:** 4ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-gcp`
```json
// Response Body
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
  "status": "Available",
  "conditions": [
    {
      "type": "Ready",
      "status": "True",
      "reason": "Available",
      "lastTransitionTime": "2026-08-31T07:42:34Z"
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
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 6ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-alibaba",
      "uid": "tbn49stuu7oo1htp85o8",
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
      "status": "Creating",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Creating",
          "message": "RDBMS creation in progress",
          "lastTransitionTime": "2026-08-31T07:42:36Z"
        },
        {
          "type": "Synced",
          "status": "False",
          "reason": "Creating",
          "lastTransitionTime": "2026-08-31T07:42:36Z"
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
      "storageType": "cloud_auto",
      "storageSize": 100,
      "adminUserName": "dbadmin",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "publicAccess": true,
      "deletionProtection": false
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
      "status": "Creating",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Creating",
          "message": "RDBMS creation in progress",
          "lastTransitionTime": "2026-08-31T07:38:35Z"
        },
        {
          "type": "Synced",
          "status": "False",
          "reason": "Creating",
          "lastTransitionTime": "2026-08-31T07:38:35Z"
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
      "storageSize": 32,
      "adminUserName": "azureuser",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "publicAccess": true,
      "deletionProtection": false
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
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:42:34Z"
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
      "endpoint": ":3306"
    },
    {
      "resourceType": "rdbms",
      "id": "test-test-rdbms-nhn",
      "uid": "tbna7alivb64t5mbface",
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
      "status": "Creating",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Creating",
          "message": "RDBMS creation in progress",
          "lastTransitionTime": "2026-08-31T07:38:46Z"
        },
        {
          "type": "Synced",
          "status": "False",
          "reason": "Creating",
          "lastTransitionTime": "2026-08-31T07:38:46Z"
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
      "storageType": "General SSD",
      "storageSize": 100,
      "adminUserName": "myadmin",
      "highAvailability": false,
      "backupRetentionDays": 7,
      "publicAccess": true,
      "nhnDBSGToAllowAllInbound": true,
      "deletionProtection": false
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
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:42:38Z"
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
- **Duration:** 1.86s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-gcp/database`
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
- **Duration:** 2.53s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-gcp/database`
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

### 13. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 179ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m32.294s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 9.762s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-gcp/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m26.47s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-gcp?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 38.53s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 1m28.714s

