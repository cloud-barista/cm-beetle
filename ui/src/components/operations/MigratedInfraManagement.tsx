'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '../../store/migrationStore';
import { TopologyMap } from '../design/TopologyMap';
import { beetleApi, artifactApi } from '../../api/client';
import { resolveTargetNodeLineage, parseDiscreteTcpPorts } from '../../utils/nodeLineage';
import { isBuiltInSampleModel } from '../../utils/modelFormat';
import { CheckSSHReadyResponse } from '../../types/migration';
import { ObjectDirectoryTree } from '../storage/ObjectDirectoryTree';
import { 
  Server, 
  Trash2, 
  RefreshCw, 
  CheckCircle2, 
  AlertTriangle, 
  Eye, 
  Download, 
  Copy, 
  Check, 
  Cloud, 
  Layers, 
  Key, 
  Globe, 
  HardDrive, 
  Sparkles,
  Loader2,
  X,
  LayoutGrid,
  List,
  Shield,
  Network,
  Radio,
  Database,
  Sliders,
  Zap,
  Activity,
  FileText,
  ExternalLink,
  GitBranch,
  ArrowRight,
  Compass
} from 'lucide-react';

const SAMPLE_INFRA_ID = 'sample-aws-infra-01';

const SAMPLE_STORAGE_ID = 'sample-aws-object-storage-01';

const SAMPLE_STORAGE_DETAIL = {
  resourceType: 'ObjectStorage',
  id: 'sample-aws-object-storage-01',
  uid: 'sample-aws-object-storage-01-uid-001',
  name: 'sample-aws-object-storage-01',
  connectionName: 'aws-ap-northeast-2',
  connectionConfig: {
    configName: 'aws-ap-northeast-2',
    providerName: 'AWS',
    driverName: 'aws-driver-v1.0',
    credentialName: 'aws-credential',
    credentialHolder: 'admin',
    regionZoneInfo: {
      assignedRegion: 'ap-northeast-2',
      assignedZone: 'ap-northeast-2a'
    }
  },
  description: 'Sample object storage for UI demo',
  status: 'Available',
  creationDate: '2026-07-23T10:00:00Z',
  contents: [
    { key: 'sample-data-01.csv', size: 1024500, storageClass: 'STANDARD', eTag: '9b2cf535f27731c974343645a3985328', lastModified: '2026-07-23T10:05:00Z' },
    { key: 'backup-config.json', size: 45200, storageClass: 'STANDARD', eTag: '4a1bf535f27731c974343645a3989912', lastModified: '2026-07-23T10:10:00Z' }
  ]
};

const MIGRATED_STORAGE_DEMO = {
  resourceType: 'ObjectStorage',
  id: 'os101-x8f2-os101-x8f2',
  uid: 'os101-x8f2-os101-x8f2',
  name: 'os101-x8f2-os101-x8f2',
  connectionName: 'aws-ap-northeast-2',
  connectionConfig: {
    configName: 'aws-ap-northeast-2',
    providerName: 'AWS',
    driverName: 'aws-driver-v1.0',
    credentialName: 'aws-credential',
    credentialHolder: 'admin',
    regionZoneInfo: {
      assignedRegion: 'ap-northeast-2',
      assignedZone: 'ap-northeast-2a'
    }
  },
  description: 'Created by CM-Beetle',
  status: 'Available',
  creationDate: '2026-07-23T12:00:00Z',
  contents: []
};

const extractCspAndRegionFromConnection = (connName?: string) => {
  if (!connName) return { csp: 'AWS', region: 'ap-northeast-2' };
  const parts = connName.split('-');
  if (parts.length >= 2) {
    const csp = parts[0].toUpperCase();
    const region = parts.slice(1).join('-');
    return { csp, region };
  }
  return { csp: connName.toUpperCase(), region: 'ap-northeast-2' };
};

const SAMPLE_INFRA_DETAIL = {
  name: 'sample-aws-infra-01',
  description: 'Sample AWS Multi-Cloud Migrated Infrastructure (Demo)',
  targetCloud: {
    csp: 'aws',
    region: 'ap-northeast-2'
  },
  targetVNet: {
    name: 'sample-vnet-aws-01',
    cidrBlock: '10.0.0.0/16',
    subnetInfoList: [
      { name: 'sample-subnet-aws-01', cidrBlock: '10.0.1.0/24' },
      { name: 'sample-subnet-aws-02', cidrBlock: '10.0.2.0/24' }
    ]
  },
  targetSecurityGroupList: [
    {
      name: 'sample-sg-aws-01',
      firewallRules: [
        { protocol: 'tcp', dstPorts: '22', cidr: '0.0.0.0/0', direction: 'inbound' },
        { protocol: 'tcp', dstPorts: '80, 443', cidr: '0.0.0.0/0', direction: 'inbound' },
        { protocol: 'tcp', dstPorts: '8080, 8056', cidr: '10.0.0.0/16', direction: 'inbound' }
      ]
    }
  ],
  targetSshKey: {
    name: 'sample-key-aws-01',
    fingerprint: 'SHA256:x9A1b2C3d4E5f6G7h8I9j0K1l2M3n4O5p6Q7r8S9t0U'
  },
  targetNlbList: [
    {
      name: 'sample-nlb-aws-01',
      listenerPort: '80',
      targetPort: '8080',
      protocol: 'TCP',
      healthCheck: 'HTTP /healthz'
    }
  ],
  vms: [
    { name: 'sample-web-node-01', specId: 'aws+ap-northeast-2+t3.small', publicIp: '54.180.12.101', privateIp: '10.0.1.15' },
    { name: 'sample-app-node-01', specId: 'aws+ap-northeast-2+t3.medium', publicIp: '54.180.12.102', privateIp: '10.0.1.25' }
  ]
};

const isGpuSpec = (specId: string, specList?: any[]): boolean => {
  if (!specId) return false;
  const s = specId.toLowerCase();
  if (Array.isArray(specList)) {
    const matched = specList.find(
      (sp: any) =>
        sp?.id?.toLowerCase() === s ||
        sp?.cspSpecName?.toLowerCase() === s ||
        (sp?.cspSpecName && s.endsWith(String(sp.cspSpecName).toLowerCase()))
    );
    if (matched) {
      return Number(matched.acceleratorCount ?? matched.AcceleratorCount ?? 0) > 0 || Boolean(matched.acceleratorModel);
    }
  }
  return (
    s.includes('gpu') ||
    s.includes('g4dn.') ||
    s.includes('g5.') ||
    s.includes('g5g.') ||
    s.includes('g6.') ||
    s.includes('g2-') ||
    s.includes('p3.') ||
    s.includes('p4d') ||
    s.includes('p5.') ||
    s.includes('a2-') ||
    s.includes('a3-') ||
    s.includes('standard_nc') ||
    s.includes('standard_nd') ||
    s.includes('standard_nv') ||
    s.includes('tesla') ||
    s.includes('a100') ||
    s.includes('h100')
  );
};

const isValidEndpointHost = (host?: string): boolean => {
  if (!host || host === 'N/A') return false;
  return /^[a-zA-Z0-9.-]+$/.test(host.trim());
};

export const MigratedInfraManagement: React.FC = () => {
  const {
    jobs,
    objectStorageJobs,
    savedCloudModels,
    selectedCloudModel,
    editedCandidate,
    recommendationCandidates,
    selectedCandidateIndex,
    savedSourceModels,
    selectedSourceModel,
    refinedSourceInfra,
    fetchSavedSourceModels,
    selectSourceModel,
    fetchSavedCloudModels,
    selectCloudModel,
    setRecommendationMode,
    setActiveTab,
    namespaceId,
    liveReportHtml,
    fetchMigrationReport,
    deletingInfrasMap,
    startInfraTeardown,
    removeInfraTeardown,
    pollInfraTeardownStatus,
    deletingStoragesMap,
    startStorageTeardown,
    removeStorageTeardown,
    pollStorageTeardownStatus,
    removeJob,
    targetOverviewInfraId,
    setTargetOverviewInfraId
  } = useMigrationStore();

  const [copiedIp, setCopiedIp] = useState<string | null>(null);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [deleteConfirmText, setDeleteConfirmText] = useState('');
  const [isDeleting, setIsDeleting] = useState(false);
  const [deleteError, setDeleteError] = useState('');
  const [deleteSuccessMsg, setDeleteSuccessMsg] = useState('');
  const [detectedNlbs, setDetectedNlbs] = useState<any[]>([]);
  const [deleteStage, setDeleteStage] = useState<'idle' | 'deleting-nlb' | 'deleting-infra'>('idle');
  const [liveNlbs, setLiveNlbs] = useState<any[]>([]);
  const [hasFetchedLiveNlbs, setHasFetchedLiveNlbs] = useState<boolean>(false);
  const [liveSecurityGroups, setLiveSecurityGroups] = useState<any[]>([]);
  const [deletingNlbId, setDeletingNlbId] = useState<string | null>(null);

  // Migrated Infra IDs fetched from GET /beetle/migration/ns/{nsId}/infra?option=id
  const [migratedInfraIds, setMigratedInfraIds] = useState<string[]>([]);
  const [isLoadingIds, setIsLoadingIds] = useState(false);
  const [selectedInfraId, setSelectedInfraId] = useState<string>('');
  const [catalogViewMode, setCatalogViewMode] = useState<'grid' | 'table'>('grid');
  const [activeSubTab, setActiveSubTab] = useState<'vms' | 'vnets' | 'sgs' | 'ssh' | 'nlb'>('vms');
  const [overviewCategory, setOverviewCategory] = useState<'infra' | 'storage'>('infra');

  const [migratedStorages, setMigratedStorages] = useState<any[]>([]);
  const [isLoadingStorages, setIsLoadingStorages] = useState<boolean>(false);
  const [storageCatalogViewMode, setStorageCatalogViewMode] = useState<'grid' | 'table'>('grid');
  const [selectedStorageId, setSelectedStorageId] = useState<string>('');
  const [showStorageDeleteConfirm, setShowStorageDeleteConfirm] = useState(false);
  const [targetStorageToDelete, setTargetStorageToDelete] = useState<string>('');
  const [storageDeleteConfirmText, setStorageDeleteConfirmText] = useState('');
  const [forceDeleteStorage, setForceDeleteStorage] = useState(false);
  const [preferAsyncStorageDelete, setPreferAsyncStorageDelete] = useState(true);
  const [isDeletingStorage, setIsDeletingStorage] = useState(false);
  const [storageDeleteError, setStorageDeleteError] = useState('');

  const handleOpenStorageDeleteModal = (sName: string) => {
    if (sName === SAMPLE_STORAGE_ID) {
      setStorageDeleteError('Sample object storage is protected and cannot be deleted.');
      return;
    }
    setTargetStorageToDelete(sName);
    setStorageDeleteConfirmText('');
    setForceDeleteStorage(false);
    setPreferAsyncStorageDelete(true);
    setStorageDeleteError('');
    setShowStorageDeleteConfirm(true);
  };

  const handleConfirmStorageDelete = async () => {
    if (!targetStorageToDelete || storageDeleteConfirmText !== targetStorageToDelete) return;
    setIsDeletingStorage(true);
    setStorageDeleteError('');

    try {
      const ns = namespaceId || 'mig01';
      const option = forceDeleteStorage ? 'force' : '';
      const res = await beetleApi.deleteMigratedObjectStorage(ns, targetStorageToDelete, option, preferAsyncStorageDelete);
      if (res.success) {
        setShowStorageDeleteConfirm(false);
        const reqId = res.reqId || `req-del-storage-${Date.now()}`;
        startStorageTeardown(targetStorageToDelete, reqId);

        if (preferAsyncStorageDelete) {
          setDeleteSuccessMsg(`⚡ Object storage "${targetStorageToDelete}" deletion task initiated asynchronously (ReqID: ${reqId}).`);
        } else {
          setDeleteSuccessMsg(`Object storage "${targetStorageToDelete}" deleted successfully.`);
        }
        setTimeout(() => setDeleteSuccessMsg(''), 6000);
        await loadMigratedStorages();
      } else {
        setStorageDeleteError(res.error || 'Failed to delete object storage.');
      }
    } catch (err: any) {
      setStorageDeleteError(err.message || 'Failed to delete object storage.');
    } finally {
      setIsDeletingStorage(false);
    }
  };

  const loadMigratedStorages = async () => {
    setIsLoadingStorages(true);
    try {
      const ns = namespaceId || 'mig01';
      const list = await beetleApi.getMigratedObjectStorages(ns);
      const storageList = Array.isArray(list) ? list : [];
      
      const map = new Map<string, any>();
      map.set(SAMPLE_STORAGE_ID, SAMPLE_STORAGE_DETAIL);

      storageList.forEach((s: any) => {
        const key = s.id || s.name || s.bucketName;
        if (key) map.set(key, s);
      });

      const combined = Array.from(map.values());
      setMigratedStorages(combined);

      // Auto-select first real storage if available, otherwise sample
      const realItem = combined.find((s: any) => s.id !== SAMPLE_STORAGE_ID);
      if (realItem) {
        setSelectedStorageId(realItem.id || realItem.name || realItem.bucketName);
      } else {
        setSelectedStorageId(SAMPLE_STORAGE_ID);
      }
    } catch (err) {
      console.warn('Failed to fetch migrated object storages:', err);
      setMigratedStorages([SAMPLE_STORAGE_DETAIL]);
      setSelectedStorageId(SAMPLE_STORAGE_ID);
    } finally {
      setIsLoadingStorages(false);
    }
  };

  // Detailed infra object fetched from GET /beetle/migration/ns/{nsId}/infra/{infraId}
  const [loadedInfraDetail, setLoadedInfraDetail] = useState<any | null>(null);
  const [isLoadingDetail, setIsLoadingDetail] = useState(false);

  // Detailed storage object fetched from GET /beetle/migration/middleware/ns/{nsId}/objectStorage/{osId}
  const [loadedStorageDetail, setLoadedStorageDetail] = useState<any | null>(null);
  const [isLoadingStorageDetail, setIsLoadingStorageDetail] = useState<boolean>(false);

  useEffect(() => {
    if (!selectedStorageId || selectedStorageId === SAMPLE_STORAGE_ID) {
      setLoadedStorageDetail(null);
      return;
    }
    const fetchDetail = async () => {
      setIsLoadingStorageDetail(true);
      try {
        const ns = namespaceId || 'mig01';
        const detail = await beetleApi.getMigratedObjectStorageDetail(ns, selectedStorageId);
        if (detail) {
          setLoadedStorageDetail(detail);
        }
      } catch (err) {
        console.warn('Failed to fetch storage detail:', err);
      } finally {
        setIsLoadingStorageDetail(false);
      }
    };
    fetchDetail();
  }, [selectedStorageId, namespaceId]);

  const fetchMigratedInfraIds = async () => {
    setIsLoadingIds(true);
    try {
      const ns = namespaceId || 'mig01';
      const ids = await beetleApi.getMigratedInfraIdList(ns);
      // Only include jobs that are actively provisioning (in-flight)
      const inFlightJobInfraIds = jobs
        .filter(j => j.status === 'Handling' || j.nlbStatus === 'Provisioning')
        .map(j => j.infraId)
        .filter(Boolean);
      const combinedIds = Array.from(new Set([SAMPLE_INFRA_ID, ...inFlightJobInfraIds, ...ids]));
      setMigratedInfraIds(combinedIds);

      // Auto-select valid ID if current selection is invalid or empty
      setSelectedInfraId(prev => (prev && combinedIds.includes(prev) ? prev : combinedIds[0] || ''));
    } catch (err) {
      console.warn('Failed to load migrated infra IDs, falling back to sample:', err);
      const inFlightJobInfraIds = jobs
        .filter(j => j.status === 'Handling' || j.nlbStatus === 'Provisioning')
        .map(j => j.infraId)
        .filter(Boolean);
      const fallbackIds = Array.from(new Set([SAMPLE_INFRA_ID, ...inFlightJobInfraIds]));
      setMigratedInfraIds(fallbackIds);
      setSelectedInfraId(prev => (prev && fallbackIds.includes(prev) ? prev : fallbackIds[0] || ''));
    } finally {
      setIsLoadingIds(false);
    }
  };

  // Persistent teardown polling effect (runs across tab navigation and triggers catalog refresh on completion!)
  useEffect(() => {
    const activeDeletingEntries = Object.entries(deletingInfrasMap);
    if (activeDeletingEntries.length === 0) return;

    const intervalId = setInterval(async () => {
      const ns = namespaceId || 'mig01';
      for (const [infraId, info] of Object.entries(deletingInfrasMap)) {
        if (!info.reqId) continue;
        const res = await pollInfraTeardownStatus(ns, infraId, info.reqId);
        if (res.completed) {
          if (res.success) {
            setDeleteSuccessMsg(`Target infrastructure "${infraId}" was completely terminated and removed.`);
            // Clean up any lingering job for this infra in store
            jobs.filter(j => j.infraId === infraId).forEach(j => removeJob(j.id));
          } else {
            if (res.error?.includes('does not exist') || res.error?.includes('not found')) {
              setDeleteSuccessMsg(`Target infrastructure "${infraId}" does not exist or was already deleted.`);
              jobs.filter(j => j.infraId === infraId).forEach(j => removeJob(j.id));
            } else {
              setDeleteError(`Infrastructure "${infraId}" deletion failed: ${res.error || 'Unknown error'}`);
            }
          }
          // Refresh list from server after teardown completes!
          await fetchMigratedInfraIds();
          if (selectedInfraId === infraId) {
            setSelectedInfraId(SAMPLE_INFRA_ID);
            setLoadedInfraDetail(SAMPLE_INFRA_DETAIL);
          }
        }
      }
    }, 3000);

    return () => clearInterval(intervalId);
  }, [deletingInfrasMap, namespaceId, selectedInfraId, jobs, removeJob]);

  // Persistent Object Storage teardown polling effect (runs across tab navigation!)
  useEffect(() => {
    const activeDeletingEntries = Object.entries(deletingStoragesMap);
    if (activeDeletingEntries.length === 0) return;

    const intervalId = setInterval(async () => {
      const ns = namespaceId || 'mig01';
      for (const [storageId, info] of Object.entries(deletingStoragesMap)) {
        if (!info.reqId) continue;
        const res = await pollStorageTeardownStatus(ns, storageId, info.reqId);
        if (res.completed) {
          if (res.success) {
            setDeleteSuccessMsg(`Object storage "${storageId}" was completely deleted.`);
          } else {
            setDeleteError(`Object storage "${storageId}" deletion failed: ${res.error || 'Unknown error'}`);
          }
          await loadMigratedStorages();
          if (selectedStorageId === storageId) {
            setSelectedStorageId('');
          }
        }
      }
    }, 3000);

    return () => clearInterval(intervalId);
  }, [deletingStoragesMap, namespaceId, selectedStorageId]);

  const [infraReportHtml, setInfraReportHtml] = useState<string>('');
  const [reportDocType, setReportDocType] = useState<'migration' | 'target-summary'>('migration');
  const [reportFormat, setReportFormat] = useState<'html' | 'md' | 'json'>('html');
  const [reportContent, setReportContent] = useState<string>('');
  const [isLoadingReport, setIsLoadingReport] = useState<boolean>(false);

  // Official CM-Beetle SSH readiness state (GET /migration/ns/{nsId}/infra/{infraId}/ssh-ready?option=detail)
  const [sshCooldownSec, setSshCooldownSec] = useState<number>(0);
  const [sshReadyData, setSshReadyData] = useState<CheckSSHReadyResponse | null>(null);
  const [isCheckingSsh, setIsCheckingSsh] = useState<boolean>(false);
  const [sshError, setSshError] = useState<string>('');
  const [showSshModal, setShowSshModal] = useState<boolean>(false);

  useEffect(() => {
    if (sshCooldownSec <= 0) return;
    const timer = setInterval(() => {
      setSshCooldownSec((prev) => (prev <= 1 ? 0 : prev - 1));
    }, 1000);
    return () => clearInterval(timer);
  }, [sshCooldownSec]);

  const handleCheckInfraSshReady = async () => {
    if (!selectedInfraId || selectedInfraId === SAMPLE_INFRA_ID) return;
    const ns = namespaceId || 'mig01';
    setShowSshModal(true);
    if (sshCooldownSec > 0) return;

    setIsCheckingSsh(true);
    setSshError('');
    try {
      const res = await beetleApi.checkInfraSshReady(ns, selectedInfraId, 'detail');
      if (res.retryAfterSeconds && res.retryAfterSeconds > 0) {
        setSshCooldownSec(res.retryAfterSeconds);
      }
      if (res.data) {
        setSshReadyData(res.data);
      }
      if (!res.success && res.error) {
        setSshError(res.error);
      }
    } catch (err: any) {
      setSshError(err.message || 'Failed to check SSH readiness.');
    } finally {
      setIsCheckingSsh(false);
    }
  };

  const fetchReportOrSummary = async (
    targetInfraId: string,
    docType: 'migration' | 'target-summary' = reportDocType,
    fmt: 'html' | 'md' | 'json' = reportFormat
  ) => {
    if (!targetInfraId || targetInfraId === SAMPLE_INFRA_ID) return;
    setIsLoadingReport(true);
    try {
      const ns = namespaceId || 'mig01';
      if (docType === 'migration') {
        const srcPayload = selectedSourceModel?.onpremiseInfraModel || refinedSourceInfra;
        const content = await beetleApi.getMigrationReport(ns, targetInfraId, srcPayload, fmt);
        setReportContent(content);
        if (fmt === 'html') {
          setInfraReportHtml(content);
        }
      } else {
        const content = await beetleApi.getTargetInfraSummary(ns, targetInfraId, fmt);
        setReportContent(content);
        if (fmt === 'html') {
          setInfraReportHtml(content);
        }
      }
    } catch (rErr) {
      console.warn('Report or summary fetch notice:', rErr);
    } finally {
      setIsLoadingReport(false);
    }
  };

  const handleLoadInfraDetail = async (targetId?: string) => {
    const infraIdToLoad = targetId || selectedInfraId;
    if (!infraIdToLoad) return;
    setSelectedInfraId(infraIdToLoad);
    setSshReadyData(null);
    setSshError('');
    setIsLoadingDetail(true);
    try {
      if (infraIdToLoad === SAMPLE_INFRA_ID) {
        setLoadedInfraDetail(SAMPLE_INFRA_DETAIL);
        setLiveNlbs([]);
        setHasFetchedLiveNlbs(false);
        setLiveSecurityGroups([]);
        setIsLoadingDetail(false);
        return;
      }
      const ns = namespaceId || 'mig01';
      const detail = await beetleApi.getMigratedInfraDetail(ns, infraIdToLoad);
      if (!detail) {
        setLoadedInfraDetail(null);
        setLiveSecurityGroups([]);
        return;
      }
      setLoadedInfraDetail(detail);
      setLiveNlbs([]);
      setHasFetchedLiveNlbs(false);

      try {
        const summaryJson = await beetleApi.getTargetInfraSummary(ns, infraIdToLoad, 'json');
        const parsedSummary = typeof summaryJson === 'string' ? JSON.parse(summaryJson) : summaryJson;
        const sgs = parsedSummary?.securityResources?.securityGroups;
        if (Array.isArray(sgs) && sgs.length > 0) {
          setLiveSecurityGroups(sgs);
        } else {
          setLiveSecurityGroups([]);
        }
      } catch {
        setLiveSecurityGroups([]);
      }

      // Fetch active NLB list and compliance report in parallel
      await Promise.all([
        (async () => {
          try {
            const nRes = await beetleApi.getNlbList(ns, infraIdToLoad);
            if (nRes.success && Array.isArray(nRes.data)) {
              setLiveNlbs(nRes.data);
              setHasFetchedLiveNlbs(true);
            } else {
              setLiveNlbs([]);
              setHasFetchedLiveNlbs(true);
            }
          } catch (nErr) {
            console.warn('NLB list fetch notice:', nErr);
            setLiveNlbs([]);
            setHasFetchedLiveNlbs(true);
          }
        })(),
        fetchReportOrSummary(infraIdToLoad, reportDocType, reportFormat)
      ]);
    } catch (err) {
      console.warn('Failed to load infra detail for', infraIdToLoad, err);
      if (infraIdToLoad === SAMPLE_INFRA_ID) {
        setLoadedInfraDetail(SAMPLE_INFRA_DETAIL);
        setLiveNlbs([]);
        setHasFetchedLiveNlbs(false);
      }
    } finally {
      setIsLoadingDetail(false);
    }
  };

  const handleOpenDeleteModalForInfra = async (infraId: string) => {
    if (infraId === SAMPLE_INFRA_ID) {
      alert('Sample infrastructure is protected and cannot be deleted.');
      return;
    }
    setSelectedInfraId(infraId);
    setShowDeleteConfirm(true);
    setDeleteConfirmText('');
    setDeleteError('');
    setDeleteStage('idle');

    // Query active NLBs for staged teardown
    try {
      const ns = namespaceId || 'mig01';
      const nRes = await beetleApi.getNlbList(ns, infraId);
      if (nRes.success && Array.isArray(nRes.data) && nRes.data.length > 0) {
        setDetectedNlbs(nRes.data);
      } else {
        const matchModel = savedCloudModels.find(m => m.id === infraId || m.name === infraId);
        const modelNlbs = (matchModel?.cloudInfraModel as any)?.targetNlbList || [];
        setDetectedNlbs(modelNlbs);
      }
    } catch {
      setDetectedNlbs([]);
    }
  };

  const handleDeleteSingleNlb = async (nlbId: string) => {
    if (!selectedInfraId) return;
    const confirm = window.confirm(`Are you sure you want to delete Managed NLB "${nlbId}"?\n\n(Compute nodes and network resources will remain intact)`);
    if (!confirm) return;

    setDeletingNlbId(nlbId);
    try {
      const ns = namespaceId || 'mig01';
      const res = await beetleApi.deleteNlb(ns, selectedInfraId, nlbId);
      if (res.success) {
        setLiveNlbs(prev => prev.filter(n => (n.id || n.name) !== nlbId));
        setHasFetchedLiveNlbs(true);
        setDeleteSuccessMsg(`✓ Managed NLB "${nlbId}" was deleted successfully.`);
      } else {
        alert(`Failed to delete NLB: ${res.error || 'Unknown error'}`);
      }
    } catch (err: any) {
      alert(`Failed to delete NLB: ${err.message || 'Network error'}`);
    } finally {
      setDeletingNlbId(null);
    }
  };

  const completedJobsCount = jobs.filter(j => j.status === 'Success').length;
  const completedStoragesCount = objectStorageJobs.filter(j => j.status === 'Success').length;
  const activeJobsCount = jobs.filter(j => j.status === 'Handling' || j.nlbStatus === 'Provisioning').length;

  useEffect(() => {
    fetchMigratedInfraIds();
    loadMigratedStorages();
  }, [namespaceId, completedJobsCount, completedStoragesCount, activeJobsCount, jobs.length]);

  // Automatically focus and load the target infra when navigated from 4. Migration Execution
  useEffect(() => {
    if (targetOverviewInfraId) {
      setOverviewCategory('infra');
      setSelectedInfraId(targetOverviewInfraId);
      handleLoadInfraDetail(targetOverviewInfraId);
      setTargetOverviewInfraId('');
    }
  }, [targetOverviewInfraId]);

  // Combine backend migrated infras, completed migration jobs from store, and saved models
  const completedJobs = jobs.filter(j => j.status === 'Success');

  const handleCopySshCommand = (ip: string, keyName: string) => {
    if (!ip || ip === 'N/A') return;
    const cmd = `ssh -i ~/.ssh/${keyName}.pem cb-user@${ip}`;
    navigator.clipboard.writeText(cmd);
    setCopiedIp(ip);
    setTimeout(() => setCopiedIp(null), 2000);
  };

  const handleDownloadReport = () => {
    const activeText = reportContent || infraReportHtml || liveReportHtml;
    if (!activeText) return;
    const mimeType =
      reportFormat === 'json'
        ? 'application/json'
        : reportFormat === 'md'
        ? 'text/markdown'
        : 'text/html';
    const ext = reportFormat;
    const prefix = reportDocType === 'target-summary' ? 'target_summary' : 'migration_report';
    const blob = new Blob([activeText], { type: mimeType });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `${prefix}_${selectedInfraId || 'infra'}.${ext}`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  };

  const handleConfirmDelete = async () => {
    if (!selectedInfraId) return;
    if (selectedInfraId === SAMPLE_INFRA_ID) {
      setDeleteError('Sample infrastructure is protected and cannot be deleted.');
      setShowDeleteConfirm(false);
      return;
    }
    const targetInfraId = selectedInfraId;
    setIsDeleting(true);
    setDeleteError('');
    setDeleteSuccessMsg('');

    const ns = namespaceId || 'mig01';

    try {
      // Step 1: If an NLB exists, delete the NLB first via flat API
      if (detectedNlbs.length > 0) {
        setDeleteStage('deleting-nlb');
        for (const nlb of detectedNlbs) {
          const nlbId = nlb.id || nlb.name;
          if (nlbId) {
            const delRes = await beetleApi.deleteNlb(ns, targetInfraId, nlbId);
            if (!delRes.success) {
              console.warn(`NLB ${nlbId} delete warning:`, delRes.error);
            }
          }
        }
      }

      // Step 2: Delete base infrastructure (MCI/VMs, SGs, VNets) via flat API
      setDeleteStage('deleting-infra');
      const res = await beetleApi.deleteMigratedInfra(ns, targetInfraId, 'terminate', true);

      if (res.success) {
        setShowDeleteConfirm(false);
        setDeleteConfirmText('');
        setDeleteStage('idle');

        if (res.reqId) {
          startInfraTeardown(targetInfraId, res.reqId);
          setDeleteSuccessMsg(
            detectedNlbs.length > 0
              ? `Target infrastructure "${targetInfraId}" termination initiated: Step 1 (NLB detached) completed, Step 2 (Infra teardown Req ID: ${res.reqId}) is in progress.`
              : `Target infrastructure "${targetInfraId}" termination initiated asynchronously (Req ID: ${res.reqId}).`
          );
        } else {
          setDeleteSuccessMsg(`Target infrastructure "${targetInfraId}" deleted successfully.`);
          jobs.filter(j => j.infraId === targetInfraId).forEach(j => removeJob(j.id));
          await fetchMigratedInfraIds();
          if (selectedInfraId === targetInfraId) {
            setSelectedInfraId(SAMPLE_INFRA_ID);
            setLoadedInfraDetail(SAMPLE_INFRA_DETAIL);
          }
        }
      } else {
        if (res.error?.includes('does not exist') || res.error?.includes('not found')) {
          setShowDeleteConfirm(false);
          setDeleteConfirmText('');
          setDeleteStage('idle');
          setDeleteSuccessMsg(`Target infrastructure "${targetInfraId}" does not exist or was already deleted.`);
          jobs.filter(j => j.infraId === targetInfraId).forEach(j => removeJob(j.id));
          await fetchMigratedInfraIds();
          setSelectedInfraId(SAMPLE_INFRA_ID);
          setLoadedInfraDetail(SAMPLE_INFRA_DETAIL);
        } else {
          setDeleteError(res.error || 'Failed to delete target infrastructure.');
        }
      }
    } catch (err: any) {
      console.error('Delete failed:', err);
      setDeleteError(err.response?.data?.error || err.message || 'Failed to delete target infrastructure.');
    } finally {
      setIsDeleting(false);
      setDeleteStage('idle');
    }
  };

  // Active infra detail from loadedInfraDetail, completed session jobs, or savedCloudModels
  const activeJobInfra = completedJobs.find(j => j.id === selectedInfraId || j.infraId === selectedInfraId);
  const targetModelIdLabel = loadedInfraDetail?.label?.['cm-target-model-id'];
  const candidateSavedModels =
    selectedInfraId === SAMPLE_INFRA_ID
      ? savedCloudModels
      : [...savedCloudModels].reverse().filter((m) => !isBuiltInSampleModel(m));
  const activeSavedModel = candidateSavedModels.find(
    (m) =>
      (targetModelIdLabel && m.id === targetModelIdLabel) ||
      m.id === selectedInfraId ||
      m.name === selectedInfraId ||
      m.cloudInfraModel?.targetInfra?.name === selectedInfraId
  );

  // Resolve 4-stage artifact lineage (cm-source-group-id -> cm-source-model-id -> cm-planning-model-id -> cm-target-model-id)
  const [lineageGraph, setLineageGraph] = useState<any | null>(null);
  const [isReTargeting, setIsReTargeting] = useState<boolean>(false);

  const resolvedSourceGroupId =
    loadedInfraDetail?.label?.['cm-source-group-id'] ||
    lineageGraph?.sourceGroupId ||
    activeSavedModel?.lineage?.sourceGroupId ||
    selectedSourceModel?.lineage?.sourceGroupId ||
    '';
  const resolvedSourceModelId =
    loadedInfraDetail?.label?.['cm-source-model-id'] ||
    lineageGraph?.sourceArtifact?.id ||
    activeSavedModel?.lineage?.sourceModelId ||
    selectedSourceModel?.lineage?.sourceModelId ||
    '';
  const resolvedPlanningModelId =
    loadedInfraDetail?.label?.['cm-planning-model-id'] ||
    lineageGraph?.planningArtifact?.id ||
    activeSavedModel?.lineage?.planningModelId ||
    '';
  const resolvedTargetModelId =
    targetModelIdLabel ||
    lineageGraph?.targetArtifact?.id ||
    activeSavedModel?.lineage?.targetModelId ||
    activeSavedModel?.id ||
    '';

  const lineageLookupSeedId = resolvedTargetModelId || resolvedPlanningModelId || resolvedSourceModelId;

  useEffect(() => {
    fetchSavedSourceModels();
  }, []);

  useEffect(() => {
    let cancelled = false;
    if (
      !lineageLookupSeedId ||
      lineageLookupSeedId.startsWith('sample-') ||
      lineageLookupSeedId.startsWith('cloud-model-') ||
      isBuiltInSampleModel(lineageLookupSeedId)
    ) {
      setLineageGraph(null);
      return;
    }
    artifactApi
      .getArtifactLineage(lineageLookupSeedId)
      .then((graph) => {
        if (!cancelled && graph) {
          setLineageGraph(graph);
        }
      })
      .catch(() => {
        if (!cancelled) {
          setLineageGraph(null);
        }
      });
    return () => {
      cancelled = true;
    };
  }, [lineageLookupSeedId]);

  const handleReTargetToAnotherCloud = async () => {
    if (isReTargeting) return;
    setIsReTargeting(true);
    try {
      const candidateArtifactId = resolvedPlanningModelId || resolvedSourceModelId;
      let modelToLoad =
        savedSourceModels.find((m) => candidateArtifactId && m.id === candidateArtifactId) || null;
      if (
        !modelToLoad &&
        candidateArtifactId &&
        !candidateArtifactId.startsWith('sample-') &&
        !isBuiltInSampleModel(candidateArtifactId)
      ) {
        modelToLoad = await artifactApi.getSourceModel(candidateArtifactId).catch(() => null);
      }
      if (!modelToLoad) {
        modelToLoad = selectedSourceModel || savedSourceModels[0] || null;
      }
      if (modelToLoad) {
        selectSourceModel(modelToLoad);
      }
      selectCloudModel(null);
      setRecommendationMode('multi');
      setActiveTab('design');
    } finally {
      setIsReTargeting(false);
    }
  };

  // Extract real VNet, Subnet, Security Groups, SSH Key from loadedInfraDetail
  const liveNodes = loadedInfraDetail?.node || [];
  const firstNode = liveNodes[0];

  const extractedVNetName = firstNode?.vNetId || 'vnet-migrated';
  const extractedSubnetName = firstNode?.subnetId || 'subnet-default';
  const extractedSshKeyName = firstNode?.sshKeyId || `key-${loadedInfraDetail?.name || selectedInfraId || 'mig01'}`;

  // Extract unique subnets across all live nodes and infer /24 CIDR from private IP when needed
  const inferSubnetCidr = (subnetName: string, idx: number): string => {
    const savedSub = activeSavedModel?.cloudInfraModel?.targetVNet?.subnetInfoList?.find(
      (s: any) => s.name === subnetName
    );
    const savedCidr = savedSub?.ipv4_CIDR || (savedSub as any)?.cidrBlock;
    if (savedCidr) {
      return savedCidr;
    }
    const nodeInSub = liveNodes.find(
      (n: any) => n.subnetId === subnetName && (n.privateIP || n.privateIp) && (n.privateIP || n.privateIp) !== 'N/A'
    );
    const privIp = nodeInSub?.privateIP || nodeInSub?.privateIp || '';
    const octets = privIp.split('.');
    if (octets.length === 4) {
      return `${octets[0]}.${octets[1]}.${octets[2]}.0/24`;
    }
    return `10.0.${idx + 1}.0/24`;
  };

  const uniqueLiveSubnets: string[] = Array.from(
    new Set(liveNodes.map((n: any) => n.subnetId).filter(Boolean))
  );
  const extractedSubnetsList =
    uniqueLiveSubnets.length > 0
      ? uniqueLiveSubnets.map((subName, idx) => ({
          name: subName,
          ipv4_CIDR: inferSubnetCidr(subName, idx),
          cidrBlock: inferSubnetCidr(subName, idx)
        }))
      : [{ name: extractedSubnetName, ipv4_CIDR: '10.0.1.0/24', cidrBlock: '10.0.1.0/24' }];

  // Extract unique security group IDs across all nodes
  const extractedSgNames: string[] = Array.from(
    new Set(
      liveNodes.flatMap((n: any) => n.securityGroupIds || []).filter(Boolean)
    )
  );

  const knownModelSecurityGroups: any[] = [
    ...(activeSavedModel?.cloudInfraModel?.targetSecurityGroupList || []),
    ...(selectedCloudModel?.cloudInfraModel?.targetSecurityGroupList || []),
    ...(editedCandidate?.targetSecurityGroupList || []),
    ...(recommendationCandidates?.[selectedCandidateIndex]?.targetSecurityGroupList || []),
    ...(loadedInfraDetail?.targetSecurityGroupList || [])
  ];

  const resolveFirewallRulesForSg = (sgName: string): { rules: any[]; isFallback: boolean } => {
    const liveMatch = liveSecurityGroups.find(
      (sg: any) =>
        sg?.name === sgName ||
        sg?.cspSecurityGroupId === sgName ||
        (sg?.name && (sgName.includes(sg.name) || sg.name.includes(sgName)))
    );
    if (liveMatch && Array.isArray(liveMatch.rules) && liveMatch.rules.length > 0) {
      const inboundRules = liveMatch.rules
        .filter((r: any) => String(r.direction || 'inbound').toLowerCase() !== 'outbound')
        .map((r: any) => ({
          protocol: r.protocol || 'tcp',
          dstPorts: r.fromPort && r.toPort && r.fromPort !== r.toPort ? `${r.fromPort}-${r.toPort}` : (r.fromPort || r.toPort || '22'),
          fromPort: r.fromPort,
          toPort: r.toPort,
          cidr: r.cidr || '0.0.0.0/0',
          direction: r.direction || 'inbound'
        }));
      if (inboundRules.length > 0) {
        return { rules: inboundRules, isFallback: false };
      }
    }

    const modelMatch = knownModelSecurityGroups.find(
      (sg: any) =>
        sg?.name === sgName ||
        (sg?.name && (sgName.includes(sg.name) || sg.name.includes(sgName)))
    );
    if (modelMatch && Array.isArray(modelMatch.firewallRules) && modelMatch.firewallRules.length > 0) {
      return { rules: modelMatch.firewallRules, isFallback: false };
    }

    return { rules: [{ protocol: 'tcp', dstPorts: '22, 80, 443' }], isFallback: true };
  };

  const targetSecurityGroupList = extractedSgNames.length > 0
    ? extractedSgNames.map((sgName: string) => {
        const resolved = resolveFirewallRulesForSg(sgName);
        return {
          name: sgName,
          firewallRules: resolved.rules,
          isFallback: resolved.isFallback
        };
      })
    : Array.isArray(loadedInfraDetail?.targetSecurityGroupList) && loadedInfraDetail.targetSecurityGroupList.length > 0
    ? loadedInfraDetail.targetSecurityGroupList.map((sg: any) => ({
        ...sg,
        isFallback: false
      }))
    : [{ name: 'sg-default', firewallRules: [{ protocol: 'tcp', dstPorts: '22, 80, 443' }], isFallback: true }];

  // Normalize model data for TopologyMap and VM Access table, synchronizing live node subnetId & deployedNodes
  const currentCloudModel = (() => {
    const liveDerivedNodeGroups =
      liveNodes.length > 0
        ? Array.from(
            new Set(
              liveNodes.map(
                (n: any) => n.nodeGroupId || n.subGroupId || n.name?.replace(/-\d+$/, '') || 'default-group'
              )
            )
          ).map((ngId: any) => {
            const ngNodes = liveNodes.filter(
              (n: any) => (n.nodeGroupId || n.subGroupId || n.name?.replace(/-\d+$/, '') || 'default-group') === ngId
            );
            const firstNgNode = ngNodes[0] || {};
            const specName = firstNgNode.spec?.cspSpecName || firstNgNode.cspSpecName || firstNgNode.specId || 'c5.large';
            const vcpu = firstNgNode.spec?.vCPU || firstNgNode.vCPU || (specName.includes('small') ? 2 : specName.includes('xlarge') ? 4 : 2);
            const memGiB = firstNgNode.spec?.memoryGiB || firstNgNode.memoryGiB || (specName.includes('xlarge') ? 16 : 4);
            const imageName = firstNgNode.cspImageName || firstNgNode.imageId || 'ubuntu-22.04';
            const diskSize = firstNgNode.rootDiskSize || 30;

            return {
              name: ngId,
              specId: specName,
              vCPU: vcpu,
              memoryGiB: memGiB,
              nodeGroupSize: ngNodes.length,
              imageId: imageName,
              rootDiskSize: diskSize,
              subnetId: firstNgNode.subnetId || extractedSubnetName,
              securityGroupIds: firstNgNode.securityGroupIds || extractedSgNames,
              deployedNodes: ngNodes
            };
          })
        : [];

    const hasProvisionedJobNlb =
      activeJobInfra?.nlbStatus === 'Success' || activeJobInfra?.nlbStatus === 'Provisioning';

    const resolveActiveNlbList = (modelNlbs?: any[]) => {
      if (selectedInfraId === SAMPLE_INFRA_ID) {
        return modelNlbs && modelNlbs.length > 0 ? modelNlbs : SAMPLE_INFRA_DETAIL.targetNlbList;
      }
      if (liveNlbs.length > 0) {
        return liveNlbs;
      }
      if (!hasFetchedLiveNlbs && hasProvisionedJobNlb && activeJobInfra?.targetNlbList?.length) {
        return activeJobInfra.targetNlbList;
      }
      return undefined;
    };

    if (activeSavedModel?.cloudInfraModel) {
      const base = activeSavedModel.cloudInfraModel;
      const baseSubnets = base.targetVNet?.subnetInfoList || [];
      const mergedSubnets = [...baseSubnets];
      extractedSubnetsList.forEach((liveSub) => {
        if (!mergedSubnets.some((s: any) => s.name === liveSub.name)) {
          mergedSubnets.push(liveSub as any);
        }
      });

      const enrichedNodeGroups = (base.targetInfra?.nodeGroups || []).map((ng: any) => {
        const ngNodes = liveNodes.filter(
          (n: any) =>
            n.nodeGroupId === ng.name ||
            n.subGroupId === ng.name ||
            n.name === ng.name ||
            (typeof n.name === 'string' && n.name.startsWith(`${ng.name}-`))
        );
        const liveSubnetId = ngNodes[0]?.subnetId;
        return {
          ...ng,
          subnetId: liveSubnetId || ng.subnetId || mergedSubnets[0]?.name || extractedSubnetName,
          deployedNodes: ngNodes
        };
      });

      const matchedLiveNodesCount = enrichedNodeGroups.reduce(
        (sum: number, ng: any) => sum + (Array.isArray(ng.deployedNodes) ? ng.deployedNodes.length : 0),
        0
      );
      const finalNodeGroups =
        liveNodes.length > 0 && matchedLiveNodesCount === 0
          ? liveDerivedNodeGroups
          : enrichedNodeGroups;

      const resolvedNlbList = resolveActiveNlbList(base.targetNlbList);

      return {
        ...base,
        targetNlbList: resolvedNlbList,
        targetVNet: {
          ...base.targetVNet,
          subnetInfoList: mergedSubnets.length > 0 ? mergedSubnets : extractedSubnetsList
        },
        targetInfra: {
          ...base.targetInfra,
          nodeGroups: finalNodeGroups
        }
      };
    }

    const fallbackNlbList = resolveActiveNlbList(activeJobInfra?.targetNlbList);

    return {
      name: loadedInfraDetail?.name || loadedInfraDetail?.id || activeJobInfra?.infraId || selectedInfraId || 'migrated-infra',
      status: 'Running',
      description: 'Migrated Infra',
      targetCloud: {
        csp: (firstNode?.region?.providerName || loadedInfraDetail?.id?.split('-')?.[1] || activeJobInfra?.csp || 'AWS').toUpperCase(),
        region: firstNode?.region?.regionName || activeJobInfra?.region || 'ap-northeast-2'
      },
      targetVNet: {
        name: extractedVNetName,
        cidrBlock: '10.0.0.0/16',
        subnetInfoList: extractedSubnetsList
      },
      targetSshKey: {
        name: extractedSshKeyName
      },
      targetSecurityGroupList,
      targetSpecList: liveNodes.map((n: any) => n.spec).filter((sp: any) => sp && typeof sp === 'object'),
      targetNlbList: fallbackNlbList,
      targetInfra: {
        name: loadedInfraDetail?.name || loadedInfraDetail?.id || activeJobInfra?.infraId || selectedInfraId || 'migrated-infra',
        nodeGroups:
          liveDerivedNodeGroups.length > 0
            ? liveDerivedNodeGroups
            : activeJobInfra?.vms
            ? [{ name: 'worker-group', specId: activeJobInfra.vms[0]?.specId || 'c5.large', vCPU: 2, memoryGiB: 4, nodeGroupSize: activeJobInfra.vms.length, imageId: 'ubuntu-22.04', rootDiskSize: 30, subnetId: extractedSubnetName, securityGroupIds: [] }]
            : [{ name: 'default-group', specId: 'c5.large', vCPU: 2, memoryGiB: 4, nodeGroupSize: 2, imageId: 'ubuntu-22.04', rootDiskSize: 30, subnetId: extractedSubnetName, securityGroupIds: [] }]
      }
    };
  })();

  // Extract VM nodes array for VM Access table
  const activeNodes: { name: string; specId: string; publicIp: string; privateIp: string; subnetId?: string; status?: string; systemMessage?: string; label?: Record<string, string>; subGroupId?: string; securityGroupIds?: string[] }[] = loadedInfraDetail?.node && loadedInfraDetail.node.length > 0
    ? loadedInfraDetail.node.map((n: any) => ({
        name: n.name || n.id || 'node',
        specId: n.specId || n.cspSpecName || n.spec?.cspSpecName || (typeof n.spec === 'string' ? n.spec : '') || 'c5.large',
        publicIp: n.publicIP || n.publicIp || 'N/A',
        privateIp: n.privateIP || n.privateIp || 'N/A',
        subnetId: n.subnetId || '',
        status: n.status || ((n.publicIP || n.publicIp || n.privateIP || n.privateIp) ? 'Running' : 'Failed'),
        systemMessage: n.systemMessage || '',
        label: n.label,
        subGroupId: n.subGroupId || n.nodeGroupId || '',
        securityGroupIds: Array.isArray(n.securityGroupIds) ? n.securityGroupIds : []
      }))
    : activeJobInfra?.vms && activeJobInfra.vms.length > 0
    ? activeJobInfra.vms
    : currentCloudModel.targetInfra?.nodeGroups.map((ng: any, idx: number) => ({
        name: `${ng.name}-${idx + 1}`,
        specId: ng.specId,
        publicIp: `54.180.${10 + idx}.${25 + idx}`,
        privateIp: `10.0.1.${100 + idx}`,
        subnetId: ng.subnetId || '',
        status: 'Running',
        systemMessage: '',
        label: ng.label,
        subGroupId: ng.name,
        securityGroupIds: Array.isArray(ng.securityGroupIds) ? ng.securityGroupIds : []
      }));

  const resolveNodeServicePorts = (node: any, lineagePorts?: number[]): number[] => {
    const portSet = new Set<number>(lineagePorts || []);
    const nodeGroups = currentCloudModel.targetInfra?.nodeGroups || [];
    const matchedNg = nodeGroups.find(
      (ng: any) =>
        ng?.name &&
        (ng.name === node.subGroupId ||
          ng.name === node.name ||
          (typeof node.name === 'string' && node.name.startsWith(`${ng.name}-`)))
    );
    const nodeSgIds: string[] =
      Array.isArray(node.securityGroupIds) && node.securityGroupIds.length > 0
        ? node.securityGroupIds
        : Array.isArray(matchedNg?.securityGroupIds) && matchedNg.securityGroupIds.length > 0
        ? matchedNg.securityGroupIds
        : [];

    const allSgs = [
      ...(targetSecurityGroupList || []),
      ...((currentCloudModel as any)?.targetSecurityGroupList || [])
    ];

    const matchedSgs =
      nodeSgIds.length > 0
        ? allSgs.filter((sg: any) =>
            nodeSgIds.some(
              (id) =>
                sg?.name === id ||
                (sg?.name && (id.includes(sg.name) || sg.name.includes(id)))
            )
          )
        : allSgs.length === 1
        ? allSgs
        : [];

    for (const sg of matchedSgs) {
      if (sg?.isFallback) continue;
      for (const rule of sg?.firewallRules || []) {
        const dir = String(rule?.direction || 'inbound').toLowerCase();
        const proto = String(rule?.protocol || 'tcp').toLowerCase();
        if (dir === 'outbound' || (proto !== 'tcp' && proto !== 'all' && proto !== '*')) continue;
        for (const p of parseDiscreteTcpPorts(rule?.dstPorts ?? rule?.fromPort, rule?.toPort)) {
          portSet.add(p);
        }
      }
    }

    return Array.from(portSet).sort((a, b) => a - b);
  };

  return (
    <div className="space-y-6 animate-fade-in">

      {/* 1. Single-Line Tab Description Box */}
      <div className="glass-panel px-6 py-4.5 rounded-2xl border border-border-main flex flex-wrap items-center gap-x-3 gap-y-1.5">
        <div className="flex items-center gap-2 shrink-0">
          <Server className="w-5 h-5 text-emerald-500" />
          <h2 className="text-base font-extrabold text-text-main tracking-tight">
            Cloud Infra Overview
          </h2>
        </div>
        <span className="text-sm text-text-muted">
          Visualize topology, inspect allocated public/private IPs, review compliance reports, or delete deployed infrastructures.
        </span>
      </div>

      {/* 2. Overview Category Sub-Tabs (Cloud Infrastructures vs Object Storages) */}
      <div className="flex flex-wrap items-center gap-2">
        <button
          onClick={() => setOverviewCategory('infra')}
          className={`px-4 py-2.5 rounded-xl text-xs font-extrabold flex items-center gap-2 transition cursor-pointer ${
            overviewCategory === 'infra'
              ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
              : 'bg-bg-panel border border-border-main text-text-muted hover:text-text-main'
          }`}
        >
          <Server className="w-4 h-4" />
          <span>Cloud Infrastructures ({migratedInfraIds.length})</span>
        </button>

        <button
          onClick={() => setOverviewCategory('storage')}
          className={`px-4 py-2.5 rounded-xl text-xs font-extrabold flex items-center gap-2 transition cursor-pointer ${
            overviewCategory === 'storage'
              ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
              : 'bg-bg-panel border border-border-main text-text-muted hover:text-text-main'
          }`}
        >
          <HardDrive className="w-4 h-4" />
          <span>Object Storages ({migratedStorages.length})</span>
        </button>
      </div>

      {overviewCategory === 'infra' && (
        <div className="space-y-6">
          {/* 3. Cloud Infrastructures Section (Cards vs Table Toggle) */}
          <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-4">
            <div className="flex flex-wrap justify-between items-center gap-3 border-b border-border-main/20 pb-3">
              <h3 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                <Server className="w-4 h-4 text-emerald-500" />
                Cloud Infrastructures ({migratedInfraIds.length})
              </h3>

          <div className="flex items-center gap-2">
            {/* View Switcher Toggle (Cards vs Table) */}
            <div className="flex items-center bg-bg-panel/80 p-0.5 rounded-lg border border-border-main">
              <button
                onClick={() => setCatalogViewMode('grid')}
                className={`p-1.5 px-2.5 rounded-md transition cursor-pointer text-xs flex items-center gap-1 font-bold font-mono ${
                  catalogViewMode === 'grid'
                    ? 'bg-emerald-500 text-slate-950 shadow-sm'
                    : 'text-text-muted hover:text-text-main'
                }`}
                title="Cards Grid View"
              >
                <LayoutGrid className="w-3.5 h-3.5" />
                <span>Cards</span>
              </button>
              <button
                onClick={() => setCatalogViewMode('table')}
                className={`p-1.5 px-2.5 rounded-md transition cursor-pointer text-xs flex items-center gap-1 font-bold font-mono ${
                  catalogViewMode === 'table'
                    ? 'bg-emerald-500 text-slate-950 shadow-sm'
                    : 'text-text-muted hover:text-text-main'
                }`}
                title="Compact Table Row View"
              >
                <List className="w-3.5 h-3.5" />
                <span>Table</span>
              </button>
            </div>

            <button
              onClick={fetchMigratedInfraIds}
              disabled={isLoadingIds}
              className="px-3 py-1.5 bg-bg-panel hover:bg-border-main/20 text-emerald-500 rounded-lg border border-border-main text-xs font-bold transition cursor-pointer flex items-center gap-1.5 font-mono"
              title="Refresh migrated infra list from API"
            >
              <RefreshCw className={`w-3.5 h-3.5 ${isLoadingIds ? 'animate-spin' : ''}`} />
              <span>Refresh</span>
            </button>
          </div>
        </div>

        {migratedInfraIds.length > 0 ? (
          catalogViewMode === 'grid' ? (
            /* Cards Grid View (Compact Width) */
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-3">
              {migratedInfraIds.map((infraId) => {
                const isSelected = selectedInfraId === infraId;
                const isTerminating = !!deletingInfrasMap[infraId];
                const matchingActiveJob = jobs.find(j => (j.infraId === infraId || j.id === infraId || (j.reqId && j.reqId === infraId)) && (j.status === 'Handling' || j.nlbStatus === 'Provisioning'));
                const isMigrating = !!matchingActiveJob;
                const matchingJob = completedJobs.find(j => j.infraId === infraId || j.id === infraId) || jobs.find(j => j.infraId === infraId || j.id === infraId);
                const csp = matchingJob?.csp || (infraId.toLowerCase().includes('aws') ? 'AWS' : infraId.toLowerCase().includes('azure') ? 'AZURE' : 'GCP');
                const region = matchingJob?.region || 'ap-northeast-2';

                return (
                  <div
                    key={infraId}
                    className={`p-4 rounded-xl border text-left transition-all duration-200 flex flex-col justify-between space-y-3 relative overflow-hidden ${
                      isTerminating || isMigrating
                        ? 'bg-amber-500/5 border-amber-500/40'
                        : isSelected
                        ? 'bg-emerald-500/10 border-emerald-500/60 shadow-lg shadow-emerald-500/10 ring-1 ring-emerald-500/40'
                        : 'bg-bg-panel/40 border-border-main/40 hover:bg-bg-panel/80 hover:border-border-main'
                    }`}
                  >
                    <div className="flex justify-between items-start gap-2">
                      <div className="space-y-1 min-w-0">
                        <div className="flex items-center gap-1.5 flex-wrap">
                          {infraId === SAMPLE_INFRA_ID && (
                            <span className="text-xs font-extrabold px-2 py-0.5 rounded-md bg-amber-500/20 text-amber-600 dark:text-amber-300 border border-amber-500/40 font-mono shadow-sm">
                              [Sample]
                            </span>
                          )}
                          <span className="text-xs font-extrabold px-2 py-0.5 rounded bg-emerald-500/10 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 font-mono">
                            {csp}
                          </span>
                          <span className="text-xs font-bold text-text-muted font-mono truncate max-w-[120px]">
                            {region}
                          </span>
                        </div>
                        <h4 className="font-extrabold text-sm text-text-main truncate max-w-[200px]" title={infraId}>
                          {infraId}
                        </h4>
                      </div>

                      {isTerminating ? (
                        <span className="text-xs font-bold px-2 py-0.5 rounded-full bg-amber-500/10 text-amber-500 border border-amber-500/30 flex items-center gap-1 font-mono shrink-0">
                          <span className="w-1.5 h-1.5 bg-amber-500 rounded-full animate-ping" />
                          Terminating...
                        </span>
                      ) : isMigrating ? (
                        <span className="text-xs font-bold px-2 py-0.5 rounded-full bg-amber-500/10 text-amber-500 border border-amber-500/30 flex items-center gap-1 font-mono shrink-0">
                          <span className="w-1.5 h-1.5 bg-amber-500 rounded-full animate-ping" />
                          {matchingActiveJob?.nlbStatus === 'Provisioning' ? 'NLB Migrating...' : 'Migrating...'}
                        </span>
                      ) : (
                        <span className="text-xs font-bold px-2 py-0.5 rounded-full bg-emerald-100 dark:bg-emerald-950/50 text-emerald-600 dark:text-emerald-400 border border-emerald-300 dark:border-emerald-800/40 flex items-center gap-1 font-mono shrink-0">
                          <span className="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse" />
                          Deployed
                        </span>
                      )}
                    </div>

                    {/* Actions: Load Detail & Delete */}
                    <div className="flex items-center justify-between gap-2 pt-2 border-t border-border-main/20">
                      <button
                        onClick={() => handleLoadInfraDetail(infraId)}
                        disabled={isTerminating || isMigrating || (isLoadingDetail && selectedInfraId === infraId)}
                        className={`flex-1 py-1.5 px-2.5 rounded-lg text-xs font-bold flex items-center justify-center gap-1.5 transition cursor-pointer font-mono ${
                          isTerminating || isMigrating
                            ? 'bg-bg-input text-text-muted border border-border-main opacity-50 cursor-not-allowed'
                            : isSelected
                            ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
                            : 'bg-bg-input hover:bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-border-main'
                        }`}
                        title={isTerminating ? 'Infrastructure is currently terminating' : isMigrating ? 'Infrastructure migration in progress' : `Load detail for ${infraId}`}
                      >
                        {isTerminating ? (
                          <Loader2 className="w-3.5 h-3.5 animate-spin text-amber-500" />
                        ) : isMigrating ? (
                          <Loader2 className="w-3.5 h-3.5 animate-spin text-amber-500" />
                        ) : isLoadingDetail && selectedInfraId === infraId ? (
                          <Loader2 className="w-3.5 h-3.5 animate-spin" />
                        ) : (
                          <Eye className="w-3.5 h-3.5" />
                        )}
                        <span>{isTerminating ? 'Terminating...' : isMigrating ? 'Migrating...' : isSelected ? 'Loaded ✓' : 'Load Detail'}</span>
                      </button>

                      <button
                        onClick={() => handleOpenDeleteModalForInfra(infraId)}
                        disabled={isTerminating || isMigrating || infraId === SAMPLE_INFRA_ID}
                        className="py-1.5 px-2.5 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-lg text-xs font-bold flex items-center justify-center gap-1 transition cursor-pointer font-mono disabled:opacity-30 disabled:cursor-not-allowed"
                        title={infraId === SAMPLE_INFRA_ID ? 'Sample infrastructure cannot be deleted' : isTerminating ? 'Termination already in progress' : isMigrating ? 'Migration in progress' : `Delete infrastructure ${infraId}`}
                      >
                        <Trash2 className="w-3.5 h-3.5" />
                        <span>Delete</span>
                      </button>
                    </div>
                  </div>
                );
              })}
            </div>
          ) : (
            /* Compact Table Row View */
            <div className="overflow-x-auto rounded-xl border border-border-main/40 bg-bg-panel/20">
              <table className="w-full text-left text-xs font-mono">
                <thead>
                  <tr className="border-b border-border-main/40 bg-bg-panel/80 text-text-muted font-bold">
                    <th className="py-2.5 px-4 w-12 text-center">#</th>
                    <th className="py-2.5 px-4">Infrastructure ID</th>
                    <th className="py-2.5 px-4">CSP / Region</th>
                    <th className="py-2.5 px-4">Status</th>
                    <th className="py-2.5 px-4 text-right">Actions</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-border-main/20">
                  {migratedInfraIds.map((infraId, idx) => {
                    const isSelected = selectedInfraId === infraId;
                    const isTerminating = !!deletingInfrasMap[infraId];
                    const matchingActiveJob = jobs.find(j => (j.infraId === infraId || j.id === infraId || (j.reqId && j.reqId === infraId)) && (j.status === 'Handling' || j.nlbStatus === 'Provisioning'));
                    const isMigrating = !!matchingActiveJob;
                    const matchingJob = completedJobs.find(j => j.infraId === infraId || j.id === infraId) || jobs.find(j => j.infraId === infraId || j.id === infraId);
                    const csp = matchingJob?.csp || (infraId.toLowerCase().includes('aws') ? 'AWS' : infraId.toLowerCase().includes('azure') ? 'AZURE' : 'GCP');
                    const region = matchingJob?.region || 'ap-northeast-2';

                    return (
                      <tr
                        key={infraId}
                        className={`transition hover:bg-bg-panel/60 ${
                          isSelected ? 'bg-emerald-500/10 font-bold' : ''
                        }`}
                      >
                        <td className="py-2.5 px-4 text-center text-text-muted">{idx + 1}</td>
                        <td className="py-2.5 px-4 font-bold text-text-main flex items-center gap-1.5">
                          {infraId === SAMPLE_INFRA_ID && (
                            <span className="text-xs font-extrabold px-2 py-0.5 rounded-md bg-amber-500/20 text-amber-600 dark:text-amber-300 border border-amber-500/40 font-mono shadow-sm">
                              [Sample]
                            </span>
                          )}
                          <span>{infraId}</span>
                        </td>
                        <td className="py-2.5 px-4">
                          <span className="text-emerald-600 dark:text-emerald-400 font-extrabold mr-2">
                            {csp}
                          </span>
                          <span className="text-text-muted">{region}</span>
                        </td>
                        <td className="py-2.5 px-4">
                          {isTerminating ? (
                            <span className="inline-flex items-center gap-1 text-amber-500 font-bold">
                              <span className="w-1.5 h-1.5 bg-amber-500 rounded-full animate-ping" />
                              Terminating...
                            </span>
                          ) : isMigrating ? (
                            <span className="inline-flex items-center gap-1 text-amber-500 font-bold">
                              <span className="w-1.5 h-1.5 bg-amber-500 rounded-full animate-ping" />
                              {matchingActiveJob?.nlbStatus === 'Provisioning' ? 'NLB Migrating...' : 'Migrating...'}
                            </span>
                          ) : (
                            <span className="inline-flex items-center gap-1 text-emerald-600 dark:text-emerald-400 font-bold">
                              <span className="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse" />
                              Deployed
                            </span>
                          )}
                        </td>
                        <td className="py-2.5 px-4 text-right">
                          <div className="flex items-center justify-end gap-2">
                            <button
                              onClick={() => handleLoadInfraDetail(infraId)}
                              disabled={isTerminating || isMigrating || (isLoadingDetail && selectedInfraId === infraId)}
                              className={`py-1 px-3 rounded-lg text-xs font-bold flex items-center gap-1.5 transition cursor-pointer ${
                                isTerminating || isMigrating
                                  ? 'bg-bg-input text-text-muted border border-border-main opacity-50 cursor-not-allowed'
                                  : isSelected
                                  ? 'bg-emerald-500 text-slate-950 shadow-sm'
                                  : 'bg-bg-input hover:bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-border-main'
                              }`}
                            >
                              {isTerminating ? (
                                <Loader2 className="w-3.5 h-3.5 animate-spin text-amber-500" />
                              ) : isMigrating ? (
                                <Loader2 className="w-3.5 h-3.5 animate-spin text-amber-500" />
                              ) : isLoadingDetail && selectedInfraId === infraId ? (
                                <Loader2 className="w-3.5 h-3.5 animate-spin" />
                              ) : (
                                <Eye className="w-3.5 h-3.5" />
                              )}
                              <span>{isTerminating ? 'Terminating...' : isMigrating ? 'Migrating...' : isSelected ? 'Loaded ✓' : 'Load Detail'}</span>
                            </button>

                            <button
                              onClick={() => handleOpenDeleteModalForInfra(infraId)}
                              disabled={isTerminating || isMigrating || infraId === SAMPLE_INFRA_ID}
                              className="py-1 px-3 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-lg text-xs font-bold flex items-center gap-1 transition cursor-pointer disabled:opacity-30 disabled:cursor-not-allowed"
                              title={infraId === SAMPLE_INFRA_ID ? 'Sample infrastructure cannot be deleted' : isTerminating ? 'Termination already in progress' : isMigrating ? 'Migration in progress' : `Delete infrastructure ${infraId}`}
                            >
                              <Trash2 className="w-3.5 h-3.5" />
                              <span>Delete</span>
                            </button>
                          </div>
                        </td>
                      </tr>
                    );
                  })}
                </tbody>
              </table>
            </div>
          )) : (
            <div className="py-8 text-center text-text-muted text-xs font-mono italic">
              No deployed infrastructures found. Execute a migration in Tab 4 to deploy new infrastructure.
            </div>
          )}
        </div>



      {deleteSuccessMsg && (
        <div className="p-4 bg-emerald-500/10 border border-emerald-500/20 text-emerald-600 dark:text-emerald-400 rounded-2xl text-sm font-bold flex items-center gap-2">
          <CheckCircle2 className="w-5 h-5" />
          <span>{deleteSuccessMsg}</span>
        </div>
      )}

      {selectedInfraId || (migratedInfraIds.length > 0) || loadedInfraDetail ? (
        <div className="space-y-6">

          {/* End-to-End Migration Lineage Chain & Multi-Cloud Re-Targeting Bar */}
          <div className="glass-panel p-5 rounded-2xl border border-emerald-500/30 space-y-3.5">
            <div className="flex flex-wrap items-center justify-between gap-3">
              <div className="flex items-center gap-2.5">
                <GitBranch className="w-5 h-5 text-emerald-500 shrink-0" />
                <div>
                  <h3 className="text-base font-extrabold text-text-main">
                    End-to-End Migration Artifact Lineage
                  </h3>
                  <p className="text-sm text-text-muted">
                    Traces provisioned cloud infrastructure back to its canonical imdl Source, Planning, and Target artifacts for repeatable Multi-Cloud re-targeting.
                  </p>
                </div>
              </div>

              <button
                type="button"
                onClick={handleReTargetToAnotherCloud}
                disabled={isReTargeting}
                className="px-4 py-2 bg-emerald-600 hover:bg-emerald-700 disabled:opacity-50 text-white rounded-xl text-sm font-extrabold transition flex items-center gap-2 cursor-pointer shrink-0 shadow-sm"
                title="Load the linked Planning or Source artifact into 2. Target Cloud Optimizer to deploy onto another CSP"
              >
                <Compass className={`w-4 h-4 ${isReTargeting ? 'animate-spin' : ''}`} />
                <span>{isReTargeting ? 'Loading Artifact...' : 'Re-Target to Another Cloud'}</span>
                <ArrowRight className="w-4 h-4" />
              </button>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-5 gap-2.5 text-sm">
              <div className="p-3 rounded-xl bg-bg-input/60 border border-border-main/60 space-y-0.5">
                <span className="block text-xs font-normal text-text-muted">1. Source Group</span>
                <span className="block font-extrabold font-mono text-text-main truncate" title={resolvedSourceGroupId || 'On-Premise Group'}>
                  {resolvedSourceGroupId || 'onprem-source-group'}
                </span>
              </div>

              <div className="p-3 rounded-xl bg-bg-input/60 border border-border-main/60 space-y-0.5">
                <span className="block text-xs font-normal text-text-muted">2. Source Model (imdl)</span>
                <span
                  className="block font-extrabold font-mono text-emerald-600 dark:text-emerald-400 truncate"
                  title={lineageGraph?.sourceArtifact?.name || resolvedSourceModelId || selectedSourceModel?.name || 'Source Artifact'}
                >
                  {lineageGraph?.sourceArtifact?.name ||
                    savedSourceModels.find((m) => m.id === resolvedSourceModelId)?.name ||
                    resolvedSourceModelId ||
                    selectedSourceModel?.name ||
                    'OnpremInfra (source)'}
                </span>
              </div>

              <div className="p-3 rounded-xl bg-bg-input/60 border border-border-main/60 space-y-0.5">
                <span className="block text-xs font-normal text-text-muted">3. Planning Model</span>
                <span
                  className="block font-extrabold font-mono text-teal-600 dark:text-teal-400 truncate"
                  title={lineageGraph?.planningArtifact?.name || resolvedPlanningModelId || 'Direct from Source'}
                >
                  {lineageGraph?.planningArtifact?.name ||
                    savedSourceModels.find((m) => m.id === resolvedPlanningModelId)?.name ||
                    resolvedPlanningModelId ||
                    'Direct Source Plan'}
                </span>
              </div>

              <div className="p-3 rounded-xl bg-bg-input/60 border border-border-main/60 space-y-0.5">
                <span className="block text-xs font-normal text-text-muted">4. Target Cloud Model</span>
                <span
                  className="block font-extrabold font-mono text-emerald-600 dark:text-emerald-400 truncate"
                  title={lineageGraph?.targetArtifact?.name || activeSavedModel?.name || resolvedTargetModelId || 'RecommendedInfra'}
                >
                  {lineageGraph?.targetArtifact?.name ||
                    activeSavedModel?.name ||
                    resolvedTargetModelId ||
                    'RecommendedInfra'}
                </span>
              </div>

              <div className="p-3 rounded-xl bg-emerald-500/10 border border-emerald-500/30 space-y-0.5">
                <span className="block text-xs font-normal text-text-muted">5. Provisioned Cloud Infra</span>
                <span
                  className="block font-extrabold font-mono text-text-main truncate"
                  title={`${currentCloudModel.name} (${currentCloudModel.targetCloud?.csp?.toUpperCase()} / ${currentCloudModel.targetCloud?.region})`}
                >
                  {currentCloudModel.name} ({currentCloudModel.targetCloud?.csp?.toUpperCase()})
                </span>
              </div>
            </div>
          </div>

          {/* Infrastructure Metrics & Resources Cards Bar (5 Resource Cards) */}
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-5">
            {/* 1. Target CSP & Region */}
            <div className="glass-panel p-5 rounded-2xl border border-border-main/60 space-y-2 flex flex-col justify-center">
              <span className="text-xs font-bold text-text-muted font-mono block">Target CSP & Region</span>
              <div className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 flex items-center gap-2 truncate">
                <Cloud className="w-4 h-4 shrink-0 text-emerald-500" />
                <span className="truncate">{currentCloudModel.targetCloud?.csp?.toUpperCase()} ({currentCloudModel.targetCloud?.region})</span>
              </div>
            </div>

            {/* 2. VNet & Subnet */}
            <div className="glass-panel p-5 rounded-2xl border border-border-main/60 space-y-2 flex flex-col justify-center">
              <span className="text-xs font-bold text-text-muted font-mono block">VNet & Subnet</span>
              <div className="text-sm font-extrabold text-text-main font-mono truncate" title={`${extractedVNetName} (${currentCloudModel.targetVNet?.cidrBlock || '10.0.0.0/16'}) / ${extractedSubnetName}`}>
                {extractedVNetName} ({currentCloudModel.targetVNet?.cidrBlock || '10.0.0.0/16'})
              </div>
            </div>

            {/* 3. Security Group */}
            <div className="glass-panel p-5 rounded-2xl border border-border-main/60 space-y-2 flex flex-col justify-center">
              <span className="text-xs font-bold text-text-muted font-mono block">Security Group</span>
              <div className="text-sm font-extrabold text-amber-600 dark:text-amber-400 font-mono truncate" title={extractedSgNames.join(', ') || 'mig-sg-01'}>
                {extractedSgNames.length > 0 ? extractedSgNames.join(', ') : (currentCloudModel.targetSecurityGroupList?.[0]?.name || 'mig-sg-01')}
              </div>
            </div>

            {/* 4. SSH Access Key */}
            <div className="glass-panel p-5 rounded-2xl border border-border-main/60 space-y-2 flex flex-col justify-center">
              <span className="text-xs font-bold text-text-muted font-mono block">SSH Access Key</span>
              <div className="text-sm font-extrabold text-teal-600 dark:text-teal-400 font-mono truncate" title={extractedSshKeyName}>
                {extractedSshKeyName || currentCloudModel.targetSshKey?.name || 'key-mig01'}
              </div>
            </div>

            {/* 5. Managed NLB */}
            <div className="glass-panel p-5 rounded-2xl border border-border-main/60 space-y-2 flex flex-col justify-center">
              <span className="text-xs font-bold text-text-muted font-mono block">Managed NLB</span>
              <div className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 font-mono truncate" title={(currentCloudModel as any).targetNlbList?.[0]?.name || 'N/A'}>
                {(currentCloudModel as any).targetNlbList && (currentCloudModel as any).targetNlbList.length > 0 ? (currentCloudModel as any).targetNlbList[0].name || 'mig-nlb-01' : 'N/A (VM Direct)'}
              </div>
            </div>
          </div>

          {/* Infrastructure Topology Diagram Visualizer */}
          <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-4">
            <div className="flex justify-between items-center border-b border-border-main/20 pb-3">
              <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                <Layers className="w-5 h-5 text-emerald-500" />
                Target Cloud Infrastructure Topology Map
              </h3>
            </div>

            {/* Topology Map Container - Fluid Dynamic Height */}
            <div className="w-full min-h-[300px] rounded-xl overflow-hidden border border-border-main/40 bg-slate-50 dark:bg-slate-950">
              <TopologyMap data={currentCloudModel as any} />
            </div>
          </div>

          {/* 4. Sub-Tab Detailed Resource Inspector Box */}
          <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-5">
            <div className="flex flex-wrap items-center justify-between gap-3 border-b border-border-main/20 pb-4">
              {/* Sub-Tab Navigation Bar (Ordered by Resource Provisioning Dependency Sequence) */}
              <div className="flex items-center gap-1.5 bg-bg-panel/80 p-1 rounded-xl border border-border-main overflow-x-auto">
                {/* 1. Network Infrastructure */}
                <button
                  onClick={() => setActiveSubTab('vnets')}
                  className={`px-3.5 py-2 rounded-lg text-xs font-bold font-mono transition flex items-center gap-2 cursor-pointer ${
                    activeSubTab === 'vnets'
                      ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
                      : 'text-text-muted hover:text-text-main hover:bg-bg-panel'
                  }`}
                >
                  <Network className="w-4 h-4" />
                  <span>1. VNet & Subnets</span>
                </button>

                {/* 2. Security Firewall */}
                <button
                  onClick={() => setActiveSubTab('sgs')}
                  className={`px-3.5 py-2 rounded-lg text-xs font-bold font-mono transition flex items-center gap-2 cursor-pointer ${
                    activeSubTab === 'sgs'
                      ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
                      : 'text-text-muted hover:text-text-main hover:bg-bg-panel'
                  }`}
                >
                  <Shield className="w-4 h-4" />
                  <span>2. Security Groups ({targetSecurityGroupList.length})</span>
                </button>

                {/* 3. SSH Key Credentials */}
                <button
                  onClick={() => setActiveSubTab('ssh')}
                  className={`px-3.5 py-2 rounded-lg text-xs font-bold font-mono transition flex items-center gap-2 cursor-pointer ${
                    activeSubTab === 'ssh'
                      ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
                      : 'text-text-muted hover:text-text-main hover:bg-bg-panel'
                  }`}
                >
                  <Key className="w-4 h-4" />
                  <span>3. SSH Key</span>
                </button>

                {/* 4. Compute Nodes */}
                <button
                  onClick={() => setActiveSubTab('vms')}
                  className={`px-3.5 py-2 rounded-lg text-xs font-bold font-mono transition flex items-center gap-2 cursor-pointer ${
                    activeSubTab === 'vms'
                      ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
                      : 'text-text-muted hover:text-text-main hover:bg-bg-panel'
                  }`}
                >
                  <Server className="w-4 h-4" />
                  <span>4. Nodes ({activeNodes.length})</span>
                </button>

                {/* 5. Network Load Balancers */}
                <button
                  onClick={() => setActiveSubTab('nlb')}
                  className={`px-3.5 py-2 rounded-lg text-xs font-bold font-mono transition flex items-center gap-2 cursor-pointer ${
                    activeSubTab === 'nlb'
                      ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
                      : 'text-text-muted hover:text-text-main hover:bg-bg-panel'
                  }`}
                >
                  <Radio className="w-4 h-4" />
                  <span>5. Managed NLB</span>
                </button>
              </div>

              <span className="text-xs text-emerald-500 font-bold bg-emerald-500/10 px-3 py-1 rounded-full border border-emerald-500/20 uppercase font-mono shrink-0">
                Active Resource View
              </span>
            </div>

            {/* Sub-Tab 4: Nodes */}
            {activeSubTab === 'vms' && (
              <div className="space-y-3">
                <div className="flex flex-wrap items-center justify-between gap-3">
                  <h4 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                    <Globe className="w-4 h-4 text-teal-400" />
                    <span>Deployed Nodes &amp; SSH Readiness</span>
                  </h4>

                  <div className="flex items-center gap-2 flex-wrap">
                    {sshReadyData && (
                      <span className={`px-2.5 py-1 rounded-lg text-xs font-mono font-bold border ${
                        sshReadyData.ready
                          ? 'bg-emerald-500/15 text-emerald-500 border-emerald-500/30'
                          : 'bg-amber-500/15 text-amber-500 border-amber-500/30'
                      }`}>
                        SSH Ready: {sshReadyData.readyNodes} / {sshReadyData.totalNodes} Nodes
                      </span>
                    )}
                    <button
                      type="button"
                      onClick={handleCheckInfraSshReady}
                      disabled={selectedInfraId === SAMPLE_INFRA_ID || isCheckingSsh}
                      className="px-3 py-1.5 bg-emerald-500/10 hover:bg-emerald-500/20 disabled:opacity-40 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 rounded-lg text-xs font-extrabold transition flex items-center gap-1.5 cursor-pointer font-mono"
                    >
                      <Activity className="w-3.5 h-3.5 text-emerald-500" />
                      <span>
                        {sshCooldownSec > 0
                          ? `SSH Cooldown (${sshCooldownSec}s)`
                          : 'Check SSH Readiness (GET /ssh-ready)'}
                      </span>
                    </button>
                  </div>
                </div>

                <div className="overflow-x-auto border border-border-main/50 rounded-xl">
                  <table className="w-full text-left border-collapse text-xs font-mono">
                    <thead>
                      <tr className="border-b border-border-main bg-bg-input/60 text-text-muted font-normal">
                        <th className="py-2.5 px-4">Node Name</th>
                        <th className="py-2.5 px-4">Node Spec</th>
                        <th className="py-2.5 px-4">Public IP Address</th>
                        <th className="py-2.5 px-4">Private IP Address</th>
                        <th className="py-2.5 px-4">Node Status</th>
                        <th className="py-2.5 px-4">SSH Status</th>
                        <th className="py-2.5 px-4">SSH Command</th>
                      </tr>
                    </thead>
                    <tbody className="divide-y divide-border-main/40 text-text-main">
                      {activeNodes.map((node, idx) => {
                        const isCopied = copiedIp === node.publicIp;
                        const matchedSsh = sshReadyData?.nodeStatus?.find(
                          (ns) => ns.name === node.name || ns.id === node.name || (node.publicIp !== 'N/A' && ns.publicIP === node.publicIp)
                        );
                        const isNodeFailed =
                          node.status?.toLowerCase() === 'failed' ||
                          node.status?.toLowerCase() === 'error' ||
                          ((!node.publicIp || node.publicIp === 'N/A') && (!node.privateIp || node.privateIp === 'N/A'));
                        const nodeLineage = resolveTargetNodeLineage(
                          node,
                          currentCloudModel.targetInfra?.nodeGroups,
                          selectedSourceModel?.onpremiseInfraModel || refinedSourceInfra
                        );
                        const nodeServicePorts = !isNodeFailed && isValidEndpointHost(node.publicIp)
                          ? resolveNodeServicePorts(node, nodeLineage?.servicePorts)
                          : [];
                        return (
                          <tr key={idx} className={`transition ${isNodeFailed ? 'bg-red-500/[0.04] hover:bg-red-500/[0.08]' : 'hover:bg-emerald-500/[0.02]'}`}>
                            <td className="py-3 px-4 font-extrabold text-text-main">
                              <div className="flex flex-col gap-0.5">
                                <span>{node.name}</span>
                                {nodeLineage?.sourceHostname && (
                                  <span className="text-xs font-normal text-text-muted">
                                    Source: <span className="font-bold text-emerald-600 dark:text-emerald-400">{nodeLineage.sourceHostname}</span>
                                  </span>
                                )}
                              </div>
                            </td>
                            <td className="py-3 px-4 text-emerald-600 dark:text-emerald-400 font-extrabold">
                              <div className="flex items-center gap-1.5">
                                <span>{node.specId}</span>
                                {isGpuSpec(node.specId, (currentCloudModel as any)?.targetSpecList) && (
                                  <span className="px-1.5 py-0.5 bg-emerald-500/15 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold rounded">
                                    GPU
                                  </span>
                                )}
                              </div>
                            </td>
                            <td className={`py-3 px-4 font-extrabold ${isNodeFailed ? 'text-red-500' : 'text-text-main'}`}>
                              <div className="flex flex-wrap items-center gap-1.5">
                                <span className="select-all">{node.publicIp}</span>
                                {nodeServicePorts.map((port) => {
                                  const scheme = port === 443 || port === 8443 ? 'https' : 'http';
                                  const endpointUrl = `${scheme}://${node.publicIp.trim()}:${port}`;
                                  return (
                                    <a
                                      key={port}
                                      href={endpointUrl}
                                      target="_blank"
                                      rel="noopener noreferrer"
                                      title={`Open ${endpointUrl} in a new tab`}
                                      className="inline-flex items-center gap-1 px-2 py-0.5 rounded-md bg-emerald-500/15 hover:bg-emerald-500/25 border border-emerald-500/40 text-emerald-600 dark:text-emerald-400 text-xs font-extrabold font-mono transition cursor-pointer"
                                    >
                                      <span>:{port}</span>
                                      <ExternalLink className="w-3 h-3 shrink-0" />
                                    </a>
                                  );
                                })}
                              </div>
                            </td>
                            <td className={`py-3 px-4 ${isNodeFailed ? 'text-red-500/80' : 'text-text-muted'}`}>
                              <div className="flex items-center gap-1.5">
                                <span>{node.privateIp}</span>
                                {node.subnetId && (
                                  <span className="px-1.5 py-0.5 rounded bg-bg-panel border border-border-main text-text-muted text-xs font-mono">
                                    {node.subnetId}
                                  </span>
                                )}
                              </div>
                            </td>
                            <td className="py-3 px-4">
                              <div className="flex flex-col gap-1">
                                <span
                                  className={`inline-flex items-center w-fit px-2 py-0.5 rounded font-bold border ${
                                    isNodeFailed
                                      ? 'bg-red-500/15 text-red-500 border-red-500/30'
                                      : 'bg-teal-500/15 text-teal-600 dark:text-teal-400 border-teal-500/30'
                                  }`}
                                  title={node.systemMessage || undefined}
                                >
                                  {isNodeFailed ? (node.status || 'Failed') : (node.status || 'Running')}
                                </span>
                                {isNodeFailed && node.systemMessage && (
                                  <span
                                    className="text-xs font-sans text-red-500 max-w-xs truncate"
                                    title={node.systemMessage}
                                  >
                                    {node.systemMessage}
                                  </span>
                                )}
                              </div>
                            </td>
                            <td className="py-3 px-4">
                              {isNodeFailed ? (
                                <span className="px-2 py-0.5 rounded font-bold border bg-red-500/10 text-red-500 border-red-500/30">
                                  N/A (Failed)
                                </span>
                              ) : matchedSsh ? (
                                <span className={`px-2 py-0.5 rounded font-bold border ${
                                  matchedSsh.sshReady
                                    ? 'bg-emerald-500/15 text-emerald-500 border-emerald-500/30'
                                    : 'bg-red-500/15 text-red-500 border-red-500/30'
                                }`}>
                                  {matchedSsh.sshReady ? `READY (${matchedSsh.sshPort || 22})` : 'UNREACHABLE'}
                                </span>
                              ) : (
                                <span className="text-text-muted">Not Checked</span>
                              )}
                            </td>
                            <td className="py-3 px-4">
                              <button
                                onClick={() => handleCopySshCommand(node.publicIp, currentCloudModel.targetSshKey?.name || 'key')}
                                disabled={!node.publicIp || node.publicIp === 'N/A' || isNodeFailed}
                                className="px-3 py-1.5 bg-bg-panel border border-border-main hover:bg-emerald-500/10 hover:border-emerald-500/30 text-emerald-600 dark:text-emerald-400 text-xs font-bold rounded-lg transition cursor-pointer flex items-center gap-1.5 disabled:opacity-40 disabled:cursor-not-allowed font-mono"
                              >
                                {isCopied ? <Check className="w-3.5 h-3.5 text-green-500" /> : <Copy className="w-3.5 h-3.5" />}
                                <span>{isCopied ? 'Copied SSH Command!' : 'Copy SSH Command'}</span>
                              </button>
                            </td>
                          </tr>
                        );
                      })}
                    </tbody>
                  </table>
                </div>
              </div>
            )}

            {/* Sub-Tab 2: VNet & Subnets */}
            {activeSubTab === 'vnets' && (
              <div className="space-y-4">
                <h4 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                  <Network className="w-4 h-4 text-emerald-500" />
                  VNet & Subnet Configurations
                </h4>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  {/* Virtual Network Info */}
                  <div className="glass-panel p-4 rounded-xl border border-border-main/60 space-y-3">
                    <span className="text-xs font-bold text-text-muted font-mono block">Virtual Network (VNet)</span>
                    <div className="space-y-1.5 text-xs font-mono">
                      <div className="flex justify-between py-1 border-b border-border-main/20">
                        <span className="text-text-muted">VNet Name:</span>
                        <span className="font-extrabold text-emerald-600 dark:text-emerald-400">{extractedVNetName}</span>
                      </div>
                      <div className="flex justify-between py-1 border-b border-border-main/20">
                        <span className="text-text-muted">CIDR Block:</span>
                        <span className="font-bold text-text-main">{currentCloudModel.targetVNet?.cidrBlock || '10.0.0.0/16'}</span>
                      </div>
                      <div className="flex justify-between py-1">
                        <span className="text-text-muted">Target Region:</span>
                        <span className="font-bold text-text-main">{currentCloudModel.targetCloud?.region || 'ap-northeast-2'}</span>
                      </div>
                    </div>
                  </div>

                  {/* Subnets List */}
                  <div className="glass-panel p-4 rounded-xl border border-border-main/60 space-y-3">
                    <span className="text-xs font-bold text-text-muted font-mono block">Subnets Allocation</span>
                    <div className="space-y-2 text-xs font-mono">
                      {(currentCloudModel.targetVNet?.subnetInfoList || [{ name: extractedSubnetName, ipv4_CIDR: '10.0.1.0/24', cidrBlock: '10.0.1.0/24' }]).map((sub: any, idx: number) => {
                        const subCidr = sub.ipv4_CIDR || sub.cidrBlock || '10.0.1.0/24';
                        const subNgCount = (currentCloudModel.targetInfra?.nodeGroups || []).filter(
                          (ng: any) => ng.subnetId === sub.name
                        ).length;
                        return (
                          <div key={idx} className="p-2.5 bg-bg-panel/60 rounded-lg border border-border-main/40 flex justify-between items-center">
                            <div>
                              <span className="font-extrabold text-text-main block">{sub.name}</span>
                              <span className="text-text-muted text-xs">Subnet ID: {sub.name} ({subNgCount} Node Group(s))</span>
                            </div>
                            <span className="px-2 py-0.5 rounded bg-emerald-500/10 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 font-bold">
                              {subCidr}
                            </span>
                          </div>
                        );
                      })}
                    </div>
                  </div>
                </div>
              </div>
            )}

            {/* Sub-Tab 3: Security Groups */}
            {activeSubTab === 'sgs' && (
              <div className="space-y-4">
                <h4 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                  <Shield className="w-4 h-4 text-amber-500" />
                  Security Group Firewall Rules
                </h4>

                <div className="overflow-x-auto border border-border-main/50 rounded-xl">
                  <table className="w-full text-left border-collapse text-xs font-mono">
                    <thead>
                      <tr className="border-b border-border-main bg-bg-input/60 text-text-muted font-bold">
                        <th className="py-2.5 px-4">Security Group Name</th>
                        <th className="py-2.5 px-4">Direction</th>
                        <th className="py-2.5 px-4">Protocol</th>
                        <th className="py-2.5 px-4">Port Range</th>
                        <th className="py-2.5 px-4">Source CIDR</th>
                      </tr>
                    </thead>
                    <tbody className="divide-y divide-border-main/40 text-text-main">
                      {targetSecurityGroupList.map((sg: any, idx: number) => (
                        <React.Fragment key={idx}>
                          {(sg.firewallRules || [{ protocol: 'tcp', dstPorts: '22, 80, 443' }]).map((rule: any, rIdx: number) => (
                            <tr key={rIdx} className="hover:bg-emerald-500/[0.02] transition">
                              <td className="py-2.5 px-4 font-bold text-amber-600 dark:text-amber-400">{sg.name}</td>
                              <td className="py-2.5 px-4">
                                <span className="px-2 py-0.5 rounded bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 font-extrabold">
                                  INBOUND
                                </span>
                              </td>
                              <td className="py-2.5 px-4 font-bold uppercase">{rule.protocol || 'tcp'}</td>
                              <td className="py-2.5 px-4 font-extrabold text-emerald-600 dark:text-emerald-400">{rule.dstPorts || rule.fromPort || '22, 80, 443'}</td>
                              <td className="py-2.5 px-4 text-text-muted">{rule.cidr || '0.0.0.0/0'}</td>
                            </tr>
                          ))}
                        </React.Fragment>
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            )}

            {/* Sub-Tab 4: SSH Key */}
            {activeSubTab === 'ssh' && (
              <div className="space-y-4">
                <h4 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                  <Key className="w-4 h-4 text-teal-400" />
                  SSH Access Key Credentials
                </h4>

                <div className="glass-panel p-5 rounded-xl border border-border-main/60 space-y-4">
                  <div className="grid grid-cols-1 md:grid-cols-3 gap-4 border-b border-border-main/20 pb-4 text-xs font-mono">
                    <div>
                      <span className="text-text-muted block font-bold">Key Pair Identifier</span>
                      <span className="text-sm font-extrabold text-teal-600 dark:text-teal-400">{extractedSshKeyName}</span>
                    </div>
                    <div>
                      <span className="text-text-muted block font-bold">Key Algorithm</span>
                      <span className="font-bold text-text-main">RSA 4096-bit</span>
                    </div>
                    <div>
                      <span className="text-text-muted block font-bold">Default System User</span>
                      <span className="font-bold text-emerald-600 dark:text-emerald-400">cb-user</span>
                    </div>
                  </div>

                  <div className="space-y-1 text-xs font-mono">
                    <span className="text-text-muted font-bold block">SSH Private Key Local Path:</span>
                    <code className="p-2.5 bg-bg-input rounded-lg border border-border-main block select-all text-emerald-600 dark:text-emerald-400">
                      ~/.ssh/{extractedSshKeyName}.pem
                    </code>
                  </div>
                </div>
              </div>
            )}

            {/* Sub-Tab 5: Managed NLB */}
            {activeSubTab === 'nlb' && (() => {
              const displayNlbs = liveNlbs.length > 0
                ? liveNlbs
                : ((currentCloudModel as any).targetNlbList || []);

              return (
                <div className="space-y-4">
                  <div className="flex items-center justify-between">
                    <h4 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                      <Radio className="w-4 h-4 text-emerald-500" />
                      Managed Network Load Balancers (NLB)
                    </h4>
                    {displayNlbs.length > 0 && (
                      <span className="text-xs font-mono text-text-muted">
                        Total: <strong className="text-text-main">{displayNlbs.length}</strong> active NLB(s)
                      </span>
                    )}
                  </div>

                  {displayNlbs.length > 0 ? (
                    <div className="overflow-x-auto border border-border-main/50 rounded-xl">
                      <table className="w-full text-left border-collapse text-xs font-mono">
                        <thead>
                          <tr className="border-b border-border-main bg-bg-input/60 text-text-muted font-bold">
                            <th className="py-2.5 px-4">NLB Name</th>
                            <th className="py-2.5 px-4">Listener Port</th>
                            <th className="py-2.5 px-4">Target Protocol</th>
                            <th className="py-2.5 px-4">Health Check</th>
                            <th className="py-2.5 px-4">Target Node Group</th>
                            <th className="py-2.5 px-4 text-right">Actions</th>
                          </tr>
                        </thead>
                        <tbody className="divide-y divide-border-main/40 text-text-main">
                          {displayNlbs.map((nlb: any, idx: number) => {
                            const nlbName = nlb.name || nlb.id || `mig-nlb-${idx + 1}`;
                            const port = nlb.listener?.port ? `Port ${nlb.listener.port}` : 'Port 80 ➔ 8080';
                            const protocol = nlb.listener?.protocol || 'TCP / HTTP';
                            const targetGroup = nlb.targetGroup?.nodeGroupId || 'ng-web-01';
                            const isDeletingNlb = deletingNlbId === nlbName;

                            return (
                              <tr key={idx} className="hover:bg-emerald-500/[0.02] transition">
                                <td className="py-3 px-4 font-bold text-emerald-600 dark:text-emerald-400">{nlbName}</td>
                                <td className="py-3 px-4 font-bold">{port}</td>
                                <td className="py-3 px-4 text-emerald-500 font-extrabold">{protocol}</td>
                                <td className="py-3 px-4 text-text-muted">HTTP /healthz (Interval: 10s)</td>
                                <td className="py-3 px-4 font-bold text-text-main">{targetGroup}</td>
                                <td className="py-3 px-4 text-right">
                                  <button
                                    onClick={() => handleDeleteSingleNlb(nlbName)}
                                    disabled={isDeletingNlb || isDeleting}
                                    className="px-2.5 py-1 text-red-500 hover:bg-red-500/10 border border-red-500/30 rounded-lg transition font-mono font-bold inline-flex items-center gap-1 cursor-pointer disabled:opacity-40"
                                    title="Delete this NLB individually"
                                  >
                                    {isDeletingNlb ? (
                                      <Loader2 className="w-3.5 h-3.5 animate-spin" />
                                    ) : (
                                      <Trash2 className="w-3.5 h-3.5" />
                                    )}
                                    <span>Delete NLB</span>
                                  </button>
                                </td>
                              </tr>
                            );
                          })}
                        </tbody>
                      </table>
                    </div>
                  ) : (
                    <div className="p-6 bg-bg-panel/40 border border-border-main/40 rounded-xl text-center text-xs font-mono text-text-muted">
                      No Dedicated Managed NLB configured for this infrastructure. Compute Nodes are configured with Direct Public IP access.
                    </div>
                  )}
                </div>
              );
            })()}
          </div>

          {/* Post-Migration Report & Target Summary Viewer (Multi-Format HTML / MD / JSON) */}
          {(selectedInfraId !== SAMPLE_INFRA_ID || reportContent || infraReportHtml || liveReportHtml) && (
            <div className="bg-white rounded-2xl border border-slate-200 shadow-sm overflow-hidden flex flex-col">
              <div className="px-6 py-4 bg-white border-b border-slate-200 flex flex-wrap items-center justify-between gap-3">
                <div className="flex flex-wrap items-center gap-3">
                  <h3 className="text-sm font-extrabold text-slate-800 flex items-center gap-2">
                    <FileText className="w-4 h-4 text-emerald-500" />
                    <span>Migration Report &amp; Target Infra Summary</span>
                  </h3>

                  {/* Document Type Switcher: Migration Report vs Target Summary */}
                  <div className="flex items-center bg-slate-100 p-0.5 rounded-lg border border-slate-300 text-xs font-mono">
                    <button
                      type="button"
                      onClick={() => {
                        setReportDocType('migration');
                        fetchReportOrSummary(selectedInfraId, 'migration', reportFormat);
                      }}
                      className={`px-2.5 py-1 rounded-md font-bold transition cursor-pointer ${
                        reportDocType === 'migration'
                          ? 'bg-emerald-500 text-slate-950 shadow-sm'
                          : 'text-slate-600 hover:text-slate-900'
                      }`}
                    >
                      Migration Report (POST /report/migration)
                    </button>
                    <button
                      type="button"
                      onClick={() => {
                        setReportDocType('target-summary');
                        fetchReportOrSummary(selectedInfraId, 'target-summary', reportFormat);
                      }}
                      className={`px-2.5 py-1 rounded-md font-bold transition cursor-pointer ${
                        reportDocType === 'target-summary'
                          ? 'bg-emerald-500 text-slate-950 shadow-sm'
                          : 'text-slate-600 hover:text-slate-900'
                      }`}
                    >
                      Target Summary (GET /summary/target)
                    </button>
                  </div>

                  {/* Format Switcher: HTML / MD / JSON */}
                  <div className="flex items-center bg-slate-100 p-0.5 rounded-lg border border-slate-300 text-xs font-mono">
                    {(['html', 'md', 'json'] as const).map((fmt) => (
                      <button
                        key={fmt}
                        type="button"
                        onClick={() => {
                          setReportFormat(fmt);
                          fetchReportOrSummary(selectedInfraId, reportDocType, fmt);
                        }}
                        className={`px-2.5 py-1 rounded-md font-extrabold uppercase transition cursor-pointer ${
                          reportFormat === fmt
                            ? 'bg-teal-500 text-slate-950 shadow-sm'
                            : 'text-slate-600 hover:text-slate-900'
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
                    onClick={() => fetchReportOrSummary(selectedInfraId, reportDocType, reportFormat)}
                    disabled={isLoadingReport || selectedInfraId === SAMPLE_INFRA_ID}
                    className="px-3 py-1.5 bg-slate-100 border border-slate-300 hover:bg-emerald-50 text-slate-700 rounded-lg text-xs font-bold flex items-center gap-1 transition cursor-pointer disabled:opacity-40"
                  >
                    <RefreshCw className={`w-3.5 h-3.5 ${isLoadingReport ? 'animate-spin' : ''}`} />
                    <span>Refresh</span>
                  </button>
                  <button
                    type="button"
                    onClick={handleDownloadReport}
                    disabled={!(reportContent || infraReportHtml || liveReportHtml)}
                    className="px-3.5 py-1.5 bg-slate-100 border border-slate-300 hover:bg-emerald-50 hover:border-emerald-500/30 text-emerald-700 rounded-lg text-xs font-extrabold flex items-center transition cursor-pointer disabled:opacity-40"
                  >
                    <Download className="w-3.5 h-3.5 mr-1" />
                    <span>Download ({reportFormat.toUpperCase()})</span>
                  </button>
                </div>
              </div>

              {isLoadingReport ? (
                <div className="p-12 text-center bg-white space-y-2">
                  <Loader2 className="w-6 h-6 animate-spin text-emerald-500 mx-auto" />
                  <p className="text-xs font-mono text-slate-600">Generating {reportDocType === 'migration' ? 'Migration Report' : 'Target Summary'} ({reportFormat.toUpperCase()})...</p>
                </div>
              ) : reportFormat === 'html' ? (
                (reportContent || infraReportHtml || liveReportHtml) ? (
                  <div className="w-full h-[1650px] bg-white rounded-b-2xl overflow-hidden">
                    <iframe
                      title="Post-Migration Compliance Report"
                      srcDoc={(reportContent || infraReportHtml || liveReportHtml)
                        .replace(/background-color:\s*#f5f5f5/gi, 'background-color: #ffffff')
                        .replace(/background:\s*#f5f5f5/gi, 'background: #ffffff')}
                      className="w-full h-full border-0 bg-white"
                      sandbox="allow-same-origin"
                    />
                  </div>
                ) : (
                  <div className="p-8 text-center text-xs font-mono text-slate-600 bg-white">
                    Click &quot;Load Detail&quot; or &quot;Refresh&quot; on a migrated infrastructure to view the report.
                  </div>
                )
              ) : (
                <div className="p-6 bg-slate-50 max-h-[850px] overflow-y-auto">
                  <pre className="text-xs font-mono text-slate-800 whitespace-pre-wrap break-words leading-relaxed">
                    {reportContent || 'No report content fetched yet. Click Refresh to generate.'}
                  </pre>
                </div>
              )}
            </div>
          )}

        </div>
      ) : (
        /* Empty State */
        <div className="glass-panel p-12 rounded-2xl text-center space-y-3 border border-border-main">
          <Server className="w-12 h-12 text-text-muted mx-auto opacity-50" />
          <h3 className="text-base font-extrabold text-text-main">No Migrated Infrastructure Selected</h3>
          <p className="text-sm text-text-muted max-w-md mx-auto">
            Choose a migrated cloud infrastructure from the dropdown above to view its live topology map, VM public/private IP addresses, and compliance reports.
          </p>
        </div>
      )}
      </div>
      )}

      {overviewCategory === 'storage' && (() => {
        const activeSelectedStorage = (loadedStorageDetail?.name || loadedStorageDetail?.id) === selectedStorageId
          ? loadedStorageDetail
          : (migratedStorages.find((s: any) => (s.name || s.bucketName || s.id) === selectedStorageId) || migratedStorages[0] || SAMPLE_STORAGE_DETAIL);

        return (
          <div className="space-y-6 animate-fade-in">
            {/* 4. Object Storages Section */}
            <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-4">
              <div className="flex flex-wrap justify-between items-center gap-3 border-b border-border-main/20 pb-3">
                <h3 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                  <HardDrive className="w-4 h-4 text-emerald-500" />
                  <span>Object Storages ({migratedStorages.length})</span>
                </h3>

                <div className="flex items-center gap-2">
                  {/* View Switcher Toggle (Cards vs Table) */}
                  <div className="flex items-center bg-bg-panel/80 p-0.5 rounded-lg border border-border-main">
                    <button
                      onClick={() => setStorageCatalogViewMode('grid')}
                      className={`p-1.5 px-2.5 rounded-md transition cursor-pointer text-xs flex items-center gap-1 font-bold font-mono ${
                        storageCatalogViewMode === 'grid'
                          ? 'bg-emerald-500 text-slate-950 shadow-sm'
                          : 'text-text-muted hover:text-text-main'
                      }`}
                      title="Cards Grid View"
                    >
                      <LayoutGrid className="w-3.5 h-3.5" />
                      <span>Cards</span>
                    </button>
                    <button
                      onClick={() => setStorageCatalogViewMode('table')}
                      className={`p-1.5 px-2.5 rounded-md transition cursor-pointer text-xs flex items-center gap-1 font-bold font-mono ${
                        storageCatalogViewMode === 'table'
                          ? 'bg-emerald-500 text-slate-950 shadow-sm'
                          : 'text-text-muted hover:text-text-main'
                      }`}
                      title="Compact Table Row View"
                    >
                      <List className="w-3.5 h-3.5" />
                      <span>Table</span>
                    </button>
                  </div>

                  <button
                    onClick={loadMigratedStorages}
                    disabled={isLoadingStorages}
                    className="px-3 py-1.5 bg-bg-panel hover:bg-border-main/20 text-emerald-500 rounded-lg border border-border-main text-xs font-bold transition cursor-pointer flex items-center gap-1.5 font-mono"
                    title="Refresh object storages list"
                  >
                    <RefreshCw className={`w-3.5 h-3.5 ${isLoadingStorages ? 'animate-spin' : ''}`} />
                    <span>Refresh</span>
                  </button>
                </div>
              </div>

              {migratedStorages.length > 0 ? (
                storageCatalogViewMode === 'grid' ? (
                  /* Cards Grid View */
                  <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-3">
                    {migratedStorages.map((storage, idx) => {
                      const sName = storage.name || storage.bucketName || storage.id || `storage-${idx + 1}`;
                      const isSelected = (selectedStorageId === sName) || (selectedStorageId === '' && idx === 0);
                      const isDeleting = !!deletingStoragesMap[sName];
                      const isProvisioning = objectStorageJobs.some(j => (j.name === sName || j.id === sName || j.reqId === sName) && j.status === 'Handling');
                      const connInfo = extractCspAndRegionFromConnection(storage.connectionName);
                      const csp = (storage.csp || storage.targetCloud?.csp || connInfo.csp).toUpperCase();
                      const region = storage.region || storage.targetCloud?.region || connInfo.region;

                      return (
                        <div
                          key={storage.id || idx}
                          onClick={() => setSelectedStorageId(sName)}
                          className={`p-4 rounded-xl border text-left transition-all duration-200 flex flex-col justify-between space-y-3 cursor-pointer relative overflow-hidden ${
                            isDeleting || isProvisioning
                              ? 'bg-amber-500/5 border-amber-500/40'
                              : isSelected
                              ? 'bg-emerald-500/10 border-emerald-500/60 shadow-lg shadow-emerald-500/10 ring-1 ring-emerald-500/40'
                              : 'bg-bg-panel/40 border-border-main/40 hover:bg-bg-panel/80 hover:border-border-main'
                          }`}
                        >
                          <div className="flex justify-between items-start gap-2">
                            <div className="space-y-1 min-w-0">
                              <div className="flex items-center gap-1.5 flex-wrap">
                                {sName === SAMPLE_STORAGE_ID && (
                                  <span className="text-xs font-extrabold px-2 py-0.5 rounded-md bg-amber-500/20 text-amber-600 dark:text-amber-300 border border-amber-500/40 font-mono shadow-sm">
                                    [Sample]
                                  </span>
                                )}
                                <span className="text-xs font-extrabold px-2 py-0.5 rounded bg-emerald-500/10 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 font-mono">
                                  {csp}
                                </span>
                                <span className="text-xs font-bold text-text-muted font-mono truncate max-w-[120px]">
                                  {region}
                                </span>
                              </div>
                              <h4 className="font-extrabold text-sm text-text-main truncate max-w-[200px]" title={sName}>
                                {sName}
                              </h4>
                            </div>

                            {isDeleting ? (
                              <span className="text-xs font-bold px-2 py-0.5 rounded-full bg-amber-500/10 text-amber-500 border border-amber-500/30 flex items-center gap-1 font-mono shrink-0">
                                <span className="w-1.5 h-1.5 bg-amber-500 rounded-full animate-ping" />
                                Deleting...
                              </span>
                            ) : isProvisioning ? (
                              <span className="text-xs font-bold px-2 py-0.5 rounded-full bg-amber-500/10 text-amber-500 border border-amber-500/30 flex items-center gap-1 font-mono shrink-0">
                                <span className="w-1.5 h-1.5 bg-amber-500 rounded-full animate-ping" />
                                Migrating...
                              </span>
                            ) : (
                              <span className="text-xs font-bold px-2 py-0.5 rounded-full bg-emerald-100 dark:bg-emerald-950/50 text-emerald-600 dark:text-emerald-400 border border-emerald-300 dark:border-emerald-800/40 flex items-center gap-1 font-mono shrink-0">
                                <span className="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse" />
                                Available
                              </span>
                            )}
                          </div>

                          <div className="flex items-center justify-between gap-2 pt-2 border-t border-border-main/20">
                            <button
                              onClick={(e) => {
                                e.stopPropagation();
                                setSelectedStorageId(sName);
                              }}
                              disabled={isDeleting}
                              className={`flex-1 py-1.5 px-2.5 rounded-lg text-xs font-bold flex items-center justify-center gap-1.5 transition cursor-pointer font-mono ${
                                isDeleting
                                  ? 'bg-bg-input text-text-muted border border-border-main opacity-50 cursor-not-allowed'
                                  : isSelected
                                  ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
                                  : 'bg-bg-input hover:bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-border-main'
                              }`}
                            >
                              <Eye className="w-3.5 h-3.5" />
                              <span>{isSelected ? 'Loaded ✓' : 'Load Detail'}</span>
                            </button>

                            <button
                               onClick={(e) => {
                                 e.stopPropagation();
                                 handleOpenStorageDeleteModal(sName);
                               }}
                               disabled={isDeleting || sName === SAMPLE_STORAGE_ID}
                               className="py-1.5 px-2.5 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-lg text-xs font-bold flex items-center justify-center gap-1 transition cursor-pointer font-mono disabled:opacity-30 disabled:cursor-not-allowed"
                               title={sName === SAMPLE_STORAGE_ID ? 'Sample object storage cannot be deleted' : isDeleting ? 'Storage deletion in progress' : `Delete object storage ${sName}`}
                             >
                               {isDeleting ? (
                                 <>
                                   <Loader2 className="w-3.5 h-3.5 animate-spin text-amber-500" />
                                   <span>Deleting...</span>
                                 </>
                               ) : (
                                 <>
                                   <Trash2 className="w-3.5 h-3.5" />
                                   <span>Delete</span>
                                 </>
                               )}
                             </button>
                          </div>
                        </div>
                      );
                    })}
                  </div>
                ) : (
                  /* Compact Table Row View */
                  <div className="overflow-x-auto rounded-xl border border-border-main/40 bg-bg-panel/20">
                    <table className="w-full text-left text-xs font-mono">
                      <thead>
                        <tr className="border-b border-border-main/40 bg-bg-panel/80 text-text-muted font-normal">
                          <th className="py-2.5 px-4 w-12 text-center font-normal">#</th>
                          <th className="py-2.5 px-4 font-normal">Object Storage Name</th>
                          <th className="py-2.5 px-4 font-normal">Target CSP / Region</th>
                          <th className="py-2.5 px-4 font-normal">Source Bucket</th>
                          <th className="py-2.5 px-4 font-normal">Status</th>
                          <th className="py-2.5 px-4 text-right font-normal">Actions</th>
                        </tr>
                      </thead>
                      <tbody className="divide-y divide-border-main/20">
                        {migratedStorages.map((storage, idx) => {
                          const sName = storage.name || storage.bucketName || storage.id || `storage-${idx + 1}`;
                          const isSelected = (selectedStorageId === sName) || (selectedStorageId === '' && idx === 0);
                          const isDeleting = !!deletingStoragesMap[sName];
                          const connInfo = extractCspAndRegionFromConnection(storage.connectionName);
                          const csp = (storage.csp || storage.targetCloud?.csp || connInfo.csp).toUpperCase();
                          const region = storage.region || storage.targetCloud?.region || connInfo.region;
                          const srcBucket = storage.sourceBucketName || 'datamold-aws-test';

                          return (
                            <tr
                              key={storage.id || idx}
                              onClick={() => setSelectedStorageId(sName)}
                              className={`transition hover:bg-bg-panel/60 cursor-pointer ${
                                isDeleting
                                  ? 'bg-amber-500/5'
                                  : isSelected
                                  ? 'bg-emerald-500/10 font-bold'
                                  : ''
                              }`}
                            >
                              <td className="py-3 px-4 text-center text-text-muted">{idx + 1}</td>
                              <td className="py-3 px-4 font-extrabold text-text-main flex items-center gap-2">
                                <HardDrive className="w-4 h-4 text-emerald-500 shrink-0" />
                                {sName === SAMPLE_STORAGE_ID && (
                                  <span className="text-xs font-extrabold px-2 py-0.5 rounded-md bg-amber-500/20 text-amber-600 dark:text-amber-300 border border-amber-500/40 font-mono shadow-sm">
                                    [Sample]
                                  </span>
                                )}
                                <span>{sName}</span>
                              </td>
                              <td className="py-3 px-4 text-text-main font-bold">
                                {csp} <span className="text-text-muted font-normal">({region})</span>
                              </td>
                              <td className="py-3 px-4 text-text-muted font-normal">
                                {srcBucket}
                              </td>
                              <td className="py-3 px-4">
                                {isDeleting ? (
                                  <span className="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full text-xs font-extrabold bg-amber-500/10 text-amber-500 border border-amber-500/30">
                                    <span className="w-1.5 h-1.5 rounded-full bg-amber-500 animate-ping"></span>
                                    <span>Deleting...</span>
                                  </span>
                                ) : (
                                  <span className="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full text-xs font-extrabold bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-emerald-500/20">
                                    <span className="w-1.5 h-1.5 rounded-full bg-emerald-500 animate-pulse"></span>
                                    <span>Available</span>
                                  </span>
                                )}
                              </td>
                              <td className="py-3 px-4 text-right">
                                <div className="flex items-center justify-end gap-2">
                                  <button
                                    onClick={(e) => {
                                      e.stopPropagation();
                                      setSelectedStorageId(sName);
                                    }}
                                    disabled={isDeleting}
                                    className={`py-1 px-3 rounded-lg text-xs font-bold flex items-center gap-1.5 transition cursor-pointer ${
                                      isDeleting
                                        ? 'bg-bg-input text-text-muted border border-border-main opacity-50 cursor-not-allowed'
                                        : isSelected
                                        ? 'bg-emerald-500 text-slate-950 shadow-sm'
                                        : 'bg-bg-input hover:bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-border-main'
                                    }`}
                                  >
                                    <Eye className="w-3.5 h-3.5" />
                                    <span>{isSelected ? 'Loaded ✓' : 'Load Detail'}</span>
                                  </button>

                                  <button
                                     onClick={(e) => {
                                       e.stopPropagation();
                                       handleOpenStorageDeleteModal(sName);
                                     }}
                                     disabled={isDeleting || sName === SAMPLE_STORAGE_ID}
                                     className="py-1 px-3 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-lg text-xs font-bold transition cursor-pointer disabled:opacity-30 disabled:cursor-not-allowed"
                                     title={sName === SAMPLE_STORAGE_ID ? 'Sample object storage cannot be deleted' : isDeleting ? 'Storage deletion in progress' : `Delete object storage ${sName}`}
                                   >
                                     {isDeleting ? (
                                       <>
                                         <Loader2 className="w-3.5 h-3.5 inline mr-1 animate-spin text-amber-500" />
                                         <span>Deleting...</span>
                                       </>
                                     ) : (
                                       <>
                                         <Trash2 className="w-3.5 h-3.5 inline mr-1" />
                                         <span>Delete</span>
                                       </>
                                     )}
                                   </button>
                                </div>
                              </td>
                            </tr>
                          );
                        })}
                      </tbody>
                    </table>
                  </div>
                )
              ) : (
                <div className="py-8 text-center text-text-muted text-xs font-mono italic">
                  No deployed object storages found.
                </div>
              )}
            </div>

            {/* Detailed Storage Information Block */}
            {activeSelectedStorage && (
              <div className="space-y-6 animate-fade-in">
                {/* Resource Metrics Bar (4 API Resource Cards) */}
                <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-5 font-sans">
                  {/* 1. Target CSP & Region */}
                  <div className="glass-panel p-5 rounded-2xl border border-border-main/60 space-y-2 flex flex-col justify-center">
                    <span className="text-xs font-normal text-text-muted">Target CSP &amp; Region</span>
                    <div className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 flex items-center gap-2 truncate">
                      <Cloud className="w-4 h-4 shrink-0 text-emerald-500" />
                      <span className="truncate">{(activeSelectedStorage.connectionConfig?.providerName || activeSelectedStorage.csp || 'AWS').toUpperCase()} ({activeSelectedStorage.connectionConfig?.regionZoneInfo?.assignedRegion || activeSelectedStorage.region || 'ap-northeast-2'})</span>
                    </div>
                  </div>

                  {/* 2. Object Storage Name */}
                  <div className="glass-panel p-5 rounded-2xl border border-border-main/60 space-y-2 flex flex-col justify-center">
                    <span className="text-xs font-normal text-text-muted">Object Storage Name</span>
                    <div className="text-sm font-extrabold text-text-main flex items-center gap-2 truncate font-mono">
                      <HardDrive className="w-4 h-4 shrink-0 text-emerald-500" />
                      <span className="truncate">{activeSelectedStorage.name || activeSelectedStorage.id}</span>
                    </div>
                  </div>

                  {/* 3. Objects Count */}
                  <div className="glass-panel p-5 rounded-2xl border border-border-main/60 space-y-2 flex flex-col justify-center">
                    <span className="text-xs font-normal text-text-muted">Total Objects</span>
                    <div className="text-sm font-extrabold text-text-main flex items-center gap-2 truncate">
                      <Database className="w-4 h-4 shrink-0 text-emerald-500" />
                      <span className="truncate">{Array.isArray(activeSelectedStorage.contents) ? activeSelectedStorage.contents.length : 0} Object(s)</span>
                    </div>
                  </div>

                  {/* 4. Lifecycle Status */}
                  <div className="glass-panel p-5 rounded-2xl border border-border-main/60 space-y-2 flex flex-col justify-center">
                    <span className="text-xs font-normal text-text-muted">Lifecycle Status</span>
                    <div className="text-sm font-extrabold text-emerald-600 dark:text-emerald-400 flex items-center gap-2">
                      <span className="w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
                      <span>{activeSelectedStorage.status || 'Available'}</span>
                    </div>
                  </div>
                </div>

                {/* Detailed Spec Configuration Panel */}
                <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-6">
                  <div className="flex items-center justify-between border-b border-border-main/30 pb-3">
                    <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
                      <Sliders className="w-5 h-5 text-emerald-500" />
                      <span>Object Storage Identifiers &amp; Spec Details</span>
                    </h3>
                    <span className="text-sm font-mono text-text-muted">
                      Storage ID: <strong className="text-emerald-500 font-extrabold">{activeSelectedStorage.name || activeSelectedStorage.id}</strong>
                    </span>
                  </div>

                  {/* Resource Identifiers Grid (API Returned Fields Only) */}
                  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5 text-sm">
                    {/* Object Storage Name (Tumblebug Managed) */}
                    <div className="p-4 bg-bg-panel/60 border border-border-main rounded-xl space-y-1.5">
                      <span className="text-text-muted font-normal block">Object Storage Name (Tumblebug)</span>
                      <div className="font-extrabold text-text-main font-mono break-all text-sm">
                        {activeSelectedStorage.name || activeSelectedStorage.id}
                      </div>
                    </div>

                    {/* CSP Bucket Name (Assigned via UID) */}
                    <div className="p-4 bg-bg-panel/60 border border-border-main rounded-xl space-y-1.5">
                      <span className="text-text-muted font-normal block">CSP Bucket Name (UID Assigned)</span>
                      <div className="font-extrabold text-emerald-500 font-mono break-all text-sm">
                        {activeSelectedStorage.uid || activeSelectedStorage.id}
                      </div>
                    </div>

                    {/* Connection Config & Driver */}
                    <div className="p-4 bg-bg-panel/60 border border-border-main rounded-xl space-y-1.5">
                      <span className="text-text-muted font-normal block">Connection &amp; Driver</span>
                      <div className="font-extrabold text-text-main font-mono text-sm">
                        {activeSelectedStorage.connectionConfig?.configName || activeSelectedStorage.connectionName || 'aws-ap-northeast-2'}
                      </div>
                    </div>

                    {/* Credential & Holder */}
                    <div className="p-4 bg-bg-panel/60 border border-border-main rounded-xl space-y-1.5">
                      <span className="text-text-muted font-normal block">Credential &amp; Holder</span>
                      <div className="font-extrabold text-text-main font-mono text-sm">
                        {activeSelectedStorage.connectionConfig?.credentialName || 'aws-credential'} <span className="text-text-muted font-normal">({activeSelectedStorage.connectionConfig?.credentialHolder || 'admin'})</span>
                      </div>
                    </div>

                    {/* Resource Creation Timestamp */}
                    <div className="p-4 bg-bg-panel/60 border border-border-main rounded-xl space-y-1.5">
                      <span className="text-text-muted font-normal block">Creation Date</span>
                      <div className="font-extrabold text-text-main font-mono text-sm">
                        {activeSelectedStorage.creationDate || 'N/A'}
                      </div>
                    </div>

                    {/* Description */}
                    <div className="p-4 bg-bg-panel/60 border border-border-main rounded-xl space-y-1.5">
                      <span className="text-text-muted font-normal block">Description</span>
                      <div className="font-extrabold text-text-main font-mono text-sm truncate">
                        {activeSelectedStorage.description || 'Created by CM-Beetle'}
                      </div>
                    </div>
                  </div>

                  {/* Bucket Objects List (Directory Tree View) */}
                  {Array.isArray(activeSelectedStorage.contents) && activeSelectedStorage.contents.length > 0 && (
                    <div className="space-y-3 pt-4 border-t border-border-main/30">
                      <div className="flex items-center justify-between">
                        <h4 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                          <Database className="w-4 h-4 text-emerald-500" />
                          <span>Bucket Contents &amp; Objects Directory ({activeSelectedStorage.contents.length})</span>
                        </h4>
                      </div>

                      <ObjectDirectoryTree
                        bucketName={activeSelectedStorage.name || activeSelectedStorage.id || 'target-bucket'}
                        objects={activeSelectedStorage.contents}
                      />
                    </div>
                  )}
                </div>
              </div>
            )}
          </div>
        );
      })()}

      {/* Delete Confirmation Modal */}
      {showDeleteConfirm && selectedInfraId && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up space-y-4">
            <div className="flex justify-between items-center border-b border-border-main/20 pb-3">
              <h3 className="text-base font-bold text-text-main flex items-center gap-2">
                <Trash2 className="w-5 h-5 text-red-500" />
                Delete Migrated Infrastructure
              </h3>
              <button
                onClick={() => setShowDeleteConfirm(false)}
                disabled={isDeleting}
                className="text-text-muted hover:text-text-main transition p-1 rounded-lg cursor-pointer"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            <p className="text-sm text-text-muted leading-relaxed">
              Are you sure you want to delete <strong className="text-text-main">"{selectedInfraId}"</strong>? This will release all cloud VMs, subnets, and VPC network allocations.
            </p>

            {/* Staged Teardown Notice when NLB is detected */}
            {detectedNlbs.length > 0 && (
              <div className="p-3.5 bg-amber-500/10 border border-amber-500/30 rounded-xl space-y-2 text-xs">
                <div className="flex items-center gap-1.5 font-bold text-amber-500">
                  <AlertTriangle className="w-4 h-4 shrink-0" />
                  <span>Attached Managed NLB Detected ({detectedNlbs.map(n => n.name || n.id).join(', ')})</span>
                </div>
                <p className="text-text-muted leading-relaxed">
                  To prevent cloud dependency locks (ENI release), Beetle UX Lab will orchestrate teardown in 2 sequential stages:
                </p>
                <div className="space-y-1 font-mono text-xs text-text-main">
                  <div className="flex items-center gap-2">
                    <span className="px-1.5 py-0.5 rounded bg-amber-500/20 text-amber-600 dark:text-amber-400 font-bold shrink-0">Stage 1</span>
                    <span>Terminate & Detach Managed NLB ({detectedNlbs.map(n => n.name || n.id).join(', ')})</span>
                  </div>
                  <div className="flex items-center gap-2">
                    <span className="px-1.5 py-0.5 rounded bg-emerald-500/20 text-emerald-600 dark:text-emerald-400 font-bold shrink-0">Stage 2</span>
                    <span>Terminate Compute Nodes, Security Groups & VNet</span>
                  </div>
                </div>
              </div>
            )}

            <div className="space-y-1.5">
              <label className="block text-xs font-bold text-text-muted">
                To confirm, type <span className="font-mono bg-bg-panel px-1 py-0.5 rounded text-text-main">{selectedInfraId}</span> in the box below:
              </label>
              <input
                type="text"
                value={deleteConfirmText}
                onChange={(e) => setDeleteConfirmText(e.target.value)}
                placeholder="Type the infra name to delete"
                className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm font-bold font-mono focus:outline-none focus:ring-1 focus:ring-red-500"
                disabled={isDeleting}
              />
            </div>

            {deleteError && (
              <div className="p-3 bg-red-500/10 border border-red-500/20 text-red-500 text-xs font-semibold rounded-xl">
                {deleteError}
              </div>
            )}

            <div className="flex justify-end gap-3 pt-2">
              <button
                onClick={() => setShowDeleteConfirm(false)}
                disabled={isDeleting}
                className="px-4 py-2 bg-bg-panel border border-border-main text-text-main rounded-xl text-sm font-semibold hover:bg-bg-input transition cursor-pointer"
              >
                Cancel
              </button>
              <button
                onClick={handleConfirmDelete}
                disabled={isDeleting || deleteConfirmText !== selectedInfraId}
                className={`px-4 py-2 rounded-xl text-sm font-bold flex items-center gap-1.5 transition ${
                  isDeleting || deleteConfirmText !== selectedInfraId
                    ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                    : 'bg-red-500 hover:bg-red-600 text-white cursor-pointer shadow-lg shadow-red-500/20'
                }`}
              >
                {isDeleting && <Loader2 className="w-4 h-4 animate-spin" />}
                {deleteStage === 'deleting-nlb'
                  ? 'Step 1/2: Detaching NLB...'
                  : deleteStage === 'deleting-infra'
                  ? 'Step 2/2: Terminating Infra...'
                  : 'Confirm Delete'}
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Object Storage Delete Confirmation Modal */}
      {showStorageDeleteConfirm && targetStorageToDelete && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-3xl border border-border-main animate-scale-up space-y-5 font-sans shadow-2xl">
            {/* Modal Header */}
            <div className="flex justify-between items-center border-b border-border-main/20 pb-3">
              <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
                <Trash2 className="w-5 h-5 text-red-500" />
                Delete Object Storage
              </h3>
              <button
                onClick={() => setShowStorageDeleteConfirm(false)}
                disabled={isDeletingStorage}
                className="text-text-muted hover:text-text-main transition p-1 rounded-lg cursor-pointer"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            <p className="text-sm text-text-muted leading-relaxed">
              Are you sure you want to delete Object Storage <strong className="text-text-main font-mono">"{targetStorageToDelete}"</strong>? This will permanently release the bucket resource and its metadata.
            </p>

            {/* Deletion Confirmation Text Input Box */}
            <div className="space-y-1.5 p-4 bg-bg-panel/50 border border-border-main/40 rounded-xl">
              <label className="block text-xs font-bold text-text-muted">
                To confirm deletion, type <span className="font-mono bg-bg-input px-1.5 py-0.5 rounded text-emerald-600 dark:text-emerald-400 font-extrabold">{targetStorageToDelete}</span> in the box below:
              </label>
              <input
                type="text"
                value={storageDeleteConfirmText}
                onChange={(e) => setStorageDeleteConfirmText(e.target.value)}
                placeholder="Type the object storage name to delete"
                className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm font-bold font-mono focus:outline-none focus:ring-1 focus:ring-red-500"
                disabled={isDeletingStorage}
              />
            </div>

            {/* Force Delete Toggle Option (CSP Console Standard: Checkbox on Left) */}
            <div className="p-4 bg-bg-panel/50 border border-border-main/40 rounded-xl">
              <label className="flex items-start gap-3.5 cursor-pointer">
                <input
                  type="checkbox"
                  checked={forceDeleteStorage}
                  onChange={(e) => setForceDeleteStorage(e.target.checked)}
                  className="w-4.5 h-4.5 accent-red-500 cursor-pointer rounded shrink-0 mt-0.5"
                />
                <div className="text-xs font-mono">
                  <span className="font-extrabold text-text-main block">Force Delete (option=force)</span>
                  <span className="text-text-muted font-sans text-xs block mt-0.5">Forcefully delete bucket along with all stored objects. (Uncheck for standard empty bucket deletion)</span>
                </div>
              </label>
            </div>

            {storageDeleteError && (
              <div className="p-3 bg-red-500/10 border border-red-500/20 text-red-500 text-xs font-semibold rounded-xl">
                {storageDeleteError}
              </div>
            )}

            {/* Footer: Prefer respond-async on Left, Cancel & Confirm Delete Buttons on Right */}
            <div className="border-t border-border-main/40 pt-4 flex flex-col sm:flex-row items-stretch sm:items-center justify-between gap-3 font-sans">
              {/* Left Side: Prefer respond-async Toggle */}
              <label className="flex items-center gap-2 cursor-pointer text-xs font-bold text-text-main p-2 px-3.5 bg-bg-panel/60 hover:bg-bg-input border border-border-main/40 rounded-xl shrink-0 transition">
                <Zap className="w-4 h-4 text-emerald-500 shrink-0" />
                <span className="font-mono">Prefer: respond-async</span>
                <input
                  type="checkbox"
                  checked={preferAsyncStorageDelete}
                  onChange={(e) => setPreferAsyncStorageDelete(e.target.checked)}
                  className="w-4 h-4 accent-emerald-500 cursor-pointer rounded ml-1"
                />
              </label>

              {/* Right Side: Cancel & Confirm Delete Buttons */}
              <div className="flex items-center justify-end space-x-3 shrink-0">
                <button
                  type="button"
                  onClick={() => setShowStorageDeleteConfirm(false)}
                  disabled={isDeletingStorage}
                  className="px-5 py-2.5 bg-bg-input hover:bg-bg-main border border-border-main text-text-main font-bold text-sm rounded-xl transition cursor-pointer"
                >
                  Cancel
                </button>
                <button
                  type="button"
                  onClick={handleConfirmStorageDelete}
                  disabled={isDeletingStorage || storageDeleteConfirmText !== targetStorageToDelete}
                  className={`px-6 py-2.5 rounded-xl text-sm font-extrabold flex items-center gap-2 transition ${
                    isDeletingStorage || storageDeleteConfirmText !== targetStorageToDelete
                      ? 'bg-bg-panel border border-border-main text-text-muted cursor-not-allowed'
                      : 'bg-red-500 hover:bg-red-600 text-white cursor-pointer shadow-lg shadow-red-500/20'
                  }`}
                >
                  {isDeletingStorage && <Loader2 className="w-4 h-4 animate-spin" />}
                  Confirm Delete
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Official CM-Beetle SSH Readiness Modal */}
      {showSshModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm p-4 animate-fade-in">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-3xl border border-border-main animate-scale-up space-y-4 max-h-[85vh] overflow-y-auto">
            <div className="flex justify-between items-center border-b border-border-main/20 pb-3">
              <h3 className="text-base font-extrabold text-text-main flex items-center gap-2">
                <Activity className="w-5 h-5 text-emerald-500" />
                <span>CM-Beetle SSH Readiness Verification ({selectedInfraId})</span>
              </h3>
              <button
                onClick={() => setShowSshModal(false)}
                className="text-text-muted hover:text-text-main transition p-1 rounded-lg cursor-pointer"
              >
                <X className="w-5 h-5" />
              </button>
            </div>

            {isCheckingSsh ? (
              <div className="p-8 text-center space-y-3 bg-bg-input rounded-xl border border-border-main/30">
                <RefreshCw className="w-6 h-6 animate-spin text-emerald-500 mx-auto" />
                <p className="text-sm font-mono text-text-muted">Checking SSH readiness across provisioned Nodes...</p>
              </div>
            ) : sshError ? (
              <div className="p-5 rounded-xl border bg-amber-500/10 border-amber-500/30 space-y-2 text-center">
                <AlertTriangle className="w-8 h-8 text-amber-500 mx-auto" />
                <h4 className="text-sm font-extrabold font-mono text-amber-500">{sshError}</h4>
                {sshCooldownSec > 0 && (
                  <p className="text-xs font-mono text-text-muted">
                    Cooldown active: {sshCooldownSec}s remaining (30s interval enforced by CM-Beetle).
                  </p>
                )}
              </div>
            ) : sshReadyData ? (
              <div className="space-y-4">
                <div className={`p-4 rounded-xl border flex flex-col sm:flex-row sm:items-center justify-between gap-3 ${
                  sshReadyData.ready
                    ? 'bg-emerald-500/10 border-emerald-500/30'
                    : 'bg-amber-500/10 border-amber-500/30'
                }`}>
                  <div className="space-y-1">
                    <div className="flex items-center gap-2">
                      <span className={`px-2.5 py-0.5 rounded-full text-xs font-extrabold ${
                        sshReadyData.ready
                          ? 'bg-emerald-500/20 text-emerald-500'
                          : 'bg-amber-500/20 text-amber-500'
                      }`}>
                        {sshReadyData.ready ? 'ALL NODES SSH READY' : 'PARTIAL OR PENDING SSH'}
                      </span>
                      <span className="text-sm font-extrabold text-text-main font-mono">
                        {sshReadyData.readyNodes} / {sshReadyData.totalNodes} Nodes Ready
                      </span>
                    </div>
                    <p className="text-xs font-mono text-text-muted">{sshReadyData.message}</p>
                  </div>
                  <div className="text-xs font-mono text-text-muted text-right">
                    <div>Checked At:</div>
                    <div className="font-extrabold text-text-main">{sshReadyData.checkedAt}</div>
                  </div>
                </div>

                {sshReadyData.nodeStatus && sshReadyData.nodeStatus.length > 0 && (
                  <div className="overflow-x-auto border border-border-main/50 rounded-xl">
                    <table className="w-full text-left border-collapse text-sm">
                      <thead>
                        <tr className="border-b border-border-main bg-bg-input/60 text-text-muted font-normal">
                          <th className="py-2.5 px-3">Node Name</th>
                          <th className="py-2.5 px-3">Public IP</th>
                          <th className="py-2.5 px-3">Private IP</th>
                          <th className="py-2.5 px-3">SSH User / Port</th>
                          <th className="py-2.5 px-3">SSH Status</th>
                          <th className="py-2.5 px-3">Diagnostic Message</th>
                        </tr>
                      </thead>
                      <tbody className="divide-y divide-border-main/40 font-mono text-xs">
                        {sshReadyData.nodeStatus.map((ns, i) => (
                          <tr key={i} className="hover:bg-bg-input/40">
                            <td className="py-2.5 px-3 font-extrabold text-text-main">{ns.name || ns.id}</td>
                            <td className="py-2.5 px-3 font-extrabold text-emerald-500">{ns.publicIP || '-'}</td>
                            <td className="py-2.5 px-3 text-text-muted">{ns.privateIP || '-'}</td>
                            <td className="py-2.5 px-3 text-text-main">{ns.username || 'cb-user'}:{ns.sshPort || 22}</td>
                            <td className="py-2.5 px-3">
                              <span className={`px-2 py-0.5 rounded font-bold border ${
                                ns.sshReady
                                  ? 'bg-emerald-500/15 text-emerald-500 border-emerald-500/30'
                                  : 'bg-red-500/15 text-red-500 border-red-500/30'
                              }`}>
                                {ns.sshReady ? 'READY' : 'NOT READY'}
                              </span>
                            </td>
                            <td className="py-2.5 px-3 text-text-muted">{ns.error || 'SSH handshake verified'}</td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                )}
              </div>
            ) : null}

            <div className="flex items-center justify-between pt-2 border-t border-border-main/20">
              <span className="text-xs font-mono text-text-muted">
                {sshCooldownSec > 0 ? `Next check available in ${sshCooldownSec}s` : 'Ready to re-check'}
              </span>
              <div className="flex items-center gap-2">
                <button
                  onClick={handleCheckInfraSshReady}
                  disabled={sshCooldownSec > 0 || isCheckingSsh}
                  className="px-4 py-2 bg-emerald-500/15 hover:bg-emerald-500/25 disabled:opacity-40 border border-emerald-500/30 text-emerald-500 rounded-xl text-xs font-bold transition cursor-pointer"
                >
                  {sshCooldownSec > 0 ? `Wait ${sshCooldownSec}s` : 'Re-Check SSH'}
                </button>
                <button
                  onClick={() => setShowSshModal(false)}
                  className="px-5 py-2 bg-bg-panel border border-border-main text-text-main hover:bg-bg-input rounded-xl text-xs font-bold transition cursor-pointer"
                >
                  Close
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

    </div>
  );
};
