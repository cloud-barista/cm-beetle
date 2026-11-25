package tbclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

func TestTumblebugTransport_PerInstanceIsolationAndCleanup(t *testing.T) {
	var hits atomic.Int32
	var shouldFail atomic.Bool

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Test-Origin", "server")
		if shouldFail.Load() {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"message":"err"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"ok"}`))
	}))
	defer srv.Close()

	cliA := NewClient(ApiConfig{
		RestUrl:       srv.URL,
		Username:      "default",
		Password:      "default",
		Timeout:       5 * time.Second,
		CacheDuration: time.Minute,
	})
	cliB := NewClient(ApiConfig{
		RestUrl:       srv.URL,
		Username:      "default",
		Password:      "default",
		Timeout:       5 * time.Second,
		CacheDuration: time.Minute,
	})

	// 1. Two independent clients must not share TumblebugTransport cache state.
	if _, _, err := cliA.NewSession().IsReady(); err != nil {
		t.Fatalf("cliA IsReady failed: %v", err)
	}
	if _, _, err := cliA.NewSession().IsReady(); err != nil {
		t.Fatalf("cliA cached IsReady failed: %v", err)
	}
	if _, _, err := cliB.NewSession().IsReady(); err != nil {
		t.Fatalf("cliB IsReady failed: %v", err)
	}
	if got := hits.Load(); got != 2 {
		t.Fatalf("expected 2 upstream hits across 2 isolated transports, got %d", got)
	}

	// 2. Circuit breaker tripped on cliA must not affect cliB.
	shouldFail.Store(true)
	for i := 0; i < circuitBreakerFailureThreshold; i++ {
		_, _, _ = cliA.NewSession().NoCache().IsReady()
	}
	trA := cliA.restyClient.GetClient().Transport.(*TumblebugTransport)
	trB := cliB.restyClient.GetClient().Transport.(*TumblebugTransport)

	trA.mu.Lock()
	aBreakerCount := len(trA.circuitBreakers)
	aCounterCount := len(trA.requestCounter)
	trA.mu.Unlock()
	trB.mu.Lock()
	bBreakerCount := len(trB.circuitBreakers)
	trB.mu.Unlock()
	if aBreakerCount == 0 || aCounterCount != 0 || bBreakerCount != 0 {
		t.Fatalf("expected cliA circuit breaker open, cliA zero counters, cliB zero breakers, got aBreakers=%d aCounters=%d bBreakers=%d",
			aBreakerCount, aCounterCount, bBreakerCount)
	}

	shouldFail.Store(false)
	if _, _, err := cliB.NewSession().NoCache().IsReady(); err != nil {
		t.Fatalf("expected cliB not to be blocked by cliA circuit breaker, got: %v", err)
	}

	// 3. Resetting circuit breaker clears the map entry on healthy success.
	trA.recordCircuitBreakerSuccess("GET_" + srv.URL + "/readyz")
	trA.mu.Lock()
	aBreakerAfterReset := len(trA.circuitBreakers)
	trA.mu.Unlock()
	if aBreakerAfterReset != 0 {
		t.Fatalf("expected circuit breaker entry to be deleted on success, got %d", aBreakerAfterReset)
	}
}

func TestTumblebugTransport_ConcurrentCounterCleanupAndNoCacheOverride(t *testing.T) {
	var hits atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		hits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"ok"}`))
	}))
	defer srv.Close()

	cli := NewClient(ApiConfig{
		RestUrl:       srv.URL,
		Username:      "default",
		Password:      "default",
		Timeout:       5 * time.Second,
		CacheDuration: time.Minute,
	})

	// Explicit Session.NoCache() must override a positive context cache duration.
	parentCtx := SetCacheDuration(context.Background(), time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _, _ = cli.NewSession().SetContext(parentCtx).NoCache().IsReady()
		}()
	}
	wg.Wait()

	if got := hits.Load(); got != 8 {
		t.Fatalf("expected NoCache() to bypass parent cache duration for all 8 calls, got %d hits", got)
	}

	tr := cli.restyClient.GetClient().Transport.(*TumblebugTransport)
	tr.mu.Lock()
	defer tr.mu.Unlock()
	if len(tr.requestCounter) != 0 {
		t.Fatalf("expected requestCounter map to be empty after requests complete, got %+v", tr.requestCounter)
	}
}

func TestClassifySubRequestResourceAndListNlbsNoCacheRetry(t *testing.T) {
	if got := classifySubRequestResource([]string{"specImagePairReview"}); got != "specImageReview" {
		t.Fatalf("expected specImageReview classification, got %q", got)
	}

	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if calls.Add(1) == 1 {
			_, _ = w.Write([]byte(`{"nlb":[]}`))
			return
		}
		_, _ = w.Write([]byte(`{"nlb":[{"id":"nlb-01","name":"nlb-01"}]}`))
	}))
	defer srv.Close()

	cli := NewClient(ApiConfig{
		RestUrl:       srv.URL,
		Username:      "default",
		Password:      "default",
		Timeout:       15 * time.Second,
		CacheDuration: time.Minute,
	})

	// ListNlbs retries after 5s and uses NoCache() so the second attempt must not hit the empty cached response.
	res, err := cli.NewSession().SetContext(context.Background()).ListNlbs("mig01", "infra-01")
	if err != nil {
		t.Fatalf("ListNlbs failed: %v", err)
	}
	if calls.Load() != 2 || len(res.NLB) != 1 || res.NLB[0].Id != "nlb-01" {
		t.Fatalf("expected 2 uncached attempts returning nlb-01, got calls=%d res=%+v", calls.Load(), res)
	}
	_ = tbmodel.NLBInfo{}
}
