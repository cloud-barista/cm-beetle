package compat

import (
	"regexp"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckAzure checks compatibility between Azure VM spec and OS image
// Primary focus: Hypervisor Generation compatibility (V1 vs V2)
func CheckAzure(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	log.Trace().Msgf("Starting Azure compatibility check for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)

	// 1. Hypervisor Generation compatibility check (most critical for Azure)
	if !isAzureHypervisorGenerationCompatible(spec, image) {
		log.Trace().Msgf("Azure hypervisor generation compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	// 2. NVMe support compatibility check (for v6 series and newer)
	// TODO: Enable when Azure provides consistent NVMe support information in VM spec and image Details
	/*
		if !isAzureNvmeSupportCompatible(spec, image) {
			log.Trace().Msgf("Azure NVMe support compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
			return false
		}
	*/

	log.Trace().Msgf("Azure compatibility check passed for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
	return true
}

// isAzureHypervisorGenerationCompatible checks hypervisor generation compatibility
func isAzureHypervisorGenerationCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	specGeneration := getAzureVmGeneration(spec.CspSpecName)
	imageGeneration := getAzureImageGeneration(image)

	log.Trace().Msgf("Azure generation check - VM: %s (%s), Image: %s (%s)",
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
		`^basic_a\d+`,           // Basic A-series (classic)
		`^standard_a\d+`,        // Standard A-series (classic)
		`^standard_a\d+_v\d+`,   // A-series versions (mostly Gen1 only)
		`^standard_g\d+`,        // G-series (legacy, Gen1 only)
		`^standard_gs\d+`,       // GS-series (legacy, Gen1 only)
		`^standard_d[1-4]_v3$`,  // Small D-v3 series (D1_v3, D2_v3, D3_v3, D4_v3) are Gen1 only
		`^standard_ds[1-4]_v3$`, // Small DS-v3 series (DS1_v3, DS2_v3, DS3_v3, DS4_v3) are Gen1 only
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

// === 2. CPU Vendor Detection ===

// azureAmdPatterns lists Azure VM size regex patterns for AMD EPYC-based series, identified by the
// "a" additive-feature letter in the size name per Azure's VM naming convention
// (https://learn.microsoft.com/en-us/azure/virtual-machines/vm-naming-conventions). Each pattern
// requires a digit immediately after the family letter(s), so e.g. the D-series pattern can never
// match a DC-series name (no digit follows "d" there) - the two lists below are disjoint by
// construction, not by ordering.
var azureAmdPatterns = []string{
	`^standard_d\d+as_v\d+$`,   // Dasv3/4/5/6  - AMD general purpose
	`^standard_d\d+ads_v\d+$`,  // Dadsv5/6     - AMD general purpose + local disk
	`^standard_e\d+as_v\d+$`,   // Easv4/5/6    - AMD memory optimized
	`^standard_e\d+ads_v\d+$`,  // Eadsv5/6     - AMD memory optimized + local disk
	`^standard_f\d+as_v\d+$`,   // Fasv6        - AMD compute optimized
	`^standard_dc\d+as_v\d+$`,  // DCasv5       - AMD confidential computing
	`^standard_dc\d+ads_v\d+$`, // DCadsv5      - AMD confidential computing + local disk
	`^standard_ec\d+as_v\d+$`,  // ECasv5       - AMD confidential memory optimized
	`^standard_ec\d+ads_v\d+$`, // ECadsv5      - AMD confidential memory optimized + local disk
}

// azureIntelPatterns lists the Intel-equivalent sibling families of azureAmdPatterns (same
// family/generation, without the "a" AMD additive letter).
//
// This is deliberately an explicit allow-list rather than "everything not matched as AMD is
// Intel": families such as HBv2/HBv3 (HPC) are AMD EPYC-based without any "a" in the name, and
// GPU/legacy/burstable/large-memory families haven't been analyzed here. A blanket
// not-AMD-therefore-Intel fallback would silently mislabel those as Intel.
var azureIntelPatterns = []string{
	`^standard_d\d+s?_v\d+$`,  // Dv4/Dsv4, Dv5/Dsv5, Dv6/Dsv6 - Intel
	`^standard_d\d+ds_v\d+$`,  // Ddsv4/Ddsv5/Ddsv6            - Intel + local disk
	`^standard_e\d+s?_v\d+$`,  // Ev4/Esv4, Ev5/Esv5, Ev6/Esv6 - Intel
	`^standard_e\d+ds_v\d+$`,  // Edsv4/Edsv5/Edsv6            - Intel + local disk
	`^standard_f\d+s?_v\d+$`,  // Fsv2 and successors         - Intel compute optimized
	`^standard_dc\d+s_v\d+$`,  // DCsv3+                      - Intel confidential computing
	`^standard_dc\d+ds_v\d+$`, // DCdsv3+                     - Intel confidential + local disk
	`^standard_ec\d+s_v\d+$`,  // ECsv5                       - Intel confidential memory optimized
	`^standard_ec\d+ds_v\d+$`, // ECdsv5                      - Intel confidential + local disk
}

// getAzureCpuVendor classifies the CPU vendor of an Azure VM size from its CspSpecName. Returns
// "amd", "intel", or "" if unclassified (e.g. classic A-series, M-series, HB/HC HPC series, GPU
// series, or any family not enumerated above). "" means "not classified", never "assumed Intel" -
// callers must treat it as unknown.
func getAzureCpuVendor(cspSpecName string) string {
	name := strings.ToLower(cspSpecName)
	for _, p := range azureAmdPatterns {
		if matched, _ := regexp.MatchString(p, name); matched {
			return "amd"
		}
	}
	for _, p := range azureIntelPatterns {
		if matched, _ := regexp.MatchString(p, name); matched {
			return "intel"
		}
	}
	return ""
}

// === 3. NVMe Support Compatibility (for v6 series and newer) ===
// TODO: Enable when Azure provides consistent NVMe support information in VM spec and image Details

/*
// isAzureNvmeSupportCompatible checks NVMe support compatibility between spec and image
func isAzureNvmeSupportCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	specNvmeSupport := extractAzureNvmeSupportFromSpecDetails(spec)
	imageNvmeSupport := extractAzureNvmeSupportFromImageDetails(image)

	log.Trace().Msgf("Azure NVMe support check - Spec: %s (%s), Image: %s (%s)",
		spec.CspSpecName, specNvmeSupport, image.CspImageName, imageNvmeSupport)

	// If no NVMe info available, assume compatible with different confidence levels
	if specNvmeSupport == "" && imageNvmeSupport == "" {
		log.Trace().Msgf("Azure NVMe support info completely missing, assuming compatible")
		return true
	} else if specNvmeSupport == "" {
		// Only image info available - be permissive since we don't know spec requirements
		log.Trace().Msgf("Azure spec NVMe support unknown, image: %s, assuming compatible", imageNvmeSupport)
		return true
	} else if imageNvmeSupport == "" {
		// Only spec info available - be permissive since most modern Azure images support NVMe
		if specNvmeSupport == "required" {
			log.Trace().Msgf("Azure spec requires NVMe but image support unknown, assuming compatible (risky)")
		} else {
			log.Trace().Msgf("Azure spec NVMe: %s, image support unknown, assuming compatible", specNvmeSupport)
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
			log.Trace().Msgf("Azure NVMe optimal - Spec supports NVMe, Image supports NVMe (optimal performance)")
			return true
		} else if imageNvmeSupport == "unsupported" {
			// WARNING: This combination is risky - NVMe hardware without NVMe drivers
			// Most modern NVMe SSDs cannot fall back to SATA/AHCI compatibility mode
			log.Trace().Msgf("Azure NVMe risky - Spec supports NVMe, Image doesn't support NVMe (may fail to boot)")
			return false
		} else {
			log.Trace().Msgf("Azure NVMe unknown - Spec supports NVMe, Image NVMe support unknown (assuming compatible)")
			return true
		}

	case "unsupported":
		// Instance doesn't support NVMe hardware
		// Images with NVMe drivers are still compatible (drivers just won't be used)
		// Only incompatible if image REQUIRES NVMe
		if imageNvmeSupport == "required" {
			log.Trace().Msgf("Azure NVMe incompatible - Spec doesn't support NVMe, but Image requires it")
			return false
		}
		log.Trace().Msgf("Azure NVMe compatible - Spec unsupported, Image %s (NVMe drivers will be unused)", imageNvmeSupport)
		return true

	default:
		log.Trace().Msgf("Unknown NVMe support value: %s", specNvmeSupport)
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
