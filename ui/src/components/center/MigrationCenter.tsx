'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '../../store/migrationStore';
import { Play, CheckCircle2, AlertTriangle, RefreshCw, Terminal, Eye, Download, Info, Server } from 'lucide-react';

export const MigrationCenter: React.FC = () => {
  const {
    savedCloudModels,
    selectedCloudModel,
    namespaceId,
    nameSeed,
    isDeploying,
    activeDeploymentId,
    deploymentStatus,
    liveReportHtml,
    fetchSavedCloudModels,
    selectCloudModel,
    startMigration,
    fetchDeploymentStatus,
    fetchMigrationReport
  } = useMigrationStore();

  const [pollingActive, setPollingActive] = useState(false);

  useEffect(() => {
    fetchSavedCloudModels();
  }, []);

  // Poll deployment status if active migration in progress
  useEffect(() => {
    let intervalId: any;
    if (pollingActive || isDeploying) {
      intervalId = setInterval(() => {
        fetchDeploymentStatus();
        fetchMigrationReport();
      }, 5000);
    }
    return () => {
      if (intervalId) clearInterval(intervalId);
    };
  }, [pollingActive, isDeploying]);

  const handleStartDeploy = async () => {
    if (!selectedCloudModel) {
      alert('Please choose a Target Cloud Design model first.');
      return;
    }
    if (!namespaceId) {
      alert('Namespace ID cannot be empty.');
      return;
    }
    setPollingActive(true);
    await startMigration(selectedCloudModel.cloudInfraModel);
  };

  const getStatusBadge = (status: string) => {
    switch (status?.toLowerCase()) {
      case 'completed':
        return (
          <span className="px-3 py-1.5 bg-green-500/10 text-green-400 border border-green-500/20 rounded-full font-bold uppercase text-xs">
            Success
          </span>
        );
      case 'failed':
        return (
          <span className="px-3 py-1.5 bg-red-500/10 text-red-400 border border-red-500/20 rounded-full font-bold uppercase text-xs">
            Failed
          </span>
        );
      case 'deploying':
      case 'processing':
        return (
          <span className="px-3 py-1.5 bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 rounded-full font-bold uppercase text-xs flex items-center justify-center space-x-1 w-28">
            <RefreshCw className="w-3.5 h-3.5 animate-spin" />
            <span>Building</span>
          </span>
        );
      default:
        return (
          <span className="px-3 py-1.5 bg-bg-input text-text-muted rounded-full font-bold uppercase text-xs">
            Idle
          </span>
        );
    }
  };

  const mockBuildLogs = deploymentStatus?.logs || [
    'System: Awaiting migration launch authorization...',
    'Tumblebug: Target namespace and virtual subnet verified.',
  ];

  const handleDownloadReport = () => {
    if (!liveReportHtml) return;
    const blob = new Blob([liveReportHtml], { type: 'text/html' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `migration_report_${activeDeploymentId || 'infra'}.html`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  };

  return (
    <div className="space-y-6">

      {/* Active IP Allocation Grid Panel */}
      {deploymentStatus && (
        <div className="glass-panel p-6 rounded-2xl animate-fade-in space-y-4">
          <div className="flex justify-between items-center">
            <h3 className="text-sm font-bold text-text-main flex items-center">
              <Server className="w-5 h-5 text-emerald-400 mr-2" />
              Active Provisioned Cloud VM Access Points
            </h3>
            {getStatusBadge(deploymentStatus.status)}
          </div>

          <div className="overflow-x-auto border border-border-main rounded-xl">
            <table className="w-full text-left border-collapse text-sm">
              <thead>
                <tr className="border-b border-border-main bg-bg-input/40 text-text-muted font-bold">
                  <th className="py-3 px-4">Cloud Node Group</th>
                  <th className="py-3 px-4">Instance Spec ID</th>
                  <th className="py-3 px-4">Public IP Address</th>
                  <th className="py-3 px-4">Private IP Address</th>
                  <th className="py-3 px-4">Security Key</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-border-main text-text-main">
                {selectedCloudModel?.cloudInfraModel.targetInfra.nodeGroups.map((ng: any, idx: number) => {
                  const vm = deploymentStatus.vms?.[idx] || {};
                  return (
                    <tr key={idx} className="hover:bg-emerald-500/[0.01] transition">
                      <td className="py-3.5 px-4 font-bold text-text-main">{ng.name}</td>
                      <td className="py-3.5 px-4 font-mono text-emerald-400">{ng.specId}</td>
                      <td className="py-3.5 px-4 font-mono select-all text-text-main font-semibold">{vm.publicIp || 'Generating...'}</td>
                      <td className="py-3.5 px-4 font-mono text-text-muted">{vm.privateIp || 'Generating...'}</td>
                      <td className="py-3.5 px-4 font-semibold text-teal-400">
                        {selectedCloudModel.cloudInfraModel.targetSshKey.name}
                      </td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
        </div>
      )}

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">

        {/* Left Control Parameter Settings */}
        <div className="glass-panel p-6 rounded-2xl flex flex-col justify-between space-y-6">
          <div className="space-y-4">
            <h2 className="text-base font-bold text-text-main flex items-center mb-2">
              <Info className="w-5 h-5 text-emerald-400 mr-2" />
              1. Deployment Specifications
            </h2>
            <p className="text-sm text-text-muted mb-4">
              Configure deployment coordinates and namespace variables for target deployment.
            </p>

            <div>
              <label className="block text-sm font-semibold text-text-muted mb-1.5">Load Target Cloud Design</label>
              <select
                value={selectedCloudModel?.id || ''}
                onChange={(e) => {
                  const model = savedCloudModels.find(m => m.id === e.target.value) || null;
                  selectCloudModel(model);
                }}
                className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-1 focus:ring-emerald-500"
              >
                <option value="">-- Choose Recommended Model --</option>
                {savedCloudModels.map((m: any) => (
                  <option key={m.id} value={m.id}>
                    {m.name} ({m.cloudInfraModel.targetCloud.csp.toUpperCase()})
                  </option>
                ))}
              </select>
            </div>

            <div className="grid grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-semibold text-text-muted mb-1.5">Namespace ID</label>
                <input
                  type="text"
                  value={namespaceId}
                  onChange={(e) => useMigrationStore.setState({ namespaceId: e.target.value })}
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500"
                />
              </div>
              <div>
                <label className="block text-sm font-semibold text-text-muted mb-1.5">NameSeed (Prefix)</label>
                <input
                  type="text"
                  value={nameSeed}
                  onChange={(e) => useMigrationStore.setState({ nameSeed: e.target.value })}
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-1 focus:ring-teal-500"
                />
              </div>
            </div>
          </div>

          <div className="pt-4 border-t border-border-main">
            <button
              onClick={handleStartDeploy}
              disabled={isDeploying || !selectedCloudModel}
              className="w-full py-3.5 bg-blue-600 hover:bg-blue-700 disabled:bg-slate-200 dark:disabled:bg-slate-800 disabled:opacity-40 text-white disabled:text-text-muted font-extrabold rounded-xl text-sm tracking-wider transition shadow-lg shadow-blue-500/10 flex items-center justify-center space-x-2 cursor-pointer"
            >
              {isDeploying ? (
                <>
                  <RefreshCw className="w-4 h-4 animate-spin" />
                  <span>Provisioning Infrastructure...</span>
                </>
              ) : (
                <>
                  <Play className="w-4 h-4" />
                  <span>Execute Cloud Deployment</span>
                </>
              )}
            </button>
          </div>
        </div>

        {/* Middle Real-time Build Console Logs */}
        <div className="glass-panel rounded-2xl overflow-hidden flex flex-col h-[400px] lg:col-span-2">
          <div className="px-5 py-4 bg-bg-input/40 border-b border-border-main flex items-center justify-between">
            <h3 className="text-sm font-bold text-text-main flex items-center">
              <Terminal className="w-4 h-4 text-teal-400 mr-2" />
              Live Provisioning logs Console
            </h3>
            {isDeploying && (
              <span className="flex h-2.5 w-2.5 relative">
                <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-teal-400 opacity-75"></span>
                <span className="relative inline-flex rounded-full h-2.5 w-2.5 bg-teal-500"></span>
              </span>
            )}
          </div>
          <div className="flex-1 bg-bg-input p-4 font-mono text-xs text-emerald-400 overflow-y-auto space-y-2 select-text leading-relaxed">
            {mockBuildLogs.map((msg: string, i: number) => (
              <div key={i} className={msg.includes('SUCCESS') ? 'text-green-400 font-semibold' : msg.includes('FAILED') ? 'text-red-400 font-semibold' : 'text-text-muted'}>
                {msg}
              </div>
            ))}
          </div>
          <div className="p-4 bg-bg-input/20 border-t border-border-main flex items-center justify-between text-sm text-text-muted">
            <span>Namespace: {namespaceId || 'None'}</span>
            <span>Target: {selectedCloudModel?.cloudInfraModel.targetCloud.csp.toUpperCase() || 'None'}</span>
          </div>
        </div>

      </div>

      {/* Deployment Status and Saved Spec Details */}
      {deploymentStatus && (
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <div className="glass-panel p-6 rounded-2xl space-y-4 lg:col-span-1">
            <h3 className="text-sm font-bold text-text-main mb-4">Migration Plan Information</h3>

            <div className="space-y-3 text-sm">
              <div className="flex justify-between pb-2 border-b border-border-main">
                <span className="text-text-muted">Infrastructure ID:</span>
                <span className="text-text-main font-bold font-mono">{activeDeploymentId || 'Checking...'}</span>
              </div>
              <div className="flex justify-between pb-2 border-b border-border-main">
                <span className="text-text-muted">Namespace Target:</span>
                <span className="text-text-main font-semibold">{namespaceId}</span>
              </div>
              <div className="flex justify-between pb-2">
                <span className="text-text-muted">Provisioning Status:</span>
                <span>{getStatusBadge(deploymentStatus.status)}</span>
              </div>
            </div>

            {/* Visual Build Success / Fail Banner */}
            {deploymentStatus.status === 'completed' ? (
              <div className="bg-bg-input border border-border-main rounded-xl p-4 flex items-center space-x-3">
                <div className="p-2 bg-green-500/10 text-green-400 rounded-lg">
                  <CheckCircle2 className="w-5 h-5" />
                </div>
                <div>
                  <h4 className="text-xs font-bold text-text-main">Provisioning Complete</h4>
                  <p className="text-[10px] text-text-muted mt-0.5">All VPC subnets and compute instances are active.</p>
                </div>
              </div>
            ) : deploymentStatus.status === 'failed' ? (
              <div className="bg-bg-input border border-border-main rounded-xl p-4 flex items-center space-x-3">
                <div className="p-2 bg-red-500/10 text-red-400 rounded-lg">
                  <AlertTriangle className="w-5 h-5" />
                </div>
                <div>
                  <h4 className="text-xs font-bold text-text-main">Provisioning Failed</h4>
                  <p className="text-[10px] text-text-muted mt-0.5">Build terminated. Check error logs.</p>
                </div>
              </div>
            ) : (
              <div className="bg-bg-input border border-border-main rounded-xl p-4 flex items-center space-x-3">
                <div className="p-2 bg-emerald-500/10 text-emerald-400 rounded-lg animate-pulse">
                  <RefreshCw className="w-5 h-5" />
                </div>
                <div>
                  <h4 className="text-xs font-bold text-text-main">Active Build In Progress</h4>
                  <p className="text-[10px] text-text-muted mt-0.5">Contacting cloud API interfaces. Average: 2-3 mins.</p>
                </div>
              </div>
            )}
          </div>

          {/* Embedded Migration Report HTML Viewer */}
          {liveReportHtml && (
            <div className="glass-panel rounded-2xl overflow-hidden lg:col-span-2 flex flex-col">
              <div className="px-5 py-3.5 bg-bg-input/40 border-b border-border-main flex items-center justify-between">
                <h3 className="text-sm font-bold text-text-main flex items-center">
                  <Eye className="w-4 h-4 text-emerald-400 mr-2" />
                  Post-Migration Compliance & Comparison Report
                </h3>
                <button
                  onClick={handleDownloadReport}
                  className="px-3.5 py-1.5 bg-bg-panel border border-border-main hover:bg-emerald-500/10 hover:border-emerald-500/20 text-emerald-400 rounded-lg text-xs font-semibold flex items-center transition cursor-pointer"
                >
                  <Download className="w-3.5 h-3.5 mr-1" />
                  Download Report
                </button>
              </div>
              <div
                className="p-6 bg-bg-input text-text-main max-h-[500px] overflow-y-auto font-sans prose prose-sm max-w-none text-sm"
                dangerouslySetInnerHTML={{ __html: liveReportHtml }}
              />
            </div>
          )}
        </div>
      )}

    </div>
  );
};
