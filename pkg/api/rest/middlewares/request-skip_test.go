package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestShouldSkipRequestTracking(t *testing.T) {
	e := echo.New()

	skippedPaths := []string{
		"/beetle/request/req-123",
		"/beetle/requests",
		"/beetle/readyz",
		"/beetle/httpVersion",
		"/beetle/api",
		"/beetle/swagger/index.html",
		"/beetle/docs/openapi.json",
		"/beetle/ui/index.html",
		"/tumblebug/ns",
		"/favicon.ico",
	}
	for _, p := range skippedPaths {
		req := httptest.NewRequest(http.MethodGet, p, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.True(t, ShouldSkipRequestTracking(c), "expected %s to skip request tracking", p)
	}

	trackedPaths := []string{
		"/beetle/recommendation/infra",
		"/beetle/migration/ns/mig01/infra",
		"/beetle/migration/ns/mig01/infra/mig-infra-01",
		"/beetle/artifact/lineage/art-123",
	}
	for _, p := range trackedPaths {
		req := httptest.NewRequest(http.MethodPost, p, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.False(t, ShouldSkipRequestTracking(c), "expected %s to be tracked", p)
	}
}
