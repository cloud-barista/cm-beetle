package migration_test

import (
	"testing"
	"time"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
)

func init() {
	tbclient.Init(tbclient.ApiConfig{
		RestUrl:  "http://127.0.0.1:1323/tumblebug",
		Username: "default",
		Password: "default",
		Timeout:  100 * time.Millisecond,
	})
}

func TestMigrateRDBMS_Validation(t *testing.T) {
	t.Run("Rejects missing adminUserPassword without fallback", func(t *testing.T) {
		req := rdbmsmodel.RecommendedRDBMS{
			Status:      "recommended",
			TargetCloud: rdbmsmodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"},
			TargetRDBMSInstances: []rdbmsmodel.TargetRDBMSInstance{
				{
					SourceInstanceName: "src-01",
					RDBMSName:          "mig-rdbms-01",
					DBEngine:           "mysql",
					DBEngineVersion:    "8.0",
					DBInstanceSpec:     "db.t3.medium",
					StorageType:        "gp3",
					StorageSize:        100,
					AdminUserName:      "dbadmin",
				},
			},
		}

		err := migration.CreateRDBMS("mig01", req, "test")
		if err == nil {
			t.Fatal("expected error when adminUserPassword is empty, got nil")
		}
	})

	t.Run("Attempts provisioning when adminUserName and adminUserPassword are provided", func(t *testing.T) {
		req := rdbmsmodel.RecommendedRDBMS{
			Status:      "recommended",
			TargetCloud: rdbmsmodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"},
			TargetRDBMSInstances: []rdbmsmodel.TargetRDBMSInstance{
				{
					SourceInstanceName: "src-01",
					RDBMSName:          "mig-rdbms-01",
					DBEngine:           "mysql",
					DBEngineVersion:    "8.0",
					DBInstanceSpec:     "db.t3.medium",
					StorageType:        "gp3",
					StorageSize:        100,
					AdminUserName:      "dbadmin",
					AdminUserPassword:  "Password123!",
					Databases: []rdbmsmodel.TargetDatabase{
						{
							DatabaseName: "sampledb",
							Charset:      "utf8mb4",
							Collation:    "utf8mb4_unicode_ci",
						},
					},
				},
			},
		}

		// Should attempt connection to tumblebug; fails gracefully on offline TB
		err := migration.CreateRDBMS("mig01", req, "test")
		if err == nil {
			t.Log("CreateRDBMS succeeded (live TB)")
		} else {
			t.Logf("CreateRDBMS failed as expected on offline TB: %v", err)
		}
	})
}
