package recommendation

import (
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"

	"github.com/rs/zerolog/log"
)

const (
	// MinMultiInfraTargets is the minimum number of CSP/region pairs a multi-target
	// recommendation request must include (a single pair is just /recommendation/infra).
	MinMultiInfraTargets = 2

	// MaxMultiInfraTargets caps the number of CSP/region pairs a multi-target recommendation
	// request may include. Set to the project's current scope of 10 supported CSPs.
	MaxMultiInfraTargets = 10

	// multiInfraCandidatesPerTarget is intentionally fixed at 1: multi-target recommendation
	// is for cross-CSP comparison (breadth), not per-CSP alternatives (depth, still served by
	// the single-target APIs). This also bounds the per-target preflight-check cost, which
	// matters because multi-target calls already multiply Tumblebug pacer load by target count.
	multiInfraCandidatesPerTarget = 1
)

// singleTargetRecommender is the shape shared by RecommendVmInfraCandidates and
// RecommendInfraWithNlbCandidates. Letting the multi-target orchestrator take either as a
// parameter avoids duplicating validation, sequencing, or sentinel-item logic per variant.
type singleTargetRecommender func(desiredCsp, desiredRegion string, srcInfra onpremmodel.OnpremInfra, limit int, minMatchRate float64) ([]cloudmodel.RecommendedInfra, error)

// RecommendMultiInfraCandidates recommends the single best-match infrastructure candidate for
// each target CSP/region pair, for cross-CSP comparison. Composes RecommendVmInfraCandidates;
// no matching/ranking logic is duplicated here.
func RecommendMultiInfraCandidates(pairs []cloudmodel.CloudProperty, srcInfra onpremmodel.OnpremInfra, minMatchRate float64) ([]cloudmodel.RecommendedInfra, error) {
	return recommendPerTarget(pairs, srcInfra, minMatchRate, RecommendVmInfraCandidates)
}

// RecommendMultiInfraWithNlbCandidates is the NLB-aware counterpart of
// RecommendMultiInfraCandidates, composing RecommendInfraWithNlbCandidates per target.
func RecommendMultiInfraWithNlbCandidates(pairs []cloudmodel.CloudProperty, srcInfra onpremmodel.OnpremInfra, minMatchRate float64) ([]cloudmodel.RecommendedInfra, error) {
	return recommendPerTarget(pairs, srcInfra, minMatchRate, RecommendInfraWithNlbCandidates)
}

// recommendPerTarget runs recommend once per target pair, sequentially. Tumblebug calls share
// a single process-wide pacer, so parallelizing targets would not reduce latency and would
// only risk 429s. Every pair yields exactly one result item — the top candidate on success, or
// a sentinel "failed"/"nothing-to-recommend" item otherwise — so len(result) == len(pairs)
// always holds, giving callers a predictable, index-free way to match results back to targets.
func recommendPerTarget(pairs []cloudmodel.CloudProperty, srcInfra onpremmodel.OnpremInfra, minMatchRate float64, recommend singleTargetRecommender) ([]cloudmodel.RecommendedInfra, error) {
	results := make([]cloudmodel.RecommendedInfra, 0, len(pairs))

	for _, pair := range pairs {
		csp := strings.ToLower(pair.Csp)
		region := strings.ToLower(pair.Region)

		if ok, err := IsValidCspAndRegion(csp, region); !ok {
			log.Warn().Err(err).Str("csp", csp).Str("region", region).Msg("skipping invalid target in multi-target recommendation")
			results = append(results, failedTargetResult(csp, region, err))
			continue
		}

		candidates, err := recommend(csp, region, srcInfra, multiInfraCandidatesPerTarget, minMatchRate)
		if err != nil {
			log.Warn().Err(err).Str("csp", csp).Str("region", region).Msg("recommendation failed for target")
			results = append(results, failedTargetResult(csp, region, err))
			continue
		}
		if len(candidates) == 0 {
			// Matches the status vocabulary of RecommendVmInfraCandidates/RecommendInfraWithNlbCandidates
			// ("highly-matched" / "partially-matched" / "nothing-to-recommend"), not the older
			// RecommendationStatus enum ("none"/"partial"/"ok") used by the dynamic-list recommender.
			results = append(results, cloudmodel.RecommendedInfra{
				Status:      "nothing-to-recommend",
				Description: "No compatible infrastructure could be recommended for this target.",
				TargetCloud: cloudmodel.CloudProperty{Csp: csp, Region: region},
			})
			continue
		}

		results = append(results, candidates[0])
	}

	return results, nil
}

// failedTargetResult builds a sentinel RecommendedInfra for a target that could not be
// processed (invalid CSP/region, or an internal recommendation error).
func failedTargetResult(csp, region string, err error) cloudmodel.RecommendedInfra {
	msg := "recommendation failed"
	if err != nil {
		msg = err.Error()
	}
	return cloudmodel.RecommendedInfra{
		Status:      "failed",
		Description: msg,
		TargetCloud: cloudmodel.CloudProperty{Csp: csp, Region: region},
	}
}
