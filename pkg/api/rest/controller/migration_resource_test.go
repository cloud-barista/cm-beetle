package controller

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync/atomic"
	"testing"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/middlewares"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTumblebugProxy_RewritesAuthAndCachesHandler(t *testing.T) {
	var hits atomic.Int32
	var lastPath, lastAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		lastPath = r.URL.Path
		lastAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"vnet":[]}`))
	}))
	defer srv.Close()

	prevTB := config.Tumblebug
	config.Tumblebug = config.TumblebugConfig{
		RestUrl: srv.URL,
		API: config.TumblebugApiConfig{
			Username: "tb-admin",
			Password: "tb-secret",
		},
	}
	t.Cleanup(func() {
		config.Tumblebug = prevTB
	})

	h1 := createTumblebugProxyHandler("/migration/ns/*/resources/vNet", "/ns/$1/resources/vNet")
	h2 := createTumblebugProxyHandler("/migration/ns/*/resources/vNet", "/ns/$1/resources/vNet")
	assert.Equal(t, reflect.ValueOf(h1).Pointer(), reflect.ValueOf(h2).Pointer(), "expected proxy handler to be cached")

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/beetle/migration/ns/mig01/resources/vNet", nil)
	rec := httptest.NewRecorder()
	require.NoError(t, ListMigratedVNets(e.NewContext(req, rec)))
	require.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "/ns/mig01/resources/vNet", lastPath)
	assert.Equal(t, middlewares.BasicAuth("tb-admin", "tb-secret"), lastAuth)

	req2 := httptest.NewRequest(http.MethodGet, "/beetle/migration/ns/mig02/resources/vNet", nil)
	rec2 := httptest.NewRecorder()
	require.NoError(t, ListMigratedVNets(e.NewContext(req2, rec2)))
	require.Equal(t, http.StatusOK, rec2.Code)
	assert.Equal(t, "/ns/mig02/resources/vNet", lastPath)
	assert.Equal(t, int32(2), hits.Load())
}
