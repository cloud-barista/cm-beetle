package recommendation

import (
	"fmt"
	"strings"
	"time"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"
)

// Honeybee v0.4.0 Kubernetes models
type KubernetesInfoList struct {
	Servers []Kubernetes `json:"servers" validate:"required"`
}

type Kubernetes struct {
	NodeCount NodeCount              `json:"node_count"`
	Nodes     []Node                 `json:"nodes"`
	Workloads map[string]interface{} `json:"workloads"`
}

type NodeCount struct {
	Total        int `json:"total"`
	ControlPlane int `json:"control_plane"`
	Worker       int `json:"worker"`
}

type NodeType string

const (
	NodeTypeControlPlane NodeType = "control-plane"
	NodeTypeWorker       NodeType = "worker"
)

type Node struct {
	Type      NodeType                 `json:"type"`
	Name      interface{}              `json:"name,omitempty"`
	Labels    interface{}              `json:"labels,omitempty"`
	Addresses interface{}              `json:"addresses,omitempty"`
	NodeSpec  NodeSpec                 `json:"node_spec,omitempty"`
	NodeInfo  onpremmodel.NodeProperty `json:"node_info,omitempty"`
}

type NodeSpec struct {
	CPU              int `json:"cpu"`               // cores
	Memory           int `json:"memory"`            // MiB
	EphemeralStorage int `json:"ephemeral_storage"` // MiB
}

type Helm struct {
	Repo    []Repo    `json:"repo"`
	Release []Release `json:"release"`
}

type Repo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Release struct {
	Name             string    `json:"name"`
	Namespace        string    `json:"namespace"`
	Revision         int       `json:"revision"`
	Updated          time.Time `json:"updated"`
	Status           string    `json:"status"`
	ChartNameVersion string    `json:"chart"`
	AppVersion       string    `json:"app_version"`
}

// Beetle defines structs

// NodeGroupInfo holds aggregated information for a node group
type NodeGroupInfo struct {
	Count       int
	TotalCPU    int
	TotalMemory int // in MiB
}

// nodeGroup holds worker nodes that share identical spec characteristics.
type nodeGroup struct {
	Nodes []onpremmodel.NodeProperty
}

// repNode returns the representative node for spec lookup (always the first).
func (g nodeGroup) repNode() onpremmodel.NodeProperty {
	return g.Nodes[0]
}

// k8sNodeGroupsOnCreation maps each known CSP to its nodeGroupsOnCreation value.
// true  → NodeGroups are created together with the cluster (k8sClusterDynamic).
// false → NodeGroups must be created separately (k8sNodeGroupDynamic) after cluster creation.
// Source: CB-Tumblebug assets/k8sclusterinfo.yaml (v0.12.13)
//
// IMPORTANT: When CB-Tumblebug adds a new CSP, add it here explicitly.
// Unknown CSPs are rejected to prevent silent misconfiguration.
var k8sNodeGroupsOnCreation = map[string]bool{
	"aws":     false,
	"alibaba": false,
	"tencent": false,
	"azure":   true,
	"gcp":     true,
	"nhn":     true,
	"ncp":     true,
	"ibm":     true,
}

// isNodeGroupsOnCreation returns whether the CSP creates NodeGroups together
// with the cluster. Returns an error for unsupported CSPs.
func isNodeGroupsOnCreation(csp string) (bool, error) {
	v, ok := k8sNodeGroupsOnCreation[strings.ToLower(csp)]
	if !ok {
		return false, fmt.Errorf("CSP %q is not supported for K8s cluster migration; update k8sNodeGroupsOnCreation map when CB-Tumblebug adds support", csp)
	}
	return v, nil
}

// nodeGroupKey is the grouping key for worker nodes.
type nodeGroupKey struct {
	VCPU         uint32
	MemoryGiB    uint64
	OSName       string
	OSVersion    string
	Architecture string
}

// filterWorkerNodes returns nodes whose Role equals "worker" (case-insensitive).
func filterWorkerNodes(nodes []onpremmodel.NodeProperty) []onpremmodel.NodeProperty {
	workers := make([]onpremmodel.NodeProperty, 0, len(nodes))
	for _, n := range nodes {
		if strings.EqualFold(n.Role, string(NodeTypeWorker)) {
			workers = append(workers, n)
		}
	}
	return workers
}

// groupNodesBySpec groups worker nodes by identical (vCPU, MemoryGiB, OS, Architecture).
// Each group becomes one K8s NodeGroup; DesiredNodeSize = len(group.Nodes).
// OnpremInfra.K8sCluster is a single pointer — all nodes belong to one cluster.
// Multi-cluster input is out of scope (Phase 3).
func groupNodesBySpec(workers []onpremmodel.NodeProperty) []nodeGroup {
	groups := make(map[nodeGroupKey]*nodeGroup)
	order := []nodeGroupKey{}

	for _, n := range workers {
		threads := n.CPU.Threads
		if threads == 0 {
			threads = 1
		}
		vcpu := n.CPU.Cpus * threads
		key := nodeGroupKey{
			VCPU:         vcpu,
			MemoryGiB:    n.Memory.TotalSize,
			OSName:       n.OS.Name,
			OSVersion:    n.OS.VersionID,
			Architecture: n.CPU.Architecture,
		}
		if _, exists := groups[key]; !exists {
			groups[key] = &nodeGroup{}
			order = append(order, key)
		}
		groups[key].Nodes = append(groups[key].Nodes, n)
	}

	result := make([]nodeGroup, 0, len(order))
	for _, k := range order {
		result = append(result, *groups[k])
	}
	return result
}

// extractMajorMinor returns "major.minor" from a full K8s version string.
// Examples: "v1.32.3" → "1.32",  "1.34.3-gke.100" → "1.34"
func extractMajorMinor(version string) string {
	clean := strings.SplitN(strings.TrimPrefix(version, "v"), "-", 2)[0]
	parts := strings.SplitN(clean, ".", 3)
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}
	return clean
}

// resolveK8sVersion selects the best matching available K8s version ID for the given CSP/region.
// Returns the newest available version if no prefix match is found.
func resolveK8sVersion(csp, region, sourceVersion string) (string, error) {
	if sourceVersion == "" {
		log.Warn().Str("csp", csp).Str("region", region).Msg("Source K8s version is empty; will use newest available")
	}

	majorMinor := extractMajorMinor(sourceVersion)

	available, err := tbclient.NewSession().ReadAvailableK8sVersion(
		strings.ToLower(csp),
		strings.ToLower(region),
	)
	if err != nil {
		return "", fmt.Errorf("failed to fetch available K8s versions: %w", err)
	}
	if len(available) == 0 {
		return "", fmt.Errorf("no available K8s versions for %s/%s", csp, region)
	}

	if majorMinor != "" {
		for _, v := range available {
			if strings.HasPrefix(v.Name, majorMinor) {
				log.Debug().Str("sourceVersion", sourceVersion).Str("matched", v.Id).Msg("Matched K8s version")
				return v.Id, nil
			}
		}
		log.Warn().Str("sourceVersion", sourceVersion).Str("majorMinor", majorMinor).
			Msg("No K8s version matching source major.minor; using newest available")
	}

	newest := available[len(available)-1].Id
	log.Debug().Str("newestVersion", newest).Msg("Using newest available K8s version")
	return newest, nil
}

// applyK8sMinimums enforces vCPU >= 2 and memoryGiB >= 4 on the given node.
// Returns a modified copy; the original is unchanged.
func applyK8sMinimums(node onpremmodel.NodeProperty) onpremmodel.NodeProperty {
	// minK8sVCPU and minK8sMemGiB are defined in resource-k8s-spec.go (package level).

	threads := node.CPU.Threads
	if threads == 0 {
		threads = 1
	}
	vcpu := node.CPU.Cpus * threads
	if vcpu < minK8sVCPU {
		log.Warn().Str("hostname", node.Hostname).Uint32("vcpu", vcpu).
			Msgf("Node vCPU (%d) below K8s minimum (%d); adjusting", vcpu, minK8sVCPU)
		node.CPU.Cpus = 1
		node.CPU.Threads = minK8sVCPU
	}

	if node.Memory.TotalSize < uint64(minK8sMemGiB) {
		log.Warn().Str("hostname", node.Hostname).Uint64("memGiB", node.Memory.TotalSize).
			Msgf("Node memory (%d GiB) below K8s minimum (%d GiB); adjusting", node.Memory.TotalSize, minK8sMemGiB)
		node.Memory.TotalSize = uint64(minK8sMemGiB)
	}

	return node
}

// RecommendK8sInfraWithDefaults builds a K8s cluster recommendation from the given source infra.
// The cluster field always carries the first worker node group's spec (required by TB for connection resolution).
// Whether that group is also created via k8sClusterDynamic depends on the CSP's nodeGroupsOnCreation:
//   - true  (Azure, GCP, …): first group created with cluster → additionalNodeGroups = groups[1:]
//   - false (AWS, Alibaba, Tencent): cluster NodeGroup is discarded by TB → additionalNodeGroups = groups[0:]
func RecommendK8sInfraWithDefaults(csp, region string, srcInfra onpremmodel.OnpremInfra) (cloudmodel.RecommendedK8sCluster, error) {
	emptyRet := cloudmodel.RecommendedK8sCluster{}

	// Step 1: validate CSP support
	nodeGroupCreatedWithCluster, err := isNodeGroupsOnCreation(csp)
	if err != nil {
		return emptyRet, err
	}

	// Step 2: filter and group worker nodes
	workers := filterWorkerNodes(srcInfra.Nodes)
	if len(workers) == 0 {
		return emptyRet, fmt.Errorf("no worker nodes found in source infra")
	}
	groups := groupNodesBySpec(workers)

	// Step 3: resolve K8s version
	srcVersion := ""
	if srcInfra.K8sCluster != nil {
		srcVersion = srcInfra.K8sCluster.Version
	}
	version, err := resolveK8sVersion(csp, region, srcVersion)
	if err != nil {
		log.Warn().Err(err).Msg("K8s version resolution failed; proceeding with empty version (TB will use default)")
	}

	// Step 4: build K8sClusterDynamicReq using first group's spec.
	// TB always needs a valid SpecId to resolve ConnectionName/Provider/Region.
	firstGroup := groups[0]
	firstSpec, firstImage, err := recommendK8sNodeGroupResources(csp, region, firstGroup.repNode())
	if err != nil {
		return emptyRet, fmt.Errorf("failed to recommend resources for first node group: %w", err)
	}

	cluster := cloudmodel.K8sClusterDynamicReq{
		Name:            fmt.Sprintf("k8s-%s-%s", strings.ToLower(csp), strings.ToLower(region)),
		NodeGroupName:   "worker-group-1",
		ConnectionName:  fmt.Sprintf("%s-%s", strings.ToLower(csp), strings.ToLower(region)),
		Version:         version,
		SpecId:          firstSpec,
		ImageId:         firstImage,
		DesiredNodeSize: len(firstGroup.Nodes),
		MinNodeSize:     1,
		MaxNodeSize:     len(firstGroup.Nodes),
		OnAutoScaling:   "false",
		RootDiskType:    "default",
		RootDiskSize:    int(firstGroup.repNode().RootDisk.TotalSize),
	}

	// Step 5: build additionalNodeGroups based on CSP nodeGroupsOnCreation.
	// - nodeGroupCreatedWithCluster=true : first group is created with the cluster → start from groups[1:]
	// - nodeGroupCreatedWithCluster=false: TB discards the cluster's NodeGroup → include all groups[0:]
	var additionalStartIdx int
	if nodeGroupCreatedWithCluster {
		additionalStartIdx = 1
	} else {
		additionalStartIdx = 0
	}

	additionalGroups := make([]cloudmodel.K8sNodeGroupDynamicReq, 0, len(groups)-additionalStartIdx)
	for i, g := range groups[additionalStartIdx:] {
		specId, imageId, err := recommendK8sNodeGroupResources(csp, region, g.repNode())
		if err != nil {
			log.Warn().Err(err).Int("groupIndex", i+additionalStartIdx+1).Msg("Failed to recommend resources for node group; skipping")
			continue
		}
		additionalGroups = append(additionalGroups, cloudmodel.K8sNodeGroupDynamicReq{
			Name:            fmt.Sprintf("worker-group-%d", i+additionalStartIdx+1),
			SpecId:          specId,
			ImageId:         imageId,
			DesiredNodeSize: len(g.Nodes),
			MinNodeSize:     1,
			MaxNodeSize:     len(g.Nodes),
			OnAutoScaling:   "false",
			RootDiskType:    "default",
			RootDiskSize:    int(g.repNode().RootDisk.TotalSize),
		})
	}

	log.Info().
		Str("csp", csp).
		Str("region", region).
		Str("version", version).
		Bool("nodeGroupsOnCreation", nodeGroupCreatedWithCluster).
		Int("nodeGroups", len(groups)).
		Int("additionalNodeGroups", len(additionalGroups)).
		Msg("K8s infrastructure recommendation completed")

	return cloudmodel.RecommendedK8sCluster{
		Cluster:              cluster,
		AdditionalNodeGroups: additionalGroups,
	}, nil
}

// recommendK8sNodeGroupResources returns the specId and imageId for a worker node.
func recommendK8sNodeGroupResources(csp, region string, node onpremmodel.NodeProperty) (specId, imageId string, err error) {
	specs, _, err := RecommendK8sSpecs(csp, region, node, 1)
	if err != nil || len(specs) == 0 {
		return "", "", fmt.Errorf("no K8s specs found for node (hostname: %s): %w", node.Hostname, err)
	}
	specId = specs[0].Id

	imageId, err = RecommendK8sOsImage(csp, region, node)
	if err != nil {
		log.Warn().Err(err).Str("hostname", node.Hostname).Msg("K8s image recommendation failed; using default")
		imageId = "default"
		err = nil
	}

	return specId, imageId, nil
}

// Deprecated: RecommendK8sControlPlane is replaced by RecommendK8sInfraWithDefaults.
func RecommendK8sControlPlane(provider, region string, k8sInfoList KubernetesInfoList) (cloudmodel.K8sClusterDynamicReq, error) {
	log.Warn().Msg("RecommendK8sControlPlane is deprecated; use RecommendK8sInfraWithDefaults instead")
	return cloudmodel.K8sClusterDynamicReq{}, fmt.Errorf("RecommendK8sControlPlane is deprecated")
}

// Deprecated: RecommendK8sNodeGroup is replaced by RecommendK8sInfraWithDefaults.
func RecommendK8sNodeGroup(provider, region string, k8sInfoList KubernetesInfoList) (cloudmodel.K8sNodeGroupDynamicReq, error) {
	log.Warn().Msg("RecommendK8sNodeGroup is deprecated; use RecommendK8sInfraWithDefaults instead")
	return cloudmodel.K8sNodeGroupDynamicReq{}, fmt.Errorf("RecommendK8sNodeGroup is deprecated")
}
