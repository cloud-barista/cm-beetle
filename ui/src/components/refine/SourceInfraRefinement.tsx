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
  ArrowRight, ArrowLeft,
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

export const SourceInfraRefinement: React.FC<{ onNext?: () => void; onBack?: () => void }> = ({ onNext, onBack }) => {
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

  // Node spec editing state
  const [isEditingActiveNode, setIsEditingActiveNode] = useState(false);

  // Network Config editing state
  const [isEditingNetworkConfig, setIsEditingNetworkConfig] = useState(false);
  const [newGwIp, setNewGwIp] = useState('');
  const [newGwIface, setNewGwIface] = useState('eth0');

  // NLB Config editing state
  const [isEditingNlbConfig, setIsEditingNlbConfig] = useState(false);
  const [tunedNlbs, setTunedNlbs] = useState<any[]>([]);
  const [newNlbPort, setNewNlbPort] = useState('80');
  const [newNlbProtocol, setNewNlbProtocol] = useState('tcp');
  const [newNlbBackendName, setNewNlbBackendName] = useState('backend_web');
  const [newNlbServerIp, setNewNlbServerIp] = useState('');
  const [newNlbServerPort, setNewNlbServerPort] = useState('80');

  useEffect(() => {
    if (selectedSourceModel?.onpremiseInfraModel?.nlbs) {
      setTunedNlbs(JSON.parse(JSON.stringify(selectedSourceModel.onpremiseInfraModel.nlbs)));
    } else {
      setTunedNlbs([]);
    }
  }, [selectedSourceModel, isModelLoaded]);

  // New Data Disk Form state
  const [newDataDiskLabel, setNewDataDiskLabel] = useState('');
  const [newDataDiskSize, setNewDataDiskSize] = useState('100');
  const [newDataDiskType, setNewDataDiskType] = useState('SSD');

  // New Network Interface Form state
  const [newIfaceName, setNewIfaceName] = useState('');
  const [newIfaceCidr, setNewIfaceCidr] = useState('');
  const [newIfaceState, setNewIfaceState] = useState('up');
  const [newIfaceSpeed, setNewIfaceSpeed] = useState('1000');
  const [newIfaceSpeedUnit, setNewIfaceSpeedUnit] = useState('Mbps');

  // Unit helper functions conforming to Cloud-Barista standards (Memory: GiB, Disk: GB)
  const toGiB = (val: number): number => {
    if (!val) return 0;
    return val > 1000000 ? Math.round(val / (1024 * 1024 * 1024)) : Math.round(val);
  };

  const toGB = (val: number): number => {
    if (!val) return 0;
    return val > 1000000 ? Math.round(val / (1024 * 1024 * 1024)) : Math.round(val);
  };

  const handleUpdateNodeCpu = (machineId: string, key: string, value: any) => {
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      return {
        ...n,
        cpu: {
          ...n.cpu,
          [key]: ['cpus', 'cores', 'threads'].includes(key) ? Math.max(1, parseInt(value, 10) || 1) : value
        }
      };
    }));
  };

  const handleUpdateNodeMemory = (machineId: string, sizeInGiB: number) => {
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      return {
        ...n,
        memory: {
          ...n.memory,
          totalSize: Math.max(1, sizeInGiB)
        }
      };
    }));
  };

  const handleUpdateNodeRootDisk = (machineId: string, key: string, value: any) => {
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      return {
        ...n,
        rootDisk: {
          ...n.rootDisk,
          [key]: key === 'totalSize' ? Math.max(1, parseInt(value, 10) || 1) : value
        }
      };
    }));
  };

  const handleAddDataDisk = (machineId: string) => {
    const size = parseInt(newDataDiskSize, 10);
    if (!size || isNaN(size)) return;
    const diskLabel = newDataDiskLabel || `data-disk-${((activeNode?.dataDisks || []).length) + 1}`;
    const newDisk = {
      label: diskLabel,
      totalSize: size,
      type: newDataDiskType
    };
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      return {
        ...n,
        dataDisks: [...(n.dataDisks || []), newDisk]
      };
    }));
    setNewDataDiskLabel('');
    setNewDataDiskSize('100');
  };

  const handleDeleteDataDisk = (machineId: string, diskIdx: number) => {
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      return {
        ...n,
        dataDisks: (n.dataDisks || []).filter((_, idx) => idx !== diskIdx)
      };
    }));
  };

  const handleUpdateNodeOs = (machineId: string, key: string, value: string) => {
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      return {
        ...n,
        os: {
          ...n.os,
          [key]: value
        }
      };
    }));
  };

  const handleUpdateHostname = (machineId: string, hostname: string) => {
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      return {
        ...n,
        hostname
      };
    }));
  };

  const handleAddNetworkInterface = (machineId: string) => {
    if (!newIfaceName) return;
    const newIface = {
      name: newIfaceName,
      state: newIfaceState,
      mtu: 1500,
      ipv4CidrBlocks: newIfaceCidr ? [newIfaceCidr] : [],
      ipv6CidrBlocks: []
    };
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      return {
        ...n,
        interfaces: [...(n.interfaces || []), newIface]
      };
    }));
    setNewIfaceName('');
    setNewIfaceCidr('');
  };

  const handleDeleteNetworkInterface = (machineId: string, ifaceIdx: number) => {
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      return {
        ...n,
        interfaces: (n.interfaces || []).filter((_, idx) => idx !== ifaceIdx)
      };
    }));
  };

  const handleAddGateway = () => {
    if (!newGwIp || !tunedNetwork) return;
    const currentGws = tunedNetwork.ipv4Networks?.defaultGateways || [];
    const newGw = {
      ip: newGwIp,
      interfaceName: newGwIface || 'eth0',
      machineId: activeTunedNodeId || 'node-gw'
    };
    const updatedNetwork = {
      ...tunedNetwork,
      ipv4Networks: {
        ...tunedNetwork.ipv4Networks,
        defaultGateways: [...currentGws, newGw]
      }
    };
    setTunedNetwork(updatedNetwork);
    setNewGwIp('');
  };

  const handleDeleteGateway = (gwIdx: number) => {
    if (!tunedNetwork) return;
    const currentGws = tunedNetwork.ipv4Networks?.defaultGateways || [];
    const updatedNetwork = {
      ...tunedNetwork,
      ipv4Networks: {
        ...tunedNetwork.ipv4Networks,
        defaultGateways: currentGws.filter((_: any, idx: number) => idx !== gwIdx)
      }
    };
    setTunedNetwork(updatedNetwork);
  };

  const handleUpdateNlbListenerPort = (nlbIdx: number, port: number) => {
    setTunedNlbs(prev => prev.map((nlb, idx) => {
      if (idx !== nlbIdx) return nlb;
      return {
        ...nlb,
        listener: {
          ...nlb.listener,
          port
        }
      };
    }));
  };

  const handleAddNlbServer = (nlbIdx: number) => {
    if (!newNlbServerIp) return;
    const srvPort = parseInt(newNlbServerPort, 10) || 80;
    const srvName = `srv-${newNlbServerIp.replaceAll('.', '-')}`;
    const newSrv = {
      name: srvName,
      ip: newNlbServerIp,
      port: srvPort
    };
    setTunedNlbs(prev => prev.map((nlb, idx) => {
      if (idx !== nlbIdx) return nlb;
      return {
        ...nlb,
        backend: {
          ...nlb.backend,
          servers: [...(nlb.backend?.servers || []), newSrv]
        }
      };
    }));
    setNewNlbServerIp('');
  };

  const handleDeleteNlbServer = (nlbIdx: number, srvIdx: number) => {
    setTunedNlbs(prev => prev.map((nlb, idx) => {
      if (idx !== nlbIdx) return nlb;
      return {
        ...nlb,
        backend: {
          ...nlb.backend,
          servers: (nlb.backend?.servers || []).filter((_: any, sIdx: number) => sIdx !== srvIdx)
        }
      };
    }));
  };

  const handleAddNlbInstance = () => {
    const port = parseInt(newNlbPort, 10) || 80;
    const newNlb = {
      software: 'haproxy',
      listener: {
        bindAddress: '*',
        port: port,
        protocol: newNlbProtocol
      },
      backend: {
        name: newNlbBackendName || 'backend_web',
        balance: 'roundrobin',
        protocol: newNlbProtocol,
        servers: []
      },
      healthCheck: {
        enabled: true,
        protocol: newNlbProtocol,
        port: port,
        interval: 10,
        timeout: 10,
        threshold: 3
      }
    };
    setTunedNlbs(prev => [...prev, newNlb]);
  };

  const handleDeleteNlbInstance = (nlbIdx: number) => {
    setTunedNlbs(prev => prev.filter((_, idx) => idx !== nlbIdx));
  };

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
      network: tunedNetwork || selectedSourceModel.onpremiseInfraModel.network,
      nlbs: tunedNlbs
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
    <div className="space-y-6 mx-auto pb-24">

      {/* Top Banner Description Box */}
      <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
        <div className="flex items-center gap-2 shrink-0">
          <Sliders className="w-5 h-5 text-emerald-500" />
          <h2 className="text-base font-extrabold text-text-main tracking-tight">
            Source Infra Refinement
          </h2>
        </div>
        <span className="text-sm text-text-muted">
          Review & refine extracted source metadata, adjust node specifications, configure network settings, and set migration parameters.
        </span>
      </div>

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
                    <div className="flex flex-col md:flex-row md:items-center justify-between border-b border-border-main/20 pb-2.5 gap-2">
                      <h4 className="text-sm font-bold text-emerald-600 dark:text-emerald-400 flex items-center">
                        <Network className="w-4 h-4 mr-1.5 text-emerald-500" />
                        1. Network Configuration
                      </h4>
                      <button
                        type="button"
                        onClick={() => setIsEditingNetworkConfig(!isEditingNetworkConfig)}
                        className={`px-3 py-1.5 rounded-xl text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer whitespace-nowrap shrink-0 shadow-md shadow-emerald-500/20 ${
                          isEditingNetworkConfig
                            ? 'bg-emerald-700 hover:bg-emerald-800 text-white border border-emerald-500/40'
                            : 'bg-emerald-600 hover:bg-emerald-700 text-white border border-emerald-500/40'
                        }`}
                      >
                        <Sliders className="w-3.5 h-3.5 text-white" />
                        <span>{isEditingNetworkConfig ? 'Done Editing Network' : 'Edit Network Config'}</span>
                      </button>
                    </div>

                    <div className="grid grid-cols-1 md:grid-cols-2 gap-6 text-sm mt-2">
                      {/* Gateways section */}
                      <div className="bg-bg-panel/20 p-3 rounded-lg border border-border-main/20 space-y-2">
                        <span className="text-text-muted font-semibold block mb-1">Gateways for VNet/Subnet estimation:</span>
                        <div className="space-y-1.5">
                          {(tunedNetwork?.ipv4Networks?.defaultGateways || []).map((gw: any, idx: number) => (
                            <div key={idx} className="flex justify-between items-center bg-bg-panel px-2.5 py-1.5 rounded-lg border border-border-main/30 font-mono text-sm">
                              <span className="text-text-muted font-normal">{gw.interfaceName}</span>
                              <div className="flex items-center gap-2">
                                <span className="text-text-main font-bold">{gw.ip}</span>
                                {isEditingNetworkConfig && (
                                  <button
                                    type="button"
                                    onClick={() => handleDeleteGateway(idx)}
                                    className="text-red-400 hover:text-red-300 font-bold cursor-pointer"
                                  >
                                    ✕
                                  </button>
                                )}
                              </div>
                            </div>
                          ))}
                          {(tunedNetwork?.ipv4Networks?.defaultGateways || []).length === 0 && (
                            <div className="text-xs text-text-muted italic">No gateways configured.</div>
                          )}
                        </div>

                        {/* Add Gateway form in Edit mode */}
                        {isEditingNetworkConfig && (
                          <div className="pt-2 border-t border-border-main/20 flex gap-2 items-center">
                            <input
                              type="text"
                              placeholder="Gateway IP (10.0.0.1)"
                              value={newGwIp}
                              onChange={(e) => setNewGwIp(e.target.value)}
                              className="flex-1 bg-bg-panel border border-border-main/60 rounded-lg px-2.5 py-1 text-xs text-text-main font-mono"
                            />
                            <input
                              type="text"
                              placeholder="Interface (eth0)"
                              value={newGwIface}
                              onChange={(e) => setNewGwIface(e.target.value)}
                              className="w-24 bg-bg-panel border border-border-main/60 rounded-lg px-2.5 py-1 text-xs text-text-main"
                            />
                            <button
                              type="button"
                              onClick={handleAddGateway}
                              className="px-3 py-1 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg text-xs font-bold transition cursor-pointer whitespace-nowrap"
                            >
                              Add GW
                            </button>
                          </div>
                        )}
                      </div>

                      {/* CIDRs section */}
                      {tunedNetwork?.ipv4Networks && (
                        <div className="bg-bg-panel/20 p-3 rounded-lg border border-border-main/20 space-y-2">
                          <span className="text-text-muted font-semibold block">Source Network CIDR Block:</span>
                          <div className="flex flex-wrap gap-1.5 pt-0.5">
                            {(tunedNetwork.ipv4Networks.cidrBlocks || []).map((cidr: string, idx: number) => (
                              <span key={idx} className="bg-emerald-500/10 border border-emerald-500/25 text-emerald-600 dark:text-emerald-400 font-mono text-sm px-2 py-1 rounded-md font-extrabold flex items-center space-x-1.5 animate-fade-in">
                                <span>{cidr}</span>
                                <button
                                  type="button"
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
                              type="button"
                              onClick={() => {
                                handleAddCidr(newCidr);
                                setNewCidr('');
                              }}
                              className="bg-emerald-600 hover:bg-emerald-700 text-white font-bold px-4 py-1.5 rounded-lg text-sm cursor-pointer transition whitespace-nowrap"
                            >
                              Add CIDR
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
                      <div className="flex flex-col md:flex-row md:items-center justify-between border-b border-border-main/20 pb-2.5 gap-2">
                        <div className="flex items-center space-x-3">
                          <h4 className="text-sm font-bold text-emerald-600 dark:text-emerald-400 flex items-center">
                            <Server className="w-4 h-4 mr-1.5 text-emerald-500" />
                            Server Details ({activeNode.hostname})
                          </h4>
                          <span className="text-xs text-text-muted font-mono bg-bg-panel px-2.5 py-0.5 rounded-md border border-border-main/30">
                            Machine ID: <span className="text-text-main font-bold">{activeNode.machineId}</span>
                          </span>
                        </div>
                        <button
                          type="button"
                          onClick={() => setIsEditingActiveNode(!isEditingActiveNode)}
                          className={`px-3.5 py-2 rounded-xl text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer whitespace-nowrap shrink-0 shadow-md shadow-emerald-500/20 ${
                            isEditingActiveNode
                              ? 'bg-emerald-700 hover:bg-emerald-800 text-white border border-emerald-500/40'
                              : 'bg-emerald-600 hover:bg-emerald-700 text-white border border-emerald-500/40'
                          }`}
                        >
                          <Sliders className="w-3.5 h-3.5 text-white" />
                          <span>{isEditingActiveNode ? 'Done Editing Spec' : 'Edit Server Spec'}</span>
                        </button>
                      </div>

                      <div className="space-y-5 mt-2">
                        {isEditingActiveNode ? (
                          /* EDIT MODE FORM UI */
                          <div className="space-y-5 animate-fade-in">
                            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

                              {/* 1. Server HW Spec Edit */}
                              <div className="bg-bg-panel/50 border border-emerald-500/30 rounded-xl p-4 space-y-3 text-sm">
                                <div className="flex justify-between items-center border-b border-border-main/20 pb-1.5">
                                  <span className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400">Server Hardware Spec (Edit)</span>
                                  <span className="text-xs text-text-muted font-mono">Standard Units: GiB / GB</span>
                                </div>
                                <div className="space-y-2">
                                  <div>
                                    <label className="block text-xs font-normal text-text-muted mb-0.5">Hostname</label>
                                    <input
                                      type="text"
                                      value={activeNode.hostname}
                                      onChange={(e) => handleUpdateHostname(activeNode.machineId, e.target.value)}
                                      className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2.5 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500"
                                    />
                                  </div>
                                  <div>
                                    <label className="block text-xs font-normal text-text-muted mb-0.5">CPU Model String</label>
                                    <input
                                      type="text"
                                      value={activeNode.cpu.model || ''}
                                      onChange={(e) => handleUpdateNodeCpu(activeNode.machineId, 'model', e.target.value)}
                                      className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2.5 py-1 text-xs text-text-main font-bold font-sans focus:ring-1 focus:ring-emerald-500"
                                    />
                                  </div>
                                  <div className="grid grid-cols-2 gap-2">
                                    <div>
                                      <label className="block text-xs font-normal text-text-muted mb-0.5">Architecture</label>
                                      <select
                                        value={activeNode.cpu.architecture || 'x86_64'}
                                        onChange={(e) => handleUpdateNodeCpu(activeNode.machineId, 'architecture', e.target.value)}
                                        className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                                      >
                                        <option value="x86_64">x86_64</option>
                                        <option value="arm64">arm64</option>
                                      </select>
                                    </div>
                                    <div>
                                      <label className="block text-xs font-normal text-text-muted mb-0.5">CPUs (Sockets)</label>
                                      <input
                                        type="number"
                                        min="1"
                                        value={activeNode.cpu.cpus}
                                        onChange={(e) => handleUpdateNodeCpu(activeNode.machineId, 'cpus', e.target.value)}
                                        className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500"
                                      />
                                    </div>
                                  </div>
                                  <div className="grid grid-cols-2 gap-2">
                                    <div>
                                      <label className="block text-xs font-normal text-text-muted mb-0.5">Cores per CPU</label>
                                      <input
                                        type="number"
                                        min="1"
                                        value={activeNode.cpu.cores}
                                        onChange={(e) => handleUpdateNodeCpu(activeNode.machineId, 'cores', e.target.value)}
                                        className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500"
                                      />
                                    </div>
                                    <div>
                                      <label className="block text-xs font-normal text-text-muted mb-0.5">Threads</label>
                                      <input
                                        type="number"
                                        min="1"
                                        value={activeNode.cpu.threads}
                                        onChange={(e) => handleUpdateNodeCpu(activeNode.machineId, 'threads', e.target.value)}
                                        className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500"
                                      />
                                    </div>
                                  </div>

                                  {/* Memory in GiB */}
                                  <div className="pt-2 border-t border-border-main/20">
                                    <label className="block text-xs font-normal text-text-muted mb-0.5">Memory RAM (GiB)</label>
                                    <div className="flex items-center gap-2">
                                      <input
                                        type="number"
                                        min="1"
                                        value={toGiB(activeNode.memory.totalSize)}
                                        onChange={(e) => handleUpdateNodeMemory(activeNode.machineId, Number(e.target.value))}
                                        className="flex-1 bg-bg-panel border border-border-main/60 rounded-lg px-2.5 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500"
                                      />
                                      <span className="text-xs font-extrabold text-emerald-600 dark:text-emerald-400">GiB</span>
                                    </div>
                                  </div>

                                  {/* Root Disk in GB */}
                                  <div className="pt-2 border-t border-border-main/20 grid grid-cols-2 gap-2">
                                    <div>
                                      <label className="block text-xs font-normal text-text-muted mb-0.5">Root Disk Size (GB)</label>
                                      <div className="flex items-center gap-1.5">
                                        <input
                                          type="number"
                                          min="1"
                                          value={toGB(activeNode.rootDisk.totalSize)}
                                          onChange={(e) => handleUpdateNodeRootDisk(activeNode.machineId, 'totalSize', e.target.value)}
                                          className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500"
                                        />
                                        <span className="text-xs font-extrabold text-emerald-600 dark:text-emerald-400">GB</span>
                                      </div>
                                    </div>
                                    <div>
                                      <label className="block text-xs font-normal text-text-muted mb-0.5">Root Disk Type</label>
                                      <select
                                        value={activeNode.rootDisk.type || 'SSD'}
                                        onChange={(e) => handleUpdateNodeRootDisk(activeNode.machineId, 'type', e.target.value)}
                                        className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                                      >
                                        <option value="SSD">SSD</option>
                                        <option value="HDD">HDD</option>
                                        <option value="NVMe">NVMe</option>
                                      </select>
                                    </div>
                                  </div>
                                </div>
                              </div>

                              {/* 2. OS Edit & Data Disk Management */}
                              <div className="space-y-4">
                                {/* OS Edit */}
                                <div className="bg-bg-panel/50 border border-emerald-500/30 rounded-xl p-4 space-y-2.5 text-sm">
                                  <span className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 block border-b border-border-main/20 pb-1.5">Operating System Info (Edit)</span>
                                  <div className="grid grid-cols-2 gap-2">
                                    <div>
                                      <label className="block text-xs font-normal text-text-muted mb-0.5">OS Name</label>
                                      <input
                                        type="text"
                                        value={activeNode.os.name || ''}
                                        onChange={(e) => handleUpdateNodeOs(activeNode.machineId, 'name', e.target.value)}
                                        className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500"
                                      />
                                    </div>
                                    <div>
                                      <label className="block text-xs font-normal text-text-muted mb-0.5">OS Version</label>
                                      <input
                                        type="text"
                                        value={activeNode.os.version || ''}
                                        onChange={(e) => handleUpdateNodeOs(activeNode.machineId, 'version', e.target.value)}
                                        className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500"
                                      />
                                    </div>
                                  </div>
                                  <div className="grid grid-cols-2 gap-2">
                                    <div>
                                      <label className="block text-xs font-normal text-text-muted mb-0.5">Version ID</label>
                                      <input
                                        type="text"
                                        value={activeNode.os.versionId || ''}
                                        onChange={(e) => handleUpdateNodeOs(activeNode.machineId, 'versionId', e.target.value)}
                                        className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs font-mono text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500"
                                      />
                                    </div>
                                    <div>
                                      <label className="block text-xs font-normal text-text-muted mb-0.5">Pretty Name</label>
                                      <input
                                        type="text"
                                        value={activeNode.os.prettyName || ''}
                                        onChange={(e) => handleUpdateNodeOs(activeNode.machineId, 'prettyName', e.target.value)}
                                        className="w-full bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-extrabold focus:ring-1 focus:ring-emerald-500"
                                      />
                                    </div>
                                  </div>
                                </div>

                                {/* Data Disks Management (GB) */}
                                <div className="bg-bg-panel/50 border border-emerald-500/30 rounded-xl p-4 space-y-2.5 text-sm">
                                  <span className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 block border-b border-border-main/20 pb-1.5">Data Disks (GB)</span>
                                  {activeNode.dataDisks && activeNode.dataDisks.length > 0 ? (
                                    <div className="space-y-1.5 max-h-32 overflow-y-auto">
                                      {activeNode.dataDisks.map((d, idx) => (
                                        <div key={idx} className="flex justify-between items-center bg-bg-panel px-3 py-1.5 rounded-lg border border-border-main/30 text-xs">
                                          <span className="font-extrabold text-text-main">{d.label}</span>
                                          <div className="flex items-center gap-2">
                                            <span className="font-extrabold text-emerald-600 dark:text-emerald-400">{toGB(d.totalSize)} GB ({d.type})</span>
                                            <button
                                              type="button"
                                              onClick={() => handleDeleteDataDisk(activeNode.machineId, idx)}
                                              className="text-red-400 hover:text-red-300 font-bold cursor-pointer"
                                            >
                                              ✕
                                            </button>
                                          </div>
                                        </div>
                                      ))}
                                    </div>
                                  ) : (
                                    <div className="text-xs text-text-muted italic">No data disks configured.</div>
                                  )}

                                  {/* Add Data Disk Form */}
                                  <div className="pt-2 border-t border-border-main/20 flex gap-2 items-center">
                                    <input
                                      type="text"
                                      placeholder="Disk Name"
                                      value={newDataDiskLabel}
                                      onChange={(e) => setNewDataDiskLabel(e.target.value)}
                                      className="flex-1 bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main"
                                    />
                                    <input
                                      type="number"
                                      min="1"
                                      placeholder="Size (GB)"
                                      value={newDataDiskSize}
                                      onChange={(e) => setNewDataDiskSize(e.target.value)}
                                      className="w-20 bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main"
                                    />
                                    <select
                                      value={newDataDiskType}
                                      onChange={(e) => setNewDataDiskType(e.target.value)}
                                      className="bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main cursor-pointer"
                                    >
                                      <option value="SSD">SSD</option>
                                      <option value="HDD">HDD</option>
                                      <option value="NVMe">NVMe</option>
                                    </select>
                                    <button
                                      type="button"
                                      onClick={() => handleAddDataDisk(activeNode.machineId)}
                                      className="px-3 py-1 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg text-xs font-bold transition cursor-pointer whitespace-nowrap"
                                    >
                                      Add Disk
                                    </button>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        ) : (
                          /* VIEW MODE CARD UI */
                          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

                            {/* 1. Server HW Spec */}
                            <div className="bg-bg-panel/30 border border-border-main/30 rounded-xl p-4 space-y-2.5 text-sm">
                              <span className="text-sm font-bold text-text-muted block mb-1 border-b border-border-main/20 pb-1">Server HW Spec</span>
                              <div className="flex justify-between items-center">
                                <span className="text-text-muted font-normal shrink-0">CPU Model:</span>
                                <span className="text-text-main font-extrabold font-sans text-right text-sm" title={`${activeNode.cpu.model} (${activeNode.cpu.vendor})`}>
                                  {activeNode.cpu.model || 'Unknown'} ({activeNode.cpu.vendor || 'Generic'})
                                </span>
                              </div>
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">Architecture:</span>
                                <span className="text-text-main font-extrabold font-mono">{activeNode.cpu.architecture || 'x86_64'}</span>
                              </div>
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">CPUs (Sockets):</span>
                                <span className="text-text-main font-extrabold">{activeNode.cpu.cpus} cpus</span>
                              </div>
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">Cores per CPU:</span>
                                <span className="text-text-main font-extrabold">{activeNode.cpu.cores} cores</span>
                              </div>
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">Threads:</span>
                                <span className="text-text-main font-extrabold">{activeNode.cpu.threads} threads</span>
                              </div>
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">Memory RAM:</span>
                                <span className="text-text-main font-extrabold text-emerald-600 dark:text-emerald-400">
                                  {toGiB(activeNode.memory.totalSize)} GiB
                                </span>
                              </div>
                              <div className="flex justify-between pt-1 border-t border-border-main/20 mt-1">
                                <span className="text-text-muted font-normal">Root Disk:</span>
                                <span className="text-text-main font-extrabold">
                                  {toGB(activeNode.rootDisk.totalSize)} GB ({activeNode.rootDisk.type || 'SSD'})
                                </span>
                              </div>
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">Data Disk:</span>
                                {activeNode.dataDisks && activeNode.dataDisks.length > 0 ? (
                                  <span className="text-text-main font-extrabold">
                                    {activeNode.dataDisks.map((d) => `${toGB(d.totalSize)} GB`).join(', ')} ({activeNode.dataDisks.length} disks)
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
                                <span className="text-text-main font-extrabold">{activeNode.os.name || 'Linux'}</span>
                              </div>
                              <div className="flex justify-between">
                                <span className="text-text-muted font-normal">Version:</span>
                                <span className="text-text-main font-extrabold">{activeNode.os.version || '—'}</span>
                              </div>
                              {activeNode.os.prettyName && (
                                <div className="flex justify-between">
                                  <span className="text-text-muted font-normal">Pretty Name:</span>
                                  <span className="text-text-main font-extrabold">{activeNode.os.prettyName}</span>
                                </div>
                              )}
                              {activeNode.os.versionId && (
                                <div className="flex justify-between">
                                  <span className="text-text-muted font-normal">Version ID:</span>
                                  <span className="text-text-main font-extrabold font-mono">{activeNode.os.versionId}</span>
                                </div>
                              )}
                              {activeNode.os.versionCodename && (
                                <div className="flex justify-between">
                                  <span className="text-text-muted font-normal">Codename:</span>
                                  <span className="text-text-main font-extrabold font-mono capitalize">{activeNode.os.versionCodename}</span>
                                </div>
                              )}
                              {activeNode.os.idLike && (
                                <div className="flex justify-between">
                                  <span className="text-text-muted font-normal">Base OS Like:</span>
                                  <span className="text-text-main font-extrabold font-mono uppercase">{activeNode.os.idLike}</span>
                                </div>
                              )}
                            </div>
                          </div>
                        )}

                        {/* Middle Row: Network Interfaces (Full-Width with Add/Remove in Edit Mode) */}
                        <div className="bg-bg-panel/30 border border-border-main/30 rounded-xl p-4 space-y-2.5 text-sm">
                          <div className="flex justify-between items-center border-b border-border-main/20 pb-1">
                            <span className="text-sm font-bold text-text-muted">Network Interfaces</span>
                            <span className="text-xs text-text-muted font-mono">Standard Units: Mbps / Gbps</span>
                          </div>
                          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3">
                            {activeNode.interfaces?.map((iface, idx) => (
                              <div key={idx} className="bg-bg-panel/60 border border-border-main/25 p-3 rounded-lg space-y-1 font-mono text-sm relative">
                                <div className="flex justify-between text-text-muted font-normal">
                                  <span className="font-extrabold text-text-main">{iface.name}</span>
                                  <div className="flex items-center gap-1.5">
                                    <span className="text-emerald-600 dark:text-emerald-400 uppercase text-xs font-extrabold">{iface.state}</span>
                                    {isEditingActiveNode && (
                                      <button
                                        type="button"
                                        onClick={() => handleDeleteNetworkInterface(activeNode.machineId, idx)}
                                        className="text-red-400 hover:text-red-300 font-bold ml-1 cursor-pointer"
                                      >
                                        ✕
                                      </button>
                                    )}
                                  </div>
                                </div>
                                {iface.ipv4CidrBlocks && iface.ipv4CidrBlocks.length > 0 && (
                                  <div className="text-text-muted text-xs">
                                    <span className="font-normal">IPv4: </span>
                                    <span className="text-text-main font-extrabold">{iface.ipv4CidrBlocks.join(', ')}</span>
                                  </div>
                                )}
                              </div>
                            ))}
                            {(!activeNode.interfaces || activeNode.interfaces.length === 0) && (
                              <div className="text-text-muted italic text-center py-2 col-span-3">No interfaces configured.</div>
                            )}
                          </div>

                          {/* Add Interface Form in Edit Mode */}
                          {isEditingActiveNode && (
                            <div className="pt-2 border-t border-border-main/20 flex flex-wrap gap-2 items-center">
                              <input
                                type="text"
                                placeholder="Interface Name (eth0)"
                                value={newIfaceName}
                                onChange={(e) => setNewIfaceName(e.target.value)}
                                className="bg-bg-panel border border-border-main/60 rounded-lg px-2.5 py-1 text-xs text-text-main"
                              />
                              <input
                                type="text"
                                placeholder="IPv4 CIDR (192.168.1.10/24)"
                                value={newIfaceCidr}
                                onChange={(e) => setNewIfaceCidr(e.target.value)}
                                className="bg-bg-panel border border-border-main/60 rounded-lg px-2.5 py-1 text-xs text-text-main font-mono"
                              />
                              <select
                                value={newIfaceState}
                                onChange={(e) => setNewIfaceState(e.target.value)}
                                className="bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main cursor-pointer"
                              >
                                <option value="up">UP</option>
                                <option value="down">DOWN</option>
                              </select>
                              <button
                                type="button"
                                onClick={() => handleAddNetworkInterface(activeNode.machineId)}
                                className="px-3 py-1 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg text-xs font-bold transition cursor-pointer whitespace-nowrap ml-auto"
                              >
                                Add Interface
                              </button>
                            </div>
                          )}
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
                                  {isEditingActiveNode && (
                                    <button
                                      type="button"
                                      onClick={() => handleDeleteFirewallRule(activeNode.machineId, idx)}
                                      className="text-red-400 hover:text-red-300 font-bold ml-2 px-1 cursor-pointer"
                                      title="Delete rule"
                                    >
                                      ✕
                                    </button>
                                  )}
                                </div>
                              ))}
                            </div>
                          ) : (
                            <div className="text-text-muted italic text-center py-4">No firewall rules configured.</div>
                          )}

                          {/* Firewall Rule Creator Form in Edit Mode */}
                          {isEditingActiveNode && (
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
                                    type="button"
                                    onClick={() => handleAddFirewallRule(activeNode.machineId)}
                                    className="w-full bg-emerald-600 hover:bg-emerald-700 text-white font-bold py-1 text-sm rounded-lg cursor-pointer transition text-center"
                                  >
                                    Add Rule
                                  </button>
                                </div>
                              </div>
                            </div>
                          )}
                        </div>

                      </div>
                    </div>
                  )}

                  {/* Row 4: Load Balancer (NLB) */}
                  <div className="space-y-3 p-5 bg-bg-input/40 border border-border-main/50 rounded-xl">
                    <div className="flex flex-col md:flex-row md:items-center justify-between border-b border-border-main/20 pb-2.5 gap-2">
                      <h4 className="text-sm font-bold text-emerald-600 dark:text-emerald-400 flex items-center">
                        <Cpu className="w-4 h-4 mr-1.5 text-emerald-500" />
                        4. Network Load Balancer (NLB) Configuration
                      </h4>
                      <button
                        type="button"
                        onClick={() => setIsEditingNlbConfig(!isEditingNlbConfig)}
                        className={`px-3 py-1.5 rounded-xl text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer whitespace-nowrap shrink-0 shadow-md shadow-emerald-500/20 ${
                          isEditingNlbConfig
                            ? 'bg-emerald-700 hover:bg-emerald-800 text-white border border-emerald-500/40'
                            : 'bg-emerald-600 hover:bg-emerald-700 text-white border border-emerald-500/40'
                        }`}
                      >
                        <Sliders className="w-3.5 h-3.5 text-white" />
                        <span>{isEditingNlbConfig ? 'Done Editing NLB' : 'Edit NLB Config'}</span>
                      </button>
                    </div>

                    {tunedNlbs && tunedNlbs.length > 0 ? (
                      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                        {tunedNlbs.map((nlb: any, idx: number) => (
                          <div key={idx} className="bg-bg-panel border border-border-main/40 rounded-xl p-4.5 space-y-3.5 leading-relaxed text-sm relative">
                            <div className="flex justify-between items-center border-b border-border-main/40 pb-2">
                              <span className="font-extrabold text-text-main text-base flex items-center">
                                <span className="w-2 h-2 bg-teal-500 rounded-full mr-2 animate-pulse" />
                                {nlb.software || 'Haproxy'} NLB
                              </span>
                              <div className="flex items-center gap-2">
                                <span className="text-xs font-extrabold px-2.5 py-0.5 bg-teal-500/10 text-teal-600 dark:text-teal-400 rounded-full border border-teal-500/30 uppercase">
                                  {nlb.listener?.protocol || 'tcp'} Mode
                                </span>
                                {isEditingNlbConfig && (
                                  <button
                                    type="button"
                                    onClick={() => handleDeleteNlbInstance(idx)}
                                    className="text-red-400 hover:text-red-300 font-bold text-xs cursor-pointer ml-1"
                                    title="Delete NLB instance"
                                  >
                                    ✕
                                  </button>
                                )}
                              </div>
                            </div>

                            <div className="space-y-3 font-mono text-sm text-text-muted font-normal">
                              {/* Traffic Ingress */}
                              <div className="flex justify-between items-center text-text-muted">
                                <span className="font-sans font-normal shrink-0">Traffic Ingress ➔</span>
                                {isEditingNlbConfig ? (
                                  <div className="flex items-center gap-1.5">
                                    <span className="font-bold text-xs">Port:</span>
                                    <input
                                      type="number"
                                      min="1"
                                      max="65535"
                                      value={nlb.listener?.port || 80}
                                      onChange={(e) => handleUpdateNlbListenerPort(idx, Number(e.target.value))}
                                      className="w-20 bg-bg-input border border-border-main/60 rounded px-2 py-0.5 text-xs text-text-main font-bold"
                                    />
                                  </div>
                                ) : (
                                  <span className="text-text-main font-bold">
                                    {(nlb.listener?.protocol || 'tcp').toUpperCase()} {nlb.listener?.bindAddress || '*'}:{nlb.listener?.port || 80}
                                  </span>
                                )}
                              </div>

                              {/* Balancing Method & Backend Targets */}
                              <div className="pl-4 border-l-2 border-teal-500/30 py-0.5 space-y-2">
                                <div className="text-xs text-teal-600 dark:text-teal-400 font-bold font-sans">
                                  ▼ Load Balancing: {nlb.backend?.balance || 'roundrobin'}
                                </div>
                                <div className="text-xs text-text-muted">
                                  Backend Group: <span className="text-text-main font-bold">{nlb.backend?.name || 'backend'}</span>
                                </div>

                                {/* Target VMs */}
                                <div className="space-y-1.5 pl-2 mt-1 border-l border-border-main/50">
                                  {(nlb.backend?.servers || []).map((srv: any, sIdx: number) => (
                                    <div key={sIdx} className="flex justify-between items-center text-xs">
                                      <span>├─ {srv.name}</span>
                                      <div className="flex items-center gap-1.5">
                                        <span className="text-text-main font-bold">{srv.ip}:{srv.port}</span>
                                        {isEditingNlbConfig && (
                                          <button
                                            type="button"
                                            onClick={() => handleDeleteNlbServer(idx, sIdx)}
                                            className="text-red-400 hover:text-red-300 font-bold text-xs cursor-pointer"
                                          >
                                            ✕
                                          </button>
                                        )}
                                      </div>
                                    </div>
                                  ))}
                                </div>

                                {/* Add Backend Server Form in Edit Mode */}
                                {isEditingNlbConfig && (
                                  <div className="pt-2 border-t border-border-main/20 flex gap-1.5 items-center">
                                    <input
                                      type="text"
                                      placeholder="Server IP (10.0.1.15)"
                                      value={newNlbServerIp}
                                      onChange={(e) => setNewNlbServerIp(e.target.value)}
                                      className="flex-1 bg-bg-input border border-border-main/60 rounded px-2 py-0.5 text-xs text-text-main"
                                    />
                                    <input
                                      type="number"
                                      placeholder="Port (80)"
                                      value={newNlbServerPort}
                                      onChange={(e) => setNewNlbServerPort(e.target.value)}
                                      className="w-16 bg-bg-input border border-border-main/60 rounded px-2 py-0.5 text-xs text-text-main"
                                    />
                                    <button
                                      type="button"
                                      onClick={() => handleAddNlbServer(idx)}
                                      className="px-2 py-0.5 bg-emerald-600 hover:bg-emerald-700 text-white rounded text-xs font-bold whitespace-nowrap cursor-pointer"
                                    >
                                      + Target
                                    </button>
                                  </div>
                                )}
                              </div>
                            </div>
                          </div>
                        ))}
                      </div>
                    ) : (
                      <div className="text-xs text-text-muted italic p-3 bg-bg-panel/20 rounded-lg border border-dashed border-border-main/50 text-center">
                        No Network Load Balancer configured.
                      </div>
                    )}

                    {/* Add NLB Instance Form in Edit Mode */}
                    {isEditingNlbConfig && (
                      <div className="pt-3 border-t border-border-main/20 flex flex-wrap gap-2 items-center bg-bg-panel/40 p-3 rounded-xl border border-border-main/20">
                        <span className="text-xs font-extrabold text-text-main">Add Custom NLB Instance:</span>
                        <input
                          type="text"
                          placeholder="Backend Name (backend_web)"
                          value={newNlbBackendName}
                          onChange={(e) => setNewNlbBackendName(e.target.value)}
                          className="bg-bg-panel border border-border-main/60 rounded px-2.5 py-1 text-xs text-text-main"
                        />
                        <input
                          type="number"
                          placeholder="Listener Port (80)"
                          value={newNlbPort}
                          onChange={(e) => setNewNlbPort(e.target.value)}
                          className="w-28 bg-bg-panel border border-border-main/60 rounded px-2.5 py-1 text-xs text-text-main"
                        />
                        <select
                          value={newNlbProtocol}
                          onChange={(e) => setNewNlbProtocol(e.target.value)}
                          className="bg-bg-panel border border-border-main/60 rounded px-2 py-1 text-xs text-text-main cursor-pointer"
                        >
                          <option value="tcp">TCP</option>
                          <option value="http">HTTP</option>
                        </select>
                        <button
                          type="button"
                          onClick={handleAddNlbInstance}
                          className="px-3.5 py-1 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg text-xs font-extrabold cursor-pointer ml-auto transition"
                        >
                          + Add NLB
                        </button>
                      </div>
                    )}
                  </div>

                  {/* Save spec and proceed button at the bottom of Step 1 */}
                  <div className="flex flex-row items-center justify-start pt-4 border-t border-border-main/20 mt-4 space-x-4">
                    <button
                      onClick={() => setShowSaveModal(true)}
                      className="px-6 py-3 bg-gradient-to-r from-emerald-400 via-teal-400 to-blue-600 hover:from-emerald-500 hover:to-blue-700 text-slate-950 rounded-xl text-sm font-extrabold flex items-center transition cursor-pointer shadow-lg shadow-emerald-500/10 shrink-0"
                    >
                      <Save className="w-4 h-4 mr-1.5 text-slate-950" /> Save Source Infra Revision
                    </button>
                    <div className="flex items-center space-x-2 text-sm text-text-muted">
                      <span className="font-bold">Model to save:</span>
                      <span className="text-emerald-600 dark:text-emerald-600 dark:text-emerald-400 font-extrabold text-sm">{selectedSourceModel.name}</span>
                      <span className="text-sm text-text-muted font-mono bg-bg-panel px-1.5 py-0.5 rounded border border-border-main/40">
                        v{selectedSourceModel.version || '1.0'}
                      </span>
                    </div>
                  </div>

                  {(onNext || onBack) && (
                    <div className="flex items-center justify-between pt-4 border-t border-border-main/20 mt-4">
                      {onBack ? (
                        <button
                          onClick={onBack}
                          className="px-4 py-2 bg-bg-input border border-border-main hover:bg-bg-main text-text-main font-bold text-xs rounded-xl transition cursor-pointer flex items-center space-x-1.5"
                        >
                          <ArrowLeft className="w-3.5 h-3.5" />
                          <span>Back to 1. Source Analysis</span>
                        </button>
                      ) : <div />}
                      {onNext && (
                        <button
                          onClick={onNext}
                          className="px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-bold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center space-x-2 cursor-pointer ml-auto"
                        >
                          <span>Next: Proceed to 3. Target Infra Optimization</span>
                          <ArrowRight className="w-4 h-4" />
                        </button>
                      )}
                    </div>
                  )}
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
        defaultName={selectedSourceModel && selectedSourceModel.id !== 'sample-source-infra-1' ? selectedSourceModel.name : ''}
        defaultDescription={selectedSourceModel && selectedSourceModel.id !== 'sample-source-infra-1' ? (selectedSourceModel.description || '') : ''}
        defaultVersion={selectedSourceModel && selectedSourceModel.id !== 'sample-source-infra-1' ? (selectedSourceModel.version || '1.0.0') : '1.0.0'}
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
