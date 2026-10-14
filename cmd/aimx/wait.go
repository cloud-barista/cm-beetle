package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/go-resty/resty/v2"
)

// ansiCursorResetRegex matches cursor-up or column-reset sequences that overwrite previous text.
var ansiCursorResetRegex = regexp.MustCompile(`\x1b\[[0-9]*[AGH]`)

// ansiEscapeRegex matches standard ANSI CSI/OSC terminal control sequences.
var ansiEscapeRegex = regexp.MustCompile(`\x1b\[[0-9;?]*[a-zA-Z]|\x1b\][^\x07]*\x07`)

// waitWithCountdown pauses execution for the specified seconds to prevent cloud/container timing races.
func waitWithCountdown(seconds int, reason string) {
	if seconds <= 0 {
		return
	}
	fmt.Printf("⏱️  [Stabilization Wait] %s (%ds)...\n", reason, seconds)
	remaining := seconds
	for remaining > 0 {
		step := 5
		if remaining < step {
			step = remaining
		}
		time.Sleep(time.Duration(step) * time.Second)
		remaining -= step
		if remaining > 0 {
			fmt.Printf("   ... %ds remaining (%s)\n", remaining, reason)
		}
	}
}

// execRemoteCmd executes remote shell commands via CB-Tumblebug's async API and streams logs live via SSE.
func execRemoteCmd(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, commands []string) (string, error) {
	return execRemoteCmdSSE(tb, tbEndpoint, nsID, infraID, nodeID, commands, true)
}

// execRemoteCmdQuiet executes remote shell commands via CB-Tumblebug's async SSE API without printing raw lines to stdout.
func execRemoteCmdQuiet(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, commands []string) (string, error) {
	return execRemoteCmdSSE(tb, tbEndpoint, nsID, infraID, nodeID, commands, false)
}

// execRemoteCmdSSE dispatches an async remote command to CB-Tumblebug and subscribes to its SSE event stream.
func execRemoteCmdSSE(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, commands []string, printLive bool) (string, error) {
	baseURL := strings.TrimRight(tbEndpoint, "/")
	cmdURL := fmt.Sprintf("%s/tumblebug/ns/%s/cmd/infra/%s", baseURL, nsID, infraID)
	reqBody := tbmodel.InfraCmdReq{
		Command:        commands,
		TimeoutMinutes: 20,
	}
	var asyncResp map[string]string

	req := tb.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("async", "true").
		SetBody(&reqBody).
		SetResult(&asyncResp)
	if nodeID != "" {
		req.SetQueryParam("nodeId", nodeID)
	}

	resp, err := req.Post(cmdURL)
	if err != nil {
		return "", fmt.Errorf("remote command HTTP error on infra %s: %w", infraID, err)
	}
	if resp.IsError() {
		return "", fmt.Errorf("remote command API status %d on infra %s: %s", resp.StatusCode(), infraID, resp.String())
	}

	xRequestID := strings.TrimSpace(asyncResp["xRequestId"])
	if xRequestID == "" {
		xRequestID = strings.TrimSpace(resp.Header().Get("X-Request-Id"))
	}
	if xRequestID == "" {
		return "", fmt.Errorf("remote command on infra %s returned empty xRequestId", infraID)
	}

	streamURL := fmt.Sprintf("%s/tumblebug/ns/%s/stream/cmd/infra/%s", baseURL, nsID, infraID)
	streamResp, err := tb.R().
		SetDoNotParseResponse(true).
		SetHeader("Accept", "text/event-stream").
		SetQueryParam("xRequestId", xRequestID).
		Get(streamURL)
	if err != nil {
		return "", fmt.Errorf("SSE stream connection error on infra %s (xRequestId=%s): %w", infraID, xRequestID, err)
	}
	rawBody := streamResp.RawBody()
	if rawBody != nil {
		defer rawBody.Close()
	}
	if streamResp.IsError() {
		return "", fmt.Errorf("SSE stream status %d on infra %s (xRequestId=%s)", streamResp.StatusCode(), infraID, xRequestID)
	}

	return consumeCommandSSEStream(rawBody, infraID, printLive)
}

// consumeCommandSSEStream reads SSE events until CommandDone and returns accumulated stdout or any execution error.
func consumeCommandSSEStream(r io.Reader, infraID string, printLive bool) (string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	var stdoutBuilder strings.Builder
	var stderrBuilder strings.Builder
	var statusErrors []string
	var lastPrinted string
	inProgressLine := false

	clearProgressLine := func() {
		if inProgressLine {
			fmt.Print("\r\x1b[K")
			inProgressLine = false
		}
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ":") || !strings.HasPrefix(line, "data:") {
			continue
		}
		payload := strings.TrimSpace(strings.TrimPrefix(line, "data:"))
		if payload == "" {
			continue
		}

		var evt tbmodel.CommandStreamEvent
		if err := json.Unmarshal([]byte(payload), &evt); err != nil {
			continue
		}

		switch evt.Type {
		case tbmodel.EventCommandLog:
			if evt.Log == nil {
				continue
			}
			isTBDiag := strings.HasPrefix(strings.TrimSpace(evt.Log.Line), "[TB]")
			if !isTBDiag {
				if evt.Log.Stream == "stderr" {
					stderrBuilder.WriteString(evt.Log.Line)
					stderrBuilder.WriteString("\n")
				} else {
					stdoutBuilder.WriteString(evt.Log.Line)
					stdoutBuilder.WriteString("\n")
				}
			}
			if !printLive {
				continue
			}

			cleanLine := sanitizeStreamLogLine(evt.Log.Line)
			if cleanLine == "" || cleanLine == lastPrinted {
				continue
			}
			lastPrinted = cleanLine

			if isTransientProgressLine(cleanLine) {
				fmt.Printf("\r   ⏳ %s\x1b[K", truncateRunes(cleanLine, 96))
				inProgressLine = true
			} else {
				clearProgressLine()
				fmt.Printf("   │ %s\n", cleanLine)
			}

		case tbmodel.EventCommandStatus:
			if evt.Status == nil {
				continue
			}
			st := evt.Status.Status
			if st == tbmodel.CommandStatusFailed || st == tbmodel.CommandStatusTimeout || st == tbmodel.CommandStatusCancelled {
				errDetail := firstNonEmpty(evt.Status.ErrorMessage, evt.Status.ResultSummary, string(st))
				statusErrors = append(statusErrors, fmt.Sprintf("node %s: %s", evt.NodeId, errDetail))
			}

		case tbmodel.EventCommandDone:
			if evt.Summary != nil {
				if evt.Summary.Error != "" {
					statusErrors = append(statusErrors, evt.Summary.Error)
				} else if evt.Summary.FailedNodes > 0 && len(statusErrors) == 0 {
					statusErrors = append(statusErrors, fmt.Sprintf("%d/%d node(s) failed on infra %s", evt.Summary.FailedNodes, evt.Summary.TotalNodes, infraID))
				}
			}
			goto streamDone
		}
	}

streamDone:
	if printLive {
		clearProgressLine()
	}
	if err := scanner.Err(); err != nil {
		return strings.TrimSpace(stdoutBuilder.String()), fmt.Errorf("SSE stream read error on infra %s: %w", infraID, err)
	}

	combinedOut := strings.TrimSpace(stdoutBuilder.String())
	if len(statusErrors) > 0 {
		stderrStr := strings.TrimSpace(stderrBuilder.String())
		if stderrStr != "" {
			return combinedOut, fmt.Errorf("remote command failed on %s: %s (stderr: %s)", infraID, strings.Join(statusErrors, "; "), stderrStr)
		}
		return combinedOut, fmt.Errorf("remote command failed on %s: %s", infraID, strings.Join(statusErrors, "; "))
	}
	return combinedOut, nil
}

// sanitizeStreamLogLine strips ANSI escape sequences and resolves carriage-return overwrites to the latest frame.
func sanitizeStreamLogLine(raw string) string {
	normalized := ansiCursorResetRegex.ReplaceAllString(raw, "\r")
	stripped := ansiEscapeRegex.ReplaceAllString(normalized, "")
	segments := strings.Split(stripped, "\r")
	for i := len(segments) - 1; i >= 0; i-- {
		seg := strings.TrimSpace(segments[i])
		if seg != "" {
			return seg
		}
	}
	return ""
}

// isTransientProgressLine returns true if a log line represents an intermediate download/extract progress update.
func isTransientProgressLine(line string) bool {
	lower := strings.ToLower(line)
	if strings.HasPrefix(line, "[TB]") || strings.ContainsAny(line, "▕▏⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏") {
		return true
	}
	if strings.HasPrefix(lower, "pulling ") || strings.HasPrefix(lower, "verifying sha256") {
		return true
	}
	for _, kw := range []string{"downloading", "extracting", "waiting", "pulling fs layer", "verifying checksum", "download complete", "pull complete"} {
		if strings.Contains(lower, kw) {
			return true
		}
	}
	return false
}

// truncateRunes limits a string to maxRunes characters so single-line progress updates do not wrap.
func truncateRunes(s string, maxRunes int) string {
	runes := []rune(s)
	if len(runes) <= maxRunes {
		return s
	}
	return string(runes[:maxRunes-3]) + "..."
}

// pollGpuVmReady retries SSH and nvidia-smi checks until the GPU VM is responsive.
func pollGpuVmReady(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, timing TimingConfig) (string, error) {
	maxAttempts := timing.SshRetryMaxAttempts
	if maxAttempts <= 0 {
		maxAttempts = 12
	}
	interval := timing.SshRetryIntervalSeconds
	if interval <= 0 {
		interval = 10
	}

	fmt.Printf("🔍 Polling SSH & NVIDIA GPU readiness on infra '%s' (max %d attempts, %ds interval)...\n", infraID, maxAttempts, interval)
	checkCmd := []string{"nvidia-smi --query-gpu=name,memory.total,driver_version --format=csv,noheader"}

	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		out, err := execRemoteCmdQuiet(tb, tbEndpoint, nsID, infraID, nodeID, checkCmd)
		if err == nil && strings.TrimSpace(out) != "" {
			fmt.Printf("✅ GPU VM '%s' is ready (attempt %d/%d): %s\n", infraID, attempt, maxAttempts, strings.Split(out, "\n")[0])
			return out, nil
		}
		lastErr = err
		fmt.Printf("   ⏳ Attempt %d/%d: GPU VM not ready yet (%v); waiting %ds...\n", attempt, maxAttempts, err, interval)
		time.Sleep(time.Duration(interval) * time.Second)
	}
	return "", fmt.Errorf("GPU VM '%s' did not become ready after %d attempts: %w", infraID, maxAttempts, lastErr)
}

// pollServiceEndpointsReady waits until both Ollama and Open WebUI respond on the GPU VM.
func pollServiceEndpointsReady(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, ollamaPort, webuiPort int, timing TimingConfig) error {
	maxAttempts := timing.SshRetryMaxAttempts
	if maxAttempts <= 0 {
		maxAttempts = 12
	}
	interval := timing.SshRetryIntervalSeconds
	if interval <= 0 {
		interval = 10
	}

	fmt.Printf("🔍 Polling Ollama (:%d) & Open WebUI (:%d) readiness on '%s'...\n", ollamaPort, webuiPort, infraID)
	checkCmd := fmt.Sprintf(
		"curl -sf http://localhost:%d/api/tags >/dev/null && code=$(curl -s -o /dev/null -w '%%{http_code}' http://localhost:%d) && [ \"$code\" -ge 200 ] && [ \"$code\" -lt 400 ] && echo \"SERVICES_OK:$code\"",
		ollamaPort, webuiPort,
	)

	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		out, err := execRemoteCmdQuiet(tb, tbEndpoint, nsID, infraID, nodeID, []string{checkCmd})
		if err == nil && strings.Contains(out, "SERVICES_OK:") {
			fmt.Printf("✅ AI services on '%s' are healthy (%s)\n", infraID, strings.TrimSpace(out))
			return nil
		}
		lastErr = err
		fmt.Printf("   ⏳ Attempt %d/%d: AI services initializing; waiting %ds...\n", attempt, maxAttempts, interval)
		time.Sleep(time.Duration(interval) * time.Second)
	}
	return fmt.Errorf("AI services on '%s' did not become healthy after %d attempts: %v", infraID, maxAttempts, lastErr)
}
