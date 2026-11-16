package tbclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

func TestRequestScopedCache_AllowlistBypassErrorsAndConcurrency(t *testing.T) {
	var (
		regionCalls        atomic.Int32
		specSearchCalls    atomic.Int32
		imageSearchCalls   atomic.Int32
		reviewPairCalls    atomic.Int32
		k8sSpecCalls       atomic.Int32
		k8sProfileCalls    atomic.Int32
		rdbmsSupportCalls  atomic.Int32
		rdbmsValidateCalls atomic.Int32
		createVNetCalls    atomic.Int32
		readInfraCalls     atomic.Int32
		readyzCalls        atomic.Int32
		failNextSpec       atomic.Bool
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/provider/aws/region/ap-northeast-2":
			regionCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"regionId":"ap-northeast-2","regionName":"ap-northeast-2"}`))

		case r.Method == http.MethodGet && r.URL.Path == "/k8sClusterInfo":
			k8sProfileCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"k8s_cluster":{"aws":{"nodeGroupsOnCreation":false}}}`))

		case r.Method == http.MethodGet && r.URL.Path == "/rdbms/support":
			rdbmsSupportCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{}`))

		case r.Method == http.MethodGet && r.URL.Path == "/ns/mig01/infra/infra01":
			n := readInfraCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprintf(w, `{"id":"infra01","status":"Running-%d"}`, n)

		case r.Method == http.MethodGet && r.URL.Path == "/readyz":
			readyzCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"ready":true,"initialized":true}`))

		case r.Method == http.MethodPost && (r.URL.Path == "/recommendSpec" || r.URL.Path == "//tumblebug/recommendSpec"):
			specSearchCalls.Add(1)
			if failNextSpec.Swap(false) {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"message":"temporary failure"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`[{"id":"aws+ap-northeast-2+t3.medium","cspSpecName":"t3.medium"}]`))

		case r.Method == http.MethodPost && r.URL.Path == "/ns/system/resources/searchImage":
			imageSearchCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"imageCount":1,"imageList":[{"id":"aws+ap-northeast-2+ubuntu-22.04"}]}`))

		case r.Method == http.MethodPost && r.URL.Path == "/specImagePairReview":
			reviewPairCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"isValid":true,"status":"OK"}`))

		case r.Method == http.MethodPost && r.URL.Path == "/k8sClusterRecommendNode":
			k8sSpecCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`[{"id":"aws+ap-northeast-2+m5.large"}]`))

		case r.Method == http.MethodPost && r.URL.Path == "/ns/mig01/resources/rdbms/validate":
			rdbmsValidateCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{}`))

		case r.Method == http.MethodPost && r.URL.Path == "/ns/mig01/resources/vNet":
			createVNetCalls.Add(1)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"id":"vnet-01","name":"vnet-01"}`))

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()

	cli := NewClient(ApiConfig{
		RestUrl:       srv.URL,
		Username:      "default",
		Password:      "default",
		Timeout:       5 * time.Second,
		CacheDuration: 10 * time.Millisecond,
	})

	ctx := WithRequestCache(context.Background())

	// 1. Transient non-2xx errors must never be cached.
	failNextSpec.Store(true)
	if _, err := cli.NewSession().SetContext(ctx).InfraRecommendSpec(`{"limit":5}`); err == nil {
		t.Fatalf("expected initial 500 error")
	}
	specs, err := cli.NewSession().SetContext(ctx).InfraRecommendSpec(`{"limit":5}`)
	if err != nil || len(specs) != 1 {
		t.Fatalf("expected retry after non-2xx failure to succeed, got len=%d err=%v", len(specs), err)
	}
	if got := specSearchCalls.Load(); got != 2 {
		t.Fatalf("expected 2 recommendSpec upstream calls after 500+200, got %d", got)
	}

	// 2. Identical query POST (including differently formatted JSON) hits request cache.
	for i := 0; i < 4; i++ {
		if _, err := cli.NewSession().SetContext(ctx).InfraRecommendSpec(`{"limit": 5}`); err != nil {
			t.Fatalf("cached InfraRecommendSpec failed: %v", err)
		}
	}
	if got := specSearchCalls.Load(); got != 2 {
		t.Fatalf("expected cached recommendSpec hits to avoid upstream calls, got %d", got)
	}

	// 3. Different body creates a separate cache entry, and double-slash /tumblebug prefix still caches.
	if _, err := cli.NewSession().SetContext(ctx).InfraRecommendSpec(`{"limit":10}`); err != nil {
		t.Fatalf("different body InfraRecommendSpec failed: %v", err)
	}
	if got := specSearchCalls.Load(); got != 3 {
		t.Fatalf("expected distinct recommendSpec body to miss cache, got %d", got)
	}
	for i := 0; i < 2; i++ {
		if _, err := cli.NewSession().SetContext(ctx).SetBody(map[string]int{"limit": 12}).Post("//tumblebug/recommendSpec"); err != nil {
			t.Fatalf("double-slash /tumblebug/recommendSpec failed: %v", err)
		}
	}
	if got := specSearchCalls.Load(); got != 4 {
		t.Fatalf("expected double-slash /tumblebug/recommendSpec to cache on repeat, got %d", got)
	}

	// 4. Other allowlisted read GETs and query POSTs are cached across repeats within the request.
	for i := 0; i < 3; i++ {
		if _, err := cli.NewSession().SetContext(ctx).ReadRegionInfo("aws", "ap-northeast-2"); err != nil {
			t.Fatalf("ReadRegionInfo failed: %v", err)
		}
		if _, err := cli.NewSession().SetContext(ctx).SearchVmOsImage("system", tbmodel.SearchImageRequest{OSType: "ubuntu 22.04"}); err != nil {
			t.Fatalf("SearchVmOsImage failed: %v", err)
		}
		if _, err := cli.NewSession().SetContext(ctx).ReviewSpecImagePair(tbmodel.SpecImagePairReviewReq{SpecId: "s1", ImageId: "i1"}); err != nil {
			t.Fatalf("ReviewSpecImagePair failed: %v", err)
		}
		if _, err := cli.NewSession().SetContext(ctx).K8sClusterRecommendNode(tbmodel.RecommendSpecReq{}); err != nil {
			t.Fatalf("K8sClusterRecommendNode failed: %v", err)
		}
		if _, err := cli.NewSession().SetContext(ctx).getK8sClusterDetailExt("aws"); err != nil {
			t.Fatalf("getK8sClusterDetailExt failed: %v", err)
		}
		if _, err := cli.NewSession().SetContext(ctx).GetRDBMSSupport("aws"); err != nil {
			t.Fatalf("GetRDBMSSupport failed: %v", err)
		}
		if _, err := cli.NewSession().SetContext(ctx).SetBody(map[string]string{"name": "db1"}).Post("/ns/mig01/resources/rdbms/validate"); err != nil {
			t.Fatalf("rdbms validate failed: %v", err)
		}
	}
	if regionCalls.Load() != 1 || imageSearchCalls.Load() != 1 || reviewPairCalls.Load() != 1 ||
		k8sSpecCalls.Load() != 1 || k8sProfileCalls.Load() != 1 || rdbmsSupportCalls.Load() != 1 || rdbmsValidateCalls.Load() != 1 {
		t.Fatalf("expected 1 upstream call per allowlisted query, got region=%d img=%d review=%d k8sSpec=%d k8sProf=%d rdbmsSup=%d rdbmsVal=%d",
			regionCalls.Load(), imageSearchCalls.Load(), reviewPairCalls.Load(),
			k8sSpecCalls.Load(), k8sProfileCalls.Load(), rdbmsSupportCalls.Load(), rdbmsValidateCalls.Load())
	}

	// 5. Mutating POSTs, mutable-resource GETs (after short global TTL), readyz, and NoCache bypass request cache.
	for i := 0; i < 2; i++ {
		if i > 0 {
			time.Sleep(25 * time.Millisecond)
		}
		if _, err := cli.NewSession().SetContext(ctx).CreateVNet("mig01", tbmodel.VNetReq{Name: "vnet-01"}); err != nil {
			t.Fatalf("CreateVNet failed: %v", err)
		}
		if _, err := cli.NewSession().SetContext(ctx).SetResult(&tbmodel.InfraInfo{}).Get("/ns/mig01/infra/infra01"); err != nil {
			t.Fatalf("mutable infra GET failed: %v", err)
		}
		if _, err := cli.NewSession().SetContext(ctx).CheckReadyz(); err != nil {
			t.Fatalf("CheckReadyz failed: %v", err)
		}
		// After 25ms (> 10ms global TTL), ReadRegionInfo still hits RequestCache while NoCache forces upstream call.
		if _, err := cli.NewSession().SetContext(ctx).ReadRegionInfo("aws", "ap-northeast-2"); err != nil {
			t.Fatalf("cached ReadRegionInfo failed: %v", err)
		}
		if _, err := cli.NewSession().SetContext(ctx).NoCache().ReadRegionInfo("aws", "ap-northeast-2"); err != nil {
			t.Fatalf("NoCache ReadRegionInfo failed: %v", err)
		}
	}
	if createVNetCalls.Load() != 2 || readInfraCalls.Load() != 2 || readyzCalls.Load() != 2 || regionCalls.Load() != 3 {
		t.Fatalf("expected mutating/polling/NoCache calls to bypass request cache, got createVNet=%d infraGet=%d readyz=%d region=%d",
			createVNetCalls.Load(), readInfraCalls.Load(), readyzCalls.Load(), regionCalls.Load())
	}

	// 6. Another request context has an isolated cache and cannot observe hits from ctx.
	otherCtx := WithRequestCache(context.Background())
	if _, err := cli.NewSession().SetContext(otherCtx).ReadRegionInfo("aws", "ap-northeast-2"); err != nil {
		t.Fatalf("isolated ReadRegionInfo failed: %v", err)
	}
	if regionCalls.Load() != 4 {
		t.Fatalf("expected isolated request cache miss on otherCtx, got regionCalls=%d", regionCalls.Load())
	}

	stats, ok := GetRequestCacheStats(ctx)
	if !ok || stats.Hits == 0 || stats.Misses == 0 || stats.Bypassed == 0 {
		t.Fatalf("expected populated request cache stats, got ok=%v stats=%+v", ok, stats)
	}
}

func TestRequestScopedCache_CanonicalBodyAndConcurrency(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var payload any
		_ = json.NewDecoder(r.Body).Decode(&payload)
		_ = json.NewEncoder(w).Encode([]tbmodel.SpecInfo{{Id: "aws+ap-northeast-2+t3.medium", CspSpecName: "t3.medium"}})
	}))
	defer srv.Close()

	cli := NewClient(ApiConfig{RestUrl: srv.URL, Username: "default", Password: "default"})
	ctx := WithRequestCache(context.Background())

	// Map keys with different insertion formatting should normalize to the same canonical JSON hash.
	if _, err := cli.NewSession().SetContext(ctx).SetBody(json.RawMessage(`{"b":2,"a":1}`)).Post("/recommendSpec"); err != nil {
		t.Fatalf("first POST failed: %v", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var res []tbmodel.SpecInfo
			resp, err := cli.NewSession().
				SetContext(ctx).
				SetBody(json.RawMessage(`{ "a" : 1 , "b" : 2 }`)).
				SetResult(&res).
				Post("/recommendSpec")
			if err != nil || resp.IsError() || len(res) != 1 {
				t.Errorf("concurrent cached POST failed: len=%d err=%v", len(res), err)
			}
		}()
	}
	wg.Wait()

	if calls.Load() != 1 {
		t.Fatalf("expected all 16 formatted/concurrent calls to hit the single cached entry, got %d upstream calls", calls.Load())
	}
}
