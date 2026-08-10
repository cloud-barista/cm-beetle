/*
Copyright 2024 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package recommendation

import (
	"fmt"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"
)

// RecommendK8sNodeGroupImages recommends a node image for every worker node group derived from the
// source workers, and assembles the L0 response.
//
// Workers are grouped by spec signature (same rule as RecommendK8sNodeGroupSpecs).
// Each group receives the architecture-appropriate node image from the provider's curated list.
//
// Returns RecommendedOsImageList. For K8s node images, only TargetOsImage.Id is populated;
// other ImageInfo fields are empty because the curated list carries only an id string.
//
// If NodeImageDesignation=false for the provider, "default" is returned for every group.
// If arm64 designation is required but no ARM image is available in the curated list, an
// error is returned so the limitation surfaces at recommendation time.
func RecommendK8sNodeGroupImages(provider, region string, srcInfra onpremmodel.OnpremInfra) (cloudmodel.RecommendedOsImageList, error) {

	ret := cloudmodel.RecommendedOsImageList{}

	workers := collectWorkerNodes(srcInfra.Nodes)
	if len(workers) == 0 {
		return ret, fmt.Errorf("no worker nodes found in source infra (nodes with role=worker are required)")
	}

	for i, g := range groupWorkersBySpec(workers) {
		members := machineIdsOf(g.nodes)
		arch := normalizeArch(g.nodes[0].CPU.Architecture)

		imageId, err := ResolveK8sNodeImageId(provider, region, arch)
		if err != nil {
			log.Warn().Err(err).Int("nodeGroupIndex", i+1).Msg("failed to resolve K8s node image")
			ret.RecommendedOsImageList = append(ret.RecommendedOsImageList, cloudmodel.RecommendedOsImage{
				Status:        string(NothingRecommended),
				SourceServers: members,
				Description:   fmt.Sprintf("failed to resolve node image for node group %d: %v", i+1, err),
				TargetOsImage: cloudmodel.ImageInfo{},
			})
			continue
		}

		desc := buildNodeImageDescription(provider, region, i+1, len(g.nodes), arch, imageId)

		ret.RecommendedOsImageList = append(ret.RecommendedOsImageList, cloudmodel.RecommendedOsImage{
			Status:        string(FullyRecommended),
			SourceServers: members,
			Description:   desc,
			TargetOsImage: cloudmodel.ImageInfo{Id: imageId},
		})
	}

	ret.Count = len(ret.RecommendedOsImageList)
	ret.Status, ret.Description = summarizeRecommendedOsImageList(ret.RecommendedOsImageList)

	log.Info().
		Str("provider", provider).Str("region", region).
		Int("workerCount", len(workers)).Int("imageCount", ret.Count).
		Msg("K8s node group image recommendation completed")

	return ret, nil
}

// ResolveK8sNodeImageId returns the node image ID appropriate for the given architecture.
// Promoted from the unexported resolveNodeImageId used internally by RecommendK8sCluster.
//
// Behavior:
//   - NodeImageDesignation=false (CSP manages the image) or x86_64 arch → "default".
//   - NodeImageDesignation=true + arm64 → pick an ARM node image from the curated list.
//   - arm64 required but no ARM image available → error.
func ResolveK8sNodeImageId(provider, region, arch string) (string, error) {
	designation, images, err := tbclient.NewSession().GetK8sNodeImages(provider, region)
	if err != nil {
		log.Warn().Err(err).Str("provider", provider).Str("region", region).
			Msg("Failed to fetch node images; using default image")
		return "default", nil
	}
	if !designation || arch != "arm64" {
		return "default", nil
	}
	for _, img := range images {
		if isArmNodeImageId(img.Id) {
			return img.Id, nil
		}
	}
	return "", fmt.Errorf("no arm64 node image available for provider %q region %q (node image "+
		"designation required); add an ARM node image to k8sclusterinfo.yaml or use an x86_64 spec", provider, region)
}

// isArmNodeImageId reports whether a curated node image identifier denotes an ARM/aarch64 image.
func isArmNodeImageId(id string) bool {
	l := strings.ToLower(id)
	return strings.Contains(l, "arm64") || strings.Contains(l, "arm_64") ||
		strings.Contains(l, "aarch64") || strings.Contains(l, "_arm_") || strings.Contains(l, "-arm-")
}

// buildNodeImageDescription constructs the per-entry description in user-friendly language.
func buildNodeImageDescription(provider, region string, groupIdx, nodeCount int, arch, imageId string) string {
	designation, _, err := tbclient.NewSession().GetK8sNodeImages(provider, region)
	if err != nil || !designation {
		return fmt.Sprintf(
			"Node group %d (%d node(s), %s) — %s manages node image selection automatically",
			groupIdx, nodeCount, arch, provider)
	}
	if arch != "arm64" {
		return fmt.Sprintf(
			"Node group %d (%d node(s), %s) — Use node image '%s' for %s %s",
			groupIdx, nodeCount, arch, imageId, provider, region)
	}
	return fmt.Sprintf(
		"Node group %d (%d node(s), %s) — Use %s for %s %s",
		groupIdx, nodeCount, arch, imageId, provider, region)
}

// summarizeRecommendedOsImageList derives the overall status and description.
func summarizeRecommendedOsImageList(list []cloudmodel.RecommendedOsImage) (status, description string) {
	countFailed := 0
	for _, entry := range list {
		if entry.Status == string(NothingRecommended) {
			countFailed++
		}
	}
	switch {
	case len(list) == 0 || countFailed == len(list):
		return string(NothingRecommended), "No K8s node images available"
	case countFailed == 0:
		return string(FullyRecommended), fmt.Sprintf("Recommended node image(s) for %d node group(s)", len(list))
	default:
		return string(PartiallyRecommended), fmt.Sprintf(
			"Recommended %d of %d node image(s)", len(list)-countFailed, len(list))
	}
}
