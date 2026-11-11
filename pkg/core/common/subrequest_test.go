package common

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSubRequestTrackerConcurrencyAndSummary(t *testing.T) {
	parentID := fmt.Sprintf("req-race-%d", time.Now().UnixNano())
	defer DiscardSubRequests(parentID)

	const workers = 16
	var wg sync.WaitGroup

	// Poll snapshots concurrently while worker goroutines start and finish sub-requests.
	stopPoll := make(chan struct{})
	pollDone := make(chan struct{})
	go func() {
		defer close(pollDone)
		for {
			select {
			case <-stopPoll:
				return
			default:
				subs, summary := SnapshotSubRequests(parentID)
				if summary != nil && len(subs) != summary.Total {
					t.Errorf("snapshot length mismatch: subs=%d total=%d", len(subs), summary.Total)
				}
			}
		}
	}()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			subID, seq := StartSubRequest(parentID, "POST", "/ns/mig01/securityGroup", "securityGroup", fmt.Sprintf("sg-%d", idx))
			if subID == "" || seq == 0 {
				t.Errorf("expected non-empty subID and seq")
				return
			}
			if idx%4 == 0 {
				FinishSubRequest(parentID, seq, 500, "simulated upstream failure")
			} else {
				FinishSubRequest(parentID, seq, 200, "")
			}
		}(i)
	}

	wg.Wait()
	close(stopPoll)
	<-pollDone

	finalSubs, finalSummary := FinalizeSubRequests(parentID)
	if len(finalSubs) != workers {
		t.Fatalf("expected %d finalized sub-requests, got %d", workers, len(finalSubs))
	}
	if finalSummary == nil || finalSummary.Total != workers || finalSummary.Handling != 0 {
		t.Fatalf("unexpected final summary: %+v", finalSummary)
	}
	if finalSummary.Error != workers/4 || finalSummary.Success != workers-(workers/4) {
		t.Fatalf("unexpected success/error counts: %+v", finalSummary)
	}

	// Verify tracker memory was released upon FinalizeSubRequests.
	afterSubs, afterSummary := SnapshotSubRequests(parentID)
	if len(afterSubs) != 0 || afterSummary != nil {
		t.Fatalf("expected empty tracker after finalize, got len=%d summary=%+v", len(afterSubs), afterSummary)
	}
}

func TestRunAsyncDetachesParentCancellationAndPreservesReqID(t *testing.T) {
	reqID := fmt.Sprintf("async-ctx-%d", time.Now().UnixNano())
	if err := SetRequest(reqID, RequestDetails{
		ReqId:     reqID,
		StartTime: time.Now(),
		Status:    RequestStatusHandling,
	}); err != nil {
		t.Fatalf("SetRequest failed: %v", err)
	}
	defer RemoveRequest(reqID)

	reqCtx, cancel := context.WithCancel(WithRequestID(context.Background(), reqID))
	done := make(chan struct{})

	started := RunAsync(reqCtx, reqID, func(bgCtx context.Context) (string, error) {
		defer close(done)
		// Wait until caller cancels HTTP request context (simulating 202 response completion).
		time.Sleep(30 * time.Millisecond)
		if bgCtx.Err() != nil {
			return "", fmt.Errorf("background context was unexpectedly cancelled: %w", bgCtx.Err())
		}
		gotID := RequestIDFromContext(bgCtx)
		if gotID != reqID {
			return "", fmt.Errorf("expected reqID %q in bgCtx, got %q", reqID, gotID)
		}
		_, seq := StartSubRequest(gotID, "POST", "/ns/mig01/vNet", "vNet", "mig01-vnet")
		FinishSubRequest(gotID, seq, 200, "")
		return "ok", nil
	})
	if !started {
		t.Fatalf("expected RunAsync to start")
	}

	// Simulate Echo cancelling HTTP request context right after writing 202 Accepted.
	cancel()

	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatalf("timed out waiting for RunAsync completion")
	}

	// Allow finalizeRequest to persist into lkvstore.
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		details, ok := GetRequest(reqID)
		if ok && details.Status == RequestStatusSuccess {
			if len(details.SubRequests) != 1 || details.SubRequests[0].ResourceType != "vNet" {
				t.Fatalf("expected finalized vNet sub-request, got %+v", details.SubRequests)
			}
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	details, _ := GetRequest(reqID)
	t.Fatalf("expected Success status, got %+v", details)
}
