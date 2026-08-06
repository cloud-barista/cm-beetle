package compat

import (
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckNcp checks compatibility between NCP VM spec and OS image
func CheckNcp(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	log.Trace().Msgf("Starting NCP compatibility check for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)

	// NCP Image ID compatibility check using CorrespondingImageIds
	if !isNcpImageCompatible(spec, image) {
		log.Trace().Msgf("NCP image compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	// Add other NCP-specific compatibility checks here if needed
	log.Trace().Msgf("NCP compatibility check passed for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
	return true
}

// === NCP Image Compatibility Functions ===

// isNcpImageCompatible checks if the image's CspImageName exactly matches one of the spec's CorrespondingImageIds.
func isNcpImageCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	// Extract CorrespondingImageIds from spec details
	correspondingImageIds := extractNcpCorrespondingImageIds(spec)
	if len(correspondingImageIds) == 0 {
		log.Trace().Msgf("No CorrespondingImageIds found for NCP spec: %s, allowing all images", spec.CspSpecName)
		return true // If no corresponding image IDs specified, allow all images
	}

	// Check if the image's CspImageName is in the corresponding image IDs list
	for _, correspondingId := range correspondingImageIds {
		if image.CspImageName == correspondingId {
			log.Trace().Msgf("NCP image %s matches corresponding image ID for spec %s", image.CspImageName, spec.CspSpecName)
			return true
		}
	}

	log.Trace().Msgf("NCP image %s not found in corresponding image IDs %v for spec %s",
		image.CspImageName, correspondingImageIds, spec.CspSpecName)
	return false
}

// extractNcpCorrespondingImageIds extracts CorrespondingImageIds from NCP spec details
func extractNcpCorrespondingImageIds(spec cloudmodel.SpecInfo) []string {
	for _, detail := range spec.Details {
		if strings.EqualFold(detail.Key, "CorrespondingImageIds") {
			// Parse comma-separated image IDs
			imageIds := strings.Split(detail.Value, ",")
			var cleanedIds []string
			for _, id := range imageIds {
				cleanedId := strings.TrimSpace(id)
				if cleanedId != "" {
					cleanedIds = append(cleanedIds, cleanedId)
				}
			}
			log.Trace().Msgf("Extracted NCP CorrespondingImageIds: %v", cleanedIds)
			return cleanedIds
		}
	}
	return []string{}
}

// === NCP VM Spec Filtering Functions ===

// FilterNcpVmSpecsByHypervisor filters NCP VM specs to include only KVM hypervisor specs
func FilterNcpVmSpecsByHypervisor(vmSpecs []cloudmodel.SpecInfo) []cloudmodel.SpecInfo {
	if len(vmSpecs) == 0 {
		return vmSpecs
	}

	log.Trace().Msgf("NCP filtering: checking %d VM specs for KVM hypervisor", len(vmSpecs))

	var filteredSpecs []cloudmodel.SpecInfo

	for _, spec := range vmSpecs {
		hasKvmHypervisor := false

		// Check if this spec has KVM hypervisor
		for _, detail := range spec.Details {
			if strings.EqualFold(detail.Key, "hypervisortype") &&
				strings.Contains(strings.ToUpper(detail.Value), "KVM") {
				hasKvmHypervisor = true
				break
			}
		}

		if hasKvmHypervisor {
			filteredSpecs = append(filteredSpecs, spec)
			log.Trace().Msgf("NCP: included VM spec %s (KVM hypervisor found)", spec.CspSpecName)
		} else {
			log.Trace().Msgf("NCP: filtered out VM spec %s (no KVM hypervisor)", spec.CspSpecName)
		}
	}

	log.Trace().Msgf("NCP filtering: %d VM specs filtered to %d KVM specs", len(vmSpecs), len(filteredSpecs))

	// If no KVM specs found, return original list with warning
	if len(filteredSpecs) == 0 {
		log.Warn().Msg("No KVM hypervisor specs found for NCP, returning all specs")
		return vmSpecs
	}

	return filteredSpecs
}
