# Managed RDBMS (MySQL) Test Report: AZURE (koreacentral)

- **Test Case:** Azure KoreaCentral MySQL Test
- **Access Mode:** Public (publicAccess=true)
- **Date & Time:** 2026-11-25 18:54:32
- **Namespace:** `default`
- **Total Duration:** 12m47.75s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** AZURE
- **Target Region:** `koreacentral`
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
- **Duration:** 5.225s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "Canonical:ubuntu-24_04-lts:server:24.04.202610210",
  "specId": "azure+koreacentral+standard_b2als_v2"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "cached": true,
    "instanceType": "Standard_B2als_v2",
    "provider": "azure",
    "queriedAt": "2026-11-25T09:54:34.484276376Z",
    "quota": {
      "fitCount": 25,
      "limit": 50,
      "name": "standardBasv2Family",
      "perInstance": 2,
      "scope": "region",
      "unit": "vCPU",
      "used": 0
    },
    "refreshAfter": "2026-11-25T09:59:34.484276376Z",
    "region": "koreacentral",
    "source": "azure:CheckSpecAvailability",
    "status": "Available",
    "ttlSeconds": 300,
    "zones": [
      {
        "available": true,
        "status": "Available",
        "zoneId": "3"
      },
      {
        "available": true,
        "status": "Available",
        "zoneId": "1"
      },
      {
        "available": true,
        "status": "Available",
        "zoneId": "2"
      }
    ]
  },
  "connectionName": "azure-koreacentral",
  "estimatedCost": "$0.0468/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "azure-australiacentral",
    "creationDate": "",
    "cspImageName": "Canonical:ubuntu-24_04-lts:server:24.04.202611170",
    "description": "Server LTS",
    "details": [
      {
        "key": "Location",
        "value": "australiacentral"
      },
      {
        "key": "Publisher",
        "value": "Canonical"
      },
      {
        "key": "Offer",
        "value": "ubuntu-24_04-lts"
      },
      {
        "key": "SKU",
        "value": "server"
      },
      {
        "key": "Version",
        "value": "24.04.202610210"
      },
      {
        "key": "ID",
        "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/ubuntu-24_04-lts/Skus/server/Versions/24.04.202610210"
      },
      {
        "key": "HyperVGeneration",
        "value": "V2"
      },
      {
        "key": "Features",
        "value": "SecurityType=TrustedLaunchSupported, IsAcceleratedNetworkSupported=True, DiskControllerTypes=SCSI, NVMe, IsHibernateSupported=True"
      },
      {
        "key": "FeatureCount",
        "value": "4"
      },
      {
        "key": "ImageDeprecationState",
        "value": "Active"
      }
    ],
    "fetchedTime": "2026.10.29 04:34:25 Thu",
    "id": "Canonical:ubuntu-24_04-lts:server:24.04.202610210",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "Canonical:ubuntu-24_04-lts:server:24.04.202610210",
    "namespace": "system",
    "osArchitecture": "x86_64",
    "osDiskSizeGB": -1,
    "osDiskType": "default",
    "osDistribution": "Canonical:ubuntu-24_04-lts:server:24.04.202610210",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "azure",
    "regionList": [
      "common"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "from-assets",
    "uid": "tbm7el5o8iv7q202l982"
  },
  "imageId": "Canonical:ubuntu-24_04-lts:server:24.04.202610210",
  "imageValidation": {
    "cspResourceId": "Canonical:ubuntu-24_04-lts:server:24.04.202611170",
    "isAvailable": true,
    "resourceId": "Canonical:ubuntu-24_04-lts:server:24.04.202610210",
    "resourceName": "Canonical:ubuntu-24_04-lts:server:24.04.202610210",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "azure",
  "regionName": "koreacentral",
  "specDetails": {
    "architecture": "x86_64",
    "availabilityCheckedAt": "2026-11-25T09:54:34.484276376Z",
    "availabilityQuota": {
      "fitCount": 25,
      "limit": 50,
      "name": "standardBasv2Family",
      "perInstance": 2,
      "scope": "region",
      "unit": "vCPU",
      "used": 0
    },
    "availabilityRefreshAfter": "2026-11-25T09:59:34.484276376Z",
    "availabilityStatus": "Available",
    "availableZones": [
      "1",
      "2",
      "3"
    ],
    "connectionName": "azure-koreacentral",
    "costPerHour": 0.0468,
    "cspSpecName": "Standard_B2als_v2",
    "details": [
      {
        "key": "MaxDataDiskCount",
        "value": "4"
      },
      {
        "key": "MemoryInMB",
        "value": "4096"
      },
      {
        "key": "Name",
        "value": "Standard_B2als_v2"
      },
      {
        "key": "NumberOfCores",
        "value": "2"
      },
      {
        "key": "OSDiskSizeInMB",
        "value": "1047552"
      },
      {
        "key": "ResourceDiskSizeInMB",
        "value": "0"
      },
      {
        "key": "MaxResourceVolumeMB",
        "value": "0"
      },
      {
        "key": "OSVhdSizeMB",
        "value": "1047552"
      },
      {
        "key": "vCPUs",
        "value": "2"
      },
      {
        "key": "MemoryPreservingMaintenanceSupported",
        "value": "True"
      },
      {
        "key": "HyperVGenerations",
        "value": "V1,V2"
      },
      {
        "key": "SupportedCapacityReservationTypes",
        "value": "Open,Targeted"
      },
      {
        "key": "MemoryGB",
        "value": "4"
      },
      {
        "key": "MaxDataDiskCount",
        "value": "4"
      },
      {
        "key": "CpuArchitectureType",
        "value": "x64"
      },
      {
        "key": "LowPriorityCapable",
        "value": "True"
      },
      {
        "key": "HibernationSupported",
        "value": "True"
      },
      {
        "key": "PremiumIO",
        "value": "True"
      },
      {
        "key": "VMDeploymentTypes",
        "value": "IaaS"
      },
      {
        "key": "vCPUsConstraintsAllowed",
        "value": "1, 2"
      },
      {
        "key": "vCPUsAvailable",
        "value": "2"
      },
      {
        "key": "vCPUsPerCore",
        "value": "2"
      },
      {
        "key": "CombinedTempDiskAndCachedIOPS",
        "value": "9000"
      },
      {
        "key": "CombinedTempDiskAndCachedReadBytesPerSecond",
        "value": "125000000"
      },
      {
        "key": "CombinedTempDiskAndCachedWriteBytesPerSecond",
        "value": "125000000"
      },
      {
        "key": "UncachedDiskIOPS",
        "value": "3750"
      },
      {
        "key": "UncachedDiskBytesPerSecond",
        "value": "85000000"
      },
      {
        "key": "EphemeralOSDiskSupported",
        "value": "False"
      },
      {
        "key": "EncryptionAtHostSupported",
        "value": "True"
      },
      {
        "key": "CapacityReservationSupported",
        "value": "True"
      },
      {
        "key": "AcceleratedNetworkingEnabled",
        "value": "True"
      },
      {
        "key": "RdmaEnabled",
        "value": "False"
      },
      {
        "key": "MaxNetworkInterfaces",
        "value": "2"
      },
      {
        "key": "UltraSSDAvailable",
        "value": "False"
      },
      {
        "key": "LocationInfo_0_Location",
        "value": "KoreaCentral"
      },
      {
        "key": "LocationInfo_0_Zone_0",
        "value": "3"
      },
      {
        "key": "LocationInfo_0_Zone_1",
        "value": "1"
      },
      {
        "key": "LocationInfo_0_Zone_2",
        "value": "2"
      },
      {
        "key": "Family",
        "value": "standardBasv2Family"
      },
      {
        "key": "Tier",
        "value": "Standard"
      },
      {
        "key": "Size",
        "value": "B2als_v2"
      },
      {
        "key": "ResourceType",
        "value": "virtualMachines"
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
    "id": "azure+koreacentral+standard_b2als_v2",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "azure+koreacentral+standard_b2als_v2",
    "namespace": "system",
    "providerName": "azure",
    "regionLatitude": 37.5665,
    "regionLongitude": 126.978,
    "regionName": "koreacentral",
    "rootDiskSize": 0,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tb4fm9rbtluo2mp32cgg",
    "vCPU": 2
  },
  "specId": "azure+koreacentral+standard_b2als_v2",
  "specValidation": {
    "cspResourceId": "Standard_B2als_v2",
    "isAvailable": true,
    "resourceId": "azure+koreacentral+standard_b2als_v2",
    "resourceName": "Standard_B2als_v2",
    "status": "Available"
  },
  "status": "OK",
  "suggestedZone": "3"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 11.784s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.1.0.0/16",
  "connectionName": "azure-koreacentral",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-azure",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.1.1.0/24",
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
  "cidrBlock": "10.1.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T09:54:47Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:47Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:47Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbob9pv4a0t0r36rukkk",
  "cspResourceName": "tbob9pv4a0t0r36rukkk",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-azure",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "ID",
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbob9pv4a0t0r36rukkk"
    },
    {
      "key": "Location",
      "value": "koreacentral"
    },
    {
      "key": "Properties",
      "value": "{addressSpace:{addressPrefixes:[10.1.0.0/16]},enableDdosProtection:false,privateEndpointVNetPolicies:Disabled,provisioningState:Succeeded,resourceGuid:07b4f7de-b0b5-401e-8163-3a664149c260,subnets:[{etag:W/\\d72c9434-a01a-4040-a52f-cd357b00e2d3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbob9pv4a0t0r36rukkk/subnets/tba98ditvus0csab9vbk,name:tba98ditvus0csab9vbk,properties:{addressPrefix:10.1.1.0/24,delegations:[],privateEndpointNetworkPolicies:Disabled,privateLinkServiceNetworkPolicies:Enabled,provisioningState:Succeeded,serviceEndpoints:[{locations:[koreacentral,koreasouth],provisioningState:Succeeded,service:Microsoft.Storage}]},type:Microsoft.Network/virtualNetworks/subnets}],virtualNetworkPeerings:[]}"
    },
    {
      "key": "Etag",
      "value": "W/\\d72c9434-a01a-4040-a52f-cd357b00e2d3\\"
    },
    {
      "key": "Name",
      "value": "tbob9pv4a0t0r36rukkk"
    },
    {
      "key": "Type",
      "value": "Microsoft.Network/virtualNetworks"
    }
  ],
  "name": "test-rdbms-vnet-azure",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:54:47Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:54:47Z",
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
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbob9pv4a0t0r36rukkk/subnets/tba98ditvus0csab9vbk",
      "cspResourceName": "tba98ditvus0csab9vbk",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbob9pv4a0t0r36rukkk",
      "cspVNetName": "tbob9pv4a0t0r36rukkk",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.1.1.0/24",
      "keyValueList": [
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbob9pv4a0t0r36rukkk/subnets/tba98ditvus0csab9vbk"
        },
        {
          "key": "Name",
          "value": "tba98ditvus0csab9vbk"
        },
        {
          "key": "Properties",
          "value": "{addressPrefix:10.1.1.0/24,delegations:[],privateEndpointNetworkPolicies:Disabled,privateLinkServiceNetworkPolicies:Enabled,provisioningState:Succeeded,serviceEndpoints:[{locations:[koreacentral,koreasouth],provisioningState:Succeeded,service:Microsoft.Storage}]}"
        },
        {
          "key": "Type",
          "value": "Microsoft.Network/virtualNetworks/subnets"
        },
        {
          "key": "Etag",
          "value": "W/\\d72c9434-a01a-4040-a52f-cd357b00e2d3\\"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tba98ditvus0csab9vbk"
    }
  ],
  "systemLabel": "",
  "uid": "tbob9pv4a0t0r36rukkk"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 3.224s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "azure-koreacentral",
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
  "name": "test-rdbms-sg-azure",
  "vNetId": "test-rdbms-vnet-azure"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig",
  "cspResourceName": "tbbrvi8emks31hq0cjig",
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
  "id": "test-rdbms-sg-azure",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "ID",
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig"
    },
    {
      "key": "Location",
      "value": "koreacentral"
    },
    {
      "key": "Properties",
      "value": "{defaultSecurityRules:[{etag:W/\\61699a49-94eb-4290-a33f-98b3490a0064\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig/defaultSecurityRules/AllowVnetInBound,name:AllowVnetInBound,properties:{access:Allow,description:Allow inbound traffic from all VMs in VNET,destinationAddressPrefix:VirtualNetwork,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65000,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:VirtualNetwork,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\61699a49-94eb-4290-a33f-98b3490a0064\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig/defaultSecurityRules/AllowAzureLoadBalancerInBound,name:AllowAzureLoadBalancerInBound,properties:{access:Allow,description:Allow inbound traffic from azure load balancer,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65001,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:AzureLoadBalancer,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\61699a49-94eb-4290-a33f-98b3490a0064\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig/defaultSecurityRules/DenyAllInBound,name:DenyAllInBound,properties:{access:Deny,description:Deny all inbound traffic,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65500,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\61699a49-94eb-4290-a33f-98b3490a0064\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig/defaultSecurityRules/AllowVnetOutBound,name:AllowVnetOutBound,properties:{access:Allow,description:Allow outbound traffic from all VMs to all VMs in VNET,destinationAddressPrefix:VirtualNetwork,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65000,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:VirtualNetwork,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\61699a49-94eb-4290-a33f-98b3490a0064\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig/defaultSecurityRules/AllowInternetOutBound,name:AllowInternetOutBound,properties:{access:Allow,description:Allow outbound traffic from all VMs to Internet,destinationAddressPrefix:Internet,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65001,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\61699a49-94eb-4290-a33f-98b3490a0064\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig/defaultSecurityRules/DenyAllOutBound,name:DenyAllOutBound,properties:{access:Deny,description:Deny all outbound traffic,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65500,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules}],provisioningState:Succeeded,resourceGuid:43febf8b-6e0d-441d-b4cf-ea166cb885ea,securityRules:[{etag:W/\\61699a49-94eb-4290-a33f-98b3490a0064\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig/securityRules/inbound-rules-3759-3306-3306-TCP,name:inbound-rules-3759-3306-3306-TCP,properties:{access:Allow,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:3306,destinationPortRanges:[],direction:Inbound,priority:100,protocol:Tcp,provisioningState:Succeeded,sourceAddressPrefix:0.0.0.0/0,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\61699a49-94eb-4290-a33f-98b3490a0064\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig/securityRules/inbound-rules-60895-22-22-TCP,name:inbound-rules-60895-22-22-TCP,properties:{access:Allow,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:22,destinationPortRanges:[],direction:Inbound,priority:101,protocol:Tcp,provisioningState:Succeeded,sourceAddressPrefix:0.0.0.0/0,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\61699a49-94eb-4290-a33f-98b3490a0064\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig/securityRules/deny-outbound,name:deny-outbound,properties:{access:Deny,destinationAddressPrefix:0.0.0.0/0,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:4096,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\61699a49-94eb-4290-a33f-98b3490a0064\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbbrvi8emks31hq0cjig/securityRules/allow-outbound,name:allow-outbound,properties:{access:Allow,destinationAddressPrefix:0.0.0.0/0,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:101,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules}]}"
    },
    {
      "key": "Etag",
      "value": "W/\\61699a49-94eb-4290-a33f-98b3490a0064\\"
    },
    {
      "key": "Name",
      "value": "tbbrvi8emks31hq0cjig"
    },
    {
      "key": "Type",
      "value": "Microsoft.Network/networkSecurityGroups"
    }
  ],
  "name": "test-rdbms-sg-azure",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbbrvi8emks31hq0cjig",
  "vNetId": "test-rdbms-vnet-azure"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 9ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=azure`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "azure": {
      "dbOperationMethod": "cspNativeApi",
      "note": "Storage type selection not supported. Storage SKU is read-only and automatically set by Azure based on compute tier (Premium SSD for General Purpose/Memory Optimized tiers, locally redundant storage for Burstable tier). dbOperationMethod is cspNativeApi (armmysqlfs.DatabasesClient). Azure uses SubnetNames only in VPC-private mode (PublicAccess=false); when PublicAccess=true, subnet is not used.",
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
- **Duration:** 5.628s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=azure-koreacentral&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "1-35",
    "connectionName": "azure-koreacentral",
    "dbEngine": "mysql",
    "dbInstanceSpecOptions": [
      "Standard_B12ms",
      "Standard_B16ms",
      "Standard_B1ms",
      "Standard_B20ms",
      "Standard_B2ms",
      "Standard_B2s",
      "Standard_B4ms",
      "Standard_B8ms",
      "Standard_D16ads_v5",
      "Standard_D16ads_v6",
      "Standard_D16ds_v4",
      "Standard_D16ds_v6",
      "Standard_D2ads_v5",
      "Standard_D2ads_v6",
      "Standard_D2ds_v4",
      "Standard_D2ds_v6",
      "Standard_D32ads_v5",
      "Standard_D32ads_v6",
      "Standard_D32ds_v4",
      "Standard_D32ds_v6",
      "Standard_D48ads_v5",
      "Standard_D48ads_v6",
      "Standard_D48ds_v4",
      "Standard_D48ds_v6",
      "Standard_D4ads_v5",
      "Standard_D4ads_v6",
      "Standard_D4ds_v4",
      "Standard_D4ds_v6",
      "Standard_D64ads_v5",
      "Standard_D64ads_v6",
      "Standard_D64ds_v4",
      "Standard_D64ds_v6",
      "Standard_D8ads_v5",
      "Standard_D8ads_v6",
      "Standard_D8ds_v4",
      "Standard_D8ds_v6",
      "Standard_D96ads_v5",
      "Standard_D96ads_v6",
      "Standard_D96ds_v6",
      "Standard_E16ads_v5",
      "Standard_E16ads_v6",
      "Standard_E16ds_v4",
      "Standard_E16ds_v5",
      "Standard_E16ds_v6",
      "Standard_E20ads_v5",
      "Standard_E20ads_v6",
      "Standard_E20ds_v4",
      "Standard_E20ds_v5",
      "Standard_E20ds_v6",
      "Standard_E2ads_v5",
      "Standard_E2ads_v6",
      "Standard_E2ds_v4",
      "Standard_E2ds_v5",
      "Standard_E2ds_v6",
      "Standard_E32ads_v5",
      "Standard_E32ads_v6",
      "Standard_E32ds_v4",
      "Standard_E32ds_v5",
      "Standard_E32ds_v6",
      "Standard_E48ads_v5",
      "Standard_E48ads_v6",
      "Standard_E48ds_v4",
      "Standard_E48ds_v5",
      "Standard_E48ds_v6",
      "Standard_E4ads_v5",
      "Standard_E4ads_v6",
      "Standard_E4ds_v4",
      "Standard_E4ds_v5",
      "Standard_E4ds_v6",
      "Standard_E64ads_v5",
      "Standard_E64ads_v6",
      "Standard_E64ds_v4",
      "Standard_E64ds_v5",
      "Standard_E64ds_v6",
      "Standard_E80ids_v4",
      "Standard_E8ads_v5",
      "Standard_E8ads_v6",
      "Standard_E8ds_v4",
      "Standard_E8ds_v5",
      "Standard_E8ds_v6",
      "Standard_E96ads_v5",
      "Standard_E96ads_v6",
      "Standard_E96ds_v5",
      "Standard_E96ds_v6"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "49152",
        "name": "Standard_B12ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "12"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_B16ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "2048",
        "name": "Standard_B1ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "1"
      },
      {
        "memSizeMiB": "81920",
        "name": "Standard_B20ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "8192",
        "name": "Standard_B2ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "4096",
        "name": "Standard_B2s",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_B4ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_B8ms",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_D16ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_D16ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_D16ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_D16ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "8192",
        "name": "Standard_D2ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "8192",
        "name": "Standard_D2ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "8192",
        "name": "Standard_D2ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "8192",
        "name": "Standard_D2ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_D32ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_D32ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_D32ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_D32ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "196608",
        "name": "Standard_D48ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "196608",
        "name": "Standard_D48ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "196608",
        "name": "Standard_D48ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "196608",
        "name": "Standard_D48ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_D4ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_D4ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_D4ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_D4ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_D64ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_D64ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_D64ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_D64ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_D8ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_D8ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_D8ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_D8ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_D96ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_D96ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_D96ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_E16ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_E16ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_E16ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_E16ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "Standard_E16ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "163840",
        "name": "Standard_E20ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "163840",
        "name": "Standard_E20ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "163840",
        "name": "Standard_E20ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "163840",
        "name": "Standard_E20ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "163840",
        "name": "Standard_E20ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "20"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_E2ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_E2ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_E2ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_E2ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "Standard_E2ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_E32ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_E32ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_E32ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_E32ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "262144",
        "name": "Standard_E32ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_E48ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_E48ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_E48ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_E48ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "393216",
        "name": "Standard_E48ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_E4ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_E4ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_E4ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_E4ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "Standard_E4ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "524288",
        "name": "Standard_E64ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "524288",
        "name": "Standard_E64ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "524288",
        "name": "Standard_E64ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "524288",
        "name": "Standard_E64ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "524288",
        "name": "Standard_E64ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "516080",
        "name": "Standard_E80ids_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "80"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_E8ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_E8ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_E8ds_v4",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_E8ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "Standard_E8ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "688128",
        "name": "Standard_E96ads_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "688128",
        "name": "Standard_E96ads_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "688128",
        "name": "Standard_E96ds_v5",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "786432",
        "name": "Standard_E96ds_v6",
        "storageSizeRangeGB": {
          "max": 0,
          "min": 0
        },
        "vCpuClockGHz": "-1",
        "vCpuCount": "96"
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
          "description": "Storage SKU is automatically determined by Azure based on compute tier. User cannot specify storage type.",
          "displayName": "Automatic (Azure-managed)",
          "recommendationLevel": "standard",
          "storageType": "NA"
        }
      ]
    },
    "providerName": "azure",
    "regionName": "koreacentral",
    "requiresSecurityGroup": false,
    "requiresSubnet": false,
    "storageSizeRangeGB": {
      "max": 35184,
      "min": 21
    },
    "storageTypeOptions": [
      "NA"
    ],
    "supportedVersions": [
      "5.7",
      "8.0.21",
      "8.4",
      "9.5"
    ],
    "supportsBackup": true,
    "supportsDeletionProtection": false,
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
- **Duration:** 5ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "azure",
    "region": "koreacentral"
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
    "adminUserName": "dbadmin",
    "backupRetentionDays": 7,
    "highAvailability": false,
    "publicAccess": true
  }
}
```
```json
// Response Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for azure (koreacentral)",
  "status": "recommended",
  "targetCloud": {
    "csp": "azure",
    "region": "koreacentral"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "dbadmin",
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
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-azure",
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-azure"
    }
  ],
  "warnings": [
    "Storage type selection is not configurable on target cloud (azure); requested storage type 'SSD' for instance 'Source MySQL 01' will be managed automatically by the provider."
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 20ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "dbadmin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "azure-koreacentral",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0.21",
  "dbInstanceSpec": "Standard_B2s",
  "name": "rdbms-azure",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-azure"
  ],
  "storageSize": 100,
  "subnetIds": [
    "subnet-1"
  ],
  "vNetId": "test-rdbms-vnet-azure"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "dbadmin",
    "adminUserPassword": "******",
    "connectionName": "azure-koreacentral",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0.21",
    "dbInstanceSpec": "Standard_B2s",
    "name": "rdbms-azure",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-azure"
    ],
    "storageSize": 100,
    "subnetIds": [
      "subnet-1"
    ],
    "vNetId": "test-rdbms-vnet-azure"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 5m21.59s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for azure (koreacentral)",
  "status": "recommended",
  "targetCloud": {
    "csp": "azure",
    "region": "koreacentral"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "dbadmin",
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
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-azure",
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-azure"
    }
  ],
  "warnings": [
    "Storage type selection is not configurable on target cloud (azure); requested storage type 'SSD' for instance 'Source MySQL 01' will be managed automatically by the provider."
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
- **Duration:** 230ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "AUTO",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T09:59:42Z",
      "reason": "Available",
      "status": "True",
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
  "status": "Available",
  "storageSize": 100,
  "storageType": "Premium_LRS",
  "subnetIds": [
    "subnet-1"
  ],
  "uid": "tbqt11ul0jr723gfs0im",
  "vNetId": "test-rdbms-vnet-azure"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 61ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:59:03Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:59:03Z",
          "reason": "Creating",
          "status": "False",
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
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "highAvailability": false,
      "id": "test-rdbms-alibaba",
      "name": "test-rdbms-alibaba",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "status": "Creating",
      "storageSize": 100,
      "storageType": "cloud_auto",
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
          "value": "configuring-enhanced-monitoring"
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
      "backupTime": "AUTO",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:59:42Z",
          "reason": "Available",
          "status": "True",
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
      "status": "Available",
      "storageSize": 100,
      "storageType": "Premium_LRS",
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
      "backupTime": "00:00-12:00",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:59:46Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:59:46Z",
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
      "cspResourceId": "cdb-q1mk2sbg",
      "cspResourceName": "tbjgc7odeet6a3manird",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "4000",
      "dbInstanceType": "NA",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "kr-cdb-q1mk2sbg.sql.tencentcdb.com:24489",
      "highAvailability": false,
      "id": "test-rdbms-tencent",
      "keyValueList": [
        {
          "key": "WanStatus",
          "value": "1"
        },
        {
          "key": "Zone",
          "value": "ap-seoul-1"
        },
        {
          "key": "InitFlag",
          "value": "1"
        },
        {
          "key": "Memory",
          "value": "4000"
        },
        {
          "key": "Status",
          "value": "1"
        },
        {
          "key": "VpcId",
          "value": "9730283"
        },
        {
          "key": "InstanceId",
          "value": "cdb-q1mk2sbg"
        },
        {
          "key": "Volume",
          "value": "100"
        },
        {
          "key": "AutoRenew",
          "value": "0"
        },
        {
          "key": "ProtectMode",
          "value": "0"
        },
        {
          "key": "SubnetId",
          "value": "820587"
        },
        {
          "key": "InstanceType",
          "value": "1"
        },
        {
          "key": "ProjectId",
          "value": "0"
        },
        {
          "key": "Region",
          "value": "ap-seoul"
        },
        {
          "key": "DeadlineTime",
          "value": "0000-00-00 00:00:00"
        },
        {
          "key": "DeployMode",
          "value": "0"
        },
        {
          "key": "TaskStatus",
          "value": "0"
        },
        {
          "key": "DeviceType",
          "value": "CLOUD_NATIVE_CLUSTER"
        },
        {
          "key": "EngineVersion",
          "value": "8.0"
        },
        {
          "key": "InstanceName",
          "value": "tbjgc7odeet6a3manird"
        },
        {
          "key": "WanDomain",
          "value": "kr-cdb-q1mk2sbg.sql.tencentcdb.com"
        },
        {
          "key": "WanPort",
          "value": "24489"
        },
        {
          "key": "PayType",
          "value": "1"
        },
        {
          "key": "CreateTime",
          "value": "2026-11-25 17:55:12"
        },
        {
          "key": "Vip",
          "value": "10.4.1.13"
        },
        {
          "key": "Vport",
          "value": "3306"
        },
        {
          "key": "CdbError",
          "value": "0"
        },
        {
          "key": "UniqVpcId",
          "value": "vpc-7yqgk5bc"
        },
        {
          "key": "UniqSubnetId",
          "value": "subnet-f4cf8dsh"
        },
        {
          "key": "Cpu",
          "value": "2"
        },
        {
          "key": "Qps",
          "value": "4400"
        },
        {
          "key": "ZoneId",
          "value": "180001"
        },
        {
          "key": "InstanceNodes",
          "value": "2"
        },
        {
          "key": "EngineType",
          "value": "InnoDB"
        },
        {
          "key": "MaxDelayTime",
          "value": "0"
        },
        {
          "key": "DiskType",
          "value": "CLOUD_HSSD"
        },
        {
          "key": "ExpandCpu",
          "value": "0"
        },
        {
          "key": "DeviceBandwidth",
          "value": "0"
        },
        {
          "key": "DestroyProtect",
          "value": "off"
        }
      ],
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
      "uid": "tbjgc7odeet6a3manird",
      "vNetId": "test-rdbms-vnet-tencent"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 19.23s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure/database`
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
- **Duration:** 1.537s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure/database`
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
- **Duration:** 1.315s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": true,
    "issuer": "CN=DigiCert Global Root G2,OU=www.digicert.com,O=DigiCert Inc,C=US",
    "notAfter": "2038-01-15T12:00:00Z",
    "pem": "-----BEGIN CERTIFICATE-----\nMIIDjjCCAnagAwIBAgIQAzrx5qcRqaC7KGSxHQn65TANBgkqhkiG9w0BAQsFADBh\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBH\nMjAeFw0xMzA4MDExMjAwMDBaFw0zODAxMTUxMjAwMDBaMGExCzAJBgNVBAYTAlVT\nMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\nb20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IEcyMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuzfNNNx7a8myaJCtSnX/RrohCgiN9RlUyfuI\n2/Ou8jqJkTx65qsGGmvPrC3oXgkkRLpimn7Wo6h+4FR1IAWsULecYxpsMNzaHxmx\n1x7e/dfgy5SDN67sH0NO3Xss0r0upS/kqbitOtSZpLYl6ZtrAGCSYP9PIUkY92eQ\nq2EGnI/yuum06ZIya7XzV+hdG82MHauVBJVJ8zUtluNJbd134/tJS7SsVQepj5Wz\ntCO7TG1F8PapspUwtP1MVYwnSlcUfIKdzXOS0xZKBgyMUNGPHgm+F6HmIcr9g+UQ\nvIOlCsRnKPZzFBQ9RnbDhxSJITRNrw9FDKZJobq7nMWxM4MphQIDAQABo0IwQDAP\nBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNVHQ4EFgQUTiJUIBiV\n5uNu5g/6+rkS7QYXjzkwDQYJKoZIhvcNAQELBQADggEBAGBnKJRvDkhj6zHd6mcY\n1Yl9PMWLSn/pvtsrF9+wX3N3KjITOYFnQoQj8kVnNeyIv/iPsGEMNKSuIEyExtv4\nNeF22d+mQrvHRAiGfzZ0JFrabA0UWTW98kndth/Jsw1HKj2ZL7tcu7XUIOGZX1NG\nFdtom/DzMNU+MeKNhJ7jitralj41E6Vf8PlwUHBHQRFXGU7Aj64GxJUTFy8bJZ91\n8rGOmaFvE7FBcf6IKshPECBV1/MUReXgRPTqh5Uykw7+U0b6LJ3/iyK5S9kJRaTe\npLiaWN0bfVKfjllDiIGknibVb63dDcY3fe0Dkhvld1927jyNxF1WW6LZZm6zNTfl\nMrY=\n-----END CERTIFICATE-----\n",
    "subject": "CN=DigiCert Global Root G2,OU=www.digicert.com,O=DigiCert Inc,C=US"
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
- **Duration:** 355ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 5m6.032s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 33.246s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 35.884s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 1.549s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 20.827s

