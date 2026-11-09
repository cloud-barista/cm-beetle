package main

import (
	"fmt"
	"strings"
	"time"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

// Execution mode constants define whether steps call live migration APIs or simulation profiles.
const (
	ExecutionModeSimulate = "simulate"
	ExecutionModeReal     = "real"
)

// Failure classification codes categorize why a CSP failed during scale-up.
const (
	FailureCodeNone              = ""
	FailureCodeQuotaExceeded     = "QUOTA_EXCEEDED"
	FailureCodeRateLimited       = "API_RATE_LIMIT"
	FailureCodeCapacityExhausted = "AZ_CAPACITY_EXHAUSTED"
	FailureCodeSpecImageMismatch = "SPEC_IMAGE_CONSTRAINT"
	FailureCodeNetworkExhausted  = "SUBNET_IP_EXHAUSTED"
	FailureCodeTimeout           = "PROVISIONING_TIMEOUT"
	FailureCodeProvisionError    = "CSP_PROVISION_ERROR"
)

// StressTestConfig holds the full configuration for the multi-CSP stress test CLI.
type StressTestConfig struct {
	Stress StressSettings `yaml:"stress" json:"stress"`
	Test   TestTargets    `yaml:"test" json:"test"`
	Beetle BeetleSettings `yaml:"beetle" json:"beetle"`
}

// StressSettings defines equal-round scaling rules and random on-premise model parameters.
type StressSettings struct {
	ExecutionMode            string          `yaml:"executionMode" json:"executionMode"`
	InitialNodes             int             `yaml:"initialNodes" json:"initialNodes"`
	StepNodes                int             `yaml:"stepNodes" json:"stepNodes"`
	MaxNodes                 int             `yaml:"maxNodes" json:"maxNodes"`
	RandomSeed               int64           `yaml:"randomSeed" json:"randomSeed"`
	MergeIdenticalNodeGroups bool            `yaml:"mergeIdenticalNodeGroups" json:"mergeIdenticalNodeGroups"`
	NameSeedPrefix           string          `yaml:"nameSeedPrefix" json:"nameSeedPrefix"`
	Generator                GeneratorConfig `yaml:"generator" json:"generator"`
}

// GeneratorConfig controls the resource bounds for generated on-premise nodes.
type GeneratorConfig struct {
	MinVCPU     int `yaml:"minVcpu" json:"minVcpu"`
	MaxVCPU     int `yaml:"maxVcpu" json:"maxVcpu"`
	MinMemoryGB int `yaml:"minMemoryGb" json:"minMemoryGb"`
	MaxMemoryGB int `yaml:"maxMemoryGb" json:"maxMemoryGb"`
	MinDiskGB   int `yaml:"minDiskGb" json:"minDiskGb"`
	MaxDiskGB   int `yaml:"maxDiskGb" json:"maxDiskGb"`
}

// TestTargets lists the candidate CSP-Region pairs for the stress test.
type TestTargets struct {
	Cases []CSPTestCase `yaml:"cases" json:"cases"`
}

// CSPTestCase configures a single CSP target and its optional simulation profile.
type CSPTestCase struct {
	cloudmodel.CloudProperty `yaml:",inline"`
	Name                     string `yaml:"name" json:"name"`
	Execute                  bool   `yaml:"execute" json:"execute"`
	SimMaxNodes              int    `yaml:"simMaxNodes,omitempty" json:"simMaxNodes,omitempty"`
	SimFailureCode           string `yaml:"simFailureCode,omitempty" json:"simFailureCode,omitempty"`
	SimFailureReason         string `yaml:"simFailureReason,omitempty" json:"simFailureReason,omitempty"`
	SimBaseLatencySec        int    `yaml:"simBaseLatencySec,omitempty" json:"simBaseLatencySec,omitempty"`
}

// BeetleSettings holds the CM-Beetle connection and namespace configuration.
type BeetleSettings struct {
	Endpoint       string `yaml:"endpoint" json:"endpoint"`
	NamespaceID    string `yaml:"namespaceId" json:"namespaceId"`
	AuthConfigFile string `yaml:"authConfigFile" json:"authConfigFile"`
}

// AuthConfig holds HTTP Basic authentication credentials for CM-Beetle.
type AuthConfig struct {
	BeetleApiUsername string `json:"beetleApiUsername"`
	BeetleApiPassword string `json:"beetleApiPassword"`
}

// StepPhaseResult records timing and status for one phase within a CSP test step.
type StepPhaseResult struct {
	Phase        string    `json:"phase"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	DurationMs   int64     `json:"durationMs"`
	SimulatedSec float64   `json:"simulatedSec"`
	Success      bool      `json:"success"`
	Skipped      bool      `json:"skipped"`
	StatusCode   int       `json:"statusCode,omitempty"`
	ErrorMessage string    `json:"errorMessage,omitempty"`
	Details      string    `json:"details,omitempty"`
}

// CleanupSummary tracks the 2-stage resource deletion verification for a step.
type CleanupSummary struct {
	Attempted                bool     `json:"attempted"`
	InfraDeleted             bool     `json:"infraDeleted"`
	SecurityGroupsCleaned    int      `json:"securityGroupsCleaned"`
	SshKeysCleaned           int      `json:"sshKeysCleaned"`
	VNetsCleaned             int      `json:"vNetsCleaned"`
	OrphanResourcesRemaining int      `json:"orphanResourcesRemaining"`
	CleanVerified            bool     `json:"cleanVerified"`
	DurationMs               int64    `json:"durationMs"`
	SimulatedSec             float64  `json:"simulatedSec"`
	Logs                     []string `json:"logs"`
}

// StepAttempt captures the complete lifecycle of one CSP at a specific node scale.
type StepAttempt struct {
	SequenceIndex       int             `json:"sequenceIndex"`
	Round               int             `json:"round"`
	CSP                 string          `json:"csp"`
	Region              string          `json:"region"`
	DisplayName         string          `json:"displayName"`
	TargetNodes         int             `json:"targetNodes"`
	GeneratedNodeGroups int             `json:"generatedNodeGroups"`
	SecurityGroupCount  int             `json:"securityGroupCount"`
	TotalVCPUs          int             `json:"totalVcpus"`
	TotalMemoryGB       int             `json:"totalMemoryGb"`
	TotalDiskGB         int             `json:"totalDiskGb"`
	AvgMatchRate        float64         `json:"avgMatchRate"`
	NameSeed            string          `json:"nameSeed"`
	InfraID             string          `json:"infraId"`
	StartTime           time.Time       `json:"startTime"`
	EndTime             time.Time       `json:"endTime"`
	DurationMs          int64           `json:"durationMs"`
	SimulatedDurationS  float64         `json:"simulatedDurationSec"`
	ElapsedSimulatedS   float64         `json:"elapsedSimulatedSec"`
	Success             bool            `json:"success"`
	Eliminated          bool            `json:"eliminated"`
	FailurePhase        string          `json:"failurePhase,omitempty"`
	FailureCode         string          `json:"failureCode,omitempty"`
	FailureReason       string          `json:"failureReason,omitempty"`
	PlanPhase           StepPhaseResult `json:"planPhase"`
	RecommendPhase      StepPhaseResult `json:"recommendPhase"`
	MigratePhase        StepPhaseResult `json:"migratePhase"`
	CleanupPhase        StepPhaseResult `json:"cleanupPhase"`
	CleanupSummary      CleanupSummary  `json:"cleanupSummary"`
}

// CSPSurvivalStatus summarizes the final standing of a single CSP across all rounds.
type CSPSurvivalStatus struct {
	CSP                   string `json:"csp"`
	Region                string `json:"region"`
	DisplayName           string `json:"displayName"`
	Status                string `json:"status"`
	MaxSucceededNodes     int    `json:"maxSucceededNodes"`
	FailedAtNodes         int    `json:"failedAtNodes,omitempty"`
	EliminatedRound       int    `json:"eliminatedRound,omitempty"`
	EliminatedStepIndex   int    `json:"eliminatedStepIndex,omitempty"`
	FailureCode           string `json:"failureCode,omitempty"`
	FailureReason         string `json:"failureReason,omitempty"`
	CleanupVerified       bool   `json:"cleanupVerified"`
	TotalStepsAttempted   int    `json:"totalStepsAttempted"`
	TotalStepsSucceeded   int    `json:"totalStepsSucceeded"`
	TotalNodesProvisioned int    `json:"totalNodesProvisioned"`
	TotalNodesCleaned     int    `json:"totalNodesCleaned"`
}

// StressTestReport is the root JSON payload persisted and rendered into the dashboard.
type StressTestReport struct {
	ReportVersion         string              `json:"reportVersion"`
	TestDate              string              `json:"testDate"`
	TestTime              string              `json:"testTime"`
	StartTime             time.Time           `json:"startTime"`
	EndTime               time.Time           `json:"endTime"`
	WallClockDurationMs   int64               `json:"wallClockDurationMs"`
	TotalSimulatedSec     float64             `json:"totalSimulatedSec"`
	ExecutionMode         string              `json:"executionMode"`
	InitialNodes          int                 `json:"initialNodes"`
	StepNodes             int                 `json:"stepNodes"`
	MaxNodes              int                 `json:"maxNodes"`
	MaxReachedNodes       int                 `json:"maxReachedNodes"`
	TotalRounds           int                 `json:"totalRounds"`
	TotalSteps            int                 `json:"totalSteps"`
	TotalNodesProvisioned int                 `json:"totalNodesProvisioned"`
	TotalNodesCleaned     int                 `json:"totalNodesCleaned"`
	OrphanResourcesLeft   int                 `json:"orphanResourcesLeft"`
	BeetleEndpoint        string              `json:"beetleEndpoint"`
	NamespaceID           string              `json:"namespaceId"`
	CSPStatuses           []CSPSurvivalStatus `json:"cspStatuses"`
	Steps                 []StepAttempt       `json:"steps"`
}

// Validate checks all configuration fields explicitly and fails fast on invalid values.
func (c *StressTestConfig) Validate() error {
	execMode := strings.ToLower(strings.TrimSpace(c.Stress.ExecutionMode))
	if execMode != ExecutionModeSimulate && execMode != ExecutionModeReal {
		return fmt.Errorf("invalid stress.executionMode %q: must be %q or %q", c.Stress.ExecutionMode, ExecutionModeSimulate, ExecutionModeReal)
	}
	c.Stress.ExecutionMode = execMode

	if c.Stress.InitialNodes <= 0 {
		return fmt.Errorf("stress.initialNodes must be > 0 (got %d)", c.Stress.InitialNodes)
	}
	if c.Stress.StepNodes <= 0 {
		return fmt.Errorf("stress.stepNodes must be > 0 (got %d)", c.Stress.StepNodes)
	}
	if c.Stress.MaxNodes < c.Stress.InitialNodes {
		return fmt.Errorf("stress.maxNodes (%d) must be >= stress.initialNodes (%d)", c.Stress.MaxNodes, c.Stress.InitialNodes)
	}

	prefix := strings.TrimSpace(c.Stress.NameSeedPrefix)
	if prefix == "" {
		return fmt.Errorf("stress.nameSeedPrefix is required")
	}
	if ok, detail := common.IsValidNameSeed(prefix); !ok {
		return fmt.Errorf("invalid stress.nameSeedPrefix %q: %s", prefix, detail)
	}
	if len(prefix) > 6 {
		return fmt.Errorf("stress.nameSeedPrefix %q must be 6 characters or fewer so composed round seeds stay within 20 characters", prefix)
	}
	c.Stress.NameSeedPrefix = strings.ToLower(prefix)

	g := c.Stress.Generator
	if g.MinVCPU <= 0 || g.MaxVCPU < g.MinVCPU {
		return fmt.Errorf("invalid stress.generator vCPU range: min=%d, max=%d", g.MinVCPU, g.MaxVCPU)
	}
	if g.MinMemoryGB <= 0 || g.MaxMemoryGB < g.MinMemoryGB {
		return fmt.Errorf("invalid stress.generator memory range: min=%d, max=%d", g.MinMemoryGB, g.MaxMemoryGB)
	}
	if g.MinDiskGB <= 0 || g.MaxDiskGB < g.MinDiskGB {
		return fmt.Errorf("invalid stress.generator disk range: min=%d, max=%d", g.MinDiskGB, g.MaxDiskGB)
	}

	activeCount := 0
	for i, tc := range c.Test.Cases {
		if !tc.Execute {
			continue
		}
		if strings.TrimSpace(tc.Csp) == "" || strings.TrimSpace(tc.Region) == "" {
			return fmt.Errorf("test.cases[%d] has empty csp or region", i)
		}
		activeCount++
	}
	if activeCount == 0 {
		return fmt.Errorf("at least one entry in test.cases must have execute: true")
	}

	if c.Stress.ExecutionMode == ExecutionModeReal {
		if strings.TrimSpace(c.Beetle.Endpoint) == "" {
			return fmt.Errorf("beetle.endpoint is required when executionMode is %q", ExecutionModeReal)
		}
		if strings.TrimSpace(c.Beetle.NamespaceID) == "" {
			return fmt.Errorf("beetle.namespaceId is required when executionMode is %q", ExecutionModeReal)
		}
	}

	return nil
}

// ClassifyFailureReason inspects an HTTP status code and error string to assign a clear failure category.
func ClassifyFailureReason(statusCode int, rawErr string) (string, string) {
	lower := strings.ToLower(rawErr)

	switch {
	case statusCode == 429 || statusCode == 503 || strings.Contains(lower, "rate limit") || strings.Contains(lower, "too many requests") || strings.Contains(lower, "throttl"):
		return FailureCodeRateLimited, "CSP or gateway API rate limit reached during concurrent resource provisioning"
	case strings.Contains(lower, "quota") || strings.Contains(lower, "vcpulimitexceeded") || strings.Contains(lower, "limitexceeded") || strings.Contains(lower, "maxvcpu"):
		return FailureCodeQuotaExceeded, "Regional vCPU or instance count quota limit exceeded on target CSP account"
	case strings.Contains(lower, "capacity") || strings.Contains(lower, "insufficientinstancecapacity") || strings.Contains(lower, "out of stock") || strings.Contains(lower, "zone"):
		return FailureCodeCapacityExhausted, "Target availability zone exhausted for the requested Node Spec"
	case strings.Contains(lower, "subnet") || strings.Contains(lower, "ip address") || strings.Contains(lower, "cidr"):
		return FailureCodeNetworkExhausted, "Subnet IP address pool or VNIC attachment limit exhausted"
	case strings.Contains(lower, "image") || strings.Contains(lower, "spec") || strings.Contains(lower, "disk") || strings.Contains(lower, "nothing-to-recommend"):
		return FailureCodeSpecImageMismatch, "Compatible Node Spec or OS Image preflight constraint failed"
	case strings.Contains(lower, "timeout") || strings.Contains(lower, "deadline exceeded"):
		return FailureCodeTimeout, "CSP provisioning exceeded maximum wait timeout at high node concurrency"
	default:
		if rawErr != "" {
			return FailureCodeProvisionError, rawErr
		}
		return FailureCodeProvisionError, fmt.Sprintf("CSP provisioning failed (HTTP %d)", statusCode)
	}
}
