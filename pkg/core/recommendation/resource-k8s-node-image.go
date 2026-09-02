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

	profile := getTargetProfile(provider, region)

	for i, g := range groupWorkersByArch(workers) {
		members := machineIdsOf(g.nodes)

		imageId, err := selectK8sNodeImage(profile, g.arch)
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

		desc := buildNodeImageDescription(profile, i+1, len(g.nodes), g.arch, imageId)

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

// selectK8sNodeImage picks the node image for the given architecture from the target profile.
//
// Behavior:
//   - NodeImageDesignation=false (the CSP manages the image) or x86_64 → "default".
//   - arm64 + designation → pick an ARM node image from the curated list.
//   - arm64 but no ARM image available → error.
//
// Note: when the profile lookup failed, nodeImageDesignation is false and every architecture gets
// "default" — including arm64, which then pairs an ARM spec with the provider's x86 image. That
// is the pre-existing behavior, kept here deliberately; see
// docs/k8s-recommendation/k8s-io-consolidation-implementation-plan.md for the proposed fix.
func selectK8sNodeImage(p targetProfile, arch string) (imageId string, err error) {
	if !p.nodeImageDesignation || arch != "arm64" {
		return "default", nil
	}
	for _, img := range p.nodeImages {
		if isArmNodeImageId(img.Id) {
			return img.Id, nil
		}
	}
	return "", fmt.Errorf(
		"no arm64 K8s node image is available for provider %q region %q. Use x86_64 worker nodes, "+
			"or check the available node images with searchImage(isKubernetesImage=true, provider=%q)",
		p.provider, p.region, p.provider)
}

// isArmNodeImageId reports whether a curated node image identifier denotes an ARM/aarch64 image.
func isArmNodeImageId(id string) bool {
	l := strings.ToLower(id)
	return strings.Contains(l, "arm64") || strings.Contains(l, "arm_64") ||
		strings.Contains(l, "aarch64") || strings.Contains(l, "_arm_") || strings.Contains(l, "-arm-")
}

// buildNodeImageDescription constructs the per-entry description in user-friendly language.
func buildNodeImageDescription(p targetProfile, groupIdx, nodeCount int, arch, imageId string) string {
	if !p.nodeImageDesignation {
		return fmt.Sprintf(
			"Architecture group %d (%d node(s), %s) — %s manages node image selection automatically",
			groupIdx, nodeCount, arch, p.provider)
	}
	return fmt.Sprintf(
		"Architecture group %d (%d node(s), %s) — Use node image '%s' for %s %s",
		groupIdx, nodeCount, arch, imageId, p.provider, p.region)
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
