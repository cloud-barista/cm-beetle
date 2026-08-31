# Managed RDBMS Test Report: AWS (ap-northeast-2)

- **Test Case:** AWS AP-Northeast-2 (Seoul) MySQL Test
- **Date & Time:** 2026-08-31 16:38:03
- **Namespace:** `default`
- **Total Duration:** 18m17.787s
- **Overall Status:** ✅ PASSED

## Environment and Scenario

### Environment
- **Target CSP:** AWS
- **Target Region:** `ap-northeast-2`
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
- **Duration:** 3.99s
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
    "queriedAt": "2026-08-31T07:38:07.138843007Z",
    "region": "ap-northeast-2",
    "source": "aws:DescribeInstanceTypeOfferings",
    "zones": [
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2c"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2b"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2d"
      },
      {
        "available": true,
        "status": "AVAILABLE",
        "zoneId": "ap-northeast-2a"
      }
    ]
  },
  "connectionName": "aws-ap-northeast-2",
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
  "isValid": true,
  "message": "Spec and image pair is valid for provisioning",
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
  "status": "OK",
  "suggestedZone": "ap-northeast-2c"
}
```

### 2. Tumblebug POST /resources/vNet (Create VNet & Subnets) [✅ SUCCESS]
- **Duration:** 3.72s
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
      "lastTransitionTime": "2026-08-31T07:38:10Z",
      "reason": "Available",
      "status": "True",
      "type": "Ready"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:10Z",
      "reason": "Available",
      "status": "True",
      "type": "Synced"
    },
    {
      "lastTransitionTime": "2026-08-31T07:38:10Z",
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
  "cspResourceId": "vpc-0a95dce99183abd55",
  "cspResourceName": "tb4cmrvs9ddgv3to0n2n",
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
      "value": "{AssociationId:vpc-cidr-assoc-0d55e131101a12a43,CidrBlock:10.0.0.0/16,CidrBlockState:{State:associated,StatusMessage:null}}"
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
      "value": "{Key:Name,Value:tb4cmrvs9ddgv3to0n2n}"
    },
    {
      "key": "VpcId",
      "value": "vpc-0a95dce99183abd55"
    }
  ],
  "name": "test-rdbms-vnet-aws",
  "resourceType": "vNet",
  "status": "Available",
  "subnetInfoList": [
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-08-31T07:38:10Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-08-31T07:38:10Z",
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
      "cspResourceId": "subnet-032cd8a3ea137052e",
      "cspResourceName": "tbsm2va1f9cmm1gbuo38",
      "cspVNetId": "vpc-0a95dce99183abd55",
      "cspVNetName": "tb4cmrvs9ddgv3to0n2n",
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
          "value": "arn:aws:ec2:ap-northeast-2:635484366616:subnet/subnet-032cd8a3ea137052e"
        },
        {
          "key": "SubnetId",
          "value": "subnet-032cd8a3ea137052e"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tbsm2va1f9cmm1gbuo38}"
        },
        {
          "key": "VpcId",
          "value": "vpc-0a95dce99183abd55"
        }
      ],
      "name": "subnet-1",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tbsm2va1f9cmm1gbuo38",
      "zone": "ap-northeast-2a"
    },
    {
      "conditions": [
        {
          "lastTransitionTime": "2026-08-31T07:38:10Z",
          "reason": "Available",
          "status": "True",
          "type": "Ready"
        },
        {
          "lastTransitionTime": "2026-08-31T07:38:10Z",
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
      "cspResourceId": "subnet-0167ed5a3499addab",
      "cspResourceName": "tb4cq1tf2ohiqg2t32cu",
      "cspVNetId": "vpc-0a95dce99183abd55",
      "cspVNetName": "tb4cmrvs9ddgv3to0n2n",
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
          "value": "arn:aws:ec2:ap-northeast-2:635484366616:subnet/subnet-0167ed5a3499addab"
        },
        {
          "key": "SubnetId",
          "value": "subnet-0167ed5a3499addab"
        },
        {
          "key": "Tags",
          "value": "{Key:Name,Value:tb4cq1tf2ohiqg2t32cu}"
        },
        {
          "key": "VpcId",
          "value": "vpc-0a95dce99183abd55"
        }
      ],
      "name": "subnet-2",
      "resourceType": "subnet",
      "status": "Available",
      "uid": "tb4cq1tf2ohiqg2t32cu",
      "zone": "ap-northeast-2c"
    }
  ],
  "systemLabel": "",
  "uid": "tb4cmrvs9ddgv3to0n2n"
}
```

### 3. Tumblebug POST /resources/securityGroup (Create SecurityGroup) [✅ SUCCESS]
- **Duration:** 1.435s
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
  "cspResourceId": "sg-0446cc88f3390b509",
  "cspResourceName": "tb109hm0lpeb112q97qv",
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
      "value": "tb109hm0lpeb112q97qv"
    },
    {
      "key": "VpcID",
      "value": "vpc-0a95dce99183abd55"
    },
    {
      "key": "OwnerID",
      "value": "635484366616"
    },
    {
      "key": "Description",
      "value": "tb109hm0lpeb112q97qv"
    }
  ],
  "name": "test-rdbms-sg-aws",
  "resourceType": "securityGroup",
  "systemLabel": "",
  "uid": "tb109hm0lpeb112q97qv",
  "vNetId": "test-rdbms-vnet-aws"
}
```

### 4. Beetle GET RDBMS Support [✅ SUCCESS]
- **Duration:** 23ms
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
- **Duration:** 31.42s
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms/capability?connectionName=aws-ap-northeast-2`
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
          "maxSize": 65536,
          "minSize": 20,
          "recommendationLevel": "legacy",
          "storageType": "gp2"
        },
        {
          "constraints": "Minimum 20GB storage.",
          "description": "Latest generation general-purpose SSD with 3,000 baseline IOPS and 125 MiB/s throughput. Recommended for most workloads.",
          "displayName": "General Purpose SSD (gp3)",
          "maxSize": 65536,
          "minSize": 20,
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
          "maxSize": 65536,
          "minSize": 100,
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
          "maxSize": 65536,
          "minSize": 100,
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
    "storageSizeRange": {
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
- **Duration:** 5ms
- **Request URL:** `http://localhost:8056/beetle/recommendation/middleware/rdbms`
```json
// Request Body
{
  "desiredCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
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
  "description": "Successfully recommended 1 managed RDBMS configuration(s) for aws (ap-northeast-2)",
  "warnings": [
    "Requested storage type 'SSD' is not supported on target cloud. Replaced with capability recommended storage for instance 'source-mysql-01'."
  ],
  "targetCloud": {
    "csp": "aws",
    "region": "ap-northeast-2"
  },
  "targetRDBMSInstances": [
    {
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "rdbmsName": "test-rdbms-aws",
      "dbEngine": "mysql",
      "dbEngineVersion": "8.0",
      "dbInstanceSpec": "db.t3.medium",
      "storageType": "gp3",
      "storageSize": 100,
      "adminUserName": "root",
      "adminUserPassword": "******",
      "vNetId": "test-rdbms-vnet-aws",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "securityGroupIds": [
        "test-rdbms-sg-aws"
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
- **Duration:** 3.418s
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
  "name": "test-rdbms-aws",
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
    "name": "test-rdbms-aws",
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
- **Duration:** 11m3.502s
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
      "rdbmsName": "test-rdbms-aws",
      "securityGroupIds": [
        "test-rdbms-sg-aws"
      ],
      "sourceInstanceName": "source-mysql-01",
      "sourceMachineId": "node-550e8400-e29b-41d4-a716-446655440000",
      "storageSize": 100,
      "storageType": "gp3",
      "subnetIds": [
        "subnet-1",
        "subnet-2"
      ],
      "vNetId": "test-rdbms-vnet-aws"
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
- **Duration:** 6ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-aws`
```json
// Response Body
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
      "key": "sys.connectionName",
      "value": "aws-ap-northeast-2"
    },
    {
      "key": "sys.manager",
      "value": "cb-tumblebug"
    },
    {
      "key": "sys.labelType",
      "value": "rdbms"
    },
    {
      "key": "sys.cspResourceId",
      "value": "tb1k45ucvlpo2hrndcpf"
    },
    {
      "key": "sys.id",
      "value": "test-test-rdbms-aws"
    },
    {
      "key": "sys.cspResourceName",
      "value": "tb1k45ucvlpo2hrndcpf"
    },
    {
      "key": "sys.description",
      "value": "Migrated by CM-Beetle from source instance source-mysql-01"
    },
    {
      "key": "sys.uid",
      "value": "tb1k45ucvlpo2hrndcpf"
    },
    {
      "key": "sys.name",
      "value": "test-test-rdbms-aws"
    },
    {
      "key": "Name",
      "value": "tb1k45ucvlpo2hrndcpf"
    },
    {
      "key": "sys.namespace",
      "value": "default"
    }
  ]
}
```

### 10. Beetle GET RDBMS List [✅ SUCCESS]
- **Duration:** 8ms
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
      "status": "Available",
      "conditions": [
        {
          "type": "Ready",
          "status": "True",
          "reason": "Available",
          "lastTransitionTime": "2026-08-31T07:46:10Z"
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
      "status": "Deleting",
      "conditions": [
        {
          "type": "Ready",
          "status": "False",
          "reason": "Deleting",
          "message": "RDBMS deletion in progress",
          "lastTransitionTime": "2026-08-31T07:47:26Z"
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
- **Duration:** 585ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-aws/database`
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
- **Duration:** 577ms
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-aws/database`
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

### 13. Data I/O Test (External Remote) [✅ SUCCESS]
- **Duration:** 90ms
```json
// Response Body
{
  "result": "External SQL write/read/verify/drop cycle succeeded"
}
```

### 14. Data I/O Test (Internal VPC VM) [✅ SUCCESS]
- **Duration:** 3m44.617s
```json
// Response Body
{
  "result": "Pass"
}
```

### 15. Beetle DELETE Logical Database [✅ SUCCESS]
- **Duration:** 13.656s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-aws/database/sampledb`

### 16. Beetle DELETE RDBMS Instance [✅ SUCCESS]
- **Duration:** 2m21.245s
- **Request URL:** `http://localhost:8056/beetle/migration/middleware/ns/default/rdbms/test-test-rdbms-aws?option=force`

### 17. Tumblebug DELETE /resources/securityGroup [✅ SUCCESS]
- **Duration:** 784ms

### 18. Tumblebug DELETE /resources/vNet [✅ SUCCESS]
- **Duration:** 8.707s

