'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '../../store/migrationStore';
import { OnpremNode, OnpremInfra, OnpremModelEnvelope } from '../../types/migration';
import sampleData from '../../data/sampleSourceInfra.json';
import { SaveRevisionModal } from '../common/SaveRevisionModal';
import {
  RefreshCw, ChevronDown, ChevronUp, Plus, Trash2, Server, Cpu,
  HardDrive, Network, Shield, Save, CheckCircle2, Loader2,
  FileText, Copy, Sparkles, GitBranch, Layers, X, Sliders,
} from 'lucide-react';

const SAMPLE_INFRA: OnpremInfra = sampleData.sourceInfra as OnpremInfra;
const SAMPLE_MODEL: OnpremModelEnvelope = {
  id: 'sample-source-infra-1',
  name: '[Sample] web-haproxy-influxdb',
  description: '1 HAProxy/App node + 2 InfluxDB nodes with NLB (sample)',
  onpremiseInfraModel: SAMPLE_INFRA,
  version: '1.0',
  updatedTime: new Date().toISOString(),
};

export const SourceInfraRefinement: React.FC = () => {
  const {
    savedSourceModels, selectedSourceModel, selectSourceModel,
    fetchSavedSourceModels, saveSourceModel, updateSourceModel, deleteSourceModel,
  } = useMigrationStore();

  const [activeTunedNodeId, setActiveTunedNodeId] = useState<string>('');
  const [tunedNodes, setTunedNodes] = useState<OnpremNode[]>([]);
  const [isJsonOpen, setIsJsonOpen] = useState(false);
  const [tunedNetwork, setTunedNetwork] = useState<any>(null);
  const [newCidr, setNewCidr] = useState('');
  const [activeStep, setActiveStep] = useState<number>(1);
  const [isModelLoaded, setIsModelLoaded] = useState(false);
  const [excludedNodeIds, setExcludedNodeIds] = useState<string[]>([]);
  const [newRuleDir, setNewRuleDir] = useState('inbound');
  const [newRuleProto, setNewRuleProto] = useState('tcp');
  const [newRulePort, setNewRulePort] = useState('');
  const [newRuleCidr, setNewRuleCidr] = useState('0.0.0.0/0');
  const [tuningSourceSaveSuccess, setTuningSourceSaveSuccess] = useState(false);
  const [showSaveModal, setShowSaveModal] = useState(false);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const [deleteError, setDeleteError] = useState('');
  const [deleteConfirmText, setDeleteConfirmText] = useState('');

  // Ensure sample model is always in the list
  const allModels: OnpremModelEnvelope[] = [
    SAMPLE_MODEL,
    ...savedSourceModels.filter(m => m.id !== 'sample-source-infra-1'),
  ];

  useEffect(() => { fetchSavedSourceModels(); }, []);

  const handleLoadModel = () => {
    if (!selectedSourceModel || !selectedSourceModel.onpremiseInfraModel) return;
    setTunedNodes(JSON.parse(JSON.stringify(selectedSourceModel.onpremiseInfraModel.nodes || [])));
    setTunedNetwork(JSON.parse(JSON.stringify(selectedSourceModel.onpremiseInfraModel.network || { ipv4Networks: {}, ipv6Networks: {} })));
    setExcludedNodeIds([]);
    setIsModelLoaded(true);
    setActiveStep(2); // Unlock Step 2: Review and Editing
    if (selectedSourceModel.onpremiseInfraModel.nodes && selectedSourceModel.onpremiseInfraModel.nodes.length > 0) {
      setActiveTunedNodeId(selectedSourceModel.onpremiseInfraModel.nodes[0]?.machineId || '');
    }
  };

  const handleDeleteModel = () => {
    if (!selectedSourceModel || selectedSourceModel.id === 'sample-source-infra-1') return;
    setDeleteConfirmText('');
    setShowDeleteConfirm(true);
  };

  const activeNode = tunedNodes.find((n) => n.machineId === activeTunedNodeId);

  const handleToggleNodeExclude = (machineId: string) => {
    setExcludedNodeIds(prev =>
      prev.includes(machineId)
        ? prev.filter(id => id !== machineId)
        : [...prev, machineId]
    );
  };

  const handleDeleteFirewallRule = (machineId: string, ruleIdx: number) => {
    setTunedNodes(prev =>
      prev.map(node =>
        node.machineId === machineId
          ? {
            ...node,
            firewallTable: (node.firewallTable || []).filter((_, idx) => idx !== ruleIdx)
          }
          : node
      )
    );
  };

  const handleAddFirewallRule = (machineId: string) => {
    if (!newRulePort) return;
    const newRule = {
      action: 'allow',
      direction: newRuleDir,
      protocol: newRuleProto,
      dstCIDR: newRuleCidr,
      dstPorts: newRulePort,
      srcCIDR: newRuleCidr,
      srcPorts: '*'
    };
    setTunedNodes(prev =>
      prev.map(node =>
        node.machineId === machineId
          ? {
            ...node,
            firewallTable: [...(node.firewallTable || []), newRule]
          }
          : node
      )
    );
    setNewRulePort('');
  };

  const handleSaveToDamselfly = async (result: { name: string; description: string; version: string; overwriteId: string | null }) => {
    if (!selectedSourceModel) return;
    const filteredNodes = tunedNodes.filter(n => !excludedNodeIds.includes(n.machineId));
    const updatedInfra = {
      ...selectedSourceModel.onpremiseInfraModel,
      nodes: filteredNodes,
      network: tunedNetwork || selectedSourceModel.onpremiseInfraModel.network
    };

    if (result.overwriteId) {
      await updateSourceModel(result.overwriteId, result.name, result.description, result.version, updatedInfra);
    } else {
      await saveSourceModel(result.name, result.description, result.version, updatedInfra);
    }

    setTunedNodes(filteredNodes);
    setExcludedNodeIds([]);
    setTuningSourceSaveSuccess(true);
    setActiveStep(3); // Advance to Step 3: Desired Cloud Target Specification
    setTimeout(() => setTuningSourceSaveSuccess(false), 2000);
  };

  const handleAddCidr = (cidr: string) => {
    if (!cidr || !tunedNetwork) return;
    const currentCidrs = tunedNetwork.ipv4Networks?.cidrBlocks || [];
    if (currentCidrs.includes(cidr)) return; // prevent duplicate
    const updatedNetwork = {
      ...tunedNetwork,
      ipv4Networks: {
        ...tunedNetwork.ipv4Networks,
        cidrBlocks: [...currentCidrs, cidr]
      }
    };
    setTunedNetwork(updatedNetwork);
  };

  const handleRemoveCidr = (cidrToRemove: string) => {
    if (!tunedNetwork) return;
    const currentCidrs = tunedNetwork.ipv4Networks?.cidrBlocks || [];
    const updatedNetwork = {
      ...tunedNetwork,
      ipv4Networks: {
        ...tunedNetwork.ipv4Networks,
        cidrBlocks: currentCidrs.filter((c: string) => c !== cidrToRemove)
      }
    };
    setTunedNetwork(updatedNetwork);
  };

  return (
    <div className="space-y-8 mx-auto pb-24">

      {/* ═══ STEP 1: Source Infrastructure Model Selection ═══ */}
      <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${activeStep >= 1 ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
        <div className="flex items-center space-x-3 mb-4 border-b border-border-main/40 pb-3">
          <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${isModelLoaded ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
            {isModelLoaded ? '✓' : '1'}
          </span>
          <h3 className="text-base font-extrabold text-text-main">Step 1: Source Infrastructure Model Selection</h3>
        </div>
        <div className="space-y-4">
          <div>
            <label className="block text-sm font-bold text-text-muted mb-2">Choose Source Model</label>
            <select
              value={selectedSourceModel?.id || ''}
              onChange={(e) => {
                const m = allModels.find(x => x.id === e.target.value) || null;
                selectSourceModel(m);
                setIsModelLoaded(false);
                setTunedNodes([]);
                setActiveStep(1);
              }}
              className="w-full max-w-md bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-3 text-sm font-bold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer mb-3"
            >
              <option value="">-- Choose Source Model --</option>
              {allModels.map(m => (
                <option key={m.id} value={m.id}>{m.name} (v{m.version || '1.0'})</option>
              ))}
            </select>
            <div className="flex items-center gap-3">
              <button
                onClick={handleLoadModel}
                disabled={!selectedSourceModel}
                className={`px-5 py-3 rounded-xl text-sm font-extrabold flex items-center transition cursor-pointer ${selectedSourceModel
                  ? 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 shadow-md shadow-emerald-500/25'
                  : 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'}`}
              >
                <RefreshCw className="w-4 h-4 mr-1.5" /> Load Model
              </button>
              {selectedSourceModel && selectedSourceModel.id !== 'sample-source-infra-1' && (
                <button
                  onClick={handleDeleteModel}
                  className="px-5 py-3 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-xl text-sm font-extrabold flex items-center transition cursor-pointer"
                >
                  <Trash2 className="w-4 h-4 mr-1.5" /> Delete Model
                </button>
              )}
            </div>
          </div>
        </div>
      </div>

      {/* -------------------------------------------------------------
          STEP 2: Source Infrastructure Review & Editing
         ------------------------------------------------------------- */}
      <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${isModelLoaded ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
        <div className="flex items-center justify-between mb-4 border-b border-border-main/40 pb-3">
          <div className="flex items-center space-x-3">
            <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${activeStep > 2 ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
              {activeStep > 2 ? '✓' : '2'}
            </span>
            <h3 className="text-base font-extrabold text-text-main">
              Step 2: Source Infrastructure Review & Editing
            </h3>
          </div>
          {selectedSourceModel && isModelLoaded && (
            <div className="flex items-center space-x-2">
              <button
                onClick={() => setIsJsonOpen(!isJsonOpen)}
                className="px-3 py-1.5 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-sm font-bold transition cursor-pointer flex items-center text-emerald-600 dark:text-emerald-400"
              >
                <Copy className="w-3.5 h-3.5 mr-1" />
                JSON View
              </button>
            </div>
          )}
        </div>



        {selectedSourceModel && isModelLoaded && (
          <div className="bg-bg-panel/40 border border-border-main/30 rounded-xl p-3.5 flex flex-col md:flex-row md:items-center justify-between text-sm space-y-2 md:space-y-0 mb-4">
            <div className="flex items-center space-x-2">
              <span className="text-text-muted font-bold">Loaded Model:</span>
              <span className="text-emerald-600 dark:text-emerald-600 dark:text-emerald-400 font-extrabold text-sm">{selectedSourceModel.name}</span>
              <span className="text-sm text-text-muted font-mono bg-bg-panel px-2 py-0.5 rounded border border-border-main/40">
                v{selectedSourceModel.version || '1.0'}
              </span>
            </div>
            {selectedSourceModel.updatedTime && (
              <div className="text-sm text-text-muted">
                Last Updated: <span className="text-text-main font-semibold">{new Date(selectedSourceModel.updatedTime).toLocaleString()}</span>
              </div>
            )}
          </div>
        )}

        {isModelLoaded && selectedSourceModel && (
          <div className="space-y-4">
            <div className="mt-2">
              {isJsonOpen ? (
                /* Raw JSON View */
                <div className="space-y-3">
                  <div className="flex justify-between items-center bg-bg-input px-3.5 py-2 rounded-xl border border-border-main/50">
                    <span className="text-sm text-text-muted font-mono">onpremiseInfraModel</span>
                    <button
                      onClick={() => {
                        navigator.clipboard.writeText(JSON.stringify(selectedSourceModel?.onpremiseInfraModel, null, 2));
                        alert('JSON copied!');
                      }}
                      className="px-3 py-1 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-xs font-bold transition flex items-center cursor-pointer text-emerald-600 dark:text-emerald-400"
                    >
                      <Copy className="w-3.5 h-3.5 mr-1" /> Copy JSON
                    </button>
                  </div>
                  <pre className="text-sm font-mono text-slate-800 dark:text-emerald-400 bg-bg-panel p-3.5 rounded-xl border border-border-main overflow-y-auto max-h-[400px] select-text">
                    {JSON.stringify(selectedSourceModel?.onpremiseInfraModel, null, 2)}
                  </pre>
                </div>
              ) : (
                /* Structured Hierarchy Spec Editor UI (Vertical 4 Rows: Network -> Server List -> Server Details/Tuning -> NLB) */
                <div className="space-y-6">

                  {/* Row 1: Network Configuration */}
                  <div className="space-y-3 p-5 bg-bg-input/40 border border-border-main/50 rounded-xl">
                    <h4 className="text-sm font-bold text-emerald-600 dark:text-emerald-400 flex items-center">
                      <Network className="w-4 h-4 mr-1.5 text-emerald-600 dark:text-emerald-400" />
                      1. Network Configuration
                    </h4>
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm mt-2">
                      {selectedSourceModel.onpremiseInfraModel.network?.ipv4Networks?.defaultGateways && (
                        <div className="bg-bg-panel/20 p-3 rounded-lg border border-border-main/20">
                          <span className="text-text-muted font-semibold block mb-2">Gateways for VNet/Subnet estimation:</span>
                          <div className="space-y-1">
                            {selectedSourceModel.onpremiseInfraModel.network.ipv4Networks.defaultGateways.map((gw, idx) => (
                              <div key={idx} className="flex justify-between bg-bg-panel px-2.5 py-1.5 rounded-lg border border-border-main/30 font-mono text-sm">
                                <span className="text-text-muted font-normal">{gw.interfaceName}</span>
                                <span className="text-text-main font-bold">{gw.ip}</span>
                              </div>
                            ))}
                          </div>
                        </div>
                      )}
                      {tunedNetwork?.ipv4Networks && (
                        <div className="bg-bg-panel/20 p-3 rounded-lg border border-border-main/20 space-y-2">
                          <span className="text-text-muted font-semibold block">Source Network CIDR Block:</span>
                          <div className="flex flex-wrap gap-1.5 pt-0.5">
                            {(tunedNetwork.ipv4Networks.cidrBlocks || []).map((cidr: string, idx: number) => (
                              <span key={idx} className="bg-emerald-500/10 border border-emerald-500/25 text-emerald-600 dark:text-emerald-600 dark:text-emerald-400 font-mono text-sm px-2 py-1 rounded-md font-extrabold flex items-center space-x-1.5 animate-fade-in">
                                <span>{cidr}</span>
                                <button
                                  onClick={() => handleRemoveCidr(cidr)}
                                  className="hover:text-red-400 font-extrabold text-sm ml-1 transition cursor-pointer"
                                >
                                  ✕
                                </button>
                              </span>
                            ))}
                            {(tunedNetwork.ipv4Networks.cidrBlocks || []).length === 0 && (
                              <span className="text-sm text-text-muted italic">No CIDR blocks declared.</span>
                            )}
                          </div>
                          <div className="flex items-center space-x-2 pt-1">
                            <input
                              type="text"
                              value={newCidr}
                              onChange={(e) => setNewCidr(e.target.value)}
                              placeholder="e.g., 10.0.0.0/16"
                              onKeyDown={(e) => {
                                if (e.key === 'Enter') {
                                  handleAddCidr(newCidr);
                                  setNewCidr('');
                                }
                              }}
                              className="flex-1 bg-bg-panel border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:ring-1 focus:ring-emerald-500"
                            />
                            <button
                              onClick={() => {
                                handleAddCidr(newCidr);
                                setNewCidr('');
                              }}
                              className="bg-emerald-500 hover:bg-emerald-600 text-slate-950 font-bold px-4 py-1.5 rounded-lg text-sm cursor-pointer transition"
                            >
                              Add
                            </button>
                          </div>
                        </div>
                      )}
                    </div>
                  </div>

                  {/* Row 2: Server List */}
                  <div className="space-y-3 p-5 bg-bg-input/40 border border-border-main/50 rounded-xl">
                    <h4 className="text-sm font-bold text-emerald-600 dark:text-emerald-400 flex items-center">
                      <Sliders className="w-4 h-4 mr-1.5 text-emerald-600 dark:text-emerald-400" />
                      2. Server List ({tunedNodes.length} Servers)
                    </h4>
                    <div className="flex flex-wrap gap-2.5 max-h-36 overflow-y-auto p-2 bg-bg-panel/20 rounded-xl border border-border-main/20">
                      {tunedNodes.map((n) => {
                        const isExcluded = excludedNodeIds.includes(n.machineId);
                        const isActive = n.machineId === activeTunedNodeId;
                        return (
                          <div
                            key={n.machineId}
                            className={`flex items-center space-x-2 px-3 py-1.5 rounded-lg text-sm font-bold transition border ${isActive
                                ? 'bg-emerald-500/10 border-emerald-500 text-emerald-600 dark:text-emerald-600 dark:text-emerald-400 shadow-md shadow-emerald-500/10'
                                : 'bg-bg-panel border-border-main text-text-muted hover:text-text-main'
                              } ${isExcluded ? 'opacity-40' : ''}`}
                          >
                            <button
                              onClick={() => setActiveTunedNodeId(n.machineId)}
                              className="flex items-center space-x-1.5 cursor-pointer focus:outline-none"
                            >
                              <HardDrive className="w-4 h-4" />
                              <span className={isExcluded ? 'line-through' : ''}>{n.hostname}</span>
                            </button>
                            <button
                              onClick={(e) => {
                                e.stopPropagation();
                                handleToggleNodeExclude(n.machineId);
                              }}
                              className={`ml-1 px-1.5 py-0.5 rounded text-sm cursor-pointer transition ${isExcluded
                                  ? 'bg-emerald-500/20 text-emerald-600 dark:text-emerald-400 hover:bg-emerald-500/30'
                                  : 'bg-red-500/15 text-red-400 hover:bg-red-500/25'
                                }`}
                              title={isExcluded ? 'Include server in recommendation' : 'Exclude server from recommendation'}
                            >
                              {isExcluded ? 'Include' : 'Exclude'}
                            </button>
                          </div>
                        );
                      })}
                    </div>
                  </div>

                  {/* Row 3: Active Server Spec Details */}
                  {activeNode && (
                    <div className="space-y-3 p-5 bg-bg-input/40 border border-border-main/50 rounded-xl">
                      <div className="flex flex-col md:flex-row md:items-center md:justify-start md:space-x-3 border-b border-border-main/20 pb-2">
                        <h4 className="text-sm font-bold text-emerald-600 dark:text-emerald-400 flex items-center">
                          <Server className="w-4 h-4 mr-1.5 text-emerald-600 dark:text-emerald-400" />
                          Server Details ({activeNode.hostname})
                        </h4>
                        <span className="text-sm text-text-muted font-mono bg-bg-panel px-2.5 py-0.5 rounded-md border border-border-main/30 mt-1 md:mt-0">
                          Machine ID: <span className="text-text-main font-bold">{activeNode.machineId}</span>
                        </span>
                      </div>
                      <div className="space-y-5 mt-2">
                        {/* Upper Row: Side-by-Side HW and OS specs */}
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

                          {/* 1. Server HW Spec */}
                          <div className="bg-bg-panel/30 border border-border-main/30 rounded-xl p-4 space-y-2.5 text-sm">
                            <span className="text-sm font-bold text-text-muted block mb-1 border-b border-border-main/20 pb-1">Server HW Spec</span>
                            <div className="flex justify-between items-center">
                              <span className="text-text-muted font-normal shrink-0">CPU Model:</span>
                              <span className="text-text-main font-bold font-sans text-right text-sm" title={`${activeNode.cpu.model} (${activeNode.cpu.vendor})`}>
                                {activeNode.cpu.model} ({activeNode.cpu.vendor})
                              </span>
                            </div>
                            <div className="flex justify-between">
                              <span className="text-text-muted font-normal">Architecture:</span>
                              <span className="text-text-main font-bold font-mono">{activeNode.cpu.architecture}</span>
                            </div>
                            <div className="flex justify-between">
                              <span className="text-text-muted font-normal">CPUs (Sockets):</span>
                              <span className="text-text-main font-bold">{activeNode.cpu.cpus} cpus</span>
                            </div>
                            <div className="flex justify-between">
                              <span className="text-text-muted font-normal">Cores per CPU:</span>
                              <span className="text-text-main font-bold">{activeNode.cpu.cores} cores</span>
                            </div>
                            <div className="flex justify-between">
                              <span className="text-text-muted font-normal">Threads:</span>
                              <span className="text-text-main font-bold">{activeNode.cpu.threads} threads</span>
                            </div>
                            <div className="flex justify-between">
                              <span className="text-text-muted font-normal">Memory RAM:</span>
                              <span className="text-text-main font-bold">
                                {activeNode.memory.totalSize < 1000000
                                  ? activeNode.memory.totalSize.toFixed(1)
                                  : (activeNode.memory.totalSize / (1024 * 1024 * 1024)).toFixed(1)} GB
                              </span>
                            </div>
                            <div className="flex justify-between pt-1 border-t border-border-main/20 mt-1">
                              <span className="text-text-muted font-normal">Root Disk:</span>
                              <span className="text-text-main font-bold">
                                {activeNode.rootDisk.totalSize < 1000000
                                  ? activeNode.rootDisk.totalSize.toFixed(0)
                                  : (activeNode.rootDisk.totalSize / (1024 * 1024 * 1024)).toFixed(0)} GB ({activeNode.rootDisk.type})
                              </span>
                            </div>
                            <div className="flex justify-between">
                              <span className="text-text-muted font-normal">Data Disk:</span>
                              {activeNode.dataDisks && activeNode.dataDisks.length > 0 ? (
                                <span className="text-text-main font-bold">
                                  {activeNode.dataDisks.map((d) => {
                                    const sz = d.totalSize < 1000000
                                      ? d.totalSize.toFixed(0)
                                      : (d.totalSize / (1024 * 1024 * 1024)).toFixed(0);
                                    return `${sz}GB`;
                                  }).join(', ')} ({activeNode.dataDisks.length} disks)
                                </span>
                              ) : (
                                <span className="text-text-muted font-normal italic">None</span>
                              )}
                            </div>
                          </div>

                          {/* 2. Operating System */}
                          <div className="bg-bg-panel/30 border border-border-main/30 rounded-xl p-4 space-y-2.5 text-sm">
                            <span className="text-sm font-bold text-text-muted block mb-1 border-b border-border-main/20 pb-1">Operating System</span>
                            <div className="flex justify-between">
                              <span className="text-text-muted font-normal">Name:</span>
                              <span className="text-text-main font-bold">{activeNode.os.name}</span>
                            </div>
                            <div className="flex justify-between">
                              <span className="text-text-muted font-normal">Version:</span>
                              <span className="text-text-main font-bold">{activeNode.os.version}</span>
                            </div>
                            {activeNode.os.prettyName && (
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">Pretty Name:</span>
                                <span className="text-text-main font-bold">{activeNode.os.prettyName}</span>
                              </div>
                            )}
                            {activeNode.os.versionId && (
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">Version ID:</span>
                                <span className="text-text-main font-bold font-mono">{activeNode.os.versionId}</span>
                              </div>
                            )}
                            {activeNode.os.versionCodename && (
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">Codename:</span>
                                <span className="text-text-main font-bold font-mono capitalize">{activeNode.os.versionCodename}</span>
                              </div>
                            )}
                            {activeNode.os.idLike && (
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">Base OS Like:</span>
                                <span className="text-text-main font-bold font-mono uppercase">{activeNode.os.idLike}</span>
                              </div>
                            )}
                          </div>
                        </div>

                        {/* Middle Row: Network Interfaces (Full-Width) */}
                        <div className="bg-bg-panel/30 border border-border-main/30 rounded-xl p-4 space-y-2.5 text-sm">
                          <span className="text-sm font-bold text-text-muted block mb-1 border-b border-border-main/20 pb-1">Network Interfaces</span>
                          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
                            {activeNode.interfaces?.map((iface, idx) => (
                              <div key={idx} className="bg-bg-panel/60 border border-border-main/25 p-3 rounded-lg space-y-1 font-mono text-sm">
                                <div className="flex justify-between text-text-muted font-normal">
                                  <span>{iface.name}</span>
                                  <span className="text-emerald-600 dark:text-emerald-400 uppercase text-sm font-extrabold">{iface.state}</span>
                                </div>
                                {iface.ipv4CidrBlocks && iface.ipv4CidrBlocks.length > 0 && (
                                  <div className="text-text-muted text-sm">
                                    <span className="font-normal">IPv4: </span>
                                    <span className="text-text-main font-bold">{iface.ipv4CidrBlocks.join(', ')}</span>
                                  </div>
                                )}
                              </div>
                            ))}
                            {(!activeNode.interfaces || activeNode.interfaces.length === 0) && (
                              <div className="text-text-muted italic text-center py-2 col-span-3">No interfaces configured.</div>
                            )}
                          </div>
                        </div>

                        {/* Lower Row: Firewall Rules (Full-Width & Editable) */}
                        <div className="bg-bg-panel/30 border border-border-main/30 rounded-xl p-4 space-y-2.5 text-sm">
                          <span className="text-sm font-bold text-text-muted block mb-1 border-b border-border-main/20 pb-1">Firewall Rules</span>
                          {activeNode.firewallTable && activeNode.firewallTable.length > 0 ? (
                            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 max-h-56 overflow-y-auto pr-1">
                              {activeNode.firewallTable.map((rule, idx) => (
                                <div key={idx} className="flex justify-between items-center text-sm bg-bg-panel p-2 rounded-lg border border-border-main/35 font-mono text-text-muted font-normal">
                                  <span className={`uppercase font-extrabold text-xs px-2 py-0.5 rounded ${rule.direction === 'inbound' ? 'bg-green-500/10 text-green-600 dark:text-green-400 border border-green-500/20' : 'bg-blue-500/10 text-blue-600 dark:text-blue-400 border border-blue-500/20'}`}>
                                    {rule.direction === 'inbound' ? 'Inbound' : 'Outbound'}
                                  </span>
                                  <span>{(rule.protocol || '').toUpperCase()}</span>
                                  <span>Port: {rule.dstPorts}</span>
                                  <span className="truncate max-w-[120px]" title={rule.srcCIDR}>{rule.srcCIDR}</span>
                                  <button
                                    onClick={() => handleDeleteFirewallRule(activeNode.machineId, idx)}
                                    className="text-red-400 hover:text-red-300 font-bold ml-2 px-1 cursor-pointer"
                                    title="Delete rule"
                                  >
                                    ✕
                                  </button>
                                </div>
                              ))}
                            </div>
                          ) : (
                            <div className="text-text-muted italic text-center py-4">No firewall rules configured.</div>
                          )}

                          {/* Firewall Rule Creator Form */}
                          <div className="pt-3 border-t border-border-main/20 mt-3 space-y-2">
                            <span className="text-text-muted font-bold text-sm block">Add Custom Firewall Rule</span>
                            <div className="grid grid-cols-2 md:grid-cols-5 gap-2 items-end bg-bg-panel/40 p-3 rounded-xl border border-border-main/20">
                              <div>
                                <label className="block text-sm font-semibold text-text-muted mb-1">Direction</label>
                                <select
                                  value={newRuleDir}
                                  onChange={(e) => setNewRuleDir(e.target.value)}
                                  className="w-full bg-bg-input border border-border-main/60 text-text-main rounded-lg px-2 py-1 text-sm focus:outline-none"
                                >
                                  <option value="inbound">Inbound</option>
                                  <option value="outbound">Outbound</option>
                                </select>
                              </div>
                              <div>
                                <label className="block text-sm font-semibold text-text-muted mb-1">Protocol</label>
                                <select
                                  value={newRuleProto}
                                  onChange={(e) => setNewRuleProto(e.target.value)}
                                  className="w-full bg-bg-input border border-border-main/60 text-text-main rounded-lg px-2 py-1 text-sm focus:outline-none"
                                >
                                  <option value="tcp">TCP</option>
                                  <option value="udp">UDP</option>
                                  <option value="icmp">ICMP</option>
                                  <option value="*">* (All)</option>
                                </select>
                              </div>
                              <div>
                                <label className="block text-sm font-semibold text-text-muted mb-1">Ports</label>
                                <input
                                  type="text"
                                  value={newRulePort}
                                  onChange={(e) => setNewRulePort(e.target.value)}
                                  placeholder="e.g., 80 or *"
                                  className="w-full bg-bg-input border border-border-main/60 text-text-main rounded-lg px-2 py-1 text-sm focus:outline-none"
                                />
                              </div>
                              <div>
                                <label className="block text-sm font-semibold text-text-muted mb-1">Source CIDR</label>
                                <input
                                  type="text"
                                  value={newRuleCidr}
                                  onChange={(e) => setNewRuleCidr(e.target.value)}
                                  placeholder="e.g., 0.0.0.0/0"
                                  className="w-full bg-bg-input border border-border-main/60 text-text-main rounded-lg px-2 py-1 text-sm focus:outline-none font-mono"
                                />
                              </div>
                              <div className="col-span-2 md:col-span-1">
                                <button
                                  onClick={() => handleAddFirewallRule(activeNode.machineId)}
                                  className="w-full bg-emerald-500 hover:bg-emerald-600 text-slate-950 font-bold py-1 text-sm rounded-lg cursor-pointer transition text-center"
                                >
                                  Add Rule
                                </button>
                              </div>
                            </div>
                          </div>
                        </div>

                      </div>
                    </div>
                  )}

                  {/* Row 4: Load Balancer (NLB) */}
                  <div className="space-y-3 p-5 bg-bg-input/40 border border-border-main/50 rounded-xl">
                    <h4 className="text-sm font-bold text-emerald-600 dark:text-emerald-400 flex items-center">
                      <Cpu className="w-4 h-4 mr-1.5 text-emerald-600 dark:text-emerald-400" />
                      4. Network Load Balancer (NLB)
                    </h4>
                    {selectedSourceModel.onpremiseInfraModel.nlbs && selectedSourceModel.onpremiseInfraModel.nlbs.length > 0 ? (
                      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                        {selectedSourceModel.onpremiseInfraModel.nlbs.map((nlb, idx) => (
                          <div key={idx} className="bg-bg-panel border border-border-main/40 rounded-xl p-4.5 space-y-3.5 leading-relaxed text-sm">
                            <div className="flex justify-between items-center border-b border-border-main/40 pb-2">
                              <span className="font-extrabold text-text-main text-base flex items-center">
                                <span className="w-2 h-2 bg-teal-500 rounded-full mr-2 animate-pulse" />
                                Haproxy NLB
                              </span>
                              <span className="text-sm font-bold px-2.5 py-1 bg-teal-100 dark:bg-teal-950/40 text-teal-600 dark:text-teal-400 rounded-full border border-teal-200 dark:border-teal-800/40 uppercase">
                                {nlb.listener.protocol} Mode
                              </span>
                            </div>

                            <div className="space-y-3 font-mono text-sm text-text-muted font-normal">
                              {/* Traffic Ingress */}
                              <div className="flex items-center space-x-2 text-text-muted">
                                <span className="font-sans font-normal shrink-0">Traffic Ingress ➔</span>
                                <span className="text-text-main font-bold">
                                  {nlb.listener.protocol.toUpperCase()} {nlb.listener.bindAddress}:{nlb.listener.port}
                                </span>
                              </div>

                              {/* Balancing Method */}
                              <div className="pl-4 border-l-2 border-teal-500/30 py-0.5 space-y-2">
                                <div className="text-sm text-teal-600 dark:text-teal-400 font-bold font-sans">
                                  ▼ Load Balancing: {nlb.backend.balance}
                                </div>
                                <div className="text-text-muted">
                                  Backend Group: <span className="text-text-main font-bold">{nlb.backend.name}</span>
                                </div>
                                {/* Target VMs */}
                                <div className="space-y-1.5 pl-3.5 mt-1 border-l border-border-main/50">
                                  {nlb.backend.servers.map((srv, sIdx) => (
                                    <div key={sIdx} className="flex justify-between items-center text-sm">
                                      <span>├─ {srv.name}</span>
                                      <span className="text-text-main font-bold">{srv.ip}:{srv.port}</span>
                                    </div>
                                  ))}
                                </div>
                              </div>

                              {/* Health Check */}
                              {nlb.healthCheck?.enabled && (
                                <div className="pt-2 border-t border-border-main/20 flex justify-between items-center text-sm text-text-muted font-sans font-semibold">
                                  <span className="flex items-center text-emerald-600 dark:text-emerald-400">
                                    <span className="w-1.5 h-1.5 bg-emerald-500 rounded-full mr-1.5" />
                                    Health Check Active
                                  </span>
                                  <span>Interval: {nlb.healthCheck.interval}s (Thresh: {nlb.healthCheck.threshold})</span>
                                </div>
                              )}
                            </div>
                          </div>
                        ))}
                      </div>
                    ) : (
                      <div className="text-sm text-text-muted italic p-2 bg-bg-panel/20 rounded-lg border border-dashed border-border-main/50 text-center">
                        No Network Load Balancer detected.
                      </div>
                    )}
                  </div>

                  {/* Save spec and proceed button at the bottom of Step 1 */}
                  <div className="flex flex-row items-center justify-start pt-4 border-t border-border-main/20 mt-4 space-x-4">
                    <button
                      onClick={() => setShowSaveModal(true)}
                      className="px-6 py-3 bg-gradient-to-r from-emerald-500 to-blue-600 hover:from-emerald-600 hover:to-blue-700 text-slate-950 rounded-xl text-sm font-extrabold flex items-center transition cursor-pointer shadow-lg shadow-emerald-500/10 shrink-0"
                    >
                      <Save className="w-4 h-4 mr-1.5" /> Save Source Infra Revision
                    </button>
                    <div className="flex items-center space-x-2 text-sm text-text-muted">
                      <span className="font-bold">Model to save:</span>
                      <span className="text-emerald-600 dark:text-emerald-600 dark:text-emerald-400 font-extrabold text-sm">{selectedSourceModel.name}</span>
                      <span className="text-sm text-text-muted font-mono bg-bg-panel px-1.5 py-0.5 rounded border border-border-main/40">
                        v{selectedSourceModel.version || '1.0'}
                      </span>
                    </div>
                  </div>
                </div>
              )}
            </div>
          </div>
        )}
      </div>

      <SaveRevisionModal
        isOpen={showSaveModal}
        onClose={() => setShowSaveModal(false)}
        title="Save Source Infra Revision"
        defaultName={selectedSourceModel?.name || ''}
        defaultDescription={selectedSourceModel?.description || ''}
        defaultVersion={selectedSourceModel?.version || '1.0.0'}
        existingRevisions={savedSourceModels
          .filter(m => m.id !== 'sample-source-infra-1')
          .map(m => ({ id: m.id, name: m.name, version: m.version }))}
        onSave={handleSaveToDamselfly}
      />

      {/* Delete Confirmation Modal */}
      {showDeleteConfirm && selectedSourceModel && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up">
            <div className="flex justify-between items-center mb-4">
              <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                <Trash2 className="w-4 h-4 text-red-500" /> Delete Model
              </h3>
              <button
                onClick={() => { setShowDeleteConfirm(false); setDeleteError(''); setDeleteConfirmText(''); }}
                disabled={isDeleting}
                className="text-text-muted hover:text-text-main transition p-1 hover:bg-bg-input rounded-lg cursor-pointer disabled:opacity-50"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            <div className="space-y-4">
              <p className="text-sm text-text-muted leading-relaxed">
                Are you sure you want to delete the model <strong className="text-text-main">"{selectedSourceModel.name}"</strong>? This action cannot be undone.
              </p>

              <div className="space-y-1.5 pt-1">
                <label className="block text-xs font-bold text-text-muted">
                  To confirm, type <span className="font-mono bg-bg-panel px-1 py-0.5 rounded border border-border-main/60 text-text-main select-all">{selectedSourceModel.name}</span> in the box below:
                </label>
                <input
                  type="text"
                  value={deleteConfirmText}
                  onChange={(e) => setDeleteConfirmText(e.target.value)}
                  placeholder="Type the model name to delete"
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-1 focus:ring-red-500 font-bold font-mono"
                  disabled={isDeleting}
                />
              </div>

              {deleteError && (
                <div className="flex items-center gap-2 bg-red-500/10 text-red-500 px-4 py-3 rounded-xl text-xs font-semibold border border-red-500/20">
                  <span>{deleteError}</span>
                </div>
              )}

              <div className="flex justify-end gap-3 pt-2">
                <button
                  onClick={() => { setShowDeleteConfirm(false); setDeleteError(''); setDeleteConfirmText(''); }}
                  disabled={isDeleting}
                  className="px-4 py-2 bg-bg-panel border border-border-main text-text-main rounded-xl text-sm font-semibold hover:bg-bg-input transition cursor-pointer disabled:opacity-50"
                >
                  Cancel
                </button>
                <button
                  onClick={async () => {
                    setIsDeleting(true);
                    setDeleteError('');
                    try {
                      await deleteSourceModel(selectedSourceModel.id);
                      setShowDeleteConfirm(false);
                      setDeleteConfirmText('');
                      setIsModelLoaded(false);
                      setTunedNodes([]);
                      setActiveStep(1);
                    } catch (err: any) {
                      setDeleteError(err.message || 'Failed to delete model');
                    } finally {
                      setIsDeleting(false);
                    }
                  }}
                  disabled={isDeleting || deleteConfirmText !== selectedSourceModel.name}
                  className={`px-4 py-2 rounded-xl text-sm font-semibold transition flex items-center gap-1.5 ${
                    isDeleting || deleteConfirmText !== selectedSourceModel.name
                      ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                      : 'bg-red-500 hover:bg-red-600 text-white cursor-pointer shadow-md shadow-red-500/20 animate-pulse'
                  }`}
                >
                  {isDeleting && <Loader2 className="w-4 h-4 animate-spin" />}
                  Confirm Delete
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Toast Notification */}
      {tuningSourceSaveSuccess && (
        <div className="fixed bottom-6 right-6 z-50 flex items-center gap-2.5 bg-slate-950/95 border border-emerald-500/40 text-emerald-600 dark:text-emerald-400 px-5 py-4.5 rounded-2xl shadow-2xl shadow-emerald-500/10 animate-fade-in font-bold text-sm backdrop-blur-md">
          <CheckCircle2 className="w-5 h-5 text-emerald-500" />
          <span>Source specification updated and saved. Ready to define target Cloud settings.</span>
        </div>
      )}

    </div>
  );
};
