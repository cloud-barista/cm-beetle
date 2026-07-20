'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '../../store/migrationStore';
import { TopologyMap } from './TopologyMap';
import { OnpremNode, OnpremInfra, OnpremModelEnvelope } from '../../types/migration';
import { Sparkles, GitBranch, Save, Layers, DollarSign, RefreshCw, Network, Server, Sliders, Cpu, ChevronDown, ChevronUp, Copy, HardDrive, X, FileText, Trash2, Loader2 } from 'lucide-react';
import { SaveRevisionModal } from '../common/SaveRevisionModal';



export const CloudInfraOptimizer: React.FC = () => {
  const {
    savedSourceModels,
    selectedSourceModel,
    desiredCsp,
    desiredRegion,
    recommendationCandidates,
    selectedCandidateIndex,
    editedCandidate,
    fetchSavedSourceModels,
    selectSourceModel,
    setDesiredCsp,
    setDesiredRegion,
    triggerRecommendation,
    selectCandidate,
    updateEditedCandidate,
    saveCloudModel,
    updateCloudModel,
    deleteCloudModel,
    savedCloudModels,
    selectedCloudModel,
    selectCloudModel,
    fetchSavedCloudModels,
    tumblebugProviders,
    tumblebugRegions,
    fetchTumblebugProviders,
    fetchTumblebugRegions
  } = useMigrationStore();

  const [activeTunedNodeId, setActiveTunedNodeId] = useState<string>('');
  const [showCompareModal, setShowCompareModal] = useState(false);
  const [showSaveTargetModal, setShowSaveTargetModal] = useState(false);

  // Tuned nodes state for left spec editor
  const [tunedNodes, setTunedNodes] = useState<OnpremNode[]>([]);
  const [tuningSourceSaveSuccess, setTuningSourceSaveSuccess] = useState(false);
  const [isJsonOpen, setIsJsonOpen] = useState(false);
  const [tunedNetwork, setTunedNetwork] = useState<any>(null);
  const [newCidr, setNewCidr] = useState('');
  const [activeStep, setActiveStep] = useState<number>(1);
  const [isModelLoaded, setIsModelLoaded] = useState(false);
  const [recommendMode, setRecommendMode] = useState<'new' | 'saved'>('new');
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const [deleteError, setDeleteError] = useState('');
  const [deleteConfirmText, setDeleteConfirmText] = useState('');

  // Excluded node IDs list for target recommendation filters
  const [excludedNodeIds, setExcludedNodeIds] = useState<string[]>([]);

  // Firewall rule creator form states
  const [newRuleDir, setNewRuleDir] = useState('inbound');
  const [newRuleProto, setNewRuleProto] = useState('tcp');
  const [newRulePort, setNewRulePort] = useState('');
  const [newRuleCidr, setNewRuleCidr] = useState('0.0.0.0/0');

  // Target Cloud Resource Editor tab states
  const [targetActiveTab, setTargetActiveTab] = useState<'network' | 'sshkey' | 'security' | 'compute' | 'storage' | 'nlb'>('network');
  // Track which node group cards are expanded in Topology (key = "subIdx-ngIdx")
  const [expandedNgCards, setExpandedNgCards] = useState<Record<string, boolean>>({});
  const toggleNgCard = (key: string) =>
    setExpandedNgCards(prev => ({ ...prev, [key]: !prev[key] }));
  const [tgtRuleDir, setTgtRuleDir] = useState('inbound');
  const [tgtRuleProto, setTgtRuleProto] = useState('tcp');
  const [tgtRulePort, setTgtRulePort] = useState('');
  const [tgtRuleCidr, setTgtRuleCidr] = useState('0.0.0.0/0');


  useEffect(() => {
    fetchSavedSourceModels();
    fetchSavedCloudModels();
    fetchTumblebugProviders();
  }, []);

  // Auto-load tunedNodes from selectedSourceModel when tab is entered
  // (source model selection & tuning now happens in SourceInfraRefinement)
  useEffect(() => {
    if (selectedSourceModel?.onpremiseInfraModel && tunedNodes.length === 0) {
      setTunedNodes(JSON.parse(JSON.stringify(selectedSourceModel.onpremiseInfraModel.nodes || [])));
      setTunedNetwork(JSON.parse(JSON.stringify(selectedSourceModel.onpremiseInfraModel.network || { ipv4Networks: {}, ipv6Networks: {} })));
      if (selectedSourceModel.onpremiseInfraModel.nodes && selectedSourceModel.onpremiseInfraModel.nodes.length > 0) {
        setActiveTunedNodeId(selectedSourceModel.onpremiseInfraModel.nodes[0]?.machineId || '');
      }
    }
  }, [selectedSourceModel]);





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

  const activeNode = tunedNodes.find((n) => n.machineId === activeTunedNodeId);

  // Exclude/Include node toggle
  const handleToggleNodeExclude = (machineId: string) => {
    setExcludedNodeIds(prev =>
      prev.includes(machineId)
        ? prev.filter(id => id !== machineId)
        : [...prev, machineId]
    );
  };

  // Firewall Rule deletion
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

  // Firewall Rule addition
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

  const handleSaveTunedSourceModel = async () => {
    if (!selectedSourceModel) return;
    try {
      const filteredNodes = tunedNodes.filter(n => !excludedNodeIds.includes(n.machineId));
      const currentVer = parseFloat(selectedSourceModel.version || '1.0');
      const nextVer = (currentVer + 0.1).toFixed(1);

      const updatedInfra = {
        ...selectedSourceModel.onpremiseInfraModel,
        nodes: filteredNodes,
        network: tunedNetwork || selectedSourceModel.onpremiseInfraModel.network
      };

      const updatedModel: OnpremModelEnvelope = {
        ...selectedSourceModel,
        onpremiseInfraModel: updatedInfra,
        version: nextVer.toString(),
        updatedTime: new Date().toISOString()
      };

      const updatedList = savedSourceModels.map(m => m.id === selectedSourceModel.id ? updatedModel : m);
      useMigrationStore.setState({
        savedSourceModels: updatedList,
        selectedSourceModel: updatedModel
      });

      setTunedNodes(filteredNodes);
      setExcludedNodeIds([]);
      setTuningSourceSaveSuccess(true);
      setActiveStep(1); // Advance to Step 3: Desired Cloud Target Specification
      setTimeout(() => setTuningSourceSaveSuccess(false), 2000);
    } catch (err) {
      console.error(err);
    }
  };

  const handleRecommend = async () => {
    const filteredNodes = tunedNodes.filter(n => !excludedNodeIds.includes(n.machineId));
    if (filteredNodes.length === 0) return;
    const sourceInfra: OnpremInfra = {
      nodes: filteredNodes,
      network: tunedNetwork || selectedSourceModel?.onpremiseInfraModel?.network || {
        ipv4Networks: {},
        ipv6Networks: {}
      },
      nlbs: selectedSourceModel?.onpremiseInfraModel?.nlbs || []
    };
    await triggerRecommendation(sourceInfra);
    setActiveStep(2);
  };

  const handleLoadDesign = () => {
    if (!selectedCloudModel) return;
    setActiveStep(2);
  };

  const handleDeleteDesign = () => {
    if (!selectedCloudModel || selectedCloudModel.id === 'cloud-demo-1') return;
    setDeleteConfirmText('');
    setDeleteError('');
    setIsDeleting(false);
    setShowDeleteConfirm(true);
  };

  const handleConfirmDeleteDesign = async () => {
    if (!selectedCloudModel) return;
    setIsDeleting(true);
    setDeleteError('');
    try {
      await deleteCloudModel(selectedCloudModel.id);
      setShowDeleteConfirm(false);
      setDeleteConfirmText('');
      setActiveStep(1);
    } catch (err: any) {
      setDeleteError(err.message || 'Failed to delete design');
    } finally {
      setIsDeleting(false);
    }
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

  // ── Spec / Image detail helpers ──────────────────────────────────────────
  // specId format: "csp+region+instanceType" → extract instance type portion
  const extractInstanceType = (specId: string) => specId.split('+').pop() || specId;

  // Lookup SpecInfo from targetSpecList by specId (try id, name, cspSpecName)
  const getSpecInfo = (candidate: typeof editedCandidate, specId: string) => {
    if (!candidate?.targetSpecList) return null;
    return candidate.targetSpecList.find((s: any) =>
      s.id === specId || s.name === specId || s.cspSpecName === extractInstanceType(specId)
    ) ?? (candidate.targetSpecList.length === 1 ? candidate.targetSpecList[0] : null);
  };

  // Lookup ImageInfo from targetOsImageList by imageId
  const getImageInfo = (candidate: typeof editedCandidate, imageId: string) => {
    if (!candidate?.targetOsImageList) return null;
    return candidate.targetOsImageList.find((img: any) =>
      img.id === imageId || img.cspImageId === imageId || img.cspImageName === imageId || img.name === imageId
    ) ?? (candidate.targetOsImageList.length === 1 ? candidate.targetOsImageList[0] : null);
  };

  const getEstimatedMonthlyCost = (candidate: typeof editedCandidate) => {
    if (!candidate?.targetInfra?.nodeGroups) return '0.00';
    let totalHourlyCost = 0;
    candidate.targetInfra.nodeGroups.forEach((ng) => {
      const spec = getSpecInfo(candidate, ng.specId);
      const costPerHour = spec ? (spec.costPerHour || 0) : 0;
      totalHourlyCost += costPerHour * (ng.nodeGroupSize || 0);
    });
    return (totalHourlyCost * 720).toFixed(2);
  };

  // Human-readable OS name: use imgInfo fields if available, otherwise extract last path component from imageId
  const formatOsName = (imgInfo: any, imageId?: string): string => {
    if (imgInfo?.osDistribution) return imgInfo.osDistribution;
    if (imgInfo?.osType)         return imgInfo.osType;
    // Fallback: extract last component from path-like imageId
    if (imageId) {
      if (imageId.includes('/')) return imageId.split('/').pop() || imageId;
      return imageId;
    }
    return '';
  };

  // Human-readable memory: show GiB or MiB
  const formatMemory = (gib: number) => {
    if (!gib && gib !== 0) return '—';
    return gib < 1 ? `${Math.round(gib * 1024)} MiB` : `${gib} GiB`;
  };
  // ─────────────────────────────────────────────────────────────────────────

  const handleTuneTargetNodeProperty = (ngIdx: number, key: string, value: any) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    if (key === 'nodeGroupSize') {
      const numericVal = parseInt(value, 10);
      updatedCandidate.targetInfra.nodeGroups[ngIdx].nodeGroupSize = isNaN(numericVal) ? 1 : Math.max(1, numericVal);
    } else if (key === 'rootDiskSize') {
      const numericVal = parseInt(value, 10);
      updatedCandidate.targetInfra.nodeGroups[ngIdx].rootDiskSize = isNaN(numericVal) ? 30 : Math.max(0, numericVal);
    } else {
      updatedCandidate.targetInfra.nodeGroups[ngIdx][key] = value;
    }
    updateEditedCandidate(updatedCandidate);
  };

  // 1. VNet & Subnet properties tuning
  const handleTuneTargetVNetProperty = (key: string, value: string) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));

    // VNet Name 변경 시 referential integrity 동기화
    if (key === 'name') {
      const oldVnetName = updatedCandidate.targetVNet.name;
      updatedCandidate.targetVNet.name = value;

      // 보안 그룹들의 vnetId 동기화
      if (updatedCandidate.targetSecurityGroupList) {
        updatedCandidate.targetSecurityGroupList.forEach((sg: any) => {
          if (sg.vnetId === oldVnetName || !sg.vnetId) {
            sg.vnetId = value;
          }
        });
      }

      // 컴퓨트 노드그룹들의 vnetId 동기화
      if (updatedCandidate.targetInfra?.nodeGroups) {
        updatedCandidate.targetInfra.nodeGroups.forEach((ng: any) => {
          if (ng.vnetId === oldVnetName || !ng.vnetId) {
            ng.vnetId = value;
          }
        });
      }
      // targetInfra 수준의 vnetId 동기화
      if (updatedCandidate.targetInfra) {
        updatedCandidate.targetInfra.vnetId = value;
      }
    } else {
      updatedCandidate.targetVNet[key] = value;
    }

    updateEditedCandidate(updatedCandidate);
  };

  const handleTuneTargetSubnetProperty = (subIdx: number, key: string, value: string) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    if (updatedCandidate.targetVNet.subnetInfoList && updatedCandidate.targetVNet.subnetInfoList[subIdx]) {
      const oldSubnetName = updatedCandidate.targetVNet.subnetInfoList[subIdx].name;
      updatedCandidate.targetVNet.subnetInfoList[subIdx][key] = value;

      // Subnet Name 변경 시 컴퓨트 노드그룹들의 subnetId 레퍼런스 동기화
      if (key === 'name' && updatedCandidate.targetInfra?.nodeGroups) {
        updatedCandidate.targetInfra.nodeGroups.forEach((ng: any) => {
          if (ng.subnetId === oldSubnetName) {
            ng.subnetId = value;
          }
        });
      }
    }
    updateEditedCandidate(updatedCandidate);
  };

  const handleAddSubnet = () => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    const n = updatedCandidate.targetVNet.subnetInfoList.length + 1;
    const vnetCidrBase = updatedCandidate.targetVNet.cidrBlock.split('.').slice(0, 2).join('.');
    updatedCandidate.targetVNet.subnetInfoList.push({
      name: `${updatedCandidate.targetVNet.name}-subnet-${n}`,
      ipv4_CIDR: `${vnetCidrBase}.${n}.0/24`,
      description: '',
      zone: '',
    });
    updateEditedCandidate(updatedCandidate);
  };

  const handleDeleteSubnet = (subIdx: number) => {
    if (!editedCandidate) return;
    if (editedCandidate.targetVNet.subnetInfoList.length <= 1) return; // keep minimum 1
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    const deleted = updatedCandidate.targetVNet.subnetInfoList[subIdx];
    updatedCandidate.targetVNet.subnetInfoList.splice(subIdx, 1);
    const fallback = updatedCandidate.targetVNet.subnetInfoList[0];
    // Reassign nodeGroups referencing the deleted subnet
    if (updatedCandidate.targetInfra?.nodeGroups) {
      updatedCandidate.targetInfra.nodeGroups.forEach((ng: any) => {
        if (ng.subnetId === deleted.name) ng.subnetId = fallback.name;
      });
    }
    updateEditedCandidate(updatedCandidate);
  };

  // NLB property tuning (supports dot-notation for nested fields, e.g. 'listener.port')
  const handleTuneTargetNlbProperty = (nlbIdx: number, path: string, value: any) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    if (!updatedCandidate.targetNlbList?.[nlbIdx]) return;
    const keys = path.split('.');
    let obj: any = updatedCandidate.targetNlbList[nlbIdx];
    for (let i = 0; i < keys.length - 1; i++) { obj = obj[keys[i]]; }
    obj[keys[keys.length - 1]] = value;
    updateEditedCandidate(updatedCandidate);
  };

  // 2. SSH Key property tuning
  const handleTuneTargetSshKeyProperty = (key: string, value: string) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    updatedCandidate.targetSshKey[key] = value;

    // SSH Key Name 변경 시 컴퓨트 노드그룹들의 sshKeyId 레퍼런스 동기화
    if (key === 'name' && updatedCandidate.targetInfra?.nodeGroups) {
      updatedCandidate.targetInfra.nodeGroups.forEach((ng: any) => {
        ng.sshKeyId = value;
      });
    }
    updateEditedCandidate(updatedCandidate);
  };

  // 3. Security Group properties tuning
  const handleTuneTargetSecurityGroupProperty = (sgIdx: number, key: string, value: string) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    if (updatedCandidate.targetSecurityGroupList && updatedCandidate.targetSecurityGroupList[sgIdx]) {
      const oldSgName = updatedCandidate.targetSecurityGroupList[sgIdx].name;
      updatedCandidate.targetSecurityGroupList[sgIdx][key] = value;

      // SG Name 변경 시 컴퓨트 노드그룹들의 securityGroupIds 레퍼런스 동기화
      if (key === 'name' && updatedCandidate.targetInfra?.nodeGroups) {
        updatedCandidate.targetInfra.nodeGroups.forEach((ng: any) => {
          if (ng.securityGroupIds) {
            ng.securityGroupIds = ng.securityGroupIds.map((sgId: string) =>
              sgId === oldSgName ? value : sgId
            );
          }
        });
      }
    }
    updateEditedCandidate(updatedCandidate);
  };

  const handleDeleteTargetFirewallRule = (sgIdx: number, ruleIdx: number) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    if (updatedCandidate.targetSecurityGroupList && updatedCandidate.targetSecurityGroupList[sgIdx]) {
      const rules = updatedCandidate.targetSecurityGroupList[sgIdx].firewallRules || [];
      updatedCandidate.targetSecurityGroupList[sgIdx].firewallRules = rules.filter((_: any, idx: number) => idx !== ruleIdx);
    }
    updateEditedCandidate(updatedCandidate);
  };

  const handleAddTargetFirewallRule = (sgIdx: number) => {
    if (!editedCandidate || !tgtRulePort) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    if (updatedCandidate.targetSecurityGroupList && updatedCandidate.targetSecurityGroupList[sgIdx]) {
      const rules = updatedCandidate.targetSecurityGroupList[sgIdx].firewallRules || [];
      const newRule = {
        Direction: tgtRuleDir,
        Protocol: tgtRuleProto.toUpperCase(),
        CIDR: tgtRuleCidr,
        Ports: tgtRulePort,
      };
      updatedCandidate.targetSecurityGroupList[sgIdx].firewallRules = [...rules, newRule];
    }
    updateEditedCandidate(updatedCandidate);
    setTgtRulePort('');
  };

  // Add a new empty Security Group to the candidate
  const handleAddSecurityGroup = () => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    const n = (updatedCandidate.targetSecurityGroupList || []).length + 1;
    const newSg = {
      name: `${updatedCandidate.targetVNet?.name || 'sg'}-sg-${n}`,
      vnetId: updatedCandidate.targetVNet?.name || '',
      description: '',
      firewallRules: [],
    };
    updatedCandidate.targetSecurityGroupList = [...(updatedCandidate.targetSecurityGroupList || []), newSg];
    updateEditedCandidate(updatedCandidate);
  };

  // Delete a Security Group and remove it from all node groups
  const handleDeleteSecurityGroup = (sgIdx: number) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    const deleted = updatedCandidate.targetSecurityGroupList[sgIdx];
    updatedCandidate.targetSecurityGroupList.splice(sgIdx, 1);
    if (updatedCandidate.targetInfra?.nodeGroups) {
      updatedCandidate.targetInfra.nodeGroups.forEach((ng: any) => {
        if (ng.securityGroupIds) {
          ng.securityGroupIds = ng.securityGroupIds.filter((id: string) => id !== deleted.name);
        }
      });
    }
    updateEditedCandidate(updatedCandidate);
  };

  // Add a Security Group to a specific node group
  const handleNodeGroupAddSg = (ngIdx: number, sgName: string) => {
    if (!editedCandidate || !sgName) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    const ng = updatedCandidate.targetInfra.nodeGroups[ngIdx];
    if (!ng.securityGroupIds) ng.securityGroupIds = [];
    if (!ng.securityGroupIds.includes(sgName)) ng.securityGroupIds.push(sgName);
    updateEditedCandidate(updatedCandidate);
  };

  // Remove a Security Group from a specific node group
  const handleNodeGroupRemoveSg = (ngIdx: number, sgName: string) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    const ng = updatedCandidate.targetInfra.nodeGroups[ngIdx];
    if (ng.securityGroupIds) {
      ng.securityGroupIds = ng.securityGroupIds.filter((id: string) => id !== sgName);
    }
    updateEditedCandidate(updatedCandidate);
  };

  const handleSaveToDamselfly = async (result: { name: string; description: string; version: string; overwriteId: string | null }) => {
    if (!editedCandidate) return;
    if (result.overwriteId) {
      await updateCloudModel(result.overwriteId, result.name, result.description, result.version, editedCandidate);
    } else {
      await saveCloudModel(result.name, result.description, result.version, editedCandidate);
    }
  };

  const { isRecommending } = useMigrationStore();

  return (
    <div className="space-y-8 mx-auto pb-24">

      {/* -------------------------------------------------------------
          STEP 1: Desired CSP and Region Selection
         ------------------------------------------------------------- */}
      <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${activeStep >= 1 ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
        <div className="flex items-center space-x-3 mb-4 border-b border-border-main/40 pb-3">
          <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${activeStep > 1 ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
            {activeStep > 1 ? '✓' : '1'}
          </span>
          <h3 className="text-base font-extrabold text-text-main">
            Step 1: Desired CSP and Region Selection
          </h3>
        </div>

        {/* Tab switch control */}
        <div className="flex border-b border-border-main/20 mb-6 max-w-4xl">
          <button
            onClick={() => { setRecommendMode('new'); selectCloudModel(null); }}
            className={`flex-1 pb-3 text-sm font-bold text-center border-b-2 transition-all cursor-pointer ${
              recommendMode === 'new'
                ? 'border-emerald-500 text-emerald-500'
                : 'border-transparent text-text-muted hover:text-text-main'
            }`}
          >
            Generate New Recommendation
          </button>
          <button
            onClick={() => { setRecommendMode('saved'); selectSourceModel(null); setTunedNodes([]); setIsModelLoaded(false); }}
            className={`flex-1 pb-3 text-sm font-bold text-center border-b-2 transition-all cursor-pointer ${
              recommendMode === 'saved'
                ? 'border-emerald-500 text-emerald-500'
                : 'border-transparent text-text-muted hover:text-text-main'
            }`}
          >
            Load Customized Cloud Infrastructure
          </button>
        </div>

        {recommendMode === 'new' ? (
          <>
            {/* Source Model selection */}
            <div className="mb-6 pb-6 border-b border-border-main/20">
              <label className="block text-sm font-bold text-emerald-600 dark:text-emerald-400 mb-2">Source Infrastructure Model</label>
              <div className="flex flex-col gap-2 max-w-4xl">
                <select
                  value={selectedSourceModel?.id || ''}
                  onChange={(e) => {
                    const m = savedSourceModels.find((x: any) => x.id === e.target.value) || null;
                    selectSourceModel(m);
                    setTunedNodes([]);
                    setIsModelLoaded(false);
                  }}
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-3 text-sm font-bold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                >
                  <option value="">-- Choose Source Model --</option>
                  {savedSourceModels.map((m: any) => (
                    <option key={m.id} value={m.id}>{m.name} (v{m.version || '1.0'})</option>
                  ))}
                </select>
                <button
                  onClick={handleLoadModel}
                  disabled={!selectedSourceModel}
                  className={`self-start px-5 py-2.5 rounded-xl text-sm font-extrabold flex items-center gap-2 transition cursor-pointer ${selectedSourceModel
                    ? 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 shadow-md shadow-emerald-500/25'
                    : 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'}`}
                >
                  <RefreshCw className="w-4 h-4" />
                  {isModelLoaded ? 'Model Loaded ✓' : 'Load Model'}
                </button>
              </div>
              {isModelLoaded && tunedNodes.length > 0 && (
                <p className="mt-2 text-xs text-emerald-600 dark:text-emerald-400 font-semibold">
                  {tunedNodes.length} node(s) ready for recommendation
                </p>
              )}
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-6 max-w-4xl">
              <div>
                <label className="block text-sm font-bold text-emerald-600 dark:text-emerald-400 mb-2">Desired CSP</label>
                <select
                  value={desiredCsp}
                  onChange={(e) => setDesiredCsp(e.target.value)}
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-3 text-sm font-bold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                >
                  {tumblebugProviders.map((csp: string) => {
                    const prettyCsp =
                      csp.toLowerCase() === 'aws' ? 'Amazon Web Services (AWS)' :
                        csp.toLowerCase() === 'azure' ? 'Microsoft Azure (Azure)' :
                          csp.toLowerCase() === 'gcp' ? 'Google Cloud Platform (GCP)' :
                            csp.toLowerCase() === 'alibaba' ? 'Alibaba Cloud (Alibaba)' :
                              csp.toLowerCase() === 'tencent' ? 'Tencent Cloud (Tencent)' :
                                csp.toLowerCase() === 'ibm' ? 'IBM Cloud (IBM)' :
                                  csp.toLowerCase() === 'ncp' || csp.toLowerCase() === 'ncloud' ? 'Naver Cloud Platform (NCP)' :
                                    csp.toLowerCase() === 'nhn' || csp.toLowerCase() === 'nhncloud' ? 'NHN Cloud (NHN)' :
                                      csp.toLowerCase() === 'kt' ? 'KT Cloud (KT)' :
                                        csp.toLowerCase() === 'openstack' ? 'OpenStack (OpenStack)' :
                                          csp.toLowerCase() === 'cloudit' ? 'Cloudit (Cloudit)' :
                                            csp.toLowerCase() === 'outscale' ? 'Outscale' : csp.toUpperCase();
                    return (
                      <option key={csp} value={csp}>
                        {prettyCsp}
                      </option>
                    );
                  })}
                </select>
              </div>
              <div>
                <label className="block text-sm font-bold text-emerald-600 dark:text-emerald-400 mb-2">Desired Region</label>
                <select
                  value={desiredRegion}
                  onChange={(e) => setDesiredRegion(e.target.value)}
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-3 text-sm font-bold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                >
                  {[...tumblebugRegions]
                    .sort((a, b) => a.id.localeCompare(b.id))
                    .map((r) => (
                      <option key={r.id} value={r.id}>
                        {r.name} ({r.id})
                      </option>
                    ))}
                </select>
              </div>
            </div>

            <div className="mt-6 flex justify-start">
              <button
                onClick={handleRecommend}
                disabled={isRecommending || tunedNodes.length === 0 || !desiredCsp || !desiredRegion}
                className="px-6 py-3.5 bg-gradient-to-r from-emerald-500 to-blue-600 hover:from-emerald-600 hover:to-blue-700 disabled:opacity-50 text-slate-950 font-extrabold rounded-xl text-sm tracking-wider transition shadow-lg shadow-emerald-500/10 flex items-center justify-center space-x-2 cursor-pointer"
              >
                {isRecommending ? (
                  <>
                    <RefreshCw className="w-4 h-4 animate-spin" />
                    <span>Recommending...</span>
                  </>
                ) : (
                  <>
                    <Sparkles className="w-4 h-4" />
                    <span>Recommend Target Cloud Infrastructure</span>
                  </>
                )}
              </button>
            </div>
          </>
        ) : (
          <div className="max-w-4xl space-y-4">
            <div>
              <label className="block text-sm font-bold text-emerald-600 dark:text-emerald-400 mb-2">Target Cloud Infrastructure Model</label>
              <select
                value={selectedCloudModel?.id || ''}
                onChange={(e) => {
                  const m = savedCloudModels.find((x: any) => x.id === e.target.value) || null;
                  selectCloudModel(m);
                }}
                className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-3 text-sm font-bold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
              >
                <option value="">-- Choose Target Cloud Infrastructure Model --</option>
                {savedCloudModels.map((m: any) => (
                  <option key={m.id} value={m.id}>{m.name} (v{m.version || '1.0.0'})</option>
                ))}
              </select>
            </div>
            <div className="flex gap-3 pt-2">
              <button
                onClick={handleLoadDesign}
                disabled={!selectedCloudModel}
                className={`px-5 py-2.5 rounded-xl text-sm font-extrabold flex items-center gap-2 transition cursor-pointer ${
                  selectedCloudModel
                    ? 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 shadow-md shadow-emerald-500/25'
                    : 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                }`}
              >
                <RefreshCw className="w-4 h-4" />
                <span>Load Design</span>
              </button>
              <button
                onClick={handleDeleteDesign}
                disabled={!selectedCloudModel || selectedCloudModel.id === 'cloud-demo-1'}
                className={`px-5 py-2.5 rounded-xl text-sm font-extrabold flex items-center gap-2 transition cursor-pointer ${
                  selectedCloudModel && selectedCloudModel.id !== 'cloud-demo-1'
                    ? 'bg-red-500 hover:bg-red-600 text-white shadow-md shadow-red-500/20'
                    : 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                }`}
              >
                <Trash2 className="w-4 h-4" />
                <span>Delete Design</span>
              </button>
            </div>
          </div>
        )}
      </div>

      {/* -------------------------------------------------------------
          STEP 2: Recommended Target Cloud Alternatives Review and Editing
         ------------------------------------------------------------- */}
      <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${activeStep >= 2 ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
        <div className="flex items-center space-x-3 mb-4 border-b border-border-main/40 pb-3">
          <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${activeStep > 2 ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
            {activeStep > 2 ? '✓' : '2'}
          </span>
          <h3 className="text-base font-extrabold text-text-main">
            Step 2: Recommended Target Cloud Alternatives Review and Editing
          </h3>
        </div>

        {recommendationCandidates.length === 0 ? (
          <div className="py-8 text-center text-text-muted text-sm italic">
            No recommendation generated yet. Complete previous steps to run recommendations.
          </div>
        ) : (
          <div className="space-y-6">

            {/* Row-based layout: Recommended Cloud Summary (Row 1) & Topology Visualization (Row 2) */}
            {editedCandidate && (
              <div className="flex flex-col space-y-6 pt-4 border-t border-border-main/20">

                {/* Row 1: Detailed specifications overview card (Recommended Cloud Summary) */}
                <div className="bg-bg-panel/30 border border-border-main/50 rounded-2xl p-5 space-y-4 w-full">
                  <div className="border-b border-border-main/30 pb-3 mb-2">
                    <span className="text-sm font-bold text-emerald-600 dark:text-emerald-400 block font-mono">Recommended Cloud Summary</span>
                  </div>

                  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">

                    {/* 1. Estimation (Match Level & Est. Cost) */}
                    <div className="bg-bg-panel/50 border border-border-main/20 p-4 rounded-xl font-mono flex flex-col justify-center">
                      <div className="space-y-2">
                        <span className="block text-sm font-bold text-emerald-500 font-sans border-b border-border-main/10 pb-1">Estimation</span>

                        <div className="flex flex-row justify-between items-center pt-2 min-h-[45px] gap-2">
                          <div className="flex items-center space-x-1.5">
                            <span className="text-xs text-text-muted font-bold font-sans">Match</span>
                            <span className="text-emerald-600 dark:text-emerald-400 font-extrabold text-xs uppercase bg-emerald-500/10 border border-emerald-500/20 px-2 py-0.5 rounded whitespace-nowrap">{editedCandidate.status}</span>
                          </div>
                          <div className="flex items-center space-x-1.5 border-l border-border-main/20 pl-3">
                            <span className="text-xs text-text-muted font-bold font-sans">Cost</span>
                            <span className="text-emerald-600 dark:text-emerald-400 font-extrabold text-lg font-mono whitespace-nowrap">
                              ${getEstimatedMonthlyCost(editedCandidate)}/month
                            </span>
                          </div>
                        </div>
                      </div>
                    </div>

                    {/* 2. Network (VNet & Subnets) */}
                    <div className="bg-bg-panel/50 border border-border-main/20 p-4 rounded-xl font-mono flex flex-col justify-center">
                      <div className="space-y-2">
                        <span className="block text-sm font-bold text-emerald-500 font-sans border-b border-border-main/10 pb-1">Network</span>

                        {/* 1 VNet(s) X Subnet(s) */}
                        <div className="text-lg font-extrabold text-text-main font-sans tracking-tight py-2">
                          1 VNet(s) {editedCandidate.targetVNet.subnetInfoList?.length || 0} Subnet(s)
                        </div>
                      </div>
                    </div>

                    {/* 3. Compute (Nodes) */}
                    <div className="bg-bg-panel/50 border border-border-main/20 p-4 rounded-xl font-mono flex flex-col justify-between">
                      <div className="space-y-2">
                        <span className="block text-sm font-bold text-emerald-500 font-sans border-b border-border-main/10 pb-1">Compute</span>

                        <div className="grid grid-cols-2 gap-2 items-start pt-2 min-h-[65px]">
                          {/* Left Side: Total Nodes Count */}
                          <div className="flex flex-col justify-center border-r border-border-main/20 pr-2">
                            <div className="text-lg font-extrabold text-text-main font-sans tracking-tight">
                              {editedCandidate.targetInfra.nodeGroups.reduce((acc, ng) => acc + ng.nodeGroupSize, 0)} Node(s)
                            </div>
                          </div>

                          {/* Right Side: Per-NodeGroup list */}
                          <div className="space-y-1.5 pl-2 max-h-[85px] overflow-y-auto w-full">
                            {editedCandidate.targetInfra.nodeGroups.map((ng, i) => (
                              <div key={i} className="bg-bg-panel border border-border-main/50 px-2 py-1 rounded-lg text-xs font-sans flex justify-between items-center space-x-2">
                                <span className="text-text-muted font-bold whitespace-nowrap">Node Group {i + 1}</span>
                                <span className="font-extrabold text-emerald-600 dark:text-emerald-400">{ng.nodeGroupSize} Nodes</span>
                              </div>
                            ))}
                          </div>
                        </div>
                      </div>
                    </div>

                    {/* 4. Security Groups */}
                    <div className="bg-bg-panel/50 border border-border-main/20 p-4 rounded-xl font-mono flex flex-col justify-between">
                      <div className="space-y-2">
                        <span className="block text-sm font-bold text-emerald-500 font-sans border-b border-border-main/10 pb-1">Security</span>

                        <div className="grid grid-cols-2 gap-2 items-start pt-2 min-h-[65px]">
                          {/* Left Side: Total SG Count */}
                          <div className="flex flex-col justify-center border-r border-border-main/20 pr-2">
                            <div className="text-lg font-extrabold text-text-main font-sans tracking-tight whitespace-nowrap">
                              {(editedCandidate.targetSecurityGroupList || []).length} Security Group(s)
                            </div>
                          </div>

                          {/* Right Side: Per-SG Rules list */}
                          <div className="space-y-1.5 pl-2 max-h-[85px] overflow-y-auto w-full">
                            {(editedCandidate.targetSecurityGroupList || []).map((sg, i) => (
                              <div key={i} className="bg-bg-panel border border-border-main/50 px-2 py-1 rounded-lg text-xs font-sans flex justify-between items-center space-x-1">
                                <span className="text-text-muted font-bold">SG {i + 1}</span>
                                <span className="font-extrabold text-emerald-600 dark:text-emerald-400">{(sg?.firewallRules || []).length} Rules</span>
                              </div>
                            ))}
                          </div>
                        </div>
                      </div>
                    </div>

                    {/* 5. SSH Key */}
                    <div className="bg-bg-panel/50 border border-border-main/20 p-4 rounded-xl font-mono flex flex-col justify-center">
                      <div className="space-y-2">
                        <span className="block text-sm font-bold text-emerald-500 font-sans border-b border-border-main/10 pb-1">SSH Key</span>

                        <div className="text-lg font-extrabold text-text-main font-sans tracking-tight py-2">
                          {editedCandidate.targetSshKey ? 1 : 0} SSH Key(s)
                        </div>
                      </div>
                    </div>

                    {/* 6. Load Balancer (NLB) */}
                    <div className="bg-bg-panel/50 border border-border-main/20 p-4 rounded-xl font-mono flex flex-col justify-center">
                      <div className="space-y-2">
                        <span className="block text-sm font-bold text-emerald-500 font-sans border-b border-border-main/10 pb-1">Load Balancer</span>

                        <div className="text-lg font-extrabold text-text-main font-sans tracking-tight py-2">
                          {(editedCandidate.targetNlbList || []).length} NLB(s)
                        </div>
                      </div>
                    </div>

                  </div>
                </div>

                {/* Row 2: Topology Visualization */}

                <div className="w-full bg-bg-panel/40 border border-border-main/50 rounded-2xl p-5 relative min-h-[300px] flex flex-col justify-between">
                  <div>
                    <div className="flex justify-between items-center mb-4">
                      <span className="text-sm font-bold text-emerald-600 dark:text-emerald-400">Topology Visualization</span>
                      <span className="text-sm text-text-muted font-mono">{editedCandidate.targetCloud.region} ({editedCandidate.targetCloud.csp.toUpperCase()})</span>
                    </div>
                    {/* Simulated Topology Drawing */}
                    <div className="flex flex-col space-y-4 pt-4 text-sm">

                      {/* VPC / VNet Container */}
                      <div className="border border-emerald-400 dark:border-emerald-800/40 bg-emerald-500/5 rounded-2xl p-5 relative">
                        <div className="flex flex-col md:flex-row justify-between md:items-center gap-3 mb-4 border-b border-emerald-200 dark:border-emerald-800/20 pb-3">
                          <span className="font-extrabold text-emerald-600 dark:text-emerald-400 flex items-center space-x-1.5 font-mono">
                            <Network className="w-4 h-4 animate-pulse" />
                            <span>VPC / VNet: {editedCandidate.targetVNet.name} ({editedCandidate.targetVNet.cidrBlock})</span>
                          </span>

                          {/* Associated SSH Key & Security Group */}
                          <div className="flex flex-wrap gap-2">
                            {editedCandidate.targetSshKey && (
                              <span className="bg-amber-100 dark:bg-yellow-950/40 border border-amber-300 dark:border-yellow-900/30 text-amber-600 dark:text-yellow-400 text-sm px-2 py-0.5 rounded font-extrabold flex items-center space-x-1 font-mono">
                                <HardDrive className="w-3 h-3" />
                                <span>Key: {editedCandidate.targetSshKey.name || 'default-key'}</span>
                              </span>
                            )}
                            {(editedCandidate.targetSecurityGroupList || []).map((sg, sgIdx) => {
                              if (!sg) return null;
                              return (
                                <span key={sgIdx} className="bg-orange-100 dark:bg-orange-950/40 border border-orange-300 dark:border-orange-900/30 text-orange-600 dark:text-orange-400 text-sm px-2 py-0.5 rounded font-extrabold flex items-center space-x-1 font-mono">
                                  <Sliders className="w-3 h-3" />
                                  <span>SG: {sg.name || 'default-sg'} ({(sg.firewallRules || []).length} Rules)</span>
                                </span>
                              );
                            })}
                          </div>
                        </div>

                        {/* NLB (If exists) - Located INSIDE the VPC container */}
                        {editedCandidate.targetNlbList && editedCandidate.targetNlbList.length > 0 && (
                          <div className="mb-4 relative z-10 flex flex-col items-center justify-center border-b border-border-main/20 pb-4">
                            {editedCandidate.targetNlbList.map((nlb, nlbIdx) => {
                              if (!nlb) return null;

                              // Find matching NodeGroup in target infrastructure to calculate simulated instances
                              // Fallback to first NodeGroup if name tuning mismatch occurs, or use original nodeGroupId
                              const matchingNg = editedCandidate.targetInfra.nodeGroups.find(
                                (ng) => ng.name === nlb.targetGroup?.nodeGroupId
                              ) || editedCandidate.targetInfra.nodeGroups[0];

                              const namePrefix = matchingNg?.name || nlb.targetGroup?.nodeGroupId || 'target-node';
                              const targetNodeCount = matchingNg?.nodeGroupSize || 1;
                              const targetNodesArray = Array.from({ length: targetNodeCount });

                              return (
                                <div key={nlbIdx} className="w-full max-w-lg bg-bg-panel border border-teal-300 dark:border-teal-800/40 rounded-xl p-4.5 space-y-3.5 shadow-md">
                                  <div className="flex justify-between items-center border-b border-border-main/40 pb-2">
                                    <span className="font-extrabold text-teal-900 dark:text-teal-300 text-sm flex items-center">
                                      <span className="w-2 h-2 bg-teal-500 rounded-full mr-2 animate-pulse" />
                                      Target Managed NLB
                                    </span>
                                    <span className="text-xs font-bold px-2 py-0.5 bg-teal-100 dark:bg-teal-950/40 text-teal-600 dark:text-teal-400 rounded-full border border-teal-200 dark:border-teal-800/40 uppercase">
                                      {nlb.type || 'PUBLIC'} Mode
                                    </span>
                                  </div>

                                  <div className="space-y-3 font-mono text-sm text-text-muted font-normal text-left">
                                    {/* Traffic Ingress */}
                                    <div className="flex items-center space-x-2 text-text-muted">
                                      <span className="font-sans font-normal shrink-0">Traffic Ingress ➔</span>
                                      <span className="text-text-main font-bold">
                                        Listener Port: {nlb.listener?.port || 'ALL'}
                                      </span>
                                    </div>

                                    {/* Balancing Method */}
                                    <div className="pl-3.5 border-l-2 border-teal-500/30 py-0.5 space-y-2">
                                      <div className="text-sm text-teal-600 dark:text-teal-400 font-bold font-sans">
                                        ▼ Target Routing Group
                                      </div>
                                      <div className="text-text-muted">
                                        Target NodeGroup: <span className="text-text-main font-bold">{namePrefix}</span>
                                      </div>
                                      {/* Target VMs */}
                                      <div className="space-y-1.5 pl-3 mt-1 border-l border-border-main/50 text-sm">
                                        {targetNodesArray.map((_, nodeIdx) => {
                                          const suffix = String(nodeIdx + 1).padStart(2, '0');
                                          const instanceName = `${namePrefix}-${suffix}`;
                                          const isLast = nodeIdx === targetNodesArray.length - 1;
                                          return (
                                            <div key={nodeIdx} className="flex justify-between items-center">
                                              <span>{isLast ? '└─' : '├─'} {instanceName}</span>
                                              <span className="text-text-main font-bold">Port: {nlb.targetGroup?.port || 'ALL'}</span>
                                            </div>
                                          );
                                        })}
                                      </div>
                                    </div>

                                    {/* Description Info */}
                                    {nlb.description && (
                                      <div className="pt-2 border-t border-border-main/20 text-sm text-text-muted font-sans font-semibold">
                                        Info: {nlb.description}
                                      </div>
                                    )}
                                  </div>
                                </div>
                              );
                            })}
                          </div>
                        )}

                        {/* Dynamic Subnet Containers — one per subnet in targetVNet.subnetInfoList */}
                        {(editedCandidate.targetVNet.subnetInfoList?.length > 0
                          ? editedCandidate.targetVNet.subnetInfoList
                          : [{ name: 'default-subnet', ipv4_CIDR: '10.0.1.0/24', description: '', zone: '' }]
                        ).map((sub, subIdx) => {
                          // NodeGroups assigned to this subnet; unassigned ones go to first subnet
                          const assignedNgs = editedCandidate.targetInfra.nodeGroups.filter(
                            ng => ng.subnetId === sub.name
                          );
                          const allSubnetNames = editedCandidate.targetVNet.subnetInfoList?.map(s => s.name) || [];
                          const unassignedNgs = editedCandidate.targetInfra.nodeGroups.filter(
                            ng => !allSubnetNames.includes(ng.subnetId)
                          );
                          const displayNgs = assignedNgs.length > 0
                            ? assignedNgs
                            : subIdx === 0 ? unassignedNgs : [];

                          return (
                        <div key={subIdx} className="border border-dashed border-emerald-400 dark:border-emerald-800/30 bg-bg-panel/40 rounded-xl p-4 space-y-3">
                          <div className="flex items-center text-sm text-text-muted font-mono pb-2 border-b border-border-main">
                            <span>
                              Subnet: {sub.name} ({sub.ipv4_CIDR})
                              {sub.zone && <span className="ml-2 text-xs text-emerald-600 dark:text-emerald-400/70">Zone: {sub.zone}</span>}
                            </span>
                          </div>

                          {/* Node Groups — accordion table, independent row expansion */}
                          {displayNgs.length === 0 ? (
                            <div className="text-xs text-text-muted italic py-2 text-center">No nodes assigned to this subnet</div>
                          ) : (
                          <div className="divide-y divide-border-main/20 rounded-xl border border-border-main/30 overflow-hidden">
                            {displayNgs.map((ng, idx) => {
                              const cardKey = `${subIdx}-${idx}`;
                              const isExpanded = !!expandedNgCards[cardKey];
                              const nodeCount = ng.nodeGroupSize || 1;
                              const nodesArray = Array.from({ length: nodeCount });

                              const specInfo = getSpecInfo(editedCandidate, ng.specId);
                              const vcpu    = specInfo?.vCPU;
                              const memGiB  = specInfo?.memoryGiB;
                              const imgInfo = getImageInfo(editedCandidate, ng.imageId);
                              const osName  = formatOsName(imgInfo, ng.imageId);

                              // Security groups for this node group
                              // Priority 1: use explicit securityGroupIds if set
                              // Priority 2: name-based heuristic for backwards compatibility
                              const ngSgs = (() => {
                                if (ng.securityGroupIds && ng.securityGroupIds.length > 0) {
                                  return (editedCandidate.targetSecurityGroupList || []).filter((sg: any) =>
                                    ng.securityGroupIds.includes(sg.name)
                                  );
                                }
                                // Heuristic: name match only (no index-based or blanket fallback)
                                return (editedCandidate.targetSecurityGroupList || []).filter((sg: any) => {
                                  if (!sg?.name) return false;
                                  const s = sg.name.toLowerCase(), n = ng.name.toLowerCase();
                                  return s.includes(n) || n.includes(s) || s.includes('default') || s.includes('common');
                                });
                              })();

                              return (
                                <div key={idx} className={isExpanded ? 'bg-bg-input/10' : 'bg-bg-panel/30 hover:bg-bg-input/10 transition'}>
                                  {/* Summary row — always visible */}
                                  <button
                                    onClick={() => toggleNgCard(cardKey)}
                                    className="w-full flex flex-wrap items-center gap-x-3 gap-y-1 px-3 py-2.5 text-left cursor-pointer"
                                  >
                                    <ChevronDown className={`w-3.5 h-3.5 text-text-muted flex-shrink-0 transition-transform ${isExpanded ? 'rotate-180' : ''}`} />

                                    {/* Node Group label + name — flex-1 to use available space */}
                                    <div className="flex items-center gap-1.5 flex-1 min-w-0 min-w-[180px]">
                                      <Layers className="w-3.5 h-3.5 text-emerald-600 dark:text-emerald-400 flex-shrink-0" />
                                      <div className="min-w-0">
                                        <span className="text-xs text-text-muted font-normal block leading-none mb-0.5">Node Group</span>
                                        <span className="text-sm font-bold text-text-main block" title={ng.name}>{ng.name}</span>
                                      </div>
                                    </div>

                                    {/* Spec: vCPU · Memory · instance type */}
                                    <div className="flex items-center gap-1 flex-shrink-0 flex-wrap">
                                      {vcpu   && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">{vcpu} vCPU</span>}
                                      {memGiB && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">{formatMemory(memGiB)}</span>}
                                      <span className="text-xs font-mono text-text-muted">{extractInstanceType(ng.specId)}</span>
                                    </div>

                                    {/* OS */}
                                    {osName && <span className="bg-teal-500/10 border border-teal-500/20 text-teal-600 dark:text-teal-400 text-xs font-extrabold px-1.5 py-0.5 rounded flex-shrink-0">{osName}</span>}

                                    {/* Disk */}
                                    <span className="text-xs font-mono text-text-muted flex-shrink-0">{ng.rootDiskSize} GB</span>

                                    {/* Security Groups */}
                                    {ngSgs.map((sg: any, sgIdx: number) => (
                                      <span key={sgIdx} className="bg-orange-50 dark:bg-orange-950/20 border border-orange-200 dark:border-orange-900/30 text-orange-600 dark:text-orange-400 text-xs px-1.5 py-0.5 rounded font-mono flex-shrink-0">
                                        {sg.name}
                                      </span>
                                    ))}

                                    {/* Node count — push to right */}
                                    <span className="ml-auto flex-shrink-0 text-xs font-bold bg-emerald-100 dark:bg-emerald-950/40 border border-emerald-300 dark:border-emerald-800/40 text-emerald-600 dark:text-emerald-400 px-1.5 py-0.5 rounded">×{nodeCount}</span>
                                  </button>

                                  {/* Expanded: detail summary + VM nodes */}
                                  {isExpanded && (
                                    <div className="border-t border-border-main/20 px-4 py-3 space-y-3 bg-bg-input/5">

                                      {/* Top summary: Spec · Image · Root Disk · SG · Size */}
                                      <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-3">
                                        {/* Spec */}
                                        <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                          <span className="block text-xs font-bold text-text-muted mb-1">Spec</span>
                                          <div className="flex flex-wrap gap-1">
                                            {vcpu   && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">{vcpu} vCPU</span>}
                                            {memGiB && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">{formatMemory(memGiB)}</span>}
                                          </div>
                                          <span className="text-xs font-mono text-text-muted mt-1 block" title={ng.specId}>{extractInstanceType(ng.specId)}</span>
                                        </div>

                                        {/* Image */}
                                        <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                          <span className="block text-xs font-bold text-text-muted mb-1">Image</span>
                                          {osName
                                            ? <span className="bg-teal-500/10 border border-teal-500/20 text-teal-600 dark:text-teal-400 text-xs font-extrabold px-1.5 py-0.5 rounded inline-block">{osName}</span>
                                            : <span className="text-xs font-mono text-text-muted break-all" title={ng.imageId}>{ng.imageId.split('/').pop() || ng.imageId}</span>
                                          }
                                        </div>

                                        {/* Root Disk */}
                                        <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                          <span className="block text-xs font-bold text-text-muted mb-1">Root Disk</span>
                                          <span className="text-sm font-extrabold text-text-main">{ng.rootDiskSize} GB</span>
                                        </div>

                                        {/* Security Group */}
                                        <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                          <span className="block text-xs font-bold text-text-muted mb-1">Security Group</span>
                                          <div className="flex flex-col gap-1">
                                            {ngSgs.length > 0
                                              ? ngSgs.map((sg: any, i: number) => (
                                                  <span key={i} className="bg-orange-50 dark:bg-orange-950/20 border border-orange-200 dark:border-orange-900/30 text-orange-600 dark:text-orange-400 text-xs px-1.5 py-0.5 rounded font-mono">{sg.name}</span>
                                                ))
                                              : <span className="text-xs text-text-muted">—</span>
                                            }
                                          </div>
                                        </div>

                                        {/* Size (node count) */}
                                        <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                          <span className="block text-xs font-bold text-text-muted mb-1">Size</span>
                                          <span className="text-2xl font-extrabold text-emerald-600 dark:text-emerald-400">{nodeCount}</span>
                                          <span className="text-xs text-text-muted ml-1">node{nodeCount > 1 ? 's' : ''}</span>
                                        </div>
                                      </div>

                                      {/* VM node list */}
                                      <div className="space-y-1">
                                        <span className="text-xs text-text-muted font-semibold">Nodes</span>
                                        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-1.5 max-h-40 overflow-y-auto">
                                          {nodesArray.map((_, nodeIdx) => {
                                            const suffix   = String(nodeIdx + 1).padStart(2, '0');
                                            const nodeName = `${ng.name}-${suffix}`;
                                            return (
                                              <div key={nodeIdx} className="bg-bg-panel border border-emerald-500/10 px-2.5 py-1.5 rounded-lg flex items-center gap-1.5 font-mono text-xs">
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
                          )}
                        </div>
                          );
                        })}
                      </div>

                      {/* Integrated Fine-Tuning & Saving Controls */}
                      <div className="mt-5 border-t border-border-main/30 pt-4 space-y-4">
                        <div className="flex flex-col space-y-4">
                          <div className="flex flex-col sm:flex-row justify-between sm:items-center gap-3 pb-2 border-b border-border-main/20">
                            <div>
                              <span className="block text-sm font-bold text-emerald-600 dark:text-emerald-400 font-mono">Review and Editing</span>
                              <span className="text-sm text-text-muted italic block mt-0.5">* Modifying resource values dynamically updates the topology diagram in real-time</span>
                            </div>
                            <div className="flex items-center gap-2 flex-wrap justify-end">
                              {editedCandidate && (
                                <span className="text-xs font-extrabold text-emerald-600 dark:text-emerald-400 font-mono bg-emerald-500/10 border border-emerald-500/20 px-2.5 py-1.5 rounded-lg flex items-center gap-1 shadow-sm">
                                  <DollarSign className="w-3.5 h-3.5 text-emerald-500" />
                                  <span>Cost: ${getEstimatedMonthlyCost(editedCandidate)}/month</span>
                                </span>
                              )}
                              {selectedCloudModel ? (
                                <span className="px-3.5 py-1.5 bg-cyan-500/10 border border-cyan-500/30 text-cyan-600 dark:text-cyan-400 rounded-xl text-xs font-extrabold font-mono">
                                  Loaded Design: {selectedCloudModel.name} (v{selectedCloudModel.version})
                                </span>
                              ) : (
                                <>
                                  <span className="text-xs font-bold text-text-muted">Alternatives:</span>
                                  {recommendationCandidates.map((c, idx) => {
                                    const isActive = selectedCandidateIndex === idx;
                                    return (
                                      <button
                                        key={idx}
                                        onClick={() => {
                                          selectCandidate(idx);
                                        }}
                                        className={`px-3 py-1.5 rounded-lg text-xs font-bold border transition cursor-pointer ${
                                          isActive
                                            ? 'bg-cyan-500/10 border-cyan-500/40 text-cyan-600 dark:text-cyan-400 font-extrabold'
                                            : 'bg-bg-panel border-border-main text-text-muted hover:text-text-main'
                                        }`}
                                      >
                                        Alt {idx + 1}
                                      </button>
                                    );
                                  })}
                                  <button
                                    onClick={() => setShowCompareModal(true)}
                                    className="px-3 py-1.5 bg-bg-panel border border-border-main hover:bg-cyan-500/10 hover:border-cyan-500/20 text-cyan-600 dark:text-cyan-400 rounded-lg text-xs font-bold transition cursor-pointer"
                                  >
                                    Compare
                                  </button>
                                </>
                              )}
                            </div>
                          </div>

                          {/* Resource Oriented Tuning Tabs Header */}
                          <div className="flex flex-wrap gap-1.5 border-b border-border-main/10 pb-2">
                            <button
                              onClick={() => setTargetActiveTab('network')}
                              className={`px-3 py-1.5 rounded-lg text-sm font-bold border transition cursor-pointer ${targetActiveTab === 'network'
                                  ? 'bg-emerald-500/10 border-emerald-500/40 text-emerald-600 dark:text-emerald-400 font-extrabold'
                                  : 'bg-bg-panel/40 border-border-main/30 text-text-muted hover:text-text-main'
                                }`}
                            >
                              Network (VNet & Subnets)
                            </button>
                            <button
                              onClick={() => setTargetActiveTab('compute')}
                              className={`px-3 py-1.5 rounded-lg text-sm font-bold border transition cursor-pointer ${targetActiveTab === 'compute'
                                  ? 'bg-emerald-500/10 border-emerald-500/40 text-emerald-600 dark:text-emerald-400 font-extrabold'
                                  : 'bg-bg-panel/40 border-border-main/30 text-text-muted hover:text-text-main'
                                }`}
                            >
                              Compute (Nodes)
                            </button>
                            <button
                              disabled
                              className="px-3 py-1.5 rounded-lg text-sm font-bold border border-border-main/10 bg-bg-panel/10 text-text-muted/40 cursor-not-allowed opacity-50 flex items-center space-x-1"
                              title="Storage Tuning is coming soon in next version"
                            >
                              <span>Disk (Soon)</span>
                            </button>
                            <button
                              onClick={() => setTargetActiveTab('security')}
                              className={`px-3 py-1.5 rounded-lg text-sm font-bold border transition cursor-pointer ${targetActiveTab === 'security'
                                  ? 'bg-emerald-500/10 border-emerald-500/40 text-emerald-600 dark:text-emerald-400 font-extrabold'
                                  : 'bg-bg-panel/40 border-border-main/30 text-text-muted hover:text-text-main'
                                }`}
                            >
                              Security Groups & Rules
                            </button>
                            <button
                              onClick={() => setTargetActiveTab('sshkey')}
                              className={`px-3 py-1.5 rounded-lg text-sm font-bold border transition cursor-pointer ${targetActiveTab === 'sshkey'
                                  ? 'bg-emerald-500/10 border-emerald-500/40 text-emerald-600 dark:text-emerald-400 font-extrabold'
                                  : 'bg-bg-panel/40 border-border-main/30 text-text-muted hover:text-text-main'
                                }`}
                            >
                              SSH Auth Key
                            </button>
                            {(editedCandidate.targetNlbList || []).length > 0 && (
                              <button
                                onClick={() => setTargetActiveTab('nlb')}
                                className={`px-3 py-1.5 rounded-lg text-sm font-bold border transition cursor-pointer ${targetActiveTab === 'nlb'
                                    ? 'bg-emerald-500/10 border-emerald-500/40 text-emerald-600 dark:text-emerald-400 font-extrabold'
                                    : 'bg-bg-panel/40 border-border-main/30 text-text-muted hover:text-text-main'
                                  }`}
                              >
                                NLB — Network Load Balancer
                              </button>
                            )}
                          </div>

                          {/* Tabs Content */}
                          <div className="space-y-3">

                            {/* TAB 1: Compute Resources (Node Groups) */}
                            {targetActiveTab === 'compute' && (
                              <div className="space-y-3.5">
                                {editedCandidate.targetInfra.nodeGroups.map((ng, ngIdx) => {
                                  const specInfo = getSpecInfo(editedCandidate, ng.specId);
                                  const imgInfo  = getImageInfo(editedCandidate, ng.imageId);
                                  const osName   = formatOsName(imgInfo, ng.imageId);
                                  const vcpu     = specInfo?.vCPU;
                                  const memGiB   = specInfo?.memoryGiB;
                                  return (
                                  <div key={ngIdx} className="bg-bg-panel/40 p-4 rounded-xl border border-border-main/30 space-y-3">
                                    <div className="flex justify-between items-center text-sm font-bold text-text-main border-b border-border-main/10 pb-1.5 font-mono">
                                      <span>Node Group #{ngIdx + 1} — {extractInstanceType(ng.specId)}</span>
                                      <span className="text-emerald-600 dark:text-emerald-400">Spec Match Type: Balanced</span>
                                    </div>

                                    {/* Spec & Image summary badges */}
                                    <div className="flex flex-wrap gap-3 p-3 bg-bg-input/20 rounded-lg border border-border-main/10">
                                      {/* Spec badges */}
                                      <div className="flex flex-col gap-1">
                                        <span className="text-xs font-bold text-text-muted">VM Spec</span>
                                        <div className="flex items-center gap-1.5 flex-wrap">
                                          {vcpu   && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-2 py-0.5 rounded">{vcpu} vCPU</span>}
                                          {memGiB && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-2 py-0.5 rounded">{formatMemory(memGiB)}</span>}
                                          <span className="text-xs text-text-muted font-mono" title={ng.specId}>{extractInstanceType(ng.specId)}</span>
                                        </div>
                                      </div>
                                      <div className="w-px bg-border-main/30 self-stretch mx-1 hidden sm:block" />
                                      {/* Image badges — flex-1 so it uses remaining width without truncation */}
                                      <div className="flex flex-col gap-1 flex-1 min-w-[180px]">
                                        <span className="text-xs font-bold text-text-muted">OS Image</span>
                                        <div className="flex flex-wrap gap-1.5 items-start">
                                          {osName && <span className="bg-teal-500/10 border border-teal-500/20 text-teal-600 dark:text-teal-400 text-xs font-extrabold px-2 py-0.5 rounded break-words">{osName}</span>}
                                          <span className="text-xs text-text-muted font-mono break-all" title={ng.imageId}>{ng.imageId.split('/').pop() || ng.imageId}</span>
                                        </div>
                                      </div>
                                    </div>

                                    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-6 gap-3.5">
                                      {/* Node Group Name */}
                                      <div>
                                        <label className="block text-sm font-bold text-text-muted mb-1 font-sans">Name</label>
                                        <input
                                          type="text"
                                          value={ng.name}
                                          onChange={(e) => handleTuneTargetNodeProperty(ngIdx, 'name', e.target.value)}
                                          className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                        />
                                      </div>
                                      {/* Spec ID */}
                                      <div>
                                        <label className="block text-sm font-bold text-text-muted mb-1 font-sans">Spec ID</label>
                                        <input
                                          type="text"
                                          value={ng.specId}
                                          onChange={(e) => handleTuneTargetNodeProperty(ngIdx, 'specId', e.target.value)}
                                          className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                        />
                                      </div>
                                      {/* Image ID */}
                                      <div>
                                        <label className="block text-sm font-bold text-text-muted mb-1 font-sans">Image ID</label>
                                        <input
                                          type="text"
                                          value={ng.imageId}
                                          onChange={(e) => handleTuneTargetNodeProperty(ngIdx, 'imageId', e.target.value)}
                                          className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                        />
                                      </div>
                                      {/* Root Disk Size */}
                                      <div>
                                        <label className="block text-sm font-bold text-text-muted mb-1 font-sans">Root Disk (GB)</label>
                                        <input
                                          type="number"
                                          min={10}
                                          value={ng.rootDiskSize}
                                          onChange={(e) => handleTuneTargetNodeProperty(ngIdx, 'rootDiskSize', e.target.value)}
                                          className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                        />
                                      </div>
                                      {/* Subnet Assignment */}
                                      <div>
                                        <label className="block text-sm font-bold text-text-muted mb-1 font-sans">Subnet</label>
                                        <select
                                          value={ng.subnetId}
                                          onChange={(e) => handleTuneTargetNodeProperty(ngIdx, 'subnetId', e.target.value)}
                                          className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40 cursor-pointer"
                                        >
                                          {(editedCandidate.targetVNet.subnetInfoList || []).map((sub, sIdx) => (
                                            <option key={sIdx} value={sub.name}>{sub.name}</option>
                                          ))}
                                        </select>
                                      </div>
                                      {/* Node Count — last column */}
                                      <div>
                                        <label className="block text-sm font-bold text-text-muted mb-1 font-sans">Node Count</label>
                                        <input
                                          type="number"
                                          min={1}
                                          value={ng.nodeGroupSize}
                                          onChange={(e) => handleTuneTargetNodeProperty(ngIdx, 'nodeGroupSize', e.target.value)}
                                          className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                        />
                                      </div>
                                    </div>

                                    {/* Security Group assignment */}
                                    <div className="border-t border-border-main/20 pt-3 space-y-2">
                                      <label className="block text-sm font-bold text-text-muted font-sans">Security Groups</label>
                                      <div className="flex flex-wrap gap-2 items-center min-h-[32px]">
                                        {/* Current SGs as removable pills */}
                                        {(ng.securityGroupIds || []).map((sgId: string, i: number) => (
                                          <span key={i} className="inline-flex items-center gap-1 bg-orange-50 dark:bg-orange-950/20 border border-orange-200 dark:border-orange-900/30 text-orange-600 dark:text-orange-400 text-xs px-2 py-1 rounded-lg font-mono">
                                            {sgId}
                                            <button
                                              onClick={() => handleNodeGroupRemoveSg(ngIdx, sgId)}
                                              className="ml-0.5 text-orange-400 hover:text-red-400 transition cursor-pointer leading-none"
                                              title={`Remove ${sgId}`}
                                            >
                                              ×
                                            </button>
                                          </span>
                                        ))}
                                        {/* Dropdown to add a SG not yet assigned */}
                                        {(editedCandidate.targetSecurityGroupList || []).filter(
                                          (sg: any) => !(ng.securityGroupIds || []).includes(sg.name)
                                        ).length > 0 && (
                                          <select
                                            defaultValue=""
                                            onChange={(e) => { if (e.target.value) { handleNodeGroupAddSg(ngIdx, e.target.value); e.target.value = ''; } }}
                                            className="bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2 py-1 text-xs font-mono focus:outline-none focus:border-emerald-500/40 cursor-pointer"
                                          >
                                            <option value="">+ Add SG</option>
                                            {(editedCandidate.targetSecurityGroupList || [])
                                              .filter((sg: any) => !(ng.securityGroupIds || []).includes(sg.name))
                                              .map((sg: any, i: number) => (
                                                <option key={i} value={sg.name}>{sg.name}</option>
                                              ))}
                                          </select>
                                        )}
                                        {(ng.securityGroupIds || []).length === 0 && (editedCandidate.targetSecurityGroupList || []).length === 0 && (
                                          <span className="text-xs text-text-muted italic">No security groups defined — add one in the Security Groups tab</span>
                                        )}
                                      </div>
                                    </div>
                                  </div>
                                  );
                                })}
                              </div>
                            )}

                            {/* TAB 2: Network Resources (VNet & Subnets) */}
                            {targetActiveTab === 'network' && (
                              <div className="bg-bg-panel/40 p-4 rounded-xl border border-border-main/30 space-y-4">
                                <div className="text-sm font-bold text-text-main border-b border-border-main/10 pb-1.5 font-mono">
                                  <span>VNet (Virtual Private Cloud) Configuration</span>
                                </div>
                                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                                  <div>
                                    <label className="block text-sm font-bold text-text-muted mb-1">VNet Resource Name</label>
                                    <input
                                      type="text"
                                      value={editedCandidate.targetVNet.name}
                                      onChange={(e) => handleTuneTargetVNetProperty('name', e.target.value)}
                                      className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-3 py-2 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                    />
                                  </div>
                                  <div>
                                    <label className="block text-sm font-bold text-text-muted mb-1">VNet CIDR Address Block</label>
                                    <input
                                      type="text"
                                      value={editedCandidate.targetVNet.cidrBlock}
                                      onChange={(e) => handleTuneTargetVNetProperty('cidrBlock', e.target.value)}
                                      className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-3 py-2 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                    />
                                  </div>
                                </div>

                                {/* Subnets list editing */}
                                <div className="space-y-3 mt-4">
                                  <div className="flex items-center justify-between">
                                    <span className="block text-sm font-bold text-text-muted font-mono">Subnet Resource Blocks</span>
                                    <button
                                      onClick={handleAddSubnet}
                                      className="px-3 py-1.5 bg-emerald-500 hover:bg-emerald-600 text-slate-950 rounded-lg text-xs font-bold transition cursor-pointer"
                                    >
                                      + Add Subnet
                                    </button>
                                  </div>
                                  {editedCandidate.targetVNet.subnetInfoList && editedCandidate.targetVNet.subnetInfoList.map((sub, subIdx) => (
                                    <div key={subIdx} className="bg-bg-input/20 border border-border-main/20 p-3 rounded-lg space-y-2.5">
                                      <div className="flex items-center justify-between mb-1">
                                        <span className="text-xs font-bold text-text-muted uppercase">Subnet #{subIdx + 1}</span>
                                        <button
                                          onClick={() => handleDeleteSubnet(subIdx)}
                                          disabled={editedCandidate.targetVNet.subnetInfoList.length <= 1}
                                          className="px-2 py-1 text-xs font-bold text-red-400 hover:bg-red-500/10 rounded disabled:opacity-30 disabled:cursor-not-allowed transition cursor-pointer"
                                          title={editedCandidate.targetVNet.subnetInfoList.length <= 1 ? 'At least 1 subnet required' : 'Delete subnet'}
                                        >
                                          Delete
                                        </button>
                                      </div>
                                      <div className="grid grid-cols-1 md:grid-cols-3 gap-3">
                                        <div>
                                          <label className="block text-xs font-bold text-text-muted mb-1">Name</label>
                                          <input
                                            type="text"
                                            value={sub.name}
                                            onChange={(e) => handleTuneTargetSubnetProperty(subIdx, 'name', e.target.value)}
                                            className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                          />
                                        </div>
                                        <div>
                                          <label className="block text-xs font-bold text-text-muted mb-1">CIDR Block</label>
                                          <input
                                            type="text"
                                            value={sub.ipv4_CIDR}
                                            onChange={(e) => handleTuneTargetSubnetProperty(subIdx, 'ipv4_CIDR', e.target.value)}
                                            className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                          />
                                        </div>
                                        <div>
                                          <label className="block text-xs font-bold text-text-muted mb-1">Zone <span className="text-text-muted font-normal normal-case">(optional)</span></label>
                                          <input
                                            type="text"
                                            placeholder="e.g. ap-northeast-2a"
                                            value={sub.zone || ''}
                                            onChange={(e) => handleTuneTargetSubnetProperty(subIdx, 'zone', e.target.value)}
                                            className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                          />
                                        </div>
                                      </div>
                                    </div>
                                  ))}
                                </div>
                              </div>
                            )}

                            {/* TAB 3: SSH Authentication Key */}
                            {targetActiveTab === 'sshkey' && (
                              <div className="bg-bg-panel/40 p-4 rounded-xl border border-border-main/30 space-y-4">
                                <div className="text-sm font-bold text-text-main border-b border-border-main/10 pb-1.5 font-mono">
                                  <span>SSH Credentials Resource Settings</span>
                                </div>
                                <div className="grid grid-cols-1 gap-4">
                                  <div>
                                    <label className="block text-sm font-bold text-text-muted mb-1">SSH Key Pair Name ID</label>
                                    <input
                                      type="text"
                                      value={editedCandidate.targetSshKey.name}
                                      onChange={(e) => handleTuneTargetSshKeyProperty('name', e.target.value)}
                                      className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-3 py-2 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                    />
                                  </div>
                                  <div>
                                    <label className="block text-sm font-bold text-text-muted mb-1">Description / Scope Tag</label>
                                    <input
                                      type="text"
                                      value={editedCandidate.targetSshKey.description || ''}
                                      onChange={(e) => handleTuneTargetSshKeyProperty('description', e.target.value)}
                                      className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-3 py-2 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                      placeholder="e.g., Key pair created for cluster authentication"
                                    />
                                  </div>
                                </div>
                              </div>
                            )}

                            {/* TAB 4: Security Groups & Rules */}
                            {targetActiveTab === 'security' && (
                              <div className="space-y-4">
                                {/* Add SG header */}
                                <div className="flex items-center justify-between pb-2 border-b border-border-main/20">
                                  <span className="text-sm font-semibold text-text-muted">
                                    {(editedCandidate.targetSecurityGroupList || []).length} Security Group(s)
                                  </span>
                                  <button
                                    onClick={handleAddSecurityGroup}
                                    className="px-3 py-1.5 bg-emerald-500 hover:bg-emerald-600 text-slate-950 rounded-lg text-xs font-bold transition cursor-pointer"
                                  >
                                    + Add Security Group
                                  </button>
                                </div>

                                {editedCandidate.targetSecurityGroupList && editedCandidate.targetSecurityGroupList.map((sg, sgIdx) => (
                                  <div key={sgIdx} className="bg-bg-panel/40 p-4 rounded-xl border border-border-main/30 space-y-4">
                                    <div className="flex flex-col md:flex-row justify-between md:items-center gap-3 border-b border-border-main/10 pb-2">
                                      <div className="flex-1">
                                        <label className="block text-sm font-bold text-text-muted mb-1 font-sans">Security Group Name</label>
                                        <input
                                          type="text"
                                          value={sg.name}
                                          onChange={(e) => handleTuneTargetSecurityGroupProperty(sgIdx, 'name', e.target.value)}
                                          className="bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1 text-sm font-mono w-full max-w-sm focus:outline-none focus:border-emerald-500/40 font-extrabold"
                                        />
                                      </div>
                                      <div className="flex items-center gap-2 self-start">
                                        <span className="text-sm text-text-muted font-semibold bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 px-2 py-0.5 rounded font-mono">
                                          {(sg.firewallRules || []).length} rules
                                        </span>
                                        <button
                                          onClick={() => handleDeleteSecurityGroup(sgIdx)}
                                          className="px-2.5 py-1 text-xs font-bold text-red-400 hover:bg-red-500/10 border border-red-400/30 hover:border-red-400/60 rounded-lg transition cursor-pointer"
                                          title="Delete this Security Group"
                                        >
                                          Delete SG
                                        </button>
                                      </div>
                                    </div>

                                    {/* Firewall rules — table list */}
                                    {(sg.firewallRules || []).length > 0 && (
                                      <div className="rounded-lg border border-border-main/30 overflow-hidden">
                                        {/* Header row */}
                                        <div className="grid grid-cols-[1fr_1fr_1fr_2.5fr_80px] bg-bg-input/40 px-5 py-2.5 text-xs font-bold text-text-muted border-b border-border-main/20 gap-6">
                                          <span>Direction</span>
                                          <span>Protocol</span>
                                          <span>Port Range</span>
                                          <span>CIDR (Source / Destination)</span>
                                          <span />
                                        </div>
                                        {/* Data rows */}
                                        <div className="divide-y divide-border-main/15">
                                          {(sg.firewallRules || []).map((rule, ruleIdx) => {
                                            if (!rule) return null;
                                            const dir       = (rule.Direction || rule.direction || 'inbound').toLowerCase();
                                            const isInbound = dir === 'inbound';
                                            const proto     = (rule.Protocol   || rule.protocol   || 'tcp').toLowerCase();
                                            const protocol  = proto.toUpperCase();
                                            const port      = rule.Ports || rule.dstPorts || rule.srcPorts || '*';
                                            const cidr      = rule.CIDR
                                              ? rule.CIDR
                                              : isInbound
                                                ? (rule.srcCIDR || rule.dstCIDR || '0.0.0.0/0')
                                                : (rule.dstCIDR || rule.srcCIDR || '0.0.0.0/0');

                                            return (
                                              <div key={ruleIdx} className="grid grid-cols-[1fr_1fr_1fr_2.5fr_80px] items-center px-5 py-3 gap-6 hover:bg-bg-input/10 transition text-sm font-mono">
                                                <span className={`inline-flex items-center px-2 py-0.5 rounded text-xs font-extrabold w-fit ${isInbound ? 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400' : 'bg-yellow-500/10 text-yellow-600 dark:text-yellow-400'}`}>
                                                  {dir.toUpperCase()}
                                                </span>
                                                <span className="text-text-muted">{protocol}</span>
                                                <span className="text-emerald-600 dark:text-emerald-400 font-extrabold">{port}</span>
                                                <span className="text-text-muted truncate" title={cidr}>
                                                  <span className="text-text-muted/50 text-xs mr-1">{isInbound ? 'Source:' : 'Destination:'}</span>{cidr}
                                                </span>
                                                <button
                                                  onClick={() => handleDeleteTargetFirewallRule(sgIdx, ruleIdx)}
                                                  className="text-xs font-bold text-red-400 hover:text-red-300 hover:bg-red-500/10 px-2 py-1 rounded transition cursor-pointer"
                                                >
                                                  Delete
                                                </button>
                                              </div>
                                            );
                                          })}
                                        </div>
                                      </div>
                                    )}

                                    {/* Add Target Firewall Rule form */}
                                    <div className="bg-bg-input/20 border border-border-main/20 p-3.5 rounded-lg space-y-3">
                                      <span className="block text-sm font-bold text-emerald-600 dark:text-emerald-400 font-mono">Add Firewall Access Rule</span>
                                      <div className="grid grid-cols-2 md:grid-cols-4 gap-3">
                                        <div>
                                          <label className="block text-sm font-bold text-text-muted mb-1">Direction</label>
                                          <select
                                            value={tgtRuleDir}
                                            onChange={(e) => setTgtRuleDir(e.target.value)}
                                            className="w-full bg-bg-input border border-border-main/40 text-text-main rounded-md px-2 py-1 text-sm focus:outline-none focus:border-emerald-500/40 cursor-pointer"
                                          >
                                            <option value="inbound">Inbound</option>
                                            <option value="outbound">Outbound</option>
                                          </select>
                                        </div>
                                        <div>
                                          <label className="block text-sm font-bold text-text-muted mb-1">Protocol</label>
                                          <select
                                            value={tgtRuleProto}
                                            onChange={(e) => setTgtRuleProto(e.target.value)}
                                            className="w-full bg-bg-input border border-border-main/40 text-text-main rounded-md px-2 py-1 text-sm focus:outline-none focus:border-emerald-500/40 cursor-pointer"
                                          >
                                            <option value="tcp">TCP</option>
                                            <option value="udp">UDP</option>
                                            <option value="icmp">ICMP</option>
                                            <option value="all">ALL</option>
                                          </select>
                                        </div>
                                        <div>
                                          <label className="block text-sm font-bold text-text-muted mb-1">Ports</label>
                                          <input
                                            type="text"
                                            value={tgtRulePort}
                                            onChange={(e) => setTgtRulePort(e.target.value)}
                                            className="w-full bg-bg-input border border-border-main/40 text-text-main rounded-md px-2 py-1 text-sm focus:outline-none focus:border-emerald-500/40 font-mono"
                                            placeholder="e.g., 80, 443, 22"
                                          />
                                        </div>
                                        <div>
                                          <label className="block text-sm font-bold text-text-muted mb-1">Source CIDR</label>
                                          <input
                                            type="text"
                                            value={tgtRuleCidr}
                                            onChange={(e) => setTgtRuleCidr(e.target.value)}
                                            className="w-full bg-bg-input border border-border-main/40 text-text-main rounded-md px-2 py-1 text-sm focus:outline-none focus:border-emerald-500/40 font-mono"
                                            placeholder="e.g., 0.0.0.0/0"
                                          />
                                        </div>
                                      </div>
                                      <button
                                        onClick={() => handleAddTargetFirewallRule(sgIdx)}
                                        className="px-3 py-1.5 bg-emerald-500 hover:bg-emerald-600 text-slate-950 rounded-lg text-sm font-bold transition cursor-pointer"
                                      >
                                        + Add Rule to SG
                                      </button>
                                    </div>
                                  </div>
                                ))}
                              </div>
                            )}

                            {/* TAB 5: NLB */}
                            {targetActiveTab === 'nlb' && (
                              <div className="space-y-3.5">
                                {(editedCandidate.targetNlbList || []).length === 0 ? (
                                  <div className="py-6 text-center text-text-muted text-sm italic">No NLBs in this candidate.</div>
                                ) : (editedCandidate.targetNlbList || []).map((nlb: any, nlbIdx: number) => (
                                  <div key={nlbIdx} className="bg-bg-panel/40 p-4 rounded-xl border border-border-main/30 space-y-3">
                                    <div className="flex justify-between items-center text-sm font-bold text-text-main border-b border-border-main/10 pb-1.5 font-mono">
                                      <span>NLB #{nlbIdx + 1}</span>
                                    </div>
                                    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3.5">
                                      {/* Type */}
                                      <div>
                                        <label className="block text-xs font-bold text-text-muted mb-1">Type</label>
                                        <select value={nlb.type || 'PUBLIC'} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'type', e.target.value)}
                                          className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40 cursor-pointer">
                                          <option value="PUBLIC">PUBLIC</option>
                                          <option value="INTERNAL">INTERNAL</option>
                                        </select>
                                      </div>
                                      {/* Scope */}
                                      <div>
                                        <label className="block text-xs font-bold text-text-muted mb-1">Scope</label>
                                        <select value={nlb.scope || 'REGION'} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'scope', e.target.value)}
                                          className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40 cursor-pointer">
                                          <option value="REGION">REGION</option>
                                          <option value="GLOBAL">GLOBAL</option>
                                        </select>
                                      </div>
                                      {/* Description */}
                                      <div>
                                        <label className="block text-xs font-bold text-text-muted mb-1">Description</label>
                                        <input type="text" value={nlb.description || ''} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'description', e.target.value)}
                                          className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40" />
                                      </div>
                                    </div>
                                    {/* Listener */}
                                    <div>
                                      <span className="block text-xs font-bold text-text-muted mb-2">Listener</span>
                                      <div className="grid grid-cols-2 gap-3">
                                        <div>
                                          <label className="block text-xs text-text-muted mb-1">Protocol</label>
                                          <select value={nlb.listener?.protocol || 'TCP'} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'listener.protocol', e.target.value)}
                                            className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40 cursor-pointer">
                                            <option value="TCP">TCP</option>
                                            <option value="UDP">UDP</option>
                                            <option value="HTTP">HTTP</option>
                                            <option value="HTTPS">HTTPS</option>
                                          </select>
                                        </div>
                                        <div>
                                          <label className="block text-xs text-text-muted mb-1">Port</label>
                                          <input type="text" value={nlb.listener?.port || ''} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'listener.port', e.target.value)}
                                            placeholder="e.g. 80" className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40" />
                                        </div>
                                      </div>
                                    </div>
                                    {/* Target Group */}
                                    <div>
                                      <span className="block text-xs font-bold text-text-muted mb-2">Target Group</span>
                                      <div className="grid grid-cols-1 sm:grid-cols-3 gap-3">
                                        <div>
                                          <label className="block text-xs text-text-muted mb-1">Node Group</label>
                                          <select value={nlb.targetGroup?.nodeGroupId || ''} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'targetGroup.nodeGroupId', e.target.value)}
                                            className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40 cursor-pointer">
                                            {editedCandidate.targetInfra.nodeGroups.map((ng: any, i: number) => (
                                              <option key={i} value={ng.name}>{ng.name}</option>
                                            ))}
                                          </select>
                                        </div>
                                        <div>
                                          <label className="block text-xs text-text-muted mb-1">Protocol</label>
                                          <select value={nlb.targetGroup?.protocol || 'TCP'} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'targetGroup.protocol', e.target.value)}
                                            className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40 cursor-pointer">
                                            <option value="TCP">TCP</option>
                                            <option value="UDP">UDP</option>
                                            <option value="HTTP">HTTP</option>
                                            <option value="HTTPS">HTTPS</option>
                                          </select>
                                        </div>
                                        <div>
                                          <label className="block text-xs text-text-muted mb-1">Port</label>
                                          <input type="text" value={nlb.targetGroup?.port || ''} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'targetGroup.port', e.target.value)}
                                            placeholder="e.g. 8080" className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40" />
                                        </div>
                                      </div>
                                    </div>
                                    {/* Health Checker */}
                                    <div>
                                      <span className="block text-xs font-bold text-text-muted mb-2">Health Checker</span>
                                      <div className="grid grid-cols-3 gap-3">
                                        <div>
                                          <label className="block text-xs text-text-muted mb-1">Interval (s)</label>
                                          <input type="number" min={1} value={nlb.healthChecker?.interval ?? 10} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'healthChecker.interval', parseInt(e.target.value))}
                                            className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40" />
                                        </div>
                                        <div>
                                          <label className="block text-xs text-text-muted mb-1">Timeout (s)</label>
                                          <input type="number" min={1} value={nlb.healthChecker?.timeout ?? 10} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'healthChecker.timeout', parseInt(e.target.value))}
                                            className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40" />
                                        </div>
                                        <div>
                                          <label className="block text-xs text-text-muted mb-1">Threshold</label>
                                          <input type="number" min={1} value={nlb.healthChecker?.threshold ?? 3} onChange={(e) => handleTuneTargetNlbProperty(nlbIdx, 'healthChecker.threshold', parseInt(e.target.value))}
                                            className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40" />
                                        </div>
                                      </div>
                                    </div>
                                  </div>
                                ))}
                              </div>
                            )}

                          </div>

                          <div className="flex flex-col md:flex-row justify-between items-center gap-4 pt-3 border-t border-border-main/20">
                            <span className="text-sm text-amber-600 dark:text-amber-400 font-sans">
                              ⚠️ Note: Resource creation may fail due to account-level CSP quota limits, regional availability, or insufficient instance stock. Verify your account quotas before deployment.
                            </span>
                            <button
                              onClick={() => setShowSaveTargetModal(true)}
                              className="w-full md:w-auto px-6 py-3 bg-gradient-to-r from-emerald-500 to-blue-600 hover:from-emerald-600 hover:to-blue-700 text-slate-950 rounded-xl text-sm font-extrabold flex items-center justify-center space-x-1.5 transition cursor-pointer shadow-lg shadow-emerald-500/10"
                            >
                              <Save className="w-4 h-4" />
                              <span>Save Target Cloud Infrastructure Model</span>
                            </button>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>

              </div>
            )}
          </div>
        )}
      </div>

      {/* Compare Candidates Modal */}
      {showCompareModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-4xl border border-border-main animate-scale-up flex flex-col max-h-[85vh]">

            {/* Modal Header */}
            <div className="flex justify-between items-center mb-4 border-b border-border-main/30 pb-3">
              <h3 className="text-base font-extrabold text-text-main flex items-center">
                <Sparkles className="w-5 h-5 text-emerald-600 dark:text-emerald-400 mr-2" />
                Side-by-Side Candidate Comparison Matrix
              </h3>
              <button
                onClick={() => setShowCompareModal(false)}
                className="text-text-muted hover:text-text-main transition p-1 hover:bg-bg-input rounded-lg cursor-pointer"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            {/* Scrollable Content Area */}
            <div className="overflow-y-auto pr-1 flex-1 space-y-4">
              <div className="overflow-x-auto border border-border-main/50 rounded-xl bg-bg-panel/40">
                <table className="w-full text-left border-collapse text-sm">
                  <thead>
                    <tr className="border-b border-border-main bg-bg-input/60 text-text-muted font-bold">
                      <th className="py-3 px-4 min-w-[180px]">Metric</th>
                      {recommendationCandidates.map((c, idx) => (
                        <th key={idx} className="py-3 px-4 min-w-[200px]">
                          Candidate {idx + 1}
                        </th>
                      ))}
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-border-main/40">
                    <tr>
                      <td className="py-3.5 px-4 font-bold text-text-muted bg-bg-input/10">Match Level</td>
                      {recommendationCandidates.map((c, idx) => (
                        <td key={idx} className="py-3.5 px-4">
                          <span className={`px-2.5 py-1 rounded-full text-sm font-bold uppercase ${c.status === 'highly-matched'
                              ? 'bg-green-500/10 text-green-600 dark:text-green-400 border border-green-500/20'
                              : 'bg-yellow-500/10 text-yellow-400 border border-yellow-500/20'
                            }`}>
                            {c.status}
                          </span>
                        </td>
                      ))}
                    </tr>
                    <tr>
                      <td className="py-3.5 px-4 font-bold text-text-muted bg-bg-input/10">Node Spec Allocation</td>
                      {recommendationCandidates.map((c, idx) => (
                        <td key={idx} className="py-3.5 px-4 text-text-main font-mono text-sm space-y-1">
                          {c.targetInfra.nodeGroups.map((ng, i) => (
                            <div key={i} className="flex justify-between border-b border-border-main/10 pb-1 last:border-0 last:pb-0">
                              <span className="text-text-muted">{ng.name}:</span>
                              <span className="font-bold text-emerald-600 dark:text-emerald-400">{ng.specId} (x{ng.nodeGroupSize})</span>
                            </div>
                          ))}
                        </td>
                      ))}
                    </tr>
                    <tr>
                      <td className="py-3.5 px-4 font-bold text-text-muted bg-bg-input/10">Estimated Cost</td>
                      {recommendationCandidates.map((c, idx) => (
                        <td key={idx} className="py-3.5 px-4 text-emerald-600 dark:text-emerald-400 font-extrabold font-mono">
                          <div className="flex items-center space-x-1">
                            <DollarSign className="w-3.5 h-3.5" />
                            <span>{getEstimatedMonthlyCost(c)}</span>
                            <span className="text-sm text-text-muted">/ month</span>
                          </div>
                        </td>
                      ))}
                    </tr>
                    <tr>
                      <td className="py-3.5 px-4 font-bold text-text-muted bg-bg-input/10">Description Summary</td>
                      {recommendationCandidates.map((c, idx) => (
                        <td key={idx} className="py-3.5 px-4 text-text-muted text-sm leading-relaxed max-w-[200px]">
                          {c.description}
                        </td>
                      ))}
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>

            {/* Modal Footer */}
            <div className="mt-4 flex justify-end border-t border-border-main/30 pt-3">
              <button
                onClick={() => setShowCompareModal(false)}
                className="px-5 py-2.5 bg-bg-panel border border-border-main hover:bg-emerald-500/10 hover:border-emerald-500/20 text-emerald-600 dark:text-emerald-400 rounded-xl text-sm font-bold cursor-pointer transition"
              >
                Close Comparison
              </button>
            </div>
          </div>
        </div>
      )}

      <SaveRevisionModal
        isOpen={showSaveTargetModal}
        onClose={() => setShowSaveTargetModal(false)}
        title="Save Target Cloud Infrastructure Model"
        defaultName={selectedCloudModel?.name || ''}
        defaultDescription={selectedCloudModel?.description || ''}
        defaultVersion={selectedCloudModel?.version || '1.0.0'}
        existingRevisions={savedCloudModels.map(m => ({ id: m.id, name: m.name, version: m.version }))}
        onSave={handleSaveToDamselfly}
        successMessage="Target Cloud Infrastructure Model saved to Damselfly Repository successfully."
      />

      {/* Delete Confirmation Modal */}
      {showDeleteConfirm && selectedCloudModel && (
        <div className="fixed inset-0 z-[60] flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up">
            <div className="flex justify-between items-center mb-4">
              <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                <Trash2 className="w-4 h-4 text-red-500" /> Delete Target Cloud Infrastructure Model
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
                Are you sure you want to delete the target cloud infrastructure model <strong className="text-text-main">"{selectedCloudModel.name}"</strong>? This action cannot be undone.
              </p>

              <div className="space-y-1.5 pt-1">
                <label className="block text-xs font-bold text-text-muted">
                  To confirm, type <span className="font-mono bg-bg-panel px-1 py-0.5 rounded border border-border-main/60 text-text-main select-all">{selectedCloudModel.name}</span> in the box below:
                </label>
                <input
                  type="text"
                  value={deleteConfirmText}
                  onChange={(e) => setDeleteConfirmText(e.target.value)}
                  placeholder="Type the design name to delete"
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
                  onClick={handleConfirmDeleteDesign}
                  disabled={isDeleting || deleteConfirmText !== selectedCloudModel.name}
                  className={`px-4 py-2 rounded-xl text-sm font-semibold transition flex items-center gap-1.5 ${
                    isDeleting || deleteConfirmText !== selectedCloudModel.name
                      ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                      : 'bg-red-500 hover:bg-red-600 text-white cursor-pointer shadow-md shadow-red-500/20'
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

    </div>
  );
};
