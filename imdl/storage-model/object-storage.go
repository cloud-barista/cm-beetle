package storagemodel

// CORSRule represents a CORS rule. Field names use singular form to match CB-Spider/CB-Tumblebug convention.
type CORSRule struct {
	AllowedOrigin []string `json:"allowedOrigin" validate:"required"` // Allowed origins (e.g., ["*"], ["https://example.com"])
	AllowedMethod []string `json:"allowedMethod" validate:"required"` // Allowed HTTP methods (e.g., ["GET", "PUT", "POST"])
	AllowedHeader []string `json:"allowedHeader,omitempty"`           // Allowed headers (e.g., ["*"], ["Content-Type"])
	ExposeHeader  []string `json:"exposeHeader,omitempty"`            // Headers exposed to the browser (e.g., ["ETag"])
	MaxAgeSeconds int      `json:"maxAgeSeconds,omitempty"`           // Preflight response cache duration in seconds
}

// Target-side property types

// BucketSpecProperty is the bucket configuration to apply when creating a target bucket.
type BucketSpecProperty struct {
	VersioningEnabled bool       `json:"versioningEnabled"` // Whether to enable versioning
	CORSEnabled       bool       `json:"corsEnabled"`       // Whether to configure CORS
	CORSRule          []CORSRule `json:"corsRule,omitempty"` // CORS rules to apply
}

// Source-side property types

// BucketFeatureProperty captures the feature configuration of a source bucket.
type BucketFeatureProperty struct {
	VersioningEnabled bool       `json:"versioningEnabled,omitempty"` // Whether versioning is enabled
	CORSEnabled       bool       `json:"corsEnabled,omitempty"`       // Whether CORS is configured
	CORSRule          []CORSRule `json:"corsRule,omitempty"`          // Active CORS rules
	EncryptionEnabled bool       `json:"encryptionEnabled,omitempty"` // Whether server-side encryption is enabled
	IsPublic          bool       `json:"isPublic,omitempty"`          // Whether the bucket allows public access
}

// BucketUsageProperty captures operational metrics of a source bucket (used for recommendation only).
type BucketUsageProperty struct {
	TotalSizeBytes  int64  `json:"totalSizeBytes,omitempty"`  // Total stored data in bytes
	ObjectCount     int64  `json:"objectCount,omitempty"`     // Total number of objects
	AccessFrequency string `json:"accessFrequency,omitempty"` // Access pattern: "frequent" | "infrequent" | "archive"
}

// BucketMetaProperty captures metadata of a source bucket.
type BucketMetaProperty struct {
	Tags         map[string]string `json:"tags,omitempty"`         // Key-value tags
	CreationDate string            `json:"creationDate,omitempty"` // Bucket creation date (RFC 3339)
}
