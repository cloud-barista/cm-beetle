package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/validation"
)

// stepValidate calls the Pre-flight Validation API before provisioning.
func stepValidate(client *resty.Client, cfg TestConfig, auth AuthConfig, report *CSPTestReport, requestBody []byte) StepResult {
	res := StepResult{
		Target:    report.DisplayName,
		Number:    2,
		Name:      "POST /beetle/validation/ns/{nsId}/k8sCluster (Pre-flight validation)",
		StartTime: time.Now(),
	}

	if report.Recommendation == nil {
		res.Duration = time.Since(res.StartTime)
		res.Success = false
		res.Error = "recommendation is missing; step 1 must succeed first"
		return res
	}

	url := fmt.Sprintf("%s/beetle/validation/ns/%s/k8sCluster", cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(report.Recommendation).
		Post(url)

	res.Duration = time.Since(res.StartTime)
	if err != nil {
		res.Error = err.Error()
		return res
	}

	res.StatusCode = resp.StatusCode()
	if resp.StatusCode() != 200 {
		res.Error = fmt.Sprintf("HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
		return res
	}

	var apiResp model.ApiResponse[validation.ValidationResult]
	if err := json.Unmarshal(resp.Body(), &apiResp); err != nil {
		res.Error = fmt.Sprintf("failed to parse validation response: %s", err)
		return res
	}

	valResult := apiResp.Data
	if !valResult.Valid {
		res.Success = false
		res.Error = fmt.Sprintf("validation failed with %d issue(s)", len(valResult.Issues))
		for _, issue := range valResult.Issues {
			res.Notes = append(res.Notes, fmt.Sprintf("❌ [%s] %s: %s", issue.Code, issue.Path, issue.Message))
		}
		return res
	}

	res.Success = true
	res.Notes = append(res.Notes, "✅ Target K8s infra model is valid for migration (0 issues)")
	return res
}

// runValidationSuite executes validation-only suite without provisioning real resources.
func runValidationSuite(client *resty.Client, cfg TestConfig, auth AuthConfig, targets []TestCase, requestBody []byte) []*CSPTestReport {
	reports := make([]*CSPTestReport, len(targets))

	fmt.Println("=========================================================")
	fmt.Println(" CM-Beetle K8s Infra Validation Test Suite")
	fmt.Printf(" %d target(s), concurrency: %s (Recommend -> Validate only)\n", len(targets), cfg.Test.Set.Mode)
	fmt.Println("=========================================================")

	if cfg.Test.Set.Mode == "sequential" {
		for i, t := range targets {
			reports[i] = runValidationLifecycle(cfg, auth, t, requestBody, caseNameSeed(cfg, t, i))
		}
		return reports
	}

	var wg sync.WaitGroup
	for i, t := range targets {
		wg.Add(1)
		go func(idx int, target TestCase) {
			defer wg.Done()
			if cfg.Test.Set.StartDelaySeconds > 0 && idx > 0 {
				time.Sleep(time.Duration(idx*cfg.Test.Set.StartDelaySeconds) * time.Second)
			}
			reports[idx] = runValidationLifecycle(cfg, auth, target, requestBody, caseNameSeed(cfg, target, idx))
		}(i, t)
	}
	wg.Wait()
	return reports
}

// runValidationLifecycle runs recommend followed by pre-flight validation.
func runValidationLifecycle(cfg TestConfig, auth AuthConfig, target TestCase, requestBody []byte, nameSeed string) *CSPTestReport {
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

	printStepStart(report.DisplayName, 1, "POST /beetle/recommendation/k8sCluster")
	resRec := stepRecommend(client, cfg, auth, report, requestBody)
	report.Steps = append(report.Steps, resRec)
	printStep(report.DisplayName, resRec)

	if !resRec.Success {
		return report
	}

	printStepStart(report.DisplayName, 2, "POST /beetle/validation/ns/{nsId}/k8sCluster")
	resVal := stepValidate(client, cfg, auth, report, requestBody)
	report.Steps = append(report.Steps, resVal)
	printStep(report.DisplayName, resVal)

	return report
}
