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

// Package recommendation contains the core logic for recommending K8s infra configurations.
package recommendation

import (
	"fmt"
	"net"
	"regexp"
	"sort"
	"strconv"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"
)

// RecommendK8sInfra builds a complete K8s infra recommendation from on-premise infra data.
func RecommendK8sInfra(provider, region string, onpremInfra onpremmodel.OnpremInfra) (cloudmodel.RecommendedInfra, error) {
	emptyRet := cloudmodel.RecommendedInfra{}

	if onpremInfra.K8sCluster == nil {
		return emptyRet, fmt.Errorf("source infra has no K8s cluster information")
	}

	workerNodes := collectWorkerNodes(onpremInfra.Nodes)
	if len(workerNodes) == 0 {
		return emptyRet, fmt.Errorf("no worker nodes found in source K8s cluster")
	}

	// Read everything the target cloud declares before any per-worker work begins. Keeping the
	// lookups here — rather than inside the helpers that consume them — is what makes the rest of
	// this function, and those helpers, free of hidden I/O.
	profile := getTargetProfile(provider, region)

	// Versions come from a different endpoint than the K8s asset, so they are read separately
	// rather than folded into the profile: the L0 APIs need the profile but not the versions.
	available, err := tbclient.NewSession().GetAvailableK8sVersions(provider, region)
	if err != nil {
		return emptyRet, fmt.Errorf("failed to fetch available K8s versions from Tumblebug: %w", err)
	}

	version, err := selectK8sVersion(available, provider, onpremInfra.K8sCluster.Version)
	if err != nil {
		return emptyRet, fmt.Errorf("K8s version selection failed: %w", err)
	}

	connectionName := fmt.Sprintf("%s-%s", strings.ToLower(provider), strings.ToLower(region))

	// Derive cluster name from source cluster name; fall back to a generic default.
	clusterName := "migrated-k8s"
	if onpremInfra.K8sCluster.Name != "" {
		clusterName = onpremInfra.K8sCluster.Name
	}

	// Recommend per worker, then merge on the resolved target — not the other way round.
	//
	// Managed K8s node groups are homogeneous (EKS managed node group, AKS agent pool, GKE node
	// pool), so a node group can hold only workers that share one spec. Grouping on the *target*
	// rather than on the source signature means two workers differing by a single vCPU that
	// resolve to the same instance type share one node group instead of producing two.
	//
	// Ask for the same candidate pool the L0 spec API uses, not just one spec. Requesting a
	// single spec would delegate the choice to Tumblebug's cost priority and leave
	// sortK8sSpecsByProximity sorting a one-element slice — so the two APIs could recommend
	// different specs for the same worker whenever the cheapest candidate in range is not the
	// closest one.
	groups, failed := recommendNodeGroups(profile, normalizeWorkers(workerNodes), GetDefaultSpecsLimit())
	logExcludedWorkers(failed)
	if len(groups) == 0 {
		return emptyRet, fmt.Errorf("no K8s worker node group could be recommended: %s", summarizeFailures(failed))
	}

	nodeGroupReqs := make([]cloudmodel.K8sNodeGroupReq, 0, len(groups))
	includedWorkers := 0
	for i, g := range groups {
		// Node group names use a fixed base ("workers") + index, normalized and checked against
		// the CSP naming rule the profile carries.
		name := buildNodeGroupName(profile, "workers", i+1)
		nodeGroupReqs = append(nodeGroupReqs, buildK8sNodeGroupReq(profile, name, g))
		includedWorkers += len(g.nodes)
	}

	vNetReq := buildK8sVNetReq(profile, connectionName, onpremInfra)
	sshKeyReq := buildK8sSshKeyReq(connectionName)
	sgReqList := buildK8sSecurityGroupReqList(connectionName, provider, region, vNetReq.CidrBlock, workerNodes, onpremInfra.K8sCluster)

	clusterReq := cloudmodel.K8sClusterReq{
		ConnectionName: connectionName,
		Name:           clusterName,
		Version:        version,
		// VNetId, SubnetIds, SecurityGroupIds are filled in by migration logic
		K8sNodeGroupList: nodeGroupReqs,
		// Count the workers actually covered, not every worker in the source: excluded ones
		// (see below) are not in any node group.
		Description: fmt.Sprintf("Migrated from on-premise K8s cluster (v%s, %d workers)", onpremInfra.K8sCluster.Version, includedWorkers),
	}

	// A worker that could not be resolved is excluded rather than failing the whole
	// recommendation, so the exclusion has to be visible in the response — dropping it silently
	// would be worse than the hard failure this replaced.
	description := fmt.Sprintf("K8s cluster recommendation for %s %s (source: v%s → target: v%s)",
		provider, region, onpremInfra.K8sCluster.Version, version)
	if note := summarizeFailures(failed); note != "" {
		description += ". " + note
	}

	result := cloudmodel.RecommendedInfra{
		Status:      "recommended",
		Description: description,
		TargetCloud: cloudmodel.CloudProperty{
			Csp:    strings.ToLower(provider),
			Region: strings.ToLower(region),
		},
		TargetVNet:              vNetReq,
		TargetSshKey:            sshKeyReq,
		TargetSecurityGroupList: sgReqList,
		TargetK8sCluster:        clusterReq,
	}

	log.Info().
		Str("provider", provider).
		Str("region", region).
		Str("version", version).
		Int("nodeGroupCount", len(nodeGroupReqs)).
		Int("workerCount", len(workerNodes)).
		Int("includedWorkers", includedWorkers).
		Int("excludedWorkers", len(failed)).
		Msg("K8s infra recommendation completed")

	return result, nil
}

// collectWorkerNodes filters nodes with role "worker" or "Worker".
func collectWorkerNodes(nodes []onpremmodel.NodeProperty) []onpremmodel.NodeProperty {
	var workers []onpremmodel.NodeProperty
	for _, n := range nodes {
		if strings.EqualFold(n.Role, "worker") {
			workers = append(workers, n)
		}
	}
	return workers
}

// sourceVcpuOf returns the source node's vCPU count. One rule serves both the memoization key
// and the spec search, so a node can never be keyed by one value and sized by another.
//
// Cpus == 0 means "totals, not per-socket": a node discovered through the Kubernetes API arrives
// as {cpus: 0, cores: 0, threads: N} because honeybee lands the API's total CPU count in Threads.
// Treat it as a single socket. Threads == 0 falls back to Cores (assume no SMT) rather than 1,
// since a bare 1 understates a multi-core socket by its full core count.
func sourceVcpuOf(n onpremmodel.NodeProperty) uint32 {
	cpus := n.CPU.Cpus
	if cpus == 0 {
		cpus = 1
	}
	threads := n.CPU.Threads
	if threads == 0 {
		threads = n.CPU.Cores
	}
	if threads == 0 {
		threads = 1
	}
	return cpus * threads
}

// workerRequirement is one source worker translated into the sizing the target must satisfy.
//
// Three things happen while building it, all of them pure — no target cloud is consulted:
//   - vCPU is computed by one rule (sourceVcpuOf), so collectors that report CPU differently
//     produce the same number for the same machine.
//   - Architecture aliases are canonicalized (amd64/x64 -> x86_64).
//   - Sizing is floored at minViableWorkerVcpu/MemGiB. This *changes* the value rather than
//     canonicalizing it: below that floor a node is not a viable managed K8s worker and AKS
//     rejects it outright. upscaleNote records the change so it never happens silently.
//
// What is NOT decided here: whether any of it can be bought. Which specs exist in a region is a
// target lookup, so it belongs to the spec search — see RecommendK8sNodeSpecs.
//
// The source node travels with the requirement because node groups report their members by
// machine id, and because the failure path needs to name the worker that could not be resolved.
type workerRequirement struct {
	node        onpremmodel.NodeProperty
	vcpu        uint32
	memGiB      uint32
	arch        string
	upscaleNote string // non-empty only when the floor raised the request
}

// normalizeWorker turns one source worker into the sizing the target must satisfy.
// See workerRequirement for what "normalize" covers, and what it deliberately does not.
func normalizeWorker(n onpremmodel.NodeProperty) workerRequirement {
	srcVcpu := sourceVcpuOf(n)
	srcMem := uint32(n.Memory.TotalSize)

	req := workerRequirement{
		node:   n,
		vcpu:   srcVcpu,
		memGiB: srcMem,
		arch:   normalizeArch(n.CPU.Architecture),
	}
	if req.vcpu < minViableWorkerVcpu {
		req.vcpu = minViableWorkerVcpu
	}
	if uint64(req.memGiB) < minViableWorkerMemGiB {
		req.memGiB = uint32(minViableWorkerMemGiB)
	}

	if req.vcpu != srcVcpu || req.memGiB != srcMem {
		req.upscaleNote = fmt.Sprintf(
			"(worker spec upscaled from source %dvCPU/%dGiB to %dvCPU/%dGiB — the minimum node size "+
				"accepted by the target K8s node recommendation. A node this small leaves little allocatable "+
				"capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods "+
				"may fail to schedule at the source size.)",
			srcVcpu, srcMem, req.vcpu, req.memGiB)
		log.Warn().
			Uint32("sourceVcpu", srcVcpu).Uint32("sourceMemGiB", srcMem).
			Uint32("minVcpu", req.vcpu).Uint32("minMemGiB", req.memGiB).
			Msg("Source worker below minimum viable K8s worker spec; upscaling recommendation")
	}
	return req
}

// normalizeWorkers applies normalizeWorker to every worker, preserving source order.
func normalizeWorkers(nodes []onpremmodel.NodeProperty) []workerRequirement {
	reqs := make([]workerRequirement, 0, len(nodes))
	for _, n := range nodes {
		reqs = append(reqs, normalizeWorker(n))
	}
	return reqs
}

// workerSpecKey is the memoization key for spec recommendation — NOT a partition key.
//
// Two workers sharing a key resolve to the same target, so the recommendation is computed once.
// Two workers with different keys may still end up in the same node group: merging happens on the
// resolved target, not here (see groupWorkersIntoNodeGroups). That is what stops a 1-unit difference in
// the source from producing a separate node group.
// Keying on the normalized requirement rather than the raw node means two workers that differ
// only below the floor (1vCPU/1GiB and 2vCPU/4GiB both become 2vCPU/4GiB) share one search.
func workerSpecKey(req workerRequirement) string {
	return fmt.Sprintf("%d|%d|%s", req.vcpu, req.memGiB, req.arch)
}

// workerTarget is one source worker resolved against the target cloud: the ranked specs it maps
// onto, the node image its architecture requires, and any note that must reach the user.
//
// A non-nil err means this worker could not be resolved. It is excluded from the node groups and
// reported back to the caller rather than aborting the whole recommendation — one worker without
// an available node image must not cost the caller the entire cluster.
type workerTarget struct {
	node        onpremmodel.NodeProperty
	specs       []cloudmodel.SpecInfo // ranked; specs[0] is the chosen one
	imageId     string
	upscaleNote string
	err         error
}

// specId returns the chosen target spec id, or "" when the worker failed to resolve.
func (t workerTarget) specId() string {
	if t.err != nil || len(t.specs) == 0 {
		return ""
	}
	return t.specs[0].Id
}

// recommendSpecsAndImages finds the target spec candidates and node image for every worker, in
// source order. Workers that could not be resolved are returned separately rather than aborting
// the whole recommendation — one worker without an available node image must not cost the caller
// the entire cluster. The exclusion is reported by logExcludedWorkers and summarizeFailures.
//
// It recommends specs and images, not node groups: it compares nothing across workers, so whether
// two of them share a node group is not decided here. That happens downstream, on the resolved
// target — see recommendNodeGroups.
//
// The recommendation is memoized by workerSpecKey so identical workers cost one Tumblebug call, which
// keeps the call count at what the previous group-then-recommend flow used. The key is a cache key
// only — see workerSpecKey.
//
// The memo lives for one call, during which limit is fixed, so limit is not part of the key. Add
// it if a future caller varies limit per worker.
func recommendSpecsAndImages(p targetProfile, reqs []workerRequirement, limit int) (candidates, failed []workerTarget) {

	memo := make(map[string]workerTarget, len(reqs))
	targets := make([]workerTarget, 0, len(reqs))

	for _, req := range reqs {
		key := workerSpecKey(req)
		t, cached := memo[key]
		if cached {
			// The resolved target is shared, but the node and its note are per worker: two
			// workers can share a key and still have been floored from different source sizes.
			t.node, t.upscaleNote = req.node, req.upscaleNote
		} else {
			t = resolveWorkerTarget(p, req, limit)
			memo[key] = t
		}
		targets = append(targets, t)
	}
	return splitResolved(targets)
}

// splitResolved separates the workers that resolved from those that did not, preserving source
// order in both. Split here rather than during grouping so the two concerns stay apart: a failed
// worker has no target to be grouped on.
func splitResolved(targets []workerTarget) (candidates, failed []workerTarget) {
	for _, t := range targets {
		if t.err != nil {
			failed = append(failed, t)
			continue
		}
		candidates = append(candidates, t)
	}
	return candidates, failed
}

// resolveWorkerTarget searches for one worker's spec candidates and node image. A non-nil err on
// the result means this worker could not be resolved; it is never returned as a Go error because
// the caller keeps going with the workers that did resolve.
func resolveWorkerTarget(p targetProfile, req workerRequirement, limit int) workerTarget {
	t := workerTarget{node: req.node, upscaleNote: req.upscaleNote}

	specs, err := RecommendK8sNodeSpecs(p.provider, p.region, req, limit)
	switch {
	case err != nil:
		// Keep machine IDs out of the error text: failures are grouped by cause and the
		// affected IDs are listed separately (see groupFailuresByCause).
		t.err = fmt.Errorf("spec recommendation failed: %w", err)
	case len(specs) == 0:
		t.err = fmt.Errorf("no target spec found within the search range")
	default:
		t.specs = specs
		// The spec search already filtered on this architecture, so the image must match it:
		// an ARM spec needs an ARM image, not the x86 default.
		imageId, imgErr := selectK8sNodeImage(p, req.arch)
		if imgErr != nil {
			t.err = fmt.Errorf("node image selection failed: %w", imgErr)
		} else {
			t.imageId = imageId
		}
	}
	return t
}

// nodeGroupAccum accumulates the source workers that resolved to one target (spec, image) pair.
type nodeGroupAccum struct {
	specs   []cloudmodel.SpecInfo // ranked; specs[0] is the chosen target spec
	imageId string
	nodes   []onpremmodel.NodeProperty
	notes   []string
}

// spec returns the node group's chosen target spec.
func (g nodeGroupAccum) spec() cloudmodel.SpecInfo {
	if len(g.specs) == 0 {
		return cloudmodel.SpecInfo{}
	}
	return g.specs[0]
}

// specId returns the node group's target spec id.
func (g nodeGroupAccum) specId() string {
	if len(g.specs) == 0 {
		return ""
	}
	return g.specs[0].Id
}

// groupWorkersIntoNodeGroups turns resolved workers into homogeneous node groups, keyed by the
// target (specId, imageId) pair — never by source attributes. Two source workers differing by a
// single vCPU that resolve to the same instance type therefore share one node group.
//
// This is a group-by, not a decision: same target means same node group, unconditionally. It is
// also the only step that changes the unit of the data, from per worker to per node group.
//
// It cannot be folded into mergeNodeGroupsByUpsizing, whose criteria ask "how much does promoting
// cost" and would wrongly refuse workers that are already on the same spec. Say a 12vCPU/24GiB
// requirement and a 16vCPU/64GiB one both land on m5.4xlarge (16/64) because no 12-vCPU spec
// exists. canMergeByUpsizing sees memory utilization 24/64 = 0.375 and refuses — leaving two node
// groups with the identical spec, which a managed node group has no reason to have. The unused
// 40GiB came from catalog rounding, not from promotion, so charging it to promotion is wrong.
//
// Architecture needs no separate key: specId implies it (a spec belongs to exactly one
// architecture), so arm64 and x86_64 workers can never merge and every member of a group shares
// one node image.
//
// Groups are ordered by first appearance in the source node list, so node group indices — and
// therefore names — are deterministic across runs.
func groupWorkersIntoNodeGroups(candidates []workerTarget) []nodeGroupAccum {
	index := make(map[string]int)
	var groups []nodeGroupAccum

	for _, t := range candidates {
		key := t.specId() + "|" + t.imageId
		if i, ok := index[key]; ok {
			groups[i].nodes = append(groups[i].nodes, t.node)
			continue
		}
		index[key] = len(groups)
		groups = append(groups, nodeGroupAccum{
			specs:   t.specs,
			imageId: t.imageId,
			nodes:   []onpremmodel.NodeProperty{t.node},
			notes:   noteSlice(t.upscaleNote),
		})
	}
	return groups
}

// recommendNodeGroups recommends the node groups for a set of normalized workers.
//
// This is where a node group recommendation is finished; the three steps below are its
// implementation, not three separate recommendations. Assembling them here rather than at the call
// sites is what stops RecommendK8sInfra and the L0 spec API from ending up with different node
// group breakdowns for the same input — they used to derive groups independently and could
// disagree.
func recommendNodeGroups(p targetProfile, reqs []workerRequirement, limit int) ([]nodeGroupAccum, []workerTarget) {
	candidates, failed := recommendSpecsAndImages(p, reqs, limit) // find specs and images     (I/O)
	groups := groupWorkersIntoNodeGroups(candidates)              // workers -> node groups    (pure)
	return mergeNodeGroupsByUpsizing(groups), failed              // node groups -> fewer ones (pure)
}

// Consolidation bounds. A worker is only promoted into a larger group's spec when the result stays
// within both of these — they guard different risks and neither subsumes the other.
//
//   - upsizeUtilFloor bounds the *ratio*. At 0.5 a worker may rise by at most one catalog
//     step, since instance families roughly double (4 → 8 → 16 vCPU): a second step would leave it
//     at 25%. Industry sizing guidance treats anything under ~60% as over-provisioned and Azure
//     Migrate's default buffer is 1.3x, so 2x is already generous; it is the price of removing a
//     node group, not a headroom target.
//   - maxWastedVcpu / maxWastedMemGiB bound the *absolute* amount. The ratio alone leaks at the top
//     of the catalog, where steps shrink to 1.5x and 1.33x: 32 → 64 vCPU spans two steps yet still
//     computes to exactly 50%, wasting 32 vCPU. The caps stop that. 16 vCPU is the largest single
//     step below 64 vCPU, so one step stays allowed through that range.
//
// Neither value is derived from a published standard; both are judgment calls to revisit once real
// cluster data exists. See docs/k8s-recommendation/k8s-node-group-consolidation-analysis.md.
const (
	upsizeUtilFloor = 0.5
	maxWastedVcpu   = 16
	maxWastedMemGiB = 64
)

// mergeNodeGroupsByUpsizing folds node groups into their larger neighbours where the smaller
// group's workers would still be reasonably sized in the larger spec.
//
// Despite sitting last, this is a recommendation and not a tidy-up: it changes which spec the
// promoted workers are bought at. It is also the only step that reduces the node group count —
// groupWorkersIntoNodeGroups decides the starting count, this decides the final one.
//
// It never consults the catalog. The promotion target is the largest spec already present in the
// run, so no worker can land on a spec that no worker resolved to, and the function stays pure.
//
// Managed node groups are homogeneous, so consolidating means promoting the smaller workers to the
// larger spec — capacity they will not use and do pay for. That is why it is bounded rather than
// unconditional: merging everything into the largest spec costs far more than it saves.
//
// Groups are sorted smallest-first and each run is extended while the whole run still passes, so a
// chain cannot slip past the bounds by merging pairwise. Merging the small end first is also where
// consolidation is cheapest in absolute terms.
func mergeNodeGroupsByUpsizing(groups []nodeGroupAccum) []nodeGroupAccum {
	if len(groups) < 2 {
		return groups
	}

	// Size order drives the merging, but the caller's order is restored afterwards: consolidation
	// should change which node groups exist, not the order they are reported in. Node group names
	// are index-based, so reordering would rename groups for no reason.
	type ranked struct {
		nodeGroupAccum
		origin int // first-appearance index, used to restore the caller's order
	}
	ordered := make([]ranked, len(groups))
	for i, g := range groups {
		ordered[i] = ranked{nodeGroupAccum: g, origin: i}
	}
	sort.SliceStable(ordered, func(i, j int) bool {
		a, b := ordered[i].spec(), ordered[j].spec()
		if a.VCPU != b.VCPU {
			return a.VCPU < b.VCPU
		}
		return a.MemoryGiB < b.MemoryGiB
	})

	bySize := make([]nodeGroupAccum, len(ordered))
	for i, r := range ordered {
		bySize[i] = r.nodeGroupAccum
	}

	var out []ranked
	for i := 0; i < len(bySize); {
		end := i
		for end+1 < len(bySize) && canMergeByUpsizing(bySize[i:end+2]) {
			end++
		}
		// A merged run takes the earliest origin among its members, so it lands where the first of
		// those groups used to be.
		origin := ordered[i].origin
		for _, r := range ordered[i : end+1] {
			if r.origin < origin {
				origin = r.origin
			}
		}
		out = append(out, ranked{nodeGroupAccum: mergeRun(bySize[i : end+1]), origin: origin})
		i = end + 1
	}

	sort.Slice(out, func(i, j int) bool { return out[i].origin < out[j].origin })
	restored := make([]nodeGroupAccum, len(out))
	for i, r := range out {
		restored[i] = r.nodeGroupAccum
	}
	return restored
}

// canMergeByUpsizing reports whether every worker in the run would still be acceptably sized in
// the run's largest spec — that is, whether the promotion is worth what it costs. The check is against the whole run, not pairs, so extending a run re-tests
// the workers already in it.
//
// Architecture is not compared: a spec belongs to exactly one architecture, so an arm64 and an
// x86_64 group can never satisfy the size checks against a single spec anyway — and the node image
// travels with the spec.
func canMergeByUpsizing(run []nodeGroupAccum) bool {
	target := run[len(run)-1].spec()
	if target.VCPU == 0 || target.MemoryGiB == 0 {
		return false
	}
	// An arm64 spec cannot host x86_64 workers or vice versa, whatever the sizes say. An unknown
	// architecture is treated the same: consolidation is optional, so guessing is never worth it.
	if target.Architecture == "" {
		return false
	}
	for _, g := range run {
		if !strings.EqualFold(g.spec().Architecture, target.Architecture) {
			return false
		}
	}

	minVcpu, minMem := minSourceSize(run)

	// Ratio: how much of the target each worker would actually use.
	if float64(minVcpu)/float64(target.VCPU) < upsizeUtilFloor {
		return false
	}
	if float64(minMem)/float64(target.MemoryGiB) < upsizeUtilFloor {
		return false
	}
	// Absolute: how much capacity is bought and left idle.
	if float64(target.VCPU)-float64(minVcpu) > maxWastedVcpu {
		return false
	}
	if float64(target.MemoryGiB)-float64(minMem) > maxWastedMemGiB {
		return false
	}
	return true
}

// minSourceSize returns the smallest source vCPU and memory across every worker in the run. The
// smallest worker has the lowest utilization, so checking it covers the rest; the two dimensions
// are taken independently because different workers may hold each minimum.
//
// Deliberately the *source* size, not the normalized requirement. Reading the requirement would
// let the two upsizing mechanisms compound: the floor has already raised a 1vCPU/2GiB worker to
// 2vCPU/4GiB, and measuring from there would let consolidation raise it once more, to 4vCPU/8GiB.
// Each step passes its own 50% gate while the worker ends up at 25% utilization — exactly what
// upsizeUtilFloor exists to prevent. The floor's own upsizing is disclosed separately, through
// workerRequirement.upscaleNote.
func minSourceSize(run []nodeGroupAccum) (vcpu uint32, memGiB uint64) {
	for _, g := range run {
		for _, n := range g.nodes {
			v, m := sourceVcpuOf(n), n.Memory.TotalSize
			if vcpu == 0 || v < vcpu {
				vcpu = v
			}
			if memGiB == 0 || m < memGiB {
				memGiB = m
			}
		}
	}
	return vcpu, memGiB
}

// mergeRun collapses a run of node groups into one, keeping the largest spec and image. Promoted
// workers are recorded in the notes so the response can say what was consolidated and why.
func mergeRun(run []nodeGroupAccum) nodeGroupAccum {
	last := run[len(run)-1]
	if len(run) == 1 {
		return last
	}

	merged := nodeGroupAccum{specs: last.specs, imageId: last.imageId}
	for _, g := range run {
		merged.nodes = append(merged.nodes, g.nodes...)
		merged.notes = append(merged.notes, g.notes...)
	}

	target := last.spec()
	var promoted []string
	for _, g := range run[:len(run)-1] {
		promoted = append(promoted, fmt.Sprintf("%s (%dvCPU/%dGiB)",
			g.spec().CspSpecName, g.spec().VCPU, uint64(g.spec().MemoryGiB)))
	}
	merged.notes = append(merged.notes, fmt.Sprintf(
		"(consolidated into %s: %s merged in, so these workers share one node group instead of %d. "+
			"They are provisioned larger than their source, which is the cost of the smaller node "+
			"group count.)",
		target.CspSpecName, strings.Join(promoted, ", "), len(run)))
	return merged
}

// noteSlice wraps a possibly-empty note so a group starts with either zero or one note.
func noteSlice(note string) []string {
	if note == "" {
		return nil
	}
	return []string{note}
}

// failureGroup is a set of workers excluded for the same reason.
type failureGroup struct {
	cause      string
	machineIds []string
}

// groupFailuresByCause groups excluded workers by error text so one shared cause is reported once
// with the affected machine IDs listed, rather than repeating the same sentence per worker.
// First-appearance order keeps the output deterministic.
func groupFailuresByCause(failed []workerTarget) []failureGroup {
	index := make(map[string]int)
	var groups []failureGroup

	for _, t := range failed {
		cause := "unknown error"
		if t.err != nil {
			cause = t.err.Error()
		}
		if i, ok := index[cause]; ok {
			groups[i].machineIds = append(groups[i].machineIds, t.node.MachineId)
			continue
		}
		index[cause] = len(groups)
		groups = append(groups, failureGroup{cause: cause, machineIds: []string{t.node.MachineId}})
	}
	return groups
}

// summarizeFailures renders excluded workers as one sentence, grouped by cause.
// Returns "" when nothing was excluded, so callers can append it unconditionally.
func summarizeFailures(failed []workerTarget) string {
	if len(failed) == 0 {
		return ""
	}
	groups := groupFailuresByCause(failed)
	parts := make([]string, 0, len(groups))
	for _, g := range groups {
		parts = append(parts, fmt.Sprintf("%s (%s)", g.cause, strings.Join(g.machineIds, ", ")))
	}
	return fmt.Sprintf("%d source worker(s) excluded: %s", len(failed), strings.Join(parts, "; "))
}

// logExcludedWorkers reports excluded workers once per cause at warn level.
func logExcludedWorkers(failed []workerTarget) {
	for _, g := range groupFailuresByCause(failed) {
		log.Warn().
			Str("cause", g.cause).
			Strs("machineIds", g.machineIds).
			Msg("Source workers excluded from K8s recommendation")
	}
}

// archAliases maps the architecture strings on-premise sources report to the canonical values
// Tumblebug stores. Tumblebug normalizes every CSP's own representation into a fixed set on spec
// registration (e.g. IBM's "amd64" becomes "x86_64"), and the spec search filter matches the
// stored value exactly (case-insensitively), so an un-normalized alias silently matches nothing.
// Notably "aarch64" — what `uname -m` prints on ARM Linux — is never a stored value.
var archAliases = map[string]string{
	"":        "x86_64",
	"amd64":   "x86_64",
	"x64":     "x86_64",
	"x86-64":  "x86_64",
	"x86_64":  "x86_64",
	"arm64":   "arm64",
	"aarch64": "arm64",
	"armv8":   "arm64",
	"armv8l":  "arm64",
}

// normalizeArch maps a source architecture string to the value Tumblebug expects for spec search
// and node image selection. Unknown values pass through lowercased rather than being forced to a
// default, so an unsupported architecture fails loudly at spec search instead of silently
// recommending a mismatched (x86) spec.
func normalizeArch(arch string) string {
	lower := strings.ToLower(strings.TrimSpace(arch))
	if canonical, ok := archAliases[lower]; ok {
		return canonical
	}
	return lower
}

// selectK8sVersion picks the minimum CSP-supported K8s version that is ≥ the source version.
// Falls back to the latest available version if none is ≥ source.
//
// Azure workaround:
// Azure's availableK8sVersion includes LTS-only versions (e.g. "1.33.12") alongside
// standard versions, without distinguishing them. Using an LTS-only version without
// explicitly enabling the LTS support plan causes a 400 K8sVersionNotSupported error.
// As a temporary measure, Azure always skips the minimum satisfying version and selects
// the next higher minor version instead.
//
// TODO: Remove this workaround once CB-Tumblebug exposes an isLts field in
// availableK8sVersion. At that point, filter by isLts=false and remove the Azure branch.
func selectK8sVersion(available []tbmodel.K8sClusterVersionDetailAvailable, provider, sourceVersion string) (string, error) {
	if len(available) == 0 {
		return "", fmt.Errorf("no K8s versions available for provider %q", provider)
	}

	sourceMajorMinor := extractMajorMinor(sourceVersion)

	// Find minimum available version ≥ source.
	// Compare by v.Name (major.minor), but return v.Id (full patch accepted by CB-Spider).
	type candidate struct{ name, id string }
	var best *candidate
	for _, v := range available {
		vMajorMinor := extractMajorMinor(v.Name)
		if compareVersion(vMajorMinor, sourceMajorMinor) >= 0 {
			if best == nil || compareVersion(vMajorMinor, extractMajorMinor(best.name)) < 0 {
				best = &candidate{name: v.Name, id: v.Id}
			}
		}
	}

	if best == nil {
		// All available versions are older than source; use the highest available.
		// Do not rely on the ordering of the API response (Tumblebug returns descending,
		// so available[len-1] would be the *oldest*, not the latest).
		latest := available[0]
		for _, v := range available {
			if compareVersion(extractMajorMinor(v.Name), extractMajorMinor(latest.Name)) > 0 {
				latest = v
			}
		}
		best = &candidate{name: latest.Name, id: latest.Id}
		log.Warn().
			Str("sourceVersion", sourceVersion).
			Str("selectedVersion", best.id).
			Msg("Source version exceeds all CSP-supported versions; using latest available")
	}

	// Azure LTS workaround: skip the minimum satisfying version and use the next minor.
	// Azure's available version list includes LTS-only versions without marking them as
	// such. The minimum version (e.g. "1.33.12") is often LTS-only in a given region,
	// so we step up one minor to get a standard (non-LTS) version (e.g. "1.34.8").
	if strings.EqualFold(provider, "azure") {
		var next *candidate
		for _, v := range available {
			vMajorMinor := extractMajorMinor(v.Name)
			if compareVersion(vMajorMinor, extractMajorMinor(best.name)) > 0 {
				if next == nil || compareVersion(vMajorMinor, extractMajorMinor(next.name)) < 0 {
					next = &candidate{name: v.Name, id: v.Id}
				}
			}
		}
		if next != nil {
			log.Info().
				Str("skipped", best.id).
				Str("selected", next.id).
				Msg("Azure: skipped minimum version (may be LTS-only); using next minor version")
			best = next
		} else {
			log.Warn().
				Str("version", best.id).
				Msg("Azure: no higher version available; using minimum (may require LTS plan)")
		}
	}

	log.Info().
		Str("sourceVersion", sourceVersion).
		Str("selectedVersion", best.id).
		Msg("K8s version selected")
	return best.id, nil
}

// extractMajorMinor extracts "major.minor" from a version string like "1.32.3" or "1.33".
func extractMajorMinor(version string) string {
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		return parts[0] + "." + parts[1]
	}
	return version
}

// compareVersion compares two "major.minor" strings numerically.
// Returns negative if a < b, 0 if equal, positive if a > b.
func compareVersion(a, b string) int {
	partsA := strings.Split(a, ".")
	partsB := strings.Split(b, ".")

	for i := 0; i < 2; i++ {
		var numA, numB int
		if i < len(partsA) {
			numA, _ = strconv.Atoi(partsA[i])
		}
		if i < len(partsB) {
			numB, _ = strconv.Atoi(partsB[i])
		}
		if numA != numB {
			return numA - numB
		}
	}
	return 0
}

// minViableWorkerVcpu / minViableWorkerMemGiB is the minimum viable managed K8s worker spec.
// Nodes below this cannot reliably host the kubelet + system daemonsets, and Azure AKS rejects
// sub-2vCPU/4GiB VMs for the mandatory system node pool. Applied as a floor to every node group
// across all CSPs so the recommendation is viable regardless of node-group ordering.
const (
	minViableWorkerVcpu   uint32 = 2
	minViableWorkerMemGiB uint64 = 4
)

// buildK8sVNetReq builds a VNetReq for the K8s cluster's VPC.
//
// It first tries to derive the CIDR from the source on-premise network topology via
// RecommendVNet, then checks that the chosen CIDR does not overlap with the cluster's
// PodCIDR or ServiceCIDR. If network data is absent or recommendation fails, it falls
// back to the hardcoded 10.10.0.0/16 default so recommendation remains resilient.
func buildK8sVNetReq(p targetProfile, connectionName string, srcInfra onpremmodel.OnpremInfra) cloudmodel.VNetReq {
	const fallbackCIDR = "10.10.0.0/16"

	provider, region := p.provider, p.region
	subnetCount := p.requiredSubnetCount
	zones := p.regionZones

	buildVNet := func(vpcCIDR string) cloudmodel.VNetReq {
		subnets := make([]cloudmodel.SubnetReq, 0, subnetCount)
		for i := 0; i < subnetCount; i++ {
			subnet := cloudmodel.SubnetReq{
				Name:      fmt.Sprintf("k8s-subnet-%c", 'a'+i),
				IPv4_CIDR: fmt.Sprintf("%s", deriveSubnetCIDR(vpcCIDR, i)),
			}
			if subnetCount > 1 && len(zones) > 0 {
				subnet.Zone = zones[i%len(zones)]
			}
			subnets = append(subnets, subnet)
		}
		return cloudmodel.VNetReq{
			ConnectionName: connectionName,
			Name:           "k8s-vpc",
			CidrBlock:      vpcCIDR,
			SubnetInfoList: subnets,
			Description:    "VPC for migrated K8s cluster",
		}
	}

	// Try dynamic recommendation when source network data is available.
	if len(srcInfra.Network.IPv4Networks.CidrBlocks) > 0 || len(srcInfra.Network.IPv4Networks.DefaultGateways) > 0 {
		if recommended, err := RecommendVNet(provider, region, srcInfra); err == nil && len(recommended) > 0 {
			vpcCIDR := recommended[0].CidrBlock
			// Check for CIDR overlap with K8s internal networks.
			if srcInfra.K8sCluster != nil {
				if cidrsOverlap(vpcCIDR, srcInfra.K8sCluster.PodCIDR) || cidrsOverlap(vpcCIDR, srcInfra.K8sCluster.ServiceCIDR) {
					log.Warn().
						Str("vpcCIDR", vpcCIDR).
						Str("podCIDR", srcInfra.K8sCluster.PodCIDR).
						Str("serviceCIDR", srcInfra.K8sCluster.ServiceCIDR).
						Msg("Recommended VNet CIDR overlaps with K8s internal CIDRs; falling back to default")
				} else {
					log.Info().Str("vpcCIDR", vpcCIDR).Msg("Using dynamically recommended VNet CIDR for K8s cluster")
					// Use only the CIDR from RecommendVNet; generate K8s-appropriate subnets internally.
					return buildVNet(vpcCIDR)
				}
			}
		} else if err != nil {
			log.Warn().Err(err).Msg("VNet recommendation failed; using hardcoded fallback for K8s cluster")
		}
	}

	// Fallback: hardcoded CIDR with fixed subnet naming.
	log.Debug().Str("cidr", fallbackCIDR).Msg("Using fallback CIDR for K8s VNet")
	return buildVNet(fallbackCIDR)
}

// deriveSubnetCIDR returns a /24 subnet CIDR from a VPC CIDR by replacing the third octet.
func deriveSubnetCIDR(vpcCIDR string, index int) string {
	// Simple heuristic: replace last two octets with index-based values.
	// Works for standard /16 VPC CIDRs (e.g. 10.10.0.0/16 → 10.10.1.0/24).
	parts := strings.Split(vpcCIDR, "/")
	if len(parts) != 2 {
		return fmt.Sprintf("10.10.%d.0/24", index+1)
	}
	octets := strings.Split(parts[0], ".")
	if len(octets) < 4 {
		return fmt.Sprintf("10.10.%d.0/24", index+1)
	}
	return fmt.Sprintf("%s.%s.%d.0/24", octets[0], octets[1], index+1)
}

// cidrsOverlap returns true when neither CIDR is empty and they share any address space.
func cidrsOverlap(cidr1, cidr2 string) bool {
	if cidr1 == "" || cidr2 == "" {
		return false
	}
	_, net1, err1 := net.ParseCIDR(cidr1)
	_, net2, err2 := net.ParseCIDR(cidr2)
	if err1 != nil || err2 != nil {
		return false
	}
	return net1.Contains(net2.IP) || net2.Contains(net1.IP)
}

// defaultRequiredSubnetCount is the per-CSP fallback used when the K8s asset does not say how many
// subnets a cluster needs. AWS EKS requires two, in distinct AZs; the other CSPs need one.
func defaultRequiredSubnetCount(provider string) int {
	if strings.EqualFold(provider, "aws") {
		return 2
	}
	return 1
}

// getRegionZones returns the availability zones of the region, from Tumblebug. Falls back to the
// conventional "<region>a"/"<region>b" names on lookup failure.
//
// This stays a lookup of its own rather than joining the K8s asset read: it hits a different
// endpoint, and only multi-subnet CSPs need it at all.
func getRegionZones(provider, region string) []string {
	info, err := tbclient.NewSession().ReadRegionInfo(provider, region)
	if err != nil || len(info.Zones) == 0 {
		fallback := []string{region + "a", region + "b"}
		log.Warn().Err(err).Str("provider", provider).Str("region", region).
			Msg("Failed to fetch region zones; using conventional <region>+letter fallback")
		return fallback
	}
	return info.Zones
}

// targetProfile is what the target cloud declares about K8s clusters, read before any per-worker
// work begins. Gathering it in one place is what keeps the rest of the recommendation free of
// scattered lookups — and what lets the functions taking a profile be tested without Tumblebug.
//
// provider and region ride along so those functions need no extra parameters.
type targetProfile struct {
	provider string
	region   string

	nodeImageDesignation bool
	nodeImages           []tbmodel.K8sClusterNodeImageDetailAvailable

	nodeGroupNamingRule string
	requiredSubnetCount int
	regionZones         []string // only fetched when requiredSubnetCount > 1
	rootDiskType        string
}

// getTargetProfile reads the target cloud's K8s declarations in one place.
//
// It does not fail. Each field falls back to what the individual lookups used before, so a
// Tumblebug hiccup degrades the recommendation instead of aborting it.
func getTargetProfile(provider, region string) targetProfile {
	p := targetProfile{provider: provider, region: region, rootDiskType: "default"}

	profile, err := tbclient.NewSession().GetK8sClusterProfile(provider, region)
	if err != nil {
		log.Warn().Err(err).Str("provider", provider).Str("region", region).
			Msg("Failed to read K8s cluster info; using per-field defaults")
		p.requiredSubnetCount = defaultRequiredSubnetCount(provider)
	} else {
		p.nodeImageDesignation = profile.NodeImageDesignation
		p.nodeImages = profile.NodeImages
		p.nodeGroupNamingRule = profile.NodeGroupNamingRule
		p.requiredSubnetCount = profile.RequiredSubnetCount
		p.rootDiskType = profile.RootDiskType
		if p.rootDiskType == "" {
			p.rootDiskType = "default"
		}
		if p.requiredSubnetCount < 1 {
			log.Warn().Str("provider", provider).Int("fallback", defaultRequiredSubnetCount(provider)).
				Msg("K8s asset declares no required subnet count; using per-provider fallback")
			p.requiredSubnetCount = defaultRequiredSubnetCount(provider)
		}
	}

	if p.requiredSubnetCount > 1 {
		p.regionZones = getRegionZones(provider, region)
	}
	return p
}

// buildK8sSshKeyReq builds a SshKeyReq for worker node access.
func buildK8sSshKeyReq(connectionName string) cloudmodel.SshKeyReq {
	return cloudmodel.SshKeyReq{
		ConnectionName: connectionName,
		Name:           "k8s-sshkey",
		Description:    "SSH key for K8s worker nodes",
	}
}

// buildK8sSecurityGroupReqList builds SecurityGroup rules for K8s cluster access.
//
// When worker nodes carry FirewallTable data (collected via SSH by cm-honeybee), it calls
// RecommendSecurityGroups to derive rules from the source firewall configuration and then
// augments the result with K8s-specific ports (kubelet, NodePort range). If FirewallTable
// is absent (e.g., honeybee mergeK8sNodes bug), it falls back to a minimal fixed rule set.
//
// Note: VNetId is intentionally empty here; migration logic fills it after VNet creation.
func buildK8sSecurityGroupReqList(connectionName, provider, region, vpcCIDR string, workers []onpremmodel.NodeProperty, k8sCluster *onpremmodel.K8sClusterProperty) []cloudmodel.SecurityGroupReq {
	const k8sKubeletPort = "10250"
	const k8sApiServerPort = "6443"
	nodePortRange := "30000-32767"
	if k8sCluster != nil && k8sCluster.NodePortRange != "" {
		nodePortRange = k8sCluster.NodePortRange
	}

	// Standard K8s access rules: API server, Kubelet, NodePort, SSH, and intra-cluster communication.
	requiredRules := []cloudmodel.FirewallRuleReq{
		{Ports: k8sApiServerPort, Protocol: "TCP", Direction: "inbound", CIDR: "0.0.0.0/0"},
		{Ports: "443", Protocol: "TCP", Direction: "inbound", CIDR: "0.0.0.0/0"},
		{Ports: "22", Protocol: "TCP", Direction: "inbound", CIDR: "0.0.0.0/0"},
		{Ports: k8sKubeletPort, Protocol: "TCP", Direction: "inbound", CIDR: "0.0.0.0/0"},
		{Ports: nodePortRange, Protocol: "TCP", Direction: "inbound", CIDR: "0.0.0.0/0"},
		{Ports: "1-65535", Protocol: "TCP", Direction: "inbound", CIDR: vpcCIDR},
		{Ports: "1-65535", Protocol: "UDP", Direction: "inbound", CIDR: vpcCIDR},
		{Ports: "1-65535", Protocol: "TCP", Direction: "inbound", CIDR: "172.16.0.0/12"},
		{Ports: "1-65535", Protocol: "UDP", Direction: "inbound", CIDR: "172.16.0.0/12"},
		{Ports: "1-65535", Protocol: "TCP", Direction: "outbound", CIDR: "0.0.0.0/0"},
		{Ports: "1-65535", Protocol: "UDP", Direction: "outbound", CIDR: "0.0.0.0/0"},
	}
	if k8sCluster != nil && k8sCluster.PodCIDR != "" && k8sCluster.PodCIDR != vpcCIDR {
		requiredRules = append(requiredRules,
			cloudmodel.FirewallRuleReq{Ports: "1-65535", Protocol: "TCP", Direction: "inbound", CIDR: k8sCluster.PodCIDR},
			cloudmodel.FirewallRuleReq{Ports: "1-65535", Protocol: "UDP", Direction: "inbound", CIDR: k8sCluster.PodCIDR},
		)
	}

	// Try to derive rules from source worker firewall tables.
	var workerWithFirewall []onpremmodel.NodeProperty
	for _, w := range workers {
		if len(w.FirewallTable) > 0 {
			workerWithFirewall = append(workerWithFirewall, w)
		}
	}

	if len(workerWithFirewall) > 0 {
		sgList, err := RecommendSecurityGroups(provider, region, workerWithFirewall)
		if err == nil && len(sgList.TargetSecurityGroupList) > 0 {
			// Take the first recommended SG and augment with K8s-specific ports.
			sg := sgList.TargetSecurityGroupList[0].TargetSecurityGroup
			sg.ConnectionName = connectionName
			sg.Name = "k8s-sg"

			existingRules := make(map[string]bool)
			if sg.FirewallRules != nil {
				for _, r := range *sg.FirewallRules {
					existingRules[r.Ports+"|"+r.Protocol+"|"+r.Direction+"|"+r.CIDR] = true
				}
			}

			// Add K8s-specific ports when not already present with matching CIDR from source rules.
			var k8sExtra []cloudmodel.FirewallRuleReq
			for _, candidate := range requiredRules {
				key := candidate.Ports + "|" + candidate.Protocol + "|" + candidate.Direction + "|" + candidate.CIDR
				if !existingRules[key] {
					k8sExtra = append(k8sExtra, candidate)
					existingRules[key] = true
				}
			}

			if sg.FirewallRules != nil {
				*sg.FirewallRules = append(*sg.FirewallRules, k8sExtra...)
			} else {
				sg.FirewallRules = &k8sExtra
			}
			log.Info().Int("workers", len(workerWithFirewall)).Msg("Using dynamically recommended SecurityGroup for K8s cluster")
			return []cloudmodel.SecurityGroupReq{sg}
		}
		log.Warn().Err(err).Msg("SecurityGroup recommendation from FirewallTable failed; using fallback rules")
	} else {
		log.Debug().Msg("No worker FirewallTable data available; using fallback K8s SecurityGroup rules")
	}

	// Fallback: minimal fixed rules covering standard K8s access patterns.
	rules := &requiredRules
	return []cloudmodel.SecurityGroupReq{{
		ConnectionName: connectionName,
		Name:           "k8s-sg",
		FirewallRules:  rules,
		Description:    "Security group for migrated K8s cluster",
	}}
}

// cspMaxNodeGroupNameLen holds CSP-specific maximum node group name lengths that are NOT
// encoded in Tumblebug's nodeGroupNamingRule regex. Azure AKS caps agent pool names at
// 12 chars, but its rule ("^[a-z][a-z0-9]*$") omits that bound.
var cspMaxNodeGroupNameLen = map[string]int{
	"azure": 12,
}

// cspRequiresFixedNodeGroupSize lists CSPs that reject MaxNodeSize=0 for a non-autoscaling
// node group and therefore need a fixed-size group pinned at desiredSize (min=max=desired):
//   - AWS EKS: "The MaxNodeSize value must be greater than or equal to 1"
//   - Tencent TKE: "MaxNodeSize cannot be smaller than 1"
//   - IBM IKS: "MaxNodeSize of Node Group ... cannot be smaller than 1"
//
// CSPs not listed here keep min/max at 0 (unset), which Azure AKS requires when auto-scaling
// is off. Verified per CSP via live migration; extend as new CSPs are tested.
var cspRequiresFixedNodeGroupSize = map[string]bool{
	"aws":     true,
	"tencent": true,
	"ibm":     true,
}

// cspRequiresAutoScalingOn lists CSPs where managed node groups are always ASG-backed
// and reject OnAutoScaling=false (e.g. AWS EKS, cb-spider issue #1496).
var cspRequiresAutoScalingOn = map[string]bool{
	"aws": true,
}

// defaultMaxNodeGroupNameLen is a conservative upper bound applied when a CSP has no
// explicit cap above. It stays within the smallest common cloud limit (e.g. GKE node pool
// names allow 40 chars) to remain safe across providers.
const defaultMaxNodeGroupNameLen = 40

// buildNodeGroupName assembles a node group name from base + index, normalizes it to satisfy
// common CSP constraints (lowercase alphanumeric, letter-initial), enforces a safe length
// cap, and validates it against the CSP naming rule fetched from Tumblebug. It never fails:
// on any error it falls back to the normalized default, keeping recommendation resilient to
// Tumblebug hiccups.
//
// The source model carries no node group name to preserve, so names are synthesized. The index
// is what keeps them distinct when heterogeneous workers resolve to several node groups.
func buildNodeGroupName(p targetProfile, base string, index int) string {
	safeName := normalizeNodeGroupName(base, index, maxNodeGroupNameLen(p.provider))

	if p.nodeGroupNamingRule == "" {
		return safeName // CSP defines no naming rule, or it could not be read
	}

	re, err := regexp.Compile(p.nodeGroupNamingRule)
	if err != nil {
		log.Warn().Err(err).Str("rule", p.nodeGroupNamingRule).Str("name", safeName).
			Msg("Invalid node group naming rule regex from Tumblebug; using normalized default")
		return safeName
	}
	if !re.MatchString(safeName) {
		log.Warn().Str("rule", p.nodeGroupNamingRule).Str("name", safeName).
			Msg("Normalized node group name does not match CSP rule; using it as best effort")
	}
	return safeName
}

// maxNodeGroupNameLen returns the safe maximum name length for the provider.
func maxNodeGroupNameLen(provider string) int {
	if n, ok := cspMaxNodeGroupNameLen[strings.ToLower(provider)]; ok {
		return n
	}
	return defaultMaxNodeGroupNameLen
}

// normalizeNodeGroupName lowercases base, strips characters outside [a-z0-9], guarantees a
// letter-initial name, appends the index suffix, and truncates to maxLen (keeping the suffix).
func normalizeNodeGroupName(base string, index, maxLen int) string {
	var b strings.Builder
	for _, r := range strings.ToLower(base) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	// Ensure the name starts with a letter; CSP rules require a letter-initial name.
	cleaned := strings.TrimLeft(b.String(), "0123456789")
	if cleaned == "" {
		cleaned = "workers"
	}

	suffix := strconv.Itoa(index)
	// Truncate the base so that base+suffix fits within maxLen.
	if maxLen > 0 && len(cleaned)+len(suffix) > maxLen {
		keep := maxLen - len(suffix)
		if keep < 1 {
			keep = 1
		}
		cleaned = cleaned[:keep]
	}
	return cleaned + suffix
}

// buildK8sNodeGroupReq builds a K8sNodeGroupReq from one merged node group.
// Note: SshKeyId is intentionally empty here; migration logic fills it after SshKey creation.
//
// Min/MaxNodeSize handling differs per CSP even when auto-scaling is disabled:
//   - Azure AKS rejects MinNodeSize when auto-scaling is off ("If MinNodeSize is
//     specified, OnAutoScaling must be enabled"), so both are left at 0.
//   - AWS EKS requires MaxNodeSize >= 1 unconditionally ("The MaxNodeSize value must
//     be greater than or equal to 1"), so a fixed-size group uses Max = desiredSize.
func buildK8sNodeGroupReq(profile targetProfile, name string, g nodeGroupAccum) cloudmodel.K8sNodeGroupReq {
	provider := profile.provider
	desiredSize := len(g.nodes)
	if desiredSize == 0 {
		desiredSize = 1
	}

	// Use the largest root disk in the group so no node is under-provisioned.
	rootDiskSize := 0
	for _, w := range g.nodes {
		if int(w.RootDisk.TotalSize) > rootDiskSize {
			rootDiskSize = int(w.RootDisk.TotalSize)
		}
	}

	// Apply CSP-specific node group sizing rules when autoscaling is off.
	minNodeSize, maxNodeSize := 0, 0
	if cspRequiresFixedNodeGroupSize[strings.ToLower(provider)] {
		minNodeSize, maxNodeSize = desiredSize, desiredSize
	}

	rootDiskType := profile.rootDiskType
	if rootDiskType == "" {
		rootDiskType = "default"
	}

	description := fmt.Sprintf("Worker node group migrated from on-premise (%d node(s))", desiredSize)
	if len(g.notes) > 0 {
		description += " " + strings.Join(g.notes, " ")
	}

	onAutoScaling := "false"
	if cspRequiresAutoScalingOn[strings.ToLower(provider)] {
		onAutoScaling = "true"
	}

	return cloudmodel.K8sNodeGroupReq{
		Name:         name,
		ImageId:      g.imageId,
		SpecId:       g.specId(),
		RootDiskType: rootDiskType,
		RootDiskSize: rootDiskSize,
		// SshKeyId filled by migration logic
		OnAutoScaling:   onAutoScaling,
		DesiredNodeSize: desiredSize,
		MinNodeSize:     minNodeSize,
		MaxNodeSize:     maxNodeSize,
		Description:     description,
	}
}
