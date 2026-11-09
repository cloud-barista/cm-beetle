package main

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

type cspSimProfile struct {
	maxNodes       int
	baseLatencySec float64
	perNodeSec     float64
	failureCode    string
	failureReason  string
	avgMatchRate   float64
}

var defaultSimProfiles = map[string]cspSimProfile{
	"aws": {
		maxNodes:       100,
		baseLatencySec: 28.0,
		perNodeSec:     1.35,
		failureCode:    FailureCodeQuotaExceeded,
		failureReason:  "Regional On-Demand Standard (A, C, D, H, I, M, R, T, Z) instance vCPU quota reached in ap-northeast-2",
		avgMatchRate:   96.4,
	},
	"gcp": {
		maxNodes:       90,
		baseLatencySec: 24.0,
		perNodeSec:     1.25,
		failureCode:    FailureCodeRateLimited,
		failureReason:  "Compute Engine regional instance insert rate quota exceeded at 100 concurrent Nodes in asia-northeast3",
		avgMatchRate:   95.1,
	},
	"azure": {
		maxNodes:       80,
		baseLatencySec: 42.0,
		perNodeSec:     1.85,
		failureCode:    FailureCodeQuotaExceeded,
		failureReason:  "Total Regional vCPUs quota limit (240 vCPUs) exceeded for Standard D-series in koreasouth",
		avgMatchRate:   94.2,
	},
	"alibaba": {
		maxNodes:       70,
		baseLatencySec: 36.0,
		perNodeSec:     1.65,
		failureCode:    FailureCodeTimeout,
		failureReason:  "ECS batch provisioning and cloud_essd disk attachment exceeded stabilization window at 80 Nodes",
		avgMatchRate:   92.8,
	},
	"tencent": {
		maxNodes:       60,
		baseLatencySec: 34.0,
		perNodeSec:     1.60,
		failureCode:    FailureCodeCapacityExhausted,
		failureReason:  "Insufficient CVM S5/SA2 instance inventory in ap-seoul availability zone for 70 concurrent Nodes",
		avgMatchRate:   91.5,
	},
	"ncp": {
		maxNodes:       50,
		baseLatencySec: 46.0,
		perNodeSec:     2.10,
		failureCode:    FailureCodeNetworkExhausted,
		failureReason:  "VPC concurrent server creation and ACG attachment limit (50 Nodes per batch) reached in KR region",
		avgMatchRate:   90.3,
	},
	"ibm": {
		maxNodes:       40,
		baseLatencySec: 48.0,
		perNodeSec:     2.25,
		failureCode:    FailureCodeRateLimited,
		failureReason:  "VPC Virtual Server API rate limit (HTTP 429 Too Many Requests) triggered at 50 concurrent Nodes in au-syd",
		avgMatchRate:   89.4,
	},
	"nhn": {
		maxNodes:       30,
		baseLatencySec: 38.0,
		perNodeSec:     1.95,
		failureCode:    FailureCodeQuotaExceeded,
		failureReason:  "Project hypervisor instance and vCPU quota limit reached at 40 Nodes in kr1 (Pangyo)",
		avgMatchRate:   88.7,
	},
}

var matchRateAvgRegex = regexp.MustCompile(`Avg=([0-9.]+)%`)

// ExecuteStressStep runs Plan, Recommend, Migrate, and Cleanup for a single CSP at targetNodes scale.
func ExecuteStressStep(
	rng *rand.Rand,
	client *resty.Client,
	cfg StressTestConfig,
	beetleAvailable bool,
	tc CSPTestCase,
	round int,
	seqIndex int,
	targetNodes int,
) StepAttempt {
	stepStart := time.Now()
	displayName := tc.Name
	if displayName == "" {
		displayName = fmt.Sprintf("%s-%s", strings.ToUpper(tc.Csp), tc.Region)
	}

	nameSeed := buildStepNameSeed(cfg.Stress.NameSeedPrefix, round, tc.Csp)

	attempt := StepAttempt{
		SequenceIndex: seqIndex,
		Round:         round,
		CSP:           strings.ToLower(tc.Csp),
		Region:        tc.Region,
		DisplayName:   displayName,
		TargetNodes:   targetNodes,
		NameSeed:      nameSeed,
		StartTime:     stepStart,
	}

	// Phase 1: Plan (Generate random onpremiseInfraModel).
	planStart := time.Now()
	onpremInfra := GenerateRandomOnpremInfra(rng, targetNodes, cfg.Stress.Generator)
	vcpus, memGB, diskGB := SummarizeOnpremResources(onpremInfra)
	planEnd := time.Now()

	attempt.TotalVCPUs = vcpus
	attempt.TotalMemoryGB = memGB
	attempt.TotalDiskGB = diskGB
	attempt.PlanPhase = StepPhaseResult{
		Phase:        "Plan",
		StartTime:    planStart,
		EndTime:      planEnd,
		DurationMs:   max(1, planEnd.Sub(planStart).Milliseconds()),
		SimulatedSec: 0.8 + float64(targetNodes)*0.02,
		Success:      true,
		StatusCode:   200,
		Details:      fmt.Sprintf("Generated onpremiseInfraModel (%d Nodes, %d vCPUs, %d GiB Memory)", targetNodes, vcpus, memGB),
	}

	// Phase 2: Recommend & Validate.
	recPhase, seededRec, avgMatch := executeRecommendPhase(rng, client, cfg, beetleAvailable, tc, onpremInfra, nameSeed)
	attempt.RecommendPhase = recPhase
	attempt.AvgMatchRate = avgMatch
	attempt.GeneratedNodeGroups = len(seededRec.TargetInfra.NodeGroups)
	attempt.SecurityGroupCount = len(seededRec.TargetSecurityGroupList)
	attempt.InfraID = seededRec.TargetInfra.Name

	if !recPhase.Success {
		attempt.Success = false
		attempt.Eliminated = true
		attempt.FailurePhase = "Recommend"
		attempt.FailureCode, attempt.FailureReason = ClassifyFailureReason(recPhase.StatusCode, recPhase.ErrorMessage)
		attempt.MigratePhase = StepPhaseResult{
			Phase:        "Migrate",
			StartTime:    time.Now(),
			EndTime:      time.Now(),
			Skipped:      true,
			ErrorMessage: "Skipped due to recommendation failure",
		}
		attempt.CleanupPhase = StepPhaseResult{
			Phase:        "Cleanup",
			StartTime:    time.Now(),
			EndTime:      time.Now(),
			Success:      true,
			Skipped:      true,
			Details:      "No cloud resources were created prior to recommendation failure",
		}
		attempt.CleanupSummary = CleanupSummary{
			Attempted:     false,
			CleanVerified: true,
			Logs:          []string{"No cloud resources allocated; clean state verified"},
		}
		finalizeStepTiming(&attempt)
		return attempt
	}

	// Phase 3 & 4: Migrate and 2-Stage Cleanup (Real or Simulated).
	if cfg.Stress.ExecutionMode == ExecutionModeReal {
		runRealMigrationAndCleanup(client, cfg, tc, &seededRec, &attempt)
	} else {
		runSimulatedMigrationAndCleanup(rng, tc, &seededRec, &attempt)
	}

	finalizeStepTiming(&attempt)
	return attempt
}

func executeRecommendPhase(
	rng *rand.Rand,
	client *resty.Client,
	cfg StressTestConfig,
	beetleAvailable bool,
	tc CSPTestCase,
	onpremInfra interface{},
	nameSeed string,
) (StepPhaseResult, cloudmodel.RecommendedInfra, float64) {
	recStart := time.Now()
	profile := resolveSimProfile(tc)

	// Call live CM-Beetle recommendation endpoint when available.
	if beetleAvailable && client != nil && cfg.Beetle.Endpoint != "" {
		url := fmt.Sprintf("%s/beetle/recommendation/infra", strings.TrimRight(cfg.Beetle.Endpoint, "/"))
		reqBody := map[string]interface{}{
			"desiredCspAndRegionPair": cloudmodel.CloudProperty{
				Csp:    tc.Csp,
				Region: tc.Region,
			},
			"onpremiseInfraModel": onpremInfra,
		}

		var apiResp model.ApiResponse[[]cloudmodel.RecommendedInfra]
		err := common.ExecuteHttpRequest(client, "POST", url, nil, true, &reqBody, &apiResp, 0)
		recEnd := time.Now()
		durMs := max(1, recEnd.Sub(recStart).Milliseconds())

		if err == nil && apiResp.Success && len(apiResp.Data) > 0 && apiResp.Data[0].Status != "nothing-to-recommend" {
			candidate := apiResp.Data[0]
			if cfg.Stress.MergeIdenticalNodeGroups {
				candidate = ConsolidateIdenticalNodeGroups(candidate)
			}
			seeded := common.ApplyNameSeed(candidate, nameSeed)
			if ok, detail := common.ValidateComposedNames(seeded); !ok {
				return StepPhaseResult{
					Phase:        "Recommend",
					StartTime:    recStart,
					EndTime:      recEnd,
					DurationMs:   durMs,
					SimulatedSec: float64(durMs) / 1000.0,
					Success:      false,
					StatusCode:   400,
					ErrorMessage: "Naming/Reference validation failed: " + detail,
				}, seeded, 0
			}
			matchRate := parseAvgMatchRate(candidate.Description, profile.avgMatchRate)

			return StepPhaseResult{
				Phase:        "Recommend",
				StartTime:    recStart,
				EndTime:      recEnd,
				DurationMs:   durMs,
				SimulatedSec: float64(durMs) / 1000.0,
				Success:      true,
				StatusCode:   200,
				Details:      fmt.Sprintf("Recommended %d NodeGroups, %d SecurityGroups (Avg Match: %.1f%%)", len(seeded.TargetInfra.NodeGroups), len(seeded.TargetSecurityGroupList), matchRate),
			}, seeded, matchRate
		}

		// In real mode, a recommendation failure is a hard step failure.
		if cfg.Stress.ExecutionMode == ExecutionModeReal {
			errMsg := "Recommendation returned no compatible candidates"
			if err != nil {
				errMsg = err.Error()
			} else if apiResp.Error != "" {
				errMsg = apiResp.Error
			}
			return StepPhaseResult{
				Phase:        "Recommend",
				StartTime:    recStart,
				EndTime:      recEnd,
				DurationMs:   durMs,
				SimulatedSec: float64(durMs) / 1000.0,
				Success:      false,
				StatusCode:   500,
				ErrorMessage: errMsg,
			}, cloudmodel.RecommendedInfra{}, 0
		}
	}

	// Build deterministic recommended model when running offline simulation.
	synth := buildSyntheticRecommendedInfra(tc, len(defaultRolePresets), nameSeed, cfg.Stress.MergeIdenticalNodeGroups)
	recEnd := time.Now()
	simRecSec := 1.6 + float64(len(synth.TargetInfra.NodeGroups))*0.12 + rng.Float64()*0.5
	jitterMatch := math.Min(99.5, math.Max(80.0, profile.avgMatchRate+(rng.Float64()-0.5)*1.4))

	return StepPhaseResult{
		Phase:        "Recommend",
		StartTime:    recStart,
		EndTime:      recEnd,
		DurationMs:   max(1, recEnd.Sub(recStart).Milliseconds()),
		SimulatedSec: math.Round(simRecSec*10) / 10,
		Success:      true,
		StatusCode:   200,
		Details:      fmt.Sprintf("Recommended %d NodeGroups, %d deduplicated SecurityGroups (Avg Match: %.1f%%)", len(synth.TargetInfra.NodeGroups), len(synth.TargetSecurityGroupList), jitterMatch),
	}, synth, math.Round(jitterMatch*10) / 10
}

func runRealMigrationAndCleanup(
	client *resty.Client,
	cfg StressTestConfig,
	tc CSPTestCase,
	seededRec *cloudmodel.RecommendedInfra,
	attempt *StepAttempt,
) {
	migStart := time.Now()
	endpoint := strings.TrimRight(cfg.Beetle.Endpoint, "/")
	nsID := cfg.Beetle.NamespaceID

	// Validate the seeded recommendation model before migration.
	valURL := fmt.Sprintf("%s/beetle/validation/ns/%s/infra", endpoint, nsID)
	valReq := controller.ValidateInfraRequest{RecommendedInfra: *seededRec}
	var valResp map[string]interface{}
	_ = common.ExecuteHttpRequest(client, "POST", valURL, nil, true, &valReq, &valResp, 0)

	// Invoke migration endpoint with fresh resource creation enforced.
	migURL := fmt.Sprintf("%s/beetle/migration/ns/%s/infra?useExisting=false", endpoint, nsID)
	migReq := controller.MigrateInfraRequest{RecommendedInfra: *seededRec}
	var migResp model.ApiResponse[controller.MigrateInfraResponse]

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&migReq).
		SetResult(&migResp).
		Post(migURL)

	migEnd := time.Now()
	migDurMs := max(1, migEnd.Sub(migStart).Milliseconds())
	statusCode := 0
	if resp != nil {
		statusCode = resp.StatusCode()
	}

	mciStatusFailed := strings.Contains(strings.ToLower(migResp.Data.Status), "failed")
	if err == nil && resp != nil && resp.IsSuccess() && migResp.Success && !mciStatusFailed {
		if migResp.Data.Id != "" {
			attempt.InfraID = migResp.Data.Id
		} else if migResp.Data.Name != "" {
			attempt.InfraID = migResp.Data.Name
		}
		attempt.Success = true
		attempt.MigratePhase = StepPhaseResult{
			Phase:        "Migrate",
			StartTime:    migStart,
			EndTime:      migEnd,
			DurationMs:   migDurMs,
			SimulatedSec: float64(migDurMs) / 1000.0,
			Success:      true,
			StatusCode:   statusCode,
			Details:      fmt.Sprintf("Provisioned %d Nodes in MCI %s", attempt.TargetNodes, attempt.InfraID),
		}
	} else {
		rawErr := ""
		if err != nil {
			rawErr = err.Error()
		} else if mciStatusFailed {
			rawErr = fmt.Sprintf("MCI status reported %s", migResp.Data.Status)
		} else if migResp.Error != "" {
			rawErr = migResp.Error
		} else if resp != nil {
			rawErr = resp.String()
		}
		code, reason := ClassifyFailureReason(statusCode, rawErr)
		attempt.Success = false
		attempt.Eliminated = true
		attempt.FailurePhase = "Migrate"
		attempt.FailureCode = code
		attempt.FailureReason = reason
		attempt.MigratePhase = StepPhaseResult{
			Phase:        "Migrate",
			StartTime:    migStart,
			EndTime:      migEnd,
			DurationMs:   migDurMs,
			SimulatedSec: float64(migDurMs) / 1000.0,
			Success:      false,
			StatusCode:   statusCode,
			ErrorMessage: reason,
			Details:      rawErr,
		}
	}

	// Always run 2-stage cleanup whether migration succeeded or failed midway.
	cleanPhase, cleanSummary := ExecuteTwoStageCleanup(client, endpoint, nsID, attempt.InfraID, seededRec)
	attempt.CleanupPhase = cleanPhase
	attempt.CleanupSummary = cleanSummary
}

func runSimulatedMigrationAndCleanup(
	rng *rand.Rand,
	tc CSPTestCase,
	seededRec *cloudmodel.RecommendedInfra,
	attempt *StepAttempt,
) {
	profile := resolveSimProfile(tc)
	migStart := time.Now()

	simProvisionSec := profile.baseLatencySec + float64(attempt.TargetNodes)*profile.perNodeSec + rng.Float64()*4.0
	simProvisionSec = math.Round(simProvisionSec*10) / 10

	sgCount := max(1, len(seededRec.TargetSecurityGroupList))
	simCleanupSec := math.Round((12.0+float64(attempt.TargetNodes)*0.45+float64(sgCount)*1.8+rng.Float64()*2.5)*10) / 10

	if attempt.TargetNodes <= profile.maxNodes {
		attempt.Success = true
		attempt.Eliminated = false
		attempt.MigratePhase = StepPhaseResult{
			Phase:        "Migrate",
			StartTime:    migStart,
			EndTime:      time.Now(),
			DurationMs:   15,
			SimulatedSec: simProvisionSec,
			Success:      true,
			StatusCode:   200,
			Details:      fmt.Sprintf("Provisioned %d Nodes across %d NodeGroups in %s (%.1fs)", attempt.TargetNodes, attempt.GeneratedNodeGroups, attempt.InfraID, simProvisionSec),
		}
		attempt.CleanupSummary = CleanupSummary{
			Attempted:                true,
			InfraDeleted:             true,
			SecurityGroupsCleaned:    sgCount,
			SshKeysCleaned:           1,
			VNetsCleaned:             1,
			OrphanResourcesRemaining: 0,
			CleanVerified:            true,
			DurationMs:               10,
			SimulatedSec:             simCleanupSec,
			Logs: []string{
				fmt.Sprintf("Terminated MCI %s (%d Nodes)", attempt.InfraID, attempt.TargetNodes),
				fmt.Sprintf("Swept %d SecurityGroups, 1 SSHKey, 1 VNet (withsubnets)", sgCount),
			},
		}
	} else {
		attempt.Success = false
		attempt.Eliminated = true
		attempt.FailurePhase = "Migrate"
		attempt.FailureCode = profile.failureCode
		attempt.FailureReason = profile.failureReason

		statusCode := 500
		if profile.failureCode == FailureCodeRateLimited {
			statusCode = 429
		} else if profile.failureCode == FailureCodeQuotaExceeded {
			statusCode = 403
		}

		attempt.MigratePhase = StepPhaseResult{
			Phase:        "Migrate",
			StartTime:    migStart,
			EndTime:      time.Now(),
			DurationMs:   15,
			SimulatedSec: math.Round(simProvisionSec*0.65*10) / 10,
			Success:      false,
			StatusCode:   statusCode,
			ErrorMessage: profile.failureReason,
			Details:      fmt.Sprintf("[%s] %s", profile.failureCode, profile.failureReason),
		}
		attempt.CleanupSummary = CleanupSummary{
			Attempted:                true,
			InfraDeleted:             true,
			SecurityGroupsCleaned:    sgCount,
			SshKeysCleaned:           1,
			VNetsCleaned:             1,
			OrphanResourcesRemaining: 0,
			CleanVerified:            true,
			DurationMs:               12,
			SimulatedSec:             simCleanupSec,
			Logs: []string{
				fmt.Sprintf("Rollback triggered on %s failure at %d Nodes", profile.failureCode, attempt.TargetNodes),
				fmt.Sprintf("Force-terminated partial MCI %s", attempt.InfraID),
				fmt.Sprintf("2-stage sweep verified: %d SecurityGroups, 1 SSHKey, 1 VNet removed (0 orphans)", sgCount),
			},
		}
	}

	attempt.CleanupPhase = StepPhaseResult{
		Phase:        "Cleanup",
		StartTime:    time.Now(),
		EndTime:      time.Now(),
		DurationMs:   attempt.CleanupSummary.DurationMs,
		SimulatedSec: attempt.CleanupSummary.SimulatedSec,
		Success:      true,
		StatusCode:   200,
		Details:      strings.Join(attempt.CleanupSummary.Logs, " | "),
	}
}

func resolveSimProfile(tc CSPTestCase) cspSimProfile {
	cspKey := strings.ToLower(strings.TrimSpace(tc.Csp))
	prof, ok := defaultSimProfiles[cspKey]
	if !ok {
		prof = cspSimProfile{
			maxNodes:       50,
			baseLatencySec: 35.0,
			perNodeSec:     1.75,
			failureCode:    FailureCodeQuotaExceeded,
			failureReason:  "Regional vCPU quota limit reached",
			avgMatchRate:   90.0,
		}
	}

	if tc.SimMaxNodes > 0 {
		prof.maxNodes = tc.SimMaxNodes
	}
	if tc.SimFailureCode != "" {
		prof.failureCode = tc.SimFailureCode
	}
	if tc.SimFailureReason != "" {
		prof.failureReason = tc.SimFailureReason
	}
	if tc.SimBaseLatencySec > 0 {
		prof.baseLatencySec = float64(tc.SimBaseLatencySec)
	}
	return prof
}

func buildSyntheticRecommendedInfra(tc CSPTestCase, roleCount int, nameSeed string, mergeGroups bool) cloudmodel.RecommendedInfra {
	connName := fmt.Sprintf("%s-%s", strings.ToLower(tc.Csp), strings.ToLower(tc.Region))
	sgs := make([]cloudmodel.SecurityGroupReq, 0, roleCount)
	for i := 0; i < roleCount; i++ {
		sgs = append(sgs, cloudmodel.SecurityGroupReq{
			Name:           fmt.Sprintf("sg-%02d", i+1),
			ConnectionName: connName,
			VNetId:         "vnet-01",
		})
	}

	ngs := make([]cloudmodel.CreateNodeGroupReq, 0, roleCount)
	for i := 0; i < roleCount; i++ {
		ngs = append(ngs, cloudmodel.CreateNodeGroupReq{
			Name:             fmt.Sprintf("ng-%02d", i+1),
			ConnectionName:   connName,
			VNetId:           "vnet-01",
			SubnetId:         "subnet-01",
			SshKeyId:         "sshkey-01",
			SecurityGroupIds: []string{fmt.Sprintf("sg-%02d", i+1)},
			NodeGroupSize:    1,
		})
	}

	rec := cloudmodel.RecommendedInfra{
		Status: "fully-recommended",
		TargetCloud: cloudmodel.CloudProperty{
			Csp:    strings.ToLower(tc.Csp),
			Region: tc.Region,
		},
		TargetVNet: cloudmodel.VNetReq{
			Name:           "vnet-01",
			ConnectionName: connName,
			SubnetInfoList: []cloudmodel.SubnetReq{
				{Name: "subnet-01", IPv4_CIDR: "10.0.1.0/24"},
				{Name: "subnet-02", IPv4_CIDR: "10.0.2.0/24"},
			},
		},
		TargetSshKey: cloudmodel.SshKeyReq{
			Name:           "sshkey-01",
			ConnectionName: connName,
		},
		TargetSecurityGroupList: sgs,
		TargetInfra: cloudmodel.InfraReq{
			Name:       "infra101",
			NodeGroups: ngs,
		},
	}

	if mergeGroups {
		rec = ConsolidateIdenticalNodeGroups(rec)
	}
	return common.ApplyNameSeed(rec, nameSeed)
}

func buildStepNameSeed(prefix string, round int, cspName string) string {
	shortCSP := strings.ToLower(strings.TrimSpace(cspName))
	if len(shortCSP) > 5 {
		shortCSP = shortCSP[:5]
	}
	seed := fmt.Sprintf("%s-r%02d-%s", prefix, round, shortCSP)
	if len(seed) > 20 {
		seed = seed[:20]
	}
	return strings.TrimRight(seed, "-")
}

func parseAvgMatchRate(desc string, fallback float64) float64 {
	matches := matchRateAvgRegex.FindStringSubmatch(desc)
	if len(matches) == 2 {
		if val, err := strconv.ParseFloat(matches[1], 64); err == nil {
			return val
		}
	}
	return fallback
}

func finalizeStepTiming(attempt *StepAttempt) {
	attempt.EndTime = time.Now()
	attempt.DurationMs = max(1, attempt.EndTime.Sub(attempt.StartTime).Milliseconds())
	totalSim := attempt.PlanPhase.SimulatedSec +
		attempt.RecommendPhase.SimulatedSec +
		attempt.MigratePhase.SimulatedSec +
		attempt.CleanupPhase.SimulatedSec
	attempt.SimulatedDurationS = math.Round(totalSim*10) / 10
}
