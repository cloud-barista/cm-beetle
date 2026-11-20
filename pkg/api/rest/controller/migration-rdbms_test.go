package controller

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTumblebugErrorStatusCode(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{
			name: "400 Bad Request preserved",
			err:  &tbclient.APIError{StatusCode: http.StatusBadRequest, Status: "400 Bad Request", Body: "bad request"},
			want: http.StatusBadRequest,
		},
		{
			name: "404 Not Found preserved through fmt.Errorf wrapping",
			err:  fmt.Errorf("wrapped: %w", &tbclient.APIError{StatusCode: http.StatusNotFound, Status: "404 Not Found", Body: "missing"}),
			want: http.StatusNotFound,
		},
		{
			name: "409 Conflict preserved",
			err:  &tbclient.APIError{StatusCode: http.StatusConflict, Status: "409 Conflict", Body: "conflict"},
			want: http.StatusConflict,
		},
		{
			name: "422 Unprocessable Entity preserved",
			err:  &tbclient.APIError{StatusCode: http.StatusUnprocessableEntity, Status: "422 Unprocessable Entity", Body: "private fast-fail"},
			want: http.StatusUnprocessableEntity,
		},
		{
			name: "500 Internal Server Error maps to 500",
			err:  &tbclient.APIError{StatusCode: http.StatusInternalServerError, Status: "500 Internal Server Error", Body: "boom"},
			want: http.StatusInternalServerError,
		},
		{
			name: "Generic error maps to 500",
			err:  errors.New("network timeout"),
			want: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tumblebugErrorStatusCode(tc.err))
		})
	}
}

func TestMigrationRDBMSControllers_PreserveTumblebug4xxStatus(t *testing.T) {
	e := echo.New()

	t.Run("GetRDBMS propagates 404 from Tumblebug", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message":"rdbms missing-db not found"}`))
		})

		req := httptest.NewRequest(http.MethodGet, "/beetle/migration/middleware/ns/mig01/rdbms/missing-db", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "rdbmsId")
		c.SetParamValues("mig01", "missing-db")

		require.NoError(t, GetRDBMS(c))
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("CreateRDBMSDatabase propagates 422 from Tumblebug", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_, _ = w.Write([]byte(`{"message":"private RDBMS requires direct SQL"}`))
		})

		body := bytes.NewBufferString(`{"databaseName":"appdb","charset":"utf8mb4","collation":"utf8mb4_unicode_ci"}`)
		req := httptest.NewRequest(http.MethodPost, "/beetle/migration/middleware/ns/mig01/rdbms/db-01/database", body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("X-Admin-User-Name", "dbadmin")
		req.Header.Set("X-Admin-User-Password", "Password123!")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "rdbmsId")
		c.SetParamValues("mig01", "db-01")

		require.NoError(t, CreateRDBMSDatabase(c))
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Contains(t, rec.Body.String(), "private RDBMS requires direct SQL")
	})

	t.Run("ListRDBMSDatabases propagates 422 from Tumblebug", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_, _ = w.Write([]byte(`{"message":"aws has no native API to list logical databases"}`))
		})

		req := httptest.NewRequest(http.MethodGet, "/beetle/migration/middleware/ns/mig01/rdbms/db-01/database", nil)
		req.Header.Set("X-Admin-User-Name", "dbadmin")
		req.Header.Set("X-Admin-User-Password", "Password123!")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "rdbmsId")
		c.SetParamValues("mig01", "db-01")

		require.NoError(t, ListRDBMSDatabases(c))
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("DeleteRDBMSDatabase propagates 422 from Tumblebug", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_, _ = w.Write([]byte(`{"message":"aws has no native API to delete logical databases"}`))
		})

		req := httptest.NewRequest(http.MethodDelete, "/beetle/migration/middleware/ns/mig01/rdbms/db-01/database/appdb", nil)
		req.Header.Set("X-Admin-User-Name", "dbadmin")
		req.Header.Set("X-Admin-User-Password", "Password123!")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "rdbmsId", "dbName")
		c.SetParamValues("mig01", "db-01", "appdb")

		require.NoError(t, DeleteRDBMSDatabase(c))
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("GetRDBMSSecureTransport propagates 422 from Tumblebug", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			_, _ = w.Write([]byte(`{"message":"secure-transport probe requires direct TLS/SQL access"}`))
		})

		req := httptest.NewRequest(http.MethodGet, "/beetle/migration/middleware/ns/mig01/rdbms/db-01/secure-transport", nil)
		req.Header.Set("X-Admin-User-Name", "dbadmin")
		req.Header.Set("X-Admin-User-Password", "Password123!")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "rdbmsId")
		c.SetParamValues("mig01", "db-01")

		require.NoError(t, GetRDBMSSecureTransport(c))
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		assert.Contains(t, rec.Body.String(), "secure-transport probe requires direct TLS/SQL access")
	})

	t.Run("CreateRDBMSDatabase propagates 400 from Tumblebug", func(t *testing.T) {
		setupMockTumblebugForController(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"message":"invalid databaseName"}`))
		})

		body := bytes.NewBufferString(`{"databaseName":"bad name"}`)
		req := httptest.NewRequest(http.MethodPost, "/beetle/migration/middleware/ns/mig01/rdbms/db-01/database", body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set("X-Admin-User-Name", "dbadmin")
		req.Header.Set("X-Admin-User-Password", "Password123!")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("nsId", "rdbmsId")
		c.SetParamValues("mig01", "db-01")

		require.NoError(t, CreateRDBMSDatabase(c))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
