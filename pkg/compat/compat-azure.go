package compat

import (
	"regexp"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckAzure checks compatibility between Azure VM spec and OS image
// Primary focus: Hypervisor Generation compatibility (V1 vs V2)
func CheckAzure(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	log.Debug().Msgf("Starting Azure compatibility check for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)

	// 1. Hypervisor Generation compatibility check (most critical for Azure)
	if !isAzureHypervisorGenerationCompatible(spec, image) {
		log.Debug().Msgf("Azure hypervisor generation compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	// 2. NVMe support compatibility check (for v6 series and newer)
	// TODO: Enable when Azure provides consistent NVMe support information in VM spec and image Details
	/*
		if !isAzureNvmeSupportCompatible(spec, image) {
			log.Debug().Msgf("Azure NVMe support compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
			return false
		}
	*/

	log.Debug().Msgf("Azure compatibility check passed for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
	return true
}

// isAzureHypervisorGenerationCompatible checks hypervisor generation compatibility
func isAzureHypervisorGenerationCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	specGeneration := getAzureVmGeneration(spec.CspSpecName)
	imageGeneration := getAzureImageGeneration(image)

	log.Debug().Msgf("Azure generation check - VM: %s (%s), Image: %s (%s)",
		spec.CspSpecName, specGeneration, image.CspImageName, imageGeneration)

	// Critical compatibility rules based on Azure documentation
	switch specGeneration {
	case "Gen1Only":
		// Gen1-only VMs can only boot Generation 1 images
		return imageGeneration == "V1" || imageGeneration == "Generation1"

	case "GenBoth":
		// VMs that support both generations can boot both Gen1 and Gen2 images
		return true

	case "Gen2Only":
		// Gen2-only VMs can only boot Generation 2 images
		return imageGeneration == "V2" || imageGeneration == "Generation2"

	default:
		log.Warn().Msgf("Unknown VM generation: %s", specGeneration)
		return true
	}
}

// getAzureVmGeneration gets VM generation support
func getAzureVmGeneration(vmSize string) string {
	vmSizeLower := strings.ToLower(vmSize)

	// Gen1-only families (legacy, classic)
	gen1OnlyPatterns := []string{
		`^basic_a\d+`,         // Basic A-series (classic)
		`^standard_a\d+`,      // Standard A-series (classic)
		`^standard_a\d+_v\d+`, // A-series versions (mostly Gen1 only)
		`^standard_g\d+`,      // G-series (legacy, Gen1 only)
		`^standard_gs\d+`,     // GS-series (legacy, Gen1 only)
	}

	for _, pattern := range gen1OnlyPatterns {
		if matched, _ := regexp.MatchString(pattern, vmSizeLower); matched {
			return "Gen1Only"
		}
	}

	// Gen2-only families (newest series)
	gen2OnlyPatterns := []string{
		`^standard_hx\d+`,          // HX series - high memory, Gen2 only
		`^standard_fx\d+`,          // FX series - high memory, Gen2 only
		`^standard_dc\d+s?_v[3-9]`, // DCv3+ series - confidential computing, Gen2 only
		`^standard_dcas\d+`,        // DCas series - AMD confidential computing, Gen2 only
		`^standard_dcads\d+`,       // DCads series - AMD confidential computing, Gen2 only
		`^standard_ecas\d+`,        // ECas series - AMD memory optimized, Gen2 only
		`^standard_ecads\d+`,       // ECads series - AMD memory optimized, Gen2 only
	}

	for _, pattern := range gen2OnlyPatterns {
		if matched, _ := regexp.MatchString(pattern, vmSizeLower); matched {
			return "Gen2Only"
		}
	}

	// Gen2 supported families: B, D, E, F, L, M, NC, ND, NV, HB, HC series
	return "GenBoth"
}

// getAzureImageGeneration extracts hypervisor generation from image
func getAzureImageGeneration(image cloudmodel.ImageInfo) string {
	// Check image SKU for generation indicators
	imageName := strings.ToLower(image.CspImageName)

	// Common patterns for Gen2 images
	if strings.Contains(imageName, "gen2") || strings.Contains(imageName, "-g2") {
		return "V2"
	}

	// Check image details if available
	for _, detail := range image.Details {
		if detail.Key == "Properties" {
			// Parse hyperVGeneration from Properties JSON string
			if strings.Contains(detail.Value, "hyperVGeneration:V1") {
				return "V1"
			}
			if strings.Contains(detail.Value, "hyperVGeneration:V2") {
				return "V2"
			}
		}
	}

	// Default to V1 if not specified (most common for older images)
	return "V1"
}

// === 2. NVMe Support Compatibility (for v6 series and newer) ===
// TODO: Enable when Azure provides consistent NVMe support information in VM spec and image Details

/*
// isAzureNvmeSupportCompatible checks NVMe support compatibility between spec and image
func isAzureNvmeSupportCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	specNvmeSupport := extractAzureNvmeSupportFromSpecDetails(spec)
	imageNvmeSupport := extractAzureNvmeSupportFromImageDetails(image)

	log.Debug().Msgf("Azure NVMe support check - Spec: %s (%s), Image: %s (%s)",
		spec.CspSpecName, specNvmeSupport, image.CspImageName, imageNvmeSupport)

	// If no NVMe info available, assume compatible with different confidence levels
	if specNvmeSupport == "" && imageNvmeSupport == "" {
		log.Debug().Msgf("Azure NVMe support info completely missing, assuming compatible")
		return true
	} else if specNvmeSupport == "" {
		// Only image info available - be permissive since we don't know spec requirements
		log.Debug().Msgf("Azure spec NVMe support unknown, image: %s, assuming compatible", imageNvmeSupport)
		return true
	} else if imageNvmeSupport == "" {
		// Only spec info available - be permissive since most modern Azure images support NVMe
		if specNvmeSupport == "required" {
			log.Debug().Msgf("Azure spec requires NVMe but image support unknown, assuming compatible (risky)")
		} else {
			log.Debug().Msgf("Azure spec NVMe: %s, image support unknown, assuming compatible", specNvmeSupport)
		}
		return true
	}

	// Apply same logic as AWS and corrected Alibaba logic
	switch specNvmeSupport {
	case "required":
		// Instance requires NVMe, image must support it
		return imageNvmeSupport == "supported" || imageNvmeSupport == "required"

	case "supported":
		// Instance supports NVMe, but compatibility depends on image driver support
		if imageNvmeSupport == "supported" || imageNvmeSupport == "required" {
			log.Debug().Msgf("Azure NVMe optimal - Spec supports NVMe, Image supports NVMe (optimal performance)")
			return true
		} else if imageNvmeSupport == "unsupported" {
			// WARNING: This combination is risky - NVMe hardware without NVMe drivers
			// Most modern NVMe SSDs cannot fall back to SATA/AHCI compatibility mode
			log.Debug().Msgf("Azure NVMe risky - Spec supports NVMe, Image doesn't support NVMe (may fail to boot)")
			return false
		} else {
			log.Debug().Msgf("Azure NVMe unknown - Spec supports NVMe, Image NVMe support unknown (assuming compatible)")
			return true
		}

	case "unsupported":
		// Instance doesn't support NVMe hardware
		// Images with NVMe drivers are still compatible (drivers just won't be used)
		// Only incompatible if image REQUIRES NVMe
		if imageNvmeSupport == "required" {
			log.Debug().Msgf("Azure NVMe incompatible - Spec doesn't support NVMe, but Image requires it")
			return false
		}
		log.Debug().Msgf("Azure NVMe compatible - Spec unsupported, Image %s (NVMe drivers will be unused)", imageNvmeSupport)
		return true

	default:
		log.Debug().Msgf("Unknown NVMe support value: %s", specNvmeSupport)
		return true
	}
}
*/

/*
// extractAzureNvmeSupportFromSpecDetails extracts NVMe support from Azure VM spec details
func extractAzureNvmeSupportFromSpecDetails(spec cloudmodel.SpecInfo) string {
	// First check direct NvmeSupport field in Details
	for _, kv := range spec.Details {
		if strings.EqualFold(kv.Key, "NvmeSupport") {
			return strings.ToLower(strings.TrimSpace(kv.Value))
		}
		// Check for other possible NVMe-related fields
		if strings.EqualFold(kv.Key, "StorageType") || strings.EqualFold(kv.Key, "DiskType") {
			value := strings.ToLower(strings.TrimSpace(kv.Value))
			if strings.Contains(value, "nvme") {
				return "supported"
			}
		}
		// Check for NVMe capability indicators
		if strings.EqualFold(kv.Key, "Capabilities") || strings.EqualFold(kv.Key, "Features") {
			value := strings.ToLower(strings.TrimSpace(kv.Value))
			if strings.Contains(value, "nvme") {
				return "supported"
			}
		}
	}

	// If no explicit NVMe information found in Details, return empty
	// This will be handled by the missing info logic in the main function
	return ""
}
*/

/*
// extractAzureNvmeSupportFromImageDetails extracts NVMe support from Azure image details
func extractAzureNvmeSupportFromImageDetails(image cloudmodel.ImageInfo) string {
	// First check direct NvmeSupport field
	for _, kv := range image.Details {
		if strings.EqualFold(kv.Key, "NvmeSupport") {
			return strings.ToLower(strings.TrimSpace(kv.Value))
		}
	}

	// Azure-specific: Check Properties for DiskControllerTypes or Features
	for _, kv := range image.Details {
		if strings.EqualFold(kv.Key, "Properties") {
			properties := kv.Value
			// Check for DiskControllerTypes containing NVMe
			if strings.Contains(properties, "DiskControllerTypes") {
				if strings.Contains(properties, "NVMe") {
					if strings.Contains(properties, "SCSI") {
						// Image supports both SCSI and NVMe
						return "supported"
					} else {
						// Image supports NVMe only
						return "required"
					}
				} else if strings.Contains(properties, "SCSI") {
					// Image supports SCSI only (no NVMe)
					return "unsupported"
				}
			}
			// Check for NVMe in features
			if strings.Contains(properties, "NVMe") {
				return "supported"
			}
		}
	}

	return ""
}
*/
