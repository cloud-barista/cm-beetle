package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
)

func main() {
	configPath := flag.String("config", "testconf/test-config.yaml", "Path to AIMX YAML configuration file")
	modeFlag := flag.String("mode", "", "Execution mode override: seq | all | setup-source | recommend | migrate | verify | cleanup")
	cleanupSourceFlag := flag.Bool("cleanup-source", false, "Also terminate Source GPU Infra and default resources during -mode=cleanup")
	flag.Parse()

	fmt.Println("================================================================================")
	fmt.Println("🤖 AIMX (AI Migration X) — AI Semiconductor & AI Service Migration Executor")
	fmt.Println("================================================================================")

	cfg, err := loadConfig(*configPath)
	if err != nil {
		log.Fatalf("❌ Configuration Error: %v", err)
	}
	if strings.TrimSpace(*modeFlag) != "" {
		cfg.Test.ExecutionMode = strings.ToLower(strings.TrimSpace(*modeFlag))
	}
	if *cleanupSourceFlag {
		cfg.Source.CleanupSourceInfra = true
	}

	auth, err := loadAuthConfig(cfg.Beetle.AuthConfigFile)
	if err != nil {
		log.Fatalf("❌ Auth Configuration Error: %v", err)
	}

	targets := selectedTargets(cfg)
	if len(targets) == 0 && cfg.Test.ExecutionMode != "setup-source" && cfg.Test.ExecutionMode != "cleanup" {
		log.Fatalf("❌ No target cloud cases have 'execute: true' in %s", *configPath)
	}

	beetleClient := newBeetleClient(auth)
	tbClient := newTumblebugClient(auth)

	state, err := loadState(cfg.Beetle.StateFile)
	if err != nil {
		log.Fatalf("❌ Failed to load state file %s: %v", cfg.Beetle.StateFile, err)
	}

	mode := strings.ToLower(cfg.Test.ExecutionMode)
	fmt.Printf("⚙️  Execution Mode: %s | Source Infra: %s (ns: %s) | Model: %s\n",
		mode, cfg.Source.InfraID, cfg.Source.NamespaceID, cfg.AIService.ModelName)

	switch mode {
	case "setup-source":
		if err := setupSourceEnvironment(tbClient, cfg, auth, state); err != nil {
			log.Fatalf("❌ [setup-source] Failed: %v", err)
		}

	case "recommend":
		if err := checkBeetleReady(beetleClient, cfg.Beetle.Endpoint); err != nil {
			log.Fatalf("❌ CM-Beetle Readiness Error: %v", err)
		}
		for _, target := range targets {
			if _, err := recommendTargetGpuInfra(beetleClient, tbClient, cfg, auth, target, state); err != nil {
				log.Fatalf("❌ [recommend] Failed for %s/%s: %v", target.Csp, target.Region, err)
			}
		}

	case "migrate":
		if err := checkBeetleReady(beetleClient, cfg.Beetle.Endpoint); err != nil {
			log.Fatalf("❌ CM-Beetle Readiness Error: %v", err)
		}
		for _, target := range targets {
			if err := migrateInfraAndAIService(beetleClient, tbClient, cfg, auth, target, state); err != nil {
				log.Fatalf("❌ [migrate] Failed for %s/%s: %v", target.Csp, target.Region, err)
			}
		}

	case "verify":
		for _, target := range targets {
			if err := verifyMigratedAIService(tbClient, cfg, auth, target, state); err != nil {
				log.Fatalf("❌ [verify] Failed for %s/%s: %v", target.Csp, target.Region, err)
			}
		}

	case "cleanup":
		if err := checkBeetleReady(beetleClient, cfg.Beetle.Endpoint); err != nil {
			log.Fatalf("❌ CM-Beetle Readiness Error: %v", err)
		}
		if err := cleanupEnvironment(beetleClient, tbClient, cfg, auth, state); err != nil {
			log.Fatalf("❌ [cleanup] Failed: %v", err)
		}

	case "seq", "all":
		if err := checkBeetleReady(beetleClient, cfg.Beetle.Endpoint); err != nil {
			log.Fatalf("❌ CM-Beetle Readiness Error: %v", err)
		}
		if err := runSequentialPipeline(beetleClient, tbClient, cfg, auth, targets, state); err != nil {
			log.Fatalf("❌ [seq] Pipeline Failed: %v", err)
		}

	default:
		log.Fatalf("❌ Unknown mode '%s' (supported: seq, all, setup-source, recommend, migrate, verify, cleanup)", mode)
	}

	fmt.Println("\n🎉 AIMX Execution Completed Successfully!")
}

// runSequentialPipeline executes Stages 1 through 4 in order with configurable transition delays.
func runSequentialPipeline(beetle, tb *resty.Client, cfg TestConfig, auth AuthConfig, targets []TestCase, state *MigrationState) error {
	if err := setupSourceEnvironment(tb, cfg, auth, state); err != nil {
		return fmt.Errorf("stage 1 (setup-source) failed: %w", err)
	}

	for _, target := range targets {
		waitWithCountdown(cfg.Timing.StepTransitionSeconds, "Stage 1 -> Stage 2 transition (Source Ready -> Target GPU Recommendation)")
		if _, err := recommendTargetGpuInfra(beetle, tb, cfg, auth, target, state); err != nil {
			return fmt.Errorf("stage 2 (recommend) failed for %s/%s: %w", target.Csp, target.Region, err)
		}

		waitWithCountdown(cfg.Timing.StepTransitionSeconds, "Stage 2 -> Stage 3 transition (Recommendation -> Infra & transx Migration)")
		if err := migrateInfraAndAIService(beetle, tb, cfg, auth, target, state); err != nil {
			return fmt.Errorf("stage 3 (migrate) failed for %s/%s: %w", target.Csp, target.Region, err)
		}

		waitWithCountdown(cfg.Timing.StepTransitionSeconds, "Stage 3 -> Stage 4 transition (Migration -> End-to-End Verification)")
		if err := verifyMigratedAIService(tb, cfg, auth, target, state); err != nil {
			return fmt.Errorf("stage 4 (verify) failed for %s/%s: %w", target.Csp, target.Region, err)
		}
	}
	return nil
}

// checkBeetleReady verifies that the CM-Beetle REST server is reachable before issuing migration calls.
func checkBeetleReady(beetle *resty.Client, endpoint string) error {
	url := fmt.Sprintf("%s/beetle/readyz", strings.TrimRight(endpoint, "/"))
	resp, err := beetle.R().Get(url)
	if err != nil {
		return fmt.Errorf("cannot reach CM-Beetle at %s: %w", url, err)
	}
	if resp.IsError() {
		return fmt.Errorf("CM-Beetle readyz returned HTTP %d: %s", resp.StatusCode(), resp.String())
	}
	return nil
}
