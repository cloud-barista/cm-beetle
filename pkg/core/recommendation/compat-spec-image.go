package recommendation

import (
	"fmt"
	"strings"

	"github.com/cloud-barista/cm-beetle/pkg/compat"
	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckSpecImageCompatibility validates compatibility between a single VM spec and VM image
func CheckSpecImageCompatibility(csp string, spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	cspLower := strings.ToLower(csp)

	log.Debug().Msgf("Checking compatibility between spec %s and image %s for CSP %s",
		spec.CspSpecName, image.CspImageName, csp)

	// Use the centralized compatibility check from compat package
	isCompatible := compat.CheckCompatibility(cspLower, spec, image)

	if isCompatible {
		log.Debug().Msgf("Spec %s and Image %s are compatible for CSP %s",
			spec.CspSpecName, image.CspImageName, csp)
	} else {
		log.Debug().Msgf("Spec %s and Image %s are incompatible for CSP %s",
			spec.CspSpecName, image.CspImageName, csp)
	}

	return isCompatible
}

// FindCompatibleSpecAndImage finds a compatible VM spec and image pair by performing CSP-specific compatibility checks
func FindCompatibleSpecAndImage(specs []cloudmodel.SpecInfo, images []cloudmodel.ImageInfo, csp string) (cloudmodel.SpecInfo, cloudmodel.ImageInfo, error) {
	var emptySpec cloudmodel.SpecInfo
	var emptyImage cloudmodel.ImageInfo

	if len(specs) == 0 {
		return emptySpec, emptyImage, fmt.Errorf("no VM specs provided")
	}
	if len(images) == 0 {
		return emptySpec, emptyImage, fmt.Errorf("no VM images provided")
	}

	log.Debug().Msgf("Finding compatible spec and image for CSP: %s, specs: %d, images: %d", csp, len(specs), len(images))

	// Pre-filter specs and images based on CSP-specific rules
	filteredSpecs, filteredImages := preFilterByCsp(csp, specs, images)

	if len(filteredSpecs) == 0 {
		return emptySpec, emptyImage, fmt.Errorf("no compatible VM specs found after CSP-specific filtering")
	}
	if len(filteredImages) == 0 {
		return emptySpec, emptyImage, fmt.Errorf("no compatible VM images found after CSP-specific filtering")
	}

	log.Debug().Msgf("After pre-filtering - specs: %d, images: %d", len(filteredSpecs), len(filteredImages))

	// Find best compatible pair without scoring
	bestSpec, bestImage, err := findCompatiblePair(csp, filteredSpecs, filteredImages)
	if err != nil {
		return emptySpec, emptyImage, fmt.Errorf("failed to find compatible spec-image pair: %w", err)
	}

	log.Info().Msgf("Found compatible pair - Spec: %s, Image: %s", bestSpec.CspSpecName, bestImage.CspImageName)
	return bestSpec, bestImage, nil
}

// FindCompatibleVmSpecAndImagePairs finds all compatible VM spec and image pairs by performing CSP-specific compatibility checks
func FindCompatibleVmSpecAndImagePairs(specs []cloudmodel.SpecInfo, images []cloudmodel.ImageInfo, csp string) ([]CompatibleSpecImagePair, error) {

	if len(specs) == 0 {
		return nil, fmt.Errorf("no VM specs provided")
	}
	if len(images) == 0 {
		return nil, fmt.Errorf("no VM images provided")
	}

	log.Debug().Msgf("Finding compatible spec and image for CSP: %s, specs: %d, images: %d", csp, len(specs), len(images))

	// Pre-filter specs and images based on CSP-specific rules
	filteredSpecs, filteredImages := preFilterByCsp(csp, specs, images)

	if len(filteredSpecs) == 0 {
		return nil, fmt.Errorf("no compatible VM specs found after CSP-specific filtering")
	}
	if len(filteredImages) == 0 {
		return nil, fmt.Errorf("no compatible VM images found after CSP-specific filtering")
	}

	log.Debug().Msgf("After pre-filtering - specs: %d, images: %d", len(filteredSpecs), len(filteredImages))

	// Find best compatible pair without scoring
	compatiblePairs, err := findCompatiblePairs(csp, filteredSpecs, filteredImages)
	if err != nil {
		return nil, fmt.Errorf("failed to find compatible spec-image pair: %w", err)
	}

	log.Info().Msgf("Found compatible pairs - Count: %d", len(compatiblePairs))
	// For now, just return the first compatible pair
	if len(compatiblePairs) > 0 {
		return compatiblePairs, nil
	}
	return nil, fmt.Errorf("no compatible spec-image pairs found")
}

// preFilterByCsp performs CSP-specific pre-filtering with integrated logic
func preFilterByCsp(csp string, specs []cloudmodel.SpecInfo, images []cloudmodel.ImageInfo) ([]cloudmodel.SpecInfo, []cloudmodel.ImageInfo) {
	cspLower := strings.ToLower(csp)

	switch cspLower {
	case "aws":
		// Filter out UEFI images for AWS
		filteredImages := make([]cloudmodel.ImageInfo, 0, len(images))
		for _, img := range images {
			if !strings.Contains(strings.ToLower(img.CspImageName), "uefi") {
				filteredImages = append(filteredImages, img)
			}
		}
		log.Debug().Msgf("AWS pre-filtering: %d images filtered to %d", len(images), len(filteredImages))
		return specs, filteredImages

	case "ncp":
		// Filter specs for KVM hypervisor only
		filteredSpecs := filterNcpVmSpecsByHypervisor(specs)
		log.Debug().Msgf("NCP pre-filtering: %d specs filtered to %d KVM specs", len(specs), len(filteredSpecs))
		return filteredSpecs, images

	default:
		// No specific filtering for GCP and others
		// Rely on comprehensive compatibility checks in findCompatiblePair
		log.Debug().Msgf("No specific pre-filtering rules for CSP: %s", csp)
		return specs, images
	}
}

// findCompatiblePair finds the first compatible spec-image pair using comprehensive compatibility checks
func findCompatiblePair(csp string, specs []cloudmodel.SpecInfo, images []cloudmodel.ImageInfo) (cloudmodel.SpecInfo, cloudmodel.ImageInfo, error) {
	var emptySpec cloudmodel.SpecInfo
	var emptyImage cloudmodel.ImageInfo

	cspLower := strings.ToLower(csp)

	// Use standard compatibility check for all CSPs
	for _, image := range images {
		for _, spec := range specs {
			if isCompatible := compat.CheckCompatibility(cspLower, spec, image); isCompatible {
				log.Info().Msgf("Found compatible pair - Spec: %s, Image: %s",
					spec.CspSpecName, image.CspImageName)
				log.Debug().Msgf("Compatible pair (spec): %v", spec)
				log.Debug().Msgf("Compatible pair (image): %v", image)
				return spec, image, nil
			}
		}
	}

	return emptySpec, emptyImage, fmt.Errorf("no compatible spec-image pairs found")
}

// findCompatiblePairs finds all compatible spec-image pairs using comprehensive compatibility checks
func findCompatiblePairs(csp string, specs []cloudmodel.SpecInfo, images []cloudmodel.ImageInfo) ([]CompatibleSpecImagePair, error) {

	var compatiblePairs []CompatibleSpecImagePair

	cspLower := strings.ToLower(csp)

	// Use standard compatibility check for all CSPs
	for _, image := range images {
		for _, spec := range specs {
			if isCompatible := compat.CheckCompatibility(cspLower, spec, image); isCompatible {
				compatiblePairs = append(compatiblePairs, CompatibleSpecImagePair{
					Spec:  spec,
					Image: image,
				})
				log.Debug().Msgf("Compatible pair - Spec: %s, Image: %s",
					spec.CspSpecName, image.CspImageName)
			} else {
				log.Debug().Msgf("Incompatible pair - Spec: %s, Image: %s",
					spec.CspSpecName, image.CspImageName)
			}
		}
	}
	if len(compatiblePairs) == 0 {
		return nil, fmt.Errorf("no compatible spec-image pairs found")
	}

	return compatiblePairs, nil
}

// filterNcpVmSpecsByHypervisor filters NCP VM specs to include only KVM hypervisor specs
func filterNcpVmSpecsByHypervisor(vmSpecs []cloudmodel.SpecInfo) []cloudmodel.SpecInfo {
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
