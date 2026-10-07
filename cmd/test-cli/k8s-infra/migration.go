package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
)

// StepResult holds one lifecycle step outcome.
type StepResult struct {
	Target     string
	Number     int
	Name       string
	StartTime  time.Time
	Duration   time.Duration
	Success    bool
	Skipped    bool
	StatusCode int
	Notes      []string
	Error      string
}

// CSPTestReport captures the full lifecycle for one CSP and region pair.
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

// ApiResponseRaw mirrors model.ApiResponse for payloads parsed in two stages.
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

// RequestDetails mirrors common.RequestDetails fields.
type RequestDetails struct {
	Status        string          `json:"status"`
	ResponseData  json.RawMessage `json:"responseData"`
	ErrorResponse string          `json:"errorResponse"`
}

// runMigrationSuite executes the full E2E migration lifecycle against real CSPs.
func runMigrationSuite(cfg TestConfig, auth AuthConfig, targets []TestCase, requestBody []byte) []*CSPTestReport {
	reports := make([]*CSPTestReport, len(targets))

	fmt.Println("=========================================================")
	fmt.Println(" CM-Beetle K8s Infra Migration Test Suite")
	fmt.Printf(" %d target(s), concurrency: %s, migration mode: %s\n", len(targets), cfg.Test.Set.Mode, cfg.Migration.Mode)
	fmt.Println(" Recommend -> Validate -> Migrate -> List -> Get -> Workload -> Delete -> Residual")
	fmt.Println("=========================================================")

	if cfg.Test.Set.Mode == "sequential" {
		for i, t := range targets {
			reports[i] = runLifecycle(cfg, auth, t, requestBody, caseNameSeed(cfg, i))
		}
		return reports
	}

	var wg sync.WaitGroup
	for i, t := range targets {
		wg.Add(1)
		go func(idx int, target TestCase) {
			defer wg.Done()
			if d := cfg.Test.Set.StartDelaySeconds; d > 0 && idx > 0 {
				time.Sleep(time.Duration(idx*d) * time.Second)
			}
			reports[idx] = runLifecycle(cfg, auth, target, requestBody, caseNameSeed(cfg, idx))
		}(i, t)
	}
	wg.Wait()
	return reports
}

// runLifecycle executes the ordered steps for one target.
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

	steps := []struct {
		number int
		name   string
		run    func(*resty.Client, TestConfig, AuthConfig, *CSPTestReport, []byte) StepResult
	}{
		{1, "POST /recommendation/k8sCluster", stepRecommend},
		{2, "POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)", stepValidate},
		{3, "POST /migration/ns/{nsId}/k8sCluster", stepMigrate},
		{4, "GET /migration/ns/{nsId}/k8sCluster", stepList},
		{5, "GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation", stepGetAndVerify},
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

	if report.ClusterID != "" && cfg.Workload.Enabled && lastStepPassed(report) {
		runStep(6, "Workload verification (kubeconfig -> K8s API -> nginx)", stepWorkload)
	}

	if report.ClusterID != "" {
		res := runStep(7, "DELETE /migration/ns/{nsId}/k8sCluster/{id}", stepDelete)
		if res.Success && cfg.Verify.ResidualResources {
			runStep(8, "Residual resource check (Tumblebug)", stepResidualCheck)
		}
	}

	return report
}

// stepRecommend calls the recommendation API for a given target.
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

// stepMigrate provisions the cluster via sync or async API path.
func stepMigrate(client *resty.Client, cfg TestConfig, _ AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 3, Name: "POST /migration/ns/{nsId}/k8sCluster", StartTime: time.Now()}

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
		return res
	}
	res.Success = true
	return res
}

// migrateSync sends a blocking migration request.
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

// migrateAsync sends an async migration request and polls until completion.
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

	var apiResp ApiResponseRaw
	if err := json.Unmarshal(details.ResponseData, &apiResp); err == nil && len(apiResp.Data) > 0 {
		return apiResp.Data, resp.StatusCode(), nil
	}
	return details.ResponseData, resp.StatusCode(), nil
}

// extractReqID extracts request ID from header or response body.
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

// pollUntilDone polls request endpoint until status is Success or Error.
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

// adoptOrphanCluster attempts to find partially created cluster for cleanup.
func adoptOrphanCluster(client *resty.Client, cfg TestConfig, report *CSPTestReport, res *StepResult) {
	if report.ClusterID != "" || report.Recommendation == nil {
		return
	}
	wantName := report.Recommendation.TargetK8sCluster.Name
	if wantName == "" {
		return
	}
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

// stepList queries the cluster ID list.
func stepList(client *resty.Client, cfg TestConfig, _ AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 4, Name: "GET /migration/ns/{nsId}/k8sCluster", StartTime: time.Now()}

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

// stepGetAndVerify checks that provisioned cluster matches recommendation.
func stepGetAndVerify(client *resty.Client, cfg TestConfig, auth AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 5, Name: "GET /migration/ns/{nsId}/k8sCluster/{id} + verify vs recommendation", StartTime: time.Now()}

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

	var specSession *tbclient.Session
	if auth.TumblebugEndpoint != "" {
		specSession = tbclient.NewClient(tbclient.ApiConfig{
			RestUrl:  auth.TumblebugEndpoint + "/tumblebug",
			Username: auth.TumblebugApiUsername,
			Password: auth.TumblebugApiPassword,
		}).NewSession()
	}

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

// specMatch checks if actual spec ID matches recommended spec name.
func specMatch(sess *tbclient.Session, got, want string) (bool, error) {
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

// stepDelete deletes the provisioned cluster and retries if needed.
func stepDelete(client *resty.Client, cfg TestConfig, _ AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 7, Name: "DELETE /migration/ns/{nsId}/k8sCluster/{id}", StartTime: time.Now()}

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

// isAlreadyGone checks if error indicates resource is already deleted.
func isAlreadyGone(body string) bool {
	lower := strings.ToLower(body)
	return strings.Contains(lower, "not exist") || strings.Contains(lower, "not found")
}

// stepResidualCheck checks whether prerequisite resources survived cluster deletion.
func stepResidualCheck(_ *resty.Client, cfg TestConfig, auth AuthConfig, report *CSPTestReport, _ []byte) StepResult {
	res := StepResult{Target: report.DisplayName, Number: 8, Name: "Residual resource check (Tumblebug)", StartTime: time.Now()}

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
	res.Success = true
	return res
}

// lastStepPassed checks if the last step in report succeeded.
func lastStepPassed(report *CSPTestReport) bool {
	if len(report.Steps) == 0 {
		return false
	}
	return report.Steps[len(report.Steps)-1].Success
}
