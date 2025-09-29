package compat

import (
	"strings"

	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckAlibaba checks compatibility between Alibaba Cloud VM spec and OS image
// Primary focus: NVMe support, Boot mode, and Disk compatibility
func CheckAlibaba(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	log.Debug().Msgf("Starting Alibaba Cloud compatibility check for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)

	// 1. NVMe support compatibility check (most critical)
	if !isAlibabaNvmeSupportCompatible(spec, image) {
		log.Debug().Msgf("Alibaba NVMe support compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	// 2. Boot mode compatibility check
	if !isAlibabaBootModeCompatible(spec, image) {
		log.Debug().Msgf("Alibaba boot mode compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	// 3. Instance category and disk compatibility check
	if !isAlibabaInstanceCategoryCompatible(spec, image) {
		log.Debug().Msgf("Alibaba instance category compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	log.Debug().Msgf("Alibaba Cloud compatibility check passed for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
	return true
}

// === 1. NVMe Support Compatibility (Most Critical) ===

// isAlibabaNvmeSupportCompatible checks NVMe support compatibility between spec and image
// This addresses the "No AvailableSystemDisk" error seen in logs
func isAlibabaNvmeSupportCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	specNvmeSupport := extractAlibabaNvmeSupportFromSpecDetails(spec)
	imageNvmeSupport := extractAlibabaNvmeSupportFromImageDetails(image)

	log.Debug().Msgf("Alibaba NVMe support check - Spec: %s (%s), Image: %s (%s)",
		spec.CspSpecName, specNvmeSupport, image.CspImageName, imageNvmeSupport)

	// If no NVMe info available, assume compatible with different confidence levels
	if specNvmeSupport == "" && imageNvmeSupport == "" {
		log.Debug().Msgf("Alibaba NVMe support info completely missing, assuming compatible")
		return true
	} else if specNvmeSupport == "" {
		// Only image info available - be permissive since we don't know spec requirements
		log.Debug().Msgf("Alibaba spec NVMe support unknown, image: %s, assuming compatible", imageNvmeSupport)
		return true
	} else if imageNvmeSupport == "" {
		// Only spec info available - be permissive since most modern images support NVMe
		if specNvmeSupport == "required" {
			log.Debug().Msgf("Alibaba spec requires NVMe but image support unknown, assuming compatible (risky)")
		} else {
			log.Debug().Msgf("Alibaba spec NVMe: %s, image support unknown, assuming compatible", specNvmeSupport)
		}
		return true
	}

	// Critical compatibility rules for Alibaba Cloud
	switch specNvmeSupport {
	case "required":
		// Instance requires NVMe, image must support it
		return imageNvmeSupport == "supported" || imageNvmeSupport == "required"

	case "supported":
		// Instance supports NVMe, but compatibility depends on image driver support
		if imageNvmeSupport == "supported" || imageNvmeSupport == "required" {
			log.Debug().Msgf("Alibaba NVMe optimal - Spec supports NVMe, Image supports NVMe (optimal performance)")
			return true
		} else if imageNvmeSupport == "unsupported" {
			// WARNING: This combination is risky - NVMe hardware without NVMe drivers
			// Most modern NVMe SSDs cannot fall back to SATA/AHCI compatibility mode
			log.Debug().Msgf("Alibaba NVMe risky - Spec supports NVMe, Image doesn't support NVMe (may fail to boot)")
			return false // Changed from true to false - this combination is actually problematic
		} else {
			log.Debug().Msgf("Alibaba NVMe unknown - Spec supports NVMe, Image NVMe support unknown (assuming compatible)")
			return true
		}

	case "unsupported":
		// Instance doesn't support NVMe hardware
		// Images with NVMe drivers are still compatible (drivers just won't be used)
		// Only incompatible if image REQUIRES NVMe (which is rare)
		if imageNvmeSupport == "required" {
			log.Debug().Msgf("Alibaba NVMe incompatible - Spec doesn't support NVMe, but Image requires it")
			return false
		}
		// Image with NVMe driver support on non-NVMe hardware is fine
		log.Debug().Msgf("Alibaba NVMe compatible - Spec unsupported, Image %s (NVMe drivers will be unused)", imageNvmeSupport)
		return true

	default:
		log.Debug().Msgf("Unknown NVMe support value: %s", specNvmeSupport)
		return true
	}
}

// === 2. Boot Mode Compatibility ===

// isAlibabaBootModeCompatible checks boot mode compatibility between spec and image
func isAlibabaBootModeCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	specBootModes := extractAlibabaSupportedBootModesFromSpecDetails(spec)
	imageBootMode := extractAlibabaBootModeFromImageDetails(image)

	log.Debug().Msgf("Alibaba boot mode check - Spec supports: %v, Image: %s",
		specBootModes, imageBootMode)

	// If no boot mode info, assume compatible
	if len(specBootModes) == 0 || imageBootMode == "" {
		log.Debug().Msgf("Alibaba boot mode info missing, assuming compatible")
		return true
	}

	// Handle UEFI-Preferred case
	if imageBootMode == "UEFI-Preferred" {
		// UEFI-Preferred images can work with both BIOS and UEFI
		for _, mode := range specBootModes {
			if strings.EqualFold(mode, "UEFI") || strings.EqualFold(mode, "BIOS") {
				return true
			}
		}
	}

	// Check if image's boot mode is supported by the spec
	for _, supportedMode := range specBootModes {
		if strings.EqualFold(imageBootMode, supportedMode) {
			return true
		}
	}

	log.Debug().Msgf("Alibaba boot mode incompatible - Spec supports: %v, Image requires: %s",
		specBootModes, imageBootMode)
	return false
}

// === 3. Instance Category and Disk Compatibility ===

// isAlibabaInstanceCategoryCompatible checks instance category and disk compatibility
func isAlibabaInstanceCategoryCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	specCategory := extractAlibabaInstanceCategoryFromSpecDetails(spec)
	specDiskQuantity := extractAlibabaDiskQuantityFromSpecDetails(spec)
	imageIoOptimized := extractAlibabaIoOptimizedFromImageDetails(image)

	log.Debug().Msgf("Alibaba instance category check - Spec: %s (disks: %s), Image IoOptimized: %s",
		specCategory, specDiskQuantity, imageIoOptimized)

	// Special handling for shared instances with limited disk support
	if specCategory == "Shared" && specDiskQuantity != "" {
		if specDiskQuantity == "17" && imageIoOptimized == "true" {
			// Shared instances with limited disk quantity might have issues with I/O optimized images
			log.Debug().Msgf("Alibaba potential compatibility issue - Shared instance with I/O optimized image")
			return true // Allow but log warning
		}
	}

	// Enhanced instances generally work with all image types
	if specCategory == "Enhanced" || specCategory == "EnterpriseLevel" {
		return true
	}

	return true
}

// === Helper Functions for Extracting Details ===

// extractAlibabaNvmeSupportFromSpecDetails extracts NVMe support from VM spec details
func extractAlibabaNvmeSupportFromSpecDetails(spec cloudmodel.SpecInfo) string {
	for _, kv := range spec.Details {
		if strings.EqualFold(kv.Key, "NvmeSupport") {
			return strings.ToLower(strings.TrimSpace(kv.Value))
		}
	}
	return ""
}

// extractAlibabaNvmeSupportFromImageDetails extracts NVMe support from image details
func extractAlibabaNvmeSupportFromImageDetails(image cloudmodel.ImageInfo) string {
	for _, kv := range image.Details {
		if strings.EqualFold(kv.Key, "Features") {
			// Parse Features JSON-like string for NvmeSupport
			features := kv.Value
			if strings.Contains(features, "NvmeSupport:supported") {
				return "supported"
			} else if strings.Contains(features, "NvmeSupport:required") {
				return "required"
			} else if strings.Contains(features, "NvmeSupport:unsupported") {
				return "unsupported"
			}
		}
	}
	return ""
}

// extractAlibabaSupportedBootModesFromSpecDetails extracts supported boot modes from VM spec details
func extractAlibabaSupportedBootModesFromSpecDetails(spec cloudmodel.SpecInfo) []string {
	for _, kv := range spec.Details {
		if strings.EqualFold(kv.Key, "SupportedBootModes") {
			// Parse format: {SupportedBootMode:[BIOS,UEFI]}
			value := kv.Value
			if strings.Contains(value, "SupportedBootMode:") {
				start := strings.Index(value, "[")
				end := strings.Index(value, "]")
				if start != -1 && end != -1 && end > start {
					modeList := value[start+1 : end]
					modes := strings.Split(modeList, ",")
					var result []string
					for _, mode := range modes {
						cleaned := strings.TrimSpace(mode)
						if cleaned != "" {
							result = append(result, cleaned)
						}
					}
					return result
				}
			}
		}
	}
	return []string{}
}

// extractAlibabaBootModeFromImageDetails extracts boot mode from image details
func extractAlibabaBootModeFromImageDetails(image cloudmodel.ImageInfo) string {
	for _, kv := range image.Details {
		if strings.EqualFold(kv.Key, "BootMode") {
			return strings.TrimSpace(kv.Value)
		}
	}
	return ""
}

// extractAlibabaInstanceCategoryFromSpecDetails extracts instance category from VM spec details
func extractAlibabaInstanceCategoryFromSpecDetails(spec cloudmodel.SpecInfo) string {
	for _, kv := range spec.Details {
		if strings.EqualFold(kv.Key, "InstanceCategory") {
			return strings.TrimSpace(kv.Value)
		}
	}
	return ""
}

// extractAlibabaDiskQuantityFromSpecDetails extracts disk quantity from VM spec details
func extractAlibabaDiskQuantityFromSpecDetails(spec cloudmodel.SpecInfo) string {
	for _, kv := range spec.Details {
		if strings.EqualFold(kv.Key, "DiskQuantity") {
			return strings.TrimSpace(kv.Value)
		}
	}
	return ""
}

// extractAlibabaIoOptimizedFromImageDetails extracts I/O optimized info from image details
func extractAlibabaIoOptimizedFromImageDetails(image cloudmodel.ImageInfo) string {
	for _, kv := range image.Details {
		if strings.EqualFold(kv.Key, "IsSupportIoOptimized") {
			return strings.ToLower(strings.TrimSpace(kv.Value))
		}
	}
	return ""
}

// === Alibaba Cloud-Specific Helper Functions ===
// Additional helper functions for future enhancements:
// - Instance family compatibility (ecs.g6, ecs.c6, ecs.e, etc.)
// - Region-specific instance type availability
// - Local disk vs cloud disk compatibility
// - GPU instance type validation
// - Network performance tier requirements
// - Enhanced network features compatibility
