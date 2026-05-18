package recommendation

import (
	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"
)

// CspProvisioningPrecheck is the result of a POST /specImagePairReview call.
type CspProvisioningPrecheck struct {
	// ResolvedCspImageName is the latest CSP image name. Set on CreateNodeGroupReq.CspImageName
	// so TumbleBug bypasses its DB cache and passes the fresh name directly to Spider.
	ResolvedCspImageName string

	// IsAvailable reports whether the image is currently available on the CSP.
	IsAvailable bool

	// SuggestedSystemDisk is the disk category confirmed available for the spec+zone.
	// Empty = no suggestion; use CSP default. Set on CreateNodeGroupReq.RootDiskType.
	SuggestedSystemDisk string
}

// PreflightCheckCspProvisioning calls POST /specImagePairReview and returns image availability,
// the resolved latest CSP image name, and the suggested system disk for the spec+zone.
func PreflightCheckCspProvisioning(specId, imageId, currentCspImageName, rootDiskType string) (CspProvisioningPrecheck, error) {
	empty := CspProvisioningPrecheck{ResolvedCspImageName: currentCspImageName}

	result, err := tbclient.NewSession().ReviewSpecImagePair(tbmodel.SpecImagePairReviewReq{
		SpecId:       specId,
		ImageId:      imageId,
		RootDiskType: rootDiskType,
	})
	if err != nil {
		return empty, err
	}

	for _, w := range result.Warnings {
		log.Warn().Msgf("spec-image review warning (specId: %s, imageId: %s): %s", specId, imageId, w)
	}

	if !result.ImageValidation.IsAvailable {
		return CspProvisioningPrecheck{
			ResolvedCspImageName: currentCspImageName,
			IsAvailable:          false,
			SuggestedSystemDisk:  result.SuggestedSystemDisk,
		}, nil
	}

	resolvedName := currentCspImageName
	if result.ImageValidation.CspResourceId != "" && result.ImageValidation.CspResourceId != currentCspImageName {
		log.Info().Msgf("CSP image resolved to latest: %s -> %s (specId: %s)",
			currentCspImageName, result.ImageValidation.CspResourceId, specId)
		resolvedName = result.ImageValidation.CspResourceId
	}

	return CspProvisioningPrecheck{
		ResolvedCspImageName: resolvedName,
		IsAvailable:          true,
		SuggestedSystemDisk:  result.SuggestedSystemDisk,
	}, nil
}
