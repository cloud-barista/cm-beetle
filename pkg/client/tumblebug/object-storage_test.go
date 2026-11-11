package tbclient

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloud-barista/cm-beetle/pkg/ratelimit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newLoopbackTestClient(t *testing.T, handler http.HandlerFunc) *TumblebugClient {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	cli := NewClient(ApiConfig{
		RestUrl:       srv.URL,
		Username:      "default",
		Password:      "default",
		HttpTransport: http.DefaultTransport,
	})
	cli.restyClient.SetRetryCount(0)
	return cli
}

func TestExistObjectStorage_StatusHandling(t *testing.T) {
	t.Run("200 OK returns true without error", func(t *testing.T) {
		cli := newLoopbackTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodHead, r.Method)
			assert.Equal(t, "/ns/mig01/resources/objectStorage/os-01", r.URL.Path)
			w.WriteHeader(http.StatusOK)
		})

		exists, err := cli.NewSession().ExistObjectStorage("mig01", "os-01")
		require.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("404 Not Found returns false without error", func(t *testing.T) {
		cli := newLoopbackTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})

		exists, err := cli.NewSession().ExistObjectStorage("mig01", "missing-os")
		require.NoError(t, err)
		assert.False(t, exists)
	})

	t.Run("429 Too Many Requests returns ratelimit.ErrLimited", func(t *testing.T) {
		cli := newLoopbackTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTooManyRequests)
		})

		exists, err := cli.NewSession().ExistObjectStorage("mig01", "os-01")
		require.Error(t, err)
		assert.False(t, exists)
		var limitedErr *ratelimit.ErrLimited
		assert.True(t, errors.As(err, &limitedErr))
	})

	t.Run("500 Internal Server Error returns explicit error instead of false", func(t *testing.T) {
		cli := newLoopbackTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})

		exists, err := cli.NewSession().ExistObjectStorage("mig01", "os-01")
		require.Error(t, err)
		assert.False(t, exists)
		assert.Contains(t, err.Error(), "500")
	})
}

func TestGetAndDeleteStorageObject_HeadersAndPathEncoding(t *testing.T) {
	t.Run("GetStorageObject normalizes nested and pre-encoded objectKey and parses headers", func(t *testing.T) {
		var capturedEscapedPaths []string
		cli := newLoopbackTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodHead, r.Method)
			capturedEscapedPaths = append(capturedEscapedPaths, r.URL.EscapedPath())
			w.Header().Set("ETag", "\"abc123\"")
			w.Header().Set("Last-Modified", "Wed, 11 Nov 2026 02:00:00 GMT")
			w.Header().Set("Content-Length", "4096")
			w.Header().Set("x-amz-storage-class", "STANDARD")
			w.WriteHeader(http.StatusOK)
		})

		obj1, err := cli.NewSession().GetStorageObject("mig01", "os-01", "folder/sub/file.txt")
		require.NoError(t, err)
		assert.Equal(t, "folder/sub/file.txt", obj1.Key)
		assert.Equal(t, int64(4096), obj1.Size)
		assert.Equal(t, "\"abc123\"", obj1.ETag)
		assert.Equal(t, "Wed, 11 Nov 2026 02:00:00 GMT", obj1.LastModified)
		assert.Equal(t, "STANDARD", obj1.StorageClass)

		obj2, err := cli.NewSession().GetStorageObject("mig01", "os-01", "folder%2Fsub%2Ffile.txt")
		require.NoError(t, err)
		assert.Equal(t, "folder/sub/file.txt", obj2.Key)

		// Verify traversal-like keys stay encoded inside the single objectKey segment.
		_, err = cli.NewSession().GetStorageObject("mig01", "os-01", "../../etc/passwd")
		require.NoError(t, err)

		require.Len(t, capturedEscapedPaths, 3)
		assert.Equal(t, "/ns/mig01/resources/objectStorage/os-01/object/folder%2Fsub%2Ffile.txt", capturedEscapedPaths[0])
		assert.Equal(t, "/ns/mig01/resources/objectStorage/os-01/object/folder%2Fsub%2Ffile.txt", capturedEscapedPaths[1])
		assert.Equal(t, "/ns/mig01/resources/objectStorage/os-01/object/..%2F..%2Fetc%2Fpasswd", capturedEscapedPaths[2])
	})

	t.Run("GetStorageObject sets Size to -1 when upstream omits Content-Length", func(t *testing.T) {
		cli := newLoopbackTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("ETag", "\"etag-only\"")
			w.WriteHeader(http.StatusOK)
		})

		obj, err := cli.NewSession().GetStorageObject("mig01", "os-01", "readme.md")
		require.NoError(t, err)
		assert.Equal(t, int64(-1), obj.Size)
	})

	t.Run("DeleteStorageObject encodes slash in objectKey as single segment", func(t *testing.T) {
		var escapedPath string
		cli := newLoopbackTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			escapedPath = r.URL.EscapedPath()
			w.WriteHeader(http.StatusNoContent)
		})

		err := cli.NewSession().DeleteStorageObject("mig01", "os-01", "logs/2026/app.log")
		require.NoError(t, err)
		assert.Equal(t, "/ns/mig01/resources/objectStorage/os-01/object/logs%2F2026%2Fapp.log", escapedPath)
	})
}
