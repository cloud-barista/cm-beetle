// Package compat provides compatibility checking functionality between VM specifications and images across different Cloud Service Providers (CSPs).
// This package centralizes all CSP-specific compatibility validation logic.
package compat

import (
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/csp"
	"github.com/rs/zerolog/log"
)

// Checker interface defines the contract for CSP-specific compatibility checkers
type Checker interface {
	CheckCompatibility(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool
}

// CheckCompatibility performs compatibility check between spec and image for the specified CSP
func CheckCompatibility(cspName string, spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {

	// 1. Architecture check for all CSPs (common check)
	if !isArchitectureCompatible(cspName, spec, image) {
		return false
	}

	// 2. CSP-specific compatibility checks using Detail information
	switch strings.ToLower(cspName) {
	case csp.AWS:
		return CheckAws(spec, image)
	case csp.GCP:
		return CheckGcp(spec, image)
	case csp.Azure:
		return CheckAzure(spec, image)
	case csp.NCP:
		return CheckNcp(spec, image)
	case csp.Alibaba:
		return CheckAlibaba(spec, image)
	case csp.Tencent:
		return CheckTencent(spec, image)
	case csp.IBM:
		return CheckIbm(spec, image)
	case csp.NHN:
		return CheckNhn(spec, image)
	case csp.KT:
		return CheckKt(spec, image)
	case csp.OpenStack:
		return CheckOpenstack(spec, image)
	default:
		log.Trace().Msgf("No specific compatibility checks for CSP: %s", cspName)
		return true
	}
}

// GetCpuVendor returns a normalized CPU vendor string ("amd", "intel", or "") for a CSP spec,
// derived from whatever signal is available for that CSP (name convention or, for Alibaba,
// structured spec Details). Returns "" for CSPs without a vendor detector or when the spec
// doesn't match a known pattern. Callers must treat "" as unknown/unclassified, never as a
// default vendor.
func GetCpuVendor(cspName string, spec cloudmodel.SpecInfo) string {
	switch strings.ToLower(cspName) {
	case csp.AWS:
		return getAwsCpuVendor(spec.CspSpecName)
	case csp.Azure:
		return getAzureCpuVendor(spec.CspSpecName)
	case csp.GCP:
		return getGcpCpuVendor(spec.CspSpecName)
	case csp.Alibaba:
		return getAlibabaCpuVendor(spec)
	case csp.IBM:
		return getIbmCpuVendor(spec.CspSpecName)
	case csp.NCP:
		return getNcpCpuVendor(spec.CspSpecName)
	default:
		// Return empty for CSPs without vendor detectors (tencent, nhn, kt, openstack).
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
