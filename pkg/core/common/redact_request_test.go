package common

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestExtractRequestInfo_RedactsSensitiveAndCaseInsensitiveHeaders(t *testing.T) {
	body := `{"MasterUserPassword":"Password123!","user":"admin"}`
	req := httptest.NewRequest(http.MethodPost, "/beetle/migration/data?password=urlSecret&ok=1", strings.NewReader(body))
	req.Header["authorization"] = []string{"Bearer secret-jwt"}
	req.Header["Cookie"] = []string{"session=top-secret"}
	req.Header["Proxy-Authorization"] = []string{"Basic c2VjcmV0"}
	req.Header["X-Api-Key"] = []string{"ak-live-999"}
	req.Header["X-Auth-Token"] = []string{"tok-xyz"}
	req.Header["accept"] = []string{"application/json"}
	req.Header["user-agent"] = []string{"test-agent"}
	req.Header["X-Empty-Slice"] = []string{}
	req.Header["X-Request-Source"] = []string{"beetle-ux-lab"}
	req.Header["Content-Type"] = []string{"application/json"}

	info := ExtractRequestInfo(req)

	for _, stripped := range []string{
		"authorization",
		"Cookie",
		"Proxy-Authorization",
		"X-Api-Key",
		"X-Auth-Token",
		"accept",
		"user-agent",
		"X-Empty-Slice",
	} {
		if _, exists := info.Header[stripped]; exists {
			t.Errorf("expected header %q to be omitted, got %v", stripped, info.Header[stripped])
		}
	}

	if info.Header["X-Request-Source"] != "beetle-ux-lab" {
		t.Errorf("expected X-Request-Source to be preserved, got %q", info.Header["X-Request-Source"])
	}
	if info.Header["Content-Type"] != "application/json" {
		t.Errorf("expected Content-Type to be preserved, got %q", info.Header["Content-Type"])
	}

	// Verify request body was restored for downstream handlers
	reRead, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("failed to re-read restored request body: %v", err)
	}
	if string(reRead) != body {
		t.Errorf("restored body = %q, want %q", string(reRead), body)
	}
}

func TestSetRequest_RedactsSensitiveDataAcrossFields(t *testing.T) {
	reqID := fmt.Sprintf("req-redact-%d", time.Now().UnixNano())
	defer RemoveRequest(reqID)

	err := SetRequest(reqID, RequestDetails{
		ReqId:     reqID,
		StartTime: time.Now(),
		Status:    RequestStatusError,
		RequestInfo: RequestInfo{
			Method: http.MethodPost,
			URL:    "/beetle/test?password=querySecret123&region=ap-northeast-2",
			Header: map[string]string{"X-Request-Source": "cli"},
			Body: map[string]any{
				"MasterUserPassword": "DbPassword!99",
				"privateKey":         "-----BEGIN RSA PRIVATE KEY-----\nMIIE\n-----END RSA PRIVATE KEY-----\n",
				"normalField":        "keep-me",
			},
		},
		ResponseData: map[string]any{
			"message": "Bearer header-secret-token",
			"status":  "failed",
		},
		ErrorResponse: "upstream failed with password=InlineSecret456",
	})
	if err != nil {
		t.Fatalf("SetRequest failed: %v", err)
	}

	got, ok := GetRequest(reqID)
	if !ok {
		t.Fatalf("GetRequest(%q) returned false", reqID)
	}

	if got.RequestInfo.URL != "/beetle/test?password=********&region=ap-northeast-2" {
		t.Errorf("unexpected redacted URL: %q", got.RequestInfo.URL)
	}
	if got.ErrorResponse != "upstream failed with password=********" {
		t.Errorf("unexpected redacted ErrorResponse: %q", got.ErrorResponse)
	}

	bodyMap, ok := got.RequestInfo.Body.(map[string]any)
	if !ok {
		t.Fatalf("expected RequestInfo.Body map, got %T", got.RequestInfo.Body)
	}
	if bodyMap["MasterUserPassword"] != "********" {
		t.Errorf("MasterUserPassword not redacted: %v", bodyMap["MasterUserPassword"])
	}
	if bodyMap["privateKey"] != "********" {
		t.Errorf("privateKey not redacted: %v", bodyMap["privateKey"])
	}
	if bodyMap["normalField"] != "keep-me" {
		t.Errorf("normalField modified: %v", bodyMap["normalField"])
	}

	respMap, ok := got.ResponseData.(map[string]any)
	if !ok {
		t.Fatalf("expected ResponseData map, got %T", got.ResponseData)
	}
	if respMap["message"] != "Bearer ********" {
		t.Errorf("ResponseData bearer token not redacted: %v", respMap["message"])
	}
	if respMap["status"] != "failed" {
		t.Errorf("ResponseData status modified: %v", respMap["status"])
	}
}

func TestFinishSubRequest_RedactsErrorMessage(t *testing.T) {
	parentID := fmt.Sprintf("sub-redact-%d", time.Now().UnixNano())
	defer DiscardSubRequests(parentID)

	_, seq := StartSubRequest(parentID, http.MethodPost, "/ns/mig01/rdbms", "rdbms", "db-01")
	FinishSubRequest(parentID, seq, http.StatusInternalServerError, "failed with MasterUserPassword: RawSecret123!")

	subs, _ := SnapshotSubRequests(parentID)
	if len(subs) != 1 {
		t.Fatalf("expected 1 live sub-request, got %d", len(subs))
	}
	if subs[0].ErrorMessage != "failed with MasterUserPassword: ********" {
		t.Errorf("unexpected live sub-request ErrorMessage: %q", subs[0].ErrorMessage)
	}
}
