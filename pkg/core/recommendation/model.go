package recommendation

import (
	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
)

// RecommendationStatus represents the status of a recommendation.
type RecommendationStatus string

const (
	NothingRecommended   RecommendationStatus = "none"
	PartiallyRecommended RecommendationStatus = "partial"
	FullyRecommended     RecommendationStatus = "ok"
)

type CompatibleSpecImagePair struct {
	Spec  cloudmodel.SpecInfo  `json:"spec"`
	Image cloudmodel.ImageInfo `json:"image"`
}

// Note: The shared models for infrastructure recommendations have been moved to
// imdl/cloud-model for sharing and use with other subsystems.


/*
 * Models for Container Infrastructure (i.e., an infrastructure for Kubernetes)
 */
type RecommendedInfraInfo struct {
	Status      string                  `json:"status"`
	Description string                  `json:"description"`
	TargetInfra tbmodel.InfraDynamicReq `json:"targetInfra"`
}
