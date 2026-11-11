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
	if seenPostReqIDs[0] != parentReqID+"-1" || seenPostReqIDs[1] != parentReqID+"-2" {
		t.Errorf("expected distinct sequential subReqIds [%s-1, %s-2], got %v", parentReqID, parentReqID, seenPostReqIDs)
	}
	for _, src := range seenPostSources {
		if src != wantSource {
			t.Errorf("expected POST X-Request-Source=%q, got %q", wantSource, src)
		}
	}

	subs, summary := common.FinalizeSubRequests(parentReqID)
	if len(subs) != 2 || summary == nil || summary.Total != 2 || summary.Success != 1 || summary.Error != 1 {
		t.Fatalf("unexpected finalized subRequests: subs=%+v summary=%+v", subs, summary)
	}
	if subs[1].ResourceType != "vNet" || subs[1].ResourceHint != "mig01-vnet" {
		t.Errorf("unexpected extracted resource metadata: %+v", subs[1])
	}
}
