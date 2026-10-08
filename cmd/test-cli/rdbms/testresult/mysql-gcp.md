# Managed RDBMS (MySQL) Test Report: GCP (us-central1)

- **Test Case:** GCP US-Central1 MySQL Test
- **Date & Time:** 2026-10-08 13:20:42
- **Namespace:** `default`
- **Total Duration:** 13m47.48s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** GCP
- **Target Region:** `us-central1`
- **Namespace:** `default`
- **Test Date:** 2026-10-08 13:20:42

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
- **Duration:** 7.527s
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
    "queriedAt": "2026-10-08T04:20:49.826943068Z",
    "region": "us-central1",
    "source": "gcp:machineTypes.aggregatedList",
    "zones": [
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-c"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-b"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-f"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "us-central1-a"
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
  "suggestedZone": "us-central1-c"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 33.939s
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
      "lastTransitionTime": "2026-10-08T04:21:23Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-08T04:21:23Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-10-08T04:21:23Z",
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
  "cspResourceId": "tb3t1kg96edp27goltjq",
  "cspResourceName": "tb3t1kg96edp27goltjq",
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
      "value": "2026-10-07T21:20:50.962-07:00"
    },
    {
      "key": "EnableUlaInternalIpv6",
      "value": "false"
    },
    {
      "key": "Id",
      "value": "3291048371810903245"
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
      "value": "tb3t1kg96edp27goltjq"
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
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb3t1kg96edp27goltjq"
    },
    {
      "key": "SelfLinkWithId",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/3291048371810903245"
    },
    {
      "key": "Subnetworks",
      "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1/subnetworks/tbmi96o0b323241lr7ek"
    }
  ],
  "name": "test-rdbms-vnet-gcp",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:23Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:23Z",
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
      "cspResourceId": "tbmi96o0b323241lr7ek",
      "cspResourceName": "tbmi96o0b323241lr7ek",
      "cspVNetId": "tb3t1kg96edp27goltjq",
      "cspVNetName": "tb3t1kg96edp27goltjq",
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
          "value": "2026-10-07T21:21:12.101-07:00"
        },
        {
          "key": "EnableFlowLogs",
          "value": "false"
        },
        {
          "key": "Fingerprint",
          "value": "Frh6S_fmQrs="
        },
        {
          "key": "GatewayAddress",
          "value": "10.2.1.1"
        },
        {
          "key": "Id",
          "value": "5029531793270416599"
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
          "value": "tbmi96o0b323241lr7ek"
        },
        {
          "key": "Network",
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb3t1kg96edp27goltjq"
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
          "value": "https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/regions/us-central1/subnetworks/tbmi96o0b323241lr7ek"
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
          "value": "tbmi96o0b323241lr7ek"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbmi96o0b323241lr7ek",
      "zone": "us-central1-a"
    }
  ],
  "systemLabel": "",
  "uid": "tb3t1kg96edp27goltjq"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 33.224s
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
  "cspResourceId": "tbbftp0p5j3hcoulde7r",
  "cspResourceName": "tbbftp0p5j3hcoulde7r",
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
      "value": "{allowed:[{IPProtocol:tcp,ports:[3306]}],creationTimestamp:2026-10-07T21:21:40.924-07:00,direction:INGRESS,id:5969872976733709499,kind:compute#firewall,logConfig:{},name:tbbftp0p5j3hcoulde7r-i-001,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb3t1kg96edp27goltjq,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tbbftp0p5j3hcoulde7r-i-001,sourceRanges:[0.0.0.0/0],targetTags:[tbbftp0p5j3hcoulde7r]}; {allowed:[{IPProtocol:tcp,ports:[22]}],creationTimestamp:2026-10-07T21:21:49.173-07:00,direction:INGRESS,id:7618672129523492019,kind:compute#firewall,logConfig:{},name:tbbftp0p5j3hcoulde7r-i-002,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb3t1kg96edp27goltjq,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tbbftp0p5j3hcoulde7r-i-002,sourceRanges:[0.0.0.0/0],targetTags:[tbbftp0p5j3hcoulde7r]}; {creationTimestamp:2026-10-07T21:21:26.035-07:00,denied:[{IPProtocol:all}],destinationRanges:[0.0.0.0/0],direction:EGRESS,id:937825332087137450,kind:compute#firewall,logConfig:{},name:tbbftp0p5j3hcoulde7r-o-001,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb3t1kg96edp27goltjq,priority:65535,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tbbftp0p5j3hcoulde7r-o-001,targetTags:[tbbftp0p5j3hcoulde7r]}; {allowed:[{IPProtocol:all}],creationTimestamp:2026-10-07T21:21:34.077-07:00,destinationRanges:[0.0.0.0/0],direction:EGRESS,id:8459023115671479458,kind:compute#firewall,logConfig:{},name:tbbftp0p5j3hcoulde7r-o-002,network:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/networks/tb3t1kg96edp27goltjq,priority:1000,selfLink:https://www.googleapis.com/compute/v1/projects/GCP_PROJECT_ID/global/firewalls/tbbftp0p5j3hcoulde7r-o-002,targetTags:[tbbftp0p5j3hcoulde7r]}"
    }
  ],
  "name": "test-rdbms-sg-gcp",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbbftp0p5j3hcoulde7r",
  "vNetId": "test-rdbms-vnet-gcp"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 17ms
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
- **Duration:** 749ms
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
- **Duration:** 38ms
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
- **Duration:** 3m26.03s
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
- **Duration:** 5ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "19:00",
  "conditions": [
    {
      "lastTransitionTime": "2026-10-08T04:25:21Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-08T04:25:21Z",
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
  "cspResourceId": "tb5ivnnbbt2sn5eqjole",
  "cspResourceName": "tb5ivnnbbt2sn5eqjole",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "db-n1-standard-2",
  "dbInstanceType": "ZONAL",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "encryption": true,
  "endpoint": "35.226.5.44:3306",
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
  "uid": "tb5ivnnbbt2sn5eqjole",
  "vNetId": "test-rdbms-vnet-gcp"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 2ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "backupTime": "18:34-19:04",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:30Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:30Z",
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
      "cspResourceId": "tb9i05t58qdva0pqii7q",
      "cspResourceName": "tb9i05t58qdva0pqii7q",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.46",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "tb9i05t58qdva0pqii7q.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com:3306",
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
          "value": "tb9i05t58qdva0pqii7q"
        }
      ],
      "uid": "tb9i05t58qdva0pqii7q",
      "vNetId": "test-rdbms-vnet-aws"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "AUTO",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:25:05Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:25:05Z",
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
      "cspResourceId": "tbncr09n28klnsbdp033",
      "cspResourceName": "tbncr09n28klnsbdp033",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "dbInstanceType": "Burstable",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "tbncr09n28klnsbdp033.mysql.database.azure.com:3306",
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
      "uid": "tbncr09n28klnsbdp033",
      "vNetId": "test-rdbms-vnet-azure"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "19:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:25:21Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:25:21Z",
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
      "cspResourceId": "tb5ivnnbbt2sn5eqjole",
      "cspResourceName": "tb5ivnnbbt2sn5eqjole",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db-n1-standard-2",
      "dbInstanceType": "ZONAL",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "encryption": true,
      "endpoint": "35.226.5.44:3306",
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
      "uid": "tb5ivnnbbt2sn5eqjole",
      "vNetId": "test-rdbms-vnet-gcp"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:12Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:12Z",
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
      "dbInstanceSpec": "b3c.4x16.encrypted",
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
      "uid": "tbjg8pdr8lgovsqj2clo",
      "vNetId": "test-rdbms-vnet-ibm"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "16:00Z-17:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:24:46Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:24:46Z",
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
      "cspResourceId": "rm-mj75n3v6blq0q6gef",
      "cspResourceName": "tb4himg14kr1h819depl",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "mariadb.n2.medium.2c",
      "dbInstanceType": "HighAvailability",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "43.108.66.234:3306",
      "highAvailability": true,
      "id": "test-rdbms-mariadb-alibaba",
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
      "uid": "tb4himg14kr1h819depl",
      "vNetId": "test-rdbms-vnet-mariadb-alibaba"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "16:33-17:03",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:22:18Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:22:18Z",
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
      "cspResourceId": "tbnon77lqh7fau59hm9p",
      "cspResourceName": "tbnon77lqh7fau59hm9p",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6.27",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-aws",
      "iops": "3000",
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
          "value": "tbnon77lqh7fau59hm9p"
        }
      ],
      "uid": "tbnon77lqh7fau59hm9p",
      "vNetId": "test-rdbms-vnet-mariadb-aws"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:25Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:25Z",
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
      "uid": "tbfcn32o0g0dc0nn4ge4",
      "vNetId": "test-rdbms-vnet-mariadb-nhn"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:21:23Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:21:23Z",
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
      "uid": "tbgv5t06psvjaa5r2i0d",
      "vNetId": "test-rdbms-vnet-nhn"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "00:00-12:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T04:24:52Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:24:52Z",
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
      "cspResourceId": "cdb-5d1t4gzi",
      "cspResourceName": "tbgg0091uj25fuq62c78",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "kr-cdb-5d1t4gzi.sql.tencentcdb.com:29206",
      "highAvailability": false,
      "id": "test-rdbms-tencent",
      "name": "test-rdbms-tencent",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-tencent"
      ],
      "status": "Available",
      "storageSize": 100,
      "storageType": "CLOUD_HSSD",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbgg0091uj25fuq62c78",
      "vNetId": "test-rdbms-vnet-tencent"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 1.512s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/database`
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
- **Duration:** 1.549s
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
- **Duration:** 2.689s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": false,
    "issuer": "CN=Cloud SQL Server CA,O=Google\\, Inc,C=US,2.5.4.46=#132431313536326439332d653762362d343331392d623736332d336533306135363034616163",
    "notAfter": "2036-10-05T04:22:41Z",
    "pem": "-----BEGIN CERTIFICATE-----\nMIIDkDCCAnigAwIBAgIEO1IXSTANBgkqhkiG9w0BAQsFADBwMS0wKwYDVQQuEyQx\nMTU2MmQ5My1lN2I2LTQzMTktYjc2My0zZTMwYTU2MDRhYWMxHDAaBgNVBAMTE0Ns\nb3VkIFNRTCBTZXJ2ZXIgQ0ExFDASBgNVBAoTC0dvb2dsZSwgSW5jMQswCQYDVQQG\nEwJVUzAeFw0yNjEwMDgwNDIxNDFaFw0zNjEwMDUwNDIyNDFaME0xKDAmBgNVBAMT\nH3lra2ltLWV0cmk6dGI1aXZubmJidDJzbjVlcWpvbGUxFDASBgNVBAoTC0dvb2ds\nZSwgSW5jMQswCQYDVQQGEwJVUzCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\nggEBALR3Y905zNsXon4MmyrvUrNLrSQhXc8570wZvWQKo9EkB2jZ5VNIzMTOnolJ\ne+2Ly50jsQVO+UxenJQ7c3fiuMCo9/XhVpLVjDGofGVMGV+un/9638uo6X3y6DQN\neSJQSAqgzTkBkvMq795iU9kY4POIPneQ5syRt4bDFc5/9+HiRUYiLdVpubgrLKBj\nRndcs6+998C3uF4TI00wif0ckWXeVxi9C5qg87hU+KaUZWxo4adkzb20WoHaS/Pw\nsVNle8AA7LZTepc5ljhgUJuO8Bn/HK3lN4+VKcXluFwO6E4SzIo0hY6OzHeQoefQ\nnEnRe62SL+4mRoj4alTawWunvVECAwEAAaNVMFMwCQYDVR0TBAIwADBGBgNVHREE\nPzA9gjsxLTM4NTYzZjE0LWVhMDUtNDNlOC1iNzA3LWYwZjY1ODMwNjNlYi51cy1j\nZW50cmFsMS5zcWwuZ29vZzANBgkqhkiG9w0BAQsFAAOCAQEAjdRw8olHgmoMFIzo\nXgV6lYITUY2AHLn/STS87ZiS/bKwg7iWZ0aZxVUWSLAx4ZrRt34x6dI/gjCqcJlx\nAvP2c94onHNyGqQXKQNabZ9zwXRneQLUYn4SQnCrEkm5RvJi9IMOOY6KjpesXe/g\nxwLLetKgNFQmEyMkf+uiz36pknx34m8TcqUea1g/1gy6Imjt81qAFKUDniJCboAk\nB6F20s1lBRDqeve8RIJ5+qaaMJf5tL+BNVgOK5efXOm39a9FLSQ2xTtjEGlDcTI+\nwsHyEOg/RT2tEkc7y6v9zBAkQ5J95S4i1NfrmygDyalqx4bIf78i7ahCudcWR0ys\n42Syxg==\n-----END CERTIFICATE-----\n",
    "subject": "CN=ykkim-etri:tb5ivnnbbt2sn5eqjole,O=Google\\, Inc,C=US"
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
- **Duration:** 178ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m0.254s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 10.636s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m42.272s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-gcp?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 30.152s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 1m36.703s

