# Managed RDBMS Test Report: AZURE (koreacentral)

- **Test Case:** Azure KoreaCentral MySQL Test
- **Date & Time:** 2026-08-31 16:38:03
- **Namespace:** `default`
- **Total Duration:** 23m2.192s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** AZURE
- **Target Region:** `koreacentral`
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
- **Duration:** 4.629s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202603120",
  "specId": "azure+koreacentral+standard_d2ps_v6"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "instanceType": "Standard_D2ps_v6",
    "provider": "azure",
    "queriedAt": "2026-08-31T07:38:08.145237344Z",
    "region": "koreacentral",
    "source": "azure:CheckSpecAvailability"
  },
  "connectionName": "azure-koreacentral",
  "estimatedCost": "$0.0914/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "azure-all",
    "creationDate": "",
    "cspImageName": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202608270",
    "description": "Server ARM64",
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
        "value": "server-arm64"
      },
      {
        "key": "Version",
        "value": "24.04.202603120"
      },
      {
        "key": "ID",
        "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/ubuntu-24_04-lts/Skus/server-arm64/Versions/24.04.202603120"
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
    "fetchedTime": "2026.08.21 14:06:39 Fri",
    "id": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202603120",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202603120",
    "namespace": "system",
    "osArchitecture": "arm64",
    "osDiskSizeGB": -1,
    "osDiskType": "default",
    "osDistribution": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202603120",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04 (ARM64)",
    "providerName": "azure",
    "regionList": [
      "common"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "from-assets",
    "uid": "tbojakvjpb90jchpaump"
  },
  "imageId": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202603120",
  "imageValidation": {
    "cspResourceId": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202608270",
    "isAvailable": true,
    "resourceId": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202603120",
    "resourceName": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202603120",
    "status": "Available"
  },
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "azure",
  "regionName": "koreacentral",
  "specDetails": {
    "architecture": "arm64",
    "connectionName": "azure-koreacentral",
    "costPerHour": 0.0914,
    "cspSpecName": "Standard_D2ps_v6",
    "details": [
      {
        "key": "MaxDataDiskCount",
        "value": "8"
      },
      {
        "key": "MemoryInMB",
        "value": "8192"
      },
      {
        "key": "Name",
        "value": "Standard_D2ps_v6"
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
        "value": "V2"
      },
      {
        "key": "DiskControllerTypes",
        "value": "SCSI"
      },
      {
        "key": "SupportedCapacityReservationTypes",
        "value": "Open,Targeted"
      },
      {
        "key": "MemoryGB",
        "value": "8"
      },
      {
        "key": "MaxDataDiskCount",
        "value": "8"
      },
      {
        "key": "CpuArchitectureType",
        "value": "Arm64"
      },
      {
        "key": "LowPriorityCapable",
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
        "value": "1"
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
        "value": "106000000"
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
        "value": "2"
      },
      {
        "key": "LocationInfo_0_Zone_1",
        "value": "3"
      },
      {
        "key": "LocationInfo_0_Zone_2",
        "value": "1"
      },
      {
        "key": "Family",
        "value": "StandardDpsv6Family"
      },
      {
        "key": "Tier",
        "value": "Standard"
      },
      {
        "key": "Size",
        "value": "D2ps_v6"
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
    "id": "azure+koreacentral+standard_d2ps_v6",
    "infraType": "node",
    "memoryGiB": 7.8125,
    "name": "azure+koreacentral+standard_d2ps_v6",
    "namespace": "system",
    "providerName": "azure",
    "regionLatitude": 37.5665,
    "regionLongitude": 126.978,
    "regionName": "koreacentral",
    "rootDiskSize": 0,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbahqjlrbu2ab7hk6kqr",
    "vCPU": 2
  },
  "specId": "azure+koreacentral+standard_d2ps_v6",
  "specValidation": {
    "cspResourceId": "Standard_D2ps_v6",
    "isAvailable": true,
    "resourceId": "azure+koreacentral+standard_d2ps_v6",
    "resourceName": "Standard_D2ps_v6",
    "status": "Available"
  },
  "status": "OK"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 10.805s
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
      "lastTransitionTime": "2026-08-31T07:38:19Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:19Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:19Z",
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
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tblod3lmkg6f8sosleem",
  "cspResourceName": "tblod3lmkg6f8sosleem",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-azure",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "ID",
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tblod3lmkg6f8sosleem"
    },
    {
      "key": "Location",
      "value": "koreacentral"
    },
    {
      "key": "Properties",
      "value": "{addressSpace:{addressPrefixes:[10.1.0.0/16]},enableDdosProtection:false,privateEndpointVNetPolicies:Disabled,provisioningState:Succeeded,resourceGuid:aa2c020d-2c82-4d65-aec2-c73864bc8f0b,subnets:[{etag:W/\\6e819e64-7e45-42e1-a477-8d3258a860b0\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tblod3lmkg6f8sosleem/subnets/tbh7bjrj2l3nla4mk4mk,name:tbh7bjrj2l3nla4mk4mk,properties:{addressPrefix:10.1.1.0/24,delegations:[],privateEndpointNetworkPolicies:Disabled,privateLinkServiceNetworkPolicies:Enabled,provisioningState:Succeeded,serviceEndpoints:[{locations:[koreacentral,koreasouth],provisioningState:Succeeded,service:Microsoft.Storage}]},type:Microsoft.Network/virtualNetworks/subnets}],virtualNetworkPeerings:[]}"
    },
    {
      "key": "Etag",
      "value": "W/\\6e819e64-7e45-42e1-a477-8d3258a860b0\\"
    },
    {
      "key": "Name",
      "value": "tblod3lmkg6f8sosleem"
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
          "lastTransitionTime": "2026-08-31T07:38:19Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-08-31T07:38:19Z",
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
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tblod3lmkg6f8sosleem/subnets/tbh7bjrj2l3nla4mk4mk",
      "cspResourceName": "tbh7bjrj2l3nla4mk4mk",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tblod3lmkg6f8sosleem",
      "cspVNetName": "tblod3lmkg6f8sosleem",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.1.1.0/24",
      "keyValueList": [
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tblod3lmkg6f8sosleem/subnets/tbh7bjrj2l3nla4mk4mk"
        },
        {
          "key": "Name",
          "value": "tbh7bjrj2l3nla4mk4mk"
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
          "value": "W/\\6e819e64-7e45-42e1-a477-8d3258a860b0\\"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbh7bjrj2l3nla4mk4mk"
    }
  ],
  "systemLabel": "",
  "uid": "tblod3lmkg6f8sosleem"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 9.582s
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
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8",
  "cspResourceName": "tbf9qvqjug8c7cv5mdp8",
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
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8"
    },
    {
      "key": "Location",
      "value": "koreacentral"
    },
    {
      "key": "Properties",
      "value": "{defaultSecurityRules:[{etag:W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8/defaultSecurityRules/AllowVnetInBound,name:AllowVnetInBound,properties:{access:Allow,description:Allow inbound traffic from all VMs in VNET,destinationAddressPrefix:VirtualNetwork,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65000,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:VirtualNetwork,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8/defaultSecurityRules/AllowAzureLoadBalancerInBound,name:AllowAzureLoadBalancerInBound,properties:{access:Allow,description:Allow inbound traffic from azure load balancer,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65001,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:AzureLoadBalancer,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8/defaultSecurityRules/DenyAllInBound,name:DenyAllInBound,properties:{access:Deny,description:Deny all inbound traffic,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65500,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8/defaultSecurityRules/AllowVnetOutBound,name:AllowVnetOutBound,properties:{access:Allow,description:Allow outbound traffic from all VMs to all VMs in VNET,destinationAddressPrefix:VirtualNetwork,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65000,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:VirtualNetwork,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8/defaultSecurityRules/AllowInternetOutBound,name:AllowInternetOutBound,properties:{access:Allow,description:Allow outbound traffic from all VMs to Internet,destinationAddressPrefix:Internet,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65001,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8/defaultSecurityRules/DenyAllOutBound,name:DenyAllOutBound,properties:{access:Deny,description:Deny all outbound traffic,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65500,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules}],provisioningState:Succeeded,resourceGuid:6f3d6799-c3ea-4d5a-8628-594c23ca7834,securityRules:[{etag:W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8/securityRules/inbound-rules-68674-3306-3306-TCP,name:inbound-rules-68674-3306-3306-TCP,properties:{access:Allow,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:3306,destinationPortRanges:[],direction:Inbound,priority:100,protocol:Tcp,provisioningState:Succeeded,sourceAddressPrefix:0.0.0.0/0,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8/securityRules/inbound-rules-27237-22-22-TCP,name:inbound-rules-27237-22-22-TCP,properties:{access:Allow,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:22,destinationPortRanges:[],direction:Inbound,priority:101,protocol:Tcp,provisioningState:Succeeded,sourceAddressPrefix:0.0.0.0/0,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8/securityRules/deny-outbound,name:deny-outbound,properties:{access:Deny,destinationAddressPrefix:0.0.0.0/0,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:4096,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbf9qvqjug8c7cv5mdp8/securityRules/allow-outbound,name:allow-outbound,properties:{access:Allow,destinationAddressPrefix:0.0.0.0/0,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:101,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules}]}"
    },
    {
      "key": "Etag",
      "value": "W/\\b929f28c-e913-4d5c-a2b8-bb412c925cb3\\"
    },
    {
      "key": "Name",
      "value": "tbf9qvqjug8c7cv5mdp8"
    },
    {
      "key": "Type",
      "value": "Microsoft.Network/networkSecurityGroups"
    }
  ],
  "name": "test-rdbms-sg-azure",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbf9qvqjug8c7cv5mdp8",
  "vNetId": "test-rdbms-vnet-azure"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 4ms
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
- **Duration:** 6.27s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=azure-koreacentral`
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
      "Standard_D16ads_v6",
      "Standard_D16ds_v4",
      "Standard_D16ds_v6",
      "Standard_D2ads_v6",
      "Standard_D2ds_v4",
      "Standard_D2ds_v6",
      "Standard_D32ads_v6",
      "Standard_D32ds_v4",
      "Standard_D32ds_v6",
      "Standard_D48ads_v6",
      "Standard_D48ds_v4",
      "Standard_D48ds_v6",
      "Standard_D4ads_v6",
      "Standard_D4ds_v4",
      "Standard_D4ds_v6",
      "Standard_D64ads_v6",
      "Standard_D64ds_v4",
      "Standard_D64ds_v6",
      "Standard_D8ads_v6",
      "Standard_D8ds_v4",
      "Standard_D8ds_v6",
      "Standard_D96ads_v6",
      "Standard_D96ds_v6",
      "Standard_E16ads_v6",
      "Standard_E16ds_v4",
      "Standard_E16ds_v5",
      "Standard_E16ds_v6",
      "Standard_E20ads_v6",
      "Standard_E20ds_v4",
      "Standard_E20ds_v5",
      "Standard_E20ds_v6",
      "Standard_E2ads_v6",
      "Standard_E2ds_v4",
      "Standard_E2ds_v5",
      "Standard_E2ds_v6",
      "Standard_E32ads_v6",
      "Standard_E32ds_v4",
      "Standard_E32ds_v5",
      "Standard_E32ds_v6",
      "Standard_E48ads_v6",
      "Standard_E48ds_v4",
      "Standard_E48ds_v5",
      "Standard_E48ds_v6",
      "Standard_E4ads_v6",
      "Standard_E4ds_v4",
      "Standard_E4ds_v5",
      "Standard_E4ds_v6",
      "Standard_E64ads_v6",
      "Standard_E64ds_v4",
      "Standard_E64ds_v5",
      "Standard_E64ds_v6",
      "Standard_E80ids_v4",
      "Standard_E8ads_v6",
      "Standard_E8ds_v4",
      "Standard_E8ds_v5",
      "Standard_E8ds_v6",
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
    "storageSizeRange": {
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
- **Duration:** 3ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "desiredCloud": {
    "csp": "azure",
    "region": "koreacentral"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for azure (koreacentral)",
  "targetCloud": {
    "csp": "azure",
    "region": "koreacentral"
  },
  "targetRDBMSInstances": [
    {
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "rdbmsName": "test-rdbms-azure",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "storageSize": 32,
      "adminUserName": "azureuser",
      "adminUserPassword": "******",
      "vNetId": "test-rdbms-vnet-azure",
      "subnetIds": [
        "subnet-1"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-azure"
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
- **Duration:** 11ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "azureuser",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "azure-koreacentral",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0.21",
  "dbInstanceSpec": "Standard_B2s",
  "name": "test-rdbms-azure",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-azure"
  ],
  "storageSize": 32,
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
    "adminUserName": "azureuser",
    "adminUserPassword": "******",
    "connectionName": "azure-koreacentral",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0.21",
    "dbInstanceSpec": "Standard_B2s",
    "name": "test-rdbms-azure",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-azure"
    ],
    "storageSize": 32,
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
- **Duration:** 14m59.232s
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
      "adminUserName": "azureuser",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.21",
      "dbInstanceSpec": "Standard_B2s",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "test-rdbms-azure",
      "securityGroupIds": [
        "test-rdbms-sg-azure"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 32,
      "subnetIds": [
        "subnet-1"
      ],
      "vNetId": "test-rdbms-vnet-azure"
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
- **Duration:** 1.395s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-azure`
```json
// Response Body
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
  "endpoint": "tbsn9vgep8ehset1mj4o.mysql.database.azure.com:3306",
  "tagList": [
    {
      "key": "sys.cspResourceId",
      "value": "tbsn9vgep8ehset1mj4o"
    },
    {
      "key": "sys.description",
      "value": "Migrated by CM-Beetle from source instance source-mysql-01"
    },
    {
      "key": "sys.manager",
      "value": "cb-tumblebug"
    },
    {
      "key": "sys.uid",
      "value": "tbsn9vgep8ehset1mj4o"
    },
    {
      "key": "sys.connectionName",
      "value": "azure-koreacentral"
    },
    {
      "key": "sys.cspResourceName",
      "value": "tbsn9vgep8ehset1mj4o"
    },
    {
      "key": "sys.id",
      "value": "test-test-rdbms-azure"
    },
    {
      "key": "sys.labelType",
      "value": "rdbms"
    },
    {
      "key": "sys.name",
      "value": "test-test-rdbms-azure"
    },
    {
      "key": "sys.namespace",
      "value": "default"
    }
  ]
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 4ms
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
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:49:43Z"
        },
        {
          "type": "Synced",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:49:43Z"
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
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 18.05s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-azure/database`
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
- **Duration:** 3.227s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-azure/database`
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
- **Duration:** 166ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 5m15.238s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 44.333s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-azure/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 38.651s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-azure?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 4.7s

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 25.891s

