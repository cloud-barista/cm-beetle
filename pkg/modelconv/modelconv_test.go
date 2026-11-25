package modelconv

import (
	"testing"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/stretchr/testify/require"
)

func TestValidateTypeCompatibilityEmbeddedAndTags(t *testing.T) {
	type innerSrc struct {
		Shared   string `json:"shared"`
		Shadowed string `json:"shadowed"`
	}
	type outerSrc struct {
		innerSrc
		Shadowed   int    `json:"shadowed"`
		Optional   string `json:",omitempty"`
		Ignored    string `json:"-"`
		unexported string
	}
	type flattenedDst struct {
		Shared   string `json:"shared"`
		Shadowed int    `json:"shadowed"`
		Optional string `json:"Optional,omitempty"`
	}
	type missingFieldDst struct {
		Ignored string `json:"ignored"`
	}

	require.NoError(t, ValidateTypeCompatibility[outerSrc, flattenedDst]())
	require.NoError(t, ValidateTypeCompatibility[outerSrc, outerSrc]())
	require.Error(t, ValidateTypeCompatibility[outerSrc, missingFieldDst]())

	// Verify copied Tumblebug models with embedded structs and json:"-" fields validate cleanly.
	require.NoError(t, ValidateTypeCompatibility[cloudmodel.InfraReq, tbmodel.InfraReq]())
	require.NoError(t, ValidateTypeCompatibility[tbmodel.InfraInfo, cloudmodel.VmInfraInfo]())
}
