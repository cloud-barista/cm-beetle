# Managed RDBMS (MySQL) Test Report: GCP (us-central1)

- **Test Case:** GCP US-Central1 MySQL Test
- **Access Mode:** Public (publicAccess=true)
- **Date & Time:** 2026-11-25 18:54:32
- **Namespace:** `default`
- **Total Duration:** 13m37.977s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** GCP
- **Target Region:** `us-central1`
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
- **Duration:** 5.676s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20261028",
  "specId": "gcp+us-central1+e2-medium"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "cached": true,
    "instanceType": "e2-medium",
    "provider": "gcp",
    "queriedAt": "2026-11-25T09:54:35.761460299Z",
    "quota": {
      "fitCount": 2500,
      "limit": 5000,
      "name": "CPUS",
      "perInstance": 2,
      "scope": "region",
      "unit": "vCPU",
      "used": 0
    },
    "refreshAfter": "2026-11-25T09:59:35.761460299Z",
    "region": "us-central1",
    "source": "gcp:machineTypes.list",
    "status": "Available",
    "ttlSeconds": 300,
    "zones": [
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-a"
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
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-f"
      }
    ]
  },
  "connectionName": "gcp-us-central1",
  "estimatedCost": "$0.0335/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "gcp-africa-south1",
    "creationDate": "",
    "cspImageName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20261028",
    "description": "Canonical, Ubuntu, 24.04 LTS Minimal, amd64 noble minimal image built on 2026-10-28",
    "details": [
      {
        "key": "Architecture",
        "value": "X86_64"
      },
      {
        "key": "ArchiveSizeBytes",
        "value": "58764723200"
      },
      {
        "key": "CreationTimestamp",
        "value": "2026-10-28T02:12:35.133-07:00"
      },
      {
        "key": "Description",
        "value": "Canonical, Ubuntu, 24.04 LTS Minimal, amd64 noble minimal image built on 2026-10-28"
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
        "value": "7016161906646581869"
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
        "value": "ubuntu-minimal-2404-noble-amd64-v20261028"
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
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20261028"
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
        "value": "europe-central2; asia-southeast3; asia-northeast2; northamerica-south1; asia; europe-west1; southamerica-east1; us; europe-southwest1; us-east4; us-west8; europe-west9; asia-southeast2; asia-southeast1; europe-north2; europe-west3; us-central1; northamerica-northeast2; us-east5; us-east1; me-west1; asia-south2; asia-east2; us-west2; asia-northeast3; me-central1; africa-south1; asia-east1; europe-west4; australia-southeast2; northamerica-northeast1; asia-south1; europe-north1; us-east7; eu; europe-west12; us-south1; us-central2; europe-west8; us-west3; australia-southeast1; europe-west2; me-central2; europe-west6; europe-west10; southamerica-west1; asia-northeast1; us-west1; europe-west15; us-west4"
      }
    ],
    "fetchedTime": "2026.10.29 04:25:11 Thu",
    "id": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20261028",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20261028",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": 10,
    "osDiskType": "NA",
    "osDistribution": "Canonical, Ubuntu, 24.04 LTS Minimal, amd64 noble minimal image built on 2026-10-28",
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
    "uid": "tbpjnqprrb0i1ece3tp8"
  },
  "imageId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20261028",
  "imageValidation": {
    "cspResourceId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20261028",
    "isAvailable": true,
    "resourceId": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20261028",
    "resourceName": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/images/ubuntu-minimal-2404-noble-amd64-v20261028",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "gcp",
  "regionName": "us-central1",
  "specDetails": {
    "architecture": "x86_64",
    "availabilityCheckedAt": "2026-11-25T09:54:35.761460299Z",
    "availabilityQuota": {
      "fitCount": 2500,
      "limit": 5000,
      "name": "CPUS",
      "perInstance": 2,
      "scope": "region",
      "unit": "vCPU",
      "used": 0
    },
    "availabilityRefreshAfter": "2026-11-25T09:59:35.761460299Z",
    "availabilityStatus": "Available",
    "availableZones": [
      "us-central1-a",
      "us-central1-b",
      "us-central1-c",
      "us-central1-f"
    ],
    "connectionName": "gcp-us-central1",
    "costPerHour": 0.033506,
    "cspSpecName": "e2-medium",
    "details": [
      {
        "key": "CreationTimestamp",
        "value": "1969-12-31T16:00:00.000-08:00"
      },
      {
        "key": "Description",
        "value": "Efficient Instance, 2 vCPU (1/2 shared physical core) and 4 GB RAM"
      },
      {
        "key": "GuestCpus",
        "value": "2"
      },
      {
        "key": "Id",
        "value": "334004"
      },
      {
        "key": "ImageSpaceGb",
        "value": "0"
      },
      {
        "key": "IsSharedCpu",
        "value": "true"
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
        "value": "4096"
      },
      {
        "key": "Name",
        "value": "e2-medium"
      },
      {
        "key": "SelfLink",
        "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/zones/us-central1-a/machineTypes/e2-medium"
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
    "id": "gcp+us-central1+e2-medium",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "gcp+us-central1+e2-medium",
    "namespace": "system",
    "providerName": "gcp",
    "regionLatitude": 41.2522,
    "regionLongitude": -95.8575,
    "regionName": "us-central1",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbs7a4v0kojhs44rb796",
    "vCPU": 2
  },
  "specId": "gcp+us-central1+e2-medium",
  "specValidation": {
    "cspResourceId": "e2-medium",
    "isAvailable": true,
    "resourceId": "gcp+us-central1+e2-medium",
    "resourceName": "e2-medium",
    "status": "Available"
  },
  "status": "OK",
  "suggestedZone": "us-central1-a"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 29.027s
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
  "cspResourceId": "tb4qst4uuhaubd2ppqnm",
  "cspResourceName": "tb4qst4uuhaubd2ppqnm",
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
      "value": "2026-11-25T01:54:37.542-08:00"
    },
    {
      "key": "EnableUlaInternalIpv6",
      "value": "false"
    },
    {
      "key": "Id",
      "value": "8082287843120713394"
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
      "value": "tb4qst4uuhaubd2ppqnm"
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
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb4qst4uuhaubd2ppqnm"
    },
    {
      "key": "SelfLinkWithId",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/8082287843120713394"
    },
    {
      "key": "Subnetworks",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1/subnetworks/tbfu536bn06le2h1kp2a"
    }
  ],
  "name": "test-rdbms-vnet-gcp",
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
      "cspResourceId": "tbfu536bn06le2h1kp2a",
      "cspResourceName": "tbfu536bn06le2h1kp2a",
      "cspVNetId": "tb4qst4uuhaubd2ppqnm",
      "cspVNetName": "tb4qst4uuhaubd2ppqnm",
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
          "value": "2026-11-25T01:54:52.824-08:00"
        },
        {
          "key": "EnableFlowLogs",
          "value": "false"
        },
        {
          "key": "Fingerprint",
          "value": "gFIYo2BkHbo="
        },
        {
          "key": "GatewayAddress",
          "value": "10.2.1.1"
        },
        {
          "key": "Id",
          "value": "5670863409511075459"
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
          "value": "tbfu536bn06le2h1kp2a"
        },
        {
          "key": "Network",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb4qst4uuhaubd2ppqnm"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1/subnetworks/tbfu536bn06le2h1kp2a"
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
          "value": "tbfu536bn06le2h1kp2a"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbfu536bn06le2h1kp2a",
      "zone": "us-central1-a"
    }
  ],
  "systemLabel": "",
  "uid": "tb4qst4uuhaubd2ppqnm"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 34.045s
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
  "cspResourceId": "tba5eldufhj9aclqa678",
  "cspResourceName": "tba5eldufhj9aclqa678",
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
      "value": "{allowed:[{IPProtocol:tcp,ports:[3306]}],creationTimestamp:2026-11-25T01:55:22.058-08:00,direction:INGRESS,id:8447724517025582694,kind:compute#firewall,logConfig:{},name:tba5eldufhj9aclqa678-i-001,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb4qst4uuhaubd2ppqnm,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tba5eldufhj9aclqa678-i-001,sourceRanges:[0.0.0.0/0],targetTags:[tba5eldufhj9aclqa678]}; {allowed:[{IPProtocol:tcp,ports:[22]}],creationTimestamp:2026-11-25T01:55:28.979-08:00,direction:INGRESS,id:2117051222185414271,kind:compute#firewall,logConfig:{},name:tba5eldufhj9aclqa678-i-002,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb4qst4uuhaubd2ppqnm,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tba5eldufhj9aclqa678-i-002,sourceRanges:[0.0.0.0/0],targetTags:[tba5eldufhj9aclqa678]}; {creationTimestamp:2026-11-25T01:55:05.675-08:00,denied:[{IPProtocol:all}],destinationRanges:[0.0.0.0/0],direction:EGRESS,id:766536366986825366,kind:compute#firewall,logConfig:{},name:tba5eldufhj9aclqa678-o-001,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb4qst4uuhaubd2ppqnm,priority:65535,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tba5eldufhj9aclqa678-o-001,targetTags:[tba5eldufhj9aclqa678]}; {allowed:[{IPProtocol:all}],creationTimestamp:2026-11-25T01:55:13.946-08:00,destinationRanges:[0.0.0.0/0],direction:EGRESS,id:4778390873512283758,kind:compute#firewall,logConfig:{},name:tba5eldufhj9aclqa678-o-002,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb4qst4uuhaubd2ppqnm,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tba5eldufhj9aclqa678-o-002,targetTags:[tba5eldufhj9aclqa678]}"
    }
  ],
  "name": "test-rdbms-sg-gcp",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tba5eldufhj9aclqa678",
  "vNetId": "test-rdbms-vnet-gcp"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 6ms
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
- **Duration:** 981ms
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
- **Duration:** 13ms
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
- **Duration:** 34ms
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
- **Duration:** 3m12.456s
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
- **Duration:** 17ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "16:00",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T09:58:37Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T09:58:37Z",
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
  "cspResourceId": "tbc0aj584ko18830vb1j",
  "cspResourceName": "tbc0aj584ko18830vb1j",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "db-n1-standard-2",
  "dbInstanceType": "ZONAL",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "encryption": true,
  "endpoint": "136.80.5.81:3306",
  "highAvailability": false,
  "id": "test-rdbms-gcp",
  "keyValueList": [
    {
      "key": "BackendType",
      "value": "SECOND_GEN"
    },
    {
      "key": "ConnectionName",
      "value": "ykkim-etri:us-central1:tbc0aj584ko18830vb1j"
    },
    {
      "key": "CreateTime",
      "value": "2026-11-25T09:55:37.790Z"
    },
    {
      "key": "CurrentDiskSize",
      "value": "0"
    },
    {
      "key": "DatabaseInstalledVersion",
      "value": "MYSQL_8_0_46"
    },
    {
      "key": "DatabaseVersion",
      "value": "MYSQL_8_0"
    },
    {
      "key": "Etag",
      "value": "74c240921252e7aa3276e6b5dc8adb3ebefa361edcf0d4a331a9eee4d0012ea1"
    },
    {
      "key": "GceZone",
      "value": "us-central1-c"
    },
    {
      "key": "GeminiConfig",
      "value": "{flagRecommenderEnabled:true}"
    },
    {
      "key": "IncludeReplicasForMajorVersionUpgrade",
      "value": "false"
    },
    {
      "key": "InstanceType",
      "value": "CLOUD_SQL_INSTANCE"
    },
    {
      "key": "IpAddresses",
      "value": "{ipAddress:136.80.5.81,type:PRIMARY}"
    },
    {
      "key": "Kind",
      "value": "sql#instance"
    },
    {
      "key": "MaintenanceVersion",
      "value": "MYSQL_8_0_46.R20261013.02_15"
    },
    {
      "key": "MaxDiskSize",
      "value": "0"
    },
    {
      "key": "Name",
      "value": "tbc0aj584ko18830vb1j"
    },
    {
      "key": "NodeCount",
      "value": "0"
    },
    {
      "key": "Project",
      "value": "ykkim-etri"
    },
    {
      "key": "Region",
      "value": "us-central1"
    },
    {
      "key": "SatisfiesPzi",
      "value": "true"
    },
    {
      "key": "SatisfiesPzs",
      "value": "false"
    },
    {
      "key": "SelfLink",
      "value": "https://sqladmin.googleapis.com/sql/v1beta4/projects/GCP_PROJECT_ID/instances/tbc0aj584ko18830vb1j"
    },
    {
      "key": "ServerCaCert",
      "value": "{cert:-----BEGIN CERTIFICATE-----\\nMIIDcTCCAlmgAwIBAgIBADANBgkqhkiG9w0BAQsFADBwMS0wKwYDVQQuEyRhODg2\\nMWU3Ni0xYjRmLTRlZmItYmM4My0zOTZiYmZlZGEwNWIxHDAaBgNVBAMTE0Nsb3Vk\\nIFNRTCBTZXJ2ZXIgQ0ExFDASBgNVBAoTC0dvb2dsZSwgSW5jMQswCQYDVQQGEwJV\\nUzAeFw0yNjExMjUwOTU1MjZaFw0zNjExMjIwOTU2MjZaMHAxLTArBgNVBC4TJGE4\\nODYxZTc2LTFiNGYtNGVmYi1iYzgzLTM5NmJiZmVkYTA1YjEcMBoGA1UEAxMTQ2xv\\ndWQgU1FMIFNlcnZlciBDQTEUMBIGA1UEChMLR29vZ2xlLCBJbmMxCzAJBgNVBAYT\\nAlVTMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAilpdp2ZUIpAlxNob\\nmpLtJjEGEMvs0V77gI+27HmvrWkNOT3ppFhxLEnvsqdzKbVzSvD3ptPYm7zqiAoX\\nthlJ0LFmSUVw5ViP6EiTtywMiG0QNbt7UUSvU849j96MYN3h2wTJZgKm0co26Vxw\\ndZ2W4TK8WKOD5BF9vCGSkJVkpF1xPzOFl231DYm0KLrBRM4rMThSwB3ICifQiUso\\nlnIgiLZtdcEd1dXm1E8SMaAva6S7J3nd1SKpypnaCr4B3STY5wUEu8OXhDvYfzId\\nlRfEH5r0xptyIZjAR6ZrfcAUhTtOHSvSakaf9aSwzZnjEaYiO6dSHx0UhwH+gIxF\\n4jWwvQIDAQABoxYwFDASBgNVHRMBAf8ECDAGAQH/AgEAMA0GCSqGSIb3DQEBCwUA\\nA4IBAQByQWdUNDPEyiGz5HCpboeKwSmCrYxuM5JzORuaj6yJY3G7TPKnYb4ltM+6\\nU5d99PArl09QgNt6RNWyVdqp+/QycPwadArcS5G+1z/Md6yfp6/KlMD1cuxtxxsS\\nl4vGbaTAwUT5ldwQ3HfyZFnQfWyMTr7f4fwsrJkyjTPetPIqYHQgg63qbyucFiut\\nT+XkL+oQU41uSH9SHZ9iqJwTK+UFYi3d0cu2ex9ouf8YDihNn60sLKZiwHkYM5aj\\n0919+nHI2xXAwZi9+Qjo8VL5qxJK4U2tUUAu06NRnWTc+khUOI9qy3D5YjglqmV1\\nMb5SVhJd+pwMe1BK/zMlEOpk2Z3D\\n-----END CERTIFICATE-----,certSerialNumber:0,commonName:C=US,O=Google\\\\, Inc,CN=Cloud SQL Server CA,dnQualifier=a8861e76-1b4f-4efb-bc83-396bbfeda05b,createTime:2026-11-25T09:55:26.377Z,expirationTime:2036-11-22T09:56:26.377Z,instance:tbc0aj584ko18830vb1j,kind:sql#sslCert,sha1Fingerprint:6a34e29353b85234b2b6769082c46082deded54a}"
    },
    {
      "key": "ServiceAccountEmailAddress",
      "value": "MASKED_EMAIL"
    },
    {
      "key": "Settings",
      "value": "{activationPolicy:ALWAYS,availabilityType:ZONAL,backupConfiguration:{backupRetentionSettings:{retainedBackups:7,retentionUnit:COUNT},backupTier:STANDARD,enabled:true,kind:sql#backupConfiguration,startTime:16:00,transactionLogRetentionDays:7,transactionalLogStorageState:TRANSACTIONAL_LOG_STORAGE_STATE_UNSPECIFIED},connectorEnforcement:NOT_REQUIRED,dataDiskSizeGb:100,dataDiskType:PD_SSD,edition:ENTERPRISE,enableDataplexIntegration:true,ipConfiguration:{authorizedNetworks:[{kind:sql#aclEntry,name:AllowAll,value:0.0.0.0/0}],ipv4Enabled:true,serverCaMode:GOOGLE_MANAGED_INTERNAL_CA,serverCertificateRotationMode:SERVER_CERTIFICATE_ROTATION_MODE_UNSPECIFIED,sslMode:ALLOW_UNENCRYPTED_AND_ENCRYPTED},kind:sql#settings,locationPreference:{kind:sql#locationPreference,zone:us-central1-c},pricingPlan:PER_USE,replicationLagMaxSeconds:31536000,replicationType:SYNCHRONOUS,settingsVersion:1,storageAutoResize:true,tier:db-n1-standard-2}"
    },
    {
      "key": "SqlNetworkArchitecture",
      "value": "NEW_NETWORK_ARCHITECTURE"
    },
    {
      "key": "State",
      "value": "RUNNABLE"
    },
    {
      "key": "SwitchTransactionLogsToCloudStorageEnabled",
      "value": "false"
    },
    {
      "key": "UpgradableDatabaseVersions",
      "value": "{displayName:MySQL 8.0.18,majorVersion:MYSQL_8_0,name:MYSQL_8_0_18}; {displayName:MySQL 8.0.26,majorVersion:MYSQL_8_0,name:MYSQL_8_0_26}; {displayName:MySQL 8.0.27,majorVersion:MYSQL_8_0,name:MYSQL_8_0_27}; {displayName:MySQL 8.0.28,majorVersion:MYSQL_8_0,name:MYSQL_8_0_28}; {displayName:MySQL 8.0.29,majorVersion:MYSQL_8_0,name:MYSQL_8_0_29}; {displayName:MySQL 8.0.30,majorVersion:MYSQL_8_0,name:MYSQL_8_0_30}; {displayName:MySQL 8.0.31,majorVersion:MYSQL_8_0,name:MYSQL_8_0_31}; {displayName:MySQL 8.0.32,majorVersion:MYSQL_8_0,name:MYSQL_8_0_32}; {displayName:MySQL 8.0.33,majorVersion:MYSQL_8_0,name:MYSQL_8_0_33}; {displayName:MySQL 8.0.34,majorVersion:MYSQL_8_0,name:MYSQL_8_0_34}; {displayName:MySQL 8.0.35,majorVersion:MYSQL_8_0,name:MYSQL_8_0_35}; {displayName:MySQL 8.0.36,majorVersion:MYSQL_8_0,name:MYSQL_8_0_36}; {displayName:MySQL 8.0.37,majorVersion:MYSQL_8_0,name:MYSQL_8_0_37}; {displayName:MySQL 8.0.39,majorVersion:MYSQL_8_0,name:MYSQL_8_0_39}; {displayName:MySQL 8.0.40,majorVersion:MYSQL_8_0,name:MYSQL_8_0_40}; {displayName:MySQL 8.0.41,majorVersion:MYSQL_8_0,name:MYSQL_8_0_41}; {displayName:MySQL 8.0.42,majorVersion:MYSQL_8_0,name:MYSQL_8_0_42}; {displayName:MySQL 8.0.43,majorVersion:MYSQL_8_0,name:MYSQL_8_0_43}; {displayName:MySQL 8.0.44,majorVersion:MYSQL_8_0,name:MYSQL_8_0_44}; {displayName:MySQL 8.0.45,majorVersion:MYSQL_8_0,name:MYSQL_8_0_45}; {displayName:MySQL 8.4,majorVersion:MYSQL_8_4,name:MYSQL_8_4}"
    }
  ],
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
  "uid": "tbc0aj584ko18830vb1j",
  "vNetId": "test-rdbms-vnet-gcp"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 27ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "backupTime": "16:23-16:53",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:55:24Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:55:24Z",
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
      "cspResourceId": "tbvtoteol9rtlis035nc",
      "cspResourceName": "tbvtoteol9rtlis035nc",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.46-rds.20260908",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
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
          "value": "creating"
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
          "key": "Iops",
          "value": "3000"
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
          "value": "{AllocatedStorage:null,BackupRetentionPeriod:null,CACertificateIdentifier:null,DBInstanceClass:null,DBInstanceIdentifier:null,DBSubnetGroupName:null,EngineVersion:null,IAMDatabaseAuthenticationEnabled:null,Iops:null,LicenseModel:null,MasterUserPassword:****,MultiAZ:null,PendingCloudwatchLogsExports:null,Port:null,ProcessorFeatures:null,StorageType:null}"
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
          "value": "tbvtoteol9rtlis035nc"
        }
      ],
      "uid": "tbvtoteol9rtlis035nc",
      "vNetId": "test-rdbms-vnet-aws"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:54:56Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:54:56Z",
          "reason": "Creating",
          "status": "False",
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-azure",
      "name": "test-rdbms-azure",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "status": "Creating",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbqt11ul0jr723gfs0im",
      "vNetId": "test-rdbms-vnet-azure"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "16:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:58:37Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:58:37Z",
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
      "cspResourceId": "tbc0aj584ko18830vb1j",
      "cspResourceName": "tbc0aj584ko18830vb1j",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "dbInstanceType": "ZONAL",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "encryption": true,
      "endpoint": "136.80.5.81:3306",
      "highAvailability": false,
      "id": "test-rdbms-gcp",
      "keyValueList": [
        {
          "key": "BackendType",
          "value": "SECOND_GEN"
        },
        {
          "key": "ConnectionName",
          "value": "ykkim-etri:us-central1:tbc0aj584ko18830vb1j"
        },
        {
          "key": "CreateTime",
          "value": "2026-11-25T09:55:37.790Z"
        },
        {
          "key": "CurrentDiskSize",
          "value": "0"
        },
        {
          "key": "DatabaseInstalledVersion",
          "value": "MYSQL_8_0_46"
        },
        {
          "key": "DatabaseVersion",
          "value": "MYSQL_8_0"
        },
        {
          "key": "Etag",
          "value": "74c240921252e7aa3276e6b5dc8adb3ebefa361edcf0d4a331a9eee4d0012ea1"
        },
        {
          "key": "GceZone",
          "value": "us-central1-c"
        },
        {
          "key": "GeminiConfig",
          "value": "{flagRecommenderEnabled:true}"
        },
        {
          "key": "IncludeReplicasForMajorVersionUpgrade",
          "value": "false"
        },
        {
          "key": "InstanceType",
          "value": "CLOUD_SQL_INSTANCE"
        },
        {
          "key": "IpAddresses",
          "value": "{ipAddress:136.80.5.81,type:PRIMARY}"
        },
        {
          "key": "Kind",
          "value": "sql#instance"
        },
        {
          "key": "MaintenanceVersion",
          "value": "MYSQL_8_0_46.R20261013.02_15"
        },
        {
          "key": "MaxDiskSize",
          "value": "0"
        },
        {
          "key": "Name",
          "value": "tbc0aj584ko18830vb1j"
        },
        {
          "key": "NodeCount",
          "value": "0"
        },
        {
          "key": "Project",
          "value": "ykkim-etri"
        },
        {
          "key": "Region",
          "value": "us-central1"
        },
        {
          "key": "SatisfiesPzi",
          "value": "true"
        },
        {
          "key": "SatisfiesPzs",
          "value": "false"
        },
        {
          "key": "SelfLink",
          "value": "https://sqladmin.googleapis.com/sql/v1beta4/projects/GCP_PROJECT_ID/instances/tbc0aj584ko18830vb1j"
        },
        {
          "key": "ServerCaCert",
          "value": "{cert:-----BEGIN CERTIFICATE-----\\nMIIDcTCCAlmgAwIBAgIBADANBgkqhkiG9w0BAQsFADBwMS0wKwYDVQQuEyRhODg2\\nMWU3Ni0xYjRmLTRlZmItYmM4My0zOTZiYmZlZGEwNWIxHDAaBgNVBAMTE0Nsb3Vk\\nIFNRTCBTZXJ2ZXIgQ0ExFDASBgNVBAoTC0dvb2dsZSwgSW5jMQswCQYDVQQGEwJV\\nUzAeFw0yNjExMjUwOTU1MjZaFw0zNjExMjIwOTU2MjZaMHAxLTArBgNVBC4TJGE4\\nODYxZTc2LTFiNGYtNGVmYi1iYzgzLTM5NmJiZmVkYTA1YjEcMBoGA1UEAxMTQ2xv\\ndWQgU1FMIFNlcnZlciBDQTEUMBIGA1UEChMLR29vZ2xlLCBJbmMxCzAJBgNVBAYT\\nAlVTMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAilpdp2ZUIpAlxNob\\nmpLtJjEGEMvs0V77gI+27HmvrWkNOT3ppFhxLEnvsqdzKbVzSvD3ptPYm7zqiAoX\\nthlJ0LFmSUVw5ViP6EiTtywMiG0QNbt7UUSvU849j96MYN3h2wTJZgKm0co26Vxw\\ndZ2W4TK8WKOD5BF9vCGSkJVkpF1xPzOFl231DYm0KLrBRM4rMThSwB3ICifQiUso\\nlnIgiLZtdcEd1dXm1E8SMaAva6S7J3nd1SKpypnaCr4B3STY5wUEu8OXhDvYfzId\\nlRfEH5r0xptyIZjAR6ZrfcAUhTtOHSvSakaf9aSwzZnjEaYiO6dSHx0UhwH+gIxF\\n4jWwvQIDAQABoxYwFDASBgNVHRMBAf8ECDAGAQH/AgEAMA0GCSqGSIb3DQEBCwUA\\nA4IBAQByQWdUNDPEyiGz5HCpboeKwSmCrYxuM5JzORuaj6yJY3G7TPKnYb4ltM+6\\nU5d99PArl09QgNt6RNWyVdqp+/QycPwadArcS5G+1z/Md6yfp6/KlMD1cuxtxxsS\\nl4vGbaTAwUT5ldwQ3HfyZFnQfWyMTr7f4fwsrJkyjTPetPIqYHQgg63qbyucFiut\\nT+XkL+oQU41uSH9SHZ9iqJwTK+UFYi3d0cu2ex9ouf8YDihNn60sLKZiwHkYM5aj\\n0919+nHI2xXAwZi9+Qjo8VL5qxJK4U2tUUAu06NRnWTc+khUOI9qy3D5YjglqmV1\\nMb5SVhJd+pwMe1BK/zMlEOpk2Z3D\\n-----END CERTIFICATE-----,certSerialNumber:0,commonName:C=US,O=Google\\\\, Inc,CN=Cloud SQL Server CA,dnQualifier=a8861e76-1b4f-4efb-bc83-396bbfeda05b,createTime:2026-11-25T09:55:26.377Z,expirationTime:2036-11-22T09:56:26.377Z,instance:tbc0aj584ko18830vb1j,kind:sql#sslCert,sha1Fingerprint:6a34e29353b85234b2b6769082c46082deded54a}"
        },
        {
          "key": "ServiceAccountEmailAddress",
          "value": "MASKED_EMAIL"
        },
        {
          "key": "Settings",
          "value": "{activationPolicy:ALWAYS,availabilityType:ZONAL,backupConfiguration:{backupRetentionSettings:{retainedBackups:7,retentionUnit:COUNT},backupTier:STANDARD,enabled:true,kind:sql#backupConfiguration,startTime:16:00,transactionLogRetentionDays:7,transactionalLogStorageState:TRANSACTIONAL_LOG_STORAGE_STATE_UNSPECIFIED},connectorEnforcement:NOT_REQUIRED,dataDiskSizeGb:100,dataDiskType:PD_SSD,edition:ENTERPRISE,enableDataplexIntegration:true,ipConfiguration:{authorizedNetworks:[{kind:sql#aclEntry,name:AllowAll,value:0.0.0.0/0}],ipv4Enabled:true,serverCaMode:GOOGLE_MANAGED_INTERNAL_CA,serverCertificateRotationMode:SERVER_CERTIFICATE_ROTATION_MODE_UNSPECIFIED,sslMode:ALLOW_UNENCRYPTED_AND_ENCRYPTED},kind:sql#settings,locationPreference:{kind:sql#locationPreference,zone:us-central1-c},pricingPlan:PER_USE,replicationLagMaxSeconds:31536000,replicationType:SYNCHRONOUS,settingsVersion:1,storageAutoResize:true,tier:db-n1-standard-2}"
        },
        {
          "key": "SqlNetworkArchitecture",
          "value": "NEW_NETWORK_ARCHITECTURE"
        },
        {
          "key": "State",
          "value": "RUNNABLE"
        },
        {
          "key": "SwitchTransactionLogsToCloudStorageEnabled",
          "value": "false"
        },
        {
          "key": "UpgradableDatabaseVersions",
          "value": "{displayName:MySQL 8.0.18,majorVersion:MYSQL_8_0,name:MYSQL_8_0_18}; {displayName:MySQL 8.0.26,majorVersion:MYSQL_8_0,name:MYSQL_8_0_26}; {displayName:MySQL 8.0.27,majorVersion:MYSQL_8_0,name:MYSQL_8_0_27}; {displayName:MySQL 8.0.28,majorVersion:MYSQL_8_0,name:MYSQL_8_0_28}; {displayName:MySQL 8.0.29,majorVersion:MYSQL_8_0,name:MYSQL_8_0_29}; {displayName:MySQL 8.0.30,majorVersion:MYSQL_8_0,name:MYSQL_8_0_30}; {displayName:MySQL 8.0.31,majorVersion:MYSQL_8_0,name:MYSQL_8_0_31}; {displayName:MySQL 8.0.32,majorVersion:MYSQL_8_0,name:MYSQL_8_0_32}; {displayName:MySQL 8.0.33,majorVersion:MYSQL_8_0,name:MYSQL_8_0_33}; {displayName:MySQL 8.0.34,majorVersion:MYSQL_8_0,name:MYSQL_8_0_34}; {displayName:MySQL 8.0.35,majorVersion:MYSQL_8_0,name:MYSQL_8_0_35}; {displayName:MySQL 8.0.36,majorVersion:MYSQL_8_0,name:MYSQL_8_0_36}; {displayName:MySQL 8.0.37,majorVersion:MYSQL_8_0,name:MYSQL_8_0_37}; {displayName:MySQL 8.0.39,majorVersion:MYSQL_8_0,name:MYSQL_8_0_39}; {displayName:MySQL 8.0.40,majorVersion:MYSQL_8_0,name:MYSQL_8_0_40}; {displayName:MySQL 8.0.41,majorVersion:MYSQL_8_0,name:MYSQL_8_0_41}; {displayName:MySQL 8.0.42,majorVersion:MYSQL_8_0,name:MYSQL_8_0_42}; {displayName:MySQL 8.0.43,majorVersion:MYSQL_8_0,name:MYSQL_8_0_43}; {displayName:MySQL 8.0.44,majorVersion:MYSQL_8_0,name:MYSQL_8_0_44}; {displayName:MySQL 8.0.45,majorVersion:MYSQL_8_0,name:MYSQL_8_0_45}; {displayName:MySQL 8.4,majorVersion:MYSQL_8_4,name:MYSQL_8_4}"
        }
      ],
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
      "uid": "tbc0aj584ko18830vb1j",
      "vNetId": "test-rdbms-vnet-gcp"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:55:07Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:55:07Z",
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
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:54:55Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:54:55Z",
          "reason": "Creating",
          "status": "False",
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.4",
      "dbInstanceSpec": "m1.medium",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-openstack",
      "name": "test-rdbms-openstack",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-openstack"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "RBD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbvg10m2pcf9d7qe09hr",
      "vNetId": "test-rdbms-vnet-openstack"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:55:01Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:55:01Z",
          "reason": "Creating",
          "status": "False",
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-tencent",
      "name": "test-rdbms-tencent",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "CLOUD_HSSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbjgc7odeet6a3manird",
      "vNetId": "test-rdbms-vnet-tencent"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 1.762s
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
- **Duration:** 1.709s
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
- **Duration:** 2.863s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": false,
    "issuer": "CN=Cloud SQL Server CA,O=Google\\, Inc,C=US,2.5.4.46=#132461383836316537362d316234662d346566622d626338332d333936626266656461303562",
    "notAfter": "2036-11-22T09:56:26Z",
    "pem": "-----BEGIN CERTIFICATE-----\nMIIDkDCCAnigAwIBAgIEet79ajANBgkqhkiG9w0BAQsFADBwMS0wKwYDVQQuEyRh\nODg2MWU3Ni0xYjRmLTRlZmItYmM4My0zOTZiYmZlZGEwNWIxHDAaBgNVBAMTE0Ns\nb3VkIFNRTCBTZXJ2ZXIgQ0ExFDASBgNVBAoTC0dvb2dsZSwgSW5jMQswCQYDVQQG\nEwJVUzAeFw0yNjExMjUwOTU1MjZaFw0zNjExMjIwOTU2MjZaME0xKDAmBgNVBAMT\nH3lra2ltLWV0cmk6dGJjMGFqNTg0a28xODgzMHZiMWoxFDASBgNVBAoTC0dvb2ds\nZSwgSW5jMQswCQYDVQQGEwJVUzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBALSS77ZmVL+cLBFsRK/5bYgzxxzJdk+F5dEezTXlvosri2NTtl7N56V7Ppkf\npkrM7YqHWomqHXtHnILsH7MJWpq79GA/1D//Eo3b+qUSVZhhfQzy3rXHppCXvyOZ\nuGUNqInOFqCn0kwnqhfcPmCf4zshTi5a4vY59VjJHMffqsDnijIvsa4N/AsQylhV\npVrOhBT8VAXDC/rtD8wZv3LBx7S0q4JVBAAHK0njOdhpwgZ0vAc7pKwjAOcTEu5W\ngJPXI4e7RwKpnKJNdo2aIo/Svsxs2FPaA4f3Zjt2m021A9GV3qBvqv6YY+Lqddri\niD/k/Uc8ovQI2nrlD5PLfDWFOpkCAwEAAaNVMFMwCQYDVR0TBAIwADBGBgNVHREE\nPzA9gjsxLTFkYWQxYTdlLWJkY2MtNGNkYy1iM2EyLTU3MWRhZDdjYTMzNi51cy1j\nZW50cmFsMS5zcWwuZ29vZzANBgkqhkiG9w0BAQsFAAOCAQEAetj+C0VkYHs9nEhQ\nZTMm89oeM25N7dpFaDS7DtqcoZv+GSdCRvMyqyCIm1GANUNfqKduQUE0fNn1UGbu\nsiFS/HaTS4Af2tKZOFzaBeoowOP5RnL7gQ5NDjaE4VHbcTpc4QBfTbdOvLPmW+mw\nXHYL+oO+Q5vpOarQ1HYSts4KJ+/Kk3u2KqD2fXuplWp9/LDHCIY8pJJUSVNy+IvG\nNfmcUtAyIuYvWlmmx/AsY6TT3rByJ2EaALHh+eVRQA+RABR86XvJkMkTl9JHUpD5\nBl7kt6QNiP0FnHvl6pmtOdtYqRkHZXQU6RAL/8k/0VHb/0rV3MqFPTkrQu0SHgIt\nY9tG+A==\n-----END CERTIFICATE-----\n",
    "subject": "CN=ykkim-etri:tbc0aj584ko18830vb1j,O=Google\\, Inc,C=US"
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
- **Duration:** 1.423s
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m18.835s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 8.13s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m11.681s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 24.294s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 2m4.999s

