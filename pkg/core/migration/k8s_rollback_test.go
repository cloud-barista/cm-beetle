package migration

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

func TestK8sPrereqsRollbackDetachesCanceledParentContext(t *testing.T) {
	var mu sync.Mutex
	var deletedPaths []string
	var seenSources []string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		deletedPaths = append(deletedPaths, r.Method+" "+r.URL.Path)
		seenSources = append(seenSources, r.Header.Get("X-Request-Source"))
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"deleted"}`))
	}))
	defer srv.Close()

	restore := tbclient.SetClient(tbclient.NewClient(tbclient.ApiConfig{
		RestUrl:  srv.URL,
		Username: "default",
		Password: "default",
		Timeout:  5 * time.Second,
	}))
	defer restore()

	parentReqID := fmt.Sprintf("k8s-rb-%d", time.Now().UnixNano())
	defer common.DiscardSubRequests(parentReqID)

	callerCtx, cancel := context.WithCancel(common.WithRequestID(context.Background(), parentReqID))
	cancel()

	prereqs := &k8sClusterPrereqs{
		nsId:     "mig01",
		vNetId:   "mig01-vnet",
		sshKeyId: "mig01-key",
		sgIds:    []string{"mig01-sg-1", "mig01-sg-2"},
	}
	prereqs.rollback(callerCtx)

	mu.Lock()
	defer mu.Unlock()

	wantPaths := []string{
		"DELETE /ns/mig01/resources/securityGroup/mig01-sg-2",
		"DELETE /ns/mig01/resources/securityGroup/mig01-sg-1",
		"DELETE /ns/mig01/resources/sshKey/mig01-key",
		"DELETE /ns/mig01/resources/vNet/mig01-vnet",
	}
	if len(deletedPaths) != len(wantPaths) {
		t.Fatalf("expected %d rollback DELETE calls despite canceled caller ctx, got %v", len(wantPaths), deletedPaths)
	}
	for i, want := range wantPaths {
		if deletedPaths[i] != want {
			t.Errorf("rollback step %d: want %q, got %q", i, want, deletedPaths[i])
		}
	}
	for _, src := range seenSources {
		if src != "cm-beetle:"+parentReqID {
			t.Errorf("expected X-Request-Source=%q on rollback call, got %q", "cm-beetle:"+parentReqID, src)
		}
	}

	subs, summary := common.FinalizeSubRequests(parentReqID)
	if len(subs) != 4 || summary == nil || summary.Success != 4 || summary.Error != 0 {
		t.Fatalf("expected 4 succeeded rollback sub-requests, got subs=%+v summary=%+v", subs, summary)
	}
}
