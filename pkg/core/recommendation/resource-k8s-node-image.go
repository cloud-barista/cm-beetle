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

// RecommendK8sNodeGroupImages recommends a node image for every architecture present among the
// source workers, and assembles the L0 response.
//
// Grouping is by normalized architecture, which is the only input to node image selection. This
// intentionally differs from the node group breakdown in RecommendK8sInfra: several node groups can
// share one architecture and therefore one node image, and grouping by spec here would cost a
// Tumblebug spec search per group for a result that does not depend on the spec.
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

	for i, g := range groupWorkersByArch(workers) {
		members := machineIdsOf(g.nodes)

		imageId, designation, err := selectK8sNodeImage(provider, region, g.arch)
		if err != nil {
			log.Warn().Err(err).Str("arch", g.arch).Msg("failed to resolve K8s node image")
			ret.RecommendedOsImageList = append(ret.RecommendedOsImageList, cloudmodel.RecommendedOsImage{
				Status:        string(NothingRecommended),
				SourceServers: members,
				Description:   fmt.Sprintf("failed to resolve node image for architecture group %d (%s): %v", i+1, g.arch, err),
				TargetOsImage: cloudmodel.ImageInfo{},
			})
			continue
		}

		desc := buildNodeImageDescription(provider, region, i+1, len(g.nodes), g.arch, imageId, designation)

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

// archGroup is the set of source workers sharing one normalized architecture.
type archGroup struct {
	arch  string
	nodes []onpremmodel.NodeProperty
}

// groupWorkersByArch groups workers by normalized architecture, preserving first-appearance order
// so group indices are deterministic.
func groupWorkersByArch(workers []onpremmodel.NodeProperty) []archGroup {
	index := make(map[string]int)
	var groups []archGroup
	for _, w := range workers {
		arch := normalizeArch(w.CPU.Architecture)
		if i, ok := index[arch]; ok {
			groups[i].nodes = append(groups[i].nodes, w)
			continue
		}
		index[arch] = len(groups)
		groups = append(groups, archGroup{arch: arch, nodes: []onpremmodel.NodeProperty{w}})
	}
	return groups
}

// selectK8sNodeImage picks the node image for the architecture and also returns the provider's
// image-designation flag, so a caller that needs the flag (to describe the choice) does not have
// to fetch the same data twice. Callers that only want the id discard it: `id, _, err := ...`.
//
// Behavior:
//   - NodeImageDesignation=false (CSP manages the image) or x86_64 arch → "default".
//   - NodeImageDesignation=true + arm64 → pick an ARM node image from the curated list.
//   - arm64 required but no ARM image available → error.
func selectK8sNodeImage(provider, region, arch string) (imageId string, designation bool, err error) {
	designation, images, err := tbclient.NewSession().GetK8sNodeImages(provider, region)
	if err != nil {
		log.Warn().Err(err).Str("provider", provider).Str("region", region).
			Msg("Failed to fetch node images; using default image")
		return "default", false, nil
	}
	if !designation || arch != "arm64" {
		return "default", designation, nil
	}
	for _, img := range images {
		if isArmNodeImageId(img.Id) {
			return img.Id, designation, nil
		}
	}
	return "", designation, fmt.Errorf(
		"no arm64 K8s node image is available for provider %q region %q. Use x86_64 worker nodes, "+
			"or check the available node images with searchImage(isKubernetesImage=true, provider=%q)",
		provider, region, provider)
}

// isArmNodeImageId reports whether a curated node image identifier denotes an ARM/aarch64 image.
func isArmNodeImageId(id string) bool {
	l := strings.ToLower(id)
	return strings.Contains(l, "arm64") || strings.Contains(l, "arm_64") ||
		strings.Contains(l, "aarch64") || strings.Contains(l, "_arm_") || strings.Contains(l, "-arm-")
}

// buildNodeImageDescription constructs the per-entry description in user-friendly language.
// designation comes from the caller's own lookup rather than a second fetch of the same data.
func buildNodeImageDescription(provider, region string, groupIdx, nodeCount int, arch, imageId string, designation bool) string {
	if !designation {
		return fmt.Sprintf(
			"Architecture group %d (%d node(s), %s) — %s manages node image selection automatically",
			groupIdx, nodeCount, arch, provider)
	}
	return fmt.Sprintf(
		"Architecture group %d (%d node(s), %s) — Use node image '%s' for %s %s",
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
