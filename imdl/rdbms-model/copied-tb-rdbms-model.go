package rdbmsmodel

// * To avoid circular/version dependencies, the following structs are copied from the cb-tumblebug framework.
// TODO: When the cb-tumblebug framework is updated, we should synchronize these structs.
// * Source: CB-Tumblebug (src/core/model/rdbms.go, src/core/model/common.go, src/core/model/net.go)
// * Synchronized: 2026-08-27 (RDBMSInfo, RDBMSListResponse, RDBMSCreateRequest, RDBMSCapabilityResponse, RDBMSSupportResponse, RDBMSDatabaseCreateReq, RDBMSDatabaseInfo, RDBMSDatabaseListResponse, etc.)

// StorageSizeRange defines the minimum and maximum storage capacity in Tumblebug format.
type StorageSizeRange struct {
	Min int `json:"min" example:"10"`
	Max int `json:"max" example:"1000"`
}

// StorageTypeNote provides user-facing guidance for a storage type.
type StorageTypeNote struct {
	StorageType         string            `json:"storageType" example:"gp3"`
	DisplayName         string            `json:"displayName" example:"General Purpose SSD v3"`
	Description         string            `json:"description" example:"Cost-effective, 3000 baseline IOPS, recommended for general workloads"`
	MinSize             int               `json:"minSize,omitempty" example:"100"`
	MaxSize             int               `json:"maxSize,omitempty" example:"65536"`
	RequiresIops        bool              `json:"requiresIops,omitempty" example:"true"`
	IopsRange           *StorageSizeRange `json:"iopsRange,omitempty"`
	Recommended         bool              `json:"recommended,omitempty" example:"true"`
	RecommendationLevel string            `json:"recommendationLevel,omitempty" example:"recommended" enums:"legacy,standard,recommended,premium"`
	CompatibleSpecs     []string          `json:"compatibleSpecs,omitempty" example:"rds.mysql.*"`
	IncompatibleSpecs   []string          `json:"incompatibleSpecs,omitempty" example:"mysql.n4.*"`
	Constraints         string            `json:"constraints,omitempty" example:"Requires 'iops' parameter (e.g., '3000')"`
}

// RDBMSAdminUserNameRequirement captures CSP-level admin username constraints.
type RDBMSAdminUserNameRequirement struct {
	FixedValue     string   `json:"fixedValue,omitempty"`
	ReservedValues []string `json:"reservedValues,omitempty"`
	Note           string   `json:"note,omitempty"`
}

// RDBMSAdminUserPasswordRequirement captures CSP-level admin password constraints.
type RDBMSAdminUserPasswordRequirement struct {
	MinLength           int    `json:"minLength,omitempty"`
	MaxLength           int    `json:"maxLength,omitempty"`
	RequiresSpecialChar bool   `json:"requiresSpecialChar,omitempty"`
	ForbidsSpecialChar  bool   `json:"forbidsSpecialChar,omitempty"`
	Note                string `json:"note,omitempty"`
}

// RDBMSStorageTypeConfig is one storage type entry under a CSP's "storageTypes" in assets/rdbmsinfo.yaml.
type RDBMSStorageTypeConfig struct {
	Name                    string            `yaml:"name"`
	Description             string            `yaml:"description"`
	RecommendationLevel     string            `yaml:"recommendationLevel"` // legacy|standard|recommended|premium
	RequiresIops            bool              `yaml:"requiresIops"`
	MinStorageSize          int               `yaml:"minStorageSize,omitempty"`
	MaxStorageSize          int               `yaml:"maxStorageSize,omitempty"`
	IopsRange               *StorageSizeRange `yaml:"iopsRange,omitempty"`
	CompatibleSpecs         []string          `yaml:"compatibleSpecs,omitempty"`
	IncompatibleSpecs       []string          `yaml:"incompatibleSpecs,omitempty"`
	CompatibleMachineSeries []string          `yaml:"compatibleMachineSeries,omitempty"`
	Note                    string            `yaml:"note,omitempty"`
}

// RDBMSDBMSRequirement is one DB-engine entry for requirements.
type RDBMSDBMSRequirement struct {
	MinStorageSize          int      `json:"minStorageSize,omitempty"`
	MaxStorageSize          int      `json:"maxStorageSize,omitempty"`
	DefaultPort             int      `json:"defaultPort,omitempty"`
	Note                    string   `json:"note,omitempty"`
	ReferenceEngineVersion  string   `json:"referenceEngineVersion,omitempty"`
	ReferenceDBInstanceSpec string   `json:"referenceDBInstanceSpec,omitempty"`
	// DeprecatedVersions lists engine versions that are deprecated by the CSP and discouraged for new deployments.
	DeprecatedVersions []string `json:"deprecatedVersions,omitempty"`
	// EndOfLifeVersions lists engine versions that have reached official End of Life (EOL) and cannot be provisioned.
	EndOfLifeVersions []string `json:"endOfLifeVersions,omitempty"`
}

// RDBMSDatabaseRequirement captures database name rules.
type RDBMSDatabaseRequirement struct {
	MaxDatabaseNameLength int      `json:"maxDatabaseNameLength,omitempty"`
	ReservedDatabaseNames []string `json:"reservedDatabaseNames,omitempty"`
	Note                  string   `json:"note,omitempty"`
}

// StaticFieldNote flags one RDBMSMetaInfo field whose value is fixed/approximate.
type StaticFieldNote struct {
	Field string `json:"field" example:"storageSizeRange"`
	Note  string `json:"note" example:"Static fallback value"`
}

// RDBMSDBInstanceSpecInfo represents individual instance specification in Tumblebug capability.
type RDBMSDBInstanceSpecInfo struct {
	Name               string           `json:"name" example:"db.t3.medium"`
	VCpuCount          string           `json:"vCpuCount,omitempty" example:"2"`
	VCpuClockGHz       string           `json:"vCpuClockGHz,omitempty" example:"2.5"`
	MemSizeMiB         string           `json:"memSizeMiB,omitempty" example:"4096"`
	StorageSizeRangeGB StorageSizeRange `json:"storageSizeRangeGB,omitempty"`
}

// RDBMSNotes holds structured guidance notes for storage types and instance specs.
type RDBMSNotes struct {
	StorageTypes []StorageTypeNote `json:"storageTypes,omitempty"`
}

// RDBMSMetaInfo represents Tumblebug-style RDBMS support capability details.
type RDBMSMetaInfo struct {
	ProviderName                     string                             `json:"providerName" example:"aws"`
	RegionName                       string                             `json:"regionName" example:"ap-northeast-2"`
	ConnectionName                   string                             `json:"connectionName" example:"aws-ap-northeast-2"`
	DBEngine                         string                             `json:"dbEngine" example:"mysql"`
	SupportedVersions                []string                           `json:"supportedVersions" example:"8.0,8.4"`
	DBInstanceSpecOptions            []string                           `json:"dbInstanceSpecOptions" example:"db.t3.medium"`
	DBInstanceSpecs                  []RDBMSDBInstanceSpecInfo          `json:"dbInstanceSpecs,omitempty"`
	LiveSupportedEngines             []string                           `json:"liveSupportedEngines,omitempty"`
	StorageTypeOptions               []string                           `json:"storageTypeOptions" example:"gp2,gp3"`
	DefaultStorageType               string                             `json:"defaultStorageType" example:"gp3"`
	StorageSizeRange                 StorageSizeRange                   `json:"storageSizeRange"`
	SupportsTag                      bool                               `json:"supportsTag" example:"true"`
	SupportsStorageTypeSelection     bool                               `json:"supportsStorageTypeSelection" example:"true"`
	SupportsStorageSizeConfiguration bool                               `json:"supportsStorageSizeConfiguration" example:"true"`
	SupportsHighAvailability         bool                               `json:"supportsHighAvailability" example:"true"`
	SupportsPublicAccess             bool                               `json:"supportsPublicAccess" example:"true"`
	SupportsBackup                   bool                               `json:"supportsBackup" example:"true"`
	SupportsEncryption               bool                               `json:"supportsEncryption" example:"true"`
	SupportsDeletionProtection       bool                               `json:"supportsDeletionProtection" example:"true"`
	RequiresSubnet                   bool                               `json:"requiresSubnet" example:"true"`
	RequiresSecurityGroup            bool                               `json:"requiresSecurityGroup" example:"true"`
	BackupRetentionRange             string                             `json:"backupRetentionRange,omitempty"`
	StorageTypeGuidance              map[string]StorageTypeNote         `json:"storageTypeGuidance,omitempty"`
	Notes                            *RDBMSNotes                        `json:"notes,omitempty"`
	AdminUserNameRequirement         *RDBMSAdminUserNameRequirement     `json:"adminUserNameRequirement,omitempty"`
	AdminUserPasswordRequirement     *RDBMSAdminUserPasswordRequirement `json:"adminUserPasswordRequirement,omitempty"`
	DBMSRequirements                 map[string]RDBMSDBMSRequirement    `json:"dbmsRequirements,omitempty"`
	DatabaseRequirements             *RDBMSDatabaseRequirement          `json:"databaseRequirements,omitempty"`
	DBOperationMethod                string                             `json:"dbOperationMethod" example:"cspNativeApi"`
	StaticFields                     []StaticFieldNote                  `json:"staticFields,omitempty"`
}

// RDBMSCapabilityResponse wraps the Tumblebug API response for GET /rdbms/capability.
type RDBMSCapabilityResponse struct {
	ResourceType string        `json:"resourceType" example:"rdbms"`
	Supports     RDBMSMetaInfo `json:"supports"`
}

// RDBMSSupportResponse wraps the Tumblebug API response for GET /rdbms/support.
type RDBMSSupportResponse struct {
	ResourceType string                         `json:"resourceType" example:"rdbms"`
	Supports     map[string]RDBMSCSPSupportInfo `json:"supports"`
}

// RDBMSCSPSupportInfo is one CSP's entry in GET /rdbms/support.
type RDBMSCSPSupportInfo struct {
	Supported             bool     `json:"supported" example:"true"`
	SupportedDBEngines    []string `json:"supportedDBEngines,omitempty" example:"mysql,mariadb"`
	DBOperationMethod     string   `json:"dbOperationMethod,omitempty" example:"cspNativeApi" enums:"cspNativeApi,sqlFallback"`
	SupportsTag           bool     `json:"supportsTag,omitempty" example:"true"`
	StorageTypeSelectable bool     `json:"storageTypeSelectable" example:"true"`
	Note                  string   `json:"note,omitempty" example:"Storage type selection is not supported on this CSP."`
}

// RDBMSCreateRequest is the Tumblebug-facing request to create an RDBMS instance.
type RDBMSCreateRequest struct {
	Name                     string     `json:"name" validate:"required" example:"rdbms-01"`
	ConnectionName           string     `json:"connectionName" validate:"required" example:"aws-ap-northeast-2"`
	VNetId                   string     `json:"vNetId" validate:"required" example:"vnet-01"`
	SubnetIds                []string   `json:"subnetIds,omitempty" example:"subnet-01"`
	SecurityGroupIds         []string   `json:"securityGroupIds,omitempty" example:"sg-01"`
	DBEngine                 string     `json:"dbEngine" validate:"required" example:"mysql" enums:"mysql,mariadb"`
	DBEngineVersion          string     `json:"dbEngineVersion,omitempty" example:"8.0"`
	DBInstanceSpec           string     `json:"dbInstanceSpec,omitempty" example:"db.t3.medium"`
	StorageType              string     `json:"storageType,omitempty" example:"gp3"`
	StorageSize              int        `json:"storageSize,omitempty" example:"100"`
	Iops                     string     `json:"iops,omitempty" example:"3000"`
	AdminUserName            string     `json:"adminUserName" validate:"required" example:"admin"`
	AdminUserPassword        string     `json:"adminUserPassword" validate:"required" example:"Password123!"`
	HighAvailability         bool       `json:"highAvailability,omitempty" example:"false"`
	BackupRetentionDays      int        `json:"backupRetentionDays,omitempty" example:"7"`
	PublicAccess             bool       `json:"publicAccess,omitempty" example:"false"`
	NHNDBSGToAllowAllInbound bool       `json:"nhnDBSGToAllowAllInbound,omitempty" example:"false"`
	DeletionProtection       bool       `json:"deletionProtection,omitempty" example:"false"`
	Description              string     `json:"description,omitempty" example:"managed by CB-Tumblebug"`
	AutoFillDefaults         bool       `json:"autoFillDefaults,omitempty" example:"false"`
	TagList                  []KeyValue `json:"tagList,omitempty"`
}

// RDBMSInfo is the Tumblebug-facing RDBMS instance resource (persisted and returned to callers).
type RDBMSInfo struct {
	ResourceType string `json:"resourceType" example:"rdbms"`

	Id              string `json:"id" example:"rdbms-01"`
	Uid             string `json:"uid,omitempty" example:"wef12awefadf1221edcf"`
	CspResourceName string `json:"cspResourceName,omitempty"`
	CspResourceId   string `json:"cspResourceId,omitempty"`

	Name             string      `json:"name" example:"rdbms-01"`
	ConnectionName   string      `json:"connectionName"`
	ConnectionConfig ConnConfig  `json:"connectionConfig"`
	Description      string      `json:"description,omitempty"`
	Status           string      `json:"status"`
	SystemMessage    string      `json:"systemMessage,omitempty"`
	Conditions       []Condition `json:"conditions,omitempty"`

	DeletionRequestedAt string `json:"deletionRequestedAt,omitempty"`

	VNetId           string   `json:"vNetId"`
	SubnetIds        []string `json:"subnetIds,omitempty"`
	SecurityGroupIds []string `json:"securityGroupIds,omitempty"`

	DBEngine                 string     `json:"dbEngine" example:"mysql"`
	DBEngineVersion          string     `json:"dbEngineVersion" example:"8.0"`
	DBInstanceSpec           string     `json:"dbInstanceSpec" example:"db.t3.medium"`
	DBInstanceType           string     `json:"dbInstanceType,omitempty" example:"Primary" enums:"Primary,ReadReplica"`
	StorageType              string     `json:"storageType,omitempty" example:"gp3"`
	StorageSize              int        `json:"storageSize" example:"100"`
	Iops                     string     `json:"iops,omitempty" example:"3000"`
	AdminUserName            string     `json:"adminUserName" example:"admin"`
	HighAvailability         bool       `json:"highAvailability" example:"false"`
	BackupRetentionDays      int        `json:"backupRetentionDays,omitempty" example:"7"`
	BackupTime               string     `json:"backupTime,omitempty" example:"03:00"`
	PublicAccess             bool       `json:"publicAccess" example:"false"`
	NHNDBSGToAllowAllInbound bool       `json:"nhnDBSGToAllowAllInbound,omitempty"`
	DeletionProtection       bool       `json:"deletionProtection" example:"false"`
	Encryption               bool       `json:"encryption,omitempty"`
	Endpoint                 string     `json:"endpoint,omitempty" example:"rdbms-01.xxxx.rds.amazonaws.com:3306"`
	TagList                  []KeyValue `json:"tagList,omitempty"`
}

// RDBMSListResponse is the response structure for listing RDBMS instances.
type RDBMSListResponse struct {
	RDBMS []RDBMSInfo `json:"rdbms"`
}

// RDBMSDatabaseCreateReq creates a logical database inside an Available RDBMS instance.
type RDBMSDatabaseCreateReq struct {
	DatabaseName      string `json:"databaseName" validate:"required" example:"sampledb"`
	AdminUserPassword string `json:"adminUserPassword" validate:"required" example:"Password123!"`
}

// RDBMSDatabaseInfo represents one logical database inside an RDBMS instance.
type RDBMSDatabaseInfo struct {
	DatabaseName string `json:"databaseName" example:"sampledb"`
}

// RDBMSDatabaseListResponse wraps the response for listing databases inside an RDBMS instance.
type RDBMSDatabaseListResponse struct {
	Databases []string `json:"databases" example:"sampledb"`
}

// IdList represents a list of resource IDs.
type IdList struct {
	IdList []string `json:"idList"`
}

// KeyValue represents a key-value pair for tags/metadata.
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Condition represents an observation about a resource's state.
type Condition struct {
	Type               ConditionType   `json:"type"`
	Status             ConditionStatus `json:"status"`
	Reason             string          `json:"reason,omitempty"`
	Message            string          `json:"message,omitempty"`
	LastTransitionTime string          `json:"lastTransitionTime,omitempty"`
}

// ConditionType represents the type of a condition.
type ConditionType string

// ConditionStatus represents the status of a condition.
type ConditionStatus string

// Location is structure for location information.
type Location struct {
	Display   string  `mapstructure:"display" json:"display"`
	Latitude  float64 `mapstructure:"latitude" json:"latitude"`
	Longitude float64 `mapstructure:"longitude" json:"longitude"`
}

// RegionZoneInfo is struct for containing region struct.
type RegionZoneInfo struct {
	AssignedRegion string `json:"assignedRegion"`
	AssignedZone   string `json:"assignedZone"`
}

// RegionDetail is structure for region information.
type RegionDetail struct {
	RegionId           string   `mapstructure:"id" json:"regionId"`
	RegionName         string   `mapstructure:"regionName" json:"regionName"`
	Description        string   `mapstructure:"description" json:"description"`
	Location           Location `mapstructure:"location" json:"location"`
	Zones              []string `mapstructure:"zone" json:"zones"`
	RepresentativeZone *string  `mapstructure:"representativeZone" json:"representativeZone,omitempty"`
}

// ConnConfig is struct for containing connection config.
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
