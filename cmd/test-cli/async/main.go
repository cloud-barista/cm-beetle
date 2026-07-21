// Package main is the starting point of CM-Beetle Async Response Test CLI.
//
// Exercises the core async flow (Prefer: respond-async -> 202 -> poll GET /request/{reqId})
// end to end: Recommend -> Poll -> Get -> Migrate -> Poll -> Get -> Delete -> Poll -> Get.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
)

// restyNoopLogger silences all Resty log output (e.g. "Basic Auth in HTTP mode" warnings).
type restyNoopLogger struct{}

func (restyNoopLogger) Errorf(_ string, _ ...interface{}) {}
func (restyNoopLogger) Warnf(_ string, _ ...interface{})  {}
func (restyNoopLogger) Debugf(_ string, _ ...interface{}) {}

// TestConfig holds the test configuration from test-config.yaml.
type TestConfig struct {
	Beetle struct {
		Endpoint        string `yaml:"endpoint"`
		NamespaceID     string `yaml:"namespaceId"`
		RequestBodyFile string `yaml:"requestBodyFile"`
		AuthConfigFile  string `yaml:"authConfigFile"`
	} `yaml:"beetle"`
	Poll struct {
		IntervalSec int `yaml:"intervalSec"` // fixed poll interval
		TimeoutSec  int `yaml:"timeoutSec"`
	} `yaml:"poll"`
}

// AuthConfig holds Beetle API credentials.
type AuthConfig struct {
	BeetleApiUsername string `json:"beetleApiUsername"`
	BeetleApiPassword string `json:"beetleApiPassword"`
}

// ApiResponse mirrors model.ApiResponse[T] — only the fields this CLI needs.
type ApiResponse struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message,omitempty"`
	Error   string          `json:"error,omitempty"`
}

// AsyncJobData mirrors model.AsyncJobResponse.
type AsyncJobData struct {
	ReqID     string `json:"reqId"`
	Status    string `json:"status"`
	StatusURL string `json:"statusUrl"`
}

// RequestDetails mirrors the fields of common.RequestDetails this CLI reads.
type RequestDetails struct {
	Status        string          `json:"status"`
	ResponseData  json.RawMessage `json:"responseData"`
	ErrorResponse string          `json:"errorResponse"`
}

var configFile = flag.String("config", "testconf/test-config.yaml", "Path to config file")

func init() {
	config.Init()
	l := logger.NewLogger(logger.Config{
		LogLevel:    config.Beetle.LogLevel,
		LogWriter:   config.Beetle.LogWriter,
		LogFilePath: config.Beetle.LogFile.Path,
		MaxSize:     config.Beetle.LogFile.MaxSize,
		MaxBackups:  config.Beetle.LogFile.MaxBackups,
		MaxAge:      config.Beetle.LogFile.MaxAge,
		Compress:    config.Beetle.LogFile.Compress,
	})
	log.Logger = *l
}

func main() {
	flag.Parse()

	cfg, err := loadConfig(*configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	auth, err := loadAuthConfig(cfg.Beetle.AuthConfigFile)
	if err != nil {
		fmt.Printf("(no auth config loaded: %v — proceeding without auth)\n", err)
	}

	client := resty.New().SetTimeout(30 * time.Second).SetLogger(restyNoopLogger{})
	if auth.BeetleApiUsername != "" {
		client.SetBasicAuth(auth.BeetleApiUsername, auth.BeetleApiPassword)
	}

	intervalSec := cfg.Poll.IntervalSec
	if intervalSec <= 0 {
		intervalSec = 10
	}
	timeoutSec := cfg.Poll.TimeoutSec
	if timeoutSec <= 0 {
		timeoutSec = 300
	}

	endpoint := cfg.Beetle.Endpoint
	nsId := cfg.Beetle.NamespaceID

	fmt.Println("=========================================================")
	fmt.Println(" CM-Beetle Async Response Test CLI")
	fmt.Println(" Recommend -> Poll -> Get -> Migrate -> Poll -> Get -> Delete -> Poll -> Get")
	fmt.Println("=========================================================")

	// [1/9] Recommend (async)
	fmt.Println("\n[1/9] POST /recommendation/infra (Prefer: respond-async)")
	reqBody, err := os.ReadFile(cfg.Beetle.RequestBodyFile)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read request body file")
	}
	recReqID := postAsync(client, endpoint+"/recommendation/infra", reqBody)
	fmt.Printf("      -> 202 Accepted, reqId=%s\n", recReqID)

	// [2/9] Poll
	fmt.Println("\n[2/9] Poll GET /request/{reqId} (recommendation)")
	recDetails := pollUntilDone(client, endpoint, recReqID, intervalSec, timeoutSec)
	if recDetails.Status != "Success" {
		log.Fatal().Str("status", recDetails.Status).Str("error", recDetails.ErrorResponse).Msg("recommendation did not succeed")
	}
	fmt.Printf("      -> status: %s\n", recDetails.Status)

	// [3/9] "Get" — recommendation results only exist inside the polled response; inspect them here.
	fmt.Println("\n[3/9] Inspect recommendation result")
	var candidates []json.RawMessage
	if err := json.Unmarshal(recDetails.ResponseData, &candidates); err != nil {
		log.Fatal().Err(err).Msg("failed to parse recommendation candidates")
	}
	if len(candidates) == 0 {
		log.Fatal().Msg("no recommendation candidates returned")
	}
	candidate := candidates[0]
	fmt.Printf("      -> %d candidate(s) recommended; using candidate #1 for migration\n", len(candidates))

	// [4/9] Migrate (async) — the recommended candidate JSON is posted as-is (RecommendedInfra shape).
	fmt.Println("\n[4/9] POST /migration/ns/{nsId}/infra (Prefer: respond-async)")
	migrateURL := fmt.Sprintf("%s/migration/ns/%s/infra", endpoint, nsId)
	migReqID := postAsync(client, migrateURL, candidate)
	fmt.Printf("      -> 202 Accepted, reqId=%s\n", migReqID)

	// [5/9] Poll
	fmt.Println("\n[5/9] Poll GET /request/{reqId} (migration)")
	migDetails := pollUntilDone(client, endpoint, migReqID, intervalSec, timeoutSec)
	if migDetails.Status != "Success" {
		log.Fatal().Str("status", migDetails.Status).Str("error", migDetails.ErrorResponse).Msg("migration did not succeed")
	}
	fmt.Printf("      -> status: %s\n", migDetails.Status)

	var infraInfo struct {
		Id string `json:"id"`
	}
	if err := json.Unmarshal(migDetails.ResponseData, &infraInfo); err != nil {
		log.Fatal().Err(err).Msg("failed to parse migrated infra info")
	}
	infraId := infraInfo.Id
	if infraId == "" {
		log.Fatal().Msg("migrated infra has no id")
	}
	fmt.Printf("      -> infraId: %s\n", infraId)

	// [6/9] Get
	fmt.Println("\n[6/9] GET /migration/ns/{nsId}/infra/{infraId}")
	getURL := fmt.Sprintf("%s/migration/ns/%s/infra/%s", endpoint, nsId, infraId)
	getResp, err := client.R().Get(getURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get migrated infra")
	}
	fmt.Printf("      -> HTTP %d: %s\n", getResp.StatusCode(), truncate(string(getResp.Body()), 300))

	// [7/9] Delete (async)
	fmt.Println("\n[7/9] DELETE /migration/ns/{nsId}/infra/{infraId} (Prefer: respond-async)")
	delReqID := deleteAsync(client, getURL+"?option=terminate")
	fmt.Printf("      -> 202 Accepted, reqId=%s\n", delReqID)

	// [8/9] Poll
	fmt.Println("\n[8/9] Poll GET /request/{reqId} (deletion)")
	delDetails := pollUntilDone(client, endpoint, delReqID, intervalSec, timeoutSec)
	if delDetails.Status != "Success" {
		log.Fatal().Str("status", delDetails.Status).Str("error", delDetails.ErrorResponse).Msg("deletion did not succeed")
	}
	fmt.Printf("      -> status: %s\n", delDetails.Status)

	// [9/9] Get again — expect 404, confirming the infra is gone.
	fmt.Println("\n[9/9] GET /migration/ns/{nsId}/infra/{infraId} (expect 404)")
	confirmResp, err := client.R().Get(getURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to confirm deletion")
	}
	if confirmResp.StatusCode() == http.StatusNotFound {
		fmt.Println("      -> HTTP 404: deletion confirmed")
	} else {
		fmt.Printf("      -> HTTP %d (expected 404): %s\n", confirmResp.StatusCode(), truncate(string(confirmResp.Body()), 300))
	}

	fmt.Println("\n=========================================================")
	fmt.Println(" Done.")
	fmt.Println("=========================================================")
}

// postAsync sends req body as a POST with Prefer: respond-async and returns the reqId from a 202 response.
func postAsync(client *resty.Client, url string, body []byte) string {
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Prefer", "respond-async").
		SetBody(body).
		Post(url)
	if err != nil {
		log.Fatal().Err(err).Str("url", url).Msg("async request failed")
	}
	if resp.StatusCode() != http.StatusAccepted {
		log.Fatal().Int("status", resp.StatusCode()).Str("body", string(resp.Body())).Msg("expected 202 Accepted")
	}
	return extractReqID(resp)
}

// deleteAsync sends a DELETE with Prefer: respond-async and returns the reqId from a 202 response.
func deleteAsync(client *resty.Client, url string) string {
	resp, err := client.R().
		SetHeader("Prefer", "respond-async").
		Delete(url)
	if err != nil {
		log.Fatal().Err(err).Str("url", url).Msg("async request failed")
	}
	if resp.StatusCode() != http.StatusAccepted {
		log.Fatal().Int("status", resp.StatusCode()).Str("body", string(resp.Body())).Msg("expected 202 Accepted")
	}
	return extractReqID(resp)
}

func extractReqID(resp *resty.Response) string {
	var apiResp ApiResponse
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		log.Fatal().Err(err).Msg("failed to parse async job response")
	}
	var job AsyncJobData
	if err := json.Unmarshal(apiResp.Data, &job); err != nil {
		log.Fatal().Err(err).Msg("failed to parse async job data")
	}
	if job.ReqID == "" {
		job.ReqID = resp.Header().Get("X-Request-Id")
	}
	if job.ReqID == "" {
		log.Fatal().Msg("no reqId in async job response")
	}
	return job.ReqID
}

// pollUntilDone polls GET /request/{reqId} on a fixed interval until status is Success or
// Error, or timeoutSec elapses.
func pollUntilDone(client *resty.Client, endpoint, reqID string, intervalSec, timeoutSec int) RequestDetails {
	statusURL := fmt.Sprintf("%s/request/%s", endpoint, reqID)
	start := time.Now()
	deadline := start.Add(time.Duration(timeoutSec) * time.Second)
	attempt := 0
	interval := time.Duration(intervalSec) * time.Second

	for time.Now().Before(deadline) {
		time.Sleep(interval)
		attempt++
		elapsed := time.Since(start).Round(time.Second)

		resp, err := client.R().Get(statusURL)
		if err != nil {
			fmt.Printf("      ... [%s elapsed, poll #%d] request error: %v\n", elapsed, attempt, err)
			continue
		}

		var apiResp ApiResponse
		if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
			fmt.Printf("      ... [%s elapsed, poll #%d] parse error: %v\n", elapsed, attempt, err)
			continue
		}
		var details RequestDetails
		if err := json.Unmarshal(apiResp.Data, &details); err != nil {
			fmt.Printf("      ... [%s elapsed, poll #%d] parse error: %v\n", elapsed, attempt, err)
			continue
		}

		fmt.Printf("      ... [%s elapsed, poll #%d] status: %s\n", elapsed, attempt, details.Status)
		if details.Status == "Success" || details.Status == "Error" {
			return details
		}
	}

	log.Fatal().Str("reqId", reqID).Msg("polling timed out")
	return RequestDetails{}
}

func loadConfig(path string) (TestConfig, error) {
	var cfg TestConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func loadAuthConfig(path string) (AuthConfig, error) {
	var auth AuthConfig
	if path == "" {
		return auth, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return auth, err
	}
	if err := json.Unmarshal(data, &auth); err != nil {
		return auth, err
	}
	return auth, nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
