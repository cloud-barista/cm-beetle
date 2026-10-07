package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
)

// TestConfig holds test configuration for K8s test CLI across all modes.
type TestConfig struct {
	Test struct {
		ExecutionMode string `yaml:"executionMode"` // migration | recommendation | validation | all
		Set           struct {
			Mode              string `yaml:"mode"`              // parallel or sequential
			StartDelaySeconds int    `yaml:"startDelaySeconds"` // stagger between parallel starts
		} `yaml:"set"`
		Cases     []TestCase `yaml:"cases"`
		Scenarios []Scenario `yaml:"scenarios"`
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
		NodeReadyTimeoutSec  int  `yaml:"nodeReadyTimeoutSec"`
		NodePollSec          int  `yaml:"nodePollSec"`
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

// Scenario is one on-premise fixture plus the expectations it must satisfy.
type Scenario struct {
	File    string `yaml:"file"`
	Name    string `yaml:"name"`
	Execute bool   `yaml:"execute"`
	Expect  Expect `yaml:"expect"`
}

// Expect declares what a scenario's recommendation response must satisfy.
type Expect struct {
	StatusCode     int    `yaml:"statusCode"`
	NodeGroupCount *int   `yaml:"nodeGroupCount"`
	TotalNodeSize  *int   `yaml:"totalNodeSize"`
	VersionPrefix  string `yaml:"versionPrefix"`
}

// AuthConfig holds Beetle and Tumblebug credentials.
type AuthConfig struct {
	BeetleApiUsername    string `json:"beetleApiUsername"`
	BeetleApiPassword    string `json:"beetleApiPassword"`
	TumblebugApiUsername string `json:"tumblebugApiUsername"`
	TumblebugApiPassword string `json:"tumblebugApiPassword"`
	TumblebugEndpoint    string `json:"tumblebugEndpoint"`
}

// loadConfig reads and parses the YAML test configuration file.
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

// loadAuthConfig reads Beetle and Tumblebug credentials from JSON.
func loadAuthConfig(path string) (AuthConfig, error) {
	var auth AuthConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return auth, err
	}
	if err := json.Unmarshal(data, &auth); err != nil {
		return auth, err
	}
	return auth, nil
}

// applyDefaults populates missing config fields with safe default values.
func applyDefaults(cfg *TestConfig) {
	if cfg.Test.ExecutionMode == "" {
		cfg.Test.ExecutionMode = "migration"
	}
	if cfg.Test.Set.Mode == "" {
		cfg.Test.Set.Mode = "parallel"
	}
	if cfg.Migration.Mode == "" {
		cfg.Migration.Mode = "async"
	}
	if cfg.Migration.TimeoutSec == 0 {
		cfg.Migration.TimeoutSec = 2400
	}
	if cfg.Poll.IntervalSec == 0 {
		cfg.Poll.IntervalSec = 15
	}
	if cfg.Poll.TimeoutSec == 0 {
		cfg.Poll.TimeoutSec = 2400
	}
	if cfg.Delete.TimeoutSec == 0 {
		cfg.Delete.TimeoutSec = 1800
	}
	if cfg.Delete.RetryIntervalSec == 0 {
		cfg.Delete.RetryIntervalSec = 10
	}
	if cfg.Delete.MaxRetries == 0 {
		cfg.Delete.MaxRetries = 3
	}
	if cfg.Workload.KubeconfigTimeoutSec == 0 {
		cfg.Workload.KubeconfigTimeoutSec = 120
	}
	if cfg.Workload.KubeconfigPollSec == 0 {
		cfg.Workload.KubeconfigPollSec = 5
	}
	if cfg.Workload.NodeReadyTimeoutSec == 0 {
		cfg.Workload.NodeReadyTimeoutSec = 600
	}
	if cfg.Workload.NodePollSec == 0 {
		cfg.Workload.NodePollSec = 10
	}
	if cfg.Workload.PodReadyTimeoutSec == 0 {
		cfg.Workload.PodReadyTimeoutSec = 300
	}
	if cfg.Workload.PodPollSec == 0 {
		cfg.Workload.PodPollSec = 5
	}
	if cfg.Workload.LbAddressTimeoutSec == 0 {
		cfg.Workload.LbAddressTimeoutSec = 300
	}
	if cfg.Workload.LbAccessTimeoutSec == 0 {
		cfg.Workload.LbAccessTimeoutSec = 180
	}
	if cfg.Workload.LbPollSec == 0 {
		cfg.Workload.LbPollSec = 10
	}
}

// selectedTargets returns enabled target cases.
func selectedTargets(cfg TestConfig) []TestCase {
	var out []TestCase
	for _, c := range cfg.Test.Cases {
		if c.Execute {
			out = append(out, c)
		}
	}
	return out
}

// selectedScenarios returns enabled scenarios.
func selectedScenarios(cfg TestConfig) []Scenario {
	var out []Scenario
	for _, s := range cfg.Test.Scenarios {
		if s.Execute {
			out = append(out, s)
		}
	}
	return out
}

// caseNameSeed determines resource prefix for each test case to prevent collision.
func caseNameSeed(cfg TestConfig, idx int) string {
	if cfg.Beetle.NameSeed == "" {
		return ""
	}
	return fmt.Sprintf("%s%02d", cfg.Beetle.NameSeed, idx+1)
}
