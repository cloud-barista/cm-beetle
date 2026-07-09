'use client';

import React, { useState, useEffect } from 'react';
import { useMigrationStore } from '../../store/migrationStore';
import { honeybeeApi } from '../../api/client';
import { Plus, Server, Play, Key, Upload, CheckCircle2, XCircle, RefreshCw, FileText, ChevronRight, Download } from 'lucide-react';

interface ServerRow {
  id: string;
  name: string;
  ip: string;
  port: string;      // empty string defaults to common port
  user: string;      // empty string defaults to common user
  privateKey: string; // empty string defaults to common privateKey
  status: 'idle' | 'checking' | 'success' | 'failed';
  statusMsg?: string;
}

export const SourceCenter: React.FC = () => {
  const {
    sourceGroups,
    activeSgId,
    connections,
    refinedSourceInfra,
    isLoadingSource,
    fetchSourceGroups,
    createSourceGroup,
    registerConnection,
    fetchRefinedInfraByGroup,
    fetchSavedSourceModels,
    saveSourceModel
  } = useMigrationStore();

  // Local UI Control states
  const [showGroupModal, setShowGroupModal] = useState(false);
  const [newGroupName, setNewGroupName] = useState('');
  const [newGroupDesc, setNewGroupDesc] = useState('');

  // 1. Shared Credentials Template (Common SSH Key definitions)
  const [useCommonCred, setUseCommonCred] = useState(true);
  const [commonUser, setCommonUser] = useState('ubuntu');
  const [commonPort, setCommonPort] = useState(22);
  const [commonKey, setCommonKey] = useState('');

  // Spreadsheet-like Grid rows
  const [serverRows, setServerRows] = useState<ServerRow[]>([
    { id: '1', name: 'Web Server 1', ip: '10.0.1.30', port: '', user: '', privateKey: '', status: 'idle' },
    { id: '2', name: 'Database 1', ip: '10.0.1.221', port: '', user: '', privateKey: '', status: 'idle' },
    { id: '3', name: 'Database 2', ip: '10.0.1.138', port: '', user: '', privateKey: '', status: 'idle' }
  ]);

  // Model Save details
  const [modelName, setModelName] = useState('onpremise-web-db-v1');
  const [modelDesc, setModelDesc] = useState('Production server cluster containing 1 Web and 2 InfluxDB DB nodes');
  const [saveSuccess, setSaveSuccess] = useState(false);

  useEffect(() => {
    fetchSourceGroups();
    fetchSavedSourceModels();
  }, []);

  // Sync rows if activeSgId changes or connections loaded
  useEffect(() => {
    if (connections && connections.length > 0) {
      setServerRows(
        connections.map((c, idx) => ({
          id: c.id || String(idx + 1),
          name: c.name,
          ip: c.ip,
          port: c.port ? String(c.port) : '',
          user: c.user || '',
          privateKey: c.privateKey || '',
          status: 'success'
        }))
      );
    } else {
      setServerRows([
        { id: '1', name: 'Web Server 1', ip: '10.0.1.30', port: '', user: '', privateKey: '', status: 'idle' },
        { id: '2', name: 'Database 1', ip: '10.0.1.221', port: '', user: '', privateKey: '', status: 'idle' },
        { id: '3', name: 'Database 2', ip: '10.0.1.138', port: '', user: '', privateKey: '', status: 'idle' }
      ]);
    }
  }, [connections]);

  const handleAddRow = () => {
    const nextId = `row-${Date.now()}-${serverRows.length + 1}`;
    setServerRows([
      ...serverRows,
      {
        id: nextId,
        name: `Server Node ${serverRows.length + 1}`,
        ip: '',
        port: '',
        user: '',
        privateKey: '',
        status: 'idle'
      }
    ]);
  };

  const handleUpdateRow = (id: string, field: keyof ServerRow, value: any) => {
    setServerRows(
      serverRows.map((row) => (row.id === id ? { ...row, [field]: value } : row))
    );
  };

  const handleRemoveRow = (id: string) => {
    if (serverRows.length > 1) {
      setServerRows(serverRows.filter((row) => row.id !== id));
    }
  };

  // Drag and drop SSH PEM Key directly to cell handler
  const handleKeyDrop = (id: string, e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    const file = e.dataTransfer.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = (event) => {
        const text = event.target?.result as string;
        handleUpdateRow(id, 'privateKey', text);
      };
      reader.readAsText(file);
    }
  };

  // Drag & drop key file for Common Credential
  const handleCommonKeyDrop = (e: React.DragEvent<HTMLTextAreaElement>) => {
    e.preventDefault();
    const file = e.dataTransfer.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = (event) => {
        setCommonKey(event.target?.result as string);
      };
      reader.readAsText(file);
    }
  };

  // JSON Import parser
  const handleJsonImport = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    const reader = new FileReader();
    reader.onload = (event) => {
      try {
        const parsed = JSON.parse(event.target?.result as string);
        if (Array.isArray(parsed)) {
          const importedRows: ServerRow[] = parsed.map((item: any, idx: number) => ({
            id: `import-${idx}-${Date.now()}`,
            name: item.name || `Imported Host ${idx + 1}`,
            ip: item.ip || '',
            port: item.port ? String(item.port) : '',
            user: item.user || '',
            privateKey: item.privateKey || '',
            status: 'idle'
          }));
          setServerRows(importedRows);
        } else {
          alert('JSON file must be an array of server configurations.');
        }
      } catch (err) {
        alert('Invalid JSON file format.');
      }
    };
    reader.readAsText(file);
  };

  // Smart Credentials Resolver
  const resolveCredentials = (row: ServerRow) => {
    const resolvedUser = row.user || (useCommonCred ? commonUser : '') || 'ubuntu';
    const resolvedPort = row.port ? Number(row.port) : (useCommonCred ? commonPort : 22);
    const resolvedKey = row.privateKey || (useCommonCred ? commonKey : '');

    return {
      user: resolvedUser,
      port: resolvedPort,
      privateKey: resolvedKey
    };
  };

  // Validate SSH connection one by one
  const handleValidateConnections = async () => {
    const updatedRows = [...serverRows];
    for (let i = 0; i < updatedRows.length; i++) {
      const row = updatedRows[i];
      handleUpdateRow(row.id, 'status', 'checking');

      const creds = resolveCredentials(row);
      const payload = {
        ip: row.ip,
        port: creds.port,
        user: creds.user,
        privateKey: creds.privateKey
      };

      const testResult = await honeybeeApi.testSshConnection(payload);
      if (testResult.success) {
        handleUpdateRow(row.id, 'status', 'success');
      } else {
        handleUpdateRow(row.id, 'status', 'failed');
        handleUpdateRow(row.id, 'statusMsg', testResult.message);
      }
    }
  };

  // Register connections under selected group and fetch refined
  const handleExtractAndSaveModel = async () => {
    let sgId = activeSgId;

    if (!sgId) {
      alert('Please select or create a Source Group first.');
      return;
    }

    try {
      for (const row of serverRows) {
        const creds = resolveCredentials(row);
        await registerConnection({
          name: row.name,
          ip: row.ip,
          port: creds.port,
          user: creds.user,
          privateKey: creds.privateKey,
          description: `Server registration under group ${sgId}`
        });
      }

      await fetchRefinedInfraByGroup(sgId);
    } catch (err) {
      console.error('Failed to extract infra info:', err);
    }
  };

  const handleSaveToDamselfly = async () => {
    try {
      await saveSourceModel(modelName, modelDesc);
      setSaveSuccess(true);
      setTimeout(() => setSaveSuccess(false), 3000);
    } catch (err) {
      console.error('Failed to save to Damselfly:', err);
    }
  };

  // Download template import JSON
  const handleDownloadTemplate = () => {
    const templateData = [
      {
        name: "Web Server 1",
        ip: "10.0.1.30",
        port: 22,
        user: "ubuntu",
        privateKey: ""
      },
      {
        name: "Database 1",
        ip: "10.0.1.221",
        port: 22,
        user: "root",
        privateKey: "-----BEGIN RSA PRIVATE KEY-----\n...\n-----END RSA PRIVATE KEY-----"
      }
    ];
    const blob = new Blob([JSON.stringify(templateData, null, 2)], { type: 'application/json' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'source_server_template.json';
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  };

  return (
    <div className="space-y-6">

      {/* 1. Group Selector & Credential Library */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">

        {/* Source Infrastructure Control Card */}
        <div className="glass-panel p-6 rounded-2xl flex flex-col justify-between">
          <div>
            <div className="flex items-center justify-between mb-4">
              <h2 className="text-base font-bold text-text-main flex items-center">
                <Server className="w-5 h-5 text-cyan-400 mr-2" />
                1. Source Infrastructures
              </h2>
              <button
                onClick={() => setShowGroupModal(true)}
                className="px-3 py-2 bg-bg-panel border border-cyan-500/40 hover:bg-cyan-500/10 hover:border-cyan-500/30 rounded-lg text-sm font-bold flex items-center transition cursor-pointer text-cyan-600 dark:text-cyan-400"
              >
                <Plus className="w-4 h-4 mr-1" /> New Infra Group
              </button>
            </div>
            <p className="text-sm text-text-muted mb-4">
              Select or define a source infrastructure context before registering server hosts.
            </p>
            <select
              value={activeSgId}
              onChange={(e) => {
                useMigrationStore.setState({ activeSgId: e.target.value });
                if (e.target.value) {
                  honeybeeApi.getConnectionInfoList(e.target.value).then((conns) => {
                    useMigrationStore.setState({ connections: conns });
                  });
                }
              }}
              className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm focus:outline-none focus:ring-1 focus:ring-cyan-500"
            >
              <option value="">-- Select Source Infrastructure --</option>
              {sourceGroups.map((g) => (
                <option key={g.id} value={g.id}>
                  {g.name}
                </option>
              ))}
            </select>
          </div>

          <div className="mt-6 pt-4 border-t border-border-main flex justify-between text-sm text-text-muted">
            <span>Infra catalogs: {sourceGroups.length}</span>
            <span>Current rows: {serverRows.length}</span>
          </div>
        </div>

        {/* 2. Shared Credentials Card */}
        <div className="glass-panel p-6 rounded-2xl lg:col-span-2">
          <div className="flex items-center justify-between mb-2">
            <h2 className="text-base font-bold text-text-main flex items-center">
              <Key className="w-5 h-5 text-purple-400 mr-2" />
              2. Shared Credentials
            </h2>
            <label className="inline-flex items-center cursor-pointer">
              <input
                type="checkbox"
                checked={useCommonCred}
                onChange={(e) => setUseCommonCred(e.target.checked)}
                className="sr-only peer"
              />
              <div className="relative w-9 h-5 bg-bg-input peer-focus:outline-none rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-slate-400 after:border-slate-300 after:border after:rounded-full after:h-4 after:w-4 after:transition-all peer-checked:bg-purple-600" />
              <span className="ms-2 text-sm text-text-muted">Enable Auto Inherit</span>
            </label>
          </div>
          <p className="text-sm text-text-muted mb-4">
            If servers share a common SSH key/password, register it below. Empty grid fields will automatically inherit these settings.
          </p>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
            <div>
              <label className="block text-sm font-semibold text-text-muted mb-1.5">Common SSH Username</label>
              <input
                type="text"
                disabled={!useCommonCred}
                value={commonUser}
                onChange={(e) => setCommonUser(e.target.value)}
                className="w-full bg-bg-input border border-border-main disabled:opacity-40 text-text-main rounded-xl px-3.5 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-purple-500"
              />
            </div>
            <div>
              <label className="block text-sm font-semibold text-text-muted mb-1.5">Common SSH Port</label>
              <input
                type="number"
                disabled={!useCommonCred}
                value={commonPort}
                onChange={(e) => setCommonPort(Number(e.target.value))}
                className="w-full bg-bg-input border border-border-main disabled:opacity-40 text-text-main rounded-xl px-3.5 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-purple-500"
              />
            </div>
            <div>
              <label className="block text-sm font-semibold text-text-muted mb-1.5">Common Private Key (PEM format)</label>
              <textarea
                disabled={!useCommonCred}
                onDragOver={(e) => e.preventDefault()}
                onDrop={handleCommonKeyDrop}
                placeholder="Paste PEM Key or Drag file here"
                value={commonKey}
                onChange={(e) => setCommonKey(e.target.value)}
                rows={1}
                className="w-full bg-bg-input border border-border-main disabled:opacity-40 text-text-main rounded-xl px-3.5 py-2 text-sm font-mono focus:outline-none focus:ring-1 focus:ring-purple-500 resize-none overflow-hidden"
              />
            </div>
          </div>
        </div>

      </div>

      {/* 2. Connection Grid Table */}
      <div className="glass-panel rounded-2xl overflow-hidden">
        <div className="px-6 py-5 border-b border-border-main flex justify-between items-center bg-bg-input/20">
          <div>
            <h2 className="text-base font-bold text-text-main">Source Server Connections</h2>
            <p className="text-sm text-text-muted mt-1">Spreadsheet list: fill details. Empty cells automatically inherit the Shared template key.</p>
          </div>
          <div className="flex items-center space-x-3">
            <button
              onClick={handleDownloadTemplate}
              className="px-4 py-2.5 bg-bg-panel border border-cyan-500/40 hover:bg-cyan-500/10 hover:border-cyan-500/30 rounded-xl text-sm font-bold transition flex items-center cursor-pointer text-cyan-600 dark:text-cyan-400"
            >
              <Download className="w-4 h-4 mr-1.5" />
              <span>Get JSON Template</span>
            </button>
            <label className="px-4 py-2.5 bg-bg-panel border border-cyan-500/40 hover:bg-cyan-500/10 hover:border-cyan-500/30 rounded-xl text-sm font-bold transition flex items-center cursor-pointer text-cyan-600 dark:text-cyan-400">
              <Upload className="w-4 h-4 mr-1.5" />
              <span>Import JSON File</span>
              <input
                type="file"
                accept=".json"
                onChange={handleJsonImport}
                className="hidden"
              />
            </label>
            <button
              onClick={handleAddRow}
              className="px-4 py-2.5 bg-bg-panel border border-cyan-500/40 hover:bg-cyan-500/10 hover:border-cyan-500/30 rounded-xl text-sm font-bold transition cursor-pointer text-cyan-600 dark:text-cyan-400"
            >
              + Add Server Row
            </button>
            <button
              onClick={handleValidateConnections}
              className="px-4 py-2.5 bg-purple-600 hover:bg-purple-700 text-white rounded-xl text-sm font-extrabold transition shadow-md shadow-purple-500/10 flex items-center cursor-pointer"
            >
              <Play className="w-4 h-4 mr-1.5" /> Validate Connections
            </button>
          </div>
        </div>

        <div className="overflow-x-auto">
          <table className="w-full text-left border-collapse text-sm">
            <thead>
              <tr className="border-b border-border-main bg-bg-input/40 text-text-muted uppercase tracking-wider font-bold">
                <th className="py-3.5 px-4 w-12 text-center">#</th>
                <th className="py-3.5 px-4 w-60">Connection Name</th>
                <th className="py-3.5 px-4 w-52">IP Address / Host</th>
                <th className="py-3.5 px-4 w-28">SSH Port</th>
                <th className="py-3.5 px-4 w-36">SSH Username</th>
                <th className="py-3.5 px-4">SSH Private Key (PEM format)</th>
                <th className="py-3.5 px-4 text-center w-36">Status</th>
                <th className="py-3.5 px-4 text-center w-24">Actions</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-border-main">
              {serverRows.map((row, idx) => (
                <tr key={row.id} className="hover:bg-cyan-500/[0.02] transition">
                  <td className="py-3.5 px-4 text-center text-text-muted">{idx + 1}</td>

                  {/* Connection Name */}
                  <td className="py-3.5 px-4">
                    <input
                      type="text"
                      value={row.name}
                      onChange={(e) => handleUpdateRow(row.id, 'name', e.target.value)}
                      className="w-full bg-bg-input border border-border-main rounded-lg px-2.5 py-1.5 focus:outline-none focus:border-cyan-500/50 text-text-main text-sm"
                    />
                  </td>

                  {/* IP Address */}
                  <td className="py-3.5 px-4">
                    <input
                      type="text"
                      placeholder="e.g. 10.0.1.30"
                      value={row.ip}
                      onChange={(e) => handleUpdateRow(row.id, 'ip', e.target.value)}
                      className="w-full bg-bg-input border border-border-main rounded-lg px-2.5 py-1.5 focus:outline-none focus:border-cyan-500/50 text-text-main text-sm"
                    />
                  </td>

                  {/* SSH Port */}
                  <td className="py-3.5 px-4">
                    <input
                      type="number"
                      placeholder={useCommonCred ? String(commonPort) : '22'}
                      value={row.port}
                      onChange={(e) => handleUpdateRow(row.id, 'port', e.target.value)}
                      className="w-full bg-bg-input border border-border-main rounded-lg px-2.5 py-1.5 focus:outline-none focus:border-cyan-500/50 text-text-main text-sm placeholder-purple-500/50"
                    />
                  </td>

                  {/* SSH Username */}
                  <td className="py-3.5 px-4">
                    <input
                      type="text"
                      placeholder={useCommonCred ? commonUser : 'ubuntu'}
                      value={row.user}
                      onChange={(e) => handleUpdateRow(row.id, 'user', e.target.value)}
                      className="w-full bg-bg-input border border-border-main rounded-lg px-2.5 py-1.5 focus:outline-none focus:border-cyan-500/50 text-text-main text-sm placeholder-purple-500/50"
                    />
                  </td>

                  {/* SSH Key */}
                  <td className="py-3.5 px-4">
                    <div
                      onDragOver={(e) => e.preventDefault()}
                      onDrop={(e) => handleKeyDrop(row.id, e)}
                      className="border border-dashed border-border-main rounded-lg p-1.5 flex items-center justify-between bg-bg-input text-sm"
                    >
                      <input
                        type="password"
                        placeholder={
                          row.privateKey
                            ? 'Individual Key Loaded'
                            : (useCommonCred && commonKey ? '🔗 Inheriting Shared Key template' : 'Drop PEM or paste Key here')
                        }
                        value={row.privateKey}
                        onChange={(e) => handleUpdateRow(row.id, 'privateKey', e.target.value)}
                        className={`w-full bg-transparent border-0 outline-none text-text-main text-sm mr-2 truncate ${!row.privateKey && useCommonCred && commonKey ? 'text-purple-400/70 font-semibold' : ''
                          }`}
                      />
                      <div className="flex items-center space-x-1 flex-shrink-0 text-xs text-text-muted bg-bg-panel px-2.5 py-1 rounded cursor-pointer">
                        <Upload className="w-3.5 h-3.5" />
                        <span>Drop</span>
                      </div>
                    </div>
                  </td>

                  {/* Status Badge */}
                  <td className="py-3.5 px-4 text-center">
                    {row.status === 'idle' && (
                      <span className="px-3 py-1.5 bg-bg-input text-text-muted rounded-full font-bold border border-border-main text-xs">Idle</span>
                    )}
                    {row.status === 'checking' && (
                      <span className="px-3 py-1.5 bg-cyan-100 dark:bg-cyan-950/40 text-cyan-600 dark:text-cyan-400 border border-cyan-300 dark:border-cyan-800/40 rounded-full font-bold flex items-center justify-center space-x-1 text-xs">
                        <RefreshCw className="w-3.5 h-3.5 animate-spin" />
                        <span>Checking...</span>
                      </span>
                    )}
                    {row.status === 'success' && (
                      <span className="px-3 py-1.5 bg-emerald-100 dark:bg-emerald-950/40 text-emerald-600 dark:text-emerald-400 border border-emerald-300 dark:border-emerald-800/40 rounded-full font-bold flex items-center justify-center space-x-1 text-xs">
                        <CheckCircle2 className="w-3.5 h-3.5" />
                        <span>Verified</span>
                      </span>
                    )}
                    {row.status === 'failed' && (
                      <span
                        title={row.statusMsg}
                        className="px-3 py-1.5 bg-red-100 dark:bg-red-950/40 text-red-600 dark:text-red-400 border border-red-300 dark:border-red-800/40 rounded-full font-bold flex items-center justify-center space-x-1 cursor-help text-xs"
                      >
                        <XCircle className="w-3.5 h-3.5" />
                        <span>Failed</span>
                      </span>
                    )}
                  </td>

                  {/* Actions */}
                  <td className="py-3.5 px-4 text-center">
                    <button
                      onClick={() => handleRemoveRow(row.id)}
                      className="text-red-600 hover:text-red-600 dark:text-red-400 dark:hover:text-red-300 font-bold hover:underline cursor-pointer text-sm"
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        <div className="px-6 py-4 bg-bg-input/20 border-t border-border-main flex justify-end">
          <button
            onClick={handleExtractAndSaveModel}
            disabled={isLoadingSource}
            className="px-6 py-3.5 bg-gradient-to-r from-cyan-500 to-sky-600 hover:from-cyan-600 hover:to-sky-700 disabled:opacity-50 text-slate-950 rounded-xl text-sm font-extrabold transition shadow-lg shadow-cyan-500/15 flex items-center cursor-pointer"
          >
            {isLoadingSource ? (
              <>
                <RefreshCw className="w-4 h-4 mr-2 animate-spin" />
                Extracting specs...
              </>
            ) : (
              <>
                Scan & Extract Network Info
                <ChevronRight className="w-4 h-4 ml-1.5" />
              </>
            )}
          </button>
        </div>
      </div>

      {/* Extracted Model & Save Form */}
      {refinedSourceInfra && (
        <div className="glass-panel p-6 rounded-2xl animate-fade-in grid grid-cols-1 lg:grid-cols-3 gap-6">
          <div className="lg:col-span-2">
            <h3 className="text-sm font-bold text-text-main mb-3 flex items-center">
              <FileText className="w-4 h-4 text-cyan-400 mr-2" />
              Extracted Source Infrastructure Model Spec
            </h3>
            <div className="bg-bg-input border border-border-main rounded-xl p-4 text-xs font-mono text-slate-800 dark:text-cyan-400 max-h-80 overflow-y-auto">
              <pre>{JSON.stringify(refinedSourceInfra, null, 2)}</pre>
            </div>
          </div>

          <div className="flex flex-col justify-between">
            <div className="space-y-4">
              <h3 className="text-sm font-bold text-text-main">
                Save as On-Premise User Model
              </h3>
              <p className="text-xs text-text-muted leading-relaxed">
                Save the extracted computing specifications to Damselfly server to use in the Simulation / Recommendation phase.
              </p>

              <div>
                <label className="block text-xs font-semibold text-text-muted mb-1.5">Model Name</label>
                <input
                  type="text"
                  value={modelName}
                  onChange={(e) => setModelName(e.target.value)}
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2 text-xs focus:outline-none focus:ring-1 focus:ring-cyan-500"
                />
              </div>

              <div>
                <label className="block text-xs font-semibold text-text-muted mb-1.5">Description</label>
                <textarea
                  value={modelDesc}
                  onChange={(e) => setModelDesc(e.target.value)}
                  rows={3}
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2 text-xs focus:outline-none focus:ring-1 focus:ring-cyan-500 resize-none"
                />
              </div>
            </div>

            <div className="mt-6 pt-4 border-t border-border-main space-y-3">
              {saveSuccess && (
                <div className="p-3 bg-green-500/10 border border-green-500/20 text-green-400 rounded-xl text-xs text-center font-medium animate-fade-in">
                  Model saved to Damselfly Repository successfully.
                </div>
              )}
              <button
                onClick={handleSaveToDamselfly}
                className="w-full py-3 bg-gradient-to-r from-emerald-500 to-teal-600 hover:from-emerald-600 hover:to-teal-700 text-slate-950 font-bold rounded-xl text-xs transition shadow-md shadow-emerald-500/10 cursor-pointer"
              >
                Save Source Model to Damselfly
              </button>
            </div>
          </div>
        </div>
      )}

      {/* New Source Infrastructure Dialog */}
      {showGroupModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/80 backdrop-blur-sm">
          <div className="glass-panel p-6 rounded-2xl w-full max-w-md border border-border-main animate-scale-up">
            <h3 className="text-base font-bold text-text-main mb-4">Register New Source Infrastructure</h3>

            <div className="space-y-4">
              <div>
                <label className="block text-sm font-semibold text-text-muted mb-1.5">Infrastructure Name</label>
                <input
                  type="text"
                  placeholder="e.g. source-infra-1"
                  value={newGroupName}
                  onChange={(e) => setNewGroupName(e.target.value)}
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2.5 text-sm focus:outline-none"
                />
              </div>
              <div>
                <label className="block text-sm font-semibold text-text-muted mb-1.5">Description</label>
                <textarea
                  placeholder="Infrastructure details..."
                  value={newGroupDesc}
                  onChange={(e) => setNewGroupDesc(e.target.value)}
                  rows={3}
                  className="w-full bg-bg-input border border-border-main text-text-main rounded-xl px-4 py-2 text-sm focus:outline-none resize-none"
                />
              </div>
            </div>

            <div className="mt-6 flex justify-end space-x-3 text-sm font-bold">
              <button
                onClick={() => setShowGroupModal(false)}
                className="px-4 py-2.5 bg-bg-input border border-border-main hover:bg-bg-panel rounded-xl cursor-pointer"
              >
                Cancel
              </button>
              <button
                onClick={async () => {
                  if (newGroupName) {
                    await createSourceGroup(newGroupName, newGroupDesc);
                    setShowGroupModal(false);
                    setNewGroupName('');
                    setNewGroupDesc('');
                  }
                }}
                className="px-4 py-2.5 bg-cyan-500 text-slate-950 hover:bg-cyan-600 rounded-xl cursor-pointer"
              >
                Create
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
