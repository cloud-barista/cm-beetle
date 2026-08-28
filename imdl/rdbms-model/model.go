package rdbmsmodel

// RecommendedRDBMSModel is the top-level wrapper for external systems.
type RecommendedRDBMSModel struct {
	RecommendedRDBMSModel RecommendedRDBMS `json:"recommendedRDBMSModel" validate:"required"`
}

// SourceRDBMSModel is the top-level wrapper for external systems.
type SourceRDBMSModel struct {
	SourceRDBMSModel SourceRDBMS `json:"sourceRDBMSModel" validate:"required"`
}

// CloudProperty identifies the target cloud provider and region.
type CloudProperty struct {
	Csp    string `json:"csp"    example:"aws"`                     // Cloud service provider (e.g., aws, azure, gcp, ncp, alibaba, etc.)
	Region string `json:"region" example:"ap-northeast-2"`          // Region identifier
	Zone   string `json:"zone,omitempty" example:"ap-northeast-2a"` // Optional Zone identifier
}
