package storagemodel

// * To avoid external framework dependencies, the following structs are synchronized with CB-Tumblebug framework.

type ObjectStorageInfo struct {
	// ResourceType is the type of this resource
	ResourceType string `json:"resourceType" example:"ObjectStorage"`

	// Id is unique identifier for the object
	Id string `json:"id" example:"globally-unique-bucket-name-12345"`
	// Uid is universally unique identifier for the object, used for labelSelector
	Uid string `json:"uid,omitempty" example:"wef12awefadf1221edcf"`

	// CspResourceName is name assigned to the CSP resource. This name is internally used to handle the resource.
	CspResourceName string `json:"cspResourceName,omitempty" example:""`
	// CspResourceId is resource identifier managed by CSP
	CspResourceId string `json:"cspResourceId,omitempty" example:""`

	// Variables for management of Object Storage resource in CB-Tumblebug
	ConnectionName   string      `json:"connectionName"`
	ConnectionConfig ConnConfig  `json:"connectionConfig"`
	Description      string      `json:"description" example:"this object storage is managed by CB-Tumblebug"`
	Status           string      `json:"status"`
	SystemMessage    string      `json:"systemMessage,omitempty"`
	Conditions       []Condition `json:"conditions,omitempty"`

	// Name is human-readable string to represent the object
	Name         string   `json:"name" example:"globally-unique-bucket-name-12345"`
	Prefix       string   `json:"prefix,omitempty" example:""`
	Marker       string   `json:"marker,omitempty" example:""`
	MaxKeys      int      `json:"maxKeys,omitempty" example:"1000"`
	IsTruncated  bool     `json:"isTruncated,omitempty" example:"false"`
	CreationDate string   `json:"creationDate,omitempty" example:"2025-09-04T04:18:06Z"`
	Contents     []Object `json:"contents,omitempty"`
}

// ObjectStorageListResponse represents the response structure for listing object storages
type ObjectStorageListResponse struct {
	ObjectStorage []ObjectStorageInfo `json:"objectStorage"`
}

// IdList represents a list of resource IDs.
type IdList struct {
	IdList []string `json:"idList"`
}

type Object struct {
	Key          string `json:"key" example:"test-object.txt"`
	LastModified string `json:"lastModified" example:"2025-09-04T04:18:06Z"`
	ETag         string `json:"eTag" example:"9b2cf535f27731c974343645a3985328"`
	Size         int64  `json:"size" example:"1024"`
	StorageClass string `json:"storageClass" example:"STANDARD"`
}

// ConditionType represents the type of a condition
type ConditionType string

// ConditionStatus represents the status of a condition
type ConditionStatus string

// Condition represents an observation about a resource's state
type Condition struct {
	Type               ConditionType   `json:"type"`
	Status             ConditionStatus `json:"status"`
	Reason             string          `json:"reason,omitempty"`
	Message            string          `json:"message,omitempty"`
	LastTransitionTime string          `json:"lastTransitionTime,omitempty"`
}

// Location is structure for location information
type Location struct {
	Display   string  `mapstructure:"display" json:"display"`
	Latitude  float64 `mapstructure:"latitude" json:"latitude"`
	Longitude float64 `mapstructure:"longitude" json:"longitude"`
}

// RegionZoneInfo is struct for containing region struct
type RegionZoneInfo struct {
	AssignedRegion string `json:"assignedRegion"`
	AssignedZone   string `json:"assignedZone"`
}

// RegionDetail is structure for region information
type RegionDetail struct {
	RegionId           string   `mapstructure:"id" json:"regionId"`
	RegionName         string   `mapstructure:"regionName" json:"regionName"`
	Description        string   `mapstructure:"description" json:"description"`
	Location           Location `mapstructure:"location" json:"location"`
	Zones              []string `mapstructure:"zone" json:"zones"`
	RepresentativeZone *string  `mapstructure:"representativeZone" json:"representativeZone,omitempty"`
}

// ConnConfig is struct for containing modified CB-Spider struct for connection config
type ConnConfig struct {
	ConfigName           string         `json:"configName"`
	ProviderName         string         `json:"providerName"`
	DriverName           string         `json:"driverName"`
	CredentialName       string         `json:"credentialName"`
	CredentialHolder     string         `json:"credentialHolder"`
	RegionZoneInfoName   string         `json:"regionZoneInfoName"`
	RegionZoneInfo       RegionZoneInfo `json:"regionZoneInfo"`
	RegionDetail         RegionDetail   `json:"regionDetail"`
	RegionRepresentative bool           `json:"regionRepresentative"`
	Verified             bool           `json:"verified"`
}
