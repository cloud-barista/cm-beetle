'use client';

import React from 'react';
import {
  Server,
  Cloud,
  CheckCircle2,
  RefreshCw,
  AlertTriangle,
  Activity,
  ArrowRight,
  Shield,
  Key,
  Network,
  Layers,
  Clock,
  Zap,
  Globe
} from 'lucide-react';
import {
  OnpremNode,
  CloudModelEnvelope,
  CheckSSHReadyResponse
} from '../../types/migration';

export interface SplitViewMigrationJob {
  id: string;
  reqId: string;
  infraId: string;
  nsId: string;
  nameSeed: string;
  csp: string;
  region: string;
  status: 'Handling' | 'Success' | 'Failed';
  startTime: string;
  elapsedSeconds: number;
  nodeGroupsCount: number;
  totalVms: number;
  logs: string[];
  vms?: { publicIp: string; privateIp: string; specId: string; name: string }[];
  error?: string;
  isSample?: boolean;
  targetNlbList?: any[];
  nlbStatus?: 'Idle' | 'Provisioning' | 'Success' | 'Failed';
}

interface MigrationSplitScreenProps {
  sourceNodes: OnpremNode[];
  sourceModelName: string;
  selectedCloudModel: CloudModelEnvelope | null;
  activeJob: SplitViewMigrationJob | undefined;
  sshReadyResponse?: CheckSSHReadyResponse;
  sshCooldownSec: number;
  sshLoading: boolean;
  onCheckSshReady: (filterNodeName?: string) => void;
}

interface PlannedTargetNode {
  name: string;
  specId: string;
  publicIp: string;
  privateIp: string;
  nodeGroupName?: string;
}

export const MigrationSplitScreen: React.FC<MigrationSplitScreenProps> = ({
  sourceNodes,
  sourceModelName,
  selectedCloudModel,
  activeJob,
  sshReadyResponse,
  sshCooldownSec,
  sshLoading,
  onCheckSshReady
}) => {
  // Build target nodes list from activeJob.vms if provisioned, or expand from targetInfra.nodeGroups
  const targetNodes: PlannedTargetNode[] = React.useMemo(() => {
    if (activeJob?.vms && activeJob.vms.length > 0) {
      return activeJob.vms.map((vm) => ({
        name: vm.name,
        specId: vm.specId || 'cloud-spec',
        publicIp: vm.publicIp || '-',
        privateIp: vm.privateIp || '-'
      }));
    }

    const nodeGroups = selectedCloudModel?.cloudInfraModel?.targetInfra?.nodeGroups || [];
    if (nodeGroups.length > 0) {
      const expanded: PlannedTargetNode[] = [];
      nodeGroups.forEach((ng) => {
        const size = Math.max(1, Number(ng.nodeGroupSize || 1));
        for (let i = 1; i <= size; i++) {
          const suffix = size > 1 ? `-${String(i).padStart(2, '0')}` : '-01';
          expanded.push({
            name: `${ng.name || 'node'}${suffix}`,
            specId: ng.specId || 'cloud-spec',
            publicIp: activeJob?.status === 'Handling' ? 'Allocating...' : 'Pending',
            privateIp: activeJob?.status === 'Handling' ? 'Allocating...' : 'Pending',
            nodeGroupName: ng.name
          });
        }
      });
      return expanded;
    }

    return sourceNodes.map((sn, idx) => ({
      name: sn.hostname || `target-node-${String(idx + 1).padStart(2, '0')}`,
      specId: 'Pending Spec',
      publicIp: 'Pending',
      privateIp: 'Pending'
    }));
  }, [activeJob, selectedCloudModel, sourceNodes]);

  const targetInfra = selectedCloudModel?.cloudInfraModel;
  const vnetName = targetInfra?.targetVNet?.name || 'Target VNet';
  const subnetCount = targetInfra?.targetVNet?.subnetInfoList?.length || 1;
  const sgCount = targetInfra?.targetSecurityGroupList?.length || 1;
  const sshKeyName = targetInfra?.targetSshKey?.name || 'Target SSH Key';
  const ngCount = targetInfra?.targetInfra?.nodeGroups?.length || activeJob?.nodeGroupsCount || 1;
  const nlbCount = activeJob?.targetNlbList?.length || targetInfra?.targetNlbList?.length || 0;

  const getStageState = (stageIndex: number): 'completed' | 'active' | 'pending' | 'failed' => {
    if (!activeJob) return 'pending';
    if (activeJob.status === 'Failed') {
      return stageIndex <= 4 ? 'failed' : 'pending';
    }
    if (activeJob.status === 'Success') {
      if (stageIndex === 6) {
        if (activeJob.nlbStatus === 'Provisioning') return 'active';
        if (activeJob.nlbStatus === 'Failed') return 'failed';
        return 'completed';
      }
      return 'completed';
    }
    // Handling state: animate progression across stages 1..5 based on elapsedSeconds
    const elapsed = activeJob.elapsedSeconds || 0;
    if (stageIndex === 1) return elapsed >= 2 ? 'completed' : 'active';
    if (stageIndex === 2) return elapsed >= 4 ? 'completed' : elapsed >= 2 ? 'active' : 'pending';
    if (stageIndex === 3) return elapsed >= 6 ? 'completed' : elapsed >= 4 ? 'active' : 'pending';
    if (stageIndex === 4) return elapsed >= 8 ? 'completed' : elapsed >= 6 ? 'active' : 'pending';
    if (stageIndex === 5) return elapsed >= 8 ? 'active' : 'pending';
    return 'pending';
  };

  const pipelineStages = [
    {
      index: 1,
      title: '1. VPC / VNet',
      detail: vnetName,
      icon: Globe
    },
    {
      index: 2,
      title: '2. Subnet(s)',
      detail: `${subnetCount} Subnet(s)`,
      icon: Layers
    },
    {
      index: 3,
      title: '3. Security Group(s)',
      detail: `${sgCount} Group(s)`,
      icon: Shield
    },
    {
      index: 4,
      title: '4. SSH Key',
      detail: sshKeyName,
      icon: Key
    },
    {
      index: 5,
      title: '5. Node Group(s)',
      detail: `${ngCount} Group(s) / ${targetNodes.length} Nodes`,
      icon: Server
    },
    {
      index: 6,
      title: '6. Network Load Balancer',
      detail: nlbCount > 0 ? `${nlbCount} NLB Configured` : 'Optional / Ready',
      icon: Network
    }
  ];

  const readyNodesCount = sshReadyResponse?.readyNodes ?? 0;

  return (
    <div className="glass-panel p-5 sm:p-6 rounded-2xl border border-border-main space-y-5 animate-fade-in">
      {/* Top Split Screen Summary Bar */}
      <div className="flex flex-wrap items-center justify-between gap-3 border-b border-border-main/30 pb-4">
        <div className="flex items-center gap-3">
          <span className="px-3 py-1 bg-emerald-500/15 border border-emerald-500/30 text-emerald-500 rounded-lg text-xs font-extrabold font-mono">
            16:9 Live Topology Split View
          </span>
          <span className="text-sm font-normal text-text-muted">
            Source Model: <strong className="font-extrabold text-text-main">{sourceModelName}</strong>
          </span>
          <ArrowRight className="w-4 h-4 text-emerald-500" />
          <span className="text-sm font-normal text-text-muted">
            Target Cloud:{' '}
            <strong className="font-extrabold text-emerald-500">
              {activeJob
                ? `${activeJob.csp} (${activeJob.region})`
                : targetInfra?.targetCloud
                ? `${targetInfra.targetCloud.csp.toUpperCase()} (${targetInfra.targetCloud.region})`
                : 'Select Target Model'}
            </strong>
          </span>
        </div>

        <div className="flex items-center gap-2.5">
          {activeJob && (
            <div className="px-3 py-1.5 bg-bg-panel border border-border-main rounded-xl text-xs font-mono flex items-center gap-2">
              <Clock className="w-3.5 h-3.5 text-emerald-500" />
              <span className="font-normal text-text-muted">Elapsed:</span>
              <span className="font-extrabold text-text-main">{activeJob.elapsedSeconds}s</span>
            </div>
          )}

          {activeJob?.status === 'Success' && (
            <button
              type="button"
              onClick={() => onCheckSshReady()}
              disabled={sshCooldownSec > 0 || sshLoading}
              className="px-3.5 py-1.5 bg-emerald-500/15 hover:bg-emerald-500/25 disabled:opacity-50 border border-emerald-500/30 text-emerald-500 rounded-xl text-xs font-extrabold transition flex items-center gap-1.5 cursor-pointer"
            >
              <Activity className="w-3.5 h-3.5 text-emerald-500" />
              <span>
                {sshCooldownSec > 0
                  ? `SSH Cooldown (${sshCooldownSec}s)`
                  : 'Verify All Nodes SSH Readiness'}
              </span>
            </button>
          )}
        </div>
      </div>

      {/* 3-Column 16:9 Split Grid: Left (Source 15 Nodes) | Center (6-Stage Pipeline) | Right (Target 15 Nodes) */}
      <div className="grid grid-cols-1 xl:grid-cols-12 gap-5 items-stretch">
        {/* LEFT COLUMN (5 cols): Source On-Premise Nodes */}
        <div className="xl:col-span-5 bg-bg-panel/40 border border-border-main/60 rounded-2xl p-4 flex flex-col justify-between space-y-3">
          <div className="space-y-3">
            <div className="flex items-center justify-between border-b border-border-main/30 pb-2.5">
              <div className="flex items-center gap-2">
                <Server className="w-4 h-4 text-emerald-500" />
                <h3 className="text-sm font-extrabold text-text-main">
                  Source On-Premise Nodes ({sourceNodes.length})
                </h3>
              </div>
              <span className="text-xs font-normal text-text-muted font-mono">
                onpremiseInfraModel
              </span>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-2.5">
              {sourceNodes.map((node, idx) => {
                const vcpu = (node.cpu?.cpus || 1) * (node.cpu?.threads || node.cpu?.cores || 2);
                const hasGpu = Array.isArray(node.gpuCards) && node.gpuCards.length > 0;
                const memGb = node.memory?.totalSize || 16;
                const rootDiskGb = node.rootDisk?.totalSize || 100;

                return (
                  <div
                    key={node.machineId || idx}
                    className={`p-2.5 rounded-xl border transition flex flex-col justify-between space-y-1.5 ${
                      hasGpu
                        ? 'bg-emerald-500/10 border-emerald-500/40'
                        : 'bg-bg-input/60 border-border-main/50'
                    }`}
                  >
                    <div className="flex items-center justify-between gap-1">
                      <span
                        className="text-sm font-extrabold text-text-main font-mono truncate"
                        title={node.hostname || node.machineId}
                      >
                        {node.hostname || node.machineId}
                      </span>
                      {hasGpu && (
                        <span className="px-1.5 py-0.5 bg-emerald-500/20 text-emerald-500 border border-emerald-500/30 rounded text-xs font-extrabold font-mono shrink-0">
                          GPU
                        </span>
                      )}
                    </div>

                    <div className="text-xs font-mono flex items-center justify-between text-text-muted">
                      <span className="font-normal text-text-muted">Spec:</span>
                      <span className="font-extrabold text-text-main">
                        {vcpu}c / {memGb}G / {rootDiskGb}G
                      </span>
                    </div>

                    {node.workloadTag && (
                      <div
                        className="text-xs font-normal text-teal-600 dark:text-teal-400 truncate border-t border-border-main/20 pt-1"
                        title={node.workloadTag}
                      >
                        {node.workloadTag}
                      </div>
                    )}
                  </div>
                );
              })}
            </div>
          </div>

          <div className="pt-2 border-t border-border-main/20 flex items-center justify-between text-xs">
            <span className="font-normal text-text-muted">Total Source Capacity:</span>
            <span className="font-extrabold text-text-main font-mono">
              {sourceNodes.length} Nodes ({sourceNodes.filter((n) => (n.gpuCards?.length || 0) > 0).length} GPU Nodes)
            </span>
          </div>
        </div>

        {/* CENTER COLUMN (2 cols): 6-Stage Migration Pipeline Stream */}
        <div className="xl:col-span-2 bg-bg-panel/60 border border-border-main/60 rounded-2xl p-4 flex flex-col justify-between space-y-3">
          <div className="space-y-3">
            <div className="flex items-center justify-between border-b border-border-main/30 pb-2.5">
              <div className="flex items-center gap-1.5">
                <Zap className="w-4 h-4 text-emerald-500" />
                <h3 className="text-sm font-extrabold text-text-main">
                  Migration Stream
                </h3>
              </div>
            </div>

            <div className="space-y-2.5">
              {pipelineStages.map((stage) => {
                const state = getStageState(stage.index);
                const IconComponent = stage.icon;

                return (
                  <div
                    key={stage.index}
                    className={`p-2.5 rounded-xl border transition flex items-center justify-between gap-2 ${
                      state === 'completed'
                        ? 'bg-emerald-500/10 border-emerald-500/40'
                        : state === 'active'
                        ? 'bg-teal-500/15 border-teal-500/50 animate-pulse'
                        : state === 'failed'
                        ? 'bg-red-500/10 border-red-500/40'
                        : 'bg-bg-input/40 border-border-main/40 opacity-75'
                    }`}
                  >
                    <div className="flex items-center gap-2 min-w-0">
                      <IconComponent
                        className={`w-4 h-4 shrink-0 ${
                          state === 'completed'
                            ? 'text-emerald-500'
                            : state === 'active'
                            ? 'text-teal-400'
                            : state === 'failed'
                            ? 'text-red-400'
                            : 'text-text-muted'
                        }`}
                      />
                      <div className="min-w-0">
                        <div className="text-xs font-extrabold text-text-main truncate">
                          {stage.title}
                        </div>
                        <div className="text-xs font-normal text-text-muted font-mono truncate">
                          {stage.detail}
                        </div>
                      </div>
                    </div>

                    <div className="shrink-0">
                      {state === 'completed' ? (
                        <CheckCircle2 className="w-4 h-4 text-emerald-500" />
                      ) : state === 'active' ? (
                        <RefreshCw className="w-4 h-4 text-teal-400 animate-spin" />
                      ) : state === 'failed' ? (
                        <AlertTriangle className="w-4 h-4 text-red-400" />
                      ) : (
                        <span className="w-2 h-2 rounded-full bg-border-main block" />
                      )}
                    </div>
                  </div>
                );
              })}
            </div>
          </div>

          <div className="pt-2 border-t border-border-main/20 space-y-1 text-xs font-mono">
            <div className="flex items-center justify-between">
              <span className="font-normal text-text-muted">Job Status:</span>
              <span
                className={`font-extrabold ${
                  activeJob?.status === 'Success'
                    ? 'text-emerald-500'
                    : activeJob?.status === 'Handling'
                    ? 'text-teal-400'
                    : activeJob?.status === 'Failed'
                    ? 'text-red-400'
                    : 'text-text-muted'
                }`}
              >
                {activeJob?.status || 'Standby'}
              </span>
            </div>
            <div className="flex items-center justify-between">
              <span className="font-normal text-text-muted">SSH Ready:</span>
              <span className="font-extrabold text-emerald-500">
                {sshReadyResponse
                  ? `${readyNodesCount} / ${sshReadyResponse.totalNodes}`
                  : activeJob?.status === 'Success'
                  ? `${targetNodes.length} Provisioned`
                  : '0 / ' + targetNodes.length}
              </span>
            </div>
          </div>
        </div>

        {/* RIGHT COLUMN (5 cols): Target Cloud Nodes */}
        <div className="xl:col-span-5 bg-bg-panel/40 border border-border-main/60 rounded-2xl p-4 flex flex-col justify-between space-y-3">
          <div className="space-y-3">
            <div className="flex items-center justify-between border-b border-border-main/30 pb-2.5">
              <div className="flex items-center gap-2">
                <Cloud className="w-4 h-4 text-teal-400" />
                <h3 className="text-sm font-extrabold text-text-main">
                  Target Cloud Nodes ({targetNodes.length})
                </h3>
              </div>
              <span className="text-xs font-normal text-text-muted font-mono">
                {activeJob?.infraId || targetInfra?.targetInfra?.name || 'cloudInfraModel'}
              </span>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-2.5">
              {targetNodes.map((tn, idx) => {
                const matchedSsh = sshReadyResponse?.nodeStatus?.find(
                  (n) =>
                    n.id === tn.name ||
                    n.name === tn.name ||
                    (tn.publicIp !== '-' && n.publicIP === tn.publicIp)
                );
                const isSshReady = Boolean(matchedSsh?.sshReady);
                const isProvisioned = activeJob?.status === 'Success';
                const isHandling = activeJob?.status === 'Handling';

                return (
                  <div
                    key={`${tn.name}-${idx}`}
                    className={`p-2.5 rounded-xl border transition flex flex-col justify-between space-y-1.5 ${
                      isSshReady
                        ? 'bg-emerald-500/10 border-emerald-500/50 shadow-sm shadow-emerald-500/10'
                        : isProvisioned
                        ? 'bg-teal-500/10 border-teal-500/40'
                        : isHandling
                        ? 'bg-bg-input/60 border-teal-500/30 animate-pulse'
                        : 'bg-bg-input/40 border-border-main/50'
                    }`}
                  >
                    <div className="flex items-center justify-between gap-1">
                      <span
                        className="text-sm font-extrabold text-text-main font-mono truncate"
                        title={tn.name}
                      >
                        {tn.name}
                      </span>

                      {isSshReady ? (
                        <span className="px-1.5 py-0.5 bg-emerald-500/20 text-emerald-500 border border-emerald-500/30 rounded text-xs font-extrabold shrink-0">
                          SSH Ready
                        </span>
                      ) : isProvisioned ? (
                        <span className="px-1.5 py-0.5 bg-teal-500/20 text-teal-400 border border-teal-500/30 rounded text-xs font-extrabold shrink-0">
                          Running
                        </span>
                      ) : isHandling ? (
                        <span className="px-1.5 py-0.5 bg-amber-500/20 text-amber-400 border border-amber-500/30 rounded text-xs font-extrabold shrink-0">
                          Creating
                        </span>
                      ) : (
                        <span className="px-1.5 py-0.5 bg-bg-panel text-text-muted border border-border-main rounded text-xs font-normal shrink-0">
                          Planned
                        </span>
                      )}
                    </div>

                    <div className="text-xs font-mono flex items-center justify-between">
                      <span className="font-normal text-text-muted">Spec:</span>
                      <span className="font-extrabold text-emerald-500 truncate max-w-[120px]" title={tn.specId}>
                        {tn.specId}
                      </span>
                    </div>

                    <div className="text-xs font-mono flex items-center justify-between border-t border-border-main/20 pt-1">
                      <span className="font-normal text-text-muted">IP:</span>
                      <span className="font-extrabold text-text-main select-all truncate">
                        {tn.publicIp}
                      </span>
                    </div>
                  </div>
                );
              })}
            </div>
          </div>

          <div className="pt-2 border-t border-border-main/20 flex items-center justify-between text-xs">
            <span className="font-normal text-text-muted">Target Provisioning Status:</span>
            <span className="font-extrabold text-emerald-500 font-mono">
              {activeJob?.status === 'Success'
                ? `${targetNodes.length} / ${targetNodes.length} Nodes Active`
                : activeJob?.status === 'Handling'
                ? 'Provisioning Cloud Nodes...'
                : `${targetNodes.length} Planned Nodes`}
            </span>
          </div>
        </div>
      </div>
    </div>
  );
};
