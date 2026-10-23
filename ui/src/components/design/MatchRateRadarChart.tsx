'use client';

import React, { useMemo } from 'react';
import {
  Radar,
  RadarChart,
  PolarGrid,
  PolarAngleAxis,
  PolarRadiusAxis,
  ResponsiveContainer,
  Legend,
  Tooltip,
} from 'recharts';
import { OnpremInfra, RecommendedInfra } from '../../types/migration';
import { Sparkles } from 'lucide-react';

interface MatchRateRadarChartProps {
  candidates: RecommendedInfra[];
  sourceInfra?: OnpremInfra | null;
}

// Brand-compliant radar series palette (Emerald, Teal, Slate)
const RADAR_PALETTE = [
  { stroke: '#10b981', fill: '#10b981', fillOpacity: 0.28 },
  { stroke: '#14b8a6', fill: '#14b8a6', fillOpacity: 0.22 },
  { stroke: '#64748b', fill: '#64748b', fillOpacity: 0.18 },
  { stroke: '#059669', fill: '#059669', fillOpacity: 0.16 },
  { stroke: '#0d9488', fill: '#0d9488', fillOpacity: 0.14 },
];

// Estimate total monthly cost for a candidate in USD
function estimateMonthlyCost(candidate: RecommendedInfra): number {
  const specMap = new Map<string, any>(
    (candidate.targetSpecList || []).map((s: any) => [s.id, s])
  );
  let hourlyTotal = 0;
  for (const ng of candidate.targetInfra?.nodeGroups || []) {
    const spec = specMap.get(ng.specId);
    const count = ng.nodeGroupSize || 1;
    if (spec && typeof spec.costPerHour === 'number' && spec.costPerHour > 0) {
      hourlyTotal += spec.costPerHour * count;
    } else {
      const vcpu = spec?.vCPU || 4;
      const mem = spec?.memoryGiB || 16;
      hourlyTotal += (vcpu * 0.035 + mem * 0.004) * count;
    }
  }
  return Math.max(10, Math.round(hourlyTotal * 730));
}

// Clamp score into [55, 99] range for clear visual comparison
function clampScore(val: number): number {
  return Math.min(99, Math.max(55, Math.round(val)));
}

export const MatchRateRadarChart: React.FC<MatchRateRadarChartProps> = ({
  candidates,
  sourceInfra,
}) => {
  const { chartData, candidateMeta } = useMemo(() => {
    const srcNodes = sourceInfra?.nodes || [];
    const srcTotalVcpu = srcNodes.reduce(
      (sum, n) => sum + (n.cpu?.cpus || 1) * (n.cpu?.threads || n.cpu?.cores || 2),
      0
    ) || 16;
    const srcTotalMemGiB = srcNodes.reduce(
      (sum, n) => sum + (n.memory?.totalSize || 8),
      0
    ) || 64;
    const srcGpuCount = srcNodes.reduce(
      (sum, n) => sum + (n.gpuCards?.length || 0),
      0
    );

    const costs = candidates.map((c) => estimateMonthlyCost(c));
    const minCost = Math.min(...costs, 1);
    const maxCost = Math.max(...costs, minCost + 1);

    const meta = candidates.slice(0, 5).map((c, idx) => {
      const key = `c${idx + 1}`;
      const csp = (c.targetCloud?.csp || 'cloud').toUpperCase();
      const region = c.targetCloud?.region || 'region';
      const label = `c ${idx + 1} (${csp} / ${region})`;
      const palette = RADAR_PALETTE[idx % RADAR_PALETTE.length];

      const specMap = new Map<string, any>(
        (c.targetSpecList || []).map((s: any) => [s.id, s])
      );
      let targetVcpu = 0;
      let targetMem = 0;
      let targetGpuCapable = 0;
      let diskMatchedCount = 0;
      const ngList = c.targetInfra?.nodeGroups || [];

      for (const ng of ngList) {
        const spec = specMap.get(ng.specId);
        const count = ng.nodeGroupSize || 1;
        targetVcpu += (spec?.vCPU || 4) * count;
        targetMem += (spec?.memoryGiB || 16) * count;
        const specLower = (ng.specId || '').toLowerCase();
        if (
          specLower.includes('g4') ||
          specLower.includes('g5') ||
          specLower.includes('p3') ||
          specLower.includes('p4') ||
          specLower.includes('nc') ||
          specLower.includes('gpu') ||
          specLower.includes('a2-') ||
          specLower.includes('g2-')
        ) {
          targetGpuCapable += count;
        }
        if ((Number(ng.rootDiskSize) || 50) >= 30 && ng.rootDiskType) {
          diskMatchedCount += 1;
        }
      }

      const vcpuRatio = Math.min(targetVcpu / srcTotalVcpu, srcTotalVcpu / Math.max(targetVcpu, 1));
      const memRatio = Math.min(targetMem / srcTotalMemGiB, srcTotalMemGiB / Math.max(targetMem, 1));
      const computeScore = clampScore((vcpuRatio * 0.55 + memRatio * 0.45) * 96 - idx * 2);

      const acceleratorScore =
        srcGpuCount > 0
          ? clampScore(targetGpuCapable >= srcGpuCount ? 95 - idx * 3 : 78 - idx * 3)
          : clampScore(92 - idx * 2);

      const costVal = costs[idx];
      const costScore =
        maxCost === minCost
          ? clampScore(90 - idx * 3)
          : clampScore(96 - ((costVal - minCost) / (maxCost - minCost)) * 24);

      const storageScore = clampScore(
        84 + (diskMatchedCount / Math.max(ngList.length, 1)) * 12 - idx * 2
      );

      const subnetCount = c.targetVNet?.subnetInfoList?.length || 1;
      const sgCount = c.targetSecurityGroupList?.length || 1;
      const networkScore = clampScore(85 + Math.min(10, subnetCount * 3 + sgCount * 2) - idx * 2);

      const regionLower = region.toLowerCase();
      const regionBase =
        regionLower.includes('northeast-2') ||
        regionLower.includes('seoul') ||
        regionLower.includes('korea') ||
        regionLower.includes('kr')
          ? 96
          : regionLower.includes('northeast') || regionLower.includes('asia') || regionLower.includes('east')
          ? 88
          : 82;
      const latencyScore = clampScore(regionBase - idx * 2);

      const avgScore = Math.round(
        (computeScore + acceleratorScore + costScore + storageScore + networkScore + latencyScore) / 6
      );

      return {
        key,
        label,
        palette,
        avgScore,
        costUsd: costVal,
        scores: {
          'Compute Match': computeScore,
          'Accelerator Fit': acceleratorScore,
          'Cost Efficiency': costScore,
          'Storage Performance': storageScore,
          'Network Topology': networkScore,
          'Region Latency': latencyScore,
        },
      };
    });

    const dimensions = [
      'Compute Match',
      'Accelerator Fit',
      'Cost Efficiency',
      'Storage Performance',
      'Network Topology',
      'Region Latency',
    ] as const;

    const rows = dimensions.map((dim) => {
      const row: Record<string, string | number> = { dimension: dim };
      for (const m of meta) {
        row[m.key] = m.scores[dim];
      }
      return row;
    });

    return { chartData: rows, candidateMeta: meta };
  }, [candidates, sourceInfra]);

  if (!candidates || candidates.length === 0) return null;

  return (
    <div className="bg-bg-panel/50 border border-border-main/60 rounded-xl p-5 space-y-4">
      <div className="flex flex-col md:flex-row md:items-center justify-between gap-3 border-b border-border-main/30 pb-3">
        <div>
          <h4 className="text-sm font-extrabold text-text-main flex items-center gap-2">
            <Sparkles className="w-4 h-4 text-emerald-500" />
            <span>6-Dimensional Candidate Fitness Radar</span>
          </h4>
          <p className="text-sm text-text-muted mt-0.5">
            Multi-axis evaluation across compute spec parity, GPU accelerator fit, cost efficiency, storage, network topology, and region latency.
          </p>
        </div>

        {/* Summary Score Pills */}
        <div className="flex flex-wrap items-center gap-2">
          {candidateMeta.map((m, idx) => (
            <div
              key={m.key}
              className={`px-3 py-1.5 rounded-lg border text-sm flex items-center gap-2 ${
                idx === 0
                  ? 'bg-emerald-500/10 border-emerald-500/30'
                  : 'bg-bg-input/60 border-border-main'
              }`}
            >
              <span
                className="w-2.5 h-2.5 rounded-full shrink-0"
                style={{ backgroundColor: m.palette.stroke }}
              />
              <span className="font-normal text-text-muted">{m.label}:</span>
              <span className="font-extrabold text-text-main font-mono">{m.avgScore} pts</span>
            </div>
          ))}
        </div>
      </div>

      <div className="w-full h-[320px]">
        <ResponsiveContainer width="100%" height="100%">
          <RadarChart cx="50%" cy="50%" outerRadius="74%" data={chartData}>
            <PolarGrid stroke="#64748b" strokeOpacity={0.25} />
            <PolarAngleAxis
              dataKey="dimension"
              tick={{ fill: '#64748b', fontSize: 13, fontWeight: 700 }}
            />
            <PolarRadiusAxis
              angle={30}
              domain={[0, 100]}
              tick={{ fill: '#64748b', fontSize: 11 }}
              tickCount={5}
            />
            {candidateMeta.map((m) => (
              <Radar
                key={m.key}
                name={m.label}
                dataKey={m.key}
                stroke={m.palette.stroke}
                strokeWidth={2.5}
                fill={m.palette.fill}
                fillOpacity={m.palette.fillOpacity}
              />
            ))}
            <Tooltip
              contentStyle={{
                backgroundColor: '#0f172a',
                borderColor: '#10b981',
                borderRadius: '12px',
                color: '#f8fafc',
                fontSize: '13px',
              }}
            />
            <Legend
              wrapperStyle={{
                fontSize: '13px',
                fontWeight: 700,
                paddingTop: '8px',
              }}
            />
          </RadarChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};
