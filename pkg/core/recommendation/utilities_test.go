package recommendation

import (
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/stretchr/testify/assert"
)

func TestIsDuplicateInfraCandidate(t *testing.T) {
	c1 := cloudmodel.RecommendedInfra{
		TargetInfra: cloudmodel.InfraReq{
			NodeGroups: []cloudmodel.CreateNodeGroupReq{
				{SpecId: "m2-g3", ImageId: "img-01"},
				{SpecId: "s2-g3", ImageId: "img-01"},
			},
		},
	}
	c2 := cloudmodel.RecommendedInfra{
		TargetInfra: cloudmodel.InfraReq{
			NodeGroups: []cloudmodel.CreateNodeGroupReq{
				{SpecId: "m2-g3a", ImageId: "img-01"},
				{SpecId: "s2-g3", ImageId: "img-01"},
			},
		},
	}
	cDuplicateOfC1 := cloudmodel.RecommendedInfra{
		TargetInfra: cloudmodel.InfraReq{
			NodeGroups: []cloudmodel.CreateNodeGroupReq{
				{SpecId: "m2-g3", ImageId: "img-01"},
				{SpecId: "s2-g3", ImageId: "img-01"},
			},
		},
	}

	existing := []cloudmodel.RecommendedInfra{c1, c2}

	// c1 duplicate should be detected
	assert.True(t, isDuplicateInfraCandidate(existing, cDuplicateOfC1))

	// c3 with different spec should not be detected as duplicate
	c3 := cloudmodel.RecommendedInfra{
		TargetInfra: cloudmodel.InfraReq{
			NodeGroups: []cloudmodel.CreateNodeGroupReq{
				{SpecId: "m2-g3", ImageId: "img-01"},
				{SpecId: "c2-g3", ImageId: "img-01"},
			},
		},
	}
	assert.False(t, isDuplicateInfraCandidate(existing, c3))
}
