package main

import (
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

func TestGenerateRandomOnpremInfraAndConsolidation(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	genCfg := GeneratorConfig{
		MinVCPU:     1,
		MaxVCPU:     4,
		MinMemoryGB: 2,
		MaxMemoryGB: 8,
		MinDiskGB:   20,
		MaxDiskGB:   40,
	}

	infra := GenerateRandomOnpremInfra(rng, 20, genCfg)
	if len(infra.Nodes) != 20 {
		t.Fatalf("expected 20 nodes, got %d", len(infra.Nodes))
	}
	if len(infra.Network.IPv4Networks.DefaultGateways) != 20 {
		t.Fatalf("expected 20 default gateways, got %d", len(infra.Network.IPv4Networks.DefaultGateways))
	}

	vcpus, memGB, diskGB := SummarizeOnpremResources(infra)
	if vcpus < 20 || memGB < 40 || diskGB < 400 {
		t.Fatalf("unexpected resource summary: vcpus=%d, mem=%d, disk=%d", vcpus, memGB, diskGB)
	}

	// Verify NodeGroup consolidation preserves total NodeGroupSize sum.
	rec := cloudmodel.RecommendedInfra{
		TargetInfra: cloudmodel.InfraReq{
			Name: "infra101",
			NodeGroups: []cloudmodel.CreateNodeGroupReq{
				{Name: "ng-1", SpecId: "aws+ap-northeast-2+t3.medium", ImageId: "ubuntu22", RootDiskType: "gp3", RootDiskSize: 20, SubnetId: "sub-1", SecurityGroupIds: []string{"sg-01"}, NodeGroupSize: 1},
				{Name: "ng-2", SpecId: "aws+ap-northeast-2+t3.medium", ImageId: "ubuntu22", RootDiskType: "gp3", RootDiskSize: 20, SubnetId: "sub-1", SecurityGroupIds: []string{"sg-01"}, NodeGroupSize: 1},
				{Name: "ng-3", SpecId: "aws+ap-northeast-2+t3.large", ImageId: "ubuntu22", RootDiskType: "gp3", RootDiskSize: 30, SubnetId: "sub-2", SecurityGroupIds: []string{"sg-02"}, NodeGroupSize: 1},
			},
		},
	}

	merged := ConsolidateIdenticalNodeGroups(rec)
	if len(merged.TargetInfra.NodeGroups) != 2 {
		t.Fatalf("expected 2 consolidated NodeGroups, got %d", len(merged.TargetInfra.NodeGroups))
	}
	totalSize := 0
	for _, ng := range merged.TargetInfra.NodeGroups {
		totalSize += ng.NodeGroupSize
	}
	if totalSize != 3 {
		t.Fatalf("expected total NodeGroupSize=3, got %d", totalSize)
	}
}

func TestRunStressSuiteEqualRound(t *testing.T) {
	cfg, err := loadStressConfig("testconf/template-test-config.yaml")
	if err != nil {
		t.Fatalf("failed to load template config: %v", err)
	}
	cfg.Beetle.Endpoint = ""

	if err := cfg.Validate(); err != nil {
		t.Fatalf("expected valid config, got error: %v", err)
	}

	report := RunStressSuite(cfg, nil)
	if report.TotalRounds != 10 {
		t.Fatalf("expected 10 rounds, got %d", report.TotalRounds)
	}
	if report.OrphanResourcesLeft != 0 {
		t.Fatalf("expected 0 orphan resources, got %d", report.OrphanResourcesLeft)
	}

	// Verify NameSeed and InfraID uniqueness and naming/reference validity across all steps.
	seenSeeds := make(map[string]bool)
	seenInfraIDs := make(map[string]bool)
	for _, step := range report.Steps {
		if ok, detail := common.IsValidNameSeed(step.NameSeed); !ok {
			t.Fatalf("invalid NameSeed %q at step %d: %s", step.NameSeed, step.SequenceIndex, detail)
		}
		if seenSeeds[step.NameSeed] {
			t.Fatalf("duplicate NameSeed detected across steps: %q", step.NameSeed)
		}
		seenSeeds[step.NameSeed] = true

		if seenInfraIDs[step.InfraID] {
			t.Fatalf("duplicate InfraID detected across steps: %q", step.InfraID)
		}
		seenInfraIDs[step.InfraID] = true

		synth := buildSyntheticRecommendedInfra(CSPTestCase{CloudProperty: cloudmodel.CloudProperty{Csp: step.CSP, Region: step.Region}}, len(defaultRolePresets), step.NameSeed, true)
		if ok, detail := common.ValidateComposedNames(synth); !ok {
			t.Fatalf("ValidateComposedNames failed for step %d (%s): %s", step.SequenceIndex, step.NameSeed, detail)
		}
	}

	// Verify NHN (simMaxNodes=30) was eliminated at 40 Nodes (Round 4) and excluded from subsequent rounds.
	var nhnStatus *CSPSurvivalStatus
	for i := range report.CSPStatuses {
		if report.CSPStatuses[i].CSP == "nhn" {
			nhnStatus = &report.CSPStatuses[i]
			break
		}
	}
	if nhnStatus == nil {
		t.Fatalf("nhn status not found")
	}
	if nhnStatus.Status != "ELIMINATED" || nhnStatus.MaxSucceededNodes != 30 || nhnStatus.FailedAtNodes != 40 {
		t.Fatalf("unexpected nhn status: %+v", *nhnStatus)
	}

	// Verify Markdown and HTML generation.
	tmpDir := t.TempDir()
	mdPath := filepath.Join(tmpDir, "summary.md")
	htmlPath := filepath.Join(tmpDir, "dashboard.html")

	if err := GenerateMarkdownReport(report, mdPath); err != nil {
		t.Fatalf("GenerateMarkdownReport failed: %v", err)
	}
	if err := GenerateInteractiveDashboardHTML(report, htmlPath); err != nil {
		t.Fatalf("GenerateInteractiveDashboardHTML failed: %v", err)
	}

	htmlBytes, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("failed to read generated HTML: %v", err)
	}
	htmlStr := string(htmlBytes)
	if strings.Contains(htmlStr, "innerHTML") {
		t.Fatalf("generated HTML must not use innerHTML")
	}
	if !strings.Contains(htmlStr, "Multi-CSP Progressive Stress Test") {
		t.Fatalf("generated HTML missing title")
	}
}
