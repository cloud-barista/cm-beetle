# Managed RDBMS (MySQL) Test Report: AZURE (koreacentral)

- **Test Case:** Azure KoreaCentral MySQL Test
- **Date & Time:** 2026-10-08 13:20:42
- **Namespace:** `default`
- **Total Duration:** 11m24.763s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** AZURE
- **Target Region:** `koreacentral`
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
- **Duration:** 4.721s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202608070",
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
    "queriedAt": "2026-10-08T04:20:46.40321181Z",
    "region": "koreacentral",
    "source": "azure:CheckSpecAvailability"
  },
  "connectionName": "azure-koreacentral",
  "estimatedCost": "$0.0914/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "azure-australiacentral",
    "creationDate": "",
    "cspImageName": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202609040",
    "description": "",
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
        "value": "24.04.202608070"
      },
      {
        "key": "ID",
        "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/Providers/Microsoft.Compute/Locations/AustraliaCentral/Publishers/Canonical/ArtifactTypes/VMImage/Offers/ubuntu-24_04-lts/Skus/server-arm64/Versions/24.04.202608070"
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
    "fetchedTime": "2026.08.21 14:01:58 Fri",
    "id": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202608070",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202608070",
    "namespace": "system",
    "osArchitecture": "arm64",
    "osDiskSizeGB": -1,
    "osDiskType": "default",
    "osDistribution": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202608070",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "azure",
    "regionList": [
      "common"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tb703qmia9laeua9huea"
  },
  "imageId": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202608070",
  "imageValidation": {
    "cspResourceId": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202609040",
    "isAvailable": true,
    "resourceId": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202608070",
    "resourceName": "Canonical:ubuntu-24_04-lts:server-arm64:24.04.202608070",
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
- **Duration:** 7.862s
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
      "lastTransitionTime": "2026-10-08T04:20:55Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-08T04:20:55Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-10-08T04:20:55Z",
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
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbanjism7o4afs5hffmm",
  "cspResourceName": "tbanjism7o4afs5hffmm",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-azure",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "ID",
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbanjism7o4afs5hffmm"
    },
    {
      "key": "Location",
      "value": "koreacentral"
    },
    {
      "key": "Properties",
      "value": "{addressSpace:{addressPrefixes:[10.1.0.0/16]},enableDdosProtection:false,privateEndpointVNetPolicies:Disabled,provisioningState:Succeeded,resourceGuid:60be6fb8-dda9-4b12-9365-9e81d63aba6a,subnets:[{etag:W/\\53e54505-941d-474d-9005-e2412dcc7a2f\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbanjism7o4afs5hffmm/subnets/tbvjaeqsan37m9aqv8ho,name:tbvjaeqsan37m9aqv8ho,properties:{addressPrefix:10.1.1.0/24,delegations:[],privateEndpointNetworkPolicies:Disabled,privateLinkServiceNetworkPolicies:Enabled,provisioningState:Succeeded,serviceEndpoints:[{locations:[koreacentral,koreasouth],provisioningState:Succeeded,service:Microsoft.Storage}]},type:Microsoft.Network/virtualNetworks/subnets}],virtualNetworkPeerings:[]}"
    },
    {
      "key": "Etag",
      "value": "W/\\53e54505-941d-474d-9005-e2412dcc7a2f\\"
    },
    {
      "key": "Name",
      "value": "tbanjism7o4afs5hffmm"
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
          "lastTransitionTime": "2026-10-08T04:20:55Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T04:20:55Z",
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
      "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbanjism7o4afs5hffmm/subnets/tbvjaeqsan37m9aqv8ho",
      "cspResourceName": "tbvjaeqsan37m9aqv8ho",
      "cspVNetId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbanjism7o4afs5hffmm",
      "cspVNetName": "tbanjism7o4afs5hffmm",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.1.1.0/24",
      "keyValueList": [
        {
          "key": "ID",
          "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/virtualNetworks/tbanjism7o4afs5hffmm/subnets/tbvjaeqsan37m9aqv8ho"
        },
        {
          "key": "Name",
          "value": "tbvjaeqsan37m9aqv8ho"
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
          "value": "W/\\53e54505-941d-474d-9005-e2412dcc7a2f\\"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbvjaeqsan37m9aqv8ho"
    }
  ],
  "systemLabel": "",
  "uid": "tbanjism7o4afs5hffmm"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 1.923s
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
  "cspResourceId": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3",
  "cspResourceName": "tbct69jaq7h87ogtakd3",
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
      "value": "/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3"
    },
    {
      "key": "Location",
      "value": "koreacentral"
    },
    {
      "key": "Properties",
      "value": "{defaultSecurityRules:[{etag:W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3/defaultSecurityRules/AllowVnetInBound,name:AllowVnetInBound,properties:{access:Allow,description:Allow inbound traffic from all VMs in VNET,destinationAddressPrefix:VirtualNetwork,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65000,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:VirtualNetwork,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3/defaultSecurityRules/AllowAzureLoadBalancerInBound,name:AllowAzureLoadBalancerInBound,properties:{access:Allow,description:Allow inbound traffic from azure load balancer,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65001,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:AzureLoadBalancer,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3/defaultSecurityRules/DenyAllInBound,name:DenyAllInBound,properties:{access:Deny,description:Deny all inbound traffic,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Inbound,priority:65500,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3/defaultSecurityRules/AllowVnetOutBound,name:AllowVnetOutBound,properties:{access:Allow,description:Allow outbound traffic from all VMs to all VMs in VNET,destinationAddressPrefix:VirtualNetwork,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65000,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:VirtualNetwork,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3/defaultSecurityRules/AllowInternetOutBound,name:AllowInternetOutBound,properties:{access:Allow,description:Allow outbound traffic from all VMs to Internet,destinationAddressPrefix:Internet,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65001,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules},{etag:W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3/defaultSecurityRules/DenyAllOutBound,name:DenyAllOutBound,properties:{access:Deny,description:Deny all outbound traffic,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:65500,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/defaultSecurityRules}],provisioningState:Succeeded,resourceGuid:704f7c3f-5e17-4d61-9fa2-a7c4238fbefd,securityRules:[{etag:W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3/securityRules/inbound-rules-28225-3306-3306-TCP,name:inbound-rules-28225-3306-3306-TCP,properties:{access:Allow,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:3306,destinationPortRanges:[],direction:Inbound,priority:100,protocol:Tcp,provisioningState:Succeeded,sourceAddressPrefix:0.0.0.0/0,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3/securityRules/inbound-rules-23398-22-22-TCP,name:inbound-rules-23398-22-22-TCP,properties:{access:Allow,destinationAddressPrefix:*,destinationAddressPrefixes:[],destinationPortRange:22,destinationPortRanges:[],direction:Inbound,priority:101,protocol:Tcp,provisioningState:Succeeded,sourceAddressPrefix:0.0.0.0/0,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3/securityRules/deny-outbound,name:deny-outbound,properties:{access:Deny,destinationAddressPrefix:0.0.0.0/0,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:4096,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules},{etag:W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\,id:/subscriptions/AZURE_SUBSCRIPTION_ID/resourceGroups/koreacentral/providers/Microsoft.Network/networkSecurityGroups/tbct69jaq7h87ogtakd3/securityRules/allow-outbound,name:allow-outbound,properties:{access:Allow,destinationAddressPrefix:0.0.0.0/0,destinationAddressPrefixes:[],destinationPortRange:*,destinationPortRanges:[],direction:Outbound,priority:101,protocol:*,provisioningState:Succeeded,sourceAddressPrefix:*,sourceAddressPrefixes:[],sourcePortRange:*,sourcePortRanges:[]},type:Microsoft.Network/networkSecurityGroups/securityRules}]}"
    },
    {
      "key": "Etag",
      "value": "W/\\37bc230d-918b-4820-8f6f-012bd68807cb\\"
    },
    {
      "key": "Name",
      "value": "tbct69jaq7h87ogtakd3"
    },
    {
      "key": "Type",
      "value": "Microsoft.Network/networkSecurityGroups"
    }
  ],
  "name": "test-rdbms-sg-azure",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbct69jaq7h87ogtakd3",
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
- **Duration:** 5.171s
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
    "adminUserName": "azureuser",
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
- **Duration:** 17ms
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
    "adminUserName": "azureuser",
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
- **Duration:** 4m20.826s
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
- **Duration:** 969ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure`
```json
// Response Body
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
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 9ms
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
- **Duration:** 17.241s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure/database`
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
- **Duration:** 1.601s
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
- **Duration:** 776ms
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
- **Duration:** 7ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 4m35.598s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 41.373s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 37.877s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-azure?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 3.704s

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 25.08s

