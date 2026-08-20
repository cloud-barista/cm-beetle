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

// Package tbclient provides client functions to interact with CB-Tumblebug API
package tbclient

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

// K8sClusterList wraps the list response for K8s clusters from Tumblebug.
type K8sClusterList struct {
	K8sClusters []tbmodel.K8sClusterInfo `json:"K8sClusterInfo"`
}

// GetAvailableK8sVersions returns the K8s versions supported by the CSP and region.
func (s *Session) GetAvailableK8sVersions(providerName, regionName string) ([]tbmodel.K8sClusterVersionDetailAvailable, error) {
	log.Debug().Str("provider", providerName).Str("region", regionName).Msg("Getting available K8s versions")

	var result []tbmodel.K8sClusterVersionDetailAvailable
	url := fmt.Sprintf("/availableK8sVersion?providerName=%s&regionName=%s", providerName, regionName)

	resp, err := s.SetResult(&result).Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get available K8s versions")
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Int("count", len(result)).Msg("Got available K8s versions")
	return result, nil
}

// K8sClusterRecommendNode finds appropriate K8s worker node specs by filtering and prioritizing.
//
// This is the K8s counterpart of InfraRecommendSpec. Tumblebug wraps the shared spec search with
// validateK8sMinimumRequirements, which rejects filter conditions below the K8s node minimums and
// injects those minimums when the request omits them.
//
// The request is passed as a struct rather than a JSON string: InfraRecommendSpec accepts a string
// only to json.Unmarshal it straight back into this same type. nsId is fixed to the system common
// namespace on the Tumblebug side, so it is not a parameter here.
func (s *Session) K8sClusterRecommendNode(req tbmodel.RecommendSpecReq) ([]tbmodel.SpecInfo, error) {
	log.Debug().Msg("K8s Cluster Recommend Node")

	var specInfoList = []tbmodel.SpecInfo{}
	var emptyRet = []tbmodel.SpecInfo{}

	resp, err := s.
		SetBody(req).
		SetResult(&specInfoList).
		Post("/k8sClusterRecommendNode")

	if err != nil {
		log.Error().Err(err).Msg("Failed to recommend K8s node specs")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Int("count", len(specInfoList)).Msg("Got K8s node spec recommendations")
	return specInfoList, nil
}

// getK8sClusterDetail fetches the per-CSP K8s cluster asset details (naming rule, required
// subnet count, etc.) from Tumblebug's k8sClusterInfo endpoint.
func (s *Session) getK8sClusterDetail(providerName string) (tbmodel.K8sClusterDetail, error) {
	emptyRet := tbmodel.K8sClusterDetail{}

	var result tbmodel.K8sClusterAssetInfo
	resp, err := s.SetResult(&result).Get("/k8sClusterInfo")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get K8s cluster info")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	detail, ok := result.CSPs[strings.ToLower(providerName)]
	if !ok {
		return emptyRet, fmt.Errorf("no K8s cluster info for provider %q", providerName)
	}
	return detail, nil
}

// GetK8sNodeGroupNamingRule returns the CSP-specific node group naming rule (a regex)
// from Tumblebug's k8sClusterInfo asset data. Returns an empty string when the CSP
// defines no rule. The regex does not always encode length limits (e.g. Azure omits
// its 12-char cap), so callers must apply their own safe upper bound in addition.
func (s *Session) GetK8sNodeGroupNamingRule(providerName string) (string, error) {
	log.Debug().Str("provider", providerName).Msg("Getting K8s node group naming rule")

	detail, err := s.getK8sClusterDetail(providerName)
	if err != nil {
		return "", err
	}

	log.Debug().Str("provider", providerName).Str("namingRule", detail.NodeGroupNamingRule).Msg("Got K8s node group naming rule")
	return detail.NodeGroupNamingRule, nil
}

// GetK8sNodeImages returns whether the CSP requires node image designation and the curated
// node image list resolved for the region (preferring an exact region match, falling back to
// the "common" region entry). Data comes from the same k8sClusterInfo asset used for naming
// rules and subnet counts, so no extra (still under-development) node-image endpoint is needed.
func (s *Session) GetK8sNodeImages(providerName, regionName string) (bool, []tbmodel.K8sClusterNodeImageDetailAvailable, error) {
	log.Debug().Str("provider", providerName).Str("region", regionName).Msg("Getting K8s node images")

	detail, err := s.getK8sClusterDetail(providerName)
	if err != nil {
		return false, nil, err
	}
	return detail.NodeImageDesignation, resolveRegionNodeImages(detail.NodeImage, regionName), nil
}

// resolveRegionNodeImages picks the node image list for the given region, preferring an exact
// region match and falling back to the "common" keyword entry (per k8sclusterinfo.yaml semantics).
func resolveRegionNodeImages(details []tbmodel.K8sClusterNodeImageDetail, regionName string) []tbmodel.K8sClusterNodeImageDetailAvailable {
	var common []tbmodel.K8sClusterNodeImageDetailAvailable
	for _, d := range details {
		for _, r := range d.Region {
			if strings.EqualFold(r, regionName) {
				return d.Available
			}
			if strings.EqualFold(r, "common") {
				common = d.Available
			}
		}
	}
	return common
}

// GetK8sRequiredSubnetCount returns the number of subnets the CSP requires to create a K8s
// cluster (e.g. AWS EKS needs 2 across distinct AZs; most CSPs need 1), from Tumblebug's
// k8sClusterInfo asset data. The value is a required minimum, not an exact/maximum.
func (s *Session) GetK8sRequiredSubnetCount(providerName string) (int, error) {
	log.Debug().Str("provider", providerName).Msg("Getting K8s required subnet count")

	detail, err := s.getK8sClusterDetail(providerName)
	if err != nil {
		return 0, err
	}

	log.Debug().Str("provider", providerName).Int("requiredSubnetCount", detail.RequiredSubnetCount).Msg("Got K8s required subnet count")
	return detail.RequiredSubnetCount, nil
}

// CreateK8sCluster creates a K8s cluster using pre-created VNet, SshKey, and SecurityGroup resources.
func (s *Session) CreateK8sCluster(nsId string, reqBody tbmodel.K8sClusterReq) (tbmodel.K8sClusterInfo, error) {
	log.Debug().Str("name", reqBody.Name).Msg("Creating K8s cluster")

	emptyRet := tbmodel.K8sClusterInfo{}

	// EKS creation can take 10-20 minutes; use a generous timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	url := fmt.Sprintf("/ns/%s/k8sCluster", nsId)
	resBody := tbmodel.K8sClusterInfo{}

	resp, err := s.
		SetContext(ctx).
		SetBody(&reqBody).
		SetResult(&resBody).
		Post(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create K8s cluster")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Str("clusterId", resBody.Id).Msg("Created K8s cluster successfully")
	return resBody, nil
}

// AddK8sNodeGroup adds a worker node group to an existing K8s cluster.
// Required for AWS (nodeGroupsOnCreation=false); universally compatible with other CSPs.
func (s *Session) AddK8sNodeGroup(nsId, clusterId string, reqBody tbmodel.K8sNodeGroupReq) (tbmodel.K8sClusterInfo, error) {
	log.Debug().Str("clusterId", clusterId).Str("nodeGroup", reqBody.Name).Msg("Adding K8s node group")

	emptyRet := tbmodel.K8sClusterInfo{}

	// Node group creation can also take several minutes.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	url := fmt.Sprintf("/ns/%s/k8sCluster/%s/k8sNodeGroup", nsId, clusterId)
	resBody := tbmodel.K8sClusterInfo{}

	resp, err := s.
		SetContext(ctx).
		SetBody(&reqBody).
		SetResult(&resBody).
		Post(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to add K8s node group")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Str("clusterId", resBody.Id).Msg("Added K8s node group successfully")
	return resBody, nil
}

// ReadK8sCluster retrieves information about a specific K8s cluster.
func (s *Session) ReadK8sCluster(nsId, clusterId string) (tbmodel.K8sClusterInfo, error) {
	log.Debug().Str("clusterId", clusterId).Msg("Reading K8s cluster")

	emptyRet := tbmodel.K8sClusterInfo{}
	url := fmt.Sprintf("/ns/%s/k8sCluster/%s", nsId, clusterId)
	resBody := tbmodel.K8sClusterInfo{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to read K8s cluster")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Str("clusterId", resBody.Id).Msg("Read K8s cluster successfully")
	return resBody, nil
}

// ReadAllK8sClusters retrieves all K8s clusters in the specified namespace.
func (s *Session) ReadAllK8sClusters(nsId string) (K8sClusterList, error) {
	log.Debug().Str("nsId", nsId).Msg("Reading all K8s clusters")

	emptyRet := K8sClusterList{}
	url := fmt.Sprintf("/ns/%s/k8sCluster", nsId)
	resBody := K8sClusterList{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to read all K8s clusters")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Int("count", len(resBody.K8sClusters)).Msg("Read all K8s clusters successfully")
	return resBody, nil
}

// ReadK8sClusterIds retrieves only the K8s cluster IDs in the namespace.
//
// Tumblebug's full cluster list refreshes every cluster's state through Spider, so its cost
// grows with the number of clusters and can exceed Tumblebug's own 120 s request timeout.
// The `option=id` variant reads IDs straight from Tumblebug's key-value store, which stays
// cheap no matter how many clusters exist.
func (s *Session) ReadK8sClusterIds(nsId string) (tbmodel.IdList, error) {
	log.Debug().Str("nsId", nsId).Msg("Reading K8s cluster IDs")

	emptyRet := tbmodel.IdList{}
	url := fmt.Sprintf("/ns/%s/k8sCluster?option=id", nsId)
	resBody := tbmodel.IdList{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to read K8s cluster IDs")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Int("count", len(resBody.IdList)).Msg("Read K8s cluster IDs successfully")
	return resBody, nil
}

// CheckK8sNodeGroupsOnCreation checks whether the given CSP requires NodeGroups
// to be included in the K8s cluster creation request (nodeGroupsOnCreation=true),
// or allows them to be added separately after cluster creation (nodeGroupsOnCreation=false).
func (s *Session) CheckK8sNodeGroupsOnCreation(providerName string) (bool, error) {
	log.Debug().Str("provider", providerName).Msg("Checking K8s nodeGroupsOnCreation")

	type result struct {
		Result string `json:"result"`
	}
	var res result
	url := fmt.Sprintf("/checkK8sNodeGroupsOnK8sCreation?providerName=%s", providerName)

	resp, err := s.SetResult(&res).Get(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check K8s nodeGroupsOnCreation")
		return false, err
	}
	if resp.IsError() {
		return false, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	required := res.Result == "true"
	log.Debug().Str("provider", providerName).Bool("nodeGroupsOnCreation", required).Msg("K8s nodeGroupsOnCreation checked")
	return required, nil
}

// DeleteK8sNodeGroup deletes a specific node group from a K8s cluster.
func (s *Session) DeleteK8sNodeGroup(nsId, clusterId, nodeGroupId string) error {
	log.Debug().Str("clusterId", clusterId).Str("nodeGroupId", nodeGroupId).Msg("Deleting K8s node group")

	url := fmt.Sprintf("/ns/%s/k8sCluster/%s/k8sNodeGroup/%s", nsId, clusterId, nodeGroupId)

	resp, err := s.Delete(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete K8s node group")
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Str("nodeGroupId", nodeGroupId).Msg("Deleted K8s node group successfully")
	return nil
}

// DeleteK8sCluster deletes a K8s cluster and all its node groups.
func (s *Session) DeleteK8sCluster(nsId, clusterId string) error {
	log.Debug().Str("clusterId", clusterId).Msg("Deleting K8s cluster")

	url := fmt.Sprintf("/ns/%s/k8sCluster/%s", nsId, clusterId)

	resp, err := s.Delete(url)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete K8s cluster")
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Str("clusterId", clusterId).Msg("Deleted K8s cluster successfully")
	return nil
}
