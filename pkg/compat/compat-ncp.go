package compat

import (
	"regexp"
	"strings"
	"unicode"

	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckNcp checks compatibility between NCP VM spec and OS image
func CheckNcp(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	log.Debug().Msgf("Starting NCP compatibility check for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)

	// NCP Image ID compatibility check using CorrespondingImageIds
	if !isNcpImageCompatible(spec, image) {
		log.Debug().Msgf("NCP image compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	// Add other NCP-specific compatibility checks here if needed
	log.Debug().Msgf("NCP compatibility check passed for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
	return true
}

// === NCP Image Compatibility Functions ===

// isNcpImageCompatible checks if NCP image is compatible with spec using CorrespondingImageIds
func isNcpImageCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	// Extract CorrespondingImageIds from spec details
	correspondingImageIds := extractNcpCorrespondingImageIds(spec)
	if len(correspondingImageIds) == 0 {
		log.Debug().Msgf("No CorrespondingImageIds found for NCP spec: %s, allowing all images", spec.CspSpecName)
		return true // If no corresponding image IDs specified, allow all images
	}

	// Extract image ID from image (could be in CspImageName or other fields)
	imageId := extractNcpImageId(image)
	if imageId == "" {
		log.Debug().Msgf("Could not extract image ID from NCP image: %s", image.CspImageName)
		return true // If we can't extract image ID, be permissive
	}

	// Check if image ID is in the corresponding image IDs list
	for _, correspondingId := range correspondingImageIds {
		if imageId == correspondingId {
			log.Debug().Msgf("NCP image ID %s matches corresponding image ID for spec %s", imageId, spec.CspSpecName)
			return true
		}
	}

	log.Debug().Msgf("NCP image ID %s not found in corresponding image IDs %v for spec %s",
		imageId, correspondingImageIds, spec.CspSpecName)
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
			log.Debug().Msgf("Extracted NCP CorrespondingImageIds: %v", cleanedIds)
			return cleanedIds
		}
	}
	return []string{}
}

// extractNcpImageId extracts image ID from NCP image info
func extractNcpImageId(image cloudmodel.ImageInfo) string {
	// Try to extract from CspImageName first (might contain the ID directly)
	if image.CspImageName != "" {
		log.Debug().Msgf("NCP image CspImageName: %s", image.CspImageName)
		// If CspImageName is numeric, use it directly
		if isNumeric(image.CspImageName) {
			return image.CspImageName
		}
	}

	// Try to extract from Details
	for _, detail := range image.Details {
		if strings.EqualFold(detail.Key, "ImageId") ||
			strings.EqualFold(detail.Key, "Id") ||
			strings.EqualFold(detail.Key, "NcpImageId") {
			if detail.Value != "" && isNumeric(detail.Value) {
				log.Debug().Msgf("Extracted NCP image ID from Details[%s]: %s", detail.Key, detail.Value)
				return detail.Value
			}
		}
	}

	// Fallback: try to extract numeric part from CspImageName
	if image.CspImageName != "" {
		// Look for numeric patterns in the image name
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(image.CspImageName, -1)
		if len(matches) > 0 {
			// Use the first (or longest) numeric match
			longestMatch := ""
			for _, match := range matches {
				if len(match) > len(longestMatch) {
					longestMatch = match
				}
			}
			if longestMatch != "" {
				log.Debug().Msgf("Extracted NCP image ID from CspImageName pattern: %s", longestMatch)
				return longestMatch
			}
		}
	}

	log.Debug().Msgf("Could not extract numeric image ID from NCP image: %s", image.CspImageName)
	return ""
}

// isNumeric checks if a string contains only digits
func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// === NCP VM Spec Filtering Functions ===

// FilterNcpVmSpecsByHypervisor filters NCP VM specs to include only KVM hypervisor specs
func FilterNcpVmSpecsByHypervisor(vmSpecs []cloudmodel.SpecInfo) []cloudmodel.SpecInfo {
	if len(vmSpecs) == 0 {
		return vmSpecs
	}

	log.Debug().Msgf("NCP filtering: checking %d VM specs for KVM hypervisor", len(vmSpecs))

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
			log.Debug().Msgf("NCP: included VM spec %s (KVM hypervisor found)", spec.CspSpecName)
		} else {
			log.Debug().Msgf("NCP: filtered out VM spec %s (no KVM hypervisor)", spec.CspSpecName)
		}
	}

	log.Debug().Msgf("NCP filtering: %d VM specs filtered to %d KVM specs", len(vmSpecs), len(filteredSpecs))

	// If no KVM specs found, return original list with warning
	if len(filteredSpecs) == 0 {
		log.Warn().Msg("No KVM hypervisor specs found for NCP, returning all specs")
		return vmSpecs
	}

	return filteredSpecs
}
