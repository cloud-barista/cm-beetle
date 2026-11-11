package migration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func withMockTumblebug(t *testing.T, handler http.HandlerFunc) {
	t.Helper()
	srv := httptest.NewServer(handler)
	restore := tbclient.SetClient(tbclient.NewClient(tbclient.ApiConfig{
		RestUrl:       srv.URL,
		Username:      "default",
		Password:      "default",
		HttpTransport: http.DefaultTransport,
	}))
	t.Cleanup(func() {
		restore()
		srv.Close()
	})
}

func TestCreateObjectStorage_IndependentBucketsAndSupportLookup(t *testing.T) {
	t.Run("Creates all buckets, normalizes uppercase CSP, and isolates versioning and CORS sessions", func(t *testing.T) {
		var mu sync.Mutex
		var supportProviders []string
		var versioningBodies []string
		var corsBodies []string

		withMockTumblebug(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/connConfig/aws-ap-northeast-2":
				_ = json.NewEncoder(w).Encode(tbmodel.ConnConfig{ConfigName: "aws-ap-northeast-2"})
			case r.Method == http.MethodPut && r.URL.Path == "/ns/mig01/resources/objectStorage":
				var req tbmodel.ObjectStorageCreateRequest
				require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
				_ = json.NewEncoder(w).Encode(tbmodel.ObjectStorageInfo{
					Id:              req.BucketName,
					Name:            req.BucketName,
					CspResourceName: "uid-" + req.BucketName,
					ConnectionName:  req.ConnectionName,
					Status:          "Available",
				})
			case r.Method == http.MethodGet && r.URL.Path == "/objectStorage/support":
				mu.Lock()
				supportProviders = append(supportProviders, r.URL.Query().Get("providerName"))
				mu.Unlock()
				_ = json.NewEncoder(w).Encode(tbmodel.ObjectStorageSupportResponse{
					Supports: map[string]tbmodel.ObjectStorageFeatureSupport{
						"aws": {Versioning: true, Cors: true},
					},
				})
			case r.Method == http.MethodPut && strings.HasSuffix(r.URL.Path, "/versioning"):
				var raw map[string]any
				require.NoError(t, json.NewDecoder(r.Body).Decode(&raw))
				b, _ := json.Marshal(raw)
				mu.Lock()
				versioningBodies = append(versioningBodies, string(b))
				mu.Unlock()
				w.WriteHeader(http.StatusOK)
			case r.Method == http.MethodPut && strings.HasSuffix(r.URL.Path, "/cors"):
				var raw map[string]any
				require.NoError(t, json.NewDecoder(r.Body).Decode(&raw))
				b, _ := json.Marshal(raw)
				mu.Lock()
				corsBodies = append(corsBodies, string(b))
				mu.Unlock()
				w.WriteHeader(http.StatusOK)
			default:
				w.WriteHeader(http.StatusNotFound)
			}
		})

		req := storagemodel.RecommendedObjectStorage{
			Status:      "recommended",
			TargetCloud: storagemodel.CloudProperty{Csp: "AWS", Region: "ap-northeast-2"},
			TargetObjectStorages: []storagemodel.TargetObjectStorage{
				{
					SourceBucketName: "src-1",
					BucketName:       "os-01",
					BucketSpecProperty: storagemodel.BucketSpecProperty{
						VersioningEnabled: true,
						CORSEnabled:       true,
						CORSRule: []storagemodel.CORSRule{
							{AllowedOrigin: []string{"*"}, AllowedMethod: []string{"GET"}},
						},
					},
				},
				{
					SourceBucketName: "src-2",
					BucketName:       "os-02",
					BucketSpecProperty: storagemodel.BucketSpecProperty{
						VersioningEnabled: true,
					},
				},
			},
		}

		res, err := migration.CreateObjectStorage("mig01", req, "dev")
		require.NoError(t, err)
		assert.Equal(t, "created", res.Status)
		assert.Empty(t, res.Warnings)
		require.Len(t, res.ObjectStorage, 2)
		assert.Equal(t, "dev-os-01", res.ObjectStorage[0].Id)
		assert.Equal(t, "dev-os-02", res.ObjectStorage[1].Id)
		assert.Equal(t, []string{"aws"}, supportProviders)
		assert.Len(t, versioningBodies, 2)
		assert.Len(t, corsBodies, 1)
	})

	t.Run("Returns partial status when one bucket collides and target CSP lacks CORS support", func(t *testing.T) {
		withMockTumblebug(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/connConfig/ncp-kr":
				_ = json.NewEncoder(w).Encode(tbmodel.ConnConfig{ConfigName: "ncp-kr"})
			case r.Method == http.MethodPut && r.URL.Path == "/ns/mig01/resources/objectStorage":
				var req tbmodel.ObjectStorageCreateRequest
				require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
				if req.BucketName == "os-01" {
					w.WriteHeader(http.StatusConflict)
					_, _ = w.Write([]byte(`{"message":"already exists"}`))
					return
				}
				_ = json.NewEncoder(w).Encode(tbmodel.ObjectStorageInfo{
					Id:   req.BucketName,
					Name: req.BucketName,
				})
			case r.Method == http.MethodGet && r.URL.Path == "/objectStorage/support":
				_ = json.NewEncoder(w).Encode(tbmodel.ObjectStorageSupportResponse{
					Supports: map[string]tbmodel.ObjectStorageFeatureSupport{
						"ncp": {Versioning: true, Cors: false},
					},
				})
			default:
				w.WriteHeader(http.StatusNotFound)
			}
		})

		req := storagemodel.RecommendedObjectStorage{
			TargetCloud: storagemodel.CloudProperty{Csp: "NCP", Region: "kr"},
			TargetObjectStorages: []storagemodel.TargetObjectStorage{
				{SourceBucketName: "src-1", BucketName: "os-01"},
				{
					SourceBucketName: "src-2",
					BucketName:       "os-02",
					BucketSpecProperty: storagemodel.BucketSpecProperty{
						CORSEnabled: true,
						CORSRule: []storagemodel.CORSRule{
							{AllowedOrigin: []string{"*"}, AllowedMethod: []string{"GET"}},
						},
					},
				},
			},
		}

		res, err := migration.CreateObjectStorage("mig01", req, "")
		require.NoError(t, err)
		assert.Equal(t, "partial", res.Status)
		require.Len(t, res.ObjectStorage, 1)
		assert.Equal(t, "os-02", res.ObjectStorage[0].Id)
		assert.Len(t, res.Warnings, 2)
	})

	t.Run("Returns error preserving 409 status when all buckets fail", func(t *testing.T) {
		withMockTumblebug(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/connConfig/aws-ap-northeast-2":
				_ = json.NewEncoder(w).Encode(tbmodel.ConnConfig{ConfigName: "aws-ap-northeast-2"})
			case r.Method == http.MethodPut && r.URL.Path == "/ns/mig01/resources/objectStorage":
				w.WriteHeader(http.StatusConflict)
				_, _ = w.Write([]byte(`{"message":"ObjectStorage os-01 already exists"}`))
			default:
				w.WriteHeader(http.StatusNotFound)
			}
		})

		req := storagemodel.RecommendedObjectStorage{
			TargetCloud: storagemodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"},
			TargetObjectStorages: []storagemodel.TargetObjectStorage{
				{SourceBucketName: "src-1", BucketName: "os-01"},
			},
		}

		res, err := migration.CreateObjectStorage("mig01", req, "")
		require.Error(t, err)
		assert.Equal(t, "failed", res.Status)
		assert.Contains(t, err.Error(), "409")
	})
}

func TestDeleteObjectStorage_NotFoundSemantics(t *testing.T) {
	t.Run("Standard and force delete treat 404 as idempotent success", func(t *testing.T) {
		withMockTumblebug(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message":"not found"}`))
		})

		require.NoError(t, migration.DeleteObjectStorage("mig01", "os-01", ""))
		require.NoError(t, migration.DeleteObjectStorage("mig01", "os-01", "force"))
	})

	t.Run("Empty option propagates 404 error when bucket does not exist", func(t *testing.T) {
		withMockTumblebug(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message":"object storage not found"}`))
		})

		err := migration.DeleteObjectStorage("mig01", "os-01", "empty")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "404")
	})

	t.Run("500 error on bucket containing 404 in its ID is not swallowed", func(t *testing.T) {
		withMockTumblebug(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"message":"internal failure on bucket os-404"}`))
		})

		err := migration.DeleteObjectStorage("mig01", "os-404", "")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "500")
	})
}
