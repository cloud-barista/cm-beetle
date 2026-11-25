package middlewares

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestCustomRoundTripper_ClonesRequestAndSetsAuth(t *testing.T) {
	var seenAuth string
	rt := &CustomRoundTripper{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			seenAuth = r.Header.Get("Authorization")
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("ok")),
				Header:     make(http.Header),
			}, nil
		}),
		BasicAuthStr: BasicAuth("tbuser", "tbpass"),
	}

	origReq := httptest.NewRequest(http.MethodGet, "http://example.com/test", nil)
	origReq.Header.Set("Authorization", "Bearer original")

	resp, err := rt.RoundTrip(origReq)
	if err != nil {
		t.Fatalf("RoundTrip failed: %v", err)
	}
	_ = resp.Body.Close()

	if seenAuth != BasicAuth("tbuser", "tbpass") {
		t.Errorf("upstream Authorization = %q, want %q", seenAuth, BasicAuth("tbuser", "tbpass"))
	}
	if got := origReq.Header.Get("Authorization"); got != "Bearer original" {
		t.Errorf("original request Header mutated to %q", got)
	}
}

func TestProxy_BasicAuthConditionalInjection(t *testing.T) {
	var upstreamAuth string
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upstreamAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
	}))
	defer upstream.Close()

	targetURL, err := url.Parse(upstream.URL)
	if err != nil {
		t.Fatalf("url.Parse failed: %v", err)
	}

	t.Run("Injects BasicAuth when credentials configured", func(t *testing.T) {
		upstreamAuth = ""
		e := echo.New()
		mw := Proxy(ProxyConfig{
			URL:      targetURL,
			Username: "default",
			Password: "secret",
		})
		h := mw(func(c echo.Context) error { return nil })

		req := httptest.NewRequest(http.MethodGet, "/ns/mig01", nil)
		rec := httptest.NewRecorder()
		if err := h(e.NewContext(req, rec)); err != nil {
			t.Fatalf("proxy handler failed: %v", err)
		}
		if upstreamAuth != BasicAuth("default", "secret") {
			t.Errorf("upstream Authorization = %q, want %q", upstreamAuth, BasicAuth("default", "secret"))
		}
	})

	t.Run("Preserves caller Authorization when credentials empty", func(t *testing.T) {
		upstreamAuth = ""
		e := echo.New()
		mw := Proxy(ProxyConfig{
			URL: targetURL,
		})
		h := mw(func(c echo.Context) error { return nil })

		req := httptest.NewRequest(http.MethodGet, "/ns/mig01", nil)
		req.Header.Set("Authorization", "Bearer caller-token")
		rec := httptest.NewRecorder()
		if err := h(e.NewContext(req, rec)); err != nil {
			t.Fatalf("proxy handler failed: %v", err)
		}
		if upstreamAuth != "Bearer caller-token" {
			t.Errorf("upstream Authorization = %q, want caller-token", upstreamAuth)
		}
	})
}
