package recommendation

import (
	"strings"
	"testing"
	"time"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
)

func init() {
	tbclient.Init(tbclient.ApiConfig{
		RestUrl:  "http://127.0.0.1:1323/tumblebug",
		Username: "default",
		Password: "default",
		Timeout:  100 * time.Millisecond,
	})
}

func TestValidateSourceRDBMS(t *testing.T) {
	tests := []struct {
		name        string
		sources     []rdbmsmodel.SourceRDBMSProperty
		expectErr   bool
		errContains string
	}{
		{
			name:        "Empty sources list",
			sources:     []rdbmsmodel.SourceRDBMSProperty{},
			expectErr:   true,
			errContains: "at least one source RDBMS instance is required",
		},
		{
			name: "Missing Engine",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					DisplayName: "db-01",
					DBNode: rdbmsmodel.DBNodeProperty{
						Hostname: "node-01",
						CPU:      rdbmsmodel.CpuProperty{Threads: 2},
						Memory:   rdbmsmodel.MemoryProperty{TotalSize: 4},
						RootDisk: rdbmsmodel.DiskProperty{TotalSize: 50},
					},
					DBEngine: rdbmsmodel.DBEngineProperty{
						Engine:        "",
						EngineVersion: "8.0",
					},
				},
			},
			expectErr:   true,
			errContains: "engine is required",
		},
		{
			name: "Missing EngineVersion",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					DisplayName: "db-01",
					DBNode: rdbmsmodel.DBNodeProperty{
						Hostname: "node-01",
						CPU:      rdbmsmodel.CpuProperty{Threads: 2},
						Memory:   rdbmsmodel.MemoryProperty{TotalSize: 4},
						RootDisk: rdbmsmodel.DiskProperty{TotalSize: 50},
					},
					DBEngine: rdbmsmodel.DBEngineProperty{
						Engine:        "mysql",
						EngineVersion: "  ",
					},
				},
			},
			expectErr:   true,
			errContains: "engineVersion is required",
		},
		{
			name: "Non-positive vCPU",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					DisplayName: "db-01",
					DBNode: rdbmsmodel.DBNodeProperty{
						Hostname: "node-01",
						CPU:      rdbmsmodel.CpuProperty{Threads: 0, Cpus: 0},
						Memory:   rdbmsmodel.MemoryProperty{TotalSize: 4},
						RootDisk: rdbmsmodel.DiskProperty{TotalSize: 50},
					},
					DBEngine: rdbmsmodel.DBEngineProperty{
						Engine:        "mysql",
						EngineVersion: "8.0",
					},
				},
			},
			expectErr:   true,
			errContains: "effective vcpu must be greater than 0",
		},
		{
			name: "Non-positive Memory",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					DisplayName: "db-01",
					DBNode: rdbmsmodel.DBNodeProperty{
						Hostname: "node-01",
						CPU:      rdbmsmodel.CpuProperty{Threads: 2},
						Memory:   rdbmsmodel.MemoryProperty{TotalSize: 0},
						RootDisk: rdbmsmodel.DiskProperty{TotalSize: 50},
					},
					DBEngine: rdbmsmodel.DBEngineProperty{
						Engine:        "mysql",
						EngineVersion: "8.0",
					},
				},
			},
			expectErr:   true,
			errContains: "memory totalSize must be greater than 0",
		},
		{
			name: "Non-positive Storage",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					DisplayName: "db-01",
					DBNode: rdbmsmodel.DBNodeProperty{
						Hostname: "node-01",
						CPU:      rdbmsmodel.CpuProperty{Threads: 2},
						Memory:   rdbmsmodel.MemoryProperty{TotalSize: 4},
						RootDisk: rdbmsmodel.DiskProperty{TotalSize: 0},
					},
					DBEngine: rdbmsmodel.DBEngineProperty{
						Engine:        "mysql",
						EngineVersion: "8.0",
					},
				},
			},
			expectErr:   true,
			errContains: "storage totalSize must be greater than 0",
		},
		{
			name: "Missing DatabaseName in inner database list",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					DisplayName: "db-01",
					DBNode: rdbmsmodel.DBNodeProperty{
						Hostname: "node-01",
						CPU:      rdbmsmodel.CpuProperty{Threads: 2},
						Memory:   rdbmsmodel.MemoryProperty{TotalSize: 4},
						RootDisk: rdbmsmodel.DiskProperty{TotalSize: 50},
					},
					DBEngine: rdbmsmodel.DBEngineProperty{
						Engine:        "mysql",
						EngineVersion: "8.0",
					},
					InnerDatabases: []rdbmsmodel.InnerDatabaseProperty{
						{DatabaseName: ""},
					},
				},
			},
			expectErr:   true,
			errContains: "databaseName is required",
		},
		{
			name: "Valid Source Instance",
			sources: []rdbmsmodel.SourceRDBMSProperty{
				{
					DisplayName: "db-01",
					DBNode: rdbmsmodel.DBNodeProperty{
						Hostname: "node-01",
						CPU:      rdbmsmodel.CpuProperty{Threads: 4},
						Memory:   rdbmsmodel.MemoryProperty{TotalSize: 8},
						RootDisk: rdbmsmodel.DiskProperty{TotalSize: 50, Type: "SSD"},
						DataDisks: []rdbmsmodel.DiskProperty{
							{TotalSize: 50, Type: "SSD"},
						},
					},
					DBEngine: rdbmsmodel.DBEngineProperty{
						Engine:        "mysql",
						EngineVersion: "8.0",
						Role:          "primary",
					},
					InnerDatabases: []rdbmsmodel.InnerDatabaseProperty{
						{DatabaseName: "app_db"},
					},
				},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSourceRDBMS(tt.sources)
			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error containing '%s', got nil", tt.errContains)
				}
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Fatalf("expected error containing '%s', got '%s'", tt.errContains, err.Error())
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestRecommendRDBMS_ValidationFailures(t *testing.T) {
	validSource := []rdbmsmodel.SourceRDBMSProperty{
		{
			DisplayName: "src-01",
			DBNode: rdbmsmodel.DBNodeProperty{
				Hostname: "node-01",
				CPU:      rdbmsmodel.CpuProperty{Threads: 2},
				Memory:   rdbmsmodel.MemoryProperty{TotalSize: 4},
				RootDisk: rdbmsmodel.DiskProperty{TotalSize: 50, Type: "SSD"},
			},
			DBEngine: rdbmsmodel.DBEngineProperty{
				Engine:        "mysql",
				EngineVersion: "8.0",
			},
		},
	}

	// Missing CSP
	_, err := RecommendRDBMS("", "ap-northeast-2", validSource, true)
	if err == nil || !strings.Contains(err.Error(), "desiredCsp and desiredRegion are required") {
		t.Errorf("expected error for missing CSP, got %v", err)
	}

	// Missing Region
	_, err = RecommendRDBMS("aws", "", validSource, true)
	if err == nil || !strings.Contains(err.Error(), "desiredCsp and desiredRegion are required") {
		t.Errorf("expected error for missing Region, got %v", err)
	}

	// Invalid Source (empty sources)
	_, err = RecommendRDBMS("aws", "ap-northeast-2", []rdbmsmodel.SourceRDBMSProperty{}, true)
	if err == nil || !strings.Contains(err.Error(), "at least one source RDBMS instance is required") {
		t.Errorf("expected error for empty sources, got %v", err)
	}
}

func TestAutoFillSourceRDBMSDefaults(t *testing.T) {
	t.Run("Empty sources list returns error", func(t *testing.T) {
		_, _, err := AutoFillSourceRDBMSDefaults([]rdbmsmodel.SourceRDBMSProperty{})
		if err == nil || !strings.Contains(err.Error(), "at least one source RDBMS instance is required") {
			t.Errorf("expected error for empty sources, got %v", err)
		}
	})

	t.Run("Missing engine returns error", func(t *testing.T) {
		sources := []rdbmsmodel.SourceRDBMSProperty{
			{
				DisplayName: "test-db",
				DBEngine:    rdbmsmodel.DBEngineProperty{Engine: ""},
			},
		}
		_, _, err := AutoFillSourceRDBMSDefaults(sources)
		if err == nil || !strings.Contains(err.Error(), "engine is required") {
			t.Errorf("expected error for missing engine, got %v", err)
		}
	})

	t.Run("Applies all sensible defaults when compute, storage, and version are 0/empty", func(t *testing.T) {
		sources := []rdbmsmodel.SourceRDBMSProperty{
			{
				DisplayName: "zero-spec-mysql",
				DBEngine: rdbmsmodel.DBEngineProperty{
					Engine: "mysql",
				},
			},
		}

		filled, warnings, err := AutoFillSourceRDBMSDefaults(sources)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(filled) != 1 {
			t.Fatalf("expected 1 filled source, got %d", len(filled))
		}

		inst := filled[0]
		// EngineVersion defaulted to DefaultMySQLVersion ("8.0")
		if inst.DBEngine.EngineVersion != DefaultMySQLVersion {
			t.Errorf("expected engineVersion '%s', got '%s'", DefaultMySQLVersion, inst.DBEngine.EngineVersion)
		}
		// Port defaulted to 3306
		if inst.DBEngine.Port != 3306 {
			t.Errorf("expected port 3306, got %d", inst.DBEngine.Port)
		}
		// Role defaulted to standalone
		if inst.DBEngine.Role != "standalone" {
			t.Errorf("expected role 'standalone', got '%s'", inst.DBEngine.Role)
		}
		// vCPU defaulted to DefaultRDBMSVcpu (2)
		if inst.DBNode.CPU.Cores != DefaultRDBMSVcpu || inst.DBNode.CPU.Threads != DefaultRDBMSVcpu {
			t.Errorf("expected vCPU cores/threads %d, got cores=%d threads=%d", DefaultRDBMSVcpu, inst.DBNode.CPU.Cores, inst.DBNode.CPU.Threads)
		}
		// Memory defaulted to DefaultRDBMSMemoryGiB (4 GiB)
		if inst.DBNode.Memory.TotalSize != DefaultRDBMSMemoryGiB {
			t.Errorf("expected memory %d GiB, got %d", DefaultRDBMSMemoryGiB, inst.DBNode.Memory.TotalSize)
		}
		// Storage defaulted to DefaultRDBMSStorageGB (100) and DefaultRDBMSStorageType ("SSD")
		if inst.DBNode.RootDisk.TotalSize != DefaultRDBMSStorageGB || inst.DBNode.RootDisk.Type != DefaultRDBMSStorageType {
			t.Errorf("expected storage %d GB %s, got %d GB %s", DefaultRDBMSStorageGB, DefaultRDBMSStorageType, inst.DBNode.RootDisk.TotalSize, inst.DBNode.RootDisk.Type)
		}

		// Verify actionable warnings were generated
		if len(warnings) < 4 {
			t.Errorf("expected at least 4 warnings (version, CPU, memory, storage), got %d: %v", len(warnings), warnings)
		}
	})

	t.Run("Engine version defaults for MariaDB and PostgreSQL", func(t *testing.T) {
		sources := []rdbmsmodel.SourceRDBMSProperty{
			{
				DisplayName: "mariadb-node",
				DBEngine:    rdbmsmodel.DBEngineProperty{Engine: "mariadb"},
			},
			{
				DisplayName: "pg-node",
				DBEngine:    rdbmsmodel.DBEngineProperty{Engine: "postgresql"},
			},
		}

		filled, _, err := AutoFillSourceRDBMSDefaults(sources)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if filled[0].DBEngine.EngineVersion != DefaultMariaDBVersion {
			t.Errorf("expected MariaDB default version '%s', got '%s'", DefaultMariaDBVersion, filled[0].DBEngine.EngineVersion)
		}
		if filled[0].DBEngine.Port != 3306 {
			t.Errorf("expected MariaDB default port 3306, got %d", filled[0].DBEngine.Port)
		}

		if filled[1].DBEngine.EngineVersion != DefaultPostgreSQLVersion {
			t.Errorf("expected PostgreSQL default version '%s', got '%s'", DefaultPostgreSQLVersion, filled[1].DBEngine.EngineVersion)
		}
		if filled[1].DBEngine.Port != 5432 {
			t.Errorf("expected PostgreSQL default port 5432, got %d", filled[1].DBEngine.Port)
		}
	})

	t.Run("Preserves explicitly specified specifications without overwriting", func(t *testing.T) {
		sources := []rdbmsmodel.SourceRDBMSProperty{
			{
				DisplayName: "custom-db",
				DBEngine: rdbmsmodel.DBEngineProperty{
					Engine:        "mysql",
					EngineVersion: "5.7",
					Port:          3307,
					Role:          "primary",
				},
				DBNode: rdbmsmodel.DBNodeProperty{
					CPU:      rdbmsmodel.CpuProperty{Threads: 8},
					Memory:   rdbmsmodel.MemoryProperty{TotalSize: 16},
					RootDisk: rdbmsmodel.DiskProperty{TotalSize: 200, Type: "HDD"},
				},
			},
		}

		filled, warnings, err := AutoFillSourceRDBMSDefaults(sources)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(warnings) != 0 {
			t.Errorf("expected 0 warnings for fully specified source, got %d: %v", len(warnings), warnings)
		}
		if filled[0].DBEngine.EngineVersion != "5.7" {
			t.Errorf("expected version '5.7', got '%s'", filled[0].DBEngine.EngineVersion)
		}
		if filled[0].DBEngine.Port != 3307 {
			t.Errorf("expected port 3307, got %d", filled[0].DBEngine.Port)
		}
		if filled[0].DBNode.CPU.Threads != 8 {
			t.Errorf("expected threads 8, got %d", filled[0].DBNode.CPU.Threads)
		}
		if filled[0].DBNode.Memory.TotalSize != 16 {
			t.Errorf("expected memory 16, got %d", filled[0].DBNode.Memory.TotalSize)
		}
		if filled[0].DBNode.RootDisk.TotalSize != 200 || filled[0].DBNode.RootDisk.Type != "HDD" {
			t.Errorf("expected rootDisk 200 HDD, got %d %s", filled[0].DBNode.RootDisk.TotalSize, filled[0].DBNode.RootDisk.Type)
		}
	})
}

func TestRecommendRDBMS_AutoFillVsStrictValidationFlag(t *testing.T) {
	// Source with 0 vCPU, 0 memory, 0 storage
	zeroSpecSource := []rdbmsmodel.SourceRDBMSProperty{
		{
			DisplayName: "zero-spec-db",
			DBEngine: rdbmsmodel.DBEngineProperty{
				Engine: "mysql",
			},
		},
	}

	// 1. When autoFillSourceDefaults = false, strict ValidateSourceRDBMS must fail immediately
	_, err := RecommendRDBMS("aws", "ap-northeast-2", zeroSpecSource, false)
	if err == nil {
		t.Fatalf("expected validation error when autoFillSourceDefaults is false, got nil")
	}
	if !strings.Contains(err.Error(), "engineVersion is required") &&
		!strings.Contains(err.Error(), "effective vcpu must be greater than 0") {
		t.Errorf("expected engineVersion or vcpu validation error, got: %v", err)
	}

	// 2. When autoFillSourceDefaults = true, validation does not fail with "vcpu must be greater than 0"
	// (it proceeds past source validation to Tumblebug network call or capability lookup)
	_, err = RecommendRDBMS("aws", "ap-northeast-2", zeroSpecSource, true)
	// If it fails, it must NOT be a SourceRDBMS validation failure (such as "vcpu must be greater than 0")
	if err != nil && (strings.Contains(err.Error(), "effective vcpu must be greater than 0") ||
		strings.Contains(err.Error(), "engineVersion is required")) {
		t.Errorf("unexpected validation failure when autoFillSourceDefaults is true: %v", err)
	}
}

func TestTargetPreferences_DefaultsAndOverrides(t *testing.T) {
	falseVal := false
	trueVal := true

	pref := &TargetPreferences{
		AdminUserName:            "dbadmin",
		HighAvailability:         &falseVal,
		PublicAccess:             &trueVal,
		BackupRetentionDays:      14,
		NHNDBSGToAllowAllInbound: true,
	}

	if pref.AdminUserName != "dbadmin" {
		t.Errorf("expected AdminUserName 'dbadmin', got '%s'", pref.AdminUserName)
	}
	if pref.HighAvailability == nil || *pref.HighAvailability != false {
		t.Errorf("expected HighAvailability false, got %v", pref.HighAvailability)
	}
	if pref.PublicAccess == nil || *pref.PublicAccess != true {
		t.Errorf("expected PublicAccess true, got %v", pref.PublicAccess)
	}
	if pref.BackupRetentionDays != 14 {
		t.Errorf("expected BackupRetentionDays 14, got %d", pref.BackupRetentionDays)
	}
	if !pref.NHNDBSGToAllowAllInbound {
		t.Errorf("expected NHNDBSGToAllowAllInbound true, got %v", pref.NHNDBSGToAllowAllInbound)
	}

	// Verify default preferences zero value
	defaultPref := &TargetPreferences{}
	if defaultPref.BackupRetentionDays != 0 {
		t.Errorf("expected default BackupRetentionDays 0, got %d", defaultPref.BackupRetentionDays)
	}
}

func TestCalculationHelpers(t *testing.T) {
	// 1. resolveSourceRDBMSInstanceName
	t.Run("resolveSourceRDBMSInstanceName", func(t *testing.T) {
		// DisplayName preferred
		src1 := rdbmsmodel.SourceRDBMSProperty{
			DisplayName: "my-custom-db",
			DBNode:      rdbmsmodel.DBNodeProperty{Hostname: "host-1", MachineId: "id-1"},
		}
		if id := resolveSourceRDBMSInstanceName(src1); id != "my-custom-db" {
			t.Errorf("expected 'my-custom-db', got '%s'", id)
		}

		// Hostname fallback
		src2 := rdbmsmodel.SourceRDBMSProperty{
			DBNode: rdbmsmodel.DBNodeProperty{Hostname: "host-1", MachineId: "id-1"},
		}
		if id := resolveSourceRDBMSInstanceName(src2); id != "host-1" {
			t.Errorf("expected 'host-1', got '%s'", id)
		}

		// MachineId fallback
		src3 := rdbmsmodel.SourceRDBMSProperty{
			DBNode: rdbmsmodel.DBNodeProperty{MachineId: "id-1"},
		}
		if id := resolveSourceRDBMSInstanceName(src3); id != "id-1" {
			t.Errorf("expected 'id-1', got '%s'", id)
		}

		// Default fallback
		src4 := rdbmsmodel.SourceRDBMSProperty{}
		if id := resolveSourceRDBMSInstanceName(src4); id != "rdbms-node" {
			t.Errorf("expected 'rdbms-node', got '%s'", id)
		}
	})

	// 2. calculateEffectiveVcpu
	t.Run("calculateEffectiveVcpu", func(t *testing.T) {
		nodeWithThreads := rdbmsmodel.DBNodeProperty{
			CPU: rdbmsmodel.CpuProperty{Threads: 8, Cpus: 1, Cores: 4},
		}
		if v := calculateEffectiveVcpu(nodeWithThreads); v != 8 {
			t.Errorf("expected 8 threads, got %d", v)
		}

		nodeWithCpusCores := rdbmsmodel.DBNodeProperty{
			CPU: rdbmsmodel.CpuProperty{Cpus: 2, Cores: 4},
		}
		if v := calculateEffectiveVcpu(nodeWithCpusCores); v != 8 {
			t.Errorf("expected 8 (2*4), got %d", v)
		}

		nodeEmpty := rdbmsmodel.DBNodeProperty{}
		if v := calculateEffectiveVcpu(nodeEmpty); v != 0 {
			t.Errorf("expected 0 for empty CPU, got %d", v)
		}
	})

	// 3. calculateEffectiveMemoryMb
	t.Run("calculateEffectiveMemoryMb", func(t *testing.T) {
		node := rdbmsmodel.DBNodeProperty{
			Memory: rdbmsmodel.MemoryProperty{TotalSize: 8},
		}
		if m := calculateEffectiveMemoryMb(node); m != 8192 {
			t.Errorf("expected 8192 MB, got %d", m)
		}

		nodeZero := rdbmsmodel.DBNodeProperty{}
		if m := calculateEffectiveMemoryMb(nodeZero); m != 0 {
			t.Errorf("expected 0 for empty memory, got %d", m)
		}
	})

	// 4. calculateEffectiveStorageSizeGb
	t.Run("calculateEffectiveStorageSizeGb", func(t *testing.T) {
		node := rdbmsmodel.DBNodeProperty{
			RootDisk:  rdbmsmodel.DiskProperty{TotalSize: 50},
			DataDisks: []rdbmsmodel.DiskProperty{{TotalSize: 100}, {TotalSize: 50}},
		}
		if s := calculateEffectiveStorageSizeGb(node); s != 200 {
			t.Errorf("expected 200 GB, got %d", s)
		}
	})

	// 5. determinePrimaryStorageType
	t.Run("determinePrimaryStorageType", func(t *testing.T) {
		// DataDisk type preferred
		node1 := rdbmsmodel.DBNodeProperty{
			RootDisk:  rdbmsmodel.DiskProperty{Type: "HDD"},
			DataDisks: []rdbmsmodel.DiskProperty{{Type: "SSD"}},
		}
		if st := determinePrimaryStorageType(node1); st != "SSD" {
			t.Errorf("expected 'SSD', got '%s'", st)
		}

		// RootDisk fallback
		node2 := rdbmsmodel.DBNodeProperty{
			RootDisk: rdbmsmodel.DiskProperty{Type: "HDD"},
		}
		if st := determinePrimaryStorageType(node2); st != "HDD" {
			t.Errorf("expected 'HDD', got '%s'", st)
		}

		// Default fallback
		node3 := rdbmsmodel.DBNodeProperty{}
		if st := determinePrimaryStorageType(node3); st != "SSD" {
			t.Errorf("expected 'SSD', got '%s'", st)
		}
	})
}

func TestRecommendDBInstanceSpec_Proximity(t *testing.T) {
	capa := rdbmsmodel.RDBMSMetaInfo{
		DBInstanceSpecs: []rdbmsmodel.RDBMSDBInstanceSpecInfo{
			{Name: "db.t3.small", VCpuCount: "2", MemSizeMiB: "2048"},
			{Name: "db.t3.medium", VCpuCount: "2", MemSizeMiB: "4096"},
			{Name: "db.t3.large", VCpuCount: "2", MemSizeMiB: "8192"},
			{Name: "db.t3.xlarge", VCpuCount: "4", MemSizeMiB: "16384"},
			{Name: "db.m5.large", VCpuCount: "2", MemSizeMiB: "8192"},
			{Name: "db.c5.xlarge", VCpuCount: "4", MemSizeMiB: "8192"},
		},
	}

	// 1. Exact match (2 vCPU, 4096 MiB, 100 GB) -> should pick db.t3.medium
	spec, err := recommendDBInstanceSpec(2, 4096, 100, "mysql", capa)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec != "db.t3.medium" {
		t.Errorf("expected 'db.t3.medium', got '%s'", spec)
	}

	// 2. Compute-intensive workload (4 vCPU, 4096 MiB = ratio 1.0) -> should favor vCPU
	spec, err = recommendDBInstanceSpec(4, 4096, 100, "mysql", capa)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec != "db.c5.xlarge" {
		t.Errorf("expected 'db.c5.xlarge', got '%s'", spec)
	}

	// 3. Exact larger match (4 vCPU, 16384 MiB) -> should pick db.t3.xlarge
	spec, err = recommendDBInstanceSpec(4, 16384, 100, "mysql", capa)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec != "db.t3.xlarge" {
		t.Errorf("expected 'db.t3.xlarge', got '%s'", spec)
	}

	// 4. Empty specs list with no fallback should return error
	emptyCapa := rdbmsmodel.RDBMSMetaInfo{
		DBInstanceSpecs: []rdbmsmodel.RDBMSDBInstanceSpecInfo{},
	}
	_, err = recommendDBInstanceSpec(2, 4096, 100, "mysql", emptyCapa)
	if err == nil {
		t.Fatalf("expected error when DBInstanceSpecs is empty, got nil")
	}
}

func TestRecommendDBInstanceSpec_DiskAware(t *testing.T) {
	capa := rdbmsmodel.RDBMSMetaInfo{
		DBInstanceSpecs: []rdbmsmodel.RDBMSDBInstanceSpecInfo{
			{
				Name:               "db.t3.medium.limited_disk",
				VCpuCount:          "2",
				MemSizeMiB:         "4096",
				StorageSizeRangeGB: rdbmsmodel.StorageSizeRange{Min: 10, Max: 50},
			},
			{
				Name:               "db.t3.medium.standard_disk",
				VCpuCount:          "2",
				MemSizeMiB:         "4096",
				StorageSizeRangeGB: rdbmsmodel.StorageSizeRange{Min: 10, Max: 1000},
			},
			{
				Name:               "db.m5.large",
				VCpuCount:          "2",
				MemSizeMiB:         "8192",
				StorageSizeRangeGB: rdbmsmodel.StorageSizeRange{Min: 10, Max: 5000},
			},
		},
	}

	// 1. Source database requires 100 GB storage with 2 vCPU, 4096 MiB RAM.
	// db.t3.medium.limited_disk only supports up to 50 GB, so db.t3.medium.standard_disk should be selected.
	spec, err := recommendDBInstanceSpec(2, 4096, 100, "mysql", capa)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec != "db.t3.medium.standard_disk" {
		t.Errorf("expected 'db.t3.medium.standard_disk', got '%s'", spec)
	}

	// 2. Source database requires 2000 GB storage with 2 vCPU, 4096 MiB RAM.
	// Both db.t3.medium specs have Max <= 1000, so db.m5.large should be selected to satisfy disk.
	spec, err = recommendDBInstanceSpec(2, 4096, 2000, "mysql", capa)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec != "db.m5.large" {
		t.Errorf("expected 'db.m5.large', got '%s'", spec)
	}
}

func TestSelectEngineVersion_Proximity(t *testing.T) {
	supported := []string{"8.0.32", "8.0.35", "8.4.0"}
	warnings := make([]string, 0)

	// Exact match
	v, err := selectEngineVersion("mysql", "8.0.32", supported, &warnings, "test-inst")
	if err != nil || v != "8.0.32" {
		t.Errorf("expected '8.0.32', got '%s', err: %v", v, err)
	}

	// Prefix match
	v, err = selectEngineVersion("mysql", "8.4", supported, &warnings, "test-inst")
	if err != nil || v != "8.4.0" {
		t.Errorf("expected '8.4.0', got '%s', err: %v", v, err)
	}

	// IBM Scenario: Requested 8.0 on target cloud offering only [8.4.0] -> recommends closest version 8.4.0 with warning
	ibmSupported := []string{"8.4.0"}
	warnings = nil
	v, err = selectEngineVersion("mysql", "8.0", ibmSupported, &warnings, "test-inst")
	if err != nil || v != "8.4.0" {
		t.Errorf("expected closest version '8.4.0', got '%s', err: %v", v, err)
	}
	if len(warnings) == 0 || !strings.Contains(warnings[0], "recommended closest available version '8.4.0'") {
		t.Errorf("expected warning about recommending closest version, got: %v", warnings)
	}

	// OpenStack Scenario: Requested 8.0 on target cloud offering only [5.7.29] -> recommends 5.7.29 with warning
	openstackSupported := []string{"5.7.29"}
	warnings = nil
	v, err = selectEngineVersion("mysql", "8.0", openstackSupported, &warnings, "test-inst")
	if err != nil || v != "5.7.29" {
		t.Errorf("expected fallback version '5.7.29', got '%s', err: %v", v, err)
	}
	if len(warnings) == 0 || !strings.Contains(warnings[0], "recommended closest available version '5.7.29'") {
		t.Errorf("expected warning about recommending closest version, got: %v", warnings)
	}
}

func TestIsDBEngineSupported(t *testing.T) {
	support := rdbmsmodel.RDBMSCSPSupportInfo{
		Supported:          true,
		SupportedDBEngines: []string{"mysql", "mariadb"},
	}

	if !isDBEngineSupported(support, true, "mysql") {
		t.Errorf("expected mysql to be supported")
	}
	if !isDBEngineSupported(support, true, "mariadb") {
		t.Errorf("expected mariadb to be supported")
	}
	if isDBEngineSupported(support, true, "postgresql") {
		t.Errorf("expected postgresql to be unsupported")
	}

	unsupportedCSP := rdbmsmodel.RDBMSCSPSupportInfo{
		Supported:          true,
		SupportedDBEngines: []string{"mysql"},
	}
	if isDBEngineSupported(unsupportedCSP, true, "mariadb") {
		t.Errorf("expected mariadb to be unsupported when not in SupportedDBEngines")
	}
}

func TestSelectStorageType_NoFallback(t *testing.T) {
	capa := rdbmsmodel.RDBMSMetaInfo{
		SupportsStorageTypeSelection: true,
		StorageTypeOptions:           []string{"gp2", "gp3", "io1"},
		Notes: &rdbmsmodel.RDBMSNotes{
			StorageTypes: []rdbmsmodel.StorageTypeNote{
				{StorageType: "gp3", Recommended: true},
				{StorageType: "gp2", Recommended: false},
			},
		},
	}

	var warnings []string

	// 1. Exact match
	st, note, err := selectStorageType("gp2", capa, &warnings, "inst-01")
	if err != nil || st != "gp2" {
		t.Errorf("expected 'gp2', got '%s', err: %v", st, err)
	}

	// 2. Generic SSD mapping to recommended SSD (gp3)
	st, note, err = selectStorageType("SSD", capa, &warnings, "inst-01")
	if err != nil || st != "gp3" || note == nil || !note.Recommended {
		t.Errorf("expected 'gp3', got '%s', err: %v", st, err)
	}

	// 3. Unsupported specific storage type -> Recommends available with warning
	warnings = nil
	st, _, err = selectStorageType("non_existent_storage", capa, &warnings, "inst-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if st != "gp3" {
		t.Errorf("expected fallback to recommended 'gp3', got '%s'", st)
	}
	if len(warnings) == 0 || !strings.Contains(warnings[0], "not supported on target cloud") {
		t.Errorf("expected warning for unsupported storage type, got: %v", warnings)
	}

	// 4. Unselectable storage type on CSP (e.g. Azure/NCP) -> returns "" and warns if requested
	unselectableCapa := rdbmsmodel.RDBMSMetaInfo{
		ProviderName:                 "azure",
		SupportsStorageTypeSelection: false,
		StorageTypeOptions:           []string{"auto"},
	}
	st, _, err = selectStorageType("SSD", unselectableCapa, &warnings, "inst-01")
	if err != nil || st != "" {
		t.Errorf("expected empty string for unselectable storage, got '%s', err: %v", st, err)
	}

	// 5. Target cloud with non-SSD named storage (e.g. OpenStack [RBD, __DEFAULT__])
	// Generic "SSD" request should map to "RBD" with an informative warning
	openstackCapa := rdbmsmodel.RDBMSMetaInfo{
		ProviderName:                 "openstack",
		SupportsStorageTypeSelection: true,
		StorageTypeOptions:           []string{"__DEFAULT__", "RBD"},
	}
	st, _, err = selectStorageType("SSD", openstackCapa, &warnings, "inst-01")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if st != "RBD" {
		t.Errorf("expected 'RBD', got '%s'", st)
	}
	if len(warnings) == 0 || !strings.Contains(warnings[len(warnings)-1], "recommended closest available storage type 'RBD'") {
		t.Errorf("expected warning about mapping SSD to RBD, got warnings: %v", warnings)
	}

	// 6. Alibaba scenario: options contain [cloud_auto, cloud_essd, cloud_essd2, cloud_essd3, local_ssd]
	// cloud_auto is Recommended. Generic "SSD" request MUST pick "cloud_auto" and NOT "cloud_essd2" (premium)
	alibabaCapa := rdbmsmodel.RDBMSMetaInfo{
		ProviderName:                 "alibaba",
		SupportsStorageTypeSelection: true,
		StorageTypeOptions:           []string{"cloud_auto", "cloud_essd", "cloud_essd2", "cloud_essd3", "local_ssd"},
		Notes: &rdbmsmodel.RDBMSNotes{
			StorageTypes: []rdbmsmodel.StorageTypeNote{
				{StorageType: "cloud_auto", DisplayName: "Auto-selected Storage Type (SSD)", Recommended: true},
				{StorageType: "cloud_essd", DisplayName: "Enhanced SSD (ESSD PL1)", RecommendationLevel: "standard"},
				{StorageType: "cloud_essd2", DisplayName: "Enhanced SSD (ESSD PL2)", RecommendationLevel: "premium", MinSize: 500},
				{StorageType: "cloud_essd3", DisplayName: "Enhanced SSD (ESSD PL3)", RecommendationLevel: "premium", MinSize: 1500},
			},
		},
	}
	st, note, err = selectStorageType("SSD", alibabaCapa, &warnings, "inst-01")
	if err != nil {
		t.Fatalf("unexpected error for alibaba: %v", err)
	}
	if st != "cloud_auto" {
		t.Errorf("expected 'cloud_auto', got '%s'", st)
	}
	if note == nil || !note.Recommended {
		t.Errorf("expected note to be recommended")
	}
}

func TestResolveTargetAdminUser(t *testing.T) {
	// 1. Unspecified user preference -> defaults to DefaultRDBMSAdminUser ("dbadmin")
	t.Run("Default username when unspecified", func(t *testing.T) {
		var warnings []string
		u := resolveTargetAdminUser("", nil, &warnings)
		if u != DefaultRDBMSAdminUser {
			t.Errorf("expected '%s', got '%s'", DefaultRDBMSAdminUser, u)
		}
		if len(warnings) != 0 {
			t.Errorf("expected 0 warnings, got %d", len(warnings))
		}
	})

	// 2. User-specified allowed username -> preserves user input
	t.Run("Preserves valid user input", func(t *testing.T) {
		var warnings []string
		req := &rdbmsmodel.RDBMSAdminUserNameRequirement{
			ReservedValues: []string{"root", "admin", "administrator"},
		}
		u := resolveTargetAdminUser("customdbmaster", req, &warnings)
		if u != "customdbmaster" {
			t.Errorf("expected 'customdbmaster', got '%s'", u)
		}
		if len(warnings) != 0 {
			t.Errorf("expected 0 warnings, got %d", len(warnings))
		}
	})

	// 3. User requested a reserved username -> replaces with compliant default and warns
	t.Run("Replaces reserved username with default and warns", func(t *testing.T) {
		var warnings []string
		req := &rdbmsmodel.RDBMSAdminUserNameRequirement{
			ReservedValues: []string{"root", "admin"},
		}
		u := resolveTargetAdminUser("root", req, &warnings)
		if u != DefaultRDBMSAdminUser {
			t.Errorf("expected '%s', got '%s'", DefaultRDBMSAdminUser, u)
		}
		if len(warnings) == 0 || !strings.Contains(warnings[0], "not allowed on target cloud") {
			t.Errorf("expected not allowed warning, got %v", warnings)
		}
	})

	// 4. CSP requires FixedValue -> enforces FixedValue and warns if user specified different
	t.Run("Enforces CSP FixedValue", func(t *testing.T) {
		var warnings []string
		req := &rdbmsmodel.RDBMSAdminUserNameRequirement{
			FixedValue: "azureuser",
		}
		u := resolveTargetAdminUser("customuser", req, &warnings)
		if u != "azureuser" {
			t.Errorf("expected 'azureuser', got '%s'", u)
		}
		if len(warnings) == 0 || !strings.Contains(warnings[0], "not allowed on target cloud") {
			t.Errorf("expected not allowed warning, got %v", warnings)
		}
	})
}
