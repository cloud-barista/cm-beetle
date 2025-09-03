package recommendation

import (
	"fmt"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"

	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"

	// "github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/infra"
	// "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"

	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

// RecommendContainer recommends appropriate K8s node specs for container workloads
func RecommendContainer(provider, region string, srcInfra onpremmodel.OnpremInfra) (RecommendedInfraInfo, error) {
	var emptyResp RecommendedInfraInfo
	var result RecommendedInfraInfo

	result.TargetInfra = tbmodel.MciDynamicReq{
		Description: "Recommended Kubernetes node configuration by CM-Beetle",
		Name:        "recommended-k8s-cluster",
		SubGroups:   []tbmodel.CreateSubGroupDynamicReq{}, // a field to contain Kubernetes node information
	}

	client := resty.New()
	client.SetBaseURL(config.Tumblebug.RestUrl)
	client.SetBasicAuth(config.Tumblebug.API.Username, config.Tumblebug.API.Password)

	// Analyze container workload resource requirements from user input (srcInfra)
	var totalCores uint32 = 0
	var totalMemory uint32 = 0

	// Calculate resource requirements if container workloads exist
	if len(srcInfra.Servers) > 0 {
		// Utilize server information hosting container workloads
		for _, server := range srcInfra.Servers {
			totalCores += server.CPU.Cores
			totalMemory += uint32(server.Memory.TotalSize)
		}
	}

	var coresMin, coresMax uint32
	var memoryMin, memoryMax uint32

	if totalCores > 0 {
		coresMax = totalCores << 1 // Double
		if totalCores > 1 {
			coresMin = totalCores >> 1 // Half
		} else {
			coresMin = 1 // Minimum 1
		}
	} else {
		// Set default values
		coresMin = 2
		coresMax = 4
	}

	if totalMemory > 0 {
		memoryMax = totalMemory << 1 // Double
		if totalMemory > 1 {
			memoryMin = totalMemory >> 1 // Half
		} else {
			memoryMin = 1 // Minimum 1
		}
	} else {
		// Set default values
		memoryMin = 4
		memoryMax = 8
	}

	log.Debug().
		Uint32("coreLowerLimit", coresMin).
		Uint32("coreUpperLimit", coresMax).
		Uint32("memoryLowerLimit (GiB)", memoryMin).
		Uint32("memoryUpperLimit (GiB)", memoryMax).
		Str("providerName", provider).
		Str("regionName", region).
		Msg("Container workload resource requirements")

	// Step 1: Recommend spec for Kubernetes node - using dynamic values based on user input
	plan := tbmodel.RecommendSpecReq{
		Filter: tbmodel.FilterInfo{Policy: []tbmodel.FilterCondition{
			{Metric: "vCPU", Condition: []tbmodel.Operation{
				{Operator: ">=", Operand: fmt.Sprintf("%d", coresMin)},
				{Operator: "<=", Operand: fmt.Sprintf("%d", coresMax)},
			}},
			{Metric: "memoryGiB", Condition: []tbmodel.Operation{
				{Operator: ">=", Operand: fmt.Sprintf("%d", memoryMin)},
				{Operator: "<=", Operand: fmt.Sprintf("%d", memoryMax)},
			}},
			{Metric: "providerName", Condition: []tbmodel.Operation{{Operand: provider}}},
			{Metric: "regionName", Condition: []tbmodel.Operation{{Operand: region}}},
			{Metric: "infraType", Condition: []tbmodel.Operation{{Operand: "k8s"}}},
		}},
		Priority: tbmodel.PriorityInfo{
			Policy: []tbmodel.PriorityCondition{
				{Metric: "performance", Weight: "0.5"},
				{Metric: "cost", Weight: "0.5"},
			},
		},
		Limit: "5",
	}

	// TbSpecInfo is a response body that contains a list of Kubernetes node specs.
	var specResp []tbmodel.SpecInfo
	resp, err := client.R().
		SetBody(plan).
		SetResult(&specResp).
		Post("/k8sClusterRecommendNode")
	if err != nil {
		log.Error().Err(err).Msg("failed to call k8sClusterRecommendNode")
		return emptyResp, fmt.Errorf("failed to call k8sClusterRecommendNode: %w", err)
	}

	if resp.StatusCode() != 200 {
		log.Error().Int("status", resp.StatusCode()).Msg("k8sClusterRecommendNode returned non-200 status")
		return emptyResp, fmt.Errorf("k8sClusterRecommendNode returned non-200 status: %d", resp.StatusCode())
	}

	if len(specResp) == 0 {
		log.Warn().Msg("no recommended specs found")
		result.Status = string(NothingRecommended)
		result.Description = "Could not find appropriate K8s node specification."
		return result, nil
	}

	// Recommend available OS images for Kubernetes node spec
	for _, specInfo := range specResp {
		// Step 2: Request the available OS image for the spec
		reqCheck := tbmodel.K8sClusterConnectionConfigCandidatesReq{
			SpecIds: []string{specInfo.Id},
		}

		var checkResp tbmodel.CheckK8sClusterDynamicReqInfo
		resp, err = client.R().
			SetBody(reqCheck).
			SetResult(&checkResp).
			Post("/k8sClusterDynamicCheckRequest")

		if err != nil {
			log.Error().Err(err).Str("specId", specInfo.Id).Msg("failed to call k8sClusterDynamicCheckRequest")
			continue
		}

		if resp.StatusCode() != 200 {
			log.Error().Int("status", resp.StatusCode()).Str("specId", specInfo.Id).Msg("k8sClusterDynamicCheckRequest returned non-200 status")
			continue
		}

		if len(checkResp.ReqCheck) == 0 {
			log.Warn().Str("specId", specInfo.Id).Msg("no compatibility info found for spec")
			continue
		}

		for _, nodeInfo := range checkResp.ReqCheck {
			commonSpec := specInfo.Id

			if len(nodeInfo.Image) == 0 {
				log.Warn().Str("spec", commonSpec).Msg("no compatible images found for spec")
				continue
			}

			imageID := "default" // * NOTE: Set to "default" because Some CSP's Kubernetes Service uses default image.
			if nodeInfo.Image[0].Id != "default" {
				imageID = nodeInfo.Image[0].Id
			}

			// Set the recommended Kubernetes node spec and OS image to the response body
			subgroup := tbmodel.CreateSubGroupDynamicReq{
				Name:        fmt.Sprintf("k8snode-%s", strings.Split(commonSpec, "-")[len(strings.Split(commonSpec, "-"))-1]),
				SpecId:      commonSpec,
				ImageId:     imageID,
				Description: "Recommended K8s node",
			}
			result.TargetInfra.SubGroups = append(result.TargetInfra.SubGroups, subgroup)
		}
	}

	if len(result.TargetInfra.SubGroups) > 0 {
		result.Status = string(FullyRecommended)
		result.Description = "K8s node configuration recommended."
	} else {
		result.Status = string(NothingRecommended)
		result.Description = "Could not find appropriate K8s node configuration."
	}

	return result, nil
}
