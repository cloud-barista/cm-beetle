package migration

import storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"

// MigratedObjectStorageInfo represents an object storage bucket created in the target cloud via CM-Beetle.
type MigratedObjectStorageInfo = storagemodel.ObjectStorageInfo

// MigratedObjectStorageListResponse is the list response for migrated object storages.
type MigratedObjectStorageListResponse = storagemodel.ObjectStorageListResponse

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
