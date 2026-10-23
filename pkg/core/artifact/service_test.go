package artifact

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremisemodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArtifact_ThreeStageLineageAndGenerics(t *testing.T) {
	tmpDir := t.TempDir()
	st, err := InitStore(tmpDir)
	require.NoError(t, err)
	require.NotNil(t, st)

	// 1. Create Source Infra Artifact using CreateAs[OnpremInfra]
	srcInfra := onpremisemodel.OnpremInfra{
		Nodes: []onpremisemodel.NodeProperty{
			{
				Hostname:  "centos-web-01",
				MachineId: "m-web-01",
				CPU: onpremisemodel.CpuProperty{
					Architecture: "x86_64",
					Cpus:         1,
					Cores:        4,
					Threads:      8,
				},
				Memory: onpremisemodel.MemoryProperty{
					TotalSize: 16,
				},
			},
		},
	}

	srcArt, err := CreateAs(CreateMetadata{
		Id:            "art-src-101",
		Name:          "ecommerce-source-v1",
		Description:   "Initial source snapshot",
		Category:      CategoryInfra,
		Stage:         StageSource,
		SourceGroupId: "sg-ecommerce-prod",
	}, srcInfra)
	require.NoError(t, err)
	assert.Equal(t, "OnpremInfra", srcArt.SchemaType)
	assert.Equal(t, "sg-ecommerce-prod", srcArt.Lineage.SourceGroupId)
	assert.Equal(t, "art-src-101", srcArt.Lineage.SourceModelId)
	assert.Empty(t, srcArt.Lineage.PlanningModelId)

	// 2. Derive Planning Infra Artifact from Source via ParentId
	planInfra := srcInfra
	planInfra.Nodes[0].Hostname = "ubuntu-web-01"
	planArt, err := CreateAs(CreateMetadata{
		Id:          "art-plan-102",
		Name:        "ecommerce-plan-ubuntu",
		Description: "Refined OS to Ubuntu 22.04",
		Category:    CategoryInfra,
		Stage:       StagePlanning,
		ParentId:    srcArt.Id,
	}, planInfra)
	require.NoError(t, err)
	assert.Equal(t, "sg-ecommerce-prod", planArt.Lineage.SourceGroupId)
	assert.Equal(t, "art-src-101", planArt.Lineage.SourceModelId)
	assert.Equal(t, "art-plan-102", planArt.Lineage.PlanningModelId)

	// 3. Derive Target Infra Artifact using legacy JSON payload to test Upcast + Lineage together
	legacyTargetPayload := json.RawMessage(`{
		"status": "recommended",
		"description": "AWS target infra",
		"targetCloud": {"csp": "aws", "region": "ap-northeast-2"},
		"targetVmInfra": {
			"name": "aws-ecommerce-infra",
			"subGroups": [
				{
					"name": "web-ng",
					"subGroupSize": 2,
					"specId": "aws+ap-northeast-2+t3.large",
					"imageId": "ami-ubuntu-2204",
					"imageType": "PublicImage",
					"vmUserName": "ubuntu"
				}
			]
		}
	}`)

	rawTarget, err := Create(CreateReq{
		Id:       "art-tgt-103",
		Name:     "ecommerce-aws-target",
		Category: CategoryInfra,
		Stage:    StageTarget,
		ParentId: planArt.Id,
		Payload:  legacyTargetPayload,
	})
	require.NoError(t, err)
	assert.Equal(t, "RecommendedInfra", rawTarget.SchemaType)
	assert.Equal(t, "sg-ecommerce-prod", rawTarget.Lineage.SourceGroupId)
	assert.Equal(t, "art-src-101", rawTarget.Lineage.SourceModelId)
	assert.Equal(t, "art-plan-102", rawTarget.Lineage.PlanningModelId)
	assert.Equal(t, "art-tgt-103", rawTarget.Lineage.TargetModelId)
	require.NotNil(t, rawTarget.ConversionReport)
	assert.True(t, rawTarget.ConversionReport.Upcasted)

	// 4. Cast and retrieve using GetAs[cloudmodel.RecommendedInfra]
	typedTarget, err := GetAs[cloudmodel.RecommendedInfra]("art-tgt-103")
	require.NoError(t, err)
	assert.Equal(t, "aws-ecommerce-infra", typedTarget.Payload.TargetInfra.Name)
	require.Len(t, typedTarget.Payload.TargetInfra.NodeGroups, 1)
	assert.Equal(t, 2, typedTarget.Payload.TargetInfra.NodeGroups[0].NodeGroupSize)
	assert.Equal(t, "ubuntu", typedTarget.Payload.TargetInfra.NodeGroups[0].NodeUserName)

	// 5. Verify Type Guard blocks casting RecommendedInfra to SourceObjectStorage
	_, err = GetAs[storagemodel.SourceObjectStorage]("art-tgt-103")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "schemaType mismatch")

	// 6. Verify LineageGraph resolution
	graph, err := GetLineageGraph("art-tgt-103")
	require.NoError(t, err)
	require.NotNil(t, graph.SourceArtifact)
	require.NotNil(t, graph.PlanningArtifact)
	require.NotNil(t, graph.TargetArtifact)
	assert.Equal(t, "art-src-101", graph.SourceArtifact.Id)
	assert.Equal(t, "art-plan-102", graph.PlanningArtifact.Id)
	assert.Equal(t, "art-tgt-103", graph.TargetArtifact.Id)
	assert.Len(t, graph.GroupArtifacts, 3)

	// 7. Verify PUT update and atomic disk persistence reload
	updatedPlan, err := Update("art-plan-102", UpdateReq{
		Name:    "ecommerce-plan-ubuntu-v2",
		Version: "2.0.0",
	})
	require.NoError(t, err)
	assert.Equal(t, "ecommerce-plan-ubuntu-v2", updatedPlan.Name)
	assert.Equal(t, "2.0.0", updatedPlan.Version)

	reloadedStore, err := NewStore(tmpDir)
	require.NoError(t, err)
	reloadedItem, err := reloadedStore.Get("art-plan-102")
	require.NoError(t, err)
	assert.Equal(t, "ecommerce-plan-ubuntu-v2", reloadedItem.Name)

	// 8. Verify path traversal protection
	_, err = reloadedStore.resolveSafeFilePath("../escape")
	require.Error(t, err)

	// 9. Verify file exists on disk and delete removes it
	jsonPath := filepath.Join(tmpDir, "art-tgt-103.json")
	_, err = os.Stat(jsonPath)
	require.NoError(t, err)
	require.NoError(t, Delete("art-tgt-103"))
	_, err = os.Stat(jsonPath)
	assert.True(t, os.IsNotExist(err))
}

