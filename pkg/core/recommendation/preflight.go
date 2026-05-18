package recommendation

import (
	"strings"

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

// spiderKnownDiskTypes returns the root disk type names CB-Spider recognizes for a CSP.
// Values are sourced from the rootdisktype field in:
// https://raw.githubusercontent.com/cloud-barista/cb-tumblebug/refs/heads/main/assets/spider/cloudos_meta.yaml
// (CB-Tumblebug synchronizes this file from CB-Spider.)
// Returns nil for unknown CSPs (no filtering applied).
func spiderKnownDiskTypes(csp string) []string {
	switch strings.ToLower(csp) {
	case "aws":
		return []string{"standard", "gp2", "gp3"}
	case "azure":
		return []string{"PremiumSSD", "StandardSSD", "StandardHDD"}
	case "gcp":
		return []string{"pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"}
	case "alibaba":
		// cloud_auto and cloud_essd_entry are valid on Alibaba but not yet listed in cloudos_meta.yaml.
		return []string{"cloud_essd", "cloud_efficiency", "cloud", "cloud_ssd"}
	case "tencent":
		return []string{"CLOUD_PREMIUM", "CLOUD_SSD"}
	case "ibm":
		return []string{"general-purpose", "sdp"}
	case "ncp", "ncpvpc":
		return []string{"HDD"}
	case "nhn":
		return []string{"General_HDD", "General_SSD"}
	case "kt", "ktvpc":
		return []string{"HDD", "SSD"}
	case "ktclassic":
		return []string{"HDD", "SSD"}
	default:
		return nil // unknown CSP — pass through unchanged
	}
}

// validateSuggestedDiskType returns diskType if CB-Spider recognizes it for the CSP,
// or "" if not, falling back to the CSP's own default.
func validateSuggestedDiskType(csp, diskType string) string {
	if diskType == "" {
		return ""
	}
	known := spiderKnownDiskTypes(csp)
	if known == nil {
		return diskType // unknown CSP — pass through unchanged
	}
	for _, t := range known {
		if t == diskType {
			return diskType
		}
	}
	log.Warn().Msgf("SuggestedSystemDisk %q is not recognized by CB-Spider for %s; falling back to CSP default", diskType, csp)
	return ""
}

// PreflightCheckCspProvisioning calls POST /specImagePairReview and returns image availability,
// the resolved latest CSP image name, and the suggested system disk for the spec+zone.
func PreflightCheckCspProvisioning(specId, imageId, currentCspImageName, rootDiskType string) (CspProvisioningPrecheck, error) {
	empty := CspProvisioningPrecheck{ResolvedCspImageName: currentCspImageName}

	// Extract CSP from specId (format: "csp+region+instanceType") for disk-type validation.
	csp := ""
	if parts := strings.SplitN(specId, "+", 2); len(parts) >= 1 {
		csp = parts[0]
	}

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

	suggestedDisk := validateSuggestedDiskType(csp, result.SuggestedSystemDisk)

	if !result.ImageValidation.IsAvailable {
		return CspProvisioningPrecheck{
			ResolvedCspImageName: currentCspImageName,
			IsAvailable:          false,
			SuggestedSystemDisk:  suggestedDisk,
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
		SuggestedSystemDisk:  suggestedDisk,
	}, nil
}
