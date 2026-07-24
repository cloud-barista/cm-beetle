'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '../../store/migrationStore';
import { TopologyMap } from '../design/TopologyMap';
import { damselflyApi, beetleApi, tumblebugApi } from '../../api/client';
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
  Sliders
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

export const MigratedInfraManagement: React.FC = () => {
  const {
    jobs,
    savedCloudModels,
    selectedCloudModel,
    fetchSavedCloudModels,
    selectCloudModel,
    namespaceId,
    liveReportHtml,
    fetchMigrationReport,
    deletingInfrasMap,
    startInfraTeardown,
    removeInfraTeardown,
    pollInfraTeardownStatus
  } = useMigrationStore();

  const [copiedIp, setCopiedIp] = useState<string | null>(null);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [deleteConfirmText, setDeleteConfirmText] = useState('');
  const [isDeleting, setIsDeleting] = useState(false);
  const [deleteError, setDeleteError] = useState('');
  const [deleteSuccessMsg, setDeleteSuccessMsg] = useState('');

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
      // Ensure SAMPLE_INFRA_ID is always present in catalog list
      const combinedIds = Array.from(new Set([SAMPLE_INFRA_ID, ...ids]));
      setMigratedInfraIds(combinedIds);

      // Auto-select first ID if available
      if (combinedIds.length > 0 && !selectedInfraId) {
        setSelectedInfraId(combinedIds[0]);
      }
    } catch (err) {
      console.warn('Failed to load migrated infra IDs, falling back to sample:', err);
      setMigratedInfraIds([SAMPLE_INFRA_ID]);
      if (!selectedInfraId) {
        setSelectedInfraId(SAMPLE_INFRA_ID);
      }
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
          } else {
            setDeleteError(`Infrastructure "${infraId}" deletion failed: ${res.error || 'Unknown error'}`);
          }
          // Refresh list from server after teardown completes!
          await fetchMigratedInfraIds();
          if (selectedInfraId === infraId) {
            setSelectedInfraId('');
            setLoadedInfraDetail(null);
          }
        }
      }
    }, 3000);

    return () => clearInterval(intervalId);
  }, [deletingInfrasMap, namespaceId, selectedInfraId]);

  const [infraReportHtml, setInfraReportHtml] = useState<string>('');

  const handleLoadInfraDetail = async (targetId?: string) => {
    const infraIdToLoad = targetId || selectedInfraId;
    if (!infraIdToLoad) return;
    setSelectedInfraId(infraIdToLoad);
    setIsLoadingDetail(true);
    try {
      if (infraIdToLoad === SAMPLE_INFRA_ID) {
        setLoadedInfraDetail(SAMPLE_INFRA_DETAIL);
        setIsLoadingDetail(false);
        return;
      }
      const ns = namespaceId || 'mig01';
      const detail = await beetleApi.getMigratedInfraDetail(ns, infraIdToLoad);
      setLoadedInfraDetail(detail || SAMPLE_INFRA_DETAIL);

      // Fetch compliance report HTML as well
      try {
        const rHtml = await beetleApi.getMigrationReport(ns, infraIdToLoad);
        setInfraReportHtml(rHtml);
      } catch (rErr) {
        console.warn('Report fetch notice:', rErr);
      }
    } catch (err) {
      console.warn('Failed to load infra detail for', infraIdToLoad, err);
      if (infraIdToLoad === SAMPLE_INFRA_ID) {
        setLoadedInfraDetail(SAMPLE_INFRA_DETAIL);
      }
    } finally {
      setIsLoadingDetail(false);
    }
  };

  const handleOpenDeleteModalForInfra = (infraId: string) => {
    if (infraId === SAMPLE_INFRA_ID) {
      alert('Sample infrastructure is protected and cannot be deleted.');
      return;
    }
    setSelectedInfraId(infraId);
    setShowDeleteConfirm(true);
    setDeleteConfirmText('');
    setDeleteError('');
  };

  useEffect(() => {
    fetchMigratedInfraIds();
    loadMigratedStorages();
  }, [namespaceId]);

  // Combine backend migrated infras, completed migration jobs from store, and saved models
  const completedJobs = jobs.filter(j => j.status === 'Success');

  const handleCopySshCommand = (ip: string, keyName: string) => {
    if (!ip || ip === 'N/A') return;
    const cmd = `ssh -i ~/.ssh/${keyName}.pem ubuntu@${ip}`;
    navigator.clipboard.writeText(cmd);
    setCopiedIp(ip);
    setTimeout(() => setCopiedIp(null), 2000);
  };

  const handleDownloadReport = () => {
    if (!liveReportHtml) return;
    const blob = new Blob([liveReportHtml], { type: 'text/html' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `migration_report_${selectedInfraId || 'infra'}.html`;
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
      // Execute Delete infrastructure API request (DELETE /beetle/migration/ns/{nsId}/infra/{infraId}?option=terminate) with Prefer: respond-async
      const res = await beetleApi.deleteMigratedInfra(ns, targetInfraId, 'terminate', true);

      if (res.success) {
        setShowDeleteConfirm(false);
        setDeleteConfirmText('');

        if (res.reqId) {
          startInfraTeardown(targetInfraId, res.reqId);
          setDeleteSuccessMsg(`Target infrastructure "${targetInfraId}" termination initiated asynchronously (Req ID: ${res.reqId}).`);
        } else {
          setDeleteSuccessMsg(`Target infrastructure "${targetInfraId}" deleted successfully.`);
          await fetchMigratedInfraIds();
          if (selectedInfraId === targetInfraId) {
            setSelectedInfraId('');
            setLoadedInfraDetail(null);
          }
        }
      } else {
        setDeleteError(res.error || 'Failed to delete target infrastructure.');
      }
    } catch (err: any) {
      console.error('Delete failed:', err);
      setDeleteError(err.response?.data?.error || err.message || 'Failed to delete target infrastructure.');
    } finally {
      setIsDeleting(false);
    }
  };

  // Active infra detail from loadedInfraDetail, completed session jobs, or savedCloudModels
  const activeJobInfra = completedJobs.find(j => j.id === selectedInfraId || j.infraId === selectedInfraId);
  const activeSavedModel = savedCloudModels.find(m => m.id === selectedInfraId || m.name === selectedInfraId);

  // Extract real VNet, Subnet, Security Groups, SSH Key from loadedInfraDetail
  const liveNodes = loadedInfraDetail?.node || [];
  const firstNode = liveNodes[0];

  const extractedVNetName = firstNode?.vNetId || 'vnet-migrated';
  const extractedSubnetName = firstNode?.subnetId || 'subnet-default';
  const extractedSshKeyName = firstNode?.sshKeyId || `key-${loadedInfraDetail?.name || selectedInfraId || 'mig01'}`;
  
  // Extract unique security group IDs across all nodes
  const extractedSgNames: string[] = Array.from(
    new Set(
      liveNodes.flatMap((n: any) => n.securityGroupIds || []).filter(Boolean)
    )
  );

  const targetSecurityGroupList = extractedSgNames.length > 0
    ? extractedSgNames.map((sgName: string) => ({
        name: sgName,
        firewallRules: [{ protocol: 'tcp', dstPorts: '22, 80, 443' }]
      }))
    : [{ name: 'sg-default', firewallRules: [{ protocol: 'tcp', dstPorts: '22, 80, 443' }] }];

  // Normalize model data for TopologyMap and VM Access table
  const currentCloudModel = activeSavedModel?.cloudInfraModel || {
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
      subnetInfoList: [{ name: extractedSubnetName, cidrBlock: '10.0.1.0/24' }]
    },
    targetSshKey: {
      name: extractedSshKeyName
    },
    targetSecurityGroupList,
    targetInfra: {
      name: loadedInfraDetail?.name || loadedInfraDetail?.id || activeJobInfra?.infraId || selectedInfraId || 'migrated-infra',
      nodeGroups: liveNodes.length > 0
        ? Array.from(new Set(liveNodes.map((n: any) => n.nodeGroupId || n.name || 'default-group'))).map((ngId: any) => {
            const ngNodes = liveNodes.filter((n: any) => (n.nodeGroupId || n.name || 'default-group') === ngId);
            const firstNode = ngNodes[0] || {};
            const specName = firstNode.spec?.cspSpecName || firstNode.cspSpecName || firstNode.specId || 'c5.large';
            const vcpu = firstNode.spec?.vCPU || firstNode.vCPU || (specName.includes('small') ? 2 : specName.includes('xlarge') ? 4 : 2);
            const memGiB = firstNode.spec?.memoryGiB || firstNode.memoryGiB || (specName.includes('xlarge') ? 16 : 4);
            const imageName = firstNode.cspImageName || firstNode.imageId || 'ubuntu-22.04';
            const diskSize = firstNode.rootDiskSize || 30;

            return {
              name: ngId,
              specId: specName,
              vCPU: vcpu,
              memoryGiB: memGiB,
              nodeGroupSize: ngNodes.length,
              imageId: imageName,
              rootDiskSize: diskSize,
              securityGroupIds: firstNode.securityGroupIds || extractedSgNames
            };
          })
        : activeJobInfra?.vms
        ? [{ name: 'worker-group', specId: activeJobInfra.vms[0]?.specId || 'c5.large', vCPU: 2, memoryGiB: 4, nodeGroupSize: activeJobInfra.vms.length, imageId: 'ubuntu-22.04', rootDiskSize: 30, securityGroupIds: [] }]
        : [{ name: 'default-group', specId: 'c5.large', vCPU: 2, memoryGiB: 4, nodeGroupSize: 2, imageId: 'ubuntu-22.04', rootDiskSize: 30, securityGroupIds: [] }]
    }
  };

  // Extract VM nodes array for VM Access table
  const activeNodes: { name: string; specId: string; publicIp: string; privateIp: string }[] = loadedInfraDetail?.node && loadedInfraDetail.node.length > 0
    ? loadedInfraDetail.node.map((n: any) => ({
        name: n.name || n.id || 'node',
        specId: n.specId || n.spec || 'c5.large',
        publicIp: n.publicIP || n.publicIp || 'N/A',
        privateIp: n.privateIP || n.privateIp || 'N/A'
      }))
    : activeJobInfra?.vms && activeJobInfra.vms.length > 0
    ? activeJobInfra.vms
    : currentCloudModel.targetInfra?.nodeGroups.map((ng: any, idx: number) => ({
        name: `${ng.name}-node-${idx + 1}`,
        specId: ng.specId,
        publicIp: `54.180.${10 + idx}.${25 + idx}`,
        privateIp: `10.0.1.${100 + idx}`
      }));

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
                const matchingJob = completedJobs.find(j => j.infraId === infraId || j.id === infraId);
                const csp = matchingJob?.csp || (infraId.toLowerCase().includes('aws') ? 'AWS' : infraId.toLowerCase().includes('azure') ? 'AZURE' : 'GCP');
                const region = matchingJob?.region || 'ap-northeast-2';

                return (
                  <div
                    key={infraId}
                    className={`p-4 rounded-xl border text-left transition-all duration-200 flex flex-col justify-between space-y-3 relative overflow-hidden ${
                      isTerminating
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
                        disabled={isTerminating || (isLoadingDetail && selectedInfraId === infraId)}
                        className={`flex-1 py-1.5 px-2.5 rounded-lg text-xs font-bold flex items-center justify-center gap-1.5 transition cursor-pointer font-mono ${
                          isTerminating
                            ? 'bg-bg-input text-text-muted border border-border-main opacity-50 cursor-not-allowed'
                            : isSelected
                            ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
                            : 'bg-bg-input hover:bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-border-main'
                        }`}
                        title={isTerminating ? 'Infrastructure is currently terminating' : `Load detail for ${infraId}`}
                      >
                        {isTerminating ? (
                          <Loader2 className="w-3.5 h-3.5 animate-spin text-amber-500" />
                        ) : isLoadingDetail && selectedInfraId === infraId ? (
                          <Loader2 className="w-3.5 h-3.5 animate-spin" />
                        ) : (
                          <Eye className="w-3.5 h-3.5" />
                        )}
                        <span>{isTerminating ? 'Terminating...' : isSelected ? 'Loaded ✓' : 'Load Detail'}</span>
                      </button>

                      <button
                        onClick={() => handleOpenDeleteModalForInfra(infraId)}
                        disabled={isTerminating || infraId === SAMPLE_INFRA_ID}
                        className="py-1.5 px-2.5 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-lg text-xs font-bold flex items-center justify-center gap-1 transition cursor-pointer font-mono disabled:opacity-30 disabled:cursor-not-allowed"
                        title={infraId === SAMPLE_INFRA_ID ? 'Sample infrastructure cannot be deleted' : isTerminating ? 'Termination already in progress' : `Delete infrastructure ${infraId}`}
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
                    const matchingJob = completedJobs.find(j => j.infraId === infraId || j.id === infraId);
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
                              disabled={isTerminating || (isLoadingDetail && selectedInfraId === infraId)}
                              className={`py-1 px-3 rounded-lg text-xs font-bold flex items-center gap-1.5 transition cursor-pointer ${
                                isTerminating
                                  ? 'bg-bg-input text-text-muted border border-border-main opacity-50 cursor-not-allowed'
                                  : isSelected
                                  ? 'bg-emerald-500 text-slate-950 shadow-sm'
                                  : 'bg-bg-input hover:bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-border-main'
                              }`}
                            >
                              {isTerminating ? (
                                <Loader2 className="w-3.5 h-3.5 animate-spin text-amber-500" />
                              ) : isLoadingDetail && selectedInfraId === infraId ? (
                                <Loader2 className="w-3.5 h-3.5 animate-spin" />
                              ) : (
                                <Eye className="w-3.5 h-3.5" />
                              )}
                              <span>{isTerminating ? 'Terminating...' : isSelected ? 'Loaded ✓' : 'Load Detail'}</span>
                            </button>

                            <button
                              onClick={() => handleOpenDeleteModalForInfra(infraId)}
                              disabled={isTerminating || infraId === SAMPLE_INFRA_ID}
                              className="py-1 px-3 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-lg text-xs font-bold flex items-center gap-1 transition cursor-pointer disabled:opacity-30 disabled:cursor-not-allowed"
                              title={infraId === SAMPLE_INFRA_ID ? 'Sample infrastructure cannot be deleted' : isTerminating ? 'Termination already in progress' : `Delete infrastructure ${infraId}`}
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
                <h4 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                  <Globe className="w-4 h-4 text-teal-400" />
                  Deployed Nodes & SSH Access
                </h4>
                <div className="overflow-x-auto border border-border-main/50 rounded-xl">
                  <table className="w-full text-left border-collapse text-xs font-mono">
                    <thead>
                      <tr className="border-b border-border-main bg-bg-input/60 text-text-muted font-bold">
                        <th className="py-2.5 px-4">VM Node Name</th>
                        <th className="py-2.5 px-4">Instance Spec ID</th>
                        <th className="py-2.5 px-4">Public IP Address</th>
                        <th className="py-2.5 px-4">Private IP Address</th>
                        <th className="py-2.5 px-4">SSH Command</th>
                      </tr>
                    </thead>
                    <tbody className="divide-y divide-border-main/40 text-text-main">
                      {activeNodes.map((node, idx) => {
                        const isCopied = copiedIp === node.publicIp;
                        return (
                          <tr key={idx} className="hover:bg-emerald-500/[0.02] transition">
                            <td className="py-3 px-4 font-bold text-text-main">{node.name}</td>
                            <td className="py-3 px-4 text-emerald-600 dark:text-emerald-400 font-bold">{node.specId}</td>
                            <td className="py-3 px-4 select-all text-text-main font-extrabold">{node.publicIp}</td>
                            <td className="py-3 px-4 text-text-muted">{node.privateIp}</td>
                            <td className="py-3 px-4">
                              <button
                                onClick={() => handleCopySshCommand(node.publicIp, currentCloudModel.targetSshKey?.name || 'key')}
                                disabled={!node.publicIp || node.publicIp === 'N/A'}
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
                    <span className="text-xs font-bold text-text-muted font-mono block uppercase">Virtual Network (VNet)</span>
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
                    <span className="text-xs font-bold text-text-muted font-mono block uppercase">Subnets Allocation</span>
                    <div className="space-y-2 text-xs font-mono">
                      {(currentCloudModel.targetVNet?.subnetInfoList || [{ name: extractedSubnetName, cidrBlock: '10.0.1.0/24' }]).map((sub: any, idx: number) => (
                        <div key={idx} className="p-2.5 bg-bg-panel/60 rounded-lg border border-border-main/40 flex justify-between items-center">
                          <div>
                            <span className="font-extrabold text-text-main block">{sub.name}</span>
                            <span className="text-text-muted text-[11px]">Subnet ID: {sub.name}</span>
                          </div>
                          <span className="px-2 py-0.5 rounded bg-emerald-500/10 border border-emerald-500/30 text-emerald-600 dark:text-emerald-400 font-bold">
                            {sub.cidrBlock}
                          </span>
                        </div>
                      ))}
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
                      <span className="font-bold text-emerald-600 dark:text-emerald-400">ubuntu / root</span>
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
            {activeSubTab === 'nlb' && (
              <div className="space-y-4">
                <h4 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                  <Radio className="w-4 h-4 text-emerald-500" />
                  Managed Network Load Balancers (NLB)
                </h4>

                {(currentCloudModel as any).targetNlbList && (currentCloudModel as any).targetNlbList.length > 0 ? (
                  <div className="overflow-x-auto border border-border-main/50 rounded-xl">
                    <table className="w-full text-left border-collapse text-xs font-mono">
                      <thead>
                        <tr className="border-b border-border-main bg-bg-input/60 text-text-muted font-bold">
                          <th className="py-2.5 px-4">NLB Name</th>
                          <th className="py-2.5 px-4">Listener Port</th>
                          <th className="py-2.5 px-4">Target Protocol</th>
                          <th className="py-2.5 px-4">Health Check</th>
                          <th className="py-2.5 px-4">Target Node Group</th>
                        </tr>
                      </thead>
                      <tbody className="divide-y divide-border-main/40 text-text-main">
                        {(currentCloudModel as any).targetNlbList.map((nlb: any, idx: number) => (
                          <tr key={idx} className="hover:bg-emerald-500/[0.02] transition">
                            <td className="py-3 px-4 font-bold text-emerald-600 dark:text-emerald-400">{nlb.name || 'mig-nlb-01'}</td>
                            <td className="py-3 px-4 font-bold">Port 80 ➔ 8080</td>
                            <td className="py-3 px-4 text-emerald-500 font-extrabold">TCP / HTTP</td>
                            <td className="py-3 px-4 text-text-muted">HTTP /healthz (Interval: 10s)</td>
                            <td className="py-3 px-4 font-bold text-text-main">ng-web-01</td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                ) : (
                  <div className="p-6 bg-bg-panel/40 border border-border-main/40 rounded-xl text-center text-xs font-mono text-text-muted">
                    No Dedicated Managed NLB configured for this infrastructure. Compute VM instances are configured with Direct Public IP access.
                  </div>
                )}
              </div>
            )}
          </div>

          {/* Compliance HTML Report Viewer */}
          {(infraReportHtml || liveReportHtml) && (
            <div className="glass-panel rounded-2xl border border-border-main overflow-hidden flex flex-col">
              <div className="px-5 py-4 bg-bg-input/40 border-b border-border-main flex items-center justify-between">
                <h3 className="text-sm font-bold text-text-main flex items-center gap-2">
                  <Eye className="w-4 h-4 text-emerald-400" />
                  Post-Migration Compliance & Comparison Report
                </h3>
                <button
                  onClick={handleDownloadReport}
                  className="px-3.5 py-1.5 bg-bg-panel border border-border-main hover:bg-emerald-500/10 hover:border-emerald-500/20 text-emerald-600 dark:text-emerald-400 rounded-lg text-xs font-extrabold flex items-center transition cursor-pointer"
                >
                  <Download className="w-3.5 h-3.5 mr-1" />
                  Download HTML Report
                </button>
              </div>
              <div
                className="p-6 bg-bg-input text-text-main max-h-[500px] overflow-y-auto font-sans prose prose-sm max-w-none text-sm"
                dangerouslySetInnerHTML={{ __html: infraReportHtml || liveReportHtml }}
              />
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
                      const connInfo = extractCspAndRegionFromConnection(storage.connectionName);
                      const csp = (storage.csp || storage.targetCloud?.csp || connInfo.csp).toUpperCase();
                      const region = storage.region || storage.targetCloud?.region || connInfo.region;

                      return (
                        <div
                          key={storage.id || idx}
                          onClick={() => setSelectedStorageId(sName)}
                          className={`p-4 rounded-xl border text-left transition-all duration-200 flex flex-col justify-between space-y-3 cursor-pointer relative overflow-hidden ${
                            isSelected
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

                            <span className="text-xs font-bold px-2 py-0.5 rounded-full bg-emerald-100 dark:bg-emerald-950/50 text-emerald-600 dark:text-emerald-400 border border-emerald-300 dark:border-emerald-800/40 flex items-center gap-1 font-mono shrink-0">
                              <span className="w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse" />
                              Available
                            </span>
                          </div>

                          <div className="flex items-center justify-between gap-2 pt-2 border-t border-border-main/20">
                            <button
                              onClick={(e) => {
                                e.stopPropagation();
                                setSelectedStorageId(sName);
                              }}
                              className={`flex-1 py-1.5 px-2.5 rounded-lg text-xs font-bold flex items-center justify-center gap-1.5 transition cursor-pointer font-mono ${
                                isSelected
                                  ? 'bg-emerald-500 text-slate-950 shadow-md shadow-emerald-500/20'
                                  : 'bg-bg-input hover:bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-border-main'
                              }`}
                            >
                              <Eye className="w-3.5 h-3.5" />
                              <span>{isSelected ? 'Loaded ✓' : 'Load Detail'}</span>
                            </button>

                            <button
                              onClick={async (e) => {
                                e.stopPropagation();
                                if (sName === SAMPLE_STORAGE_ID) return;
                                if (!confirm(`Delete object storage '${sName}'?`)) return;
                                try {
                                  await beetleApi.deleteMigratedObjectStorage(namespaceId || 'mig01', sName);
                                  await loadMigratedStorages();
                                } catch (err) {
                                  console.warn('Failed to delete storage', err);
                                }
                              }}
                              disabled={sName === SAMPLE_STORAGE_ID}
                              className="py-1.5 px-2.5 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-lg text-xs font-bold flex items-center justify-center gap-1 transition cursor-pointer font-mono disabled:opacity-30 disabled:cursor-not-allowed"
                              title={sName === SAMPLE_STORAGE_ID ? 'Sample object storage cannot be deleted' : `Delete object storage ${sName}`}
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
                          const connInfo = extractCspAndRegionFromConnection(storage.connectionName);
                          const csp = (storage.csp || storage.targetCloud?.csp || connInfo.csp).toUpperCase();
                          const region = storage.region || storage.targetCloud?.region || connInfo.region;
                          const srcBucket = storage.sourceBucketName || 'datamold-aws-test';

                          return (
                            <tr
                              key={storage.id || idx}
                              onClick={() => setSelectedStorageId(sName)}
                              className={`transition hover:bg-bg-panel/60 cursor-pointer ${
                                isSelected ? 'bg-emerald-500/10 font-bold' : ''
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
                                <span className="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full text-xs font-extrabold bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-emerald-500/20">
                                  <span className="w-1.5 h-1.5 rounded-full bg-emerald-500 animate-pulse"></span>
                                  <span>Available</span>
                                </span>
                              </td>
                              <td className="py-3 px-4 text-right">
                                <div className="flex items-center justify-end gap-2">
                                  <button
                                    onClick={(e) => {
                                      e.stopPropagation();
                                      setSelectedStorageId(sName);
                                    }}
                                    className={`py-1 px-3 rounded-lg text-xs font-bold flex items-center gap-1.5 transition cursor-pointer ${
                                      isSelected
                                        ? 'bg-emerald-500 text-slate-950 shadow-sm'
                                        : 'bg-bg-input hover:bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border border-border-main'
                                    }`}
                                  >
                                    <Eye className="w-3.5 h-3.5" />
                                    <span>{isSelected ? 'Loaded ✓' : 'Load Detail'}</span>
                                  </button>

                                  <button
                                    onClick={async (e) => {
                                      e.stopPropagation();
                                      if (sName === SAMPLE_STORAGE_ID) return;
                                      if (!confirm(`Delete object storage '${sName}'?`)) return;
                                      try {
                                        await beetleApi.deleteMigratedObjectStorage(namespaceId || 'mig01', sName);
                                        await loadMigratedStorages();
                                      } catch (err) {
                                        console.warn('Failed to delete storage', err);
                                      }
                                    }}
                                    disabled={sName === SAMPLE_STORAGE_ID}
                                    className="py-1 px-3 bg-red-500/10 hover:bg-red-500/20 text-red-500 border border-red-500/20 rounded-lg text-xs font-bold transition cursor-pointer disabled:opacity-30 disabled:cursor-not-allowed"
                                    title={sName === SAMPLE_STORAGE_ID ? 'Sample object storage cannot be deleted' : `Delete object storage ${sName}`}
                                  >
                                    <Trash2 className="w-3.5 h-3.5 inline mr-1" />
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

                  {/* Bucket Objects List (Contents) */}
                  {Array.isArray(activeSelectedStorage.contents) && activeSelectedStorage.contents.length > 0 && (
                    <div className="space-y-3 pt-4 border-t border-border-main/30">
                      <div className="flex items-center justify-between">
                        <h4 className="text-sm font-extrabold text-text-main flex items-center gap-2">
                          <Database className="w-4 h-4 text-emerald-500" />
                          <span>Bucket Contents &amp; Objects ({activeSelectedStorage.contents.length})</span>
                        </h4>
                      </div>

                      <div className="overflow-x-auto rounded-xl border border-border-main/50 bg-bg-panel/40">
                        <table className="w-full text-left text-sm font-mono">
                          <thead className="bg-bg-panel/90 text-text-muted font-normal border-b border-border-main/50">
                            <tr>
                              <th className="p-3.5 font-normal">Object Key</th>
                              <th className="p-3.5 font-normal">Size (Bytes)</th>
                              <th className="p-3.5 font-normal">Storage Class</th>
                              <th className="p-3.5 font-normal">ETag</th>
                              <th className="p-3.5 font-normal">Last Modified</th>
                            </tr>
                          </thead>
                          <tbody className="divide-y divide-border-main/20 text-text-main">
                            {activeSelectedStorage.contents.map((obj: any, oIdx: number) => (
                              <tr key={oIdx} className="hover:bg-bg-panel/60 transition">
                                <td className="p-3.5 font-extrabold text-emerald-500">{obj.key}</td>
                                <td className="p-3.5 font-extrabold">{obj.size} B</td>
                                <td className="p-3.5 text-text-muted">{obj.storageClass || 'STANDARD'}</td>
                                <td className="p-3.5 text-text-muted font-mono">{obj.eTag || '-'}</td>
                                <td className="p-3.5 text-text-main">{obj.lastModified || '-'}</td>
                              </tr>
                            ))}
                          </tbody>
                        </table>
                      </div>
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
                Confirm Delete
              </button>
            </div>
          </div>
        </div>
      )}

    </div>
  );
};
