package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v3"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
)

// restyNoopLogger silences verbose Resty debug output.
type restyNoopLogger struct{}

func (restyNoopLogger) Errorf(_ string, _ ...interface{}) {}
func (restyNoopLogger) Warnf(_ string, _ ...interface{})  {}
func (restyNoopLogger) Debugf(_ string, _ ...interface{}) {}

// TestConfig defines the root YAML configuration for AIMX.
type TestConfig struct {
	Test struct {
		ExecutionMode string     `yaml:"executionMode"`
		Cases         []TestCase `yaml:"cases"`
	} `yaml:"test"`
	Beetle    BeetleConfig    `yaml:"beetle"`
	Source    SourceConfig    `yaml:"source"`
	AIService AIServiceConfig `yaml:"aiService"`
	Timing    TimingConfig    `yaml:"timing"`
}

// TestCase represents one target CSP and region migration target.
type TestCase struct {
	cloudmodel.CloudProperty `yaml:",inline"`
	Name                     string `yaml:"name"`
	Execute                  bool   `yaml:"execute"`
	NameSeed                 string `yaml:"nameSeed,omitempty"`
}

// BeetleConfig holds CM-Beetle endpoint and file paths.
type BeetleConfig struct {
	Endpoint       string `yaml:"endpoint"`
	NamespaceID    string `yaml:"namespaceId"`
	AuthConfigFile string `yaml:"authConfigFile"`
	BaseInfraFile  string `yaml:"baseInfraFile"`
	StateFile      string `yaml:"stateFile"`
}

// SourceConfig holds CB-Tumblebug source GPU VM settings for direct infraDynamic provisioning.
type SourceConfig struct {
	NamespaceID            string `yaml:"namespaceId"`
	InfraID                string `yaml:"infraId"`
	NodeGroupName          string `yaml:"nodeGroupName"`
	SpecID                 string `yaml:"specId"`
	ImageID                string `yaml:"imageId"`
	RootDiskType           string `yaml:"rootDiskType"`
	RootDiskSize           int    `yaml:"rootDiskSize"`
	AutoProvisionIfMissing bool   `yaml:"autoProvisionIfMissing"`
	OpenServicePorts       bool   `yaml:"openServicePortsOnSource"`
	CleanupSourceInfra     bool   `yaml:"cleanupSourceInfra"`
}

// TransxFilterConfig maps custom include/exclude glob patterns to transx.FilterOption.
type TransxFilterConfig struct {
	Include []string `yaml:"include"`
	Exclude []string `yaml:"exclude"`
}

// AIServiceConfig defines the Gemma model, Open WebUI stack, and transx transfer parameters.
type AIServiceConfig struct {
	RootDir              string             `yaml:"rootDir"`
	ModelName            string             `yaml:"modelName"`
	OllamaPort           int                `yaml:"ollamaPort"`
	OpenWebUIPort        int                `yaml:"openWebUIPort"`
	TestPrompt           string             `yaml:"testPrompt"`
	TransferModelWeights bool               `yaml:"transferModelWeights"`
	TransxStrategy       string             `yaml:"transxStrategy"`
	Filter               TransxFilterConfig `yaml:"filter"`
	UseDataEncryption    bool               `yaml:"useDataEncryption"`
}

// TimingConfig controls stabilization delays and readiness polling intervals across stages.
type TimingConfig struct {
	StepTransitionSeconds     int `yaml:"stepTransitionSeconds"`
	VmBootStabilizeSeconds    int `yaml:"vmBootStabilizeSeconds"`
	SshRetryMaxAttempts       int `yaml:"sshRetryMaxAttempts"`
	SshRetryIntervalSeconds   int `yaml:"sshRetryIntervalSeconds"`
	ServiceInitWaitSeconds    int `yaml:"serviceInitWaitSeconds"`
	DbFlushWaitSeconds        int `yaml:"dbFlushWaitSeconds"`
	TargetServiceReadySeconds int `yaml:"targetServiceReadyWaitSeconds"`
	RemoteCmdTimeoutMinutes   int `yaml:"remoteCmdTimeoutMinutes"`
}

// AuthConfig holds Basic Auth credentials and endpoints for CM-Beetle and CB-Tumblebug.
type AuthConfig struct {
	BeetleApiUsername    string `json:"beetleApiUsername"`
	BeetleApiPassword    string `json:"beetleApiPassword"`
	TumblebugApiUsername string `json:"tumblebugApiUsername"`
	TumblebugApiPassword string `json:"tumblebugApiPassword"`
	TumblebugEndpoint    string `json:"tumblebugEndpoint"`
}

// EndpointMetrics captures hardware, model inference, and application metrics on a GPU VM.
type EndpointMetrics struct {
	InfraID                string  `json:"infraId"`
	NodeID                 string  `json:"nodeId"`
	PublicIP               string  `json:"publicIp"`
	SshKeyID               string  `json:"sshKeyId"`
	SshUsername            string  `json:"sshUsername"`
	Csp                    string  `json:"csp"`
	Region                 string  `json:"region"`
	SpecID                 string  `json:"specId"`
	ImageID                string  `json:"imageId"`
	GpuModel               string  `json:"gpuModel"`
	GpuVendor              string  `json:"gpuVendor"`
	GpuCount               int     `json:"gpuCount"`
	GpuVramGB              float32 `json:"gpuVramGb"`
	DriverVersion          string  `json:"driverVersion"`
	CudaVersion            string  `json:"cudaVersion"`
	ModelName              string  `json:"modelName"`
	OllamaProcessor        string  `json:"ollamaProcessor"`
	EvalTokensPerSec       float64 `json:"evalTokensPerSec"`
	SampleResponse         string  `json:"sampleResponse"`
	WebUIUrl               string  `json:"webUiUrl"`
	WebUIVersion           string  `json:"webUiVersion"`
	WebUIConfigOK          bool    `json:"webUiConfigOk"`
	ExternalWebUIReachable bool    `json:"externalWebUiReachable"`
	WebUIDbSha256          string  `json:"webUiDbSha256"`
	WebUIHttpStatus        int     `json:"webUiHttpStatus"`
	VolumeSizeHuman        string  `json:"volumeSizeHuman"`
}

// MigrationState persists intermediate artifacts across individual CLI mode invocations.
type MigrationState struct {
	UpdatedAt              time.Time                     `json:"updatedAt"`
	Source                 EndpointMetrics               `json:"source"`
	DiscoveredOnpremInfra  onpremmodel.OnpremInfra       `json:"discoveredOnpremInfra"`
	RecommendedCandidates  []cloudmodel.RecommendedInfra `json:"recommendedCandidates"`
	SelectedRecommendation *cloudmodel.RecommendedInfra  `json:"selectedRecommendation,omitempty"`
	TargetInfra            *cloudmodel.VmInfraInfo       `json:"targetInfra,omitempty"`
	Target                 EndpointMetrics               `json:"target"`
	DataTransferDuration   string                        `json:"dataTransferDuration"`
	DataTransferStrategy   string                        `json:"dataTransferStrategy"`
	StageDurations         map[string]string             `json:"stageDurations,omitempty"`
}

// loadConfig reads, parses, and validates the YAML configuration file without fallback.
func loadConfig(path string) (TestConfig, error) {
	var cfg TestConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("failed to read config file %s: %w", path, err)
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse YAML config %s: %w", path, err)
	}
	if err := validateConfig(cfg); err != nil {
		return cfg, fmt.Errorf("invalid config in %s: %w", path, err)
	}
	return cfg, nil
}

// validateConfig verifies that all required configuration fields are explicitly provided.
func validateConfig(cfg TestConfig) error {
	requiredStrings := map[string]string{
		"test.executionMode":       cfg.Test.ExecutionMode,
		"beetle.endpoint":          cfg.Beetle.Endpoint,
		"beetle.namespaceId":       cfg.Beetle.NamespaceID,
		"beetle.authConfigFile":    cfg.Beetle.AuthConfigFile,
		"beetle.baseInfraFile":     cfg.Beetle.BaseInfraFile,
		"beetle.stateFile":         cfg.Beetle.StateFile,
		"source.namespaceId":       cfg.Source.NamespaceID,
		"source.infraId":           cfg.Source.InfraID,
		"source.nodeGroupName":     cfg.Source.NodeGroupName,
		"source.specId":            cfg.Source.SpecID,
		"source.imageId":           cfg.Source.ImageID,
		"aiService.rootDir":        cfg.AIService.RootDir,
		"aiService.modelName":      cfg.AIService.ModelName,
		"aiService.testPrompt":     cfg.AIService.TestPrompt,
		"aiService.transxStrategy": cfg.AIService.TransxStrategy,
	}
	for field, val := range requiredStrings {
		if val == "" {
			return fmt.Errorf("required field '%s' is missing or empty", field)
		}
	}
	if cfg.Source.RootDiskSize <= 0 {
		return fmt.Errorf("source.rootDiskSize (%d) must be a positive integer", cfg.Source.RootDiskSize)
	}
	if cfg.AIService.OllamaPort <= 0 || cfg.AIService.OpenWebUIPort <= 0 {
		return fmt.Errorf("aiService.ollamaPort (%d) and aiService.openWebUIPort (%d) must be positive integers",
			cfg.AIService.OllamaPort, cfg.AIService.OpenWebUIPort)
	}
	if cfg.Timing.SshRetryMaxAttempts <= 0 || cfg.Timing.SshRetryIntervalSeconds <= 0 || cfg.Timing.RemoteCmdTimeoutMinutes <= 0 {
		return fmt.Errorf("timing.sshRetryMaxAttempts, timing.sshRetryIntervalSeconds, and timing.remoteCmdTimeoutMinutes must be > 0")
	}
	return nil
}

// loadAuthConfig reads Beetle and Tumblebug credentials from JSON without fallback.
func loadAuthConfig(path string) (AuthConfig, error) {
	var auth AuthConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return auth, fmt.Errorf("failed to read auth config %s: %w", path, err)
	}
	if err := json.Unmarshal(data, &auth); err != nil {
		return auth, fmt.Errorf("failed to parse auth JSON %s: %w", path, err)
	}
	if auth.TumblebugEndpoint == "" {
		return auth, fmt.Errorf("tumblebugEndpoint is required in %s", path)
	}
	return auth, nil
}

// selectedTargets returns all TestCase items marked with Execute: true.
func selectedTargets(cfg TestConfig) []TestCase {
	var out []TestCase
	for _, tc := range cfg.Test.Cases {
		if tc.Execute {
			out = append(out, tc)
		}
	}
	return out
}

// newBeetleClient creates a Resty HTTP client configured for CM-Beetle API requests.
func newBeetleClient(auth AuthConfig) *resty.Client {
	c := resty.New().SetTimeout(45 * time.Minute).SetLogger(restyNoopLogger{})
	if auth.BeetleApiUsername != "" {
		c.SetBasicAuth(auth.BeetleApiUsername, auth.BeetleApiPassword)
	}
	return c
}

// newTumblebugClient creates a Resty HTTP client configured for CB-Tumblebug API requests.
func newTumblebugClient(auth AuthConfig) *resty.Client {
	c := resty.New().SetTimeout(45 * time.Minute).SetLogger(restyNoopLogger{})
	if auth.TumblebugApiUsername != "" {
		c.SetBasicAuth(auth.TumblebugApiUsername, auth.TumblebugApiPassword)
	}
	return c
}

// loadState loads the persisted MigrationState from disk or initializes a new one.
func loadState(path string) (*MigrationState, error) {
	st := &MigrationState{
		StageDurations: make(map[string]string),
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return st, nil
		}
		return nil, fmt.Errorf("failed to read state file %s: %w", path, err)
	}
	if err := json.Unmarshal(data, st); err != nil {
		return nil, fmt.Errorf("failed to parse state file %s: %w", path, err)
	}
	if st.StageDurations == nil {
		st.StageDurations = make(map[string]string)
	}
	return st, nil
}

// saveState writes the current MigrationState to disk for cross-mode continuity.
func saveState(path string, st *MigrationState) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create state directory: %w", err)
	}
	st.UpdatedAt = time.Now()
	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal migration state: %w", err)
	}
	return os.WriteFile(path, data, 0644)
}
