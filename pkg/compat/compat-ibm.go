package compat

import (
	"regexp"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckIbm checks compatibility between IBM Cloud VM spec and OS image
func CheckIbm(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	log.Trace().Msgf("Starting IBM Cloud compatibility check for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)

	// TODO: Add IBM Cloud-specific compatibility checks using Detail information
	log.Debug().Msg("IBM Cloud compatibility validation is planned for future implementation")

	log.Trace().Msgf("IBM Cloud compatibility check passed for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
	return true
}

// === IBM Cloud-Specific Helper Functions ===
// TODO: Add IBM Cloud-specific compatibility validation functions as needed
// Examples:
// - Virtual Server Instance profile compatibility
// - Bare metal vs virtual server requirements
// - GPU and FPGA support validation
// - Network interface requirements
// - Storage type compatibility (Block, File, Object)

// === CPU Vendor Detection ===

// ibmAmdPatterns lists IBM VPC instance profiles confirmed AMD-based (IBM Cloud VPC docs: High
// Frequency profiles run on AMD 5th Gen EPYC 9575F).
var ibmAmdPatterns = []string{
	`^hx\d+a`,  // hx4a  - High Frequency, AMD
	`^hx\d+da`, // hx4da - High Frequency + local disk, AMD
}

// ibmAmbiguousPrefixes lists Flex profiles, which run on a mixed Intel/AMD host pool per IBM's
// docs - vendor isn't determinable from the name, so these must not fall through to the Intel
// default below (they'd otherwise match it: "bxf" starts with "bx").
var ibmAmbiguousPrefixes = []string{"nxf", "bxf", "cxf", "mxf"}

// ibmIntelPrefixes lists profile family prefixes IBM's VPC docs confirm run exclusively on Intel
// Xeon. An explicit allow-list, not a "not-AMD-therefore-Intel" fallback - GPU profiles (gx*)
// aren't documented either way and are deliberately left unclassified.
var ibmIntelPrefixes = []string{"bx", "cx", "mx", "vx", "ux", "ox"}

// getIbmCpuVendor classifies the CPU vendor of an IBM VPC instance profile from its name. Returns
// "amd", "intel", or "" if unclassified (Flex profiles, GPU profiles, or any unrecognized family).
func getIbmCpuVendor(cspSpecName string) string {
	name := strings.ToLower(cspSpecName)

	for _, p := range ibmAmdPatterns {
		if matched, _ := regexp.MatchString(p, name); matched {
			return "amd"
		}
	}
	for _, prefix := range ibmAmbiguousPrefixes {
		if strings.HasPrefix(name, prefix) {
			return ""
		}
	}
	for _, prefix := range ibmIntelPrefixes {
		if strings.HasPrefix(name, prefix) {
			return "intel"
		}
	}
	return ""
}
