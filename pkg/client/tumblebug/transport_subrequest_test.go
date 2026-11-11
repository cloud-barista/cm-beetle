package tbclient

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

func TestTransportSubRequestCorrelationAnd429RetryUniqueIDs(t *testing.T) {
	var postAttempts atomic.Int32
	var mu sync.Mutex
	var seenGetSource string
	var seenGetReqID string
	var seenPostReqIDs []string
	var seenPostSources []string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet {
			mu.Lock()
			seenGetSource = r.Header.Get("X-Request-Source")
			seenGetReqID = r.Header.Get("X-Request-Id")
			mu.Unlock()
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"message":"ready"}`))
			return
		}

		attempt := postAttempts.Add(1)
		mu.Lock()
		seenPostReqIDs = append(seenPostReqIDs, r.Header.Get("X-Request-Id"))
		seenPostSources = append(seenPostSources, r.Header.Get("X-Request-Source"))
		mu.Unlock()

		if attempt == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"message":"rate limit exceeded"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"mig01-vnet","name":"mig01-vnet"}`))
	}))
	defer srv.Close()

	cli := NewClient(ApiConfig{
		RestUrl:       srv.URL,
		Username:      "default",
		Password:      "default",
		Timeout:       10 * time.Second,
		CacheDuration: 10 * time.Millisecond,
	})

	parentReqID := fmt.Sprintf("tb-corr-%d", time.Now().UnixNano())
	defer common.DiscardSubRequests(parentReqID)
	ctx := common.WithRequestID(context.Background(), parentReqID)

	// 1. GET call should propagate X-Request-Source without recording a subRequest.
	ready, _, err := cli.NewSession().SetContext(ctx).IsReady()
	if err != nil || !ready {
		t.Fatalf("IsReady failed: ready=%v err=%v", ready, err)
	}
	subsAfterGet, _ := common.SnapshotSubRequests(parentReqID)
	if len(subsAfterGet) != 0 {
		t.Fatalf("expected GET not to be recorded in subRequests, got %+v", subsAfterGet)
	}

	// 2. POST call that hits 429 once and succeeds on retry should issue distinct subReqIds.
	vNetInfo, err := cli.NewSession().SetContext(ctx).CreateVNet("mig01", tbmodel.VNetReq{Name: "mig01-vnet"})
	if err != nil {
		t.Fatalf("CreateVNet failed after retry: %v", err)
	}
	if vNetInfo.Id != "mig01-vnet" {
		t.Fatalf("unexpected vNetInfo: %+v", vNetInfo)
	}

	mu.Lock()
	defer mu.Unlock()

	wantSource := "cm-beetle:" + parentReqID
	if seenGetSource != wantSource {
		t.Errorf("expected GET X-Request-Source=%q, got %q", wantSource, seenGetSource)
	}
	if seenGetReqID != "" {
		t.Errorf("expected GET not to stamp parent X-Request-Id, got %q", seenGetReqID)
	}
	if len(seenPostReqIDs) != 2 {
		t.Fatalf("expected 2 POST attempts (429 + 200), got %v", seenPostReqIDs)
	}
	if seenPostReqIDs[0] != parentReqID+"-1" || seenPostReqIDs[1] != parentReqID+"-1-r2" {
		t.Errorf("expected distinct retry subReqIds [%s-1, %s-1-r2], got %v", parentReqID, parentReqID, seenPostReqIDs)
	}
	for _, src := range seenPostSources {
		if src != wantSource {
			t.Errorf("expected POST X-Request-Source=%q, got %q", wantSource, src)
		}
	}

	subs, summary := common.FinalizeSubRequests(parentReqID)
	if len(subs) != 1 || summary == nil || summary.Total != 1 || summary.Success != 1 || summary.Error != 0 {
		t.Fatalf("unexpected finalized subRequests: subs=%+v summary=%+v", subs, summary)
	}
	if subs[0].SubReqId != parentReqID+"-1-r2" || subs[0].ResourceType != "vNet" || subs[0].ResourceHint != "mig01-vnet" {
		t.Errorf("unexpected extracted resource metadata: %+v", subs[0])
	}
}

func TestSessionSetContextPreservesCacheDurationAndReadyz(t *testing.T) {
	var readyCalls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		call := readyCalls.Add(1)
		w.WriteHeader(http.StatusOK)
		if call == 1 {
			_, _ = w.Write([]byte(`{"message":"initializing","ready":true,"initialized":false}`))
			return
		}
		_, _ = w.Write([]byte(`{"message":"ready","ready":true,"initialized":true}`))
	}))
	defer srv.Close()

	cli := NewClient(ApiConfig{
		RestUrl:       srv.URL,
		Username:      "default",
		Password:      "default",
		Timeout:       5 * time.Second,
		CacheDuration: 5 * time.Second,
	})

	parentReqID := fmt.Sprintf("cache-ctx-%d", time.Now().UnixNano())
	ctx := common.WithRequestID(context.Background(), parentReqID)

	sess := cli.NewSession().NoCache().SetContext(ctx)
	d, ok := GetCacheDuration(sess.Context())
	if !ok || d != 0 {
		t.Fatalf("expected NoCache duration 0 to survive SetContext, got d=%v ok=%v", d, ok)
	}
	if gotID := common.RequestIDFromContext(sess.Context()); gotID != parentReqID {
		t.Fatalf("expected reqID %q in session context, got %q", parentReqID, gotID)
	}

	r1, err := cli.NewSession().SetContext(ctx).CheckReadyz()
	if err != nil || !r1.Ready || r1.Initialized {
		t.Fatalf("unexpected first CheckReadyz: %+v err=%v", r1, err)
	}
	r2, err := cli.NewSession().SetContext(ctx).CheckReadyz()
	if err != nil || !r2.Ready || !r2.Initialized {
		t.Fatalf("expected second CheckReadyz to bypass cache and report initialized=true: %+v err=%v", r2, err)
	}
}
