package storagemodel

// RecommendedObjectStorageModel is the top-level wrapper for external systems (e.g., cm-damselfly).
type RecommendedObjectStorageModel struct {
	RecommendedObjectStorageModel RecommendedObjectStorage `json:"recommendedObjectStorageModel" validate:"required"`
}

// SourceObjectStorageModel is the top-level wrapper for external systems (e.g., cm-damselfly).
type SourceObjectStorageModel struct {
	SourceObjectStorageModel SourceObjectStorage `json:"sourceObjectStorageModel" validate:"required"`
}

// RecommendedObjectStorage is the recommendation result and direct input to the migration API.
type RecommendedObjectStorage struct {
	Status               string                `json:"status"`             // e.g., "recommended", "partial", "failed"
	Description          string                `json:"description"`        // Human-readable summary
	Warnings             []string              `json:"warnings,omitempty"` // CSP feature-support warnings
	TargetCloud          CloudProperty         `json:"targetCloud"`
	TargetObjectStorages []TargetObjectStorage `json:"targetObjectStorages"`
}

// TargetObjectStorage is the bucket specification to create in the target cloud.
type TargetObjectStorage struct {
	SourceBucketName string `json:"sourceBucketName"` // Originating source bucket name
	BucketName       string `json:"bucketName"`       // Globally unique target bucket name

	BucketSpecProperty // Target configuration derived from source features
}

// SourceObjectStorage describes a bucket as observed in the source environment.
type SourceObjectStorage struct {
	BucketName string `json:"bucketName" validate:"required"` // Source bucket name

	BucketFeatureProperty // Feature configuration observed in the source
	BucketUsageProperty   // Operational metrics (recommendation only)
	BucketMetaProperty    // Metadata observed in the source
}

// CloudProperty identifies the target cloud provider and region.
type CloudProperty struct {
	Csp    string `json:"csp"    example:"aws"`            // Cloud service provider (e.g., aws, azure, gcp, ncp, alibaba)
	Region string `json:"region" example:"ap-northeast-2"` // Region identifier
}

