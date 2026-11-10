'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '../../store/migrationStore';
import { TopologyMap } from './TopologyMap';
import { MatchRateRadarChart } from './MatchRateRadarChart';
import { OnpremNode, OnpremInfra, OnpremModelEnvelope, RecommendedInfra, CloudModelEnvelope } from '../../types/migration';
import { Sparkles, GitBranch, Save, Layers, DollarSign, RefreshCw, Network, Server, Sliders, Cpu, ChevronDown, ChevronUp, Copy, HardDrive, X, FileText, Trash2, Loader2, Compass, ArrowRight, ArrowLeft, Plus, AlertTriangle, ShieldCheck, CheckCircle2, Globe, Target, Zap, Undo2 } from 'lucide-react';
import { SaveRevisionModal } from '../common/SaveRevisionModal';
import { ActiveModelBanner } from '../common/ActiveModelBanner';
import { DeleteModelConfirmModal } from '../common/DeleteModelConfirmModal';
import { UnsavedConfirmModal } from '../common/UnsavedConfirmModal';
import { ModelToast } from '../common/ModelToast';
import {
  buildDefaultModelName,
  cleanModelDisplayName,
  formatSourceOptionText,
  formatTargetOptionText,
  getStageBadgeMeta,
  isBuiltInSampleModel,
  resolveModelStageKind,
  resolveValidTargetInfraName,
} from '../../utils/modelFormat';
import { tumblebugApi, beetleApi, pickPreferredKoreaRegion, isKoreaRegion, PREFERRED_KOREA_REGIONS_BY_CSP } from '../../api/client';
import sampleGpuData from '../../data/sampleSourceGpuInfra.json';



export const CloudInfraOptimizer: React.FC<{ onNext?: () => void; onBack?: () => void }> = ({ onNext, onBack }) => {
  const {
    savedSourceModels,
    selectedSourceModel,
    desiredCsp,
    desiredRegion,
    targetPairs,
    recommendationMode,
    setRecommendationMode,
    cachedMultiCandidates,
    setCachedMultiCandidates,
    cachedMultiPairs,
    setCachedMultiPairs,
    recommendationLimit,
    setRecommendationLimit,
    minMatchRate,
    setMinMatchRate,
    recommendationCandidates,
    selectedCandidateIndex,
    editedCandidate,
    fetchSavedSourceModels,
    selectSourceModel,
    setDesiredCsp,
    setDesiredRegion,
    setTargetPairs,
    addTargetPair,
    removeTargetPair,
    updateTargetPair,
    triggerRecommendation,
    selectCandidate,
    updateEditedCandidate: storeUpdateEditedCandidate,
    saveCloudModel,
    updateCloudModel,
    deleteCloudModel,
    clearCustomCloudModels,
    savedCloudModels,
    selectedCloudModel,
    selectCloudModel,
    fetchSavedCloudModels,
    tumblebugProviders,
    tumblebugRegions,
    fetchTumblebugProviders,
    fetchTumblebugRegions,
    namespaceId
  } = useMigrationStore();

  const [isValidatingDryRun, setIsValidatingDryRun] = useState(false);
  const [dryRunResult, setDryRunResult] = useState<{
    open: boolean;
    valid?: boolean;
    issues: Array<{ code: string; severity: string; path: string; message: string }>;
    error?: string;
  }>({ open: false, issues: [] });

  const handleDryRunValidate = async () => {
    if (!editedCandidate) return;
    setIsValidatingDryRun(true);
    try {
      const res = await beetleApi.validateInfra(
        namespaceId || 'mig01',
        editedCandidate,
        true
      );
      setDryRunResult({
        open: true,
        valid: res.valid,
        issues: res.issues || [],
        error: res.error
      });
    } catch (err: any) {
      setDryRunResult({
        open: true,
        valid: false,
        issues: [],
        error: err.message || 'Validation request failed'
      });
    } finally {
      setIsValidatingDryRun(false);
    }
  };

  const [activeTunedNodeId, setActiveTunedNodeId] = useState<string>('');
  const [pendingSourceModelId, setPendingSourceModelId] = useState<string>(selectedSourceModel?.id || '');
  const [loadedSourceModelId, setLoadedSourceModelId] = useState<string>('');
  const [pendingCloudModelId, setPendingCloudModelId] = useState<string>(selectedCloudModel?.id || '');
  const [loadedCloudModelId, setLoadedCloudModelId] = useState<string>('');
  const [hasUnsavedChanges, setHasUnsavedChanges] = useState(false);
  const [showCompareModal, setShowCompareModal] = useState(false);
  const [showSaveTargetModal, setShowSaveTargetModal] = useState(false);
  const [showUnsavedNextModal, setShowUnsavedNextModal] = useState(false);
  const [showClearAllConfirm, setShowClearAllConfirm] = useState(false);
  const [toastMsg, setToastMsg] = useState<string | null>(null);
  const [toastTone, setToastTone] = useState<'success' | 'error'>('success');

  const updateEditedCandidate = (updated: RecommendedInfra) => {
    setHasUnsavedChanges(true);
    storeUpdateEditedCandidate(updated);
  };

  const executeProceedToMigration = () => {
    if (editedCandidate) {
      const isSavedModel = Boolean(selectedCloudModel && !String(selectedCloudModel.id).startsWith('candidate-'));
      if (isSavedModel && !hasUnsavedChanges && selectedCloudModel) {
        useMigrationStore.setState({ selectedCloudModel });
      } else {
        const modelName =
          (isSavedModel ? selectedCloudModel?.name : '') ||
          buildDefaultModelName('target', selectedSourceModel?.name, editedCandidate.targetCloud?.csp);
        useMigrationStore.setState({
          selectedCloudModel: {
            id: `candidate-${Date.now()}`,
            name: modelName,
            description:
              selectedCloudModel?.description ||
              `Optimized target cloud model for ${editedCandidate.targetCloud?.csp?.toUpperCase() || 'Cloud'} (${editedCandidate.targetCloud?.region || ''})`,
            cloudInfraModel: editedCandidate,
            version: selectedCloudModel?.version || '1.0.0',
            stage: 'target',
            updatedTime: new Date().toISOString(),
          },
        });
      }
    }
    if (onNext) {
      onNext();
    }
  };

  const handleProceedToMigration = () => {
    if (editedCandidate && hasUnsavedChanges) {
      setShowUnsavedNextModal(true);
      return;
    }
    executeProceedToMigration();
  };

  const [pairRegionsMap, setPairRegionsMap] = useState<Record<string, { id: string; name: string }[]>>({});

  // Tuned nodes state for left spec editor
  const [tunedNodes, setTunedNodes] = useState<OnpremNode[]>([]);
  const [isJsonOpen, setIsJsonOpen] = useState(false);
  const [tunedNetwork, setTunedNetwork] = useState<any>(null);
  const [newCidr, setNewCidr] = useState('');
  const [activeStep, setActiveStep] = useState<number>(() =>
    editedCandidate && recommendationCandidates.length > 0 ? 2 : 1
  );
  const [isModelLoaded, setIsModelLoaded] = useState(false);
  const [recommendMode, setRecommendMode] = useState<'new' | 'saved'>('new');
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);

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

  // Auto-fetch regions for all target CSPs in targetPairs
  useEffect(() => {
    (targetPairs || []).forEach(async (pair) => {
      if (pair.csp && !pairRegionsMap[pair.csp]) {
        try {
          const regions = await tumblebugApi.getRegions(pair.csp);
          setPairRegionsMap(prev => ({ ...prev, [pair.csp]: regions }));
        } catch (err) {
          console.warn(`Failed to fetch regions for ${pair.csp}:`, err);
        }
      }
    });
  }, [targetPairs]);

  const getCspDefaultRegion = async (csp: string): Promise<string> => {
    const key = (csp || '').toLowerCase();
    let regions = pairRegionsMap[key] || pairRegionsMap[csp];
    if (!regions || regions.length === 0) {
      try {
        regions = await tumblebugApi.getRegions(key);
        if (regions && regions.length > 0) {
          setPairRegionsMap(prev => ({ ...prev, [key]: regions, [csp]: regions }));
        }
      } catch (err) {
        console.warn(`Failed to fetch regions for ${key}:`, err);
      }
    }
    return pickPreferredKoreaRegion(key, regions);
  };

  const handlePairCspChange = async (index: number, newCsp: string) => {
    const defaultRegion = await getCspDefaultRegion(newCsp);
    updateTargetPair(index, newCsp, defaultRegion);
  };

  const handleAddNextPair = async () => {
    if ((targetPairs || []).length >= 10) return;
    const availableCsps = tumblebugProviders && tumblebugProviders.length > 0
      ? tumblebugProviders
      : ['aws', 'azure', 'gcp', 'alibaba', 'ncp'];
    const usedCsps = (targetPairs || []).map(p => p.csp.toLowerCase());
    const nextCsp = availableCsps.find(c => !usedCsps.includes(c.toLowerCase())) || availableCsps[0];
    const defaultRegion = await getCspDefaultRegion(nextCsp);
    addTargetPair(nextCsp, defaultRegion);
  };

  const hasDuplicatePairs = (() => {
    const seen = new Set<string>();
    for (const pair of targetPairs || []) {
      const key = `${pair.csp.toLowerCase()}/${pair.region.toLowerCase()}`;
      if (seen.has(key)) return true;
      seen.add(key);
    }
    return false;
  })();

  const handleSwitchMode = (mode: 'multi' | 'single') => {
    if (mode === recommendationMode) return;
    setRecommendationMode(mode);
    if (mode === 'single') {
      if ((targetPairs || []).length > 1) {
        setCachedMultiPairs(targetPairs);
      }
      const firstPair = (targetPairs && targetPairs.length > 0) ? targetPairs[0] : { csp: 'aws', region: 'ap-northeast-2' };
      setTargetPairs([firstPair]);
    } else {
      // Cross-CSP Exploration requires at least 2 pairs
      if (cachedMultiPairs && cachedMultiPairs.length >= 2) {
        setTargetPairs(cachedMultiPairs);
      } else if ((targetPairs || []).length < 2) {
        const first = (targetPairs && targetPairs.length > 0) ? targetPairs[0] : { csp: 'aws', region: 'ap-northeast-2' };
        const secondCsp = first.csp.toLowerCase() === 'aws' ? 'gcp' : 'aws';
        const secondRegion = secondCsp === 'gcp' ? 'asia-northeast3' : 'ap-northeast-2';
        setTargetPairs([first, { csp: secondCsp, region: secondRegion }]);
      }
    }
  };

  const handleApplyGlobalBig3 = () => {
    const global3 = [
      { csp: 'aws', region: 'ap-northeast-2' },
      { csp: 'gcp', region: 'asia-northeast3' },
      { csp: 'azure', region: 'koreacentral' }
    ];
    setTargetPairs(global3);
    global3.forEach(p => {
      tumblebugApi.getRegions(p.csp).then(regs => {
        if (regs && regs.length > 0) {
          setPairRegionsMap(prev => ({ ...prev, [p.csp]: regs }));
        }
      }).catch(() => {});
    });
  };

  const handleApplyDomesticBig2 = async () => {
    const ncpRegion = await getCspDefaultRegion('ncp');
    const nhnRegion = await getCspDefaultRegion('nhn');
    setTargetPairs([
      { csp: 'ncp', region: ncpRegion },
      { csp: 'nhn', region: nhnRegion }
    ]);
  };

  const handleResetToDefaultMultiPairs = () => {
    const default2 = [
      { csp: 'aws', region: 'ap-northeast-2' },
      { csp: 'gcp', region: 'asia-northeast3' }
    ];
    setTargetPairs(default2);
  };

  const handleDeepDiveWithCsp = async (csp: string, region: string) => {
    setCachedMultiCandidates(recommendationCandidates);
    setCachedMultiPairs(targetPairs);
    setRecommendationMode('single');
    setTargetPairs([{ csp, region }]);
    setRecommendationLimit(5);
    const filteredNodes = tunedNodes.filter(n => !excludedNodeIds.includes(n.machineId));
    const sourceNlbs = selectedSourceModel?.onpremiseInfraModel?.nlbs;
    const sourceInfra: OnpremInfra = {
      nodes: filteredNodes.length > 0 ? filteredNodes : (selectedSourceModel?.onpremiseInfraModel?.nodes || []),
      network: tunedNetwork || selectedSourceModel?.onpremiseInfraModel?.network || {
        ipv4Networks: {},
        ipv6Networks: {}
      },
      nlbs: sourceNlbs && sourceNlbs.length > 0 ? sourceNlbs : undefined
    };
    useMigrationStore.setState({ isRecommending: true, recommendationCandidates: [], editedCandidate: null });
    try {
      const candidates = await beetleApi.getRecommendations(sourceInfra, csp, region, 5, minMatchRate);
      useMigrationStore.setState({
        recommendationCandidates: candidates,
        selectedCandidateIndex: 0,
        editedCandidate: candidates.length > 0 ? candidates[0] : null,
        isRecommending: false
      });
      setHasUnsavedChanges(candidates.length > 0);
      setActiveStep(2);
    } catch (err) {
      console.error('Failed to get deep-dive recommendations:', err);
      useMigrationStore.setState({ isRecommending: false });
      setToastTone('error');
      setToastMsg('Failed to get deep-dive recommendations for ' + csp.toUpperCase());
    }
  };

  const handleRestoreMultiCandidates = () => {
    if (cachedMultiCandidates && cachedMultiCandidates.length > 0) {
      useMigrationStore.setState({
        recommendationCandidates: cachedMultiCandidates,
        selectedCandidateIndex: 0,
        editedCandidate: cachedMultiCandidates[0],
        targetPairs: cachedMultiPairs && cachedMultiPairs.length > 0 ? cachedMultiPairs : targetPairs,
        recommendationMode: 'multi',
        cachedMultiCandidates: null,
        cachedMultiPairs: null,
      });
      setHasUnsavedChanges(true);
      setActiveStep(2);
    }
  };

  useEffect(() => {
    if (selectedSourceModel?.id) {
      setPendingSourceModelId(selectedSourceModel.id);
    }
  }, [selectedSourceModel?.id]);

  useEffect(() => {
    if (pendingSourceModelId && !savedSourceModels.some(m => m.id === pendingSourceModelId)) {
      setPendingSourceModelId(selectedSourceModel?.id || savedSourceModels[0]?.id || '');
    }
    if (loadedSourceModelId && !savedSourceModels.some(m => m.id === loadedSourceModelId)) {
      setLoadedSourceModelId('');
      setTunedNodes([]);
      setTunedNetwork(null);
      if (recommendMode === 'new') {
        setIsModelLoaded(false);
      }
    }
  }, [savedSourceModels, pendingSourceModelId, loadedSourceModelId, selectedSourceModel?.id, recommendMode]);

  const userPlanningModels = savedSourceModels.filter(m => !isBuiltInSampleModel(m) && m.stage === 'planning');
  const userSourceModels = savedSourceModels.filter(m => !isBuiltInSampleModel(m) && m.stage !== 'planning');
  const sampleSourceModels = savedSourceModels.filter(m => isBuiltInSampleModel(m));
  const pendingSourceModel = savedSourceModels.find(m => m.id === pendingSourceModelId) || null;
  const isSourceAlreadyLoaded =
    recommendMode === 'new' && Boolean(loadedSourceModelId) && Boolean(pendingSourceModel) && loadedSourceModelId === pendingSourceModel?.id;

  const validSavedCloudModels = savedCloudModels.filter(m => m && m.name !== 'Unnamed Design');
  const userTargetModels = validSavedCloudModels.filter(m => !isBuiltInSampleModel(m));
  const sampleTargetModels = validSavedCloudModels.filter(m => isBuiltInSampleModel(m));

  useEffect(() => {
    if (selectedCloudModel?.id && !String(selectedCloudModel.id).startsWith('candidate-')) {
      setPendingCloudModelId(selectedCloudModel.id);
    } else if (
      (!pendingCloudModelId || isBuiltInSampleModel(pendingCloudModelId)) &&
      !loadedCloudModelId &&
      userTargetModels.length > 0
    ) {
      setPendingCloudModelId(userTargetModels[userTargetModels.length - 1].id);
    } else if (!pendingCloudModelId && validSavedCloudModels.length > 0) {
      setPendingCloudModelId(validSavedCloudModels[0].id);
    }
  }, [selectedCloudModel?.id, userTargetModels.length, validSavedCloudModels.length]);

  const pendingCloudModel = validSavedCloudModels.find(m => m.id === pendingCloudModelId) || null;
  const isTargetAlreadyLoaded =
    recommendMode === 'saved' &&
    Boolean(loadedCloudModelId) &&
    Boolean(pendingCloudModel) &&
    loadedCloudModelId === pendingCloudModel?.id &&
    !hasUnsavedChanges;

  const handleLoadModel = () => {
    if (!pendingSourceModel || !pendingSourceModel.onpremiseInfraModel) return;
    selectSourceModel(pendingSourceModel);
    setLoadedSourceModelId(pendingSourceModel.id);
    setTunedNodes(JSON.parse(JSON.stringify(pendingSourceModel.onpremiseInfraModel.nodes || [])));
    setTunedNetwork(JSON.parse(JSON.stringify(pendingSourceModel.onpremiseInfraModel.network || { ipv4Networks: {}, ipv6Networks: {} })));
    setExcludedNodeIds([]);
    setIsModelLoaded(true);
    if (pendingSourceModel.onpremiseInfraModel.nodes && pendingSourceModel.onpremiseInfraModel.nodes.length > 0) {
      setActiveTunedNodeId(pendingSourceModel.onpremiseInfraModel.nodes[0]?.machineId || '');
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

  const handleRecommend = async () => {
    selectCloudModel(null);
    setLoadedCloudModelId('');
    const filteredNodes = tunedNodes.filter(n => !excludedNodeIds.includes(n.machineId));
    if (filteredNodes.length === 0) return;
    const sourceNlbs = selectedSourceModel?.onpremiseInfraModel?.nlbs;
    const sourceInfra: OnpremInfra = {
      nodes: filteredNodes,
      network: tunedNetwork || selectedSourceModel?.onpremiseInfraModel?.network || {
        ipv4Networks: {},
        ipv6Networks: {}
      },
      nlbs: sourceNlbs && sourceNlbs.length > 0 ? sourceNlbs : undefined
    };
    await triggerRecommendation(sourceInfra);
    const latestCandidates = useMigrationStore.getState().recommendationCandidates;
    setHasUnsavedChanges(latestCandidates.length > 0);
    setActiveStep(2);
  };

  const handleLoadDesign = () => {
    if (!pendingCloudModel) return;
    selectCloudModel(pendingCloudModel);
    setLoadedCloudModelId(pendingCloudModel.id);
    const linkedSource = useMigrationStore.getState().selectedSourceModel;
    if (linkedSource?.onpremiseInfraModel) {
      setLoadedSourceModelId(linkedSource.id);
      setPendingSourceModelId(linkedSource.id);
      setTunedNodes(JSON.parse(JSON.stringify(linkedSource.onpremiseInfraModel.nodes || [])));
      setTunedNetwork(JSON.parse(JSON.stringify(linkedSource.onpremiseInfraModel.network || { ipv4Networks: {}, ipv6Networks: {} })));
      if (linkedSource.onpremiseInfraModel.nodes && linkedSource.onpremiseInfraModel.nodes.length > 0) {
        setActiveTunedNodeId(linkedSource.onpremiseInfraModel.nodes[0]?.machineId || '');
      }
    }
    let rawInfra = pendingCloudModel.cloudInfraModel || (pendingCloudModel as any).cloud_infra_model || (pendingCloudModel as any).recommendedInfra;
    if (typeof rawInfra === 'string') {
      try { rawInfra = JSON.parse(rawInfra); } catch {}
    }
    const infra = rawInfra ? JSON.parse(JSON.stringify(rawInfra)) : null;
    if (infra) {
      storeUpdateEditedCandidate(infra);
      useMigrationStore.setState({
        recommendationCandidates: [infra],
        selectedCandidateIndex: 0
      });
      if (infra.targetCloud?.csp || (infra as any).csp) {
        setDesiredCsp(infra.targetCloud?.csp || (infra as any).csp);
      }
      if (infra.targetCloud?.region || (infra as any).region) {
        setDesiredRegion(infra.targetCloud?.region || (infra as any).region);
      }
    }
    setIsModelLoaded(true);
    setHasUnsavedChanges(false);
    setActiveStep(2);
  };

  const handleDeleteDesign = () => {
    if (!pendingCloudModel || isBuiltInSampleModel(pendingCloudModel)) return;
    setShowDeleteConfirm(true);
  };

  useEffect(() => {
    if (
      savedCloudModels.some(
        m =>
          m.name === 'Unnamed Design' ||
          m.id.startsWith('cloud-model-') ||
          m.id.startsWith('art-tgt-') ||
          !m.cloudInfraModel
      )
    ) {
      fetchSavedCloudModels();
    }
  }, [savedCloudModels, fetchSavedCloudModels]);

  const handleClearAllCustomDesigns = () => {
    const customCount = savedCloudModels.filter(m => !isBuiltInSampleModel(m)).length;
    if (customCount === 0) return;
    setShowClearAllConfirm(true);
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

  // Lookup SpecInfo from targetSpecList by specId (try id, name, cspSpecName case-insensitively)
  const getSpecInfo = (candidate: typeof editedCandidate, specId: string) => {
    if (!candidate?.targetSpecList || !specId) return null;
    const lowerId = specId.toLowerCase();
    const extractedLower = extractInstanceType(specId).toLowerCase();
    return candidate.targetSpecList.find((s: any) =>
      s.id?.toLowerCase() === lowerId ||
      s.name?.toLowerCase() === lowerId ||
      s.cspSpecName?.toLowerCase() === extractedLower
    ) ?? (candidate.targetSpecList.length === 1 ? candidate.targetSpecList[0] : null);
  };

  // Helper to find source node by machineId/hostname strictly within active model
  const findSourceNode = (sourceId: string): OnpremNode | null => {
    if (!sourceId) return null;
    const cleanId = sourceId.trim().toLowerCase();

    // 1. Check selectedSourceModel
    const currentNodes = (selectedSourceModel?.onpremiseInfraModel?.nodes ||
                         (selectedSourceModel as any)?.sourceInfra?.nodes || []) as OnpremNode[];
    const match1 = currentNodes.find((n: any) =>
      n.machineId?.toLowerCase() === cleanId ||
      n.hostname?.toLowerCase() === cleanId
    );
    if (match1) return match1;

    // 2. Fallback to sample GPU data ONLY if current model is explicitly a GPU workload
    const isGpuWorkload = Boolean(
      selectedSourceModel?.id?.toLowerCase().includes('gpu') ||
      selectedSourceModel?.name?.toLowerCase().includes('gpu')
    );
    if (isGpuWorkload) {
      const sampleNodes = ((sampleGpuData as any)?.nodes || []) as OnpremNode[];
      const matchSample = sampleNodes.find((n: any) =>
        n.machineId?.toLowerCase() === cleanId ||
        n.hostname?.toLowerCase() === cleanId
      );
      if (matchSample) return matchSample;
    }

    return null;
  };

  const [isTargetJsonOpen, setIsTargetJsonOpen] = useState(false);
  const [isTargetJsonCopied, setIsTargetJsonCopied] = useState(false);

  // Helper to determine if target node / spec has GPU hardware
  const getGpuInfo = (specInfo: any, specId?: string, ng?: any): { hasGpu: boolean; label: string } | null => {
    const matchByInstanceName = (rawName: string): { hasGpu: boolean; label: string } | null => {
      const s = rawName.toLowerCase();
      if (s.includes('g4dn.') || (s.includes('standard_nc') && s.includes('t4'))) return { hasGpu: true, label: 'Tesla T4' };
      if (s.includes('g5.') || s.includes('g5g.') || (s.includes('standard_nv') && s.includes('a10'))) return { hasGpu: true, label: 'A10G' };
      if (s.includes('g6.') || s.includes('g2-')) return { hasGpu: true, label: 'L4' };
      if (s.includes('p3.') || s.includes('standard_nc6s_v3') || s.includes('standard_nc12s_v3') || s.includes('standard_nc24s_v3')) return { hasGpu: true, label: 'V100' };
      if (s.includes('p4d') || s.includes('a2-') || (s.includes('standard_nc') && s.includes('a100')) || (s.includes('standard_nd') && s.includes('a100'))) return { hasGpu: true, label: 'A100' };
      if (s.includes('p5.') || s.includes('a3-') || (s.includes('standard_nc') && s.includes('h100')) || (s.includes('standard_nd') && s.includes('h100'))) return { hasGpu: true, label: 'H100' };
      if (s.includes('standard_nc') || s.includes('standard_nd') || s.includes('standard_nv')) return { hasGpu: true, label: 'GPU' };
      return null;
    };

    // 1. Authoritative: CB-Tumblebug SpecInfo metadata
    if (specInfo) {
      const count = Number(specInfo.acceleratorCount ?? specInfo.AcceleratorCount ?? specInfo.gpuCount ?? specInfo.GpuCount ?? 0);
      const model = String(specInfo.acceleratorModel ?? specInfo.AcceleratorModel ?? specInfo.gpuModel ?? specInfo.GpuModel ?? '').trim();
      const memGb = Number(specInfo.acceleratorMemoryGB ?? specInfo.AcceleratorMemoryGB ?? specInfo.gpuMemoryGiB ?? 0);
      if (count > 0 || model) {
        const prefix = count > 1 ? `${count}x ` : '';
        const memSuffix = memGb > 0 ? ` (${memGb}GB)` : '';
        const label = model ? `${prefix}${model}${memSuffix}` : `${count || 1}x GPU${memSuffix}`;
        return { hasGpu: true, label };
      }
      if (Array.isArray(specInfo.details)) {
        const gpuDetail = specInfo.details.find((d: any) =>
          d.key === 'GpuInfo' || d.key === 'AcceleratorInfo' || d.key === 'InferenceAcceleratorInfo'
        );
        if (gpuDetail) return { hasGpu: true, label: String(gpuDetail.value || 'GPU') };
      }
      const bySpecName = matchByInstanceName(String(specInfo.cspSpecName || specInfo.name || specId || ''));
      if (bySpecName) return bySpecName;
      return null;
    }

    // 2. CSP Instance Spec naming pattern when SpecInfo object is unavailable
    if (specId) {
      return matchByInstanceName(specId);
    }

    // 3. Fallback when neither specInfo nor specId is populated yet
    if (ng && (ng.name || '').toLowerCase().includes('gpu')) {
      return { hasGpu: true, label: 'GPU' };
    }

    return null;
  };

  // Candidate-level GPU accelerators summary
  const candidateGpuSummary = React.useMemo(() => {
    if (!editedCandidate?.targetInfra?.nodeGroups) {
      return { hasGpu: false, gpuNgCount: 0, gpuNodeCount: 0 };
    }

    const gpuNgs = editedCandidate.targetInfra.nodeGroups.filter(ng => {
      const specInfo = getSpecInfo(editedCandidate, ng.specId);
      return !!getGpuInfo(specInfo, ng.specId, ng);
    });

    const hasGpu = gpuNgs.length > 0;
    const gpuNgCount = gpuNgs.length;
    const gpuNodeCount = gpuNgs.reduce((acc, ng) => acc + (ng.nodeGroupSize || 0), 0);
    return { hasGpu, gpuNgCount, gpuNodeCount };
  }, [editedCandidate]);

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

  // Human-readable memory: show GiB or MiB with field-standard rounding (max 2 decimal places for fractional GiB)
  const formatMemory = (gib: number) => {
    if (!gib && gib !== 0) return '-';
    if (gib < 1) {
      return `${Math.round(gib * 1024)} MiB`;
    }
    const rounded = Number.isInteger(gib) ? gib : Math.round(gib * 100) / 100;
    return `${rounded} GiB`;
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
    const existingSubnets = updatedCandidate.targetVNet.subnetInfoList || [];
    const existingNames = new Set(existingSubnets.map((s: any) => s.name));
    const existingCidrs = new Set(existingSubnets.map((s: any) => s.ipv4_CIDR));

    // Match the prefix convention of the first recommended subnet (e.g. "subnet-01" -> "subnet-", "mig-subnet-01" -> "mig-subnet-")
    const firstSubnetName = existingSubnets[0]?.name || 'subnet-01';
    const prefixMatch = firstSubnetName.match(/^(.*?)(\d+)$/);
    const namePrefix = prefixMatch ? prefixMatch[1] : 'subnet-';
    const padWidth = prefixMatch ? Math.max(prefixMatch[2].length, 2) : 2;

    let n = existingSubnets.length + 1;
    let nextName = `${namePrefix}${String(n).padStart(padWidth, '0')}`;
    while (existingNames.has(nextName)) {
      n += 1;
      nextName = `${namePrefix}${String(n).padStart(padWidth, '0')}`;
    }

    const vnetCidrBase = (updatedCandidate.targetVNet.cidrBlock || '10.0.0.0/16').split('.').slice(0, 2).join('.');
    let cidrOctet = n;
    let nextCidr = `${vnetCidrBase}.${cidrOctet}.0/24`;
    while (existingCidrs.has(nextCidr) && cidrOctet < 255) {
      cidrOctet += 1;
      nextCidr = `${vnetCidrBase}.${cidrOctet}.0/24`;
    }

    updatedCandidate.targetVNet.subnetInfoList.push({
      name: nextName,
      ipv4_CIDR: nextCidr,
      description: 'a recommended subnet for migration',
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

  // Naming Alignment & Validation (CM-Beetle v0.5.12+)
  const [namingFocusOldName, setNamingFocusOldName] = useState<string>('');
  const [namingValidation, setNamingValidation] = useState<{
    loading?: boolean;
    valid?: boolean;
    message?: string;
  }>({});

  const handleAlignResourceNameOnBlur = async (
    resourceType: 'vNet' | 'subnet' | 'sshKey' | 'securityGroup' | 'infra',
    oldName: string,
    newName: string
  ) => {
    if (!editedCandidate || !oldName || !newName || oldName === newName) return;
    setNamingValidation({ loading: true, message: 'Aligning resource references...' });
    try {
      const alignRes = await beetleApi.alignResourceName(resourceType, oldName, newName, editedCandidate);
      const nextModel = alignRes.success && alignRes.data ? alignRes.data : editedCandidate;
      if (alignRes.success && alignRes.data) {
        updateEditedCandidate(alignRes.data);
      }
      const valRes = await beetleApi.validateNaming(nextModel);
      setNamingValidation({
        loading: false,
        valid: valRes.valid,
        message: valRes.message || valRes.error || (valRes.valid ? 'Resource naming references aligned' : 'Naming reference mismatch detected')
      });
    } catch (err: any) {
      setNamingValidation({
        loading: false,
        valid: false,
        message: err.message || 'Naming alignment request failed'
      });
    }
  };

  const handleValidateNaming = async () => {
    if (!editedCandidate) return;
    setNamingValidation({ loading: true, message: 'Validating resource naming references...' });
    try {
      const valRes = await beetleApi.validateNaming(editedCandidate);
      setNamingValidation({
        loading: false,
        valid: valRes.valid,
        message: valRes.message || valRes.error || (valRes.valid ? 'All resource naming references are valid' : 'Naming reference issue detected')
      });
    } catch (err: any) {
      setNamingValidation({
        loading: false,
        valid: false,
        message: err.message || 'Naming validation failed'
      });
    }
  };

  // Post-Deployment Bootstrap Commands (targetInfra.postCommands & postCommandAsync)
  const handleTogglePostCommandAsync = (checked: boolean) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    if (!updatedCandidate.targetInfra) return;
    updatedCandidate.targetInfra.postCommandAsync = checked;
    updateEditedCandidate(updatedCandidate);
  };

  const handleAddPostCommand = () => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    if (!updatedCandidate.targetInfra) return;
    const existing = updatedCandidate.targetInfra.postCommands || [];
    updatedCandidate.targetInfra.postCommands = [
      ...existing,
      {
        commands: ['sudo apt-get update -y'],
        nodeGroupName: '',
        nodeName: ''
      }
    ];
    updateEditedCandidate(updatedCandidate);
  };

  const handleRemovePostCommand = (cmdIdx: number) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    if (!updatedCandidate.targetInfra?.postCommands) return;
    updatedCandidate.targetInfra.postCommands = updatedCandidate.targetInfra.postCommands.filter(
      (_: any, idx: number) => idx !== cmdIdx
    );
    updateEditedCandidate(updatedCandidate);
  };

  const handleUpdatePostCommand = (
    cmdIdx: number,
    field: 'nodeGroupName' | 'nodeName' | 'commandsText',
    value: string
  ) => {
    if (!editedCandidate) return;
    const updatedCandidate = JSON.parse(JSON.stringify(editedCandidate));
    if (!updatedCandidate.targetInfra?.postCommands?.[cmdIdx]) return;
    if (field === 'commandsText') {
      updatedCandidate.targetInfra.postCommands[cmdIdx].commands = value
        .split('\n')
        .map((line: string) => line.trim())
        .filter(Boolean);
    } else {
      updatedCandidate.targetInfra.postCommands[cmdIdx][field] = value;
    }
    updateEditedCandidate(updatedCandidate);
  };

  useEffect(() => {
    if (targetActiveTab === 'nlb' && (!editedCandidate?.targetNlbList || editedCandidate.targetNlbList.length === 0)) {
      setTargetActiveTab('compute');
    }
  }, [editedCandidate, targetActiveTab]);

  const handleSaveArtifact = async (result: { name: string; description: string; version: string; overwriteId: string | null }) => {
    if (!editedCandidate) return;
    const validInfraName = resolveValidTargetInfraName(
      editedCandidate.targetInfra?.name,
      result.name,
      editedCandidate.targetCloud?.csp
    );
    const targetModelToSave: RecommendedInfra = {
      ...editedCandidate,
      targetNlbList:
        editedCandidate.targetNlbList && editedCandidate.targetNlbList.length > 0
          ? editedCandidate.targetNlbList
          : undefined,
      targetInfra: {
        ...editedCandidate.targetInfra,
        name: validInfraName
      }
    };
    storeUpdateEditedCandidate(targetModelToSave);

    const savedModel = result.overwriteId
      ? await updateCloudModel(result.overwriteId, result.name, result.description, result.version, targetModelToSave)
      : await saveCloudModel(result.name, result.description, result.version, targetModelToSave);

    if (savedModel) {
      selectCloudModel(savedModel);
      setPendingCloudModelId(savedModel.id);
      setLoadedCloudModelId(savedModel.id);
    }
    setHasUnsavedChanges(false);
    setToastTone('success');
    setToastMsg(
      result.overwriteId
        ? `Target cloud model "${savedModel.name}" (v${savedModel.version || result.version}) updated successfully.`
        : `Target cloud model "${savedModel.name}" (v${savedModel.version || result.version}) saved successfully.`
    );
  };

  const { isRecommending } = useMigrationStore();

  return (
    <div className="space-y-6 mx-auto pb-24">

      {/* Top Banner Description Box */}
      <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
        <div className="flex items-center gap-2 shrink-0">
          <Compass className="w-5 h-5 text-emerald-500" />
          <h2 className="text-base font-extrabold text-text-main tracking-tight">
            Target Cloud Optimizer
          </h2>
        </div>
        <span className="text-sm text-text-muted">
          Generate optimized target cloud recommendations, compare multi-CSP specs & cost estimates, and customize target cloud models.
        </span>
      </div>

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
            type="button"
            onClick={() => { setRecommendMode('new'); setIsModelLoaded(Boolean(selectedSourceModel && tunedNodes.length > 0)); }}
            className={`flex-1 pb-3 text-sm font-extrabold text-center border-b-2 transition-all cursor-pointer ${
              recommendMode === 'new'
                ? 'border-emerald-500 text-emerald-500'
                : 'border-transparent text-text-muted hover:text-text-main'
            }`}
          >
            Generate New Recommendation
          </button>
          <button
            type="button"
            onClick={() => { setRecommendMode('saved'); setIsModelLoaded(Boolean(selectedCloudModel && !String(selectedCloudModel.id).startsWith('candidate-'))); }}
            className={`flex-1 pb-3 text-sm font-extrabold text-center border-b-2 transition-all cursor-pointer ${
              recommendMode === 'saved'
                ? 'border-emerald-500 text-emerald-500'
                : 'border-transparent text-text-muted hover:text-text-main'
            }`}
          >
            Load Saved Target Cloud Model
          </button>
        </div>

        {recommendMode === 'new' ? (
          <>
            {/* Source Model selection */}
            <div className="mb-6 pb-6 border-b border-border-main/20">
              <label className="block text-sm font-normal text-text-muted mb-2">Source or Planning Infrastructure Model</label>
              <div className="flex flex-col gap-3 max-w-2xl">
                <div className="flex flex-wrap items-center gap-2.5">
                  <select
                    value={pendingSourceModelId}
                    onChange={(e) => setPendingSourceModelId(e.target.value)}
                    className="flex-1 min-w-[260px] bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-3 text-sm font-extrabold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                  >
                    <option value="">-- Choose Infrastructure Model --</option>
                    {userPlanningModels.length > 0 && (
                      <optgroup label="Planning Models">
                        {userPlanningModels.map((m) => (
                          <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                        ))}
                      </optgroup>
                    )}
                    {userSourceModels.length > 0 && (
                      <optgroup label="Source Models">
                        {userSourceModels.map((m) => (
                          <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                        ))}
                      </optgroup>
                    )}
                    <optgroup label="Sample Models">
                      {sampleSourceModels.map((m) => (
                        <option key={m.id} value={m.id}>{formatSourceOptionText(m)}</option>
                      ))}
                    </optgroup>
                  </select>
                  {pendingSourceModel && (() => {
                    const badge = getStageBadgeMeta(resolveModelStageKind(pendingSourceModel, 'source'));
                    return (
                      <span className={`px-3 py-1.5 rounded-xl text-xs font-extrabold border shrink-0 ${badge.className}`}>
                        {badge.label}
                      </span>
                    );
                  })()}
                </div>

                {pendingSourceModel && (
                  <div className="text-xs bg-bg-input/50 border border-border-main/50 rounded-xl px-3.5 py-2.5 flex flex-wrap items-center gap-x-4 gap-y-1">
                    <span>
                      <span className="font-normal text-text-muted">Nodes: </span>
                      <span className="font-extrabold text-text-main">{pendingSourceModel.onpremiseInfraModel?.nodes?.length || 0}</span>
                    </span>
                    {pendingSourceModel.description && (
                      <span className="truncate max-w-md">
                        <span className="font-normal text-text-muted">Info: </span>
                        <span className="font-extrabold text-text-main">{pendingSourceModel.description}</span>
                      </span>
                    )}
                  </div>
                )}

                <div className="flex flex-wrap items-center gap-3">
                  <button
                    type="button"
                    onClick={handleLoadModel}
                    disabled={!pendingSourceModel}
                    className={`px-5 py-2.5 rounded-xl text-sm font-extrabold flex items-center gap-2 transition cursor-pointer ${
                      !pendingSourceModel
                        ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                        : isSourceAlreadyLoaded
                          ? 'bg-emerald-500/20 border border-emerald-500/40 text-emerald-600 dark:text-emerald-400'
                          : 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 shadow-md shadow-emerald-500/25'
                    }`}
                  >
                    {isSourceAlreadyLoaded ? (
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
                  {isModelLoaded && tunedNodes.length > 0 && selectedSourceModel && (
                    <span className="text-xs font-extrabold text-emerald-600 dark:text-emerald-400">
                      {tunedNodes.length} Node(s) loaded from &quot;{cleanModelDisplayName(selectedSourceModel.name)}&quot;
                    </span>
                  )}
                </div>
              </div>
            </div>

            {/* Target Cloud Pair(s) Selection Container */}
            <div className="max-w-4xl space-y-4">
              {/* Mode Selector Tab (Standalone row, fixed height to prevent UI shift) */}
              <div className="flex items-center gap-1.5 p-1 bg-bg-panel border border-border-main rounded-2xl w-fit">
                <button
                  type="button"
                  onClick={() => handleSwitchMode('multi')}
                  className={`px-4 py-2.5 rounded-xl text-xs font-extrabold flex items-center gap-2 transition cursor-pointer ${
                    recommendationMode === 'multi'
                      ? 'bg-emerald-500 text-slate-950 shadow-sm'
                      : 'text-text-muted hover:text-text-main'
                  }`}
                >
                  <Globe className="w-4 h-4" />
                  <span>Cross-CSP Exploration (Multi-Cloud)</span>
                  {recommendationMode === 'multi' && (targetPairs || []).length > 1 && (
                    <span className="px-1.5 py-0.5 rounded-full bg-slate-950/20 text-slate-950 text-xs font-mono">
                      {(targetPairs || []).length}
                    </span>
                  )}
                </button>
                <button
                  type="button"
                  onClick={() => handleSwitchMode('single')}
                  className={`px-4 py-2.5 rounded-xl text-xs font-extrabold flex items-center gap-2 transition cursor-pointer ${
                    recommendationMode === 'single'
                      ? 'bg-emerald-500 text-slate-950 shadow-sm'
                      : 'text-text-muted hover:text-text-main'
                  }`}
                >
                  <Target className="w-4 h-4" />
                  <span>Single-CSP Deep Dive (Architecture Options)</span>
                  {recommendationMode === 'single' && (
                    <span className="px-1.5 py-0.5 rounded-full bg-slate-950/20 text-slate-950 text-xs font-mono">
                      {recommendationLimit}
                    </span>
                  )}
                </button>
              </div>

              {/* Header Label and Add Pair Button */}
              <div className="flex flex-wrap items-center justify-between gap-2 pt-1">
                <div>
                  <label className="block text-sm font-normal text-text-muted">
                    {recommendationMode === 'multi' ? 'Target Cloud Pair(s)' : 'Target Cloud Provider & Region'}
                  </label>
                  <span className="text-xs text-text-muted font-normal">
                    {recommendationMode === 'multi'
                      ? 'Select 2~10 distinct CSP and region pairs to compare top architectures across clouds'
                      : 'Select a target cloud provider and region to explore multiple alternative candidate specifications'}
                  </span>
                </div>
                {recommendationMode === 'multi' && (
                  <button
                    onClick={handleAddNextPair}
                    disabled={(targetPairs || []).length >= 10}
                    className="px-3.5 py-1.5 bg-emerald-500/10 border border-emerald-500/30 hover:bg-emerald-500/20 text-emerald-600 dark:text-emerald-400 disabled:opacity-40 rounded-xl text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer"
                  >
                    <Plus className="w-3.5 h-3.5" />
                    <span>Add Target Cloud Pair ({(targetPairs || []).length}/10)</span>
                  </button>
                )}
              </div>

              {/* Mode-specific Toolbar Strip (Presets for Multi-Cloud, Recommendation Count for Single-Cloud, plus Min Match Rate) */}
              {recommendationMode === 'multi' ? (
                <div className="flex flex-wrap items-center justify-between gap-2.5 px-4 py-2.5 bg-bg-panel/50 border border-border-main/50 rounded-xl">
                  <div className="flex flex-wrap items-center gap-2">
                    <span className="text-xs font-normal text-text-muted flex items-center gap-1 mr-1">
                      <Zap className="w-3.5 h-3.5 text-amber-500 shrink-0" />
                      <span>Quick Presets:</span>
                    </span>
                    <button
                      type="button"
                      onClick={handleApplyGlobalBig3}
                      className="px-2.5 py-1 bg-bg-panel hover:bg-emerald-500/10 border border-border-main hover:border-emerald-500/30 text-text-main hover:text-emerald-500 rounded-lg text-xs font-extrabold transition cursor-pointer flex items-center gap-1.5 shadow-sm"
                      title="AWS (ap-northeast-2), GCP (asia-northeast3), Azure (koreacentral)"
                    >
                      <span>+ Global Big 3</span>
                      <span className="text-xs font-normal text-text-muted">(AWS, GCP, Azure)</span>
                    </button>
                    <button
                      type="button"
                      onClick={handleApplyDomesticBig2}
                      className="px-2.5 py-1 bg-bg-panel hover:bg-emerald-500/10 border border-border-main hover:border-emerald-500/30 text-text-main hover:text-emerald-500 rounded-lg text-xs font-extrabold transition cursor-pointer flex items-center gap-1.5 shadow-sm"
                      title="NCP (kr), NHN Cloud (kr1)"
                    >
                      <span>+ Domestic 2</span>
                      <span className="text-xs font-normal text-text-muted">(NCP, NHN)</span>
                    </button>
                  </div>
                  <div className="flex flex-wrap items-center gap-3">
                    <div className="flex items-center gap-1.5">
                      <span className="text-xs font-normal text-text-muted">Min Match Rate (%):</span>
                      <input
                        type="number"
                        min={0}
                        max={100}
                        value={minMatchRate}
                        onChange={(e) => {
                          const val = parseInt(e.target.value, 10);
                          setMinMatchRate(isNaN(val) ? 85 : Math.min(100, Math.max(0, val)));
                        }}
                        className="w-16 bg-bg-input border border-border-main text-text-main rounded-lg px-2 py-1 text-xs font-extrabold font-mono focus:outline-none focus:ring-1 focus:ring-emerald-500"
                        title="Minimum spec match rate threshold (0-100%, default 85%)"
                      />
                    </div>
                    {(targetPairs || []).length > 2 && (
                      <button
                        type="button"
                        onClick={handleResetToDefaultMultiPairs}
                        className="px-2.5 py-1 text-text-muted hover:text-red-400 rounded-lg text-xs font-normal transition cursor-pointer"
                      >
                        Reset to 2 Pairs
                      </button>
                    )}
                  </div>
                </div>
              ) : (
                <div className="flex flex-wrap items-center justify-between gap-3 px-4 py-2 bg-bg-panel/50 border border-border-main/50 rounded-xl">
                  <div className="flex items-center gap-3">
                    <span className="text-xs font-normal text-text-muted flex items-center gap-1.5 shrink-0">
                      <Sliders className="w-3.5 h-3.5 text-emerald-500 shrink-0" />
                      <span>Recommendation Count:</span>
                    </span>
                    <select
                      value={recommendationLimit}
                      onChange={(e) => setRecommendationLimit(Number(e.target.value))}
                      className="bg-bg-input border border-border-main text-text-main rounded-lg px-3 py-1.5 text-xs font-extrabold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                    >
                      {[1, 2, 3, 4, 5, 6, 7, 8, 9, 10].map((num) => (
                        <option key={num} value={num}>
                          {num} {num === 1 ? 'Candidate' : 'Candidates'} {num === 5 ? '(Default)' : ''}
                        </option>
                      ))}
                    </select>
                    <span className="text-xs text-text-muted font-normal hidden sm:inline">
                      Generate {recommendationLimit} architecture candidate options for comparison
                    </span>
                  </div>
                  <div className="flex items-center gap-1.5">
                    <span className="text-xs font-normal text-text-muted">Min Match Rate (%):</span>
                    <input
                      type="number"
                      min={0}
                      max={100}
                      value={minMatchRate}
                      onChange={(e) => {
                        const val = parseInt(e.target.value, 10);
                        setMinMatchRate(isNaN(val) ? 85 : Math.min(100, Math.max(0, val)));
                      }}
                      className="w-16 bg-bg-input border border-border-main text-text-main rounded-lg px-2 py-1 text-xs font-extrabold font-mono focus:outline-none focus:ring-1 focus:ring-emerald-500"
                      title="Minimum spec match rate threshold (0-100%, default 85%)"
                    />
                  </div>
                </div>
              )}

              {recommendationMode === 'multi' && (targetPairs || []).length < 2 && (
                <div className="p-3 bg-amber-500/10 border border-amber-500/30 text-amber-600 dark:text-amber-400 rounded-xl text-xs font-bold flex items-center gap-2">
                  <AlertTriangle className="w-4 h-4 shrink-0" />
                  <span>Cross-CSP Exploration requires at least 2 distinct target cloud pairs to compare architectures.</span>
                </div>
              )}

              {hasDuplicatePairs && (targetPairs || []).length > 1 && (
                <div className="p-3 bg-amber-500/10 border border-amber-500/30 text-amber-600 dark:text-amber-400 rounded-xl text-xs font-bold flex items-center gap-2">
                  <AlertTriangle className="w-4 h-4 shrink-0" />
                  <span>Duplicate target CSP & Region pair detected. Multi-CSP recommendation requires unique CSP/Region pairs.</span>
                </div>
              )}

              <div className="space-y-3">
                {(targetPairs || []).map((pair, idx) => {
                  const cspKey = (pair.csp || '').toLowerCase();
                  const availableRegions = pairRegionsMap[pair.csp] || pairRegionsMap[cspKey] || tumblebugRegions;
                  const prefs = (PREFERRED_KOREA_REGIONS_BY_CSP[cspKey] || []).map((p) => p.toLowerCase());
                  const sortedRegions = [...availableRegions].sort((a, b) => {
                    const aPref = prefs.indexOf(a.id.toLowerCase());
                    const bPref = prefs.indexOf(b.id.toLowerCase());
                    if (aPref !== -1 || bPref !== -1) {
                      if (aPref === -1) return 1;
                      if (bPref === -1) return -1;
                      return aPref - bPref;
                    }
                    const aKr = isKoreaRegion(a) ? 0 : 1;
                    const bKr = isKoreaRegion(b) ? 0 : 1;
                    if (aKr !== bKr) return aKr - bKr;
                    return a.id.localeCompare(b.id);
                  });
                  const selectedRegionValue =
                    sortedRegions.find((r) => r.id.toLowerCase() === (pair.region || '').toLowerCase())?.id ||
                    pair.region;

                  return (
                    <div key={idx} className="bg-bg-panel/50 border border-border-main/50 p-4 rounded-xl flex flex-col md:flex-row items-stretch md:items-center gap-4 transition hover:border-emerald-500/30">
                      <div className="flex items-center gap-2 shrink-0 md:w-28">
                        <span className="w-6 h-6 rounded-full bg-emerald-500/10 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 flex items-center justify-center text-xs font-extrabold">
                          {idx + 1}
                        </span>
                        <span className="text-sm font-normal text-text-muted">Target Pair</span>
                      </div>

                      <div className="grid grid-cols-1 sm:grid-cols-2 gap-3 flex-1">
                        <div>
                          <label className="block text-xs font-normal text-text-muted mb-1">Provider (CSP)</label>
                          <select
                            value={pair.csp}
                            onChange={(e) => handlePairCspChange(idx, e.target.value)}
                            className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-3 py-2.5 text-sm font-extrabold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
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
                          <label className="block text-xs font-normal text-text-muted mb-1">Region</label>
                          <select
                            value={selectedRegionValue}
                            onChange={(e) => updateTargetPair(idx, pair.csp, e.target.value)}
                            className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-3 py-2.5 text-sm font-extrabold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
                          >
                            {sortedRegions.map((r) => (
                              <option key={r.id} value={r.id}>
                                {r.name} ({r.id})
                              </option>
                            ))}
                          </select>
                        </div>
                      </div>

                      {recommendationMode === 'multi' && (targetPairs || []).length > 2 && (
                        <button
                          onClick={() => removeTargetPair(idx)}
                          className="p-2 text-text-muted hover:text-red-500 hover:bg-red-500/10 rounded-xl transition cursor-pointer self-end md:self-center"
                          title="Remove target pair"
                        >
                          <Trash2 className="w-4 h-4" />
                        </button>
                      )}
                    </div>
                  );
                })}
              </div>
            </div>

            <div className="mt-6 flex justify-start">
              <button
                onClick={handleRecommend}
                disabled={
                  isRecommending ||
                  tunedNodes.length === 0 ||
                  (recommendationMode === 'multi' && (targetPairs || []).length < 2) ||
                  (hasDuplicatePairs && (targetPairs || []).length > 1)
                }
                className="px-6 py-3.5 bg-gradient-to-r from-emerald-500 to-teal-600 hover:from-emerald-600 hover:to-teal-700 disabled:opacity-50 text-slate-950 font-extrabold rounded-xl text-sm tracking-wider transition shadow-lg shadow-emerald-500/10 flex items-center justify-center space-x-2 cursor-pointer"
              >
                {isRecommending ? (
                  <>
                    <RefreshCw className="w-4 h-4 animate-spin" />
                    <span>Recommending...</span>
                  </>
                ) : (
                  <>
                    <Sparkles className="w-4 h-4" />
                    <span>
                      {recommendationMode === 'multi'
                        ? (targetPairs || []).length < 2
                          ? 'Select at least 2 Target Pairs to Compare'
                          : `Recommend Cross-CSP Target Infrastructure (${(targetPairs || []).length} Pairs)`
                        : `Recommend Target Cloud Infrastructure (${recommendationLimit} Candidates)`}
                    </span>
                  </>
                )}
              </button>
            </div>
          </>
        ) : (
          <div className="max-w-2xl space-y-3">
            <label className="block text-sm font-normal text-text-muted">Target Cloud Infrastructure Model</label>
            <div className="flex flex-wrap items-center gap-2.5">
              <select
                value={pendingCloudModelId}
                onChange={(e) => setPendingCloudModelId(e.target.value)}
                className="flex-1 min-w-[260px] bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-3 text-sm font-extrabold focus:outline-none focus:ring-1 focus:ring-emerald-500 cursor-pointer"
              >
                <option value="">-- Choose Target Cloud Model --</option>
                {userTargetModels.length > 0 && (
                  <optgroup label="Target Models">
                    {userTargetModels.map((m) => (
                      <option key={m.id} value={m.id}>
                        {formatTargetOptionText(m, m.cloudInfraModel?.targetCloud?.csp)}
                      </option>
                    ))}
                  </optgroup>
                )}
                <optgroup label="Sample Models">
                  {sampleTargetModels.map((m) => (
                    <option key={m.id} value={m.id}>
                      {formatTargetOptionText(m, m.cloudInfraModel?.targetCloud?.csp)}
                    </option>
                  ))}
                </optgroup>
              </select>
              {pendingCloudModel && (() => {
                const badge = getStageBadgeMeta(resolveModelStageKind(pendingCloudModel, 'target'));
                return (
                  <span className={`px-3 py-1.5 rounded-xl text-xs font-extrabold border shrink-0 ${badge.className}`}>
                    {badge.label}
                  </span>
                );
              })()}
            </div>

            {pendingCloudModel && (
              <div className="text-xs bg-bg-input/50 border border-border-main/50 rounded-xl px-3.5 py-2.5 flex flex-wrap items-center gap-x-4 gap-y-1">
                <span>
                  <span className="font-normal text-text-muted">Cloud: </span>
                  <span className="font-extrabold text-emerald-600 dark:text-emerald-400">
                    {(pendingCloudModel.cloudInfraModel?.targetCloud?.csp || 'CLOUD').toUpperCase()} ({pendingCloudModel.cloudInfraModel?.targetCloud?.region || '-'})
                  </span>
                </span>
                <span>
                  <span className="font-normal text-text-muted">Node Groups: </span>
                  <span className="font-extrabold text-text-main">
                    {pendingCloudModel.cloudInfraModel?.targetInfra?.nodeGroups?.length || 0}
                  </span>
                </span>
                {pendingCloudModel.description && (
                  <span className="truncate max-w-md">
                    <span className="font-normal text-text-muted">Info: </span>
                    <span className="font-extrabold text-text-main">{pendingCloudModel.description}</span>
                  </span>
                )}
              </div>
            )}

            <div className="flex flex-wrap gap-3 pt-1">
              <button
                type="button"
                onClick={handleLoadDesign}
                disabled={!pendingCloudModel}
                className={`px-5 py-2.5 rounded-xl text-sm font-extrabold flex items-center gap-2 transition cursor-pointer ${
                  !pendingCloudModel
                    ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                    : isTargetAlreadyLoaded
                      ? 'bg-emerald-500/20 border border-emerald-500/40 text-emerald-600 dark:text-emerald-400'
                      : 'bg-emerald-500 hover:bg-emerald-600 text-slate-950 shadow-md shadow-emerald-500/25'
                }`}
              >
                {isTargetAlreadyLoaded ? (
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
              {pendingCloudModel && !isBuiltInSampleModel(pendingCloudModel) && (
                <button
                  type="button"
                  onClick={handleDeleteDesign}
                  className="px-5 py-2.5 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-xl text-sm font-extrabold flex items-center gap-2 transition cursor-pointer"
                >
                  <Trash2 className="w-4 h-4" />
                  <span>Delete Model</span>
                </button>
              )}
              {userTargetModels.length > 0 && (
                <button
                  type="button"
                  onClick={handleClearAllCustomDesigns}
                  className="px-5 py-2.5 rounded-xl text-sm font-extrabold flex items-center gap-2 transition cursor-pointer bg-red-500/10 hover:bg-red-500/20 text-red-600 dark:text-red-400 border border-red-500/20"
                >
                  <Trash2 className="w-4 h-4" />
                  <span>Delete All Custom Models</span>
                </button>
              )}
            </div>
          </div>
        )}
      </div>

      {/* -------------------------------------------------------------
          STEP 2: Recommended Target Cloud Alternatives Review and Editing
         ------------------------------------------------------------- */}
      <div className={`glass-panel p-6 rounded-2xl transition-all duration-300 ${activeStep >= 2 || (Boolean(editedCandidate) && recommendationCandidates.length > 0) ? 'opacity-100' : 'opacity-40 pointer-events-none'}`}>
        <div className="flex items-center justify-between mb-4 border-b border-border-main/40 pb-3">
          <div className="flex items-center space-x-3">
            <span className={`w-7 h-7 rounded-full flex items-center justify-center text-sm font-extrabold ${activeStep > 2 ? 'bg-green-500/20 text-green-600 dark:text-green-400 border border-green-500/30' : 'bg-emerald-500 text-slate-950'}`}>
              {activeStep > 2 ? '✓' : '2'}
            </span>
            <h3 className="text-base font-extrabold text-text-main">
              Step 2: Recommended Target Cloud Alternatives Review and Editing
            </h3>
          </div>
          {editedCandidate && (
            <button
              type="button"
              onClick={() => setIsTargetJsonOpen(!isTargetJsonOpen)}
              className="px-3 py-1.5 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-sm font-bold transition cursor-pointer flex items-center text-emerald-600 dark:text-emerald-400"
            >
              <Copy className="w-3.5 h-3.5 mr-1" />
              JSON View
            </button>
          )}
        </div>

        {isTargetJsonOpen && editedCandidate && (
          <div className="space-y-3 mb-6">
            <div className="flex justify-between items-center bg-bg-input px-3.5 py-2 rounded-xl border border-border-main/50">
              <span className="text-sm text-text-muted font-mono">cloudInfraModel</span>
              <button
                type="button"
                onClick={() => {
                  const normalizedTarget = {
                    ...editedCandidate,
                    targetNlbList:
                      editedCandidate.targetNlbList && editedCandidate.targetNlbList.length > 0
                        ? editedCandidate.targetNlbList
                        : undefined,
                  };
                  navigator.clipboard.writeText(JSON.stringify(normalizedTarget, null, 2));
                  setIsTargetJsonCopied(true);
                  setTimeout(() => setIsTargetJsonCopied(false), 2000);
                }}
                className="px-3 py-1 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 hover:border-emerald-500/30 rounded-lg text-xs font-bold transition flex items-center cursor-pointer text-emerald-600 dark:text-emerald-400"
              >
                <Copy className="w-3.5 h-3.5 mr-1" /> {isTargetJsonCopied ? 'Copied!' : 'Copy JSON'}
              </button>
            </div>
            <pre className="text-sm font-mono text-slate-800 dark:text-emerald-400 bg-bg-panel p-3.5 rounded-xl border border-border-main overflow-y-auto max-h-[400px] select-text">
              {JSON.stringify(
                {
                  ...editedCandidate,
                  targetNlbList:
                    editedCandidate.targetNlbList && editedCandidate.targetNlbList.length > 0
                      ? editedCandidate.targetNlbList
                      : undefined,
                },
                null,
                2
              )}
            </pre>
          </div>
        )}

        {recommendationCandidates.length === 0 ? (
          <div className="py-8 text-center text-text-muted text-sm italic">
            No recommendation generated yet. Complete previous steps to run recommendations.
          </div>
        ) : (
          <div className="space-y-6">

            {/* Row-based layout: Candidates (Row 1), Recommended Cloud Summary (Row 2), Topology Visualization (Row 3) */}
            {editedCandidate && (
              editedCandidate.status === 'failed' ||
              editedCandidate.status === 'nothing-to-recommend' ||
              !editedCandidate.targetInfra?.nodeGroups ||
              !Array.isArray(editedCandidate.targetInfra.nodeGroups) ||
              editedCandidate.targetInfra.nodeGroups.length === 0 ? (
                <div className="bg-amber-500/10 border border-amber-500/30 rounded-2xl p-6 text-center space-y-3 my-4">
                  <AlertTriangle className="w-8 h-8 text-amber-500 mx-auto" />
                  <h4 className="text-base font-extrabold text-amber-600 dark:text-amber-400">
                    No Infrastructure Recommended for {editedCandidate.targetCloud?.csp ? editedCandidate.targetCloud.csp.toUpperCase() : 'Target Cloud'} ({editedCandidate.targetCloud?.region || 'Region'})
                  </h4>
                  <p className="text-sm text-text-muted max-w-xl mx-auto">
                    {editedCandidate.description || 'The target cloud provider or region could not find matching VM specifications/images for the requested source workloads.'}
                  </p>
                </div>
              ) : (
                <div className="flex flex-col space-y-6 pt-4 border-t border-border-main/20">

                  <ActiveModelBanner
                    mode={selectedCloudModel && !String(selectedCloudModel.id).startsWith('candidate-') ? 'loaded' : 'draft'}
                    name={
                      selectedCloudModel && !String(selectedCloudModel.id).startsWith('candidate-')
                        ? selectedCloudModel.name
                        : buildDefaultModelName('target', selectedSourceModel?.name, editedCandidate.targetCloud?.csp)
                    }
                    version={
                      selectedCloudModel && !String(selectedCloudModel.id).startsWith('candidate-')
                        ? selectedCloudModel.version
                        : undefined
                    }
                    stageKind={
                      selectedCloudModel && !String(selectedCloudModel.id).startsWith('candidate-')
                        ? resolveModelStageKind(selectedCloudModel, 'target')
                        : 'draft'
                    }
                    hasUnsavedChanges={hasUnsavedChanges}
                    updatedTime={selectedCloudModel?.updatedTime}
                    metrics={[
                      {
                        label: 'Cloud',
                        value: `${(editedCandidate.targetCloud?.csp || 'CLOUD').toUpperCase()} / ${editedCandidate.targetCloud?.region || '-'}`,
                        tone: 'emerald',
                      },
                      {
                        label: 'Node Groups',
                        value: String(editedCandidate.targetInfra?.nodeGroups?.length || 0),
                      },
                      {
                        label: 'Est. Cost',
                        value: `$${getEstimatedMonthlyCost(editedCandidate)}/mo`,
                        tone: 'teal',
                      },
                    ]}
                  />

                  {/* Row 1: Recommended Target Cloud Candidates (Horizontal Scroll) */}
                  <div className="bg-bg-panel/30 border border-border-main/50 rounded-2xl p-5 space-y-3 w-full">
                    <div className="flex flex-col sm:flex-row justify-between sm:items-center gap-3 border-b border-border-main/30 pb-3 mb-2">
                      <div>
                        <div className="flex items-center gap-2">
                          <Layers className="w-4 h-4 text-emerald-500" />
                          <span className="text-sm font-bold text-emerald-600 dark:text-emerald-400 font-mono">
                            Recommended Target Cloud Candidates
                          </span>
                          <span className="text-xs px-2 py-0.5 rounded-full bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 font-extrabold border border-emerald-500/20 font-mono">
                            {recommendationCandidates.length} Candidate{recommendationCandidates.length > 1 ? 's' : ''}
                          </span>
                        </div>
                        <span className="text-sm text-text-muted mt-0.5 block">
                          Select a recommended candidate to examine its topology and customize infrastructure settings.
                        </span>
                      </div>

                      {recommendationCandidates.length > 1 && (
                        <button
                          type="button"
                          onClick={() => setShowCompareModal(true)}
                          className="px-3.5 py-2 bg-emerald-500/10 border border-emerald-500/30 hover:bg-emerald-500/20 text-emerald-600 dark:text-emerald-400 rounded-xl text-sm font-bold transition cursor-pointer flex items-center gap-2 shrink-0 self-start sm:self-auto shadow-sm"
                        >
                          <Layers className="w-4 h-4 text-emerald-500" />
                          <span>Compare Matrix ({recommendationCandidates.length})</span>
                        </button>
                      )}
                    </div>

                    {/* Multi-CSP rollback banner if viewing single-CSP deep dive */}
                    {cachedMultiCandidates && cachedMultiCandidates.length > 0 && (
                      <div className="flex flex-wrap items-center justify-between gap-3 p-3.5 bg-emerald-500/10 border border-emerald-500/30 rounded-2xl mb-2">
                        <div className="flex items-center gap-2 text-xs font-bold text-emerald-600 dark:text-emerald-400">
                          <Sparkles className="w-4 h-4 shrink-0" />
                          <span>
                            Currently viewing Single-CSP Deep Dive for {targetPairs[0]?.csp?.toUpperCase()}. Previous Multi-CSP comparison results are saved.
                          </span>
                        </div>
                        <button
                          type="button"
                          onClick={handleRestoreMultiCandidates}
                          className="px-3.5 py-1.5 bg-bg-panel border border-border-main hover:border-emerald-500/40 text-text-main rounded-xl text-xs font-extrabold flex items-center gap-1.5 transition cursor-pointer shadow-sm hover:shadow"
                        >
                          <Undo2 className="w-3.5 h-3.5 text-emerald-500" />
                          <span>Back to Multi-CSP Comparison ({cachedMultiCandidates.length} CSPs)</span>
                        </button>
                      </div>
                    )}

                    {/* Horizontal scrollable candidate cards */}
                    <div className="overflow-x-auto pb-3 pt-1 flex items-stretch gap-4 scroll-smooth">
                      {(() => {
                        const isMultiCspResults = recommendationCandidates.length > 1 &&
                          new Set(recommendationCandidates.map(c => c.targetCloud?.csp?.toLowerCase())).size > 1;

                        return recommendationCandidates.map((c, idx) => {
                          const isActive = selectedCandidateIndex === idx;
                          const cspName = c.targetCloud?.csp ? c.targetCloud.csp.toUpperCase() : '';
                          const regionName = c.targetCloud?.region || '-';
                          const cost = getEstimatedMonthlyCost(c);
                          const nodeGroups = c.targetInfra?.nodeGroups || [];
                          const nodeCount = nodeGroups.reduce((acc, ng) => acc + (ng.nodeGroupSize || 0), 0);
                          const nodeGroupCount = nodeGroups.length;
                          const candNlbCount = (c.targetNlbList || []).length;
                          const candGpuNodeCount = nodeGroups.reduce((acc, ng) => {
                            const sInfo = getSpecInfo(c, ng.specId);
                            return getGpuInfo(sInfo, ng.specId, ng) ? acc + (ng.nodeGroupSize || 0) : acc;
                          }, 0);

                          return (
                            <button
                              key={idx}
                              type="button"
                              onClick={() => selectCandidate(idx)}
                              className={`p-5 rounded-2xl border text-left transition-all cursor-pointer flex flex-col justify-between space-y-3.5 relative shrink-0 w-[290px] min-w-[280px] ${
                                isActive
                                  ? 'bg-emerald-500/10 border-emerald-500 text-text-main shadow-md shadow-emerald-500/10 ring-1 ring-emerald-500'
                                  : 'bg-bg-panel/70 border-border-main text-text-muted hover:border-emerald-500/40 hover:bg-bg-panel hover:shadow-sm'
                              }`}
                            >
                              {/* Card Header: Candidate Title & CSP Badge */}
                              <div className="flex items-center justify-between gap-2 w-full">
                                <div className="flex items-center gap-2.5 min-w-0">
                                  <span className={`w-6 h-6 rounded-full flex items-center justify-center text-xs font-mono font-extrabold shrink-0 ${
                                    isActive
                                      ? 'bg-emerald-500 text-slate-950 shadow-sm'
                                      : 'bg-bg-input text-text-muted border border-border-main'
                                  }`}>
                                    {idx + 1}
                                  </span>
                                  <span className="text-sm font-extrabold text-text-main truncate">
                                    Candidate {idx + 1}
                                  </span>
                                </div>
                                {cspName && (
                                  <span className={`text-xs px-2.5 py-0.5 rounded font-mono font-extrabold uppercase shrink-0 ${
                                    isActive
                                      ? 'bg-emerald-500/20 text-emerald-600 dark:text-emerald-300 border border-emerald-500/30'
                                      : 'bg-bg-input text-text-muted border border-border-main'
                                  }`}>
                                    {cspName}
                                  </span>
                                )}
                              </div>

                              {/* Card Body: Values Only (No Keys!) */}
                              <div className="space-y-2 w-full pt-1.5 border-t border-border-main/15">
                                {/* Target Region Value */}
                                <div className="font-extrabold text-text-main text-sm truncate" title={regionName}>
                                  {regionName}
                                </div>

                                {/* Node Allocation Value */}
                                <div className="font-extrabold text-text-main text-sm truncate">
                                  {nodeGroupCount} Node Group(s), {nodeCount} Node(s)
                                </div>

                                {/* Match Level, GPU, and NLB Badges */}
                                <div className="flex flex-wrap items-center gap-1.5">
                                  <span className={`inline-block text-xs uppercase font-extrabold px-2.5 py-0.5 rounded ${
                                    c.status === 'highly-matched'
                                      ? 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-emerald-500/20'
                                      : 'bg-amber-500/10 text-amber-600 dark:text-amber-400 border border-amber-500/20'
                                  }`}>
                                    {c.status || 'matched'}
                                  </span>
                                  {candGpuNodeCount > 0 && (
                                    <span className="inline-block text-xs font-extrabold px-2 py-0.5 rounded bg-emerald-500/15 text-emerald-600 dark:text-emerald-400 border border-emerald-500/30">
                                      GPU ({candGpuNodeCount})
                                    </span>
                                  )}
                                  {candNlbCount > 0 && (
                                    <span className="inline-block text-xs font-extrabold px-2 py-0.5 rounded bg-teal-500/15 text-teal-600 dark:text-teal-400 border border-teal-500/30">
                                      NLB ({candNlbCount})
                                    </span>
                                  )}
                                </div>
                              </div>

                              {/* Card Footer: Est. Cost Value Only & Select State */}
                              <div className="pt-2.5 border-t border-border-main/20 flex items-center justify-between w-full">
                                <span className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 tabular-nums">
                                  ${cost}/month
                                </span>

                                <span className={`text-xs font-bold px-2.5 py-1 rounded-lg transition ${
                                  isActive
                                    ? 'bg-emerald-500 text-slate-950 font-extrabold shadow-sm flex items-center gap-1'
                                    : 'bg-bg-input text-text-muted border border-border-main hover:text-text-main'
                                }`}>
                                  {isActive && <CheckCircle2 className="w-3.5 h-3.5" />}
                                  {isActive ? 'Selected' : 'Select'}
                                </span>
                              </div>

                              {/* Deep Dive Action (Visible in Multi-CSP mode or when results contain multiple CSPs) */}
                              {(recommendationMode === 'multi' || isMultiCspResults) && c.targetCloud?.csp && (
                                <div
                                  onClick={(e) => {
                                    e.stopPropagation();
                                    handleDeepDiveWithCsp(c.targetCloud.csp, c.targetCloud.region || 'ap-northeast-2');
                                  }}
                                  className="w-full pt-1"
                                >
                                  <span className="w-full py-1.5 px-2.5 bg-emerald-500/10 hover:bg-emerald-500/20 border border-emerald-500/30 hover:border-emerald-500/50 text-emerald-600 dark:text-emerald-400 rounded-xl text-xs font-extrabold flex items-center justify-center gap-1.5 transition cursor-pointer shadow-sm">
                                    <Sparkles className="w-3.5 h-3.5 text-emerald-500" />
                                    <span>Deep Dive with {cspName} (5 Candidates) ➔</span>
                                  </span>
                                </div>
                              )}
                            </button>
                          );
                        });
                      })()}
                    </div>
                  </div>

                  {/* Row 2: 2-Line Summary Bar */}
                  <div className="bg-bg-panel/40 border border-border-main/50 rounded-xl p-4 sm:px-5 sm:py-3.5 w-full flex flex-col gap-2.5 text-sm shadow-sm">
                    {/* Line 1: Summary, Candidate, Match, Cost */}
                    <div className="flex flex-wrap items-center gap-x-5 gap-y-2">
                      <div className="inline-flex items-center gap-2 shrink-0 pr-4 border-r border-border-main/30">
                        <Compass className="w-4 h-4 text-emerald-500 shrink-0" />
                        <span className="text-sm font-extrabold text-text-main">
                          Summary
                        </span>
                        <span className="text-xs px-2 py-0.5 rounded-full bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 font-extrabold border border-emerald-500/20">
                          Candidate {selectedCandidateIndex + 1}
                        </span>
                      </div>

                      {/* Match */}
                      <div className="inline-flex items-center gap-1.5 shrink-0">
                        <span className="text-sm font-normal text-text-muted leading-none">Match</span>
                        <span className="text-xs uppercase font-extrabold px-2 py-0.5 rounded-md bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-emerald-500/20 leading-tight">
                          {editedCandidate.status}
                        </span>
                      </div>

                      <div className="h-4 w-px bg-border-main/30 shrink-0 hidden sm:block self-center" />

                      {/* Cost */}
                      <div className="inline-flex items-baseline gap-1.5 shrink-0">
                        <span className="text-sm font-normal text-text-muted">Cost</span>
                        <span className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 tabular-nums">
                          ${getEstimatedMonthlyCost(editedCandidate)}/month
                        </span>
                      </div>
                    </div>

                    <div className="h-px w-full bg-border-main/20" />

                    {/* Line 2: Network, Overall Compute, Accelerator, Security, SSH Key, NLB */}
                    <div className="flex flex-wrap items-center gap-x-5 gap-y-2">
                      {/* Network */}
                      <div className="inline-flex items-baseline gap-1.5 shrink-0">
                        <span className="text-sm font-normal text-text-muted">Network</span>
                        <span className="text-sm font-extrabold text-text-main tabular-nums">
                          1 VNet, {editedCandidate.targetVNet?.subnetInfoList?.length || 0} Subnet(s)
                        </span>
                      </div>

                      <div className="h-4 w-px bg-border-main/30 shrink-0 hidden sm:block self-center" />

                      {/* Overall Compute */}
                      <div className="inline-flex items-baseline gap-1.5 shrink-0">
                        <span className="text-sm font-normal text-text-muted">Overall Compute</span>
                        <span className="text-sm font-extrabold text-text-main tabular-nums">
                          {(editedCandidate.targetInfra?.nodeGroups || []).length} Node Group(s), {(editedCandidate.targetInfra?.nodeGroups || []).reduce((acc, ng) => acc + (ng.nodeGroupSize || 0), 0)} Node(s)
                        </span>
                      </div>

                      <div className="h-4 w-px bg-border-main/30 shrink-0 hidden sm:block self-center" />

                      {/* Accelerator */}
                      <div className="inline-flex items-baseline gap-1.5 shrink-0">
                        <span className="text-sm font-normal text-text-muted">Accelerator</span>
                        {candidateGpuSummary.hasGpu && candidateGpuSummary.gpuNgCount > 0 ? (
                          <span className="text-sm font-extrabold text-text-main tabular-nums">
                            {candidateGpuSummary.gpuNgCount} Node Group(s), {candidateGpuSummary.gpuNodeCount} Node(s)
                          </span>
                        ) : (
                          <span className="text-sm font-extrabold text-text-muted tabular-nums">
                            0 Node(s)
                          </span>
                        )}
                      </div>

                      <div className="h-4 w-px bg-border-main/30 shrink-0 hidden sm:block self-center" />

                      {/* Security */}
                      <div className="inline-flex items-baseline gap-1.5 shrink-0">
                        <span className="text-sm font-normal text-text-muted">Security</span>
                        <span className="text-sm font-extrabold text-text-main tabular-nums">
                          {(editedCandidate.targetSecurityGroupList || []).length} SG(s), {(editedCandidate.targetSecurityGroupList || []).reduce((acc, sg) => acc + (sg?.firewallRules || []).length, 0)} Rule(s)
                        </span>
                      </div>

                      <div className="h-4 w-px bg-border-main/30 shrink-0 hidden sm:block self-center" />

                      {/* SSH Key */}
                      <div className="inline-flex items-baseline gap-1.5 shrink-0">
                        <span className="text-sm font-normal text-text-muted">SSH Key</span>
                        <span className="text-sm font-extrabold text-text-main tabular-nums">
                          {editedCandidate.targetSshKey ? 1 : 0} Key(s)
                        </span>
                      </div>

                      <div className="h-4 w-px bg-border-main/30 shrink-0 hidden sm:block self-center" />

                      {/* Load Balancer */}
                      <div className="inline-flex items-baseline gap-1.5 shrink-0">
                        <span className="text-sm font-normal text-text-muted">NLB</span>
                        <span className="text-sm font-extrabold text-text-main tabular-nums">
                          {(editedCandidate.targetNlbList || []).length} NLB(s)
                        </span>
                      </div>
                    </div>
                  </div>

                  {/* Row 3: Topology Visualization */}

                  <div className="w-full bg-bg-panel/40 border border-border-main/50 rounded-2xl p-5 relative min-h-[300px] flex flex-col justify-between">
                    <div>
                      <div className="flex justify-between items-center mb-4">
                        <span className="text-sm font-bold text-emerald-600 dark:text-emerald-400">Topology Visualization</span>
                        <span className="text-sm text-text-muted font-mono">{editedCandidate.targetCloud?.region} ({editedCandidate.targetCloud?.csp?.toUpperCase()})</span>
                      </div>
                      {/* Simulated Topology Drawing */}
                      <div className="flex flex-col space-y-4 pt-4 text-sm">

                        {/* VPC / VNet Container */}
                        <div className="border border-emerald-400 dark:border-emerald-800/40 bg-emerald-500/5 rounded-2xl p-5 relative">
                          <div className="flex flex-col md:flex-row justify-between md:items-center gap-3 mb-4 border-b border-emerald-200 dark:border-emerald-800/20 pb-3">
                            <span className="font-extrabold text-emerald-600 dark:text-emerald-400 flex items-center space-x-1.5 font-mono">
                              <Network className="w-4 h-4 animate-pulse" />
                              <span>VPC / VNet: {editedCandidate.targetVNet?.name || 'vnet'} ({editedCandidate.targetVNet?.cidrBlock || ''})</span>
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
                                    <span>SG: {sg.name || 'default-sg'} ({(sg.firewallRules || []).length} Rule(s))</span>
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

                                const matchingNg = (editedCandidate.targetInfra?.nodeGroups || []).find(
                                  (ng) => ng.name === nlb.targetGroup?.nodeGroupId
                                ) || (editedCandidate.targetInfra?.nodeGroups || [])[0];

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
                            ng => !ng.subnetId || !allSubnetNames.includes(ng.subnetId)
                          );
                          const displayNgs = subIdx === 0
                            ? [...assignedNgs, ...unassignedNgs]
                            : assignedNgs;

                          return (
                        <div key={subIdx} className="border border-dashed border-emerald-400 dark:border-emerald-800/30 bg-bg-panel/40 rounded-xl p-4 space-y-3">
                          <div className="flex items-center justify-between text-sm text-text-muted font-mono pb-2 border-b border-border-main">
                            <span>
                              Subnet: {sub.name} ({sub.ipv4_CIDR})
                              {sub.zone && <span className="ml-2 text-xs text-emerald-600 dark:text-emerald-400/70">Zone: {sub.zone}</span>}
                            </span>
                            <span className="px-2 py-0.5 rounded bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold">
                              {displayNgs.length} Node Group(s)
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
                              const gpuInfo = getGpuInfo(specInfo, ng.specId, ng);
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

                                    {/* Spec: vCPU / Memory / GPU / instance type */}
                                    <div className="flex items-center gap-1 flex-shrink-0 flex-wrap">
                                      {vcpu   && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">{vcpu} vCPU</span>}
                                      {memGiB && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">{formatMemory(memGiB)}</span>}
                                      {gpuInfo && (
                                        <span className="bg-amber-500/10 border border-amber-500/20 text-amber-600 dark:text-amber-400 text-xs font-bold px-1.5 py-0.5 rounded font-mono flex items-center gap-1" title={`Target Cloud GPU: ${gpuInfo.label}`}>
                                          <span>{gpuInfo.label !== 'GPU' ? `GPU: ${gpuInfo.label}` : 'GPU'}</span>
                                        </span>
                                      )}
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

                                    {/* Node count - push to right */}
                                    <span className="ml-auto flex-shrink-0 text-xs font-bold bg-emerald-100 dark:bg-emerald-950/40 border border-emerald-300 dark:border-emerald-800/40 text-emerald-600 dark:text-emerald-400 px-1.5 py-0.5 rounded">×{nodeCount}</span>
                                  </button>

                                  {/* Expanded: detail summary + VM nodes */}
                                  {isExpanded && (
                                    <div className="border-t border-border-main/20 px-4 py-3 space-y-3 bg-bg-input/5">

                                      {/* Top summary: Spec / Image / Root Disk / SG / Size */}
                                      <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-3">
                                        {/* Spec */}
                                        <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                          <span className="block text-xs font-bold text-text-muted mb-1">Spec</span>
                                          <div className="flex flex-wrap gap-1">
                                            {vcpu   && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">{vcpu} vCPU</span>}
                                            {memGiB && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-1.5 py-0.5 rounded">{formatMemory(memGiB)}</span>}
                                            {gpuInfo && (
                                              <span className="bg-amber-500/10 border border-amber-500/20 text-amber-600 dark:text-amber-400 text-xs font-bold px-1.5 py-0.5 rounded font-mono" title={`Target Cloud GPU: ${gpuInfo.label}`}>
                                                {gpuInfo.label !== 'GPU' ? `GPU: ${gpuInfo.label}` : 'GPU'}
                                              </span>
                                            )}
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
                                          {ng.cspImageName && (
                                            <span className="block text-xs font-mono text-text-muted mt-1 break-all" title={ng.cspImageName}>
                                              CSP: {ng.cspImageName}
                                            </span>
                                          )}
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
                                              : <span className="text-xs text-text-muted">-</span>
                                            }
                                          </div>
                                        </div>

                                        {/* Size (node count) */}
                                        <div className="bg-bg-panel border border-border-main/30 rounded-lg px-3 py-2">
                                          <span className="block text-xs font-bold text-text-muted mb-1">Size</span>
                                          <span className="text-2xl font-extrabold text-emerald-600 dark:text-emerald-400">{nodeCount}</span>
                                          <span className="text-xs text-text-muted ml-1">node{nodeCount > 1 ? 's' : ''}</span>
                                          {(ng.label?.['cm-source-machine-ids'] || ng.label?.sourceMachineIds) && (
                                            <span className="block text-xs font-mono text-emerald-600 dark:text-emerald-400 mt-1 break-all">
                                              Source: {ng.label?.['cm-source-machine-ids'] || ng.label?.sourceMachineIds}
                                            </span>
                                          )}
                                        </div>
                                      </div>

                                      {/* VM node list */}
                                      <div className="space-y-1">
                                        <span className="text-xs text-text-muted font-semibold">Nodes</span>
                                        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-1.5 max-h-40 overflow-y-auto">
                                          {nodesArray.map((_, nodeIdx) => {
                                            const nodeName = `${ng.name}-${nodeIdx + 1}`;
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
                                <button
                                  type="button"
                                  onClick={handleValidateNaming}
                                  disabled={namingValidation.loading}
                                  className="px-2.5 py-1.5 bg-bg-panel border border-border-main hover:border-emerald-500/40 text-text-main hover:text-emerald-500 rounded-lg text-xs font-extrabold transition cursor-pointer flex items-center gap-1.5"
                                  title="Validate referential integrity of resource names across VNet, Subnets, SSH Key, and Security Groups"
                                >
                                  {namingValidation.loading ? (
                                    <Loader2 className="w-3.5 h-3.5 animate-spin text-emerald-500" />
                                  ) : (
                                    <ShieldCheck className="w-3.5 h-3.5 text-emerald-500" />
                                  )}
                                  <span>Validate Naming</span>
                                </button>
                              )}
                              {namingValidation.message && (
                                <span
                                  className={`px-2.5 py-1 rounded-lg text-xs font-extrabold border flex items-center gap-1 ${
                                    namingValidation.valid === false
                                      ? 'bg-red-500/10 border-red-500/30 text-red-500'
                                      : 'bg-emerald-500/10 border-emerald-500/30 text-emerald-600 dark:text-emerald-400'
                                  }`}
                                >
                                  <span>{namingValidation.message}</span>
                                </span>
                              )}
                              {editedCandidate && (
                                <span className="text-xs font-extrabold text-emerald-600 dark:text-emerald-400 font-mono bg-emerald-500/10 border border-emerald-500/20 px-2.5 py-1.5 rounded-lg flex items-center gap-1 shadow-sm">
                                  <DollarSign className="w-3.5 h-3.5 text-emerald-500" />
                                  <span>Cost: ${getEstimatedMonthlyCost(editedCandidate)}/month</span>
                                </span>
                              )}
                              {selectedCloudModel ? (
                                <span className="px-3.5 py-1.5 bg-emerald-500/10 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 rounded-xl text-xs font-extrabold font-mono">
                                  Loaded Design: {selectedCloudModel.name} (v{selectedCloudModel.version})
                                </span>
                              ) : (
                                <>
                                  <span className="text-xs font-normal text-text-muted">Candidates:</span>
                                  {recommendationCandidates.map((c, idx) => {
                                    const isActive = selectedCandidateIndex === idx;
                                    const cspName = c.targetCloud?.csp ? c.targetCloud.csp.toUpperCase() : '';
                                    const regionName = c.targetCloud?.region || '';
                                    return (
                                      <button
                                        key={idx}
                                        onClick={() => {
                                          selectCandidate(idx);
                                        }}
                                        className={`px-3 py-1.5 rounded-lg text-xs font-bold border transition cursor-pointer flex items-center gap-1.5 ${
                                          isActive
                                            ? 'bg-emerald-500/10 border-emerald-500/40 text-emerald-600 dark:text-emerald-400 font-extrabold shadow-sm'
                                            : 'bg-bg-panel border-border-main text-text-muted hover:text-text-main'
                                        }`}
                                      >
                                        <span>c {idx + 1}</span>
                                        {cspName && (
                                          <span className={`text-xs px-1.5 py-0.5 rounded font-mono font-extrabold ${
                                            isActive ? 'bg-emerald-500/20 text-emerald-600 dark:text-emerald-300' : 'bg-bg-input text-text-muted'
                                          }`}>
                                            {cspName}
                                          </span>
                                        )}
                                      </button>
                                    );
                                  })}
                                  {recommendationCandidates.length > 1 && (
                                    <button
                                      onClick={() => setShowCompareModal(true)}
                                      className="px-3 py-1.5 bg-bg-panel border border-border-main hover:bg-emerald-500/10 hover:border-emerald-500/30 text-emerald-600 dark:text-emerald-400 rounded-lg text-xs font-bold transition cursor-pointer"
                                    >
                                      Compare
                                    </button>
                                  )}
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
                                NLB - Network Load Balancer
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
                                  const sourceMachineIds = ng.label?.['cm-source-machine-ids'] || ng.label?.sourceMachineIds || '';
                                  return (
                                  <div key={ngIdx} className="bg-bg-panel/40 p-4 rounded-xl border border-border-main/30 space-y-3">
                                    <div className="flex flex-wrap justify-between items-center gap-2 text-sm font-bold text-text-main border-b border-border-main/10 pb-1.5 font-mono">
                                      <div className="flex items-center gap-2 flex-wrap">
                                        <span>Node Group #{ngIdx + 1} - {extractInstanceType(ng.specId)}</span>
                                        {sourceMachineIds && (
                                          <span className="px-2 py-0.5 rounded bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold">
                                            Source: {sourceMachineIds}
                                          </span>
                                        )}
                                      </div>
                                      <span className="text-emerald-600 dark:text-emerald-400">Spec Match Type: Balanced</span>
                                    </div>

                                    {/* Spec & Image summary badges */}
                                    <div className="flex flex-wrap gap-3 p-3 bg-bg-input/20 rounded-lg border border-border-main/10">
                                      {/* Spec badges */}
                                      <div className="flex flex-col gap-1">
                                        <span className="text-xs font-bold text-text-muted">Node Spec</span>
                                        <div className="flex items-center gap-1.5 flex-wrap">
                                          {vcpu   && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-2 py-0.5 rounded">{vcpu} vCPU</span>}
                                          {memGiB && <span className="bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold px-2 py-0.5 rounded">{formatMemory(memGiB)}</span>}
                                          {(() => {
                                            const step3Gpu = getGpuInfo(specInfo, ng.specId, ng);
                                            return step3Gpu ? (
                                              <span className="bg-amber-500/10 border border-amber-500/20 text-amber-600 dark:text-amber-400 text-xs font-bold px-2 py-0.5 rounded font-mono" title={`Target Cloud GPU: ${step3Gpu.label}`}>
                                                {step3Gpu.label !== 'GPU' ? `GPU: ${step3Gpu.label}` : 'GPU'}
                                              </span>
                                            ) : null;
                                          })()}
                                          <span className="text-xs text-text-muted font-mono" title={ng.specId}>{extractInstanceType(ng.specId)}</span>
                                        </div>
                                      </div>
                                      <div className="w-px bg-border-main/30 self-stretch mx-1 hidden sm:block" />
                                      {/* Image badges - flex-1 so it uses remaining width without truncation */}
                                      <div className="flex flex-col gap-1 flex-1 min-w-[180px]">
                                        <span className="text-xs font-bold text-text-muted">OS Image</span>
                                        <div className="flex flex-wrap gap-1.5 items-start">
                                          {osName && <span className="bg-teal-500/10 border border-teal-500/20 text-teal-600 dark:text-teal-400 text-xs font-extrabold px-2 py-0.5 rounded break-words">{osName}</span>}
                                          <span className="text-xs text-text-muted font-mono break-all" title={ng.imageId}>{ng.imageId.split('/').pop() || ng.imageId}</span>
                                          {ng.cspImageName && (
                                            <span className="text-xs font-mono text-emerald-600 dark:text-emerald-400 break-all">
                                              (CSP Image: {ng.cspImageName})
                                            </span>
                                          )}
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
                                      {/* Node Count - last column */}
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
                                          <span className="text-xs text-text-muted italic">No security groups defined - add one in the Security Groups tab</span>
                                        )}
                                      </div>
                                    </div>
                                  </div>
                                  );
                                })}

                                {/* Post-Deployment Bootstrap Commands (targetInfra.postCommands & postCommandAsync) */}
                                <div className="bg-bg-panel/40 p-4 rounded-xl border border-border-main/30 space-y-3">
                                  <div className="flex flex-wrap items-center justify-between gap-2 border-b border-border-main/10 pb-2">
                                    <div>
                                      <span className="text-sm font-extrabold text-text-main block">Post-Deployment Bootstrap Commands (postCommands)</span>
                                      <span className="text-xs text-text-muted">Configure remote commands executed on target nodes after provisioning</span>
                                    </div>
                                    <div className="flex items-center gap-3">
                                      <label className="inline-flex items-center gap-2 text-xs text-text-main cursor-pointer">
                                        <input
                                          type="checkbox"
                                          checked={Boolean(editedCandidate.targetInfra.postCommandAsync)}
                                          onChange={(e) => handleTogglePostCommandAsync(e.target.checked)}
                                          className="rounded border-border-main text-emerald-500 focus:ring-emerald-500"
                                        />
                                        <span className="font-normal text-text-muted">Async Execution (</span>
                                        <span className="font-extrabold text-text-main font-mono">postCommandAsync</span>
                                        <span className="font-normal text-text-muted">)</span>
                                      </label>
                                      <button
                                        type="button"
                                        onClick={handleAddPostCommand}
                                        className="px-3 py-1.5 bg-emerald-500 hover:bg-emerald-600 text-slate-950 rounded-lg text-xs font-extrabold transition cursor-pointer"
                                      >
                                        + Add Command Block
                                      </button>
                                    </div>
                                  </div>

                                  {(editedCandidate.targetInfra.postCommands || []).length === 0 ? (
                                    <p className="text-xs text-text-muted italic">
                                      No post-deployment commands configured. Click &quot;+ Add Command Block&quot; to attach bootstrap scripts to all nodes, a specific Node Group, or a specific Node.
                                    </p>
                                  ) : (
                                    <div className="space-y-3">
                                      {(editedCandidate.targetInfra.postCommands || []).map((cmdBlock, cmdIdx) => (
                                        <div key={cmdIdx} className="bg-bg-input/20 border border-border-main/20 p-3 rounded-lg space-y-2.5">
                                          <div className="flex items-center justify-between">
                                            <span className="text-xs font-extrabold text-text-main font-mono">
                                              Command Block #{cmdIdx + 1}
                                            </span>
                                            <button
                                              type="button"
                                              onClick={() => handleRemovePostCommand(cmdIdx)}
                                              className="px-2 py-1 text-xs font-bold text-red-400 hover:bg-red-500/10 rounded transition cursor-pointer"
                                            >
                                              Remove
                                            </button>
                                          </div>
                                          <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
                                            <div>
                                              <label className="block text-xs font-normal text-text-muted mb-1">
                                                Target Node Group (optional)
                                              </label>
                                              <input
                                                type="text"
                                                placeholder="Leave empty for all Node Groups"
                                                value={cmdBlock.nodeGroupName || ''}
                                                onChange={(e) => handleUpdatePostCommand(cmdIdx, 'nodeGroupName', e.target.value)}
                                                className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                              />
                                            </div>
                                            <div>
                                              <label className="block text-xs font-normal text-text-muted mb-1">
                                                Target Node Name (optional)
                                              </label>
                                              <input
                                                type="text"
                                                placeholder="Leave empty for all Nodes in group"
                                                value={cmdBlock.nodeName || ''}
                                                onChange={(e) => handleUpdatePostCommand(cmdIdx, 'nodeName', e.target.value)}
                                                className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                              />
                                            </div>
                                          </div>
                                          <div>
                                            <label className="block text-xs font-normal text-text-muted mb-1">
                                              Shell Commands (one command per line)
                                            </label>
                                            <textarea
                                              rows={3}
                                              value={(cmdBlock.commands || []).join('\n')}
                                              onChange={(e) => handleUpdatePostCommand(cmdIdx, 'commandsText', e.target.value)}
                                              className="w-full bg-bg-input border border-border-main/50 text-text-main rounded-lg px-2.5 py-1.5 text-sm font-mono focus:outline-none focus:border-emerald-500/40"
                                              placeholder="sudo apt-get update -y"
                                            />
                                          </div>
                                        </div>
                                      ))}
                                    </div>
                                  )}
                                </div>
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
                                      onFocus={(e) => setNamingFocusOldName(e.target.value)}
                                      onChange={(e) => handleTuneTargetVNetProperty('name', e.target.value)}
                                      onBlur={(e) => handleAlignResourceNameOnBlur('vNet', namingFocusOldName, e.target.value)}
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
                                        <span className="text-xs font-bold text-text-muted">Subnet #{subIdx + 1}</span>
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
                                            onFocus={(e) => setNamingFocusOldName(e.target.value)}
                                            onChange={(e) => handleTuneTargetSubnetProperty(subIdx, 'name', e.target.value)}
                                            onBlur={(e) => handleAlignResourceNameOnBlur('subnet', namingFocusOldName, e.target.value)}
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
                                      onFocus={(e) => setNamingFocusOldName(e.target.value)}
                                      onChange={(e) => handleTuneTargetSshKeyProperty('name', e.target.value)}
                                      onBlur={(e) => handleAlignResourceNameOnBlur('sshKey', namingFocusOldName, e.target.value)}
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
                                          onFocus={(e) => setNamingFocusOldName(e.target.value)}
                                          onChange={(e) => handleTuneTargetSecurityGroupProperty(sgIdx, 'name', e.target.value)}
                                          onBlur={(e) => handleAlignResourceNameOnBlur('securityGroup', namingFocusOldName, e.target.value)}
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

                          {/* Full-width CSP Quota & Deployment Warning Banner */}
                          <div className="flex items-center gap-2.5 px-4 py-3 rounded-xl bg-amber-500/10 border border-amber-500/20 text-amber-700 dark:text-amber-400 text-xs">
                            <span className="font-bold shrink-0">⚠️ Note:</span>
                            <span>Resource creation may fail due to account-level CSP quota limits, regional availability, or insufficient instance stock. Verify your account quotas before deployment.</span>
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
                                <span>Back to 2. Migration Planning</span>
                              </button>
                            ) : <div />}

                            <div className="flex flex-wrap items-center gap-2.5 w-full sm:w-auto justify-end">
                              <button
                                type="button"
                                onClick={handleDryRunValidate}
                                disabled={isValidatingDryRun}
                                className="w-full sm:w-auto px-4 py-2.5 bg-bg-panel border border-emerald-500/40 hover:bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 rounded-xl text-xs font-extrabold flex items-center justify-center space-x-1.5 transition cursor-pointer shadow-sm disabled:opacity-50"
                              >
                                {isValidatingDryRun ? (
                                  <RefreshCw className="w-3.5 h-3.5 animate-spin text-emerald-500" />
                                ) : (
                                  <ShieldCheck className="w-3.5 h-3.5 text-emerald-500" />
                                )}
                                <span>{isValidatingDryRun ? 'Checking Consistency...' : 'Pre-Flight Validation (Dry-Run)'}</span>
                              </button>

                              <button
                                type="button"
                                onClick={() => setShowSaveTargetModal(true)}
                                className="w-full sm:w-auto px-5 py-2.5 bg-gradient-to-r from-emerald-400 via-teal-400 to-teal-500 hover:from-emerald-500 hover:to-teal-600 text-slate-950 rounded-xl text-xs font-extrabold flex items-center justify-center space-x-1.5 transition cursor-pointer shadow-md shadow-emerald-500/20"
                              >
                                <Save className="w-3.5 h-3.5 text-slate-950" />
                                <span>Save Target Cloud Model</span>
                              </button>

                              {onNext && (
                                <button
                                  type="button"
                                  onClick={handleProceedToMigration}
                                  className="w-full sm:w-auto px-5 py-2.5 bg-emerald-600 hover:bg-emerald-700 text-white font-extrabold text-xs rounded-xl shadow-lg shadow-emerald-500/20 transition flex items-center justify-center space-x-1.5 cursor-pointer"
                                >
                                  <span>Next: Proceed to 4. Migration Execution</span>
                                  <ArrowRight className="w-3.5 h-3.5" />
                                </button>
                              )}
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>

              </div>
            ))}
          </div>
        )}
      </div>

      {/* Compare Candidates Modal */}
      {showCompareModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-[92vw] border border-border-main animate-scale-up flex flex-col max-h-[85vh]">

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
              <MatchRateRadarChart
                candidates={recommendationCandidates}
                sourceInfra={selectedSourceModel?.onpremiseInfraModel}
              />
              <div className="overflow-x-auto border border-border-main/50 rounded-xl bg-bg-panel/40">
                <table className="w-full text-left border-collapse text-sm">
                  <thead>
                    <tr className="border-b border-border-main bg-bg-input/60 text-text-muted font-normal">
                      <th className="py-3 px-4 min-w-[170px] whitespace-nowrap font-normal text-text-muted">Metric</th>
                      {recommendationCandidates.map((c, idx) => (
                        <th key={idx} className="py-3 px-4 min-w-[320px] whitespace-nowrap font-extrabold text-text-main">
                          c {idx + 1} {c.targetCloud?.csp ? `(${c.targetCloud.csp.toUpperCase()})` : ''}
                        </th>
                      ))}
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-border-main/40">
                    <tr>
                      <td className="py-3.5 px-4 align-top font-normal text-text-muted bg-bg-input/10 whitespace-nowrap">Target Cloud</td>
                      {recommendationCandidates.map((c, idx) => (
                        <td key={idx} className="py-3.5 px-4 align-top text-emerald-600 dark:text-emerald-400 font-extrabold font-mono text-sm whitespace-nowrap">
                          {c.targetCloud?.csp ? c.targetCloud.csp.toUpperCase() : '-'} / {c.targetCloud?.region || '-'}
                        </td>
                      ))}
                    </tr>
                    <tr>
                      <td className="py-3.5 px-4 align-top font-normal text-text-muted bg-bg-input/10 whitespace-nowrap">Match Level</td>
                      {recommendationCandidates.map((c, idx) => (
                        <td key={idx} className="py-3.5 px-4 align-top whitespace-nowrap">
                          <span className={`px-2.5 py-1 rounded-full text-sm font-extrabold uppercase ${c.status === 'highly-matched'
                              ? 'bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-emerald-500/20'
                              : 'bg-amber-500/10 text-amber-600 dark:text-amber-400 border border-amber-500/20'
                            }`}>
                            {c.status}
                          </span>
                        </td>
                      ))}
                    </tr>
                    <tr>
                      <td className="py-3.5 px-4 align-top font-normal text-text-muted bg-bg-input/10 whitespace-nowrap">Node Spec Allocation</td>
                      {recommendationCandidates.map((c, idx) => (
                        <td key={idx} className="py-3.5 px-4 align-top text-text-main font-mono text-sm">
                          <div className="divide-y divide-border-main/15">
                            {c.targetInfra.nodeGroups.map((ng, i) => (
                              <div key={i} className="flex items-center justify-between gap-4 py-1.5 min-h-[34px] first:pt-0 last:pb-0 whitespace-nowrap">
                                <span className="text-text-muted font-normal shrink-0">{ng.name}:</span>
                                <div className="flex items-center justify-end gap-1.5 shrink-0 text-right">
                                  <span className="font-extrabold text-emerald-600 dark:text-emerald-400">{ng.specId} (x{ng.nodeGroupSize})</span>
                                  {(() => {
                                    const matrixGpu = getGpuInfo(getSpecInfo(c, ng.specId), ng.specId, ng);
                                    return matrixGpu ? (
                                      <span className="bg-amber-500/10 border border-amber-500/20 text-amber-600 dark:text-amber-400 text-xs font-extrabold px-1.5 py-0.5 rounded font-mono shrink-0" title={`Target Cloud GPU: ${matrixGpu.label}`}>
                                        {matrixGpu.label !== 'GPU' ? `GPU: ${matrixGpu.label}` : 'GPU'}
                                      </span>
                                    ) : null;
                                  })()}
                                </div>
                              </div>
                            ))}
                          </div>
                        </td>
                      ))}
                    </tr>
                    <tr>
                      <td className="py-3.5 px-4 align-top font-normal text-text-muted bg-bg-input/10 whitespace-nowrap">Estimated Cost</td>
                      {recommendationCandidates.map((c, idx) => (
                        <td key={idx} className="py-3.5 px-4 align-top text-emerald-600 dark:text-emerald-400 font-extrabold font-mono whitespace-nowrap">
                          <div className="flex items-center space-x-1">
                            <DollarSign className="w-3.5 h-3.5" />
                            <span>{getEstimatedMonthlyCost(c)}</span>
                            <span className="text-sm font-normal text-text-muted">/ month</span>
                          </div>
                        </td>
                      ))}
                    </tr>
                    <tr>
                      <td className="py-3.5 px-4 align-top font-normal text-text-muted bg-bg-input/10 whitespace-nowrap">Description Summary</td>
                      {recommendationCandidates.map((c, idx) => (
                        <td key={idx} className="py-3.5 px-4 align-top text-text-main font-extrabold text-sm leading-relaxed">
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
        title="Save Target Cloud Model"
        stageKind="target"
        defaultName={
          selectedCloudModel && !isBuiltInSampleModel(selectedCloudModel) && !String(selectedCloudModel.id).startsWith('candidate-')
            ? cleanModelDisplayName(selectedCloudModel.name)
            : buildDefaultModelName('target', selectedSourceModel?.name || editedCandidate?.targetInfra?.name, editedCandidate?.targetCloud?.csp)
        }
        defaultDescription={
          selectedCloudModel && !isBuiltInSampleModel(selectedCloudModel) && !String(selectedCloudModel.id).startsWith('candidate-')
            ? selectedCloudModel.description || ''
            : `Optimized target cloud model for ${editedCandidate?.targetCloud?.csp?.toUpperCase() || 'Cloud'} (${editedCandidate?.targetCloud?.region || ''})`
        }
        defaultVersion={
          selectedCloudModel && !isBuiltInSampleModel(selectedCloudModel) && !String(selectedCloudModel.id).startsWith('candidate-')
            ? selectedCloudModel.version || '1.0.0'
            : '1.0.0'
        }
        defaultOverwriteId={
          selectedCloudModel && !isBuiltInSampleModel(selectedCloudModel) && !String(selectedCloudModel.id).startsWith('candidate-')
            ? selectedCloudModel.id
            : null
        }
        existingRevisions={userTargetModels.map(m => ({ id: m.id, name: m.name, description: m.description, version: m.version }))}
        onSave={handleSaveArtifact}
      />

      <DeleteModelConfirmModal
        isOpen={showDeleteConfirm && Boolean(pendingCloudModel)}
        onClose={() => setShowDeleteConfirm(false)}
        title="Delete Target Cloud Model"
        modelName={pendingCloudModel?.name || ''}
        onConfirm={async () => {
          if (!pendingCloudModel) return;
          const deletedName = cleanModelDisplayName(pendingCloudModel.name);
          const deletedId = pendingCloudModel.id;
          await deleteCloudModel(deletedId);
          const remaining = validSavedCloudModels.filter(m => m.id !== deletedId);
          setPendingCloudModelId(remaining[0]?.id || '');
          if (loadedCloudModelId === deletedId || selectedCloudModel?.id === deletedId) {
            setLoadedCloudModelId('');
            setIsModelLoaded(false);
            setHasUnsavedChanges(false);
            setActiveStep(1);
          }
          setToastTone('success');
          setToastMsg(`Target cloud model "${deletedName}" deleted.`);
        }}
      />

      <DeleteModelConfirmModal
        isOpen={showClearAllConfirm}
        onClose={() => setShowClearAllConfirm(false)}
        title="Delete All Custom Models"
        modelName="DELETE"
        descriptionText={`Are you sure you want to delete all ${userTargetModels.length} custom target cloud model(s)? Built-in sample models will be preserved.`}
        onConfirm={async () => {
          const count = await clearCustomCloudModels();
          setPendingCloudModelId(sampleTargetModels[0]?.id || '');
          setActiveStep(1);
          setHasUnsavedChanges(false);
          setToastTone('success');
          setToastMsg(`Deleted ${count} custom target cloud model(s).`);
        }}
      />

      <UnsavedConfirmModal
        isOpen={showUnsavedNextModal}
        onStay={() => setShowUnsavedNextModal(false)}
        onProceed={executeProceedToMigration}
      />

      <ModelToast
        message={toastMsg}
        tone={toastTone}
        onClose={() => setToastMsg(null)}
      />

      {/* Pre-Flight Validation Modal */}
      {dryRunResult.open && (
        <div className="fixed inset-0 z-[60] flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-2xl border border-border-main animate-scale-up space-y-4">
            <div className="flex justify-between items-center border-b border-border-main/20 pb-3">
              <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
                <ShieldCheck className="w-5 h-5 text-emerald-500" />
                Pre-Flight Infrastructure Validation (Dry-Run)
              </h3>
              <button
                onClick={() => setDryRunResult({ open: false, issues: [] })}
                className="text-text-muted hover:text-text-main transition p-1 hover:bg-bg-input rounded-lg cursor-pointer"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            <p className="text-xs text-text-muted leading-relaxed">
              CM-Beetle dry-run validation checks 12 core consistency rules (referential integrity, spec-image architecture matching, CSP connection names, and CIDR subnet collisions) before initiating cloud provisioning.
            </p>

            {dryRunResult.error ? (
              <div className="p-4 bg-red-500/10 border border-red-500/30 rounded-xl text-xs space-y-1">
                <div className="flex items-center gap-1.5 font-bold text-red-600 dark:text-red-400">
                  <AlertTriangle className="w-4 h-4 text-red-500" />
                  <span>Validation Request Error</span>
                </div>
                <p className="text-text-muted">{dryRunResult.error}</p>
              </div>
            ) : dryRunResult.valid ? (
              <div className="p-4 bg-emerald-500/10 border border-emerald-500/30 rounded-xl space-y-2">
                <div className="flex items-center gap-2 font-extrabold text-emerald-600 dark:text-emerald-400 text-sm">
                  <CheckCircle2 className="w-5 h-5 text-emerald-500 shrink-0" />
                  <span>All 12 Pre-Flight Consistency Checks Passed!</span>
                </div>
                <p className="text-xs text-text-muted leading-relaxed">
                  Target cloud specifications, OS image configurations, VNet/Subnet address spaces, and security rules are verified and ready for deployment.
                </p>
                {dryRunResult.issues && dryRunResult.issues.length > 0 && (
                  <div className="pt-2 border-t border-emerald-500/20 space-y-1 font-mono text-xs">
                    <span className="text-amber-500 font-bold">Advisory Warnings ({dryRunResult.issues.length}):</span>
                    {dryRunResult.issues.map((iss, i) => (
                      <div key={i} className="text-amber-400 flex items-start gap-1">
                        <span>• [{iss.code}]</span>
                        <span>{iss.message}</span>
                      </div>
                    ))}
                  </div>
                )}
              </div>
            ) : (
              <div className="p-4 bg-red-500/10 border border-red-500/30 rounded-xl space-y-3">
                <div className="flex items-center gap-2 font-extrabold text-red-600 dark:text-red-400 text-sm">
                  <AlertTriangle className="w-5 h-5 text-red-500 shrink-0" />
                  <span>Validation Issues Detected ({dryRunResult.issues.length} issue(s))</span>
                </div>
                <div className="max-h-60 overflow-y-auto space-y-2 pr-1 font-mono text-xs">
                  {dryRunResult.issues.map((iss, i) => (
                    <div key={i} className={`p-3 rounded-lg border space-y-1 ${
                      iss.severity === 'error'
                        ? 'bg-red-500/15 border-red-500/30 text-red-400'
                        : 'bg-amber-500/15 border-amber-500/30 text-amber-400'
                    }`}>
                      <div className="flex items-center justify-between font-bold">
                        <span>[{iss.code}]</span>
                        <span className="uppercase text-xs px-1.5 py-0.5 rounded font-sans bg-bg-panel text-text-main font-extrabold">{iss.severity}</span>
                      </div>
                      {iss.path && <div className="text-xs text-text-muted">Path: {iss.path}</div>}
                      <div className="font-sans font-normal text-text-main">{iss.message}</div>
                    </div>
                  ))}
                </div>
              </div>
            )}

            <div className="flex justify-end pt-2">
              <button
                type="button"
                onClick={() => setDryRunResult({ open: false, issues: [] })}
                className="px-5 py-2 bg-bg-panel border border-border-main hover:bg-bg-input text-text-main rounded-xl text-xs font-bold transition cursor-pointer"
              >
                Close
              </button>
            </div>
          </div>
        </div>
      )}

    </div>
  );
};
