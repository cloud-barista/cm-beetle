/*
Copyright 2024 The Cloud-Barista Authors.
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

package tbclient

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

type sessionAttemptKey struct{}

// sessionSubRequestAttempt tracks 429 retry state for a single Session request.
type sessionSubRequestAttempt struct {
	mu             sync.Mutex
	seq            uint64
	attempt        int
	lastStatusCode int
	lastMethod     string
	lastPath       string
}

func withSessionAttempt(ctx context.Context, state *sessionSubRequestAttempt) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if state == nil {
		state = &sessionSubRequestAttempt{}
	}
	return context.WithValue(ctx, sessionAttemptKey{}, state)
}

func getSessionAttempt(ctx context.Context) *sessionSubRequestAttempt {
	if ctx == nil {
		return nil
	}
	state, _ := ctx.Value(sessionAttemptKey{}).(*sessionSubRequestAttempt)
	return state
}

// beginSubRequestTrace stamps correlation headers and starts sub-request tracking when applicable.
func beginSubRequestTrace(req *http.Request, bodyBytes []byte) func(statusCode int, respBody []byte, callErr error) {
	noop := func(int, []byte, error) {}
	parentReqID := common.RequestIDFromContext(req.Context())
	if parentReqID == "" {
		return noop
	}

	if req.Header.Get("X-Request-Source") == "" {
		req.Header.Set("X-Request-Source", common.FormatTumblebugSource(parentReqID))
	}
	if !shouldTrackSubRequest(req.Method, req.URL.Path) {
		return noop
	}

	targetPath, resourceType, resourceHint := extractSubRequestTarget(req.URL.Path, bodyBytes)
	attemptTracker := getSessionAttempt(req.Context())

	var subReqID string
	var seq uint64
	if attemptTracker != nil {
		attemptTracker.mu.Lock()
		is429Retry := attemptTracker.seq > 0 &&
			attemptTracker.lastStatusCode == http.StatusTooManyRequests &&
			attemptTracker.lastMethod == req.Method &&
			attemptTracker.lastPath == req.URL.Path &&
			attemptTracker.attempt <= retryOn429Count
		if is429Retry {
			attemptTracker.attempt++
			subReqID = common.RestartSubRequestAttempt(parentReqID, attemptTracker.seq, attemptTracker.attempt)
			seq = attemptTracker.seq
		} else {
			subReqID, seq = common.StartSubRequest(parentReqID, req.Method, targetPath, resourceType, resourceHint)
			attemptTracker.seq = seq
			attemptTracker.attempt = 1
			attemptTracker.lastStatusCode = 0
			attemptTracker.lastMethod = req.Method
			attemptTracker.lastPath = req.URL.Path
		}
		attemptTracker.mu.Unlock()
	} else {
		subReqID, seq = common.StartSubRequest(parentReqID, req.Method, targetPath, resourceType, resourceHint)
	}

	if subReqID == "" || seq == 0 {
		return noop
	}
	req.Header.Set("X-Request-Id", subReqID)

	return func(statusCode int, respBody []byte, callErr error) {
		if attemptTracker != nil {
			attemptTracker.mu.Lock()
			attemptTracker.lastStatusCode = statusCode
			attemptTracker.mu.Unlock()
		}
		errMsg := ""
		if callErr != nil {
			errMsg = callErr.Error()
		} else if statusCode >= http.StatusBadRequest {
			errMsg = extractUpstreamErrorSummary(respBody, http.StatusText(statusCode))
		}
		common.FinishSubRequest(parentReqID, seq, statusCode, errMsg)
	}
}

// shouldTrackSubRequest returns true for mutating or recommendation calls tracked in Tumblebug.
func shouldTrackSubRequest(method, urlPath string) bool {
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodDelete:
	default:
		return false
	}
	normalized := strings.TrimPrefix(urlPath, "/tumblebug")
	if normalized == "/readyz" || normalized == "/request" || strings.HasPrefix(normalized, "/request/") || normalized == "/requests" {
		return false
	}
	return true
}

// extractSubRequestTarget derives normalized path, resource category, and name or ID hint.
func extractSubRequestTarget(urlPath string, bodyBytes []byte) (string, string, string) {
	targetPath := strings.TrimPrefix(urlPath, "/tumblebug")
	if targetPath == "" {
		targetPath = "/"
	}

	var hint string
	if len(bodyBytes) > 0 {
		var meta struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		}
		if err := json.Unmarshal(bodyBytes, &meta); err == nil {
			if strings.TrimSpace(meta.Name) != "" {
				hint = strings.TrimSpace(meta.Name)
			} else if strings.TrimSpace(meta.Id) != "" {
				hint = strings.TrimSpace(meta.Id)
			}
		}
	}

	parts := strings.Split(strings.Trim(targetPath, "/"), "/")
	resourceType := classifySubRequestResource(parts)
	if hint == "" {
		hint = extractPathResourceHint(parts)
	}
	return targetPath, resourceType, hint
}

func classifySubRequestResource(parts []string) string {
	if len(parts) == 0 || parts[0] == "" {
		return ""
	}
	if parts[0] == "recommendSpec" {
		return "recommendSpec"
	}
	if parts[0] == "k8sClusterRecommendNode" {
		return "k8sNodeSpec"
	}
	if parts[0] == "reviewSpecImagePair" {
		return "specImageReview"
	}
	if parts[0] == "ns" && len(parts) >= 3 {
		switch parts[2] {
		case "vNet":
			return "vNet"
		case "sshKey":
			return "sshKey"
		case "securityGroup":
			return "securityGroup"
		case "infra", "infraDynamic":
			if len(parts) >= 5 && parts[4] == "nlb" {
				return "nlb"
			}
			return "infra"
		case "nlb":
			return "nlb"
		case "cmd":
			return "remoteCommand"
		case "k8sCluster":
			for _, seg := range parts[3:] {
				if seg == "k8sNodeGroup" {
					return "k8sNodeGroup"
				}
			}
			return "k8sCluster"
		case "os", "objectStorage":
			return "objectStorage"
		case "resources":
			if len(parts) >= 4 {
				return parts[3]
			}
		default:
			return parts[2]
		}
	}
	return parts[len(parts)-1]
}

func extractPathResourceHint(parts []string) string {
	if len(parts) >= 4 && parts[0] == "ns" {
		if parts[2] == "resources" && len(parts) >= 5 {
			return parts[4]
		}
		if parts[2] == "k8sCluster" && len(parts) >= 6 && parts[4] == "k8sNodeGroup" {
			return parts[5]
		}
		if parts[2] != "resources" && parts[2] != "cmd" {
			return parts[3]
		}
	}
	return ""
}

func extractUpstreamErrorSummary(respBody []byte, fallbackStatus string) string {
	if len(respBody) > 0 {
		var msgObj struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}
		if err := json.Unmarshal(respBody, &msgObj); err == nil {
			if strings.TrimSpace(msgObj.Error) != "" {
				return truncateString(strings.TrimSpace(msgObj.Error), 512)
			}
			if strings.TrimSpace(msgObj.Message) != "" {
				return truncateString(strings.TrimSpace(msgObj.Message), 512)
			}
		}
		trimmed := strings.TrimSpace(string(respBody))
		if trimmed != "" {
			return truncateString(trimmed, 512)
		}
	}
	return fallbackStatus
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
