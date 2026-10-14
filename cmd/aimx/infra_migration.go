package main

import (
	"fmt"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/go-resty/resty/v2"
)

// recommendTargetGpuInfra inspects the source GPU VM and requests target GPU infra candidates from CM-Beetle.
func recommendTargetGpuInfra(beetle, tb *resty.Client, cfg TestConfig, auth AuthConfig, target TestCase, state *MigrationState) (*cloudmodel.RecommendedInfra, error) {
	start := time.Now()
	fmt.Println("\n================================================================================")
	fmt.Printf("🧠 [Stage 2] Hybrid Source Discovery & Target GPU Recommendation (%s / %s)\n", target.Csp, target.Region)
	fmt.Println("================================================================================")

	srcInfra, err := discoverSourceInfra(tb, cfg, auth, state)
	if err != nil {
		return nil, err
	}

	reqBody := controller.RecommendInfraRequest{
		DesiredCspAndRegionPair: cloudmodel.CloudProperty{
			Csp:    target.Csp,
			Region: target.Region,
		},
		OnpremiseInfraModel: srcInfra,
	}

	url := fmt.Sprintf("%s/beetle/recommendation/infra", strings.TrimRight(cfg.Beetle.Endpoint, "/"))
	var apiResp model.ApiResponse[[]cloudmodel.RecommendedInfra]

	fmt.Printf("📡 Requesting GPU infrastructure recommendation from CM-Beetle (%s)...\n", url)
	resp, err := beetle.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("desiredCsp", target.Csp).
		SetQueryParam("desiredRegion", target.Region).
		SetQueryParam("limit", "5").
		SetBody(&reqBody).
		SetResult(&apiResp).
		Post(url)
	if err != nil {
		return nil, fmt.Errorf("recommendation HTTP error: %w", err)
	}
	if resp.IsError() || !apiResp.Success {
		return nil, fmt.Errorf("recommendation API failed (status %d): %s", resp.StatusCode(), resp.String())
	}

	candidates := apiResp.Data
	if len(candidates) == 0 {
		return nil, fmt.Errorf("no GPU infrastructure candidates returned for %s/%s", target.Csp, target.Region)
	}

	fmt.Printf("✅ Received %d GPU infrastructure candidate(s):\n", len(candidates))
	for i := range candidates {
		c := candidates[i]
		specSummary := findCandidateSpecSummary(c)
		imgSummary := "N/A"
		if len(c.TargetInfra.NodeGroups) > 0 {
			imgSummary = firstNonEmpty(c.TargetInfra.NodeGroups[0].CspImageName, c.TargetInfra.NodeGroups[0].ImageId, "N/A")
		}
		fmt.Printf("   [%d] %s | Spec: %s | Image: %s\n", i+1, c.Description, specSummary, imgSummary)
	}

	selected := candidates[0]
	if len(selected.TargetInfra.NodeGroups) == 0 || strings.TrimSpace(selected.TargetInfra.NodeGroups[0].ImageId) == "" {
		return nil, fmt.Errorf("recommended GPU candidate for %s/%s has no valid OS imageId", target.Csp, target.Region)
	}
	state.RecommendedCandidates = candidates
	state.SelectedRecommendation = &selected
	state.StageDurations["recommend"] = time.Since(start).Round(time.Second).String()
	if err := saveState(cfg.Beetle.StateFile, state); err != nil {
		return nil, err
	}

	return &selected, nil
}

// findCandidateSpecSummary extracts the SpecInfo matching the candidate's first NodeGroup SpecId.
func findCandidateSpecSummary(c cloudmodel.RecommendedInfra) string {
	targetSpecID := ""
	if len(c.TargetInfra.NodeGroups) > 0 {
		targetSpecID = c.TargetInfra.NodeGroups[0].SpecId
	}
	for _, s := range c.TargetSpecList {
		if s.Id == targetSpecID {
			return fmt.Sprintf("%s (%d vCPU, %.0f GiB, %dx %s %.0fGB, $%.4f/hr)",
				s.CspSpecName, s.VCPU, s.MemoryGiB, s.AcceleratorCount, s.AcceleratorModel, s.AcceleratorMemoryGB, s.CostPerHour)
		}
	}
	if len(c.TargetSpecList) > 0 {
		s := c.TargetSpecList[len(c.TargetSpecList)-1]
		return fmt.Sprintf("%s (%d vCPU, %.0f GiB, %dx %s %.0fGB, $%.4f/hr)",
			s.CspSpecName, s.VCPU, s.MemoryGiB, s.AcceleratorCount, s.AcceleratorModel, s.AcceleratorMemoryGB, s.CostPerHour)
	}
	if targetSpecID != "" {
		return targetSpecID
	}
	return "N/A"
}

// provisionTargetGpuInfra calls CM-Beetle's POST /beetle/migration/ns/{nsId}/infra to provision or reuse the Target GPU VM.
func provisionTargetGpuInfra(beetle, tb *resty.Client, cfg TestConfig, auth AuthConfig, target TestCase, state *MigrationState) (*cloudmodel.VmInfraInfo, error) {
	if state.Target.InfraID != "" && state.Target.Csp == target.Csp && state.Target.Region == target.Region {
		baseURL := strings.TrimRight(auth.TumblebugEndpoint, "/")
		var existing tbmodel.InfraInfo
		getResp, err := tb.R().SetResult(&existing).Get(fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s", baseURL, cfg.Beetle.NamespaceID, state.Target.InfraID))
		if err == nil && !getResp.IsError() && isInfraHealthyAndRunning(existing) {
			fmt.Printf("\n✅ [Stage 3-1] Reusing existing running Target GPU Infrastructure '%s' (node=%s, ip=%s, spec=%s)\n",
				state.Target.InfraID, state.Target.NodeID, state.Target.PublicIP, state.Target.SpecID)
			return state.TargetInfra, nil
		}
	}

	if state.SelectedRecommendation == nil {
		if _, err := recommendTargetGpuInfra(beetle, tb, cfg, auth, target, state); err != nil {
			return nil, err
		}
	}

	specLabel := findCandidateSpecSummary(*state.SelectedRecommendation)
	fmt.Printf("\n🚀 [Stage 3-1] Provisioning Target GPU Infrastructure via CM-Beetle (%s / %s) — %s...\n",
		target.Csp, target.Region, specLabel)

	reqBody := controller.MigrateInfraRequest{
		RecommendedInfra: *state.SelectedRecommendation,
	}

	url := fmt.Sprintf("%s/beetle/migration/ns/%s/infra", strings.TrimRight(cfg.Beetle.Endpoint, "/"), cfg.Beetle.NamespaceID)
	var apiResp model.ApiResponse[cloudmodel.VmInfraInfo]

	req := beetle.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&reqBody).
		SetResult(&apiResp)
	if target.NameSeed != "" {
		req.SetQueryParam("nameSeed", target.NameSeed)
	}

	resp, err := req.Post(url)
	if err != nil {
		return nil, fmt.Errorf("target GPU infra migration HTTP error: %w", err)
	}
	if resp.IsError() || !apiResp.Success {
		return nil, fmt.Errorf("target GPU infra migration failed (status %d): %s", resp.StatusCode(), resp.String())
	}

	vmInfra := apiResp.Data
	if len(vmInfra.Node) == 0 {
		return nil, fmt.Errorf("target GPU infra '%s' provisioned with 0 nodes", vmInfra.Id)
	}

	node := vmInfra.Node[0]
	state.TargetInfra = &vmInfra
	state.Target = EndpointMetrics{
		InfraID:     vmInfra.Id,
		NodeID:      node.Id,
		PublicIP:    node.PublicIP,
		SshKeyID:    node.SshKeyId,
		SshUsername: node.NodeUserName,
		Csp:         target.Csp,
		Region:      target.Region,
		SpecID:      node.SpecId,
		ImageID:     node.ImageId,
		ModelName:   cfg.AIService.ModelName,
		WebUIUrl:    fmt.Sprintf("http://%s:%d", node.PublicIP, cfg.AIService.OpenWebUIPort),
	}
	if err := saveState(cfg.Beetle.StateFile, state); err != nil {
		return nil, err
	}

	if len(node.SecurityGroupIds) > 0 {
		openSecurityGroupServicePorts(tb, auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, node.SecurityGroupIds, cfg.AIService)
	}

	fmt.Printf("✅ Target GPU VM Provisioned: infraId=%s, nodeId=%s, publicIp=%s, spec=%s\n",
		state.Target.InfraID, state.Target.NodeID, state.Target.PublicIP, state.Target.SpecID)

	waitWithCountdown(cfg.Timing.VmBootStabilizeSeconds, "Waiting for newly provisioned Target GPU VM boot & NVIDIA driver init")
	if _, err := pollGpuVmReady(tb, auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, state.Target.InfraID, state.Target.NodeID, cfg.Timing); err != nil {
		return nil, err
	}

	return &vmInfra, nil
}
