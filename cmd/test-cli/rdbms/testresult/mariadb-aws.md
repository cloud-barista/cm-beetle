# Managed RDBMS (MariaDB) Test Report: AWS (ap-northeast-2)

- **Test Case:** AWS AP-Northeast-2 (Seoul) MariaDB Test
- **Access Mode:** Public (publicAccess=true)
- **Date & Time:** 2026-11-25 19:27:54
- **Namespace:** `default`
- **Total Duration:** 17m54.391s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** AWS
- **Target Region:** `ap-northeast-2`
- **Namespace:** `default`
- **Test Date:** 2026-11-25 19:27:54

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
- **Duration:** 1.475s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "ami-0e0a5fdc771990a70",
  "specId": "aws+ap-northeast-2+t4g.medium"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "cached": true,
    "instanceType": "t4g.medium",
    "provider": "aws",
    "queriedAt": "2026-11-25T10:27:54.492180037Z",
    "quota": {
      "fitCount": 1204,
      "limit": 2421,
      "name": "Running On-Demand Standard (A, C, D, H, I, M, R, T, Z) instances",
      "perInstance": 2,
      "scope": "region",
      "unit": "vCPU",
      "used": 12
    },
    "refreshAfter": "2026-11-25T10:32:54.492180037Z",
    "region": "ap-northeast-2",
    "source": "aws:DescribeInstanceTypeOfferings",
    "status": "Available",
    "ttlSeconds": 300,
    "zones": [
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2a"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2b"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2c"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2d"
      }
    ]
  },
  "connectionName": "aws-ap-northeast-2",
  "estimatedCost": "$0.0416/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "aws-ap-northeast-2",
    "creationDate": "2026-10-23T10:41:44.000Z",
    "cspImageName": "ami-0e0a5fdc771990a70",
    "description": "Canonical, Ubuntu, 24.04, arm64 noble image",
    "details": [
      {
        "key": "Architecture",
        "value": "arm64"
      },
      {
        "key": "BlockDeviceMappings",
        "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-00c4672f066104e67,Throughput:null,VolumeSize:8,VolumeType:gp3},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
      },
      {
        "key": "BootMode",
        "value": "uefi"
      },
      {
        "key": "CreationDate",
        "value": "2026-10-23T10:41:44.000Z"
      },
      {
        "key": "DeprecationTime",
        "value": "2028-10-23T10:41:44.000Z"
      },
      {
        "key": "Description",
        "value": "Canonical, Ubuntu, 24.04, arm64 noble image"
      },
      {
        "key": "EnaSupport",
        "value": "true"
      },
      {
        "key": "Hypervisor",
        "value": "xen"
      },
      {
        "key": "ImageId",
        "value": "ami-0e0a5fdc771990a70"
      },
      {
        "key": "ImageLocation",
        "value": "amazon/ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-arm64-server-20261023"
      },
      {
        "key": "ImageOwnerAlias",
        "value": "amazon"
      },
      {
        "key": "ImageType",
        "value": "machine"
      },
      {
        "key": "Name",
        "value": "ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-arm64-server-20261023"
      },
      {
        "key": "OwnerId",
        "value": "099720109477"
      },
      {
        "key": "PlatformDetails",
        "value": "Linux/UNIX"
      },
      {
        "key": "Public",
        "value": "true"
      },
      {
        "key": "RootDeviceName",
        "value": "/dev/sda1"
      },
      {
        "key": "RootDeviceType",
        "value": "ebs"
      },
      {
        "key": "SriovNetSupport",
        "value": "simple"
      },
      {
        "key": "State",
        "value": "available"
      },
      {
        "key": "UsageOperation",
        "value": "RunInstances"
      },
      {
        "key": "VirtualizationType",
        "value": "hvm"
      }
    ],
    "fetchedTime": "2026.10.29 04:25:18 Thu",
    "id": "ami-0e0a5fdc771990a70",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "ami-0e0a5fdc771990a70",
    "namespace": "system",
    "osArchitecture": "arm64",
    "osDiskSizeGB": -1,
    "osDiskType": "ebs",
    "osDistribution": "ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-arm64-server-20261023",
    "osPlatform": "Linux/UNIX",
    "osType": "Ubuntu 24.04",
    "providerName": "aws",
    "regionList": [
      "ap-northeast-2"
    ],
    "resourceType": "image",
    "sourceCspImageName": "",
    "sourceNodeUid": "",
    "systemLabel": "",
    "uid": "tb3o288nb1ucm6n194qc"
  },
  "imageId": "ami-0e0a5fdc771990a70",
  "imageValidation": {
    "cspResourceId": "ami-0e0a5fdc771990a70",
    "isAvailable": true,
    "resourceId": "ami-0e0a5fdc771990a70",
    "resourceName": "ami-0e0a5fdc771990a70",
    "status": "Available"
  },
  "info": [
    "Provisioning history: Image 'ami-0e0a5fdc771990a70' has succeeded with spec 'aws+ap-northeast-2+t4g.medium' before (no failures)"
  ],
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "aws",
  "regionName": "ap-northeast-2",
  "specDetails": {
    "architecture": "arm64",
    "availabilityCheckedAt": "2026-11-25T10:27:54.492180037Z",
    "availabilityQuota": {
      "fitCount": 1204,
      "limit": 2421,
      "name": "Running On-Demand Standard (A, C, D, H, I, M, R, T, Z) instances",
      "perInstance": 2,
      "scope": "region",
      "unit": "vCPU",
      "used": 12
    },
    "availabilityRefreshAfter": "2026-11-25T10:32:54.492180037Z",
    "availabilityStatus": "Available",
    "availableZones": [
      "ap-northeast-2a",
      "ap-northeast-2b",
      "ap-northeast-2c",
      "ap-northeast-2d"
    ],
    "connectionName": "aws-ap-northeast-2",
    "costPerHour": 0.0416,
    "cspSpecName": "t4g.medium",
    "details": [
      {
        "key": "AutoRecoverySupported",
        "value": "true"
      },
      {
        "key": "BareMetal",
        "value": "false"
      },
      {
        "key": "BurstablePerformanceSupported",
        "value": "true"
      },
      {
        "key": "CurrentGeneration",
        "value": "true"
      },
      {
        "key": "DedicatedHostsSupported",
        "value": "false"
      },
      {
        "key": "EbsInfo",
        "value": "{AttachmentLimitType:shared,EbsCards:null,EbsOptimizedInfo:{BaselineBandwidthInMbps:347,BaselineIops:2000,BaselineThroughputInMBps:43.375,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,MaximumEbsAttachments:27,MaximumEbsCards:null,NvmeSupport:required}"
      },
      {
        "key": "FreeTierEligible",
        "value": "false"
      },
      {
        "key": "HibernationSupported",
        "value": "true"
      },
      {
        "key": "Hypervisor",
        "value": "nitro"
      },
      {
        "key": "InstanceStorageSupported",
        "value": "false"
      },
      {
        "key": "InstanceType",
        "value": "t4g.medium"
      },
      {
        "key": "MemoryInfo",
        "value": "{SizeInMiB:4096}"
      },
      {
        "key": "NetworkInfo",
        "value": "{BandwidthWeightings:null,ConnectionTrackingConfiguration:{DefaultTcpEstablishedTimeout:432000,DefaultUdpStreamTimeout:180,DefaultUdpTimeout:30},DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSrdSupported:false,EnaSupport:required,EncryptionInTransitSupported:false,FlexibleEnaQueuesSupport:unsupported,Ipv4AddressesPerInterface:6,Ipv4AddressesPerSecondaryInterface:null,Ipv6AddressesPerInterface:6,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,MaximumSecondaryNetworkInterfaces:null,NetworkCards:[{AdditionalFlexibleNetworkInterfaces:null,BaselineBandwidthInGbps:0.256,DefaultEnaQueueCountPerInterface:2,MaximumEnaQueueCount:null,MaximumEnaQueueCountPerInterface:null,MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit,PeakBandwidthInGbps:5}],NetworkPerformance:Up to 5 Gigabit,SecondaryNetworkSupported:false}"
      },
      {
        "key": "NitroEnclavesSupport",
        "value": "unsupported"
      },
      {
        "key": "NitroTpmInfo",
        "value": "{SupportedVersions:[2.0]}"
      },
      {
        "key": "NitroTpmSupport",
        "value": "supported"
      },
      {
        "key": "PhcSupport",
        "value": "unsupported"
      },
      {
        "key": "PlacementGroupInfo",
        "value": "{SupportedStrategies:[partition,spread]}"
      },
      {
        "key": "ProcessorInfo",
        "value": "{Manufacturer:AWS,SupportedArchitectures:[arm64],SupportedFeatures:null,SustainedClockSpeedInGhz:2.5}"
      },
      {
        "key": "RebootMigrationSupport",
        "value": "supported"
      },
      {
        "key": "SupportedBootModes",
        "value": "uefi"
      },
      {
        "key": "SupportedInRegion",
        "value": "true"
      },
      {
        "key": "SupportedRootDeviceTypes",
        "value": "ebs"
      },
      {
        "key": "SupportedUsageClasses",
        "value": "on-demand; spot"
      },
      {
        "key": "SupportedVirtualizationTypes",
        "value": "hvm"
      },
      {
        "key": "VCpuInfo",
        "value": "{DefaultCores:2,DefaultThreadsPerCore:1,DefaultVCpus:2,ValidCores:[1,2],ValidThreadsPerCore:[1]}"
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
    "id": "aws+ap-northeast-2+t4g.medium",
    "infraType": "node",
    "memoryGiB": 4,
    "name": "aws+ap-northeast-2+t4g.medium",
    "namespace": "system",
    "providerName": "aws",
    "regionLatitude": 37.36,
    "regionLongitude": 126.78,
    "regionName": "ap-northeast-2",
    "rootDiskSize": -1,
    "rootDiskType": "",
    "systemLabel": "auto-gen",
    "uid": "tbep68pth62epiftrjg9",
    "vCPU": 2
  },
  "specId": "aws+ap-northeast-2+t4g.medium",
  "specValidation": {
    "cspResourceId": "t4g.medium",
    "isAvailable": true,
    "resourceId": "aws+ap-northeast-2+t4g.medium",
    "resourceName": "t4g.medium",
    "status": "Available"
  },
  "status": "OK",
  "suggestedZone": "ap-northeast-2a"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 3.295s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.10.0.0/16",
  "connectionName": "aws-ap-northeast-2",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-mariadb-aws",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.10.1.0/24",
      "name": "subnet-1",
      "zone": "ap-northeast-2a"
    },
    {
      "ipv4_CIDR": "10.10.2.0/24",
      "name": "subnet-2",
      "zone": "ap-northeast-2c"
    }
  ]
}
```
```json
// Response Body
{
  "associatedObjectList": null,
  "cidrBlock": "10.10.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T10:27:58Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T10:27:58Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-11-25T10:27:58Z",
      "reason": "AllReady",
      "status": "True",
      "type": "ChildrenReady"
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
  "cspResourceId": "vpc-0df32b14dbcdda989",
  "cspResourceName": "tbkn4dvv4rtl3f2eevp9",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-mariadb-aws",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "CidrBlock",
      "value": "10.10.0.0/16"
    },
    {
      "key": "CidrBlockAssociationSet",
      "value": "{AssociationId:vpc-cidr-assoc-0feec249e2801e2ca,CidrBlock:10.10.0.0/16,CidrBlockState:{State:associated,StatusMessage:null}}"
    },
    {
      "key": "DhcpOptionsId",
      "value": "dopt-fa6b9492"
    },
    {
      "key": "InstanceTenancy",
      "value": "default"
    },
    {
      "key": "IsDefault",
      "value": "false"
    },
    {
      "key": "OwnerId",
      "value": "635484366616"
    },
    {
      "key": "State",
      "value": "pending"
    },
    {
      "key": "Tags",
      "value": "{Key:Name,Value:tbkn4dvv4rtl3f2eevp9}"
    },
    {
      "key": "VpcId",
      "value": "vpc-0df32b14dbcdda989"
    }
  ],
  "name": "test-rdbms-vnet-mariadb-aws",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:27:58Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:27:58Z",
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
      "cspResourceId": "subnet-0332b3b067d7b2be3",
      "cspResourceName": "tbg20p2dsfjcpt29tate",
      "cspVNetId": "vpc-0df32b14dbcdda989",
      "cspVNetName": "tbkn4dvv4rtl3f2eevp9",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.10.1.0/24",
      "keyValueList": [
        {
          "key": "AssignIpv6AddressOnCreation",
          "value": "false"
        },
        {
          "key": "AvailabilityZone",
          "value": "ap-northeast-2a"
        },
        {
          "key": "AvailabilityZoneId",
          "value": "apne2-az1"
        },
        {
          "key": "AvailableIpAddressCount",
          "value": "251"
        },
        {
          "key": "CidrBlock",
          "value": "10.10.1.0/24"
        },
        {
          "key": "DefaultForAz",
          "value": "false"
        },
        {
          "key": "MapCustomerOwnedIpOnLaunch",
          "value": "false"
        },
        {
          "key": "MapPublicIpOnLaunch",
          "value": "false"
        },
        {
          "key": "OwnerId",
          "value": "635484366616"
        },
        {
          "key": "State",
          "value": "available"
        },
        {
          "key": "SubnetArn",
          "value": "arn:aws:ec2:ap-northeast-2:635484366616:subnet/subnet-0332b3b067d7b2be3"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0332b3b067d7b2be3"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbg20p2dsfjcpt29tate}"
        },
        {
          "key": "VpcId",
          "value": "vpc-0df32b14dbcdda989"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbg20p2dsfjcpt29tate",
      "zone": "ap-northeast-2a"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:27:58Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:27:58Z",
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
      "cspResourceId": "subnet-0fcd6322d6b4d1720",
      "cspResourceName": "tb8ojnbshcaqum9vpndt",
      "cspVNetId": "vpc-0df32b14dbcdda989",
      "cspVNetName": "tbkn4dvv4rtl3f2eevp9",
      "description": "",
      "id": "subnet-2",
      "ipv4_CIDR": "10.10.2.0/24",
      "keyValueList": [
        {
          "key": "AssignIpv6AddressOnCreation",
          "value": "false"
        },
        {
          "key": "AvailabilityZone",
          "value": "ap-northeast-2c"
        },
        {
          "key": "AvailabilityZoneId",
          "value": "apne2-az3"
        },
        {
          "key": "AvailableIpAddressCount",
          "value": "251"
        },
        {
          "key": "CidrBlock",
          "value": "10.10.2.0/24"
        },
        {
          "key": "DefaultForAz",
          "value": "false"
        },
        {
          "key": "MapCustomerOwnedIpOnLaunch",
          "value": "false"
        },
        {
          "key": "MapPublicIpOnLaunch",
          "value": "false"
        },
        {
          "key": "OwnerId",
          "value": "635484366616"
        },
        {
          "key": "State",
          "value": "available"
        },
        {
          "key": "SubnetArn",
          "value": "arn:aws:ec2:ap-northeast-2:635484366616:subnet/subnet-0fcd6322d6b4d1720"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0fcd6322d6b4d1720"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tb8ojnbshcaqum9vpndt}"
        },
        {
          "key": "VpcId",
          "value": "vpc-0df32b14dbcdda989"
        }
      ],
      "name": "subnet-2",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb8ojnbshcaqum9vpndt",
      "zone": "ap-northeast-2c"
    }
  ],
  "systemLabel": "",
  "uid": "tbkn4dvv4rtl3f2eevp9"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 1.408s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/securityGroup`
```json
// Request Body
{
  "connectionName": "aws-ap-northeast-2",
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
  "name": "test-rdbms-sg-mariadb-aws",
  "vNetId": "test-rdbms-vnet-mariadb-aws"
}
```
```json
// Response Body
{
  "associatedObjectList": [],
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
  "cspResourceId": "sg-0dec51a4099d5aa2c",
  "cspResourceName": "tbhbqd0ok297ddpvgpsm",
  "description": "Pre-requisite SecurityGroup for CM-Beetle RDBMS test",
  "firewallRules": [
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "22",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "inbound",
      "Port": "3306",
      "Protocol": "TCP"
    },
    {
      "CIDR": "0.0.0.0/0",
      "Direction": "outbound",
      "Port": "",
      "Protocol": "ALL"
    }
  ],
  "id": "test-rdbms-sg-mariadb-aws",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "GroupName",
      "value": "tbhbqd0ok297ddpvgpsm"
    },
    {
      "key": "VpcID",
      "value": "vpc-0df32b14dbcdda989"
    },
    {
      "key": "OwnerID",
      "value": "635484366616"
    },
    {
      "key": "Description",
      "value": "tbhbqd0ok297ddpvgpsm"
    }
  ],
  "name": "test-rdbms-sg-mariadb-aws",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbhbqd0ok297ddpvgpsm",
  "vNetId": "test-rdbms-vnet-mariadb-aws"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 7ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/support?providerName=aws`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "aws": {
      "dbOperationMethod": "sqlFallback",
      "storageTypeSelectable": true,
      "supported": true,
      "supportedDBEngines": [
        "mysql",
        "mariadb"
      ],
      "supportsTag": true
    }
  }
}
```

### 5. Beetle GET RDBMS Capability [✅ SUCCESS]
- **Duration:** 1m5.974s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=aws-ap-northeast-2&dbEngine=mariadb`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "0-35",
    "connectionName": "aws-ap-northeast-2",
    "dbEngine": "mariadb",
    "dbInstanceSpecOptions": [
      "db.m5.12xlarge",
      "db.m5.16xlarge",
      "db.m5.24xlarge",
      "db.m5.2xlarge",
      "db.m5.4xlarge",
      "db.m5.8xlarge",
      "db.m5.large",
      "db.m5.xlarge",
      "db.m5d.12xlarge",
      "db.m5d.16xlarge",
      "db.m5d.24xlarge",
      "db.m5d.2xlarge",
      "db.m5d.4xlarge",
      "db.m5d.8xlarge",
      "db.m5d.large",
      "db.m5d.xlarge",
      "db.m6g.12xlarge",
      "db.m6g.16xlarge",
      "db.m6g.2xlarge",
      "db.m6g.4xlarge",
      "db.m6g.8xlarge",
      "db.m6g.large",
      "db.m6g.xlarge",
      "db.m6gd.12xlarge",
      "db.m6gd.16xlarge",
      "db.m6gd.2xlarge",
      "db.m6gd.4xlarge",
      "db.m6gd.8xlarge",
      "db.m6gd.large",
      "db.m6gd.xlarge",
      "db.m6i.12xlarge",
      "db.m6i.16xlarge",
      "db.m6i.24xlarge",
      "db.m6i.2xlarge",
      "db.m6i.32xlarge",
      "db.m6i.4xlarge",
      "db.m6i.8xlarge",
      "db.m6i.large",
      "db.m6i.xlarge",
      "db.m7g.12xlarge",
      "db.m7g.16xlarge",
      "db.m7g.2xlarge",
      "db.m7g.4xlarge",
      "db.m7g.8xlarge",
      "db.m7g.large",
      "db.m7g.xlarge",
      "db.m7i.12xlarge",
      "db.m7i.16xlarge",
      "db.m7i.24xlarge",
      "db.m7i.2xlarge",
      "db.m7i.48xlarge",
      "db.m7i.4xlarge",
      "db.m7i.8xlarge",
      "db.m7i.large",
      "db.m7i.xlarge",
      "db.m8g.12xlarge",
      "db.m8g.16xlarge",
      "db.m8g.24xlarge",
      "db.m8g.2xlarge",
      "db.m8g.48xlarge",
      "db.m8g.4xlarge",
      "db.m8g.8xlarge",
      "db.m8g.large",
      "db.m8g.xlarge",
      "db.r5.12xlarge",
      "db.r5.16xlarge",
      "db.r5.24xlarge",
      "db.r5.2xlarge",
      "db.r5.4xlarge",
      "db.r5.8xlarge",
      "db.r5.large",
      "db.r5.xlarge",
      "db.r5d.12xlarge",
      "db.r5d.16xlarge",
      "db.r5d.24xlarge",
      "db.r5d.2xlarge",
      "db.r5d.4xlarge",
      "db.r5d.8xlarge",
      "db.r5d.large",
      "db.r5d.xlarge",
      "db.r6g.12xlarge",
      "db.r6g.16xlarge",
      "db.r6g.2xlarge",
      "db.r6g.4xlarge",
      "db.r6g.8xlarge",
      "db.r6g.large",
      "db.r6g.xlarge",
      "db.r6i.12xlarge",
      "db.r6i.16xlarge",
      "db.r6i.24xlarge",
      "db.r6i.2xlarge",
      "db.r6i.32xlarge",
      "db.r6i.4xlarge",
      "db.r6i.8xlarge",
      "db.r6i.large",
      "db.r6i.xlarge",
      "db.r7g.12xlarge",
      "db.r7g.16xlarge",
      "db.r7g.2xlarge",
      "db.r7g.4xlarge",
      "db.r7g.8xlarge",
      "db.r7g.large",
      "db.r7g.xlarge",
      "db.r7i.12xlarge",
      "db.r7i.16xlarge",
      "db.r7i.24xlarge",
      "db.r7i.2xlarge",
      "db.r7i.48xlarge",
      "db.r7i.4xlarge",
      "db.r7i.8xlarge",
      "db.r7i.large",
      "db.r7i.xlarge",
      "db.r8g.12xlarge",
      "db.r8g.16xlarge",
      "db.r8g.24xlarge",
      "db.r8g.2xlarge",
      "db.r8g.48xlarge",
      "db.r8g.4xlarge",
      "db.r8g.8xlarge",
      "db.r8g.large",
      "db.r8g.xlarge",
      "db.r8gd.12xlarge",
      "db.r8gd.16xlarge",
      "db.r8gd.24xlarge",
      "db.r8gd.2xlarge",
      "db.r8gd.48xlarge",
      "db.r8gd.4xlarge",
      "db.r8gd.8xlarge",
      "db.r8gd.large",
      "db.r8gd.xlarge",
      "db.t3.2xlarge",
      "db.t3.large",
      "db.t3.medium",
      "db.t3.micro",
      "db.t3.small",
      "db.t3.xlarge",
      "db.t4g.2xlarge",
      "db.t4g.large",
      "db.t4g.medium",
      "db.t4g.micro",
      "db.t4g.small",
      "db.t4g.xlarge"
    ],
    "dbInstanceSpecs": [
      {
        "memSizeMiB": "196608",
        "name": "db.m5.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m5.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.m5.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m5.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m5.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m5.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m5.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m5.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m5d.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m5d.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.m5d.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m5d.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m5d.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m5d.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m5d.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m5d.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m6g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m6g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m6g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m6g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m6g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m6g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m6g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m6gd.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m6gd.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m6gd.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m6gd.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m6gd.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m6gd.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m6gd.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m6i.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m6i.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.m6i.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m6i.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.m6i.32xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "128"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m6i.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m6i.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m6i.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m6i.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m7g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m7g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m7g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m7g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m7g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m7g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m7g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m7i.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m7i.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.m7i.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m7i.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.m7i.48xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "192"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m7i.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m7i.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m7i.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m7i.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "196608",
        "name": "db.m8g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.m8g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.m8g.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.m8g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.m8g.48xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.7",
        "vCpuCount": "192"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.m8g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.m8g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.m8g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.m8g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r5.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r5.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r5.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r5.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r5.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r5.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r5.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r5.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r5d.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r5d.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r5d.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r5d.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r5d.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r5d.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r5d.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r5d.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.1",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r6g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r6g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r6g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r6g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r6g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r6g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r6g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r6i.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r6i.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r6i.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r6i.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "1048576",
        "name": "db.r6i.32xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "128"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r6i.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r6i.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r6i.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r6i.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r7g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r7g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r7g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r7g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r7g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r7g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r7g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.6",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r7i.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r7i.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r7i.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r7i.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "1572864",
        "name": "db.r7i.48xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "192"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r7i.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r7i.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r7i.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r7i.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "3.2",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r8g.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r8g.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r8g.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r8g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "1572864",
        "name": "db.r8g.48xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.7",
        "vCpuCount": "192"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r8g.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r8g.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r8g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r8g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "393216",
        "name": "db.r8gd.12xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "48"
      },
      {
        "memSizeMiB": "524288",
        "name": "db.r8gd.16xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "64"
      },
      {
        "memSizeMiB": "786432",
        "name": "db.r8gd.24xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "96"
      },
      {
        "memSizeMiB": "65536",
        "name": "db.r8gd.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "1572864",
        "name": "db.r8gd.48xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.7",
        "vCpuCount": "192"
      },
      {
        "memSizeMiB": "131072",
        "name": "db.r8gd.4xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "16"
      },
      {
        "memSizeMiB": "262144",
        "name": "db.r8gd.8xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "32"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.r8gd.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.r8gd.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.8",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.t3.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.t3.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "4096",
        "name": "db.t3.medium",
        "storageSizeRangeGB": {
          "max": 17592,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "1024",
        "name": "db.t3.micro",
        "storageSizeRangeGB": {
          "max": 6597,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "2048",
        "name": "db.t3.small",
        "storageSizeRangeGB": {
          "max": 17592,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.t3.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "4"
      },
      {
        "memSizeMiB": "32768",
        "name": "db.t4g.2xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "8"
      },
      {
        "memSizeMiB": "8192",
        "name": "db.t4g.large",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "4096",
        "name": "db.t4g.medium",
        "storageSizeRangeGB": {
          "max": 17592,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "1024",
        "name": "db.t4g.micro",
        "storageSizeRangeGB": {
          "max": 6597,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "2048",
        "name": "db.t4g.small",
        "storageSizeRangeGB": {
          "max": 17592,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "2"
      },
      {
        "memSizeMiB": "16384",
        "name": "db.t4g.xlarge",
        "storageSizeRangeGB": {
          "max": 70369,
          "min": 5
        },
        "vCpuClockGHz": "2.5",
        "vCpuCount": "4"
      }
    ],
    "dbOperationMethod": "",
    "defaultStorageType": "",
    "liveSupportedEngines": [
      "mysql",
      "mariadb"
    ],
    "notes": {
      "storageTypes": [
        {
          "constraints": "Minimum 20GB storage.",
          "description": "Legacy general-purpose SSD storage. Consider gp3 for better cost/performance.",
          "displayName": "General Purpose SSD (gp2)",
          "maxSizeGB": 70369,
          "minSizeGB": 20,
          "recommendationLevel": "legacy",
          "storageType": "gp2"
        },
        {
          "constraints": "Minimum 20GB storage.",
          "description": "Latest generation general-purpose SSD with 3,000 baseline IOPS and 125 MiB/s throughput. Recommended for most workloads.",
          "displayName": "General Purpose SSD (gp3)",
          "maxSizeGB": 70369,
          "minSizeGB": 20,
          "recommendationLevel": "recommended",
          "recommended": true,
          "storageType": "gp3"
        },
        {
          "constraints": "Requires 'iops' parameter (range: 1000-64000). Minimum 100GB storage.",
          "description": "High-performance SSD for I/O-intensive workloads. Requires 'iops' parameter (e.g., '3000'). Minimum 100 GB storage.",
          "displayName": "Provisioned IOPS SSD (io1)",
          "iopsRange": {
            "max": 64000,
            "min": 1000
          },
          "maxSizeGB": 70369,
          "minSizeGB": 100,
          "recommendationLevel": "premium",
          "requiresIops": true,
          "storageType": "io1"
        },
        {
          "constraints": "Requires 'iops' parameter (range: 1000-256000). Minimum 100GB storage.",
          "description": "Next-generation high-performance SSD with higher durability (99.999%). Requires 'iops' parameter. Minimum 100 GB storage.",
          "displayName": "Provisioned IOPS SSD (io2)",
          "iopsRange": {
            "max": 256000,
            "min": 1000
          },
          "maxSizeGB": 70369,
          "minSizeGB": 100,
          "recommendationLevel": "premium",
          "requiresIops": true,
          "storageType": "io2"
        }
      ]
    },
    "providerName": "aws",
    "regionName": "ap-northeast-2",
    "requiresSecurityGroup": true,
    "requiresSubnet": true,
    "storageSizeRangeGB": {
      "max": 70369,
      "min": 5
    },
    "storageTypeOptions": [
      "gp2",
      "gp3",
      "io1",
      "io2"
    ],
    "supportedVersions": [
      "10.5",
      "10.6",
      "10.11",
      "11.4",
      "11.8",
      "12.3"
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
- **Duration:** 6ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
  },
  "sourceRDBMSInstances": [
    {
      "dbEngine": {
        "engine": "mariadb",
        "engineVersion": "10.6",
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
    "adminUserName": "root",
    "backupRetentionDays": 7,
    "highAvailability": false,
    "publicAccess": true
  }
}
```
```json
// Response Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for aws (ap-northeast-2)",
  "status": "recommended",
  "targetCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "root",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "charset": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "db.t3.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-mariadb-aws",
      "securityGroupIds": [
        "test-rdbms-sg-mariadb-aws"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "gp3",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "vNetId": "test-rdbms-vnet-mariadb-aws"
    }
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 8.314s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "root",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "aws-ap-northeast-2",
  "dbEngine": "mariadb",
  "dbEngineVersion": "10.6",
  "dbInstanceSpec": "db.t3.medium",
  "name": "rdbms-mariadb-aws",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-mariadb-aws"
  ],
  "storageSize": 100,
  "storageType": "gp3",
  "subnetIds": [
    "subnet-1",
    "subnet-2"
  ],
  "vNetId": "test-rdbms-vnet-mariadb-aws"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "root",
    "adminUserPassword": "******",
    "connectionName": "aws-ap-northeast-2",
    "dbEngine": "mariadb",
    "dbEngineVersion": "10.6",
    "dbInstanceSpec": "db.t3.medium",
    "name": "rdbms-mariadb-aws",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-mariadb-aws"
    ],
    "storageSize": 100,
    "storageType": "gp3",
    "subnetIds": [
      "subnet-1",
      "subnet-2"
    ],
    "vNetId": "test-rdbms-vnet-mariadb-aws"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 9m56.784s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms?nameSeed=test`
```json
// Request Body
{
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for aws (ap-northeast-2)",
  "status": "recommended",
  "targetCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
  },
  "targetRDBMSInstances": [
    {
      "adminUserName": "root",
      "adminUserPassword": "******",
      "backupRetentionDays": 7,
      "databases": [
        {
          "charset": "utf8mb4",
          "collation": "utf8mb4_unicode_ci",
          "databaseName": "sampledb"
        }
      ],
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "db.t3.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-mariadb-aws",
      "securityGroupIds": [
        "test-rdbms-sg-mariadb-aws"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "gp3",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "vNetId": "test-rdbms-vnet-mariadb-aws"
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
- **Duration:** 12ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-aws`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "15:06-15:36",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T10:38:32Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T10:38:32Z",
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
  "cspResourceId": "tb55rm4aqne7bvtmh109",
  "cspResourceName": "tb55rm4aqne7bvtmh109",
  "dbEngine": "mariadb",
  "dbEngineVersion": "10.6.28",
  "dbInstanceSpec": "db.t3.medium",
  "dbInstanceType": "Primary",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "tb55rm4aqne7bvtmh109.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com:3306",
  "highAvailability": false,
  "id": "test-rdbms-mariadb-aws",
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
      "value": "arn:aws:rds:ap-northeast-2:635484366616:db:tb55rm4aqne7bvtmh109"
    },
    {
      "key": "DBInstanceClass",
      "value": "db.t3.medium"
    },
    {
      "key": "DBInstanceIdentifier",
      "value": "tb55rm4aqne7bvtmh109"
    },
    {
      "key": "DBInstanceStatus",
      "value": "available"
    },
    {
      "key": "DBParameterGroups",
      "value": "{DBParameterGroupName:default.mariadb10.6,ParameterApplyStatus:in-sync}"
    },
    {
      "key": "DBSubnetGroup",
      "value": "{DBSubnetGroupArn:null,DBSubnetGroupDescription:CB-Spider RDBMS subnet group for cb-spider-tb55rm4aqne7bvtmh109,DBSubnetGroupName:cb-spider-tb55rm4aqne7bvtmh109,SubnetGroupStatus:Complete,Subnets:[{SubnetAvailabilityZone:{Name:ap-northeast-2a},SubnetIdentifier:subnet-0332b3b067d7b2be3,SubnetOutpost:{Arn:null},SubnetStatus:Active},{SubnetAvailabilityZone:{Name:ap-northeast-2c},SubnetIdentifier:subnet-0fcd6322d6b4d1720,SubnetOutpost:{Arn:null},SubnetStatus:Active}],VpcId:vpc-0df32b14dbcdda989}"
    },
    {
      "key": "DbInstancePort",
      "value": "0"
    },
    {
      "key": "DbiResourceId",
      "value": "db-6S54HILSY73L2WFC6GXJJLPUJM"
    },
    {
      "key": "DeletionProtection",
      "value": "false"
    },
    {
      "key": "Endpoint",
      "value": "{Address:tb55rm4aqne7bvtmh109.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com,HostedZoneId:ZLA2NUCOLGUUR,Port:3306}"
    },
    {
      "key": "Engine",
      "value": "mariadb"
    },
    {
      "key": "EngineVersion",
      "value": "10.6.28"
    },
    {
      "key": "IAMDatabaseAuthenticationEnabled",
      "value": "false"
    },
    {
      "key": "InstanceCreateTime",
      "value": "2026-11-25T10:32:42.286Z"
    },
    {
      "key": "Iops",
      "value": "3000"
    },
    {
      "key": "LatestRestorableTime",
      "value": "2026-11-25T10:34:00.588Z"
    },
    {
      "key": "LicenseModel",
      "value": "general-public-license"
    },
    {
      "key": "MasterUsername",
      "value": "root"
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
      "value": "{OptionGroupName:default:mariadb-10-6,Status:in-sync}"
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
      "value": "15:06-15:36"
    },
    {
      "key": "PreferredMaintenanceWindow",
      "value": "sun:16:40-sun:17:10"
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
      "value": "{Key:Name,Value:tb55rm4aqne7bvtmh109}"
    },
    {
      "key": "VpcSecurityGroups",
      "value": "{Status:active,VpcSecurityGroupId:sg-0dec51a4099d5aa2c}"
    }
  ],
  "name": "test-rdbms-mariadb-aws",
  "publicAccess": true,
  "resourceType": "rdbms",
  "securityGroupIds": [
    "test-rdbms-sg-mariadb-aws"
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
      "value": "tb55rm4aqne7bvtmh109"
    }
  ],
  "uid": "tb55rm4aqne7bvtmh109",
  "vNetId": "test-rdbms-vnet-mariadb-aws"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 21ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "backupTime": "09:00Z-10:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:34:16Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:31:22Z",
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
      "cspResourceId": "rm-mj7yv170t22mnj1bz",
      "cspResourceName": "tbg504bm99cilsq7lbb2",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6",
      "dbInstanceSpec": "mariadb.n2.medium.2c",
      "dbInstanceType": "HighAvailability",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "43.108.66.115:3306",
      "highAvailability": true,
      "id": "test-rdbms-mariadb-alibaba",
      "keyValueList": [
        {
          "key": "Engine",
          "value": "MariaDB"
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
          "value": "rm-mj7yv170t22mnj1bz-20261125182909"
        },
        {
          "key": "DBInstanceNetType",
          "value": "Intranet"
        },
        {
          "key": "DBInstanceClass",
          "value": "mariadb.n2.medium.2c"
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
          "value": "rm-mj7yv170t22mnj1bz"
        },
        {
          "key": "VpcId",
          "value": "vpc-mj7fosavcqa9tajq8qsyy"
        },
        {
          "key": "DBInstanceMemory",
          "value": "4096"
        },
        {
          "key": "DBInstanceStorageType",
          "value": "cloud_essd"
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
          "value": "10.6"
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
          "value": "mariadb106_20260531"
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
          "key": "MasterZone",
          "value": "ap-northeast-2a"
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
          "value": "HighAvailability"
        },
        {
          "key": "InsId",
          "value": "1"
        },
        {
          "key": "DBInstanceDescription",
          "value": "tbg504bm99cilsq7lbb2"
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
          "value": "0"
        },
        {
          "key": "ResourceGroupId",
          "value": "rg-acfnvekhilw5kmy"
        },
        {
          "key": "DispenseMode",
          "value": "MultiAVZDispenseMode"
        },
        {
          "key": "TempUpgradeRecoveryMemory",
          "value": "0"
        },
        {
          "key": "ProxyType",
          "value": "0"
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
          "value": "vsw-mj7mlhi0t2mhxca0s54s7"
        },
        {
          "key": "CreationTime",
          "value": "2026-11-25T10:29:13Z"
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
          "value": "rm-mj7yv170t22mnj1bz.mariadb.ap-northeast-2.rds.aliyuncs.com"
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
          "value": "{SlaveZone:[{ZoneId:ap-northeast-2a}]}"
        }
      ],
      "name": "test-rdbms-mariadb-alibaba",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-mariadb-alibaba"
      ],
      "status": "Deleting",
      "storageSize": 100,
      "storageType": "cloud_essd",
      "subnetIds": [
        "subnet-1"
      ],
      "uid": "tbg504bm99cilsq7lbb2",
      "vNetId": "test-rdbms-vnet-mariadb-alibaba"
    },
    {
      "backupRetentionDays": 7,
      "backupTime": "15:06-15:36",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:38:32Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:38:32Z",
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
      "cspResourceId": "tb55rm4aqne7bvtmh109",
      "cspResourceName": "tb55rm4aqne7bvtmh109",
      "dbEngine": "mariadb",
      "dbEngineVersion": "10.6.28",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "tb55rm4aqne7bvtmh109.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-mariadb-aws",
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
          "value": "arn:aws:rds:ap-northeast-2:635484366616:db:tb55rm4aqne7bvtmh109"
        },
        {
          "key": "DBInstanceClass",
          "value": "db.t3.medium"
        },
        {
          "key": "DBInstanceIdentifier",
          "value": "tb55rm4aqne7bvtmh109"
        },
        {
          "key": "DBInstanceStatus",
          "value": "available"
        },
        {
          "key": "DBParameterGroups",
          "value": "{DBParameterGroupName:default.mariadb10.6,ParameterApplyStatus:in-sync}"
        },
        {
          "key": "DBSubnetGroup",
          "value": "{DBSubnetGroupArn:null,DBSubnetGroupDescription:CB-Spider RDBMS subnet group for cb-spider-tb55rm4aqne7bvtmh109,DBSubnetGroupName:cb-spider-tb55rm4aqne7bvtmh109,SubnetGroupStatus:Complete,Subnets:[{SubnetAvailabilityZone:{Name:ap-northeast-2a},SubnetIdentifier:subnet-0332b3b067d7b2be3,SubnetOutpost:{Arn:null},SubnetStatus:Active},{SubnetAvailabilityZone:{Name:ap-northeast-2c},SubnetIdentifier:subnet-0fcd6322d6b4d1720,SubnetOutpost:{Arn:null},SubnetStatus:Active}],VpcId:vpc-0df32b14dbcdda989}"
        },
        {
          "key": "DbInstancePort",
          "value": "0"
        },
        {
          "key": "DbiResourceId",
          "value": "db-6S54HILSY73L2WFC6GXJJLPUJM"
        },
        {
          "key": "DeletionProtection",
          "value": "false"
        },
        {
          "key": "Endpoint",
          "value": "{Address:tb55rm4aqne7bvtmh109.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com,HostedZoneId:ZLA2NUCOLGUUR,Port:3306}"
        },
        {
          "key": "Engine",
          "value": "mariadb"
        },
        {
          "key": "EngineVersion",
          "value": "10.6.28"
        },
        {
          "key": "IAMDatabaseAuthenticationEnabled",
          "value": "false"
        },
        {
          "key": "InstanceCreateTime",
          "value": "2026-11-25T10:32:42.286Z"
        },
        {
          "key": "Iops",
          "value": "3000"
        },
        {
          "key": "LatestRestorableTime",
          "value": "2026-11-25T10:34:00.588Z"
        },
        {
          "key": "LicenseModel",
          "value": "general-public-license"
        },
        {
          "key": "MasterUsername",
          "value": "root"
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
          "value": "{OptionGroupName:default:mariadb-10-6,Status:in-sync}"
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
          "value": "15:06-15:36"
        },
        {
          "key": "PreferredMaintenanceWindow",
          "value": "sun:16:40-sun:17:10"
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
          "value": "{Key:Name,Value:tb55rm4aqne7bvtmh109}"
        },
        {
          "key": "VpcSecurityGroups",
          "value": "{Status:active,VpcSecurityGroupId:sg-0dec51a4099d5aa2c}"
        }
      ],
      "name": "test-rdbms-mariadb-aws",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-mariadb-aws"
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
          "value": "tb55rm4aqne7bvtmh109"
        }
      ],
      "uid": "tb55rm4aqne7bvtmh109",
      "vNetId": "test-rdbms-vnet-mariadb-aws"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:28:37Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:28:37Z",
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
      "uid": "tb5jnso3laaus5n58f8i",
      "vNetId": "test-rdbms-vnet-mariadb-nhn"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 488ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-aws/database`
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
- **Duration:** 350ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-aws/database`
```json
// Response Body
{
  "databases": [
    "information_schema",
    "innodb",
    "mysql",
    "performance_schema",
    "sampledb",
    "sampledb_dyn",
    "sys"
  ]
}
```

### 13. Beetle GET Secure Transport Info [✅ SUCCESS]
- **Duration:** 410ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-aws/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": false,
    "issuer": "CN=Amazon RDS ap-northeast-2 Root CA RSA2048 G1,OU=Amazon RDS,O=Amazon Web Services\\, Inc.,L=Seattle,ST=WA,C=US",
    "notAfter": "2032-12-08T19:00:47Z",
    "pem": "-----BEGIN CERTIFICATE-----\nMIIENjCCAx6gAwIBAgIRANDpdGHcX1ynlOSQsstVSoMwDQYJKoZIhvcNAQELBQAw\ngZwxCzAJBgNVBAYTAlVTMSIwIAYDVQQKDBlBbWF6b24gV2ViIFNlcnZpY2VzLCBJ\nbmMuMRMwEQYDVQQLDApBbWF6b24gUkRTMQswCQYDVQQIDAJXQTE1MDMGA1UEAwws\nQW1hem9uIFJEUyBhcC1ub3J0aGVhc3QtMiBSb290IENBIFJTQTIwNDggRzExEDAO\nBgNVBAcMB1NlYXR0bGUwHhcNMjIxMjA4MTgwMDQ3WhcNMzIxMjA4MTkwMDQ3WjCB\npzELMAkGA1UEBhMCVVMxIjAgBgNVBAoMGUFtYXpvbiBXZWIgU2VydmljZXMsIElu\nYy4xEzARBgNVBAsMCkFtYXpvbiBSRFMxCzAJBgNVBAgMAldBMUAwPgYDVQQDDDdB\nbWF6b24gUkRTIGFwLW5vcnRoZWFzdC0yIFN1Ym9yZGluYXRlIENBIFJTQTIwNDgg\nRzEuQS4yMRAwDgYDVQQHDAdTZWF0dGxlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A\nMIIBCgKCAQEAtjQWiONj80aWDZxXoaUrMqWW9r8PmuZUN9Hd++7uJzwkjFkOeznu\nANOwsfF/6J599LKWkacH77qK38bV3QUawED4zcqJGNQbhLB1530w9IwevOXzueD8\naon5My/AM22ukvA47Jb/MVTvUEcqSFtsfK0Jp4RByfqj5GL61uhe2QBu7G2SFSAs\ndfeNgk95Q32MiIlizAz+zqlVshPBejs51SNxvqeRiF21oedAFWz/xmveBnVmriRt\nKKZdWWeY8eHOsxO/Ej1GDpX2P4qDzuiaCDTCa6TaEmDjIrQYA78/nycZop4HrTm9\nHGAkGyq+HeBUGCUgB+Aj8Lr3rNagezSJZwIDAQABo2YwZDASBgNVHRMBAf8ECDAG\nAQH/AgEAMB8GA1UdIwQYMBaAFCSBNYWw2DrADcdtuKENDi+m3UTbMB0GA1UdDgQW\nBBQtSHOI2CXyhkBtUOlIiXh9urm4hTAOBgNVHQ8BAf8EBAMCAYYwDQYJKoZIhvcN\nAQELBQADggEBAHKY/bnyQ66j7Dd+T4pav+0AiMxixpI5tvez2DpYKimCzoz3gcoc\npIWMLZFwrCsoFRUJ7FFNp/GHwDfI7RB8tO1yFoy5y9IZHos14zHl8H+gSSM/SZeq\nhBEk56o0kpfhEI+cffCejRWOutthi3jesoXNDfchP6sClTI12kS9W3gY1t5BLZZw\nIJdqOIbDpm/xQKs0LPJfWG89xYW79DvUMTj3fiIgyb3bGmmhr/y+ioym3i/xNyMo\nD0bv4+eMgJt6Rcz89MX6Ckic7jct3Y9r6lXcFKj3+JwkYQ/TtJ8ggMi1eZm9rH8r\nPqCE8yDY6GmPZP9vatKB/WQdcOi3XnUaGCM=\n-----END CERTIFICATE-----\n",
    "subject": "CN=Amazon RDS ap-northeast-2 Subordinate CA RSA2048 G1.A.2,OU=Amazon RDS,O=Amazon Web Services\\, Inc.,L=Seattle,ST=WA,C=US"
  },
  "enforced": false,
  "engine": "mariadb",
  "recommendedSSLMode": "VERIFY_IDENTITY",
  "requireSecureTransport": "OFF",
  "rules": "",
  "tlsCipher": "TLS_AES_128_GCM_SHA256",
  "tlsInUse": true
}
```

### 14. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 25ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m50.314s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 13.294s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-aws/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m21.278s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-mariadb-aws?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 839ms

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 10.1s

