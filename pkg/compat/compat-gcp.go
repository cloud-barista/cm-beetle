package compat

import (
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckGcp checks compatibility between GCP VM spec and OS image
func CheckGcp(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	log.Trace().Msgf("Starting GCP compatibility check for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)

	// TODO: Add GCP-specific compatibility checks using Detail information
	log.Debug().Msg("GCP compatibility validation is planned for future implementation")

	log.Trace().Msgf("GCP compatibility check passed for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
	return true
}

// === GCP-Specific Helper Functions ===
// TODO: Add GCP-specific compatibility validation functions as needed
// Examples:
// - Machine type family compatibility
// - GPU availability checks
// - Regional capability validation
// - Disk type compatibility
// - Network performance requirements

// === CPU Vendor Detection ===

// gcpAmdSeriesPrefixes lists GCP machine type series known to run on AMD EPYC.
var gcpAmdSeriesPrefixes = map[string]bool{"n2d": true, "c2d": true, "c3d": true, "c4d": true, "t2d": true}

// gcpIntelSeriesPrefixes lists GCP machine type series known to run on Intel Xeon.
//
// e2, t2a, a2, a3, and g2 are deliberately excluded: e2's underlying CPU vendor is chosen
// automatically by GCP with no documented mapping, t2a is Ampere (ARM, out of scope for
// amd/intel), and a2/a3/g2 are GPU-attached families whose host CPU vendor hasn't been analyzed.
var gcpIntelSeriesPrefixes = map[string]bool{
	"n1": true, "n2": true, "n4": true,
	"c2": true, "c3": true, "c4": true,
	"m1": true, "m2": true, "m3": true,
}

// getGcpCpuVendor classifies the CPU vendor of a GCP machine type from its series prefix (the
// part before the first "-", e.g. "n2d" in "n2d-standard-4"). Returns "amd", "intel", or ""
// (unclassified - includes e2, ARM, and unanalyzed GPU families).
func getGcpCpuVendor(cspSpecName string) string {
	name := strings.ToLower(cspSpecName)
	series := name
	if idx := strings.Index(name, "-"); idx >= 0 {
		series = name[:idx]
	}

	switch {
	case gcpAmdSeriesPrefixes[series]:
		return "amd"
	case gcpIntelSeriesPrefixes[series]:
		return "intel"
	default:
		return ""
	}
}
