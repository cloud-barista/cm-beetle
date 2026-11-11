package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockTumblebugForController(t *testing.T, handler http.HandlerFunc) {
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

func TestDeleteObjectStorageController_OptionAllowListAndStatuses(t *testing.T) {
	e := echo.New()

	t.Run("Rejects removed reconcile option and unknown option with 400", func(t *testing.T) {
		for _, badOpt := range []string{"reconcile", "drop", "true"} {
			req := httptest.NewRequest(http.MethodDelete, "/beetle/migration/middleware/ns/mig01/objectStorage/os-01?option="+badOpt, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("nsId", "osId")
			c.SetParamValues("mig01", "os-01")

			require.NoError(t, DeleteObjectStorage(c))
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Rejects invalid option even when Prefer respond-async header is set", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/beetle/migration/middleware/ns/mig01/objectStorage/os-01?option=reconcile", nil)
		req.Header.Set("Prefer", "respond-async")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "osId")
		c.SetParamValues("mig01", "os-01")

		require.NoError(t, DeleteObjectStorage(c))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("Returns 404 when option=empty targets a non-existent bucket", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "empty", r.URL.Query().Get("option"))
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message":"object storage os-01 not found"}`))
		})

		req := httptest.NewRequest(http.MethodDelete, "/beetle/migration/middleware/ns/mig01/objectStorage/os-01?option=empty", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "osId")
		c.SetParamValues("mig01", "os-01")

		require.NoError(t, DeleteObjectStorage(c))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("Returns 409 when standard delete fails because bucket is not empty", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"message":"BucketNotEmpty: The bucket you tried to delete is not empty"}`))
		})

		req := httptest.NewRequest(http.MethodDelete, "/beetle/migration/middleware/ns/mig01/objectStorage/os-01", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "osId")
		c.SetParamValues("mig01", "os-01")

		require.NoError(t, DeleteObjectStorage(c))
		assert.Equal(t, http.StatusConflict, rec.Code)
	})
}

func TestMigrateObjectStorageController_CreatedBodyAndConflict(t *testing.T) {
	e := echo.New()

	t.Run("Returns 201 Created with MigratedObjectStorageResult JSON body", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			switch {
			case r.Method == http.MethodGet && r.URL.Path == "/connConfig/aws-ap-northeast-2":
				_ = json.NewEncoder(w).Encode(tbmodel.ConnConfig{ConfigName: "aws-ap-northeast-2"})
			case r.Method == http.MethodPut && r.URL.Path == "/ns/mig01/resources/objectStorage":
				var createReq tbmodel.ObjectStorageCreateRequest
				require.NoError(t, json.NewDecoder(r.Body).Decode(&createReq))
				_ = json.NewEncoder(w).Encode(tbmodel.ObjectStorageInfo{
					Id:              createReq.BucketName,
					Name:            createReq.BucketName,
					CspResourceName: "csp-" + createReq.BucketName,
					ConnectionName:  createReq.ConnectionName,
					Status:          "Available",
				})
			default:
				w.WriteHeader(http.StatusNotFound)
			}
		})

		body := `{
			"status": "recommended",
			"targetCloud": {"csp": "aws", "region": "ap-northeast-2"},
			"targetObjectStorages": [
				{"sourceBucketName": "src-01", "bucketName": "os-01", "versioningEnabled": false, "corsEnabled": false}
			]
		}`
		req := httptest.NewRequest(http.MethodPost, "/beetle/migration/middleware/ns/mig01/objectStorage?nameSeed=test", bytes.NewBufferString(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId")
		c.SetParamValues("mig01")

		require.NoError(t, MigrateObjectStorage(c))
		assert.Equal(t, http.StatusCreated, rec.Code)

		var resp model.ApiResponse[migration.MigratedObjectStorageResult]
		require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
		assert.True(t, resp.Success)
		assert.Equal(t, "created", resp.Data.Status)
		require.Len(t, resp.Data.ObjectStorage, 1)
		assert.Equal(t, "test-os-01", resp.Data.ObjectStorage[0].Id)
	})

	t.Run("Returns 409 Conflict when all target buckets collide", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
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

		body := `{
			"targetCloud": {"csp": "aws", "region": "ap-northeast-2"},
			"targetObjectStorages": [
				{"sourceBucketName": "src-01", "bucketName": "os-01"}
			]
		}`
		req := httptest.NewRequest(http.MethodPost, "/beetle/migration/middleware/ns/mig01/objectStorage", bytes.NewBufferString(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId")
		c.SetParamValues("mig01")

		require.NoError(t, MigrateObjectStorage(c))
		assert.Equal(t, http.StatusConflict, rec.Code)
	})
}

func TestGetStorageObjectController_HeadResponseHeaders(t *testing.T) {
	e := echo.New()

	t.Run("Sets ETag, Last-Modified, Content-Length, and X-Storage-Class headers with empty body", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodHead, r.Method)
			assert.Equal(t, "/ns/mig01/resources/objectStorage/os-01/object/folder%2Freport.pdf", r.URL.EscapedPath())
			w.Header().Set("ETag", "\"etag-789\"")
			w.Header().Set("Last-Modified", "Wed, 11 Nov 2026 03:00:00 GMT")
			w.Header().Set("Content-Length", "8192")
			w.Header().Set("x-amz-storage-class", "STANDARD")
			w.WriteHeader(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodHead, "/beetle/migration/middleware/ns/mig01/objectStorage/os-01/object/folder%2Freport.pdf", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "osId", "objectKey")
		c.SetParamValues("mig01", "os-01", "folder%2Freport.pdf")

		require.NoError(t, GetStorageObject(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 0, rec.Body.Len())
		assert.Equal(t, "\"etag-789\"", rec.Header().Get("ETag"))
		assert.Equal(t, "Wed, 11 Nov 2026 03:00:00 GMT", rec.Header().Get("Last-Modified"))
		assert.Equal(t, "8192", rec.Header().Get("Content-Length"))
		assert.Equal(t, "STANDARD", rec.Header().Get("X-Storage-Class"))
	})

	t.Run("Omits Content-Length header when upstream does not provide it", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("ETag", "\"etag-legacy\"")
			w.WriteHeader(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodHead, "/beetle/migration/middleware/ns/mig01/objectStorage/os-01/object/file.txt", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "osId", "objectKey")
		c.SetParamValues("mig01", "os-01", "file.txt")

		require.NoError(t, GetStorageObject(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "", rec.Header().Get("Content-Length"))
	})
}

func TestTestClientHelper_NoFallbackAndRegionValidation(t *testing.T) {
	e := echo.New()

	t.Run("Resolves supported S3 endpoints", func(t *testing.T) {
		epAws, err := getS3EndpointForCsp("AWS", "ap-northeast-2")
		require.NoError(t, err)
		assert.Equal(t, "s3.ap-northeast-2.amazonaws.com", epAws)

		epNcp, err := getS3EndpointForCsp("ncp", "kr")
		require.NoError(t, err)
		assert.Equal(t, "kr.object.ncloudstorage.com", epNcp)

		epIbm, err := getS3EndpointForCsp("ibm", "au-syd")
		require.NoError(t, err)
		assert.Equal(t, "s3.au-syd.cloud-object-storage.appdomain.cloud", epIbm)
	})

	t.Run("Rejects unsupported CSP and malformed region without fallback", func(t *testing.T) {
		_, err := getS3EndpointForCsp("azure", "koreacentral")
		require.Error(t, err)

		_, err = getS3EndpointForCsp("aws", "ap-northeast-2/evil")
		require.Error(t, err)

		scanBody := `{"csp":"azure","region":"koreacentral","accessKeyId":"k","secretAccessKey":"s"}`
		req := httptest.NewRequest(http.MethodPost, "/beetle/migration/middleware/objectStorage/scan", bytes.NewBufferString(scanBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		require.NoError(t, ScanObjectStorage(c))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestRecommendObjectStorageController_CaseInsensitiveSupport(t *testing.T) {
	e := echo.New()

	setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "ncp", r.URL.Query().Get("providerName"))
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(tbmodel.ObjectStorageSupportResponse{
			Supports: map[string]tbmodel.ObjectStorageFeatureSupport{
				"ncp": {Versioning: true, Cors: false},
			},
		})
	})

	body := `{
		"desiredCloud": {"csp": "aws", "region": "ap-northeast-2"},
		"sourceObjectStorages": [
			{
				"bucketName": "src-bucket",
				"versioningEnabled": true,
				"corsEnabled": true,
				"corsRule": [{"allowedOrigin": ["*"], "allowedMethod": ["GET"]}]
			}
		]
	}`
	req := httptest.NewRequest(http.MethodPost, "/beetle/recommendation/middleware/objectStorage?desiredCsp=NCP&desiredRegion=kr", bytes.NewBufferString(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	require.NoError(t, RecommendObjectStorage(c))
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp model.ApiResponse[storagemodel.RecommendedObjectStorage]
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.Equal(t, "partial", resp.Data.Status)
	assert.Equal(t, "ncp", resp.Data.TargetCloud.Csp)
	assert.Equal(t, "kr", resp.Data.TargetCloud.Region)
	require.Len(t, resp.Data.TargetObjectStorages, 1)
	assert.False(t, resp.Data.TargetObjectStorages[0].CORSEnabled)
	assert.Len(t, resp.Data.Warnings, 1)
}
