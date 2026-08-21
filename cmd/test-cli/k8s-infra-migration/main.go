// Package main is the starting point of CM-Beetle K8s Infra Migration Test CLI.
//
// Runs the full K8s migration lifecycle against real CSPs: recommend -> migrate -> list ->
// get (verified against the recommendation) -> delete -> residual resource check.
//
// The CLI does not track per-CSP provisioning times. Beetle already polls until the cluster
// reaches Active (see pkg/core/migration/k8s-infra.go), so the only thing the CLI must solve
// is how to wait for a call that can take tens of minutes. It does that with the async job
// pattern (Prefer: respond-async -> poll GET /request/{reqId}), the same flow exercised by
// cmd/test-cli/async. Set migration.mode to "sync" to exercise the blocking path instead.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
)

// restyNoopLogger silences all Resty log output (e.g. "Basic Auth in HTTP mode" warnings).
type restyNoopLogger struct{}

func (restyNoopLogger) Errorf(_ string, _ ...interface{}) {}
func (restyNoopLogger) Warnf(_ string, _ ...interface{})  {}
func (restyNoopLogger) Debugf(_ string, _ ...interface{}) {}

// TestConfig holds test configuration.
type TestConfig struct {
	Test struct {
		Set struct {
			Mode              string `yaml:"mode"`              // parallel or sequential
			StartDelaySeconds int    `yaml:"startDelaySeconds"` // stagger between parallel starts
		} `yaml:"set"`
		Cases []TestCase `yaml:"cases"`
	} `yaml:"test"`
	Beetle struct {
		Endpoint        string `yaml:"endpoint"`
		NamespaceID     string `yaml:"namespaceId"`
		RequestBodyFile string `yaml:"requestBodyFile"`
		AuthConfigFile  string `yaml:"authConfigFile"`
		NameSeed        string `yaml:"nameSeed"`
	} `yaml:"beetle"`
	Migration struct {
		Mode       string `yaml:"mode"`       // async (default) or sync
		TimeoutSec int    `yaml:"timeoutSec"` // HTTP timeout for the sync path
	} `yaml:"migration"`
	Poll struct {
		IntervalSec int `yaml:"intervalSec"`
		TimeoutSec  int `yaml:"timeoutSec"`
	} `yaml:"poll"`
	Delete struct {
		TimeoutSec       int `yaml:"timeoutSec"`
		RetryIntervalSec int `yaml:"retryIntervalSec"`
		MaxRetries       int `yaml:"maxRetries"`
	} `yaml:"delete"`
	Verify struct {
		ResidualResources bool `yaml:"residualResources"`
	} `yaml:"verify"`
	Workload struct {
		Enabled              bool `yaml:"enabled"`
		KubeconfigTimeoutSec int  `yaml:"kubeconfigTimeoutSec"`
		KubeconfigPollSec    int  `yaml:"kubeconfigPollSec"`
		PodReadyTimeoutSec   int  `yaml:"podReadyTimeoutSec"`
		PodPollSec           int  `yaml:"podPollSec"`
		LoadBalancerEnabled  bool `yaml:"loadBalancerEnabled"`
		LbAddressTimeoutSec  int  `yaml:"lbAddressTimeoutSec"`
		LbAccessTimeoutSec   int  `yaml:"lbAccessTimeoutSec"`
		LbPollSec            int  `yaml:"lbPollSec"`
	} `yaml:"workload"`
}

// TestCase is one target CSP/region pair; only entries with Execute: true are used.
type TestCase struct {
	cloudmodel.CloudProperty `yaml:",inline"`
	Name                     string `yaml:"name"`
	Execute                  bool   `yaml:"execute"`
}

// AuthConfig holds Beetle and Tumblebug credentials. Tumblebug is called directly only for
// the residual resource check, since Beetle exposes no API for it.
type AuthConfig struct {
	BeetleApiUsername    string `json:"beetleApiUsername"`
	BeetleApiPassword    string `json:"beetleApiPassword"`
	TumblebugApiUsername string `json:"tumblebugApiUsername"`
	TumblebugApiPassword string `json:"tumblebugApiPassword"`
	TumblebugEndpoint    string `json:"tumblebugEndpoint"`
}

// progressf prints an in-flight progress line. Targets run concurrently and their output
// interleaves, so every line carries the target name — without it a stalled CSP cannot be told
// apart from a healthy one.
func progressf(target, format string, a ...interface{}) {
	fmt.Printf("      [%s] %s\n", target, fmt.Sprintf(format, a...))
}

// StepResult holds one lifecycle step's outcome.
type StepResult struct {
	Target     string
	Number     int
	Name       string
	StartTime  time.Time
	Duration   time.Duration
	Success    bool
	Skipped    bool
	StatusCode int
	Notes      []string // check outcomes, prefixed with ✅ / ❌ / ℹ️
	Error      string
}

// CSPTestReport captures the whole lifecycle for one CSP/region pair.
type CSPTestReport struct {
	CSP          string
	Region       string
	DisplayName  string
	TestDateTime time.Time

	BeetleURL     string
	NamespaceID   string
	BeetleVersion string
	GitHash       string

	NameSeed       string
	Recommendation *cloudmodel.RecommendedInfra
	ClusterID      string
	ClusterInfo    *tbmodel.K8sClusterInfo

	Steps []StepResult
}

// ApiResponseRaw mirrors model.ApiResponse[T] for payloads parsed in two stages.
type ApiResponseRaw struct {
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
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	applyDefaults(&cfg)

	targets := selectedTargets(cfg)
	if len(targets) == 0 {
		log.Fatal().Msg("At least 1 test case must have execute: true")
	}

	requestBody, err := os.ReadFile(cfg.Beetle.RequestBodyFile)
	if err != nil {
		log.Fatal().Err(err).Str("file", cfg.Beetle.RequestBodyFile).Msg("Failed to read request body file")
	}

	auth, err := loadAuthConfig(cfg.Beetle.AuthConfigFile)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to load auth config; proceeding without auth")
	}

	client := newClient(cfg, auth)
	if err := checkBeetleReadiness(client, cfg.Beetle.Endpoint); err != nil {
		log.Fatal().Err(err).Msg("CM-Beetle readiness check failed")
	}

	fmt.Println("=========================================================")
	fmt.Println(" CM-Beetle K8s Infra Migration Test CLI")
	fmt.Printf(" %d target(s), mode: %s, migration: %s\n", len(targets), cfg.Test.Set.Mode, cfg.Migration.Mode)
	fmt.Println(" Recommend -> Migrate -> List -> Get -> Delete -> Residual check")
	fmt.Println("=========================================================")

	reports := runAllTargets(cfg, auth, targets, requestBody)

	for _, r := range reports {
		if err := generateCSPReport(r); err != nil {
			log.Warn().Err(err).Str("csp", r.CSP).Msg("Failed to generate CSP report")
		}
	}
	if err := generateSummaryReport(reports); err != nil {
		log.Warn().Err(err).Msg("Failed to generate summary report")
	}

	if !printFinalSummary(reports) {
		os.Exit(1)
	}
}

func runAllTargets(cfg TestConfig, auth AuthConfig, targets []TestCase, requestBody []byte) []*CSPTestReport {
	reports := make([]*CSPTestReport, len(targets))

	if cfg.Test.Set.Mode == "sequential" {
		for i, t := range targets {
			reports[i] = runLifecycle(cfg, auth, t, requestBody, caseNameSeed(cfg, i))
		}
		return reports
	}

	// Each goroutine writes only its own slot, so results stay in target order without a mutex.
	var wg sync.WaitGroup
	for i, t := range targets {
		wg.Add(1)
		go func(idx int, target TestCase) {
			defer wg.Done()
			// Stagger starts so concurrent runs do not burst Tumblebug's read APIs at once.
			if d := cfg.Test.Set.StartDelaySeconds; d > 0 && idx > 0 {
				time.Sleep(time.Duration(idx*d) * time.Second)
			}
			reports[idx] = runLifecycle(cfg, auth, target, requestBody, caseNameSeed(cfg, idx))
		}(i, t)
	}
	wg.Wait()
	return reports
}

// runLifecycle executes the ordered steps for one target. A failed step skips the remaining
// ones, except cleanup: deletion always runs once a cluster ID exists, so a failure mid-run
// never leaves a billable cluster behind.
func runLifecycle(cfg TestConfig, auth AuthConfig, target TestCase, requestBody []byte, nameSeed string) *CSPTestReport {
	report := &CSPTestReport{
		CSP:           target.Csp,
		Region:        target.Region,
		DisplayName:   target.Name,
		TestDateTime:  time.Now(),
		BeetleURL:     cfg.Beetle.Endpoint,
		NamespaceID:   cfg.Beetle.NamespaceID,
		BeetleVersion: getBeetleVersion(),
		GitHash:       getGitHash(),
		NameSeed:      nameSeed,
	}
	client := newClient(cfg, auth)

	printBanner(target)

	// The name is declared here rather than read back from the result, so it can be announced
	// before the step runs — a step that takes minutes needs its heading up front.
	steps := []struct {
		number int
		name   string
		run    func(*resty.Client, TestConfig, AuthConfig, *CSPTestReport, []byte) StepResult
	}{
		{1, "POST /recommendation/k8sCluster", stepRecommend},
		{2, "POST /migration/ns/{nsId}/k8sCluster", stepMigrate},
		{3, "GET /migration/ns/{nsId}/k8sCluster", stepList},
		{4, "GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation", stepGetAndVerify},
	}

	runStep := func(number int, name string, fn func(*resty.Client, TestConfig, AuthConfig, *CSPTestReport, []byte) StepResult) StepResult {
		printStepStart(report.DisplayName, number, name)
		res := fn(client, cfg, auth, report, requestBody)
		report.Steps = append(report.Steps, res)
		printStep(report.DisplayName, res)
		return res
	}

	for _, st := range steps {
		if res := runStep(st.number, st.name, st.run); !res.Success {
			break
		}
	}

	// Workload verification runs only on a healthy cluster, and always before cleanup.
	if report.ClusterID != "" && cfg.Workload.Enabled && lastStepPassed(report) {
		runStep(5, "Workload verification (kubeconfig -> K8s API -> nginx)", stepWorkload)
	}

	// Cleanup always runs when there is something to clean up.
	if report.ClusterID != "" {
		res := runStep(6, "DELETE /migration/ns/{nsId}/k8sCluster/{id}", stepDelete)
		if res.Success && cfg.Verify.ResidualResources {
			runStep(7, "Residual resource check (Tumblebug)", stepResidualCheck)
		}
	}

	return report
}

// --- Steps -----------------------------------------------------------------

func stepRecommend(client *resty.Client, cfg TestConfig, _ AuthConfig, report *CSPTestReport, requestBody []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 1, Name: "POST /recommendation/k8sCluster", StartTime: time.Now()}

	url := cfg.Beetle.Endpoint + "/beetle/recommendation/k8sCluster"
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("desiredProvider", report.CSP).
		SetQueryParam("desiredRegion", report.Region).
		SetBody(requestBody).
		Post(url)
	res.Duration = time.Since(res.StartTime)

	if err != nil {
		res.Error = err.Error()
		return res
	}
	res.StatusCode = resp.StatusCode()
	if resp.StatusCode() != http.StatusOK {
		res.Error = fmt.Sprintf("expected 200, got %d: %s", resp.StatusCode(), truncate(string(resp.Body()), 300))
		return res
	}

	var apiResp model.ApiResponse[cloudmodel.RecommendedInfra]
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		res.Error = fmt.Sprintf("failed to parse recommendation: %s", err)
		return res
	}
	if !apiResp.Success {
		res.Error = apiResp.Error
		return res
	}

	report.Recommendation = &apiResp.Data
	cluster := apiResp.Data.TargetK8sCluster
	res.Notes = append(res.Notes,
		fmt.Sprintf("ℹ️  cluster: %s (version %s)", cluster.Name, cluster.Version),
		fmt.Sprintf("ℹ️  node groups: %d", len(cluster.K8sNodeGroupList)))
	for i, ng := range cluster.K8sNodeGroupList {
		res.Notes = append(res.Notes, fmt.Sprintf("ℹ️  node group[%d] %q spec=%s nodes=%d", i, ng.Name, ng.SpecId, ng.DesiredNodeSize))
	}
	res.Success = true
	return res
}

// stepMigrate provisions the cluster. Beetle blocks until the cluster is Active — up to 40
// minutes by its own polling budget — so the async path keeps the HTTP connection short and
// prints progress while waiting.
func stepMigrate(client *resty.Client, cfg TestConfig, _ AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 2, Name: "POST /migration/ns/{nsId}/k8sCluster", StartTime: time.Now()}

	body, err := json.Marshal(report.Recommendation)
	if err != nil {
		res.Duration = time.Since(res.StartTime)
		res.Error = fmt.Sprintf("failed to marshal recommendation: %s", err)
		return res
	}
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/k8sCluster", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID)
	if report.NameSeed != "" {
		url += "?nameSeed=" + report.NameSeed
		res.Notes = append(res.Notes, fmt.Sprintf("ℹ️  nameSeed: %s", report.NameSeed))
	}

	var clusterJSON []byte
	if cfg.Migration.Mode == "sync" {
		clusterJSON, res.StatusCode, err = migrateSync(client, cfg, url, body, &res)
	} else {
		clusterJSON, res.StatusCode, err = migrateAsync(client, cfg, url, body, &res)
	}
	res.Duration = time.Since(res.StartTime)
	if err != nil {
		res.Error = err.Error()
		// Migration can fail after the cluster itself was created (e.g. a node group fails).
		// Recover the ID by name so cleanup still runs and no billable cluster is orphaned.
		adoptOrphanCluster(client, cfg, report, &res)
		return res
	}

	var info tbmodel.K8sClusterInfo
	if err := json.Unmarshal(clusterJSON, &info); err != nil {
		res.Error = fmt.Sprintf("failed to parse cluster info: %s", err)
		adoptOrphanCluster(client, cfg, report, &res)
		return res
	}
	report.ClusterInfo = &info
	report.ClusterID = info.Id

	res.Notes = append(res.Notes,
		fmt.Sprintf("ℹ️  cluster id: %s", info.Id),
		fmt.Sprintf("ℹ️  elapsed: %v", res.Duration.Truncate(time.Second)))
	if string(info.Status) == "Active" {
		res.Notes = append(res.Notes, "✅ status: Active")
	} else {
		res.Notes = append(res.Notes, fmt.Sprintf("❌ status: %s (want Active)", info.Status))
		res.Error = fmt.Sprintf("cluster status is %s", info.Status)
		return res // ClusterID is already set, so cleanup still runs
	}
	res.Success = true
	return res
}

func migrateSync(client *resty.Client, cfg TestConfig, url string, body []byte, res *StepResult) ([]byte, int, error) {
	syncClient := *client
	sc := (&syncClient).SetTimeout(time.Duration(cfg.Migration.TimeoutSec) * time.Second)

	progressf(res.Target, "... sync mode: holding the connection for up to %ds", cfg.Migration.TimeoutSec)
	resp, err := sc.R().SetHeader("Content-Type", "application/json").SetBody(body).Post(url)
	if err != nil {
		return nil, 0, err
	}
	if resp.StatusCode() != http.StatusCreated {
		return nil, resp.StatusCode(), fmt.Errorf("expected 201, got %d: %s", resp.StatusCode(), truncate(string(resp.Body()), 300))
	}
	var apiResp ApiResponseRaw
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return nil, resp.StatusCode(), fmt.Errorf("failed to parse response: %w", err)
	}
	return apiResp.Data, resp.StatusCode(), nil
}

// migrateAsync posts with Prefer: respond-async and polls the request record. A 503 means
// Beetle's async job pool (20 concurrent) is full, so it is retried rather than failed.
func migrateAsync(client *resty.Client, cfg TestConfig, url string, body []byte, res *StepResult) ([]byte, int, error) {
	const maxAdmissionRetries = 5
	var resp *resty.Response
	var err error

	for attempt := 1; attempt <= maxAdmissionRetries; attempt++ {
		resp, err = client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Prefer", "respond-async").
			SetBody(body).
			Post(url)
		if err != nil {
			return nil, 0, err
		}
		if resp.StatusCode() != http.StatusServiceUnavailable {
			break
		}
		wait := 30 * time.Second
		progressf(res.Target, "... async job pool full (503), retrying in %v (%d/%d)", wait, attempt, maxAdmissionRetries)
		time.Sleep(wait)
	}

	if resp.StatusCode() != http.StatusAccepted {
		return nil, resp.StatusCode(), fmt.Errorf("expected 202, got %d: %s", resp.StatusCode(), truncate(string(resp.Body()), 300))
	}

	reqID, err := extractReqID(resp)
	if err != nil {
		return nil, resp.StatusCode(), err
	}
	res.Notes = append(res.Notes, fmt.Sprintf("ℹ️  async reqId: %s", reqID))
	progressf(res.Target, "-> 202 Accepted, reqId=%s", reqID)

	details, err := pollUntilDone(client, res.Target, cfg.Beetle.Endpoint, reqID, cfg.Poll.IntervalSec, cfg.Poll.TimeoutSec)
	if err != nil {
		return nil, resp.StatusCode(), err
	}
	if details.Status != "Success" {
		return nil, resp.StatusCode(), fmt.Errorf("async job ended with status %s: %s", details.Status, truncate(details.ErrorResponse, 300))
	}

	// The recorded response body is the same ApiResponse envelope the sync path returns.
	var apiResp ApiResponseRaw
	if err := json.Unmarshal(details.ResponseData, &apiResp); err == nil && len(apiResp.Data) > 0 {
		return apiResp.Data, resp.StatusCode(), nil
	}
	return details.ResponseData, resp.StatusCode(), nil
}

// caseNameSeed gives each target its own name prefix so concurrent runs against different
// CSPs do not fight over the fixed resource names a recommendation produces. Mirrors the
// per-case seeding cmd/test-cli/infra uses.
func caseNameSeed(cfg TestConfig, index int) string {
	if cfg.Beetle.NameSeed == "" {
		return ""
	}
	return fmt.Sprintf("%s%02d", cfg.Beetle.NameSeed, index+1)
}

// lastStepPassed reports whether the most recent step succeeded, so workload verification is
// attempted only on a cluster that passed the earlier checks.
// specMatch reports whether a created node group's spec is the one that was recommended.
//
// Tumblebug normally echoes back its own namespaced spec ID ("aws+ap-northeast-2+c5a.xlarge"),
// but for some CSPs the stored ID is empty and the CSP's native name surfaces instead
// ("Standard_B4as_v2" on Azure). Rather than loosening the comparison to accommodate that, the
// recommended ID is resolved through Tumblebug and its recorded cspSpecName is compared
// exactly — so both renderings are checked against the same authoritative source.
//
// matched=false with err!=nil means the check could not be performed, which is reported as
// unknown rather than as a mismatch: a failed lookup is not evidence of a wrong spec.
func specMatch(sess *tbclient.Session, got, want string) (matched bool, err error) {
	if got == want {
		return true, nil
	}
	if sess == nil {
		return false, fmt.Errorf("no Tumblebug endpoint configured to resolve the spec")
	}
	spec, err := sess.ReadVmSpec("system", want)
	if err != nil {
		return false, fmt.Errorf("could not resolve spec %q via Tumblebug: %w", want, err)
	}
	if spec.CspSpecName == "" {
		return false, fmt.Errorf("Tumblebug has no cspSpecName recorded for spec %q", want)
	}
	return got == spec.CspSpecName, nil
}

func lastStepPassed(report *CSPTestReport) bool {
	if len(report.Steps) == 0 {
		return false
	}
	return report.Steps[len(report.Steps)-1].Success
}

// adoptOrphanCluster looks up a cluster by the recommended name and records its ID on the
// report. Beetle creates the cluster before its node groups, so a migration that fails
// partway can leave a running, billable cluster whose ID never reached the caller. Recording
// it here is what lets the cleanup step tear it down.
func adoptOrphanCluster(client *resty.Client, cfg TestConfig, report *CSPTestReport, res *StepResult) {
	if report.ClusterID != "" || report.Recommendation == nil {
		return
	}
	wantName := report.Recommendation.TargetK8sCluster.Name
	if wantName == "" {
		return
	}
	// Beetle applies the seed server-side, so the created cluster carries the prefixed name.
	if report.NameSeed != "" {
		wantName = report.NameSeed + "-" + wantName
	}

	url := fmt.Sprintf("%s/beetle/migration/ns/%s/k8sCluster", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID)
	resp, err := client.R().Get(url)
	if err != nil || resp.StatusCode() != http.StatusOK {
		res.Notes = append(res.Notes, "❌ could not list clusters to check for a partially created one — verify manually")
		return
	}

	var apiResp model.ApiResponse[[]tbmodel.K8sClusterInfo]
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		res.Notes = append(res.Notes, "❌ could not parse cluster list to check for a partially created one — verify manually")
		return
	}
	for _, c := range apiResp.Data {
		if c.Name == wantName || c.Id == wantName {
			cluster := c
			report.ClusterID = cluster.Id
			report.ClusterInfo = &cluster
			res.Notes = append(res.Notes,
				fmt.Sprintf("⚠️  cluster %s exists despite the failure (status %s) — scheduling cleanup", cluster.Id, cluster.Status))
			return
		}
	}
	res.Notes = append(res.Notes, "ℹ️  no partially created cluster found — nothing to clean up")
}

func stepList(client *resty.Client, cfg TestConfig, _ AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 3, Name: "GET /migration/ns/{nsId}/k8sCluster", StartTime: time.Now()}

	// option=id keeps this cheap: the full listing refreshes every cluster through the CSP and
	// exceeds Tumblebug's 120 s request timeout once several clusters exist. Presence of the ID
	// is all this step needs.
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/k8sCluster?option=id", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID)
	resp, err := client.R().Get(url)
	res.Duration = time.Since(res.StartTime)
	if err != nil {
		res.Error = err.Error()
		return res
	}
	res.StatusCode = resp.StatusCode()
	if resp.StatusCode() != http.StatusOK {
		res.Error = fmt.Sprintf("expected 200, got %d", resp.StatusCode())
		return res
	}

	var apiResp model.ApiResponse[cloudmodel.IdList]
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		res.Error = fmt.Sprintf("failed to parse cluster ID list: %s", err)
		return res
	}

	found := false
	for _, id := range apiResp.Data.IdList {
		if id == report.ClusterID {
			found = true
			break
		}
	}
	if found {
		res.Notes = append(res.Notes, fmt.Sprintf("✅ migrated cluster present in list (%d total)", len(apiResp.Data.IdList)))
		res.Success = true
	} else {
		res.Notes = append(res.Notes, fmt.Sprintf("❌ cluster %s not found in list of %d", report.ClusterID, len(apiResp.Data.IdList)))
		res.Error = "migrated cluster missing from list"
	}
	return res
}

// stepGetAndVerify is this CLI's distinctive check: it compares what was actually created
// against what the recommendation asked for. A migration that succeeds but silently drops a
// node group or substitutes a spec would otherwise go unnoticed.
func stepGetAndVerify(client *resty.Client, cfg TestConfig, auth AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 4, Name: "GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation", StartTime: time.Now()}

	url := fmt.Sprintf("%s/beetle/migration/ns/%s/k8sCluster/%s", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, report.ClusterID)
	resp, err := client.R().Get(url)
	res.Duration = time.Since(res.StartTime)
	if err != nil {
		res.Error = err.Error()
		return res
	}
	res.StatusCode = resp.StatusCode()
	if resp.StatusCode() != http.StatusOK {
		res.Error = fmt.Sprintf("expected 200, got %d", resp.StatusCode())
		return res
	}

	var apiResp model.ApiResponse[tbmodel.K8sClusterInfo]
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		res.Error = fmt.Sprintf("failed to parse cluster: %s", err)
		return res
	}
	actual := apiResp.Data
	report.ClusterInfo = &actual

	want := report.Recommendation.TargetK8sCluster
	ok := true

	if string(actual.Status) == "Active" {
		res.Notes = append(res.Notes, "✅ status: Active")
	} else {
		ok = false
		res.Notes = append(res.Notes, fmt.Sprintf("❌ status: %s (want Active)", actual.Status))
	}

	if len(actual.K8sNodeGroupList) == len(want.K8sNodeGroupList) {
		res.Notes = append(res.Notes, fmt.Sprintf("✅ node group count matches recommendation: %d", len(actual.K8sNodeGroupList)))
	} else {
		ok = false
		res.Notes = append(res.Notes, fmt.Sprintf("❌ node group count %d, recommendation asked for %d",
			len(actual.K8sNodeGroupList), len(want.K8sNodeGroupList)))
	}

	// Spec IDs are resolved through Tumblebug (see specMatch); Beetle exposes no spec API.
	// A nil session leaves the spec check unverified rather than failing it.
	var specSession *tbclient.Session
	if auth.TumblebugEndpoint != "" {
		specSession = tbclient.NewClient(tbclient.ApiConfig{
			RestUrl:  auth.TumblebugEndpoint + "/tumblebug",
			Username: auth.TumblebugApiUsername,
			Password: auth.TumblebugApiPassword,
		}).NewSession()
	}

	// Match by name: CSPs may reorder node groups, so index-based comparison is unreliable.
	wantByName := make(map[string]cloudmodel.K8sNodeGroupReq, len(want.K8sNodeGroupList))
	for _, ng := range want.K8sNodeGroupList {
		wantByName[ng.Name] = ng
	}
	for _, got := range actual.K8sNodeGroupList {
		w, exists := wantByName[got.Name]
		if !exists {
			ok = false
			res.Notes = append(res.Notes, fmt.Sprintf("❌ node group %q was not in the recommendation", got.Name))
			continue
		}
		specOK, specErr := specMatch(specSession, got.SpecId, w.SpecId)
		switch {
		case specErr != nil:
			res.Notes = append(res.Notes, fmt.Sprintf("ℹ️  node group %q spec=%s — not verified: %s", got.Name, got.SpecId, specErr))
		case !specOK:
			ok = false
			res.Notes = append(res.Notes, fmt.Sprintf("❌ node group %q spec=%s, recommended %s", got.Name, got.SpecId, w.SpecId))
		}
		if got.DesiredNodeSize != w.DesiredNodeSize {
			ok = false
			res.Notes = append(res.Notes, fmt.Sprintf("❌ node group %q desiredNodeSize=%d, recommended %d",
				got.Name, got.DesiredNodeSize, w.DesiredNodeSize))
		}
		if specOK && got.DesiredNodeSize == w.DesiredNodeSize {
			res.Notes = append(res.Notes, fmt.Sprintf("✅ node group %q matches (spec=%s, nodes=%d)", got.Name, got.SpecId, got.DesiredNodeSize))
		}
	}

	// Version is reported by the CSP and may carry a provider suffix (e.g. "1.30.1-aliyun.1"),
	// so a prefix match is the meaningful comparison.
	if want.Version != "" {
		if strings.HasPrefix(actual.Version, want.Version) || strings.HasPrefix(want.Version, actual.Version) {
			res.Notes = append(res.Notes, fmt.Sprintf("✅ version: %s (recommended %s)", actual.Version, want.Version))
		} else {
			res.Notes = append(res.Notes, fmt.Sprintf("ℹ️  version: %s (recommended %s) — provider-specific rendering", actual.Version, want.Version))
		}
	}

	res.Success = ok
	if !ok {
		res.Error = "created cluster does not match the recommendation"
	}
	return res
}

// stepDelete removes the cluster. Deletion has no async mode, and Beetle waits for every node
// group to disappear before deleting the cluster, so this call is long-running. Some CSPs
// (observed on NCP) reject deletion for several minutes after reporting Active, so failures
// are retried with a fixed backoff rather than matched against specific error strings.
func stepDelete(client *resty.Client, cfg TestConfig, _ AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 6, Name: "DELETE /migration/ns/{nsId}/k8sCluster/{id}", StartTime: time.Now()}

	deleteClient := *client
	dc := (&deleteClient).SetTimeout(time.Duration(cfg.Delete.TimeoutSec) * time.Second)
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/k8sCluster/%s", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, report.ClusterID)

	for attempt := 1; attempt <= cfg.Delete.MaxRetries; attempt++ {
		resp, err := dc.R().Delete(url)
		if err != nil {
			res.Notes = append(res.Notes, fmt.Sprintf("ℹ️  attempt %d: %s", attempt, err))
		} else {
			res.StatusCode = resp.StatusCode()
			body := string(resp.Body())

			if resp.StatusCode() == http.StatusOK {
				res.Duration = time.Since(res.StartTime)
				res.Notes = append(res.Notes, fmt.Sprintf("✅ deleted on attempt %d (%v)", attempt, res.Duration.Truncate(time.Second)))
				res.Success = true
				return res
			}
			// Deletion of an already-absent cluster is a no-op success for cleanup purposes.
			if isAlreadyGone(body) {
				res.Duration = time.Since(res.StartTime)
				res.Notes = append(res.Notes,
					"✅ cluster already absent — treated as cleanup success",
					"ℹ️  known gap: Beetle returns 500 for a missing cluster instead of a no-op success")
				res.Success = true
				return res
			}
			res.Notes = append(res.Notes, fmt.Sprintf("ℹ️  attempt %d: HTTP %d — %s", attempt, resp.StatusCode(), truncate(body, 200)))
		}

		if attempt < cfg.Delete.MaxRetries {
			wait := time.Duration(cfg.Delete.RetryIntervalSec) * time.Second
			progressf(res.Target, "... deletion not accepted yet, retrying in %v (%d/%d)", wait, attempt, cfg.Delete.MaxRetries)
			time.Sleep(wait)
		}
	}

	res.Duration = time.Since(res.StartTime)
	res.Error = fmt.Sprintf("deletion failed after %d attempts — cluster %s may still exist and incur cost",
		cfg.Delete.MaxRetries, report.ClusterID)
	return res
}

// stepResidualCheck reports whether the prerequisite resources survive cluster deletion.
// Leftover VNet/SecurityGroup/SshKey is current known behaviour, so this is recorded rather
// than failed: cleanup scope is tracked as a separate Beetle improvement.
func stepResidualCheck(_ *resty.Client, cfg TestConfig, auth AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 7, Name: "Residual resource check (Tumblebug)", StartTime: time.Now()}

	if auth.TumblebugEndpoint == "" {
		res.Duration = time.Since(res.StartTime)
		res.Skipped = true
		res.Success = true
		res.Notes = append(res.Notes, "ℹ️  skipped: tumblebugEndpoint not set in auth config")
		return res
	}
	if report.ClusterInfo == nil {
		res.Duration = time.Since(res.StartTime)
		res.Skipped = true
		res.Success = true
		res.Notes = append(res.Notes, "ℹ️  skipped: no cluster info captured before deletion")
		return res
	}

	sess := tbclient.NewClient(tbclient.ApiConfig{
		RestUrl:  auth.TumblebugEndpoint + "/tumblebug",
		Username: auth.TumblebugApiUsername,
		Password: auth.TumblebugApiPassword,
	}).NewSession()

	ns := cfg.Beetle.NamespaceID
	net := report.ClusterInfo.Network

	if net.VNetId != "" {
		if _, err := sess.ReadVNet(ns, net.VNetId); err == nil {
			res.Notes = append(res.Notes, fmt.Sprintf("ℹ️  VNet %s still exists (known gap)", net.VNetId))
		} else {
			res.Notes = append(res.Notes, fmt.Sprintf("✅ VNet %s removed", net.VNetId))
		}
	}
	for _, sgID := range net.SecurityGroupIds {
		if _, err := sess.ReadSecurityGroup(ns, sgID); err == nil {
			res.Notes = append(res.Notes, fmt.Sprintf("ℹ️  SecurityGroup %s still exists (known gap)", sgID))
		} else {
			res.Notes = append(res.Notes, fmt.Sprintf("✅ SecurityGroup %s removed", sgID))
		}
	}
	seenKeys := map[string]bool{}
	for _, ng := range report.ClusterInfo.K8sNodeGroupList {
		if ng.SshKeyId == "" || seenKeys[ng.SshKeyId] {
			continue
		}
		seenKeys[ng.SshKeyId] = true
		if _, err := sess.ReadSshKey(ns, ng.SshKeyId); err == nil {
			res.Notes = append(res.Notes, fmt.Sprintf("ℹ️  SshKey %s still exists (known gap)", ng.SshKeyId))
		} else {
			res.Notes = append(res.Notes, fmt.Sprintf("✅ SshKey %s removed", ng.SshKeyId))
		}
	}

	res.Duration = time.Since(res.StartTime)
	res.Success = true // informational only
	return res
}

// --- Helpers ---------------------------------------------------------------

func isAlreadyGone(body string) bool {
	lower := strings.ToLower(body)
	return strings.Contains(lower, "not exist") || strings.Contains(lower, "not found")
}

func extractReqID(resp *resty.Response) (string, error) {
	var apiResp ApiResponseRaw
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		return "", fmt.Errorf("failed to parse async job response: %w", err)
	}
	var job AsyncJobData
	if err := json.Unmarshal(apiResp.Data, &job); err != nil {
		return "", fmt.Errorf("failed to parse async job data: %w", err)
	}
	if job.ReqID == "" {
		job.ReqID = resp.Header().Get("X-Request-Id")
	}
	if job.ReqID == "" {
		return "", fmt.Errorf("no reqId in async job response")
	}
	return job.ReqID, nil
}

// pollUntilDone polls GET /request/{reqId} until the job finishes or the timeout elapses.
func pollUntilDone(client *resty.Client, target, endpoint, reqID string, intervalSec, timeoutSec int) (RequestDetails, error) {
	statusURL := fmt.Sprintf("%s/beetle/request/%s", endpoint, reqID)
	start := time.Now()
	deadline := start.Add(time.Duration(timeoutSec) * time.Second)
	interval := time.Duration(intervalSec) * time.Second
	attempt := 0

	for time.Now().Before(deadline) {
		time.Sleep(interval)
		attempt++
		elapsed := time.Since(start).Round(time.Second)

		resp, err := client.R().Get(statusURL)
		if err != nil {
			progressf(target, "... [%s elapsed, poll #%d] request error: %v", elapsed, attempt, err)
			continue
		}
		var apiResp ApiResponseRaw
		if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
			progressf(target, "... [%s elapsed, poll #%d] parse error: %v", elapsed, attempt, err)
			continue
		}
		var details RequestDetails
		if err := json.Unmarshal(apiResp.Data, &details); err != nil {
			progressf(target, "... [%s elapsed, poll #%d] parse error: %v", elapsed, attempt, err)
			continue
		}

		progressf(target, "... [%s elapsed, poll #%d] status: %s", elapsed, attempt, details.Status)
		if details.Status == "Success" || details.Status == "Error" {
			return details, nil
		}
	}
	return RequestDetails{}, fmt.Errorf("polling timed out after %ds", timeoutSec)
}

func newClient(cfg TestConfig, auth AuthConfig) *resty.Client {
	c := resty.New().SetTimeout(2 * time.Minute).SetLogger(restyNoopLogger{})
	if auth.BeetleApiUsername != "" {
		c.SetBasicAuth(auth.BeetleApiUsername, auth.BeetleApiPassword)
	}
	return c
}

func applyDefaults(cfg *TestConfig) {
	if cfg.Test.Set.Mode != "sequential" {
		cfg.Test.Set.Mode = "parallel"
	}
	if cfg.Migration.Mode != "sync" {
		cfg.Migration.Mode = "async"
	}
	setIfZero(&cfg.Migration.TimeoutSec, 3600) // 60 min: Beetle polls up to 40 min for Active
	setIfZero(&cfg.Poll.IntervalSec, 30)
	setIfZero(&cfg.Poll.TimeoutSec, 3600)
	setIfZero(&cfg.Delete.TimeoutSec, 1800) // 30 min: node group teardown is polled for up to 20 min
	setIfZero(&cfg.Delete.RetryIntervalSec, 60)
	setIfZero(&cfg.Delete.MaxRetries, 10)
	setIfZero(&cfg.Test.Set.StartDelaySeconds, 10)
	setIfZero(&cfg.Workload.KubeconfigTimeoutSec, 600) // kubeconfig lags Active on some CSPs
	setIfZero(&cfg.Workload.KubeconfigPollSec, 60)
	setIfZero(&cfg.Workload.PodReadyTimeoutSec, 180)
	setIfZero(&cfg.Workload.PodPollSec, 10)
	setIfZero(&cfg.Workload.LbAddressTimeoutSec, 900) // CSP load balancer provisioning is slow
	setIfZero(&cfg.Workload.LbAccessTimeoutSec, 300)  // health checks and DNS settle after the address appears
	setIfZero(&cfg.Workload.LbPollSec, 15)
	if cfg.Beetle.NamespaceID == "" {
		cfg.Beetle.NamespaceID = "mig01"
	}
}

func setIfZero(v *int, def int) {
	if *v <= 0 {
		*v = def
	}
}

func selectedTargets(cfg TestConfig) []TestCase {
	var out []TestCase
	for _, c := range cfg.Test.Cases {
		if c.Execute {
			out = append(out, c)
		}
	}
	return out
}

func truncate(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func printBanner(t TestCase) {
	fmt.Printf("\n=========================================================\n")
	fmt.Printf(" Target: %s (%s / %s)\n", t.Name, t.Csp, t.Region)
	fmt.Printf("=========================================================\n")
}

// printStepStart announces a step before it runs, so long-running work (cluster creation,
// load balancer provisioning) has a heading above its progress lines instead of appearing
// under the previous step's result.
func printStepStart(target string, number int, name string) {
	fmt.Printf("▶  [%s] Step %d: %s ...\n", target, number, name)
}

// printStep reports a finished step. The duration belongs here rather than on the start line,
// because it is only known once the step ends.
func printStep(target string, res StepResult) {
	icon := "✅"
	switch {
	case res.Skipped:
		icon = "⏭️"
	case !res.Success:
		icon = "❌"
	}
	fmt.Printf("%s [%s] Step %d: %s — done in %v\n", icon, target, res.Number, res.Name,
		res.Duration.Truncate(time.Millisecond))
	for _, n := range res.Notes {
		fmt.Printf("      [%s] %s\n", target, n)
	}
	if res.Error != "" {
		fmt.Printf("      [%s] error: %s\n", target, res.Error)
	}
}

func printFinalSummary(reports []*CSPTestReport) bool {
	fmt.Println("\n=========================================================")
	fmt.Println(" OVERALL TEST SUMMARY")
	fmt.Println("=========================================================")

	allPassed := true
	for _, r := range reports {
		passed, total := countSteps(r)
		status := "✅"
		if passed != total {
			status = "❌"
			allPassed = false
		}
		fmt.Printf(" %s %-22s %d/%d steps passed\n", status, r.DisplayName, passed, total)
	}
	fmt.Println("=========================================================")
	return allPassed
}

func countSteps(r *CSPTestReport) (passed, total int) {
	for _, s := range r.Steps {
		total++
		if s.Success {
			passed++
		}
	}
	return passed, total
}

// checkBeetleReadiness checks if CM-Beetle is ready using GET /beetle/readyz.
func checkBeetleReadiness(client *resty.Client, beetleURL string) error {
	fmt.Println("🔍 Checking CM-Beetle readiness...")
	url := beetleURL + "/beetle/readyz"

	var response map[string]interface{}
	var emptyBody interface{} = common.NoBody
	err := common.ExecuteHttpRequest(client, "GET", url, nil, common.SetUseBody(emptyBody), &emptyBody, &response, 0)
	if err != nil {
		return fmt.Errorf("CM-Beetle readiness check failed: %w", err)
	}
	if message, ok := response["message"].(string); ok && strings.Contains(message, "NOT ready") {
		return fmt.Errorf("CM-Beetle is not ready: %s", message)
	}
	fmt.Println("✅ CM-Beetle is ready!")
	return nil
}

func loadConfig(configPath string) (TestConfig, error) {
	var cfg TestConfig
	file, err := os.Open(configPath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func loadAuthConfig(authConfigPath string) (AuthConfig, error) {
	var auth AuthConfig
	file, err := os.Open(authConfigPath)
	if err != nil {
		return auth, err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&auth); err != nil {
		return auth, fmt.Errorf("failed to decode auth config: %w", err)
	}
	return auth, nil
}

// getGitHash returns the current git commit hash.
func getGitHash() string {
	out, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

// getBeetleVersion returns the CM-Beetle version from git tags or commit hash.
// Only beetle release tags (v[0-9]*) are matched to avoid picking up imdl/*, transx/* tags.
func getBeetleVersion() string {
	commitHash := getGitHash()

	if exactTag, err := exec.Command("git", "describe", "--tags", "--exact-match", "--match", "v[0-9]*").Output(); err == nil {
		return strings.TrimSpace(string(exactTag))
	}
	if tag, err := exec.Command("git", "describe", "--tags", "--abbrev=0", "--match", "v[0-9]*").Output(); err == nil {
		if tagStr := strings.TrimSpace(string(tag)); tagStr != "" {
			if commitHash != "unknown" {
				return fmt.Sprintf("%s+ (%s)", tagStr, commitHash)
			}
			return tagStr + "+"
		}
	}
	if commitHash != "unknown" {
		return fmt.Sprintf("main (%s)", commitHash)
	}
	return "main (unknown)"
}

// maskSensitiveInfo removes CSP account identifiers from report content.
func maskSensitiveInfo(content string) string {
	reSub := regexp.MustCompile(`(?i)/subscriptions/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)
	content = reSub.ReplaceAllString(content, "/subscriptions/AZURE_SUBSCRIPTION_ID")

	reGCP := regexp.MustCompile(`projects/([a-z0-9\-]+)/`)
	content = reGCP.ReplaceAllStringFunc(content, func(match string) string {
		parts := strings.Split(match, "/")
		if len(parts) >= 2 && (parts[1] == "compute" || parts[1] == "v1") {
			return match
		}
		return "projects/GCP_PROJECT_ID/"
	})

	reEmail := regexp.MustCompile(`[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}`)
	return reEmail.ReplaceAllString(content, "MASKED_EMAIL")
}

func testResultDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(cwd, "testresult")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

func generateCSPReport(report *CSPTestReport) error {
	dir, err := testResultDir()
	if err != nil {
		return err
	}
	path := filepath.Join(dir, fmt.Sprintf("k8s-infra-migration-test-results-%s.md", strings.ToLower(report.CSP)))
	if err := os.WriteFile(path, []byte(maskSensitiveInfo(buildCSPMarkdown(report))), 0644); err != nil {
		return err
	}
	log.Info().Str("file", path).Msg("✅ CSP test report generated and saved")
	return nil
}

func buildCSPMarkdown(report *CSPTestReport) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# CM-Beetle K8s Infra Migration Test Results — %s\n\n", report.DisplayName))
	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> Full lifecycle against a real CSP: recommend → migrate → list → get (verified against\n")
	sb.WriteString("> the recommendation) → delete → residual resource check.\n\n")

	sb.WriteString("## Environment\n\n")
	sb.WriteString(fmt.Sprintf("- CSP / Region: %s / %s\n", report.CSP, report.Region))
	sb.WriteString(fmt.Sprintf("- CM-Beetle URL: %s\n", report.BeetleURL))
	sb.WriteString(fmt.Sprintf("- CM-Beetle Version: %s\n", report.BeetleVersion))
	sb.WriteString(fmt.Sprintf("- Git Commit: %s\n", report.GitHash))
	sb.WriteString(fmt.Sprintf("- Namespace: %s\n", report.NamespaceID))
	sb.WriteString(fmt.Sprintf("- Test Date: %s\n", report.TestDateTime.Format("2006-01-02 15:04:05 MST")))
	if report.ClusterID != "" {
		sb.WriteString(fmt.Sprintf("- Cluster ID: %s\n", report.ClusterID))
	}
	sb.WriteString("\n")

	passed, total := countSteps(report)
	sb.WriteString("## Test Results Summary\n\n")
	sb.WriteString("| Step | Description | Status | Duration |\n")
	sb.WriteString("|------|-------------|--------|----------|\n")
	var totalDuration time.Duration
	for _, s := range report.Steps {
		status := "✅ **PASS**"
		switch {
		case s.Skipped:
			status = "⏭️ **SKIP**"
		case !s.Success:
			status = "❌ **FAIL**"
		}
		sb.WriteString(fmt.Sprintf("| %d | %s | %s | %v |\n", s.Number, s.Name, status, s.Duration.Truncate(time.Millisecond)))
		totalDuration += s.Duration
	}
	sb.WriteString(fmt.Sprintf("\n**Overall Result**: %d/%d steps passed", passed, total))
	if passed == total {
		sb.WriteString(" ✅\n\n")
	} else {
		sb.WriteString(" ❌\n\n")
	}
	sb.WriteString(fmt.Sprintf("**Total Duration**: %v\n\n---\n\n", totalDuration.Truncate(time.Second)))

	sb.WriteString("## Step Details\n\n")
	for _, s := range report.Steps {
		sb.WriteString(fmt.Sprintf("### Step %d — %s\n\n", s.Number, s.Name))
		sb.WriteString(fmt.Sprintf("- **Duration**: %v\n", s.Duration.Truncate(time.Millisecond)))
		if s.StatusCode != 0 {
			sb.WriteString(fmt.Sprintf("- **Status Code**: %d\n", s.StatusCode))
		}
		if s.Error != "" {
			sb.WriteString(fmt.Sprintf("- **Error**: %s\n", s.Error))
		}
		sb.WriteString("\n")
		for _, n := range s.Notes {
			sb.WriteString(fmt.Sprintf("- %s\n", n))
		}
		sb.WriteString("\n")
	}

	if report.Recommendation != nil {
		sb.WriteString("## Recommendation (input to migration)\n\n")
		sb.WriteString("<details>\n  <summary> <ins>Click to see the recommendation</ins> </summary>\n\n```json\n")
		b, _ := json.MarshalIndent(report.Recommendation, "", "  ")
		sb.WriteString(string(b))
		sb.WriteString("\n```\n\n</details>\n\n")
	}
	if report.ClusterInfo != nil {
		sb.WriteString("## Created Cluster\n\n")
		sb.WriteString("<details>\n  <summary> <ins>Click to see the cluster info</ins> </summary>\n\n```json\n")
		b, _ := json.MarshalIndent(report.ClusterInfo, "", "  ")
		sb.WriteString(string(b))
		sb.WriteString("\n```\n\n</details>\n\n")
	}

	return sb.String()
}

func generateSummaryReport(reports []*CSPTestReport) error {
	dir, err := testResultDir()
	if err != nil {
		return err
	}

	var sb strings.Builder
	sb.WriteString("# CM-Beetle K8s Infra Migration Test Summary\n\n")
	if len(reports) > 0 {
		sb.WriteString(fmt.Sprintf("- CM-Beetle Version: %s\n", reports[0].BeetleVersion))
		sb.WriteString(fmt.Sprintf("- Test Date: %s\n\n", reports[0].TestDateTime.Format("2006-01-02 15:04:05 MST")))
	}

	sb.WriteString("| Target | CSP / Region | Steps Passed | Cluster ID | Result |\n")
	sb.WriteString("|--------|--------------|--------------|------------|--------|\n")
	for _, r := range reports {
		passed, total := countSteps(r)
		result := "✅ PASS"
		if passed != total {
			result = "❌ FAIL"
		}
		clusterID := r.ClusterID
		if clusterID == "" {
			clusterID = "—"
		}
		sb.WriteString(fmt.Sprintf("| %s | %s / %s | %d/%d | `%s` | %s |\n",
			r.DisplayName, r.CSP, r.Region, passed, total, clusterID, result))
	}
	sb.WriteString("\n## Per-CSP Reports\n\n")
	for _, r := range reports {
		sb.WriteString(fmt.Sprintf("- [%s](k8s-infra-migration-test-results-%s.md)\n", r.DisplayName, strings.ToLower(r.CSP)))
	}

	path := filepath.Join(dir, "k8s-infra-migration-test-summary-all.md")
	if err := os.WriteFile(path, []byte(maskSensitiveInfo(sb.String())), 0644); err != nil {
		return err
	}
	log.Info().Str("file", path).Msg("✅ Summary report generated and saved")
	return nil
}
