package recommendation_test

import (
	"testing"
	"time"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
)

func init() {
	tbclient.Init(tbclient.ApiConfig{
		RestUrl:  "http://127.0.0.1:1323/tumblebug",
		Username: "default",
		Password: "default",
		Timeout:  100 * time.Millisecond,
	})
}

func TestRecommendRDBMS_BasicMySQL(t *testing.T) {
	sources := []rdbmsmodel.SourceRDBMSProperty{
		{
			InstanceName:  "src-mysql-01",
			Engine:        "mysql",
			EngineVersion: "8.0",
			Vcpu:          4,
			MemoryMb:      16384,
			StorageSizeGb: 150,
			StorageType:   "SSD",
			PublicAccess:  true,
			Databases: []rdbmsmodel.SourceDatabaseProperty{
				{DatabaseName: "order_db", CharacterSet: "utf8mb4"},
			},
		},
	}

	result, err := recommendation.RecommendRDBMS("aws", "ap-northeast-2", sources)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Status != "recommended" {
		t.Errorf("expected status 'recommended', got '%s'", result.Status)
	}

	if len(result.TargetRDBMSInstances) != 1 {
		t.Fatalf("expected 1 target instance, got %d", len(result.TargetRDBMSInstances))
	}

	target := result.TargetRDBMSInstances[0]
	if target.DBEngine != "mysql" {
		t.Errorf("expected engine 'mysql', got '%s'", target.DBEngine)
	}
	if target.StorageSize < 100 {
		t.Errorf("expected storage size >= 100 for AWS, got %d", target.StorageSize)
	}
	if len(target.Databases) != 1 || target.Databases[0].DatabaseName != "order_db" {
		t.Errorf("expected database 'order_db' to be mapped, got %v", target.Databases)
	}
}

func TestRecommendRDBMS_MariaDBFallback(t *testing.T) {
	sources := []rdbmsmodel.SourceRDBMSProperty{
		{
			InstanceName:  "src-maria-01",
			Engine:        "mariadb",
			EngineVersion: "10.5",
			Vcpu:          2,
			MemoryMb:      4096,
			StorageSizeGb: 10,
		},
	}

	// GCP does not support MariaDB
	result, err := recommendation.RecommendRDBMS("gcp", "asia-northeast3", sources)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	target := result.TargetRDBMSInstances[0]
	if target.DBEngine != "mysql" {
		t.Errorf("expected engine fallback to 'mysql', got '%s'", target.DBEngine)
	}

	// Storage size adjustment check (10GB -> min 20GB)
	if target.StorageSize < 20 {
		t.Errorf("expected storage size adjusted to >= 20GB, got %d", target.StorageSize)
	}
}
