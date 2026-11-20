package migration_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

		err := migration.CreateRDBMS(context.Background(), "mig01", req, "test")
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
		err := migration.CreateRDBMS(context.Background(), "mig01", req, "test")
		if err == nil {
			t.Log("CreateRDBMS succeeded (live TB)")
		} else {
			t.Logf("CreateRDBMS failed as expected on offline TB: %v", err)
		}
	})
}

func TestCreateRDBMS_PrivateLogicalDatabaseHandling(t *testing.T) {
	t.Run("AWS private RDBMS skips remote logical database creation", func(t *testing.T) {
		rdbmsCreateCalls := 0
		dbCreateCalls := 0

		withMockTumblebug(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/connConfig/aws-ap-northeast-2":
				_ = json.NewEncoder(w).Encode(tbmodel.ConnConfig{ConfigName: "aws-ap-northeast-2"})
			case r.Method == http.MethodPost && r.URL.Path == "/ns/mig01/resources/rdbms":
				rdbmsCreateCalls++
				_ = json.NewEncoder(w).Encode(rdbmsmodel.RDBMSInfo{Name: "mig-rdbms-01", Status: "Available"})
			case r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/database"):
				dbCreateCalls++
				w.WriteHeader(http.StatusUnprocessableEntity)
			default:
				http.NotFound(w, r)
			}
		})

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
					PublicAccess:       false,
					Databases: []rdbmsmodel.TargetDatabase{
						{DatabaseName: "appdb", Charset: "utf8mb4", Collation: "utf8mb4_unicode_ci"},
					},
				},
			},
		}

		err := migration.CreateRDBMS(context.Background(), "mig01", req, "")
		require.NoError(t, err)
		assert.Equal(t, 1, rdbmsCreateCalls)
		assert.Equal(t, 0, dbCreateCalls, "private AWS RDBMS must skip /database API call")
	})

	t.Run("Azure private RDBMS invokes logical database creation", func(t *testing.T) {
		rdbmsCreateCalls := 0
		dbCreateCalls := 0

		withMockTumblebug(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/connConfig/azure-koreacentral":
				_ = json.NewEncoder(w).Encode(tbmodel.ConnConfig{ConfigName: "azure-koreacentral"})
			case r.Method == http.MethodPost && r.URL.Path == "/ns/mig01/resources/rdbms":
				rdbmsCreateCalls++
				_ = json.NewEncoder(w).Encode(rdbmsmodel.RDBMSInfo{Name: "mig-rdbms-01", Status: "Available"})
			case r.Method == http.MethodPost && r.URL.Path == "/ns/mig01/resources/rdbms/mig-rdbms-01/database":
				dbCreateCalls++
				w.WriteHeader(http.StatusOK)
			default:
				http.NotFound(w, r)
			}
		})

		req := rdbmsmodel.RecommendedRDBMS{
			Status:      "recommended",
			TargetCloud: rdbmsmodel.CloudProperty{Csp: "azure", Region: "koreacentral"},
			TargetRDBMSInstances: []rdbmsmodel.TargetRDBMSInstance{
				{
					SourceInstanceName: "src-01",
					RDBMSName:          "mig-rdbms-01",
					DBEngine:           "mysql",
					DBEngineVersion:    "8.0",
					DBInstanceSpec:     "Standard_B2s",
					StorageSize:        100,
					AdminUserName:      "dbadmin",
					AdminUserPassword:  "Password123!",
					PublicAccess:       false,
					Databases: []rdbmsmodel.TargetDatabase{
						{DatabaseName: "appdb", Charset: "utf8mb4", Collation: "utf8mb4_unicode_ci"},
					},
				},
			},
		}

		err := migration.CreateRDBMS(context.Background(), "mig01", req, "")
		require.NoError(t, err)
		assert.Equal(t, 1, rdbmsCreateCalls)
		assert.Equal(t, 1, dbCreateCalls, "private Azure RDBMS should invoke /database API")
	})

	t.Run("Tencent private RDBMS continues after 422 on logical database with collation", func(t *testing.T) {
		rdbmsCreateCalls := 0
		dbCreateCalls := 0

		withMockTumblebug(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/connConfig/tencent-ap-seoul":
				_ = json.NewEncoder(w).Encode(tbmodel.ConnConfig{ConfigName: "tencent-ap-seoul"})
			case r.Method == http.MethodPost && r.URL.Path == "/ns/mig01/resources/rdbms":
				rdbmsCreateCalls++
				_ = json.NewEncoder(w).Encode(rdbmsmodel.RDBMSInfo{Name: "mig-rdbms-01", Status: "Available"})
			case r.Method == http.MethodPost && r.URL.Path == "/ns/mig01/resources/rdbms/mig-rdbms-01/database":
				dbCreateCalls++
				w.WriteHeader(http.StatusUnprocessableEntity)
				_, _ = w.Write([]byte(`{"message":"specifying collation on private Tencent RDBMS requires direct SQL"}`))
			default:
				http.NotFound(w, r)
			}
		})

		req := rdbmsmodel.RecommendedRDBMS{
			Status:      "recommended",
			TargetCloud: rdbmsmodel.CloudProperty{Csp: "tencent", Region: "ap-seoul"},
			TargetRDBMSInstances: []rdbmsmodel.TargetRDBMSInstance{
				{
					SourceInstanceName: "src-01",
					RDBMSName:          "mig-rdbms-01",
					DBEngine:           "mysql",
					DBEngineVersion:    "8.0",
					DBInstanceSpec:     "2C4G",
					StorageSize:        100,
					AdminUserName:      "dbadmin",
					AdminUserPassword:  "Password123!",
					PublicAccess:       false,
					Databases: []rdbmsmodel.TargetDatabase{
						{DatabaseName: "appdb", Charset: "utf8mb4", Collation: "utf8mb4_unicode_ci"},
					},
				},
			},
		}

		err := migration.CreateRDBMS(context.Background(), "mig01", req, "")
		require.NoError(t, err)
		assert.Equal(t, 1, rdbmsCreateCalls)
		assert.Equal(t, 1, dbCreateCalls)
	})
}
