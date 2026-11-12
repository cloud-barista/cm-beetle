'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '../../store/migrationStore';
import { OnpremNode, OnpremInfra, OnpremModelEnvelope } from '../../types/migration';
import {
  SAMPLE_15_NODE_MODEL,
  SAMPLE_9_NODE_TESTBED_MODEL,
  SAMPLE_MIGRATION_TESTBED_MODEL,
} from '../../data/sample15NodeInfra';
import { SaveRevisionModal } from '../common/SaveRevisionModal';
import { ActiveModelBanner } from '../common/ActiveModelBanner';
import { DeleteModelConfirmModal } from '../common/DeleteModelConfirmModal';
import { UnsavedConfirmModal } from '../common/UnsavedConfirmModal';
import { ModelToast } from '../common/ModelToast';
import {
  buildDefaultModelName,
  cleanModelDisplayName,
  formatSourceOptionText,
  getStageBadgeMeta,
  isBuiltInSampleModel,
  resolveModelStageKind,
} from '../../utils/modelFormat';
import { beetleApi } from '../../api/client';
import {
  RefreshCw, ChevronDown, ChevronUp, Plus, Trash2, Server, Cpu,
  HardDrive, Network, Shield, Save, CheckCircle2, Loader2,
  FileText, Copy, Sparkles, GitBranch, Layers, X, Sliders,
  ArrowRight, ArrowLeft, Download,
} from 'lucide-react';

type ClusterGroupFilter = 'all' | 'ai-inference' | 'analytics-compute' | 'backend-db';

// Infer cluster group from explicit field or node hostname/GPU characteristics
const inferNodeClusterGroup = (n: OnpremNode): 'ai-inference' | 'analytics-compute' | 'backend-db' | null => {
  if (n.clusterGroup) return n.clusterGroup;
  const host = (n.hostname || '').toLowerCase();
  if ((n.gpuCards && n.gpuCards.length > 0) || /^(ai|gpu|llm|ollama|vllm)/.test(host)) {
    return 'ai-inference';
  }
  if (/^(data|analytics|spark|etl|airflow|worker)/.test(host)) {
    return 'analytics-compute';
  }
  if (/^(db|postgres|mysql|maria|backend|api|storage|minio|nfs|cache|redis|haproxy|nginx|tomcat)/.test(host)) {
    return 'backend-db';
  }
  return null;
};

export const SourceInfraRefinement: React.FC<{ onNext?: () => void; onBack?: () => void }> = ({ onNext, onBack }) => {
  const {
    savedSourceModels, selectedSourceModel, selectSourceModel,
    unsavedSourceDraft, setUnsavedPlanningDraft,
    fetchSavedSourceModels, saveSourceModel, updateSourceModel, deleteSourceModel,
  } = useMigrationStore();

  const [pendingModelId, setPendingModelId] = useState<string>(selectedSourceModel?.id || '');
  const [loadedModelId, setLoadedModelId] = useState<string>('');
  const [loadedModel, setLoadedModel] = useState<OnpremModelEnvelope | null>(null);
  const [activeTunedNodeId, setActiveTunedNodeId] = useState<string>('');
  const [tunedNodes, setTunedNodes] = useState<OnpremNode[]>([]);
  const [selectedClusterGroup, setSelectedClusterGroup] = useState<ClusterGroupFilter>('all');
  const [isJsonOpen, setIsJsonOpen] = useState(false);
  const [tunedNetwork, setTunedNetwork] = useState<any>(null);
  const [newCidr, setNewCidr] = useState('');
  const [activeStep, setActiveStep] = useState<number>(1);
  const [isModelLoaded, setIsModelLoaded] = useState(false);
  const [hasUnsavedChanges, setHasUnsavedChanges] = useState(false);
  const [excludedNodeIds, setExcludedNodeIds] = useState<string[]>([]);
  const [newRuleDir, setNewRuleDir] = useState('inbound');
  const [newRuleProto, setNewRuleProto] = useState('tcp');
  const [newRulePort, setNewRulePort] = useState('');
  const [newRuleCidr, setNewRuleCidr] = useState('0.0.0.0/0');
  const [showSaveModal, setShowSaveModal] = useState(false);
  const [showUnsavedNextModal, setShowUnsavedNextModal] = useState(false);
  const [showOverwriteLoadModal, setShowOverwriteLoadModal] = useState(false);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [toastMsg, setToastMsg] = useState<string | null>(null);
  const [toastTone, setToastTone] = useState<'success' | 'error'>('success');

  // Source Infra Summary (POST /beetle/summary/source) state
  const [showSourceSummaryModal, setShowSourceSummaryModal] = useState(false);
  const [sourceSummaryFormat, setSourceSummaryFormat] = useState<'html' | 'md' | 'json'>('html');
  const [sourceSummaryContent, setSourceSummaryContent] = useState<string>('');
  const [isLoadingSourceSummary, setIsLoadingSourceSummary] = useState(false);

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

  // New Data Disk Form state
  const [newDataDiskLabel, setNewDataDiskLabel] = useState('');
  const [newDataDiskSize, setNewDataDiskSize] = useState('100');
  const [newDataDiskType, setNewDataDiskType] = useState('SSD');

  // New GPU Card Form state
  const [newGpuVendor, setNewGpuVendor] = useState('NVIDIA');
  const [newGpuModel, setNewGpuModel] = useState('Tesla T4');
  const [newGpuVram, setNewGpuVram] = useState('16');
  const [newGpuDriver, setNewGpuDriver] = useState('535.183.01');
  const [newGpuCuda, setNewGpuCuda] = useState('12.2');

  // New Network Interface Form state
  const [newIfaceName, setNewIfaceName] = useState('');
  const [newIfaceCidr, setNewIfaceCidr] = useState('');
  const [newIfaceState, setNewIfaceState] = useState('up');
  const [newIfaceSpeed, setNewIfaceSpeed] = useState('1000');
  const [newIfaceSpeedUnit, setNewIfaceSpeedUnit] = useState('Mbps');

  // Unit helper functions conforming to Cloud-Barista standards (Memory: GiB, Disk: GB)
  const toGiB = (val: number): number => {
    if (!val) return 0;
    const raw = val > 1000000 ? val / (1024 * 1024 * 1024) : val;
    return Math.round(raw * 100) / 100;
  };

  const toGB = (val: number): number => {
    if (!val) return 0;
    return val > 1000000 ? Math.round(val / (1024 * 1024 * 1024)) : Math.round(val);
  };

  const handleUpdateNodeCpu = (machineId: string, key: string, value: any) => {
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      return {
        ...n,
        dataDisks: (n.dataDisks || []).filter((_, idx) => idx !== diskIdx)
      };
    }));
  };

  const handleAddGpuCard = (machineId: string) => {
    if (!newGpuModel.trim()) return;
    setHasUnsavedChanges(true);
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      const currentCards = n.gpuCards || [];
      return {
        ...n,
        gpuCards: [
          ...currentCards,
          {
            vendor: newGpuVendor,
            model: newGpuModel.trim(),
            memoryTotalGB: Number(newGpuVram) || 16,
            driverVersion: newGpuDriver.trim() || undefined,
            cudaVersion: newGpuCuda.trim() || undefined,
            driverIndex: String(currentCards.length)
          }
        ]
      };
    }));
    setNewGpuModel('Tesla T4');
    setNewGpuVram('16');
  };

  const handleDeleteGpuCard = (machineId: string, idx: number) => {
    setHasUnsavedChanges(true);
    setTunedNodes(prev => prev.map(n => {
      if (n.machineId !== machineId) return n;
      const filtered = (n.gpuCards || []).filter((_, i) => i !== idx);
      return {
        ...n,
        gpuCards: filtered.length > 0 ? filtered : undefined
      };
    }));
  };

  const handleUpdateNodeOs = (machineId: string, key: string, value: string) => {
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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
    const cleanIp = newNlbServerIp.trim();
    if (!cleanIp) return;
    setHasUnsavedChanges(true);
    const srvPort = parseInt(newNlbServerPort, 10) || 80;
    const matchedNode = tunedNodes.find(n =>
      n.interfaces?.some(iface => iface.ipv4CidrBlocks?.some(cidr => cidr.split('/')[0].trim() === cleanIp))
    );
    const srvName = matchedNode?.hostname || `srv-${cleanIp.replaceAll('.', '-')}`;
    const newSrv = {
      name: srvName,
      ip: cleanIp,
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
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
    setTunedNlbs(prev => prev.filter((_, idx) => idx !== nlbIdx));
  };

  const sampleModels: OnpremModelEnvelope[] = [
    SAMPLE_MIGRATION_TESTBED_MODEL,
    SAMPLE_9_NODE_TESTBED_MODEL,
    SAMPLE_15_NODE_MODEL,
  ];
  const draftModels: OnpremModelEnvelope[] = unsavedSourceDraft ? [unsavedSourceDraft] : [];
  const userModels = savedSourceModels.filter(m => !isBuiltInSampleModel(m));
  const userPlanningModels = userModels.filter(m => m.stage === 'planning');
  const userSourceModels = userModels.filter(m => m.stage !== 'planning');
  const allModels: OnpremModelEnvelope[] = [
    ...draftModels,
    ...userPlanningModels,
    ...userSourceModels,
    ...sampleModels,
  ];

  useEffect(() => { fetchSavedSourceModels(); }, []);

  useEffect(() => {
    if (selectedSourceModel?.id && selectedSourceModel.id !== 'planning-draft') {
      setPendingModelId(selectedSourceModel.id);
    }
  }, [selectedSourceModel?.id]);

  useEffect(() => {
    if (pendingModelId === 'source-draft' && !unsavedSourceDraft) {
      const nextId =
        (selectedSourceModel && selectedSourceModel.id !== 'planning-draft' ? selectedSourceModel.id : '') ||
        allModels[0]?.id ||
        '';
      setPendingModelId(nextId);
    }
    if (loadedModelId === 'source-draft' && !unsavedSourceDraft && selectedSourceModel && !selectedSourceModel.isDraft) {
      setLoadedModelId(selectedSourceModel.id);
      setLoadedModel(selectedSourceModel);
    }
  }, [pendingModelId, loadedModelId, unsavedSourceDraft, selectedSourceModel, allModels]);

  const pendingModel = allModels.find(m => m.id === pendingModelId) || null;
  const isPendingAlreadyLoaded =
    isModelLoaded && Boolean(pendingModel) && loadedModelId === pendingModel?.id && !hasUnsavedChanges;

  const executeLoadModel = () => {
    if (!pendingModel || !pendingModel.onpremiseInfraModel) return;
    selectSourceModel(pendingModel);
    setLoadedModel(pendingModel);
    setLoadedModelId(pendingModel.id);
    setTunedNodes(JSON.parse(JSON.stringify(pendingModel.onpremiseInfraModel.nodes || [])));
    setTunedNetwork(JSON.parse(JSON.stringify(pendingModel.onpremiseInfraModel.network || { ipv4Networks: {}, ipv6Networks: {} })));
    setTunedNlbs(
      pendingModel.onpremiseInfraModel.nlbs
        ? JSON.parse(JSON.stringify(pendingModel.onpremiseInfraModel.nlbs))
        : []
    );
    setExcludedNodeIds([]);
    setSelectedClusterGroup('all');
    setIsModelLoaded(true);
    setHasUnsavedChanges(false);
    setActiveStep(2);
    if (pendingModel.onpremiseInfraModel.nodes && pendingModel.onpremiseInfraModel.nodes.length > 0) {
      setActiveTunedNodeId(pendingModel.onpremiseInfraModel.nodes[0]?.machineId || '');
    }
  };

  const handleLoadModel = () => {
    if (!pendingModel || !pendingModel.onpremiseInfraModel) return;
    if (isModelLoaded && hasUnsavedChanges) {
      setShowOverwriteLoadModal(true);
      return;
    }
    executeLoadModel();
  };

  const handleDeleteModel = () => {
    if (!pendingModel || isBuiltInSampleModel(pendingModel)) return;
    setShowDeleteConfirm(true);
  };

  const activeNode = tunedNodes.find((n) => n.machineId === activeTunedNodeId);

  const handleToggleNodeExclude = (machineId: string) => {
    setHasUnsavedChanges(true);
    setExcludedNodeIds(prev =>
      prev.includes(machineId)
        ? prev.filter(id => id !== machineId)
        : [...prev, machineId]
    );
  };

  const handleDeleteFirewallRule = (machineId: string, ruleIdx: number) => {
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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

  const [isJsonCopied, setIsJsonCopied] = useState(false);

  const filteredPlanningNodes = tunedNodes
    .filter(n => !excludedNodeIds.includes(n.machineId))
    .map(n => ({
      ...n,
      gpuCards: n.gpuCards && n.gpuCards.length > 0 ? n.gpuCards : undefined,
    }));

  const totalPlanningGpus = filteredPlanningNodes.reduce(
    (sum, n) => sum + (n.gpuCards ? n.gpuCards.length : 0),
    0
  );

  const effectivePlanningInfra = loadedModel
    ? {
        ...loadedModel.onpremiseInfraModel,
        nodes: filteredPlanningNodes,
        network: tunedNetwork || loadedModel.onpremiseInfraModel.network,
        nlbs: tunedNlbs.length > 0 ? tunedNlbs : undefined,
      }
    : null;

  useEffect(() => {
    if (isModelLoaded && hasUnsavedChanges && loadedModel && effectivePlanningInfra && filteredPlanningNodes.length > 0) {
      const draftName =
        loadedModel.stage === 'planning' && !isBuiltInSampleModel(loadedModel) && !loadedModel.isDraft
          ? cleanModelDisplayName(loadedModel.name)
          : buildDefaultModelName('planning', loadedModel.name);
      setUnsavedPlanningDraft({
        id: 'planning-draft',
        name: draftName,
        description: `Refined planning infrastructure model (${filteredPlanningNodes.length} nodes)`,
        onpremiseInfraModel: JSON.parse(JSON.stringify(effectivePlanningInfra)),
        version: loadedModel.version || '1.0.0',
        stage: 'planning',
        category: 'infra',
        schemaType: 'OnpremInfra',
        lineage: {
          sourceGroupId: loadedModel.lineage?.sourceGroupId,
          sourceModelId: loadedModel.isDraft ? undefined : (loadedModel.lineage?.sourceModelId || loadedModel.id),
        },
        isDraft: true,
      });
    } else {
      setUnsavedPlanningDraft(null);
    }
  }, [isModelLoaded, hasUnsavedChanges, loadedModel, tunedNodes, excludedNodeIds, tunedNetwork, tunedNlbs, setUnsavedPlanningDraft]);

  const handleSaveArtifact = async (result: { name: string; description: string; version: string; overwriteId: string | null }) => {
    if (!loadedModel || !effectivePlanningInfra) return;
    const filteredNodes = filteredPlanningNodes;
    const updatedInfra = effectivePlanningInfra;

    const resolvedSourceParentId = loadedModel.isDraft
      ? undefined
      : loadedModel.stage === 'planning'
        ? (loadedModel.lineage?.sourceModelId || loadedModel.id)
        : loadedModel.id;

    const saved = result.overwriteId
      ? await updateSourceModel(result.overwriteId, result.name, result.description, result.version, updatedInfra, {
          stage: 'planning',
          parentId: resolvedSourceParentId,
          sourceGroupId: loadedModel.lineage?.sourceGroupId
        })
      : await saveSourceModel(result.name, result.description, result.version, updatedInfra, {
          stage: 'planning',
          parentId: resolvedSourceParentId,
          sourceGroupId: loadedModel.lineage?.sourceGroupId
        });

    setPendingModelId(saved.id);
    setLoadedModel(saved);
    setLoadedModelId(saved.id);
    setTunedNodes(filteredNodes);
    setExcludedNodeIds([]);
    setHasUnsavedChanges(false);
    setActiveStep(3);
    setToastTone('success');
    setToastMsg(
      result.overwriteId
        ? `Planning model "${saved.name}" (v${saved.version || result.version}) updated successfully.`
        : `Planning model "${saved.name}" (v${saved.version || result.version}) saved successfully.`
    );
  };

  const handleProceedNext = () => {
    if (!onNext) return;
    if (isModelLoaded && hasUnsavedChanges) {
      setShowUnsavedNextModal(true);
      return;
    }
    onNext();
  };

  const handleAddCidr = (cidr: string) => {
    if (!cidr || !tunedNetwork) return;
    const currentCidrs = tunedNetwork.ipv4Networks?.cidrBlocks || [];
    if (currentCidrs.includes(cidr)) return;
    setHasUnsavedChanges(true);
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
    setHasUnsavedChanges(true);
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

  const handleFetchSourceSummary = async (fmt: 'html' | 'md' | 'json' = sourceSummaryFormat) => {
    if (!loadedModel) return;
    setSourceSummaryFormat(fmt);
    setShowSourceSummaryModal(true);
    setIsLoadingSourceSummary(true);
    try {
      const currentInfra = {
        ...loadedModel.onpremiseInfraModel,
        nodes: filteredPlanningNodes.length > 0 ? filteredPlanningNodes : loadedModel.onpremiseInfraModel.nodes,
        network: tunedNetwork || loadedModel.onpremiseInfraModel.network,
        nlbs: tunedNlbs.length > 0 ? tunedNlbs : undefined,
      };
      const content = await beetleApi.getSourceInfraSummary(currentInfra, fmt);
      setSourceSummaryContent(content);
    } catch (err: any) {
      setSourceSummaryContent(`Failed to generate source infrastructure summary: ${err.message || 'Unknown error'}`);
    } finally {
      setIsLoadingSourceSummary(false);
    }
  };

  const handleDownloadSourceSummary = () => {
    if (!sourceSummaryContent) return;
    const mimeType =
      sourceSummaryFormat === 'json'
        ? 'application/json'
        : sourceSummaryFormat === 'md'
        ? 'text/markdown'
        : 'text/html';
    const blob = new Blob([sourceSummaryContent], { type: mimeType });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `source_infra_summary_${loadedModel?.name?.replace(/\s+/g, '_') || 'model'}.${sourceSummaryFormat}`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  };

  return (
    <div className="space-y-6 mx-auto pb-24">

      {/* Top Banner Description Box */}
      <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
        <div className="flex items-center gap-2 shrink-0">
          <Sliders className="w-5 h-5 text-emerald-500" />
          <h2 className="text-base font-extrabold text-text-main tracking-tight">
            Migration Planning
          </h2>
        </div>
        <span className="text-sm text-text-muted">
          Plan and refine extracted infrastructure models, adjust node and network specifications, and save planning artifact revisions.
        </span>
      </div>

      {/* ═══ STEP 1: Source / Planning Infrastructure Model Selection ═══ */}
      <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${activeStep >= 1 ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
        <div className="flex items-center space-x-3 mb-4 border-b border-border-main/40 pb-3">
          <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${isModelLoaded ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
            {isModelLoaded ? '✓' : '1'}
          </span>
          <h3 className="text-base font-extrabold text-text-main">Step 1: Source / Planning Infrastructure Model Selection</h3>
        </div>
        <div className="space-y-4">
          <div className="max-w-2xl space-y-3">
            <label className="block text-sm font-normal text-text-muted">Source or Planning Infrastructure Model</label>
            <div className="flex flex-wrap items-center gap-2.5">
              <select
                value={pendingModelId}
                onChange={(e) => setPendingModelId(e.target.value)}
                className="flex-1 min-w-[260px] bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-3 text-sm font-extrabold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
              >
                <option value="">-- Choose Infrastructure Model --</option>
                {draftModels.length > 0 && (
                  <optgroup label="Draft Model">
                    {draftModels.map(m => (
                      <option key={m.id} value={m.id}>{formatSourceOptionText(m, '1.0', true)}</option>
                    ))}
                  </optgroup>
                )}
                {userPlanningModels.length > 0 && (
                  <optgroup label="Planning Models">
                    {userPlanningModels.map(m => (
                      <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                    ))}
                  </optgroup>
                )}
                {userSourceModels.length > 0 && (
                  <optgroup label="Source Models">
                    {userSourceModels.map(m => (
                      <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                    ))}
                  </optgroup>
                )}
                <optgroup label="Sample Models">
                  {sampleModels.map(m => (
                    <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                  ))}
                </optgroup>
              </select>
              {pendingModel && (() => {
                const badge = getStageBadgeMeta(resolveModelStageKind(pendingModel, 'source'));
                return (
                  <span className={`px-3 py-1.5 rounded-xl text-xs font-extrabold border shrink-0 ${badge.className}`}>
                    {badge.label}
                  </span>
                );
              })()}
            </div>

            {pendingModel && (
              <div className="text-xs bg-bg-input/50 border border-border-main/50 rounded-xl px-3.5 py-2.5 flex flex-wrap items-center gap-x-4 gap-y-1">
                <span>
                  <span className="font-normal text-text-muted">Nodes: </span>
                  <span className="font-extrabold text-text-main">{pendingModel.onpremiseInfraModel?.nodes?.length || 0}</span>
                </span>
                {pendingModel.description && (
                  <span className="truncate max-w-md">
                    <span className="font-normal text-text-muted">Info: </span>
                    <span className="font-extrabold text-text-main">{pendingModel.description}</span>
                  </span>
                )}
              </div>
            )}

            <div className="flex flex-wrap items-center gap-3 pt-1">
              <button
                type="button"
                onClick={handleLoadModel}
                disabled={!pendingModel}
                className={`px-5 py-2.5 rounded-xl text-sm font-extrabold flex items-center gap-1.5 transition cursor-pointer ${
                  !pendingModel
                    ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                    : isPendingAlreadyLoaded
                      ? 'bg-emerald-500/20 border border-emerald-500/40 text-emerald-600 dark:text-emerald-400'
                      : 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 shadow-md shadow-emerald-500/25'
                }`}
              >
                {isPendingAlreadyLoaded ? (
                  <>
                    <CheckCircle2 className="w-4 h-4" />
                    <span>Model Loaded</span>
                  </>
                ) : (
                  <>
                    <RefreshCw className="w-4 h-4" />
                    <span>Load Model</span>
                  </>
                )}
              </button>
              {pendingModel && !isBuiltInSampleModel(pendingModel) && !pendingModel.isDraft && (
                <button
                  type="button"
                  onClick={handleDeleteModel}
                  className="px-5 py-2.5 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-xl text-sm font-extrabold flex items-center gap-1.5 transition cursor-pointer"
                >
                  <Trash2 className="w-4 h-4" />
                  <span>Delete Model</span>
                </button>
              )}
            </div>
          </div>
        </div>
      </div>

      {/* -------------------------------------------------------------
          STEP 2: Planning Infrastructure Specification & Refinement
         ------------------------------------------------------------- */}
      <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${isModelLoaded ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
        <div className="flex items-center justify-between mb-4 border-b border-border-main/40 pb-3">
          <div className="flex items-center space-x-3">
            <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${activeStep > 2 ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
              {activeStep > 2 ? '✓' : '2'}
            </span>
            <h3 className="text-base font-extrabold text-text-main">
              Step 2: Planning Infrastructure Specification & Refinement
            </h3>
          </div>
          {loadedModel && isModelLoaded && (
            <div className="flex items-center space-x-2">
              <button
                type="button"
                onClick={() => handleFetchSourceSummary(sourceSummaryFormat)}
                className="px-3 py-1.5 bg-teal-500/10 border border-teal-500/40 hover:bg-teal-500/20 rounded-lg text-sm font-bold transition cursor-pointer flex items-center text-teal-600 dark:text-teal-400"
              >
                <FileText className="w-3.5 h-3.5 mr-1.5" />
                Source Summary (POST /summary/source)
              </button>
              <button
                type="button"
                onClick={() => setIsJsonOpen(!isJsonOpen)}
                className="px-3 py-1.5 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-sm font-bold transition cursor-pointer flex items-center text-emerald-600 dark:text-emerald-400"
              >
                <Copy className="w-3.5 h-3.5 mr-1" />
                JSON View
              </button>
            </div>
          )}
        </div>

        {loadedModel && isModelLoaded && (
          <div className="mb-4">
            <ActiveModelBanner
              mode={loadedModel.isDraft ? 'draft' : 'loaded'}
              name={loadedModel.name}
              version={loadedModel.isDraft ? undefined : loadedModel.version}
              stageKind={resolveModelStageKind(loadedModel, 'source')}
              hasUnsavedChanges={hasUnsavedChanges || Boolean(loadedModel.isDraft)}
              updatedTime={loadedModel.updatedTime}
              metrics={[
                { label: 'Nodes', value: String(filteredPlanningNodes.length) },
                {
                  label: 'GPU',
                  value: totalPlanningGpus > 0 ? `${totalPlanningGpus} Card(s)` : 'CPU Only',
                  tone: totalPlanningGpus > 0 ? 'emerald' : 'default',
                },
                {
                  label: 'NLB',
                  value: tunedNlbs.length > 0 ? `${tunedNlbs.length} Configured` : 'No NLB',
                  tone: tunedNlbs.length > 0 ? 'teal' : 'default',
                },
              ]}
            />
          </div>
        )}

        {isModelLoaded && loadedModel && (
          <div className="space-y-4">
            {isJsonOpen && (
              <div className="space-y-3 mb-4">
                <div className="flex justify-between items-center bg-bg-input px-3.5 py-2 rounded-xl border border-border-main/50">
                  <span className="text-sm text-text-muted font-mono">onpremiseInfraModel</span>
                  <button
                    type="button"
                    onClick={() => {
                      navigator.clipboard.writeText(JSON.stringify(effectivePlanningInfra, null, 2));
                      setIsJsonCopied(true);
                      setTimeout(() => setIsJsonCopied(false), 2000);
                    }}
                    className="px-3 py-1 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-xs font-bold transition flex items-center cursor-pointer text-emerald-600 dark:text-emerald-400"
                  >
                    <Copy className="w-3.5 h-3.5 mr-1" /> {isJsonCopied ? 'Copied!' : 'Copy JSON'}
                  </button>
                </div>
                <pre className="text-sm font-mono text-slate-800 dark:text-emerald-400 bg-bg-panel p-3.5 rounded-xl border border-border-main overflow-y-auto max-h-[400px] select-text">
                  {JSON.stringify(effectivePlanningInfra, null, 2)}
                </pre>
              </div>
            )}

            <div className="mt-2">
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
                    <div className="flex flex-wrap items-center justify-between gap-3">
                      <div className="flex flex-wrap items-center gap-2.5">
                        <h4 className="text-sm font-bold text-emerald-600 dark:text-emerald-400 flex items-center">
                          <Sliders className="w-4 h-4 mr-1.5 text-emerald-600 dark:text-emerald-400" />
                          2. Node List ({tunedNodes.length} Nodes)
                        </h4>
                        <span className="px-2.5 py-0.5 rounded-lg bg-emerald-500/15 border border-emerald-500/30 text-xs font-extrabold font-mono text-emerald-600 dark:text-emerald-400">
                          Included: {tunedNodes.length - excludedNodeIds.length} / {tunedNodes.length}
                        </span>
                      </div>

                      {/* Cluster Group Filter Pills */}
                      <div className="flex flex-wrap items-center gap-2">
                        {tunedNodes.some((n) => inferNodeClusterGroup(n) !== null) && (() => {
                          const aiNodes = tunedNodes.filter((n) => inferNodeClusterGroup(n) === 'ai-inference');
                          const analyticsNodes = tunedNodes.filter((n) => inferNodeClusterGroup(n) === 'analytics-compute');
                          const backendNodes = tunedNodes.filter((n) => inferNodeClusterGroup(n) === 'backend-db');
                          const pills: { id: ClusterGroupFilter; label: string; count: number; nodes: OnpremNode[] }[] = [
                            { id: 'all', label: 'All Nodes', count: tunedNodes.length, nodes: tunedNodes },
                            { id: 'ai-inference', label: 'AI Inference', count: aiNodes.length, nodes: aiNodes },
                            { id: 'analytics-compute', label: 'Analytics Compute', count: analyticsNodes.length, nodes: analyticsNodes },
                            { id: 'backend-db', label: 'Backend & DB', count: backendNodes.length, nodes: backendNodes },
                          ];
                          return (
                            <div className="flex flex-wrap items-center gap-1.5">
                              {pills.filter((p) => p.id === 'all' || p.count > 0).map((pill) => {
                                const isSelected = selectedClusterGroup === pill.id;
                                return (
                                  <button
                                    key={pill.id}
                                    type="button"
                                    onClick={() => {
                                      setSelectedClusterGroup(pill.id);
                                      if (pill.nodes.length > 0 && !pill.nodes.some((n) => n.machineId === activeTunedNodeId)) {
                                        setActiveTunedNodeId(pill.nodes[0].machineId);
                                      }
                                    }}
                                    className={`px-3 py-1.5 rounded-lg text-sm font-bold transition cursor-pointer border ${
                                      isSelected
                                        ? 'bg-emerald-500 text-slate-950 border-emerald-500 shadow-sm shadow-emerald-500/20'
                                        : 'bg-bg-panel border-border-main text-text-muted hover:text-text-main hover:border-emerald-500/40'
                                    }`}
                                  >
                                    {pill.label} ({pill.count})
                                  </button>
                                );
                              })}
                            </div>
                          );
                        })()}

                        {excludedNodeIds.length > 0 && (
                          <button
                            type="button"
                            onClick={() => {
                              setHasUnsavedChanges(true);
                              setExcludedNodeIds([]);
                            }}
                            className="px-3 py-1.5 rounded-lg text-xs font-extrabold transition cursor-pointer border bg-bg-panel hover:bg-bg-input border-border-main text-text-muted hover:text-text-main"
                            title="Restore all excluded nodes"
                          >
                            Include All ({tunedNodes.length})
                          </button>
                        )}
                      </div>
                    </div>
                    <div className="flex flex-wrap gap-2.5 max-h-44 overflow-y-auto p-2 bg-bg-panel/20 rounded-xl border border-border-main/20">
                      {tunedNodes
                        .filter((n) => selectedClusterGroup === 'all' || inferNodeClusterGroup(n) === selectedClusterGroup)
                        .map((n) => {
                          const isExcluded = excludedNodeIds.includes(n.machineId);
                          const isActive = n.machineId === activeTunedNodeId;
                          const vcpus = (n.cpu?.cpus || 1) * (n.cpu?.threads || n.cpu?.cores || 1);
                          return (
                            <div
                              key={n.machineId}
                              className={`flex items-center space-x-2 px-3 py-1.5 rounded-lg text-sm font-bold transition border ${isActive
                                  ? 'bg-emerald-500/10 border-emerald-500 text-emerald-600 dark:text-emerald-400 shadow-md shadow-emerald-500/10'
                                  : 'bg-bg-panel border-border-main text-text-muted hover:text-text-main'
                                } ${isExcluded ? 'opacity-40' : ''}`}
                            >
                              <button
                                onClick={() => setActiveTunedNodeId(n.machineId)}
                                className="flex items-center space-x-1.5 cursor-pointer focus:outline-none"
                              >
                                <HardDrive className="w-4 h-4" />
                                <span className={isExcluded ? 'line-through' : ''}>{n.hostname}</span>
                                <span className="text-xs font-mono text-text-muted">
                                  ({vcpus}vCPU/{toGiB(n.memory?.totalSize || 0)}GiB)
                                </span>
                                {n.gpuCards && n.gpuCards.length > 0 && (
                                  <span className="ml-1 px-1.5 py-0.5 bg-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs rounded font-extrabold border border-emerald-500/30">
                                    GPU
                                  </span>
                                )}
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
                                title={isExcluded ? 'Include node in recommendation' : 'Exclude node from recommendation'}
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

                                  {/* GPU Cards Management */}
                                  <div className="bg-bg-panel/50 border border-emerald-500/30 rounded-xl p-4 space-y-2.5 text-sm">
                                    <div className="flex justify-between items-center border-b border-border-main/20 pb-1.5">
                                      <div className="flex items-center gap-2">
                                        <span className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400">GPU Cards</span>
                                        <span className="text-xs text-text-muted">(Source Hardware Metadata)</span>
                                      </div>
                                      <span className="text-xs text-text-muted font-mono">Discrete VRAM (GB)</span>
                                    </div>
                                    {activeNode.gpuCards && activeNode.gpuCards.length > 0 ? (
                                      <div className="space-y-1.5 max-h-32 overflow-y-auto">
                                        {activeNode.gpuCards.map((g, idx) => (
                                          <div key={idx} className="flex justify-between items-center bg-bg-panel px-3 py-1.5 rounded-lg border border-border-main/30 text-xs">
                                            <div className="flex flex-wrap items-center gap-2">
                                              <span className="font-extrabold text-text-main">{g.vendor} {g.model}</span>
                                              {g.driverVersion && (
                                                <span className="text-teal-600 dark:text-teal-400 font-mono text-xs font-bold">Driver {g.driverVersion}</span>
                                              )}
                                              {g.cudaVersion && (
                                                <span className="text-text-muted font-mono text-xs">CUDA {g.cudaVersion}</span>
                                              )}
                                            </div>
                                            <div className="flex items-center gap-2">
                                              <span className="font-extrabold text-emerald-600 dark:text-emerald-400">{g.memoryTotalGB || 0} GB VRAM</span>
                                              <button
                                                type="button"
                                                onClick={() => handleDeleteGpuCard(activeNode.machineId, idx)}
                                                className="text-red-400 hover:text-red-300 font-bold cursor-pointer"
                                                title="Remove GPU card"
                                              >
                                                ✕
                                              </button>
                                            </div>
                                          </div>
                                        ))}
                                      </div>
                                    ) : (
                                      <div className="text-xs text-text-muted italic">No GPU cards configured.</div>
                                    )}

                                    {/* Add GPU Card Form */}
                                    <div className="pt-2 border-t border-border-main/20 flex flex-wrap gap-2 items-center">
                                      <select
                                        value={newGpuVendor}
                                        onChange={(e) => setNewGpuVendor(e.target.value)}
                                        className="bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main cursor-pointer"
                                      >
                                        <option value="NVIDIA">NVIDIA</option>
                                        <option value="AMD">AMD</option>
                                        <option value="Intel">Intel</option>
                                        <option value="Google">Google</option>
                                      </select>
                                      <input
                                        type="text"
                                        placeholder="Model (e.g. Tesla T4, A100)"
                                        value={newGpuModel}
                                        onChange={(e) => setNewGpuModel(e.target.value)}
                                        className="flex-1 min-w-[120px] bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-bold"
                                      />
                                      <div className="flex items-center gap-1">
                                        <input
                                          type="number"
                                          min="1"
                                          placeholder="VRAM"
                                          value={newGpuVram}
                                          onChange={(e) => setNewGpuVram(e.target.value)}
                                          className="w-16 bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-bold"
                                        />
                                        <span className="text-xs text-text-muted font-bold">GB</span>
                                      </div>
                                      <input
                                        type="text"
                                        placeholder="Driver (e.g. 535.183)"
                                        value={newGpuDriver}
                                        onChange={(e) => setNewGpuDriver(e.target.value)}
                                        className="w-28 bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-mono"
                                      />
                                      <input
                                        type="text"
                                        placeholder="CUDA (e.g. 12.2)"
                                        value={newGpuCuda}
                                        onChange={(e) => setNewGpuCuda(e.target.value)}
                                        className="w-24 bg-bg-panel border border-border-main/60 rounded-lg px-2 py-1 text-xs text-text-main font-mono"
                                      />
                                      <button
                                        type="button"
                                        onClick={() => handleAddGpuCard(activeNode.machineId)}
                                        className="px-3 py-1 bg-emerald-600 hover:bg-emerald-700 text-white rounded-lg text-xs font-bold transition cursor-pointer whitespace-nowrap"
                                      >
                                        Add GPU
                                      </button>
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
                              <div className="flex justify-between items-start pt-1.5 border-t border-border-main/20 mt-1">
                                <span className="text-text-muted font-normal whitespace-nowrap">GPU Cards:</span>
                                {activeNode.gpuCards && activeNode.gpuCards.length > 0 ? (
                                  <div className="flex flex-col items-end gap-1 text-right">
                                    {activeNode.gpuCards.map((g, idx) => {
                                      const displayName = (g.vendor && g.model?.startsWith(g.vendor))
                                        ? g.model
                                        : `${g.vendor ? `${g.vendor} ` : ''}${g.model || 'GPU'}`;
                                      return (
                                        <div key={idx} className="flex flex-col items-end gap-0.5">
                                          <div className="flex items-center gap-1.5">
                                            <span className="text-emerald-600 dark:text-emerald-400 font-extrabold text-xs sm:text-sm">
                                              {displayName}
                                            </span>
                                            <span className="text-text-muted font-medium text-xs">
                                              ({g.memoryTotalGB || 0}GB VRAM)
                                            </span>
                                          </div>
                                          {(g.driverVersion || g.cudaVersion) && (
                                            <span className="text-teal-600 dark:text-teal-400 font-mono text-xs font-bold">
                                              {[
                                                g.driverVersion ? `Driver ${g.driverVersion}` : '',
                                                g.cudaVersion ? `CUDA ${g.cudaVersion}` : ''
                                              ]
                                                .filter(Boolean)
                                                .join(' / ')}
                                            </span>
                                          )}
                                        </div>
                                      );
                                    })}
                                  </div>
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
                                <span className="text-text-main font-extrabold">{activeNode.os.version || '-'}</span>
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
                                    <datalist id={`nlb-node-ips-${idx}`}>
                                      {tunedNodes.flatMap(n =>
                                        (n.interfaces || []).flatMap(iface =>
                                          (iface.ipv4CidrBlocks || [])
                                            .map(cidr => cidr.split('/')[0].trim())
                                            .filter(ip => ip && !ip.startsWith('127.'))
                                            .map(ip => (
                                              <option key={`${n.machineId}-${ip}`} value={ip}>
                                                {`${n.hostname} (${ip})`}
                                              </option>
                                            ))
                                        )
                                      )}
                                    </datalist>
                                    <input
                                      type="text"
                                      list={`nlb-node-ips-${idx}`}
                                      placeholder="Select Node or Enter IP"
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

                  {/* Unified Bottom Action & Navigation Bar */}
                  <div className="flex flex-col sm:flex-row items-center justify-between gap-3 pt-4 border-t border-border-main/20 mt-4">
                    {onBack ? (
                      <button
                        type="button"
                        onClick={onBack}
                        className="w-full sm:w-auto px-4 py-2.5 bg-bg-input border border-border-main hover:bg-bg-main text-text-main font-extrabold text-xs rounded-xl transition cursor-pointer flex items-center justify-center space-x-1.5"
                      >
                        <ArrowLeft className="w-3.5 h-3.5" />
                        <span>Back to 1. Source Analysis</span>
                      </button>
                    ) : <div />}

                    <div className="flex flex-wrap items-center gap-2.5 w-full sm:w-auto justify-end">
                      <button
                        type="button"
                        onClick={() => setShowSaveModal(true)}
                        className="w-full sm:w-auto px-5 py-2.5 bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-500 hover:from-emerald-500 hover:to-teal-600 text-slate-950 rounded-xl text-xs font-extrabold flex items-center justify-center gap-1.5 transition cursor-pointer shadow-md shadow-emerald-500/20"
                      >
                        <Save className="w-3.5 h-3.5 text-slate-950" />
                        <span>Save Planning Model</span>
                      </button>
                      {onNext && (
                        <button
                          type="button"
                          onClick={handleProceedNext}
                          className="w-full sm:w-auto px-6 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-extrabold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center justify-center gap-2 cursor-pointer"
                        >
                          <span>Next: Proceed to 3. Target Cloud Optimizer</span>
                          <ArrowRight className="w-4 h-4" />
                        </button>
                      )}
                    </div>
                  </div>
              </div>
            </div>
          </div>
        )}
      </div>

      <SaveRevisionModal
        isOpen={showSaveModal}
        onClose={() => setShowSaveModal(false)}
        title="Save Planning Model"
        stageKind="planning"
        defaultName={
          loadedModel
            ? loadedModel.stage === 'planning' && !isBuiltInSampleModel(loadedModel) && !loadedModel.isDraft
              ? cleanModelDisplayName(loadedModel.name)
              : buildDefaultModelName('planning', loadedModel.name)
            : ''
        }
        defaultDescription={
          loadedModel
            ? loadedModel.stage === 'planning' && !isBuiltInSampleModel(loadedModel) && !loadedModel.isDraft
              ? loadedModel.description || 'Refined planning infrastructure model'
              : `Refined planning infrastructure model derived from ${cleanModelDisplayName(loadedModel.name)}`
            : ''
        }
        defaultVersion={
          loadedModel && loadedModel.stage === 'planning' && !isBuiltInSampleModel(loadedModel) && !loadedModel.isDraft
            ? loadedModel.version || '1.0.0'
            : '1.0.0'
        }
        defaultOverwriteId={
          loadedModel && loadedModel.stage === 'planning' && !isBuiltInSampleModel(loadedModel) && !loadedModel.isDraft
            ? loadedModel.id
            : null
        }
        existingRevisions={savedSourceModels
          .filter(m => !isBuiltInSampleModel(m) && m.stage === 'planning')
          .map(m => ({ id: m.id, name: m.name, description: m.description, version: m.version }))}
        onSave={handleSaveArtifact}
      />

      <DeleteModelConfirmModal
        isOpen={showDeleteConfirm && Boolean(pendingModel)}
        onClose={() => setShowDeleteConfirm(false)}
        title="Delete Model"
        modelName={pendingModel?.name || ''}
        onConfirm={async () => {
          if (!pendingModel) return;
          const deletedName = cleanModelDisplayName(pendingModel.name);
          const deletedId = pendingModel.id;
          await deleteSourceModel(deletedId);
          const remaining = allModels.filter(m => m.id !== deletedId);
          setPendingModelId(remaining[0]?.id || '');
          if (loadedModelId === deletedId || loadedModel?.id === deletedId) {
            setLoadedModelId('');
            setLoadedModel(null);
            setIsModelLoaded(false);
            setTunedNodes([]);
            setTunedNlbs([]);
            setHasUnsavedChanges(false);
            setActiveStep(1);
          }
          setToastTone('success');
          setToastMsg(`Model "${deletedName}" deleted.`);
        }}
      />

      <UnsavedConfirmModal
        isOpen={showUnsavedNextModal}
        onStay={() => setShowUnsavedNextModal(false)}
        onProceed={() => {
          const currentPlanDraft = useMigrationStore.getState().unsavedPlanningDraft;
          if (currentPlanDraft) {
            selectSourceModel(currentPlanDraft);
          }
          if (onNext) onNext();
        }}
      />

      <UnsavedConfirmModal
        isOpen={showOverwriteLoadModal}
        onStay={() => setShowOverwriteLoadModal(false)}
        onProceed={executeLoadModel}
        description="You have unsaved edits in the current editor. Loading another model will overwrite your current unsaved changes. Continue?"
        stayLabel="Cancel"
        proceedLabel="Load and Overwrite"
      />

      {/* Source Infrastructure Summary Modal (POST /beetle/summary/source) */}
      {showSourceSummaryModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-5xl border border-border-main animate-scale-up space-y-4 max-h-[90vh] flex flex-col">
            <div className="flex flex-wrap justify-between items-center gap-3 border-b border-border-main/20 pb-3 shrink-0">
              <div className="flex flex-wrap items-center gap-3">
                <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
                  <FileText className="w-5 h-5 text-emerald-500" />
                  <span>Source Infrastructure Summary (POST /summary/source)</span>
                </h3>
                <div className="flex items-center bg-bg-panel/80 p-0.5 rounded-lg border border-border-main text-xs font-mono">
                  {(['html', 'md', 'json'] as const).map((fmt) => (
                    <button
                      key={fmt}
                      type="button"
                      onClick={() => handleFetchSourceSummary(fmt)}
                      className={`px-2.5 py-1 rounded-md font-extrabold uppercase transition cursor-pointer ${
                        sourceSummaryFormat === fmt
                          ? 'bg-emerald-500 text-slate-950 shadow-sm'
                          : 'text-text-muted hover:text-text-main'
                      }`}
                    >
                      {fmt}
                    </button>
                  ))}
                </div>
              </div>

              <div className="flex items-center gap-2">
                <button
                  type="button"
                  onClick={handleDownloadSourceSummary}
                  disabled={!sourceSummaryContent || isLoadingSourceSummary}
                  className="px-3.5 py-1.5 bg-emerald-500/15 hover:bg-emerald-500/25 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 rounded-lg text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer disabled:opacity-40"
                >
                  <Download className="w-3.5 h-3.5" />
                  <span>Download ({sourceSummaryFormat.toUpperCase()})</span>
                </button>
                <button
                  type="button"
                  onClick={() => setShowSourceSummaryModal(false)}
                  className="text-text-muted hover:text-text-main transition p-1 rounded-lg cursor-pointer"
                >
                  <X className="w-5 h-5" />
                </button>
              </div>
            </div>

            <div className="flex-1 overflow-y-auto rounded-xl border border-border-main/40 bg-white min-h-[450px]">
              {isLoadingSourceSummary ? (
                <div className="p-12 text-center space-y-2 bg-bg-input h-full flex flex-col items-center justify-center">
                  <Loader2 className="w-6 h-6 animate-spin text-emerald-500" />
                  <p className="text-sm font-mono text-text-muted">Generating Source Infrastructure Summary ({sourceSummaryFormat.toUpperCase()})...</p>
                </div>
              ) : sourceSummaryFormat === 'html' ? (
                <iframe
                  title="Source Infrastructure Summary"
                  srcDoc={sourceSummaryContent
                    .replace(/background-color:\s*#f5f5f5/gi, 'background-color: #ffffff')
                    .replace(/background:\s*#f5f5f5/gi, 'background: #ffffff')}
                  className="w-full h-[650px] border-0 bg-white"
                  sandbox="allow-same-origin"
                />
              ) : (
                <div className="p-5 bg-slate-50 h-full overflow-y-auto">
                  <pre className="text-xs font-mono text-slate-800 whitespace-pre-wrap break-words leading-relaxed">
                    {sourceSummaryContent}
                  </pre>
                </div>
              )}
            </div>
          </div>
        </div>
      )}

      <ModelToast
        message={toastMsg}
        tone={toastTone}
        onClose={() => setToastMsg(null)}
      />

    </div>
  );
};
