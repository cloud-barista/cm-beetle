// Package compat provides compatibility checking functionality between VM specifications and images across different Cloud Service Providers (CSPs).
// This package centralizes all CSP-specific compatibility validation logic.
package compat

import (
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/rs/zerolog/log"
)

// Checker interface defines the contract for CSP-specific compatibility checkers
type Checker interface {
	CheckCompatibility(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool
}

// CheckCompatibility performs compatibility check between spec and image for the specified CSP
func CheckCompatibility(csp string, spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {

	// 1. Architecture check for all CSPs (common check)
	if !isArchitectureCompatible(csp, spec, image) {
		return false
	}

	// 2. CSP-specific compatibility checks using Detail information
	switch strings.ToLower(csp) {
	case "aws":
		return CheckAws(spec, image)
	case "gcp":
		return CheckGcp(spec, image)
	case "azure":
		return CheckAzure(spec, image)
	case "ncp":
		return CheckNcp(spec, image)
	case "alibaba":
		return CheckAlibaba(spec, image)
	case "tencent":
		return CheckTencent(spec, image)
	case "ibm":
		return CheckIbm(spec, image)
	case "nhn":
		return CheckNhn(spec, image)
	case "kt":
		return CheckKt(spec, image)
	case "openstack":
		return CheckOpenstack(spec, image)
	default:
		log.Trace().Msgf("No specific compatibility checks for CSP: %s", csp)
		return true
	}
}

// GetCpuVendor returns a normalized CPU vendor string ("amd", "intel", or "") for a CSP spec,
// derived from whatever signal is available for that CSP (name convention or, for Alibaba,
// structured spec Details). Returns "" for CSPs without a vendor detector or when the spec
// doesn't match a known pattern. Callers must treat "" as unknown/unclassified, never as a
// default vendor.
func GetCpuVendor(csp string, spec cloudmodel.SpecInfo) string {
	switch strings.ToLower(csp) {
	case "aws":
		return getAwsCpuVendor(spec.CspSpecName)
	case "azure":
		return getAzureCpuVendor(spec.CspSpecName)
	case "gcp":
		return getGcpCpuVendor(spec.CspSpecName)
	case "alibaba":
		return getAlibabaCpuVendor(spec)
	case "ibm":
		return getIbmCpuVendor(spec.CspSpecName)
	case "ncp":
		return getNcpCpuVendor(spec.CspSpecName)
	default:
		// tencent, nhn, kt, openstack: not currently enabled recommendation CSPs (see
		// isSupportedCSP in pkg/core/recommendation/infra.go) and not researched - revisit if
		// one of them is enabled.
		return ""
	}
}

// isArchitectureCompatible checks CPU architecture compatibility for all CSPs
func isArchitectureCompatible(csp string, spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	if spec.Architecture != "" && string(image.OSArchitecture) != "" {
		if spec.Architecture != string(image.OSArchitecture) {
			log.Trace().Msgf("%s architecture mismatch - Spec: %s (%s), Image: %s (%s)",
				strings.ToUpper(csp), spec.CspSpecName, spec.Architecture, image.CspImageName, string(image.OSArchitecture))
			return false
		}
		log.Trace().Msgf("%s architecture match - %s", strings.ToUpper(csp), spec.Architecture)
	}
	return true
}
