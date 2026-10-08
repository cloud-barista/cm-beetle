package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
)

// restyNoopLogger silences verbose Resty log messages.
type restyNoopLogger struct{}

func (restyNoopLogger) Errorf(_ string, _ ...interface{}) {}
func (restyNoopLogger) Warnf(_ string, _ ...interface{})  {}
func (restyNoopLogger) Debugf(_ string, _ ...interface{}) {}

var (
	configFile = flag.String("config", "testconf/test-config.yaml", "Path to config file")
	modeFlag   = flag.String("mode", "", "Execution mode: migration (default), recommendation, validation, all")
)

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

// newClient creates configured HTTP client with credentials and silent logger.
func newClient(cfg TestConfig, auth AuthConfig) *resty.Client {
	c := resty.New().SetTimeout(2 * time.Minute).SetLogger(restyNoopLogger{})
	if auth.BeetleApiUsername != "" {
		c.SetBasicAuth(auth.BeetleApiUsername, auth.BeetleApiPassword)
	}
	return c
}

// checkBeetleReadiness verifies CM-Beetle readyz endpoint.
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

func main() {
	flag.Parse()

	cfg, err := loadConfig(*configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	if *modeFlag != "" {
		cfg.Test.ExecutionMode = *modeFlag
	}
	applyDefaults(&cfg)

	auth, err := loadAuthConfig(cfg.Beetle.AuthConfigFile)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to load auth config; proceeding without auth")
	}

	client := newClient(cfg, auth)
	if err := checkBeetleReadiness(client, cfg.Beetle.Endpoint); err != nil {
		log.Fatal().Err(err).Msg("CM-Beetle readiness check failed")
	}

	execMode := strings.ToLower(cfg.Test.ExecutionMode)
	switch execMode {
	case "recommendation", "rec":
		targets := selectedTargets(cfg)
		if len(targets) == 0 {
			log.Fatal().Msg("At least 1 test case must have execute: true")
		}
		scenarios := selectedScenarios(cfg)
		if len(scenarios) == 0 {
			log.Fatal().Msg("At least 1 scenario must have execute: true")
		}

		report := runRecommendationSuite(client, cfg, auth, targets, scenarios)
		if err := generateRecommendationMarkdownReport(report); err != nil {
			log.Warn().Err(err).Msg("Failed to generate recommendation markdown report")
		}
		if !printRecommendationFinalSummary(report) {
			os.Exit(1)
		}

	case "validation", "val":
		targets := selectedTargets(cfg)
		if len(targets) == 0 {
			log.Fatal().Msg("At least 1 test case must have execute: true")
		}
		requestBody, err := os.ReadFile(cfg.Beetle.RequestBodyFile)
		if err != nil {
			log.Fatal().Err(err).Str("file", cfg.Beetle.RequestBodyFile).Msg("Failed to read request body file")
		}

		reports := runValidationSuite(client, cfg, auth, targets, requestBody)
		for _, r := range reports {
			if err := generateCSPReport(r); err != nil {
				log.Warn().Err(err).Str("csp", r.CSP).Msg("Failed to generate CSP report")
			}
		}
		if err := generateSummaryReport(reports); err != nil {
			log.Warn().Err(err).Msg("Failed to generate summary report")
		}
		if !printMigrationFinalSummary(reports) {
			os.Exit(1)
		}

	case "migration", "mig":
		targets := selectedTargets(cfg)
		if len(targets) == 0 {
			log.Fatal().Msg("At least 1 test case must have execute: true")
		}
		requestBody, err := os.ReadFile(cfg.Beetle.RequestBodyFile)
		if err != nil {
			log.Fatal().Err(err).Str("file", cfg.Beetle.RequestBodyFile).Msg("Failed to read request body file")
		}

		reports := runMigrationSuite(cfg, auth, targets, requestBody)
		for _, r := range reports {
			if err := generateCSPReport(r); err != nil {
				log.Warn().Err(err).Str("csp", r.CSP).Msg("Failed to generate CSP report")
			}
		}
		if err := generateSummaryReport(reports); err != nil {
			log.Warn().Err(err).Msg("Failed to generate summary report")
		}
		if !printMigrationFinalSummary(reports) {
			os.Exit(1)
		}

	case "all":
		targets := selectedTargets(cfg)
		if len(targets) == 0 {
			log.Fatal().Msg("At least 1 test case must have execute: true")
		}
		scenarios := selectedScenarios(cfg)
		if len(scenarios) == 0 {
			log.Fatal().Msg("At least 1 scenario must have execute: true")
		}
		requestBody, err := os.ReadFile(cfg.Beetle.RequestBodyFile)
		if err != nil {
			log.Fatal().Err(err).Str("file", cfg.Beetle.RequestBodyFile).Msg("Failed to read request body file")
		}

		recReport := runRecommendationSuite(client, cfg, auth, targets, scenarios)
		if err := generateRecommendationMarkdownReport(recReport); err != nil {
			log.Warn().Err(err).Msg("Failed to generate recommendation markdown report")
		}
		if !printRecommendationFinalSummary(recReport) {
			log.Error().Msg("Recommendation phase failed; aborting migration phase")
			os.Exit(1)
		}

		migReports := runMigrationSuite(cfg, auth, targets, requestBody)
		for _, r := range migReports {
			if err := generateCSPReport(r); err != nil {
				log.Warn().Err(err).Str("csp", r.CSP).Msg("Failed to generate CSP report")
			}
		}
		if err := generateSummaryReport(migReports); err != nil {
			log.Warn().Err(err).Msg("Failed to generate summary report")
		}
		if !printMigrationFinalSummary(migReports) {
			os.Exit(1)
		}

	default:
		log.Fatal().Msgf("Unknown execution mode: %s (supported: migration, recommendation, validation, all)", cfg.Test.ExecutionMode)
	}
}
