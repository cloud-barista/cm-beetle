/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TumblebugSourcePrefix identifies CM-Beetle as the upstream caller in X-Request-Source.
const TumblebugSourcePrefix = "cm-beetle:"

type requestIDContextKey struct{}

// WithRequestID attaches the Beetle request ID to ctx for downstream correlation.
func WithRequestID(ctx context.Context, reqID string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if reqID == "" {
		return ctx
	}
	return context.WithValue(ctx, requestIDContextKey{}, reqID)
}

// RequestIDFromContext extracts the Beetle request ID from ctx, or returns empty string.
func RequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	v, _ := ctx.Value(requestIDContextKey{}).(string)
	return v
}

// FormatTumblebugSource builds the canonical X-Request-Source value ("cm-beetle:{reqId}").
func FormatTumblebugSource(reqID string) string {
	if reqID == "" {
		return ""
	}
	return TumblebugSourcePrefix + reqID
}

// SubRequest records one outbound state-changing or recommendation call to CB-Tumblebug.
type SubRequest struct {
	Seq              uint64    `json:"seq"`
	SubReqId         string    `json:"subReqId"`
	Method           string    `json:"method"`
	TargetPath       string    `json:"targetPath"`
	ResourceType     string    `json:"resourceType,omitempty"`
	ResourceHint     string    `json:"resourceHint,omitempty"`
	Status           string    `json:"status"`
	StatusCode       int       `json:"statusCode,omitempty"`
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime,omitempty"`
	DurationMs       int64     `json:"durationMs,omitempty"`
	ErrorMessage     string    `json:"errorMessage,omitempty"`
	UpstreamStatus   string    `json:"upstreamStatus,omitempty"`
	UpstreamProgress any       `json:"upstreamProgress,omitempty"`
	UpstreamError    string    `json:"upstreamError,omitempty"`
}

// SubRequestSummary aggregates sub-request counts and active step description.
type SubRequestSummary struct {
	Total         int    `json:"total"`
	Handling      int    `json:"handling"`
	Success       int    `json:"success"`
	Error         int    `json:"error"`
	CurrentAction string `json:"currentAction,omitempty"`
}

type trackedParent struct {
	createdAt time.Time
	nextSeq   uint64
	items     []SubRequest
}

type subRequestTracker struct {
	mu      sync.Mutex
	entries map[string]*trackedParent
}

var defaultSubRequestTracker = &subRequestTracker{
	entries: make(map[string]*trackedParent),
}

// StartSubRequest allocates a unique "{reqId}-{seq}" sub-request ID and records Handling state.
func StartSubRequest(parentReqID, method, targetPath, resourceType, resourceHint string) (string, uint64) {
	if parentReqID == "" {
		return "", 0
	}
	defaultSubRequestTracker.mu.Lock()
	defer defaultSubRequestTracker.mu.Unlock()

	parent, ok := defaultSubRequestTracker.entries[parentReqID]
	if !ok {
		parent = &trackedParent{createdAt: time.Now()}
		defaultSubRequestTracker.entries[parentReqID] = parent
	}
	parent.nextSeq++
	seq := parent.nextSeq
	subReqID := fmt.Sprintf("%s-%d", parentReqID, seq)
	parent.items = append(parent.items, SubRequest{
		Seq:          seq,
		SubReqId:     subReqID,
		Method:       method,
		TargetPath:   targetPath,
		ResourceType: resourceType,
		ResourceHint: resourceHint,
		Status:       RequestStatusHandling,
		StartTime:    time.Now(),
	})
	return subReqID, seq
}

// FinishSubRequest updates a recorded sub-request with its final status, HTTP code, and latency.
func FinishSubRequest(parentReqID string, seq uint64, statusCode int, errMsg string) {
	if parentReqID == "" || seq == 0 {
		return
	}
	defaultSubRequestTracker.mu.Lock()
	defer defaultSubRequestTracker.mu.Unlock()

	parent, ok := defaultSubRequestTracker.entries[parentReqID]
	if !ok {
		return
	}
	for i := len(parent.items) - 1; i >= 0; i-- {
		if parent.items[i].Seq == seq {
			now := time.Now()
			parent.items[i].EndTime = now
			parent.items[i].DurationMs = now.Sub(parent.items[i].StartTime).Milliseconds()
			parent.items[i].StatusCode = statusCode
			if errMsg == "" && statusCode > 0 && statusCode < 400 {
				parent.items[i].Status = RequestStatusSuccess
			} else {
				parent.items[i].Status = RequestStatusError
				parent.items[i].ErrorMessage = errMsg
			}
			return
		}
	}
}

// SnapshotSubRequests returns a race-free copy of live sub-requests and summary for polling.
func SnapshotSubRequests(parentReqID string) ([]SubRequest, *SubRequestSummary) {
	if parentReqID == "" {
		return nil, nil
	}
	defaultSubRequestTracker.mu.Lock()
	parent, ok := defaultSubRequestTracker.entries[parentReqID]
	if !ok || len(parent.items) == 0 {
		defaultSubRequestTracker.mu.Unlock()
		return nil, nil
	}
	snapshot := append([]SubRequest(nil), parent.items...)
	defaultSubRequestTracker.mu.Unlock()

	return snapshot, SummarizeSubRequests(snapshot)
}

// FinalizeSubRequests detaches and returns the final sub-requests for persistence in lkvstore.
func FinalizeSubRequests(parentReqID string) ([]SubRequest, *SubRequestSummary) {
	if parentReqID == "" {
		return nil, nil
	}
	defaultSubRequestTracker.mu.Lock()
	parent, ok := defaultSubRequestTracker.entries[parentReqID]
	if !ok || len(parent.items) == 0 {
		delete(defaultSubRequestTracker.entries, parentReqID)
		defaultSubRequestTracker.mu.Unlock()
		return nil, nil
	}
	snapshot := append([]SubRequest(nil), parent.items...)
	delete(defaultSubRequestTracker.entries, parentReqID)
	defaultSubRequestTracker.mu.Unlock()

	return snapshot, SummarizeSubRequests(snapshot)
}

// DiscardSubRequests removes in-memory sub-request entries for a deleted request ID.
func DiscardSubRequests(parentReqID string) {
	if parentReqID == "" {
		return
	}
	defaultSubRequestTracker.mu.Lock()
	delete(defaultSubRequestTracker.entries, parentReqID)
	defaultSubRequestTracker.mu.Unlock()
}

// DiscardAllSubRequests clears all in-memory sub-request entries.
func DiscardAllSubRequests() {
	defaultSubRequestTracker.mu.Lock()
	defaultSubRequestTracker.entries = make(map[string]*trackedParent)
	defaultSubRequestTracker.mu.Unlock()
}

// SummarizeSubRequests computes counts and active action label from a sub-request slice.
func SummarizeSubRequests(items []SubRequest) *SubRequestSummary {
	if len(items) == 0 {
		return nil
	}
	summary := &SubRequestSummary{
		Total: len(items),
	}
	for _, item := range items {
		switch item.Status {
		case RequestStatusHandling:
			summary.Handling++
			summary.CurrentAction = formatSubRequestAction(item)
		case RequestStatusSuccess:
			summary.Success++
		case RequestStatusError:
			summary.Error++
		}
	}
	return summary
}

func formatSubRequestAction(item SubRequest) string {
	target := item.TargetPath
	if item.ResourceType != "" {
		target = item.ResourceType
		if item.ResourceHint != "" {
			target = fmt.Sprintf("%s (%s)", item.ResourceType, item.ResourceHint)
		}
	}
	return fmt.Sprintf("%s %s", item.Method, target)
}
