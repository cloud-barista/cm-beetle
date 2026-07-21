'use client';

import React, { useState } from 'react';
import { RecommendedInfra } from '../../types/migration';
import { 
  Network, 
  Server, 
  ChevronDown, 
  Globe, 
  Lock, 
  ShieldCheck, 
  Key, 
  Layers,
  HardDrive,
  Sliders,
  DollarSign
} from 'lucide-react';

interface TopologyMapProps {
  data: RecommendedInfra | null;
}

const getSpecHourlyCost = (specId: string): number => {
  const lower = (specId || '').toLowerCase();
  if (lower.includes('nano')) return 0.0052;
  if (lower.includes('micro')) return 0.0104;
  if (lower.includes('small')) return 0.0208;
  if (lower.includes('medium')) return 0.0416;
  if (lower.includes('2xlarge')) return 0.3328;
  if (lower.includes('xlarge')) return 0.1664;
  if (lower.includes('large')) return 0.0832;
  return 0.0416;
};

export const TopologyMap: React.FC<TopologyMapProps> = ({ data }) => {
  const [expandedNgCards, setExpandedNgCards] = useState<Record<string, boolean>>({});

  const toggleNgCard = (key: string) => {
    setExpandedNgCards(prev => ({ ...prev, [key]: !prev[key] }));
  };

  if (!data) {
    return (
      <div className="w-full h-full bg-slate-50 dark:bg-slate-950 border border-border-main rounded-xl p-8 flex items-center justify-center text-sm font-bold text-text-muted">
        No cloud architecture recommendation loaded.
      </div>
    );
  }

  // Safely parse data structures
  const targetVNet = data.targetVNet || {
    name: 'vnet-migrated',
    cidrBlock: '10.0.0.0/16',
    subnetInfoList: [{ name: 'subnet-default', ipv4_CIDR: '10.0.1.0/24' }]
  };
  const targetSshKey = data.targetSshKey || { name: 'key-mig01' };
  const targetSecurityGroupList = data.targetSecurityGroupList || [];
  const targetInfra = data.targetInfra || { name: 'infra', nodeGroups: [] };
  const targetNlbList = data.targetNlbList || [];

  const nodeGroups = targetInfra.nodeGroups || [];
  const subnets = (targetVNet.subnetInfoList && targetVNet.subnetInfoList.length > 0)
    ? targetVNet.subnetInfoList
    : [{ name: 'subnet-default', ipv4_CIDR: '10.0.1.0/24' }];

  // Calculate infrastructure estimated monthly & hourly cost
  const totalMonthlyCost = nodeGroups.reduce((sum, ng) => {
    const specId = ng.specId || '';
    const nodeCount = ng.nodeGroupSize || 1;
    return sum + (getSpecHourlyCost(specId) * nodeCount * 720);
  }, 0).toFixed(2);

  const totalHourlyCost = nodeGroups.reduce((sum, ng) => {
    const specId = ng.specId || '';
    const nodeCount = ng.nodeGroupSize || 1;
    return sum + (getSpecHourlyCost(specId) * nodeCount);
  }, 0).toFixed(3);

  return (
    <div className="w-full min-h-[300px] bg-slate-50 dark:bg-slate-950 border border-border-main rounded-xl overflow-hidden relative flex flex-col">
      {/* Topology Header Bar */}
      <div className="bg-white/80 dark:bg-slate-900/80 backdrop-blur-md px-5 py-3 border-b border-border-main flex justify-between items-center z-10 shrink-0">
        <div className="flex items-center space-x-2">
          <Network className="w-5 h-5 text-emerald-600 dark:text-emerald-400" />
          <span className="font-extrabold text-sm text-text-main">Topology Visualization</span>
        </div>

        <div className="flex items-center space-x-2 font-mono text-xs">
          {/* Estimated Monthly Cost Badge */}
          <span className="bg-emerald-500/10 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 px-3 py-1 rounded-full font-bold flex items-center space-x-1">
            <DollarSign className="w-3.5 h-3.5" />
            <span>Est. Cost: ${totalMonthlyCost}/mo (${totalHourlyCost}/hr)</span>
          </span>

          <span className="bg-emerald-100 dark:bg-emerald-950 text-emerald-700 dark:text-emerald-400 px-2.5 py-1 rounded-full font-bold">
            {nodeGroups.length} NodeGroup(s)
          </span>
          <span className="bg-blue-100 dark:bg-blue-950 text-blue-700 dark:text-blue-400 px-2.5 py-1 rounded-full font-bold">
            {subnets.length} Subnet(s)
          </span>
        </div>
      </div>

      {/* Main Structured Architecture Container View (Tab 3 NodeGroup List Style) */}
      <div className="p-5 space-y-5 font-sans text-sm">
        {/* Outer VPC / VNet Box */}
        <div className="border-2 border-emerald-400 dark:border-emerald-800/40 bg-emerald-500/5 rounded-2xl p-5 space-y-5">
          
          {/* VPC Header Bar */}
          <div className="flex flex-wrap justify-between items-center gap-3 border-b border-emerald-200 dark:border-emerald-800/20 pb-3">
            <span className="font-extrabold text-emerald-600 dark:text-emerald-400 flex items-center space-x-1.5 font-mono">
              <Network className="w-4 h-4 animate-pulse" />
              <span>VPC / VNet: {targetVNet.name || 'vnet-migrated'} ({targetVNet.cidrBlock || '10.0.0.0/16'})</span>
            </span>

            {/* Associated Credentials & Firewalls */}
            <div className="flex flex-wrap gap-2">
              {targetSshKey && (
                <span className="bg-amber-100 dark:bg-yellow-950/40 border border-amber-300 dark:border-yellow-900/30 text-amber-600 dark:text-yellow-400 text-xs px-2 py-0.5 rounded font-extrabold flex items-center space-x-1 font-mono">
                  <HardDrive className="w-3 h-3" />
                  <span>Key: {targetSshKey.name || 'default-key'}</span>
                </span>
              )}

              {targetSecurityGroupList.map((sg: any, sgIdx: number) => (
                <span key={sgIdx} className="bg-orange-100 dark:bg-orange-950/40 border border-orange-300 dark:border-orange-900/30 text-orange-600 dark:text-orange-400 text-xs px-2 py-0.5 rounded font-extrabold flex items-center space-x-1 font-mono">
                  <Sliders className="w-3 h-3" />
                  <span>SG: {sg.name} ({sg.firewallRules?.length || 0} Rules)</span>
                </span>
              ))}
            </div>
          </div>

          {/* Target Managed NLB (If exists) */}
          {targetNlbList && targetNlbList.length > 0 && (
            <div className="mb-4 relative z-10 flex flex-col items-center justify-center border-b border-border-main/20 pb-4">
              {targetNlbList.map((nlb, nlbIdx) => {
                const targetNgName = nlb.targetGroup?.nodeGroupId || 'VM NodeGroup';
                return (
                  <div key={nlbIdx} className="w-full max-w-xl bg-bg-panel border border-teal-300 dark:border-teal-800/40 rounded-xl p-4 space-y-3 shadow-sm">
                    <div className="flex justify-between items-center border-b border-border-main/40 pb-2">
                      <span className="font-extrabold text-teal-900 dark:text-teal-300 text-sm flex items-center">
                        <span className="w-2 h-2 bg-teal-500 rounded-full mr-2 animate-pulse" />
                        Target Managed NLB
                      </span>
                      <span className="text-xs font-bold px-2 py-0.5 bg-teal-100 dark:bg-teal-950/40 text-teal-600 dark:text-teal-400 rounded-full border border-teal-200 dark:border-teal-800/40 uppercase">
                        {nlb.type || 'PUBLIC'} MODE
                      </span>
                    </div>

                    <div className="space-y-2 font-mono text-xs text-text-muted">
                      <div className="flex items-center space-x-2">
                        <span className="font-sans font-normal shrink-0">Traffic Ingress ➔</span>
                        <span className="text-text-main font-bold">Listener Port: {nlb.listener?.port || 'ALL'}</span>
                      </div>
                      <div className="pl-3.5 border-l-2 border-teal-500/30 py-0.5 space-y-1">
                        <div className="text-xs text-teal-600 dark:text-teal-400 font-bold font-sans">
                          ▼ Target Routing Group
                        </div>
                        <div className="text-text-muted">
                          Target NodeGroup: <span className="text-text-main font-bold">{targetNgName}</span>
                        </div>
                      </div>
                    </div>
                  </div>
                );
              })}
            </div>
          )}

          {/* Subnet Zone Container */}
          {subnets.map((sub: any, subIdx: number) => {
            const subnetCidr = sub.ipv4_CIDR || sub.cidrBlock || '10.0.1.0/24';

            return (
              <div key={subIdx} className="border border-dashed border-emerald-400 dark:border-emerald-800/30 bg-bg-panel/40 rounded-xl p-4 space-y-3">
                <div className="flex items-center text-xs text-text-muted font-mono pb-2 border-b border-border-main">
                  <span>Subnet: {sub.name} ({subnetCidr})</span>
                </div>

                {/* Node Groups Accordion List (Tab 3 Format) */}
                <div className="divide-y divide-border-main/20 rounded-xl border border-border-main/30 overflow-hidden">
                  {nodeGroups.map((ng, idx) => {
                    const cardKey = `${subIdx}-${idx}`;
                    const isExpanded = !!expandedNgCards[cardKey];
                    const nodeCount = ng.nodeGroupSize || 1;
                    const nodesArray = Array.from({ length: nodeCount });
                    const rootDiskSize = (ng as any).rootDiskSize || '30';
                    const specId = ng.specId || 'custom-spec';
                    const imageId = ng.imageId || 'ubuntu-22.04';
                    const vcpu = (ng as any).vCPU || (specId.includes('small') ? 2 : specId.includes('xlarge') ? 4 : 2);
                    const memGiB = (ng as any).memoryGiB || (specId.includes('xlarge') ? 16 : 4);
                    const ngHourlyCost = getSpecHourlyCost(specId) * nodeCount;
                    const ngMonthlyCost = (ngHourlyCost * 720).toFixed(2);

                    // Associated security groups
                    const ngSgs = (targetSecurityGroupList || []).filter((sg: any) => {
                      const sgIds = ng.securityGroupIds || [];
                      return sgIds.includes(sg.name) || sg.name?.toLowerCase().includes(ng.name.toLowerCase());
                    });

                    return (
                      <div key={idx} className={isExpanded ? 'bg-bg-input/10' : 'bg-bg-panel/30 hover:bg-bg-input/10 transition'}>
                        {/* Summary Bar - Always Visible */}
                        <button
                          onClick={() => toggleNgCard(cardKey)}
                          className="w-full flex flex-wrap items-center gap-x-3 gap-y-1.5 px-3 py-2.5 text-left cursor-pointer"
                        >
                          <ChevronDown className={`w-3.5 h-3.5 text-text-muted flex-shrink-0 transition-transform ${isExpanded ? 'rotate-180' : ''}`} />

                          {/* Node Group label + name */}
                          <div className="flex items-center gap-1.5 flex-1 min-w-0 min-w-[180px]">
                            <Layers className="w-3.5 h-3.5 text-emerald-600 dark:text-emerald-400 flex-shrink-0" />
                            <div className="min-w-0">
                              <span className="text-xs text-text-muted font-normal block leading-none mb-0.5">Node Group</span>
                              <span className="text-sm font-bold text-text-main block" title={ng.name}>{ng.name}</span>
                            </div>
                          </div>

                          {/* Spec: vCPU · Memory · instance type */}
                          <div className="flex items-center gap-1.5 flex-shrink-0 flex-wrap font-mono">
                            <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">
                              {vcpu} vCPU
                            </span>
                            <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">
                              {memGiB} GiB
                            </span>
                            <span className="text-xs font-mono text-text-muted">
                              {specId.split('-').pop() || specId}
                            </span>
                          </div>

                          {/* OS Image */}
                          <span className="bg-teal-500/10 border border-teal-500/20 text-teal-600 dark:text-teal-400 text-xs font-extrabold px-1.5 py-0.5 rounded flex-shrink-0 font-mono truncate max-w-[280px]" title={imageId}>
                            {imageId}
                          </span>

                          {/* Root Disk */}
                          <span className="text-xs font-mono text-text-muted flex-shrink-0">{rootDiskSize} GB</span>

                          {/* Security Groups */}
                          {ngSgs.map((sg: any, sgIdx: number) => (
                            <span key={sgIdx} className="bg-orange-50 dark:bg-orange-950/20 border border-orange-200 dark:border-orange-900/30 text-orange-600 dark:text-orange-400 text-xs px-1.5 py-0.5 rounded font-mono flex-shrink-0">
                              {sg.name}
                            </span>
                          ))}

                          {/* Estimated Monthly Cost */}
                          <span className="bg-emerald-500/10 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded flex-shrink-0 font-mono">
                            ${ngMonthlyCost}/mo
                          </span>

                          {/* Node Count Badge */}
                          <span className="ml-auto flex-shrink-0 text-xs font-bold bg-emerald-100 dark:bg-emerald-950/40 border border-emerald-300 dark:border-emerald-800/40 text-emerald-600 dark:text-emerald-400 px-1.5 py-0.5 rounded">
                            ×{nodeCount}
                          </span>
                        </button>

                        {/* Expanded Detail Panel */}
                        {isExpanded && (
                          <div className="border-t border-border-main/20 px-4 py-3 space-y-3 bg-bg-input/5">
                            {/* Summary Grid: Spec · Image · Root Disk · Security Group · Size · Cost */}
                            <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-6 gap-3">
                              {/* Spec Card */}
                              <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                <span className="block text-xs font-bold text-text-muted mb-1">Spec</span>
                                <div className="flex flex-wrap gap-1 font-mono">
                                  <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">
                                    {vcpu} vCPU
                                  </span>
                                  <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">
                                    {memGiB} GiB
                                  </span>
                                </div>
                                <span className="text-xs font-mono text-text-muted mt-1 block" title={specId}>
                                  {specId}
                                </span>
                              </div>

                              {/* Image Card */}
                              <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                <span className="block text-xs font-bold text-text-muted mb-1">Image</span>
                                <span className="bg-teal-500/10 border border-teal-500/20 text-teal-600 dark:text-teal-400 text-xs font-extrabold px-1.5 py-0.5 rounded font-mono inline-block break-all" title={imageId}>
                                  {imageId}
                                </span>
                              </div>

                              {/* Root Disk Card */}
                              <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                <span className="block text-xs font-bold text-text-muted mb-1">Root Disk</span>
                                <span className="text-sm font-extrabold text-text-main font-mono">{rootDiskSize} GB</span>
                              </div>

                              {/* Security Group Card */}
                              <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                <span className="block text-xs font-bold text-text-muted mb-1">Security Group</span>
                                <div className="flex flex-col gap-1 font-mono">
                                  {ngSgs.length > 0 ? (
                                    ngSgs.map((sg: any, i: number) => (
                                      <span key={i} className="bg-orange-50 dark:bg-orange-950/20 border border-orange-200 dark:border-orange-900/30 text-orange-600 dark:text-orange-400 text-xs px-1.5 py-0.5 rounded">
                                        {sg.name}
                                      </span>
                                    ))
                                  ) : (
                                    <span className="text-xs text-text-muted">—</span>
                                  )}
                                </div>
                              </div>

                              {/* Estimated Cost Card */}
                              <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                <span className="block text-xs font-bold text-text-muted mb-1">Est. Cost</span>
                                <span className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 font-mono">${ngMonthlyCost}/mo</span>
                                <span className="text-xs text-text-muted block font-mono">${ngHourlyCost.toFixed(3)}/hr</span>
                              </div>

                              {/* Size Card */}
                              <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                <span className="block text-xs font-bold text-text-muted mb-1">Size</span>
                                <span className="text-2xl font-extrabold text-emerald-600 dark:text-emerald-400 font-mono">{nodeCount}</span>
                                <span className="text-xs text-text-muted ml-1 font-mono">node{nodeCount > 1 ? 's' : ''}</span>
                              </div>
                            </div>

                            {/* Node Instances Grid */}
                            <div className="space-y-1">
                              <span className="text-xs text-text-muted font-semibold">Nodes</span>
                              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-1.5 font-mono text-xs">
                                {nodesArray.map((_, nodeIdx) => {
                                  const suffix = String(nodeIdx + 1).padStart(2, '0');
                                  const nodeName = `${ng.name}-${suffix}`;
                                  return (
                                    <div key={nodeIdx} className="bg-bg-panel border border-emerald-500/10 px-2.5 py-1.5 rounded-lg flex items-center gap-1.5">
                                      <Server className="w-3 h-3 text-emerald-600 dark:text-emerald-400 flex-shrink-0" />
                                      <span className="text-text-muted truncate">{nodeName}</span>
                                    </div>
                                  );
                                })}
                              </div>
                            </div>
                          </div>
                        )}
                      </div>
                    );
                  })}
                </div>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
};
