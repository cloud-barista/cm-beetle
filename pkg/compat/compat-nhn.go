package compat

import (
	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	"github.com/rs/zerolog/log"
)

// CheckNhn checks compatibility between NHN Cloud VM spec and OS image
func CheckNhn(spec cloudmodel.SpecInfo, image cloudmodel.ImageInfo) bool {
	log.Debug().Msgf("Starting NHN Cloud compatibility check for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)

	// TODO: Add NHN Cloud-specific compatibility checks using Detail information
	log.Info().Msg("NHN Cloud compatibility validation is planned for future implementation")

	log.Debug().Msgf("NHN Cloud compatibility check passed for Spec: %s, Image: %s", spec.CspSpecName, image.CspImageName)
	return true
}

// === NHN Cloud-Specific Helper Functions ===
// TODO: Add NHN Cloud-specific compatibility validation functions as needed
// Examples:
// - Instance type family compatibility
// - GPU instance support validation
// - Local SSD vs network storage requirements
// - Load balancer compatibility
// - Auto scaling group requirements
