package modelconv

import (
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/imdl/lineage"
	onpremisemodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/cloud-barista/cm-beetle/transx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpcastTo_LegacyRecommendedInfra(t *testing.T) {
	legacyJSON := []byte(`{
		"status": "recommended",
		"description": "Legacy cloud recommendation payload",
		"legacyMetaFlag": "should-be-dropped",
		"targetCloud": {
			"csp": "aws",
			"region": "ap-northeast-2"
		},
		"targetVmInfra": {
			"name": "mig-infra-01",
			"label": {
				"sourceMachineId": "machine-01,machine-02",
				"sourceGroupId": "sg-prod-01"
			},
			"subGroups": [
				{
					"name": "web-group",
					"subGroupSize": "3",
					"specId": "aws+ap-northeast-2+t3.medium",
					"imageId": "ami-12345678",
					"imageType": "PublicImage",
					"vmUserName": "ubuntu",
					"rootDiskSize": "50",
					"label": {
						"sourceMachineId": "machine-01"
					}
				}
			]
		},
		"targetVmSpecList": [
			{
				"id": "aws+ap-northeast-2+t3.medium",
				"cspSpecName": "t3.medium",
				"vCPU": 2,
				"memoryGiB": 4
			}
		],
		"targetVmOsImageList": [
			{
				"id": "ami-12345678",
				"cspImageName": "ubuntu-22.04"
			}
		]
	}`)

	upcasted, report, err := UpcastTo[cloudmodel.RecommendedInfra](legacyJSON)
	require.NoError(t, err)
	require.NotNil(t, report)

	assert.True(t, report.Upcasted)
	assert.Equal(t, "RecommendedInfra", report.TargetSchemaType)
	assert.Equal(t, "mig-infra-01", upcasted.TargetInfra.Name)
	assert.Equal(t, "machine-01,machine-02", upcasted.TargetInfra.Label[lineage.LabelSourceMachineIds])
	assert.Equal(t, "sg-prod-01", upcasted.TargetInfra.Label[lineage.LabelSourceGroupId])

	require.Len(t, upcasted.TargetInfra.NodeGroups, 1)
	ng := upcasted.TargetInfra.NodeGroups[0]
	assert.Equal(t, "web-group", ng.Name)
	assert.Equal(t, 3, ng.NodeGroupSize)
	assert.Equal(t, "ubuntu", ng.NodeUserName)
	assert.Equal(t, 50, ng.RootDiskSize)
	assert.Equal(t, "machine-01", ng.Label[lineage.LabelSourceMachineIds])

	require.Len(t, upcasted.TargetSpecList, 1)
	assert.Equal(t, "t3.medium", upcasted.TargetSpecList[0].CspSpecName)

	require.Len(t, upcasted.TargetOsImageList, 1)
	assert.Equal(t, "ubuntu-22.04", upcasted.TargetOsImageList[0].CspImageName)

	assert.Contains(t, report.DroppedFields, "legacyMetaFlag")
	assert.Contains(t, report.DroppedFields, "targetInfra.nodeGroups[0].imageType")
}

func TestUpcastTo_LegacyWrappedOnpremInfra(t *testing.T) {
	legacyJSON := []byte(`{
		"onpremiseInfraModel": {
			"servers": [
				{
					"hostname": "web-01",
					"machineId": "m-001",
					"cpu": {
						"architecture": "x86_64",
						"cpus": 1,
						"cores": 4,
						"threads": 8
					},
					"memory": {
						"totalSize": 16
					}
				}
			],
			"obsoleteField": true
		}
	}`)

	infra, report, err := UpcastTo[onpremisemodel.OnpremInfra](legacyJSON)
	require.NoError(t, err)
	require.NotNil(t, report)

	assert.True(t, report.Upcasted)
	require.Len(t, infra.Nodes, 1)
	assert.Equal(t, "web-01", infra.Nodes[0].Hostname)
	assert.Equal(t, uint32(4), infra.Nodes[0].CPU.Cores)
	assert.Contains(t, report.DroppedFields, "obsoleteField")
}

func TestUpcastTo_StorageAndRDBMSAndDataModels(t *testing.T) {
	storageJSON := []byte(`{
		"sourceObjectStorageModel": {
			"description": "Legacy bucket list",
			"buckets": [
				{
					"bucketName": "media-assets-prod"
				}
			]
		}
	}`)
	storageModel, report, err := UpcastTo[storagemodel.SourceObjectStorage](storageJSON)
	require.NoError(t, err)
	assert.True(t, report.Upcasted)
	require.Len(t, storageModel.SourceObjectStorages, 1)
	assert.Equal(t, "media-assets-prod", storageModel.SourceObjectStorages[0].BucketName)

	rdbmsJSON := []byte(`{
		"description": "Legacy RDBMS list",
		"rdbmsList": [
			{
				"displayName": "orders-db",
				"dbNode": {
					"hostname": "db-01"
				},
				"dbEngine": {
					"engine": "mysql",
					"engineVersion": "8.0"
				}
			}
		]
	}`)
	rdbmsModel, report, err := UpcastTo[rdbmsmodel.SourceRDBMS](rdbmsJSON)
	require.NoError(t, err)
	assert.True(t, report.Upcasted)
	require.Len(t, rdbmsModel.SourceRDBMSInstances, 1)
	assert.Equal(t, "orders-db", rdbmsModel.SourceRDBMSInstances[0].DisplayName)

	dataJSON := []byte(`{
		"strategy": "auto",
		"source": {
			"storageType": "filesystem",
			"path": "/data/src",
			"filesystem": {
				"accessType": "local"
			}
		},
		"destination": {
			"storageType": "filesystem",
			"path": "/data/dst",
			"filesystem": {
				"accessType": "local"
			}
		}
	}`)
	dataModel, report, err := UpcastTo[transx.DataMigrationModel](dataJSON)
	require.NoError(t, err)
	assert.False(t, report.Upcasted)
	assert.Equal(t, "auto", dataModel.Strategy)
}

func TestUpcastTo_FailFastOnInvalidType(t *testing.T) {
	invalidJSON := []byte(`{
		"targetInfra": {
			"name": "bad-infra",
			"nodeGroups": [
				{
					"name": "ng-1",
					"nodeGroupSize": "not-a-number"
				}
			]
		}
	}`)

	_, _, err := UpcastTo[cloudmodel.RecommendedInfra](invalidJSON)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "expected integer")
}
