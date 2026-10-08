# Managed RDBMS (MySQL) Test Report: AWS (ap-northeast-2)

- **Test Case:** AWS AP-Northeast-2 (Seoul) MySQL Test
- **Date & Time:** 2026-10-08 18:20:26
- **Namespace:** `default`
- **Total Duration:** 15m33s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** AWS
- **Target Region:** `ap-northeast-2`
- **Namespace:** `default`
- **Test Date:** 2026-10-08 18:20:26

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
- **Duration:** 2.537s
- **Request URL:** `http://localhost:1323/tumblebug/specImagePairReview`
```json
// Request Body
{
  "imageId": "ami-04e3ca2324a305ad0",
  "specId": "aws+ap-northeast-2+t4g.medium"
}
```
```json
// Response Body
{
  "availability": {
    "available": true,
    "instanceType": "t4g.medium",
    "provider": "aws",
    "queriedAt": "2026-10-08T09:20:28.823687714Z",
    "region": "ap-northeast-2",
    "source": "aws:DescribeInstanceTypeOfferings",
    "zones": [
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2b"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2a"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2d"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2c"
      }
    ]
  },
  "connectionName": "aws-ap-northeast-2",
  "errors": [
    "High provisioning failure risk: Spec 'aws+ap-northeast-2+t4g.medium' has 100% failure rate (1 failures out of 1 attempts) | Image 'ami-04e3ca2324a305ad0' has FAILED with spec 'aws+ap-northeast-2+t4g.medium' before (never succeeded). Recent failures: VM provisioning failed (status: Terminating) - check CSP console for details"
  ],
  "estimatedCost": "$0.0416/hour",
  "imageDetails": {
    "commandHistory": null,
    "connectionName": "aws-ap-northeast-2",
    "creationDate": "2026-07-14T11:54:44.000Z",
    "cspImageName": "ami-04e3ca2324a305ad0",
    "description": "Canonical, Ubuntu, 24.04, arm64 noble image",
    "details": [
      {
        "key": "Architecture",
        "value": "arm64"
      },
      {
        "key": "BlockDeviceMappings",
        "value": "{DeviceName:/dev/sda1,Ebs:{DeleteOnTermination:true,Encrypted:false,Iops:null,KmsKeyId:null,OutpostArn:null,SnapshotId:snap-0f98c4fc07815c093,Throughput:null,VolumeSize:8,VolumeType:gp3},NoDevice:null,VirtualName:null}; {DeviceName:/dev/sdb,Ebs:null,NoDevice:null,VirtualName:ephemeral0}; {DeviceName:/dev/sdc,Ebs:null,NoDevice:null,VirtualName:ephemeral1}"
      },
      {
        "key": "BootMode",
        "value": "uefi"
      },
      {
        "key": "CreationDate",
        "value": "2026-07-14T11:54:44.000Z"
      },
      {
        "key": "DeprecationTime",
        "value": "2028-07-14T11:54:44.000Z"
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
        "value": "ami-04e3ca2324a305ad0"
      },
      {
        "key": "ImageLocation",
        "value": "amazon/ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-arm64-server-20260714"
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
        "value": "ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-arm64-server-20260714"
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
    "fetchedTime": "2026.08.21 14:20:39 Fri",
    "id": "ami-04e3ca2324a305ad0",
    "imageStatus": "Available",
    "infraType": "",
    "isBasicGpuImage": false,
    "isBasicImage": true,
    "isGPUImage": false,
    "isKubernetesImage": false,
    "name": "ami-04e3ca2324a305ad0",
    "namespace": "system",
    "osArchitecture": "arm64",
    "osDiskSizeGB": -1,
    "osDiskType": "ebs",
    "osDistribution": "ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-arm64-server-20260714",
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
    "uid": "tbk8m6abcnq0ogukdn90"
  },
  "imageId": "ami-04e3ca2324a305ad0",
  "imageValidation": {
    "cspResourceId": "ami-04e3ca2324a305ad0",
    "isAvailable": true,
    "resourceId": "ami-04e3ca2324a305ad0",
    "resourceName": "ami-04e3ca2324a305ad0",
    "status": "Available"
  },
  "isValid": false,
  "message": "High provisioning failure risk detected",
  "providerName": "aws",
  "regionName": "ap-northeast-2",
  "specDetails": {
    "architecture": "arm64",
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
        "value": "{EbsOptimizedInfo:{BaselineBandwidthInMbps:347,BaselineIops:2000,BaselineThroughputInMBps:43.375,MaximumBandwidthInMbps:2085,MaximumIops:11800,MaximumThroughputInMBps:260.625},EbsOptimizedSupport:default,EncryptionSupport:supported,NvmeSupport:required}"
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
        "value": "{DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:required,Ipv4AddressesPerInterface:6,Ipv6AddressesPerInterface:6,Ipv6Supported:true,MaximumNetworkCards:1,MaximumNetworkInterfaces:3,NetworkCards:[{MaximumNetworkInterfaces:3,NetworkCardIndex:0,NetworkPerformance:Up to 5 Gigabit}],NetworkPerformance:Up to 5 Gigabit}"
      },
      {
        "key": "PlacementGroupInfo",
        "value": "{SupportedStrategies:[partition,spread]}"
      },
      {
        "key": "ProcessorInfo",
        "value": "{SupportedArchitectures:[arm64],SustainedClockSpeedInGhz:2.5}"
      },
      {
        "key": "SupportedBootModes",
        "value": "uefi"
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
    "uid": "tbk5s7k571ba0vqjtqkg",
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
  "status": "Error",
  "suggestedZone": "ap-northeast-2b"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 2.588s
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
      "lastTransitionTime": "2026-10-08T09:20:31Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-08T09:20:31Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-10-08T09:20:31Z",
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
  "cspResourceId": "vpc-0b899cd218f25b23f",
  "cspResourceName": "tbp57geoj1pbnvld0pea",
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
      "value": "{AssociationId:vpc-cidr-assoc-0405f77106431455f,CidrBlock:10.0.0.0/16,CidrBlockState:{State:associated,StatusMessage:null}}"
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
      "value": "{Key:Name,Value:tbp57geoj1pbnvld0pea}"
    },
    {
      "key": "VpcId",
      "value": "vpc-0b899cd218f25b23f"
    }
  ],
  "name": "test-rdbms-vnet-aws",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T09:20:31Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T09:20:31Z",
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
      "cspResourceId": "subnet-0c1c9eddc92f9e209",
      "cspResourceName": "tbha74rblmpu2pjltikf",
      "cspVNetId": "vpc-0b899cd218f25b23f",
      "cspVNetName": "tbp57geoj1pbnvld0pea",
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
          "value": "arn:aws:ec2:ap-northeast-2:635484366616:subnet/subnet-0c1c9eddc92f9e209"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0c1c9eddc92f9e209"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbha74rblmpu2pjltikf}"
        },
        {
          "key": "VpcId",
          "value": "vpc-0b899cd218f25b23f"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbha74rblmpu2pjltikf",
      "zone": "ap-northeast-2a"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T09:20:31Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T09:20:31Z",
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
      "cspResourceId": "subnet-030e0d8fcc05b4650",
      "cspResourceName": "tbh0a8kkp2usdl3k0pcu",
      "cspVNetId": "vpc-0b899cd218f25b23f",
      "cspVNetName": "tbp57geoj1pbnvld0pea",
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
          "value": "arn:aws:ec2:ap-northeast-2:635484366616:subnet/subnet-030e0d8fcc05b4650"
        },
        {
          "key": "SubnetId",
          "value": "subnet-030e0d8fcc05b4650"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbh0a8kkp2usdl3k0pcu}"
        },
        {
          "key": "VpcId",
          "value": "vpc-0b899cd218f25b23f"
        }
      ],
      "name": "subnet-2",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbh0a8kkp2usdl3k0pcu",
      "zone": "ap-northeast-2c"
    }
  ],
  "systemLabel": "",
  "uid": "tbp57geoj1pbnvld0pea"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 1.304s
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
  "cspResourceId": "sg-08ac07a395b3f3307",
  "cspResourceName": "tbn3ifc5uuuiaq5tg2r7",
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
      "value": "tbn3ifc5uuuiaq5tg2r7"
    },
    {
      "key": "VpcID",
      "value": "vpc-0b899cd218f25b23f"
    },
    {
      "key": "OwnerID",
      "value": "635484366616"
    },
    {
      "key": "Description",
      "value": "tbn3ifc5uuuiaq5tg2r7"
    }
  ],
  "name": "test-rdbms-sg-aws",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tbn3ifc5uuuiaq5tg2r7",
  "vNetId": "test-rdbms-vnet-aws"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 11ms
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
- **Duration:** 35.66s
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
      "db.t4g.micro",
      "db.t4g.small"
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
      "8.4"
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
- **Duration:** 3.522s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/validate?nsId=default`
```json
// Request Body
{
  "adminUserName": "root",
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
    "adminUserName": "root",
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
- **Duration:** 7m53.658s
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
- **Duration:** 7ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws`
```json
// Response Body
{
  "backupRetentionDays": 7,
  "backupTime": "16:40-17:10",
  "conditions": [
    {
      "lastTransitionTime": "2026-10-08T09:29:05Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-10-08T09:29:05Z",
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
  "cspResourceId": "tb0asqfaousfd3ga29bb",
  "cspResourceName": "tb0asqfaousfd3ga29bb",
  "dbEngine": "mysql",
  "dbEngineVersion": "8.0.46",
  "dbInstanceSpec": "db.t3.medium",
  "dbInstanceType": "Primary",
  "deletionProtection": false,
  "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
  "endpoint": "tb0asqfaousfd3ga29bb.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com:3306",
  "highAvailability": false,
  "id": "test-rdbms-aws",
  "iops": "3000",
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
      "value": "tb0asqfaousfd3ga29bb"
    }
  ],
  "uid": "tb0asqfaousfd3ga29bb",
  "vNetId": "test-rdbms-vnet-aws"
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
      "backupRetentionDays": 7,
      "backupTime": "16:40-17:10",
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T09:29:05Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T09:29:05Z",
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
      "cspResourceId": "tb0asqfaousfd3ga29bb",
      "cspResourceName": "tb0asqfaousfd3ga29bb",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0.46",
      "dbInstanceSpec": "db.t3.medium",
      "dbInstanceType": "Primary",
      "deletionProtection": false,
      "description": "Migrated by CM-Beetle from source instance Source MySQL 01",
      "endpoint": "tb0asqfaousfd3ga29bb.chrkjg2ktom1.ap-northeast-2.rds.amazonaws.com:3306",
      "highAvailability": false,
      "id": "test-rdbms-aws",
      "iops": "3000",
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
          "value": "tb0asqfaousfd3ga29bb"
        }
      ],
      "uid": "tb0asqfaousfd3ga29bb",
      "vNetId": "test-rdbms-vnet-aws"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T09:20:51Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T09:20:51Z",
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
      "uid": "tbqddlsqavt5vbr0i279",
      "vNetId": "test-rdbms-vnet-ibm"
    },
    {
      "backupRetentionDays": 7,
      "conditions": [
        {
          "lastTransitionTime": "2026-10-08T09:21:06Z",
          "message": "RDBMS creation in progress",
          "reason": "Creating",
          "status": "False",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-10-08T09:21:06Z",
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
      "uid": "tbbb699o1omdfd3jvvpt",
      "vNetId": "test-rdbms-vnet-nhn"
    }
  ]
}
```

### 11. Beetle POST Create Logical Database [✅ SUCCESS]
- **Duration:** 528ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws/database`
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
- **Duration:** 421ms
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
- **Duration:** 470ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws/secure-transport`
```json
// Response Body
{
  "caCertificate": {
    "isSelfSigned": true,
    "issuer": "CN=Amazon RDS ap-northeast-2 Root CA RSA2048 G1,OU=Amazon RDS,O=Amazon Web Services\\, Inc.,L=Seattle,ST=WA,C=US",
    "notAfter": "2061-05-20T17:28:41Z",
    "pem": "-----BEGIN CERTIFICATE-----\nMIIECTCCAvGgAwIBAgIRAPQAvihfjBg/JDbj6U64K98wDQYJKoZIhvcNAQELBQAw\ngZwxCzAJBgNVBAYTAlVTMSIwIAYDVQQKDBlBbWF6b24gV2ViIFNlcnZpY2VzLCBJ\nbmMuMRMwEQYDVQQLDApBbWF6b24gUkRTMQswCQYDVQQIDAJXQTE1MDMGA1UEAwws\nQW1hem9uIFJEUyBhcC1ub3J0aGVhc3QtMiBSb290IENBIFJTQTIwNDggRzExEDAO\nBgNVBAcMB1NlYXR0bGUwIBcNMjEwNTIwMTYyODQxWhgPMjA2MTA1MjAxNzI4NDFa\nMIGcMQswCQYDVQQGEwJVUzEiMCAGA1UECgwZQW1hem9uIFdlYiBTZXJ2aWNlcywg\nSW5jLjETMBEGA1UECwwKQW1hem9uIFJEUzELMAkGA1UECAwCV0ExNTAzBgNVBAMM\nLEFtYXpvbiBSRFMgYXAtbm9ydGhlYXN0LTIgUm9vdCBDQSBSU0EyMDQ4IEcxMRAw\nDgYDVQQHDAdTZWF0dGxlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA\nvJ9lgyksCxkBlY40qOzI1TCj/Q0FVGuPL/Z1Mw2YN0l+41BDv0FHApjTUkIKOeIP\nnwDwpXTa3NjYbk3cOZ/fpH2rYJ++Fte6PNDGPgKppVCUh6x3jiVZ1L7wOgnTdK1Q\nTrw8440IDS5eLykRHvz8OmwvYDl0iIrt832V0QyOlHTGt6ZJ/aTQKl12Fy3QBLv7\nstClPzvHTrgWqVU6uidSYoDtzHbU7Vda7YH0wD9IUoMBf7Tu0rqcE4uH47s2XYkc\nSdLEoOg/Ngs7Y9B1y1GCyj3Ux7hnyvCoRTw014QyNB7dTatFMDvYlrRDGG14KeiU\nUL7Vo/+EejWI31eXNLw84wIDAQABo0IwQDAPBgNVHRMBAf8EBTADAQH/MB0GA1Ud\nDgQWBBQkgTWFsNg6wA3HbbihDQ4vpt1E2zAOBgNVHQ8BAf8EBAMCAYYwDQYJKoZI\nhvcNAQELBQADggEBAGz1Asiw7hn5WYUj8RpOCzpE0h/oBZcnxP8wulzZ5Xd0YxWO\n0jYUcUk3tTQy1QvoY+Q5aCjg6vFv+oFBAxkib/SmZzp4xLisZIGlzpJQuAgRkwWA\n6BVMgRS+AaOMQ6wKPgz1x4v6T0cIELZEPq3piGxvvqkcLZKdCaeC3wCS6sxuafzZ\n4qA3zMwWuLOzRftgX2hQto7d/2YkRXga7jSvQl3id/EI+xrYoH6zIWgjdU1AUaNq\nNGT7DIo47vVMfnd9HFZNhREsd4GJE83I+JhTqIxiKPNxrKgESzyADmNPt0gXDnHo\ntbV1pMZz5HpJtjnP/qVZhEK5oB0tqlKPv9yx074=\n-----END CERTIFICATE-----\n",
    "subject": "CN=Amazon RDS ap-northeast-2 Root CA RSA2048 G1,OU=Amazon RDS,O=Amazon Web Services\\, Inc.,L=Seattle,ST=WA,C=US"
  },
  "enforced": false,
  "engine": "mysql",
  "recommendedSSLMode": "VERIFY_IDENTITY",
  "requireSecureTransport": "OFF",
  "rules": "",
  "tlsCipher": "TLS_AES_128_GCM_SHA256",
  "tlsInUse": true
}
```

### 14. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 20.02s
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 15. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m47.138s
```json
// Response Body
{
  "result": "Pass"
}
```

### 16. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 13.053s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws/database/sampledb`

### 17. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m21.317s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-rdbms-aws?option=force`

### 18. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 769ms

### 19. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 9.984s

