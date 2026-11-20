package tbclient

import (
	"errors"
	"net/http"
	"testing"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	"github.com/cloud-barista/cm-beetle/pkg/ratelimit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRDBMSClient_APIErrorStatusPreservation(t *testing.T) {
	cases := []struct {
		name       string
		statusCode int
		body       string
		invoke     func(s *Session) error
	}{
		{
			name:       "CreateRDBMSDatabase preserves 422",
			statusCode: http.StatusUnprocessableEntity,
			body:       `{"message":"private RDBMS requires direct SQL"}`,
			invoke: func(s *Session) error {
				return s.CreateRDBMSDatabase("mig01", "db-01", "dbadmin", "secret", rdbmsmodel.RDBMSDatabaseCreateReq{DatabaseName: "appdb"})
			},
		},
		{
			name:       "ListRDBMSDatabases preserves 422",
			statusCode: http.StatusUnprocessableEntity,
			body:       `{"message":"aws has no native API to list logical databases"}`,
			invoke: func(s *Session) error {
				_, err := s.ListRDBMSDatabases("mig01", "db-01", "dbadmin", "secret")
				return err
			},
		},
		{
			name:       "DeleteRDBMSDatabase preserves 422",
			statusCode: http.StatusUnprocessableEntity,
			body:       `{"message":"aws has no native API to delete logical databases"}`,
			invoke: func(s *Session) error {
				return s.DeleteRDBMSDatabase("mig01", "db-01", "appdb", "dbadmin", "secret")
			},
		},
		{
			name:       "GetRDBMSSecureTransport preserves 422",
			statusCode: http.StatusUnprocessableEntity,
			body:       `{"message":"secure-transport probe requires direct TLS/SQL access"}`,
			invoke: func(s *Session) error {
				_, err := s.GetRDBMSSecureTransport("mig01", "db-01", "dbadmin", "secret")
				return err
			},
		},
		{
			name:       "GetRDBMS preserves 404 and string format",
			statusCode: http.StatusNotFound,
			body:       `{"message":"rdbms not found"}`,
			invoke: func(s *Session) error {
				_, err := s.GetRDBMS("mig01", "missing-db")
				return err
			},
		},
		{
			name:       "DeleteRDBMS preserves 404 and string format",
			statusCode: http.StatusNotFound,
			body:       `{"message":"rdbms not found"}`,
			invoke: func(s *Session) error {
				return s.DeleteRDBMS("mig01", "missing-db", "")
			},
		},
		{
			name:       "CreateRDBMSDatabase preserves 400",
			statusCode: http.StatusBadRequest,
			body:       `{"message":"invalid databaseName"}`,
			invoke: func(s *Session) error {
				return s.CreateRDBMSDatabase("mig01", "db-01", "dbadmin", "secret", rdbmsmodel.RDBMSDatabaseCreateReq{DatabaseName: "bad name"})
			},
		},
		{
			name:       "GetRDBMSSecureTransport preserves 500",
			statusCode: http.StatusInternalServerError,
			body:       `{"message":"internal error"}`,
			invoke: func(s *Session) error {
				_, err := s.GetRDBMSSecureTransport("mig01", "db-01", "dbadmin", "secret")
				return err
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cli := newLoopbackTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.statusCode)
				_, _ = w.Write([]byte(tc.body))
			})

			err := tc.invoke(cli.NewSession())
			require.Error(t, err)

			var apiErr *APIError
			require.True(t, errors.As(err, &apiErr), "expected *APIError, got %T", err)
			assert.Equal(t, tc.statusCode, apiErr.StatusCode)
			assert.Equal(t, tc.body, apiErr.Body)
			assert.Contains(t, err.Error(), "API Error:")
			assert.Contains(t, err.Error(), tc.body)
		})
	}

	t.Run("429 Too Many Requests still returns ratelimit.ErrLimited", func(t *testing.T) {
		cli := newLoopbackTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTooManyRequests)
		})

		_, err := cli.NewSession().GetRDBMSSecureTransport("mig01", "db-01", "dbadmin", "secret")
		require.Error(t, err)
		var limitedErr *ratelimit.ErrLimited
		assert.True(t, errors.As(err, &limitedErr))
	})
}
