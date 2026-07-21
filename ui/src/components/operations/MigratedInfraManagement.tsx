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
  Radio
} from 'lucide-react';

const SAMPLE_INFRA_ID = 'sample-aws-infra-01';

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

  // Detailed infra object fetched from GET /beetle/migration/ns/{nsId}/infra/{infraId}
  const [loadedInfraDetail, setLoadedInfraDetail] = useState<any | null>(null);
  const [isLoadingDetail, setIsLoadingDetail] = useState(false);

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

      {/* 2. Deployed Infrastructures (Cards vs Table Toggle) */}
      <div className="glass-panel p-6 rounded-2xl border border-border-main space-y-4">
        <div className="flex flex-wrap justify-between items-center gap-3 border-b border-border-main/20 pb-3">
          <h3 className="text-sm font-extrabold text-text-main flex items-center gap-2">
            <Server className="w-4 h-4 text-emerald-500" />
            Deployed Infrastructures ({migratedInfraIds.length})
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
          )
        ) : (
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
