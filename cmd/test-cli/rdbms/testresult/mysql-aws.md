# Managed RDBMS (MySQL) Test Report: AWS (ap-northeast-2)

- **Test Case:** AWS AP-Northeast-2 (Seoul) MySQL Test
- **Access Mode:** Public (publicAccess=true)
- **Date & Time:** 2026-11-25 18:54:32
- **Namespace:** `default`
- **Total Duration:** 18m31.647s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** AWS
- **Target Region:** `ap-northeast-2`
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
- **Duration:** 3.693s
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
    "queriedAt": "2026-11-25T09:54:33.21011855Z",
    "quota": {
      "fitCount": 1204,
      "limit": 2421,
      "name": "Running On-Demand Standard (A, C, D, H, I, M, R, T, Z) instances",
      "perInstance": 2,
      "scope": "region",
      "unit": "vCPU",
      "used": 12
    },
    "refreshAfter": "2026-11-25T09:59:33.21011855Z",
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
    "Provisioning history: Image 'ami-0e0a5fdc771990a70' + Spec 'aws+ap-northeast-2+t4g.medium' is a new combination (no history)"
  ],
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
  "providerName": "aws",
  "regionName": "ap-northeast-2",
  "specDetails": {
    "architecture": "arm64",
    "availabilityCheckedAt": "2026-11-25T09:54:33.21011855Z",
    "availabilityQuota": {
      "fitCount": 1204,
      "limit": 2421,
      "name": "Running On-Demand Standard (A, C, D, H, I, M, R, T, Z) instances",
      "perInstance": 2,
      "scope": "region",
      "unit": "vCPU",
      "used": 12
    },
    "availabilityRefreshAfter": "2026-11-25T09:59:33.21011855Z",
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
- **Duration:** 3.688s
- **Request URL:** `http://localhost:1323/tumblebug/ns/default/resources/vNet`
```json
// Request Body
{
  "cidrBlock": "10.0.0.0/16",
  "connectionName": "aws-ap-northeast-2",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "name": "test-rdbms-vnet-aws",
  "subnetInfoList": [
    {
      "ipv4_CIDR": "10.0.1.0/24",
      "name": "subnet-1",
      "zone": "ap-northeast-2a"
    },
    {
      "ipv4_CIDR": "10.0.2.0/24",
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
  "cidrBlock": "10.0.0.0/16",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T09:54:37Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:37Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-11-25T09:54:37Z",
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
  "cspResourceId": "vpc-014472542cc87bfa5",
  "cspResourceName": "tbt5alk98e6u93kche7b",
  "description": "Pre-requisite VNet for CM-Beetle RDBMS test",
  "id": "test-rdbms-vnet-aws",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "CidrBlock",
      "value": "10.0.0.0/16"
    },
    {
      "key": "CidrBlockAssociationSet",
      "value": "{AssociationId:vpc-cidr-assoc-083cd7050019a1efd,CidrBlock:10.0.0.0/16,CidrBlockState:{State:associated,StatusMessage:null}}"
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
      "value": "{Key:Name,Value:tbt5alk98e6u93kche7b}"
    },
    {
      "key": "VpcId",
      "value": "vpc-014472542cc87bfa5"
    }
  ],
  "name": "test-rdbms-vnet-aws",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:54:37Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:54:37Z",
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
      "cspResourceId": "subnet-0cc10b17b6ec493e1",
      "cspResourceName": "tbigsnr2e3t5uc65i2no",
      "cspVNetId": "vpc-014472542cc87bfa5",
      "cspVNetName": "tbt5alk98e6u93kche7b",
      "description": "",
      "id": "subnet-1",
      "ipv4_CIDR": "10.0.1.0/24",
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
          "value": "10.0.1.0/24"
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
          "value": "arn:aws:ec2:ap-northeast-2:635484366616:subnet/subnet-0cc10b17b6ec493e1"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0cc10b17b6ec493e1"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbigsnr2e3t5uc65i2no}"
        },
        {
          "key": "VpcId",
          "value": "vpc-014472542cc87bfa5"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbigsnr2e3t5uc65i2no",
      "zone": "ap-northeast-2a"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T09:54:37Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T09:54:37Z",
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
      "cspResourceId": "subnet-0d95aacbfb1b0dbeb",
      "cspResourceName": "tbbmdrn1tarngstmna7k",
      "cspVNetId": "vpc-014472542cc87bfa5",
      "cspVNetName": "tbt5alk98e6u93kche7b",
      "description": "",
      "id": "subnet-2",
      "ipv4_CIDR": "10.0.2.0/24",
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
          "value": "10.0.2.0/24"
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
          "value": "arn:aws:ec2:ap-northeast-2:635484366616:subnet/subnet-0d95aacbfb1b0dbeb"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0d95aacbfb1b0dbeb"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbbmdrn1tarngstmna7k}"
        },
        {
          "key": "VpcId",
          "value": "vpc-014472542cc87bfa5"
        }
      ],
      "name": "subnet-2",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbbmdrn1tarngstmna7k",
      "zone": "ap-northeast-2c"
    }
  ],
  "systemLabel": "",
  "uid": "tbt5alk98e6u93kche7b"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 1.978s
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
  "name": "test-rdbms-sg-aws",
  "vNetId": "test-rdbms-vnet-aws"
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
  "cspResourceId": "sg-00fbbe9cd2de98ea2",
  "cspResourceName": "tburflq9g2o29alb00se",
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
  "id": "test-rdbms-sg-aws",
  "isAutoGenerated": false,
  "keyValueList": [
    {
      "key": "GroupName",
      "value": "tburflq9g2o29alb00se"
    },
    {
      "key": "VpcID",
      "value": "vpc-014472542cc87bfa5"
    },
    {
      "key": "OwnerID",
      "value": "635484366616"
    },
    {
      "key": "Description",
      "value": "tburflq9g2o29alb00se"
    }
  ],
  "name": "test-rdbms-sg-aws",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tburflq9g2o29alb00se",
  "vNetId": "test-rdbms-vnet-aws"
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
- **Duration:** 40.245s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=aws-ap-northeast-2&dbEngine=mysql`
```json
// Response Body
{
  "resourceType": "rdbms",
  "supports": {
    "backupRetentionRange": "0-35",
    "connectionName": "aws-ap-northeast-2",
    "dbEngine": "mysql",
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
      "5.7",
      "8.0",
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
  "autoFillSourceDefaults": true,
  "desiredCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for aws (ap-northeast-2)",
  "status": "recommended",
  "targetCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
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
      "dbInstanceSpec": "db.t3.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-aws",
      "securityGroupIds": [
        "test-rdbms-sg-aws"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "gp3",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "vNetId": "test-rdbms-vnet-aws"
    }
  ]
}
```

### 7. Beetle POST Validate RDBMS Recommendation [✅ SUCCESS]
- **Duration:** 6.537s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "admin",
  "adminUserPassword": "******",
  "autoFillDefaults": true,
  "connectionName": "aws-ap-northeast-2",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0",
  "dbInstanceSpec": "db.t3.medium",
  "name": "rdbms-aws",
  "publicAccess": true,
  "securityGroupIds": [
    "test-rdbms-sg-aws"
  ],
  "storageSize": 100,
  "storageType": "gp3",
  "subnetIds": [
    "subnet-1",
    "subnet-2"
  ],
  "vNetId": "test-rdbms-vnet-aws"
}
```
```json
// Response Body
{
  "data": {
    "adminUserName": "admin",
    "adminUserPassword": "******",
    "connectionName": "aws-ap-northeast-2",
    "dbEngine": "mysql",
    "dbEngineVersion": "8.0",
    "dbInstanceSpec": "db.t3.medium",
    "name": "rdbms-aws",
    "publicAccess": true,
    "securityGroupIds": [
      "test-rdbms-sg-aws"
    ],
    "storageSize": 100,
    "storageType": "gp3",
    "subnetIds": [
      "subnet-1",
      "subnet-2"
    ],
    "vNetId": "test-rdbms-vnet-aws"
  },
  "message": "RDBMS configuration is valid",
  "success": true
}
```

### 8. Beetle POST Migrate RDBMS (Provisioning) [✅ SUCCESS]
- **Duration:** 10m3.202s
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
      "dbInstanceSpec": "db.t3.medium",
      "highAvailability": false,
      "publicAccess": true,
      "rdbmsName": "rdbms-aws",
      "securityGroupIds": [
        "test-rdbms-sg-aws"
      ],
      "sourceInstanceName": "Source MySQL 01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "gp3",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "vNetId": "test-rdbms-vnet-aws"
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
- **Duration:** 9ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "16:23-16:53",
  "conditions": [
    {
      "lastTransitionTime": "2026-11-25T10:04:51Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-11-25T10:04:51Z",
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
      "value": "available"
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
      "key": "LatestRestorableTime",
      "value": "2026-11-25T10:00:09.563Z"
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
      "value": "tbvtoteol9rtlis035nc"
    }
  ],
  "uid": "tbvtoteol9rtlis035nc",
  "vNetId": "test-rdbms-vnet-aws"
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 30ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms`
```json
// Response Body
{
  "rdbms": [
    {
      "backupRetentionDays": 7,
      "backupTime": "07:00Z-08:00Z",
      "conditions": [
        {
          "lastTransitionTime": "2026-11-25T10:04:29Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:01:41Z",
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
      "cspResourceId": "rm-mj7uar0nh1r8n558k",
      "cspResourceName": "tbc066ftkd9ijporkvrq",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "mysql.n2.medium.1",
      "dbInstanceType": "Basic",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "43.108.66.129:3306",
      "highAvailability": false,
      "id": "test-rdbms-alibaba",
      "keyValueList": [
        {
          "key": "Engine",
          "value": "MySQL"
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
          "value": "rm-mj7uar0nh1r8n558k-20261125175911"
        },
        {
          "key": "DBInstanceNetType",
          "value": "Intranet"
        },
        {
          "key": "DBInstanceClass",
          "value": "mysql.n2.medium.1"
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
          "value": "rm-mj7uar0nh1r8n558k"
        },
        {
          "key": "VpcId",
          "value": "vpc-mj7t8uchc521yb50jec8v"
        },
        {
          "key": "DBInstanceMemory",
          "value": "4096"
        },
        {
          "key": "DBInstanceStorageType",
          "value": "general_essd"
        },
        {
          "key": "LatestKernelVersion",
          "value": "rds_20260620"
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
          "value": "8.0"
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
          "value": "rds_20260620"
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
          "key": "LockMode",
          "value": "Unlock"
        },
        {
          "key": "CanTempUpgrade",
          "value": "false"
        },
        {
          "key": "Category",
          "value": "Basic"
        },
        {
          "key": "InsId",
          "value": "1"
        },
        {
          "key": "DBInstanceDescription",
          "value": "tbc066ftkd9ijporkvrq"
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
          "value": "2"
        },
        {
          "key": "ResourceGroupId",
          "value": "rg-acfnvekhilw5kmy"
        },
        {
          "key": "DispenseMode",
          "value": "ClassicDispenseMode"
        },
        {
          "key": "TempUpgradeRecoveryMemory",
          "value": "0"
        },
        {
          "key": "ProxyType",
          "value": "2"
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
          "value": "vsw-mj7yfvd5kk99kumix6i1u"
        },
        {
          "key": "CreationTime",
          "value": "2026-11-25T09:59:16Z"
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
          "value": "rm-mj7uar0nh1r8n558k.mysql.ap-northeast-2.rds.aliyuncs.com"
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
          "key": "AutoUpgradeMinorVersion",
          "value": "Auto"
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
          "value": "{SlaveZone:[]}"
        }
      ],
      "name": "test-rdbms-alibaba",
      "publicAccess": true,
      "resourceType": "rdbms",
      "securityGroupIds": [
        "test-rdbms-sg-alibaba"
      ],
      "status": "Deleting",
      "storageSize": 100,
      "storageType": "general_essd",
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
          "lastTransitionTime": "2026-11-25T10:04:51Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-11-25T10:04:51Z",
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
          "value": "available"
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
          "key": "LatestRestorableTime",
          "value": "2026-11-25T10:00:09.563Z"
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
          "lastTransitionTime": "2026-11-25T10:02:58Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
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
      "status": "Deleting",
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
          "lastTransitionTime": "2026-11-25T10:03:23Z",
          "message": "RDBMS deletion in progress",
          "reason": "Deleting",
          "status": "False",
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
      "status": "Deleting",
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
- **Duration:** 769ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws/database`
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
- **Duration:** 1.404s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws/database`
```json
// Response Body
{
  "databases": [
    "information_schema",
    "mysql",
    "performance_schema",
    "sampledb",
    "sampledb_dyn",
    "sys"
  ]
}
```

### 13. Beetle GET Secure Transport Info [✅ SUCCESS]
- **Duration:** 19.443s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": false,
    "issuer": "",
    "notAfter": "0001-01-01T00:00:00Z",
    "pem": "",
    "subject": ""
  },
  "caCertificateError": "after 2 attempts, last error: dial: dial tcp 43.200.17.225:3306: i/o timeout",
  "enforced": false,
  "engine": "mysql",
  "recommendedSSLMode": "",
  "requireSecureTransport": "OFF",
  "rules": "",
  "tlsCipher": "TLS_AES_128_GCM_SHA256",
  "tlsInUse": true
}
```

### 14. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 47.438s
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m37.273s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 13.22s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m21.401s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 825ms

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 10.481s

