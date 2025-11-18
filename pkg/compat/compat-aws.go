// compat-aws.go provides AWS-specific compatibility checking functionality
package compat

import (
	"regexp"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckAws checks AWS-specific compatibility following the 4-step validation
// Note: CPU Architecture compatibility is already checked in the common function
// 1. Virtualization Type compatibility
// 2. Hypervisor and Driver compatibility (ENA + NVMe) - Most Critical
// 3. Boot Mode compatibility
// 4. Xen-on-Nitro compatibility (when applicable)
func CheckAws(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	log.Trace().Msgf("Starting AWS compatibility check for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)

	// 1. Virtualization Type compatibility check
	if !isAwsVirtualizationTypeCompatible(spec, image) {
		log.Trace().Msgf("AWS virtualization type compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	// 2. Hypervisor and Driver compatibility check (most critical)
	if !isAwsHypervisorAndDriverCompatible(spec, image) {
		log.Trace().Msgf("AWS hypervisor and driver compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	// 3. Boot Mode compatibility check
	if !isAwsBootModeCompatible(spec, image) {
		log.Trace().Msgf("AWS boot mode compatibility failed - Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
		return false
	}

	log.Trace().Msgf("AWS compatibility check passed for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
	return true
}

// === 1. Virtualization Type Compatibility ===

// isAwsVirtualizationTypeCompatible checks virtualization type compatibility between spec and image
func isAwsVirtualizationTypeCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	supportedVirtTypes := extractAwsSupportedVirtualizationTypesFromSpecDetails(spec)
	imageVirtType := extractAwsVirtualizationTypeFromImageDetails(image)

	if len(supportedVirtTypes) == 0 || imageVirtType == "" {
		log.Trace().Msgf("AWS virtualization type check - insufficient info, assuming compatible")
		return true
	}

	// Check if image's virtualization type is supported by the spec
	for _, supportedType := range supportedVirtTypes {
		if imageVirtType == supportedType {
			log.Trace().Msgf("AWS virtualization type compatible - Spec supports: %v, Image: %s",
				supportedVirtTypes, imageVirtType)
			return true
		}
	}

	log.Trace().Msgf("AWS virtualization type incompatible - Spec supports: %v, Image: %s",
		supportedVirtTypes, imageVirtType)
	return false
}

// extractAwsVirtualizationTypeFromImageDetails extracts virtualization type from VM image details
func extractAwsVirtualizationTypeFromImageDetails(image cloudmodel.ImageInfo) string {
	for _, kv := range image.Details {
		if strings.EqualFold(kv.Key, "virtualizationtype") {
			return strings.ToLower(strings.TrimSpace(kv.Value))
		}
	}
	return ""
}

// extractAwsSupportedVirtualizationTypesFromSpecDetails extracts supported virtualization types from VM spec details
func extractAwsSupportedVirtualizationTypesFromSpecDetails(spec cloudmodel.SpecInfo) []string {
	for _, kv := range spec.Details {
		if strings.EqualFold(kv.Key, "supportedvirtualizationtypes") {
			// Parse supported types (format: "hvm", "hvm; pv", etc.)
			virtTypes := strings.Split(kv.Value, ";")
			var result []string
			for _, vt := range virtTypes {
				cleaned := strings.ToLower(strings.TrimSpace(vt))
				if cleaned != "" {
					result = append(result, cleaned)
				}
			}
			return result
		}
	}
	return []string{}
}

// === 2. Hypervisor and Driver Compatibility (Most Critical) ===

// isAwsHypervisorAndDriverCompatible checks hypervisor and driver compatibility (most critical for AWS)
func isAwsHypervisorAndDriverCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	// Extract hypervisor information
	specHypervisor := extractAwsHypervisorFromSpecDetails(spec)
	imageHypervisor := extractAwsHypervisorFromImageDetails(image)

	log.Trace().Msgf("AWS hypervisor and driver compatibility check - Spec: %s (%s), Image: %s (%s)",
		spec.CspSpecName, specHypervisor, image.CspImageName, imageHypervisor)

	// If no hypervisor info, assume compatible (basic check)
	if specHypervisor == "" || imageHypervisor == "" {
		log.Trace().Msgf("AWS hypervisor info missing, checking drivers only")
		return isAwsDriverCompatible(spec, image)
	}

	// Normalize hypervisor names
	specHyp := normalizeAwsHypervisor(specHypervisor)
	imageHyp := normalizeAwsHypervisor(imageHypervisor)

	// Critical compatibility analysis for Nitro instances + Xen AMI
	if specHyp == "nitro" && imageHyp == "xen" {
		log.Trace().Msgf("AWS CRITICAL: Nitro instance with Xen AMI - checking Xen-on-Nitro compatibility")

		// For Nitro instances with Xen AMIs (Xen-on-Nitro)
		// ENA is required, but NVMe is optional - many Xen AMIs work without NVMe drivers
		enaCompatible := isAwsEnaDriverCompatible(spec, image)
		nvmeCompatible := isAwsNvmeDriverCompatible(spec, image)

		if !enaCompatible {
			log.Trace().Msgf("AWS Xen-on-Nitro incompatible - ENA required but not supported")
			return false
		}

		if !nvmeCompatible {
			log.Trace().Msgf("AWS Xen-on-Nitro compatible - ENA: %t, NVMe: %t (NVMe optional for Xen AMIs)", enaCompatible, nvmeCompatible)
		} else {
			log.Trace().Msgf("AWS Xen-on-Nitro compatible - ENA: %t, NVMe: %t", enaCompatible, nvmeCompatible)
		}
		return true
	}

	// Other combinations: check basic driver compatibility
	return isAwsDriverCompatible(spec, image)
}

// isAwsDriverCompatible checks overall driver compatibility (ENA + NVMe)
func isAwsDriverCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	enaCompatible := isAwsEnaDriverCompatible(spec, image)
	nvmeCompatible := isAwsNvmeDriverCompatible(spec, image)

	log.Trace().Msgf("AWS driver compatibility - ENA: %t, NVMe: %t", enaCompatible, nvmeCompatible)
	return enaCompatible && nvmeCompatible
}

// isAwsEnaDriverCompatible checks ENA driver compatibility between spec and image
func isAwsEnaDriverCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	specEnaSupport := extractAwsEnaSupportFromSpecDetails(spec)
	imageEnaSupport := extractAwsEnaSupportFromImageDetails(image)

	// If spec requires ENA but image doesn't support it: INCOMPATIBLE
	if specEnaSupport == "required" && (imageEnaSupport == "unsupported" || imageEnaSupport == "false") {
		log.Trace().Msgf("AWS ENA incompatible - Spec requires ENA, Image doesn't support it")
		return false
	}

	// Be more permissive: if we don't have clear incompatibility, assume compatible
	// Most modern AMIs support ENA even if not explicitly documented
	if specEnaSupport == "unknown" || imageEnaSupport == "unknown" {
		log.Trace().Msgf("AWS ENA compatibility unknown, assuming compatible")
		return true
	}

	return true
}

// isAwsNvmeDriverCompatible checks NVMe driver compatibility between spec and image
func isAwsNvmeDriverCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	specNvmeSupport := extractAwsNvmeSupportFromSpecDetails(spec)
	imageNvmeSupport := extractAwsNvmeSupportFromImageDetails(image)

	// If spec requires NVMe but image doesn't support it: INCOMPATIBLE
	if specNvmeSupport == "required" && (imageNvmeSupport == "unsupported" || imageNvmeSupport == "false") {
		log.Trace().Msgf("AWS NVMe incompatible - Spec requires NVMe, Image doesn't support it")
		return false
	}

	// Be more permissive: if we don't have clear incompatibility, assume compatible
	// Most modern AMIs support NVMe even if not explicitly documented
	if specNvmeSupport == "unknown" || imageNvmeSupport == "unknown" {
		log.Trace().Msgf("AWS NVMe compatibility unknown, assuming compatible")
		return true
	}

	return true
}

// normalizeAwsHypervisor normalizes hypervisor names for consistent comparison
func normalizeAwsHypervisor(hypervisor string) string {
	hypervisor = strings.ToLower(strings.TrimSpace(hypervisor))
	if strings.Contains(hypervisor, "nitro") {
		return "nitro"
	}
	if strings.Contains(hypervisor, "xen") {
		return "xen"
	}
	return hypervisor
}

// extractAwsHypervisorFromImageDetails extracts hypervisor information from VM image details
func extractAwsHypervisorFromImageDetails(image cloudmodel.ImageInfo) string {
	for _, kv := range image.Details {
		key := strings.ToLower(kv.Key)
		if key == "hypervisor" || key == "hypervisortype" || key == "virtualizationtype" {
			return strings.ToLower(strings.TrimSpace(kv.Value))
		}
	}
	return ""
}

// extractAwsHypervisorFromSpecDetails extracts hypervisor information from VM spec details
func extractAwsHypervisorFromSpecDetails(spec cloudmodel.SpecInfo) string {
	for _, kv := range spec.Details {
		key := strings.ToLower(kv.Key)
		if key == "hypervisor" || key == "hypervisortype" || key == "virtualizationtype" {
			return strings.ToLower(strings.TrimSpace(kv.Value))
		}
	}
	return ""
}

// === 3. ENA Driver Support ===

// extractAwsEnaSupportFromImageDetails extracts ENA support from image details
// Image has direct "EnaSupport" key with value like "true"
func extractAwsEnaSupportFromImageDetails(image cloudmodel.ImageInfo) string {
	for _, kv := range image.Details {
		key := strings.ToLower(strings.TrimSpace(kv.Key))
		value := strings.TrimSpace(kv.Value)

		// Direct ENA support check (EnaSupport key)
		if key == "enasupport" {
			log.Trace().Msgf("Found ENA support in image %s: %s = %s", image.CspImageName, kv.Key, kv.Value)
			return isEnaSupport(value)
		}
	}

	log.Trace().Msgf("No ENA support information found in image %s", image.CspImageName)
	return "unknown"
}

// extractAwsEnaSupportFromSpecDetails extracts ENA support from spec details
// Spec has "NetworkInfo" key with value containing "EnaSupport:required"
func extractAwsEnaSupportFromSpecDetails(spec cloudmodel.SpecInfo) string {
	for _, kv := range spec.Details {
		key := strings.ToLower(strings.TrimSpace(kv.Key))
		value := strings.TrimSpace(kv.Value)

		// Check nested NetworkInfo structure
		if key == "networkinfo" {
			// Look for EnaSupport in the NetworkInfo JSON-like string
			// Example: {DefaultNetworkCardIndex:0,EfaInfo:null,EfaSupported:false,EnaSupport:unsupported,...}

			// Find EnaSupport field
			enaPattern := `EnaSupport:([^,}]+)`
			re := regexp.MustCompile(enaPattern)
			matches := re.FindStringSubmatch(value)

			if len(matches) >= 2 {
				enaValue := strings.TrimSpace(matches[1])
				log.Trace().Msgf("Found ENA support in NetworkInfo of %s: EnaSupport = %s", spec.CspSpecName, enaValue)
				return isEnaSupport(enaValue)
			}
		}
	}

	log.Trace().Msgf("No ENA support information found in spec %s", spec.CspSpecName)
	return "unknown"
}

// isEnaSupport parses ENA support value and returns normalized string
func isEnaSupport(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))

	switch value {
	case "required":
		return "required" // Strict requirement for ENA
	case "true", "supported", "enabled":
		return "supported" // ENA is supported but not required
	case "false", "unsupported", "disabled":
		return "unsupported"
	default:
		log.Trace().Msgf("Unknown ENA support value: %s", value)
		return "unknown"
	}
}

// === 4. NVMe Driver Support ===

// extractAwsNvmeSupportFromSpecDetails extracts NVMe support from spec details
func extractAwsNvmeSupportFromSpecDetails(spec cloudmodel.SpecInfo) string {
	for _, kv := range spec.Details {
		key := strings.ToLower(kv.Key)
		if strings.Contains(key, "nvmesupport") || strings.Contains(key, "ebsinfo") {
			value := strings.ToLower(strings.TrimSpace(kv.Value))
			if strings.Contains(value, "required") {
				return "required"
			}
			if strings.Contains(value, "supported") {
				return "supported"
			}
			if strings.Contains(value, "unsupported") {
				return "unsupported"
			}
		}
	}
	return "unknown"
}

// extractAwsNvmeSupportFromImageDetails extracts NVMe support from image details
func extractAwsNvmeSupportFromImageDetails(image cloudmodel.ImageInfo) string {
	for _, kv := range image.Details {
		key := strings.ToLower(kv.Key)
		if strings.Contains(key, "nvmesupport") {
			value := strings.ToLower(strings.TrimSpace(kv.Value))
			if value == "true" || value == "supported" {
				return "supported"
			}
			if value == "false" || value == "unsupported" {
				return "unsupported"
			}
		}
	}
	// For Xen-based AMIs, assume NVMe is not supported unless explicitly stated
	hypervisor := extractAwsHypervisorFromImageDetails(image)
	if normalizeAwsHypervisor(hypervisor) == "xen" {
		return "unsupported"
	}
	return "unknown"
}

// === 5. Boot Mode Compatibility ===

// isAwsBootModeCompatible checks boot mode compatibility between spec and image
func isAwsBootModeCompatible(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	supportedBootModes := extractAwsSupportedBootModesFromSpecDetails(spec)
	imageBootMode := extractAwsBootModeFromImageDetails(image)

	if len(supportedBootModes) == 0 || imageBootMode == "" {
		log.Trace().Msgf("AWS boot mode check - insufficient info, assuming compatible")
		return true
	}

	// Handle uefi-preferred: compatible with both uefi and legacy-bios specs
	if imageBootMode == "uefi-preferred" {
		for _, mode := range supportedBootModes {
			if mode == "uefi" || mode == "legacy-bios" {
				log.Trace().Msgf("AWS boot mode compatible - Image: %s, Spec supports: %v",
					imageBootMode, supportedBootModes)
				return true
			}
		}
	} else {
		// Direct mode matching
		for _, mode := range supportedBootModes {
			if mode == imageBootMode {
				log.Trace().Msgf("AWS boot mode compatible - Image: %s, Spec supports: %v",
					imageBootMode, supportedBootModes)
				return true
			}
		}
	}

	log.Trace().Msgf("AWS boot mode incompatible - Image: %s, Spec supports: %v",
		imageBootMode, supportedBootModes)
	return false
}

// extractAwsSupportedBootModesFromSpecDetails extracts supported boot modes from spec details
func extractAwsSupportedBootModesFromSpecDetails(spec cloudmodel.SpecInfo) []string {
	for _, kv := range spec.Details {
		if strings.Contains(strings.ToLower(kv.Key), "supportedbootmodes") {
			// Parse boot modes (e.g., "legacy-bios; uefi")
			modes := strings.Split(kv.Value, ";")
			var result []string
			for _, mode := range modes {
				cleaned := strings.ToLower(strings.TrimSpace(mode))
				if cleaned != "" {
					result = append(result, cleaned)
				}
			}
			return result
		}
	}
	return []string{}
}

// extractAwsBootModeFromImageDetails extracts boot mode from image details
func extractAwsBootModeFromImageDetails(image cloudmodel.ImageInfo) string {
	for _, kv := range image.Details {
		if strings.EqualFold(kv.Key, "bootmode") {
			return strings.ToLower(strings.TrimSpace(kv.Value))
		}
	}
	return ""
}

// === 6. AWS Xen-on-Nitro Compatibility ===
// Note: Simplified approach - modern Nitro instances generally support Xen AMIs
// with proper driver compatibility checks
