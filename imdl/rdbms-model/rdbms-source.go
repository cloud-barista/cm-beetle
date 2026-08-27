package rdbmsmodel

// SourceRDBMS describes the source RDBMS environment/group input for recommendation and model persistence.
type SourceRDBMS struct {
	Description          string                `json:"description,omitempty"`
	SourceCloud          *CloudProperty        `json:"sourceCloud,omitempty"`
	SourceRDBMSInstances []SourceRDBMSProperty `json:"sourceRDBMSInstances" validate:"required,min=1"`
}

// SourceRDBMSProperty describes one RDBMS instance observed in the source environment.
type SourceRDBMSProperty struct {
	// InstanceName is the identifier for the DB instance.
	InstanceName string `json:"instanceName" validate:"required" example:"prod-mysql-01"`
	// MachineId is the optional host machine identifier (e.g., node UUID) for infra traceability.
	MachineId string `json:"machineId,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`

	// Engine & Version
	Engine        string `json:"engine" validate:"required" example:"mysql"`       // "mysql" or "mariadb"
	EngineVersion string `json:"engineVersion" validate:"required" example:"8.0"` // e.g. "8.0", "10.5"

	// Compute & Memory
	Vcpu     int `json:"vcpu" validate:"required" example:"4"`       // Number of vCPU cores
	MemoryMb int `json:"memoryMb" validate:"required" example:"8192"` // Memory in MB

	// Storage
	StorageSizeGb int    `json:"storageSizeGb" validate:"required" example:"100"` // Disk size in GB
	StorageType   string `json:"storageType,omitempty" example:"SSD"`             // Storage type (e.g. SSD, HDD)
	Iops          int    `json:"iops,omitempty" example:"3000"`                   // Observed IOPS

	// Network & Topology
	Port             int  `json:"port,omitempty" example:"3306"`               // DB listening port (default: 3306)
	HighAvailability bool `json:"highAvailability,omitempty" example:"false"` // HA / Replication mode
	PublicAccess     bool `json:"publicAccess,omitempty" example:"false"`      // External public access enabled

	// Backup & Policy
	BackupRetentionDays int `json:"backupRetentionDays,omitempty" example:"7"` // Backup retention in days

	// Logical Databases (Inner DBs)
	Databases []SourceDatabaseProperty `json:"databases,omitempty"`
}

// SourceDatabaseProperty describes a logical database inside a source RDBMS instance.
type SourceDatabaseProperty struct {
	DatabaseName string  `json:"databaseName" validate:"required" example:"order_db"`
	CharacterSet string  `json:"characterSet,omitempty" example:"utf8mb4"`
	Collation    string  `json:"collation,omitempty" example:"utf8mb4_unicode_ci"`
	SizeMb       float64 `json:"sizeMb,omitempty" example:"512.5"`
	TableCount   int     `json:"tableCount,omitempty" example:"24"`
}
