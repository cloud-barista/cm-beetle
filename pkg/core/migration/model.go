package migration

// MigratedObjectStorageInfo represents an object storage bucket created in the target cloud via CM-Beetle.
type MigratedObjectStorageInfo struct {
	Id             string `json:"id"`                     // Bucket ID (unique identifier within the namespace)
	Name           string `json:"name"`                   // Bucket name in the target cloud
	Status         string `json:"status"`                 // Current status (e.g., "Available")
	Description    string `json:"description"`            // Description
	ConnectionName string `json:"connectionName"`         // Connection identifier (format: "{csp}-{region}")
	CreationDate   string `json:"creationDate,omitempty"` // Bucket creation date (RFC 3339)
}

// MigratedObjectStorageListResponse is the list response for migrated object storages.
type MigratedObjectStorageListResponse struct {
	ObjectStorages []MigratedObjectStorageInfo `json:"objectStorages"`
}
