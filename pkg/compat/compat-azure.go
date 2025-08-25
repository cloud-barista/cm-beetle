package compat

import (
	"regexp"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckAzure checks compatibility between Azure VM spec and OS image
// Primary focus: Hypervisor Generation compatibility (V1 vs V2)
func CheckAzure(spec cloudmodel.TbSpecInfo, image cloudmodel.TbImageInfo) bool {
	log.Debug().Msgf("Starting Azure compatibility check for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)

	// 1. Hypervisor Generation compatibility check (most critical for Azure)
	if !isAzureHypervisorGenerationCompatible(spec, image) {
		log.Debug().Msgf("Azure hypervisor generation compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	log.Debug().Msgf("Azure compatibility check passed for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
	return true
}

// isAzureHypervisorGenerationCompatible checks hypervisor generation compatibility
func isAzureHypervisorGenerationCompatible(spec cloudmodel.TbSpecInfo, image cloudmodel.TbImageInfo) bool {
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
func getAzureImageGeneration(image cloudmodel.TbImageInfo) string {
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
