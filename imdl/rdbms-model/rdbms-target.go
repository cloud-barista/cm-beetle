package rdbmsmodel

// RecommendedRDBMS is the recommendation result and direct input to the migration API.
type RecommendedRDBMS struct {
	Status               string                `json:"status"`             // "recommended", "partial", "failed"
	Description          string                `json:"description"`        // Human-readable summary
	Warnings             []string              `json:"warnings,omitempty"` // CSP feature-support warnings
	TargetCloud          CloudProperty         `json:"targetCloud"`
	TargetRDBMSInstances []TargetRDBMSInstance `json:"targetRDBMSInstances"`
}

// TargetRDBMSInstance is the DB instance specification to create in the target cloud.
type TargetRDBMSInstance struct {
	// Source traceability
	SourceInstanceName string `json:"sourceInstanceName" example:"prod-mysql-01"`
	SourceMachineId    string `json:"sourceMachineId,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`

	// Target instance name (supports Late-Binding nameSeed)
	RDBMSName string `json:"rdbmsName" validate:"required" example:"mig-rdbms-01"`

	// Target Engine & Version
	DBEngine        string `json:"dbEngine" validate:"required" example:"mysql"` // "mysql" or "mariadb"
	DBEngineVersion string `json:"dbEngineVersion,omitempty" example:"8.0"`

	// Target Spec & Storage
	DBInstanceSpec string `json:"dbInstanceSpec,omitempty" example:"db.t3.medium"`
	StorageType    string `json:"storageType,omitempty" example:"gp3"`
	StorageSize    int    `json:"storageSize" example:"100"` // Disk size in GB
	Iops           string `json:"iops,omitempty" example:"3000"`

	// Admin Credentials
	AdminUserName     string `json:"adminUserName" validate:"required" example:"cbuser"`
	AdminUserPassword string `json:"adminUserPassword,omitempty" example:"Password123!"`

	// Network & Access
	VNetId           string   `json:"vNetId,omitempty" example:"vnet-01"`
	SubnetIds        []string `json:"subnetIds,omitempty" example:"[\"subnet-01\",\"subnet-02\"]"`
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" example:"[\"sg-01\"]"`
	PublicAccess     bool     `json:"publicAccess" example:"true"`

	// Availability & Protection
	HighAvailability    bool `json:"highAvailability" example:"false"`
	BackupRetentionDays int  `json:"backupRetentionDays,omitempty" example:"7"`
	DeletionProtection  bool `json:"deletionProtection,omitempty" example:"false"`

	// Optional Inner Databases to create after instance provisioning
	Databases []TargetDatabase `json:"databases,omitempty"`
}

// TargetDatabase represents a logical database to create inside the target RDBMS instance.
type TargetDatabase struct {
	DatabaseName string `json:"databaseName" validate:"required" example:"order_db"`
	CharacterSet string `json:"characterSet,omitempty" example:"utf8mb4"`
}
