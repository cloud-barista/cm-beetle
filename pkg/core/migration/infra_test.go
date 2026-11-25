package migration

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestCheckSSHReadinessWithDetails_InvalidParamsAndNoBusyLoop(t *testing.T) {
	ctx := context.Background()

	if _, err := CheckSSHReadinessWithDetails(ctx, "mig01", "infra-01", 0, time.Second); err == nil || !strings.Contains(err.Error(), "must be positive") {
		t.Fatalf("expected error for non-positive maxWaitTime, got: %v", err)
	}
	if _, err := CheckSSHReadinessWithDetails(ctx, "mig01", "infra-01", time.Second, 0); err == nil || !strings.Contains(err.Error(), "must be positive") {
		t.Fatalf("expected error for non-positive checkInterval, got: %v", err)
	}

	canceledCtx, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := CheckSSHReadinessWithDetails(canceledCtx, "mig01", "infra-01", 100*time.Millisecond, 50*time.Millisecond); err == nil {
		t.Fatalf("expected canceled context to return error")
	}
}
