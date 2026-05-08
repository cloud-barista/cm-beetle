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

// StorageObjectInfo represents a single object in an object storage bucket.
type StorageObjectInfo struct {
	Key          string `json:"key"`                    // Object key (relative path within the bucket)
	Size         int64  `json:"size"`                   // Size in bytes
	LastModified string `json:"lastModified,omitempty"` // Last modified timestamp (RFC 3339)
	ETag         string `json:"eTag,omitempty"`         // Entity tag (content hash)
	StorageClass string `json:"storageClass,omitempty"` // Storage class (e.g., STANDARD)
}

// StorageObjectListResponse is the list response for objects in an object storage bucket.
type StorageObjectListResponse struct {
	OsId    string              `json:"osId"`    // Object storage ID
	Count   int                 `json:"count"`   // Total number of objects
	Objects []StorageObjectInfo `json:"objects"` // List of objects
}

// StorageObjectMetadata represents the metadata of a single object returned by a HEAD request.
type StorageObjectMetadata struct {
	Key          string `json:"key"`                    // Object key
	Size         int64  `json:"size"`                   // Size in bytes
	LastModified string `json:"lastModified,omitempty"` // Last modified timestamp
	ETag         string `json:"eTag,omitempty"`         // Entity tag (content hash)
	StorageClass string `json:"storageClass,omitempty"` // Storage class
}
