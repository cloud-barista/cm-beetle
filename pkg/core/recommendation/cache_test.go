package recommendation

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockTumblebugCounters struct {
	regionCalls        atomic.Int64
	recommendSpecCalls atomic.Int64
	searchImageCalls   atomic.Int64
	reviewPairCalls    atomic.Int64
	failNextSpec       atomic.Bool
}

func newMockTumblebugHandler(counters *mockTumblebugCounters) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/provider/"):
			counters.regionCalls.Add(1)
			parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
			regionID := "ap-northeast-2"
			if len(parts) >= 4 {
				regionID = parts[3]
			}
			_ = json.NewEncoder(w).Encode(tbmodel.RegionDetail{
				RegionId:   regionID,
				RegionName: regionID,
				Zones:      []string{regionID + "a", regionID + "c"},
			})

		case r.Method == http.MethodPost && r.URL.Path == "/recommendSpec":
			counters.recommendSpecCalls.Add(1)
			if counters.failNextSpec.Swap(false) {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"message":"transient error"}`))
				return
			}
			var req tbmodel.RecommendSpecReq
			_ = json.NewDecoder(r.Body).Decode(&req)
			provider := "aws"
			region := "ap-northeast-2"
			arch := "x86_64"
			isGpu := false
			var vcpu uint16 = 4
			var mem float32 = 8
			for _, p := range req.Filter.Policy {
				switch p.Metric {
				case "providerName":
					if len(p.Condition) > 0 {
						provider = p.Condition[0].Operand
					}
				case "regionName":
					if len(p.Condition) > 0 {
						region = p.Condition[0].Operand
					}
				case "architecture":
					if len(p.Condition) > 0 {
						arch = p.Condition[0].Operand
					}
				case "acceleratorCount":
					isGpu = true
				case "vCPU":
					if len(p.Condition) > 0 {
						var parsed int
						_, _ = fmt.Sscanf(p.Condition[0].Operand, "%d", &parsed)
						if parsed > 0 {
							vcpu = uint16(parsed)
						}
					}
				case "memoryGiB":
					if len(p.Condition) > 0 {
						var parsed int
						_, _ = fmt.Sscanf(p.Condition[0].Operand, "%d", &parsed)
						if parsed > 0 {
							mem = float32(parsed)
						}
					}
				}
			}

			var specs []tbmodel.SpecInfo
			if isGpu {
				specs = []tbmodel.SpecInfo{
					{
						Id: fmt.Sprintf("%s+%s+g6.xlarge", provider, region), CspSpecName: "g6.xlarge",
						ProviderName: provider, RegionName: region, Architecture: arch,
						VCPU: 4, MemoryGiB: 16, AcceleratorCount: 1, AcceleratorMemoryGB: 24, AcceleratorModel: "NVIDIA L4",
						CostPerHour: 0.80, AvailabilityStatus: tbmodel.AvailabilityAvailable,
						AvailableZones: []string{region + "a"},
					},
					{
						Id: fmt.Sprintf("%s+%s+g5.xlarge", provider, region), CspSpecName: "g5.xlarge",
						ProviderName: provider, RegionName: region, Architecture: arch,
						VCPU: 4, MemoryGiB: 16, AcceleratorCount: 1, AcceleratorMemoryGB: 24, AcceleratorModel: "NVIDIA A10G",
						CostPerHour: 1.00, AvailabilityStatus: tbmodel.AvailabilityAvailable,
						AvailableZones: []string{region + "a"},
					},
					{
						Id: fmt.Sprintf("%s+%s+g4dn.xlarge", provider, region), CspSpecName: "g4dn.xlarge",
						ProviderName: provider, RegionName: region, Architecture: arch,
						VCPU: 4, MemoryGiB: 16, AcceleratorCount: 1, AcceleratorMemoryGB: 16, AcceleratorModel: "NVIDIA T4",
						CostPerHour: 0.52, AvailabilityStatus: tbmodel.AvailabilityAvailable,
						AvailableZones: []string{region + "a"},
					},
				}
			} else {
				specs = []tbmodel.SpecInfo{
					{
						Id: fmt.Sprintf("%s+%s+m6i-%dc-%dg", provider, region, vcpu, int(mem)), CspSpecName: fmt.Sprintf("m6i.%d.%d", vcpu, int(mem)),
						ProviderName: provider, RegionName: region, Architecture: arch,
						VCPU: vcpu, MemoryGiB: mem, CostPerHour: 0.12,
						AvailabilityStatus: tbmodel.AvailabilityAvailable,
						AvailableZones:     []string{region + "a", region + "c"},
					},
					{
						Id: fmt.Sprintf("%s+%s+m6a-%dc-%dg", provider, region, vcpu, int(mem)), CspSpecName: fmt.Sprintf("m6a.%d.%d", vcpu, int(mem)),
						ProviderName: provider, RegionName: region, Architecture: arch,
						VCPU: vcpu, MemoryGiB: mem, CostPerHour: 0.10,
						AvailabilityStatus: tbmodel.AvailabilityAvailable,
						AvailableZones:     []string{region + "a", region + "c"},
					},
					{
						Id: fmt.Sprintf("%s+%s+c6i-%dc-%dg", provider, region, vcpu, int(mem)), CspSpecName: fmt.Sprintf("c6i.%d.%d", vcpu, int(mem)),
						ProviderName: provider, RegionName: region, Architecture: arch,
						VCPU: vcpu + 2, MemoryGiB: mem + 4, CostPerHour: 0.18,
						AvailabilityStatus: tbmodel.AvailabilityAvailable,
						AvailableZones:     []string{region + "a"},
					},
				}
			}
			_ = json.NewEncoder(w).Encode(specs)

		case r.Method == http.MethodPost && r.URL.Path == "/ns/system/resources/searchImage":
			counters.searchImageCalls.Add(1)
			var req tbmodel.SearchImageRequest
			_ = json.NewDecoder(r.Body).Decode(&req)
			provider := req.ProviderName
			if provider == "" {
				provider = "aws"
			}
			region := req.RegionName
			if region == "" {
				region = "ap-northeast-2"
			}
			arch := req.OSArchitecture
			if arch == "" {
				arch = tbmodel.OSArchitecture("x86_64")
			}
			resp := tbmodel.SearchImageResponse{
				ImageCount: 1,
				ImageList: []tbmodel.ImageInfo{
					{
						Id:             fmt.Sprintf("%s+%s+ubuntu-22.04-%s", provider, region, req.OSType),
						Name:           "ubuntu-22.04",
						ProviderName:   provider,
						CspImageName:   fmt.Sprintf("ami-ubuntu-2204-%s", strings.ReplaceAll(req.OSType, " ", "-")),
						OSType:         req.OSType,
						OSArchitecture: arch,
						OSDistribution: "Ubuntu 22.04 LTS",
						OSDiskType:     "SSD",
					},
				},
			}
			_ = json.NewEncoder(w).Encode(resp)

		case r.Method == http.MethodPost && r.URL.Path == "/specImagePairReview":
			counters.reviewPairCalls.Add(1)
			var req tbmodel.SpecImagePairReviewReq
			_ = json.NewDecoder(r.Body).Decode(&req)
			_ = json.NewEncoder(w).Encode(tbmodel.SpecImagePairReviewResult{
				IsValid: true,
				Status:  "OK",
				SpecId:  req.SpecId,
				ImageId: req.ImageId,
				ImageValidation: tbmodel.ReviewResourceValidation{
					ResourceId:    req.ImageId,
					IsAvailable:   true,
					Status:        "Available",
					CspResourceId: req.ImageId + "-latest",
				},
				SuggestedSystemDisk: "gp3",
			})

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

func setupMockTumblebugServer(t *testing.T) (*mockTumblebugCounters, *httptest.Server) {
	t.Helper()
	prevLevel := zerolog.GlobalLevel()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	t.Cleanup(func() { zerolog.SetGlobalLevel(prevLevel) })

	counters := &mockTumblebugCounters{}
	srv := httptest.NewServer(newMockTumblebugHandler(counters))

	restore := tbclient.SetClient(tbclient.NewClient(tbclient.ApiConfig{
		RestUrl:  srv.URL,
		Username: "default",
		Password: "default",
	}))
	t.Cleanup(func() {
		restore()
		srv.Close()
	})

	return counters, srv
}

func buildSynthetic1000NodeInfra(count int) onpremmodel.OnpremInfra {
	type hwProfile struct {
		cpus   uint32
		memGB  uint64
		diskGB uint64
	}
	// 12 distinct (vCPU, Memory, RootDisk) combinations matching stress test distribution.
	profiles := []hwProfile{
		{2, 4, 30}, {2, 8, 50}, {2, 16, 100}, {2, 32, 100},
		{4, 4, 30}, {4, 8, 50}, {4, 16, 100}, {4, 32, 100},
		{8, 8, 50}, {8, 16, 50}, {8, 32, 100}, {8, 64, 100},
	}

	rolePorts := [][]string{
		{"22", "80", "443"},
		{"22", "8080", "8443"},
		{"22", "3306", "5432"},
		{"22", "9092"},
	}

	numSubnets := (count + 239) / 240
	if numSubnets < 1 {
		numSubnets = 1
	}
	cidrs := make([]string, numSubnets)
	for s := 0; s < numSubnets; s++ {
		cidrs[s] = fmt.Sprintf("10.1.%d.0/24", s+1)
	}
	nodes := make([]onpremmodel.NodeProperty, 0, count)
	for i := 0; i < count; i++ {
		prof := profiles[i%len(profiles)]
		ports := rolePorts[i%len(rolePorts)]
		fw := make([]onpremmodel.FirewallRuleProperty, 0, len(ports))
		for _, p := range ports {
			fw = append(fw, onpremmodel.FirewallRuleProperty{
				Action:    "allow",
				Direction: "inbound",
				Protocol:  "tcp",
				SrcCIDR:   "0.0.0.0/0",
				DstPorts:  p,
			})
		}

		subnetIdx := i / 240
		if subnetIdx >= len(cidrs) {
			subnetIdx = len(cidrs) - 1
		}
		hostOctet := (i % 240) + 10
		nodes = append(nodes, onpremmodel.NodeProperty{
			MachineId: fmt.Sprintf("00000000-0000-0000-0000-%012d", i+1),
			Hostname:  fmt.Sprintf("srv-%04d", i+1),
			CPU: onpremmodel.CpuProperty{
				Cpus:         prof.cpus,
				Cores:        prof.cpus,
				Threads:      1,
				Architecture: "x86_64",
				Vendor:       "GenuineIntel",
			},
			Memory: onpremmodel.MemoryProperty{
				TotalSize: prof.memGB,
			},
			RootDisk: onpremmodel.DiskProperty{
				Type:      "SSD",
				TotalSize: prof.diskGB,
			},
			OS: onpremmodel.OsProperty{
				ID:              "ubuntu",
				VersionID:       "22.04",
				VersionCodename: "jammy",
				Name:            "Ubuntu",
				Version:         "22.04.5 LTS",
			},
			Interfaces: []onpremmodel.NetworkInterfaceProperty{
				{
					Name:           "eth0",
					IPv4CidrBlocks: []string{fmt.Sprintf("10.1.%d.%d/24", subnetIdx+1, hostOctet)},
				},
			},
			FirewallTable: fw,
		})
	}

	return onpremmodel.OnpremInfra{
		Network: onpremmodel.NetworkProperty{
			IPv4Networks: onpremmodel.NetworkDetail{
				CidrBlocks: cidrs,
			},
		},
		Nodes: nodes,
	}
}

func TestRecommendInfraCandidates_RequestScopedCache_1000Nodes(t *testing.T) {
	counters, _ := setupMockTumblebugServer(t)

	reqID := fmt.Sprintf("req-cache-1000-%d", time.Now().UnixNano())
	t.Cleanup(func() { common.DiscardSubRequests(reqID) })
	ctx := tbclient.WithRequestCache(common.WithRequestID(context.Background(), reqID))

	srcInfra := buildSynthetic1000NodeInfra(1000)

	start := time.Now()
	candidates, err := RecommendInfraCandidates(ctx, "aws", "ap-northeast-2", srcInfra, 3, 85.0)
	elapsed := time.Since(start)

	require.NoError(t, err)
	require.Len(t, candidates, 3)

	// Verify topology integrity across 1,000 nodes (5 subnets, 4 deduplicated SGs).
	for _, cand := range candidates {
		assert.Len(t, cand.TargetInfra.NodeGroups, 1000)
		assert.Len(t, cand.TargetVNet.SubnetInfoList, 5)
		assert.Len(t, cand.TargetSecurityGroupList, 4)
		for _, ng := range cand.TargetInfra.NodeGroups {
			assert.NotEmpty(t, ng.SpecId)
			assert.NotEmpty(t, ng.ImageId)
			assert.Equal(t, "gp3", ng.RootDiskType)
			assert.NotEmpty(t, ng.CspImageName)
		}
	}

	// Verify API calls dropped from ~5,000 POST calls (+1,001 GET calls) to <= 60 total.
	regionCalls := counters.regionCalls.Load()
	specCalls := counters.recommendSpecCalls.Load()
	imageCalls := counters.searchImageCalls.Load()
	reviewCalls := counters.reviewPairCalls.Load()
	totalPostCalls := specCalls + imageCalls + reviewCalls
	stats, hasStats := tbclient.GetRequestCacheStats(ctx)

	t.Logf("1,000-node cached run completed in %v: region=%d, spec=%d, image=%d, review=%d, totalPost=%d, stats=%+v",
		elapsed, regionCalls, specCalls, imageCalls, reviewCalls, totalPostCalls, stats)

	assert.True(t, hasStats)
	assert.Greater(t, stats.Hits, 5000)
	assert.Equal(t, int64(1), regionCalls, "region validation should be called once across all 1,000 nodes")
	assert.LessOrEqual(t, specCalls, int64(12), "unique spec profiles should cap recommendSpec calls")
	assert.Equal(t, int64(1), imageCalls, "identical OS profile should call searchImage once")
	assert.LessOrEqual(t, reviewCalls, int64(36), "unique spec-image pairs across 3 candidates should cap preflight calls")
	assert.LessOrEqual(t, totalPostCalls, int64(60), "total downstream POST calls for 1,000 nodes must stay within 60")

	subs, summary := common.FinalizeSubRequests(reqID)
	require.NotNil(t, summary)
	assert.Equal(t, int(totalPostCalls), summary.Total)
	assert.Equal(t, int(totalPostCalls), summary.Success)
	assert.Equal(t, 0, summary.Error)
	assert.Len(t, subs, int(totalPostCalls))
}

func TestRecommendInfraCandidates_CacheMatchesUncachedOutput(t *testing.T) {
	setupMockTumblebugServer(t)

	// 24 nodes covering 12 hardware profiles x 2 passes.
	srcInfra := buildSynthetic1000NodeInfra(24)

	cachedCandidates, err := RecommendInfraCandidates(tbclient.WithRequestCache(context.Background()), "aws", "ap-northeast-2", srcInfra, 3, 85.0)
	require.NoError(t, err)
	require.Len(t, cachedCandidates, 3)

	// Compare against independent uncached single-node runs per node.
	for nodeIdx, node := range srcInfra.Nodes {
		singleSpecs, _, err := RecommendNodeSpecs(context.Background(), "aws", "ap-northeast-2", node, GetDefaultSpecsLimit())
		require.NoError(t, err)
		singleImages, err := RecommendVmOsImages(context.Background(), "aws", "ap-northeast-2", node, GetDefaultImagesLimit())
		require.NoError(t, err)
		pairs, err := FindCompatibleVmSpecAndImagePairs(singleSpecs, singleImages, "aws")
		require.NoError(t, err)

		for candIdx := range cachedCandidates {
			expectedPair := pairs[candIdx]
			precheck, err := PreflightCheckCspProvisioning(
				context.Background(), expectedPair.Spec.Id, expectedPair.Image.Id, expectedPair.Image.CspImageName, "",
			)
			require.NoError(t, err)

			actualNG := cachedCandidates[candIdx].TargetInfra.NodeGroups[nodeIdx]
			assert.Equal(t, expectedPair.Spec.Id, actualNG.SpecId)
			assert.Equal(t, expectedPair.Image.Id, actualNG.ImageId)
			assert.Equal(t, precheck.SuggestedSystemDisk, actualNG.RootDiskType)
			assert.Equal(t, precheck.ResolvedCspImageName, actualNG.CspImageName)
		}
	}
}

func TestRecommendationCache_KeyDifferentiationAndErrorSafety(t *testing.T) {
	counters, _ := setupMockTumblebugServer(t)
	ctx := tbclient.WithRequestCache(context.Background())

	baseNode := onpremmodel.NodeProperty{
		MachineId: "n1",
		CPU: onpremmodel.CpuProperty{
			Cpus: 2, Cores: 2, Threads: 1, Architecture: "x86_64", Vendor: "GenuineIntel",
		},
		Memory:   onpremmodel.MemoryProperty{TotalSize: 8},
		RootDisk: onpremmodel.DiskProperty{Type: "SSD", TotalSize: 50},
		OS:       onpremmodel.OsProperty{ID: "ubuntu", VersionID: "22.04", VersionCodename: "jammy"},
	}

	// 1. Transient upstream error must not be cached; retry must invoke Tumblebug again and succeed.
	counters.failNextSpec.Store(true)
	_, _, err := RecommendNodeSpecs(ctx, "aws", "ap-northeast-2", baseNode, 10)
	require.Error(t, err)
	assert.Equal(t, int64(1), counters.recommendSpecCalls.Load())

	specsIntel, _, err := RecommendNodeSpecs(ctx, "aws", "ap-northeast-2", baseNode, 10)
	require.NoError(t, err)
	require.NotEmpty(t, specsIntel)
	assert.Equal(t, int64(2), counters.recommendSpecCalls.Load())
	assert.True(t, strings.HasPrefix(specsIntel[0].CspSpecName, "m6i."), "expected Intel spec first, got %s", specsIntel[0].CspSpecName)
	expectedIntelSpec := specsIntel[0].CspSpecName

	// 2. Identical profile with different machineId hits cache, and caller mutation does not poison cache.
	cloneNode := baseNode
	cloneNode.MachineId = "n2"
	specsIntelCached, _, err := RecommendNodeSpecs(ctx, "aws", "ap-northeast-2", cloneNode, 10)
	require.NoError(t, err)
	assert.Equal(t, int64(2), counters.recommendSpecCalls.Load())
	specsIntelCached[0].CspSpecName = "mutated-by-caller"

	specsIntelAgain, _, err := RecommendNodeSpecs(ctx, "aws", "ap-northeast-2", cloneNode, 10)
	require.NoError(t, err)
	assert.Equal(t, expectedIntelSpec, specsIntelAgain[0].CspSpecName)

	// 3. AMD CPU vendor differs from Intel locally without duplicate Tumblebug call (m6a ranks before m6i).
	amdNode := baseNode
	amdNode.MachineId = "n-amd"
	amdNode.CPU.Vendor = "AuthenticAMD"
	specsAmd, _, err := RecommendNodeSpecs(ctx, "aws", "ap-northeast-2", amdNode, 10)
	require.NoError(t, err)
	assert.Equal(t, int64(2), counters.recommendSpecCalls.Load())
	assert.True(t, strings.HasPrefix(specsAmd[0].CspSpecName, "m6a."), "expected AMD spec first, got %s", specsAmd[0].CspSpecName)

	// Verify caller mutation on cached response doesn't corrupt subsequent vendor ranking either.
	specsAmdCached, _, err := RecommendNodeSpecs(ctx, "aws", "ap-northeast-2", amdNode, 10)
	require.NoError(t, err)
	assert.Equal(t, int64(2), counters.recommendSpecCalls.Load())
	assert.True(t, strings.HasPrefix(specsAmdCached[0].CspSpecName, "m6a."))

	// 4. Different OS version and RootDisk.Type differentiate image cache keys.
	_, err = RecommendVmOsImages(ctx, "aws", "ap-northeast-2", baseNode, 20)
	require.NoError(t, err)
	_, err = RecommendVmOsImages(ctx, "aws", "ap-northeast-2", cloneNode, 20)
	require.NoError(t, err)
	assert.Equal(t, int64(1), counters.searchImageCalls.Load())

	debianNode := baseNode
	debianNode.OS = onpremmodel.OsProperty{ID: "debian", VersionID: "12", VersionCodename: "bookworm"}
	_, err = RecommendVmOsImages(ctx, "aws", "ap-northeast-2", debianNode, 20)
	require.NoError(t, err)
	assert.Equal(t, int64(2), counters.searchImageCalls.Load())

	// 5. Different RootDiskType differentiates preflight cache keys.
	_, err = PreflightCheckCspProvisioning(ctx, "spec-1", "img-1", "ami-1", "")
	require.NoError(t, err)
	_, err = PreflightCheckCspProvisioning(ctx, "spec-1", "img-1", "ami-1", "")
	require.NoError(t, err)
	_, err = PreflightCheckCspProvisioning(ctx, "spec-1", "img-1", "ami-1", "gp3")
	require.NoError(t, err)
	assert.Equal(t, int64(2), counters.reviewPairCalls.Load())
}

func TestRecommendationCache_OtherOrchestrators(t *testing.T) {
	t.Run("RecommendInfraWithDefaults deduplicates across specs and nodes", func(t *testing.T) {
		counters, _ := setupMockTumblebugServer(t)
		infra := buildSynthetic1000NodeInfra(10)
		// Force all 10 nodes to share the same profile.
		for i := range infra.Nodes {
			infra.Nodes[i].CPU = infra.Nodes[0].CPU
			infra.Nodes[i].Memory = infra.Nodes[0].Memory
		}

		res, err := RecommendInfraWithDefaults(tbclient.WithRequestCache(context.Background()), "aws", "ap-northeast-2", infra)
		require.NoError(t, err)
		assert.Greater(t, res.Count, 0)
		assert.Equal(t, int64(1), counters.recommendSpecCalls.Load())
		assert.Equal(t, int64(1), counters.searchImageCalls.Load())
	})

	t.Run("RecommendGpuInfraCandidates caches homogeneous GPU nodes", func(t *testing.T) {
		counters, _ := setupMockTumblebugServer(t)
		infra := buildSynthetic1000NodeInfra(10)
		for i := range infra.Nodes {
			infra.Nodes[i].CPU = onpremmodel.CpuProperty{Cpus: 4, Cores: 4, Threads: 1, Architecture: "x86_64"}
			infra.Nodes[i].Memory = onpremmodel.MemoryProperty{TotalSize: 16}
			infra.Nodes[i].GPUCards = []onpremmodel.GpuCardProperty{
				{Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24},
			}
		}

		cands, err := RecommendGpuInfraCandidates(tbclient.WithRequestCache(context.Background()), "aws", "ap-northeast-2", infra, 3, 85.0)
		require.NoError(t, err)
		require.NotEmpty(t, cands)
		assert.Equal(t, int64(1), counters.recommendSpecCalls.Load())
		assert.Equal(t, int64(1), counters.searchImageCalls.Load())
		assert.LessOrEqual(t, counters.reviewPairCalls.Load(), int64(3))
	})

	t.Run("RecommendMultiInfraCandidates isolates cache per target cloud", func(t *testing.T) {
		counters, _ := setupMockTumblebugServer(t)
		infra := buildSynthetic1000NodeInfra(10)
		for i := range infra.Nodes {
			infra.Nodes[i].CPU = infra.Nodes[0].CPU
			infra.Nodes[i].Memory = infra.Nodes[0].Memory
		}

		res, err := RecommendMultiInfraCandidates(
			tbclient.WithRequestCache(context.Background()),
			[]cloudmodel.CloudProperty{
				{Csp: "aws", Region: "ap-northeast-2"},
				{Csp: "gcp", Region: "asia-northeast3"},
			},
			infra,
			85.0,
		)
		require.NoError(t, err)
		require.Len(t, res, 2)
		assert.Equal(t, int64(2), counters.regionCalls.Load())
		assert.Equal(t, int64(2), counters.recommendSpecCalls.Load())
		assert.Equal(t, int64(2), counters.searchImageCalls.Load())
		assert.Equal(t, int64(2), counters.reviewPairCalls.Load())
	})

	t.Run("RecommendInfraWithNlbCandidates caches standalone nodes", func(t *testing.T) {
		counters, _ := setupMockTumblebugServer(t)
		infra := buildSynthetic1000NodeInfra(10)
		for i := range infra.Nodes {
			infra.Nodes[i].CPU = infra.Nodes[0].CPU
			infra.Nodes[i].Memory = infra.Nodes[0].Memory
		}
		infra.NLBs = []onpremmodel.NlbProperty{
			{
				HostMachineId: infra.Nodes[0].MachineId,
				Software:      "haproxy",
				Listener:      onpremmodel.NlbListenerProperty{BindAddress: "*", Port: 80, Protocol: "tcp"},
				Backend: onpremmodel.NlbBackendProperty{
					Name:     "web-pool",
					Protocol: "tcp",
					Servers: []onpremmodel.NlbServerProperty{
						{Name: "s1", IP: "10.1.1.10", Port: 80},
						{Name: "s2", IP: "10.1.1.11", Port: 80},
					},
				},
			},
		}

		cands, err := RecommendInfraWithNlbCandidates(tbclient.WithRequestCache(context.Background()), "aws", "ap-northeast-2", infra, 3, 85.0)
		require.NoError(t, err)
		require.NotEmpty(t, cands)
		assert.Equal(t, int64(1), counters.regionCalls.Load())
		assert.Equal(t, int64(1), counters.recommendSpecCalls.Load())
		assert.Equal(t, int64(1), counters.searchImageCalls.Load())
		assert.LessOrEqual(t, counters.reviewPairCalls.Load(), int64(3))
	})

	t.Run("Identical PreflightCheckCspProvisioning calls share cache regardless of caller", func(t *testing.T) {
		counters, _ := setupMockTumblebugServer(t)
		ctx := tbclient.WithRequestCache(context.Background())

		for i := 0; i < 10; i++ {
			precheck, err := PreflightCheckCspProvisioning(ctx, "aws+ap-northeast-2+t3.medium", "aws+ap-northeast-2+ubuntu-22.04", "", "gp3")
			require.NoError(t, err)
			require.True(t, precheck.IsAvailable)
		}
		assert.Equal(t, int64(1), counters.reviewPairCalls.Load())
	})
}

func readProcessRssMB() float64 {
	data, err := os.ReadFile("/proc/self/status")
	if err != nil {
		return 0
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "VmRSS:") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				kb, err := strconv.ParseFloat(fields[1], 64)
				if err == nil {
					return kb / 1024.0
				}
			}
		}
	}
	return 0
}

func BenchmarkRecommendInfraCandidates_Scale(b *testing.B) {
	prevLevel := zerolog.GlobalLevel()
	prevLogger := log.Logger
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Logger.Output(io.Discard)
	b.Cleanup(func() {
		zerolog.SetGlobalLevel(prevLevel)
		log.Logger = prevLogger
	})

	counters := &mockTumblebugCounters{}
	srv := httptest.NewServer(newMockTumblebugHandler(counters))
	restore := tbclient.SetClient(tbclient.NewClient(tbclient.ApiConfig{
		RestUrl: srv.URL, Username: "default", Password: "default",
	}))
	b.Cleanup(func() {
		restore()
		srv.Close()
	})

	nodeCounts := []int{
		1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000,
	}
	for _, n := range nodeCounts {
		srcInfra := buildSynthetic1000NodeInfra(n)
		reqBytes, _ := json.Marshal(srcInfra)
		b.Run(fmt.Sprintf("nodes=%d", n), func(b *testing.B) {
			var lastStats tbclient.RequestCacheStats
			var respLen int
			var subnetCount int
			var sgCount int
			var heapMB float64
			var rssMB float64
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ctx := tbclient.WithRequestCache(context.Background())
				cands, err := RecommendInfraCandidates(ctx, "aws", "ap-northeast-2", srcInfra, 3, 85.0)
				if err != nil {
					b.Fatalf("RecommendInfraCandidates(N=%d) failed: %v", n, err)
				}
				if i == b.N-1 {
					b.StopTimer()
					lastStats, _ = tbclient.GetRequestCacheStats(ctx)
					respBytes, _ := json.Marshal(cands)
					respLen = len(respBytes)
					if len(cands) > 0 {
						subnetCount = len(cands[0].TargetVNet.SubnetInfoList)
						sgCount = len(cands[0].TargetSecurityGroupList)
					}
					var ms runtime.MemStats
					runtime.ReadMemStats(&ms)
					heapMB = float64(ms.HeapAlloc) / (1024.0 * 1024.0)
					rssMB = readProcessRssMB()
					b.StartTimer()
				}
			}
			b.ReportMetric(float64(lastStats.Misses), "http_misses")
			b.ReportMetric(float64(lastStats.Hits), "cache_hits")
			b.ReportMetric(float64(subnetCount), "subnets")
			b.ReportMetric(float64(sgCount), "sgs")
			b.ReportMetric(float64(len(reqBytes))/(1024.0*1024.0), "req_mb")
			b.ReportMetric(float64(respLen)/(1024.0*1024.0), "resp_mb")
			b.ReportMetric(heapMB, "heap_mb")
			b.ReportMetric(rssMB, "rss_mb")
		})
	}
}
