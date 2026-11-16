package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/labstack/echo/v4"
)

func TestRequestIdAndDetailsIssuer_AttachesRequestCacheSyncAndAsync(t *testing.T) {
	var specCalls atomic.Int32
	tbSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/recommendSpec" {
			specCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`[{"id":"aws+ap-northeast-2+t3.medium"}]`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer tbSrv.Close()

	restore := tbclient.SetClient(tbclient.NewClient(tbclient.ApiConfig{
		RestUrl:  tbSrv.URL,
		Username: "default",
		Password: "default",
	}))
	defer restore()

	e := echo.New()
	e.Use(RequestIdAndDetailsIssuer)

	e.GET("/beetle/test-sync", func(c echo.Context) error {
		ctx := c.Request().Context()
		for i := 0; i < 3; i++ {
			if _, err := tbclient.NewSession().SetContext(ctx).InfraRecommendSpec(`{"limit":5}`); err != nil {
				return err
			}
		}
		return c.NoContent(http.StatusOK)
	})

	var asyncStats atomic.Pointer[tbclient.RequestCacheStats]
	asyncDone := make(chan struct{})
	e.POST("/beetle/test-async", func(c echo.Context) error {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(c.Request().Context(), reqID, func(ctx context.Context) (string, error) {
			defer close(asyncDone)
			for i := 0; i < 4; i++ {
				if _, err := tbclient.NewSession().SetContext(ctx).InfraRecommendSpec(`{"limit":5}`); err != nil {
					return "", err
				}
			}
			if st, ok := tbclient.GetRequestCacheStats(ctx); ok {
				asyncStats.Store(&st)
			}
			return "ok", nil
		})
		if !started {
			return c.NoContent(http.StatusServiceUnavailable)
		}
		return c.NoContent(http.StatusAccepted)
	})

	// 1. Synchronous Echo request automatically deduplicates within the request.
	reqSync := httptest.NewRequest(http.MethodGet, "/beetle/test-sync", nil)
	recSync := httptest.NewRecorder()
	e.ServeHTTP(recSync, reqSync)
	if recSync.Code != http.StatusOK {
		t.Fatalf("expected 200 for sync request, got %d", recSync.Code)
	}
	if got := specCalls.Load(); got != 1 {
		t.Fatalf("expected 1 upstream call inside sync middleware request, got %d", got)
	}
	common.DiscardSubRequests(recSync.Header().Get(echo.HeaderXRequestID))

	// 2. Async RunAsync request inherits an isolated RequestCache from the HTTP middleware context.
	reqAsync := httptest.NewRequest(http.MethodPost, "/beetle/test-async", nil)
	recAsync := httptest.NewRecorder()
	e.ServeHTTP(recAsync, reqAsync)
	if recAsync.Code != http.StatusAccepted {
		t.Fatalf("expected 202 for async request, got %d", recAsync.Code)
	}

	select {
	case <-asyncDone:
	case <-time.After(3 * time.Second):
		t.Fatalf("timed out waiting for async job completion")
	}

	if got := specCalls.Load(); got != 2 {
		t.Fatalf("expected isolated async request to make 1 fresh upstream call (total 2), got %d", got)
	}
	st := asyncStats.Load()
	if st == nil || st.Misses != 1 || st.Hits != 3 {
		t.Fatalf("expected async RunAsync stats misses=1 hits=3, got %+v", st)
	}
	common.DiscardSubRequests(recAsync.Header().Get(echo.HeaderXRequestID))
}
